// File generated from our OpenAPI spec by Stainless.

package cloudflare

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

// ZoneSettingTLSClientAuthService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingTLSClientAuthService] method instead.
type ZoneSettingTLSClientAuthService struct {
	Options []option.RequestOption
}

// NewZoneSettingTLSClientAuthService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingTLSClientAuthService(opts ...option.RequestOption) (r *ZoneSettingTLSClientAuthService) {
	r = &ZoneSettingTLSClientAuthService{}
	r.Options = opts
	return
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
func (r *ZoneSettingTLSClientAuthService) Edit(ctx context.Context, params ZoneSettingTLSClientAuthEditParams, opts ...option.RequestOption) (res *ZoneSettingTLSClientAuthEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingTLSClientAuthEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/tls_client_auth", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
func (r *ZoneSettingTLSClientAuthService) Get(ctx context.Context, query ZoneSettingTLSClientAuthGetParams, opts ...option.RequestOption) (res *ZoneSettingTLSClientAuthGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingTLSClientAuthGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/tls_client_auth", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingTLSClientAuthEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingTLSClientAuthEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingTLSClientAuthEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingTLSClientAuthEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingTLSClientAuthEditResponseJSON `json:"-"`
}

// zoneSettingTLSClientAuthEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingTLSClientAuthEditResponse]
type zoneSettingTLSClientAuthEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingTLSClientAuthEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingTLSClientAuthEditResponseID string

const (
	ZoneSettingTLSClientAuthEditResponseIDTLSClientAuth ZoneSettingTLSClientAuthEditResponseID = "tls_client_auth"
)

// Current value of the zone setting.
type ZoneSettingTLSClientAuthEditResponseValue string

const (
	ZoneSettingTLSClientAuthEditResponseValueOn  ZoneSettingTLSClientAuthEditResponseValue = "on"
	ZoneSettingTLSClientAuthEditResponseValueOff ZoneSettingTLSClientAuthEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingTLSClientAuthEditResponseEditable bool

const (
	ZoneSettingTLSClientAuthEditResponseEditableTrue  ZoneSettingTLSClientAuthEditResponseEditable = true
	ZoneSettingTLSClientAuthEditResponseEditableFalse ZoneSettingTLSClientAuthEditResponseEditable = false
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingTLSClientAuthGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingTLSClientAuthGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingTLSClientAuthGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingTLSClientAuthGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingTLSClientAuthGetResponseJSON `json:"-"`
}

// zoneSettingTLSClientAuthGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingTLSClientAuthGetResponse]
type zoneSettingTLSClientAuthGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingTLSClientAuthGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingTLSClientAuthGetResponseID string

const (
	ZoneSettingTLSClientAuthGetResponseIDTLSClientAuth ZoneSettingTLSClientAuthGetResponseID = "tls_client_auth"
)

// Current value of the zone setting.
type ZoneSettingTLSClientAuthGetResponseValue string

const (
	ZoneSettingTLSClientAuthGetResponseValueOn  ZoneSettingTLSClientAuthGetResponseValue = "on"
	ZoneSettingTLSClientAuthGetResponseValueOff ZoneSettingTLSClientAuthGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingTLSClientAuthGetResponseEditable bool

const (
	ZoneSettingTLSClientAuthGetResponseEditableTrue  ZoneSettingTLSClientAuthGetResponseEditable = true
	ZoneSettingTLSClientAuthGetResponseEditableFalse ZoneSettingTLSClientAuthGetResponseEditable = false
)

type ZoneSettingTLSClientAuthEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// value of the zone setting.
	Value param.Field[ZoneSettingTLSClientAuthEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingTLSClientAuthEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// value of the zone setting.
type ZoneSettingTLSClientAuthEditParamsValue string

const (
	ZoneSettingTLSClientAuthEditParamsValueOn  ZoneSettingTLSClientAuthEditParamsValue = "on"
	ZoneSettingTLSClientAuthEditParamsValueOff ZoneSettingTLSClientAuthEditParamsValue = "off"
)

type ZoneSettingTLSClientAuthEditResponseEnvelope struct {
	Errors   []ZoneSettingTLSClientAuthEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingTLSClientAuthEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result ZoneSettingTLSClientAuthEditResponse             `json:"result"`
	JSON   zoneSettingTLSClientAuthEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingTLSClientAuthEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingTLSClientAuthEditResponseEnvelope]
type zoneSettingTLSClientAuthEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingTLSClientAuthEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingTLSClientAuthEditResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingTLSClientAuthEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingTLSClientAuthEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingTLSClientAuthEditResponseEnvelopeErrors]
type zoneSettingTLSClientAuthEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingTLSClientAuthEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingTLSClientAuthEditResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingTLSClientAuthEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingTLSClientAuthEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingTLSClientAuthEditResponseEnvelopeMessages]
type zoneSettingTLSClientAuthEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingTLSClientAuthEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingTLSClientAuthGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingTLSClientAuthGetResponseEnvelope struct {
	Errors   []ZoneSettingTLSClientAuthGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingTLSClientAuthGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result ZoneSettingTLSClientAuthGetResponse             `json:"result"`
	JSON   zoneSettingTLSClientAuthGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingTLSClientAuthGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingTLSClientAuthGetResponseEnvelope]
type zoneSettingTLSClientAuthGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingTLSClientAuthGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingTLSClientAuthGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneSettingTLSClientAuthGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingTLSClientAuthGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingTLSClientAuthGetResponseEnvelopeErrors]
type zoneSettingTLSClientAuthGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingTLSClientAuthGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingTLSClientAuthGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingTLSClientAuthGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingTLSClientAuthGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingTLSClientAuthGetResponseEnvelopeMessages]
type zoneSettingTLSClientAuthGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingTLSClientAuthGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
