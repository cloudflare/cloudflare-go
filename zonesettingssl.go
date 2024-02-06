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
func (r *ZoneSettingSSLService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingSSLUpdateParams, opts ...option.RequestOption) (res *ZoneSettingSSLUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ssl", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
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
func (r *ZoneSettingSSLService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingSSLListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ssl", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingSSLUpdateResponse struct {
	Errors   []ZoneSettingSSLUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingSSLUpdateResponseMessage `json:"messages"`
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
	Result ZoneSettingSSLUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                             `json:"success"`
	JSON    zoneSettingSSLUpdateResponseJSON `json:"-"`
}

// zoneSettingSSLUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingSSLUpdateResponse]
type zoneSettingSSLUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSSLUpdateResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingSSLUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingSSLUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingSSLUpdateResponseError]
type zoneSettingSSLUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSSLUpdateResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingSSLUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingSSLUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingSSLUpdateResponseMessage]
type zoneSettingSSLUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
type ZoneSettingSSLUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingSSLUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSSLUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value ZoneSettingSSLUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingSSLUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingSSLUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingSSLUpdateResponseResult]
type zoneSettingSSLUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingSSLUpdateResponseResultID string

const (
	ZoneSettingSSLUpdateResponseResultIDSSL ZoneSettingSSLUpdateResponseResultID = "ssl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSSLUpdateResponseResultEditable bool

const (
	ZoneSettingSSLUpdateResponseResultEditableTrue  ZoneSettingSSLUpdateResponseResultEditable = true
	ZoneSettingSSLUpdateResponseResultEditableFalse ZoneSettingSSLUpdateResponseResultEditable = false
)

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingSSLUpdateResponseResultValue string

const (
	ZoneSettingSSLUpdateResponseResultValueOff      ZoneSettingSSLUpdateResponseResultValue = "off"
	ZoneSettingSSLUpdateResponseResultValueFlexible ZoneSettingSSLUpdateResponseResultValue = "flexible"
	ZoneSettingSSLUpdateResponseResultValueFull     ZoneSettingSSLUpdateResponseResultValue = "full"
	ZoneSettingSSLUpdateResponseResultValueStrict   ZoneSettingSSLUpdateResponseResultValue = "strict"
)

type ZoneSettingSSLListResponse struct {
	Errors   []ZoneSettingSSLListResponseError   `json:"errors"`
	Messages []ZoneSettingSSLListResponseMessage `json:"messages"`
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
	Result ZoneSettingSSLListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                           `json:"success"`
	JSON    zoneSettingSSLListResponseJSON `json:"-"`
}

// zoneSettingSSLListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingSSLListResponse]
type zoneSettingSSLListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSSLListResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneSettingSSLListResponseErrorJSON `json:"-"`
}

// zoneSettingSSLListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingSSLListResponseError]
type zoneSettingSSLListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSSLListResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingSSLListResponseMessageJSON `json:"-"`
}

// zoneSettingSSLListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSettingSSLListResponseMessage]
type zoneSettingSSLListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
type ZoneSettingSSLListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingSSLListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSSLListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value ZoneSettingSSLListResponseResultValue `json:"value"`
	JSON  zoneSettingSSLListResponseResultJSON  `json:"-"`
}

// zoneSettingSSLListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingSSLListResponseResult]
type zoneSettingSSLListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSSLListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingSSLListResponseResultID string

const (
	ZoneSettingSSLListResponseResultIDSSL ZoneSettingSSLListResponseResultID = "ssl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSSLListResponseResultEditable bool

const (
	ZoneSettingSSLListResponseResultEditableTrue  ZoneSettingSSLListResponseResultEditable = true
	ZoneSettingSSLListResponseResultEditableFalse ZoneSettingSSLListResponseResultEditable = false
)

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingSSLListResponseResultValue string

const (
	ZoneSettingSSLListResponseResultValueOff      ZoneSettingSSLListResponseResultValue = "off"
	ZoneSettingSSLListResponseResultValueFlexible ZoneSettingSSLListResponseResultValue = "flexible"
	ZoneSettingSSLListResponseResultValueFull     ZoneSettingSSLListResponseResultValue = "full"
	ZoneSettingSSLListResponseResultValueStrict   ZoneSettingSSLListResponseResultValue = "strict"
)

type ZoneSettingSSLUpdateParams struct {
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value param.Field[ZoneSettingSSLUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingSSLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingSSLUpdateParamsValue string

const (
	ZoneSettingSSLUpdateParamsValueOff      ZoneSettingSSLUpdateParamsValue = "off"
	ZoneSettingSSLUpdateParamsValueFlexible ZoneSettingSSLUpdateParamsValue = "flexible"
	ZoneSettingSSLUpdateParamsValueFull     ZoneSettingSSLUpdateParamsValue = "full"
	ZoneSettingSSLUpdateParamsValueStrict   ZoneSettingSSLUpdateParamsValue = "strict"
)
