// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
func (r *DevicePostureService) New(ctx context.Context, accountID interface{}, body DevicePostureNewParams, opts ...option.RequestOption) (res *DevicePostureNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a device posture rule.
func (r *DevicePostureService) Update(ctx context.Context, accountID interface{}, ruleID string, body DevicePostureUpdateParams, opts ...option.RequestOption) (res *DevicePostureUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches device posture rules for a Zero Trust account.
func (r *DevicePostureService) List(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *[]DevicePostureListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a device posture rule.
func (r *DevicePostureService) Delete(ctx context.Context, accountID interface{}, ruleID string, opts ...option.RequestOption) (res *DevicePostureDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single device posture rule.
func (r *DevicePostureService) Get(ctx context.Context, accountID interface{}, ruleID string, opts ...option.RequestOption) (res *DevicePostureGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePostureNewResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input DevicePostureNewResponseInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []DevicePostureNewResponseMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type DevicePostureNewResponseType `json:"type"`
	JSON devicePostureNewResponseJSON `json:"-"`
}

// devicePostureNewResponseJSON contains the JSON metadata for the struct
// [DevicePostureNewResponse]
type devicePostureNewResponseJSON struct {
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

func (r *DevicePostureNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by [DevicePostureNewResponseInputTeamsDevicesFileInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesFirewallInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesSentineloneInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesApplicationInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesClientCertificateInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesIntuneInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesKolideInputRequest],
// [DevicePostureNewResponseInputTeamsDevicesTaniumInputRequest] or
// [DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureNewResponseInput interface {
	implementsDevicePostureNewResponseInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DevicePostureNewResponseInput)(nil)).Elem(), "")
}

type DevicePostureNewResponseInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                        `json:"thumbprint"`
	JSON       devicePostureNewResponseInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesFileInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesFileInputRequest]
type devicePostureNewResponseInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesFileInputRequest) implementsDevicePostureNewResponseInput() {
}

// Operating system
type DevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type DevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem DevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequest]
type devicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequest) implementsDevicePostureNewResponseInput() {
}

// Operating System
type DevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type DevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                `json:"domain"`
	JSON   devicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequest]
type devicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequest) implementsDevicePostureNewResponseInput() {
}

// Operating System
type DevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                                             `json:"os_version_extra"`
	JSON           devicePostureNewResponseInputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesOSVersionInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequest]
type devicePostureNewResponseInputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequest) implementsDevicePostureNewResponseInput() {
}

// Operating System
type DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator string

const (
	DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown5 DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator = "<"
	DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown6 DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator = "<="
	DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown7 DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator = ">"
	DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown8 DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator = ">="
	DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown9 DevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type DevicePostureNewResponseInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem DevicePostureNewResponseInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureNewResponseInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesFirewallInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesFirewallInputRequest]
type devicePostureNewResponseInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesFirewallInputRequest) implementsDevicePostureNewResponseInput() {
}

// Operating System
type DevicePostureNewResponseInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureNewResponseInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureNewResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureNewResponseInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureNewResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type DevicePostureNewResponseInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                               `json:"thumbprint"`
	JSON       devicePostureNewResponseInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesSentineloneInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesSentineloneInputRequest]
type devicePostureNewResponseInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesSentineloneInputRequest) implementsDevicePostureNewResponseInput() {
}

// Operating system
type DevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type DevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                               `json:"thumbprint"`
	JSON       devicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequest]
type devicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequest) implementsDevicePostureNewResponseInput() {
}

// Operating system
type DevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type DevicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                    `json:"requireAll"`
	JSON       devicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequest]
type devicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequest) implementsDevicePostureNewResponseInput() {
}

type DevicePostureNewResponseInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                               `json:"thumbprint"`
	JSON       devicePostureNewResponseInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesApplicationInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesApplicationInputRequest]
type devicePostureNewResponseInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesApplicationInputRequest) implementsDevicePostureNewResponseInput() {
}

// Operating system
type DevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type DevicePostureNewResponseInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                     `json:"cn,required"`
	JSON devicePostureNewResponseInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesClientCertificateInputRequest]
type devicePostureNewResponseInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesClientCertificateInputRequest) implementsDevicePostureNewResponseInput() {
}

type DevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                `json:"connection_id,required"`
	JSON         devicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequest]
type devicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequest) implementsDevicePostureNewResponseInput() {
}

// Compliance Status
type DevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Operator
	Operator DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            devicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequest]
type devicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
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

func (r *DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequest) implementsDevicePostureNewResponseInput() {
}

