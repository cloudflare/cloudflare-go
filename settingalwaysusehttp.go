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

// SettingAlwaysUseHTTPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingAlwaysUseHTTPService]
// method instead.
type SettingAlwaysUseHTTPService struct {
	Options []option.RequestOption
}

// NewSettingAlwaysUseHTTPService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingAlwaysUseHTTPService(opts ...option.RequestOption) (r *SettingAlwaysUseHTTPService) {
	r = &SettingAlwaysUseHTTPService{}
	r.Options = opts
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
func (r *SettingAlwaysUseHTTPService) Update(ctx context.Context, zoneID string, body SettingAlwaysUseHTTPUpdateParams, opts ...option.RequestOption) (res *SettingAlwaysUseHTTPUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAlwaysUseHTTPUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_use_https", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
func (r *SettingAlwaysUseHTTPService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingAlwaysUseHTTPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAlwaysUseHTTPGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_use_https", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type SettingAlwaysUseHTTPUpdateResponse struct {
	// ID of the zone setting.
	ID SettingAlwaysUseHTTPUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingAlwaysUseHTTPUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingAlwaysUseHTTPUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingAlwaysUseHTTPUpdateResponseJSON `json:"-"`
}

// settingAlwaysUseHTTPUpdateResponseJSON contains the JSON metadata for the struct
// [SettingAlwaysUseHTTPUpdateResponse]
type settingAlwaysUseHTTPUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingAlwaysUseHTTPUpdateResponseID string

const (
	SettingAlwaysUseHTTPUpdateResponseIDAlwaysUseHTTPs SettingAlwaysUseHTTPUpdateResponseID = "always_use_https"
)

// Current value of the zone setting.
type SettingAlwaysUseHTTPUpdateResponseValue string

const (
	SettingAlwaysUseHTTPUpdateResponseValueOn  SettingAlwaysUseHTTPUpdateResponseValue = "on"
	SettingAlwaysUseHTTPUpdateResponseValueOff SettingAlwaysUseHTTPUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingAlwaysUseHTTPUpdateResponseEditable bool

const (
	SettingAlwaysUseHTTPUpdateResponseEditableTrue  SettingAlwaysUseHTTPUpdateResponseEditable = true
	SettingAlwaysUseHTTPUpdateResponseEditableFalse SettingAlwaysUseHTTPUpdateResponseEditable = false
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type SettingAlwaysUseHTTPGetResponse struct {
	// ID of the zone setting.
	ID SettingAlwaysUseHTTPGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingAlwaysUseHTTPGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingAlwaysUseHTTPGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingAlwaysUseHTTPGetResponseJSON `json:"-"`
}

// settingAlwaysUseHTTPGetResponseJSON contains the JSON metadata for the struct
// [SettingAlwaysUseHTTPGetResponse]
type settingAlwaysUseHTTPGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingAlwaysUseHTTPGetResponseID string

const (
	SettingAlwaysUseHTTPGetResponseIDAlwaysUseHTTPs SettingAlwaysUseHTTPGetResponseID = "always_use_https"
)

// Current value of the zone setting.
type SettingAlwaysUseHTTPGetResponseValue string

const (
	SettingAlwaysUseHTTPGetResponseValueOn  SettingAlwaysUseHTTPGetResponseValue = "on"
	SettingAlwaysUseHTTPGetResponseValueOff SettingAlwaysUseHTTPGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingAlwaysUseHTTPGetResponseEditable bool

const (
	SettingAlwaysUseHTTPGetResponseEditableTrue  SettingAlwaysUseHTTPGetResponseEditable = true
	SettingAlwaysUseHTTPGetResponseEditableFalse SettingAlwaysUseHTTPGetResponseEditable = false
)

type SettingAlwaysUseHTTPUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingAlwaysUseHTTPUpdateParamsValue] `json:"value,required"`
}

func (r SettingAlwaysUseHTTPUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingAlwaysUseHTTPUpdateParamsValue string

const (
	SettingAlwaysUseHTTPUpdateParamsValueOn  SettingAlwaysUseHTTPUpdateParamsValue = "on"
	SettingAlwaysUseHTTPUpdateParamsValueOff SettingAlwaysUseHTTPUpdateParamsValue = "off"
)

type SettingAlwaysUseHTTPUpdateResponseEnvelope struct {
	Errors   []SettingAlwaysUseHTTPUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAlwaysUseHTTPUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result SettingAlwaysUseHTTPUpdateResponse             `json:"result"`
	JSON   settingAlwaysUseHTTPUpdateResponseEnvelopeJSON `json:"-"`
}

// settingAlwaysUseHTTPUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingAlwaysUseHTTPUpdateResponseEnvelope]
type settingAlwaysUseHTTPUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysUseHTTPUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingAlwaysUseHTTPUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAlwaysUseHTTPUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingAlwaysUseHTTPUpdateResponseEnvelopeErrors]
type settingAlwaysUseHTTPUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysUseHTTPUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingAlwaysUseHTTPUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAlwaysUseHTTPUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingAlwaysUseHTTPUpdateResponseEnvelopeMessages]
type settingAlwaysUseHTTPUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysUseHTTPGetResponseEnvelope struct {
	Errors   []SettingAlwaysUseHTTPGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAlwaysUseHTTPGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result SettingAlwaysUseHTTPGetResponse             `json:"result"`
	JSON   settingAlwaysUseHTTPGetResponseEnvelopeJSON `json:"-"`
}

// settingAlwaysUseHTTPGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAlwaysUseHTTPGetResponseEnvelope]
type settingAlwaysUseHTTPGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysUseHTTPGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingAlwaysUseHTTPGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAlwaysUseHTTPGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingAlwaysUseHTTPGetResponseEnvelopeErrors]
type settingAlwaysUseHTTPGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAlwaysUseHTTPGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingAlwaysUseHTTPGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAlwaysUseHTTPGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAlwaysUseHTTPGetResponseEnvelopeMessages]
type settingAlwaysUseHTTPGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysUseHTTPGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
