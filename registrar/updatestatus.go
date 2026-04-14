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
func (r *UpdateStatusService) Get(ctx context.Context, domainName string, query UpdateStatusGetParams, opts ...option.RequestOption) (res *WorkflowStatus, err error) {
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
	path := fmt.Sprintf("accounts/%s/registrar/registrations/%s/update-status", query.AccountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type UpdateStatusGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type UpdateStatusGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Status of an async registration workflow.
	Result WorkflowStatus `json:"result" api:"required"`
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
