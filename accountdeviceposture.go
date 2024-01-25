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

// AccountDevicePostureService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDevicePostureService]
// method instead.
type AccountDevicePostureService struct {
	Options      []option.RequestOption
	Integrations *AccountDevicePostureIntegrationService
}

// NewAccountDevicePostureService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDevicePostureService(opts ...option.RequestOption) (r *AccountDevicePostureService) {
	r = &AccountDevicePostureService{}
	r.Options = opts
	r.Integrations = NewAccountDevicePostureIntegrationService(opts...)
	return
}

// Fetches a single device posture rule.
func (r *AccountDevicePostureService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDevicePostureGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a device posture rule.
func (r *AccountDevicePostureService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountDevicePostureUpdateParams, opts ...option.RequestOption) (res *AccountDevicePostureUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a device posture rule.
func (r *AccountDevicePostureService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDevicePostureDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new device posture rule.
func (r *AccountDevicePostureService) DevicePostureRulesNewDevicePostureRule(ctx context.Context, identifier interface{}, body AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParams, opts ...option.RequestOption) (res *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches device posture rules for a Zero Trust account.
func (r *AccountDevicePostureService) DevicePostureRulesListDevicePostureRules(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountDevicePostureGetResponse struct {
	Errors   []AccountDevicePostureGetResponseError   `json:"errors"`
	Messages []AccountDevicePostureGetResponseMessage `json:"messages"`
	Result   AccountDevicePostureGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDevicePostureGetResponseSuccess `json:"success"`
	JSON    accountDevicePostureGetResponseJSON    `json:"-"`
}

// accountDevicePostureGetResponseJSON contains the JSON metadata for the struct
// [AccountDevicePostureGetResponse]
type accountDevicePostureGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureGetResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountDevicePostureGetResponseErrorJSON `json:"-"`
}

// accountDevicePostureGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountDevicePostureGetResponseError]
type accountDevicePostureGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureGetResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountDevicePostureGetResponseMessageJSON `json:"-"`
}

// accountDevicePostureGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountDevicePostureGetResponseMessage]
type accountDevicePostureGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureGetResponseResult struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input AccountDevicePostureGetResponseResultInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []AccountDevicePostureGetResponseResultMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type AccountDevicePostureGetResponseResultType `json:"type"`
	JSON accountDevicePostureGetResponseResultJSON `json:"-"`
}

// accountDevicePostureGetResponseResultJSON contains the JSON metadata for the
// struct [AccountDevicePostureGetResponseResult]
type accountDevicePostureGetResponseResultJSON struct {
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

func (r *AccountDevicePostureGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhFileInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhUniqueClientIDInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhDomainJoinedInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhFirewallInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhCarbonblackInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhDiskEncryptionInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhApplicationInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhClientCertificateInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhWorkspaceOneInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequest],
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequest] or
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequest].
type AccountDevicePostureGetResponseResultInput interface {
	implementsAccountDevicePostureGetResponseResultInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountDevicePostureGetResponseResultInput)(nil)).Elem(), "")
}

type AccountDevicePostureGetResponseResultInputZR7Sv6YhFileInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureGetResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                 `json:"thumbprint"`
	JSON       accountDevicePostureGetResponseResultInputZr7Sv6YhFileInputRequestJSON `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhFileInputRequestJSON contains
// the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhFileInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhFileInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Operating system
type AccountDevicePostureGetResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhFileInputRequestOperatingSystemWindows AccountDevicePostureGetResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem = "windows"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhFileInputRequestOperatingSystemLinux   AccountDevicePostureGetResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem = "linux"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhFileInputRequestOperatingSystemMac     AccountDevicePostureGetResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureGetResponseResultInputZR7Sv6YhUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem AccountDevicePostureGetResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            accountDevicePostureGetResponseResultInputZr7Sv6YhUniqueClientIDInputRequestJSON            `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhUniqueClientIDInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhUniqueClientIDInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Operating System
type AccountDevicePostureGetResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemAndroid  AccountDevicePostureGetResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "android"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemIos      AccountDevicePostureGetResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "ios"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemChromeos AccountDevicePostureGetResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type AccountDevicePostureGetResponseResultInputZR7Sv6YhDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem AccountDevicePostureGetResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                         `json:"domain"`
	JSON   accountDevicePostureGetResponseResultInputZr7Sv6YhDomainJoinedInputRequestJSON `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhDomainJoinedInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhDomainJoinedInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Operating System
type AccountDevicePostureGetResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystemWindows AccountDevicePostureGetResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem = "windows"
)

type AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequest struct {
	// Operating System
	OperatingSystem AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision string `json:"os_distro_revision"`
	// Product Verison Extra that Mac OS uses (mac only)
	OsVersionExtra string                                                                      `json:"os_version_extra"`
	JSON           accountDevicePostureGetResponseResultInputZr7Sv6YhOsVersionInputRequestJSON `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhOsVersionInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhOsVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Operating System
type AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystem string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystemWindows AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperator string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown5 AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = "<"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown6 AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = "<="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown7 AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = ">"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown8 AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = ">="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown9 AccountDevicePostureGetResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = "=="
)

type AccountDevicePostureGetResponseResultInputZR7Sv6YhFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem AccountDevicePostureGetResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            accountDevicePostureGetResponseResultInputZr7Sv6YhFirewallInputRequestJSON            `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhFirewallInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhFirewallInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhFirewallInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Operating System
type AccountDevicePostureGetResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystemWindows AccountDevicePostureGetResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem = "windows"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystemMac     AccountDevicePostureGetResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                        `json:"thumbprint"`
	JSON       accountDevicePostureGetResponseResultInputZr7Sv6YhSentineloneInputRequestJSON `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhSentineloneInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Operating system
type AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystemWindows AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "windows"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystemLinux   AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "linux"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystemMac     AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureGetResponseResultInputZR7Sv6YhCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureGetResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                        `json:"thumbprint"`
	JSON       accountDevicePostureGetResponseResultInputZr7Sv6YhCarbonblackInputRequestJSON `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhCarbonblackInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhCarbonblackInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhCarbonblackInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Operating system
type AccountDevicePostureGetResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystemWindows AccountDevicePostureGetResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "windows"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystemLinux   AccountDevicePostureGetResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "linux"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystemMac     AccountDevicePostureGetResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureGetResponseResultInputZR7Sv6YhDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                             `json:"requireAll"`
	JSON       accountDevicePostureGetResponseResultInputZr7Sv6YhDiskEncryptionInputRequestJSON `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhDiskEncryptionInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhDiskEncryptionInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

type AccountDevicePostureGetResponseResultInputZR7Sv6YhApplicationInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureGetResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                        `json:"thumbprint"`
	JSON       accountDevicePostureGetResponseResultInputZr7Sv6YhApplicationInputRequestJSON `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhApplicationInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhApplicationInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhApplicationInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Operating system
type AccountDevicePostureGetResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystemWindows AccountDevicePostureGetResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem = "windows"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystemLinux   AccountDevicePostureGetResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem = "linux"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystemMac     AccountDevicePostureGetResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureGetResponseResultInputZR7Sv6YhClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                              `json:"cn,required"`
	JSON accountDevicePostureGetResponseResultInputZr7Sv6YhClientCertificateInputRequestJSON `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhClientCertificateInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhClientCertificateInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

