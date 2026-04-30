// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// NamespaceInstanceItemService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceInstanceItemService] method instead.
type NamespaceInstanceItemService struct {
	Options []option.RequestOption
}

// NewNamespaceInstanceItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNamespaceInstanceItemService(opts ...option.RequestOption) (r *NamespaceInstanceItemService) {
	r = &NamespaceInstanceItemService{}
	r.Options = opts
	return
}

// Lists indexed items in an AI Search instance.
func (r *NamespaceInstanceItemService) List(ctx context.Context, name string, id string, params NamespaceInstanceItemListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[NamespaceInstanceItemListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/items", params.AccountID, name, id)
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

// Lists indexed items in an AI Search instance.
func (r *NamespaceInstanceItemService) ListAutoPaging(ctx context.Context, name string, id string, params NamespaceInstanceItemListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[NamespaceInstanceItemListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, name, id, params, opts...))
}

// Deletes a file from a managed AI Search instance and triggers a reindex.
func (r *NamespaceInstanceItemService) Delete(ctx context.Context, name string, id string, itemID string, body NamespaceInstanceItemDeleteParams, opts ...option.RequestOption) (res *NamespaceInstanceItemDeleteResponse, err error) {
	var env NamespaceInstanceItemDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/items/%s", body.AccountID, name, id, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists chunks for a specific item in an AI Search instance, including their text
// content.
func (r *NamespaceInstanceItemService) Chunks(ctx context.Context, name string, id string, itemID string, params NamespaceInstanceItemChunksParams, opts ...option.RequestOption) (res *[]NamespaceInstanceItemChunksResponse, err error) {
	var env NamespaceInstanceItemChunksResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/items/%s/chunks", params.AccountID, name, id, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Creates or updates an indexed item in an AI Search instance.
func (r *NamespaceInstanceItemService) NewOrUpdate(ctx context.Context, name string, id string, params NamespaceInstanceItemNewOrUpdateParams, opts ...option.RequestOption) (res *NamespaceInstanceItemNewOrUpdateResponse, err error) {
	var env NamespaceInstanceItemNewOrUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/items", params.AccountID, name, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Downloads the raw file content for a specific item from the managed AI Search
// instance storage.
func (r *NamespaceInstanceItemService) Download(ctx context.Context, name string, id string, itemID string, query NamespaceInstanceItemDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/items/%s/download", query.AccountID, name, id, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves a specific indexed item from an AI Search instance.
func (r *NamespaceInstanceItemService) Get(ctx context.Context, name string, id string, itemID string, query NamespaceInstanceItemGetParams, opts ...option.RequestOption) (res *NamespaceInstanceItemGetResponse, err error) {
	var env NamespaceInstanceItemGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/items/%s", query.AccountID, name, id, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists processing logs for a specific item in an AI Search instance.
func (r *NamespaceInstanceItemService) Logs(ctx context.Context, name string, id string, itemID string, params NamespaceInstanceItemLogsParams, opts ...option.RequestOption) (res *[]NamespaceInstanceItemLogsResponse, err error) {
	var env NamespaceInstanceItemLogsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/items/%s/logs", params.AccountID, name, id, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Syncs an item to an AI Search instance index.
func (r *NamespaceInstanceItemService) Sync(ctx context.Context, name string, id string, itemID string, params NamespaceInstanceItemSyncParams, opts ...option.RequestOption) (res *NamespaceInstanceItemSyncResponse, err error) {
	var env NamespaceInstanceItemSyncResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/items/%s", params.AccountID, name, id, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Uploads a file to a managed AI Search instance via multipart/form-data (max
// 4MB).
func (r *NamespaceInstanceItemService) Upload(ctx context.Context, name string, id string, params NamespaceInstanceItemUploadParams, opts ...option.RequestOption) (res *NamespaceInstanceItemUploadResponse, err error) {
	var env NamespaceInstanceItemUploadResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/items", params.AccountID, name, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type NamespaceInstanceItemListResponse struct {
	ID          string                                      `json:"id" api:"required"`
	Checksum    string                                      `json:"checksum" api:"required"`
	ChunksCount int64                                       `json:"chunks_count" api:"required,nullable"`
	CreatedAt   time.Time                                   `json:"created_at" api:"required" format:"date-time"`
	FileSize    float64                                     `json:"file_size" api:"required,nullable"`
	Key         string                                      `json:"key" api:"required"`
	LastSeenAt  time.Time                                   `json:"last_seen_at" api:"required" format:"date-time"`
	Namespace   string                                      `json:"namespace" api:"required"`
	NextAction  NamespaceInstanceItemListResponseNextAction `json:"next_action" api:"required,nullable"`
	// Identifies which data source this item belongs to. "builtin" for uploaded files,
	// "{type}:{source}" for external sources, null for legacy items.
	SourceID string                                  `json:"source_id" api:"required,nullable"`
	Status   NamespaceInstanceItemListResponseStatus `json:"status" api:"required"`
	Error    string                                  `json:"error"`
	JSON     namespaceInstanceItemListResponseJSON   `json:"-"`
}

// namespaceInstanceItemListResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceItemListResponse]
type namespaceInstanceItemListResponseJSON struct {
	ID          apijson.Field
	Checksum    apijson.Field
	ChunksCount apijson.Field
	CreatedAt   apijson.Field
	FileSize    apijson.Field
	Key         apijson.Field
	LastSeenAt  apijson.Field
	Namespace   apijson.Field
	NextAction  apijson.Field
	SourceID    apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemListResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemListResponseNextAction string

const (
	NamespaceInstanceItemListResponseNextActionIndex  NamespaceInstanceItemListResponseNextAction = "INDEX"
	NamespaceInstanceItemListResponseNextActionDelete NamespaceInstanceItemListResponseNextAction = "DELETE"
)

func (r NamespaceInstanceItemListResponseNextAction) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemListResponseNextActionIndex, NamespaceInstanceItemListResponseNextActionDelete:
		return true
	}
	return false
}

type NamespaceInstanceItemListResponseStatus string

const (
	NamespaceInstanceItemListResponseStatusQueued    NamespaceInstanceItemListResponseStatus = "queued"
	NamespaceInstanceItemListResponseStatusRunning   NamespaceInstanceItemListResponseStatus = "running"
	NamespaceInstanceItemListResponseStatusCompleted NamespaceInstanceItemListResponseStatus = "completed"
	NamespaceInstanceItemListResponseStatusError     NamespaceInstanceItemListResponseStatus = "error"
	NamespaceInstanceItemListResponseStatusSkipped   NamespaceInstanceItemListResponseStatus = "skipped"
	NamespaceInstanceItemListResponseStatusOutdated  NamespaceInstanceItemListResponseStatus = "outdated"
)

func (r NamespaceInstanceItemListResponseStatus) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemListResponseStatusQueued, NamespaceInstanceItemListResponseStatusRunning, NamespaceInstanceItemListResponseStatusCompleted, NamespaceInstanceItemListResponseStatusError, NamespaceInstanceItemListResponseStatusSkipped, NamespaceInstanceItemListResponseStatusOutdated:
		return true
	}
	return false
}

type NamespaceInstanceItemDeleteResponse struct {
	Key  string                                  `json:"key" api:"required"`
	JSON namespaceInstanceItemDeleteResponseJSON `json:"-"`
}

// namespaceInstanceItemDeleteResponseJSON contains the JSON metadata for the
// struct [NamespaceInstanceItemDeleteResponse]
type namespaceInstanceItemDeleteResponseJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemChunksResponse struct {
	ID        string                                  `json:"id" api:"required"`
	Item      NamespaceInstanceItemChunksResponseItem `json:"item" api:"required"`
	Text      string                                  `json:"text" api:"required"`
	EndByte   float64                                 `json:"end_byte"`
	StartByte float64                                 `json:"start_byte"`
	JSON      namespaceInstanceItemChunksResponseJSON `json:"-"`
}

// namespaceInstanceItemChunksResponseJSON contains the JSON metadata for the
// struct [NamespaceInstanceItemChunksResponse]
type namespaceInstanceItemChunksResponseJSON struct {
	ID          apijson.Field
	Item        apijson.Field
	Text        apijson.Field
	EndByte     apijson.Field
	StartByte   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemChunksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemChunksResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemChunksResponseItem struct {
	Key       string                                      `json:"key" api:"required"`
	Metadata  map[string]interface{}                      `json:"metadata"`
	Timestamp float64                                     `json:"timestamp"`
	JSON      namespaceInstanceItemChunksResponseItemJSON `json:"-"`
}

// namespaceInstanceItemChunksResponseItemJSON contains the JSON metadata for the
// struct [NamespaceInstanceItemChunksResponseItem]
type namespaceInstanceItemChunksResponseItemJSON struct {
	Key         apijson.Field
	Metadata    apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemChunksResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemChunksResponseItemJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemNewOrUpdateResponse struct {
	ID          string                                             `json:"id" api:"required"`
	Checksum    string                                             `json:"checksum" api:"required"`
	ChunksCount int64                                              `json:"chunks_count" api:"required,nullable"`
	CreatedAt   time.Time                                          `json:"created_at" api:"required" format:"date-time"`
	FileSize    float64                                            `json:"file_size" api:"required,nullable"`
	Key         string                                             `json:"key" api:"required"`
	LastSeenAt  time.Time                                          `json:"last_seen_at" api:"required" format:"date-time"`
	Namespace   string                                             `json:"namespace" api:"required"`
	NextAction  NamespaceInstanceItemNewOrUpdateResponseNextAction `json:"next_action" api:"required,nullable"`
	// Identifies which data source this item belongs to. "builtin" for uploaded files,
	// "{type}:{source}" for external sources, null for legacy items.
	SourceID string                                         `json:"source_id" api:"required,nullable"`
	Status   NamespaceInstanceItemNewOrUpdateResponseStatus `json:"status" api:"required"`
	Error    string                                         `json:"error"`
	JSON     namespaceInstanceItemNewOrUpdateResponseJSON   `json:"-"`
}

// namespaceInstanceItemNewOrUpdateResponseJSON contains the JSON metadata for the
// struct [NamespaceInstanceItemNewOrUpdateResponse]
type namespaceInstanceItemNewOrUpdateResponseJSON struct {
	ID          apijson.Field
	Checksum    apijson.Field
	ChunksCount apijson.Field
	CreatedAt   apijson.Field
	FileSize    apijson.Field
	Key         apijson.Field
	LastSeenAt  apijson.Field
	Namespace   apijson.Field
	NextAction  apijson.Field
	SourceID    apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemNewOrUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemNewOrUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemNewOrUpdateResponseNextAction string

const (
	NamespaceInstanceItemNewOrUpdateResponseNextActionIndex  NamespaceInstanceItemNewOrUpdateResponseNextAction = "INDEX"
	NamespaceInstanceItemNewOrUpdateResponseNextActionDelete NamespaceInstanceItemNewOrUpdateResponseNextAction = "DELETE"
)

func (r NamespaceInstanceItemNewOrUpdateResponseNextAction) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemNewOrUpdateResponseNextActionIndex, NamespaceInstanceItemNewOrUpdateResponseNextActionDelete:
		return true
	}
	return false
}

type NamespaceInstanceItemNewOrUpdateResponseStatus string

const (
	NamespaceInstanceItemNewOrUpdateResponseStatusQueued    NamespaceInstanceItemNewOrUpdateResponseStatus = "queued"
	NamespaceInstanceItemNewOrUpdateResponseStatusRunning   NamespaceInstanceItemNewOrUpdateResponseStatus = "running"
	NamespaceInstanceItemNewOrUpdateResponseStatusCompleted NamespaceInstanceItemNewOrUpdateResponseStatus = "completed"
	NamespaceInstanceItemNewOrUpdateResponseStatusError     NamespaceInstanceItemNewOrUpdateResponseStatus = "error"
	NamespaceInstanceItemNewOrUpdateResponseStatusSkipped   NamespaceInstanceItemNewOrUpdateResponseStatus = "skipped"
	NamespaceInstanceItemNewOrUpdateResponseStatusOutdated  NamespaceInstanceItemNewOrUpdateResponseStatus = "outdated"
)

func (r NamespaceInstanceItemNewOrUpdateResponseStatus) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemNewOrUpdateResponseStatusQueued, NamespaceInstanceItemNewOrUpdateResponseStatusRunning, NamespaceInstanceItemNewOrUpdateResponseStatusCompleted, NamespaceInstanceItemNewOrUpdateResponseStatusError, NamespaceInstanceItemNewOrUpdateResponseStatusSkipped, NamespaceInstanceItemNewOrUpdateResponseStatusOutdated:
		return true
	}
	return false
}

type NamespaceInstanceItemGetResponse struct {
	ID          string                                     `json:"id" api:"required"`
	Checksum    string                                     `json:"checksum" api:"required"`
	ChunksCount int64                                      `json:"chunks_count" api:"required,nullable"`
	CreatedAt   time.Time                                  `json:"created_at" api:"required" format:"date-time"`
	FileSize    float64                                    `json:"file_size" api:"required,nullable"`
	Key         string                                     `json:"key" api:"required"`
	LastSeenAt  time.Time                                  `json:"last_seen_at" api:"required" format:"date-time"`
	Namespace   string                                     `json:"namespace" api:"required"`
	NextAction  NamespaceInstanceItemGetResponseNextAction `json:"next_action" api:"required,nullable"`
	// Identifies which data source this item belongs to. "builtin" for uploaded files,
	// "{type}:{source}" for external sources, null for legacy items.
	SourceID string                                 `json:"source_id" api:"required,nullable"`
	Status   NamespaceInstanceItemGetResponseStatus `json:"status" api:"required"`
	Error    string                                 `json:"error"`
	JSON     namespaceInstanceItemGetResponseJSON   `json:"-"`
}

// namespaceInstanceItemGetResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceItemGetResponse]
type namespaceInstanceItemGetResponseJSON struct {
	ID          apijson.Field
	Checksum    apijson.Field
	ChunksCount apijson.Field
	CreatedAt   apijson.Field
	FileSize    apijson.Field
	Key         apijson.Field
	LastSeenAt  apijson.Field
	Namespace   apijson.Field
	NextAction  apijson.Field
	SourceID    apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemGetResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemGetResponseNextAction string

const (
	NamespaceInstanceItemGetResponseNextActionIndex  NamespaceInstanceItemGetResponseNextAction = "INDEX"
	NamespaceInstanceItemGetResponseNextActionDelete NamespaceInstanceItemGetResponseNextAction = "DELETE"
)

func (r NamespaceInstanceItemGetResponseNextAction) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemGetResponseNextActionIndex, NamespaceInstanceItemGetResponseNextActionDelete:
		return true
	}
	return false
}

type NamespaceInstanceItemGetResponseStatus string

const (
	NamespaceInstanceItemGetResponseStatusQueued    NamespaceInstanceItemGetResponseStatus = "queued"
	NamespaceInstanceItemGetResponseStatusRunning   NamespaceInstanceItemGetResponseStatus = "running"
	NamespaceInstanceItemGetResponseStatusCompleted NamespaceInstanceItemGetResponseStatus = "completed"
	NamespaceInstanceItemGetResponseStatusError     NamespaceInstanceItemGetResponseStatus = "error"
	NamespaceInstanceItemGetResponseStatusSkipped   NamespaceInstanceItemGetResponseStatus = "skipped"
	NamespaceInstanceItemGetResponseStatusOutdated  NamespaceInstanceItemGetResponseStatus = "outdated"
)

func (r NamespaceInstanceItemGetResponseStatus) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemGetResponseStatusQueued, NamespaceInstanceItemGetResponseStatusRunning, NamespaceInstanceItemGetResponseStatusCompleted, NamespaceInstanceItemGetResponseStatusError, NamespaceInstanceItemGetResponseStatusSkipped, NamespaceInstanceItemGetResponseStatusOutdated:
		return true
	}
	return false
}

type NamespaceInstanceItemLogsResponse struct {
	Action           string                                `json:"action" api:"required"`
	ChunkCount       int64                                 `json:"chunkCount" api:"required,nullable"`
	ErrorType        string                                `json:"errorType" api:"required,nullable"`
	FileKey          string                                `json:"fileKey" api:"required"`
	Message          string                                `json:"message" api:"required,nullable"`
	ProcessingTimeMs int64                                 `json:"processingTimeMs" api:"required,nullable"`
	Timestamp        time.Time                             `json:"timestamp" api:"required" format:"date-time"`
	JSON             namespaceInstanceItemLogsResponseJSON `json:"-"`
}

// namespaceInstanceItemLogsResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceItemLogsResponse]
type namespaceInstanceItemLogsResponseJSON struct {
	Action           apijson.Field
	ChunkCount       apijson.Field
	ErrorType        apijson.Field
	FileKey          apijson.Field
	Message          apijson.Field
	ProcessingTimeMs apijson.Field
	Timestamp        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NamespaceInstanceItemLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemLogsResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemSyncResponse struct {
	ID          string                                      `json:"id" api:"required"`
	Checksum    string                                      `json:"checksum" api:"required"`
	ChunksCount int64                                       `json:"chunks_count" api:"required,nullable"`
	CreatedAt   time.Time                                   `json:"created_at" api:"required" format:"date-time"`
	FileSize    float64                                     `json:"file_size" api:"required,nullable"`
	Key         string                                      `json:"key" api:"required"`
	LastSeenAt  time.Time                                   `json:"last_seen_at" api:"required" format:"date-time"`
	Namespace   string                                      `json:"namespace" api:"required"`
	NextAction  NamespaceInstanceItemSyncResponseNextAction `json:"next_action" api:"required,nullable"`
	// Identifies which data source this item belongs to. "builtin" for uploaded files,
	// "{type}:{source}" for external sources, null for legacy items.
	SourceID string                                  `json:"source_id" api:"required,nullable"`
	Status   NamespaceInstanceItemSyncResponseStatus `json:"status" api:"required"`
	Error    string                                  `json:"error"`
	JSON     namespaceInstanceItemSyncResponseJSON   `json:"-"`
}

// namespaceInstanceItemSyncResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceItemSyncResponse]
type namespaceInstanceItemSyncResponseJSON struct {
	ID          apijson.Field
	Checksum    apijson.Field
	ChunksCount apijson.Field
	CreatedAt   apijson.Field
	FileSize    apijson.Field
	Key         apijson.Field
	LastSeenAt  apijson.Field
	Namespace   apijson.Field
	NextAction  apijson.Field
	SourceID    apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemSyncResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemSyncResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemSyncResponseNextAction string

const (
	NamespaceInstanceItemSyncResponseNextActionIndex  NamespaceInstanceItemSyncResponseNextAction = "INDEX"
	NamespaceInstanceItemSyncResponseNextActionDelete NamespaceInstanceItemSyncResponseNextAction = "DELETE"
)

func (r NamespaceInstanceItemSyncResponseNextAction) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemSyncResponseNextActionIndex, NamespaceInstanceItemSyncResponseNextActionDelete:
		return true
	}
	return false
}

type NamespaceInstanceItemSyncResponseStatus string

const (
	NamespaceInstanceItemSyncResponseStatusQueued    NamespaceInstanceItemSyncResponseStatus = "queued"
	NamespaceInstanceItemSyncResponseStatusRunning   NamespaceInstanceItemSyncResponseStatus = "running"
	NamespaceInstanceItemSyncResponseStatusCompleted NamespaceInstanceItemSyncResponseStatus = "completed"
	NamespaceInstanceItemSyncResponseStatusError     NamespaceInstanceItemSyncResponseStatus = "error"
	NamespaceInstanceItemSyncResponseStatusSkipped   NamespaceInstanceItemSyncResponseStatus = "skipped"
	NamespaceInstanceItemSyncResponseStatusOutdated  NamespaceInstanceItemSyncResponseStatus = "outdated"
)

func (r NamespaceInstanceItemSyncResponseStatus) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemSyncResponseStatusQueued, NamespaceInstanceItemSyncResponseStatusRunning, NamespaceInstanceItemSyncResponseStatusCompleted, NamespaceInstanceItemSyncResponseStatusError, NamespaceInstanceItemSyncResponseStatusSkipped, NamespaceInstanceItemSyncResponseStatusOutdated:
		return true
	}
	return false
}

