// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRegistrarDomainService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRegistrarDomainService]
// method instead.
type AccountRegistrarDomainService struct {
	Options []option.RequestOption
}

// NewAccountRegistrarDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRegistrarDomainService(opts ...option.RequestOption) (r *AccountRegistrarDomainService) {
	r = &AccountRegistrarDomainService{}
	r.Options = opts
	return
}

// Show individual domain.
func (r *AccountRegistrarDomainService) Get(ctx context.Context, accountIdentifier string, domainName string, opts ...option.RequestOption) (res *AccountRegistrarDomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/registrar/domains/%s", accountIdentifier, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update individual domain.
func (r *AccountRegistrarDomainService) Update(ctx context.Context, accountIdentifier string, domainName string, body AccountRegistrarDomainUpdateParams, opts ...option.RequestOption) (res *AccountRegistrarDomainUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/registrar/domains/%s", accountIdentifier, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List domains handled by Registrar.
func (r *AccountRegistrarDomainService) RegistrarDomainsListDomains(ctx context.Context, accountIdentifier string, body AccountRegistrarDomainRegistrarDomainsListDomainsParams, opts ...option.RequestOption) (res *AccountRegistrarDomainRegistrarDomainsListDomainsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/registrar/domains", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountRegistrarDomainGetResponse struct {
	Errors   []AccountRegistrarDomainGetResponseError   `json:"errors"`
	Messages []AccountRegistrarDomainGetResponseMessage `json:"messages"`
	Result   interface{}                                `json:"result"`
	// Whether the API call was successful
	Success AccountRegistrarDomainGetResponseSuccess `json:"success"`
	JSON    accountRegistrarDomainGetResponseJSON    `json:"-"`
}

// accountRegistrarDomainGetResponseJSON contains the JSON metadata for the struct
// [AccountRegistrarDomainGetResponse]
type accountRegistrarDomainGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRegistrarDomainGetResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountRegistrarDomainGetResponseErrorJSON `json:"-"`
}

// accountRegistrarDomainGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountRegistrarDomainGetResponseError]
type accountRegistrarDomainGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRegistrarDomainGetResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountRegistrarDomainGetResponseMessageJSON `json:"-"`
}

// accountRegistrarDomainGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountRegistrarDomainGetResponseMessage]
type accountRegistrarDomainGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRegistrarDomainGetResponseSuccess bool

const (
	AccountRegistrarDomainGetResponseSuccessTrue AccountRegistrarDomainGetResponseSuccess = true
)

type AccountRegistrarDomainUpdateResponse struct {
	Errors   []AccountRegistrarDomainUpdateResponseError   `json:"errors"`
	Messages []AccountRegistrarDomainUpdateResponseMessage `json:"messages"`
	Result   interface{}                                   `json:"result"`
	// Whether the API call was successful
	Success AccountRegistrarDomainUpdateResponseSuccess `json:"success"`
	JSON    accountRegistrarDomainUpdateResponseJSON    `json:"-"`
}

// accountRegistrarDomainUpdateResponseJSON contains the JSON metadata for the
// struct [AccountRegistrarDomainUpdateResponse]
type accountRegistrarDomainUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRegistrarDomainUpdateResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountRegistrarDomainUpdateResponseErrorJSON `json:"-"`
}

// accountRegistrarDomainUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountRegistrarDomainUpdateResponseError]
type accountRegistrarDomainUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRegistrarDomainUpdateResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountRegistrarDomainUpdateResponseMessageJSON `json:"-"`
}

// accountRegistrarDomainUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountRegistrarDomainUpdateResponseMessage]
type accountRegistrarDomainUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRegistrarDomainUpdateResponseSuccess bool

const (
	AccountRegistrarDomainUpdateResponseSuccessTrue AccountRegistrarDomainUpdateResponseSuccess = true
)

type AccountRegistrarDomainRegistrarDomainsListDomainsResponse struct {
	Errors     []AccountRegistrarDomainRegistrarDomainsListDomainsResponseError    `json:"errors"`
	Messages   []AccountRegistrarDomainRegistrarDomainsListDomainsResponseMessage  `json:"messages"`
	Result     []AccountRegistrarDomainRegistrarDomainsListDomainsResponseResult   `json:"result"`
	ResultInfo AccountRegistrarDomainRegistrarDomainsListDomainsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountRegistrarDomainRegistrarDomainsListDomainsResponseSuccess `json:"success"`
	JSON    accountRegistrarDomainRegistrarDomainsListDomainsResponseJSON    `json:"-"`
}