type AccountDevicePostureGetResponseResultInputZR7Sv6YhWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus AccountDevicePostureGetResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                         `json:"connection_id,required"`
	JSON         accountDevicePostureGetResponseResultInputZr7Sv6YhWorkspaceOneInputRequestJSON `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhWorkspaceOneInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhWorkspaceOneInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Compliance Status
type AccountDevicePostureGetResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusCompliant    AccountDevicePostureGetResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "compliant"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusNoncompliant AccountDevicePostureGetResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusUnknown      AccountDevicePostureGetResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Operator
	Operator AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	Os string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            accountDevicePostureGetResponseResultInputZr7Sv6YhCrowdstrikeInputRequestJSON            `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhCrowdstrikeInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhCrowdstrikeInputRequestJSON struct {
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

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Operator
type AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown15 AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = "<"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown16 AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = "<="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown17 AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = ">"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown18 AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = ">="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown19 AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown25 AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "<"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown26 AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "<="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown27 AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = ">"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown28 AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = ">="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown29 AccountDevicePostureGetResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "=="
)

type AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         accountDevicePostureGetResponseResultInputZr7Sv6YhIntuneInputRequestJSON `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhIntuneInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Compliance Status
type AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusCompliant     AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "compliant"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusNoncompliant  AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "noncompliant"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusUnknown       AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "unknown"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusNotapplicable AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "notapplicable"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusIngraceperiod AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "ingraceperiod"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusError         AccountDevicePostureGetResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "error"
)

type AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                                   `json:"issue_count,required"`
	JSON       accountDevicePostureGetResponseResultInputZr7Sv6YhKolideInputRequestJSON `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhKolideInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Count Operator
type AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequestCountOperator string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown35 AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = "<"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown36 AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = "<="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown37 AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = ">"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown38 AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = ">="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown39 AccountDevicePostureGetResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = "=="
)

type AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                                  `json:"total_score"`
	JSON       accountDevicePostureGetResponseResultInputZr7Sv6YhTaniumInputRequestJSON `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhTaniumInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestOperator string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown45 AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestOperator = "<"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown46 AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestOperator = "<="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown47 AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestOperator = ">"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown48 AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestOperator = ">="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown49 AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelLow      AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "low"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelMedium   AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "medium"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelHigh     AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "high"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelCritical AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown55 AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = "<"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown56 AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = "<="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown57 AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = ">"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown58 AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = ">="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown59 AccountDevicePostureGetResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = "=="
)

type AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     accountDevicePostureGetResponseResultInputZr7Sv6YhSentineloneS2sInputRequestJSON     `json:"-"`
}

// accountDevicePostureGetResponseResultInputZr7Sv6YhSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequest]
type accountDevicePostureGetResponseResultInputZr7Sv6YhSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequest) implementsAccountDevicePostureGetResponseResultInput() {
}

// Network status of device.
type AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusConnected     AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "connected"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusDisconnected  AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "disconnected"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusDisconnecting AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusConnecting    AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator string

const (
	AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown65 AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = "<"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown66 AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = "<="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown67 AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = ">"
	AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown68 AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = ">="
	AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown69 AccountDevicePostureGetResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = "=="
)

type AccountDevicePostureGetResponseResultMatch struct {
	Platform AccountDevicePostureGetResponseResultMatchPlatform `json:"platform"`
	JSON     accountDevicePostureGetResponseResultMatchJSON     `json:"-"`
}

// accountDevicePostureGetResponseResultMatchJSON contains the JSON metadata for
// the struct [AccountDevicePostureGetResponseResultMatch]
type accountDevicePostureGetResponseResultMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureGetResponseResultMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureGetResponseResultMatchPlatform string

const (
	AccountDevicePostureGetResponseResultMatchPlatformWindows AccountDevicePostureGetResponseResultMatchPlatform = "windows"
	AccountDevicePostureGetResponseResultMatchPlatformMac     AccountDevicePostureGetResponseResultMatchPlatform = "mac"
	AccountDevicePostureGetResponseResultMatchPlatformLinux   AccountDevicePostureGetResponseResultMatchPlatform = "linux"
	AccountDevicePostureGetResponseResultMatchPlatformAndroid AccountDevicePostureGetResponseResultMatchPlatform = "android"
	AccountDevicePostureGetResponseResultMatchPlatformIos     AccountDevicePostureGetResponseResultMatchPlatform = "ios"
)

// The type of device posture rule.
type AccountDevicePostureGetResponseResultType string

const (
	AccountDevicePostureGetResponseResultTypeFile              AccountDevicePostureGetResponseResultType = "file"
	AccountDevicePostureGetResponseResultTypeApplication       AccountDevicePostureGetResponseResultType = "application"
	AccountDevicePostureGetResponseResultTypeTanium            AccountDevicePostureGetResponseResultType = "tanium"
	AccountDevicePostureGetResponseResultTypeGateway           AccountDevicePostureGetResponseResultType = "gateway"
	AccountDevicePostureGetResponseResultTypeWarp              AccountDevicePostureGetResponseResultType = "warp"
	AccountDevicePostureGetResponseResultTypeDiskEncryption    AccountDevicePostureGetResponseResultType = "disk_encryption"
	AccountDevicePostureGetResponseResultTypeSentinelone       AccountDevicePostureGetResponseResultType = "sentinelone"
	AccountDevicePostureGetResponseResultTypeCarbonblack       AccountDevicePostureGetResponseResultType = "carbonblack"
	AccountDevicePostureGetResponseResultTypeFirewall          AccountDevicePostureGetResponseResultType = "firewall"
	AccountDevicePostureGetResponseResultTypeOsVersion         AccountDevicePostureGetResponseResultType = "os_version"
	AccountDevicePostureGetResponseResultTypeDomainJoined      AccountDevicePostureGetResponseResultType = "domain_joined"
	AccountDevicePostureGetResponseResultTypeClientCertificate AccountDevicePostureGetResponseResultType = "client_certificate"
	AccountDevicePostureGetResponseResultTypeUniqueClientID    AccountDevicePostureGetResponseResultType = "unique_client_id"
	AccountDevicePostureGetResponseResultTypeKolide            AccountDevicePostureGetResponseResultType = "kolide"
	AccountDevicePostureGetResponseResultTypeTaniumS2s         AccountDevicePostureGetResponseResultType = "tanium_s2s"
	AccountDevicePostureGetResponseResultTypeCrowdstrikeS2s    AccountDevicePostureGetResponseResultType = "crowdstrike_s2s"
	AccountDevicePostureGetResponseResultTypeIntune            AccountDevicePostureGetResponseResultType = "intune"
	AccountDevicePostureGetResponseResultTypeWorkspaceOne      AccountDevicePostureGetResponseResultType = "workspace_one"
	AccountDevicePostureGetResponseResultTypeSentineloneS2s    AccountDevicePostureGetResponseResultType = "sentinelone_s2s"
)

// Whether the API call was successful.
type AccountDevicePostureGetResponseSuccess bool

const (
	AccountDevicePostureGetResponseSuccessTrue AccountDevicePostureGetResponseSuccess = true
)

type AccountDevicePostureUpdateResponse struct {
	Errors   []AccountDevicePostureUpdateResponseError   `json:"errors"`
	Messages []AccountDevicePostureUpdateResponseMessage `json:"messages"`
	Result   AccountDevicePostureUpdateResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDevicePostureUpdateResponseSuccess `json:"success"`
	JSON    accountDevicePostureUpdateResponseJSON    `json:"-"`
}

// accountDevicePostureUpdateResponseJSON contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponse]
type accountDevicePostureUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureUpdateResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountDevicePostureUpdateResponseErrorJSON `json:"-"`
}

// accountDevicePostureUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountDevicePostureUpdateResponseError]
type accountDevicePostureUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureUpdateResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountDevicePostureUpdateResponseMessageJSON `json:"-"`
}

// accountDevicePostureUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountDevicePostureUpdateResponseMessage]
type accountDevicePostureUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureUpdateResponseResult struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input AccountDevicePostureUpdateResponseResultInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []AccountDevicePostureUpdateResponseResultMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type AccountDevicePostureUpdateResponseResultType `json:"type"`
	JSON accountDevicePostureUpdateResponseResultJSON `json:"-"`
}

// accountDevicePostureUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountDevicePostureUpdateResponseResult]
type accountDevicePostureUpdateResponseResultJSON struct {
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

func (r *AccountDevicePostureUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFileInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhUniqueClientIDInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDomainJoinedInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFirewallInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCarbonblackInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDiskEncryptionInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhApplicationInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhClientCertificateInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhWorkspaceOneInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequest],
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequest] or
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequest].
type AccountDevicePostureUpdateResponseResultInput interface {
	implementsAccountDevicePostureUpdateResponseResultInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountDevicePostureUpdateResponseResultInput)(nil)).Elem(), "")
}

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFileInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                    `json:"thumbprint"`
	JSON       accountDevicePostureUpdateResponseResultInputZr7Sv6YhFileInputRequestJSON `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhFileInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFileInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFileInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Operating system
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFileInputRequestOperatingSystemWindows AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem = "windows"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFileInputRequestOperatingSystemLinux   AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem = "linux"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFileInputRequestOperatingSystemMac     AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem AccountDevicePostureUpdateResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            accountDevicePostureUpdateResponseResultInputZr7Sv6YhUniqueClientIDInputRequestJSON            `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhUniqueClientIDInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhUniqueClientIDInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Operating System
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemAndroid  AccountDevicePostureUpdateResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "android"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemIos      AccountDevicePostureUpdateResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "ios"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemChromeos AccountDevicePostureUpdateResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                            `json:"domain"`
	JSON   accountDevicePostureUpdateResponseResultInputZr7Sv6YhDomainJoinedInputRequestJSON `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDomainJoinedInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDomainJoinedInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Operating System
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystemWindows AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem = "windows"
)

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequest struct {
	// Operating System
	OperatingSystem AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision string `json:"os_distro_revision"`
	// Product Verison Extra that Mac OS uses (mac only)
	OsVersionExtra string                                                                         `json:"os_version_extra"`
	JSON           accountDevicePostureUpdateResponseResultInputZr7Sv6YhOsVersionInputRequestJSON `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhOsVersionInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhOsVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Operating System
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystemWindows AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperator string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown75 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = "<"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown76 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = "<="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown77 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = ">"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown78 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = ">="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown79 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = "=="
)

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            accountDevicePostureUpdateResponseResultInputZr7Sv6YhFirewallInputRequestJSON            `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhFirewallInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFirewallInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFirewallInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Operating System
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystemWindows AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem = "windows"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystemMac     AccountDevicePostureUpdateResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                           `json:"thumbprint"`
	JSON       accountDevicePostureUpdateResponseResultInputZr7Sv6YhSentineloneInputRequestJSON `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhSentineloneInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Operating system
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystemWindows AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "windows"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystemLinux   AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "linux"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystemMac     AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                           `json:"thumbprint"`
	JSON       accountDevicePostureUpdateResponseResultInputZr7Sv6YhCarbonblackInputRequestJSON `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhCarbonblackInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCarbonblackInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCarbonblackInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Operating system
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystemWindows AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "windows"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystemLinux   AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "linux"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystemMac     AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                                `json:"requireAll"`
	JSON       accountDevicePostureUpdateResponseResultInputZr7Sv6YhDiskEncryptionInputRequestJSON `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDiskEncryptionInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhDiskEncryptionInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhApplicationInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureUpdateResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                           `json:"thumbprint"`
	JSON       accountDevicePostureUpdateResponseResultInputZr7Sv6YhApplicationInputRequestJSON `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhApplicationInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhApplicationInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhApplicationInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Operating system
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystemWindows AccountDevicePostureUpdateResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem = "windows"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystemLinux   AccountDevicePostureUpdateResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem = "linux"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystemMac     AccountDevicePostureUpdateResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                                 `json:"cn,required"`
	JSON accountDevicePostureUpdateResponseResultInputZr7Sv6YhClientCertificateInputRequestJSON `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhClientCertificateInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhClientCertificateInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus AccountDevicePostureUpdateResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                            `json:"connection_id,required"`
	JSON         accountDevicePostureUpdateResponseResultInputZr7Sv6YhWorkspaceOneInputRequestJSON `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhWorkspaceOneInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhWorkspaceOneInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Compliance Status
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusCompliant    AccountDevicePostureUpdateResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "compliant"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusNoncompliant AccountDevicePostureUpdateResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusUnknown      AccountDevicePostureUpdateResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Operator
	Operator AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	Os string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            accountDevicePostureUpdateResponseResultInputZr7Sv6YhCrowdstrikeInputRequestJSON            `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhCrowdstrikeInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhCrowdstrikeInputRequestJSON struct {
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

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Operator
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown85 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = "<"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown86 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = "<="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown87 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = ">"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown88 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = ">="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown89 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown95 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "<"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown96 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "<="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown97 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = ">"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown98 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = ">="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown99 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "=="
)

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                      `json:"connection_id,required"`
	JSON         accountDevicePostureUpdateResponseResultInputZr7Sv6YhIntuneInputRequestJSON `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhIntuneInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Compliance Status
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusCompliant     AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "compliant"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusNoncompliant  AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "noncompliant"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusUnknown       AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "unknown"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusNotapplicable AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "notapplicable"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusIngraceperiod AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "ingraceperiod"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusError         AccountDevicePostureUpdateResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "error"
)

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                                      `json:"issue_count,required"`
	JSON       accountDevicePostureUpdateResponseResultInputZr7Sv6YhKolideInputRequestJSON `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhKolideInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Count Operator
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequestCountOperator string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown105 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = "<"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown106 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = "<="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown107 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = ">"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown108 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = ">="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown109 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = "=="
)

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                                     `json:"total_score"`
	JSON       accountDevicePostureUpdateResponseResultInputZr7Sv6YhTaniumInputRequestJSON `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhTaniumInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestOperator string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown115 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestOperator = "<"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown116 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestOperator = "<="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown117 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestOperator = ">"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown118 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestOperator = ">="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown119 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelLow      AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "low"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelMedium   AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "medium"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelHigh     AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "high"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelCritical AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown125 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = "<"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown126 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = "<="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown127 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = ">"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown128 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = ">="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown129 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = "=="
)

