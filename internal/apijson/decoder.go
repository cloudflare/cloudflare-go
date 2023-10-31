package apijson

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"time"
	"unsafe"

	"github.com/tidwall/gjson"
)

var decoders sync.Map // map[reflect.Type]decoderFunc

func Unmarshal(raw []byte, to any) error {
	d := &decoder{dateFormat: time.RFC3339}
	return d.unmarshal(raw, to)
}

func UnmarshalRoot(raw []byte, to any) error {
	d := &decoder{dateFormat: time.RFC3339, root: true}
	return d.unmarshal(raw, to)
}

type decoder struct {
	root       bool
	dateFormat string
}

type decoderFunc func(node gjson.Result, value reflect.Value) error

type decoderField struct {
	tag    parsedStructTag
	fn     decoderFunc
	idx    []int
	goname string
}

type decoderEntry struct {
	reflect.Type
	dateFormat string
	root       bool
}

func (d *decoder) unmarshal(raw []byte, to any) error {
	value := reflect.ValueOf(to).Elem()
	result := gjson.ParseBytes(raw)
	if !value.IsValid() {
		return fmt.Errorf("apijson: cannot marshal into invalid value")
	}
	return d.typeDecoder(value.Type())(result, value)
}

func (d *decoder) typeDecoder(t reflect.Type) decoderFunc {
	entry := decoderEntry{
		Type:       t,
		dateFormat: d.dateFormat,
		root:       d.root,
	}

	if fi, ok := decoders.Load(entry); ok {
		return fi.(decoderFunc)
	}

	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  decoderFunc
	)
	wg.Add(1)
	fi, loaded := decoders.LoadOrStore(entry, decoderFunc(func(node gjson.Result, v reflect.Value) error {
		wg.Wait()
		return f(node, v)
	}))
	if loaded {
		return fi.(decoderFunc)
	}

	// Compute the real decoder and replace the indirect func with it.
	f = d.newTypeDecoder(t)
	wg.Done()
	decoders.Store(entry, f)
	return f
}

func unmarshalerDecoder(n gjson.Result, v reflect.Value) error {
	return v.Interface().(json.Unmarshaler).UnmarshalJSON([]byte(n.Raw))
}

func (d *decoder) newTypeDecoder(t reflect.Type) decoderFunc {
	if t.ConvertibleTo(reflect.TypeOf(time.Time{})) {
		return d.newTimeTypeDecoder(t)
	}
	if !d.root && t.Implements(reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()) {
		return unmarshalerDecoder
	}
	d.root = false

	if _, ok := unionRegistry[t]; ok {
		return d.newUnionDecoder(t)
	}

	switch t.Kind() {
	case reflect.Pointer:
		inner := t.Elem()
		innerDecoder := d.typeDecoder(inner)

		return func(n gjson.Result, v reflect.Value) error {
			if !v.IsValid() {
				return fmt.Errorf("apijson: unexpected invalid reflection value %+#v", v)
			}

			newValue := reflect.New(inner).Elem()
			err := innerDecoder(n, newValue)
			if err != nil {
				return err
			}

			v.Set(newValue.Addr())
			return nil
		}
	case reflect.Struct:
		return d.newStructTypeDecoder(t)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return d.newArrayTypeDecoder(t)
	case reflect.Map:
		return d.newMapDecoder(t)
	case reflect.Interface:
		return func(node gjson.Result, value reflect.Value) error {
			if !value.IsValid() {
				return fmt.Errorf("apijson: unexpected invalid value %+#v", value)
			}
			if node.Value() != nil {
				value.Set(reflect.ValueOf(node.Value()))
			}
			return nil
		}
	default:
		return d.newPrimitiveTypeDecoder(t)
	}
}

func (d *decoder) newUnionDecoder(t reflect.Type) decoderFunc {
	unionEntry, ok := unionRegistry[t]
	if !ok {
		panic("apijson: couldn't find union of type " + t.String() + " in union registry")
	}
	decoders := []decoderFunc{}
	for _, variant := range unionEntry.variants {
		decoder := d.typeDecoder(variant.Type)
		decoders = append(decoders, decoder)
	}
	return func(n gjson.Result, v reflect.Value) error {
		for idx, variant := range unionEntry.variants {
			decoder := decoders[idx]
			if variant.TypeFilter != n.Type {
				continue
			}
			if len(unionEntry.discriminatorKey) != 0 && n.Get(unionEntry.discriminatorKey).Value() != variant.DiscriminatorValue {
				continue
			}
			inner := reflect.New(variant.Type).Elem()
			err := decoder(n, inner)
			if err != nil {
				return err
			}
			v.Set(inner)
		}
		return errors.New("apijson: was not able to coerce type as union")
	}
}

