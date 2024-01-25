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

// AccountLoadBalancerRegionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountLoadBalancerRegionService] method instead.
type AccountLoadBalancerRegionService struct {
	Options []option.RequestOption
}

// NewAccountLoadBalancerRegionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerRegionService(opts ...option.RequestOption) (r *AccountLoadBalancerRegionService) {
	r = &AccountLoadBalancerRegionService{}
	r.Options = opts
	return
}

// Get a single region mapping.
func (r *AccountLoadBalancerRegionService) Get(ctx context.Context, accountIdentifier string, regionCode AccountLoadBalancerRegionGetParamsRegionCode, opts ...option.RequestOption) (res *AccountLoadBalancerRegionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/regions/%v", accountIdentifier, regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all region mappings.
func (r *AccountLoadBalancerRegionService) LoadBalancerRegionsListRegions(ctx context.Context, accountIdentifier string, query AccountLoadBalancerRegionLoadBalancerRegionsListRegionsParams, opts ...option.RequestOption) (res *AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/regions", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountLoadBalancerRegionGetResponse struct {
	Errors   []AccountLoadBalancerRegionGetResponseError   `json:"errors"`
	Messages []AccountLoadBalancerRegionGetResponseMessage `json:"messages"`
	// A list of countries and subdivisions mapped to a region.
	Result interface{} `json:"result"`
	// Whether the API call was successful
	Success AccountLoadBalancerRegionGetResponseSuccess `json:"success"`
	JSON    accountLoadBalancerRegionGetResponseJSON    `json:"-"`
}

// accountLoadBalancerRegionGetResponseJSON contains the JSON metadata for the
// struct [AccountLoadBalancerRegionGetResponse]
type accountLoadBalancerRegionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerRegionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerRegionGetResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountLoadBalancerRegionGetResponseErrorJSON `json:"-"`
}

// accountLoadBalancerRegionGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountLoadBalancerRegionGetResponseError]
type accountLoadBalancerRegionGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerRegionGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerRegionGetResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountLoadBalancerRegionGetResponseMessageJSON `json:"-"`
}

// accountLoadBalancerRegionGetResponseMessageJSON contains the JSON metadata for
// the struct [AccountLoadBalancerRegionGetResponseMessage]
type accountLoadBalancerRegionGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerRegionGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLoadBalancerRegionGetResponseSuccess bool

const (
	AccountLoadBalancerRegionGetResponseSuccessTrue AccountLoadBalancerRegionGetResponseSuccess = true
)

type AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponse struct {
	Errors   []AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseError   `json:"errors"`
	Messages []AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseMessage `json:"messages"`
	Result   interface{}                                                              `json:"result"`
	// Whether the API call was successful
	Success AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseSuccess `json:"success"`
	JSON    accountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseJSON    `json:"-"`
}

// accountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseJSON contains the
// JSON metadata for the struct
// [AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponse]
type accountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseError struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    accountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseErrorJSON `json:"-"`
}

// accountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseError]
type accountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseMessage struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseMessageJSON `json:"-"`
}

// accountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseMessage]
type accountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseSuccess bool

const (
	AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseSuccessTrue AccountLoadBalancerRegionLoadBalancerRegionsListRegionsResponseSuccess = true
)

// A list of Cloudflare regions. WNAM: Western North America, ENAM: Eastern North
// America, WEU: Western Europe, EEU: Eastern Europe, NSAM: Northern South America,
// SSAM: Southern South America, OC: Oceania, ME: Middle East, NAF: North Africa,
// SAF: South Africa, SAS: Southern Asia, SEAS: South East Asia, NEAS: North East
// Asia).
type AccountLoadBalancerRegionGetParamsRegionCode string

const (
	AccountLoadBalancerRegionGetParamsRegionCodeWnam AccountLoadBalancerRegionGetParamsRegionCode = "WNAM"
	AccountLoadBalancerRegionGetParamsRegionCodeEnam AccountLoadBalancerRegionGetParamsRegionCode = "ENAM"
	AccountLoadBalancerRegionGetParamsRegionCodeWeu  AccountLoadBalancerRegionGetParamsRegionCode = "WEU"
	AccountLoadBalancerRegionGetParamsRegionCodeEeu  AccountLoadBalancerRegionGetParamsRegionCode = "EEU"
	AccountLoadBalancerRegionGetParamsRegionCodeNsam AccountLoadBalancerRegionGetParamsRegionCode = "NSAM"
	AccountLoadBalancerRegionGetParamsRegionCodeSsam AccountLoadBalancerRegionGetParamsRegionCode = "SSAM"
	AccountLoadBalancerRegionGetParamsRegionCodeOc   AccountLoadBalancerRegionGetParamsRegionCode = "OC"
	AccountLoadBalancerRegionGetParamsRegionCodeMe   AccountLoadBalancerRegionGetParamsRegionCode = "ME"
	AccountLoadBalancerRegionGetParamsRegionCodeNaf  AccountLoadBalancerRegionGetParamsRegionCode = "NAF"
	AccountLoadBalancerRegionGetParamsRegionCodeSaf  AccountLoadBalancerRegionGetParamsRegionCode = "SAF"
	AccountLoadBalancerRegionGetParamsRegionCodeSas  AccountLoadBalancerRegionGetParamsRegionCode = "SAS"
	AccountLoadBalancerRegionGetParamsRegionCodeSeas AccountLoadBalancerRegionGetParamsRegionCode = "SEAS"
	AccountLoadBalancerRegionGetParamsRegionCodeNeas AccountLoadBalancerRegionGetParamsRegionCode = "NEAS"
)

type AccountLoadBalancerRegionLoadBalancerRegionsListRegionsParams struct {
	// Two-letter alpha-2 country code followed in ISO 3166-1.
	CountryCodeA2 param.Field[string] `query:"country_code_a2"`
	// Two-letter subdivision code followed in ISO 3166-2.
	SubdivisionCode param.Field[string] `query:"subdivision_code"`
	// Two-letter subdivision code followed in ISO 3166-2.
	SubdivisionCodeA2 param.Field[string] `query:"subdivision_code_a2"`
}

// URLQuery serializes
// [AccountLoadBalancerRegionLoadBalancerRegionsListRegionsParams]'s query
// parameters as `url.Values`.
func (r AccountLoadBalancerRegionLoadBalancerRegionsListRegionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
