package form

type Marshaler interface {
	MarshalMultipart() ([]byte, error)
}
