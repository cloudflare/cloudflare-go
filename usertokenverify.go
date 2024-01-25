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
	path := "user/tokens/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserTokenVerifyUserAPITokensVerifyTokenResponse struct {
	Errors   []UserTokenVerifyUserAPITokensVerifyTokenResponseError   `json:"errors"`
	Messages []UserTokenVerifyUserAPITokensVerifyTokenResponseMessage `json:"messages"`
	Result   UserTokenVerifyUserAPITokensVerifyTokenResponseResult    `json:"result"`
	// Whether the API call was successful
	Success UserTokenVerifyUserAPITokensVerifyTokenResponseSuccess `json:"success"`
	JSON    userTokenVerifyUserAPITokensVerifyTokenResponseJSON    `json:"-"`
}

// userTokenVerifyUserAPITokensVerifyTokenResponseJSON contains the JSON metadata
// for the struct [UserTokenVerifyUserAPITokensVerifyTokenResponse]
type userTokenVerifyUserAPITokensVerifyTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyUserAPITokensVerifyTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenVerifyUserAPITokensVerifyTokenResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    userTokenVerifyUserAPITokensVerifyTokenResponseErrorJSON `json:"-"`
}

// userTokenVerifyUserAPITokensVerifyTokenResponseErrorJSON contains the JSON
// metadata for the struct [UserTokenVerifyUserAPITokensVerifyTokenResponseError]
type userTokenVerifyUserAPITokensVerifyTokenResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyUserAPITokensVerifyTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenVerifyUserAPITokensVerifyTokenResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    userTokenVerifyUserAPITokensVerifyTokenResponseMessageJSON `json:"-"`
}

// userTokenVerifyUserAPITokensVerifyTokenResponseMessageJSON contains the JSON
// metadata for the struct [UserTokenVerifyUserAPITokensVerifyTokenResponseMessage]
type userTokenVerifyUserAPITokensVerifyTokenResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyUserAPITokensVerifyTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenVerifyUserAPITokensVerifyTokenResponseResult struct {
	// Token identifier tag.
	ID string `json:"id,required"`
	// Status of the token.
	Status UserTokenVerifyUserAPITokensVerifyTokenResponseResultStatus `json:"status,required"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time                                                 `json:"not_before" format:"date-time"`
	JSON      userTokenVerifyUserAPITokensVerifyTokenResponseResultJSON `json:"-"`
}

// userTokenVerifyUserAPITokensVerifyTokenResponseResultJSON contains the JSON
// metadata for the struct [UserTokenVerifyUserAPITokensVerifyTokenResponseResult]
type userTokenVerifyUserAPITokensVerifyTokenResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	ExpiresOn   apijson.Field
	NotBefore   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyUserAPITokensVerifyTokenResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the token.
type UserTokenVerifyUserAPITokensVerifyTokenResponseResultStatus string

const (
	UserTokenVerifyUserAPITokensVerifyTokenResponseResultStatusActive   UserTokenVerifyUserAPITokensVerifyTokenResponseResultStatus = "active"
	UserTokenVerifyUserAPITokensVerifyTokenResponseResultStatusDisabled UserTokenVerifyUserAPITokensVerifyTokenResponseResultStatus = "disabled"
	UserTokenVerifyUserAPITokensVerifyTokenResponseResultStatusExpired  UserTokenVerifyUserAPITokensVerifyTokenResponseResultStatus = "expired"
)

// Whether the API call was successful
type UserTokenVerifyUserAPITokensVerifyTokenResponseSuccess bool

const (
	UserTokenVerifyUserAPITokensVerifyTokenResponseSuccessTrue UserTokenVerifyUserAPITokensVerifyTokenResponseSuccess = true
)
