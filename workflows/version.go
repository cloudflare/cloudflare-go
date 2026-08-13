// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflows

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// VersionService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVersionService] method instead.
type VersionService struct {
	Options []option.RequestOption
}

// NewVersionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVersionService(opts ...option.RequestOption) (r *VersionService) {
	r = &VersionService{}
	r.Options = opts
	return
}

// Lists all deployed versions of a workflow.
func (r *VersionService) List(ctx context.Context, workflowName string, params VersionListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[VersionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/versions", params.AccountID, workflowName)
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

// Lists all deployed versions of a workflow.
func (r *VersionService) ListAutoPaging(ctx context.Context, workflowName string, params VersionListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[VersionListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, workflowName, params, opts...))
}

// Retrieves details for a specific deployed workflow version.
func (r *VersionService) Get(ctx context.Context, workflowName string, versionID string, query VersionGetParams, opts ...option.RequestOption) (res *VersionGetResponse, err error) {
	var env VersionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/versions/%s", query.AccountID, workflowName, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves the graph visualization of a workflow version.
func (r *VersionService) Graph(ctx context.Context, workflowName string, versionID string, query VersionGraphParams, opts ...option.RequestOption) (res *VersionGraphResponse, err error) {
	var env VersionGraphResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/versions/%s/graph", query.AccountID, workflowName, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type VersionListResponse struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	ClassName string    `json:"class_name" api:"required"`
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	HasDag    bool      `json:"has_dag" api:"required"`
	// The programming language of the workflow implementation
	Language   VersionListResponseLanguage `json:"language" api:"required"`
	ModifiedOn time.Time                   `json:"modified_on" api:"required" format:"date-time"`
	WorkflowID string                      `json:"workflow_id" api:"required" format:"uuid"`
	Limits     VersionListResponseLimits   `json:"limits"`
	JSON       versionListResponseJSON     `json:"-"`
}

// versionListResponseJSON contains the JSON metadata for the struct
// [VersionListResponse]
type versionListResponseJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	HasDag      apijson.Field
	Language    apijson.Field
	ModifiedOn  apijson.Field
	WorkflowID  apijson.Field
	Limits      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseJSON) RawJSON() string {
	return r.raw
}

// The programming language of the workflow implementation
type VersionListResponseLanguage string

const (
	VersionListResponseLanguageJavascript VersionListResponseLanguage = "javascript"
	VersionListResponseLanguagePython     VersionListResponseLanguage = "python"
)

func (r VersionListResponseLanguage) IsKnown() bool {
	switch r {
	case VersionListResponseLanguageJavascript, VersionListResponseLanguagePython:
		return true
	}
	return false
}

type VersionListResponseLimits struct {
	Steps int64                         `json:"steps"`
	JSON  versionListResponseLimitsJSON `json:"-"`
}

// versionListResponseLimitsJSON contains the JSON metadata for the struct
// [VersionListResponseLimits]
type versionListResponseLimitsJSON struct {
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponseLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseLimitsJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponse struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	ClassName string    `json:"class_name" api:"required"`
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	HasDag    bool      `json:"has_dag" api:"required"`
	// The programming language of the workflow implementation
	Language   VersionGetResponseLanguage `json:"language" api:"required"`
	ModifiedOn time.Time                  `json:"modified_on" api:"required" format:"date-time"`
	WorkflowID string                     `json:"workflow_id" api:"required" format:"uuid"`
	Limits     VersionGetResponseLimits   `json:"limits"`
	JSON       versionGetResponseJSON     `json:"-"`
}

// versionGetResponseJSON contains the JSON metadata for the struct
// [VersionGetResponse]
type versionGetResponseJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	HasDag      apijson.Field
	Language    apijson.Field
	ModifiedOn  apijson.Field
	WorkflowID  apijson.Field
	Limits      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseJSON) RawJSON() string {
	return r.raw
}

