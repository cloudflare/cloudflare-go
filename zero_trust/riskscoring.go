// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// RiskScoringService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRiskScoringService] method instead.
type RiskScoringService struct {
	Options      []option.RequestOption
	Behaviours   *RiskScoringBehaviourService
	Summary      *RiskScoringSummaryService
	Integrations *RiskScoringIntegrationService
}

// NewRiskScoringService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRiskScoringService(opts ...option.RequestOption) (r *RiskScoringService) {
	r = &RiskScoringService{}
	r.Options = opts
	r.Behaviours = NewRiskScoringBehaviourService(opts...)
	r.Summary = NewRiskScoringSummaryService(opts...)
	r.Integrations = NewRiskScoringIntegrationService(opts...)
	return
}

// Get risk event/score information for a specific user
func (r *RiskScoringService) Get(ctx context.Context, accountIdentifier string, userID string, query RiskScoringGetParams, opts ...option.RequestOption) (res *RiskScoringGetResponse, err error) {
	var env RiskScoringGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/%s", accountIdentifier, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Clear the risk score for a particular user
func (r *RiskScoringService) Reset(ctx context.Context, accountIdentifier string, userID string, opts ...option.RequestOption) (res *RiskScoringResetResponseUnion, err error) {
	var env RiskScoringResetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/%s/reset", accountIdentifier, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RiskScoringGetResponse struct {
	Email         string                          `json:"email"`
	Events        []RiskScoringGetResponseEvent   `json:"events"`
	LastResetTime time.Time                       `json:"last_reset_time,nullable" format:"date-time"`
	Name          string                          `json:"name"`
	RiskLevel     RiskScoringGetResponseRiskLevel `json:"risk_level,nullable"`
	JSON          riskScoringGetResponseJSON      `json:"-"`
}

// riskScoringGetResponseJSON contains the JSON metadata for the struct
// [RiskScoringGetResponse]
type riskScoringGetResponseJSON struct {
	Email         apijson.Field
	Events        apijson.Field
	LastResetTime apijson.Field
	Name          apijson.Field
	RiskLevel     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RiskScoringGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringGetResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringGetResponseEvent struct {
	ID           string                                `json:"id,required"`
	Name         string                                `json:"name,required"`
	RiskLevel    RiskScoringGetResponseEventsRiskLevel `json:"risk_level,required"`
	Timestamp    time.Time                             `json:"timestamp,required" format:"date-time"`
	EventDetails interface{}                           `json:"event_details"`
	JSON         riskScoringGetResponseEventJSON       `json:"-"`
}

// riskScoringGetResponseEventJSON contains the JSON metadata for the struct
// [RiskScoringGetResponseEvent]
type riskScoringGetResponseEventJSON struct {
	ID           apijson.Field
	Name         apijson.Field
	RiskLevel    apijson.Field
	Timestamp    apijson.Field
	EventDetails apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RiskScoringGetResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringGetResponseEventJSON) RawJSON() string {
	return r.raw
}

type RiskScoringGetResponseEventsRiskLevel string

const (
	RiskScoringGetResponseEventsRiskLevelLow    RiskScoringGetResponseEventsRiskLevel = "low"
	RiskScoringGetResponseEventsRiskLevelMedium RiskScoringGetResponseEventsRiskLevel = "medium"
	RiskScoringGetResponseEventsRiskLevelHigh   RiskScoringGetResponseEventsRiskLevel = "high"
)

func (r RiskScoringGetResponseEventsRiskLevel) IsKnown() bool {
	switch r {
	case RiskScoringGetResponseEventsRiskLevelLow, RiskScoringGetResponseEventsRiskLevelMedium, RiskScoringGetResponseEventsRiskLevelHigh:
		return true
	}
	return false
}

type RiskScoringGetResponseRiskLevel string

const (
	RiskScoringGetResponseRiskLevelLow    RiskScoringGetResponseRiskLevel = "low"
	RiskScoringGetResponseRiskLevelMedium RiskScoringGetResponseRiskLevel = "medium"
	RiskScoringGetResponseRiskLevelHigh   RiskScoringGetResponseRiskLevel = "high"
)

func (r RiskScoringGetResponseRiskLevel) IsKnown() bool {
	switch r {
	case RiskScoringGetResponseRiskLevelLow, RiskScoringGetResponseRiskLevelMedium, RiskScoringGetResponseRiskLevelHigh:
		return true
	}
	return false
}

// Union satisfied by [zero_trust.RiskScoringResetResponseUnknown] or
// [shared.UnionString].
type RiskScoringResetResponseUnion interface {
	ImplementsZeroTrustRiskScoringResetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RiskScoringResetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RiskScoringGetParams struct {
	Direction param.Field[RiskScoringGetParamsDirection] `query:"direction"`
	OrderBy   param.Field[RiskScoringGetParamsOrderBy]   `query:"order_by"`
	Page      param.Field[int64]                         `query:"page"`
	PerPage   param.Field[int64]                         `query:"per_page"`
}

// URLQuery serializes [RiskScoringGetParams]'s query parameters as `url.Values`.
func (r RiskScoringGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RiskScoringGetParamsDirection string

const (
	RiskScoringGetParamsDirectionDesc RiskScoringGetParamsDirection = "desc"
	RiskScoringGetParamsDirectionAsc  RiskScoringGetParamsDirection = "asc"
)

func (r RiskScoringGetParamsDirection) IsKnown() bool {
	switch r {
	case RiskScoringGetParamsDirectionDesc, RiskScoringGetParamsDirectionAsc:
		return true
	}
	return false
}

type RiskScoringGetParamsOrderBy string

const (
	RiskScoringGetParamsOrderByTimestamp RiskScoringGetParamsOrderBy = "timestamp"
	RiskScoringGetParamsOrderByRiskLevel RiskScoringGetParamsOrderBy = "risk_level"
)

func (r RiskScoringGetParamsOrderBy) IsKnown() bool {
	switch r {
	case RiskScoringGetParamsOrderByTimestamp, RiskScoringGetParamsOrderByRiskLevel:
		return true
	}
	return false
}

type RiskScoringGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   RiskScoringGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success    RiskScoringGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RiskScoringGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       riskScoringGetResponseEnvelopeJSON       `json:"-"`
}

// riskScoringGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RiskScoringGetResponseEnvelope]
type riskScoringGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RiskScoringGetResponseEnvelopeSuccess bool

const (
	RiskScoringGetResponseEnvelopeSuccessTrue RiskScoringGetResponseEnvelopeSuccess = true
)

func (r RiskScoringGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RiskScoringGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RiskScoringGetResponseEnvelopeResultInfo struct {
	Count      int64                                        `json:"count,required"`
	Page       int64                                        `json:"page,required"`
	PerPage    int64                                        `json:"per_page,required"`
	TotalCount int64                                        `json:"total_count,required"`
	JSON       riskScoringGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// riskScoringGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [RiskScoringGetResponseEnvelopeResultInfo]
type riskScoringGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RiskScoringResetResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   RiskScoringResetResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success RiskScoringResetResponseEnvelopeSuccess `json:"success,required"`
	JSON    riskScoringResetResponseEnvelopeJSON    `json:"-"`
}

// riskScoringResetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RiskScoringResetResponseEnvelope]
type riskScoringResetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringResetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringResetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RiskScoringResetResponseEnvelopeSuccess bool

const (
	RiskScoringResetResponseEnvelopeSuccessTrue RiskScoringResetResponseEnvelopeSuccess = true
)

func (r RiskScoringResetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RiskScoringResetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
