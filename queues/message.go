// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// MessageService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r *MessageService) {
	r = &MessageService{}
	r.Options = opts
	return
}

// Acknowledge + Retry messages from a Queue.
func (r *MessageService) Ack(ctx context.Context, queueID string, params MessageAckParams, opts ...option.RequestOption) (res *MessageAckResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MessageAckResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/messages/ack", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Pull a batch of messages from a Queue.
func (r *MessageService) Pull(ctx context.Context, queueID string, params MessagePullParams, opts ...option.RequestOption) (res *[]MessagePullResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MessagePullResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/messages/pull", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MessageAckResponse struct {
	// The number of messages that were succesfully acknowledged
	AckCount float64 `json:"ackCount"`
	// The number of messages that were succesfully retried
	RetryCount float64                `json:"retryCount"`
	Warnings   []string               `json:"warnings"`
	JSON       messageAckResponseJSON `json:"-"`
}

// messageAckResponseJSON contains the JSON metadata for the struct
// [MessageAckResponse]
type messageAckResponseJSON struct {
	AckCount    apijson.Field
	RetryCount  apijson.Field
	Warnings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageAckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageAckResponseJSON) RawJSON() string {
	return r.raw
}

type MessagePullResponse struct {
	ID          string                  `json:"id"`
	Attempts    float64                 `json:"attempts"`
	Body        string                  `json:"body"`
	LeaseID     string                  `json:"lease_id"`
	Metadata    interface{}             `json:"metadata"`
	TimestampMs float64                 `json:"timestamp_ms"`
	JSON        messagePullResponseJSON `json:"-"`
}

// messagePullResponseJSON contains the JSON metadata for the struct
// [MessagePullResponse]
type messagePullResponseJSON struct {
	ID          apijson.Field
	Attempts    apijson.Field
	Body        apijson.Field
	LeaseID     apijson.Field
	Metadata    apijson.Field
	TimestampMs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagePullResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePullResponseJSON) RawJSON() string {
	return r.raw
}

type MessageAckParams struct {
	// Identifier
	AccountID param.Field[string]                  `path:"account_id,required"`
	Acks      param.Field[[]MessageAckParamsAck]   `json:"acks"`
	Retries   param.Field[[]MessageAckParamsRetry] `json:"retries"`
}

func (r MessageAckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageAckParamsAck struct {
	// Lease ID for a message to acknowledge.
	LeaseID param.Field[string] `json:"lease_id"`
}

func (r MessageAckParamsAck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageAckParamsRetry struct {
	// The number of seconds to delay before making the message available for another
	// attempt.
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
	// Lease ID for a message to retry.
	LeaseID param.Field[string] `json:"lease_id"`
}

func (r MessageAckParamsRetry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageAckResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   MessageAckResponse    `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    MessageAckResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo MessageAckResponseEnvelopeResultInfo `json:"result_info"`
	JSON       messageAckResponseEnvelopeJSON       `json:"-"`
}

// messageAckResponseEnvelopeJSON contains the JSON metadata for the struct
// [MessageAckResponseEnvelope]
type messageAckResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageAckResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageAckResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MessageAckResponseEnvelopeSuccess bool

const (
	MessageAckResponseEnvelopeSuccessTrue MessageAckResponseEnvelopeSuccess = true
)

func (r MessageAckResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MessageAckResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MessageAckResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       messageAckResponseEnvelopeResultInfoJSON `json:"-"`
}

// messageAckResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [MessageAckResponseEnvelopeResultInfo]
type messageAckResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageAckResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageAckResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type MessagePullParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The maximum number of messages to include in a batch
	BatchSize param.Field[float64] `json:"batch_size"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs param.Field[float64] `json:"visibility_timeout_ms"`
}

func (r MessagePullParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessagePullResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []MessagePullResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    MessagePullResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo MessagePullResponseEnvelopeResultInfo `json:"result_info"`
	JSON       messagePullResponseEnvelopeJSON       `json:"-"`
}

// messagePullResponseEnvelopeJSON contains the JSON metadata for the struct
// [MessagePullResponseEnvelope]
type messagePullResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagePullResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePullResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MessagePullResponseEnvelopeSuccess bool

const (
	MessagePullResponseEnvelopeSuccessTrue MessagePullResponseEnvelopeSuccess = true
)

func (r MessagePullResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MessagePullResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MessagePullResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       messagePullResponseEnvelopeResultInfoJSON `json:"-"`
}

// messagePullResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [MessagePullResponseEnvelopeResultInfo]
type messagePullResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagePullResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePullResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
