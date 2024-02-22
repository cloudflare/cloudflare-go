// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DcvDelegationUUIDService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDcvDelegationUUIDService] method
// instead.
type DcvDelegationUUIDService struct {
	Options []option.RequestOption
}

// NewDcvDelegationUUIDService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDcvDelegationUUIDService(opts ...option.RequestOption) (r *DcvDelegationUUIDService) {
	r = &DcvDelegationUUIDService{}
	r.Options = opts
	return
}

// Retrieve the account and zone specific unique identifier used as part of the
// CNAME target for DCV Delegation.
func (r *DcvDelegationUUIDService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *DcvDelegationUUIDGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DcvDelegationUUIDGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dcv_delegation/uuid", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DcvDelegationUUIDGetResponse struct {
	// The DCV Delegation unique identifier.
	UUID string                           `json:"uuid"`
	JSON dcvDelegationUUIDGetResponseJSON `json:"-"`
}

// dcvDelegationUUIDGetResponseJSON contains the JSON metadata for the struct
// [DcvDelegationUUIDGetResponse]
type dcvDelegationUUIDGetResponseJSON struct {
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DcvDelegationUUIDGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DcvDelegationUUIDGetResponseEnvelope struct {
	Errors   []DcvDelegationUUIDGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DcvDelegationUUIDGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DcvDelegationUUIDGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DcvDelegationUUIDGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dcvDelegationUUIDGetResponseEnvelopeJSON    `json:"-"`
}

// dcvDelegationUUIDGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DcvDelegationUUIDGetResponseEnvelope]
type dcvDelegationUUIDGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DcvDelegationUUIDGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DcvDelegationUUIDGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    dcvDelegationUUIDGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dcvDelegationUUIDGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DcvDelegationUUIDGetResponseEnvelopeErrors]
type dcvDelegationUUIDGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DcvDelegationUUIDGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DcvDelegationUUIDGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dcvDelegationUUIDGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dcvDelegationUUIDGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DcvDelegationUUIDGetResponseEnvelopeMessages]
type dcvDelegationUUIDGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DcvDelegationUUIDGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DcvDelegationUUIDGetResponseEnvelopeSuccess bool

const (
	DcvDelegationUUIDGetResponseEnvelopeSuccessTrue DcvDelegationUUIDGetResponseEnvelopeSuccess = true
)
