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

// ZeroTrustDevicePostureService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDevicePostureService]
// method instead.
type ZeroTrustDevicePostureService struct {
	Options      []option.RequestOption
	Integrations *ZeroTrustDevicePostureIntegrationService
}

// NewZeroTrustDevicePostureService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDevicePostureService(opts ...option.RequestOption) (r *ZeroTrustDevicePostureService) {
	r = &ZeroTrustDevicePostureService{}
	r.Options = opts
	r.Integrations = NewZeroTrustDevicePostureIntegrationService(opts...)
	return
}

// Creates a new device posture rule.
func (r *ZeroTrustDevicePostureService) New(ctx context.Context, params ZeroTrustDevicePostureNewParams, opts ...option.RequestOption) (res *ZeroTrustDevicePostureNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePostureNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a device posture rule.
func (r *ZeroTrustDevicePostureService) Update(ctx context.Context, ruleID string, params ZeroTrustDevicePostureUpdateParams, opts ...option.RequestOption) (res *ZeroTrustDevicePostureUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePostureUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches device posture rules for a Zero Trust account.
func (r *ZeroTrustDevicePostureService) List(ctx context.Context, query ZeroTrustDevicePostureListParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePostureListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePostureListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a device posture rule.
func (r *ZeroTrustDevicePostureService) Delete(ctx context.Context, ruleID string, body ZeroTrustDevicePostureDeleteParams, opts ...option.RequestOption) (res *ZeroTrustDevicePostureDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePostureDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single device posture rule.
func (r *ZeroTrustDevicePostureService) Get(ctx context.Context, ruleID string, query ZeroTrustDevicePostureGetParams, opts ...option.RequestOption) (res *ZeroTrustDevicePostureGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePostureGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDevicePostureNewResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input ZeroTrustDevicePostureNewResponseInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []ZeroTrustDevicePostureNewResponseMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type ZeroTrustDevicePostureNewResponseType `json:"type"`
	JSON zeroTrustDevicePostureNewResponseJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponse]
type zeroTrustDevicePostureNewResponseJSON struct {
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

func (r *ZeroTrustDevicePostureNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesClientCertificateInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequest],
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequest] or
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequest].
type ZeroTrustDevicePostureNewResponseInput interface {
	implementsZeroTrustDevicePostureNewResponseInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustDevicePostureNewResponseInput)(nil)).Elem(), "")
}

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                 `json:"thumbprint"`
	JSON       zeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequestJSON contains
// the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystemLinux   ZeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystemMac     ZeroTrustDevicePostureNewResponseInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem ZeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            zeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  ZeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      ZeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos ZeroTrustDevicePostureNewResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem ZeroTrustDevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                         `json:"domain"`
	JSON   zeroTrustDevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                                                      `json:"os_version_extra"`
	JSON           zeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown5 ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator = "<"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown6 ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator = "<="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown7 ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator = ">"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown8 ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator = ">="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown9 ZeroTrustDevicePostureNewResponseInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem ZeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            zeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequestOperatingSystemMac     ZeroTrustDevicePostureNewResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                        `json:"thumbprint"`
	JSON       zeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                        `json:"thumbprint"`
	JSON       zeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   ZeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     ZeroTrustDevicePostureNewResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                             `json:"requireAll"`
	JSON       zeroTrustDevicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                        `json:"thumbprint"`
	JSON       zeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   ZeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystemMac     ZeroTrustDevicePostureNewResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                              `json:"cn,required"`
	JSON zeroTrustDevicePostureNewResponseInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesClientCertificateInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus ZeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                         `json:"connection_id,required"`
	JSON         zeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Compliance Status
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    ZeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant ZeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      ZeroTrustDevicePostureNewResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Operator
	Operator ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            zeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
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

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Operator
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown15 ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown16 ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown17 ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown18 ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown19 ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestStateOnline  ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestStateOffline ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestStateUnknown ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown25 ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown26 ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown27 ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown28 ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown29 ZeroTrustDevicePostureNewResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         zeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Compliance Status
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatusError         ZeroTrustDevicePostureNewResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                                   `json:"issue_count,required"`
	JSON       zeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Count Operator
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown35 ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown36 ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown37 ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown38 ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown39 ZeroTrustDevicePostureNewResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                                  `json:"total_score"`
	JSON       zeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown45 ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown46 ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown47 ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown48 ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown49 ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevelLow      ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevelMedium   ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevelHigh     ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevelCritical ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown55 ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown56 ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown57 ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown58 ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown59 ZeroTrustDevicePostureNewResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     zeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// zeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequest]
type zeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureNewResponseInput() {
}

// Network status of device.
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown65 ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown66 ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown67 ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown68 ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown69 ZeroTrustDevicePostureNewResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type ZeroTrustDevicePostureNewResponseMatch struct {
	Platform ZeroTrustDevicePostureNewResponseMatchPlatform `json:"platform"`
	JSON     zeroTrustDevicePostureNewResponseMatchJSON     `json:"-"`
}

// zeroTrustDevicePostureNewResponseMatchJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePostureNewResponseMatch]
type zeroTrustDevicePostureNewResponseMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureNewResponseMatchPlatform string

const (
	ZeroTrustDevicePostureNewResponseMatchPlatformWindows ZeroTrustDevicePostureNewResponseMatchPlatform = "windows"
	ZeroTrustDevicePostureNewResponseMatchPlatformMac     ZeroTrustDevicePostureNewResponseMatchPlatform = "mac"
	ZeroTrustDevicePostureNewResponseMatchPlatformLinux   ZeroTrustDevicePostureNewResponseMatchPlatform = "linux"
	ZeroTrustDevicePostureNewResponseMatchPlatformAndroid ZeroTrustDevicePostureNewResponseMatchPlatform = "android"
	ZeroTrustDevicePostureNewResponseMatchPlatformIos     ZeroTrustDevicePostureNewResponseMatchPlatform = "ios"
)

