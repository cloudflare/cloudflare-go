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

// Fetches details for a single device.
func (r *AccountDeviceService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDeviceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a list of enrolled devices.
func (r *AccountDeviceService) DevicesListDevices(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountDeviceDevicesListDevicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountDeviceGetResponse struct {
	Errors   []AccountDeviceGetResponseError   `json:"errors"`
	Messages []AccountDeviceGetResponseMessage `json:"messages"`
	Result   interface{}                       `json:"result"`
	// Whether the API call was successful.
	Success AccountDeviceGetResponseSuccess `json:"success"`
	JSON    accountDeviceGetResponseJSON    `json:"-"`
}

// accountDeviceGetResponseJSON contains the JSON metadata for the struct
// [AccountDeviceGetResponse]
type accountDeviceGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceGetResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    accountDeviceGetResponseErrorJSON `json:"-"`
}

// accountDeviceGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountDeviceGetResponseError]
type accountDeviceGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceGetResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    accountDeviceGetResponseMessageJSON `json:"-"`
}

// accountDeviceGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountDeviceGetResponseMessage]
type accountDeviceGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDeviceGetResponseSuccess bool

const (
	AccountDeviceGetResponseSuccessTrue AccountDeviceGetResponseSuccess = true
)

type AccountDeviceDevicesListDevicesResponse struct {
	Errors     []AccountDeviceDevicesListDevicesResponseError    `json:"errors"`
	Messages   []AccountDeviceDevicesListDevicesResponseMessage  `json:"messages"`
	Result     []AccountDeviceDevicesListDevicesResponseResult   `json:"result"`
	ResultInfo AccountDeviceDevicesListDevicesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDeviceDevicesListDevicesResponseSuccess `json:"success"`
	JSON    accountDeviceDevicesListDevicesResponseJSON    `json:"-"`
}

// accountDeviceDevicesListDevicesResponseJSON contains the JSON metadata for the
// struct [AccountDeviceDevicesListDevicesResponse]
type accountDeviceDevicesListDevicesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDevicesListDevicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDevicesListDevicesResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountDeviceDevicesListDevicesResponseErrorJSON `json:"-"`
}

// accountDeviceDevicesListDevicesResponseErrorJSON contains the JSON metadata for
// the struct [AccountDeviceDevicesListDevicesResponseError]
type accountDeviceDevicesListDevicesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDevicesListDevicesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDevicesListDevicesResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accountDeviceDevicesListDevicesResponseMessageJSON `json:"-"`
}

// accountDeviceDevicesListDevicesResponseMessageJSON contains the JSON metadata
// for the struct [AccountDeviceDevicesListDevicesResponseMessage]
type accountDeviceDevicesListDevicesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDevicesListDevicesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDevicesListDevicesResponseResult struct {
	// Device ID.
	ID string `json:"id"`
	// When the device was created.
	Created time.Time `json:"created" format:"date-time"`
	// True if the device was deleted.
	Deleted    bool                                                    `json:"deleted"`
	DeviceType AccountDeviceDevicesListDevicesResponseResultDeviceType `json:"device_type"`
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
	// The operating system version extra parameter.
	OsVersionExtra string `json:"os_version_extra"`
	// When the device was revoked.
	RevokedAt time.Time `json:"revoked_at" format:"date-time"`
	// The device serial number.
	SerialNumber string `json:"serial_number"`
	// When the device was updated.
	Updated time.Time                                         `json:"updated" format:"date-time"`
	User    AccountDeviceDevicesListDevicesResponseResultUser `json:"user"`
	// The WARP client version.
	Version string                                            `json:"version"`
	JSON    accountDeviceDevicesListDevicesResponseResultJSON `json:"-"`
}

// accountDeviceDevicesListDevicesResponseResultJSON contains the JSON metadata for
// the struct [AccountDeviceDevicesListDevicesResponseResult]
type accountDeviceDevicesListDevicesResponseResultJSON struct {
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
	OsVersionExtra   apijson.Field
	RevokedAt        apijson.Field
	SerialNumber     apijson.Field
	Updated          apijson.Field
	User             apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDeviceDevicesListDevicesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDevicesListDevicesResponseResultDeviceType string

const (
	AccountDeviceDevicesListDevicesResponseResultDeviceTypeWindows AccountDeviceDevicesListDevicesResponseResultDeviceType = "windows"
	AccountDeviceDevicesListDevicesResponseResultDeviceTypeMac     AccountDeviceDevicesListDevicesResponseResultDeviceType = "mac"
	AccountDeviceDevicesListDevicesResponseResultDeviceTypeLinux   AccountDeviceDevicesListDevicesResponseResultDeviceType = "linux"
	AccountDeviceDevicesListDevicesResponseResultDeviceTypeAndroid AccountDeviceDevicesListDevicesResponseResultDeviceType = "android"
	AccountDeviceDevicesListDevicesResponseResultDeviceTypeIos     AccountDeviceDevicesListDevicesResponseResultDeviceType = "ios"
)

type AccountDeviceDevicesListDevicesResponseResultUser struct {
	// UUID
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The enrolled device user's name.
	Name string                                                `json:"name"`
	JSON accountDeviceDevicesListDevicesResponseResultUserJSON `json:"-"`
}

// accountDeviceDevicesListDevicesResponseResultUserJSON contains the JSON metadata
// for the struct [AccountDeviceDevicesListDevicesResponseResultUser]
type accountDeviceDevicesListDevicesResponseResultUserJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDevicesListDevicesResponseResultUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceDevicesListDevicesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       accountDeviceDevicesListDevicesResponseResultInfoJSON `json:"-"`
}

// accountDeviceDevicesListDevicesResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountDeviceDevicesListDevicesResponseResultInfo]
type accountDeviceDevicesListDevicesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDevicesListDevicesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDeviceDevicesListDevicesResponseSuccess bool

const (
	AccountDeviceDevicesListDevicesResponseSuccessTrue AccountDeviceDevicesListDevicesResponseSuccess = true
)
