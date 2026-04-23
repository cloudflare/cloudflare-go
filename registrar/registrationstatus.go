// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
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
func (r *RegistrationStatusService) Get(ctx context.Context, domainName string, query RegistrationStatusGetParams, opts ...option.RequestOption) (res *WorkflowStatus, err error) {
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
	path := fmt.Sprintf("accounts/%s/registrar/registrations/%s/registration-status", query.AccountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type RegistrationStatusGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RegistrationStatusGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Status of an async registration workflow.
	Result WorkflowStatus `json:"result" api:"required"`
	// Whether the API call was successful
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

// Whether the API call was successful
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
