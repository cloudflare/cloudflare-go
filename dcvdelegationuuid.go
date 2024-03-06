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

// DCVDelegationUUIDService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDCVDelegationUUIDService] method
// instead.
type DCVDelegationUUIDService struct {
	Options []option.RequestOption
}

// NewDCVDelegationUUIDService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDCVDelegationUUIDService(opts ...option.RequestOption) (r *DCVDelegationUUIDService) {
	r = &DCVDelegationUUIDService{}
	r.Options = opts
	return
}

// Retrieve the account and zone specific unique identifier used as part of the
// CNAME target for DCV Delegation.
func (r *DCVDelegationUUIDService) Get(ctx context.Context, query DCVDelegationUUIDGetParams, opts ...option.RequestOption) (res *DCVDelegationUUIDGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DCVDelegationUUIDGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dcv_delegation/uuid", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DCVDelegationUUIDGetResponse struct {
	// The DCV Delegation unique identifier.
	UUID string                           `json:"uuid"`
	JSON dcvDelegationUUIDGetResponseJSON `json:"-"`
}

// dcvDelegationUUIDGetResponseJSON contains the JSON metadata for the struct
// [DCVDelegationUUIDGetResponse]
type dcvDelegationUUIDGetResponseJSON struct {
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DCVDelegationUUIDGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DCVDelegationUUIDGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DCVDelegationUUIDGetResponseEnvelope struct {
	Errors   []DCVDelegationUUIDGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DCVDelegationUUIDGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DCVDelegationUUIDGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DCVDelegationUUIDGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dcvDelegationUUIDGetResponseEnvelopeJSON    `json:"-"`
}

// dcvDelegationUUIDGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DCVDelegationUUIDGetResponseEnvelope]
type dcvDelegationUUIDGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DCVDelegationUUIDGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DCVDelegationUUIDGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    dcvDelegationUUIDGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dcvDelegationUUIDGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DCVDelegationUUIDGetResponseEnvelopeErrors]
type dcvDelegationUUIDGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DCVDelegationUUIDGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DCVDelegationUUIDGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dcvDelegationUUIDGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dcvDelegationUUIDGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DCVDelegationUUIDGetResponseEnvelopeMessages]
type dcvDelegationUUIDGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DCVDelegationUUIDGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DCVDelegationUUIDGetResponseEnvelopeSuccess bool

const (
	DCVDelegationUUIDGetResponseEnvelopeSuccessTrue DCVDelegationUUIDGetResponseEnvelopeSuccess = true
)
