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
	"github.com/cloudflare/cloudflare-go/v7/shared"
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
	// Whether the workflow has reached a terminal state. `true` when `state` is
	// `succeeded` or `failed`. `false` for `pending`, `in_progress`,
	// `action_required`, and `blocked`.
	Completed bool                         `json:"completed" api:"required"`
	CreatedAt time.Time                    `json:"created_at" api:"required" format:"date-time"`
	Links     UpdateStatusGetResponseLinks `json:"links" api:"required"`
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
	State     UpdateStatusGetResponseState `json:"state" api:"required"`
	UpdatedAt time.Time                    `json:"updated_at" api:"required" format:"date-time"`
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

// Error details when a workflow reaches the `failed` state. The specific error
// codes and messages depend on the workflow type (registration, update, etc.) and
// the underlying registry response. These workflow error codes are separate from
// immediate HTTP error `errors[].code` values returned by non-2xx responses.
// Surface `error.message` to the user for context.
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
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type UpdateStatusGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Status of an async registration workflow.
	Result UpdateStatusGetResponse `json:"result" api:"required"`
	// Whether the API call was successful
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

// Whether the API call was successful
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
