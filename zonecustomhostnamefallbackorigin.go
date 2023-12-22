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
func (r *ZoneCustomHostnameFallbackOriginService) Delete(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *FallbackOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get Fallback Origin for Custom Hostnames
func (r *ZoneCustomHostnameFallbackOriginService) CustomHostnameFallbackOriginForAZoneGetFallbackOriginForCustomHostnames(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *FallbackOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Fallback Origin for Custom Hostnames
func (r *ZoneCustomHostnameFallbackOriginService) CustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnames(ctx context.Context, zoneIdentifier string, body ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesParams, opts ...option.RequestOption) (res *FallbackOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type FallbackOriginResponse struct {
	Errors   []FallbackOriginResponseError   `json:"errors"`
	Messages []FallbackOriginResponseMessage `json:"messages"`
	Result   interface{}                     `json:"result"`
	// Whether the API call was successful
	Success FallbackOriginResponseSuccess `json:"success"`
	JSON    fallbackOriginResponseJSON    `json:"-"`
}

// fallbackOriginResponseJSON contains the JSON metadata for the struct
// [FallbackOriginResponse]
type fallbackOriginResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FallbackOriginResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    fallbackOriginResponseErrorJSON `json:"-"`
}

// fallbackOriginResponseErrorJSON contains the JSON metadata for the struct
// [FallbackOriginResponseError]
type fallbackOriginResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FallbackOriginResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    fallbackOriginResponseMessageJSON `json:"-"`
}

// fallbackOriginResponseMessageJSON contains the JSON metadata for the struct
// [FallbackOriginResponseMessage]
type fallbackOriginResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FallbackOriginResponseSuccess bool

const (
	FallbackOriginResponseSuccessTrue FallbackOriginResponseSuccess = true
)

type ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesParams struct {
	// Your origin hostname that requests to your custom hostnames will be sent to.
	Origin param.Field[string] `json:"origin,required"`
}

func (r ZoneCustomHostnameFallbackOriginCustomHostnameFallbackOriginForAZoneUpdateFallbackOriginForCustomHostnamesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
