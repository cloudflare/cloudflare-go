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

// ZoneSettingAutomaticHTTPSRewriteService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingAutomaticHTTPSRewriteService] method instead.
type ZoneSettingAutomaticHTTPSRewriteService struct {
	Options []option.RequestOption
}

// NewZoneSettingAutomaticHTTPSRewriteService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingAutomaticHTTPSRewriteService(opts ...option.RequestOption) (r *ZoneSettingAutomaticHTTPSRewriteService) {
	r = &ZoneSettingAutomaticHTTPSRewriteService{}
	r.Options = opts
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
func (r *ZoneSettingAutomaticHTTPSRewriteService) Edit(ctx context.Context, params ZoneSettingAutomaticHTTPSRewriteEditParams, opts ...option.RequestOption) (res *ZoneSettingAutomaticHTTPSRewriteEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingAutomaticHTTPSRewriteEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
func (r *ZoneSettingAutomaticHTTPSRewriteService) Get(ctx context.Context, query ZoneSettingAutomaticHTTPSRewriteGetParams, opts ...option.RequestOption) (res *ZoneSettingAutomaticHTTPSRewriteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingAutomaticHTTPSRewriteGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingAutomaticHTTPSRewriteEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingAutomaticHTTPSRewriteEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingAutomaticHTTPSRewriteEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAutomaticHTTPSRewriteEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingAutomaticHTTPSRewriteEditResponseJSON `json:"-"`
}

// zoneSettingAutomaticHTTPSRewriteEditResponseJSON contains the JSON metadata for
// the struct [ZoneSettingAutomaticHTTPSRewriteEditResponse]
type zoneSettingAutomaticHTTPSRewriteEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPSRewriteEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingAutomaticHTTPSRewriteEditResponseID string

const (
	ZoneSettingAutomaticHTTPSRewriteEditResponseIDAutomaticHTTPSRewrites ZoneSettingAutomaticHTTPSRewriteEditResponseID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type ZoneSettingAutomaticHTTPSRewriteEditResponseValue string

const (
	ZoneSettingAutomaticHTTPSRewriteEditResponseValueOn  ZoneSettingAutomaticHTTPSRewriteEditResponseValue = "on"
	ZoneSettingAutomaticHTTPSRewriteEditResponseValueOff ZoneSettingAutomaticHTTPSRewriteEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAutomaticHTTPSRewriteEditResponseEditable bool

const (
	ZoneSettingAutomaticHTTPSRewriteEditResponseEditableTrue  ZoneSettingAutomaticHTTPSRewriteEditResponseEditable = true
	ZoneSettingAutomaticHTTPSRewriteEditResponseEditableFalse ZoneSettingAutomaticHTTPSRewriteEditResponseEditable = false
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingAutomaticHTTPSRewriteGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingAutomaticHTTPSRewriteGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingAutomaticHTTPSRewriteGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAutomaticHTTPSRewriteGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingAutomaticHTTPSRewriteGetResponseJSON `json:"-"`
}

// zoneSettingAutomaticHTTPSRewriteGetResponseJSON contains the JSON metadata for
// the struct [ZoneSettingAutomaticHTTPSRewriteGetResponse]
type zoneSettingAutomaticHTTPSRewriteGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPSRewriteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingAutomaticHTTPSRewriteGetResponseID string

const (
	ZoneSettingAutomaticHTTPSRewriteGetResponseIDAutomaticHTTPSRewrites ZoneSettingAutomaticHTTPSRewriteGetResponseID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type ZoneSettingAutomaticHTTPSRewriteGetResponseValue string

const (
	ZoneSettingAutomaticHTTPSRewriteGetResponseValueOn  ZoneSettingAutomaticHTTPSRewriteGetResponseValue = "on"
	ZoneSettingAutomaticHTTPSRewriteGetResponseValueOff ZoneSettingAutomaticHTTPSRewriteGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAutomaticHTTPSRewriteGetResponseEditable bool

const (
	ZoneSettingAutomaticHTTPSRewriteGetResponseEditableTrue  ZoneSettingAutomaticHTTPSRewriteGetResponseEditable = true
	ZoneSettingAutomaticHTTPSRewriteGetResponseEditableFalse ZoneSettingAutomaticHTTPSRewriteGetResponseEditable = false
)

type ZoneSettingAutomaticHTTPSRewriteEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingAutomaticHTTPSRewriteEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingAutomaticHTTPSRewriteEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingAutomaticHTTPSRewriteEditParamsValue string

const (
	ZoneSettingAutomaticHTTPSRewriteEditParamsValueOn  ZoneSettingAutomaticHTTPSRewriteEditParamsValue = "on"
	ZoneSettingAutomaticHTTPSRewriteEditParamsValueOff ZoneSettingAutomaticHTTPSRewriteEditParamsValue = "off"
)

type ZoneSettingAutomaticHTTPSRewriteEditResponseEnvelope struct {
	Errors   []ZoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result ZoneSettingAutomaticHTTPSRewriteEditResponse             `json:"result"`
	JSON   zoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZoneSettingAutomaticHTTPSRewriteEditResponseEnvelope]
type zoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPSRewriteEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeErrors]
type zoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeMessages]
type zoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPSRewriteEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticHTTPSRewriteGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingAutomaticHTTPSRewriteGetResponseEnvelope struct {
	Errors   []ZoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result ZoneSettingAutomaticHTTPSRewriteGetResponse             `json:"result"`
	JSON   zoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZoneSettingAutomaticHTTPSRewriteGetResponseEnvelope]
type zoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPSRewriteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeErrors]
type zoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeMessages]
type zoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPSRewriteGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
