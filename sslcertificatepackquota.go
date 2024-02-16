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
func (r *SSLCertificatePackQuotaService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SSLCertificatePackQuotaGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLCertificatePackQuotaGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/quota", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SSLCertificatePackQuotaGetResponse struct {
	Advanced SSLCertificatePackQuotaGetResponseAdvanced `json:"advanced"`
	JSON     sslCertificatePackQuotaGetResponseJSON     `json:"-"`
}

// sslCertificatePackQuotaGetResponseJSON contains the JSON metadata for the struct
// [SSLCertificatePackQuotaGetResponse]
type sslCertificatePackQuotaGetResponseJSON struct {
	Advanced    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackQuotaGetResponseAdvanced struct {
	// Quantity Allocated.
	Allocated int64 `json:"allocated"`
	// Quantity Used.
	Used int64                                          `json:"used"`
	JSON sslCertificatePackQuotaGetResponseAdvancedJSON `json:"-"`
}

// sslCertificatePackQuotaGetResponseAdvancedJSON contains the JSON metadata for
// the struct [SSLCertificatePackQuotaGetResponseAdvanced]
type sslCertificatePackQuotaGetResponseAdvancedJSON struct {
	Allocated   apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaGetResponseAdvanced) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackQuotaGetResponseEnvelope struct {
	Errors   []SSLCertificatePackQuotaGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLCertificatePackQuotaGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLCertificatePackQuotaGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLCertificatePackQuotaGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslCertificatePackQuotaGetResponseEnvelopeJSON    `json:"-"`
}

// sslCertificatePackQuotaGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SSLCertificatePackQuotaGetResponseEnvelope]
type sslCertificatePackQuotaGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackQuotaGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    sslCertificatePackQuotaGetResponseEnvelopeErrorsJSON `json:"-"`
}

// sslCertificatePackQuotaGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SSLCertificatePackQuotaGetResponseEnvelopeErrors]
type sslCertificatePackQuotaGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLCertificatePackQuotaGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    sslCertificatePackQuotaGetResponseEnvelopeMessagesJSON `json:"-"`
}

// sslCertificatePackQuotaGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SSLCertificatePackQuotaGetResponseEnvelopeMessages]
type sslCertificatePackQuotaGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLCertificatePackQuotaGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLCertificatePackQuotaGetResponseEnvelopeSuccess bool

const (
	SSLCertificatePackQuotaGetResponseEnvelopeSuccessTrue SSLCertificatePackQuotaGetResponseEnvelopeSuccess = true
)
