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

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// ConsumerService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConsumerService] method instead.
type ConsumerService struct {
	Options []option.RequestOption
}

// NewConsumerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConsumerService(opts ...option.RequestOption) (r *ConsumerService) {
	r = &ConsumerService{}
	r.Options = opts
	return
}

// Creates a consumer for a Queue.
func (r *ConsumerService) New(ctx context.Context, queueID string, params ConsumerNewParams, opts ...option.RequestOption) (res *Consumer, err error) {
	var env ConsumerNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces a Queue consumer, or creates it if it does not exist.
func (r *ConsumerService) Update(ctx context.Context, queueID string, consumerID string, params ConsumerUpdateParams, opts ...option.RequestOption) (res *Consumer, err error) {
	var env ConsumerUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return nil, err
	}
	if consumerID == "" {
		err = errors.New("missing required consumer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers/%s", params.AccountID, queueID, consumerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the consumers configured for a Queue.
func (r *ConsumerService) List(ctx context.Context, queueID string, query ConsumerListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Consumer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers", query.AccountID, queueID)
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

// Returns the consumers configured for a Queue.
func (r *ConsumerService) ListAutoPaging(ctx context.Context, queueID string, query ConsumerListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Consumer] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, queueID, query, opts...))
}

// Deletes a consumer from a Queue.
func (r *ConsumerService) Delete(ctx context.Context, queueID string, consumerID string, body ConsumerDeleteParams, opts ...option.RequestOption) (res *ConsumerDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return nil, err
	}
	if consumerID == "" {
		err = errors.New("missing required consumer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers/%s", body.AccountID, queueID, consumerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns a Queue consumer by identifier.
func (r *ConsumerService) Get(ctx context.Context, queueID string, consumerID string, query ConsumerGetParams, opts ...option.RequestOption) (res *Consumer, err error) {
	var env ConsumerGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return nil, err
	}
	if consumerID == "" {
		err = errors.New("missing required consumer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers/%s", query.AccountID, queueID, consumerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Response body representing a consumer
type Consumer struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string `json:"dead_letter_queue"`
	QueueName       string `json:"queue_name"`
	// Name of a Worker
	ScriptName string `json:"script_name"`
	// This field can have the runtime type of
	// [ConsumerMqWorkerConsumerResponseSettings],
	// [ConsumerMqHTTPConsumerResponseSettings],
	// [ConsumerMqNotificationConsumerResponseSettings].
	Settings interface{}  `json:"settings"`
	Type     ConsumerType `json:"type"`
	JSON     consumerJSON `json:"-"`
	union    ConsumerUnion
}

// consumerJSON contains the JSON metadata for the struct [Consumer]
type consumerJSON struct {
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

func (r consumerJSON) RawJSON() string {
	return r.raw
}

func (r *Consumer) UnmarshalJSON(data []byte) (err error) {
	*r = Consumer{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConsumerUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [ConsumerMqWorkerConsumerResponse],
// [ConsumerMqHTTPConsumerResponse], [ConsumerMqNotificationConsumerResponse].
func (r Consumer) AsUnion() ConsumerUnion {
	return r.union
}

// Response body representing a consumer
//
// Union satisfied by [ConsumerMqWorkerConsumerResponse],
// [ConsumerMqHTTPConsumerResponse] or [ConsumerMqNotificationConsumerResponse].
type ConsumerUnion interface {
	implementsConsumer()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConsumerUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConsumerMqWorkerConsumerResponse{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConsumerMqHTTPConsumerResponse{}),
			DiscriminatorValue: "http_pull",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConsumerMqNotificationConsumerResponse{}),
			DiscriminatorValue: "notification",
		},
	)
}

type ConsumerMqWorkerConsumerResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string `json:"dead_letter_queue"`
	QueueName       string `json:"queue_name"`
	// Name of a Worker
	ScriptName string                                   `json:"script_name"`
	Settings   ConsumerMqWorkerConsumerResponseSettings `json:"settings"`
	Type       ConsumerMqWorkerConsumerResponseType     `json:"type"`
	JSON       consumerMqWorkerConsumerResponseJSON     `json:"-"`
}

// consumerMqWorkerConsumerResponseJSON contains the JSON metadata for the struct
// [ConsumerMqWorkerConsumerResponse]
type consumerMqWorkerConsumerResponseJSON struct {
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

func (r *ConsumerMqWorkerConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerMqWorkerConsumerResponseJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerMqWorkerConsumerResponse) implementsConsumer() {}

type ConsumerMqWorkerConsumerResponseSettings struct {
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
	RetryDelay float64                                      `json:"retry_delay"`
	JSON       consumerMqWorkerConsumerResponseSettingsJSON `json:"-"`
}

// consumerMqWorkerConsumerResponseSettingsJSON contains the JSON metadata for the
// struct [ConsumerMqWorkerConsumerResponseSettings]
type consumerMqWorkerConsumerResponseSettingsJSON struct {
	BatchSize      apijson.Field
	MaxConcurrency apijson.Field
	MaxRetries     apijson.Field
	MaxWaitTimeMs  apijson.Field
	RetryDelay     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConsumerMqWorkerConsumerResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerMqWorkerConsumerResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerMqWorkerConsumerResponseType string

const (
	ConsumerMqWorkerConsumerResponseTypeWorker ConsumerMqWorkerConsumerResponseType = "worker"
)

func (r ConsumerMqWorkerConsumerResponseType) IsKnown() bool {
	switch r {
	case ConsumerMqWorkerConsumerResponseTypeWorker:
		return true
	}
	return false
}

type ConsumerMqHTTPConsumerResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string                                 `json:"dead_letter_queue"`
	QueueName       string                                 `json:"queue_name"`
	Settings        ConsumerMqHTTPConsumerResponseSettings `json:"settings"`
	Type            ConsumerMqHTTPConsumerResponseType     `json:"type"`
	JSON            consumerMqHTTPConsumerResponseJSON     `json:"-"`
}

// consumerMqHTTPConsumerResponseJSON contains the JSON metadata for the struct
// [ConsumerMqHTTPConsumerResponse]
type consumerMqHTTPConsumerResponseJSON struct {
	ConsumerID      apijson.Field
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	QueueName       apijson.Field
	Settings        apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConsumerMqHTTPConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerMqHTTPConsumerResponseJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerMqHTTPConsumerResponse) implementsConsumer() {}

type ConsumerMqHTTPConsumerResponseSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize float64 `json:"batch_size"`
	// The maximum number of retries
	MaxRetries float64 `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay float64 `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs float64                                    `json:"visibility_timeout_ms"`
	JSON                consumerMqHTTPConsumerResponseSettingsJSON `json:"-"`
}

// consumerMqHTTPConsumerResponseSettingsJSON contains the JSON metadata for the
// struct [ConsumerMqHTTPConsumerResponseSettings]
type consumerMqHTTPConsumerResponseSettingsJSON struct {
	BatchSize           apijson.Field
	MaxRetries          apijson.Field
	RetryDelay          apijson.Field
	VisibilityTimeoutMs apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ConsumerMqHTTPConsumerResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerMqHTTPConsumerResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerMqHTTPConsumerResponseType string

const (
	ConsumerMqHTTPConsumerResponseTypeHTTPPull ConsumerMqHTTPConsumerResponseType = "http_pull"
)

func (r ConsumerMqHTTPConsumerResponseType) IsKnown() bool {
	switch r {
	case ConsumerMqHTTPConsumerResponseTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerMqNotificationConsumerResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured.
	DeadLetterQueue string `json:"dead_letter_queue"`
	QueueName       string `json:"queue_name"`
	// Notification destinations for a Queue. At least one email, webhook, or PagerDuty
	// destination is required.
	Settings ConsumerMqNotificationConsumerResponseSettings `json:"settings"`
	Type     ConsumerMqNotificationConsumerResponseType     `json:"type"`
	JSON     consumerMqNotificationConsumerResponseJSON     `json:"-"`
}

// consumerMqNotificationConsumerResponseJSON contains the JSON metadata for the
// struct [ConsumerMqNotificationConsumerResponse]
type consumerMqNotificationConsumerResponseJSON struct {
	ConsumerID      apijson.Field
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	QueueName       apijson.Field
	Settings        apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConsumerMqNotificationConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerMqNotificationConsumerResponseJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerMqNotificationConsumerResponse) implementsConsumer() {}

// Notification destinations for a Queue. At least one email, webhook, or PagerDuty
// destination is required.
type ConsumerMqNotificationConsumerResponseSettings struct {
	// This field can have the runtime type of
	// [[]ConsumerMqNotificationConsumerResponseSettingsObjectEmail].
	Email interface{} `json:"email"`
	// This field can have the runtime type of
	// [[]ConsumerMqNotificationConsumerResponseSettingsObjectPagerduty].
	Pagerduty interface{} `json:"pagerduty"`
	// This field can have the runtime type of
	// [[]ConsumerMqNotificationConsumerResponseSettingsObjectWebhook].
	Webhooks interface{}                                        `json:"webhooks"`
	JSON     consumerMqNotificationConsumerResponseSettingsJSON `json:"-"`
	union    ConsumerMqNotificationConsumerResponseSettingsUnion
}

// consumerMqNotificationConsumerResponseSettingsJSON contains the JSON metadata
// for the struct [ConsumerMqNotificationConsumerResponseSettings]
type consumerMqNotificationConsumerResponseSettingsJSON struct {
	Email       apijson.Field
	Pagerduty   apijson.Field
	Webhooks    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r consumerMqNotificationConsumerResponseSettingsJSON) RawJSON() string {
	return r.raw
}

func (r *ConsumerMqNotificationConsumerResponseSettings) UnmarshalJSON(data []byte) (err error) {
	*r = ConsumerMqNotificationConsumerResponseSettings{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConsumerMqNotificationConsumerResponseSettingsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConsumerMqNotificationConsumerResponseSettingsObject],
// [ConsumerMqNotificationConsumerResponseSettingsObject],
// [ConsumerMqNotificationConsumerResponseSettingsObject].
func (r ConsumerMqNotificationConsumerResponseSettings) AsUnion() ConsumerMqNotificationConsumerResponseSettingsUnion {
	return r.union
}

// Notification destinations for a Queue. At least one email, webhook, or PagerDuty
// destination is required.
//
// Union satisfied by [ConsumerMqNotificationConsumerResponseSettingsObject],
// [ConsumerMqNotificationConsumerResponseSettingsObject] or
// [ConsumerMqNotificationConsumerResponseSettingsObject].
type ConsumerMqNotificationConsumerResponseSettingsUnion interface {
	implementsConsumerMqNotificationConsumerResponseSettings()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConsumerMqNotificationConsumerResponseSettingsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConsumerMqNotificationConsumerResponseSettingsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConsumerMqNotificationConsumerResponseSettingsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConsumerMqNotificationConsumerResponseSettingsObject{}),
		},
	)
}

type ConsumerMqNotificationConsumerResponseSettingsObject struct {
	Email []ConsumerMqNotificationConsumerResponseSettingsObjectEmail `json:"email" api:"required"`
	// PagerDuty notification destinations.
	Pagerduty []ConsumerMqNotificationConsumerResponseSettingsObjectPagerduty `json:"pagerduty"`
	// Webhook notification destinations.
	Webhooks []ConsumerMqNotificationConsumerResponseSettingsObjectWebhook `json:"webhooks"`
	JSON     consumerMqNotificationConsumerResponseSettingsObjectJSON      `json:"-"`
}

// consumerMqNotificationConsumerResponseSettingsObjectJSON contains the JSON
// metadata for the struct [ConsumerMqNotificationConsumerResponseSettingsObject]
type consumerMqNotificationConsumerResponseSettingsObjectJSON struct {
	Email       apijson.Field
	Pagerduty   apijson.Field
	Webhooks    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerMqNotificationConsumerResponseSettingsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerMqNotificationConsumerResponseSettingsObjectJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerMqNotificationConsumerResponseSettingsObject) implementsConsumerMqNotificationConsumerResponseSettings() {
}

type ConsumerMqNotificationConsumerResponseSettingsObjectEmail struct {
	// The email address.
	ID   string                                                        `json:"id" api:"required"`
	JSON consumerMqNotificationConsumerResponseSettingsObjectEmailJSON `json:"-"`
}

// consumerMqNotificationConsumerResponseSettingsObjectEmailJSON contains the JSON
// metadata for the struct
// [ConsumerMqNotificationConsumerResponseSettingsObjectEmail]
type consumerMqNotificationConsumerResponseSettingsObjectEmailJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerMqNotificationConsumerResponseSettingsObjectEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerMqNotificationConsumerResponseSettingsObjectEmailJSON) RawJSON() string {
	return r.raw
}

type ConsumerMqNotificationConsumerResponseSettingsObjectPagerduty struct {
	// UUID.
	ID   string                                                            `json:"id" api:"required"`
	JSON consumerMqNotificationConsumerResponseSettingsObjectPagerdutyJSON `json:"-"`
}

// consumerMqNotificationConsumerResponseSettingsObjectPagerdutyJSON contains the
// JSON metadata for the struct
// [ConsumerMqNotificationConsumerResponseSettingsObjectPagerduty]
type consumerMqNotificationConsumerResponseSettingsObjectPagerdutyJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerMqNotificationConsumerResponseSettingsObjectPagerduty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerMqNotificationConsumerResponseSettingsObjectPagerdutyJSON) RawJSON() string {
	return r.raw
}

type ConsumerMqNotificationConsumerResponseSettingsObjectWebhook struct {
	// UUID.
	ID   string                                                          `json:"id" api:"required"`
	JSON consumerMqNotificationConsumerResponseSettingsObjectWebhookJSON `json:"-"`
}

// consumerMqNotificationConsumerResponseSettingsObjectWebhookJSON contains the
// JSON metadata for the struct
// [ConsumerMqNotificationConsumerResponseSettingsObjectWebhook]
type consumerMqNotificationConsumerResponseSettingsObjectWebhookJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerMqNotificationConsumerResponseSettingsObjectWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerMqNotificationConsumerResponseSettingsObjectWebhookJSON) RawJSON() string {
	return r.raw
}

