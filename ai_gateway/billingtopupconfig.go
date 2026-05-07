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

// BillingTopupConfigService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingTopupConfigService] method instead.
type BillingTopupConfigService struct {
	Options []option.RequestOption
}

// NewBillingTopupConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBillingTopupConfigService(opts ...option.RequestOption) (r *BillingTopupConfigService) {
	r = &BillingTopupConfigService{}
	r.Options = opts
	return
}

// Configure auto top-up with a balance threshold and top-up amount.
func (r *BillingTopupConfigService) New(ctx context.Context, params BillingTopupConfigNewParams, opts ...option.RequestOption) (res *BillingTopupConfigNewResponse, err error) {
	var env BillingTopupConfigNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/billing/topup/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Remove the auto top-up configuration for the account.
func (r *BillingTopupConfigService) Delete(ctx context.Context, body BillingTopupConfigDeleteParams, opts ...option.RequestOption) (res *BillingTopupConfigDeleteResponse, err error) {
	var env BillingTopupConfigDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/billing/topup/config", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve the current auto top-up threshold, amount, and any error state.
func (r *BillingTopupConfigService) Get(ctx context.Context, query BillingTopupConfigGetParams, opts ...option.RequestOption) (res *BillingTopupConfigGetResponse, err error) {
	var env BillingTopupConfigGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/billing/topup/config", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type BillingTopupConfigNewResponse struct {
	Amount    float64                           `json:"amount" api:"required"`
	Threshold float64                           `json:"threshold" api:"required"`
	JSON      billingTopupConfigNewResponseJSON `json:"-"`
}

// billingTopupConfigNewResponseJSON contains the JSON metadata for the struct
// [BillingTopupConfigNewResponse]
type billingTopupConfigNewResponseJSON struct {
	Amount      apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigNewResponseJSON) RawJSON() string {
	return r.raw
}

type BillingTopupConfigDeleteResponse = interface{}

type BillingTopupConfigGetResponse struct {
	Amount         float64                           `json:"amount" api:"required,nullable"`
	DisabledReason string                            `json:"disabledReason" api:"required,nullable"`
	Error          string                            `json:"error" api:"required,nullable"`
	LastFailedAt   float64                           `json:"lastFailedAt" api:"required,nullable"`
	Threshold      float64                           `json:"threshold" api:"required,nullable"`
	JSON           billingTopupConfigGetResponseJSON `json:"-"`
}

// billingTopupConfigGetResponseJSON contains the JSON metadata for the struct
// [BillingTopupConfigGetResponse]
type billingTopupConfigGetResponseJSON struct {
	Amount         apijson.Field
	DisabledReason apijson.Field
	Error          apijson.Field
	LastFailedAt   apijson.Field
	Threshold      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BillingTopupConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigGetResponseJSON) RawJSON() string {
	return r.raw
}

type BillingTopupConfigNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Auto top-up amount in cents (min 1000).
	Amount param.Field[int64] `json:"amount" api:"required"`
	// Balance threshold in cents that triggers auto top-up (min 500).
	Threshold param.Field[int64] `json:"threshold" api:"required"`
}

func (r BillingTopupConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingTopupConfigNewResponseEnvelope struct {
	Errors     []BillingTopupConfigNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []BillingTopupConfigNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     BillingTopupConfigNewResponse                   `json:"result" api:"required"`
	Success    BillingTopupConfigNewResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo BillingTopupConfigNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       billingTopupConfigNewResponseEnvelopeJSON       `json:"-"`
}

// billingTopupConfigNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [BillingTopupConfigNewResponseEnvelope]
type billingTopupConfigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BillingTopupConfigNewResponseEnvelopeErrors struct {
	Code    float64                                         `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	JSON    billingTopupConfigNewResponseEnvelopeErrorsJSON `json:"-"`
}

// billingTopupConfigNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BillingTopupConfigNewResponseEnvelopeErrors]
type billingTopupConfigNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupConfigNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BillingTopupConfigNewResponseEnvelopeMessages struct {
	Code    float64                                           `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	JSON    billingTopupConfigNewResponseEnvelopeMessagesJSON `json:"-"`
}

// billingTopupConfigNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BillingTopupConfigNewResponseEnvelopeMessages]
type billingTopupConfigNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupConfigNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BillingTopupConfigNewResponseEnvelopeSuccess bool

const (
	BillingTopupConfigNewResponseEnvelopeSuccessTrue BillingTopupConfigNewResponseEnvelopeSuccess = true
)

func (r BillingTopupConfigNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingTopupConfigNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BillingTopupConfigNewResponseEnvelopeResultInfo struct {
	HasMore    bool                                                `json:"has_more" api:"required"`
	Page       float64                                             `json:"page" api:"required"`
	PerPage    float64                                             `json:"per_page" api:"required"`
	TotalCount float64                                             `json:"total_count" api:"required"`
	JSON       billingTopupConfigNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// billingTopupConfigNewResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [BillingTopupConfigNewResponseEnvelopeResultInfo]
type billingTopupConfigNewResponseEnvelopeResultInfoJSON struct {
	HasMore     apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupConfigNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type BillingTopupConfigDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BillingTopupConfigDeleteResponseEnvelope struct {
	Errors     []BillingTopupConfigDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []BillingTopupConfigDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     BillingTopupConfigDeleteResponse                   `json:"result" api:"required"`
	Success    BillingTopupConfigDeleteResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo BillingTopupConfigDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       billingTopupConfigDeleteResponseEnvelopeJSON       `json:"-"`
}

// billingTopupConfigDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [BillingTopupConfigDeleteResponseEnvelope]
type billingTopupConfigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BillingTopupConfigDeleteResponseEnvelopeErrors struct {
	Code    float64                                            `json:"code" api:"required"`
	Message string                                             `json:"message" api:"required"`
	JSON    billingTopupConfigDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// billingTopupConfigDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [BillingTopupConfigDeleteResponseEnvelopeErrors]
type billingTopupConfigDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupConfigDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BillingTopupConfigDeleteResponseEnvelopeMessages struct {
	Code    float64                                              `json:"code" api:"required"`
	Message string                                               `json:"message" api:"required"`
	JSON    billingTopupConfigDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// billingTopupConfigDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [BillingTopupConfigDeleteResponseEnvelopeMessages]
type billingTopupConfigDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupConfigDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BillingTopupConfigDeleteResponseEnvelopeSuccess bool

const (
	BillingTopupConfigDeleteResponseEnvelopeSuccessTrue BillingTopupConfigDeleteResponseEnvelopeSuccess = true
)

func (r BillingTopupConfigDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingTopupConfigDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BillingTopupConfigDeleteResponseEnvelopeResultInfo struct {
	HasMore    bool                                                   `json:"has_more" api:"required"`
	Page       float64                                                `json:"page" api:"required"`
	PerPage    float64                                                `json:"per_page" api:"required"`
	TotalCount float64                                                `json:"total_count" api:"required"`
	JSON       billingTopupConfigDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// billingTopupConfigDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [BillingTopupConfigDeleteResponseEnvelopeResultInfo]
type billingTopupConfigDeleteResponseEnvelopeResultInfoJSON struct {
	HasMore     apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupConfigDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type BillingTopupConfigGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BillingTopupConfigGetResponseEnvelope struct {
	Errors     []BillingTopupConfigGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []BillingTopupConfigGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     BillingTopupConfigGetResponse                   `json:"result" api:"required"`
	Success    BillingTopupConfigGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo BillingTopupConfigGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       billingTopupConfigGetResponseEnvelopeJSON       `json:"-"`
}

// billingTopupConfigGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [BillingTopupConfigGetResponseEnvelope]
type billingTopupConfigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BillingTopupConfigGetResponseEnvelopeErrors struct {
	Code    float64                                         `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	JSON    billingTopupConfigGetResponseEnvelopeErrorsJSON `json:"-"`
}

// billingTopupConfigGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BillingTopupConfigGetResponseEnvelopeErrors]
type billingTopupConfigGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BillingTopupConfigGetResponseEnvelopeMessages struct {
	Code    float64                                           `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	JSON    billingTopupConfigGetResponseEnvelopeMessagesJSON `json:"-"`
}

// billingTopupConfigGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BillingTopupConfigGetResponseEnvelopeMessages]
type billingTopupConfigGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BillingTopupConfigGetResponseEnvelopeSuccess bool

const (
	BillingTopupConfigGetResponseEnvelopeSuccessTrue BillingTopupConfigGetResponseEnvelopeSuccess = true
)

func (r BillingTopupConfigGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingTopupConfigGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BillingTopupConfigGetResponseEnvelopeResultInfo struct {
	HasMore    bool                                                `json:"has_more" api:"required"`
	Page       float64                                             `json:"page" api:"required"`
	PerPage    float64                                             `json:"per_page" api:"required"`
	TotalCount float64                                             `json:"total_count" api:"required"`
	JSON       billingTopupConfigGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// billingTopupConfigGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [BillingTopupConfigGetResponseEnvelopeResultInfo]
type billingTopupConfigGetResponseEnvelopeResultInfoJSON struct {
	HasMore     apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupConfigGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupConfigGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
