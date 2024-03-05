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

// ZoneSettingOrangeToOrangeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingOrangeToOrangeService] method instead.
type ZoneSettingOrangeToOrangeService struct {
	Options []option.RequestOption
}

// NewZoneSettingOrangeToOrangeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingOrangeToOrangeService(opts ...option.RequestOption) (r *ZoneSettingOrangeToOrangeService) {
	r = &ZoneSettingOrangeToOrangeService{}
	r.Options = opts
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
func (r *ZoneSettingOrangeToOrangeService) Edit(ctx context.Context, params ZoneSettingOrangeToOrangeEditParams, opts ...option.RequestOption) (res *ZonesOrangeToOrange, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingOrangeToOrangeEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/orange_to_orange", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
func (r *ZoneSettingOrangeToOrangeService) Get(ctx context.Context, query ZoneSettingOrangeToOrangeGetParams, opts ...option.RequestOption) (res *ZonesOrangeToOrange, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingOrangeToOrangeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/orange_to_orange", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID ZonesOrangeToOrangeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesOrangeToOrangeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesOrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time               `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesOrangeToOrangeJSON `json:"-"`
}

// zonesOrangeToOrangeJSON contains the JSON metadata for the struct
// [ZonesOrangeToOrange]
type zonesOrangeToOrangeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesOrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesOrangeToOrange) implementsZoneSettingEditResponse() {}

func (r ZonesOrangeToOrange) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesOrangeToOrangeID string

const (
	ZonesOrangeToOrangeIDOrangeToOrange ZonesOrangeToOrangeID = "orange_to_orange"
)

// Current value of the zone setting.
type ZonesOrangeToOrangeValue string

const (
	ZonesOrangeToOrangeValueOn  ZonesOrangeToOrangeValue = "on"
	ZonesOrangeToOrangeValueOff ZonesOrangeToOrangeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesOrangeToOrangeEditable bool

const (
	ZonesOrangeToOrangeEditableTrue  ZonesOrangeToOrangeEditable = true
	ZonesOrangeToOrangeEditableFalse ZonesOrangeToOrangeEditable = false
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZonesOrangeToOrangeParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesOrangeToOrangeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesOrangeToOrangeValue] `json:"value,required"`
}

func (r ZonesOrangeToOrangeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesOrangeToOrangeParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingOrangeToOrangeEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Value param.Field[ZonesOrangeToOrangeParam] `json:"value,required"`
}

func (r ZoneSettingOrangeToOrangeEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingOrangeToOrangeEditResponseEnvelope struct {
	Errors   []ZoneSettingOrangeToOrangeEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingOrangeToOrangeEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result ZonesOrangeToOrange                               `json:"result"`
	JSON   zoneSettingOrangeToOrangeEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingOrangeToOrangeEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingOrangeToOrangeEditResponseEnvelope]
type zoneSettingOrangeToOrangeEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOrangeToOrangeEditResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingOrangeToOrangeEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingOrangeToOrangeEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingOrangeToOrangeEditResponseEnvelopeErrors]
type zoneSettingOrangeToOrangeEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOrangeToOrangeEditResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingOrangeToOrangeEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingOrangeToOrangeEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingOrangeToOrangeEditResponseEnvelopeMessages]
type zoneSettingOrangeToOrangeEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOrangeToOrangeGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingOrangeToOrangeGetResponseEnvelope struct {
	Errors   []ZoneSettingOrangeToOrangeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingOrangeToOrangeGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result ZonesOrangeToOrange                              `json:"result"`
	JSON   zoneSettingOrangeToOrangeGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingOrangeToOrangeGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingOrangeToOrangeGetResponseEnvelope]
type zoneSettingOrangeToOrangeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOrangeToOrangeGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingOrangeToOrangeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingOrangeToOrangeGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingOrangeToOrangeGetResponseEnvelopeErrors]
type zoneSettingOrangeToOrangeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOrangeToOrangeGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingOrangeToOrangeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingOrangeToOrangeGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingOrangeToOrangeGetResponseEnvelopeMessages]
type zoneSettingOrangeToOrangeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
