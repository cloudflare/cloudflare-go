// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// RegistrarService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistrarService] method instead.
type RegistrarService struct {
	Options            []option.RequestOption
	Domains            *DomainService
	Registrations      *RegistrationService
	RegistrationStatus *RegistrationStatusService
	UpdateStatus       *UpdateStatusService
}

// NewRegistrarService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegistrarService(opts ...option.RequestOption) (r *RegistrarService) {
	r = &RegistrarService{}
	r.Options = opts
	r.Domains = NewDomainService(opts...)
	r.Registrations = NewRegistrationService(opts...)
	r.RegistrationStatus = NewRegistrationStatusService(opts...)
	r.UpdateStatus = NewUpdateStatusService(opts...)
	return
}

// Performs real-time, authoritative availability checks directly against domain
// registries. Use this endpoint to verify a domain is available before attempting
// registration via `POST /registrations`.
//
// **Important:** Unlike the Search endpoint, these results are authoritative and
// reflect current registry status. Always check availability immediately before
// registration as domain status can change rapidly.
//
// **Note:** This endpoint uses POST to accept a list of domains in the request
// body. It is a read-only operation — it does not create, modify, or reserve any
// domains.
//
// ### Extension support
//
// Only domains on extensions supported for programmatic registration by this API
// can be registered. If you check a domain on an unsupported extension, the
// response will include `registrable: false` with a `reason` field explaining why:
//
//   - `extension_not_supported_via_api` — Cloudflare Registrar supports this
//     extension in the dashboard, but it is not yet available for programmatic
//     registration via this API. Register via
//     `https://dash.cloudflare.com/{account_id}/domains/registrations` instead.
//   - `extension_not_supported` — This extension is not supported by Cloudflare
//     Registrar.
//   - `extension_disallows_registration` — The extension's registry has temporarily
//     or permanently frozen new registrations. No registrar can register domains on
//     this extension at this time.
//   - `domain_premium` — The domain is premium priced. Premium registration is not
//     currently supported by this API.
//   - `domain_unavailable` — The domain is already registered, reserved, or
//     otherwise not available for registration on a supported extension.
//
// The `reason` field is only present when `registrable` is `false`.
//
// ### Behavior
//
// - Maximum 20 domains per request
// - Pricing is only returned for domains where `registrable: true`
// - Results are not cached; each request queries the registry
//
// ### Workflow
//
//  1. Call this endpoint with domains the user wants to register.
//  2. For each domain where `registrable: true`, present pricing to the user.
//  3. If `tier: premium`, note that premium registration is not currently supported
//     by this API and do not proceed to `POST /registrations`.
//  4. Proceed to `POST /registrations` only for supported non-premium domains.
func (r *RegistrarService) Check(ctx context.Context, params RegistrarCheckParams, opts ...option.RequestOption) (res *RegistrarCheckResponse, err error) {
	var env RegistrarCheckResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar/domain-check", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Searches for domain name suggestions based on a keyword, phrase, or partial
// domain name. Returns a list of potentially available domains with pricing
// information.
//
// **Important:** Results are non-authoritative and based on cached data. Always
// use the `/domain-check` endpoint to verify real-time availability before
// attempting registration.
//
// Suggestions are scoped to extensions supported for programmatic registration via
// this API (`POST /registrations`). Domains on unsupported extensions will not
// appear in results, even if they are available at the registry level.
//
// ### Use cases
//
//   - Brand name discovery (e.g., "acme corp" → acmecorp.com, acmecorp.dev)
//   - Keyword-based suggestions (e.g., "coffee shop" → coffeeshop.com,
//     mycoffeeshop.net)
//   - Alternative extension discovery (e.g., "example.com" → example.com,
//     example.app, example.xyz)
//
// ### Workflow
//
//  1. Call this endpoint with a keyword or domain name.
//  2. Present suggestions to the user.
//  3. Call `/domain-check` with the user's chosen domains to confirm real-time
//     availability and pricing.
//  4. Proceed to `POST /registrations` only for supported non-premium domains where
//     the Check response returns `registrable: true`.
//
// **Note:** Searching with just a domain extension (e.g., "com" or ".app") is not
// supported. Provide a keyword or domain name.
func (r *RegistrarService) Search(ctx context.Context, params RegistrarSearchParams, opts ...option.RequestOption) (res *RegistrarSearchResponse, err error) {
	var env RegistrarSearchResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar/domain-search", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A domain registration resource representing the current state of a registered
// domain.
type Registration struct {
	// Whether the domain will be automatically renewed before expiration.
	AutoRenew bool `json:"auto_renew" api:"required"`
	// When the domain was registered. Present when the registration resource exists.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Fully qualified domain name (FQDN) including the extension (e.g., `example.com`,
	// `mybrand.app`). The domain name uniquely identifies a registration — the same
	// domain cannot be registered twice, making it a natural idempotency key for
	// registration requests.
	DomainName string `json:"domain_name" api:"required"`
	// When the domain registration expires. Present when the registration is ready;
	// may be null only while `status` is `registration_pending`.
	ExpiresAt time.Time `json:"expires_at" api:"required,nullable" format:"date-time"`
	// Whether the domain is locked for transfer.
	Locked bool `json:"locked" api:"required"`
	// Current WHOIS privacy mode for the registration.
	PrivacyMode RegistrationPrivacyMode `json:"privacy_mode" api:"required"`
	// Current registration status.
	//
	// - `active`: Domain is registered and operational
	// - `registration_pending`: Registration is in progress
	// - `expired`: Domain has expired
	// - `suspended`: Domain is suspended by the registry
	// - `redemption_period`: Domain is in the redemption grace period
	// - `pending_delete`: Domain is pending deletion by the registry
	Status RegistrationStatus `json:"status" api:"required"`
	JSON   registrationJSON   `json:"-"`
}

// registrationJSON contains the JSON metadata for the struct [Registration]
type registrationJSON struct {
	AutoRenew   apijson.Field
	CreatedAt   apijson.Field
	DomainName  apijson.Field
	ExpiresAt   apijson.Field
	Locked      apijson.Field
	PrivacyMode apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Registration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationJSON) RawJSON() string {
	return r.raw
}

// Current WHOIS privacy mode for the registration.
type RegistrationPrivacyMode string

const (
	RegistrationPrivacyModeRedaction RegistrationPrivacyMode = "redaction"
)

func (r RegistrationPrivacyMode) IsKnown() bool {
	switch r {
	case RegistrationPrivacyModeRedaction:
		return true
	}
	return false
}

// Current registration status.
//
// - `active`: Domain is registered and operational
// - `registration_pending`: Registration is in progress
// - `expired`: Domain has expired
// - `suspended`: Domain is suspended by the registry
// - `redemption_period`: Domain is in the redemption grace period
// - `pending_delete`: Domain is pending deletion by the registry
type RegistrationStatus string

const (
	RegistrationStatusActive              RegistrationStatus = "active"
	RegistrationStatusRegistrationPending RegistrationStatus = "registration_pending"
	RegistrationStatusExpired             RegistrationStatus = "expired"
	RegistrationStatusSuspended           RegistrationStatus = "suspended"
	RegistrationStatusRedemptionPeriod    RegistrationStatus = "redemption_period"
	RegistrationStatusPendingDelete       RegistrationStatus = "pending_delete"
)

func (r RegistrationStatus) IsKnown() bool {
	switch r {
	case RegistrationStatusActive, RegistrationStatusRegistrationPending, RegistrationStatusExpired, RegistrationStatusSuspended, RegistrationStatusRedemptionPeriod, RegistrationStatusPendingDelete:
		return true
	}
	return false
}

// Status of an async registration workflow.
type WorkflowStatus struct {
	// Whether the workflow has reached a terminal state. `true` when `state` is
	// `succeeded` or `failed`. `false` for `pending`, `in_progress`,
	// `action_required`, and `blocked`.
	Completed bool                `json:"completed" api:"required"`
	CreatedAt time.Time           `json:"created_at" api:"required" format:"date-time"`
	Links     WorkflowStatusLinks `json:"links" api:"required"`
	// Workflow lifecycle state.
	//
	//   - `pending`: Workflow has been created but not yet started processing.
	//   - `in_progress`: Actively processing. Continue polling `links.self`. The
	//     workflow has an internal deadline and will not remain in this state
	//     indefinitely.
	//   - `action_required`: Paused — requires action by the user (not the system). See
	//     `context.action` for what is needed. An automated polling loop must break on
	//     this state; it will not resolve on its own without user intervention.
	//   - `blocked`: The workflow cannot make progress due to a third party such as the
	//     domain extension's registry or a losing registrar. No user action will help.
	//     Continue polling — the block may resolve when the third party responds.
	//   - `succeeded`: Terminal. The operation completed successfully. `completed` will
	//     be `true`. For registrations, `context.registration` contains the resulting
	//     registration resource.
	//   - `failed`: Terminal. The operation failed. `completed` will be `true`. See
	//     `error.code` and `error.message` for the reason. Do not auto-retry without
	//     user review.
	State     WorkflowStatusState `json:"state" api:"required"`
	UpdatedAt time.Time           `json:"updated_at" api:"required" format:"date-time"`
	// Workflow-specific data for this workflow.
	//
	// The workflow subject is identified by `context.domain_name` for domain-centric
	// workflows.
	Context map[string]interface{} `json:"context"`
	// Error details when a workflow reaches the `failed` state. The specific error
	// codes and messages depend on the workflow type (registration, update, etc.) and
	// the underlying registry response. These workflow error codes are separate from
	// immediate HTTP error `errors[].code` values returned by non-2xx responses.
	// Surface `error.message` to the user for context.
	Error WorkflowStatusError `json:"error" api:"nullable"`
	JSON  workflowStatusJSON  `json:"-"`
}

// workflowStatusJSON contains the JSON metadata for the struct [WorkflowStatus]
type workflowStatusJSON struct {
	Completed   apijson.Field
	CreatedAt   apijson.Field
	Links       apijson.Field
	State       apijson.Field
	UpdatedAt   apijson.Field
	Context     apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowStatusJSON) RawJSON() string {
	return r.raw
}

type WorkflowStatusLinks struct {
	// URL to this status resource.
	Self string `json:"self" api:"required"`
	// URL to the domain resource.
	Resource string                  `json:"resource"`
	JSON     workflowStatusLinksJSON `json:"-"`
}

// workflowStatusLinksJSON contains the JSON metadata for the struct
// [WorkflowStatusLinks]
type workflowStatusLinksJSON struct {
	Self        apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowStatusLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowStatusLinksJSON) RawJSON() string {
	return r.raw
}

// Workflow lifecycle state.
//
//   - `pending`: Workflow has been created but not yet started processing.
//   - `in_progress`: Actively processing. Continue polling `links.self`. The
//     workflow has an internal deadline and will not remain in this state
//     indefinitely.
//   - `action_required`: Paused — requires action by the user (not the system). See
//     `context.action` for what is needed. An automated polling loop must break on
//     this state; it will not resolve on its own without user intervention.
//   - `blocked`: The workflow cannot make progress due to a third party such as the
//     domain extension's registry or a losing registrar. No user action will help.
//     Continue polling — the block may resolve when the third party responds.
//   - `succeeded`: Terminal. The operation completed successfully. `completed` will
//     be `true`. For registrations, `context.registration` contains the resulting
//     registration resource.
//   - `failed`: Terminal. The operation failed. `completed` will be `true`. See
//     `error.code` and `error.message` for the reason. Do not auto-retry without
//     user review.
type WorkflowStatusState string

const (
	WorkflowStatusStatePending        WorkflowStatusState = "pending"
	WorkflowStatusStateInProgress     WorkflowStatusState = "in_progress"
	WorkflowStatusStateActionRequired WorkflowStatusState = "action_required"
	WorkflowStatusStateBlocked        WorkflowStatusState = "blocked"
	WorkflowStatusStateSucceeded      WorkflowStatusState = "succeeded"
	WorkflowStatusStateFailed         WorkflowStatusState = "failed"
)

func (r WorkflowStatusState) IsKnown() bool {
	switch r {
	case WorkflowStatusStatePending, WorkflowStatusStateInProgress, WorkflowStatusStateActionRequired, WorkflowStatusStateBlocked, WorkflowStatusStateSucceeded, WorkflowStatusStateFailed:
		return true
	}
	return false
}

// Error details when a workflow reaches the `failed` state. The specific error
// codes and messages depend on the workflow type (registration, update, etc.) and
// the underlying registry response. These workflow error codes are separate from
// immediate HTTP error `errors[].code` values returned by non-2xx responses.
// Surface `error.message` to the user for context.
type WorkflowStatusError struct {
	// Machine-readable error code identifying the failure reason.
	Code string `json:"code" api:"required"`
	// Human-readable explanation of the failure. May include registry-specific
	// details.
	Message string                  `json:"message" api:"required"`
	JSON    workflowStatusErrorJSON `json:"-"`
}

// workflowStatusErrorJSON contains the JSON metadata for the struct
// [WorkflowStatusError]
type workflowStatusErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowStatusError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowStatusErrorJSON) RawJSON() string {
	return r.raw
}

// Contains the availability check results.
type RegistrarCheckResponse struct {
	// Array of domain availability results. Domains on unsupported extensions are
	// included with `registrable: false` and a `reason` field. Malformed domain names
	// may be omitted.
	Domains []RegistrarCheckResponseDomain `json:"domains" api:"required"`
	JSON    registrarCheckResponseJSON     `json:"-"`
}

// registrarCheckResponseJSON contains the JSON metadata for the struct
// [RegistrarCheckResponse]
type registrarCheckResponseJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarCheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarCheckResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a single authoritative domain availability result returned by the
// Check endpoint. Check results reflect current registry status and should be used
// immediately before registration.
type RegistrarCheckResponseDomain struct {
	// The fully qualified domain name (FQDN) in punycode format for internationalized
	// domain names (IDNs).
	Name string `json:"name" api:"required"`
	// Indicates whether this domain can be registered programmatically through this
	// API based on a real-time registry check.
	//
	//   - `true`: Domain is available for registration. The `pricing` object will be
	//     included.
	//   - `false`: Domain is not available. See the `reason` field for why. `tier` may
	//     still be present on some non-registrable results, such as premium domains.
	Registrable bool `json:"registrable" api:"required"`
	// Annual pricing information for a registrable domain. This object is only present
	// when `registrable` is `true`. All prices are per year and returned as strings to
	// preserve decimal precision.
	//
	// `registration_cost` and `renewal_cost` are frequently the same value, but may
	// differ — especially for premium domains where registries set different rates for
	// initial registration vs. renewal. For a multi-year registration (e.g., 4 years),
	// the first year is charged at `registration_cost` and each subsequent year at
	// `renewal_cost`. Registry pricing may change over time; the values returned here
	// reflect the current registry rate. Premium pricing may be surfaced by Search and
	// Check, but premium registration is not currently supported by this API.
	Pricing RegistrarCheckResponseDomainsPricing `json:"pricing"`
	// Present only when `registrable` is `false`. Explains why the domain cannot be
	// registered via this API.
	//
	//   - `extension_not_supported_via_api`: Cloudflare Registrar supports this
	//     extension in the dashboard but it is not yet available for programmatic
	//     registration via this API. The user can register via
	//     `https://dash.cloudflare.com/{account_id}/domains/registrations`.
	//   - `extension_not_supported`: This extension is not supported by Cloudflare
	//     Registrar at all.
	//   - `extension_disallows_registration`: The extension's registry has temporarily
	//     or permanently frozen new registrations. No registrar can register domains on
	//     this extension at this time.
	//   - `domain_premium`: The domain is premium priced. Premium registration is not
	//     currently supported by this API.
	//   - `domain_unavailable`: The domain is already registered, reserved, or otherwise
	//     not available on a supported extension.
	Reason RegistrarCheckResponseDomainsReason `json:"reason"`
	// The pricing tier for this domain. Always present when `registrable` is `true`;
	// defaults to `standard` for most domains. May be absent when `registrable` is
	// `false`.
	//
	// - `standard`: Standard registry pricing
	// - `premium`: Premium domain with higher pricing set by the registry
	Tier RegistrarCheckResponseDomainsTier `json:"tier"`
	JSON registrarCheckResponseDomainJSON  `json:"-"`
}

// registrarCheckResponseDomainJSON contains the JSON metadata for the struct
// [RegistrarCheckResponseDomain]
type registrarCheckResponseDomainJSON struct {
	Name        apijson.Field
	Registrable apijson.Field
	Pricing     apijson.Field
	Reason      apijson.Field
	Tier        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarCheckResponseDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarCheckResponseDomainJSON) RawJSON() string {
	return r.raw
}

// Annual pricing information for a registrable domain. This object is only present
// when `registrable` is `true`. All prices are per year and returned as strings to
// preserve decimal precision.
//
// `registration_cost` and `renewal_cost` are frequently the same value, but may
// differ — especially for premium domains where registries set different rates for
// initial registration vs. renewal. For a multi-year registration (e.g., 4 years),
// the first year is charged at `registration_cost` and each subsequent year at
// `renewal_cost`. Registry pricing may change over time; the values returned here
// reflect the current registry rate. Premium pricing may be surfaced by Search and
// Check, but premium registration is not currently supported by this API.
type RegistrarCheckResponseDomainsPricing struct {
	// ISO-4217 currency code for the prices (e.g., "USD", "EUR", "GBP").
	Currency string `json:"currency" api:"required"`
	// The first-year cost to register this domain. For premium domains
	// (`tier: premium`), this price is set by the registry and may be significantly
	// higher than standard pricing. For multi-year registrations, this cost applies to
	// the first year only; subsequent years are charged at `renewal_cost`.
	RegistrationCost string `json:"registration_cost" api:"required"`
	// Per-year renewal cost for this domain. Applied to each year beyond the first
	// year of a multi-year registration, and to each annual auto-renewal thereafter.
	// May differ from `registration_cost`, especially for premium domains where
	// initial registration often costs more than renewals.
	RenewalCost string                                   `json:"renewal_cost" api:"required"`
	JSON        registrarCheckResponseDomainsPricingJSON `json:"-"`
}

// registrarCheckResponseDomainsPricingJSON contains the JSON metadata for the
// struct [RegistrarCheckResponseDomainsPricing]
type registrarCheckResponseDomainsPricingJSON struct {
	Currency         apijson.Field
	RegistrationCost apijson.Field
	RenewalCost      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegistrarCheckResponseDomainsPricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarCheckResponseDomainsPricingJSON) RawJSON() string {
	return r.raw
}

// Present only when `registrable` is `false`. Explains why the domain cannot be
// registered via this API.
//
//   - `extension_not_supported_via_api`: Cloudflare Registrar supports this
//     extension in the dashboard but it is not yet available for programmatic
//     registration via this API. The user can register via
//     `https://dash.cloudflare.com/{account_id}/domains/registrations`.
//   - `extension_not_supported`: This extension is not supported by Cloudflare
//     Registrar at all.
//   - `extension_disallows_registration`: The extension's registry has temporarily
//     or permanently frozen new registrations. No registrar can register domains on
//     this extension at this time.
//   - `domain_premium`: The domain is premium priced. Premium registration is not
//     currently supported by this API.
//   - `domain_unavailable`: The domain is already registered, reserved, or otherwise
//     not available on a supported extension.
type RegistrarCheckResponseDomainsReason string

const (
	RegistrarCheckResponseDomainsReasonExtensionNotSupportedViaAPI    RegistrarCheckResponseDomainsReason = "extension_not_supported_via_api"
	RegistrarCheckResponseDomainsReasonExtensionNotSupported          RegistrarCheckResponseDomainsReason = "extension_not_supported"
	RegistrarCheckResponseDomainsReasonExtensionDisallowsRegistration RegistrarCheckResponseDomainsReason = "extension_disallows_registration"
	RegistrarCheckResponseDomainsReasonDomainPremium                  RegistrarCheckResponseDomainsReason = "domain_premium"
	RegistrarCheckResponseDomainsReasonDomainUnavailable              RegistrarCheckResponseDomainsReason = "domain_unavailable"
)

func (r RegistrarCheckResponseDomainsReason) IsKnown() bool {
	switch r {
	case RegistrarCheckResponseDomainsReasonExtensionNotSupportedViaAPI, RegistrarCheckResponseDomainsReasonExtensionNotSupported, RegistrarCheckResponseDomainsReasonExtensionDisallowsRegistration, RegistrarCheckResponseDomainsReasonDomainPremium, RegistrarCheckResponseDomainsReasonDomainUnavailable:
		return true
	}
	return false
}

// The pricing tier for this domain. Always present when `registrable` is `true`;
// defaults to `standard` for most domains. May be absent when `registrable` is
// `false`.
//
// - `standard`: Standard registry pricing
// - `premium`: Premium domain with higher pricing set by the registry
type RegistrarCheckResponseDomainsTier string

const (
	RegistrarCheckResponseDomainsTierStandard RegistrarCheckResponseDomainsTier = "standard"
	RegistrarCheckResponseDomainsTierPremium  RegistrarCheckResponseDomainsTier = "premium"
)

func (r RegistrarCheckResponseDomainsTier) IsKnown() bool {
	switch r {
	case RegistrarCheckResponseDomainsTierStandard, RegistrarCheckResponseDomainsTierPremium:
		return true
	}
	return false
}

// Contains the search results.
type RegistrarSearchResponse struct {
	// Array of domain suggestions sorted by relevance. May be empty if no domains
	// match the search criteria.
	Domains []RegistrarSearchResponseDomain `json:"domains" api:"required"`
	JSON    registrarSearchResponseJSON     `json:"-"`
}

// registrarSearchResponseJSON contains the JSON metadata for the struct
// [RegistrarSearchResponse]
type registrarSearchResponseJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSearchResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a single domain suggestion returned by the Search endpoint. Search
// results are non-authoritative and may be based on cached data. Use POST
// /domain-check to confirm real-time availability and pricing before registration.
type RegistrarSearchResponseDomain struct {
	// The fully qualified domain name (FQDN) in punycode format for internationalized
	// domain names (IDNs).
	Name string `json:"name" api:"required"`
	// Indicates whether this domain appears available based on search data. Search
	// results are non-authoritative and may be stale. - `true`: The domain appears
	// available. Use POST /domain-check to confirm before registration.
	//
	// - `false`: The domain does not appear available in search results.
	Registrable bool `json:"registrable" api:"required"`
	// Annual pricing information for a registrable domain. This object is only present
	// when `registrable` is `true`. All prices are per year and returned as strings to
	// preserve decimal precision.
	//
	// `registration_cost` and `renewal_cost` are frequently the same value, but may
	// differ — especially for premium domains where registries set different rates for
	// initial registration vs. renewal. For a multi-year registration (e.g., 4 years),
	// the first year is charged at `registration_cost` and each subsequent year at
	// `renewal_cost`. Registry pricing may change over time; the values returned here
	// reflect the current registry rate. Premium pricing may be surfaced by Search and
	// Check, but premium registration is not currently supported by this API.
	Pricing RegistrarSearchResponseDomainsPricing `json:"pricing"`
	// Present only when `registrable` is `false` on search results. Explains why the
	// domain does not appear registrable through this API. These values are advisory;
	// use POST /domain-check for authoritative status.
	//
	//   - `extension_not_supported_via_api`: Cloudflare Registrar supports this
	//     extension in the dashboard but it is not yet available for programmatic
	//     registration via this API.
	//   - `extension_not_supported`: This extension is not supported by Cloudflare
	//     Registrar at all.
	//   - `extension_disallows_registration`: The extension's registry has temporarily
	//     or permanently frozen new registrations.
	//   - `domain_premium`: The domain is premium priced. Premium registration is not
	//     currently supported by this API.
	//   - `domain_unavailable`: The domain appears unavailable.
	Reason RegistrarSearchResponseDomainsReason `json:"reason"`
	// The pricing tier for this domain. Always present when `registrable` is `true`;
	// defaults to `standard` for most domains. May be absent when `registrable` is
	// `false`.
	//
	// - `standard`: Standard registry pricing
	// - `premium`: Premium domain with higher pricing set by the registry
	Tier RegistrarSearchResponseDomainsTier `json:"tier"`
	JSON registrarSearchResponseDomainJSON  `json:"-"`
}

// registrarSearchResponseDomainJSON contains the JSON metadata for the struct
// [RegistrarSearchResponseDomain]
type registrarSearchResponseDomainJSON struct {
	Name        apijson.Field
	Registrable apijson.Field
	Pricing     apijson.Field
	Reason      apijson.Field
	Tier        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSearchResponseDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSearchResponseDomainJSON) RawJSON() string {
	return r.raw
}

// Annual pricing information for a registrable domain. This object is only present
// when `registrable` is `true`. All prices are per year and returned as strings to
// preserve decimal precision.
//
// `registration_cost` and `renewal_cost` are frequently the same value, but may
// differ — especially for premium domains where registries set different rates for
// initial registration vs. renewal. For a multi-year registration (e.g., 4 years),
// the first year is charged at `registration_cost` and each subsequent year at
// `renewal_cost`. Registry pricing may change over time; the values returned here
// reflect the current registry rate. Premium pricing may be surfaced by Search and
// Check, but premium registration is not currently supported by this API.
type RegistrarSearchResponseDomainsPricing struct {
	// ISO-4217 currency code for the prices (e.g., "USD", "EUR", "GBP").
	Currency string `json:"currency" api:"required"`
	// The first-year cost to register this domain. For premium domains
	// (`tier: premium`), this price is set by the registry and may be significantly
	// higher than standard pricing. For multi-year registrations, this cost applies to
	// the first year only; subsequent years are charged at `renewal_cost`.
	RegistrationCost string `json:"registration_cost" api:"required"`
	// Per-year renewal cost for this domain. Applied to each year beyond the first
	// year of a multi-year registration, and to each annual auto-renewal thereafter.
	// May differ from `registration_cost`, especially for premium domains where
	// initial registration often costs more than renewals.
	RenewalCost string                                    `json:"renewal_cost" api:"required"`
	JSON        registrarSearchResponseDomainsPricingJSON `json:"-"`
}

// registrarSearchResponseDomainsPricingJSON contains the JSON metadata for the
// struct [RegistrarSearchResponseDomainsPricing]
type registrarSearchResponseDomainsPricingJSON struct {
	Currency         apijson.Field
	RegistrationCost apijson.Field
	RenewalCost      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegistrarSearchResponseDomainsPricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSearchResponseDomainsPricingJSON) RawJSON() string {
	return r.raw
}

// Present only when `registrable` is `false` on search results. Explains why the
// domain does not appear registrable through this API. These values are advisory;
// use POST /domain-check for authoritative status.
//
//   - `extension_not_supported_via_api`: Cloudflare Registrar supports this
//     extension in the dashboard but it is not yet available for programmatic
//     registration via this API.
//   - `extension_not_supported`: This extension is not supported by Cloudflare
//     Registrar at all.
//   - `extension_disallows_registration`: The extension's registry has temporarily
//     or permanently frozen new registrations.
//   - `domain_premium`: The domain is premium priced. Premium registration is not
//     currently supported by this API.
//   - `domain_unavailable`: The domain appears unavailable.
type RegistrarSearchResponseDomainsReason string

const (
	RegistrarSearchResponseDomainsReasonExtensionNotSupportedViaAPI    RegistrarSearchResponseDomainsReason = "extension_not_supported_via_api"
	RegistrarSearchResponseDomainsReasonExtensionNotSupported          RegistrarSearchResponseDomainsReason = "extension_not_supported"
	RegistrarSearchResponseDomainsReasonExtensionDisallowsRegistration RegistrarSearchResponseDomainsReason = "extension_disallows_registration"
	RegistrarSearchResponseDomainsReasonDomainPremium                  RegistrarSearchResponseDomainsReason = "domain_premium"
	RegistrarSearchResponseDomainsReasonDomainUnavailable              RegistrarSearchResponseDomainsReason = "domain_unavailable"
)

func (r RegistrarSearchResponseDomainsReason) IsKnown() bool {
	switch r {
	case RegistrarSearchResponseDomainsReasonExtensionNotSupportedViaAPI, RegistrarSearchResponseDomainsReasonExtensionNotSupported, RegistrarSearchResponseDomainsReasonExtensionDisallowsRegistration, RegistrarSearchResponseDomainsReasonDomainPremium, RegistrarSearchResponseDomainsReasonDomainUnavailable:
		return true
	}
	return false
}

// The pricing tier for this domain. Always present when `registrable` is `true`;
// defaults to `standard` for most domains. May be absent when `registrable` is
// `false`.
//
// - `standard`: Standard registry pricing
// - `premium`: Premium domain with higher pricing set by the registry
type RegistrarSearchResponseDomainsTier string

const (
	RegistrarSearchResponseDomainsTierStandard RegistrarSearchResponseDomainsTier = "standard"
	RegistrarSearchResponseDomainsTierPremium  RegistrarSearchResponseDomainsTier = "premium"
)

func (r RegistrarSearchResponseDomainsTier) IsKnown() bool {
	switch r {
	case RegistrarSearchResponseDomainsTierStandard, RegistrarSearchResponseDomainsTierPremium:
		return true
	}
	return false
}

type RegistrarCheckParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// List of fully qualified domain names (FQDNs) to check for availability. Each
	// domain must include the extension.
	//
	//   - Minimum: 1 domain
	//   - Maximum: 20 domains per request
	//   - Domains on unsupported extensions are returned with `registrable: false` and a
	//     `reason` field
	//   - Malformed domain names (e.g., missing extension) may be omitted from the
	//     response
	Domains param.Field[[]string] `json:"domains" api:"required"`
}

