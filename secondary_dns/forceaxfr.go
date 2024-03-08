// File generated from our OpenAPI spec by Stainless.

package secondary_dns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ForceAxfrService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewForceAxfrService] method instead.
type ForceAxfrService struct {
	Options []option.RequestOption
}

// NewForceAxfrService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewForceAxfrService(opts ...option.RequestOption) (r *ForceAxfrService) {
	r = &ForceAxfrService{}
	r.Options = opts
	return
}

// Sends AXFR zone transfer request to primary nameserver(s).
func (r *ForceAxfrService) New(ctx context.Context, body ForceAxfrNewParams, opts ...option.RequestOption) (res *SecondaryDNSForceResult, err error) {
	opts = append(r.Options[:], opts...)
	var env ForceAxfrNewResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/force_axfr", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSForceResult = string

type ForceAxfrNewParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
}

type ForceAxfrNewResponseEnvelope struct {
	Errors   []ForceAxfrNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ForceAxfrNewResponseEnvelopeMessages `json:"messages,required"`
	// When force_axfr query parameter is set to true, the response is a simple string
	Result SecondaryDNSForceResult `json:"result,required"`
	// Whether the API call was successful
	Success ForceAxfrNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    forceAxfrNewResponseEnvelopeJSON    `json:"-"`
}

// forceAxfrNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ForceAxfrNewResponseEnvelope]
type forceAxfrNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ForceAxfrNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r forceAxfrNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ForceAxfrNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    forceAxfrNewResponseEnvelopeErrorsJSON `json:"-"`
}

// forceAxfrNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ForceAxfrNewResponseEnvelopeErrors]
type forceAxfrNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ForceAxfrNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r forceAxfrNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ForceAxfrNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    forceAxfrNewResponseEnvelopeMessagesJSON `json:"-"`
}

// forceAxfrNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ForceAxfrNewResponseEnvelopeMessages]
type forceAxfrNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ForceAxfrNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r forceAxfrNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ForceAxfrNewResponseEnvelopeSuccess bool

const (
	ForceAxfrNewResponseEnvelopeSuccessTrue ForceAxfrNewResponseEnvelopeSuccess = true
)
