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

// AccountDlpPatternValidateService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountDlpPatternValidateService] method instead.
type AccountDlpPatternValidateService struct {
	Options []option.RequestOption
}

// NewAccountDlpPatternValidateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDlpPatternValidateService(opts ...option.RequestOption) (r *AccountDlpPatternValidateService) {
	r = &AccountDlpPatternValidateService{}
	r.Options = opts
	return
}

// Validates whether this pattern is a valid regular expression. Rejects it if the
// regular expression is too complex or can match an unbounded-length string. Your
// regex will be rejected if it uses the Kleene Star -- be sure to bound the
// maximum number of characters that can be matched.
func (r *AccountDlpPatternValidateService) DlpPatternValidationValidatePattern(ctx context.Context, accountIdentifier string, body AccountDlpPatternValidateDlpPatternValidationValidatePatternParams, opts ...option.RequestOption) (res *AccountDlpPatternValidateDlpPatternValidationValidatePatternResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/patterns/validate", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountDlpPatternValidateDlpPatternValidationValidatePatternResponse struct {
	Errors   []AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseError   `json:"errors"`
	Messages []AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseMessage `json:"messages"`
	Result   AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseSuccess `json:"success"`
	JSON    accountDlpPatternValidateDlpPatternValidationValidatePatternResponseJSON    `json:"-"`
}

// accountDlpPatternValidateDlpPatternValidationValidatePatternResponseJSON
// contains the JSON metadata for the struct
// [AccountDlpPatternValidateDlpPatternValidationValidatePatternResponse]
type accountDlpPatternValidateDlpPatternValidationValidatePatternResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPatternValidateDlpPatternValidationValidatePatternResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseError struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    accountDlpPatternValidateDlpPatternValidationValidatePatternResponseErrorJSON `json:"-"`
}

// accountDlpPatternValidateDlpPatternValidationValidatePatternResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseError]
type accountDlpPatternValidateDlpPatternValidationValidatePatternResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseMessage struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accountDlpPatternValidateDlpPatternValidationValidatePatternResponseMessageJSON `json:"-"`
}

// accountDlpPatternValidateDlpPatternValidationValidatePatternResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseMessage]
type accountDlpPatternValidateDlpPatternValidationValidatePatternResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseResult struct {
	Valid bool                                                                           `json:"valid"`
	JSON  accountDlpPatternValidateDlpPatternValidationValidatePatternResponseResultJSON `json:"-"`
}

// accountDlpPatternValidateDlpPatternValidationValidatePatternResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseResult]
type accountDlpPatternValidateDlpPatternValidationValidatePatternResponseResultJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseSuccess bool

const (
	AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseSuccessTrue AccountDlpPatternValidateDlpPatternValidationValidatePatternResponseSuccess = true
)

type AccountDlpPatternValidateDlpPatternValidationValidatePatternParams struct {
	// The regex pattern.
	Regex param.Field[string] `json:"regex,required"`
}

func (r AccountDlpPatternValidateDlpPatternValidationValidatePatternParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
