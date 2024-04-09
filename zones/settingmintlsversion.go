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

// SettingMinTLSVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingMinTLSVersionService]
// method instead.
type SettingMinTLSVersionService struct {
	Options []option.RequestOption
}

// NewSettingMinTLSVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingMinTLSVersionService(opts ...option.RequestOption) (r *SettingMinTLSVersionService) {
	r = &SettingMinTLSVersionService{}
	r.Options = opts
	return
}

// Changes Minimum TLS Version setting.
func (r *SettingMinTLSVersionService) Edit(ctx context.Context, params SettingMinTLSVersionEditParams, opts ...option.RequestOption) (res *MinTLSVersion, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMinTLSVersionEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets Minimum TLS Version setting.
func (r *SettingMinTLSVersionService) Get(ctx context.Context, query SettingMinTLSVersionGetParams, opts ...option.RequestOption) (res *MinTLSVersion, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMinTLSVersionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type MinTLSVersion struct {
	// ID of the zone setting.
	ID MinTLSVersionID `json:"id,required"`
	// Current value of the zone setting.
	Value MinTLSVersionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable MinTLSVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time         `json:"modified_on,nullable" format:"date-time"`
	JSON       minTLSVersionJSON `json:"-"`
}

// minTLSVersionJSON contains the JSON metadata for the struct [MinTLSVersion]
type minTLSVersionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MinTLSVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r minTLSVersionJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type MinTLSVersionID string

const (
	MinTLSVersionIDMinTLSVersion MinTLSVersionID = "min_tls_version"
)

func (r MinTLSVersionID) IsKnown() bool {
	switch r {
	case MinTLSVersionIDMinTLSVersion:
		return true
	}
	return false
}

// Current value of the zone setting.
type MinTLSVersionValue string

const (
	MinTLSVersionValue1_0 MinTLSVersionValue = "1.0"
	MinTLSVersionValue1_1 MinTLSVersionValue = "1.1"
	MinTLSVersionValue1_2 MinTLSVersionValue = "1.2"
	MinTLSVersionValue1_3 MinTLSVersionValue = "1.3"
)

func (r MinTLSVersionValue) IsKnown() bool {
	switch r {
	case MinTLSVersionValue1_0, MinTLSVersionValue1_1, MinTLSVersionValue1_2, MinTLSVersionValue1_3:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type MinTLSVersionEditable bool

const (
	MinTLSVersionEditableTrue  MinTLSVersionEditable = true
	MinTLSVersionEditableFalse MinTLSVersionEditable = false
)

func (r MinTLSVersionEditable) IsKnown() bool {
	switch r {
	case MinTLSVersionEditableTrue, MinTLSVersionEditableFalse:
		return true
	}
	return false
}

type SettingMinTLSVersionEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingMinTLSVersionEditParamsValue] `json:"value,required"`
}

func (r SettingMinTLSVersionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingMinTLSVersionEditParamsValue string

const (
	SettingMinTLSVersionEditParamsValue1_0 SettingMinTLSVersionEditParamsValue = "1.0"
	SettingMinTLSVersionEditParamsValue1_1 SettingMinTLSVersionEditParamsValue = "1.1"
	SettingMinTLSVersionEditParamsValue1_2 SettingMinTLSVersionEditParamsValue = "1.2"
	SettingMinTLSVersionEditParamsValue1_3 SettingMinTLSVersionEditParamsValue = "1.3"
)

func (r SettingMinTLSVersionEditParamsValue) IsKnown() bool {
	switch r {
	case SettingMinTLSVersionEditParamsValue1_0, SettingMinTLSVersionEditParamsValue1_1, SettingMinTLSVersionEditParamsValue1_2, SettingMinTLSVersionEditParamsValue1_3:
		return true
	}
	return false
}

type SettingMinTLSVersionEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result MinTLSVersion                                `json:"result"`
	JSON   settingMinTLSVersionEditResponseEnvelopeJSON `json:"-"`
}

// settingMinTLSVersionEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingMinTLSVersionEditResponseEnvelope]
type settingMinTLSVersionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMinTLSVersionEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingMinTLSVersionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingMinTLSVersionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result MinTLSVersion                               `json:"result"`
	JSON   settingMinTLSVersionGetResponseEnvelopeJSON `json:"-"`
}

// settingMinTLSVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingMinTLSVersionGetResponseEnvelope]
type settingMinTLSVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMinTLSVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