// The programming language of the workflow implementation
type VersionGetResponseLanguage string

const (
	VersionGetResponseLanguageJavascript VersionGetResponseLanguage = "javascript"
	VersionGetResponseLanguagePython     VersionGetResponseLanguage = "python"
)

func (r VersionGetResponseLanguage) IsKnown() bool {
	switch r {
	case VersionGetResponseLanguageJavascript, VersionGetResponseLanguagePython:
		return true
	}
	return false
}

type VersionGetResponseLimits struct {
	Steps int64                        `json:"steps"`
	JSON  versionGetResponseLimitsJSON `json:"-"`
}

// versionGetResponseLimitsJSON contains the JSON metadata for the struct
// [VersionGetResponseLimits]
type versionGetResponseLimitsJSON struct {
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseLimitsJSON) RawJSON() string {
	return r.raw
}

type VersionGraphResponse struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	ClassName string    `json:"class_name" api:"required"`
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Versioned workflow graph payload.
	Graph      VersionGraphResponseGraph `json:"graph" api:"required,nullable"`
	ModifiedOn time.Time                 `json:"modified_on" api:"required" format:"date-time"`
	WorkflowID string                    `json:"workflow_id" api:"required" format:"uuid"`
	JSON       versionGraphResponseJSON  `json:"-"`
}

// versionGraphResponseJSON contains the JSON metadata for the struct
// [VersionGraphResponse]
type versionGraphResponseJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	Graph       apijson.Field
	ModifiedOn  apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGraphResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGraphResponseJSON) RawJSON() string {
	return r.raw
}

// Versioned workflow graph payload.
type VersionGraphResponseGraph struct {
	Version float64 `json:"version" api:"required"`
	// A parsed workflow entrypoint with its step graph.
	Workflow VersionGraphResponseGraphWorkflow `json:"workflow" api:"required"`
	JSON     versionGraphResponseGraphJSON     `json:"-"`
}

// versionGraphResponseGraphJSON contains the JSON metadata for the struct
// [VersionGraphResponseGraph]
type versionGraphResponseGraphJSON struct {
	Version     apijson.Field
	Workflow    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGraphResponseGraph) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGraphResponseGraphJSON) RawJSON() string {
	return r.raw
}

// A parsed workflow entrypoint with its step graph.
type VersionGraphResponseGraphWorkflow struct {
	ClassName string                                               `json:"class_name" api:"required"`
	Functions map[string]VersionGraphResponseGraphWorkflowFunction `json:"functions" api:"required"`
	Nodes     []VersionGraphResponseGraphWorkflowNode              `json:"nodes" api:"required"`
	// Shape descriptor for JSON payloads.
	Payload VersionGraphResponseGraphWorkflowPayload `json:"payload"`
	JSON    versionGraphResponseGraphWorkflowJSON    `json:"-"`
}

// versionGraphResponseGraphWorkflowJSON contains the JSON metadata for the struct
// [VersionGraphResponseGraphWorkflow]
type versionGraphResponseGraphWorkflowJSON struct {
	ClassName   apijson.Field
	Functions   apijson.Field
	Nodes       apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGraphResponseGraphWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGraphResponseGraphWorkflowJSON) RawJSON() string {
	return r.raw
}

type VersionGraphResponseGraphWorkflowFunction struct {
	Name string `json:"name" api:"required"`
	// Child nodes (recursive).
	Nodes []interface{}                                  `json:"nodes" api:"required"`
	Type  VersionGraphResponseGraphWorkflowFunctionsType `json:"type" api:"required"`
	JSON  versionGraphResponseGraphWorkflowFunctionJSON  `json:"-"`
}

// versionGraphResponseGraphWorkflowFunctionJSON contains the JSON metadata for the
// struct [VersionGraphResponseGraphWorkflowFunction]
type versionGraphResponseGraphWorkflowFunctionJSON struct {
	Name        apijson.Field
	Nodes       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGraphResponseGraphWorkflowFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGraphResponseGraphWorkflowFunctionJSON) RawJSON() string {
	return r.raw
}