type NamespaceInstanceItemUploadResponse struct {
	ID          string                                        `json:"id" api:"required"`
	Checksum    string                                        `json:"checksum" api:"required"`
	ChunksCount int64                                         `json:"chunks_count" api:"required,nullable"`
	CreatedAt   time.Time                                     `json:"created_at" api:"required" format:"date-time"`
	FileSize    float64                                       `json:"file_size" api:"required,nullable"`
	Key         string                                        `json:"key" api:"required"`
	LastSeenAt  time.Time                                     `json:"last_seen_at" api:"required" format:"date-time"`
	Namespace   string                                        `json:"namespace" api:"required"`
	NextAction  NamespaceInstanceItemUploadResponseNextAction `json:"next_action" api:"required,nullable"`
	// Identifies which data source this item belongs to. "builtin" for uploaded files,
	// "{type}:{source}" for external sources, null for legacy items.
	SourceID string                                    `json:"source_id" api:"required,nullable"`
	Status   NamespaceInstanceItemUploadResponseStatus `json:"status" api:"required"`
	Error    string                                    `json:"error"`
	JSON     namespaceInstanceItemUploadResponseJSON   `json:"-"`
}

// namespaceInstanceItemUploadResponseJSON contains the JSON metadata for the
// struct [NamespaceInstanceItemUploadResponse]
type namespaceInstanceItemUploadResponseJSON struct {
	ID          apijson.Field
	Checksum    apijson.Field
	ChunksCount apijson.Field
	CreatedAt   apijson.Field
	FileSize    apijson.Field
	Key         apijson.Field
	LastSeenAt  apijson.Field
	Namespace   apijson.Field
	NextAction  apijson.Field
	SourceID    apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemUploadResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemUploadResponseNextAction string

const (
	NamespaceInstanceItemUploadResponseNextActionIndex  NamespaceInstanceItemUploadResponseNextAction = "INDEX"
	NamespaceInstanceItemUploadResponseNextActionDelete NamespaceInstanceItemUploadResponseNextAction = "DELETE"
)

func (r NamespaceInstanceItemUploadResponseNextAction) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemUploadResponseNextActionIndex, NamespaceInstanceItemUploadResponseNextActionDelete:
		return true
	}
	return false
}

