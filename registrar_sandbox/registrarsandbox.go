// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_sandbox

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// Use the Registrar Sandbox API to test domain search, availability checks,
// registration, and domain management flows without buying real domains.
//
// **This API is a test environment for the production Registrar API.**
//
// ## Prerequisites
//
// Before using this API, make sure you have:
//
//  1. **Cloudflare account** — the caller must have a valid Cloudflare account.
//  2. **API authentication** — create an API token with Registrar Sandbox
//     permissions.
//
// ## How the Sandbox API differs from the production Registrar API
//
// Because the Sandbox API is intended for testing, it behaves differently from the
// production Registrar API in a few important ways:
//
//  1. **No billing** — you will not be charged real money for purchasing a domain.
//  2. **No real domains** — purchased domains are test records and will not be
//     reachable on the Internet.
//  3. **No DNS zones** — purchasing a domain does not create a zone resource.
//  4. **No Registration Express Mode** — you must provide full contact data.
//
// Sandbox purchases are still persisted. If you purchase a domain in the sandbox,
// that domain will not be available for others to purchase in the sandbox.
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
// The Sandbox API currently supports programmatic registration for these
// extensions:
//
// `com`, `net`
//
// The production Registrar API supports 40+ extensions.
//
// Cloudflare Registrar supports 400+ extensions in the dashboard. Extensions not
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
//     currently supported by this API. The Sandbox API currently supports only
//     `com` and `net`, which do not have premium registrations, but clients should
//     still handle this response for consistency with the production Registrar API.
//     Surface the premium pricing to the user, but do not proceed to
//     `POST /registrations` for that domain.
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
// RegistrarSandboxService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistrarSandboxService] method instead.
type RegistrarSandboxService struct {
	Options            []option.RequestOption
	Registrations      *RegistrationService
	RegistrationStatus *RegistrationStatusService
	UpdateStatus       *UpdateStatusService
	Extensions         *ExtensionService
}

