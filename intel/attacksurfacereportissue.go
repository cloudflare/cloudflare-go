// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// AttackSurfaceReportIssueService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAttackSurfaceReportIssueService] method instead.
type AttackSurfaceReportIssueService struct {
	Options []option.RequestOption
}

// NewAttackSurfaceReportIssueService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAttackSurfaceReportIssueService(opts ...option.RequestOption) (r *AttackSurfaceReportIssueService) {
	r = &AttackSurfaceReportIssueService{}
	r.Options = opts
	return
}

// Get Security Center Issues
func (r *AttackSurfaceReportIssueService) List(ctx context.Context, params AttackSurfaceReportIssueListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[AttackSurfaceReportIssueListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issues", params.AccountID)
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

// Get Security Center Issues
func (r *AttackSurfaceReportIssueService) ListAutoPaging(ctx context.Context, params AttackSurfaceReportIssueListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[AttackSurfaceReportIssueListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

// Get Security Center Issue Counts by Class
func (r *AttackSurfaceReportIssueService) Class(ctx context.Context, params AttackSurfaceReportIssueClassParams, opts ...option.RequestOption) (res *[]AttackSurfaceReportIssueClassResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackSurfaceReportIssueClassResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issues/class", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Archive Security Center Insight
func (r *AttackSurfaceReportIssueService) Dismiss(ctx context.Context, issueID string, params AttackSurfaceReportIssueDismissParams, opts ...option.RequestOption) (res *AttackSurfaceReportIssueDismissResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackSurfaceReportIssueDismissResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/%s/dismiss", params.AccountID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Security Center Issue Counts by Severity
func (r *AttackSurfaceReportIssueService) Severity(ctx context.Context, params AttackSurfaceReportIssueSeverityParams, opts ...option.RequestOption) (res *[]AttackSurfaceReportIssueSeverityResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackSurfaceReportIssueSeverityResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issues/severity", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Security Center Issue Counts by Type
func (r *AttackSurfaceReportIssueService) Type(ctx context.Context, params AttackSurfaceReportIssueTypeParams, opts ...option.RequestOption) (res *[]AttackSurfaceReportIssueTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackSurfaceReportIssueTypeResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issues/type", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IssueType string

const (
	IssueTypeComplianceViolation   IssueType = "compliance_violation"
	IssueTypeEmailSecurity         IssueType = "email_security"
	IssueTypeExposedInfrastructure IssueType = "exposed_infrastructure"
	IssueTypeInsecureConfiguration IssueType = "insecure_configuration"
	IssueTypeWeakAuthentication    IssueType = "weak_authentication"
)

func (r IssueType) IsKnown() bool {
	switch r {
	case IssueTypeComplianceViolation, IssueTypeEmailSecurity, IssueTypeExposedInfrastructure, IssueTypeInsecureConfiguration, IssueTypeWeakAuthentication:
		return true
	}
	return false
}

type SeverityQueryParam string

const (
	SeverityQueryParamLow      SeverityQueryParam = "low"
	SeverityQueryParamModerate SeverityQueryParam = "moderate"
	SeverityQueryParamCritical SeverityQueryParam = "critical"
)

func (r SeverityQueryParam) IsKnown() bool {
	switch r {
	case SeverityQueryParamLow, SeverityQueryParamModerate, SeverityQueryParamCritical:
		return true
	}
	return false
}

type AttackSurfaceReportIssueListResponse struct {
	Errors   []shared.ResponseInfo                      `json:"errors,required"`
	Messages []shared.ResponseInfo                      `json:"messages,required"`
	Result   AttackSurfaceReportIssueListResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AttackSurfaceReportIssueListResponseSuccess `json:"success,required"`
	JSON    attackSurfaceReportIssueListResponseJSON    `json:"-"`
}

// attackSurfaceReportIssueListResponseJSON contains the JSON metadata for the
// struct [AttackSurfaceReportIssueListResponse]
type attackSurfaceReportIssueListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueListResponseJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueListResponseResult struct {
	// Total number of results
	Count  int64                                             `json:"count"`
	Issues []AttackSurfaceReportIssueListResponseResultIssue `json:"issues"`
	// Current page within paginated list of results
	Page int64 `json:"page"`
	// Number of results per page of results
	PerPage int64                                          `json:"per_page"`
	JSON    attackSurfaceReportIssueListResponseResultJSON `json:"-"`
}

// attackSurfaceReportIssueListResponseResultJSON contains the JSON metadata for
// the struct [AttackSurfaceReportIssueListResponseResult]
type attackSurfaceReportIssueListResponseResultJSON struct {
	Count       apijson.Field
	Issues      apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueListResponseResultIssue struct {
	ID          string                                                    `json:"id"`
	Dismissed   bool                                                      `json:"dismissed"`
	IssueClass  string                                                    `json:"issue_class"`
	IssueType   AttackSurfaceReportIssueListResponseResultIssuesIssueType `json:"issue_type"`
	Payload     interface{}                                               `json:"payload"`
	ResolveLink string                                                    `json:"resolve_link"`
	ResolveText string                                                    `json:"resolve_text"`
	Severity    AttackSurfaceReportIssueListResponseResultIssuesSeverity  `json:"severity"`
	Since       time.Time                                                 `json:"since" format:"date-time"`
	Subject     string                                                    `json:"subject"`
	Timestamp   time.Time                                                 `json:"timestamp" format:"date-time"`
	JSON        attackSurfaceReportIssueListResponseResultIssueJSON       `json:"-"`
}

// attackSurfaceReportIssueListResponseResultIssueJSON contains the JSON metadata
// for the struct [AttackSurfaceReportIssueListResponseResultIssue]
type attackSurfaceReportIssueListResponseResultIssueJSON struct {
	ID          apijson.Field
	Dismissed   apijson.Field
	IssueClass  apijson.Field
	IssueType   apijson.Field
	Payload     apijson.Field
	ResolveLink apijson.Field
	ResolveText apijson.Field
	Severity    apijson.Field
	Since       apijson.Field
	Subject     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueListResponseResultIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueListResponseResultIssueJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueListResponseResultIssuesIssueType string

const (
	AttackSurfaceReportIssueListResponseResultIssuesIssueTypeComplianceViolation   AttackSurfaceReportIssueListResponseResultIssuesIssueType = "compliance_violation"
	AttackSurfaceReportIssueListResponseResultIssuesIssueTypeEmailSecurity         AttackSurfaceReportIssueListResponseResultIssuesIssueType = "email_security"
	AttackSurfaceReportIssueListResponseResultIssuesIssueTypeExposedInfrastructure AttackSurfaceReportIssueListResponseResultIssuesIssueType = "exposed_infrastructure"
	AttackSurfaceReportIssueListResponseResultIssuesIssueTypeInsecureConfiguration AttackSurfaceReportIssueListResponseResultIssuesIssueType = "insecure_configuration"
	AttackSurfaceReportIssueListResponseResultIssuesIssueTypeWeakAuthentication    AttackSurfaceReportIssueListResponseResultIssuesIssueType = "weak_authentication"
)

func (r AttackSurfaceReportIssueListResponseResultIssuesIssueType) IsKnown() bool {
	switch r {
	case AttackSurfaceReportIssueListResponseResultIssuesIssueTypeComplianceViolation, AttackSurfaceReportIssueListResponseResultIssuesIssueTypeEmailSecurity, AttackSurfaceReportIssueListResponseResultIssuesIssueTypeExposedInfrastructure, AttackSurfaceReportIssueListResponseResultIssuesIssueTypeInsecureConfiguration, AttackSurfaceReportIssueListResponseResultIssuesIssueTypeWeakAuthentication:
		return true
	}
	return false
}

type AttackSurfaceReportIssueListResponseResultIssuesSeverity string

const (
	AttackSurfaceReportIssueListResponseResultIssuesSeverityLow      AttackSurfaceReportIssueListResponseResultIssuesSeverity = "Low"
	AttackSurfaceReportIssueListResponseResultIssuesSeverityModerate AttackSurfaceReportIssueListResponseResultIssuesSeverity = "Moderate"
	AttackSurfaceReportIssueListResponseResultIssuesSeverityCritical AttackSurfaceReportIssueListResponseResultIssuesSeverity = "Critical"
)

func (r AttackSurfaceReportIssueListResponseResultIssuesSeverity) IsKnown() bool {
	switch r {
	case AttackSurfaceReportIssueListResponseResultIssuesSeverityLow, AttackSurfaceReportIssueListResponseResultIssuesSeverityModerate, AttackSurfaceReportIssueListResponseResultIssuesSeverityCritical:
		return true
	}
	return false
}

// Whether the API call was successful
type AttackSurfaceReportIssueListResponseSuccess bool

const (
	AttackSurfaceReportIssueListResponseSuccessTrue AttackSurfaceReportIssueListResponseSuccess = true
)

func (r AttackSurfaceReportIssueListResponseSuccess) IsKnown() bool {
	switch r {
	case AttackSurfaceReportIssueListResponseSuccessTrue:
		return true
	}
	return false
}

type AttackSurfaceReportIssueClassResponse struct {
	Count int64                                     `json:"count"`
	Value string                                    `json:"value"`
	JSON  attackSurfaceReportIssueClassResponseJSON `json:"-"`
}

// attackSurfaceReportIssueClassResponseJSON contains the JSON metadata for the
// struct [AttackSurfaceReportIssueClassResponse]
type attackSurfaceReportIssueClassResponseJSON struct {
	Count       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueClassResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueClassResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [intel.AttackSurfaceReportIssueDismissResponseUnknown] or
// [shared.UnionString].
type AttackSurfaceReportIssueDismissResponseUnion interface {
	ImplementsIntelAttackSurfaceReportIssueDismissResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AttackSurfaceReportIssueDismissResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AttackSurfaceReportIssueSeverityResponse struct {
	Count int64                                        `json:"count"`
	Value string                                       `json:"value"`
	JSON  attackSurfaceReportIssueSeverityResponseJSON `json:"-"`
}

// attackSurfaceReportIssueSeverityResponseJSON contains the JSON metadata for the
// struct [AttackSurfaceReportIssueSeverityResponse]
type attackSurfaceReportIssueSeverityResponseJSON struct {
	Count       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueSeverityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueSeverityResponseJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueTypeResponse struct {
	Count int64                                    `json:"count"`
	Value string                                   `json:"value"`
	JSON  attackSurfaceReportIssueTypeResponseJSON `json:"-"`
}

// attackSurfaceReportIssueTypeResponseJSON contains the JSON metadata for the
// struct [AttackSurfaceReportIssueTypeResponse]
type attackSurfaceReportIssueTypeResponseJSON struct {
	Count       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueTypeResponseJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueListParams struct {
	// Identifier
	AccountID     param.Field[string]      `path:"account_id,required"`
	Dismissed     param.Field[bool]        `query:"dismissed"`
	IssueClass    param.Field[[]string]    `query:"issue_class"`
	IssueClassNeq param.Field[[]string]    `query:"issue_class~neq"`
	IssueType     param.Field[[]IssueType] `query:"issue_type"`
	IssueTypeNeq  param.Field[[]IssueType] `query:"issue_type~neq"`
	// Current page within paginated list of results
	Page param.Field[int64] `query:"page"`
	// Number of results per page of results
	PerPage     param.Field[int64]                `query:"per_page"`
	Product     param.Field[[]string]             `query:"product"`
	ProductNeq  param.Field[[]string]             `query:"product~neq"`
	Severity    param.Field[[]SeverityQueryParam] `query:"severity"`
	SeverityNeq param.Field[[]SeverityQueryParam] `query:"severity~neq"`
	Subject     param.Field[[]string]             `query:"subject"`
	SubjectNeq  param.Field[[]string]             `query:"subject~neq"`
}

// URLQuery serializes [AttackSurfaceReportIssueListParams]'s query parameters as
// `url.Values`.
func (r AttackSurfaceReportIssueListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttackSurfaceReportIssueClassParams struct {
	// Identifier
	AccountID     param.Field[string]               `path:"account_id,required"`
	Dismissed     param.Field[bool]                 `query:"dismissed"`
	IssueClass    param.Field[[]string]             `query:"issue_class"`
	IssueClassNeq param.Field[[]string]             `query:"issue_class~neq"`
	IssueType     param.Field[[]IssueType]          `query:"issue_type"`
	IssueTypeNeq  param.Field[[]IssueType]          `query:"issue_type~neq"`
	Product       param.Field[[]string]             `query:"product"`
	ProductNeq    param.Field[[]string]             `query:"product~neq"`
	Severity      param.Field[[]SeverityQueryParam] `query:"severity"`
	SeverityNeq   param.Field[[]SeverityQueryParam] `query:"severity~neq"`
	Subject       param.Field[[]string]             `query:"subject"`
	SubjectNeq    param.Field[[]string]             `query:"subject~neq"`
}

// URLQuery serializes [AttackSurfaceReportIssueClassParams]'s query parameters as
// `url.Values`.
func (r AttackSurfaceReportIssueClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttackSurfaceReportIssueClassResponseEnvelope struct {
	Errors   []shared.ResponseInfo                   `json:"errors,required"`
	Messages []shared.ResponseInfo                   `json:"messages,required"`
	Result   []AttackSurfaceReportIssueClassResponse `json:"result,required"`
	// Whether the API call was successful
	Success AttackSurfaceReportIssueClassResponseEnvelopeSuccess `json:"success,required"`
	JSON    attackSurfaceReportIssueClassResponseEnvelopeJSON    `json:"-"`
}

// attackSurfaceReportIssueClassResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackSurfaceReportIssueClassResponseEnvelope]
type attackSurfaceReportIssueClassResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueClassResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueClassResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AttackSurfaceReportIssueClassResponseEnvelopeSuccess bool

const (
	AttackSurfaceReportIssueClassResponseEnvelopeSuccessTrue AttackSurfaceReportIssueClassResponseEnvelopeSuccess = true
)

func (r AttackSurfaceReportIssueClassResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AttackSurfaceReportIssueClassResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AttackSurfaceReportIssueDismissParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Dismiss   param.Field[bool]   `json:"dismiss"`
}

func (r AttackSurfaceReportIssueDismissParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AttackSurfaceReportIssueDismissResponseEnvelope struct {
	Errors   []shared.ResponseInfo                        `json:"errors,required"`
	Messages []shared.ResponseInfo                        `json:"messages,required"`
	Result   AttackSurfaceReportIssueDismissResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success AttackSurfaceReportIssueDismissResponseEnvelopeSuccess `json:"success,required"`
	JSON    attackSurfaceReportIssueDismissResponseEnvelopeJSON    `json:"-"`
}

// attackSurfaceReportIssueDismissResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackSurfaceReportIssueDismissResponseEnvelope]
type attackSurfaceReportIssueDismissResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueDismissResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueDismissResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AttackSurfaceReportIssueDismissResponseEnvelopeSuccess bool

const (
	AttackSurfaceReportIssueDismissResponseEnvelopeSuccessTrue AttackSurfaceReportIssueDismissResponseEnvelopeSuccess = true
)

func (r AttackSurfaceReportIssueDismissResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AttackSurfaceReportIssueDismissResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AttackSurfaceReportIssueSeverityParams struct {
	// Identifier
	AccountID     param.Field[string]               `path:"account_id,required"`
	Dismissed     param.Field[bool]                 `query:"dismissed"`
	IssueClass    param.Field[[]string]             `query:"issue_class"`
	IssueClassNeq param.Field[[]string]             `query:"issue_class~neq"`
	IssueType     param.Field[[]IssueType]          `query:"issue_type"`
	IssueTypeNeq  param.Field[[]IssueType]          `query:"issue_type~neq"`
	Product       param.Field[[]string]             `query:"product"`
	ProductNeq    param.Field[[]string]             `query:"product~neq"`
	Severity      param.Field[[]SeverityQueryParam] `query:"severity"`
	SeverityNeq   param.Field[[]SeverityQueryParam] `query:"severity~neq"`
	Subject       param.Field[[]string]             `query:"subject"`
	SubjectNeq    param.Field[[]string]             `query:"subject~neq"`
}

// URLQuery serializes [AttackSurfaceReportIssueSeverityParams]'s query parameters
// as `url.Values`.
func (r AttackSurfaceReportIssueSeverityParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttackSurfaceReportIssueSeverityResponseEnvelope struct {
	Errors   []shared.ResponseInfo                      `json:"errors,required"`
	Messages []shared.ResponseInfo                      `json:"messages,required"`
	Result   []AttackSurfaceReportIssueSeverityResponse `json:"result,required"`
	// Whether the API call was successful
	Success AttackSurfaceReportIssueSeverityResponseEnvelopeSuccess `json:"success,required"`
	JSON    attackSurfaceReportIssueSeverityResponseEnvelopeJSON    `json:"-"`
}

// attackSurfaceReportIssueSeverityResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackSurfaceReportIssueSeverityResponseEnvelope]
type attackSurfaceReportIssueSeverityResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueSeverityResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueSeverityResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AttackSurfaceReportIssueSeverityResponseEnvelopeSuccess bool

const (
	AttackSurfaceReportIssueSeverityResponseEnvelopeSuccessTrue AttackSurfaceReportIssueSeverityResponseEnvelopeSuccess = true
)

func (r AttackSurfaceReportIssueSeverityResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AttackSurfaceReportIssueSeverityResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AttackSurfaceReportIssueTypeParams struct {
	// Identifier
	AccountID     param.Field[string]               `path:"account_id,required"`
	Dismissed     param.Field[bool]                 `query:"dismissed"`
	IssueClass    param.Field[[]string]             `query:"issue_class"`
	IssueClassNeq param.Field[[]string]             `query:"issue_class~neq"`
	IssueType     param.Field[[]IssueType]          `query:"issue_type"`
	IssueTypeNeq  param.Field[[]IssueType]          `query:"issue_type~neq"`
	Product       param.Field[[]string]             `query:"product"`
	ProductNeq    param.Field[[]string]             `query:"product~neq"`
	Severity      param.Field[[]SeverityQueryParam] `query:"severity"`
	SeverityNeq   param.Field[[]SeverityQueryParam] `query:"severity~neq"`
	Subject       param.Field[[]string]             `query:"subject"`
	SubjectNeq    param.Field[[]string]             `query:"subject~neq"`
}

// URLQuery serializes [AttackSurfaceReportIssueTypeParams]'s query parameters as
// `url.Values`.
func (r AttackSurfaceReportIssueTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttackSurfaceReportIssueTypeResponseEnvelope struct {
	Errors   []shared.ResponseInfo                  `json:"errors,required"`
	Messages []shared.ResponseInfo                  `json:"messages,required"`
	Result   []AttackSurfaceReportIssueTypeResponse `json:"result,required"`
	// Whether the API call was successful
	Success AttackSurfaceReportIssueTypeResponseEnvelopeSuccess `json:"success,required"`
	JSON    attackSurfaceReportIssueTypeResponseEnvelopeJSON    `json:"-"`
}

// attackSurfaceReportIssueTypeResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackSurfaceReportIssueTypeResponseEnvelope]
type attackSurfaceReportIssueTypeResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueTypeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AttackSurfaceReportIssueTypeResponseEnvelopeSuccess bool

const (
	AttackSurfaceReportIssueTypeResponseEnvelopeSuccessTrue AttackSurfaceReportIssueTypeResponseEnvelopeSuccess = true
)

func (r AttackSurfaceReportIssueTypeResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AttackSurfaceReportIssueTypeResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
