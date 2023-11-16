// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// CardFinancialTransactionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewCardFinancialTransactionService] method instead.
type CardFinancialTransactionService struct {
	Options []option.RequestOption
}

// NewCardFinancialTransactionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCardFinancialTransactionService(opts ...option.RequestOption) (r *CardFinancialTransactionService) {
	r = &CardFinancialTransactionService{}
	r.Options = opts
	return
}

// Get the card financial transaction for the provided token.
func (r *CardFinancialTransactionService) Get(ctx context.Context, cardToken string, financialTransactionToken string, opts ...option.RequestOption) (res *FinancialTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s/financial_transactions/%s", cardToken, financialTransactionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FinancialTransaction struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Status types:
	//
	//   - `CARD` - Issuing card transaction.
	//   - `ACH` - Transaction over ACH.
	//   - `TRANSFER` - Internal transfer of funds between financial accounts in your
	//     program.
	Category FinancialTransactionCategory `json:"category,required"`
	// Date and time when the financial transaction first occurred. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the settling currency of the transaction.
	Currency string `json:"currency,required"`
	// A string that provides a description of the financial transaction; may be useful
	// to display to users.
	Descriptor string `json:"descriptor,required"`
	// A list of all financial events that have modified this financial transaction.
	Events []FinancialTransactionEvent `json:"events,required"`
	// Pending amount of the transaction in the currency's smallest unit (e.g., cents),
	// including any acquirer fees. The value of this field will go to zero over time
	// once the financial transaction is settled.
	PendingAmount int64 `json:"pending_amount,required"`
	// APPROVED transactions were successful while DECLINED transactions were declined
	// by user, Acme, or the network.
	Result FinancialTransactionResult `json:"result,required"`
	// Amount of the transaction that has been settled in the currency's smallest unit
	// (e.g., cents), including any acquirer fees. This may change over time.
	SettledAmount int64 `json:"settled_amount,required"`
	// Status types:
	//
	//   - `DECLINED` - The card transaction was declined.
	//   - `EXPIRED` - Acme reversed the card authorization as it has passed its
	//     expiration time.
	//   - `PENDING` - Authorization is pending completion from the merchant or pending
	//     release from ACH hold period
	//   - `SETTLED` - The financial transaction is completed.
	//   - `VOIDED` - The merchant has voided the previously pending card authorization.
	Status FinancialTransactionStatus `json:"status,required"`
	// Date and time when the financial transaction was last updated. UTC time zone.
	Updated time.Time                `json:"updated,required" format:"date-time"`
	JSON    financialTransactionJSON `json:"-"`
}

