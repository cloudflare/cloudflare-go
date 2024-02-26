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

// SettingTLSClientAuthService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingTLSClientAuthService]
// method instead.
type SettingTLSClientAuthService struct {
	Options []option.RequestOption
}

// NewSettingTLSClientAuthService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingTLSClientAuthService(opts ...option.RequestOption) (r *SettingTLSClientAuthService) {
	r = &SettingTLSClientAuthService{}
	r.Options = opts
	return
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
func (r *SettingTLSClientAuthService) Edit(ctx context.Context, params SettingTLSClientAuthEditParams, opts ...option.RequestOption) (res *SettingTLSClientAuthEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTLSClientAuthEditResponseEnvelope
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
func (r *SettingTLSClientAuthService) Get(ctx context.Context, query SettingTLSClientAuthGetParams, opts ...option.RequestOption) (res *SettingTLSClientAuthGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTLSClientAuthGetResponseEnvelope
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
type SettingTLSClientAuthEditResponse struct {
	// ID of the zone setting.
	ID SettingTLSClientAuthEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingTLSClientAuthEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingTLSClientAuthEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingTLSClientAuthEditResponseJSON `json:"-"`
}

// settingTLSClientAuthEditResponseJSON contains the JSON metadata for the struct
// [SettingTLSClientAuthEditResponse]
type settingTLSClientAuthEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSClientAuthEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingTLSClientAuthEditResponseID string

const (
	SettingTLSClientAuthEditResponseIDTLSClientAuth SettingTLSClientAuthEditResponseID = "tls_client_auth"
)

// Current value of the zone setting.
type SettingTLSClientAuthEditResponseValue string

const (
	SettingTLSClientAuthEditResponseValueOn  SettingTLSClientAuthEditResponseValue = "on"
	SettingTLSClientAuthEditResponseValueOff SettingTLSClientAuthEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingTLSClientAuthEditResponseEditable bool

const (
	SettingTLSClientAuthEditResponseEditableTrue  SettingTLSClientAuthEditResponseEditable = true
	SettingTLSClientAuthEditResponseEditableFalse SettingTLSClientAuthEditResponseEditable = false
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type SettingTLSClientAuthGetResponse struct {
	// ID of the zone setting.
	ID SettingTLSClientAuthGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingTLSClientAuthGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingTLSClientAuthGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingTLSClientAuthGetResponseJSON `json:"-"`
}

// settingTLSClientAuthGetResponseJSON contains the JSON metadata for the struct
// [SettingTLSClientAuthGetResponse]
type settingTLSClientAuthGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSClientAuthGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingTLSClientAuthGetResponseID string

const (
	SettingTLSClientAuthGetResponseIDTLSClientAuth SettingTLSClientAuthGetResponseID = "tls_client_auth"
)

// Current value of the zone setting.
type SettingTLSClientAuthGetResponseValue string

const (
	SettingTLSClientAuthGetResponseValueOn  SettingTLSClientAuthGetResponseValue = "on"
	SettingTLSClientAuthGetResponseValueOff SettingTLSClientAuthGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingTLSClientAuthGetResponseEditable bool

const (
	SettingTLSClientAuthGetResponseEditableTrue  SettingTLSClientAuthGetResponseEditable = true
	SettingTLSClientAuthGetResponseEditableFalse SettingTLSClientAuthGetResponseEditable = false
)

type SettingTLSClientAuthEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// value of the zone setting.
	Value param.Field[SettingTLSClientAuthEditParamsValue] `json:"value,required"`
}

func (r SettingTLSClientAuthEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// value of the zone setting.
type SettingTLSClientAuthEditParamsValue string

const (
	SettingTLSClientAuthEditParamsValueOn  SettingTLSClientAuthEditParamsValue = "on"
	SettingTLSClientAuthEditParamsValueOff SettingTLSClientAuthEditParamsValue = "off"
)

type SettingTLSClientAuthEditResponseEnvelope struct {
	Errors   []SettingTLSClientAuthEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTLSClientAuthEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result SettingTLSClientAuthEditResponse             `json:"result"`
	JSON   settingTLSClientAuthEditResponseEnvelopeJSON `json:"-"`
}

// settingTLSClientAuthEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingTLSClientAuthEditResponseEnvelope]
type settingTLSClientAuthEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSClientAuthEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTLSClientAuthEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingTLSClientAuthEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingTLSClientAuthEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingTLSClientAuthEditResponseEnvelopeErrors]
type settingTLSClientAuthEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSClientAuthEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTLSClientAuthEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingTLSClientAuthEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingTLSClientAuthEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingTLSClientAuthEditResponseEnvelopeMessages]
type settingTLSClientAuthEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSClientAuthEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTLSClientAuthGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingTLSClientAuthGetResponseEnvelope struct {
	Errors   []SettingTLSClientAuthGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTLSClientAuthGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result SettingTLSClientAuthGetResponse             `json:"result"`
	JSON   settingTLSClientAuthGetResponseEnvelopeJSON `json:"-"`
}

// settingTLSClientAuthGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingTLSClientAuthGetResponseEnvelope]
type settingTLSClientAuthGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSClientAuthGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTLSClientAuthGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingTLSClientAuthGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingTLSClientAuthGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingTLSClientAuthGetResponseEnvelopeErrors]
type settingTLSClientAuthGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSClientAuthGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTLSClientAuthGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingTLSClientAuthGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingTLSClientAuthGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingTLSClientAuthGetResponseEnvelopeMessages]
type settingTLSClientAuthGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSClientAuthGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
