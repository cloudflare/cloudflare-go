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

// ZoneSettingIPV6Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingIPV6Service] method
// instead.
type ZoneSettingIPV6Service struct {
	Options []option.RequestOption
}

// NewZoneSettingIPV6Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingIPV6Service(opts ...option.RequestOption) (r *ZoneSettingIPV6Service) {
	r = &ZoneSettingIPV6Service{}
	r.Options = opts
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
func (r *ZoneSettingIPV6Service) Edit(ctx context.Context, params ZoneSettingIPV6EditParams, opts ...option.RequestOption) (res *ZoneSettingIPV6EditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingIPV6EditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ipv6", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
func (r *ZoneSettingIPV6Service) Get(ctx context.Context, query ZoneSettingIPV6GetParams, opts ...option.RequestOption) (res *ZoneSettingIPV6GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingIPV6GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ipv6", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZoneSettingIPV6EditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingIPV6EditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingIPV6EditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingIPV6EditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingIPV6EditResponseJSON `json:"-"`
}

// zoneSettingIPV6EditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingIPV6EditResponse]
type zoneSettingIPV6EditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6EditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingIPV6EditResponseID string

const (
	ZoneSettingIPV6EditResponseIDIPV6 ZoneSettingIPV6EditResponseID = "ipv6"
)

// Current value of the zone setting.
type ZoneSettingIPV6EditResponseValue string

const (
	ZoneSettingIPV6EditResponseValueOff ZoneSettingIPV6EditResponseValue = "off"
	ZoneSettingIPV6EditResponseValueOn  ZoneSettingIPV6EditResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingIPV6EditResponseEditable bool

const (
	ZoneSettingIPV6EditResponseEditableTrue  ZoneSettingIPV6EditResponseEditable = true
	ZoneSettingIPV6EditResponseEditableFalse ZoneSettingIPV6EditResponseEditable = false
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZoneSettingIPV6GetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingIPV6GetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingIPV6GetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingIPV6GetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingIPV6GetResponseJSON `json:"-"`
}

// zoneSettingIPV6GetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingIPV6GetResponse]
type zoneSettingIPV6GetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingIPV6GetResponseID string

const (
	ZoneSettingIPV6GetResponseIDIPV6 ZoneSettingIPV6GetResponseID = "ipv6"
)

// Current value of the zone setting.
type ZoneSettingIPV6GetResponseValue string

const (
	ZoneSettingIPV6GetResponseValueOff ZoneSettingIPV6GetResponseValue = "off"
	ZoneSettingIPV6GetResponseValueOn  ZoneSettingIPV6GetResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingIPV6GetResponseEditable bool

const (
	ZoneSettingIPV6GetResponseEditableTrue  ZoneSettingIPV6GetResponseEditable = true
	ZoneSettingIPV6GetResponseEditableFalse ZoneSettingIPV6GetResponseEditable = false
)

type ZoneSettingIPV6EditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingIPV6EditParamsValue] `json:"value,required"`
}

func (r ZoneSettingIPV6EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingIPV6EditParamsValue string

const (
	ZoneSettingIPV6EditParamsValueOff ZoneSettingIPV6EditParamsValue = "off"
	ZoneSettingIPV6EditParamsValueOn  ZoneSettingIPV6EditParamsValue = "on"
)

type ZoneSettingIPV6EditResponseEnvelope struct {
	Errors   []ZoneSettingIPV6EditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingIPV6EditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IPv6 on all subdomains that are Cloudflare enabled.
	// (https://support.cloudflare.com/hc/en-us/articles/200168586).
	Result ZoneSettingIPV6EditResponse             `json:"result"`
	JSON   zoneSettingIPV6EditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingIPV6EditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingIPV6EditResponseEnvelope]
type zoneSettingIPV6EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPV6EditResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingIPV6EditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingIPV6EditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingIPV6EditResponseEnvelopeErrors]
type zoneSettingIPV6EditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6EditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPV6EditResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingIPV6EditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingIPV6EditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingIPV6EditResponseEnvelopeMessages]
type zoneSettingIPV6EditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6EditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPV6GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingIPV6GetResponseEnvelope struct {
	Errors   []ZoneSettingIPV6GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingIPV6GetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IPv6 on all subdomains that are Cloudflare enabled.
	// (https://support.cloudflare.com/hc/en-us/articles/200168586).
	Result ZoneSettingIPV6GetResponse             `json:"result"`
	JSON   zoneSettingIPV6GetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingIPV6GetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneSettingIPV6GetResponseEnvelope]
type zoneSettingIPV6GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPV6GetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSettingIPV6GetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingIPV6GetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingIPV6GetResponseEnvelopeErrors]
type zoneSettingIPV6GetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6GetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPV6GetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingIPV6GetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingIPV6GetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingIPV6GetResponseEnvelopeMessages]
type zoneSettingIPV6GetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6GetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
