package cloudflare

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"time"

	"github.com/pkg/errors"
)

// WorkerRequestParams provides parameters for worker requests for both enterprise and standard requests
type WorkerRequestParams struct {
	ZoneID     string
	ScriptName string
}

// WorkerScriptParams provides a worker script and the associated bindings
type WorkerScriptParams struct {
	Script string

	// Bindings should be a map where the keys are the binding name, and the
	// values are the binding content
	Bindings map[string]WorkerBinding
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

// Workers supports multiple types of bindings, e.g. KV namespaces or WebAssembly modules, and each type
// of binding will be represented differently in the upload request body. At a high-level, every binding
// will specify metadata, which is a JSON object with the properties "name" and "type". Some types of bindings
// will also have additional metadata properties. For example, KV bindings also specify the KV namespace.
// In addition to the metadata, some binding types may need to include additional data as part of the
// multipart form. For example, WebAssembly bindings will include the contents of the WebAssembly module.

// WorkerBinding is the generic interface implemented by all of
// the various binding types
type WorkerBinding interface {
	// serialize is responsible for returning the binding metadata as well as an optionally
	// returning a function that can modify the multipart form body. For example, this is used
	// by WebAssembly bindings to add a new part containing the WebAssembly module contents.
	serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error)
}

// workerBindingMeta is the metadata portion of the binding
type workerBindingMeta = map[string]interface{}

// workerBindingBodyWriter allows for a binding to add additional parts to the multipart body
type workerBindingBodyWriter func(*multipart.Writer) error

// WorkerInheritBinding will just persist whatever binding content was previously uploaded
type WorkerInheritBinding struct {
	// Optional parameter that allows for renaming a binding without changing
	// its contents. If `OldName` is empty, the binding name will not be changed.
	OldName string
}

func (b WorkerInheritBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	meta := workerBindingMeta{
		"name": bindingName,
		"type": "inherit",
	}

	if b.OldName != "" {
		meta["old_name"] = b.OldName
	}

	return meta, nil, nil
}

// WorkerKvNamespaceBinding is a binding to a Workers KV Namespace
//
// https://developers.cloudflare.com/workers/archive/api/resource-bindings/kv-namespaces/
type WorkerKvNamespaceBinding struct {
	NamespaceID string
}

func (b WorkerKvNamespaceBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.NamespaceID == "" {
		return nil, nil, errors.Errorf(`NamespaceID for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name":         bindingName,
		"type":         "kv_namespace",
		"namespace_id": b.NamespaceID,
	}, nil, nil
}

// WorkerWebAssemblyBinding is a binding to a WebAssembly module
//
// https://developers.cloudflare.com/workers/archive/api/resource-bindings/webassembly-modules/
type WorkerWebAssemblyBinding struct {
	Module io.Reader
}

func (b WorkerWebAssemblyBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	partName := getRandomPartName()

	bodyWriter := func(mpw *multipart.Writer) error {
		var hdr = textproto.MIMEHeader{}
		hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, partName))
		hdr.Set("content-type", "application/wasm")
		pw, err := mpw.CreatePart(hdr)
		if err != nil {
			return err
		}
		_, err = io.Copy(pw, b.Module)
		return err
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": "wasm_module",
		"part": partName,
	}, bodyWriter, nil
}

// Each binding that adds a part to the multipart form body will need
// a unique part name so we just generate a random 128bit hex string
func getRandomPartName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes)
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
		return api.uploadWorkerWithName(requestParams.ScriptName, "application/javascript", []byte(data))
	}
	return api.uploadWorkerForZone(requestParams.ZoneID, "application/javascript", []byte(data))
}

// UploadWorkerWithBindings push raw script content and bindings for your worker
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorkerWithBindings(requestParams *WorkerRequestParams, data *WorkerScriptParams) (WorkerScriptResponse, error) {
	contentType, body, err := formatMultipartBody(data)
	if err != nil {
		return WorkerScriptResponse{}, err
	}
	if requestParams.ScriptName != "" {
		return api.uploadWorkerWithName(requestParams.ScriptName, contentType, body)
	}
	return api.uploadWorkerForZone(requestParams.ZoneID, contentType, body)
}

func (api *API) uploadWorkerForZone(zoneID, contentType string, body []byte) (WorkerScriptResponse, error) {
	uri := "/zones/" + zoneID + "/workers/script"
	headers := make(http.Header)
	headers.Set("Content-Type", contentType)
	res, err := api.makeRequestWithHeaders("PUT", uri, body, headers)
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

func (api *API) uploadWorkerWithName(scriptName, contentType string, body []byte) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required for enterprise only request")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	headers := make(http.Header)
	headers.Set("Content-Type", contentType)
	res, err := api.makeRequestWithHeaders("PUT", uri, body, headers)
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

// Returns content-type, body, error
func formatMultipartBody(params *WorkerScriptParams) (string, []byte, error) {
	var buf = &bytes.Buffer{}
	var mpw = multipart.NewWriter(buf)
	defer mpw.Close()

	// Write metadata part
	scriptPartName := "script"
	meta := struct {
		BodyPart string              `json:"body_part"`
		Bindings []workerBindingMeta `json:"bindings"`
	}{
		BodyPart: scriptPartName,
		Bindings: make([]workerBindingMeta, 0, len(params.Bindings)),
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
	hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, scriptPartName))
	hdr.Set("content-type", "application/javascript")
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

// CreateWorkerRoute creates worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-filters-create-filter, https://api.cloudflare.com/#worker-routes-create-route
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
// API reference: https://api.cloudflare.com/#worker-routes-delete-route
func (api *API) DeleteWorkerRoute(zoneID string, routeID string) (WorkerRouteResponse, error) {
	uri := "/zones/" + zoneID + "/workers/routes/" + routeID
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
// API reference: https://api.cloudflare.com/#worker-filters-list-filters, https://api.cloudflare.com/#worker-routes-list-routes
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
// API reference: https://api.cloudflare.com/#worker-filters-update-filter, https://api.cloudflare.com/#worker-routes-update-route
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
