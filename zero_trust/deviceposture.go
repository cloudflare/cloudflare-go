// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// DevicePostureService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDevicePostureService] method
// instead.
type DevicePostureService struct {
	Options      []option.RequestOption
	Integrations *DevicePostureIntegrationService
}

// NewDevicePostureService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDevicePostureService(opts ...option.RequestOption) (r *DevicePostureService) {
	r = &DevicePostureService{}
	r.Options = opts
	r.Integrations = NewDevicePostureIntegrationService(opts...)
	return
}

// Creates a new device posture rule.
func (r *DevicePostureService) New(ctx context.Context, params DevicePostureNewParams, opts ...option.RequestOption) (res *DevicePostureRules, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a device posture rule.
func (r *DevicePostureService) Update(ctx context.Context, ruleID string, params DevicePostureUpdateParams, opts ...option.RequestOption) (res *DevicePostureRules, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches device posture rules for a Zero Trust account.
func (r *DevicePostureService) List(ctx context.Context, query DevicePostureListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DevicePostureRules], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/devices/posture", query.AccountID)
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

// Fetches device posture rules for a Zero Trust account.
func (r *DevicePostureService) ListAutoPaging(ctx context.Context, query DevicePostureListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DevicePostureRules] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a device posture rule.
func (r *DevicePostureService) Delete(ctx context.Context, ruleID string, params DevicePostureDeleteParams, opts ...option.RequestOption) (res *DevicePostureDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single device posture rule.
func (r *DevicePostureService) Get(ctx context.Context, ruleID string, query DevicePostureGetParams, opts ...option.RequestOption) (res *DevicePostureRules, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePostureRules struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input DevicePostureRulesInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []DevicePostureRulesMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type DevicePostureRulesType `json:"type"`
	JSON devicePostureRulesJSON `json:"-"`
}

// devicePostureRulesJSON contains the JSON metadata for the struct
// [DevicePostureRules]
type devicePostureRulesJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expiration  apijson.Field
	Input       apijson.Field
	Match       apijson.Field
	Name        apijson.Field
	Schedule    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesJSON) RawJSON() string {
	return r.raw
}

// The value to be checked against.
//
// Union satisfied by
// [zero_trust.DevicePostureRulesInputTeamsDevicesFileInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesOSVersionInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesFirewallInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesSentineloneInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesCarbonblackInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesApplicationInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesClientCertificateInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesIntuneInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesKolideInputRequest],
// [zero_trust.DevicePostureRulesInputTeamsDevicesTaniumInputRequest] or
// [zero_trust.DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureRulesInput interface {
	implementsZeroTrustDevicePostureRulesInput()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DevicePostureRulesInput)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesFileInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesOSVersionInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesFirewallInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesSentineloneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesCarbonblackInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesApplicationInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesClientCertificateInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesIntuneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesKolideInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesTaniumInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest{}),
		},
	)
}

type DevicePostureRulesInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                  `json:"thumbprint"`
	JSON       devicePostureRulesInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesFileInputRequestJSON contains the JSON
// metadata for the struct [DevicePostureRulesInputTeamsDevicesFileInputRequest]
type devicePostureRulesInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesFileInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operating system
type DevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

