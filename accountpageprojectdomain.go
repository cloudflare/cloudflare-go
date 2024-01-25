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
func (r *AccountPageProjectDomainService) Get(ctx context.Context, accountIdentifier string, projectName string, domainName string, opts ...option.RequestOption) (res *AccountPageProjectDomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountIdentifier, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retry the validation status of a single domain.
func (r *AccountPageProjectDomainService) Update(ctx context.Context, accountIdentifier string, projectName string, domainName string, opts ...option.RequestOption) (res *AccountPageProjectDomainUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountIdentifier, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Delete a Pages project's domain.
func (r *AccountPageProjectDomainService) Delete(ctx context.Context, accountIdentifier string, projectName string, domainName string, opts ...option.RequestOption) (res *AccountPageProjectDomainDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountIdentifier, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add a new domain for the Pages project.
func (r *AccountPageProjectDomainService) PagesDomainsAddDomain(ctx context.Context, accountIdentifier string, projectName string, body AccountPageProjectDomainPagesDomainsAddDomainParams, opts ...option.RequestOption) (res *AccountPageProjectDomainPagesDomainsAddDomainResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a list of all domains associated with a Pages project.
func (r *AccountPageProjectDomainService) PagesDomainsGetDomains(ctx context.Context, accountIdentifier string, projectName string, opts ...option.RequestOption) (res *AccountPageProjectDomainPagesDomainsGetDomainsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountPageProjectDomainGetResponse struct {
	Errors   []AccountPageProjectDomainGetResponseError   `json:"errors"`
	Messages []AccountPageProjectDomainGetResponseMessage `json:"messages"`
	Result   interface{}                                  `json:"result"`
	// Whether the API call was successful
	Success AccountPageProjectDomainGetResponseSuccess `json:"success"`
	JSON    accountPageProjectDomainGetResponseJSON    `json:"-"`
}

// accountPageProjectDomainGetResponseJSON contains the JSON metadata for the
// struct [AccountPageProjectDomainGetResponse]
type accountPageProjectDomainGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDomainGetResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountPageProjectDomainGetResponseErrorJSON `json:"-"`
}

// accountPageProjectDomainGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountPageProjectDomainGetResponseError]
type accountPageProjectDomainGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDomainGetResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountPageProjectDomainGetResponseMessageJSON `json:"-"`
}

// accountPageProjectDomainGetResponseMessageJSON contains the JSON metadata for
// the struct [AccountPageProjectDomainGetResponseMessage]
type accountPageProjectDomainGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectDomainGetResponseSuccess bool

const (
	AccountPageProjectDomainGetResponseSuccessTrue AccountPageProjectDomainGetResponseSuccess = true
)

type AccountPageProjectDomainUpdateResponse struct {
	Errors   []AccountPageProjectDomainUpdateResponseError   `json:"errors"`
	Messages []AccountPageProjectDomainUpdateResponseMessage `json:"messages"`
	Result   interface{}                                     `json:"result"`
	// Whether the API call was successful
	Success AccountPageProjectDomainUpdateResponseSuccess `json:"success"`
	JSON    accountPageProjectDomainUpdateResponseJSON    `json:"-"`
}

// accountPageProjectDomainUpdateResponseJSON contains the JSON metadata for the
// struct [AccountPageProjectDomainUpdateResponse]
type accountPageProjectDomainUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDomainUpdateResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountPageProjectDomainUpdateResponseErrorJSON `json:"-"`
}

// accountPageProjectDomainUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountPageProjectDomainUpdateResponseError]
type accountPageProjectDomainUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDomainUpdateResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountPageProjectDomainUpdateResponseMessageJSON `json:"-"`
}

// accountPageProjectDomainUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountPageProjectDomainUpdateResponseMessage]
type accountPageProjectDomainUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectDomainUpdateResponseSuccess bool

const (
	AccountPageProjectDomainUpdateResponseSuccessTrue AccountPageProjectDomainUpdateResponseSuccess = true
)

type AccountPageProjectDomainDeleteResponse = interface{}

type AccountPageProjectDomainPagesDomainsAddDomainResponse struct {
	Errors   []AccountPageProjectDomainPagesDomainsAddDomainResponseError   `json:"errors"`
	Messages []AccountPageProjectDomainPagesDomainsAddDomainResponseMessage `json:"messages"`
	Result   interface{}                                                    `json:"result"`
	// Whether the API call was successful
	Success AccountPageProjectDomainPagesDomainsAddDomainResponseSuccess `json:"success"`
	JSON    accountPageProjectDomainPagesDomainsAddDomainResponseJSON    `json:"-"`
}

// accountPageProjectDomainPagesDomainsAddDomainResponseJSON contains the JSON
// metadata for the struct [AccountPageProjectDomainPagesDomainsAddDomainResponse]
type accountPageProjectDomainPagesDomainsAddDomainResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainPagesDomainsAddDomainResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDomainPagesDomainsAddDomainResponseError struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    accountPageProjectDomainPagesDomainsAddDomainResponseErrorJSON `json:"-"`
}

