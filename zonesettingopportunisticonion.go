// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSettingOpportunisticOnionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingOpportunisticOnionService] method instead.
type ZoneSettingOpportunisticOnionService struct {
	Options []option.RequestOption
}

// NewZoneSettingOpportunisticOnionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingOpportunisticOnionService(opts ...option.RequestOption) (r *ZoneSettingOpportunisticOnionService) {
	r = &ZoneSettingOpportunisticOnionService{}
	r.Options = opts
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
func (r *ZoneSettingOpportunisticOnionService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingOpportunisticOnionUpdateParams, opts ...option.RequestOption) (res *ZoneSettingOpportunisticOnionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/opportunistic_onion", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
func (r *ZoneSettingOpportunisticOnionService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingOpportunisticOnionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/opportunistic_onion", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingOpportunisticOnionUpdateResponse struct {
	Errors   []ZoneSettingOpportunisticOnionUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingOpportunisticOnionUpdateResponseMessage `json:"messages"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result ZoneSettingOpportunisticOnionUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                            `json:"success"`
	JSON    zoneSettingOpportunisticOnionUpdateResponseJSON `json:"-"`
}

// zoneSettingOpportunisticOnionUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneSettingOpportunisticOnionUpdateResponse]
type zoneSettingOpportunisticOnionUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticOnionUpdateResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingOpportunisticOnionUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingOpportunisticOnionUpdateResponseErrorJSON contains the JSON metadata
// for the struct [ZoneSettingOpportunisticOnionUpdateResponseError]
type zoneSettingOpportunisticOnionUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticOnionUpdateResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingOpportunisticOnionUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingOpportunisticOnionUpdateResponseMessageJSON contains the JSON
// metadata for the struct [ZoneSettingOpportunisticOnionUpdateResponseMessage]
type zoneSettingOpportunisticOnionUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZoneSettingOpportunisticOnionUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingOpportunisticOnionUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingOpportunisticOnionUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingOpportunisticOnionUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingOpportunisticOnionUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingOpportunisticOnionUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingOpportunisticOnionUpdateResponseResult]
type zoneSettingOpportunisticOnionUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingOpportunisticOnionUpdateResponseResultID string

const (
	ZoneSettingOpportunisticOnionUpdateResponseResultIDOpportunisticOnion ZoneSettingOpportunisticOnionUpdateResponseResultID = "opportunistic_onion"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOpportunisticOnionUpdateResponseResultEditable bool

const (
	ZoneSettingOpportunisticOnionUpdateResponseResultEditableTrue  ZoneSettingOpportunisticOnionUpdateResponseResultEditable = true
	ZoneSettingOpportunisticOnionUpdateResponseResultEditableFalse ZoneSettingOpportunisticOnionUpdateResponseResultEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingOpportunisticOnionUpdateResponseResultValue string

const (
	ZoneSettingOpportunisticOnionUpdateResponseResultValueOn  ZoneSettingOpportunisticOnionUpdateResponseResultValue = "on"
	ZoneSettingOpportunisticOnionUpdateResponseResultValueOff ZoneSettingOpportunisticOnionUpdateResponseResultValue = "off"
)

type ZoneSettingOpportunisticOnionListResponse struct {
	Errors   []ZoneSettingOpportunisticOnionListResponseError   `json:"errors"`
	Messages []ZoneSettingOpportunisticOnionListResponseMessage `json:"messages"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result ZoneSettingOpportunisticOnionListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                          `json:"success"`
	JSON    zoneSettingOpportunisticOnionListResponseJSON `json:"-"`
}

// zoneSettingOpportunisticOnionListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingOpportunisticOnionListResponse]
type zoneSettingOpportunisticOnionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticOnionListResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingOpportunisticOnionListResponseErrorJSON `json:"-"`
}

// zoneSettingOpportunisticOnionListResponseErrorJSON contains the JSON metadata
// for the struct [ZoneSettingOpportunisticOnionListResponseError]
type zoneSettingOpportunisticOnionListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticOnionListResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingOpportunisticOnionListResponseMessageJSON `json:"-"`
}

// zoneSettingOpportunisticOnionListResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingOpportunisticOnionListResponseMessage]
type zoneSettingOpportunisticOnionListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZoneSettingOpportunisticOnionListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingOpportunisticOnionListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingOpportunisticOnionListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingOpportunisticOnionListResponseResultValue `json:"value"`
	JSON  zoneSettingOpportunisticOnionListResponseResultJSON  `json:"-"`
}

// zoneSettingOpportunisticOnionListResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingOpportunisticOnionListResponseResult]
type zoneSettingOpportunisticOnionListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticOnionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingOpportunisticOnionListResponseResultID string

const (
	ZoneSettingOpportunisticOnionListResponseResultIDOpportunisticOnion ZoneSettingOpportunisticOnionListResponseResultID = "opportunistic_onion"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOpportunisticOnionListResponseResultEditable bool

const (
	ZoneSettingOpportunisticOnionListResponseResultEditableTrue  ZoneSettingOpportunisticOnionListResponseResultEditable = true
	ZoneSettingOpportunisticOnionListResponseResultEditableFalse ZoneSettingOpportunisticOnionListResponseResultEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingOpportunisticOnionListResponseResultValue string

const (
	ZoneSettingOpportunisticOnionListResponseResultValueOn  ZoneSettingOpportunisticOnionListResponseResultValue = "on"
	ZoneSettingOpportunisticOnionListResponseResultValueOff ZoneSettingOpportunisticOnionListResponseResultValue = "off"
)

type ZoneSettingOpportunisticOnionUpdateParams struct {
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingOpportunisticOnionUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingOpportunisticOnionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingOpportunisticOnionUpdateParamsValue string

const (
	ZoneSettingOpportunisticOnionUpdateParamsValueOn  ZoneSettingOpportunisticOnionUpdateParamsValue = "on"
	ZoneSettingOpportunisticOnionUpdateParamsValueOff ZoneSettingOpportunisticOnionUpdateParamsValue = "off"
)
