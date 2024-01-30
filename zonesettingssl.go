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

// ZoneSettingSslService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingSslService] method
// instead.
type ZoneSettingSslService struct {
	Options []option.RequestOption
}

// NewZoneSettingSslService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingSslService(opts ...option.RequestOption) (r *ZoneSettingSslService) {
	r = &ZoneSettingSslService{}
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
func (r *ZoneSettingSslService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingSslUpdateParams, opts ...option.RequestOption) (res *ZoneSettingSslUpdateResponse, err error) {
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
func (r *ZoneSettingSslService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingSslListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ssl", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingSslUpdateResponse struct {
	Errors   []ZoneSettingSslUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingSslUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                               `json:"success,required"`
	Result  ZoneSettingSslUpdateResponseResult `json:"result"`
	JSON    zoneSettingSslUpdateResponseJSON   `json:"-"`
}

// zoneSettingSslUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingSslUpdateResponse]
type zoneSettingSslUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSslUpdateResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingSslUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingSslUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingSslUpdateResponseError]
type zoneSettingSslUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSslUpdateResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingSslUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingSslUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingSslUpdateResponseMessage]
type zoneSettingSslUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSslUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingSslUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value ZoneSettingSslUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSslUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSslUpdateResponseResultJSON `json:"-"`
}

// zoneSettingSslUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingSslUpdateResponseResult]
type zoneSettingSslUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingSslUpdateResponseResultID string

const (
	ZoneSettingSslUpdateResponseResultIDSsl ZoneSettingSslUpdateResponseResultID = "ssl"
)

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingSslUpdateResponseResultValue string

const (
	ZoneSettingSslUpdateResponseResultValueOff      ZoneSettingSslUpdateResponseResultValue = "off"
	ZoneSettingSslUpdateResponseResultValueFlexible ZoneSettingSslUpdateResponseResultValue = "flexible"
	ZoneSettingSslUpdateResponseResultValueFull     ZoneSettingSslUpdateResponseResultValue = "full"
	ZoneSettingSslUpdateResponseResultValueStrict   ZoneSettingSslUpdateResponseResultValue = "strict"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSslUpdateResponseResultEditable bool

const (
	ZoneSettingSslUpdateResponseResultEditableTrue  ZoneSettingSslUpdateResponseResultEditable = true
	ZoneSettingSslUpdateResponseResultEditableFalse ZoneSettingSslUpdateResponseResultEditable = false
)

type ZoneSettingSslListResponse struct {
	Errors   []ZoneSettingSslListResponseError   `json:"errors,required"`
	Messages []ZoneSettingSslListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                             `json:"success,required"`
	Result  ZoneSettingSslListResponseResult `json:"result"`
	JSON    zoneSettingSslListResponseJSON   `json:"-"`
}

// zoneSettingSslListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingSslListResponse]
type zoneSettingSslListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSslListResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneSettingSslListResponseErrorJSON `json:"-"`
}

// zoneSettingSslListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingSslListResponseError]
type zoneSettingSslListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSslListResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingSslListResponseMessageJSON `json:"-"`
}

// zoneSettingSslListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSettingSslListResponseMessage]
type zoneSettingSslListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSslListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingSslListResponseResultID `json:"id,required"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value ZoneSettingSslListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSslListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSslListResponseResultJSON `json:"-"`
}

// zoneSettingSslListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingSslListResponseResult]
type zoneSettingSslListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingSslListResponseResultID string

const (
	ZoneSettingSslListResponseResultIDSsl ZoneSettingSslListResponseResultID = "ssl"
)

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingSslListResponseResultValue string

const (
	ZoneSettingSslListResponseResultValueOff      ZoneSettingSslListResponseResultValue = "off"
	ZoneSettingSslListResponseResultValueFlexible ZoneSettingSslListResponseResultValue = "flexible"
	ZoneSettingSslListResponseResultValueFull     ZoneSettingSslListResponseResultValue = "full"
	ZoneSettingSslListResponseResultValueStrict   ZoneSettingSslListResponseResultValue = "strict"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSslListResponseResultEditable bool

const (
	ZoneSettingSslListResponseResultEditableTrue  ZoneSettingSslListResponseResultEditable = true
	ZoneSettingSslListResponseResultEditableFalse ZoneSettingSslListResponseResultEditable = false
)

type ZoneSettingSslUpdateParams struct {
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value param.Field[ZoneSettingSslUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingSslUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingSslUpdateParamsValue string

const (
	ZoneSettingSslUpdateParamsValueOff      ZoneSettingSslUpdateParamsValue = "off"
	ZoneSettingSslUpdateParamsValueFlexible ZoneSettingSslUpdateParamsValue = "flexible"
	ZoneSettingSslUpdateParamsValueFull     ZoneSettingSslUpdateParamsValue = "full"
	ZoneSettingSslUpdateParamsValueStrict   ZoneSettingSslUpdateParamsValue = "strict"
)
