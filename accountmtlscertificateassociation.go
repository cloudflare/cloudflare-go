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
func (r *AccountMtlsCertificateAssociationService) MTlsCertificateManagementListMTlsCertificateAssociations(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s/associations", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponse struct {
	Errors     []AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseError    `json:"errors"`
	Messages   []AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseMessage  `json:"messages"`
	Result     []AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResult   `json:"result"`
	ResultInfo AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseSuccess `json:"success"`
	JSON    accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseJSON    `json:"-"`
}

// accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponse]
type accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseError struct {
	Code    int64                                                                                                      `json:"code,required"`
	Message string                                                                                                     `json:"message,required"`
	JSON    accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseErrorJSON `json:"-"`
}

// accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseError]
type accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseMessage struct {
	Code    int64                                                                                                        `json:"code,required"`
	Message string                                                                                                       `json:"message,required"`
	JSON    accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseMessageJSON `json:"-"`
}

// accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseMessage]
type accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResult struct {
	// The service using the certificate.
	Service string `json:"service"`
	// Certificate deployment status for the given service.
	Status string                                                                                                      `json:"status"`
	JSON   accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResultJSON `json:"-"`
}

// accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResult]
type accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResultJSON struct {
	Service     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                         `json:"total_count"`
	JSON       accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResultInfoJSON `json:"-"`
}

// accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResultInfo]
type accountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseSuccess bool

const (
	AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseSuccessTrue AccountMtlsCertificateAssociationMTlsCertificateManagementListMTlsCertificateAssociationsResponseSuccess = true
)
