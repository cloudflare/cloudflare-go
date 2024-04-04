// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssl

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// UniversalSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUniversalSettingService] method
// instead.
type UniversalSettingService struct {
	Options []option.RequestOption
}

// NewUniversalSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUniversalSettingService(opts ...option.RequestOption) (r *UniversalSettingService) {
	r = &UniversalSettingService{}
	r.Options = opts
	return
}

// Patch Universal SSL Settings for a Zone.
func (r *UniversalSettingService) Edit(ctx context.Context, params UniversalSettingEditParams, opts ...option.RequestOption) (res *UniversalSSLSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env UniversalSettingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/universal/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Universal SSL Settings for a Zone.
func (r *UniversalSettingService) Get(ctx context.Context, query UniversalSettingGetParams, opts ...option.RequestOption) (res *UniversalSSLSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env UniversalSettingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/universal/settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UniversalSSLSettings struct {
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
	Enabled bool                     `json:"enabled"`
	JSON    universalSSLSettingsJSON `json:"-"`
}

// universalSSLSettingsJSON contains the JSON metadata for the struct
// [UniversalSSLSettings]
type universalSSLSettingsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UniversalSSLSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r universalSSLSettingsJSON) RawJSON() string {
	return r.raw
}

type UniversalSettingEditParams struct {
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

func (r UniversalSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UniversalSettingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   UniversalSSLSettings  `json:"result,required"`
	// Whether the API call was successful
	Success UniversalSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    universalSettingEditResponseEnvelopeJSON    `json:"-"`
}

// universalSettingEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [UniversalSettingEditResponseEnvelope]
type universalSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UniversalSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r universalSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UniversalSettingEditResponseEnvelopeSuccess bool

const (
	UniversalSettingEditResponseEnvelopeSuccessTrue UniversalSettingEditResponseEnvelopeSuccess = true
)

func (r UniversalSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UniversalSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UniversalSettingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type UniversalSettingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   UniversalSSLSettings  `json:"result,required"`
	// Whether the API call was successful
	Success UniversalSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    universalSettingGetResponseEnvelopeJSON    `json:"-"`
}

// universalSettingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [UniversalSettingGetResponseEnvelope]
type universalSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UniversalSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r universalSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UniversalSettingGetResponseEnvelopeSuccess bool

const (
	UniversalSettingGetResponseEnvelopeSuccessTrue UniversalSettingGetResponseEnvelopeSuccess = true
)

func (r UniversalSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UniversalSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
