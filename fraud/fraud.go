// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fraud

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// FraudService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFraudService] method instead.
type FraudService struct {
	Options []option.RequestOption
}

// NewFraudService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFraudService(opts ...option.RequestOption) (r *FraudService) {
	r = &FraudService{}
	r.Options = opts
	return
}
