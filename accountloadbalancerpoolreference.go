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

// AccountLoadBalancerPoolReferenceService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountLoadBalancerPoolReferenceService] method instead.
type AccountLoadBalancerPoolReferenceService struct {
	Options []option.RequestOption
}

// NewAccountLoadBalancerPoolReferenceService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerPoolReferenceService(opts ...option.RequestOption) (r *AccountLoadBalancerPoolReferenceService) {
	r = &AccountLoadBalancerPoolReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided pool.
func (r *AccountLoadBalancerPoolReferenceService) AccountLoadBalancerPoolsListPoolReferences(ctx context.Context, accountIdentifier string, identifier interface{}, opts ...option.RequestOption) (res *SchemasReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%v/references", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SchemasReferencesResponse struct {
	Errors   []SchemasReferencesResponseError   `json:"errors"`
	Messages []SchemasReferencesResponseMessage `json:"messages"`
	// List of resources that reference a given pool.
	Result     []SchemasReferencesResponseResult   `json:"result"`
	ResultInfo SchemasReferencesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SchemasReferencesResponseSuccess `json:"success"`
	JSON    schemasReferencesResponseJSON    `json:"-"`
}

// schemasReferencesResponseJSON contains the JSON metadata for the struct
// [SchemasReferencesResponse]
type schemasReferencesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasReferencesResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    schemasReferencesResponseErrorJSON `json:"-"`
}

// schemasReferencesResponseErrorJSON contains the JSON metadata for the struct
// [SchemasReferencesResponseError]
type schemasReferencesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasReferencesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasReferencesResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    schemasReferencesResponseMessageJSON `json:"-"`
}

// schemasReferencesResponseMessageJSON contains the JSON metadata for the struct
// [SchemasReferencesResponseMessage]
type schemasReferencesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasReferencesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasReferencesResponseResult struct {
	ReferenceType SchemasReferencesResponseResultReferenceType `json:"reference_type"`
	ResourceID    string                                       `json:"resource_id"`
	ResourceName  string                                       `json:"resource_name"`
	ResourceType  string                                       `json:"resource_type"`
	JSON          schemasReferencesResponseResultJSON          `json:"-"`
}

// schemasReferencesResponseResultJSON contains the JSON metadata for the struct
// [SchemasReferencesResponseResult]
type schemasReferencesResponseResultJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SchemasReferencesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasReferencesResponseResultReferenceType string

const (
	SchemasReferencesResponseResultReferenceTypeStar     SchemasReferencesResponseResultReferenceType = "*"
	SchemasReferencesResponseResultReferenceTypeReferral SchemasReferencesResponseResultReferenceType = "referral"
	SchemasReferencesResponseResultReferenceTypeReferrer SchemasReferencesResponseResultReferenceType = "referrer"
)

type SchemasReferencesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       schemasReferencesResponseResultInfoJSON `json:"-"`
}

// schemasReferencesResponseResultInfoJSON contains the JSON metadata for the
// struct [SchemasReferencesResponseResultInfo]
type schemasReferencesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasReferencesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasReferencesResponseSuccess bool

const (
	SchemasReferencesResponseSuccessTrue SchemasReferencesResponseSuccess = true
)
