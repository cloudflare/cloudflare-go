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

// SettingSSLService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSettingSSLService] method instead.
type SettingSSLService struct {
	Options []option.RequestOption
}

// NewSettingSSLService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingSSLService(opts ...option.RequestOption) (r *SettingSSLService) {
	r = &SettingSSLService{}
	r.Options = opts
	return
}

// SSL encrypts your visitor's connection and safeguards credit card numbers and
// other personal data to and from your website. SSL can take up to 5 minutes to
// fully activate. Requires Cloudflare active on your root domain or www domain.
// Off: no SSL between the visitor and Cloudflare, and no SSL between Cloudflare
// and your web server (all HTTP traffic). Flexible: SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, but no SSL between Cloudflare and
// your web server. You don't need to have an SSL cert on your web server, but your
// vistors will still see the site as being HTTPS enabled. Full: SSL between the
// visitor and Cloudflare -- visitor sees HTTPS on your site, and SSL between
// Cloudflare and your web server. You'll need to have your own SSL cert or
// self-signed cert at the very least. Full (Strict): SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, and SSL between Cloudflare and
// your web server. You'll need to have a valid SSL certificate installed on your
// web server. This certificate must be signed by a certificate authority, have an
// expiration date in the future, and respond for the request domain name
// (hostname). (https://support.cloudflare.com/hc/en-us/articles/200170416).
func (r *SettingSSLService) Update(ctx context.Context, zoneID string, body SettingSSLUpdateParams, opts ...option.RequestOption) (res *SettingSSLUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSSLUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ssl", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// SSL encrypts your visitor's connection and safeguards credit card numbers and
// other personal data to and from your website. SSL can take up to 5 minutes to
// fully activate. Requires Cloudflare active on your root domain or www domain.
// Off: no SSL between the visitor and Cloudflare, and no SSL between Cloudflare
// and your web server (all HTTP traffic). Flexible: SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, but no SSL between Cloudflare and
// your web server. You don't need to have an SSL cert on your web server, but your
// vistors will still see the site as being HTTPS enabled. Full: SSL between the
// visitor and Cloudflare -- visitor sees HTTPS on your site, and SSL between
// Cloudflare and your web server. You'll need to have your own SSL cert or
// self-signed cert at the very least. Full (Strict): SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, and SSL between Cloudflare and
// your web server. You'll need to have a valid SSL certificate installed on your
// web server. This certificate must be signed by a certificate authority, have an
// expiration date in the future, and respond for the request domain name
// (hostname). (https://support.cloudflare.com/hc/en-us/articles/200170416).
func (r *SettingSSLService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingSSLGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSSLGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ssl", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// SSL encrypts your visitor's connection and safeguards credit card numbers and
// other personal data to and from your website. SSL can take up to 5 minutes to
// fully activate. Requires Cloudflare active on your root domain or www domain.
// Off: no SSL between the visitor and Cloudflare, and no SSL between Cloudflare
// and your web server (all HTTP traffic). Flexible: SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, but no SSL between Cloudflare and
// your web server. You don't need to have an SSL cert on your web server, but your
// vistors will still see the site as being HTTPS enabled. Full: SSL between the
// visitor and Cloudflare -- visitor sees HTTPS on your site, and SSL between
// Cloudflare and your web server. You'll need to have your own SSL cert or
// self-signed cert at the very least. Full (Strict): SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, and SSL between Cloudflare and
// your web server. You'll need to have a valid SSL certificate installed on your
// web server. This certificate must be signed by a certificate authority, have an
// expiration date in the future, and respond for the request domain name
// (hostname). (https://support.cloudflare.com/hc/en-us/articles/200170416).
type SettingSSLUpdateResponse struct {
	// ID of the zone setting.
	ID SettingSSLUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingSSLUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingSSLUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingSSLUpdateResponseJSON `json:"-"`
}

// settingSSLUpdateResponseJSON contains the JSON metadata for the struct
// [SettingSSLUpdateResponse]
type settingSSLUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingSSLUpdateResponseID string

const (
	SettingSSLUpdateResponseIDSSL SettingSSLUpdateResponseID = "ssl"
)

// Current value of the zone setting.
type SettingSSLUpdateResponseValue string

const (
	SettingSSLUpdateResponseValueOff      SettingSSLUpdateResponseValue = "off"
	SettingSSLUpdateResponseValueFlexible SettingSSLUpdateResponseValue = "flexible"
	SettingSSLUpdateResponseValueFull     SettingSSLUpdateResponseValue = "full"
	SettingSSLUpdateResponseValueStrict   SettingSSLUpdateResponseValue = "strict"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingSSLUpdateResponseEditable bool

const (
	SettingSSLUpdateResponseEditableTrue  SettingSSLUpdateResponseEditable = true
	SettingSSLUpdateResponseEditableFalse SettingSSLUpdateResponseEditable = false
)

// SSL encrypts your visitor's connection and safeguards credit card numbers and
// other personal data to and from your website. SSL can take up to 5 minutes to
// fully activate. Requires Cloudflare active on your root domain or www domain.
// Off: no SSL between the visitor and Cloudflare, and no SSL between Cloudflare
// and your web server (all HTTP traffic). Flexible: SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, but no SSL between Cloudflare and
// your web server. You don't need to have an SSL cert on your web server, but your
// vistors will still see the site as being HTTPS enabled. Full: SSL between the
// visitor and Cloudflare -- visitor sees HTTPS on your site, and SSL between
// Cloudflare and your web server. You'll need to have your own SSL cert or
// self-signed cert at the very least. Full (Strict): SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, and SSL between Cloudflare and
// your web server. You'll need to have a valid SSL certificate installed on your
// web server. This certificate must be signed by a certificate authority, have an
// expiration date in the future, and respond for the request domain name
// (hostname). (https://support.cloudflare.com/hc/en-us/articles/200170416).
type SettingSSLGetResponse struct {
	// ID of the zone setting.
	ID SettingSSLGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingSSLGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingSSLGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingSSLGetResponseJSON `json:"-"`
}

// settingSSLGetResponseJSON contains the JSON metadata for the struct
// [SettingSSLGetResponse]
type settingSSLGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingSSLGetResponseID string

const (
	SettingSSLGetResponseIDSSL SettingSSLGetResponseID = "ssl"
)

// Current value of the zone setting.
type SettingSSLGetResponseValue string

const (
	SettingSSLGetResponseValueOff      SettingSSLGetResponseValue = "off"
	SettingSSLGetResponseValueFlexible SettingSSLGetResponseValue = "flexible"
	SettingSSLGetResponseValueFull     SettingSSLGetResponseValue = "full"
	SettingSSLGetResponseValueStrict   SettingSSLGetResponseValue = "strict"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingSSLGetResponseEditable bool

const (
	SettingSSLGetResponseEditableTrue  SettingSSLGetResponseEditable = true
	SettingSSLGetResponseEditableFalse SettingSSLGetResponseEditable = false
)

type SettingSSLUpdateParams struct {
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value param.Field[SettingSSLUpdateParamsValue] `json:"value,required"`
}

func (r SettingSSLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Depends on the zone's plan level
type SettingSSLUpdateParamsValue string

const (
	SettingSSLUpdateParamsValueOff      SettingSSLUpdateParamsValue = "off"
	SettingSSLUpdateParamsValueFlexible SettingSSLUpdateParamsValue = "flexible"
	SettingSSLUpdateParamsValueFull     SettingSSLUpdateParamsValue = "full"
	SettingSSLUpdateParamsValueStrict   SettingSSLUpdateParamsValue = "strict"
)

type SettingSSLUpdateResponseEnvelope struct {
	Errors   []SettingSSLUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSSLUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// SSL encrypts your visitor's connection and safeguards credit card numbers and
	// other personal data to and from your website. SSL can take up to 5 minutes to
	// fully activate. Requires Cloudflare active on your root domain or www domain.
	// Off: no SSL between the visitor and Cloudflare, and no SSL between Cloudflare
	// and your web server (all HTTP traffic). Flexible: SSL between the visitor and
	// Cloudflare -- visitor sees HTTPS on your site, but no SSL between Cloudflare and
	// your web server. You don't need to have an SSL cert on your web server, but your
	// vistors will still see the site as being HTTPS enabled. Full: SSL between the
	// visitor and Cloudflare -- visitor sees HTTPS on your site, and SSL between
	// Cloudflare and your web server. You'll need to have your own SSL cert or
	// self-signed cert at the very least. Full (Strict): SSL between the visitor and
	// Cloudflare -- visitor sees HTTPS on your site, and SSL between Cloudflare and
	// your web server. You'll need to have a valid SSL certificate installed on your
	// web server. This certificate must be signed by a certificate authority, have an
	// expiration date in the future, and respond for the request domain name
	// (hostname). (https://support.cloudflare.com/hc/en-us/articles/200170416).
	Result SettingSSLUpdateResponse             `json:"result"`
	JSON   settingSSLUpdateResponseEnvelopeJSON `json:"-"`
}

// settingSSLUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingSSLUpdateResponseEnvelope]
type settingSSLUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSSLUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingSSLUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSSLUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingSSLUpdateResponseEnvelopeErrors]
type settingSSLUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSSLUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingSSLUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSSLUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingSSLUpdateResponseEnvelopeMessages]
type settingSSLUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSSLGetResponseEnvelope struct {
	Errors   []SettingSSLGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSSLGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// SSL encrypts your visitor's connection and safeguards credit card numbers and
	// other personal data to and from your website. SSL can take up to 5 minutes to
	// fully activate. Requires Cloudflare active on your root domain or www domain.
	// Off: no SSL between the visitor and Cloudflare, and no SSL between Cloudflare
	// and your web server (all HTTP traffic). Flexible: SSL between the visitor and
	// Cloudflare -- visitor sees HTTPS on your site, but no SSL between Cloudflare and
	// your web server. You don't need to have an SSL cert on your web server, but your
	// vistors will still see the site as being HTTPS enabled. Full: SSL between the
	// visitor and Cloudflare -- visitor sees HTTPS on your site, and SSL between
	// Cloudflare and your web server. You'll need to have your own SSL cert or
	// self-signed cert at the very least. Full (Strict): SSL between the visitor and
	// Cloudflare -- visitor sees HTTPS on your site, and SSL between Cloudflare and
	// your web server. You'll need to have a valid SSL certificate installed on your
	// web server. This certificate must be signed by a certificate authority, have an
	// expiration date in the future, and respond for the request domain name
	// (hostname). (https://support.cloudflare.com/hc/en-us/articles/200170416).
	Result SettingSSLGetResponse             `json:"result"`
	JSON   settingSSLGetResponseEnvelopeJSON `json:"-"`
}

// settingSSLGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingSSLGetResponseEnvelope]
type settingSSLGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSSLGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    settingSSLGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSSLGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingSSLGetResponseEnvelopeErrors]
type settingSSLGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSSLGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    settingSSLGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSSLGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingSSLGetResponseEnvelopeMessages]
type settingSSLGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
