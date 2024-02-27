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

// ZoneSettingCipherService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingCipherService] method
// instead.
type ZoneSettingCipherService struct {
	Options []option.RequestOption
}

// NewZoneSettingCipherService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingCipherService(opts ...option.RequestOption) (r *ZoneSettingCipherService) {
	r = &ZoneSettingCipherService{}
	r.Options = opts
	return
}

// Changes ciphers setting.
func (r *ZoneSettingCipherService) Edit(ctx context.Context, params ZoneSettingCipherEditParams, opts ...option.RequestOption) (res *ZoneSettingCipherEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingCipherEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ciphers", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets ciphers setting.
func (r *ZoneSettingCipherService) Get(ctx context.Context, query ZoneSettingCipherGetParams, opts ...option.RequestOption) (res *ZoneSettingCipherGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingCipherGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ciphers", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingCipherEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingCipherEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value []string `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingCipherEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingCipherEditResponseJSON `json:"-"`
}

// zoneSettingCipherEditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingCipherEditResponse]
type zoneSettingCipherEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingCipherEditResponseID string

const (
	ZoneSettingCipherEditResponseIDCiphers ZoneSettingCipherEditResponseID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingCipherEditResponseEditable bool

const (
	ZoneSettingCipherEditResponseEditableTrue  ZoneSettingCipherEditResponseEditable = true
	ZoneSettingCipherEditResponseEditableFalse ZoneSettingCipherEditResponseEditable = false
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingCipherGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingCipherGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value []string `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingCipherGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingCipherGetResponseJSON `json:"-"`
}

// zoneSettingCipherGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingCipherGetResponse]
type zoneSettingCipherGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingCipherGetResponseID string

const (
	ZoneSettingCipherGetResponseIDCiphers ZoneSettingCipherGetResponseID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingCipherGetResponseEditable bool

const (
	ZoneSettingCipherGetResponseEditableTrue  ZoneSettingCipherGetResponseEditable = true
	ZoneSettingCipherGetResponseEditableFalse ZoneSettingCipherGetResponseEditable = false
)

type ZoneSettingCipherEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[[]string] `json:"value,required"`
}

func (r ZoneSettingCipherEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingCipherEditResponseEnvelope struct {
	Errors   []ZoneSettingCipherEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingCipherEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Result ZoneSettingCipherEditResponse             `json:"result"`
	JSON   zoneSettingCipherEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingCipherEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingCipherEditResponseEnvelope]
type zoneSettingCipherEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCipherEditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingCipherEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingCipherEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingCipherEditResponseEnvelopeErrors]
type zoneSettingCipherEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCipherEditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingCipherEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingCipherEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingCipherEditResponseEnvelopeMessages]
type zoneSettingCipherEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCipherGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingCipherGetResponseEnvelope struct {
	Errors   []ZoneSettingCipherGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingCipherGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Result ZoneSettingCipherGetResponse             `json:"result"`
	JSON   zoneSettingCipherGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingCipherGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingCipherGetResponseEnvelope]
type zoneSettingCipherGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCipherGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingCipherGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingCipherGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingCipherGetResponseEnvelopeErrors]
type zoneSettingCipherGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCipherGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingCipherGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingCipherGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingCipherGetResponseEnvelopeMessages]
type zoneSettingCipherGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
