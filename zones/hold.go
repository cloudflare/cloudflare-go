// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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

// Enforce a zone hold on the zone, blocking the creation and activation of zones
// with this zone's hostname.
func (r *HoldService) New(ctx context.Context, params HoldNewParams, opts ...option.RequestOption) (res *HoldNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HoldNewResponseEnvelope
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
func (r *HoldService) Delete(ctx context.Context, params HoldDeleteParams, opts ...option.RequestOption) (res *HoldDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HoldDeleteResponseEnvelope
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
func (r *HoldService) Get(ctx context.Context, query HoldGetParams, opts ...option.RequestOption) (res *HoldGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HoldGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/hold", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HoldNewResponse struct {
	Hold              bool                `json:"hold"`
	HoldAfter         string              `json:"hold_after"`
	IncludeSubdomains string              `json:"include_subdomains"`
	JSON              holdNewResponseJSON `json:"-"`
}

// holdNewResponseJSON contains the JSON metadata for the struct [HoldNewResponse]
type holdNewResponseJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *HoldNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r holdNewResponseJSON) RawJSON() string {
	return r.raw
}

type HoldDeleteResponse struct {
	Hold              bool                   `json:"hold"`
	HoldAfter         string                 `json:"hold_after"`
	IncludeSubdomains string                 `json:"include_subdomains"`
	JSON              holdDeleteResponseJSON `json:"-"`
}

// holdDeleteResponseJSON contains the JSON metadata for the struct
// [HoldDeleteResponse]
type holdDeleteResponseJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *HoldDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r holdDeleteResponseJSON) RawJSON() string {
	return r.raw
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

func (r holdGetResponseJSON) RawJSON() string {
	return r.raw
}

type HoldNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// If provided, the zone hold will extend to block any subdomain of the given zone,
	// as well as SSL4SaaS Custom Hostnames. For example, a zone hold on a zone with
	// the hostname 'example.com' and include_subdomains=true will block 'example.com',
	// 'staging.example.com', 'api.staging.example.com', etc.
	IncludeSubdomains param.Field[bool] `query:"include_subdomains"`
}

// URLQuery serializes [HoldNewParams]'s query parameters as `url.Values`.
func (r HoldNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HoldNewResponseEnvelope struct {
	Errors   []HoldNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HoldNewResponseEnvelopeMessages `json:"messages,required"`
	Result   HoldNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HoldNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    holdNewResponseEnvelopeJSON    `json:"-"`
}

// holdNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [HoldNewResponseEnvelope]
type holdNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r holdNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HoldNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    holdNewResponseEnvelopeErrorsJSON `json:"-"`
}

// holdNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [HoldNewResponseEnvelopeErrors]
type holdNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r holdNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HoldNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    holdNewResponseEnvelopeMessagesJSON `json:"-"`
}

// holdNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [HoldNewResponseEnvelopeMessages]
type holdNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r holdNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HoldNewResponseEnvelopeSuccess bool

const (
	HoldNewResponseEnvelopeSuccessTrue HoldNewResponseEnvelopeSuccess = true
)

func (r HoldNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HoldNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HoldDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// If `hold_after` is provided, the hold will be temporarily disabled, then
	// automatically re-enabled by the system at the time specified in this
	// RFC3339-formatted timestamp. Otherwise, the hold will be disabled indefinitely.
	HoldAfter param.Field[string] `query:"hold_after"`
}

// URLQuery serializes [HoldDeleteParams]'s query parameters as `url.Values`.
func (r HoldDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HoldDeleteResponseEnvelope struct {
	Result HoldDeleteResponse             `json:"result"`
	JSON   holdDeleteResponseEnvelopeJSON `json:"-"`
}

// holdDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [HoldDeleteResponseEnvelope]
type holdDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r holdDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HoldGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HoldGetResponseEnvelope struct {
	Errors   []HoldGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HoldGetResponseEnvelopeMessages `json:"messages,required"`
	Result   HoldGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HoldGetResponseEnvelopeSuccess `json:"success,required"`
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

func (r holdGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r holdGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r holdGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HoldGetResponseEnvelopeSuccess bool

const (
	HoldGetResponseEnvelopeSuccessTrue HoldGetResponseEnvelopeSuccess = true
)

func (r HoldGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HoldGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
