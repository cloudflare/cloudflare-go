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
func (r *SettingMinTLSVersionService) Edit(ctx context.Context, params SettingMinTLSVersionEditParams, opts ...option.RequestOption) (res *ZonesMinTLSVersion, err error) {
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
func (r *SettingMinTLSVersionService) Get(ctx context.Context, query SettingMinTLSVersionGetParams, opts ...option.RequestOption) (res *ZonesMinTLSVersion, err error) {
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
type ZonesMinTLSVersion struct {
	// ID of the zone setting.
	ID ZonesMinTLSVersionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesMinTLSVersionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesMinTLSVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time              `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesMinTLSVersionJSON `json:"-"`
}

// zonesMinTLSVersionJSON contains the JSON metadata for the struct
// [ZonesMinTLSVersion]
type zonesMinTLSVersionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesMinTLSVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesMinTLSVersionJSON) RawJSON() string {
	return r.raw
}

func (r ZonesMinTLSVersion) implementsZonesSettingEditResponse() {}

func (r ZonesMinTLSVersion) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesMinTLSVersionID string

const (
	ZonesMinTLSVersionIDMinTLSVersion ZonesMinTLSVersionID = "min_tls_version"
)

// Current value of the zone setting.
type ZonesMinTLSVersionValue string

const (
	ZonesMinTLSVersionValue1_0 ZonesMinTLSVersionValue = "1.0"
	ZonesMinTLSVersionValue1_1 ZonesMinTLSVersionValue = "1.1"
	ZonesMinTLSVersionValue1_2 ZonesMinTLSVersionValue = "1.2"
	ZonesMinTLSVersionValue1_3 ZonesMinTLSVersionValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesMinTLSVersionEditable bool

const (
	ZonesMinTLSVersionEditableTrue  ZonesMinTLSVersionEditable = true
	ZonesMinTLSVersionEditableFalse ZonesMinTLSVersionEditable = false
)

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZonesMinTLSVersionParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesMinTLSVersionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesMinTLSVersionValue] `json:"value,required"`
}

func (r ZonesMinTLSVersionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesMinTLSVersionParam) implementsZonesSettingEditParamsItem() {}

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

type SettingMinTLSVersionEditResponseEnvelope struct {
	Errors   []SettingMinTLSVersionEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMinTLSVersionEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result ZonesMinTLSVersion                           `json:"result"`
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

type SettingMinTLSVersionEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingMinTLSVersionEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMinTLSVersionEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingMinTLSVersionEditResponseEnvelopeErrors]
type settingMinTLSVersionEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMinTLSVersionEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingMinTLSVersionEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingMinTLSVersionEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMinTLSVersionEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingMinTLSVersionEditResponseEnvelopeMessages]
type settingMinTLSVersionEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMinTLSVersionEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingMinTLSVersionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingMinTLSVersionGetResponseEnvelope struct {
	Errors   []SettingMinTLSVersionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMinTLSVersionGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result ZonesMinTLSVersion                          `json:"result"`
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

type SettingMinTLSVersionGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingMinTLSVersionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMinTLSVersionGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingMinTLSVersionGetResponseEnvelopeErrors]
type settingMinTLSVersionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMinTLSVersionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingMinTLSVersionGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingMinTLSVersionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMinTLSVersionGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingMinTLSVersionGetResponseEnvelopeMessages]
type settingMinTLSVersionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMinTLSVersionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
