// File generated from our OpenAPI spec by Stainless.

package cloudflare

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

// ZoneSettingHTTP2Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingHTTP2Service] method
// instead.
type ZoneSettingHTTP2Service struct {
	Options []option.RequestOption
}

// NewZoneSettingHTTP2Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingHTTP2Service(opts ...option.RequestOption) (r *ZoneSettingHTTP2Service) {
	r = &ZoneSettingHTTP2Service{}
	r.Options = opts
	return
}

// Value of the HTTP2 setting.
func (r *ZoneSettingHTTP2Service) Edit(ctx context.Context, params ZoneSettingHTTP2EditParams, opts ...option.RequestOption) (res *ZoneSettingHTTP2EditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingHTTP2EditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http2", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Value of the HTTP2 setting.
func (r *ZoneSettingHTTP2Service) Get(ctx context.Context, query ZoneSettingHTTP2GetParams, opts ...option.RequestOption) (res *ZoneSettingHTTP2GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingHTTP2GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http2", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// HTTP2 enabled for this zone.
type ZoneSettingHTTP2EditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingHTTP2EditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingHTTP2EditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHTTP2EditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingHTTP2EditResponseJSON `json:"-"`
}

// zoneSettingHTTP2EditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingHTTP2EditResponse]
type zoneSettingHTTP2EditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2EditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingHTTP2EditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingHTTP2EditResponseID string

const (
	ZoneSettingHTTP2EditResponseIDHTTP2 ZoneSettingHTTP2EditResponseID = "http2"
)

// Current value of the zone setting.
type ZoneSettingHTTP2EditResponseValue string

const (
	ZoneSettingHTTP2EditResponseValueOn  ZoneSettingHTTP2EditResponseValue = "on"
	ZoneSettingHTTP2EditResponseValueOff ZoneSettingHTTP2EditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHTTP2EditResponseEditable bool

const (
	ZoneSettingHTTP2EditResponseEditableTrue  ZoneSettingHTTP2EditResponseEditable = true
	ZoneSettingHTTP2EditResponseEditableFalse ZoneSettingHTTP2EditResponseEditable = false
)

// HTTP2 enabled for this zone.
type ZoneSettingHTTP2GetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingHTTP2GetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingHTTP2GetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHTTP2GetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingHTTP2GetResponseJSON `json:"-"`
}

// zoneSettingHTTP2GetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingHTTP2GetResponse]
type zoneSettingHTTP2GetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingHTTP2GetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingHTTP2GetResponseID string

const (
	ZoneSettingHTTP2GetResponseIDHTTP2 ZoneSettingHTTP2GetResponseID = "http2"
)

// Current value of the zone setting.
type ZoneSettingHTTP2GetResponseValue string

const (
	ZoneSettingHTTP2GetResponseValueOn  ZoneSettingHTTP2GetResponseValue = "on"
	ZoneSettingHTTP2GetResponseValueOff ZoneSettingHTTP2GetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHTTP2GetResponseEditable bool

const (
	ZoneSettingHTTP2GetResponseEditableTrue  ZoneSettingHTTP2GetResponseEditable = true
	ZoneSettingHTTP2GetResponseEditableFalse ZoneSettingHTTP2GetResponseEditable = false
)

type ZoneSettingHTTP2EditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the HTTP2 setting.
	Value param.Field[ZoneSettingHTTP2EditParamsValue] `json:"value,required"`
}

func (r ZoneSettingHTTP2EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the HTTP2 setting.
type ZoneSettingHTTP2EditParamsValue string

const (
	ZoneSettingHTTP2EditParamsValueOn  ZoneSettingHTTP2EditParamsValue = "on"
	ZoneSettingHTTP2EditParamsValueOff ZoneSettingHTTP2EditParamsValue = "off"
)

type ZoneSettingHTTP2EditResponseEnvelope struct {
	Errors   []ZoneSettingHTTP2EditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingHTTP2EditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP2 enabled for this zone.
	Result ZoneSettingHTTP2EditResponse             `json:"result"`
	JSON   zoneSettingHTTP2EditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingHTTP2EditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP2EditResponseEnvelope]
type zoneSettingHTTP2EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingHTTP2EditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingHTTP2EditResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingHTTP2EditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingHTTP2EditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingHTTP2EditResponseEnvelopeErrors]
type zoneSettingHTTP2EditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2EditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingHTTP2EditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingHTTP2EditResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingHTTP2EditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingHTTP2EditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingHTTP2EditResponseEnvelopeMessages]
type zoneSettingHTTP2EditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2EditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingHTTP2EditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingHTTP2GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingHTTP2GetResponseEnvelope struct {
	Errors   []ZoneSettingHTTP2GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingHTTP2GetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP2 enabled for this zone.
	Result ZoneSettingHTTP2GetResponse             `json:"result"`
	JSON   zoneSettingHTTP2GetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingHTTP2GetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP2GetResponseEnvelope]
type zoneSettingHTTP2GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingHTTP2GetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingHTTP2GetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingHTTP2GetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingHTTP2GetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP2GetResponseEnvelopeErrors]
type zoneSettingHTTP2GetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2GetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingHTTP2GetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingHTTP2GetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingHTTP2GetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingHTTP2GetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingHTTP2GetResponseEnvelopeMessages]
type zoneSettingHTTP2GetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2GetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingHTTP2GetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
