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

// ZoneAPIGatewaySchemaService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAPIGatewaySchemaService]
// method instead.
type ZoneAPIGatewaySchemaService struct {
	Options []option.RequestOption
}

// NewZoneAPIGatewaySchemaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAPIGatewaySchemaService(opts ...option.RequestOption) (r *ZoneAPIGatewaySchemaService) {
	r = &ZoneAPIGatewaySchemaService{}
	r.Options = opts
	return
}

// Retrieve operations and features as OpenAPI schemas
func (r *ZoneAPIGatewaySchemaService) APIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemas(ctx context.Context, zoneID string, query ZoneAPIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParams, opts ...option.RequestOption) (res *SchemaResponseWithThreshold, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/schemas", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SchemaResponseWithThreshold struct {
	Errors   []SchemaResponseWithThresholdError   `json:"errors"`
	Messages []SchemaResponseWithThresholdMessage `json:"messages"`
	Result   SchemaResponseWithThresholdResult    `json:"result"`
	// Whether the API call was successful
	Success SchemaResponseWithThresholdSuccess `json:"success"`
	JSON    schemaResponseWithThresholdJSON    `json:"-"`
}

// schemaResponseWithThresholdJSON contains the JSON metadata for the struct
// [SchemaResponseWithThreshold]
type schemaResponseWithThresholdJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaResponseWithThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemaResponseWithThresholdError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    schemaResponseWithThresholdErrorJSON `json:"-"`
}

// schemaResponseWithThresholdErrorJSON contains the JSON metadata for the struct
// [SchemaResponseWithThresholdError]
type schemaResponseWithThresholdErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaResponseWithThresholdError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemaResponseWithThresholdMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    schemaResponseWithThresholdMessageJSON `json:"-"`
}

// schemaResponseWithThresholdMessageJSON contains the JSON metadata for the struct
// [SchemaResponseWithThresholdMessage]
type schemaResponseWithThresholdMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaResponseWithThresholdMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemaResponseWithThresholdResult struct {
	Schemas   []interface{}                         `json:"schemas"`
	Timestamp string                                `json:"timestamp"`
	JSON      schemaResponseWithThresholdResultJSON `json:"-"`
}

// schemaResponseWithThresholdResultJSON contains the JSON metadata for the struct
// [SchemaResponseWithThresholdResult]
type schemaResponseWithThresholdResultJSON struct {
	Schemas     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaResponseWithThresholdResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemaResponseWithThresholdSuccess bool

const (
	SchemaResponseWithThresholdSuccessTrue SchemaResponseWithThresholdSuccess = true
)

type ZoneAPIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParams struct {
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]ZoneAPIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeature] `query:"feature"`
	// Receive schema only for the given host(s).
	Host param.Field[[]string] `query:"host"`
}

// URLQuery serializes
// [ZoneAPIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParams]'s
// query parameters as `url.Values`.
func (r ZoneAPIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneAPIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeature string

const (
	ZoneAPIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeatureThresholds       ZoneAPIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeature = "thresholds"
	ZoneAPIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeatureParameterSchemas ZoneAPIGatewaySchemaAPIShieldEndpointManagementGetOperationsAndFeaturesAsOpenAPISchemasParamsFeature = "parameter_schemas"
)