func (r RegistrarCheckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegistrarCheckResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Contains the availability check results.
	Result RegistrarCheckResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success RegistrarCheckResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    registrarCheckResponseEnvelopeJSON    `json:"-"`
}

// registrarCheckResponseEnvelopeJSON contains the JSON metadata for the struct
// [RegistrarCheckResponseEnvelope]
type registrarCheckResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarCheckResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarCheckResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegistrarCheckResponseEnvelopeSuccess bool

const (
	RegistrarCheckResponseEnvelopeSuccessTrue RegistrarCheckResponseEnvelopeSuccess = true
)

func (r RegistrarCheckResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegistrarCheckResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RegistrarSearchParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The search term to find domain suggestions. Accepts keywords, phrases, or full
	// domain names.
	//
	//   - Phrases: "coffee shop" returns coffeeshop.com, mycoffeeshop.net, etc.
	//   - Domain names: "example.com" returns example.com and variations across
	//     extensions
	Q param.Field[string] `query:"q" api:"required"`
	// Limits results to specific domain extensions from the supported set. If not
	// specified, returns results across all supported extensions. Extensions not in
	// the supported set are silently ignored.
	Extensions param.Field[[]string] `query:"extensions"`
	// Maximum number of domain suggestions to return. Defaults to 20 if not specified.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [RegistrarSearchParams]'s query parameters as `url.Values`.
func (r RegistrarSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RegistrarSearchResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Contains the search results.
	Result RegistrarSearchResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success RegistrarSearchResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    registrarSearchResponseEnvelopeJSON    `json:"-"`
}

// registrarSearchResponseEnvelopeJSON contains the JSON metadata for the struct
// [RegistrarSearchResponseEnvelope]
type registrarSearchResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSearchResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSearchResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegistrarSearchResponseEnvelopeSuccess bool

const (
	RegistrarSearchResponseEnvelopeSuccessTrue RegistrarSearchResponseEnvelopeSuccess = true
)

func (r RegistrarSearchResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegistrarSearchResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
