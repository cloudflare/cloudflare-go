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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *DevicePostureService) New(ctx context.Context, params DevicePostureNewParams, opts ...option.RequestOption) (res *DevicePostureRule, err error) {
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
func (r *DevicePostureService) Update(ctx context.Context, ruleID string, params DevicePostureUpdateParams, opts ...option.RequestOption) (res *DevicePostureRule, err error) {
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
func (r *DevicePostureService) List(ctx context.Context, query DevicePostureListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DevicePostureRule], err error) {
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
func (r *DevicePostureService) ListAutoPaging(ctx context.Context, query DevicePostureListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DevicePostureRule] {
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
func (r *DevicePostureService) Get(ctx context.Context, ruleID string, query DevicePostureGetParams, opts ...option.RequestOption) (res *DevicePostureRule, err error) {
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

type DevicePostureRule struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input Input `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []MatchItem `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type DevicePostureRuleType `json:"type"`
	JSON devicePostureRuleJSON `json:"-"`
}

// devicePostureRuleJSON contains the JSON metadata for the struct
// [DevicePostureRule]
type devicePostureRuleJSON struct {
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

func (r *DevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRuleJSON) RawJSON() string {
	return r.raw
}

// The type of device posture rule.
type DevicePostureRuleType string

const (
	DevicePostureRuleTypeFile              DevicePostureRuleType = "file"
	DevicePostureRuleTypeApplication       DevicePostureRuleType = "application"
	DevicePostureRuleTypeTanium            DevicePostureRuleType = "tanium"
	DevicePostureRuleTypeGateway           DevicePostureRuleType = "gateway"
	DevicePostureRuleTypeWARP              DevicePostureRuleType = "warp"
	DevicePostureRuleTypeDiskEncryption    DevicePostureRuleType = "disk_encryption"
	DevicePostureRuleTypeSentinelone       DevicePostureRuleType = "sentinelone"
	DevicePostureRuleTypeCarbonblack       DevicePostureRuleType = "carbonblack"
	DevicePostureRuleTypeFirewall          DevicePostureRuleType = "firewall"
	DevicePostureRuleTypeOSVersion         DevicePostureRuleType = "os_version"
	DevicePostureRuleTypeDomainJoined      DevicePostureRuleType = "domain_joined"
	DevicePostureRuleTypeClientCertificate DevicePostureRuleType = "client_certificate"
	DevicePostureRuleTypeUniqueClientID    DevicePostureRuleType = "unique_client_id"
	DevicePostureRuleTypeKolide            DevicePostureRuleType = "kolide"
	DevicePostureRuleTypeTaniumS2s         DevicePostureRuleType = "tanium_s2s"
	DevicePostureRuleTypeCrowdstrikeS2s    DevicePostureRuleType = "crowdstrike_s2s"
	DevicePostureRuleTypeIntune            DevicePostureRuleType = "intune"
	DevicePostureRuleTypeWorkspaceOne      DevicePostureRuleType = "workspace_one"
	DevicePostureRuleTypeSentineloneS2s    DevicePostureRuleType = "sentinelone_s2s"
)

func (r DevicePostureRuleType) IsKnown() bool {
	switch r {
	case DevicePostureRuleTypeFile, DevicePostureRuleTypeApplication, DevicePostureRuleTypeTanium, DevicePostureRuleTypeGateway, DevicePostureRuleTypeWARP, DevicePostureRuleTypeDiskEncryption, DevicePostureRuleTypeSentinelone, DevicePostureRuleTypeCarbonblack, DevicePostureRuleTypeFirewall, DevicePostureRuleTypeOSVersion, DevicePostureRuleTypeDomainJoined, DevicePostureRuleTypeClientCertificate, DevicePostureRuleTypeUniqueClientID, DevicePostureRuleTypeKolide, DevicePostureRuleTypeTaniumS2s, DevicePostureRuleTypeCrowdstrikeS2s, DevicePostureRuleTypeIntune, DevicePostureRuleTypeWorkspaceOne, DevicePostureRuleTypeSentineloneS2s:
		return true
	}
	return false
}

// The value to be checked against.
type Input struct {
	// Whether or not file exists
	Exists bool `json:"exists"`
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system"`
	// File path.
	Path string `json:"path"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string `json:"thumbprint"`
	// List ID.
	ID string `json:"id"`
	// Domain
	Domain string `json:"domain"`
	// operator
	Operator UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 `json:"operator"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string `json:"os_version_extra"`
	// Version of OS
	Version string `json:"version"`
	// Enabled
	Enabled    bool        `json:"enabled"`
	CheckDisks interface{} `json:"checkDisks,required"`
	// Whether to check all disks for encryption.
	RequireAll bool `json:"requireAll"`
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id"`
	// Common Name that is protected by the certificate
	Cn string `json:"cn"`
	// Compliance Status
	ComplianceStatus InputComplianceStatus `json:"compliance_status"`
	// Posture Integration ID.
	ConnectionID string `json:"connection_id"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State InputState `json:"state"`
	// Version Operator
	VersionOperator InputVersionOperator `json:"versionOperator"`
	// Count Operator
	CountOperator InputCountOperator `json:"countOperator"`
	// The Number of Issues.
	IssueCount string `json:"issue_count"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel InputRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator InputScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64 `json:"total_score"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus InputNetworkStatus `json:"network_status"`
	JSON          inputJSON          `json:"-"`
	union         InputUnion
}

// inputJSON contains the JSON metadata for the struct [Input]
type inputJSON struct {
	Exists           apijson.Field
	OperatingSystem  apijson.Field
	Path             apijson.Field
	Sha256           apijson.Field
	Thumbprint       apijson.Field
	ID               apijson.Field
	Domain           apijson.Field
	Operator         apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	Version          apijson.Field
	Enabled          apijson.Field
	CheckDisks       apijson.Field
	RequireAll       apijson.Field
	CertificateID    apijson.Field
	Cn               apijson.Field
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	LastSeen         apijson.Field
	OS               apijson.Field
	Overall          apijson.Field
	SensorConfig     apijson.Field
	State            apijson.Field
	VersionOperator  apijson.Field
	CountOperator    apijson.Field
	IssueCount       apijson.Field
	EidLastSeen      apijson.Field
	RiskLevel        apijson.Field
	ScoreOperator    apijson.Field
	TotalScore       apijson.Field
	ActiveThreats    apijson.Field
	Infected         apijson.Field
	IsActive         apijson.Field
	NetworkStatus    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r inputJSON) RawJSON() string {
	return r.raw
}

func (r *Input) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r Input) AsUnion() InputUnion {
	return r.union
}

// The value to be checked against.
//
// Union satisfied by [zero_trust.InputTeamsDevicesFileInputRequest],
// [zero_trust.InputTeamsDevicesUniqueClientIDInputRequest],
// [zero_trust.InputTeamsDevicesDomainJoinedInputRequest],
// [zero_trust.InputTeamsDevicesOSVersionInputRequest],
// [zero_trust.InputTeamsDevicesFirewallInputRequest],
// [zero_trust.InputTeamsDevicesSentineloneInputRequest],
// [zero_trust.InputTeamsDevicesCarbonblackInputRequest],
// [zero_trust.InputTeamsDevicesDiskEncryptionInputRequest],
// [zero_trust.InputTeamsDevicesApplicationInputRequest],
// [zero_trust.InputTeamsDevicesClientCertificateInputRequest],
// [zero_trust.InputTeamsDevicesWorkspaceOneInputRequest],
// [zero_trust.InputTeamsDevicesCrowdstrikeInputRequest],
// [zero_trust.InputTeamsDevicesIntuneInputRequest],
// [zero_trust.InputTeamsDevicesKolideInputRequest],
// [zero_trust.InputTeamsDevicesTaniumInputRequest] or
// [zero_trust.InputTeamsDevicesSentineloneS2sInputRequest].
type InputUnion interface {
	implementsZeroTrustInput()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InputUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesFileInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesUniqueClientIDInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesDomainJoinedInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesOSVersionInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesFirewallInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesSentineloneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesCarbonblackInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesDiskEncryptionInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesApplicationInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesClientCertificateInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesWorkspaceOneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesCrowdstrikeInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesIntuneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesKolideInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesTaniumInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputTeamsDevicesSentineloneS2sInputRequest{}),
		},
	)
}

type InputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                `json:"thumbprint"`
	JSON       inputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// inputTeamsDevicesFileInputRequestJSON contains the JSON metadata for the struct
// [InputTeamsDevicesFileInputRequest]
type inputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesFileInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesFileInputRequest) implementsZeroTrustInput() {}

type InputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem InputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            inputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// inputTeamsDevicesUniqueClientIDInputRequestJSON contains the JSON metadata for
// the struct [InputTeamsDevicesUniqueClientIDInputRequest]
type inputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesUniqueClientIDInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustInput() {}

// Operating System
type InputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	InputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  InputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	InputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      InputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	InputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos InputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

func (r InputTeamsDevicesUniqueClientIDInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case InputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid, InputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos, InputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos:
		return true
	}
	return false
}

type InputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem InputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                        `json:"domain"`
	JSON   inputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// inputTeamsDevicesDomainJoinedInputRequestJSON contains the JSON metadata for the
// struct [InputTeamsDevicesDomainJoinedInputRequest]
type inputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesDomainJoinedInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustInput() {}

// Operating System
type InputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	InputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows InputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

func (r InputTeamsDevicesDomainJoinedInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case InputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type InputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem InputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// operator
	Operator UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                     `json:"os_version_extra"`
	JSON           inputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// inputTeamsDevicesOSVersionInputRequestJSON contains the JSON metadata for the
// struct [InputTeamsDevicesOSVersionInputRequest]
type inputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesOSVersionInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesOSVersionInputRequest) implementsZeroTrustInput() {}

// Operating System
type InputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	InputTeamsDevicesOSVersionInputRequestOperatingSystemWindows InputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

func (r InputTeamsDevicesOSVersionInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case InputTeamsDevicesOSVersionInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type InputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem InputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            inputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// inputTeamsDevicesFirewallInputRequestJSON contains the JSON metadata for the
// struct [InputTeamsDevicesFirewallInputRequest]
type inputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesFirewallInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesFirewallInputRequest) implementsZeroTrustInput() {}

// Operating System
type InputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	InputTeamsDevicesFirewallInputRequestOperatingSystemWindows InputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	InputTeamsDevicesFirewallInputRequestOperatingSystemMac     InputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

func (r InputTeamsDevicesFirewallInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case InputTeamsDevicesFirewallInputRequestOperatingSystemWindows, InputTeamsDevicesFirewallInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type InputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                       `json:"thumbprint"`
	JSON       inputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// inputTeamsDevicesSentineloneInputRequestJSON contains the JSON metadata for the
// struct [InputTeamsDevicesSentineloneInputRequest]
type inputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesSentineloneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesSentineloneInputRequest) implementsZeroTrustInput() {}

type InputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                       `json:"thumbprint"`
	JSON       inputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// inputTeamsDevicesCarbonblackInputRequestJSON contains the JSON metadata for the
// struct [InputTeamsDevicesCarbonblackInputRequest]
type inputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesCarbonblackInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustInput() {}

type InputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                            `json:"requireAll"`
	JSON       inputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// inputTeamsDevicesDiskEncryptionInputRequestJSON contains the JSON metadata for
// the struct [InputTeamsDevicesDiskEncryptionInputRequest]
type inputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesDiskEncryptionInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustInput() {}

type InputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                       `json:"thumbprint"`
	JSON       inputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// inputTeamsDevicesApplicationInputRequestJSON contains the JSON metadata for the
// struct [InputTeamsDevicesApplicationInputRequest]
type inputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesApplicationInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesApplicationInputRequest) implementsZeroTrustInput() {}

type InputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                             `json:"cn,required"`
	JSON inputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// inputTeamsDevicesClientCertificateInputRequestJSON contains the JSON metadata
// for the struct [InputTeamsDevicesClientCertificateInputRequest]
type inputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesClientCertificateInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustInput() {}

type InputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus InputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                        `json:"connection_id,required"`
	JSON         inputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// inputTeamsDevicesWorkspaceOneInputRequestJSON contains the JSON metadata for the
// struct [InputTeamsDevicesWorkspaceOneInputRequest]
type inputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesWorkspaceOneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustInput() {}

// Compliance Status
type InputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	InputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    InputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	InputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant InputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	InputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      InputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

func (r InputTeamsDevicesWorkspaceOneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case InputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant, InputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant, InputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown:
		return true
	}
	return false
}

type InputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// operator
	Operator UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State InputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator InputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            inputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// inputTeamsDevicesCrowdstrikeInputRequestJSON contains the JSON metadata for the
// struct [InputTeamsDevicesCrowdstrikeInputRequest]
type inputTeamsDevicesCrowdstrikeInputRequestJSON struct {
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

func (r *InputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesCrowdstrikeInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustInput() {}

// For more details on state, please refer to the Crowdstrike documentation.
type InputTeamsDevicesCrowdstrikeInputRequestState string

const (
	InputTeamsDevicesCrowdstrikeInputRequestStateOnline  InputTeamsDevicesCrowdstrikeInputRequestState = "online"
	InputTeamsDevicesCrowdstrikeInputRequestStateOffline InputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	InputTeamsDevicesCrowdstrikeInputRequestStateUnknown InputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

func (r InputTeamsDevicesCrowdstrikeInputRequestState) IsKnown() bool {
	switch r {
	case InputTeamsDevicesCrowdstrikeInputRequestStateOnline, InputTeamsDevicesCrowdstrikeInputRequestStateOffline, InputTeamsDevicesCrowdstrikeInputRequestStateUnknown:
		return true
	}
	return false
}

// Version Operator
type InputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	InputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess            InputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	InputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals    InputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	InputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater         InputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	InputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals InputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	InputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals          InputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

func (r InputTeamsDevicesCrowdstrikeInputRequestVersionOperator) IsKnown() bool {
	switch r {
	case InputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess, InputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals, InputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater, InputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals, InputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals:
		return true
	}
	return false
}

type InputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus InputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                  `json:"connection_id,required"`
	JSON         inputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// inputTeamsDevicesIntuneInputRequestJSON contains the JSON metadata for the
// struct [InputTeamsDevicesIntuneInputRequest]
type inputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesIntuneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesIntuneInputRequest) implementsZeroTrustInput() {}

// Compliance Status
type InputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	InputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     InputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	InputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  InputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	InputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       InputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	InputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable InputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	InputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod InputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	InputTeamsDevicesIntuneInputRequestComplianceStatusError         InputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

func (r InputTeamsDevicesIntuneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case InputTeamsDevicesIntuneInputRequestComplianceStatusCompliant, InputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant, InputTeamsDevicesIntuneInputRequestComplianceStatusUnknown, InputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable, InputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod, InputTeamsDevicesIntuneInputRequestComplianceStatusError:
		return true
	}
	return false
}

type InputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator InputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                  `json:"issue_count,required"`
	JSON       inputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// inputTeamsDevicesKolideInputRequestJSON contains the JSON metadata for the
// struct [InputTeamsDevicesKolideInputRequest]
type inputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesKolideInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesKolideInputRequest) implementsZeroTrustInput() {}

// Count Operator
type InputTeamsDevicesKolideInputRequestCountOperator string

const (
	InputTeamsDevicesKolideInputRequestCountOperatorLess            InputTeamsDevicesKolideInputRequestCountOperator = "<"
	InputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals    InputTeamsDevicesKolideInputRequestCountOperator = "<="
	InputTeamsDevicesKolideInputRequestCountOperatorGreater         InputTeamsDevicesKolideInputRequestCountOperator = ">"
	InputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals InputTeamsDevicesKolideInputRequestCountOperator = ">="
	InputTeamsDevicesKolideInputRequestCountOperatorEquals          InputTeamsDevicesKolideInputRequestCountOperator = "=="
)

func (r InputTeamsDevicesKolideInputRequestCountOperator) IsKnown() bool {
	switch r {
	case InputTeamsDevicesKolideInputRequestCountOperatorLess, InputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals, InputTeamsDevicesKolideInputRequestCountOperatorGreater, InputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals, InputTeamsDevicesKolideInputRequestCountOperatorEquals:
		return true
	}
	return false
}

type InputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator InputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel InputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator InputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                 `json:"total_score"`
	JSON       inputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// inputTeamsDevicesTaniumInputRequestJSON contains the JSON metadata for the
// struct [InputTeamsDevicesTaniumInputRequest]
type inputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesTaniumInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesTaniumInputRequest) implementsZeroTrustInput() {}

// Operator to evaluate risk_level or eid_last_seen.
type InputTeamsDevicesTaniumInputRequestOperator string

const (
	InputTeamsDevicesTaniumInputRequestOperatorLess            InputTeamsDevicesTaniumInputRequestOperator = "<"
	InputTeamsDevicesTaniumInputRequestOperatorLessOrEquals    InputTeamsDevicesTaniumInputRequestOperator = "<="
	InputTeamsDevicesTaniumInputRequestOperatorGreater         InputTeamsDevicesTaniumInputRequestOperator = ">"
	InputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals InputTeamsDevicesTaniumInputRequestOperator = ">="
	InputTeamsDevicesTaniumInputRequestOperatorEquals          InputTeamsDevicesTaniumInputRequestOperator = "=="
)

func (r InputTeamsDevicesTaniumInputRequestOperator) IsKnown() bool {
	switch r {
	case InputTeamsDevicesTaniumInputRequestOperatorLess, InputTeamsDevicesTaniumInputRequestOperatorLessOrEquals, InputTeamsDevicesTaniumInputRequestOperatorGreater, InputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals, InputTeamsDevicesTaniumInputRequestOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type InputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	InputTeamsDevicesTaniumInputRequestRiskLevelLow      InputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	InputTeamsDevicesTaniumInputRequestRiskLevelMedium   InputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	InputTeamsDevicesTaniumInputRequestRiskLevelHigh     InputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	InputTeamsDevicesTaniumInputRequestRiskLevelCritical InputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

func (r InputTeamsDevicesTaniumInputRequestRiskLevel) IsKnown() bool {
	switch r {
	case InputTeamsDevicesTaniumInputRequestRiskLevelLow, InputTeamsDevicesTaniumInputRequestRiskLevelMedium, InputTeamsDevicesTaniumInputRequestRiskLevelHigh, InputTeamsDevicesTaniumInputRequestRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type InputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	InputTeamsDevicesTaniumInputRequestScoreOperatorLess            InputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	InputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals    InputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	InputTeamsDevicesTaniumInputRequestScoreOperatorGreater         InputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	InputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals InputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	InputTeamsDevicesTaniumInputRequestScoreOperatorEquals          InputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

func (r InputTeamsDevicesTaniumInputRequestScoreOperator) IsKnown() bool {
	switch r {
	case InputTeamsDevicesTaniumInputRequestScoreOperatorLess, InputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals, InputTeamsDevicesTaniumInputRequestScoreOperatorGreater, InputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals, InputTeamsDevicesTaniumInputRequestScoreOperatorEquals:
		return true
	}
	return false
}

type InputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus InputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 `json:"operator"`
	JSON     inputTeamsDevicesSentineloneS2sInputRequestJSON  `json:"-"`
}

// inputTeamsDevicesSentineloneS2sInputRequestJSON contains the JSON metadata for
// the struct [InputTeamsDevicesSentineloneS2sInputRequest]
type inputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTeamsDevicesSentineloneS2sInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustInput() {}

// Network status of device.
type InputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	InputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     InputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	InputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  InputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	InputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting InputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	InputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    InputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

func (r InputTeamsDevicesSentineloneS2sInputRequestNetworkStatus) IsKnown() bool {
	switch r {
	case InputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected, InputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected, InputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting, InputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting:
		return true
	}
	return false
}

// Compliance Status
type InputComplianceStatus string

const (
	InputComplianceStatusCompliant     InputComplianceStatus = "compliant"
	InputComplianceStatusNoncompliant  InputComplianceStatus = "noncompliant"
	InputComplianceStatusUnknown       InputComplianceStatus = "unknown"
	InputComplianceStatusNotapplicable InputComplianceStatus = "notapplicable"
	InputComplianceStatusIngraceperiod InputComplianceStatus = "ingraceperiod"
	InputComplianceStatusError         InputComplianceStatus = "error"
)

func (r InputComplianceStatus) IsKnown() bool {
	switch r {
	case InputComplianceStatusCompliant, InputComplianceStatusNoncompliant, InputComplianceStatusUnknown, InputComplianceStatusNotapplicable, InputComplianceStatusIngraceperiod, InputComplianceStatusError:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type InputState string

const (
	InputStateOnline  InputState = "online"
	InputStateOffline InputState = "offline"
	InputStateUnknown InputState = "unknown"
)

func (r InputState) IsKnown() bool {
	switch r {
	case InputStateOnline, InputStateOffline, InputStateUnknown:
		return true
	}
	return false
}

// Version Operator
type InputVersionOperator string

const (
	InputVersionOperatorLess            InputVersionOperator = "<"
	InputVersionOperatorLessOrEquals    InputVersionOperator = "<="
	InputVersionOperatorGreater         InputVersionOperator = ">"
	InputVersionOperatorGreaterOrEquals InputVersionOperator = ">="
	InputVersionOperatorEquals          InputVersionOperator = "=="
)

func (r InputVersionOperator) IsKnown() bool {
	switch r {
	case InputVersionOperatorLess, InputVersionOperatorLessOrEquals, InputVersionOperatorGreater, InputVersionOperatorGreaterOrEquals, InputVersionOperatorEquals:
		return true
	}
	return false
}

// Count Operator
type InputCountOperator string

const (
	InputCountOperatorLess            InputCountOperator = "<"
	InputCountOperatorLessOrEquals    InputCountOperator = "<="
	InputCountOperatorGreater         InputCountOperator = ">"
	InputCountOperatorGreaterOrEquals InputCountOperator = ">="
	InputCountOperatorEquals          InputCountOperator = "=="
)

func (r InputCountOperator) IsKnown() bool {
	switch r {
	case InputCountOperatorLess, InputCountOperatorLessOrEquals, InputCountOperatorGreater, InputCountOperatorGreaterOrEquals, InputCountOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type InputRiskLevel string

const (
	InputRiskLevelLow      InputRiskLevel = "low"
	InputRiskLevelMedium   InputRiskLevel = "medium"
	InputRiskLevelHigh     InputRiskLevel = "high"
	InputRiskLevelCritical InputRiskLevel = "critical"
)

func (r InputRiskLevel) IsKnown() bool {
	switch r {
	case InputRiskLevelLow, InputRiskLevelMedium, InputRiskLevelHigh, InputRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type InputScoreOperator string

const (
	InputScoreOperatorLess            InputScoreOperator = "<"
	InputScoreOperatorLessOrEquals    InputScoreOperator = "<="
	InputScoreOperatorGreater         InputScoreOperator = ">"
	InputScoreOperatorGreaterOrEquals InputScoreOperator = ">="
	InputScoreOperatorEquals          InputScoreOperator = "=="
)

func (r InputScoreOperator) IsKnown() bool {
	switch r {
	case InputScoreOperatorLess, InputScoreOperatorLessOrEquals, InputScoreOperatorGreater, InputScoreOperatorGreaterOrEquals, InputScoreOperatorEquals:
		return true
	}
	return false
}

// Network status of device.
type InputNetworkStatus string

const (
	InputNetworkStatusConnected     InputNetworkStatus = "connected"
	InputNetworkStatusDisconnected  InputNetworkStatus = "disconnected"
	InputNetworkStatusDisconnecting InputNetworkStatus = "disconnecting"
	InputNetworkStatusConnecting    InputNetworkStatus = "connecting"
)

func (r InputNetworkStatus) IsKnown() bool {
	switch r {
	case InputNetworkStatusConnected, InputNetworkStatusDisconnected, InputNetworkStatusDisconnecting, InputNetworkStatusConnecting:
		return true
	}
	return false
}

// The value to be checked against.
type InputParam struct {
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system"`
	// File path.
	Path param.Field[string] `json:"path"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
	// List ID.
	ID param.Field[string] `json:"id"`
	// Domain
	Domain param.Field[string] `json:"domain"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator"`
	// Operating System Distribution Name (linux only)
	OSDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra param.Field[string] `json:"os_version_extra"`
	// Version of OS
	Version param.Field[string] `json:"version"`
	// Enabled
	Enabled    param.Field[bool]        `json:"enabled"`
	CheckDisks param.Field[interface{}] `json:"checkDisks,required"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn"`
	// Compliance Status
	ComplianceStatus param.Field[InputComplianceStatus] `json:"compliance_status"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// Os Version
	OS param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[InputState] `json:"state"`
	// Version Operator
	VersionOperator param.Field[InputVersionOperator] `json:"versionOperator"`
	// Count Operator
	CountOperator param.Field[InputCountOperator] `json:"countOperator"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[InputRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[InputScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[InputNetworkStatus] `json:"network_status"`
}

func (r InputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputParam) implementsZeroTrustInputUnionParam() {}

// The value to be checked against.
//
// Satisfied by [zero_trust.InputTeamsDevicesFileInputRequestParam],
// [zero_trust.InputTeamsDevicesUniqueClientIDInputRequestParam],
// [zero_trust.InputTeamsDevicesDomainJoinedInputRequestParam],
// [zero_trust.InputTeamsDevicesOSVersionInputRequestParam],
// [zero_trust.InputTeamsDevicesFirewallInputRequestParam],
// [zero_trust.InputTeamsDevicesSentineloneInputRequestParam],
// [zero_trust.InputTeamsDevicesCarbonblackInputRequestParam],
// [zero_trust.InputTeamsDevicesDiskEncryptionInputRequestParam],
// [zero_trust.InputTeamsDevicesApplicationInputRequestParam],
// [zero_trust.InputTeamsDevicesClientCertificateInputRequestParam],
// [zero_trust.InputTeamsDevicesWorkspaceOneInputRequestParam],
// [zero_trust.InputTeamsDevicesCrowdstrikeInputRequestParam],
// [zero_trust.InputTeamsDevicesIntuneInputRequestParam],
// [zero_trust.InputTeamsDevicesKolideInputRequestParam],
// [zero_trust.InputTeamsDevicesTaniumInputRequestParam],
// [zero_trust.InputTeamsDevicesSentineloneS2sInputRequestParam], [InputParam].
type InputUnionParam interface {
	implementsZeroTrustInputUnionParam()
}

type InputTeamsDevicesFileInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r InputTeamsDevicesFileInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesFileInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesUniqueClientIDInputRequestParam struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[InputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r InputTeamsDevicesUniqueClientIDInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesUniqueClientIDInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesDomainJoinedInputRequestParam struct {
	// Operating System
	OperatingSystem param.Field[InputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r InputTeamsDevicesDomainJoinedInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesDomainJoinedInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesOSVersionInputRequestParam struct {
	// Operating System
	OperatingSystem param.Field[InputTeamsDevicesOSVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator,required"`
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

func (r InputTeamsDevicesOSVersionInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesOSVersionInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesFirewallInputRequestParam struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[InputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r InputTeamsDevicesFirewallInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesFirewallInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesSentineloneInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r InputTeamsDevicesSentineloneInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesSentineloneInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesCarbonblackInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r InputTeamsDevicesCarbonblackInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesCarbonblackInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesDiskEncryptionInputRequestParam struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r InputTeamsDevicesDiskEncryptionInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesDiskEncryptionInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesApplicationInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r InputTeamsDevicesApplicationInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesApplicationInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesClientCertificateInputRequestParam struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r InputTeamsDevicesClientCertificateInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesClientCertificateInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesWorkspaceOneInputRequestParam struct {
	// Compliance Status
	ComplianceStatus param.Field[InputTeamsDevicesWorkspaceOneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r InputTeamsDevicesWorkspaceOneInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesWorkspaceOneInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesCrowdstrikeInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator"`
	// Os Version
	OS param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[InputTeamsDevicesCrowdstrikeInputRequestState] `json:"state"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[InputTeamsDevicesCrowdstrikeInputRequestVersionOperator] `json:"versionOperator"`
}

func (r InputTeamsDevicesCrowdstrikeInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesCrowdstrikeInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesIntuneInputRequestParam struct {
	// Compliance Status
	ComplianceStatus param.Field[InputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r InputTeamsDevicesIntuneInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesIntuneInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesKolideInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[InputTeamsDevicesKolideInputRequestCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r InputTeamsDevicesKolideInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesKolideInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesTaniumInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[InputTeamsDevicesTaniumInputRequestOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[InputTeamsDevicesTaniumInputRequestRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[InputTeamsDevicesTaniumInputRequestScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r InputTeamsDevicesTaniumInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesTaniumInputRequestParam) implementsZeroTrustInputUnionParam() {}

type InputTeamsDevicesSentineloneS2sInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[InputTeamsDevicesSentineloneS2sInputRequestNetworkStatus] `json:"network_status"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator"`
}

func (r InputTeamsDevicesSentineloneS2sInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputTeamsDevicesSentineloneS2sInputRequestParam) implementsZeroTrustInputUnionParam() {}

type MatchItem struct {
	Platform MatchItemPlatform `json:"platform"`
	JSON     matchItemJSON     `json:"-"`
}

// matchItemJSON contains the JSON metadata for the struct [MatchItem]
type matchItemJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MatchItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matchItemJSON) RawJSON() string {
	return r.raw
}

type MatchItemPlatform string

const (
	MatchItemPlatformWindows MatchItemPlatform = "windows"
	MatchItemPlatformMac     MatchItemPlatform = "mac"
	MatchItemPlatformLinux   MatchItemPlatform = "linux"
	MatchItemPlatformAndroid MatchItemPlatform = "android"
	MatchItemPlatformIos     MatchItemPlatform = "ios"
)

func (r MatchItemPlatform) IsKnown() bool {
	switch r {
	case MatchItemPlatformWindows, MatchItemPlatformMac, MatchItemPlatformLinux, MatchItemPlatformAndroid, MatchItemPlatformIos:
		return true
	}
	return false
}

type MatchItemParam struct {
	Platform param.Field[MatchItemPlatform] `json:"platform"`
}

func (r MatchItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// operator
type UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 string

const (
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Less            UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = "<"
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930LessOrEquals    UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = "<="
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Greater         UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = ">"
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930GreaterOrEquals UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = ">="
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Equals          UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = "=="
)

func (r UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Less, UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930LessOrEquals, UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Greater, UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930GreaterOrEquals, UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Equals:
		return true
	}
	return false
}

// Operating system
type UnnamedSchemaRef41885dd46b9e0294254c49305a273681 string

const (
	UnnamedSchemaRef41885dd46b9e0294254c49305a273681Windows UnnamedSchemaRef41885dd46b9e0294254c49305a273681 = "windows"
	UnnamedSchemaRef41885dd46b9e0294254c49305a273681Linux   UnnamedSchemaRef41885dd46b9e0294254c49305a273681 = "linux"
	UnnamedSchemaRef41885dd46b9e0294254c49305a273681Mac     UnnamedSchemaRef41885dd46b9e0294254c49305a273681 = "mac"
)

func (r UnnamedSchemaRef41885dd46b9e0294254c49305a273681) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef41885dd46b9e0294254c49305a273681Windows, UnnamedSchemaRef41885dd46b9e0294254c49305a273681Linux, UnnamedSchemaRef41885dd46b9e0294254c49305a273681Mac:
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
	Input param.Field[InputUnionParam] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]MatchItemParam] `json:"match"`
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

type DevicePostureNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DevicePostureRule                                         `json:"result,required,nullable"`
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
	Input param.Field[InputUnionParam] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]MatchItemParam] `json:"match"`
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

type DevicePostureUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DevicePostureRule                                         `json:"result,required,nullable"`
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DevicePostureDeleteResponse                               `json:"result,required,nullable"`
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DevicePostureRule                                         `json:"result,required,nullable"`
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
