// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
	"github.com/tidwall/gjson"
)

// SubscriptionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r *SubscriptionService) {
	r = &SubscriptionService{}
	r.Options = opts
	return
}

// Create a new event subscription for a queue
func (r *SubscriptionService) New(ctx context.Context, params SubscriptionNewParams, opts ...option.RequestOption) (res *SubscriptionNewResponse, err error) {
	var env SubscriptionNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_subscriptions/subscriptions", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing event subscription
func (r *SubscriptionService) Update(ctx context.Context, subscriptionID string, params SubscriptionUpdateParams, opts ...option.RequestOption) (res *SubscriptionUpdateResponse, err error) {
	var env SubscriptionUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_subscriptions/subscriptions/%s", params.AccountID, subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a paginated list of event subscriptions with optional sorting and filtering
func (r *SubscriptionService) List(ctx context.Context, params SubscriptionListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SubscriptionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_subscriptions/subscriptions", params.AccountID)
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

// Get a paginated list of event subscriptions with optional sorting and filtering
func (r *SubscriptionService) ListAutoPaging(ctx context.Context, params SubscriptionListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SubscriptionListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete an existing event subscription
func (r *SubscriptionService) Delete(ctx context.Context, subscriptionID string, body SubscriptionDeleteParams, opts ...option.RequestOption) (res *SubscriptionDeleteResponse, err error) {
	var env SubscriptionDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_subscriptions/subscriptions/%s", body.AccountID, subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details about an existing event subscription
func (r *SubscriptionService) Get(ctx context.Context, subscriptionID string, query SubscriptionGetParams, opts ...option.RequestOption) (res *SubscriptionGetResponse, err error) {
	var env SubscriptionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_subscriptions/subscriptions/%s", query.AccountID, subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SubscriptionNewResponse struct {
	// Unique identifier for the subscription
	ID string `json:"id,required"`
	// When the subscription was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Destination configuration for the subscription
	Destination SubscriptionNewResponseDestination `json:"destination,required"`
	// Whether the subscription is active
	Enabled bool `json:"enabled,required"`
	// List of event types this subscription handles
	Events []string `json:"events,required"`
	// When the subscription was last modified
	ModifiedAt time.Time `json:"modified_at,required" format:"date-time"`
	// Name of the subscription
	Name string `json:"name,required"`
	// Source configuration for the subscription
	Source SubscriptionNewResponseSource `json:"source,required"`
	JSON   subscriptionNewResponseJSON   `json:"-"`
}

// subscriptionNewResponseJSON contains the JSON metadata for the struct
// [SubscriptionNewResponse]
type subscriptionNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Destination apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseJSON) RawJSON() string {
	return r.raw
}

// Destination configuration for the subscription
type SubscriptionNewResponseDestination struct {
	// ID of the target queue
	QueueID string `json:"queue_id,required"`
	// Type of destination
	Type SubscriptionNewResponseDestinationType `json:"type,required"`
	JSON subscriptionNewResponseDestinationJSON `json:"-"`
}

// subscriptionNewResponseDestinationJSON contains the JSON metadata for the struct
// [SubscriptionNewResponseDestination]
type subscriptionNewResponseDestinationJSON struct {
	QueueID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseDestinationJSON) RawJSON() string {
	return r.raw
}

// Type of destination
type SubscriptionNewResponseDestinationType string

const (
	SubscriptionNewResponseDestinationTypeQueuesQueue SubscriptionNewResponseDestinationType = "queues.queue"
)

func (r SubscriptionNewResponseDestinationType) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseDestinationTypeQueuesQueue:
		return true
	}
	return false
}

// Source configuration for the subscription
type SubscriptionNewResponseSource struct {
	// Name of the Workers AI model
	ModelName string `json:"model_name"`
	// Type of source
	Type SubscriptionNewResponseSourceType `json:"type"`
	// Name of the worker
	WorkerName string `json:"worker_name"`
	// Name of the workflow
	WorkflowName string                            `json:"workflow_name"`
	JSON         subscriptionNewResponseSourceJSON `json:"-"`
	union        SubscriptionNewResponseSourceUnion
}

// subscriptionNewResponseSourceJSON contains the JSON metadata for the struct
// [SubscriptionNewResponseSource]
type subscriptionNewResponseSourceJSON struct {
	ModelName    apijson.Field
	Type         apijson.Field
	WorkerName   apijson.Field
	WorkflowName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r subscriptionNewResponseSourceJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionNewResponseSource) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionNewResponseSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SubscriptionNewResponseSourceUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionNewResponseSourceMqEventSourceImages],
// [SubscriptionNewResponseSourceMqEventSourceKV],
// [SubscriptionNewResponseSourceMqEventSourceR2],
// [SubscriptionNewResponseSourceMqEventSourceSuperSlurper],
// [SubscriptionNewResponseSourceMqEventSourceVectorize],
// [SubscriptionNewResponseSourceMqEventSourceWorkersAIModel],
// [SubscriptionNewResponseSourceMqEventSourceWorkersBuildsWorker],
// [SubscriptionNewResponseSourceMqEventSourceWorkflowsWorkflow].
func (r SubscriptionNewResponseSource) AsUnion() SubscriptionNewResponseSourceUnion {
	return r.union
}

// Source configuration for the subscription
//
// Union satisfied by [SubscriptionNewResponseSourceMqEventSourceImages],
// [SubscriptionNewResponseSourceMqEventSourceKV],
// [SubscriptionNewResponseSourceMqEventSourceR2],
// [SubscriptionNewResponseSourceMqEventSourceSuperSlurper],
// [SubscriptionNewResponseSourceMqEventSourceVectorize],
// [SubscriptionNewResponseSourceMqEventSourceWorkersAIModel],
// [SubscriptionNewResponseSourceMqEventSourceWorkersBuildsWorker] or
// [SubscriptionNewResponseSourceMqEventSourceWorkflowsWorkflow].
type SubscriptionNewResponseSourceUnion interface {
	implementsSubscriptionNewResponseSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionNewResponseSourceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionNewResponseSourceMqEventSourceImages{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionNewResponseSourceMqEventSourceKV{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionNewResponseSourceMqEventSourceR2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionNewResponseSourceMqEventSourceSuperSlurper{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionNewResponseSourceMqEventSourceVectorize{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionNewResponseSourceMqEventSourceWorkersAIModel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionNewResponseSourceMqEventSourceWorkersBuildsWorker{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionNewResponseSourceMqEventSourceWorkflowsWorkflow{}),
		},
	)
}

type SubscriptionNewResponseSourceMqEventSourceImages struct {
	// Type of source
	Type SubscriptionNewResponseSourceMqEventSourceImagesType `json:"type"`
	JSON subscriptionNewResponseSourceMqEventSourceImagesJSON `json:"-"`
}

// subscriptionNewResponseSourceMqEventSourceImagesJSON contains the JSON metadata
// for the struct [SubscriptionNewResponseSourceMqEventSourceImages]
type subscriptionNewResponseSourceMqEventSourceImagesJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseSourceMqEventSourceImages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseSourceMqEventSourceImagesJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionNewResponseSourceMqEventSourceImages) implementsSubscriptionNewResponseSource() {}

// Type of source
type SubscriptionNewResponseSourceMqEventSourceImagesType string

const (
	SubscriptionNewResponseSourceMqEventSourceImagesTypeImages SubscriptionNewResponseSourceMqEventSourceImagesType = "images"
)

func (r SubscriptionNewResponseSourceMqEventSourceImagesType) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseSourceMqEventSourceImagesTypeImages:
		return true
	}
	return false
}

type SubscriptionNewResponseSourceMqEventSourceKV struct {
	// Type of source
	Type SubscriptionNewResponseSourceMqEventSourceKVType `json:"type"`
	JSON subscriptionNewResponseSourceMqEventSourceKVJSON `json:"-"`
}

// subscriptionNewResponseSourceMqEventSourceKVJSON contains the JSON metadata for
// the struct [SubscriptionNewResponseSourceMqEventSourceKV]
type subscriptionNewResponseSourceMqEventSourceKVJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseSourceMqEventSourceKV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseSourceMqEventSourceKVJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionNewResponseSourceMqEventSourceKV) implementsSubscriptionNewResponseSource() {}

// Type of source
type SubscriptionNewResponseSourceMqEventSourceKVType string

const (
	SubscriptionNewResponseSourceMqEventSourceKVTypeKV SubscriptionNewResponseSourceMqEventSourceKVType = "kv"
)

func (r SubscriptionNewResponseSourceMqEventSourceKVType) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseSourceMqEventSourceKVTypeKV:
		return true
	}
	return false
}

type SubscriptionNewResponseSourceMqEventSourceR2 struct {
	// Type of source
	Type SubscriptionNewResponseSourceMqEventSourceR2Type `json:"type"`
	JSON subscriptionNewResponseSourceMqEventSourceR2JSON `json:"-"`
}

// subscriptionNewResponseSourceMqEventSourceR2JSON contains the JSON metadata for
// the struct [SubscriptionNewResponseSourceMqEventSourceR2]
type subscriptionNewResponseSourceMqEventSourceR2JSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseSourceMqEventSourceR2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseSourceMqEventSourceR2JSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionNewResponseSourceMqEventSourceR2) implementsSubscriptionNewResponseSource() {}

// Type of source
type SubscriptionNewResponseSourceMqEventSourceR2Type string

const (
	SubscriptionNewResponseSourceMqEventSourceR2TypeR2 SubscriptionNewResponseSourceMqEventSourceR2Type = "r2"
)

func (r SubscriptionNewResponseSourceMqEventSourceR2Type) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseSourceMqEventSourceR2TypeR2:
		return true
	}
	return false
}

type SubscriptionNewResponseSourceMqEventSourceSuperSlurper struct {
	// Type of source
	Type SubscriptionNewResponseSourceMqEventSourceSuperSlurperType `json:"type"`
	JSON subscriptionNewResponseSourceMqEventSourceSuperSlurperJSON `json:"-"`
}

// subscriptionNewResponseSourceMqEventSourceSuperSlurperJSON contains the JSON
// metadata for the struct [SubscriptionNewResponseSourceMqEventSourceSuperSlurper]
type subscriptionNewResponseSourceMqEventSourceSuperSlurperJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseSourceMqEventSourceSuperSlurper) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseSourceMqEventSourceSuperSlurperJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionNewResponseSourceMqEventSourceSuperSlurper) implementsSubscriptionNewResponseSource() {
}

// Type of source
type SubscriptionNewResponseSourceMqEventSourceSuperSlurperType string

const (
	SubscriptionNewResponseSourceMqEventSourceSuperSlurperTypeSuperSlurper SubscriptionNewResponseSourceMqEventSourceSuperSlurperType = "superSlurper"
)

func (r SubscriptionNewResponseSourceMqEventSourceSuperSlurperType) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseSourceMqEventSourceSuperSlurperTypeSuperSlurper:
		return true
	}
	return false
}

type SubscriptionNewResponseSourceMqEventSourceVectorize struct {
	// Type of source
	Type SubscriptionNewResponseSourceMqEventSourceVectorizeType `json:"type"`
	JSON subscriptionNewResponseSourceMqEventSourceVectorizeJSON `json:"-"`
}

// subscriptionNewResponseSourceMqEventSourceVectorizeJSON contains the JSON
// metadata for the struct [SubscriptionNewResponseSourceMqEventSourceVectorize]
type subscriptionNewResponseSourceMqEventSourceVectorizeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseSourceMqEventSourceVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseSourceMqEventSourceVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionNewResponseSourceMqEventSourceVectorize) implementsSubscriptionNewResponseSource() {
}

// Type of source
type SubscriptionNewResponseSourceMqEventSourceVectorizeType string

const (
	SubscriptionNewResponseSourceMqEventSourceVectorizeTypeVectorize SubscriptionNewResponseSourceMqEventSourceVectorizeType = "vectorize"
)

func (r SubscriptionNewResponseSourceMqEventSourceVectorizeType) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseSourceMqEventSourceVectorizeTypeVectorize:
		return true
	}
	return false
}

type SubscriptionNewResponseSourceMqEventSourceWorkersAIModel struct {
	// Name of the Workers AI model
	ModelName string `json:"model_name"`
	// Type of source
	Type SubscriptionNewResponseSourceMqEventSourceWorkersAIModelType `json:"type"`
	JSON subscriptionNewResponseSourceMqEventSourceWorkersAIModelJSON `json:"-"`
}

