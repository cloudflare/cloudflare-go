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
func (r *ZoneSettingBrowserCheckService) Edit(ctx context.Context, params ZoneSettingBrowserCheckEditParams, opts ...option.RequestOption) (res *ZoneSettingBrowserCheckEditResponse, err error) {
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
func (r *ZoneSettingBrowserCheckService) Get(ctx context.Context, query ZoneSettingBrowserCheckGetParams, opts ...option.RequestOption) (res *ZoneSettingBrowserCheckGetResponse, err error) {
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
type ZoneSettingBrowserCheckEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCheckEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingBrowserCheckEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCheckEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingBrowserCheckEditResponseJSON `json:"-"`
}

// zoneSettingBrowserCheckEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingBrowserCheckEditResponse]
type zoneSettingBrowserCheckEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBrowserCheckEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingBrowserCheckEditResponseID string

const (
	ZoneSettingBrowserCheckEditResponseIDBrowserCheck ZoneSettingBrowserCheckEditResponseID = "browser_check"
)

// Current value of the zone setting.
type ZoneSettingBrowserCheckEditResponseValue string

const (
	ZoneSettingBrowserCheckEditResponseValueOn  ZoneSettingBrowserCheckEditResponseValue = "on"
	ZoneSettingBrowserCheckEditResponseValueOff ZoneSettingBrowserCheckEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCheckEditResponseEditable bool

const (
	ZoneSettingBrowserCheckEditResponseEditableTrue  ZoneSettingBrowserCheckEditResponseEditable = true
	ZoneSettingBrowserCheckEditResponseEditableFalse ZoneSettingBrowserCheckEditResponseEditable = false
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingBrowserCheckGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCheckGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingBrowserCheckGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCheckGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingBrowserCheckGetResponseJSON `json:"-"`
}

// zoneSettingBrowserCheckGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingBrowserCheckGetResponse]
type zoneSettingBrowserCheckGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBrowserCheckGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingBrowserCheckGetResponseID string

const (
	ZoneSettingBrowserCheckGetResponseIDBrowserCheck ZoneSettingBrowserCheckGetResponseID = "browser_check"
)

// Current value of the zone setting.
type ZoneSettingBrowserCheckGetResponseValue string

const (
	ZoneSettingBrowserCheckGetResponseValueOn  ZoneSettingBrowserCheckGetResponseValue = "on"
	ZoneSettingBrowserCheckGetResponseValueOff ZoneSettingBrowserCheckGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCheckGetResponseEditable bool

const (
	ZoneSettingBrowserCheckGetResponseEditableTrue  ZoneSettingBrowserCheckGetResponseEditable = true
	ZoneSettingBrowserCheckGetResponseEditableFalse ZoneSettingBrowserCheckGetResponseEditable = false
)

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
	Result ZoneSettingBrowserCheckEditResponse             `json:"result"`
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

func (r zoneSettingBrowserCheckEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingBrowserCheckEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingBrowserCheckEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result ZoneSettingBrowserCheckGetResponse             `json:"result"`
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

func (r zoneSettingBrowserCheckGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingBrowserCheckGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingBrowserCheckGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
