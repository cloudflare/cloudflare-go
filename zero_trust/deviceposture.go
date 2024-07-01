// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// DevicePostureService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePostureService] method instead.
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
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
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
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
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
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *DevicePostureService) Delete(ctx context.Context, ruleID string, body DevicePostureDeleteParams, opts ...option.RequestOption) (res *DevicePostureDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureDeleteResponseEnvelope
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", body.AccountID, ruleID)
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
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CarbonblackInput = string

type CarbonblackInputParam = string

type ClientCertificateInput struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                     `json:"cn,required"`
	JSON clientCertificateInputJSON `json:"-"`
}

// clientCertificateInputJSON contains the JSON metadata for the struct
// [ClientCertificateInput]
type clientCertificateInputJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ClientCertificateInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateInputJSON) RawJSON() string {
	return r.raw
}

func (r ClientCertificateInput) implementsZeroTrustDeviceInput() {}

type ClientCertificateInputParam struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r ClientCertificateInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ClientCertificateInputParam) implementsZeroTrustDeviceInputUnionParam() {}

type CrowdstrikeInput struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// operator
	Operator CrowdstrikeInputOperator `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State CrowdstrikeInputState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator CrowdstrikeInputVersionOperator `json:"versionOperator"`
	JSON            crowdstrikeInputJSON            `json:"-"`
}

// crowdstrikeInputJSON contains the JSON metadata for the struct
// [CrowdstrikeInput]
type crowdstrikeInputJSON struct {
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

func (r *CrowdstrikeInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crowdstrikeInputJSON) RawJSON() string {
	return r.raw
}

func (r CrowdstrikeInput) implementsZeroTrustDeviceInput() {}

// operator
type CrowdstrikeInputOperator string

const (
	CrowdstrikeInputOperatorLess            CrowdstrikeInputOperator = "<"
	CrowdstrikeInputOperatorLessOrEquals    CrowdstrikeInputOperator = "<="
	CrowdstrikeInputOperatorGreater         CrowdstrikeInputOperator = ">"
	CrowdstrikeInputOperatorGreaterOrEquals CrowdstrikeInputOperator = ">="
	CrowdstrikeInputOperatorEquals          CrowdstrikeInputOperator = "=="
)

func (r CrowdstrikeInputOperator) IsKnown() bool {
	switch r {
	case CrowdstrikeInputOperatorLess, CrowdstrikeInputOperatorLessOrEquals, CrowdstrikeInputOperatorGreater, CrowdstrikeInputOperatorGreaterOrEquals, CrowdstrikeInputOperatorEquals:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type CrowdstrikeInputState string

const (
	CrowdstrikeInputStateOnline  CrowdstrikeInputState = "online"
	CrowdstrikeInputStateOffline CrowdstrikeInputState = "offline"
	CrowdstrikeInputStateUnknown CrowdstrikeInputState = "unknown"
)

func (r CrowdstrikeInputState) IsKnown() bool {
	switch r {
	case CrowdstrikeInputStateOnline, CrowdstrikeInputStateOffline, CrowdstrikeInputStateUnknown:
		return true
	}
	return false
}

// Version Operator
type CrowdstrikeInputVersionOperator string

const (
	CrowdstrikeInputVersionOperatorLess            CrowdstrikeInputVersionOperator = "<"
	CrowdstrikeInputVersionOperatorLessOrEquals    CrowdstrikeInputVersionOperator = "<="
	CrowdstrikeInputVersionOperatorGreater         CrowdstrikeInputVersionOperator = ">"
	CrowdstrikeInputVersionOperatorGreaterOrEquals CrowdstrikeInputVersionOperator = ">="
	CrowdstrikeInputVersionOperatorEquals          CrowdstrikeInputVersionOperator = "=="
)

func (r CrowdstrikeInputVersionOperator) IsKnown() bool {
	switch r {
	case CrowdstrikeInputVersionOperatorLess, CrowdstrikeInputVersionOperatorLessOrEquals, CrowdstrikeInputVersionOperatorGreater, CrowdstrikeInputVersionOperatorGreaterOrEquals, CrowdstrikeInputVersionOperatorEquals:
		return true
	}
	return false
}

type CrowdstrikeInputParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// operator
	Operator param.Field[CrowdstrikeInputOperator] `json:"operator"`
	// Os Version
	OS param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[CrowdstrikeInputState] `json:"state"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[CrowdstrikeInputVersionOperator] `json:"versionOperator"`
}

func (r CrowdstrikeInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CrowdstrikeInputParam) implementsZeroTrustDeviceInputUnionParam() {}

