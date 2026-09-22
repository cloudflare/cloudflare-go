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

// UpdateStatusService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUpdateStatusService] method instead.
type UpdateStatusService struct {
	Options []option.RequestOption
}

// NewUpdateStatusService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUpdateStatusService(opts ...option.RequestOption) (r *UpdateStatusService) {
	r = &UpdateStatusService{}
	r.Options = opts
	return
}

// Returns the current status of a domain update workflow.
//
// Use this endpoint to poll for completion when the PATCH response returned
// `202 Accepted`. The URL is provided in the `links.self` field of the workflow
// status response.
//
// Poll this endpoint until the workflow reaches a terminal state or a state that
// requires user attention.
//
// Use increasing backoff between polls. When the workflow remains blocked on a
// third party, use a longer polling interval and do not poll indefinitely.
func (r *UpdateStatusService) Get(ctx context.Context, domainName string, query UpdateStatusGetParams, opts ...option.RequestOption) (res *UpdateStatusGetResponse, err error) {
	var env UpdateStatusGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar-sandbox/registrations/%s/update-status", query.AccountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Status of an async registration workflow.
type UpdateStatusGetResponse struct {
	// Indicates whether the workflow reached a terminal state. A `succeeded` or
	// `failed` state returns `true`; `pending`, `in_progress`, `action_required`, and
	// `blocked` return `false`.
	Completed bool                         `json:"completed" api:"required"`
	CreatedAt time.Time                    `json:"created_at" api:"required" format:"date-time"`
	Links     UpdateStatusGetResponseLinks `json:"links" api:"required"`
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
	State     UpdateStatusGetResponseState `json:"state" api:"required"`
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
	Error UpdateStatusGetResponseError `json:"error" api:"nullable"`
	JSON  updateStatusGetResponseJSON  `json:"-"`
}

// updateStatusGetResponseJSON contains the JSON metadata for the struct
// [UpdateStatusGetResponse]
type updateStatusGetResponseJSON struct {
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

func (r *UpdateStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateStatusGetResponseJSON) RawJSON() string {
	return r.raw
}

type UpdateStatusGetResponseLinks struct {
	// URL to this status resource.
	Self string `json:"self" api:"required"`
	// URL to the domain resource.
	Resource string                           `json:"resource"`
	JSON     updateStatusGetResponseLinksJSON `json:"-"`
}

// updateStatusGetResponseLinksJSON contains the JSON metadata for the struct
// [UpdateStatusGetResponseLinks]
type updateStatusGetResponseLinksJSON struct {
	Self        apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateStatusGetResponseLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateStatusGetResponseLinksJSON) RawJSON() string {
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
type UpdateStatusGetResponseState string

const (
	UpdateStatusGetResponseStatePending        UpdateStatusGetResponseState = "pending"
	UpdateStatusGetResponseStateInProgress     UpdateStatusGetResponseState = "in_progress"
	UpdateStatusGetResponseStateActionRequired UpdateStatusGetResponseState = "action_required"
	UpdateStatusGetResponseStateBlocked        UpdateStatusGetResponseState = "blocked"
	UpdateStatusGetResponseStateSucceeded      UpdateStatusGetResponseState = "succeeded"
	UpdateStatusGetResponseStateFailed         UpdateStatusGetResponseState = "failed"
)

func (r UpdateStatusGetResponseState) IsKnown() bool {
	switch r {
	case UpdateStatusGetResponseStatePending, UpdateStatusGetResponseStateInProgress, UpdateStatusGetResponseStateActionRequired, UpdateStatusGetResponseStateBlocked, UpdateStatusGetResponseStateSucceeded, UpdateStatusGetResponseStateFailed:
		return true
	}
	return false
}

// Provides error details when a workflow reaches the `failed` state. The workflow
// type (registration, update, etc.) and underlying registry response determine the
// specific codes and messages. Workflow error codes differ from immediate HTTP
// error `errors[].code` values in non-2xx responses. Surface `error.message` to
// the user for context.
type UpdateStatusGetResponseError struct {
	// Machine-readable error code identifying the failure reason.
	Code string `json:"code" api:"required"`
	// Human-readable explanation of the failure. May include registry-specific
	// details.
	Message string                           `json:"message" api:"required"`
	JSON    updateStatusGetResponseErrorJSON `json:"-"`
}

// updateStatusGetResponseErrorJSON contains the JSON metadata for the struct
// [UpdateStatusGetResponseError]
type updateStatusGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateStatusGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateStatusGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type UpdateStatusGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type UpdateStatusGetResponseEnvelope struct {
	Errors   []UpdateStatusGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []UpdateStatusGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Status of an async registration workflow.
	Result UpdateStatusGetResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success UpdateStatusGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    updateStatusGetResponseEnvelopeJSON    `json:"-"`
}

// updateStatusGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UpdateStatusGetResponseEnvelope]
type updateStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateStatusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UpdateStatusGetResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source UpdateStatusGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   updateStatusGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// updateStatusGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UpdateStatusGetResponseEnvelopeErrors]
type updateStatusGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateStatusGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateStatusGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type UpdateStatusGetResponseEnvelopeErrorsSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    updateStatusGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// updateStatusGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [UpdateStatusGetResponseEnvelopeErrorsSource]
type updateStatusGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateStatusGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateStatusGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UpdateStatusGetResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source UpdateStatusGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   updateStatusGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// updateStatusGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UpdateStatusGetResponseEnvelopeMessages]
type updateStatusGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateStatusGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateStatusGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type UpdateStatusGetResponseEnvelopeMessagesSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    updateStatusGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// updateStatusGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [UpdateStatusGetResponseEnvelopeMessagesSource]
type updateStatusGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateStatusGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateStatusGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UpdateStatusGetResponseEnvelopeSuccess bool

const (
	UpdateStatusGetResponseEnvelopeSuccessTrue UpdateStatusGetResponseEnvelopeSuccess = true
)

func (r UpdateStatusGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UpdateStatusGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
