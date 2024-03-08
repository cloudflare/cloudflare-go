// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// ZeroTrustAccessLogService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustAccessLogService] method
// instead.
type ZeroTrustAccessLogService struct {
	Options        []option.RequestOption
	AccessRequests *ZeroTrustAccessLogAccessRequestService
}

// NewZeroTrustAccessLogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustAccessLogService(opts ...option.RequestOption) (r *ZeroTrustAccessLogService) {
	r = &ZeroTrustAccessLogService{}
	r.Options = opts
	r.AccessRequests = NewZeroTrustAccessLogAccessRequestService(opts...)
	return
}
