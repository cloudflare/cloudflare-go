package apiform

import (
	"bytes"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"strings"
	"testing"
	"time"
)

func P[T any](v T) *T { return &v }

type Primitives struct {
	A bool    `form:"a"`
	B int     `form:"b"`
	C uint    `form:"c"`
	D float64 `form:"d"`
	E float32 `form:"e"`
	F []int   `form:"f"`
}

type PrimitivePointers struct {
	A *bool    `form:"a"`
	B *int     `form:"b"`
	C *uint    `form:"c"`
	D *float64 `form:"d"`
	E *float32 `form:"e"`
	F *[]int   `form:"f"`
}

type Slices struct {
	Slice []Primitives `form:"slices"`
}

type DateTime struct {
	Date     time.Time `form:"date" format:"date"`
	DateTime time.Time `form:"date-time" format:"date-time"`
}

type AdditionalProperties struct {
	A      bool                   `form:"a"`
	Extras map[string]interface{} `form:"-" api:"extrafields"`
}

type TypedAdditionalProperties struct {
	A      bool           `form:"a"`
	Extras map[string]int `form:"-" api:"extrafields"`
}

type EmbeddedStructs struct {
	AdditionalProperties
	A      *int                   `form:"number2"`
	Extras map[string]interface{} `form:"-" api:"extrafields"`
}

type Recursive struct {
	Name  string     `form:"name"`
	Child *Recursive `form:"child"`
}

type UnknownStruct struct {
	Unknown interface{} `form:"unknown"`
}

type UnionStruct struct {
	Union Union `form:"union" format:"date"`
}

type Union interface {
	union()
}

type UnionInteger int64

func (UnionInteger) union() {}

type UnionStructA struct {
	Type string `form:"type"`
	A    string `form:"a"`
	B    string `form:"b"`
}

func (UnionStructA) union() {}

type UnionStructB struct {
	Type string `form:"type"`
	A    string `form:"a"`
}

func (UnionStructB) union() {}

type UnionTime time.Time

func (UnionTime) union() {}

type ReaderStruct struct {
	Files []io.Reader `form:"files"`
}

type nameReader struct {
	io.Reader
	name string
}

func (r *nameReader) Name() string { return r.name }

type filenameReader struct {
	io.Reader
	filename    string
	contentType string
}

func (r *filenameReader) Filename() string    { return r.filename }
func (r *filenameReader) ContentType() string { return r.contentType }

var tests = map[string]struct {
	buf string
	val interface{}
}{
	"map_string": {
		`--xxx
Content-Disposition: form-data; name="foo"

bar
--xxx--
`,
		map[string]string{"foo": "bar"},
	},

	"map_interface": {
		`--xxx
Content-Disposition: form-data; name="a"

1
--xxx
Content-Disposition: form-data; name="b"

str
--xxx
Content-Disposition: form-data; name="c"

false
--xxx--
`,
		map[string]interface{}{"a": float64(1), "b": "str", "c": false},
	},

	"primitive_struct": {
		`--xxx
Content-Disposition: form-data; name="a"

false
--xxx
Content-Disposition: form-data; name="b"

237628372683
--xxx
Content-Disposition: form-data; name="c"

654
--xxx
Content-Disposition: form-data; name="d"

9999.43
--xxx
Content-Disposition: form-data; name="e"

43.76
--xxx
Content-Disposition: form-data; name="f.0"

1
--xxx
Content-Disposition: form-data; name="f.1"

2
--xxx
Content-Disposition: form-data; name="f.2"

3
--xxx
Content-Disposition: form-data; name="f.3"

4
--xxx--
`,
		Primitives{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
	},

	"slices": {
		`--xxx
Content-Disposition: form-data; name="slices.0.a"

false
--xxx
Content-Disposition: form-data; name="slices.0.b"

237628372683
--xxx
Content-Disposition: form-data; name="slices.0.c"

654
--xxx
Content-Disposition: form-data; name="slices.0.d"

9999.43
--xxx
Content-Disposition: form-data; name="slices.0.e"

43.76
--xxx
Content-Disposition: form-data; name="slices.0.f.0"

1
--xxx
Content-Disposition: form-data; name="slices.0.f.1"

2
--xxx
Content-Disposition: form-data; name="slices.0.f.2"

3
--xxx
Content-Disposition: form-data; name="slices.0.f.3"

4
--xxx--
`,
		Slices{
			Slice: []Primitives{{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}}},
		},
	},

	"primitive_pointer_struct": {
		`--xxx
Content-Disposition: form-data; name="a"

false
--xxx
Content-Disposition: form-data; name="b"

237628372683
--xxx
Content-Disposition: form-data; name="c"

654
--xxx
Content-Disposition: form-data; name="d"

9999.43
--xxx
Content-Disposition: form-data; name="e"

43.76
--xxx
Content-Disposition: form-data; name="f.0"

1
--xxx
Content-Disposition: form-data; name="f.1"

2
--xxx
Content-Disposition: form-data; name="f.2"

3
--xxx
Content-Disposition: form-data; name="f.3"

4
--xxx
Content-Disposition: form-data; name="f.4"

5
--xxx--
`,
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
		`--xxx
Content-Disposition: form-data; name="date"

2006-01-02
--xxx
Content-Disposition: form-data; name="date-time"

2006-01-02T15:04:05Z
--xxx--
`,
		DateTime{
			Date:     time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
			DateTime: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
		},
	},

	"additional_properties": {
		`--xxx
Content-Disposition: form-data; name="a"

true
--xxx
Content-Disposition: form-data; name="bar"

value
--xxx
Content-Disposition: form-data; name="foo"

true
--xxx--
`,
		AdditionalProperties{
			A: true,
			Extras: map[string]interface{}{
				"bar": "value",
				"foo": true,
			},
		},
	},

	"recursive_struct": {
		`--xxx
Content-Disposition: form-data; name="child.name"

Alex
--xxx
Content-Disposition: form-data; name="name"

Robert
--xxx--
`,
		Recursive{Name: "Robert", Child: &Recursive{Name: "Alex"}},
	},

	"unknown_struct_number": {
		`--xxx
Content-Disposition: form-data; name="unknown"

12
--xxx--
`,
		UnknownStruct{
			Unknown: 12.,
		},
	},

	"unknown_struct_map": {
		`--xxx
Content-Disposition: form-data; name="unknown.foo"

bar
--xxx--
`,
		UnknownStruct{
			Unknown: map[string]interface{}{
				"foo": "bar",
			},
		},
	},

	"union_integer": {
		`--xxx
Content-Disposition: form-data; name="union"

12
--xxx--
`,
		UnionStruct{
			Union: UnionInteger(12),
		},
	},

	"union_struct_discriminated_a": {
		`--xxx
Content-Disposition: form-data; name="union.a"

foo
--xxx
Content-Disposition: form-data; name="union.b"

bar
--xxx
Content-Disposition: form-data; name="union.type"

typeA
--xxx--
`,

		UnionStruct{
			Union: UnionStructA{
				Type: "typeA",
				A:    "foo",
				B:    "bar",
			},
		},
	},

	"union_struct_discriminated_b": {
		`--xxx
Content-Disposition: form-data; name="union.a"

foo
--xxx
Content-Disposition: form-data; name="union.type"

typeB
--xxx--
`,
		UnionStruct{
			Union: UnionStructB{
				Type: "typeB",
				A:    "foo",
			},
		},
	},

	"union_struct_time": {
		`--xxx
Content-Disposition: form-data; name="union"

2010-05-23
--xxx--
`,
		UnionStruct{
			Union: UnionTime(time.Date(2010, 05, 23, 0, 0, 0, 0, time.UTC)),
		},
	},
}

