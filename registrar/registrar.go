// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/tidwall/gjson"
)

// Registrar API for searching, checking, registering, and managing domains through
// Cloudflare Registrar.
//
// ## Prerequisites
//
// Before using this API, ensure:
//
//  1. **Cloudflare account** — the caller must have a valid Cloudflare account.
//  2. **Billing profile** — the account must have a billing profile with a valid,
//     current default payment method (credit card or other accepted method). This
//     cannot be set up via API — the account owner must configure billing at
//     `https://dash.cloudflare.com/{account_id}/billing/payment-info` before
//     calling `POST /registrations`.
//  3. **API authentication** — use an API token or API key with the appropriate
//     Registrar permissions for the operations you are calling.
//
// ## Terminology: domain extension
//
// Throughout this API, "extension" refers to the domain extension part of a fully
// qualified domain name — the portion after the registrable label. For example, in
// `example.co.uk`, the extension is `co.uk` (not just `uk`). This covers both
// top-level domains like `com` and multi-level extensions like `co.uk`. This is
// distinct from other uses of the word "extension" (e.g., EPP extensions).
//
// ## Supported extensions
//
// This API supports programmatic registration for all extensions supported by the
// dashboard experience, with the following exceptions:
//
// `giving`, `mom`, `inc`, `lol`, `sh`, `link`, `cc`, `new`
//
// Cloudflare Registrar supports 400+ extensions in the dashboard. Extensions
// listed above can be registered at
// `https://dash.cloudflare.com/{account_id}/domains/registrations`.
//
// ## Typical workflow
//
//  1. **Search** — call `GET /domain-search?q={keyword}` to discover available
//     domains.
//  2. **Check** — call `POST /domain-check` with candidate domains to verify
//     real-time availability and pricing.
//  3. **Review the response** — if `registrable: false`, inspect `reason` to
//     understand whether the domain is unavailable, the extension is not supported
//     by this API, the extension is not supported by Cloudflare Registrar at all,
//     or the extension's registry has frozen new registrations.
//  4. **Handle premium domains** — if `tier: premium`, premium registration is not
//     currently supported by this API. Surface the premium pricing to the user, but
//     do not proceed to `POST /registrations` for that domain.
//  5. **Observe the registration schema** — call `GET /extensions/:extension_name`
//     to discover the required values for registering this extension.
//  6. **Register** — call `POST /registrations` with the chosen domain name for
//     supported non-premium registrations.
//  7. **Confirm completion** — if the response is `201 Created`, registration
//     completed within the default timeout and no polling is needed.
//  8. **Poll when needed** — if the response is `202 Accepted`, poll `links.self`
//     from the workflow response.
//  9. **Stop for user action** — if `state: action_required`, stop polling and
//     surface `context.action` to the user. The workflow will not resolve on its
//     own.
//  10. **Continue when blocked** — if `state: blocked`, continue polling and inform
//     the user that a third party, such as the extension registry or losing
//     registrar, is delaying progress.
//  11. **Review failures before retrying** — if `state: failed`, review
//     `error.code` and `error.message`, then decide whether user action or a new
//     Check call is needed.
//
// **All successful domain registrations are non-refundable.** Once the
// registration workflow completes with `state: succeeded`, the charge cannot be
// reversed. Confirm pricing and domain choice with the user before calling
// `POST /registrations`.
//
// ## Default behavior for mutating operations
//
// By default, mutating operations such as create and update hold the connection
// for a bounded, server-defined amount of time while the operation completes. In
// most cases, the response contains a completed workflow status and no polling is
// required.
//
//   - **Completed within the synchronous wait window:** Returns `201` (create) or
//     `200` (update) with a `workflow_status` where `state: succeeded` and
//     `completed: true`.
//   - **Still processing after the synchronous wait window:** Returns `202 Accepted`
//     with a `workflow_status` where `completed: false`. Use the `links.self` URL to
//     poll for completion.
//
// ## Non-blocking mode
//
// To receive an immediate `202 Accepted` response without waiting, send the
// `Prefer: respond-async` request header (RFC 7240). The server will acknowledge
// it with a `Preference-Applied: respond-async` response header.
//
// ## Polling
//
// When the response is `202`, poll the workflow status endpoint indicated by
// `links.self` in the response body until the workflow reaches a terminal state or
// requires user action.
//
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
	Extensions         *ExtensionService
	TransferIn         *TransferInService
	TransferInStatus   *TransferInStatusService
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
	r.Extensions = NewExtensionService(opts...)
	r.TransferIn = NewTransferInService(opts...)
	r.TransferInStatus = NewTransferInStatusService(opts...)
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