type NamespaceInstanceItemUploadResponseStatus string

const (
	NamespaceInstanceItemUploadResponseStatusQueued    NamespaceInstanceItemUploadResponseStatus = "queued"
	NamespaceInstanceItemUploadResponseStatusRunning   NamespaceInstanceItemUploadResponseStatus = "running"
	NamespaceInstanceItemUploadResponseStatusCompleted NamespaceInstanceItemUploadResponseStatus = "completed"
	NamespaceInstanceItemUploadResponseStatusError     NamespaceInstanceItemUploadResponseStatus = "error"
	NamespaceInstanceItemUploadResponseStatusSkipped   NamespaceInstanceItemUploadResponseStatus = "skipped"
	NamespaceInstanceItemUploadResponseStatusOutdated  NamespaceInstanceItemUploadResponseStatus = "outdated"
)

func (r NamespaceInstanceItemUploadResponseStatus) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemUploadResponseStatusQueued, NamespaceInstanceItemUploadResponseStatusRunning, NamespaceInstanceItemUploadResponseStatusCompleted, NamespaceInstanceItemUploadResponseStatusError, NamespaceInstanceItemUploadResponseStatusSkipped, NamespaceInstanceItemUploadResponseStatusOutdated:
		return true
	}
	return false
}

type NamespaceInstanceItemListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter items by their unique ID. Returns at most one item.
	ItemID param.Field[string] `query:"item_id"`
	// JSON-encoded metadata filter using Vectorize filter syntax. Examples:
	// {"folder":"reports/"},
	// {"timestamp":{"$gte":1700000000000}}, {"folder":{"$in":["docs/","reports/"]}}
	MetadataFilter param.Field[string] `query:"metadata_filter"`
	Page           param.Field[int64]  `query:"page"`
	PerPage        param.Field[int64]  `query:"per_page"`
	Search         param.Field[string] `query:"search"`
	// Sort order for items. "status" (default) sorts by status priority then
	// last_seen_at. "modified_at" sorts by file modification time (most recent first),
	// falling back to created_at.
	SortBy param.Field[NamespaceInstanceItemListParamsSortBy] `query:"sort_by"`
	// Filter items by source_id. Use "builtin" for uploaded files, or a source
	// identifier like "web-crawler:https://example.com".
	Source param.Field[string]                                `query:"source"`
	Status param.Field[NamespaceInstanceItemListParamsStatus] `query:"status"`
}

