// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountImageV2Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountImageV2Service] method
// instead.
type AccountImageV2Service struct {
	Options       []option.RequestOption
	DirectUploads *AccountImageV2DirectUploadService
}

// NewAccountImageV2Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountImageV2Service(opts ...option.RequestOption) (r *AccountImageV2Service) {
	r = &AccountImageV2Service{}
	r.Options = opts
	r.DirectUploads = NewAccountImageV2DirectUploadService(opts...)
	return
}
