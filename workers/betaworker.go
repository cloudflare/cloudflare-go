// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// BetaWorkerService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaWorkerService] method instead.
type BetaWorkerService struct {
	Options  []option.RequestOption
	Versions *BetaWorkerVersionService
}

// NewBetaWorkerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaWorkerService(opts ...option.RequestOption) (r *BetaWorkerService) {
	r = &BetaWorkerService{}
	r.Options = opts
	r.Versions = NewBetaWorkerVersionService(opts...)
	return
}

// Create a new Worker.
func (r *BetaWorkerService) New(ctx context.Context, params BetaWorkerNewParams, opts ...option.RequestOption) (res *Worker, err error) {
	var env BetaWorkerNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/workers", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing Worker.
func (r *BetaWorkerService) Update(ctx context.Context, workerID string, params BetaWorkerUpdateParams, opts ...option.RequestOption) (res *Worker, err error) {
	var env BetaWorkerUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s", params.AccountID, workerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Workers for an account.
func (r *BetaWorkerService) List(ctx context.Context, params BetaWorkerListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Worker], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/workers", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List all Workers for an account.
func (r *BetaWorkerService) ListAutoPaging(ctx context.Context, params BetaWorkerListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Worker] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a Worker and all its associated resources (versions, deployments, etc.).
func (r *BetaWorkerService) Delete(ctx context.Context, workerID string, body BetaWorkerDeleteParams, opts ...option.RequestOption) (res *BetaWorkerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s", body.AccountID, workerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get details about a specific Worker.
func (r *BetaWorkerService) Get(ctx context.Context, workerID string, query BetaWorkerGetParams, opts ...option.RequestOption) (res *Worker, err error) {
	var env BetaWorkerGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s", query.AccountID, workerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Worker struct {
	// Identifier.
	ID string `json:"id,required"`
	// When the Worker was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Name of the Worker.
	Name string `json:"name,required"`
	// When the Worker was most recently updated.
	UpdatedOn time.Time `json:"updated_on,required" format:"date-time"`
	// Whether logpush is enabled for the Worker.
	Logpush bool `json:"logpush"`
	// Observability settings for the Worker.
	Observability WorkerObservability `json:"observability"`
	// Subdomain settings for the Worker.
	Subdomain WorkerSubdomain `json:"subdomain"`
	// Tags associated with the Worker.
	Tags []string `json:"tags"`
	// Other Workers that should consume logs from the Worker.
	TailConsumers []WorkerTailConsumer `json:"tail_consumers"`
	JSON          workerJSON           `json:"-"`
}

// workerJSON contains the JSON metadata for the struct [Worker]
type workerJSON struct {
	ID            apijson.Field
	CreatedOn     apijson.Field
	Name          apijson.Field
	UpdatedOn     apijson.Field
	Logpush       apijson.Field
	Observability apijson.Field
	Subdomain     apijson.Field
	Tags          apijson.Field
	TailConsumers apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Worker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerJSON) RawJSON() string {
	return r.raw
}

// Observability settings for the Worker.
type WorkerObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled bool `json:"enabled"`
	// The sampling rate for observability. From 0 to 1 (1 = 100%, 0.1 = 10%).
	HeadSamplingRate float64 `json:"head_sampling_rate"`
	// Log settings for the Worker.
	Logs WorkerObservabilityLogs `json:"logs"`
	JSON workerObservabilityJSON `json:"-"`
}

// workerObservabilityJSON contains the JSON metadata for the struct
// [WorkerObservability]
type workerObservabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	Logs             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WorkerObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerObservabilityJSON) RawJSON() string {
	return r.raw
}

// Log settings for the Worker.
type WorkerObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled bool `json:"enabled"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%).
	HeadSamplingRate float64 `json:"head_sampling_rate"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs bool                        `json:"invocation_logs"`
	JSON           workerObservabilityLogsJSON `json:"-"`
}

// workerObservabilityLogsJSON contains the JSON metadata for the struct
// [WorkerObservabilityLogs]
type workerObservabilityLogsJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	InvocationLogs   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WorkerObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// Subdomain settings for the Worker.
type WorkerSubdomain struct {
	// Whether the \*.workers.dev subdomain is enabled for the Worker.
	Enabled bool `json:"enabled"`
	// Whether
	// [preview URLs](https://developers.cloudflare.com/workers/configuration/previews/)
	// are enabled for the Worker.
	PreviewsEnabled bool                `json:"previews_enabled"`
	JSON            workerSubdomainJSON `json:"-"`
}

