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
func (r *ZoneSettingMinTLSVersionService) Edit(ctx context.Context, params ZoneSettingMinTLSVersionEditParams, opts ...option.RequestOption) (res *ZonesMinTLSVersion, err error) {
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
func (r *ZoneSettingMinTLSVersionService) Get(ctx context.Context, query ZoneSettingMinTLSVersionGetParams, opts ...option.RequestOption) (res *ZonesMinTLSVersion, err error) {
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

func (r ZonesMinTLSVersion) implementsZoneSettingEditResponse() {}

func (r ZonesMinTLSVersion) implementsZoneSettingGetResponse() {}

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

func (r ZonesMinTLSVersionParam) implementsZoneSettingEditParamsItem() {}

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
	Result ZonesMinTLSVersion                               `json:"result"`
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
	Result ZonesMinTLSVersion                              `json:"result"`
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
