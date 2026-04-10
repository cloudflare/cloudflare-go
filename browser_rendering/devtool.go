// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// DevtoolService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevtoolService] method instead.
type DevtoolService struct {
	Options []option.RequestOption
	Session *DevtoolSessionService
	Browser *DevtoolBrowserService
}

// NewDevtoolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDevtoolService(opts ...option.RequestOption) (r *DevtoolService) {
	r = &DevtoolService{}
	r.Options = opts
	r.Session = NewDevtoolSessionService(opts...)
	r.Browser = NewDevtoolBrowserService(opts...)
	return
}
