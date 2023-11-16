// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// AccountService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccountService] method instead.
type AccountService struct {
	Options             []option.RequestOption
	CreditConfiguration *AccountCreditConfigurationService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.CreditConfiguration = NewAccountCreditConfigurationService(opts...)
	return
}

// Get account configuration such as spend limits.
func (r *AccountService) Get(ctx context.Context, accountToken string, opts ...option.RequestOption) (res *AccountConfiguration, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", accountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update account configuration such as spend limits and verification address. Can
// only be run on accounts that are part of the program managed by this API key.
//
// Accounts that are in the `PAUSED` state will not be able to transact or create
// new cards.
func (r *AccountService) Update(ctx context.Context, accountToken string, body AccountUpdateParams, opts ...option.RequestOption) (res *AccountConfiguration, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", accountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AccountConfiguration struct {
	// Globally unique identifier for the account. This is the same as the
	// account_token returned by the enroll endpoint. If using this parameter, do not
	// include pagination.
	Token string `json:"token,required" format:"uuid"`
	// Spend limit information for the user containing the daily, monthly, and lifetime
	// spend limit of the account. Any charges to a card owned by this account will be
	// declined once their transaction volume has surpassed the value in the applicable
	// time limit (rolling). A lifetime limit of 0 indicates that the lifetime limit
	// feature is disabled.
	SpendLimit AccountConfigurationSpendLimit `json:"spend_limit,required"`
	// Account state:
	//
	//   - `ACTIVE` - Account is able to transact and create new cards.
	//   - `PAUSED` - Account will not be able to transact or create new cards. It can be
	//     set back to `ACTIVE`.
	//   - `CLOSED` - Account will permanently not be able to transact or create new
	//     cards.
	State         AccountConfigurationState         `json:"state,required"`
	AccountHolder AccountConfigurationAccountHolder `json:"account_holder"`
	// List of identifiers for the Auth Rule(s) that are applied on the account.
	AuthRuleTokens      []string                                `json:"auth_rule_tokens"`
	VerificationAddress AccountConfigurationVerificationAddress `json:"verification_address"`
	JSON                accountConfigurationJSON                `json:"-"`
}

// accountConfigurationJSON contains the JSON metadata for the struct
// [AccountConfiguration]
type accountConfigurationJSON struct {
	Token               apijson.Field
	SpendLimit          apijson.Field
	State               apijson.Field
	AccountHolder       apijson.Field
	AuthRuleTokens      apijson.Field
	VerificationAddress apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Spend limit information for the user containing the daily, monthly, and lifetime
// spend limit of the account. Any charges to a card owned by this account will be
// declined once their transaction volume has surpassed the value in the applicable
// time limit (rolling). A lifetime limit of 0 indicates that the lifetime limit
// feature is disabled.
type AccountConfigurationSpendLimit struct {
	// Daily spend limit (in cents).
	Daily int64 `json:"daily,required"`
	// Total spend limit over account lifetime (in cents).
	Lifetime int64 `json:"lifetime,required"`
	// Monthly spend limit (in cents).
	Monthly int64                              `json:"monthly,required"`
	JSON    accountConfigurationSpendLimitJSON `json:"-"`
}

// accountConfigurationSpendLimitJSON contains the JSON metadata for the struct
// [AccountConfigurationSpendLimit]
type accountConfigurationSpendLimitJSON struct {
	Daily       apijson.Field
	Lifetime    apijson.Field
	Monthly     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountConfigurationSpendLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Account state:
//
//   - `ACTIVE` - Account is able to transact and create new cards.
//   - `PAUSED` - Account will not be able to transact or create new cards. It can be
//     set back to `ACTIVE`.
//   - `CLOSED` - Account will permanently not be able to transact or create new
//     cards.
type AccountConfigurationState string

const (
	AccountConfigurationStateActive AccountConfigurationState = "ACTIVE"
	AccountConfigurationStatePaused AccountConfigurationState = "PAUSED"
	AccountConfigurationStateClosed AccountConfigurationState = "CLOSED"
)

type AccountConfigurationAccountHolder struct {
	// Globally unique identifier for the account holder.
	Token string `json:"token,required"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Account_token of the enrolled business associated with an
	// enrolled AUTHORIZED_USER individual.
	BusinessAccountToken string `json:"business_account_token,required"`
	// Email address.
	Email string `json:"email,required"`
	// Phone number of the individual.
	PhoneNumber string                                `json:"phone_number,required"`
	JSON        accountConfigurationAccountHolderJSON `json:"-"`
}

// accountConfigurationAccountHolderJSON contains the JSON metadata for the struct
// [AccountConfigurationAccountHolder]
type accountConfigurationAccountHolderJSON struct {
	Token                apijson.Field
	BusinessAccountToken apijson.Field
	Email                apijson.Field
	PhoneNumber          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountConfigurationAccountHolder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountConfigurationVerificationAddress struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
	// City name.
	City string `json:"city,required"`
	// Country name. Only USA is currently supported.
	Country string `json:"country,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode string `json:"postal_code,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State string `json:"state,required"`
	// Unit or apartment number (if applicable).
	Address2 string                                      `json:"address2"`
	JSON     accountConfigurationVerificationAddressJSON `json:"-"`
}

// accountConfigurationVerificationAddressJSON contains the JSON metadata for the
// struct [AccountConfigurationVerificationAddress]
type accountConfigurationVerificationAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountConfigurationVerificationAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateParams struct {
	// Amount (in cents) for the account's daily spend limit. By default the daily
	// spend limit is set to $1,250.
	DailySpendLimit param.Field[int64] `json:"daily_spend_limit"`
	// Amount (in cents) for the account's lifetime spend limit. Once this limit is
	// reached, no transactions will be accepted on any card created for this account
	// until the limit is updated. Note that a spend limit of 0 is effectively no
	// limit, and should only be used to reset or remove a prior limit. Only a limit of
	// 1 or above will result in declined transactions due to checks against the
	// account limit. This behavior differs from the daily spend limit and the monthly
	// spend limit.
	LifetimeSpendLimit param.Field[int64] `json:"lifetime_spend_limit"`
	// Amount (in cents) for the account's monthly spend limit. By default the monthly
	// spend limit is set to $5,000.
	MonthlySpendLimit param.Field[int64] `json:"monthly_spend_limit"`
	// Account states.
	State param.Field[AccountUpdateParamsState] `json:"state"`
	// Address used during Address Verification Service (AVS) checks during
	// transactions if enabled via Auth Rules.
	VerificationAddress param.Field[AccountUpdateParamsVerificationAddress] `json:"verification_address"`
}

func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account states.
type AccountUpdateParamsState string

const (
	AccountUpdateParamsStateActive AccountUpdateParamsState = "ACTIVE"
	AccountUpdateParamsStatePaused AccountUpdateParamsState = "PAUSED"
)

// Address used during Address Verification Service (AVS) checks during
// transactions if enabled via Auth Rules.
type AccountUpdateParamsVerificationAddress struct {
	Address1   param.Field[string] `json:"address1"`
	Address2   param.Field[string] `json:"address2"`
	City       param.Field[string] `json:"city"`
	Country    param.Field[string] `json:"country"`
	PostalCode param.Field[string] `json:"postal_code"`
	State      param.Field[string] `json:"state"`
}

func (r AccountUpdateParamsVerificationAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
