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
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SettingProxyReadTimeoutService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingProxyReadTimeoutService] method instead.
type SettingProxyReadTimeoutService struct {
	Options []option.RequestOption
}

// NewSettingProxyReadTimeoutService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingProxyReadTimeoutService(opts ...option.RequestOption) (r *SettingProxyReadTimeoutService) {
	r = &SettingProxyReadTimeoutService{}
	r.Options = opts
	return
}

// Maximum time between two read operations from origin.
func (r *SettingProxyReadTimeoutService) Edit(ctx context.Context, params SettingProxyReadTimeoutEditParams, opts ...option.RequestOption) (res *ProxyReadTimeout, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingProxyReadTimeoutEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/proxy_read_timeout", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Maximum time between two read operations from origin.
func (r *SettingProxyReadTimeoutService) Get(ctx context.Context, query SettingProxyReadTimeoutGetParams, opts ...option.RequestOption) (res *ProxyReadTimeout, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingProxyReadTimeoutGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/proxy_read_timeout", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Maximum time between two read operations from origin.
type ProxyReadTimeout struct {
	// ID of the zone setting.
	ID ProxyReadTimeoutID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time            `json:"modified_on,nullable" format:"date-time"`
	JSON       proxyReadTimeoutJSON `json:"-"`
}

// proxyReadTimeoutJSON contains the JSON metadata for the struct
// [ProxyReadTimeout]
type proxyReadTimeoutJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r proxyReadTimeoutJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ProxyReadTimeoutID string

const (
	ProxyReadTimeoutIDProxyReadTimeout ProxyReadTimeoutID = "proxy_read_timeout"
)

func (r ProxyReadTimeoutID) IsKnown() bool {
	switch r {
	case ProxyReadTimeoutIDProxyReadTimeout:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ProxyReadTimeoutEditable bool

const (
	ProxyReadTimeoutEditableTrue  ProxyReadTimeoutEditable = true
	ProxyReadTimeoutEditableFalse ProxyReadTimeoutEditable = false
)

func (r ProxyReadTimeoutEditable) IsKnown() bool {
	switch r {
	case ProxyReadTimeoutEditableTrue, ProxyReadTimeoutEditableFalse:
		return true
	}
	return false
}

// Maximum time between two read operations from origin.
type ProxyReadTimeoutParam struct {
	// ID of the zone setting.
	ID param.Field[ProxyReadTimeoutID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[float64] `json:"value,required"`
}

func (r ProxyReadTimeoutParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingProxyReadTimeoutEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Maximum time between two read operations from origin.
	Value param.Field[ProxyReadTimeoutParam] `json:"value,required"`
}

func (r SettingProxyReadTimeoutEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingProxyReadTimeoutEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Maximum time between two read operations from origin.
	Result ProxyReadTimeout                                `json:"result"`
	JSON   settingProxyReadTimeoutEditResponseEnvelopeJSON `json:"-"`
}

// settingProxyReadTimeoutEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingProxyReadTimeoutEditResponseEnvelope]
type settingProxyReadTimeoutEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingProxyReadTimeoutEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingProxyReadTimeoutGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingProxyReadTimeoutGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Maximum time between two read operations from origin.
	Result ProxyReadTimeout                               `json:"result"`
	JSON   settingProxyReadTimeoutGetResponseEnvelopeJSON `json:"-"`
}

// settingProxyReadTimeoutGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingProxyReadTimeoutGetResponseEnvelope]
type settingProxyReadTimeoutGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingProxyReadTimeoutGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
