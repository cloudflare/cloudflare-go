// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// BillingProfileService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewBillingProfileService] method
// instead.
type BillingProfileService struct {
	Options []option.RequestOption
}

// NewBillingProfileService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBillingProfileService(opts ...option.RequestOption) (r *BillingProfileService) {
	r = &BillingProfileService{}
	r.Options = opts
	return
}

// Gets the current billing profile for the account.
func (r *BillingProfileService) AccountBillingProfileBillingProfileDetails(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *BillingProfileAccountBillingProfileBillingProfileDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelope
	path := fmt.Sprintf("accounts/%v/billing/profile", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [BillingProfileAccountBillingProfileBillingProfileDetailsResponseUnknown] or
// [shared.UnionString].
type BillingProfileAccountBillingProfileBillingProfileDetailsResponse interface {
	ImplementsBillingProfileAccountBillingProfileBillingProfileDetailsResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingProfileAccountBillingProfileBillingProfileDetailsResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelope struct {
	Errors   []BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeMessages `json:"messages,required"`
	Result   BillingProfileAccountBillingProfileBillingProfileDetailsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeSuccess `json:"success,required"`
	JSON    billingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeJSON    `json:"-"`
}

// billingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelope]
type billingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeErrors struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    billingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeErrorsJSON `json:"-"`
}

// billingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeErrors]
type billingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeMessages struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    billingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeMessagesJSON `json:"-"`
}

// billingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeMessages]
type billingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeSuccess bool

const (
	BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeSuccessTrue BillingProfileAccountBillingProfileBillingProfileDetailsResponseEnvelopeSuccess = true
)
