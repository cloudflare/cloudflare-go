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
func (r *SettingBrowserCacheTTLService) Update(ctx context.Context, zoneID string, body SettingBrowserCacheTTLUpdateParams, opts ...option.RequestOption) (res *SettingBrowserCacheTTLUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrowserCacheTTLUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
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
func (r *SettingBrowserCacheTTLService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingBrowserCacheTTLGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrowserCacheTTLGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
type SettingBrowserCacheTTLUpdateResponse struct {
	// ID of the zone setting.
	ID SettingBrowserCacheTTLUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingBrowserCacheTTLUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingBrowserCacheTTLUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingBrowserCacheTTLUpdateResponseJSON `json:"-"`
}

// settingBrowserCacheTTLUpdateResponseJSON contains the JSON metadata for the
// struct [SettingBrowserCacheTTLUpdateResponse]
type settingBrowserCacheTTLUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingBrowserCacheTTLUpdateResponseID string

const (
	SettingBrowserCacheTTLUpdateResponseIDBrowserCacheTTL SettingBrowserCacheTTLUpdateResponseID = "browser_cache_ttl"
)

// Current value of the zone setting.
type SettingBrowserCacheTTLUpdateResponseValue float64

const (
	SettingBrowserCacheTTLUpdateResponseValue0        SettingBrowserCacheTTLUpdateResponseValue = 0
	SettingBrowserCacheTTLUpdateResponseValue30       SettingBrowserCacheTTLUpdateResponseValue = 30
	SettingBrowserCacheTTLUpdateResponseValue60       SettingBrowserCacheTTLUpdateResponseValue = 60
	SettingBrowserCacheTTLUpdateResponseValue120      SettingBrowserCacheTTLUpdateResponseValue = 120
	SettingBrowserCacheTTLUpdateResponseValue300      SettingBrowserCacheTTLUpdateResponseValue = 300
	SettingBrowserCacheTTLUpdateResponseValue1200     SettingBrowserCacheTTLUpdateResponseValue = 1200
	SettingBrowserCacheTTLUpdateResponseValue1800     SettingBrowserCacheTTLUpdateResponseValue = 1800
	SettingBrowserCacheTTLUpdateResponseValue3600     SettingBrowserCacheTTLUpdateResponseValue = 3600
	SettingBrowserCacheTTLUpdateResponseValue7200     SettingBrowserCacheTTLUpdateResponseValue = 7200
	SettingBrowserCacheTTLUpdateResponseValue10800    SettingBrowserCacheTTLUpdateResponseValue = 10800
	SettingBrowserCacheTTLUpdateResponseValue14400    SettingBrowserCacheTTLUpdateResponseValue = 14400
	SettingBrowserCacheTTLUpdateResponseValue18000    SettingBrowserCacheTTLUpdateResponseValue = 18000
	SettingBrowserCacheTTLUpdateResponseValue28800    SettingBrowserCacheTTLUpdateResponseValue = 28800
	SettingBrowserCacheTTLUpdateResponseValue43200    SettingBrowserCacheTTLUpdateResponseValue = 43200
	SettingBrowserCacheTTLUpdateResponseValue57600    SettingBrowserCacheTTLUpdateResponseValue = 57600
	SettingBrowserCacheTTLUpdateResponseValue72000    SettingBrowserCacheTTLUpdateResponseValue = 72000
	SettingBrowserCacheTTLUpdateResponseValue86400    SettingBrowserCacheTTLUpdateResponseValue = 86400
	SettingBrowserCacheTTLUpdateResponseValue172800   SettingBrowserCacheTTLUpdateResponseValue = 172800
	SettingBrowserCacheTTLUpdateResponseValue259200   SettingBrowserCacheTTLUpdateResponseValue = 259200
	SettingBrowserCacheTTLUpdateResponseValue345600   SettingBrowserCacheTTLUpdateResponseValue = 345600
	SettingBrowserCacheTTLUpdateResponseValue432000   SettingBrowserCacheTTLUpdateResponseValue = 432000
	SettingBrowserCacheTTLUpdateResponseValue691200   SettingBrowserCacheTTLUpdateResponseValue = 691200
	SettingBrowserCacheTTLUpdateResponseValue1382400  SettingBrowserCacheTTLUpdateResponseValue = 1382400
	SettingBrowserCacheTTLUpdateResponseValue2073600  SettingBrowserCacheTTLUpdateResponseValue = 2073600
	SettingBrowserCacheTTLUpdateResponseValue2678400  SettingBrowserCacheTTLUpdateResponseValue = 2678400
	SettingBrowserCacheTTLUpdateResponseValue5356800  SettingBrowserCacheTTLUpdateResponseValue = 5356800
	SettingBrowserCacheTTLUpdateResponseValue16070400 SettingBrowserCacheTTLUpdateResponseValue = 16070400
	SettingBrowserCacheTTLUpdateResponseValue31536000 SettingBrowserCacheTTLUpdateResponseValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingBrowserCacheTTLUpdateResponseEditable bool

const (
	SettingBrowserCacheTTLUpdateResponseEditableTrue  SettingBrowserCacheTTLUpdateResponseEditable = true
	SettingBrowserCacheTTLUpdateResponseEditableFalse SettingBrowserCacheTTLUpdateResponseEditable = false
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

type SettingBrowserCacheTTLUpdateParams struct {
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value param.Field[SettingBrowserCacheTTLUpdateParamsValue] `json:"value,required"`
}

func (r SettingBrowserCacheTTLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type SettingBrowserCacheTTLUpdateParamsValue float64

const (
	SettingBrowserCacheTTLUpdateParamsValue0        SettingBrowserCacheTTLUpdateParamsValue = 0
	SettingBrowserCacheTTLUpdateParamsValue30       SettingBrowserCacheTTLUpdateParamsValue = 30
	SettingBrowserCacheTTLUpdateParamsValue60       SettingBrowserCacheTTLUpdateParamsValue = 60
	SettingBrowserCacheTTLUpdateParamsValue120      SettingBrowserCacheTTLUpdateParamsValue = 120
	SettingBrowserCacheTTLUpdateParamsValue300      SettingBrowserCacheTTLUpdateParamsValue = 300
	SettingBrowserCacheTTLUpdateParamsValue1200     SettingBrowserCacheTTLUpdateParamsValue = 1200
	SettingBrowserCacheTTLUpdateParamsValue1800     SettingBrowserCacheTTLUpdateParamsValue = 1800
	SettingBrowserCacheTTLUpdateParamsValue3600     SettingBrowserCacheTTLUpdateParamsValue = 3600
	SettingBrowserCacheTTLUpdateParamsValue7200     SettingBrowserCacheTTLUpdateParamsValue = 7200
	SettingBrowserCacheTTLUpdateParamsValue10800    SettingBrowserCacheTTLUpdateParamsValue = 10800
	SettingBrowserCacheTTLUpdateParamsValue14400    SettingBrowserCacheTTLUpdateParamsValue = 14400
	SettingBrowserCacheTTLUpdateParamsValue18000    SettingBrowserCacheTTLUpdateParamsValue = 18000
	SettingBrowserCacheTTLUpdateParamsValue28800    SettingBrowserCacheTTLUpdateParamsValue = 28800
	SettingBrowserCacheTTLUpdateParamsValue43200    SettingBrowserCacheTTLUpdateParamsValue = 43200
	SettingBrowserCacheTTLUpdateParamsValue57600    SettingBrowserCacheTTLUpdateParamsValue = 57600
	SettingBrowserCacheTTLUpdateParamsValue72000    SettingBrowserCacheTTLUpdateParamsValue = 72000
	SettingBrowserCacheTTLUpdateParamsValue86400    SettingBrowserCacheTTLUpdateParamsValue = 86400
	SettingBrowserCacheTTLUpdateParamsValue172800   SettingBrowserCacheTTLUpdateParamsValue = 172800
	SettingBrowserCacheTTLUpdateParamsValue259200   SettingBrowserCacheTTLUpdateParamsValue = 259200
	SettingBrowserCacheTTLUpdateParamsValue345600   SettingBrowserCacheTTLUpdateParamsValue = 345600
	SettingBrowserCacheTTLUpdateParamsValue432000   SettingBrowserCacheTTLUpdateParamsValue = 432000
	SettingBrowserCacheTTLUpdateParamsValue691200   SettingBrowserCacheTTLUpdateParamsValue = 691200
	SettingBrowserCacheTTLUpdateParamsValue1382400  SettingBrowserCacheTTLUpdateParamsValue = 1382400
	SettingBrowserCacheTTLUpdateParamsValue2073600  SettingBrowserCacheTTLUpdateParamsValue = 2073600
	SettingBrowserCacheTTLUpdateParamsValue2678400  SettingBrowserCacheTTLUpdateParamsValue = 2678400
	SettingBrowserCacheTTLUpdateParamsValue5356800  SettingBrowserCacheTTLUpdateParamsValue = 5356800
	SettingBrowserCacheTTLUpdateParamsValue16070400 SettingBrowserCacheTTLUpdateParamsValue = 16070400
	SettingBrowserCacheTTLUpdateParamsValue31536000 SettingBrowserCacheTTLUpdateParamsValue = 31536000
)

type SettingBrowserCacheTTLUpdateResponseEnvelope struct {
	Errors   []SettingBrowserCacheTTLUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingBrowserCacheTTLUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result SettingBrowserCacheTTLUpdateResponse             `json:"result"`
	JSON   settingBrowserCacheTTLUpdateResponseEnvelopeJSON `json:"-"`
}

// settingBrowserCacheTTLUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingBrowserCacheTTLUpdateResponseEnvelope]
type settingBrowserCacheTTLUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingBrowserCacheTTLUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingBrowserCacheTTLUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingBrowserCacheTTLUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingBrowserCacheTTLUpdateResponseEnvelopeErrors]
type settingBrowserCacheTTLUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingBrowserCacheTTLUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    settingBrowserCacheTTLUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingBrowserCacheTTLUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingBrowserCacheTTLUpdateResponseEnvelopeMessages]
type settingBrowserCacheTTLUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