func (d *decoder) newMapDecoder(t reflect.Type) decoderFunc {
	keyType := t.Key()
	itemType := t.Elem()
	itemDecoder := d.typeDecoder(itemType)

	return func(node gjson.Result, value reflect.Value) (err error) {
		mapValue := reflect.MakeMapWithSize(t, len(node.Map()))

		node.ForEach(func(key, value gjson.Result) bool {
			// It's fine for us to just use `ValueOf` here because the key types will
			// always be primitive types so we don't need to decode it using the standard pattern
			keyValue := reflect.ValueOf(key.Value())
			if !keyValue.IsValid() {
				if err == nil {
					err = fmt.Errorf("apijson: received invalid key type %v", keyValue.String())
				}
				return false
			}
			if keyValue.Type() != keyType {
				if err == nil {
					err = fmt.Errorf("apijson: expected key type %v but got %v", keyType, keyValue.Type())
				}
				return false
			}

			itemValue := reflect.New(itemType).Elem()
			itemerr := itemDecoder(value, itemValue)
			if itemerr != nil {
				if err == nil {
					err = itemerr
				}
				return false
			}

			mapValue.SetMapIndex(keyValue, itemValue)
			return true
		})

		if err != nil {
			return err
		}
		value.Set(mapValue)
		return nil
	}
}

func (d *decoder) newArrayTypeDecoder(t reflect.Type) decoderFunc {
	itemDecoder := d.typeDecoder(t.Elem())

	return func(node gjson.Result, value reflect.Value) (err error) {
		arrayNode := node.Array()

		arrayValue := reflect.MakeSlice(reflect.SliceOf(t.Elem()), len(arrayNode), len(arrayNode))
		for i, itemNode := range arrayNode {
			err = itemDecoder(itemNode, arrayValue.Index(i))
			if err != nil {
				return err
			}
		}

		value.Set(arrayValue)
		return nil
	}
}

