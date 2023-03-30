package json

import "reflect"

type status uint8

const (
	missing status = iota
	null
	invalid
	valid
)

type Metadata struct {
	raw    []byte
	status status
}

func (j Metadata) IsNull() bool    { return j.status <= null }
func (j Metadata) IsMissing() bool { return j.status == missing }
func (j Metadata) IsInvalid() bool { return j.status == invalid }
func (j Metadata) Raw() []byte     { return j.raw }

func getMetadataField(root reflect.Value, index []int, name string) reflect.Value {
	strct := root.FieldByIndex(index[:len(index)-1])
	if !strct.IsValid() {
		panic("couldn't find encapsulating struct for field " + name)
	}
	meta := strct.FieldByName("JSON")
	if !meta.IsValid() {
		return reflect.Value{}
	}
	field := meta.FieldByName(name)
	if !field.IsValid() {
		return reflect.Value{}
	}
	return field
}
