// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *SettingIPV6Service) Edit(ctx context.Context, params SettingIPV6EditParams, opts ...option.RequestOption) (res *ZoneSettingIPV6, err error) {
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
func (r *SettingIPV6Service) Get(ctx context.Context, query SettingIPV6GetParams, opts ...option.RequestOption) (res *ZoneSettingIPV6, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingIPV6GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ipv6", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZoneSettingIPV6 struct {
	// ID of the zone setting.
	ID ZoneSettingIPV6ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingIPV6Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingIPV6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingIPV6JSON `json:"-"`
}

// zoneSettingIPV6JSON contains the JSON metadata for the struct [ZoneSettingIPV6]
type zoneSettingIPV6JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingIPV6JSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingIPV6) implementsZonesSettingEditResponse() {}

func (r ZoneSettingIPV6) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingIPV6ID string

const (
	ZoneSettingIPV6IDIPV6 ZoneSettingIPV6ID = "ipv6"
)

func (r ZoneSettingIPV6ID) IsKnown() bool {
	switch r {
	case ZoneSettingIPV6IDIPV6:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingIPV6Value string

const (
	ZoneSettingIPV6ValueOff ZoneSettingIPV6Value = "off"
	ZoneSettingIPV6ValueOn  ZoneSettingIPV6Value = "on"
)

func (r ZoneSettingIPV6Value) IsKnown() bool {
	switch r {
	case ZoneSettingIPV6ValueOff, ZoneSettingIPV6ValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingIPV6Editable bool

const (
	ZoneSettingIPV6EditableTrue  ZoneSettingIPV6Editable = true
	ZoneSettingIPV6EditableFalse ZoneSettingIPV6Editable = false
)

func (r ZoneSettingIPV6Editable) IsKnown() bool {
	switch r {
	case ZoneSettingIPV6EditableTrue, ZoneSettingIPV6EditableFalse:
		return true
	}
	return false
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZoneSettingIPV6Param struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingIPV6ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingIPV6Value] `json:"value,required"`
}

func (r ZoneSettingIPV6Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingIPV6Param) implementsZonesSettingEditParamsItem() {}

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

func (r SettingIPV6EditParamsValue) IsKnown() bool {
	switch r {
	case SettingIPV6EditParamsValueOff, SettingIPV6EditParamsValueOn:
		return true
	}
	return false
}

type SettingIPV6EditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IPv6 on all subdomains that are Cloudflare enabled.
	// (https://support.cloudflare.com/hc/en-us/articles/200168586).
	Result ZoneSettingIPV6                     `json:"result"`
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

type SettingIPV6GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingIPV6GetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IPv6 on all subdomains that are Cloudflare enabled.
	// (https://support.cloudflare.com/hc/en-us/articles/200168586).
	Result ZoneSettingIPV6                    `json:"result"`
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