type VersionGraphResponseGraphWorkflowFunctionsType string

const (
	VersionGraphResponseGraphWorkflowFunctionsTypeFunctionDef VersionGraphResponseGraphWorkflowFunctionsType = "function_def"
)

func (r VersionGraphResponseGraphWorkflowFunctionsType) IsKnown() bool {
	switch r {
	case VersionGraphResponseGraphWorkflowFunctionsTypeFunctionDef:
		return true
	}
	return false
}

type VersionGraphResponseGraphWorkflowNode struct {
	Type VersionGraphResponseGraphWorkflowNodesType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [[]VersionGraphResponseGraphWorkflowNodesObjectBranch].
	Branches interface{} `json:"branches"`
	// This field can have the runtime type of
	// [VersionGraphResponseGraphWorkflowNodesObjectCatchBlock].
	CatchBlock interface{} `json:"catch_block"`
	ClassName  string      `json:"class_name"`
	// This field can have the runtime type of
	// [VersionGraphResponseGraphWorkflowNodesObjectConfig].
	Config       interface{} `json:"config"`
	Discriminant string      `json:"discriminant"`
	// This field can have the runtime type of
	// [VersionGraphResponseGraphWorkflowNodesObjectDurationUnion].
	Duration interface{} `json:"duration"`
	// This field can have the runtime type of
	// [VersionGraphResponseGraphWorkflowNodesObjectFinallyBlock].
	FinallyBlock interface{} `json:"finally_block"`
	// This field can have the runtime type of
	// [map[string]VersionGraphResponseGraphWorkflowNodesObjectFunction].
	Functions interface{} `json:"functions"`
	// Parallel execution strategy.
	Kind VersionGraphResponseGraphWorkflowNodesKind `json:"kind"`
	Name string                                     `json:"name"`
	// This field can have the runtime type of [[]interface{}].
	Nodes interface{} `json:"nodes"`
	// This field can have the runtime type of
	// [VersionGraphResponseGraphWorkflowNodesObjectOptions].
	Options interface{} `json:"options"`
	// This field can have the runtime type of
	// [VersionGraphResponseGraphWorkflowNodesObjectPayload].
	Payload   interface{} `json:"payload"`
	Resolves  float64     `json:"resolves"`
	Starts    float64     `json:"starts"`
	Timestamp string      `json:"timestamp"`
	// This field can have the runtime type of
	// [VersionGraphResponseGraphWorkflowNodesObjectTryBlock].
	TryBlock interface{}                               `json:"try_block"`
	JSON     versionGraphResponseGraphWorkflowNodeJSON `json:"-"`
	union    VersionGraphResponseGraphWorkflowNodesUnion
}

// versionGraphResponseGraphWorkflowNodeJSON contains the JSON metadata for the
// struct [VersionGraphResponseGraphWorkflowNode]
type versionGraphResponseGraphWorkflowNodeJSON struct {
	Type         apijson.Field
	Branches     apijson.Field
	CatchBlock   apijson.Field
	ClassName    apijson.Field
	Config       apijson.Field
	Discriminant apijson.Field
	Duration     apijson.Field
	FinallyBlock apijson.Field
	Functions    apijson.Field
	Kind         apijson.Field
	Name         apijson.Field
	Nodes        apijson.Field
	Options      apijson.Field
	Payload      apijson.Field
	Resolves     apijson.Field
	Starts       apijson.Field
	Timestamp    apijson.Field
	TryBlock     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r versionGraphResponseGraphWorkflowNodeJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGraphResponseGraphWorkflowNode) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGraphResponseGraphWorkflowNode{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [VersionGraphResponseGraphWorkflowNodesUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject].
func (r VersionGraphResponseGraphWorkflowNode) AsUnion() VersionGraphResponseGraphWorkflowNodesUnion {
	return r.union
}

// Union satisfied by [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject],
// [VersionGraphResponseGraphWorkflowNodesObject] or
// [VersionGraphResponseGraphWorkflowNodesObject].
type VersionGraphResponseGraphWorkflowNodesUnion interface {
	implementsVersionGraphResponseGraphWorkflowNode()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGraphResponseGraphWorkflowNodesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowNodesObject{}),
		},
	)
}