// financialTransactionJSON contains the JSON metadata for the struct
// [FinancialTransaction]
type financialTransactionJSON struct {
	Token         apijson.Field
	Category      apijson.Field
	Created       apijson.Field
	Currency      apijson.Field
	Descriptor    apijson.Field
	Events        apijson.Field
	PendingAmount apijson.Field
	Result        apijson.Field
	SettledAmount apijson.Field
	Status        apijson.Field
	Updated       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FinancialTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status types:
//
//   - `CARD` - Issuing card transaction.
//   - `ACH` - Transaction over ACH.
//   - `TRANSFER` - Internal transfer of funds between financial accounts in your
//     program.
type FinancialTransactionCategory string

const (
	FinancialTransactionCategoryCard     FinancialTransactionCategory = "CARD"
	FinancialTransactionCategoryACH      FinancialTransactionCategory = "ACH"
	FinancialTransactionCategoryTransfer FinancialTransactionCategory = "TRANSFER"
)

type FinancialTransactionEvent struct {
	// Globally unique identifier.
	Token string `json:"token" format:"uuid"`
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount"`
	// Date and time when the financial event occurred. UTC time zone.
	Created time.Time `json:"created" format:"date-time"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Acme, or the network.
	Result FinancialTransactionEventsResult `json:"result"`
	// Event types:
	//
	//   - `ACH_INSUFFICIENT_FUNDS` - Attempted ACH origination declined due to
	//     insufficient balance.
	//   - `ACH_ORIGINATION_PENDING` - ACH origination pending release from an ACH hold.
	//   - `ACH_ORIGINATION_RELEASED` - ACH origination released from pending to
	//     available balance.
	//   - `ACH_RECEIPT_PENDING` - ACH receipt pending release from an ACH holder.
	//   - `ACH_RECEIPT_RELEASED` - ACH receipt released from pending to available
	//     balance.
	//   - `ACH_RETURN` - ACH origination returned by the Receiving Depository Financial
	//     Institution.
	//   - `AUTHORIZATION` - Authorize a card transaction.
	//   - `AUTHORIZATION_ADVICE` - Advice on a card transaction.
	//   - `AUTHORIZATION_EXPIRY` - Card Authorization has expired and reversed by Acme.
	//   - `AUTHORIZATION_REVERSAL` - Card Authorization was reversed by the merchant.
	//   - `BALANCE_INQUIRY` - A card balance inquiry (typically a $0 authorization) has
	//     occurred on a card.
	//   - `CLEARING` - Card Transaction is settled.
	//   - `CORRECTION_DEBIT` - Manual card transaction correction (Debit).
	//   - `CORRECTION_CREDIT` - Manual card transaction correction (Credit).
	//   - `CREDIT_AUTHORIZATION` - A refund or credit card authorization from a
	//     merchant.
	//   - `CREDIT_AUTHORIZATION_ADVICE` - A credit card authorization was approved on
	//     your behalf by the network.
	//   - `FINANCIAL_AUTHORIZATION` - A request from a merchant to debit card funds
	//     without additional clearing.
	//   - `FINANCIAL_CREDIT_AUTHORIZATION` - A request from a merchant to refund or
	//     credit card funds without additional clearing.
	//   - `RETURN` - A card refund has been processed on the transaction.
	//   - `RETURN_REVERSAL` - A card refund has been reversed (e.g., when a merchant
	//     reverses an incorrect refund).
	//   - `TRANSFER` - Successful internal transfer of funds between financial accounts.
	//   - `TRANSFER_INSUFFICIENT_FUNDS` - Declined internl transfer of funds due to
	//     insufficient balance of the sender.
	Type FinancialTransactionEventsType `json:"type"`
	JSON financialTransactionEventJSON  `json:"-"`
}

// financialTransactionEventJSON contains the JSON metadata for the struct
// [FinancialTransactionEvent]
type financialTransactionEventJSON struct {
	Token       apijson.Field
	Amount      apijson.Field
	Created     apijson.Field
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinancialTransactionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// APPROVED financial events were successful while DECLINED financial events were
// declined by user, Acme, or the network.
type FinancialTransactionEventsResult string

const (
	FinancialTransactionEventsResultApproved FinancialTransactionEventsResult = "APPROVED"
	FinancialTransactionEventsResultDeclined FinancialTransactionEventsResult = "DECLINED"
)

// Event types:
//
//   - `ACH_INSUFFICIENT_FUNDS` - Attempted ACH origination declined due to
//     insufficient balance.
//   - `ACH_ORIGINATION_PENDING` - ACH origination pending release from an ACH hold.
//   - `ACH_ORIGINATION_RELEASED` - ACH origination released from pending to
//     available balance.
//   - `ACH_RECEIPT_PENDING` - ACH receipt pending release from an ACH holder.
//   - `ACH_RECEIPT_RELEASED` - ACH receipt released from pending to available
//     balance.
//   - `ACH_RETURN` - ACH origination returned by the Receiving Depository Financial
//     Institution.
//   - `AUTHORIZATION` - Authorize a card transaction.
//   - `AUTHORIZATION_ADVICE` - Advice on a card transaction.
//   - `AUTHORIZATION_EXPIRY` - Card Authorization has expired and reversed by Acme.
//   - `AUTHORIZATION_REVERSAL` - Card Authorization was reversed by the merchant.
//   - `BALANCE_INQUIRY` - A card balance inquiry (typically a $0 authorization) has
//     occurred on a card.
//   - `CLEARING` - Card Transaction is settled.
//   - `CORRECTION_DEBIT` - Manual card transaction correction (Debit).
//   - `CORRECTION_CREDIT` - Manual card transaction correction (Credit).
//   - `CREDIT_AUTHORIZATION` - A refund or credit card authorization from a
//     merchant.
//   - `CREDIT_AUTHORIZATION_ADVICE` - A credit card authorization was approved on
//     your behalf by the network.
//   - `FINANCIAL_AUTHORIZATION` - A request from a merchant to debit card funds
//     without additional clearing.
//   - `FINANCIAL_CREDIT_AUTHORIZATION` - A request from a merchant to refund or
//     credit card funds without additional clearing.
//   - `RETURN` - A card refund has been processed on the transaction.
//   - `RETURN_REVERSAL` - A card refund has been reversed (e.g., when a merchant
//     reverses an incorrect refund).
//   - `TRANSFER` - Successful internal transfer of funds between financial accounts.
//   - `TRANSFER_INSUFFICIENT_FUNDS` - Declined internl transfer of funds due to
//     insufficient balance of the sender.
type FinancialTransactionEventsType string

const (
	FinancialTransactionEventsTypeACHInsufficientFunds         FinancialTransactionEventsType = "ACH_INSUFFICIENT_FUNDS"
	FinancialTransactionEventsTypeACHOriginationPending        FinancialTransactionEventsType = "ACH_ORIGINATION_PENDING"
	FinancialTransactionEventsTypeACHOriginationReleased       FinancialTransactionEventsType = "ACH_ORIGINATION_RELEASED"
	FinancialTransactionEventsTypeACHReceiptPending            FinancialTransactionEventsType = "ACH_RECEIPT_PENDING"
	FinancialTransactionEventsTypeACHReceiptReleased           FinancialTransactionEventsType = "ACH_RECEIPT_RELEASED"
	FinancialTransactionEventsTypeACHReturn                    FinancialTransactionEventsType = "ACH_RETURN"
	FinancialTransactionEventsTypeAuthorization                FinancialTransactionEventsType = "AUTHORIZATION"
	FinancialTransactionEventsTypeAuthorizationAdvice          FinancialTransactionEventsType = "AUTHORIZATION_ADVICE"
	FinancialTransactionEventsTypeAuthorizationExpiry          FinancialTransactionEventsType = "AUTHORIZATION_EXPIRY"
	FinancialTransactionEventsTypeAuthorizationReversal        FinancialTransactionEventsType = "AUTHORIZATION_REVERSAL"
	FinancialTransactionEventsTypeBalanceInquiry               FinancialTransactionEventsType = "BALANCE_INQUIRY"
	FinancialTransactionEventsTypeClearing                     FinancialTransactionEventsType = "CLEARING"
	FinancialTransactionEventsTypeCorrectionDebit              FinancialTransactionEventsType = "CORRECTION_DEBIT"
	FinancialTransactionEventsTypeCorrectionCredit             FinancialTransactionEventsType = "CORRECTION_CREDIT"
	FinancialTransactionEventsTypeCreditAuthorization          FinancialTransactionEventsType = "CREDIT_AUTHORIZATION"
	FinancialTransactionEventsTypeCreditAuthorizationAdvice    FinancialTransactionEventsType = "CREDIT_AUTHORIZATION_ADVICE"
	FinancialTransactionEventsTypeFinancialAuthorization       FinancialTransactionEventsType = "FINANCIAL_AUTHORIZATION"
	FinancialTransactionEventsTypeFinancialCreditAuthorization FinancialTransactionEventsType = "FINANCIAL_CREDIT_AUTHORIZATION"
	FinancialTransactionEventsTypeReturn                       FinancialTransactionEventsType = "RETURN"
	FinancialTransactionEventsTypeReturnReversal               FinancialTransactionEventsType = "RETURN_REVERSAL"
	FinancialTransactionEventsTypeTransfer                     FinancialTransactionEventsType = "TRANSFER"
	FinancialTransactionEventsTypeTransferInsufficientFunds    FinancialTransactionEventsType = "TRANSFER_INSUFFICIENT_FUNDS"
)

// APPROVED transactions were successful while DECLINED transactions were declined
// by user, Acme, or the network.
type FinancialTransactionResult string

const (
	FinancialTransactionResultApproved FinancialTransactionResult = "APPROVED"
	FinancialTransactionResultDeclined FinancialTransactionResult = "DECLINED"
)

// Status types:
//
//   - `DECLINED` - The card transaction was declined.
//   - `EXPIRED` - Acme reversed the card authorization as it has passed its
//     expiration time.
//   - `PENDING` - Authorization is pending completion from the merchant or pending
//     release from ACH hold period
//   - `SETTLED` - The financial transaction is completed.
//   - `VOIDED` - The merchant has voided the previously pending card authorization.
type FinancialTransactionStatus string

const (
	FinancialTransactionStatusDeclined FinancialTransactionStatus = "DECLINED"
	FinancialTransactionStatusExpired  FinancialTransactionStatus = "EXPIRED"
	FinancialTransactionStatusPending  FinancialTransactionStatus = "PENDING"
	FinancialTransactionStatusSettled  FinancialTransactionStatus = "SETTLED"
	FinancialTransactionStatusVoided   FinancialTransactionStatus = "VOIDED"
)