// Operator
type DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown15 DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown16 DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown17 DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown18 DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown19 DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown25 DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown26 DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown27 DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown28 DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown29 DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type DevicePostureNewResponseInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                          `json:"connection_id,required"`
	JSON         devicePostureNewResponseInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesIntuneInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesIntuneInputRequest]
type devicePostureNewResponseInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesIntuneInputRequest) implementsDevicePostureNewResponseInput() {
}

// Compliance Status
type DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type DevicePostureNewResponseInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator DevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                          `json:"issue_count,required"`
	JSON       devicePostureNewResponseInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesKolideInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesKolideInputRequest]
type devicePostureNewResponseInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesKolideInputRequest) implementsDevicePostureNewResponseInput() {
}

// Count Operator
type DevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown35 DevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown36 DevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown37 DevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown38 DevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown39 DevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type DevicePostureNewResponseInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                         `json:"total_score"`
	JSON       devicePostureNewResponseInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesTaniumInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesTaniumInputRequest]
type devicePostureNewResponseInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesTaniumInputRequest) implementsDevicePostureNewResponseInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown45 DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown46 DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown47 DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown48 DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown49 DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown55 DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown56 DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown57 DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown58 DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown59 DevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     devicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequest]
type devicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequest) implementsDevicePostureNewResponseInput() {
}

// Network status of device.
type DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown65 DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown66 DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown67 DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown68 DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown69 DevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type DevicePostureNewResponseMatch struct {
	Platform DevicePostureNewResponseMatchPlatform `json:"platform"`
	JSON     devicePostureNewResponseMatchJSON     `json:"-"`
}

// devicePostureNewResponseMatchJSON contains the JSON metadata for the struct
// [DevicePostureNewResponseMatch]
type devicePostureNewResponseMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureNewResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureNewResponseMatchPlatform string

const (
	DevicePostureNewResponseMatchPlatformWindows DevicePostureNewResponseMatchPlatform = "windows"
	DevicePostureNewResponseMatchPlatformMac     DevicePostureNewResponseMatchPlatform = "mac"
	DevicePostureNewResponseMatchPlatformLinux   DevicePostureNewResponseMatchPlatform = "linux"
	DevicePostureNewResponseMatchPlatformAndroid DevicePostureNewResponseMatchPlatform = "android"
	DevicePostureNewResponseMatchPlatformIos     DevicePostureNewResponseMatchPlatform = "ios"
)

// The type of device posture rule.
type DevicePostureNewResponseType string

const (
	DevicePostureNewResponseTypeFile              DevicePostureNewResponseType = "file"
	DevicePostureNewResponseTypeApplication       DevicePostureNewResponseType = "application"
	DevicePostureNewResponseTypeTanium            DevicePostureNewResponseType = "tanium"
	DevicePostureNewResponseTypeGateway           DevicePostureNewResponseType = "gateway"
	DevicePostureNewResponseTypeWARP              DevicePostureNewResponseType = "warp"
	DevicePostureNewResponseTypeDiskEncryption    DevicePostureNewResponseType = "disk_encryption"
	DevicePostureNewResponseTypeSentinelone       DevicePostureNewResponseType = "sentinelone"
	DevicePostureNewResponseTypeCarbonblack       DevicePostureNewResponseType = "carbonblack"
	DevicePostureNewResponseTypeFirewall          DevicePostureNewResponseType = "firewall"
	DevicePostureNewResponseTypeOSVersion         DevicePostureNewResponseType = "os_version"
	DevicePostureNewResponseTypeDomainJoined      DevicePostureNewResponseType = "domain_joined"
	DevicePostureNewResponseTypeClientCertificate DevicePostureNewResponseType = "client_certificate"
	DevicePostureNewResponseTypeUniqueClientID    DevicePostureNewResponseType = "unique_client_id"
	DevicePostureNewResponseTypeKolide            DevicePostureNewResponseType = "kolide"
	DevicePostureNewResponseTypeTaniumS2s         DevicePostureNewResponseType = "tanium_s2s"
	DevicePostureNewResponseTypeCrowdstrikeS2s    DevicePostureNewResponseType = "crowdstrike_s2s"
	DevicePostureNewResponseTypeIntune            DevicePostureNewResponseType = "intune"
	DevicePostureNewResponseTypeWorkspaceOne      DevicePostureNewResponseType = "workspace_one"
	DevicePostureNewResponseTypeSentineloneS2s    DevicePostureNewResponseType = "sentinelone_s2s"
)

type DevicePostureUpdateResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input DevicePostureUpdateResponseInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []DevicePostureUpdateResponseMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type DevicePostureUpdateResponseType `json:"type"`
	JSON devicePostureUpdateResponseJSON `json:"-"`
}

