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

// InstanceJobService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceJobService] method instead.
type InstanceJobService struct {
	Options []option.RequestOption
}

// NewInstanceJobService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceJobService(opts ...option.RequestOption) (r *InstanceJobService) {
	r = &InstanceJobService{}
	r.Options = opts
	return
}

// Creates a new indexing job for an AI Search instance.
func (r *InstanceJobService) New(ctx context.Context, id string, body InstanceJobNewParams, opts ...option.RequestOption) (res *InstanceJobNewResponse, err error) {
	var env InstanceJobNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s/jobs", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists indexing jobs for an AI Search instance.
func (r *InstanceJobService) List(ctx context.Context, id string, params InstanceJobListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[InstanceJobListResponse], err error) {
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
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s/jobs", params.AccountID, id)
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
func (r *InstanceJobService) ListAutoPaging(ctx context.Context, id string, params InstanceJobListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[InstanceJobListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, id, params, opts...))
}

// Retrieves details for a specific AI Search indexing job.
func (r *InstanceJobService) Get(ctx context.Context, id string, jobID string, query InstanceJobGetParams, opts ...option.RequestOption) (res *InstanceJobGetResponse, err error) {
	var env InstanceJobGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s/jobs/%s", query.AccountID, id, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists log entries for an AI Search indexing job.
func (r *InstanceJobService) Logs(ctx context.Context, id string, jobID string, params InstanceJobLogsParams, opts ...option.RequestOption) (res *[]InstanceJobLogsResponse, err error) {
	var env InstanceJobLogsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s/jobs/%s/logs", params.AccountID, id, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InstanceJobNewResponse struct {
	ID         string                       `json:"id,required"`
	Source     InstanceJobNewResponseSource `json:"source,required"`
	EndReason  string                       `json:"end_reason"`
	EndedAt    string                       `json:"ended_at"`
	LastSeenAt string                       `json:"last_seen_at"`
	StartedAt  string                       `json:"started_at"`
	JSON       instanceJobNewResponseJSON   `json:"-"`
}

// instanceJobNewResponseJSON contains the JSON metadata for the struct
// [InstanceJobNewResponse]
type instanceJobNewResponseJSON struct {
	ID          apijson.Field
	Source      apijson.Field
	EndReason   apijson.Field
	EndedAt     apijson.Field
	LastSeenAt  apijson.Field
	StartedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceJobNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceJobNewResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceJobNewResponseSource string

const (
	InstanceJobNewResponseSourceUser     InstanceJobNewResponseSource = "user"
	InstanceJobNewResponseSourceSchedule InstanceJobNewResponseSource = "schedule"
)

func (r InstanceJobNewResponseSource) IsKnown() bool {
	switch r {
	case InstanceJobNewResponseSourceUser, InstanceJobNewResponseSourceSchedule:
		return true
	}
	return false
}

type InstanceJobListResponse struct {
	ID         string                        `json:"id,required"`
	Source     InstanceJobListResponseSource `json:"source,required"`
	EndReason  string                        `json:"end_reason"`
	EndedAt    string                        `json:"ended_at"`
	LastSeenAt string                        `json:"last_seen_at"`
	StartedAt  string                        `json:"started_at"`
	JSON       instanceJobListResponseJSON   `json:"-"`
}

// instanceJobListResponseJSON contains the JSON metadata for the struct
// [InstanceJobListResponse]
type instanceJobListResponseJSON struct {
	ID          apijson.Field
	Source      apijson.Field
	EndReason   apijson.Field
	EndedAt     apijson.Field
	LastSeenAt  apijson.Field
	StartedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceJobListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceJobListResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceJobListResponseSource string

const (
	InstanceJobListResponseSourceUser     InstanceJobListResponseSource = "user"
	InstanceJobListResponseSourceSchedule InstanceJobListResponseSource = "schedule"
)

func (r InstanceJobListResponseSource) IsKnown() bool {
	switch r {
	case InstanceJobListResponseSourceUser, InstanceJobListResponseSourceSchedule:
		return true
	}
	return false
}

type InstanceJobGetResponse struct {
	ID         string                       `json:"id,required"`
	Source     InstanceJobGetResponseSource `json:"source,required"`
	EndReason  string                       `json:"end_reason"`
	EndedAt    string                       `json:"ended_at"`
	LastSeenAt string                       `json:"last_seen_at"`
	StartedAt  string                       `json:"started_at"`
	JSON       instanceJobGetResponseJSON   `json:"-"`
}

// instanceJobGetResponseJSON contains the JSON metadata for the struct
// [InstanceJobGetResponse]
type instanceJobGetResponseJSON struct {
	ID          apijson.Field
	Source      apijson.Field
	EndReason   apijson.Field
	EndedAt     apijson.Field
	LastSeenAt  apijson.Field
	StartedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceJobGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceJobGetResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceJobGetResponseSource string

const (
	InstanceJobGetResponseSourceUser     InstanceJobGetResponseSource = "user"
	InstanceJobGetResponseSourceSchedule InstanceJobGetResponseSource = "schedule"
)

func (r InstanceJobGetResponseSource) IsKnown() bool {
	switch r {
	case InstanceJobGetResponseSourceUser, InstanceJobGetResponseSourceSchedule:
		return true
	}
	return false
}

type InstanceJobLogsResponse struct {
	ID          int64                       `json:"id,required"`
	CreatedAt   float64                     `json:"created_at,required"`
	Message     string                      `json:"message,required"`
	MessageType int64                       `json:"message_type,required"`
	JSON        instanceJobLogsResponseJSON `json:"-"`
}

// instanceJobLogsResponseJSON contains the JSON metadata for the struct
// [InstanceJobLogsResponse]
type instanceJobLogsResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Message     apijson.Field
	MessageType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceJobLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceJobLogsResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceJobNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type InstanceJobNewResponseEnvelope struct {
	Result  InstanceJobNewResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    instanceJobNewResponseEnvelopeJSON `json:"-"`
}

// instanceJobNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceJobNewResponseEnvelope]
type instanceJobNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceJobNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceJobNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InstanceJobListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [InstanceJobListParams]'s query parameters as `url.Values`.
func (r InstanceJobListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceJobGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type InstanceJobGetResponseEnvelope struct {
	Result  InstanceJobGetResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    instanceJobGetResponseEnvelopeJSON `json:"-"`
}

// instanceJobGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceJobGetResponseEnvelope]
type instanceJobGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceJobGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceJobGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InstanceJobLogsParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [InstanceJobLogsParams]'s query parameters as `url.Values`.
func (r InstanceJobLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceJobLogsResponseEnvelope struct {
	Result     []InstanceJobLogsResponse                 `json:"result,required"`
	ResultInfo InstanceJobLogsResponseEnvelopeResultInfo `json:"result_info,required"`
	Success    bool                                      `json:"success,required"`
	JSON       instanceJobLogsResponseEnvelopeJSON       `json:"-"`
}

// instanceJobLogsResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceJobLogsResponseEnvelope]
type instanceJobLogsResponseEnvelopeJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceJobLogsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceJobLogsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InstanceJobLogsResponseEnvelopeResultInfo struct {
	Count      int64                                         `json:"count,required"`
	Page       int64                                         `json:"page,required"`
	PerPage    int64                                         `json:"per_page,required"`
	TotalCount int64                                         `json:"total_count,required"`
	JSON       instanceJobLogsResponseEnvelopeResultInfoJSON `json:"-"`
}

// instanceJobLogsResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [InstanceJobLogsResponseEnvelopeResultInfo]
type instanceJobLogsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceJobLogsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceJobLogsResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
