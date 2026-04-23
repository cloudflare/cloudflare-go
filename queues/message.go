// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues

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
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/messages/ack", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Push a batch of message to a Queue
func (r *MessageService) BulkPush(ctx context.Context, queueID string, params MessageBulkPushParams, opts ...option.RequestOption) (res *MessageBulkPushResponse, err error) {
	var env MessageBulkPushResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/messages/batch", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Pull a batch of messages from a Queue
func (r *MessageService) Pull(ctx context.Context, queueID string, params MessagePullParams, opts ...option.RequestOption) (res *MessagePullResponse, err error) {
	var env MessagePullResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/messages/pull", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Push a message to a Queue
func (r *MessageService) Push(ctx context.Context, queueID string, params MessagePushParams, opts ...option.RequestOption) (res *MessagePushResponse, err error) {
	var env MessagePushResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/messages", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type MessageAckResponse struct {
	// The number of messages that were succesfully acknowledged.
	AckCount float64 `json:"ackCount"`
	// The number of messages that were succesfully retried.
	RetryCount float64 `json:"retryCount"`
	// Map of lease IDs to warning messages encountered during acknowledgement.
	Warnings map[string]string      `json:"warnings"`
	JSON     messageAckResponseJSON `json:"-"`
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

type MessageBulkPushResponse struct {
	Metadata MessageBulkPushResponseMetadata `json:"metadata"`
	JSON     messageBulkPushResponseJSON     `json:"-"`
}

// messageBulkPushResponseJSON contains the JSON metadata for the struct
// [MessageBulkPushResponse]
type messageBulkPushResponseJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBulkPushResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBulkPushResponseJSON) RawJSON() string {
	return r.raw
}

type MessageBulkPushResponseMetadata struct {
	// Best-effort metrics for the queue. Values may be approximate due to the
	// distributed nature of queues.
	Metrics MessageBulkPushResponseMetadataMetrics `json:"metrics"`
	JSON    messageBulkPushResponseMetadataJSON    `json:"-"`
}

// messageBulkPushResponseMetadataJSON contains the JSON metadata for the struct
// [MessageBulkPushResponseMetadata]
type messageBulkPushResponseMetadataJSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBulkPushResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBulkPushResponseMetadataJSON) RawJSON() string {
	return r.raw
}

// Best-effort metrics for the queue. Values may be approximate due to the
// distributed nature of queues.
type MessageBulkPushResponseMetadataMetrics struct {
	// The size in bytes of unacknowledged messages in the queue.
	BacklogBytes float64 `json:"backlog_bytes" api:"required"`
	// The number of unacknowledged messages in the queue.
	BacklogCount float64 `json:"backlog_count" api:"required"`
	// Unix timestamp in milliseconds of the oldest unacknowledged message in the
	// queue. Returns 0 if unknown.
	OldestMessageTimestampMs float64                                    `json:"oldest_message_timestamp_ms" api:"required"`
	JSON                     messageBulkPushResponseMetadataMetricsJSON `json:"-"`
}

// messageBulkPushResponseMetadataMetricsJSON contains the JSON metadata for the
// struct [MessageBulkPushResponseMetadataMetrics]
type messageBulkPushResponseMetadataMetricsJSON struct {
	BacklogBytes             apijson.Field
	BacklogCount             apijson.Field
	OldestMessageTimestampMs apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *MessageBulkPushResponseMetadataMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBulkPushResponseMetadataMetricsJSON) RawJSON() string {
	return r.raw
}

type MessagePullResponse struct {
	// The number of unacknowledged messages in the queue.
	MessageBacklogCount float64                      `json:"message_backlog_count"`
	Messages            []MessagePullResponseMessage `json:"messages"`
	Metadata            MessagePullResponseMetadata  `json:"metadata"`
	JSON                messagePullResponseJSON      `json:"-"`
}

// messagePullResponseJSON contains the JSON metadata for the struct
// [MessagePullResponse]
type messagePullResponseJSON struct {
	MessageBacklogCount apijson.Field
	Messages            apijson.Field
	Metadata            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MessagePullResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePullResponseJSON) RawJSON() string {
	return r.raw
}

type MessagePullResponseMessage struct {
	ID       string  `json:"id"`
	Attempts float64 `json:"attempts"`
	Body     string  `json:"body"`
	// An ID that represents an "in-flight" message that has been pulled from a Queue.
	// You must hold on to this ID and use it to acknowledge this message.
	LeaseID     string                         `json:"lease_id"`
	Metadata    interface{}                    `json:"metadata"`
	TimestampMs float64                        `json:"timestamp_ms"`
	JSON        messagePullResponseMessageJSON `json:"-"`
}

// messagePullResponseMessageJSON contains the JSON metadata for the struct
// [MessagePullResponseMessage]
type messagePullResponseMessageJSON struct {
	ID          apijson.Field
	Attempts    apijson.Field
	Body        apijson.Field
	LeaseID     apijson.Field
	Metadata    apijson.Field
	TimestampMs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagePullResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePullResponseMessageJSON) RawJSON() string {
	return r.raw
}

type MessagePullResponseMetadata struct {
	// Best-effort metrics for the queue. Values may be approximate due to the
	// distributed nature of queues.
	Metrics MessagePullResponseMetadataMetrics `json:"metrics"`
	JSON    messagePullResponseMetadataJSON    `json:"-"`
}

// messagePullResponseMetadataJSON contains the JSON metadata for the struct
// [MessagePullResponseMetadata]
type messagePullResponseMetadataJSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagePullResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePullResponseMetadataJSON) RawJSON() string {
	return r.raw
}

