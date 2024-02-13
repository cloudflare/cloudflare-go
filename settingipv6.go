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

// SettingIPV6Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingIPV6Service] method
// instead.
type SettingIPV6Service struct {
	Options []option.RequestOption
}

// NewSettingIPV6Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingIPV6Service(opts ...option.RequestOption) (r *SettingIPV6Service) {
	r = &SettingIPV6Service{}
	r.Options = opts
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
func (r *SettingIPV6Service) Update(ctx context.Context, zoneID string, body SettingIPV6UpdateParams, opts ...option.RequestOption) (res *SettingIPV6UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingIPV6UpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ipv6", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
func (r *SettingIPV6Service) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingIPV6GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingIPV6GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ipv6", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type SettingIPV6UpdateResponse struct {
	// ID of the zone setting.
	ID SettingIPV6UpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingIPV6UpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingIPV6UpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingIPV6UpdateResponseJSON `json:"-"`
}

// settingIPV6UpdateResponseJSON contains the JSON metadata for the struct
// [SettingIPV6UpdateResponse]
type settingIPV6UpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingIPV6UpdateResponseID string

const (
	SettingIPV6UpdateResponseIDIPV6 SettingIPV6UpdateResponseID = "ipv6"
)

// Current value of the zone setting.
type SettingIPV6UpdateResponseValue string

const (
	SettingIPV6UpdateResponseValueOff SettingIPV6UpdateResponseValue = "off"
	SettingIPV6UpdateResponseValueOn  SettingIPV6UpdateResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingIPV6UpdateResponseEditable bool

const (
	SettingIPV6UpdateResponseEditableTrue  SettingIPV6UpdateResponseEditable = true
	SettingIPV6UpdateResponseEditableFalse SettingIPV6UpdateResponseEditable = false
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type SettingIPV6GetResponse struct {
	// ID of the zone setting.
	ID SettingIPV6GetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingIPV6GetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingIPV6GetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingIPV6GetResponseJSON `json:"-"`
}

// settingIPV6GetResponseJSON contains the JSON metadata for the struct
// [SettingIPV6GetResponse]
type settingIPV6GetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingIPV6GetResponseID string

const (
	SettingIPV6GetResponseIDIPV6 SettingIPV6GetResponseID = "ipv6"
)

// Current value of the zone setting.
type SettingIPV6GetResponseValue string

const (
	SettingIPV6GetResponseValueOff SettingIPV6GetResponseValue = "off"
	SettingIPV6GetResponseValueOn  SettingIPV6GetResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingIPV6GetResponseEditable bool

const (
	SettingIPV6GetResponseEditableTrue  SettingIPV6GetResponseEditable = true
	SettingIPV6GetResponseEditableFalse SettingIPV6GetResponseEditable = false
)

type SettingIPV6UpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingIPV6UpdateParamsValue] `json:"value,required"`
}

func (r SettingIPV6UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingIPV6UpdateParamsValue string

const (
	SettingIPV6UpdateParamsValueOff SettingIPV6UpdateParamsValue = "off"
	SettingIPV6UpdateParamsValueOn  SettingIPV6UpdateParamsValue = "on"
)

type SettingIPV6UpdateResponseEnvelope struct {
	Errors   []SettingIPV6UpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingIPV6UpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IPv6 on all subdomains that are Cloudflare enabled.
	// (https://support.cloudflare.com/hc/en-us/articles/200168586).
	Result SettingIPV6UpdateResponse             `json:"result"`
	JSON   settingIPV6UpdateResponseEnvelopeJSON `json:"-"`
}

// settingIPV6UpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingIPV6UpdateResponseEnvelope]
type settingIPV6UpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6UpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPV6UpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingIPV6UpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingIPV6UpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingIPV6UpdateResponseEnvelopeErrors]
type settingIPV6UpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6UpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPV6UpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingIPV6UpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingIPV6UpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingIPV6UpdateResponseEnvelopeMessages]
type settingIPV6UpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6UpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPV6GetResponseEnvelope struct {
	Errors   []SettingIPV6GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingIPV6GetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IPv6 on all subdomains that are Cloudflare enabled.
	// (https://support.cloudflare.com/hc/en-us/articles/200168586).
	Result SettingIPV6GetResponse             `json:"result"`
	JSON   settingIPV6GetResponseEnvelopeJSON `json:"-"`
}

// settingIPV6GetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingIPV6GetResponseEnvelope]
type settingIPV6GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPV6GetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    settingIPV6GetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingIPV6GetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingIPV6GetResponseEnvelopeErrors]
type settingIPV6GetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6GetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPV6GetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingIPV6GetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingIPV6GetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingIPV6GetResponseEnvelopeMessages]
type settingIPV6GetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6GetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
