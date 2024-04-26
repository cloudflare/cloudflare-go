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
	"github.com/cloudflare/cloudflare-go/v2/shared"
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
func (r *SettingAlwaysUseHTTPSService) Edit(ctx context.Context, params SettingAlwaysUseHTTPSEditParams, opts ...option.RequestOption) (res *AlwaysUseHTTPS, err error) {
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
func (r *SettingAlwaysUseHTTPSService) Get(ctx context.Context, query SettingAlwaysUseHTTPSGetParams, opts ...option.RequestOption) (res *AlwaysUseHTTPS, err error) {
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
type AlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID AlwaysUseHTTPSID `json:"id,required"`
	// Current value of the zone setting.
	Value AlwaysUseHTTPSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable AlwaysUseHTTPSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time          `json:"modified_on,nullable" format:"date-time"`
	JSON       alwaysUseHTTPSJSON `json:"-"`
}

// alwaysUseHTTPSJSON contains the JSON metadata for the struct [AlwaysUseHTTPS]
type alwaysUseHTTPSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alwaysUseHTTPSJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type AlwaysUseHTTPSID string

const (
	AlwaysUseHTTPSIDAlwaysUseHTTPS AlwaysUseHTTPSID = "always_use_https"
)

func (r AlwaysUseHTTPSID) IsKnown() bool {
	switch r {
	case AlwaysUseHTTPSIDAlwaysUseHTTPS:
		return true
	}
	return false
}

// Current value of the zone setting.
type AlwaysUseHTTPSValue string

const (
	AlwaysUseHTTPSValueOn  AlwaysUseHTTPSValue = "on"
	AlwaysUseHTTPSValueOff AlwaysUseHTTPSValue = "off"
)

func (r AlwaysUseHTTPSValue) IsKnown() bool {
	switch r {
	case AlwaysUseHTTPSValueOn, AlwaysUseHTTPSValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type AlwaysUseHTTPSEditable bool

const (
	AlwaysUseHTTPSEditableTrue  AlwaysUseHTTPSEditable = true
	AlwaysUseHTTPSEditableFalse AlwaysUseHTTPSEditable = false
)

func (r AlwaysUseHTTPSEditable) IsKnown() bool {
	switch r {
	case AlwaysUseHTTPSEditableTrue, AlwaysUseHTTPSEditableFalse:
		return true
	}
	return false
}

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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result AlwaysUseHTTPS                                `json:"result"`
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

type SettingAlwaysUseHTTPSGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingAlwaysUseHTTPSGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result AlwaysUseHTTPS                               `json:"result"`
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
