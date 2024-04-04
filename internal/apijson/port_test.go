package apijson

import (
	"reflect"
	"testing"
)

type Metadata struct {
	CreatedAt string `json:"created_at"`
}

// Card is the "combined" type of CardVisa and CardMastercard
type Card struct {
	Processor CardProcessor `json:"processor"`
	Data      any           `json:"data"`
	IsFoo     bool          `json:"is_foo"`
	IsBar     bool          `json:"is_bar"`
	Metadata  Metadata      `json:"metadata"`

	JSON cardJSON
}

type cardJSON struct {
	Processor Field
	Data      Field
	IsFoo     Field
	IsBar     Field
	Metadata  Field
	Extras    map[string]Field
	raw       string
}

func (r cardJSON) RawJSON() string { return r.raw }

type CardProcessor string

// CardVisa
type CardVisa struct {
	Processor CardVisaProcessor `json:"processor"`
	Data      CardVisaData      `json:"data"`
	IsFoo     bool              `json:"is_foo"`
	Metadata  Metadata          `json:"metadata"`

	JSON cardVisaJSON
}

type cardVisaJSON struct {
	Processor Field
	Data      Field
	IsFoo     Field
	Metadata  Field
	Extras    map[string]Field
	raw       string
}

func (r cardVisaJSON) RawJSON() string { return r.raw }

type CardVisaProcessor string

type CardVisaData struct {
	Foo string `json:"foo"`
}

// CardMastercard
type CardMastercard struct {
	Processor CardMastercardProcessor `json:"processor"`
	Data      CardMastercardData      `json:"data"`
	IsBar     bool                    `json:"is_bar"`
	Metadata  Metadata                `json:"metadata"`

	JSON cardMastercardJSON
}

type cardMastercardJSON struct {
	Processor Field
	Data      Field
	IsBar     Field
	Metadata  Field
	Extras    map[string]Field
	raw       string
}

func (r cardMastercardJSON) RawJSON() string { return r.raw }

type CardMastercardProcessor string

type CardMastercardData struct {
	Bar int64 `json:"bar"`
}

var portTests = map[string]struct {
	from any
	to   any
}{
	"visa to card": {
		CardVisa{
			Processor: "visa",
			IsFoo:     true,
			Data: CardVisaData{
				Foo: "foo",
			},
			Metadata: Metadata{
				CreatedAt: "Mar 29 2024",
			},
			JSON: cardVisaJSON{
				raw:       `{"processor":"visa","is_foo":true,"data":{"foo":"foo"}}`,
				Processor: Field{raw: `"visa"`, status: valid},
				IsFoo:     Field{raw: `true`, status: valid},
				Data:      Field{raw: `{"foo":"foo"}`, status: valid},
				Extras:    map[string]Field{"extra": {raw: `"yo"`, status: valid}},
			},
		},
		Card{
			Processor: "visa",
			IsFoo:     true,
			IsBar:     false,
			Data: CardVisaData{
				Foo: "foo",
			},
			Metadata: Metadata{
				CreatedAt: "Mar 29 2024",
			},
			JSON: cardJSON{
				raw:       `{"processor":"visa","is_foo":true,"data":{"foo":"foo"}}`,
				Processor: Field{raw: `"visa"`, status: valid},
				IsFoo:     Field{raw: `true`, status: valid},
				Data:      Field{raw: `{"foo":"foo"}`, status: valid},
				Extras:    map[string]Field{"extra": {raw: `"yo"`, status: valid}},
			},
		},
	},
	"mastercard to card": {
		CardMastercard{
			Processor: "visa",
			IsBar:     true,
			Data: CardMastercardData{
				Bar: 13,
			},
		},
		Card{
			Processor: "visa",
			IsFoo:     false,
			IsBar:     true,
			Data: CardMastercardData{
				Bar: 13,
			},
		},
	},
}

func TestPort(t *testing.T) {
	for name, test := range portTests {
		t.Run(name, func(t *testing.T) {
			toVal := reflect.New(reflect.TypeOf(test.to))

			err := Port(test.from, toVal.Interface())
			if err != nil {
				t.Fatalf("port of %v failed with error %v", test.from, err)
			}

			if !reflect.DeepEqual(toVal.Elem().Interface(), test.to) {
				t.Fatalf("expected:\n%+#v\n\nto port to:\n%+#v\n\nbut got:\n%+#v", test.from, test.to, toVal.Elem().Interface())
			}
		})
	}
}
