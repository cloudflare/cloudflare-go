package query

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

var encoders sync.Map // map[reflect.Type]encoderFunc

type encoder struct {
	settings QuerySettings
}

type encoderEntry struct {
	reflect.Type
	settings QuerySettings
}

type encoderFunc func(key string, value reflect.Value) []Pair

type Pair struct {
	key   string
	value string
}

type parsedStructTag struct {
	name      string
	omitempty bool
	inline    bool
}

func parseStructTag(raw string) (tag parsedStructTag) {
	parts := strings.Split(raw, ",")
	if len(parts) == 0 {
		return
	}
	tag.name = parts[0]
	for _, part := range parts[1:] {
		switch part {
		case "omitempty":
			tag.omitempty = true
		case "inline":
			tag.inline = true
		}
	}
	return
}

type structField struct {
	parsedStructTag
	encoderFunc
}

func (e *encoder) typeEncoder(t reflect.Type) encoderFunc {
	entry := encoderEntry{t, e.settings}
	if fi, ok := encoders.Load(entry); ok {
		return fi.(encoderFunc)
	}

	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  encoderFunc
	)
	wg.Add(1)
	fi, loaded := encoders.LoadOrStore(entry, encoderFunc(func(key string, v reflect.Value) []Pair {
		wg.Wait()
		return f(key, v)
	}))
	if loaded {
		return fi.(encoderFunc)
	}

	// Compute the real encoder and replace the indirect func with it.
	f = e.newTypeEncoder(t)
	wg.Done()
	encoders.Store(entry, f)
	return f
}

func (e *encoder) newTypeEncoder(t reflect.Type) encoderFunc {
	switch t.Kind() {
	case reflect.Pointer:
		encoder := e.typeEncoder(t.Elem())
		return func(key string, value reflect.Value) (pairs []Pair) {
			if !value.IsValid() || value.IsNil() {
				return
			}
			pairs = encoder(key, value.Elem())
			return
		}
	case reflect.Struct:
		return e.newStructTypeEncoder(t)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return e.newArrayTypeEncoder(t)
	case reflect.Map:
		return e.newMapEncoder(t)
	case reflect.Interface:
		return func(key string, value reflect.Value) []Pair {
			value = value.Elem()
			if !value.IsValid() {
				return nil
			}
			return e.typeEncoder(value.Type())(key, value)
		}
	default:
		return e.newPrimitiveTypeEncoder(t)
	}
}

func (e *encoder) newStructTypeEncoder(t reflect.Type) encoderFunc {
	if t.Implements(reflect.TypeOf((*fields.FieldLike)(nil)).Elem()) {
		return e.newFieldTypeEncoder(t)
	}

	fieldEncoders := map[int]structField{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			t := field.Type
			if !field.IsExported() && t.Kind() != reflect.Struct {
				// Ignore embedded fields of unexported non-struct types.
				continue
			}
			// Do not ignore embedded fields of unexported struct types
			// since they may have exported fields.
		}
		if !field.IsExported() {
			continue
		}
		tag, ok := field.Tag.Lookup(queryStructTag)
		if !ok {
			continue
		}
		fieldEncoders[i] = structField{parseStructTag(tag), e.typeEncoder(field.Type)}
	}

	return func(key string, value reflect.Value) (pairs []Pair) {
		for i, field := range fieldEncoders {
			if field.omitempty {
				if !value.IsValid() || value.IsZero() {
					continue
				}
			}
			var subkey string = field.name
			if field.inline {
				subkey = key
			} else {
				subkey = e.renderKeyPath(key, subkey)
			}
			for _, pair := range field.encoderFunc(subkey, value.Field(i)) {
				if field.omitempty && len(pair.value) == 0 {
					continue
				}
				pairs = append(pairs, pair)
			}
		}
		return
	}
}

func (e *encoder) newMapEncoder(t reflect.Type) encoderFunc {
	keyEncoder := e.typeEncoder(t.Key())
	elementEncoder := e.typeEncoder(t.Elem())
	return func(key string, value reflect.Value) (pairs []Pair) {
		iter := value.MapRange()
		for iter.Next() {
			encodedKey := keyEncoder("", iter.Key())
			if len(encodedKey) != 1 {
				panic("Unexpected number of parts for encoded map key. Are you using a non-primitive for this map?")
			}
			subkey := encodedKey[0].value
			keyPath := e.renderKeyPath(key, subkey)
			pairs = append(pairs, elementEncoder(keyPath, iter.Value())...)
		}
		return
	}
}

