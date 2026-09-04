// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package billing

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

// CreditService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditService] method instead.
type CreditService struct {
	Options []option.RequestOption
}

// NewCreditService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCreditService(opts ...option.RequestOption) (r *CreditService) {
	r = &CreditService{}
	r.Options = opts
	return
}

// Gets the credit balance and eligibility for an account.
func (r *CreditService) Get(ctx context.Context, query CreditGetParams, opts ...option.RequestOption) (res *CreditGetResponse, err error) {
	var env CreditGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/billing/credits", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type CreditGetResponse struct {
	// The confirmed credit balance in cents.
	ConfirmedBalanceCents int64 `json:"confirmed_balance_cents"`
	// Currency of the credit balance.
	Currency string `json:"currency"`
	// Days remaining until the credits expire.
	DaysRemaining int64 `json:"days_remaining"`
	// Whether the account is eligible to receive credits.
	Eligible bool `json:"eligible"`
	// Whether a credit record exists for the account.
	HasRecord bool `json:"has_record"`
	// The original credit amount in cents.
	OriginalAmountCents int64 `json:"original_amount_cents"`
	// Percentage of the original credit amount consumed.
	PercentConsumed float64 `json:"percent_consumed"`
	// Projected date when the credits will be depleted.
	ProjectedDepletionDate time.Time `json:"projected_depletion_date" format:"date-time"`
	// When the credits become valid.
	ValidFrom time.Time `json:"valid_from" format:"date-time"`
	// When the credits expire.
	ValidTo time.Time             `json:"valid_to" format:"date-time"`
	JSON    creditGetResponseJSON `json:"-"`
}

// creditGetResponseJSON contains the JSON metadata for the struct
// [CreditGetResponse]
type creditGetResponseJSON struct {
	ConfirmedBalanceCents  apijson.Field
	Currency               apijson.Field
	DaysRemaining          apijson.Field
	Eligible               apijson.Field
	HasRecord              apijson.Field
	OriginalAmountCents    apijson.Field
	PercentConsumed        apijson.Field
	ProjectedDepletionDate apijson.Field
	ValidFrom              apijson.Field
	ValidTo                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CreditGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGetResponseJSON) RawJSON() string {
	return r.raw
}

type CreditGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CreditGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   CreditGetResponse     `json:"result" api:"required"`
	// Whether the API call was successful
	Success CreditGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    creditGetResponseEnvelopeJSON    `json:"-"`
}

// creditGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CreditGetResponseEnvelope]
type creditGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CreditGetResponseEnvelopeSuccess bool

const (
	CreditGetResponseEnvelopeSuccessTrue CreditGetResponseEnvelopeSuccess = true
)

func (r CreditGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CreditGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
