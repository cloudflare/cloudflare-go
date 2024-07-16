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
	var env OperationNewResponseEnvelope
	opts = append(r.Options[:], opts...)
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
func (r *OperationService) List(ctx context.Context, params OperationListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[OperationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
func (r *OperationService) ListAutoPaging(ctx context.Context, params OperationListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[OperationListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete an operation
func (r *OperationService) Delete(ctx context.Context, operationID string, body OperationDeleteParams, opts ...option.RequestOption) (res *OperationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s", body.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieve information about an operation
func (r *OperationService) Get(ctx context.Context, operationID string, params OperationGetParams, opts ...option.RequestOption) (res *OperationGetResponse, err error) {
	var env OperationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
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
	// UUID
	OperationID string            `json:"operation_id,required"`
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
	ParameterSchemas interface{} `json:"parameter_schemas,required"`
	// This field can have the runtime type of
	// [APIShieldFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting].
	APIRouting interface{} `json:"api_routing,required"`
	// This field can have the runtime type of
	// [APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals].
	ConfidenceIntervals interface{} `json:"confidence_intervals,required"`
	// This field can have the runtime type of
	// [APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo].
	SchemaInfo interface{}           `json:"schema_info,required"`
	JSON       apiShieldFeaturesJSON `json:"-"`
	union      APIShieldFeaturesUnion
}

// apiShieldFeaturesJSON contains the JSON metadata for the struct
// [APIShieldFeatures]
type apiShieldFeaturesJSON struct {
	Thresholds          apijson.Field
	ParameterSchemas    apijson.Field
	APIRouting          apijson.Field
	ConfidenceIntervals apijson.Field
	SchemaInfo          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r apiShieldFeaturesJSON) RawJSON() string {
	return r.raw
}

func (r *APIShieldFeatures) UnmarshalJSON(data []byte) (err error) {
	*r = APIShieldFeatures{}
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
// [api_gateway.APIShieldFeaturesAPIShieldOperationFeatureParameterSchemas],
// [api_gateway.APIShieldFeaturesAPIShieldOperationFeatureAPIRouting],
// [api_gateway.APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervals],
// [api_gateway.APIShieldFeaturesAPIShieldOperationFeatureSchemaInfo].
func (r APIShieldFeatures) AsUnion() APIShieldFeaturesUnion {
	return r.union
}

// Union satisfied by
// [api_gateway.APIShieldFeaturesAPIShieldOperationFeatureThresholds],
// [api_gateway.APIShieldFeaturesAPIShieldOperationFeatureParameterSchemas],
// [api_gateway.APIShieldFeaturesAPIShieldOperationFeatureAPIRouting],
// [api_gateway.APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervals] or
// [api_gateway.APIShieldFeaturesAPIShieldOperationFeatureSchemaInfo].
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
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIShieldFeaturesAPIShieldOperationFeatureAPIRouting{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervals{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIShieldFeaturesAPIShieldOperationFeatureSchemaInfo{}),
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

type APIShieldFeaturesAPIShieldOperationFeatureAPIRouting struct {
	// API Routing settings on endpoint.
	APIRouting APIShieldFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting `json:"api_routing"`
	JSON       apiShieldFeaturesAPIShieldOperationFeatureAPIRoutingJSON       `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureAPIRoutingJSON contains the JSON
// metadata for the struct [APIShieldFeaturesAPIShieldOperationFeatureAPIRouting]
type apiShieldFeaturesAPIShieldOperationFeatureAPIRoutingJSON struct {
	APIRouting  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureAPIRoutingJSON) RawJSON() string {
	return r.raw
}

func (r APIShieldFeaturesAPIShieldOperationFeatureAPIRouting) implementsAPIGatewayAPIShieldFeatures() {
}

// API Routing settings on endpoint.
type APIShieldFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// Target route.
	Route string                                                             `json:"route"`
	JSON  apiShieldFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON contains the
// JSON metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting]
type apiShieldFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON struct {
	LastUpdated apijson.Field
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON) RawJSON() string {
	return r.raw
}

type APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervals struct {
	ConfidenceIntervals APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals `json:"confidence_intervals"`
	JSON                apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON                `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON contains the
// JSON metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervals]
type apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON struct {
	ConfidenceIntervals apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

func (r APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervals) implementsAPIGatewayAPIShieldFeatures() {
}

type APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals struct {
	LastUpdated        time.Time                                                                                          `json:"last_updated" format:"date-time"`
	SuggestedThreshold APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold `json:"suggested_threshold"`
	JSON               apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON               `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals]
type apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON struct {
	LastUpdated        apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

type APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold struct {
	ConfidenceIntervals APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals `json:"confidence_intervals"`
	// Suggested threshold.
	Mean float64                                                                                                `json:"mean"`
	JSON apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON
// contains the JSON metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold]
type apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON struct {
	ConfidenceIntervals apijson.Field
	Mean                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON) RawJSON() string {
	return r.raw
}

type APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals struct {
	// Upper and lower bound for percentile estimate
	P90 APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 `json:"p90"`
	// Upper and lower bound for percentile estimate
	P95 APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 `json:"p95"`
	// Upper and lower bound for percentile estimate
	P99  APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99  `json:"p99"`
	JSON apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals]
type apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON struct {
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                      `json:"upper"`
	JSON  apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON
// contains the JSON metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90]
type apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                      `json:"upper"`
	JSON  apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON
// contains the JSON metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95]
type apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                      `json:"upper"`
	JSON  apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON
// contains the JSON metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99]
type apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON) RawJSON() string {
	return r.raw
}

type APIShieldFeaturesAPIShieldOperationFeatureSchemaInfo struct {
	SchemaInfo APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo `json:"schema_info"`
	JSON       apiShieldFeaturesAPIShieldOperationFeatureSchemaInfoJSON       `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureSchemaInfoJSON contains the JSON
// metadata for the struct [APIShieldFeaturesAPIShieldOperationFeatureSchemaInfo]
type apiShieldFeaturesAPIShieldOperationFeatureSchemaInfoJSON struct {
	SchemaInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureSchemaInfoJSON) RawJSON() string {
	return r.raw
}

func (r APIShieldFeaturesAPIShieldOperationFeatureSchemaInfo) implementsAPIGatewayAPIShieldFeatures() {
}

type APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo struct {
	// Schema active on endpoint.
	ActiveSchema APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema `json:"active_schema"`
	// True if a Cloudflare-provided learned schema is available for this endpoint.
	LearnedAvailable bool `json:"learned_available"`
	// Action taken on requests failing validation.
	MitigationAction APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction `json:"mitigation_action,nullable"`
	JSON             apiShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON             `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON contains the
// JSON metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo]
type apiShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON struct {
	ActiveSchema     apijson.Field
	LearnedAvailable apijson.Field
	MitigationAction apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON) RawJSON() string {
	return r.raw
}

// Schema active on endpoint.
type APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// True if schema is Cloudflare-provided.
	IsLearned bool `json:"is_learned"`
	// Schema file name.
	Name string                                                                         `json:"name"`
	JSON apiShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON `json:"-"`
}

// apiShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON
// contains the JSON metadata for the struct
// [APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema]
type apiShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IsLearned   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON) RawJSON() string {
	return r.raw
}

// Action taken on requests failing validation.
type APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction string

const (
	APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone  APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "none"
	APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog   APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "log"
	APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "block"
)

func (r APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction) IsKnown() bool {
	switch r {
	case APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone, APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog, APIShieldFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock:
		return true
	}
	return false
}

type OperationListResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method OperationListResponseMethod `json:"method,required"`
	// UUID
	OperationID string                        `json:"operation_id,required"`
	Features    OperationListResponseFeatures `json:"features"`
	JSON        operationListResponseJSON     `json:"-"`
}

// operationListResponseJSON contains the JSON metadata for the struct
// [OperationListResponse]
type operationListResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type OperationListResponseMethod string

const (
	OperationListResponseMethodGet     OperationListResponseMethod = "GET"
	OperationListResponseMethodPost    OperationListResponseMethod = "POST"
	OperationListResponseMethodHead    OperationListResponseMethod = "HEAD"
	OperationListResponseMethodOptions OperationListResponseMethod = "OPTIONS"
	OperationListResponseMethodPut     OperationListResponseMethod = "PUT"
	OperationListResponseMethodDelete  OperationListResponseMethod = "DELETE"
	OperationListResponseMethodConnect OperationListResponseMethod = "CONNECT"
	OperationListResponseMethodPatch   OperationListResponseMethod = "PATCH"
	OperationListResponseMethodTrace   OperationListResponseMethod = "TRACE"
)

func (r OperationListResponseMethod) IsKnown() bool {
	switch r {
	case OperationListResponseMethodGet, OperationListResponseMethodPost, OperationListResponseMethodHead, OperationListResponseMethodOptions, OperationListResponseMethodPut, OperationListResponseMethodDelete, OperationListResponseMethodConnect, OperationListResponseMethodPatch, OperationListResponseMethodTrace:
		return true
	}
	return false
}

type OperationListResponseFeatures struct {
	// This field can have the runtime type of
	// [OperationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds].
	Thresholds interface{} `json:"thresholds,required"`
	// This field can have the runtime type of
	// [OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas].
	ParameterSchemas interface{} `json:"parameter_schemas,required"`
	// This field can have the runtime type of
	// [OperationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting].
	APIRouting interface{} `json:"api_routing,required"`
	// This field can have the runtime type of
	// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals].
	ConfidenceIntervals interface{} `json:"confidence_intervals,required"`
	// This field can have the runtime type of
	// [OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo].
	SchemaInfo interface{}                       `json:"schema_info,required"`
	JSON       operationListResponseFeaturesJSON `json:"-"`
	union      OperationListResponseFeaturesUnion
}

// operationListResponseFeaturesJSON contains the JSON metadata for the struct
// [OperationListResponseFeatures]
type operationListResponseFeaturesJSON struct {
	Thresholds          apijson.Field
	ParameterSchemas    apijson.Field
	APIRouting          apijson.Field
	ConfidenceIntervals apijson.Field
	SchemaInfo          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r operationListResponseFeaturesJSON) RawJSON() string {
	return r.raw
}

func (r *OperationListResponseFeatures) UnmarshalJSON(data []byte) (err error) {
	*r = OperationListResponseFeatures{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [OperationListResponseFeaturesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [api_gateway.OperationListResponseFeaturesAPIShieldOperationFeatureThresholds],
// [api_gateway.OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas],
// [api_gateway.OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting],
// [api_gateway.OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals],
// [api_gateway.OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo].
func (r OperationListResponseFeatures) AsUnion() OperationListResponseFeaturesUnion {
	return r.union
}

// Union satisfied by
// [api_gateway.OperationListResponseFeaturesAPIShieldOperationFeatureThresholds],
// [api_gateway.OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas],
// [api_gateway.OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting],
// [api_gateway.OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals]
// or
// [api_gateway.OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo].
type OperationListResponseFeaturesUnion interface {
	implementsAPIGatewayOperationListResponseFeatures()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OperationListResponseFeaturesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationListResponseFeaturesAPIShieldOperationFeatureThresholds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo{}),
		},
	)
}

type OperationListResponseFeaturesAPIShieldOperationFeatureThresholds struct {
	Thresholds OperationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds `json:"thresholds"`
	JSON       operationListResponseFeaturesAPIShieldOperationFeatureThresholdsJSON       `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureThresholdsJSON contains
// the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureThresholds]
type operationListResponseFeaturesAPIShieldOperationFeatureThresholdsJSON struct {
	Thresholds  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureThresholdsJSON) RawJSON() string {
	return r.raw
}

func (r OperationListResponseFeaturesAPIShieldOperationFeatureThresholds) implementsAPIGatewayOperationListResponseFeatures() {
}

type OperationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds struct {
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
	SuggestedThreshold int64                                                                          `json:"suggested_threshold"`
	JSON               operationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds]
type operationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON struct {
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

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON) RawJSON() string {
	return r.raw
}

type OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas struct {
	ParameterSchemas OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas `json:"parameter_schemas,required"`
	JSON             operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON             `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas]
type operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON struct {
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON) RawJSON() string {
	return r.raw
}

func (r OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas) implementsAPIGatewayOperationListResponseFeatures() {
}

type OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An operation schema object containing a response.
	ParameterSchemas OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas `json:"parameter_schemas"`
	JSON             operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON             `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas]
type operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON struct {
	LastUpdated      apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

// An operation schema object containing a response.
type OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas struct {
	// An array containing the learned parameter schemas.
	Parameters []interface{} `json:"parameters"`
	// An empty response object. This field is required to yield a valid operation
	// schema.
	Responses interface{}                                                                                                `json:"responses,nullable"`
	JSON      operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas]
type operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON struct {
	Parameters  apijson.Field
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

type OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting struct {
	// API Routing settings on endpoint.
	APIRouting OperationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting `json:"api_routing"`
	JSON       operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON       `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON contains
// the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting]
type operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON struct {
	APIRouting  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON) RawJSON() string {
	return r.raw
}

func (r OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting) implementsAPIGatewayOperationListResponseFeatures() {
}

// API Routing settings on endpoint.
type OperationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// Target route.
	Route string                                                                         `json:"route"`
	JSON  operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting]
type operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON struct {
	LastUpdated apijson.Field
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON) RawJSON() string {
	return r.raw
}

type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals struct {
	ConfidenceIntervals OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals `json:"confidence_intervals"`
	JSON                operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON                `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON struct {
	ConfidenceIntervals apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

func (r OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals) implementsAPIGatewayOperationListResponseFeatures() {
}

type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals struct {
	LastUpdated        time.Time                                                                                                      `json:"last_updated" format:"date-time"`
	SuggestedThreshold OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold `json:"suggested_threshold"`
	JSON               operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON               `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON struct {
	LastUpdated        apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold struct {
	ConfidenceIntervals OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals `json:"confidence_intervals"`
	// Suggested threshold.
	Mean float64                                                                                                            `json:"mean"`
	JSON operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON struct {
	ConfidenceIntervals apijson.Field
	Mean                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON) RawJSON() string {
	return r.raw
}

type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals struct {
	// Upper and lower bound for percentile estimate
	P90 OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 `json:"p90"`
	// Upper and lower bound for percentile estimate
	P95 OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 `json:"p95"`
	// Upper and lower bound for percentile estimate
	P99  OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99  `json:"p99"`
	JSON operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON struct {
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                  `json:"upper"`
	JSON  operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                  `json:"upper"`
	JSON  operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                  `json:"upper"`
	JSON  operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON) RawJSON() string {
	return r.raw
}

type OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo struct {
	SchemaInfo OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo `json:"schema_info"`
	JSON       operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON       `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON contains
// the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo]
type operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON struct {
	SchemaInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON) RawJSON() string {
	return r.raw
}

func (r OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo) implementsAPIGatewayOperationListResponseFeatures() {
}

type OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo struct {
	// Schema active on endpoint.
	ActiveSchema OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema `json:"active_schema"`
	// True if a Cloudflare-provided learned schema is available for this endpoint.
	LearnedAvailable bool `json:"learned_available"`
	// Action taken on requests failing validation.
	MitigationAction OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction `json:"mitigation_action,nullable"`
	JSON             operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON             `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo]
type operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON struct {
	ActiveSchema     apijson.Field
	LearnedAvailable apijson.Field
	MitigationAction apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON) RawJSON() string {
	return r.raw
}

// Schema active on endpoint.
type OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// True if schema is Cloudflare-provided.
	IsLearned bool `json:"is_learned"`
	// Schema file name.
	Name string                                                                                     `json:"name"`
	JSON operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema]
type operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IsLearned   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON) RawJSON() string {
	return r.raw
}

// Action taken on requests failing validation.
type OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction string

const (
	OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone  OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "none"
	OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog   OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "log"
	OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "block"
)

func (r OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction) IsKnown() bool {
	switch r {
	case OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone, OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog, OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock:
		return true
	}
	return false
}

type OperationDeleteResponse struct {
	Errors   Message `json:"errors,required"`
	Messages Message `json:"messages,required"`
	// Whether the API call was successful
	Success OperationDeleteResponseSuccess `json:"success,required"`
	JSON    operationDeleteResponseJSON    `json:"-"`
}

// operationDeleteResponseJSON contains the JSON metadata for the struct
// [OperationDeleteResponse]
type operationDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OperationDeleteResponseSuccess bool

const (
	OperationDeleteResponseSuccessTrue OperationDeleteResponseSuccess = true
)

func (r OperationDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case OperationDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type OperationGetResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method OperationGetResponseMethod `json:"method,required"`
	// UUID
	OperationID string                       `json:"operation_id,required"`
	Features    OperationGetResponseFeatures `json:"features"`
	JSON        operationGetResponseJSON     `json:"-"`
}

// operationGetResponseJSON contains the JSON metadata for the struct
// [OperationGetResponse]
type operationGetResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type OperationGetResponseMethod string

const (
	OperationGetResponseMethodGet     OperationGetResponseMethod = "GET"
	OperationGetResponseMethodPost    OperationGetResponseMethod = "POST"
	OperationGetResponseMethodHead    OperationGetResponseMethod = "HEAD"
	OperationGetResponseMethodOptions OperationGetResponseMethod = "OPTIONS"
	OperationGetResponseMethodPut     OperationGetResponseMethod = "PUT"
	OperationGetResponseMethodDelete  OperationGetResponseMethod = "DELETE"
	OperationGetResponseMethodConnect OperationGetResponseMethod = "CONNECT"
	OperationGetResponseMethodPatch   OperationGetResponseMethod = "PATCH"
	OperationGetResponseMethodTrace   OperationGetResponseMethod = "TRACE"
)

func (r OperationGetResponseMethod) IsKnown() bool {
	switch r {
	case OperationGetResponseMethodGet, OperationGetResponseMethodPost, OperationGetResponseMethodHead, OperationGetResponseMethodOptions, OperationGetResponseMethodPut, OperationGetResponseMethodDelete, OperationGetResponseMethodConnect, OperationGetResponseMethodPatch, OperationGetResponseMethodTrace:
		return true
	}
	return false
}

type OperationGetResponseFeatures struct {
	// This field can have the runtime type of
	// [OperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds].
	Thresholds interface{} `json:"thresholds,required"`
	// This field can have the runtime type of
	// [OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas].
	ParameterSchemas interface{} `json:"parameter_schemas,required"`
	// This field can have the runtime type of
	// [OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting].
	APIRouting interface{} `json:"api_routing,required"`
	// This field can have the runtime type of
	// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals].
	ConfidenceIntervals interface{} `json:"confidence_intervals,required"`
	// This field can have the runtime type of
	// [OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo].
	SchemaInfo interface{}                      `json:"schema_info,required"`
	JSON       operationGetResponseFeaturesJSON `json:"-"`
	union      OperationGetResponseFeaturesUnion
}

// operationGetResponseFeaturesJSON contains the JSON metadata for the struct
// [OperationGetResponseFeatures]
type operationGetResponseFeaturesJSON struct {
	Thresholds          apijson.Field
	ParameterSchemas    apijson.Field
	APIRouting          apijson.Field
	ConfidenceIntervals apijson.Field
	SchemaInfo          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r operationGetResponseFeaturesJSON) RawJSON() string {
	return r.raw
}

func (r *OperationGetResponseFeatures) UnmarshalJSON(data []byte) (err error) {
	*r = OperationGetResponseFeatures{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [OperationGetResponseFeaturesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [api_gateway.OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds],
// [api_gateway.OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas],
// [api_gateway.OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting],
// [api_gateway.OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals],
// [api_gateway.OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo].
func (r OperationGetResponseFeatures) AsUnion() OperationGetResponseFeaturesUnion {
	return r.union
}

// Union satisfied by
// [api_gateway.OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds],
// [api_gateway.OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas],
// [api_gateway.OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting],
// [api_gateway.OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals]
// or
// [api_gateway.OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo].
type OperationGetResponseFeaturesUnion interface {
	implementsAPIGatewayOperationGetResponseFeatures()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OperationGetResponseFeaturesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo{}),
		},
	)
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds struct {
	Thresholds OperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds `json:"thresholds"`
	JSON       operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsJSON       `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsJSON contains the
// JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds]
type operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsJSON struct {
	Thresholds  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsJSON) RawJSON() string {
	return r.raw
}

func (r OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds) implementsAPIGatewayOperationGetResponseFeatures() {
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds struct {
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
	SuggestedThreshold int64                                                                         `json:"suggested_threshold"`
	JSON               operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds]
type operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON struct {
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

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON) RawJSON() string {
	return r.raw
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas struct {
	ParameterSchemas OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas `json:"parameter_schemas,required"`
	JSON             operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON             `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas]
type operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON struct {
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON) RawJSON() string {
	return r.raw
}

func (r OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas) implementsAPIGatewayOperationGetResponseFeatures() {
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An operation schema object containing a response.
	ParameterSchemas OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas `json:"parameter_schemas"`
	JSON             operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON             `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas]
type operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON struct {
	LastUpdated      apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

// An operation schema object containing a response.
type OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas struct {
	// An array containing the learned parameter schemas.
	Parameters []interface{} `json:"parameters"`
	// An empty response object. This field is required to yield a valid operation
	// schema.
	Responses interface{}                                                                                               `json:"responses,nullable"`
	JSON      operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas]
type operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON struct {
	Parameters  apijson.Field
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting struct {
	// API Routing settings on endpoint.
	APIRouting OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting `json:"api_routing"`
	JSON       operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON       `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON contains the
// JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting]
type operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON struct {
	APIRouting  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON) RawJSON() string {
	return r.raw
}

func (r OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting) implementsAPIGatewayOperationGetResponseFeatures() {
}

// API Routing settings on endpoint.
type OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// Target route.
	Route string                                                                        `json:"route"`
	JSON  operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting]
type operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON struct {
	LastUpdated apijson.Field
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON) RawJSON() string {
	return r.raw
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals struct {
	ConfidenceIntervals OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals `json:"confidence_intervals"`
	JSON                operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON                `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON struct {
	ConfidenceIntervals apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

func (r OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals) implementsAPIGatewayOperationGetResponseFeatures() {
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals struct {
	LastUpdated        time.Time                                                                                                     `json:"last_updated" format:"date-time"`
	SuggestedThreshold OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold `json:"suggested_threshold"`
	JSON               operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON               `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON struct {
	LastUpdated        apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold struct {
	ConfidenceIntervals OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals `json:"confidence_intervals"`
	// Suggested threshold.
	Mean float64                                                                                                           `json:"mean"`
	JSON operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON struct {
	ConfidenceIntervals apijson.Field
	Mean                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON) RawJSON() string {
	return r.raw
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals struct {
	// Upper and lower bound for percentile estimate
	P90 OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 `json:"p90"`
	// Upper and lower bound for percentile estimate
	P95 OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 `json:"p95"`
	// Upper and lower bound for percentile estimate
	P99  OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99  `json:"p99"`
	JSON operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON struct {
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                 `json:"upper"`
	JSON  operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                 `json:"upper"`
	JSON  operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                 `json:"upper"`
	JSON  operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON) RawJSON() string {
	return r.raw
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo struct {
	SchemaInfo OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo `json:"schema_info"`
	JSON       operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON       `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON contains the
// JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo]
type operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON struct {
	SchemaInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON) RawJSON() string {
	return r.raw
}

func (r OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo) implementsAPIGatewayOperationGetResponseFeatures() {
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo struct {
	// Schema active on endpoint.
	ActiveSchema OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema `json:"active_schema"`
	// True if a Cloudflare-provided learned schema is available for this endpoint.
	LearnedAvailable bool `json:"learned_available"`
	// Action taken on requests failing validation.
	MitigationAction OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction `json:"mitigation_action,nullable"`
	JSON             operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON             `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo]
type operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON struct {
	ActiveSchema     apijson.Field
	LearnedAvailable apijson.Field
	MitigationAction apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON) RawJSON() string {
	return r.raw
}

// Schema active on endpoint.
type OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// True if schema is Cloudflare-provided.
	IsLearned bool `json:"is_learned"`
	// Schema file name.
	Name string                                                                                    `json:"name"`
	JSON operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema]
type operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IsLearned   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON) RawJSON() string {
	return r.raw
}

// Action taken on requests failing validation.
type OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction string

const (
	OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone  OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "none"
	OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog   OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "log"
	OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "block"
)

func (r OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction) IsKnown() bool {
	switch r {
	case OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone, OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog, OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock:
		return true
	}
	return false
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
	Result   []APIShield `json:"result,required"`
	// Whether the API call was successful
	Success OperationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    operationNewResponseEnvelopeJSON    `json:"-"`
}

// operationNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OperationNewResponseEnvelope]
type operationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OperationNewResponseEnvelopeSuccess bool

const (
	OperationNewResponseEnvelopeSuccessTrue OperationNewResponseEnvelopeSuccess = true
)

func (r OperationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OperationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
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
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
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
	Errors   Message              `json:"errors,required"`
	Messages Message              `json:"messages,required"`
	Result   OperationGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success OperationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    operationGetResponseEnvelopeJSON    `json:"-"`
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

// Whether the API call was successful
type OperationGetResponseEnvelopeSuccess bool

const (
	OperationGetResponseEnvelopeSuccessTrue OperationGetResponseEnvelopeSuccess = true
)

func (r OperationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OperationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
