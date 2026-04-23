// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// RegistrationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistrationService] method instead.
type RegistrationService struct {
	Options []option.RequestOption
}

// NewRegistrationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegistrationService(opts ...option.RequestOption) (r *RegistrationService) {
	r = &RegistrationService{}
	r.Options = opts
	return
}

// Starts a domain registration workflow. This is a billable operation — successful
// registration charges the account's default payment method. All successful domain
// registrations are non-refundable — once the workflow completes with
// `state: succeeded`, the charge cannot be reversed.
//
// ### Prerequisites
//
//   - The account must have a billing profile with a valid default payment method.
//     Set this up at
//     `https://dash.cloudflare.com/{account_id}/billing/payment-info`.
//   - The account must not already be at the maximum supported domain limit. A
//     single account may own up to 100 domains in total across registrations created
//     through either the dashboard or this API.
//   - The domain must be on a supported extension for programmatic registration.
//   - Use `POST /domain-check` immediately before calling this endpoint to confirm
//     real-time availability and pricing.
//
// ### Supported extensions
//
// In this API, "extension" means the full registrable suffix after the domain
// label. For example, in `example.co.uk`, the extension is `co.uk`.
//
// Programmatic registration is currently supported for:
//
// `com`, `org`, `net`, `app`, `dev`, `cc`, `xyz`, `info`, `cloud`, `studio`,
// `live`, `link`, `pro`, `tech`, `fyi`, `shop`, `online`, `tools`, `run`, `games`,
// `build`, `systems`, `world`, `news`, `site`, `network`, `chat`, `space`,
// `family`, `page`, `life`, `group`, `email`, `solutions`, `day`, `blog`, `ing`,
// `icu`, `academy`, `today`
//
// Cloudflare Registrar supports 400+ extensions in the dashboard. Extensions not
// listed above can still be registered at
// `https://dash.cloudflare.com/{account_id}/domains/registrations`.
//
// ### Express mode
//
// The only required field is `domain_name`. If `contacts` is omitted, the system
// uses the account's default address book entry as the registrant. If no default
// exists and no contact is provided, the request fails. Set up a default address
// book entry and accept the required agreement at
// `https://dash.cloudflare.com/{account_id}/domains/registrations`.
//
// ### Defaults
//
//   - `years`: defaults to the extension's minimum registration period (1 year for
//     most extensions, but varies — for example, `.ai` (if supported) requires a
//     minimum of 2 years).
//   - `auto_renew`: defaults to `false`. Setting it to `true` is an explicit opt-in
//     authorizing Cloudflare to charge the account's default payment method up to 30
//     days before domain expiry to renew the registration. Renewal pricing may
//     change over time based on registry pricing.
//   - `privacy_mode`: defaults to `redaction`.
//
// ### Premium domains
//
// Premium domain registration is not currently supported by this API. If
// `POST /domain-check` returns `tier: premium`, do not call this endpoint for that
// domain.
//
// ### Response behavior
//
// By default, the server holds the connection for a bounded, server-defined amount
// of time while the registration completes. Most registrations finish within this
// window and return `201 Created` with a completed workflow status.
//
// If the registration is still processing after this synchronous wait window, the
// server returns `202 Accepted`. Poll the URL in `links.self` to track progress.
//
// To skip the wait and receive an immediate `202`, send `Prefer: respond-async`.
func (r *RegistrationService) New(ctx context.Context, params RegistrationNewParams, opts ...option.RequestOption) (res *WorkflowStatus, err error) {
	var env RegistrationNewResponseEnvelope
	if params.Prefer.Present {
		opts = append(opts, option.WithHeader("Prefer", fmt.Sprintf("%v", params.Prefer)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar/registrations", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of domain registrations owned by the account.
//
// This endpoint uses cursor-based pagination. Results are ordered by registration
// date by default. To fetch the next page, pass the `cursor` value from the
// `result_info` object in the response as the `cursor` query parameter in your
// next request. An empty `cursor` string indicates there are no more pages.
func (r *RegistrationService) List(ctx context.Context, params RegistrationListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Registration], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar/registrations", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a paginated list of domain registrations owned by the account.
//
// This endpoint uses cursor-based pagination. Results are ordered by registration
// date by default. To fetch the next page, pass the `cursor` value from the
// `result_info` object in the response as the `cursor` query parameter in your
// next request. An empty `cursor` string indicates there are no more pages.
func (r *RegistrationService) ListAutoPaging(ctx context.Context, params RegistrationListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Registration] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Updates an existing domain registration.
//
// By default, the server holds the connection for a bounded, server-defined amount
// of time while the update completes. Most updates finish within this window and
// return `200 OK` with a completed workflow status.
//
// If the update is still processing after this synchronous wait window, the server
// returns `202 Accepted`. Poll the URL in `links.self` to track progress.
//
// To skip the wait and receive an immediate `202`, send `Prefer: respond-async`.
//
// This endpoint currently supports updating `auto_renew` only.
func (r *RegistrationService) Edit(ctx context.Context, domainName string, params RegistrationEditParams, opts ...option.RequestOption) (res *WorkflowStatus, err error) {
	var env RegistrationEditResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/registrar/registrations/%s", params.AccountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the current state of a domain registration.
//
// This is the canonical read endpoint for a domain you own. It returns the full
// registration resource including current settings and expiration. When the
// registration resource is ready, both `created_at` and `expires_at` are present
// in the response.
func (r *RegistrationService) Get(ctx context.Context, domainName string, query RegistrationGetParams, opts ...option.RequestOption) (res *Registration, err error) {
	var env RegistrationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar/registrations/%s", query.AccountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type RegistrationNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Fully qualified domain name (FQDN) including the extension (e.g., `example.com`,
	// `mybrand.app`). The domain name uniquely identifies a registration — the same
	// domain cannot be registered twice, making it a natural idempotency key for
	// registration requests.
	DomainName param.Field[string] `json:"domain_name" api:"required"`
	// Enable or disable automatic renewal. Defaults to `false` if omitted. Setting
	// this field to `true` is an explicit opt-in authorizing Cloudflare to charge the
	// account's default payment method up to 30 days before domain expiry to renew the
	// domain automatically. Renewal pricing may change over time based on registry
	// pricing.
	AutoRenew param.Field[bool] `json:"auto_renew"`
	// Contact data for the registration request.
	//
	// If the `contacts` object is omitted entirely from the request, or if
	// `contacts.registrant` is not provided, the system will use the account's default
	// address book entry as the registrant contact. This default must be
	// pre-configured by the account owner at
	// `https://dash.cloudflare.com/{account_id}/domains/registrations`, where they can
	// create or update the address book entry and accept the required agreement. No
	// API exists for managing address book entries at this time.
	//
	// If no default address book entry exists and no registrant contact is provided,
	// the registration request will fail with a validation error.
	Contacts param.Field[RegistrationNewParamsContacts] `json:"contacts"`
	// WHOIS privacy mode for the registration. Defaults to `redaction`.
	//
	//   - `off`: Do not request WHOIS privacy.
	//   - `redaction`: Request WHOIS redaction where supported by the extension. Some
	//     extensions do not support privacy/redaction.
	PrivacyMode param.Field[RegistrationNewParamsPrivacyMode] `json:"privacy_mode"`
	// Number of years to register (1–10). If omitted, defaults to the minimum
	// registration period required by the registry for this extension. For most
	// extensions this is 1 year, but some extensions require longer minimum terms
	// (e.g., `.ai` requires a minimum of 2 years).
	//
	// The registry for each extension may also enforce its own maximum registration
	// term. If the requested value exceeds the registry's maximum, the registration
	// will be rejected. When in doubt, use the default by omitting this field.
	Years  param.Field[int64]  `json:"years"`
	Prefer param.Field[string] `header:"Prefer"`
}

func (r RegistrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Contact data for the registration request.
//
// If the `contacts` object is omitted entirely from the request, or if
// `contacts.registrant` is not provided, the system will use the account's default
// address book entry as the registrant contact. This default must be
// pre-configured by the account owner at
// `https://dash.cloudflare.com/{account_id}/domains/registrations`, where they can
// create or update the address book entry and accept the required agreement. No
// API exists for managing address book entries at this time.
//
// If no default address book entry exists and no registrant contact is provided,
// the registration request will fail with a validation error.
type RegistrationNewParamsContacts struct {
	// Registrant contact data for the domain registration. This information is
	// submitted to the domain registry and, depending on extension and privacy
	// settings, may appear in public WHOIS records.
	Registrant param.Field[RegistrationNewParamsContactsRegistrant] `json:"registrant"`
}

func (r RegistrationNewParamsContacts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Registrant contact data for the domain registration. This information is
// submitted to the domain registry and, depending on extension and privacy
// settings, may appear in public WHOIS records.
type RegistrationNewParamsContactsRegistrant struct {
	// Email address for the registrant. Used for domain-related communications from
	// the registry, including ownership verification and renewal notices.
	Email param.Field[string] `json:"email" api:"required" format:"email"`
	// Phone number in E.164 format: `+{country_code}.{number}` with no spaces or
	// dashes. Examples: `+1.5555555555` (US), `+44.2071234567` (UK), `+81.312345678`
	// (Japan).
	Phone param.Field[string] `json:"phone" api:"required"`
	// Postal/mailing information for the registrant contact.
	PostalInfo param.Field[RegistrationNewParamsContactsRegistrantPostalInfo] `json:"postal_info" api:"required"`
	// Fax number in E.164 format (e.g., `+1.5555555555`). Optional. Most registrations
	// do not require a fax number.
	Fax param.Field[string] `json:"fax"`
}

func (r RegistrationNewParamsContactsRegistrant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Postal/mailing information for the registrant contact.
type RegistrationNewParamsContactsRegistrantPostalInfo struct {
	// Physical mailing address for the registrant contact.
	Address param.Field[RegistrationNewParamsContactsRegistrantPostalInfoAddress] `json:"address" api:"required"`
	// Full legal name of the registrant (individual or authorized representative).
	Name param.Field[string] `json:"name" api:"required"`
	// Organization or company name. Optional for individual registrants.
	Organization param.Field[string] `json:"organization"`
}

func (r RegistrationNewParamsContactsRegistrantPostalInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Physical mailing address for the registrant contact.
type RegistrationNewParamsContactsRegistrantPostalInfoAddress struct {
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

func (r RegistrationNewParamsContactsRegistrantPostalInfoAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WHOIS privacy mode for the registration. Defaults to `redaction`.
//
//   - `off`: Do not request WHOIS privacy.
//   - `redaction`: Request WHOIS redaction where supported by the extension. Some
//     extensions do not support privacy/redaction.
type RegistrationNewParamsPrivacyMode string

const (
	RegistrationNewParamsPrivacyModeRedaction RegistrationNewParamsPrivacyMode = "redaction"
)

func (r RegistrationNewParamsPrivacyMode) IsKnown() bool {
	switch r {
	case RegistrationNewParamsPrivacyModeRedaction:
		return true
	}
	return false
}

type RegistrationNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Status of an async registration workflow.
	Result WorkflowStatus `json:"result" api:"required"`
	// Whether the API call was successful
	Success RegistrationNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    registrationNewResponseEnvelopeJSON    `json:"-"`
}

// registrationNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RegistrationNewResponseEnvelope]
type registrationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegistrationNewResponseEnvelopeSuccess bool

const (
	RegistrationNewResponseEnvelopeSuccessTrue RegistrationNewResponseEnvelopeSuccess = true
)

func (r RegistrationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegistrationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RegistrationListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Opaque token from a previous response's `result_info.cursor`. Pass this value to
	// fetch the next page of results. Omit (or pass an empty string) for the first
	// page.
	Cursor param.Field[string] `query:"cursor"`
	// Sort direction for results. Defaults to ascending order.
	Direction param.Field[RegistrationListParamsDirection] `query:"direction"`
	// Number of items to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Column to sort results by. Defaults to registration date (`registry_created_at`)
	// when omitted.
	SortBy param.Field[RegistrationListParamsSortBy] `query:"sort_by"`
}

// URLQuery serializes [RegistrationListParams]'s query parameters as `url.Values`.
func (r RegistrationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction for results. Defaults to ascending order.
type RegistrationListParamsDirection string

const (
	RegistrationListParamsDirectionAsc  RegistrationListParamsDirection = "asc"
	RegistrationListParamsDirectionDesc RegistrationListParamsDirection = "desc"
)

func (r RegistrationListParamsDirection) IsKnown() bool {
	switch r {
	case RegistrationListParamsDirectionAsc, RegistrationListParamsDirectionDesc:
		return true
	}
	return false
}

// Column to sort results by. Defaults to registration date (`registry_created_at`)
// when omitted.
type RegistrationListParamsSortBy string

const (
	RegistrationListParamsSortByRegistryCreatedAt RegistrationListParamsSortBy = "registry_created_at"
	RegistrationListParamsSortByRegistryExpiresAt RegistrationListParamsSortBy = "registry_expires_at"
	RegistrationListParamsSortByName              RegistrationListParamsSortBy = "name"
)

func (r RegistrationListParamsSortBy) IsKnown() bool {
	switch r {
	case RegistrationListParamsSortByRegistryCreatedAt, RegistrationListParamsSortByRegistryExpiresAt, RegistrationListParamsSortByName:
		return true
	}
	return false
}

type RegistrationEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Enable or disable automatic renewal. Setting this field to `true` authorizes
	// Cloudflare to charge the account's default payment method up to 30 days before
	// domain expiry to renew the domain automatically. Renewal pricing may change over
	// time based on registry pricing.
	AutoRenew param.Field[bool]                         `json:"auto_renew"`
	Prefer    param.Field[RegistrationEditParamsPrefer] `header:"Prefer"`
}

func (r RegistrationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegistrationEditParamsPrefer string

const (
	RegistrationEditParamsPreferRespondAsync RegistrationEditParamsPrefer = "respond-async"
)

func (r RegistrationEditParamsPrefer) IsKnown() bool {
	switch r {
	case RegistrationEditParamsPreferRespondAsync:
		return true
	}
	return false
}

type RegistrationEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Status of an async registration workflow.
	Result WorkflowStatus `json:"result" api:"required"`
	// Whether the API call was successful
	Success RegistrationEditResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    registrationEditResponseEnvelopeJSON    `json:"-"`
}

// registrationEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RegistrationEditResponseEnvelope]
type registrationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegistrationEditResponseEnvelopeSuccess bool

const (
	RegistrationEditResponseEnvelopeSuccessTrue RegistrationEditResponseEnvelopeSuccess = true
)

func (r RegistrationEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegistrationEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RegistrationGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RegistrationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// A domain registration resource representing the current state of a registered
	// domain.
	Result Registration `json:"result" api:"required"`
	// Whether the API call was successful
	Success RegistrationGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    registrationGetResponseEnvelopeJSON    `json:"-"`
}

// registrationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RegistrationGetResponseEnvelope]
type registrationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegistrationGetResponseEnvelopeSuccess bool

const (
	RegistrationGetResponseEnvelopeSuccessTrue RegistrationGetResponseEnvelopeSuccess = true
)

func (r RegistrationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegistrationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
