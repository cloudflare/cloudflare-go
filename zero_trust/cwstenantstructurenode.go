// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// CWSTenantStructureNodeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCWSTenantStructureNodeService] method instead.
type CWSTenantStructureNodeService struct {
	Options []option.RequestOption
}

// NewCWSTenantStructureNodeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCWSTenantStructureNodeService(opts ...option.RequestOption) (r *CWSTenantStructureNodeService) {
	r = &CWSTenantStructureNodeService{}
	r.Options = opts
	return
}

// Adds a validated cloud-native node to the tenant hierarchy.
func (r *CWSTenantStructureNodeService) New(ctx context.Context, structureID string, params CWSTenantStructureNodeNewParams, opts ...option.RequestOption) (res *TenantStructureNode, err error) {
	var env CWSTenantStructureNodeNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if structureID == "" {
		err = errors.New("missing required structure_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/tenant-structures/%s/nodes", params.AccountID, structureID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates the node's display name or parent while preserving its immutable
// identity fields.
func (r *CWSTenantStructureNodeService) Update(ctx context.Context, structureID string, nodeID string, params CWSTenantStructureNodeUpdateParams, opts ...option.RequestOption) (res *TenantStructureNode, err error) {
	var env CWSTenantStructureNodeUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if structureID == "" {
		err = errors.New("missing required structure_id parameter")
		return nil, err
	}
	if nodeID == "" {
		err = errors.New("missing required node_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/tenant-structures/%s/nodes/%s", params.AccountID, structureID, nodeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of nodes in the tenant structure.
func (r *CWSTenantStructureNodeService) List(ctx context.Context, structureID string, params CWSTenantStructureNodeListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[TenantStructureNode], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if structureID == "" {
		err = errors.New("missing required structure_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/tenant-structures/%s/nodes", params.AccountID, structureID)
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

// Returns a paginated list of nodes in the tenant structure.
func (r *CWSTenantStructureNodeService) ListAutoPaging(ctx context.Context, structureID string, params CWSTenantStructureNodeListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[TenantStructureNode] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, structureID, params, opts...))
}

// Deletes a childless node from the tenant hierarchy.
func (r *CWSTenantStructureNodeService) Delete(ctx context.Context, structureID string, nodeID string, body CWSTenantStructureNodeDeleteParams, opts ...option.RequestOption) (res *CWSTenantStructureNodeDeleteResponse, err error) {
	var env CWSTenantStructureNodeDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if structureID == "" {
		err = errors.New("missing required structure_id parameter")
		return nil, err
	}
	if nodeID == "" {
		err = errors.New("missing required node_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/tenant-structures/%s/nodes/%s", body.AccountID, structureID, nodeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Atomically validates and applies 1 to 500 unordered node operations. Creates may
// reference new parents but cannot reuse a native_id deleted in the same batch.
// Maximum body size: 1 MiB.
func (r *CWSTenantStructureNodeService) Batch(ctx context.Context, structureID string, params CWSTenantStructureNodeBatchParams, opts ...option.RequestOption) (res *CWSTenantStructureNodeBatchResponse, err error) {
	var env CWSTenantStructureNodeBatchResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if structureID == "" {
		err = errors.New("missing required structure_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/tenant-structures/%s/nodes/batch", params.AccountID, structureID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the tenant node with the specified ID.
func (r *CWSTenantStructureNodeService) Get(ctx context.Context, structureID string, nodeID string, query CWSTenantStructureNodeGetParams, opts ...option.RequestOption) (res *TenantStructureNode, err error) {
	var env CWSTenantStructureNodeGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if structureID == "" {
		err = errors.New("missing required structure_id parameter")
		return nil, err
	}
	if nodeID == "" {
		err = errors.New("missing required node_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/tenant-structures/%s/nodes/%s", query.AccountID, structureID, nodeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type TenantStructureNode struct {
	// Server-assigned UUID
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Cloud-native identifier of this node, unique within the structure; the
	// parent-reference target. Shape depends on the structure's application. Immutable
	// after create.
	NativeID string `json:"native_id" api:"required"`
	// Position in the tree. Legal values depend on the structure's application (aws:
	// root/ou/account; cloudflare: org/account/zone). At most one root per structure.
	// Immutable after create.
	NodeKind  string    `json:"node_kind" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Display name (e.g. an OU or account name), mutable. Optional; empty by default.
	Name string `json:"name"`
	// Native ID of this node's parent. Required for non-root nodes and omitted for the
	// root. The parent must be in the same structure and able to contain this node
	// kind.
	Parent string                  `json:"parent"`
	JSON   tenantStructureNodeJSON `json:"-"`
}

// tenantStructureNodeJSON contains the JSON metadata for the struct
// [TenantStructureNode]
type tenantStructureNodeJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	NativeID    apijson.Field
	NodeKind    apijson.Field
	UpdatedAt   apijson.Field
	Name        apijson.Field
	Parent      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantStructureNode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantStructureNodeJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeDeleteResponse struct {
	ID   string                                   `json:"id" api:"required" format:"uuid"`
	JSON cwsTenantStructureNodeDeleteResponseJSON `json:"-"`
}

// cwsTenantStructureNodeDeleteResponseJSON contains the JSON metadata for the
// struct [CWSTenantStructureNodeDeleteResponse]
type cwsTenantStructureNodeDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeBatchResponse struct {
	Created []TenantStructureNode                        `json:"created" api:"required"`
	Deleted []CWSTenantStructureNodeBatchResponseDeleted `json:"deleted" api:"required"`
	Updated []TenantStructureNode                        `json:"updated" api:"required"`
	JSON    cwsTenantStructureNodeBatchResponseJSON      `json:"-"`
}

// cwsTenantStructureNodeBatchResponseJSON contains the JSON metadata for the
// struct [CWSTenantStructureNodeBatchResponse]
type cwsTenantStructureNodeBatchResponseJSON struct {
	Created     apijson.Field
	Deleted     apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeBatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeBatchResponseJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeBatchResponseDeleted struct {
	ID   string                                         `json:"id" api:"required" format:"uuid"`
	JSON cwsTenantStructureNodeBatchResponseDeletedJSON `json:"-"`
}

// cwsTenantStructureNodeBatchResponseDeletedJSON contains the JSON metadata for
// the struct [CWSTenantStructureNodeBatchResponseDeleted]
type cwsTenantStructureNodeBatchResponseDeletedJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeBatchResponseDeleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeBatchResponseDeletedJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeNewParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Cloud-native identifier of this node, unique within the structure; the
	// parent-reference target. Shape depends on the structure's application. Immutable
	// after create.
	NativeID param.Field[string] `json:"native_id" api:"required"`
	// Position in the tree. Legal values depend on the structure's application (aws:
	// root/ou/account; cloudflare: org/account/zone). At most one root per structure.
	// Immutable after create.
	NodeKind param.Field[string] `json:"node_kind" api:"required"`
	// Display name (e.g. an OU or account name), mutable. Optional; empty by default.
	Name param.Field[string] `json:"name"`
	// Native ID of this node's parent. Required for non-root nodes and omitted for the
	// root. The parent must be in the same structure and able to contain this node
	// kind.
	Parent param.Field[string] `json:"parent"`
}

func (r CWSTenantStructureNodeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSTenantStructureNodeNewResponseEnvelope struct {
	Errors   []CWSTenantStructureNodeNewResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSTenantStructureNodeNewResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   TenantStructureNode                                 `json:"result" api:"required"`
	Success  bool                                                `json:"success" api:"required"`
	JSON     cwsTenantStructureNodeNewResponseEnvelopeJSON       `json:"-"`
}

// cwsTenantStructureNodeNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSTenantStructureNodeNewResponseEnvelope]
type cwsTenantStructureNodeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeNewResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code" api:"required"`
	Message string                                                `json:"message" api:"required"`
	Source  CWSTenantStructureNodeNewResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsTenantStructureNodeNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsTenantStructureNodeNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CWSTenantStructureNodeNewResponseEnvelopeErrors]
type cwsTenantStructureNodeNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer" api:"required"`
	JSON    cwsTenantStructureNodeNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsTenantStructureNodeNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CWSTenantStructureNodeNewResponseEnvelopeErrorsSource]
type cwsTenantStructureNodeNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeNewResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code" api:"required"`
	Message string                                                  `json:"message" api:"required"`
	Source  CWSTenantStructureNodeNewResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsTenantStructureNodeNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsTenantStructureNodeNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CWSTenantStructureNodeNewResponseEnvelopeMessages]
type cwsTenantStructureNodeNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer" api:"required"`
	JSON    cwsTenantStructureNodeNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsTenantStructureNodeNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CWSTenantStructureNodeNewResponseEnvelopeMessagesSource]
type cwsTenantStructureNodeNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeUpdateParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Cloud-native identifier of this node, unique within the structure; the
	// parent-reference target. Shape depends on the structure's application. Immutable
	// after create.
	NativeID param.Field[string] `json:"native_id" api:"required"`
	// Position in the tree. Legal values depend on the structure's application (aws:
	// root/ou/account; cloudflare: org/account/zone). At most one root per structure.
	// Immutable after create.
	NodeKind param.Field[string] `json:"node_kind" api:"required"`
	// Display name (e.g. an OU or account name), mutable. Optional; empty by default.
	Name param.Field[string] `json:"name"`
	// Native ID of this node's parent. Required for non-root nodes and omitted for the
	// root. The parent must be in the same structure and able to contain this node
	// kind.
	Parent param.Field[string] `json:"parent"`
}

func (r CWSTenantStructureNodeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSTenantStructureNodeUpdateResponseEnvelope struct {
	Errors   []CWSTenantStructureNodeUpdateResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSTenantStructureNodeUpdateResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   TenantStructureNode                                    `json:"result" api:"required"`
	Success  bool                                                   `json:"success" api:"required"`
	JSON     cwsTenantStructureNodeUpdateResponseEnvelopeJSON       `json:"-"`
}

// cwsTenantStructureNodeUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [CWSTenantStructureNodeUpdateResponseEnvelope]
type cwsTenantStructureNodeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeUpdateResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code" api:"required"`
	Message string                                                   `json:"message" api:"required"`
	Source  CWSTenantStructureNodeUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsTenantStructureNodeUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsTenantStructureNodeUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CWSTenantStructureNodeUpdateResponseEnvelopeErrors]
type cwsTenantStructureNodeUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer" api:"required"`
	JSON    cwsTenantStructureNodeUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsTenantStructureNodeUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [CWSTenantStructureNodeUpdateResponseEnvelopeErrorsSource]
type cwsTenantStructureNodeUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeUpdateResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code" api:"required"`
	Message string                                                     `json:"message" api:"required"`
	Source  CWSTenantStructureNodeUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsTenantStructureNodeUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsTenantStructureNodeUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CWSTenantStructureNodeUpdateResponseEnvelopeMessages]
type cwsTenantStructureNodeUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer" api:"required"`
	JSON    cwsTenantStructureNodeUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsTenantStructureNodeUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CWSTenantStructureNodeUpdateResponseEnvelopeMessagesSource]
type cwsTenantStructureNodeUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeListParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Sort direction (default: asc)
	Direction param.Field[CWSTenantStructureNodeListParamsDirection] `query:"direction"`
	// Filter by node name (case-insensitive substring match)
	Name param.Field[string] `query:"name"`
	// Filter by exact cloud-native identifier
	NativeID param.Field[string] `query:"native_id"`
	// Filter by exact node kind
	NodeKind param.Field[string] `query:"node_kind"`
	// Field to order results by (default: created_at)
	Order param.Field[CWSTenantStructureNodeListParamsOrder] `query:"order"`
	// Filter by exact parent native ID
	Parent param.Field[string] `query:"parent"`
	// Number of items per page (1-100, default 20)
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [CWSTenantStructureNodeListParams]'s query parameters as
// `url.Values`.
func (r CWSTenantStructureNodeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction (default: asc)
type CWSTenantStructureNodeListParamsDirection string

const (
	CWSTenantStructureNodeListParamsDirectionAsc  CWSTenantStructureNodeListParamsDirection = "asc"
	CWSTenantStructureNodeListParamsDirectionDesc CWSTenantStructureNodeListParamsDirection = "desc"
)

func (r CWSTenantStructureNodeListParamsDirection) IsKnown() bool {
	switch r {
	case CWSTenantStructureNodeListParamsDirectionAsc, CWSTenantStructureNodeListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by (default: created_at)
type CWSTenantStructureNodeListParamsOrder string

const (
	CWSTenantStructureNodeListParamsOrderName      CWSTenantStructureNodeListParamsOrder = "name"
	CWSTenantStructureNodeListParamsOrderCreatedAt CWSTenantStructureNodeListParamsOrder = "created_at"
	CWSTenantStructureNodeListParamsOrderUpdatedAt CWSTenantStructureNodeListParamsOrder = "updated_at"
)

func (r CWSTenantStructureNodeListParamsOrder) IsKnown() bool {
	switch r {
	case CWSTenantStructureNodeListParamsOrderName, CWSTenantStructureNodeListParamsOrderCreatedAt, CWSTenantStructureNodeListParamsOrderUpdatedAt:
		return true
	}
	return false
}

type CWSTenantStructureNodeDeleteParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSTenantStructureNodeDeleteResponseEnvelope struct {
	Errors   []CWSTenantStructureNodeDeleteResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSTenantStructureNodeDeleteResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   CWSTenantStructureNodeDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                                   `json:"success" api:"required"`
	JSON     cwsTenantStructureNodeDeleteResponseEnvelopeJSON       `json:"-"`
}

// cwsTenantStructureNodeDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [CWSTenantStructureNodeDeleteResponseEnvelope]
type cwsTenantStructureNodeDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeDeleteResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code" api:"required"`
	Message string                                                   `json:"message" api:"required"`
	Source  CWSTenantStructureNodeDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsTenantStructureNodeDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsTenantStructureNodeDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CWSTenantStructureNodeDeleteResponseEnvelopeErrors]
type cwsTenantStructureNodeDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer" api:"required"`
	JSON    cwsTenantStructureNodeDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsTenantStructureNodeDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [CWSTenantStructureNodeDeleteResponseEnvelopeErrorsSource]
type cwsTenantStructureNodeDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeDeleteResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code" api:"required"`
	Message string                                                     `json:"message" api:"required"`
	Source  CWSTenantStructureNodeDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsTenantStructureNodeDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsTenantStructureNodeDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CWSTenantStructureNodeDeleteResponseEnvelopeMessages]
type cwsTenantStructureNodeDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer" api:"required"`
	JSON    cwsTenantStructureNodeDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsTenantStructureNodeDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CWSTenantStructureNodeDeleteResponseEnvelopeMessagesSource]
type cwsTenantStructureNodeDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeBatchParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Nodes to create; parents may be created elsewhere in this request
	Creates param.Field[[]CWSTenantStructureNodeBatchParamsCreate] `json:"creates" api:"required"`
	// Nodes to delete; every descendant that will not be reparented must also be
	// deleted
	Deletes param.Field[[]CWSTenantStructureNodeBatchParamsDelete] `json:"deletes" api:"required"`
	// Nodes whose mutable fields should be fully replaced
	Updates param.Field[[]CWSTenantStructureNodeBatchParamsUpdate] `json:"updates" api:"required"`
}

func (r CWSTenantStructureNodeBatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSTenantStructureNodeBatchParamsCreate struct {
	// Cloud-native identifier of this node, unique within the structure; the
	// parent-reference target. Shape depends on the structure's application. Immutable
	// after create.
	NativeID param.Field[string] `json:"native_id" api:"required"`
	// Position in the tree. Legal values depend on the structure's application (aws:
	// root/ou/account; cloudflare: org/account/zone). At most one root per structure.
	// Immutable after create.
	NodeKind param.Field[string] `json:"node_kind" api:"required"`
	// Display name (e.g. an OU or account name), mutable. Optional; empty by default.
	Name param.Field[string] `json:"name"`
	// Native ID of this node's parent. Required for non-root nodes and omitted for the
	// root. The parent must be in the same structure and able to contain this node
	// kind.
	Parent param.Field[string] `json:"parent"`
}

func (r CWSTenantStructureNodeBatchParamsCreate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSTenantStructureNodeBatchParamsDelete struct {
	// ID of the node to delete
	ID param.Field[string] `json:"id" api:"required" format:"uuid"`
}

func (r CWSTenantStructureNodeBatchParamsDelete) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSTenantStructureNodeBatchParamsUpdate struct {
	// ID of the node to update
	ID param.Field[string] `json:"id" api:"required" format:"uuid"`
	// Full replacement display name
	Name param.Field[string] `json:"name"`
	// Full replacement parent native ID; omit for the root
	Parent param.Field[string] `json:"parent"`
}

func (r CWSTenantStructureNodeBatchParamsUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSTenantStructureNodeBatchResponseEnvelope struct {
	Errors   []CWSTenantStructureNodeBatchResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSTenantStructureNodeBatchResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   CWSTenantStructureNodeBatchResponse                   `json:"result" api:"required"`
	Success  bool                                                  `json:"success" api:"required"`
	JSON     cwsTenantStructureNodeBatchResponseEnvelopeJSON       `json:"-"`
}

// cwsTenantStructureNodeBatchResponseEnvelopeJSON contains the JSON metadata for
// the struct [CWSTenantStructureNodeBatchResponseEnvelope]
type cwsTenantStructureNodeBatchResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeBatchResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeBatchResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeBatchResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code" api:"required"`
	Message string                                                  `json:"message" api:"required"`
	Source  CWSTenantStructureNodeBatchResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsTenantStructureNodeBatchResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsTenantStructureNodeBatchResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CWSTenantStructureNodeBatchResponseEnvelopeErrors]
type cwsTenantStructureNodeBatchResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeBatchResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeBatchResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeBatchResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer" api:"required"`
	JSON    cwsTenantStructureNodeBatchResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsTenantStructureNodeBatchResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [CWSTenantStructureNodeBatchResponseEnvelopeErrorsSource]
type cwsTenantStructureNodeBatchResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeBatchResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeBatchResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeBatchResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code" api:"required"`
	Message string                                                    `json:"message" api:"required"`
	Source  CWSTenantStructureNodeBatchResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsTenantStructureNodeBatchResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsTenantStructureNodeBatchResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CWSTenantStructureNodeBatchResponseEnvelopeMessages]
type cwsTenantStructureNodeBatchResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeBatchResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeBatchResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeBatchResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer" api:"required"`
	JSON    cwsTenantStructureNodeBatchResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsTenantStructureNodeBatchResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CWSTenantStructureNodeBatchResponseEnvelopeMessagesSource]
type cwsTenantStructureNodeBatchResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeBatchResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeBatchResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeGetParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSTenantStructureNodeGetResponseEnvelope struct {
	Errors   []CWSTenantStructureNodeGetResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSTenantStructureNodeGetResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   TenantStructureNode                                 `json:"result" api:"required"`
	Success  bool                                                `json:"success" api:"required"`
	JSON     cwsTenantStructureNodeGetResponseEnvelopeJSON       `json:"-"`
}

// cwsTenantStructureNodeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSTenantStructureNodeGetResponseEnvelope]
type cwsTenantStructureNodeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code" api:"required"`
	Message string                                                `json:"message" api:"required"`
	Source  CWSTenantStructureNodeGetResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsTenantStructureNodeGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsTenantStructureNodeGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CWSTenantStructureNodeGetResponseEnvelopeErrors]
type cwsTenantStructureNodeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer" api:"required"`
	JSON    cwsTenantStructureNodeGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsTenantStructureNodeGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CWSTenantStructureNodeGetResponseEnvelopeErrorsSource]
type cwsTenantStructureNodeGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code" api:"required"`
	Message string                                                  `json:"message" api:"required"`
	Source  CWSTenantStructureNodeGetResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsTenantStructureNodeGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsTenantStructureNodeGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CWSTenantStructureNodeGetResponseEnvelopeMessages]
type cwsTenantStructureNodeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNodeGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer" api:"required"`
	JSON    cwsTenantStructureNodeGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsTenantStructureNodeGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CWSTenantStructureNodeGetResponseEnvelopeMessagesSource]
type cwsTenantStructureNodeGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNodeGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNodeGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
