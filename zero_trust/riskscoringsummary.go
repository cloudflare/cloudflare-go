// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// RiskScoringSummaryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRiskScoringSummaryService] method
// instead.
type RiskScoringSummaryService struct {
	Options []option.RequestOption
}

// NewRiskScoringSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRiskScoringSummaryService(opts ...option.RequestOption) (r *RiskScoringSummaryService) {
	r = &RiskScoringSummaryService{}
	r.Options = opts
	return
}

// Get risk score info for all users in the account
func (r *RiskScoringSummaryService) Get(ctx context.Context, accountIdentifier string, query RiskScoringSummaryGetParams, opts ...option.RequestOption) (res *RiskScoringSummaryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RiskScoringSummaryGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/summary", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RiskScoringSummaryGetResponse struct {
	Users []RiskScoringSummaryGetResponseUser `json:"users"`
	JSON  riskScoringSummaryGetResponseJSON   `json:"-"`
}

// riskScoringSummaryGetResponseJSON contains the JSON metadata for the struct
// [RiskScoringSummaryGetResponse]
type riskScoringSummaryGetResponseJSON struct {
	Users       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringSummaryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringSummaryGetResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringSummaryGetResponseUser struct {
	Email        string                                         `json:"email,required"`
	EventCount   int64                                          `json:"event_count,required"`
	LastEvent    time.Time                                      `json:"last_event,required" format:"date-time"`
	MaxRiskLevel RiskScoringSummaryGetResponseUsersMaxRiskLevel `json:"max_risk_level,required,nullable"`
	Name         string                                         `json:"name,required"`
	// The ID for a user
	UserID string                                `json:"user_id,required"`
	JSON   riskScoringSummaryGetResponseUserJSON `json:"-"`
}

// riskScoringSummaryGetResponseUserJSON contains the JSON metadata for the struct
// [RiskScoringSummaryGetResponseUser]
type riskScoringSummaryGetResponseUserJSON struct {
	Email        apijson.Field
	EventCount   apijson.Field
	LastEvent    apijson.Field
	MaxRiskLevel apijson.Field
	Name         apijson.Field
	UserID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RiskScoringSummaryGetResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringSummaryGetResponseUserJSON) RawJSON() string {
	return r.raw
}

type RiskScoringSummaryGetResponseUsersMaxRiskLevel string

const (
	RiskScoringSummaryGetResponseUsersMaxRiskLevelLow    RiskScoringSummaryGetResponseUsersMaxRiskLevel = "low"
	RiskScoringSummaryGetResponseUsersMaxRiskLevelMedium RiskScoringSummaryGetResponseUsersMaxRiskLevel = "medium"
	RiskScoringSummaryGetResponseUsersMaxRiskLevelHigh   RiskScoringSummaryGetResponseUsersMaxRiskLevel = "high"
)

func (r RiskScoringSummaryGetResponseUsersMaxRiskLevel) IsKnown() bool {
	switch r {
	case RiskScoringSummaryGetResponseUsersMaxRiskLevelLow, RiskScoringSummaryGetResponseUsersMaxRiskLevelMedium, RiskScoringSummaryGetResponseUsersMaxRiskLevelHigh:
		return true
	}
	return false
}

type RiskScoringSummaryGetParams struct {
	Direction param.Field[RiskScoringSummaryGetParamsDirection] `query:"direction"`
	OrderBy   param.Field[RiskScoringSummaryGetParamsOrderBy]   `query:"order_by"`
	Page      param.Field[int64]                                `query:"page"`
	PerPage   param.Field[int64]                                `query:"per_page"`
}

// URLQuery serializes [RiskScoringSummaryGetParams]'s query parameters as
// `url.Values`.
func (r RiskScoringSummaryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RiskScoringSummaryGetParamsDirection string

const (
	RiskScoringSummaryGetParamsDirectionDesc RiskScoringSummaryGetParamsDirection = "desc"
	RiskScoringSummaryGetParamsDirectionAsc  RiskScoringSummaryGetParamsDirection = "asc"
)

func (r RiskScoringSummaryGetParamsDirection) IsKnown() bool {
	switch r {
	case RiskScoringSummaryGetParamsDirectionDesc, RiskScoringSummaryGetParamsDirectionAsc:
		return true
	}
	return false
}

type RiskScoringSummaryGetParamsOrderBy string

const (
	RiskScoringSummaryGetParamsOrderByTimestamp    RiskScoringSummaryGetParamsOrderBy = "timestamp"
	RiskScoringSummaryGetParamsOrderByEventCount   RiskScoringSummaryGetParamsOrderBy = "event_count"
	RiskScoringSummaryGetParamsOrderByMaxRiskLevel RiskScoringSummaryGetParamsOrderBy = "max_risk_level"
)

func (r RiskScoringSummaryGetParamsOrderBy) IsKnown() bool {
	switch r {
	case RiskScoringSummaryGetParamsOrderByTimestamp, RiskScoringSummaryGetParamsOrderByEventCount, RiskScoringSummaryGetParamsOrderByMaxRiskLevel:
		return true
	}
	return false
}

type RiskScoringSummaryGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   RiskScoringSummaryGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success    RiskScoringSummaryGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RiskScoringSummaryGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       riskScoringSummaryGetResponseEnvelopeJSON       `json:"-"`
}

// riskScoringSummaryGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RiskScoringSummaryGetResponseEnvelope]
type riskScoringSummaryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringSummaryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringSummaryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RiskScoringSummaryGetResponseEnvelopeSuccess bool

const (
	RiskScoringSummaryGetResponseEnvelopeSuccessTrue RiskScoringSummaryGetResponseEnvelopeSuccess = true
)

func (r RiskScoringSummaryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RiskScoringSummaryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RiskScoringSummaryGetResponseEnvelopeResultInfo struct {
	Count      int64                                               `json:"count,required"`
	Page       int64                                               `json:"page,required"`
	PerPage    int64                                               `json:"per_page,required"`
	TotalCount int64                                               `json:"total_count,required"`
	JSON       riskScoringSummaryGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// riskScoringSummaryGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [RiskScoringSummaryGetResponseEnvelopeResultInfo]
type riskScoringSummaryGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringSummaryGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringSummaryGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
