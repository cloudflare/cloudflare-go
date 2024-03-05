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

// ZoneSettingSSLService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingSSLService] method
// instead.
type ZoneSettingSSLService struct {
	Options []option.RequestOption
}

// NewZoneSettingSSLService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingSSLService(opts ...option.RequestOption) (r *ZoneSettingSSLService) {
	r = &ZoneSettingSSLService{}
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
func (r *ZoneSettingSSLService) Edit(ctx context.Context, params ZoneSettingSSLEditParams, opts ...option.RequestOption) (res *ZonesSSL, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingSSLEditResponseEnvelope
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
func (r *ZoneSettingSSLService) Get(ctx context.Context, query ZoneSettingSSLGetParams, opts ...option.RequestOption) (res *ZonesSSL, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingSSLGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ssl", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
type ZonesSSL struct {
	// ID of the zone setting.
	ID ZonesSSLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesSSLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesSSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time    `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesSSLJSON `json:"-"`
}

// zonesSSLJSON contains the JSON metadata for the struct [ZonesSSL]
type zonesSSLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesSSL) implementsZoneSettingEditResponse() {}

func (r ZonesSSL) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesSSLID string

const (
	ZonesSSLIDSSL ZonesSSLID = "ssl"
)

// Current value of the zone setting.
type ZonesSSLValue string

const (
	ZonesSSLValueOff      ZonesSSLValue = "off"
	ZonesSSLValueFlexible ZonesSSLValue = "flexible"
	ZonesSSLValueFull     ZonesSSLValue = "full"
	ZonesSSLValueStrict   ZonesSSLValue = "strict"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesSSLEditable bool

const (
	ZonesSSLEditableTrue  ZonesSSLEditable = true
	ZonesSSLEditableFalse ZonesSSLEditable = false
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
type ZonesSSLParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSSLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesSSLValue] `json:"value,required"`
}

func (r ZonesSSLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSSLParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingSSLEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value param.Field[ZoneSettingSSLEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingSSLEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingSSLEditParamsValue string

const (
	ZoneSettingSSLEditParamsValueOff      ZoneSettingSSLEditParamsValue = "off"
	ZoneSettingSSLEditParamsValueFlexible ZoneSettingSSLEditParamsValue = "flexible"
	ZoneSettingSSLEditParamsValueFull     ZoneSettingSSLEditParamsValue = "full"
	ZoneSettingSSLEditParamsValueStrict   ZoneSettingSSLEditParamsValue = "strict"
)

type ZoneSettingSSLEditResponseEnvelope struct {
	Errors   []ZoneSettingSSLEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingSSLEditResponseEnvelopeMessages `json:"messages,required"`
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
	Result ZonesSSL                               `json:"result"`
	JSON   zoneSettingSSLEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingSSLEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneSettingSSLEditResponseEnvelope]
type zoneSettingSSLEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSSLEditResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSettingSSLEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingSSLEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingSSLEditResponseEnvelopeErrors]
type zoneSettingSSLEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSSLEditResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingSSLEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingSSLEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingSSLEditResponseEnvelopeMessages]
type zoneSettingSSLEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSSLGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingSSLGetResponseEnvelope struct {
	Errors   []ZoneSettingSSLGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingSSLGetResponseEnvelopeMessages `json:"messages,required"`
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
	Result ZonesSSL                              `json:"result"`
	JSON   zoneSettingSSLGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingSSLGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneSettingSSLGetResponseEnvelope]
type zoneSettingSSLGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSSLGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSettingSSLGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingSSLGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingSSLGetResponseEnvelopeErrors]
type zoneSettingSSLGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSSLGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingSSLGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingSSLGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZoneSettingSSLGetResponseEnvelopeMessages]
type zoneSettingSSLGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
