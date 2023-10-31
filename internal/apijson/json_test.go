package apijson

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/tidwall/gjson"
)

func P[T any](v T) *T { return &v }

type Primitives struct {
	A bool    `json:"a"`
	B int     `json:"b"`
	C uint    `json:"c"`
	D float64 `json:"d"`
	E float32 `json:"e"`
	F []int   `json:"f"`
}

type PrimitivePointers struct {
	A *bool    `json:"a"`
	B *int     `json:"b"`
	C *uint    `json:"c"`
	D *float64 `json:"d"`
	E *float32 `json:"e"`
	F *[]int   `json:"f"`
}

type Slices struct {
	Slice []Primitives `json:"slices"`
}

type DateTime struct {
	Date     time.Time `json:"date" format:"date"`
	DateTime time.Time `json:"date-time" format:"date-time"`
}

type AdditionalProperties struct {
	A      bool                   `json:"a"`
	Extras map[string]interface{} `json:"-,extras"`
}

type TypedAdditionalProperties struct {
	A      bool           `json:"a"`
	Extras map[string]int `json:"-,extras"`
}

type EmbeddedStructs struct {
	AdditionalProperties
	A      *int                   `json:"number2"`
	Extras map[string]interface{} `json:"-,extras"`
}

type Recursive struct {
	Name  string     `json:"name"`
	Child *Recursive `json:"child"`
}

type JSONFieldStruct struct {
	A      bool                `json:"a"`
	B      int64               `json:"b"`
	C      string              `json:"c"`
	D      string              `json:"d"`
	Extras map[string]int64    `json:"-,extras"`
	JSON   JSONFieldStructJSON `json:"-,metadata"`
}

type JSONFieldStructJSON struct {
	A      Field
	B      Field
	C      Field
	D      Field
	Extras map[string]Field
	raw    string
}

type UnknownStruct struct {
	Unknown interface{} `json:"unknown"`
}

type UnionStruct struct {
	Union Union `json:"union" format:"date"`
}

type Union interface {
	union()
}

type Inline struct {
	InlineField Primitives `json:"-,inline"`
	JSON        InlineJSON `json:"-,metadata"`
}

type InlineArray struct {
	InlineField []string   `json:"-,inline"`
	JSON        InlineJSON `json:"-,metadata"`
}

type InlineJSON struct {
	InlineField Field
	raw         string
}

func init() {
	RegisterUnion(reflect.TypeOf((*Union)(nil)).Elem(), "type",
		UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionTime{}),
		},
		UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(UnionInteger(0)),
		},
		UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "typeA",
			Type:               reflect.TypeOf(UnionStructA{}),
		},
		UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "typeB",
			Type:               reflect.TypeOf(UnionStructB{}),
		},
	)
}

type UnionInteger int64

func (UnionInteger) union() {}

type UnionStructA struct {
	Type string `json:"type"`
	A    string `json:"a"`
	B    string `json:"b"`
}

func (UnionStructA) union() {}

type UnionStructB struct {
	Type string `json:"type"`
	A    string `json:"a"`
}

func (UnionStructB) union() {}

type UnionTime time.Time

func (UnionTime) union() {}

