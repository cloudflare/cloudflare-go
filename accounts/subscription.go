// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

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
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// SubscriptionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r *SubscriptionService) {
	r = &SubscriptionService{}
	r.Options = opts
	return
}

// Updates an account subscription.
func (r *SubscriptionService) Update(ctx context.Context, subscriptionIdentifier string, params SubscriptionUpdateParams, opts ...option.RequestOption) (res *shared.Subscription, err error) {
	var env SubscriptionUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if subscriptionIdentifier == "" {
		err = errors.New("missing required subscription_identifier parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/subscriptions/%s", params.AccountID, subscriptionIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Deletes an account's subscription.
func (r *SubscriptionService) Delete(ctx context.Context, subscriptionIdentifier string, body SubscriptionDeleteParams, opts ...option.RequestOption) (res *SubscriptionDeleteResponse, err error) {
	var env SubscriptionDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if subscriptionIdentifier == "" {
		err = errors.New("missing required subscription_identifier parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/subscriptions/%s", body.AccountID, subscriptionIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SubscriptionDeleteResponse struct {
	// Subscription identifier tag.
	SubscriptionID string                         `json:"subscription_id"`
	JSON           subscriptionDeleteResponseJSON `json:"-"`
}

// subscriptionDeleteResponseJSON contains the JSON metadata for the struct
// [SubscriptionDeleteResponse]
type subscriptionDeleteResponseJSON struct {
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionUpdateParams struct {
	// Identifier
	AccountID    param.Field[string]      `path:"account_id" api:"required"`
	Subscription shared.SubscriptionParam `json:"subscription" api:"required"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Subscription)
}

type SubscriptionUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   shared.Subscription   `json:"result" api:"required"`
	// Whether the API call was successful
	Success SubscriptionUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    subscriptionUpdateResponseEnvelopeJSON    `json:"-"`
}

// subscriptionUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionUpdateResponseEnvelope]
type subscriptionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionUpdateResponseEnvelopeSuccess bool

const (
	SubscriptionUpdateResponseEnvelopeSuccessTrue SubscriptionUpdateResponseEnvelopeSuccess = true
)

func (r SubscriptionUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubscriptionDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SubscriptionDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors" api:"required"`
	Messages []shared.ResponseInfo      `json:"messages" api:"required"`
	Result   SubscriptionDeleteResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success SubscriptionDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    subscriptionDeleteResponseEnvelopeJSON    `json:"-"`
}

// subscriptionDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionDeleteResponseEnvelope]
type subscriptionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionDeleteResponseEnvelopeSuccess bool

const (
	SubscriptionDeleteResponseEnvelopeSuccessTrue SubscriptionDeleteResponseEnvelopeSuccess = true
)

func (r SubscriptionDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