// accountRegistrarDomainRegistrarDomainsListDomainsResponseJSON contains the JSON
// metadata for the struct
// [AccountRegistrarDomainRegistrarDomainsListDomainsResponse]
type accountRegistrarDomainRegistrarDomainsListDomainsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainRegistrarDomainsListDomainsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRegistrarDomainRegistrarDomainsListDomainsResponseError struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accountRegistrarDomainRegistrarDomainsListDomainsResponseErrorJSON `json:"-"`
}

// accountRegistrarDomainRegistrarDomainsListDomainsResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountRegistrarDomainRegistrarDomainsListDomainsResponseError]
type accountRegistrarDomainRegistrarDomainsListDomainsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainRegistrarDomainsListDomainsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRegistrarDomainRegistrarDomainsListDomainsResponseMessage struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    accountRegistrarDomainRegistrarDomainsListDomainsResponseMessageJSON `json:"-"`
}

// accountRegistrarDomainRegistrarDomainsListDomainsResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountRegistrarDomainRegistrarDomainsListDomainsResponseMessage]
type accountRegistrarDomainRegistrarDomainsListDomainsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainRegistrarDomainsListDomainsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRegistrarDomainRegistrarDomainsListDomainsResponseResult struct {
	// Domain identifier.
	ID string `json:"id"`
	// Shows if a domain is available for transferring into Cloudflare Registrar.
	Available bool `json:"available"`
	// Indicates if the domain can be registered as a new domain.
	CanRegister bool `json:"can_register"`
	// Shows time of creation.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Shows name of current registrar.
	CurrentRegistrar string `json:"current_registrar"`
	// Shows when domain name registration expires.
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Shows whether a registrar lock is in place for a domain.
	Locked            bool                                                                             `json:"locked"`
	RegistrantContact AccountRegistrarDomainRegistrarDomainsListDomainsResponseResultRegistrantContact `json:"registrant_contact"`
	// A comma-separated list of registry status codes. A full list of status codes can
	// be found at
	// [EPP Status Codes](https://www.icann.org/resources/pages/epp-status-codes-2014-06-16-en).
	RegistryStatuses string `json:"registry_statuses"`
	// Whether a particular TLD is currently supported by Cloudflare Registrar. Refer
	// to [TLD Policies](https://www.cloudflare.com/tld-policies/) for a list of
	// supported TLDs.
	SupportedTld bool `json:"supported_tld"`
	// Statuses for domain transfers into Cloudflare Registrar.
	TransferIn AccountRegistrarDomainRegistrarDomainsListDomainsResponseResultTransferIn `json:"transfer_in"`
	// Last updated.
	UpdatedAt time.Time                                                           `json:"updated_at" format:"date-time"`
	JSON      accountRegistrarDomainRegistrarDomainsListDomainsResponseResultJSON `json:"-"`
}

