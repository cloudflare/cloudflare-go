// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *SecondaryDNSForceAxfrService) New(ctx context.Context, body SecondaryDNSForceAxfrNewParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSForceAxfrNewResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/force_axfr", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSForceAxfrNewParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
}

type SecondaryDNSForceAxfrNewResponseEnvelope struct {
	Errors   []SecondaryDNSForceAxfrNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSForceAxfrNewResponseEnvelopeMessages `json:"messages,required"`
	// When force_axfr query parameter is set to true, the response is a simple string
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSForceAxfrNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSForceAxfrNewResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSForceAxfrNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSForceAxfrNewResponseEnvelope]
type secondaryDNSForceAxfrNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSForceAxfrNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDNSForceAxfrNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SecondaryDNSForceAxfrNewResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    secondaryDNSForceAxfrNewResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSForceAxfrNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SecondaryDNSForceAxfrNewResponseEnvelopeErrors]
type secondaryDNSForceAxfrNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSForceAxfrNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDNSForceAxfrNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SecondaryDNSForceAxfrNewResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    secondaryDNSForceAxfrNewResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSForceAxfrNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SecondaryDNSForceAxfrNewResponseEnvelopeMessages]
type secondaryDNSForceAxfrNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSForceAxfrNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDNSForceAxfrNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SecondaryDNSForceAxfrNewResponseEnvelopeSuccess bool

const (
	SecondaryDNSForceAxfrNewResponseEnvelopeSuccessTrue SecondaryDNSForceAxfrNewResponseEnvelopeSuccess = true
)
