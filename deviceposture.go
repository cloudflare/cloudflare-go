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

// Updates a device posture rule.
func (r *DevicePostureService) Update(ctx context.Context, identifier interface{}, uuid string, body DevicePostureUpdateParams, opts ...option.RequestOption) (res *DevicePostureUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
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

// Creates a new device posture rule.
func (r *DevicePostureService) DevicePostureRulesNewDevicePostureRule(ctx context.Context, identifier interface{}, body DevicePostureDevicePostureRulesNewDevicePostureRuleParams, opts ...option.RequestOption) (res *DevicePostureDevicePostureRulesNewDevicePostureRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches device posture rules for a Zero Trust account.
func (r *DevicePostureService) DevicePostureRulesListDevicePostureRules(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DevicePostureDevicePostureRulesListDevicePostureRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
// [DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequest],
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

type DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OsVersionExtra string                                                                `json:"os_version_extra"`
	JSON           devicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestJSON `json:"-"`
}

// devicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestJSON contains
// the JSON metadata for the struct
// [DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequest]
type devicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequest) implementsDevicePostureUpdateResponseInput() {
}

// Operating System
type DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem string

const (
	DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperatingSystemWindows DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperator string

const (
	DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown5 DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown6 DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown7 DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown8 DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown9 DevicePostureUpdateResponseInputTeamsDevicesOsVersionInputRequestOperator = "=="
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
	// Operator
	Operator DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	Os string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
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
	Operator        apijson.Field
	Os              apijson.Field
	Overall         apijson.Field
	SensorConfig    apijson.Field
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
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown15 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown16 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown17 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown18 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown19 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown25 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown26 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown27 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown28 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown29 DevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
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
	DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown35 DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown36 DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown37 DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown38 DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown39 DevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
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
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown45 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown46 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown47 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown48 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown49 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
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
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown55 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown56 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown57 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown58 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown59 DevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
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
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown65 DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown66 DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown67 DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown68 DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown69 DevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
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
	DevicePostureUpdateResponseTypeWarp              DevicePostureUpdateResponseType = "warp"
	DevicePostureUpdateResponseTypeDiskEncryption    DevicePostureUpdateResponseType = "disk_encryption"
	DevicePostureUpdateResponseTypeSentinelone       DevicePostureUpdateResponseType = "sentinelone"
	DevicePostureUpdateResponseTypeCarbonblack       DevicePostureUpdateResponseType = "carbonblack"
	DevicePostureUpdateResponseTypeFirewall          DevicePostureUpdateResponseType = "firewall"
	DevicePostureUpdateResponseTypeOsVersion         DevicePostureUpdateResponseType = "os_version"
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

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType `json:"type"`
	JSON devicePostureDevicePostureRulesNewDevicePostureRuleResponseJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseJSON contains the
// JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponse]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseJSON struct {
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

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDomainJoinedInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDiskEncryptionInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesClientCertificateInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequest]
// or
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput interface {
	implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput)(nil)).Elem(), "")
}

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                           `json:"thumbprint"`
	JSON       devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Operating system
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Operating System
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                                                   `json:"domain"`
	JSON   devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDomainJoinedInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDomainJoinedInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Operating System
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OsVersionExtra string                                                                                                `json:"os_version_extra"`
	JSON           devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Operating System
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown75 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown76 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown77 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown78 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown79 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesOsVersionInputRequestOperator = "=="
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Operating System
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                  `json:"thumbprint"`
	JSON       devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Operating system
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                  `json:"thumbprint"`
	JSON       devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Operating system
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                                                       `json:"requireAll"`
	JSON       devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDiskEncryptionInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesDiskEncryptionInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                  `json:"thumbprint"`
	JSON       devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Operating system
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                                                        `json:"cn,required"`
	JSON devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesClientCertificateInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesClientCertificateInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                                                   `json:"connection_id,required"`
	JSON         devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Compliance Status
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Operator
	Operator DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	Os string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
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

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Operator
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown85 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown86 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown87 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown88 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown89 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown95 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown96 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown97 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown98 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown99 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                                             `json:"connection_id,required"`
	JSON         devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Compliance Status
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                                                             `json:"issue_count,required"`
	JSON       devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Count Operator
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown105 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown106 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown107 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown108 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown109 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                                                            `json:"total_score"`
	JSON       devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown115 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown116 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown117 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown118 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown119 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown125 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown126 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown127 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown128 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown129 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequest]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleResponseInput() {
}

