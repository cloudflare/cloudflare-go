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
func (r *SettingBrowserCheckService) Edit(ctx context.Context, params SettingBrowserCheckEditParams, opts ...option.RequestOption) (res *ZonesBrowserCheck, err error) {
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
func (r *SettingBrowserCheckService) Get(ctx context.Context, query SettingBrowserCheckGetParams, opts ...option.RequestOption) (res *ZonesBrowserCheck, err error) {
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
type ZonesBrowserCheck struct {
	// ID of the zone setting.
	ID ZonesBrowserCheckID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesBrowserCheckValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesBrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time             `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesBrowserCheckJSON `json:"-"`
}

// zonesBrowserCheckJSON contains the JSON metadata for the struct
// [ZonesBrowserCheck]
type zonesBrowserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesBrowserCheckJSON) RawJSON() string {
	return r.raw
}

func (r ZonesBrowserCheck) implementsZonesSettingEditResponse() {}

func (r ZonesBrowserCheck) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesBrowserCheckID string

const (
	ZonesBrowserCheckIDBrowserCheck ZonesBrowserCheckID = "browser_check"
)

func (r ZonesBrowserCheckID) IsKnown() bool {
	switch r {
	case ZonesBrowserCheckIDBrowserCheck:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZonesBrowserCheckValue string

const (
	ZonesBrowserCheckValueOn  ZonesBrowserCheckValue = "on"
	ZonesBrowserCheckValueOff ZonesBrowserCheckValue = "off"
)

func (r ZonesBrowserCheckValue) IsKnown() bool {
	switch r {
	case ZonesBrowserCheckValueOn, ZonesBrowserCheckValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesBrowserCheckEditable bool

const (
	ZonesBrowserCheckEditableTrue  ZonesBrowserCheckEditable = true
	ZonesBrowserCheckEditableFalse ZonesBrowserCheckEditable = false
)

func (r ZonesBrowserCheckEditable) IsKnown() bool {
	switch r {
	case ZonesBrowserCheckEditableTrue, ZonesBrowserCheckEditableFalse:
		return true
	}
	return false
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZonesBrowserCheckParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesBrowserCheckID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesBrowserCheckValue] `json:"value,required"`
}

func (r ZonesBrowserCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesBrowserCheckParam) implementsZonesSettingEditParamsItem() {}

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
	Errors   []SettingBrowserCheckEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingBrowserCheckEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
	// headers abused most commonly by spammers and denies access to your page. It will
	// also challenge visitors that do not have a user agent or a non standard user
	// agent (also commonly used by abuse bots, crawlers or visitors).
	// (https://support.cloudflare.com/hc/en-us/articles/200170086).
	Result ZonesBrowserCheck                           `json:"result"`
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

type SettingBrowserCheckEditResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingBrowserCheckEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingBrowserCheckEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingBrowserCheckEditResponseEnvelopeErrors]
type settingBrowserCheckEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCheckEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCheckEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCheckEditResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingBrowserCheckEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingBrowserCheckEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingBrowserCheckEditResponseEnvelopeMessages]
type settingBrowserCheckEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCheckEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCheckEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCheckGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingBrowserCheckGetResponseEnvelope struct {
	Errors   []SettingBrowserCheckGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingBrowserCheckGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
	// headers abused most commonly by spammers and denies access to your page. It will
	// also challenge visitors that do not have a user agent or a non standard user
	// agent (also commonly used by abuse bots, crawlers or visitors).
	// (https://support.cloudflare.com/hc/en-us/articles/200170086).
	Result ZonesBrowserCheck                          `json:"result"`
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

type SettingBrowserCheckGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingBrowserCheckGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingBrowserCheckGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingBrowserCheckGetResponseEnvelopeErrors]
type settingBrowserCheckGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCheckGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCheckGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCheckGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingBrowserCheckGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingBrowserCheckGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingBrowserCheckGetResponseEnvelopeMessages]
type settingBrowserCheckGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCheckGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCheckGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
