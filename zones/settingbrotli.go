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
func (r *SettingBrotliService) Edit(ctx context.Context, params SettingBrotliEditParams, opts ...option.RequestOption) (res *ZonesBrotli, err error) {
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
func (r *SettingBrotliService) Get(ctx context.Context, query SettingBrotliGetParams, opts ...option.RequestOption) (res *ZonesBrotli, err error) {
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
type ZonesBrotli struct {
	// ID of the zone setting.
	ID ZonesBrotliID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesBrotliValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesBrotliEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time       `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesBrotliJSON `json:"-"`
}

// zonesBrotliJSON contains the JSON metadata for the struct [ZonesBrotli]
type zonesBrotliJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesBrotli) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesBrotliJSON) RawJSON() string {
	return r.raw
}

func (r ZonesBrotli) implementsZonesSettingEditResponse() {}

func (r ZonesBrotli) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesBrotliID string

const (
	ZonesBrotliIDBrotli ZonesBrotliID = "brotli"
)

// Current value of the zone setting.
type ZonesBrotliValue string

const (
	ZonesBrotliValueOff ZonesBrotliValue = "off"
	ZonesBrotliValueOn  ZonesBrotliValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesBrotliEditable bool

const (
	ZonesBrotliEditableTrue  ZonesBrotliEditable = true
	ZonesBrotliEditableFalse ZonesBrotliEditable = false
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZonesBrotliParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesBrotliID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesBrotliValue] `json:"value,required"`
}

func (r ZonesBrotliParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesBrotliParam) implementsZonesSettingEditParamsItem() {}

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

type SettingBrotliEditResponseEnvelope struct {
	Errors   []SettingBrotliEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingBrotliEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result ZonesBrotli                           `json:"result"`
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

type SettingBrotliEditResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingBrotliEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingBrotliEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingBrotliEditResponseEnvelopeErrors]
type settingBrotliEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrotliEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingBrotliEditResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingBrotliEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingBrotliEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingBrotliEditResponseEnvelopeMessages]
type settingBrotliEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrotliEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingBrotliGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingBrotliGetResponseEnvelope struct {
	Errors   []SettingBrotliGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingBrotliGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result ZonesBrotli                          `json:"result"`
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

type SettingBrotliGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingBrotliGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingBrotliGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingBrotliGetResponseEnvelopeErrors]
type settingBrotliGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrotliGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingBrotliGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingBrotliGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingBrotliGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingBrotliGetResponseEnvelopeMessages]
type settingBrotliGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrotliGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