// Network status of device.
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown135 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown136 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown137 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown138 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown139 DevicePostureDevicePostureRulesNewDevicePostureRuleResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatch struct {
	Platform DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchPlatform `json:"platform"`
	JSON     devicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchJSON     `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchJSON contains
// the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatch]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchPlatform string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchPlatformWindows DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchPlatform = "windows"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchPlatformMac     DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchPlatform = "mac"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchPlatformLinux   DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchPlatform = "linux"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchPlatformAndroid DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchPlatform = "android"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchPlatformIos     DevicePostureDevicePostureRulesNewDevicePostureRuleResponseMatchPlatform = "ios"
)

// The type of device posture rule.
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeFile              DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "file"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeApplication       DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "application"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeTanium            DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "tanium"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeGateway           DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "gateway"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeWarp              DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "warp"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeDiskEncryption    DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "disk_encryption"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeSentinelone       DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "sentinelone"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeCarbonblack       DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "carbonblack"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeFirewall          DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "firewall"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeOsVersion         DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "os_version"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeDomainJoined      DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "domain_joined"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeClientCertificate DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "client_certificate"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeUniqueClientID    DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "unique_client_id"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeKolide            DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "kolide"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeTaniumS2s         DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "tanium_s2s"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeCrowdstrikeS2s    DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "crowdstrike_s2s"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeIntune            DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "intune"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeWorkspaceOne      DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "workspace_one"
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseTypeSentineloneS2s    DevicePostureDevicePostureRulesNewDevicePostureRuleResponseType = "sentinelone_s2s"
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input DevicePostureDevicePostureRulesListDevicePostureRulesResponseInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type DevicePostureDevicePostureRulesListDevicePostureRulesResponseType `json:"type"`
	JSON devicePostureDevicePostureRulesListDevicePostureRulesResponseJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseJSON contains the
// JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponse]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseJSON struct {
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

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDomainJoinedInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDiskEncryptionInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesClientCertificateInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequest],
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequest]
// or
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInput interface {
	implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DevicePostureDevicePostureRulesListDevicePostureRulesResponseInput)(nil)).Elem(), "")
}

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                             `json:"thumbprint"`
	JSON       devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Operating system
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Operating System
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                                                     `json:"domain"`
	JSON   devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDomainJoinedInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDomainJoinedInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Operating System
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequest struct {
	// Operating System
	OperatingSystem DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OsVersionExtra string                                                                                                  `json:"os_version_extra"`
	JSON           devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Operating System
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperator string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown145 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperator = "<"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown146 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperator = "<="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown147 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperator = ">"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown148 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperator = ">="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown149 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesOsVersionInputRequestOperator = "=="
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Operating System
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                    `json:"thumbprint"`
	JSON       devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Operating system
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                    `json:"thumbprint"`
	JSON       devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Operating system
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                                                         `json:"requireAll"`
	JSON       devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDiskEncryptionInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesDiskEncryptionInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                    `json:"thumbprint"`
	JSON       devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Operating system
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                                                          `json:"cn,required"`
	JSON devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesClientCertificateInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesClientCertificateInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                                                     `json:"connection_id,required"`
	JSON         devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Compliance Status
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Operator
	Operator DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	Os string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
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

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Operator
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown155 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown156 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown157 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown158 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown159 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown165 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown166 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown167 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown168 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown169 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                                               `json:"connection_id,required"`
	JSON         devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Compliance Status
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                                                               `json:"issue_count,required"`
	JSON       devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Count Operator
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown175 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown176 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown177 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown178 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown179 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                                                              `json:"total_score"`
	JSON       devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown185 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown186 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown187 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown188 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown189 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown195 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown196 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown197 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown198 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown199 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequest]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequest) implementsDevicePostureDevicePostureRulesListDevicePostureRulesResponseInput() {
}

// Network status of device.
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown205 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown206 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown207 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown208 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown209 DevicePostureDevicePostureRulesListDevicePostureRulesResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatch struct {
	Platform DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatchPlatform `json:"platform"`
	JSON     devicePostureDevicePostureRulesListDevicePostureRulesResponseMatchJSON     `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseMatchJSON contains
// the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatch]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatchPlatform string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatchPlatformWindows DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatchPlatform = "windows"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatchPlatformMac     DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatchPlatform = "mac"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatchPlatformLinux   DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatchPlatform = "linux"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatchPlatformAndroid DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatchPlatform = "android"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatchPlatformIos     DevicePostureDevicePostureRulesListDevicePostureRulesResponseMatchPlatform = "ios"
)

// The type of device posture rule.
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseType string

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeFile              DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "file"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeApplication       DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "application"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeTanium            DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "tanium"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeGateway           DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "gateway"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeWarp              DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "warp"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeDiskEncryption    DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "disk_encryption"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeSentinelone       DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "sentinelone"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeCarbonblack       DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "carbonblack"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeFirewall          DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "firewall"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeOsVersion         DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "os_version"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeDomainJoined      DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "domain_joined"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeClientCertificate DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "client_certificate"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeUniqueClientID    DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "unique_client_id"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeKolide            DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "kolide"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeTaniumS2s         DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "tanium_s2s"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeCrowdstrikeS2s    DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "crowdstrike_s2s"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeIntune            DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "intune"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeWorkspaceOne      DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "workspace_one"
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseTypeSentineloneS2s    DevicePostureDevicePostureRulesListDevicePostureRulesResponseType = "sentinelone_s2s"
)

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
	DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown215 DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown216 DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown217 DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown218 DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperatorUnknown219 DevicePostureGetResponseInputTeamsDevicesOsVersionInputRequestOperator = "=="
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
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown225 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown226 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown227 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown228 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown229 DevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
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
	DevicePostureUpdateParamsTypeWarp              DevicePostureUpdateParamsType = "warp"
	DevicePostureUpdateParamsTypeDiskEncryption    DevicePostureUpdateParamsType = "disk_encryption"
	DevicePostureUpdateParamsTypeSentinelone       DevicePostureUpdateParamsType = "sentinelone"
	DevicePostureUpdateParamsTypeCarbonblack       DevicePostureUpdateParamsType = "carbonblack"
	DevicePostureUpdateParamsTypeFirewall          DevicePostureUpdateParamsType = "firewall"
	DevicePostureUpdateParamsTypeOsVersion         DevicePostureUpdateParamsType = "os_version"
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
// [DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequest],
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

type DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// Operator
	Operator param.Field[DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperator] `json:"operator,required"`
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

func (r DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequest) implementsDevicePostureUpdateParamsInput() {
}

// Operating System
type DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperatingSystem string

const (
	DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperatingSystemWindows DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown285 DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown286 DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown287 DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown288 DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown289 DevicePostureUpdateParamsInputTeamsDevicesOsVersionInputRequestOperator = "=="
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
	// Operator
	Operator param.Field[DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator] `json:"operator"`
	// Os Version
	Os param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
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
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown295 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown296 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown297 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown298 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown299 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown305 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown306 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown307 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown308 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown309 DevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
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
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown315 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown316 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown317 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown318 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown319 DevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
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
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown325 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown326 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown327 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown328 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown329 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
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
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown335 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown336 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown337 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown338 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown339 DevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
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
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown345 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown346 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown347 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown348 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown349 DevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
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

type DevicePostureDevicePostureRulesNewDevicePostureRuleParams struct {
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatch] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeFile              DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "file"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeApplication       DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "application"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeTanium            DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "tanium"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeGateway           DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "gateway"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeWarp              DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "warp"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeDiskEncryption    DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "disk_encryption"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeSentinelone       DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "sentinelone"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeCarbonblack       DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "carbonblack"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeFirewall          DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "firewall"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeOsVersion         DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "os_version"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeDomainJoined      DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "domain_joined"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeClientCertificate DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "client_certificate"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeUniqueClientID    DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "unique_client_id"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeKolide            DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "kolide"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeTaniumS2s         DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "tanium_s2s"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeCrowdstrikeS2s    DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "crowdstrike_s2s"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeIntune            DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "intune"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeWorkspaceOne      DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "workspace_one"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeSentineloneS2s    DevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "sentinelone_s2s"
)

// The value to be checked against.
//
// Satisfied by
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFileInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesUniqueClientIDInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesDomainJoinedInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFirewallInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCarbonblackInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesDiskEncryptionInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesApplicationInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesClientCertificateInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesWorkspaceOneInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequest],
// [DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequest].
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput interface {
	implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput()
}

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFileInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFileInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFileInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating system
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFileInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux   DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFileInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesUniqueClientIDInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating System
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesDomainJoinedInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating System
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// Operator
	Operator param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperator] `json:"operator,required"`
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

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating System
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown355 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown356 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown357 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown358 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperatorUnknown359 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesOsVersionInputRequestOperator = "=="
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFirewallInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating System
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating system
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCarbonblackInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCarbonblackInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating system
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesDiskEncryptionInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesApplicationInputRequestOperatingSystem] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesApplicationInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesApplicationInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating system
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesClientCertificateInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesWorkspaceOneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesWorkspaceOneInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Compliance Status
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Operator
	Operator param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestOperator] `json:"operator"`
	// Os Version
	Os param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator] `json:"versionOperator"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operator
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown365 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown366 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown367 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown368 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown369 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown375 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown376 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown377 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown378 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown379 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Compliance Status
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatusError         DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequestCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Count Operator
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown385 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown386 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown387 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown388 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown389 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown395 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown396 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown397 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown398 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown399 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow      DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown405 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown406 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown407 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown408 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown409 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus] `json:"network_status"`
	// operator
	Operator param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestOperator] `json:"operator"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequest) implementsDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Network status of device.
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown415 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown416 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown417 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown418 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown419 DevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatch struct {
	Platform param.Field[DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform] `json:"platform"`
}

func (r DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform string

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformWindows DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "windows"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformMac     DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "mac"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformLinux   DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "linux"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformAndroid DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "android"
	DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformIos     DevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "ios"
)

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelope struct {
	Errors   []DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureDevicePostureRulesNewDevicePostureRuleResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeJSON    `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelope]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeErrors struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    devicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeErrors]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeMessages struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    devicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeMessages]
type devicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeSuccess bool

const (
	DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeSuccessTrue DevicePostureDevicePostureRulesNewDevicePostureRuleResponseEnvelopeSuccess = true
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelope struct {
	Errors   []DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePostureDevicePostureRulesListDevicePostureRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeJSON       `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelope]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeErrors struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    devicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeErrors]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeMessages struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    devicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeMessages]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeSuccess bool

const (
	DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeSuccessTrue DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeSuccess = true
)

type DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                             `json:"total_count"`
	JSON       devicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeResultInfo]
type devicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDevicePostureRulesListDevicePostureRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
