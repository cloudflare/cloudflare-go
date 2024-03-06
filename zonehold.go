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
func (r *ZoneHoldService) New(ctx context.Context, params ZoneHoldNewParams, opts ...option.RequestOption) (res *ZoneHoldNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneHoldNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/hold", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Stop enforcement of a zone hold on the zone, permanently or temporarily,
// allowing the creation and activation of zones with this zone's hostname.
func (r *ZoneHoldService) Delete(ctx context.Context, params ZoneHoldDeleteParams, opts ...option.RequestOption) (res *ZoneHoldDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneHoldDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/hold", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve whether the zone is subject to a zone hold, and metadata about the
// hold.
func (r *ZoneHoldService) Get(ctx context.Context, query ZoneHoldGetParams, opts ...option.RequestOption) (res *ZoneHoldGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneHoldGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/hold", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZoneHoldNewResponse struct {
	Hold              bool                    `json:"hold"`
	HoldAfter         string                  `json:"hold_after"`
	IncludeSubdomains string                  `json:"include_subdomains"`
	JSON              zoneHoldNewResponseJSON `json:"-"`
}

// zoneHoldNewResponseJSON contains the JSON metadata for the struct
// [ZoneHoldNewResponse]
type zoneHoldNewResponseJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneHoldNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneHoldDeleteResponse struct {
	Hold              bool                       `json:"hold"`
	HoldAfter         string                     `json:"hold_after"`
	IncludeSubdomains string                     `json:"include_subdomains"`
	JSON              zoneHoldDeleteResponseJSON `json:"-"`
}

// zoneHoldDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneHoldDeleteResponse]
type zoneHoldDeleteResponseJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneHoldDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldDeleteResponseJSON) RawJSON() string {
	return r.raw
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

func (r zoneHoldGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneHoldNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// If provided, the zone hold will extend to block any subdomain of the given zone,
	// as well as SSL4SaaS Custom Hostnames. For example, a zone hold on a zone with
	// the hostname 'example.com' and include_subdomains=true will block 'example.com',
	// 'staging.example.com', 'api.staging.example.com', etc.
	IncludeSubdomains param.Field[bool] `query:"include_subdomains"`
}

// URLQuery serializes [ZoneHoldNewParams]'s query parameters as `url.Values`.
func (r ZoneHoldNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneHoldNewResponseEnvelope struct {
	Errors   []ZoneHoldNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneHoldNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZoneHoldNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZoneHoldNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zoneHoldNewResponseEnvelopeJSON    `json:"-"`
}

// zoneHoldNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneHoldNewResponseEnvelope]
type zoneHoldNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneHoldNewResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneHoldNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneHoldNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ZoneHoldNewResponseEnvelopeErrors]
type zoneHoldNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneHoldNewResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneHoldNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneHoldNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZoneHoldNewResponseEnvelopeMessages]
type zoneHoldNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneHoldNewResponseEnvelopeSuccess bool

const (
	ZoneHoldNewResponseEnvelopeSuccessTrue ZoneHoldNewResponseEnvelopeSuccess = true
)

type ZoneHoldDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// If `hold_after` is provided, the hold will be temporarily disabled, then
	// automatically re-enabled by the system at the time specified in this
	// RFC3339-formatted timestamp. Otherwise, the hold will be disabled indefinitely.
	HoldAfter param.Field[string] `query:"hold_after"`
}

// URLQuery serializes [ZoneHoldDeleteParams]'s query parameters as `url.Values`.
func (r ZoneHoldDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneHoldDeleteResponseEnvelope struct {
	Result ZoneHoldDeleteResponse             `json:"result"`
	JSON   zoneHoldDeleteResponseEnvelopeJSON `json:"-"`
}

// zoneHoldDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneHoldDeleteResponseEnvelope]
type zoneHoldDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneHoldGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

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

func (r zoneHoldGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneHoldGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneHoldGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneHoldGetResponseEnvelopeSuccess bool

const (
	ZoneHoldGetResponseEnvelopeSuccessTrue ZoneHoldGetResponseEnvelopeSuccess = true
)
