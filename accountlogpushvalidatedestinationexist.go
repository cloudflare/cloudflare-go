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

// AccountLogpushValidateDestinationExistService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLogpushValidateDestinationExistService] method instead.
type AccountLogpushValidateDestinationExistService struct {
	Options []option.RequestOption
}

// NewAccountLogpushValidateDestinationExistService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountLogpushValidateDestinationExistService(opts ...option.RequestOption) (r *AccountLogpushValidateDestinationExistService) {
	r = &AccountLogpushValidateDestinationExistService{}
	r.Options = opts
	return
}

// Checks if there is an existing job with a destination.
func (r *AccountLogpushValidateDestinationExistService) DeleteAccountsAccountIdentifierLogpushValidateDestinationExists(ctx context.Context, accountIdentifier string, body AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsParams, opts ...option.RequestOption) (res *AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/validate/destination/exists", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponse struct {
	Errors   []AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseError   `json:"errors"`
	Messages []AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseMessage `json:"messages"`
	Result   AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseSuccess `json:"success"`
	JSON    accountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseJSON    `json:"-"`
}

// accountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseJSON
// contains the JSON metadata for the struct
// [AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponse]
type accountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseError struct {
	Code    int64                                                                                                                  `json:"code,required"`
	Message string                                                                                                                 `json:"message,required"`
	JSON    accountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseErrorJSON `json:"-"`
}

// accountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseError]
type accountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseMessage struct {
	Code    int64                                                                                                                    `json:"code,required"`
	Message string                                                                                                                   `json:"message,required"`
	JSON    accountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseMessageJSON `json:"-"`
}

// accountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseMessage]
type accountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseResult struct {
	Exists bool                                                                                                                    `json:"exists"`
	JSON   accountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseResultJSON `json:"-"`
}

// accountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseResult]
type accountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseResultJSON struct {
	Exists      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseSuccess bool

const (
	AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseSuccessTrue AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseSuccess = true
)

type AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r AccountLogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
