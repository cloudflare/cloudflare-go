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
func (r *SettingBrowserCacheTTLService) Edit(ctx context.Context, params SettingBrowserCacheTTLEditParams, opts ...option.RequestOption) (res *BrowserCacheTTL, err error) {
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
func (r *SettingBrowserCacheTTLService) Get(ctx context.Context, query SettingBrowserCacheTTLGetParams, opts ...option.RequestOption) (res *BrowserCacheTTL, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrowserCacheTTLGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", query.ZoneID)
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
type BrowserCacheTTL struct {
	// ID of the zone setting.
	ID BrowserCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value BrowserCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable BrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       browserCacheTTLJSON `json:"-"`
}

// browserCacheTTLJSON contains the JSON metadata for the struct [BrowserCacheTTL]
type browserCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserCacheTTLJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type BrowserCacheTTLID string

const (
	BrowserCacheTTLIDBrowserCacheTTL BrowserCacheTTLID = "browser_cache_ttl"
)

func (r BrowserCacheTTLID) IsKnown() bool {
	switch r {
	case BrowserCacheTTLIDBrowserCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type BrowserCacheTTLValue float64

const (
	BrowserCacheTTLValue0        BrowserCacheTTLValue = 0
	BrowserCacheTTLValue30       BrowserCacheTTLValue = 30
	BrowserCacheTTLValue60       BrowserCacheTTLValue = 60
	BrowserCacheTTLValue120      BrowserCacheTTLValue = 120
	BrowserCacheTTLValue300      BrowserCacheTTLValue = 300
	BrowserCacheTTLValue1200     BrowserCacheTTLValue = 1200
	BrowserCacheTTLValue1800     BrowserCacheTTLValue = 1800
	BrowserCacheTTLValue3600     BrowserCacheTTLValue = 3600
	BrowserCacheTTLValue7200     BrowserCacheTTLValue = 7200
	BrowserCacheTTLValue10800    BrowserCacheTTLValue = 10800
	BrowserCacheTTLValue14400    BrowserCacheTTLValue = 14400
	BrowserCacheTTLValue18000    BrowserCacheTTLValue = 18000
	BrowserCacheTTLValue28800    BrowserCacheTTLValue = 28800
	BrowserCacheTTLValue43200    BrowserCacheTTLValue = 43200
	BrowserCacheTTLValue57600    BrowserCacheTTLValue = 57600
	BrowserCacheTTLValue72000    BrowserCacheTTLValue = 72000
	BrowserCacheTTLValue86400    BrowserCacheTTLValue = 86400
	BrowserCacheTTLValue172800   BrowserCacheTTLValue = 172800
	BrowserCacheTTLValue259200   BrowserCacheTTLValue = 259200
	BrowserCacheTTLValue345600   BrowserCacheTTLValue = 345600
	BrowserCacheTTLValue432000   BrowserCacheTTLValue = 432000
	BrowserCacheTTLValue691200   BrowserCacheTTLValue = 691200
	BrowserCacheTTLValue1382400  BrowserCacheTTLValue = 1382400
	BrowserCacheTTLValue2073600  BrowserCacheTTLValue = 2073600
	BrowserCacheTTLValue2678400  BrowserCacheTTLValue = 2678400
	BrowserCacheTTLValue5356800  BrowserCacheTTLValue = 5356800
	BrowserCacheTTLValue16070400 BrowserCacheTTLValue = 16070400
	BrowserCacheTTLValue31536000 BrowserCacheTTLValue = 31536000
)

func (r BrowserCacheTTLValue) IsKnown() bool {
	switch r {
	case BrowserCacheTTLValue0, BrowserCacheTTLValue30, BrowserCacheTTLValue60, BrowserCacheTTLValue120, BrowserCacheTTLValue300, BrowserCacheTTLValue1200, BrowserCacheTTLValue1800, BrowserCacheTTLValue3600, BrowserCacheTTLValue7200, BrowserCacheTTLValue10800, BrowserCacheTTLValue14400, BrowserCacheTTLValue18000, BrowserCacheTTLValue28800, BrowserCacheTTLValue43200, BrowserCacheTTLValue57600, BrowserCacheTTLValue72000, BrowserCacheTTLValue86400, BrowserCacheTTLValue172800, BrowserCacheTTLValue259200, BrowserCacheTTLValue345600, BrowserCacheTTLValue432000, BrowserCacheTTLValue691200, BrowserCacheTTLValue1382400, BrowserCacheTTLValue2073600, BrowserCacheTTLValue2678400, BrowserCacheTTLValue5356800, BrowserCacheTTLValue16070400, BrowserCacheTTLValue31536000:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type BrowserCacheTTLEditable bool

const (
	BrowserCacheTTLEditableTrue  BrowserCacheTTLEditable = true
	BrowserCacheTTLEditableFalse BrowserCacheTTLEditable = false
)

func (r BrowserCacheTTLEditable) IsKnown() bool {
	switch r {
	case BrowserCacheTTLEditableTrue, BrowserCacheTTLEditableFalse:
		return true
	}
	return false
}

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

func (r SettingBrowserCacheTTLEditParamsValue) IsKnown() bool {
	switch r {
	case SettingBrowserCacheTTLEditParamsValue0, SettingBrowserCacheTTLEditParamsValue30, SettingBrowserCacheTTLEditParamsValue60, SettingBrowserCacheTTLEditParamsValue120, SettingBrowserCacheTTLEditParamsValue300, SettingBrowserCacheTTLEditParamsValue1200, SettingBrowserCacheTTLEditParamsValue1800, SettingBrowserCacheTTLEditParamsValue3600, SettingBrowserCacheTTLEditParamsValue7200, SettingBrowserCacheTTLEditParamsValue10800, SettingBrowserCacheTTLEditParamsValue14400, SettingBrowserCacheTTLEditParamsValue18000, SettingBrowserCacheTTLEditParamsValue28800, SettingBrowserCacheTTLEditParamsValue43200, SettingBrowserCacheTTLEditParamsValue57600, SettingBrowserCacheTTLEditParamsValue72000, SettingBrowserCacheTTLEditParamsValue86400, SettingBrowserCacheTTLEditParamsValue172800, SettingBrowserCacheTTLEditParamsValue259200, SettingBrowserCacheTTLEditParamsValue345600, SettingBrowserCacheTTLEditParamsValue432000, SettingBrowserCacheTTLEditParamsValue691200, SettingBrowserCacheTTLEditParamsValue1382400, SettingBrowserCacheTTLEditParamsValue2073600, SettingBrowserCacheTTLEditParamsValue2678400, SettingBrowserCacheTTLEditParamsValue5356800, SettingBrowserCacheTTLEditParamsValue16070400, SettingBrowserCacheTTLEditParamsValue31536000:
		return true
	}
	return false
}

type SettingBrowserCacheTTLEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result BrowserCacheTTL                                `json:"result"`
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

type SettingBrowserCacheTTLGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingBrowserCacheTTLGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result BrowserCacheTTL                               `json:"result"`
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
