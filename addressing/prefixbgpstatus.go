// File generated from our OpenAPI spec by Stainless.

package addressing

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// PrefixBGPStatusService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPrefixBGPStatusService] method
// instead.
type PrefixBGPStatusService struct {
	Options []option.RequestOption
}

// NewPrefixBGPStatusService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrefixBGPStatusService(opts ...option.RequestOption) (r *PrefixBGPStatusService) {
	r = &PrefixBGPStatusService{}
	r.Options = opts
	return
}

// Advertise or withdraw BGP route for a prefix.
func (r *PrefixBGPStatusService) Edit(ctx context.Context, prefixID string, params PrefixBGPStatusEditParams, opts ...option.RequestOption) (res *PrefixBGPStatusEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PrefixBGPStatusEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List the current advertisement state for a prefix.
func (r *PrefixBGPStatusService) Get(ctx context.Context, prefixID string, query PrefixBGPStatusGetParams, opts ...option.RequestOption) (res *PrefixBGPStatusGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PrefixBGPStatusGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", query.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PrefixBGPStatusEditResponse struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                       `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 prefixBGPStatusEditResponseJSON `json:"-"`
}

// prefixBGPStatusEditResponseJSON contains the JSON metadata for the struct
// [PrefixBGPStatusEditResponse]
type prefixBGPStatusEditResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PrefixBGPStatusEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPStatusEditResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPStatusGetResponse struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                      `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 prefixBGPStatusGetResponseJSON `json:"-"`
}

// prefixBGPStatusGetResponseJSON contains the JSON metadata for the struct
// [PrefixBGPStatusGetResponse]
type prefixBGPStatusGetResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PrefixBGPStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPStatusGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPStatusEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Enablement of prefix advertisement to the Internet.
	Advertised param.Field[bool] `json:"advertised,required"`
}

func (r PrefixBGPStatusEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixBGPStatusEditResponseEnvelope struct {
	Errors   []PrefixBGPStatusEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixBGPStatusEditResponseEnvelopeMessages `json:"messages,required"`
	Result   PrefixBGPStatusEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PrefixBGPStatusEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    prefixBGPStatusEditResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPStatusEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixBGPStatusEditResponseEnvelope]
type prefixBGPStatusEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPStatusEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPStatusEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPStatusEditResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    prefixBGPStatusEditResponseEnvelopeErrorsJSON `json:"-"`
}

// prefixBGPStatusEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrefixBGPStatusEditResponseEnvelopeErrors]
type prefixBGPStatusEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPStatusEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPStatusEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPStatusEditResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    prefixBGPStatusEditResponseEnvelopeMessagesJSON `json:"-"`
}

// prefixBGPStatusEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PrefixBGPStatusEditResponseEnvelopeMessages]
type prefixBGPStatusEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPStatusEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPStatusEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPStatusEditResponseEnvelopeSuccess bool

const (
	PrefixBGPStatusEditResponseEnvelopeSuccessTrue PrefixBGPStatusEditResponseEnvelopeSuccess = true
)

type PrefixBGPStatusGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixBGPStatusGetResponseEnvelope struct {
	Errors   []PrefixBGPStatusGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixBGPStatusGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PrefixBGPStatusGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PrefixBGPStatusGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    prefixBGPStatusGetResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPStatusGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PrefixBGPStatusGetResponseEnvelope]
type prefixBGPStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPStatusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPStatusGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    prefixBGPStatusGetResponseEnvelopeErrorsJSON `json:"-"`
}

// prefixBGPStatusGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrefixBGPStatusGetResponseEnvelopeErrors]
type prefixBGPStatusGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPStatusGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPStatusGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPStatusGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    prefixBGPStatusGetResponseEnvelopeMessagesJSON `json:"-"`
}

// prefixBGPStatusGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PrefixBGPStatusGetResponseEnvelopeMessages]
type prefixBGPStatusGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPStatusGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPStatusGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPStatusGetResponseEnvelopeSuccess bool

const (
	PrefixBGPStatusGetResponseEnvelopeSuccessTrue PrefixBGPStatusGetResponseEnvelopeSuccess = true
)
