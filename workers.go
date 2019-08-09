package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// WorkerRequestParams provides parameters for worker requests for both enterprise and standard requests
type WorkerRequestParams struct {
	ZoneID     string
	ScriptName string
}

// WorkerRoute aka filters are patterns used to enable or disable workers that match requests.
//
// API reference: https://api.cloudflare.com/#worker-filters-properties
type WorkerRoute struct {
	ID      string `json:"id,omitempty"`
	Pattern string `json:"pattern"`
	Enabled bool   `json:"enabled"`
	Script  string `json:"script,omitempty"`
}

// WorkerRoutesResponse embeds Response struct and slice of WorkerRoutes
type WorkerRoutesResponse struct {
	Response
	Routes []WorkerRoute `json:"result"`
}

// WorkerRouteResponse embeds Response struct and a single WorkerRoute
type WorkerRouteResponse struct {
	Response
	WorkerRoute `json:"result"`
}

// WorkerScript Cloudflare Worker struct with metadata
type WorkerScript struct {
	WorkerMetaData
	Script string `json:"script"`
}

// WorkerMetaData contains worker script information such as size, creation & modification dates
type WorkerMetaData struct {
	ID         string    `json:"id,omitempty"`
	ETAG       string    `json:"etag,omitempty"`
	Size       int       `json:"size,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}

// WorkerListResponse wrapper struct for API response to worker script list API call
type WorkerListResponse struct {
	Response
	WorkerList []WorkerMetaData `json:"result"`
}

// WorkerScriptResponse wrapper struct for API response to worker script calls
type WorkerScriptResponse struct {
	Response
	WorkerScript `json:"result"`
}

// WorkerResourceBindings defines lists of resources to bind to a worker script using the Resource Binding API
//
// API reference: https://developers.cloudflare.com/workers/api/resource-bindings/
type WorkerResourceBindings struct {
	KVBindings   []*WorkerKVBinding
	WASMBindings []*WorkerWASMBinding
}

// MarshalJSON implements Marshaller interface
// Coerces simplified structure into the correct json format
func (w *WorkerResourceBindings) MarshalJSON() ([]byte, error) {
	kv, err := json.Marshal(w.KVBindings)
	if err != nil {
		return nil, err
	}
	wasm, err := json.Marshal(w.WASMBindings)
	if err != nil {
		return nil, err
	}

	var b []byte
	switch {
	case w.KVBindings == nil && w.WASMBindings != nil:
		b = wasm
	case w.KVBindings != nil && w.WASMBindings == nil:
		b = kv
	case w.KVBindings != nil && w.WASMBindings != nil:
		// Merge the two json-ified arrays into one
		b = bytes.Replace(append(kv, wasm...), []byte("]["), []byte(","), 1)
	default:
		b = []byte("[]")
	}

	return json.Marshal(struct {
		BodyPart string          `json:"body_part"`
		Bindings json.RawMessage `json:"bindings"`
	}{"script", b})
}

// WorkerKVBinding defines the binding between a variable name and a KV namespace
//
// API reference: https://developers.cloudflare.com/workers/api/resource-bindings/kv-namespaces/
type WorkerKVBinding struct {
	Name        string `json:"name"`
	NamespaceID string `json:"namespace_id"`
}

// MarshalJSON implements Marshaller interface
// Adds a type field when marshalling into json
func (w *WorkerKVBinding) MarshalJSON() ([]byte, error) {
	type Copy WorkerKVBinding
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Copy
	}{
		Type: "kv_namespace",
		Copy: (*Copy)(w),
	})
}

// WorkerWASMBinding binds WASM modules to worker scripts
//
// API reference: https://developers.cloudflare.com/workers/api/resource-bindings/webassembly-modules/
type WorkerWASMBinding struct {
	Name   string    `json:"name"`
	Module io.Reader `json:"-"`
}

// MarshalJSON implements Marshaller interface
// Adds a type field when marshalling into json
func (w *WorkerWASMBinding) MarshalJSON() ([]byte, error) {
	type Copy WorkerWASMBinding
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Copy
	}{
		Type: "wasm_module",
		Copy: (*Copy)(w),
	})
}

// DeleteWorker deletes worker for a zone.
//
// API reference: https://api.cloudflare.com/#worker-script-delete-worker
func (api *API) DeleteWorker(requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	// if ScriptName is provided we will treat as org request
	if requestParams.ScriptName != "" {
		return api.deleteWorkerWithName(requestParams.ScriptName)
	}
	uri := "/zones/" + requestParams.ZoneID + "/workers/script"
	res, err := api.makeRequest("DELETE", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DeleteWorkerWithName deletes worker for a zone.
// This is an enterprise only feature https://developers.cloudflare.com/workers/api/config-api-for-enterprise
// account must be specified as api option https://godoc.org/github.com/cloudflare/cloudflare-go#UsingAccount
//
// API reference: https://api.cloudflare.com/#worker-script-delete-worker
func (api *API) deleteWorkerWithName(scriptName string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required for enterprise only request")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	res, err := api.makeRequest("DELETE", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DownloadWorker fetch raw script content for your worker returns []byte containing worker code js
//
// API reference: https://api.cloudflare.com/#worker-script-download-worker
func (api *API) DownloadWorker(requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	if requestParams.ScriptName != "" {
		return api.downloadWorkerWithName(requestParams.ScriptName)
	}
	uri := "/zones/" + requestParams.ZoneID + "/workers/script"
	res, err := api.makeRequest("GET", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	r.Script = string(res)
	r.Success = true
	return r, nil
}

// DownloadWorkerWithName fetch raw script content for your worker returns string containing worker code js
// This is an enterprise only feature https://developers.cloudflare.com/workers/api/config-api-for-enterprise/
//
// API reference: https://api.cloudflare.com/#worker-script-download-worker
func (api *API) downloadWorkerWithName(scriptName string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required for enterprise only request")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	res, err := api.makeRequest("GET", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	r.Script = string(res)
	r.Success = true
	return r, nil
}

// ListWorkerScripts returns list of worker scripts for given account.
//
// This is an enterprise only feature https://developers.cloudflare.com/workers/api/config-api-for-enterprise
//
// API reference: https://developers.cloudflare.com/workers/api/config-api-for-enterprise/
func (api *API) ListWorkerScripts() (WorkerListResponse, error) {
	if api.AccountID == "" {
		return WorkerListResponse{}, errors.New("account ID required for enterprise only request")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return WorkerListResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerListResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// UploadWorker push raw script content for your worker.
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorker(requestParams *WorkerRequestParams, data string) (WorkerScriptResponse, error) {
	if requestParams.ScriptName != "" {
		return api.uploadWorkerWithName(requestParams.ScriptName, data)
	}
	uri := "/zones/" + requestParams.ZoneID + "/workers/script"
	headers := make(http.Header)
	headers.Set("Content-Type", "application/javascript")
	res, err := api.makeRequestWithHeaders("PUT", uri, []byte(data), headers)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// UploadWorkerWithName push raw script content for your worker.
//
// This is an enterprise only feature https://developers.cloudflare.com/workers/api/config-api-for-enterprise/
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) uploadWorkerWithName(scriptName string, data string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required for enterprise only request")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	headers := make(http.Header)
	headers.Set("Content-Type", "application/javascript")
	res, err := api.makeRequestWithHeaders("PUT", uri, []byte(data), headers)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// UploadWorkerWithBindings push raw script content for your worker with resource bindings
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker and https://developers.cloudflare.com/workers/api/resource-bindings/
func (api *API) UploadWorkerWithBindings(requestParams *WorkerRequestParams, resourceBindings *WorkerResourceBindings, data string) (WorkerScriptResponse, error) {
	var r WorkerScriptResponse
	formData, contentType, err := api.createBindingFormData(resourceBindings, data)
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	if requestParams.ScriptName != "" {
		return api.uploadWorkerWithBindingsAndName(requestParams.ScriptName, formData, contentType)
	}
	uri := "/zones/" + requestParams.ZoneID + "/workers/script"
	headers := make(http.Header)
	headers.Set("Content-Type", contentType)
	res, err := api.makeRequestWithHeaders("PUT", uri, formData, headers)
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// uploadWorkerWithBindingsAndName push raw script content for your worker with resource bindings
// This is an enterprise only feature https://developers.cloudflare.com/workers/api/config-api-for-enterprise/
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker and https://developers.cloudflare.com/workers/api/resource-bindings/
func (api *API) uploadWorkerWithBindingsAndName(scriptName string, formData *bytes.Buffer, contentType string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required for enterprise only request")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	headers := make(http.Header)
	headers.Set("Content-Type", contentType)
	var r WorkerScriptResponse
	res, err := api.makeRequestWithHeaders("PUT", uri, formData, headers)
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

func (api *API) createBindingFormData(resourceBindings *WorkerResourceBindings, data string) (*bytes.Buffer, string, error) {
	var b bytes.Buffer

	bindJSON, err := json.Marshal(resourceBindings)
	if err != nil {
		return &b, "", fmt.Errorf("error marshalling bindings to json: %s", err)
	}
	parts := map[string]io.Reader{
		"script":   strings.NewReader(data),
		"metadata": bytes.NewReader(bindJSON),
	}
	// Add WASM bytecode
	for i, wasm := range resourceBindings.WASMBindings {
		key := fmt.Sprintf("wasm_resource%d", i)
		if _, ok := parts[key]; ok {
			return &b, "", fmt.Errorf("attempting to add duplicate form key '%s'", err)
		}
		parts[key] = wasm.Module
	}
	form := multipart.NewWriter(&b)
	for key, reader := range parts {
		var fw io.Writer
		if fw, err = form.CreateFormField(key); err != nil {
			return &b, "", fmt.Errorf("error creating form field for bindings api: %s", err)
		}
		if _, err := io.Copy(fw, reader); err != nil {
			return &b, "", fmt.Errorf("error copying form-data for bindings api: %s", err)
		}
	}
	if err := form.Close(); err != nil {
		return &b, "", fmt.Errorf("error closing form-data stream: %s", err)
	}
	return &b, form.FormDataContentType(), nil
}

// CreateWorkerRoute creates worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-filters-create-filter
func (api *API) CreateWorkerRoute(zoneID string, route WorkerRoute) (WorkerRouteResponse, error) {
	// Check whether a script name is defined in order to determine whether
	// to use the single-script or multi-script endpoint.
	pathComponent := "filters"
	if route.Script != "" {
		if api.AccountID == "" {
			return WorkerRouteResponse{}, errors.New("account ID required for enterprise only request")
		}
		pathComponent = "routes"
	}

	uri := "/zones/" + zoneID + "/workers/" + pathComponent
	res, err := api.makeRequest("POST", uri, route)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DeleteWorkerRoute deletes worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-filters-delete-filter
func (api *API) DeleteWorkerRoute(zoneID string, routeID string) (WorkerRouteResponse, error) {
	// For deleting a route, it doesn't matter whether we use the
	// single-script or multi-script endpoint
	uri := "/zones/" + zoneID + "/workers/filters/" + routeID
	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// ListWorkerRoutes returns list of worker routes
//
// API reference: https://api.cloudflare.com/#worker-filters-list-filters
func (api *API) ListWorkerRoutes(zoneID string) (WorkerRoutesResponse, error) {
	pathComponent := "filters"
	if api.AccountID != "" {
		pathComponent = "routes"
	}
	uri := "/zones/" + zoneID + "/workers/" + pathComponent
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return WorkerRoutesResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRoutesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRoutesResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	for i := range r.Routes {
		route := &r.Routes[i]
		// The Enabled flag will not be set in the multi-script API response
		// so we manually set it to true if the script name is not empty
		// in case any multi-script customers rely on the Enabled field
		if route.Script != "" {
			route.Enabled = true
		}
	}
	return r, nil
}

// UpdateWorkerRoute updates worker route for a zone.
//
// API reference: https://api.cloudflare.com/#worker-filters-update-filter
func (api *API) UpdateWorkerRoute(zoneID string, routeID string, route WorkerRoute) (WorkerRouteResponse, error) {
	// Check whether a script name is defined in order to determine whether
	// to use the single-script or multi-script endpoint.
	pathComponent := "filters"
	if route.Script != "" {
		if api.AccountID == "" {
			return WorkerRouteResponse{}, errors.New("account ID required for enterprise only request")
		}
		pathComponent = "routes"
	}
	uri := "/zones/" + zoneID + "/workers/" + pathComponent + "/" + routeID
	res, err := api.makeRequest("PUT", uri, route)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}
