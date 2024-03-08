// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// D1Service contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewD1Service] method instead.
type D1Service struct {
	Options  []option.RequestOption
	Database *D1DatabaseService
}

// NewD1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewD1Service(opts ...option.RequestOption) (r *D1Service) {
	r = &D1Service{}
	r.Options = opts
	r.Database = NewD1DatabaseService(opts...)
	return
}
