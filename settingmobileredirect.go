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

// SettingMobileRedirectService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingMobileRedirectService]
// method instead.
type SettingMobileRedirectService struct {
	Options []option.RequestOption
}

// NewSettingMobileRedirectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingMobileRedirectService(opts ...option.RequestOption) (r *SettingMobileRedirectService) {
	r = &SettingMobileRedirectService{}
	r.Options = opts
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
func (r *SettingMobileRedirectService) Edit(ctx context.Context, params SettingMobileRedirectEditParams, opts ...option.RequestOption) (res *SettingMobileRedirectEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMobileRedirectEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mobile_redirect", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
func (r *SettingMobileRedirectService) Get(ctx context.Context, query SettingMobileRedirectGetParams, opts ...option.RequestOption) (res *SettingMobileRedirectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMobileRedirectGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mobile_redirect", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type SettingMobileRedirectEditResponse struct {
	// Identifier of the zone setting.
	ID SettingMobileRedirectEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingMobileRedirectEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingMobileRedirectEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingMobileRedirectEditResponseJSON `json:"-"`
}

// settingMobileRedirectEditResponseJSON contains the JSON metadata for the struct
// [SettingMobileRedirectEditResponse]
type settingMobileRedirectEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of the zone setting.
type SettingMobileRedirectEditResponseID string

const (
	SettingMobileRedirectEditResponseIDMobileRedirect SettingMobileRedirectEditResponseID = "mobile_redirect"
)

// Current value of the zone setting.
type SettingMobileRedirectEditResponseValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status SettingMobileRedirectEditResponseValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                       `json:"strip_uri"`
	JSON     settingMobileRedirectEditResponseValueJSON `json:"-"`
}

// settingMobileRedirectEditResponseValueJSON contains the JSON metadata for the
// struct [SettingMobileRedirectEditResponseValue]
type settingMobileRedirectEditResponseValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingMobileRedirectEditResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type SettingMobileRedirectEditResponseValueStatus string

const (
	SettingMobileRedirectEditResponseValueStatusOn  SettingMobileRedirectEditResponseValueStatus = "on"
	SettingMobileRedirectEditResponseValueStatusOff SettingMobileRedirectEditResponseValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingMobileRedirectEditResponseEditable bool

const (
	SettingMobileRedirectEditResponseEditableTrue  SettingMobileRedirectEditResponseEditable = true
	SettingMobileRedirectEditResponseEditableFalse SettingMobileRedirectEditResponseEditable = false
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type SettingMobileRedirectGetResponse struct {
	// Identifier of the zone setting.
	ID SettingMobileRedirectGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingMobileRedirectGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingMobileRedirectGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingMobileRedirectGetResponseJSON `json:"-"`
}

// settingMobileRedirectGetResponseJSON contains the JSON metadata for the struct
// [SettingMobileRedirectGetResponse]
type settingMobileRedirectGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of the zone setting.
type SettingMobileRedirectGetResponseID string

const (
	SettingMobileRedirectGetResponseIDMobileRedirect SettingMobileRedirectGetResponseID = "mobile_redirect"
)

// Current value of the zone setting.
type SettingMobileRedirectGetResponseValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status SettingMobileRedirectGetResponseValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                      `json:"strip_uri"`
	JSON     settingMobileRedirectGetResponseValueJSON `json:"-"`
}

// settingMobileRedirectGetResponseValueJSON contains the JSON metadata for the
// struct [SettingMobileRedirectGetResponseValue]
type settingMobileRedirectGetResponseValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingMobileRedirectGetResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type SettingMobileRedirectGetResponseValueStatus string

const (
	SettingMobileRedirectGetResponseValueStatusOn  SettingMobileRedirectGetResponseValueStatus = "on"
	SettingMobileRedirectGetResponseValueStatusOff SettingMobileRedirectGetResponseValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingMobileRedirectGetResponseEditable bool

const (
	SettingMobileRedirectGetResponseEditableTrue  SettingMobileRedirectGetResponseEditable = true
	SettingMobileRedirectGetResponseEditableFalse SettingMobileRedirectGetResponseEditable = false
)

type SettingMobileRedirectEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingMobileRedirectEditParamsValue] `json:"value,required"`
}

func (r SettingMobileRedirectEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingMobileRedirectEditParamsValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain param.Field[string] `json:"mobile_subdomain"`
	// Whether or not mobile redirect is enabled.
	Status param.Field[SettingMobileRedirectEditParamsValueStatus] `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri param.Field[bool] `json:"strip_uri"`
}

func (r SettingMobileRedirectEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not mobile redirect is enabled.
type SettingMobileRedirectEditParamsValueStatus string

const (
	SettingMobileRedirectEditParamsValueStatusOn  SettingMobileRedirectEditParamsValueStatus = "on"
	SettingMobileRedirectEditParamsValueStatusOff SettingMobileRedirectEditParamsValueStatus = "off"
)

type SettingMobileRedirectEditResponseEnvelope struct {
	Errors   []SettingMobileRedirectEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMobileRedirectEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically redirect visitors on mobile devices to a mobile-optimized
	// subdomain. Refer to
	// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
	// for more information.
	Result SettingMobileRedirectEditResponse             `json:"result"`
	JSON   settingMobileRedirectEditResponseEnvelopeJSON `json:"-"`
}

// settingMobileRedirectEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingMobileRedirectEditResponseEnvelope]
type settingMobileRedirectEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMobileRedirectEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingMobileRedirectEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMobileRedirectEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingMobileRedirectEditResponseEnvelopeErrors]
type settingMobileRedirectEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMobileRedirectEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingMobileRedirectEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMobileRedirectEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingMobileRedirectEditResponseEnvelopeMessages]
type settingMobileRedirectEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMobileRedirectGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingMobileRedirectGetResponseEnvelope struct {
	Errors   []SettingMobileRedirectGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMobileRedirectGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically redirect visitors on mobile devices to a mobile-optimized
	// subdomain. Refer to
	// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
	// for more information.
	Result SettingMobileRedirectGetResponse             `json:"result"`
	JSON   settingMobileRedirectGetResponseEnvelopeJSON `json:"-"`
}

// settingMobileRedirectGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingMobileRedirectGetResponseEnvelope]
type settingMobileRedirectGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMobileRedirectGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingMobileRedirectGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMobileRedirectGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingMobileRedirectGetResponseEnvelopeErrors]
type settingMobileRedirectGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMobileRedirectGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingMobileRedirectGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMobileRedirectGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingMobileRedirectGetResponseEnvelopeMessages]
type settingMobileRedirectGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