// accountRegistrarDomainRegistrarDomainsListDomainsResponseResultJSON contains the
// JSON metadata for the struct
// [AccountRegistrarDomainRegistrarDomainsListDomainsResponseResult]
type accountRegistrarDomainRegistrarDomainsListDomainsResponseResultJSON struct {
	ID                apijson.Field
	Available         apijson.Field
	CanRegister       apijson.Field
	CreatedAt         apijson.Field
	CurrentRegistrar  apijson.Field
	ExpiresAt         apijson.Field
	Locked            apijson.Field
	RegistrantContact apijson.Field
	RegistryStatuses  apijson.Field
	SupportedTld      apijson.Field
	TransferIn        apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountRegistrarDomainRegistrarDomainsListDomainsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRegistrarDomainRegistrarDomainsListDomainsResponseResultRegistrantContact struct {
	// Address.
	Address string `json:"address,required"`
	// City.
	City string `json:"city,required"`
	// The country in which the user lives.
	Country string `json:"country,required,nullable"`
	// User's first name
	FirstName string `json:"first_name,required,nullable"`
	// User's last name
	LastName string `json:"last_name,required,nullable"`
	// Name of organization.
	Organization string `json:"organization,required"`
	// User's telephone number
	Phone string `json:"phone,required,nullable"`
	// State.
	State string `json:"state,required"`
	// The zipcode or postal code where the user lives.
	Zip string `json:"zip,required,nullable"`
	// Contact Identifier.
	ID string `json:"id"`
	// Optional address line for unit, floor, suite, etc.
	Address2 string `json:"address2"`
	// The contact email address of the user.
	Email string `json:"email"`
	// Contact fax number.
	Fax  string                                                                               `json:"fax"`
	JSON accountRegistrarDomainRegistrarDomainsListDomainsResponseResultRegistrantContactJSON `json:"-"`
}

// accountRegistrarDomainRegistrarDomainsListDomainsResponseResultRegistrantContactJSON
// contains the JSON metadata for the struct
// [AccountRegistrarDomainRegistrarDomainsListDomainsResponseResultRegistrantContact]
type accountRegistrarDomainRegistrarDomainsListDomainsResponseResultRegistrantContactJSON struct {
	Address      apijson.Field
	City         apijson.Field
	Country      apijson.Field
	FirstName    apijson.Field
	LastName     apijson.Field
	Organization apijson.Field
	Phone        apijson.Field
	State        apijson.Field
	Zip          apijson.Field
	ID           apijson.Field
	Address2     apijson.Field
	Email        apijson.Field
	Fax          apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountRegistrarDomainRegistrarDomainsListDomainsResponseResultRegistrantContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Statuses for domain transfers into Cloudflare Registrar.
type AccountRegistrarDomainRegistrarDomainsListDomainsResponseResultTransferIn struct {
	// Form of authorization has been accepted by the registrant.
	AcceptFoa interface{} `json:"accept_foa"`
	// Shows transfer status with the registry.
	ApproveTransfer interface{} `json:"approve_transfer"`
	// Indicates if cancellation is still possible.
	CanCancelTransfer bool `json:"can_cancel_transfer"`
	// Privacy guards are disabled at the foreign registrar.
	DisablePrivacy interface{} `json:"disable_privacy"`
	// Auth code has been entered and verified.
	EnterAuthCode interface{} `json:"enter_auth_code"`
	// Domain is unlocked at the foreign registrar.
	UnlockDomain interface{}                                                                   `json:"unlock_domain"`
	JSON         accountRegistrarDomainRegistrarDomainsListDomainsResponseResultTransferInJSON `json:"-"`
}

// accountRegistrarDomainRegistrarDomainsListDomainsResponseResultTransferInJSON
// contains the JSON metadata for the struct
// [AccountRegistrarDomainRegistrarDomainsListDomainsResponseResultTransferIn]
type accountRegistrarDomainRegistrarDomainsListDomainsResponseResultTransferInJSON struct {
	AcceptFoa         apijson.Field
	ApproveTransfer   apijson.Field
	CanCancelTransfer apijson.Field
	DisablePrivacy    apijson.Field
	EnterAuthCode     apijson.Field
	UnlockDomain      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountRegistrarDomainRegistrarDomainsListDomainsResponseResultTransferIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRegistrarDomainRegistrarDomainsListDomainsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                 `json:"total_count"`
	JSON       accountRegistrarDomainRegistrarDomainsListDomainsResponseResultInfoJSON `json:"-"`
}

// accountRegistrarDomainRegistrarDomainsListDomainsResponseResultInfoJSON contains
// the JSON metadata for the struct
// [AccountRegistrarDomainRegistrarDomainsListDomainsResponseResultInfo]
type accountRegistrarDomainRegistrarDomainsListDomainsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainRegistrarDomainsListDomainsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRegistrarDomainRegistrarDomainsListDomainsResponseSuccess bool

const (
	AccountRegistrarDomainRegistrarDomainsListDomainsResponseSuccessTrue AccountRegistrarDomainRegistrarDomainsListDomainsResponseSuccess = true
)

type AccountRegistrarDomainUpdateParams struct {
	// Auto-renew controls whether subscription is automatically renewed upon domain
	// expiration.
	AutoRenew param.Field[bool] `json:"auto_renew"`
	// Shows whether a registrar lock is in place for a domain.
	Locked param.Field[bool] `json:"locked"`
	// List of name servers.
	NameServers param.Field[[]string] `json:"name_servers"`
	// Privacy option controls redacting WHOIS information.
	Privacy param.Field[bool] `json:"privacy"`
}

func (r AccountRegistrarDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRegistrarDomainRegistrarDomainsListDomainsParams struct {
	// List of domain names.
	ID param.Field[[]string] `json:"id,required"`
}

func (r AccountRegistrarDomainRegistrarDomainsListDomainsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
