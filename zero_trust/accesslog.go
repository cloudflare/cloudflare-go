// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// AccessLogService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccessLogService] method instead.
type AccessLogService struct {
	Options        []option.RequestOption
	AccessRequests *AccessLogAccessRequestService
}

// NewAccessLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessLogService(opts ...option.RequestOption) (r *AccessLogService) {
	r = &AccessLogService{}
	r.Options = opts
	r.AccessRequests = NewAccessLogAccessRequestService(opts...)
	return
}
