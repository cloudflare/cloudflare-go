// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflows

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

// List deployed Workflow versions
func (r *VersionService) List(ctx context.Context, workflowName string, params VersionListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[VersionListResponse], err error) {
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
func (r *VersionService) ListAutoPaging(ctx context.Context, workflowName string, params VersionListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[VersionListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, workflowName, params, opts...))
}

// Get Workflow version details
func (r *VersionService) Get(ctx context.Context, workflowName string, versionID string, query VersionGetParams, opts ...option.RequestOption) (res *VersionGetResponse, err error) {
	var env VersionGetResponseEnvelope
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

type VersionListResponse struct {
	Errors     []VersionListResponseError    `json:"errors,required"`
	Messages   []VersionListResponseMessage  `json:"messages,required"`
	Result     VersionListResponseResult     `json:"result,required"`
	Success    VersionListResponseSuccess    `json:"success,required"`
	ResultInfo VersionListResponseResultInfo `json:"result_info"`
	JSON       versionListResponseJSON       `json:"-"`
}

// versionListResponseJSON contains the JSON metadata for the struct
// [VersionListResponse]
type versionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseJSON) RawJSON() string {
	return r.raw
}

type VersionListResponseError struct {
	Code    float64                      `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    versionListResponseErrorJSON `json:"-"`
}

// versionListResponseErrorJSON contains the JSON metadata for the struct
// [VersionListResponseError]
type versionListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type VersionListResponseMessage struct {
	Code    float64                        `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    versionListResponseMessageJSON `json:"-"`
}

// versionListResponseMessageJSON contains the JSON metadata for the struct
// [VersionListResponseMessage]
type versionListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type VersionListResponseResult struct {
	ID         string                        `json:"id,required" format:"uuid"`
	ClassName  string                        `json:"class_name,required"`
	CreatedOn  time.Time                     `json:"created_on,required" format:"date-time"`
	ModifiedOn time.Time                     `json:"modified_on,required" format:"date-time"`
	WorkflowID string                        `json:"workflow_id,required" format:"uuid"`
	JSON       versionListResponseResultJSON `json:"-"`
}

// versionListResponseResultJSON contains the JSON metadata for the struct
// [VersionListResponseResult]
type versionListResponseResultJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseResultJSON) RawJSON() string {
	return r.raw
}

type VersionListResponseSuccess bool

const (
	VersionListResponseSuccessTrue VersionListResponseSuccess = true
)

func (r VersionListResponseSuccess) IsKnown() bool {
	switch r {
	case VersionListResponseSuccessTrue:
		return true
	}
	return false
}

type VersionListResponseResultInfo struct {
	Count      float64                           `json:"count,required"`
	Page       float64                           `json:"page,required"`
	PerPage    float64                           `json:"per_page,required"`
	TotalCount float64                           `json:"total_count,required"`
	JSON       versionListResponseResultInfoJSON `json:"-"`
}

// versionListResponseResultInfoJSON contains the JSON metadata for the struct
// [VersionListResponseResultInfo]
type versionListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponse struct {
	ID         string                 `json:"id,required" format:"uuid"`
	ClassName  string                 `json:"class_name,required"`
	CreatedOn  time.Time              `json:"created_on,required" format:"date-time"`
	ModifiedOn time.Time              `json:"modified_on,required" format:"date-time"`
	WorkflowID string                 `json:"workflow_id,required" format:"uuid"`
	JSON       versionGetResponseJSON `json:"-"`
}

// versionGetResponseJSON contains the JSON metadata for the struct
// [VersionGetResponse]
type versionGetResponseJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseJSON) RawJSON() string {
	return r.raw
}

type VersionListParams struct {
	AccountID param.Field[string]  `path:"account_id,required"`
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type VersionGetResponseEnvelope struct {
	Errors     []VersionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []VersionGetResponseEnvelopeMessages `json:"messages,required"`
	Result     VersionGetResponse                   `json:"result,required"`
	Success    VersionGetResponseEnvelopeSuccess    `json:"success,required"`
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
	Code    float64                              `json:"code,required"`
	Message string                               `json:"message,required"`
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
	Code    float64                                `json:"code,required"`
	Message string                                 `json:"message,required"`
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
	Count      float64                                  `json:"count,required"`
	Page       float64                                  `json:"page,required"`
	PerPage    float64                                  `json:"per_page,required"`
	TotalCount float64                                  `json:"total_count,required"`
	JSON       versionGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// versionGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [VersionGetResponseEnvelopeResultInfo]
type versionGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
