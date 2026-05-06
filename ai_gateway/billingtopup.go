// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// BillingTopupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingTopupService] method instead.
type BillingTopupService struct {
	Options []option.RequestOption
	Config  *BillingTopupConfigService
}

// NewBillingTopupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBillingTopupService(opts ...option.RequestOption) (r *BillingTopupService) {
	r = &BillingTopupService{}
	r.Options = opts
	r.Config = NewBillingTopupConfigService(opts...)
	return
}

// Create a credit top-up via Stripe PaymentIntent for the given account.
func (r *BillingTopupService) New(ctx context.Context, params BillingTopupNewParams, opts ...option.RequestOption) (res *BillingTopupNewResponse, err error) {
	var env BillingTopupNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/billing/topup", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get the payment processing status of a top-up by its invoice ID.
func (r *BillingTopupService) Status(ctx context.Context, params BillingTopupStatusParams, opts ...option.RequestOption) (res *BillingTopupStatusResponse, err error) {
	var env BillingTopupStatusResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/billing/topup/status", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type BillingTopupNewResponse struct {
	// Stripe PaymentIntent client secret.
	ClientSecret string `json:"client_secret" api:"required,nullable"`
	// Whether the user was already onboarded.
	Onboarding bool `json:"onboarding" api:"required"`
	// Stripe invoice ID.
	PaymentIntentID string `json:"payment_intent_id" api:"required"`
	// Card brand (visa, mastercard, etc.).
	Brand string `json:"brand"`
	// Last 4 digits of card.
	Last4 string                      `json:"last4"`
	JSON  billingTopupNewResponseJSON `json:"-"`
}

// billingTopupNewResponseJSON contains the JSON metadata for the struct
// [BillingTopupNewResponse]
type billingTopupNewResponseJSON struct {
	ClientSecret    apijson.Field
	Onboarding      apijson.Field
	PaymentIntentID apijson.Field
	Brand           apijson.Field
	Last4           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BillingTopupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupNewResponseJSON) RawJSON() string {
	return r.raw
}

type BillingTopupStatusResponse struct {
	PaymentIntentID string                           `json:"payment_intent_id" api:"required"`
	Status          BillingTopupStatusResponseStatus `json:"status" api:"required"`
	JSON            billingTopupStatusResponseJSON   `json:"-"`
}

// billingTopupStatusResponseJSON contains the JSON metadata for the struct
// [BillingTopupStatusResponse]
type billingTopupStatusResponseJSON struct {
	PaymentIntentID apijson.Field
	Status          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BillingTopupStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupStatusResponseJSON) RawJSON() string {
	return r.raw
}

type BillingTopupStatusResponseStatus string

const (
	BillingTopupStatusResponseStatusCompleted BillingTopupStatusResponseStatus = "completed"
	BillingTopupStatusResponseStatusPending   BillingTopupStatusResponseStatus = "pending"
)

func (r BillingTopupStatusResponseStatus) IsKnown() bool {
	switch r {
	case BillingTopupStatusResponseStatusCompleted, BillingTopupStatusResponseStatusPending:
		return true
	}
	return false
}

type BillingTopupNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Top-up amount in cents (min 1000).
	Amount param.Field[int64] `json:"amount" api:"required"`
}

func (r BillingTopupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingTopupNewResponseEnvelope struct {
	Errors     []BillingTopupNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []BillingTopupNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     BillingTopupNewResponse                   `json:"result" api:"required"`
	Success    BillingTopupNewResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo BillingTopupNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       billingTopupNewResponseEnvelopeJSON       `json:"-"`
}

// billingTopupNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [BillingTopupNewResponseEnvelope]
type billingTopupNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BillingTopupNewResponseEnvelopeErrors struct {
	Code    float64                                   `json:"code" api:"required"`
	Message string                                    `json:"message" api:"required"`
	JSON    billingTopupNewResponseEnvelopeErrorsJSON `json:"-"`
}

// billingTopupNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BillingTopupNewResponseEnvelopeErrors]
type billingTopupNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BillingTopupNewResponseEnvelopeMessages struct {
	Code    float64                                     `json:"code" api:"required"`
	Message string                                      `json:"message" api:"required"`
	JSON    billingTopupNewResponseEnvelopeMessagesJSON `json:"-"`
}

// billingTopupNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [BillingTopupNewResponseEnvelopeMessages]
type billingTopupNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BillingTopupNewResponseEnvelopeSuccess bool

const (
	BillingTopupNewResponseEnvelopeSuccessTrue BillingTopupNewResponseEnvelopeSuccess = true
)

func (r BillingTopupNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingTopupNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BillingTopupNewResponseEnvelopeResultInfo struct {
	HasMore    bool                                          `json:"has_more" api:"required"`
	Page       float64                                       `json:"page" api:"required"`
	PerPage    float64                                       `json:"per_page" api:"required"`
	TotalCount float64                                       `json:"total_count" api:"required"`
	JSON       billingTopupNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// billingTopupNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [BillingTopupNewResponseEnvelopeResultInfo]
type billingTopupNewResponseEnvelopeResultInfoJSON struct {
	HasMore     apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type BillingTopupStatusParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Stripe invoice ID to check status for.
	PaymentIntentID param.Field[string] `json:"payment_intent_id" api:"required"`
}

func (r BillingTopupStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingTopupStatusResponseEnvelope struct {
	Errors     []BillingTopupStatusResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []BillingTopupStatusResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     BillingTopupStatusResponse                   `json:"result" api:"required"`
	Success    BillingTopupStatusResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo BillingTopupStatusResponseEnvelopeResultInfo `json:"result_info"`
	JSON       billingTopupStatusResponseEnvelopeJSON       `json:"-"`
}

// billingTopupStatusResponseEnvelopeJSON contains the JSON metadata for the struct
// [BillingTopupStatusResponseEnvelope]
type billingTopupStatusResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupStatusResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupStatusResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BillingTopupStatusResponseEnvelopeErrors struct {
	Code    float64                                      `json:"code" api:"required"`
	Message string                                       `json:"message" api:"required"`
	JSON    billingTopupStatusResponseEnvelopeErrorsJSON `json:"-"`
}

// billingTopupStatusResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BillingTopupStatusResponseEnvelopeErrors]
type billingTopupStatusResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupStatusResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupStatusResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BillingTopupStatusResponseEnvelopeMessages struct {
	Code    float64                                        `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	JSON    billingTopupStatusResponseEnvelopeMessagesJSON `json:"-"`
}

// billingTopupStatusResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BillingTopupStatusResponseEnvelopeMessages]
type billingTopupStatusResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupStatusResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupStatusResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BillingTopupStatusResponseEnvelopeSuccess bool

const (
	BillingTopupStatusResponseEnvelopeSuccessTrue BillingTopupStatusResponseEnvelopeSuccess = true
)

func (r BillingTopupStatusResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingTopupStatusResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BillingTopupStatusResponseEnvelopeResultInfo struct {
	HasMore    bool                                             `json:"has_more" api:"required"`
	Page       float64                                          `json:"page" api:"required"`
	PerPage    float64                                          `json:"per_page" api:"required"`
	TotalCount float64                                          `json:"total_count" api:"required"`
	JSON       billingTopupStatusResponseEnvelopeResultInfoJSON `json:"-"`
}

// billingTopupStatusResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [BillingTopupStatusResponseEnvelopeResultInfo]
type billingTopupStatusResponseEnvelopeResultInfoJSON struct {
	HasMore     apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingTopupStatusResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingTopupStatusResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
