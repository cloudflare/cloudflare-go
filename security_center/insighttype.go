// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_center

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/intel"
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// InsightTypeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightTypeService] method instead.
type InsightTypeService struct {
	Options []option.RequestOption
}

// NewInsightTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInsightTypeService(opts ...option.RequestOption) (r *InsightTypeService) {
	r = &InsightTypeService{}
	r.Options = opts
	return
}

// Get Security Center Insight Counts by Type
func (r *InsightTypeService) Get(ctx context.Context, params InsightTypeGetParams, opts ...option.RequestOption) (res *[]InsightTypeGetResponse, err error) {
	var env InsightTypeGetResponseEnvelope
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
	path := fmt.Sprintf("%s/%s/security-center/insights/type", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InsightTypeGetResponse struct {
	Count int64                      `json:"count"`
	Value string                     `json:"value"`
	JSON  insightTypeGetResponseJSON `json:"-"`
}

// insightTypeGetResponseJSON contains the JSON metadata for the struct
// [InsightTypeGetResponse]
type insightTypeGetResponseJSON struct {
	Count       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightTypeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightTypeGetResponseJSON) RawJSON() string {
	return r.raw
}

type InsightTypeGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID        param.Field[string]                     `path:"zone_id"`
	Dismissed     param.Field[bool]                       `query:"dismissed"`
	IssueClass    param.Field[[]string]                   `query:"issue_class"`
	IssueClassNeq param.Field[[]string]                   `query:"issue_class~neq"`
	IssueType     param.Field[[]intel.IssueType]          `query:"issue_type"`
	IssueTypeNeq  param.Field[[]intel.IssueType]          `query:"issue_type~neq"`
	Product       param.Field[[]string]                   `query:"product"`
	ProductNeq    param.Field[[]string]                   `query:"product~neq"`
	Severity      param.Field[[]intel.SeverityQueryParam] `query:"severity"`
	SeverityNeq   param.Field[[]intel.SeverityQueryParam] `query:"severity~neq"`
	Subject       param.Field[[]string]                   `query:"subject"`
	SubjectNeq    param.Field[[]string]                   `query:"subject~neq"`
}

// URLQuery serializes [InsightTypeGetParams]'s query parameters as `url.Values`.
func (r InsightTypeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InsightTypeGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success InsightTypeGetResponseEnvelopeSuccess `json:"success,required"`
	Result  []InsightTypeGetResponse              `json:"result"`
	JSON    insightTypeGetResponseEnvelopeJSON    `json:"-"`
}

// insightTypeGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InsightTypeGetResponseEnvelope]
type insightTypeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightTypeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightTypeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type InsightTypeGetResponseEnvelopeSuccess bool

const (
	InsightTypeGetResponseEnvelopeSuccessTrue InsightTypeGetResponseEnvelopeSuccess = true
)

func (r InsightTypeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InsightTypeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