func (r DevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystemWindows, DevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystemLinux, DevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// devicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest]
type devicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operating System
type DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

func (r DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid, DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos, DevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                          `json:"domain"`
	JSON   devicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest]
type devicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operating System
type DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

func (r DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                                       `json:"os_version_extra"`
	JSON           devicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureRulesInputTeamsDevicesOSVersionInputRequest]
type devicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operating System
type DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

func (r DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

// Operator
type DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator string

const (
	DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorLess            DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = "<"
	DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorLessOrEquals    DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = "<="
	DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorGreater         DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = ">"
	DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorGreaterOrEquals DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = ">="
	DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorEquals          DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

func (r DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorLess, DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorLessOrEquals, DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorGreater, DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorGreaterOrEquals, DevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorEquals:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureRulesInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// devicePostureRulesInputTeamsDevicesFirewallInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureRulesInputTeamsDevicesFirewallInputRequest]
type devicePostureRulesInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesFirewallInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operating System
type DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

func (r DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystemWindows, DevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                         `json:"thumbprint"`
	JSON       devicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureRulesInputTeamsDevicesSentineloneInputRequest]
type devicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operating system
type DevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

func (r DevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows, DevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux, DevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                         `json:"thumbprint"`
	JSON       devicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureRulesInputTeamsDevicesCarbonblackInputRequest]
type devicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operating system
type DevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

func (r DevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows, DevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux, DevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                              `json:"requireAll"`
	JSON       devicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest]
type devicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

type DevicePostureRulesInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                         `json:"thumbprint"`
	JSON       devicePostureRulesInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesApplicationInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureRulesInputTeamsDevicesApplicationInputRequest]
type devicePostureRulesInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesApplicationInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operating system
type DevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

func (r DevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystemWindows, DevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystemLinux, DevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                               `json:"cn,required"`
	JSON devicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureRulesInputTeamsDevicesClientCertificateInputRequest]
type devicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

type DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                          `json:"connection_id,required"`
	JSON         devicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest]
type devicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Compliance Status
type DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

func (r DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant, DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant, DevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Operator
	Operator DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            devicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// devicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest]
type devicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
	ConnectionID    apijson.Field
	LastSeen        apijson.Field
	Operator        apijson.Field
	OS              apijson.Field
	Overall         apijson.Field
	SensorConfig    apijson.Field
	State           apijson.Field
	Version         apijson.Field
	VersionOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operator
type DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorLess            DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorLessOrEquals    DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorGreater         DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorGreaterOrEquals DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorEquals          DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

func (r DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorLess, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorLessOrEquals, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorGreater, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorGreaterOrEquals, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorEquals:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

func (r DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateOnline, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateOffline, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateUnknown:
		return true
	}
	return false
}

// Version Operator
type DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess            DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals    DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater         DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals          DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

func (r DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals, DevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                    `json:"connection_id,required"`
	JSON         devicePostureRulesInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesIntuneInputRequestJSON contains the JSON
// metadata for the struct [DevicePostureRulesInputTeamsDevicesIntuneInputRequest]
type devicePostureRulesInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesIntuneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Compliance Status
type DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

func (r DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant, DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant, DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown, DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable, DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod, DevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusError:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                    `json:"issue_count,required"`
	JSON       devicePostureRulesInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesKolideInputRequestJSON contains the JSON
// metadata for the struct [DevicePostureRulesInputTeamsDevicesKolideInputRequest]
type devicePostureRulesInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesKolideInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Count Operator
type DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorLess            DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals    DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorGreater         DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorEquals          DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

func (r DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorLess, DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals, DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorGreater, DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals, DevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorEquals:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                   `json:"total_score"`
	JSON       devicePostureRulesInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// devicePostureRulesInputTeamsDevicesTaniumInputRequestJSON contains the JSON
// metadata for the struct [DevicePostureRulesInputTeamsDevicesTaniumInputRequest]
type devicePostureRulesInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesTaniumInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorLess            DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals    DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorGreater         DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorEquals          DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = "=="
)

func (r DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorLess, DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals, DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorGreater, DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals, DevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

func (r DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelLow, DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelMedium, DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelHigh, DevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorLess            DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals    DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorGreater         DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorEquals          DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

func (r DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorLess, DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals, DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorGreater, DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals, DevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorEquals:
		return true
	}
	return false
}

type DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     devicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// devicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest]
type devicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureRulesInput() {
}

// Network status of device.
type DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

func (r DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected, DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected, DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting, DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting:
		return true
	}
	return false
}

// operator
type DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorLess            DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorLessOrEquals    DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorGreater         DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorGreaterOrEquals DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorEquals          DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

func (r DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorLess, DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorLessOrEquals, DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorGreater, DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorGreaterOrEquals, DevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorEquals:
		return true
	}
	return false
}

type DevicePostureRulesMatch struct {
	Platform DevicePostureRulesMatchPlatform `json:"platform"`
	JSON     devicePostureRulesMatchJSON     `json:"-"`
}

// devicePostureRulesMatchJSON contains the JSON metadata for the struct
// [DevicePostureRulesMatch]
type devicePostureRulesMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureRulesMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesMatchJSON) RawJSON() string {
	return r.raw
}

type DevicePostureRulesMatchPlatform string

const (
	DevicePostureRulesMatchPlatformWindows DevicePostureRulesMatchPlatform = "windows"
	DevicePostureRulesMatchPlatformMac     DevicePostureRulesMatchPlatform = "mac"
	DevicePostureRulesMatchPlatformLinux   DevicePostureRulesMatchPlatform = "linux"
	DevicePostureRulesMatchPlatformAndroid DevicePostureRulesMatchPlatform = "android"
	DevicePostureRulesMatchPlatformIos     DevicePostureRulesMatchPlatform = "ios"
)

func (r DevicePostureRulesMatchPlatform) IsKnown() bool {
	switch r {
	case DevicePostureRulesMatchPlatformWindows, DevicePostureRulesMatchPlatformMac, DevicePostureRulesMatchPlatformLinux, DevicePostureRulesMatchPlatformAndroid, DevicePostureRulesMatchPlatformIos:
		return true
	}
	return false
}

// The type of device posture rule.
type DevicePostureRulesType string

const (
	DevicePostureRulesTypeFile              DevicePostureRulesType = "file"
	DevicePostureRulesTypeApplication       DevicePostureRulesType = "application"
	DevicePostureRulesTypeTanium            DevicePostureRulesType = "tanium"
	DevicePostureRulesTypeGateway           DevicePostureRulesType = "gateway"
	DevicePostureRulesTypeWARP              DevicePostureRulesType = "warp"
	DevicePostureRulesTypeDiskEncryption    DevicePostureRulesType = "disk_encryption"
	DevicePostureRulesTypeSentinelone       DevicePostureRulesType = "sentinelone"
	DevicePostureRulesTypeCarbonblack       DevicePostureRulesType = "carbonblack"
	DevicePostureRulesTypeFirewall          DevicePostureRulesType = "firewall"
	DevicePostureRulesTypeOSVersion         DevicePostureRulesType = "os_version"
	DevicePostureRulesTypeDomainJoined      DevicePostureRulesType = "domain_joined"
	DevicePostureRulesTypeClientCertificate DevicePostureRulesType = "client_certificate"
	DevicePostureRulesTypeUniqueClientID    DevicePostureRulesType = "unique_client_id"
	DevicePostureRulesTypeKolide            DevicePostureRulesType = "kolide"
	DevicePostureRulesTypeTaniumS2s         DevicePostureRulesType = "tanium_s2s"
	DevicePostureRulesTypeCrowdstrikeS2s    DevicePostureRulesType = "crowdstrike_s2s"
	DevicePostureRulesTypeIntune            DevicePostureRulesType = "intune"
	DevicePostureRulesTypeWorkspaceOne      DevicePostureRulesType = "workspace_one"
	DevicePostureRulesTypeSentineloneS2s    DevicePostureRulesType = "sentinelone_s2s"
)

func (r DevicePostureRulesType) IsKnown() bool {
	switch r {
	case DevicePostureRulesTypeFile, DevicePostureRulesTypeApplication, DevicePostureRulesTypeTanium, DevicePostureRulesTypeGateway, DevicePostureRulesTypeWARP, DevicePostureRulesTypeDiskEncryption, DevicePostureRulesTypeSentinelone, DevicePostureRulesTypeCarbonblack, DevicePostureRulesTypeFirewall, DevicePostureRulesTypeOSVersion, DevicePostureRulesTypeDomainJoined, DevicePostureRulesTypeClientCertificate, DevicePostureRulesTypeUniqueClientID, DevicePostureRulesTypeKolide, DevicePostureRulesTypeTaniumS2s, DevicePostureRulesTypeCrowdstrikeS2s, DevicePostureRulesTypeIntune, DevicePostureRulesTypeWorkspaceOne, DevicePostureRulesTypeSentineloneS2s:
		return true
	}
	return false
}

type DevicePostureDeleteResponse struct {
	// API UUID.
	ID   string                          `json:"id"`
	JSON devicePostureDeleteResponseJSON `json:"-"`
}

// devicePostureDeleteResponseJSON contains the JSON metadata for the struct
// [DevicePostureDeleteResponse]
type devicePostureDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DevicePostureNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[DevicePostureNewParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[DevicePostureNewParamsInput] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]DevicePostureNewParamsMatch] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r DevicePostureNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type DevicePostureNewParamsType string

const (
	DevicePostureNewParamsTypeFile              DevicePostureNewParamsType = "file"
	DevicePostureNewParamsTypeApplication       DevicePostureNewParamsType = "application"
	DevicePostureNewParamsTypeTanium            DevicePostureNewParamsType = "tanium"
	DevicePostureNewParamsTypeGateway           DevicePostureNewParamsType = "gateway"
	DevicePostureNewParamsTypeWARP              DevicePostureNewParamsType = "warp"
	DevicePostureNewParamsTypeDiskEncryption    DevicePostureNewParamsType = "disk_encryption"
	DevicePostureNewParamsTypeSentinelone       DevicePostureNewParamsType = "sentinelone"
	DevicePostureNewParamsTypeCarbonblack       DevicePostureNewParamsType = "carbonblack"
	DevicePostureNewParamsTypeFirewall          DevicePostureNewParamsType = "firewall"
	DevicePostureNewParamsTypeOSVersion         DevicePostureNewParamsType = "os_version"
	DevicePostureNewParamsTypeDomainJoined      DevicePostureNewParamsType = "domain_joined"
	DevicePostureNewParamsTypeClientCertificate DevicePostureNewParamsType = "client_certificate"
	DevicePostureNewParamsTypeUniqueClientID    DevicePostureNewParamsType = "unique_client_id"
	DevicePostureNewParamsTypeKolide            DevicePostureNewParamsType = "kolide"
	DevicePostureNewParamsTypeTaniumS2s         DevicePostureNewParamsType = "tanium_s2s"
	DevicePostureNewParamsTypeCrowdstrikeS2s    DevicePostureNewParamsType = "crowdstrike_s2s"
	DevicePostureNewParamsTypeIntune            DevicePostureNewParamsType = "intune"
	DevicePostureNewParamsTypeWorkspaceOne      DevicePostureNewParamsType = "workspace_one"
	DevicePostureNewParamsTypeSentineloneS2s    DevicePostureNewParamsType = "sentinelone_s2s"
)

func (r DevicePostureNewParamsType) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsTypeFile, DevicePostureNewParamsTypeApplication, DevicePostureNewParamsTypeTanium, DevicePostureNewParamsTypeGateway, DevicePostureNewParamsTypeWARP, DevicePostureNewParamsTypeDiskEncryption, DevicePostureNewParamsTypeSentinelone, DevicePostureNewParamsTypeCarbonblack, DevicePostureNewParamsTypeFirewall, DevicePostureNewParamsTypeOSVersion, DevicePostureNewParamsTypeDomainJoined, DevicePostureNewParamsTypeClientCertificate, DevicePostureNewParamsTypeUniqueClientID, DevicePostureNewParamsTypeKolide, DevicePostureNewParamsTypeTaniumS2s, DevicePostureNewParamsTypeCrowdstrikeS2s, DevicePostureNewParamsTypeIntune, DevicePostureNewParamsTypeWorkspaceOne, DevicePostureNewParamsTypeSentineloneS2s:
		return true
	}
	return false
}

