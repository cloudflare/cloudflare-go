// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// ObservabilityAgentTraceService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilityAgentTraceService] method instead.
type ObservabilityAgentTraceService struct {
	Options []option.RequestOption
}

// NewObservabilityAgentTraceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObservabilityAgentTraceService(opts ...option.RequestOption) (r *ObservabilityAgentTraceService) {
	r = &ObservabilityAgentTraceService{}
	r.Options = opts
	return
}

// Lists Traces with a named Agent invocation in the window, ordered by
// lastActivityMs, newest first, including work without a Conversation ID. The
// window defaults to the last 24 hours. Summaries include the Agent names and
// Conversation IDs seen, authorized Trace duration, chat input/output/total
// tokens, models, providers, LLM/tool call counts and span-error counts within the
// effective window.
func (r *ObservabilityAgentTraceService) List(ctx context.Context, params ObservabilityAgentTraceListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObservabilityAgentTraceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/agents/traces", params.AccountID)
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

// Lists Traces with a named Agent invocation in the window, ordered by
// lastActivityMs, newest first, including work without a Conversation ID. The
// window defaults to the last 24 hours. Summaries include the Agent names and
// Conversation IDs seen, authorized Trace duration, chat input/output/total
// tokens, models, providers, LLM/tool call counts and span-error counts within the
// effective window.
func (r *ObservabilityAgentTraceService) ListAutoPaging(ctx context.Context, params ObservabilityAgentTraceListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObservabilityAgentTraceListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Returns a page of authorized canonical otel events for a Trace containing named
// Agent activity, using the existing events-query flow. Searches the full 7-day
// retention unless from or to narrow it. Events are ordered by event timestamp and
// ID, newest first. Continue with result_info.cursors.after; one page is not a
// complete span tree.
func (r *ObservabilityAgentTraceService) Events(ctx context.Context, traceID string, params ObservabilityAgentTraceEventsParams, opts ...option.RequestOption) (res *[]ObservabilityAgentTraceEventsResponse, err error) {
	var env ObservabilityAgentTraceEventsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if traceID == "" {
		err = errors.New("missing required traceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/agents/traces/%s/events", params.AccountID, traceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the summary of one Trace containing named Agent activity, with the same
// fields as the Trace list. Searches the full 7-day retention unless from or to
// narrow it. List the Trace's events with /agents/traces/{traceId}/events.
func (r *ObservabilityAgentTraceService) Get(ctx context.Context, traceID string, params ObservabilityAgentTraceGetParams, opts ...option.RequestOption) (res *ObservabilityAgentTraceGetResponse, err error) {
	var env ObservabilityAgentTraceGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if traceID == "" {
		err = errors.New("missing required traceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/agents/traces/%s", params.AccountID, traceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ObservabilityAgentTraceListResponse struct {
	// Sorted distinct gen_ai.agent.name values of authorized Agent invocations in the
	// Trace within the effective window. Includes every Agent in the Trace, not only
	// those matching the agent filter.
	Agents []string `json:"agents" api:"required"`
	// Sorted distinct gen_ai.conversation.id values of authorized Agent invocations in
	// the Trace within the effective window. Empty when none are recorded.
	ConversationIDs []string `json:"conversationIds" api:"required"`
	// Recorded spans with a nonempty error or ERROR status. Counts spans, not distinct
	// messages or incidents.
	ErrorCount int64 `json:"errorCount" api:"required"`
	// Input tokens reported by chat spans in the authorized Trace window; null when
	// none report input usage. Informational, not billing data.
	InputTokens float64 `json:"inputTokens" api:"required,nullable"`
	// Latest start of an authorized Agent invocation in the Trace within the effective
	// window, as Unix epoch milliseconds. The Trace list is ordered by this value,
	// newest first.
	LastActivityMs int64 `json:"lastActivityMs" api:"required"`
	// Recorded chat spans in the authorized Trace window.
	LlmCallCount int64 `json:"llmCallCount" api:"required"`
	// Recorded chat spans with a nonempty error or ERROR status.
	LlmErrorCount int64 `json:"llmErrorCount" api:"required"`
	// Sorted distinct response models, falling back to request models, from chat spans
	// in the authorized Trace window.
	Models []string `json:"models" api:"required"`
	// Output tokens reported by chat spans in the authorized Trace window; null when
	// none report output usage. Informational, not billing data.
	OutputTokens float64 `json:"outputTokens" api:"required,nullable"`
	// Sorted distinct GenAI providers from chat spans in the authorized Trace window.
	Providers []string `json:"providers" api:"required"`
	// Name of the root span that initiated the trace.
	RootSpanName string `json:"rootSpanName" api:"required"`
	// Logical transaction name for the root span.
	RootTransactionName string `json:"rootTransactionName" api:"required"`
	// Authorized Worker services represented in the trace.
	Services []string `json:"services" api:"required"`
	// Total number of spans in the trace.
	Spans float64 `json:"spans" api:"required"`
	// Recorded execute_tool spans in the authorized Trace window. Custom-named
	// approval spans are not counted.
	ToolCallCount int64 `json:"toolCallCount" api:"required"`
	// Recorded execute_tool spans with a nonempty error or ERROR status.
	ToolErrorCount int64 `json:"toolErrorCount" api:"required"`
	// Input plus output tokens; null if either direction is unreported. Does not add
	// cache or reasoning breakdowns.
	TotalTokens float64 `json:"totalTokens" api:"required,nullable"`
	// Total duration of the trace in milliseconds.
	TraceDurationMs float64 `json:"traceDurationMs" api:"required"`
	// Trace end time as a Unix epoch in milliseconds.
	TraceEndMs float64 `json:"traceEndMs" api:"required"`
	// Unique identifier for the distributed trace.
	TraceID string `json:"traceId" api:"required"`
	// Trace start time as a Unix epoch in milliseconds.
	TraceStartMs float64 `json:"traceStartMs" api:"required"`
	// Error messages encountered during the trace, if any.
	Errors []string                                `json:"errors"`
	JSON   observabilityAgentTraceListResponseJSON `json:"-"`
}

// observabilityAgentTraceListResponseJSON contains the JSON metadata for the
// struct [ObservabilityAgentTraceListResponse]
type observabilityAgentTraceListResponseJSON struct {
	Agents              apijson.Field
	ConversationIDs     apijson.Field
	ErrorCount          apijson.Field
	InputTokens         apijson.Field
	LastActivityMs      apijson.Field
	LlmCallCount        apijson.Field
	LlmErrorCount       apijson.Field
	Models              apijson.Field
	OutputTokens        apijson.Field
	Providers           apijson.Field
	RootSpanName        apijson.Field
	RootTransactionName apijson.Field
	Services            apijson.Field
	Spans               apijson.Field
	ToolCallCount       apijson.Field
	ToolErrorCount      apijson.Field
	TotalTokens         apijson.Field
	TraceDurationMs     apijson.Field
	TraceEndMs          apijson.Field
	TraceID             apijson.Field
	TraceStartMs        apijson.Field
	Errors              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ObservabilityAgentTraceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceListResponseJSON) RawJSON() string {
	return r.raw
}

// A single telemetry event representing a log line, span, or metric data point
// emitted by a Worker.
type ObservabilityAgentTraceEventsResponse struct {
	// Structured metadata extracted from the event. These fields are indexed and
	// available for filtering and aggregation.
	Metadata ObservabilityAgentTraceEventsResponseMetadata `json:"$metadata" api:"required"`
	// The dataset this event belongs to (e.g. cloudflare-workers).
	Dataset string `json:"dataset" api:"required"`
	// Raw log payload. May be a string or a structured object depending on how the log
	// was emitted.
	Source ObservabilityAgentTraceEventsResponseSourceUnion `json:"source" api:"required"`
	// Event timestamp as a Unix epoch in milliseconds.
	Timestamp int64 `json:"timestamp" api:"required"`
	// Cloudflare Containers event information that enriches your logs for identifying
	// and debugging issues.
	Containers map[string]interface{} `json:"$containers"`
	// Cloudflare Workers event information that enriches your logs for identifying and
	// debugging issues.
	Workers ObservabilityAgentTraceEventsResponseWorkers `json:"$workers"`
	JSON    observabilityAgentTraceEventsResponseJSON    `json:"-"`
}

// observabilityAgentTraceEventsResponseJSON contains the JSON metadata for the
// struct [ObservabilityAgentTraceEventsResponse]
type observabilityAgentTraceEventsResponseJSON struct {
	Metadata    apijson.Field
	Dataset     apijson.Field
	Source      apijson.Field
	Timestamp   apijson.Field
	Containers  apijson.Field
	Workers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityAgentTraceEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceEventsResponseJSON) RawJSON() string {
	return r.raw
}

// Structured metadata extracted from the event. These fields are indexed and
// available for filtering and aggregation.
type ObservabilityAgentTraceEventsResponseMetadata struct {
	// Unique event ID. Use as the cursor value for offset-based pagination.
	ID string `json:"id" api:"required"`
	// Cloudflare account identifier.
	Account string `json:"account"`
	// Cloudflare product that generated this event (e.g. workers, pages).
	CloudService string `json:"cloudService"`
	ColdStart    int64  `json:"coldStart"`
	// Estimated cost units for this invocation.
	Cost int64 `json:"cost"`
	// Span duration in milliseconds.
	Duration int64 `json:"duration"`
	// Span end time as a Unix epoch in milliseconds.
	EndTime int64 `json:"endTime"`
	// Span end time as a Unix epoch in nanoseconds.
	EndTimeNS string `json:"endTimeNs"`
	// Error message, present when the log represents an error.
	Error string `json:"error"`
	// Templatized version of the error message used for grouping similar errors.
	ErrorTemplate string `json:"errorTemplate"`
	// Size of the stored telemetry event in bytes.
	EventSize int64 `json:"eventSize"`
	// Content-based fingerprint used to group similar events.
	Fingerprint string `json:"fingerprint"`
	// Log level (e.g. log, debug, info, warn, error).
	Level string `json:"level"`
	// Log message text.
	Message string `json:"message"`
	// Templatized version of the log message used for grouping similar messages.
	MessageTemplate string `json:"messageTemplate"`
	// Metric name when the event represents a metric data point.
	MetricName string `json:"metricName"`
	// Origin of the event (e.g. fetch, scheduled, queue).
	Origin string `json:"origin"`
	// Span ID of the parent span in the trace hierarchy.
	ParentSpanID string `json:"parentSpanId"`
	// Infrastructure provider identifier.
	Provider string `json:"provider"`
	// Cloudflare Ray ID from the `cf-ray` header of the request that triggered the
	// invocation.
	RayID string `json:"rayId"`
	// Cloudflare data center / region that handled the request.
	Region string `json:"region"`
	// Cloudflare request ID that ties all logs from a single invocation together.
	RequestID string `json:"requestId"`
	// Worker script name that produced this event.
	Service string `json:"service"`
	// Span ID for this individual unit of work within a trace.
	SpanID string `json:"spanId"`
	// Human-readable name for this span.
	SpanName string `json:"spanName"`
	// OpenTelemetry span status. Present only on spans. `unset` when the span did not
	// set a status, including spans recorded before status was captured.
	SpanStatus ObservabilityAgentTraceEventsResponseMetadataSpanStatus `json:"spanStatus"`
	// OpenTelemetry span status message. Present only on spans when the status is
	// `error` and a message was set.
	SpanStatusMessage string `json:"spanStatusMessage"`
	// Stack / deployment identifier.
	StackID string `json:"stackId"`
	// Span start time as a Unix epoch in milliseconds.
	StartTime int64 `json:"startTime"`
	// Span start time as a Unix epoch in nanoseconds.
	StartTimeNS string `json:"startTimeNs"`
	// HTTP response status code returned by the Worker.
	StatusCode int64 `json:"statusCode"`
	// Event time as a Unix epoch in nanoseconds.
	TimestampNS string `json:"timestampNs"`
	// Total duration of the entire trace in milliseconds.
	TraceDuration int64 `json:"traceDuration"`
	// Distributed trace ID linking spans across services.
	TraceID string `json:"traceId"`
	// Logical transaction name for this request.
	TransactionName string `json:"transactionName"`
	// What triggered the invocation (e.g. GET /users, POST /orders, queue message).
	Trigger string `json:"trigger"`
	// Event type classifier (e.g. cf-worker-event, cf-worker-log).
	Type string `json:"type"`
	// Request URL that triggered the Worker invocation.
	URL  string                                            `json:"url"`
	JSON observabilityAgentTraceEventsResponseMetadataJSON `json:"-"`
}

// observabilityAgentTraceEventsResponseMetadataJSON contains the JSON metadata for
// the struct [ObservabilityAgentTraceEventsResponseMetadata]
type observabilityAgentTraceEventsResponseMetadataJSON struct {
	ID                apijson.Field
	Account           apijson.Field
	CloudService      apijson.Field
	ColdStart         apijson.Field
	Cost              apijson.Field
	Duration          apijson.Field
	EndTime           apijson.Field
	EndTimeNS         apijson.Field
	Error             apijson.Field
	ErrorTemplate     apijson.Field
	EventSize         apijson.Field
	Fingerprint       apijson.Field
	Level             apijson.Field
	Message           apijson.Field
	MessageTemplate   apijson.Field
	MetricName        apijson.Field
	Origin            apijson.Field
	ParentSpanID      apijson.Field
	Provider          apijson.Field
	RayID             apijson.Field
	Region            apijson.Field
	RequestID         apijson.Field
	Service           apijson.Field
	SpanID            apijson.Field
	SpanName          apijson.Field
	SpanStatus        apijson.Field
	SpanStatusMessage apijson.Field
	StackID           apijson.Field
	StartTime         apijson.Field
	StartTimeNS       apijson.Field
	StatusCode        apijson.Field
	TimestampNS       apijson.Field
	TraceDuration     apijson.Field
	TraceID           apijson.Field
	TransactionName   apijson.Field
	Trigger           apijson.Field
	Type              apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityAgentTraceEventsResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceEventsResponseMetadataJSON) RawJSON() string {
	return r.raw
}

// OpenTelemetry span status. Present only on spans. `unset` when the span did not
// set a status, including spans recorded before status was captured.
type ObservabilityAgentTraceEventsResponseMetadataSpanStatus string

const (
	ObservabilityAgentTraceEventsResponseMetadataSpanStatusUnset ObservabilityAgentTraceEventsResponseMetadataSpanStatus = "unset"
	ObservabilityAgentTraceEventsResponseMetadataSpanStatusOk    ObservabilityAgentTraceEventsResponseMetadataSpanStatus = "ok"
	ObservabilityAgentTraceEventsResponseMetadataSpanStatusError ObservabilityAgentTraceEventsResponseMetadataSpanStatus = "error"
)

func (r ObservabilityAgentTraceEventsResponseMetadataSpanStatus) IsKnown() bool {
	switch r {
	case ObservabilityAgentTraceEventsResponseMetadataSpanStatusUnset, ObservabilityAgentTraceEventsResponseMetadataSpanStatusOk, ObservabilityAgentTraceEventsResponseMetadataSpanStatusError:
		return true
	}
	return false
}

// Raw log payload. May be a string or a structured object depending on how the log
// was emitted.
//
// Union satisfied by [shared.UnionString] or
// [ObservabilityAgentTraceEventsResponseSourceMap].
type ObservabilityAgentTraceEventsResponseSourceUnion interface {
	ImplementsObservabilityAgentTraceEventsResponseSourceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityAgentTraceEventsResponseSourceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityAgentTraceEventsResponseSourceMap{}),
		},
	)
}

type ObservabilityAgentTraceEventsResponseSourceMap map[string]interface{}

func (r ObservabilityAgentTraceEventsResponseSourceMap) ImplementsObservabilityAgentTraceEventsResponseSourceUnion() {
}

// Cloudflare Workers event information that enriches your logs for identifying and
// debugging issues.
type ObservabilityAgentTraceEventsResponseWorkers struct {
	EventType  ObservabilityAgentTraceEventsResponseWorkersEventType `json:"eventType" api:"required"`
	ScriptName string                                                `json:"scriptName" api:"required"`
	CPUTimeMs  float64                                               `json:"cpuTimeMs"`
	// This field can have the runtime type of
	// [[]ObservabilityAgentTraceEventsResponseWorkersObjectDiagnosticsChannelEvent].
	DiagnosticsChannelEvents interface{} `json:"diagnosticsChannelEvents"`
	DispatchNamespace        string      `json:"dispatchNamespace"`
	DurableObjectID          string      `json:"durableObjectId"`
	Entrypoint               string      `json:"entrypoint"`
	// This field can have the runtime type of [map[string]interface{}].
	Event          interface{}                                                `json:"event"`
	ExecutionModel ObservabilityAgentTraceEventsResponseWorkersExecutionModel `json:"executionModel"`
	// This field can have the runtime type of
	// [ObservabilityAgentTraceEventsResponseWorkersObjectGateway].
	Gateway interface{} `json:"gateway"`
	Outcome string      `json:"outcome"`
	// This field can have the runtime type of
	// [ObservabilityAgentTraceEventsResponseWorkersObjectPreview].
	Preview   interface{} `json:"preview"`
	RequestID string      `json:"requestId"`
	// This field can have the runtime type of
	// [ObservabilityAgentTraceEventsResponseWorkersObjectScriptVersion].
	ScriptVersion interface{}                                      `json:"scriptVersion"`
	SpanID        string                                           `json:"spanId"`
	TraceID       string                                           `json:"traceId"`
	Truncated     bool                                             `json:"truncated"`
	WallTimeMs    float64                                          `json:"wallTimeMs"`
	JSON          observabilityAgentTraceEventsResponseWorkersJSON `json:"-"`
	union         ObservabilityAgentTraceEventsResponseWorkersUnion
}

// observabilityAgentTraceEventsResponseWorkersJSON contains the JSON metadata for
// the struct [ObservabilityAgentTraceEventsResponseWorkers]
type observabilityAgentTraceEventsResponseWorkersJSON struct {
	EventType                apijson.Field
	ScriptName               apijson.Field
	CPUTimeMs                apijson.Field
	DiagnosticsChannelEvents apijson.Field
	DispatchNamespace        apijson.Field
	DurableObjectID          apijson.Field
	Entrypoint               apijson.Field
	Event                    apijson.Field
	ExecutionModel           apijson.Field
	Gateway                  apijson.Field
	Outcome                  apijson.Field
	Preview                  apijson.Field
	RequestID                apijson.Field
	ScriptVersion            apijson.Field
	SpanID                   apijson.Field
	TraceID                  apijson.Field
	Truncated                apijson.Field
	WallTimeMs               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r observabilityAgentTraceEventsResponseWorkersJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilityAgentTraceEventsResponseWorkers) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilityAgentTraceEventsResponseWorkers{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObservabilityAgentTraceEventsResponseWorkersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ObservabilityAgentTraceEventsResponseWorkersObject],
// [ObservabilityAgentTraceEventsResponseWorkersObject].
func (r ObservabilityAgentTraceEventsResponseWorkers) AsUnion() ObservabilityAgentTraceEventsResponseWorkersUnion {
	return r.union
}

// Cloudflare Workers event information that enriches your logs for identifying and
// debugging issues.
//
// Union satisfied by [ObservabilityAgentTraceEventsResponseWorkersObject] or
// [ObservabilityAgentTraceEventsResponseWorkersObject].
type ObservabilityAgentTraceEventsResponseWorkersUnion interface {
	implementsObservabilityAgentTraceEventsResponseWorkers()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityAgentTraceEventsResponseWorkersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityAgentTraceEventsResponseWorkersObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityAgentTraceEventsResponseWorkersObject{}),
		},
	)
}

type ObservabilityAgentTraceEventsResponseWorkersObject struct {
	EventType       ObservabilityAgentTraceEventsResponseWorkersObjectEventType      `json:"eventType" api:"required"`
	ScriptName      string                                                           `json:"scriptName" api:"required"`
	DurableObjectID string                                                           `json:"durableObjectId"`
	Entrypoint      string                                                           `json:"entrypoint"`
	Event           map[string]interface{}                                           `json:"event"`
	ExecutionModel  ObservabilityAgentTraceEventsResponseWorkersObjectExecutionModel `json:"executionModel"`
	Gateway         ObservabilityAgentTraceEventsResponseWorkersObjectGateway        `json:"gateway"`
	Outcome         string                                                           `json:"outcome"`
	Preview         ObservabilityAgentTraceEventsResponseWorkersObjectPreview        `json:"preview"`
	RequestID       string                                                           `json:"requestId"`
	ScriptVersion   ObservabilityAgentTraceEventsResponseWorkersObjectScriptVersion  `json:"scriptVersion"`
	SpanID          string                                                           `json:"spanId"`
	TraceID         string                                                           `json:"traceId"`
	Truncated       bool                                                             `json:"truncated"`
	JSON            observabilityAgentTraceEventsResponseWorkersObjectJSON           `json:"-"`
}

// observabilityAgentTraceEventsResponseWorkersObjectJSON contains the JSON
// metadata for the struct [ObservabilityAgentTraceEventsResponseWorkersObject]
type observabilityAgentTraceEventsResponseWorkersObjectJSON struct {
	EventType       apijson.Field
	ScriptName      apijson.Field
	DurableObjectID apijson.Field
	Entrypoint      apijson.Field
	Event           apijson.Field
	ExecutionModel  apijson.Field
	Gateway         apijson.Field
	Outcome         apijson.Field
	Preview         apijson.Field
	RequestID       apijson.Field
	ScriptVersion   apijson.Field
	SpanID          apijson.Field
	TraceID         apijson.Field
	Truncated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilityAgentTraceEventsResponseWorkersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceEventsResponseWorkersObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityAgentTraceEventsResponseWorkersObject) implementsObservabilityAgentTraceEventsResponseWorkers() {
}

type ObservabilityAgentTraceEventsResponseWorkersObjectEventType string

const (
	ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeFetch     ObservabilityAgentTraceEventsResponseWorkersObjectEventType = "fetch"
	ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeScheduled ObservabilityAgentTraceEventsResponseWorkersObjectEventType = "scheduled"
	ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeAlarm     ObservabilityAgentTraceEventsResponseWorkersObjectEventType = "alarm"
	ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeCron      ObservabilityAgentTraceEventsResponseWorkersObjectEventType = "cron"
	ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeQueue     ObservabilityAgentTraceEventsResponseWorkersObjectEventType = "queue"
	ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeEmail     ObservabilityAgentTraceEventsResponseWorkersObjectEventType = "email"
	ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeTail      ObservabilityAgentTraceEventsResponseWorkersObjectEventType = "tail"
	ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeRpc       ObservabilityAgentTraceEventsResponseWorkersObjectEventType = "rpc"
	ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeJsrpc     ObservabilityAgentTraceEventsResponseWorkersObjectEventType = "jsrpc"
	ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeWebsocket ObservabilityAgentTraceEventsResponseWorkersObjectEventType = "websocket"
	ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeWorkflow  ObservabilityAgentTraceEventsResponseWorkersObjectEventType = "workflow"
	ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeUnknown   ObservabilityAgentTraceEventsResponseWorkersObjectEventType = "unknown"
)

func (r ObservabilityAgentTraceEventsResponseWorkersObjectEventType) IsKnown() bool {
	switch r {
	case ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeFetch, ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeScheduled, ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeAlarm, ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeCron, ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeQueue, ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeEmail, ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeTail, ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeRpc, ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeJsrpc, ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeWebsocket, ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeWorkflow, ObservabilityAgentTraceEventsResponseWorkersObjectEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilityAgentTraceEventsResponseWorkersObjectExecutionModel string

const (
	ObservabilityAgentTraceEventsResponseWorkersObjectExecutionModelDurableObject ObservabilityAgentTraceEventsResponseWorkersObjectExecutionModel = "durableObject"
	ObservabilityAgentTraceEventsResponseWorkersObjectExecutionModelStateless     ObservabilityAgentTraceEventsResponseWorkersObjectExecutionModel = "stateless"
)

func (r ObservabilityAgentTraceEventsResponseWorkersObjectExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityAgentTraceEventsResponseWorkersObjectExecutionModelDurableObject, ObservabilityAgentTraceEventsResponseWorkersObjectExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityAgentTraceEventsResponseWorkersObjectGateway struct {
	RequestID   string                                                        `json:"requestId" api:"required"`
	ExtraFields map[string]interface{}                                        `json:"-" api:"extrafields"`
	JSON        observabilityAgentTraceEventsResponseWorkersObjectGatewayJSON `json:"-"`
}

// observabilityAgentTraceEventsResponseWorkersObjectGatewayJSON contains the JSON
// metadata for the struct
// [ObservabilityAgentTraceEventsResponseWorkersObjectGateway]
type observabilityAgentTraceEventsResponseWorkersObjectGatewayJSON struct {
	RequestID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityAgentTraceEventsResponseWorkersObjectGateway) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceEventsResponseWorkersObjectGatewayJSON) RawJSON() string {
	return r.raw
}

type ObservabilityAgentTraceEventsResponseWorkersObjectPreview struct {
	ID   string                                                        `json:"id"`
	Name string                                                        `json:"name"`
	Slug string                                                        `json:"slug"`
	JSON observabilityAgentTraceEventsResponseWorkersObjectPreviewJSON `json:"-"`
}

// observabilityAgentTraceEventsResponseWorkersObjectPreviewJSON contains the JSON
// metadata for the struct
// [ObservabilityAgentTraceEventsResponseWorkersObjectPreview]
type observabilityAgentTraceEventsResponseWorkersObjectPreviewJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityAgentTraceEventsResponseWorkersObjectPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceEventsResponseWorkersObjectPreviewJSON) RawJSON() string {
	return r.raw
}

type ObservabilityAgentTraceEventsResponseWorkersObjectScriptVersion struct {
	ID      string                                                              `json:"id"`
	Message string                                                              `json:"message"`
	Tag     string                                                              `json:"tag"`
	JSON    observabilityAgentTraceEventsResponseWorkersObjectScriptVersionJSON `json:"-"`
}

// observabilityAgentTraceEventsResponseWorkersObjectScriptVersionJSON contains the
// JSON metadata for the struct
// [ObservabilityAgentTraceEventsResponseWorkersObjectScriptVersion]
type observabilityAgentTraceEventsResponseWorkersObjectScriptVersionJSON struct {
	ID          apijson.Field
	Message     apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityAgentTraceEventsResponseWorkersObjectScriptVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceEventsResponseWorkersObjectScriptVersionJSON) RawJSON() string {
	return r.raw
}

type ObservabilityAgentTraceEventsResponseWorkersEventType string

const (
	ObservabilityAgentTraceEventsResponseWorkersEventTypeFetch     ObservabilityAgentTraceEventsResponseWorkersEventType = "fetch"
	ObservabilityAgentTraceEventsResponseWorkersEventTypeScheduled ObservabilityAgentTraceEventsResponseWorkersEventType = "scheduled"
	ObservabilityAgentTraceEventsResponseWorkersEventTypeAlarm     ObservabilityAgentTraceEventsResponseWorkersEventType = "alarm"
	ObservabilityAgentTraceEventsResponseWorkersEventTypeCron      ObservabilityAgentTraceEventsResponseWorkersEventType = "cron"
	ObservabilityAgentTraceEventsResponseWorkersEventTypeQueue     ObservabilityAgentTraceEventsResponseWorkersEventType = "queue"
	ObservabilityAgentTraceEventsResponseWorkersEventTypeEmail     ObservabilityAgentTraceEventsResponseWorkersEventType = "email"
	ObservabilityAgentTraceEventsResponseWorkersEventTypeTail      ObservabilityAgentTraceEventsResponseWorkersEventType = "tail"
	ObservabilityAgentTraceEventsResponseWorkersEventTypeRpc       ObservabilityAgentTraceEventsResponseWorkersEventType = "rpc"
	ObservabilityAgentTraceEventsResponseWorkersEventTypeJsrpc     ObservabilityAgentTraceEventsResponseWorkersEventType = "jsrpc"
	ObservabilityAgentTraceEventsResponseWorkersEventTypeWebsocket ObservabilityAgentTraceEventsResponseWorkersEventType = "websocket"
	ObservabilityAgentTraceEventsResponseWorkersEventTypeWorkflow  ObservabilityAgentTraceEventsResponseWorkersEventType = "workflow"
	ObservabilityAgentTraceEventsResponseWorkersEventTypeUnknown   ObservabilityAgentTraceEventsResponseWorkersEventType = "unknown"
)

func (r ObservabilityAgentTraceEventsResponseWorkersEventType) IsKnown() bool {
	switch r {
	case ObservabilityAgentTraceEventsResponseWorkersEventTypeFetch, ObservabilityAgentTraceEventsResponseWorkersEventTypeScheduled, ObservabilityAgentTraceEventsResponseWorkersEventTypeAlarm, ObservabilityAgentTraceEventsResponseWorkersEventTypeCron, ObservabilityAgentTraceEventsResponseWorkersEventTypeQueue, ObservabilityAgentTraceEventsResponseWorkersEventTypeEmail, ObservabilityAgentTraceEventsResponseWorkersEventTypeTail, ObservabilityAgentTraceEventsResponseWorkersEventTypeRpc, ObservabilityAgentTraceEventsResponseWorkersEventTypeJsrpc, ObservabilityAgentTraceEventsResponseWorkersEventTypeWebsocket, ObservabilityAgentTraceEventsResponseWorkersEventTypeWorkflow, ObservabilityAgentTraceEventsResponseWorkersEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilityAgentTraceEventsResponseWorkersExecutionModel string

const (
	ObservabilityAgentTraceEventsResponseWorkersExecutionModelDurableObject ObservabilityAgentTraceEventsResponseWorkersExecutionModel = "durableObject"
	ObservabilityAgentTraceEventsResponseWorkersExecutionModelStateless     ObservabilityAgentTraceEventsResponseWorkersExecutionModel = "stateless"
)

func (r ObservabilityAgentTraceEventsResponseWorkersExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityAgentTraceEventsResponseWorkersExecutionModelDurableObject, ObservabilityAgentTraceEventsResponseWorkersExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityAgentTraceGetResponse struct {
	// Sorted distinct gen_ai.agent.name values of authorized Agent invocations in the
	// Trace within the effective window. Includes every Agent in the Trace, not only
	// those matching the agent filter.
	Agents []string `json:"agents" api:"required"`
	// Sorted distinct gen_ai.conversation.id values of authorized Agent invocations in
	// the Trace within the effective window. Empty when none are recorded.
	ConversationIDs []string `json:"conversationIds" api:"required"`
	// Recorded spans with a nonempty error or ERROR status. Counts spans, not distinct
	// messages or incidents.
	ErrorCount int64 `json:"errorCount" api:"required"`
	// Input tokens reported by chat spans in the authorized Trace window; null when
	// none report input usage. Informational, not billing data.
	InputTokens float64 `json:"inputTokens" api:"required,nullable"`
	// Latest start of an authorized Agent invocation in the Trace within the effective
	// window, as Unix epoch milliseconds. The Trace list is ordered by this value,
	// newest first.
	LastActivityMs int64 `json:"lastActivityMs" api:"required"`
	// Recorded chat spans in the authorized Trace window.
	LlmCallCount int64 `json:"llmCallCount" api:"required"`
	// Recorded chat spans with a nonempty error or ERROR status.
	LlmErrorCount int64 `json:"llmErrorCount" api:"required"`
	// Sorted distinct response models, falling back to request models, from chat spans
	// in the authorized Trace window.
	Models []string `json:"models" api:"required"`
	// Output tokens reported by chat spans in the authorized Trace window; null when
	// none report output usage. Informational, not billing data.
	OutputTokens float64 `json:"outputTokens" api:"required,nullable"`
	// Sorted distinct GenAI providers from chat spans in the authorized Trace window.
	Providers []string `json:"providers" api:"required"`
	// Name of the root span that initiated the trace.
	RootSpanName string `json:"rootSpanName" api:"required"`
	// Logical transaction name for the root span.
	RootTransactionName string `json:"rootTransactionName" api:"required"`
	// Authorized Worker services represented in the trace.
	Services []string `json:"services" api:"required"`
	// Total number of spans in the trace.
	Spans float64 `json:"spans" api:"required"`
	// Recorded execute_tool spans in the authorized Trace window. Custom-named
	// approval spans are not counted.
	ToolCallCount int64 `json:"toolCallCount" api:"required"`
	// Recorded execute_tool spans with a nonempty error or ERROR status.
	ToolErrorCount int64 `json:"toolErrorCount" api:"required"`
	// Input plus output tokens; null if either direction is unreported. Does not add
	// cache or reasoning breakdowns.
	TotalTokens float64 `json:"totalTokens" api:"required,nullable"`
	// Total duration of the trace in milliseconds.
	TraceDurationMs float64 `json:"traceDurationMs" api:"required"`
	// Trace end time as a Unix epoch in milliseconds.
	TraceEndMs float64 `json:"traceEndMs" api:"required"`
	// Unique identifier for the distributed trace.
	TraceID string `json:"traceId" api:"required"`
	// Trace start time as a Unix epoch in milliseconds.
	TraceStartMs float64 `json:"traceStartMs" api:"required"`
	// Error messages encountered during the trace, if any.
	Errors []string                               `json:"errors"`
	JSON   observabilityAgentTraceGetResponseJSON `json:"-"`
}

// observabilityAgentTraceGetResponseJSON contains the JSON metadata for the struct
// [ObservabilityAgentTraceGetResponse]
type observabilityAgentTraceGetResponseJSON struct {
	Agents              apijson.Field
	ConversationIDs     apijson.Field
	ErrorCount          apijson.Field
	InputTokens         apijson.Field
	LastActivityMs      apijson.Field
	LlmCallCount        apijson.Field
	LlmErrorCount       apijson.Field
	Models              apijson.Field
	OutputTokens        apijson.Field
	Providers           apijson.Field
	RootSpanName        apijson.Field
	RootTransactionName apijson.Field
	Services            apijson.Field
	Spans               apijson.Field
	ToolCallCount       apijson.Field
	ToolErrorCount      apijson.Field
	TotalTokens         apijson.Field
	TraceDurationMs     apijson.Field
	TraceEndMs          apijson.Field
	TraceID             apijson.Field
	TraceStartMs        apijson.Field
	Errors              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ObservabilityAgentTraceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceGetResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityAgentTraceListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Only include work in which an Agent with this gen_ai.agent.name took part. With
	// service, both must match the same Agent.
	Agent param.Field[string] `query:"agent"`
	// Opaque cursor from result_info.cursors.after in the previous response. Send it
	// with the same filters; its window replaces from and to.
	Cursor param.Field[string] `query:"cursor"`
	// Start of the search window, as Unix epoch milliseconds. Defaults to 24 hours
	// before to. Clamped to the 7-day retention boundary; the effective value is in
	// result_info.from.
	From param.Field[int64] `query:"from"`
	// Maximum items to return per page. Defaults to 20.
	PerPage param.Field[int64] `query:"per_page"`
	// Only include work in which an Agent of this Worker service took part.
	Service param.Field[string] `query:"service"`
	// End of the search window, as Unix epoch milliseconds. Defaults to now.
	To param.Field[int64] `query:"to"`
}

// URLQuery serializes [ObservabilityAgentTraceListParams]'s query parameters as
// `url.Values`.
func (r ObservabilityAgentTraceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ObservabilityAgentTraceEventsParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Opaque cursor from result_info.cursors.after in the previous response. Send it
	// with the same filters; its window replaces from and to.
	Cursor param.Field[string] `query:"cursor"`
	// Optional. Narrows the search, which by default covers the full 7-day retention,
	// so any retained Trace ID can be fetched without a window. Start as Unix epoch
	// milliseconds; the effective value is in result_info.from.
	From param.Field[int64] `query:"from"`
	// Maximum items to return per page. Defaults to 100.
	PerPage param.Field[int64] `query:"per_page"`
	// Optional. Narrows the search; end as Unix epoch milliseconds. Defaults to now.
	To param.Field[int64] `query:"to"`
}

// URLQuery serializes [ObservabilityAgentTraceEventsParams]'s query parameters as
// `url.Values`.
func (r ObservabilityAgentTraceEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ObservabilityAgentTraceEventsResponseEnvelope struct {
	Errors     []ObservabilityAgentTraceEventsResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []ObservabilityAgentTraceEventsResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     []ObservabilityAgentTraceEventsResponse                 `json:"result" api:"required"`
	ResultInfo ObservabilityAgentTraceEventsResponseEnvelopeResultInfo `json:"result_info" api:"required"`
	Success    ObservabilityAgentTraceEventsResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON       observabilityAgentTraceEventsResponseEnvelopeJSON       `json:"-"`
}

// observabilityAgentTraceEventsResponseEnvelopeJSON contains the JSON metadata for
// the struct [ObservabilityAgentTraceEventsResponseEnvelope]
type observabilityAgentTraceEventsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityAgentTraceEventsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceEventsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityAgentTraceEventsResponseEnvelopeErrors struct {
	Message string                                                  `json:"message" api:"required"`
	JSON    observabilityAgentTraceEventsResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityAgentTraceEventsResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ObservabilityAgentTraceEventsResponseEnvelopeErrors]
type observabilityAgentTraceEventsResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityAgentTraceEventsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceEventsResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityAgentTraceEventsResponseEnvelopeMessages struct {
	Message ObservabilityAgentTraceEventsResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityAgentTraceEventsResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityAgentTraceEventsResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilityAgentTraceEventsResponseEnvelopeMessages]
type observabilityAgentTraceEventsResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityAgentTraceEventsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceEventsResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityAgentTraceEventsResponseEnvelopeMessagesMessage string

const (
	ObservabilityAgentTraceEventsResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityAgentTraceEventsResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityAgentTraceEventsResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityAgentTraceEventsResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityAgentTraceEventsResponseEnvelopeResultInfo struct {
	// Items returned on this page. Not a total: a page can be short or empty while
	// more results remain.
	Count   int64                                                          `json:"count" api:"required"`
	Cursors ObservabilityAgentTraceEventsResponseEnvelopeResultInfoCursors `json:"cursors" api:"required"`
	// Effective start of the time window, after defaults and retention clamping.
	From int64 `json:"from" api:"required"`
	// Maximum items requested for this page.
	PerPage int64 `json:"per_page" api:"required"`
	// Effective end of the time window.
	To   int64                                                       `json:"to" api:"required"`
	JSON observabilityAgentTraceEventsResponseEnvelopeResultInfoJSON `json:"-"`
}

// observabilityAgentTraceEventsResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ObservabilityAgentTraceEventsResponseEnvelopeResultInfo]
type observabilityAgentTraceEventsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Cursors     apijson.Field
	From        apijson.Field
	PerPage     apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityAgentTraceEventsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceEventsResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ObservabilityAgentTraceEventsResponseEnvelopeResultInfoCursors struct {
	// Opaque cursor for the next page. Only null means there are no more results;
	// earlier pages can be short or empty.
	After string                                                             `json:"after" api:"required,nullable"`
	JSON  observabilityAgentTraceEventsResponseEnvelopeResultInfoCursorsJSON `json:"-"`
}

// observabilityAgentTraceEventsResponseEnvelopeResultInfoCursorsJSON contains the
// JSON metadata for the struct
// [ObservabilityAgentTraceEventsResponseEnvelopeResultInfoCursors]
type observabilityAgentTraceEventsResponseEnvelopeResultInfoCursorsJSON struct {
	After       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityAgentTraceEventsResponseEnvelopeResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceEventsResponseEnvelopeResultInfoCursorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityAgentTraceEventsResponseEnvelopeSuccess bool

const (
	ObservabilityAgentTraceEventsResponseEnvelopeSuccessTrue ObservabilityAgentTraceEventsResponseEnvelopeSuccess = true
)

func (r ObservabilityAgentTraceEventsResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityAgentTraceEventsResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityAgentTraceGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Optional. Narrows the search, which by default covers the full 7-day retention,
	// so any retained Trace ID can be fetched without a window. Start as Unix epoch
	// milliseconds; the effective value is in result_info.from.
	From param.Field[int64] `query:"from"`
	// Optional. Narrows the search; end as Unix epoch milliseconds. Defaults to now.
	To param.Field[int64] `query:"to"`
}

// URLQuery serializes [ObservabilityAgentTraceGetParams]'s query parameters as
// `url.Values`.
func (r ObservabilityAgentTraceGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ObservabilityAgentTraceGetResponseEnvelope struct {
	Errors     []ObservabilityAgentTraceGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []ObservabilityAgentTraceGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     ObservabilityAgentTraceGetResponse                   `json:"result" api:"required"`
	ResultInfo ObservabilityAgentTraceGetResponseEnvelopeResultInfo `json:"result_info" api:"required"`
	Success    ObservabilityAgentTraceGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON       observabilityAgentTraceGetResponseEnvelopeJSON       `json:"-"`
}

// observabilityAgentTraceGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ObservabilityAgentTraceGetResponseEnvelope]
type observabilityAgentTraceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityAgentTraceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityAgentTraceGetResponseEnvelopeErrors struct {
	Message string                                               `json:"message" api:"required"`
	JSON    observabilityAgentTraceGetResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityAgentTraceGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ObservabilityAgentTraceGetResponseEnvelopeErrors]
type observabilityAgentTraceGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityAgentTraceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityAgentTraceGetResponseEnvelopeMessages struct {
	Message ObservabilityAgentTraceGetResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityAgentTraceGetResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityAgentTraceGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilityAgentTraceGetResponseEnvelopeMessages]
type observabilityAgentTraceGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityAgentTraceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityAgentTraceGetResponseEnvelopeMessagesMessage string

const (
	ObservabilityAgentTraceGetResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityAgentTraceGetResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityAgentTraceGetResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityAgentTraceGetResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityAgentTraceGetResponseEnvelopeResultInfo struct {
	// Effective start of the time window, after defaults and retention clamping.
	From int64 `json:"from" api:"required"`
	// Effective end of the time window.
	To   int64                                                    `json:"to" api:"required"`
	JSON observabilityAgentTraceGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// observabilityAgentTraceGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ObservabilityAgentTraceGetResponseEnvelopeResultInfo]
type observabilityAgentTraceGetResponseEnvelopeResultInfoJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityAgentTraceGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityAgentTraceGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ObservabilityAgentTraceGetResponseEnvelopeSuccess bool

const (
	ObservabilityAgentTraceGetResponseEnvelopeSuccessTrue ObservabilityAgentTraceGetResponseEnvelopeSuccess = true
)

func (r ObservabilityAgentTraceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityAgentTraceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