// Performs real-time, authoritative eligibility checks directly against needed
// requirements. Use this endpoint to verify a domain is available before
// attempting a transfer via `POST /registrations/:domain_name/transfer-in`.
//
// **Note:** This endpoint uses POST to accept a list of domains in the request
// body. It is a read-only operation — it does not create, modify, or reserve any
// domains.
//
// ### Behavior
//
//   - Maximum 10 domains per request
//   - Pricing is only returned for domains where `transferable: true`
//   - Results are not cached; each request queries the registry & other needed
//     upstreams
//
// ## Extension Support
//
// All `.uk` extensions (`.uk`, `.co.uk`, etc) do not support auth codes. As such,
// Cloudflare will ignore the `auth_code` section of this request for `.uk`
// domains.
//
// This means that a `.uk` domain depends on public data to obtain domain
// information, so it might be a few minutes outdated.
//
// ### Workflow
//
//  1. Call this endpoint with domains the user wants to transfer.
//  2. For each domain where `transferable: true`, present pricing to the user.
//  3. For each domain where `transferable: false`, present reasons to the user
//  4. Proceed to `POST /registrations/:domain_name/transfer-in` only for the
//     `transferable: true` domains.
func (r *RegistrarService) TransferCheck(ctx context.Context, params RegistrarTransferCheckParams, opts ...option.RequestOption) (res *RegistrarTransferCheckResponse, err error) {
	var env RegistrarTransferCheckResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar/domain-transfer-check", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A domain registration resource representing the current state of a registered
// domain.
type Registration struct {
	// Whether automatic renewal occurs before expiration.
	AutoRenew bool `json:"auto_renew" api:"required"`
	// When the domain was registered. Present when the registration resource exists.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Provides a fully qualified domain name (FQDN), including the extension (e.g.,
	// `example.com`, `mybrand.app`). The domain name uniquely identifies a
	// registration. Cloudflare permits only one registration per domain, making the
	// domain name a natural idempotency key for registration requests.
	DomainName string `json:"domain_name" api:"required"`
	// When the domain registration expires. Ready registrations include this value;
	// only `registration_pending` and `transfer_pending` may return null.
	ExpiresAt time.Time `json:"expires_at" api:"required,nullable" format:"date-time"`
	// Whether the domain is locked for transfer.
	Locked bool `json:"locked" api:"required"`
	// Current WHOIS privacy mode for the registration.
	PrivacyMode RegistrationPrivacyMode `json:"privacy_mode" api:"required"`
	// Current registration status.
	//
	// - `active`: The domain operates with an active registration.
	// - `registration_pending`: Registration remains in progress.
	// - `transfer_pending`: Domain transfer is in progress.
	// - `expired`: The domain registration expired.
	// - `suspended`: The registry suspended the domain.
	// - `redemption_period`: The domain entered the redemption grace period.
	// - `pending_delete`: The registry scheduled the domain for deletion.
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
	RegistrationPrivacyModeOff       RegistrationPrivacyMode = "off"
	RegistrationPrivacyModeRedaction RegistrationPrivacyMode = "redaction"
)

func (r RegistrationPrivacyMode) IsKnown() bool {
	switch r {
	case RegistrationPrivacyModeOff, RegistrationPrivacyModeRedaction:
		return true
	}
	return false
}

// Current registration status.
//
// - `active`: The domain operates with an active registration.
// - `registration_pending`: Registration remains in progress.
// - `transfer_pending`: Domain transfer is in progress.
// - `expired`: The domain registration expired.
// - `suspended`: The registry suspended the domain.
// - `redemption_period`: The domain entered the redemption grace period.
// - `pending_delete`: The registry scheduled the domain for deletion.
type RegistrationStatus string

const (
	RegistrationStatusActive              RegistrationStatus = "active"
	RegistrationStatusRegistrationPending RegistrationStatus = "registration_pending"
	RegistrationStatusTransferPending     RegistrationStatus = "transfer_pending"
	RegistrationStatusExpired             RegistrationStatus = "expired"
	RegistrationStatusSuspended           RegistrationStatus = "suspended"
	RegistrationStatusRedemptionPeriod    RegistrationStatus = "redemption_period"
	RegistrationStatusPendingDelete       RegistrationStatus = "pending_delete"
)

func (r RegistrationStatus) IsKnown() bool {
	switch r {
	case RegistrationStatusActive, RegistrationStatusRegistrationPending, RegistrationStatusTransferPending, RegistrationStatusExpired, RegistrationStatusSuspended, RegistrationStatusRedemptionPeriod, RegistrationStatusPendingDelete:
		return true
	}
	return false
}

// Status of an async registration workflow.
type WorkflowStatus struct {
	// Indicates whether the workflow reached a terminal state. A `succeeded` or
	// `failed` state returns `true`; `pending`, `in_progress`, `action_required`, and
	// `blocked` return `false`.
	Completed bool                `json:"completed" api:"required"`
	CreatedAt time.Time           `json:"created_at" api:"required" format:"date-time"`
	Links     WorkflowStatusLinks `json:"links" api:"required"`
	// Describes the workflow lifecycle state.
	//
	// - `pending`: The workflow awaits processing.
	// - `in_progress`: Processing started. Continue polling `links.self`. An internal
	//   deadline limits the duration of this state.
	// - `action_required`: The workflow pauses for user action. See `context.action`
	//   for details. Stop automated polling until the user completes the required
	//   action.
	// - `blocked`: A third party, such as the domain extension's registry or a losing
	//   registrar, prevents progress. Continue polling because the block may resolve
	//   when the third party responds.
	// - `succeeded`: Terminal state. The operation completed successfully. `completed`
	//   equals `true`. For registrations, `context.registration` contains the
	//   resulting registration resource.
	// - `failed`: Terminal state. The operation failed. `completed` equals `true`. See
	//   `error.code` and `error.message` for the reason. Require user review before
	//   retrying.
	State     WorkflowStatusState `json:"state" api:"required"`
	UpdatedAt time.Time           `json:"updated_at" api:"required" format:"date-time"`
	// Provides workflow-specific data.
	//
	// For domain-centric workflows, `context.domain_name` identifies the workflow
	// subject.
	Context map[string]interface{} `json:"context"`
	// Provides error details when a workflow reaches the `failed` state. The workflow
	// type (registration, update, etc.) and underlying registry response determine the
	// specific codes and messages. Workflow error codes differ from immediate HTTP
	// error `errors[].code` values in non-2xx responses. Surface `error.message` to
	// the user for context.
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

// Describes the workflow lifecycle state.
//
//   - `pending`: The workflow awaits processing.
//   - `in_progress`: Processing started. Continue polling `links.self`. An internal
//     deadline limits the duration of this state.
//   - `action_required`: The workflow pauses for user action. See `context.action`
//     for details. Stop automated polling until the user completes the required
//     action.
//   - `blocked`: A third party, such as the domain extension's registry or a losing
//     registrar, prevents progress. Continue polling because the block may resolve
//     when the third party responds.
//   - `succeeded`: Terminal state. The operation completed successfully. `completed`
//     equals `true`. For registrations, `context.registration` contains the
//     resulting registration resource.
//   - `failed`: Terminal state. The operation failed. `completed` equals `true`. See
//     `error.code` and `error.message` for the reason. Require user review before
//     retrying.
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

// Provides error details when a workflow reaches the `failed` state. The workflow
// type (registration, update, etc.) and underlying registry response determine the
// specific codes and messages. Workflow error codes differ from immediate HTTP
// error `errors[].code` values in non-2xx responses. Surface `error.message` to
// the user for context.
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
	// Array of domain availability results. Results for unsupported extensions contain
	// `registrable: false` and a `reason` field. The response may omit malformed
	// domain names.
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

// Describes a single authoritative domain availability result from the Check
// endpoint. Check results reflect current registry status; use them immediately
// before registration.
type RegistrarCheckResponseDomain struct {
	// The fully qualified domain name (FQDN) in punycode format for internationalized
	// domain names (IDNs).
	Name string `json:"name" api:"required"`
	// Indicates programmatic registration eligibility according to a real-time
	// registry check.
	//
	// - `true`: The domain is available for registration. The response includes the
	//   `pricing` object.
	// - `false`: A restriction prevents registration. See the `reason` field for
	//   details. Some results, such as premium domains, may still include `tier`.
	Registrable bool `json:"registrable" api:"required"`
	// Provides annual pricing information for a given domain. The API returns all
	// per-year prices as strings to preserve decimal precision.
	//
	// `renewal_cost` and `registration_cost` or `transfer_cost` are frequently the
	// same value, but may differ due to premium rates for certain domains.
	//
	// For a multi-year operations, the operation's cost applies to the first year and
	// `renewal_cost` applies to each subsequent year. The values reflect the current
	// registry rate, which can change over time.
	Pricing RegistrarCheckResponseDomainsPricing `json:"pricing"`
	// Appears only when `registrable` is `false` and explains the result.
	//
	// - `extension_not_supported_via_api`: Cloudflare Registrar supports this
	//   extension in the dashboard but currently excludes it from programmatic
	//   registration through this API. The user can register via
	//   `https://dash.cloudflare.com/{account_id}/domains/registrations`.
	// - `extension_not_supported`: Cloudflare Registrar excludes this extension
	//   entirely.
	// - `extension_disallows_registration`: The extension's registry temporarily or
	//   permanently freezes new registrations. Registrars currently cannot register
	//   domains on this extension.
	// - `domain_premium`: The domain carries premium pricing. This API currently
	//   supports standard registrations only.
	// - `domain_unavailable`: An existing registration, reservation, or other registry
	//   restriction makes the domain unavailable on a supported extension.
	Reason RegistrarCheckResponseDomainsReason `json:"reason"`
	// The pricing tier for this domain. A `registrable` value of `true` always
	// includes this field, which defaults to `standard` for most domains. A
	// `registrable` value of `false` may omit it.
	//
	// - `standard`: Standard registry pricing.
	// - `premium`: Premium domain with higher pricing from the registry.
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

// Provides annual pricing information for a given domain. The API returns all
// per-year prices as strings to preserve decimal precision.
//
// `renewal_cost` and `registration_cost` or `transfer_cost` are frequently the
// same value, but may differ due to premium rates for certain domains.
//
// For a multi-year operations, the operation's cost applies to the first year and
// `renewal_cost` applies to each subsequent year. The values reflect the current
// registry rate, which can change over time.
type RegistrarCheckResponseDomainsPricing struct {
	// ISO-4217 currency code for the prices (e.g., "USD", "EUR", "GBP").
	Currency string `json:"currency" api:"required"`
	// The first-year cost to register this domain.
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

// Appears only when `registrable` is `false` and explains the result.
//
//   - `extension_not_supported_via_api`: Cloudflare Registrar supports this
//     extension in the dashboard but currently excludes it from programmatic
//     registration through this API. The user can register via
//     `https://dash.cloudflare.com/{account_id}/domains/registrations`.
//   - `extension_not_supported`: Cloudflare Registrar excludes this extension
//     entirely.
//   - `extension_disallows_registration`: The extension's registry temporarily or
//     permanently freezes new registrations. Registrars currently cannot register
//     domains on this extension.
//   - `domain_premium`: The domain carries premium pricing. This API currently
//     supports standard registrations only.
//   - `domain_unavailable`: An existing registration, reservation, or other registry
//     restriction makes the domain unavailable on a supported extension.
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

// The pricing tier for this domain. A `registrable` value of `true` always
// includes this field, which defaults to `standard` for most domains. A
// `registrable` value of `false` may omit it.
//
// - `standard`: Standard registry pricing.
// - `premium`: Premium domain with higher pricing from the registry.
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
	// Lists domain suggestions in relevance order. An empty array indicates that the
	// search criteria matched zero domains.
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

// Describes a single domain suggestion from the Search endpoint. Search results
// use non-authoritative data that may come from a cache. Use POST /domain-check to
// confirm real-time availability and pricing before registration.
type RegistrarSearchResponseDomain struct {
	// The fully qualified domain name (FQDN) in punycode format for internationalized
	// domain names (IDNs).
	Name string `json:"name" api:"required"`
	// Indicates domain availability according to potentially stale, non-authoritative
	// search data.
	//
	// - `true`: The domain appears available. Use POST /domain-check to confirm before
	//   registration.
	// - `false`: Search results mark the domain ineligible for registration through
	//   this API. See `reason` for details.
	Registrable bool `json:"registrable" api:"required"`
	// Provides annual pricing information for a given domain. The API returns all
	// per-year prices as strings to preserve decimal precision.
	//
	// `renewal_cost` and `registration_cost` or `transfer_cost` are frequently the
	// same value, but may differ due to premium rates for certain domains.
	//
	// For a multi-year operations, the operation's cost applies to the first year and
	// `renewal_cost` applies to each subsequent year. The values reflect the current
	// registry rate, which can change over time.
	Pricing RegistrarSearchResponseDomainsPricing `json:"pricing"`
	// Appears only when `registrable` is `false` and explains the advisory search
	// result. Use POST /domain-check for authoritative status.
	//
	// - `extension_not_supported_via_api`: Cloudflare Registrar supports this
	//   extension in the dashboard but currently excludes it from programmatic
	//   registration through this API.
	// - `extension_not_supported`: Cloudflare Registrar excludes this extension
	//   entirely.
	// - `extension_disallows_registration`: The extension's registry temporarily or
	//   permanently freezes new registrations.
	// - `domain_premium`: The domain carries premium pricing. This API currently
	//   supports standard registrations only.
	// - `domain_unavailable`: The domain appears unavailable.
	Reason RegistrarSearchResponseDomainsReason `json:"reason"`
	// The pricing tier for this domain. A `registrable` value of `true` always
	// includes this field, which defaults to `standard` for most domains. A
	// `registrable` value of `false` may omit it.
	//
	// - `standard`: Standard registry pricing.
	// - `premium`: Premium domain with higher pricing from the registry.
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

// Provides annual pricing information for a given domain. The API returns all
// per-year prices as strings to preserve decimal precision.
//
// `renewal_cost` and `registration_cost` or `transfer_cost` are frequently the
// same value, but may differ due to premium rates for certain domains.
//
// For a multi-year operations, the operation's cost applies to the first year and
// `renewal_cost` applies to each subsequent year. The values reflect the current
// registry rate, which can change over time.
type RegistrarSearchResponseDomainsPricing struct {
	// ISO-4217 currency code for the prices (e.g., "USD", "EUR", "GBP").
	Currency string `json:"currency" api:"required"`
	// The first-year cost to register this domain.
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

// Appears only when `registrable` is `false` and explains the advisory search
// result. Use POST /domain-check for authoritative status.
//
//   - `extension_not_supported_via_api`: Cloudflare Registrar supports this
//     extension in the dashboard but currently excludes it from programmatic
//     registration through this API.
//   - `extension_not_supported`: Cloudflare Registrar excludes this extension
//     entirely.
//   - `extension_disallows_registration`: The extension's registry temporarily or
//     permanently freezes new registrations.
//   - `domain_premium`: The domain carries premium pricing. This API currently
//     supports standard registrations only.
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

// The pricing tier for this domain. A `registrable` value of `true` always
// includes this field, which defaults to `standard` for most domains. A
// `registrable` value of `false` may omit it.
//
// - `standard`: Standard registry pricing.
// - `premium`: Premium domain with higher pricing from the registry.
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

// Contains the transfer eligibility results.
type RegistrarTransferCheckResponse struct {
	// Maps domain names to transfer eligibility results. Each value contains `name`,
	// `transferable`, and `reasons`.
	Domains map[string]RegistrarTransferCheckResponseDomain `json:"domains" api:"required"`
	JSON    registrarTransferCheckResponseJSON              `json:"-"`
}

// registrarTransferCheckResponseJSON contains the JSON metadata for the struct
// [RegistrarTransferCheckResponse]
type registrarTransferCheckResponseJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarTransferCheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarTransferCheckResponseJSON) RawJSON() string {
	return r.raw
}

// Transfer eligibility for a single domain. `reasons` is always present:
//
// - Empty when `transferable` is `true`.
// - One or more reason objects when `transferable` is `false`.
type RegistrarTransferCheckResponseDomain struct {
	Transferable RegistrarTransferCheckResponseDomainsTransferable `json:"transferable" api:"required"`
	// The check evaluates this domain name.
	Name string `json:"name"`
	// This field can have the runtime type of
	// [RegistrarTransferCheckResponseDomainsTransferableResultPricing],
	// [RegistrarTransferCheckResponseDomainsNonTransferableResultPricing].
	Pricing interface{} `json:"pricing"`
	// This field can have the runtime type of
	// [[]RegistrarTransferCheckResponseDomainsTransferableResultReason],
	// [[]RegistrarTransferCheckResponseDomainsNonTransferableResultReason].
	Reasons interface{}                              `json:"reasons"`
	JSON    registrarTransferCheckResponseDomainJSON `json:"-"`
	union   RegistrarTransferCheckResponseDomainsUnion
}

// registrarTransferCheckResponseDomainJSON contains the JSON metadata for the
// struct [RegistrarTransferCheckResponseDomain]
type registrarTransferCheckResponseDomainJSON struct {
	Transferable apijson.Field
	Name         apijson.Field
	Pricing      apijson.Field
	Reasons      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r registrarTransferCheckResponseDomainJSON) RawJSON() string {
	return r.raw
}

func (r *RegistrarTransferCheckResponseDomain) UnmarshalJSON(data []byte) (err error) {
	*r = RegistrarTransferCheckResponseDomain{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RegistrarTransferCheckResponseDomainsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RegistrarTransferCheckResponseDomainsTransferableResult],
// [RegistrarTransferCheckResponseDomainsNonTransferableResult].
func (r RegistrarTransferCheckResponseDomain) AsUnion() RegistrarTransferCheckResponseDomainsUnion {
	return r.union
}

// Transfer eligibility for a single domain. `reasons` is always present:
//
// - Empty when `transferable` is `true`.
// - One or more reason objects when `transferable` is `false`.
//
// Union satisfied by [RegistrarTransferCheckResponseDomainsTransferableResult] or
// [RegistrarTransferCheckResponseDomainsNonTransferableResult].
type RegistrarTransferCheckResponseDomainsUnion interface {
	implementsRegistrarTransferCheckResponseDomain()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RegistrarTransferCheckResponseDomainsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RegistrarTransferCheckResponseDomainsTransferableResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RegistrarTransferCheckResponseDomainsNonTransferableResult{}),
		},
	)
}

type RegistrarTransferCheckResponseDomainsTransferableResult struct {
	// Provides annual pricing information for a given domain. The API returns all
	// per-year prices as strings to preserve decimal precision.
	//
	// `renewal_cost` and `registration_cost` or `transfer_cost` are frequently the
	// same value, but may differ due to premium rates for certain domains.
	//
	// For a multi-year operations, the operation's cost applies to the first year and
	// `renewal_cost` applies to each subsequent year. The values reflect the current
	// registry rate, which can change over time.
	Pricing      RegistrarTransferCheckResponseDomainsTransferableResultPricing      `json:"pricing" api:"required"`
	Transferable RegistrarTransferCheckResponseDomainsTransferableResultTransferable `json:"transferable" api:"required"`
	// The check evaluates this domain name.
	Name    string                                                          `json:"name"`
	Reasons []RegistrarTransferCheckResponseDomainsTransferableResultReason `json:"reasons"`
	JSON    registrarTransferCheckResponseDomainsTransferableResultJSON     `json:"-"`
}

// registrarTransferCheckResponseDomainsTransferableResultJSON contains the JSON
// metadata for the struct
// [RegistrarTransferCheckResponseDomainsTransferableResult]
type registrarTransferCheckResponseDomainsTransferableResultJSON struct {
	Pricing      apijson.Field
	Transferable apijson.Field
	Name         apijson.Field
	Reasons      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RegistrarTransferCheckResponseDomainsTransferableResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarTransferCheckResponseDomainsTransferableResultJSON) RawJSON() string {
	return r.raw
}

func (r RegistrarTransferCheckResponseDomainsTransferableResult) implementsRegistrarTransferCheckResponseDomain() {
}

// Provides annual pricing information for a given domain. The API returns all
// per-year prices as strings to preserve decimal precision.
//
// `renewal_cost` and `registration_cost` or `transfer_cost` are frequently the
// same value, but may differ due to premium rates for certain domains.
//
// For a multi-year operations, the operation's cost applies to the first year and
// `renewal_cost` applies to each subsequent year. The values reflect the current
// registry rate, which can change over time.
type RegistrarTransferCheckResponseDomainsTransferableResultPricing struct {
	// ISO-4217 currency code for the prices (e.g., "USD", "EUR", "GBP").
	Currency string `json:"currency" api:"required"`
	// Per-year renewal cost for this domain. Applied to each year beyond the first
	// year of a multi-year registration, and to each annual auto-renewal thereafter.
	// May differ from `registration_cost`, especially for premium domains where
	// initial registration often costs more than renewals.
	RenewalCost string `json:"renewal_cost" api:"required"`
	// The first-year cost to transfer this domain.
	TransferCost string                                                             `json:"transfer_cost" api:"required"`
	JSON         registrarTransferCheckResponseDomainsTransferableResultPricingJSON `json:"-"`
}

// registrarTransferCheckResponseDomainsTransferableResultPricingJSON contains the
// JSON metadata for the struct
// [RegistrarTransferCheckResponseDomainsTransferableResultPricing]
type registrarTransferCheckResponseDomainsTransferableResultPricingJSON struct {
	Currency     apijson.Field
	RenewalCost  apijson.Field
	TransferCost apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RegistrarTransferCheckResponseDomainsTransferableResultPricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarTransferCheckResponseDomainsTransferableResultPricingJSON) RawJSON() string {
	return r.raw
}

type RegistrarTransferCheckResponseDomainsTransferableResultTransferable bool

const (
	RegistrarTransferCheckResponseDomainsTransferableResultTransferableTrue RegistrarTransferCheckResponseDomainsTransferableResultTransferable = true
)

func (r RegistrarTransferCheckResponseDomainsTransferableResultTransferable) IsKnown() bool {
	switch r {
	case RegistrarTransferCheckResponseDomainsTransferableResultTransferableTrue:
		return true
	}
	return false
}

type RegistrarTransferCheckResponseDomainsTransferableResultReason struct {
	// Transfer eligibility reason code.
	//
	// - `extension_not_supported_via_api`: This API excludes the extension; dashboard
	//   flows support it.
	// - `extension_not_supported`: Cloudflare Registrar excludes the extension.
	// - `domain_premium`: This API currently excludes premium transfers.
	// - `extension_disallows_transfer`: Extension currently blocks transfer
	//   operations.
	// - `domain_not_exists`: No registration record exists for the domain.
	// - `domain_on_cloudflare`: Cloudflare already serves as the domain's registrar.
	// - `domain_locked`: Losing registrar reports transfer-prohibited lock status.
	// - `registry_status`: Registry status currently blocks transfer (for example,
	//   pending transfer or deletion state).
	// - `domain_outside_transfer_window`: Domain is within a transfer wait window (for
	//   example, recently registered).
	// - `domain_max_term`: Completing transfer would exceed the registry maximum term.
	// - `invalid_auth_code`: The provided auth code is incorrect.
	// - `invalid_auth_code_format`: Auth code fails Base64 validation.
	// - `dnssec_enabled`: DNSSEC is enabled. It must be disabled before transfer.
	// - `zone_not_found`: The target account lacks a Cloudflare zone for the domain.
	// - `zone_status_invalid`: The Cloudflare zone cannot transfer in its current
	//   state.
	// - `invalid_zone_plan`: The zone plan fails transfer requirements.
	// - `domain_unsupported`: This endpoint rejects the domain name format.
	Code RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode `json:"code" api:"required"`
	JSON registrarTransferCheckResponseDomainsTransferableResultReasonJSON  `json:"-"`
}

// registrarTransferCheckResponseDomainsTransferableResultReasonJSON contains the
// JSON metadata for the struct
// [RegistrarTransferCheckResponseDomainsTransferableResultReason]
type registrarTransferCheckResponseDomainsTransferableResultReasonJSON struct {
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarTransferCheckResponseDomainsTransferableResultReason) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarTransferCheckResponseDomainsTransferableResultReasonJSON) RawJSON() string {
	return r.raw
}

// Transfer eligibility reason code.
//
//   - `extension_not_supported_via_api`: This API excludes the extension; dashboard
//     flows support it.
//   - `extension_not_supported`: Cloudflare Registrar excludes the extension.
//   - `domain_premium`: This API currently excludes premium transfers.
//   - `extension_disallows_transfer`: Extension currently blocks transfer
//     operations.
//   - `domain_not_exists`: No registration record exists for the domain.
//   - `domain_on_cloudflare`: Cloudflare already serves as the domain's registrar.
//   - `domain_locked`: Losing registrar reports transfer-prohibited lock status.
//   - `registry_status`: Registry status currently blocks transfer (for example,
//     pending transfer or deletion state).
//   - `domain_outside_transfer_window`: Domain is within a transfer wait window (for
//     example, recently registered).
//   - `domain_max_term`: Completing transfer would exceed the registry maximum term.
//   - `invalid_auth_code`: The provided auth code is incorrect.
//   - `invalid_auth_code_format`: Auth code fails Base64 validation.
//   - `dnssec_enabled`: DNSSEC is enabled. It must be disabled before transfer.
//   - `zone_not_found`: The target account lacks a Cloudflare zone for the domain.
//   - `zone_status_invalid`: The Cloudflare zone cannot transfer in its current
//     state.
//   - `invalid_zone_plan`: The zone plan fails transfer requirements.
//   - `domain_unsupported`: This endpoint rejects the domain name format.
type RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode string

const (
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeExtensionNotSupportedViaAPI RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "extension_not_supported_via_api"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeExtensionNotSupported       RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "extension_not_supported"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainPremium               RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "domain_premium"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeExtensionDisallowsTransfer  RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "extension_disallows_transfer"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainNotExists             RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "domain_not_exists"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainOnCloudflare          RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "domain_on_cloudflare"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainLocked                RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "domain_locked"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeRegistryStatus              RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "registry_status"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainOutsideTransferWindow RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "domain_outside_transfer_window"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainMaxTerm               RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "domain_max_term"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeInvalidAuthCode             RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "invalid_auth_code"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeInvalidAuthCodeFormat       RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "invalid_auth_code_format"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDNSSECEnabled               RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "dnssec_enabled"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeZoneNotFound                RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "zone_not_found"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeZoneStatusInvalid           RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "zone_status_invalid"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeInvalidZonePlan             RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "invalid_zone_plan"
	RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainUnsupported           RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode = "domain_unsupported"
)

func (r RegistrarTransferCheckResponseDomainsTransferableResultReasonsCode) IsKnown() bool {
	switch r {
	case RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeExtensionNotSupportedViaAPI, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeExtensionNotSupported, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainPremium, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeExtensionDisallowsTransfer, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainNotExists, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainOnCloudflare, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainLocked, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeRegistryStatus, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainOutsideTransferWindow, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainMaxTerm, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeInvalidAuthCode, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeInvalidAuthCodeFormat, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDNSSECEnabled, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeZoneNotFound, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeZoneStatusInvalid, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeInvalidZonePlan, RegistrarTransferCheckResponseDomainsTransferableResultReasonsCodeDomainUnsupported:
		return true
	}
	return false
}

type RegistrarTransferCheckResponseDomainsNonTransferableResult struct {
	Transferable RegistrarTransferCheckResponseDomainsNonTransferableResultTransferable `json:"transferable" api:"required"`
	// The check evaluates this domain name.
	Name string `json:"name"`
	// Provides annual pricing information for a given domain. The API returns all
	// per-year prices as strings to preserve decimal precision.
	//
	// `renewal_cost` and `registration_cost` or `transfer_cost` are frequently the
	// same value, but may differ due to premium rates for certain domains.
	//
	// For a multi-year operations, the operation's cost applies to the first year and
	// `renewal_cost` applies to each subsequent year. The values reflect the current
	// registry rate, which can change over time.
	Pricing RegistrarTransferCheckResponseDomainsNonTransferableResultPricing  `json:"pricing"`
	Reasons []RegistrarTransferCheckResponseDomainsNonTransferableResultReason `json:"reasons"`
	JSON    registrarTransferCheckResponseDomainsNonTransferableResultJSON     `json:"-"`
}

// registrarTransferCheckResponseDomainsNonTransferableResultJSON contains the JSON
// metadata for the struct
// [RegistrarTransferCheckResponseDomainsNonTransferableResult]
type registrarTransferCheckResponseDomainsNonTransferableResultJSON struct {
	Transferable apijson.Field
	Name         apijson.Field
	Pricing      apijson.Field
	Reasons      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RegistrarTransferCheckResponseDomainsNonTransferableResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarTransferCheckResponseDomainsNonTransferableResultJSON) RawJSON() string {
	return r.raw
}

func (r RegistrarTransferCheckResponseDomainsNonTransferableResult) implementsRegistrarTransferCheckResponseDomain() {
}

type RegistrarTransferCheckResponseDomainsNonTransferableResultTransferable bool

const (
	RegistrarTransferCheckResponseDomainsNonTransferableResultTransferableFalse RegistrarTransferCheckResponseDomainsNonTransferableResultTransferable = false
)

func (r RegistrarTransferCheckResponseDomainsNonTransferableResultTransferable) IsKnown() bool {
	switch r {
	case RegistrarTransferCheckResponseDomainsNonTransferableResultTransferableFalse:
		return true
	}
	return false
}

// Provides annual pricing information for a given domain. The API returns all
// per-year prices as strings to preserve decimal precision.
//
// `renewal_cost` and `registration_cost` or `transfer_cost` are frequently the
// same value, but may differ due to premium rates for certain domains.
//
// For a multi-year operations, the operation's cost applies to the first year and
// `renewal_cost` applies to each subsequent year. The values reflect the current
// registry rate, which can change over time.
type RegistrarTransferCheckResponseDomainsNonTransferableResultPricing struct {
	// ISO-4217 currency code for the prices (e.g., "USD", "EUR", "GBP").
	Currency string `json:"currency" api:"required"`
	// Per-year renewal cost for this domain. Applied to each year beyond the first
	// year of a multi-year registration, and to each annual auto-renewal thereafter.
	// May differ from `registration_cost`, especially for premium domains where
	// initial registration often costs more than renewals.
	RenewalCost string `json:"renewal_cost" api:"required"`
	// The first-year cost to transfer this domain.
	TransferCost string                                                                `json:"transfer_cost" api:"required"`
	JSON         registrarTransferCheckResponseDomainsNonTransferableResultPricingJSON `json:"-"`
}

// registrarTransferCheckResponseDomainsNonTransferableResultPricingJSON contains
// the JSON metadata for the struct
// [RegistrarTransferCheckResponseDomainsNonTransferableResultPricing]
type registrarTransferCheckResponseDomainsNonTransferableResultPricingJSON struct {
	Currency     apijson.Field
	RenewalCost  apijson.Field
	TransferCost apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RegistrarTransferCheckResponseDomainsNonTransferableResultPricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarTransferCheckResponseDomainsNonTransferableResultPricingJSON) RawJSON() string {
	return r.raw
}

type RegistrarTransferCheckResponseDomainsNonTransferableResultReason struct {
	// Transfer eligibility reason code.
	//
	// - `extension_not_supported_via_api`: This API excludes the extension; dashboard
	//   flows support it.
	// - `extension_not_supported`: Cloudflare Registrar excludes the extension.
	// - `domain_premium`: This API currently excludes premium transfers.
	// - `extension_disallows_transfer`: Extension currently blocks transfer
	//   operations.
	// - `domain_not_exists`: No registration record exists for the domain.
	// - `domain_on_cloudflare`: Cloudflare already serves as the domain's registrar.
	// - `domain_locked`: Losing registrar reports transfer-prohibited lock status.
	// - `registry_status`: Registry status currently blocks transfer (for example,
	//   pending transfer or deletion state).
	// - `domain_outside_transfer_window`: Domain is within a transfer wait window (for
	//   example, recently registered).
	// - `domain_max_term`: Completing transfer would exceed the registry maximum term.
	// - `invalid_auth_code`: The provided auth code is incorrect.
	// - `invalid_auth_code_format`: Auth code fails Base64 validation.
	// - `dnssec_enabled`: DNSSEC is enabled. It must be disabled before transfer.
	// - `zone_not_found`: The target account lacks a Cloudflare zone for the domain.
	// - `zone_status_invalid`: The Cloudflare zone cannot transfer in its current
	//   state.
	// - `invalid_zone_plan`: The zone plan fails transfer requirements.
	// - `domain_unsupported`: This endpoint rejects the domain name format.
	Code RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode `json:"code" api:"required"`
	JSON registrarTransferCheckResponseDomainsNonTransferableResultReasonJSON  `json:"-"`
}

// registrarTransferCheckResponseDomainsNonTransferableResultReasonJSON contains
// the JSON metadata for the struct
// [RegistrarTransferCheckResponseDomainsNonTransferableResultReason]
type registrarTransferCheckResponseDomainsNonTransferableResultReasonJSON struct {
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarTransferCheckResponseDomainsNonTransferableResultReason) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarTransferCheckResponseDomainsNonTransferableResultReasonJSON) RawJSON() string {
	return r.raw
}

// Transfer eligibility reason code.
//
//   - `extension_not_supported_via_api`: This API excludes the extension; dashboard
//     flows support it.
//   - `extension_not_supported`: Cloudflare Registrar excludes the extension.
//   - `domain_premium`: This API currently excludes premium transfers.
//   - `extension_disallows_transfer`: Extension currently blocks transfer
//     operations.
//   - `domain_not_exists`: No registration record exists for the domain.
//   - `domain_on_cloudflare`: Cloudflare already serves as the domain's registrar.
//   - `domain_locked`: Losing registrar reports transfer-prohibited lock status.
//   - `registry_status`: Registry status currently blocks transfer (for example,
//     pending transfer or deletion state).
//   - `domain_outside_transfer_window`: Domain is within a transfer wait window (for
//     example, recently registered).
//   - `domain_max_term`: Completing transfer would exceed the registry maximum term.
//   - `invalid_auth_code`: The provided auth code is incorrect.
//   - `invalid_auth_code_format`: Auth code fails Base64 validation.
//   - `dnssec_enabled`: DNSSEC is enabled. It must be disabled before transfer.
//   - `zone_not_found`: The target account lacks a Cloudflare zone for the domain.
//   - `zone_status_invalid`: The Cloudflare zone cannot transfer in its current
//     state.
//   - `invalid_zone_plan`: The zone plan fails transfer requirements.
//   - `domain_unsupported`: This endpoint rejects the domain name format.
type RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode string

const (
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeExtensionNotSupportedViaAPI RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "extension_not_supported_via_api"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeExtensionNotSupported       RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "extension_not_supported"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainPremium               RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "domain_premium"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeExtensionDisallowsTransfer  RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "extension_disallows_transfer"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainNotExists             RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "domain_not_exists"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainOnCloudflare          RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "domain_on_cloudflare"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainLocked                RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "domain_locked"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeRegistryStatus              RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "registry_status"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainOutsideTransferWindow RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "domain_outside_transfer_window"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainMaxTerm               RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "domain_max_term"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeInvalidAuthCode             RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "invalid_auth_code"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeInvalidAuthCodeFormat       RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "invalid_auth_code_format"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDNSSECEnabled               RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "dnssec_enabled"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeZoneNotFound                RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "zone_not_found"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeZoneStatusInvalid           RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "zone_status_invalid"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeInvalidZonePlan             RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "invalid_zone_plan"
	RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainUnsupported           RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode = "domain_unsupported"
)

func (r RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCode) IsKnown() bool {
	switch r {
	case RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeExtensionNotSupportedViaAPI, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeExtensionNotSupported, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainPremium, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeExtensionDisallowsTransfer, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainNotExists, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainOnCloudflare, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainLocked, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeRegistryStatus, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainOutsideTransferWindow, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainMaxTerm, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeInvalidAuthCode, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeInvalidAuthCodeFormat, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDNSSECEnabled, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeZoneNotFound, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeZoneStatusInvalid, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeInvalidZonePlan, RegistrarTransferCheckResponseDomainsNonTransferableResultReasonsCodeDomainUnsupported:
		return true
	}
	return false
}

type RegistrarTransferCheckResponseDomainsTransferable bool

const (
	RegistrarTransferCheckResponseDomainsTransferableTrue  RegistrarTransferCheckResponseDomainsTransferable = true
	RegistrarTransferCheckResponseDomainsTransferableFalse RegistrarTransferCheckResponseDomainsTransferable = false
)

func (r RegistrarTransferCheckResponseDomainsTransferable) IsKnown() bool {
	switch r {
	case RegistrarTransferCheckResponseDomainsTransferableTrue, RegistrarTransferCheckResponseDomainsTransferableFalse:
		return true
	}
	return false
}

type RegistrarCheckParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// List of fully qualified domain names (FQDNs) to check for availability. Each
	// domain must include the extension.
	//
	// - Minimum: 1 domain.
	// - Maximum: 20 domains per request.
	// - The response returns domains on unsupported extensions with
	//   `registrable: false` and a `reason` field.
	// - The response may omit malformed domain names (e.g., names missing an
	//   extension).
	Domains param.Field[[]string] `json:"domains" api:"required"`
}

