// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// DeviceFleetStatusService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceFleetStatusService] method instead.
type DeviceFleetStatusService struct {
	Options []option.RequestOption
}

// NewDeviceFleetStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceFleetStatusService(opts ...option.RequestOption) (r *DeviceFleetStatusService) {
	r = &DeviceFleetStatusService{}
	r.Options = opts
	return
}

// Get the live status of a latest device given device_id from the device_state
// table
func (r *DeviceFleetStatusService) Get(ctx context.Context, deviceID string, params DeviceFleetStatusGetParams, opts ...option.RequestOption) (res *DeviceFleetStatusGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/%s/fleet-status/live", params.AccountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type DeviceFleetStatusGetResponse struct {
	// Cloudflare colo airport code.
	Colo string `json:"colo" api:"required"`
	// Device identifier (UUID v4)
	DeviceID string `json:"deviceId" api:"required"`
	// The mode under which the WARP client is run.
	Mode string `json:"mode" api:"required"`
	// Operating system.
	Platform string `json:"platform" api:"required"`
	// Network status.
	Status    string `json:"status" api:"required"`
	Timestamp string `json:"timestamp" api:"required"`
	// WARP client version.
	Version         string                                    `json:"version" api:"required"`
	AlwaysOn        bool                                      `json:"alwaysOn" api:"nullable"`
	BatteryCharging bool                                      `json:"batteryCharging" api:"nullable"`
	BatteryCycles   int64                                     `json:"batteryCycles" api:"nullable"`
	BatteryPct      float64                                   `json:"batteryPct" api:"nullable"`
	ConnectionType  string                                    `json:"connectionType" api:"nullable"`
	CPUPct          float64                                   `json:"cpuPct" api:"nullable"`
	CPUPctByApp     []DeviceFleetStatusGetResponseCPUPctByApp `json:"cpuPctByApp" api:"nullable"`
	DeviceIPV4      DeviceFleetStatusGetResponseDeviceIPV4    `json:"deviceIpv4" api:"nullable"`
	DeviceIPV6      DeviceFleetStatusGetResponseDeviceIPV6    `json:"deviceIpv6" api:"nullable"`
	// Device identifier (human readable).
	DeviceName string `json:"deviceName"`
	// Deprecated: use registrationId. Device registration identifier (UUID).
	//
	// Deprecated: Use `registrationId` instead.
	DeviceRegistration string                                  `json:"deviceRegistration" api:"nullable"`
	DiskReadBps        int64                                   `json:"diskReadBps" api:"nullable"`
	DiskUsagePct       float64                                 `json:"diskUsagePct" api:"nullable"`
	DiskWriteBps       int64                                   `json:"diskWriteBps" api:"nullable"`
	DOHSubdomain       string                                  `json:"dohSubdomain" api:"nullable"`
	EstimatedLossPct   float64                                 `json:"estimatedLossPct" api:"nullable"`
	FirewallEnabled    bool                                    `json:"firewallEnabled" api:"nullable"`
	GatewayIPV4        DeviceFleetStatusGetResponseGatewayIPV4 `json:"gatewayIpv4" api:"nullable"`
	GatewayIPV6        DeviceFleetStatusGetResponseGatewayIPV6 `json:"gatewayIpv6" api:"nullable"`
	HandshakeLatencyMs float64                                 `json:"handshakeLatencyMs" api:"nullable"`
	ISPIPV4            DeviceFleetStatusGetResponseISPIPV4     `json:"ispIpv4" api:"nullable"`
	ISPIPV6            DeviceFleetStatusGetResponseISPIPV6     `json:"ispIpv6" api:"nullable"`
	Metal              string                                  `json:"metal" api:"nullable"`
	NetworkRcvdBps     int64                                   `json:"networkRcvdBps" api:"nullable"`
	NetworkSentBps     int64                                   `json:"networkSentBps" api:"nullable"`
	NetworkSsid        string                                  `json:"networkSsid" api:"nullable"`
	// User contact email address
	PersonEmail     string                                        `json:"personEmail"`
	RamAvailableKB  int64                                         `json:"ramAvailableKb" api:"nullable"`
	RamUsedPct      float64                                       `json:"ramUsedPct" api:"nullable"`
	RamUsedPctByApp []DeviceFleetStatusGetResponseRamUsedPctByApp `json:"ramUsedPctByApp" api:"nullable"`
	// Device registration identifier (UUID v4). On multi-user devices, this uniquely
	// identifies a user's registration on the device.
	RegistrationID string `json:"registrationId" api:"nullable"`
	// Round-trip time statistics for the WARP tunnel.
	RTT          DeviceFleetStatusGetResponseRTT `json:"rtt" api:"nullable"`
	SwitchLocked bool                            `json:"switchLocked" api:"nullable"`
	// WARP tunnel packet and byte counters.
	TunnelStats     DeviceFleetStatusGetResponseTunnelStats `json:"tunnelStats" api:"nullable"`
	TunnelType      string                                  `json:"tunnelType" api:"nullable"`
	WifiStrengthDbm int64                                   `json:"wifiStrengthDbm" api:"nullable"`
	JSON            deviceFleetStatusGetResponseJSON        `json:"-"`
}

