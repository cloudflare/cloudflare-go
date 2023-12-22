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
func (r *AccountLogpushOwnershipValidateService) PostAccountsAccountIdentifierLogpushOwnershipValidate(ctx context.Context, accountIdentifier string, body AccountLogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateParams, opts ...option.RequestOption) (res *ValidateOwnershipResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/ownership/validate", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ValidateOwnershipResponse struct {
	Errors   []ValidateOwnershipResponseError   `json:"errors"`
	Messages []ValidateOwnershipResponseMessage `json:"messages"`
	Result   ValidateOwnershipResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ValidateOwnershipResponseSuccess `json:"success"`
	JSON    validateOwnershipResponseJSON    `json:"-"`
}

// validateOwnershipResponseJSON contains the JSON metadata for the struct
// [ValidateOwnershipResponse]
type validateOwnershipResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateOwnershipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ValidateOwnershipResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    validateOwnershipResponseErrorJSON `json:"-"`
}

// validateOwnershipResponseErrorJSON contains the JSON metadata for the struct
// [ValidateOwnershipResponseError]
type validateOwnershipResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateOwnershipResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ValidateOwnershipResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    validateOwnershipResponseMessageJSON `json:"-"`
}

// validateOwnershipResponseMessageJSON contains the JSON metadata for the struct
// [ValidateOwnershipResponseMessage]
type validateOwnershipResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateOwnershipResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ValidateOwnershipResponseResult struct {
	Valid bool                                `json:"valid"`
	JSON  validateOwnershipResponseResultJSON `json:"-"`
}

// validateOwnershipResponseResultJSON contains the JSON metadata for the struct
// [ValidateOwnershipResponseResult]
type validateOwnershipResponseResultJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateOwnershipResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ValidateOwnershipResponseSuccess bool

const (
	ValidateOwnershipResponseSuccessTrue ValidateOwnershipResponseSuccess = true
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