var tests = map[string]struct {
	buf string
	val interface{}
}{
	"true":               {"true", true},
	"false":              {"false", false},
	"int":                {"1", 1},
	"int_bigger":         {"12324", 12324},
	"int_string_coerce":  {`"65"`, 65},
	"int_boolean_coerce": {"true", 1},
	"int64":              {"1", int64(1)},
	"int64_huge":         {"123456789123456789", int64(123456789123456789)},
	"uint":               {"1", uint(1)},
	"uint_bigger":        {"12324", uint(12324)},
	"uint_coerce":        {`"65"`, uint(65)},
	"float_1.54":         {"1.54", float32(1.54)},
	"float_1.89":         {"1.89", float64(1.89)},
	"string":             {`"str"`, "str"},
	"string_int_coerce":  {`12`, "12"},
	"array_string":       {`["foo","bar"]`, []string{"foo", "bar"}},
	"array_int":          {`[1,2]`, []int{1, 2}},
	"array_int_coerce":   {`["1",2]`, []int{1, 2}},

	"ptr_true":               {"true", P(true)},
	"ptr_false":              {"false", P(false)},
	"ptr_int":                {"1", P(1)},
	"ptr_int_bigger":         {"12324", P(12324)},
	"ptr_int_string_coerce":  {`"65"`, P(65)},
	"ptr_int_boolean_coerce": {"true", P(1)},
	"ptr_int64":              {"1", P(int64(1))},
	"ptr_int64_huge":         {"123456789123456789", P(int64(123456789123456789))},
	"ptr_uint":               {"1", P(uint(1))},
	"ptr_uint_bigger":        {"12324", P(uint(12324))},
	"ptr_uint_coerce":        {`"65"`, P(uint(65))},
	"ptr_float_1.54":         {"1.54", P(float32(1.54))},
	"ptr_float_1.89":         {"1.89", P(float64(1.89))},

	"date_time": {`"2007-03-01T13:00:00Z"`, time.Date(2007, time.March, 1, 13, 0, 0, 0, time.UTC)},

	"map_string":    {`{"foo":"bar"}`, map[string]string{"foo": "bar"}},
	"map_interface": {`{"a":1,"b":"str","c":false}`, map[string]interface{}{"a": float64(1), "b": "str", "c": false}},

	"primitive_struct": {
		`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}`,
		Primitives{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
	},

	"slices": {
		`{"slices":[{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}]}`,
		Slices{
			Slice: []Primitives{{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}}},
		},
	},

	"primitive_pointer_struct": {
		`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4,5]}`,
		PrimitivePointers{
			A: P(false),
			B: P(237628372683),
			C: P(uint(654)),
			D: P(9999.43),
			E: P(float32(43.76)),
			F: &[]int{1, 2, 3, 4, 5},
		},
	},

	"datetime_struct": {
		`{"date":"2006-01-02","date-time":"2006-01-02T15:04:05Z"}`,
		DateTime{
			Date:     time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
			DateTime: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
		},
	},

	"additional_properties": {
		`{"a":true,"bar":"value","foo":true}`,
		AdditionalProperties{
			A: true,
			Extras: map[string]interface{}{
				"bar": "value",
				"foo": true,
			},
		},
	},

	"recursive_struct": {
		`{"child":{"name":"Alex"},"name":"Robert"}`,
		Recursive{Name: "Robert", Child: &Recursive{Name: "Alex"}},
	},

	"metadata_coerce": {
		`{"a":"12","b":"12","c":null,"extra_typed":12,"extra_untyped":{"foo":"bar"}}`,
		JSONFieldStruct{
			A: false,
			B: 12,
			C: "",
			JSON: JSONFieldStructJSON{
				raw: `{"a":"12","b":"12","c":null,"extra_typed":12,"extra_untyped":{"foo":"bar"}}`,
				A:   Field{raw: `"12"`, status: invalid},
				B:   Field{raw: `"12"`, status: valid},
				C:   Field{raw: "null", status: null},
				D:   Field{raw: "", status: missing},
				Extras: map[string]Field{
					"extra_typed": {
						raw:    "12",
						status: valid,
					},
					"extra_untyped": {
						raw:    `{"foo":"bar"}`,
						status: invalid,
					},
				},
			},
			Extras: map[string]int64{
				"extra_typed":   12,
				"extra_untyped": 0,
			},
		},
	},

	"unknown_struct_number": {
		`{"unknown":12}`,
		UnknownStruct{
			Unknown: 12.,
		},
	},

	"unknown_struct_map": {
		`{"unknown":{"foo":"bar"}}`,
		UnknownStruct{
			Unknown: map[string]interface{}{
				"foo": "bar",
			},
		},
	},

	"union_integer": {
		`{"union":12}`,
		UnionStruct{
			Union: UnionInteger(12),
		},
	},

	"union_struct_discriminated_a": {
		`{"union":{"a":"foo","b":"bar","type":"typeA"}}`,
		UnionStruct{
			Union: UnionStructA{
				Type: "typeA",
				A:    "foo",
				B:    "bar",
			},
		},
	},

	"union_struct_discriminated_b": {
		`{"union":{"a":"foo","type":"typeB"}}`,
		UnionStruct{
			Union: UnionStructB{
				Type: "typeB",
				A:    "foo",
			},
		},
	},

	"union_struct_time": {
		`{"union":"2010-05-23"}`,
		UnionStruct{
			Union: UnionTime(time.Date(2010, 05, 23, 0, 0, 0, 0, time.UTC)),
		},
	},

	"inline_coerce": {
		`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}`,
		Inline{
			InlineField: Primitives{A: false, B: 237628372683, C: 0x28e, D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
			JSON: InlineJSON{
				InlineField: Field{raw: "{\"a\":false,\"b\":237628372683,\"c\":654,\"d\":9999.43,\"e\":43.76,\"f\":[1,2,3,4]}", status: 3},
				raw:         "{\"a\":false,\"b\":237628372683,\"c\":654,\"d\":9999.43,\"e\":43.76,\"f\":[1,2,3,4]}",
			},
		},
	},

	"inline_array_coerce": {
		`["Hello","foo","bar"]`,
		InlineArray{
			InlineField: []string{"Hello", "foo", "bar"},
			JSON: InlineJSON{
				InlineField: Field{raw: `["Hello","foo","bar"]`, status: 3},
				raw:         `["Hello","foo","bar"]`,
			},
		},
	},
}

func TestDecode(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := reflect.New(reflect.TypeOf(test.val))
			if err := Unmarshal([]byte(test.buf), result.Interface()); err != nil {
				t.Fatalf("deserialization of %v failed with error %v", result, err)
			}
			if !reflect.DeepEqual(result.Elem().Interface(), test.val) {
				t.Fatalf("expected '%s' to deserialize to \n%#v\nbut got\n%#v", test.buf, test.val, result.Elem().Interface())
			}
		})
	}
}

func TestEncode(t *testing.T) {
	for name, test := range tests {
		if strings.HasSuffix(name, "_coerce") {
			continue
		}
		t.Run(name, func(t *testing.T) {
			raw, err := Marshal(test.val)
			if err != nil {
				t.Fatalf("serialization of %v failed with error %v", test.val, err)
			}
			if string(raw) != test.buf {
				t.Fatalf("expected %+#v to serialize to %s but got %s", test.val, test.buf, string(raw))
			}
		})
	}
}
