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

// HoldService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewHoldService] method instead.
type HoldService struct {
	Options []option.RequestOption
}

// NewHoldService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHoldService(opts ...option.RequestOption) (r *HoldService) {
	r = &HoldService{}
	r.Options = opts
	return
}

// Retrieve whether the zone is subject to a zone hold, and metadata about the
// hold.
func (r *HoldService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *HoldGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enforce a zone hold on the zone, blocking the creation and activation of zones
// with this zone's hostname.
func (r *HoldService) Enforce(ctx context.Context, zoneID string, body HoldEnforceParams, opts ...option.RequestOption) (res *HoldEnforceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stop enforcement of a zone hold on the zone, permanently or temporarily,
// allowing the creation and activation of zones with this zone's hostname.
func (r *HoldService) Remove(ctx context.Context, zoneID string, body HoldRemoveParams, opts ...option.RequestOption) (res *HoldRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HoldRemoveResponseEnvelope
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HoldGetResponse struct {
	Errors   []HoldGetResponseError   `json:"errors"`
	Messages []HoldGetResponseMessage `json:"messages"`
	Result   HoldGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success HoldGetResponseSuccess `json:"success"`
	JSON    holdGetResponseJSON    `json:"-"`
}

// holdGetResponseJSON contains the JSON metadata for the struct [HoldGetResponse]
type holdGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HoldGetResponseError struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    holdGetResponseErrorJSON `json:"-"`
}

// holdGetResponseErrorJSON contains the JSON metadata for the struct
// [HoldGetResponseError]
type holdGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HoldGetResponseMessage struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    holdGetResponseMessageJSON `json:"-"`
}

// holdGetResponseMessageJSON contains the JSON metadata for the struct
// [HoldGetResponseMessage]
type holdGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HoldGetResponseResult struct {
	Hold              bool                      `json:"hold"`
	HoldAfter         string                    `json:"hold_after"`
	IncludeSubdomains string                    `json:"include_subdomains"`
	JSON              holdGetResponseResultJSON `json:"-"`
}

// holdGetResponseResultJSON contains the JSON metadata for the struct
// [HoldGetResponseResult]
type holdGetResponseResultJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *HoldGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HoldGetResponseSuccess bool

const (
	HoldGetResponseSuccessTrue HoldGetResponseSuccess = true
)

type HoldEnforceResponse struct {
	Errors   []HoldEnforceResponseError   `json:"errors"`
	Messages []HoldEnforceResponseMessage `json:"messages"`
	Result   HoldEnforceResponseResult    `json:"result"`
	// Whether the API call was successful
	Success HoldEnforceResponseSuccess `json:"success"`
	JSON    holdEnforceResponseJSON    `json:"-"`
}

// holdEnforceResponseJSON contains the JSON metadata for the struct
// [HoldEnforceResponse]
type holdEnforceResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldEnforceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HoldEnforceResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    holdEnforceResponseErrorJSON `json:"-"`
}

// holdEnforceResponseErrorJSON contains the JSON metadata for the struct
// [HoldEnforceResponseError]
type holdEnforceResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldEnforceResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HoldEnforceResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    holdEnforceResponseMessageJSON `json:"-"`
}

// holdEnforceResponseMessageJSON contains the JSON metadata for the struct
// [HoldEnforceResponseMessage]
type holdEnforceResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldEnforceResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HoldEnforceResponseResult struct {
	Hold              bool                          `json:"hold"`
	HoldAfter         string                        `json:"hold_after"`
	IncludeSubdomains string                        `json:"include_subdomains"`
	JSON              holdEnforceResponseResultJSON `json:"-"`
}

// holdEnforceResponseResultJSON contains the JSON metadata for the struct
// [HoldEnforceResponseResult]
type holdEnforceResponseResultJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *HoldEnforceResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HoldEnforceResponseSuccess bool

const (
	HoldEnforceResponseSuccessTrue HoldEnforceResponseSuccess = true
)

type HoldRemoveResponse struct {
	Hold              bool                   `json:"hold"`
	HoldAfter         string                 `json:"hold_after"`
	IncludeSubdomains string                 `json:"include_subdomains"`
	JSON              holdRemoveResponseJSON `json:"-"`
}

// holdRemoveResponseJSON contains the JSON metadata for the struct
// [HoldRemoveResponse]
type holdRemoveResponseJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *HoldRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HoldEnforceParams struct {
	// If provided, the zone hold will extend to block any subdomain of the given zone,
	// as well as SSL4SaaS Custom Hostnames. For example, a zone hold on a zone with
	// the hostname 'example.com' and include_subdomains=true will block 'example.com',
	// 'staging.example.com', 'api.staging.example.com', etc.
	IncludeSubdomains param.Field[bool] `query:"include_subdomains"`
}

// URLQuery serializes [HoldEnforceParams]'s query parameters as `url.Values`.
func (r HoldEnforceParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HoldRemoveParams struct {
	// If `hold_after` is provided, the hold will be temporarily disabled, then
	// automatically re-enabled by the system at the time specified in this
	// RFC3339-formatted timestamp. Otherwise, the hold will be disabled indefinitely.
	HoldAfter param.Field[string] `query:"hold_after"`
}

// URLQuery serializes [HoldRemoveParams]'s query parameters as `url.Values`.
func (r HoldRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HoldRemoveResponseEnvelope struct {
	Result HoldRemoveResponse             `json:"result"`
	JSON   holdRemoveResponseEnvelopeJSON `json:"-"`
}

// holdRemoveResponseEnvelopeJSON contains the JSON metadata for the struct
// [HoldRemoveResponseEnvelope]
type holdRemoveResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldRemoveResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
