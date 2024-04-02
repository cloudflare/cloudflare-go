// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// WaitingRoomService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWaitingRoomService] method
// instead.
type WaitingRoomService struct {
	Options  []option.RequestOption
	Page     *PageService
	Events   *EventService
	Rules    *RuleService
	Statuses *StatusService
	Settings *SettingService
}

// NewWaitingRoomService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWaitingRoomService(opts ...option.RequestOption) (r *WaitingRoomService) {
	r = &WaitingRoomService{}
	r.Options = opts
	r.Page = NewPageService(opts...)
	r.Events = NewEventService(opts...)
	r.Rules = NewRuleService(opts...)
	r.Statuses = NewStatusService(opts...)
	r.Settings = NewSettingService(opts...)
	return
}
