// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// QueueService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewQueueService] method instead.
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
func (r *QueueService) New(ctx context.Context, params QueueNewParams, opts ...option.RequestOption) (res *QueueNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env QueueNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/queues", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a queue.
func (r *QueueService) Update(ctx context.Context, queueID string, params QueueUpdateParams, opts ...option.RequestOption) (res *QueueUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env QueueUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/queues/%s", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the queues owned by an account.
func (r *QueueService) List(ctx context.Context, query QueueListParams, opts ...option.RequestOption) (res *[]QueueListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env QueueListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/queues", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a queue.
func (r *QueueService) Delete(ctx context.Context, queueID string, body QueueDeleteParams, opts ...option.RequestOption) (res *QueueDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env QueueDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/queues/%s", body.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific queue.
func (r *QueueService) Get(ctx context.Context, queueID string, query QueueGetParams, opts ...option.RequestOption) (res *QueueGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env QueueGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/queues/%s", query.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type QueueNewResponse struct {
	CreatedOn  interface{}          `json:"created_on"`
	ModifiedOn interface{}          `json:"modified_on"`
	QueueID    interface{}          `json:"queue_id"`
	QueueName  string               `json:"queue_name"`
	JSON       queueNewResponseJSON `json:"-"`
}

// queueNewResponseJSON contains the JSON metadata for the struct
// [QueueNewResponse]
type queueNewResponseJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	QueueID     apijson.Field
	QueueName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueNewResponseJSON) RawJSON() string {
	return r.raw
}

type QueueUpdateResponse struct {
	CreatedOn  interface{}             `json:"created_on"`
	ModifiedOn interface{}             `json:"modified_on"`
	QueueID    interface{}             `json:"queue_id"`
	QueueName  string                  `json:"queue_name"`
	JSON       queueUpdateResponseJSON `json:"-"`
}

// queueUpdateResponseJSON contains the JSON metadata for the struct
// [QueueUpdateResponse]
type queueUpdateResponseJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	QueueID     apijson.Field
	QueueName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type QueueListResponse struct {
	Consumers           interface{}           `json:"consumers"`
	ConsumersTotalCount interface{}           `json:"consumers_total_count"`
	CreatedOn           interface{}           `json:"created_on"`
	ModifiedOn          interface{}           `json:"modified_on"`
	Producers           interface{}           `json:"producers"`
	ProducersTotalCount interface{}           `json:"producers_total_count"`
	QueueID             string                `json:"queue_id"`
	QueueName           string                `json:"queue_name"`
	JSON                queueListResponseJSON `json:"-"`
}

// queueListResponseJSON contains the JSON metadata for the struct
// [QueueListResponse]
type queueListResponseJSON struct {
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

func (r *QueueListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueListResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [queues.QueueDeleteResponseUnknown],
// [queues.QueueDeleteResponseArray] or [shared.UnionString].
type QueueDeleteResponse interface {
	ImplementsQueuesQueueDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*QueueDeleteResponse)(nil)).Elem(),
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

func (r QueueDeleteResponseArray) ImplementsQueuesQueueDeleteResponse() {}

type QueueGetResponse struct {
	Consumers           interface{}          `json:"consumers"`
	ConsumersTotalCount interface{}          `json:"consumers_total_count"`
	CreatedOn           interface{}          `json:"created_on"`
	ModifiedOn          interface{}          `json:"modified_on"`
	Producers           interface{}          `json:"producers"`
	ProducersTotalCount interface{}          `json:"producers_total_count"`
	QueueID             string               `json:"queue_id"`
	QueueName           string               `json:"queue_name"`
	JSON                queueGetResponseJSON `json:"-"`
}

// queueGetResponseJSON contains the JSON metadata for the struct
// [QueueGetResponse]
type queueGetResponseJSON struct {
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

func (r *QueueGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueGetResponseJSON) RawJSON() string {
	return r.raw
}

type QueueNewParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r QueueNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type QueueNewResponseEnvelope struct {
	Errors   []QueueNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []QueueNewResponseEnvelopeMessages `json:"messages,required"`
	Result   QueueNewResponse                   `json:"result,required,nullable"`
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

type QueueNewResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    queueNewResponseEnvelopeErrorsJSON `json:"-"`
}

// queueNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [QueueNewResponseEnvelopeErrors]
type queueNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type QueueNewResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    queueNewResponseEnvelopeMessagesJSON `json:"-"`
}

// queueNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [QueueNewResponseEnvelopeMessages]
type queueNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueNewResponseEnvelopeMessagesJSON) RawJSON() string {
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
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
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
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r QueueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type QueueUpdateResponseEnvelope struct {
	Errors   []QueueUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []QueueUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   QueueUpdateResponse                   `json:"result,required,nullable"`
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

type QueueUpdateResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    queueUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// queueUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [QueueUpdateResponseEnvelopeErrors]
type queueUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type QueueUpdateResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    queueUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// queueUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [QueueUpdateResponseEnvelopeMessages]
type queueUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
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
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
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

type QueueListResponseEnvelope struct {
	Errors   []QueueListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []QueueListResponseEnvelopeMessages `json:"messages,required"`
	Result   []QueueListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    QueueListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo QueueListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       queueListResponseEnvelopeJSON       `json:"-"`
}

// queueListResponseEnvelopeJSON contains the JSON metadata for the struct
// [QueueListResponseEnvelope]
type queueListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type QueueListResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    queueListResponseEnvelopeErrorsJSON `json:"-"`
}

// queueListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [QueueListResponseEnvelopeErrors]
type queueListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type QueueListResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    queueListResponseEnvelopeMessagesJSON `json:"-"`
}

// queueListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [QueueListResponseEnvelopeMessages]
type queueListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type QueueListResponseEnvelopeSuccess bool

const (
	QueueListResponseEnvelopeSuccessTrue QueueListResponseEnvelopeSuccess = true
)

func (r QueueListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case QueueListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type QueueListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	TotalPages float64                                 `json:"total_pages"`
	JSON       queueListResponseEnvelopeResultInfoJSON `json:"-"`
}

// queueListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [QueueListResponseEnvelopeResultInfo]
type queueListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type QueueDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type QueueDeleteResponseEnvelope struct {
	Errors   []QueueDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []QueueDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   QueueDeleteResponse                   `json:"result,required,nullable"`
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

type QueueDeleteResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    queueDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// queueDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [QueueDeleteResponseEnvelopeErrors]
type queueDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type QueueDeleteResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    queueDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// queueDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [QueueDeleteResponseEnvelopeMessages]
type queueDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
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
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
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
	Errors   []QueueGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []QueueGetResponseEnvelopeMessages `json:"messages,required"`
	Result   QueueGetResponse                   `json:"result,required,nullable"`
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

type QueueGetResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    queueGetResponseEnvelopeErrorsJSON `json:"-"`
}

// queueGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [QueueGetResponseEnvelopeErrors]
type queueGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type QueueGetResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    queueGetResponseEnvelopeMessagesJSON `json:"-"`
}

// queueGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [QueueGetResponseEnvelopeMessages]
type queueGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueGetResponseEnvelopeMessagesJSON) RawJSON() string {
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
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
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
