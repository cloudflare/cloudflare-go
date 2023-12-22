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
func (r *UserTokenValueService) UserAPITokensRollToken(ctx context.Context, identifier interface{}, body UserTokenValueUserAPITokensRollTokenParams, opts ...option.RequestOption) (res *ResponseSingleValue, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/tokens/%v/value", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ResponseSingleValue struct {
	Errors   []ResponseSingleValueError   `json:"errors"`
	Messages []ResponseSingleValueMessage `json:"messages"`
	// The token value.
	Result string `json:"result"`
	// Whether the API call was successful
	Success ResponseSingleValueSuccess `json:"success"`
	JSON    responseSingleValueJSON    `json:"-"`
}

// responseSingleValueJSON contains the JSON metadata for the struct
// [ResponseSingleValue]
type responseSingleValueJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseSingleValueError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    responseSingleValueErrorJSON `json:"-"`
}

// responseSingleValueErrorJSON contains the JSON metadata for the struct
// [ResponseSingleValueError]
type responseSingleValueErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleValueError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseSingleValueMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    responseSingleValueMessageJSON `json:"-"`
}

// responseSingleValueMessageJSON contains the JSON metadata for the struct
// [ResponseSingleValueMessage]
type responseSingleValueMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleValueMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseSingleValueSuccess bool

const (
	ResponseSingleValueSuccessTrue ResponseSingleValueSuccess = true
)

type UserTokenValueUserAPITokensRollTokenParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r UserTokenValueUserAPITokensRollTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
