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

// SettingSSLService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingSSLService] method instead.
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
func (r *SettingSSLService) Edit(ctx context.Context, params SettingSSLEditParams, opts ...option.RequestOption) (res *SSL, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSSLEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ssl", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
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
func (r *SettingSSLService) Get(ctx context.Context, query SettingSSLGetParams, opts ...option.RequestOption) (res *SSL, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSSLGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ssl", query.ZoneID)
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
type SSL struct {
	// ID of the zone setting.
	ID SSLID `json:"id,required"`
	// Current value of the zone setting.
	Value SSLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       sslJSON   `json:"-"`
}

// sslJSON contains the JSON metadata for the struct [SSL]
type sslJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sslJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SSLID string

const (
	SSLIDSSL SSLID = "ssl"
)

func (r SSLID) IsKnown() bool {
	switch r {
	case SSLIDSSL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SSLValue string

const (
	SSLValueOff      SSLValue = "off"
	SSLValueFlexible SSLValue = "flexible"
	SSLValueFull     SSLValue = "full"
	SSLValueStrict   SSLValue = "strict"
)

func (r SSLValue) IsKnown() bool {
	switch r {
	case SSLValueOff, SSLValueFlexible, SSLValueFull, SSLValueStrict:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SSLEditable bool

const (
	SSLEditableTrue  SSLEditable = true
	SSLEditableFalse SSLEditable = false
)

func (r SSLEditable) IsKnown() bool {
	switch r {
	case SSLEditableTrue, SSLEditableFalse:
		return true
	}
	return false
}

type SettingSSLEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value param.Field[SettingSSLEditParamsValue] `json:"value,required"`
}

func (r SettingSSLEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Depends on the zone's plan level
type SettingSSLEditParamsValue string

const (
	SettingSSLEditParamsValueOff      SettingSSLEditParamsValue = "off"
	SettingSSLEditParamsValueFlexible SettingSSLEditParamsValue = "flexible"
	SettingSSLEditParamsValueFull     SettingSSLEditParamsValue = "full"
	SettingSSLEditParamsValueStrict   SettingSSLEditParamsValue = "strict"
)

func (r SettingSSLEditParamsValue) IsKnown() bool {
	switch r {
	case SettingSSLEditParamsValueOff, SettingSSLEditParamsValueFlexible, SettingSSLEditParamsValueFull, SettingSSLEditParamsValueStrict:
		return true
	}
	return false
}

type SettingSSLEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
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
	Result SSL                                `json:"result"`
	JSON   settingSSLEditResponseEnvelopeJSON `json:"-"`
}

// settingSSLEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingSSLEditResponseEnvelope]
type settingSSLEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSSLEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSSLEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingSSLGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingSSLGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
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
	Result SSL                               `json:"result"`
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

func (r settingSSLGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
