// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
	Posture       *DevicePostureService
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
	r.Posture = NewDevicePostureService(opts...)
	r.Revokes = NewDeviceRevokeService(opts...)
	r.Settings = NewDeviceSettingService(opts...)
	r.Unrevokes = NewDeviceUnrevokeService(opts...)
	r.OverrideCodes = NewDeviceOverrideCodeService(opts...)
	return
}

// Fetches a list of enrolled devices.
func (r *DeviceService) List(ctx context.Context, query DeviceListParams, opts ...option.RequestOption) (res *[]TeamsDevicesDevices, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single device.
func (r *DeviceService) Get(ctx context.Context, deviceID string, query DeviceGetParams, opts ...option.RequestOption) (res *DeviceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/%s", query.AccountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TeamsDevicesDevices struct {
	// Device ID.
	ID string `json:"id"`
	// When the device was created.
	Created time.Time `json:"created" format:"date-time"`
	// True if the device was deleted.
	Deleted    bool                          `json:"deleted"`
	DeviceType TeamsDevicesDevicesDeviceType `json:"device_type"`
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
	Updated time.Time               `json:"updated" format:"date-time"`
	User    TeamsDevicesDevicesUser `json:"user"`
	// The WARP client version.
	Version string                  `json:"version"`
	JSON    teamsDevicesDevicesJSON `json:"-"`
}

// teamsDevicesDevicesJSON contains the JSON metadata for the struct
// [TeamsDevicesDevices]
type teamsDevicesDevicesJSON struct {
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

func (r *TeamsDevicesDevices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicesJSON) RawJSON() string {
	return r.raw
}

type TeamsDevicesDevicesDeviceType string

const (
	TeamsDevicesDevicesDeviceTypeWindows TeamsDevicesDevicesDeviceType = "windows"
	TeamsDevicesDevicesDeviceTypeMac     TeamsDevicesDevicesDeviceType = "mac"
	TeamsDevicesDevicesDeviceTypeLinux   TeamsDevicesDevicesDeviceType = "linux"
	TeamsDevicesDevicesDeviceTypeAndroid TeamsDevicesDevicesDeviceType = "android"
	TeamsDevicesDevicesDeviceTypeIos     TeamsDevicesDevicesDeviceType = "ios"
)

type TeamsDevicesDevicesUser struct {
	// UUID
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The enrolled device user's name.
	Name string                      `json:"name"`
	JSON teamsDevicesDevicesUserJSON `json:"-"`
}

// teamsDevicesDevicesUserJSON contains the JSON metadata for the struct
// [TeamsDevicesDevicesUser]
type teamsDevicesDevicesUserJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamsDevicesDevicesUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDevicesUserJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.DeviceGetResponseUnknown] or
// [shared.UnionString].
type DeviceGetResponse interface {
	ImplementsZeroTrustDeviceGetResponse()
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

type DeviceListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type DeviceListResponseEnvelope struct {
	Errors   []DeviceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceListResponseEnvelopeMessages `json:"messages,required"`
	Result   []TeamsDevicesDevices                `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DeviceListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DeviceListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       deviceListResponseEnvelopeJSON       `json:"-"`
}

// deviceListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceListResponseEnvelope]
type deviceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DeviceListResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    deviceListResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DeviceListResponseEnvelopeErrors]
type deviceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DeviceListResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    deviceListResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DeviceListResponseEnvelopeMessages]
type deviceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceListResponseEnvelopeSuccess bool

const (
	DeviceListResponseEnvelopeSuccessTrue DeviceListResponseEnvelopeSuccess = true
)

type DeviceListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       deviceListResponseEnvelopeResultInfoJSON `json:"-"`
}

// deviceListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [DeviceListResponseEnvelopeResultInfo]
type deviceListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DeviceGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
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

func (r deviceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r deviceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r deviceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceGetResponseEnvelopeSuccess bool

const (
	DeviceGetResponseEnvelopeSuccessTrue DeviceGetResponseEnvelopeSuccess = true
)
