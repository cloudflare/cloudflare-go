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
func (r *AccountDlpPatternValidateService) DlpPatternValidationValidatePattern(ctx context.Context, accountIdentifier string, body AccountDlpPatternValidateDlpPatternValidationValidatePatternParams, opts ...option.RequestOption) (res *ValidateResponseLbtsZt0d, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/patterns/validate", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ValidateResponseLbtsZt0d struct {
	Errors   []ValidateResponseLbtsZt0dError   `json:"errors"`
	Messages []ValidateResponseLbtsZt0dMessage `json:"messages"`
	Result   ValidateResponseLbtsZt0dResult    `json:"result"`
	// Whether the API call was successful
	Success ValidateResponseLbtsZt0dSuccess `json:"success"`
	JSON    validateResponseLbtsZt0dJSON    `json:"-"`
}

// validateResponseLbtsZt0dJSON contains the JSON metadata for the struct
// [ValidateResponseLbtsZt0d]
type validateResponseLbtsZt0dJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateResponseLbtsZt0d) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ValidateResponseLbtsZt0dError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    validateResponseLbtsZt0dErrorJSON `json:"-"`
}

// validateResponseLbtsZt0dErrorJSON contains the JSON metadata for the struct
// [ValidateResponseLbtsZt0dError]
type validateResponseLbtsZt0dErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateResponseLbtsZt0dError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ValidateResponseLbtsZt0dMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    validateResponseLbtsZt0dMessageJSON `json:"-"`
}

// validateResponseLbtsZt0dMessageJSON contains the JSON metadata for the struct
// [ValidateResponseLbtsZt0dMessage]
type validateResponseLbtsZt0dMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateResponseLbtsZt0dMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ValidateResponseLbtsZt0dResult struct {
	Valid bool                               `json:"valid"`
	JSON  validateResponseLbtsZt0dResultJSON `json:"-"`
}

// validateResponseLbtsZt0dResultJSON contains the JSON metadata for the struct
// [ValidateResponseLbtsZt0dResult]
type validateResponseLbtsZt0dResultJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateResponseLbtsZt0dResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ValidateResponseLbtsZt0dSuccess bool

const (
	ValidateResponseLbtsZt0dSuccessTrue ValidateResponseLbtsZt0dSuccess = true
)

type AccountDlpPatternValidateDlpPatternValidationValidatePatternParams struct {
	// The regex pattern.
	Regex param.Field[string] `json:"regex,required"`
}

func (r AccountDlpPatternValidateDlpPatternValidationValidatePatternParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
