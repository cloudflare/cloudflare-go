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

// AccountLogpushOwnershipService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountLogpushOwnershipService] method instead.
type AccountLogpushOwnershipService struct {
	Options   []option.RequestOption
	Validates *AccountLogpushOwnershipValidateService
}

// NewAccountLogpushOwnershipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLogpushOwnershipService(opts ...option.RequestOption) (r *AccountLogpushOwnershipService) {
	r = &AccountLogpushOwnershipService{}
	r.Options = opts
	r.Validates = NewAccountLogpushOwnershipValidateService(opts...)
	return
}

// Gets a new ownership challenge sent to your destination.
func (r *AccountLogpushOwnershipService) PostAccountsAccountIdentifierLogpushOwnership(ctx context.Context, accountIdentifier string, body AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipParams, opts ...option.RequestOption) (res *AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/ownership", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponse struct {
	Errors   []AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseError   `json:"errors"`
	Messages []AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseMessage `json:"messages"`
	Result   AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseSuccess `json:"success"`
	JSON    accountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseJSON    `json:"-"`
}

// accountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseJSON
// contains the JSON metadata for the struct
// [AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponse]
type accountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseError struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseErrorJSON `json:"-"`
}

// accountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseError]
type accountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseMessage struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    accountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseMessageJSON `json:"-"`
}

// accountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseMessage]
type accountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseResult struct {
	Filename string                                                                                 `json:"filename"`
	Message  string                                                                                 `json:"message"`
	Valid    bool                                                                                   `json:"valid"`
	JSON     accountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseResultJSON `json:"-"`
}

// accountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseResultJSON
// contains the JSON metadata for the struct
// [AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseResult]
type accountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseResultJSON struct {
	Filename    apijson.Field
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseSuccess bool

const (
	AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseSuccessTrue AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseSuccess = true
)

type AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
