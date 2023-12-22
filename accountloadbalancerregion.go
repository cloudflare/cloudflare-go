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
func (r *AccountLoadBalancerRegionService) Get(ctx context.Context, accountIdentifier string, regionCode AccountLoadBalancerRegionGetParamsRegionCode, opts ...option.RequestOption) (res *RegionSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/regions/%v", accountIdentifier, regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all region mappings.
func (r *AccountLoadBalancerRegionService) LoadBalancerRegionsListRegions(ctx context.Context, accountIdentifier string, query AccountLoadBalancerRegionLoadBalancerRegionsListRegionsParams, opts ...option.RequestOption) (res *RegionResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/regions", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RegionResponseCollection struct {
	Errors   []RegionResponseCollectionError   `json:"errors"`
	Messages []RegionResponseCollectionMessage `json:"messages"`
	Result   interface{}                       `json:"result"`
	// Whether the API call was successful
	Success RegionResponseCollectionSuccess `json:"success"`
	JSON    regionResponseCollectionJSON    `json:"-"`
}

// regionResponseCollectionJSON contains the JSON metadata for the struct
// [RegionResponseCollection]
type regionResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RegionResponseCollectionError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    regionResponseCollectionErrorJSON `json:"-"`
}

// regionResponseCollectionErrorJSON contains the JSON metadata for the struct
// [RegionResponseCollectionError]
type regionResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RegionResponseCollectionMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    regionResponseCollectionMessageJSON `json:"-"`
}

// regionResponseCollectionMessageJSON contains the JSON metadata for the struct
// [RegionResponseCollectionMessage]
type regionResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RegionResponseCollectionSuccess bool

const (
	RegionResponseCollectionSuccessTrue RegionResponseCollectionSuccess = true
)

type RegionSingleResponse struct {
	Errors   []RegionSingleResponseError   `json:"errors"`
	Messages []RegionSingleResponseMessage `json:"messages"`
	// A list of countries and subdivisions mapped to a region.
	Result interface{} `json:"result"`
	// Whether the API call was successful
	Success RegionSingleResponseSuccess `json:"success"`
	JSON    regionSingleResponseJSON    `json:"-"`
}

// regionSingleResponseJSON contains the JSON metadata for the struct
// [RegionSingleResponse]
type regionSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RegionSingleResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    regionSingleResponseErrorJSON `json:"-"`
}

// regionSingleResponseErrorJSON contains the JSON metadata for the struct
// [RegionSingleResponseError]
type regionSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RegionSingleResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    regionSingleResponseMessageJSON `json:"-"`
}

// regionSingleResponseMessageJSON contains the JSON metadata for the struct
// [RegionSingleResponseMessage]
type regionSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RegionSingleResponseSuccess bool

const (
	RegionSingleResponseSuccessTrue RegionSingleResponseSuccess = true
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
