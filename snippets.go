package cloudflare

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type SnippetsResponse struct {
	Response
	Result []Snippet `json:"result"`
}

type Snippet struct {
	CreatedOn   *time.Time `json:"created_on"`
	ModifiedOn  *time.Time `json:"modified_on"`
	SnippetName string     `json:"snippet_name"`
}

type SnippetFile struct {
	FileName string `json:"file_name"`
	Content  string `json:"content"`
}

type SnippetRequest struct {
	SnippetName string        `json:"snippet_name"`
	MainFile    string        `json:"main_file"`
	Files       []SnippetFile `json:"files"`
}

func (api *API) ListZoneSnippets(ctx context.Context, rc *ResourceContainer) ([]Snippet, error) {
	if rc.Identifier == "" {
		return nil, ErrMissingZoneID
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/snippets", rc.Identifier), nil)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	result := SnippetsResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

type SnippetMetadata struct {
	MainModule string `json:"main_module"`
}

func snippetMultipartBody(request SnippetRequest) (string, *bytes.Buffer, error) {
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	defer mw.Close()

	for _, file := range request.Files {
		tp, err := mw.CreateFormFile(file.FileName, file.FileName)
		if err != nil {
			return "", nil, err
		}
		_, err = tp.Write([]byte(file.Content))
		if err != nil {
			return "", nil, err
		}
	}

	tp, err := mw.CreateFormField("metadata")
	if err != nil {
		return "", nil, err
	}

	if err = json.NewEncoder(tp).Encode(SnippetMetadata{
		MainModule: request.MainFile,
	}); err != nil {
		return "", nil, err
	}

	return mw.Boundary(), body, nil
}

func (api *API) UpdateZoneSnippet(ctx context.Context, rc *ResourceContainer, params SnippetRequest) ([]Snippet, error) {
	if rc.Identifier == "" {
		return nil, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/snippets/%s", rc.Identifier, params.SnippetName)

	boundary, body, err := snippetMultipartBody(params)
	if err != nil {
		return nil, err
	}

	res, err := api.makeRequestContextWithHeaders(ctx, http.MethodPut, uri, body, http.Header{
		"Content-Type": []string{fmt.Sprintf("multipart/form-data; boundary=%s", boundary)},
	})
	if err != nil {
		return nil, err
	}

	result := SnippetsResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}
