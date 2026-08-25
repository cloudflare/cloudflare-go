// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// NELService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNELService] method instead.
type NELService struct {
	Options []option.RequestOption
}

// NewNELService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNELService(opts ...option.RequestOption) (r *NELService) {
	r = &NELService{}
	r.Options = opts
	return
}

// Updates the Network Error Logging (NEL) setting for a zone. Requires the NEL
// product feature to be enabled for the zone. The setting controls whether
// browsers report network errors to Cloudflare's NEL endpoint.
func (r *NELService) Edit(ctx context.Context, params NELEditParams, opts ...option.RequestOption) (res *Setting, err error) {
	var env NELEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/nel", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetches the Network Error Logging (NEL) setting for a zone. NEL allows browsers
// to report network errors to a configured endpoint. The setting is enabled by
// default for free and pro zones, and disabled by default for business and
// enterprise zones unless the NEL product feature is enabled.
func (r *NELService) Get(ctx context.Context, query NELGetParams, opts ...option.RequestOption) (res *Setting, err error) {
	var env NELGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/nel", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A zone-scoped NEL configuration setting.
type Setting struct {
	// Zone setting identifier.
	ID SettingID `json:"id" api:"required"`
	// Whether the setting is editable. This is false when the zone's plan does not
	// include NEL or the NEL product feature is not enabled.
	Editable bool `json:"editable" api:"required"`
	// When the setting was last modified. A zero value (0001-01-01T00:00:00Z)
	// indicates the setting has never been explicitly set and is using the default
	// value.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The NEL configuration value.
	Value SettingValue `json:"value" api:"required"`
	JSON  settingJSON  `json:"-"`
}

// settingJSON contains the JSON metadata for the struct [Setting]
type settingJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Setting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingJSON) RawJSON() string {
	return r.raw
}

// Zone setting identifier.
type SettingID string

const (
	SettingIDNEL SettingID = "nel"
)

func (r SettingID) IsKnown() bool {
	switch r {
	case SettingIDNEL:
		return true
	}
	return false
}

// The NEL configuration value.
type SettingValue struct {
	// Whether Network Error Logging is enabled for the zone. When enabled, browsers
	// report network errors to Cloudflare's NEL endpoint.
	Enabled bool             `json:"enabled" api:"required"`
	JSON    settingValueJSON `json:"-"`
}

// settingValueJSON contains the JSON metadata for the struct [SettingValue]
type settingValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingValueJSON) RawJSON() string {
	return r.raw
}

type NELEditParams struct {
	// Identifier of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// The NEL configuration value.
	Value param.Field[NELEditParamsValue] `json:"value" api:"required"`
}

func (r NELEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The NEL configuration value.
type NELEditParamsValue struct {
	// Whether Network Error Logging is enabled for the zone. When enabled, browsers
	// report network errors to Cloudflare's NEL endpoint.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r NELEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response envelope for a single NEL setting.
type NELEditResponseEnvelope struct {
	Errors   []NELEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []NELEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// A zone-scoped NEL configuration setting.
	Result Setting `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                        `json:"success" api:"required"`
	JSON    nelEditResponseEnvelopeJSON `json:"-"`
}

// nelEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [NELEditResponseEnvelope]
type nelEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NELEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nelEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// An API response message or error.
type NELEditResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string                            `json:"message" api:"required"`
	JSON    nelEditResponseEnvelopeErrorsJSON `json:"-"`
}

// nelEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [NELEditResponseEnvelopeErrors]
type nelEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NELEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nelEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// An API response message or error.
type NELEditResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string                              `json:"message" api:"required"`
	JSON    nelEditResponseEnvelopeMessagesJSON `json:"-"`
}

// nelEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [NELEditResponseEnvelopeMessages]
type nelEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NELEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nelEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type NELGetParams struct {
	// Identifier of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

// Response envelope for a single NEL setting.
type NELGetResponseEnvelope struct {
	Errors   []NELGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []NELGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// A zone-scoped NEL configuration setting.
	Result Setting `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                       `json:"success" api:"required"`
	JSON    nelGetResponseEnvelopeJSON `json:"-"`
}

// nelGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [NELGetResponseEnvelope]
type nelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NELGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// An API response message or error.
type NELGetResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string                           `json:"message" api:"required"`
	JSON    nelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// nelGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [NELGetResponseEnvelopeErrors]
type nelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NELGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nelGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// An API response message or error.
type NELGetResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string                             `json:"message" api:"required"`
	JSON    nelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// nelGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [NELGetResponseEnvelopeMessages]
type nelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NELGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nelGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
