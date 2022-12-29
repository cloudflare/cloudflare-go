package cloudflare

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
	"time"
)

// WorkerRequestParams provides parameters for worker requests for both enterprise and standard requests.
type WorkerRequestParams struct {
	ZoneID     string
	ScriptName string
}

type CreateWorkerParams struct {
	ScriptName string
	Script     string

	// Module changes the Content-Type header to specify the script is an
	// ES Module syntax script.
	Module bool

	// Logpush opts the worker into Workers Logpush logging. A nil value leaves the current setting unchanged.
	//  https://developers.cloudflare.com/workers/platform/logpush/
	Logpush *bool

	// Bindings should be a map where the keys are the binding name, and the
	// values are the binding content
	Bindings map[string]WorkerBinding
}

// WorkerScriptParams provides a worker script and the associated bindings.
type WorkerScriptParams struct {
	ScriptName string

	// Module changes the Content-Type header to specify the script is an
	// ES Module syntax script.
	Module bool

	// Bindings should be a map where the keys are the binding name, and the
	// values are the binding content
	Bindings map[string]WorkerBinding
}

// WorkerRoute is used to map traffic matching a URL pattern to a workers
//
// API reference: https://api.cloudflare.com/#worker-routes-properties
type WorkerRoute struct {
	ID         string `json:"id,omitempty"`
	Pattern    string `json:"pattern"`
	ScriptName string `json:"script,omitempty"`
}

// WorkerRoutesResponse embeds Response struct and slice of WorkerRoutes.
type WorkerRoutesResponse struct {
	Response
	Routes []WorkerRoute `json:"result"`
}

// WorkerRouteResponse embeds Response struct and a single WorkerRoute.
type WorkerRouteResponse struct {
	Response
	WorkerRoute `json:"result"`
}

// WorkerScript Cloudflare Worker struct with metadata.
type WorkerScript struct {
	WorkerMetaData
	Script     string `json:"script"`
	UsageModel string `json:"usage_model,omitempty"`
}