// subscriptionNewResponseSourceMqEventSourceWorkersAIModelJSON contains the JSON
// metadata for the struct
// [SubscriptionNewResponseSourceMqEventSourceWorkersAIModel]
type subscriptionNewResponseSourceMqEventSourceWorkersAIModelJSON struct {
	ModelName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseSourceMqEventSourceWorkersAIModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseSourceMqEventSourceWorkersAIModelJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionNewResponseSourceMqEventSourceWorkersAIModel) implementsSubscriptionNewResponseSource() {
}

// Type of source
type SubscriptionNewResponseSourceMqEventSourceWorkersAIModelType string

const (
	SubscriptionNewResponseSourceMqEventSourceWorkersAIModelTypeWorkersAIModel SubscriptionNewResponseSourceMqEventSourceWorkersAIModelType = "workersAi.model"
)

func (r SubscriptionNewResponseSourceMqEventSourceWorkersAIModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseSourceMqEventSourceWorkersAIModelTypeWorkersAIModel:
		return true
	}
	return false
}

type SubscriptionNewResponseSourceMqEventSourceWorkersBuildsWorker struct {
	// Type of source
	Type SubscriptionNewResponseSourceMqEventSourceWorkersBuildsWorkerType `json:"type"`
	// Name of the worker
	WorkerName string                                                            `json:"worker_name"`
	JSON       subscriptionNewResponseSourceMqEventSourceWorkersBuildsWorkerJSON `json:"-"`
}

// subscriptionNewResponseSourceMqEventSourceWorkersBuildsWorkerJSON contains the
// JSON metadata for the struct
// [SubscriptionNewResponseSourceMqEventSourceWorkersBuildsWorker]
type subscriptionNewResponseSourceMqEventSourceWorkersBuildsWorkerJSON struct {
	Type        apijson.Field
	WorkerName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseSourceMqEventSourceWorkersBuildsWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseSourceMqEventSourceWorkersBuildsWorkerJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionNewResponseSourceMqEventSourceWorkersBuildsWorker) implementsSubscriptionNewResponseSource() {
}

// Type of source
type SubscriptionNewResponseSourceMqEventSourceWorkersBuildsWorkerType string

const (
	SubscriptionNewResponseSourceMqEventSourceWorkersBuildsWorkerTypeWorkersBuildsWorker SubscriptionNewResponseSourceMqEventSourceWorkersBuildsWorkerType = "workersBuilds.worker"
)

func (r SubscriptionNewResponseSourceMqEventSourceWorkersBuildsWorkerType) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseSourceMqEventSourceWorkersBuildsWorkerTypeWorkersBuildsWorker:
		return true
	}
	return false
}

type SubscriptionNewResponseSourceMqEventSourceWorkflowsWorkflow struct {
	// Type of source
	Type SubscriptionNewResponseSourceMqEventSourceWorkflowsWorkflowType `json:"type"`
	// Name of the workflow
	WorkflowName string                                                          `json:"workflow_name"`
	JSON         subscriptionNewResponseSourceMqEventSourceWorkflowsWorkflowJSON `json:"-"`
}

// subscriptionNewResponseSourceMqEventSourceWorkflowsWorkflowJSON contains the
// JSON metadata for the struct
// [SubscriptionNewResponseSourceMqEventSourceWorkflowsWorkflow]
type subscriptionNewResponseSourceMqEventSourceWorkflowsWorkflowJSON struct {
	Type         apijson.Field
	WorkflowName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionNewResponseSourceMqEventSourceWorkflowsWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseSourceMqEventSourceWorkflowsWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionNewResponseSourceMqEventSourceWorkflowsWorkflow) implementsSubscriptionNewResponseSource() {
}

// Type of source
type SubscriptionNewResponseSourceMqEventSourceWorkflowsWorkflowType string

const (
	SubscriptionNewResponseSourceMqEventSourceWorkflowsWorkflowTypeWorkflowsWorkflow SubscriptionNewResponseSourceMqEventSourceWorkflowsWorkflowType = "workflows.workflow"
)

func (r SubscriptionNewResponseSourceMqEventSourceWorkflowsWorkflowType) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseSourceMqEventSourceWorkflowsWorkflowTypeWorkflowsWorkflow:
		return true
	}
	return false
}

// Type of source
type SubscriptionNewResponseSourceType string

const (
	SubscriptionNewResponseSourceTypeImages              SubscriptionNewResponseSourceType = "images"
	SubscriptionNewResponseSourceTypeKV                  SubscriptionNewResponseSourceType = "kv"
	SubscriptionNewResponseSourceTypeR2                  SubscriptionNewResponseSourceType = "r2"
	SubscriptionNewResponseSourceTypeSuperSlurper        SubscriptionNewResponseSourceType = "superSlurper"
	SubscriptionNewResponseSourceTypeVectorize           SubscriptionNewResponseSourceType = "vectorize"
	SubscriptionNewResponseSourceTypeWorkersAIModel      SubscriptionNewResponseSourceType = "workersAi.model"
	SubscriptionNewResponseSourceTypeWorkersBuildsWorker SubscriptionNewResponseSourceType = "workersBuilds.worker"
	SubscriptionNewResponseSourceTypeWorkflowsWorkflow   SubscriptionNewResponseSourceType = "workflows.workflow"
)

func (r SubscriptionNewResponseSourceType) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseSourceTypeImages, SubscriptionNewResponseSourceTypeKV, SubscriptionNewResponseSourceTypeR2, SubscriptionNewResponseSourceTypeSuperSlurper, SubscriptionNewResponseSourceTypeVectorize, SubscriptionNewResponseSourceTypeWorkersAIModel, SubscriptionNewResponseSourceTypeWorkersBuildsWorker, SubscriptionNewResponseSourceTypeWorkflowsWorkflow:
		return true
	}
	return false
}

type SubscriptionUpdateResponse struct {
	// Unique identifier for the subscription
	ID string `json:"id,required"`
	// When the subscription was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Destination configuration for the subscription
	Destination SubscriptionUpdateResponseDestination `json:"destination,required"`
	// Whether the subscription is active
	Enabled bool `json:"enabled,required"`
	// List of event types this subscription handles
	Events []string `json:"events,required"`
	// When the subscription was last modified
	ModifiedAt time.Time `json:"modified_at,required" format:"date-time"`
	// Name of the subscription
	Name string `json:"name,required"`
	// Source configuration for the subscription
	Source SubscriptionUpdateResponseSource `json:"source,required"`
	JSON   subscriptionUpdateResponseJSON   `json:"-"`
}

// subscriptionUpdateResponseJSON contains the JSON metadata for the struct
// [SubscriptionUpdateResponse]
type subscriptionUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Destination apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Destination configuration for the subscription
type SubscriptionUpdateResponseDestination struct {
	// ID of the target queue
	QueueID string `json:"queue_id,required"`
	// Type of destination
	Type SubscriptionUpdateResponseDestinationType `json:"type,required"`
	JSON subscriptionUpdateResponseDestinationJSON `json:"-"`
}

// subscriptionUpdateResponseDestinationJSON contains the JSON metadata for the
// struct [SubscriptionUpdateResponseDestination]
type subscriptionUpdateResponseDestinationJSON struct {
	QueueID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseDestinationJSON) RawJSON() string {
	return r.raw
}

// Type of destination
type SubscriptionUpdateResponseDestinationType string

const (
	SubscriptionUpdateResponseDestinationTypeQueuesQueue SubscriptionUpdateResponseDestinationType = "queues.queue"
)

func (r SubscriptionUpdateResponseDestinationType) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseDestinationTypeQueuesQueue:
		return true
	}
	return false
}

// Source configuration for the subscription
type SubscriptionUpdateResponseSource struct {
	// Name of the Workers AI model
	ModelName string `json:"model_name"`
	// Type of source
	Type SubscriptionUpdateResponseSourceType `json:"type"`
	// Name of the worker
	WorkerName string `json:"worker_name"`
	// Name of the workflow
	WorkflowName string                               `json:"workflow_name"`
	JSON         subscriptionUpdateResponseSourceJSON `json:"-"`
	union        SubscriptionUpdateResponseSourceUnion
}

// subscriptionUpdateResponseSourceJSON contains the JSON metadata for the struct
// [SubscriptionUpdateResponseSource]
type subscriptionUpdateResponseSourceJSON struct {
	ModelName    apijson.Field
	Type         apijson.Field
	WorkerName   apijson.Field
	WorkflowName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r subscriptionUpdateResponseSourceJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionUpdateResponseSource) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionUpdateResponseSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SubscriptionUpdateResponseSourceUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionUpdateResponseSourceMqEventSourceImages],
// [SubscriptionUpdateResponseSourceMqEventSourceKV],
// [SubscriptionUpdateResponseSourceMqEventSourceR2],
// [SubscriptionUpdateResponseSourceMqEventSourceSuperSlurper],
// [SubscriptionUpdateResponseSourceMqEventSourceVectorize],
// [SubscriptionUpdateResponseSourceMqEventSourceWorkersAIModel],
// [SubscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorker],
// [SubscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflow].
func (r SubscriptionUpdateResponseSource) AsUnion() SubscriptionUpdateResponseSourceUnion {
	return r.union
}

// Source configuration for the subscription
//
// Union satisfied by [SubscriptionUpdateResponseSourceMqEventSourceImages],
// [SubscriptionUpdateResponseSourceMqEventSourceKV],
// [SubscriptionUpdateResponseSourceMqEventSourceR2],
// [SubscriptionUpdateResponseSourceMqEventSourceSuperSlurper],
// [SubscriptionUpdateResponseSourceMqEventSourceVectorize],
// [SubscriptionUpdateResponseSourceMqEventSourceWorkersAIModel],
// [SubscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorker] or
// [SubscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflow].
type SubscriptionUpdateResponseSourceUnion interface {
	implementsSubscriptionUpdateResponseSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionUpdateResponseSourceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionUpdateResponseSourceMqEventSourceImages{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionUpdateResponseSourceMqEventSourceKV{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionUpdateResponseSourceMqEventSourceR2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionUpdateResponseSourceMqEventSourceSuperSlurper{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionUpdateResponseSourceMqEventSourceVectorize{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionUpdateResponseSourceMqEventSourceWorkersAIModel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorker{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflow{}),
		},
	)
}

type SubscriptionUpdateResponseSourceMqEventSourceImages struct {
	// Type of source
	Type SubscriptionUpdateResponseSourceMqEventSourceImagesType `json:"type"`
	JSON subscriptionUpdateResponseSourceMqEventSourceImagesJSON `json:"-"`
}

// subscriptionUpdateResponseSourceMqEventSourceImagesJSON contains the JSON
// metadata for the struct [SubscriptionUpdateResponseSourceMqEventSourceImages]
type subscriptionUpdateResponseSourceMqEventSourceImagesJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseSourceMqEventSourceImages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseSourceMqEventSourceImagesJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionUpdateResponseSourceMqEventSourceImages) implementsSubscriptionUpdateResponseSource() {
}

// Type of source
type SubscriptionUpdateResponseSourceMqEventSourceImagesType string

const (
	SubscriptionUpdateResponseSourceMqEventSourceImagesTypeImages SubscriptionUpdateResponseSourceMqEventSourceImagesType = "images"
)

func (r SubscriptionUpdateResponseSourceMqEventSourceImagesType) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseSourceMqEventSourceImagesTypeImages:
		return true
	}
	return false
}

type SubscriptionUpdateResponseSourceMqEventSourceKV struct {
	// Type of source
	Type SubscriptionUpdateResponseSourceMqEventSourceKVType `json:"type"`
	JSON subscriptionUpdateResponseSourceMqEventSourceKVJSON `json:"-"`
}

// subscriptionUpdateResponseSourceMqEventSourceKVJSON contains the JSON metadata
// for the struct [SubscriptionUpdateResponseSourceMqEventSourceKV]
type subscriptionUpdateResponseSourceMqEventSourceKVJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseSourceMqEventSourceKV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseSourceMqEventSourceKVJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionUpdateResponseSourceMqEventSourceKV) implementsSubscriptionUpdateResponseSource() {
}

