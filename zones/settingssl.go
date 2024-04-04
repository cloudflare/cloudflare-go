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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *SettingSSLService) Edit(ctx context.Context, params SettingSSLEditParams, opts ...option.RequestOption) (res *ZoneSettingSSL, err error) {
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
func (r *SettingSSLService) Get(ctx context.Context, query SettingSSLGetParams, opts ...option.RequestOption) (res *ZoneSettingSSL, err error) {
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
type ZoneSettingSSL struct {
	// ID of the zone setting.
	ID ZoneSettingSSLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingSSLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time          `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSSLJSON `json:"-"`
}

// zoneSettingSSLJSON contains the JSON metadata for the struct [ZoneSettingSSL]
type zoneSettingSSLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSSLJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingSSL) implementsZonesSettingEditResponse() {}

func (r ZoneSettingSSL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingSSLID string

const (
	ZoneSettingSSLIDSSL ZoneSettingSSLID = "ssl"
)

func (r ZoneSettingSSLID) IsKnown() bool {
	switch r {
	case ZoneSettingSSLIDSSL:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingSSLValue string

const (
	ZoneSettingSSLValueOff      ZoneSettingSSLValue = "off"
	ZoneSettingSSLValueFlexible ZoneSettingSSLValue = "flexible"
	ZoneSettingSSLValueFull     ZoneSettingSSLValue = "full"
	ZoneSettingSSLValueStrict   ZoneSettingSSLValue = "strict"
)

func (r ZoneSettingSSLValue) IsKnown() bool {
	switch r {
	case ZoneSettingSSLValueOff, ZoneSettingSSLValueFlexible, ZoneSettingSSLValueFull, ZoneSettingSSLValueStrict:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSSLEditable bool

const (
	ZoneSettingSSLEditableTrue  ZoneSettingSSLEditable = true
	ZoneSettingSSLEditableFalse ZoneSettingSSLEditable = false
)

func (r ZoneSettingSSLEditable) IsKnown() bool {
	switch r {
	case ZoneSettingSSLEditableTrue, ZoneSettingSSLEditableFalse:
		return true
	}
	return false
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
type ZoneSettingSSLParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingSSLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingSSLValue] `json:"value,required"`
}

func (r ZoneSettingSSLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingSSLParam) implementsZonesSettingEditParamsItem() {}

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
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
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
	Result ZoneSettingSSL                     `json:"result"`
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
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
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
	Result ZoneSettingSSL                    `json:"result"`
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