type VersionGraphResponseGraphWorkflowNodesObject struct {
	// Duration as milliseconds (number) or human-readable string.
	Duration VersionGraphResponseGraphWorkflowNodesObjectDurationUnion `json:"duration" api:"required"`
	Name     string                                                    `json:"name" api:"required"`
	Type     VersionGraphResponseGraphWorkflowNodesObjectType          `json:"type" api:"required"`
	Resolves float64                                                   `json:"resolves"`
	Starts   float64                                                   `json:"starts"`
	JSON     versionGraphResponseGraphWorkflowNodesObjectJSON          `json:"-"`
}

// versionGraphResponseGraphWorkflowNodesObjectJSON contains the JSON metadata for
// the struct [VersionGraphResponseGraphWorkflowNodesObject]
type versionGraphResponseGraphWorkflowNodesObjectJSON struct {
	Duration    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Resolves    apijson.Field
	Starts      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGraphResponseGraphWorkflowNodesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGraphResponseGraphWorkflowNodesObjectJSON) RawJSON() string {
	return r.raw
}

func (r VersionGraphResponseGraphWorkflowNodesObject) implementsVersionGraphResponseGraphWorkflowNode() {
}

// Duration as milliseconds (number) or human-readable string.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type VersionGraphResponseGraphWorkflowNodesObjectDurationUnion interface {
	ImplementsVersionGraphResponseGraphWorkflowNodesObjectDurationUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGraphResponseGraphWorkflowNodesObjectDurationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type VersionGraphResponseGraphWorkflowNodesObjectType string

const (
	VersionGraphResponseGraphWorkflowNodesObjectTypeStepSleep VersionGraphResponseGraphWorkflowNodesObjectType = "step_sleep"
)

func (r VersionGraphResponseGraphWorkflowNodesObjectType) IsKnown() bool {
	switch r {
	case VersionGraphResponseGraphWorkflowNodesObjectTypeStepSleep:
		return true
	}
	return false
}

type VersionGraphResponseGraphWorkflowNodesType string

const (
	VersionGraphResponseGraphWorkflowNodesTypeStepSleep        VersionGraphResponseGraphWorkflowNodesType = "step_sleep"
	VersionGraphResponseGraphWorkflowNodesTypeStepDo           VersionGraphResponseGraphWorkflowNodesType = "step_do"
	VersionGraphResponseGraphWorkflowNodesTypeStepWaitForEvent VersionGraphResponseGraphWorkflowNodesType = "step_wait_for_event"
	VersionGraphResponseGraphWorkflowNodesTypeStepSleepUntil   VersionGraphResponseGraphWorkflowNodesType = "step_sleep_until"
	VersionGraphResponseGraphWorkflowNodesTypeLoop             VersionGraphResponseGraphWorkflowNodesType = "loop"
	VersionGraphResponseGraphWorkflowNodesTypeParallel         VersionGraphResponseGraphWorkflowNodesType = "parallel"
	VersionGraphResponseGraphWorkflowNodesTypeTry              VersionGraphResponseGraphWorkflowNodesType = "try"
	VersionGraphResponseGraphWorkflowNodesTypeBlock            VersionGraphResponseGraphWorkflowNodesType = "block"
	VersionGraphResponseGraphWorkflowNodesTypeIf               VersionGraphResponseGraphWorkflowNodesType = "if"
	VersionGraphResponseGraphWorkflowNodesTypeSwitch           VersionGraphResponseGraphWorkflowNodesType = "switch"
	VersionGraphResponseGraphWorkflowNodesTypeStart            VersionGraphResponseGraphWorkflowNodesType = "start"
	VersionGraphResponseGraphWorkflowNodesTypeFunctionCall     VersionGraphResponseGraphWorkflowNodesType = "function_call"
	VersionGraphResponseGraphWorkflowNodesTypeFunctionDef      VersionGraphResponseGraphWorkflowNodesType = "function_def"
	VersionGraphResponseGraphWorkflowNodesTypeBreak            VersionGraphResponseGraphWorkflowNodesType = "break"
)

