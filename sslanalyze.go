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
func (r *SSLAnalyzeService) New(ctx context.Context, params SSLAnalyzeNewParams, opts ...option.RequestOption) (res *SSLAnalyzeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLAnalyzeNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/analyze", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [SSLAnalyzeNewResponseUnknown] or [shared.UnionString].
type SSLAnalyzeNewResponse interface {
	ImplementsSSLAnalyzeNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SSLAnalyzeNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SSLAnalyzeNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[SSLAnalyzeNewParamsBundleMethod] `json:"bundle_method"`
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate"`
}

func (r SSLAnalyzeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type SSLAnalyzeNewParamsBundleMethod string

const (
	SSLAnalyzeNewParamsBundleMethodUbiquitous SSLAnalyzeNewParamsBundleMethod = "ubiquitous"
	SSLAnalyzeNewParamsBundleMethodOptimal    SSLAnalyzeNewParamsBundleMethod = "optimal"
	SSLAnalyzeNewParamsBundleMethodForce      SSLAnalyzeNewParamsBundleMethod = "force"
)

type SSLAnalyzeNewResponseEnvelope struct {
	Errors   []SSLAnalyzeNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLAnalyzeNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLAnalyzeNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SSLAnalyzeNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslAnalyzeNewResponseEnvelopeJSON    `json:"-"`
}

// sslAnalyzeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SSLAnalyzeNewResponseEnvelope]
type sslAnalyzeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLAnalyzeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLAnalyzeNewResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    sslAnalyzeNewResponseEnvelopeErrorsJSON `json:"-"`
}

// sslAnalyzeNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SSLAnalyzeNewResponseEnvelopeErrors]
type sslAnalyzeNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLAnalyzeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLAnalyzeNewResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    sslAnalyzeNewResponseEnvelopeMessagesJSON `json:"-"`
}

// sslAnalyzeNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SSLAnalyzeNewResponseEnvelopeMessages]
type sslAnalyzeNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLAnalyzeNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLAnalyzeNewResponseEnvelopeSuccess bool

const (
	SSLAnalyzeNewResponseEnvelopeSuccessTrue SSLAnalyzeNewResponseEnvelopeSuccess = true
)
