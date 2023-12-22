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

// AccountPageProjectDomainService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountPageProjectDomainService] method instead.
type AccountPageProjectDomainService struct {
	Options []option.RequestOption
}

// NewAccountPageProjectDomainService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountPageProjectDomainService(opts ...option.RequestOption) (r *AccountPageProjectDomainService) {
	r = &AccountPageProjectDomainService{}
	r.Options = opts
	return
}

// Fetch a single domain.
func (r *AccountPageProjectDomainService) Get(ctx context.Context, accountIdentifier string, projectName string, domainName string, opts ...option.RequestOption) (res *DomainResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountIdentifier, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retry the validation status of a single domain.
func (r *AccountPageProjectDomainService) Update(ctx context.Context, accountIdentifier string, projectName string, domainName string, opts ...option.RequestOption) (res *DomainResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountIdentifier, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Delete a domain.
func (r *AccountPageProjectDomainService) Delete(ctx context.Context, accountIdentifier string, projectName string, domainName string, opts ...option.RequestOption) (res *AccountPageProjectDomainDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountIdentifier, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add a new domain for the Pages project.
func (r *AccountPageProjectDomainService) PagesDomainsAddDomain(ctx context.Context, accountIdentifier string, projectName string, body AccountPageProjectDomainPagesDomainsAddDomainParams, opts ...option.RequestOption) (res *DomainResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a list of all domains.
func (r *AccountPageProjectDomainService) PagesDomainsGetDomains(ctx context.Context, accountIdentifier string, projectName string, opts ...option.RequestOption) (res *DomainResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DomainResponseCollection struct {
	Errors     []DomainResponseCollectionError    `json:"errors"`
	Messages   []DomainResponseCollectionMessage  `json:"messages"`
	Result     []interface{}                      `json:"result"`
	ResultInfo DomainResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DomainResponseCollectionSuccess `json:"success"`
	JSON    domainResponseCollectionJSON    `json:"-"`
}

// domainResponseCollectionJSON contains the JSON metadata for the struct
// [DomainResponseCollection]
type domainResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainResponseCollectionError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    domainResponseCollectionErrorJSON `json:"-"`
}

// domainResponseCollectionErrorJSON contains the JSON metadata for the struct
// [DomainResponseCollectionError]
type domainResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainResponseCollectionMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    domainResponseCollectionMessageJSON `json:"-"`
}

// domainResponseCollectionMessageJSON contains the JSON metadata for the struct
// [DomainResponseCollectionMessage]
type domainResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainResponseCollectionResultInfo struct {
	Count      interface{}                            `json:"count"`
	Page       interface{}                            `json:"page"`
	PerPage    interface{}                            `json:"per_page"`
	TotalCount interface{}                            `json:"total_count"`
	JSON       domainResponseCollectionResultInfoJSON `json:"-"`
}

// domainResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [DomainResponseCollectionResultInfo]
type domainResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DomainResponseCollectionSuccess bool

const (
	DomainResponseCollectionSuccessTrue DomainResponseCollectionSuccess = true
)

type DomainResponseSingle struct {
	Errors   []DomainResponseSingleError   `json:"errors"`
	Messages []DomainResponseSingleMessage `json:"messages"`
	Result   interface{}                   `json:"result"`
	// Whether the API call was successful
	Success DomainResponseSingleSuccess `json:"success"`
	JSON    domainResponseSingleJSON    `json:"-"`
}

// domainResponseSingleJSON contains the JSON metadata for the struct
// [DomainResponseSingle]
type domainResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainResponseSingleError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    domainResponseSingleErrorJSON `json:"-"`
}

// domainResponseSingleErrorJSON contains the JSON metadata for the struct
// [DomainResponseSingleError]
type domainResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainResponseSingleMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    domainResponseSingleMessageJSON `json:"-"`
}

// domainResponseSingleMessageJSON contains the JSON metadata for the struct
// [DomainResponseSingleMessage]
type domainResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DomainResponseSingleSuccess bool

const (
	DomainResponseSingleSuccessTrue DomainResponseSingleSuccess = true
)

type AccountPageProjectDomainDeleteResponse = interface{}

type AccountPageProjectDomainPagesDomainsAddDomainParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountPageProjectDomainPagesDomainsAddDomainParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
