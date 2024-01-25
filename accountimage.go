// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountImageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountImageService] method
// instead.
type AccountImageService struct {
	Options []option.RequestOption
	V1s     *AccountImageV1Service
	V2s     *AccountImageV2Service
}

// NewAccountImageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountImageService(opts ...option.RequestOption) (r *AccountImageService) {
	r = &AccountImageService{}
	r.Options = opts
	r.V1s = NewAccountImageV1Service(opts...)
	r.V2s = NewAccountImageV2Service(opts...)
	return
}
