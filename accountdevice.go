// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDeviceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDeviceService] method
// instead.
type AccountDeviceService struct {
	Options       []option.RequestOption
	DexTests      *AccountDeviceDexTestService
	Networks      *AccountDeviceNetworkService
	Policies      *AccountDevicePolicyService
	Postures      *AccountDevicePostureService
	Revokes       *AccountDeviceRevokeService
	Settings      *AccountDeviceSettingService
	Unrevokes     *AccountDeviceUnrevokeService
	OverrideCodes *AccountDeviceOverrideCodeService
}

// NewAccountDeviceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountDeviceService(opts ...option.RequestOption) (r *AccountDeviceService) {
	r = &AccountDeviceService{}
	r.Options = opts
	r.DexTests = NewAccountDeviceDexTestService(opts...)
	r.Networks = NewAccountDeviceNetworkService(opts...)
	r.Policies = NewAccountDevicePolicyService(opts...)
	r.Postures = NewAccountDevicePostureService(opts...)
	r.Revokes = NewAccountDeviceRevokeService(opts...)
	r.Settings = NewAccountDeviceSettingService(opts...)
	r.Unrevokes = NewAccountDeviceUnrevokeService(opts...)
	r.OverrideCodes = NewAccountDeviceOverrideCodeService(opts...)
	return
}

// Fetch a single Device.
func (r *AccountDeviceService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *DeviceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Enrolled Devices.
func (r *AccountDeviceService) DevicesListDevices(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *DevicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DeviceResponse struct {
	Errors   []DeviceResponseError   `json:"errors"`
	Messages []DeviceResponseMessage `json:"messages"`
	Result   interface{}             `json:"result"`
	// Whether the API call was successful
	Success DeviceResponseSuccess `json:"success"`
	JSON    deviceResponseJSON    `json:"-"`
}

// deviceResponseJSON contains the JSON metadata for the struct [DeviceResponse]
type deviceResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceResponseError struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    deviceResponseErrorJSON `json:"-"`
}

// deviceResponseErrorJSON contains the JSON metadata for the struct
// [DeviceResponseError]
type deviceResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceResponseMessage struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    deviceResponseMessageJSON `json:"-"`
}

// deviceResponseMessageJSON contains the JSON metadata for the struct
// [DeviceResponseMessage]
type deviceResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DeviceResponseSuccess bool

const (
	DeviceResponseSuccessTrue DeviceResponseSuccess = true
)

type DevicesResponse struct {
	Errors     []DevicesResponseError    `json:"errors"`
	Messages   []DevicesResponseMessage  `json:"messages"`
	Result     []DevicesResponseResult   `json:"result"`
	ResultInfo DevicesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DevicesResponseSuccess `json:"success"`
	JSON    devicesResponseJSON    `json:"-"`
}

// devicesResponseJSON contains the JSON metadata for the struct [DevicesResponse]
type devicesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicesResponseError struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    devicesResponseErrorJSON `json:"-"`
}

// devicesResponseErrorJSON contains the JSON metadata for the struct
// [DevicesResponseError]
type devicesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicesResponseMessage struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    devicesResponseMessageJSON `json:"-"`
}

// devicesResponseMessageJSON contains the JSON metadata for the struct
// [DevicesResponseMessage]
type devicesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicesResponseResult struct {
	// Device ID.
	ID string `json:"id"`
	// When the device was created.
	Created time.Time `json:"created" format:"date-time"`
	// True if the device was deleted.
	Deleted    bool                            `json:"deleted"`
	DeviceType DevicesResponseResultDeviceType `json:"device_type"`
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
	OsDistroName string `json:"os_distro_name"`
	// The Linux distro revision.
	OsDistroRevision string `json:"os_distro_revision"`
	// The operating system version.
	OsVersion string `json:"os_version"`
	// When the device was revoked.
	RevokedAt time.Time `json:"revoked_at" format:"date-time"`
	// The device serial number.
	SerialNumber string `json:"serial_number"`
	// When the device was updated.
	Updated time.Time                 `json:"updated" format:"date-time"`
	User    DevicesResponseResultUser `json:"user"`
	// The WARP client version.
	Version string                    `json:"version"`
	JSON    devicesResponseResultJSON `json:"-"`
}

// devicesResponseResultJSON contains the JSON metadata for the struct
// [DevicesResponseResult]
type devicesResponseResultJSON struct {
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
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersion        apijson.Field
	RevokedAt        apijson.Field
	SerialNumber     apijson.Field
	Updated          apijson.Field
	User             apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicesResponseResultDeviceType string

const (
	DevicesResponseResultDeviceTypeWindows DevicesResponseResultDeviceType = "windows"
	DevicesResponseResultDeviceTypeMac     DevicesResponseResultDeviceType = "mac"
	DevicesResponseResultDeviceTypeLinux   DevicesResponseResultDeviceType = "linux"
	DevicesResponseResultDeviceTypeAndroid DevicesResponseResultDeviceType = "android"
	DevicesResponseResultDeviceTypeIos     DevicesResponseResultDeviceType = "ios"
)

type DevicesResponseResultUser struct {
	// UUID
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The enrolled device user's name.
	Name string                        `json:"name"`
	JSON devicesResponseResultUserJSON `json:"-"`
}

// devicesResponseResultUserJSON contains the JSON metadata for the struct
// [DevicesResponseResultUser]
type devicesResponseResultUserJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicesResponseResultUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                       `json:"total_count"`
	JSON       devicesResponseResultInfoJSON `json:"-"`
}

// devicesResponseResultInfoJSON contains the JSON metadata for the struct
// [DevicesResponseResultInfo]
type devicesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DevicesResponseSuccess bool

const (
	DevicesResponseSuccessTrue DevicesResponseSuccess = true
)