// The value to be checked against.
//
// Satisfied by
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesFileInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesApplicationInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesKolideInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesTaniumInputRequest],
// [zero_trust.DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureNewParamsInput interface {
	implementsZeroTrustDevicePostureNewParamsInput()
}

type DevicePostureNewParamsInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureNewParamsInputTeamsDevicesFileInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating system
type DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

func (r DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemWindows, DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux, DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

func (r DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid, DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos, DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

func (r DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// Operator
	Operator param.Field[DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator] `json:"operator,required"`
	// Version of OS
	Version param.Field[string] `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra param.Field[string] `json:"os_version_extra"`
}

func (r DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

func (r DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

// Operator
type DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorLess            DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorLessOrEquals    DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorGreater         DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorGreaterOrEquals DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorEquals          DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

func (r DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorLess, DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorLessOrEquals, DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorGreater, DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorGreaterOrEquals, DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorEquals:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

func (r DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows, DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating system
type DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows, DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux, DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating system
type DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

func (r DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows, DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux, DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

type DevicePostureNewParamsInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureNewParamsInputTeamsDevicesApplicationInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating system
type DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

func (r DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemWindows, DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemLinux, DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

type DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Compliance Status
type DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

func (r DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant, DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant, DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// Operator
	Operator param.Field[DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator] `json:"operator"`
	// Os Version
	OS param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState] `json:"state"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator] `json:"versionOperator"`
}

func (r DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operator
type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorLess            DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorLessOrEquals    DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorGreater         DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorGreaterOrEquals DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorEquals          DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

func (r DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorLess, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorLessOrEquals, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorGreater, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorGreaterOrEquals, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorEquals:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

func (r DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown:
		return true
	}
	return false
}

// Version Operator
type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess            DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals    DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater         DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals          DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

func (r DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals, DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Compliance Status
type DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

func (r DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant, DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant, DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown, DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable, DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod, DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusError:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesKolideInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Count Operator
type DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorLess            DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals    DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorGreater         DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorEquals          DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

func (r DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorLess, DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals, DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorGreater, DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals, DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorEquals:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r DevicePostureNewParamsInputTeamsDevicesTaniumInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorLess            DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals    DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorGreater         DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorEquals          DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
)

func (r DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorLess, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorGreater, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

func (r DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLess            DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals    DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreater         DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorEquals          DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

func (r DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLess, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreater, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals, DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorEquals:
		return true
	}
	return false
}

type DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus] `json:"network_status"`
	// operator
	Operator param.Field[DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator] `json:"operator"`
}

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Network status of device.
type DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected, DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected, DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting, DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting:
		return true
	}
	return false
}

// operator
type DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorLess            DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorLessOrEquals    DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorGreater         DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorGreaterOrEquals DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorEquals          DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorLess, DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorLessOrEquals, DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorGreater, DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorGreaterOrEquals, DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorEquals:
		return true
	}
	return false
}