// workerSubdomainJSON contains the JSON metadata for the struct [WorkerSubdomain]
type workerSubdomainJSON struct {
	Enabled         apijson.Field
	PreviewsEnabled apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WorkerSubdomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerSubdomainJSON) RawJSON() string {
	return r.raw
}

type WorkerTailConsumer struct {
	// Name of the consumer Worker.
	Name string                 `json:"name,required"`
	JSON workerTailConsumerJSON `json:"-"`
}

// workerTailConsumerJSON contains the JSON metadata for the struct
// [WorkerTailConsumer]
type workerTailConsumerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerTailConsumerJSON) RawJSON() string {
	return r.raw
}

type WorkerParam struct {
	// Name of the Worker.
	Name param.Field[string] `json:"name,required"`
	// Whether logpush is enabled for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Observability settings for the Worker.
	Observability param.Field[WorkerObservabilityParam] `json:"observability"`
	// Subdomain settings for the Worker.
	Subdomain param.Field[WorkerSubdomainParam] `json:"subdomain"`
	// Tags associated with the Worker.
	Tags param.Field[[]string] `json:"tags"`
	// Other Workers that should consume logs from the Worker.
	TailConsumers param.Field[[]WorkerTailConsumerParam] `json:"tail_consumers"`
}

func (r WorkerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Observability settings for the Worker.
type WorkerObservabilityParam struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled"`
	// The sampling rate for observability. From 0 to 1 (1 = 100%, 0.1 = 10%).
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Log settings for the Worker.
	Logs param.Field[WorkerObservabilityLogsParam] `json:"logs"`
}

func (r WorkerObservabilityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Log settings for the Worker.
type WorkerObservabilityLogsParam struct {
	// Whether logs are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%).
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs param.Field[bool] `json:"invocation_logs"`
}

func (r WorkerObservabilityLogsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Subdomain settings for the Worker.
type WorkerSubdomainParam struct {
	// Whether the \*.workers.dev subdomain is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled"`
	// Whether
	// [preview URLs](https://developers.cloudflare.com/workers/configuration/previews/)
	// are enabled for the Worker.
	PreviewsEnabled param.Field[bool] `json:"previews_enabled"`
}

func (r WorkerSubdomainParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerTailConsumerParam struct {
	// Name of the consumer Worker.
	Name param.Field[string] `json:"name,required"`
}

func (r WorkerTailConsumerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaWorkerDeleteResponse struct {
	Errors   []BetaWorkerDeleteResponseError   `json:"errors,required"`
	Messages []BetaWorkerDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success BetaWorkerDeleteResponseSuccess `json:"success,required"`
	JSON    betaWorkerDeleteResponseJSON    `json:"-"`
}

// betaWorkerDeleteResponseJSON contains the JSON metadata for the struct
// [BetaWorkerDeleteResponse]
type betaWorkerDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerDeleteResponseError struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           BetaWorkerDeleteResponseErrorsSource `json:"source"`
	JSON             betaWorkerDeleteResponseErrorJSON    `json:"-"`
}

// betaWorkerDeleteResponseErrorJSON contains the JSON metadata for the struct
// [BetaWorkerDeleteResponseError]
type betaWorkerDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerDeleteResponseErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    betaWorkerDeleteResponseErrorsSourceJSON `json:"-"`
}

// betaWorkerDeleteResponseErrorsSourceJSON contains the JSON metadata for the
// struct [BetaWorkerDeleteResponseErrorsSource]
type betaWorkerDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerDeleteResponseMessage struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           BetaWorkerDeleteResponseMessagesSource `json:"source"`
	JSON             betaWorkerDeleteResponseMessageJSON    `json:"-"`
}

// betaWorkerDeleteResponseMessageJSON contains the JSON metadata for the struct
// [BetaWorkerDeleteResponseMessage]
type betaWorkerDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerDeleteResponseMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    betaWorkerDeleteResponseMessagesSourceJSON `json:"-"`
}

// betaWorkerDeleteResponseMessagesSourceJSON contains the JSON metadata for the
// struct [BetaWorkerDeleteResponseMessagesSource]
type betaWorkerDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BetaWorkerDeleteResponseSuccess bool

const (
	BetaWorkerDeleteResponseSuccessTrue BetaWorkerDeleteResponseSuccess = true
)

func (r BetaWorkerDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case BetaWorkerDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type BetaWorkerNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Worker    WorkerParam         `json:"worker,required"`
}

