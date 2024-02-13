// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// APIGatewayOperationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAPIGatewayOperationService]
// method instead.
type APIGatewayOperationService struct {
	Options []option.RequestOption
}

// NewAPIGatewayOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIGatewayOperationService(opts ...option.RequestOption) (r *APIGatewayOperationService) {
	r = &APIGatewayOperationService{}
	r.Options = opts
	return
}

// Update the `state` on a discovered operation
func (r *APIGatewayOperationService) Update(ctx context.Context, zoneID string, operationID string, body APIGatewayOperationUpdateParams, opts ...option.RequestOption) (res *APIGatewayOperationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewayOperationUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations/%s", zoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve the most up to date view of discovered operations
func (r *APIGatewayOperationService) List(ctx context.Context, zoneID string, query APIGatewayOperationListParams, opts ...option.RequestOption) (res *[]APIGatewayOperationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewayOperationListResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an operation
func (r *APIGatewayOperationService) Delete(ctx context.Context, zoneID string, operationID string, opts ...option.RequestOption) (res *APIGatewayOperationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIShieldAPIResponseSingle
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s", zoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Add one or more operations to a zone. Endpoints can contain path variables.
// Host, method, endpoint will be normalized to a canoncial form when creating an
// operation and must be unique on the zone. Inserting an operation that matches an
// existing one will return the record of the already existing operation and update
// its last_updated date.
func (r *APIGatewayOperationService) APIShieldEndpointManagementAddOperationsToAZone(ctx context.Context, zoneID string, body APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParams, opts ...option.RequestOption) (res *[]APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/operations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve information about all operations on a zone
func (r *APIGatewayOperationService) APIShieldEndpointManagementGetInformationAboutAllOperationsOnAZone(ctx context.Context, zoneID string, query APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParams, opts ...option.RequestOption) (res *[]APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/operations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve information about an operation
func (r *APIGatewayOperationService) Get(ctx context.Context, zoneID string, operationID string, query APIGatewayOperationGetParams, opts ...option.RequestOption) (res *APIGatewayOperationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewayOperationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s", zoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type APIGatewayOperationUpdateResponse struct {
	// State of operation in API Discovery
	//
	// - `review` - Operation is not saved into API Shield Endpoint Management
	// - `saved` - Operation is saved into API Shield Endpoint Management
	// - `ignored` - Operation is marked as ignored
	State APIGatewayOperationUpdateResponseState `json:"state"`
	JSON  apiGatewayOperationUpdateResponseJSON  `json:"-"`
}

// apiGatewayOperationUpdateResponseJSON contains the JSON metadata for the struct
// [APIGatewayOperationUpdateResponse]
type apiGatewayOperationUpdateResponseJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of operation in API Discovery
//
// - `review` - Operation is not saved into API Shield Endpoint Management
// - `saved` - Operation is saved into API Shield Endpoint Management
// - `ignored` - Operation is marked as ignored
type APIGatewayOperationUpdateResponseState string

const (
	APIGatewayOperationUpdateResponseStateReview  APIGatewayOperationUpdateResponseState = "review"
	APIGatewayOperationUpdateResponseStateSaved   APIGatewayOperationUpdateResponseState = "saved"
	APIGatewayOperationUpdateResponseStateIgnored APIGatewayOperationUpdateResponseState = "ignored"
)

type APIGatewayOperationListResponse struct {
	// UUID identifier
	ID string `json:"id,required" format:"uuid"`
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method APIGatewayOperationListResponseMethod `json:"method,required"`
	// API discovery engine(s) that discovered this operation
	Origin []APIGatewayOperationListResponseOrigin `json:"origin,required"`
	// State of operation in API Discovery
	//
	// - `review` - Operation is not saved into API Shield Endpoint Management
	// - `saved` - Operation is saved into API Shield Endpoint Management
	// - `ignored` - Operation is marked as ignored
	State    APIGatewayOperationListResponseState    `json:"state,required"`
	Features APIGatewayOperationListResponseFeatures `json:"features"`
	JSON     apiGatewayOperationListResponseJSON     `json:"-"`
}

// apiGatewayOperationListResponseJSON contains the JSON metadata for the struct
// [APIGatewayOperationListResponse]
type apiGatewayOperationListResponseJSON struct {
	ID          apijson.Field
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	Origin      apijson.Field
	State       apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method used to access the endpoint.
type APIGatewayOperationListResponseMethod string

const (
	APIGatewayOperationListResponseMethodGet     APIGatewayOperationListResponseMethod = "GET"
	APIGatewayOperationListResponseMethodPost    APIGatewayOperationListResponseMethod = "POST"
	APIGatewayOperationListResponseMethodHead    APIGatewayOperationListResponseMethod = "HEAD"
	APIGatewayOperationListResponseMethodOptions APIGatewayOperationListResponseMethod = "OPTIONS"
	APIGatewayOperationListResponseMethodPut     APIGatewayOperationListResponseMethod = "PUT"
	APIGatewayOperationListResponseMethodDelete  APIGatewayOperationListResponseMethod = "DELETE"
	APIGatewayOperationListResponseMethodConnect APIGatewayOperationListResponseMethod = "CONNECT"
	APIGatewayOperationListResponseMethodPatch   APIGatewayOperationListResponseMethod = "PATCH"
	APIGatewayOperationListResponseMethodTrace   APIGatewayOperationListResponseMethod = "TRACE"
)

//   - `ML` - Discovered operation was sourced using ML API Discovery \*
//     `SessionIdentifier` - Discovered operation was sourced using Session
//     Identifier API Discovery
type APIGatewayOperationListResponseOrigin string

const (
	APIGatewayOperationListResponseOriginMl                APIGatewayOperationListResponseOrigin = "ML"
	APIGatewayOperationListResponseOriginSessionIdentifier APIGatewayOperationListResponseOrigin = "SessionIdentifier"
)

// State of operation in API Discovery
//
// - `review` - Operation is not saved into API Shield Endpoint Management
// - `saved` - Operation is saved into API Shield Endpoint Management
// - `ignored` - Operation is marked as ignored
type APIGatewayOperationListResponseState string

const (
	APIGatewayOperationListResponseStateReview  APIGatewayOperationListResponseState = "review"
	APIGatewayOperationListResponseStateSaved   APIGatewayOperationListResponseState = "saved"
	APIGatewayOperationListResponseStateIgnored APIGatewayOperationListResponseState = "ignored"
)

type APIGatewayOperationListResponseFeatures struct {
	TrafficStats APIGatewayOperationListResponseFeaturesTrafficStats `json:"traffic_stats"`
	JSON         apiGatewayOperationListResponseFeaturesJSON         `json:"-"`
}

// apiGatewayOperationListResponseFeaturesJSON contains the JSON metadata for the
// struct [APIGatewayOperationListResponseFeatures]
type apiGatewayOperationListResponseFeaturesJSON struct {
	TrafficStats apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *APIGatewayOperationListResponseFeatures) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayOperationListResponseFeaturesTrafficStats struct {
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The period in seconds these statistics were computed over
	PeriodSeconds int64 `json:"period_seconds,required"`
	// The average number of requests seen during this period
	Requests float64                                                 `json:"requests,required"`
	JSON     apiGatewayOperationListResponseFeaturesTrafficStatsJSON `json:"-"`
}

// apiGatewayOperationListResponseFeaturesTrafficStatsJSON contains the JSON
// metadata for the struct [APIGatewayOperationListResponseFeaturesTrafficStats]
type apiGatewayOperationListResponseFeaturesTrafficStatsJSON struct {
	LastUpdated   apijson.Field
	PeriodSeconds apijson.Field
	Requests      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIGatewayOperationListResponseFeaturesTrafficStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [APIGatewayOperationDeleteResponseUnknown] or
// [shared.UnionString].
type APIGatewayOperationDeleteResponse interface {
	ImplementsAPIGatewayOperationDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIGatewayOperationDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponse struct {
	Features interface{}                                                                    `json:"features"`
	JSON     apiGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseJSON `json:"-"`
}

// apiGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponse]
type apiGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseJSON struct {
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethod `json:"method,required"`
	// UUID identifier
	OperationID string                                                                                                `json:"operation_id,required" format:"uuid"`
	Features    APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeatures `json:"features"`
	JSON        apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseJSON     `json:"-"`
}

// apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponse]
type apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method used to access the endpoint.
type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethod string

const (
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethodGet     APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethod = "GET"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethodPost    APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethod = "POST"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethodHead    APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethod = "HEAD"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethodOptions APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethod = "OPTIONS"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethodPut     APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethod = "PUT"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethodDelete  APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethod = "DELETE"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethodConnect APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethod = "CONNECT"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethodPatch   APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethod = "PATCH"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethodTrace   APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseMethod = "TRACE"
)

// Union satisfied by
// [APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholds]
// or
// [APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemas].
type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeatures interface {
	implementsAPIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeatures()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeatures)(nil)).Elem(), "")
}

type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholds struct {
	Thresholds APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds `json:"thresholds"`
	JSON       apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholdsJSON       `json:"-"`
}

// apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholdsJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholds]
type apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholdsJSON struct {
	Thresholds  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholds) implementsAPIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeatures() {
}

type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds struct {
	// The total number of auth-ids seen across this calculation.
	AuthIDTokens int64 `json:"auth_id_tokens"`
	// The number of data points used for the threshold suggestion calculation.
	DataPoints  int64     `json:"data_points"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The p50 quantile of requests (in period_seconds).
	P50 int64 `json:"p50"`
	// The p90 quantile of requests (in period_seconds).
	P90 int64 `json:"p90"`
	// The p99 quantile of requests (in period_seconds).
	P99 int64 `json:"p99"`
	// The period over which this threshold is suggested.
	PeriodSeconds int64 `json:"period_seconds"`
	// The estimated number of requests covered by these calculations.
	Requests int64 `json:"requests"`
	// The suggested threshold in requests done by the same auth_id or period_seconds.
	SuggestedThreshold int64                                                                                                                                                  `json:"suggested_threshold"`
	JSON               apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON `json:"-"`
}

// apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds]
type apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON struct {
	AuthIDTokens       apijson.Field
	DataPoints         apijson.Field
	LastUpdated        apijson.Field
	P50                apijson.Field
	P90                apijson.Field
	P99                apijson.Field
	PeriodSeconds      apijson.Field
	Requests           apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemas struct {
	ParameterSchemas APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas `json:"parameter_schemas,required"`
	JSON             apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON             `json:"-"`
}

// apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemas]
type apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON struct {
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemas) implementsAPIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeatures() {
}

type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An operation schema object containing a response.
	ParameterSchemas APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas `json:"parameter_schemas"`
	JSON             apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON             `json:"-"`
}

// apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas]
type apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON struct {
	LastUpdated      apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An operation schema object containing a response.
type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas struct {
	// An array containing the learned parameter schemas.
	Parameters []interface{} `json:"parameters"`
	// An empty response object. This field is required to yield a valid operation
	// schema.
	Responses interface{}                                                                                                                                                                        `json:"responses,nullable"`
	JSON      apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON `json:"-"`
}

// apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas]
type apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON struct {
	Parameters  apijson.Field
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayOperationGetResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method APIGatewayOperationGetResponseMethod `json:"method,required"`
	// UUID identifier
	OperationID string                                 `json:"operation_id,required" format:"uuid"`
	Features    APIGatewayOperationGetResponseFeatures `json:"features"`
	JSON        apiGatewayOperationGetResponseJSON     `json:"-"`
}

// apiGatewayOperationGetResponseJSON contains the JSON metadata for the struct
// [APIGatewayOperationGetResponse]
type apiGatewayOperationGetResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method used to access the endpoint.
type APIGatewayOperationGetResponseMethod string

const (
	APIGatewayOperationGetResponseMethodGet     APIGatewayOperationGetResponseMethod = "GET"
	APIGatewayOperationGetResponseMethodPost    APIGatewayOperationGetResponseMethod = "POST"
	APIGatewayOperationGetResponseMethodHead    APIGatewayOperationGetResponseMethod = "HEAD"
	APIGatewayOperationGetResponseMethodOptions APIGatewayOperationGetResponseMethod = "OPTIONS"
	APIGatewayOperationGetResponseMethodPut     APIGatewayOperationGetResponseMethod = "PUT"
	APIGatewayOperationGetResponseMethodDelete  APIGatewayOperationGetResponseMethod = "DELETE"
	APIGatewayOperationGetResponseMethodConnect APIGatewayOperationGetResponseMethod = "CONNECT"
	APIGatewayOperationGetResponseMethodPatch   APIGatewayOperationGetResponseMethod = "PATCH"
	APIGatewayOperationGetResponseMethodTrace   APIGatewayOperationGetResponseMethod = "TRACE"
)

// Union satisfied by
// [APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholds] or
// [APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas].
type APIGatewayOperationGetResponseFeatures interface {
	implementsAPIGatewayOperationGetResponseFeatures()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*APIGatewayOperationGetResponseFeatures)(nil)).Elem(), "")
}

type APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholds struct {
	Thresholds APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds `json:"thresholds"`
	JSON       apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsJSON       `json:"-"`
}

// apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholds]
type apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsJSON struct {
	Thresholds  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholds) implementsAPIGatewayOperationGetResponseFeatures() {
}

type APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds struct {
	// The total number of auth-ids seen across this calculation.
	AuthIDTokens int64 `json:"auth_id_tokens"`
	// The number of data points used for the threshold suggestion calculation.
	DataPoints  int64     `json:"data_points"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The p50 quantile of requests (in period_seconds).
	P50 int64 `json:"p50"`
	// The p90 quantile of requests (in period_seconds).
	P90 int64 `json:"p90"`
	// The p99 quantile of requests (in period_seconds).
	P99 int64 `json:"p99"`
	// The period over which this threshold is suggested.
	PeriodSeconds int64 `json:"period_seconds"`
	// The estimated number of requests covered by these calculations.
	Requests int64 `json:"requests"`
	// The suggested threshold in requests done by the same auth_id or period_seconds.
	SuggestedThreshold int64                                                                                   `json:"suggested_threshold"`
	JSON               apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON `json:"-"`
}

// apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds]
type apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON struct {
	AuthIDTokens       apijson.Field
	DataPoints         apijson.Field
	LastUpdated        apijson.Field
	P50                apijson.Field
	P90                apijson.Field
	P99                apijson.Field
	PeriodSeconds      apijson.Field
	Requests           apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas struct {
	ParameterSchemas APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas `json:"parameter_schemas,required"`
	JSON             apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON             `json:"-"`
}

// apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas]
type apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON struct {
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas) implementsAPIGatewayOperationGetResponseFeatures() {
}

type APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An operation schema object containing a response.
	ParameterSchemas APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas `json:"parameter_schemas"`
	JSON             apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON             `json:"-"`
}

// apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas]
type apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON struct {
	LastUpdated      apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An operation schema object containing a response.
type APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas struct {
	// An array containing the learned parameter schemas.
	Parameters []interface{} `json:"parameters"`
	// An empty response object. This field is required to yield a valid operation
	// schema.
	Responses interface{}                                                                                                         `json:"responses,nullable"`
	JSON      apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON `json:"-"`
}

// apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas]
type apiGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON struct {
	Parameters  apijson.Field
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayOperationUpdateParams struct {
	// Mark state of operation in API Discovery
	//
	// - `review` - Mark operation as for review
	// - `ignored` - Mark operation as ignored
	State param.Field[APIGatewayOperationUpdateParamsState] `json:"state"`
}

func (r APIGatewayOperationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark state of operation in API Discovery
//
// - `review` - Mark operation as for review
// - `ignored` - Mark operation as ignored
type APIGatewayOperationUpdateParamsState string

const (
	APIGatewayOperationUpdateParamsStateReview  APIGatewayOperationUpdateParamsState = "review"
	APIGatewayOperationUpdateParamsStateIgnored APIGatewayOperationUpdateParamsState = "ignored"
)

type APIGatewayOperationUpdateResponseEnvelope struct {
	Errors   APIShieldMessages                 `json:"errors,required"`
	Messages APIShieldMessages                 `json:"messages,required"`
	Result   APIGatewayOperationUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                          `json:"success,required"`
	JSON    apiGatewayOperationUpdateResponseEnvelopeJSON `json:"-"`
}

// apiGatewayOperationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [APIGatewayOperationUpdateResponseEnvelope]
type apiGatewayOperationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayOperationListParams struct {
	// When `true`, only return API Discovery results that are not saved into API
	// Shield Endpoint Management
	Diff param.Field[bool] `query:"diff"`
	// Direction to order results.
	Direction param.Field[APIGatewayOperationListParamsDirection] `query:"direction"`
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Field to order by
	Order param.Field[APIGatewayOperationListParamsOrder] `query:"order"`
	// Filter results to only include discovery results sourced from a particular
	// discovery engine
	//
	//   - `ML` - Discovered operations that were sourced using ML API Discovery
	//   - `SessionIdentifier` - Discovered operations that were sourced using Session
	//     Identifier API Discovery
	Origin param.Field[APIGatewayOperationListParamsOrigin] `query:"origin"`
	// Page number of paginated results.
	Page param.Field[interface{}] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[interface{}] `query:"per_page"`
	// Filter results to only include discovery results in a particular state. States
	// are as follows
	//
	//   - `review` - Discovered operations that are not saved into API Shield Endpoint
	//     Management
	//   - `saved` - Discovered operations that are already saved into API Shield
	//     Endpoint Management
	//   - `ignored` - Discovered operations that have been marked as ignored
	State param.Field[APIGatewayOperationListParamsState] `query:"state"`
}

// URLQuery serializes [APIGatewayOperationListParams]'s query parameters as
// `url.Values`.
func (r APIGatewayOperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type APIGatewayOperationListParamsDirection string

const (
	APIGatewayOperationListParamsDirectionAsc  APIGatewayOperationListParamsDirection = "asc"
	APIGatewayOperationListParamsDirectionDesc APIGatewayOperationListParamsDirection = "desc"
)

// Field to order by
type APIGatewayOperationListParamsOrder string

const (
	APIGatewayOperationListParamsOrderHost                    APIGatewayOperationListParamsOrder = "host"
	APIGatewayOperationListParamsOrderMethod                  APIGatewayOperationListParamsOrder = "method"
	APIGatewayOperationListParamsOrderEndpoint                APIGatewayOperationListParamsOrder = "endpoint"
	APIGatewayOperationListParamsOrderTrafficStatsRequests    APIGatewayOperationListParamsOrder = "traffic_stats.requests"
	APIGatewayOperationListParamsOrderTrafficStatsLastUpdated APIGatewayOperationListParamsOrder = "traffic_stats.last_updated"
)

// Filter results to only include discovery results sourced from a particular
// discovery engine
//
//   - `ML` - Discovered operations that were sourced using ML API Discovery
//   - `SessionIdentifier` - Discovered operations that were sourced using Session
//     Identifier API Discovery
type APIGatewayOperationListParamsOrigin string

const (
	APIGatewayOperationListParamsOriginMl                APIGatewayOperationListParamsOrigin = "ML"
	APIGatewayOperationListParamsOriginSessionIdentifier APIGatewayOperationListParamsOrigin = "SessionIdentifier"
)

// Filter results to only include discovery results in a particular state. States
// are as follows
//
//   - `review` - Discovered operations that are not saved into API Shield Endpoint
//     Management
//   - `saved` - Discovered operations that are already saved into API Shield
//     Endpoint Management
//   - `ignored` - Discovered operations that have been marked as ignored
type APIGatewayOperationListParamsState string

const (
	APIGatewayOperationListParamsStateReview  APIGatewayOperationListParamsState = "review"
	APIGatewayOperationListParamsStateSaved   APIGatewayOperationListParamsState = "saved"
	APIGatewayOperationListParamsStateIgnored APIGatewayOperationListParamsState = "ignored"
)

type APIGatewayOperationListResponseEnvelope struct {
	Errors   APIShieldMessages                 `json:"errors,required"`
	Messages APIShieldMessages                 `json:"messages,required"`
	Result   []APIGatewayOperationListResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    bool                                              `json:"success,required"`
	ResultInfo APIGatewayOperationListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       apiGatewayOperationListResponseEnvelopeJSON       `json:"-"`
}

// apiGatewayOperationListResponseEnvelopeJSON contains the JSON metadata for the
// struct [APIGatewayOperationListResponseEnvelope]
type apiGatewayOperationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayOperationListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       apiGatewayOperationListResponseEnvelopeResultInfoJSON `json:"-"`
}

// apiGatewayOperationListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [APIGatewayOperationListResponseEnvelopeResultInfo]
type apiGatewayOperationListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIShieldAPIResponseSingle struct {
	Errors   APIShieldMessages                 `json:"errors,required"`
	Messages APIShieldMessages                 `json:"messages,required"`
	Result   APIGatewayOperationDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                           `json:"success,required"`
	JSON    apiShieldAPIResponseSingleJSON `json:"-"`
}

// apiShieldAPIResponseSingleJSON contains the JSON metadata for the struct
// [APIShieldAPIResponseSingle]
type apiShieldAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParams struct {
	Body param.Field[[]APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBody] `json:"body,required"`
}

func (r APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBody struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint param.Field[string] `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host param.Field[string] `json:"host,required" format:"hostname"`
	// The HTTP method used to access the endpoint.
	Method param.Field[APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethod] `json:"method,required"`
}

func (r APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method used to access the endpoint.
type APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethod string

const (
	APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethodGet     APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethod = "GET"
	APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethodPost    APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethod = "POST"
	APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethodHead    APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethod = "HEAD"
	APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethodOptions APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethod = "OPTIONS"
	APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethodPut     APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethod = "PUT"
	APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethodDelete  APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethod = "DELETE"
	APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethodConnect APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethod = "CONNECT"
	APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethodPatch   APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethod = "PATCH"
	APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethodTrace   APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethod = "TRACE"
)

type APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelope struct {
	Errors   APIShieldMessages                                                            `json:"errors,required"`
	Messages APIShieldMessages                                                            `json:"messages,required"`
	Result   []APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    bool                                                                                         `json:"success,required"`
	ResultInfo APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelopeResultInfo `json:"result_info"`
	JSON       apiGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelopeJSON       `json:"-"`
}

// apiGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelope]
type apiGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                          `json:"total_count"`
	JSON       apiGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelopeResultInfoJSON `json:"-"`
}

// apiGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelopeResultInfo]
type apiGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParams struct {
	// Direction to order results.
	Direction param.Field[APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsDirection] `query:"direction"`
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsFeature] `query:"feature"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Field to order by. When requesting a feature, the feature keys are available for
	// ordering as well, e.g., `thresholds.suggested_threshold`.
	Order param.Field[APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[interface{}] `query:"page"`
	// Number of results to return per page
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParams]'s
// query parameters as `url.Values`.
func (r APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsDirection string

const (
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsDirectionAsc  APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsDirection = "asc"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsDirectionDesc APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsDirection = "desc"
)

type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsFeature string

const (
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsFeatureThresholds       APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsFeature = "thresholds"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsFeatureParameterSchemas APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsFeature = "parameter_schemas"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsFeatureSchemaInfo       APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsFeature = "schema_info"
)

// Field to order by. When requesting a feature, the feature keys are available for
// ordering as well, e.g., `thresholds.suggested_threshold`.
type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsOrder string

const (
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsOrderMethod        APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsOrder = "method"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsOrderHost          APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsOrder = "host"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsOrderEndpoint      APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsOrder = "endpoint"
	APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsOrderThresholdsKey APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsOrder = "thresholds.$key"
)

type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelope struct {
	Errors   APIShieldMessages                                                                               `json:"errors,required"`
	Messages APIShieldMessages                                                                               `json:"messages,required"`
	Result   []APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    bool                                                                                                            `json:"success,required"`
	ResultInfo APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelopeResultInfo `json:"result_info"`
	JSON       apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelopeJSON       `json:"-"`
}

// apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelope]
type apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                             `json:"total_count"`
	JSON       apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelopeResultInfoJSON `json:"-"`
}

// apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelopeResultInfo]
type apiGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayOperationGetParams struct {
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]APIGatewayOperationGetParamsFeature] `query:"feature"`
}

// URLQuery serializes [APIGatewayOperationGetParams]'s query parameters as
// `url.Values`.
func (r APIGatewayOperationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIGatewayOperationGetParamsFeature string

const (
	APIGatewayOperationGetParamsFeatureThresholds       APIGatewayOperationGetParamsFeature = "thresholds"
	APIGatewayOperationGetParamsFeatureParameterSchemas APIGatewayOperationGetParamsFeature = "parameter_schemas"
	APIGatewayOperationGetParamsFeatureSchemaInfo       APIGatewayOperationGetParamsFeature = "schema_info"
)

type APIGatewayOperationGetResponseEnvelope struct {
	Errors   APIShieldMessages              `json:"errors,required"`
	Messages APIShieldMessages              `json:"messages,required"`
	Result   APIGatewayOperationGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                       `json:"success,required"`
	JSON    apiGatewayOperationGetResponseEnvelopeJSON `json:"-"`
}

// apiGatewayOperationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [APIGatewayOperationGetResponseEnvelope]
type apiGatewayOperationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayOperationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