// The value to be checked against.
type DeviceInput struct {
	// Whether or not file exists
	Exists bool `json:"exists"`
	// Operating system
	OperatingSystem DeviceInputOperatingSystem `json:"operating_system"`
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
	Operator DeviceInputOperator `json:"operator"`
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
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [[]CarbonblackInput].
	CheckDisks interface{} `json:"checkDisks,required"`
	// Whether to check all disks for encryption.
	RequireAll bool `json:"requireAll"`
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id"`
	// Common Name that is protected by the certificate
	Cn string `json:"cn"`
	// Compliance Status
	ComplianceStatus DeviceInputComplianceStatus `json:"compliance_status"`
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
	State DeviceInputState `json:"state"`
	// Version Operator
	VersionOperator DeviceInputVersionOperator `json:"versionOperator"`
	// Count Operator
	CountOperator DeviceInputCountOperator `json:"countOperator"`
	// The Number of Issues.
	IssueCount string `json:"issue_count"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DeviceInputRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator DeviceInputScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64 `json:"total_score"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus DeviceInputNetworkStatus `json:"network_status"`
	JSON          deviceInputJSON          `json:"-"`
	union         DeviceInputUnion
}

// deviceInputJSON contains the JSON metadata for the struct [DeviceInput]
type deviceInputJSON struct {
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

func (r deviceInputJSON) RawJSON() string {
	return r.raw
}

func (r *DeviceInput) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DeviceInputUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [zero_trust.FileInput],
// [zero_trust.UniqueClientIDInput], [zero_trust.DomainJoinedInput],
// [zero_trust.OSVersionInput], [zero_trust.FirewallInput],
// [zero_trust.SentineloneInput],
// [zero_trust.DeviceInputTeamsDevicesCarbonblackInputRequest],
// [zero_trust.DiskEncryptionInput],
// [zero_trust.DeviceInputTeamsDevicesApplicationInputRequest],
// [zero_trust.ClientCertificateInput], [zero_trust.WorkspaceOneInput],
// [zero_trust.CrowdstrikeInput], [zero_trust.IntuneInput],
// [zero_trust.KolideInput], [zero_trust.TaniumInput],
// [zero_trust.SentineloneS2sInput].
func (r DeviceInput) AsUnion() DeviceInputUnion {
	return r.union
}

// The value to be checked against.
//
// Union satisfied by [zero_trust.FileInput], [zero_trust.UniqueClientIDInput],
// [zero_trust.DomainJoinedInput], [zero_trust.OSVersionInput],
// [zero_trust.FirewallInput], [zero_trust.SentineloneInput],
// [zero_trust.DeviceInputTeamsDevicesCarbonblackInputRequest],
// [zero_trust.DiskEncryptionInput],
// [zero_trust.DeviceInputTeamsDevicesApplicationInputRequest],
// [zero_trust.ClientCertificateInput], [zero_trust.WorkspaceOneInput],
// [zero_trust.CrowdstrikeInput], [zero_trust.IntuneInput],
// [zero_trust.KolideInput], [zero_trust.TaniumInput] or
// [zero_trust.SentineloneS2sInput].
type DeviceInputUnion interface {
	implementsZeroTrustDeviceInput()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeviceInputUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FileInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UniqueClientIDInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DomainJoinedInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OSVersionInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SentineloneInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesCarbonblackInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DiskEncryptionInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesApplicationInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ClientCertificateInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkspaceOneInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CrowdstrikeInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntuneInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(KolideInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TaniumInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SentineloneS2sInput{}),
		},
	)
}

type DeviceInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                             `json:"thumbprint"`
	JSON       deviceInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesCarbonblackInputRequestJSON contains the JSON metadata
// for the struct [DeviceInputTeamsDevicesCarbonblackInputRequest]
type deviceInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesCarbonblackInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDeviceInput() {}

// Operating system
type DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

func (r DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows, DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux, DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                             `json:"thumbprint"`
	JSON       deviceInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesApplicationInputRequestJSON contains the JSON metadata
// for the struct [DeviceInputTeamsDevicesApplicationInputRequest]
type deviceInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesApplicationInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDeviceInput() {}

// Operating system
type DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DeviceInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DeviceInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DeviceInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

func (r DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesApplicationInputRequestOperatingSystemWindows, DeviceInputTeamsDevicesApplicationInputRequestOperatingSystemLinux, DeviceInputTeamsDevicesApplicationInputRequestOperatingSystemMac:
		return true
	}
	return false
}

// Operating system
type DeviceInputOperatingSystem string

const (
	DeviceInputOperatingSystemWindows  DeviceInputOperatingSystem = "windows"
	DeviceInputOperatingSystemLinux    DeviceInputOperatingSystem = "linux"
	DeviceInputOperatingSystemMac      DeviceInputOperatingSystem = "mac"
	DeviceInputOperatingSystemAndroid  DeviceInputOperatingSystem = "android"
	DeviceInputOperatingSystemIos      DeviceInputOperatingSystem = "ios"
	DeviceInputOperatingSystemChromeos DeviceInputOperatingSystem = "chromeos"
)

