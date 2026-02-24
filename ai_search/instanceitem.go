// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search

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

// InstanceItemService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceItemService] method instead.
type InstanceItemService struct {
	Options []option.RequestOption
}

// NewInstanceItemService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceItemService(opts ...option.RequestOption) (r *InstanceItemService) {
	r = &InstanceItemService{}
	r.Options = opts
	return
}

// Lists indexed items in an AI Search instance.
func (r *InstanceItemService) List(ctx context.Context, id string, params InstanceItemListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[InstanceItemListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s/items", params.AccountID, id)
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
func (r *InstanceItemService) ListAutoPaging(ctx context.Context, id string, params InstanceItemListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[InstanceItemListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, id, params, opts...))
}

// Retrieves a specific indexed item from an AI Search instance.
func (r *InstanceItemService) Get(ctx context.Context, id string, itemID string, query InstanceItemGetParams, opts ...option.RequestOption) (res *InstanceItemGetResponse, err error) {
	var env InstanceItemGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s/items/%s", query.AccountID, id, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InstanceItemListResponse struct {
	ID          float64                            `json:"id,required"`
	Checksum    string                             `json:"checksum,required"`
	ChunksCount int64                              `json:"chunks_count,required,nullable"`
	CreatedAt   time.Time                          `json:"created_at,required" format:"date-time"`
	FileSize    float64                            `json:"file_size,required,nullable"`
	Key         string                             `json:"key,required"`
	LastSeenAt  time.Time                          `json:"last_seen_at,required" format:"date-time"`
	Namespace   string                             `json:"namespace,required"`
	NextAction  InstanceItemListResponseNextAction `json:"next_action,required,nullable"`
	Status      InstanceItemListResponseStatus     `json:"status,required"`
	Error       string                             `json:"error"`
	PublicID    string                             `json:"public_id"`
	JSON        instanceItemListResponseJSON       `json:"-"`
}

// instanceItemListResponseJSON contains the JSON metadata for the struct
// [InstanceItemListResponse]
type instanceItemListResponseJSON struct {
	ID          apijson.Field
	Checksum    apijson.Field
	ChunksCount apijson.Field
	CreatedAt   apijson.Field
	FileSize    apijson.Field
	Key         apijson.Field
	LastSeenAt  apijson.Field
	Namespace   apijson.Field
	NextAction  apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	PublicID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceItemListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceItemListResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceItemListResponseNextAction string

const (
	InstanceItemListResponseNextActionIndex  InstanceItemListResponseNextAction = "INDEX"
	InstanceItemListResponseNextActionDelete InstanceItemListResponseNextAction = "DELETE"
)

func (r InstanceItemListResponseNextAction) IsKnown() bool {
	switch r {
	case InstanceItemListResponseNextActionIndex, InstanceItemListResponseNextActionDelete:
		return true
	}
	return false
}

type InstanceItemListResponseStatus string

const (
	InstanceItemListResponseStatusQueued    InstanceItemListResponseStatus = "queued"
	InstanceItemListResponseStatusRunning   InstanceItemListResponseStatus = "running"
	InstanceItemListResponseStatusCompleted InstanceItemListResponseStatus = "completed"
	InstanceItemListResponseStatusError     InstanceItemListResponseStatus = "error"
	InstanceItemListResponseStatusSkipped   InstanceItemListResponseStatus = "skipped"
)

func (r InstanceItemListResponseStatus) IsKnown() bool {
	switch r {
	case InstanceItemListResponseStatusQueued, InstanceItemListResponseStatusRunning, InstanceItemListResponseStatusCompleted, InstanceItemListResponseStatusError, InstanceItemListResponseStatusSkipped:
		return true
	}
	return false
}

type InstanceItemGetResponse struct {
	ID          float64                           `json:"id,required"`
	Checksum    string                            `json:"checksum,required"`
	ChunksCount int64                             `json:"chunks_count,required,nullable"`
	CreatedAt   time.Time                         `json:"created_at,required" format:"date-time"`
	FileSize    float64                           `json:"file_size,required,nullable"`
	Key         string                            `json:"key,required"`
	LastSeenAt  time.Time                         `json:"last_seen_at,required" format:"date-time"`
	Namespace   string                            `json:"namespace,required"`
	NextAction  InstanceItemGetResponseNextAction `json:"next_action,required,nullable"`
	Status      InstanceItemGetResponseStatus     `json:"status,required"`
	Error       string                            `json:"error"`
	PublicID    string                            `json:"public_id"`
	JSON        instanceItemGetResponseJSON       `json:"-"`
}

// instanceItemGetResponseJSON contains the JSON metadata for the struct
// [InstanceItemGetResponse]
type instanceItemGetResponseJSON struct {
	ID          apijson.Field
	Checksum    apijson.Field
	ChunksCount apijson.Field
	CreatedAt   apijson.Field
	FileSize    apijson.Field
	Key         apijson.Field
	LastSeenAt  apijson.Field
	Namespace   apijson.Field
	NextAction  apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	PublicID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceItemGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceItemGetResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceItemGetResponseNextAction string

const (
	InstanceItemGetResponseNextActionIndex  InstanceItemGetResponseNextAction = "INDEX"
	InstanceItemGetResponseNextActionDelete InstanceItemGetResponseNextAction = "DELETE"
)

func (r InstanceItemGetResponseNextAction) IsKnown() bool {
	switch r {
	case InstanceItemGetResponseNextActionIndex, InstanceItemGetResponseNextActionDelete:
		return true
	}
	return false
}

type InstanceItemGetResponseStatus string

const (
	InstanceItemGetResponseStatusQueued    InstanceItemGetResponseStatus = "queued"
	InstanceItemGetResponseStatusRunning   InstanceItemGetResponseStatus = "running"
	InstanceItemGetResponseStatusCompleted InstanceItemGetResponseStatus = "completed"
	InstanceItemGetResponseStatusError     InstanceItemGetResponseStatus = "error"
	InstanceItemGetResponseStatusSkipped   InstanceItemGetResponseStatus = "skipped"
)

func (r InstanceItemGetResponseStatus) IsKnown() bool {
	switch r {
	case InstanceItemGetResponseStatusQueued, InstanceItemGetResponseStatusRunning, InstanceItemGetResponseStatusCompleted, InstanceItemGetResponseStatusError, InstanceItemGetResponseStatusSkipped:
		return true
	}
	return false
}

type InstanceItemListParams struct {
	AccountID param.Field[string]                       `path:"account_id,required"`
	Page      param.Field[int64]                        `query:"page"`
	PerPage   param.Field[int64]                        `query:"per_page"`
	Search    param.Field[string]                       `query:"search"`
	Status    param.Field[InstanceItemListParamsStatus] `query:"status"`
}

// URLQuery serializes [InstanceItemListParams]'s query parameters as `url.Values`.
func (r InstanceItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceItemListParamsStatus string

const (
	InstanceItemListParamsStatusQueued    InstanceItemListParamsStatus = "queued"
	InstanceItemListParamsStatusRunning   InstanceItemListParamsStatus = "running"
	InstanceItemListParamsStatusCompleted InstanceItemListParamsStatus = "completed"
	InstanceItemListParamsStatusError     InstanceItemListParamsStatus = "error"
	InstanceItemListParamsStatusSkipped   InstanceItemListParamsStatus = "skipped"
)

func (r InstanceItemListParamsStatus) IsKnown() bool {
	switch r {
	case InstanceItemListParamsStatusQueued, InstanceItemListParamsStatusRunning, InstanceItemListParamsStatusCompleted, InstanceItemListParamsStatusError, InstanceItemListParamsStatusSkipped:
		return true
	}
	return false
}

type InstanceItemGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type InstanceItemGetResponseEnvelope struct {
	Result  InstanceItemGetResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    instanceItemGetResponseEnvelopeJSON `json:"-"`
}

// instanceItemGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceItemGetResponseEnvelope]
type instanceItemGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceItemGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceItemGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
