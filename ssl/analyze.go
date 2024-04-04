// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssl

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// AnalyzeService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAnalyzeService] method instead.
type AnalyzeService struct {
	Options []option.RequestOption
}

// NewAnalyzeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAnalyzeService(opts ...option.RequestOption) (r *AnalyzeService) {
	r = &AnalyzeService{}
	r.Options = opts
	return
}

// Returns the set of hostnames, the signature algorithm, and the expiration date
// of the certificate.
func (r *AnalyzeService) New(ctx context.Context, params AnalyzeNewParams, opts ...option.RequestOption) (res *AnalyzeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AnalyzeNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/analyze", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ssl.AnalyzeNewResponseUnknown] or [shared.UnionString].
type AnalyzeNewResponse interface {
	ImplementsSSLAnalyzeNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyzeNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AnalyzeNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[shared.UnnamedSchemaRef78] `json:"bundle_method"`
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate"`
}

func (r AnalyzeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AnalyzeNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   AnalyzeNewResponse    `json:"result,required"`
	// Whether the API call was successful
	Success AnalyzeNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    analyzeNewResponseEnvelopeJSON    `json:"-"`
}

// analyzeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AnalyzeNewResponseEnvelope]
type analyzeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyzeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyzeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AnalyzeNewResponseEnvelopeSuccess bool

const (
	AnalyzeNewResponseEnvelopeSuccessTrue AnalyzeNewResponseEnvelopeSuccess = true
)

func (r AnalyzeNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AnalyzeNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