func (r DeviceInputOperatingSystem) IsKnown() bool {
	switch r {
	case DeviceInputOperatingSystemWindows, DeviceInputOperatingSystemLinux, DeviceInputOperatingSystemMac, DeviceInputOperatingSystemAndroid, DeviceInputOperatingSystemIos, DeviceInputOperatingSystemChromeos:
		return true
	}
	return false
}

// operator
type DeviceInputOperator string

const (
	DeviceInputOperatorLess            DeviceInputOperator = "<"
	DeviceInputOperatorLessOrEquals    DeviceInputOperator = "<="
	DeviceInputOperatorGreater         DeviceInputOperator = ">"
	DeviceInputOperatorGreaterOrEquals DeviceInputOperator = ">="
	DeviceInputOperatorEquals          DeviceInputOperator = "=="
)

func (r DeviceInputOperator) IsKnown() bool {
	switch r {
	case DeviceInputOperatorLess, DeviceInputOperatorLessOrEquals, DeviceInputOperatorGreater, DeviceInputOperatorGreaterOrEquals, DeviceInputOperatorEquals:
		return true
	}
	return false
}

// Compliance Status
type DeviceInputComplianceStatus string

const (
	DeviceInputComplianceStatusCompliant     DeviceInputComplianceStatus = "compliant"
	DeviceInputComplianceStatusNoncompliant  DeviceInputComplianceStatus = "noncompliant"
	DeviceInputComplianceStatusUnknown       DeviceInputComplianceStatus = "unknown"
	DeviceInputComplianceStatusNotapplicable DeviceInputComplianceStatus = "notapplicable"
	DeviceInputComplianceStatusIngraceperiod DeviceInputComplianceStatus = "ingraceperiod"
	DeviceInputComplianceStatusError         DeviceInputComplianceStatus = "error"
)

