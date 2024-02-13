// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AlertingV3Service contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAlertingV3Service] method instead.
type AlertingV3Service struct {
	Options      []option.RequestOption
	Destinations *AlertingV3DestinationService
	Histories    *AlertingV3HistoryService
	Policies     *AlertingV3PolicyService
}

// NewAlertingV3Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlertingV3Service(opts ...option.RequestOption) (r *AlertingV3Service) {
	r = &AlertingV3Service{}
	r.Options = opts
	r.Destinations = NewAlertingV3DestinationService(opts...)
	r.Histories = NewAlertingV3HistoryService(opts...)
	r.Policies = NewAlertingV3PolicyService(opts...)
	return
}
