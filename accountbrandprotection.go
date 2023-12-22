// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountBrandProtectionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountBrandProtectionService]
// method instead.
type AccountBrandProtectionService struct {
	Options  []option.RequestOption
	Submits  *AccountBrandProtectionSubmitService
	URLInfos *AccountBrandProtectionURLInfoService
}

// NewAccountBrandProtectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountBrandProtectionService(opts ...option.RequestOption) (r *AccountBrandProtectionService) {
	r = &AccountBrandProtectionService{}
	r.Options = opts
	r.Submits = NewAccountBrandProtectionSubmitService(opts...)
	r.URLInfos = NewAccountBrandProtectionURLInfoService(opts...)
	return
}
