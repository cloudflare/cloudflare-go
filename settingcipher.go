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

// SettingCipherService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingCipherService] method
// instead.
type SettingCipherService struct {
	Options []option.RequestOption
}

// NewSettingCipherService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingCipherService(opts ...option.RequestOption) (r *SettingCipherService) {
	r = &SettingCipherService{}
	r.Options = opts
	return
}

// Changes ciphers setting.
func (r *SettingCipherService) Update(ctx context.Context, zoneID string, body SettingCipherUpdateParams, opts ...option.RequestOption) (res *SettingCipherUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingCipherUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ciphers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets ciphers setting.
func (r *SettingCipherService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingCipherGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingCipherGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ciphers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type SettingCipherUpdateResponse struct {
	// ID of the zone setting.
	ID SettingCipherUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value []string `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingCipherUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingCipherUpdateResponseJSON `json:"-"`
}

// settingCipherUpdateResponseJSON contains the JSON metadata for the struct
// [SettingCipherUpdateResponse]
type settingCipherUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCipherUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingCipherUpdateResponseID string

const (
	SettingCipherUpdateResponseIDCiphers SettingCipherUpdateResponseID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingCipherUpdateResponseEditable bool

const (
	SettingCipherUpdateResponseEditableTrue  SettingCipherUpdateResponseEditable = true
	SettingCipherUpdateResponseEditableFalse SettingCipherUpdateResponseEditable = false
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type SettingCipherGetResponse struct {
	// ID of the zone setting.
	ID SettingCipherGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value []string `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingCipherGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingCipherGetResponseJSON `json:"-"`
}

// settingCipherGetResponseJSON contains the JSON metadata for the struct
// [SettingCipherGetResponse]
type settingCipherGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCipherGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingCipherGetResponseID string

const (
	SettingCipherGetResponseIDCiphers SettingCipherGetResponseID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingCipherGetResponseEditable bool

const (
	SettingCipherGetResponseEditableTrue  SettingCipherGetResponseEditable = true
	SettingCipherGetResponseEditableFalse SettingCipherGetResponseEditable = false
)

type SettingCipherUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[[]string] `json:"value,required"`
}

func (r SettingCipherUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingCipherUpdateResponseEnvelope struct {
	Errors   []SettingCipherUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingCipherUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Result SettingCipherUpdateResponse             `json:"result"`
	JSON   settingCipherUpdateResponseEnvelopeJSON `json:"-"`
}

// settingCipherUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingCipherUpdateResponseEnvelope]
type settingCipherUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCipherUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingCipherUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingCipherUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingCipherUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingCipherUpdateResponseEnvelopeErrors]
type settingCipherUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCipherUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingCipherUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    settingCipherUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingCipherUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingCipherUpdateResponseEnvelopeMessages]
type settingCipherUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCipherUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingCipherGetResponseEnvelope struct {
	Errors   []SettingCipherGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingCipherGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Result SettingCipherGetResponse             `json:"result"`
	JSON   settingCipherGetResponseEnvelopeJSON `json:"-"`
}

// settingCipherGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingCipherGetResponseEnvelope]
type settingCipherGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCipherGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingCipherGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingCipherGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingCipherGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingCipherGetResponseEnvelopeErrors]
type settingCipherGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCipherGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingCipherGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingCipherGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingCipherGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingCipherGetResponseEnvelopeMessages]
type settingCipherGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCipherGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
