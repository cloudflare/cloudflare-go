// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// OperationService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationService] method instead.
type OperationService struct {
	Options          []option.RequestOption
	SchemaValidation *OperationSchemaValidationService
}

// NewOperationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOperationService(opts ...option.RequestOption) (r *OperationService) {
	r = &OperationService{}
	r.Options = opts
	r.SchemaValidation = NewOperationSchemaValidationService(opts...)
	return
}

// Add one or more operations to a zone. Endpoints can contain path variables.
// Host, method, endpoint will be normalized to a canoncial form when creating an
// operation and must be unique on the zone. Inserting an operation that matches an
// existing one will return the record of the already existing operation and update
// its last_updated date.
func (r *OperationService) New(ctx context.Context, params OperationNewParams, opts ...option.RequestOption) (res *[]APIShield, err error) {
	opts = append(r.Options[:], opts...)
	var env OperationNewResponseEnvelope
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve information about all operations on a zone
func (r *OperationService) List(ctx context.Context, params OperationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[APIShield], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/operations", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve information about all operations on a zone
func (r *OperationService) ListAutoPaging(ctx context.Context, params OperationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[APIShield] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Delete an operation
func (r *OperationService) Delete(ctx context.Context, operationID string, body OperationDeleteParams, opts ...option.RequestOption) (res *OperationDeleteResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env OperationDeleteResponseEnvelope
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s", body.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve information about an operation
func (r *OperationService) Get(ctx context.Context, operationID string, params OperationGetParams, opts ...option.RequestOption) (res *APIShield, err error) {
	opts = append(r.Options[:], opts...)
	var env OperationGetResponseEnvelope
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s", params.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type APIShield struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method APIShieldMethod `json:"method,required"`
	// UUID identifier
	OperationID string            `json:"operation_id,required" format:"uuid"`
	Features    APIShieldFeatures `json:"features"`
	JSON        apiShieldJSON     `json:"-"`
}

// apiShieldJSON contains the JSON metadata for the struct [APIShield]
type apiShieldJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShield) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldJSON) RawJSON() string {
	return r.raw
}

func (r APIShield) implementsAPIGatewayUserSchemaOperationListResponse() {}

// The HTTP method used to access the endpoint.
type APIShieldMethod string

const (
	APIShieldMethodGet     APIShieldMethod = "GET"
	APIShieldMethodPost    APIShieldMethod = "POST"
	APIShieldMethodHead    APIShieldMethod = "HEAD"
	APIShieldMethodOptions APIShieldMethod = "OPTIONS"
	APIShieldMethodPut     APIShieldMethod = "PUT"
	APIShieldMethodDelete  APIShieldMethod = "DELETE"
	APIShieldMethodConnect APIShieldMethod = "CONNECT"
	APIShieldMethodPatch   APIShieldMethod = "PATCH"
	APIShieldMethodTrace   APIShieldMethod = "TRACE"
)

func (r APIShieldMethod) IsKnown() bool {
	switch r {
	case APIShieldMethodGet, APIShieldMethodPost, APIShieldMethodHead, APIShieldMethodOptions, APIShieldMethodPut, APIShieldMethodDelete, APIShieldMethodConnect, APIShieldMethodPatch, APIShieldMethodTrace:
		return true
	}
	return false
}

type APIShieldFeatures struct {
	// This field can have the runtime type of
	// [APIShieldFeaturesAPIShieldOperationFeatureThresholdsThresholds].
	Thresholds interface{} `json:"thresholds,required"`
	// This field can have the runtime type of
	// [APIShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas].
	ParameterSchemas interface{}           `json:"parameter_schemas,required"`
	JSON             apiShieldFeaturesJSON `json:"-"`
	union            APIShieldFeaturesUnion
}

// apiShieldFeaturesJSON contains the JSON metadata for the struct
// [APIShieldFeatures]
type apiShieldFeaturesJSON struct {
	Thresholds       apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r apiShieldFeaturesJSON) RawJSON() string {
	return r.raw
}

func (r *APIShieldFeatures) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [APIShieldFeaturesUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [api_gateway.APIShieldFeaturesAPIShieldOperationFeatureThresholds],
// [api_gateway.APIShieldFeaturesAPIShieldOperationFeatureParameterSchemas].
func (r APIShieldFeatures) AsUnion() APIShieldFeaturesUnion {
	return r.union
}

// Union satisfied by
// [api_gateway.APIShieldFeaturesAPIShieldOperationFeatureThresholds] or
// [api_gateway.APIShieldFeaturesAPIShieldOperationFeatureParameterSchemas].
type APIShieldFeaturesUnion interface {
	implementsAPIGatewayAPIShieldFeatures()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIShieldFeaturesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIShieldFeaturesAPIShieldOperationFeatureThresholds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIShieldFeaturesAPIShieldOperationFeatureParameterSchemas{}),
		},
	)
}