func (r BetaWorkerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Worker)
}

type BetaWorkerNewResponseEnvelope struct {
	Errors   []BetaWorkerNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BetaWorkerNewResponseEnvelopeMessages `json:"messages,required"`
	Result   Worker                                  `json:"result,required"`
	// Whether the API call was successful.
	Success BetaWorkerNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    betaWorkerNewResponseEnvelopeJSON    `json:"-"`
}

// betaWorkerNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [BetaWorkerNewResponseEnvelope]
type betaWorkerNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerNewResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           BetaWorkerNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             betaWorkerNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// betaWorkerNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BetaWorkerNewResponseEnvelopeErrors]
type betaWorkerNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerNewResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    betaWorkerNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// betaWorkerNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [BetaWorkerNewResponseEnvelopeErrorsSource]
type betaWorkerNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerNewResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           BetaWorkerNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             betaWorkerNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// betaWorkerNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [BetaWorkerNewResponseEnvelopeMessages]
type betaWorkerNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerNewResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    betaWorkerNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// betaWorkerNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [BetaWorkerNewResponseEnvelopeMessagesSource]
type betaWorkerNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BetaWorkerNewResponseEnvelopeSuccess bool

const (
	BetaWorkerNewResponseEnvelopeSuccessTrue BetaWorkerNewResponseEnvelopeSuccess = true
)

func (r BetaWorkerNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BetaWorkerNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BetaWorkerUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Worker    WorkerParam         `json:"worker,required"`
}

func (r BetaWorkerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Worker)
}

type BetaWorkerUpdateResponseEnvelope struct {
	Errors   []BetaWorkerUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BetaWorkerUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   Worker                                     `json:"result,required"`
	// Whether the API call was successful.
	Success BetaWorkerUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    betaWorkerUpdateResponseEnvelopeJSON    `json:"-"`
}

// betaWorkerUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [BetaWorkerUpdateResponseEnvelope]
type betaWorkerUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerUpdateResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           BetaWorkerUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             betaWorkerUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// betaWorkerUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BetaWorkerUpdateResponseEnvelopeErrors]
type betaWorkerUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    betaWorkerUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// betaWorkerUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [BetaWorkerUpdateResponseEnvelopeErrorsSource]
type betaWorkerUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerUpdateResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           BetaWorkerUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             betaWorkerUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// betaWorkerUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [BetaWorkerUpdateResponseEnvelopeMessages]
type betaWorkerUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    betaWorkerUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// betaWorkerUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [BetaWorkerUpdateResponseEnvelopeMessagesSource]
type betaWorkerUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BetaWorkerUpdateResponseEnvelopeSuccess bool

const (
	BetaWorkerUpdateResponseEnvelopeSuccessTrue BetaWorkerUpdateResponseEnvelopeSuccess = true
)

func (r BetaWorkerUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BetaWorkerUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BetaWorkerListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Current page.
	Page param.Field[int64] `query:"page"`
	// Items per-page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [BetaWorkerListParams]'s query parameters as `url.Values`.
func (r BetaWorkerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BetaWorkerDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type BetaWorkerGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type BetaWorkerGetResponseEnvelope struct {
	Errors   []BetaWorkerGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BetaWorkerGetResponseEnvelopeMessages `json:"messages,required"`
	Result   Worker                                  `json:"result,required"`
	// Whether the API call was successful.
	Success BetaWorkerGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    betaWorkerGetResponseEnvelopeJSON    `json:"-"`
}

// betaWorkerGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BetaWorkerGetResponseEnvelope]
type betaWorkerGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerGetResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           BetaWorkerGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             betaWorkerGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// betaWorkerGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BetaWorkerGetResponseEnvelopeErrors]
type betaWorkerGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    betaWorkerGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// betaWorkerGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [BetaWorkerGetResponseEnvelopeErrorsSource]
type betaWorkerGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           BetaWorkerGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             betaWorkerGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// betaWorkerGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [BetaWorkerGetResponseEnvelopeMessages]
type betaWorkerGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    betaWorkerGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// betaWorkerGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [BetaWorkerGetResponseEnvelopeMessagesSource]
type betaWorkerGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BetaWorkerGetResponseEnvelopeSuccess bool

const (
	BetaWorkerGetResponseEnvelopeSuccessTrue BetaWorkerGetResponseEnvelopeSuccess = true
)

func (r BetaWorkerGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BetaWorkerGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
