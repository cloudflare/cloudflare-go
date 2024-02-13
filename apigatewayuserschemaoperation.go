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
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// APIGatewayUserSchemaOperationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAPIGatewayUserSchemaOperationService] method instead.
type APIGatewayUserSchemaOperationService struct {
	Options []option.RequestOption
}

// NewAPIGatewayUserSchemaOperationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAPIGatewayUserSchemaOperationService(opts ...option.RequestOption) (r *APIGatewayUserSchemaOperationService) {
	r = &APIGatewayUserSchemaOperationService{}
	r.Options = opts
	return
}

// Retrieves all operations from the schema. Operations that already exist in API
// Shield Endpoint Management will be returned as full operations.
func (r *APIGatewayUserSchemaOperationService) List(ctx context.Context, zoneID string, schemaID string, query APIGatewayUserSchemaOperationListParams, opts ...option.RequestOption) (res *[]APIGatewayUserSchemaOperationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewayUserSchemaOperationListResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s/operations", zoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [APIGatewayUserSchemaOperationListResponseAPIShieldOperation]
// or [APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperation].
type APIGatewayUserSchemaOperationListResponse interface {
	implementsAPIGatewayUserSchemaOperationListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*APIGatewayUserSchemaOperationListResponse)(nil)).Elem(), "")
}

type APIGatewayUserSchemaOperationListResponseAPIShieldOperation struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethod `json:"method,required"`
	// UUID identifier
	OperationID string                                                              `json:"operation_id,required" format:"uuid"`
	Features    APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeatures `json:"features"`
	JSON        apiGatewayUserSchemaOperationListResponseAPIShieldOperationJSON     `json:"-"`
}

// apiGatewayUserSchemaOperationListResponseAPIShieldOperationJSON contains the
// JSON metadata for the struct
// [APIGatewayUserSchemaOperationListResponseAPIShieldOperation]
type apiGatewayUserSchemaOperationListResponseAPIShieldOperationJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayUserSchemaOperationListResponseAPIShieldOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r APIGatewayUserSchemaOperationListResponseAPIShieldOperation) implementsAPIGatewayUserSchemaOperationListResponse() {
}

// The HTTP method used to access the endpoint.
type APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethod string

const (
	APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethodGet     APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethod = "GET"
	APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethodPost    APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethod = "POST"
	APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethodHead    APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethod = "HEAD"
	APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethodOptions APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethod = "OPTIONS"
	APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethodPut     APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethod = "PUT"
	APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethodDelete  APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethod = "DELETE"
	APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethodConnect APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethod = "CONNECT"
	APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethodPatch   APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethod = "PATCH"
	APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethodTrace   APIGatewayUserSchemaOperationListResponseAPIShieldOperationMethod = "TRACE"
)

// Union satisfied by
// [APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholds]
// or
// [APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemas].
type APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeatures interface {
	implementsAPIGatewayUserSchemaOperationListResponseAPIShieldOperationFeatures()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeatures)(nil)).Elem(), "")
}

type APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholds struct {
	Thresholds APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholds `json:"thresholds"`
	JSON       apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsJSON       `json:"-"`
}

// apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsJSON
// contains the JSON metadata for the struct
// [APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholds]
type apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsJSON struct {
	Thresholds  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholds) implementsAPIGatewayUserSchemaOperationListResponseAPIShieldOperationFeatures() {
}

type APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholds struct {
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
	SuggestedThreshold int64                                                                                                                `json:"suggested_threshold"`
	JSON               apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON `json:"-"`
}

// apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON
// contains the JSON metadata for the struct
// [APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholds]
type apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON struct {
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

func (r *APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemas struct {
	ParameterSchemas APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas `json:"parameter_schemas,required"`
	JSON             apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasJSON             `json:"-"`
}

// apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasJSON
// contains the JSON metadata for the struct
// [APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemas]
type apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasJSON struct {
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemas) implementsAPIGatewayUserSchemaOperationListResponseAPIShieldOperationFeatures() {
}

type APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An operation schema object containing a response.
	ParameterSchemas APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas `json:"parameter_schemas"`
	JSON             apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON             `json:"-"`
}

// apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas]
type apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON struct {
	LastUpdated      apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An operation schema object containing a response.
type APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas struct {
	// An array containing the learned parameter schemas.
	Parameters []interface{} `json:"parameters"`
	// An empty response object. This field is required to yield a valid operation
	// schema.
	Responses interface{}                                                                                                                                      `json:"responses,nullable"`
	JSON      apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON `json:"-"`
}

// apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas]
type apiGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON struct {
	Parameters  apijson.Field
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayUserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperation struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host string `json:"host,required" format:"hostname"`
	// The HTTP method used to access the endpoint.
	Method APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethod `json:"method,required"`
	JSON   apiGatewayUserSchemaOperationListResponseAPIShieldBasicOperationJSON   `json:"-"`
}

// apiGatewayUserSchemaOperationListResponseAPIShieldBasicOperationJSON contains
// the JSON metadata for the struct
// [APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperation]
type apiGatewayUserSchemaOperationListResponseAPIShieldBasicOperationJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperation) implementsAPIGatewayUserSchemaOperationListResponse() {
}

// The HTTP method used to access the endpoint.
type APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethod string

const (
	APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethodGet     APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethod = "GET"
	APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethodPost    APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethod = "POST"
	APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethodHead    APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethod = "HEAD"
	APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethodOptions APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethod = "OPTIONS"
	APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethodPut     APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethod = "PUT"
	APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethodDelete  APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethod = "DELETE"
	APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethodConnect APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethod = "CONNECT"
	APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethodPatch   APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethod = "PATCH"
	APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethodTrace   APIGatewayUserSchemaOperationListResponseAPIShieldBasicOperationMethod = "TRACE"
)

type APIGatewayUserSchemaOperationListParams struct {
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]APIGatewayUserSchemaOperationListParamsFeature] `query:"feature"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Filter results by whether operations exist in API Shield Endpoint Management or
	// not. `new` will just return operations from the schema that do not exist in API
	// Shield Endpoint Management. `existing` will just return operations from the
	// schema that already exist in API Shield Endpoint Management.
	OperationStatus param.Field[APIGatewayUserSchemaOperationListParamsOperationStatus] `query:"operation_status"`
	// Page number of paginated results.
	Page param.Field[interface{}] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[interface{}] `query:"per_page"`
}

// URLQuery serializes [APIGatewayUserSchemaOperationListParams]'s query parameters
// as `url.Values`.
func (r APIGatewayUserSchemaOperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIGatewayUserSchemaOperationListParamsFeature string

const (
	APIGatewayUserSchemaOperationListParamsFeatureThresholds       APIGatewayUserSchemaOperationListParamsFeature = "thresholds"
	APIGatewayUserSchemaOperationListParamsFeatureParameterSchemas APIGatewayUserSchemaOperationListParamsFeature = "parameter_schemas"
	APIGatewayUserSchemaOperationListParamsFeatureSchemaInfo       APIGatewayUserSchemaOperationListParamsFeature = "schema_info"
)

// Filter results by whether operations exist in API Shield Endpoint Management or
// not. `new` will just return operations from the schema that do not exist in API
// Shield Endpoint Management. `existing` will just return operations from the
// schema that already exist in API Shield Endpoint Management.
type APIGatewayUserSchemaOperationListParamsOperationStatus string

const (
	APIGatewayUserSchemaOperationListParamsOperationStatusNew      APIGatewayUserSchemaOperationListParamsOperationStatus = "new"
	APIGatewayUserSchemaOperationListParamsOperationStatusExisting APIGatewayUserSchemaOperationListParamsOperationStatus = "existing"
)

type APIGatewayUserSchemaOperationListResponseEnvelope struct {
	Errors   APIShieldMessages                           `json:"errors,required"`
	Messages APIShieldMessages                           `json:"messages,required"`
	Result   []APIGatewayUserSchemaOperationListResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    bool                                                        `json:"success,required"`
	ResultInfo APIGatewayUserSchemaOperationListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       apiGatewayUserSchemaOperationListResponseEnvelopeJSON       `json:"-"`
}

// apiGatewayUserSchemaOperationListResponseEnvelopeJSON contains the JSON metadata
// for the struct [APIGatewayUserSchemaOperationListResponseEnvelope]
type apiGatewayUserSchemaOperationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayUserSchemaOperationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayUserSchemaOperationListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       apiGatewayUserSchemaOperationListResponseEnvelopeResultInfoJSON `json:"-"`
}

// apiGatewayUserSchemaOperationListResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [APIGatewayUserSchemaOperationListResponseEnvelopeResultInfo]
type apiGatewayUserSchemaOperationListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayUserSchemaOperationListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
