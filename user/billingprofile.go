// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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

// Accesses your billing profile object.
func (r *BillingProfileService) Get(ctx context.Context, opts ...option.RequestOption) (res *BillingProfileGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env BillingProfileGetResponseEnvelope
	path := "user/billing/profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [user.BillingProfileGetResponseUnknown] or
// [shared.UnionString].
type BillingProfileGetResponseUnion interface {
	ImplementsUserBillingProfileGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingProfileGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type BillingProfileGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   BillingProfileGetResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success BillingProfileGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    billingProfileGetResponseEnvelopeJSON    `json:"-"`
}

// billingProfileGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BillingProfileGetResponseEnvelope]
type billingProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingProfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BillingProfileGetResponseEnvelopeSuccess bool

const (
	BillingProfileGetResponseEnvelopeSuccessTrue BillingProfileGetResponseEnvelopeSuccess = true
)

func (r BillingProfileGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingProfileGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
