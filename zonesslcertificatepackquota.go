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

// ZoneSslCertificatePackQuotaService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSslCertificatePackQuotaService] method instead.
type ZoneSslCertificatePackQuotaService struct {
	Options []option.RequestOption
}

// NewZoneSslCertificatePackQuotaService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSslCertificatePackQuotaService(opts ...option.RequestOption) (r *ZoneSslCertificatePackQuotaService) {
	r = &ZoneSslCertificatePackQuotaService{}
	r.Options = opts
	return
}

// For a given zone, list certificate pack quotas.
func (r *ZoneSslCertificatePackQuotaService) CertificatePacksGetCertificatePackQuotas(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *CertificatePackQuotaResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/quota", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CertificatePackQuotaResponse struct {
	Errors   []CertificatePackQuotaResponseError   `json:"errors"`
	Messages []CertificatePackQuotaResponseMessage `json:"messages"`
	Result   CertificatePackQuotaResponseResult    `json:"result"`
	// Whether the API call was successful
	Success CertificatePackQuotaResponseSuccess `json:"success"`
	JSON    certificatePackQuotaResponseJSON    `json:"-"`
}

// certificatePackQuotaResponseJSON contains the JSON metadata for the struct
// [CertificatePackQuotaResponse]
type certificatePackQuotaResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatePackQuotaResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    certificatePackQuotaResponseErrorJSON `json:"-"`
}

// certificatePackQuotaResponseErrorJSON contains the JSON metadata for the struct
// [CertificatePackQuotaResponseError]
type certificatePackQuotaResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackQuotaResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatePackQuotaResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    certificatePackQuotaResponseMessageJSON `json:"-"`
}

// certificatePackQuotaResponseMessageJSON contains the JSON metadata for the
// struct [CertificatePackQuotaResponseMessage]
type certificatePackQuotaResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackQuotaResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatePackQuotaResponseResult struct {
	Advanced CertificatePackQuotaResponseResultAdvanced `json:"advanced"`
	JSON     certificatePackQuotaResponseResultJSON     `json:"-"`
}

// certificatePackQuotaResponseResultJSON contains the JSON metadata for the struct
// [CertificatePackQuotaResponseResult]
type certificatePackQuotaResponseResultJSON struct {
	Advanced    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackQuotaResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificatePackQuotaResponseResultAdvanced struct {
	// Quantity Allocated.
	Allocated int64 `json:"allocated"`
	// Quantity Used.
	Used int64                                          `json:"used"`
	JSON certificatePackQuotaResponseResultAdvancedJSON `json:"-"`
}

// certificatePackQuotaResponseResultAdvancedJSON contains the JSON metadata for
// the struct [CertificatePackQuotaResponseResultAdvanced]
type certificatePackQuotaResponseResultAdvancedJSON struct {
	Allocated   apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackQuotaResponseResultAdvanced) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificatePackQuotaResponseSuccess bool

const (
	CertificatePackQuotaResponseSuccessTrue CertificatePackQuotaResponseSuccess = true
)