// URLQuery serializes [NamespaceInstanceItemListParams]'s query parameters as
// `url.Values`.
func (r NamespaceInstanceItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort order for items. "status" (default) sorts by status priority then
// last_seen_at. "modified_at" sorts by file modification time (most recent first),
// falling back to created_at.
type NamespaceInstanceItemListParamsSortBy string

const (
	NamespaceInstanceItemListParamsSortByStatus     NamespaceInstanceItemListParamsSortBy = "status"
	NamespaceInstanceItemListParamsSortByModifiedAt NamespaceInstanceItemListParamsSortBy = "modified_at"
)

func (r NamespaceInstanceItemListParamsSortBy) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemListParamsSortByStatus, NamespaceInstanceItemListParamsSortByModifiedAt:
		return true
	}
	return false
}

type NamespaceInstanceItemListParamsStatus string

const (
	NamespaceInstanceItemListParamsStatusQueued    NamespaceInstanceItemListParamsStatus = "queued"
	NamespaceInstanceItemListParamsStatusRunning   NamespaceInstanceItemListParamsStatus = "running"
	NamespaceInstanceItemListParamsStatusCompleted NamespaceInstanceItemListParamsStatus = "completed"
	NamespaceInstanceItemListParamsStatusError     NamespaceInstanceItemListParamsStatus = "error"
	NamespaceInstanceItemListParamsStatusSkipped   NamespaceInstanceItemListParamsStatus = "skipped"
	NamespaceInstanceItemListParamsStatusOutdated  NamespaceInstanceItemListParamsStatus = "outdated"
)