type ConsumerMqNotificationConsumerResponseType string

const (
	ConsumerMqNotificationConsumerResponseTypeNotification ConsumerMqNotificationConsumerResponseType = "notification"
)

func (r ConsumerMqNotificationConsumerResponseType) IsKnown() bool {
	switch r {
	case ConsumerMqNotificationConsumerResponseTypeNotification:
		return true
	}
	return false
}

type ConsumerType string

const (
	ConsumerTypeWorker       ConsumerType = "worker"
	ConsumerTypeHTTPPull     ConsumerType = "http_pull"
	ConsumerTypeNotification ConsumerType = "notification"
)

func (r ConsumerType) IsKnown() bool {
	switch r {
	case ConsumerTypeWorker, ConsumerTypeHTTPPull, ConsumerTypeNotification:
		return true
	}
	return false
}

// Response body representing a consumer
type ConsumerParam struct {
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue param.Field[string] `json:"dead_letter_queue"`
	QueueName       param.Field[string] `json:"queue_name"`
	// Name of a Worker
	ScriptName param.Field[string]       `json:"script_name"`
	Settings   param.Field[interface{}]  `json:"settings"`
	Type       param.Field[ConsumerType] `json:"type"`
}

func (r ConsumerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerParam) implementsConsumerUnionParam() {}

