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

// SecondaryDNSOutgoingForceNotifyService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSecondaryDNSOutgoingForceNotifyService] method instead.
type SecondaryDNSOutgoingForceNotifyService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSOutgoingForceNotifyService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSecondaryDNSOutgoingForceNotifyService(opts ...option.RequestOption) (r *SecondaryDNSOutgoingForceNotifyService) {
	r = &SecondaryDNSOutgoingForceNotifyService{}
	r.Options = opts
	return
}

// Notifies the secondary nameserver(s) and clears IXFR backlog of primary zone.
func (r *SecondaryDNSOutgoingForceNotifyService) SecondaryDNSPrimaryZoneForceDNSNotify(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/force_notify", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeMessages `json:"messages,required"`
	// When force_notify query parameter is set to true, the response is a simple
	// string
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelope]
type secondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeErrors struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    secondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeErrors]
type secondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeMessages struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    secondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeMessages]
type secondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeSuccessTrue SecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseEnvelopeSuccess = true
)