// The type of device posture rule.
type ZeroTrustDevicePostureNewResponseType string

const (
	ZeroTrustDevicePostureNewResponseTypeFile              ZeroTrustDevicePostureNewResponseType = "file"
	ZeroTrustDevicePostureNewResponseTypeApplication       ZeroTrustDevicePostureNewResponseType = "application"
	ZeroTrustDevicePostureNewResponseTypeTanium            ZeroTrustDevicePostureNewResponseType = "tanium"
	ZeroTrustDevicePostureNewResponseTypeGateway           ZeroTrustDevicePostureNewResponseType = "gateway"
	ZeroTrustDevicePostureNewResponseTypeWARP              ZeroTrustDevicePostureNewResponseType = "warp"
	ZeroTrustDevicePostureNewResponseTypeDiskEncryption    ZeroTrustDevicePostureNewResponseType = "disk_encryption"
	ZeroTrustDevicePostureNewResponseTypeSentinelone       ZeroTrustDevicePostureNewResponseType = "sentinelone"
	ZeroTrustDevicePostureNewResponseTypeCarbonblack       ZeroTrustDevicePostureNewResponseType = "carbonblack"
	ZeroTrustDevicePostureNewResponseTypeFirewall          ZeroTrustDevicePostureNewResponseType = "firewall"
	ZeroTrustDevicePostureNewResponseTypeOSVersion         ZeroTrustDevicePostureNewResponseType = "os_version"
	ZeroTrustDevicePostureNewResponseTypeDomainJoined      ZeroTrustDevicePostureNewResponseType = "domain_joined"
	ZeroTrustDevicePostureNewResponseTypeClientCertificate ZeroTrustDevicePostureNewResponseType = "client_certificate"
	ZeroTrustDevicePostureNewResponseTypeUniqueClientID    ZeroTrustDevicePostureNewResponseType = "unique_client_id"
	ZeroTrustDevicePostureNewResponseTypeKolide            ZeroTrustDevicePostureNewResponseType = "kolide"
	ZeroTrustDevicePostureNewResponseTypeTaniumS2s         ZeroTrustDevicePostureNewResponseType = "tanium_s2s"
	ZeroTrustDevicePostureNewResponseTypeCrowdstrikeS2s    ZeroTrustDevicePostureNewResponseType = "crowdstrike_s2s"
	ZeroTrustDevicePostureNewResponseTypeIntune            ZeroTrustDevicePostureNewResponseType = "intune"
	ZeroTrustDevicePostureNewResponseTypeWorkspaceOne      ZeroTrustDevicePostureNewResponseType = "workspace_one"
	ZeroTrustDevicePostureNewResponseTypeSentineloneS2s    ZeroTrustDevicePostureNewResponseType = "sentinelone_s2s"
)

type ZeroTrustDevicePostureUpdateResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input ZeroTrustDevicePostureUpdateResponseInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []ZeroTrustDevicePostureUpdateResponseMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type ZeroTrustDevicePostureUpdateResponseType `json:"type"`
	JSON zeroTrustDevicePostureUpdateResponseJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePostureUpdateResponse]
type zeroTrustDevicePostureUpdateResponseJSON struct {
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

func (r *ZeroTrustDevicePostureUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequest],
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequest] or
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequest].
type ZeroTrustDevicePostureUpdateResponseInput interface {
	implementsZeroTrustDevicePostureUpdateResponseInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustDevicePostureUpdateResponseInput)(nil)).Elem(), "")
}

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                    `json:"thumbprint"`
	JSON       zeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystemLinux   ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystemMac     ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            zeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                            `json:"domain"`
	JSON   zeroTrustDevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                                                         `json:"os_version_extra"`
	JSON           zeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown75 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator = "<"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown76 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator = "<="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown77 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator = ">"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown78 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator = ">="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown79 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            zeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestOperatingSystemMac     ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                           `json:"thumbprint"`
	JSON       zeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                           `json:"thumbprint"`
	JSON       zeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                                `json:"requireAll"`
	JSON       zeroTrustDevicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                           `json:"thumbprint"`
	JSON       zeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystemMac     ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                                 `json:"cn,required"`
	JSON zeroTrustDevicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                            `json:"connection_id,required"`
	JSON         zeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Compliance Status
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Operator
	Operator ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            zeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
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

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Operator
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown85 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown86 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown87 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown88 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown89 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestStateOnline  ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestStateOffline ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestStateUnknown ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown95 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown96 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown97 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown98 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown99 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                      `json:"connection_id,required"`
	JSON         zeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Compliance Status
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatusError         ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                                      `json:"issue_count,required"`
	JSON       zeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Count Operator
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown105 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown106 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown107 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown108 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown109 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                                     `json:"total_score"`
	JSON       zeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown115 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown116 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown117 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown118 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown119 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevelLow      ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevelMedium   ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevelHigh     ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevelCritical ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown125 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown126 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown127 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown128 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown129 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     zeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequest]
type zeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureUpdateResponseInput() {
}

