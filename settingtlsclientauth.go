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
func (r *SettingTLSClientAuthService) Update(ctx context.Context, zoneID string, body SettingTLSClientAuthUpdateParams, opts ...option.RequestOption) (res *SettingTLSClientAuthUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTLSClientAuthUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/tls_client_auth", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
func (r *SettingTLSClientAuthService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingTLSClientAuthGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTLSClientAuthGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/tls_client_auth", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type SettingTLSClientAuthUpdateResponse struct {
	// ID of the zone setting.
	ID SettingTLSClientAuthUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingTLSClientAuthUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingTLSClientAuthUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingTLSClientAuthUpdateResponseJSON `json:"-"`
}

// settingTLSClientAuthUpdateResponseJSON contains the JSON metadata for the struct
// [SettingTLSClientAuthUpdateResponse]
type settingTLSClientAuthUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSClientAuthUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingTLSClientAuthUpdateResponseID string

const (
	SettingTLSClientAuthUpdateResponseIDTLSClientAuth SettingTLSClientAuthUpdateResponseID = "tls_client_auth"
)

// Current value of the zone setting.
type SettingTLSClientAuthUpdateResponseValue string

const (
	SettingTLSClientAuthUpdateResponseValueOn  SettingTLSClientAuthUpdateResponseValue = "on"
	SettingTLSClientAuthUpdateResponseValueOff SettingTLSClientAuthUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingTLSClientAuthUpdateResponseEditable bool

const (
	SettingTLSClientAuthUpdateResponseEditableTrue  SettingTLSClientAuthUpdateResponseEditable = true
	SettingTLSClientAuthUpdateResponseEditableFalse SettingTLSClientAuthUpdateResponseEditable = false
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

type SettingTLSClientAuthUpdateParams struct {
	// value of the zone setting.
	Value param.Field[SettingTLSClientAuthUpdateParamsValue] `json:"value,required"`
}

func (r SettingTLSClientAuthUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// value of the zone setting.
type SettingTLSClientAuthUpdateParamsValue string

const (
	SettingTLSClientAuthUpdateParamsValueOn  SettingTLSClientAuthUpdateParamsValue = "on"
	SettingTLSClientAuthUpdateParamsValueOff SettingTLSClientAuthUpdateParamsValue = "off"
)

type SettingTLSClientAuthUpdateResponseEnvelope struct {
	Errors   []SettingTLSClientAuthUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTLSClientAuthUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result SettingTLSClientAuthUpdateResponse             `json:"result"`
	JSON   settingTLSClientAuthUpdateResponseEnvelopeJSON `json:"-"`
}

// settingTLSClientAuthUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingTLSClientAuthUpdateResponseEnvelope]
type settingTLSClientAuthUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSClientAuthUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTLSClientAuthUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingTLSClientAuthUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingTLSClientAuthUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingTLSClientAuthUpdateResponseEnvelopeErrors]
type settingTLSClientAuthUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSClientAuthUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTLSClientAuthUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingTLSClientAuthUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingTLSClientAuthUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingTLSClientAuthUpdateResponseEnvelopeMessages]
type settingTLSClientAuthUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSClientAuthUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
