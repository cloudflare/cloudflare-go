// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
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
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/workers", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Perform a complete replacement of a Worker, where omitted properties are set to
// their default values. This is the exact same as the Create Worker endpoint, but
// operates on an existing Worker. To perform a partial update instead, use the
// Edit Worker endpoint.
func (r *BetaWorkerService) Update(ctx context.Context, workerID string, params BetaWorkerUpdateParams, opts ...option.RequestOption) (res *Worker, err error) {
	var env BetaWorkerUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s", params.AccountID, workerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all Workers for an account.
func (r *BetaWorkerService) List(ctx context.Context, params BetaWorkerListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Worker], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
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
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s", body.AccountID, workerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Perform a partial update on a Worker, where omitted properties are left
// unchanged from their current values.
func (r *BetaWorkerService) Edit(ctx context.Context, workerID string, params BetaWorkerEditParams, opts ...option.RequestOption) (res *Worker, err error) {
	var env BetaWorkerEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s", params.AccountID, workerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get details about a specific Worker.
func (r *BetaWorkerService) Get(ctx context.Context, workerID string, query BetaWorkerGetParams, opts ...option.RequestOption) (res *Worker, err error) {
	var env BetaWorkerGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s", query.AccountID, workerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type Worker struct {
	// Immutable ID of the Worker.
	ID string `json:"id" api:"required"`
	// When the Worker was created.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Whether logpush is enabled for the Worker.
	Logpush bool `json:"logpush" api:"required"`
	// Name of the Worker.
	Name string `json:"name" api:"required"`
	// Observability settings for the Worker.
	Observability WorkerObservability `json:"observability" api:"required"`
	// Other resources that reference the Worker and depend on it existing.
	References WorkerReferences `json:"references" api:"required"`
	// Subdomain settings for the Worker.
	Subdomain WorkerSubdomain `json:"subdomain" api:"required"`
	// Tags associated with the Worker.
	Tags []string `json:"tags" api:"required"`
	// Other Workers that should consume logs from the Worker.
	TailConsumers []WorkerTailConsumer `json:"tail_consumers" api:"required"`
	// When the Worker was most recently updated.
	UpdatedOn time.Time `json:"updated_on" api:"required" format:"date-time"`
	// When the Worker's most recent deployment was created. `null` if the Worker has
	// never been deployed.
	DeployedOn time.Time  `json:"deployed_on" api:"nullable" format:"date-time"`
	JSON       workerJSON `json:"-"`
}

// workerJSON contains the JSON metadata for the struct [Worker]
type workerJSON struct {
	ID            apijson.Field
	CreatedOn     apijson.Field
	Logpush       apijson.Field
	Name          apijson.Field
	Observability apijson.Field
	References    apijson.Field
	Subdomain     apijson.Field
	Tags          apijson.Field
	TailConsumers apijson.Field
	UpdatedOn     apijson.Field
	DeployedOn    apijson.Field
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
	// Trace settings for the Worker.
	Traces WorkerObservabilityTraces `json:"traces"`
	JSON   workerObservabilityJSON   `json:"-"`
}

// workerObservabilityJSON contains the JSON metadata for the struct
// [WorkerObservability]
type workerObservabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	Logs             apijson.Field
	Traces           apijson.Field
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
	// A list of destinations where logs will be exported to.
	Destinations []string `json:"destinations"`
	// Whether logs are enabled for the Worker.
	Enabled bool `json:"enabled"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%).
	HeadSamplingRate float64 `json:"head_sampling_rate"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs bool `json:"invocation_logs"`
	// Whether log persistence is enabled for the Worker.
	Persist bool                        `json:"persist"`
	JSON    workerObservabilityLogsJSON `json:"-"`
}

// workerObservabilityLogsJSON contains the JSON metadata for the struct
// [WorkerObservabilityLogs]
type workerObservabilityLogsJSON struct {
	Destinations     apijson.Field
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	InvocationLogs   apijson.Field
	Persist          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WorkerObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// Trace settings for the Worker.
type WorkerObservabilityTraces struct {
	// A list of destinations where traces will be exported to.
	Destinations []string `json:"destinations"`
	// Whether traces are enabled for the Worker.
	Enabled bool `json:"enabled"`
	// The sampling rate for traces. From 0 to 1 (1 = 100%, 0.1 = 10%).
	HeadSamplingRate float64 `json:"head_sampling_rate"`
	// Whether trace persistence is enabled for the Worker.
	Persist bool                          `json:"persist"`
	JSON    workerObservabilityTracesJSON `json:"-"`
}

// workerObservabilityTracesJSON contains the JSON metadata for the struct
// [WorkerObservabilityTraces]
type workerObservabilityTracesJSON struct {
	Destinations     apijson.Field
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	Persist          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WorkerObservabilityTraces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerObservabilityTracesJSON) RawJSON() string {
	return r.raw
}

// Other resources that reference the Worker and depend on it existing.
type WorkerReferences struct {
	// Other Workers that reference the Worker as an outbound for a dispatch namespace.
	DispatchNamespaceOutbounds []WorkerReferencesDispatchNamespaceOutbound `json:"dispatch_namespace_outbounds" api:"required"`
	// Custom domains connected to the Worker.
	Domains []WorkerReferencesDomain `json:"domains" api:"required"`
	// Other Workers that reference Durable Object classes implemented by the Worker.
	DurableObjects []WorkerReferencesDurableObject `json:"durable_objects" api:"required"`
	// Queues that send messages to the Worker.
	Queues []WorkerReferencesQueue `json:"queues" api:"required"`
	// Other Workers that reference the Worker using
	// [service bindings](https://developers.cloudflare.com/workers/runtime-apis/bindings/service-bindings/).
	Workers []WorkerReferencesWorker `json:"workers" api:"required"`
	JSON    workerReferencesJSON     `json:"-"`
}

// workerReferencesJSON contains the JSON metadata for the struct
// [WorkerReferences]
type workerReferencesJSON struct {
	DispatchNamespaceOutbounds apijson.Field
	Domains                    apijson.Field
	DurableObjects             apijson.Field
	Queues                     apijson.Field
	Workers                    apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *WorkerReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerReferencesJSON) RawJSON() string {
	return r.raw
}

type WorkerReferencesDispatchNamespaceOutbound struct {
	// ID of the dispatch namespace.
	NamespaceID string `json:"namespace_id" api:"required"`
	// Name of the dispatch namespace.
	NamespaceName string `json:"namespace_name" api:"required"`
	// ID of the Worker using the dispatch namespace.
	WorkerID string `json:"worker_id" api:"required"`
	// Name of the Worker using the dispatch namespace.
	WorkerName string                                        `json:"worker_name" api:"required"`
	JSON       workerReferencesDispatchNamespaceOutboundJSON `json:"-"`
}

// workerReferencesDispatchNamespaceOutboundJSON contains the JSON metadata for the
// struct [WorkerReferencesDispatchNamespaceOutbound]
type workerReferencesDispatchNamespaceOutboundJSON struct {
	NamespaceID   apijson.Field
	NamespaceName apijson.Field
	WorkerID      apijson.Field
	WorkerName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerReferencesDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerReferencesDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

type WorkerReferencesDomain struct {
	// ID of the custom domain.
	ID string `json:"id" api:"required"`
	// ID of the TLS certificate issued for the custom domain.
	CertificateID string `json:"certificate_id" api:"required"`
	// Full hostname of the custom domain, including the zone name.
	Hostname string `json:"hostname" api:"required"`
	// ID of the zone.
	ZoneID string `json:"zone_id" api:"required"`
	// Name of the zone.
	ZoneName string                     `json:"zone_name" api:"required"`
	JSON     workerReferencesDomainJSON `json:"-"`
}

// workerReferencesDomainJSON contains the JSON metadata for the struct
// [WorkerReferencesDomain]
type workerReferencesDomainJSON struct {
	ID            apijson.Field
	CertificateID apijson.Field
	Hostname      apijson.Field
	ZoneID        apijson.Field
	ZoneName      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerReferencesDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerReferencesDomainJSON) RawJSON() string {
	return r.raw
}

type WorkerReferencesDurableObject struct {
	// ID of the Durable Object namespace being used.
	NamespaceID string `json:"namespace_id" api:"required"`
	// Name of the Durable Object namespace being used.
	NamespaceName string `json:"namespace_name" api:"required"`
	// ID of the Worker using the Durable Object implementation.
	WorkerID string `json:"worker_id" api:"required"`
	// Name of the Worker using the Durable Object implementation.
	WorkerName string                            `json:"worker_name" api:"required"`
	JSON       workerReferencesDurableObjectJSON `json:"-"`
}

// workerReferencesDurableObjectJSON contains the JSON metadata for the struct
// [WorkerReferencesDurableObject]
type workerReferencesDurableObjectJSON struct {
	NamespaceID   apijson.Field
	NamespaceName apijson.Field
	WorkerID      apijson.Field
	WorkerName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerReferencesDurableObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerReferencesDurableObjectJSON) RawJSON() string {
	return r.raw
}

type WorkerReferencesQueue struct {
	// ID of the queue consumer configuration.
	QueueConsumerID string `json:"queue_consumer_id" api:"required"`
	// ID of the queue.
	QueueID string `json:"queue_id" api:"required"`
	// Name of the queue.
	QueueName string                    `json:"queue_name" api:"required"`
	JSON      workerReferencesQueueJSON `json:"-"`
}

// workerReferencesQueueJSON contains the JSON metadata for the struct
// [WorkerReferencesQueue]
type workerReferencesQueueJSON struct {
	QueueConsumerID apijson.Field
	QueueID         apijson.Field
	QueueName       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WorkerReferencesQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerReferencesQueueJSON) RawJSON() string {
	return r.raw
}

type WorkerReferencesWorker struct {
	// ID of the referencing Worker.
	ID string `json:"id" api:"required"`
	// Name of the referencing Worker.
	Name string                     `json:"name" api:"required"`
	JSON workerReferencesWorkerJSON `json:"-"`
}

// workerReferencesWorkerJSON contains the JSON metadata for the struct
// [WorkerReferencesWorker]
type workerReferencesWorkerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerReferencesWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerReferencesWorkerJSON) RawJSON() string {
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
	Name string                 `json:"name" api:"required"`
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
	// Whether logpush is enabled for the Worker.
	Logpush param.Field[bool] `json:"logpush" api:"required"`
	// Name of the Worker.
	Name param.Field[string] `json:"name" api:"required"`
	// Observability settings for the Worker.
	Observability param.Field[WorkerObservabilityParam] `json:"observability" api:"required"`
	// Subdomain settings for the Worker.
	Subdomain param.Field[WorkerSubdomainParam] `json:"subdomain" api:"required"`
	// Tags associated with the Worker.
	Tags param.Field[[]string] `json:"tags" api:"required"`
	// Other Workers that should consume logs from the Worker.
	TailConsumers param.Field[[]WorkerTailConsumerParam] `json:"tail_consumers" api:"required"`
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
	// Trace settings for the Worker.
	Traces param.Field[WorkerObservabilityTracesParam] `json:"traces"`
}

func (r WorkerObservabilityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Log settings for the Worker.
type WorkerObservabilityLogsParam struct {
	// A list of destinations where logs will be exported to.
	Destinations param.Field[[]string] `json:"destinations"`
	// Whether logs are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%).
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs param.Field[bool] `json:"invocation_logs"`
	// Whether log persistence is enabled for the Worker.
	Persist param.Field[bool] `json:"persist"`
}

func (r WorkerObservabilityLogsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Trace settings for the Worker.
type WorkerObservabilityTracesParam struct {
	// A list of destinations where traces will be exported to.
	Destinations param.Field[[]string] `json:"destinations"`
	// Whether traces are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled"`
	// The sampling rate for traces. From 0 to 1 (1 = 100%, 0.1 = 10%).
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Whether trace persistence is enabled for the Worker.
	Persist param.Field[bool] `json:"persist"`
}

func (r WorkerObservabilityTracesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Other resources that reference the Worker and depend on it existing.
type WorkerReferencesParam struct {
	// Other Workers that reference the Worker as an outbound for a dispatch namespace.
	DispatchNamespaceOutbounds param.Field[[]WorkerReferencesDispatchNamespaceOutboundParam] `json:"dispatch_namespace_outbounds" api:"required"`
	// Custom domains connected to the Worker.
	Domains param.Field[[]WorkerReferencesDomainParam] `json:"domains" api:"required"`
	// Other Workers that reference Durable Object classes implemented by the Worker.
	DurableObjects param.Field[[]WorkerReferencesDurableObjectParam] `json:"durable_objects" api:"required"`
	// Queues that send messages to the Worker.
	Queues param.Field[[]WorkerReferencesQueueParam] `json:"queues" api:"required"`
	// Other Workers that reference the Worker using
	// [service bindings](https://developers.cloudflare.com/workers/runtime-apis/bindings/service-bindings/).
	Workers param.Field[[]WorkerReferencesWorkerParam] `json:"workers" api:"required"`
}

func (r WorkerReferencesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerReferencesDispatchNamespaceOutboundParam struct {
	// ID of the dispatch namespace.
	NamespaceID param.Field[string] `json:"namespace_id" api:"required"`
	// Name of the dispatch namespace.
	NamespaceName param.Field[string] `json:"namespace_name" api:"required"`
	// ID of the Worker using the dispatch namespace.
	WorkerID param.Field[string] `json:"worker_id" api:"required"`
	// Name of the Worker using the dispatch namespace.
	WorkerName param.Field[string] `json:"worker_name" api:"required"`
}

func (r WorkerReferencesDispatchNamespaceOutboundParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerReferencesDomainParam struct {
	// ID of the custom domain.
	ID param.Field[string] `json:"id" api:"required"`
	// ID of the TLS certificate issued for the custom domain.
	CertificateID param.Field[string] `json:"certificate_id" api:"required"`
	// Full hostname of the custom domain, including the zone name.
	Hostname param.Field[string] `json:"hostname" api:"required"`
	// ID of the zone.
	ZoneID param.Field[string] `json:"zone_id" api:"required"`
	// Name of the zone.
	ZoneName param.Field[string] `json:"zone_name" api:"required"`
}

func (r WorkerReferencesDomainParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerReferencesDurableObjectParam struct {
	// ID of the Durable Object namespace being used.
	NamespaceID param.Field[string] `json:"namespace_id" api:"required"`
	// Name of the Durable Object namespace being used.
	NamespaceName param.Field[string] `json:"namespace_name" api:"required"`
	// ID of the Worker using the Durable Object implementation.
	WorkerID param.Field[string] `json:"worker_id" api:"required"`
	// Name of the Worker using the Durable Object implementation.
	WorkerName param.Field[string] `json:"worker_name" api:"required"`
}

func (r WorkerReferencesDurableObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerReferencesQueueParam struct {
	// ID of the queue consumer configuration.
	QueueConsumerID param.Field[string] `json:"queue_consumer_id" api:"required"`
	// ID of the queue.
	QueueID param.Field[string] `json:"queue_id" api:"required"`
	// Name of the queue.
	QueueName param.Field[string] `json:"queue_name" api:"required"`
}

func (r WorkerReferencesQueueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerReferencesWorkerParam struct {
	// ID of the referencing Worker.
	ID param.Field[string] `json:"id" api:"required"`
	// Name of the referencing Worker.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r WorkerReferencesWorkerParam) MarshalJSON() (data []byte, err error) {
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
	Name param.Field[string] `json:"name" api:"required"`
}

func (r WorkerTailConsumerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaWorkerDeleteResponse struct {
	Errors   []BetaWorkerDeleteResponseError   `json:"errors" api:"required"`
	Messages []BetaWorkerDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success BetaWorkerDeleteResponseSuccess `json:"success" api:"required"`
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
	Code             int64                                `json:"code" api:"required"`
	Message          string                               `json:"message" api:"required"`
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
	Code             int64                                  `json:"code" api:"required"`
	Message          string                                 `json:"message" api:"required"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Worker    WorkerParam         `json:"worker" api:"required"`
}

func (r BetaWorkerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Worker)
}

type BetaWorkerNewResponseEnvelope struct {
	Errors   []BetaWorkerNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []BetaWorkerNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   Worker                                  `json:"result" api:"required"`
	// Whether the API call was successful.
	Success BetaWorkerNewResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
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
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Worker    WorkerParam         `json:"worker" api:"required"`
}

func (r BetaWorkerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Worker)
}

type BetaWorkerUpdateResponseEnvelope struct {
	Errors   []BetaWorkerUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []BetaWorkerUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   Worker                                     `json:"result" api:"required"`
	// Whether the API call was successful.
	Success BetaWorkerUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
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
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Sort direction.
	Order param.Field[BetaWorkerListParamsOrder] `query:"order"`
	// Property to sort results by.
	OrderBy param.Field[BetaWorkerListParamsOrderBy] `query:"order_by"`
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

// Sort direction.
type BetaWorkerListParamsOrder string

const (
	BetaWorkerListParamsOrderAsc  BetaWorkerListParamsOrder = "asc"
	BetaWorkerListParamsOrderDesc BetaWorkerListParamsOrder = "desc"
)

func (r BetaWorkerListParamsOrder) IsKnown() bool {
	switch r {
	case BetaWorkerListParamsOrderAsc, BetaWorkerListParamsOrderDesc:
		return true
	}
	return false
}

// Property to sort results by.
type BetaWorkerListParamsOrderBy string

const (
	BetaWorkerListParamsOrderByDeployedOn BetaWorkerListParamsOrderBy = "deployed_on"
	BetaWorkerListParamsOrderByUpdatedOn  BetaWorkerListParamsOrderBy = "updated_on"
	BetaWorkerListParamsOrderByCreatedOn  BetaWorkerListParamsOrderBy = "created_on"
	BetaWorkerListParamsOrderByName       BetaWorkerListParamsOrderBy = "name"
)

func (r BetaWorkerListParamsOrderBy) IsKnown() bool {
	switch r {
	case BetaWorkerListParamsOrderByDeployedOn, BetaWorkerListParamsOrderByUpdatedOn, BetaWorkerListParamsOrderByCreatedOn, BetaWorkerListParamsOrderByName:
		return true
	}
	return false
}

type BetaWorkerDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BetaWorkerEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Worker    WorkerParam         `json:"worker" api:"required"`
}

func (r BetaWorkerEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Worker)
}

type BetaWorkerEditResponseEnvelope struct {
	Errors   []BetaWorkerEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []BetaWorkerEditResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   Worker                                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success BetaWorkerEditResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    betaWorkerEditResponseEnvelopeJSON    `json:"-"`
}

// betaWorkerEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [BetaWorkerEditResponseEnvelope]
type betaWorkerEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerEditResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           BetaWorkerEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             betaWorkerEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// betaWorkerEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BetaWorkerEditResponseEnvelopeErrors]
type betaWorkerEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerEditResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    betaWorkerEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// betaWorkerEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [BetaWorkerEditResponseEnvelopeErrorsSource]
type betaWorkerEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerEditResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           BetaWorkerEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             betaWorkerEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// betaWorkerEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [BetaWorkerEditResponseEnvelopeMessages]
type betaWorkerEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerEditResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    betaWorkerEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// betaWorkerEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [BetaWorkerEditResponseEnvelopeMessagesSource]
type betaWorkerEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BetaWorkerEditResponseEnvelopeSuccess bool

const (
	BetaWorkerEditResponseEnvelopeSuccessTrue BetaWorkerEditResponseEnvelopeSuccess = true
)

func (r BetaWorkerEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BetaWorkerEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BetaWorkerGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BetaWorkerGetResponseEnvelope struct {
	Errors   []BetaWorkerGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []BetaWorkerGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   Worker                                  `json:"result" api:"required"`
	// Whether the API call was successful.
	Success BetaWorkerGetResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
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
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
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
