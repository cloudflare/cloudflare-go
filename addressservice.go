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

// AddressServiceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressServiceService] method
// instead.
type AddressServiceService struct {
	Options []option.RequestOption
}

// NewAddressServiceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressServiceService(opts ...option.RequestOption) (r *AddressServiceService) {
	r = &AddressServiceService{}
	r.Options = opts
	return
}

// Bring-Your-Own IP (BYOIP) prefixes onboarded to Cloudflare must be bound to a
// service running on the Cloudflare network to enable a Cloudflare product on the
// IP addresses. This endpoint can be used as a reference of available services on
// the Cloudflare network, and their service IDs.
func (r *AddressServiceService) List(ctx context.Context, query AddressServiceListParams, opts ...option.RequestOption) (res *[]AddressServiceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressServiceListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/services", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressServiceListResponse struct {
	// Identifier
	ID string `json:"id"`
	// Name of a service running on the Cloudflare network
	Name string                         `json:"name"`
	JSON addressServiceListResponseJSON `json:"-"`
}

// addressServiceListResponseJSON contains the JSON metadata for the struct
// [AddressServiceListResponse]
type addressServiceListResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressServiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressServiceListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressServiceListResponseEnvelope struct {
	Errors   []AddressServiceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressServiceListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressServiceListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success AddressServiceListResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressServiceListResponseEnvelopeJSON    `json:"-"`
}

// addressServiceListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressServiceListResponseEnvelope]
type addressServiceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressServiceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressServiceListResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    addressServiceListResponseEnvelopeErrorsJSON `json:"-"`
}

// addressServiceListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AddressServiceListResponseEnvelopeErrors]
type addressServiceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressServiceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressServiceListResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    addressServiceListResponseEnvelopeMessagesJSON `json:"-"`
}

// addressServiceListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AddressServiceListResponseEnvelopeMessages]
type addressServiceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressServiceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressServiceListResponseEnvelopeSuccess bool

const (
	AddressServiceListResponseEnvelopeSuccessTrue AddressServiceListResponseEnvelopeSuccess = true
)