// devicePostureUpdateResponseJSON contains the JSON metadata for the struct
// [DevicePostureUpdateResponse]
type devicePostureUpdateResponseJSON struct {
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

func (r *DevicePostureUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by
// [DevicePostureUpdateResponseInputTeamsDevicesFileInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequest],
// [DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequest] or
// [DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureUpdateResponseInput interface {
	implementsDevicePostureUpdateResponseInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DevicePostureUpdateResponseInput)(nil)).Elem(), "")
}

type DevicePostureUpdateResponseInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                           `json:"thumbprint"`
	JSON       devicePostureUpdateResponseInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesFileInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesFileInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesFileInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Operating system
type DevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type DevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem DevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Operating System
type DevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type DevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                   `json:"domain"`
	JSON   devicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Operating System
type DevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                                                `json:"os_version_extra"`
	JSON           devicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Operating System
type DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator string

const (
	DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown75 DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown76 DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown77 DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown78 DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown79 DevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type DevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem DevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Operating System
type DevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type DevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                  `json:"thumbprint"`
	JSON       devicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Operating system
type DevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type DevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                  `json:"thumbprint"`
	JSON       devicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Operating system
type DevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type DevicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                       `json:"requireAll"`
	JSON       devicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequest) implementsDevicePostureUpdateResponseInput() {
}

type DevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                  `json:"thumbprint"`
	JSON       devicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Operating system
type DevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type DevicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                        `json:"cn,required"`
	JSON devicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequest) implementsDevicePostureUpdateResponseInput() {
}

type DevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         devicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Compliance Status
type DevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Operator
	Operator DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            devicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
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

func (r *DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Operator
type DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown85 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown86 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown87 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown88 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown89 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown95 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown96 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown97 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown98 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown99 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                             `json:"connection_id,required"`
	JSON         devicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Compliance Status
type DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                             `json:"issue_count,required"`
	JSON       devicePostureUpdateResponseInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesKolideInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Count Operator
type DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown105 DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown106 DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown107 DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown108 DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown109 DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                            `json:"total_score"`
	JSON       devicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown115 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown116 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown117 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown118 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown119 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown125 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown126 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown127 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown128 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown129 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     devicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Network status of device.
type DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown135 DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown136 DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown137 DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown138 DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown139 DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type DevicePostureUpdateResponseMatch struct {
	Platform DevicePostureUpdateResponseMatchPlatform `json:"platform"`
	JSON     devicePostureUpdateResponseMatchJSON     `json:"-"`
}

// devicePostureUpdateResponseMatchJSON contains the JSON metadata for the struct
// [DevicePostureUpdateResponseMatch]
type devicePostureUpdateResponseMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureUpdateResponseMatchPlatform string

const (
	DevicePostureUpdateResponseMatchPlatformWindows DevicePostureUpdateResponseMatchPlatform = "windows"
	DevicePostureUpdateResponseMatchPlatformMac     DevicePostureUpdateResponseMatchPlatform = "mac"
	DevicePostureUpdateResponseMatchPlatformLinux   DevicePostureUpdateResponseMatchPlatform = "linux"
	DevicePostureUpdateResponseMatchPlatformAndroid DevicePostureUpdateResponseMatchPlatform = "android"
	DevicePostureUpdateResponseMatchPlatformIos     DevicePostureUpdateResponseMatchPlatform = "ios"
)

// The type of device posture rule.
type DevicePostureUpdateResponseType string

const (
	DevicePostureUpdateResponseTypeFile              DevicePostureUpdateResponseType = "file"
	DevicePostureUpdateResponseTypeApplication       DevicePostureUpdateResponseType = "application"
	DevicePostureUpdateResponseTypeTanium            DevicePostureUpdateResponseType = "tanium"
	DevicePostureUpdateResponseTypeGateway           DevicePostureUpdateResponseType = "gateway"
	DevicePostureUpdateResponseTypeWARP              DevicePostureUpdateResponseType = "warp"
	DevicePostureUpdateResponseTypeDiskEncryption    DevicePostureUpdateResponseType = "disk_encryption"
	DevicePostureUpdateResponseTypeSentinelone       DevicePostureUpdateResponseType = "sentinelone"
	DevicePostureUpdateResponseTypeCarbonblack       DevicePostureUpdateResponseType = "carbonblack"
	DevicePostureUpdateResponseTypeFirewall          DevicePostureUpdateResponseType = "firewall"
	DevicePostureUpdateResponseTypeOSVersion         DevicePostureUpdateResponseType = "os_version"
	DevicePostureUpdateResponseTypeDomainJoined      DevicePostureUpdateResponseType = "domain_joined"
	DevicePostureUpdateResponseTypeClientCertificate DevicePostureUpdateResponseType = "client_certificate"
	DevicePostureUpdateResponseTypeUniqueClientID    DevicePostureUpdateResponseType = "unique_client_id"
	DevicePostureUpdateResponseTypeKolide            DevicePostureUpdateResponseType = "kolide"
	DevicePostureUpdateResponseTypeTaniumS2s         DevicePostureUpdateResponseType = "tanium_s2s"
	DevicePostureUpdateResponseTypeCrowdstrikeS2s    DevicePostureUpdateResponseType = "crowdstrike_s2s"
	DevicePostureUpdateResponseTypeIntune            DevicePostureUpdateResponseType = "intune"
	DevicePostureUpdateResponseTypeWorkspaceOne      DevicePostureUpdateResponseType = "workspace_one"
	DevicePostureUpdateResponseTypeSentineloneS2s    DevicePostureUpdateResponseType = "sentinelone_s2s"
)

type DevicePostureListResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input DevicePostureListResponseInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []DevicePostureListResponseMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type DevicePostureListResponseType `json:"type"`
	JSON devicePostureListResponseJSON `json:"-"`
}

