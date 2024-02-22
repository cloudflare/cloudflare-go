// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DeviceService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDeviceService] method instead.
type DeviceService struct {
	Options       []option.RequestOption
	DEXTests      *DeviceDEXTestService
	Networks      *DeviceNetworkService
	Policies      *DevicePolicyService
	Postures      *DevicePostureService
	Revokes       *DeviceRevokeService
	Settings      *DeviceSettingService
	Unrevokes     *DeviceUnrevokeService
	OverrideCodes *DeviceOverrideCodeService
}

// NewDeviceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDeviceService(opts ...option.RequestOption) (r *DeviceService) {
	r = &DeviceService{}
	r.Options = opts
	r.DEXTests = NewDeviceDEXTestService(opts...)
	r.Networks = NewDeviceNetworkService(opts...)
	r.Policies = NewDevicePolicyService(opts...)
	r.Postures = NewDevicePostureService(opts...)
	r.Revokes = NewDeviceRevokeService(opts...)
	r.Settings = NewDeviceSettingService(opts...)
	r.Unrevokes = NewDeviceUnrevokeService(opts...)
	r.OverrideCodes = NewDeviceOverrideCodeService(opts...)
	return
}
