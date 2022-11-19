package cloudflare

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	ErrMissingQueueName         = errors.New("required queue name is missing")
	ErrMissingQueueConsumerName = errors.New("required queue consumer name is missing")
)

type Queue struct {
	ID                  string          `json:"queue_id,omitempty"`
	Name                string          `json:"queue_name,omitempty"`
	CreatedOn           *time.Time      `json:"created_on,omitempty"`
	ModifiedOn          *time.Time      `json:"modified_on,omitempty"`
	ProducersTotalCount int             `json:"producers_total_count,omitempty"`
	Producers           []QueueProducer `json:"producers,omitempty"`
	ConsumersTotalCount int             `json:"consumers_total_count,omitempty"`
	Consumers           []QueueConsumer `json:"consumers,omitempty"`
}

type QueueProducer struct {
	Service     string `json:"service,omitempty"`
	Environment string `json:"environment,omitempty"`
}

type QueueConsumer struct {
	Service         string                `json:"service,omitempty"`
	ScriptName      string                `json:"script_name,omitempty"`
	Environment     string                `json:"environment,omitempty"`
	Settings        QueueConsumerSettings `json:"settings,omitempty"`
	QueueName       string                `json:"queue_name,omitempty"`
	CreatedOn       *time.Time            `json:"created_on,omitempty"`
	DeadLetterQueue string                `json:"dead_letter_queue,omitempty"`
}

type QueueConsumerSettings struct {
	BatchSize   int `json:"batch_size,omitempty"`
	MaxRetires  int `json:"max_retries,omitempty"`
	MaxWaitTime int `json:"max_wait_time_ms,omitempty"`
}

type QueueListResponse struct {
	Response
	ResultInfo `json:"result_info"`
	Result     []Queue `json:"result"`
}

type QueueCreateParams struct {
	Name string `json:"queue_name"`
}

type QueueResponse struct {
	Response
	Result Queue `json:"result"`
}

type QueueListConsumersResponse struct {
	Response
	ResultInfo `json:"result_info"`
	Result     []QueueConsumer `json:"result"`
}

type QueueConsumerResponse struct {
	Response
	Result QueueConsumer `json:"result"`
}

// QueueList returns the queues owned by an account.
//
// API reference: https://api.cloudflare.com/#queue-list-queues
func (api *API) QueueList(ctx context.Context, rc *ResourceContainer) ([]Queue, error) {
	if rc.Identifier == "" {
		return []Queue{}, ErrMissingAccountID
	}
	uri := fmt.Sprintf("/accounts/%s/workers/queues", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []Queue{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}
	var r QueueListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []Queue{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// QueueCreate creates a new queue.
//
// API reference: https://api.cloudflare.com/#queue-create-queue
func (api *API) QueueCreate(ctx context.Context, rc *ResourceContainer, queue QueueCreateParams) (Queue, error) {
	if rc.Identifier == "" {
		return Queue{}, ErrMissingAccountID
	}

	if queue.Name == "" {
		return Queue{}, ErrMissingQueueName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, queue)
	if err != nil {
		return Queue{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r QueueResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return Queue{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// QueueDelete deletes a queue.
//
// API reference: https://api.cloudflare.com/#queue-delete-queue
func (api *API) QueueDelete(ctx context.Context, rc *ResourceContainer, queueName string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}
	if queueName == "" {
		return ErrMissingQueueName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues/%s", rc.Identifier, queueName)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", errMakeRequestError, err)
	}
	return nil
}

// QueueGet returns a queue.
//
// API reference: https://api.cloudflare.com/#queue-get-queue
func (api *API) QueueGet(ctx context.Context, rc *ResourceContainer, queueName string) (Queue, error) {
	if rc.Identifier == "" {
		return Queue{}, ErrMissingAccountID
	}
	if queueName == "" {
		return Queue{}, ErrMissingQueueName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues/%s", rc.Identifier, queueName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return Queue{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r QueueResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return Queue{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// QueueUpdate updates a queue.
//
// API reference: https://api.cloudflare.com/#queue-update-queue
func (api *API) QueueUpdate(ctx context.Context, rc *ResourceContainer, queueName string, queue Queue) (Queue, error) {
	if rc.Identifier == "" {
		return Queue{}, ErrMissingAccountID
	}

	if queueName == "" {
		return Queue{}, ErrMissingQueueName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues/%s", rc.Identifier, queueName)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, queue)
	if err != nil {
		return Queue{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r QueueResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return Queue{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// QueueListConsumers returns the consumers of a queue.
//
// API reference: https://api.cloudflare.com/#queue-list-queue-consumers
func (api *API) QueueListConsumers(ctx context.Context, rc *ResourceContainer, queueName string) ([]QueueConsumer, error) {
	if rc.Identifier == "" {
		return []QueueConsumer{}, ErrMissingAccountID
	}

	if queueName == "" {
		return []QueueConsumer{}, ErrMissingQueueName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues/%s/consumers", rc.Identifier, queueName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []QueueConsumer{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r QueueListConsumersResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []QueueConsumer{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// QueueCreateConsumer creates a new consumer for a queue.
//
// API reference: https://api.cloudflare.com/#queue-create-queue-consumer
func (api *API) QueueCreateConsumer(ctx context.Context, rc *ResourceContainer, queueName string, consumer QueueConsumer) (QueueConsumer, error) {
	if rc.Identifier == "" {
		return QueueConsumer{}, ErrMissingAccountID
	}

	if queueName == "" {
		return QueueConsumer{}, ErrMissingQueueName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues/%s/consumers", rc.Identifier, queueName)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, consumer)
	if err != nil {
		return QueueConsumer{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r QueueConsumerResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return QueueConsumer{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// QueueDeleteConsumer deletes the consumer for a queue..
//
// API reference: https://api.cloudflare.com/#queue-delete-queue-consumer
func (api *API) QueueDeleteConsumer(ctx context.Context, rc *ResourceContainer, queueName, consumerName string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	if queueName == "" {
		return ErrMissingQueueName
	}

	if consumerName == "" {
		return ErrMissingQueueConsumerName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues/%s/consumers/%s", rc.Identifier, queueName, consumerName)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	return nil
}

// QueueUpdateConsumer updates the consumer for a queue, or creates one if it does not exist..
//
// API reference: https://api.cloudflare.com/#queue-update-queue-consumer
func (api *API) QueueUpdateConsumer(ctx context.Context, rc *ResourceContainer, queueName, consumerName string, consumer QueueConsumer) (QueueConsumer, error) {
	if rc.Identifier == "" {
		return QueueConsumer{}, ErrMissingAccountID
	}

	if queueName == "" {
		return QueueConsumer{}, ErrMissingQueueName
	}

	if consumerName == "" {
		return QueueConsumer{}, ErrMissingQueueConsumerName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues/%s/consumers/%s", rc.Identifier, queueName, consumerName)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, consumer)
	if err != nil {
		return QueueConsumer{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r QueueConsumerResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return QueueConsumer{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}