// Type of source
type SubscriptionUpdateResponseSourceMqEventSourceKVType string

const (
	SubscriptionUpdateResponseSourceMqEventSourceKVTypeKV SubscriptionUpdateResponseSourceMqEventSourceKVType = "kv"
)

func (r SubscriptionUpdateResponseSourceMqEventSourceKVType) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseSourceMqEventSourceKVTypeKV:
		return true
	}
	return false
}

type SubscriptionUpdateResponseSourceMqEventSourceR2 struct {
	// Type of source
	Type SubscriptionUpdateResponseSourceMqEventSourceR2Type `json:"type"`
	JSON subscriptionUpdateResponseSourceMqEventSourceR2JSON `json:"-"`
}

// subscriptionUpdateResponseSourceMqEventSourceR2JSON contains the JSON metadata
// for the struct [SubscriptionUpdateResponseSourceMqEventSourceR2]
type subscriptionUpdateResponseSourceMqEventSourceR2JSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseSourceMqEventSourceR2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseSourceMqEventSourceR2JSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionUpdateResponseSourceMqEventSourceR2) implementsSubscriptionUpdateResponseSource() {
}

// Type of source
type SubscriptionUpdateResponseSourceMqEventSourceR2Type string

const (
	SubscriptionUpdateResponseSourceMqEventSourceR2TypeR2 SubscriptionUpdateResponseSourceMqEventSourceR2Type = "r2"
)

func (r SubscriptionUpdateResponseSourceMqEventSourceR2Type) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseSourceMqEventSourceR2TypeR2:
		return true
	}
	return false
}

type SubscriptionUpdateResponseSourceMqEventSourceSuperSlurper struct {
	// Type of source
	Type SubscriptionUpdateResponseSourceMqEventSourceSuperSlurperType `json:"type"`
	JSON subscriptionUpdateResponseSourceMqEventSourceSuperSlurperJSON `json:"-"`
}

// subscriptionUpdateResponseSourceMqEventSourceSuperSlurperJSON contains the JSON
// metadata for the struct
// [SubscriptionUpdateResponseSourceMqEventSourceSuperSlurper]
type subscriptionUpdateResponseSourceMqEventSourceSuperSlurperJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseSourceMqEventSourceSuperSlurper) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseSourceMqEventSourceSuperSlurperJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionUpdateResponseSourceMqEventSourceSuperSlurper) implementsSubscriptionUpdateResponseSource() {
}

// Type of source
type SubscriptionUpdateResponseSourceMqEventSourceSuperSlurperType string

const (
	SubscriptionUpdateResponseSourceMqEventSourceSuperSlurperTypeSuperSlurper SubscriptionUpdateResponseSourceMqEventSourceSuperSlurperType = "superSlurper"
)

func (r SubscriptionUpdateResponseSourceMqEventSourceSuperSlurperType) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseSourceMqEventSourceSuperSlurperTypeSuperSlurper:
		return true
	}
	return false
}

type SubscriptionUpdateResponseSourceMqEventSourceVectorize struct {
	// Type of source
	Type SubscriptionUpdateResponseSourceMqEventSourceVectorizeType `json:"type"`
	JSON subscriptionUpdateResponseSourceMqEventSourceVectorizeJSON `json:"-"`
}

// subscriptionUpdateResponseSourceMqEventSourceVectorizeJSON contains the JSON
// metadata for the struct [SubscriptionUpdateResponseSourceMqEventSourceVectorize]
type subscriptionUpdateResponseSourceMqEventSourceVectorizeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseSourceMqEventSourceVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseSourceMqEventSourceVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionUpdateResponseSourceMqEventSourceVectorize) implementsSubscriptionUpdateResponseSource() {
}

// Type of source
type SubscriptionUpdateResponseSourceMqEventSourceVectorizeType string

const (
	SubscriptionUpdateResponseSourceMqEventSourceVectorizeTypeVectorize SubscriptionUpdateResponseSourceMqEventSourceVectorizeType = "vectorize"
)

func (r SubscriptionUpdateResponseSourceMqEventSourceVectorizeType) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseSourceMqEventSourceVectorizeTypeVectorize:
		return true
	}
	return false
}

type SubscriptionUpdateResponseSourceMqEventSourceWorkersAIModel struct {
	// Name of the Workers AI model
	ModelName string `json:"model_name"`
	// Type of source
	Type SubscriptionUpdateResponseSourceMqEventSourceWorkersAIModelType `json:"type"`
	JSON subscriptionUpdateResponseSourceMqEventSourceWorkersAIModelJSON `json:"-"`
}

// subscriptionUpdateResponseSourceMqEventSourceWorkersAIModelJSON contains the
// JSON metadata for the struct
// [SubscriptionUpdateResponseSourceMqEventSourceWorkersAIModel]
type subscriptionUpdateResponseSourceMqEventSourceWorkersAIModelJSON struct {
	ModelName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseSourceMqEventSourceWorkersAIModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseSourceMqEventSourceWorkersAIModelJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionUpdateResponseSourceMqEventSourceWorkersAIModel) implementsSubscriptionUpdateResponseSource() {
}

// Type of source
type SubscriptionUpdateResponseSourceMqEventSourceWorkersAIModelType string

const (
	SubscriptionUpdateResponseSourceMqEventSourceWorkersAIModelTypeWorkersAIModel SubscriptionUpdateResponseSourceMqEventSourceWorkersAIModelType = "workersAi.model"
)

func (r SubscriptionUpdateResponseSourceMqEventSourceWorkersAIModelType) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseSourceMqEventSourceWorkersAIModelTypeWorkersAIModel:
		return true
	}
	return false
}

type SubscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorker struct {
	// Type of source
	Type SubscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorkerType `json:"type"`
	// Name of the worker
	WorkerName string                                                               `json:"worker_name"`
	JSON       subscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorkerJSON `json:"-"`
}

// subscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorkerJSON contains
// the JSON metadata for the struct
// [SubscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorker]
type subscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorkerJSON struct {
	Type        apijson.Field
	WorkerName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorkerJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorker) implementsSubscriptionUpdateResponseSource() {
}

// Type of source
type SubscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorkerType string

const (
	SubscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorkerTypeWorkersBuildsWorker SubscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorkerType = "workersBuilds.worker"
)

func (r SubscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorkerType) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseSourceMqEventSourceWorkersBuildsWorkerTypeWorkersBuildsWorker:
		return true
	}
	return false
}

type SubscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflow struct {
	// Type of source
	Type SubscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflowType `json:"type"`
	// Name of the workflow
	WorkflowName string                                                             `json:"workflow_name"`
	JSON         subscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflowJSON `json:"-"`
}

// subscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflowJSON contains the
// JSON metadata for the struct
// [SubscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflow]
type subscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflowJSON struct {
	Type         apijson.Field
	WorkflowName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflow) implementsSubscriptionUpdateResponseSource() {
}

// Type of source
type SubscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflowType string

const (
	SubscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflowTypeWorkflowsWorkflow SubscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflowType = "workflows.workflow"
)

func (r SubscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflowType) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseSourceMqEventSourceWorkflowsWorkflowTypeWorkflowsWorkflow:
		return true
	}
	return false
}

// Type of source
type SubscriptionUpdateResponseSourceType string

const (
	SubscriptionUpdateResponseSourceTypeImages              SubscriptionUpdateResponseSourceType = "images"
	SubscriptionUpdateResponseSourceTypeKV                  SubscriptionUpdateResponseSourceType = "kv"
	SubscriptionUpdateResponseSourceTypeR2                  SubscriptionUpdateResponseSourceType = "r2"
	SubscriptionUpdateResponseSourceTypeSuperSlurper        SubscriptionUpdateResponseSourceType = "superSlurper"
	SubscriptionUpdateResponseSourceTypeVectorize           SubscriptionUpdateResponseSourceType = "vectorize"
	SubscriptionUpdateResponseSourceTypeWorkersAIModel      SubscriptionUpdateResponseSourceType = "workersAi.model"
	SubscriptionUpdateResponseSourceTypeWorkersBuildsWorker SubscriptionUpdateResponseSourceType = "workersBuilds.worker"
	SubscriptionUpdateResponseSourceTypeWorkflowsWorkflow   SubscriptionUpdateResponseSourceType = "workflows.workflow"
)

func (r SubscriptionUpdateResponseSourceType) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseSourceTypeImages, SubscriptionUpdateResponseSourceTypeKV, SubscriptionUpdateResponseSourceTypeR2, SubscriptionUpdateResponseSourceTypeSuperSlurper, SubscriptionUpdateResponseSourceTypeVectorize, SubscriptionUpdateResponseSourceTypeWorkersAIModel, SubscriptionUpdateResponseSourceTypeWorkersBuildsWorker, SubscriptionUpdateResponseSourceTypeWorkflowsWorkflow:
		return true
	}
	return false
}

type SubscriptionListResponse struct {
	// Unique identifier for the subscription
	ID string `json:"id,required"`
	// When the subscription was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Destination configuration for the subscription
	Destination SubscriptionListResponseDestination `json:"destination,required"`
	// Whether the subscription is active
	Enabled bool `json:"enabled,required"`
	// List of event types this subscription handles
	Events []string `json:"events,required"`
	// When the subscription was last modified
	ModifiedAt time.Time `json:"modified_at,required" format:"date-time"`
	// Name of the subscription
	Name string `json:"name,required"`
	// Source configuration for the subscription
	Source SubscriptionListResponseSource `json:"source,required"`
	JSON   subscriptionListResponseJSON   `json:"-"`
}

// subscriptionListResponseJSON contains the JSON metadata for the struct
// [SubscriptionListResponse]
type subscriptionListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Destination apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseJSON) RawJSON() string {
	return r.raw
}

// Destination configuration for the subscription
type SubscriptionListResponseDestination struct {
	// ID of the target queue
	QueueID string `json:"queue_id,required"`
	// Type of destination
	Type SubscriptionListResponseDestinationType `json:"type,required"`
	JSON subscriptionListResponseDestinationJSON `json:"-"`
}

// subscriptionListResponseDestinationJSON contains the JSON metadata for the
// struct [SubscriptionListResponseDestination]
type subscriptionListResponseDestinationJSON struct {
	QueueID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseDestinationJSON) RawJSON() string {
	return r.raw
}

// Type of destination
type SubscriptionListResponseDestinationType string

const (
	SubscriptionListResponseDestinationTypeQueuesQueue SubscriptionListResponseDestinationType = "queues.queue"
)

func (r SubscriptionListResponseDestinationType) IsKnown() bool {
	switch r {
	case SubscriptionListResponseDestinationTypeQueuesQueue:
		return true
	}
	return false
}

// Source configuration for the subscription
type SubscriptionListResponseSource struct {
	// Name of the Workers AI model
	ModelName string `json:"model_name"`
	// Type of source
	Type SubscriptionListResponseSourceType `json:"type"`
	// Name of the worker
	WorkerName string `json:"worker_name"`
	// Name of the workflow
	WorkflowName string                             `json:"workflow_name"`
	JSON         subscriptionListResponseSourceJSON `json:"-"`
	union        SubscriptionListResponseSourceUnion
}

// subscriptionListResponseSourceJSON contains the JSON metadata for the struct
// [SubscriptionListResponseSource]
type subscriptionListResponseSourceJSON struct {
	ModelName    apijson.Field
	Type         apijson.Field
	WorkerName   apijson.Field
	WorkflowName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r subscriptionListResponseSourceJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionListResponseSource) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionListResponseSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SubscriptionListResponseSourceUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionListResponseSourceMqEventSourceImages],
// [SubscriptionListResponseSourceMqEventSourceKV],
// [SubscriptionListResponseSourceMqEventSourceR2],
// [SubscriptionListResponseSourceMqEventSourceSuperSlurper],
// [SubscriptionListResponseSourceMqEventSourceVectorize],
// [SubscriptionListResponseSourceMqEventSourceWorkersAIModel],
// [SubscriptionListResponseSourceMqEventSourceWorkersBuildsWorker],
// [SubscriptionListResponseSourceMqEventSourceWorkflowsWorkflow].
func (r SubscriptionListResponseSource) AsUnion() SubscriptionListResponseSourceUnion {
	return r.union
}

