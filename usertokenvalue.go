// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *UserTokenValueService) Update(ctx context.Context, tokenID interface{}, body UserTokenValueUpdateParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env UserTokenValueUpdateResponseEnvelope
	path := fmt.Sprintf("user/tokens/%v/value", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserTokenValueUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r UserTokenValueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type UserTokenValueUpdateResponseEnvelope struct {
	Errors   []UserTokenValueUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserTokenValueUpdateResponseEnvelopeMessages `json:"messages,required"`
	// The token value.
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success UserTokenValueUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    userTokenValueUpdateResponseEnvelopeJSON    `json:"-"`
}

// userTokenValueUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserTokenValueUpdateResponseEnvelope]
type userTokenValueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenValueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userTokenValueUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserTokenValueUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    userTokenValueUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// userTokenValueUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UserTokenValueUpdateResponseEnvelopeErrors]
type userTokenValueUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenValueUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userTokenValueUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserTokenValueUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    userTokenValueUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// userTokenValueUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [UserTokenValueUpdateResponseEnvelopeMessages]
type userTokenValueUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenValueUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userTokenValueUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserTokenValueUpdateResponseEnvelopeSuccess bool

const (
	UserTokenValueUpdateResponseEnvelopeSuccessTrue UserTokenValueUpdateResponseEnvelopeSuccess = true
)
