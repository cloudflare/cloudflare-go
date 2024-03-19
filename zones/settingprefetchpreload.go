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
func (r *SettingPrefetchPreloadService) Edit(ctx context.Context, params SettingPrefetchPreloadEditParams, opts ...option.RequestOption) (res *ZonesPrefetchPreload, err error) {
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
func (r *SettingPrefetchPreloadService) Get(ctx context.Context, query SettingPrefetchPreloadGetParams, opts ...option.RequestOption) (res *ZonesPrefetchPreload, err error) {
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
type ZonesPrefetchPreload struct {
	// ID of the zone setting.
	ID ZonesPrefetchPreloadID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesPrefetchPreloadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesPrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesPrefetchPreloadJSON `json:"-"`
}

// zonesPrefetchPreloadJSON contains the JSON metadata for the struct
// [ZonesPrefetchPreload]
type zonesPrefetchPreloadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesPrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesPrefetchPreloadJSON) RawJSON() string {
	return r.raw
}

func (r ZonesPrefetchPreload) implementsZonesSettingEditResponse() {}

func (r ZonesPrefetchPreload) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesPrefetchPreloadID string

const (
	ZonesPrefetchPreloadIDPrefetchPreload ZonesPrefetchPreloadID = "prefetch_preload"
)

// Current value of the zone setting.
type ZonesPrefetchPreloadValue string

const (
	ZonesPrefetchPreloadValueOn  ZonesPrefetchPreloadValue = "on"
	ZonesPrefetchPreloadValueOff ZonesPrefetchPreloadValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesPrefetchPreloadEditable bool

const (
	ZonesPrefetchPreloadEditableTrue  ZonesPrefetchPreloadEditable = true
	ZonesPrefetchPreloadEditableFalse ZonesPrefetchPreloadEditable = false
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZonesPrefetchPreloadParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesPrefetchPreloadID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesPrefetchPreloadValue] `json:"value,required"`
}

func (r ZonesPrefetchPreloadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesPrefetchPreloadParam) implementsZonesSettingEditParamsItem() {}

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

type SettingPrefetchPreloadEditResponseEnvelope struct {
	Errors   []SettingPrefetchPreloadEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingPrefetchPreloadEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result ZonesPrefetchPreload                           `json:"result"`
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

type SettingPrefetchPreloadEditResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingPrefetchPreloadEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingPrefetchPreloadEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingPrefetchPreloadEditResponseEnvelopeErrors]
type settingPrefetchPreloadEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingPrefetchPreloadEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingPrefetchPreloadEditResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingPrefetchPreloadEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingPrefetchPreloadEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingPrefetchPreloadEditResponseEnvelopeMessages]
type settingPrefetchPreloadEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingPrefetchPreloadEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingPrefetchPreloadGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingPrefetchPreloadGetResponseEnvelope struct {
	Errors   []SettingPrefetchPreloadGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingPrefetchPreloadGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result ZonesPrefetchPreload                          `json:"result"`
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

type SettingPrefetchPreloadGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingPrefetchPreloadGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingPrefetchPreloadGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingPrefetchPreloadGetResponseEnvelopeErrors]
type settingPrefetchPreloadGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingPrefetchPreloadGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingPrefetchPreloadGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingPrefetchPreloadGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingPrefetchPreloadGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingPrefetchPreloadGetResponseEnvelopeMessages]
type settingPrefetchPreloadGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingPrefetchPreloadGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
