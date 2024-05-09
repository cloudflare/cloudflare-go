// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event_notifications

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// EventNotificationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventNotificationService] method instead.
type EventNotificationService struct {
	Options []option.RequestOption
	R2      *R2Service
}

// NewEventNotificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEventNotificationService(opts ...option.RequestOption) (r *EventNotificationService) {
	r = &EventNotificationService{}
	r.Options = opts
	r.R2 = NewR2Service(opts...)
	return
}
