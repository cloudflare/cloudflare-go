package fields

import (
	"fmt"
)

type FieldLike interface {
	field()
}

type Field[T any] struct {
	Value   T
	Null    bool
	Present bool
	Raw     any
}

func F[T any](value T) Field[T]          { return Field[T]{Value: value, Present: true} }
func NullField[T any]() Field[T]         { return Field[T]{Null: true, Present: true} }
func RawField[T any](value any) Field[T] { return Field[T]{Raw: value, Present: true} }

func (f Field[T]) field() {}

func (f Field[T]) String() string {
	if s, ok := any(f.Value).(fmt.Stringer); ok {
		return s.String()
	}
	return fmt.Sprintf("%#v", f.Value)
}