func (r VersionGraphResponseGraphWorkflowNodesType) IsKnown() bool {
	switch r {
	case VersionGraphResponseGraphWorkflowNodesTypeStepSleep, VersionGraphResponseGraphWorkflowNodesTypeStepDo, VersionGraphResponseGraphWorkflowNodesTypeStepWaitForEvent, VersionGraphResponseGraphWorkflowNodesTypeStepSleepUntil, VersionGraphResponseGraphWorkflowNodesTypeLoop, VersionGraphResponseGraphWorkflowNodesTypeParallel, VersionGraphResponseGraphWorkflowNodesTypeTry, VersionGraphResponseGraphWorkflowNodesTypeBlock, VersionGraphResponseGraphWorkflowNodesTypeIf, VersionGraphResponseGraphWorkflowNodesTypeSwitch, VersionGraphResponseGraphWorkflowNodesTypeStart, VersionGraphResponseGraphWorkflowNodesTypeFunctionCall, VersionGraphResponseGraphWorkflowNodesTypeFunctionDef, VersionGraphResponseGraphWorkflowNodesTypeBreak:
		return true
	}
	return false
}

// Parallel execution strategy.
type VersionGraphResponseGraphWorkflowNodesKind string

const (
	VersionGraphResponseGraphWorkflowNodesKindAll        VersionGraphResponseGraphWorkflowNodesKind = "all"
	VersionGraphResponseGraphWorkflowNodesKindAny        VersionGraphResponseGraphWorkflowNodesKind = "any"
	VersionGraphResponseGraphWorkflowNodesKindAllSettled VersionGraphResponseGraphWorkflowNodesKind = "all_settled"
	VersionGraphResponseGraphWorkflowNodesKindRace       VersionGraphResponseGraphWorkflowNodesKind = "race"
	VersionGraphResponseGraphWorkflowNodesKindBreak      VersionGraphResponseGraphWorkflowNodesKind = "break"
	VersionGraphResponseGraphWorkflowNodesKindReturn     VersionGraphResponseGraphWorkflowNodesKind = "return"
)

func (r VersionGraphResponseGraphWorkflowNodesKind) IsKnown() bool {
	switch r {
	case VersionGraphResponseGraphWorkflowNodesKindAll, VersionGraphResponseGraphWorkflowNodesKindAny, VersionGraphResponseGraphWorkflowNodesKindAllSettled, VersionGraphResponseGraphWorkflowNodesKindRace, VersionGraphResponseGraphWorkflowNodesKindBreak, VersionGraphResponseGraphWorkflowNodesKindReturn:
		return true
	}
	return false
}

// Shape descriptor for JSON payloads.
type VersionGraphResponseGraphWorkflowPayload struct {
	Type VersionGraphResponseGraphWorkflowPayloadType `json:"type" api:"required"`
	// This field can have the runtime type of [map[string]interface{}].
	Fields interface{}                                  `json:"fields"`
	JSON   versionGraphResponseGraphWorkflowPayloadJSON `json:"-"`
	union  VersionGraphResponseGraphWorkflowPayloadUnion
}

