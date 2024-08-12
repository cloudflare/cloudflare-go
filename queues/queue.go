// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// QueueService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueService] method instead.
type QueueService struct {
	Options   []option.RequestOption
	Consumers *ConsumerService
	Messages  *MessageService
}

// NewQueueService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQueueService(opts ...option.RequestOption) (r *QueueService) {
	r = &QueueService{}
	r.Options = opts
	r.Consumers = NewConsumerService(opts...)
	r.Messages = NewMessageService(opts...)
	return
}

// Creates a new queue.
func (r *QueueService) New(ctx context.Context, params QueueNewParams, opts ...option.RequestOption) (res *QueueCreated, err error) {
	var env QueueNewResponseEnvelope
	opts = append(r.Options[:], opts...)
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

// Updates a queue.
func (r *QueueService) Update(ctx context.Context, queueID string, params QueueUpdateParams, opts ...option.RequestOption) (res *QueueUpdated, err error) {
	var env QueueUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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

// Deletes a queue.
func (r *QueueService) Delete(ctx context.Context, queueID string, body QueueDeleteParams, opts ...option.RequestOption) (res *QueueDeleteResponseUnion, err error) {
	var env QueueDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s", body.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific queue.
func (r *QueueService) Get(ctx context.Context, queueID string, query QueueGetParams, opts ...option.RequestOption) (res *Queue, err error) {
	var env QueueGetResponseEnvelope
	opts = append(r.Options[:], opts...)
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
	Consumers           interface{} `json:"consumers"`
	ConsumersTotalCount float64     `json:"consumers_total_count"`
	CreatedOn           string      `json:"created_on"`
	ModifiedOn          string      `json:"modified_on"`
	Producers           interface{} `json:"producers"`
	ProducersTotalCount float64     `json:"producers_total_count"`
	QueueID             string      `json:"queue_id"`
	QueueName           string      `json:"queue_name"`
	JSON                queueJSON   `json:"-"`
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
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Queue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueJSON) RawJSON() string {
	return r.raw
}

type QueueCreated struct {
	CreatedOn  string           `json:"created_on"`
	ModifiedOn string           `json:"modified_on"`
	QueueID    string           `json:"queue_id"`
	QueueName  string           `json:"queue_name"`
	JSON       queueCreatedJSON `json:"-"`
}

// queueCreatedJSON contains the JSON metadata for the struct [QueueCreated]
type queueCreatedJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	QueueID     apijson.Field
	QueueName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueCreatedJSON) RawJSON() string {
	return r.raw
}

type QueueUpdated struct {
	CreatedOn  string           `json:"created_on"`
	ModifiedOn string           `json:"modified_on"`
	QueueID    string           `json:"queue_id"`
	QueueName  string           `json:"queue_name"`
	JSON       queueUpdatedJSON `json:"-"`
}

// queueUpdatedJSON contains the JSON metadata for the struct [QueueUpdated]
type queueUpdatedJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	QueueID     apijson.Field
	QueueName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueUpdatedJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [queues.QueueDeleteResponseUnknown],
// [queues.QueueDeleteResponseArray] or [shared.UnionString].
type QueueDeleteResponseUnion interface {
	ImplementsQueuesQueueDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*QueueDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(QueueDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type QueueDeleteResponseArray []interface{}

func (r QueueDeleteResponseArray) ImplementsQueuesQueueDeleteResponseUnion() {}

type QueueNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	QueueName param.Field[string] `json:"queue_name,required"`
}

func (r QueueNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QueueNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   QueueCreated          `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    QueueNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo QueueNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       queueNewResponseEnvelopeJSON       `json:"-"`
}

// queueNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [QueueNewResponseEnvelope]
type queueNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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

type QueueNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                `json:"total_count"`
	JSON       queueNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// queueNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [QueueNewResponseEnvelopeResultInfo]
type queueNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type QueueUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r QueueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type QueueUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   QueueUpdated          `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    QueueUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo QueueUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       queueUpdateResponseEnvelopeJSON       `json:"-"`
}

// queueUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [QueueUpdateResponseEnvelope]
type queueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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

type QueueUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                   `json:"total_count"`
	JSON       queueUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// queueUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [QueueUpdateResponseEnvelopeResultInfo]
type queueUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type QueueListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type QueueDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type QueueDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   QueueDeleteResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    QueueDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo QueueDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       queueDeleteResponseEnvelopeJSON       `json:"-"`
}

// queueDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [QueueDeleteResponseEnvelope]
type queueDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type QueueDeleteResponseEnvelopeSuccess bool

const (
	QueueDeleteResponseEnvelopeSuccessTrue QueueDeleteResponseEnvelopeSuccess = true
)

func (r QueueDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case QueueDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type QueueDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                   `json:"total_count"`
	JSON       queueDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// queueDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [QueueDeleteResponseEnvelopeResultInfo]
type queueDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type QueueGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type QueueGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Queue                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    QueueGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo QueueGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       queueGetResponseEnvelopeJSON       `json:"-"`
}

// queueGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [QueueGetResponseEnvelope]
type queueGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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

type QueueGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                `json:"total_count"`
	JSON       queueGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// queueGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [QueueGetResponseEnvelopeResultInfo]
type queueGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