func (r DeviceInputComplianceStatus) IsKnown() bool {
	switch r {
	case DeviceInputComplianceStatusCompliant, DeviceInputComplianceStatusNoncompliant, DeviceInputComplianceStatusUnknown, DeviceInputComplianceStatusNotapplicable, DeviceInputComplianceStatusIngraceperiod, DeviceInputComplianceStatusError:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type DeviceInputState string

const (
	DeviceInputStateOnline  DeviceInputState = "online"
	DeviceInputStateOffline DeviceInputState = "offline"
	DeviceInputStateUnknown DeviceInputState = "unknown"
)

func (r DeviceInputState) IsKnown() bool {
	switch r {
	case DeviceInputStateOnline, DeviceInputStateOffline, DeviceInputStateUnknown:
		return true
	}
	return false
}

// Version Operator
type DeviceInputVersionOperator string

const (
	DeviceInputVersionOperatorLess            DeviceInputVersionOperator = "<"
	DeviceInputVersionOperatorLessOrEquals    DeviceInputVersionOperator = "<="
	DeviceInputVersionOperatorGreater         DeviceInputVersionOperator = ">"
	DeviceInputVersionOperatorGreaterOrEquals DeviceInputVersionOperator = ">="
	DeviceInputVersionOperatorEquals          DeviceInputVersionOperator = "=="
)

func (r DeviceInputVersionOperator) IsKnown() bool {
	switch r {
	case DeviceInputVersionOperatorLess, DeviceInputVersionOperatorLessOrEquals, DeviceInputVersionOperatorGreater, DeviceInputVersionOperatorGreaterOrEquals, DeviceInputVersionOperatorEquals:
		return true
	}
	return false
}

// Count Operator
type DeviceInputCountOperator string

const (
	DeviceInputCountOperatorLess            DeviceInputCountOperator = "<"
	DeviceInputCountOperatorLessOrEquals    DeviceInputCountOperator = "<="
	DeviceInputCountOperatorGreater         DeviceInputCountOperator = ">"
	DeviceInputCountOperatorGreaterOrEquals DeviceInputCountOperator = ">="
	DeviceInputCountOperatorEquals          DeviceInputCountOperator = "=="
)

func (r DeviceInputCountOperator) IsKnown() bool {
	switch r {
	case DeviceInputCountOperatorLess, DeviceInputCountOperatorLessOrEquals, DeviceInputCountOperatorGreater, DeviceInputCountOperatorGreaterOrEquals, DeviceInputCountOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type DeviceInputRiskLevel string

const (
	DeviceInputRiskLevelLow      DeviceInputRiskLevel = "low"
	DeviceInputRiskLevelMedium   DeviceInputRiskLevel = "medium"
	DeviceInputRiskLevelHigh     DeviceInputRiskLevel = "high"
	DeviceInputRiskLevelCritical DeviceInputRiskLevel = "critical"
)

func (r DeviceInputRiskLevel) IsKnown() bool {
	switch r {
	case DeviceInputRiskLevelLow, DeviceInputRiskLevelMedium, DeviceInputRiskLevelHigh, DeviceInputRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type DeviceInputScoreOperator string

const (
	DeviceInputScoreOperatorLess            DeviceInputScoreOperator = "<"
	DeviceInputScoreOperatorLessOrEquals    DeviceInputScoreOperator = "<="
	DeviceInputScoreOperatorGreater         DeviceInputScoreOperator = ">"
	DeviceInputScoreOperatorGreaterOrEquals DeviceInputScoreOperator = ">="
	DeviceInputScoreOperatorEquals          DeviceInputScoreOperator = "=="
)

func (r DeviceInputScoreOperator) IsKnown() bool {
	switch r {
	case DeviceInputScoreOperatorLess, DeviceInputScoreOperatorLessOrEquals, DeviceInputScoreOperatorGreater, DeviceInputScoreOperatorGreaterOrEquals, DeviceInputScoreOperatorEquals:
		return true
	}
	return false
}

// Network status of device.
type DeviceInputNetworkStatus string

const (
	DeviceInputNetworkStatusConnected     DeviceInputNetworkStatus = "connected"
	DeviceInputNetworkStatusDisconnected  DeviceInputNetworkStatus = "disconnected"
	DeviceInputNetworkStatusDisconnecting DeviceInputNetworkStatus = "disconnecting"
	DeviceInputNetworkStatusConnecting    DeviceInputNetworkStatus = "connecting"
)

func (r DeviceInputNetworkStatus) IsKnown() bool {
	switch r {
	case DeviceInputNetworkStatusConnected, DeviceInputNetworkStatusDisconnected, DeviceInputNetworkStatusDisconnecting, DeviceInputNetworkStatusConnecting:
		return true
	}
	return false
}

// The value to be checked against.
type DeviceInputParam struct {
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// Operating system
	OperatingSystem param.Field[DeviceInputOperatingSystem] `json:"operating_system"`
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
	Operator param.Field[DeviceInputOperator] `json:"operator"`
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
	ComplianceStatus param.Field[DeviceInputComplianceStatus] `json:"compliance_status"`
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
	State param.Field[DeviceInputState] `json:"state"`
	// Version Operator
	VersionOperator param.Field[DeviceInputVersionOperator] `json:"versionOperator"`
	// Count Operator
	CountOperator param.Field[DeviceInputCountOperator] `json:"countOperator"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[DeviceInputRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[DeviceInputScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[DeviceInputNetworkStatus] `json:"network_status"`
}

func (r DeviceInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputParam) implementsZeroTrustDeviceInputUnionParam() {}

// The value to be checked against.
//
// Satisfied by [zero_trust.FileInputParam], [zero_trust.UniqueClientIDInputParam],
// [zero_trust.DomainJoinedInputParam], [zero_trust.OSVersionInputParam],
// [zero_trust.FirewallInputParam], [zero_trust.SentineloneInputParam],
// [zero_trust.DeviceInputTeamsDevicesCarbonblackInputRequestParam],
// [zero_trust.DiskEncryptionInputParam],
// [zero_trust.DeviceInputTeamsDevicesApplicationInputRequestParam],
// [zero_trust.ClientCertificateInputParam], [zero_trust.WorkspaceOneInputParam],
// [zero_trust.CrowdstrikeInputParam], [zero_trust.IntuneInputParam],
// [zero_trust.KolideInputParam], [zero_trust.TaniumInputParam],
// [zero_trust.SentineloneS2sInputParam], [DeviceInputParam].
type DeviceInputUnionParam interface {
	implementsZeroTrustDeviceInputUnionParam()
}

type DeviceInputTeamsDevicesCarbonblackInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DeviceInputTeamsDevicesCarbonblackInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesCarbonblackInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesApplicationInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DeviceInputTeamsDevicesApplicationInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesApplicationInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceMatch struct {
	Platform DeviceMatchPlatform `json:"platform"`
	JSON     deviceMatchJSON     `json:"-"`
}

// deviceMatchJSON contains the JSON metadata for the struct [DeviceMatch]
type deviceMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceMatchJSON) RawJSON() string {
	return r.raw
}

type DeviceMatchPlatform string

const (
	DeviceMatchPlatformWindows DeviceMatchPlatform = "windows"
	DeviceMatchPlatformMac     DeviceMatchPlatform = "mac"
	DeviceMatchPlatformLinux   DeviceMatchPlatform = "linux"
	DeviceMatchPlatformAndroid DeviceMatchPlatform = "android"
	DeviceMatchPlatformIos     DeviceMatchPlatform = "ios"
)

func (r DeviceMatchPlatform) IsKnown() bool {
	switch r {
	case DeviceMatchPlatformWindows, DeviceMatchPlatformMac, DeviceMatchPlatformLinux, DeviceMatchPlatformAndroid, DeviceMatchPlatformIos:
		return true
	}
	return false
}

type DeviceMatchParam struct {
	Platform param.Field[DeviceMatchPlatform] `json:"platform"`
}

func (r DeviceMatchParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Input DeviceInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []DeviceMatch `json:"match"`
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

type DiskEncryptionInput struct {
	// List of volume names to be checked for encryption.
	CheckDisks []CarbonblackInput `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                    `json:"requireAll"`
	JSON       diskEncryptionInputJSON `json:"-"`
}

// diskEncryptionInputJSON contains the JSON metadata for the struct
// [DiskEncryptionInput]
type diskEncryptionInputJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiskEncryptionInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r diskEncryptionInputJSON) RawJSON() string {
	return r.raw
}

func (r DiskEncryptionInput) implementsZeroTrustDeviceInput() {}

type DiskEncryptionInputParam struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]CarbonblackInputParam] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r DiskEncryptionInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiskEncryptionInputParam) implementsZeroTrustDeviceInputUnionParam() {}

type DomainJoinedInput struct {
	// Operating System
	OperatingSystem DomainJoinedInputOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                `json:"domain"`
	JSON   domainJoinedInputJSON `json:"-"`
}

// domainJoinedInputJSON contains the JSON metadata for the struct
// [DomainJoinedInput]
type domainJoinedInputJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DomainJoinedInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainJoinedInputJSON) RawJSON() string {
	return r.raw
}

func (r DomainJoinedInput) implementsZeroTrustDeviceInput() {}

// Operating System
type DomainJoinedInputOperatingSystem string

const (
	DomainJoinedInputOperatingSystemWindows DomainJoinedInputOperatingSystem = "windows"
)

func (r DomainJoinedInputOperatingSystem) IsKnown() bool {
	switch r {
	case DomainJoinedInputOperatingSystemWindows:
		return true
	}
	return false
}

type DomainJoinedInputParam struct {
	// Operating System
	OperatingSystem param.Field[DomainJoinedInputOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r DomainJoinedInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DomainJoinedInputParam) implementsZeroTrustDeviceInputUnionParam() {}

type FileInput struct {
	// Operating system
	OperatingSystem FileInputOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string        `json:"thumbprint"`
	JSON       fileInputJSON `json:"-"`
}

// fileInputJSON contains the JSON metadata for the struct [FileInput]
type fileInputJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FileInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileInputJSON) RawJSON() string {
	return r.raw
}

func (r FileInput) implementsZeroTrustDeviceInput() {}

// Operating system
type FileInputOperatingSystem string

const (
	FileInputOperatingSystemWindows FileInputOperatingSystem = "windows"
	FileInputOperatingSystemLinux   FileInputOperatingSystem = "linux"
	FileInputOperatingSystemMac     FileInputOperatingSystem = "mac"
)

func (r FileInputOperatingSystem) IsKnown() bool {
	switch r {
	case FileInputOperatingSystemWindows, FileInputOperatingSystemLinux, FileInputOperatingSystemMac:
		return true
	}
	return false
}

type FileInputParam struct {
	// Operating system
	OperatingSystem param.Field[FileInputOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r FileInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FileInputParam) implementsZeroTrustDeviceInputUnionParam() {}

type FirewallInput struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem FirewallInputOperatingSystem `json:"operating_system,required"`
	JSON            firewallInputJSON            `json:"-"`
}

// firewallInputJSON contains the JSON metadata for the struct [FirewallInput]
type firewallInputJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FirewallInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallInputJSON) RawJSON() string {
	return r.raw
}

