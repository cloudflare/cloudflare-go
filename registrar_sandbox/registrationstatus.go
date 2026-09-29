// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_sandbox

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// RegistrationStatusService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistrationStatusService] method instead.
type RegistrationStatusService struct {
	Options []option.RequestOption
}

// NewRegistrationStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegistrationStatusService(opts ...option.RequestOption) (r *RegistrationStatusService) {
	r = &RegistrationStatusService{}
	r.Options = opts
	return
}

// Returns the current status of a domain registration workflow.
//
// Use this endpoint to poll for completion when the POST response returned
// `202 Accepted`. The URL is provided in the `links.self` field of the workflow
// status response.
//
// Poll this endpoint until the workflow reaches a terminal state or a state that
// requires user attention.
//
// **Terminal states:** `succeeded` and `failed` are terminal and always have
// `completed: true`.
//
// **Non-terminal states:**
//
//   - `action_required` has `completed: false` and will not resolve on its own. The
//     workflow is paused pending user intervention.
//   - `blocked` has `completed: false` and indicates the workflow is waiting on a
//     third party such as the extension registry or losing registrar. Continue
//     polling while informing the user of the delay.
//
// Use increasing backoff between polls. When `state: blocked`, use a longer
// polling interval and do not poll indefinitely.
//
// A naive polling loop that only checks `completed` can run indefinitely when
// `state: action_required`. Break explicitly on `action_required`:
//
// ```js
// let status;
//
//	do {
//	  await new Promise((r) => setTimeout(r, 2000));
//	  status = await cloudflare.request({
//	    method: "GET",
//	    path: reg.result.links.self,
//	  });
//	} while (!status.result.completed && status.result.state !== "action_required");
//
//	if (status.result.state === "action_required") {
//	  // Surface context.action and context.confirmation_sent_to to the user.
//	  // Do not re-submit the registration request.
//	}
//
// ```
func (r *RegistrationStatusService) Get(ctx context.Context, domainName string, query RegistrationStatusGetParams, opts ...option.RequestOption) (res *RegistrationStatusGetResponse, err error) {
	var env RegistrationStatusGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar-sandbox/registrations/%s/registration-status", query.AccountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Status of an async registration workflow.
type RegistrationStatusGetResponse struct {
	// Indicates whether the workflow reached a terminal state. A `succeeded` or
	// `failed` state returns `true`; `pending`, `in_progress`, `action_required`, and
	// `blocked` return `false`.
	Completed bool                               `json:"completed" api:"required"`
	CreatedAt time.Time                          `json:"created_at" api:"required" format:"date-time"`
	Links     RegistrationStatusGetResponseLinks `json:"links" api:"required"`
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
	State     RegistrationStatusGetResponseState `json:"state" api:"required"`
	UpdatedAt time.Time                          `json:"updated_at" api:"required" format:"date-time"`
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
	Error RegistrationStatusGetResponseError `json:"error" api:"nullable"`
	JSON  registrationStatusGetResponseJSON  `json:"-"`
}

// registrationStatusGetResponseJSON contains the JSON metadata for the struct
// [RegistrationStatusGetResponse]
type registrationStatusGetResponseJSON struct {
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

func (r *RegistrationStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationStatusGetResponseJSON) RawJSON() string {
	return r.raw
}

type RegistrationStatusGetResponseLinks struct {
	// URL to this status resource.
	Self string `json:"self" api:"required"`
	// URL to the domain resource.
	Resource string                                 `json:"resource"`
	JSON     registrationStatusGetResponseLinksJSON `json:"-"`
}

// registrationStatusGetResponseLinksJSON contains the JSON metadata for the struct
// [RegistrationStatusGetResponseLinks]
type registrationStatusGetResponseLinksJSON struct {
	Self        apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationStatusGetResponseLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationStatusGetResponseLinksJSON) RawJSON() string {
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
type RegistrationStatusGetResponseState string

const (
	RegistrationStatusGetResponseStatePending        RegistrationStatusGetResponseState = "pending"
	RegistrationStatusGetResponseStateInProgress     RegistrationStatusGetResponseState = "in_progress"
	RegistrationStatusGetResponseStateActionRequired RegistrationStatusGetResponseState = "action_required"
	RegistrationStatusGetResponseStateBlocked        RegistrationStatusGetResponseState = "blocked"
	RegistrationStatusGetResponseStateSucceeded      RegistrationStatusGetResponseState = "succeeded"
	RegistrationStatusGetResponseStateFailed         RegistrationStatusGetResponseState = "failed"
)

func (r RegistrationStatusGetResponseState) IsKnown() bool {
	switch r {
	case RegistrationStatusGetResponseStatePending, RegistrationStatusGetResponseStateInProgress, RegistrationStatusGetResponseStateActionRequired, RegistrationStatusGetResponseStateBlocked, RegistrationStatusGetResponseStateSucceeded, RegistrationStatusGetResponseStateFailed:
		return true
	}
	return false
}

// Provides error details when a workflow reaches the `failed` state. The workflow
// type (registration, update, etc.) and underlying registry response determine the
// specific codes and messages. Workflow error codes differ from immediate HTTP
// error `errors[].code` values in non-2xx responses. Surface `error.message` to
// the user for context.
type RegistrationStatusGetResponseError struct {
	// Machine-readable error code identifying the failure reason.
	Code string `json:"code" api:"required"`
	// Human-readable explanation of the failure. May include registry-specific
	// details.
	Message string                                 `json:"message" api:"required"`
	JSON    registrationStatusGetResponseErrorJSON `json:"-"`
}

// registrationStatusGetResponseErrorJSON contains the JSON metadata for the struct
// [RegistrationStatusGetResponseError]
type registrationStatusGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationStatusGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationStatusGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type RegistrationStatusGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RegistrationStatusGetResponseEnvelope struct {
	Errors   []RegistrationStatusGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RegistrationStatusGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Status of an async registration workflow.
	Result RegistrationStatusGetResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success RegistrationStatusGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    registrationStatusGetResponseEnvelopeJSON    `json:"-"`
}

// registrationStatusGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegistrationStatusGetResponseEnvelope]
type registrationStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationStatusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegistrationStatusGetResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrationStatusGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   registrationStatusGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// registrationStatusGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RegistrationStatusGetResponseEnvelopeErrors]
type registrationStatusGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationStatusGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationStatusGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrationStatusGetResponseEnvelopeErrorsSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                                `json:"pointer" api:"required"`
	JSON    registrationStatusGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// registrationStatusGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RegistrationStatusGetResponseEnvelopeErrorsSource]
type registrationStatusGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationStatusGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationStatusGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegistrationStatusGetResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source RegistrationStatusGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   registrationStatusGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// registrationStatusGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RegistrationStatusGetResponseEnvelopeMessages]
type registrationStatusGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationStatusGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationStatusGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type RegistrationStatusGetResponseEnvelopeMessagesSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                                  `json:"pointer" api:"required"`
	JSON    registrationStatusGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// registrationStatusGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RegistrationStatusGetResponseEnvelopeMessagesSource]
type registrationStatusGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrationStatusGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrationStatusGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RegistrationStatusGetResponseEnvelopeSuccess bool

const (
	RegistrationStatusGetResponseEnvelopeSuccessTrue RegistrationStatusGetResponseEnvelopeSuccess = true
)

func (r RegistrationStatusGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegistrationStatusGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
