// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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
func (r *SSLCertificatePackQuotaService) CertificatePacksGetCertificatePackQuotas(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/quota", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponse struct {
	Advanced SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseAdvanced `json:"advanced"`
	JSON     sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseJSON     `json:"-"`
}

// sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseJSON
// contains the JSON metadata for the struct
// [SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponse]
type sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseJSON struct {
	Advanced    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseAdvanced struct {
	// Quantity Allocated.
	Allocated int64 `json:"allocated"`
	// Quantity Used.
	Used int64                                                                               `json:"used"`
	JSON sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseAdvancedJSON `json:"-"`
}

// sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseAdvancedJSON
// contains the JSON metadata for the struct
// [SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseAdvanced]
type sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseAdvancedJSON struct {
	Allocated   apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseAdvanced) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelope struct {
	Errors   []SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeJSON    `json:"-"`
}

// sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelope]
type sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeErrors struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeErrorsJSON `json:"-"`
}

// sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeErrors]
type sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeMessages struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeMessagesJSON `json:"-"`
}

// sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeMessages]
type sslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeSuccess bool

const (
	SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeSuccessTrue SSLCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseEnvelopeSuccess = true
)
