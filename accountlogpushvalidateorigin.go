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

// AccountLogpushValidateOriginService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountLogpushValidateOriginService] method instead.
type AccountLogpushValidateOriginService struct {
	Options []option.RequestOption
}

// NewAccountLogpushValidateOriginService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLogpushValidateOriginService(opts ...option.RequestOption) (r *AccountLogpushValidateOriginService) {
	r = &AccountLogpushValidateOriginService{}
	r.Options = opts
	return
}

// Validates logpull origin with logpull_options.
func (r *AccountLogpushValidateOriginService) PostAccountsAccountIdentifierLogpushValidateOrigin(ctx context.Context, accountIdentifier string, body AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginParams, opts ...option.RequestOption) (res *AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/validate/origin", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponse struct {
	Errors   []AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseError   `json:"errors"`
	Messages []AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseMessage `json:"messages"`
	Result   AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseSuccess `json:"success"`
	JSON    accountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseJSON    `json:"-"`
}

// accountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseJSON
// contains the JSON metadata for the struct
// [AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponse]
type accountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseError struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    accountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseErrorJSON `json:"-"`
}

// accountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseError]
type accountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseMessage struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    accountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseMessageJSON `json:"-"`
}

// accountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseMessage]
type accountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseResult struct {
	Message string                                                                                           `json:"message"`
	Valid   bool                                                                                             `json:"valid"`
	JSON    accountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseResultJSON `json:"-"`
}

// accountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseResultJSON
// contains the JSON metadata for the struct
// [AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseResult]
type accountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseResultJSON struct {
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseSuccess bool

const (
	AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseSuccessTrue AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseSuccess = true
)

type AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginParams struct {
	// Configuration string. It specifies things like requested fields and timestamp
	// formats. If migrating from the logpull api, copy the url (full url or just the
	// query string) of your call here, and logpush will keep on making this call for
	// you, setting start and end times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options,required" format:"uri-reference"`
}

func (r AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