// Best-effort metrics for the queue. Values may be approximate due to the
// distributed nature of queues.
type MessagePullResponseMetadataMetrics struct {
	// The size in bytes of unacknowledged messages in the queue.
	BacklogBytes float64 `json:"backlog_bytes" api:"required"`
	// The number of unacknowledged messages in the queue.
	BacklogCount float64 `json:"backlog_count" api:"required"`
	// Unix timestamp in milliseconds of the oldest unacknowledged message in the
	// queue. Returns 0 if unknown.
	OldestMessageTimestampMs float64                                `json:"oldest_message_timestamp_ms" api:"required"`
	JSON                     messagePullResponseMetadataMetricsJSON `json:"-"`
}

// messagePullResponseMetadataMetricsJSON contains the JSON metadata for the struct
// [MessagePullResponseMetadataMetrics]
type messagePullResponseMetadataMetricsJSON struct {
	BacklogBytes             apijson.Field
	BacklogCount             apijson.Field
	OldestMessageTimestampMs apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *MessagePullResponseMetadataMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePullResponseMetadataMetricsJSON) RawJSON() string {
	return r.raw
}

type MessagePushResponse struct {
	Metadata MessagePushResponseMetadata `json:"metadata"`
	JSON     messagePushResponseJSON     `json:"-"`
}

// messagePushResponseJSON contains the JSON metadata for the struct
// [MessagePushResponse]
type messagePushResponseJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagePushResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePushResponseJSON) RawJSON() string {
	return r.raw
}

type MessagePushResponseMetadata struct {
	// Best-effort metrics for the queue. Values may be approximate due to the
	// distributed nature of queues.
	Metrics MessagePushResponseMetadataMetrics `json:"metrics"`
	JSON    messagePushResponseMetadataJSON    `json:"-"`
}

// messagePushResponseMetadataJSON contains the JSON metadata for the struct
// [MessagePushResponseMetadata]
type messagePushResponseMetadataJSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagePushResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePushResponseMetadataJSON) RawJSON() string {
	return r.raw
}

// Best-effort metrics for the queue. Values may be approximate due to the
// distributed nature of queues.
type MessagePushResponseMetadataMetrics struct {
	// The size in bytes of unacknowledged messages in the queue.
	BacklogBytes float64 `json:"backlog_bytes" api:"required"`
	// The number of unacknowledged messages in the queue.
	BacklogCount float64 `json:"backlog_count" api:"required"`
	// Unix timestamp in milliseconds of the oldest unacknowledged message in the
	// queue. Returns 0 if unknown.
	OldestMessageTimestampMs float64                                `json:"oldest_message_timestamp_ms" api:"required"`
	JSON                     messagePushResponseMetadataMetricsJSON `json:"-"`
}

// messagePushResponseMetadataMetricsJSON contains the JSON metadata for the struct
// [MessagePushResponseMetadataMetrics]
type messagePushResponseMetadataMetricsJSON struct {
	BacklogBytes             apijson.Field
	BacklogCount             apijson.Field
	OldestMessageTimestampMs apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *MessagePushResponseMetadataMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePushResponseMetadataMetricsJSON) RawJSON() string {
	return r.raw
}

