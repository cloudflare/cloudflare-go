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

// Creates a new consumer for a Queue
func (r *ConsumerService) New(ctx context.Context, queueID string, params ConsumerNewParams, opts ...option.RequestOption) (res *ConsumerNewResponse, err error) {
	var env ConsumerNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the consumer for a queue, or creates one if it does not exist.
func (r *ConsumerService) Update(ctx context.Context, queueID string, consumerID string, params ConsumerUpdateParams, opts ...option.RequestOption) (res *ConsumerUpdateResponse, err error) {
	var env ConsumerUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	if consumerID == "" {
		err = errors.New("missing required consumer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers/%s", params.AccountID, queueID, consumerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the consumers for a Queue
func (r *ConsumerService) List(ctx context.Context, queueID string, query ConsumerListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ConsumerListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
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

// Returns the consumers for a Queue
func (r *ConsumerService) ListAutoPaging(ctx context.Context, queueID string, query ConsumerListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ConsumerListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, queueID, query, opts...))
}

// Deletes the consumer for a queue.
func (r *ConsumerService) Delete(ctx context.Context, queueID string, consumerID string, body ConsumerDeleteParams, opts ...option.RequestOption) (res *ConsumerDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	if consumerID == "" {
		err = errors.New("missing required consumer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers/%s", body.AccountID, queueID, consumerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetches the consumer for a queue by consumer id
func (r *ConsumerService) Get(ctx context.Context, queueID string, consumerID string, query ConsumerGetParams, opts ...option.RequestOption) (res *ConsumerGetResponse, err error) {
	var env ConsumerGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	if consumerID == "" {
		err = errors.New("missing required consumer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers/%s", query.AccountID, queueID, consumerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Response body representing a consumer
type ConsumerNewResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string `json:"dead_letter_queue"`
	QueueName       string `json:"queue_name"`
	// Name of a Worker
	ScriptName string `json:"script_name"`
	// This field can have the runtime type of
	// [ConsumerNewResponseMqWorkerConsumerResponseSettings],
	// [ConsumerNewResponseMqHTTPConsumerResponseSettings].
	Settings interface{}             `json:"settings"`
	Type     ConsumerNewResponseType `json:"type"`
	JSON     consumerNewResponseJSON `json:"-"`
	union    ConsumerNewResponseUnion
}

// consumerNewResponseJSON contains the JSON metadata for the struct
// [ConsumerNewResponse]
type consumerNewResponseJSON struct {
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

func (r consumerNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ConsumerNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ConsumerNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConsumerNewResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [ConsumerNewResponseMqWorkerConsumerResponse],
// [ConsumerNewResponseMqHTTPConsumerResponse].
func (r ConsumerNewResponse) AsUnion() ConsumerNewResponseUnion {
	return r.union
}

// Response body representing a consumer
//
// Union satisfied by [ConsumerNewResponseMqWorkerConsumerResponse] or
// [ConsumerNewResponseMqHTTPConsumerResponse].
type ConsumerNewResponseUnion interface {
	implementsConsumerNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConsumerNewResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConsumerNewResponseMqWorkerConsumerResponse{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConsumerNewResponseMqHTTPConsumerResponse{}),
			DiscriminatorValue: "http_pull",
		},
	)
}

type ConsumerNewResponseMqWorkerConsumerResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string `json:"dead_letter_queue"`
	QueueName       string `json:"queue_name"`
	// Name of a Worker
	ScriptName string                                              `json:"script_name"`
	Settings   ConsumerNewResponseMqWorkerConsumerResponseSettings `json:"settings"`
	Type       ConsumerNewResponseMqWorkerConsumerResponseType     `json:"type"`
	JSON       consumerNewResponseMqWorkerConsumerResponseJSON     `json:"-"`
}

// consumerNewResponseMqWorkerConsumerResponseJSON contains the JSON metadata for
// the struct [ConsumerNewResponseMqWorkerConsumerResponse]
type consumerNewResponseMqWorkerConsumerResponseJSON struct {
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

func (r *ConsumerNewResponseMqWorkerConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerNewResponseMqWorkerConsumerResponseJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerNewResponseMqWorkerConsumerResponse) implementsConsumerNewResponse() {}

type ConsumerNewResponseMqWorkerConsumerResponseSettings struct {
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
	RetryDelay float64                                                 `json:"retry_delay"`
	JSON       consumerNewResponseMqWorkerConsumerResponseSettingsJSON `json:"-"`
}

// consumerNewResponseMqWorkerConsumerResponseSettingsJSON contains the JSON
// metadata for the struct [ConsumerNewResponseMqWorkerConsumerResponseSettings]
type consumerNewResponseMqWorkerConsumerResponseSettingsJSON struct {
	BatchSize      apijson.Field
	MaxConcurrency apijson.Field
	MaxRetries     apijson.Field
	MaxWaitTimeMs  apijson.Field
	RetryDelay     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConsumerNewResponseMqWorkerConsumerResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerNewResponseMqWorkerConsumerResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerNewResponseMqWorkerConsumerResponseType string

const (
	ConsumerNewResponseMqWorkerConsumerResponseTypeWorker ConsumerNewResponseMqWorkerConsumerResponseType = "worker"
)

func (r ConsumerNewResponseMqWorkerConsumerResponseType) IsKnown() bool {
	switch r {
	case ConsumerNewResponseMqWorkerConsumerResponseTypeWorker:
		return true
	}
	return false
}

type ConsumerNewResponseMqHTTPConsumerResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string                                            `json:"dead_letter_queue"`
	QueueName       string                                            `json:"queue_name"`
	Settings        ConsumerNewResponseMqHTTPConsumerResponseSettings `json:"settings"`
	Type            ConsumerNewResponseMqHTTPConsumerResponseType     `json:"type"`
	JSON            consumerNewResponseMqHTTPConsumerResponseJSON     `json:"-"`
}

// consumerNewResponseMqHTTPConsumerResponseJSON contains the JSON metadata for the
// struct [ConsumerNewResponseMqHTTPConsumerResponse]
type consumerNewResponseMqHTTPConsumerResponseJSON struct {
	ConsumerID      apijson.Field
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	QueueName       apijson.Field
	Settings        apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConsumerNewResponseMqHTTPConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerNewResponseMqHTTPConsumerResponseJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerNewResponseMqHTTPConsumerResponse) implementsConsumerNewResponse() {}

type ConsumerNewResponseMqHTTPConsumerResponseSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize float64 `json:"batch_size"`
	// The maximum number of retries
	MaxRetries float64 `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay float64 `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs float64                                               `json:"visibility_timeout_ms"`
	JSON                consumerNewResponseMqHTTPConsumerResponseSettingsJSON `json:"-"`
}

// consumerNewResponseMqHTTPConsumerResponseSettingsJSON contains the JSON metadata
// for the struct [ConsumerNewResponseMqHTTPConsumerResponseSettings]
type consumerNewResponseMqHTTPConsumerResponseSettingsJSON struct {
	BatchSize           apijson.Field
	MaxRetries          apijson.Field
	RetryDelay          apijson.Field
	VisibilityTimeoutMs apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ConsumerNewResponseMqHTTPConsumerResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerNewResponseMqHTTPConsumerResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerNewResponseMqHTTPConsumerResponseType string

const (
	ConsumerNewResponseMqHTTPConsumerResponseTypeHTTPPull ConsumerNewResponseMqHTTPConsumerResponseType = "http_pull"
)

func (r ConsumerNewResponseMqHTTPConsumerResponseType) IsKnown() bool {
	switch r {
	case ConsumerNewResponseMqHTTPConsumerResponseTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerNewResponseType string

const (
	ConsumerNewResponseTypeWorker   ConsumerNewResponseType = "worker"
	ConsumerNewResponseTypeHTTPPull ConsumerNewResponseType = "http_pull"
)

func (r ConsumerNewResponseType) IsKnown() bool {
	switch r {
	case ConsumerNewResponseTypeWorker, ConsumerNewResponseTypeHTTPPull:
		return true
	}
	return false
}

// Response body representing a consumer
type ConsumerUpdateResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string `json:"dead_letter_queue"`
	QueueName       string `json:"queue_name"`
	// Name of a Worker
	ScriptName string `json:"script_name"`
	// This field can have the runtime type of
	// [ConsumerUpdateResponseMqWorkerConsumerResponseSettings],
	// [ConsumerUpdateResponseMqHTTPConsumerResponseSettings].
	Settings interface{}                `json:"settings"`
	Type     ConsumerUpdateResponseType `json:"type"`
	JSON     consumerUpdateResponseJSON `json:"-"`
	union    ConsumerUpdateResponseUnion
}

// consumerUpdateResponseJSON contains the JSON metadata for the struct
// [ConsumerUpdateResponse]
type consumerUpdateResponseJSON struct {
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

func (r consumerUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ConsumerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ConsumerUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConsumerUpdateResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConsumerUpdateResponseMqWorkerConsumerResponse],
// [ConsumerUpdateResponseMqHTTPConsumerResponse].
func (r ConsumerUpdateResponse) AsUnion() ConsumerUpdateResponseUnion {
	return r.union
}

// Response body representing a consumer
//
// Union satisfied by [ConsumerUpdateResponseMqWorkerConsumerResponse] or
// [ConsumerUpdateResponseMqHTTPConsumerResponse].
type ConsumerUpdateResponseUnion interface {
	implementsConsumerUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConsumerUpdateResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConsumerUpdateResponseMqWorkerConsumerResponse{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConsumerUpdateResponseMqHTTPConsumerResponse{}),
			DiscriminatorValue: "http_pull",
		},
	)
}

type ConsumerUpdateResponseMqWorkerConsumerResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string `json:"dead_letter_queue"`
	QueueName       string `json:"queue_name"`
	// Name of a Worker
	ScriptName string                                                 `json:"script_name"`
	Settings   ConsumerUpdateResponseMqWorkerConsumerResponseSettings `json:"settings"`
	Type       ConsumerUpdateResponseMqWorkerConsumerResponseType     `json:"type"`
	JSON       consumerUpdateResponseMqWorkerConsumerResponseJSON     `json:"-"`
}

// consumerUpdateResponseMqWorkerConsumerResponseJSON contains the JSON metadata
// for the struct [ConsumerUpdateResponseMqWorkerConsumerResponse]
type consumerUpdateResponseMqWorkerConsumerResponseJSON struct {
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

func (r *ConsumerUpdateResponseMqWorkerConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerUpdateResponseMqWorkerConsumerResponseJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerUpdateResponseMqWorkerConsumerResponse) implementsConsumerUpdateResponse() {}

type ConsumerUpdateResponseMqWorkerConsumerResponseSettings struct {
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
	RetryDelay float64                                                    `json:"retry_delay"`
	JSON       consumerUpdateResponseMqWorkerConsumerResponseSettingsJSON `json:"-"`
}

// consumerUpdateResponseMqWorkerConsumerResponseSettingsJSON contains the JSON
// metadata for the struct [ConsumerUpdateResponseMqWorkerConsumerResponseSettings]
type consumerUpdateResponseMqWorkerConsumerResponseSettingsJSON struct {
	BatchSize      apijson.Field
	MaxConcurrency apijson.Field
	MaxRetries     apijson.Field
	MaxWaitTimeMs  apijson.Field
	RetryDelay     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConsumerUpdateResponseMqWorkerConsumerResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerUpdateResponseMqWorkerConsumerResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerUpdateResponseMqWorkerConsumerResponseType string

const (
	ConsumerUpdateResponseMqWorkerConsumerResponseTypeWorker ConsumerUpdateResponseMqWorkerConsumerResponseType = "worker"
)

func (r ConsumerUpdateResponseMqWorkerConsumerResponseType) IsKnown() bool {
	switch r {
	case ConsumerUpdateResponseMqWorkerConsumerResponseTypeWorker:
		return true
	}
	return false
}

type ConsumerUpdateResponseMqHTTPConsumerResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string                                               `json:"dead_letter_queue"`
	QueueName       string                                               `json:"queue_name"`
	Settings        ConsumerUpdateResponseMqHTTPConsumerResponseSettings `json:"settings"`
	Type            ConsumerUpdateResponseMqHTTPConsumerResponseType     `json:"type"`
	JSON            consumerUpdateResponseMqHTTPConsumerResponseJSON     `json:"-"`
}

// consumerUpdateResponseMqHTTPConsumerResponseJSON contains the JSON metadata for
// the struct [ConsumerUpdateResponseMqHTTPConsumerResponse]
type consumerUpdateResponseMqHTTPConsumerResponseJSON struct {
	ConsumerID      apijson.Field
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	QueueName       apijson.Field
	Settings        apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConsumerUpdateResponseMqHTTPConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerUpdateResponseMqHTTPConsumerResponseJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerUpdateResponseMqHTTPConsumerResponse) implementsConsumerUpdateResponse() {}

type ConsumerUpdateResponseMqHTTPConsumerResponseSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize float64 `json:"batch_size"`
	// The maximum number of retries
	MaxRetries float64 `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay float64 `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs float64                                                  `json:"visibility_timeout_ms"`
	JSON                consumerUpdateResponseMqHTTPConsumerResponseSettingsJSON `json:"-"`
}

// consumerUpdateResponseMqHTTPConsumerResponseSettingsJSON contains the JSON
// metadata for the struct [ConsumerUpdateResponseMqHTTPConsumerResponseSettings]
type consumerUpdateResponseMqHTTPConsumerResponseSettingsJSON struct {
	BatchSize           apijson.Field
	MaxRetries          apijson.Field
	RetryDelay          apijson.Field
	VisibilityTimeoutMs apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ConsumerUpdateResponseMqHTTPConsumerResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerUpdateResponseMqHTTPConsumerResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerUpdateResponseMqHTTPConsumerResponseType string

const (
	ConsumerUpdateResponseMqHTTPConsumerResponseTypeHTTPPull ConsumerUpdateResponseMqHTTPConsumerResponseType = "http_pull"
)

func (r ConsumerUpdateResponseMqHTTPConsumerResponseType) IsKnown() bool {
	switch r {
	case ConsumerUpdateResponseMqHTTPConsumerResponseTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerUpdateResponseType string

const (
	ConsumerUpdateResponseTypeWorker   ConsumerUpdateResponseType = "worker"
	ConsumerUpdateResponseTypeHTTPPull ConsumerUpdateResponseType = "http_pull"
)

func (r ConsumerUpdateResponseType) IsKnown() bool {
	switch r {
	case ConsumerUpdateResponseTypeWorker, ConsumerUpdateResponseTypeHTTPPull:
		return true
	}
	return false
}

// Response body representing a consumer
type ConsumerListResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string `json:"dead_letter_queue"`
	QueueName       string `json:"queue_name"`
	// Name of a Worker
	ScriptName string `json:"script_name"`
	// This field can have the runtime type of
	// [ConsumerListResponseMqWorkerConsumerResponseSettings],
	// [ConsumerListResponseMqHTTPConsumerResponseSettings].
	Settings interface{}              `json:"settings"`
	Type     ConsumerListResponseType `json:"type"`
	JSON     consumerListResponseJSON `json:"-"`
	union    ConsumerListResponseUnion
}

// consumerListResponseJSON contains the JSON metadata for the struct
// [ConsumerListResponse]
type consumerListResponseJSON struct {
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

func (r consumerListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ConsumerListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ConsumerListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConsumerListResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConsumerListResponseMqWorkerConsumerResponse],
// [ConsumerListResponseMqHTTPConsumerResponse].
func (r ConsumerListResponse) AsUnion() ConsumerListResponseUnion {
	return r.union
}

// Response body representing a consumer
//
// Union satisfied by [ConsumerListResponseMqWorkerConsumerResponse] or
// [ConsumerListResponseMqHTTPConsumerResponse].
type ConsumerListResponseUnion interface {
	implementsConsumerListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConsumerListResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConsumerListResponseMqWorkerConsumerResponse{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConsumerListResponseMqHTTPConsumerResponse{}),
			DiscriminatorValue: "http_pull",
		},
	)
}

type ConsumerListResponseMqWorkerConsumerResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string `json:"dead_letter_queue"`
	QueueName       string `json:"queue_name"`
	// Name of a Worker
	ScriptName string                                               `json:"script_name"`
	Settings   ConsumerListResponseMqWorkerConsumerResponseSettings `json:"settings"`
	Type       ConsumerListResponseMqWorkerConsumerResponseType     `json:"type"`
	JSON       consumerListResponseMqWorkerConsumerResponseJSON     `json:"-"`
}

// consumerListResponseMqWorkerConsumerResponseJSON contains the JSON metadata for
// the struct [ConsumerListResponseMqWorkerConsumerResponse]
type consumerListResponseMqWorkerConsumerResponseJSON struct {
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

func (r *ConsumerListResponseMqWorkerConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerListResponseMqWorkerConsumerResponseJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerListResponseMqWorkerConsumerResponse) implementsConsumerListResponse() {}

type ConsumerListResponseMqWorkerConsumerResponseSettings struct {
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
	RetryDelay float64                                                  `json:"retry_delay"`
	JSON       consumerListResponseMqWorkerConsumerResponseSettingsJSON `json:"-"`
}

// consumerListResponseMqWorkerConsumerResponseSettingsJSON contains the JSON
// metadata for the struct [ConsumerListResponseMqWorkerConsumerResponseSettings]
type consumerListResponseMqWorkerConsumerResponseSettingsJSON struct {
	BatchSize      apijson.Field
	MaxConcurrency apijson.Field
	MaxRetries     apijson.Field
	MaxWaitTimeMs  apijson.Field
	RetryDelay     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConsumerListResponseMqWorkerConsumerResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerListResponseMqWorkerConsumerResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerListResponseMqWorkerConsumerResponseType string

const (
	ConsumerListResponseMqWorkerConsumerResponseTypeWorker ConsumerListResponseMqWorkerConsumerResponseType = "worker"
)

func (r ConsumerListResponseMqWorkerConsumerResponseType) IsKnown() bool {
	switch r {
	case ConsumerListResponseMqWorkerConsumerResponseTypeWorker:
		return true
	}
	return false
}

type ConsumerListResponseMqHTTPConsumerResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string                                             `json:"dead_letter_queue"`
	QueueName       string                                             `json:"queue_name"`
	Settings        ConsumerListResponseMqHTTPConsumerResponseSettings `json:"settings"`
	Type            ConsumerListResponseMqHTTPConsumerResponseType     `json:"type"`
	JSON            consumerListResponseMqHTTPConsumerResponseJSON     `json:"-"`
}

// consumerListResponseMqHTTPConsumerResponseJSON contains the JSON metadata for
// the struct [ConsumerListResponseMqHTTPConsumerResponse]
type consumerListResponseMqHTTPConsumerResponseJSON struct {
	ConsumerID      apijson.Field
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	QueueName       apijson.Field
	Settings        apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConsumerListResponseMqHTTPConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerListResponseMqHTTPConsumerResponseJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerListResponseMqHTTPConsumerResponse) implementsConsumerListResponse() {}

type ConsumerListResponseMqHTTPConsumerResponseSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize float64 `json:"batch_size"`
	// The maximum number of retries
	MaxRetries float64 `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay float64 `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs float64                                                `json:"visibility_timeout_ms"`
	JSON                consumerListResponseMqHTTPConsumerResponseSettingsJSON `json:"-"`
}

// consumerListResponseMqHTTPConsumerResponseSettingsJSON contains the JSON
// metadata for the struct [ConsumerListResponseMqHTTPConsumerResponseSettings]
type consumerListResponseMqHTTPConsumerResponseSettingsJSON struct {
	BatchSize           apijson.Field
	MaxRetries          apijson.Field
	RetryDelay          apijson.Field
	VisibilityTimeoutMs apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ConsumerListResponseMqHTTPConsumerResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerListResponseMqHTTPConsumerResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerListResponseMqHTTPConsumerResponseType string

const (
	ConsumerListResponseMqHTTPConsumerResponseTypeHTTPPull ConsumerListResponseMqHTTPConsumerResponseType = "http_pull"
)

func (r ConsumerListResponseMqHTTPConsumerResponseType) IsKnown() bool {
	switch r {
	case ConsumerListResponseMqHTTPConsumerResponseTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerListResponseType string

const (
	ConsumerListResponseTypeWorker   ConsumerListResponseType = "worker"
	ConsumerListResponseTypeHTTPPull ConsumerListResponseType = "http_pull"
)

func (r ConsumerListResponseType) IsKnown() bool {
	switch r {
	case ConsumerListResponseTypeWorker, ConsumerListResponseTypeHTTPPull:
		return true
	}
	return false
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

// Response body representing a consumer
type ConsumerGetResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string `json:"dead_letter_queue"`
	QueueName       string `json:"queue_name"`
	// Name of a Worker
	ScriptName string `json:"script_name"`
	// This field can have the runtime type of
	// [ConsumerGetResponseMqWorkerConsumerResponseSettings],
	// [ConsumerGetResponseMqHTTPConsumerResponseSettings].
	Settings interface{}             `json:"settings"`
	Type     ConsumerGetResponseType `json:"type"`
	JSON     consumerGetResponseJSON `json:"-"`
	union    ConsumerGetResponseUnion
}

// consumerGetResponseJSON contains the JSON metadata for the struct
// [ConsumerGetResponse]
type consumerGetResponseJSON struct {
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

func (r consumerGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ConsumerGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ConsumerGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConsumerGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [ConsumerGetResponseMqWorkerConsumerResponse],
// [ConsumerGetResponseMqHTTPConsumerResponse].
func (r ConsumerGetResponse) AsUnion() ConsumerGetResponseUnion {
	return r.union
}

// Response body representing a consumer
//
// Union satisfied by [ConsumerGetResponseMqWorkerConsumerResponse] or
// [ConsumerGetResponseMqHTTPConsumerResponse].
type ConsumerGetResponseUnion interface {
	implementsConsumerGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConsumerGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConsumerGetResponseMqWorkerConsumerResponse{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConsumerGetResponseMqHTTPConsumerResponse{}),
			DiscriminatorValue: "http_pull",
		},
	)
}

type ConsumerGetResponseMqWorkerConsumerResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string `json:"dead_letter_queue"`
	QueueName       string `json:"queue_name"`
	// Name of a Worker
	ScriptName string                                              `json:"script_name"`
	Settings   ConsumerGetResponseMqWorkerConsumerResponseSettings `json:"settings"`
	Type       ConsumerGetResponseMqWorkerConsumerResponseType     `json:"type"`
	JSON       consumerGetResponseMqWorkerConsumerResponseJSON     `json:"-"`
}

// consumerGetResponseMqWorkerConsumerResponseJSON contains the JSON metadata for
// the struct [ConsumerGetResponseMqWorkerConsumerResponse]
type consumerGetResponseMqWorkerConsumerResponseJSON struct {
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

func (r *ConsumerGetResponseMqWorkerConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerGetResponseMqWorkerConsumerResponseJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerGetResponseMqWorkerConsumerResponse) implementsConsumerGetResponse() {}

type ConsumerGetResponseMqWorkerConsumerResponseSettings struct {
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
	RetryDelay float64                                                 `json:"retry_delay"`
	JSON       consumerGetResponseMqWorkerConsumerResponseSettingsJSON `json:"-"`
}

// consumerGetResponseMqWorkerConsumerResponseSettingsJSON contains the JSON
// metadata for the struct [ConsumerGetResponseMqWorkerConsumerResponseSettings]
type consumerGetResponseMqWorkerConsumerResponseSettingsJSON struct {
	BatchSize      apijson.Field
	MaxConcurrency apijson.Field
	MaxRetries     apijson.Field
	MaxWaitTimeMs  apijson.Field
	RetryDelay     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConsumerGetResponseMqWorkerConsumerResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerGetResponseMqWorkerConsumerResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerGetResponseMqWorkerConsumerResponseType string

const (
	ConsumerGetResponseMqWorkerConsumerResponseTypeWorker ConsumerGetResponseMqWorkerConsumerResponseType = "worker"
)

func (r ConsumerGetResponseMqWorkerConsumerResponseType) IsKnown() bool {
	switch r {
	case ConsumerGetResponseMqWorkerConsumerResponseTypeWorker:
		return true
	}
	return false
}

type ConsumerGetResponseMqHTTPConsumerResponse struct {
	// A Resource identifier.
	ConsumerID string    `json:"consumer_id"`
	CreatedOn  time.Time `json:"created_on" format:"date-time"`
	// Name of the dead letter queue, or empty string if not configured
	DeadLetterQueue string                                            `json:"dead_letter_queue"`
	QueueName       string                                            `json:"queue_name"`
	Settings        ConsumerGetResponseMqHTTPConsumerResponseSettings `json:"settings"`
	Type            ConsumerGetResponseMqHTTPConsumerResponseType     `json:"type"`
	JSON            consumerGetResponseMqHTTPConsumerResponseJSON     `json:"-"`
}

// consumerGetResponseMqHTTPConsumerResponseJSON contains the JSON metadata for the
// struct [ConsumerGetResponseMqHTTPConsumerResponse]
type consumerGetResponseMqHTTPConsumerResponseJSON struct {
	ConsumerID      apijson.Field
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	QueueName       apijson.Field
	Settings        apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConsumerGetResponseMqHTTPConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerGetResponseMqHTTPConsumerResponseJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerGetResponseMqHTTPConsumerResponse) implementsConsumerGetResponse() {}

type ConsumerGetResponseMqHTTPConsumerResponseSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize float64 `json:"batch_size"`
	// The maximum number of retries
	MaxRetries float64 `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay float64 `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs float64                                               `json:"visibility_timeout_ms"`
	JSON                consumerGetResponseMqHTTPConsumerResponseSettingsJSON `json:"-"`
}

// consumerGetResponseMqHTTPConsumerResponseSettingsJSON contains the JSON metadata
// for the struct [ConsumerGetResponseMqHTTPConsumerResponseSettings]
type consumerGetResponseMqHTTPConsumerResponseSettingsJSON struct {
	BatchSize           apijson.Field
	MaxRetries          apijson.Field
	RetryDelay          apijson.Field
	VisibilityTimeoutMs apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ConsumerGetResponseMqHTTPConsumerResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerGetResponseMqHTTPConsumerResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerGetResponseMqHTTPConsumerResponseType string

const (
	ConsumerGetResponseMqHTTPConsumerResponseTypeHTTPPull ConsumerGetResponseMqHTTPConsumerResponseType = "http_pull"
)

func (r ConsumerGetResponseMqHTTPConsumerResponseType) IsKnown() bool {
	switch r {
	case ConsumerGetResponseMqHTTPConsumerResponseTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerGetResponseType string

const (
	ConsumerGetResponseTypeWorker   ConsumerGetResponseType = "worker"
	ConsumerGetResponseTypeHTTPPull ConsumerGetResponseType = "http_pull"
)

func (r ConsumerGetResponseType) IsKnown() bool {
	switch r {
	case ConsumerGetResponseTypeWorker, ConsumerGetResponseTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerNewParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Request body for creating or updating a consumer
	Body ConsumerNewParamsBodyUnion `json:"body"`
}

func (r ConsumerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request body for creating or updating a consumer
type ConsumerNewParamsBody struct {
	Type            param.Field[ConsumerNewParamsBodyType] `json:"type,required"`
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
// [queues.ConsumerNewParamsBodyMqHTTPConsumerRequest], [ConsumerNewParamsBody].
type ConsumerNewParamsBodyUnion interface {
	implementsConsumerNewParamsBodyUnion()
}

type ConsumerNewParamsBodyMqWorkerConsumerRequest struct {
	// Name of a Worker
	ScriptName      param.Field[string]                                               `json:"script_name,required"`
	Type            param.Field[ConsumerNewParamsBodyMqWorkerConsumerRequestType]     `json:"type,required"`
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
	Type            param.Field[ConsumerNewParamsBodyMqHTTPConsumerRequestType]     `json:"type,required"`
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

type ConsumerNewParamsBodyType string

const (
	ConsumerNewParamsBodyTypeWorker   ConsumerNewParamsBodyType = "worker"
	ConsumerNewParamsBodyTypeHTTPPull ConsumerNewParamsBodyType = "http_pull"
)

func (r ConsumerNewParamsBodyType) IsKnown() bool {
	switch r {
	case ConsumerNewParamsBodyTypeWorker, ConsumerNewParamsBodyTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	// Response body representing a consumer
	Result ConsumerNewResponse `json:"result"`
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
	AccountID param.Field[string] `path:"account_id,required"`
	// Request body for creating or updating a consumer
	Body ConsumerUpdateParamsBodyUnion `json:"body,required"`
}

func (r ConsumerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request body for creating or updating a consumer
type ConsumerUpdateParamsBody struct {
	Type            param.Field[ConsumerUpdateParamsBodyType] `json:"type,required"`
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
// [ConsumerUpdateParamsBody].
type ConsumerUpdateParamsBodyUnion interface {
	implementsConsumerUpdateParamsBodyUnion()
}

type ConsumerUpdateParamsBodyMqWorkerConsumerRequest struct {
	// Name of a Worker
	ScriptName      param.Field[string]                                                  `json:"script_name,required"`
	Type            param.Field[ConsumerUpdateParamsBodyMqWorkerConsumerRequestType]     `json:"type,required"`
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
	Type            param.Field[ConsumerUpdateParamsBodyMqHTTPConsumerRequestType]     `json:"type,required"`
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

type ConsumerUpdateParamsBodyType string

const (
	ConsumerUpdateParamsBodyTypeWorker   ConsumerUpdateParamsBodyType = "worker"
	ConsumerUpdateParamsBodyTypeHTTPPull ConsumerUpdateParamsBodyType = "http_pull"
)

func (r ConsumerUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case ConsumerUpdateParamsBodyTypeWorker, ConsumerUpdateParamsBodyTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	// Response body representing a consumer
	Result ConsumerUpdateResponse `json:"result"`
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConsumerDeleteParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConsumerGetParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConsumerGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	// Response body representing a consumer
	Result ConsumerGetResponse `json:"result"`
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
