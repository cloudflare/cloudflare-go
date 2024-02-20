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

// SettingAlwaysOnlineService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingAlwaysOnlineService]
// method instead.
type SettingAlwaysOnlineService struct {
	Options []option.RequestOption
}

// NewSettingAlwaysOnlineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingAlwaysOnlineService(opts ...option.RequestOption) (r *SettingAlwaysOnlineService) {
	r = &SettingAlwaysOnlineService{}
	r.Options = opts
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
func (r *SettingAlwaysOnlineService) Edit(ctx context.Context, zoneID string, body SettingAlwaysOnlineEditParams, opts ...option.RequestOption) (res *SettingAlwaysOnlineEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAlwaysOnlineEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_online", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
func (r *SettingAlwaysOnlineService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingAlwaysOnlineGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAlwaysOnlineGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_online", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type SettingAlwaysOnlineEditResponse struct {
	// ID of the zone setting.
	ID SettingAlwaysOnlineEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingAlwaysOnlineEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingAlwaysOnlineEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingAlwaysOnlineEditResponseJSON `json:"-"`
}

// settingAlwaysOnlineEditResponseJSON contains the JSON metadata for the struct
// [SettingAlwaysOnlineEditResponse]
type settingAlwaysOnlineEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingAlwaysOnlineEditResponseID string

const (
	SettingAlwaysOnlineEditResponseIDAlwaysOnline SettingAlwaysOnlineEditResponseID = "always_online"
)

// Current value of the zone setting.
type SettingAlwaysOnlineEditResponseValue string

const (
	SettingAlwaysOnlineEditResponseValueOn  SettingAlwaysOnlineEditResponseValue = "on"
	SettingAlwaysOnlineEditResponseValueOff SettingAlwaysOnlineEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingAlwaysOnlineEditResponseEditable bool

const (
	SettingAlwaysOnlineEditResponseEditableTrue  SettingAlwaysOnlineEditResponseEditable = true
	SettingAlwaysOnlineEditResponseEditableFalse SettingAlwaysOnlineEditResponseEditable = false
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type SettingAlwaysOnlineGetResponse struct {
	// ID of the zone setting.
	ID SettingAlwaysOnlineGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingAlwaysOnlineGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingAlwaysOnlineGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingAlwaysOnlineGetResponseJSON `json:"-"`
}

// settingAlwaysOnlineGetResponseJSON contains the JSON metadata for the struct
// [SettingAlwaysOnlineGetResponse]
type settingAlwaysOnlineGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingAlwaysOnlineGetResponseID string

const (
	SettingAlwaysOnlineGetResponseIDAlwaysOnline SettingAlwaysOnlineGetResponseID = "always_online"
)

// Current value of the zone setting.
type SettingAlwaysOnlineGetResponseValue string

const (
	SettingAlwaysOnlineGetResponseValueOn  SettingAlwaysOnlineGetResponseValue = "on"
	SettingAlwaysOnlineGetResponseValueOff SettingAlwaysOnlineGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingAlwaysOnlineGetResponseEditable bool

const (
	SettingAlwaysOnlineGetResponseEditableTrue  SettingAlwaysOnlineGetResponseEditable = true
	SettingAlwaysOnlineGetResponseEditableFalse SettingAlwaysOnlineGetResponseEditable = false
)

type SettingAlwaysOnlineEditParams struct {
	// Value of the zone setting.
	Value param.Field[SettingAlwaysOnlineEditParamsValue] `json:"value,required"`
}

func (r SettingAlwaysOnlineEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingAlwaysOnlineEditParamsValue string

const (
	SettingAlwaysOnlineEditParamsValueOn  SettingAlwaysOnlineEditParamsValue = "on"
	SettingAlwaysOnlineEditParamsValueOff SettingAlwaysOnlineEditParamsValue = "off"
)

type SettingAlwaysOnlineEditResponseEnvelope struct {
	Errors   []SettingAlwaysOnlineEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAlwaysOnlineEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result SettingAlwaysOnlineEditResponse             `json:"result"`
	JSON   settingAlwaysOnlineEditResponseEnvelopeJSON `json:"-"`
}

// settingAlwaysOnlineEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAlwaysOnlineEditResponseEnvelope]
type settingAlwaysOnlineEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysOnlineEditResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingAlwaysOnlineEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAlwaysOnlineEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingAlwaysOnlineEditResponseEnvelopeErrors]
type settingAlwaysOnlineEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysOnlineEditResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingAlwaysOnlineEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAlwaysOnlineEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAlwaysOnlineEditResponseEnvelopeMessages]
type settingAlwaysOnlineEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysOnlineGetResponseEnvelope struct {
	Errors   []SettingAlwaysOnlineGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAlwaysOnlineGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result SettingAlwaysOnlineGetResponse             `json:"result"`
	JSON   settingAlwaysOnlineGetResponseEnvelopeJSON `json:"-"`
}

// settingAlwaysOnlineGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAlwaysOnlineGetResponseEnvelope]
type settingAlwaysOnlineGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysOnlineGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingAlwaysOnlineGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAlwaysOnlineGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingAlwaysOnlineGetResponseEnvelopeErrors]
type settingAlwaysOnlineGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysOnlineGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingAlwaysOnlineGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAlwaysOnlineGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAlwaysOnlineGetResponseEnvelopeMessages]
type settingAlwaysOnlineGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
