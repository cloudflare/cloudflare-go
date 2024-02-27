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

// ZoneSettingHTTP3Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingHTTP3Service] method
// instead.
type ZoneSettingHTTP3Service struct {
	Options []option.RequestOption
}

// NewZoneSettingHTTP3Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingHTTP3Service(opts ...option.RequestOption) (r *ZoneSettingHTTP3Service) {
	r = &ZoneSettingHTTP3Service{}
	r.Options = opts
	return
}

// Value of the HTTP3 setting.
func (r *ZoneSettingHTTP3Service) Edit(ctx context.Context, params ZoneSettingHTTP3EditParams, opts ...option.RequestOption) (res *ZoneSettingHTTP3EditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingHTTP3EditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http3", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Value of the HTTP3 setting.
func (r *ZoneSettingHTTP3Service) Get(ctx context.Context, query ZoneSettingHTTP3GetParams, opts ...option.RequestOption) (res *ZoneSettingHTTP3GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingHTTP3GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http3", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// HTTP3 enabled for this zone.
type ZoneSettingHTTP3EditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingHTTP3EditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingHTTP3EditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHTTP3EditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingHTTP3EditResponseJSON `json:"-"`
}

// zoneSettingHTTP3EditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingHTTP3EditResponse]
type zoneSettingHTTP3EditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3EditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingHTTP3EditResponseID string

const (
	ZoneSettingHTTP3EditResponseIDHTTP3 ZoneSettingHTTP3EditResponseID = "http3"
)

// Current value of the zone setting.
type ZoneSettingHTTP3EditResponseValue string

const (
	ZoneSettingHTTP3EditResponseValueOn  ZoneSettingHTTP3EditResponseValue = "on"
	ZoneSettingHTTP3EditResponseValueOff ZoneSettingHTTP3EditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHTTP3EditResponseEditable bool

const (
	ZoneSettingHTTP3EditResponseEditableTrue  ZoneSettingHTTP3EditResponseEditable = true
	ZoneSettingHTTP3EditResponseEditableFalse ZoneSettingHTTP3EditResponseEditable = false
)

// HTTP3 enabled for this zone.
type ZoneSettingHTTP3GetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingHTTP3GetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingHTTP3GetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHTTP3GetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingHTTP3GetResponseJSON `json:"-"`
}

// zoneSettingHTTP3GetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingHTTP3GetResponse]
type zoneSettingHTTP3GetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingHTTP3GetResponseID string

const (
	ZoneSettingHTTP3GetResponseIDHTTP3 ZoneSettingHTTP3GetResponseID = "http3"
)

// Current value of the zone setting.
type ZoneSettingHTTP3GetResponseValue string

const (
	ZoneSettingHTTP3GetResponseValueOn  ZoneSettingHTTP3GetResponseValue = "on"
	ZoneSettingHTTP3GetResponseValueOff ZoneSettingHTTP3GetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHTTP3GetResponseEditable bool

const (
	ZoneSettingHTTP3GetResponseEditableTrue  ZoneSettingHTTP3GetResponseEditable = true
	ZoneSettingHTTP3GetResponseEditableFalse ZoneSettingHTTP3GetResponseEditable = false
)

type ZoneSettingHTTP3EditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the HTTP3 setting.
	Value param.Field[ZoneSettingHTTP3EditParamsValue] `json:"value,required"`
}

func (r ZoneSettingHTTP3EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the HTTP3 setting.
type ZoneSettingHTTP3EditParamsValue string

const (
	ZoneSettingHTTP3EditParamsValueOn  ZoneSettingHTTP3EditParamsValue = "on"
	ZoneSettingHTTP3EditParamsValueOff ZoneSettingHTTP3EditParamsValue = "off"
)

type ZoneSettingHTTP3EditResponseEnvelope struct {
	Errors   []ZoneSettingHTTP3EditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingHTTP3EditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP3 enabled for this zone.
	Result ZoneSettingHTTP3EditResponse             `json:"result"`
	JSON   zoneSettingHTTP3EditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingHTTP3EditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP3EditResponseEnvelope]
type zoneSettingHTTP3EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP3EditResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingHTTP3EditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingHTTP3EditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingHTTP3EditResponseEnvelopeErrors]
type zoneSettingHTTP3EditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3EditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP3EditResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingHTTP3EditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingHTTP3EditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingHTTP3EditResponseEnvelopeMessages]
type zoneSettingHTTP3EditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3EditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP3GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingHTTP3GetResponseEnvelope struct {
	Errors   []ZoneSettingHTTP3GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingHTTP3GetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP3 enabled for this zone.
	Result ZoneSettingHTTP3GetResponse             `json:"result"`
	JSON   zoneSettingHTTP3GetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingHTTP3GetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP3GetResponseEnvelope]
type zoneSettingHTTP3GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP3GetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingHTTP3GetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingHTTP3GetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP3GetResponseEnvelopeErrors]
type zoneSettingHTTP3GetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3GetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP3GetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingHTTP3GetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingHTTP3GetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingHTTP3GetResponseEnvelopeMessages]
type zoneSettingHTTP3GetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3GetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
