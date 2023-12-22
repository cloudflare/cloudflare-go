// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAlertingV3DestinationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAlertingV3DestinationService] method instead.
type AccountAlertingV3DestinationService struct {
	Options     []option.RequestOption
	Eligibles   *AccountAlertingV3DestinationEligibleService
	Pagerduties *AccountAlertingV3DestinationPagerdutyService
	Webhooks    *AccountAlertingV3DestinationWebhookService
}

// NewAccountAlertingV3DestinationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAlertingV3DestinationService(opts ...option.RequestOption) (r *AccountAlertingV3DestinationService) {
	r = &AccountAlertingV3DestinationService{}
	r.Options = opts
	r.Eligibles = NewAccountAlertingV3DestinationEligibleService(opts...)
	r.Pagerduties = NewAccountAlertingV3DestinationPagerdutyService(opts...)
	r.Webhooks = NewAccountAlertingV3DestinationWebhookService(opts...)
	return
}