type MessageAckParams struct {
	// A Resource identifier.
	AccountID param.Field[string]                  `path:"account_id" api:"required"`
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

type MessageBulkPushParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The number of seconds to wait for attempting to deliver this batch to consumers
	DelaySeconds param.Field[float64]                             `json:"delay_seconds"`
	Messages     param.Field[[]MessageBulkPushParamsMessageUnion] `json:"messages"`
}

func (r MessageBulkPushParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBulkPushParamsMessage struct {
	Body        param.Field[interface{}]                              `json:"body"`
	ContentType param.Field[MessageBulkPushParamsMessagesContentType] `json:"content_type"`
	// The number of seconds to wait for attempting to deliver this message to
	// consumers
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
}

func (r MessageBulkPushParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBulkPushParamsMessage) implementsMessageBulkPushParamsMessageUnion() {}

// Satisfied by [queues.MessageBulkPushParamsMessagesMqQueueMessageText],
// [queues.MessageBulkPushParamsMessagesMqQueueMessageJson],
// [MessageBulkPushParamsMessage].
type MessageBulkPushParamsMessageUnion interface {
	implementsMessageBulkPushParamsMessageUnion()
}

type MessageBulkPushParamsMessagesMqQueueMessageText struct {
	Body        param.Field[string]                                                     `json:"body"`
	ContentType param.Field[MessageBulkPushParamsMessagesMqQueueMessageTextContentType] `json:"content_type"`
	// The number of seconds to wait for attempting to deliver this message to
	// consumers
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
}

func (r MessageBulkPushParamsMessagesMqQueueMessageText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBulkPushParamsMessagesMqQueueMessageText) implementsMessageBulkPushParamsMessageUnion() {
}

type MessageBulkPushParamsMessagesMqQueueMessageTextContentType string

const (
	MessageBulkPushParamsMessagesMqQueueMessageTextContentTypeText MessageBulkPushParamsMessagesMqQueueMessageTextContentType = "text"
)

func (r MessageBulkPushParamsMessagesMqQueueMessageTextContentType) IsKnown() bool {
	switch r {
	case MessageBulkPushParamsMessagesMqQueueMessageTextContentTypeText:
		return true
	}
	return false
}

type MessageBulkPushParamsMessagesMqQueueMessageJson struct {
	Body        param.Field[interface{}]                                                `json:"body"`
	ContentType param.Field[MessageBulkPushParamsMessagesMqQueueMessageJsonContentType] `json:"content_type"`
	// The number of seconds to wait for attempting to deliver this message to
	// consumers
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
}

func (r MessageBulkPushParamsMessagesMqQueueMessageJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBulkPushParamsMessagesMqQueueMessageJson) implementsMessageBulkPushParamsMessageUnion() {
}

type MessageBulkPushParamsMessagesMqQueueMessageJsonContentType string

const (
	MessageBulkPushParamsMessagesMqQueueMessageJsonContentTypeJson MessageBulkPushParamsMessagesMqQueueMessageJsonContentType = "json"
)

func (r MessageBulkPushParamsMessagesMqQueueMessageJsonContentType) IsKnown() bool {
	switch r {
	case MessageBulkPushParamsMessagesMqQueueMessageJsonContentTypeJson:
		return true
	}
	return false
}

type MessageBulkPushParamsMessagesContentType string

const (
	MessageBulkPushParamsMessagesContentTypeText MessageBulkPushParamsMessagesContentType = "text"
	MessageBulkPushParamsMessagesContentTypeJson MessageBulkPushParamsMessagesContentType = "json"
)

func (r MessageBulkPushParamsMessagesContentType) IsKnown() bool {
	switch r {
	case MessageBulkPushParamsMessagesContentTypeText, MessageBulkPushParamsMessagesContentTypeJson:
		return true
	}
	return false
}

type MessageBulkPushResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors"`
	Messages []string                `json:"messages"`
	Result   MessageBulkPushResponse `json:"result"`
	// Indicates if the API call was successful or not.
	Success MessageBulkPushResponseEnvelopeSuccess `json:"success"`
	JSON    messageBulkPushResponseEnvelopeJSON    `json:"-"`
}