type APIShieldFeaturesAPIShieldOperationFeatureThresholds struct {
	Thresholds APIShieldFeaturesAPIShieldOperationFeatureThresholdsThresholds `json:"thresholds"`
	JSON       apiShieldFeaturesAPIShieldOperationFeatureThresholdsJSON       `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureThresholdsJSON contains the JSON
// metadata for the struct [APIShieldFeaturesAPIShieldOperationFeatureThresholds]
type apiShieldFeaturesAPIShieldOperationFeatureThresholdsJSON struct {
	Thresholds  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureThresholdsJSON) RawJSON() string {
	return r.raw
}

func (r APIShieldFeaturesAPIShieldOperationFeatureThresholds) implementsAPIGatewayAPIShieldFeatures() {
}

type APIShieldFeaturesAPIShieldOperationFeatureThresholdsThresholds struct {
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
	SuggestedThreshold int64                                                              `json:"suggested_threshold"`
	JSON               apiShieldFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON contains the
// JSON metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureThresholdsThresholds]
type apiShieldFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON struct {
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

func (r *APIShieldFeaturesAPIShieldOperationFeatureThresholdsThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON) RawJSON() string {
	return r.raw
}

type APIShieldFeaturesAPIShieldOperationFeatureParameterSchemas struct {
	ParameterSchemas APIShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas `json:"parameter_schemas,required"`
	JSON             apiShieldFeaturesAPIShieldOperationFeatureParameterSchemasJSON             `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureParameterSchemasJSON contains the JSON
// metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureParameterSchemas]
type apiShieldFeaturesAPIShieldOperationFeatureParameterSchemasJSON struct {
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureParameterSchemasJSON) RawJSON() string {
	return r.raw
}

func (r APIShieldFeaturesAPIShieldOperationFeatureParameterSchemas) implementsAPIGatewayAPIShieldFeatures() {
}

type APIShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An operation schema object containing a response.
	ParameterSchemas APIShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas `json:"parameter_schemas"`
	JSON             apiShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON             `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas]
type apiShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON struct {
	LastUpdated      apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

// An operation schema object containing a response.
type APIShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas struct {
	// An array containing the learned parameter schemas.
	Parameters []interface{} `json:"parameters"`
	// An empty response object. This field is required to yield a valid operation
	// schema.
	Responses interface{}                                                                                    `json:"responses,nullable"`
	JSON      apiShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas]
type apiShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON struct {
	Parameters  apijson.Field
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [api_gateway.OperationDeleteResponseUnknown] or
// [shared.UnionString].
type OperationDeleteResponseUnion interface {
	ImplementsAPIGatewayOperationDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OperationDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type OperationNewParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   []OperationNewParamsBody `json:"body,required"`
}

func (r OperationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type OperationNewParamsBody struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint param.Field[string] `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host param.Field[string] `json:"host,required" format:"hostname"`
	// The HTTP method used to access the endpoint.
	Method param.Field[OperationNewParamsBodyMethod] `json:"method,required"`
}

