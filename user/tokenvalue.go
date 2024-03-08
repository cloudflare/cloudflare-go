// File generated from our OpenAPI spec by Stainless.

package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// TokenValueService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewTokenValueService] method instead.
type TokenValueService struct {
	Options []option.RequestOption
}

// NewTokenValueService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTokenValueService(opts ...option.RequestOption) (r *TokenValueService) {
	r = &TokenValueService{}
	r.Options = opts
	return
}

// Roll the token secret.
func (r *TokenValueService) Update(ctx context.Context, tokenID interface{}, body TokenValueUpdateParams, opts ...option.RequestOption) (res *IamValue, err error) {
	opts = append(r.Options[:], opts...)
	var env TokenValueUpdateResponseEnvelope
	path := fmt.Sprintf("user/tokens/%v/value", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IamValue = string

type TokenValueUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r TokenValueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type TokenValueUpdateResponseEnvelope struct {
	Errors   []TokenValueUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenValueUpdateResponseEnvelopeMessages `json:"messages,required"`
	// The token value.
	Result IamValue `json:"result,required"`
	// Whether the API call was successful
	Success TokenValueUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    tokenValueUpdateResponseEnvelopeJSON    `json:"-"`
}

// tokenValueUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenValueUpdateResponseEnvelope]
type tokenValueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenValueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenValueUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenValueUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    tokenValueUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// tokenValueUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TokenValueUpdateResponseEnvelopeErrors]
type tokenValueUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenValueUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenValueUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenValueUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    tokenValueUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// tokenValueUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TokenValueUpdateResponseEnvelopeMessages]
type tokenValueUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenValueUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenValueUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TokenValueUpdateResponseEnvelopeSuccess bool

const (
	TokenValueUpdateResponseEnvelopeSuccessTrue TokenValueUpdateResponseEnvelopeSuccess = true
)
