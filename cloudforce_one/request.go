// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// RequestService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestService] method instead.
type RequestService struct {
	Options  []option.RequestOption
	Message  *RequestMessageService
	Priority *RequestPriorityService
	Assets   *RequestAssetService
}

// NewRequestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRequestService(opts ...option.RequestOption) (r *RequestService) {
	r = &RequestService{}
	r.Options = opts
	r.Message = NewRequestMessageService(opts...)
	r.Priority = NewRequestPriorityService(opts...)
	r.Assets = NewRequestAssetService(opts...)
	return
}
