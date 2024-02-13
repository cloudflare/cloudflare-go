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
func (r *SettingMobileRedirectService) Update(ctx context.Context, zoneID string, body SettingMobileRedirectUpdateParams, opts ...option.RequestOption) (res *SettingMobileRedirectUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMobileRedirectUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mobile_redirect", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
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
func (r *SettingMobileRedirectService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingMobileRedirectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMobileRedirectGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mobile_redirect", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
type SettingMobileRedirectUpdateResponse struct {
	// Identifier of the zone setting.
	ID SettingMobileRedirectUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingMobileRedirectUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingMobileRedirectUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingMobileRedirectUpdateResponseJSON `json:"-"`
}

// settingMobileRedirectUpdateResponseJSON contains the JSON metadata for the
// struct [SettingMobileRedirectUpdateResponse]
type settingMobileRedirectUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of the zone setting.
type SettingMobileRedirectUpdateResponseID string

const (
	SettingMobileRedirectUpdateResponseIDMobileRedirect SettingMobileRedirectUpdateResponseID = "mobile_redirect"
)

// Current value of the zone setting.
type SettingMobileRedirectUpdateResponseValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status SettingMobileRedirectUpdateResponseValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                         `json:"strip_uri"`
	JSON     settingMobileRedirectUpdateResponseValueJSON `json:"-"`
}

// settingMobileRedirectUpdateResponseValueJSON contains the JSON metadata for the
// struct [SettingMobileRedirectUpdateResponseValue]
type settingMobileRedirectUpdateResponseValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingMobileRedirectUpdateResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type SettingMobileRedirectUpdateResponseValueStatus string

const (
	SettingMobileRedirectUpdateResponseValueStatusOn  SettingMobileRedirectUpdateResponseValueStatus = "on"
	SettingMobileRedirectUpdateResponseValueStatusOff SettingMobileRedirectUpdateResponseValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingMobileRedirectUpdateResponseEditable bool

const (
	SettingMobileRedirectUpdateResponseEditableTrue  SettingMobileRedirectUpdateResponseEditable = true
	SettingMobileRedirectUpdateResponseEditableFalse SettingMobileRedirectUpdateResponseEditable = false
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

type SettingMobileRedirectUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingMobileRedirectUpdateParamsValue] `json:"value,required"`
}

func (r SettingMobileRedirectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingMobileRedirectUpdateParamsValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain param.Field[string] `json:"mobile_subdomain"`
	// Whether or not mobile redirect is enabled.
	Status param.Field[SettingMobileRedirectUpdateParamsValueStatus] `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri param.Field[bool] `json:"strip_uri"`
}

func (r SettingMobileRedirectUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not mobile redirect is enabled.
type SettingMobileRedirectUpdateParamsValueStatus string

const (
	SettingMobileRedirectUpdateParamsValueStatusOn  SettingMobileRedirectUpdateParamsValueStatus = "on"
	SettingMobileRedirectUpdateParamsValueStatusOff SettingMobileRedirectUpdateParamsValueStatus = "off"
)

type SettingMobileRedirectUpdateResponseEnvelope struct {
	Errors   []SettingMobileRedirectUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMobileRedirectUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically redirect visitors on mobile devices to a mobile-optimized
	// subdomain. Refer to
	// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
	// for more information.
	Result SettingMobileRedirectUpdateResponse             `json:"result"`
	JSON   settingMobileRedirectUpdateResponseEnvelopeJSON `json:"-"`
}

// settingMobileRedirectUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingMobileRedirectUpdateResponseEnvelope]
type settingMobileRedirectUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMobileRedirectUpdateResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingMobileRedirectUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMobileRedirectUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingMobileRedirectUpdateResponseEnvelopeErrors]
type settingMobileRedirectUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMobileRedirectUpdateResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingMobileRedirectUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMobileRedirectUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingMobileRedirectUpdateResponseEnvelopeMessages]
type settingMobileRedirectUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