type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     accountDevicePostureUpdateResponseResultInputZr7Sv6YhSentineloneS2sInputRequestJSON     `json:"-"`
}

// accountDevicePostureUpdateResponseResultInputZr7Sv6YhSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequest]
type accountDevicePostureUpdateResponseResultInputZr7Sv6YhSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequest) implementsAccountDevicePostureUpdateResponseResultInput() {
}

// Network status of device.
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusConnected     AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "connected"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusDisconnected  AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "disconnected"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusDisconnecting AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusConnecting    AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator string

const (
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown135 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = "<"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown136 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = "<="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown137 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = ">"
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown138 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = ">="
	AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown139 AccountDevicePostureUpdateResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = "=="
)

type AccountDevicePostureUpdateResponseResultMatch struct {
	Platform AccountDevicePostureUpdateResponseResultMatchPlatform `json:"platform"`
	JSON     accountDevicePostureUpdateResponseResultMatchJSON     `json:"-"`
}

// accountDevicePostureUpdateResponseResultMatchJSON contains the JSON metadata for
// the struct [AccountDevicePostureUpdateResponseResultMatch]
type accountDevicePostureUpdateResponseResultMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureUpdateResponseResultMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureUpdateResponseResultMatchPlatform string

const (
	AccountDevicePostureUpdateResponseResultMatchPlatformWindows AccountDevicePostureUpdateResponseResultMatchPlatform = "windows"
	AccountDevicePostureUpdateResponseResultMatchPlatformMac     AccountDevicePostureUpdateResponseResultMatchPlatform = "mac"
	AccountDevicePostureUpdateResponseResultMatchPlatformLinux   AccountDevicePostureUpdateResponseResultMatchPlatform = "linux"
	AccountDevicePostureUpdateResponseResultMatchPlatformAndroid AccountDevicePostureUpdateResponseResultMatchPlatform = "android"
	AccountDevicePostureUpdateResponseResultMatchPlatformIos     AccountDevicePostureUpdateResponseResultMatchPlatform = "ios"
)

// The type of device posture rule.
type AccountDevicePostureUpdateResponseResultType string

const (
	AccountDevicePostureUpdateResponseResultTypeFile              AccountDevicePostureUpdateResponseResultType = "file"
	AccountDevicePostureUpdateResponseResultTypeApplication       AccountDevicePostureUpdateResponseResultType = "application"
	AccountDevicePostureUpdateResponseResultTypeTanium            AccountDevicePostureUpdateResponseResultType = "tanium"
	AccountDevicePostureUpdateResponseResultTypeGateway           AccountDevicePostureUpdateResponseResultType = "gateway"
	AccountDevicePostureUpdateResponseResultTypeWarp              AccountDevicePostureUpdateResponseResultType = "warp"
	AccountDevicePostureUpdateResponseResultTypeDiskEncryption    AccountDevicePostureUpdateResponseResultType = "disk_encryption"
	AccountDevicePostureUpdateResponseResultTypeSentinelone       AccountDevicePostureUpdateResponseResultType = "sentinelone"
	AccountDevicePostureUpdateResponseResultTypeCarbonblack       AccountDevicePostureUpdateResponseResultType = "carbonblack"
	AccountDevicePostureUpdateResponseResultTypeFirewall          AccountDevicePostureUpdateResponseResultType = "firewall"
	AccountDevicePostureUpdateResponseResultTypeOsVersion         AccountDevicePostureUpdateResponseResultType = "os_version"
	AccountDevicePostureUpdateResponseResultTypeDomainJoined      AccountDevicePostureUpdateResponseResultType = "domain_joined"
	AccountDevicePostureUpdateResponseResultTypeClientCertificate AccountDevicePostureUpdateResponseResultType = "client_certificate"
	AccountDevicePostureUpdateResponseResultTypeUniqueClientID    AccountDevicePostureUpdateResponseResultType = "unique_client_id"
	AccountDevicePostureUpdateResponseResultTypeKolide            AccountDevicePostureUpdateResponseResultType = "kolide"
	AccountDevicePostureUpdateResponseResultTypeTaniumS2s         AccountDevicePostureUpdateResponseResultType = "tanium_s2s"
	AccountDevicePostureUpdateResponseResultTypeCrowdstrikeS2s    AccountDevicePostureUpdateResponseResultType = "crowdstrike_s2s"
	AccountDevicePostureUpdateResponseResultTypeIntune            AccountDevicePostureUpdateResponseResultType = "intune"
	AccountDevicePostureUpdateResponseResultTypeWorkspaceOne      AccountDevicePostureUpdateResponseResultType = "workspace_one"
	AccountDevicePostureUpdateResponseResultTypeSentineloneS2s    AccountDevicePostureUpdateResponseResultType = "sentinelone_s2s"
)

// Whether the API call was successful.
type AccountDevicePostureUpdateResponseSuccess bool

const (
	AccountDevicePostureUpdateResponseSuccessTrue AccountDevicePostureUpdateResponseSuccess = true
)

type AccountDevicePostureDeleteResponse struct {
	Errors   []AccountDevicePostureDeleteResponseError   `json:"errors"`
	Messages []AccountDevicePostureDeleteResponseMessage `json:"messages"`
	Result   AccountDevicePostureDeleteResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDevicePostureDeleteResponseSuccess `json:"success"`
	JSON    accountDevicePostureDeleteResponseJSON    `json:"-"`
}

// accountDevicePostureDeleteResponseJSON contains the JSON metadata for the struct
// [AccountDevicePostureDeleteResponse]
type accountDevicePostureDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureDeleteResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountDevicePostureDeleteResponseErrorJSON `json:"-"`
}

// accountDevicePostureDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountDevicePostureDeleteResponseError]
type accountDevicePostureDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureDeleteResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountDevicePostureDeleteResponseMessageJSON `json:"-"`
}

// accountDevicePostureDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountDevicePostureDeleteResponseMessage]
type accountDevicePostureDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureDeleteResponseResult struct {
	// API UUID.
	ID   string                                       `json:"id"`
	JSON accountDevicePostureDeleteResponseResultJSON `json:"-"`
}

// accountDevicePostureDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountDevicePostureDeleteResponseResult]
type accountDevicePostureDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePostureDeleteResponseSuccess bool

