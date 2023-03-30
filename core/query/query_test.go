package query

import (
	"net/url"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go/core/pointers"
)

type EmptyTestC struct {
	Field *string `query:"field"`
}

type EmptyTestB struct {
	C EmptyTestC `query:"c"`
}

type EmptyTestA struct {
	B EmptyTestB `query:"b"`
}

type EmptyTest struct {
	A EmptyTestA `query:"a"`
}

func TestEmpty(t *testing.T) {
	assert(t, EmptyTest{}, "", QuerySettings{})
}

type BasicTest struct {
	A interface{} `query:"a"`
}

func TestBasic(t *testing.T) {
	assert(t, BasicTest{A: nil}, "", QuerySettings{})
	assert(t, BasicTest{A: 1}, "a=1", QuerySettings{})
	assert(t, BasicTest{A: "b"}, "a=b", QuerySettings{})
	assert(t, BasicTest{A: true}, "a=true", QuerySettings{})
	assert(t, BasicTest{A: false}, "a=false", QuerySettings{})
	assert(t, BasicTest{A: pointers.P(1.23456)}, "a=1.23456", QuerySettings{})
}

func serialize(v interface{}, options QuerySettings) string {
	escaped := MarshalWithSettings(v, options).Encode()
	unescaped, err := url.QueryUnescape(escaped)
	if err != nil {
		panic(err)
	}
	return unescaped
}

func assert(t *testing.T, v interface{}, expectation string, settings QuerySettings) {
	if serialized := serialize(v, settings); serialized != expectation {
		t.Fatalf("Expected %v to serialize to %s but got %s", v, expectation, serialized)
	}
}

type NestedTestDepth1A struct {
	B *string `query:"b"`
	D *string `query:"d"`
	F *string `query:"f"`
}

type NestedTestDepth1 struct {
	A *NestedTestDepth1A `query:"a"`
}

type NestedTestDepth3C struct {
	D *string `query:"d"`
}

type NestedTestDepth3B struct {
	C *NestedTestDepth3C `query:"c"`
}

type NestedTestDepth3A struct {
	B *NestedTestDepth3B `query:"b"`
}

type NestedTestDepth3 struct {
	A *NestedTestDepth3A `query:"a"`
}

func TestNestedDotted(t *testing.T) {
	settings := QuerySettings{NestedFormat: NestedQueryFormatDots}

	assert(t, NestedTestDepth1{&NestedTestDepth1A{B: pointers.P("c")}}, "a.b=c", settings)
	assert(t, NestedTestDepth1{&NestedTestDepth1A{B: pointers.P("c"), D: pointers.P("e"), F: pointers.P("g")}}, "a.b=c&a.d=e&a.f=g", settings)
	assert(t, NestedTestDepth3{&NestedTestDepth3A{&NestedTestDepth3B{&NestedTestDepth3C{D: pointers.P("e")}}}}, "a.b.c.d=e", settings)
}

func TestNestedBrackets(t *testing.T) {
	settings := QuerySettings{NestedFormat: NestedQueryFormatBrackets}

	assert(t, NestedTestDepth1{&NestedTestDepth1A{B: pointers.P("c")}}, "a[b]=c", settings)
	assert(t, NestedTestDepth1{&NestedTestDepth1A{B: pointers.P("c"), D: pointers.P("e"), F: pointers.P("g")}}, "a[b]=c&a[d]=e&a[f]=g", settings)
	assert(t, NestedTestDepth3{&NestedTestDepth3A{&NestedTestDepth3B{&NestedTestDepth3C{D: pointers.P("e")}}}}, "a[b][c][d]=e", settings)
}

type ArrayTestDepth0 struct {
	In []string `query:"in"`
}

type ArrayTestDepth1A struct {
	B []*bool `query:"b"`
}

type ArrayTestDepth1 struct {
	A *ArrayTestDepth1A `query:"a"`
}

func TestArrayComma(t *testing.T) {
	settings := QuerySettings{ArrayFormat: ArrayQueryFormatComma}

	assert(t, ArrayTestDepth0{[]string{"foo", "bar"}}, "in=foo,bar", settings)
	assert(t, ArrayTestDepth1{&ArrayTestDepth1A{B: []*bool{pointers.P(true), pointers.P(false)}}}, "a[b]=true,false", settings)
	assert(t, ArrayTestDepth1{&ArrayTestDepth1A{B: []*bool{pointers.P(true), pointers.P(false), nil, pointers.P(true)}}}, "a[b]=true,false,true", settings)
}

type ArrayRepeatMixed struct {
	In []interface{} `query:"in"`
}

type ArrayRepeatMixedB struct {
	C []string `query:"c"`
}

type ArrayRepeatMixedA struct {
	B ArrayRepeatMixedB `query:"b"`
}

func TestArrayRepeat(t *testing.T) {
	settings := QuerySettings{ArrayFormat: ArrayQueryFormatRepeat}

	assert(t, ArrayTestDepth0{[]string{"foo", "bar"}}, "in=foo&in=bar", settings)
	assert(t, ArrayTestDepth1{&ArrayTestDepth1A{B: []*bool{pointers.P(true), pointers.P(false)}}}, "a[b]=true&a[b]=false", settings)
	assert(t, ArrayTestDepth1{&ArrayTestDepth1A{B: []*bool{pointers.P(true), pointers.P(false), nil, pointers.P(true)}}}, "a[b]=true&a[b]=false&a[b]=true", settings)
	assert(t, ArrayRepeatMixed{In: []interface{}{"foof", ArrayRepeatMixedA{B: ArrayRepeatMixedB{C: []string{"d", "e"}}}}}, "in=foof&in[b][c]=d&in[b][c]=e", settings)
}

func TestMapMarshal(t *testing.T) {
	settings := QuerySettings{}

	assert(t, map[string]interface{}{"hello": "world", "goodbye": map[string]string{"friend": "yes"}}, "goodbye[friend]=yes&hello=world", settings)
}

type InlineExtrasTest struct {
	A      *string           `query:"a"`
	Extras map[string]string `query:"-,inline"`
}

func TestInlineExtras(t *testing.T) {
	settings := QuerySettings{}

	assert(t, InlineExtrasTest{pointers.P("hi"), map[string]string{"there": "neighbor"}}, "a=hi&there=neighbor", settings)
}
