// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
)

// WorkflowVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowVersionService] method instead.
type WorkflowVersionService struct {
	Options []option.RequestOption
}

// NewWorkflowVersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowVersionService(opts ...option.RequestOption) (r *WorkflowVersionService) {
	r = &WorkflowVersionService{}
	r.Options = opts
	return
}

// List deployed Workflow versions
func (r *WorkflowVersionService) List(ctx context.Context, workflowName string, params WorkflowVersionListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[WorkflowVersionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
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

// List deployed Workflow versions
func (r *WorkflowVersionService) ListAutoPaging(ctx context.Context, workflowName string, params WorkflowVersionListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[WorkflowVersionListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, workflowName, params, opts...))
}

// Get Workflow version details
func (r *WorkflowVersionService) Get(ctx context.Context, workflowName string, versionID string, query WorkflowVersionGetParams, opts ...option.RequestOption) (res *WorkflowVersionGetResponse, err error) {
	var env WorkflowVersionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/versions/%s", query.AccountID, workflowName, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkflowVersionListResponse struct {
	Errors     []WorkflowVersionListResponseError    `json:"errors,required"`
	Messages   []WorkflowVersionListResponseMessage  `json:"messages,required"`
	Result     WorkflowVersionListResponseResult     `json:"result,required"`
	Success    WorkflowVersionListResponseSuccess    `json:"success,required"`
	ResultInfo WorkflowVersionListResponseResultInfo `json:"result_info"`
	JSON       workflowVersionListResponseJSON       `json:"-"`
}

// workflowVersionListResponseJSON contains the JSON metadata for the struct
// [WorkflowVersionListResponse]
type workflowVersionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowVersionListResponseJSON) RawJSON() string {
	return r.raw
}

type WorkflowVersionListResponseError struct {
	Code    float64                              `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    workflowVersionListResponseErrorJSON `json:"-"`
}

// workflowVersionListResponseErrorJSON contains the JSON metadata for the struct
// [WorkflowVersionListResponseError]
type workflowVersionListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowVersionListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowVersionListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type WorkflowVersionListResponseMessage struct {
	Code    float64                                `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    workflowVersionListResponseMessageJSON `json:"-"`
}

// workflowVersionListResponseMessageJSON contains the JSON metadata for the struct
// [WorkflowVersionListResponseMessage]
type workflowVersionListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowVersionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowVersionListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type WorkflowVersionListResponseResult struct {
	ID         string                                `json:"id,required" format:"uuid"`
	ClassName  string                                `json:"class_name,required"`
	CreatedOn  time.Time                             `json:"created_on,required" format:"date-time"`
	ModifiedOn time.Time                             `json:"modified_on,required" format:"date-time"`
	WorkflowID string                                `json:"workflow_id,required" format:"uuid"`
	JSON       workflowVersionListResponseResultJSON `json:"-"`
}

// workflowVersionListResponseResultJSON contains the JSON metadata for the struct
// [WorkflowVersionListResponseResult]
type workflowVersionListResponseResultJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowVersionListResponseResultJSON) RawJSON() string {
	return r.raw
}

type WorkflowVersionListResponseSuccess bool

const (
	WorkflowVersionListResponseSuccessTrue WorkflowVersionListResponseSuccess = true
)

func (r WorkflowVersionListResponseSuccess) IsKnown() bool {
	switch r {
	case WorkflowVersionListResponseSuccessTrue:
		return true
	}
	return false
}

type WorkflowVersionListResponseResultInfo struct {
	Count      float64                                   `json:"count,required"`
	Page       float64                                   `json:"page,required"`
	PerPage    float64                                   `json:"per_page,required"`
	TotalCount float64                                   `json:"total_count,required"`
	JSON       workflowVersionListResponseResultInfoJSON `json:"-"`
}

// workflowVersionListResponseResultInfoJSON contains the JSON metadata for the
// struct [WorkflowVersionListResponseResultInfo]
type workflowVersionListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowVersionListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowVersionListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type WorkflowVersionGetResponse struct {
	ID         string                         `json:"id,required" format:"uuid"`
	ClassName  string                         `json:"class_name,required"`
	CreatedOn  time.Time                      `json:"created_on,required" format:"date-time"`
	ModifiedOn time.Time                      `json:"modified_on,required" format:"date-time"`
	WorkflowID string                         `json:"workflow_id,required" format:"uuid"`
	JSON       workflowVersionGetResponseJSON `json:"-"`
}

// workflowVersionGetResponseJSON contains the JSON metadata for the struct
// [WorkflowVersionGetResponse]
type workflowVersionGetResponseJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type WorkflowVersionListParams struct {
	AccountID param.Field[string]  `path:"account_id,required"`
	Page      param.Field[float64] `query:"page"`
	PerPage   param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [WorkflowVersionListParams]'s query parameters as
// `url.Values`.
func (r WorkflowVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type WorkflowVersionGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkflowVersionGetResponseEnvelope struct {
	Errors     []WorkflowVersionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []WorkflowVersionGetResponseEnvelopeMessages `json:"messages,required"`
	Result     WorkflowVersionGetResponse                   `json:"result,required"`
	Success    WorkflowVersionGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkflowVersionGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workflowVersionGetResponseEnvelopeJSON       `json:"-"`
}

// workflowVersionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkflowVersionGetResponseEnvelope]
type workflowVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkflowVersionGetResponseEnvelopeErrors struct {
	Code    float64                                      `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    workflowVersionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workflowVersionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkflowVersionGetResponseEnvelopeErrors]
type workflowVersionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowVersionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkflowVersionGetResponseEnvelopeMessages struct {
	Code    float64                                        `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    workflowVersionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workflowVersionGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkflowVersionGetResponseEnvelopeMessages]
type workflowVersionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowVersionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type WorkflowVersionGetResponseEnvelopeSuccess bool

const (
	WorkflowVersionGetResponseEnvelopeSuccessTrue WorkflowVersionGetResponseEnvelopeSuccess = true
)

func (r WorkflowVersionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WorkflowVersionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WorkflowVersionGetResponseEnvelopeResultInfo struct {
	Count      float64                                          `json:"count,required"`
	Page       float64                                          `json:"page,required"`
	PerPage    float64                                          `json:"per_page,required"`
	TotalCount float64                                          `json:"total_count,required"`
	JSON       workflowVersionGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// workflowVersionGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [WorkflowVersionGetResponseEnvelopeResultInfo]
type workflowVersionGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowVersionGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowVersionGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