func (r RegistrarCheckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegistrarCheckResponseEnvelope struct {
	Errors   []RegistrarCheckResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RegistrarCheckResponseEnvelopeMessages `json:"messages" api:"required"`
	// Contains the availability check results.
	Result RegistrarCheckResponse `json:"result" api:"required"`
	// Whether the API call was successful.
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

type RegistrarCheckResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrarCheckResponseEnvelopeErrorsSource `json:"source"`
	JSON   registrarCheckResponseEnvelopeErrorsJSON   `json:"-"`
}

// registrarCheckResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RegistrarCheckResponseEnvelopeErrors]
type registrarCheckResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarCheckResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarCheckResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrarCheckResponseEnvelopeErrorsSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                         `json:"pointer" api:"required"`
	JSON    registrarCheckResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// registrarCheckResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RegistrarCheckResponseEnvelopeErrorsSource]
type registrarCheckResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarCheckResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarCheckResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegistrarCheckResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrarCheckResponseEnvelopeMessagesSource `json:"source"`
	JSON   registrarCheckResponseEnvelopeMessagesJSON   `json:"-"`
}

// registrarCheckResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RegistrarCheckResponseEnvelopeMessages]
type registrarCheckResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarCheckResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarCheckResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrarCheckResponseEnvelopeMessagesSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                           `json:"pointer" api:"required"`
	JSON    registrarCheckResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// registrarCheckResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RegistrarCheckResponseEnvelopeMessagesSource]
type registrarCheckResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarCheckResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarCheckResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The search term to find domain suggestions. Accepts keywords, phrases, or full
	// domain names.
	//
	// - Phrases: "coffee shop" returns coffeeshop.com, mycoffeeshop.net, etc.
	// - Domain names: "example.com" returns example.com and variations across
	//   extensions
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
	Errors   []RegistrarSearchResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RegistrarSearchResponseEnvelopeMessages `json:"messages" api:"required"`
	// Contains the search results.
	Result RegistrarSearchResponse `json:"result" api:"required"`
	// Whether the API call was successful.
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

type RegistrarSearchResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrarSearchResponseEnvelopeErrorsSource `json:"source"`
	JSON   registrarSearchResponseEnvelopeErrorsJSON   `json:"-"`
}

// registrarSearchResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RegistrarSearchResponseEnvelopeErrors]
type registrarSearchResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSearchResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSearchResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrarSearchResponseEnvelopeErrorsSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    registrarSearchResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// registrarSearchResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RegistrarSearchResponseEnvelopeErrorsSource]
type registrarSearchResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSearchResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSearchResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegistrarSearchResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrarSearchResponseEnvelopeMessagesSource `json:"source"`
	JSON   registrarSearchResponseEnvelopeMessagesJSON   `json:"-"`
}

// registrarSearchResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RegistrarSearchResponseEnvelopeMessages]
type registrarSearchResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSearchResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSearchResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrarSearchResponseEnvelopeMessagesSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    registrarSearchResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// registrarSearchResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RegistrarSearchResponseEnvelopeMessagesSource]
type registrarSearchResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSearchResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSearchResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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

type RegistrarTransferCheckParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// List of domain objects to evaluate for transfer eligibility.
	Domains param.Field[[]RegistrarTransferCheckParamsDomain] `json:"domains" api:"required"`
}

func (r RegistrarTransferCheckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegistrarTransferCheckParamsDomain struct {
	// Fully qualified domain name (FQDN) to check for transfer eligibility.
	DomainName param.Field[string] `json:"domain_name" api:"required"`
	// Base64-encoded auth/EPP code from the current registrar. Required for most TLDs.
	// `.uk` namespaces do not use auth codes.
	AuthCode param.Field[string] `json:"auth_code" format:"byte"`
}

func (r RegistrarTransferCheckParamsDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegistrarTransferCheckResponseEnvelope struct {
	Errors   []RegistrarTransferCheckResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RegistrarTransferCheckResponseEnvelopeMessages `json:"messages" api:"required"`
	// Contains the transfer eligibility results.
	Result RegistrarTransferCheckResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success RegistrarTransferCheckResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    registrarTransferCheckResponseEnvelopeJSON    `json:"-"`
}

// registrarTransferCheckResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegistrarTransferCheckResponseEnvelope]
type registrarTransferCheckResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarTransferCheckResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarTransferCheckResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegistrarTransferCheckResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrarTransferCheckResponseEnvelopeErrorsSource `json:"source"`
	JSON   registrarTransferCheckResponseEnvelopeErrorsJSON   `json:"-"`
}

// registrarTransferCheckResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RegistrarTransferCheckResponseEnvelopeErrors]
type registrarTransferCheckResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarTransferCheckResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarTransferCheckResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrarTransferCheckResponseEnvelopeErrorsSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                                 `json:"pointer" api:"required"`
	JSON    registrarTransferCheckResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// registrarTransferCheckResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [RegistrarTransferCheckResponseEnvelopeErrorsSource]
type registrarTransferCheckResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarTransferCheckResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarTransferCheckResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegistrarTransferCheckResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrarTransferCheckResponseEnvelopeMessagesSource `json:"source"`
	JSON   registrarTransferCheckResponseEnvelopeMessagesJSON   `json:"-"`
}

// registrarTransferCheckResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RegistrarTransferCheckResponseEnvelopeMessages]
type registrarTransferCheckResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarTransferCheckResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarTransferCheckResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrarTransferCheckResponseEnvelopeMessagesSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                                   `json:"pointer" api:"required"`
	JSON    registrarTransferCheckResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// registrarTransferCheckResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RegistrarTransferCheckResponseEnvelopeMessagesSource]
type registrarTransferCheckResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarTransferCheckResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarTransferCheckResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RegistrarTransferCheckResponseEnvelopeSuccess bool

const (
	RegistrarTransferCheckResponseEnvelopeSuccessTrue RegistrarTransferCheckResponseEnvelopeSuccess = true
)

func (r RegistrarTransferCheckResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegistrarTransferCheckResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
