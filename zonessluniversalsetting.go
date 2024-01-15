// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSslUniversalSettingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSslUniversalSettingService] method instead.
type ZoneSslUniversalSettingService struct {
	Options []option.RequestOption
}

// NewZoneSslUniversalSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSslUniversalSettingService(opts ...option.RequestOption) (r *ZoneSslUniversalSettingService) {
	r = &ZoneSslUniversalSettingService{}
	r.Options = opts
	return
}

// Patch Universal SSL Settings for a Zone.
func (r *ZoneSslUniversalSettingService) UniversalSslSettingsForAZoneEditUniversalSslSettings(ctx context.Context, zoneIdentifier string, body ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsParams, opts ...option.RequestOption) (res *ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/universal/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get Universal SSL Settings for a Zone.
func (r *ZoneSslUniversalSettingService) UniversalSslSettingsForAZoneUniversalSslSettingsDetails(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/universal/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponse struct {
	Errors   []ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseError   `json:"errors"`
	Messages []ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseMessage `json:"messages"`
	Result   ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseSuccess `json:"success"`
	JSON    zoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseJSON    `json:"-"`
}

// zoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseJSON
// contains the JSON metadata for the struct
// [ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponse]
type zoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseError struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    zoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseErrorJSON `json:"-"`
}

// zoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseError]
type zoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseMessage struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    zoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseMessageJSON `json:"-"`
}

// zoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseMessage]
type zoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseResult struct {
	// Disabling Universal SSL removes any currently active Universal SSL certificates
	// for your zone from the edge and prevents any future Universal SSL certificates
	// from being ordered. If there are no advanced certificates or custom certificates
	// uploaded for the domain, visitors will be unable to access the domain over
	// HTTPS.
	//
	// By disabling Universal SSL, you understand that the following Cloudflare
	// settings and preferences will result in visitors being unable to visit your
	// domain unless you have uploaded a custom certificate or purchased an advanced
	// certificate.
	//
	// - HSTS
	// - Always Use HTTPS
	// - Opportunistic Encryption
	// - Onion Routing
	// - Any Page Rules redirecting traffic to HTTPS
	//
	// Similarly, any HTTP redirect to HTTPS at the origin while the Cloudflare proxy
	// is enabled will result in users being unable to visit your site without a valid
	// certificate at Cloudflare's edge.
	//
	// If you do not have a valid custom or advanced certificate at Cloudflare's edge
	// and are unsure if any of the above Cloudflare settings are enabled, or if any
	// HTTP redirects exist at your origin, we advise leaving Universal SSL enabled for
	// your domain.
	Enabled bool                                                                                          `json:"enabled"`
	JSON    zoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseResultJSON `json:"-"`
}

// zoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseResult]
type zoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseResultJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseSuccess bool

const (
	ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseSuccessTrue ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsResponseSuccess = true
)

type ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponse struct {
	Errors   []ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseError   `json:"errors"`
	Messages []ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseMessage `json:"messages"`
	Result   ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseSuccess `json:"success"`
	JSON    zoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseJSON    `json:"-"`
}

// zoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseJSON
// contains the JSON metadata for the struct
// [ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponse]
type zoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseError struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    zoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseErrorJSON `json:"-"`
}

// zoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseError]
type zoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseMessage struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    zoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseMessageJSON `json:"-"`
}

// zoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseMessage]
type zoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseResult struct {
	// Disabling Universal SSL removes any currently active Universal SSL certificates
	// for your zone from the edge and prevents any future Universal SSL certificates
	// from being ordered. If there are no advanced certificates or custom certificates
	// uploaded for the domain, visitors will be unable to access the domain over
	// HTTPS.
	//
	// By disabling Universal SSL, you understand that the following Cloudflare
	// settings and preferences will result in visitors being unable to visit your
	// domain unless you have uploaded a custom certificate or purchased an advanced
	// certificate.
	//
	// - HSTS
	// - Always Use HTTPS
	// - Opportunistic Encryption
	// - Onion Routing
	// - Any Page Rules redirecting traffic to HTTPS
	//
	// Similarly, any HTTP redirect to HTTPS at the origin while the Cloudflare proxy
	// is enabled will result in users being unable to visit your site without a valid
	// certificate at Cloudflare's edge.
	//
	// If you do not have a valid custom or advanced certificate at Cloudflare's edge
	// and are unsure if any of the above Cloudflare settings are enabled, or if any
	// HTTP redirects exist at your origin, we advise leaving Universal SSL enabled for
	// your domain.
	Enabled bool                                                                                             `json:"enabled"`
	JSON    zoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseResultJSON `json:"-"`
}

// zoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseResult]
type zoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseResultJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseSuccess bool

const (
	ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseSuccessTrue ZoneSslUniversalSettingUniversalSslSettingsForAZoneUniversalSslSettingsDetailsResponseSuccess = true
)

type ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsParams struct {
	// Disabling Universal SSL removes any currently active Universal SSL certificates
	// for your zone from the edge and prevents any future Universal SSL certificates
	// from being ordered. If there are no advanced certificates or custom certificates
	// uploaded for the domain, visitors will be unable to access the domain over
	// HTTPS.
	//
	// By disabling Universal SSL, you understand that the following Cloudflare
	// settings and preferences will result in visitors being unable to visit your
	// domain unless you have uploaded a custom certificate or purchased an advanced
	// certificate.
	//
	// - HSTS
	// - Always Use HTTPS
	// - Opportunistic Encryption
	// - Onion Routing
	// - Any Page Rules redirecting traffic to HTTPS
	//
	// Similarly, any HTTP redirect to HTTPS at the origin while the Cloudflare proxy
	// is enabled will result in users being unable to visit your site without a valid
	// certificate at Cloudflare's edge.
	//
	// If you do not have a valid custom or advanced certificate at Cloudflare's edge
	// and are unsure if any of the above Cloudflare settings are enabled, or if any
	// HTTP redirects exist at your origin, we advise leaving Universal SSL enabled for
	// your domain.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
