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

// SettingHotlinkProtectionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingHotlinkProtectionService] method instead.
type SettingHotlinkProtectionService struct {
	Options []option.RequestOption
}

// NewSettingHotlinkProtectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingHotlinkProtectionService(opts ...option.RequestOption) (r *SettingHotlinkProtectionService) {
	r = &SettingHotlinkProtectionService{}
	r.Options = opts
	return
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
func (r *SettingHotlinkProtectionService) Edit(ctx context.Context, params SettingHotlinkProtectionEditParams, opts ...option.RequestOption) (res *SettingHotlinkProtectionEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHotlinkProtectionEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/hotlink_protection", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
func (r *SettingHotlinkProtectionService) Get(ctx context.Context, query SettingHotlinkProtectionGetParams, opts ...option.RequestOption) (res *SettingHotlinkProtectionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHotlinkProtectionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/hotlink_protection", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type SettingHotlinkProtectionEditResponse struct {
	// ID of the zone setting.
	ID SettingHotlinkProtectionEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingHotlinkProtectionEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingHotlinkProtectionEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingHotlinkProtectionEditResponseJSON `json:"-"`
}

// settingHotlinkProtectionEditResponseJSON contains the JSON metadata for the
// struct [SettingHotlinkProtectionEditResponse]
type settingHotlinkProtectionEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHotlinkProtectionEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingHotlinkProtectionEditResponseID string

const (
	SettingHotlinkProtectionEditResponseIDHotlinkProtection SettingHotlinkProtectionEditResponseID = "hotlink_protection"
)

// Current value of the zone setting.
type SettingHotlinkProtectionEditResponseValue string

const (
	SettingHotlinkProtectionEditResponseValueOn  SettingHotlinkProtectionEditResponseValue = "on"
	SettingHotlinkProtectionEditResponseValueOff SettingHotlinkProtectionEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingHotlinkProtectionEditResponseEditable bool

const (
	SettingHotlinkProtectionEditResponseEditableTrue  SettingHotlinkProtectionEditResponseEditable = true
	SettingHotlinkProtectionEditResponseEditableFalse SettingHotlinkProtectionEditResponseEditable = false
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type SettingHotlinkProtectionGetResponse struct {
	// ID of the zone setting.
	ID SettingHotlinkProtectionGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingHotlinkProtectionGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingHotlinkProtectionGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingHotlinkProtectionGetResponseJSON `json:"-"`
}

// settingHotlinkProtectionGetResponseJSON contains the JSON metadata for the
// struct [SettingHotlinkProtectionGetResponse]
type settingHotlinkProtectionGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHotlinkProtectionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingHotlinkProtectionGetResponseID string

const (
	SettingHotlinkProtectionGetResponseIDHotlinkProtection SettingHotlinkProtectionGetResponseID = "hotlink_protection"
)

// Current value of the zone setting.
type SettingHotlinkProtectionGetResponseValue string

const (
	SettingHotlinkProtectionGetResponseValueOn  SettingHotlinkProtectionGetResponseValue = "on"
	SettingHotlinkProtectionGetResponseValueOff SettingHotlinkProtectionGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingHotlinkProtectionGetResponseEditable bool

const (
	SettingHotlinkProtectionGetResponseEditableTrue  SettingHotlinkProtectionGetResponseEditable = true
	SettingHotlinkProtectionGetResponseEditableFalse SettingHotlinkProtectionGetResponseEditable = false
)

type SettingHotlinkProtectionEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingHotlinkProtectionEditParamsValue] `json:"value,required"`
}

func (r SettingHotlinkProtectionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingHotlinkProtectionEditParamsValue string

const (
	SettingHotlinkProtectionEditParamsValueOn  SettingHotlinkProtectionEditParamsValue = "on"
	SettingHotlinkProtectionEditParamsValueOff SettingHotlinkProtectionEditParamsValue = "off"
)

type SettingHotlinkProtectionEditResponseEnvelope struct {
	Errors   []SettingHotlinkProtectionEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingHotlinkProtectionEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, the Hotlink Protection option ensures that other sites cannot suck
	// up your bandwidth by building pages that use images hosted on your site. Anytime
	// a request for an image on your site hits Cloudflare, we check to ensure that
	// it's not another site requesting them. People will still be able to download and
	// view images from your page, but other sites won't be able to steal them for use
	// on their own pages.
	// (https://support.cloudflare.com/hc/en-us/articles/200170026).
	Result SettingHotlinkProtectionEditResponse             `json:"result"`
	JSON   settingHotlinkProtectionEditResponseEnvelopeJSON `json:"-"`
}

// settingHotlinkProtectionEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingHotlinkProtectionEditResponseEnvelope]
type settingHotlinkProtectionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHotlinkProtectionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHotlinkProtectionEditResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingHotlinkProtectionEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingHotlinkProtectionEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingHotlinkProtectionEditResponseEnvelopeErrors]
type settingHotlinkProtectionEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHotlinkProtectionEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHotlinkProtectionEditResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    settingHotlinkProtectionEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingHotlinkProtectionEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingHotlinkProtectionEditResponseEnvelopeMessages]
type settingHotlinkProtectionEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHotlinkProtectionEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHotlinkProtectionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingHotlinkProtectionGetResponseEnvelope struct {
	Errors   []SettingHotlinkProtectionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingHotlinkProtectionGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, the Hotlink Protection option ensures that other sites cannot suck
	// up your bandwidth by building pages that use images hosted on your site. Anytime
	// a request for an image on your site hits Cloudflare, we check to ensure that
	// it's not another site requesting them. People will still be able to download and
	// view images from your page, but other sites won't be able to steal them for use
	// on their own pages.
	// (https://support.cloudflare.com/hc/en-us/articles/200170026).
	Result SettingHotlinkProtectionGetResponse             `json:"result"`
	JSON   settingHotlinkProtectionGetResponseEnvelopeJSON `json:"-"`
}

// settingHotlinkProtectionGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingHotlinkProtectionGetResponseEnvelope]
type settingHotlinkProtectionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHotlinkProtectionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHotlinkProtectionGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingHotlinkProtectionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingHotlinkProtectionGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingHotlinkProtectionGetResponseEnvelopeErrors]
type settingHotlinkProtectionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHotlinkProtectionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHotlinkProtectionGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingHotlinkProtectionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingHotlinkProtectionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingHotlinkProtectionGetResponseEnvelopeMessages]
type settingHotlinkProtectionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHotlinkProtectionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
