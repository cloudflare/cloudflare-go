package core_test

import (
	"fmt"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	"github.com/cloudflare/cloudflare-sdk-go/core/pointers"
)

func assert(t *testing.T, got, expected string) {
	if got != expected {
		t.Errorf(`expected "%s"
got "%s"`, expected, got)
	}
}

type ShipmentResponse struct {
	Shipments *[]Shipment
	Names     *[]string
}

func (r ShipmentResponse) String() (result string) {
	return fmt.Sprintf("&ShipmentResponse{Shipments:%s Names:%s}", core.Fmt(r.Shipments), core.Fmt(r.Names))
}

type Shipment struct {
	Name    *string  `json:"name"`
	Weight  *int     `json:"weight"`
	Fragile *bool    `json:"fragile"`
	Address *Address `json:"address"`
}

func (r Shipment) String() (result string) {
	return fmt.Sprintf("&Shipment{Name:%s Weight:%s Fragile:%s Address:%s}", core.FmtP(r.Name), core.FmtP(r.Weight), core.FmtP(r.Fragile), r.Address)
}

type Address struct {
	Line1 *string `json:"line1"`
	Line2 *string `json:"line2"`
}

func (r Address) String() (result string) {
	return fmt.Sprintf("&Address{Line1:%s Line2:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2))
}

func TestFP(t *testing.T) {
	assert(t, core.FmtP(pointers.P("hello")), `"hello"`)
	assert(t, core.FmtP(pointers.P(false)), `false`)
	assert(t, core.FmtP(pointers.P(62)), `62`)
}

func TestFR(t *testing.T) {
	sa := Shipment{}
	assert(t, fmt.Sprintf("%s", &sa), "&Shipment{Name:<nil> Weight:<nil> Fragile:<nil> Address:<nil>}")
	sa.Name = pointers.P("hello")
	sa.Weight = pointers.P(12)
	sa.Fragile = pointers.P(true)
	sa.Address = &Address{Line1: pointers.P("1811 Francisco St."), Line2: pointers.P("Unit 2")}
	assert(t, fmt.Sprintf("%s", &sa), `&Shipment{Name:"hello" Weight:12 Fragile:true Address:&Address{Line1:"1811 Francisco St." Line2:"Unit 2"}}`)
}

func TestFSlice(t *testing.T) {
	sr := ShipmentResponse{
		Shipments: &[]Shipment{
			{
				Name:    pointers.P("hello"),
				Weight:  pointers.P(12),
				Fragile: pointers.P(true),
				Address: &Address{Line1: pointers.P("1811 Francisco St."), Line2: pointers.P("Unit 2")},
			},
			{
				Name:    pointers.P("hello"),
				Weight:  pointers.P(12),
				Fragile: pointers.P(true),
				Address: &Address{Line1: pointers.P("1811 Francisco St."), Line2: pointers.P("Unit 2")},
			},
		},
		Names: &[]string{"hello", "I exist therefore I am"},
	}
	assert(t, fmt.Sprintf("%s", &sr), `&ShipmentResponse{Shipments:[&Shipment{Name:"hello" Weight:12 Fragile:true Address:&Address{Line1:"1811 Francisco St." Line2:"Unit 2"}} &Shipment{Name:"hello" Weight:12 Fragile:true Address:&Address{Line1:"1811 Francisco St." Line2:"Unit 2"}}] Names:["hello" "I exist therefore I am"]}`)
}

type Table struct {
	Rows      *[][]string
	Addresses *[][]Address
	Records   *map[string]string
}

func (r Table) String() string {
	return fmt.Sprintf("&Table{Rows:%s Addresses:%s Records:%s}", core.Fmt(r.Rows), core.Fmt(r.Addresses), core.Fmt(r.Records))
}

func TestFSlice2(t *testing.T) {
	tb := &Table{
		Rows:      &[][]string{{"hello", "right on"}, {"yo"}},
		Addresses: &[][]Address{{{Line1: pointers.P("1811 Francisco St.")}, {Line1: pointers.P("1823 Francisco St.")}}, {{Line1: pointers.P("1833 Francisco St.")}}},
		Records:   &map[string]string{"foo": "bar"},
	}
	assert(t, fmt.Sprintf("%s", tb), `&Table{Rows:[["hello" "right on"] ["yo"]] Addresses:[[&Address{Line1:"1811 Francisco St." Line2:<nil>} &Address{Line1:"1823 Francisco St." Line2:<nil>}] [&Address{Line1:"1833 Francisco St." Line2:<nil>}]] Records:{"foo":"bar"}}`)
}
