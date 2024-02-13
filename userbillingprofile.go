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
func (r *UserBillingProfileService) UserBillingProfileBillingProfileDetails(ctx context.Context, opts ...option.RequestOption) (res *UserBillingProfileUserBillingProfileBillingProfileDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelope
	path := "user/billing/profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [UserBillingProfileUserBillingProfileBillingProfileDetailsResponseUnknown] or
// [shared.UnionString].
type UserBillingProfileUserBillingProfileBillingProfileDetailsResponse interface {
	ImplementsUserBillingProfileUserBillingProfileBillingProfileDetailsResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserBillingProfileUserBillingProfileBillingProfileDetailsResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelope struct {
	Errors   []UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeMessages `json:"messages,required"`
	Result   UserBillingProfileUserBillingProfileBillingProfileDetailsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeSuccess `json:"success,required"`
	JSON    userBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeJSON    `json:"-"`
}

// userBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelope]
type userBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeErrors struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    userBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeErrorsJSON `json:"-"`
}

// userBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeErrors]
type userBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeMessages struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    userBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeMessagesJSON `json:"-"`
}

// userBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeMessages]
type userBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeSuccess bool

const (
	UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeSuccessTrue UserBillingProfileUserBillingProfileBillingProfileDetailsResponseEnvelopeSuccess = true
)
