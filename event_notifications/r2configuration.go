// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event_notifications

import (
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// R2ConfigurationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewR2ConfigurationService] method instead.
type R2ConfigurationService struct {
	Options []option.RequestOption
	Queues  *R2ConfigurationQueueService
}

// NewR2ConfigurationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewR2ConfigurationService(opts ...option.RequestOption) (r *R2ConfigurationService) {
	r = &R2ConfigurationService{}
	r.Options = opts
	r.Queues = NewR2ConfigurationQueueService(opts...)
	return
}
