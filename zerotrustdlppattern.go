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

// ZeroTrustDLPPatternService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDLPPatternService]
// method instead.
type ZeroTrustDLPPatternService struct {
	Options []option.RequestOption
}

// NewZeroTrustDLPPatternService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDLPPatternService(opts ...option.RequestOption) (r *ZeroTrustDLPPatternService) {
	r = &ZeroTrustDLPPatternService{}
	r.Options = opts
	return
}

// Validates whether this pattern is a valid regular expression. Rejects it if the
// regular expression is too complex or can match an unbounded-length string. Your
// regex will be rejected if it uses the Kleene Star -- be sure to bound the
// maximum number of characters that can be matched.
func (r *ZeroTrustDLPPatternService) Validate(ctx context.Context, params ZeroTrustDLPPatternValidateParams, opts ...option.RequestOption) (res *ZeroTrustDLPPatternValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPPatternValidateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/patterns/validate", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDLPPatternValidateResponse struct {
	Valid bool                                    `json:"valid"`
	JSON  zeroTrustDLPPatternValidateResponseJSON `json:"-"`
}

// zeroTrustDLPPatternValidateResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDLPPatternValidateResponse]
type zeroTrustDLPPatternValidateResponseJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPPatternValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPPatternValidateResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPPatternValidateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The regex pattern.
	Regex param.Field[string] `json:"regex,required"`
}

func (r ZeroTrustDLPPatternValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDLPPatternValidateResponseEnvelope struct {
	Errors   []ZeroTrustDLPPatternValidateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDLPPatternValidateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDLPPatternValidateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDLPPatternValidateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDLPPatternValidateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDLPPatternValidateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDLPPatternValidateResponseEnvelope]
type zeroTrustDLPPatternValidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPPatternValidateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPPatternValidateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPPatternValidateResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustDLPPatternValidateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPPatternValidateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDLPPatternValidateResponseEnvelopeErrors]
type zeroTrustDLPPatternValidateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPPatternValidateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPPatternValidateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPPatternValidateResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustDLPPatternValidateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPPatternValidateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDLPPatternValidateResponseEnvelopeMessages]
type zeroTrustDLPPatternValidateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPPatternValidateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPPatternValidateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustDLPPatternValidateResponseEnvelopeSuccess bool

const (
	ZeroTrustDLPPatternValidateResponseEnvelopeSuccessTrue ZeroTrustDLPPatternValidateResponseEnvelopeSuccess = true
)
