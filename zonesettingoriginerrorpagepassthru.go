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

// ZoneSettingOriginErrorPagePassThruService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingOriginErrorPagePassThruService] method instead.
type ZoneSettingOriginErrorPagePassThruService struct {
	Options []option.RequestOption
}

// NewZoneSettingOriginErrorPagePassThruService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneSettingOriginErrorPagePassThruService(opts ...option.RequestOption) (r *ZoneSettingOriginErrorPagePassThruService) {
	r = &ZoneSettingOriginErrorPagePassThruService{}
	r.Options = opts
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
func (r *ZoneSettingOriginErrorPagePassThruService) Edit(ctx context.Context, params ZoneSettingOriginErrorPagePassThruEditParams, opts ...option.RequestOption) (res *ZonesOriginErrorPagePassThru, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingOriginErrorPagePassThruEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/origin_error_page_pass_thru", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
func (r *ZoneSettingOriginErrorPagePassThruService) Get(ctx context.Context, query ZoneSettingOriginErrorPagePassThruGetParams, opts ...option.RequestOption) (res *ZonesOriginErrorPagePassThru, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingOriginErrorPagePassThruGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/origin_error_page_pass_thru", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type ZonesOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID ZonesOriginErrorPagePassThruID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesOriginErrorPagePassThruValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesOriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesOriginErrorPagePassThruJSON `json:"-"`
}

// zonesOriginErrorPagePassThruJSON contains the JSON metadata for the struct
// [ZonesOriginErrorPagePassThru]
type zonesOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesOriginErrorPagePassThru) implementsZoneSettingEditResponse() {}

func (r ZonesOriginErrorPagePassThru) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesOriginErrorPagePassThruID string

const (
	ZonesOriginErrorPagePassThruIDOriginErrorPagePassThru ZonesOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

// Current value of the zone setting.
type ZonesOriginErrorPagePassThruValue string

const (
	ZonesOriginErrorPagePassThruValueOn  ZonesOriginErrorPagePassThruValue = "on"
	ZonesOriginErrorPagePassThruValueOff ZonesOriginErrorPagePassThruValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesOriginErrorPagePassThruEditable bool

const (
	ZonesOriginErrorPagePassThruEditableTrue  ZonesOriginErrorPagePassThruEditable = true
	ZonesOriginErrorPagePassThruEditableFalse ZonesOriginErrorPagePassThruEditable = false
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type ZonesOriginErrorPagePassThruParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesOriginErrorPagePassThruID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesOriginErrorPagePassThruValue] `json:"value,required"`
}

func (r ZonesOriginErrorPagePassThruParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesOriginErrorPagePassThruParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingOriginErrorPagePassThruEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingOriginErrorPagePassThruEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingOriginErrorPagePassThruEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingOriginErrorPagePassThruEditParamsValue string

const (
	ZoneSettingOriginErrorPagePassThruEditParamsValueOn  ZoneSettingOriginErrorPagePassThruEditParamsValue = "on"
	ZoneSettingOriginErrorPagePassThruEditParamsValueOff ZoneSettingOriginErrorPagePassThruEditParamsValue = "off"
)

type ZoneSettingOriginErrorPagePassThruEditResponseEnvelope struct {
	Errors   []ZoneSettingOriginErrorPagePassThruEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingOriginErrorPagePassThruEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will proxy customer error pages on any 502,504 errors on origin
	// server instead of showing a default Cloudflare error page. This does not apply
	// to 522 errors and is limited to Enterprise Zones.
	Result ZonesOriginErrorPagePassThru                               `json:"result"`
	JSON   zoneSettingOriginErrorPagePassThruEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingOriginErrorPagePassThruEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZoneSettingOriginErrorPagePassThruEditResponseEnvelope]
type zoneSettingOriginErrorPagePassThruEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThruEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginErrorPagePassThruEditResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zoneSettingOriginErrorPagePassThruEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingOriginErrorPagePassThruEditResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZoneSettingOriginErrorPagePassThruEditResponseEnvelopeErrors]
type zoneSettingOriginErrorPagePassThruEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThruEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginErrorPagePassThruEditResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    zoneSettingOriginErrorPagePassThruEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingOriginErrorPagePassThruEditResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZoneSettingOriginErrorPagePassThruEditResponseEnvelopeMessages]
type zoneSettingOriginErrorPagePassThruEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThruEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginErrorPagePassThruGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingOriginErrorPagePassThruGetResponseEnvelope struct {
	Errors   []ZoneSettingOriginErrorPagePassThruGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingOriginErrorPagePassThruGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will proxy customer error pages on any 502,504 errors on origin
	// server instead of showing a default Cloudflare error page. This does not apply
	// to 522 errors and is limited to Enterprise Zones.
	Result ZonesOriginErrorPagePassThru                              `json:"result"`
	JSON   zoneSettingOriginErrorPagePassThruGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingOriginErrorPagePassThruGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZoneSettingOriginErrorPagePassThruGetResponseEnvelope]
type zoneSettingOriginErrorPagePassThruGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThruGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginErrorPagePassThruGetResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zoneSettingOriginErrorPagePassThruGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingOriginErrorPagePassThruGetResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZoneSettingOriginErrorPagePassThruGetResponseEnvelopeErrors]
type zoneSettingOriginErrorPagePassThruGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThruGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginErrorPagePassThruGetResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zoneSettingOriginErrorPagePassThruGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingOriginErrorPagePassThruGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZoneSettingOriginErrorPagePassThruGetResponseEnvelopeMessages]
type zoneSettingOriginErrorPagePassThruGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThruGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
