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

// UserTokenValueService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserTokenValueService] method
// instead.
type UserTokenValueService struct {
	Options []option.RequestOption
}

// NewUserTokenValueService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserTokenValueService(opts ...option.RequestOption) (r *UserTokenValueService) {
	r = &UserTokenValueService{}
	r.Options = opts
	return
}

// Roll the token secret.
func (r *UserTokenValueService) UserAPITokensRollToken(ctx context.Context, tokenID interface{}, body UserTokenValueUserAPITokensRollTokenParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env UserTokenValueUserAPITokensRollTokenResponseEnvelope
	path := fmt.Sprintf("user/tokens/%v/value", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserTokenValueUserAPITokensRollTokenParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r UserTokenValueUserAPITokensRollTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type UserTokenValueUserAPITokensRollTokenResponseEnvelope struct {
	Errors   []UserTokenValueUserAPITokensRollTokenResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserTokenValueUserAPITokensRollTokenResponseEnvelopeMessages `json:"messages,required"`
	// The token value.
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success UserTokenValueUserAPITokensRollTokenResponseEnvelopeSuccess `json:"success,required"`
	JSON    userTokenValueUserAPITokensRollTokenResponseEnvelopeJSON    `json:"-"`
}

// userTokenValueUserAPITokensRollTokenResponseEnvelopeJSON contains the JSON
// metadata for the struct [UserTokenValueUserAPITokensRollTokenResponseEnvelope]
type userTokenValueUserAPITokensRollTokenResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenValueUserAPITokensRollTokenResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenValueUserAPITokensRollTokenResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    userTokenValueUserAPITokensRollTokenResponseEnvelopeErrorsJSON `json:"-"`
}

// userTokenValueUserAPITokensRollTokenResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [UserTokenValueUserAPITokensRollTokenResponseEnvelopeErrors]
type userTokenValueUserAPITokensRollTokenResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenValueUserAPITokensRollTokenResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenValueUserAPITokensRollTokenResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    userTokenValueUserAPITokensRollTokenResponseEnvelopeMessagesJSON `json:"-"`
}

// userTokenValueUserAPITokensRollTokenResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [UserTokenValueUserAPITokensRollTokenResponseEnvelopeMessages]
type userTokenValueUserAPITokensRollTokenResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenValueUserAPITokensRollTokenResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenValueUserAPITokensRollTokenResponseEnvelopeSuccess bool

const (
	UserTokenValueUserAPITokensRollTokenResponseEnvelopeSuccessTrue UserTokenValueUserAPITokensRollTokenResponseEnvelopeSuccess = true
)
