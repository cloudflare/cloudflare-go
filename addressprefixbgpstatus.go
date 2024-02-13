// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AddressPrefixBGPStatusService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressPrefixBGPStatusService]
// method instead.
type AddressPrefixBGPStatusService struct {
	Options []option.RequestOption
}

// NewAddressPrefixBGPStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressPrefixBGPStatusService(opts ...option.RequestOption) (r *AddressPrefixBGPStatusService) {
	r = &AddressPrefixBGPStatusService{}
	r.Options = opts
	return
}

// List the current advertisement state for a prefix.
func (r *AddressPrefixBGPStatusService) IPAddressManagementDynamicAdvertisementGetAdvertisementStatus(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Advertise or withdraw BGP route for a prefix.
func (r *AddressPrefixBGPStatusService) IPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatus(ctx context.Context, accountIdentifier string, prefixIdentifier string, body AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusParams, opts ...option.RequestOption) (res *AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponse struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                                                                                       `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseJSON `json:"-"`
}

// addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseJSON
// contains the JSON metadata for the struct
// [AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponse]
type addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponse struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                                                                                                       `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseJSON `json:"-"`
}

// addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseJSON
// contains the JSON metadata for the struct
// [AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponse]
type addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelope struct {
	Errors   []AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelope]
type addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeErrors struct {
	Code    int64                                                                                                         `json:"code,required"`
	Message string                                                                                                        `json:"message,required"`
	JSON    addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeErrors]
type addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeMessages struct {
	Code    int64                                                                                                           `json:"code,required"`
	Message string                                                                                                          `json:"message,required"`
	JSON    addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeMessages]
type addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeSuccess bool

const (
	AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeSuccessTrue AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseEnvelopeSuccess = true
)

type AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusParams struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised param.Field[bool] `json:"advertised,required"`
}

func (r AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelope struct {
	Errors   []AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelope]
type addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeErrors struct {
	Code    int64                                                                                                                         `json:"code,required"`
	Message string                                                                                                                        `json:"message,required"`
	JSON    addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeErrors]
type addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeMessages struct {
	Code    int64                                                                                                                           `json:"code,required"`
	Message string                                                                                                                          `json:"message,required"`
	JSON    addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeMessages]
type addressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeSuccess bool

const (
	AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeSuccessTrue AddressPrefixBGPStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseEnvelopeSuccess = true
)
