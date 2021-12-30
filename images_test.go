package cloudflare

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func timeMustParse(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}

var expectedImageStruct = Image{
	ID:       "ZxR0pLaXRldlBtaFhhO2FiZGVnaA",
	Filename: "avatar.png",
	Metadata: map[string]interface{}{
		"meta": "metaID",
	},
	RequireSignedURLs: true,
	Variants: []string{
		"https://imagedelivery.net/MTt4OTd0b0w5aj/ZxR0pLaXRldlBtaFhhO2FiZGVnaA/hero",
		"https://imagedelivery.net/MTt4OTd0b0w5aj/ZxR0pLaXRldlBtaFhhO2FiZGVnaA/original",
		"https://imagedelivery.net/MTt4OTd0b0w5aj/ZxR0pLaXRldlBtaFhhO2FiZGVnaA/thumbnail",
	},
	Uploaded: timeMustParse(time.RFC3339, "2014-01-02T02:20:00Z"),
}

func TestUploadImage(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)

		u, err := parseImageMultipartUpload(r)
		if !assert.NoError(t, err) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		assert.Equal(t, u.RequireSignedURLs, true)
		assert.Equal(t, u.Metadata, map[string]interface{}{"meta": "metaID"})
		assert.Equal(t, u.File, []byte("this is definitely an image"))

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "ZxR0pLaXRldlBtaFhhO2FiZGVnaA",
					"filename": "avatar.png",
					"metadata": {
						"meta": "metaID"
					},
					"requireSignedURLs": true,
					"variants": [
						"https://imagedelivery.net/MTt4OTd0b0w5aj/ZxR0pLaXRldlBtaFhhO2FiZGVnaA/hero",
						"https://imagedelivery.net/MTt4OTd0b0w5aj/ZxR0pLaXRldlBtaFhhO2FiZGVnaA/original",
						"https://imagedelivery.net/MTt4OTd0b0w5aj/ZxR0pLaXRldlBtaFhhO2FiZGVnaA/thumbnail"
					],
					"uploaded": "2014-01-02T02:20:00Z"
				}
			]
		}
		`)
	}

	mux.HandleFunc("/accounts/foo/images/v1", handler)
	want := expectedImageStruct

	actual, err := client.UploadImage(context.Background(), client.AccountID, ImageUploadRequest{
		File: fakeFile{
			Buffer: bytes.NewBufferString("this is definitely an image"),
		},
		Name: "avatar.png",
		RequireSignedURLs: true,
		Metadata: map[string]interface{}{
			"meta": "metaID",
		},
	})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestImages(t *testing.T) {
	setup(UsingAccount("foo"))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "ZxR0pLaXRldlBtaFhhO2FiZGVnaA",
					"filename": "avatar.png",
					"metadata": {
						"meta": "metaID"
					},
					"requireSignedURLs": true,
					"variants": [
						"https://imagedelivery.net/MTt4OTd0b0w5aj/ZxR0pLaXRldlBtaFhhO2FiZGVnaA/hero",
						"https://imagedelivery.net/MTt4OTd0b0w5aj/ZxR0pLaXRldlBtaFhhO2FiZGVnaA/original",
						"https://imagedelivery.net/MTt4OTd0b0w5aj/ZxR0pLaXRldlBtaFhhO2FiZGVnaA/thumbnail"
					],
					"uploaded": "2014-01-02T02:20:00Z"
				}
			]
		}
		`)
	}

	mux.HandleFunc("/accounts/foo/images/v1", handler)
	want := []Image{expectedImageStruct}

	actual, err := client.Images(context.Background(), client.AccountID, PaginationOptions{})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

type fakeFile struct {
	*bytes.Buffer
}

func (f fakeFile) Close() error {
	return nil
}

type imageMultipartUpload struct {
	// this is for testing, never read an entire file into memory,
	// especially when being done on a per-http request basis.
	File              []byte
	RequireSignedURLs bool
	Metadata          map[string]interface{}
}

func parseImageMultipartUpload(r *http.Request) (imageMultipartUpload, error) {
	var u imageMultipartUpload
	mdBytes, err := getImageFormValue(r, "metadata")
	if err != nil {
		if !strings.HasPrefix(err.Error(), "no value found for key") {
			return u, err
		}
	}
	if mdBytes != nil {
		err = json.Unmarshal(mdBytes, &u.Metadata)
		if err != nil {
			return u, err
		}
	}

	rsuBytes, err := getImageFormValue(r, "requireSignedURLs")
	if err != nil {
		if !strings.HasPrefix(err.Error(), "no value found for key") {
			return u, err
		}
	}
	if rsuBytes != nil {
		if bytes.Equal(rsuBytes, []byte("true")) {
			u.RequireSignedURLs = true
		}
	}

	f, _, err := r.FormFile("file")
	if err != nil {
		return u, err
	}
	defer f.Close()

	u.File, err = ioutil.ReadAll(f)
	if err != nil {
		return u, err
	}

	return u, nil
}

// See getFormValue for more information, the only difference between
// getFormValue and this one is the max memory.
func getImageFormValue(r *http.Request, key string) ([]byte, error) {
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		return nil, err
	}

	if values, ok := r.MultipartForm.Value[key]; ok {
		return []byte(values[0]), nil
	}

	if fileHeaders, ok := r.MultipartForm.File[key]; ok {
		file, err := fileHeaders[0].Open()
		if err != nil {
			return nil, err
		}
		return ioutil.ReadAll(file)
	}

	return nil, fmt.Errorf("no value found for key %v", key)
}