// NewRegistrarSandboxService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegistrarSandboxService(opts ...option.RequestOption) (r *RegistrarSandboxService) {
	r = &RegistrarSandboxService{}
	r.Options = opts
	r.Registrations = NewRegistrationService(opts...)
	r.RegistrationStatus = NewRegistrationStatusService(opts...)
	r.UpdateStatus = NewUpdateStatusService(opts...)
	r.Extensions = NewExtensionService(opts...)
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
func (r *RegistrarSandboxService) Check(ctx context.Context, params RegistrarSandboxCheckParams, opts ...option.RequestOption) (res *RegistrarSandboxCheckResponse, err error) {
	var env RegistrarSandboxCheckResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar-sandbox/domain-check", params.AccountID)
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
func (r *RegistrarSandboxService) Search(ctx context.Context, params RegistrarSandboxSearchParams, opts ...option.RequestOption) (res *RegistrarSandboxSearchResponse, err error) {
	var env RegistrarSandboxSearchResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar-sandbox/domain-search", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Contains the availability check results.
type RegistrarSandboxCheckResponse struct {
	// Array of domain availability results. Results for unsupported extensions contain
	// `registrable: false` and a `reason` field. The response may omit malformed
	// domain names.
	Domains []RegistrarSandboxCheckResponseDomain `json:"domains" api:"required"`
	JSON    registrarSandboxCheckResponseJSON     `json:"-"`
}

// registrarSandboxCheckResponseJSON contains the JSON metadata for the struct
// [RegistrarSandboxCheckResponse]
type registrarSandboxCheckResponseJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxCheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxCheckResponseJSON) RawJSON() string {
	return r.raw
}

// Describes a single authoritative domain availability result from the Check
// endpoint. Check results reflect current registry status; use them immediately
// before registration.
type RegistrarSandboxCheckResponseDomain struct {
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
	// Provides annual pricing information for a registrable domain. This object
	// appears only when `registrable` is `true`. The API returns all per-year prices
	// as strings to preserve decimal precision.
	//
	// `registration_cost` and `renewal_cost` frequently have the same value, but may
	// differ, especially when registries set different premium rates for initial
	// registration and renewal. For a multi-year registration (e.g., 4 years),
	// `registration_cost` applies to the first year and `renewal_cost` applies to each
	// subsequent year. The values reflect the current registry rate, which may change
	// over time. Search and Check may surface premium pricing, but this API currently
	// supports standard registrations only.
	Pricing RegistrarSandboxCheckResponseDomainsPricing `json:"pricing"`
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
	Reason RegistrarSandboxCheckResponseDomainsReason `json:"reason"`
	// The pricing tier for this domain. A `registrable` value of `true` always
	// includes this field, which defaults to `standard` for most domains. A
	// `registrable` value of `false` may omit it.
	//
	// - `standard`: Standard registry pricing.
	// - `premium`: Premium domain with higher pricing from the registry.
	Tier RegistrarSandboxCheckResponseDomainsTier `json:"tier"`
	JSON registrarSandboxCheckResponseDomainJSON  `json:"-"`
}

// registrarSandboxCheckResponseDomainJSON contains the JSON metadata for the
// struct [RegistrarSandboxCheckResponseDomain]
type registrarSandboxCheckResponseDomainJSON struct {
	Name        apijson.Field
	Registrable apijson.Field
	Pricing     apijson.Field
	Reason      apijson.Field
	Tier        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxCheckResponseDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxCheckResponseDomainJSON) RawJSON() string {
	return r.raw
}

// Provides annual pricing information for a registrable domain. This object
// appears only when `registrable` is `true`. The API returns all per-year prices
// as strings to preserve decimal precision.
//
// `registration_cost` and `renewal_cost` frequently have the same value, but may
// differ, especially when registries set different premium rates for initial
// registration and renewal. For a multi-year registration (e.g., 4 years),
// `registration_cost` applies to the first year and `renewal_cost` applies to each
// subsequent year. The values reflect the current registry rate, which may change
// over time. Search and Check may surface premium pricing, but this API currently
// supports standard registrations only.
type RegistrarSandboxCheckResponseDomainsPricing struct {
	// ISO-4217 currency code for the prices (e.g., "USD", "EUR", "GBP").
	Currency string `json:"currency" api:"required"`
	// The first-year cost to register this domain. For premium domains
	// (`tier: premium`), the registry sets this price, which may significantly exceed
	// standard pricing. For multi-year registrations, this cost applies to the first
	// year only; `renewal_cost` applies to subsequent years.
	RegistrationCost string `json:"registration_cost" api:"required"`
	// Per-year renewal cost for this domain. Applied to each year beyond the first
	// year of a multi-year registration, and to each annual auto-renewal thereafter.
	// May differ from `registration_cost`, especially for premium domains where
	// initial registration often costs more than renewals.
	RenewalCost string                                          `json:"renewal_cost" api:"required"`
	JSON        registrarSandboxCheckResponseDomainsPricingJSON `json:"-"`
}

// registrarSandboxCheckResponseDomainsPricingJSON contains the JSON metadata for
// the struct [RegistrarSandboxCheckResponseDomainsPricing]
type registrarSandboxCheckResponseDomainsPricingJSON struct {
	Currency         apijson.Field
	RegistrationCost apijson.Field
	RenewalCost      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegistrarSandboxCheckResponseDomainsPricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxCheckResponseDomainsPricingJSON) RawJSON() string {
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
type RegistrarSandboxCheckResponseDomainsReason string

const (
	RegistrarSandboxCheckResponseDomainsReasonExtensionNotSupportedViaAPI    RegistrarSandboxCheckResponseDomainsReason = "extension_not_supported_via_api"
	RegistrarSandboxCheckResponseDomainsReasonExtensionNotSupported          RegistrarSandboxCheckResponseDomainsReason = "extension_not_supported"
	RegistrarSandboxCheckResponseDomainsReasonExtensionDisallowsRegistration RegistrarSandboxCheckResponseDomainsReason = "extension_disallows_registration"
	RegistrarSandboxCheckResponseDomainsReasonDomainPremium                  RegistrarSandboxCheckResponseDomainsReason = "domain_premium"
	RegistrarSandboxCheckResponseDomainsReasonDomainUnavailable              RegistrarSandboxCheckResponseDomainsReason = "domain_unavailable"
)

func (r RegistrarSandboxCheckResponseDomainsReason) IsKnown() bool {
	switch r {
	case RegistrarSandboxCheckResponseDomainsReasonExtensionNotSupportedViaAPI, RegistrarSandboxCheckResponseDomainsReasonExtensionNotSupported, RegistrarSandboxCheckResponseDomainsReasonExtensionDisallowsRegistration, RegistrarSandboxCheckResponseDomainsReasonDomainPremium, RegistrarSandboxCheckResponseDomainsReasonDomainUnavailable:
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
type RegistrarSandboxCheckResponseDomainsTier string

const (
	RegistrarSandboxCheckResponseDomainsTierStandard RegistrarSandboxCheckResponseDomainsTier = "standard"
	RegistrarSandboxCheckResponseDomainsTierPremium  RegistrarSandboxCheckResponseDomainsTier = "premium"
)

func (r RegistrarSandboxCheckResponseDomainsTier) IsKnown() bool {
	switch r {
	case RegistrarSandboxCheckResponseDomainsTierStandard, RegistrarSandboxCheckResponseDomainsTierPremium:
		return true
	}
	return false
}

// Contains the search results.
type RegistrarSandboxSearchResponse struct {
	// Lists domain suggestions in relevance order. An empty array indicates that the
	// search criteria matched zero domains.
	Domains []RegistrarSandboxSearchResponseDomain `json:"domains" api:"required"`
	JSON    registrarSandboxSearchResponseJSON     `json:"-"`
}

// registrarSandboxSearchResponseJSON contains the JSON metadata for the struct
// [RegistrarSandboxSearchResponse]
type registrarSandboxSearchResponseJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxSearchResponseJSON) RawJSON() string {
	return r.raw
}

// Describes a single domain suggestion from the Search endpoint. Search results
// use non-authoritative data that may come from a cache. Use POST /domain-check to
// confirm real-time availability and pricing before registration.
type RegistrarSandboxSearchResponseDomain struct {
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
	// Provides annual pricing information for a registrable domain. This object
	// appears only when `registrable` is `true`. The API returns all per-year prices
	// as strings to preserve decimal precision.
	//
	// `registration_cost` and `renewal_cost` frequently have the same value, but may
	// differ, especially when registries set different premium rates for initial
	// registration and renewal. For a multi-year registration (e.g., 4 years),
	// `registration_cost` applies to the first year and `renewal_cost` applies to each
	// subsequent year. The values reflect the current registry rate, which may change
	// over time. Search and Check may surface premium pricing, but this API currently
	// supports standard registrations only.
	Pricing RegistrarSandboxSearchResponseDomainsPricing `json:"pricing"`
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
	Reason RegistrarSandboxSearchResponseDomainsReason `json:"reason"`
	// The pricing tier for this domain. A `registrable` value of `true` always
	// includes this field, which defaults to `standard` for most domains. A
	// `registrable` value of `false` may omit it.
	//
	// - `standard`: Standard registry pricing.
	// - `premium`: Premium domain with higher pricing from the registry.
	Tier RegistrarSandboxSearchResponseDomainsTier `json:"tier"`
	JSON registrarSandboxSearchResponseDomainJSON  `json:"-"`
}

// registrarSandboxSearchResponseDomainJSON contains the JSON metadata for the
// struct [RegistrarSandboxSearchResponseDomain]
type registrarSandboxSearchResponseDomainJSON struct {
	Name        apijson.Field
	Registrable apijson.Field
	Pricing     apijson.Field
	Reason      apijson.Field
	Tier        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxSearchResponseDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxSearchResponseDomainJSON) RawJSON() string {
	return r.raw
}

// Provides annual pricing information for a registrable domain. This object
// appears only when `registrable` is `true`. The API returns all per-year prices
// as strings to preserve decimal precision.
//
// `registration_cost` and `renewal_cost` frequently have the same value, but may
// differ, especially when registries set different premium rates for initial
// registration and renewal. For a multi-year registration (e.g., 4 years),
// `registration_cost` applies to the first year and `renewal_cost` applies to each
// subsequent year. The values reflect the current registry rate, which may change
// over time. Search and Check may surface premium pricing, but this API currently
// supports standard registrations only.
type RegistrarSandboxSearchResponseDomainsPricing struct {
	// ISO-4217 currency code for the prices (e.g., "USD", "EUR", "GBP").
	Currency string `json:"currency" api:"required"`
	// The first-year cost to register this domain. For premium domains
	// (`tier: premium`), the registry sets this price, which may significantly exceed
	// standard pricing. For multi-year registrations, this cost applies to the first
	// year only; `renewal_cost` applies to subsequent years.
	RegistrationCost string `json:"registration_cost" api:"required"`
	// Per-year renewal cost for this domain. Applied to each year beyond the first
	// year of a multi-year registration, and to each annual auto-renewal thereafter.
	// May differ from `registration_cost`, especially for premium domains where
	// initial registration often costs more than renewals.
	RenewalCost string                                           `json:"renewal_cost" api:"required"`
	JSON        registrarSandboxSearchResponseDomainsPricingJSON `json:"-"`
}

// registrarSandboxSearchResponseDomainsPricingJSON contains the JSON metadata for
// the struct [RegistrarSandboxSearchResponseDomainsPricing]
type registrarSandboxSearchResponseDomainsPricingJSON struct {
	Currency         apijson.Field
	RegistrationCost apijson.Field
	RenewalCost      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegistrarSandboxSearchResponseDomainsPricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxSearchResponseDomainsPricingJSON) RawJSON() string {
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
type RegistrarSandboxSearchResponseDomainsReason string

const (
	RegistrarSandboxSearchResponseDomainsReasonExtensionNotSupportedViaAPI    RegistrarSandboxSearchResponseDomainsReason = "extension_not_supported_via_api"
	RegistrarSandboxSearchResponseDomainsReasonExtensionNotSupported          RegistrarSandboxSearchResponseDomainsReason = "extension_not_supported"
	RegistrarSandboxSearchResponseDomainsReasonExtensionDisallowsRegistration RegistrarSandboxSearchResponseDomainsReason = "extension_disallows_registration"
	RegistrarSandboxSearchResponseDomainsReasonDomainPremium                  RegistrarSandboxSearchResponseDomainsReason = "domain_premium"
	RegistrarSandboxSearchResponseDomainsReasonDomainUnavailable              RegistrarSandboxSearchResponseDomainsReason = "domain_unavailable"
)

func (r RegistrarSandboxSearchResponseDomainsReason) IsKnown() bool {
	switch r {
	case RegistrarSandboxSearchResponseDomainsReasonExtensionNotSupportedViaAPI, RegistrarSandboxSearchResponseDomainsReasonExtensionNotSupported, RegistrarSandboxSearchResponseDomainsReasonExtensionDisallowsRegistration, RegistrarSandboxSearchResponseDomainsReasonDomainPremium, RegistrarSandboxSearchResponseDomainsReasonDomainUnavailable:
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
type RegistrarSandboxSearchResponseDomainsTier string

const (
	RegistrarSandboxSearchResponseDomainsTierStandard RegistrarSandboxSearchResponseDomainsTier = "standard"
	RegistrarSandboxSearchResponseDomainsTierPremium  RegistrarSandboxSearchResponseDomainsTier = "premium"
)

func (r RegistrarSandboxSearchResponseDomainsTier) IsKnown() bool {
	switch r {
	case RegistrarSandboxSearchResponseDomainsTierStandard, RegistrarSandboxSearchResponseDomainsTierPremium:
		return true
	}
	return false
}

type RegistrarSandboxCheckParams struct {
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

func (r RegistrarSandboxCheckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegistrarSandboxCheckResponseEnvelope struct {
	Errors   []RegistrarSandboxCheckResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RegistrarSandboxCheckResponseEnvelopeMessages `json:"messages" api:"required"`
	// Contains the availability check results.
	Result RegistrarSandboxCheckResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success RegistrarSandboxCheckResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    registrarSandboxCheckResponseEnvelopeJSON    `json:"-"`
}

// registrarSandboxCheckResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegistrarSandboxCheckResponseEnvelope]
type registrarSandboxCheckResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxCheckResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxCheckResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegistrarSandboxCheckResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrarSandboxCheckResponseEnvelopeErrorsSource `json:"source"`
	JSON   registrarSandboxCheckResponseEnvelopeErrorsJSON   `json:"-"`
}

// registrarSandboxCheckResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RegistrarSandboxCheckResponseEnvelopeErrors]
type registrarSandboxCheckResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxCheckResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxCheckResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrarSandboxCheckResponseEnvelopeErrorsSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                                `json:"pointer" api:"required"`
	JSON    registrarSandboxCheckResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// registrarSandboxCheckResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RegistrarSandboxCheckResponseEnvelopeErrorsSource]
type registrarSandboxCheckResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxCheckResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxCheckResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegistrarSandboxCheckResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrarSandboxCheckResponseEnvelopeMessagesSource `json:"source"`
	JSON   registrarSandboxCheckResponseEnvelopeMessagesJSON   `json:"-"`
}

// registrarSandboxCheckResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RegistrarSandboxCheckResponseEnvelopeMessages]
type registrarSandboxCheckResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxCheckResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxCheckResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrarSandboxCheckResponseEnvelopeMessagesSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                                  `json:"pointer" api:"required"`
	JSON    registrarSandboxCheckResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// registrarSandboxCheckResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RegistrarSandboxCheckResponseEnvelopeMessagesSource]
type registrarSandboxCheckResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxCheckResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxCheckResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RegistrarSandboxCheckResponseEnvelopeSuccess bool

const (
	RegistrarSandboxCheckResponseEnvelopeSuccessTrue RegistrarSandboxCheckResponseEnvelopeSuccess = true
)

func (r RegistrarSandboxCheckResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegistrarSandboxCheckResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RegistrarSandboxSearchParams struct {
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

// URLQuery serializes [RegistrarSandboxSearchParams]'s query parameters as
// `url.Values`.
func (r RegistrarSandboxSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RegistrarSandboxSearchResponseEnvelope struct {
	Errors   []RegistrarSandboxSearchResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RegistrarSandboxSearchResponseEnvelopeMessages `json:"messages" api:"required"`
	// Contains the search results.
	Result RegistrarSandboxSearchResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success RegistrarSandboxSearchResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    registrarSandboxSearchResponseEnvelopeJSON    `json:"-"`
}

// registrarSandboxSearchResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegistrarSandboxSearchResponseEnvelope]
type registrarSandboxSearchResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxSearchResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxSearchResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegistrarSandboxSearchResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrarSandboxSearchResponseEnvelopeErrorsSource `json:"source"`
	JSON   registrarSandboxSearchResponseEnvelopeErrorsJSON   `json:"-"`
}

// registrarSandboxSearchResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RegistrarSandboxSearchResponseEnvelopeErrors]
type registrarSandboxSearchResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxSearchResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxSearchResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrarSandboxSearchResponseEnvelopeErrorsSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                                 `json:"pointer" api:"required"`
	JSON    registrarSandboxSearchResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// registrarSandboxSearchResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [RegistrarSandboxSearchResponseEnvelopeErrorsSource]
type registrarSandboxSearchResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxSearchResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxSearchResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegistrarSandboxSearchResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrarSandboxSearchResponseEnvelopeMessagesSource `json:"source"`
	JSON   registrarSandboxSearchResponseEnvelopeMessagesJSON   `json:"-"`
}

// registrarSandboxSearchResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RegistrarSandboxSearchResponseEnvelopeMessages]
type registrarSandboxSearchResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxSearchResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxSearchResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrarSandboxSearchResponseEnvelopeMessagesSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                                   `json:"pointer" api:"required"`
	JSON    registrarSandboxSearchResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// registrarSandboxSearchResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RegistrarSandboxSearchResponseEnvelopeMessagesSource]
type registrarSandboxSearchResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarSandboxSearchResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarSandboxSearchResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RegistrarSandboxSearchResponseEnvelopeSuccess bool

const (
	RegistrarSandboxSearchResponseEnvelopeSuccessTrue RegistrarSandboxSearchResponseEnvelopeSuccess = true
)

func (r RegistrarSandboxSearchResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegistrarSandboxSearchResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