func (r NamespaceInstanceItemListParamsStatus) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemListParamsStatusQueued, NamespaceInstanceItemListParamsStatusRunning, NamespaceInstanceItemListParamsStatusCompleted, NamespaceInstanceItemListParamsStatusError, NamespaceInstanceItemListParamsStatusSkipped, NamespaceInstanceItemListParamsStatusOutdated:
		return true
	}
	return false
}

type NamespaceInstanceItemDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type NamespaceInstanceItemDeleteResponseEnvelope struct {
	Result  NamespaceInstanceItemDeleteResponse             `json:"result" api:"required"`
	Success bool                                            `json:"success" api:"required"`
	JSON    namespaceInstanceItemDeleteResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceItemDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [NamespaceInstanceItemDeleteResponseEnvelope]
type namespaceInstanceItemDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemChunksParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Limit     param.Field[int64]  `query:"limit"`
	Offset    param.Field[int64]  `query:"offset"`
}

// URLQuery serializes [NamespaceInstanceItemChunksParams]'s query parameters as
// `url.Values`.
func (r NamespaceInstanceItemChunksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NamespaceInstanceItemChunksResponseEnvelope struct {
	Result     []NamespaceInstanceItemChunksResponse                 `json:"result" api:"required"`
	ResultInfo NamespaceInstanceItemChunksResponseEnvelopeResultInfo `json:"result_info" api:"required"`
	Success    bool                                                  `json:"success" api:"required"`
	JSON       namespaceInstanceItemChunksResponseEnvelopeJSON       `json:"-"`
}

// namespaceInstanceItemChunksResponseEnvelopeJSON contains the JSON metadata for
// the struct [NamespaceInstanceItemChunksResponseEnvelope]
type namespaceInstanceItemChunksResponseEnvelopeJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemChunksResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemChunksResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemChunksResponseEnvelopeResultInfo struct {
	Count  int64                                                     `json:"count" api:"required"`
	Limit  int64                                                     `json:"limit" api:"required"`
	Offset int64                                                     `json:"offset" api:"required"`
	Total  int64                                                     `json:"total" api:"required"`
	JSON   namespaceInstanceItemChunksResponseEnvelopeResultInfoJSON `json:"-"`
}

// namespaceInstanceItemChunksResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [NamespaceInstanceItemChunksResponseEnvelopeResultInfo]
type namespaceInstanceItemChunksResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemChunksResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemChunksResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemNewOrUpdateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Item key / filename. Must not exceed 128 characters.
	Key        param.Field[string]                                           `json:"key" api:"required"`
	NextAction param.Field[NamespaceInstanceItemNewOrUpdateParamsNextAction] `json:"next_action" api:"required"`
}

func (r NamespaceInstanceItemNewOrUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceItemNewOrUpdateParamsNextAction string

const (
	NamespaceInstanceItemNewOrUpdateParamsNextActionIndex NamespaceInstanceItemNewOrUpdateParamsNextAction = "INDEX"
)

func (r NamespaceInstanceItemNewOrUpdateParamsNextAction) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemNewOrUpdateParamsNextActionIndex:
		return true
	}
	return false
}

type NamespaceInstanceItemNewOrUpdateResponseEnvelope struct {
	Result  NamespaceInstanceItemNewOrUpdateResponse             `json:"result" api:"required"`
	Success bool                                                 `json:"success" api:"required"`
	JSON    namespaceInstanceItemNewOrUpdateResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceItemNewOrUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [NamespaceInstanceItemNewOrUpdateResponseEnvelope]
type namespaceInstanceItemNewOrUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemNewOrUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemNewOrUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemDownloadParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type NamespaceInstanceItemGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type NamespaceInstanceItemGetResponseEnvelope struct {
	Result  NamespaceInstanceItemGetResponse             `json:"result" api:"required"`
	Success bool                                         `json:"success" api:"required"`
	JSON    namespaceInstanceItemGetResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceItemGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceInstanceItemGetResponseEnvelope]
type namespaceInstanceItemGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemLogsParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Cursor    param.Field[string] `query:"cursor"`
	Limit     param.Field[int64]  `query:"limit"`
}

