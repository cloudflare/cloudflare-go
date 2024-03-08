// File generated from our OpenAPI spec by Stainless.

package alerting

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// V3DestinationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewV3DestinationService] method
// instead.
type V3DestinationService struct {
	Options   []option.RequestOption
	Eligible  *V3DestinationEligibleService
	Pagerduty *V3DestinationPagerdutyService
	Webhooks  *V3DestinationWebhookService
}

// NewV3DestinationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV3DestinationService(opts ...option.RequestOption) (r *V3DestinationService) {
	r = &V3DestinationService{}
	r.Options = opts
	r.Eligible = NewV3DestinationEligibleService(opts...)
	r.Pagerduty = NewV3DestinationPagerdutyService(opts...)
	r.Webhooks = NewV3DestinationWebhookService(opts...)
	return
}
