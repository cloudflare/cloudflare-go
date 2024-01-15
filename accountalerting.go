// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAlertingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountAlertingService] method
// instead.
type AccountAlertingService struct {
	Options []option.RequestOption
	V3      *AccountAlertingV3Service
}

// NewAccountAlertingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountAlertingService(opts ...option.RequestOption) (r *AccountAlertingService) {
	r = &AccountAlertingService{}
	r.Options = opts
	r.V3 = NewAccountAlertingV3Service(opts...)
	return
}