// Source configuration for the subscription
//
// Union satisfied by [SubscriptionListResponseSourceMqEventSourceImages],
// [SubscriptionListResponseSourceMqEventSourceKV],
// [SubscriptionListResponseSourceMqEventSourceR2],
// [SubscriptionListResponseSourceMqEventSourceSuperSlurper],
// [SubscriptionListResponseSourceMqEventSourceVectorize],
// [SubscriptionListResponseSourceMqEventSourceWorkersAIModel],
// [SubscriptionListResponseSourceMqEventSourceWorkersBuildsWorker] or
// [SubscriptionListResponseSourceMqEventSourceWorkflowsWorkflow].
type SubscriptionListResponseSourceUnion interface {
	implementsSubscriptionListResponseSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionListResponseSourceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionListResponseSourceMqEventSourceImages{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionListResponseSourceMqEventSourceKV{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionListResponseSourceMqEventSourceR2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionListResponseSourceMqEventSourceSuperSlurper{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionListResponseSourceMqEventSourceVectorize{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionListResponseSourceMqEventSourceWorkersAIModel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionListResponseSourceMqEventSourceWorkersBuildsWorker{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionListResponseSourceMqEventSourceWorkflowsWorkflow{}),
		},
	)
}

type SubscriptionListResponseSourceMqEventSourceImages struct {
	// Type of source
	Type SubscriptionListResponseSourceMqEventSourceImagesType `json:"type"`
	JSON subscriptionListResponseSourceMqEventSourceImagesJSON `json:"-"`
}

// subscriptionListResponseSourceMqEventSourceImagesJSON contains the JSON metadata
// for the struct [SubscriptionListResponseSourceMqEventSourceImages]
type subscriptionListResponseSourceMqEventSourceImagesJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseSourceMqEventSourceImages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseSourceMqEventSourceImagesJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionListResponseSourceMqEventSourceImages) implementsSubscriptionListResponseSource() {
}

// Type of source
type SubscriptionListResponseSourceMqEventSourceImagesType string

const (
	SubscriptionListResponseSourceMqEventSourceImagesTypeImages SubscriptionListResponseSourceMqEventSourceImagesType = "images"
)

func (r SubscriptionListResponseSourceMqEventSourceImagesType) IsKnown() bool {
	switch r {
	case SubscriptionListResponseSourceMqEventSourceImagesTypeImages:
		return true
	}
	return false
}

type SubscriptionListResponseSourceMqEventSourceKV struct {
	// Type of source
	Type SubscriptionListResponseSourceMqEventSourceKVType `json:"type"`
	JSON subscriptionListResponseSourceMqEventSourceKVJSON `json:"-"`
}

// subscriptionListResponseSourceMqEventSourceKVJSON contains the JSON metadata for
// the struct [SubscriptionListResponseSourceMqEventSourceKV]
type subscriptionListResponseSourceMqEventSourceKVJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseSourceMqEventSourceKV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseSourceMqEventSourceKVJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionListResponseSourceMqEventSourceKV) implementsSubscriptionListResponseSource() {}

// Type of source
type SubscriptionListResponseSourceMqEventSourceKVType string

const (
	SubscriptionListResponseSourceMqEventSourceKVTypeKV SubscriptionListResponseSourceMqEventSourceKVType = "kv"
)

func (r SubscriptionListResponseSourceMqEventSourceKVType) IsKnown() bool {
	switch r {
	case SubscriptionListResponseSourceMqEventSourceKVTypeKV:
		return true
	}
	return false
}

type SubscriptionListResponseSourceMqEventSourceR2 struct {
	// Type of source
	Type SubscriptionListResponseSourceMqEventSourceR2Type `json:"type"`
	JSON subscriptionListResponseSourceMqEventSourceR2JSON `json:"-"`
}

// subscriptionListResponseSourceMqEventSourceR2JSON contains the JSON metadata for
// the struct [SubscriptionListResponseSourceMqEventSourceR2]
type subscriptionListResponseSourceMqEventSourceR2JSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseSourceMqEventSourceR2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseSourceMqEventSourceR2JSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionListResponseSourceMqEventSourceR2) implementsSubscriptionListResponseSource() {}

// Type of source
type SubscriptionListResponseSourceMqEventSourceR2Type string

const (
	SubscriptionListResponseSourceMqEventSourceR2TypeR2 SubscriptionListResponseSourceMqEventSourceR2Type = "r2"
)

func (r SubscriptionListResponseSourceMqEventSourceR2Type) IsKnown() bool {
	switch r {
	case SubscriptionListResponseSourceMqEventSourceR2TypeR2:
		return true
	}
	return false
}

type SubscriptionListResponseSourceMqEventSourceSuperSlurper struct {
	// Type of source
	Type SubscriptionListResponseSourceMqEventSourceSuperSlurperType `json:"type"`
	JSON subscriptionListResponseSourceMqEventSourceSuperSlurperJSON `json:"-"`
}

// subscriptionListResponseSourceMqEventSourceSuperSlurperJSON contains the JSON
// metadata for the struct
// [SubscriptionListResponseSourceMqEventSourceSuperSlurper]
type subscriptionListResponseSourceMqEventSourceSuperSlurperJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseSourceMqEventSourceSuperSlurper) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseSourceMqEventSourceSuperSlurperJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionListResponseSourceMqEventSourceSuperSlurper) implementsSubscriptionListResponseSource() {
}

// Type of source
type SubscriptionListResponseSourceMqEventSourceSuperSlurperType string

const (
	SubscriptionListResponseSourceMqEventSourceSuperSlurperTypeSuperSlurper SubscriptionListResponseSourceMqEventSourceSuperSlurperType = "superSlurper"
)

func (r SubscriptionListResponseSourceMqEventSourceSuperSlurperType) IsKnown() bool {
	switch r {
	case SubscriptionListResponseSourceMqEventSourceSuperSlurperTypeSuperSlurper:
		return true
	}
	return false
}

type SubscriptionListResponseSourceMqEventSourceVectorize struct {
	// Type of source
	Type SubscriptionListResponseSourceMqEventSourceVectorizeType `json:"type"`
	JSON subscriptionListResponseSourceMqEventSourceVectorizeJSON `json:"-"`
}

// subscriptionListResponseSourceMqEventSourceVectorizeJSON contains the JSON
// metadata for the struct [SubscriptionListResponseSourceMqEventSourceVectorize]
type subscriptionListResponseSourceMqEventSourceVectorizeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseSourceMqEventSourceVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseSourceMqEventSourceVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionListResponseSourceMqEventSourceVectorize) implementsSubscriptionListResponseSource() {
}

// Type of source
type SubscriptionListResponseSourceMqEventSourceVectorizeType string

const (
	SubscriptionListResponseSourceMqEventSourceVectorizeTypeVectorize SubscriptionListResponseSourceMqEventSourceVectorizeType = "vectorize"
)

func (r SubscriptionListResponseSourceMqEventSourceVectorizeType) IsKnown() bool {
	switch r {
	case SubscriptionListResponseSourceMqEventSourceVectorizeTypeVectorize:
		return true
	}
	return false
}

type SubscriptionListResponseSourceMqEventSourceWorkersAIModel struct {
	// Name of the Workers AI model
	ModelName string `json:"model_name"`
	// Type of source
	Type SubscriptionListResponseSourceMqEventSourceWorkersAIModelType `json:"type"`
	JSON subscriptionListResponseSourceMqEventSourceWorkersAIModelJSON `json:"-"`
}

// subscriptionListResponseSourceMqEventSourceWorkersAIModelJSON contains the JSON
// metadata for the struct
// [SubscriptionListResponseSourceMqEventSourceWorkersAIModel]
type subscriptionListResponseSourceMqEventSourceWorkersAIModelJSON struct {
	ModelName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseSourceMqEventSourceWorkersAIModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseSourceMqEventSourceWorkersAIModelJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionListResponseSourceMqEventSourceWorkersAIModel) implementsSubscriptionListResponseSource() {
}

// Type of source
type SubscriptionListResponseSourceMqEventSourceWorkersAIModelType string

const (
	SubscriptionListResponseSourceMqEventSourceWorkersAIModelTypeWorkersAIModel SubscriptionListResponseSourceMqEventSourceWorkersAIModelType = "workersAi.model"
)

func (r SubscriptionListResponseSourceMqEventSourceWorkersAIModelType) IsKnown() bool {
	switch r {
	case SubscriptionListResponseSourceMqEventSourceWorkersAIModelTypeWorkersAIModel:
		return true
	}
	return false
}

type SubscriptionListResponseSourceMqEventSourceWorkersBuildsWorker struct {
	// Type of source
	Type SubscriptionListResponseSourceMqEventSourceWorkersBuildsWorkerType `json:"type"`
	// Name of the worker
	WorkerName string                                                             `json:"worker_name"`
	JSON       subscriptionListResponseSourceMqEventSourceWorkersBuildsWorkerJSON `json:"-"`
}

// subscriptionListResponseSourceMqEventSourceWorkersBuildsWorkerJSON contains the
// JSON metadata for the struct
// [SubscriptionListResponseSourceMqEventSourceWorkersBuildsWorker]
type subscriptionListResponseSourceMqEventSourceWorkersBuildsWorkerJSON struct {
	Type        apijson.Field
	WorkerName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseSourceMqEventSourceWorkersBuildsWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseSourceMqEventSourceWorkersBuildsWorkerJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionListResponseSourceMqEventSourceWorkersBuildsWorker) implementsSubscriptionListResponseSource() {
}

// Type of source
type SubscriptionListResponseSourceMqEventSourceWorkersBuildsWorkerType string

const (
	SubscriptionListResponseSourceMqEventSourceWorkersBuildsWorkerTypeWorkersBuildsWorker SubscriptionListResponseSourceMqEventSourceWorkersBuildsWorkerType = "workersBuilds.worker"
)

func (r SubscriptionListResponseSourceMqEventSourceWorkersBuildsWorkerType) IsKnown() bool {
	switch r {
	case SubscriptionListResponseSourceMqEventSourceWorkersBuildsWorkerTypeWorkersBuildsWorker:
		return true
	}
	return false
}

type SubscriptionListResponseSourceMqEventSourceWorkflowsWorkflow struct {
	// Type of source
	Type SubscriptionListResponseSourceMqEventSourceWorkflowsWorkflowType `json:"type"`
	// Name of the workflow
	WorkflowName string                                                           `json:"workflow_name"`
	JSON         subscriptionListResponseSourceMqEventSourceWorkflowsWorkflowJSON `json:"-"`
}

// subscriptionListResponseSourceMqEventSourceWorkflowsWorkflowJSON contains the
// JSON metadata for the struct
// [SubscriptionListResponseSourceMqEventSourceWorkflowsWorkflow]
type subscriptionListResponseSourceMqEventSourceWorkflowsWorkflowJSON struct {
	Type         apijson.Field
	WorkflowName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionListResponseSourceMqEventSourceWorkflowsWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseSourceMqEventSourceWorkflowsWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionListResponseSourceMqEventSourceWorkflowsWorkflow) implementsSubscriptionListResponseSource() {
}

// Type of source
type SubscriptionListResponseSourceMqEventSourceWorkflowsWorkflowType string

const (
	SubscriptionListResponseSourceMqEventSourceWorkflowsWorkflowTypeWorkflowsWorkflow SubscriptionListResponseSourceMqEventSourceWorkflowsWorkflowType = "workflows.workflow"
)

func (r SubscriptionListResponseSourceMqEventSourceWorkflowsWorkflowType) IsKnown() bool {
	switch r {
	case SubscriptionListResponseSourceMqEventSourceWorkflowsWorkflowTypeWorkflowsWorkflow:
		return true
	}
	return false
}

// Type of source
type SubscriptionListResponseSourceType string

