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

// ZoneSettingBrowserCheckService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingBrowserCheckService] method instead.
type ZoneSettingBrowserCheckService struct {
	Options []option.RequestOption
}

// NewZoneSettingBrowserCheckService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingBrowserCheckService(opts ...option.RequestOption) (r *ZoneSettingBrowserCheckService) {
	r = &ZoneSettingBrowserCheckService{}
	r.Options = opts
	return
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
func (r *ZoneSettingBrowserCheckService) Edit(ctx context.Context, params ZoneSettingBrowserCheckEditParams, opts ...option.RequestOption) (res *ZonesBrowserCheck, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingBrowserCheckEditResponseEnvelope
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
func (r *ZoneSettingBrowserCheckService) Get(ctx context.Context, query ZoneSettingBrowserCheckGetParams, opts ...option.RequestOption) (res *ZonesBrowserCheck, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingBrowserCheckGetResponseEnvelope
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

func (r ZonesBrowserCheck) implementsZoneSettingEditResponse() {}

func (r ZonesBrowserCheck) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesBrowserCheckID string

const (
	ZonesBrowserCheckIDBrowserCheck ZonesBrowserCheckID = "browser_check"
)

// Current value of the zone setting.
type ZonesBrowserCheckValue string

const (
	ZonesBrowserCheckValueOn  ZonesBrowserCheckValue = "on"
	ZonesBrowserCheckValueOff ZonesBrowserCheckValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesBrowserCheckEditable bool

const (
	ZonesBrowserCheckEditableTrue  ZonesBrowserCheckEditable = true
	ZonesBrowserCheckEditableFalse ZonesBrowserCheckEditable = false
)

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

func (r ZonesBrowserCheckParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingBrowserCheckEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingBrowserCheckEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingBrowserCheckEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingBrowserCheckEditParamsValue string

const (
	ZoneSettingBrowserCheckEditParamsValueOn  ZoneSettingBrowserCheckEditParamsValue = "on"
	ZoneSettingBrowserCheckEditParamsValueOff ZoneSettingBrowserCheckEditParamsValue = "off"
)

type ZoneSettingBrowserCheckEditResponseEnvelope struct {
	Errors   []ZoneSettingBrowserCheckEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingBrowserCheckEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
	// headers abused most commonly by spammers and denies access to your page. It will
	// also challenge visitors that do not have a user agent or a non standard user
	// agent (also commonly used by abuse bots, crawlers or visitors).
	// (https://support.cloudflare.com/hc/en-us/articles/200170086).
	Result ZonesBrowserCheck                               `json:"result"`
	JSON   zoneSettingBrowserCheckEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingBrowserCheckEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCheckEditResponseEnvelope]
type zoneSettingBrowserCheckEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCheckEditResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneSettingBrowserCheckEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingBrowserCheckEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingBrowserCheckEditResponseEnvelopeErrors]
type zoneSettingBrowserCheckEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCheckEditResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingBrowserCheckEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingBrowserCheckEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingBrowserCheckEditResponseEnvelopeMessages]
type zoneSettingBrowserCheckEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCheckGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingBrowserCheckGetResponseEnvelope struct {
	Errors   []ZoneSettingBrowserCheckGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingBrowserCheckGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
	// headers abused most commonly by spammers and denies access to your page. It will
	// also challenge visitors that do not have a user agent or a non standard user
	// agent (also commonly used by abuse bots, crawlers or visitors).
	// (https://support.cloudflare.com/hc/en-us/articles/200170086).
	Result ZonesBrowserCheck                              `json:"result"`
	JSON   zoneSettingBrowserCheckGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingBrowserCheckGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCheckGetResponseEnvelope]
type zoneSettingBrowserCheckGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCheckGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingBrowserCheckGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingBrowserCheckGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingBrowserCheckGetResponseEnvelopeErrors]
type zoneSettingBrowserCheckGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCheckGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingBrowserCheckGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingBrowserCheckGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingBrowserCheckGetResponseEnvelopeMessages]
type zoneSettingBrowserCheckGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