// devicePostureListResponseJSON contains the JSON metadata for the struct
// [DevicePostureListResponse]
type devicePostureListResponseJSON struct {
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

func (r *DevicePostureListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by [DevicePostureListResponseInputTeamsDevicesFileInputRequest],
// [DevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequest],
// [DevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequest],
// [DevicePostureListResponseInputTeamsDevicesOSVersionInputRequest],
// [DevicePostureListResponseInputTeamsDevicesFirewallInputRequest],
// [DevicePostureListResponseInputTeamsDevicesSentineloneInputRequest],
// [DevicePostureListResponseInputTeamsDevicesCarbonblackInputRequest],
// [DevicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequest],
// [DevicePostureListResponseInputTeamsDevicesApplicationInputRequest],
// [DevicePostureListResponseInputTeamsDevicesClientCertificateInputRequest],
// [DevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequest],
// [DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequest],
// [DevicePostureListResponseInputTeamsDevicesIntuneInputRequest],
// [DevicePostureListResponseInputTeamsDevicesKolideInputRequest],
// [DevicePostureListResponseInputTeamsDevicesTaniumInputRequest] or
// [DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureListResponseInput interface {
	implementsDevicePostureListResponseInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DevicePostureListResponseInput)(nil)).Elem(), "")
}

type DevicePostureListResponseInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                         `json:"thumbprint"`
	JSON       devicePostureListResponseInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesFileInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesFileInputRequest]
type devicePostureListResponseInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesFileInputRequest) implementsDevicePostureListResponseInput() {
}

// Operating system
type DevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type DevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem DevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequest]
type devicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequest) implementsDevicePostureListResponseInput() {
}

// Operating System
type DevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type DevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                 `json:"domain"`
	JSON   devicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequest]
type devicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequest) implementsDevicePostureListResponseInput() {
}

// Operating System
type DevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type DevicePostureListResponseInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                                              `json:"os_version_extra"`
	JSON           devicePostureListResponseInputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesOSVersionInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesOSVersionInputRequest]
type devicePostureListResponseInputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesOSVersionInputRequest) implementsDevicePostureListResponseInput() {
}

// Operating System
type DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator string

const (
	DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown145 DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator = "<"
	DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown146 DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator = "<="
	DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown147 DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator = ">"
	DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown148 DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator = ">="
	DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown149 DevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type DevicePostureListResponseInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem DevicePostureListResponseInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureListResponseInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesFirewallInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesFirewallInputRequest]
type devicePostureListResponseInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesFirewallInputRequest) implementsDevicePostureListResponseInput() {
}

// Operating System
type DevicePostureListResponseInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureListResponseInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureListResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureListResponseInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureListResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type DevicePostureListResponseInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                `json:"thumbprint"`
	JSON       devicePostureListResponseInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesSentineloneInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesSentineloneInputRequest]
type devicePostureListResponseInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesSentineloneInputRequest) implementsDevicePostureListResponseInput() {
}

// Operating system
type DevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type DevicePostureListResponseInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                `json:"thumbprint"`
	JSON       devicePostureListResponseInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesCarbonblackInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesCarbonblackInputRequest]
type devicePostureListResponseInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesCarbonblackInputRequest) implementsDevicePostureListResponseInput() {
}

// Operating system
type DevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type DevicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                     `json:"requireAll"`
	JSON       devicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequest]
type devicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequest) implementsDevicePostureListResponseInput() {
}

type DevicePostureListResponseInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                `json:"thumbprint"`
	JSON       devicePostureListResponseInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesApplicationInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesApplicationInputRequest]
type devicePostureListResponseInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesApplicationInputRequest) implementsDevicePostureListResponseInput() {
}

// Operating system
type DevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type DevicePostureListResponseInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                      `json:"cn,required"`
	JSON devicePostureListResponseInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesClientCertificateInputRequest]
type devicePostureListResponseInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesClientCertificateInputRequest) implementsDevicePostureListResponseInput() {
}

type DevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                 `json:"connection_id,required"`
	JSON         devicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequest]
type devicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequest) implementsDevicePostureListResponseInput() {
}

// Compliance Status
type DevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Operator
	Operator DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            devicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequest]
type devicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
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

func (r *DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequest) implementsDevicePostureListResponseInput() {
}

// Operator
type DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown155 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown156 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown157 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown158 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown159 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown165 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown166 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown167 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown168 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown169 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type DevicePostureListResponseInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                           `json:"connection_id,required"`
	JSON         devicePostureListResponseInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesIntuneInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesIntuneInputRequest]
type devicePostureListResponseInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesIntuneInputRequest) implementsDevicePostureListResponseInput() {
}

// Compliance Status
type DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type DevicePostureListResponseInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                           `json:"issue_count,required"`
	JSON       devicePostureListResponseInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesKolideInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesKolideInputRequest]
type devicePostureListResponseInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesKolideInputRequest) implementsDevicePostureListResponseInput() {
}

// Count Operator
type DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown175 DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown176 DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown177 DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown178 DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown179 DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type DevicePostureListResponseInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                          `json:"total_score"`
	JSON       devicePostureListResponseInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesTaniumInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesTaniumInputRequest]
type devicePostureListResponseInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesTaniumInputRequest) implementsDevicePostureListResponseInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown185 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown186 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown187 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown188 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown189 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown195 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown196 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown197 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown198 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown199 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     devicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequest]
type devicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequest) implementsDevicePostureListResponseInput() {
}

// Network status of device.
type DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown205 DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown206 DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown207 DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown208 DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown209 DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type DevicePostureListResponseMatch struct {
	Platform DevicePostureListResponseMatchPlatform `json:"platform"`
	JSON     devicePostureListResponseMatchJSON     `json:"-"`
}

// devicePostureListResponseMatchJSON contains the JSON metadata for the struct
// [DevicePostureListResponseMatch]
type devicePostureListResponseMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureListResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureListResponseMatchPlatform string

const (
	DevicePostureListResponseMatchPlatformWindows DevicePostureListResponseMatchPlatform = "windows"
	DevicePostureListResponseMatchPlatformMac     DevicePostureListResponseMatchPlatform = "mac"
	DevicePostureListResponseMatchPlatformLinux   DevicePostureListResponseMatchPlatform = "linux"
	DevicePostureListResponseMatchPlatformAndroid DevicePostureListResponseMatchPlatform = "android"
	DevicePostureListResponseMatchPlatformIos     DevicePostureListResponseMatchPlatform = "ios"
)

// The type of device posture rule.
type DevicePostureListResponseType string

const (
	DevicePostureListResponseTypeFile              DevicePostureListResponseType = "file"
	DevicePostureListResponseTypeApplication       DevicePostureListResponseType = "application"
	DevicePostureListResponseTypeTanium            DevicePostureListResponseType = "tanium"
	DevicePostureListResponseTypeGateway           DevicePostureListResponseType = "gateway"
	DevicePostureListResponseTypeWARP              DevicePostureListResponseType = "warp"
	DevicePostureListResponseTypeDiskEncryption    DevicePostureListResponseType = "disk_encryption"
	DevicePostureListResponseTypeSentinelone       DevicePostureListResponseType = "sentinelone"
	DevicePostureListResponseTypeCarbonblack       DevicePostureListResponseType = "carbonblack"
	DevicePostureListResponseTypeFirewall          DevicePostureListResponseType = "firewall"
	DevicePostureListResponseTypeOSVersion         DevicePostureListResponseType = "os_version"
	DevicePostureListResponseTypeDomainJoined      DevicePostureListResponseType = "domain_joined"
	DevicePostureListResponseTypeClientCertificate DevicePostureListResponseType = "client_certificate"
	DevicePostureListResponseTypeUniqueClientID    DevicePostureListResponseType = "unique_client_id"
	DevicePostureListResponseTypeKolide            DevicePostureListResponseType = "kolide"
	DevicePostureListResponseTypeTaniumS2s         DevicePostureListResponseType = "tanium_s2s"
	DevicePostureListResponseTypeCrowdstrikeS2s    DevicePostureListResponseType = "crowdstrike_s2s"
	DevicePostureListResponseTypeIntune            DevicePostureListResponseType = "intune"
	DevicePostureListResponseTypeWorkspaceOne      DevicePostureListResponseType = "workspace_one"
	DevicePostureListResponseTypeSentineloneS2s    DevicePostureListResponseType = "sentinelone_s2s"
)

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

type DevicePostureGetResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input DevicePostureGetResponseInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []DevicePostureGetResponseMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type DevicePostureGetResponseType `json:"type"`
	JSON devicePostureGetResponseJSON `json:"-"`
}

// devicePostureGetResponseJSON contains the JSON metadata for the struct
// [DevicePostureGetResponse]
type devicePostureGetResponseJSON struct {
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

func (r *DevicePostureGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by [DevicePostureGetResponseInputTeamsDevicesFileInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesFirewallInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesSentineloneInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesApplicationInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesClientCertificateInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesIntuneInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesKolideInputRequest],
// [DevicePostureGetResponseInputTeamsDevicesTaniumInputRequest] or
// [DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureGetResponseInput interface {
	implementsDevicePostureGetResponseInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DevicePostureGetResponseInput)(nil)).Elem(), "")
}

type DevicePostureGetResponseInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                        `json:"thumbprint"`
	JSON       devicePostureGetResponseInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesFileInputRequestJSON contains the JSON
// metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesFileInputRequest]
type devicePostureGetResponseInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesFileInputRequest) implementsDevicePostureGetResponseInput() {
}

// Operating system
type DevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type DevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem DevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequest]
type devicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequest) implementsDevicePostureGetResponseInput() {
}

// Operating System
type DevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type DevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                `json:"domain"`
	JSON   devicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequest]
type devicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequest) implementsDevicePostureGetResponseInput() {
}

// Operating System
type DevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                                             `json:"os_version_extra"`
	JSON           devicePostureGetResponseInputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesOSVersionInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequest]
type devicePostureGetResponseInputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequest) implementsDevicePostureGetResponseInput() {
}

// Operating System
type DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator string

const (
	DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown215 DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown216 DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown217 DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown218 DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown219 DevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type DevicePostureGetResponseInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem DevicePostureGetResponseInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureGetResponseInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesFirewallInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesFirewallInputRequest]
type devicePostureGetResponseInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesFirewallInputRequest) implementsDevicePostureGetResponseInput() {
}

// Operating System
type DevicePostureGetResponseInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureGetResponseInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureGetResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureGetResponseInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureGetResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type DevicePostureGetResponseInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                               `json:"thumbprint"`
	JSON       devicePostureGetResponseInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesSentineloneInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesSentineloneInputRequest]
type devicePostureGetResponseInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesSentineloneInputRequest) implementsDevicePostureGetResponseInput() {
}

// Operating system
type DevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type DevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                               `json:"thumbprint"`
	JSON       devicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequest]
type devicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequest) implementsDevicePostureGetResponseInput() {
}

// Operating system
type DevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type DevicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                    `json:"requireAll"`
	JSON       devicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequest]
type devicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequest) implementsDevicePostureGetResponseInput() {
}

type DevicePostureGetResponseInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                               `json:"thumbprint"`
	JSON       devicePostureGetResponseInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesApplicationInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesApplicationInputRequest]
type devicePostureGetResponseInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesApplicationInputRequest) implementsDevicePostureGetResponseInput() {
}

// Operating system
type DevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type DevicePostureGetResponseInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                     `json:"cn,required"`
	JSON devicePostureGetResponseInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesClientCertificateInputRequest]
type devicePostureGetResponseInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesClientCertificateInputRequest) implementsDevicePostureGetResponseInput() {
}

type DevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                `json:"connection_id,required"`
	JSON         devicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequest]
type devicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequest) implementsDevicePostureGetResponseInput() {
}

// Compliance Status
type DevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Operator
	Operator DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            devicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequest]
type devicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
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

func (r *DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequest) implementsDevicePostureGetResponseInput() {
}

// Operator
type DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown225 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown226 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown227 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown228 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown229 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown235 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown236 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown237 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown238 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown239 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type DevicePostureGetResponseInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                          `json:"connection_id,required"`
	JSON         devicePostureGetResponseInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesIntuneInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesIntuneInputRequest]
type devicePostureGetResponseInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesIntuneInputRequest) implementsDevicePostureGetResponseInput() {
}

// Compliance Status
type DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type DevicePostureGetResponseInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                          `json:"issue_count,required"`
	JSON       devicePostureGetResponseInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesKolideInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesKolideInputRequest]
type devicePostureGetResponseInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesKolideInputRequest) implementsDevicePostureGetResponseInput() {
}

// Count Operator
type DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown245 DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown246 DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown247 DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown248 DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown249 DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type DevicePostureGetResponseInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                         `json:"total_score"`
	JSON       devicePostureGetResponseInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesTaniumInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesTaniumInputRequest]
type devicePostureGetResponseInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesTaniumInputRequest) implementsDevicePostureGetResponseInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown255 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown256 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown257 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown258 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown259 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown265 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown266 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown267 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown268 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown269 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     devicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequest]
type devicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequest) implementsDevicePostureGetResponseInput() {
}

// Network status of device.
type DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown275 DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown276 DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown277 DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown278 DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown279 DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type DevicePostureGetResponseMatch struct {
	Platform DevicePostureGetResponseMatchPlatform `json:"platform"`
	JSON     devicePostureGetResponseMatchJSON     `json:"-"`
}

// devicePostureGetResponseMatchJSON contains the JSON metadata for the struct
// [DevicePostureGetResponseMatch]
type devicePostureGetResponseMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureGetResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureGetResponseMatchPlatform string

const (
	DevicePostureGetResponseMatchPlatformWindows DevicePostureGetResponseMatchPlatform = "windows"
	DevicePostureGetResponseMatchPlatformMac     DevicePostureGetResponseMatchPlatform = "mac"
	DevicePostureGetResponseMatchPlatformLinux   DevicePostureGetResponseMatchPlatform = "linux"
	DevicePostureGetResponseMatchPlatformAndroid DevicePostureGetResponseMatchPlatform = "android"
	DevicePostureGetResponseMatchPlatformIos     DevicePostureGetResponseMatchPlatform = "ios"
)

// The type of device posture rule.
type DevicePostureGetResponseType string

const (
	DevicePostureGetResponseTypeFile              DevicePostureGetResponseType = "file"
	DevicePostureGetResponseTypeApplication       DevicePostureGetResponseType = "application"
	DevicePostureGetResponseTypeTanium            DevicePostureGetResponseType = "tanium"
	DevicePostureGetResponseTypeGateway           DevicePostureGetResponseType = "gateway"
	DevicePostureGetResponseTypeWARP              DevicePostureGetResponseType = "warp"
	DevicePostureGetResponseTypeDiskEncryption    DevicePostureGetResponseType = "disk_encryption"
	DevicePostureGetResponseTypeSentinelone       DevicePostureGetResponseType = "sentinelone"
	DevicePostureGetResponseTypeCarbonblack       DevicePostureGetResponseType = "carbonblack"
	DevicePostureGetResponseTypeFirewall          DevicePostureGetResponseType = "firewall"
	DevicePostureGetResponseTypeOSVersion         DevicePostureGetResponseType = "os_version"
	DevicePostureGetResponseTypeDomainJoined      DevicePostureGetResponseType = "domain_joined"
	DevicePostureGetResponseTypeClientCertificate DevicePostureGetResponseType = "client_certificate"
	DevicePostureGetResponseTypeUniqueClientID    DevicePostureGetResponseType = "unique_client_id"
	DevicePostureGetResponseTypeKolide            DevicePostureGetResponseType = "kolide"
	DevicePostureGetResponseTypeTaniumS2s         DevicePostureGetResponseType = "tanium_s2s"
	DevicePostureGetResponseTypeCrowdstrikeS2s    DevicePostureGetResponseType = "crowdstrike_s2s"
	DevicePostureGetResponseTypeIntune            DevicePostureGetResponseType = "intune"
	DevicePostureGetResponseTypeWorkspaceOne      DevicePostureGetResponseType = "workspace_one"
	DevicePostureGetResponseTypeSentineloneS2s    DevicePostureGetResponseType = "sentinelone_s2s"
)

type DevicePostureNewParams struct {
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

// The value to be checked against.
//
// Satisfied by [DevicePostureNewParamsInputTeamsDevicesFileInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesApplicationInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesKolideInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesTaniumInputRequest],
// [DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureNewParamsInput interface {
	implementsDevicePostureNewParamsInput()
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

func (r DevicePostureNewParamsInputTeamsDevicesFileInputRequest) implementsDevicePostureNewParamsInput() {
}

// Operating system
type DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest) implementsDevicePostureNewParamsInput() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest) implementsDevicePostureNewParamsInput() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

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

func (r DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest) implementsDevicePostureNewParamsInput() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown285 DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown286 DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown287 DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown288 DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown289 DevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesFirewallInputRequest) implementsDevicePostureNewParamsInput() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

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

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest) implementsDevicePostureNewParamsInput() {
}

// Operating system
type DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

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

