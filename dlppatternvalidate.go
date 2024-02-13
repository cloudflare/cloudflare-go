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

// DLPPatternValidateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDLPPatternValidateService] method
// instead.
type DLPPatternValidateService struct {
	Options []option.RequestOption
}

// NewDLPPatternValidateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPPatternValidateService(opts ...option.RequestOption) (r *DLPPatternValidateService) {
	r = &DLPPatternValidateService{}
	r.Options = opts
	return
}

// Validates whether this pattern is a valid regular expression. Rejects it if the
// regular expression is too complex or can match an unbounded-length string. Your
// regex will be rejected if it uses the Kleene Star -- be sure to bound the
// maximum number of characters that can be matched.
func (r *DLPPatternValidateService) DLPPatternValidationValidatePattern(ctx context.Context, accountID string, body DLPPatternValidateDLPPatternValidationValidatePatternParams, opts ...option.RequestOption) (res *DLPPatternValidateDLPPatternValidationValidatePatternResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/patterns/validate", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPPatternValidateDLPPatternValidationValidatePatternResponse struct {
	Valid bool                                                              `json:"valid"`
	JSON  dlpPatternValidateDLPPatternValidationValidatePatternResponseJSON `json:"-"`
}

// dlpPatternValidateDLPPatternValidationValidatePatternResponseJSON contains the
// JSON metadata for the struct
// [DLPPatternValidateDLPPatternValidationValidatePatternResponse]
type dlpPatternValidateDLPPatternValidationValidatePatternResponseJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPatternValidateDLPPatternValidationValidatePatternResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPatternValidateDLPPatternValidationValidatePatternParams struct {
	// The regex pattern.
	Regex param.Field[string] `json:"regex,required"`
}

func (r DLPPatternValidateDLPPatternValidationValidatePatternParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelope struct {
	Errors   []DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeMessages `json:"messages,required"`
	Result   DLPPatternValidateDLPPatternValidationValidatePatternResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeJSON    `json:"-"`
}

// dlpPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelope]
type dlpPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeErrors struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    dlpPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeErrors]
type dlpPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeMessages struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    dlpPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeMessages]
type dlpPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeSuccess bool

const (
	DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeSuccessTrue DLPPatternValidateDLPPatternValidationValidatePatternResponseEnvelopeSuccess = true
)
