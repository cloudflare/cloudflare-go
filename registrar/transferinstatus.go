// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// TransferInStatusService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransferInStatusService] method instead.
type TransferInStatusService struct {
	Options []option.RequestOption
}

// NewTransferInStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransferInStatusService(opts ...option.RequestOption) (r *TransferInStatusService) {
	r = &TransferInStatusService{}
	r.Options = opts
	return
}

// Returns the current status of a domain transfer workflow.
//
// Use this endpoint to poll transfer progress after initiating a transfer with
// `POST /accounts/{account_id}/registrar/registrations/{domain_name}/transfer-in`.
// The URL is provided in the `links.self` field of the transfer response.
//
// ### Transfer timelines
//
// Transfers typically take 1–10 days due to ICANN-mandated approval windows.
//
// ### Workflow states
//
// **Terminal states:** `succeeded` and `failed` are terminal and always have
// `completed: true`.
//
// **Non-terminal states:**
//
//   - `in_progress`: Transfer has been submitted to the registry and is being
//     processed. Continue polling.
//   - `blocked`: The workflow is waiting on the losing registrar or registry to
//     release the domain. This is the **most common state** for transfers and is
//     entirely normal — it means the ICANN transfer approval window is in effect.
//     The losing registrar has up to 5 days to approve or reject. Continue polling
//     with longer intervals (e.g., every 30–60 minutes).
//   - `action_required`: The user needs to take action (e.g., the FOA email needs to
//     be accepted). See `context` for details on what is needed.
//   - `pending`: Transfer workflow created but not yet started processing.
//
// ### Polling guidance
//
// Adjust your polling interval based on the current workflow state:
//
//   - `pending` or `in_progress`: Poll every 30 seconds.
//   - `blocked`: The transfer is waiting on a third party (e.g., losing registrar
//     approval). Poll every 30–60 minutes.
//   - `action_required`: Stop polling. The workflow will not advance until the user
//     takes action. Check `context` for details on what is needed.
//   - `succeeded` or `failed`: Terminal — stop polling.
func (r *TransferInStatusService) Get(ctx context.Context, domainName string, query TransferInStatusGetParams, opts ...option.RequestOption) (res *WorkflowStatus, err error) {
	var env TransferInStatusGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar/registrations/%s/transfer-in-status", query.AccountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type TransferInStatusGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type TransferInStatusGetResponseEnvelope struct {
	Errors   []TransferInStatusGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []TransferInStatusGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Status of an async registration workflow.
	Result WorkflowStatus `json:"result" api:"required"`
	// Whether the API call was successful.
	Success TransferInStatusGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    transferInStatusGetResponseEnvelopeJSON    `json:"-"`
}

// transferInStatusGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [TransferInStatusGetResponseEnvelope]
type transferInStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferInStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferInStatusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TransferInStatusGetResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source TransferInStatusGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   transferInStatusGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// transferInStatusGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TransferInStatusGetResponseEnvelopeErrors]
type transferInStatusGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferInStatusGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferInStatusGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type TransferInStatusGetResponseEnvelopeErrorsSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                              `json:"pointer" api:"required"`
	JSON    transferInStatusGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// transferInStatusGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [TransferInStatusGetResponseEnvelopeErrorsSource]
type transferInStatusGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferInStatusGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferInStatusGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TransferInStatusGetResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source TransferInStatusGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   transferInStatusGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// transferInStatusGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [TransferInStatusGetResponseEnvelopeMessages]
type transferInStatusGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferInStatusGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferInStatusGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type TransferInStatusGetResponseEnvelopeMessagesSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                                `json:"pointer" api:"required"`
	JSON    transferInStatusGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// transferInStatusGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [TransferInStatusGetResponseEnvelopeMessagesSource]
type transferInStatusGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferInStatusGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferInStatusGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TransferInStatusGetResponseEnvelopeSuccess bool

const (
	TransferInStatusGetResponseEnvelopeSuccessTrue TransferInStatusGetResponseEnvelopeSuccess = true
)

func (r TransferInStatusGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TransferInStatusGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
