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

// AccountMtlsCertificateAssociationService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountMtlsCertificateAssociationService] method instead.
type AccountMtlsCertificateAssociationService struct {
	Options []option.RequestOption
}

// NewAccountMtlsCertificateAssociationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountMtlsCertificateAssociationService(opts ...option.RequestOption) (r *AccountMtlsCertificateAssociationService) {
	r = &AccountMtlsCertificateAssociationService{}
	r.Options = opts
	return
}

// Lists all active associations between the certificate and Cloudflare services.
func (r *AccountMtlsCertificateAssociationService) MTlsCertificateManagementListMTlsCertificateAssociations(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AssociationResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s/associations", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AssociationResponseCollection struct {
	Errors     []AssociationResponseCollectionError    `json:"errors"`
	Messages   []AssociationResponseCollectionMessage  `json:"messages"`
	Result     []AssociationResponseCollectionResult   `json:"result"`
	ResultInfo AssociationResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AssociationResponseCollectionSuccess `json:"success"`
	JSON    associationResponseCollectionJSON    `json:"-"`
}

// associationResponseCollectionJSON contains the JSON metadata for the struct
// [AssociationResponseCollection]
type associationResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssociationResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationResponseCollectionError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    associationResponseCollectionErrorJSON `json:"-"`
}

// associationResponseCollectionErrorJSON contains the JSON metadata for the struct
// [AssociationResponseCollectionError]
type associationResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssociationResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationResponseCollectionMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    associationResponseCollectionMessageJSON `json:"-"`
}

// associationResponseCollectionMessageJSON contains the JSON metadata for the
// struct [AssociationResponseCollectionMessage]
type associationResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssociationResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationResponseCollectionResult struct {
	// The service using the certificate.
	Service string `json:"service"`
	// Certificate deployment status for the given service.
	Status string                                  `json:"status"`
	JSON   associationResponseCollectionResultJSON `json:"-"`
}

// associationResponseCollectionResultJSON contains the JSON metadata for the
// struct [AssociationResponseCollectionResult]
type associationResponseCollectionResultJSON struct {
	Service     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssociationResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       associationResponseCollectionResultInfoJSON `json:"-"`
}

// associationResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [AssociationResponseCollectionResultInfo]
type associationResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssociationResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AssociationResponseCollectionSuccess bool

const (
	AssociationResponseCollectionSuccessTrue AssociationResponseCollectionSuccess = true
)