func (r OperationNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method used to access the endpoint.
type OperationNewParamsBodyMethod string

const (
	OperationNewParamsBodyMethodGet     OperationNewParamsBodyMethod = "GET"
	OperationNewParamsBodyMethodPost    OperationNewParamsBodyMethod = "POST"
	OperationNewParamsBodyMethodHead    OperationNewParamsBodyMethod = "HEAD"
	OperationNewParamsBodyMethodOptions OperationNewParamsBodyMethod = "OPTIONS"
	OperationNewParamsBodyMethodPut     OperationNewParamsBodyMethod = "PUT"
	OperationNewParamsBodyMethodDelete  OperationNewParamsBodyMethod = "DELETE"
	OperationNewParamsBodyMethodConnect OperationNewParamsBodyMethod = "CONNECT"
	OperationNewParamsBodyMethodPatch   OperationNewParamsBodyMethod = "PATCH"
	OperationNewParamsBodyMethodTrace   OperationNewParamsBodyMethod = "TRACE"
)

func (r OperationNewParamsBodyMethod) IsKnown() bool {
	switch r {
	case OperationNewParamsBodyMethodGet, OperationNewParamsBodyMethodPost, OperationNewParamsBodyMethodHead, OperationNewParamsBodyMethodOptions, OperationNewParamsBodyMethodPut, OperationNewParamsBodyMethodDelete, OperationNewParamsBodyMethodConnect, OperationNewParamsBodyMethodPatch, OperationNewParamsBodyMethodTrace:
		return true
	}
	return false
}

type OperationNewResponseEnvelope struct {
	Errors   Message     `json:"errors,required"`
	Messages Message     `json:"messages,required"`
	Result   []APIShield `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    bool                                   `json:"success,required"`
	ResultInfo OperationNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       operationNewResponseEnvelopeJSON       `json:"-"`
}

// operationNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OperationNewResponseEnvelope]
type operationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OperationNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       operationNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// operationNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [OperationNewResponseEnvelopeResultInfo]
type operationNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type OperationListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Direction to order results.
	Direction param.Field[OperationListParamsDirection] `query:"direction"`
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]OperationListParamsFeature] `query:"feature"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Field to order by. When requesting a feature, the feature keys are available for
	// ordering as well, e.g., `thresholds.suggested_threshold`.
	Order param.Field[OperationListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[interface{}] `query:"page"`
	// Number of results to return per page
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [OperationListParams]'s query parameters as `url.Values`.
func (r OperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type OperationListParamsDirection string

const (
	OperationListParamsDirectionAsc  OperationListParamsDirection = "asc"
	OperationListParamsDirectionDesc OperationListParamsDirection = "desc"
)

func (r OperationListParamsDirection) IsKnown() bool {
	switch r {
	case OperationListParamsDirectionAsc, OperationListParamsDirectionDesc:
		return true
	}
	return false
}

type OperationListParamsFeature string

const (
	OperationListParamsFeatureThresholds       OperationListParamsFeature = "thresholds"
	OperationListParamsFeatureParameterSchemas OperationListParamsFeature = "parameter_schemas"
	OperationListParamsFeatureSchemaInfo       OperationListParamsFeature = "schema_info"
)

func (r OperationListParamsFeature) IsKnown() bool {
	switch r {
	case OperationListParamsFeatureThresholds, OperationListParamsFeatureParameterSchemas, OperationListParamsFeatureSchemaInfo:
		return true
	}
	return false
}

// Field to order by. When requesting a feature, the feature keys are available for
// ordering as well, e.g., `thresholds.suggested_threshold`.
type OperationListParamsOrder string

const (
	OperationListParamsOrderMethod        OperationListParamsOrder = "method"
	OperationListParamsOrderHost          OperationListParamsOrder = "host"
	OperationListParamsOrderEndpoint      OperationListParamsOrder = "endpoint"
	OperationListParamsOrderThresholdsKey OperationListParamsOrder = "thresholds.$key"
)

func (r OperationListParamsOrder) IsKnown() bool {
	switch r {
	case OperationListParamsOrderMethod, OperationListParamsOrderHost, OperationListParamsOrderEndpoint, OperationListParamsOrderThresholdsKey:
		return true
	}
	return false
}

type OperationDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OperationDeleteResponseEnvelope struct {
	Errors   Message                      `json:"errors,required"`
	Messages Message                      `json:"messages,required"`
	Result   OperationDeleteResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success bool                                `json:"success,required"`
	JSON    operationDeleteResponseEnvelopeJSON `json:"-"`
}

// operationDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [OperationDeleteResponseEnvelope]
type operationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OperationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]OperationGetParamsFeature] `query:"feature"`
}

// URLQuery serializes [OperationGetParams]'s query parameters as `url.Values`.
func (r OperationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type OperationGetParamsFeature string

const (
	OperationGetParamsFeatureThresholds       OperationGetParamsFeature = "thresholds"
	OperationGetParamsFeatureParameterSchemas OperationGetParamsFeature = "parameter_schemas"
	OperationGetParamsFeatureSchemaInfo       OperationGetParamsFeature = "schema_info"
)

func (r OperationGetParamsFeature) IsKnown() bool {
	switch r {
	case OperationGetParamsFeatureThresholds, OperationGetParamsFeatureParameterSchemas, OperationGetParamsFeatureSchemaInfo:
		return true
	}
	return false
}

type OperationGetResponseEnvelope struct {
	Errors   Message   `json:"errors,required"`
	Messages Message   `json:"messages,required"`
	Result   APIShield `json:"result,required"`
	// Whether the API call was successful
	Success bool                             `json:"success,required"`
	JSON    operationGetResponseEnvelopeJSON `json:"-"`
}

// operationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [OperationGetResponseEnvelope]
type operationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
