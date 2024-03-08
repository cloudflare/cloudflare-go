// File generated from our OpenAPI spec by Stainless.

package secondary_dns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// OutgoingStatusService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOutgoingStatusService] method
// instead.
type OutgoingStatusService struct {
	Options []option.RequestOption
}

// NewOutgoingStatusService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOutgoingStatusService(opts ...option.RequestOption) (r *OutgoingStatusService) {
	r = &OutgoingStatusService{}
	r.Options = opts
	return
}

// Get primary zone transfer status.
func (r *OutgoingStatusService) Get(ctx context.Context, query OutgoingStatusGetParams, opts ...option.RequestOption) (res *SecondaryDNSEnableTransferResult, err error) {
	opts = append(r.Options[:], opts...)
	var env OutgoingStatusGetResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/status", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OutgoingStatusGetParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
}

type OutgoingStatusGetResponseEnvelope struct {
	Errors   []OutgoingStatusGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OutgoingStatusGetResponseEnvelopeMessages `json:"messages,required"`
	// The zone transfer status of a primary zone
	Result SecondaryDNSEnableTransferResult `json:"result,required"`
	// Whether the API call was successful
	Success OutgoingStatusGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    outgoingStatusGetResponseEnvelopeJSON    `json:"-"`
}

// outgoingStatusGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [OutgoingStatusGetResponseEnvelope]
type outgoingStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutgoingStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingStatusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OutgoingStatusGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    outgoingStatusGetResponseEnvelopeErrorsJSON `json:"-"`
}

// outgoingStatusGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OutgoingStatusGetResponseEnvelopeErrors]
type outgoingStatusGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutgoingStatusGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingStatusGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OutgoingStatusGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    outgoingStatusGetResponseEnvelopeMessagesJSON `json:"-"`
}

// outgoingStatusGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OutgoingStatusGetResponseEnvelopeMessages]
type outgoingStatusGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutgoingStatusGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingStatusGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OutgoingStatusGetResponseEnvelopeSuccess bool

const (
	OutgoingStatusGetResponseEnvelopeSuccessTrue OutgoingStatusGetResponseEnvelopeSuccess = true
)