// accountPageProjectDomainPagesDomainsAddDomainResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountPageProjectDomainPagesDomainsAddDomainResponseError]
type accountPageProjectDomainPagesDomainsAddDomainResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainPagesDomainsAddDomainResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDomainPagesDomainsAddDomainResponseMessage struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    accountPageProjectDomainPagesDomainsAddDomainResponseMessageJSON `json:"-"`
}

// accountPageProjectDomainPagesDomainsAddDomainResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountPageProjectDomainPagesDomainsAddDomainResponseMessage]
type accountPageProjectDomainPagesDomainsAddDomainResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainPagesDomainsAddDomainResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectDomainPagesDomainsAddDomainResponseSuccess bool

const (
	AccountPageProjectDomainPagesDomainsAddDomainResponseSuccessTrue AccountPageProjectDomainPagesDomainsAddDomainResponseSuccess = true
)

type AccountPageProjectDomainPagesDomainsGetDomainsResponse struct {
	Errors     []AccountPageProjectDomainPagesDomainsGetDomainsResponseError    `json:"errors"`
	Messages   []AccountPageProjectDomainPagesDomainsGetDomainsResponseMessage  `json:"messages"`
	Result     []interface{}                                                    `json:"result"`
	ResultInfo AccountPageProjectDomainPagesDomainsGetDomainsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountPageProjectDomainPagesDomainsGetDomainsResponseSuccess `json:"success"`
	JSON    accountPageProjectDomainPagesDomainsGetDomainsResponseJSON    `json:"-"`
}

// accountPageProjectDomainPagesDomainsGetDomainsResponseJSON contains the JSON
// metadata for the struct [AccountPageProjectDomainPagesDomainsGetDomainsResponse]
type accountPageProjectDomainPagesDomainsGetDomainsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainPagesDomainsGetDomainsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDomainPagesDomainsGetDomainsResponseError struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountPageProjectDomainPagesDomainsGetDomainsResponseErrorJSON `json:"-"`
}

// accountPageProjectDomainPagesDomainsGetDomainsResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountPageProjectDomainPagesDomainsGetDomainsResponseError]
type accountPageProjectDomainPagesDomainsGetDomainsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainPagesDomainsGetDomainsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDomainPagesDomainsGetDomainsResponseMessage struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountPageProjectDomainPagesDomainsGetDomainsResponseMessageJSON `json:"-"`
}

// accountPageProjectDomainPagesDomainsGetDomainsResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountPageProjectDomainPagesDomainsGetDomainsResponseMessage]
type accountPageProjectDomainPagesDomainsGetDomainsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainPagesDomainsGetDomainsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDomainPagesDomainsGetDomainsResponseResultInfo struct {
	Count      interface{}                                                          `json:"count"`
	Page       interface{}                                                          `json:"page"`
	PerPage    interface{}                                                          `json:"per_page"`
	TotalCount interface{}                                                          `json:"total_count"`
	JSON       accountPageProjectDomainPagesDomainsGetDomainsResponseResultInfoJSON `json:"-"`
}

// accountPageProjectDomainPagesDomainsGetDomainsResponseResultInfoJSON contains
// the JSON metadata for the struct
// [AccountPageProjectDomainPagesDomainsGetDomainsResponseResultInfo]
type accountPageProjectDomainPagesDomainsGetDomainsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainPagesDomainsGetDomainsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectDomainPagesDomainsGetDomainsResponseSuccess bool

const (
	AccountPageProjectDomainPagesDomainsGetDomainsResponseSuccessTrue AccountPageProjectDomainPagesDomainsGetDomainsResponseSuccess = true
)

type AccountPageProjectDomainPagesDomainsAddDomainParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountPageProjectDomainPagesDomainsAddDomainParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