// Network status of device.
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown135 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown136 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown137 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown138 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown139 ZeroTrustDevicePostureUpdateResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type ZeroTrustDevicePostureUpdateResponseMatch struct {
	Platform ZeroTrustDevicePostureUpdateResponseMatchPlatform `json:"platform"`
	JSON     zeroTrustDevicePostureUpdateResponseMatchJSON     `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseMatchJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePostureUpdateResponseMatch]
type zeroTrustDevicePostureUpdateResponseMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureUpdateResponseMatchPlatform string

const (
	ZeroTrustDevicePostureUpdateResponseMatchPlatformWindows ZeroTrustDevicePostureUpdateResponseMatchPlatform = "windows"
	ZeroTrustDevicePostureUpdateResponseMatchPlatformMac     ZeroTrustDevicePostureUpdateResponseMatchPlatform = "mac"
	ZeroTrustDevicePostureUpdateResponseMatchPlatformLinux   ZeroTrustDevicePostureUpdateResponseMatchPlatform = "linux"
	ZeroTrustDevicePostureUpdateResponseMatchPlatformAndroid ZeroTrustDevicePostureUpdateResponseMatchPlatform = "android"
	ZeroTrustDevicePostureUpdateResponseMatchPlatformIos     ZeroTrustDevicePostureUpdateResponseMatchPlatform = "ios"
)

// The type of device posture rule.
type ZeroTrustDevicePostureUpdateResponseType string

const (
	ZeroTrustDevicePostureUpdateResponseTypeFile              ZeroTrustDevicePostureUpdateResponseType = "file"
	ZeroTrustDevicePostureUpdateResponseTypeApplication       ZeroTrustDevicePostureUpdateResponseType = "application"
	ZeroTrustDevicePostureUpdateResponseTypeTanium            ZeroTrustDevicePostureUpdateResponseType = "tanium"
	ZeroTrustDevicePostureUpdateResponseTypeGateway           ZeroTrustDevicePostureUpdateResponseType = "gateway"
	ZeroTrustDevicePostureUpdateResponseTypeWARP              ZeroTrustDevicePostureUpdateResponseType = "warp"
	ZeroTrustDevicePostureUpdateResponseTypeDiskEncryption    ZeroTrustDevicePostureUpdateResponseType = "disk_encryption"
	ZeroTrustDevicePostureUpdateResponseTypeSentinelone       ZeroTrustDevicePostureUpdateResponseType = "sentinelone"
	ZeroTrustDevicePostureUpdateResponseTypeCarbonblack       ZeroTrustDevicePostureUpdateResponseType = "carbonblack"
	ZeroTrustDevicePostureUpdateResponseTypeFirewall          ZeroTrustDevicePostureUpdateResponseType = "firewall"
	ZeroTrustDevicePostureUpdateResponseTypeOSVersion         ZeroTrustDevicePostureUpdateResponseType = "os_version"
	ZeroTrustDevicePostureUpdateResponseTypeDomainJoined      ZeroTrustDevicePostureUpdateResponseType = "domain_joined"
	ZeroTrustDevicePostureUpdateResponseTypeClientCertificate ZeroTrustDevicePostureUpdateResponseType = "client_certificate"
	ZeroTrustDevicePostureUpdateResponseTypeUniqueClientID    ZeroTrustDevicePostureUpdateResponseType = "unique_client_id"
	ZeroTrustDevicePostureUpdateResponseTypeKolide            ZeroTrustDevicePostureUpdateResponseType = "kolide"
	ZeroTrustDevicePostureUpdateResponseTypeTaniumS2s         ZeroTrustDevicePostureUpdateResponseType = "tanium_s2s"
	ZeroTrustDevicePostureUpdateResponseTypeCrowdstrikeS2s    ZeroTrustDevicePostureUpdateResponseType = "crowdstrike_s2s"
	ZeroTrustDevicePostureUpdateResponseTypeIntune            ZeroTrustDevicePostureUpdateResponseType = "intune"
	ZeroTrustDevicePostureUpdateResponseTypeWorkspaceOne      ZeroTrustDevicePostureUpdateResponseType = "workspace_one"
	ZeroTrustDevicePostureUpdateResponseTypeSentineloneS2s    ZeroTrustDevicePostureUpdateResponseType = "sentinelone_s2s"
)

type ZeroTrustDevicePostureListResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input ZeroTrustDevicePostureListResponseInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []ZeroTrustDevicePostureListResponseMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type ZeroTrustDevicePostureListResponseType `json:"type"`
	JSON zeroTrustDevicePostureListResponseJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponse]
type zeroTrustDevicePostureListResponseJSON struct {
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

func (r *ZeroTrustDevicePostureListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesClientCertificateInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequest],
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequest] or
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequest].
type ZeroTrustDevicePostureListResponseInput interface {
	implementsZeroTrustDevicePostureListResponseInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustDevicePostureListResponseInput)(nil)).Elem(), "")
}

type ZeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                  `json:"thumbprint"`
	JSON       zeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequestJSON contains
// the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystemWindows ZeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystemLinux   ZeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystemMac     ZeroTrustDevicePostureListResponseInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem ZeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            zeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  ZeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      ZeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos ZeroTrustDevicePostureListResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type ZeroTrustDevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem ZeroTrustDevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                          `json:"domain"`
	JSON   zeroTrustDevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows ZeroTrustDevicePostureListResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                                                       `json:"os_version_extra"`
	JSON           zeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown145 ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator = "<"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown146 ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator = "<="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown147 ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator = ">"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown148 ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator = ">="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown149 ZeroTrustDevicePostureListResponseInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type ZeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem ZeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            zeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequestOperatingSystemWindows ZeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequestOperatingSystemMac     ZeroTrustDevicePostureListResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                         `json:"thumbprint"`
	JSON       zeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                         `json:"thumbprint"`
	JSON       zeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows ZeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   ZeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     ZeroTrustDevicePostureListResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                              `json:"requireAll"`
	JSON       zeroTrustDevicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

type ZeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                         `json:"thumbprint"`
	JSON       zeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystemWindows ZeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   ZeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystemMac     ZeroTrustDevicePostureListResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureListResponseInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                               `json:"cn,required"`
	JSON zeroTrustDevicePostureListResponseInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesClientCertificateInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

type ZeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus ZeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                          `json:"connection_id,required"`
	JSON         zeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Compliance Status
type ZeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    ZeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant ZeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      ZeroTrustDevicePostureListResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Operator
	Operator ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            zeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
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

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Operator
type ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown155 ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown156 ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown157 ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown158 ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown159 ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestStateOnline  ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestStateOffline ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestStateUnknown ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown165 ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown166 ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown167 ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown168 ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown169 ZeroTrustDevicePostureListResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                    `json:"connection_id,required"`
	JSON         zeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Compliance Status
type ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatusError         ZeroTrustDevicePostureListResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                                    `json:"issue_count,required"`
	JSON       zeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Count Operator
type ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown175 ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown176 ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown177 ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown178 ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown179 ZeroTrustDevicePostureListResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                                   `json:"total_score"`
	JSON       zeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown185 ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown186 ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown187 ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown188 ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown189 ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevelLow      ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevelMedium   ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevelHigh     ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevelCritical ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown195 ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown196 ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown197 ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown198 ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown199 ZeroTrustDevicePostureListResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     zeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// zeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequest]
type zeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureListResponseInput() {
}

// Network status of device.
type ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown205 ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown206 ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown207 ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown208 ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown209 ZeroTrustDevicePostureListResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type ZeroTrustDevicePostureListResponseMatch struct {
	Platform ZeroTrustDevicePostureListResponseMatchPlatform `json:"platform"`
	JSON     zeroTrustDevicePostureListResponseMatchJSON     `json:"-"`
}

// zeroTrustDevicePostureListResponseMatchJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePostureListResponseMatch]
type zeroTrustDevicePostureListResponseMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureListResponseMatchPlatform string

const (
	ZeroTrustDevicePostureListResponseMatchPlatformWindows ZeroTrustDevicePostureListResponseMatchPlatform = "windows"
	ZeroTrustDevicePostureListResponseMatchPlatformMac     ZeroTrustDevicePostureListResponseMatchPlatform = "mac"
	ZeroTrustDevicePostureListResponseMatchPlatformLinux   ZeroTrustDevicePostureListResponseMatchPlatform = "linux"
	ZeroTrustDevicePostureListResponseMatchPlatformAndroid ZeroTrustDevicePostureListResponseMatchPlatform = "android"
	ZeroTrustDevicePostureListResponseMatchPlatformIos     ZeroTrustDevicePostureListResponseMatchPlatform = "ios"
)

// The type of device posture rule.
type ZeroTrustDevicePostureListResponseType string

const (
	ZeroTrustDevicePostureListResponseTypeFile              ZeroTrustDevicePostureListResponseType = "file"
	ZeroTrustDevicePostureListResponseTypeApplication       ZeroTrustDevicePostureListResponseType = "application"
	ZeroTrustDevicePostureListResponseTypeTanium            ZeroTrustDevicePostureListResponseType = "tanium"
	ZeroTrustDevicePostureListResponseTypeGateway           ZeroTrustDevicePostureListResponseType = "gateway"
	ZeroTrustDevicePostureListResponseTypeWARP              ZeroTrustDevicePostureListResponseType = "warp"
	ZeroTrustDevicePostureListResponseTypeDiskEncryption    ZeroTrustDevicePostureListResponseType = "disk_encryption"
	ZeroTrustDevicePostureListResponseTypeSentinelone       ZeroTrustDevicePostureListResponseType = "sentinelone"
	ZeroTrustDevicePostureListResponseTypeCarbonblack       ZeroTrustDevicePostureListResponseType = "carbonblack"
	ZeroTrustDevicePostureListResponseTypeFirewall          ZeroTrustDevicePostureListResponseType = "firewall"
	ZeroTrustDevicePostureListResponseTypeOSVersion         ZeroTrustDevicePostureListResponseType = "os_version"
	ZeroTrustDevicePostureListResponseTypeDomainJoined      ZeroTrustDevicePostureListResponseType = "domain_joined"
	ZeroTrustDevicePostureListResponseTypeClientCertificate ZeroTrustDevicePostureListResponseType = "client_certificate"
	ZeroTrustDevicePostureListResponseTypeUniqueClientID    ZeroTrustDevicePostureListResponseType = "unique_client_id"
	ZeroTrustDevicePostureListResponseTypeKolide            ZeroTrustDevicePostureListResponseType = "kolide"
	ZeroTrustDevicePostureListResponseTypeTaniumS2s         ZeroTrustDevicePostureListResponseType = "tanium_s2s"
	ZeroTrustDevicePostureListResponseTypeCrowdstrikeS2s    ZeroTrustDevicePostureListResponseType = "crowdstrike_s2s"
	ZeroTrustDevicePostureListResponseTypeIntune            ZeroTrustDevicePostureListResponseType = "intune"
	ZeroTrustDevicePostureListResponseTypeWorkspaceOne      ZeroTrustDevicePostureListResponseType = "workspace_one"
	ZeroTrustDevicePostureListResponseTypeSentineloneS2s    ZeroTrustDevicePostureListResponseType = "sentinelone_s2s"
)

type ZeroTrustDevicePostureDeleteResponse struct {
	// API UUID.
	ID   string                                   `json:"id"`
	JSON zeroTrustDevicePostureDeleteResponseJSON `json:"-"`
}

// zeroTrustDevicePostureDeleteResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePostureDeleteResponse]
type zeroTrustDevicePostureDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureGetResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input ZeroTrustDevicePostureGetResponseInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []ZeroTrustDevicePostureGetResponseMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type ZeroTrustDevicePostureGetResponseType `json:"type"`
	JSON zeroTrustDevicePostureGetResponseJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponse]
type zeroTrustDevicePostureGetResponseJSON struct {
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

func (r *ZeroTrustDevicePostureGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesClientCertificateInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequest],
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequest] or
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequest].
type ZeroTrustDevicePostureGetResponseInput interface {
	implementsZeroTrustDevicePostureGetResponseInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustDevicePostureGetResponseInput)(nil)).Elem(), "")
}

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                 `json:"thumbprint"`
	JSON       zeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequestJSON contains
// the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystemWindows ZeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystemLinux   ZeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystemMac     ZeroTrustDevicePostureGetResponseInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem ZeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            zeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  ZeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      ZeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos ZeroTrustDevicePostureGetResponseInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem ZeroTrustDevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                         `json:"domain"`
	JSON   zeroTrustDevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows ZeroTrustDevicePostureGetResponseInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                                                      `json:"os_version_extra"`
	JSON           zeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown215 ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator = "<"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown216 ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator = "<="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown217 ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator = ">"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown218 ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator = ">="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperatorUnknown219 ZeroTrustDevicePostureGetResponseInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem ZeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            zeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Operating System
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequestOperatingSystemWindows ZeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequestOperatingSystemMac     ZeroTrustDevicePostureGetResponseInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                        `json:"thumbprint"`
	JSON       zeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                        `json:"thumbprint"`
	JSON       zeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows ZeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   ZeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     ZeroTrustDevicePostureGetResponseInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                             `json:"requireAll"`
	JSON       zeroTrustDevicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem ZeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                        `json:"thumbprint"`
	JSON       zeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Operating system
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystemWindows ZeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   ZeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystemMac     ZeroTrustDevicePostureGetResponseInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                              `json:"cn,required"`
	JSON zeroTrustDevicePostureGetResponseInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesClientCertificateInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus ZeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                         `json:"connection_id,required"`
	JSON         zeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Compliance Status
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    ZeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant ZeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      ZeroTrustDevicePostureGetResponseInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Operator
	Operator ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            zeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
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

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Operator
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown225 ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown226 ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown227 ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown228 ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown229 ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestStateOnline  ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestStateOffline ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestStateUnknown ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown235 ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown236 ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown237 ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown238 ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown239 ZeroTrustDevicePostureGetResponseInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         zeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Compliance Status
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatusError         ZeroTrustDevicePostureGetResponseInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                                   `json:"issue_count,required"`
	JSON       zeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Count Operator
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown245 ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = "<"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown246 ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = "<="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown247 ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = ">"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown248 ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = ">="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperatorUnknown249 ZeroTrustDevicePostureGetResponseInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                                  `json:"total_score"`
	JSON       zeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown255 ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = "<"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown256 ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = "<="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown257 ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = ">"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown258 ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = ">="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperatorUnknown259 ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevelLow      ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevelMedium   ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevelHigh     ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevelCritical ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown265 ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown266 ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown267 ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown268 ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown269 ZeroTrustDevicePostureGetResponseInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     zeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// zeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequest]
type zeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureGetResponseInput() {
}

// Network status of device.
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown275 ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown276 ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown277 ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown278 ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown279 ZeroTrustDevicePostureGetResponseInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type ZeroTrustDevicePostureGetResponseMatch struct {
	Platform ZeroTrustDevicePostureGetResponseMatchPlatform `json:"platform"`
	JSON     zeroTrustDevicePostureGetResponseMatchJSON     `json:"-"`
}

// zeroTrustDevicePostureGetResponseMatchJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePostureGetResponseMatch]
type zeroTrustDevicePostureGetResponseMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureGetResponseMatchPlatform string

const (
	ZeroTrustDevicePostureGetResponseMatchPlatformWindows ZeroTrustDevicePostureGetResponseMatchPlatform = "windows"
	ZeroTrustDevicePostureGetResponseMatchPlatformMac     ZeroTrustDevicePostureGetResponseMatchPlatform = "mac"
	ZeroTrustDevicePostureGetResponseMatchPlatformLinux   ZeroTrustDevicePostureGetResponseMatchPlatform = "linux"
	ZeroTrustDevicePostureGetResponseMatchPlatformAndroid ZeroTrustDevicePostureGetResponseMatchPlatform = "android"
	ZeroTrustDevicePostureGetResponseMatchPlatformIos     ZeroTrustDevicePostureGetResponseMatchPlatform = "ios"
)

// The type of device posture rule.
type ZeroTrustDevicePostureGetResponseType string

const (
	ZeroTrustDevicePostureGetResponseTypeFile              ZeroTrustDevicePostureGetResponseType = "file"
	ZeroTrustDevicePostureGetResponseTypeApplication       ZeroTrustDevicePostureGetResponseType = "application"
	ZeroTrustDevicePostureGetResponseTypeTanium            ZeroTrustDevicePostureGetResponseType = "tanium"
	ZeroTrustDevicePostureGetResponseTypeGateway           ZeroTrustDevicePostureGetResponseType = "gateway"
	ZeroTrustDevicePostureGetResponseTypeWARP              ZeroTrustDevicePostureGetResponseType = "warp"
	ZeroTrustDevicePostureGetResponseTypeDiskEncryption    ZeroTrustDevicePostureGetResponseType = "disk_encryption"
	ZeroTrustDevicePostureGetResponseTypeSentinelone       ZeroTrustDevicePostureGetResponseType = "sentinelone"
	ZeroTrustDevicePostureGetResponseTypeCarbonblack       ZeroTrustDevicePostureGetResponseType = "carbonblack"
	ZeroTrustDevicePostureGetResponseTypeFirewall          ZeroTrustDevicePostureGetResponseType = "firewall"
	ZeroTrustDevicePostureGetResponseTypeOSVersion         ZeroTrustDevicePostureGetResponseType = "os_version"
	ZeroTrustDevicePostureGetResponseTypeDomainJoined      ZeroTrustDevicePostureGetResponseType = "domain_joined"
	ZeroTrustDevicePostureGetResponseTypeClientCertificate ZeroTrustDevicePostureGetResponseType = "client_certificate"
	ZeroTrustDevicePostureGetResponseTypeUniqueClientID    ZeroTrustDevicePostureGetResponseType = "unique_client_id"
	ZeroTrustDevicePostureGetResponseTypeKolide            ZeroTrustDevicePostureGetResponseType = "kolide"
	ZeroTrustDevicePostureGetResponseTypeTaniumS2s         ZeroTrustDevicePostureGetResponseType = "tanium_s2s"
	ZeroTrustDevicePostureGetResponseTypeCrowdstrikeS2s    ZeroTrustDevicePostureGetResponseType = "crowdstrike_s2s"
	ZeroTrustDevicePostureGetResponseTypeIntune            ZeroTrustDevicePostureGetResponseType = "intune"
	ZeroTrustDevicePostureGetResponseTypeWorkspaceOne      ZeroTrustDevicePostureGetResponseType = "workspace_one"
	ZeroTrustDevicePostureGetResponseTypeSentineloneS2s    ZeroTrustDevicePostureGetResponseType = "sentinelone_s2s"
)

type ZeroTrustDevicePostureNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[ZeroTrustDevicePostureNewParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[ZeroTrustDevicePostureNewParamsInput] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]ZeroTrustDevicePostureNewParamsMatch] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r ZeroTrustDevicePostureNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type ZeroTrustDevicePostureNewParamsType string

const (
	ZeroTrustDevicePostureNewParamsTypeFile              ZeroTrustDevicePostureNewParamsType = "file"
	ZeroTrustDevicePostureNewParamsTypeApplication       ZeroTrustDevicePostureNewParamsType = "application"
	ZeroTrustDevicePostureNewParamsTypeTanium            ZeroTrustDevicePostureNewParamsType = "tanium"
	ZeroTrustDevicePostureNewParamsTypeGateway           ZeroTrustDevicePostureNewParamsType = "gateway"
	ZeroTrustDevicePostureNewParamsTypeWARP              ZeroTrustDevicePostureNewParamsType = "warp"
	ZeroTrustDevicePostureNewParamsTypeDiskEncryption    ZeroTrustDevicePostureNewParamsType = "disk_encryption"
	ZeroTrustDevicePostureNewParamsTypeSentinelone       ZeroTrustDevicePostureNewParamsType = "sentinelone"
	ZeroTrustDevicePostureNewParamsTypeCarbonblack       ZeroTrustDevicePostureNewParamsType = "carbonblack"
	ZeroTrustDevicePostureNewParamsTypeFirewall          ZeroTrustDevicePostureNewParamsType = "firewall"
	ZeroTrustDevicePostureNewParamsTypeOSVersion         ZeroTrustDevicePostureNewParamsType = "os_version"
	ZeroTrustDevicePostureNewParamsTypeDomainJoined      ZeroTrustDevicePostureNewParamsType = "domain_joined"
	ZeroTrustDevicePostureNewParamsTypeClientCertificate ZeroTrustDevicePostureNewParamsType = "client_certificate"
	ZeroTrustDevicePostureNewParamsTypeUniqueClientID    ZeroTrustDevicePostureNewParamsType = "unique_client_id"
	ZeroTrustDevicePostureNewParamsTypeKolide            ZeroTrustDevicePostureNewParamsType = "kolide"
	ZeroTrustDevicePostureNewParamsTypeTaniumS2s         ZeroTrustDevicePostureNewParamsType = "tanium_s2s"
	ZeroTrustDevicePostureNewParamsTypeCrowdstrikeS2s    ZeroTrustDevicePostureNewParamsType = "crowdstrike_s2s"
	ZeroTrustDevicePostureNewParamsTypeIntune            ZeroTrustDevicePostureNewParamsType = "intune"
	ZeroTrustDevicePostureNewParamsTypeWorkspaceOne      ZeroTrustDevicePostureNewParamsType = "workspace_one"
	ZeroTrustDevicePostureNewParamsTypeSentineloneS2s    ZeroTrustDevicePostureNewParamsType = "sentinelone_s2s"
)

// The value to be checked against.
//
// Satisfied by [ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesFirewallInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesApplicationInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequest],
// [ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest].
type ZeroTrustDevicePostureNewParamsInput interface {
	implementsZeroTrustDevicePostureNewParamsInput()
}

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating system
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux   ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemMac     ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating System
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  ZeroTrustDevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      ZeroTrustDevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos ZeroTrustDevicePostureNewParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating System
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// Operator
	Operator param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator] `json:"operator,required"`
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

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating System
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown285 ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown286 ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown287 ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown288 ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown289 ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating System
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac     ZeroTrustDevicePostureNewParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating system
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating system
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   ZeroTrustDevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     ZeroTrustDevicePostureNewParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesApplicationInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operating system
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemWindows ZeroTrustDevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   ZeroTrustDevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystemMac     ZeroTrustDevicePostureNewParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Compliance Status
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    ZeroTrustDevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant ZeroTrustDevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      ZeroTrustDevicePostureNewParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// Operator
	Operator param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator] `json:"operator"`
	// Os Version
	OS param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState] `json:"state"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator] `json:"versionOperator"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operator
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown295 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown296 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown297 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown298 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown299 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline  ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown305 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown306 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown307 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown308 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown309 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Compliance Status
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatusError         ZeroTrustDevicePostureNewParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Count Operator
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown315 ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown316 ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown317 ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown318 ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown319 ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown325 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown326 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown327 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown328 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown329 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow      ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium   ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh     ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown335 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown336 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown337 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown338 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown339 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus] `json:"network_status"`
	// operator
	Operator param.Field[ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator] `json:"operator"`
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureNewParamsInput() {
}

// Network status of device.
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown345 ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown346 ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown347 ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown348 ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown349 ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type ZeroTrustDevicePostureNewParamsMatch struct {
	Platform param.Field[ZeroTrustDevicePostureNewParamsMatchPlatform] `json:"platform"`
}

func (r ZeroTrustDevicePostureNewParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDevicePostureNewParamsMatchPlatform string

const (
	ZeroTrustDevicePostureNewParamsMatchPlatformWindows ZeroTrustDevicePostureNewParamsMatchPlatform = "windows"
	ZeroTrustDevicePostureNewParamsMatchPlatformMac     ZeroTrustDevicePostureNewParamsMatchPlatform = "mac"
	ZeroTrustDevicePostureNewParamsMatchPlatformLinux   ZeroTrustDevicePostureNewParamsMatchPlatform = "linux"
	ZeroTrustDevicePostureNewParamsMatchPlatformAndroid ZeroTrustDevicePostureNewParamsMatchPlatform = "android"
	ZeroTrustDevicePostureNewParamsMatchPlatformIos     ZeroTrustDevicePostureNewParamsMatchPlatform = "ios"
)

type ZeroTrustDevicePostureNewResponseEnvelope struct {
	Errors   []ZeroTrustDevicePostureNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePostureNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDevicePostureNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDevicePostureNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDevicePostureNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDevicePostureNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePostureNewResponseEnvelope]
type zeroTrustDevicePostureNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureNewResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDevicePostureNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePostureNewResponseEnvelopeErrors]
type zeroTrustDevicePostureNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureNewResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustDevicePostureNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePostureNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePostureNewResponseEnvelopeMessages]
type zeroTrustDevicePostureNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePostureNewResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePostureNewResponseEnvelopeSuccessTrue ZeroTrustDevicePostureNewResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePostureUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[ZeroTrustDevicePostureUpdateParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[ZeroTrustDevicePostureUpdateParamsInput] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]ZeroTrustDevicePostureUpdateParamsMatch] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r ZeroTrustDevicePostureUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type ZeroTrustDevicePostureUpdateParamsType string

const (
	ZeroTrustDevicePostureUpdateParamsTypeFile              ZeroTrustDevicePostureUpdateParamsType = "file"
	ZeroTrustDevicePostureUpdateParamsTypeApplication       ZeroTrustDevicePostureUpdateParamsType = "application"
	ZeroTrustDevicePostureUpdateParamsTypeTanium            ZeroTrustDevicePostureUpdateParamsType = "tanium"
	ZeroTrustDevicePostureUpdateParamsTypeGateway           ZeroTrustDevicePostureUpdateParamsType = "gateway"
	ZeroTrustDevicePostureUpdateParamsTypeWARP              ZeroTrustDevicePostureUpdateParamsType = "warp"
	ZeroTrustDevicePostureUpdateParamsTypeDiskEncryption    ZeroTrustDevicePostureUpdateParamsType = "disk_encryption"
	ZeroTrustDevicePostureUpdateParamsTypeSentinelone       ZeroTrustDevicePostureUpdateParamsType = "sentinelone"
	ZeroTrustDevicePostureUpdateParamsTypeCarbonblack       ZeroTrustDevicePostureUpdateParamsType = "carbonblack"
	ZeroTrustDevicePostureUpdateParamsTypeFirewall          ZeroTrustDevicePostureUpdateParamsType = "firewall"
	ZeroTrustDevicePostureUpdateParamsTypeOSVersion         ZeroTrustDevicePostureUpdateParamsType = "os_version"
	ZeroTrustDevicePostureUpdateParamsTypeDomainJoined      ZeroTrustDevicePostureUpdateParamsType = "domain_joined"
	ZeroTrustDevicePostureUpdateParamsTypeClientCertificate ZeroTrustDevicePostureUpdateParamsType = "client_certificate"
	ZeroTrustDevicePostureUpdateParamsTypeUniqueClientID    ZeroTrustDevicePostureUpdateParamsType = "unique_client_id"
	ZeroTrustDevicePostureUpdateParamsTypeKolide            ZeroTrustDevicePostureUpdateParamsType = "kolide"
	ZeroTrustDevicePostureUpdateParamsTypeTaniumS2s         ZeroTrustDevicePostureUpdateParamsType = "tanium_s2s"
	ZeroTrustDevicePostureUpdateParamsTypeCrowdstrikeS2s    ZeroTrustDevicePostureUpdateParamsType = "crowdstrike_s2s"
	ZeroTrustDevicePostureUpdateParamsTypeIntune            ZeroTrustDevicePostureUpdateParamsType = "intune"
	ZeroTrustDevicePostureUpdateParamsTypeWorkspaceOne      ZeroTrustDevicePostureUpdateParamsType = "workspace_one"
	ZeroTrustDevicePostureUpdateParamsTypeSentineloneS2s    ZeroTrustDevicePostureUpdateParamsType = "sentinelone_s2s"
)

// The value to be checked against.
//
// Satisfied by
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest],
// [ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest].
type ZeroTrustDevicePostureUpdateParamsInput interface {
	implementsZeroTrustDevicePostureUpdateParamsInput()
}

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating system
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux   ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemMac     ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating System
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating System
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// Operator
	Operator param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator] `json:"operator,required"`
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

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating System
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown355 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown356 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown357 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown358 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown359 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating System
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystemMac     ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating system
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating system
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operating system
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemWindows ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystemMac     ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Compliance Status
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// Operator
	Operator param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator] `json:"operator"`
	// Os Version
	OS param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState] `json:"state"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator] `json:"versionOperator"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operator
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown365 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown366 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown367 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown368 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown369 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOnline  ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateOffline ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestStateUnknown ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown375 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown376 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown377 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown378 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown379 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Compliance Status
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatusError         ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Count Operator
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown385 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown386 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown387 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown388 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown389 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown395 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown396 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown397 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown398 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown399 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelLow      ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelMedium   ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelHigh     ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevelCritical ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown405 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown406 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown407 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown408 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown409 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus] `json:"network_status"`
	// operator
	Operator param.Field[ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator] `json:"operator"`
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDevicePostureUpdateParamsInput() {
}

// Network status of device.
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown415 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown416 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown417 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown418 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown419 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type ZeroTrustDevicePostureUpdateParamsMatch struct {
	Platform param.Field[ZeroTrustDevicePostureUpdateParamsMatchPlatform] `json:"platform"`
}

func (r ZeroTrustDevicePostureUpdateParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDevicePostureUpdateParamsMatchPlatform string

const (
	ZeroTrustDevicePostureUpdateParamsMatchPlatformWindows ZeroTrustDevicePostureUpdateParamsMatchPlatform = "windows"
	ZeroTrustDevicePostureUpdateParamsMatchPlatformMac     ZeroTrustDevicePostureUpdateParamsMatchPlatform = "mac"
	ZeroTrustDevicePostureUpdateParamsMatchPlatformLinux   ZeroTrustDevicePostureUpdateParamsMatchPlatform = "linux"
	ZeroTrustDevicePostureUpdateParamsMatchPlatformAndroid ZeroTrustDevicePostureUpdateParamsMatchPlatform = "android"
	ZeroTrustDevicePostureUpdateParamsMatchPlatformIos     ZeroTrustDevicePostureUpdateParamsMatchPlatform = "ios"
)

type ZeroTrustDevicePostureUpdateResponseEnvelope struct {
	Errors   []ZeroTrustDevicePostureUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePostureUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDevicePostureUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDevicePostureUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDevicePostureUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDevicePostureUpdateResponseEnvelope]
type zeroTrustDevicePostureUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDevicePostureUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureUpdateResponseEnvelopeErrors]
type zeroTrustDevicePostureUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustDevicePostureUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePostureUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureUpdateResponseEnvelopeMessages]
type zeroTrustDevicePostureUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePostureUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePostureUpdateResponseEnvelopeSuccessTrue ZeroTrustDevicePostureUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePostureListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePostureListResponseEnvelope struct {
	Errors   []ZeroTrustDevicePostureListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePostureListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePostureListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePostureListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePostureListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePostureListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePostureListResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDevicePostureListResponseEnvelope]
type zeroTrustDevicePostureListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureListResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustDevicePostureListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePostureListResponseEnvelopeErrors]
type zeroTrustDevicePostureListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureListResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDevicePostureListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureListResponseEnvelopeMessages]
type zeroTrustDevicePostureListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePostureListResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePostureListResponseEnvelopeSuccessTrue ZeroTrustDevicePostureListResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePostureListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       zeroTrustDevicePostureListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePostureListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureListResponseEnvelopeResultInfo]
type zeroTrustDevicePostureListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePostureDeleteResponseEnvelope struct {
	Errors   []ZeroTrustDevicePostureDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePostureDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDevicePostureDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDevicePostureDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDevicePostureDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDevicePostureDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDevicePostureDeleteResponseEnvelope]
type zeroTrustDevicePostureDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureDeleteResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDevicePostureDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePostureDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureDeleteResponseEnvelopeErrors]
type zeroTrustDevicePostureDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureDeleteResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustDevicePostureDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePostureDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureDeleteResponseEnvelopeMessages]
type zeroTrustDevicePostureDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePostureDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePostureDeleteResponseEnvelopeSuccessTrue ZeroTrustDevicePostureDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePostureGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePostureGetResponseEnvelope struct {
	Errors   []ZeroTrustDevicePostureGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePostureGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDevicePostureGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDevicePostureGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDevicePostureGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDevicePostureGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePostureGetResponseEnvelope]
type zeroTrustDevicePostureGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDevicePostureGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePostureGetResponseEnvelopeErrors]
type zeroTrustDevicePostureGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustDevicePostureGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePostureGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePostureGetResponseEnvelopeMessages]
type zeroTrustDevicePostureGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePostureGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePostureGetResponseEnvelopeSuccessTrue ZeroTrustDevicePostureGetResponseEnvelopeSuccess = true
)
