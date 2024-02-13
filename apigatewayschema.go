// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// APIGatewaySchemaService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAPIGatewaySchemaService] method
// instead.
type APIGatewaySchemaService struct {
	Options []option.RequestOption
}

// NewAPIGatewaySchemaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIGatewaySchemaService(opts ...option.RequestOption) (r *APIGatewaySchemaService) {
	r = &APIGatewaySchemaService{}
	r.Options = opts
	return
}

// Updates operation-level schema validation settings on the zone
func (r *APIGatewaySchemaService) Update(ctx context.Context, zoneID string, operationID string, body APIGatewaySchemaUpdateParams, opts ...option.RequestOption) (res *APIGatewaySchemaUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s/schema_validation", zoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve operations and features as OpenAPI schemas
func (r *APIGatewaySchemaService) APIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemas(ctx context.Context, zoneID string, query APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParams, opts ...option.RequestOption) (res *APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIShieldAPIResponseSingle
	path := fmt.Sprintf("zones/%s/api_gateway/schemas", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves operation-level schema validation settings on the zone
func (r *APIGatewaySchemaService) Get(ctx context.Context, zoneID string, operationID string, opts ...option.RequestOption) (res *APIGatewaySchemaGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s/schema_validation", zoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves zone level schema validation settings currently set on the zone
func (r *APIGatewaySchemaService) GetIncremental(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *APIShieldZoneSchemaValidationSettings, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/settings/schema_validation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates multiple operation-level schema validation settings on the zone
func (r *APIGatewaySchemaService) UpdateMultiple(ctx context.Context, zoneID string, body APIGatewaySchemaUpdateMultipleParams, opts ...option.RequestOption) (res *APIGatewaySchemaUpdateMultipleResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewaySchemaUpdateMultipleResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/operations/schema_validation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type APIGatewaySchemaUpdateResponse struct {
	// When set, this applies a mitigation action to this operation
	//
	//   - `log` log request when request does not conform to schema for this operation
	//   - `block` deny access to the site when request does not conform to schema for
	//     this operation
	//   - `none` will skip mitigation for this operation
	//   - `null` indicates that no operation level mitigation is in place, see Zone
	//     Level Schema Validation Settings for mitigation action that will be applied
	MitigationAction APIGatewaySchemaUpdateResponseMitigationAction `json:"mitigation_action,nullable"`
	JSON             apiGatewaySchemaUpdateResponseJSON             `json:"-"`
}

// apiGatewaySchemaUpdateResponseJSON contains the JSON metadata for the struct
// [APIGatewaySchemaUpdateResponse]
type apiGatewaySchemaUpdateResponseJSON struct {
	MitigationAction apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIGatewaySchemaUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set, this applies a mitigation action to this operation
//
//   - `log` log request when request does not conform to schema for this operation
//   - `block` deny access to the site when request does not conform to schema for
//     this operation
//   - `none` will skip mitigation for this operation
//   - `null` indicates that no operation level mitigation is in place, see Zone
//     Level Schema Validation Settings for mitigation action that will be applied
type APIGatewaySchemaUpdateResponseMitigationAction string

const (
	APIGatewaySchemaUpdateResponseMitigationActionLog   APIGatewaySchemaUpdateResponseMitigationAction = "log"
	APIGatewaySchemaUpdateResponseMitigationActionBlock APIGatewaySchemaUpdateResponseMitigationAction = "block"
	APIGatewaySchemaUpdateResponseMitigationActionNone  APIGatewaySchemaUpdateResponseMitigationAction = "none"
)

type APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasResponse struct {
	Schemas   []interface{}                                                                                   `json:"schemas"`
	Timestamp string                                                                                          `json:"timestamp"`
	JSON      apiGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasResponseJSON `json:"-"`
}

// apiGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasResponseJSON
// contains the JSON metadata for the struct
// [APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasResponse]
type apiGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasResponseJSON struct {
	Schemas     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewaySchemaGetResponse struct {
	// When set, this applies a mitigation action to this operation
	//
	//   - `log` log request when request does not conform to schema for this operation
	//   - `block` deny access to the site when request does not conform to schema for
	//     this operation
	//   - `none` will skip mitigation for this operation
	//   - `null` indicates that no operation level mitigation is in place, see Zone
	//     Level Schema Validation Settings for mitigation action that will be applied
	MitigationAction APIGatewaySchemaGetResponseMitigationAction `json:"mitigation_action,nullable"`
	JSON             apiGatewaySchemaGetResponseJSON             `json:"-"`
}

// apiGatewaySchemaGetResponseJSON contains the JSON metadata for the struct
// [APIGatewaySchemaGetResponse]
type apiGatewaySchemaGetResponseJSON struct {
	MitigationAction apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIGatewaySchemaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set, this applies a mitigation action to this operation
//
//   - `log` log request when request does not conform to schema for this operation
//   - `block` deny access to the site when request does not conform to schema for
//     this operation
//   - `none` will skip mitigation for this operation
//   - `null` indicates that no operation level mitigation is in place, see Zone
//     Level Schema Validation Settings for mitigation action that will be applied
type APIGatewaySchemaGetResponseMitigationAction string

const (
	APIGatewaySchemaGetResponseMitigationActionLog   APIGatewaySchemaGetResponseMitigationAction = "log"
	APIGatewaySchemaGetResponseMitigationActionBlock APIGatewaySchemaGetResponseMitigationAction = "block"
	APIGatewaySchemaGetResponseMitigationActionNone  APIGatewaySchemaGetResponseMitigationAction = "none"
)

type APIGatewaySchemaUpdateMultipleResponse map[string]APIGatewaySchemaUpdateMultipleResponse

type APIGatewaySchemaUpdateParams struct {
	// When set, this applies a mitigation action to this operation
	//
	//   - `log` log request when request does not conform to schema for this operation
	//   - `block` deny access to the site when request does not conform to schema for
	//     this operation
	//   - `none` will skip mitigation for this operation
	//   - `null` indicates that no operation level mitigation is in place, see Zone
	//     Level Schema Validation Settings for mitigation action that will be applied
	MitigationAction param.Field[APIGatewaySchemaUpdateParamsMitigationAction] `json:"mitigation_action"`
}

func (r APIGatewaySchemaUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When set, this applies a mitigation action to this operation
//
//   - `log` log request when request does not conform to schema for this operation
//   - `block` deny access to the site when request does not conform to schema for
//     this operation
//   - `none` will skip mitigation for this operation
//   - `null` indicates that no operation level mitigation is in place, see Zone
//     Level Schema Validation Settings for mitigation action that will be applied
type APIGatewaySchemaUpdateParamsMitigationAction string

const (
	APIGatewaySchemaUpdateParamsMitigationActionLog   APIGatewaySchemaUpdateParamsMitigationAction = "log"
	APIGatewaySchemaUpdateParamsMitigationActionBlock APIGatewaySchemaUpdateParamsMitigationAction = "block"
	APIGatewaySchemaUpdateParamsMitigationActionNone  APIGatewaySchemaUpdateParamsMitigationAction = "none"
)

type APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParams struct {
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeature] `query:"feature"`
	// Receive schema only for the given host(s).
	Host param.Field[[]string] `query:"host"`
}

// URLQuery serializes
// [APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParams]'s
// query parameters as `url.Values`.
func (r APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeature string

const (
	APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeatureThresholds       APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeature = "thresholds"
	APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeatureParameterSchemas APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeature = "parameter_schemas"
	APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeatureSchemaInfo       APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeature = "schema_info"
)

type APIShieldAPIResponseSingle struct {
	Result APIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasResponse `json:"result,required"`
	JSON   apiShieldAPIResponseSingleJSON                                                              `json:"-"`
}

// apiShieldAPIResponseSingleJSON contains the JSON metadata for the struct
// [APIShieldAPIResponseSingle]
type apiShieldAPIResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewaySchemaUpdateMultipleParams struct {
	Body param.Field[map[string]APIGatewaySchemaUpdateMultipleParamsBody] `json:"body,required"`
}

func (r APIGatewaySchemaUpdateMultipleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Operation ID to mitigation action mappings
type APIGatewaySchemaUpdateMultipleParamsBody struct {
	// When set, this applies a mitigation action to this operation
	//
	//   - `log` log request when request does not conform to schema for this operation
	//   - `block` deny access to the site when request does not conform to schema for
	//     this operation
	//   - `none` will skip mitigation for this operation
	//   - `null` indicates that no operation level mitigation is in place, see Zone
	//     Level Schema Validation Settings for mitigation action that will be applied
	MitigationAction param.Field[APIGatewaySchemaUpdateMultipleParamsBodyMitigationAction] `json:"mitigation_action"`
}

func (r APIGatewaySchemaUpdateMultipleParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When set, this applies a mitigation action to this operation
//
//   - `log` log request when request does not conform to schema for this operation
//   - `block` deny access to the site when request does not conform to schema for
//     this operation
//   - `none` will skip mitigation for this operation
//   - `null` indicates that no operation level mitigation is in place, see Zone
//     Level Schema Validation Settings for mitigation action that will be applied
type APIGatewaySchemaUpdateMultipleParamsBodyMitigationAction string

const (
	APIGatewaySchemaUpdateMultipleParamsBodyMitigationActionLog   APIGatewaySchemaUpdateMultipleParamsBodyMitigationAction = "log"
	APIGatewaySchemaUpdateMultipleParamsBodyMitigationActionBlock APIGatewaySchemaUpdateMultipleParamsBodyMitigationAction = "block"
	APIGatewaySchemaUpdateMultipleParamsBodyMitigationActionNone  APIGatewaySchemaUpdateMultipleParamsBodyMitigationAction = "none"
)

type APIGatewaySchemaUpdateMultipleResponseEnvelope struct {
	Errors   APIShieldMessages                      `json:"errors,required"`
	Messages APIShieldMessages                      `json:"messages,required"`
	Result   APIGatewaySchemaUpdateMultipleResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                               `json:"success,required"`
	JSON    apiGatewaySchemaUpdateMultipleResponseEnvelopeJSON `json:"-"`
}

// apiGatewaySchemaUpdateMultipleResponseEnvelopeJSON contains the JSON metadata
// for the struct [APIGatewaySchemaUpdateMultipleResponseEnvelope]
type apiGatewaySchemaUpdateMultipleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewaySchemaUpdateMultipleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
