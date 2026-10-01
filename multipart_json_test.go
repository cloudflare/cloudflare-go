package cloudflare_test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"reflect"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/snippets"
	"github.com/cloudflare/cloudflare-go/v7/workers"
	"github.com/cloudflare/cloudflare-go/v7/workers_for_platforms"
)

func TestMultipartJSONMetadata(t *testing.T) {
	type multipartParams interface {
		MarshalMultipart() ([]byte, string, error)
	}
	for _, tc := range []struct {
		name     string
		params   multipartParams
		partName string
		want     map[string]interface{}
		files    []string
	}{
		{
			name: "worker settings",
			params: workers.ScriptScriptAndVersionSettingEditParams{
				AccountID: cloudflare.F("account"),
				Settings: cloudflare.F(workers.ScriptScriptAndVersionSettingEditParamsSettings{
					CompatibilityDate: cloudflare.F("2026-01-01"),
					Bindings: cloudflare.F([]workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingUnion{
						workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainText{
							Name: cloudflare.F("NAME"), Text: cloudflare.F("value"),
							Type: cloudflare.F(workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypePlainText),
						},
					}),
				}),
			},
			partName: "settings",
			want:     map[string]interface{}{"compatibility_date": "2026-01-01", "bindings": []interface{}{map[string]interface{}{"name": "NAME", "text": "value", "type": "plain_text"}}},
		},
		{
			name: "dispatch settings",
			params: workers_for_platforms.DispatchNamespaceScriptSettingEditParams{
				AccountID: cloudflare.F("account"),
				Settings: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettings{
					CompatibilityDate: cloudflare.F("2026-01-01"),
				}),
			},
			partName: "settings", want: map[string]interface{}{"compatibility_date": "2026-01-01"},
		},
		{
			name: "snippet metadata",
			params: snippets.SnippetUpdateParams{
				ZoneID:   cloudflare.F("zone"),
				Metadata: cloudflare.F(snippets.SnippetUpdateParamsMetadata{MainModule: cloudflare.F("index.js")}),
			},
			partName: "metadata", want: map[string]interface{}{"main_module": "index.js"},
		},
		{
			name: "worker content",
			params: workers.ScriptContentUpdateParams{
				AccountID: cloudflare.F("account"),
				Metadata:  cloudflare.F(workers.ScriptContentUpdateParamsMetadata{MainModule: cloudflare.F("index.js")}),
				Files:     cloudflare.F([]io.Reader{strings.NewReader("first"), strings.NewReader("second")}),
			},
			partName: "metadata", want: map[string]interface{}{"main_module": "index.js"}, files: []string{"first", "second"},
		},
		{
			name: "dispatch content",
			params: workers_for_platforms.DispatchNamespaceScriptContentUpdateParams{
				AccountID: cloudflare.F("account"),
				Metadata:  cloudflare.F(workers.WorkerMetadataParam{MainModule: cloudflare.F("index.js")}),
				Files:     cloudflare.F([]io.Reader{strings.NewReader("first"), strings.NewReader("second")}),
			},
			partName: "metadata", want: map[string]interface{}{"main_module": "index.js"}, files: []string{"first", "second"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			data, contentType, err := tc.params.MarshalMultipart()
			if err != nil {
				t.Fatal(err)
			}
			_, params, err := mime.ParseMediaType(contentType)
			if err != nil {
				t.Fatal(err)
			}
			reader := multipart.NewReader(bytes.NewReader(data), params["boundary"])
			seenJSON := false
			var files []string
			for {
				part, err := reader.NextPart()
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Fatal(err)
				}
				body, err := io.ReadAll(part)
				if err != nil {
					t.Fatal(err)
				}
				switch part.FormName() {
				case tc.partName:
					if seenJSON {
						t.Fatal("duplicate JSON part")
					}
					seenJSON = true
					if got := part.Header.Get("Content-Type"); got != "application/json" {
						t.Errorf("content type = %q", got)
					}
					var got map[string]interface{}
					if err := json.Unmarshal(body, &got); err != nil {
						t.Fatal(err)
					}
					if !reflect.DeepEqual(got, tc.want) {
						t.Errorf("JSON = %#v, want %#v", got, tc.want)
					}
				case "files":
					if part.FileName() != "anonymous_file" {
						t.Errorf("filename = %q", part.FileName())
					}
					files = append(files, string(body))
				default:
					t.Errorf("unexpected part %q", part.FormName())
				}
			}
			if !seenJSON {
				t.Errorf("missing %q JSON part", tc.partName)
			}
			if !reflect.DeepEqual(files, tc.files) {
				t.Errorf("files = %#v, want %#v", files, tc.files)
			}
		})
	}
}
