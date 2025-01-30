// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_center

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/intel"
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// InsightService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightService] method instead.
type InsightService struct {
	Options  []option.RequestOption
	Class    *InsightClassService
	Severity *InsightSeverityService
	Type     *InsightTypeService
}

// NewInsightService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInsightService(opts ...option.RequestOption) (r *InsightService) {
	r = &InsightService{}
	r.Options = opts
	r.Class = NewInsightClassService(opts...)
	r.Severity = NewInsightSeverityService(opts...)
	r.Type = NewInsightTypeService(opts...)
	return
}

// Archive Security Center Insight
func (r *InsightService) Dismiss(ctx context.Context, issueID string, params InsightDismissParams, opts ...option.RequestOption) (res *InsightDismissResponse, err error) {
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if issueID == "" {
		err = errors.New("missing required issue_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/security-center/insights/%s/dismiss", accountOrZone, accountOrZoneID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Get Security Center Insights
func (r *InsightService) Get(ctx context.Context, params InsightGetParams, opts ...option.RequestOption) (res *InsightGetResponse, err error) {
	var env InsightGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/security-center/insights", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InsightDismissResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success InsightDismissResponseSuccess `json:"success,required"`
	JSON    insightDismissResponseJSON    `json:"-"`
}

// insightDismissResponseJSON contains the JSON metadata for the struct
// [InsightDismissResponse]
type insightDismissResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightDismissResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightDismissResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type InsightDismissResponseSuccess bool

const (
	InsightDismissResponseSuccessTrue InsightDismissResponseSuccess = true
)

func (r InsightDismissResponseSuccess) IsKnown() bool {
	switch r {
	case InsightDismissResponseSuccessTrue:
		return true
	}
	return false
}

type InsightGetResponse struct {
	// Total number of results
	Count  int64                     `json:"count"`
	Issues []InsightGetResponseIssue `json:"issues"`
	// Current page within paginated list of results
	Page int64 `json:"page"`
	// Number of results per page of results
	PerPage int64                  `json:"per_page"`
	JSON    insightGetResponseJSON `json:"-"`
}

// insightGetResponseJSON contains the JSON metadata for the struct
// [InsightGetResponse]
type insightGetResponseJSON struct {
	Count       apijson.Field
	Issues      apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightGetResponseJSON) RawJSON() string {
	return r.raw
}

type InsightGetResponseIssue struct {
	ID          string                           `json:"id"`
	Dismissed   bool                             `json:"dismissed"`
	IssueClass  string                           `json:"issue_class"`
	IssueType   intel.IssueType                  `json:"issue_type"`
	Payload     interface{}                      `json:"payload"`
	ResolveLink string                           `json:"resolve_link"`
	ResolveText string                           `json:"resolve_text"`
	Severity    InsightGetResponseIssuesSeverity `json:"severity"`
	Since       time.Time                        `json:"since" format:"date-time"`
	Subject     string                           `json:"subject"`
	Timestamp   time.Time                        `json:"timestamp" format:"date-time"`
	JSON        insightGetResponseIssueJSON      `json:"-"`
}

// insightGetResponseIssueJSON contains the JSON metadata for the struct
// [InsightGetResponseIssue]
type insightGetResponseIssueJSON struct {
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

func (r *InsightGetResponseIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightGetResponseIssueJSON) RawJSON() string {
	return r.raw
}

type InsightGetResponseIssuesSeverity string

const (
	InsightGetResponseIssuesSeverityLow      InsightGetResponseIssuesSeverity = "Low"
	InsightGetResponseIssuesSeverityModerate InsightGetResponseIssuesSeverity = "Moderate"
	InsightGetResponseIssuesSeverityCritical InsightGetResponseIssuesSeverity = "Critical"
)

func (r InsightGetResponseIssuesSeverity) IsKnown() bool {
	switch r {
	case InsightGetResponseIssuesSeverityLow, InsightGetResponseIssuesSeverityModerate, InsightGetResponseIssuesSeverityCritical:
		return true
	}
	return false
}

type InsightDismissParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID  param.Field[string] `path:"zone_id"`
	Dismiss param.Field[bool]   `json:"dismiss"`
}

func (r InsightDismissParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InsightGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID        param.Field[string]            `path:"zone_id"`
	Dismissed     param.Field[bool]              `query:"dismissed"`
	IssueClass    param.Field[[]string]          `query:"issue_class"`
	IssueClassNeq param.Field[[]string]          `query:"issue_class~neq"`
	IssueType     param.Field[[]intel.IssueType] `query:"issue_type"`
	IssueTypeNeq  param.Field[[]intel.IssueType] `query:"issue_type~neq"`
	// Current page within paginated list of results
	Page param.Field[int64] `query:"page"`
	// Number of results per page of results
	PerPage     param.Field[int64]                      `query:"per_page"`
	Product     param.Field[[]string]                   `query:"product"`
	ProductNeq  param.Field[[]string]                   `query:"product~neq"`
	Severity    param.Field[[]intel.SeverityQueryParam] `query:"severity"`
	SeverityNeq param.Field[[]intel.SeverityQueryParam] `query:"severity~neq"`
	Subject     param.Field[[]string]                   `query:"subject"`
	SubjectNeq  param.Field[[]string]                   `query:"subject~neq"`
}

// URLQuery serializes [InsightGetParams]'s query parameters as `url.Values`.
func (r InsightGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InsightGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success InsightGetResponseEnvelopeSuccess `json:"success,required"`
	Result  InsightGetResponse                `json:"result"`
	JSON    insightGetResponseEnvelopeJSON    `json:"-"`
}

// insightGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InsightGetResponseEnvelope]
type insightGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type InsightGetResponseEnvelopeSuccess bool

const (
	InsightGetResponseEnvelopeSuccessTrue InsightGetResponseEnvelopeSuccess = true
)

func (r InsightGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InsightGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
