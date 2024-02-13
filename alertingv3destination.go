// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AlertingV3DestinationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAlertingV3DestinationService]
// method instead.
type AlertingV3DestinationService struct {
	Options     []option.RequestOption
	Eligibles   *AlertingV3DestinationEligibleService
	Pagerduties *AlertingV3DestinationPagerdutyService
	Webhooks    *AlertingV3DestinationWebhookService
}

// NewAlertingV3DestinationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAlertingV3DestinationService(opts ...option.RequestOption) (r *AlertingV3DestinationService) {
	r = &AlertingV3DestinationService{}
	r.Options = opts
	r.Eligibles = NewAlertingV3DestinationEligibleService(opts...)
	r.Pagerduties = NewAlertingV3DestinationPagerdutyService(opts...)
	r.Webhooks = NewAlertingV3DestinationWebhookService(opts...)
	return
}