func (r FirewallInput) implementsZeroTrustDeviceInput() {}

// Operating System
type FirewallInputOperatingSystem string

const (
	FirewallInputOperatingSystemWindows FirewallInputOperatingSystem = "windows"
	FirewallInputOperatingSystemMac     FirewallInputOperatingSystem = "mac"
)

func (r FirewallInputOperatingSystem) IsKnown() bool {
	switch r {
	case FirewallInputOperatingSystemWindows, FirewallInputOperatingSystemMac:
		return true
	}
	return false
}

type FirewallInputParam struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[FirewallInputOperatingSystem] `json:"operating_system,required"`
}

func (r FirewallInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallInputParam) implementsZeroTrustDeviceInputUnionParam() {}

type IntuneInput struct {
	// Compliance Status
	ComplianceStatus IntuneInputComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string          `json:"connection_id,required"`
	JSON         intuneInputJSON `json:"-"`
}

// intuneInputJSON contains the JSON metadata for the struct [IntuneInput]
type intuneInputJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IntuneInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intuneInputJSON) RawJSON() string {
	return r.raw
}

func (r IntuneInput) implementsZeroTrustDeviceInput() {}

// Compliance Status
type IntuneInputComplianceStatus string

const (
	IntuneInputComplianceStatusCompliant     IntuneInputComplianceStatus = "compliant"
	IntuneInputComplianceStatusNoncompliant  IntuneInputComplianceStatus = "noncompliant"
	IntuneInputComplianceStatusUnknown       IntuneInputComplianceStatus = "unknown"
	IntuneInputComplianceStatusNotapplicable IntuneInputComplianceStatus = "notapplicable"
	IntuneInputComplianceStatusIngraceperiod IntuneInputComplianceStatus = "ingraceperiod"
	IntuneInputComplianceStatusError         IntuneInputComplianceStatus = "error"
)