const (
	AccountDevicePostureDeleteResponseSuccessTrue AccountDevicePostureDeleteResponseSuccess = true
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponse struct {
	Errors   []AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseError   `json:"errors"`
	Messages []AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseMessage `json:"messages"`
	Result   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseSuccess `json:"success"`
	JSON    accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseJSON    `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseJSON contains
// the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponse]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseError struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseErrorJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseError]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseMessage struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseMessageJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseMessage]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResult struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType `json:"type"`
	JSON accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResult]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultJSON struct {
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

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFileInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhUniqueClientIDInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDomainJoinedInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFirewallInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCarbonblackInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDiskEncryptionInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhApplicationInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhClientCertificateInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhWorkspaceOneInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequest]
// or
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequest].
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput interface {
	implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput)(nil)).Elem(), "")
}

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFileInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                    `json:"thumbprint"`
	JSON       accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhFileInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhFileInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFileInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFileInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Operating system
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFileInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFileInputRequestOperatingSystemLinux   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem = "linux"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFileInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhUniqueClientIDInputRequestJSON            `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhUniqueClientIDInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhUniqueClientIDInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Operating System
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemAndroid  AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "android"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemIos      AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "ios"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemChromeos AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                                                            `json:"domain"`
	JSON   accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhDomainJoinedInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDomainJoinedInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDomainJoinedInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Operating System
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem = "windows"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequest struct {
	// Operating System
	OperatingSystem AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision string `json:"os_distro_revision"`
	// Product Verison Extra that Mac OS uses (mac only)
	OsVersionExtra string                                                                                                         `json:"os_version_extra"`
	JSON           accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhOsVersionInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhOsVersionInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhOsVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Operating System
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown145 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown146 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown147 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown148 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown149 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = "=="
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhFirewallInputRequestJSON            `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhFirewallInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFirewallInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFirewallInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Operating System
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                           `json:"thumbprint"`
	JSON       accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhSentineloneInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhSentineloneInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Operating system
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystemLinux   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "linux"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                           `json:"thumbprint"`
	JSON       accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhCarbonblackInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhCarbonblackInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCarbonblackInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCarbonblackInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Operating system
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystemLinux   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "linux"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                                                                `json:"requireAll"`
	JSON       accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhDiskEncryptionInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDiskEncryptionInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhDiskEncryptionInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhApplicationInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                           `json:"thumbprint"`
	JSON       accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhApplicationInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhApplicationInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhApplicationInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhApplicationInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Operating system
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystemLinux   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem = "linux"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                                                                 `json:"cn,required"`
	JSON accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhClientCertificateInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhClientCertificateInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhClientCertificateInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                                                            `json:"connection_id,required"`
	JSON         accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhWorkspaceOneInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhWorkspaceOneInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhWorkspaceOneInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Compliance Status
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusCompliant    AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "compliant"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusNoncompliant AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusUnknown      AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Operator
	Operator AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	Os string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhCrowdstrikeInputRequestJSON            `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhCrowdstrikeInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhCrowdstrikeInputRequestJSON struct {
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

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Operator
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown155 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown156 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown157 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown158 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown159 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown165 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown166 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown167 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown168 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown169 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "=="
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                                                      `json:"connection_id,required"`
	JSON         accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhIntuneInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhIntuneInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Compliance Status
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusCompliant     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "compliant"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusNoncompliant  AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "noncompliant"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusUnknown       AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "unknown"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusNotapplicable AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "notapplicable"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusIngraceperiod AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "ingraceperiod"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusError         AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "error"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                                                                      `json:"issue_count,required"`
	JSON       accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhKolideInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhKolideInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Count Operator
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequestCountOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown175 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown176 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown177 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown178 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown179 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = "=="
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                                                                     `json:"total_score"`
	JSON       accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhTaniumInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhTaniumInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown185 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown186 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown187 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown188 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown189 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelLow      AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "low"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelMedium   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "medium"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelHigh     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "high"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelCritical AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown195 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown196 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown197 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown198 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown199 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = "=="
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhSentineloneS2sInputRequestJSON     `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequest]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZr7Sv6YhSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInput() {
}

// Network status of device.
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusConnected     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "connected"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusDisconnected  AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "disconnected"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusDisconnecting AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusConnecting    AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown205 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown206 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown207 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown208 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown209 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = "=="
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatch struct {
	Platform AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchPlatform `json:"platform"`
	JSON     accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchJSON     `json:"-"`
}

// accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatch]
type accountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchPlatform string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchPlatformWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchPlatform = "windows"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchPlatformMac     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchPlatform = "mac"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchPlatformLinux   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchPlatform = "linux"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchPlatformAndroid AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchPlatform = "android"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchPlatformIos     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultMatchPlatform = "ios"
)

// The type of device posture rule.
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeFile              AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "file"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeApplication       AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "application"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeTanium            AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "tanium"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeGateway           AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "gateway"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeWarp              AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "warp"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeDiskEncryption    AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "disk_encryption"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeSentinelone       AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "sentinelone"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeCarbonblack       AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "carbonblack"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeFirewall          AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "firewall"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeOsVersion         AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "os_version"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeDomainJoined      AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "domain_joined"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeClientCertificate AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "client_certificate"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeUniqueClientID    AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "unique_client_id"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeKolide            AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "kolide"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeTaniumS2s         AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "tanium_s2s"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeCrowdstrikeS2s    AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "crowdstrike_s2s"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeIntune            AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "intune"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeWorkspaceOne      AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "workspace_one"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultTypeSentineloneS2s    AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseResultType = "sentinelone_s2s"
)

// Whether the API call was successful.
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseSuccess bool

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseSuccessTrue AccountDevicePostureDevicePostureRulesNewDevicePostureRuleResponseSuccess = true
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponse struct {
	Errors     []AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseError    `json:"errors"`
	Messages   []AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseMessage  `json:"messages"`
	Result     []AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResult   `json:"result"`
	ResultInfo AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseSuccess `json:"success"`
	JSON    accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseJSON    `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponse]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseError struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseErrorJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseError]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseMessage struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseMessageJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseMessage]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResult struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType `json:"type"`
	JSON accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResult]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultJSON struct {
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

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
//
// Union satisfied by
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFileInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhUniqueClientIDInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDomainJoinedInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFirewallInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCarbonblackInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDiskEncryptionInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhApplicationInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhClientCertificateInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhWorkspaceOneInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequest],
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequest]
// or
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequest].
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput interface {
	implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput)(nil)).Elem(), "")
}

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFileInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                      `json:"thumbprint"`
	JSON       accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhFileInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhFileInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFileInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFileInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Operating system
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFileInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFileInputRequestOperatingSystemLinux   AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem = "linux"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFileInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFileInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhUniqueClientIDInputRequestJSON            `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhUniqueClientIDInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhUniqueClientIDInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhUniqueClientIDInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Operating System
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemAndroid  AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "android"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemIos      AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "ios"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemChromeos AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                                                                                              `json:"domain"`
	JSON   accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhDomainJoinedInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhDomainJoinedInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDomainJoinedInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDomainJoinedInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Operating System
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem = "windows"
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequest struct {
	// Operating System
	OperatingSystem AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision string `json:"os_distro_revision"`
	// Product Verison Extra that Mac OS uses (mac only)
	OsVersionExtra string                                                                                                           `json:"os_version_extra"`
	JSON           accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhOsVersionInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhOsVersionInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhOsVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Operating System
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperator string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown215 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = "<"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown216 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = "<="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown217 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = ">"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown218 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = ">="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperatorUnknown219 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhOsVersionInputRequestOperator = "=="
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhFirewallInputRequestJSON            `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhFirewallInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFirewallInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFirewallInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Operating System
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhFirewallInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                             `json:"thumbprint"`
	JSON       accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhSentineloneInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhSentineloneInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Operating system
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystemLinux   AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "linux"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                             `json:"thumbprint"`
	JSON       accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhCarbonblackInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhCarbonblackInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCarbonblackInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCarbonblackInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Operating system
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystemLinux   AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "linux"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                                                                                  `json:"requireAll"`
	JSON       accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhDiskEncryptionInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhDiskEncryptionInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDiskEncryptionInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhDiskEncryptionInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhApplicationInputRequest struct {
	// Operating system
	OperatingSystem AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                                                                                             `json:"thumbprint"`
	JSON       accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhApplicationInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhApplicationInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhApplicationInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhApplicationInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Operating system
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystemLinux   AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem = "linux"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhApplicationInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                                                                                   `json:"cn,required"`
	JSON accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhClientCertificateInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhClientCertificateInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhClientCertificateInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhClientCertificateInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                                                              `json:"connection_id,required"`
	JSON         accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhWorkspaceOneInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhWorkspaceOneInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhWorkspaceOneInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhWorkspaceOneInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Compliance Status
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusCompliant    AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "compliant"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusNoncompliant AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusUnknown      AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Operator
	Operator AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	Os string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhCrowdstrikeInputRequestJSON            `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhCrowdstrikeInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhCrowdstrikeInputRequestJSON struct {
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

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Operator
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown225 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = "<"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown226 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = "<="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown227 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = ">"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown228 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = ">="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown229 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown235 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "<"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown236 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "<="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown237 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = ">"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown238 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = ">="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown239 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "=="
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                                                                                        `json:"connection_id,required"`
	JSON         accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhIntuneInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhIntuneInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Compliance Status
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusCompliant     AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "compliant"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusNoncompliant  AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "noncompliant"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusUnknown       AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "unknown"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusNotapplicable AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "notapplicable"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusIngraceperiod AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "ingraceperiod"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatusError         AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhIntuneInputRequestComplianceStatus = "error"
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                                                                                        `json:"issue_count,required"`
	JSON       accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhKolideInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhKolideInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Count Operator
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequestCountOperator string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown245 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = "<"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown246 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = "<="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown247 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = ">"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown248 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = ">="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequestCountOperatorUnknown249 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhKolideInputRequestCountOperator = "=="
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                                                                                       `json:"total_score"`
	JSON       accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhTaniumInputRequestJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhTaniumInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestOperator string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown255 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestOperator = "<"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown256 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestOperator = "<="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown257 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestOperator = ">"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown258 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestOperator = ">="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestOperatorUnknown259 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelLow      AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "low"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelMedium   AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "medium"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelHigh     AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "high"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevelCritical AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown265 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = "<"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown266 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = "<="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown267 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = ">"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown268 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = ">="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown269 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhTaniumInputRequestScoreOperator = "=="
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhSentineloneS2sInputRequestJSON     `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhSentineloneS2sInputRequestJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequest]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZr7Sv6YhSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequest) implementsAccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInput() {
}

