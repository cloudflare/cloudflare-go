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
func (r *AccountLogpushValidateOriginService) PostAccountsAccountIdentifierLogpushValidateOrigin(ctx context.Context, accountIdentifier string, body AccountLogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginParams, opts ...option.RequestOption) (res *ValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/validate/origin", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ValidateResponse struct {
	Errors   []ValidateResponseError   `json:"errors"`
	Messages []ValidateResponseMessage `json:"messages"`
	Result   ValidateResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ValidateResponseSuccess `json:"success"`
	JSON    validateResponseJSON    `json:"-"`
}

// validateResponseJSON contains the JSON metadata for the struct
// [ValidateResponse]
type validateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ValidateResponseError struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    validateResponseErrorJSON `json:"-"`
}

// validateResponseErrorJSON contains the JSON metadata for the struct
// [ValidateResponseError]
type validateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ValidateResponseMessage struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    validateResponseMessageJSON `json:"-"`
}

// validateResponseMessageJSON contains the JSON metadata for the struct
// [ValidateResponseMessage]
type validateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ValidateResponseResult struct {
	Message string                     `json:"message"`
	Valid   bool                       `json:"valid"`
	JSON    validateResponseResultJSON `json:"-"`
}

// validateResponseResultJSON contains the JSON metadata for the struct
// [ValidateResponseResult]
type validateResponseResultJSON struct {
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ValidateResponseSuccess bool

const (
	ValidateResponseSuccessTrue ValidateResponseSuccess = true
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
