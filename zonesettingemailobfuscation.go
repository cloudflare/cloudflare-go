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

// ZoneSettingEmailObfuscationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingEmailObfuscationService] method instead.
type ZoneSettingEmailObfuscationService struct {
	Options []option.RequestOption
}

// NewZoneSettingEmailObfuscationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingEmailObfuscationService(opts ...option.RequestOption) (r *ZoneSettingEmailObfuscationService) {
	r = &ZoneSettingEmailObfuscationService{}
	r.Options = opts
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
func (r *ZoneSettingEmailObfuscationService) Edit(ctx context.Context, params ZoneSettingEmailObfuscationEditParams, opts ...option.RequestOption) (res *ZoneSettingEmailObfuscationEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingEmailObfuscationEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/email_obfuscation", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
func (r *ZoneSettingEmailObfuscationService) Get(ctx context.Context, query ZoneSettingEmailObfuscationGetParams, opts ...option.RequestOption) (res *ZoneSettingEmailObfuscationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingEmailObfuscationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/email_obfuscation", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZoneSettingEmailObfuscationEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingEmailObfuscationEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEmailObfuscationEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEmailObfuscationEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEmailObfuscationEditResponseJSON `json:"-"`
}

// zoneSettingEmailObfuscationEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingEmailObfuscationEditResponse]
type zoneSettingEmailObfuscationEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingEmailObfuscationEditResponseID string

const (
	ZoneSettingEmailObfuscationEditResponseIDEmailObfuscation ZoneSettingEmailObfuscationEditResponseID = "email_obfuscation"
)

// Current value of the zone setting.
type ZoneSettingEmailObfuscationEditResponseValue string

const (
	ZoneSettingEmailObfuscationEditResponseValueOn  ZoneSettingEmailObfuscationEditResponseValue = "on"
	ZoneSettingEmailObfuscationEditResponseValueOff ZoneSettingEmailObfuscationEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEmailObfuscationEditResponseEditable bool

const (
	ZoneSettingEmailObfuscationEditResponseEditableTrue  ZoneSettingEmailObfuscationEditResponseEditable = true
	ZoneSettingEmailObfuscationEditResponseEditableFalse ZoneSettingEmailObfuscationEditResponseEditable = false
)

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZoneSettingEmailObfuscationGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingEmailObfuscationGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEmailObfuscationGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEmailObfuscationGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEmailObfuscationGetResponseJSON `json:"-"`
}

// zoneSettingEmailObfuscationGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingEmailObfuscationGetResponse]
type zoneSettingEmailObfuscationGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingEmailObfuscationGetResponseID string

const (
	ZoneSettingEmailObfuscationGetResponseIDEmailObfuscation ZoneSettingEmailObfuscationGetResponseID = "email_obfuscation"
)

// Current value of the zone setting.
type ZoneSettingEmailObfuscationGetResponseValue string

const (
	ZoneSettingEmailObfuscationGetResponseValueOn  ZoneSettingEmailObfuscationGetResponseValue = "on"
	ZoneSettingEmailObfuscationGetResponseValueOff ZoneSettingEmailObfuscationGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEmailObfuscationGetResponseEditable bool

const (
	ZoneSettingEmailObfuscationGetResponseEditableTrue  ZoneSettingEmailObfuscationGetResponseEditable = true
	ZoneSettingEmailObfuscationGetResponseEditableFalse ZoneSettingEmailObfuscationGetResponseEditable = false
)

type ZoneSettingEmailObfuscationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEmailObfuscationEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingEmailObfuscationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingEmailObfuscationEditParamsValue string

const (
	ZoneSettingEmailObfuscationEditParamsValueOn  ZoneSettingEmailObfuscationEditParamsValue = "on"
	ZoneSettingEmailObfuscationEditParamsValueOff ZoneSettingEmailObfuscationEditParamsValue = "off"
)

type ZoneSettingEmailObfuscationEditResponseEnvelope struct {
	Errors   []ZoneSettingEmailObfuscationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingEmailObfuscationEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result ZoneSettingEmailObfuscationEditResponse             `json:"result"`
	JSON   zoneSettingEmailObfuscationEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingEmailObfuscationEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingEmailObfuscationEditResponseEnvelope]
type zoneSettingEmailObfuscationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEmailObfuscationEditResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingEmailObfuscationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingEmailObfuscationEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingEmailObfuscationEditResponseEnvelopeErrors]
type zoneSettingEmailObfuscationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEmailObfuscationEditResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zoneSettingEmailObfuscationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingEmailObfuscationEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingEmailObfuscationEditResponseEnvelopeMessages]
type zoneSettingEmailObfuscationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEmailObfuscationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingEmailObfuscationGetResponseEnvelope struct {
	Errors   []ZoneSettingEmailObfuscationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingEmailObfuscationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result ZoneSettingEmailObfuscationGetResponse             `json:"result"`
	JSON   zoneSettingEmailObfuscationGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingEmailObfuscationGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingEmailObfuscationGetResponseEnvelope]
type zoneSettingEmailObfuscationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEmailObfuscationGetResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingEmailObfuscationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingEmailObfuscationGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingEmailObfuscationGetResponseEnvelopeErrors]
type zoneSettingEmailObfuscationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEmailObfuscationGetResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneSettingEmailObfuscationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingEmailObfuscationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingEmailObfuscationGetResponseEnvelopeMessages]
type zoneSettingEmailObfuscationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
