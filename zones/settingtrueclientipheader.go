// File generated from our OpenAPI spec by Stainless.

package zones

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

// SettingTrueClientIPHeaderService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingTrueClientIPHeaderService] method instead.
type SettingTrueClientIPHeaderService struct {
	Options []option.RequestOption
}

// NewSettingTrueClientIPHeaderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingTrueClientIPHeaderService(opts ...option.RequestOption) (r *SettingTrueClientIPHeaderService) {
	r = &SettingTrueClientIPHeaderService{}
	r.Options = opts
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
func (r *SettingTrueClientIPHeaderService) Edit(ctx context.Context, params SettingTrueClientIPHeaderEditParams, opts ...option.RequestOption) (res *ZonesTrueClientIPHeader, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTrueClientIPHeaderEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/true_client_ip_header", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
func (r *SettingTrueClientIPHeaderService) Get(ctx context.Context, query SettingTrueClientIPHeaderGetParams, opts ...option.RequestOption) (res *ZonesTrueClientIPHeader, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTrueClientIPHeaderGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/true_client_ip_header", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZonesTrueClientIPHeader struct {
	// ID of the zone setting.
	ID ZonesTrueClientIPHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesTrueClientIPHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesTrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesTrueClientIPHeaderJSON `json:"-"`
}

// zonesTrueClientIPHeaderJSON contains the JSON metadata for the struct
// [ZonesTrueClientIPHeader]
type zonesTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesTrueClientIPHeaderJSON) RawJSON() string {
	return r.raw
}

func (r ZonesTrueClientIPHeader) implementsZonesSettingEditResponse() {}

func (r ZonesTrueClientIPHeader) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesTrueClientIPHeaderID string

const (
	ZonesTrueClientIPHeaderIDTrueClientIPHeader ZonesTrueClientIPHeaderID = "true_client_ip_header"
)

// Current value of the zone setting.
type ZonesTrueClientIPHeaderValue string

const (
	ZonesTrueClientIPHeaderValueOn  ZonesTrueClientIPHeaderValue = "on"
	ZonesTrueClientIPHeaderValueOff ZonesTrueClientIPHeaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesTrueClientIPHeaderEditable bool

const (
	ZonesTrueClientIPHeaderEditableTrue  ZonesTrueClientIPHeaderEditable = true
	ZonesTrueClientIPHeaderEditableFalse ZonesTrueClientIPHeaderEditable = false
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZonesTrueClientIPHeaderParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesTrueClientIPHeaderID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesTrueClientIPHeaderValue] `json:"value,required"`
}

func (r ZonesTrueClientIPHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesTrueClientIPHeaderParam) implementsZonesSettingEditParamsItem() {}

type SettingTrueClientIPHeaderEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingTrueClientIPHeaderEditParamsValue] `json:"value,required"`
}

func (r SettingTrueClientIPHeaderEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingTrueClientIPHeaderEditParamsValue string

const (
	SettingTrueClientIPHeaderEditParamsValueOn  SettingTrueClientIPHeaderEditParamsValue = "on"
	SettingTrueClientIPHeaderEditParamsValueOff SettingTrueClientIPHeaderEditParamsValue = "off"
)

type SettingTrueClientIPHeaderEditResponseEnvelope struct {
	Errors   []SettingTrueClientIPHeaderEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTrueClientIPHeaderEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result ZonesTrueClientIPHeader                           `json:"result"`
	JSON   settingTrueClientIPHeaderEditResponseEnvelopeJSON `json:"-"`
}

// settingTrueClientIPHeaderEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingTrueClientIPHeaderEditResponseEnvelope]
type settingTrueClientIPHeaderEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrueClientIPHeaderEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTrueClientIPHeaderEditResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingTrueClientIPHeaderEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingTrueClientIPHeaderEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingTrueClientIPHeaderEditResponseEnvelopeErrors]
type settingTrueClientIPHeaderEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrueClientIPHeaderEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingTrueClientIPHeaderEditResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    settingTrueClientIPHeaderEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingTrueClientIPHeaderEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingTrueClientIPHeaderEditResponseEnvelopeMessages]
type settingTrueClientIPHeaderEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrueClientIPHeaderEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingTrueClientIPHeaderGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingTrueClientIPHeaderGetResponseEnvelope struct {
	Errors   []SettingTrueClientIPHeaderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTrueClientIPHeaderGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result ZonesTrueClientIPHeader                          `json:"result"`
	JSON   settingTrueClientIPHeaderGetResponseEnvelopeJSON `json:"-"`
}

// settingTrueClientIPHeaderGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingTrueClientIPHeaderGetResponseEnvelope]
type settingTrueClientIPHeaderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrueClientIPHeaderGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTrueClientIPHeaderGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingTrueClientIPHeaderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingTrueClientIPHeaderGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingTrueClientIPHeaderGetResponseEnvelopeErrors]
type settingTrueClientIPHeaderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrueClientIPHeaderGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingTrueClientIPHeaderGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    settingTrueClientIPHeaderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingTrueClientIPHeaderGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingTrueClientIPHeaderGetResponseEnvelopeMessages]
type settingTrueClientIPHeaderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrueClientIPHeaderGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
