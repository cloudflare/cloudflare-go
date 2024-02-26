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

// SSLUniversalSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSSLUniversalSettingService]
// method instead.
type SSLUniversalSettingService struct {
	Options []option.RequestOption
}

// NewSSLUniversalSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSSLUniversalSettingService(opts ...option.RequestOption) (r *SSLUniversalSettingService) {
	r = &SSLUniversalSettingService{}
	r.Options = opts
	return
}

// Patch Universal SSL Settings for a Zone.
func (r *SSLUniversalSettingService) Edit(ctx context.Context, params SSLUniversalSettingEditParams, opts ...option.RequestOption) (res *SSLUniversalSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLUniversalSettingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/universal/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Universal SSL Settings for a Zone.
func (r *SSLUniversalSettingService) Get(ctx context.Context, query SSLUniversalSettingGetParams, opts ...option.RequestOption) (res *SSLUniversalSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLUniversalSettingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/universal/settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SSLUniversalSettingEditResponse struct {
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
	Enabled bool                                `json:"enabled"`
	JSON    sslUniversalSettingEditResponseJSON `json:"-"`
}

// sslUniversalSettingEditResponseJSON contains the JSON metadata for the struct
// [SSLUniversalSettingEditResponse]
type sslUniversalSettingEditResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLUniversalSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLUniversalSettingGetResponse struct {
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
	Enabled bool                               `json:"enabled"`
	JSON    sslUniversalSettingGetResponseJSON `json:"-"`
}

// sslUniversalSettingGetResponseJSON contains the JSON metadata for the struct
// [SSLUniversalSettingGetResponse]
type sslUniversalSettingGetResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLUniversalSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLUniversalSettingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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

func (r SSLUniversalSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SSLUniversalSettingEditResponseEnvelope struct {
	Errors   []SSLUniversalSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLUniversalSettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLUniversalSettingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLUniversalSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslUniversalSettingEditResponseEnvelopeJSON    `json:"-"`
}

// sslUniversalSettingEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SSLUniversalSettingEditResponseEnvelope]
type sslUniversalSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLUniversalSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLUniversalSettingEditResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    sslUniversalSettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// sslUniversalSettingEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SSLUniversalSettingEditResponseEnvelopeErrors]
type sslUniversalSettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLUniversalSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLUniversalSettingEditResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    sslUniversalSettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// sslUniversalSettingEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SSLUniversalSettingEditResponseEnvelopeMessages]
type sslUniversalSettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLUniversalSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLUniversalSettingEditResponseEnvelopeSuccess bool

const (
	SSLUniversalSettingEditResponseEnvelopeSuccessTrue SSLUniversalSettingEditResponseEnvelopeSuccess = true
)

type SSLUniversalSettingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SSLUniversalSettingGetResponseEnvelope struct {
	Errors   []SSLUniversalSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLUniversalSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLUniversalSettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLUniversalSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslUniversalSettingGetResponseEnvelopeJSON    `json:"-"`
}

// sslUniversalSettingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SSLUniversalSettingGetResponseEnvelope]
type sslUniversalSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLUniversalSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLUniversalSettingGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    sslUniversalSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// sslUniversalSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SSLUniversalSettingGetResponseEnvelopeErrors]
type sslUniversalSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLUniversalSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLUniversalSettingGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    sslUniversalSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// sslUniversalSettingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SSLUniversalSettingGetResponseEnvelopeMessages]
type sslUniversalSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLUniversalSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLUniversalSettingGetResponseEnvelopeSuccess bool

const (
	SSLUniversalSettingGetResponseEnvelopeSuccessTrue SSLUniversalSettingGetResponseEnvelopeSuccess = true
)
