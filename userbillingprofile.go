// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// UserBillingProfileService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserBillingProfileService] method
// instead.
type UserBillingProfileService struct {
	Options []option.RequestOption
}

// NewUserBillingProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserBillingProfileService(opts ...option.RequestOption) (r *UserBillingProfileService) {
	r = &UserBillingProfileService{}
	r.Options = opts
	return
}

// Accesses your billing profile object.
func (r *UserBillingProfileService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserBillingProfileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserBillingProfileGetResponseEnvelope
	path := "user/billing/profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [UserBillingProfileGetResponseUnknown] or
// [shared.UnionString].
type UserBillingProfileGetResponse interface {
	ImplementsUserBillingProfileGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserBillingProfileGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserBillingProfileGetResponseEnvelope struct {
	Errors   []UserBillingProfileGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserBillingProfileGetResponseEnvelopeMessages `json:"messages,required"`
	Result   UserBillingProfileGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserBillingProfileGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    userBillingProfileGetResponseEnvelopeJSON    `json:"-"`
}

// userBillingProfileGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserBillingProfileGetResponseEnvelope]
type userBillingProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBillingProfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserBillingProfileGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    userBillingProfileGetResponseEnvelopeErrorsJSON `json:"-"`
}

// userBillingProfileGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UserBillingProfileGetResponseEnvelopeErrors]
type userBillingProfileGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingProfileGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBillingProfileGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserBillingProfileGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    userBillingProfileGetResponseEnvelopeMessagesJSON `json:"-"`
}

// userBillingProfileGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [UserBillingProfileGetResponseEnvelopeMessages]
type userBillingProfileGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingProfileGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBillingProfileGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserBillingProfileGetResponseEnvelopeSuccess bool

const (
	UserBillingProfileGetResponseEnvelopeSuccessTrue UserBillingProfileGetResponseEnvelopeSuccess = true
)
