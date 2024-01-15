// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneCustomHostnameFallbackOriginService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneCustomHostnameFallbackOriginService] method instead.
type ZoneCustomHostnameFallbackOriginService struct {
	Options []option.RequestOption
}

// NewZoneCustomHostnameFallbackOriginService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneCustomHostnameFallbackOriginService(opts ...option.RequestOption) (r *ZoneCustomHostnameFallbackOriginService) {
	r = &ZoneCustomHostnameFallbackOriginService{}
	r.Options = opts
	return
}

// Delete Fallback Origin for Custom Hostnames
func (r *ZoneCustomHostnameFallbackOriginService) Delete(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneCustomHostnameFallbackOriginDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get Fallback Origin for Custom Hostnames
func (r *ZoneCustomHostnameFallbackOriginService) CustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnames(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Fallback Origin for Custom Hostnames
func (r *ZoneCustomHostnameFallbackOriginService) CustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnames(ctx context.Context, zoneIdentifier string, body ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesParams, opts ...option.RequestOption) (res *ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneCustomHostnameFallbackOriginDeleteResponse struct {
	Errors   []ZoneCustomHostnameFallbackOriginDeleteResponseError   `json:"errors"`
	Messages []ZoneCustomHostnameFallbackOriginDeleteResponseMessage `json:"messages"`
	Result   interface{}                                             `json:"result"`
	// Whether the API call was successful
	Success ZoneCustomHostnameFallbackOriginDeleteResponseSuccess `json:"success"`
	JSON    zoneCustomHostnameFallbackOriginDeleteResponseJSON    `json:"-"`
}

// zoneCustomHostnameFallbackOriginDeleteResponseJSON contains the JSON metadata
// for the struct [ZoneCustomHostnameFallbackOriginDeleteResponse]
type zoneCustomHostnameFallbackOriginDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameFallbackOriginDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameFallbackOriginDeleteResponseError struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneCustomHostnameFallbackOriginDeleteResponseErrorJSON `json:"-"`
}

// zoneCustomHostnameFallbackOriginDeleteResponseErrorJSON contains the JSON
// metadata for the struct [ZoneCustomHostnameFallbackOriginDeleteResponseError]
type zoneCustomHostnameFallbackOriginDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameFallbackOriginDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameFallbackOriginDeleteResponseMessage struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneCustomHostnameFallbackOriginDeleteResponseMessageJSON `json:"-"`
}

// zoneCustomHostnameFallbackOriginDeleteResponseMessageJSON contains the JSON
// metadata for the struct [ZoneCustomHostnameFallbackOriginDeleteResponseMessage]
type zoneCustomHostnameFallbackOriginDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameFallbackOriginDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCustomHostnameFallbackOriginDeleteResponseSuccess bool

const (
	ZoneCustomHostnameFallbackOriginDeleteResponseSuccessTrue ZoneCustomHostnameFallbackOriginDeleteResponseSuccess = true
)

type ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponse struct {
	Errors   []ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseError   `json:"errors"`
	Messages []ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseMessage `json:"messages"`
	Result   interface{}                                                                                                              `json:"result"`
	// Whether the API call was successful
	Success ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseSuccess `json:"success"`
	JSON    zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseJSON    `json:"-"`
}

// zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponse]
type zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseError struct {
	Code    int64                                                                                                                    `json:"code,required"`
	Message string                                                                                                                   `json:"message,required"`
	JSON    zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseErrorJSON `json:"-"`
}

// zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseError]
type zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseMessage struct {
	Code    int64                                                                                                                      `json:"code,required"`
	Message string                                                                                                                     `json:"message,required"`
	JSON    zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseMessageJSON `json:"-"`
}

// zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseMessage]
type zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseSuccess bool

const (
	ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseSuccessTrue ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnamesResponseSuccess = true
)

type ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponse struct {
	Errors   []ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseError   `json:"errors"`
	Messages []ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseMessage `json:"messages"`
	Result   interface{}                                                                                                                 `json:"result"`
	// Whether the API call was successful
	Success ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseSuccess `json:"success"`
	JSON    zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseJSON    `json:"-"`
}

// zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponse]
type zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseError struct {
	Code    int64                                                                                                                       `json:"code,required"`
	Message string                                                                                                                      `json:"message,required"`
	JSON    zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseErrorJSON `json:"-"`
}

// zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseError]
type zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseMessage struct {
	Code    int64                                                                                                                         `json:"code,required"`
	Message string                                                                                                                        `json:"message,required"`
	JSON    zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseMessageJSON `json:"-"`
}

// zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseMessage]
type zoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseSuccess bool

const (
	ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseSuccessTrue ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesResponseSuccess = true
)

type ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesParams struct {
	// Your origin hostname that requests to your custom hostnames will be sent to.
	Origin param.Field[string] `json:"origin,required"`
}

func (r ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
