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

// SettingAlwaysUseHTTPSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingAlwaysUseHTTPSService]
// method instead.
type SettingAlwaysUseHTTPSService struct {
	Options []option.RequestOption
}

// NewSettingAlwaysUseHTTPSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingAlwaysUseHTTPSService(opts ...option.RequestOption) (r *SettingAlwaysUseHTTPSService) {
	r = &SettingAlwaysUseHTTPSService{}
	r.Options = opts
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
func (r *SettingAlwaysUseHTTPSService) Edit(ctx context.Context, params SettingAlwaysUseHTTPSEditParams, opts ...option.RequestOption) (res *SettingAlwaysUseHTTPSEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAlwaysUseHTTPSEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_use_https", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
func (r *SettingAlwaysUseHTTPSService) Get(ctx context.Context, query SettingAlwaysUseHTTPSGetParams, opts ...option.RequestOption) (res *SettingAlwaysUseHTTPSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAlwaysUseHTTPSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_use_https", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type SettingAlwaysUseHTTPSEditResponse struct {
	// ID of the zone setting.
	ID SettingAlwaysUseHTTPSEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingAlwaysUseHTTPSEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingAlwaysUseHTTPSEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingAlwaysUseHTTPSEditResponseJSON `json:"-"`
}

// settingAlwaysUseHTTPSEditResponseJSON contains the JSON metadata for the struct
// [SettingAlwaysUseHTTPSEditResponse]
type settingAlwaysUseHTTPSEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingAlwaysUseHTTPSEditResponseID string

const (
	SettingAlwaysUseHTTPSEditResponseIDAlwaysUseHTTPS SettingAlwaysUseHTTPSEditResponseID = "always_use_https"
)

// Current value of the zone setting.
type SettingAlwaysUseHTTPSEditResponseValue string

const (
	SettingAlwaysUseHTTPSEditResponseValueOn  SettingAlwaysUseHTTPSEditResponseValue = "on"
	SettingAlwaysUseHTTPSEditResponseValueOff SettingAlwaysUseHTTPSEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingAlwaysUseHTTPSEditResponseEditable bool

const (
	SettingAlwaysUseHTTPSEditResponseEditableTrue  SettingAlwaysUseHTTPSEditResponseEditable = true
	SettingAlwaysUseHTTPSEditResponseEditableFalse SettingAlwaysUseHTTPSEditResponseEditable = false
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type SettingAlwaysUseHTTPSGetResponse struct {
	// ID of the zone setting.
	ID SettingAlwaysUseHTTPSGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingAlwaysUseHTTPSGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingAlwaysUseHTTPSGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingAlwaysUseHTTPSGetResponseJSON `json:"-"`
}

// settingAlwaysUseHTTPSGetResponseJSON contains the JSON metadata for the struct
// [SettingAlwaysUseHTTPSGetResponse]
type settingAlwaysUseHTTPSGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingAlwaysUseHTTPSGetResponseID string

const (
	SettingAlwaysUseHTTPSGetResponseIDAlwaysUseHTTPS SettingAlwaysUseHTTPSGetResponseID = "always_use_https"
)

// Current value of the zone setting.
type SettingAlwaysUseHTTPSGetResponseValue string

const (
	SettingAlwaysUseHTTPSGetResponseValueOn  SettingAlwaysUseHTTPSGetResponseValue = "on"
	SettingAlwaysUseHTTPSGetResponseValueOff SettingAlwaysUseHTTPSGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingAlwaysUseHTTPSGetResponseEditable bool

const (
	SettingAlwaysUseHTTPSGetResponseEditableTrue  SettingAlwaysUseHTTPSGetResponseEditable = true
	SettingAlwaysUseHTTPSGetResponseEditableFalse SettingAlwaysUseHTTPSGetResponseEditable = false
)

type SettingAlwaysUseHTTPSEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingAlwaysUseHTTPSEditParamsValue] `json:"value,required"`
}

func (r SettingAlwaysUseHTTPSEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingAlwaysUseHTTPSEditParamsValue string

const (
	SettingAlwaysUseHTTPSEditParamsValueOn  SettingAlwaysUseHTTPSEditParamsValue = "on"
	SettingAlwaysUseHTTPSEditParamsValueOff SettingAlwaysUseHTTPSEditParamsValue = "off"
)

type SettingAlwaysUseHTTPSEditResponseEnvelope struct {
	Errors   []SettingAlwaysUseHTTPSEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAlwaysUseHTTPSEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result SettingAlwaysUseHTTPSEditResponse             `json:"result"`
	JSON   settingAlwaysUseHTTPSEditResponseEnvelopeJSON `json:"-"`
}

// settingAlwaysUseHTTPSEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAlwaysUseHTTPSEditResponseEnvelope]
type settingAlwaysUseHTTPSEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysUseHTTPSEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingAlwaysUseHTTPSEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAlwaysUseHTTPSEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingAlwaysUseHTTPSEditResponseEnvelopeErrors]
type settingAlwaysUseHTTPSEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysUseHTTPSEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingAlwaysUseHTTPSEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAlwaysUseHTTPSEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAlwaysUseHTTPSEditResponseEnvelopeMessages]
type settingAlwaysUseHTTPSEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysUseHTTPSGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingAlwaysUseHTTPSGetResponseEnvelope struct {
	Errors   []SettingAlwaysUseHTTPSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAlwaysUseHTTPSGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result SettingAlwaysUseHTTPSGetResponse             `json:"result"`
	JSON   settingAlwaysUseHTTPSGetResponseEnvelopeJSON `json:"-"`
}

// settingAlwaysUseHTTPSGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAlwaysUseHTTPSGetResponseEnvelope]
type settingAlwaysUseHTTPSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysUseHTTPSGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingAlwaysUseHTTPSGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAlwaysUseHTTPSGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingAlwaysUseHTTPSGetResponseEnvelopeErrors]
type settingAlwaysUseHTTPSGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysUseHTTPSGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingAlwaysUseHTTPSGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAlwaysUseHTTPSGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAlwaysUseHTTPSGetResponseEnvelopeMessages]
type settingAlwaysUseHTTPSGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
