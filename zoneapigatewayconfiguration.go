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
func (r *ZoneAPIGatewayConfigurationService) APIShieldSettingsGetInformationAboutSpecificConfigurationProperties(ctx context.Context, zoneID string, query ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParams, opts ...option.RequestOption) (res *ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/configuration", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Set configuration properties
func (r *ZoneAPIGatewayConfigurationService) APIShieldSettingsSetConfigurationProperties(ctx context.Context, zoneID string, body ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParams, opts ...option.RequestOption) (res *ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/configuration", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponse struct {
	Errors   []ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseError   `json:"errors"`
	Messages []ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseMessage `json:"messages"`
	Result   ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                                                                                                       `json:"success"`
	JSON    zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseJSON `json:"-"`
}

// zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponse]
type zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseError struct {
	Code    int64                                                                                                           `json:"code,required"`
	Message string                                                                                                          `json:"message,required"`
	JSON    zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseErrorJSON `json:"-"`
}

// zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseError]
type zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseMessage struct {
	Code    int64                                                                                                             `json:"code,required"`
	Message string                                                                                                            `json:"message,required"`
	JSON    zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseMessageJSON `json:"-"`
}

// zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseMessage]
type zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResult struct {
	AuthIDCharacteristics []ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultAuthIDCharacteristic `json:"auth_id_characteristics"`
	JSON                  zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultJSON                   `json:"-"`
}

// zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResult]
type zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultJSON struct {
	AuthIDCharacteristics apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultAuthIDCharacteristic struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	Name string `json:"name,required"`
	// The type of characteristic.
	Type ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultAuthIDCharacteristicsType `json:"type,required"`
	JSON zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultAuthIDCharacteristicJSON  `json:"-"`
}

// zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultAuthIDCharacteristicJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultAuthIDCharacteristic]
type zoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultAuthIDCharacteristicJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultAuthIDCharacteristic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of characteristic.
type ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultAuthIDCharacteristicsType string

const (
	ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultAuthIDCharacteristicsTypeHeader ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultAuthIDCharacteristicsType = "header"
	ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultAuthIDCharacteristicsTypeCookie ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesResponseResultAuthIDCharacteristicsType = "cookie"
)

type ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponse struct {
	Errors   []ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseError   `json:"errors"`
	Messages []ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseMessage `json:"messages"`
	Result   ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                                                                               `json:"success"`
	JSON    zoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseJSON `json:"-"`
}

// zoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponse]
type zoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseError struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    zoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseErrorJSON `json:"-"`
}

// zoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseError]
type zoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseMessage struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    zoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseMessageJSON `json:"-"`
}

// zoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseMessage]
type zoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseResultUnknown]
// or [shared.UnionString].
type ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseResult interface {
	ImplementsZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

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