type DevicePostureNewParamsMatch struct {
	Platform param.Field[DevicePostureNewParamsMatchPlatform] `json:"platform"`
}

func (r DevicePostureNewParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePostureNewParamsMatchPlatform string

const (
	DevicePostureNewParamsMatchPlatformWindows DevicePostureNewParamsMatchPlatform = "windows"
	DevicePostureNewParamsMatchPlatformMac     DevicePostureNewParamsMatchPlatform = "mac"
	DevicePostureNewParamsMatchPlatformLinux   DevicePostureNewParamsMatchPlatform = "linux"
	DevicePostureNewParamsMatchPlatformAndroid DevicePostureNewParamsMatchPlatform = "android"
	DevicePostureNewParamsMatchPlatformIos     DevicePostureNewParamsMatchPlatform = "ios"
)

func (r DevicePostureNewParamsMatchPlatform) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsMatchPlatformWindows, DevicePostureNewParamsMatchPlatformMac, DevicePostureNewParamsMatchPlatformLinux, DevicePostureNewParamsMatchPlatformAndroid, DevicePostureNewParamsMatchPlatformIos:
		return true
	}
	return false
}

type DevicePostureNewResponseEnvelope struct {
	Errors   []DevicePostureNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureRules                         `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureNewResponseEnvelopeJSON    `json:"-"`
}

// devicePostureNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePostureNewResponseEnvelope]
type devicePostureNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePostureNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    devicePostureNewResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePostureNewResponseEnvelopeErrors]
type devicePostureNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePostureNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    devicePostureNewResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DevicePostureNewResponseEnvelopeMessages]
type devicePostureNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureNewResponseEnvelopeSuccess bool

const (
	DevicePostureNewResponseEnvelopeSuccessTrue DevicePostureNewResponseEnvelopeSuccess = true
)

func (r DevicePostureNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[DevicePostureUpdateParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[DevicePostureUpdateParamsInput] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]DevicePostureUpdateParamsMatch] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r DevicePostureUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type DevicePostureUpdateParamsType string

const (
	DevicePostureUpdateParamsTypeFile              DevicePostureUpdateParamsType = "file"
	DevicePostureUpdateParamsTypeApplication       DevicePostureUpdateParamsType = "application"
	DevicePostureUpdateParamsTypeTanium            DevicePostureUpdateParamsType = "tanium"
	DevicePostureUpdateParamsTypeGateway           DevicePostureUpdateParamsType = "gateway"
	DevicePostureUpdateParamsTypeWARP              DevicePostureUpdateParamsType = "warp"
	DevicePostureUpdateParamsTypeDiskEncryption    DevicePostureUpdateParamsType = "disk_encryption"
	DevicePostureUpdateParamsTypeSentinelone       DevicePostureUpdateParamsType = "sentinelone"
	DevicePostureUpdateParamsTypeCarbonblack       DevicePostureUpdateParamsType = "carbonblack"
	DevicePostureUpdateParamsTypeFirewall          DevicePostureUpdateParamsType = "firewall"
	DevicePostureUpdateParamsTypeOSVersion         DevicePostureUpdateParamsType = "os_version"
	DevicePostureUpdateParamsTypeDomainJoined      DevicePostureUpdateParamsType = "domain_joined"
	DevicePostureUpdateParamsTypeClientCertificate DevicePostureUpdateParamsType = "client_certificate"
	DevicePostureUpdateParamsTypeUniqueClientID    DevicePostureUpdateParamsType = "unique_client_id"
	DevicePostureUpdateParamsTypeKolide            DevicePostureUpdateParamsType = "kolide"
	DevicePostureUpdateParamsTypeTaniumS2s         DevicePostureUpdateParamsType = "tanium_s2s"
	DevicePostureUpdateParamsTypeCrowdstrikeS2s    DevicePostureUpdateParamsType = "crowdstrike_s2s"
	DevicePostureUpdateParamsTypeIntune            DevicePostureUpdateParamsType = "intune"
	DevicePostureUpdateParamsTypeWorkspaceOne      DevicePostureUpdateParamsType = "workspace_one"
	DevicePostureUpdateParamsTypeSentineloneS2s    DevicePostureUpdateParamsType = "sentinelone_s2s"
)

func (r DevicePostureUpdateParamsType) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsTypeFile, DevicePostureUpdateParamsTypeApplication, DevicePostureUpdateParamsTypeTanium, DevicePostureUpdateParamsTypeGateway, DevicePostureUpdateParamsTypeWARP, DevicePostureUpdateParamsTypeDiskEncryption, DevicePostureUpdateParamsTypeSentinelone, DevicePostureUpdateParamsTypeCarbonblack, DevicePostureUpdateParamsTypeFirewall, DevicePostureUpdateParamsTypeOSVersion, DevicePostureUpdateParamsTypeDomainJoined, DevicePostureUpdateParamsTypeClientCertificate, DevicePostureUpdateParamsTypeUniqueClientID, DevicePostureUpdateParamsTypeKolide, DevicePostureUpdateParamsTypeTaniumS2s, DevicePostureUpdateParamsTypeCrowdstrikeS2s, DevicePostureUpdateParamsTypeIntune, DevicePostureUpdateParamsTypeWorkspaceOne, DevicePostureUpdateParamsTypeSentineloneS2s:
		return true
	}
	return false
}

// The value to be checked against.
//
// Satisfied by
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest],
// [zero_trust.DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureUpdateParamsInput interface {
	implementsZeroTrustDevicePostureUpdateParamsInput()
}

type DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating system
type DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemWindows, DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux, DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid, DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos, DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// Operator
	Operator param.Field[DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator] `json:"operator,required"`
	// Version of OS
	Version param.Field[string] `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra param.Field[string] `json:"os_version_extra"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

// Operator
type DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorLess            DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorLessOrEquals    DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorGreater         DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorGreaterOrEquals DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorEquals          DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

func (r DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorLess, DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorLessOrEquals, DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorGreater, DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorGreaterOrEquals, DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorEquals:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows, DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating system
type DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows, DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux, DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating system
type DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows, DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux, DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

type DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating system
type DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemWindows, DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemLinux, DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

type DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Compliance Status
type DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant, DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant, DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// Operator
	Operator param.Field[DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator] `json:"operator"`
	// Os Version
	OS param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState] `json:"state"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator] `json:"versionOperator"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operator
type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorLess            DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorLessOrEquals    DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorGreater         DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorGreaterOrEquals DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorEquals          DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

func (r DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorLess, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorLessOrEquals, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorGreater, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorGreaterOrEquals, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorEquals:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown:
		return true
	}
	return false
}

// Version Operator
type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess            DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals    DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater         DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals          DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

func (r DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals, DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Compliance Status
type DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant, DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant, DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown, DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable, DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod, DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusError:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Count Operator
type DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorLess            DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals    DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorGreater         DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorEquals          DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

func (r DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorLess, DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals, DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorGreater, DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals, DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorEquals:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorLess            DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals    DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorGreater         DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorEquals          DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
)

func (r DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorLess, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorGreater, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLess            DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals    DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreater         DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorEquals          DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

func (r DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLess, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreater, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals, DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorEquals:
		return true
	}
	return false
}

type DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus] `json:"network_status"`
	// operator
	Operator param.Field[DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator] `json:"operator"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Network status of device.
type DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected, DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected, DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting, DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting:
		return true
	}
	return false
}

// operator
type DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorLess            DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorLessOrEquals    DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorGreater         DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorGreaterOrEquals DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorEquals          DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorLess, DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorLessOrEquals, DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorGreater, DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorGreaterOrEquals, DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorEquals:
		return true
	}
	return false
}

type DevicePostureUpdateParamsMatch struct {
	Platform param.Field[DevicePostureUpdateParamsMatchPlatform] `json:"platform"`
}

func (r DevicePostureUpdateParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePostureUpdateParamsMatchPlatform string

const (
	DevicePostureUpdateParamsMatchPlatformWindows DevicePostureUpdateParamsMatchPlatform = "windows"
	DevicePostureUpdateParamsMatchPlatformMac     DevicePostureUpdateParamsMatchPlatform = "mac"
	DevicePostureUpdateParamsMatchPlatformLinux   DevicePostureUpdateParamsMatchPlatform = "linux"
	DevicePostureUpdateParamsMatchPlatformAndroid DevicePostureUpdateParamsMatchPlatform = "android"
	DevicePostureUpdateParamsMatchPlatformIos     DevicePostureUpdateParamsMatchPlatform = "ios"
)

func (r DevicePostureUpdateParamsMatchPlatform) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsMatchPlatformWindows, DevicePostureUpdateParamsMatchPlatformMac, DevicePostureUpdateParamsMatchPlatformLinux, DevicePostureUpdateParamsMatchPlatformAndroid, DevicePostureUpdateParamsMatchPlatformIos:
		return true
	}
	return false
}

