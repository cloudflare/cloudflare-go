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

// ZoneSettingEarlyHintService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingEarlyHintService]
// method instead.
type ZoneSettingEarlyHintService struct {
	Options []option.RequestOption
}

// NewZoneSettingEarlyHintService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingEarlyHintService(opts ...option.RequestOption) (r *ZoneSettingEarlyHintService) {
	r = &ZoneSettingEarlyHintService{}
	r.Options = opts
	return
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
func (r *ZoneSettingEarlyHintService) Edit(ctx context.Context, params ZoneSettingEarlyHintEditParams, opts ...option.RequestOption) (res *ZonesEarlyHints, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingEarlyHintEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/early_hints", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
func (r *ZoneSettingEarlyHintService) Get(ctx context.Context, query ZoneSettingEarlyHintGetParams, opts ...option.RequestOption) (res *ZonesEarlyHints, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingEarlyHintGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/early_hints", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZonesEarlyHints struct {
	// ID of the zone setting.
	ID ZonesEarlyHintsID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesEarlyHintsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesEarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesEarlyHintsJSON `json:"-"`
}

// zonesEarlyHintsJSON contains the JSON metadata for the struct [ZonesEarlyHints]
type zonesEarlyHintsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesEarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesEarlyHints) implementsZoneSettingEditResponse() {}

func (r ZonesEarlyHints) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesEarlyHintsID string

const (
	ZonesEarlyHintsIDEarlyHints ZonesEarlyHintsID = "early_hints"
)

// Current value of the zone setting.
type ZonesEarlyHintsValue string

const (
	ZonesEarlyHintsValueOn  ZonesEarlyHintsValue = "on"
	ZonesEarlyHintsValueOff ZonesEarlyHintsValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesEarlyHintsEditable bool

const (
	ZonesEarlyHintsEditableTrue  ZonesEarlyHintsEditable = true
	ZonesEarlyHintsEditableFalse ZonesEarlyHintsEditable = false
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZonesEarlyHintsParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesEarlyHintsID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesEarlyHintsValue] `json:"value,required"`
}

func (r ZonesEarlyHintsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesEarlyHintsParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingEarlyHintEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEarlyHintEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingEarlyHintEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingEarlyHintEditParamsValue string

const (
	ZoneSettingEarlyHintEditParamsValueOn  ZoneSettingEarlyHintEditParamsValue = "on"
	ZoneSettingEarlyHintEditParamsValueOff ZoneSettingEarlyHintEditParamsValue = "off"
)

type ZoneSettingEarlyHintEditResponseEnvelope struct {
	Errors   []ZoneSettingEarlyHintEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingEarlyHintEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result ZonesEarlyHints                              `json:"result"`
	JSON   zoneSettingEarlyHintEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingEarlyHintEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingEarlyHintEditResponseEnvelope]
type zoneSettingEarlyHintEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEarlyHintEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingEarlyHintEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingEarlyHintEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingEarlyHintEditResponseEnvelopeErrors]
type zoneSettingEarlyHintEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEarlyHintEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingEarlyHintEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingEarlyHintEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneSettingEarlyHintEditResponseEnvelopeMessages]
type zoneSettingEarlyHintEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEarlyHintGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingEarlyHintGetResponseEnvelope struct {
	Errors   []ZoneSettingEarlyHintGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingEarlyHintGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result ZonesEarlyHints                             `json:"result"`
	JSON   zoneSettingEarlyHintGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingEarlyHintGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingEarlyHintGetResponseEnvelope]
type zoneSettingEarlyHintGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEarlyHintGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingEarlyHintGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingEarlyHintGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingEarlyHintGetResponseEnvelopeErrors]
type zoneSettingEarlyHintGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEarlyHintGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneSettingEarlyHintGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingEarlyHintGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneSettingEarlyHintGetResponseEnvelopeMessages]
type zoneSettingEarlyHintGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
