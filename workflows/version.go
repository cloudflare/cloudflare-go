// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflows

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