func (r IntuneInputComplianceStatus) IsKnown() bool {
	switch r {
	case IntuneInputComplianceStatusCompliant, IntuneInputComplianceStatusNoncompliant, IntuneInputComplianceStatusUnknown, IntuneInputComplianceStatusNotapplicable, IntuneInputComplianceStatusIngraceperiod, IntuneInputComplianceStatusError:
		return true
	}
	return false
}

type IntuneInputParam struct {
	// Compliance Status
	ComplianceStatus param.Field[IntuneInputComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r IntuneInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntuneInputParam) implementsZeroTrustDeviceInputUnionParam() {}

type KolideInput struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator KolideInputCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string          `json:"issue_count,required"`
	JSON       kolideInputJSON `json:"-"`
}

// kolideInputJSON contains the JSON metadata for the struct [KolideInput]
type kolideInputJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *KolideInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kolideInputJSON) RawJSON() string {
	return r.raw
}

func (r KolideInput) implementsZeroTrustDeviceInput() {}

// Count Operator
type KolideInputCountOperator string

const (
	KolideInputCountOperatorLess            KolideInputCountOperator = "<"
	KolideInputCountOperatorLessOrEquals    KolideInputCountOperator = "<="
	KolideInputCountOperatorGreater         KolideInputCountOperator = ">"
	KolideInputCountOperatorGreaterOrEquals KolideInputCountOperator = ">="
	KolideInputCountOperatorEquals          KolideInputCountOperator = "=="
)

func (r KolideInputCountOperator) IsKnown() bool {
	switch r {
	case KolideInputCountOperatorLess, KolideInputCountOperatorLessOrEquals, KolideInputCountOperatorGreater, KolideInputCountOperatorGreaterOrEquals, KolideInputCountOperatorEquals:
		return true
	}
	return false
}

type KolideInputParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[KolideInputCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r KolideInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r KolideInputParam) implementsZeroTrustDeviceInputUnionParam() {}

type OSVersionInput struct {
	// Operating System
	OperatingSystem OSVersionInputOperatingSystem `json:"operating_system,required"`
	// operator
	Operator OSVersionInputOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string             `json:"os_version_extra"`
	JSON           osVersionInputJSON `json:"-"`
}

// osVersionInputJSON contains the JSON metadata for the struct [OSVersionInput]
type osVersionInputJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OSVersionInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r osVersionInputJSON) RawJSON() string {
	return r.raw
}

func (r OSVersionInput) implementsZeroTrustDeviceInput() {}

// Operating System
type OSVersionInputOperatingSystem string

const (
	OSVersionInputOperatingSystemWindows OSVersionInputOperatingSystem = "windows"
)

func (r OSVersionInputOperatingSystem) IsKnown() bool {
	switch r {
	case OSVersionInputOperatingSystemWindows:
		return true
	}
	return false
}

// operator
type OSVersionInputOperator string

const (
	OSVersionInputOperatorLess            OSVersionInputOperator = "<"
	OSVersionInputOperatorLessOrEquals    OSVersionInputOperator = "<="
	OSVersionInputOperatorGreater         OSVersionInputOperator = ">"
	OSVersionInputOperatorGreaterOrEquals OSVersionInputOperator = ">="
	OSVersionInputOperatorEquals          OSVersionInputOperator = "=="
)

func (r OSVersionInputOperator) IsKnown() bool {
	switch r {
	case OSVersionInputOperatorLess, OSVersionInputOperatorLessOrEquals, OSVersionInputOperatorGreater, OSVersionInputOperatorGreaterOrEquals, OSVersionInputOperatorEquals:
		return true
	}
	return false
}

type OSVersionInputParam struct {
	// Operating System
	OperatingSystem param.Field[OSVersionInputOperatingSystem] `json:"operating_system,required"`
	// operator
	Operator param.Field[OSVersionInputOperator] `json:"operator,required"`
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

func (r OSVersionInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OSVersionInputParam) implementsZeroTrustDeviceInputUnionParam() {}

type SentineloneInput struct {
	// Operating system
	OperatingSystem SentineloneInputOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string               `json:"thumbprint"`
	JSON       sentineloneInputJSON `json:"-"`
}

// sentineloneInputJSON contains the JSON metadata for the struct
// [SentineloneInput]
type sentineloneInputJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SentineloneInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sentineloneInputJSON) RawJSON() string {
	return r.raw
}

func (r SentineloneInput) implementsZeroTrustDeviceInput() {}

// Operating system
type SentineloneInputOperatingSystem string

