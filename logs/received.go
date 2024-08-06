// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ReceivedService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReceivedService] method instead.
type ReceivedService struct {
	Options []option.RequestOption
	Fields  *ReceivedFieldService
}

// NewReceivedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReceivedService(opts ...option.RequestOption) (r *ReceivedService) {
	r = &ReceivedService{}
	r.Options = opts
	r.Fields = NewReceivedFieldService(opts...)
	return
}
