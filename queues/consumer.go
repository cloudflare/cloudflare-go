// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
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

// Creates a new consumer for a queue.
func (r *ConsumerService) New(ctx context.Context, queueID string, params ConsumerNewParams, opts ...option.RequestOption) (res *ConsumerNewResponse, err error) {
	var env ConsumerNewResponseEnvelope
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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

// Deletes the consumer for a queue.
func (r *ConsumerService) Delete(ctx context.Context, queueID string, consumerID string, body ConsumerDeleteParams, opts ...option.RequestOption) (res *ConsumerDeleteResponseUnion, err error) {
	var env ConsumerDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the consumers for a queue.
func (r *ConsumerService) Get(ctx context.Context, queueID string, query ConsumerGetParams, opts ...option.RequestOption) (res *[]Consumer, err error) {
	var env ConsumerGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers", query.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Consumer struct {
	CreatedOn   string           `json:"created_on"`
	Environment string           `json:"environment"`
	QueueName   string           `json:"queue_name"`
	Service     string           `json:"service"`
	Settings    ConsumerSettings `json:"settings"`
	JSON        consumerJSON     `json:"-"`
}

// consumerJSON contains the JSON metadata for the struct [Consumer]
type consumerJSON struct {
	CreatedOn   apijson.Field
	Environment apijson.Field
	QueueName   apijson.Field
	Service     apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Consumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerJSON) RawJSON() string {
	return r.raw
}

type ConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize float64 `json:"batch_size"`
	// The maximum number of retries
	MaxRetries    float64              `json:"max_retries"`
	MaxWaitTimeMs float64              `json:"max_wait_time_ms"`
	JSON          consumerSettingsJSON `json:"-"`
}

// consumerSettingsJSON contains the JSON metadata for the struct
// [ConsumerSettings]
type consumerSettingsJSON struct {
	BatchSize     apijson.Field
	MaxRetries    apijson.Field
	MaxWaitTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ConsumerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerNewResponse struct {
	CreatedOn string `json:"created_on"`
	// The name of the dead letter queue.
	DeadLetterQueue string                      `json:"dead_letter_queue"`
	Environment     string                      `json:"environment"`
	QueueName       string                      `json:"queue_name"`
	ScriptName      string                      `json:"script_name"`
	Settings        ConsumerNewResponseSettings `json:"settings"`
	JSON            consumerNewResponseJSON     `json:"-"`
}

// consumerNewResponseJSON contains the JSON metadata for the struct
// [ConsumerNewResponse]
type consumerNewResponseJSON struct {
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	Environment     apijson.Field
	QueueName       apijson.Field
	ScriptName      apijson.Field
	Settings        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConsumerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerNewResponseJSON) RawJSON() string {
	return r.raw
}

type ConsumerNewResponseSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize float64 `json:"batch_size"`
	// The maximum number of retries
	MaxRetries    float64                         `json:"max_retries"`
	MaxWaitTimeMs float64                         `json:"max_wait_time_ms"`
	JSON          consumerNewResponseSettingsJSON `json:"-"`
}

// consumerNewResponseSettingsJSON contains the JSON metadata for the struct
// [ConsumerNewResponseSettings]
type consumerNewResponseSettingsJSON struct {
	BatchSize     apijson.Field
	MaxRetries    apijson.Field
	MaxWaitTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ConsumerNewResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerNewResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerUpdateResponse struct {
	CreatedOn       string                         `json:"created_on"`
	DeadLetterQueue string                         `json:"dead_letter_queue"`
	Environment     string                         `json:"environment"`
	QueueName       string                         `json:"queue_name"`
	ScriptName      string                         `json:"script_name"`
	Settings        ConsumerUpdateResponseSettings `json:"settings"`
	JSON            consumerUpdateResponseJSON     `json:"-"`
}

// consumerUpdateResponseJSON contains the JSON metadata for the struct
// [ConsumerUpdateResponse]
type consumerUpdateResponseJSON struct {
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	Environment     apijson.Field
	QueueName       apijson.Field
	ScriptName      apijson.Field
	Settings        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConsumerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ConsumerUpdateResponseSettings struct {
	BatchSize float64 `json:"batch_size"`
	// The maximum number of retries
	MaxRetries    float64                            `json:"max_retries"`
	MaxWaitTimeMs float64                            `json:"max_wait_time_ms"`
	JSON          consumerUpdateResponseSettingsJSON `json:"-"`
}

// consumerUpdateResponseSettingsJSON contains the JSON metadata for the struct
// [ConsumerUpdateResponseSettings]
type consumerUpdateResponseSettingsJSON struct {
	BatchSize     apijson.Field
	MaxRetries    apijson.Field
	MaxWaitTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ConsumerUpdateResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerUpdateResponseSettingsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [queues.ConsumerDeleteResponseUnknown],
// [queues.ConsumerDeleteResponseArray] or [shared.UnionString].
type ConsumerDeleteResponseUnion interface {
	ImplementsQueuesConsumerDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConsumerDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConsumerDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ConsumerDeleteResponseArray []interface{}

func (r ConsumerDeleteResponseArray) ImplementsQueuesConsumerDeleteResponseUnion() {}

type ConsumerNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r ConsumerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ConsumerNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ConsumerNewResponse   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ConsumerNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ConsumerNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       consumerNewResponseEnvelopeJSON       `json:"-"`
}

// consumerNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConsumerNewResponseEnvelope]
type consumerNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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

type ConsumerNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                   `json:"total_count"`
	JSON       consumerNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// consumerNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [ConsumerNewResponseEnvelopeResultInfo]
type consumerNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ConsumerUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r ConsumerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ConsumerUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   ConsumerUpdateResponse `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ConsumerUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ConsumerUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       consumerUpdateResponseEnvelopeJSON       `json:"-"`
}

// consumerUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConsumerUpdateResponseEnvelope]
type consumerUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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

type ConsumerUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                      `json:"total_count"`
	JSON       consumerUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// consumerUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [ConsumerUpdateResponseEnvelopeResultInfo]
type consumerUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ConsumerDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConsumerDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   ConsumerDeleteResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ConsumerDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ConsumerDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       consumerDeleteResponseEnvelopeJSON       `json:"-"`
}

// consumerDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConsumerDeleteResponseEnvelope]
type consumerDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConsumerDeleteResponseEnvelopeSuccess bool

const (
	ConsumerDeleteResponseEnvelopeSuccessTrue ConsumerDeleteResponseEnvelopeSuccess = true
)

func (r ConsumerDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConsumerDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConsumerDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                      `json:"total_count"`
	JSON       consumerDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// consumerDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [ConsumerDeleteResponseEnvelopeResultInfo]
type consumerDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ConsumerGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConsumerGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []Consumer            `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ConsumerGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ConsumerGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       consumerGetResponseEnvelopeJSON       `json:"-"`
}

// consumerGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConsumerGetResponseEnvelope]
type consumerGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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

type ConsumerGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                   `json:"total_count"`
	TotalPages float64                                   `json:"total_pages"`
	JSON       consumerGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// consumerGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [ConsumerGetResponseEnvelopeResultInfo]
type consumerGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
