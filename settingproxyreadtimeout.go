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
func (r *SettingProxyReadTimeoutService) Update(ctx context.Context, zoneID string, body SettingProxyReadTimeoutUpdateParams, opts ...option.RequestOption) (res *SettingProxyReadTimeoutUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingProxyReadTimeoutUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/proxy_read_timeout", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Maximum time between two read operations from origin.
func (r *SettingProxyReadTimeoutService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingProxyReadTimeoutGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingProxyReadTimeoutGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/proxy_read_timeout", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Maximum time between two read operations from origin.
type SettingProxyReadTimeoutUpdateResponse struct {
	// ID of the zone setting.
	ID SettingProxyReadTimeoutUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingProxyReadTimeoutUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingProxyReadTimeoutUpdateResponseJSON `json:"-"`
}

// settingProxyReadTimeoutUpdateResponseJSON contains the JSON metadata for the
// struct [SettingProxyReadTimeoutUpdateResponse]
type settingProxyReadTimeoutUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingProxyReadTimeoutUpdateResponseID string

const (
	SettingProxyReadTimeoutUpdateResponseIDProxyReadTimeout SettingProxyReadTimeoutUpdateResponseID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingProxyReadTimeoutUpdateResponseEditable bool

const (
	SettingProxyReadTimeoutUpdateResponseEditableTrue  SettingProxyReadTimeoutUpdateResponseEditable = true
	SettingProxyReadTimeoutUpdateResponseEditableFalse SettingProxyReadTimeoutUpdateResponseEditable = false
)

// Maximum time between two read operations from origin.
type SettingProxyReadTimeoutGetResponse struct {
	// ID of the zone setting.
	ID SettingProxyReadTimeoutGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingProxyReadTimeoutGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingProxyReadTimeoutGetResponseJSON `json:"-"`
}

// settingProxyReadTimeoutGetResponseJSON contains the JSON metadata for the struct
// [SettingProxyReadTimeoutGetResponse]
type settingProxyReadTimeoutGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingProxyReadTimeoutGetResponseID string

const (
	SettingProxyReadTimeoutGetResponseIDProxyReadTimeout SettingProxyReadTimeoutGetResponseID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingProxyReadTimeoutGetResponseEditable bool

const (
	SettingProxyReadTimeoutGetResponseEditableTrue  SettingProxyReadTimeoutGetResponseEditable = true
	SettingProxyReadTimeoutGetResponseEditableFalse SettingProxyReadTimeoutGetResponseEditable = false
)

type SettingProxyReadTimeoutUpdateParams struct {
	// Maximum time between two read operations from origin.
	Value param.Field[SettingProxyReadTimeoutUpdateParamsValue] `json:"value,required"`
}

func (r SettingProxyReadTimeoutUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Maximum time between two read operations from origin.
type SettingProxyReadTimeoutUpdateParamsValue struct {
	// ID of the zone setting.
	ID param.Field[SettingProxyReadTimeoutUpdateParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[float64] `json:"value,required"`
}

func (r SettingProxyReadTimeoutUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type SettingProxyReadTimeoutUpdateParamsValueID string

const (
	SettingProxyReadTimeoutUpdateParamsValueIDProxyReadTimeout SettingProxyReadTimeoutUpdateParamsValueID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingProxyReadTimeoutUpdateParamsValueEditable bool

const (
	SettingProxyReadTimeoutUpdateParamsValueEditableTrue  SettingProxyReadTimeoutUpdateParamsValueEditable = true
	SettingProxyReadTimeoutUpdateParamsValueEditableFalse SettingProxyReadTimeoutUpdateParamsValueEditable = false
)

type SettingProxyReadTimeoutUpdateResponseEnvelope struct {
	Errors   []SettingProxyReadTimeoutUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingProxyReadTimeoutUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Maximum time between two read operations from origin.
	Result SettingProxyReadTimeoutUpdateResponse             `json:"result"`
	JSON   settingProxyReadTimeoutUpdateResponseEnvelopeJSON `json:"-"`
}

// settingProxyReadTimeoutUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingProxyReadTimeoutUpdateResponseEnvelope]
type settingProxyReadTimeoutUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingProxyReadTimeoutUpdateResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingProxyReadTimeoutUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingProxyReadTimeoutUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingProxyReadTimeoutUpdateResponseEnvelopeErrors]
type settingProxyReadTimeoutUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingProxyReadTimeoutUpdateResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    settingProxyReadTimeoutUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingProxyReadTimeoutUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingProxyReadTimeoutUpdateResponseEnvelopeMessages]
type settingProxyReadTimeoutUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingProxyReadTimeoutGetResponseEnvelope struct {
	Errors   []SettingProxyReadTimeoutGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingProxyReadTimeoutGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Maximum time between two read operations from origin.
	Result SettingProxyReadTimeoutGetResponse             `json:"result"`
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

type SettingProxyReadTimeoutGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingProxyReadTimeoutGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingProxyReadTimeoutGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingProxyReadTimeoutGetResponseEnvelopeErrors]
type settingProxyReadTimeoutGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingProxyReadTimeoutGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingProxyReadTimeoutGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingProxyReadTimeoutGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingProxyReadTimeoutGetResponseEnvelopeMessages]
type settingProxyReadTimeoutGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
