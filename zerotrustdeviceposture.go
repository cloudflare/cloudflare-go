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
func (r *ZeroTrustDevicePostureService) New(ctx context.Context, params ZeroTrustDevicePostureNewParams, opts ...option.RequestOption) (res *TeamsDevicesDevicePostureRules, err error) {
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
func (r *ZeroTrustDevicePostureService) Update(ctx context.Context, ruleID string, params ZeroTrustDevicePostureUpdateParams, opts ...option.RequestOption) (res *TeamsDevicesDevicePostureRules, err error) {
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
func (r *ZeroTrustDevicePostureService) List(ctx context.Context, query ZeroTrustDevicePostureListParams, opts ...option.RequestOption) (res *[]TeamsDevicesDevicePostureRules, err error) {
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
func (r *ZeroTrustDevicePostureService) Get(ctx context.Context, ruleID string, query ZeroTrustDevicePostureGetParams, opts ...option.RequestOption) (res *TeamsDevicesDevicePostureRules, err error) {
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

type TeamsDevicesDevicePostureRules struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input TeamsDevicesDevicePostureRulesInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []TeamsDevicesDevicePostureRulesMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type TeamsDevicesDevicePostureRulesType `json:"type"`
	JSON teamsDevicesDevicePostureRulesJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesJSON contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRules]
type teamsDevicesDevicePostureRulesJSON struct {
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

func (r *TeamsDevicesDevicePostureRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequest],
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequest] or
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest].
type TeamsDevicesDevicePostureRulesInput interface {
	implementsTeamsDevicesDevicePostureRulesInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*TeamsDevicesDevicePostureRulesInput)(nil)).Elem(), "")
}

type TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                              `json:"thumbprint"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestJSON contains the
// JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Operating system
type TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem = "windows"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystemLinux   TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem = "linux"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystemMac     TeamsDevicesDevicePostureRulesInputTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            teamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Operating System
type TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos TeamsDevicesDevicePostureRulesInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                      `json:"domain"`
	JSON   teamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Operating System
type TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                                                   `json:"os_version_extra"`
	JSON           teamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Operating System
type TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

// Operator
type TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorUnknown5 TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorUnknown6 TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorUnknown7 TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorUnknown8 TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperatorUnknown9 TeamsDevicesDevicePostureRulesInputTeamsDevicesOSVersionInputRequestOperator = "=="
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            teamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestJSON contains
// the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Operating System
type TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystemMac     TeamsDevicesDevicePostureRulesInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                     `json:"thumbprint"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Operating system
type TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystemLinux   TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystemMac     TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                     `json:"thumbprint"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Operating system
type TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     TeamsDevicesDevicePostureRulesInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                          `json:"requireAll"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesDiskEncryptionInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

type TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                     `json:"thumbprint"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Operating system
type TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystemWindows TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystemMac     TeamsDevicesDevicePostureRulesInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                           `json:"cn,required"`
	JSON teamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesClientCertificateInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

type TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                      `json:"connection_id,required"`
	JSON         teamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Compliance Status
type TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      TeamsDevicesDevicePostureRulesInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Operator
	Operator TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            teamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
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

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Operator
type TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown15 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown16 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown17 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown18 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown19 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

// For more details on state, please refer to the Crowdstrike documentation.
type TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateOnline  TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateOffline TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestStateUnknown TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

// Version Operator
type TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown25 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown26 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown27 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown28 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown29 TeamsDevicesDevicePostureRulesInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                `json:"connection_id,required"`
	JSON         teamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestJSON contains
// the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Compliance Status
type TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatusError         TeamsDevicesDevicePostureRulesInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                                `json:"issue_count,required"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestJSON contains
// the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Count Operator
type TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorUnknown35 TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorUnknown36 TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorUnknown37 TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorUnknown38 TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperatorUnknown39 TeamsDevicesDevicePostureRulesInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                               `json:"total_score"`
	JSON       teamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestJSON contains
// the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorUnknown45 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorUnknown46 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorUnknown47 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorUnknown48 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperatorUnknown49 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelLow      TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelMedium   TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelHigh     TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevelCritical TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown55 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown56 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown57 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown58 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown59 TeamsDevicesDevicePostureRulesInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

type TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest]
type teamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequest) implementsTeamsDevicesDevicePostureRulesInput() {
}

// Network status of device.
type TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown65 TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown66 TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown67 TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown68 TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown69 TeamsDevicesDevicePostureRulesInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

type TeamsDevicesDevicePostureRulesMatch struct {
	Platform TeamsDevicesDevicePostureRulesMatchPlatform `json:"platform"`
	JSON     teamsDevicesDevicePostureRulesMatchJSON     `json:"-"`
}

// teamsDevicesDevicePostureRulesMatchJSON contains the JSON metadata for the
// struct [TeamsDevicesDevicePostureRulesMatch]
type teamsDevicesDevicePostureRulesMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamsDevicesDevicePostureRulesMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamsDevicesDevicePostureRulesMatchPlatform string

