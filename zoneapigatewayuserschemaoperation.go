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

// ZoneAPIGatewayUserSchemaOperationService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneAPIGatewayUserSchemaOperationService] method instead.
type ZoneAPIGatewayUserSchemaOperationService struct {
	Options []option.RequestOption
}

// NewZoneAPIGatewayUserSchemaOperationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAPIGatewayUserSchemaOperationService(opts ...option.RequestOption) (r *ZoneAPIGatewayUserSchemaOperationService) {
	r = &ZoneAPIGatewayUserSchemaOperationService{}
	r.Options = opts
	return
}

// Retrieves all operations from the schema. Operations that already exist in API
// Shield Endpoint Management will be returned as full operations.
func (r *ZoneAPIGatewayUserSchemaOperationService) List(ctx context.Context, zoneID string, schemaID string, query ZoneAPIGatewayUserSchemaOperationListParams, opts ...option.RequestOption) (res *ZoneAPIGatewayUserSchemaOperationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s/operations", zoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneAPIGatewayUserSchemaOperationListResponse struct {
	Errors     []ZoneAPIGatewayUserSchemaOperationListResponseError    `json:"errors"`
	Messages   []ZoneAPIGatewayUserSchemaOperationListResponseMessage  `json:"messages"`
	Result     []ZoneAPIGatewayUserSchemaOperationListResponseResult   `json:"result"`
	ResultInfo ZoneAPIGatewayUserSchemaOperationListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success bool                                              `json:"success"`
	JSON    zoneAPIGatewayUserSchemaOperationListResponseJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaOperationListResponseJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayUserSchemaOperationListResponse]
type zoneAPIGatewayUserSchemaOperationListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaOperationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaOperationListResponseError struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneAPIGatewayUserSchemaOperationListResponseErrorJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaOperationListResponseErrorJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayUserSchemaOperationListResponseError]
type zoneAPIGatewayUserSchemaOperationListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaOperationListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaOperationListResponseMessage struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneAPIGatewayUserSchemaOperationListResponseMessageJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaOperationListResponseMessageJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayUserSchemaOperationListResponseMessage]
type zoneAPIGatewayUserSchemaOperationListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaOperationListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperation] or
// [ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperation].
type ZoneAPIGatewayUserSchemaOperationListResponseResult interface {
	implementsZoneAPIGatewayUserSchemaOperationListResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAPIGatewayUserSchemaOperationListResponseResult)(nil)).Elem(), "")
}

type ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperation struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string                                                                       `json:"endpoint" format:"uri-template"`
	Features ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeatures `json:"features"`
	// RFC3986-compliant host.
	Host        string    `json:"host" format:"hostname"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethod `json:"method"`
	// UUID identifier
	OperationID string                                                                   `json:"operation_id" format:"uuid"`
	JSON        zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperation]
type zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationJSON struct {
	Endpoint    apijson.Field
	Features    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperation) implementsZoneAPIGatewayUserSchemaOperationListResponseResult() {
}

// Union satisfied by
// [ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholds]
// or
// [ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemas].
type ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeatures interface {
	implementsZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeatures()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeatures)(nil)).Elem(), "")
}

type ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholds struct {
	Thresholds ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholdsThresholds `json:"thresholds"`
	JSON       zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholdsJSON       `json:"-"`
}

// zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholdsJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholds]
type zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholdsJSON struct {
	Thresholds  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholds) implementsZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeatures() {
}

type ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholdsThresholds struct {
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
	SuggestedThreshold int64                                                                                                                        `json:"suggested_threshold"`
	JSON               zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholdsThresholdsJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholdsThresholdsJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholdsThresholds]
type zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholdsThresholdsJSON struct {
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

func (r *ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureThresholdsThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemas struct {
	ParameterSchemas ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemas `json:"parameter_schemas,required"`
	JSON             zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasJSON             `json:"-"`
}

// zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemas]
type zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasJSON struct {
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemas) implementsZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeatures() {
}

type ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemas struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An operation schema object containing a response.
	ParameterSchemas ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemasParameterSchemas `json:"parameter_schemas"`
	JSON             zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemasJSON             `json:"-"`
}

// zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemas]
type zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemasJSON struct {
	LastUpdated      apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An operation schema object containing a response.
type ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemasParameterSchemas struct {
	// An array containing the learned parameter schemas.
	Parameters []interface{} `json:"parameters"`
	// An empty response object. This field is required to yield a valid operation
	// schema.
	Responses interface{}                                                                                                                                              `json:"responses,nullable"`
	JSON      zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemasParameterSchemas]
type zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON struct {
	Parameters  apijson.Field
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationFeaturesAvFlf4mpOperationFeatureParameterSchemasParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method used to access the endpoint.
type ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethod string

const (
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethodGet     ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethod = "GET"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethodPost    ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethod = "POST"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethodHead    ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethod = "HEAD"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethodOptions ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethod = "OPTIONS"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethodPut     ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethod = "PUT"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethodDelete  ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethod = "DELETE"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethodConnect ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethod = "CONNECT"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethodPatch   ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethod = "PATCH"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethodTrace   ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpOperationMethod = "TRACE"
)

type ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperation struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host string `json:"host,required" format:"hostname"`
	// The HTTP method used to access the endpoint.
	Method ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethod `json:"method,required"`
	JSON   zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationJSON   `json:"-"`
}

// zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperation]
type zoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperation) implementsZoneAPIGatewayUserSchemaOperationListResponseResult() {
}

// The HTTP method used to access the endpoint.
type ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethod string

const (
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethodGet     ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethod = "GET"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethodPost    ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethod = "POST"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethodHead    ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethod = "HEAD"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethodOptions ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethod = "OPTIONS"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethodPut     ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethod = "PUT"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethodDelete  ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethod = "DELETE"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethodConnect ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethod = "CONNECT"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethodPatch   ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethod = "PATCH"
	ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethodTrace   ZoneAPIGatewayUserSchemaOperationListResponseResultAvFlf4mpBasicOperationMethod = "TRACE"
)

type ZoneAPIGatewayUserSchemaOperationListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       zoneAPIGatewayUserSchemaOperationListResponseResultInfoJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaOperationListResponseResultInfoJSON contains the JSON
// metadata for the struct
// [ZoneAPIGatewayUserSchemaOperationListResponseResultInfo]
type zoneAPIGatewayUserSchemaOperationListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaOperationListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaOperationListParams struct {
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]ZoneAPIGatewayUserSchemaOperationListParamsFeature] `query:"feature"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Filter results by whether operations exist in API Shield Endpoint Management or
	// not. `new` will just return operations from the schema that do not exist in API
	// Shield Endpoint Management. `existing` will just return operations from the
	// schema that already exist in API Shield Endpoint Management.
	OperationStatus param.Field[ZoneAPIGatewayUserSchemaOperationListParamsOperationStatus] `query:"operation_status"`
	// Page number of paginated results.
	Page param.Field[interface{}] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[interface{}] `query:"per_page"`
}

// URLQuery serializes [ZoneAPIGatewayUserSchemaOperationListParams]'s query
// parameters as `url.Values`.
func (r ZoneAPIGatewayUserSchemaOperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneAPIGatewayUserSchemaOperationListParamsFeature string

const (
	ZoneAPIGatewayUserSchemaOperationListParamsFeatureThresholds       ZoneAPIGatewayUserSchemaOperationListParamsFeature = "thresholds"
	ZoneAPIGatewayUserSchemaOperationListParamsFeatureParameterSchemas ZoneAPIGatewayUserSchemaOperationListParamsFeature = "parameter_schemas"
	ZoneAPIGatewayUserSchemaOperationListParamsFeatureSchemaInfo       ZoneAPIGatewayUserSchemaOperationListParamsFeature = "schema_info"
)

// Filter results by whether operations exist in API Shield Endpoint Management or
// not. `new` will just return operations from the schema that do not exist in API
// Shield Endpoint Management. `existing` will just return operations from the
// schema that already exist in API Shield Endpoint Management.
type ZoneAPIGatewayUserSchemaOperationListParamsOperationStatus string

const (
	ZoneAPIGatewayUserSchemaOperationListParamsOperationStatusNew      ZoneAPIGatewayUserSchemaOperationListParamsOperationStatus = "new"
	ZoneAPIGatewayUserSchemaOperationListParamsOperationStatusExisting ZoneAPIGatewayUserSchemaOperationListParamsOperationStatus = "existing"
)
