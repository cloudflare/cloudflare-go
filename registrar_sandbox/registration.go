// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_sandbox

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
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

// Starts a domain registration workflow.
//
// ### Prerequisites
//
//   - The account must not already be at the maximum supported domain limit. A
//     single account may own up to 500 domains in total across registrations created
//     through either the dashboard or this API.
//   - The domain must be on a supported extension for programmatic registration.
//   - Use `POST /domain-check` immediately before calling this endpoint to confirm
//     real-time availability and pricing.
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
func (r *RegistrationService) New(ctx context.Context, params RegistrationNewParams, opts ...option.RequestOption) (res *RegistrationNewResponse, err error) {
	var env RegistrationNewResponseEnvelope
	if params.Prefer.Present {
		opts = append(opts, option.WithHeader("Prefer", fmt.Sprintf("%v", params.Prefer)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar-sandbox/registrations", params.AccountID)
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
func (r *RegistrationService) List(ctx context.Context, params RegistrationListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[RegistrationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar-sandbox/registrations", params.AccountID)
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
func (r *RegistrationService) ListAutoPaging(ctx context.Context, params RegistrationListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[RegistrationListResponse] {
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
func (r *RegistrationService) Edit(ctx context.Context, domainName string, params RegistrationEditParams, opts ...option.RequestOption) (res *RegistrationEditResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/registrar-sandbox/registrations/%s", params.AccountID, domainName)
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
func (r *RegistrationService) Get(ctx context.Context, domainName string, query RegistrationGetParams, opts ...option.RequestOption) (res *RegistrationGetResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/registrar-sandbox/registrations/%s", query.AccountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Status of an async registration workflow.
type RegistrationNewResponse struct {
	// Indicates whether the workflow reached a terminal state. A `succeeded` or
	// `failed` state returns `true`; `pending`, `in_progress`, `action_required`, and
	// `blocked` return `false`.
	Completed bool                         `json:"completed" api:"required"`
	CreatedAt time.Time                    `json:"created_at" api:"required" format:"date-time"`
	Links     RegistrationNewResponseLinks `json:"links" api:"required"`
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
	State     RegistrationNewResponseState `json:"state" api:"required"`
	UpdatedAt time.Time                    `json:"updated_at" api:"required" format:"date-time"`
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
	Error RegistrationNewResponseError `json:"error" api:"nullable"`
	JSON  registrationNewResponseJSON  `json:"-"`
}

// registrationNewResponseJSON contains the JSON metadata for the struct
// [RegistrationNewResponse]
type registrationNewResponseJSON struct {
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

func (r *RegistrationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationNewResponseJSON) RawJSON() string {
	return r.raw
}

type RegistrationNewResponseLinks struct {
	// URL to this status resource.
	Self string `json:"self" api:"required"`
	// URL to the domain resource.
	Resource string                           `json:"resource"`
	JSON     registrationNewResponseLinksJSON `json:"-"`
}

// registrationNewResponseLinksJSON contains the JSON metadata for the struct
// [RegistrationNewResponseLinks]
type registrationNewResponseLinksJSON struct {
	Self        apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationNewResponseLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationNewResponseLinksJSON) RawJSON() string {
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
type RegistrationNewResponseState string

const (
	RegistrationNewResponseStatePending        RegistrationNewResponseState = "pending"
	RegistrationNewResponseStateInProgress     RegistrationNewResponseState = "in_progress"
	RegistrationNewResponseStateActionRequired RegistrationNewResponseState = "action_required"
	RegistrationNewResponseStateBlocked        RegistrationNewResponseState = "blocked"
	RegistrationNewResponseStateSucceeded      RegistrationNewResponseState = "succeeded"
	RegistrationNewResponseStateFailed         RegistrationNewResponseState = "failed"
)

func (r RegistrationNewResponseState) IsKnown() bool {
	switch r {
	case RegistrationNewResponseStatePending, RegistrationNewResponseStateInProgress, RegistrationNewResponseStateActionRequired, RegistrationNewResponseStateBlocked, RegistrationNewResponseStateSucceeded, RegistrationNewResponseStateFailed:
		return true
	}
	return false
}

// Provides error details when a workflow reaches the `failed` state. The workflow
// type (registration, update, etc.) and underlying registry response determine the
// specific codes and messages. Workflow error codes differ from immediate HTTP
// error `errors[].code` values in non-2xx responses. Surface `error.message` to
// the user for context.
type RegistrationNewResponseError struct {
	// Machine-readable error code identifying the failure reason.
	Code string `json:"code" api:"required"`
	// Human-readable explanation of the failure. May include registry-specific
	// details.
	Message string                           `json:"message" api:"required"`
	JSON    registrationNewResponseErrorJSON `json:"-"`
}

// registrationNewResponseErrorJSON contains the JSON metadata for the struct
// [RegistrationNewResponseError]
type registrationNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

// A domain registration resource representing the current state of a registered
// domain.
type RegistrationListResponse struct {
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
	// only `registration_pending` may return null.
	ExpiresAt time.Time `json:"expires_at" api:"required,nullable" format:"date-time"`
	// Whether the domain is locked for transfer.
	Locked bool `json:"locked" api:"required"`
	// Current WHOIS privacy mode for the registration.
	PrivacyMode RegistrationListResponsePrivacyMode `json:"privacy_mode" api:"required"`
	// Current registration status.
	//
	// - `active`: The domain operates with an active registration.
	// - `registration_pending`: Registration remains in progress.
	// - `expired`: The domain registration expired.
	// - `suspended`: The registry suspended the domain.
	// - `redemption_period`: The domain entered the redemption grace period.
	// - `pending_delete`: The registry scheduled the domain for deletion.
	Status RegistrationListResponseStatus `json:"status" api:"required"`
	JSON   registrationListResponseJSON   `json:"-"`
}

// registrationListResponseJSON contains the JSON metadata for the struct
// [RegistrationListResponse]
type registrationListResponseJSON struct {
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

func (r *RegistrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationListResponseJSON) RawJSON() string {
	return r.raw
}

// Current WHOIS privacy mode for the registration.
type RegistrationListResponsePrivacyMode string

const (
	RegistrationListResponsePrivacyModeOff       RegistrationListResponsePrivacyMode = "off"
	RegistrationListResponsePrivacyModeRedaction RegistrationListResponsePrivacyMode = "redaction"
)

func (r RegistrationListResponsePrivacyMode) IsKnown() bool {
	switch r {
	case RegistrationListResponsePrivacyModeOff, RegistrationListResponsePrivacyModeRedaction:
		return true
	}
	return false
}

// Current registration status.
//
// - `active`: The domain operates with an active registration.
// - `registration_pending`: Registration remains in progress.
// - `expired`: The domain registration expired.
// - `suspended`: The registry suspended the domain.
// - `redemption_period`: The domain entered the redemption grace period.
// - `pending_delete`: The registry scheduled the domain for deletion.
type RegistrationListResponseStatus string

const (
	RegistrationListResponseStatusActive              RegistrationListResponseStatus = "active"
	RegistrationListResponseStatusRegistrationPending RegistrationListResponseStatus = "registration_pending"
	RegistrationListResponseStatusExpired             RegistrationListResponseStatus = "expired"
	RegistrationListResponseStatusSuspended           RegistrationListResponseStatus = "suspended"
	RegistrationListResponseStatusRedemptionPeriod    RegistrationListResponseStatus = "redemption_period"
	RegistrationListResponseStatusPendingDelete       RegistrationListResponseStatus = "pending_delete"
)

func (r RegistrationListResponseStatus) IsKnown() bool {
	switch r {
	case RegistrationListResponseStatusActive, RegistrationListResponseStatusRegistrationPending, RegistrationListResponseStatusExpired, RegistrationListResponseStatusSuspended, RegistrationListResponseStatusRedemptionPeriod, RegistrationListResponseStatusPendingDelete:
		return true
	}
	return false
}

// Status of an async registration workflow.
type RegistrationEditResponse struct {
	// Indicates whether the workflow reached a terminal state. A `succeeded` or
	// `failed` state returns `true`; `pending`, `in_progress`, `action_required`, and
	// `blocked` return `false`.
	Completed bool                          `json:"completed" api:"required"`
	CreatedAt time.Time                     `json:"created_at" api:"required" format:"date-time"`
	Links     RegistrationEditResponseLinks `json:"links" api:"required"`
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
	State     RegistrationEditResponseState `json:"state" api:"required"`
	UpdatedAt time.Time                     `json:"updated_at" api:"required" format:"date-time"`
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
	Error RegistrationEditResponseError `json:"error" api:"nullable"`
	JSON  registrationEditResponseJSON  `json:"-"`
}

// registrationEditResponseJSON contains the JSON metadata for the struct
// [RegistrationEditResponse]
type registrationEditResponseJSON struct {
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

func (r *RegistrationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationEditResponseJSON) RawJSON() string {
	return r.raw
}

type RegistrationEditResponseLinks struct {
	// URL to this status resource.
	Self string `json:"self" api:"required"`
	// URL to the domain resource.
	Resource string                            `json:"resource"`
	JSON     registrationEditResponseLinksJSON `json:"-"`
}

// registrationEditResponseLinksJSON contains the JSON metadata for the struct
// [RegistrationEditResponseLinks]
type registrationEditResponseLinksJSON struct {
	Self        apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationEditResponseLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationEditResponseLinksJSON) RawJSON() string {
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
type RegistrationEditResponseState string

const (
	RegistrationEditResponseStatePending        RegistrationEditResponseState = "pending"
	RegistrationEditResponseStateInProgress     RegistrationEditResponseState = "in_progress"
	RegistrationEditResponseStateActionRequired RegistrationEditResponseState = "action_required"
	RegistrationEditResponseStateBlocked        RegistrationEditResponseState = "blocked"
	RegistrationEditResponseStateSucceeded      RegistrationEditResponseState = "succeeded"
	RegistrationEditResponseStateFailed         RegistrationEditResponseState = "failed"
)

func (r RegistrationEditResponseState) IsKnown() bool {
	switch r {
	case RegistrationEditResponseStatePending, RegistrationEditResponseStateInProgress, RegistrationEditResponseStateActionRequired, RegistrationEditResponseStateBlocked, RegistrationEditResponseStateSucceeded, RegistrationEditResponseStateFailed:
		return true
	}
	return false
}

// Provides error details when a workflow reaches the `failed` state. The workflow
// type (registration, update, etc.) and underlying registry response determine the
// specific codes and messages. Workflow error codes differ from immediate HTTP
// error `errors[].code` values in non-2xx responses. Surface `error.message` to
// the user for context.
type RegistrationEditResponseError struct {
	// Machine-readable error code identifying the failure reason.
	Code string `json:"code" api:"required"`
	// Human-readable explanation of the failure. May include registry-specific
	// details.
	Message string                            `json:"message" api:"required"`
	JSON    registrationEditResponseErrorJSON `json:"-"`
}

// registrationEditResponseErrorJSON contains the JSON metadata for the struct
// [RegistrationEditResponseError]
type registrationEditResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationEditResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationEditResponseErrorJSON) RawJSON() string {
	return r.raw
}

// A domain registration resource representing the current state of a registered
// domain.
type RegistrationGetResponse struct {
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
	// only `registration_pending` may return null.
	ExpiresAt time.Time `json:"expires_at" api:"required,nullable" format:"date-time"`
	// Whether the domain is locked for transfer.
	Locked bool `json:"locked" api:"required"`
	// Current WHOIS privacy mode for the registration.
	PrivacyMode RegistrationGetResponsePrivacyMode `json:"privacy_mode" api:"required"`
	// Current registration status.
	//
	// - `active`: The domain operates with an active registration.
	// - `registration_pending`: Registration remains in progress.
	// - `expired`: The domain registration expired.
	// - `suspended`: The registry suspended the domain.
	// - `redemption_period`: The domain entered the redemption grace period.
	// - `pending_delete`: The registry scheduled the domain for deletion.
	Status RegistrationGetResponseStatus `json:"status" api:"required"`
	JSON   registrationGetResponseJSON   `json:"-"`
}

// registrationGetResponseJSON contains the JSON metadata for the struct
// [RegistrationGetResponse]
type registrationGetResponseJSON struct {
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

func (r *RegistrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Current WHOIS privacy mode for the registration.
type RegistrationGetResponsePrivacyMode string

const (
	RegistrationGetResponsePrivacyModeOff       RegistrationGetResponsePrivacyMode = "off"
	RegistrationGetResponsePrivacyModeRedaction RegistrationGetResponsePrivacyMode = "redaction"
)

func (r RegistrationGetResponsePrivacyMode) IsKnown() bool {
	switch r {
	case RegistrationGetResponsePrivacyModeOff, RegistrationGetResponsePrivacyModeRedaction:
		return true
	}
	return false
}

// Current registration status.
//
// - `active`: The domain operates with an active registration.
// - `registration_pending`: Registration remains in progress.
// - `expired`: The domain registration expired.
// - `suspended`: The registry suspended the domain.
// - `redemption_period`: The domain entered the redemption grace period.
// - `pending_delete`: The registry scheduled the domain for deletion.
type RegistrationGetResponseStatus string

const (
	RegistrationGetResponseStatusActive              RegistrationGetResponseStatus = "active"
	RegistrationGetResponseStatusRegistrationPending RegistrationGetResponseStatus = "registration_pending"
	RegistrationGetResponseStatusExpired             RegistrationGetResponseStatus = "expired"
	RegistrationGetResponseStatusSuspended           RegistrationGetResponseStatus = "suspended"
	RegistrationGetResponseStatusRedemptionPeriod    RegistrationGetResponseStatus = "redemption_period"
	RegistrationGetResponseStatusPendingDelete       RegistrationGetResponseStatus = "pending_delete"
)

func (r RegistrationGetResponseStatus) IsKnown() bool {
	switch r {
	case RegistrationGetResponseStatusActive, RegistrationGetResponseStatusRegistrationPending, RegistrationGetResponseStatusExpired, RegistrationGetResponseStatusSuspended, RegistrationGetResponseStatusRedemptionPeriod, RegistrationGetResponseStatusPendingDelete:
		return true
	}
	return false
}

type RegistrationNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Provides a fully qualified domain name (FQDN), including the extension (e.g.,
	// `example.com`, `mybrand.app`). The domain name uniquely identifies a
	// registration. Cloudflare permits only one registration per domain, making the
	// domain name a natural idempotency key for registration requests.
	DomainName param.Field[string] `json:"domain_name" api:"required"`
	// Provides user acknowledgements for a specific extension or premium registration
	// flow. The extension registration schema from the extension discovery endpoint
	// identifies the required keys.
	Acknowledgements param.Field[map[string]interface{}] `json:"acknowledgements"`
	// Enable or disable automatic renewal. Defaults to `false` if omitted. Setting
	// this field to `true` is an explicit opt-in authorizing Cloudflare to charge the
	// account's default payment method up to 30 days before domain expiry to renew the
	// domain automatically. Renewal pricing may change over time based on registry
	// pricing.
	AutoRenew param.Field[bool] `json:"auto_renew"`
	// Provides registry-specific contact extension values for the registrant.
	// `GET /accounts/{account_id}/registrar/extensions/{extension}` identifies the
	// required keys and allowed values for each extension in the
	// `registration_schema.properties.contact_extensions` object.
	//
	// Examples include `.us` nexus fields, `.uk` registrant type fields, and `.ca`
	// legal type fields. Omit this object when the extension's registration schema
	// excludes `contact_extensions`.
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
	Contacts param.Field[RegistrationNewParamsContacts] `json:"contacts"`
	// Sets the WHOIS privacy mode for the registration. Defaults to `redaction`.
	//
	// - `off`: Disables WHOIS privacy.
	// - `redaction`: Requests WHOIS redaction where the extension supports it. Some
	//   extensions exclude privacy and redaction.
	PrivacyMode param.Field[RegistrationNewParamsPrivacyMode] `json:"privacy_mode"`
	// Sets the registration term from 1 to 10 years. When omitted, this field defaults
	// to the registry's minimum registration period for the extension. Most extensions
	// require 1 year, while some require longer minimum terms (e.g., `.ai` requires 2
	// years).
	//
	// Each registry may also enforce its own maximum registration term. A request
	// above that maximum fails. When uncertain, omit this field to use the default.
	Years  param.Field[int64]  `json:"years"`
	Prefer param.Field[string] `header:"Prefer"`
}

func (r RegistrationNewParams) MarshalJSON() (data []byte, err error) {
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
type RegistrationNewParamsContacts struct {
	// Optional administrator contact. Accepted only when the extension schema includes
	// this role. When the registry requires an omitted contact, Cloudflare may derive
	// it from `contacts.registrant`.
	Administrator param.Field[RegistrationNewParamsContactsAdministrator] `json:"administrator"`
	// Optional billing contact. Accepted only when the extension schema includes this
	// role. When the registry requires an omitted contact, Cloudflare may derive it
	// from `contacts.registrant`.
	Billing param.Field[RegistrationNewParamsContactsBilling] `json:"billing"`
	// Optional registrant contact. If omitted, the account's default address book
	// entry is used instead.
	Registrant param.Field[RegistrationNewParamsContactsRegistrant] `json:"registrant"`
	// Optional technical contact. Accepted only when the extension schema includes
	// this role. When the registry requires an omitted contact, Cloudflare may derive
	// it from `contacts.registrant`.
	Technical param.Field[RegistrationNewParamsContactsTechnical] `json:"technical"`
}

func (r RegistrationNewParamsContacts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional administrator contact. Accepted only when the extension schema includes
// this role. When the registry requires an omitted contact, Cloudflare may derive
// it from `contacts.registrant`.
type RegistrationNewParamsContactsAdministrator struct {
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
	PostalInfo param.Field[RegistrationNewParamsContactsAdministratorPostalInfo] `json:"postal_info" api:"required"`
	// Fax number in E.164 format (e.g., `+1.5555555555`). Optional. Most registrations
	// do not require a fax number.
	Fax param.Field[string] `json:"fax"`
}

func (r RegistrationNewParamsContactsAdministrator) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Postal/mailing information for the contact. The `name` field is the complete
// contact name in one string. Some registries require a complete personal name,
// including a family or last name where applicable, but this API does not accept
// separate first-name and last-name fields for registration contacts.
type RegistrationNewParamsContactsAdministratorPostalInfo struct {
	// Physical mailing address for the registrant contact.
	Address param.Field[RegistrationNewParamsContactsAdministratorPostalInfoAddress] `json:"address" api:"required"`
	// Full legal name of the contact, including all required name components for an
	// individual or authorized representative. Some registries require a complete
	// personal name that includes a family or last name where applicable. Provide the
	// complete name in this single field, for example `Ada Lovelace`; do not send
	// separate first-name or last-name fields.
	Name param.Field[string] `json:"name" api:"required"`
	// Organization or company name. Optional for individual registrants.
	Organization param.Field[string] `json:"organization"`
}

func (r RegistrationNewParamsContactsAdministratorPostalInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Physical mailing address for the registrant contact.
type RegistrationNewParamsContactsAdministratorPostalInfoAddress struct {
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

func (r RegistrationNewParamsContactsAdministratorPostalInfoAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional billing contact. Accepted only when the extension schema includes this
// role. When the registry requires an omitted contact, Cloudflare may derive it
// from `contacts.registrant`.
type RegistrationNewParamsContactsBilling struct {
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
	PostalInfo param.Field[RegistrationNewParamsContactsBillingPostalInfo] `json:"postal_info" api:"required"`
	// Fax number in E.164 format (e.g., `+1.5555555555`). Optional. Most registrations
	// do not require a fax number.
	Fax param.Field[string] `json:"fax"`
}

func (r RegistrationNewParamsContactsBilling) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Postal/mailing information for the contact. The `name` field is the complete
// contact name in one string. Some registries require a complete personal name,
// including a family or last name where applicable, but this API does not accept
// separate first-name and last-name fields for registration contacts.
type RegistrationNewParamsContactsBillingPostalInfo struct {
	// Physical mailing address for the registrant contact.
	Address param.Field[RegistrationNewParamsContactsBillingPostalInfoAddress] `json:"address" api:"required"`
	// Full legal name of the contact, including all required name components for an
	// individual or authorized representative. Some registries require a complete
	// personal name that includes a family or last name where applicable. Provide the
	// complete name in this single field, for example `Ada Lovelace`; do not send
	// separate first-name or last-name fields.
	Name param.Field[string] `json:"name" api:"required"`
	// Organization or company name. Optional for individual registrants.
	Organization param.Field[string] `json:"organization"`
}

func (r RegistrationNewParamsContactsBillingPostalInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Physical mailing address for the registrant contact.
type RegistrationNewParamsContactsBillingPostalInfoAddress struct {
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

func (r RegistrationNewParamsContactsBillingPostalInfoAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional registrant contact. If omitted, the account's default address book
// entry is used instead.
type RegistrationNewParamsContactsRegistrant struct {
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
	PostalInfo param.Field[RegistrationNewParamsContactsRegistrantPostalInfo] `json:"postal_info" api:"required"`
	// Fax number in E.164 format (e.g., `+1.5555555555`). Optional. Most registrations
	// do not require a fax number.
	Fax param.Field[string] `json:"fax"`
}

func (r RegistrationNewParamsContactsRegistrant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Postal/mailing information for the contact. The `name` field is the complete
// contact name in one string. Some registries require a complete personal name,
// including a family or last name where applicable, but this API does not accept
// separate first-name and last-name fields for registration contacts.
type RegistrationNewParamsContactsRegistrantPostalInfo struct {
	// Physical mailing address for the registrant contact.
	Address param.Field[RegistrationNewParamsContactsRegistrantPostalInfoAddress] `json:"address" api:"required"`
	// Full legal name of the contact, including all required name components for an
	// individual or authorized representative. Some registries require a complete
	// personal name that includes a family or last name where applicable. Provide the
	// complete name in this single field, for example `Ada Lovelace`; do not send
	// separate first-name or last-name fields.
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

// Optional technical contact. Accepted only when the extension schema includes
// this role. When the registry requires an omitted contact, Cloudflare may derive
// it from `contacts.registrant`.
type RegistrationNewParamsContactsTechnical struct {
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
	PostalInfo param.Field[RegistrationNewParamsContactsTechnicalPostalInfo] `json:"postal_info" api:"required"`
	// Fax number in E.164 format (e.g., `+1.5555555555`). Optional. Most registrations
	// do not require a fax number.
	Fax param.Field[string] `json:"fax"`
}

func (r RegistrationNewParamsContactsTechnical) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Postal/mailing information for the contact. The `name` field is the complete
// contact name in one string. Some registries require a complete personal name,
// including a family or last name where applicable, but this API does not accept
// separate first-name and last-name fields for registration contacts.
type RegistrationNewParamsContactsTechnicalPostalInfo struct {
	// Physical mailing address for the registrant contact.
	Address param.Field[RegistrationNewParamsContactsTechnicalPostalInfoAddress] `json:"address" api:"required"`
	// Full legal name of the contact, including all required name components for an
	// individual or authorized representative. Some registries require a complete
	// personal name that includes a family or last name where applicable. Provide the
	// complete name in this single field, for example `Ada Lovelace`; do not send
	// separate first-name or last-name fields.
	Name param.Field[string] `json:"name" api:"required"`
	// Organization or company name. Optional for individual registrants.
	Organization param.Field[string] `json:"organization"`
}

func (r RegistrationNewParamsContactsTechnicalPostalInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Physical mailing address for the registrant contact.
type RegistrationNewParamsContactsTechnicalPostalInfoAddress struct {
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

func (r RegistrationNewParamsContactsTechnicalPostalInfoAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sets the WHOIS privacy mode for the registration. Defaults to `redaction`.
//
//   - `off`: Disables WHOIS privacy.
//   - `redaction`: Requests WHOIS redaction where the extension supports it. Some
//     extensions exclude privacy and redaction.
type RegistrationNewParamsPrivacyMode string

const (
	RegistrationNewParamsPrivacyModeOff       RegistrationNewParamsPrivacyMode = "off"
	RegistrationNewParamsPrivacyModeRedaction RegistrationNewParamsPrivacyMode = "redaction"
)

func (r RegistrationNewParamsPrivacyMode) IsKnown() bool {
	switch r {
	case RegistrationNewParamsPrivacyModeOff, RegistrationNewParamsPrivacyModeRedaction:
		return true
	}
	return false
}

type RegistrationNewResponseEnvelope struct {
	Errors   []RegistrationNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RegistrationNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Status of an async registration workflow.
	Result RegistrationNewResponse `json:"result" api:"required"`
	// Whether the API call was successful.
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

type RegistrationNewResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrationNewResponseEnvelopeErrorsSource `json:"source"`
	JSON   registrationNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// registrationNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RegistrationNewResponseEnvelopeErrors]
type registrationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrationNewResponseEnvelopeErrorsSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    registrationNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// registrationNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RegistrationNewResponseEnvelopeErrorsSource]
type registrationNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegistrationNewResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrationNewResponseEnvelopeMessagesSource `json:"source"`
	JSON   registrationNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// registrationNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RegistrationNewResponseEnvelopeMessages]
type registrationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrationNewResponseEnvelopeMessagesSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    registrationNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// registrationNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RegistrationNewResponseEnvelopeMessagesSource]
type registrationNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
	// Identifier.
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
	// Identifier.
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
	Errors   []RegistrationEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RegistrationEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Status of an async registration workflow.
	Result RegistrationEditResponse `json:"result" api:"required"`
	// Whether the API call was successful.
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

type RegistrationEditResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrationEditResponseEnvelopeErrorsSource `json:"source"`
	JSON   registrationEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// registrationEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RegistrationEditResponseEnvelopeErrors]
type registrationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrationEditResponseEnvelopeErrorsSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                           `json:"pointer" api:"required"`
	JSON    registrationEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// registrationEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RegistrationEditResponseEnvelopeErrorsSource]
type registrationEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegistrationEditResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrationEditResponseEnvelopeMessagesSource `json:"source"`
	JSON   registrationEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// registrationEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RegistrationEditResponseEnvelopeMessages]
type registrationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrationEditResponseEnvelopeMessagesSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                             `json:"pointer" api:"required"`
	JSON    registrationEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// registrationEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RegistrationEditResponseEnvelopeMessagesSource]
type registrationEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RegistrationGetResponseEnvelope struct {
	Errors   []RegistrationGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RegistrationGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// A domain registration resource representing the current state of a registered
	// domain.
	Result RegistrationGetResponse `json:"result" api:"required"`
	// Whether the API call was successful.
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

type RegistrationGetResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrationGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   registrationGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// registrationGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RegistrationGetResponseEnvelopeErrors]
type registrationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrationGetResponseEnvelopeErrorsSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    registrationGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// registrationGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RegistrationGetResponseEnvelopeErrorsSource]
type registrationGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegistrationGetResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrationGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   registrationGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// registrationGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RegistrationGetResponseEnvelopeMessages]
type registrationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrationGetResponseEnvelopeMessagesSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    registrationGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// registrationGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RegistrationGetResponseEnvelopeMessagesSource]
type registrationGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
