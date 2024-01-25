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

// ZoneArgoTieredCachingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneArgoTieredCachingService]
// method instead.
type ZoneArgoTieredCachingService struct {
	Options []option.RequestOption
}

// NewZoneArgoTieredCachingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneArgoTieredCachingService(opts ...option.RequestOption) (r *ZoneArgoTieredCachingService) {
	r = &ZoneArgoTieredCachingService{}
	r.Options = opts
	return
}

// Get Tiered Caching setting
func (r *ZoneArgoTieredCachingService) TieredCachingGetTieredCachingSetting(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates enablement of Tiered Caching
func (r *ZoneArgoTieredCachingService) TieredCachingPatchTieredCachingSetting(ctx context.Context, zoneIdentifier string, body ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParams, opts ...option.RequestOption) (res *ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponse struct {
	Errors   []ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseError   `json:"errors"`
	Messages []ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseMessage `json:"messages"`
	Result   interface{}                                                                `json:"result"`
	// Whether the API call was successful
	Success ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseSuccess `json:"success"`
	JSON    zoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseJSON    `json:"-"`
}

// zoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseJSON contains
// the JSON metadata for the struct
// [ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponse]
type zoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseError struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    zoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseErrorJSON `json:"-"`
}

// zoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseError]
type zoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseMessage struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    zoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseMessageJSON `json:"-"`
}

// zoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseMessage]
type zoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseSuccess bool

const (
	ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseSuccessTrue ZoneArgoTieredCachingTieredCachingGetTieredCachingSettingResponseSuccess = true
)

type ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponse struct {
	Errors   []ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseError   `json:"errors"`
	Messages []ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseMessage `json:"messages"`
	Result   interface{}                                                                  `json:"result"`
	// Whether the API call was successful
	Success ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseSuccess `json:"success"`
	JSON    zoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseJSON    `json:"-"`
}

// zoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseJSON contains
// the JSON metadata for the struct
// [ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponse]
type zoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseError struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    zoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseErrorJSON `json:"-"`
}

// zoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseError]
type zoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseMessage struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    zoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseMessageJSON `json:"-"`
}

// zoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseMessage]
type zoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseSuccess bool

const (
	ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseSuccessTrue ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseSuccess = true
)

type ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParams struct {
	// Enables Tiered Caching.
	Value param.Field[ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValue] `json:"value,required"`
}

func (r ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Caching.
type ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValue string

const (
	ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValueOn  ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValue = "on"
	ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValueOff ZoneArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValue = "off"
)