// deviceFleetStatusGetResponseJSON contains the JSON metadata for the struct
// [DeviceFleetStatusGetResponse]
type deviceFleetStatusGetResponseJSON struct {
	Colo               apijson.Field
	DeviceID           apijson.Field
	Mode               apijson.Field
	Platform           apijson.Field
	Status             apijson.Field
	Timestamp          apijson.Field
	Version            apijson.Field
	AlwaysOn           apijson.Field
	BatteryCharging    apijson.Field
	BatteryCycles      apijson.Field
	BatteryPct         apijson.Field
	ConnectionType     apijson.Field
	CPUPct             apijson.Field
	CPUPctByApp        apijson.Field
	DeviceIPV4         apijson.Field
	DeviceIPV6         apijson.Field
	DeviceName         apijson.Field
	DeviceRegistration apijson.Field
	DiskReadBps        apijson.Field
	DiskUsagePct       apijson.Field
	DiskWriteBps       apijson.Field
	DOHSubdomain       apijson.Field
	EstimatedLossPct   apijson.Field
	FirewallEnabled    apijson.Field
	GatewayIPV4        apijson.Field
	GatewayIPV6        apijson.Field
	HandshakeLatencyMs apijson.Field
	ISPIPV4            apijson.Field
	ISPIPV6            apijson.Field
	Metal              apijson.Field
	NetworkRcvdBps     apijson.Field
	NetworkSentBps     apijson.Field
	NetworkSsid        apijson.Field
	PersonEmail        apijson.Field
	RamAvailableKB     apijson.Field
	RamUsedPct         apijson.Field
	RamUsedPctByApp    apijson.Field
	RegistrationID     apijson.Field
	RTT                apijson.Field
	SwitchLocked       apijson.Field
	TunnelStats        apijson.Field
	TunnelType         apijson.Field
	WifiStrengthDbm    apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseCPUPctByApp struct {
	// CPU usage percentage, on a scale of 0 to 100.
	CPUPct float64 `json:"cpu_pct" api:"nullable"`
	// Application name.
	Name string                                      `json:"name" api:"nullable"`
	JSON deviceFleetStatusGetResponseCPUPctByAppJSON `json:"-"`
}

// deviceFleetStatusGetResponseCPUPctByAppJSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseCPUPctByApp]
type deviceFleetStatusGetResponseCPUPctByAppJSON struct {
	CPUPct      apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseCPUPctByApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseCPUPctByAppJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseDeviceIPV4 struct {
	Address  string                                         `json:"address" api:"nullable"`
	ASN      int64                                          `json:"asn" api:"nullable"`
	Aso      string                                         `json:"aso" api:"nullable"`
	Location DeviceFleetStatusGetResponseDeviceIPV4Location `json:"location"`
	Name     string                                         `json:"name" api:"nullable"`
	Netmask  string                                         `json:"netmask" api:"nullable"`
	// IP version (`1` for IPv4, `2` for IPv6, `0` if unknown).
	Version int64                                      `json:"version"`
	JSON    deviceFleetStatusGetResponseDeviceIPV4JSON `json:"-"`
}

// deviceFleetStatusGetResponseDeviceIPV4JSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseDeviceIPV4]
type deviceFleetStatusGetResponseDeviceIPV4JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseDeviceIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseDeviceIPV4JSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseDeviceIPV4Location struct {
	City       string                                             `json:"city" api:"nullable"`
	CountryISO string                                             `json:"country_iso" api:"nullable"`
	StateISO   string                                             `json:"state_iso" api:"nullable"`
	Zip        string                                             `json:"zip" api:"nullable"`
	JSON       deviceFleetStatusGetResponseDeviceIPV4LocationJSON `json:"-"`
}

// deviceFleetStatusGetResponseDeviceIPV4LocationJSON contains the JSON metadata
// for the struct [DeviceFleetStatusGetResponseDeviceIPV4Location]
type deviceFleetStatusGetResponseDeviceIPV4LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseDeviceIPV4Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseDeviceIPV4LocationJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseDeviceIPV6 struct {
	Address  string                                         `json:"address" api:"nullable"`
	ASN      int64                                          `json:"asn" api:"nullable"`
	Aso      string                                         `json:"aso" api:"nullable"`
	Location DeviceFleetStatusGetResponseDeviceIPV6Location `json:"location"`
	Name     string                                         `json:"name" api:"nullable"`
	Netmask  string                                         `json:"netmask" api:"nullable"`
	// IP version (`1` for IPv4, `2` for IPv6, `0` if unknown).
	Version int64                                      `json:"version"`
	JSON    deviceFleetStatusGetResponseDeviceIPV6JSON `json:"-"`
}

// deviceFleetStatusGetResponseDeviceIPV6JSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseDeviceIPV6]
type deviceFleetStatusGetResponseDeviceIPV6JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseDeviceIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseDeviceIPV6JSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseDeviceIPV6Location struct {
	City       string                                             `json:"city" api:"nullable"`
	CountryISO string                                             `json:"country_iso" api:"nullable"`
	StateISO   string                                             `json:"state_iso" api:"nullable"`
	Zip        string                                             `json:"zip" api:"nullable"`
	JSON       deviceFleetStatusGetResponseDeviceIPV6LocationJSON `json:"-"`
}

// deviceFleetStatusGetResponseDeviceIPV6LocationJSON contains the JSON metadata
// for the struct [DeviceFleetStatusGetResponseDeviceIPV6Location]
type deviceFleetStatusGetResponseDeviceIPV6LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseDeviceIPV6Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseDeviceIPV6LocationJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseGatewayIPV4 struct {
	Address  string                                          `json:"address" api:"nullable"`
	ASN      int64                                           `json:"asn" api:"nullable"`
	Aso      string                                          `json:"aso" api:"nullable"`
	Location DeviceFleetStatusGetResponseGatewayIPV4Location `json:"location"`
	Name     string                                          `json:"name" api:"nullable"`
	Netmask  string                                          `json:"netmask" api:"nullable"`
	// IP version (`1` for IPv4, `2` for IPv6, `0` if unknown).
	Version int64                                       `json:"version"`
	JSON    deviceFleetStatusGetResponseGatewayIPV4JSON `json:"-"`
}

// deviceFleetStatusGetResponseGatewayIPV4JSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseGatewayIPV4]
type deviceFleetStatusGetResponseGatewayIPV4JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseGatewayIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseGatewayIPV4JSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseGatewayIPV4Location struct {
	City       string                                              `json:"city" api:"nullable"`
	CountryISO string                                              `json:"country_iso" api:"nullable"`
	StateISO   string                                              `json:"state_iso" api:"nullable"`
	Zip        string                                              `json:"zip" api:"nullable"`
	JSON       deviceFleetStatusGetResponseGatewayIPV4LocationJSON `json:"-"`
}

// deviceFleetStatusGetResponseGatewayIPV4LocationJSON contains the JSON metadata
// for the struct [DeviceFleetStatusGetResponseGatewayIPV4Location]
type deviceFleetStatusGetResponseGatewayIPV4LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseGatewayIPV4Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseGatewayIPV4LocationJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseGatewayIPV6 struct {
	Address  string                                          `json:"address" api:"nullable"`
	ASN      int64                                           `json:"asn" api:"nullable"`
	Aso      string                                          `json:"aso" api:"nullable"`
	Location DeviceFleetStatusGetResponseGatewayIPV6Location `json:"location"`
	Name     string                                          `json:"name" api:"nullable"`
	Netmask  string                                          `json:"netmask" api:"nullable"`
	// IP version (`1` for IPv4, `2` for IPv6, `0` if unknown).
	Version int64                                       `json:"version"`
	JSON    deviceFleetStatusGetResponseGatewayIPV6JSON `json:"-"`
}

// deviceFleetStatusGetResponseGatewayIPV6JSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseGatewayIPV6]
type deviceFleetStatusGetResponseGatewayIPV6JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseGatewayIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseGatewayIPV6JSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseGatewayIPV6Location struct {
	City       string                                              `json:"city" api:"nullable"`
	CountryISO string                                              `json:"country_iso" api:"nullable"`
	StateISO   string                                              `json:"state_iso" api:"nullable"`
	Zip        string                                              `json:"zip" api:"nullable"`
	JSON       deviceFleetStatusGetResponseGatewayIPV6LocationJSON `json:"-"`
}

// deviceFleetStatusGetResponseGatewayIPV6LocationJSON contains the JSON metadata
// for the struct [DeviceFleetStatusGetResponseGatewayIPV6Location]
type deviceFleetStatusGetResponseGatewayIPV6LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseGatewayIPV6Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseGatewayIPV6LocationJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseISPIPV4 struct {
	Address  string                                      `json:"address" api:"nullable"`
	ASN      int64                                       `json:"asn" api:"nullable"`
	Aso      string                                      `json:"aso" api:"nullable"`
	Location DeviceFleetStatusGetResponseISPIPV4Location `json:"location"`
	Name     string                                      `json:"name" api:"nullable"`
	Netmask  string                                      `json:"netmask" api:"nullable"`
	// IP version (`1` for IPv4, `2` for IPv6, `0` if unknown).
	Version int64                                   `json:"version"`
	JSON    deviceFleetStatusGetResponseIspipv4JSON `json:"-"`
}

// deviceFleetStatusGetResponseIspipv4JSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseISPIPV4]
type deviceFleetStatusGetResponseIspipv4JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseISPIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseIspipv4JSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseISPIPV4Location struct {
	City       string                                          `json:"city" api:"nullable"`
	CountryISO string                                          `json:"country_iso" api:"nullable"`
	StateISO   string                                          `json:"state_iso" api:"nullable"`
	Zip        string                                          `json:"zip" api:"nullable"`
	JSON       deviceFleetStatusGetResponseIspipv4LocationJSON `json:"-"`
}

// deviceFleetStatusGetResponseIspipv4LocationJSON contains the JSON metadata for
// the struct [DeviceFleetStatusGetResponseISPIPV4Location]
type deviceFleetStatusGetResponseIspipv4LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseISPIPV4Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseIspipv4LocationJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseISPIPV6 struct {
	Address  string                                      `json:"address" api:"nullable"`
	ASN      int64                                       `json:"asn" api:"nullable"`
	Aso      string                                      `json:"aso" api:"nullable"`
	Location DeviceFleetStatusGetResponseISPIPV6Location `json:"location"`
	Name     string                                      `json:"name" api:"nullable"`
	Netmask  string                                      `json:"netmask" api:"nullable"`
	// IP version (`1` for IPv4, `2` for IPv6, `0` if unknown).
	Version int64                                   `json:"version"`
	JSON    deviceFleetStatusGetResponseIspipv6JSON `json:"-"`
}

// deviceFleetStatusGetResponseIspipv6JSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseISPIPV6]
type deviceFleetStatusGetResponseIspipv6JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseISPIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseIspipv6JSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseISPIPV6Location struct {
	City       string                                          `json:"city" api:"nullable"`
	CountryISO string                                          `json:"country_iso" api:"nullable"`
	StateISO   string                                          `json:"state_iso" api:"nullable"`
	Zip        string                                          `json:"zip" api:"nullable"`
	JSON       deviceFleetStatusGetResponseIspipv6LocationJSON `json:"-"`
}

// deviceFleetStatusGetResponseIspipv6LocationJSON contains the JSON metadata for
// the struct [DeviceFleetStatusGetResponseISPIPV6Location]
type deviceFleetStatusGetResponseIspipv6LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseISPIPV6Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseIspipv6LocationJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseRamUsedPctByApp struct {
	// Application name.
	Name string `json:"name" api:"nullable"`
	// RAM usage percentage, on a scale of 0 to 100.
	RamUsedPct float64                                         `json:"ram_used_pct" api:"nullable"`
	JSON       deviceFleetStatusGetResponseRamUsedPctByAppJSON `json:"-"`
}

// deviceFleetStatusGetResponseRamUsedPctByAppJSON contains the JSON metadata for
// the struct [DeviceFleetStatusGetResponseRamUsedPctByApp]
type deviceFleetStatusGetResponseRamUsedPctByAppJSON struct {
	Name        apijson.Field
	RamUsedPct  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseRamUsedPctByApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseRamUsedPctByAppJSON) RawJSON() string {
	return r.raw
}

// Round-trip time statistics for the WARP tunnel.
type DeviceFleetStatusGetResponseRTT struct {
	// Minimum round-trip time in microseconds.
	MinRTTUs DeviceFleetStatusGetResponseRTTMinRTTUs `json:"minRttUs" api:"nullable"`
	// Round-trip time in microseconds.
	RTTUs DeviceFleetStatusGetResponseRTTRTTUs `json:"rttUs" api:"nullable"`
	// Round-trip time variance in microseconds.
	RTTVarUs DeviceFleetStatusGetResponseRTTRTTVarUs `json:"rttVarUs" api:"nullable"`
	JSON     deviceFleetStatusGetResponseRTTJSON     `json:"-"`
}

// deviceFleetStatusGetResponseRTTJSON contains the JSON metadata for the struct
// [DeviceFleetStatusGetResponseRTT]
type deviceFleetStatusGetResponseRTTJSON struct {
	MinRTTUs    apijson.Field
	RTTUs       apijson.Field
	RTTVarUs    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseRTT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseRTTJSON) RawJSON() string {
	return r.raw
}

// Minimum round-trip time in microseconds.
type DeviceFleetStatusGetResponseRTTMinRTTUs struct {
	Downstream int64                                       `json:"downstream" api:"nullable"`
	Upstream   int64                                       `json:"upstream" api:"nullable"`
	JSON       deviceFleetStatusGetResponseRTTMinRTTUsJSON `json:"-"`
}

// deviceFleetStatusGetResponseRTTMinRTTUsJSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseRTTMinRTTUs]
type deviceFleetStatusGetResponseRTTMinRTTUsJSON struct {
	Downstream  apijson.Field
	Upstream    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseRTTMinRTTUs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseRTTMinRTTUsJSON) RawJSON() string {
	return r.raw
}

// Round-trip time in microseconds.
type DeviceFleetStatusGetResponseRTTRTTUs struct {
	Downstream int64                                    `json:"downstream" api:"nullable"`
	Upstream   int64                                    `json:"upstream" api:"nullable"`
	JSON       deviceFleetStatusGetResponseRttrttUsJSON `json:"-"`
}

// deviceFleetStatusGetResponseRttrttUsJSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseRTTRTTUs]
type deviceFleetStatusGetResponseRttrttUsJSON struct {
	Downstream  apijson.Field
	Upstream    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseRTTRTTUs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseRttrttUsJSON) RawJSON() string {
	return r.raw
}

// Round-trip time variance in microseconds.
type DeviceFleetStatusGetResponseRTTRTTVarUs struct {
	Downstream int64                                       `json:"downstream" api:"nullable"`
	Upstream   int64                                       `json:"upstream" api:"nullable"`
	JSON       deviceFleetStatusGetResponseRttrttVarUsJSON `json:"-"`
}

// deviceFleetStatusGetResponseRttrttVarUsJSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseRTTRTTVarUs]
type deviceFleetStatusGetResponseRttrttVarUsJSON struct {
	Downstream  apijson.Field
	Upstream    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseRTTRTTVarUs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseRttrttVarUsJSON) RawJSON() string {
	return r.raw
}

// WARP tunnel packet and byte counters.
type DeviceFleetStatusGetResponseTunnelStats struct {
	// Number of bytes lost, split by direction.
	BytesLost DeviceFleetStatusGetResponseTunnelStatsBytesLost `json:"bytesLost" api:"nullable"`
	// Number of bytes received, split by direction.
	BytesReceived DeviceFleetStatusGetResponseTunnelStatsBytesReceived `json:"bytesReceived" api:"nullable"`
	// Number of bytes retransmitted, split by direction.
	BytesRetransmitted DeviceFleetStatusGetResponseTunnelStatsBytesRetransmitted `json:"bytesRetransmitted" api:"nullable"`
	// Number of bytes sent, split by direction.
	BytesSent DeviceFleetStatusGetResponseTunnelStatsBytesSent `json:"bytesSent" api:"nullable"`
	// Number of packets lost, split by direction.
	PacketsLost DeviceFleetStatusGetResponseTunnelStatsPacketsLost `json:"packetsLost" api:"nullable"`
	// Number of packets received, split by direction.
	PacketsReceived DeviceFleetStatusGetResponseTunnelStatsPacketsReceived `json:"packetsReceived" api:"nullable"`
	// Number of packets retransmitted, split by direction.
	PacketsRetransmitted DeviceFleetStatusGetResponseTunnelStatsPacketsRetransmitted `json:"packetsRetransmitted" api:"nullable"`
	// Number of packets sent, split by direction.
	PacketsSent DeviceFleetStatusGetResponseTunnelStatsPacketsSent `json:"packetsSent" api:"nullable"`
	// The measurement window duration in milliseconds.
	StatsWindowMs int64                                       `json:"statsWindowMs" api:"nullable"`
	JSON          deviceFleetStatusGetResponseTunnelStatsJSON `json:"-"`
}

// deviceFleetStatusGetResponseTunnelStatsJSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseTunnelStats]
type deviceFleetStatusGetResponseTunnelStatsJSON struct {
	BytesLost            apijson.Field
	BytesReceived        apijson.Field
	BytesRetransmitted   apijson.Field
	BytesSent            apijson.Field
	PacketsLost          apijson.Field
	PacketsReceived      apijson.Field
	PacketsRetransmitted apijson.Field
	PacketsSent          apijson.Field
	StatsWindowMs        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseTunnelStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseTunnelStatsJSON) RawJSON() string {
	return r.raw
}

// Number of bytes lost, split by direction.
type DeviceFleetStatusGetResponseTunnelStatsBytesLost struct {
	Downstream int64                                                `json:"downstream" api:"nullable"`
	Upstream   int64                                                `json:"upstream" api:"nullable"`
	JSON       deviceFleetStatusGetResponseTunnelStatsBytesLostJSON `json:"-"`
}

// deviceFleetStatusGetResponseTunnelStatsBytesLostJSON contains the JSON metadata
// for the struct [DeviceFleetStatusGetResponseTunnelStatsBytesLost]
type deviceFleetStatusGetResponseTunnelStatsBytesLostJSON struct {
	Downstream  apijson.Field
	Upstream    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseTunnelStatsBytesLost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseTunnelStatsBytesLostJSON) RawJSON() string {
	return r.raw
}

// Number of bytes received, split by direction.
type DeviceFleetStatusGetResponseTunnelStatsBytesReceived struct {
	Downstream int64                                                    `json:"downstream" api:"nullable"`
	Upstream   int64                                                    `json:"upstream" api:"nullable"`
	JSON       deviceFleetStatusGetResponseTunnelStatsBytesReceivedJSON `json:"-"`
}

// deviceFleetStatusGetResponseTunnelStatsBytesReceivedJSON contains the JSON
// metadata for the struct [DeviceFleetStatusGetResponseTunnelStatsBytesReceived]
type deviceFleetStatusGetResponseTunnelStatsBytesReceivedJSON struct {
	Downstream  apijson.Field
	Upstream    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseTunnelStatsBytesReceived) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseTunnelStatsBytesReceivedJSON) RawJSON() string {
	return r.raw
}

// Number of bytes retransmitted, split by direction.
type DeviceFleetStatusGetResponseTunnelStatsBytesRetransmitted struct {
	Downstream int64                                                         `json:"downstream" api:"nullable"`
	Upstream   int64                                                         `json:"upstream" api:"nullable"`
	JSON       deviceFleetStatusGetResponseTunnelStatsBytesRetransmittedJSON `json:"-"`
}

// deviceFleetStatusGetResponseTunnelStatsBytesRetransmittedJSON contains the JSON
// metadata for the struct
// [DeviceFleetStatusGetResponseTunnelStatsBytesRetransmitted]
type deviceFleetStatusGetResponseTunnelStatsBytesRetransmittedJSON struct {
	Downstream  apijson.Field
	Upstream    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseTunnelStatsBytesRetransmitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseTunnelStatsBytesRetransmittedJSON) RawJSON() string {
	return r.raw
}

// Number of bytes sent, split by direction.
type DeviceFleetStatusGetResponseTunnelStatsBytesSent struct {
	Downstream int64                                                `json:"downstream" api:"nullable"`
	Upstream   int64                                                `json:"upstream" api:"nullable"`
	JSON       deviceFleetStatusGetResponseTunnelStatsBytesSentJSON `json:"-"`
}

// deviceFleetStatusGetResponseTunnelStatsBytesSentJSON contains the JSON metadata
// for the struct [DeviceFleetStatusGetResponseTunnelStatsBytesSent]
type deviceFleetStatusGetResponseTunnelStatsBytesSentJSON struct {
	Downstream  apijson.Field
	Upstream    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseTunnelStatsBytesSent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseTunnelStatsBytesSentJSON) RawJSON() string {
	return r.raw
}

// Number of packets lost, split by direction.
type DeviceFleetStatusGetResponseTunnelStatsPacketsLost struct {
	Downstream int64                                                  `json:"downstream" api:"nullable"`
	Upstream   int64                                                  `json:"upstream" api:"nullable"`
	JSON       deviceFleetStatusGetResponseTunnelStatsPacketsLostJSON `json:"-"`
}

// deviceFleetStatusGetResponseTunnelStatsPacketsLostJSON contains the JSON
// metadata for the struct [DeviceFleetStatusGetResponseTunnelStatsPacketsLost]
type deviceFleetStatusGetResponseTunnelStatsPacketsLostJSON struct {
	Downstream  apijson.Field
	Upstream    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseTunnelStatsPacketsLost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseTunnelStatsPacketsLostJSON) RawJSON() string {
	return r.raw
}

// Number of packets received, split by direction.
type DeviceFleetStatusGetResponseTunnelStatsPacketsReceived struct {
	Downstream int64                                                      `json:"downstream" api:"nullable"`
	Upstream   int64                                                      `json:"upstream" api:"nullable"`
	JSON       deviceFleetStatusGetResponseTunnelStatsPacketsReceivedJSON `json:"-"`
}

// deviceFleetStatusGetResponseTunnelStatsPacketsReceivedJSON contains the JSON
// metadata for the struct [DeviceFleetStatusGetResponseTunnelStatsPacketsReceived]
type deviceFleetStatusGetResponseTunnelStatsPacketsReceivedJSON struct {
	Downstream  apijson.Field
	Upstream    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseTunnelStatsPacketsReceived) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseTunnelStatsPacketsReceivedJSON) RawJSON() string {
	return r.raw
}