// URLQuery serializes [NamespaceInstanceItemLogsParams]'s query parameters as
// `url.Values`.
func (r NamespaceInstanceItemLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NamespaceInstanceItemLogsResponseEnvelope struct {
	Result     []NamespaceInstanceItemLogsResponse                 `json:"result" api:"required"`
	ResultInfo NamespaceInstanceItemLogsResponseEnvelopeResultInfo `json:"result_info" api:"required"`
	Success    bool                                                `json:"success" api:"required"`
	JSON       namespaceInstanceItemLogsResponseEnvelopeJSON       `json:"-"`
}

// namespaceInstanceItemLogsResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceInstanceItemLogsResponseEnvelope]
type namespaceInstanceItemLogsResponseEnvelopeJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemLogsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemLogsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemLogsResponseEnvelopeResultInfo struct {
	Count     int64                                                   `json:"count" api:"required"`
	Cursor    string                                                  `json:"cursor" api:"required,nullable"`
	PerPage   int64                                                   `json:"per_page" api:"required"`
	Truncated bool                                                    `json:"truncated" api:"required"`
	JSON      namespaceInstanceItemLogsResponseEnvelopeResultInfoJSON `json:"-"`
}

// namespaceInstanceItemLogsResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [NamespaceInstanceItemLogsResponseEnvelopeResultInfo]
type namespaceInstanceItemLogsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	PerPage     apijson.Field
	Truncated   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemLogsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemLogsResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemSyncParams struct {
	AccountID  param.Field[string]                                    `path:"account_id" api:"required"`
	NextAction param.Field[NamespaceInstanceItemSyncParamsNextAction] `json:"next_action" api:"required"`
}

func (r NamespaceInstanceItemSyncParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceItemSyncParamsNextAction string

const (
	NamespaceInstanceItemSyncParamsNextActionIndex NamespaceInstanceItemSyncParamsNextAction = "INDEX"
)

func (r NamespaceInstanceItemSyncParamsNextAction) IsKnown() bool {
	switch r {
	case NamespaceInstanceItemSyncParamsNextActionIndex:
		return true
	}
	return false
}

type NamespaceInstanceItemSyncResponseEnvelope struct {
	Result  NamespaceInstanceItemSyncResponse             `json:"result" api:"required"`
	Success bool                                          `json:"success" api:"required"`
	JSON    namespaceInstanceItemSyncResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceItemSyncResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceInstanceItemSyncResponseEnvelope]
type namespaceInstanceItemSyncResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemSyncResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemSyncResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceItemUploadParams struct {
	AccountID param.Field[string]                   `path:"account_id" api:"required"`
	File      NamespaceInstanceItemUploadParamsFile `json:"file" api:"required"`
}

func (r NamespaceInstanceItemUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r.File, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type NamespaceInstanceItemUploadParamsFile struct {
	// The file to upload (max 4MB). Filename must not exceed 128 characters.
	File param.Field[io.Reader] `json:"file" api:"required" format:"binary"`
	// JSON string of custom metadata key-value pairs.
	Metadata param.Field[string] `json:"metadata"`
	// Wait for indexing to complete before responding. Defaults to false.
	WaitForCompletion param.Field[bool] `json:"wait_for_completion"`
}

func (r NamespaceInstanceItemUploadParamsFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceItemUploadResponseEnvelope struct {
	Result  NamespaceInstanceItemUploadResponse             `json:"result" api:"required"`
	Success bool                                            `json:"success" api:"required"`
	JSON    namespaceInstanceItemUploadResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceItemUploadResponseEnvelopeJSON contains the JSON metadata for
// the struct [NamespaceInstanceItemUploadResponseEnvelope]
type namespaceInstanceItemUploadResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceItemUploadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceItemUploadResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