// versionGraphResponseGraphWorkflowPayloadJSON contains the JSON metadata for the
// struct [VersionGraphResponseGraphWorkflowPayload]
type versionGraphResponseGraphWorkflowPayloadJSON struct {
	Type        apijson.Field
	Fields      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r versionGraphResponseGraphWorkflowPayloadJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGraphResponseGraphWorkflowPayload) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGraphResponseGraphWorkflowPayload{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [VersionGraphResponseGraphWorkflowPayloadUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGraphResponseGraphWorkflowPayloadType],
// [VersionGraphResponseGraphWorkflowPayloadObject].
func (r VersionGraphResponseGraphWorkflowPayload) AsUnion() VersionGraphResponseGraphWorkflowPayloadUnion {
	return r.union
}

// Shape descriptor for JSON payloads.
//
// Union satisfied by [VersionGraphResponseGraphWorkflowPayloadType] or
// [VersionGraphResponseGraphWorkflowPayloadObject].
type VersionGraphResponseGraphWorkflowPayloadUnion interface {
	implementsVersionGraphResponseGraphWorkflowPayload()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGraphResponseGraphWorkflowPayloadUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowPayloadType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGraphResponseGraphWorkflowPayloadObject{}),
		},
	)
}

type VersionGraphResponseGraphWorkflowPayloadType struct {
	Type VersionGraphResponseGraphWorkflowPayloadTypeType `json:"type" api:"required"`
	JSON versionGraphResponseGraphWorkflowPayloadTypeJSON `json:"-"`
}

// versionGraphResponseGraphWorkflowPayloadTypeJSON contains the JSON metadata for
// the struct [VersionGraphResponseGraphWorkflowPayloadType]
type versionGraphResponseGraphWorkflowPayloadTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGraphResponseGraphWorkflowPayloadType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGraphResponseGraphWorkflowPayloadTypeJSON) RawJSON() string {
	return r.raw
}

func (r VersionGraphResponseGraphWorkflowPayloadType) implementsVersionGraphResponseGraphWorkflowPayload() {
}

type VersionGraphResponseGraphWorkflowPayloadTypeType string

const (
	VersionGraphResponseGraphWorkflowPayloadTypeTypeUnknown VersionGraphResponseGraphWorkflowPayloadTypeType = "unknown"
)

func (r VersionGraphResponseGraphWorkflowPayloadTypeType) IsKnown() bool {
	switch r {
	case VersionGraphResponseGraphWorkflowPayloadTypeTypeUnknown:
		return true
	}
	return false
}

type VersionGraphResponseGraphWorkflowPayloadObject struct {
	// Nested JsonShape fields (recursive structure).
	Fields map[string]interface{}                             `json:"fields" api:"required"`
	Type   VersionGraphResponseGraphWorkflowPayloadObjectType `json:"type" api:"required"`
	JSON   versionGraphResponseGraphWorkflowPayloadObjectJSON `json:"-"`
}

// versionGraphResponseGraphWorkflowPayloadObjectJSON contains the JSON metadata
// for the struct [VersionGraphResponseGraphWorkflowPayloadObject]
type versionGraphResponseGraphWorkflowPayloadObjectJSON struct {
	Fields      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGraphResponseGraphWorkflowPayloadObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGraphResponseGraphWorkflowPayloadObjectJSON) RawJSON() string {
	return r.raw
}

func (r VersionGraphResponseGraphWorkflowPayloadObject) implementsVersionGraphResponseGraphWorkflowPayload() {
}

type VersionGraphResponseGraphWorkflowPayloadObjectType string

const (
	VersionGraphResponseGraphWorkflowPayloadObjectTypeObject VersionGraphResponseGraphWorkflowPayloadObjectType = "object"
)

func (r VersionGraphResponseGraphWorkflowPayloadObjectType) IsKnown() bool {
	switch r {
	case VersionGraphResponseGraphWorkflowPayloadObjectTypeObject:
		return true
	}
	return false
}