// Network status of device.
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusConnected     AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "connected"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusDisconnected  AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "disconnected"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusDisconnecting AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusConnecting    AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown275 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = "<"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown276 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = "<="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown277 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = ">"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown278 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = ">="
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown279 AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInputZR7Sv6YhSentineloneS2sInputRequestOperator = "=="
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatch struct {
	Platform AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchPlatform `json:"platform"`
	JSON     accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchJSON     `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatch]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchPlatform string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchPlatformWindows AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchPlatform = "windows"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchPlatformMac     AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchPlatform = "mac"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchPlatformLinux   AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchPlatform = "linux"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchPlatformAndroid AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchPlatform = "android"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchPlatformIos     AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultMatchPlatform = "ios"
)

// The type of device posture rule.
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType string

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeFile              AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "file"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeApplication       AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "application"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeTanium            AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "tanium"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeGateway           AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "gateway"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeWarp              AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "warp"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeDiskEncryption    AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "disk_encryption"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeSentinelone       AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "sentinelone"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeCarbonblack       AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "carbonblack"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeFirewall          AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "firewall"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeOsVersion         AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "os_version"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeDomainJoined      AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "domain_joined"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeClientCertificate AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "client_certificate"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeUniqueClientID    AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "unique_client_id"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeKolide            AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "kolide"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeTaniumS2s         AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "tanium_s2s"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeCrowdstrikeS2s    AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "crowdstrike_s2s"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeIntune            AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "intune"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeWorkspaceOne      AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "workspace_one"
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultTypeSentineloneS2s    AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultType = "sentinelone_s2s"
)

type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                            `json:"total_count"`
	JSON       accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInfoJSON `json:"-"`
}

// accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInfo]
type accountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseSuccess bool

const (
	AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseSuccessTrue AccountDevicePostureDevicePostureRulesListDevicePostureRulesResponseSuccess = true
)

type AccountDevicePostureUpdateParams struct {
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[AccountDevicePostureUpdateParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[AccountDevicePostureUpdateParamsInput] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]AccountDevicePostureUpdateParamsMatch] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r AccountDevicePostureUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type AccountDevicePostureUpdateParamsType string

const (
	AccountDevicePostureUpdateParamsTypeFile              AccountDevicePostureUpdateParamsType = "file"
	AccountDevicePostureUpdateParamsTypeApplication       AccountDevicePostureUpdateParamsType = "application"
	AccountDevicePostureUpdateParamsTypeTanium            AccountDevicePostureUpdateParamsType = "tanium"
	AccountDevicePostureUpdateParamsTypeGateway           AccountDevicePostureUpdateParamsType = "gateway"
	AccountDevicePostureUpdateParamsTypeWarp              AccountDevicePostureUpdateParamsType = "warp"
	AccountDevicePostureUpdateParamsTypeDiskEncryption    AccountDevicePostureUpdateParamsType = "disk_encryption"
	AccountDevicePostureUpdateParamsTypeSentinelone       AccountDevicePostureUpdateParamsType = "sentinelone"
	AccountDevicePostureUpdateParamsTypeCarbonblack       AccountDevicePostureUpdateParamsType = "carbonblack"
	AccountDevicePostureUpdateParamsTypeFirewall          AccountDevicePostureUpdateParamsType = "firewall"
	AccountDevicePostureUpdateParamsTypeOsVersion         AccountDevicePostureUpdateParamsType = "os_version"
	AccountDevicePostureUpdateParamsTypeDomainJoined      AccountDevicePostureUpdateParamsType = "domain_joined"
	AccountDevicePostureUpdateParamsTypeClientCertificate AccountDevicePostureUpdateParamsType = "client_certificate"
	AccountDevicePostureUpdateParamsTypeUniqueClientID    AccountDevicePostureUpdateParamsType = "unique_client_id"
	AccountDevicePostureUpdateParamsTypeKolide            AccountDevicePostureUpdateParamsType = "kolide"
	AccountDevicePostureUpdateParamsTypeTaniumS2s         AccountDevicePostureUpdateParamsType = "tanium_s2s"
	AccountDevicePostureUpdateParamsTypeCrowdstrikeS2s    AccountDevicePostureUpdateParamsType = "crowdstrike_s2s"
	AccountDevicePostureUpdateParamsTypeIntune            AccountDevicePostureUpdateParamsType = "intune"
	AccountDevicePostureUpdateParamsTypeWorkspaceOne      AccountDevicePostureUpdateParamsType = "workspace_one"
	AccountDevicePostureUpdateParamsTypeSentineloneS2s    AccountDevicePostureUpdateParamsType = "sentinelone_s2s"
)

// The value to be checked against.
//
// Satisfied by [AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhUniqueClientIDInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhDomainJoinedInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhFirewallInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhCarbonblackInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhDiskEncryptionInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhApplicationInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhClientCertificateInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhWorkspaceOneInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequest],
// [AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequest].
type AccountDevicePostureUpdateParamsInput interface {
	implementsAccountDevicePostureUpdateParamsInput()
}

type AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequest struct {
	// Operating system
	OperatingSystem param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Operating system
type AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequestOperatingSystemWindows AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequestOperatingSystem = "windows"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequestOperatingSystemLinux   AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequestOperatingSystem = "linux"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequestOperatingSystemMac     AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureUpdateParamsInputZR7Sv6YhUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhUniqueClientIDInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Operating System
type AccountDevicePostureUpdateParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemAndroid  AccountDevicePostureUpdateParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "android"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemIos      AccountDevicePostureUpdateParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "ios"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemChromeos AccountDevicePostureUpdateParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type AccountDevicePostureUpdateParamsInputZR7Sv6YhDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhDomainJoinedInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Operating System
type AccountDevicePostureUpdateParamsInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhDomainJoinedInputRequestOperatingSystemWindows AccountDevicePostureUpdateParamsInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem = "windows"
)

type AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// Operator
	Operator param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperator] `json:"operator,required"`
	// Version of OS
	Version param.Field[string] `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Product Verison Extra that Mac OS uses (mac only)
	OsVersionExtra param.Field[string] `json:"os_version_extra"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Operating System
type AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperatingSystemWindows AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperator string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperatorUnknown285 AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperator = "<"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperatorUnknown286 AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperator = "<="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperatorUnknown287 AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperator = ">"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperatorUnknown288 AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperator = ">="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperatorUnknown289 AccountDevicePostureUpdateParamsInputZR7Sv6YhOsVersionInputRequestOperator = "=="
)

type AccountDevicePostureUpdateParamsInputZR7Sv6YhFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhFirewallInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Operating System
type AccountDevicePostureUpdateParamsInputZR7Sv6YhFirewallInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhFirewallInputRequestOperatingSystemWindows AccountDevicePostureUpdateParamsInputZR7Sv6YhFirewallInputRequestOperatingSystem = "windows"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhFirewallInputRequestOperatingSystemMac     AccountDevicePostureUpdateParamsInputZR7Sv6YhFirewallInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneInputRequest struct {
	// Operating system
	OperatingSystem param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Operating system
type AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystemWindows AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "windows"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystemLinux   AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "linux"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystemMac     AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureUpdateParamsInputZR7Sv6YhCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhCarbonblackInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhCarbonblackInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Operating system
type AccountDevicePostureUpdateParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystemWindows AccountDevicePostureUpdateParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "windows"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystemLinux   AccountDevicePostureUpdateParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "linux"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystemMac     AccountDevicePostureUpdateParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureUpdateParamsInputZR7Sv6YhDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhDiskEncryptionInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

type AccountDevicePostureUpdateParamsInputZR7Sv6YhApplicationInputRequest struct {
	// Operating system
	OperatingSystem param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhApplicationInputRequestOperatingSystem] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhApplicationInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhApplicationInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Operating system
type AccountDevicePostureUpdateParamsInputZR7Sv6YhApplicationInputRequestOperatingSystem string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhApplicationInputRequestOperatingSystemWindows AccountDevicePostureUpdateParamsInputZR7Sv6YhApplicationInputRequestOperatingSystem = "windows"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhApplicationInputRequestOperatingSystemLinux   AccountDevicePostureUpdateParamsInputZR7Sv6YhApplicationInputRequestOperatingSystem = "linux"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhApplicationInputRequestOperatingSystemMac     AccountDevicePostureUpdateParamsInputZR7Sv6YhApplicationInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureUpdateParamsInputZR7Sv6YhClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhClientCertificateInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

type AccountDevicePostureUpdateParamsInputZR7Sv6YhWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhWorkspaceOneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhWorkspaceOneInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Compliance Status
type AccountDevicePostureUpdateParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusCompliant    AccountDevicePostureUpdateParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "compliant"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusNoncompliant AccountDevicePostureUpdateParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusUnknown      AccountDevicePostureUpdateParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Operator
	Operator param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator] `json:"operator"`
	// Os Version
	Os param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator] `json:"versionOperator"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Operator
type AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown295 AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator = "<"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown296 AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator = "<="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown297 AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator = ">"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown298 AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator = ">="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown299 AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown305 AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "<"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown306 AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "<="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown307 AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = ">"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown308 AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = ">="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown309 AccountDevicePostureUpdateParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "=="
)

type AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Compliance Status
type AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatusCompliant     AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus = "compliant"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatusNoncompliant  AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus = "noncompliant"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatusUnknown       AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus = "unknown"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatusNotapplicable AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus = "notapplicable"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatusIngraceperiod AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus = "ingraceperiod"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatusError         AccountDevicePostureUpdateParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus = "error"
)

type AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequestCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Count Operator
type AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequestCountOperator string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequestCountOperatorUnknown315 AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequestCountOperator = "<"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequestCountOperatorUnknown316 AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequestCountOperator = "<="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequestCountOperatorUnknown317 AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequestCountOperator = ">"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequestCountOperatorUnknown318 AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequestCountOperator = ">="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequestCountOperatorUnknown319 AccountDevicePostureUpdateParamsInputZR7Sv6YhKolideInputRequestCountOperator = "=="
)

type AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestOperator string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestOperatorUnknown325 AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestOperator = "<"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestOperatorUnknown326 AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestOperator = "<="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestOperatorUnknown327 AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestOperator = ">"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestOperatorUnknown328 AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestOperator = ">="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestOperatorUnknown329 AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestRiskLevel string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestRiskLevelLow      AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestRiskLevel = "low"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestRiskLevelMedium   AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestRiskLevel = "medium"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestRiskLevelHigh     AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestRiskLevel = "high"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestRiskLevelCritical AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestScoreOperator string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown335 AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestScoreOperator = "<"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown336 AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestScoreOperator = "<="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown337 AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestScoreOperator = ">"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown338 AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestScoreOperator = ">="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown339 AccountDevicePostureUpdateParamsInputZR7Sv6YhTaniumInputRequestScoreOperator = "=="
)

type AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus] `json:"network_status"`
	// operator
	Operator param.Field[AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator] `json:"operator"`
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequest) implementsAccountDevicePostureUpdateParamsInput() {
}

