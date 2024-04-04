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
func (r *SettingTLSClientAuthService) Edit(ctx context.Context, params SettingTLSClientAuthEditParams, opts ...option.RequestOption) (res *ZoneSettingTLSClientAuth, err error) {
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
func (r *SettingTLSClientAuthService) Get(ctx context.Context, query SettingTLSClientAuthGetParams, opts ...option.RequestOption) (res *ZoneSettingTLSClientAuth, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTLSClientAuthGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/tls_client_auth", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingTLSClientAuth struct {
	// ID of the zone setting.
	ID ZoneSettingTLSClientAuthID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingTLSClientAuthValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingTLSClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingTLSClientAuthJSON `json:"-"`
}

// zoneSettingTLSClientAuthJSON contains the JSON metadata for the struct
// [ZoneSettingTLSClientAuth]
type zoneSettingTLSClientAuthJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingTLSClientAuthJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingTLSClientAuthID string

const (
	ZoneSettingTLSClientAuthIDTLSClientAuth ZoneSettingTLSClientAuthID = "tls_client_auth"
)

func (r ZoneSettingTLSClientAuthID) IsKnown() bool {
	switch r {
	case ZoneSettingTLSClientAuthIDTLSClientAuth:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingTLSClientAuthValue string

const (
	ZoneSettingTLSClientAuthValueOn  ZoneSettingTLSClientAuthValue = "on"
	ZoneSettingTLSClientAuthValueOff ZoneSettingTLSClientAuthValue = "off"
)

func (r ZoneSettingTLSClientAuthValue) IsKnown() bool {
	switch r {
	case ZoneSettingTLSClientAuthValueOn, ZoneSettingTLSClientAuthValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingTLSClientAuthEditable bool

const (
	ZoneSettingTLSClientAuthEditableTrue  ZoneSettingTLSClientAuthEditable = true
	ZoneSettingTLSClientAuthEditableFalse ZoneSettingTLSClientAuthEditable = false
)

func (r ZoneSettingTLSClientAuthEditable) IsKnown() bool {
	switch r {
	case ZoneSettingTLSClientAuthEditableTrue, ZoneSettingTLSClientAuthEditableFalse:
		return true
	}
	return false
}

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

func (r SettingTLSClientAuthEditParamsValue) IsKnown() bool {
	switch r {
	case SettingTLSClientAuthEditParamsValueOn, SettingTLSClientAuthEditParamsValueOff:
		return true
	}
	return false
}

type SettingTLSClientAuthEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result ZoneSettingTLSClientAuth                     `json:"result"`
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

func (r settingTLSClientAuthEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTLSClientAuthGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingTLSClientAuthGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result ZoneSettingTLSClientAuth                    `json:"result"`
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

func (r settingTLSClientAuthGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
