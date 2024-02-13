// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// SSLAnalyzeService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSSLAnalyzeService] method instead.
type SSLAnalyzeService struct {
	Options []option.RequestOption
}

// NewSSLAnalyzeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSSLAnalyzeService(opts ...option.RequestOption) (r *SSLAnalyzeService) {
	r = &SSLAnalyzeService{}
	r.Options = opts
	return
}

// Returns the set of hostnames, the signature algorithm, and the expiration date
// of the certificate.
func (r *SSLAnalyzeService) AnalyzeCertificateAnalyzeCertificate(ctx context.Context, zoneID string, body SSLAnalyzeAnalyzeCertificateAnalyzeCertificateParams, opts ...option.RequestOption) (res *SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/analyze", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseUnknown] or
// [shared.UnionString].
type SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponse interface {
	ImplementsSSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SSLAnalyzeAnalyzeCertificateAnalyzeCertificateParams struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[SSLAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethod] `json:"bundle_method"`
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate"`
}

func (r SSLAnalyzeAnalyzeCertificateAnalyzeCertificateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type SSLAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethod string

const (
	SSLAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethodUbiquitous SSLAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethod = "ubiquitous"
	SSLAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethodOptimal    SSLAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethod = "optimal"
	SSLAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethodForce      SSLAnalyzeAnalyzeCertificateAnalyzeCertificateParamsBundleMethod = "force"
)

type SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelope struct {
	Errors   []SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeJSON    `json:"-"`
}

// sslAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelope]
type sslAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeErrors struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    sslAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeErrorsJSON `json:"-"`
}

// sslAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeErrors]
type sslAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeMessages struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    sslAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeMessagesJSON `json:"-"`
}

// sslAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeMessages]
type sslAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeSuccess bool

const (
	SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeSuccessTrue SSLAnalyzeAnalyzeCertificateAnalyzeCertificateResponseEnvelopeSuccess = true
)
