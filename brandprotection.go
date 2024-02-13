// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// BrandProtectionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewBrandProtectionService] method
// instead.
type BrandProtectionService struct {
	Options  []option.RequestOption
	Submits  *BrandProtectionSubmitService
	URLInfos *BrandProtectionURLInfoService
}

// NewBrandProtectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBrandProtectionService(opts ...option.RequestOption) (r *BrandProtectionService) {
	r = &BrandProtectionService{}
	r.Options = opts
	r.Submits = NewBrandProtectionSubmitService(opts...)
	r.URLInfos = NewBrandProtectionURLInfoService(opts...)
	return
}