// Network status of device.
type AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusConnected     AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "connected"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusDisconnected  AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "disconnected"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusDisconnecting AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusConnecting    AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator string

const (
	AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown345 AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator = "<"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown346 AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator = "<="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown347 AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator = ">"
	AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown348 AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator = ">="
	AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown349 AccountDevicePostureUpdateParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator = "=="
)

type AccountDevicePostureUpdateParamsMatch struct {
	Platform param.Field[AccountDevicePostureUpdateParamsMatchPlatform] `json:"platform"`
}

func (r AccountDevicePostureUpdateParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePostureUpdateParamsMatchPlatform string

const (
	AccountDevicePostureUpdateParamsMatchPlatformWindows AccountDevicePostureUpdateParamsMatchPlatform = "windows"
	AccountDevicePostureUpdateParamsMatchPlatformMac     AccountDevicePostureUpdateParamsMatchPlatform = "mac"
	AccountDevicePostureUpdateParamsMatchPlatformLinux   AccountDevicePostureUpdateParamsMatchPlatform = "linux"
	AccountDevicePostureUpdateParamsMatchPlatformAndroid AccountDevicePostureUpdateParamsMatchPlatform = "android"
	AccountDevicePostureUpdateParamsMatchPlatformIos     AccountDevicePostureUpdateParamsMatchPlatform = "ios"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParams struct {
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatch] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeFile              AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "file"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeApplication       AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "application"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeTanium            AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "tanium"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeGateway           AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "gateway"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeWarp              AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "warp"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeDiskEncryption    AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "disk_encryption"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeSentinelone       AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "sentinelone"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeCarbonblack       AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "carbonblack"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeFirewall          AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "firewall"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeOsVersion         AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "os_version"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeDomainJoined      AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "domain_joined"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeClientCertificate AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "client_certificate"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeUniqueClientID    AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "unique_client_id"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeKolide            AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "kolide"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeTaniumS2s         AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "tanium_s2s"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeCrowdstrikeS2s    AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "crowdstrike_s2s"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeIntune            AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "intune"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeWorkspaceOne      AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "workspace_one"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeSentineloneS2s    AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "sentinelone_s2s"
)

// The value to be checked against.
//
// Satisfied by
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhUniqueClientIDInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhDomainJoinedInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFirewallInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCarbonblackInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhDiskEncryptionInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhApplicationInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhClientCertificateInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhWorkspaceOneInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequest],
// [AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequest].
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput interface {
	implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput()
}

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequest struct {
	// Operating system
	OperatingSystem param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating system
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequestOperatingSystemLinux   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequestOperatingSystem = "linux"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhUniqueClientIDInputRequest struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhUniqueClientIDInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhUniqueClientIDInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating System
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemAndroid  AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "android"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemIos      AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "ios"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystemChromeos AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhDomainJoinedInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhDomainJoinedInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating System
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhDomainJoinedInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhDomainJoinedInputRequestOperatingSystem = "windows"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequest struct {
	// Operating System
	OperatingSystem param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// Operator
	Operator param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperator] `json:"operator,required"`
	// Version of OS
	Version param.Field[string] `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Product Verison Extra that Mac OS uses (mac only)
	OsVersionExtra param.Field[string] `json:"os_version_extra"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating System
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperatingSystem = "windows"
)

// Operator
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperatorUnknown355 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperatorUnknown356 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperatorUnknown357 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperatorUnknown358 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperatorUnknown359 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhOsVersionInputRequestOperator = "=="
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFirewallInputRequest struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFirewallInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFirewallInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating System
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFirewallInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFirewallInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFirewallInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFirewallInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFirewallInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneInputRequest struct {
	// Operating system
	OperatingSystem param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating system
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystemLinux   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "linux"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCarbonblackInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCarbonblackInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating system
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystemLinux   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "linux"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCarbonblackInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhDiskEncryptionInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhDiskEncryptionInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhApplicationInputRequest struct {
	// Operating system
	OperatingSystem param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhApplicationInputRequestOperatingSystem] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhApplicationInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhApplicationInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operating system
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhApplicationInputRequestOperatingSystem string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhApplicationInputRequestOperatingSystemWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhApplicationInputRequestOperatingSystem = "windows"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhApplicationInputRequestOperatingSystemLinux   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhApplicationInputRequestOperatingSystem = "linux"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhApplicationInputRequestOperatingSystemMac     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhApplicationInputRequestOperatingSystem = "mac"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhClientCertificateInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhClientCertificateInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhWorkspaceOneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhWorkspaceOneInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Compliance Status
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusCompliant    AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "compliant"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusNoncompliant AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatusUnknown      AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhWorkspaceOneInputRequestComplianceStatus = "unknown"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Operator
	Operator param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator] `json:"operator"`
	// Os Version
	Os param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator] `json:"versionOperator"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operator
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown365 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown366 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown367 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown368 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestOperatorUnknown369 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestOperator = "=="
)

// Version Operator
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown375 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown376 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown377 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown378 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperatorUnknown379 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhCrowdstrikeInputRequestVersionOperator = "=="
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Compliance Status
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatusCompliant     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus = "compliant"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatusNoncompliant  AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus = "noncompliant"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatusUnknown       AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus = "unknown"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatusNotapplicable AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus = "notapplicable"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatusIngraceperiod AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus = "ingraceperiod"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatusError         AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhIntuneInputRequestComplianceStatus = "error"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequestCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Count Operator
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequestCountOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequestCountOperatorUnknown385 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequestCountOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequestCountOperatorUnknown386 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequestCountOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequestCountOperatorUnknown387 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequestCountOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequestCountOperatorUnknown388 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequestCountOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequestCountOperatorUnknown389 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhKolideInputRequestCountOperator = "=="
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Operator to evaluate risk_level or eid_last_seen.
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestOperatorUnknown395 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestOperatorUnknown396 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestOperatorUnknown397 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestOperatorUnknown398 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestOperatorUnknown399 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestOperator = "=="
)

// For more details on risk level, refer to the Tanium documentation.
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestRiskLevel string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestRiskLevelLow      AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestRiskLevel = "low"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestRiskLevelMedium   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestRiskLevel = "medium"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestRiskLevelHigh     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestRiskLevel = "high"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestRiskLevelCritical AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestRiskLevel = "critical"
)

// Score Operator
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestScoreOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown405 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestScoreOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown406 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestScoreOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown407 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestScoreOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown408 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestScoreOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestScoreOperatorUnknown409 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhTaniumInputRequestScoreOperator = "=="
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus] `json:"network_status"`
	// operator
	Operator param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator] `json:"operator"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequest) implementsAccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput() {
}

// Network status of device.
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusConnected     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "connected"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusDisconnected  AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "disconnected"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusDisconnecting AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatusConnecting    AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestNetworkStatus = "connecting"
)

// operator
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown415 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator = "<"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown416 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator = "<="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown417 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator = ">"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown418 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator = ">="
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestOperatorUnknown419 AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhSentineloneS2sInputRequestOperator = "=="
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatch struct {
	Platform param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform] `json:"platform"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "windows"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformMac     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "mac"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformLinux   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "linux"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformAndroid AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "android"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformIos     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "ios"
)
