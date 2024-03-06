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

// MTLSCertificateAssociationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewMTLSCertificateAssociationService] method instead.
type MTLSCertificateAssociationService struct {
	Options []option.RequestOption
}

// NewMTLSCertificateAssociationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMTLSCertificateAssociationService(opts ...option.RequestOption) (r *MTLSCertificateAssociationService) {
	r = &MTLSCertificateAssociationService{}
	r.Options = opts
	return
}

// Lists all active associations between the certificate and Cloudflare services.
func (r *MTLSCertificateAssociationService) Get(ctx context.Context, mtlsCertificateID string, query MTLSCertificateAssociationGetParams, opts ...option.RequestOption) (res *[]MTLSCertificateAssociationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MTLSCertificateAssociationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s/associations", query.AccountID, mtlsCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MTLSCertificateAssociationGetResponse struct {
	// The service using the certificate.
	Service string `json:"service"`
	// Certificate deployment status for the given service.
	Status string                                    `json:"status"`
	JSON   mtlsCertificateAssociationGetResponseJSON `json:"-"`
}

// mtlsCertificateAssociationGetResponseJSON contains the JSON metadata for the
// struct [MTLSCertificateAssociationGetResponse]
type mtlsCertificateAssociationGetResponseJSON struct {
	Service     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MTLSCertificateAssociationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mtlsCertificateAssociationGetResponseJSON) RawJSON() string {
	return r.raw
}

type MTLSCertificateAssociationGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type MTLSCertificateAssociationGetResponseEnvelope struct {
	Errors   []MTLSCertificateAssociationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MTLSCertificateAssociationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []MTLSCertificateAssociationGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    MTLSCertificateAssociationGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo MTLSCertificateAssociationGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       mtlsCertificateAssociationGetResponseEnvelopeJSON       `json:"-"`
}

// mtlsCertificateAssociationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [MTLSCertificateAssociationGetResponseEnvelope]
type mtlsCertificateAssociationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MTLSCertificateAssociationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mtlsCertificateAssociationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MTLSCertificateAssociationGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    mtlsCertificateAssociationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// mtlsCertificateAssociationGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MTLSCertificateAssociationGetResponseEnvelopeErrors]
type mtlsCertificateAssociationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MTLSCertificateAssociationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mtlsCertificateAssociationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MTLSCertificateAssociationGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    mtlsCertificateAssociationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// mtlsCertificateAssociationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MTLSCertificateAssociationGetResponseEnvelopeMessages]
type mtlsCertificateAssociationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MTLSCertificateAssociationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mtlsCertificateAssociationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MTLSCertificateAssociationGetResponseEnvelopeSuccess bool

const (
	MTLSCertificateAssociationGetResponseEnvelopeSuccessTrue MTLSCertificateAssociationGetResponseEnvelopeSuccess = true
)

type MTLSCertificateAssociationGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       mtlsCertificateAssociationGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// mtlsCertificateAssociationGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [MTLSCertificateAssociationGetResponseEnvelopeResultInfo]
type mtlsCertificateAssociationGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MTLSCertificateAssociationGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mtlsCertificateAssociationGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
