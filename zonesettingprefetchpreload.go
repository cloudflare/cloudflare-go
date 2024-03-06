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

// ZoneSettingPrefetchPreloadService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingPrefetchPreloadService] method instead.
type ZoneSettingPrefetchPreloadService struct {
	Options []option.RequestOption
}

// NewZoneSettingPrefetchPreloadService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingPrefetchPreloadService(opts ...option.RequestOption) (r *ZoneSettingPrefetchPreloadService) {
	r = &ZoneSettingPrefetchPreloadService{}
	r.Options = opts
	return
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
func (r *ZoneSettingPrefetchPreloadService) Edit(ctx context.Context, params ZoneSettingPrefetchPreloadEditParams, opts ...option.RequestOption) (res *ZoneSettingPrefetchPreloadEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingPrefetchPreloadEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/prefetch_preload", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
func (r *ZoneSettingPrefetchPreloadService) Get(ctx context.Context, query ZoneSettingPrefetchPreloadGetParams, opts ...option.RequestOption) (res *ZoneSettingPrefetchPreloadGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingPrefetchPreloadGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/prefetch_preload", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingPrefetchPreloadEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingPrefetchPreloadEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingPrefetchPreloadEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPrefetchPreloadEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingPrefetchPreloadEditResponseJSON `json:"-"`
}

// zoneSettingPrefetchPreloadEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingPrefetchPreloadEditResponse]
type zoneSettingPrefetchPreloadEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingPrefetchPreloadEditResponseID string

const (
	ZoneSettingPrefetchPreloadEditResponseIDPrefetchPreload ZoneSettingPrefetchPreloadEditResponseID = "prefetch_preload"
)

// Current value of the zone setting.
type ZoneSettingPrefetchPreloadEditResponseValue string

const (
	ZoneSettingPrefetchPreloadEditResponseValueOn  ZoneSettingPrefetchPreloadEditResponseValue = "on"
	ZoneSettingPrefetchPreloadEditResponseValueOff ZoneSettingPrefetchPreloadEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPrefetchPreloadEditResponseEditable bool

const (
	ZoneSettingPrefetchPreloadEditResponseEditableTrue  ZoneSettingPrefetchPreloadEditResponseEditable = true
	ZoneSettingPrefetchPreloadEditResponseEditableFalse ZoneSettingPrefetchPreloadEditResponseEditable = false
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingPrefetchPreloadGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingPrefetchPreloadGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingPrefetchPreloadGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPrefetchPreloadGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingPrefetchPreloadGetResponseJSON `json:"-"`
}

// zoneSettingPrefetchPreloadGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingPrefetchPreloadGetResponse]
type zoneSettingPrefetchPreloadGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingPrefetchPreloadGetResponseID string

const (
	ZoneSettingPrefetchPreloadGetResponseIDPrefetchPreload ZoneSettingPrefetchPreloadGetResponseID = "prefetch_preload"
)

// Current value of the zone setting.
type ZoneSettingPrefetchPreloadGetResponseValue string

const (
	ZoneSettingPrefetchPreloadGetResponseValueOn  ZoneSettingPrefetchPreloadGetResponseValue = "on"
	ZoneSettingPrefetchPreloadGetResponseValueOff ZoneSettingPrefetchPreloadGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPrefetchPreloadGetResponseEditable bool

const (
	ZoneSettingPrefetchPreloadGetResponseEditableTrue  ZoneSettingPrefetchPreloadGetResponseEditable = true
	ZoneSettingPrefetchPreloadGetResponseEditableFalse ZoneSettingPrefetchPreloadGetResponseEditable = false
)

type ZoneSettingPrefetchPreloadEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingPrefetchPreloadEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingPrefetchPreloadEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingPrefetchPreloadEditParamsValue string

const (
	ZoneSettingPrefetchPreloadEditParamsValueOn  ZoneSettingPrefetchPreloadEditParamsValue = "on"
	ZoneSettingPrefetchPreloadEditParamsValueOff ZoneSettingPrefetchPreloadEditParamsValue = "off"
)

type ZoneSettingPrefetchPreloadEditResponseEnvelope struct {
	Errors   []ZoneSettingPrefetchPreloadEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingPrefetchPreloadEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result ZoneSettingPrefetchPreloadEditResponse             `json:"result"`
	JSON   zoneSettingPrefetchPreloadEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingPrefetchPreloadEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingPrefetchPreloadEditResponseEnvelope]
type zoneSettingPrefetchPreloadEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPrefetchPreloadEditResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingPrefetchPreloadEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingPrefetchPreloadEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingPrefetchPreloadEditResponseEnvelopeErrors]
type zoneSettingPrefetchPreloadEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPrefetchPreloadEditResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneSettingPrefetchPreloadEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingPrefetchPreloadEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingPrefetchPreloadEditResponseEnvelopeMessages]
type zoneSettingPrefetchPreloadEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPrefetchPreloadGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingPrefetchPreloadGetResponseEnvelope struct {
	Errors   []ZoneSettingPrefetchPreloadGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingPrefetchPreloadGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result ZoneSettingPrefetchPreloadGetResponse             `json:"result"`
	JSON   zoneSettingPrefetchPreloadGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingPrefetchPreloadGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingPrefetchPreloadGetResponseEnvelope]
type zoneSettingPrefetchPreloadGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPrefetchPreloadGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingPrefetchPreloadGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingPrefetchPreloadGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingPrefetchPreloadGetResponseEnvelopeErrors]
type zoneSettingPrefetchPreloadGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPrefetchPreloadGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingPrefetchPreloadGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingPrefetchPreloadGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingPrefetchPreloadGetResponseEnvelopeMessages]
type zoneSettingPrefetchPreloadGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
