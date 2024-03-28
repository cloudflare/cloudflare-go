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
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingAlwaysUseHTTPSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingAlwaysUseHTTPSService]
// method instead.
type SettingAlwaysUseHTTPSService struct {
	Options []option.RequestOption
}

// NewSettingAlwaysUseHTTPSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingAlwaysUseHTTPSService(opts ...option.RequestOption) (r *SettingAlwaysUseHTTPSService) {
	r = &SettingAlwaysUseHTTPSService{}
	r.Options = opts
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
func (r *SettingAlwaysUseHTTPSService) Edit(ctx context.Context, params SettingAlwaysUseHTTPSEditParams, opts ...option.RequestOption) (res *ZoneSettingAlwaysUseHTTPS, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAlwaysUseHTTPSEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_use_https", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
func (r *SettingAlwaysUseHTTPSService) Get(ctx context.Context, query SettingAlwaysUseHTTPSGetParams, opts ...option.RequestOption) (res *ZoneSettingAlwaysUseHTTPS, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAlwaysUseHTTPSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_use_https", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZoneSettingAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID ZoneSettingAlwaysUseHTTPSID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingAlwaysUseHTTPSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAlwaysUseHTTPSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingAlwaysUseHTTPSJSON `json:"-"`
}

// zoneSettingAlwaysUseHTTPSJSON contains the JSON metadata for the struct
// [ZoneSettingAlwaysUseHTTPS]
type zoneSettingAlwaysUseHTTPSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAlwaysUseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingAlwaysUseHTTPS) implementsZonesSettingEditResponse() {}

func (r ZoneSettingAlwaysUseHTTPS) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingAlwaysUseHTTPSID string

const (
	ZoneSettingAlwaysUseHTTPSIDAlwaysUseHTTPS ZoneSettingAlwaysUseHTTPSID = "always_use_https"
)

func (r ZoneSettingAlwaysUseHTTPSID) IsKnown() bool {
	switch r {
	case ZoneSettingAlwaysUseHTTPSIDAlwaysUseHTTPS:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingAlwaysUseHTTPSValue string

const (
	ZoneSettingAlwaysUseHTTPSValueOn  ZoneSettingAlwaysUseHTTPSValue = "on"
	ZoneSettingAlwaysUseHTTPSValueOff ZoneSettingAlwaysUseHTTPSValue = "off"
)

func (r ZoneSettingAlwaysUseHTTPSValue) IsKnown() bool {
	switch r {
	case ZoneSettingAlwaysUseHTTPSValueOn, ZoneSettingAlwaysUseHTTPSValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAlwaysUseHTTPSEditable bool

const (
	ZoneSettingAlwaysUseHTTPSEditableTrue  ZoneSettingAlwaysUseHTTPSEditable = true
	ZoneSettingAlwaysUseHTTPSEditableFalse ZoneSettingAlwaysUseHTTPSEditable = false
)

func (r ZoneSettingAlwaysUseHTTPSEditable) IsKnown() bool {
	switch r {
	case ZoneSettingAlwaysUseHTTPSEditableTrue, ZoneSettingAlwaysUseHTTPSEditableFalse:
		return true
	}
	return false
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZoneSettingAlwaysUseHTTPSParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingAlwaysUseHTTPSID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingAlwaysUseHTTPSValue] `json:"value,required"`
}

func (r ZoneSettingAlwaysUseHTTPSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingAlwaysUseHTTPSParam) implementsZonesSettingEditParamsItem() {}

type SettingAlwaysUseHTTPSEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingAlwaysUseHTTPSEditParamsValue] `json:"value,required"`
}

func (r SettingAlwaysUseHTTPSEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingAlwaysUseHTTPSEditParamsValue string

const (
	SettingAlwaysUseHTTPSEditParamsValueOn  SettingAlwaysUseHTTPSEditParamsValue = "on"
	SettingAlwaysUseHTTPSEditParamsValueOff SettingAlwaysUseHTTPSEditParamsValue = "off"
)

func (r SettingAlwaysUseHTTPSEditParamsValue) IsKnown() bool {
	switch r {
	case SettingAlwaysUseHTTPSEditParamsValueOn, SettingAlwaysUseHTTPSEditParamsValueOff:
		return true
	}
	return false
}

type SettingAlwaysUseHTTPSEditResponseEnvelope struct {
	Errors   []SettingAlwaysUseHTTPSEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAlwaysUseHTTPSEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result ZoneSettingAlwaysUseHTTPS                     `json:"result"`
	JSON   settingAlwaysUseHTTPSEditResponseEnvelopeJSON `json:"-"`
}

// settingAlwaysUseHTTPSEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAlwaysUseHTTPSEditResponseEnvelope]
type settingAlwaysUseHTTPSEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysUseHTTPSEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAlwaysUseHTTPSEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingAlwaysUseHTTPSEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAlwaysUseHTTPSEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingAlwaysUseHTTPSEditResponseEnvelopeErrors]
type settingAlwaysUseHTTPSEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysUseHTTPSEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAlwaysUseHTTPSEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingAlwaysUseHTTPSEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAlwaysUseHTTPSEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAlwaysUseHTTPSEditResponseEnvelopeMessages]
type settingAlwaysUseHTTPSEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysUseHTTPSEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAlwaysUseHTTPSGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingAlwaysUseHTTPSGetResponseEnvelope struct {
	Errors   []SettingAlwaysUseHTTPSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAlwaysUseHTTPSGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result ZoneSettingAlwaysUseHTTPS                    `json:"result"`
	JSON   settingAlwaysUseHTTPSGetResponseEnvelopeJSON `json:"-"`
}

// settingAlwaysUseHTTPSGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAlwaysUseHTTPSGetResponseEnvelope]
type settingAlwaysUseHTTPSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysUseHTTPSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAlwaysUseHTTPSGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingAlwaysUseHTTPSGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAlwaysUseHTTPSGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingAlwaysUseHTTPSGetResponseEnvelopeErrors]
type settingAlwaysUseHTTPSGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysUseHTTPSGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAlwaysUseHTTPSGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingAlwaysUseHTTPSGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAlwaysUseHTTPSGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAlwaysUseHTTPSGetResponseEnvelopeMessages]
type settingAlwaysUseHTTPSGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysUseHTTPSGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
