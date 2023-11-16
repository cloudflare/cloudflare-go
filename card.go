// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// CardService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCardService] method instead.
type CardService struct {
	Options               []option.RequestOption
	FinancialTransactions *CardFinancialTransactionService
}

// NewCardService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCardService(opts ...option.RequestOption) (r *CardService) {
	r = &CardService{}
	r.Options = opts
	r.FinancialTransactions = NewCardFinancialTransactionService(opts...)
	return
}

// Create a new virtual or physical card. Parameters `pin`, `shipping_address`, and
// `product_id` only apply to physical cards.
func (r *CardService) New(ctx context.Context, params CardNewParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := "cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get card configuration such as spend limit and state.
func (r *CardService) Get(ctx context.Context, cardToken string, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the specified properties of the card. Unsupplied properties will remain
// unchanged. `pin` parameter only applies to physical cards.
//
// _Note: setting a card to a `CLOSED` state is a final action that cannot be
// undone._
func (r *CardService) Update(ctx context.Context, cardToken string, body CardUpdateParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Allow your cardholders to directly add payment cards to the device's digital
// wallet (e.g. Apple Pay) with one touch from your app.
//
// This requires some additional setup and configuration. Please
// [Contact Us](https://acme.com/contact) or your Customer Success representative
// for more information.
func (r *CardService) Provision(ctx context.Context, cardToken string, params CardProvisionParams, opts ...option.RequestOption) (res *CardProvisionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s/provision", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type Card struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// An RFC 3339 timestamp for when the card was created. UTC time zone.
	Created time.Time   `json:"created,required" format:"date-time"`
	Funding CardFunding `json:"funding,required"`
	// Last four digits of the card number.
	LastFour string `json:"last_four,required"`
	// Amount (in cents) to limit approved authorizations. Transaction requests above
	// the spend limit will be declined.
	SpendLimit int64 `json:"spend_limit,required"`
	// Spend limit duration values:
	//
	//   - `ANNUALLY` - Card will authorize transactions up to spend limit in a calendar
	//     year.
	//   - `FOREVER` - Card will authorize only up to spend limit for the entire lifetime
	//     of the card.
	//   - `MONTHLY` - Card will authorize transactions up to spend limit for the
	//     trailing month. Month is calculated as this calendar date one month prior.
	//   - `TRANSACTION` - Card will authorize multiple transactions if each individual
	//     transaction is under the spend limit.
	SpendLimitDuration CardSpendLimitDuration `json:"spend_limit_duration,required"`
	// Card state values:
	//
	//   - `CLOSED` - Card will no longer approve authorizations. Closing a card cannot
	//     be undone.
	//   - `OPEN` - Card will approve authorizations (if they match card and account
	//     parameters).
	//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
	//     time.
	//   - `PENDING_FULFILLMENT` - The initial state for cards of type `PHYSICAL`. The
	//     card is provisioned pending manufacturing and fulfillment. Cards in this state
	//     can accept authorizations for e-commerce purchases, but not for "Card Present"
	//     purchases where the physical card itself is present.
	//   - `PENDING_ACTIVATION` - Each business day at 2pm Eastern Time Zone (ET), cards
	//     of type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card
	//     production warehouse and updated to state `PENDING_ACTIVATION` . Similar to
	//     `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
	//     transactions. API clients should update the card's state to `OPEN` only after
	//     the cardholder confirms receipt of the card.
	//
	// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
	// manufactured.
	State CardState `json:"state,required"`
	// Card types:
	//
	//   - `VIRTUAL` - Card will authorize at any merchant and can be added to a digital
	//     wallet like Apple Pay or Google Pay (if the card program is digital
	//     wallet-enabled).
	//   - `PHYSICAL` - Manufactured and sent to the cardholder. We offer white label
	//     branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe functionality.
	//     Reach out at [acme.com/contact](https://acme.com/contact) for more
	//     information.
	//   - `SINGLE_USE` - Card is closed upon first successful authorization.
	//   - `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first merchant that
	//     successfully authorizes the card.
	Type CardType `json:"type,required"`
	// List of identifiers for the Auth Rule(s) that are applied on the card.
	AuthRuleTokens []string `json:"auth_rule_tokens"`
	// Three digit cvv printed on the back of the card.
	Cvv string `json:"cvv"`
	// Specifies the digital card art to be displayed in the user’s digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Acme to use. See
	// [Flexible Card Art Guide](https://docs.acme.com/docs/about-digital-wallets#flexible-card-art).
	DigitalCardArtToken string `json:"digital_card_art_token" format:"uuid"`
	// Two digit (MM) expiry month.
	ExpMonth string `json:"exp_month"`
	// Four digit (yyyy) expiry year.
	ExpYear string `json:"exp_year"`
	// Hostname of card’s locked merchant (will be empty if not applicable).
	Hostname string `json:"hostname"`
	// Friendly name to identify the card. We recommend against using this field to
	// store JSON data as it can cause unexpected behavior.
	Memo string `json:"memo"`
	// Primary Account Number (PAN) (i.e. the card number). Customers must be PCI
	// compliant to have PAN returned as a field in production. Please contact
	// [support@acme.com](mailto:support@acme.com) for questions.
	Pan  string   `json:"pan"`
	JSON cardJSON `json:"-"`
}

// cardJSON contains the JSON metadata for the struct [Card]
type cardJSON struct {
	Token               apijson.Field
	Created             apijson.Field
	Funding             apijson.Field
	LastFour            apijson.Field
	SpendLimit          apijson.Field
	SpendLimitDuration  apijson.Field
	State               apijson.Field
	Type                apijson.Field
	AuthRuleTokens      apijson.Field
	Cvv                 apijson.Field
	DigitalCardArtToken apijson.Field
	ExpMonth            apijson.Field
	ExpYear             apijson.Field
	Hostname            apijson.Field
	Memo                apijson.Field
	Pan                 apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Card) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardFunding struct {
	// A globally unique identifier for this FundingAccount.
	Token string `json:"token,required" format:"uuid"`
	// An RFC 3339 string representing when this funding source was added to the Acme
	// account. This may be `null`. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// The last 4 digits of the account (e.g. bank account, debit card) associated with
	// this FundingAccount. This may be null.
	LastFour string `json:"last_four,required"`
	// State of funding source.
	//
	// Funding source states:
	//
	//   - `ENABLED` - The funding account is available to use for card creation and
	//     transactions.
	//   - `PENDING` - The funding account is still being verified e.g. bank
	//     micro-deposits verification.
	//   - `DELETED` - The founding account has been deleted.
	State CardFundingState `json:"state,required"`
	// Types of funding source:
	//
	// - `DEPOSITORY_CHECKING` - Bank checking account.
	// - `DEPOSITORY_SAVINGS` - Bank savings account.
	Type CardFundingType `json:"type,required"`
	// Account name identifying the funding source. This may be `null`.
	AccountName string `json:"account_name"`
	// The nickname given to the `FundingAccount` or `null` if it has no nickname.
	Nickname string          `json:"nickname"`
	JSON     cardFundingJSON `json:"-"`
}

// cardFundingJSON contains the JSON metadata for the struct [CardFunding]
type cardFundingJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	LastFour    apijson.Field
	State       apijson.Field
	Type        apijson.Field
	AccountName apijson.Field
	Nickname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of funding source.
//
// Funding source states:
//
//   - `ENABLED` - The funding account is available to use for card creation and
//     transactions.
//   - `PENDING` - The funding account is still being verified e.g. bank
//     micro-deposits verification.
//   - `DELETED` - The founding account has been deleted.
type CardFundingState string

const (
	CardFundingStateEnabled CardFundingState = "ENABLED"
	CardFundingStatePending CardFundingState = "PENDING"
	CardFundingStateDeleted CardFundingState = "DELETED"
)

// Types of funding source:
//
// - `DEPOSITORY_CHECKING` - Bank checking account.
// - `DEPOSITORY_SAVINGS` - Bank savings account.
type CardFundingType string

const (
	CardFundingTypeDepositoryChecking CardFundingType = "DEPOSITORY_CHECKING"
	CardFundingTypeDepositorySavings  CardFundingType = "DEPOSITORY_SAVINGS"
)

// Spend limit duration values:
//
//   - `ANNUALLY` - Card will authorize transactions up to spend limit in a calendar
//     year.
//   - `FOREVER` - Card will authorize only up to spend limit for the entire lifetime
//     of the card.
//   - `MONTHLY` - Card will authorize transactions up to spend limit for the
//     trailing month. Month is calculated as this calendar date one month prior.
//   - `TRANSACTION` - Card will authorize multiple transactions if each individual
//     transaction is under the spend limit.
type CardSpendLimitDuration string

const (
	CardSpendLimitDurationAnnually    CardSpendLimitDuration = "ANNUALLY"
	CardSpendLimitDurationForever     CardSpendLimitDuration = "FOREVER"
	CardSpendLimitDurationMonthly     CardSpendLimitDuration = "MONTHLY"
	CardSpendLimitDurationTransaction CardSpendLimitDuration = "TRANSACTION"
)

// Card state values:
//
//   - `CLOSED` - Card will no longer approve authorizations. Closing a card cannot
//     be undone.
//   - `OPEN` - Card will approve authorizations (if they match card and account
//     parameters).
//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
//     time.
//   - `PENDING_FULFILLMENT` - The initial state for cards of type `PHYSICAL`. The
//     card is provisioned pending manufacturing and fulfillment. Cards in this state
//     can accept authorizations for e-commerce purchases, but not for "Card Present"
//     purchases where the physical card itself is present.
//   - `PENDING_ACTIVATION` - Each business day at 2pm Eastern Time Zone (ET), cards
//     of type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card
//     production warehouse and updated to state `PENDING_ACTIVATION` . Similar to
//     `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
//     transactions. API clients should update the card's state to `OPEN` only after
//     the cardholder confirms receipt of the card.
//
// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
// manufactured.
type CardState string

const (
	CardStateClosed             CardState = "CLOSED"
	CardStateOpen               CardState = "OPEN"
	CardStatePaused             CardState = "PAUSED"
	CardStatePendingActivation  CardState = "PENDING_ACTIVATION"
	CardStatePendingFulfillment CardState = "PENDING_FULFILLMENT"
)

// Card types:
//
//   - `VIRTUAL` - Card will authorize at any merchant and can be added to a digital
//     wallet like Apple Pay or Google Pay (if the card program is digital
//     wallet-enabled).
//   - `PHYSICAL` - Manufactured and sent to the cardholder. We offer white label
//     branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe functionality.
//     Reach out at [acme.com/contact](https://acme.com/contact) for more
//     information.
//   - `SINGLE_USE` - Card is closed upon first successful authorization.
//   - `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first merchant that
//     successfully authorizes the card.
type CardType string

const (
	CardTypeVirtual        CardType = "VIRTUAL"
	CardTypePhysical       CardType = "PHYSICAL"
	CardTypeMerchantLocked CardType = "MERCHANT_LOCKED"
	CardTypeSingleUse      CardType = "SINGLE_USE"
)

type CardProvisionResponse struct {
	ProvisioningPayload string                    `json:"provisioning_payload"`
	JSON                cardProvisionResponseJSON `json:"-"`
}

// cardProvisionResponseJSON contains the JSON metadata for the struct
// [CardProvisionResponse]
type cardProvisionResponseJSON struct {
	ProvisioningPayload apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardProvisionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardNewParams struct {
	// Card types:
	//
	//   - `VIRTUAL` - Card will authorize at any merchant and can be added to a digital
	//     wallet like Apple Pay or Google Pay (if the card program is digital
	//     wallet-enabled).
	//   - `PHYSICAL` - Manufactured and sent to the cardholder. We offer white label
	//     branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe functionality.
	//     Reach out at [acme.com/contact](https://acme.com/contact) for more
	//     information.
	//   - `SINGLE_USE` - Card is closed upon first successful authorization.
	//   - `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first merchant that
	//     successfully authorizes the card.
	Type param.Field[CardNewParamsType] `json:"type,required"`
	// Globally unique identifier for the account that the card will be associated
	// with. Required for programs enrolling users using the
	// [/account_holders endpoint](https://docs.acme.com/docs/account-holders-kyc). See
	// [Managing Your Program](doc:managing-your-program) for more information.
	AccountToken param.Field[string] `json:"account_token" format:"uuid"`
	// For card programs with more than one BIN range. This must be configured with
	// Acme before use. Identifies the card program/BIN range under which to create the
	// card. If omitted, will utilize the program's default `card_program_token`. In
	// Sandbox, use 00000000-0000-0000-1000-000000000000 and
	// 00000000-0000-0000-2000-000000000000 to test creating cards on specific card
	// programs.
	CardProgramToken param.Field[string]               `json:"card_program_token" format:"uuid"`
	Carrier          param.Field[CardNewParamsCarrier] `json:"carrier"`
	// Specifies the digital card art to be displayed in the user’s digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Acme to use. See
	// [Flexible Card Art Guide](https://docs.acme.com/docs/about-digital-wallets#flexible-card-art).
	DigitalCardArtToken param.Field[string] `json:"digital_card_art_token" format:"uuid"`
	// Two digit (MM) expiry month. If neither `exp_month` nor `exp_year` is provided,
	// an expiration date will be generated.
	ExpMonth param.Field[string] `json:"exp_month"`
	// Four digit (yyyy) expiry year. If neither `exp_month` nor `exp_year` is
	// provided, an expiration date will be generated.
	ExpYear param.Field[string] `json:"exp_year"`
	// Friendly name to identify the card. We recommend against using this field to
	// store JSON data as it can cause unexpected behavior.
	Memo param.Field[string] `json:"memo"`
	// Encrypted PIN block (in base64). Only applies to cards of type `PHYSICAL` and
	// `VIRTUAL`. See
	// [Encrypted PIN Block](https://docs.acme.com/docs/cards#encrypted-pin-block-enterprise).
	Pin param.Field[string] `json:"pin"`
	// Only applicable to cards of type `PHYSICAL`. This must be configured with Acme
	// before use. Specifies the configuration (i.e., physical card art) that the card
	// should be manufactured with.
	ProductID       param.Field[string]                       `json:"product_id"`
	ShippingAddress param.Field[CardNewParamsShippingAddress] `json:"shipping_address"`
	// Shipping method for the card. Only applies to cards of type PHYSICAL. Use of
	// options besides `STANDARD` require additional permissions.
	//
	//   - `STANDARD` - USPS regular mail or similar international option, with no
	//     tracking
	//   - `STANDARD_WITH_TRACKING` - USPS regular mail or similar international option,
	//     with tracking
	//   - `PRIORITY` - USPS Priority, 1-3 day shipping, with tracking
	//   - `EXPRESS` - FedEx Express, 3-day shipping, with tracking
	//   - `2_DAY` - FedEx 2-day shipping, with tracking
	//   - `EXPEDITED` - FedEx Standard Overnight or similar international option, with
	//     tracking
	ShippingMethod param.Field[CardNewParamsShippingMethod] `json:"shipping_method"`
	// Amount (in cents) to limit approved authorizations. Transaction requests above
	// the spend limit will be declined. Note that a spend limit of 0 is effectively no
	// limit, and should only be used to reset or remove a prior limit. Only a limit of
	// 1 or above will result in declined transactions due to checks against the card
	// limit.
	SpendLimit param.Field[int64] `json:"spend_limit"`
	// Spend limit duration values:
	//
	//   - `ANNUALLY` - Card will authorize transactions up to spend limit in a calendar
	//     year.
	//   - `FOREVER` - Card will authorize only up to spend limit for the entire lifetime
	//     of the card.
	//   - `MONTHLY` - Card will authorize transactions up to spend limit for the
	//     trailing month. Month is calculated as this calendar date one month prior.
	//   - `TRANSACTION` - Card will authorize multiple transactions if each individual
	//     transaction is under the spend limit.
	SpendLimitDuration param.Field[CardNewParamsSpendLimitDuration] `json:"spend_limit_duration"`
	// Card state values:
	//
	//   - `OPEN` - Card will approve authorizations (if they match card and account
	//     parameters).
	//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
	//     time.
	State          param.Field[CardNewParamsState] `json:"state"`
	IdempotencyKey param.Field[string]             `header:"Idempotency-Key"`
}

func (r CardNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Card types:
//
//   - `VIRTUAL` - Card will authorize at any merchant and can be added to a digital
//     wallet like Apple Pay or Google Pay (if the card program is digital
//     wallet-enabled).
//   - `PHYSICAL` - Manufactured and sent to the cardholder. We offer white label
//     branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe functionality.
//     Reach out at [acme.com/contact](https://acme.com/contact) for more
//     information.
//   - `SINGLE_USE` - Card is closed upon first successful authorization.
//   - `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first merchant that
//     successfully authorizes the card.
type CardNewParamsType string

const (
	CardNewParamsTypeVirtual        CardNewParamsType = "VIRTUAL"
	CardNewParamsTypePhysical       CardNewParamsType = "PHYSICAL"
	CardNewParamsTypeMerchantLocked CardNewParamsType = "MERCHANT_LOCKED"
	CardNewParamsTypeSingleUse      CardNewParamsType = "SINGLE_USE"
)

type CardNewParamsCarrier struct {
	// QR code url to display on the card carrier
	QrCodeURL param.Field[string] `json:"qr_code_url"`
}

func (r CardNewParamsCarrier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardNewParamsShippingAddress struct {
	// Valid USPS routable address.
	Address1 param.Field[string] `json:"address1,required"`
	// City
	City param.Field[string] `json:"city,required"`
	// Uppercase ISO 3166-1 alpha-3 three character abbreviation.
	Country param.Field[string] `json:"country,required"`
	// Customer's first name. This will be the first name printed on the physical card.
	FirstName param.Field[string] `json:"first_name,required"`
	// Customer's surname (family name). This will be the last name printed on the
	// physical card.
	LastName param.Field[string] `json:"last_name,required"`
	// Postal code (formerly zipcode). For US addresses, either five-digit zipcode or
	// nine-digit "ZIP+4".
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Uppercase ISO 3166-2 two character abbreviation for US and CA. Optional with a
	// limit of 24 characters for other countries.
	State param.Field[string] `json:"state,required"`
	// Unit number (if applicable).
	Address2 param.Field[string] `json:"address2"`
	// Email address to be contacted for expedited shipping process purposes. Required
	// if `shipping_method` is `EXPEDITED`.
	Email param.Field[string] `json:"email"`
	// Text to be printed on line two of the physical card. Use of this field requires
	// additional permissions.
	Line2Text param.Field[string] `json:"line2_text"`
	// Cardholder's phone number in E.164 format to be contacted for expedited shipping
	// process purposes. Required if `shipping_method` is `EXPEDITED`.
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r CardNewParamsShippingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Shipping method for the card. Only applies to cards of type PHYSICAL. Use of
// options besides `STANDARD` require additional permissions.
//
//   - `STANDARD` - USPS regular mail or similar international option, with no
//     tracking
//   - `STANDARD_WITH_TRACKING` - USPS regular mail or similar international option,
//     with tracking
//   - `PRIORITY` - USPS Priority, 1-3 day shipping, with tracking
//   - `EXPRESS` - FedEx Express, 3-day shipping, with tracking
//   - `2_DAY` - FedEx 2-day shipping, with tracking
//   - `EXPEDITED` - FedEx Standard Overnight or similar international option, with
//     tracking
type CardNewParamsShippingMethod string

const (
	CardNewParamsShippingMethodStandard             CardNewParamsShippingMethod = "STANDARD"
	CardNewParamsShippingMethodStandardWithTracking CardNewParamsShippingMethod = "STANDARD_WITH_TRACKING"
	CardNewParamsShippingMethodPriority             CardNewParamsShippingMethod = "PRIORITY"
	CardNewParamsShippingMethodExpress              CardNewParamsShippingMethod = "EXPRESS"
	CardNewParamsShippingMethod2Day                 CardNewParamsShippingMethod = "2_DAY"
	CardNewParamsShippingMethodExpedited            CardNewParamsShippingMethod = "EXPEDITED"
)

// Spend limit duration values:
//
//   - `ANNUALLY` - Card will authorize transactions up to spend limit in a calendar
//     year.
//   - `FOREVER` - Card will authorize only up to spend limit for the entire lifetime
//     of the card.
//   - `MONTHLY` - Card will authorize transactions up to spend limit for the
//     trailing month. Month is calculated as this calendar date one month prior.
//   - `TRANSACTION` - Card will authorize multiple transactions if each individual
//     transaction is under the spend limit.
type CardNewParamsSpendLimitDuration string

const (
	CardNewParamsSpendLimitDurationAnnually    CardNewParamsSpendLimitDuration = "ANNUALLY"
	CardNewParamsSpendLimitDurationForever     CardNewParamsSpendLimitDuration = "FOREVER"
	CardNewParamsSpendLimitDurationMonthly     CardNewParamsSpendLimitDuration = "MONTHLY"
	CardNewParamsSpendLimitDurationTransaction CardNewParamsSpendLimitDuration = "TRANSACTION"
)

// Card state values:
//
//   - `OPEN` - Card will approve authorizations (if they match card and account
//     parameters).
//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
//     time.
type CardNewParamsState string

const (
	CardNewParamsStateOpen   CardNewParamsState = "OPEN"
	CardNewParamsStatePaused CardNewParamsState = "PAUSED"
)

type CardUpdateParams struct {
	// Identifier for any Auth Rules that will be applied to transactions taking place
	// with the card.
	AuthRuleToken param.Field[string] `json:"auth_rule_token"`
	// Specifies the digital card art to be displayed in the user’s digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Acme to use. See
	// [Flexible Card Art Guide](https://docs.acme.com/docs/about-digital-wallets#flexible-card-art).
	DigitalCardArtToken param.Field[string] `json:"digital_card_art_token" format:"uuid"`
	// Friendly name to identify the card. We recommend against using this field to
	// store JSON data as it can cause unexpected behavior.
	Memo param.Field[string] `json:"memo"`
	// Encrypted PIN block (in base64). Only applies to cards of type `PHYSICAL` and
	// `VIRTUAL`. See
	// [Encrypted PIN Block](https://docs.acme.com/docs/cards#encrypted-pin-block-enterprise).
	Pin param.Field[string] `json:"pin"`
	// Amount (in cents) to limit approved authorizations. Transaction requests above
	// the spend limit will be declined. Note that a spend limit of 0 is effectively no
	// limit, and should only be used to reset or remove a prior limit. Only a limit of
	// 1 or above will result in declined transactions due to checks against the card
	// limit.
	SpendLimit param.Field[int64] `json:"spend_limit"`
	// Spend limit duration values:
	//
	//   - `ANNUALLY` - Card will authorize transactions up to spend limit in a calendar
	//     year.
	//   - `FOREVER` - Card will authorize only up to spend limit for the entire lifetime
	//     of the card.
	//   - `MONTHLY` - Card will authorize transactions up to spend limit for the
	//     trailing month. Month is calculated as this calendar date one month prior.
	//   - `TRANSACTION` - Card will authorize multiple transactions if each individual
	//     transaction is under the spend limit.
	SpendLimitDuration param.Field[CardUpdateParamsSpendLimitDuration] `json:"spend_limit_duration"`
	// Card state values:
	//
	//   - `CLOSED` - Card will no longer approve authorizations. Closing a card cannot
	//     be undone.
	//   - `OPEN` - Card will approve authorizations (if they match card and account
	//     parameters).
	//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
	//     time.
	State param.Field[CardUpdateParamsState] `json:"state"`
}

func (r CardUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Spend limit duration values:
//
//   - `ANNUALLY` - Card will authorize transactions up to spend limit in a calendar
//     year.
//   - `FOREVER` - Card will authorize only up to spend limit for the entire lifetime
//     of the card.
//   - `MONTHLY` - Card will authorize transactions up to spend limit for the
//     trailing month. Month is calculated as this calendar date one month prior.
//   - `TRANSACTION` - Card will authorize multiple transactions if each individual
//     transaction is under the spend limit.
type CardUpdateParamsSpendLimitDuration string

const (
	CardUpdateParamsSpendLimitDurationAnnually    CardUpdateParamsSpendLimitDuration = "ANNUALLY"
	CardUpdateParamsSpendLimitDurationForever     CardUpdateParamsSpendLimitDuration = "FOREVER"
	CardUpdateParamsSpendLimitDurationMonthly     CardUpdateParamsSpendLimitDuration = "MONTHLY"
	CardUpdateParamsSpendLimitDurationTransaction CardUpdateParamsSpendLimitDuration = "TRANSACTION"
)

// Card state values:
//
//   - `CLOSED` - Card will no longer approve authorizations. Closing a card cannot
//     be undone.
//   - `OPEN` - Card will approve authorizations (if they match card and account
//     parameters).
//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
//     time.
type CardUpdateParamsState string

const (
	CardUpdateParamsStateClosed CardUpdateParamsState = "CLOSED"
	CardUpdateParamsStateOpen   CardUpdateParamsState = "OPEN"
	CardUpdateParamsStatePaused CardUpdateParamsState = "PAUSED"
)

type CardProvisionParams struct {
	// Only applicable if `digital_wallet` is `APPLE_PAY`. Omit to receive only
	// `activationData` in the response. Apple's public leaf certificate. Base64
	// encoded in PEM format with headers `(-----BEGIN CERTIFICATE-----)` and trailers
	// omitted. Provided by the device's wallet.
	Certificate param.Field[string] `json:"certificate" format:"byte"`
	// Name of digital wallet provider.
	DigitalWallet param.Field[CardProvisionParamsDigitalWallet] `json:"digital_wallet"`
	// Only applicable if `digital_wallet` is `APPLE_PAY`. Omit to receive only
	// `activationData` in the response. Base64 cryptographic nonce provided by the
	// device's wallet.
	Nonce param.Field[string] `json:"nonce" format:"byte"`
	// Only applicable if `digital_wallet` is `APPLE_PAY`. Omit to receive only
	// `activationData` in the response. Base64 cryptographic nonce provided by the
	// device's wallet.
	NonceSignature param.Field[string] `json:"nonce_signature" format:"byte"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

func (r CardProvisionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Name of digital wallet provider.
type CardProvisionParamsDigitalWallet string

const (
	CardProvisionParamsDigitalWalletApplePay   CardProvisionParamsDigitalWallet = "APPLE_PAY"
	CardProvisionParamsDigitalWalletGooglePay  CardProvisionParamsDigitalWallet = "GOOGLE_PAY"
	CardProvisionParamsDigitalWalletSamsungPay CardProvisionParamsDigitalWallet = "SAMSUNG_PAY"
)
