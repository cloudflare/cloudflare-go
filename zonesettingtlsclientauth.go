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
func (r *ZoneSettingTLSClientAuthService) Edit(ctx context.Context, params ZoneSettingTLSClientAuthEditParams, opts ...option.RequestOption) (res *ZonesTLSClientAuth, err error) {
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
func (r *ZoneSettingTLSClientAuthService) Get(ctx context.Context, query ZoneSettingTLSClientAuthGetParams, opts ...option.RequestOption) (res *ZonesTLSClientAuth, err error) {
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
type ZonesTLSClientAuth struct {
	// ID of the zone setting.
	ID ZonesTLSClientAuthID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesTLSClientAuthValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesTLSClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time              `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesTLSClientAuthJSON `json:"-"`
}

// zonesTLSClientAuthJSON contains the JSON metadata for the struct
// [ZonesTLSClientAuth]
type zonesTLSClientAuthJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesTLSClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesTLSClientAuth) implementsZoneSettingEditResponse() {}

func (r ZonesTLSClientAuth) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesTLSClientAuthID string

const (
	ZonesTLSClientAuthIDTLSClientAuth ZonesTLSClientAuthID = "tls_client_auth"
)

// Current value of the zone setting.
type ZonesTLSClientAuthValue string

const (
	ZonesTLSClientAuthValueOn  ZonesTLSClientAuthValue = "on"
	ZonesTLSClientAuthValueOff ZonesTLSClientAuthValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesTLSClientAuthEditable bool

const (
	ZonesTLSClientAuthEditableTrue  ZonesTLSClientAuthEditable = true
	ZonesTLSClientAuthEditableFalse ZonesTLSClientAuthEditable = false
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZonesTLSClientAuthParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesTLSClientAuthID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesTLSClientAuthValue] `json:"value,required"`
}

func (r ZonesTLSClientAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesTLSClientAuthParam) implementsZoneSettingEditParamsItem() {}

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
	Result ZonesTLSClientAuth                               `json:"result"`
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
	Result ZonesTLSClientAuth                              `json:"result"`
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