type DevicePostureUpdateResponseEnvelope struct {
	Errors   []DevicePostureUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureRules                            `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureUpdateResponseEnvelopeJSON    `json:"-"`
}

// devicePostureUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePostureUpdateResponseEnvelope]
type devicePostureUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePostureUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    devicePostureUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePostureUpdateResponseEnvelopeErrors]
type devicePostureUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePostureUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    devicePostureUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DevicePostureUpdateResponseEnvelopeMessages]
type devicePostureUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureUpdateResponseEnvelopeSuccess bool

const (
	DevicePostureUpdateResponseEnvelopeSuccessTrue DevicePostureUpdateResponseEnvelopeSuccess = true
)

func (r DevicePostureUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureDeleteParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r DevicePostureDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePostureDeleteResponseEnvelope struct {
	Errors   []DevicePostureDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureDeleteResponseEnvelopeJSON    `json:"-"`
}

// devicePostureDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePostureDeleteResponseEnvelope]
type devicePostureDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePostureDeleteResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    devicePostureDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePostureDeleteResponseEnvelopeErrors]
type devicePostureDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePostureDeleteResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    devicePostureDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DevicePostureDeleteResponseEnvelopeMessages]
type devicePostureDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureDeleteResponseEnvelopeSuccess bool

const (
	DevicePostureDeleteResponseEnvelopeSuccessTrue DevicePostureDeleteResponseEnvelopeSuccess = true
)

func (r DevicePostureDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureGetResponseEnvelope struct {
	Errors   []DevicePostureGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureRules                         `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureGetResponseEnvelopeJSON    `json:"-"`
}

// devicePostureGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePostureGetResponseEnvelope]
type devicePostureGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePostureGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    devicePostureGetResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePostureGetResponseEnvelopeErrors]
type devicePostureGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePostureGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    devicePostureGetResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DevicePostureGetResponseEnvelopeMessages]
type devicePostureGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureGetResponseEnvelopeSuccess bool

const (
	DevicePostureGetResponseEnvelopeSuccessTrue DevicePostureGetResponseEnvelopeSuccess = true
)

func (r DevicePostureGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
