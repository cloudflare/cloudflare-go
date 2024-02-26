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

// SSLCertificatePackQuotaService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSSLCertificatePackQuotaService] method instead.
type SSLCertificatePackQuotaService struct {
	Options []option.RequestOption
}

// NewSSLCertificatePackQuotaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSSLCertificatePackQuotaService(opts ...option.RequestOption) (r *SSLCertificatePackQuotaService) {
	r = &SSLCertificatePackQuotaService{}
	r.Options = opts
	return
}

// For a given zone, list certificate pack quotas.
func (r *SSLCertificatePackQuotaService) List(ctx context.Context, query SSLCertificatePackQuotaListParams, opts ...option.RequestOption) (res *SSLCertificatePackQuotaListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLCertificatePackQuotaListResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/quota", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SSLCertificatePackQuotaListResponse struct {
	Advanced SSLCertificatePackQuotaListResponseAdvanced `json:"advanced"`
	JSON     sslCertificatePackQuotaListResponseJSON     `json:"-"`
}

// sslCertificatePackQuotaListResponseJSON contains the JSON metadata for the
// struct [SSLCertificatePackQuotaListResponse]
type sslCertificatePackQuotaListResponseJSON struct {
	Advanced    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackQuotaListResponseAdvanced struct {
	// Quantity Allocated.
	Allocated int64 `json:"allocated"`
	// Quantity Used.
	Used int64                                           `json:"used"`
	JSON sslCertificatePackQuotaListResponseAdvancedJSON `json:"-"`
}

// sslCertificatePackQuotaListResponseAdvancedJSON contains the JSON metadata for
// the struct [SSLCertificatePackQuotaListResponseAdvanced]
type sslCertificatePackQuotaListResponseAdvancedJSON struct {
	Allocated   apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaListResponseAdvanced) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackQuotaListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SSLCertificatePackQuotaListResponseEnvelope struct {
	Errors   []SSLCertificatePackQuotaListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLCertificatePackQuotaListResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLCertificatePackQuotaListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLCertificatePackQuotaListResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslCertificatePackQuotaListResponseEnvelopeJSON    `json:"-"`
}

// sslCertificatePackQuotaListResponseEnvelopeJSON contains the JSON metadata for
// the struct [SSLCertificatePackQuotaListResponseEnvelope]
type sslCertificatePackQuotaListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackQuotaListResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    sslCertificatePackQuotaListResponseEnvelopeErrorsJSON `json:"-"`
}

// sslCertificatePackQuotaListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SSLCertificatePackQuotaListResponseEnvelopeErrors]
type sslCertificatePackQuotaListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackQuotaListResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    sslCertificatePackQuotaListResponseEnvelopeMessagesJSON `json:"-"`
}

// sslCertificatePackQuotaListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SSLCertificatePackQuotaListResponseEnvelopeMessages]
type sslCertificatePackQuotaListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLCertificatePackQuotaListResponseEnvelopeSuccess bool

const (
	SSLCertificatePackQuotaListResponseEnvelopeSuccessTrue SSLCertificatePackQuotaListResponseEnvelopeSuccess = true
)
