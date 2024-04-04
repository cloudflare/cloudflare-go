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
func (r *SettingCipherService) Edit(ctx context.Context, params SettingCipherEditParams, opts ...option.RequestOption) (res *ZoneSettingCiphers, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingCipherEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ciphers", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets ciphers setting.
func (r *SettingCipherService) Get(ctx context.Context, query SettingCipherGetParams, opts ...option.RequestOption) (res *ZoneSettingCiphers, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingCipherGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ciphers", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingCiphers struct {
	// ID of the zone setting.
	ID ZoneSettingCiphersID `json:"id,required"`
	// Current value of the zone setting.
	Value []string `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingCiphersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingCiphersJSON `json:"-"`
}

// zoneSettingCiphersJSON contains the JSON metadata for the struct
// [ZoneSettingCiphers]
type zoneSettingCiphersJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCiphers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingCiphersJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingCiphers) implementsZonesSettingEditResponse() {}

func (r ZoneSettingCiphers) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingCiphersID string

const (
	ZoneSettingCiphersIDCiphers ZoneSettingCiphersID = "ciphers"
)

func (r ZoneSettingCiphersID) IsKnown() bool {
	switch r {
	case ZoneSettingCiphersIDCiphers:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingCiphersEditable bool

const (
	ZoneSettingCiphersEditableTrue  ZoneSettingCiphersEditable = true
	ZoneSettingCiphersEditableFalse ZoneSettingCiphersEditable = false
)

func (r ZoneSettingCiphersEditable) IsKnown() bool {
	switch r {
	case ZoneSettingCiphersEditableTrue, ZoneSettingCiphersEditableFalse:
		return true
	}
	return false
}

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingCiphersParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingCiphersID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[[]string] `json:"value,required"`
}

func (r ZoneSettingCiphersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingCiphersParam) implementsZonesSettingEditParamsItem() {}

type SettingCipherEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[[]string] `json:"value,required"`
}

func (r SettingCipherEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingCipherEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Result ZoneSettingCiphers                    `json:"result"`
	JSON   settingCipherEditResponseEnvelopeJSON `json:"-"`
}

// settingCipherEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingCipherEditResponseEnvelope]
type settingCipherEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCipherEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingCipherEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingCipherGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingCipherGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Result ZoneSettingCiphers                   `json:"result"`
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

func (r settingCipherGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
