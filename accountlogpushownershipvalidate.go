// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountLogpushOwnershipValidateService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountLogpushOwnershipValidateService] method instead.
type AccountLogpushOwnershipValidateService struct {
	Options []option.RequestOption
}

// NewAccountLogpushOwnershipValidateService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLogpushOwnershipValidateService(opts ...option.RequestOption) (r *AccountLogpushOwnershipValidateService) {
	r = &AccountLogpushOwnershipValidateService{}
	r.Options = opts
	return
}

// Validates ownership challenge of the destination.
func (r *AccountLogpushOwnershipValidateService) PostAccountsAccountIdentifierLogpushOwnershipValidate(ctx context.Context, accountIdentifier string, body AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateParams, opts ...option.RequestOption) (res *AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/ownership/validate", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponse struct {
	Errors   []AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseError   `json:"errors"`
	Messages []AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseMessage `json:"messages"`
	Result   AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseSuccess `json:"success"`
	JSON    accountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseJSON    `json:"-"`
}

// accountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseJSON
// contains the JSON metadata for the struct
// [AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponse]
type accountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseError struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    accountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseErrorJSON `json:"-"`
}

// accountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseError]
type accountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseMessage struct {
	Code    int64                                                                                                   `json:"code,required"`
	Message string                                                                                                  `json:"message,required"`
	JSON    accountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseMessageJSON `json:"-"`
}

// accountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseMessage]
type accountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseResult struct {
	Valid bool                                                                                                   `json:"valid"`
	JSON  accountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseResultJSON `json:"-"`
}

// accountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseResultJSON
// contains the JSON metadata for the struct
// [AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseResult]
type accountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseResultJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseSuccess bool

const (
	AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseSuccessTrue AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseSuccess = true
)

type AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
