// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// PostQuantumTLSService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPostQuantumTLSService] method instead.
type PostQuantumTLSService struct {
	Options []option.RequestOption
}

// NewPostQuantumTLSService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPostQuantumTLSService(opts ...option.RequestOption) (r *PostQuantumTLSService) {
	r = &PostQuantumTLSService{}
	r.Options = opts
	return
}

// Tests whether a hostname or IP address supports Post-Quantum (PQ) TLS key
// exchange. Returns information about the negotiated key exchange algorithm and
// whether it uses PQ cryptography.
func (r *PostQuantumTLSService) Support(ctx context.Context, query PostQuantumTLSSupportParams, opts ...option.RequestOption) (res *PostQuantumTLSSupportResponse, err error) {
	var env PostQuantumTLSSupportResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/post_quantum/tls/support"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PostQuantumTLSSupportResponse struct {
	// The host that was tested
	Host string `json:"host,required"`
	// TLS CurveID of the negotiated key exchange
	Kex float64 `json:"kex,required"`
	// Human-readable name of the key exchange algorithm
	KexName string `json:"kexName,required"`
	// Whether the negotiated key exchange uses Post-Quantum cryptography
	Pq   bool                              `json:"pq,required"`
	JSON postQuantumTLSSupportResponseJSON `json:"-"`
}

// postQuantumTLSSupportResponseJSON contains the JSON metadata for the struct
// [PostQuantumTLSSupportResponse]
type postQuantumTLSSupportResponseJSON struct {
	Host        apijson.Field
	Kex         apijson.Field
	KexName     apijson.Field
	Pq          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostQuantumTLSSupportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumTLSSupportResponseJSON) RawJSON() string {
	return r.raw
}

type PostQuantumTLSSupportParams struct {
	// Hostname or IP address to test for Post-Quantum TLS support, optionally with
	// port (defaults to 443).
	Host param.Field[string] `query:"host,required"`
}

// URLQuery serializes [PostQuantumTLSSupportParams]'s query parameters as
// `url.Values`.
func (r PostQuantumTLSSupportParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PostQuantumTLSSupportResponseEnvelope struct {
	Result  PostQuantumTLSSupportResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    postQuantumTLSSupportResponseEnvelopeJSON `json:"-"`
}

// postQuantumTLSSupportResponseEnvelopeJSON contains the JSON metadata for the
// struct [PostQuantumTLSSupportResponseEnvelope]
type postQuantumTLSSupportResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostQuantumTLSSupportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumTLSSupportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
