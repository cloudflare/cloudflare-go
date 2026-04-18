// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
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

// Retrieves the detailed risk score breakdown for a specific user, including
// contributing factors.
func (r *RiskScoringService) Get(ctx context.Context, userID string, query RiskScoringGetParams, opts ...option.RequestOption) (res *RiskScoringGetResponse, err error) {
	var env RiskScoringGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/%s", query.AccountID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Resets risk scores for specified users, clearing their accumulated risk history.
func (r *RiskScoringService) Reset(ctx context.Context, userID string, body RiskScoringResetParams, opts ...option.RequestOption) (res *RiskScoringResetResponse, err error) {
	var env RiskScoringResetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/%s/reset", body.AccountID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type RiskScoringGetResponse struct {
	Email         string                          `json:"email" api:"required"`
	Events        []RiskScoringGetResponseEvent   `json:"events" api:"required"`
	Name          string                          `json:"name" api:"required"`
	LastResetTime time.Time                       `json:"last_reset_time" api:"nullable" format:"date-time"`
	RiskLevel     RiskScoringGetResponseRiskLevel `json:"risk_level"`
	JSON          riskScoringGetResponseJSON      `json:"-"`
}

// riskScoringGetResponseJSON contains the JSON metadata for the struct
// [RiskScoringGetResponse]
type riskScoringGetResponseJSON struct {
	Email         apijson.Field
	Events        apijson.Field
	Name          apijson.Field
	LastResetTime apijson.Field
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
	ID           string                                `json:"id" api:"required"`
	Name         string                                `json:"name" api:"required"`
	RiskLevel    RiskScoringGetResponseEventsRiskLevel `json:"risk_level" api:"required"`
	Timestamp    time.Time                             `json:"timestamp" api:"required" format:"date-time"`
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

type RiskScoringResetResponse = interface{}

type RiskScoringGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RiskScoringGetResponseEnvelope struct {
	Errors   []RiskScoringGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RiskScoringGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success    RiskScoringGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	Result     RiskScoringGetResponse                   `json:"result"`
	ResultInfo RiskScoringGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       riskScoringGetResponseEnvelopeJSON       `json:"-"`
}

// riskScoringGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RiskScoringGetResponseEnvelope]
type riskScoringGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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

type RiskScoringGetResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           RiskScoringGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             riskScoringGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// riskScoringGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RiskScoringGetResponseEnvelopeErrors]
type riskScoringGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RiskScoringGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RiskScoringGetResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    riskScoringGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// riskScoringGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RiskScoringGetResponseEnvelopeErrorsSource]
type riskScoringGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RiskScoringGetResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           RiskScoringGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             riskScoringGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// riskScoringGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RiskScoringGetResponseEnvelopeMessages]
type riskScoringGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RiskScoringGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RiskScoringGetResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    riskScoringGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// riskScoringGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RiskScoringGetResponseEnvelopeMessagesSource]
type riskScoringGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64 `json:"total_count"`
	// The number of total pages in the entire result set.
	TotalPages float64                                      `json:"total_pages"`
	JSON       riskScoringGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// riskScoringGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [RiskScoringGetResponseEnvelopeResultInfo]
type riskScoringGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RiskScoringResetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RiskScoringResetResponseEnvelope struct {
	Errors   []RiskScoringResetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RiskScoringResetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success RiskScoringResetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  RiskScoringResetResponse                `json:"result" api:"nullable"`
	JSON    riskScoringResetResponseEnvelopeJSON    `json:"-"`
}

// riskScoringResetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RiskScoringResetResponseEnvelope]
type riskScoringResetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringResetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringResetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RiskScoringResetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           RiskScoringResetResponseEnvelopeErrorsSource `json:"source"`
	JSON             riskScoringResetResponseEnvelopeErrorsJSON   `json:"-"`
}

// riskScoringResetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RiskScoringResetResponseEnvelopeErrors]
type riskScoringResetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RiskScoringResetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringResetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RiskScoringResetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    riskScoringResetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// riskScoringResetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RiskScoringResetResponseEnvelopeErrorsSource]
type riskScoringResetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringResetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringResetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RiskScoringResetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           RiskScoringResetResponseEnvelopeMessagesSource `json:"source"`
	JSON             riskScoringResetResponseEnvelopeMessagesJSON   `json:"-"`
}

// riskScoringResetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RiskScoringResetResponseEnvelopeMessages]
type riskScoringResetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RiskScoringResetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringResetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RiskScoringResetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    riskScoringResetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// riskScoringResetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RiskScoringResetResponseEnvelopeMessagesSource]
type riskScoringResetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringResetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringResetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
