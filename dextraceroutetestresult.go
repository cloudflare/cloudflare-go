// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DexTracerouteTestResultService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewDexTracerouteTestResultService] method instead.
type DexTracerouteTestResultService struct {
	Options     []option.RequestOption
	NetworkPath *DexTracerouteTestResultNetworkPathService
}

// NewDexTracerouteTestResultService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDexTracerouteTestResultService(opts ...option.RequestOption) (r *DexTracerouteTestResultService) {
	r = &DexTracerouteTestResultService{}
	r.Options = opts
	r.NetworkPath = NewDexTracerouteTestResultNetworkPathService(opts...)
	return
}
