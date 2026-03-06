// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// InstanceItemService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceItemService] method instead.
type InstanceItemService struct {
	Options []option.RequestOption
}

// NewInstanceItemService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceItemService(opts ...option.RequestOption) (r *InstanceItemService) {
	r = &InstanceItemService{}
	r.Options = opts
	return
}
