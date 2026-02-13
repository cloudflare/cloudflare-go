// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
	"github.com/tidwall/gjson"
)

// QueueService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueService] method instead.
type QueueService struct {
	Options       []option.RequestOption
	Messages      *MessageService
	Purge         *PurgeService
	Consumers     *ConsumerService
	Subscriptions *SubscriptionService
}

// NewQueueService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQueueService(opts ...option.RequestOption) (r *QueueService) {
	r = &QueueService{}
	r.Options = opts
	r.Messages = NewMessageService(opts...)
	r.Purge = NewPurgeService(opts...)
	r.Consumers = NewConsumerService(opts...)
	r.Subscriptions = NewSubscriptionService(opts...)
	return
}

// Create a new queue
func (r *QueueService) New(ctx context.Context, params QueueNewParams, opts ...option.RequestOption) (res *Queue, err error) {
	var env QueueNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a Queue. Note that this endpoint does not support partial updates. If
// successful, the Queue's configuration is overwritten with the supplied
// configuration.
func (r *QueueService) Update(ctx context.Context, queueID string, params QueueUpdateParams, opts ...option.RequestOption) (res *Queue, err error) {
	var env QueueUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the queues owned by an account.
func (r *QueueService) List(ctx context.Context, query QueueListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Queue], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Returns the queues owned by an account.
func (r *QueueService) ListAutoPaging(ctx context.Context, query QueueListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Queue] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a queue
func (r *QueueService) Delete(ctx context.Context, queueID string, body QueueDeleteParams, opts ...option.RequestOption) (res *QueueDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s", body.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Updates a Queue.
func (r *QueueService) Edit(ctx context.Context, queueID string, params QueueEditParams, opts ...option.RequestOption) (res *Queue, err error) {
	var env QueueEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details about a specific queue.
func (r *QueueService) Get(ctx context.Context, queueID string, query QueueGetParams, opts ...option.RequestOption) (res *Queue, err error) {
	var env QueueGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s", query.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Queue struct {
	Consumers           []QueueConsumer `json:"consumers"`
	ConsumersTotalCount float64         `json:"consumers_total_count"`
	CreatedOn           string          `json:"created_on"`
	ModifiedOn          string          `json:"modified_on"`
	Producers           []QueueProducer `json:"producers"`
	ProducersTotalCount float64         `json:"producers_total_count"`
	QueueID             string          `json:"queue_id"`
	QueueName           string          `json:"queue_name"`
	Settings            QueueSettings   `json:"settings"`
	JSON                queueJSON       `json:"-"`
}

// queueJSON contains the JSON metadata for the struct [Queue]
type queueJSON struct {
	Consumers           apijson.Field
	ConsumersTotalCount apijson.Field
	CreatedOn           apijson.Field
	ModifiedOn          apijson.Field
	Producers           apijson.Field
	ProducersTotalCount apijson.Field
	QueueID             apijson.Field
	QueueName           apijson.Field
	Settings            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Queue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueJSON) RawJSON() string {
	return r.raw
}

// Response body representing a consumer
type QueueConsumer struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string `json:"dead_letter_queue"`
	QueueName       string `json:"queue_name"`
	// Name of a Worker
	ScriptName string `json:"script_name"`
	// This field can have the runtime type of
	// [QueueConsumersMqWorkerConsumerResponseSettings],
	// [QueueConsumersMqHTTPConsumerResponseSettings].
	Settings interface{}        `json:"settings"`
	Type     QueueConsumersType `json:"type"`
	JSON     queueConsumerJSON  `json:"-"`
	union    QueueConsumersUnion
}

// queueConsumerJSON contains the JSON metadata for the struct [QueueConsumer]
type queueConsumerJSON struct {
	ConsumerID      apijson.Field
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	QueueName       apijson.Field
	ScriptName      apijson.Field
	Settings        apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r queueConsumerJSON) RawJSON() string {
	return r.raw
}

func (r *QueueConsumer) UnmarshalJSON(data []byte) (err error) {
	*r = QueueConsumer{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [QueueConsumersUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [QueueConsumersMqWorkerConsumerResponse],
// [QueueConsumersMqHTTPConsumerResponse].
func (r QueueConsumer) AsUnion() QueueConsumersUnion {
	return r.union
}

// Response body representing a consumer
//
// Union satisfied by [QueueConsumersMqWorkerConsumerResponse] or
// [QueueConsumersMqHTTPConsumerResponse].
type QueueConsumersUnion interface {
	implementsQueueConsumer()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*QueueConsumersUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(QueueConsumersMqWorkerConsumerResponse{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(QueueConsumersMqHTTPConsumerResponse{}),
			DiscriminatorValue: "http_pull",
		},
	)
}

type QueueConsumersMqWorkerConsumerResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string `json:"dead_letter_queue"`
	QueueName       string `json:"queue_name"`
	// Name of a Worker
	ScriptName string                                         `json:"script_name"`
	Settings   QueueConsumersMqWorkerConsumerResponseSettings `json:"settings"`
	Type       QueueConsumersMqWorkerConsumerResponseType     `json:"type"`
	JSON       queueConsumersMqWorkerConsumerResponseJSON     `json:"-"`
}

// queueConsumersMqWorkerConsumerResponseJSON contains the JSON metadata for the
// struct [QueueConsumersMqWorkerConsumerResponse]
type queueConsumersMqWorkerConsumerResponseJSON struct {
	ConsumerID      apijson.Field
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	QueueName       apijson.Field
	ScriptName      apijson.Field
	Settings        apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *QueueConsumersMqWorkerConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueConsumersMqWorkerConsumerResponseJSON) RawJSON() string {
	return r.raw
}

func (r QueueConsumersMqWorkerConsumerResponse) implementsQueueConsumer() {}

type QueueConsumersMqWorkerConsumerResponseSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize float64 `json:"batch_size"`
	// Maximum number of concurrent consumers that may consume from this Queue. Set to
	// `null` to automatically opt in to the platform's maximum (recommended).
	MaxConcurrency float64 `json:"max_concurrency"`
	// The maximum number of retries
	MaxRetries float64 `json:"max_retries"`
	// The number of milliseconds to wait for a batch to fill up before attempting to
	// deliver it
	MaxWaitTimeMs float64 `json:"max_wait_time_ms"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay float64                                            `json:"retry_delay"`
	JSON       queueConsumersMqWorkerConsumerResponseSettingsJSON `json:"-"`
}

// queueConsumersMqWorkerConsumerResponseSettingsJSON contains the JSON metadata
// for the struct [QueueConsumersMqWorkerConsumerResponseSettings]
type queueConsumersMqWorkerConsumerResponseSettingsJSON struct {
	BatchSize      apijson.Field
	MaxConcurrency apijson.Field
	MaxRetries     apijson.Field
	MaxWaitTimeMs  apijson.Field
	RetryDelay     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *QueueConsumersMqWorkerConsumerResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueConsumersMqWorkerConsumerResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type QueueConsumersMqWorkerConsumerResponseType string

const (
	QueueConsumersMqWorkerConsumerResponseTypeWorker QueueConsumersMqWorkerConsumerResponseType = "worker"
)

func (r QueueConsumersMqWorkerConsumerResponseType) IsKnown() bool {
	switch r {
	case QueueConsumersMqWorkerConsumerResponseTypeWorker:
		return true
	}
	return false
}

type QueueConsumersMqHTTPConsumerResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string                                       `json:"dead_letter_queue"`
	QueueName       string                                       `json:"queue_name"`
	Settings        QueueConsumersMqHTTPConsumerResponseSettings `json:"settings"`
	Type            QueueConsumersMqHTTPConsumerResponseType     `json:"type"`
	JSON            queueConsumersMqHTTPConsumerResponseJSON     `json:"-"`
}

// queueConsumersMqHTTPConsumerResponseJSON contains the JSON metadata for the
// struct [QueueConsumersMqHTTPConsumerResponse]
type queueConsumersMqHTTPConsumerResponseJSON struct {
	ConsumerID      apijson.Field
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	QueueName       apijson.Field
	Settings        apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *QueueConsumersMqHTTPConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueConsumersMqHTTPConsumerResponseJSON) RawJSON() string {
	return r.raw
}

func (r QueueConsumersMqHTTPConsumerResponse) implementsQueueConsumer() {}

type QueueConsumersMqHTTPConsumerResponseSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize float64 `json:"batch_size"`
	// The maximum number of retries
	MaxRetries float64 `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay float64 `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs float64                                          `json:"visibility_timeout_ms"`
	JSON                queueConsumersMqHTTPConsumerResponseSettingsJSON `json:"-"`
}

// queueConsumersMqHTTPConsumerResponseSettingsJSON contains the JSON metadata for
// the struct [QueueConsumersMqHTTPConsumerResponseSettings]
type queueConsumersMqHTTPConsumerResponseSettingsJSON struct {
	BatchSize           apijson.Field
	MaxRetries          apijson.Field
	RetryDelay          apijson.Field
	VisibilityTimeoutMs apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *QueueConsumersMqHTTPConsumerResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueConsumersMqHTTPConsumerResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type QueueConsumersMqHTTPConsumerResponseType string

const (
	QueueConsumersMqHTTPConsumerResponseTypeHTTPPull QueueConsumersMqHTTPConsumerResponseType = "http_pull"
)

func (r QueueConsumersMqHTTPConsumerResponseType) IsKnown() bool {
	switch r {
	case QueueConsumersMqHTTPConsumerResponseTypeHTTPPull:
		return true
	}
	return false
}

type QueueConsumersType string

const (
	QueueConsumersTypeWorker   QueueConsumersType = "worker"
	QueueConsumersTypeHTTPPull QueueConsumersType = "http_pull"
)

func (r QueueConsumersType) IsKnown() bool {
	switch r {
	case QueueConsumersTypeWorker, QueueConsumersTypeHTTPPull:
		return true
	}
	return false
}

type QueueProducer struct {
	BucketName string             `json:"bucket_name"`
	Script     string             `json:"script"`
	Type       QueueProducersType `json:"type"`
	JSON       queueProducerJSON  `json:"-"`
	union      QueueProducersUnion
}

// queueProducerJSON contains the JSON metadata for the struct [QueueProducer]
type queueProducerJSON struct {
	BucketName  apijson.Field
	Script      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r queueProducerJSON) RawJSON() string {
	return r.raw
}

func (r *QueueProducer) UnmarshalJSON(data []byte) (err error) {
	*r = QueueProducer{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [QueueProducersUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [QueueProducersMqWorkerProducer],
// [QueueProducersMqR2Producer].
func (r QueueProducer) AsUnion() QueueProducersUnion {
	return r.union
}

// Union satisfied by [QueueProducersMqWorkerProducer] or
// [QueueProducersMqR2Producer].
type QueueProducersUnion interface {
	implementsQueueProducer()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*QueueProducersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(QueueProducersMqWorkerProducer{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(QueueProducersMqR2Producer{}),
		},
	)
}

type QueueProducersMqWorkerProducer struct {
	Script string                             `json:"script"`
	Type   QueueProducersMqWorkerProducerType `json:"type"`
	JSON   queueProducersMqWorkerProducerJSON `json:"-"`
}

// queueProducersMqWorkerProducerJSON contains the JSON metadata for the struct
// [QueueProducersMqWorkerProducer]
type queueProducersMqWorkerProducerJSON struct {
	Script      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueProducersMqWorkerProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueProducersMqWorkerProducerJSON) RawJSON() string {
	return r.raw
}

func (r QueueProducersMqWorkerProducer) implementsQueueProducer() {}

type QueueProducersMqWorkerProducerType string

const (
	QueueProducersMqWorkerProducerTypeWorker QueueProducersMqWorkerProducerType = "worker"
)

func (r QueueProducersMqWorkerProducerType) IsKnown() bool {
	switch r {
	case QueueProducersMqWorkerProducerTypeWorker:
		return true
	}
	return false
}

type QueueProducersMqR2Producer struct {
	BucketName string                         `json:"bucket_name"`
	Type       QueueProducersMqR2ProducerType `json:"type"`
	JSON       queueProducersMqR2ProducerJSON `json:"-"`
}

// queueProducersMqR2ProducerJSON contains the JSON metadata for the struct
// [QueueProducersMqR2Producer]
type queueProducersMqR2ProducerJSON struct {
	BucketName  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueProducersMqR2Producer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueProducersMqR2ProducerJSON) RawJSON() string {
	return r.raw
}

func (r QueueProducersMqR2Producer) implementsQueueProducer() {}

type QueueProducersMqR2ProducerType string

const (
	QueueProducersMqR2ProducerTypeR2Bucket QueueProducersMqR2ProducerType = "r2_bucket"
)

func (r QueueProducersMqR2ProducerType) IsKnown() bool {
	switch r {
	case QueueProducersMqR2ProducerTypeR2Bucket:
		return true
	}
	return false
}

type QueueProducersType string

const (
	QueueProducersTypeWorker   QueueProducersType = "worker"
	QueueProducersTypeR2Bucket QueueProducersType = "r2_bucket"
)

func (r QueueProducersType) IsKnown() bool {
	switch r {
	case QueueProducersTypeWorker, QueueProducersTypeR2Bucket:
		return true
	}
	return false
}

type QueueSettings struct {
	// Number of seconds to delay delivery of all messages to consumers.
	DeliveryDelay float64 `json:"delivery_delay"`
	// Indicates if message delivery to consumers is currently paused.
	DeliveryPaused bool `json:"delivery_paused"`
	// Number of seconds after which an unconsumed message will be delayed.
	MessageRetentionPeriod float64           `json:"message_retention_period"`
	JSON                   queueSettingsJSON `json:"-"`
}

// queueSettingsJSON contains the JSON metadata for the struct [QueueSettings]
type queueSettingsJSON struct {
	DeliveryDelay          apijson.Field
	DeliveryPaused         apijson.Field
	MessageRetentionPeriod apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *QueueSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueSettingsJSON) RawJSON() string {
	return r.raw
}

type QueueParam struct {
	QueueName param.Field[string]             `json:"queue_name"`
	Settings  param.Field[QueueSettingsParam] `json:"settings"`
}

func (r QueueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response body representing a consumer
type QueueConsumerParam struct {
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue param.Field[string] `json:"dead_letter_queue"`
	QueueName       param.Field[string] `json:"queue_name"`
	// Name of a Worker
	ScriptName param.Field[string]             `json:"script_name"`
	Settings   param.Field[interface{}]        `json:"settings"`
	Type       param.Field[QueueConsumersType] `json:"type"`
}

func (r QueueConsumerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r QueueConsumerParam) implementsQueueConsumersUnionParam() {}

// Response body representing a consumer
//
// Satisfied by [queues.QueueConsumersMqWorkerConsumerResponseParam],
// [queues.QueueConsumersMqHTTPConsumerResponseParam], [QueueConsumerParam].
type QueueConsumersUnionParam interface {
	implementsQueueConsumersUnionParam()
}

type QueueConsumersMqWorkerConsumerResponseParam struct {
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue param.Field[string] `json:"dead_letter_queue"`
	QueueName       param.Field[string] `json:"queue_name"`
	// Name of a Worker
	ScriptName param.Field[string]                                              `json:"script_name"`
	Settings   param.Field[QueueConsumersMqWorkerConsumerResponseSettingsParam] `json:"settings"`
	Type       param.Field[QueueConsumersMqWorkerConsumerResponseType]          `json:"type"`
}

func (r QueueConsumersMqWorkerConsumerResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r QueueConsumersMqWorkerConsumerResponseParam) implementsQueueConsumersUnionParam() {}

type QueueConsumersMqWorkerConsumerResponseSettingsParam struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// Maximum number of concurrent consumers that may consume from this Queue. Set to
	// `null` to automatically opt in to the platform's maximum (recommended).
	MaxConcurrency param.Field[float64] `json:"max_concurrency"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of milliseconds to wait for a batch to fill up before attempting to
	// deliver it
	MaxWaitTimeMs param.Field[float64] `json:"max_wait_time_ms"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
}

func (r QueueConsumersMqWorkerConsumerResponseSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QueueConsumersMqHTTPConsumerResponseParam struct {
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue param.Field[string]                                            `json:"dead_letter_queue"`
	QueueName       param.Field[string]                                            `json:"queue_name"`
	Settings        param.Field[QueueConsumersMqHTTPConsumerResponseSettingsParam] `json:"settings"`
	Type            param.Field[QueueConsumersMqHTTPConsumerResponseType]          `json:"type"`
}

func (r QueueConsumersMqHTTPConsumerResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r QueueConsumersMqHTTPConsumerResponseParam) implementsQueueConsumersUnionParam() {}

type QueueConsumersMqHTTPConsumerResponseSettingsParam struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs param.Field[float64] `json:"visibility_timeout_ms"`
}

func (r QueueConsumersMqHTTPConsumerResponseSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QueueProducerParam struct {
	BucketName param.Field[string]             `json:"bucket_name"`
	Script     param.Field[string]             `json:"script"`
	Type       param.Field[QueueProducersType] `json:"type"`
}

func (r QueueProducerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r QueueProducerParam) implementsQueueProducersUnionParam() {}

// Satisfied by [queues.QueueProducersMqWorkerProducerParam],
// [queues.QueueProducersMqR2ProducerParam], [QueueProducerParam].
type QueueProducersUnionParam interface {
	implementsQueueProducersUnionParam()
}

type QueueProducersMqWorkerProducerParam struct {
	Script param.Field[string]                             `json:"script"`
	Type   param.Field[QueueProducersMqWorkerProducerType] `json:"type"`
}

func (r QueueProducersMqWorkerProducerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r QueueProducersMqWorkerProducerParam) implementsQueueProducersUnionParam() {}

type QueueProducersMqR2ProducerParam struct {
	BucketName param.Field[string]                         `json:"bucket_name"`
	Type       param.Field[QueueProducersMqR2ProducerType] `json:"type"`
}

func (r QueueProducersMqR2ProducerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r QueueProducersMqR2ProducerParam) implementsQueueProducersUnionParam() {}

type QueueSettingsParam struct {
	// Number of seconds to delay delivery of all messages to consumers.
	DeliveryDelay param.Field[float64] `json:"delivery_delay"`
	// Indicates if message delivery to consumers is currently paused.
	DeliveryPaused param.Field[bool] `json:"delivery_paused"`
	// Number of seconds after which an unconsumed message will be delayed.
	MessageRetentionPeriod param.Field[float64] `json:"message_retention_period"`
}

func (r QueueSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QueueDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	// Indicates if the API call was successful or not.
	Success QueueDeleteResponseSuccess `json:"success"`
	JSON    queueDeleteResponseJSON    `json:"-"`
}

// queueDeleteResponseJSON contains the JSON metadata for the struct
// [QueueDeleteResponse]
type queueDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type QueueDeleteResponseSuccess bool

const (
	QueueDeleteResponseSuccessTrue QueueDeleteResponseSuccess = true
)

func (r QueueDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case QueueDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type QueueNewParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	QueueName param.Field[string] `json:"queue_name,required"`
}

func (r QueueNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QueueNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   Queue                 `json:"result"`
	// Indicates if the API call was successful or not.
	Success QueueNewResponseEnvelopeSuccess `json:"success"`
	JSON    queueNewResponseEnvelopeJSON    `json:"-"`
}

// queueNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [QueueNewResponseEnvelope]
type queueNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type QueueNewResponseEnvelopeSuccess bool

const (
	QueueNewResponseEnvelopeSuccessTrue QueueNewResponseEnvelopeSuccess = true
)

func (r QueueNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case QueueNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type QueueUpdateParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Queue     QueueParam          `json:"queue"`
}

func (r QueueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Queue)
}

type QueueUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   Queue                 `json:"result"`
	// Indicates if the API call was successful or not.
	Success QueueUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    queueUpdateResponseEnvelopeJSON    `json:"-"`
}

// queueUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [QueueUpdateResponseEnvelope]
type queueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type QueueUpdateResponseEnvelopeSuccess bool

const (
	QueueUpdateResponseEnvelopeSuccessTrue QueueUpdateResponseEnvelopeSuccess = true
)

func (r QueueUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case QueueUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type QueueListParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type QueueDeleteParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type QueueEditParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Queue     QueueParam          `json:"queue"`
}

func (r QueueEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Queue)
}

type QueueEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   Queue                 `json:"result"`
	// Indicates if the API call was successful or not.
	Success QueueEditResponseEnvelopeSuccess `json:"success"`
	JSON    queueEditResponseEnvelopeJSON    `json:"-"`
}

// queueEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [QueueEditResponseEnvelope]
type queueEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type QueueEditResponseEnvelopeSuccess bool

const (
	QueueEditResponseEnvelopeSuccessTrue QueueEditResponseEnvelopeSuccess = true
)

func (r QueueEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case QueueEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type QueueGetParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type QueueGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   Queue                 `json:"result"`
	// Indicates if the API call was successful or not.
	Success QueueGetResponseEnvelopeSuccess `json:"success"`
	JSON    queueGetResponseEnvelopeJSON    `json:"-"`
}

// queueGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [QueueGetResponseEnvelope]
type queueGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type QueueGetResponseEnvelopeSuccess bool

const (
	QueueGetResponseEnvelopeSuccessTrue QueueGetResponseEnvelopeSuccess = true
)

func (r QueueGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case QueueGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
