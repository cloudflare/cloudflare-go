// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// NamespaceInstanceJobService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceInstanceJobService] method instead.
type NamespaceInstanceJobService struct {
	Options []option.RequestOption
}

// NewNamespaceInstanceJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNamespaceInstanceJobService(opts ...option.RequestOption) (r *NamespaceInstanceJobService) {
	r = &NamespaceInstanceJobService{}
	r.Options = opts
	return
}

// Creates a new indexing job for an AI Search instance.
func (r *NamespaceInstanceJobService) New(ctx context.Context, name string, id string, params NamespaceInstanceJobNewParams, opts ...option.RequestOption) (res *NamespaceInstanceJobNewResponse, err error) {
	var env NamespaceInstanceJobNewResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/jobs", params.AccountID, name, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates the status of an AI Search indexing job.
func (r *NamespaceInstanceJobService) Update(ctx context.Context, name string, id string, jobID string, params NamespaceInstanceJobUpdateParams, opts ...option.RequestOption) (res *NamespaceInstanceJobUpdateResponse, err error) {
	var env NamespaceInstanceJobUpdateResponseEnvelope
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
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/jobs/%s", params.AccountID, name, id, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists indexing jobs for an AI Search instance.
func (r *NamespaceInstanceJobService) List(ctx context.Context, name string, id string, params NamespaceInstanceJobListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[NamespaceInstanceJobListResponse], err error) {
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
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/jobs", params.AccountID, name, id)
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

// Lists indexing jobs for an AI Search instance.
func (r *NamespaceInstanceJobService) ListAutoPaging(ctx context.Context, name string, id string, params NamespaceInstanceJobListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[NamespaceInstanceJobListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, name, id, params, opts...))
}

// Retrieves details for a specific AI Search indexing job.
func (r *NamespaceInstanceJobService) Get(ctx context.Context, name string, id string, jobID string, query NamespaceInstanceJobGetParams, opts ...option.RequestOption) (res *NamespaceInstanceJobGetResponse, err error) {
	var env NamespaceInstanceJobGetResponseEnvelope
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
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/jobs/%s", query.AccountID, name, id, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists log entries for an AI Search indexing job.
func (r *NamespaceInstanceJobService) Logs(ctx context.Context, name string, id string, jobID string, params NamespaceInstanceJobLogsParams, opts ...option.RequestOption) (res *[]NamespaceInstanceJobLogsResponse, err error) {
	var env NamespaceInstanceJobLogsResponseEnvelope
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
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/jobs/%s/logs", params.AccountID, name, id, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type NamespaceInstanceJobNewResponse struct {
	ID          string                                `json:"id" api:"required"`
	Source      NamespaceInstanceJobNewResponseSource `json:"source" api:"required"`
	Description string                                `json:"description"`
	EndReason   string                                `json:"end_reason"`
	EndedAt     string                                `json:"ended_at"`
	LastSeenAt  string                                `json:"last_seen_at"`
	StartedAt   string                                `json:"started_at"`
	JSON        namespaceInstanceJobNewResponseJSON   `json:"-"`
}

// namespaceInstanceJobNewResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceJobNewResponse]
type namespaceInstanceJobNewResponseJSON struct {
	ID          apijson.Field
	Source      apijson.Field
	Description apijson.Field
	EndReason   apijson.Field
	EndedAt     apijson.Field
	LastSeenAt  apijson.Field
	StartedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceJobNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceJobNewResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceJobNewResponseSource string

const (
	NamespaceInstanceJobNewResponseSourceUser     NamespaceInstanceJobNewResponseSource = "user"
	NamespaceInstanceJobNewResponseSourceSchedule NamespaceInstanceJobNewResponseSource = "schedule"
)

func (r NamespaceInstanceJobNewResponseSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceJobNewResponseSourceUser, NamespaceInstanceJobNewResponseSourceSchedule:
		return true
	}
	return false
}

type NamespaceInstanceJobUpdateResponse struct {
	ID          string                                   `json:"id" api:"required"`
	Source      NamespaceInstanceJobUpdateResponseSource `json:"source" api:"required"`
	Description string                                   `json:"description"`
	EndReason   string                                   `json:"end_reason"`
	EndedAt     string                                   `json:"ended_at"`
	LastSeenAt  string                                   `json:"last_seen_at"`
	StartedAt   string                                   `json:"started_at"`
	JSON        namespaceInstanceJobUpdateResponseJSON   `json:"-"`
}

// namespaceInstanceJobUpdateResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceJobUpdateResponse]
type namespaceInstanceJobUpdateResponseJSON struct {
	ID          apijson.Field
	Source      apijson.Field
	Description apijson.Field
	EndReason   apijson.Field
	EndedAt     apijson.Field
	LastSeenAt  apijson.Field
	StartedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceJobUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceJobUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceJobUpdateResponseSource string

const (
	NamespaceInstanceJobUpdateResponseSourceUser     NamespaceInstanceJobUpdateResponseSource = "user"
	NamespaceInstanceJobUpdateResponseSourceSchedule NamespaceInstanceJobUpdateResponseSource = "schedule"
)

func (r NamespaceInstanceJobUpdateResponseSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceJobUpdateResponseSourceUser, NamespaceInstanceJobUpdateResponseSourceSchedule:
		return true
	}
	return false
}

type NamespaceInstanceJobListResponse struct {
	ID          string                                 `json:"id" api:"required"`
	Source      NamespaceInstanceJobListResponseSource `json:"source" api:"required"`
	Description string                                 `json:"description"`
	EndReason   string                                 `json:"end_reason"`
	EndedAt     string                                 `json:"ended_at"`
	LastSeenAt  string                                 `json:"last_seen_at"`
	StartedAt   string                                 `json:"started_at"`
	JSON        namespaceInstanceJobListResponseJSON   `json:"-"`
}

// namespaceInstanceJobListResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceJobListResponse]
type namespaceInstanceJobListResponseJSON struct {
	ID          apijson.Field
	Source      apijson.Field
	Description apijson.Field
	EndReason   apijson.Field
	EndedAt     apijson.Field
	LastSeenAt  apijson.Field
	StartedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceJobListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceJobListResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceJobListResponseSource string

const (
	NamespaceInstanceJobListResponseSourceUser     NamespaceInstanceJobListResponseSource = "user"
	NamespaceInstanceJobListResponseSourceSchedule NamespaceInstanceJobListResponseSource = "schedule"
)

func (r NamespaceInstanceJobListResponseSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceJobListResponseSourceUser, NamespaceInstanceJobListResponseSourceSchedule:
		return true
	}
	return false
}

type NamespaceInstanceJobGetResponse struct {
	ID          string                                `json:"id" api:"required"`
	Source      NamespaceInstanceJobGetResponseSource `json:"source" api:"required"`
	Description string                                `json:"description"`
	EndReason   string                                `json:"end_reason"`
	EndedAt     string                                `json:"ended_at"`
	LastSeenAt  string                                `json:"last_seen_at"`
	StartedAt   string                                `json:"started_at"`
	JSON        namespaceInstanceJobGetResponseJSON   `json:"-"`
}

// namespaceInstanceJobGetResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceJobGetResponse]
type namespaceInstanceJobGetResponseJSON struct {
	ID          apijson.Field
	Source      apijson.Field
	Description apijson.Field
	EndReason   apijson.Field
	EndedAt     apijson.Field
	LastSeenAt  apijson.Field
	StartedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceJobGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceJobGetResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceJobGetResponseSource string

const (
	NamespaceInstanceJobGetResponseSourceUser     NamespaceInstanceJobGetResponseSource = "user"
	NamespaceInstanceJobGetResponseSourceSchedule NamespaceInstanceJobGetResponseSource = "schedule"
)

func (r NamespaceInstanceJobGetResponseSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceJobGetResponseSourceUser, NamespaceInstanceJobGetResponseSourceSchedule:
		return true
	}
	return false
}

type NamespaceInstanceJobLogsResponse struct {
	ID          int64                                `json:"id" api:"required"`
	CreatedAt   float64                              `json:"created_at" api:"required"`
	Message     string                               `json:"message" api:"required"`
	MessageType int64                                `json:"message_type" api:"required"`
	JSON        namespaceInstanceJobLogsResponseJSON `json:"-"`
}

// namespaceInstanceJobLogsResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceJobLogsResponse]
type namespaceInstanceJobLogsResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Message     apijson.Field
	MessageType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceJobLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceJobLogsResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceJobNewParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Description param.Field[string] `json:"description"`
}

func (r NamespaceInstanceJobNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceJobNewResponseEnvelope struct {
	Result  NamespaceInstanceJobNewResponse             `json:"result" api:"required"`
	Success bool                                        `json:"success" api:"required"`
	JSON    namespaceInstanceJobNewResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceJobNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceInstanceJobNewResponseEnvelope]
type namespaceInstanceJobNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceJobNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceJobNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceJobUpdateParams struct {
	AccountID param.Field[string]                                 `path:"account_id" api:"required"`
	Action    param.Field[NamespaceInstanceJobUpdateParamsAction] `json:"action" api:"required"`
}

func (r NamespaceInstanceJobUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceJobUpdateParamsAction string

const (
	NamespaceInstanceJobUpdateParamsActionCancel NamespaceInstanceJobUpdateParamsAction = "cancel"
)

func (r NamespaceInstanceJobUpdateParamsAction) IsKnown() bool {
	switch r {
	case NamespaceInstanceJobUpdateParamsActionCancel:
		return true
	}
	return false
}

type NamespaceInstanceJobUpdateResponseEnvelope struct {
	Result  NamespaceInstanceJobUpdateResponse             `json:"result" api:"required"`
	Success bool                                           `json:"success" api:"required"`
	JSON    namespaceInstanceJobUpdateResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceJobUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [NamespaceInstanceJobUpdateResponseEnvelope]
type namespaceInstanceJobUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceJobUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceJobUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceJobListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [NamespaceInstanceJobListParams]'s query parameters as
// `url.Values`.
func (r NamespaceInstanceJobListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NamespaceInstanceJobGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type NamespaceInstanceJobGetResponseEnvelope struct {
	Result  NamespaceInstanceJobGetResponse             `json:"result" api:"required"`
	Success bool                                        `json:"success" api:"required"`
	JSON    namespaceInstanceJobGetResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceJobGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceInstanceJobGetResponseEnvelope]
type namespaceInstanceJobGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceJobGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceJobGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceJobLogsParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [NamespaceInstanceJobLogsParams]'s query parameters as
// `url.Values`.
func (r NamespaceInstanceJobLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NamespaceInstanceJobLogsResponseEnvelope struct {
	Result     []NamespaceInstanceJobLogsResponse                 `json:"result" api:"required"`
	ResultInfo NamespaceInstanceJobLogsResponseEnvelopeResultInfo `json:"result_info" api:"required"`
	Success    bool                                               `json:"success" api:"required"`
	JSON       namespaceInstanceJobLogsResponseEnvelopeJSON       `json:"-"`
}

// namespaceInstanceJobLogsResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceInstanceJobLogsResponseEnvelope]
type namespaceInstanceJobLogsResponseEnvelopeJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceJobLogsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceJobLogsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceJobLogsResponseEnvelopeResultInfo struct {
	Count      int64                                                  `json:"count" api:"required"`
	Page       int64                                                  `json:"page" api:"required"`
	PerPage    int64                                                  `json:"per_page" api:"required"`
	TotalCount int64                                                  `json:"total_count" api:"required"`
	JSON       namespaceInstanceJobLogsResponseEnvelopeResultInfoJSON `json:"-"`
}

// namespaceInstanceJobLogsResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [NamespaceInstanceJobLogsResponseEnvelopeResultInfo]
type namespaceInstanceJobLogsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceJobLogsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceJobLogsResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
