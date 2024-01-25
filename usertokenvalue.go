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
func (r *UserTokenValueService) UserAPITokensRollToken(ctx context.Context, identifier interface{}, body UserTokenValueUserAPITokensRollTokenParams, opts ...option.RequestOption) (res *UserTokenValueUserAPITokensRollTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/tokens/%v/value", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type UserTokenValueUserAPITokensRollTokenResponse struct {
	Errors   []UserTokenValueUserAPITokensRollTokenResponseError   `json:"errors"`
	Messages []UserTokenValueUserAPITokensRollTokenResponseMessage `json:"messages"`
	// The token value.
	Result string `json:"result"`
	// Whether the API call was successful
	Success UserTokenValueUserAPITokensRollTokenResponseSuccess `json:"success"`
	JSON    userTokenValueUserAPITokensRollTokenResponseJSON    `json:"-"`
}

// userTokenValueUserAPITokensRollTokenResponseJSON contains the JSON metadata for
// the struct [UserTokenValueUserAPITokensRollTokenResponse]
type userTokenValueUserAPITokensRollTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenValueUserAPITokensRollTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenValueUserAPITokensRollTokenResponseError struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    userTokenValueUserAPITokensRollTokenResponseErrorJSON `json:"-"`
}

// userTokenValueUserAPITokensRollTokenResponseErrorJSON contains the JSON metadata
// for the struct [UserTokenValueUserAPITokensRollTokenResponseError]
type userTokenValueUserAPITokensRollTokenResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenValueUserAPITokensRollTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenValueUserAPITokensRollTokenResponseMessage struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    userTokenValueUserAPITokensRollTokenResponseMessageJSON `json:"-"`
}

// userTokenValueUserAPITokensRollTokenResponseMessageJSON contains the JSON
// metadata for the struct [UserTokenValueUserAPITokensRollTokenResponseMessage]
type userTokenValueUserAPITokensRollTokenResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenValueUserAPITokensRollTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenValueUserAPITokensRollTokenResponseSuccess bool

const (
	UserTokenValueUserAPITokensRollTokenResponseSuccessTrue UserTokenValueUserAPITokensRollTokenResponseSuccess = true
)

type UserTokenValueUserAPITokensRollTokenParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r UserTokenValueUserAPITokensRollTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