const (
	TeamsDevicesDevicePostureRulesMatchPlatformWindows TeamsDevicesDevicePostureRulesMatchPlatform = "windows"
	TeamsDevicesDevicePostureRulesMatchPlatformMac     TeamsDevicesDevicePostureRulesMatchPlatform = "mac"
	TeamsDevicesDevicePostureRulesMatchPlatformLinux   TeamsDevicesDevicePostureRulesMatchPlatform = "linux"
	TeamsDevicesDevicePostureRulesMatchPlatformAndroid TeamsDevicesDevicePostureRulesMatchPlatform = "android"
	TeamsDevicesDevicePostureRulesMatchPlatformIos     TeamsDevicesDevicePostureRulesMatchPlatform = "ios"
)

// The type of device posture rule.
type TeamsDevicesDevicePostureRulesType string

const (
	TeamsDevicesDevicePostureRulesTypeFile              TeamsDevicesDevicePostureRulesType = "file"
	TeamsDevicesDevicePostureRulesTypeApplication       TeamsDevicesDevicePostureRulesType = "application"
	TeamsDevicesDevicePostureRulesTypeTanium            TeamsDevicesDevicePostureRulesType = "tanium"
	TeamsDevicesDevicePostureRulesTypeGateway           TeamsDevicesDevicePostureRulesType = "gateway"
	TeamsDevicesDevicePostureRulesTypeWARP              TeamsDevicesDevicePostureRulesType = "warp"
	TeamsDevicesDevicePostureRulesTypeDiskEncryption    TeamsDevicesDevicePostureRulesType = "disk_encryption"
	TeamsDevicesDevicePostureRulesTypeSentinelone       TeamsDevicesDevicePostureRulesType = "sentinelone"
	TeamsDevicesDevicePostureRulesTypeCarbonblack       TeamsDevicesDevicePostureRulesType = "carbonblack"
	TeamsDevicesDevicePostureRulesTypeFirewall          TeamsDevicesDevicePostureRulesType = "firewall"
	TeamsDevicesDevicePostureRulesTypeOSVersion         TeamsDevicesDevicePostureRulesType = "os_version"
	TeamsDevicesDevicePostureRulesTypeDomainJoined      TeamsDevicesDevicePostureRulesType = "domain_joined"
	TeamsDevicesDevicePostureRulesTypeClientCertificate TeamsDevicesDevicePostureRulesType = "client_certificate"
	TeamsDevicesDevicePostureRulesTypeUniqueClientID    TeamsDevicesDevicePostureRulesType = "unique_client_id"
	TeamsDevicesDevicePostureRulesTypeKolide            TeamsDevicesDevicePostureRulesType = "kolide"
	TeamsDevicesDevicePostureRulesTypeTaniumS2s         TeamsDevicesDevicePostureRulesType = "tanium_s2s"
	TeamsDevicesDevicePostureRulesTypeCrowdstrikeS2s    TeamsDevicesDevicePostureRulesType = "crowdstrike_s2s"
	TeamsDevicesDevicePostureRulesTypeIntune            TeamsDevicesDevicePostureRulesType = "intune"
	TeamsDevicesDevicePostureRulesTypeWorkspaceOne      TeamsDevicesDevicePostureRulesType = "workspace_one"
	TeamsDevicesDevicePostureRulesTypeSentineloneS2s    TeamsDevicesDevicePostureRulesType = "sentinelone_s2s"
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
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown75 ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown76 ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown77 ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown78 ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown79 ZeroTrustDevicePostureNewParamsInputTeamsDevicesOSVersionInputRequestOperator = "=="
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
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown85 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown86 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown87 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown88 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown89 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
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
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown95 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown96 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown97 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown98 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown99 ZeroTrustDevicePostureNewParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
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
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown105 ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown106 ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown107 ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown108 ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown109 ZeroTrustDevicePostureNewParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
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
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown115 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown116 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown117 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown118 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown119 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
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
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown125 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown126 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown127 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown128 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown129 ZeroTrustDevicePostureNewParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
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
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown135 ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown136 ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown137 ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown138 ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown139 ZeroTrustDevicePostureNewParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
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
	Result   TeamsDevicesDevicePostureRules                      `json:"result,required,nullable"`
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
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown145 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown146 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown147 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown148 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperatorUnknown149 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesOSVersionInputRequestOperator = "=="
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
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown155 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown156 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown157 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown158 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperatorUnknown159 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestOperator = "=="
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
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown165 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown166 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown167 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown168 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorUnknown169 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
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
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown175 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown176 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown177 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown178 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperatorUnknown179 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesKolideInputRequestCountOperator = "=="
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
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown185 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown186 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown187 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown188 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperatorUnknown189 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestOperator = "=="
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
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown195 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown196 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown197 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown198 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperatorUnknown199 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
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
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown205 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown206 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown207 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown208 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperatorUnknown209 ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesSentineloneS2sInputRequestOperator = "=="
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
	Result   TeamsDevicesDevicePostureRules                         `json:"result,required,nullable"`
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
	Result   []TeamsDevicesDevicePostureRules                     `json:"result,required,nullable"`
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
	Result   TeamsDevicesDevicePostureRules                      `json:"result,required,nullable"`
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
