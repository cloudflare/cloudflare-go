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

// SecondaryDNSOutgoingDisableService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSecondaryDNSOutgoingDisableService] method instead.
type SecondaryDNSOutgoingDisableService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSOutgoingDisableService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSecondaryDNSOutgoingDisableService(opts ...option.RequestOption) (r *SecondaryDNSOutgoingDisableService) {
	r = &SecondaryDNSOutgoingDisableService{}
	r.Options = opts
	return
}

// Disable outgoing zone transfers for primary zone and clears IXFR backlog of
// primary zone.
func (r *SecondaryDNSOutgoingDisableService) SecondaryDNSPrimaryZoneDisableOutgoingZoneTransfers(ctx context.Context, zoneID interface{}, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/disable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeMessages `json:"messages,required"`
	// The zone transfer status of a primary zone
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelope]
type secondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeErrors struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    secondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeErrors]
type secondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeMessages struct {
	Code    int64                                                                                                      `json:"code,required"`
	Message string                                                                                                     `json:"message,required"`
	JSON    secondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeMessages]
type secondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeSuccessTrue SecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseEnvelopeSuccess = true
)
