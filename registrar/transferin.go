// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// TransferInService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransferInService] method instead.
type TransferInService struct {
	Options []option.RequestOption
}

// NewTransferInService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransferInService(opts ...option.RequestOption) (r *TransferInService) {
	r = &TransferInService{}
	r.Options = opts
	return
}

// Starts a domain transfer-in workflow. This is typically a billable operation —
// successful transfers charge the account's default payment method, except for
// extensions with zero transfer pricing (e.g. UK extensions). All successful
// domain transfers are non-refundable.
//
// ### How transfers work
//
// Domain transfers move a domain from another registrar to Cloudflare. Transfers
// typically take 1-10 days due to ICANN-mandated approval windows.
//
// ### Prerequisites
//
//   - The domain must already have a zone in the Cloudflare account (added through
//     the dashboard or zone API).
//   - The zone must have DNSSec disabled.
//   - For billable transfers (i.e. extensions with non-zero transfer pricing), the
//     account must have a billing profile with a valid default payment method. Set
//     this up at `https://dash.cloudflare.com/{account_id}/billing/payment-info`.
//   - The domain must be unlocked at the current registrar.
//   - An authorization/EPP code from the current registrar is required, except for
//     UK extensions — see Auth code below.
//
// ### Auth code
//
// An authorization code (also called EPP code, transfer key, or auth-info code) is
// required for most extensions, with the exception of UK extensions. Obtain this
// from your current registrar's control panel.
//
// The auth code in the request body must be base64-encoded per RFC 4648 §4
// (standard alphabet, no line breaks).
//
// ### Response behavior
//
// Successful transfer initiation returns `202 Accepted`. Validation or initiation
// failures return the documented `4XX` responses. Poll
// `GET /accounts/{account_id}/registrar/registrations/{domain_name}/transfer-in-status`
// to track progress.
//
// ### Premium domains
//
// Premium domain transfers are not currently supported by this API. Please use the
// [dashboard](https://dash.cloudflare.com/) for now.
//
// ### Billing
//
// The account's default payment method is charged upon successful transfer
// completion, unless the extension has zero transfer pricing (e.g. UK extensions).
// The transfer adds time to the domain's existing expiration date (typically 1
// year).
func (r *TransferInService) New(ctx context.Context, domainName string, params TransferInNewParams, opts ...option.RequestOption) (res *WorkflowStatus, err error) {
	var env TransferInNewResponseEnvelope
	if params.Prefer.Present {
		opts = append(opts, option.WithHeader("Prefer", fmt.Sprintf("%v", params.Prefer)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar/registrations/%s/transfer-in", params.AccountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type TransferInNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The EPP/authorization code from your current registrar, base64-encoded per RFC
	// 4648 §4. Obtain this from your current registrar's control panel. Required for
	// all extensions, except for UK.
	AuthCode param.Field[string] `json:"auth_code" format:"byte"`
	// Enable or disable automatic renewal after transfer. Defaults to `false` if
	// omitted.
	AutoRenew param.Field[bool] `json:"auto_renew"`
	// Registry-specific contact extension values for the registrant.
	// `GET /accounts/{account_id}/registrar/extensions/{extension}` documents the
	// required keys and allowed values for each extension in the
	// `transfer_schema.properties.contact_extensions` object.
	//
	// Examples include `.us` nexus fields, `.uk` registrant type fields, and `.ca`
	// legal type fields. Include this object only when the extension's transfer schema
	// defines `contact_extensions`.
	ContactExtensions param.Field[map[string]interface{}] `json:"contact_extensions"`
	// Provides contact data for the registration request.
	//
	// The per-extension schema from
	// `GET /accounts/{account_id}/registrar/extensions/{extension}` defines the
	// accepted contact roles. Every currently supported extension requires only
	// `contacts.registrant` from API callers. Callers may provide additional roles
	// such as `technical`, `administrator`, and `billing` when the extension schema
	// includes them. When a registry requires an omitted role, Cloudflare may derive
	// that contact from `contacts.registrant`.
	//
	// When the request omits either the entire `contacts` object or
	// `contacts.registrant`, the system uses the account's default address book entry
	// as the registrant contact. The account owner must configure this default at
	// `https://dash.cloudflare.com/{account_id}/domains/registrations`, where they can
	// create or update the address book entry and accept the required agreement.
	// Dashboard settings currently provide the only way to manage address book
	// entries.
	//
	// Without either a default address book entry or a registrant contact, the
	// registration request fails validation.
	Contacts param.Field[TransferInNewParamsContacts] `json:"contacts"`
	// WHOIS privacy mode to apply after transfer completes. Defaults to the
	// extension's default privacy mode (typically `redaction`).
	PrivacyMode param.Field[TransferInNewParamsPrivacyMode] `json:"privacy_mode"`
	Prefer      param.Field[string]                         `header:"Prefer"`
}

func (r TransferInNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Provides contact data for the registration request.
//
// The per-extension schema from
// `GET /accounts/{account_id}/registrar/extensions/{extension}` defines the
// accepted contact roles. Every currently supported extension requires only
// `contacts.registrant` from API callers. Callers may provide additional roles
// such as `technical`, `administrator`, and `billing` when the extension schema
// includes them. When a registry requires an omitted role, Cloudflare may derive
// that contact from `contacts.registrant`.
//
// When the request omits either the entire `contacts` object or
// `contacts.registrant`, the system uses the account's default address book entry
// as the registrant contact. The account owner must configure this default at
// `https://dash.cloudflare.com/{account_id}/domains/registrations`, where they can
// create or update the address book entry and accept the required agreement.
// Dashboard settings currently provide the only way to manage address book
// entries.
//
// Without either a default address book entry or a registrant contact, the
// registration request fails validation.
type TransferInNewParamsContacts struct {
	// Optional administrator contact. Accepted only when the extension schema includes
	// this role. When the registry requires an omitted contact, Cloudflare may derive
	// it from `contacts.registrant`.
	Administrator param.Field[TransferInNewParamsContactsAdministrator] `json:"administrator"`
	// Optional billing contact. Accepted only when the extension schema includes this
	// role. When the registry requires an omitted contact, Cloudflare may derive it
	// from `contacts.registrant`.
	Billing param.Field[TransferInNewParamsContactsBilling] `json:"billing"`
	// Optional registrant contact. If omitted, the account's default address book
	// entry is used instead.
	Registrant param.Field[TransferInNewParamsContactsRegistrant] `json:"registrant"`
	// Optional technical contact. Accepted only when the extension schema includes
	// this role. When the registry requires an omitted contact, Cloudflare may derive
	// it from `contacts.registrant`.
	Technical param.Field[TransferInNewParamsContactsTechnical] `json:"technical"`
}

func (r TransferInNewParamsContacts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional administrator contact. Accepted only when the extension schema includes
// this role. When the registry requires an omitted contact, Cloudflare may derive
// it from `contacts.registrant`.
type TransferInNewParamsContactsAdministrator struct {
	// Email address for the registrant. Used for domain-related communications from
	// the registry, including ownership verification and renewal notices.
	Email param.Field[string] `json:"email" api:"required" format:"email"`
	// Phone number in E.164 format: `+{country_code}.{number}` without spaces or
	// dashes. Examples: `+1.5555555555` (US), `+44.2071234567` (UK), `+81.312345678`
	// (Japan).
	Phone param.Field[string] `json:"phone" api:"required"`
	// Postal/mailing information for the contact. The `name` field is the complete
	// contact name in one string. Some registries require a complete personal name,
	// including a family or last name where applicable, but this API does not accept
	// separate first-name and last-name fields for registration contacts.
	PostalInfo param.Field[TransferInNewParamsContactsAdministratorPostalInfo] `json:"postal_info" api:"required"`
	// Fax number in E.164 format (e.g., `+1.5555555555`). Optional. Most registrations
	// do not require a fax number.
	Fax param.Field[string] `json:"fax"`
}

func (r TransferInNewParamsContactsAdministrator) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Postal/mailing information for the contact. The `name` field is the complete
// contact name in one string. Some registries require a complete personal name,
// including a family or last name where applicable, but this API does not accept
// separate first-name and last-name fields for registration contacts.
type TransferInNewParamsContactsAdministratorPostalInfo struct {
	// Physical mailing address for the registrant contact.
	Address param.Field[TransferInNewParamsContactsAdministratorPostalInfoAddress] `json:"address" api:"required"`
	// Full legal name of the contact, including all required name components for an
	// individual or authorized representative. Some registries require a complete
	// personal name that includes a family or last name where applicable. Provide the
	// complete name in this single field, for example `Ada Lovelace`; do not send
	// separate first-name or last-name fields.
	Name param.Field[string] `json:"name" api:"required"`
	// Organization or company name. Optional for individual registrants.
	Organization param.Field[string] `json:"organization"`
}

func (r TransferInNewParamsContactsAdministratorPostalInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Physical mailing address for the registrant contact.
type TransferInNewParamsContactsAdministratorPostalInfoAddress struct {
	// City or locality name.
	City param.Field[string] `json:"city" api:"required"`
	// Two-letter country code per ISO 3166-1 alpha-2 (e.g., `US`, `GB`, `CA`, `DE`).
	CountryCode param.Field[string] `json:"country_code" api:"required"`
	// Postal or ZIP code.
	PostalCode param.Field[string] `json:"postal_code" api:"required"`
	// State, province, or region. Use the standard abbreviation where applicable
	// (e.g., `TX` for Texas, `ON` for Ontario).
	State param.Field[string] `json:"state" api:"required"`
	// Street address including building/suite number.
	Street param.Field[string] `json:"street" api:"required"`
}

func (r TransferInNewParamsContactsAdministratorPostalInfoAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional billing contact. Accepted only when the extension schema includes this
// role. When the registry requires an omitted contact, Cloudflare may derive it
// from `contacts.registrant`.
type TransferInNewParamsContactsBilling struct {
	// Email address for the registrant. Used for domain-related communications from
	// the registry, including ownership verification and renewal notices.
	Email param.Field[string] `json:"email" api:"required" format:"email"`
	// Phone number in E.164 format: `+{country_code}.{number}` without spaces or
	// dashes. Examples: `+1.5555555555` (US), `+44.2071234567` (UK), `+81.312345678`
	// (Japan).
	Phone param.Field[string] `json:"phone" api:"required"`
	// Postal/mailing information for the contact. The `name` field is the complete
	// contact name in one string. Some registries require a complete personal name,
	// including a family or last name where applicable, but this API does not accept
	// separate first-name and last-name fields for registration contacts.
	PostalInfo param.Field[TransferInNewParamsContactsBillingPostalInfo] `json:"postal_info" api:"required"`
	// Fax number in E.164 format (e.g., `+1.5555555555`). Optional. Most registrations
	// do not require a fax number.
	Fax param.Field[string] `json:"fax"`
}

func (r TransferInNewParamsContactsBilling) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Postal/mailing information for the contact. The `name` field is the complete
// contact name in one string. Some registries require a complete personal name,
// including a family or last name where applicable, but this API does not accept
// separate first-name and last-name fields for registration contacts.
type TransferInNewParamsContactsBillingPostalInfo struct {
	// Physical mailing address for the registrant contact.
	Address param.Field[TransferInNewParamsContactsBillingPostalInfoAddress] `json:"address" api:"required"`
	// Full legal name of the contact, including all required name components for an
	// individual or authorized representative. Some registries require a complete
	// personal name that includes a family or last name where applicable. Provide the
	// complete name in this single field, for example `Ada Lovelace`; do not send
	// separate first-name or last-name fields.
	Name param.Field[string] `json:"name" api:"required"`
	// Organization or company name. Optional for individual registrants.
	Organization param.Field[string] `json:"organization"`
}

func (r TransferInNewParamsContactsBillingPostalInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Physical mailing address for the registrant contact.
type TransferInNewParamsContactsBillingPostalInfoAddress struct {
	// City or locality name.
	City param.Field[string] `json:"city" api:"required"`
	// Two-letter country code per ISO 3166-1 alpha-2 (e.g., `US`, `GB`, `CA`, `DE`).
	CountryCode param.Field[string] `json:"country_code" api:"required"`
	// Postal or ZIP code.
	PostalCode param.Field[string] `json:"postal_code" api:"required"`
	// State, province, or region. Use the standard abbreviation where applicable
	// (e.g., `TX` for Texas, `ON` for Ontario).
	State param.Field[string] `json:"state" api:"required"`
	// Street address including building/suite number.
	Street param.Field[string] `json:"street" api:"required"`
}

func (r TransferInNewParamsContactsBillingPostalInfoAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional registrant contact. If omitted, the account's default address book
// entry is used instead.
type TransferInNewParamsContactsRegistrant struct {
	// Email address for the registrant. Used for domain-related communications from
	// the registry, including ownership verification and renewal notices.
	Email param.Field[string] `json:"email" api:"required" format:"email"`
	// Phone number in E.164 format: `+{country_code}.{number}` without spaces or
	// dashes. Examples: `+1.5555555555` (US), `+44.2071234567` (UK), `+81.312345678`
	// (Japan).
	Phone param.Field[string] `json:"phone" api:"required"`
	// Postal/mailing information for the contact. The `name` field is the complete
	// contact name in one string. Some registries require a complete personal name,
	// including a family or last name where applicable, but this API does not accept
	// separate first-name and last-name fields for registration contacts.
	PostalInfo param.Field[TransferInNewParamsContactsRegistrantPostalInfo] `json:"postal_info" api:"required"`
	// Fax number in E.164 format (e.g., `+1.5555555555`). Optional. Most registrations
	// do not require a fax number.
	Fax param.Field[string] `json:"fax"`
}

func (r TransferInNewParamsContactsRegistrant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Postal/mailing information for the contact. The `name` field is the complete
// contact name in one string. Some registries require a complete personal name,
// including a family or last name where applicable, but this API does not accept
// separate first-name and last-name fields for registration contacts.
type TransferInNewParamsContactsRegistrantPostalInfo struct {
	// Physical mailing address for the registrant contact.
	Address param.Field[TransferInNewParamsContactsRegistrantPostalInfoAddress] `json:"address" api:"required"`
	// Full legal name of the contact, including all required name components for an
	// individual or authorized representative. Some registries require a complete
	// personal name that includes a family or last name where applicable. Provide the
	// complete name in this single field, for example `Ada Lovelace`; do not send
	// separate first-name or last-name fields.
	Name param.Field[string] `json:"name" api:"required"`
	// Organization or company name. Optional for individual registrants.
	Organization param.Field[string] `json:"organization"`
}

func (r TransferInNewParamsContactsRegistrantPostalInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Physical mailing address for the registrant contact.
type TransferInNewParamsContactsRegistrantPostalInfoAddress struct {
	// City or locality name.
	City param.Field[string] `json:"city" api:"required"`
	// Two-letter country code per ISO 3166-1 alpha-2 (e.g., `US`, `GB`, `CA`, `DE`).
	CountryCode param.Field[string] `json:"country_code" api:"required"`
	// Postal or ZIP code.
	PostalCode param.Field[string] `json:"postal_code" api:"required"`
	// State, province, or region. Use the standard abbreviation where applicable
	// (e.g., `TX` for Texas, `ON` for Ontario).
	State param.Field[string] `json:"state" api:"required"`
	// Street address including building/suite number.
	Street param.Field[string] `json:"street" api:"required"`
}

func (r TransferInNewParamsContactsRegistrantPostalInfoAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional technical contact. Accepted only when the extension schema includes
// this role. When the registry requires an omitted contact, Cloudflare may derive
// it from `contacts.registrant`.
type TransferInNewParamsContactsTechnical struct {
	// Email address for the registrant. Used for domain-related communications from
	// the registry, including ownership verification and renewal notices.
	Email param.Field[string] `json:"email" api:"required" format:"email"`
	// Phone number in E.164 format: `+{country_code}.{number}` without spaces or
	// dashes. Examples: `+1.5555555555` (US), `+44.2071234567` (UK), `+81.312345678`
	// (Japan).
	Phone param.Field[string] `json:"phone" api:"required"`
	// Postal/mailing information for the contact. The `name` field is the complete
	// contact name in one string. Some registries require a complete personal name,
	// including a family or last name where applicable, but this API does not accept
	// separate first-name and last-name fields for registration contacts.
	PostalInfo param.Field[TransferInNewParamsContactsTechnicalPostalInfo] `json:"postal_info" api:"required"`
	// Fax number in E.164 format (e.g., `+1.5555555555`). Optional. Most registrations
	// do not require a fax number.
	Fax param.Field[string] `json:"fax"`
}

func (r TransferInNewParamsContactsTechnical) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Postal/mailing information for the contact. The `name` field is the complete
// contact name in one string. Some registries require a complete personal name,
// including a family or last name where applicable, but this API does not accept
// separate first-name and last-name fields for registration contacts.
type TransferInNewParamsContactsTechnicalPostalInfo struct {
	// Physical mailing address for the registrant contact.
	Address param.Field[TransferInNewParamsContactsTechnicalPostalInfoAddress] `json:"address" api:"required"`
	// Full legal name of the contact, including all required name components for an
	// individual or authorized representative. Some registries require a complete
	// personal name that includes a family or last name where applicable. Provide the
	// complete name in this single field, for example `Ada Lovelace`; do not send
	// separate first-name or last-name fields.
	Name param.Field[string] `json:"name" api:"required"`
	// Organization or company name. Optional for individual registrants.
	Organization param.Field[string] `json:"organization"`
}

func (r TransferInNewParamsContactsTechnicalPostalInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Physical mailing address for the registrant contact.
type TransferInNewParamsContactsTechnicalPostalInfoAddress struct {
	// City or locality name.
	City param.Field[string] `json:"city" api:"required"`
	// Two-letter country code per ISO 3166-1 alpha-2 (e.g., `US`, `GB`, `CA`, `DE`).
	CountryCode param.Field[string] `json:"country_code" api:"required"`
	// Postal or ZIP code.
	PostalCode param.Field[string] `json:"postal_code" api:"required"`
	// State, province, or region. Use the standard abbreviation where applicable
	// (e.g., `TX` for Texas, `ON` for Ontario).
	State param.Field[string] `json:"state" api:"required"`
	// Street address including building/suite number.
	Street param.Field[string] `json:"street" api:"required"`
}

func (r TransferInNewParamsContactsTechnicalPostalInfoAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WHOIS privacy mode to apply after transfer completes. Defaults to the
// extension's default privacy mode (typically `redaction`).
type TransferInNewParamsPrivacyMode string

const (
	TransferInNewParamsPrivacyModeOff       TransferInNewParamsPrivacyMode = "off"
	TransferInNewParamsPrivacyModeRedaction TransferInNewParamsPrivacyMode = "redaction"
)

func (r TransferInNewParamsPrivacyMode) IsKnown() bool {
	switch r {
	case TransferInNewParamsPrivacyModeOff, TransferInNewParamsPrivacyModeRedaction:
		return true
	}
	return false
}

type TransferInNewResponseEnvelope struct {
	Errors   []TransferInNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []TransferInNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Status of an async registration workflow.
	Result WorkflowStatus `json:"result" api:"required"`
	// Whether the API call was successful.
	Success TransferInNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    transferInNewResponseEnvelopeJSON    `json:"-"`
}

// transferInNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TransferInNewResponseEnvelope]
type transferInNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferInNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferInNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TransferInNewResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source TransferInNewResponseEnvelopeErrorsSource `json:"source"`
	JSON   transferInNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// transferInNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TransferInNewResponseEnvelopeErrors]
type transferInNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferInNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferInNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type TransferInNewResponseEnvelopeErrorsSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                        `json:"pointer" api:"required"`
	JSON    transferInNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// transferInNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [TransferInNewResponseEnvelopeErrorsSource]
type transferInNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferInNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferInNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TransferInNewResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source TransferInNewResponseEnvelopeMessagesSource `json:"source"`
	JSON   transferInNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// transferInNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TransferInNewResponseEnvelopeMessages]
type transferInNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferInNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferInNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type TransferInNewResponseEnvelopeMessagesSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    transferInNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// transferInNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [TransferInNewResponseEnvelopeMessagesSource]
type transferInNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferInNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferInNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TransferInNewResponseEnvelopeSuccess bool

const (
	TransferInNewResponseEnvelopeSuccessTrue TransferInNewResponseEnvelopeSuccess = true
)

func (r TransferInNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TransferInNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
