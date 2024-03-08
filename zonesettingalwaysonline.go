// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *ZoneSettingAlwaysOnlineService) Edit(ctx context.Context, params ZoneSettingAlwaysOnlineEditParams, opts ...option.RequestOption) (res *ZoneSettingAlwaysOnlineEditResponse, err error) {
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
func (r *ZoneSettingAlwaysOnlineService) Get(ctx context.Context, query ZoneSettingAlwaysOnlineGetParams, opts ...option.RequestOption) (res *ZoneSettingAlwaysOnlineGetResponse, err error) {
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
type ZoneSettingAlwaysOnlineEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingAlwaysOnlineEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingAlwaysOnlineEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAlwaysOnlineEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingAlwaysOnlineEditResponseJSON `json:"-"`
}

// zoneSettingAlwaysOnlineEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingAlwaysOnlineEditResponse]
type zoneSettingAlwaysOnlineEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAlwaysOnlineEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingAlwaysOnlineEditResponseID string

const (
	ZoneSettingAlwaysOnlineEditResponseIDAlwaysOnline ZoneSettingAlwaysOnlineEditResponseID = "always_online"
)

// Current value of the zone setting.
type ZoneSettingAlwaysOnlineEditResponseValue string

const (
	ZoneSettingAlwaysOnlineEditResponseValueOn  ZoneSettingAlwaysOnlineEditResponseValue = "on"
	ZoneSettingAlwaysOnlineEditResponseValueOff ZoneSettingAlwaysOnlineEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAlwaysOnlineEditResponseEditable bool

const (
	ZoneSettingAlwaysOnlineEditResponseEditableTrue  ZoneSettingAlwaysOnlineEditResponseEditable = true
	ZoneSettingAlwaysOnlineEditResponseEditableFalse ZoneSettingAlwaysOnlineEditResponseEditable = false
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingAlwaysOnlineGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingAlwaysOnlineGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingAlwaysOnlineGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAlwaysOnlineGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingAlwaysOnlineGetResponseJSON `json:"-"`
}

// zoneSettingAlwaysOnlineGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingAlwaysOnlineGetResponse]
type zoneSettingAlwaysOnlineGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAlwaysOnlineGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingAlwaysOnlineGetResponseID string

const (
	ZoneSettingAlwaysOnlineGetResponseIDAlwaysOnline ZoneSettingAlwaysOnlineGetResponseID = "always_online"
)

// Current value of the zone setting.
type ZoneSettingAlwaysOnlineGetResponseValue string

const (
	ZoneSettingAlwaysOnlineGetResponseValueOn  ZoneSettingAlwaysOnlineGetResponseValue = "on"
	ZoneSettingAlwaysOnlineGetResponseValueOff ZoneSettingAlwaysOnlineGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAlwaysOnlineGetResponseEditable bool

const (
	ZoneSettingAlwaysOnlineGetResponseEditableTrue  ZoneSettingAlwaysOnlineGetResponseEditable = true
	ZoneSettingAlwaysOnlineGetResponseEditableFalse ZoneSettingAlwaysOnlineGetResponseEditable = false
)

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
	Result ZoneSettingAlwaysOnlineEditResponse             `json:"result"`
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

func (r zoneSettingAlwaysOnlineEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingAlwaysOnlineEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingAlwaysOnlineEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result ZoneSettingAlwaysOnlineGetResponse             `json:"result"`
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

func (r zoneSettingAlwaysOnlineGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingAlwaysOnlineGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingAlwaysOnlineGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
