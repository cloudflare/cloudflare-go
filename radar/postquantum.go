// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// PostQuantumService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPostQuantumService] method instead.
type PostQuantumService struct {
	Options []option.RequestOption
	Origin  *PostQuantumOriginService
	TLS     *PostQuantumTLSService
}

// NewPostQuantumService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPostQuantumService(opts ...option.RequestOption) (r *PostQuantumService) {
	r = &PostQuantumService{}
	r.Options = opts
	r.Origin = NewPostQuantumOriginService(opts...)
	r.TLS = NewPostQuantumTLSService(opts...)
	return
}