func (r DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest) implementsDevicePostureNewParamsInput() {
}

// Operating system
type DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest) implementsDevicePostureNewParamsInput() {
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

func (r DevicePostureNewParamsInputTeamsDevicesApplicationInputRequest) implementsDevicePostureNewParamsInput() {
}

// Operating system
type DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest) implementsDevicePostureNewParamsInput() {
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

func (r DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest) implementsDevicePostureNewParamsInput() {
}

// Compliance Status
type DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

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

func (r DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest) implementsDevicePostureNewParamsInput() {
}

// Operator
type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown295 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown296 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown297 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown298 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown299 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown305 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown306 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown307 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown308 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown309 DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesIntuneInputRequest) implementsDevicePostureNewParamsInput() {
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

func (r DevicePostureNewParamsInputTeamsDevicesKolideInputRequest) implementsDevicePostureNewParamsInput() {
}

// Count Operator
type DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown315 DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown316 DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown317 DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown318 DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown319 DevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

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

func (r DevicePostureNewParamsInputTeamsDevicesTaniumInputRequest) implementsDevicePostureNewParamsInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown325 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown326 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown327 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown328 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown329 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown335 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown336 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown337 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown338 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown339 DevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

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

func (r DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest) implementsDevicePostureNewParamsInput() {
}

// Network status of device.
type DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown345 DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown346 DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown347 DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown348 DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown349 DevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

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

type DevicePostureNewResponseEnvelope struct {
	Errors   []DevicePostureNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureNewResponse                   `json:"result,required,nullable"`
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

// Whether the API call was successful.
type DevicePostureNewResponseEnvelopeSuccess bool

const (
	DevicePostureNewResponseEnvelopeSuccessTrue DevicePostureNewResponseEnvelopeSuccess = true
)

type DevicePostureUpdateParams struct {
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

// The value to be checked against.
//
// Satisfied by [DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest],
// [DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureUpdateParamsInput interface {
	implementsDevicePostureUpdateParamsInput()
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

func (r DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Operating system
type DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

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

func (r DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown355 DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown356 DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown357 DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown358 DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown359 DevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

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

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Operating system
type DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

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

func (r DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Operating system
type DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest) implementsDevicePostureUpdateParamsInput() {
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

func (r DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Operating system
type DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest) implementsDevicePostureUpdateParamsInput() {
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

func (r DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Compliance Status
type DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

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

func (r DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Operator
type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown365 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown366 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown367 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown368 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown369 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown375 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown376 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown377 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown378 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown379 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest) implementsDevicePostureUpdateParamsInput() {
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

func (r DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Count Operator
type DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown385 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown386 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown387 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown388 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown389 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

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

func (r DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown395 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown396 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown397 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown398 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown399 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown405 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown406 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown407 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown408 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown409 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

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

func (r DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Network status of device.
type DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown415 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown416 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown417 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown418 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown419 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

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

type DevicePostureUpdateResponseEnvelope struct {
	Errors   []DevicePostureUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureUpdateResponse                   `json:"result,required,nullable"`
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

// Whether the API call was successful.
type DevicePostureUpdateResponseEnvelopeSuccess bool

const (
	DevicePostureUpdateResponseEnvelopeSuccessTrue DevicePostureUpdateResponseEnvelopeSuccess = true
)

type DevicePostureListResponseEnvelope struct {
	Errors   []DevicePostureListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePostureListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePostureListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePostureListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePostureListResponseEnvelopeJSON       `json:"-"`
}

// devicePostureListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePostureListResponseEnvelope]
type devicePostureListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureListResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    devicePostureListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePostureListResponseEnvelopeErrors]
type devicePostureListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureListResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    devicePostureListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DevicePostureListResponseEnvelopeMessages]
type devicePostureListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePostureListResponseEnvelopeSuccess bool

const (
	DevicePostureListResponseEnvelopeSuccessTrue DevicePostureListResponseEnvelopeSuccess = true
)

type DevicePostureListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       devicePostureListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePostureListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DevicePostureListResponseEnvelopeResultInfo]
type devicePostureListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// Whether the API call was successful.
type DevicePostureDeleteResponseEnvelopeSuccess bool

const (
	DevicePostureDeleteResponseEnvelopeSuccessTrue DevicePostureDeleteResponseEnvelopeSuccess = true
)

type DevicePostureGetResponseEnvelope struct {
	Errors   []DevicePostureGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureGetResponse                   `json:"result,required,nullable"`
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

// Whether the API call was successful.
type DevicePostureGetResponseEnvelopeSuccess bool

const (
	DevicePostureGetResponseEnvelopeSuccessTrue DevicePostureGetResponseEnvelopeSuccess = true
)
