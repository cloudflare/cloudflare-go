// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package k2

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// K2Service contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewK2Service] method instead.
type K2Service struct {
	Options []option.RequestOption
	Streams *StreamService
}

// NewK2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewK2Service(opts ...option.RequestOption) (r *K2Service) {
	r = &K2Service{}
	r.Options = opts
	r.Streams = NewStreamService(opts...)
	return
}
