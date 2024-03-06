// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZeroTrustDeviceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDeviceService] method
// instead.
type ZeroTrustDeviceService struct {
	Options       []option.RequestOption
	DEXTests      *ZeroTrustDeviceDEXTestService
	Networks      *ZeroTrustDeviceNetworkService
	Policies      *ZeroTrustDevicePolicyService
	Postures      *ZeroTrustDevicePostureService
	Revokes       *ZeroTrustDeviceRevokeService
	Settings      *ZeroTrustDeviceSettingService
	Unrevokes     *ZeroTrustDeviceUnrevokeService
	OverrideCodes *ZeroTrustDeviceOverrideCodeService
}

// NewZeroTrustDeviceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZeroTrustDeviceService(opts ...option.RequestOption) (r *ZeroTrustDeviceService) {
	r = &ZeroTrustDeviceService{}
	r.Options = opts
	r.DEXTests = NewZeroTrustDeviceDEXTestService(opts...)
	r.Networks = NewZeroTrustDeviceNetworkService(opts...)
	r.Policies = NewZeroTrustDevicePolicyService(opts...)
	r.Postures = NewZeroTrustDevicePostureService(opts...)
	r.Revokes = NewZeroTrustDeviceRevokeService(opts...)
	r.Settings = NewZeroTrustDeviceSettingService(opts...)
	r.Unrevokes = NewZeroTrustDeviceUnrevokeService(opts...)
	r.OverrideCodes = NewZeroTrustDeviceOverrideCodeService(opts...)
	return
}

// Fetches a list of enrolled devices.
func (r *ZeroTrustDeviceService) List(ctx context.Context, query ZeroTrustDeviceListParams, opts ...option.RequestOption) (res *[]ZeroTrustDeviceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single device.
func (r *ZeroTrustDeviceService) Get(ctx context.Context, deviceID string, query ZeroTrustDeviceGetParams, opts ...option.RequestOption) (res *ZeroTrustDeviceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/%s", query.AccountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDeviceListResponse struct {
	// Device ID.
	ID string `json:"id"`
	// When the device was created.
	Created time.Time `json:"created" format:"date-time"`
	// True if the device was deleted.
	Deleted    bool                                  `json:"deleted"`
	DeviceType ZeroTrustDeviceListResponseDeviceType `json:"device_type"`
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
	Updated time.Time                       `json:"updated" format:"date-time"`
	User    ZeroTrustDeviceListResponseUser `json:"user"`
	// The WARP client version.
	Version string                          `json:"version"`
	JSON    zeroTrustDeviceListResponseJSON `json:"-"`
}

// zeroTrustDeviceListResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDeviceListResponse]
type zeroTrustDeviceListResponseJSON struct {
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

func (r *ZeroTrustDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceListResponseDeviceType string

const (
	ZeroTrustDeviceListResponseDeviceTypeWindows ZeroTrustDeviceListResponseDeviceType = "windows"
	ZeroTrustDeviceListResponseDeviceTypeMac     ZeroTrustDeviceListResponseDeviceType = "mac"
	ZeroTrustDeviceListResponseDeviceTypeLinux   ZeroTrustDeviceListResponseDeviceType = "linux"
	ZeroTrustDeviceListResponseDeviceTypeAndroid ZeroTrustDeviceListResponseDeviceType = "android"
	ZeroTrustDeviceListResponseDeviceTypeIos     ZeroTrustDeviceListResponseDeviceType = "ios"
)

type ZeroTrustDeviceListResponseUser struct {
	// UUID
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The enrolled device user's name.
	Name string                              `json:"name"`
	JSON zeroTrustDeviceListResponseUserJSON `json:"-"`
}

// zeroTrustDeviceListResponseUserJSON contains the JSON metadata for the struct
// [ZeroTrustDeviceListResponseUser]
type zeroTrustDeviceListResponseUserJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceListResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceListResponseUserJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ZeroTrustDeviceGetResponseUnknown] or [shared.UnionString].
type ZeroTrustDeviceGetResponse interface {
	ImplementsZeroTrustDeviceGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustDeviceGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustDeviceListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDeviceListResponseEnvelope struct {
	Errors   []ZeroTrustDeviceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDeviceListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDeviceListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDeviceListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDeviceListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDeviceListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceListResponseEnvelope]
type zeroTrustDeviceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zeroTrustDeviceListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceListResponseEnvelopeErrors]
type zeroTrustDeviceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zeroTrustDeviceListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceListResponseEnvelopeMessages]
type zeroTrustDeviceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceListResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceListResponseEnvelopeSuccessTrue ZeroTrustDeviceListResponseEnvelopeSuccess = true
)

type ZeroTrustDeviceListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       zeroTrustDeviceListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDeviceListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceListResponseEnvelopeResultInfo]
type zeroTrustDeviceListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDeviceGetResponseEnvelope struct {
	Errors   []ZeroTrustDeviceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDeviceGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDeviceGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDeviceGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDeviceGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZeroTrustDeviceGetResponseEnvelope]
type zeroTrustDeviceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zeroTrustDeviceGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceGetResponseEnvelopeErrors]
type zeroTrustDeviceGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zeroTrustDeviceGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceGetResponseEnvelopeMessages]
type zeroTrustDeviceGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceGetResponseEnvelopeSuccessTrue ZeroTrustDeviceGetResponseEnvelopeSuccess = true
)
