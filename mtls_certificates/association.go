// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AssociationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAssociationService] method
// instead.
type AssociationService struct {
	Options []option.RequestOption
}

// NewAssociationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssociationService(opts ...option.RequestOption) (r *AssociationService) {
	r = &AssociationService{}
	r.Options = opts
	return
}

// Lists all active associations between the certificate and Cloudflare services.
func (r *AssociationService) Get(ctx context.Context, mtlsCertificateID string, query AssociationGetParams, opts ...option.RequestOption) (res *[]MTLSCertificateAsssociation, err error) {
	opts = append(r.Options[:], opts...)
	var env AssociationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s/associations", query.AccountID, mtlsCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MTLSCertificateAsssociation struct {
	// The service using the certificate.
	Service string `json:"service"`
	// Certificate deployment status for the given service.
	Status string                          `json:"status"`
	JSON   mtlsCertificateAsssociationJSON `json:"-"`
}

// mtlsCertificateAsssociationJSON contains the JSON metadata for the struct
// [MTLSCertificateAsssociation]
type mtlsCertificateAsssociationJSON struct {
	Service     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MTLSCertificateAsssociation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mtlsCertificateAsssociationJSON) RawJSON() string {
	return r.raw
}

type AssociationGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AssociationGetResponseEnvelope struct {
	Errors   []AssociationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AssociationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []MTLSCertificateAsssociation            `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AssociationGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AssociationGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       associationGetResponseEnvelopeJSON       `json:"-"`
}

// associationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AssociationGetResponseEnvelope]
type associationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssociationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r associationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AssociationGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    associationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// associationGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AssociationGetResponseEnvelopeErrors]
type associationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssociationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r associationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AssociationGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    associationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// associationGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AssociationGetResponseEnvelopeMessages]
type associationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssociationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r associationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AssociationGetResponseEnvelopeSuccess bool

const (
	AssociationGetResponseEnvelopeSuccessTrue AssociationGetResponseEnvelopeSuccess = true
)

func (r AssociationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AssociationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AssociationGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       associationGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// associationGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [AssociationGetResponseEnvelopeResultInfo]
type associationGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssociationGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r associationGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
