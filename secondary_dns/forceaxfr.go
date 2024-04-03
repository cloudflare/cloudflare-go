// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// ForceAXFRService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewForceAXFRService] method instead.
type ForceAXFRService struct {
	Options []option.RequestOption
}

// NewForceAXFRService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewForceAXFRService(opts ...option.RequestOption) (r *ForceAXFRService) {
	r = &ForceAXFRService{}
	r.Options = opts
	return
}

// Sends AXFR zone transfer request to primary nameserver(s).
func (r *ForceAXFRService) New(ctx context.Context, params ForceAXFRNewParams, opts ...option.RequestOption) (res *SecondaryDNSForce, err error) {
	opts = append(r.Options[:], opts...)
	var env ForceAXFRNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/secondary_dns/force_axfr", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSForce = string

type ForceAXFRNewParams struct {
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r ForceAXFRNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ForceAXFRNewResponseEnvelope struct {
	Errors   []ForceAXFRNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ForceAXFRNewResponseEnvelopeMessages `json:"messages,required"`
	// When force_axfr query parameter is set to true, the response is a simple string
	Result SecondaryDNSForce `json:"result,required"`
	// Whether the API call was successful
	Success ForceAXFRNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    forceAXFRNewResponseEnvelopeJSON    `json:"-"`
}

// forceAXFRNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ForceAXFRNewResponseEnvelope]
type forceAXFRNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ForceAXFRNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r forceAXFRNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ForceAXFRNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    forceAXFRNewResponseEnvelopeErrorsJSON `json:"-"`
}

// forceAXFRNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ForceAXFRNewResponseEnvelopeErrors]
type forceAXFRNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ForceAXFRNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r forceAXFRNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ForceAXFRNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    forceAXFRNewResponseEnvelopeMessagesJSON `json:"-"`
}

// forceAXFRNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ForceAXFRNewResponseEnvelopeMessages]
type forceAXFRNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ForceAXFRNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r forceAXFRNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ForceAXFRNewResponseEnvelopeSuccess bool

const (
	ForceAXFRNewResponseEnvelopeSuccessTrue ForceAXFRNewResponseEnvelopeSuccess = true
)

func (r ForceAXFRNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ForceAXFRNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