const (
	SubscriptionListResponseSourceTypeImages              SubscriptionListResponseSourceType = "images"
	SubscriptionListResponseSourceTypeKV                  SubscriptionListResponseSourceType = "kv"
	SubscriptionListResponseSourceTypeR2                  SubscriptionListResponseSourceType = "r2"
	SubscriptionListResponseSourceTypeSuperSlurper        SubscriptionListResponseSourceType = "superSlurper"
	SubscriptionListResponseSourceTypeVectorize           SubscriptionListResponseSourceType = "vectorize"
	SubscriptionListResponseSourceTypeWorkersAIModel      SubscriptionListResponseSourceType = "workersAi.model"
	SubscriptionListResponseSourceTypeWorkersBuildsWorker SubscriptionListResponseSourceType = "workersBuilds.worker"
	SubscriptionListResponseSourceTypeWorkflowsWorkflow   SubscriptionListResponseSourceType = "workflows.workflow"
)

func (r SubscriptionListResponseSourceType) IsKnown() bool {
	switch r {
	case SubscriptionListResponseSourceTypeImages, SubscriptionListResponseSourceTypeKV, SubscriptionListResponseSourceTypeR2, SubscriptionListResponseSourceTypeSuperSlurper, SubscriptionListResponseSourceTypeVectorize, SubscriptionListResponseSourceTypeWorkersAIModel, SubscriptionListResponseSourceTypeWorkersBuildsWorker, SubscriptionListResponseSourceTypeWorkflowsWorkflow:
		return true
	}
	return false
}

type SubscriptionDeleteResponse struct {
	// Unique identifier for the subscription
	ID string `json:"id,required"`
	// When the subscription was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Destination configuration for the subscription
	Destination SubscriptionDeleteResponseDestination `json:"destination,required"`
	// Whether the subscription is active
	Enabled bool `json:"enabled,required"`
	// List of event types this subscription handles
	Events []string `json:"events,required"`
	// When the subscription was last modified
	ModifiedAt time.Time `json:"modified_at,required" format:"date-time"`
	// Name of the subscription
	Name string `json:"name,required"`
	// Source configuration for the subscription
	Source SubscriptionDeleteResponseSource `json:"source,required"`
	JSON   subscriptionDeleteResponseJSON   `json:"-"`
}

// subscriptionDeleteResponseJSON contains the JSON metadata for the struct
// [SubscriptionDeleteResponse]
type subscriptionDeleteResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Destination apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Destination configuration for the subscription
type SubscriptionDeleteResponseDestination struct {
	// ID of the target queue
	QueueID string `json:"queue_id,required"`
	// Type of destination
	Type SubscriptionDeleteResponseDestinationType `json:"type,required"`
	JSON subscriptionDeleteResponseDestinationJSON `json:"-"`
}

// subscriptionDeleteResponseDestinationJSON contains the JSON metadata for the
// struct [SubscriptionDeleteResponseDestination]
type subscriptionDeleteResponseDestinationJSON struct {
	QueueID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseDestinationJSON) RawJSON() string {
	return r.raw
}

// Type of destination
type SubscriptionDeleteResponseDestinationType string

const (
	SubscriptionDeleteResponseDestinationTypeQueuesQueue SubscriptionDeleteResponseDestinationType = "queues.queue"
)

func (r SubscriptionDeleteResponseDestinationType) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseDestinationTypeQueuesQueue:
		return true
	}
	return false
}

// Source configuration for the subscription
type SubscriptionDeleteResponseSource struct {
	// Name of the Workers AI model
	ModelName string `json:"model_name"`
	// Type of source
	Type SubscriptionDeleteResponseSourceType `json:"type"`
	// Name of the worker
	WorkerName string `json:"worker_name"`
	// Name of the workflow
	WorkflowName string                               `json:"workflow_name"`
	JSON         subscriptionDeleteResponseSourceJSON `json:"-"`
	union        SubscriptionDeleteResponseSourceUnion
}

// subscriptionDeleteResponseSourceJSON contains the JSON metadata for the struct
// [SubscriptionDeleteResponseSource]
type subscriptionDeleteResponseSourceJSON struct {
	ModelName    apijson.Field
	Type         apijson.Field
	WorkerName   apijson.Field
	WorkflowName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r subscriptionDeleteResponseSourceJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionDeleteResponseSource) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionDeleteResponseSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SubscriptionDeleteResponseSourceUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionDeleteResponseSourceMqEventSourceImages],
// [SubscriptionDeleteResponseSourceMqEventSourceKV],
// [SubscriptionDeleteResponseSourceMqEventSourceR2],
// [SubscriptionDeleteResponseSourceMqEventSourceSuperSlurper],
// [SubscriptionDeleteResponseSourceMqEventSourceVectorize],
// [SubscriptionDeleteResponseSourceMqEventSourceWorkersAIModel],
// [SubscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorker],
// [SubscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflow].
func (r SubscriptionDeleteResponseSource) AsUnion() SubscriptionDeleteResponseSourceUnion {
	return r.union
}

// Source configuration for the subscription
//
// Union satisfied by [SubscriptionDeleteResponseSourceMqEventSourceImages],
// [SubscriptionDeleteResponseSourceMqEventSourceKV],
// [SubscriptionDeleteResponseSourceMqEventSourceR2],
// [SubscriptionDeleteResponseSourceMqEventSourceSuperSlurper],
// [SubscriptionDeleteResponseSourceMqEventSourceVectorize],
// [SubscriptionDeleteResponseSourceMqEventSourceWorkersAIModel],
// [SubscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorker] or
// [SubscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflow].
type SubscriptionDeleteResponseSourceUnion interface {
	implementsSubscriptionDeleteResponseSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionDeleteResponseSourceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionDeleteResponseSourceMqEventSourceImages{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionDeleteResponseSourceMqEventSourceKV{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionDeleteResponseSourceMqEventSourceR2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionDeleteResponseSourceMqEventSourceSuperSlurper{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionDeleteResponseSourceMqEventSourceVectorize{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionDeleteResponseSourceMqEventSourceWorkersAIModel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorker{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflow{}),
		},
	)
}

type SubscriptionDeleteResponseSourceMqEventSourceImages struct {
	// Type of source
	Type SubscriptionDeleteResponseSourceMqEventSourceImagesType `json:"type"`
	JSON subscriptionDeleteResponseSourceMqEventSourceImagesJSON `json:"-"`
}

// subscriptionDeleteResponseSourceMqEventSourceImagesJSON contains the JSON
// metadata for the struct [SubscriptionDeleteResponseSourceMqEventSourceImages]
type subscriptionDeleteResponseSourceMqEventSourceImagesJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseSourceMqEventSourceImages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseSourceMqEventSourceImagesJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionDeleteResponseSourceMqEventSourceImages) implementsSubscriptionDeleteResponseSource() {
}

// Type of source
type SubscriptionDeleteResponseSourceMqEventSourceImagesType string

const (
	SubscriptionDeleteResponseSourceMqEventSourceImagesTypeImages SubscriptionDeleteResponseSourceMqEventSourceImagesType = "images"
)

func (r SubscriptionDeleteResponseSourceMqEventSourceImagesType) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseSourceMqEventSourceImagesTypeImages:
		return true
	}
	return false
}

type SubscriptionDeleteResponseSourceMqEventSourceKV struct {
	// Type of source
	Type SubscriptionDeleteResponseSourceMqEventSourceKVType `json:"type"`
	JSON subscriptionDeleteResponseSourceMqEventSourceKVJSON `json:"-"`
}

// subscriptionDeleteResponseSourceMqEventSourceKVJSON contains the JSON metadata
// for the struct [SubscriptionDeleteResponseSourceMqEventSourceKV]
type subscriptionDeleteResponseSourceMqEventSourceKVJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseSourceMqEventSourceKV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseSourceMqEventSourceKVJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionDeleteResponseSourceMqEventSourceKV) implementsSubscriptionDeleteResponseSource() {
}

// Type of source
type SubscriptionDeleteResponseSourceMqEventSourceKVType string

const (
	SubscriptionDeleteResponseSourceMqEventSourceKVTypeKV SubscriptionDeleteResponseSourceMqEventSourceKVType = "kv"
)

func (r SubscriptionDeleteResponseSourceMqEventSourceKVType) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseSourceMqEventSourceKVTypeKV:
		return true
	}
	return false
}

type SubscriptionDeleteResponseSourceMqEventSourceR2 struct {
	// Type of source
	Type SubscriptionDeleteResponseSourceMqEventSourceR2Type `json:"type"`
	JSON subscriptionDeleteResponseSourceMqEventSourceR2JSON `json:"-"`
}

// subscriptionDeleteResponseSourceMqEventSourceR2JSON contains the JSON metadata
// for the struct [SubscriptionDeleteResponseSourceMqEventSourceR2]
type subscriptionDeleteResponseSourceMqEventSourceR2JSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseSourceMqEventSourceR2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseSourceMqEventSourceR2JSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionDeleteResponseSourceMqEventSourceR2) implementsSubscriptionDeleteResponseSource() {
}

// Type of source
type SubscriptionDeleteResponseSourceMqEventSourceR2Type string

const (
	SubscriptionDeleteResponseSourceMqEventSourceR2TypeR2 SubscriptionDeleteResponseSourceMqEventSourceR2Type = "r2"
)

func (r SubscriptionDeleteResponseSourceMqEventSourceR2Type) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseSourceMqEventSourceR2TypeR2:
		return true
	}
	return false
}

type SubscriptionDeleteResponseSourceMqEventSourceSuperSlurper struct {
	// Type of source
	Type SubscriptionDeleteResponseSourceMqEventSourceSuperSlurperType `json:"type"`
	JSON subscriptionDeleteResponseSourceMqEventSourceSuperSlurperJSON `json:"-"`
}

// subscriptionDeleteResponseSourceMqEventSourceSuperSlurperJSON contains the JSON
// metadata for the struct
// [SubscriptionDeleteResponseSourceMqEventSourceSuperSlurper]
type subscriptionDeleteResponseSourceMqEventSourceSuperSlurperJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseSourceMqEventSourceSuperSlurper) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseSourceMqEventSourceSuperSlurperJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionDeleteResponseSourceMqEventSourceSuperSlurper) implementsSubscriptionDeleteResponseSource() {
}

// Type of source
type SubscriptionDeleteResponseSourceMqEventSourceSuperSlurperType string

const (
	SubscriptionDeleteResponseSourceMqEventSourceSuperSlurperTypeSuperSlurper SubscriptionDeleteResponseSourceMqEventSourceSuperSlurperType = "superSlurper"
)

func (r SubscriptionDeleteResponseSourceMqEventSourceSuperSlurperType) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseSourceMqEventSourceSuperSlurperTypeSuperSlurper:
		return true
	}
	return false
}

type SubscriptionDeleteResponseSourceMqEventSourceVectorize struct {
	// Type of source
	Type SubscriptionDeleteResponseSourceMqEventSourceVectorizeType `json:"type"`
	JSON subscriptionDeleteResponseSourceMqEventSourceVectorizeJSON `json:"-"`
}

// subscriptionDeleteResponseSourceMqEventSourceVectorizeJSON contains the JSON
// metadata for the struct [SubscriptionDeleteResponseSourceMqEventSourceVectorize]
type subscriptionDeleteResponseSourceMqEventSourceVectorizeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseSourceMqEventSourceVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseSourceMqEventSourceVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionDeleteResponseSourceMqEventSourceVectorize) implementsSubscriptionDeleteResponseSource() {
}

// Type of source
type SubscriptionDeleteResponseSourceMqEventSourceVectorizeType string

const (
	SubscriptionDeleteResponseSourceMqEventSourceVectorizeTypeVectorize SubscriptionDeleteResponseSourceMqEventSourceVectorizeType = "vectorize"
)

func (r SubscriptionDeleteResponseSourceMqEventSourceVectorizeType) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseSourceMqEventSourceVectorizeTypeVectorize:
		return true
	}
	return false
}

type SubscriptionDeleteResponseSourceMqEventSourceWorkersAIModel struct {
	// Name of the Workers AI model
	ModelName string `json:"model_name"`
	// Type of source
	Type SubscriptionDeleteResponseSourceMqEventSourceWorkersAIModelType `json:"type"`
	JSON subscriptionDeleteResponseSourceMqEventSourceWorkersAIModelJSON `json:"-"`
}

// subscriptionDeleteResponseSourceMqEventSourceWorkersAIModelJSON contains the
// JSON metadata for the struct
// [SubscriptionDeleteResponseSourceMqEventSourceWorkersAIModel]
type subscriptionDeleteResponseSourceMqEventSourceWorkersAIModelJSON struct {
	ModelName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseSourceMqEventSourceWorkersAIModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseSourceMqEventSourceWorkersAIModelJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionDeleteResponseSourceMqEventSourceWorkersAIModel) implementsSubscriptionDeleteResponseSource() {
}

