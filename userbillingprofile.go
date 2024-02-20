// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
func (r *UserBillingProfileService) List(ctx context.Context, opts ...option.RequestOption) (res *UserBillingProfileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserBillingProfileListResponseEnvelope
	path := "user/billing/profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [UserBillingProfileListResponseUnknown] or
// [shared.UnionString].
type UserBillingProfileListResponse interface {
	ImplementsUserBillingProfileListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserBillingProfileListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserBillingProfileListResponseEnvelope struct {
	Errors   []UserBillingProfileListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserBillingProfileListResponseEnvelopeMessages `json:"messages,required"`
	Result   UserBillingProfileListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserBillingProfileListResponseEnvelopeSuccess `json:"success,required"`
	JSON    userBillingProfileListResponseEnvelopeJSON    `json:"-"`
}

// userBillingProfileListResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserBillingProfileListResponseEnvelope]
type userBillingProfileListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingProfileListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingProfileListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    userBillingProfileListResponseEnvelopeErrorsJSON `json:"-"`
}

// userBillingProfileListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UserBillingProfileListResponseEnvelopeErrors]
type userBillingProfileListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingProfileListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingProfileListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    userBillingProfileListResponseEnvelopeMessagesJSON `json:"-"`
}

// userBillingProfileListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [UserBillingProfileListResponseEnvelopeMessages]
type userBillingProfileListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingProfileListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserBillingProfileListResponseEnvelopeSuccess bool

const (
	UserBillingProfileListResponseEnvelopeSuccessTrue UserBillingProfileListResponseEnvelopeSuccess = true
)
