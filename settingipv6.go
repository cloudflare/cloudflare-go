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
func (r *SettingIPV6Service) Edit(ctx context.Context, params SettingIPV6EditParams, opts ...option.RequestOption) (res *SettingIPV6EditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingIPV6EditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ipv6", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
func (r *SettingIPV6Service) Get(ctx context.Context, query SettingIPV6GetParams, opts ...option.RequestOption) (res *SettingIPV6GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingIPV6GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ipv6", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type SettingIPV6EditResponse struct {
	// ID of the zone setting.
	ID SettingIPV6EditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingIPV6EditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingIPV6EditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingIPV6EditResponseJSON `json:"-"`
}

// settingIPV6EditResponseJSON contains the JSON metadata for the struct
// [SettingIPV6EditResponse]
type settingIPV6EditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6EditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingIPV6EditResponseID string

const (
	SettingIPV6EditResponseIDIPV6 SettingIPV6EditResponseID = "ipv6"
)

// Current value of the zone setting.
type SettingIPV6EditResponseValue string

const (
	SettingIPV6EditResponseValueOff SettingIPV6EditResponseValue = "off"
	SettingIPV6EditResponseValueOn  SettingIPV6EditResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingIPV6EditResponseEditable bool

const (
	SettingIPV6EditResponseEditableTrue  SettingIPV6EditResponseEditable = true
	SettingIPV6EditResponseEditableFalse SettingIPV6EditResponseEditable = false
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

type SettingIPV6EditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingIPV6EditParamsValue] `json:"value,required"`
}

func (r SettingIPV6EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingIPV6EditParamsValue string

const (
	SettingIPV6EditParamsValueOff SettingIPV6EditParamsValue = "off"
	SettingIPV6EditParamsValueOn  SettingIPV6EditParamsValue = "on"
)

type SettingIPV6EditResponseEnvelope struct {
	Errors   []SettingIPV6EditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingIPV6EditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IPv6 on all subdomains that are Cloudflare enabled.
	// (https://support.cloudflare.com/hc/en-us/articles/200168586).
	Result SettingIPV6EditResponse             `json:"result"`
	JSON   settingIPV6EditResponseEnvelopeJSON `json:"-"`
}

// settingIPV6EditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingIPV6EditResponseEnvelope]
type settingIPV6EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPV6EditResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    settingIPV6EditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingIPV6EditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingIPV6EditResponseEnvelopeErrors]
type settingIPV6EditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6EditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPV6EditResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingIPV6EditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingIPV6EditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingIPV6EditResponseEnvelopeMessages]
type settingIPV6EditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPV6EditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPV6GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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
