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

// SettingBrotliService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingBrotliService] method
// instead.
type SettingBrotliService struct {
	Options []option.RequestOption
}

// NewSettingBrotliService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingBrotliService(opts ...option.RequestOption) (r *SettingBrotliService) {
	r = &SettingBrotliService{}
	r.Options = opts
	return
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
func (r *SettingBrotliService) Edit(ctx context.Context, params SettingBrotliEditParams, opts ...option.RequestOption) (res *ZoneSettingBrotli, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrotliEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/brotli", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
func (r *SettingBrotliService) Get(ctx context.Context, query SettingBrotliGetParams, opts ...option.RequestOption) (res *ZoneSettingBrotli, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrotliGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/brotli", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingBrotli struct {
	// ID of the zone setting.
	ID ZoneSettingBrotliID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingBrotliValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrotliEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingBrotliJSON `json:"-"`
}

// zoneSettingBrotliJSON contains the JSON metadata for the struct
// [ZoneSettingBrotli]
type zoneSettingBrotliJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotli) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBrotliJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingBrotli) implementsZonesSettingEditResponse() {}

func (r ZoneSettingBrotli) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingBrotliID string

const (
	ZoneSettingBrotliIDBrotli ZoneSettingBrotliID = "brotli"
)

func (r ZoneSettingBrotliID) IsKnown() bool {
	switch r {
	case ZoneSettingBrotliIDBrotli:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingBrotliValue string

const (
	ZoneSettingBrotliValueOff ZoneSettingBrotliValue = "off"
	ZoneSettingBrotliValueOn  ZoneSettingBrotliValue = "on"
)

func (r ZoneSettingBrotliValue) IsKnown() bool {
	switch r {
	case ZoneSettingBrotliValueOff, ZoneSettingBrotliValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrotliEditable bool

const (
	ZoneSettingBrotliEditableTrue  ZoneSettingBrotliEditable = true
	ZoneSettingBrotliEditableFalse ZoneSettingBrotliEditable = false
)

func (r ZoneSettingBrotliEditable) IsKnown() bool {
	switch r {
	case ZoneSettingBrotliEditableTrue, ZoneSettingBrotliEditableFalse:
		return true
	}
	return false
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingBrotliParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingBrotliID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingBrotliValue] `json:"value,required"`
}

func (r ZoneSettingBrotliParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingBrotliParam) implementsZonesSettingEditParamsItemUnion() {}

type SettingBrotliEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingBrotliEditParamsValue] `json:"value,required"`
}

func (r SettingBrotliEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingBrotliEditParamsValue string

const (
	SettingBrotliEditParamsValueOff SettingBrotliEditParamsValue = "off"
	SettingBrotliEditParamsValueOn  SettingBrotliEditParamsValue = "on"
)

func (r SettingBrotliEditParamsValue) IsKnown() bool {
	switch r {
	case SettingBrotliEditParamsValueOff, SettingBrotliEditParamsValueOn:
		return true
	}
	return false
}

type SettingBrotliEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result ZoneSettingBrotli                     `json:"result"`
	JSON   settingBrotliEditResponseEnvelopeJSON `json:"-"`
}

// settingBrotliEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingBrotliEditResponseEnvelope]
type settingBrotliEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrotliEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingBrotliGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingBrotliGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result ZoneSettingBrotli                    `json:"result"`
	JSON   settingBrotliGetResponseEnvelopeJSON `json:"-"`
}

// settingBrotliGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingBrotliGetResponseEnvelope]
type settingBrotliGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrotliGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
