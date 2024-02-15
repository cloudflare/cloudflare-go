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

// SecondaryDNSForceAxfrService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSecondaryDNSForceAxfrService]
// method instead.
type SecondaryDNSForceAxfrService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSForceAxfrService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSecondaryDNSForceAxfrService(opts ...option.RequestOption) (r *SecondaryDNSForceAxfrService) {
	r = &SecondaryDNSForceAxfrService{}
	r.Options = opts
	return
}

// Sends AXFR zone transfer request to primary nameserver(s).
func (r *SecondaryDNSForceAxfrService) SecondaryDNSSecondaryZoneForceAxfr(ctx context.Context, zoneID interface{}, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/force_axfr", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelope struct {
	Errors   []SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeMessages `json:"messages,required"`
	// When force_axfr query parameter is set to true, the response is a simple string
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelope]
type secondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeErrors struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    secondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeErrors]
type secondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeMessages struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    secondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeMessages]
type secondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeSuccess bool

const (
	SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeSuccessTrue SecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseEnvelopeSuccess = true
)
