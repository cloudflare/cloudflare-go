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
func (r *SettingPrefetchPreloadService) Edit(ctx context.Context, params SettingPrefetchPreloadEditParams, opts ...option.RequestOption) (res *PrefetchPreload, err error) {
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
func (r *SettingPrefetchPreloadService) Get(ctx context.Context, query SettingPrefetchPreloadGetParams, opts ...option.RequestOption) (res *PrefetchPreload, err error) {
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
type PrefetchPreload struct {
	// ID of the zone setting.
	ID PrefetchPreloadID `json:"id,required"`
	// Current value of the zone setting.
	Value PrefetchPreloadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable PrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       prefetchPreloadJSON `json:"-"`
}

// prefetchPreloadJSON contains the JSON metadata for the struct [PrefetchPreload]
type prefetchPreloadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefetchPreloadJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type PrefetchPreloadID string

const (
	PrefetchPreloadIDPrefetchPreload PrefetchPreloadID = "prefetch_preload"
)

func (r PrefetchPreloadID) IsKnown() bool {
	switch r {
	case PrefetchPreloadIDPrefetchPreload:
		return true
	}
	return false
}

// Current value of the zone setting.
type PrefetchPreloadValue string

const (
	PrefetchPreloadValueOn  PrefetchPreloadValue = "on"
	PrefetchPreloadValueOff PrefetchPreloadValue = "off"
)

func (r PrefetchPreloadValue) IsKnown() bool {
	switch r {
	case PrefetchPreloadValueOn, PrefetchPreloadValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type PrefetchPreloadEditable bool

const (
	PrefetchPreloadEditableTrue  PrefetchPreloadEditable = true
	PrefetchPreloadEditableFalse PrefetchPreloadEditable = false
)

func (r PrefetchPreloadEditable) IsKnown() bool {
	switch r {
	case PrefetchPreloadEditableTrue, PrefetchPreloadEditableFalse:
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result PrefetchPreload                                `json:"result"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result PrefetchPreload                               `json:"result"`
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