// Number of packets retransmitted, split by direction.
type DeviceFleetStatusGetResponseTunnelStatsPacketsRetransmitted struct {
	Downstream int64                                                           `json:"downstream" api:"nullable"`
	Upstream   int64                                                           `json:"upstream" api:"nullable"`
	JSON       deviceFleetStatusGetResponseTunnelStatsPacketsRetransmittedJSON `json:"-"`
}

// deviceFleetStatusGetResponseTunnelStatsPacketsRetransmittedJSON contains the
// JSON metadata for the struct
// [DeviceFleetStatusGetResponseTunnelStatsPacketsRetransmitted]
type deviceFleetStatusGetResponseTunnelStatsPacketsRetransmittedJSON struct {
	Downstream  apijson.Field
	Upstream    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseTunnelStatsPacketsRetransmitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseTunnelStatsPacketsRetransmittedJSON) RawJSON() string {
	return r.raw
}

// Number of packets sent, split by direction.
type DeviceFleetStatusGetResponseTunnelStatsPacketsSent struct {
	Downstream int64                                                  `json:"downstream" api:"nullable"`
	Upstream   int64                                                  `json:"upstream" api:"nullable"`
	JSON       deviceFleetStatusGetResponseTunnelStatsPacketsSentJSON `json:"-"`
}

// deviceFleetStatusGetResponseTunnelStatsPacketsSentJSON contains the JSON
// metadata for the struct [DeviceFleetStatusGetResponseTunnelStatsPacketsSent]
type deviceFleetStatusGetResponseTunnelStatsPacketsSentJSON struct {
	Downstream  apijson.Field
	Upstream    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseTunnelStatsPacketsSent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseTunnelStatsPacketsSentJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetParams struct {
	// Unique identifier linked to an account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Number of minutes before current time.
	SinceMinutes param.Field[float64] `query:"since_minutes" api:"required"`
	// List of data centers to filter results.
	Colo param.Field[string] `query:"colo"`
	// Current time in ISO format.
	TimeNow param.Field[string] `query:"time_now"`
}

// URLQuery serializes [DeviceFleetStatusGetParams]'s query parameters as
// `url.Values`.
func (r DeviceFleetStatusGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
