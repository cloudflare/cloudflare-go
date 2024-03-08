// File generated from our OpenAPI spec by Stainless.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// SettingBrowserCacheTTLService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingBrowserCacheTTLService]
// method instead.
type SettingBrowserCacheTTLService struct {
	Options []option.RequestOption
}

// NewSettingBrowserCacheTTLService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingBrowserCacheTTLService(opts ...option.RequestOption) (r *SettingBrowserCacheTTLService) {
	r = &SettingBrowserCacheTTLService{}
	r.Options = opts
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
func (r *SettingBrowserCacheTTLService) Edit(ctx context.Context, params SettingBrowserCacheTTLEditParams, opts ...option.RequestOption) (res *SettingBrowserCacheTTLEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrowserCacheTTLEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
func (r *SettingBrowserCacheTTLService) Get(ctx context.Context, query SettingBrowserCacheTTLGetParams, opts ...option.RequestOption) (res *SettingBrowserCacheTTLGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrowserCacheTTLGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type SettingBrowserCacheTTLEditResponse struct {
	// ID of the zone setting.
	ID SettingBrowserCacheTTLEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingBrowserCacheTTLEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingBrowserCacheTTLEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingBrowserCacheTTLEditResponseJSON `json:"-"`
}

// settingBrowserCacheTTLEditResponseJSON contains the JSON metadata for the struct
// [SettingBrowserCacheTTLEditResponse]
type settingBrowserCacheTTLEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingBrowserCacheTTLEditResponseID string

const (
	SettingBrowserCacheTTLEditResponseIDBrowserCacheTTL SettingBrowserCacheTTLEditResponseID = "browser_cache_ttl"
)

// Current value of the zone setting.
type SettingBrowserCacheTTLEditResponseValue float64

const (
	SettingBrowserCacheTTLEditResponseValue0        SettingBrowserCacheTTLEditResponseValue = 0
	SettingBrowserCacheTTLEditResponseValue30       SettingBrowserCacheTTLEditResponseValue = 30
	SettingBrowserCacheTTLEditResponseValue60       SettingBrowserCacheTTLEditResponseValue = 60
	SettingBrowserCacheTTLEditResponseValue120      SettingBrowserCacheTTLEditResponseValue = 120
	SettingBrowserCacheTTLEditResponseValue300      SettingBrowserCacheTTLEditResponseValue = 300
	SettingBrowserCacheTTLEditResponseValue1200     SettingBrowserCacheTTLEditResponseValue = 1200
	SettingBrowserCacheTTLEditResponseValue1800     SettingBrowserCacheTTLEditResponseValue = 1800
	SettingBrowserCacheTTLEditResponseValue3600     SettingBrowserCacheTTLEditResponseValue = 3600
	SettingBrowserCacheTTLEditResponseValue7200     SettingBrowserCacheTTLEditResponseValue = 7200
	SettingBrowserCacheTTLEditResponseValue10800    SettingBrowserCacheTTLEditResponseValue = 10800
	SettingBrowserCacheTTLEditResponseValue14400    SettingBrowserCacheTTLEditResponseValue = 14400
	SettingBrowserCacheTTLEditResponseValue18000    SettingBrowserCacheTTLEditResponseValue = 18000
	SettingBrowserCacheTTLEditResponseValue28800    SettingBrowserCacheTTLEditResponseValue = 28800
	SettingBrowserCacheTTLEditResponseValue43200    SettingBrowserCacheTTLEditResponseValue = 43200
	SettingBrowserCacheTTLEditResponseValue57600    SettingBrowserCacheTTLEditResponseValue = 57600
	SettingBrowserCacheTTLEditResponseValue72000    SettingBrowserCacheTTLEditResponseValue = 72000
	SettingBrowserCacheTTLEditResponseValue86400    SettingBrowserCacheTTLEditResponseValue = 86400
	SettingBrowserCacheTTLEditResponseValue172800   SettingBrowserCacheTTLEditResponseValue = 172800
	SettingBrowserCacheTTLEditResponseValue259200   SettingBrowserCacheTTLEditResponseValue = 259200
	SettingBrowserCacheTTLEditResponseValue345600   SettingBrowserCacheTTLEditResponseValue = 345600
	SettingBrowserCacheTTLEditResponseValue432000   SettingBrowserCacheTTLEditResponseValue = 432000
	SettingBrowserCacheTTLEditResponseValue691200   SettingBrowserCacheTTLEditResponseValue = 691200
	SettingBrowserCacheTTLEditResponseValue1382400  SettingBrowserCacheTTLEditResponseValue = 1382400
	SettingBrowserCacheTTLEditResponseValue2073600  SettingBrowserCacheTTLEditResponseValue = 2073600
	SettingBrowserCacheTTLEditResponseValue2678400  SettingBrowserCacheTTLEditResponseValue = 2678400
	SettingBrowserCacheTTLEditResponseValue5356800  SettingBrowserCacheTTLEditResponseValue = 5356800
	SettingBrowserCacheTTLEditResponseValue16070400 SettingBrowserCacheTTLEditResponseValue = 16070400
	SettingBrowserCacheTTLEditResponseValue31536000 SettingBrowserCacheTTLEditResponseValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingBrowserCacheTTLEditResponseEditable bool

const (
	SettingBrowserCacheTTLEditResponseEditableTrue  SettingBrowserCacheTTLEditResponseEditable = true
	SettingBrowserCacheTTLEditResponseEditableFalse SettingBrowserCacheTTLEditResponseEditable = false
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type SettingBrowserCacheTTLGetResponse struct {
	// ID of the zone setting.
	ID SettingBrowserCacheTTLGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingBrowserCacheTTLGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingBrowserCacheTTLGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingBrowserCacheTTLGetResponseJSON `json:"-"`
}

// settingBrowserCacheTTLGetResponseJSON contains the JSON metadata for the struct
// [SettingBrowserCacheTTLGetResponse]
type settingBrowserCacheTTLGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingBrowserCacheTTLGetResponseID string

const (
	SettingBrowserCacheTTLGetResponseIDBrowserCacheTTL SettingBrowserCacheTTLGetResponseID = "browser_cache_ttl"
)

// Current value of the zone setting.
type SettingBrowserCacheTTLGetResponseValue float64

const (
	SettingBrowserCacheTTLGetResponseValue0        SettingBrowserCacheTTLGetResponseValue = 0
	SettingBrowserCacheTTLGetResponseValue30       SettingBrowserCacheTTLGetResponseValue = 30
	SettingBrowserCacheTTLGetResponseValue60       SettingBrowserCacheTTLGetResponseValue = 60
	SettingBrowserCacheTTLGetResponseValue120      SettingBrowserCacheTTLGetResponseValue = 120
	SettingBrowserCacheTTLGetResponseValue300      SettingBrowserCacheTTLGetResponseValue = 300
	SettingBrowserCacheTTLGetResponseValue1200     SettingBrowserCacheTTLGetResponseValue = 1200
	SettingBrowserCacheTTLGetResponseValue1800     SettingBrowserCacheTTLGetResponseValue = 1800
	SettingBrowserCacheTTLGetResponseValue3600     SettingBrowserCacheTTLGetResponseValue = 3600
	SettingBrowserCacheTTLGetResponseValue7200     SettingBrowserCacheTTLGetResponseValue = 7200
	SettingBrowserCacheTTLGetResponseValue10800    SettingBrowserCacheTTLGetResponseValue = 10800
	SettingBrowserCacheTTLGetResponseValue14400    SettingBrowserCacheTTLGetResponseValue = 14400
	SettingBrowserCacheTTLGetResponseValue18000    SettingBrowserCacheTTLGetResponseValue = 18000
	SettingBrowserCacheTTLGetResponseValue28800    SettingBrowserCacheTTLGetResponseValue = 28800
	SettingBrowserCacheTTLGetResponseValue43200    SettingBrowserCacheTTLGetResponseValue = 43200
	SettingBrowserCacheTTLGetResponseValue57600    SettingBrowserCacheTTLGetResponseValue = 57600
	SettingBrowserCacheTTLGetResponseValue72000    SettingBrowserCacheTTLGetResponseValue = 72000
	SettingBrowserCacheTTLGetResponseValue86400    SettingBrowserCacheTTLGetResponseValue = 86400
	SettingBrowserCacheTTLGetResponseValue172800   SettingBrowserCacheTTLGetResponseValue = 172800
	SettingBrowserCacheTTLGetResponseValue259200   SettingBrowserCacheTTLGetResponseValue = 259200
	SettingBrowserCacheTTLGetResponseValue345600   SettingBrowserCacheTTLGetResponseValue = 345600
	SettingBrowserCacheTTLGetResponseValue432000   SettingBrowserCacheTTLGetResponseValue = 432000
	SettingBrowserCacheTTLGetResponseValue691200   SettingBrowserCacheTTLGetResponseValue = 691200
	SettingBrowserCacheTTLGetResponseValue1382400  SettingBrowserCacheTTLGetResponseValue = 1382400
	SettingBrowserCacheTTLGetResponseValue2073600  SettingBrowserCacheTTLGetResponseValue = 2073600
	SettingBrowserCacheTTLGetResponseValue2678400  SettingBrowserCacheTTLGetResponseValue = 2678400
	SettingBrowserCacheTTLGetResponseValue5356800  SettingBrowserCacheTTLGetResponseValue = 5356800
	SettingBrowserCacheTTLGetResponseValue16070400 SettingBrowserCacheTTLGetResponseValue = 16070400
	SettingBrowserCacheTTLGetResponseValue31536000 SettingBrowserCacheTTLGetResponseValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingBrowserCacheTTLGetResponseEditable bool

const (
	SettingBrowserCacheTTLGetResponseEditableTrue  SettingBrowserCacheTTLGetResponseEditable = true
	SettingBrowserCacheTTLGetResponseEditableFalse SettingBrowserCacheTTLGetResponseEditable = false
)

type SettingBrowserCacheTTLEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value param.Field[SettingBrowserCacheTTLEditParamsValue] `json:"value,required"`
}

func (r SettingBrowserCacheTTLEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type SettingBrowserCacheTTLEditParamsValue float64

const (
	SettingBrowserCacheTTLEditParamsValue0        SettingBrowserCacheTTLEditParamsValue = 0
	SettingBrowserCacheTTLEditParamsValue30       SettingBrowserCacheTTLEditParamsValue = 30
	SettingBrowserCacheTTLEditParamsValue60       SettingBrowserCacheTTLEditParamsValue = 60
	SettingBrowserCacheTTLEditParamsValue120      SettingBrowserCacheTTLEditParamsValue = 120
	SettingBrowserCacheTTLEditParamsValue300      SettingBrowserCacheTTLEditParamsValue = 300
	SettingBrowserCacheTTLEditParamsValue1200     SettingBrowserCacheTTLEditParamsValue = 1200
	SettingBrowserCacheTTLEditParamsValue1800     SettingBrowserCacheTTLEditParamsValue = 1800
	SettingBrowserCacheTTLEditParamsValue3600     SettingBrowserCacheTTLEditParamsValue = 3600
	SettingBrowserCacheTTLEditParamsValue7200     SettingBrowserCacheTTLEditParamsValue = 7200
	SettingBrowserCacheTTLEditParamsValue10800    SettingBrowserCacheTTLEditParamsValue = 10800
	SettingBrowserCacheTTLEditParamsValue14400    SettingBrowserCacheTTLEditParamsValue = 14400
	SettingBrowserCacheTTLEditParamsValue18000    SettingBrowserCacheTTLEditParamsValue = 18000
	SettingBrowserCacheTTLEditParamsValue28800    SettingBrowserCacheTTLEditParamsValue = 28800
	SettingBrowserCacheTTLEditParamsValue43200    SettingBrowserCacheTTLEditParamsValue = 43200
	SettingBrowserCacheTTLEditParamsValue57600    SettingBrowserCacheTTLEditParamsValue = 57600
	SettingBrowserCacheTTLEditParamsValue72000    SettingBrowserCacheTTLEditParamsValue = 72000
	SettingBrowserCacheTTLEditParamsValue86400    SettingBrowserCacheTTLEditParamsValue = 86400
	SettingBrowserCacheTTLEditParamsValue172800   SettingBrowserCacheTTLEditParamsValue = 172800
	SettingBrowserCacheTTLEditParamsValue259200   SettingBrowserCacheTTLEditParamsValue = 259200
	SettingBrowserCacheTTLEditParamsValue345600   SettingBrowserCacheTTLEditParamsValue = 345600
	SettingBrowserCacheTTLEditParamsValue432000   SettingBrowserCacheTTLEditParamsValue = 432000
	SettingBrowserCacheTTLEditParamsValue691200   SettingBrowserCacheTTLEditParamsValue = 691200
	SettingBrowserCacheTTLEditParamsValue1382400  SettingBrowserCacheTTLEditParamsValue = 1382400
	SettingBrowserCacheTTLEditParamsValue2073600  SettingBrowserCacheTTLEditParamsValue = 2073600
	SettingBrowserCacheTTLEditParamsValue2678400  SettingBrowserCacheTTLEditParamsValue = 2678400
	SettingBrowserCacheTTLEditParamsValue5356800  SettingBrowserCacheTTLEditParamsValue = 5356800
	SettingBrowserCacheTTLEditParamsValue16070400 SettingBrowserCacheTTLEditParamsValue = 16070400
	SettingBrowserCacheTTLEditParamsValue31536000 SettingBrowserCacheTTLEditParamsValue = 31536000
)

type SettingBrowserCacheTTLEditResponseEnvelope struct {
	Errors   []SettingBrowserCacheTTLEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingBrowserCacheTTLEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result SettingBrowserCacheTTLEditResponse             `json:"result"`
	JSON   settingBrowserCacheTTLEditResponseEnvelopeJSON `json:"-"`
}

// settingBrowserCacheTTLEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingBrowserCacheTTLEditResponseEnvelope]
type settingBrowserCacheTTLEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCacheTTLEditResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingBrowserCacheTTLEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingBrowserCacheTTLEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingBrowserCacheTTLEditResponseEnvelopeErrors]
type settingBrowserCacheTTLEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCacheTTLEditResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingBrowserCacheTTLEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingBrowserCacheTTLEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingBrowserCacheTTLEditResponseEnvelopeMessages]
type settingBrowserCacheTTLEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCacheTTLGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingBrowserCacheTTLGetResponseEnvelope struct {
	Errors   []SettingBrowserCacheTTLGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingBrowserCacheTTLGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result SettingBrowserCacheTTLGetResponse             `json:"result"`
	JSON   settingBrowserCacheTTLGetResponseEnvelopeJSON `json:"-"`
}

// settingBrowserCacheTTLGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingBrowserCacheTTLGetResponseEnvelope]
type settingBrowserCacheTTLGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCacheTTLGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingBrowserCacheTTLGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingBrowserCacheTTLGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingBrowserCacheTTLGetResponseEnvelopeErrors]
type settingBrowserCacheTTLGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCacheTTLGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingBrowserCacheTTLGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingBrowserCacheTTLGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingBrowserCacheTTLGetResponseEnvelopeMessages]
type settingBrowserCacheTTLGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
