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
func (r *SettingMinTLSVersionService) Edit(ctx context.Context, params SettingMinTLSVersionEditParams, opts ...option.RequestOption) (res *ZoneSettingMinTLSVersion, err error) {
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
func (r *SettingMinTLSVersionService) Get(ctx context.Context, query SettingMinTLSVersionGetParams, opts ...option.RequestOption) (res *ZoneSettingMinTLSVersion, err error) {
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
type ZoneSettingMinTLSVersion struct {
	// ID of the zone setting.
	ID ZoneSettingMinTLSVersionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingMinTLSVersionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMinTLSVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMinTLSVersionJSON `json:"-"`
}

// zoneSettingMinTLSVersionJSON contains the JSON metadata for the struct
// [ZoneSettingMinTLSVersion]
type zoneSettingMinTLSVersionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMinTLSVersionJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingMinTLSVersion) implementsZonesSettingEditResponse() {}

func (r ZoneSettingMinTLSVersion) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingMinTLSVersionID string

const (
	ZoneSettingMinTLSVersionIDMinTLSVersion ZoneSettingMinTLSVersionID = "min_tls_version"
)

func (r ZoneSettingMinTLSVersionID) IsKnown() bool {
	switch r {
	case ZoneSettingMinTLSVersionIDMinTLSVersion:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingMinTLSVersionValue string

const (
	ZoneSettingMinTLSVersionValue1_0 ZoneSettingMinTLSVersionValue = "1.0"
	ZoneSettingMinTLSVersionValue1_1 ZoneSettingMinTLSVersionValue = "1.1"
	ZoneSettingMinTLSVersionValue1_2 ZoneSettingMinTLSVersionValue = "1.2"
	ZoneSettingMinTLSVersionValue1_3 ZoneSettingMinTLSVersionValue = "1.3"
)

func (r ZoneSettingMinTLSVersionValue) IsKnown() bool {
	switch r {
	case ZoneSettingMinTLSVersionValue1_0, ZoneSettingMinTLSVersionValue1_1, ZoneSettingMinTLSVersionValue1_2, ZoneSettingMinTLSVersionValue1_3:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMinTLSVersionEditable bool

const (
	ZoneSettingMinTLSVersionEditableTrue  ZoneSettingMinTLSVersionEditable = true
	ZoneSettingMinTLSVersionEditableFalse ZoneSettingMinTLSVersionEditable = false
)

func (r ZoneSettingMinTLSVersionEditable) IsKnown() bool {
	switch r {
	case ZoneSettingMinTLSVersionEditableTrue, ZoneSettingMinTLSVersionEditableFalse:
		return true
	}
	return false
}

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZoneSettingMinTLSVersionParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingMinTLSVersionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingMinTLSVersionValue] `json:"value,required"`
}

func (r ZoneSettingMinTLSVersionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingMinTLSVersionParam) implementsZonesSettingEditParamsItemUnion() {}

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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result ZoneSettingMinTLSVersion                     `json:"result"`
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result ZoneSettingMinTLSVersion                    `json:"result"`
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
