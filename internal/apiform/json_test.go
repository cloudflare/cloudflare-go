package apiform

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"reflect"
	"testing"
)

func TestMarshalRootJSONModeCache(t *testing.T) {
	type metadata struct {
		Name string `json:"name"`
	}
	type payload struct {
		Metadata metadata `json:"metadata"`
	}
	value := payload{Metadata: metadata{Name: "value"}}
	// Both modes use the same reflected types; warming one must not change the other.
	for _, asJSON := range []bool{false, true, false, true} {
		buf := new(bytes.Buffer)
		writer := multipart.NewWriter(buf)
		var err error
		if asJSON {
			err = MarshalRootWithJSON(value, writer)
		} else {
			err = MarshalRoot(value, writer)
		}
		if err != nil {
			t.Fatal(err)
		}
		if err := writer.Close(); err != nil {
			t.Fatal(err)
		}
		reader := multipart.NewReader(buf, writer.Boundary())
		part, err := reader.NextPart()
		if err != nil {
			t.Fatal(err)
		}
		data, err := io.ReadAll(part)
		if err != nil {
			t.Fatal(err)
		}
		if asJSON {
			if part.FormName() != "metadata" || part.Header.Get("Content-Type") != "application/json" {
				t.Fatalf("JSON mode emitted %q with headers %v", part.FormName(), part.Header)
			}
			var got metadata
			if err := json.Unmarshal(data, &got); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, value.Metadata) {
				t.Errorf("metadata = %#v", got)
			}
		} else if part.FormName() != "metadata.name" || string(data) != "value" {
			t.Fatalf("ordinary mode emitted %q = %q", part.FormName(), data)
		}
		if _, err := reader.NextPart(); err != io.EOF {
			t.Errorf("expected end of form, got %v", err)
		}
	}
}
