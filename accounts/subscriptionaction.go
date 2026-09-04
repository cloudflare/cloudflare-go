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

// SubscriptionActionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionActionService] method instead.
type SubscriptionActionService struct {
	Options []option.RequestOption
}

// NewSubscriptionActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSubscriptionActionService(opts ...option.RequestOption) (r *SubscriptionActionService) {
	r = &SubscriptionActionService{}
	r.Options = opts
	return
}

// Smartly applies the incoming subscription into the lifecycle of the
// subscription.
func (r *SubscriptionActionService) Append(ctx context.Context, subscriptionIdentifier string, params SubscriptionActionAppendParams, opts ...option.RequestOption) (res *shared.Subscription, err error) {
	var env SubscriptionActionAppendResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if subscriptionIdentifier == "" {
		err = errors.New("missing required subscription_identifier parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/subscriptions/%s/action/append", params.AccountID, subscriptionIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SubscriptionActionAppendParams struct {
	// Identifier
	AccountID    param.Field[string]      `path:"account_id" api:"required"`
	Subscription shared.SubscriptionParam `json:"subscription" api:"required"`
}

func (r SubscriptionActionAppendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Subscription)
}

type SubscriptionActionAppendResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   shared.Subscription   `json:"result" api:"required"`
	// Whether the API call was successful
	Success SubscriptionActionAppendResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    subscriptionActionAppendResponseEnvelopeJSON    `json:"-"`
}

// subscriptionActionAppendResponseEnvelopeJSON contains the JSON metadata for the
// struct [SubscriptionActionAppendResponseEnvelope]
type subscriptionActionAppendResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionActionAppendResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionActionAppendResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionActionAppendResponseEnvelopeSuccess bool

const (
	SubscriptionActionAppendResponseEnvelopeSuccessTrue SubscriptionActionAppendResponseEnvelopeSuccess = true
)

func (r SubscriptionActionAppendResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionActionAppendResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