// WorkerMetaData contains worker script information such as size, creation & modification dates.
type WorkerMetaData struct {
	ID         string    `json:"id,omitempty"`
	ETAG       string    `json:"etag,omitempty"`
	Size       int       `json:"size,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}

// WorkerListResponse wrapper struct for API response to worker script list API call.
type WorkerListResponse struct {
	Response
	ResultInfo
	WorkerList []WorkerMetaData `json:"result"`
}

// WorkerScriptResponse wrapper struct for API response to worker script calls.
type WorkerScriptResponse struct {
	Response
	Module       bool
	WorkerScript `json:"result"`
}

type ListWorkersParams struct{}

type DeleteWorkerParams struct {
	ScriptName string
}

// DeleteWorker deletes a single Worker.
//
// API reference: https://api.cloudflare.com/#worker-script-delete-worker
func (api *API) DeleteWorker(ctx context.Context, rc *ResourceContainer, params DeleteWorkerParams) error {
	if rc.Level != AccountRouteLevel {
		return ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s", rc.Identifier, params.ScriptName)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)

	var r WorkerScriptResponse
	if err != nil {
		return err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}

// GetWorker fetch raw script content for your worker returns string containing
// worker code js.
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) GetWorker(ctx context.Context, rc *ResourceContainer, scriptName string) (WorkerScriptResponse, error) {
	if rc.Level != AccountRouteLevel {
		return WorkerScriptResponse{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return WorkerScriptResponse{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s", rc.Identifier, scriptName)
	res, err := api.makeRequestContextWithHeadersComplete(ctx, http.MethodGet, uri, nil, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}

	// Check if the response type is multipart, in which case this was a module worker
	mediaType, mediaParams, _ := mime.ParseMediaType(res.Headers.Get("content-type"))
	if strings.HasPrefix(mediaType, "multipart/") {
		bytesReader := bytes.NewReader(res.Body)
		mimeReader := multipart.NewReader(bytesReader, mediaParams["boundary"])
		mimePart, err := mimeReader.NextPart()
		if err != nil {
			return r, fmt.Errorf("could not get multipart response body: %w", err)
		}
		mimePartBody, err := io.ReadAll(mimePart)
		if err != nil {
			return r, fmt.Errorf("could not read multipart response body: %w", err)
		}
		r.Script = string(mimePartBody)
		r.Module = true
	} else {
		r.Script = string(res.Body)
		r.Module = false
	}

	r.Success = true
	return r, nil
}

// ListWorkers returns list of Workers for given account.
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) ListWorkers(ctx context.Context, rc *ResourceContainer, params ListWorkersParams) (WorkerListResponse, *ResultInfo, error) {
	if rc.Level != AccountRouteLevel {
		return WorkerListResponse{}, &ResultInfo{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return WorkerListResponse{}, &ResultInfo{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkerListResponse{}, &ResultInfo{}, err
	}

	var r WorkerListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerListResponse{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r, &r.ResultInfo, nil
}

// UploadWorker pushes raw script content for your Worker.
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorker(ctx context.Context, rc *ResourceContainer, params CreateWorkerParams) (WorkerScriptResponse, error) {
	if rc.Level != AccountRouteLevel {
		return WorkerScriptResponse{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return WorkerScriptResponse{}, ErrMissingAccountID
	}

	body := []byte(params.Script)
	var (
		contentType = "application/javascript"
		err         error
	)

	if params.Module || params.Logpush != nil || len(params.Bindings) > 0 {
		contentType, body, err = formatMultipartBody(params)
		if err != nil {
			return WorkerScriptResponse{}, err
		}
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s", rc.Identifier, params.ScriptName)
	headers := make(http.Header)
	headers.Set("Content-Type", contentType)
	res, err := api.makeRequestContextWithHeaders(ctx, http.MethodPut, uri, body, headers)

	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r, nil
}

// Returns content-type, body, error.
func formatMultipartBody(params CreateWorkerParams) (string, []byte, error) {
	var buf = &bytes.Buffer{}
	var mpw = multipart.NewWriter(buf)
	defer mpw.Close()

	// Write metadata part
	var scriptPartName string
	meta := struct {
		BodyPart   string              `json:"body_part,omitempty"`
		MainModule string              `json:"main_module,omitempty"`
		Bindings   []workerBindingMeta `json:"bindings"`
		Logpush    *bool               `json:"logpush,omitempty"`
	}{
		Bindings: make([]workerBindingMeta, 0, len(params.Bindings)),
		Logpush:  params.Logpush,
	}

	if params.Module {
		scriptPartName = "worker.mjs"
		meta.MainModule = scriptPartName
	} else {
		scriptPartName = "script"
		meta.BodyPart = scriptPartName
	}

	bodyWriters := make([]workerBindingBodyWriter, 0, len(params.Bindings))
	for name, b := range params.Bindings {
		bindingMeta, bodyWriter, err := b.serialize(name)
		if err != nil {
			return "", nil, err
		}

		meta.Bindings = append(meta.Bindings, bindingMeta)
		bodyWriters = append(bodyWriters, bodyWriter)
	}

	var hdr = textproto.MIMEHeader{}
	hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, "metadata"))
	hdr.Set("content-type", "application/json")
	pw, err := mpw.CreatePart(hdr)
	if err != nil {
		return "", nil, err
	}
	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return "", nil, err
	}
	_, err = pw.Write(metaJSON)
	if err != nil {
		return "", nil, err
	}

	// Write script part
	hdr = textproto.MIMEHeader{}

	contentType := "application/javascript"
	if params.Module {
		contentType = "application/javascript+module"
		hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"; filename="%[1]s"`, scriptPartName))
	} else {
		hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, scriptPartName))
	}
	hdr.Set("content-type", contentType)

	pw, err = mpw.CreatePart(hdr)
	if err != nil {
		return "", nil, err
	}
	_, err = pw.Write([]byte(params.Script))
	if err != nil {
		return "", nil, err
	}

	// Write other bindings with parts
	for _, w := range bodyWriters {
		if w != nil {
			err = w(mpw)
			if err != nil {
				return "", nil, err
			}
		}
	}

	mpw.Close()

	return mpw.FormDataContentType(), buf.Bytes(), nil
}
