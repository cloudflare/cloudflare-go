// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// SpeedSettingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpeedSettingService] method instead.
type SpeedSettingService struct {
	Options         []option.RequestOption
	Transformations *SpeedSettingTransformationService
}

// NewSpeedSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpeedSettingService(opts ...option.RequestOption) (r *SpeedSettingService) {
	r = &SpeedSettingService{}
	r.Options = opts
	r.Transformations = NewSpeedSettingTransformationService(opts...)
	return
}
