package workers_for_platforms_test

import (
	"bytes"
	"io"
	"mime"
	"mime/multipart"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/workers"
	"github.com/cloudflare/cloudflare-go/v7/workers_for_platforms"
)

type namedReader struct {
	io.Reader
	filename string
}

func (r *namedReader) Name() string {
	return r.filename
}

func TestDispatchNamespaceScriptSubdirectoryFilenames(t *testing.T) {
	scriptBody := `export default { fetch() { return new Response("ok"); } };`

	t.Run("DispatchNamespaceScriptUpdateParams with Name() reader", func(t *testing.T) {
		params := workers_for_platforms.DispatchNamespaceScriptUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Metadata: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadata{
				MainModule: cloudflare.F("worker/index.js"),
			}),
			Files: cloudflare.F([]io.Reader{
				&namedReader{Reader: strings.NewReader(scriptBody), filename: "worker/index.js"},
			}),
		}

		data, contentType, err := params.MarshalMultipart()
		if err != nil {
			t.Fatalf("MarshalMultipart failed: %v", err)
		}

		if !bytes.Contains(data, []byte(`Content-Disposition: form-data; name="files.0"; filename="worker/index.js"`)) {
			t.Fatalf("expected wire output to preserve subdirectory path in Content-Disposition, got:\n%s", string(data))
		}

		_, mediaParams, err := mime.ParseMediaType(contentType)
		if err != nil {
			t.Fatalf("ParseMediaType failed: %v", err)
		}
		reader := multipart.NewReader(bytes.NewReader(data), mediaParams["boundary"])
		var gotFilename, gotName, gotContentType, gotBody string
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatalf("NextPart failed: %v", err)
			}
			body, err := io.ReadAll(part)
			if err != nil {
				t.Fatalf("ReadAll failed: %v", err)
			}
			_, dispParams, err := mime.ParseMediaType(part.Header.Get("Content-Disposition"))
			if err != nil {
				t.Fatalf("ParseMediaType(Content-Disposition) failed: %v", err)
			}
			if dispParams["name"] == "files.0" {
				gotName = dispParams["name"]
				gotFilename = dispParams["filename"]
				gotContentType = part.Header.Get("Content-Type")
				gotBody = string(body)
			}
		}
		if gotName != "files.0" {
			t.Errorf("part name = %q, want %q", gotName, "files.0")
		}
		if gotFilename != "worker/index.js" {
			t.Errorf("part filename = %q, want %q", gotFilename, "worker/index.js")
		}
		if gotContentType != "application/octet-stream" {
			t.Errorf("part Content-Type = %q, want %q", gotContentType, "application/octet-stream")
		}
		if gotBody != scriptBody {
			t.Errorf("part body = %q, want %q", gotBody, scriptBody)
		}
	})

	t.Run("DispatchNamespaceScriptUpdateParams with FileParam", func(t *testing.T) {
		params := workers_for_platforms.DispatchNamespaceScriptUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Metadata: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadata{
				MainModule: cloudflare.F("worker/index.js"),
			}),
			Files: cloudflare.F([]io.Reader{
				cloudflare.FileParam(strings.NewReader(scriptBody), "worker/index.js", "application/javascript+module").Value,
			}),
		}

		data, contentType, err := params.MarshalMultipart()
		if err != nil {
			t.Fatalf("MarshalMultipart failed: %v", err)
		}

		if !bytes.Contains(data, []byte(`Content-Disposition: form-data; name="files.0"; filename="worker/index.js"`)) {
			t.Fatalf("expected wire output to preserve subdirectory path in Content-Disposition, got:\n%s", string(data))
		}

		_, mediaParams, err := mime.ParseMediaType(contentType)
		if err != nil {
			t.Fatalf("ParseMediaType failed: %v", err)
		}
		reader := multipart.NewReader(bytes.NewReader(data), mediaParams["boundary"])
		var gotFilename, gotContentType string
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatalf("NextPart failed: %v", err)
			}
			_, dispParams, err := mime.ParseMediaType(part.Header.Get("Content-Disposition"))
			if err != nil {
				t.Fatalf("ParseMediaType(Content-Disposition) failed: %v", err)
			}
			if dispParams["name"] == "files.0" {
				gotFilename = dispParams["filename"]
				gotContentType = part.Header.Get("Content-Type")
			}
		}
		if gotFilename != "worker/index.js" {
			t.Errorf("part filename = %q, want %q", gotFilename, "worker/index.js")
		}
		if gotContentType != "application/javascript+module" {
			t.Errorf("part Content-Type = %q, want %q", gotContentType, "application/javascript+module")
		}
	})

	t.Run("DispatchNamespaceScriptContentUpdateParams with Name() and FileParam readers", func(t *testing.T) {
		params := workers_for_platforms.DispatchNamespaceScriptContentUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Metadata: cloudflare.F(workers.WorkerMetadataParam{
				MainModule: cloudflare.F("worker/index.js"),
			}),
			Files: cloudflare.F([]io.Reader{
				&namedReader{Reader: strings.NewReader(scriptBody), filename: "worker/index.js"},
				cloudflare.FileParam(strings.NewReader("export const x = 1;"), "worker/lib/api.js", "application/javascript+module").Value,
			}),
		}

		data, contentType, err := params.MarshalMultipart()
		if err != nil {
			t.Fatalf("MarshalMultipart failed: %v", err)
		}

		_, mediaParams, err := mime.ParseMediaType(contentType)
		if err != nil {
			t.Fatalf("ParseMediaType failed: %v", err)
		}
		reader := multipart.NewReader(bytes.NewReader(data), mediaParams["boundary"])
		var filenames []string
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatalf("NextPart failed: %v", err)
			}
			_, dispParams, err := mime.ParseMediaType(part.Header.Get("Content-Disposition"))
			if err != nil {
				t.Fatalf("ParseMediaType(Content-Disposition) failed: %v", err)
			}
			if dispParams["name"] == "files" {
				filenames = append(filenames, dispParams["filename"])
			}
		}
		if len(filenames) != 2 || filenames[0] != "worker/index.js" || filenames[1] != "worker/lib/api.js" {
			t.Errorf("filenames = %v, want [worker/index.js worker/lib/api.js]", filenames)
		}
	})
}
