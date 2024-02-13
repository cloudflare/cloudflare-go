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

// DcvDelegationUuidService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDcvDelegationUuidService] method
// instead.
type DcvDelegationUuidService struct {
	Options []option.RequestOption
}

// NewDcvDelegationUuidService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDcvDelegationUuidService(opts ...option.RequestOption) (r *DcvDelegationUuidService) {
	r = &DcvDelegationUuidService{}
	r.Options = opts
	return
}

// Retrieve the account and zone specific unique identifier used as part of the
// CNAME target for DCV Delegation.
func (r *DcvDelegationUuidService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *DcvDelegationUuidGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DcvDelegationUuidGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dcv_delegation/uuid", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DcvDelegationUuidGetResponse struct {
	// The DCV Delegation unique identifier.
	Uuid string                           `json:"uuid"`
	JSON dcvDelegationUuidGetResponseJSON `json:"-"`
}

// dcvDelegationUuidGetResponseJSON contains the JSON metadata for the struct
// [DcvDelegationUuidGetResponse]
type dcvDelegationUuidGetResponseJSON struct {
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DcvDelegationUuidGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DcvDelegationUuidGetResponseEnvelope struct {
	Errors   []DcvDelegationUuidGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DcvDelegationUuidGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DcvDelegationUuidGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DcvDelegationUuidGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dcvDelegationUuidGetResponseEnvelopeJSON    `json:"-"`
}

// dcvDelegationUuidGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DcvDelegationUuidGetResponseEnvelope]
type dcvDelegationUuidGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DcvDelegationUuidGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DcvDelegationUuidGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    dcvDelegationUuidGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dcvDelegationUuidGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DcvDelegationUuidGetResponseEnvelopeErrors]
type dcvDelegationUuidGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DcvDelegationUuidGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DcvDelegationUuidGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dcvDelegationUuidGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dcvDelegationUuidGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DcvDelegationUuidGetResponseEnvelopeMessages]
type dcvDelegationUuidGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DcvDelegationUuidGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DcvDelegationUuidGetResponseEnvelopeSuccess bool

const (
	DcvDelegationUuidGetResponseEnvelopeSuccessTrue DcvDelegationUuidGetResponseEnvelopeSuccess = true
)
