// File generated from our OpenAPI spec by Stainless.

package ssl

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// CertificatePackQuotaService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCertificatePackQuotaService]
// method instead.
type CertificatePackQuotaService struct {
	Options []option.RequestOption
}

// NewCertificatePackQuotaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCertificatePackQuotaService(opts ...option.RequestOption) (r *CertificatePackQuotaService) {
	r = &CertificatePackQuotaService{}
	r.Options = opts
	return
}

// For a given zone, list certificate pack quotas.
func (r *CertificatePackQuotaService) Get(ctx context.Context, query CertificatePackQuotaGetParams, opts ...option.RequestOption) (res *CertificatePackQuotaGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificatePackQuotaGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/quota", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CertificatePackQuotaGetResponse struct {
	Advanced CertificatePackQuotaGetResponseAdvanced `json:"advanced"`
	JSON     certificatePackQuotaGetResponseJSON     `json:"-"`
}

// certificatePackQuotaGetResponseJSON contains the JSON metadata for the struct
// [CertificatePackQuotaGetResponse]
type certificatePackQuotaGetResponseJSON struct {
	Advanced    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackQuotaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackQuotaGetResponseJSON) RawJSON() string {
	return r.raw
}

type CertificatePackQuotaGetResponseAdvanced struct {
	// Quantity Allocated.
	Allocated int64 `json:"allocated"`
	// Quantity Used.
	Used int64                                       `json:"used"`
	JSON certificatePackQuotaGetResponseAdvancedJSON `json:"-"`
}

// certificatePackQuotaGetResponseAdvancedJSON contains the JSON metadata for the
// struct [CertificatePackQuotaGetResponseAdvanced]
type certificatePackQuotaGetResponseAdvancedJSON struct {
	Allocated   apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackQuotaGetResponseAdvanced) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackQuotaGetResponseAdvancedJSON) RawJSON() string {
	return r.raw
}

type CertificatePackQuotaGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CertificatePackQuotaGetResponseEnvelope struct {
	Errors   []CertificatePackQuotaGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificatePackQuotaGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CertificatePackQuotaGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CertificatePackQuotaGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    certificatePackQuotaGetResponseEnvelopeJSON    `json:"-"`
}

// certificatePackQuotaGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CertificatePackQuotaGetResponseEnvelope]
type certificatePackQuotaGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackQuotaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackQuotaGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificatePackQuotaGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    certificatePackQuotaGetResponseEnvelopeErrorsJSON `json:"-"`
}

// certificatePackQuotaGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CertificatePackQuotaGetResponseEnvelopeErrors]
type certificatePackQuotaGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackQuotaGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackQuotaGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificatePackQuotaGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    certificatePackQuotaGetResponseEnvelopeMessagesJSON `json:"-"`
}

// certificatePackQuotaGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CertificatePackQuotaGetResponseEnvelopeMessages]
type certificatePackQuotaGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackQuotaGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackQuotaGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CertificatePackQuotaGetResponseEnvelopeSuccess bool

const (
	CertificatePackQuotaGetResponseEnvelopeSuccessTrue CertificatePackQuotaGetResponseEnvelopeSuccess = true
)
