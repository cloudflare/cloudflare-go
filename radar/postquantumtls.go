// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
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
// exchange. Returns information about the negotiated key exchange algorithm,
// whether it uses PQ cryptography, and any detected TLS implementation bugs (Split
// ClientHello, HRR failure, etc.).
func (r *PostQuantumTLSService) Support(ctx context.Context, query PostQuantumTLSSupportParams, opts ...option.RequestOption) (res *PostQuantumTLSSupportResponse, err error) {
	var env PostQuantumTLSSupportResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/post_quantum/tls/support"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type PostQuantumTLSSupportResponse struct {
	Bugs PostQuantumTLSSupportResponseBugs `json:"bugs" api:"required"`
	// The host that was tested
	Host string `json:"host" api:"required"`
	// TLS CurveID of the negotiated key exchange
	Kex float64 `json:"kex" api:"required"`
	// Human-readable name of the key exchange algorithm
	KexName string `json:"kexName" api:"required"`
	// Whether the negotiated key exchange uses Post-Quantum cryptography (specifically
	// X25519MLKEM768)
	Pq   bool                              `json:"pq" api:"required"`
	JSON postQuantumTLSSupportResponseJSON `json:"-"`
}

// postQuantumTLSSupportResponseJSON contains the JSON metadata for the struct
// [PostQuantumTLSSupportResponse]
type postQuantumTLSSupportResponseJSON struct {
	Bugs        apijson.Field
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

type PostQuantumTLSSupportResponseBugs struct {
	// Server sends a HelloRetryRequest but fails to complete the handshake after the
	// client sends the second ClientHello. Often caused by non-compliant TLS 1.3
	// implementations on shared hosting providers.
	HrrFailure bool `json:"hrrFailure" api:"required"`
	// Server rejects fragmented ClientHello caused by large PQ keyshare, but accepts
	// classical (non-PQ) handshakes. Typically caused by middleboxes or firewalls that
	// cannot reassemble split TLS ClientHello messages.
	SplitClientHello bool `json:"splitClientHello" api:"required"`
	// Server cannot handle an unknown key exchange algorithm in the ClientHello
	// keyshare extension. Compliant servers should respond with HelloRetryRequest for
	// a supported algorithm.
	UnknownKeyshare bool                                  `json:"unknownKeyshare" api:"required"`
	JSON            postQuantumTLSSupportResponseBugsJSON `json:"-"`
}

// postQuantumTLSSupportResponseBugsJSON contains the JSON metadata for the struct
// [PostQuantumTLSSupportResponseBugs]
type postQuantumTLSSupportResponseBugsJSON struct {
	HrrFailure       apijson.Field
	SplitClientHello apijson.Field
	UnknownKeyshare  apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PostQuantumTLSSupportResponseBugs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumTLSSupportResponseBugsJSON) RawJSON() string {
	return r.raw
}

type PostQuantumTLSSupportParams struct {
	// Hostname or IP address to test for Post-Quantum TLS support, optionally with
	// port (defaults to 443).
	Host param.Field[string] `query:"host" api:"required"`
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
	Result  PostQuantumTLSSupportResponse             `json:"result" api:"required"`
	Success bool                                      `json:"success" api:"required"`
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
