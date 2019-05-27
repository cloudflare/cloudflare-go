package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
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

// WorkerResource represents a resource that can bind to a WorkerScript.
//
// API reference: https://api.cloudflare.com/#worker-binding-properties
type WorkerResource struct {
	Type          string `json:"type"`
	Name          string `json:"name"`
	KVNamespaceID string `json:"namespace_id,omitempty"`
	WasmPart      string `json:"part,omitempty"`
}

// WorkerResourceMetaData provides parameters to bind WorkerResources to a WorkerScript.
type WorkerResourceMetaData struct {
	BodyPart string            `json:"body_part"`
	Bindings []*WorkerResource `json:"bindings"`
}

// WorkerResourceListResponse is a response when listing WorkerResources of a WorkerScript.
type WorkerResourceListResponse struct {
	Result []WorkerResource `json:"result"`
}

// WorkerUploadRequest provides parameters to upload a WorkerScript and WorkerBinding.
//
// Script and MetaData are provided for convience,
// they are converted into a WorkerUploadPart.
//
// If you want to upload additional form-data in the request,
// use a WorkerUploadPart (eg. use this for uploading Wasm).
type WorkerUploadRequest struct {
	Script   string
	MetaData *WorkerResourceMetaData
	Parts    []*WorkerUploadPart
}

// WorkerUploadPart provides parameters for custom data to be bundled with a WorkerScript upload.
type WorkerUploadPart struct {
	Name    string
	File    string
	Content []byte
	Type    string
}

