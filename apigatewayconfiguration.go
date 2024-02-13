// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// APIGatewayConfigurationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAPIGatewayConfigurationService] method instead.
type APIGatewayConfigurationService struct {
	Options []option.RequestOption
}

// NewAPIGatewayConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIGatewayConfigurationService(opts ...option.RequestOption) (r *APIGatewayConfigurationService) {
	r = &APIGatewayConfigurationService{}
	r.Options = opts
	return
}

// Retrieve information about specific configuration properties
func (r *APIGatewayConfigurationService) APIShieldSettingsGetInformationAboutSpecificConfigurationProperties(ctx context.Context, zoneID string, query APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParams, opts ...option.RequestOption) (res *APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/configuration", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Set configuration properties
func (r *APIGatewayConfigurationService) APIShieldSettingsSetConfigurationProperties(ctx context.Context, zoneID string, body APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParams, opts ...option.RequestOption) (res *APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIShieldAPIResponseSingle
	path := fmt.Sprintf("zones/%s/api_gateway/configuration", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponse struct {
	AuthIDCharacteristics []APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseAuthIDCharacteristic `json:"auth_id_characteristics"`
	JSON                  apiGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseJSON                   `json:"-"`
}

// apiGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseJSON
// contains the JSON metadata for the struct
// [APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponse]
type apiGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseJSON struct {
	AuthIDCharacteristics apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseAuthIDCharacteristic struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	Name string `json:"name,required"`
	// The type of characteristic.
	Type APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseAuthIDCharacteristicsType `json:"type,required"`
	JSON apiGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseAuthIDCharacteristicJSON  `json:"-"`
}

// apiGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseAuthIDCharacteristicJSON
// contains the JSON metadata for the struct
// [APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseAuthIDCharacteristic]
type apiGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseAuthIDCharacteristicJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseAuthIDCharacteristic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of characteristic.
type APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseAuthIDCharacteristicsType string

const (
	APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseAuthIDCharacteristicsTypeHeader APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseAuthIDCharacteristicsType = "header"
	APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseAuthIDCharacteristicsTypeCookie APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseAuthIDCharacteristicsType = "cookie"
)

// Union satisfied by
// [APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseUnknown]
// or [shared.UnionString].
type APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponse interface {
	ImplementsAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParams struct {
	// Requests information about certain properties.
	Properties param.Field[[]APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParamsProperty] `query:"properties"`
}

// URLQuery serializes
// [APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParams]'s
// query parameters as `url.Values`.
func (r APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParamsProperty string

const (
	APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParamsPropertyAuthIDCharacteristics APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParamsProperty = "auth_id_characteristics"
)

type APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseEnvelope struct {
	Errors   APIShieldMessages                                                                                  `json:"errors,required"`
	Messages APIShieldMessages                                                                                  `json:"messages,required"`
	Result   APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                                                                                           `json:"success,required"`
	JSON    apiGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseEnvelopeJSON `json:"-"`
}

// apiGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseEnvelope]
type apiGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParams struct {
	AuthIDCharacteristics param.Field[[]APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristic] `json:"auth_id_characteristics"`
}

func (r APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristic struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	Name param.Field[string] `json:"name,required"`
	// The type of characteristic.
	Type param.Field[APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristicsType] `json:"type,required"`
}

func (r APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of characteristic.
type APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristicsType string

const (
	APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristicsTypeHeader APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristicsType = "header"
	APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristicsTypeCookie APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristicsType = "cookie"
)

type APIShieldAPIResponseSingle struct {
	Errors   APIShieldMessages                                                          `json:"errors,required"`
	Messages APIShieldMessages                                                          `json:"messages,required"`
	Result   APIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponse `json:"result,required"`
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