// Type of source
type SubscriptionDeleteResponseSourceMqEventSourceWorkersAIModelType string

const (
	SubscriptionDeleteResponseSourceMqEventSourceWorkersAIModelTypeWorkersAIModel SubscriptionDeleteResponseSourceMqEventSourceWorkersAIModelType = "workersAi.model"
)

func (r SubscriptionDeleteResponseSourceMqEventSourceWorkersAIModelType) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseSourceMqEventSourceWorkersAIModelTypeWorkersAIModel:
		return true
	}
	return false
}

type SubscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorker struct {
	// Type of source
	Type SubscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorkerType `json:"type"`
	// Name of the worker
	WorkerName string                                                               `json:"worker_name"`
	JSON       subscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorkerJSON `json:"-"`
}

// subscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorkerJSON contains
// the JSON metadata for the struct
// [SubscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorker]
type subscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorkerJSON struct {
	Type        apijson.Field
	WorkerName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorkerJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorker) implementsSubscriptionDeleteResponseSource() {
}

// Type of source
type SubscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorkerType string

const (
	SubscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorkerTypeWorkersBuildsWorker SubscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorkerType = "workersBuilds.worker"
)

func (r SubscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorkerType) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseSourceMqEventSourceWorkersBuildsWorkerTypeWorkersBuildsWorker:
		return true
	}
	return false
}

type SubscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflow struct {
	// Type of source
	Type SubscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflowType `json:"type"`
	// Name of the workflow
	WorkflowName string                                                             `json:"workflow_name"`
	JSON         subscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflowJSON `json:"-"`
}

// subscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflowJSON contains the
// JSON metadata for the struct
// [SubscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflow]
type subscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflowJSON struct {
	Type         apijson.Field
	WorkflowName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflow) implementsSubscriptionDeleteResponseSource() {
}

// Type of source
type SubscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflowType string

const (
	SubscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflowTypeWorkflowsWorkflow SubscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflowType = "workflows.workflow"
)

func (r SubscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflowType) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseSourceMqEventSourceWorkflowsWorkflowTypeWorkflowsWorkflow:
		return true
	}
	return false
}

// Type of source
type SubscriptionDeleteResponseSourceType string

const (
	SubscriptionDeleteResponseSourceTypeImages              SubscriptionDeleteResponseSourceType = "images"
	SubscriptionDeleteResponseSourceTypeKV                  SubscriptionDeleteResponseSourceType = "kv"
	SubscriptionDeleteResponseSourceTypeR2                  SubscriptionDeleteResponseSourceType = "r2"
	SubscriptionDeleteResponseSourceTypeSuperSlurper        SubscriptionDeleteResponseSourceType = "superSlurper"
	SubscriptionDeleteResponseSourceTypeVectorize           SubscriptionDeleteResponseSourceType = "vectorize"
	SubscriptionDeleteResponseSourceTypeWorkersAIModel      SubscriptionDeleteResponseSourceType = "workersAi.model"
	SubscriptionDeleteResponseSourceTypeWorkersBuildsWorker SubscriptionDeleteResponseSourceType = "workersBuilds.worker"
	SubscriptionDeleteResponseSourceTypeWorkflowsWorkflow   SubscriptionDeleteResponseSourceType = "workflows.workflow"
)

func (r SubscriptionDeleteResponseSourceType) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseSourceTypeImages, SubscriptionDeleteResponseSourceTypeKV, SubscriptionDeleteResponseSourceTypeR2, SubscriptionDeleteResponseSourceTypeSuperSlurper, SubscriptionDeleteResponseSourceTypeVectorize, SubscriptionDeleteResponseSourceTypeWorkersAIModel, SubscriptionDeleteResponseSourceTypeWorkersBuildsWorker, SubscriptionDeleteResponseSourceTypeWorkflowsWorkflow:
		return true
	}
	return false
}

type SubscriptionGetResponse struct {
	// Unique identifier for the subscription
	ID string `json:"id,required"`
	// When the subscription was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Destination configuration for the subscription
	Destination SubscriptionGetResponseDestination `json:"destination,required"`
	// Whether the subscription is active
	Enabled bool `json:"enabled,required"`
	// List of event types this subscription handles
	Events []string `json:"events,required"`
	// When the subscription was last modified
	ModifiedAt time.Time `json:"modified_at,required" format:"date-time"`
	// Name of the subscription
	Name string `json:"name,required"`
	// Source configuration for the subscription
	Source SubscriptionGetResponseSource `json:"source,required"`
	JSON   subscriptionGetResponseJSON   `json:"-"`
}

// subscriptionGetResponseJSON contains the JSON metadata for the struct
// [SubscriptionGetResponse]
type subscriptionGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Destination apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Destination configuration for the subscription
type SubscriptionGetResponseDestination struct {
	// ID of the target queue
	QueueID string `json:"queue_id,required"`
	// Type of destination
	Type SubscriptionGetResponseDestinationType `json:"type,required"`
	JSON subscriptionGetResponseDestinationJSON `json:"-"`
}

// subscriptionGetResponseDestinationJSON contains the JSON metadata for the struct
// [SubscriptionGetResponseDestination]
type subscriptionGetResponseDestinationJSON struct {
	QueueID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseDestinationJSON) RawJSON() string {
	return r.raw
}

// Type of destination
type SubscriptionGetResponseDestinationType string

const (
	SubscriptionGetResponseDestinationTypeQueuesQueue SubscriptionGetResponseDestinationType = "queues.queue"
)

func (r SubscriptionGetResponseDestinationType) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseDestinationTypeQueuesQueue:
		return true
	}
	return false
}

// Source configuration for the subscription
type SubscriptionGetResponseSource struct {
	// Name of the Workers AI model
	ModelName string `json:"model_name"`
	// Type of source
	Type SubscriptionGetResponseSourceType `json:"type"`
	// Name of the worker
	WorkerName string `json:"worker_name"`
	// Name of the workflow
	WorkflowName string                            `json:"workflow_name"`
	JSON         subscriptionGetResponseSourceJSON `json:"-"`
	union        SubscriptionGetResponseSourceUnion
}

// subscriptionGetResponseSourceJSON contains the JSON metadata for the struct
// [SubscriptionGetResponseSource]
type subscriptionGetResponseSourceJSON struct {
	ModelName    apijson.Field
	Type         apijson.Field
	WorkerName   apijson.Field
	WorkflowName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r subscriptionGetResponseSourceJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionGetResponseSource) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionGetResponseSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SubscriptionGetResponseSourceUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionGetResponseSourceMqEventSourceImages],
// [SubscriptionGetResponseSourceMqEventSourceKV],
// [SubscriptionGetResponseSourceMqEventSourceR2],
// [SubscriptionGetResponseSourceMqEventSourceSuperSlurper],
// [SubscriptionGetResponseSourceMqEventSourceVectorize],
// [SubscriptionGetResponseSourceMqEventSourceWorkersAIModel],
// [SubscriptionGetResponseSourceMqEventSourceWorkersBuildsWorker],
// [SubscriptionGetResponseSourceMqEventSourceWorkflowsWorkflow].
func (r SubscriptionGetResponseSource) AsUnion() SubscriptionGetResponseSourceUnion {
	return r.union
}

// Source configuration for the subscription
//
// Union satisfied by [SubscriptionGetResponseSourceMqEventSourceImages],
// [SubscriptionGetResponseSourceMqEventSourceKV],
// [SubscriptionGetResponseSourceMqEventSourceR2],
// [SubscriptionGetResponseSourceMqEventSourceSuperSlurper],
// [SubscriptionGetResponseSourceMqEventSourceVectorize],
// [SubscriptionGetResponseSourceMqEventSourceWorkersAIModel],
// [SubscriptionGetResponseSourceMqEventSourceWorkersBuildsWorker] or
// [SubscriptionGetResponseSourceMqEventSourceWorkflowsWorkflow].
type SubscriptionGetResponseSourceUnion interface {
	implementsSubscriptionGetResponseSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionGetResponseSourceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionGetResponseSourceMqEventSourceImages{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionGetResponseSourceMqEventSourceKV{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionGetResponseSourceMqEventSourceR2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionGetResponseSourceMqEventSourceSuperSlurper{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionGetResponseSourceMqEventSourceVectorize{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionGetResponseSourceMqEventSourceWorkersAIModel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionGetResponseSourceMqEventSourceWorkersBuildsWorker{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionGetResponseSourceMqEventSourceWorkflowsWorkflow{}),
		},
	)
}

type SubscriptionGetResponseSourceMqEventSourceImages struct {
	// Type of source
	Type SubscriptionGetResponseSourceMqEventSourceImagesType `json:"type"`
	JSON subscriptionGetResponseSourceMqEventSourceImagesJSON `json:"-"`
}

// subscriptionGetResponseSourceMqEventSourceImagesJSON contains the JSON metadata
// for the struct [SubscriptionGetResponseSourceMqEventSourceImages]
type subscriptionGetResponseSourceMqEventSourceImagesJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseSourceMqEventSourceImages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseSourceMqEventSourceImagesJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionGetResponseSourceMqEventSourceImages) implementsSubscriptionGetResponseSource() {}

// Type of source
type SubscriptionGetResponseSourceMqEventSourceImagesType string

const (
	SubscriptionGetResponseSourceMqEventSourceImagesTypeImages SubscriptionGetResponseSourceMqEventSourceImagesType = "images"
)

func (r SubscriptionGetResponseSourceMqEventSourceImagesType) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseSourceMqEventSourceImagesTypeImages:
		return true
	}
	return false
}

type SubscriptionGetResponseSourceMqEventSourceKV struct {
	// Type of source
	Type SubscriptionGetResponseSourceMqEventSourceKVType `json:"type"`
	JSON subscriptionGetResponseSourceMqEventSourceKVJSON `json:"-"`
}

// subscriptionGetResponseSourceMqEventSourceKVJSON contains the JSON metadata for
// the struct [SubscriptionGetResponseSourceMqEventSourceKV]
type subscriptionGetResponseSourceMqEventSourceKVJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseSourceMqEventSourceKV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseSourceMqEventSourceKVJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionGetResponseSourceMqEventSourceKV) implementsSubscriptionGetResponseSource() {}

// Type of source
type SubscriptionGetResponseSourceMqEventSourceKVType string

const (
	SubscriptionGetResponseSourceMqEventSourceKVTypeKV SubscriptionGetResponseSourceMqEventSourceKVType = "kv"
)

func (r SubscriptionGetResponseSourceMqEventSourceKVType) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseSourceMqEventSourceKVTypeKV:
		return true
	}
	return false
}

type SubscriptionGetResponseSourceMqEventSourceR2 struct {
	// Type of source
	Type SubscriptionGetResponseSourceMqEventSourceR2Type `json:"type"`
	JSON subscriptionGetResponseSourceMqEventSourceR2JSON `json:"-"`
}

// subscriptionGetResponseSourceMqEventSourceR2JSON contains the JSON metadata for
// the struct [SubscriptionGetResponseSourceMqEventSourceR2]
type subscriptionGetResponseSourceMqEventSourceR2JSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseSourceMqEventSourceR2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseSourceMqEventSourceR2JSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionGetResponseSourceMqEventSourceR2) implementsSubscriptionGetResponseSource() {}

// Type of source
type SubscriptionGetResponseSourceMqEventSourceR2Type string

const (
	SubscriptionGetResponseSourceMqEventSourceR2TypeR2 SubscriptionGetResponseSourceMqEventSourceR2Type = "r2"
)

func (r SubscriptionGetResponseSourceMqEventSourceR2Type) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseSourceMqEventSourceR2TypeR2:
		return true
	}
	return false
}

type SubscriptionGetResponseSourceMqEventSourceSuperSlurper struct {
	// Type of source
	Type SubscriptionGetResponseSourceMqEventSourceSuperSlurperType `json:"type"`
	JSON subscriptionGetResponseSourceMqEventSourceSuperSlurperJSON `json:"-"`
}

