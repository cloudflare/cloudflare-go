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

// ZoneHoldService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZoneHoldService] method instead.
type ZoneHoldService struct {
	Options []option.RequestOption
}

// NewZoneHoldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneHoldService(opts ...option.RequestOption) (r *ZoneHoldService) {
	r = &ZoneHoldService{}
	r.Options = opts
	return
}

// Enforce a zone hold on the zone, blocking the creation and activation of zones
// with this zone's hostname.
func (r *ZoneHoldService) Enforce(ctx context.Context, zoneID string, body ZoneHoldEnforceParams, opts ...option.RequestOption) (res *ZoneHoldEnforceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneHoldEnforceResponseEnvelope
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve whether the zone is subject to a zone hold, and metadata about the
// hold.
func (r *ZoneHoldService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneHoldGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneHoldGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Stop enforcement of a zone hold on the zone, permanently or temporarily,
// allowing the creation and activation of zones with this zone's hostname.
func (r *ZoneHoldService) Remove(ctx context.Context, zoneID string, body ZoneHoldRemoveParams, opts ...option.RequestOption) (res *ZoneHoldRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneHoldRemoveResponseEnvelope
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZoneHoldEnforceResponse struct {
	Hold              bool                        `json:"hold"`
	HoldAfter         string                      `json:"hold_after"`
	IncludeSubdomains string                      `json:"include_subdomains"`
	JSON              zoneHoldEnforceResponseJSON `json:"-"`
}

// zoneHoldEnforceResponseJSON contains the JSON metadata for the struct
// [ZoneHoldEnforceResponse]
type zoneHoldEnforceResponseJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneHoldEnforceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldGetResponse struct {
	Hold              bool                    `json:"hold"`
	HoldAfter         string                  `json:"hold_after"`
	IncludeSubdomains string                  `json:"include_subdomains"`
	JSON              zoneHoldGetResponseJSON `json:"-"`
}

// zoneHoldGetResponseJSON contains the JSON metadata for the struct
// [ZoneHoldGetResponse]
type zoneHoldGetResponseJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneHoldGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldRemoveResponse struct {
	Hold              bool                       `json:"hold"`
	HoldAfter         string                     `json:"hold_after"`
	IncludeSubdomains string                     `json:"include_subdomains"`
	JSON              zoneHoldRemoveResponseJSON `json:"-"`
}

// zoneHoldRemoveResponseJSON contains the JSON metadata for the struct
// [ZoneHoldRemoveResponse]
type zoneHoldRemoveResponseJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneHoldRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldEnforceParams struct {
	// If provided, the zone hold will extend to block any subdomain of the given zone,
	// as well as SSL4SaaS Custom Hostnames. For example, a zone hold on a zone with
	// the hostname 'example.com' and include_subdomains=true will block 'example.com',
	// 'staging.example.com', 'api.staging.example.com', etc.
	IncludeSubdomains param.Field[bool] `query:"include_subdomains"`
}

// URLQuery serializes [ZoneHoldEnforceParams]'s query parameters as `url.Values`.
func (r ZoneHoldEnforceParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneHoldEnforceResponseEnvelope struct {
	Errors   []ZoneHoldEnforceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneHoldEnforceResponseEnvelopeMessages `json:"messages,required"`
	Result   ZoneHoldEnforceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZoneHoldEnforceResponseEnvelopeSuccess `json:"success,required"`
	JSON    zoneHoldEnforceResponseEnvelopeJSON    `json:"-"`
}

// zoneHoldEnforceResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneHoldEnforceResponseEnvelope]
type zoneHoldEnforceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldEnforceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldEnforceResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneHoldEnforceResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneHoldEnforceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneHoldEnforceResponseEnvelopeErrors]
type zoneHoldEnforceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldEnforceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldEnforceResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneHoldEnforceResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneHoldEnforceResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZoneHoldEnforceResponseEnvelopeMessages]
type zoneHoldEnforceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldEnforceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneHoldEnforceResponseEnvelopeSuccess bool

const (
	ZoneHoldEnforceResponseEnvelopeSuccessTrue ZoneHoldEnforceResponseEnvelopeSuccess = true
)

type ZoneHoldGetResponseEnvelope struct {
	Errors   []ZoneHoldGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneHoldGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZoneHoldGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZoneHoldGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zoneHoldGetResponseEnvelopeJSON    `json:"-"`
}

// zoneHoldGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneHoldGetResponseEnvelope]
type zoneHoldGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneHoldGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneHoldGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ZoneHoldGetResponseEnvelopeErrors]
type zoneHoldGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldGetResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneHoldGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneHoldGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZoneHoldGetResponseEnvelopeMessages]
type zoneHoldGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneHoldGetResponseEnvelopeSuccess bool

const (
	ZoneHoldGetResponseEnvelopeSuccessTrue ZoneHoldGetResponseEnvelopeSuccess = true
)

type ZoneHoldRemoveParams struct {
	// If `hold_after` is provided, the hold will be temporarily disabled, then
	// automatically re-enabled by the system at the time specified in this
	// RFC3339-formatted timestamp. Otherwise, the hold will be disabled indefinitely.
	HoldAfter param.Field[string] `query:"hold_after"`
}

// URLQuery serializes [ZoneHoldRemoveParams]'s query parameters as `url.Values`.
func (r ZoneHoldRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneHoldRemoveResponseEnvelope struct {
	Result ZoneHoldRemoveResponse             `json:"result"`
	JSON   zoneHoldRemoveResponseEnvelopeJSON `json:"-"`
}

// zoneHoldRemoveResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneHoldRemoveResponseEnvelope]
type zoneHoldRemoveResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldRemoveResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
