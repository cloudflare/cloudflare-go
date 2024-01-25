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

// ZoneArgoSmartRoutingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneArgoSmartRoutingService]
// method instead.
type ZoneArgoSmartRoutingService struct {
	Options []option.RequestOption
}

// NewZoneArgoSmartRoutingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneArgoSmartRoutingService(opts ...option.RequestOption) (r *ZoneArgoSmartRoutingService) {
	r = &ZoneArgoSmartRoutingService{}
	r.Options = opts
	return
}

// Get Argo Smart Routing setting
func (r *ZoneArgoSmartRoutingService) ArgoSmartRoutingGetArgoSmartRoutingSetting(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/argo/smart_routing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates enablement of Argo Smart Routing.
func (r *ZoneArgoSmartRoutingService) ArgoSmartRoutingPatchArgoSmartRoutingSetting(ctx context.Context, zoneIdentifier string, body ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParams, opts ...option.RequestOption) (res *ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/argo/smart_routing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponse struct {
	Errors   []ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseError   `json:"errors"`
	Messages []ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseMessage `json:"messages"`
	Result   interface{}                                                                     `json:"result"`
	// Whether the API call was successful
	Success ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseSuccess `json:"success"`
	JSON    zoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseJSON    `json:"-"`
}

// zoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseJSON
// contains the JSON metadata for the struct
// [ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponse]
type zoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseError struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    zoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseErrorJSON `json:"-"`
}

// zoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseError]
type zoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseMessage struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    zoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseMessageJSON `json:"-"`
}

// zoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseMessage]
type zoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseSuccess bool

const (
	ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseSuccessTrue ZoneArgoSmartRoutingArgoSmartRoutingGetArgoSmartRoutingSettingResponseSuccess = true
)

type ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponse struct {
	Errors   []ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseError   `json:"errors"`
	Messages []ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseMessage `json:"messages"`
	Result   interface{}                                                                       `json:"result"`
	// Whether the API call was successful
	Success ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseSuccess `json:"success"`
	JSON    zoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseJSON    `json:"-"`
}

// zoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseJSON
// contains the JSON metadata for the struct
// [ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponse]
type zoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseError struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    zoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseErrorJSON `json:"-"`
}

// zoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseError]
type zoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseMessage struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    zoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseMessageJSON `json:"-"`
}

// zoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseMessage]
type zoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseSuccess bool

const (
	ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseSuccessTrue ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingResponseSuccess = true
)

type ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParams struct {
	// Enables Argo Smart Routing.
	Value param.Field[ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParamsValue] `json:"value,required"`
}

func (r ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Argo Smart Routing.
type ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParamsValue string

const (
	ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParamsValueOn  ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParamsValue = "on"
	ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParamsValueOff ZoneArgoSmartRoutingArgoSmartRoutingPatchArgoSmartRoutingSettingParamsValue = "off"
)