func TestEncode(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			writer := multipart.NewWriter(buf)
			writer.SetBoundary("xxx")
			err := Marshal(test.val, writer)
			if err != nil {
				t.Errorf("serialization of %v failed with error %v", test.val, err)
			}
			err = writer.Close()
			if err != nil {
				t.Errorf("serialization of %v failed with error %v", test.val, err)
			}
			raw := buf.Bytes()
			if string(raw) != strings.ReplaceAll(test.buf, "\n", "\r\n") {
				t.Errorf("expected %+#v to serialize to '%s' but got '%s'", test.val, test.buf, string(raw))
			}
		})
	}
}

func TestEncodeReaderFilenames(t *testing.T) {
	t.Run("preserves subdirectory paths for Name() and Filename() readers", func(t *testing.T) {
		val := ReaderStruct{
			Files: []io.Reader{
				&nameReader{Reader: strings.NewReader("first"), name: "worker/index.js"},
				&filenameReader{Reader: strings.NewReader("second"), filename: "worker/lib/api.js", contentType: "application/javascript+module"},
				strings.NewReader("third"),
			},
		}

		buf := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(buf)
		writer.SetBoundary("xxx")
		if err := Marshal(val, writer); err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		if err := writer.Close(); err != nil {
			t.Fatalf("Close failed: %v", err)
		}

		want := strings.ReplaceAll(`--xxx
Content-Disposition: form-data; name="files.0"; filename="worker/index.js"
Content-Type: application/octet-stream

first
--xxx
Content-Disposition: form-data; name="files.1"; filename="worker/lib/api.js"
Content-Type: application/javascript+module

second
--xxx
Content-Disposition: form-data; name="files.2"; filename="anonymous_file"
Content-Type: application/octet-stream

third
--xxx--
`, "\n", "\r\n")

		if got := buf.String(); got != want {
			t.Errorf("got:\n%s\nwant:\n%s", got, want)
		}

		reader := multipart.NewReader(bytes.NewReader(buf.Bytes()), "xxx")
		wantFilenames := []string{"worker/index.js", "worker/lib/api.js", "anonymous_file"}
		for i, wantFilename := range wantFilenames {
			part, err := reader.NextPart()
			if err != nil {
				t.Fatalf("part %d: %v", i, err)
			}
			_, dispParams, err := mime.ParseMediaType(part.Header.Get("Content-Disposition"))
			if err != nil {
				t.Fatalf("part %d ParseMediaType: %v", i, err)
			}
			if dispParams["filename"] != wantFilename {
				t.Errorf("part %d filename = %q, want %q", i, dispParams["filename"], wantFilename)
			}
		}
	})

	t.Run("preserves relative path for os.Open file", func(t *testing.T) {
		tmp := t.TempDir()
		t.Chdir(tmp)
		if err := os.MkdirAll("worker", 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile("worker/index.js", []byte("export default {}"), 0o644); err != nil {
			t.Fatal(err)
		}
		f, err := os.Open("worker/index.js")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		buf := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(buf)
		writer.SetBoundary("xxx")
		if err := Marshal(ReaderStruct{Files: []io.Reader{f}}, writer); err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		if err := writer.Close(); err != nil {
			t.Fatalf("Close failed: %v", err)
		}

		wantSubstr := `Content-Disposition: form-data; name="files.0"; filename="worker/index.js"`
		if !strings.Contains(buf.String(), wantSubstr) {
			t.Errorf("expected %q in:\n%s", wantSubstr, buf.String())
		}
	})
}
