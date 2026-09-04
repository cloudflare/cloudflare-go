// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// SubscriptionCancelReasonService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionCancelReasonService] method instead.
type SubscriptionCancelReasonService struct {
	Options []option.RequestOption
}

// NewSubscriptionCancelReasonService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSubscriptionCancelReasonService(opts ...option.RequestOption) (r *SubscriptionCancelReasonService) {
	r = &SubscriptionCancelReasonService{}
	r.Options = opts
	return
}

// Records a cancellation reason for an account subscription.
func (r *SubscriptionCancelReasonService) New(ctx context.Context, subscriptionIdentifier string, params SubscriptionCancelReasonNewParams, opts ...option.RequestOption) (res *SubscriptionCancelReasonNewResponse, err error) {
	var env SubscriptionCancelReasonNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if subscriptionIdentifier == "" {
		err = errors.New("missing required subscription_identifier parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/subscriptions/%s/cancel-reason", params.AccountID, subscriptionIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets the cancellation reason for an account subscription.
func (r *SubscriptionCancelReasonService) Get(ctx context.Context, subscriptionIdentifier string, query SubscriptionCancelReasonGetParams, opts ...option.RequestOption) (res *SubscriptionCancelReasonGetResponse, err error) {
	var env SubscriptionCancelReasonGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if subscriptionIdentifier == "" {
		err = errors.New("missing required subscription_identifier parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/subscriptions/%s/cancel-reason", query.AccountID, subscriptionIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SubscriptionCancelReasonNewResponse struct {
	// The cancel reason identifier.
	ID string `json:"id"`
	// Additional cancellation details.
	Other string `json:"other"`
	// The cancellation reason codes.
	ReasonCode []string `json:"reason_code"`
	// When the cancel reason was submitted.
	Submitted time.Time `json:"submitted" format:"date-time"`
	// The subscription identifier.
	SubscriptionID string                                  `json:"subscription_id"`
	JSON           subscriptionCancelReasonNewResponseJSON `json:"-"`
}

// subscriptionCancelReasonNewResponseJSON contains the JSON metadata for the
// struct [SubscriptionCancelReasonNewResponse]
type subscriptionCancelReasonNewResponseJSON struct {
	ID             apijson.Field
	Other          apijson.Field
	ReasonCode     apijson.Field
	Submitted      apijson.Field
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionCancelReasonNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCancelReasonNewResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionCancelReasonGetResponse struct {
	// The cancel reason identifier.
	ID string `json:"id"`
	// Additional cancellation details.
	Other string `json:"other"`
	// The cancellation reason codes.
	ReasonCode []string `json:"reason_code"`
	// When the cancel reason was submitted.
	Submitted time.Time `json:"submitted" format:"date-time"`
	// The subscription identifier.
	SubscriptionID string                                  `json:"subscription_id"`
	JSON           subscriptionCancelReasonGetResponseJSON `json:"-"`
}

// subscriptionCancelReasonGetResponseJSON contains the JSON metadata for the
// struct [SubscriptionCancelReasonGetResponse]
type subscriptionCancelReasonGetResponseJSON struct {
	ID             apijson.Field
	Other          apijson.Field
	ReasonCode     apijson.Field
	Submitted      apijson.Field
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionCancelReasonGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCancelReasonGetResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionCancelReasonNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Additional cancellation details.
	Other param.Field[string] `json:"other"`
	// The cancellation reason codes.
	ReasonCode param.Field[[]string] `json:"reason_code"`
}

func (r SubscriptionCancelReasonNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionCancelReasonNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo               `json:"errors" api:"required"`
	Messages []shared.ResponseInfo               `json:"messages" api:"required"`
	Result   SubscriptionCancelReasonNewResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success SubscriptionCancelReasonNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    subscriptionCancelReasonNewResponseEnvelopeJSON    `json:"-"`
}

// subscriptionCancelReasonNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [SubscriptionCancelReasonNewResponseEnvelope]
type subscriptionCancelReasonNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionCancelReasonNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCancelReasonNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionCancelReasonNewResponseEnvelopeSuccess bool

const (
	SubscriptionCancelReasonNewResponseEnvelopeSuccessTrue SubscriptionCancelReasonNewResponseEnvelopeSuccess = true
)

func (r SubscriptionCancelReasonNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionCancelReasonNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubscriptionCancelReasonGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SubscriptionCancelReasonGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo               `json:"errors" api:"required"`
	Messages []shared.ResponseInfo               `json:"messages" api:"required"`
	Result   SubscriptionCancelReasonGetResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success SubscriptionCancelReasonGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    subscriptionCancelReasonGetResponseEnvelopeJSON    `json:"-"`
}

// subscriptionCancelReasonGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SubscriptionCancelReasonGetResponseEnvelope]
type subscriptionCancelReasonGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionCancelReasonGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCancelReasonGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionCancelReasonGetResponseEnvelopeSuccess bool

const (
	SubscriptionCancelReasonGetResponseEnvelopeSuccessTrue SubscriptionCancelReasonGetResponseEnvelopeSuccess = true
)

func (r SubscriptionCancelReasonGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionCancelReasonGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
