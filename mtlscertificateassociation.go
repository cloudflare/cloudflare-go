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

// MtlsCertificateAssociationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewMtlsCertificateAssociationService] method instead.
type MtlsCertificateAssociationService struct {
	Options []option.RequestOption
}

// NewMtlsCertificateAssociationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMtlsCertificateAssociationService(opts ...option.RequestOption) (r *MtlsCertificateAssociationService) {
	r = &MtlsCertificateAssociationService{}
	r.Options = opts
	return
}

// Lists all active associations between the certificate and Cloudflare services.
func (r *MtlsCertificateAssociationService) List(ctx context.Context, accountID string, mtlsCertificateID string, opts ...option.RequestOption) (res *[]MtlsCertificateAssociationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MtlsCertificateAssociationListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s/associations", accountID, mtlsCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MtlsCertificateAssociationListResponse struct {
	// The service using the certificate.
	Service string `json:"service"`
	// Certificate deployment status for the given service.
	Status string                                     `json:"status"`
	JSON   mtlsCertificateAssociationListResponseJSON `json:"-"`
}

// mtlsCertificateAssociationListResponseJSON contains the JSON metadata for the
// struct [MtlsCertificateAssociationListResponse]
type mtlsCertificateAssociationListResponseJSON struct {
	Service     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateAssociationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateAssociationListResponseEnvelope struct {
	Errors   []MtlsCertificateAssociationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MtlsCertificateAssociationListResponseEnvelopeMessages `json:"messages,required"`
	Result   []MtlsCertificateAssociationListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    MtlsCertificateAssociationListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo MtlsCertificateAssociationListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       mtlsCertificateAssociationListResponseEnvelopeJSON       `json:"-"`
}

// mtlsCertificateAssociationListResponseEnvelopeJSON contains the JSON metadata
// for the struct [MtlsCertificateAssociationListResponseEnvelope]
type mtlsCertificateAssociationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateAssociationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateAssociationListResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    mtlsCertificateAssociationListResponseEnvelopeErrorsJSON `json:"-"`
}

// mtlsCertificateAssociationListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MtlsCertificateAssociationListResponseEnvelopeErrors]
type mtlsCertificateAssociationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateAssociationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MtlsCertificateAssociationListResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    mtlsCertificateAssociationListResponseEnvelopeMessagesJSON `json:"-"`
}

// mtlsCertificateAssociationListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MtlsCertificateAssociationListResponseEnvelopeMessages]
type mtlsCertificateAssociationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateAssociationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MtlsCertificateAssociationListResponseEnvelopeSuccess bool

const (
	MtlsCertificateAssociationListResponseEnvelopeSuccessTrue MtlsCertificateAssociationListResponseEnvelopeSuccess = true
)

type MtlsCertificateAssociationListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       mtlsCertificateAssociationListResponseEnvelopeResultInfoJSON `json:"-"`
}

// mtlsCertificateAssociationListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [MtlsCertificateAssociationListResponseEnvelopeResultInfo]
type mtlsCertificateAssociationListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MtlsCertificateAssociationListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