func (d *decoder) newStructTypeDecoder(t reflect.Type) decoderFunc {
	// map of json field name to struct field decoders
	decoderFields := map[string]decoderField{}
	extraDecoder := (*decoderField)(nil)
	inlineDecoder := (*decoderField)(nil)

	// This helper allows us to recursively collect field encoders into a flat
	// array. The parameter `index` keeps track of the access patterns necessary
	// to get to some field.
	var collectFieldDecoders func(r reflect.Type, index []int)
	collectFieldDecoders = func(r reflect.Type, index []int) {
		for i := 0; i < r.NumField(); i++ {
			idx := append(index, i)
			field := t.FieldByIndex(idx)
			if !field.IsExported() {
				continue
			}
			// If this is an embedded struct, traverse one level deeper to extract
			// the fields and get their encoders as well.
			if field.Anonymous {
				collectFieldDecoders(field.Type, idx)
				continue
			}
			// If json tag is not present, then we skip, which is intentionally
			// different behavior from the stdlib.
			ptag, ok := parseJSONStructTag(field)
			if !ok {
				continue
			}
			// We only want to support unexported fields if they're tagged with
			// `extras` because that field shouldn't be part of the public API. We
			// also want to only keep the top level extras
			if ptag.extras && len(index) == 0 {
				extraDecoder = &decoderField{ptag, d.typeDecoder(field.Type.Elem()), idx, field.Name}
				continue
			}
			if ptag.inline && len(index) == 0 {
				inlineDecoder = &decoderField{ptag, d.typeDecoder(field.Type), idx, field.Name}
				continue
			}
			if ptag.metadata {
				continue
			}

			oldFormat := d.dateFormat
			dateFormat, ok := parseFormatStructTag(field)
			if ok {
				switch dateFormat {
				case "date-time":
					d.dateFormat = time.RFC3339
				case "date":
					d.dateFormat = "2006-01-02"
				}
			}
			decoderFields[ptag.name] = decoderField{ptag, d.typeDecoder(field.Type), idx, field.Name}
			d.dateFormat = oldFormat
		}
	}
	collectFieldDecoders(t, []int{})

	return func(node gjson.Result, value reflect.Value) (err error) {
		if field := value.FieldByName("JSON"); field.IsValid() {
			if raw := field.FieldByName("raw"); raw.IsValid() {
				setUnexportedField(raw, node.Raw)
			}
		}

		if inlineDecoder != nil {
			var meta Field
			dest := value.FieldByIndex(inlineDecoder.idx)
			isValid := false
			if dest.IsValid() && node.Type != gjson.Null {
				err = inlineDecoder.fn(node, dest)
				if err == nil {
					isValid = true
				}
			}

			if node.Type == gjson.Null {
				meta = Field{
					raw:    node.Raw,
					status: null,
				}
			} else if !isValid {
				meta = Field{
					raw:    node.Raw,
					status: invalid,
				}
			} else if isValid {
				meta = Field{
					raw:    node.Raw,
					status: valid,
				}
			}
			if metadata := getSubField(value, inlineDecoder.idx, inlineDecoder.goname); metadata.IsValid() {
				metadata.Set(reflect.ValueOf(meta))
			}
			return err
		}

		typedExtraType := reflect.Type(nil)
		typedExtraFields := reflect.Value{}
		if extraDecoder != nil {
			typedExtraType = value.FieldByIndex(extraDecoder.idx).Type()
			typedExtraFields = reflect.MakeMap(typedExtraType)
		}
		untypedExtraFields := map[string]Field{}

		for fieldName, itemNode := range node.Map() {
			df, explicit := decoderFields[fieldName]
			var (
				dest reflect.Value
				fn   decoderFunc
				meta Field
			)
			if explicit {
				fn = df.fn
				dest = value.FieldByIndex(df.idx)
			}
			if !explicit && extraDecoder != nil {
				dest = reflect.New(typedExtraType.Elem()).Elem()
				fn = extraDecoder.fn
			}

			isValid := false
			if dest.IsValid() && itemNode.Type != gjson.Null {
				err = fn(itemNode, dest)
				if err == nil {
					isValid = true
				}
			}

			if itemNode.Type == gjson.Null {
				meta = Field{
					raw:    itemNode.Raw,
					status: null,
				}
			} else if !isValid {
				meta = Field{
					raw:    itemNode.Raw,
					status: invalid,
				}
			} else if isValid {
				meta = Field{
					raw:    itemNode.Raw,
					status: valid,
				}
			}

			if explicit {
				if metadata := getSubField(value, df.idx, df.goname); metadata.IsValid() {
					metadata.Set(reflect.ValueOf(meta))
				}
			}
			if !explicit {
				untypedExtraFields[fieldName] = meta
			}
			if !explicit && extraDecoder != nil {
				typedExtraFields.SetMapIndex(reflect.ValueOf(fieldName), dest)
			}
		}

		if extraDecoder != nil && typedExtraFields.Len() > 0 {
			value.FieldByIndex(extraDecoder.idx).Set(typedExtraFields)
		}
		if metadata := getSubField(value, []int{-1}, "Extras"); metadata.IsValid() && len(untypedExtraFields) > 0 {
			metadata.Set(reflect.ValueOf(untypedExtraFields))
		}
		return nil
	}
}

func (d *decoder) newPrimitiveTypeDecoder(t reflect.Type) decoderFunc {
	switch t.Kind() {
	case reflect.String:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetString(n.String())
			if n.Type == gjson.JSON {
				return fmt.Errorf("apijson: failed to parse string")
			}
			return nil
		}
	case reflect.Bool:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetBool(n.Bool())
			if n.Type == gjson.String && (n.Raw != "true" && n.Raw != "false") || n.Type == gjson.JSON {
				return fmt.Errorf("apijson: failed to parse bool")
			}
			return nil
		}
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetInt(n.Int())
			_, err := strconv.ParseFloat(n.Str, 64)
			if n.Type == gjson.JSON || (n.Type == gjson.String && err != nil) {
				return fmt.Errorf("apijson: failed to parse int")
			}
			return nil
		}
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetUint(n.Uint())
			_, err := strconv.ParseFloat(n.Str, 64)
			if n.Type == gjson.JSON || (n.Type == gjson.String && err != nil) {
				return fmt.Errorf("apijson: failed to parse uint")
			}
			return nil
		}
	case reflect.Float32, reflect.Float64:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetFloat(n.Float())
			_, err := strconv.ParseFloat(n.Str, 64)
			if n.Type == gjson.JSON || (n.Type == gjson.String && err != nil) {
				return fmt.Errorf("apijson: failed to parse float")
			}
			return nil
		}
	default:
		return func(node gjson.Result, v reflect.Value) error {
			return fmt.Errorf("unknown type received at primitive decoder: %s", t.String())
		}
	}
}

func (d *decoder) newTimeTypeDecoder(t reflect.Type) decoderFunc {
	format := d.dateFormat
	return func(n gjson.Result, v reflect.Value) error {
		parsed, err := time.Parse(format, n.Str)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(parsed).Convert(t))
		return nil
	}
}

func setUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
}
