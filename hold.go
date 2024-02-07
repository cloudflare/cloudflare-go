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
	var env HoldGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enforce a zone hold on the zone, blocking the creation and activation of zones
// with this zone's hostname.
func (r *HoldService) Enforce(ctx context.Context, zoneID string, body HoldEnforceParams, opts ...option.RequestOption) (res *HoldEnforceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HoldEnforceResponseEnvelope
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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
	Hold              bool                `json:"hold"`
	HoldAfter         string              `json:"hold_after"`
	IncludeSubdomains string              `json:"include_subdomains"`
	JSON              holdGetResponseJSON `json:"-"`
}

// holdGetResponseJSON contains the JSON metadata for the struct [HoldGetResponse]
type holdGetResponseJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *HoldGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HoldEnforceResponse struct {
	Hold              bool                    `json:"hold"`
	HoldAfter         string                  `json:"hold_after"`
	IncludeSubdomains string                  `json:"include_subdomains"`
	JSON              holdEnforceResponseJSON `json:"-"`
}

// holdEnforceResponseJSON contains the JSON metadata for the struct
// [HoldEnforceResponse]
type holdEnforceResponseJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *HoldEnforceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type HoldGetResponseEnvelope struct {
	Errors   []HoldGetResponseEnvelopeErrors   `json:"errors"`
	Messages []HoldGetResponseEnvelopeMessages `json:"messages"`
	Result   HoldGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success HoldGetResponseEnvelopeSuccess `json:"success"`
	JSON    holdGetResponseEnvelopeJSON    `json:"-"`
}

// holdGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HoldGetResponseEnvelope]
type holdGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HoldGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    holdGetResponseEnvelopeErrorsJSON `json:"-"`
}

// holdGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [HoldGetResponseEnvelopeErrors]
type holdGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HoldGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    holdGetResponseEnvelopeMessagesJSON `json:"-"`
}

// holdGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [HoldGetResponseEnvelopeMessages]
type holdGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HoldGetResponseEnvelopeSuccess bool

const (
	HoldGetResponseEnvelopeSuccessTrue HoldGetResponseEnvelopeSuccess = true
)

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

type HoldEnforceResponseEnvelope struct {
	Errors   []HoldEnforceResponseEnvelopeErrors   `json:"errors"`
	Messages []HoldEnforceResponseEnvelopeMessages `json:"messages"`
	Result   HoldEnforceResponse                   `json:"result"`
	// Whether the API call was successful
	Success HoldEnforceResponseEnvelopeSuccess `json:"success"`
	JSON    holdEnforceResponseEnvelopeJSON    `json:"-"`
}

// holdEnforceResponseEnvelopeJSON contains the JSON metadata for the struct
// [HoldEnforceResponseEnvelope]
type holdEnforceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldEnforceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HoldEnforceResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    holdEnforceResponseEnvelopeErrorsJSON `json:"-"`
}

// holdEnforceResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [HoldEnforceResponseEnvelopeErrors]
type holdEnforceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldEnforceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HoldEnforceResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    holdEnforceResponseEnvelopeMessagesJSON `json:"-"`
}

// holdEnforceResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HoldEnforceResponseEnvelopeMessages]
type holdEnforceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldEnforceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HoldEnforceResponseEnvelopeSuccess bool

const (
	HoldEnforceResponseEnvelopeSuccessTrue HoldEnforceResponseEnvelopeSuccess = true
)

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
