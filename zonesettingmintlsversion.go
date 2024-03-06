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

// ZoneSettingMinTLSVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingMinTLSVersionService] method instead.
type ZoneSettingMinTLSVersionService struct {
	Options []option.RequestOption
}

// NewZoneSettingMinTLSVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingMinTLSVersionService(opts ...option.RequestOption) (r *ZoneSettingMinTLSVersionService) {
	r = &ZoneSettingMinTLSVersionService{}
	r.Options = opts
	return
}

// Changes Minimum TLS Version setting.
func (r *ZoneSettingMinTLSVersionService) Edit(ctx context.Context, params ZoneSettingMinTLSVersionEditParams, opts ...option.RequestOption) (res *ZoneSettingMinTLSVersionEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingMinTLSVersionEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets Minimum TLS Version setting.
func (r *ZoneSettingMinTLSVersionService) Get(ctx context.Context, query ZoneSettingMinTLSVersionGetParams, opts ...option.RequestOption) (res *ZoneSettingMinTLSVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingMinTLSVersionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZoneSettingMinTLSVersionEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingMinTLSVersionEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingMinTLSVersionEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMinTLSVersionEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMinTLSVersionEditResponseJSON `json:"-"`
}

// zoneSettingMinTLSVersionEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingMinTLSVersionEditResponse]
type zoneSettingMinTLSVersionEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMinTLSVersionEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingMinTLSVersionEditResponseID string

const (
	ZoneSettingMinTLSVersionEditResponseIDMinTLSVersion ZoneSettingMinTLSVersionEditResponseID = "min_tls_version"
)

// Current value of the zone setting.
type ZoneSettingMinTLSVersionEditResponseValue string

const (
	ZoneSettingMinTLSVersionEditResponseValue1_0 ZoneSettingMinTLSVersionEditResponseValue = "1.0"
	ZoneSettingMinTLSVersionEditResponseValue1_1 ZoneSettingMinTLSVersionEditResponseValue = "1.1"
	ZoneSettingMinTLSVersionEditResponseValue1_2 ZoneSettingMinTLSVersionEditResponseValue = "1.2"
	ZoneSettingMinTLSVersionEditResponseValue1_3 ZoneSettingMinTLSVersionEditResponseValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMinTLSVersionEditResponseEditable bool

const (
	ZoneSettingMinTLSVersionEditResponseEditableTrue  ZoneSettingMinTLSVersionEditResponseEditable = true
	ZoneSettingMinTLSVersionEditResponseEditableFalse ZoneSettingMinTLSVersionEditResponseEditable = false
)

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZoneSettingMinTLSVersionGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingMinTLSVersionGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingMinTLSVersionGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMinTLSVersionGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMinTLSVersionGetResponseJSON `json:"-"`
}

// zoneSettingMinTLSVersionGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingMinTLSVersionGetResponse]
type zoneSettingMinTLSVersionGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMinTLSVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingMinTLSVersionGetResponseID string

const (
	ZoneSettingMinTLSVersionGetResponseIDMinTLSVersion ZoneSettingMinTLSVersionGetResponseID = "min_tls_version"
)

// Current value of the zone setting.
type ZoneSettingMinTLSVersionGetResponseValue string

const (
	ZoneSettingMinTLSVersionGetResponseValue1_0 ZoneSettingMinTLSVersionGetResponseValue = "1.0"
	ZoneSettingMinTLSVersionGetResponseValue1_1 ZoneSettingMinTLSVersionGetResponseValue = "1.1"
	ZoneSettingMinTLSVersionGetResponseValue1_2 ZoneSettingMinTLSVersionGetResponseValue = "1.2"
	ZoneSettingMinTLSVersionGetResponseValue1_3 ZoneSettingMinTLSVersionGetResponseValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMinTLSVersionGetResponseEditable bool

const (
	ZoneSettingMinTLSVersionGetResponseEditableTrue  ZoneSettingMinTLSVersionGetResponseEditable = true
	ZoneSettingMinTLSVersionGetResponseEditableFalse ZoneSettingMinTLSVersionGetResponseEditable = false
)

type ZoneSettingMinTLSVersionEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingMinTLSVersionEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingMinTLSVersionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingMinTLSVersionEditParamsValue string

const (
	ZoneSettingMinTLSVersionEditParamsValue1_0 ZoneSettingMinTLSVersionEditParamsValue = "1.0"
	ZoneSettingMinTLSVersionEditParamsValue1_1 ZoneSettingMinTLSVersionEditParamsValue = "1.1"
	ZoneSettingMinTLSVersionEditParamsValue1_2 ZoneSettingMinTLSVersionEditParamsValue = "1.2"
	ZoneSettingMinTLSVersionEditParamsValue1_3 ZoneSettingMinTLSVersionEditParamsValue = "1.3"
)

type ZoneSettingMinTLSVersionEditResponseEnvelope struct {
	Errors   []ZoneSettingMinTLSVersionEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingMinTLSVersionEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result ZoneSettingMinTLSVersionEditResponse             `json:"result"`
	JSON   zoneSettingMinTLSVersionEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingMinTLSVersionEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingMinTLSVersionEditResponseEnvelope]
type zoneSettingMinTLSVersionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMinTLSVersionEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingMinTLSVersionEditResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingMinTLSVersionEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingMinTLSVersionEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingMinTLSVersionEditResponseEnvelopeErrors]
type zoneSettingMinTLSVersionEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMinTLSVersionEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingMinTLSVersionEditResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingMinTLSVersionEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingMinTLSVersionEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingMinTLSVersionEditResponseEnvelopeMessages]
type zoneSettingMinTLSVersionEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMinTLSVersionEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingMinTLSVersionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingMinTLSVersionGetResponseEnvelope struct {
	Errors   []ZoneSettingMinTLSVersionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingMinTLSVersionGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result ZoneSettingMinTLSVersionGetResponse             `json:"result"`
	JSON   zoneSettingMinTLSVersionGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingMinTLSVersionGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingMinTLSVersionGetResponseEnvelope]
type zoneSettingMinTLSVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMinTLSVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingMinTLSVersionGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneSettingMinTLSVersionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingMinTLSVersionGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingMinTLSVersionGetResponseEnvelopeErrors]
type zoneSettingMinTLSVersionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMinTLSVersionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingMinTLSVersionGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingMinTLSVersionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingMinTLSVersionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingMinTLSVersionGetResponseEnvelopeMessages]
type zoneSettingMinTLSVersionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMinTLSVersionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
