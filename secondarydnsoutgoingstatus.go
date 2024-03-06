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

// SecondaryDNSOutgoingStatusService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSecondaryDNSOutgoingStatusService] method instead.
type SecondaryDNSOutgoingStatusService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSOutgoingStatusService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSecondaryDNSOutgoingStatusService(opts ...option.RequestOption) (r *SecondaryDNSOutgoingStatusService) {
	r = &SecondaryDNSOutgoingStatusService{}
	r.Options = opts
	return
}

// Get primary zone transfer status.
func (r *SecondaryDNSOutgoingStatusService) Get(ctx context.Context, query SecondaryDNSOutgoingStatusGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingStatusGetResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/status", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSOutgoingStatusGetParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
}

type SecondaryDNSOutgoingStatusGetResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingStatusGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingStatusGetResponseEnvelopeMessages `json:"messages,required"`
	// The zone transfer status of a primary zone
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingStatusGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingStatusGetResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingStatusGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SecondaryDNSOutgoingStatusGetResponseEnvelope]
type secondaryDNSOutgoingStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingStatusGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    secondaryDNSOutgoingStatusGetResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingStatusGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SecondaryDNSOutgoingStatusGetResponseEnvelopeErrors]
type secondaryDNSOutgoingStatusGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingStatusGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingStatusGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    secondaryDNSOutgoingStatusGetResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingStatusGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SecondaryDNSOutgoingStatusGetResponseEnvelopeMessages]
type secondaryDNSOutgoingStatusGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingStatusGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingStatusGetResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingStatusGetResponseEnvelopeSuccessTrue SecondaryDNSOutgoingStatusGetResponseEnvelopeSuccess = true
)
