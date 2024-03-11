// File generated from our OpenAPI spec by Stainless.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingIPV6Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingIPV6Service] method
// instead.
type SettingIPV6Service struct {
	Options []option.RequestOption
}

// NewSettingIPV6Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingIPV6Service(opts ...option.RequestOption) (r *SettingIPV6Service) {
	r = &SettingIPV6Service{}
	r.Options = opts
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
func (r *SettingIPV6Service) Edit(ctx context.Context, params SettingIPV6EditParams, opts ...option.RequestOption) (res *ZonesIPV6, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingIPV6EditResponseEnvelope
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
func (r *SettingIPV6Service) Get(ctx context.Context, query SettingIPV6GetParams, opts ...option.RequestOption) (res *ZonesIPV6, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingIPV6GetResponseEnvelope
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

func (r zonesIPV6JSON) RawJSON() string {
	return r.raw
}

func (r ZonesIPV6) implementsZonesSettingEditResponse() {}

func (r ZonesIPV6) implementsZonesSettingGetResponse() {}

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

func (r ZonesIPV6Param) implementsZonesSettingEditParamsItem() {}

type SettingIPV6EditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingIPV6EditParamsValue] `json:"value,required"`
}

func (r SettingIPV6EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingIPV6EditParamsValue string

const (
	SettingIPV6EditParamsValueOff SettingIPV6EditParamsValue = "off"
	SettingIPV6EditParamsValueOn  SettingIPV6EditParamsValue = "on"
)

type SettingIPV6EditResponseEnvelope struct {
	Errors   []SettingIPV6EditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingIPV6EditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IPv6 on all subdomains that are Cloudflare enabled.
	// (https://support.cloudflare.com/hc/en-us/articles/200168586).
	Result ZonesIPV6                           `json:"result"`
	JSON   settingIPV6EditResponseEnvelopeJSON `json:"-"`
}

// settingIPV6EditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingIPV6EditResponseEnvelope]
type settingIPV6EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingIPV6EditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingIPV6EditResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    settingIPV6EditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingIPV6EditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingIPV6EditResponseEnvelopeErrors]
type settingIPV6EditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6EditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingIPV6EditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingIPV6EditResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingIPV6EditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingIPV6EditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingIPV6EditResponseEnvelopeMessages]
type settingIPV6EditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6EditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingIPV6EditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingIPV6GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingIPV6GetResponseEnvelope struct {
	Errors   []SettingIPV6GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingIPV6GetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IPv6 on all subdomains that are Cloudflare enabled.
	// (https://support.cloudflare.com/hc/en-us/articles/200168586).
	Result ZonesIPV6                          `json:"result"`
	JSON   settingIPV6GetResponseEnvelopeJSON `json:"-"`
}

// settingIPV6GetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingIPV6GetResponseEnvelope]
type settingIPV6GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingIPV6GetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingIPV6GetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    settingIPV6GetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingIPV6GetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingIPV6GetResponseEnvelopeErrors]
type settingIPV6GetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6GetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingIPV6GetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingIPV6GetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingIPV6GetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingIPV6GetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingIPV6GetResponseEnvelopeMessages]
type settingIPV6GetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6GetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingIPV6GetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
