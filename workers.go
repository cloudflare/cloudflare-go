package cloudflare

import (
	"encoding/json"
	"github.com/pkg/errors"
	"time"
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
	Script  string `json:"script"`
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

// DeleteWorker deletes worker for a zone.
//
// API reference: https://api.cloudflare.com/#worker-script-delete-worker
func (api *API) DeleteWorker(requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	// if organization ID is defined and ScriptName is provided we will treat as org request
	if api.organizationID != "" && requestParams.ScriptName != "" {
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
// organizationID must be specified as api option https://godoc.org/github.com/cloudflare/cloudflare-go#UsingOrganization
//
// API reference: https://api.cloudflare.com/#worker-script-delete-worker
func (api *API) deleteWorkerWithName(scriptName string) (WorkerScriptResponse, error) {
	if api.organizationID == "" {
		return WorkerScriptResponse{}, errors.New("organization ID required for enterprise only request")
	}
	uri := "/accounts/" + api.organizationID + "/workers/scripts/" + scriptName
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
	if api.organizationID != "" && requestParams.ScriptName != "" {
		return api.downloadWorkerWithName(requestParams.ScriptName)
	}
	uri := "/zones/" + requestParams.ZoneID + "/workers/script"
	api.headers.Add("Content-Type", "application/javascript")
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
	if api.organizationID == "" {
		return WorkerScriptResponse{}, errors.New("organization ID required for enterprise only request")
	}
	uri := "/accounts/" + api.organizationID + "/workers/scripts/" + scriptName
	api.headers.Add("Content-Type", "application/javascript")
	res, err := api.makeRequest("GET", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	r.Script = string(res)
	r.Success = true
	return r, nil
}

// ListWorkerScripts returns list of worker scripts for given organization
// This is an enterprise only feature https://developers.cloudflare.com/workers/api/config-api-for-enterprise
//
// API reference: https://developers.cloudflare.com/workers/api/config-api-for-enterprise/
func (api *API) ListWorkerScripts() (WorkerListResponse, error) {
	if api.organizationID == "" {
		return WorkerListResponse{}, errors.New("organization ID required for enterprise only request")
	}
	uri := "/accounts/" + api.organizationID + "/workers/scripts"
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
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorker(requestParams *WorkerRequestParams, data string) (WorkerScriptResponse, error) {
	if api.organizationID != "" && requestParams.ScriptName != "" {
		return api.uploadWorkerWithName(requestParams.ScriptName, data)
	}
	uri := "/zones/" + requestParams.ZoneID + "/workers/script"
	api.headers.Add("Content-Type", "application/javascript")
	res, err := api.makeRequest("PUT", uri, []byte(data))
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

// UploadWorkerWithName push raw script content for your worker
// This is an enterprise only feature https://developers.cloudflare.com/workers/api/config-api-for-enterprise/
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) uploadWorkerWithName(scriptName string, data string) (WorkerScriptResponse, error) {
	if api.organizationID == "" {
		return WorkerScriptResponse{}, errors.New("organization ID required for enterprise only request")
	}
	uri := "/accounts/" + api.organizationID + "/workers/scripts/" + scriptName
	api.headers.Add("Content-Type", "application/javascript")
	res, err := api.makeRequest("PUT", uri, []byte(data))
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

// CreateWorkerRoute creates worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-filters-create-filter
func (api *API) CreateWorkerRoute(zoneID string, route WorkerRoute) (WorkerRouteResponse, error) {
	uri := "/zones/" + zoneID + "/workers/filters"
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
	uri := "/zones/" + zoneID + "/workers/filters"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return WorkerRoutesResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRoutesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRoutesResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// UpdateWorkerRoute updates worker route for a zone.
//
// API reference: https://api.cloudflare.com/#worker-filters-update-filter
func (api *API) UpdateWorkerRoute(zoneID string, routeID string, route WorkerRoute) (WorkerRouteResponse, error) {
	uri := "/zones/" + zoneID + "/workers/filters/" + routeID
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