// subscriptionGetResponseSourceMqEventSourceSuperSlurperJSON contains the JSON
// metadata for the struct [SubscriptionGetResponseSourceMqEventSourceSuperSlurper]
type subscriptionGetResponseSourceMqEventSourceSuperSlurperJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseSourceMqEventSourceSuperSlurper) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseSourceMqEventSourceSuperSlurperJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionGetResponseSourceMqEventSourceSuperSlurper) implementsSubscriptionGetResponseSource() {
}

// Type of source
type SubscriptionGetResponseSourceMqEventSourceSuperSlurperType string

const (
	SubscriptionGetResponseSourceMqEventSourceSuperSlurperTypeSuperSlurper SubscriptionGetResponseSourceMqEventSourceSuperSlurperType = "superSlurper"
)

func (r SubscriptionGetResponseSourceMqEventSourceSuperSlurperType) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseSourceMqEventSourceSuperSlurperTypeSuperSlurper:
		return true
	}
	return false
}

type SubscriptionGetResponseSourceMqEventSourceVectorize struct {
	// Type of source
	Type SubscriptionGetResponseSourceMqEventSourceVectorizeType `json:"type"`
	JSON subscriptionGetResponseSourceMqEventSourceVectorizeJSON `json:"-"`
}

// subscriptionGetResponseSourceMqEventSourceVectorizeJSON contains the JSON
// metadata for the struct [SubscriptionGetResponseSourceMqEventSourceVectorize]
type subscriptionGetResponseSourceMqEventSourceVectorizeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseSourceMqEventSourceVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseSourceMqEventSourceVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionGetResponseSourceMqEventSourceVectorize) implementsSubscriptionGetResponseSource() {
}

// Type of source
type SubscriptionGetResponseSourceMqEventSourceVectorizeType string

const (
	SubscriptionGetResponseSourceMqEventSourceVectorizeTypeVectorize SubscriptionGetResponseSourceMqEventSourceVectorizeType = "vectorize"
)

func (r SubscriptionGetResponseSourceMqEventSourceVectorizeType) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseSourceMqEventSourceVectorizeTypeVectorize:
		return true
	}
	return false
}

type SubscriptionGetResponseSourceMqEventSourceWorkersAIModel struct {
	// Name of the Workers AI model
	ModelName string `json:"model_name"`
	// Type of source
	Type SubscriptionGetResponseSourceMqEventSourceWorkersAIModelType `json:"type"`
	JSON subscriptionGetResponseSourceMqEventSourceWorkersAIModelJSON `json:"-"`
}

// subscriptionGetResponseSourceMqEventSourceWorkersAIModelJSON contains the JSON
// metadata for the struct
// [SubscriptionGetResponseSourceMqEventSourceWorkersAIModel]
type subscriptionGetResponseSourceMqEventSourceWorkersAIModelJSON struct {
	ModelName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseSourceMqEventSourceWorkersAIModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseSourceMqEventSourceWorkersAIModelJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionGetResponseSourceMqEventSourceWorkersAIModel) implementsSubscriptionGetResponseSource() {
}

// Type of source
type SubscriptionGetResponseSourceMqEventSourceWorkersAIModelType string

const (
	SubscriptionGetResponseSourceMqEventSourceWorkersAIModelTypeWorkersAIModel SubscriptionGetResponseSourceMqEventSourceWorkersAIModelType = "workersAi.model"
)

func (r SubscriptionGetResponseSourceMqEventSourceWorkersAIModelType) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseSourceMqEventSourceWorkersAIModelTypeWorkersAIModel:
		return true
	}
	return false
}

type SubscriptionGetResponseSourceMqEventSourceWorkersBuildsWorker struct {
	// Type of source
	Type SubscriptionGetResponseSourceMqEventSourceWorkersBuildsWorkerType `json:"type"`
	// Name of the worker
	WorkerName string                                                            `json:"worker_name"`
	JSON       subscriptionGetResponseSourceMqEventSourceWorkersBuildsWorkerJSON `json:"-"`
}

// subscriptionGetResponseSourceMqEventSourceWorkersBuildsWorkerJSON contains the
// JSON metadata for the struct
// [SubscriptionGetResponseSourceMqEventSourceWorkersBuildsWorker]
type subscriptionGetResponseSourceMqEventSourceWorkersBuildsWorkerJSON struct {
	Type        apijson.Field
	WorkerName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseSourceMqEventSourceWorkersBuildsWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseSourceMqEventSourceWorkersBuildsWorkerJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionGetResponseSourceMqEventSourceWorkersBuildsWorker) implementsSubscriptionGetResponseSource() {
}

// Type of source
type SubscriptionGetResponseSourceMqEventSourceWorkersBuildsWorkerType string

const (
	SubscriptionGetResponseSourceMqEventSourceWorkersBuildsWorkerTypeWorkersBuildsWorker SubscriptionGetResponseSourceMqEventSourceWorkersBuildsWorkerType = "workersBuilds.worker"
)

func (r SubscriptionGetResponseSourceMqEventSourceWorkersBuildsWorkerType) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseSourceMqEventSourceWorkersBuildsWorkerTypeWorkersBuildsWorker:
		return true
	}
	return false
}

type SubscriptionGetResponseSourceMqEventSourceWorkflowsWorkflow struct {
	// Type of source
	Type SubscriptionGetResponseSourceMqEventSourceWorkflowsWorkflowType `json:"type"`
	// Name of the workflow
	WorkflowName string                                                          `json:"workflow_name"`
	JSON         subscriptionGetResponseSourceMqEventSourceWorkflowsWorkflowJSON `json:"-"`
}

// subscriptionGetResponseSourceMqEventSourceWorkflowsWorkflowJSON contains the
// JSON metadata for the struct
// [SubscriptionGetResponseSourceMqEventSourceWorkflowsWorkflow]
type subscriptionGetResponseSourceMqEventSourceWorkflowsWorkflowJSON struct {
	Type         apijson.Field
	WorkflowName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionGetResponseSourceMqEventSourceWorkflowsWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseSourceMqEventSourceWorkflowsWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionGetResponseSourceMqEventSourceWorkflowsWorkflow) implementsSubscriptionGetResponseSource() {
}

// Type of source
type SubscriptionGetResponseSourceMqEventSourceWorkflowsWorkflowType string

const (
	SubscriptionGetResponseSourceMqEventSourceWorkflowsWorkflowTypeWorkflowsWorkflow SubscriptionGetResponseSourceMqEventSourceWorkflowsWorkflowType = "workflows.workflow"
)

func (r SubscriptionGetResponseSourceMqEventSourceWorkflowsWorkflowType) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseSourceMqEventSourceWorkflowsWorkflowTypeWorkflowsWorkflow:
		return true
	}
	return false
}

// Type of source
type SubscriptionGetResponseSourceType string

const (
	SubscriptionGetResponseSourceTypeImages              SubscriptionGetResponseSourceType = "images"
	SubscriptionGetResponseSourceTypeKV                  SubscriptionGetResponseSourceType = "kv"
	SubscriptionGetResponseSourceTypeR2                  SubscriptionGetResponseSourceType = "r2"
	SubscriptionGetResponseSourceTypeSuperSlurper        SubscriptionGetResponseSourceType = "superSlurper"
	SubscriptionGetResponseSourceTypeVectorize           SubscriptionGetResponseSourceType = "vectorize"
	SubscriptionGetResponseSourceTypeWorkersAIModel      SubscriptionGetResponseSourceType = "workersAi.model"
	SubscriptionGetResponseSourceTypeWorkersBuildsWorker SubscriptionGetResponseSourceType = "workersBuilds.worker"
	SubscriptionGetResponseSourceTypeWorkflowsWorkflow   SubscriptionGetResponseSourceType = "workflows.workflow"
)

func (r SubscriptionGetResponseSourceType) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseSourceTypeImages, SubscriptionGetResponseSourceTypeKV, SubscriptionGetResponseSourceTypeR2, SubscriptionGetResponseSourceTypeSuperSlurper, SubscriptionGetResponseSourceTypeVectorize, SubscriptionGetResponseSourceTypeWorkersAIModel, SubscriptionGetResponseSourceTypeWorkersBuildsWorker, SubscriptionGetResponseSourceTypeWorkflowsWorkflow:
		return true
	}
	return false
}

type SubscriptionNewParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Destination configuration for the subscription
	Destination param.Field[SubscriptionNewParamsDestination] `json:"destination"`
	// Whether the subscription is active
	Enabled param.Field[bool] `json:"enabled"`
	// List of event types this subscription handles
	Events param.Field[[]string] `json:"events"`
	// Name of the subscription
	Name param.Field[string] `json:"name"`
	// Source configuration for the subscription
	Source param.Field[SubscriptionNewParamsSourceUnion] `json:"source"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Destination configuration for the subscription
type SubscriptionNewParamsDestination struct {
	// ID of the target queue
	QueueID param.Field[string] `json:"queue_id,required"`
	// Type of destination
	Type param.Field[SubscriptionNewParamsDestinationType] `json:"type,required"`
}

func (r SubscriptionNewParamsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of destination
type SubscriptionNewParamsDestinationType string

const (
	SubscriptionNewParamsDestinationTypeQueuesQueue SubscriptionNewParamsDestinationType = "queues.queue"
)

func (r SubscriptionNewParamsDestinationType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsDestinationTypeQueuesQueue:
		return true
	}
	return false
}

// Source configuration for the subscription
type SubscriptionNewParamsSource struct {
	// Name of the Workers AI model
	ModelName param.Field[string] `json:"model_name"`
	// Type of source
	Type param.Field[SubscriptionNewParamsSourceType] `json:"type"`
	// Name of the worker
	WorkerName param.Field[string] `json:"worker_name"`
	// Name of the workflow
	WorkflowName param.Field[string] `json:"workflow_name"`
}

func (r SubscriptionNewParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsSource) implementsSubscriptionNewParamsSourceUnion() {}

// Source configuration for the subscription
//
// Satisfied by [queues.SubscriptionNewParamsSourceMqEventSourceImages],
// [queues.SubscriptionNewParamsSourceMqEventSourceKV],
// [queues.SubscriptionNewParamsSourceMqEventSourceR2],
// [queues.SubscriptionNewParamsSourceMqEventSourceSuperSlurper],
// [queues.SubscriptionNewParamsSourceMqEventSourceVectorize],
// [queues.SubscriptionNewParamsSourceMqEventSourceWorkersAIModel],
// [queues.SubscriptionNewParamsSourceMqEventSourceWorkersBuildsWorker],
// [queues.SubscriptionNewParamsSourceMqEventSourceWorkflowsWorkflow],
// [SubscriptionNewParamsSource].
type SubscriptionNewParamsSourceUnion interface {
	implementsSubscriptionNewParamsSourceUnion()
}

type SubscriptionNewParamsSourceMqEventSourceImages struct {
	// Type of source
	Type param.Field[SubscriptionNewParamsSourceMqEventSourceImagesType] `json:"type"`
}

func (r SubscriptionNewParamsSourceMqEventSourceImages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsSourceMqEventSourceImages) implementsSubscriptionNewParamsSourceUnion() {
}

// Type of source
type SubscriptionNewParamsSourceMqEventSourceImagesType string

const (
	SubscriptionNewParamsSourceMqEventSourceImagesTypeImages SubscriptionNewParamsSourceMqEventSourceImagesType = "images"
)

func (r SubscriptionNewParamsSourceMqEventSourceImagesType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsSourceMqEventSourceImagesTypeImages:
		return true
	}
	return false
}

type SubscriptionNewParamsSourceMqEventSourceKV struct {
	// Type of source
	Type param.Field[SubscriptionNewParamsSourceMqEventSourceKVType] `json:"type"`
}

func (r SubscriptionNewParamsSourceMqEventSourceKV) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsSourceMqEventSourceKV) implementsSubscriptionNewParamsSourceUnion() {}

// Type of source
type SubscriptionNewParamsSourceMqEventSourceKVType string

const (
	SubscriptionNewParamsSourceMqEventSourceKVTypeKV SubscriptionNewParamsSourceMqEventSourceKVType = "kv"
)

func (r SubscriptionNewParamsSourceMqEventSourceKVType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsSourceMqEventSourceKVTypeKV:
		return true
	}
	return false
}

type SubscriptionNewParamsSourceMqEventSourceR2 struct {
	// Type of source
	Type param.Field[SubscriptionNewParamsSourceMqEventSourceR2Type] `json:"type"`
}

func (r SubscriptionNewParamsSourceMqEventSourceR2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsSourceMqEventSourceR2) implementsSubscriptionNewParamsSourceUnion() {}

// Type of source
type SubscriptionNewParamsSourceMqEventSourceR2Type string

const (
	SubscriptionNewParamsSourceMqEventSourceR2TypeR2 SubscriptionNewParamsSourceMqEventSourceR2Type = "r2"
)

func (r SubscriptionNewParamsSourceMqEventSourceR2Type) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsSourceMqEventSourceR2TypeR2:
		return true
	}
	return false
}

type SubscriptionNewParamsSourceMqEventSourceSuperSlurper struct {
	// Type of source
	Type param.Field[SubscriptionNewParamsSourceMqEventSourceSuperSlurperType] `json:"type"`
}

func (r SubscriptionNewParamsSourceMqEventSourceSuperSlurper) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsSourceMqEventSourceSuperSlurper) implementsSubscriptionNewParamsSourceUnion() {
}

// Type of source
type SubscriptionNewParamsSourceMqEventSourceSuperSlurperType string

const (
	SubscriptionNewParamsSourceMqEventSourceSuperSlurperTypeSuperSlurper SubscriptionNewParamsSourceMqEventSourceSuperSlurperType = "superSlurper"
)

func (r SubscriptionNewParamsSourceMqEventSourceSuperSlurperType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsSourceMqEventSourceSuperSlurperTypeSuperSlurper:
		return true
	}
	return false
}

type SubscriptionNewParamsSourceMqEventSourceVectorize struct {
	// Type of source
	Type param.Field[SubscriptionNewParamsSourceMqEventSourceVectorizeType] `json:"type"`
}

func (r SubscriptionNewParamsSourceMqEventSourceVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsSourceMqEventSourceVectorize) implementsSubscriptionNewParamsSourceUnion() {
}

// Type of source
type SubscriptionNewParamsSourceMqEventSourceVectorizeType string

const (
	SubscriptionNewParamsSourceMqEventSourceVectorizeTypeVectorize SubscriptionNewParamsSourceMqEventSourceVectorizeType = "vectorize"
)

func (r SubscriptionNewParamsSourceMqEventSourceVectorizeType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsSourceMqEventSourceVectorizeTypeVectorize:
		return true
	}
	return false
}

type SubscriptionNewParamsSourceMqEventSourceWorkersAIModel struct {
	// Name of the Workers AI model
	ModelName param.Field[string] `json:"model_name"`
	// Type of source
	Type param.Field[SubscriptionNewParamsSourceMqEventSourceWorkersAIModelType] `json:"type"`
}

func (r SubscriptionNewParamsSourceMqEventSourceWorkersAIModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsSourceMqEventSourceWorkersAIModel) implementsSubscriptionNewParamsSourceUnion() {
}

// Type of source
type SubscriptionNewParamsSourceMqEventSourceWorkersAIModelType string

const (
	SubscriptionNewParamsSourceMqEventSourceWorkersAIModelTypeWorkersAIModel SubscriptionNewParamsSourceMqEventSourceWorkersAIModelType = "workersAi.model"
)

func (r SubscriptionNewParamsSourceMqEventSourceWorkersAIModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsSourceMqEventSourceWorkersAIModelTypeWorkersAIModel:
		return true
	}
	return false
}

type SubscriptionNewParamsSourceMqEventSourceWorkersBuildsWorker struct {
	// Type of source
	Type param.Field[SubscriptionNewParamsSourceMqEventSourceWorkersBuildsWorkerType] `json:"type"`
	// Name of the worker
	WorkerName param.Field[string] `json:"worker_name"`
}

func (r SubscriptionNewParamsSourceMqEventSourceWorkersBuildsWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsSourceMqEventSourceWorkersBuildsWorker) implementsSubscriptionNewParamsSourceUnion() {
}

// Type of source
type SubscriptionNewParamsSourceMqEventSourceWorkersBuildsWorkerType string

const (
	SubscriptionNewParamsSourceMqEventSourceWorkersBuildsWorkerTypeWorkersBuildsWorker SubscriptionNewParamsSourceMqEventSourceWorkersBuildsWorkerType = "workersBuilds.worker"
)

func (r SubscriptionNewParamsSourceMqEventSourceWorkersBuildsWorkerType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsSourceMqEventSourceWorkersBuildsWorkerTypeWorkersBuildsWorker:
		return true
	}
	return false
}

type SubscriptionNewParamsSourceMqEventSourceWorkflowsWorkflow struct {
	// Type of source
	Type param.Field[SubscriptionNewParamsSourceMqEventSourceWorkflowsWorkflowType] `json:"type"`
	// Name of the workflow
	WorkflowName param.Field[string] `json:"workflow_name"`
}

func (r SubscriptionNewParamsSourceMqEventSourceWorkflowsWorkflow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsSourceMqEventSourceWorkflowsWorkflow) implementsSubscriptionNewParamsSourceUnion() {
}

// Type of source
type SubscriptionNewParamsSourceMqEventSourceWorkflowsWorkflowType string

const (
	SubscriptionNewParamsSourceMqEventSourceWorkflowsWorkflowTypeWorkflowsWorkflow SubscriptionNewParamsSourceMqEventSourceWorkflowsWorkflowType = "workflows.workflow"
)

func (r SubscriptionNewParamsSourceMqEventSourceWorkflowsWorkflowType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsSourceMqEventSourceWorkflowsWorkflowTypeWorkflowsWorkflow:
		return true
	}
	return false
}

// Type of source
type SubscriptionNewParamsSourceType string

const (
	SubscriptionNewParamsSourceTypeImages              SubscriptionNewParamsSourceType = "images"
	SubscriptionNewParamsSourceTypeKV                  SubscriptionNewParamsSourceType = "kv"
	SubscriptionNewParamsSourceTypeR2                  SubscriptionNewParamsSourceType = "r2"
	SubscriptionNewParamsSourceTypeSuperSlurper        SubscriptionNewParamsSourceType = "superSlurper"
	SubscriptionNewParamsSourceTypeVectorize           SubscriptionNewParamsSourceType = "vectorize"
	SubscriptionNewParamsSourceTypeWorkersAIModel      SubscriptionNewParamsSourceType = "workersAi.model"
	SubscriptionNewParamsSourceTypeWorkersBuildsWorker SubscriptionNewParamsSourceType = "workersBuilds.worker"
	SubscriptionNewParamsSourceTypeWorkflowsWorkflow   SubscriptionNewParamsSourceType = "workflows.workflow"
)

func (r SubscriptionNewParamsSourceType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsSourceTypeImages, SubscriptionNewParamsSourceTypeKV, SubscriptionNewParamsSourceTypeR2, SubscriptionNewParamsSourceTypeSuperSlurper, SubscriptionNewParamsSourceTypeVectorize, SubscriptionNewParamsSourceTypeWorkersAIModel, SubscriptionNewParamsSourceTypeWorkersBuildsWorker, SubscriptionNewParamsSourceTypeWorkflowsWorkflow:
		return true
	}
	return false
}

type SubscriptionNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors"`
	Messages []string                `json:"messages"`
	Result   SubscriptionNewResponse `json:"result"`
	// Indicates if the API call was successful or not.
	Success SubscriptionNewResponseEnvelopeSuccess `json:"success"`
	JSON    subscriptionNewResponseEnvelopeJSON    `json:"-"`
}

// subscriptionNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionNewResponseEnvelope]
type subscriptionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type SubscriptionNewResponseEnvelopeSuccess bool

const (
	SubscriptionNewResponseEnvelopeSuccessTrue SubscriptionNewResponseEnvelopeSuccess = true
)

func (r SubscriptionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubscriptionUpdateParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Destination configuration for the subscription
	Destination param.Field[SubscriptionUpdateParamsDestination] `json:"destination"`
	// Whether the subscription is active
	Enabled param.Field[bool] `json:"enabled"`
	// List of event types this subscription handles
	Events param.Field[[]string] `json:"events"`
	// Name of the subscription
	Name param.Field[string] `json:"name"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Destination configuration for the subscription
type SubscriptionUpdateParamsDestination struct {
	// ID of the target queue
	QueueID param.Field[string] `json:"queue_id,required"`
	// Type of destination
	Type param.Field[SubscriptionUpdateParamsDestinationType] `json:"type,required"`
}

func (r SubscriptionUpdateParamsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of destination
type SubscriptionUpdateParamsDestinationType string

const (
	SubscriptionUpdateParamsDestinationTypeQueuesQueue SubscriptionUpdateParamsDestinationType = "queues.queue"
)

func (r SubscriptionUpdateParamsDestinationType) IsKnown() bool {
	switch r {
	case SubscriptionUpdateParamsDestinationTypeQueuesQueue:
		return true
	}
	return false
}

type SubscriptionUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors"`
	Messages []string                   `json:"messages"`
	Result   SubscriptionUpdateResponse `json:"result"`
	// Indicates if the API call was successful or not.
	Success SubscriptionUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    subscriptionUpdateResponseEnvelopeJSON    `json:"-"`
}

// subscriptionUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionUpdateResponseEnvelope]
type subscriptionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type SubscriptionUpdateResponseEnvelopeSuccess bool

const (
	SubscriptionUpdateResponseEnvelopeSuccessTrue SubscriptionUpdateResponseEnvelopeSuccess = true
)

func (r SubscriptionUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubscriptionListParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Sort direction
	Direction param.Field[SubscriptionListParamsDirection] `query:"direction"`
	// Field to sort by
	Order param.Field[SubscriptionListParamsOrder] `query:"order"`
	// Page number for pagination
	Page param.Field[int64] `query:"page"`
	// Number of items per page
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [SubscriptionListParams]'s query parameters as `url.Values`.
func (r SubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction
type SubscriptionListParamsDirection string

const (
	SubscriptionListParamsDirectionAsc  SubscriptionListParamsDirection = "asc"
	SubscriptionListParamsDirectionDesc SubscriptionListParamsDirection = "desc"
)

func (r SubscriptionListParamsDirection) IsKnown() bool {
	switch r {
	case SubscriptionListParamsDirectionAsc, SubscriptionListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to sort by
type SubscriptionListParamsOrder string

const (
	SubscriptionListParamsOrderCreatedAt SubscriptionListParamsOrder = "created_at"
	SubscriptionListParamsOrderName      SubscriptionListParamsOrder = "name"
	SubscriptionListParamsOrderEnabled   SubscriptionListParamsOrder = "enabled"
	SubscriptionListParamsOrderSource    SubscriptionListParamsOrder = "source"
)

func (r SubscriptionListParamsOrder) IsKnown() bool {
	switch r {
	case SubscriptionListParamsOrderCreatedAt, SubscriptionListParamsOrderName, SubscriptionListParamsOrderEnabled, SubscriptionListParamsOrderSource:
		return true
	}
	return false
}

type SubscriptionDeleteParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SubscriptionDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors"`
	Messages []string                   `json:"messages"`
	Result   SubscriptionDeleteResponse `json:"result"`
	// Indicates if the API call was successful or not.
	Success SubscriptionDeleteResponseEnvelopeSuccess `json:"success"`
	JSON    subscriptionDeleteResponseEnvelopeJSON    `json:"-"`
}

// subscriptionDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionDeleteResponseEnvelope]
type subscriptionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type SubscriptionDeleteResponseEnvelopeSuccess bool

const (
	SubscriptionDeleteResponseEnvelopeSuccessTrue SubscriptionDeleteResponseEnvelopeSuccess = true
)

func (r SubscriptionDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubscriptionGetParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SubscriptionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors"`
	Messages []string                `json:"messages"`
	Result   SubscriptionGetResponse `json:"result"`
	// Indicates if the API call was successful or not.
	Success SubscriptionGetResponseEnvelopeSuccess `json:"success"`
	JSON    subscriptionGetResponseEnvelopeJSON    `json:"-"`
}

// subscriptionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionGetResponseEnvelope]
type subscriptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type SubscriptionGetResponseEnvelopeSuccess bool

const (
	SubscriptionGetResponseEnvelopeSuccessTrue SubscriptionGetResponseEnvelopeSuccess = true
)

func (r SubscriptionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
