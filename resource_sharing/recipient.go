// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_sharing

import (
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// RecipientService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecipientService] method instead.
type RecipientService struct {
	Options []option.RequestOption
}

// NewRecipientService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecipientService(opts ...option.RequestOption) (r *RecipientService) {
	r = &RecipientService{}
	r.Options = opts
	return
}
