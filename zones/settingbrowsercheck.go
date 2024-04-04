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

// SettingBrowserCheckService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingBrowserCheckService]
// method instead.
type SettingBrowserCheckService struct {
	Options []option.RequestOption
}

// NewSettingBrowserCheckService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingBrowserCheckService(opts ...option.RequestOption) (r *SettingBrowserCheckService) {
	r = &SettingBrowserCheckService{}
	r.Options = opts
	return
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
func (r *SettingBrowserCheckService) Edit(ctx context.Context, params SettingBrowserCheckEditParams, opts ...option.RequestOption) (res *ZoneSettingBrowserCheck, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrowserCheckEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/browser_check", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
func (r *SettingBrowserCheckService) Get(ctx context.Context, query SettingBrowserCheckGetParams, opts ...option.RequestOption) (res *ZoneSettingBrowserCheck, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrowserCheckGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/browser_check", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingBrowserCheck struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCheckID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingBrowserCheckValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingBrowserCheckJSON `json:"-"`
}

// zoneSettingBrowserCheckJSON contains the JSON metadata for the struct
// [ZoneSettingBrowserCheck]
type zoneSettingBrowserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBrowserCheckJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingBrowserCheck) implementsZonesSettingEditResponse() {}

func (r ZoneSettingBrowserCheck) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingBrowserCheckID string

const (
	ZoneSettingBrowserCheckIDBrowserCheck ZoneSettingBrowserCheckID = "browser_check"
)

func (r ZoneSettingBrowserCheckID) IsKnown() bool {
	switch r {
	case ZoneSettingBrowserCheckIDBrowserCheck:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingBrowserCheckValue string

const (
	ZoneSettingBrowserCheckValueOn  ZoneSettingBrowserCheckValue = "on"
	ZoneSettingBrowserCheckValueOff ZoneSettingBrowserCheckValue = "off"
)

func (r ZoneSettingBrowserCheckValue) IsKnown() bool {
	switch r {
	case ZoneSettingBrowserCheckValueOn, ZoneSettingBrowserCheckValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCheckEditable bool

const (
	ZoneSettingBrowserCheckEditableTrue  ZoneSettingBrowserCheckEditable = true
	ZoneSettingBrowserCheckEditableFalse ZoneSettingBrowserCheckEditable = false
)

func (r ZoneSettingBrowserCheckEditable) IsKnown() bool {
	switch r {
	case ZoneSettingBrowserCheckEditableTrue, ZoneSettingBrowserCheckEditableFalse:
		return true
	}
	return false
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingBrowserCheckParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingBrowserCheckID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingBrowserCheckValue] `json:"value,required"`
}

func (r ZoneSettingBrowserCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingBrowserCheckParam) implementsZonesSettingEditParamsItemUnion() {}

type SettingBrowserCheckEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingBrowserCheckEditParamsValue] `json:"value,required"`
}

func (r SettingBrowserCheckEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingBrowserCheckEditParamsValue string

const (
	SettingBrowserCheckEditParamsValueOn  SettingBrowserCheckEditParamsValue = "on"
	SettingBrowserCheckEditParamsValueOff SettingBrowserCheckEditParamsValue = "off"
)

func (r SettingBrowserCheckEditParamsValue) IsKnown() bool {
	switch r {
	case SettingBrowserCheckEditParamsValueOn, SettingBrowserCheckEditParamsValueOff:
		return true
	}
	return false
}

type SettingBrowserCheckEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
	// headers abused most commonly by spammers and denies access to your page. It will
	// also challenge visitors that do not have a user agent or a non standard user
	// agent (also commonly used by abuse bots, crawlers or visitors).
	// (https://support.cloudflare.com/hc/en-us/articles/200170086).
	Result ZoneSettingBrowserCheck                     `json:"result"`
	JSON   settingBrowserCheckEditResponseEnvelopeJSON `json:"-"`
}

// settingBrowserCheckEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingBrowserCheckEditResponseEnvelope]
type settingBrowserCheckEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCheckEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCheckEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCheckGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingBrowserCheckGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
	// headers abused most commonly by spammers and denies access to your page. It will
	// also challenge visitors that do not have a user agent or a non standard user
	// agent (also commonly used by abuse bots, crawlers or visitors).
	// (https://support.cloudflare.com/hc/en-us/articles/200170086).
	Result ZoneSettingBrowserCheck                    `json:"result"`
	JSON   settingBrowserCheckGetResponseEnvelopeJSON `json:"-"`
}

// settingBrowserCheckGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingBrowserCheckGetResponseEnvelope]
type settingBrowserCheckGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCheckGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCheckGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
