// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// BillingSpendingLimitService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingSpendingLimitService] method instead.
type BillingSpendingLimitService struct {
	Options []option.RequestOption
}

// NewBillingSpendingLimitService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBillingSpendingLimitService(opts ...option.RequestOption) (r *BillingSpendingLimitService) {
	r = &BillingSpendingLimitService{}
	r.Options = opts
	return
}

// Configure a spending limit with amount, strategy, and duration.
func (r *BillingSpendingLimitService) New(ctx context.Context, params BillingSpendingLimitNewParams, opts ...option.RequestOption) (res *BillingSpendingLimitNewResponse, err error) {
	var env BillingSpendingLimitNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/billing/spending-limit", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Remove the spending limit for the account.
func (r *BillingSpendingLimitService) Delete(ctx context.Context, body BillingSpendingLimitDeleteParams, opts ...option.RequestOption) (res *BillingSpendingLimitDeleteResponse, err error) {
	var env BillingSpendingLimitDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/billing/spending-limit", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve the current spending limit configuration for the account.
func (r *BillingSpendingLimitService) Get(ctx context.Context, query BillingSpendingLimitGetParams, opts ...option.RequestOption) (res *BillingSpendingLimitGetResponse, err error) {
	var env BillingSpendingLimitGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/billing/spending-limit", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type BillingSpendingLimitNewResponse = interface{}

type BillingSpendingLimitDeleteResponse = interface{}

type BillingSpendingLimitGetResponse struct {
	Config  BillingSpendingLimitGetResponseConfig `json:"config" api:"required"`
	Enabled bool                                  `json:"enabled" api:"required"`
	JSON    billingSpendingLimitGetResponseJSON   `json:"-"`
}

// billingSpendingLimitGetResponseJSON contains the JSON metadata for the struct
// [BillingSpendingLimitGetResponse]
type billingSpendingLimitGetResponseJSON struct {
	Config      apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitGetResponseJSON) RawJSON() string {
	return r.raw
}

type BillingSpendingLimitGetResponseConfig struct {
	Amount   float64                                   `json:"amount" api:"required,nullable"`
	Duration string                                    `json:"duration" api:"required,nullable"`
	Strategy string                                    `json:"strategy" api:"required,nullable"`
	JSON     billingSpendingLimitGetResponseConfigJSON `json:"-"`
}

// billingSpendingLimitGetResponseConfigJSON contains the JSON metadata for the
// struct [BillingSpendingLimitGetResponseConfig]
type billingSpendingLimitGetResponseConfigJSON struct {
	Amount      apijson.Field
	Duration    apijson.Field
	Strategy    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitGetResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitGetResponseConfigJSON) RawJSON() string {
	return r.raw
}

type BillingSpendingLimitNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Spending limit amount in cents (min 100).
	Amount param.Field[int64] `json:"amount" api:"required"`
	// Spending limit duration.
	Duration param.Field[BillingSpendingLimitNewParamsDuration] `json:"duration" api:"required"`
	// Spending limit strategy.
	Strategy param.Field[BillingSpendingLimitNewParamsStrategy] `json:"strategy" api:"required"`
}

func (r BillingSpendingLimitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Spending limit duration.
type BillingSpendingLimitNewParamsDuration string

const (
	BillingSpendingLimitNewParamsDurationDaily   BillingSpendingLimitNewParamsDuration = "daily"
	BillingSpendingLimitNewParamsDurationWeekly  BillingSpendingLimitNewParamsDuration = "weekly"
	BillingSpendingLimitNewParamsDurationMonthly BillingSpendingLimitNewParamsDuration = "monthly"
)

func (r BillingSpendingLimitNewParamsDuration) IsKnown() bool {
	switch r {
	case BillingSpendingLimitNewParamsDurationDaily, BillingSpendingLimitNewParamsDurationWeekly, BillingSpendingLimitNewParamsDurationMonthly:
		return true
	}
	return false
}

// Spending limit strategy.
type BillingSpendingLimitNewParamsStrategy string

const (
	BillingSpendingLimitNewParamsStrategyFixed   BillingSpendingLimitNewParamsStrategy = "fixed"
	BillingSpendingLimitNewParamsStrategySliding BillingSpendingLimitNewParamsStrategy = "sliding"
)

func (r BillingSpendingLimitNewParamsStrategy) IsKnown() bool {
	switch r {
	case BillingSpendingLimitNewParamsStrategyFixed, BillingSpendingLimitNewParamsStrategySliding:
		return true
	}
	return false
}

type BillingSpendingLimitNewResponseEnvelope struct {
	Errors     []BillingSpendingLimitNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []BillingSpendingLimitNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     BillingSpendingLimitNewResponse                   `json:"result" api:"required"`
	Success    BillingSpendingLimitNewResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo BillingSpendingLimitNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       billingSpendingLimitNewResponseEnvelopeJSON       `json:"-"`
}

// billingSpendingLimitNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [BillingSpendingLimitNewResponseEnvelope]
type billingSpendingLimitNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BillingSpendingLimitNewResponseEnvelopeErrors struct {
	Code    float64                                           `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	JSON    billingSpendingLimitNewResponseEnvelopeErrorsJSON `json:"-"`
}

// billingSpendingLimitNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BillingSpendingLimitNewResponseEnvelopeErrors]
type billingSpendingLimitNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BillingSpendingLimitNewResponseEnvelopeMessages struct {
	Code    float64                                             `json:"code" api:"required"`
	Message string                                              `json:"message" api:"required"`
	JSON    billingSpendingLimitNewResponseEnvelopeMessagesJSON `json:"-"`
}

// billingSpendingLimitNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [BillingSpendingLimitNewResponseEnvelopeMessages]
type billingSpendingLimitNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BillingSpendingLimitNewResponseEnvelopeSuccess bool

const (
	BillingSpendingLimitNewResponseEnvelopeSuccessTrue BillingSpendingLimitNewResponseEnvelopeSuccess = true
)

func (r BillingSpendingLimitNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingSpendingLimitNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BillingSpendingLimitNewResponseEnvelopeResultInfo struct {
	HasMore    bool                                                  `json:"has_more" api:"required"`
	Page       float64                                               `json:"page" api:"required"`
	PerPage    float64                                               `json:"per_page" api:"required"`
	TotalCount float64                                               `json:"total_count" api:"required"`
	JSON       billingSpendingLimitNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// billingSpendingLimitNewResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [BillingSpendingLimitNewResponseEnvelopeResultInfo]
type billingSpendingLimitNewResponseEnvelopeResultInfoJSON struct {
	HasMore     apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type BillingSpendingLimitDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BillingSpendingLimitDeleteResponseEnvelope struct {
	Errors     []BillingSpendingLimitDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []BillingSpendingLimitDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     BillingSpendingLimitDeleteResponse                   `json:"result" api:"required"`
	Success    BillingSpendingLimitDeleteResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo BillingSpendingLimitDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       billingSpendingLimitDeleteResponseEnvelopeJSON       `json:"-"`
}

// billingSpendingLimitDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [BillingSpendingLimitDeleteResponseEnvelope]
type billingSpendingLimitDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BillingSpendingLimitDeleteResponseEnvelopeErrors struct {
	Code    float64                                              `json:"code" api:"required"`
	Message string                                               `json:"message" api:"required"`
	JSON    billingSpendingLimitDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// billingSpendingLimitDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [BillingSpendingLimitDeleteResponseEnvelopeErrors]
type billingSpendingLimitDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BillingSpendingLimitDeleteResponseEnvelopeMessages struct {
	Code    float64                                                `json:"code" api:"required"`
	Message string                                                 `json:"message" api:"required"`
	JSON    billingSpendingLimitDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// billingSpendingLimitDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [BillingSpendingLimitDeleteResponseEnvelopeMessages]
type billingSpendingLimitDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BillingSpendingLimitDeleteResponseEnvelopeSuccess bool

const (
	BillingSpendingLimitDeleteResponseEnvelopeSuccessTrue BillingSpendingLimitDeleteResponseEnvelopeSuccess = true
)

func (r BillingSpendingLimitDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingSpendingLimitDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BillingSpendingLimitDeleteResponseEnvelopeResultInfo struct {
	HasMore    bool                                                     `json:"has_more" api:"required"`
	Page       float64                                                  `json:"page" api:"required"`
	PerPage    float64                                                  `json:"per_page" api:"required"`
	TotalCount float64                                                  `json:"total_count" api:"required"`
	JSON       billingSpendingLimitDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// billingSpendingLimitDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [BillingSpendingLimitDeleteResponseEnvelopeResultInfo]
type billingSpendingLimitDeleteResponseEnvelopeResultInfoJSON struct {
	HasMore     apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type BillingSpendingLimitGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BillingSpendingLimitGetResponseEnvelope struct {
	Errors     []BillingSpendingLimitGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []BillingSpendingLimitGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     BillingSpendingLimitGetResponse                   `json:"result" api:"required"`
	Success    BillingSpendingLimitGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo BillingSpendingLimitGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       billingSpendingLimitGetResponseEnvelopeJSON       `json:"-"`
}

// billingSpendingLimitGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [BillingSpendingLimitGetResponseEnvelope]
type billingSpendingLimitGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BillingSpendingLimitGetResponseEnvelopeErrors struct {
	Code    float64                                           `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	JSON    billingSpendingLimitGetResponseEnvelopeErrorsJSON `json:"-"`
}

// billingSpendingLimitGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BillingSpendingLimitGetResponseEnvelopeErrors]
type billingSpendingLimitGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BillingSpendingLimitGetResponseEnvelopeMessages struct {
	Code    float64                                             `json:"code" api:"required"`
	Message string                                              `json:"message" api:"required"`
	JSON    billingSpendingLimitGetResponseEnvelopeMessagesJSON `json:"-"`
}

// billingSpendingLimitGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [BillingSpendingLimitGetResponseEnvelopeMessages]
type billingSpendingLimitGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BillingSpendingLimitGetResponseEnvelopeSuccess bool

const (
	BillingSpendingLimitGetResponseEnvelopeSuccessTrue BillingSpendingLimitGetResponseEnvelopeSuccess = true
)

func (r BillingSpendingLimitGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingSpendingLimitGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BillingSpendingLimitGetResponseEnvelopeResultInfo struct {
	HasMore    bool                                                  `json:"has_more" api:"required"`
	Page       float64                                               `json:"page" api:"required"`
	PerPage    float64                                               `json:"per_page" api:"required"`
	TotalCount float64                                               `json:"total_count" api:"required"`
	JSON       billingSpendingLimitGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// billingSpendingLimitGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [BillingSpendingLimitGetResponseEnvelopeResultInfo]
type billingSpendingLimitGetResponseEnvelopeResultInfoJSON struct {
	HasMore     apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingSpendingLimitGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingSpendingLimitGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