func (e *encoder) renderKeyPath(key string, subkey string) string {
	if len(key) == 0 {
		return subkey
	}
	if e.settings.NestedFormat == NestedQueryFormatDots {
		return fmt.Sprintf("%s.%s", key, subkey)
	}
	return fmt.Sprintf("%s[%s]", key, subkey)
}

func (e *encoder) newArrayTypeEncoder(t reflect.Type) encoderFunc {
	switch e.settings.ArrayFormat {
	case ArrayQueryFormatComma:
		innerEncoder := e.typeEncoder(t.Elem())
		return func(key string, v reflect.Value) []Pair {
			elements := []string{}
			for i := 0; i < v.Len(); i++ {
				for _, pair := range innerEncoder("", v.Index(i)) {
					elements = append(elements, pair.value)
				}
			}
			if len(elements) == 0 {
				return []Pair{}
			}
			return []Pair{{key, strings.Join(elements, ",")}}
		}
	case ArrayQueryFormatRepeat:
		innerEncoder := e.typeEncoder(t.Elem())
		return func(key string, value reflect.Value) (pairs []Pair) {
			for i := 0; i < value.Len(); i++ {
				pairs = append(pairs, innerEncoder(key, value.Index(i))...)
			}
			return pairs
		}
	case ArrayQueryFormatIndices:
		panic("The array indices format is not supported yet")
	case ArrayQueryFormatBrackets:
		innerEncoder := e.typeEncoder(t.Elem())
		return func(key string, value reflect.Value) []Pair {
			pairs := []Pair{}
			for i := 0; i < value.Len(); i++ {
				pairs = append(pairs, innerEncoder(key+"[]", value)...)
			}
			return pairs
		}
	default:
		panic(fmt.Sprintf("Unknown ArrayFormat value: %d", e.settings.ArrayFormat))
	}
}

func (e *encoder) newPrimitiveTypeEncoder(t reflect.Type) encoderFunc {
	switch t.Kind() {
	case reflect.Pointer:
		inner := t.Elem()

		innerEncoder := e.newPrimitiveTypeEncoder(inner)
		return func(key string, v reflect.Value) []Pair {
			if !v.IsValid() || v.IsNil() {
				return nil
			}
			return innerEncoder(key, v.Elem())
		}
	case reflect.String:
		return func(key string, v reflect.Value) []Pair {
			return []Pair{{key, v.String()}}
		}
	case reflect.Bool:
		return func(key string, v reflect.Value) []Pair {
			if v.Bool() {
				return []Pair{{key, "true"}}
			}
			return []Pair{{key, "false"}}
		}
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(key string, v reflect.Value) []Pair {
			return []Pair{{key, strconv.FormatInt(v.Int(), 10)}}
		}
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(key string, v reflect.Value) []Pair {
			return []Pair{{key, strconv.FormatUint(v.Uint(), 10)}}
		}
	case reflect.Float32, reflect.Float64:
		return func(key string, v reflect.Value) []Pair {
			return []Pair{{key, strconv.FormatFloat(v.Float(), 'f', -1, 64)}}
		}
	case reflect.Complex64, reflect.Complex128:
		bitSize := 64
		if t.Kind() == reflect.Complex128 {
			bitSize = 128
		}
		return func(key string, v reflect.Value) []Pair {
			return []Pair{{key, strconv.FormatComplex(v.Complex(), 'f', -1, bitSize)}}
		}
	default:
		return func(key string, v reflect.Value) []Pair {
			return nil
		}
	}
}

func (e *encoder) newFieldTypeEncoder(t reflect.Type) encoderFunc {
	f, _ := t.FieldByName("Value")
	enc := e.typeEncoder(f.Type)

	return func(key string, value reflect.Value) []Pair {
		present := value.FieldByName("Present")
		if !present.Bool() {
			return nil
		}
		null := value.FieldByName("Null")
		if null.Bool() {
			// TODO: Error?
			return nil
		}
		raw := value.FieldByName("Raw")
		if !raw.IsNil() {
			return e.typeEncoder(raw.Type())(key, raw)
		}
		return enc(key, value.FieldByName("Value"))
	}
}
