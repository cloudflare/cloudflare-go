// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserTokenVerifyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserTokenVerifyService] method
// instead.
type UserTokenVerifyService struct {
	Options []option.RequestOption
}

// NewUserTokenVerifyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserTokenVerifyService(opts ...option.RequestOption) (r *UserTokenVerifyService) {
	r = &UserTokenVerifyService{}
	r.Options = opts
	return
}

// Test whether a token works.
func (r *UserTokenVerifyService) UserAPITokensVerifyToken(ctx context.Context, opts ...option.RequestOption) (res *UserTokenVerifyUserAPITokensVerifyTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelope
	path := "user/tokens/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserTokenVerifyUserAPITokensVerifyTokenResponse struct {
	// Token identifier tag.
	ID string `json:"id,required"`
	// Status of the token.
	Status UserTokenVerifyUserAPITokensVerifyTokenResponseStatus `json:"status,required"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time                                           `json:"not_before" format:"date-time"`
	JSON      userTokenVerifyUserAPITokensVerifyTokenResponseJSON `json:"-"`
}

// userTokenVerifyUserAPITokensVerifyTokenResponseJSON contains the JSON metadata
// for the struct [UserTokenVerifyUserAPITokensVerifyTokenResponse]
type userTokenVerifyUserAPITokensVerifyTokenResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	ExpiresOn   apijson.Field
	NotBefore   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyUserAPITokensVerifyTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the token.
type UserTokenVerifyUserAPITokensVerifyTokenResponseStatus string

const (
	UserTokenVerifyUserAPITokensVerifyTokenResponseStatusActive   UserTokenVerifyUserAPITokensVerifyTokenResponseStatus = "active"
	UserTokenVerifyUserAPITokensVerifyTokenResponseStatusDisabled UserTokenVerifyUserAPITokensVerifyTokenResponseStatus = "disabled"
	UserTokenVerifyUserAPITokensVerifyTokenResponseStatusExpired  UserTokenVerifyUserAPITokensVerifyTokenResponseStatus = "expired"
)

type UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelope struct {
	Errors   []UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeMessages `json:"messages,required"`
	Result   UserTokenVerifyUserAPITokensVerifyTokenResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeSuccess `json:"success,required"`
	JSON    userTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeJSON    `json:"-"`
}

// userTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelope]
type userTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    userTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeErrorsJSON `json:"-"`
}

// userTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeErrors]
type userTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    userTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeMessagesJSON `json:"-"`
}

// userTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeMessages]
type userTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeSuccess bool

const (
	UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeSuccessTrue UserTokenVerifyUserAPITokensVerifyTokenResponseEnvelopeSuccess = true
)
