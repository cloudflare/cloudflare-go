// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
	Posture       *DevicePostureService
	Revoke        *DeviceRevokeService
	Settings      *DeviceSettingService
	Unrevoke      *DeviceUnrevokeService
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
	r.Posture = NewDevicePostureService(opts...)
	r.Revoke = NewDeviceRevokeService(opts...)
	r.Settings = NewDeviceSettingService(opts...)
	r.Unrevoke = NewDeviceUnrevokeService(opts...)
	r.OverrideCodes = NewDeviceOverrideCodeService(opts...)
	return
}

// Fetches a list of enrolled devices.
func (r *DeviceService) List(ctx context.Context, query DeviceListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ZeroTrustDevices], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/devices", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches a list of enrolled devices.
func (r *DeviceService) ListAutoPaging(ctx context.Context, query DeviceListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ZeroTrustDevices] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Fetches details for a single device.
func (r *DeviceService) Get(ctx context.Context, deviceID string, query DeviceGetParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef173, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/%s", query.AccountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDevices struct {
	// Device ID.
	ID string `json:"id"`
	// When the device was created.
	Created time.Time `json:"created" format:"date-time"`
	// True if the device was deleted.
	Deleted    bool                       `json:"deleted"`
	DeviceType ZeroTrustDevicesDeviceType `json:"device_type"`
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// The device's public key.
	Key string `json:"key"`
	// When the device last connected to Cloudflare services.
	LastSeen time.Time `json:"last_seen" format:"date-time"`
	// The device mac address.
	MacAddress string `json:"mac_address"`
	// The device manufacturer name.
	Manufacturer string `json:"manufacturer"`
	// The device model name.
	Model string `json:"model"`
	// The device name.
	Name string `json:"name"`
	// The Linux distro name.
	OSDistroName string `json:"os_distro_name"`
	// The Linux distro revision.
	OSDistroRevision string `json:"os_distro_revision"`
	// The operating system version.
	OSVersion string `json:"os_version"`
	// The operating system version extra parameter.
	OSVersionExtra string `json:"os_version_extra"`
	// When the device was revoked.
	RevokedAt time.Time `json:"revoked_at" format:"date-time"`
	// The device serial number.
	SerialNumber string `json:"serial_number"`
	// When the device was updated.
	Updated time.Time            `json:"updated" format:"date-time"`
	User    ZeroTrustDevicesUser `json:"user"`
	// The WARP client version.
	Version string               `json:"version"`
	JSON    zeroTrustDevicesJSON `json:"-"`
}

// zeroTrustDevicesJSON contains the JSON metadata for the struct
// [ZeroTrustDevices]
type zeroTrustDevicesJSON struct {
	ID               apijson.Field
	Created          apijson.Field
	Deleted          apijson.Field
	DeviceType       apijson.Field
	IP               apijson.Field
	Key              apijson.Field
	LastSeen         apijson.Field
	MacAddress       apijson.Field
	Manufacturer     apijson.Field
	Model            apijson.Field
	Name             apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersion        apijson.Field
	OSVersionExtra   apijson.Field
	RevokedAt        apijson.Field
	SerialNumber     apijson.Field
	Updated          apijson.Field
	User             apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustDevices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicesJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDevicesDeviceType string

const (
	ZeroTrustDevicesDeviceTypeWindows ZeroTrustDevicesDeviceType = "windows"
	ZeroTrustDevicesDeviceTypeMac     ZeroTrustDevicesDeviceType = "mac"
	ZeroTrustDevicesDeviceTypeLinux   ZeroTrustDevicesDeviceType = "linux"
	ZeroTrustDevicesDeviceTypeAndroid ZeroTrustDevicesDeviceType = "android"
	ZeroTrustDevicesDeviceTypeIos     ZeroTrustDevicesDeviceType = "ios"
)

func (r ZeroTrustDevicesDeviceType) IsKnown() bool {
	switch r {
	case ZeroTrustDevicesDeviceTypeWindows, ZeroTrustDevicesDeviceTypeMac, ZeroTrustDevicesDeviceTypeLinux, ZeroTrustDevicesDeviceTypeAndroid, ZeroTrustDevicesDeviceTypeIos:
		return true
	}
	return false
}

type ZeroTrustDevicesUser struct {
	// UUID
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The enrolled device user's name.
	Name string                   `json:"name"`
	JSON zeroTrustDevicesUserJSON `json:"-"`
}

// zeroTrustDevicesUserJSON contains the JSON metadata for the struct
// [ZeroTrustDevicesUser]
type zeroTrustDevicesUserJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicesUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDevicesUserJSON) RawJSON() string {
	return r.raw
}

type DeviceListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   shared.UnnamedSchemaRef173   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceGetResponseEnvelopeJSON    `json:"-"`
}

// deviceGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceGetResponseEnvelope]
type deviceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceGetResponseEnvelopeSuccess bool

const (
	DeviceGetResponseEnvelopeSuccessTrue DeviceGetResponseEnvelopeSuccess = true
)

func (r DeviceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
