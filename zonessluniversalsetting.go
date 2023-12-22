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
func (r *ZoneSslUniversalSettingService) UniversalSslSettingsForAZoneEditUniversalSslSettings(ctx context.Context, zoneIdentifier string, body ZoneSslUniversalSettingUniversalSslSettingsForAZoneEditUniversalSslSettingsParams, opts ...option.RequestOption) (res *SslUniversalSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/universal/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get Universal SSL Settings for a Zone.
func (r *ZoneSslUniversalSettingService) UniversalSslSettingsForAZoneUniversalSslSettingsDetails(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *SslUniversalSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/universal/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SslUniversalSettingsResponse struct {
	Errors   []SslUniversalSettingsResponseError   `json:"errors"`
	Messages []SslUniversalSettingsResponseMessage `json:"messages"`
	Result   SslUniversalSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success SslUniversalSettingsResponseSuccess `json:"success"`
	JSON    sslUniversalSettingsResponseJSON    `json:"-"`
}

// sslUniversalSettingsResponseJSON contains the JSON metadata for the struct
// [SslUniversalSettingsResponse]
type sslUniversalSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SslUniversalSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SslUniversalSettingsResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    sslUniversalSettingsResponseErrorJSON `json:"-"`
}

// sslUniversalSettingsResponseErrorJSON contains the JSON metadata for the struct
// [SslUniversalSettingsResponseError]
type sslUniversalSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SslUniversalSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SslUniversalSettingsResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    sslUniversalSettingsResponseMessageJSON `json:"-"`
}

// sslUniversalSettingsResponseMessageJSON contains the JSON metadata for the
// struct [SslUniversalSettingsResponseMessage]
type sslUniversalSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SslUniversalSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SslUniversalSettingsResponseResult struct {
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
	Enabled bool                                   `json:"enabled"`
	JSON    sslUniversalSettingsResponseResultJSON `json:"-"`
}

// sslUniversalSettingsResponseResultJSON contains the JSON metadata for the struct
// [SslUniversalSettingsResponseResult]
type sslUniversalSettingsResponseResultJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SslUniversalSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SslUniversalSettingsResponseSuccess bool

const (
	SslUniversalSettingsResponseSuccessTrue SslUniversalSettingsResponseSuccess = true
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