// Response body representing a consumer
//
// Satisfied by [queues.ConsumerMqWorkerConsumerResponseParam],
// [queues.ConsumerMqHTTPConsumerResponseParam],
// [queues.ConsumerMqNotificationConsumerResponseParam], [ConsumerParam].
type ConsumerUnionParam interface {
	implementsConsumerUnionParam()
}

type ConsumerMqWorkerConsumerResponseParam struct {
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue param.Field[string] `json:"dead_letter_queue"`
	QueueName       param.Field[string] `json:"queue_name"`
	// Name of a Worker
	ScriptName param.Field[string]                                        `json:"script_name"`
	Settings   param.Field[ConsumerMqWorkerConsumerResponseSettingsParam] `json:"settings"`
	Type       param.Field[ConsumerMqWorkerConsumerResponseType]          `json:"type"`
}

func (r ConsumerMqWorkerConsumerResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerMqWorkerConsumerResponseParam) implementsConsumerUnionParam() {}

type ConsumerMqWorkerConsumerResponseSettingsParam struct {
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

func (r ConsumerMqWorkerConsumerResponseSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerMqHTTPConsumerResponseParam struct {
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue param.Field[string]                                      `json:"dead_letter_queue"`
	QueueName       param.Field[string]                                      `json:"queue_name"`
	Settings        param.Field[ConsumerMqHTTPConsumerResponseSettingsParam] `json:"settings"`
	Type            param.Field[ConsumerMqHTTPConsumerResponseType]          `json:"type"`
}

func (r ConsumerMqHTTPConsumerResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerMqHTTPConsumerResponseParam) implementsConsumerUnionParam() {}

type ConsumerMqHTTPConsumerResponseSettingsParam struct {
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

func (r ConsumerMqHTTPConsumerResponseSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerMqNotificationConsumerResponseParam struct {
	// Name of the dead letter queue, or empty string if not configured.
	DeadLetterQueue param.Field[string] `json:"dead_letter_queue"`
	QueueName       param.Field[string] `json:"queue_name"`
	// Notification destinations for a Queue. At least one email, webhook, or PagerDuty
	// destination is required.
	Settings param.Field[ConsumerMqNotificationConsumerResponseSettingsUnionParam] `json:"settings"`
	Type     param.Field[ConsumerMqNotificationConsumerResponseType]               `json:"type"`
}

func (r ConsumerMqNotificationConsumerResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerMqNotificationConsumerResponseParam) implementsConsumerUnionParam() {}

// Notification destinations for a Queue. At least one email, webhook, or PagerDuty
// destination is required.
type ConsumerMqNotificationConsumerResponseSettingsParam struct {
	Email     param.Field[interface{}] `json:"email"`
	Pagerduty param.Field[interface{}] `json:"pagerduty"`
	Webhooks  param.Field[interface{}] `json:"webhooks"`
}

func (r ConsumerMqNotificationConsumerResponseSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerMqNotificationConsumerResponseSettingsParam) implementsConsumerMqNotificationConsumerResponseSettingsUnionParam() {
}

// Notification destinations for a Queue. At least one email, webhook, or PagerDuty
// destination is required.
//
// Satisfied by [queues.ConsumerMqNotificationConsumerResponseSettingsObjectParam],
// [queues.ConsumerMqNotificationConsumerResponseSettingsObjectParam],
// [queues.ConsumerMqNotificationConsumerResponseSettingsObjectParam],
// [ConsumerMqNotificationConsumerResponseSettingsParam].
type ConsumerMqNotificationConsumerResponseSettingsUnionParam interface {
	implementsConsumerMqNotificationConsumerResponseSettingsUnionParam()
}

type ConsumerMqNotificationConsumerResponseSettingsObjectParam struct {
	Email param.Field[[]ConsumerMqNotificationConsumerResponseSettingsObjectEmailParam] `json:"email" api:"required"`
	// PagerDuty notification destinations.
	Pagerduty param.Field[[]ConsumerMqNotificationConsumerResponseSettingsObjectPagerdutyParam] `json:"pagerduty"`
	// Webhook notification destinations.
	Webhooks param.Field[[]ConsumerMqNotificationConsumerResponseSettingsObjectWebhookParam] `json:"webhooks"`
}

func (r ConsumerMqNotificationConsumerResponseSettingsObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerMqNotificationConsumerResponseSettingsObjectParam) implementsConsumerMqNotificationConsumerResponseSettingsUnionParam() {
}

type ConsumerMqNotificationConsumerResponseSettingsObjectEmailParam struct {
	// The email address.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r ConsumerMqNotificationConsumerResponseSettingsObjectEmailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerMqNotificationConsumerResponseSettingsObjectPagerdutyParam struct {
	// UUID.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r ConsumerMqNotificationConsumerResponseSettingsObjectPagerdutyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerMqNotificationConsumerResponseSettingsObjectWebhookParam struct {
	// UUID.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r ConsumerMqNotificationConsumerResponseSettingsObjectWebhookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	// Indicates if the API call was successful or not.
	Success ConsumerDeleteResponseSuccess `json:"success"`
	JSON    consumerDeleteResponseJSON    `json:"-"`
}

// consumerDeleteResponseJSON contains the JSON metadata for the struct
// [ConsumerDeleteResponse]
type consumerDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type ConsumerDeleteResponseSuccess bool

const (
	ConsumerDeleteResponseSuccessTrue ConsumerDeleteResponseSuccess = true
)

func (r ConsumerDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ConsumerDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ConsumerNewParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Request body for creating or updating a consumer
	Body ConsumerNewParamsBodyUnion `json:"body" api:"required"`
}

func (r ConsumerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request body for creating or updating a consumer
type ConsumerNewParamsBody struct {
	Type            param.Field[ConsumerNewParamsBodyType] `json:"type" api:"required"`
	DeadLetterQueue param.Field[string]                    `json:"dead_letter_queue"`
	// Name of a Worker
	ScriptName param.Field[string]      `json:"script_name"`
	Settings   param.Field[interface{}] `json:"settings"`
}

func (r ConsumerNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerNewParamsBody) implementsConsumerNewParamsBodyUnion() {}

// Request body for creating or updating a consumer
//
// Satisfied by [queues.ConsumerNewParamsBodyMqWorkerConsumerRequest],
// [queues.ConsumerNewParamsBodyMqHTTPConsumerRequest],
// [queues.ConsumerNewParamsBodyMqNotificationConsumerRequest],
// [ConsumerNewParamsBody].
type ConsumerNewParamsBodyUnion interface {
	implementsConsumerNewParamsBodyUnion()
}

type ConsumerNewParamsBodyMqWorkerConsumerRequest struct {
	// Name of a Worker
	ScriptName      param.Field[string]                                               `json:"script_name" api:"required"`
	Type            param.Field[ConsumerNewParamsBodyMqWorkerConsumerRequestType]     `json:"type" api:"required"`
	DeadLetterQueue param.Field[string]                                               `json:"dead_letter_queue"`
	Settings        param.Field[ConsumerNewParamsBodyMqWorkerConsumerRequestSettings] `json:"settings"`
}

func (r ConsumerNewParamsBodyMqWorkerConsumerRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerNewParamsBodyMqWorkerConsumerRequest) implementsConsumerNewParamsBodyUnion() {}

type ConsumerNewParamsBodyMqWorkerConsumerRequestType string

const (
	ConsumerNewParamsBodyMqWorkerConsumerRequestTypeWorker ConsumerNewParamsBodyMqWorkerConsumerRequestType = "worker"
)

func (r ConsumerNewParamsBodyMqWorkerConsumerRequestType) IsKnown() bool {
	switch r {
	case ConsumerNewParamsBodyMqWorkerConsumerRequestTypeWorker:
		return true
	}
	return false
}

type ConsumerNewParamsBodyMqWorkerConsumerRequestSettings struct {
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

func (r ConsumerNewParamsBodyMqWorkerConsumerRequestSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerNewParamsBodyMqHTTPConsumerRequest struct {
	Type            param.Field[ConsumerNewParamsBodyMqHTTPConsumerRequestType]     `json:"type" api:"required"`
	DeadLetterQueue param.Field[string]                                             `json:"dead_letter_queue"`
	Settings        param.Field[ConsumerNewParamsBodyMqHTTPConsumerRequestSettings] `json:"settings"`
}

func (r ConsumerNewParamsBodyMqHTTPConsumerRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerNewParamsBodyMqHTTPConsumerRequest) implementsConsumerNewParamsBodyUnion() {}

type ConsumerNewParamsBodyMqHTTPConsumerRequestType string

const (
	ConsumerNewParamsBodyMqHTTPConsumerRequestTypeHTTPPull ConsumerNewParamsBodyMqHTTPConsumerRequestType = "http_pull"
)

func (r ConsumerNewParamsBodyMqHTTPConsumerRequestType) IsKnown() bool {
	switch r {
	case ConsumerNewParamsBodyMqHTTPConsumerRequestTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerNewParamsBodyMqHTTPConsumerRequestSettings struct {
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

func (r ConsumerNewParamsBodyMqHTTPConsumerRequestSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Creates or updates the notification destinations for a Queue. Only one
// notification consumer may exist per Queue. Updates replace all previously
// configured destinations.
type ConsumerNewParamsBodyMqNotificationConsumerRequest struct {
	// Notification destinations for a Queue. At least one email, webhook, or PagerDuty
	// destination is required.
	Settings        param.Field[ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsUnion] `json:"settings" api:"required"`
	Type            param.Field[ConsumerNewParamsBodyMqNotificationConsumerRequestType]          `json:"type" api:"required"`
	DeadLetterQueue param.Field[string]                                                          `json:"dead_letter_queue"`
}

func (r ConsumerNewParamsBodyMqNotificationConsumerRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerNewParamsBodyMqNotificationConsumerRequest) implementsConsumerNewParamsBodyUnion() {}

// Notification destinations for a Queue. At least one email, webhook, or PagerDuty
// destination is required.
type ConsumerNewParamsBodyMqNotificationConsumerRequestSettings struct {
	Email     param.Field[interface{}] `json:"email"`
	Pagerduty param.Field[interface{}] `json:"pagerduty"`
	Webhooks  param.Field[interface{}] `json:"webhooks"`
}

func (r ConsumerNewParamsBodyMqNotificationConsumerRequestSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerNewParamsBodyMqNotificationConsumerRequestSettings) implementsConsumerNewParamsBodyMqNotificationConsumerRequestSettingsUnion() {
}

// Notification destinations for a Queue. At least one email, webhook, or PagerDuty
// destination is required.
//
// Satisfied by
// [queues.ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObject],
// [queues.ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObject],
// [queues.ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObject],
// [ConsumerNewParamsBodyMqNotificationConsumerRequestSettings].
type ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsUnion interface {
	implementsConsumerNewParamsBodyMqNotificationConsumerRequestSettingsUnion()
}

type ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObject struct {
	Email param.Field[[]ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObjectEmail] `json:"email" api:"required"`
	// PagerDuty notification destinations.
	Pagerduty param.Field[[]ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObjectPagerduty] `json:"pagerduty"`
	// Webhook notification destinations.
	Webhooks param.Field[[]ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObjectWebhook] `json:"webhooks"`
}

func (r ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObject) implementsConsumerNewParamsBodyMqNotificationConsumerRequestSettingsUnion() {
}

type ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObjectEmail struct {
	// The email address.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObjectEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObjectPagerduty struct {
	// UUID.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObjectPagerduty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObjectWebhook struct {
	// UUID.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObjectWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerNewParamsBodyMqNotificationConsumerRequestType string

const (
	ConsumerNewParamsBodyMqNotificationConsumerRequestTypeNotification ConsumerNewParamsBodyMqNotificationConsumerRequestType = "notification"
)

func (r ConsumerNewParamsBodyMqNotificationConsumerRequestType) IsKnown() bool {
	switch r {
	case ConsumerNewParamsBodyMqNotificationConsumerRequestTypeNotification:
		return true
	}
	return false
}

type ConsumerNewParamsBodyType string

const (
	ConsumerNewParamsBodyTypeWorker       ConsumerNewParamsBodyType = "worker"
	ConsumerNewParamsBodyTypeHTTPPull     ConsumerNewParamsBodyType = "http_pull"
	ConsumerNewParamsBodyTypeNotification ConsumerNewParamsBodyType = "notification"
)

func (r ConsumerNewParamsBodyType) IsKnown() bool {
	switch r {
	case ConsumerNewParamsBodyTypeWorker, ConsumerNewParamsBodyTypeHTTPPull, ConsumerNewParamsBodyTypeNotification:
		return true
	}
	return false
}

type ConsumerNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	// Response body representing a consumer
	Result Consumer `json:"result"`
	// Indicates if the API call was successful or not.
	Success ConsumerNewResponseEnvelopeSuccess `json:"success"`
	JSON    consumerNewResponseEnvelopeJSON    `json:"-"`
}

// consumerNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConsumerNewResponseEnvelope]
type consumerNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type ConsumerNewResponseEnvelopeSuccess bool

const (
	ConsumerNewResponseEnvelopeSuccessTrue ConsumerNewResponseEnvelopeSuccess = true
)

func (r ConsumerNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConsumerNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConsumerUpdateParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Request body for creating or updating a consumer
	Body ConsumerUpdateParamsBodyUnion `json:"body" api:"required"`
}

func (r ConsumerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request body for creating or updating a consumer
type ConsumerUpdateParamsBody struct {
	Type            param.Field[ConsumerUpdateParamsBodyType] `json:"type" api:"required"`
	DeadLetterQueue param.Field[string]                       `json:"dead_letter_queue"`
	// Name of a Worker
	ScriptName param.Field[string]      `json:"script_name"`
	Settings   param.Field[interface{}] `json:"settings"`
}

func (r ConsumerUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerUpdateParamsBody) implementsConsumerUpdateParamsBodyUnion() {}

// Request body for creating or updating a consumer
//
// Satisfied by [queues.ConsumerUpdateParamsBodyMqWorkerConsumerRequest],
// [queues.ConsumerUpdateParamsBodyMqHTTPConsumerRequest],
// [queues.ConsumerUpdateParamsBodyMqNotificationConsumerRequest],
// [ConsumerUpdateParamsBody].
type ConsumerUpdateParamsBodyUnion interface {
	implementsConsumerUpdateParamsBodyUnion()
}

type ConsumerUpdateParamsBodyMqWorkerConsumerRequest struct {
	// Name of a Worker
	ScriptName      param.Field[string]                                                  `json:"script_name" api:"required"`
	Type            param.Field[ConsumerUpdateParamsBodyMqWorkerConsumerRequestType]     `json:"type" api:"required"`
	DeadLetterQueue param.Field[string]                                                  `json:"dead_letter_queue"`
	Settings        param.Field[ConsumerUpdateParamsBodyMqWorkerConsumerRequestSettings] `json:"settings"`
}

func (r ConsumerUpdateParamsBodyMqWorkerConsumerRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerUpdateParamsBodyMqWorkerConsumerRequest) implementsConsumerUpdateParamsBodyUnion() {}

type ConsumerUpdateParamsBodyMqWorkerConsumerRequestType string

const (
	ConsumerUpdateParamsBodyMqWorkerConsumerRequestTypeWorker ConsumerUpdateParamsBodyMqWorkerConsumerRequestType = "worker"
)

func (r ConsumerUpdateParamsBodyMqWorkerConsumerRequestType) IsKnown() bool {
	switch r {
	case ConsumerUpdateParamsBodyMqWorkerConsumerRequestTypeWorker:
		return true
	}
	return false
}

type ConsumerUpdateParamsBodyMqWorkerConsumerRequestSettings struct {
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

func (r ConsumerUpdateParamsBodyMqWorkerConsumerRequestSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerUpdateParamsBodyMqHTTPConsumerRequest struct {
	Type            param.Field[ConsumerUpdateParamsBodyMqHTTPConsumerRequestType]     `json:"type" api:"required"`
	DeadLetterQueue param.Field[string]                                                `json:"dead_letter_queue"`
	Settings        param.Field[ConsumerUpdateParamsBodyMqHTTPConsumerRequestSettings] `json:"settings"`
}

func (r ConsumerUpdateParamsBodyMqHTTPConsumerRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerUpdateParamsBodyMqHTTPConsumerRequest) implementsConsumerUpdateParamsBodyUnion() {}

type ConsumerUpdateParamsBodyMqHTTPConsumerRequestType string

const (
	ConsumerUpdateParamsBodyMqHTTPConsumerRequestTypeHTTPPull ConsumerUpdateParamsBodyMqHTTPConsumerRequestType = "http_pull"
)

func (r ConsumerUpdateParamsBodyMqHTTPConsumerRequestType) IsKnown() bool {
	switch r {
	case ConsumerUpdateParamsBodyMqHTTPConsumerRequestTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerUpdateParamsBodyMqHTTPConsumerRequestSettings struct {
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

func (r ConsumerUpdateParamsBodyMqHTTPConsumerRequestSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Creates or updates the notification destinations for a Queue. Only one
// notification consumer may exist per Queue. Updates replace all previously
// configured destinations.
type ConsumerUpdateParamsBodyMqNotificationConsumerRequest struct {
	// Notification destinations for a Queue. At least one email, webhook, or PagerDuty
	// destination is required.
	Settings        param.Field[ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsUnion] `json:"settings" api:"required"`
	Type            param.Field[ConsumerUpdateParamsBodyMqNotificationConsumerRequestType]          `json:"type" api:"required"`
	DeadLetterQueue param.Field[string]                                                             `json:"dead_letter_queue"`
}

func (r ConsumerUpdateParamsBodyMqNotificationConsumerRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerUpdateParamsBodyMqNotificationConsumerRequest) implementsConsumerUpdateParamsBodyUnion() {
}

// Notification destinations for a Queue. At least one email, webhook, or PagerDuty
// destination is required.
type ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettings struct {
	Email     param.Field[interface{}] `json:"email"`
	Pagerduty param.Field[interface{}] `json:"pagerduty"`
	Webhooks  param.Field[interface{}] `json:"webhooks"`
}

func (r ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettings) implementsConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsUnion() {
}

// Notification destinations for a Queue. At least one email, webhook, or PagerDuty
// destination is required.
//
// Satisfied by
// [queues.ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObject],
// [queues.ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObject],
// [queues.ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObject],
// [ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettings].
type ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsUnion interface {
	implementsConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsUnion()
}

type ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObject struct {
	Email param.Field[[]ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObjectEmail] `json:"email" api:"required"`
	// PagerDuty notification destinations.
	Pagerduty param.Field[[]ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObjectPagerduty] `json:"pagerduty"`
	// Webhook notification destinations.
	Webhooks param.Field[[]ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObjectWebhook] `json:"webhooks"`
}

func (r ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObject) implementsConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsUnion() {
}

type ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObjectEmail struct {
	// The email address.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObjectEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObjectPagerduty struct {
	// UUID.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObjectPagerduty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObjectWebhook struct {
	// UUID.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObjectWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerUpdateParamsBodyMqNotificationConsumerRequestType string

const (
	ConsumerUpdateParamsBodyMqNotificationConsumerRequestTypeNotification ConsumerUpdateParamsBodyMqNotificationConsumerRequestType = "notification"
)

func (r ConsumerUpdateParamsBodyMqNotificationConsumerRequestType) IsKnown() bool {
	switch r {
	case ConsumerUpdateParamsBodyMqNotificationConsumerRequestTypeNotification:
		return true
	}
	return false
}

type ConsumerUpdateParamsBodyType string

const (
	ConsumerUpdateParamsBodyTypeWorker       ConsumerUpdateParamsBodyType = "worker"
	ConsumerUpdateParamsBodyTypeHTTPPull     ConsumerUpdateParamsBodyType = "http_pull"
	ConsumerUpdateParamsBodyTypeNotification ConsumerUpdateParamsBodyType = "notification"
)

func (r ConsumerUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case ConsumerUpdateParamsBodyTypeWorker, ConsumerUpdateParamsBodyTypeHTTPPull, ConsumerUpdateParamsBodyTypeNotification:
		return true
	}
	return false
}

type ConsumerUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	// Response body representing a consumer
	Result Consumer `json:"result"`
	// Indicates if the API call was successful or not.
	Success ConsumerUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    consumerUpdateResponseEnvelopeJSON    `json:"-"`
}

// consumerUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConsumerUpdateResponseEnvelope]
type consumerUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type ConsumerUpdateResponseEnvelopeSuccess bool

const (
	ConsumerUpdateResponseEnvelopeSuccessTrue ConsumerUpdateResponseEnvelopeSuccess = true
)

func (r ConsumerUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConsumerUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConsumerListParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ConsumerDeleteParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ConsumerGetParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ConsumerGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	// Response body representing a consumer
	Result Consumer `json:"result"`
	// Indicates if the API call was successful or not.
	Success ConsumerGetResponseEnvelopeSuccess `json:"success"`
	JSON    consumerGetResponseEnvelopeJSON    `json:"-"`
}

// consumerGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConsumerGetResponseEnvelope]
type consumerGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type ConsumerGetResponseEnvelopeSuccess bool

const (
	ConsumerGetResponseEnvelopeSuccessTrue ConsumerGetResponseEnvelopeSuccess = true
)

func (r ConsumerGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConsumerGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
