// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
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

// Fetches a list of enrolled devices.
func (r *DeviceService) DevicesListDevices(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DeviceDevicesListDevicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceDevicesListDevicesResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single device.
func (r *DeviceService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *DeviceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceDevicesListDevicesResponse struct {
	// Device ID.
	ID string `json:"id"`
	// When the device was created.
	Created time.Time `json:"created" format:"date-time"`
	// True if the device was deleted.
	Deleted    bool                                       `json:"deleted"`
	DeviceType DeviceDevicesListDevicesResponseDeviceType `json:"device_type"`
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
	Updated time.Time                            `json:"updated" format:"date-time"`
	User    DeviceDevicesListDevicesResponseUser `json:"user"`
	// The WARP client version.
	Version string                               `json:"version"`
	JSON    deviceDevicesListDevicesResponseJSON `json:"-"`
}

// deviceDevicesListDevicesResponseJSON contains the JSON metadata for the struct
// [DeviceDevicesListDevicesResponse]
type deviceDevicesListDevicesResponseJSON struct {
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

func (r *DeviceDevicesListDevicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDevicesListDevicesResponseDeviceType string

const (
	DeviceDevicesListDevicesResponseDeviceTypeWindows DeviceDevicesListDevicesResponseDeviceType = "windows"
	DeviceDevicesListDevicesResponseDeviceTypeMac     DeviceDevicesListDevicesResponseDeviceType = "mac"
	DeviceDevicesListDevicesResponseDeviceTypeLinux   DeviceDevicesListDevicesResponseDeviceType = "linux"
	DeviceDevicesListDevicesResponseDeviceTypeAndroid DeviceDevicesListDevicesResponseDeviceType = "android"
	DeviceDevicesListDevicesResponseDeviceTypeIos     DeviceDevicesListDevicesResponseDeviceType = "ios"
)

type DeviceDevicesListDevicesResponseUser struct {
	// UUID
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The enrolled device user's name.
	Name string                                   `json:"name"`
	JSON deviceDevicesListDevicesResponseUserJSON `json:"-"`
}

// deviceDevicesListDevicesResponseUserJSON contains the JSON metadata for the
// struct [DeviceDevicesListDevicesResponseUser]
type deviceDevicesListDevicesResponseUserJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDevicesListDevicesResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [DeviceGetResponseUnknown] or [shared.UnionString].
type DeviceGetResponse interface {
	ImplementsDeviceGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeviceGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DeviceDevicesListDevicesResponseEnvelope struct {
	Errors   []DeviceDevicesListDevicesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDevicesListDevicesResponseEnvelopeMessages `json:"messages,required"`
	Result   []DeviceDevicesListDevicesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DeviceDevicesListDevicesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DeviceDevicesListDevicesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       deviceDevicesListDevicesResponseEnvelopeJSON       `json:"-"`
}

// deviceDevicesListDevicesResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceDevicesListDevicesResponseEnvelope]
type deviceDevicesListDevicesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDevicesListDevicesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDevicesListDevicesResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    deviceDevicesListDevicesResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDevicesListDevicesResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DeviceDevicesListDevicesResponseEnvelopeErrors]
type deviceDevicesListDevicesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDevicesListDevicesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDevicesListDevicesResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    deviceDevicesListDevicesResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDevicesListDevicesResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DeviceDevicesListDevicesResponseEnvelopeMessages]
type deviceDevicesListDevicesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDevicesListDevicesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceDevicesListDevicesResponseEnvelopeSuccess bool

const (
	DeviceDevicesListDevicesResponseEnvelopeSuccessTrue DeviceDevicesListDevicesResponseEnvelopeSuccess = true
)

type DeviceDevicesListDevicesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       deviceDevicesListDevicesResponseEnvelopeResultInfoJSON `json:"-"`
}

// deviceDevicesListDevicesResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DeviceDevicesListDevicesResponseEnvelopeResultInfo]
type deviceDevicesListDevicesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDevicesListDevicesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceGetResponseEnvelope struct {
	Errors   []DeviceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceGetResponse                   `json:"result,required,nullable"`
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

type DeviceGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    deviceGetResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DeviceGetResponseEnvelopeErrors]
type deviceGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceGetResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    deviceGetResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DeviceGetResponseEnvelopeMessages]
type deviceGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceGetResponseEnvelopeSuccess bool

const (
	DeviceGetResponseEnvelopeSuccessTrue DeviceGetResponseEnvelopeSuccess = true
)