// DeleteWorker deletes worker for a zone.
//
// If your account has multiple scripts, you must specify an OrganizationID:
// https://godoc.org/github.com/cloudflare/cloudflare-go#UsingOrganization
//
// API reference: https://api.cloudflare.com/#worker-script-delete-worker
func (api *API) DeleteWorker(requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	uri, err := api.uriOfWorkerScript(requestParams)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	res, err := api.makeRequest("DELETE", uri, nil)
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
// If your account has multiple scripts, you must specify an OrganizationID:
// https://godoc.org/github.com/cloudflare/cloudflare-go#UsingOrganization
//
// API reference: https://api.cloudflare.com/#worker-script-download-worker
func (api *API) DownloadWorker(requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	uri, err := api.uriOfWorkerScript(requestParams)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	r.Script = string(res)
	r.Success = true
	return r, nil
}

// ListWorkerScripts returns list of worker scripts for given organization
//
// This is an enterprise only feature https://developers.cloudflare.com/workers/api/config-api-for-enterprise
//
// API reference: https://developers.cloudflare.com/workers/api/config-api-for-enterprise/
func (api *API) ListWorkerScripts() (WorkerListResponse, error) {
	if api.OrganizationID == "" {
		return WorkerListResponse{}, errors.New("organization ID required for enterprise only request")
	}
	uri := "/accounts/" + api.OrganizationID + "/workers/scripts"
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

// UploadWorker push raw script content for your worker
//
// If your account has multiple scripts, you must specify an OrganizationID:
// https://godoc.org/github.com/cloudflare/cloudflare-go#UsingOrganization
//
// If you want to upload metadata like WorkerBinding, use UploadWorkerBundle.
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorker(requestParams *WorkerRequestParams, data string) (WorkerScriptResponse, error) {
	uri, err := api.uriOfWorkerScript(requestParams)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	headers := make(http.Header)
	headers.Set("Content-Type", "application/javascript")
	res, err := api.makeRequestWithHeaders("PUT", uri, []byte(data), headers)
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// UploadWorkerMultiPart uploads a WorkerScript with additional WorkerUploadParts.
func (api *API) UploadWorkerMultiPart(requestParams *WorkerRequestParams, upload *WorkerUploadRequest) (WorkerScriptResponse, error) {
	uri, err := api.uriOfWorkerScript(requestParams)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}

	// Conveience logic in this code block.
	//
	// If specified, convert Script and MetaData fields into a WorkerUploadPart.
	// Otherwise, trust the client has added their own custom parts.
	parts := upload.Parts
	if upload.MetaData != nil {
		if upload.MetaData.BodyPart == "" {
			upload.MetaData.BodyPart = "script"
		}
		meta, err := json.Marshal(upload.MetaData)
		if err != nil {
			return r, errors.Wrap(err, errMakeRequestError)
		}
		parts = append(parts, &WorkerUploadPart{
			Name:    "metadata",
			File:    "metadata.json",
			Content: meta,
			Type:    "application/json"})
	}
	if upload.Script != "" {
		parts = append(parts, &WorkerUploadPart{
			Name:    "script",
			File:    "script.js",
			Content: []byte(upload.Script),
			Type:    "application/javascript"})
	}

	// Encode each of the WorkerUploadParts into a multi part file.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for i := 0; i < len(parts); i++ {
		part := parts[i]

		// Create the MIME header with part name, file name, and content type.
		header := make(textproto.MIMEHeader)
		escaper := strings.NewReplacer("\\", "\\\\", `"`, "\\\"")
		header.Set("Content-Disposition",
			fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
				escaper.Replace(part.Name), escaper.Replace(part.File)))
		header.Set("Content-Type", part.Type)

		// Commit the part to the write buffer and copy over the content.
		form, err := writer.CreatePart(header)
		if err != nil {
			return r, errors.Wrap(err, errMakeRequestError)
		}
		io.Copy(form, bytes.NewReader(part.Content))
	}

	// Close the writer which finalizes the body boundry.
	err = writer.Close()
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	headers := make(http.Header)
	headers.Set("Content-Type", writer.FormDataContentType())

	// Make a single put request with all of the multi part files.
	res, err := api.makeRequestWithHeaders(http.MethodPut, uri, body.Bytes(), headers)
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// CreateWorkerRoute creates worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-filters-create-filter
func (api *API) CreateWorkerRoute(zoneID string, route WorkerRoute) (WorkerRouteResponse, error) {
	// Check whether a script name is defined in order to determine whether
	// to use the single-script or multi-script endpoint.
	uri, err := api.uriOfWorkerRoute(&WorkerRequestParams{ZoneID: zoneID, ScriptName: route.Script})
	var r WorkerRouteResponse
	if err != nil {
		return r, err
	}
	res, err := api.makeRequest("POST", uri, route)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}
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
	if api.OrganizationID != "" {
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
	uri, err := api.uriOfWorkerRoute(&WorkerRequestParams{ZoneID: zoneID, ScriptName: route.Script})
	var r WorkerRouteResponse
	if err != nil {
		return r, err
	}
	res, err := api.makeRequest("PUT", uri+"/"+routeID, route)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// ListWorkerBindings gets the list of WorkerResources bound to the WorkerScript.
//
// If your account has multiple scripts, you must specify an OrganizationID:
// https://godoc.org/github.com/cloudflare/cloudflare-go#UsingOrganization
//
// API reference: https://api.cloudflare.com/#worker-binding-list-bindings
func (api *API) ListWorkerBindings(requestParams *WorkerRequestParams) (WorkerResourceListResponse, error) {
	uri, err := api.uriOfWorkerScript(requestParams)
	var r WorkerResourceListResponse
	if err != nil {
		return r, err
	}
	headers := make(http.Header)
	headers.Set("Content-Type", "application/javascript")
	res, err := api.makeRequestWithHeaders("GET", uri+"/bindings", nil, headers)
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// uriOfWorkerHelper is a helper that determines the API endpoints.
func (api *API) uriOfWorkerHelper(requestParams *WorkerRequestParams, singletonScript, multiScript string) (string, error) {
	if requestParams.ScriptName == "" {
		if requestParams.ZoneID == "" {
			return "", errors.New("zone or organization ID required for request")
		}
		return singletonScript, nil
	}
	if api.OrganizationID == "" {
		return "", errors.New("organization ID required for enterprise only request")
	}
	return multiScript, nil
}

// uriOfWorkerScript is a helper that determines the API endpoint for WorkerScript.
func (api *API) uriOfWorkerScript(requestParams *WorkerRequestParams) (string, error) {
	return api.uriOfWorkerHelper(requestParams,
		"/zones/"+requestParams.ZoneID+"/workers/script",
		"/accounts/"+api.OrganizationID+"/workers/scripts/"+requestParams.ScriptName)
}

// uriOfWorkerRoute is a helper that determines the API endpoint for WorkerRoute.
func (api *API) uriOfWorkerRoute(requestParams *WorkerRequestParams) (string, error) {
	return api.uriOfWorkerHelper(requestParams,
		"/zones/"+requestParams.ZoneID+"/workers/filters",
		"/zones/"+api.OrganizationID+"/workers/routes")
}
