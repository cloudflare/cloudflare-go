package core

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// Returns the first string or *string in the argument that is valid and not an
// empty string
func CoalesceStrings[P ~string | ~*string](strings ...P) string {
	for _, str := range strings {
		if str, ok := any(str).(string); ok && len(str) != 0 {
			return str
		}
		if str, ok := any(str).(*string); ok && len(*str) != 0 {
			return *str
		}
	}
	return ""
}

type Primitive interface {
	~bool | ~byte | ~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~string | ~uint | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~complex64 | ~complex128 | time.Time
}

// Fast path for pretty printing pointer primitives
func FmtP[T Primitive](v *T) string {
	if v == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%#v", *v)
}

// Pretty printer that recursively prints structs, arrays, maps, and primitives.
func Fmt(v interface{}) string {
	if v == nil {
		return "<nil>"
	}

	switch v := v.(type) {
	case *bool, *byte, *float32, *float64, *int, *int8, *int16, *int32, *int64, *string, *uint, *uint16, *uint32, *uint64, *uintptr, *complex64, *complex128:
		return fmt.Sprintf("%#v", *(v).(*interface{}))
	case bool, byte, float32, float64, int, int8, int16, int32, int64, string, uint, uint16, uint32, uint64, uintptr, complex64, complex128:
		return fmt.Sprintf("%#v", v)
	case fmt.Stringer:
		return v.String()
	case *fmt.Stringer:
		return (*v).String()
	}

	r := reflect.ValueOf(v)
	for r.Kind() == reflect.Ptr {
		r = r.Elem()
	}

	switch r.Kind() {
	case reflect.Slice:
		sb := strings.Builder{}
		sb.WriteRune('[')
		for i := 0; i < r.Len(); i++ {
			if i != 0 {
				sb.WriteString(" ")
			}
			sb.WriteString(Fmt(r.Index(i).Interface()))
		}
		sb.WriteRune(']')
		return sb.String()

	case reflect.Map:
		sb := strings.Builder{}
		sb.WriteRune('{')
		for i, k := range r.MapKeys() {
			if i != 0 {
				sb.WriteString(" ")
			}
			sb.WriteString(fmt.Sprintf("%s:%s", Fmt(k.Interface()), Fmt(r.MapIndex(k).Interface())))
		}
		sb.WriteRune('}')
		return sb.String()
	default:
		return r.String()
	}
}
