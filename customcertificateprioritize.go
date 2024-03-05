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

// CustomCertificatePrioritizeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewCustomCertificatePrioritizeService] method instead.
type CustomCertificatePrioritizeService struct {
	Options []option.RequestOption
}

// NewCustomCertificatePrioritizeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCustomCertificatePrioritizeService(opts ...option.RequestOption) (r *CustomCertificatePrioritizeService) {
	r = &CustomCertificatePrioritizeService{}
	r.Options = opts
	return
}

// If a zone has multiple SSL certificates, you can set the order in which they
// should be used during a request. The higher priority will break ties across
// overlapping 'legacy_custom' certificates.
func (r *CustomCertificatePrioritizeService) Update(ctx context.Context, params CustomCertificatePrioritizeUpdateParams, opts ...option.RequestOption) (res *[]TLSCertificatesAndHostnamesCustomCertificate, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomCertificatePrioritizeUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_certificates/prioritize", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CustomCertificatePrioritizeUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Array of ordered certificates.
	Certificates param.Field[[]CustomCertificatePrioritizeUpdateParamsCertificate] `json:"certificates,required"`
}

func (r CustomCertificatePrioritizeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomCertificatePrioritizeUpdateParamsCertificate struct {
	// The order/priority in which the certificate will be used in a request. The
	// higher priority will break ties across overlapping 'legacy_custom' certificates,
	// but 'legacy_custom' certificates will always supercede 'sni_custom'
	// certificates.
	Priority param.Field[float64] `json:"priority"`
}

func (r CustomCertificatePrioritizeUpdateParamsCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomCertificatePrioritizeUpdateResponseEnvelope struct {
	Errors   []CustomCertificatePrioritizeUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomCertificatePrioritizeUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []TLSCertificatesAndHostnamesCustomCertificate              `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CustomCertificatePrioritizeUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CustomCertificatePrioritizeUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       customCertificatePrioritizeUpdateResponseEnvelopeJSON       `json:"-"`
}

// customCertificatePrioritizeUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [CustomCertificatePrioritizeUpdateResponseEnvelope]
type customCertificatePrioritizeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificatePrioritizeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificatePrioritizeUpdateResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    customCertificatePrioritizeUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// customCertificatePrioritizeUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CustomCertificatePrioritizeUpdateResponseEnvelopeErrors]
type customCertificatePrioritizeUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificatePrioritizeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomCertificatePrioritizeUpdateResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    customCertificatePrioritizeUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// customCertificatePrioritizeUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CustomCertificatePrioritizeUpdateResponseEnvelopeMessages]
type customCertificatePrioritizeUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificatePrioritizeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomCertificatePrioritizeUpdateResponseEnvelopeSuccess bool

const (
	CustomCertificatePrioritizeUpdateResponseEnvelopeSuccessTrue CustomCertificatePrioritizeUpdateResponseEnvelopeSuccess = true
)

type CustomCertificatePrioritizeUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       customCertificatePrioritizeUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// customCertificatePrioritizeUpdateResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [CustomCertificatePrioritizeUpdateResponseEnvelopeResultInfo]
type customCertificatePrioritizeUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCertificatePrioritizeUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
