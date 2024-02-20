// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// RegistrarDomainService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRegistrarDomainService] method
// instead.
type RegistrarDomainService struct {
	Options []option.RequestOption
}

// NewRegistrarDomainService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegistrarDomainService(opts ...option.RequestOption) (r *RegistrarDomainService) {
	r = &RegistrarDomainService{}
	r.Options = opts
	return
}

// List domains handled by Registrar.
func (r *RegistrarDomainService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]RegistrarDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RegistrarDomainListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/registrar/domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show individual domain.
func (r *RegistrarDomainService) Get(ctx context.Context, accountID string, domainName string, opts ...option.RequestOption) (res *RegistrarDomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RegistrarDomainGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/registrar/domains/%s", accountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update individual domain.
func (r *RegistrarDomainService) Replace(ctx context.Context, accountID string, domainName string, body RegistrarDomainReplaceParams, opts ...option.RequestOption) (res *RegistrarDomainReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RegistrarDomainReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/registrar/domains/%s", accountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RegistrarDomainListResponse struct {
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
	Locked bool `json:"locked"`
	// Shows contact information for domain registrant.
	RegistrantContact RegistrarDomainListResponseRegistrantContact `json:"registrant_contact"`
	// A comma-separated list of registry status codes. A full list of status codes can
	// be found at
	// [EPP Status Codes](https://www.icann.org/resources/pages/epp-status-codes-2014-06-16-en).
	RegistryStatuses string `json:"registry_statuses"`
	// Whether a particular TLD is currently supported by Cloudflare Registrar. Refer
	// to [TLD Policies](https://www.cloudflare.com/tld-policies/) for a list of
	// supported TLDs.
	SupportedTld bool `json:"supported_tld"`
	// Statuses for domain transfers into Cloudflare Registrar.
	TransferIn RegistrarDomainListResponseTransferIn `json:"transfer_in"`
	// Last updated.
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      registrarDomainListResponseJSON `json:"-"`
}

// registrarDomainListResponseJSON contains the JSON metadata for the struct
// [RegistrarDomainListResponse]
type registrarDomainListResponseJSON struct {
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

func (r *RegistrarDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Shows contact information for domain registrant.
type RegistrarDomainListResponseRegistrantContact struct {
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
	Fax  string                                           `json:"fax"`
	JSON registrarDomainListResponseRegistrantContactJSON `json:"-"`
}

// registrarDomainListResponseRegistrantContactJSON contains the JSON metadata for
// the struct [RegistrarDomainListResponseRegistrantContact]
type registrarDomainListResponseRegistrantContactJSON struct {
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

func (r *RegistrarDomainListResponseRegistrantContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Statuses for domain transfers into Cloudflare Registrar.
type RegistrarDomainListResponseTransferIn struct {
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
	UnlockDomain interface{}                               `json:"unlock_domain"`
	JSON         registrarDomainListResponseTransferInJSON `json:"-"`
}

// registrarDomainListResponseTransferInJSON contains the JSON metadata for the
// struct [RegistrarDomainListResponseTransferIn]
type registrarDomainListResponseTransferInJSON struct {
	AcceptFoa         apijson.Field
	ApproveTransfer   apijson.Field
	CanCancelTransfer apijson.Field
	DisablePrivacy    apijson.Field
	EnterAuthCode     apijson.Field
	UnlockDomain      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RegistrarDomainListResponseTransferIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [RegistrarDomainGetResponseUnknown],
// [RegistrarDomainGetResponseArray] or [shared.UnionString].
type RegistrarDomainGetResponse interface {
	ImplementsRegistrarDomainGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RegistrarDomainGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RegistrarDomainGetResponseArray []interface{}

func (r RegistrarDomainGetResponseArray) ImplementsRegistrarDomainGetResponse() {}

// Union satisfied by [RegistrarDomainReplaceResponseUnknown],
// [RegistrarDomainReplaceResponseArray] or [shared.UnionString].
type RegistrarDomainReplaceResponse interface {
	ImplementsRegistrarDomainReplaceResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RegistrarDomainReplaceResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RegistrarDomainReplaceResponseArray []interface{}

func (r RegistrarDomainReplaceResponseArray) ImplementsRegistrarDomainReplaceResponse() {}

type RegistrarDomainListResponseEnvelope struct {
	Errors   []RegistrarDomainListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RegistrarDomainListResponseEnvelopeMessages `json:"messages,required"`
	Result   []RegistrarDomainListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RegistrarDomainListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RegistrarDomainListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       registrarDomainListResponseEnvelopeJSON       `json:"-"`
}

// registrarDomainListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegistrarDomainListResponseEnvelope]
type registrarDomainListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarDomainListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RegistrarDomainListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    registrarDomainListResponseEnvelopeErrorsJSON `json:"-"`
}

// registrarDomainListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RegistrarDomainListResponseEnvelopeErrors]
type registrarDomainListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarDomainListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RegistrarDomainListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    registrarDomainListResponseEnvelopeMessagesJSON `json:"-"`
}

// registrarDomainListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RegistrarDomainListResponseEnvelopeMessages]
type registrarDomainListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarDomainListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RegistrarDomainListResponseEnvelopeSuccess bool

const (
	RegistrarDomainListResponseEnvelopeSuccessTrue RegistrarDomainListResponseEnvelopeSuccess = true
)

type RegistrarDomainListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       registrarDomainListResponseEnvelopeResultInfoJSON `json:"-"`
}

// registrarDomainListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [RegistrarDomainListResponseEnvelopeResultInfo]
type registrarDomainListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarDomainListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RegistrarDomainGetResponseEnvelope struct {
	Errors   []RegistrarDomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RegistrarDomainGetResponseEnvelopeMessages `json:"messages,required"`
	Result   RegistrarDomainGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RegistrarDomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    registrarDomainGetResponseEnvelopeJSON    `json:"-"`
}

// registrarDomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RegistrarDomainGetResponseEnvelope]
type registrarDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RegistrarDomainGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    registrarDomainGetResponseEnvelopeErrorsJSON `json:"-"`
}

// registrarDomainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RegistrarDomainGetResponseEnvelopeErrors]
type registrarDomainGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarDomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RegistrarDomainGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    registrarDomainGetResponseEnvelopeMessagesJSON `json:"-"`
}

// registrarDomainGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RegistrarDomainGetResponseEnvelopeMessages]
type registrarDomainGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarDomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RegistrarDomainGetResponseEnvelopeSuccess bool

const (
	RegistrarDomainGetResponseEnvelopeSuccessTrue RegistrarDomainGetResponseEnvelopeSuccess = true
)

type RegistrarDomainReplaceParams struct {
	// Auto-renew controls whether subscription is automatically renewed upon domain
	// expiration.
	AutoRenew param.Field[bool] `json:"auto_renew"`
	// Shows whether a registrar lock is in place for a domain.
	Locked param.Field[bool] `json:"locked"`
	// Privacy option controls redacting WHOIS information.
	Privacy param.Field[bool] `json:"privacy"`
}

func (r RegistrarDomainReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegistrarDomainReplaceResponseEnvelope struct {
	Errors   []RegistrarDomainReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RegistrarDomainReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   RegistrarDomainReplaceResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RegistrarDomainReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    registrarDomainReplaceResponseEnvelopeJSON    `json:"-"`
}

// registrarDomainReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegistrarDomainReplaceResponseEnvelope]
type registrarDomainReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarDomainReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RegistrarDomainReplaceResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    registrarDomainReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// registrarDomainReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RegistrarDomainReplaceResponseEnvelopeErrors]
type registrarDomainReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarDomainReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RegistrarDomainReplaceResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    registrarDomainReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// registrarDomainReplaceResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RegistrarDomainReplaceResponseEnvelopeMessages]
type registrarDomainReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarDomainReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RegistrarDomainReplaceResponseEnvelopeSuccess bool

const (
	RegistrarDomainReplaceResponseEnvelopeSuccessTrue RegistrarDomainReplaceResponseEnvelopeSuccess = true
)
