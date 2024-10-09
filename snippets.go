package cloudflare

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/goccy/go-json"
)

type Snippet struct {
	CreatedOn   string `json:"created_on"`
	ModifiedOn  string `json:"modified_on"`
	SnippetName string `json:"snippet_name"`
}

type ListSnippetsParams struct{}

type ListSnippetsResponse struct {
	Response
	ResultInfo
	Snippets []Snippet `json:"result"`
}

// ListSnippets returns a list of Snippets for a given zone.
//
// API reference: https://developers.cloudflare.com/api/operations/zone-snippets
func (api *API) ListSnippets(ctx context.Context, rc *ResourceContainer, params ListSnippetsParams) (ListSnippetsResponse, *ResultInfo, error) {
	if rc.Level != ZoneRouteLevel {
		return ListSnippetsResponse{}, &ResultInfo{}, ErrRequiredZoneLevelResourceContainer
	}

	if rc.Identifier == "" {
		return ListSnippetsResponse{}, &ResultInfo{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/snippets", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ListSnippetsResponse{}, &ResultInfo{}, err
	}

	var r ListSnippetsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return ListSnippetsResponse{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r, &r.ResultInfo, nil
}

type GetSnippetParams struct {
	SnippetName string
}

type GetSnippetResponse struct {
	Response
	Snippet Snippet `json:"result"`
}

// GetSnippet returns a single Snippet for a given zone.
//
// API reference: https://developers.cloudflare.com/api/operations/zone-snippets-snippet
func (api *API) GetSnippet(ctx context.Context, rc *ResourceContainer, params GetSnippetParams) (GetSnippetResponse, error) {
	if rc.Level != ZoneRouteLevel {
		return GetSnippetResponse{}, ErrRequiredZoneLevelResourceContainer
	}

	if rc.Identifier == "" {
		return GetSnippetResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/snippets/%s", rc.Identifier, params.SnippetName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return GetSnippetResponse{}, err
	}

	var r GetSnippetResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return GetSnippetResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r, nil
}

type DeleteSnippetParams struct {
	SnippetName string
}

type DeleteSnippetResponse struct {
	Response
}

// DeleteSnippet deletes a single Snippet for a given zone.
//
// API reference: https://developers.cloudflare.com/api/operations/zone-snippets-snippet
func (api *API) DeleteSnippet(ctx context.Context, rc *ResourceContainer, params DeleteSnippetParams) (DeleteSnippetResponse, error) {
	if rc.Level != ZoneRouteLevel {
		return DeleteSnippetResponse{}, ErrRequiredZoneLevelResourceContainer
	}

	if rc.Identifier == "" {
		return DeleteSnippetResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/snippets/%s", rc.Identifier, params.SnippetName)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return DeleteSnippetResponse{}, err
	}

	var r DeleteSnippetResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return DeleteSnippetResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r, nil
}

type PutSnippetParams struct {
	SnippetName string
	Files       map[string]string
	MainModule  string
}

type PutSnippetParamsMetadata struct {
	MainModule string `json:"main_module"`
}

type PutSnippetResponse struct {
	Response
	Snippet Snippet `json:"result"`
}

// PutSnippet updates a Snippet for a given zone.
//
// API reference: https://developers.cloudflare.com/api/operations/zone-snippets-snippet
func (api *API) PutSnippet(ctx context.Context, rc *ResourceContainer, params PutSnippetParams) (PutSnippetResponse, error) {
	if rc.Level != ZoneRouteLevel {
		return PutSnippetResponse{}, ErrRequiredZoneLevelResourceContainer
	}

	if rc.Identifier == "" {
		return PutSnippetResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/snippets/%s", rc.Identifier, params.SnippetName)

	header, body, err := formatSnippetMultipartBody(params)
	if err != nil {
		return PutSnippetResponse{}, err
	}

	headers := make(http.Header)
	headers.Set("Content-Type", header)

	res, err := api.makeRequestContextWithHeaders(ctx, http.MethodPut, uri, body, headers)
	if err != nil {
		return PutSnippetResponse{}, err
	}

	var r PutSnippetResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PutSnippetResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r, nil
}

func formatSnippetMultipartBody(params PutSnippetParams) (string, []byte, error) {
	var buf = &bytes.Buffer{}
	var mpw = multipart.NewWriter(buf)
	defer mpw.Close()

	metadataParam := PutSnippetParamsMetadata{
		MainModule: params.MainModule,
	}

	metadata, err := json.Marshal(metadataParam)
	if err != nil {
		return "", nil, err
	}

	mpw.WriteField("metadata", string(metadata))

	for k, v := range params.Files {
		fw, err := mpw.CreateFormFile(k, k)
		if err != nil {
			return "", nil, err
		}

		_, err = fw.Write([]byte(v))
		if err != nil {
			return "", nil, err
		}
	}

	mpw.Close()

	return mpw.FormDataContentType(), buf.Bytes(), nil
}
