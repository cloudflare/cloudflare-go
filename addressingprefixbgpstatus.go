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

// AddressingPrefixBGPStatusService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAddressingPrefixBGPStatusService] method instead.
type AddressingPrefixBGPStatusService struct {
	Options []option.RequestOption
}

// NewAddressingPrefixBGPStatusService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAddressingPrefixBGPStatusService(opts ...option.RequestOption) (r *AddressingPrefixBGPStatusService) {
	r = &AddressingPrefixBGPStatusService{}
	r.Options = opts
	return
}

// Advertise or withdraw BGP route for a prefix.
func (r *AddressingPrefixBGPStatusService) Edit(ctx context.Context, prefixID string, params AddressingPrefixBGPStatusEditParams, opts ...option.RequestOption) (res *AddressingPrefixBGPStatusEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBGPStatusEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List the current advertisement state for a prefix.
func (r *AddressingPrefixBGPStatusService) Get(ctx context.Context, prefixID string, query AddressingPrefixBGPStatusGetParams, opts ...option.RequestOption) (res *AddressingPrefixBGPStatusGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBGPStatusGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", query.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressingPrefixBGPStatusEditResponse struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                                 `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 addressingPrefixBGPStatusEditResponseJSON `json:"-"`
}

// addressingPrefixBGPStatusEditResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBGPStatusEditResponse]
type addressingPrefixBGPStatusEditResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixBGPStatusEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBGPStatusGetResponse struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                                `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 addressingPrefixBGPStatusGetResponseJSON `json:"-"`
}

// addressingPrefixBGPStatusGetResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBGPStatusGetResponse]
type addressingPrefixBGPStatusGetResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixBGPStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBGPStatusEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Enablement of prefix advertisement to the Internet.
	Advertised param.Field[bool] `json:"advertised,required"`
}

func (r AddressingPrefixBGPStatusEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressingPrefixBGPStatusEditResponseEnvelope struct {
	Errors   []AddressingPrefixBGPStatusEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixBGPStatusEditResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixBGPStatusEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixBGPStatusEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixBGPStatusEditResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBGPStatusEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixBGPStatusEditResponseEnvelope]
type addressingPrefixBGPStatusEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPStatusEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBGPStatusEditResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressingPrefixBGPStatusEditResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBGPStatusEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPStatusEditResponseEnvelopeErrors]
type addressingPrefixBGPStatusEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPStatusEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBGPStatusEditResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    addressingPrefixBGPStatusEditResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBGPStatusEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPStatusEditResponseEnvelopeMessages]
type addressingPrefixBGPStatusEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPStatusEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixBGPStatusEditResponseEnvelopeSuccess bool

const (
	AddressingPrefixBGPStatusEditResponseEnvelopeSuccessTrue AddressingPrefixBGPStatusEditResponseEnvelopeSuccess = true
)

type AddressingPrefixBGPStatusGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingPrefixBGPStatusGetResponseEnvelope struct {
	Errors   []AddressingPrefixBGPStatusGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixBGPStatusGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixBGPStatusGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixBGPStatusGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixBGPStatusGetResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBGPStatusGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixBGPStatusGetResponseEnvelope]
type addressingPrefixBGPStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBGPStatusGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressingPrefixBGPStatusGetResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBGPStatusGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPStatusGetResponseEnvelopeErrors]
type addressingPrefixBGPStatusGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPStatusGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBGPStatusGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    addressingPrefixBGPStatusGetResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBGPStatusGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPStatusGetResponseEnvelopeMessages]
type addressingPrefixBGPStatusGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPStatusGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixBGPStatusGetResponseEnvelopeSuccess bool

const (
	AddressingPrefixBGPStatusGetResponseEnvelopeSuccessTrue AddressingPrefixBGPStatusGetResponseEnvelopeSuccess = true
)
