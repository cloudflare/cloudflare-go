// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
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

// Acknowledge + Retry messages from a Queue
func (r *MessageService) Ack(ctx context.Context, queueID string, params MessageAckParams, opts ...option.RequestOption) (res *MessageAckResponse, err error) {
	var env MessageAckResponseEnvelope
	opts = append(r.Options[:], opts...)
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

// Pull a batch of messages from a Queue
func (r *MessageService) Pull(ctx context.Context, queueID string, params MessagePullParams, opts ...option.RequestOption) (res *[]MessagePullResponse, err error) {
	var env MessagePullResponseEnvelope
	opts = append(r.Options[:], opts...)
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
	// The number of messages that were succesfully acknowledged.
	AckCount float64 `json:"ackCount"`
	// The number of messages that were succesfully retried.
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
	ID       string  `json:"id"`
	Attempts float64 `json:"attempts"`
	Body     string  `json:"body"`
	// An ID that represents an "in-flight" message that has been pulled from a Queue.
	// You must hold on to this ID and use it to acknowledge this message.
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
	// A Resource identifier.
	AccountID param.Field[string]                  `path:"account_id,required"`
	Acks      param.Field[[]MessageAckParamsAck]   `json:"acks"`
	Retries   param.Field[[]MessageAckParamsRetry] `json:"retries"`
}

func (r MessageAckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageAckParamsAck struct {
	// An ID that represents an "in-flight" message that has been pulled from a Queue.
	// You must hold on to this ID and use it to acknowledge this message.
	LeaseID param.Field[string] `json:"lease_id"`
}

func (r MessageAckParamsAck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageAckParamsRetry struct {
	// The number of seconds to delay before making the message available for another
	// attempt.
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
	// An ID that represents an "in-flight" message that has been pulled from a Queue.
	// You must hold on to this ID and use it to acknowledge this message.
	LeaseID param.Field[string] `json:"lease_id"`
}

func (r MessageAckParamsRetry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageAckResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   MessageAckResponse    `json:"result"`
	// Indicates if the API call was successful or not.
	Success MessageAckResponseEnvelopeSuccess `json:"success"`
	JSON    messageAckResponseEnvelopeJSON    `json:"-"`
}

// messageAckResponseEnvelopeJSON contains the JSON metadata for the struct
// [MessageAckResponseEnvelope]
type messageAckResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageAckResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageAckResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
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

type MessagePullParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs param.Field[float64] `json:"visibility_timeout_ms"`
}

func (r MessagePullParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessagePullResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   []MessagePullResponse `json:"result"`
	// Indicates if the API call was successful or not.
	Success MessagePullResponseEnvelopeSuccess `json:"success"`
	JSON    messagePullResponseEnvelopeJSON    `json:"-"`
}

// messagePullResponseEnvelopeJSON contains the JSON metadata for the struct
// [MessagePullResponseEnvelope]
type messagePullResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagePullResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePullResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
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
