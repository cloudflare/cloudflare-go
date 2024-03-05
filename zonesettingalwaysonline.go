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

// ZoneSettingAlwaysOnlineService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingAlwaysOnlineService] method instead.
type ZoneSettingAlwaysOnlineService struct {
	Options []option.RequestOption
}

// NewZoneSettingAlwaysOnlineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingAlwaysOnlineService(opts ...option.RequestOption) (r *ZoneSettingAlwaysOnlineService) {
	r = &ZoneSettingAlwaysOnlineService{}
	r.Options = opts
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
func (r *ZoneSettingAlwaysOnlineService) Edit(ctx context.Context, params ZoneSettingAlwaysOnlineEditParams, opts ...option.RequestOption) (res *ZonesAlwaysOnline, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingAlwaysOnlineEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_online", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
func (r *ZoneSettingAlwaysOnlineService) Get(ctx context.Context, query ZoneSettingAlwaysOnlineGetParams, opts ...option.RequestOption) (res *ZonesAlwaysOnline, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingAlwaysOnlineGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_online", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID ZonesAlwaysOnlineID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesAlwaysOnlineValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesAlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time             `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesAlwaysOnlineJSON `json:"-"`
}

// zonesAlwaysOnlineJSON contains the JSON metadata for the struct
// [ZonesAlwaysOnline]
type zonesAlwaysOnlineJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesAlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesAlwaysOnline) implementsZoneSettingEditResponse() {}

func (r ZonesAlwaysOnline) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesAlwaysOnlineID string

const (
	ZonesAlwaysOnlineIDAlwaysOnline ZonesAlwaysOnlineID = "always_online"
)

// Current value of the zone setting.
type ZonesAlwaysOnlineValue string

const (
	ZonesAlwaysOnlineValueOn  ZonesAlwaysOnlineValue = "on"
	ZonesAlwaysOnlineValueOff ZonesAlwaysOnlineValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesAlwaysOnlineEditable bool

const (
	ZonesAlwaysOnlineEditableTrue  ZonesAlwaysOnlineEditable = true
	ZonesAlwaysOnlineEditableFalse ZonesAlwaysOnlineEditable = false
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZonesAlwaysOnlineParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesAlwaysOnlineID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesAlwaysOnlineValue] `json:"value,required"`
}

func (r ZonesAlwaysOnlineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesAlwaysOnlineParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingAlwaysOnlineEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingAlwaysOnlineEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingAlwaysOnlineEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingAlwaysOnlineEditParamsValue string

const (
	ZoneSettingAlwaysOnlineEditParamsValueOn  ZoneSettingAlwaysOnlineEditParamsValue = "on"
	ZoneSettingAlwaysOnlineEditParamsValueOff ZoneSettingAlwaysOnlineEditParamsValue = "off"
)

type ZoneSettingAlwaysOnlineEditResponseEnvelope struct {
	Errors   []ZoneSettingAlwaysOnlineEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingAlwaysOnlineEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result ZonesAlwaysOnline                               `json:"result"`
	JSON   zoneSettingAlwaysOnlineEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingAlwaysOnlineEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingAlwaysOnlineEditResponseEnvelope]
type zoneSettingAlwaysOnlineEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysOnlineEditResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneSettingAlwaysOnlineEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingAlwaysOnlineEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingAlwaysOnlineEditResponseEnvelopeErrors]
type zoneSettingAlwaysOnlineEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysOnlineEditResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingAlwaysOnlineEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingAlwaysOnlineEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingAlwaysOnlineEditResponseEnvelopeMessages]
type zoneSettingAlwaysOnlineEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysOnlineGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingAlwaysOnlineGetResponseEnvelope struct {
	Errors   []ZoneSettingAlwaysOnlineGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingAlwaysOnlineGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result ZonesAlwaysOnline                              `json:"result"`
	JSON   zoneSettingAlwaysOnlineGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingAlwaysOnlineGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingAlwaysOnlineGetResponseEnvelope]
type zoneSettingAlwaysOnlineGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysOnlineGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingAlwaysOnlineGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingAlwaysOnlineGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingAlwaysOnlineGetResponseEnvelopeErrors]
type zoneSettingAlwaysOnlineGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysOnlineGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingAlwaysOnlineGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingAlwaysOnlineGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingAlwaysOnlineGetResponseEnvelopeMessages]
type zoneSettingAlwaysOnlineGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
