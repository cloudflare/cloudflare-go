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

// SettingPrefetchPreloadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingPrefetchPreloadService]
// method instead.
type SettingPrefetchPreloadService struct {
	Options []option.RequestOption
}

// NewSettingPrefetchPreloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingPrefetchPreloadService(opts ...option.RequestOption) (r *SettingPrefetchPreloadService) {
	r = &SettingPrefetchPreloadService{}
	r.Options = opts
	return
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
func (r *SettingPrefetchPreloadService) Edit(ctx context.Context, params SettingPrefetchPreloadEditParams, opts ...option.RequestOption) (res *ZoneSettingPrefetchPreload, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPrefetchPreloadEditResponseEnvelope
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
func (r *SettingPrefetchPreloadService) Get(ctx context.Context, query SettingPrefetchPreloadGetParams, opts ...option.RequestOption) (res *ZoneSettingPrefetchPreload, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPrefetchPreloadGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/prefetch_preload", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingPrefetchPreload struct {
	// ID of the zone setting.
	ID ZoneSettingPrefetchPreloadID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingPrefetchPreloadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingPrefetchPreloadJSON `json:"-"`
}

// zoneSettingPrefetchPreloadJSON contains the JSON metadata for the struct
// [ZoneSettingPrefetchPreload]
type zoneSettingPrefetchPreloadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingPrefetchPreloadJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingPrefetchPreloadID string

const (
	ZoneSettingPrefetchPreloadIDPrefetchPreload ZoneSettingPrefetchPreloadID = "prefetch_preload"
)

func (r ZoneSettingPrefetchPreloadID) IsKnown() bool {
	switch r {
	case ZoneSettingPrefetchPreloadIDPrefetchPreload:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingPrefetchPreloadValue string

const (
	ZoneSettingPrefetchPreloadValueOn  ZoneSettingPrefetchPreloadValue = "on"
	ZoneSettingPrefetchPreloadValueOff ZoneSettingPrefetchPreloadValue = "off"
)

func (r ZoneSettingPrefetchPreloadValue) IsKnown() bool {
	switch r {
	case ZoneSettingPrefetchPreloadValueOn, ZoneSettingPrefetchPreloadValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPrefetchPreloadEditable bool

const (
	ZoneSettingPrefetchPreloadEditableTrue  ZoneSettingPrefetchPreloadEditable = true
	ZoneSettingPrefetchPreloadEditableFalse ZoneSettingPrefetchPreloadEditable = false
)

func (r ZoneSettingPrefetchPreloadEditable) IsKnown() bool {
	switch r {
	case ZoneSettingPrefetchPreloadEditableTrue, ZoneSettingPrefetchPreloadEditableFalse:
		return true
	}
	return false
}

type SettingPrefetchPreloadEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingPrefetchPreloadEditParamsValue] `json:"value,required"`
}

func (r SettingPrefetchPreloadEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingPrefetchPreloadEditParamsValue string

const (
	SettingPrefetchPreloadEditParamsValueOn  SettingPrefetchPreloadEditParamsValue = "on"
	SettingPrefetchPreloadEditParamsValueOff SettingPrefetchPreloadEditParamsValue = "off"
)

func (r SettingPrefetchPreloadEditParamsValue) IsKnown() bool {
	switch r {
	case SettingPrefetchPreloadEditParamsValueOn, SettingPrefetchPreloadEditParamsValueOff:
		return true
	}
	return false
}

type SettingPrefetchPreloadEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result ZoneSettingPrefetchPreload                     `json:"result"`
	JSON   settingPrefetchPreloadEditResponseEnvelopeJSON `json:"-"`
}

// settingPrefetchPreloadEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingPrefetchPreloadEditResponseEnvelope]
type settingPrefetchPreloadEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingPrefetchPreloadEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingPrefetchPreloadGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingPrefetchPreloadGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result ZoneSettingPrefetchPreload                    `json:"result"`
	JSON   settingPrefetchPreloadGetResponseEnvelopeJSON `json:"-"`
}

// settingPrefetchPreloadGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingPrefetchPreloadGetResponseEnvelope]
type settingPrefetchPreloadGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingPrefetchPreloadGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
