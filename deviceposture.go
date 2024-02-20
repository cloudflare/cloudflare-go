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
func (r *DevicePostureService) New(ctx context.Context, identifier interface{}, body DevicePostureNewParams, opts ...option.RequestOption) (res *DevicePostureNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches device posture rules for a Zero Trust account.
func (r *DevicePostureService) List(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DevicePostureListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a device posture rule.
func (r *DevicePostureService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *DevicePostureDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single device posture rule.
func (r *DevicePostureService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *DevicePostureGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a device posture rule.
func (r *DevicePostureService) Replace(ctx context.Context, identifier interface{}, uuid string, body DevicePostureReplaceParams, opts ...option.RequestOption) (res *DevicePostureReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
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
// [DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequest],
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

type DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OsVersionExtra string                                                             `json:"os_version_extra"`
	JSON           devicePostureNewResponseInputTeamsDevicesOsVersionInputRequestJSON `json:"-"`
}

// devicePostureNewResponseInputTeamsDevicesOsVersionInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequest]
type devicePostureNewResponseInputTeamsDevicesOsVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequest) implementsDevicePostureNewResponseInput() {
}

// Operating System
type DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem string

const (
	DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperatingSystemWindows DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperator string

const (
	DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown5 DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperator = "<"
	DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown6 DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperator = "<="
	DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown7 DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperator = ">"
	DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown8 DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperator = ">="
	DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown9 DevicePostureNewResponseInputTeamsDevicesOsVersionInputRequestOperator = "=="
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
	// Operator
	Operator DevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	Os string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
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
	Operator        apijson.Field
	Os              apijson.Field
	Overall         apijson.Field
	SensorConfig    apijson.Field
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
	DevicePostureNewResponseTypeWarp              DevicePostureNewResponseType = "warp"
	DevicePostureNewResponseTypeDiskEncryption    DevicePostureNewResponseType = "disk_encryption"
	DevicePostureNewResponseTypeSentinelone       DevicePostureNewResponseType = "sentinelone"
	DevicePostureNewResponseTypeCarbonblack       DevicePostureNewResponseType = "carbonblack"
	DevicePostureNewResponseTypeFirewall          DevicePostureNewResponseType = "firewall"
	DevicePostureNewResponseTypeOsVersion         DevicePostureNewResponseType = "os_version"
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
// [DevicePostureListResponseInputTeamsDevicesOsVersionInputRequest],
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

type DevicePostureListResponseInputTeamsDevicesOsVersionInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OsVersionExtra string                                                              `json:"os_version_extra"`
	JSON           devicePostureListResponseInputTeamsDevicesOsVersionInputRequestJSON `json:"-"`
}

// devicePostureListResponseInputTeamsDevicesOsVersionInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureListResponseInputTeamsDevicesOsVersionInputRequest]
type devicePostureListResponseInputTeamsDevicesOsVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureListResponseInputTeamsDevicesOsVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureListResponseInputTeamsDevicesOsVersionInputRequest) implementsDevicePostureListResponseInput() {
}

// Operating System
type DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem string

const (
	DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperatingSystemWindows DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperator string

const (
	DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown75 DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperator = "<"
	DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown76 DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperator = "<="
	DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown77 DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperator = ">"
	DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown78 DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperator = ">="
	DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown79 DevicePostureListResponseInputTeamsDevicesOsVersionInputRequestOperator = "=="
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
	// Operator
	Operator DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	Os string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
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
	Operator        apijson.Field
	Os              apijson.Field
	Overall         apijson.Field
	SensorConfig    apijson.Field
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
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown85 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown86 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown87 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown88 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown89 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown95 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown96 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown97 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown98 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown99 DevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
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
	DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown105 DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown106 DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown107 DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown108 DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown109 DevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
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
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown115 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown116 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown117 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown118 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown119 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
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
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown125 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown126 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown127 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown128 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown129 DevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
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
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown135 DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown136 DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown137 DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown138 DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown139 DevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
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
	DevicePostureListResponseTypeWarp              DevicePostureListResponseType = "warp"
	DevicePostureListResponseTypeDiskEncryption    DevicePostureListResponseType = "disk_encryption"
	DevicePostureListResponseTypeSentinelone       DevicePostureListResponseType = "sentinelone"
	DevicePostureListResponseTypeCarbonblack       DevicePostureListResponseType = "carbonblack"
	DevicePostureListResponseTypeFirewall          DevicePostureListResponseType = "firewall"
	DevicePostureListResponseTypeOsVersion         DevicePostureListResponseType = "os_version"
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
// [DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequest],
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

type DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OsVersionExtra string                                                             `json:"os_version_extra"`
	JSON           devicePostureGetResponseInputTeamsDevicesOsVersionInputRequestJSON `json:"-"`
}

// devicePostureGetResponseInputTeamsDevicesOsVersionInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequest]
type devicePostureGetResponseInputTeamsDevicesOsVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequest) implementsDevicePostureGetResponseInput() {
}

// Operating System
type DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem string

const (
	DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatingSystemWindows DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperator string

const (
	DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown145 DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown146 DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown147 DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown148 DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown149 DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperator = "=="
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
	// Operator
	Operator DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	Os string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
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
	Operator        apijson.Field
	Os              apijson.Field
	Overall         apijson.Field
	SensorConfig    apijson.Field
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
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown155 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown156 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown157 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown158 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown159 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown165 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown166 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown167 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown168 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown169 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
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
	DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown175 DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown176 DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown177 DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown178 DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown179 DevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
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
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown185 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown186 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown187 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown188 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown189 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
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
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown195 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown196 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown197 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown198 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown199 DevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
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
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown205 DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown206 DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown207 DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown208 DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown209 DevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
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
	DevicePostureGetResponseTypeWarp              DevicePostureGetResponseType = "warp"
	DevicePostureGetResponseTypeDiskEncryption    DevicePostureGetResponseType = "disk_encryption"
	DevicePostureGetResponseTypeSentinelone       DevicePostureGetResponseType = "sentinelone"
	DevicePostureGetResponseTypeCarbonblack       DevicePostureGetResponseType = "carbonblack"
	DevicePostureGetResponseTypeFirewall          DevicePostureGetResponseType = "firewall"
	DevicePostureGetResponseTypeOsVersion         DevicePostureGetResponseType = "os_version"
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

type DevicePostureReplaceResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input DevicePostureReplaceResponseInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []DevicePostureReplaceResponseMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type DevicePostureReplaceResponseType `json:"type"`
	JSON devicePostureReplaceResponseJSON `json:"-"`
}

// devicePostureReplaceResponseJSON contains the JSON metadata for the struct
// [DevicePostureReplaceResponse]
type devicePostureReplaceResponseJSON struct {
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

func (r *DevicePostureReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by
// [DevicePostureReplaceResponseInputTeamsDevicesFileInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesDomainJoinedInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesFirewallInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesDiskEncryptionInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesApplicationInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesClientCertificateInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequest],
// [DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequest] or
// [DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureReplaceResponseInput interface {
	implementsDevicePostureReplaceResponseInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DevicePostureReplaceResponseInput)(nil)).Elem(), "")
}

type DevicePostureReplaceResponseInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureReplaceResponseInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                            `json:"thumbprint"`
	JSON       devicePostureReplaceResponseInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesFileInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesFileInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesFileInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Operating system
type DevicePostureReplaceResponseInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureReplaceResponseInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureReplaceResponseInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureReplaceResponseInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureReplaceResponseInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureReplaceResponseInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureReplaceResponseInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type DevicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem DevicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Operating System
type DevicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureReplaceResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type DevicePostureReplaceResponseInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureReplaceResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                    `json:"domain"`
	JSON   devicePostureReplaceResponseInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesDomainJoinedInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesDomainJoinedInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Operating System
type DevicePostureReplaceResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureReplaceResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureReplaceResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OsVersionExtra string                                                                 `json:"os_version_extra"`
	JSON           devicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestJSON `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Operating System
type DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem string

const (
	DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperatingSystemWindows DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperator string

const (
	DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown215 DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperator = "<"
	DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown216 DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperator = "<="
	DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown217 DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperator = ">"
	DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown218 DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperator = ">="
	DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown219 DevicePostureReplaceResponseInputTeamsDevicesOsVersionInputRequestOperator = "=="
)

type DevicePostureReplaceResponseInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem DevicePostureReplaceResponseInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureReplaceResponseInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesFirewallInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesFirewallInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesFirewallInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Operating System
type DevicePostureReplaceResponseInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureReplaceResponseInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureReplaceResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureReplaceResponseInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureReplaceResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type DevicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                   `json:"thumbprint"`
	JSON       devicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Operating system
type DevicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureReplaceResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type DevicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                   `json:"thumbprint"`
	JSON       devicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Operating system
type DevicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureReplaceResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type DevicePostureReplaceResponseInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                        `json:"requireAll"`
	JSON       devicePostureReplaceResponseInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesDiskEncryptionInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesDiskEncryptionInputRequest) implementsDevicePostureReplaceResponseInput() {
}

type DevicePostureReplaceResponseInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureReplaceResponseInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                   `json:"thumbprint"`
	JSON       devicePostureReplaceResponseInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesApplicationInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesApplicationInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesApplicationInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Operating system
type DevicePostureReplaceResponseInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureReplaceResponseInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureReplaceResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureReplaceResponseInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureReplaceResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureReplaceResponseInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureReplaceResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type DevicePostureReplaceResponseInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                         `json:"cn,required"`
	JSON devicePostureReplaceResponseInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesClientCertificateInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesClientCertificateInputRequest) implementsDevicePostureReplaceResponseInput() {
}

type DevicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                    `json:"connection_id,required"`
	JSON         devicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Compliance Status
type DevicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureReplaceResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Operator
	Operator DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	Os string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            devicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
	ConnectionID    apijson.Field
	Operator        apijson.Field
	Os              apijson.Field
	Overall         apijson.Field
	SensorConfig    apijson.Field
	Version         apijson.Field
	VersionOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Operator
type DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown225 DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown226 DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown227 DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown228 DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown229 DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown235 DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown236 DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown237 DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown238 DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown239 DevicePostureReplaceResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                              `json:"connection_id,required"`
	JSON         devicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Compliance Status
type DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureReplaceResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                              `json:"issue_count,required"`
	JSON       devicePostureReplaceResponseInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesKolideInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Count Operator
type DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown245 DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown246 DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown247 DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown248 DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown249 DevicePostureReplaceResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                             `json:"total_score"`
	JSON       devicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestJSON contains the
// JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown255 DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown256 DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown257 DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown258 DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown259 DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown265 DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown266 DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown267 DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown268 DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown269 DevicePostureReplaceResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     devicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// devicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequest]
type devicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequest) implementsDevicePostureReplaceResponseInput() {
}

// Network status of device.
type DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown275 DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown276 DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown277 DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown278 DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown279 DevicePostureReplaceResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type DevicePostureReplaceResponseMatch struct {
	Platform DevicePostureReplaceResponseMatchPlatform `json:"platform"`
	JSON     devicePostureReplaceResponseMatchJSON     `json:"-"`
}

// devicePostureReplaceResponseMatchJSON contains the JSON metadata for the struct
// [DevicePostureReplaceResponseMatch]
type devicePostureReplaceResponseMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureReplaceResponseMatchPlatform string

const (
	DevicePostureReplaceResponseMatchPlatformWindows DevicePostureReplaceResponseMatchPlatform = "windows"
	DevicePostureReplaceResponseMatchPlatformMac     DevicePostureReplaceResponseMatchPlatform = "mac"
	DevicePostureReplaceResponseMatchPlatformLinux   DevicePostureReplaceResponseMatchPlatform = "linux"
	DevicePostureReplaceResponseMatchPlatformAndroid DevicePostureReplaceResponseMatchPlatform = "android"
	DevicePostureReplaceResponseMatchPlatformIos     DevicePostureReplaceResponseMatchPlatform = "ios"
)

// The type of device posture rule.
type DevicePostureReplaceResponseType string

const (
	DevicePostureReplaceResponseTypeFile              DevicePostureReplaceResponseType = "file"
	DevicePostureReplaceResponseTypeApplication       DevicePostureReplaceResponseType = "application"
	DevicePostureReplaceResponseTypeTanium            DevicePostureReplaceResponseType = "tanium"
	DevicePostureReplaceResponseTypeGateway           DevicePostureReplaceResponseType = "gateway"
	DevicePostureReplaceResponseTypeWarp              DevicePostureReplaceResponseType = "warp"
	DevicePostureReplaceResponseTypeDiskEncryption    DevicePostureReplaceResponseType = "disk_encryption"
	DevicePostureReplaceResponseTypeSentinelone       DevicePostureReplaceResponseType = "sentinelone"
	DevicePostureReplaceResponseTypeCarbonblack       DevicePostureReplaceResponseType = "carbonblack"
	DevicePostureReplaceResponseTypeFirewall          DevicePostureReplaceResponseType = "firewall"
	DevicePostureReplaceResponseTypeOsVersion         DevicePostureReplaceResponseType = "os_version"
	DevicePostureReplaceResponseTypeDomainJoined      DevicePostureReplaceResponseType = "domain_joined"
	DevicePostureReplaceResponseTypeClientCertificate DevicePostureReplaceResponseType = "client_certificate"
	DevicePostureReplaceResponseTypeUniqueClientID    DevicePostureReplaceResponseType = "unique_client_id"
	DevicePostureReplaceResponseTypeKolide            DevicePostureReplaceResponseType = "kolide"
	DevicePostureReplaceResponseTypeTaniumS2s         DevicePostureReplaceResponseType = "tanium_s2s"
	DevicePostureReplaceResponseTypeCrowdstrikeS2s    DevicePostureReplaceResponseType = "crowdstrike_s2s"
	DevicePostureReplaceResponseTypeIntune            DevicePostureReplaceResponseType = "intune"
	DevicePostureReplaceResponseTypeWorkspaceOne      DevicePostureReplaceResponseType = "workspace_one"
	DevicePostureReplaceResponseTypeSentineloneS2s    DevicePostureReplaceResponseType = "sentinelone_s2s"
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
	DevicePostureNewParamsTypeWarp              DevicePostureNewParamsType = "warp"
	DevicePostureNewParamsTypeDiskEncryption    DevicePostureNewParamsType = "disk_encryption"
	DevicePostureNewParamsTypeSentinelone       DevicePostureNewParamsType = "sentinelone"
	DevicePostureNewParamsTypeCarbonblack       DevicePostureNewParamsType = "carbonblack"
	DevicePostureNewParamsTypeFirewall          DevicePostureNewParamsType = "firewall"
	DevicePostureNewParamsTypeOsVersion         DevicePostureNewParamsType = "os_version"
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
// [DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequest],
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

type DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// Operator
	Operator param.Field[DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperator] `json:"operator,required"`
	// Version of OS
	Version param.Field[string] `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OsVersionExtra param.Field[string] `json:"os_version_extra"`
}

func (r DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequest) implementsDevicePostureNewParamsInput() {
}

// Operating System
type DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperatingSystem string

const (
	DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperatingSystemWindows DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperator string

const (
	DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown285 DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperator = "<"
	DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown286 DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperator = "<="
	DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown287 DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperator = ">"
	DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown288 DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperator = ">="
	DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown289 DevicePostureNewParamsInputTeamsDevicesOsVersionInputRequestOperator = "=="
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
	// Operator
	Operator param.Field[DevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator] `json:"operator"`
	// Os Version
	Os param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
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

type DevicePostureReplaceParams struct {
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[DevicePostureReplaceParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[DevicePostureReplaceParamsInput] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]DevicePostureReplaceParamsMatch] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r DevicePostureReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type DevicePostureReplaceParamsType string

const (
	DevicePostureReplaceParamsTypeFile              DevicePostureReplaceParamsType = "file"
	DevicePostureReplaceParamsTypeApplication       DevicePostureReplaceParamsType = "application"
	DevicePostureReplaceParamsTypeTanium            DevicePostureReplaceParamsType = "tanium"
	DevicePostureReplaceParamsTypeGateway           DevicePostureReplaceParamsType = "gateway"
	DevicePostureReplaceParamsTypeWarp              DevicePostureReplaceParamsType = "warp"
	DevicePostureReplaceParamsTypeDiskEncryption    DevicePostureReplaceParamsType = "disk_encryption"
	DevicePostureReplaceParamsTypeSentinelone       DevicePostureReplaceParamsType = "sentinelone"
	DevicePostureReplaceParamsTypeCarbonblack       DevicePostureReplaceParamsType = "carbonblack"
	DevicePostureReplaceParamsTypeFirewall          DevicePostureReplaceParamsType = "firewall"
	DevicePostureReplaceParamsTypeOsVersion         DevicePostureReplaceParamsType = "os_version"
	DevicePostureReplaceParamsTypeDomainJoined      DevicePostureReplaceParamsType = "domain_joined"
	DevicePostureReplaceParamsTypeClientCertificate DevicePostureReplaceParamsType = "client_certificate"
	DevicePostureReplaceParamsTypeUniqueClientID    DevicePostureReplaceParamsType = "unique_client_id"
	DevicePostureReplaceParamsTypeKolide            DevicePostureReplaceParamsType = "kolide"
	DevicePostureReplaceParamsTypeTaniumS2s         DevicePostureReplaceParamsType = "tanium_s2s"
	DevicePostureReplaceParamsTypeCrowdstrikeS2s    DevicePostureReplaceParamsType = "crowdstrike_s2s"
	DevicePostureReplaceParamsTypeIntune            DevicePostureReplaceParamsType = "intune"
	DevicePostureReplaceParamsTypeWorkspaceOne      DevicePostureReplaceParamsType = "workspace_one"
	DevicePostureReplaceParamsTypeSentineloneS2s    DevicePostureReplaceParamsType = "sentinelone_s2s"
)

// The value to be checked against.
//
// Satisfied by [DevicePostureReplaceParamsInputTeamsDevicesFileInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesUniqueClientIDInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesDomainJoinedInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesFirewallInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesSentineloneInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesCarbonblackInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesDiskEncryptionInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesApplicationInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesClientCertificateInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesWorkspaceOneInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequest],
// [DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureReplaceParamsInput interface {
	implementsDevicePostureReplaceParamsInput()
}

type DevicePostureReplaceParamsInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureReplaceParamsInputTeamsDevicesFileInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesFileInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesFileInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Operating system
type DevicePostureReplaceParamsInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureReplaceParamsInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureReplaceParamsInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureReplaceParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureReplaceParamsInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureReplaceParamsInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureReplaceParamsInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type DevicePostureReplaceParamsInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureReplaceParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesUniqueClientIDInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Operating System
type DevicePostureReplaceParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureReplaceParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureReplaceParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureReplaceParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureReplaceParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureReplaceParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureReplaceParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type DevicePostureReplaceParamsInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureReplaceParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesDomainJoinedInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Operating System
type DevicePostureReplaceParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureReplaceParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureReplaceParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// Operator
	Operator param.Field[DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperator] `json:"operator,required"`
	// Version of OS
	Version param.Field[string] `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OsVersionExtra param.Field[string] `json:"os_version_extra"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Operating System
type DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperatingSystem string

const (
	DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperatingSystemWindows DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperator string

const (
	DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown355 DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperator = "<"
	DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown356 DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperator = "<="
	DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown357 DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperator = ">"
	DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown358 DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperator = ">="
	DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown359 DevicePostureReplaceParamsInputTeamsDevicesOsVersionInputRequestOperator = "=="
)

type DevicePostureReplaceParamsInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureReplaceParamsInputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesFirewallInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Operating System
type DevicePostureReplaceParamsInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureReplaceParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureReplaceParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureReplaceParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureReplaceParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type DevicePostureReplaceParamsInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureReplaceParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesSentineloneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesSentineloneInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Operating system
type DevicePostureReplaceParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureReplaceParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureReplaceParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureReplaceParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureReplaceParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureReplaceParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureReplaceParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type DevicePostureReplaceParamsInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureReplaceParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesCarbonblackInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesCarbonblackInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Operating system
type DevicePostureReplaceParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureReplaceParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureReplaceParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureReplaceParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureReplaceParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureReplaceParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureReplaceParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type DevicePostureReplaceParamsInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesDiskEncryptionInputRequest) implementsDevicePostureReplaceParamsInput() {
}

type DevicePostureReplaceParamsInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureReplaceParamsInputTeamsDevicesApplicationInputRequestOperatingSystem] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesApplicationInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesApplicationInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Operating system
type DevicePostureReplaceParamsInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureReplaceParamsInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureReplaceParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureReplaceParamsInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureReplaceParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureReplaceParamsInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureReplaceParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type DevicePostureReplaceParamsInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesClientCertificateInputRequest) implementsDevicePostureReplaceParamsInput() {
}

type DevicePostureReplaceParamsInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureReplaceParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesWorkspaceOneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesWorkspaceOneInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Compliance Status
type DevicePostureReplaceParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureReplaceParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureReplaceParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureReplaceParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureReplaceParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureReplaceParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureReplaceParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Operator
	Operator param.Field[DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestOperator] `json:"operator"`
	// Os Version
	Os param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator] `json:"versionOperator"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Operator
type DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown365 DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown366 DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown367 DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown368 DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown369 DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown375 DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown376 DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown377 DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown378 DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown379 DevicePostureReplaceParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Compliance Status
type DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureReplaceParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequestCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Count Operator
type DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown385 DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown386 DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown387 DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown388 DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown389 DevicePostureReplaceParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown395 DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown396 DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown397 DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown398 DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown399 DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown405 DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown406 DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown407 DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown408 DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown409 DevicePostureReplaceParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus] `json:"network_status"`
	// operator
	Operator param.Field[DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestOperator] `json:"operator"`
}

func (r DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequest) implementsDevicePostureReplaceParamsInput() {
}

// Network status of device.
type DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown415 DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown416 DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown417 DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown418 DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown419 DevicePostureReplaceParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type DevicePostureReplaceParamsMatch struct {
	Platform param.Field[DevicePostureReplaceParamsMatchPlatform] `json:"platform"`
}

func (r DevicePostureReplaceParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePostureReplaceParamsMatchPlatform string

const (
	DevicePostureReplaceParamsMatchPlatformWindows DevicePostureReplaceParamsMatchPlatform = "windows"
	DevicePostureReplaceParamsMatchPlatformMac     DevicePostureReplaceParamsMatchPlatform = "mac"
	DevicePostureReplaceParamsMatchPlatformLinux   DevicePostureReplaceParamsMatchPlatform = "linux"
	DevicePostureReplaceParamsMatchPlatformAndroid DevicePostureReplaceParamsMatchPlatform = "android"
	DevicePostureReplaceParamsMatchPlatformIos     DevicePostureReplaceParamsMatchPlatform = "ios"
)

type DevicePostureReplaceResponseEnvelope struct {
	Errors   []DevicePostureReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureReplaceResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureReplaceResponseEnvelopeJSON    `json:"-"`
}

// devicePostureReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePostureReplaceResponseEnvelope]
type devicePostureReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureReplaceResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    devicePostureReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DevicePostureReplaceResponseEnvelopeErrors]
type devicePostureReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureReplaceResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    devicePostureReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DevicePostureReplaceResponseEnvelopeMessages]
type devicePostureReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePostureReplaceResponseEnvelopeSuccess bool

const (
	DevicePostureReplaceResponseEnvelopeSuccessTrue DevicePostureReplaceResponseEnvelopeSuccess = true
)
