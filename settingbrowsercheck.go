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
func (r *SettingBrowserCheckService) Edit(ctx context.Context, params SettingBrowserCheckEditParams, opts ...option.RequestOption) (res *SettingBrowserCheckEditResponse, err error) {
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
func (r *SettingBrowserCheckService) Get(ctx context.Context, query SettingBrowserCheckGetParams, opts ...option.RequestOption) (res *SettingBrowserCheckGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrowserCheckGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/browser_check", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
type SettingBrowserCheckEditResponse struct {
	// ID of the zone setting.
	ID SettingBrowserCheckEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingBrowserCheckEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingBrowserCheckEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingBrowserCheckEditResponseJSON `json:"-"`
}

// settingBrowserCheckEditResponseJSON contains the JSON metadata for the struct
// [SettingBrowserCheckEditResponse]
type settingBrowserCheckEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCheckEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingBrowserCheckEditResponseID string

const (
	SettingBrowserCheckEditResponseIDBrowserCheck SettingBrowserCheckEditResponseID = "browser_check"
)

// Current value of the zone setting.
type SettingBrowserCheckEditResponseValue string

const (
	SettingBrowserCheckEditResponseValueOn  SettingBrowserCheckEditResponseValue = "on"
	SettingBrowserCheckEditResponseValueOff SettingBrowserCheckEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingBrowserCheckEditResponseEditable bool

const (
	SettingBrowserCheckEditResponseEditableTrue  SettingBrowserCheckEditResponseEditable = true
	SettingBrowserCheckEditResponseEditableFalse SettingBrowserCheckEditResponseEditable = false
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type SettingBrowserCheckGetResponse struct {
	// ID of the zone setting.
	ID SettingBrowserCheckGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingBrowserCheckGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingBrowserCheckGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingBrowserCheckGetResponseJSON `json:"-"`
}

// settingBrowserCheckGetResponseJSON contains the JSON metadata for the struct
// [SettingBrowserCheckGetResponse]
type settingBrowserCheckGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCheckGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingBrowserCheckGetResponseID string

const (
	SettingBrowserCheckGetResponseIDBrowserCheck SettingBrowserCheckGetResponseID = "browser_check"
)

// Current value of the zone setting.
type SettingBrowserCheckGetResponseValue string

const (
	SettingBrowserCheckGetResponseValueOn  SettingBrowserCheckGetResponseValue = "on"
	SettingBrowserCheckGetResponseValueOff SettingBrowserCheckGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingBrowserCheckGetResponseEditable bool

const (
	SettingBrowserCheckGetResponseEditableTrue  SettingBrowserCheckGetResponseEditable = true
	SettingBrowserCheckGetResponseEditableFalse SettingBrowserCheckGetResponseEditable = false
)

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
	Result SettingBrowserCheckEditResponse             `json:"result"`
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
	Result SettingBrowserCheckGetResponse             `json:"result"`
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
