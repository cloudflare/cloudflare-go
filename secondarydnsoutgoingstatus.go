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
func (r *SecondaryDNSOutgoingStatusService) SecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatus(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/status", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeMessages `json:"messages,required"`
	// The zone transfer status of a primary zone
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelope]
type secondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeErrors struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    secondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeErrors]
type secondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeMessages struct {
	Code    int64                                                                                                      `json:"code,required"`
	Message string                                                                                                     `json:"message,required"`
	JSON    secondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeMessages]
type secondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeSuccessTrue SecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseEnvelopeSuccess = true
)
