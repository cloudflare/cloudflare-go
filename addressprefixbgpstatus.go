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

// Advertise or withdraw BGP route for a prefix.
func (r *AddressPrefixBGPStatusService) Edit(ctx context.Context, accountID string, prefixID string, body AddressPrefixBGPStatusEditParams, opts ...option.RequestOption) (res *AddressPrefixBGPStatusEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixBGPStatusEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List the current advertisement state for a prefix.
func (r *AddressPrefixBGPStatusService) Get(ctx context.Context, accountID string, prefixID string, opts ...option.RequestOption) (res *AddressPrefixBGPStatusGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixBGPStatusGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressPrefixBGPStatusEditResponse struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                              `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 addressPrefixBGPStatusEditResponseJSON `json:"-"`
}

// addressPrefixBGPStatusEditResponseJSON contains the JSON metadata for the struct
// [AddressPrefixBGPStatusEditResponse]
type addressPrefixBGPStatusEditResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPStatusGetResponse struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                             `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 addressPrefixBGPStatusGetResponseJSON `json:"-"`
}

// addressPrefixBGPStatusGetResponseJSON contains the JSON metadata for the struct
// [AddressPrefixBGPStatusGetResponse]
type addressPrefixBGPStatusGetResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPStatusEditParams struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised param.Field[bool] `json:"advertised,required"`
}

func (r AddressPrefixBGPStatusEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressPrefixBGPStatusEditResponseEnvelope struct {
	Errors   []AddressPrefixBGPStatusEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixBGPStatusEditResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixBGPStatusEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixBGPStatusEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixBGPStatusEditResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixBGPStatusEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressPrefixBGPStatusEditResponseEnvelope]
type addressPrefixBGPStatusEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPStatusEditResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    addressPrefixBGPStatusEditResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixBGPStatusEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressPrefixBGPStatusEditResponseEnvelopeErrors]
type addressPrefixBGPStatusEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPStatusEditResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressPrefixBGPStatusEditResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixBGPStatusEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressPrefixBGPStatusEditResponseEnvelopeMessages]
type addressPrefixBGPStatusEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixBGPStatusEditResponseEnvelopeSuccess bool

const (
	AddressPrefixBGPStatusEditResponseEnvelopeSuccessTrue AddressPrefixBGPStatusEditResponseEnvelopeSuccess = true
)

type AddressPrefixBGPStatusGetResponseEnvelope struct {
	Errors   []AddressPrefixBGPStatusGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixBGPStatusGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixBGPStatusGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixBGPStatusGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixBGPStatusGetResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixBGPStatusGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressPrefixBGPStatusGetResponseEnvelope]
type addressPrefixBGPStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPStatusGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    addressPrefixBGPStatusGetResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixBGPStatusGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressPrefixBGPStatusGetResponseEnvelopeErrors]
type addressPrefixBGPStatusGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPStatusGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    addressPrefixBGPStatusGetResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixBGPStatusGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AddressPrefixBGPStatusGetResponseEnvelopeMessages]
type addressPrefixBGPStatusGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPStatusGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixBGPStatusGetResponseEnvelopeSuccess bool

const (
	AddressPrefixBGPStatusGetResponseEnvelopeSuccessTrue AddressPrefixBGPStatusGetResponseEnvelopeSuccess = true
)
