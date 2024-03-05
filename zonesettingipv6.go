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
func (r *ZoneSettingIPV6Service) Edit(ctx context.Context, params ZoneSettingIPV6EditParams, opts ...option.RequestOption) (res *ZonesIPV6, err error) {
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
func (r *ZoneSettingIPV6Service) Get(ctx context.Context, query ZoneSettingIPV6GetParams, opts ...option.RequestOption) (res *ZonesIPV6, err error) {
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
type ZonesIPV6 struct {
	// ID of the zone setting.
	ID ZonesIPV6ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesIPV6Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesIPV6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time     `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesIPV6JSON `json:"-"`
}

// zonesIPV6JSON contains the JSON metadata for the struct [ZonesIPV6]
type zonesIPV6JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesIPV6) implementsZoneSettingEditResponse() {}

func (r ZonesIPV6) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesIPV6ID string

const (
	ZonesIPV6IDIPV6 ZonesIPV6ID = "ipv6"
)

// Current value of the zone setting.
type ZonesIPV6Value string

const (
	ZonesIPV6ValueOff ZonesIPV6Value = "off"
	ZonesIPV6ValueOn  ZonesIPV6Value = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesIPV6Editable bool

const (
	ZonesIPV6EditableTrue  ZonesIPV6Editable = true
	ZonesIPV6EditableFalse ZonesIPV6Editable = false
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZonesIPV6Param struct {
	// ID of the zone setting.
	ID param.Field[ZonesIPV6ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesIPV6Value] `json:"value,required"`
}

func (r ZonesIPV6Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesIPV6Param) implementsZoneSettingEditParamsItem() {}

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
	Result ZonesIPV6                               `json:"result"`
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
	Result ZonesIPV6                              `json:"result"`
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
