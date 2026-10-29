// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// SubscriptionBulkService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionBulkService] method instead.
type SubscriptionBulkService struct {
	Options []option.RequestOption
}

// NewSubscriptionBulkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSubscriptionBulkService(opts ...option.RequestOption) (r *SubscriptionBulkService) {
	r = &SubscriptionBulkService{}
	r.Options = opts
	return
}

// Creates multiple subscriptions for an account in a single request.
func (r *SubscriptionBulkService) New(ctx context.Context, params SubscriptionBulkNewParams, opts ...option.RequestOption) (res *SubscriptionBulkNewResponse, err error) {
	var env SubscriptionBulkNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/bulk/subscriptions", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SubscriptionBulkNewResponse struct {
	ClientSecrets []string                        `json:"client_secrets" api:"nullable"`
	Subscriptions []shared.Subscription           `json:"subscriptions"`
	JSON          subscriptionBulkNewResponseJSON `json:"-"`
}

// subscriptionBulkNewResponseJSON contains the JSON metadata for the struct
// [SubscriptionBulkNewResponse]
type subscriptionBulkNewResponseJSON struct {
	ClientSecrets apijson.Field
	Subscriptions apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SubscriptionBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionBulkNewParams struct {
	// Identifier
	AccountID       param.Field[string]                     `path:"account_id" api:"required"`
	IdempKey        param.Field[string]                     `query:"idemp_key"`
	CouponCode      param.Field[string]                     `json:"coupon_code"`
	PaymentHoldID   param.Field[int64]                      `json:"payment_hold_id"`
	Subscriptions   param.Field[[]shared.SubscriptionParam] `json:"subscriptions"`
	UserIsOnSession param.Field[bool]                       `json:"user_is_on_session"`
}

func (r SubscriptionBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SubscriptionBulkNewParams]'s query parameters as
// `url.Values`.
func (r SubscriptionBulkNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SubscriptionBulkNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors" api:"required"`
	Messages []shared.ResponseInfo       `json:"messages" api:"required"`
	Result   SubscriptionBulkNewResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success SubscriptionBulkNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    subscriptionBulkNewResponseEnvelopeJSON    `json:"-"`
}

// subscriptionBulkNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SubscriptionBulkNewResponseEnvelope]
type subscriptionBulkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionBulkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionBulkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionBulkNewResponseEnvelopeSuccess bool

const (
	SubscriptionBulkNewResponseEnvelopeSuccessTrue SubscriptionBulkNewResponseEnvelopeSuccess = true
)

func (r SubscriptionBulkNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionBulkNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
