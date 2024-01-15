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
func (r *ZoneHoldService) New(ctx context.Context, zoneID string, body ZoneHoldNewParams, opts ...option.RequestOption) (res *ZoneHoldNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve whether the zone is subject to a zone hold, and metadata about the
// hold.
func (r *ZoneHoldService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneHoldGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Stop enforcement of a zone hold on the zone, permanently or temporarily,
// allowing the creation and activation of zones with this zone's hostname.
func (r *ZoneHoldService) Delete(ctx context.Context, zoneID string, body ZoneHoldDeleteParams, opts ...option.RequestOption) (res *ZoneHoldDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type ZoneHoldNewResponse struct {
	Errors   []ZoneHoldNewResponseError   `json:"errors"`
	Messages []ZoneHoldNewResponseMessage `json:"messages"`
	Result   ZoneHoldNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneHoldNewResponseSuccess `json:"success"`
	JSON    zoneHoldNewResponseJSON    `json:"-"`
}

// zoneHoldNewResponseJSON contains the JSON metadata for the struct
// [ZoneHoldNewResponse]
type zoneHoldNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldNewResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    zoneHoldNewResponseErrorJSON `json:"-"`
}

// zoneHoldNewResponseErrorJSON contains the JSON metadata for the struct
// [ZoneHoldNewResponseError]
type zoneHoldNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldNewResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    zoneHoldNewResponseMessageJSON `json:"-"`
}

// zoneHoldNewResponseMessageJSON contains the JSON metadata for the struct
// [ZoneHoldNewResponseMessage]
type zoneHoldNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldNewResponseResult struct {
	Hold              bool                          `json:"hold"`
	HoldAfter         string                        `json:"hold_after"`
	IncludeSubdomains string                        `json:"include_subdomains"`
	JSON              zoneHoldNewResponseResultJSON `json:"-"`
}

// zoneHoldNewResponseResultJSON contains the JSON metadata for the struct
// [ZoneHoldNewResponseResult]
type zoneHoldNewResponseResultJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneHoldNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneHoldNewResponseSuccess bool

const (
	ZoneHoldNewResponseSuccessTrue ZoneHoldNewResponseSuccess = true
)

type ZoneHoldGetResponse struct {
	Errors   []ZoneHoldGetResponseError   `json:"errors"`
	Messages []ZoneHoldGetResponseMessage `json:"messages"`
	Result   ZoneHoldGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneHoldGetResponseSuccess `json:"success"`
	JSON    zoneHoldGetResponseJSON    `json:"-"`
}

// zoneHoldGetResponseJSON contains the JSON metadata for the struct
// [ZoneHoldGetResponse]
type zoneHoldGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldGetResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    zoneHoldGetResponseErrorJSON `json:"-"`
}

// zoneHoldGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneHoldGetResponseError]
type zoneHoldGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldGetResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    zoneHoldGetResponseMessageJSON `json:"-"`
}

// zoneHoldGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneHoldGetResponseMessage]
type zoneHoldGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldGetResponseResult struct {
	Hold              bool                          `json:"hold"`
	HoldAfter         string                        `json:"hold_after"`
	IncludeSubdomains string                        `json:"include_subdomains"`
	JSON              zoneHoldGetResponseResultJSON `json:"-"`
}

// zoneHoldGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneHoldGetResponseResult]
type zoneHoldGetResponseResultJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneHoldGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneHoldGetResponseSuccess bool

const (
	ZoneHoldGetResponseSuccessTrue ZoneHoldGetResponseSuccess = true
)

type ZoneHoldDeleteResponse struct {
	Result ZoneHoldDeleteResponseResult `json:"result"`
	JSON   zoneHoldDeleteResponseJSON   `json:"-"`
}

// zoneHoldDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneHoldDeleteResponse]
type zoneHoldDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldDeleteResponseResult struct {
	Hold              bool                             `json:"hold"`
	HoldAfter         string                           `json:"hold_after"`
	IncludeSubdomains string                           `json:"include_subdomains"`
	JSON              zoneHoldDeleteResponseResultJSON `json:"-"`
}

// zoneHoldDeleteResponseResultJSON contains the JSON metadata for the struct
// [ZoneHoldDeleteResponseResult]
type zoneHoldDeleteResponseResultJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneHoldDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHoldNewParams struct {
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

type ZoneHoldDeleteParams struct {
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