type VersionListParams struct {
	AccountID param.Field[string]  `path:"account_id" api:"required"`
	Page      param.Field[float64] `query:"page"`
	PerPage   param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [VersionListParams]'s query parameters as `url.Values`.
func (r VersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type VersionGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type VersionGetResponseEnvelope struct {
	Errors     []VersionGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []VersionGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     VersionGetResponse                   `json:"result" api:"required"`
	Success    VersionGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo VersionGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       versionGetResponseEnvelopeJSON       `json:"-"`
}

// versionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [VersionGetResponseEnvelope]
type versionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseEnvelopeErrors struct {
	Code    float64                              `json:"code" api:"required"`
	Message string                               `json:"message" api:"required"`
	JSON    versionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// versionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [VersionGetResponseEnvelopeErrors]
type versionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseEnvelopeMessages struct {
	Code    float64                                `json:"code" api:"required"`
	Message string                                 `json:"message" api:"required"`
	JSON    versionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// versionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [VersionGetResponseEnvelopeMessages]
type versionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseEnvelopeSuccess bool

const (
	VersionGetResponseEnvelopeSuccessTrue VersionGetResponseEnvelopeSuccess = true
)

func (r VersionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VersionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type VersionGetResponseEnvelopeResultInfo struct {
	Count      float64                                  `json:"count" api:"required"`
	PerPage    float64                                  `json:"per_page" api:"required"`
	TotalCount float64                                  `json:"total_count" api:"required"`
	Cursor     string                                   `json:"cursor"`
	Page       float64                                  `json:"page"`
	TotalPages float64                                  `json:"total_pages"`
	JSON       versionGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// versionGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [VersionGetResponseEnvelopeResultInfo]
type versionGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	Cursor      apijson.Field
	Page        apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type VersionGraphParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type VersionGraphResponseEnvelope struct {
	Errors     []VersionGraphResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []VersionGraphResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     VersionGraphResponse                   `json:"result" api:"required"`
	Success    VersionGraphResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo VersionGraphResponseEnvelopeResultInfo `json:"result_info"`
	JSON       versionGraphResponseEnvelopeJSON       `json:"-"`
}

// versionGraphResponseEnvelopeJSON contains the JSON metadata for the struct
// [VersionGraphResponseEnvelope]
type versionGraphResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGraphResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGraphResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type VersionGraphResponseEnvelopeErrors struct {
	Code    float64                                `json:"code" api:"required"`
	Message string                                 `json:"message" api:"required"`
	JSON    versionGraphResponseEnvelopeErrorsJSON `json:"-"`
}

// versionGraphResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [VersionGraphResponseEnvelopeErrors]
type versionGraphResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGraphResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGraphResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type VersionGraphResponseEnvelopeMessages struct {
	Code    float64                                  `json:"code" api:"required"`
	Message string                                   `json:"message" api:"required"`
	JSON    versionGraphResponseEnvelopeMessagesJSON `json:"-"`
}

// versionGraphResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VersionGraphResponseEnvelopeMessages]
type versionGraphResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGraphResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGraphResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type VersionGraphResponseEnvelopeSuccess bool

const (
	VersionGraphResponseEnvelopeSuccessTrue VersionGraphResponseEnvelopeSuccess = true
)

func (r VersionGraphResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VersionGraphResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type VersionGraphResponseEnvelopeResultInfo struct {
	Count      float64                                    `json:"count" api:"required"`
	PerPage    float64                                    `json:"per_page" api:"required"`
	TotalCount float64                                    `json:"total_count" api:"required"`
	Cursor     string                                     `json:"cursor"`
	Page       float64                                    `json:"page"`
	TotalPages float64                                    `json:"total_pages"`
	JSON       versionGraphResponseEnvelopeResultInfoJSON `json:"-"`
}

// versionGraphResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [VersionGraphResponseEnvelopeResultInfo]
type versionGraphResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	Cursor      apijson.Field
	Page        apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGraphResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGraphResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
