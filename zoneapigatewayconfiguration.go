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

// ZoneAPIGatewayConfigurationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneAPIGatewayConfigurationService] method instead.
type ZoneAPIGatewayConfigurationService struct {
	Options []option.RequestOption
}

// NewZoneAPIGatewayConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAPIGatewayConfigurationService(opts ...option.RequestOption) (r *ZoneAPIGatewayConfigurationService) {
	r = &ZoneAPIGatewayConfigurationService{}
	r.Options = opts
	return
}

// Retrieve information about specific configuration properties
func (r *ZoneAPIGatewayConfigurationService) APIShieldSettingsGetInformationAboutSpecificConfigurationProperties(ctx context.Context, zoneID string, query ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParams, opts ...option.RequestOption) (res *SingleResponsePwGei7T9, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/configuration", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Set configuration properties
func (r *ZoneAPIGatewayConfigurationService) APIShieldSettingsSetConfigurationProperties(ctx context.Context, zoneID string, body ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/configuration", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SingleResponsePwGei7T9 struct {
	Errors   []SingleResponsePwGei7T9Error   `json:"errors"`
	Messages []SingleResponsePwGei7T9Message `json:"messages"`
	Result   SingleResponsePwGei7T9Result    `json:"result"`
	// Whether the API call was successful
	Success SingleResponsePwGei7T9Success `json:"success"`
	JSON    singleResponsePwGei7T9JSON    `json:"-"`
}

// singleResponsePwGei7T9JSON contains the JSON metadata for the struct
// [SingleResponsePwGei7T9]
type singleResponsePwGei7T9JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponsePwGei7T9) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponsePwGei7T9Error struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleResponsePwGei7T9ErrorJSON `json:"-"`
}

// singleResponsePwGei7T9ErrorJSON contains the JSON metadata for the struct
// [SingleResponsePwGei7T9Error]
type singleResponsePwGei7T9ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponsePwGei7T9Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponsePwGei7T9Message struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    singleResponsePwGei7T9MessageJSON `json:"-"`
}

// singleResponsePwGei7T9MessageJSON contains the JSON metadata for the struct
// [SingleResponsePwGei7T9Message]
type singleResponsePwGei7T9MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponsePwGei7T9Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponsePwGei7T9Result struct {
	AuthIDCharacteristics []SingleResponsePwGei7T9ResultAuthIDCharacteristic `json:"auth_id_characteristics"`
	JSON                  singleResponsePwGei7T9ResultJSON                   `json:"-"`
}

// singleResponsePwGei7T9ResultJSON contains the JSON metadata for the struct
// [SingleResponsePwGei7T9Result]
type singleResponsePwGei7T9ResultJSON struct {
	AuthIDCharacteristics apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SingleResponsePwGei7T9Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponsePwGei7T9ResultAuthIDCharacteristic struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	Name string `json:"name,required"`
	// The type of characteristic.
	Type SingleResponsePwGei7T9ResultAuthIDCharacteristicsType `json:"type,required"`
	JSON singleResponsePwGei7T9ResultAuthIDCharacteristicJSON  `json:"-"`
}

// singleResponsePwGei7T9ResultAuthIDCharacteristicJSON contains the JSON metadata
// for the struct [SingleResponsePwGei7T9ResultAuthIDCharacteristic]
type singleResponsePwGei7T9ResultAuthIDCharacteristicJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponsePwGei7T9ResultAuthIDCharacteristic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of characteristic.
type SingleResponsePwGei7T9ResultAuthIDCharacteristicsType string

const (
	SingleResponsePwGei7T9ResultAuthIDCharacteristicsTypeHeader SingleResponsePwGei7T9ResultAuthIDCharacteristicsType = "header"
	SingleResponsePwGei7T9ResultAuthIDCharacteristicsTypeCookie SingleResponsePwGei7T9ResultAuthIDCharacteristicsType = "cookie"
)

// Whether the API call was successful
type SingleResponsePwGei7T9Success bool

const (
	SingleResponsePwGei7T9SuccessTrue SingleResponsePwGei7T9Success = true
)

type ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParams struct {
	// Requests information about certain properties.
	Properties param.Field[[]ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParamsProperty] `query:"properties"`
}

// URLQuery serializes
// [ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParams]'s
// query parameters as `url.Values`.
func (r ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParamsProperty string

const (
	ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParamsPropertyAuthIDCharacteristics ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParamsProperty = "auth_id_characteristics"
)

type ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParams struct {
	AuthIDCharacteristics param.Field[[]ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristic] `json:"auth_id_characteristics"`
}

func (r ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristic struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	Name param.Field[string] `json:"name,required"`
	// The type of characteristic.
	Type param.Field[ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristicsType] `json:"type,required"`
}

func (r ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of characteristic.
type ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristicsType string

const (
	ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristicsTypeHeader ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristicsType = "header"
	ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristicsTypeCookie ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristicsType = "cookie"
)