// messageBulkPushResponseEnvelopeJSON contains the JSON metadata for the struct
// [MessageBulkPushResponseEnvelope]
type messageBulkPushResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBulkPushResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBulkPushResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type MessageBulkPushResponseEnvelopeSuccess bool

const (
	MessageBulkPushResponseEnvelopeSuccessTrue MessageBulkPushResponseEnvelopeSuccess = true
)

func (r MessageBulkPushResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MessageBulkPushResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MessagePullParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
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
	Result   MessagePullResponse   `json:"result"`
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

type MessagePushParams struct {
	// A Resource identifier.
	AccountID param.Field[string]        `path:"account_id" api:"required"`
	Body      MessagePushParamsBodyUnion `json:"body"`
}

func (r MessagePushParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MessagePushParamsBody struct {
	Body        param.Field[interface{}]                      `json:"body"`
	ContentType param.Field[MessagePushParamsBodyContentType] `json:"content_type"`
	// The number of seconds to wait for attempting to deliver this message to
	// consumers
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
}

func (r MessagePushParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessagePushParamsBody) implementsMessagePushParamsBodyUnion() {}

// Satisfied by [queues.MessagePushParamsBodyMqQueueMessageText],
// [queues.MessagePushParamsBodyMqQueueMessageJson], [MessagePushParamsBody].
type MessagePushParamsBodyUnion interface {
	implementsMessagePushParamsBodyUnion()
}

type MessagePushParamsBodyMqQueueMessageText struct {
	Body        param.Field[string]                                             `json:"body"`
	ContentType param.Field[MessagePushParamsBodyMqQueueMessageTextContentType] `json:"content_type"`
	// The number of seconds to wait for attempting to deliver this message to
	// consumers
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
}

func (r MessagePushParamsBodyMqQueueMessageText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessagePushParamsBodyMqQueueMessageText) implementsMessagePushParamsBodyUnion() {}

type MessagePushParamsBodyMqQueueMessageTextContentType string

const (
	MessagePushParamsBodyMqQueueMessageTextContentTypeText MessagePushParamsBodyMqQueueMessageTextContentType = "text"
)

func (r MessagePushParamsBodyMqQueueMessageTextContentType) IsKnown() bool {
	switch r {
	case MessagePushParamsBodyMqQueueMessageTextContentTypeText:
		return true
	}
	return false
}

type MessagePushParamsBodyMqQueueMessageJson struct {
	Body        param.Field[interface{}]                                        `json:"body"`
	ContentType param.Field[MessagePushParamsBodyMqQueueMessageJsonContentType] `json:"content_type"`
	// The number of seconds to wait for attempting to deliver this message to
	// consumers
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
}

func (r MessagePushParamsBodyMqQueueMessageJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessagePushParamsBodyMqQueueMessageJson) implementsMessagePushParamsBodyUnion() {}

type MessagePushParamsBodyMqQueueMessageJsonContentType string

const (
	MessagePushParamsBodyMqQueueMessageJsonContentTypeJson MessagePushParamsBodyMqQueueMessageJsonContentType = "json"
)

func (r MessagePushParamsBodyMqQueueMessageJsonContentType) IsKnown() bool {
	switch r {
	case MessagePushParamsBodyMqQueueMessageJsonContentTypeJson:
		return true
	}
	return false
}

type MessagePushParamsBodyContentType string

const (
	MessagePushParamsBodyContentTypeText MessagePushParamsBodyContentType = "text"
	MessagePushParamsBodyContentTypeJson MessagePushParamsBodyContentType = "json"
)

func (r MessagePushParamsBodyContentType) IsKnown() bool {
	switch r {
	case MessagePushParamsBodyContentTypeText, MessagePushParamsBodyContentTypeJson:
		return true
	}
	return false
}

type MessagePushResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   MessagePushResponse   `json:"result"`
	// Indicates if the API call was successful or not.
	Success MessagePushResponseEnvelopeSuccess `json:"success"`
	JSON    messagePushResponseEnvelopeJSON    `json:"-"`
}

// messagePushResponseEnvelopeJSON contains the JSON metadata for the struct
// [MessagePushResponseEnvelope]
type messagePushResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagePushResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePushResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type MessagePushResponseEnvelopeSuccess bool

const (
	MessagePushResponseEnvelopeSuccessTrue MessagePushResponseEnvelopeSuccess = true
)

func (r MessagePushResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MessagePushResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