const (
	SentineloneInputOperatingSystemWindows SentineloneInputOperatingSystem = "windows"
	SentineloneInputOperatingSystemLinux   SentineloneInputOperatingSystem = "linux"
	SentineloneInputOperatingSystemMac     SentineloneInputOperatingSystem = "mac"
)

func (r SentineloneInputOperatingSystem) IsKnown() bool {
	switch r {
	case SentineloneInputOperatingSystemWindows, SentineloneInputOperatingSystemLinux, SentineloneInputOperatingSystemMac:
		return true
	}
	return false
}

type SentineloneInputParam struct {
	// Operating system
	OperatingSystem param.Field[SentineloneInputOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r SentineloneInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SentineloneInputParam) implementsZeroTrustDeviceInputUnionParam() {}

type SentineloneS2sInput struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus SentineloneS2sInputNetworkStatus `json:"network_status"`
	// operator
	Operator SentineloneS2sInputOperator `json:"operator"`
	JSON     sentineloneS2sInputJSON     `json:"-"`
}

// sentineloneS2sInputJSON contains the JSON metadata for the struct
// [SentineloneS2sInput]
type sentineloneS2sInputJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SentineloneS2sInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sentineloneS2sInputJSON) RawJSON() string {
	return r.raw
}

func (r SentineloneS2sInput) implementsZeroTrustDeviceInput() {}

// Network status of device.
type SentineloneS2sInputNetworkStatus string

const (
	SentineloneS2sInputNetworkStatusConnected     SentineloneS2sInputNetworkStatus = "connected"
	SentineloneS2sInputNetworkStatusDisconnected  SentineloneS2sInputNetworkStatus = "disconnected"
	SentineloneS2sInputNetworkStatusDisconnecting SentineloneS2sInputNetworkStatus = "disconnecting"
	SentineloneS2sInputNetworkStatusConnecting    SentineloneS2sInputNetworkStatus = "connecting"
)

func (r SentineloneS2sInputNetworkStatus) IsKnown() bool {
	switch r {
	case SentineloneS2sInputNetworkStatusConnected, SentineloneS2sInputNetworkStatusDisconnected, SentineloneS2sInputNetworkStatusDisconnecting, SentineloneS2sInputNetworkStatusConnecting:
		return true
	}
	return false
}

// operator
type SentineloneS2sInputOperator string

const (
	SentineloneS2sInputOperatorLess            SentineloneS2sInputOperator = "<"
	SentineloneS2sInputOperatorLessOrEquals    SentineloneS2sInputOperator = "<="
	SentineloneS2sInputOperatorGreater         SentineloneS2sInputOperator = ">"
	SentineloneS2sInputOperatorGreaterOrEquals SentineloneS2sInputOperator = ">="
	SentineloneS2sInputOperatorEquals          SentineloneS2sInputOperator = "=="
)

func (r SentineloneS2sInputOperator) IsKnown() bool {
	switch r {
	case SentineloneS2sInputOperatorLess, SentineloneS2sInputOperatorLessOrEquals, SentineloneS2sInputOperatorGreater, SentineloneS2sInputOperatorGreaterOrEquals, SentineloneS2sInputOperatorEquals:
		return true
	}
	return false
}

type SentineloneS2sInputParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[SentineloneS2sInputNetworkStatus] `json:"network_status"`
	// operator
	Operator param.Field[SentineloneS2sInputOperator] `json:"operator"`
}

func (r SentineloneS2sInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SentineloneS2sInputParam) implementsZeroTrustDeviceInputUnionParam() {}

type TaniumInput struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator TaniumInputOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel TaniumInputRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator TaniumInputScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64         `json:"total_score"`
	JSON       taniumInputJSON `json:"-"`
}

// taniumInputJSON contains the JSON metadata for the struct [TaniumInput]
type taniumInputJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TaniumInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taniumInputJSON) RawJSON() string {
	return r.raw
}

func (r TaniumInput) implementsZeroTrustDeviceInput() {}

// Operator to evaluate risk_level or eid_last_seen.
type TaniumInputOperator string

const (
	TaniumInputOperatorLess            TaniumInputOperator = "<"
	TaniumInputOperatorLessOrEquals    TaniumInputOperator = "<="
	TaniumInputOperatorGreater         TaniumInputOperator = ">"
	TaniumInputOperatorGreaterOrEquals TaniumInputOperator = ">="
	TaniumInputOperatorEquals          TaniumInputOperator = "=="
)

