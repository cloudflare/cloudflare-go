// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingAlwaysOnlineService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingAlwaysOnlineService]
// method instead.
type SettingAlwaysOnlineService struct {
	Options []option.RequestOption
}

// NewSettingAlwaysOnlineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingAlwaysOnlineService(opts ...option.RequestOption) (r *SettingAlwaysOnlineService) {
	r = &SettingAlwaysOnlineService{}
	r.Options = opts
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
func (r *SettingAlwaysOnlineService) Edit(ctx context.Context, params SettingAlwaysOnlineEditParams, opts ...option.RequestOption) (res *ZoneSettingAlwaysOnline, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAlwaysOnlineEditResponseEnvelope
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
func (r *SettingAlwaysOnlineService) Get(ctx context.Context, query SettingAlwaysOnlineGetParams, opts ...option.RequestOption) (res *ZoneSettingAlwaysOnline, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAlwaysOnlineGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_online", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
type ZoneSettingAlwaysOnline struct {
	// ID of the zone setting.
	ID ZoneSettingAlwaysOnlineID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingAlwaysOnlineValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingAlwaysOnlineJSON `json:"-"`
}

// zoneSettingAlwaysOnlineJSON contains the JSON metadata for the struct
// [ZoneSettingAlwaysOnline]
type zoneSettingAlwaysOnlineJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAlwaysOnlineJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingAlwaysOnline) implementsZonesSettingEditResponse() {}

func (r ZoneSettingAlwaysOnline) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingAlwaysOnlineID string

const (
	ZoneSettingAlwaysOnlineIDAlwaysOnline ZoneSettingAlwaysOnlineID = "always_online"
)

func (r ZoneSettingAlwaysOnlineID) IsKnown() bool {
	switch r {
	case ZoneSettingAlwaysOnlineIDAlwaysOnline:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingAlwaysOnlineValue string

const (
	ZoneSettingAlwaysOnlineValueOn  ZoneSettingAlwaysOnlineValue = "on"
	ZoneSettingAlwaysOnlineValueOff ZoneSettingAlwaysOnlineValue = "off"
)

func (r ZoneSettingAlwaysOnlineValue) IsKnown() bool {
	switch r {
	case ZoneSettingAlwaysOnlineValueOn, ZoneSettingAlwaysOnlineValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAlwaysOnlineEditable bool

const (
	ZoneSettingAlwaysOnlineEditableTrue  ZoneSettingAlwaysOnlineEditable = true
	ZoneSettingAlwaysOnlineEditableFalse ZoneSettingAlwaysOnlineEditable = false
)

func (r ZoneSettingAlwaysOnlineEditable) IsKnown() bool {
	switch r {
	case ZoneSettingAlwaysOnlineEditableTrue, ZoneSettingAlwaysOnlineEditableFalse:
		return true
	}
	return false
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingAlwaysOnlineParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingAlwaysOnlineID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingAlwaysOnlineValue] `json:"value,required"`
}

func (r ZoneSettingAlwaysOnlineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingAlwaysOnlineParam) implementsZonesSettingEditParamsItemUnion() {}

type SettingAlwaysOnlineEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingAlwaysOnlineEditParamsValue] `json:"value,required"`
}

func (r SettingAlwaysOnlineEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingAlwaysOnlineEditParamsValue string

const (
	SettingAlwaysOnlineEditParamsValueOn  SettingAlwaysOnlineEditParamsValue = "on"
	SettingAlwaysOnlineEditParamsValueOff SettingAlwaysOnlineEditParamsValue = "off"
)

func (r SettingAlwaysOnlineEditParamsValue) IsKnown() bool {
	switch r {
	case SettingAlwaysOnlineEditParamsValueOn, SettingAlwaysOnlineEditParamsValueOff:
		return true
	}
	return false
}

type SettingAlwaysOnlineEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result ZoneSettingAlwaysOnline                     `json:"result"`
	JSON   settingAlwaysOnlineEditResponseEnvelopeJSON `json:"-"`
}

// settingAlwaysOnlineEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAlwaysOnlineEditResponseEnvelope]
type settingAlwaysOnlineEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysOnlineEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAlwaysOnlineGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingAlwaysOnlineGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result ZoneSettingAlwaysOnline                    `json:"result"`
	JSON   settingAlwaysOnlineGetResponseEnvelopeJSON `json:"-"`
}

// settingAlwaysOnlineGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAlwaysOnlineGetResponseEnvelope]
type settingAlwaysOnlineGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysOnlineGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
