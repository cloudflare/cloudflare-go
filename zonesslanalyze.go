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

// ZoneSslAnalyzeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSslAnalyzeService] method
// instead.
type ZoneSslAnalyzeService struct {
	Options []option.RequestOption
}

// NewZoneSslAnalyzeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSslAnalyzeService(opts ...option.RequestOption) (r *ZoneSslAnalyzeService) {
	r = &ZoneSslAnalyzeService{}
	r.Options = opts
	return
}

// Returns the set of hostnames, the signature algorithm, and the expiration date
// of the certificate.
func (r *ZoneSslAnalyzeService) AnalyzeCertificateAnalyzeCertificate(ctx context.Context, identifier string, body ZoneSslAnalyzeAnalyzeCertificateAnalyzeCertificateParams, opts ...option.RequestOption) (res *CertificateAnalyzeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/analyze", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CertificateAnalyzeResponse struct {
	Errors   []CertificateAnalyzeResponseError   `json:"errors"`
	Messages []CertificateAnalyzeResponseMessage `json:"messages"`
	Result   interface{}                         `json:"result"`
	// Whether the API call was successful
	Success CertificateAnalyzeResponseSuccess `json:"success"`
	JSON    certificateAnalyzeResponseJSON    `json:"-"`
}

// certificateAnalyzeResponseJSON contains the JSON metadata for the struct
// [CertificateAnalyzeResponse]
type certificateAnalyzeResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAnalyzeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAnalyzeResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    certificateAnalyzeResponseErrorJSON `json:"-"`
}

// certificateAnalyzeResponseErrorJSON contains the JSON metadata for the struct
// [CertificateAnalyzeResponseError]
type certificateAnalyzeResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAnalyzeResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAnalyzeResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    certificateAnalyzeResponseMessageJSON `json:"-"`
}

// certificateAnalyzeResponseMessageJSON contains the JSON metadata for the struct
// [CertificateAnalyzeResponseMessage]
type certificateAnalyzeResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAnalyzeResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateAnalyzeResponseSuccess bool

const (
	CertificateAnalyzeResponseSuccessTrue CertificateAnalyzeResponseSuccess = true
)

type ZoneSslAnalyzeAnalyzeCertificateAnalyzeCertificateParams struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[ZoneSslAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethod] `json:"bundle_method"`
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate"`
}

func (r ZoneSslAnalyzeAnalyzeCertificateAnalyzeCertificateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type ZoneSslAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethod string

const (
	ZoneSslAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethodUbiquitous ZoneSslAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethod = "ubiquitous"
	ZoneSslAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethodOptimal    ZoneSslAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethod = "optimal"
	ZoneSslAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethodForce      ZoneSslAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethod = "force"
)