func (r TaniumInputOperator) IsKnown() bool {
	switch r {
	case TaniumInputOperatorLess, TaniumInputOperatorLessOrEquals, TaniumInputOperatorGreater, TaniumInputOperatorGreaterOrEquals, TaniumInputOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type TaniumInputRiskLevel string

const (
	TaniumInputRiskLevelLow      TaniumInputRiskLevel = "low"
	TaniumInputRiskLevelMedium   TaniumInputRiskLevel = "medium"
	TaniumInputRiskLevelHigh     TaniumInputRiskLevel = "high"
	TaniumInputRiskLevelCritical TaniumInputRiskLevel = "critical"
)

func (r TaniumInputRiskLevel) IsKnown() bool {
	switch r {
	case TaniumInputRiskLevelLow, TaniumInputRiskLevelMedium, TaniumInputRiskLevelHigh, TaniumInputRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type TaniumInputScoreOperator string

const (
	TaniumInputScoreOperatorLess            TaniumInputScoreOperator = "<"
	TaniumInputScoreOperatorLessOrEquals    TaniumInputScoreOperator = "<="
	TaniumInputScoreOperatorGreater         TaniumInputScoreOperator = ">"
	TaniumInputScoreOperatorGreaterOrEquals TaniumInputScoreOperator = ">="
	TaniumInputScoreOperatorEquals          TaniumInputScoreOperator = "=="
)

func (r TaniumInputScoreOperator) IsKnown() bool {
	switch r {
	case TaniumInputScoreOperatorLess, TaniumInputScoreOperatorLessOrEquals, TaniumInputScoreOperatorGreater, TaniumInputScoreOperatorGreaterOrEquals, TaniumInputScoreOperatorEquals:
		return true
	}
	return false
}

type TaniumInputParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[TaniumInputOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[TaniumInputRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[TaniumInputScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r TaniumInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TaniumInputParam) implementsZeroTrustDeviceInputUnionParam() {}

type UniqueClientIDInput struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem UniqueClientIDInputOperatingSystem `json:"operating_system,required"`
	JSON            uniqueClientIDInputJSON            `json:"-"`
}

// uniqueClientIDInputJSON contains the JSON metadata for the struct
// [UniqueClientIDInput]
type uniqueClientIDInputJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UniqueClientIDInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uniqueClientIDInputJSON) RawJSON() string {
	return r.raw
}

func (r UniqueClientIDInput) implementsZeroTrustDeviceInput() {}

// Operating System
type UniqueClientIDInputOperatingSystem string

const (
	UniqueClientIDInputOperatingSystemAndroid  UniqueClientIDInputOperatingSystem = "android"
	UniqueClientIDInputOperatingSystemIos      UniqueClientIDInputOperatingSystem = "ios"
	UniqueClientIDInputOperatingSystemChromeos UniqueClientIDInputOperatingSystem = "chromeos"
)

func (r UniqueClientIDInputOperatingSystem) IsKnown() bool {
	switch r {
	case UniqueClientIDInputOperatingSystemAndroid, UniqueClientIDInputOperatingSystemIos, UniqueClientIDInputOperatingSystemChromeos:
		return true
	}
	return false
}

type UniqueClientIDInputParam struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[UniqueClientIDInputOperatingSystem] `json:"operating_system,required"`
}

func (r UniqueClientIDInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UniqueClientIDInputParam) implementsZeroTrustDeviceInputUnionParam() {}

type WorkspaceOneInput struct {
	// Compliance Status
	ComplianceStatus WorkspaceOneInputComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                `json:"connection_id,required"`
	JSON         workspaceOneInputJSON `json:"-"`
}

// workspaceOneInputJSON contains the JSON metadata for the struct
// [WorkspaceOneInput]
type workspaceOneInputJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WorkspaceOneInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceOneInputJSON) RawJSON() string {
	return r.raw
}

func (r WorkspaceOneInput) implementsZeroTrustDeviceInput() {}

// Compliance Status
type WorkspaceOneInputComplianceStatus string

const (
	WorkspaceOneInputComplianceStatusCompliant    WorkspaceOneInputComplianceStatus = "compliant"
	WorkspaceOneInputComplianceStatusNoncompliant WorkspaceOneInputComplianceStatus = "noncompliant"
	WorkspaceOneInputComplianceStatusUnknown      WorkspaceOneInputComplianceStatus = "unknown"
)

func (r WorkspaceOneInputComplianceStatus) IsKnown() bool {
	switch r {
	case WorkspaceOneInputComplianceStatusCompliant, WorkspaceOneInputComplianceStatusNoncompliant, WorkspaceOneInputComplianceStatusUnknown:
		return true
	}
	return false
}

type WorkspaceOneInputParam struct {
	// Compliance Status
	ComplianceStatus param.Field[WorkspaceOneInputComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r WorkspaceOneInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkspaceOneInputParam) implementsZeroTrustDeviceInputUnionParam() {}

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
	Input param.Field[DeviceInputUnionParam] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]DeviceMatchParam] `json:"match"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DevicePostureRule     `json:"result,required,nullable"`
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
	Input param.Field[DeviceInputUnionParam] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]DeviceMatchParam] `json:"match"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DevicePostureRule     `json:"result,required,nullable"`
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   DevicePostureDeleteResponse `json:"result,required,nullable"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DevicePostureRule     `json:"result,required,nullable"`
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
