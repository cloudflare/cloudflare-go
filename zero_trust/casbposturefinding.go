// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// CasbPostureFindingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbPostureFindingService] method instead.
type CasbPostureFindingService struct {
	Options   []option.RequestOption
	Instances *CasbPostureFindingInstanceService
}

// NewCasbPostureFindingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCasbPostureFindingService(opts ...option.RequestOption) (r *CasbPostureFindingService) {
	r = &CasbPostureFindingService{}
	r.Options = opts
	r.Instances = NewCasbPostureFindingInstanceService(opts...)
	return
}

// List all security findings that have been identified as being problematic. This
// will return a list of findings regardless if they have been ignored or not.
func (r *CasbPostureFindingService) List(ctx context.Context, params CasbPostureFindingListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[CasbPostureFindingListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/findings", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List all security findings that have been identified as being problematic. This
// will return a list of findings regardless if they have been ignored or not.
func (r *CasbPostureFindingService) ListAutoPaging(ctx context.Context, params CasbPostureFindingListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[CasbPostureFindingListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Creates a CSV export for findings and accepts optional filters in the payload.
func (r *CasbPostureFindingService) Export(ctx context.Context, params CasbPostureFindingExportParams, opts ...option.RequestOption) (res *CasbPostureFindingExportResponse, err error) {
	var env CasbPostureFindingExportResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/findings/export", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets a security Finding that has been identified as being problematic.
func (r *CasbPostureFindingService) Get(ctx context.Context, findingID string, query CasbPostureFindingGetParams, opts ...option.RequestOption) (res *CasbPostureFindingGetResponse, err error) {
	var env CasbPostureFindingGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if findingID == "" {
		err = errors.New("missing required finding_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/findings/%s", query.AccountID, findingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Given a list of findings, mark as ignored. Does nothing if Finding is already
// ignored.
func (r *CasbPostureFindingService) Ignore(ctx context.Context, params CasbPostureFindingIgnoreParams, opts ...option.RequestOption) (res *CasbPostureFindingIgnoreResponse, err error) {
	var env CasbPostureFindingIgnoreResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/findings/ignore", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// If a Finding's severity has been changed, reset it back to default value. Does
// nothing if no override exists.
func (r *CasbPostureFindingService) ResetSeverity(ctx context.Context, findingID string, body CasbPostureFindingResetSeverityParams, opts ...option.RequestOption) (res *CasbPostureFindingResetSeverityResponse, err error) {
	var env CasbPostureFindingResetSeverityResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if findingID == "" {
		err = errors.New("missing required finding_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/findings/%s/reset_finding_severity", body.AccountID, findingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update the severity of a Finding. This will update the `severity_override` field
// on the Finding payload with the new severity value.
func (r *CasbPostureFindingService) TuneSeverity(ctx context.Context, findingID string, params CasbPostureFindingTuneSeverityParams, opts ...option.RequestOption) (res *CasbPostureFindingTuneSeverityResponse, err error) {
	var env CasbPostureFindingTuneSeverityResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if findingID == "" {
		err = errors.New("missing required finding_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/findings/%s/tune_finding_severity", params.AccountID, findingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Ability to un-ignore a Finding if it's previously been ignored. Does nothing if
// the Finding is not ignored.
func (r *CasbPostureFindingService) Unignore(ctx context.Context, params CasbPostureFindingUnignoreParams, opts ...option.RequestOption) (res *CasbPostureFindingUnignoreResponse, err error) {
	var env CasbPostureFindingUnignoreResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/findings/unignore", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Aggregated finding information with counts and metadata. This is optimized for
// list API queries and represents a finding along with its instance statistics.
type CasbPostureFindingListResponse struct {
	// Base64 encoded identifier of the security finding.
	ID string `json:"id" api:"required" format:"byte"`
	// Number of active problematic instances identified in the security finding.
	ActiveCount int64 `json:"active_count" api:"required"`
	// Number of archived instances identified in the security finding.
	ArchivedCount int64 `json:"archived_count" api:"required"`
	// Basic finding type information.
	Finding CasbPostureFindingListResponseFinding `json:"finding" api:"required"`
	// Determines if finding is currently ignored.
	Ignored bool `json:"ignored" api:"required"`
	// Number of total (Active or archived) problematic instances identified in the
	// security finding.
	InstanceCount int64 `json:"instance_count" api:"required"`
	// Summary information about an integration.
	Integration CasbPostureFindingListResponseIntegration `json:"integration" api:"required"`
	// Timestamp of the latest affliction date of an active finding.
	LatestAfflictionDate time.Time `json:"latest_affliction_date" api:"required" format:"date-time"`
	// Override information for finding severity.
	SeverityOverride CasbPostureFindingListResponseSeverityOverride `json:"severity_override"`
	JSON             casbPostureFindingListResponseJSON             `json:"-"`
}

// casbPostureFindingListResponseJSON contains the JSON metadata for the struct
// [CasbPostureFindingListResponse]
type casbPostureFindingListResponseJSON struct {
	ID                   apijson.Field
	ActiveCount          apijson.Field
	ArchivedCount        apijson.Field
	Finding              apijson.Field
	Ignored              apijson.Field
	InstanceCount        apijson.Field
	Integration          apijson.Field
	LatestAfflictionDate apijson.Field
	SeverityOverride     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CasbPostureFindingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingListResponseJSON) RawJSON() string {
	return r.raw
}

// Basic finding type information.
type CasbPostureFindingListResponseFinding struct {
	// The unique identifier of the finding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Category information for a finding.
	Category CasbPostureFindingListResponseFindingCategory `json:"category" api:"required"`
	// The name of the finding.
	Name string `json:"name" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingListResponseFindingSeverity `json:"severity" api:"required"`
	// The SaaS/Cloud vendor of the platform with which the finding is associated.
	Vendor string `json:"vendor" api:"required"`
	// Detailed description of the finding.
	Description string `json:"description" api:"nullable"`
	// Remediation guide information for a finding.
	Remediation CasbPostureFindingListResponseFindingRemediation `json:"remediation" api:"nullable"`
	JSON        casbPostureFindingListResponseFindingJSON        `json:"-"`
}

// casbPostureFindingListResponseFindingJSON contains the JSON metadata for the
// struct [CasbPostureFindingListResponseFinding]
type casbPostureFindingListResponseFindingJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	Name        apijson.Field
	Severity    apijson.Field
	Vendor      apijson.Field
	Description apijson.Field
	Remediation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingListResponseFinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingListResponseFindingJSON) RawJSON() string {
	return r.raw
}

// Category information for a finding.
type CasbPostureFindingListResponseFindingCategory struct {
	// The type of the observation.
	Observation CasbPostureFindingListResponseFindingCategoryObservation `json:"observation" api:"required"`
	// The product category.
	Product CasbPostureFindingListResponseFindingCategoryProduct `json:"product" api:"required"`
	// The type of the finding category.
	Type CasbPostureFindingListResponseFindingCategoryType `json:"type" api:"required"`
	JSON casbPostureFindingListResponseFindingCategoryJSON `json:"-"`
}

// casbPostureFindingListResponseFindingCategoryJSON contains the JSON metadata for
// the struct [CasbPostureFindingListResponseFindingCategory]
type casbPostureFindingListResponseFindingCategoryJSON struct {
	Observation apijson.Field
	Product     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingListResponseFindingCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingListResponseFindingCategoryJSON) RawJSON() string {
	return r.raw
}

// The type of the observation.
type CasbPostureFindingListResponseFindingCategoryObservation string

const (
	CasbPostureFindingListResponseFindingCategoryObservationIssue    CasbPostureFindingListResponseFindingCategoryObservation = "Issue"
	CasbPostureFindingListResponseFindingCategoryObservationInsight  CasbPostureFindingListResponseFindingCategoryObservation = "Insight"
	CasbPostureFindingListResponseFindingCategoryObservationActivity CasbPostureFindingListResponseFindingCategoryObservation = "Activity"
)

func (r CasbPostureFindingListResponseFindingCategoryObservation) IsKnown() bool {
	switch r {
	case CasbPostureFindingListResponseFindingCategoryObservationIssue, CasbPostureFindingListResponseFindingCategoryObservationInsight, CasbPostureFindingListResponseFindingCategoryObservationActivity:
		return true
	}
	return false
}

// The product category.
type CasbPostureFindingListResponseFindingCategoryProduct string

const (
	CasbPostureFindingListResponseFindingCategoryProductSaaS  CasbPostureFindingListResponseFindingCategoryProduct = "SaaS"
	CasbPostureFindingListResponseFindingCategoryProductCloud CasbPostureFindingListResponseFindingCategoryProduct = "Cloud"
)

func (r CasbPostureFindingListResponseFindingCategoryProduct) IsKnown() bool {
	switch r {
	case CasbPostureFindingListResponseFindingCategoryProductSaaS, CasbPostureFindingListResponseFindingCategoryProductCloud:
		return true
	}
	return false
}

// The type of the finding category.
type CasbPostureFindingListResponseFindingCategoryType string

const (
	CasbPostureFindingListResponseFindingCategoryTypeContent CasbPostureFindingListResponseFindingCategoryType = "Content"
	CasbPostureFindingListResponseFindingCategoryTypePosture CasbPostureFindingListResponseFindingCategoryType = "Posture"
)

func (r CasbPostureFindingListResponseFindingCategoryType) IsKnown() bool {
	switch r {
	case CasbPostureFindingListResponseFindingCategoryTypeContent, CasbPostureFindingListResponseFindingCategoryTypePosture:
		return true
	}
	return false
}

// The severity level of a finding.
type CasbPostureFindingListResponseFindingSeverity string

const (
	CasbPostureFindingListResponseFindingSeverityCritical CasbPostureFindingListResponseFindingSeverity = "Critical"
	CasbPostureFindingListResponseFindingSeverityHigh     CasbPostureFindingListResponseFindingSeverity = "High"
	CasbPostureFindingListResponseFindingSeverityMedium   CasbPostureFindingListResponseFindingSeverity = "Medium"
	CasbPostureFindingListResponseFindingSeverityLow      CasbPostureFindingListResponseFindingSeverity = "Low"
)

func (r CasbPostureFindingListResponseFindingSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingListResponseFindingSeverityCritical, CasbPostureFindingListResponseFindingSeverityHigh, CasbPostureFindingListResponseFindingSeverityMedium, CasbPostureFindingListResponseFindingSeverityLow:
		return true
	}
	return false
}

// Remediation guide information for a finding.
type CasbPostureFindingListResponseFindingRemediation struct {
	// Remediation Id.
	ID string `json:"id" api:"required" format:"uuid"`
	// Relevant Compliance Frameworks.
	Frameworks []string `json:"frameworks" api:"required"`
	// Remediation guide text.
	Guide string `json:"guide" api:"required"`
	// Description of the potential impact.
	Impact string `json:"impact" api:"required"`
	// I18N Locale.
	Locale string `json:"locale" api:"required"`
	// Description of the threat.
	Threat string                                               `json:"threat" api:"required"`
	JSON   casbPostureFindingListResponseFindingRemediationJSON `json:"-"`
}

// casbPostureFindingListResponseFindingRemediationJSON contains the JSON metadata
// for the struct [CasbPostureFindingListResponseFindingRemediation]
type casbPostureFindingListResponseFindingRemediationJSON struct {
	ID          apijson.Field
	Frameworks  apijson.Field
	Guide       apijson.Field
	Impact      apijson.Field
	Locale      apijson.Field
	Threat      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingListResponseFindingRemediation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingListResponseFindingRemediationJSON) RawJSON() string {
	return r.raw
}

// Summary information about an integration.
type CasbPostureFindingListResponseIntegration struct {
	// When entity was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// When were the integration credentials last updated.
	LastHydrated time.Time `json:"last_hydrated" api:"required" format:"date-time"`
	// Name of the integration.
	Name string `json:"name" api:"required"`
	// The vendor-specific permissions associated with the integration.
	Permissions []string `json:"permissions" api:"required"`
	// Policy configuration for an integration.
	Policy CasbPostureFindingListResponseIntegrationPolicy `json:"policy" api:"required"`
	// Current status of the integration.
	Status string `json:"status" api:"required"`
	// Last entity was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Whether the integrations permissions can be updated.
	Upgradable bool `json:"upgradable" api:"required"`
	// Information about a vendor/service provider.
	Vendor CasbPostureFindingListResponseIntegrationVendor `json:"vendor" api:"required"`
	// Zero Trust products associated with this integration.
	ZtEnrollments []CasbPostureFindingListResponseIntegrationZtEnrollment `json:"zt_enrollments" api:"required"`
	// Integration ID.
	ID string `json:"id" format:"uuid"`
	// Health status of integration credentials.
	CredentialHealthStatus CasbPostureFindingListResponseIntegrationCredentialHealthStatus `json:"credential_health_status"`
	// The date and time when the integration credentials will expire.
	CredentialsExpiry time.Time `json:"credentials_expiry" api:"nullable" format:"date-time"`
	// Whether the given integration is paused by the user.
	IsPaused bool `json:"is_paused"`
	// UI State as to whether a potential permissions upgrade has been dismissed.
	UpgradeDismissed bool                                          `json:"upgrade_dismissed"`
	JSON             casbPostureFindingListResponseIntegrationJSON `json:"-"`
}

// casbPostureFindingListResponseIntegrationJSON contains the JSON metadata for the
// struct [CasbPostureFindingListResponseIntegration]
type casbPostureFindingListResponseIntegrationJSON struct {
	Created                apijson.Field
	LastHydrated           apijson.Field
	Name                   apijson.Field
	Permissions            apijson.Field
	Policy                 apijson.Field
	Status                 apijson.Field
	Updated                apijson.Field
	Upgradable             apijson.Field
	Vendor                 apijson.Field
	ZtEnrollments          apijson.Field
	ID                     apijson.Field
	CredentialHealthStatus apijson.Field
	CredentialsExpiry      apijson.Field
	IsPaused               apijson.Field
	UpgradeDismissed       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CasbPostureFindingListResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingListResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

// Policy configuration for an integration.
type CasbPostureFindingListResponseIntegrationPolicy struct {
	// Policy identifier.
	ID string `json:"id" format:"uuid"`
	// OAuth client ID for the policy.
	ClientID string `json:"client_id" api:"nullable"`
	// Compliance level for the policy.
	ComplianceLevel string `json:"compliance_level"`
	// Whether DLP is enabled for this policy.
	DLPEnabled bool `json:"dlp_enabled"`
	// Link to policy documentation.
	Link string `json:"link" api:"nullable" format:"uri"`
	// Policy name.
	Name string `json:"name"`
	// List of permissions included in the policy.
	Permissions []string                                            `json:"permissions"`
	JSON        casbPostureFindingListResponseIntegrationPolicyJSON `json:"-"`
}

// casbPostureFindingListResponseIntegrationPolicyJSON contains the JSON metadata
// for the struct [CasbPostureFindingListResponseIntegrationPolicy]
type casbPostureFindingListResponseIntegrationPolicyJSON struct {
	ID              apijson.Field
	ClientID        apijson.Field
	ComplianceLevel apijson.Field
	DLPEnabled      apijson.Field
	Link            apijson.Field
	Name            apijson.Field
	Permissions     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CasbPostureFindingListResponseIntegrationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingListResponseIntegrationPolicyJSON) RawJSON() string {
	return r.raw
}

// Information about a vendor/service provider.
type CasbPostureFindingListResponseIntegrationVendor struct {
	// The id of the vendor.
	ID string `json:"id" api:"required"`
	// Detailed information about what kinds of issues are detected for this vendor.
	Description string `json:"description" api:"required,nullable"`
	// The display name of the vendor.
	DisplayName string `json:"display_name" api:"required"`
	// Logo URL for the vendor.
	Logo string `json:"logo" api:"required" format:"uri"`
	// The name of the vendor.
	Name string `json:"name" api:"required"`
	// Static logo URL for the vendor.
	StaticLogo string `json:"static_logo" api:"required" format:"uri"`
	// The vendor's compatible Zero Trust products.
	ZtEnrollments []string `json:"zt_enrollments" api:"required"`
	// The policies related to the vendor.
	Policies []map[string]interface{}                            `json:"policies"`
	JSON     casbPostureFindingListResponseIntegrationVendorJSON `json:"-"`
}

// casbPostureFindingListResponseIntegrationVendorJSON contains the JSON metadata
// for the struct [CasbPostureFindingListResponseIntegrationVendor]
type casbPostureFindingListResponseIntegrationVendorJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Logo          apijson.Field
	Name          apijson.Field
	StaticLogo    apijson.Field
	ZtEnrollments apijson.Field
	Policies      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CasbPostureFindingListResponseIntegrationVendor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingListResponseIntegrationVendorJSON) RawJSON() string {
	return r.raw
}

// Information about a Zero Trust product integration.
type CasbPostureFindingListResponseIntegrationZtEnrollment struct {
	// The internal identifier of the Zero Trust Product.
	ID string `json:"id"`
	// Brief description of the Zero Trust Product.
	Description string `json:"description"`
	// The verbose name of the Zero Trust Product.
	DisplayName string `json:"display_name"`
	// Flag to enable/disable access to the listed integration from the corresponding
	// Cloudflare product.
	Enabled bool                                                      `json:"enabled"`
	JSON    casbPostureFindingListResponseIntegrationZtEnrollmentJSON `json:"-"`
}

// casbPostureFindingListResponseIntegrationZtEnrollmentJSON contains the JSON
// metadata for the struct [CasbPostureFindingListResponseIntegrationZtEnrollment]
type casbPostureFindingListResponseIntegrationZtEnrollmentJSON struct {
	ID          apijson.Field
	Description apijson.Field
	DisplayName apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingListResponseIntegrationZtEnrollment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingListResponseIntegrationZtEnrollmentJSON) RawJSON() string {
	return r.raw
}

// Health status of integration credentials.
type CasbPostureFindingListResponseIntegrationCredentialHealthStatus string

const (
	CasbPostureFindingListResponseIntegrationCredentialHealthStatusInitializing CasbPostureFindingListResponseIntegrationCredentialHealthStatus = "Initializing"
	CasbPostureFindingListResponseIntegrationCredentialHealthStatusHealthy      CasbPostureFindingListResponseIntegrationCredentialHealthStatus = "Healthy"
	CasbPostureFindingListResponseIntegrationCredentialHealthStatusUnhealthy    CasbPostureFindingListResponseIntegrationCredentialHealthStatus = "Unhealthy"
)

func (r CasbPostureFindingListResponseIntegrationCredentialHealthStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingListResponseIntegrationCredentialHealthStatusInitializing, CasbPostureFindingListResponseIntegrationCredentialHealthStatusHealthy, CasbPostureFindingListResponseIntegrationCredentialHealthStatusUnhealthy:
		return true
	}
	return false
}

// Override information for finding severity.
type CasbPostureFindingListResponseSeverityOverride struct {
	// User ID who created the override.
	CreatedBy string `json:"created_by" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingListResponseSeverityOverrideSeverity `json:"severity" api:"required"`
	JSON     casbPostureFindingListResponseSeverityOverrideJSON     `json:"-"`
}

// casbPostureFindingListResponseSeverityOverrideJSON contains the JSON metadata
// for the struct [CasbPostureFindingListResponseSeverityOverride]
type casbPostureFindingListResponseSeverityOverrideJSON struct {
	CreatedBy   apijson.Field
	Severity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingListResponseSeverityOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingListResponseSeverityOverrideJSON) RawJSON() string {
	return r.raw
}

// The severity level of a finding.
type CasbPostureFindingListResponseSeverityOverrideSeverity string

const (
	CasbPostureFindingListResponseSeverityOverrideSeverityCritical CasbPostureFindingListResponseSeverityOverrideSeverity = "Critical"
	CasbPostureFindingListResponseSeverityOverrideSeverityHigh     CasbPostureFindingListResponseSeverityOverrideSeverity = "High"
	CasbPostureFindingListResponseSeverityOverrideSeverityMedium   CasbPostureFindingListResponseSeverityOverrideSeverity = "Medium"
	CasbPostureFindingListResponseSeverityOverrideSeverityLow      CasbPostureFindingListResponseSeverityOverrideSeverity = "Low"
)

func (r CasbPostureFindingListResponseSeverityOverrideSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingListResponseSeverityOverrideSeverityCritical, CasbPostureFindingListResponseSeverityOverrideSeverityHigh, CasbPostureFindingListResponseSeverityOverrideSeverityMedium, CasbPostureFindingListResponseSeverityOverrideSeverityLow:
		return true
	}
	return false
}

// Information about an export job.
type CasbPostureFindingExportResponse struct {
	// Unique identifier for the export job.
	ID string `json:"id" api:"required" format:"uuid"`
	// Status of an export job.
	Status CasbPostureFindingExportResponseStatus `json:"status" api:"required"`
	// Type of export job.
	Type CasbPostureFindingExportResponseType `json:"type" api:"required"`
	// ID of the export-requesting user.
	UserID string `json:"user_id" api:"required"`
	// The URL by which the successfully created export can be downloaded by the end
	// users.
	DownloadURL string `json:"download_url" api:"nullable" format:"uri"`
	// Contains information on errors which may have occurred during export creation.
	Errors string `json:"errors" api:"nullable"`
	// The base name of the file that is/was generated by the export job.
	FileName string `json:"file_name" api:"nullable"`
	// The full path of the file that is stored within external storage (currently R2).
	FilePath string                               `json:"file_path" api:"nullable"`
	JSON     casbPostureFindingExportResponseJSON `json:"-"`
}

// casbPostureFindingExportResponseJSON contains the JSON metadata for the struct
// [CasbPostureFindingExportResponse]
type casbPostureFindingExportResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	UserID      apijson.Field
	DownloadURL apijson.Field
	Errors      apijson.Field
	FileName    apijson.Field
	FilePath    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingExportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingExportResponseJSON) RawJSON() string {
	return r.raw
}

// Status of an export job.
type CasbPostureFindingExportResponseStatus string

const (
	CasbPostureFindingExportResponseStatusPending     CasbPostureFindingExportResponseStatus = "Pending"
	CasbPostureFindingExportResponseStatusSuccess     CasbPostureFindingExportResponseStatus = "Success"
	CasbPostureFindingExportResponseStatusFailure     CasbPostureFindingExportResponseStatus = "Failure"
	CasbPostureFindingExportResponseStatusRescheduled CasbPostureFindingExportResponseStatus = "Rescheduled"
	CasbPostureFindingExportResponseStatusInProgress  CasbPostureFindingExportResponseStatus = "In-Progress"
)

func (r CasbPostureFindingExportResponseStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingExportResponseStatusPending, CasbPostureFindingExportResponseStatusSuccess, CasbPostureFindingExportResponseStatusFailure, CasbPostureFindingExportResponseStatusRescheduled, CasbPostureFindingExportResponseStatusInProgress:
		return true
	}
	return false
}

// Type of export job.
type CasbPostureFindingExportResponseType string

const (
	CasbPostureFindingExportResponseTypeFinding         CasbPostureFindingExportResponseType = "finding"
	CasbPostureFindingExportResponseTypeFindingInstance CasbPostureFindingExportResponseType = "findingInstance"
	CasbPostureFindingExportResponseTypeContent         CasbPostureFindingExportResponseType = "content"
	CasbPostureFindingExportResponseTypeRemediationJob  CasbPostureFindingExportResponseType = "remediationJob"
)

func (r CasbPostureFindingExportResponseType) IsKnown() bool {
	switch r {
	case CasbPostureFindingExportResponseTypeFinding, CasbPostureFindingExportResponseTypeFindingInstance, CasbPostureFindingExportResponseTypeContent, CasbPostureFindingExportResponseTypeRemediationJob:
		return true
	}
	return false
}

// Aggregated finding information with counts and metadata. This is optimized for
// list API queries and represents a finding along with its instance statistics.
type CasbPostureFindingGetResponse struct {
	// Base64 encoded identifier of the security finding.
	ID string `json:"id" api:"required" format:"byte"`
	// Number of active problematic instances identified in the security finding.
	ActiveCount int64 `json:"active_count" api:"required"`
	// Number of archived instances identified in the security finding.
	ArchivedCount int64 `json:"archived_count" api:"required"`
	// Basic finding type information.
	Finding CasbPostureFindingGetResponseFinding `json:"finding" api:"required"`
	// Determines if finding is currently ignored.
	Ignored bool `json:"ignored" api:"required"`
	// Number of total (Active or archived) problematic instances identified in the
	// security finding.
	InstanceCount int64 `json:"instance_count" api:"required"`
	// Summary information about an integration.
	Integration CasbPostureFindingGetResponseIntegration `json:"integration" api:"required"`
	// Timestamp of the latest affliction date of an active finding.
	LatestAfflictionDate time.Time `json:"latest_affliction_date" api:"required" format:"date-time"`
	// Override information for finding severity.
	SeverityOverride CasbPostureFindingGetResponseSeverityOverride `json:"severity_override"`
	JSON             casbPostureFindingGetResponseJSON             `json:"-"`
}

// casbPostureFindingGetResponseJSON contains the JSON metadata for the struct
// [CasbPostureFindingGetResponse]
type casbPostureFindingGetResponseJSON struct {
	ID                   apijson.Field
	ActiveCount          apijson.Field
	ArchivedCount        apijson.Field
	Finding              apijson.Field
	Ignored              apijson.Field
	InstanceCount        apijson.Field
	Integration          apijson.Field
	LatestAfflictionDate apijson.Field
	SeverityOverride     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseJSON) RawJSON() string {
	return r.raw
}

// Basic finding type information.
type CasbPostureFindingGetResponseFinding struct {
	// The unique identifier of the finding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Category information for a finding.
	Category CasbPostureFindingGetResponseFindingCategory `json:"category" api:"required"`
	// The name of the finding.
	Name string `json:"name" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingGetResponseFindingSeverity `json:"severity" api:"required"`
	// The SaaS/Cloud vendor of the platform with which the finding is associated.
	Vendor string `json:"vendor" api:"required"`
	// Detailed description of the finding.
	Description string `json:"description" api:"nullable"`
	// Remediation guide information for a finding.
	Remediation CasbPostureFindingGetResponseFindingRemediation `json:"remediation" api:"nullable"`
	JSON        casbPostureFindingGetResponseFindingJSON        `json:"-"`
}

// casbPostureFindingGetResponseFindingJSON contains the JSON metadata for the
// struct [CasbPostureFindingGetResponseFinding]
type casbPostureFindingGetResponseFindingJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	Name        apijson.Field
	Severity    apijson.Field
	Vendor      apijson.Field
	Description apijson.Field
	Remediation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponseFinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseFindingJSON) RawJSON() string {
	return r.raw
}

// Category information for a finding.
type CasbPostureFindingGetResponseFindingCategory struct {
	// The type of the observation.
	Observation CasbPostureFindingGetResponseFindingCategoryObservation `json:"observation" api:"required"`
	// The product category.
	Product CasbPostureFindingGetResponseFindingCategoryProduct `json:"product" api:"required"`
	// The type of the finding category.
	Type CasbPostureFindingGetResponseFindingCategoryType `json:"type" api:"required"`
	JSON casbPostureFindingGetResponseFindingCategoryJSON `json:"-"`
}

// casbPostureFindingGetResponseFindingCategoryJSON contains the JSON metadata for
// the struct [CasbPostureFindingGetResponseFindingCategory]
type casbPostureFindingGetResponseFindingCategoryJSON struct {
	Observation apijson.Field
	Product     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponseFindingCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseFindingCategoryJSON) RawJSON() string {
	return r.raw
}

// The type of the observation.
type CasbPostureFindingGetResponseFindingCategoryObservation string

const (
	CasbPostureFindingGetResponseFindingCategoryObservationIssue    CasbPostureFindingGetResponseFindingCategoryObservation = "Issue"
	CasbPostureFindingGetResponseFindingCategoryObservationInsight  CasbPostureFindingGetResponseFindingCategoryObservation = "Insight"
	CasbPostureFindingGetResponseFindingCategoryObservationActivity CasbPostureFindingGetResponseFindingCategoryObservation = "Activity"
)

func (r CasbPostureFindingGetResponseFindingCategoryObservation) IsKnown() bool {
	switch r {
	case CasbPostureFindingGetResponseFindingCategoryObservationIssue, CasbPostureFindingGetResponseFindingCategoryObservationInsight, CasbPostureFindingGetResponseFindingCategoryObservationActivity:
		return true
	}
	return false
}

// The product category.
type CasbPostureFindingGetResponseFindingCategoryProduct string

const (
	CasbPostureFindingGetResponseFindingCategoryProductSaaS  CasbPostureFindingGetResponseFindingCategoryProduct = "SaaS"
	CasbPostureFindingGetResponseFindingCategoryProductCloud CasbPostureFindingGetResponseFindingCategoryProduct = "Cloud"
)

func (r CasbPostureFindingGetResponseFindingCategoryProduct) IsKnown() bool {
	switch r {
	case CasbPostureFindingGetResponseFindingCategoryProductSaaS, CasbPostureFindingGetResponseFindingCategoryProductCloud:
		return true
	}
	return false
}

// The type of the finding category.
type CasbPostureFindingGetResponseFindingCategoryType string

const (
	CasbPostureFindingGetResponseFindingCategoryTypeContent CasbPostureFindingGetResponseFindingCategoryType = "Content"
	CasbPostureFindingGetResponseFindingCategoryTypePosture CasbPostureFindingGetResponseFindingCategoryType = "Posture"
)

func (r CasbPostureFindingGetResponseFindingCategoryType) IsKnown() bool {
	switch r {
	case CasbPostureFindingGetResponseFindingCategoryTypeContent, CasbPostureFindingGetResponseFindingCategoryTypePosture:
		return true
	}
	return false
}

// The severity level of a finding.
type CasbPostureFindingGetResponseFindingSeverity string

const (
	CasbPostureFindingGetResponseFindingSeverityCritical CasbPostureFindingGetResponseFindingSeverity = "Critical"
	CasbPostureFindingGetResponseFindingSeverityHigh     CasbPostureFindingGetResponseFindingSeverity = "High"
	CasbPostureFindingGetResponseFindingSeverityMedium   CasbPostureFindingGetResponseFindingSeverity = "Medium"
	CasbPostureFindingGetResponseFindingSeverityLow      CasbPostureFindingGetResponseFindingSeverity = "Low"
)

func (r CasbPostureFindingGetResponseFindingSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingGetResponseFindingSeverityCritical, CasbPostureFindingGetResponseFindingSeverityHigh, CasbPostureFindingGetResponseFindingSeverityMedium, CasbPostureFindingGetResponseFindingSeverityLow:
		return true
	}
	return false
}

// Remediation guide information for a finding.
type CasbPostureFindingGetResponseFindingRemediation struct {
	// Remediation Id.
	ID string `json:"id" api:"required" format:"uuid"`
	// Relevant Compliance Frameworks.
	Frameworks []string `json:"frameworks" api:"required"`
	// Remediation guide text.
	Guide string `json:"guide" api:"required"`
	// Description of the potential impact.
	Impact string `json:"impact" api:"required"`
	// I18N Locale.
	Locale string `json:"locale" api:"required"`
	// Description of the threat.
	Threat string                                              `json:"threat" api:"required"`
	JSON   casbPostureFindingGetResponseFindingRemediationJSON `json:"-"`
}

// casbPostureFindingGetResponseFindingRemediationJSON contains the JSON metadata
// for the struct [CasbPostureFindingGetResponseFindingRemediation]
type casbPostureFindingGetResponseFindingRemediationJSON struct {
	ID          apijson.Field
	Frameworks  apijson.Field
	Guide       apijson.Field
	Impact      apijson.Field
	Locale      apijson.Field
	Threat      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponseFindingRemediation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseFindingRemediationJSON) RawJSON() string {
	return r.raw
}

// Summary information about an integration.
type CasbPostureFindingGetResponseIntegration struct {
	// When entity was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// When were the integration credentials last updated.
	LastHydrated time.Time `json:"last_hydrated" api:"required" format:"date-time"`
	// Name of the integration.
	Name string `json:"name" api:"required"`
	// The vendor-specific permissions associated with the integration.
	Permissions []string `json:"permissions" api:"required"`
	// Policy configuration for an integration.
	Policy CasbPostureFindingGetResponseIntegrationPolicy `json:"policy" api:"required"`
	// Current status of the integration.
	Status string `json:"status" api:"required"`
	// Last entity was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Whether the integrations permissions can be updated.
	Upgradable bool `json:"upgradable" api:"required"`
	// Information about a vendor/service provider.
	Vendor CasbPostureFindingGetResponseIntegrationVendor `json:"vendor" api:"required"`
	// Zero Trust products associated with this integration.
	ZtEnrollments []CasbPostureFindingGetResponseIntegrationZtEnrollment `json:"zt_enrollments" api:"required"`
	// Integration ID.
	ID string `json:"id" format:"uuid"`
	// Health status of integration credentials.
	CredentialHealthStatus CasbPostureFindingGetResponseIntegrationCredentialHealthStatus `json:"credential_health_status"`
	// The date and time when the integration credentials will expire.
	CredentialsExpiry time.Time `json:"credentials_expiry" api:"nullable" format:"date-time"`
	// Whether the given integration is paused by the user.
	IsPaused bool `json:"is_paused"`
	// UI State as to whether a potential permissions upgrade has been dismissed.
	UpgradeDismissed bool                                         `json:"upgrade_dismissed"`
	JSON             casbPostureFindingGetResponseIntegrationJSON `json:"-"`
}

// casbPostureFindingGetResponseIntegrationJSON contains the JSON metadata for the
// struct [CasbPostureFindingGetResponseIntegration]
type casbPostureFindingGetResponseIntegrationJSON struct {
	Created                apijson.Field
	LastHydrated           apijson.Field
	Name                   apijson.Field
	Permissions            apijson.Field
	Policy                 apijson.Field
	Status                 apijson.Field
	Updated                apijson.Field
	Upgradable             apijson.Field
	Vendor                 apijson.Field
	ZtEnrollments          apijson.Field
	ID                     apijson.Field
	CredentialHealthStatus apijson.Field
	CredentialsExpiry      apijson.Field
	IsPaused               apijson.Field
	UpgradeDismissed       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

// Policy configuration for an integration.
type CasbPostureFindingGetResponseIntegrationPolicy struct {
	// Policy identifier.
	ID string `json:"id" format:"uuid"`
	// OAuth client ID for the policy.
	ClientID string `json:"client_id" api:"nullable"`
	// Compliance level for the policy.
	ComplianceLevel string `json:"compliance_level"`
	// Whether DLP is enabled for this policy.
	DLPEnabled bool `json:"dlp_enabled"`
	// Link to policy documentation.
	Link string `json:"link" api:"nullable" format:"uri"`
	// Policy name.
	Name string `json:"name"`
	// List of permissions included in the policy.
	Permissions []string                                           `json:"permissions"`
	JSON        casbPostureFindingGetResponseIntegrationPolicyJSON `json:"-"`
}

// casbPostureFindingGetResponseIntegrationPolicyJSON contains the JSON metadata
// for the struct [CasbPostureFindingGetResponseIntegrationPolicy]
type casbPostureFindingGetResponseIntegrationPolicyJSON struct {
	ID              apijson.Field
	ClientID        apijson.Field
	ComplianceLevel apijson.Field
	DLPEnabled      apijson.Field
	Link            apijson.Field
	Name            apijson.Field
	Permissions     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponseIntegrationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseIntegrationPolicyJSON) RawJSON() string {
	return r.raw
}

// Information about a vendor/service provider.
type CasbPostureFindingGetResponseIntegrationVendor struct {
	// The id of the vendor.
	ID string `json:"id" api:"required"`
	// Detailed information about what kinds of issues are detected for this vendor.
	Description string `json:"description" api:"required,nullable"`
	// The display name of the vendor.
	DisplayName string `json:"display_name" api:"required"`
	// Logo URL for the vendor.
	Logo string `json:"logo" api:"required" format:"uri"`
	// The name of the vendor.
	Name string `json:"name" api:"required"`
	// Static logo URL for the vendor.
	StaticLogo string `json:"static_logo" api:"required" format:"uri"`
	// The vendor's compatible Zero Trust products.
	ZtEnrollments []string `json:"zt_enrollments" api:"required"`
	// The policies related to the vendor.
	Policies []map[string]interface{}                           `json:"policies"`
	JSON     casbPostureFindingGetResponseIntegrationVendorJSON `json:"-"`
}

// casbPostureFindingGetResponseIntegrationVendorJSON contains the JSON metadata
// for the struct [CasbPostureFindingGetResponseIntegrationVendor]
type casbPostureFindingGetResponseIntegrationVendorJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Logo          apijson.Field
	Name          apijson.Field
	StaticLogo    apijson.Field
	ZtEnrollments apijson.Field
	Policies      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponseIntegrationVendor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseIntegrationVendorJSON) RawJSON() string {
	return r.raw
}

// Information about a Zero Trust product integration.
type CasbPostureFindingGetResponseIntegrationZtEnrollment struct {
	// The internal identifier of the Zero Trust Product.
	ID string `json:"id"`
	// Brief description of the Zero Trust Product.
	Description string `json:"description"`
	// The verbose name of the Zero Trust Product.
	DisplayName string `json:"display_name"`
	// Flag to enable/disable access to the listed integration from the corresponding
	// Cloudflare product.
	Enabled bool                                                     `json:"enabled"`
	JSON    casbPostureFindingGetResponseIntegrationZtEnrollmentJSON `json:"-"`
}

// casbPostureFindingGetResponseIntegrationZtEnrollmentJSON contains the JSON
// metadata for the struct [CasbPostureFindingGetResponseIntegrationZtEnrollment]
type casbPostureFindingGetResponseIntegrationZtEnrollmentJSON struct {
	ID          apijson.Field
	Description apijson.Field
	DisplayName apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponseIntegrationZtEnrollment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseIntegrationZtEnrollmentJSON) RawJSON() string {
	return r.raw
}

// Health status of integration credentials.
type CasbPostureFindingGetResponseIntegrationCredentialHealthStatus string

const (
	CasbPostureFindingGetResponseIntegrationCredentialHealthStatusInitializing CasbPostureFindingGetResponseIntegrationCredentialHealthStatus = "Initializing"
	CasbPostureFindingGetResponseIntegrationCredentialHealthStatusHealthy      CasbPostureFindingGetResponseIntegrationCredentialHealthStatus = "Healthy"
	CasbPostureFindingGetResponseIntegrationCredentialHealthStatusUnhealthy    CasbPostureFindingGetResponseIntegrationCredentialHealthStatus = "Unhealthy"
)

func (r CasbPostureFindingGetResponseIntegrationCredentialHealthStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingGetResponseIntegrationCredentialHealthStatusInitializing, CasbPostureFindingGetResponseIntegrationCredentialHealthStatusHealthy, CasbPostureFindingGetResponseIntegrationCredentialHealthStatusUnhealthy:
		return true
	}
	return false
}

// Override information for finding severity.
type CasbPostureFindingGetResponseSeverityOverride struct {
	// User ID who created the override.
	CreatedBy string `json:"created_by" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingGetResponseSeverityOverrideSeverity `json:"severity" api:"required"`
	JSON     casbPostureFindingGetResponseSeverityOverrideJSON     `json:"-"`
}

// casbPostureFindingGetResponseSeverityOverrideJSON contains the JSON metadata for
// the struct [CasbPostureFindingGetResponseSeverityOverride]
type casbPostureFindingGetResponseSeverityOverrideJSON struct {
	CreatedBy   apijson.Field
	Severity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponseSeverityOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseSeverityOverrideJSON) RawJSON() string {
	return r.raw
}

// The severity level of a finding.
type CasbPostureFindingGetResponseSeverityOverrideSeverity string

const (
	CasbPostureFindingGetResponseSeverityOverrideSeverityCritical CasbPostureFindingGetResponseSeverityOverrideSeverity = "Critical"
	CasbPostureFindingGetResponseSeverityOverrideSeverityHigh     CasbPostureFindingGetResponseSeverityOverrideSeverity = "High"
	CasbPostureFindingGetResponseSeverityOverrideSeverityMedium   CasbPostureFindingGetResponseSeverityOverrideSeverity = "Medium"
	CasbPostureFindingGetResponseSeverityOverrideSeverityLow      CasbPostureFindingGetResponseSeverityOverrideSeverity = "Low"
)

func (r CasbPostureFindingGetResponseSeverityOverrideSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingGetResponseSeverityOverrideSeverityCritical, CasbPostureFindingGetResponseSeverityOverrideSeverityHigh, CasbPostureFindingGetResponseSeverityOverrideSeverityMedium, CasbPostureFindingGetResponseSeverityOverrideSeverityLow:
		return true
	}
	return false
}

// Aggregated finding information with counts and metadata. This is optimized for
// list API queries and represents a finding along with its instance statistics.
type CasbPostureFindingIgnoreResponse struct {
	// Base64 encoded identifier of the security finding.
	ID string `json:"id" api:"required" format:"byte"`
	// Number of active problematic instances identified in the security finding.
	ActiveCount int64 `json:"active_count" api:"required"`
	// Number of archived instances identified in the security finding.
	ArchivedCount int64 `json:"archived_count" api:"required"`
	// Basic finding type information.
	Finding CasbPostureFindingIgnoreResponseFinding `json:"finding" api:"required"`
	// Determines if finding is currently ignored.
	Ignored bool `json:"ignored" api:"required"`
	// Number of total (Active or archived) problematic instances identified in the
	// security finding.
	InstanceCount int64 `json:"instance_count" api:"required"`
	// Summary information about an integration.
	Integration CasbPostureFindingIgnoreResponseIntegration `json:"integration" api:"required"`
	// Timestamp of the latest affliction date of an active finding.
	LatestAfflictionDate time.Time `json:"latest_affliction_date" api:"required" format:"date-time"`
	// Override information for finding severity.
	SeverityOverride CasbPostureFindingIgnoreResponseSeverityOverride `json:"severity_override"`
	JSON             casbPostureFindingIgnoreResponseJSON             `json:"-"`
}

// casbPostureFindingIgnoreResponseJSON contains the JSON metadata for the struct
// [CasbPostureFindingIgnoreResponse]
type casbPostureFindingIgnoreResponseJSON struct {
	ID                   apijson.Field
	ActiveCount          apijson.Field
	ArchivedCount        apijson.Field
	Finding              apijson.Field
	Ignored              apijson.Field
	InstanceCount        apijson.Field
	Integration          apijson.Field
	LatestAfflictionDate apijson.Field
	SeverityOverride     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseJSON) RawJSON() string {
	return r.raw
}

// Basic finding type information.
type CasbPostureFindingIgnoreResponseFinding struct {
	// The unique identifier of the finding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Category information for a finding.
	Category CasbPostureFindingIgnoreResponseFindingCategory `json:"category" api:"required"`
	// The name of the finding.
	Name string `json:"name" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingIgnoreResponseFindingSeverity `json:"severity" api:"required"`
	// The SaaS/Cloud vendor of the platform with which the finding is associated.
	Vendor string `json:"vendor" api:"required"`
	// Detailed description of the finding.
	Description string `json:"description" api:"nullable"`
	// Remediation guide information for a finding.
	Remediation CasbPostureFindingIgnoreResponseFindingRemediation `json:"remediation" api:"nullable"`
	JSON        casbPostureFindingIgnoreResponseFindingJSON        `json:"-"`
}

// casbPostureFindingIgnoreResponseFindingJSON contains the JSON metadata for the
// struct [CasbPostureFindingIgnoreResponseFinding]
type casbPostureFindingIgnoreResponseFindingJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	Name        apijson.Field
	Severity    apijson.Field
	Vendor      apijson.Field
	Description apijson.Field
	Remediation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponseFinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseFindingJSON) RawJSON() string {
	return r.raw
}

// Category information for a finding.
type CasbPostureFindingIgnoreResponseFindingCategory struct {
	// The type of the observation.
	Observation CasbPostureFindingIgnoreResponseFindingCategoryObservation `json:"observation" api:"required"`
	// The product category.
	Product CasbPostureFindingIgnoreResponseFindingCategoryProduct `json:"product" api:"required"`
	// The type of the finding category.
	Type CasbPostureFindingIgnoreResponseFindingCategoryType `json:"type" api:"required"`
	JSON casbPostureFindingIgnoreResponseFindingCategoryJSON `json:"-"`
}

// casbPostureFindingIgnoreResponseFindingCategoryJSON contains the JSON metadata
// for the struct [CasbPostureFindingIgnoreResponseFindingCategory]
type casbPostureFindingIgnoreResponseFindingCategoryJSON struct {
	Observation apijson.Field
	Product     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponseFindingCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseFindingCategoryJSON) RawJSON() string {
	return r.raw
}

// The type of the observation.
type CasbPostureFindingIgnoreResponseFindingCategoryObservation string

const (
	CasbPostureFindingIgnoreResponseFindingCategoryObservationIssue    CasbPostureFindingIgnoreResponseFindingCategoryObservation = "Issue"
	CasbPostureFindingIgnoreResponseFindingCategoryObservationInsight  CasbPostureFindingIgnoreResponseFindingCategoryObservation = "Insight"
	CasbPostureFindingIgnoreResponseFindingCategoryObservationActivity CasbPostureFindingIgnoreResponseFindingCategoryObservation = "Activity"
)

func (r CasbPostureFindingIgnoreResponseFindingCategoryObservation) IsKnown() bool {
	switch r {
	case CasbPostureFindingIgnoreResponseFindingCategoryObservationIssue, CasbPostureFindingIgnoreResponseFindingCategoryObservationInsight, CasbPostureFindingIgnoreResponseFindingCategoryObservationActivity:
		return true
	}
	return false
}

// The product category.
type CasbPostureFindingIgnoreResponseFindingCategoryProduct string

const (
	CasbPostureFindingIgnoreResponseFindingCategoryProductSaaS  CasbPostureFindingIgnoreResponseFindingCategoryProduct = "SaaS"
	CasbPostureFindingIgnoreResponseFindingCategoryProductCloud CasbPostureFindingIgnoreResponseFindingCategoryProduct = "Cloud"
)

func (r CasbPostureFindingIgnoreResponseFindingCategoryProduct) IsKnown() bool {
	switch r {
	case CasbPostureFindingIgnoreResponseFindingCategoryProductSaaS, CasbPostureFindingIgnoreResponseFindingCategoryProductCloud:
		return true
	}
	return false
}

// The type of the finding category.
type CasbPostureFindingIgnoreResponseFindingCategoryType string

const (
	CasbPostureFindingIgnoreResponseFindingCategoryTypeContent CasbPostureFindingIgnoreResponseFindingCategoryType = "Content"
	CasbPostureFindingIgnoreResponseFindingCategoryTypePosture CasbPostureFindingIgnoreResponseFindingCategoryType = "Posture"
)

func (r CasbPostureFindingIgnoreResponseFindingCategoryType) IsKnown() bool {
	switch r {
	case CasbPostureFindingIgnoreResponseFindingCategoryTypeContent, CasbPostureFindingIgnoreResponseFindingCategoryTypePosture:
		return true
	}
	return false
}

// The severity level of a finding.
type CasbPostureFindingIgnoreResponseFindingSeverity string

const (
	CasbPostureFindingIgnoreResponseFindingSeverityCritical CasbPostureFindingIgnoreResponseFindingSeverity = "Critical"
	CasbPostureFindingIgnoreResponseFindingSeverityHigh     CasbPostureFindingIgnoreResponseFindingSeverity = "High"
	CasbPostureFindingIgnoreResponseFindingSeverityMedium   CasbPostureFindingIgnoreResponseFindingSeverity = "Medium"
	CasbPostureFindingIgnoreResponseFindingSeverityLow      CasbPostureFindingIgnoreResponseFindingSeverity = "Low"
)

func (r CasbPostureFindingIgnoreResponseFindingSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingIgnoreResponseFindingSeverityCritical, CasbPostureFindingIgnoreResponseFindingSeverityHigh, CasbPostureFindingIgnoreResponseFindingSeverityMedium, CasbPostureFindingIgnoreResponseFindingSeverityLow:
		return true
	}
	return false
}

// Remediation guide information for a finding.
type CasbPostureFindingIgnoreResponseFindingRemediation struct {
	// Remediation Id.
	ID string `json:"id" api:"required" format:"uuid"`
	// Relevant Compliance Frameworks.
	Frameworks []string `json:"frameworks" api:"required"`
	// Remediation guide text.
	Guide string `json:"guide" api:"required"`
	// Description of the potential impact.
	Impact string `json:"impact" api:"required"`
	// I18N Locale.
	Locale string `json:"locale" api:"required"`
	// Description of the threat.
	Threat string                                                 `json:"threat" api:"required"`
	JSON   casbPostureFindingIgnoreResponseFindingRemediationJSON `json:"-"`
}

// casbPostureFindingIgnoreResponseFindingRemediationJSON contains the JSON
// metadata for the struct [CasbPostureFindingIgnoreResponseFindingRemediation]
type casbPostureFindingIgnoreResponseFindingRemediationJSON struct {
	ID          apijson.Field
	Frameworks  apijson.Field
	Guide       apijson.Field
	Impact      apijson.Field
	Locale      apijson.Field
	Threat      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponseFindingRemediation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseFindingRemediationJSON) RawJSON() string {
	return r.raw
}

// Summary information about an integration.
type CasbPostureFindingIgnoreResponseIntegration struct {
	// When entity was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// When were the integration credentials last updated.
	LastHydrated time.Time `json:"last_hydrated" api:"required" format:"date-time"`
	// Name of the integration.
	Name string `json:"name" api:"required"`
	// The vendor-specific permissions associated with the integration.
	Permissions []string `json:"permissions" api:"required"`
	// Policy configuration for an integration.
	Policy CasbPostureFindingIgnoreResponseIntegrationPolicy `json:"policy" api:"required"`
	// Current status of the integration.
	Status string `json:"status" api:"required"`
	// Last entity was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Whether the integrations permissions can be updated.
	Upgradable bool `json:"upgradable" api:"required"`
	// Information about a vendor/service provider.
	Vendor CasbPostureFindingIgnoreResponseIntegrationVendor `json:"vendor" api:"required"`
	// Zero Trust products associated with this integration.
	ZtEnrollments []CasbPostureFindingIgnoreResponseIntegrationZtEnrollment `json:"zt_enrollments" api:"required"`
	// Integration ID.
	ID string `json:"id" format:"uuid"`
	// Health status of integration credentials.
	CredentialHealthStatus CasbPostureFindingIgnoreResponseIntegrationCredentialHealthStatus `json:"credential_health_status"`
	// The date and time when the integration credentials will expire.
	CredentialsExpiry time.Time `json:"credentials_expiry" api:"nullable" format:"date-time"`
	// Whether the given integration is paused by the user.
	IsPaused bool `json:"is_paused"`
	// UI State as to whether a potential permissions upgrade has been dismissed.
	UpgradeDismissed bool                                            `json:"upgrade_dismissed"`
	JSON             casbPostureFindingIgnoreResponseIntegrationJSON `json:"-"`
}

// casbPostureFindingIgnoreResponseIntegrationJSON contains the JSON metadata for
// the struct [CasbPostureFindingIgnoreResponseIntegration]
type casbPostureFindingIgnoreResponseIntegrationJSON struct {
	Created                apijson.Field
	LastHydrated           apijson.Field
	Name                   apijson.Field
	Permissions            apijson.Field
	Policy                 apijson.Field
	Status                 apijson.Field
	Updated                apijson.Field
	Upgradable             apijson.Field
	Vendor                 apijson.Field
	ZtEnrollments          apijson.Field
	ID                     apijson.Field
	CredentialHealthStatus apijson.Field
	CredentialsExpiry      apijson.Field
	IsPaused               apijson.Field
	UpgradeDismissed       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

// Policy configuration for an integration.
type CasbPostureFindingIgnoreResponseIntegrationPolicy struct {
	// Policy identifier.
	ID string `json:"id" format:"uuid"`
	// OAuth client ID for the policy.
	ClientID string `json:"client_id" api:"nullable"`
	// Compliance level for the policy.
	ComplianceLevel string `json:"compliance_level"`
	// Whether DLP is enabled for this policy.
	DLPEnabled bool `json:"dlp_enabled"`
	// Link to policy documentation.
	Link string `json:"link" api:"nullable" format:"uri"`
	// Policy name.
	Name string `json:"name"`
	// List of permissions included in the policy.
	Permissions []string                                              `json:"permissions"`
	JSON        casbPostureFindingIgnoreResponseIntegrationPolicyJSON `json:"-"`
}

// casbPostureFindingIgnoreResponseIntegrationPolicyJSON contains the JSON metadata
// for the struct [CasbPostureFindingIgnoreResponseIntegrationPolicy]
type casbPostureFindingIgnoreResponseIntegrationPolicyJSON struct {
	ID              apijson.Field
	ClientID        apijson.Field
	ComplianceLevel apijson.Field
	DLPEnabled      apijson.Field
	Link            apijson.Field
	Name            apijson.Field
	Permissions     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponseIntegrationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseIntegrationPolicyJSON) RawJSON() string {
	return r.raw
}

// Information about a vendor/service provider.
type CasbPostureFindingIgnoreResponseIntegrationVendor struct {
	// The id of the vendor.
	ID string `json:"id" api:"required"`
	// Detailed information about what kinds of issues are detected for this vendor.
	Description string `json:"description" api:"required,nullable"`
	// The display name of the vendor.
	DisplayName string `json:"display_name" api:"required"`
	// Logo URL for the vendor.
	Logo string `json:"logo" api:"required" format:"uri"`
	// The name of the vendor.
	Name string `json:"name" api:"required"`
	// Static logo URL for the vendor.
	StaticLogo string `json:"static_logo" api:"required" format:"uri"`
	// The vendor's compatible Zero Trust products.
	ZtEnrollments []string `json:"zt_enrollments" api:"required"`
	// The policies related to the vendor.
	Policies []map[string]interface{}                              `json:"policies"`
	JSON     casbPostureFindingIgnoreResponseIntegrationVendorJSON `json:"-"`
}

// casbPostureFindingIgnoreResponseIntegrationVendorJSON contains the JSON metadata
// for the struct [CasbPostureFindingIgnoreResponseIntegrationVendor]
type casbPostureFindingIgnoreResponseIntegrationVendorJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Logo          apijson.Field
	Name          apijson.Field
	StaticLogo    apijson.Field
	ZtEnrollments apijson.Field
	Policies      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponseIntegrationVendor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseIntegrationVendorJSON) RawJSON() string {
	return r.raw
}

// Information about a Zero Trust product integration.
type CasbPostureFindingIgnoreResponseIntegrationZtEnrollment struct {
	// The internal identifier of the Zero Trust Product.
	ID string `json:"id"`
	// Brief description of the Zero Trust Product.
	Description string `json:"description"`
	// The verbose name of the Zero Trust Product.
	DisplayName string `json:"display_name"`
	// Flag to enable/disable access to the listed integration from the corresponding
	// Cloudflare product.
	Enabled bool                                                        `json:"enabled"`
	JSON    casbPostureFindingIgnoreResponseIntegrationZtEnrollmentJSON `json:"-"`
}

// casbPostureFindingIgnoreResponseIntegrationZtEnrollmentJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingIgnoreResponseIntegrationZtEnrollment]
type casbPostureFindingIgnoreResponseIntegrationZtEnrollmentJSON struct {
	ID          apijson.Field
	Description apijson.Field
	DisplayName apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponseIntegrationZtEnrollment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseIntegrationZtEnrollmentJSON) RawJSON() string {
	return r.raw
}

// Health status of integration credentials.
type CasbPostureFindingIgnoreResponseIntegrationCredentialHealthStatus string

const (
	CasbPostureFindingIgnoreResponseIntegrationCredentialHealthStatusInitializing CasbPostureFindingIgnoreResponseIntegrationCredentialHealthStatus = "Initializing"
	CasbPostureFindingIgnoreResponseIntegrationCredentialHealthStatusHealthy      CasbPostureFindingIgnoreResponseIntegrationCredentialHealthStatus = "Healthy"
	CasbPostureFindingIgnoreResponseIntegrationCredentialHealthStatusUnhealthy    CasbPostureFindingIgnoreResponseIntegrationCredentialHealthStatus = "Unhealthy"
)

func (r CasbPostureFindingIgnoreResponseIntegrationCredentialHealthStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingIgnoreResponseIntegrationCredentialHealthStatusInitializing, CasbPostureFindingIgnoreResponseIntegrationCredentialHealthStatusHealthy, CasbPostureFindingIgnoreResponseIntegrationCredentialHealthStatusUnhealthy:
		return true
	}
	return false
}

// Override information for finding severity.
type CasbPostureFindingIgnoreResponseSeverityOverride struct {
	// User ID who created the override.
	CreatedBy string `json:"created_by" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingIgnoreResponseSeverityOverrideSeverity `json:"severity" api:"required"`
	JSON     casbPostureFindingIgnoreResponseSeverityOverrideJSON     `json:"-"`
}

// casbPostureFindingIgnoreResponseSeverityOverrideJSON contains the JSON metadata
// for the struct [CasbPostureFindingIgnoreResponseSeverityOverride]
type casbPostureFindingIgnoreResponseSeverityOverrideJSON struct {
	CreatedBy   apijson.Field
	Severity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponseSeverityOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseSeverityOverrideJSON) RawJSON() string {
	return r.raw
}

// The severity level of a finding.
type CasbPostureFindingIgnoreResponseSeverityOverrideSeverity string

const (
	CasbPostureFindingIgnoreResponseSeverityOverrideSeverityCritical CasbPostureFindingIgnoreResponseSeverityOverrideSeverity = "Critical"
	CasbPostureFindingIgnoreResponseSeverityOverrideSeverityHigh     CasbPostureFindingIgnoreResponseSeverityOverrideSeverity = "High"
	CasbPostureFindingIgnoreResponseSeverityOverrideSeverityMedium   CasbPostureFindingIgnoreResponseSeverityOverrideSeverity = "Medium"
	CasbPostureFindingIgnoreResponseSeverityOverrideSeverityLow      CasbPostureFindingIgnoreResponseSeverityOverrideSeverity = "Low"
)

func (r CasbPostureFindingIgnoreResponseSeverityOverrideSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingIgnoreResponseSeverityOverrideSeverityCritical, CasbPostureFindingIgnoreResponseSeverityOverrideSeverityHigh, CasbPostureFindingIgnoreResponseSeverityOverrideSeverityMedium, CasbPostureFindingIgnoreResponseSeverityOverrideSeverityLow:
		return true
	}
	return false
}

// Aggregated finding information with counts and metadata. This is optimized for
// list API queries and represents a finding along with its instance statistics.
type CasbPostureFindingResetSeverityResponse struct {
	// Base64 encoded identifier of the security finding.
	ID string `json:"id" api:"required" format:"byte"`
	// Number of active problematic instances identified in the security finding.
	ActiveCount int64 `json:"active_count" api:"required"`
	// Number of archived instances identified in the security finding.
	ArchivedCount int64 `json:"archived_count" api:"required"`
	// Basic finding type information.
	Finding CasbPostureFindingResetSeverityResponseFinding `json:"finding" api:"required"`
	// Determines if finding is currently ignored.
	Ignored bool `json:"ignored" api:"required"`
	// Number of total (Active or archived) problematic instances identified in the
	// security finding.
	InstanceCount int64 `json:"instance_count" api:"required"`
	// Summary information about an integration.
	Integration CasbPostureFindingResetSeverityResponseIntegration `json:"integration" api:"required"`
	// Timestamp of the latest affliction date of an active finding.
	LatestAfflictionDate time.Time `json:"latest_affliction_date" api:"required" format:"date-time"`
	// Override information for finding severity.
	SeverityOverride CasbPostureFindingResetSeverityResponseSeverityOverride `json:"severity_override"`
	JSON             casbPostureFindingResetSeverityResponseJSON             `json:"-"`
}

// casbPostureFindingResetSeverityResponseJSON contains the JSON metadata for the
// struct [CasbPostureFindingResetSeverityResponse]
type casbPostureFindingResetSeverityResponseJSON struct {
	ID                   apijson.Field
	ActiveCount          apijson.Field
	ArchivedCount        apijson.Field
	Finding              apijson.Field
	Ignored              apijson.Field
	InstanceCount        apijson.Field
	Integration          apijson.Field
	LatestAfflictionDate apijson.Field
	SeverityOverride     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseJSON) RawJSON() string {
	return r.raw
}

// Basic finding type information.
type CasbPostureFindingResetSeverityResponseFinding struct {
	// The unique identifier of the finding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Category information for a finding.
	Category CasbPostureFindingResetSeverityResponseFindingCategory `json:"category" api:"required"`
	// The name of the finding.
	Name string `json:"name" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingResetSeverityResponseFindingSeverity `json:"severity" api:"required"`
	// The SaaS/Cloud vendor of the platform with which the finding is associated.
	Vendor string `json:"vendor" api:"required"`
	// Detailed description of the finding.
	Description string `json:"description" api:"nullable"`
	// Remediation guide information for a finding.
	Remediation CasbPostureFindingResetSeverityResponseFindingRemediation `json:"remediation" api:"nullable"`
	JSON        casbPostureFindingResetSeverityResponseFindingJSON        `json:"-"`
}

// casbPostureFindingResetSeverityResponseFindingJSON contains the JSON metadata
// for the struct [CasbPostureFindingResetSeverityResponseFinding]
type casbPostureFindingResetSeverityResponseFindingJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	Name        apijson.Field
	Severity    apijson.Field
	Vendor      apijson.Field
	Description apijson.Field
	Remediation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponseFinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseFindingJSON) RawJSON() string {
	return r.raw
}

// Category information for a finding.
type CasbPostureFindingResetSeverityResponseFindingCategory struct {
	// The type of the observation.
	Observation CasbPostureFindingResetSeverityResponseFindingCategoryObservation `json:"observation" api:"required"`
	// The product category.
	Product CasbPostureFindingResetSeverityResponseFindingCategoryProduct `json:"product" api:"required"`
	// The type of the finding category.
	Type CasbPostureFindingResetSeverityResponseFindingCategoryType `json:"type" api:"required"`
	JSON casbPostureFindingResetSeverityResponseFindingCategoryJSON `json:"-"`
}

// casbPostureFindingResetSeverityResponseFindingCategoryJSON contains the JSON
// metadata for the struct [CasbPostureFindingResetSeverityResponseFindingCategory]
type casbPostureFindingResetSeverityResponseFindingCategoryJSON struct {
	Observation apijson.Field
	Product     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponseFindingCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseFindingCategoryJSON) RawJSON() string {
	return r.raw
}

// The type of the observation.
type CasbPostureFindingResetSeverityResponseFindingCategoryObservation string

const (
	CasbPostureFindingResetSeverityResponseFindingCategoryObservationIssue    CasbPostureFindingResetSeverityResponseFindingCategoryObservation = "Issue"
	CasbPostureFindingResetSeverityResponseFindingCategoryObservationInsight  CasbPostureFindingResetSeverityResponseFindingCategoryObservation = "Insight"
	CasbPostureFindingResetSeverityResponseFindingCategoryObservationActivity CasbPostureFindingResetSeverityResponseFindingCategoryObservation = "Activity"
)

func (r CasbPostureFindingResetSeverityResponseFindingCategoryObservation) IsKnown() bool {
	switch r {
	case CasbPostureFindingResetSeverityResponseFindingCategoryObservationIssue, CasbPostureFindingResetSeverityResponseFindingCategoryObservationInsight, CasbPostureFindingResetSeverityResponseFindingCategoryObservationActivity:
		return true
	}
	return false
}

// The product category.
type CasbPostureFindingResetSeverityResponseFindingCategoryProduct string

const (
	CasbPostureFindingResetSeverityResponseFindingCategoryProductSaaS  CasbPostureFindingResetSeverityResponseFindingCategoryProduct = "SaaS"
	CasbPostureFindingResetSeverityResponseFindingCategoryProductCloud CasbPostureFindingResetSeverityResponseFindingCategoryProduct = "Cloud"
)

func (r CasbPostureFindingResetSeverityResponseFindingCategoryProduct) IsKnown() bool {
	switch r {
	case CasbPostureFindingResetSeverityResponseFindingCategoryProductSaaS, CasbPostureFindingResetSeverityResponseFindingCategoryProductCloud:
		return true
	}
	return false
}

// The type of the finding category.
type CasbPostureFindingResetSeverityResponseFindingCategoryType string

const (
	CasbPostureFindingResetSeverityResponseFindingCategoryTypeContent CasbPostureFindingResetSeverityResponseFindingCategoryType = "Content"
	CasbPostureFindingResetSeverityResponseFindingCategoryTypePosture CasbPostureFindingResetSeverityResponseFindingCategoryType = "Posture"
)

func (r CasbPostureFindingResetSeverityResponseFindingCategoryType) IsKnown() bool {
	switch r {
	case CasbPostureFindingResetSeverityResponseFindingCategoryTypeContent, CasbPostureFindingResetSeverityResponseFindingCategoryTypePosture:
		return true
	}
	return false
}

// The severity level of a finding.
type CasbPostureFindingResetSeverityResponseFindingSeverity string

const (
	CasbPostureFindingResetSeverityResponseFindingSeverityCritical CasbPostureFindingResetSeverityResponseFindingSeverity = "Critical"
	CasbPostureFindingResetSeverityResponseFindingSeverityHigh     CasbPostureFindingResetSeverityResponseFindingSeverity = "High"
	CasbPostureFindingResetSeverityResponseFindingSeverityMedium   CasbPostureFindingResetSeverityResponseFindingSeverity = "Medium"
	CasbPostureFindingResetSeverityResponseFindingSeverityLow      CasbPostureFindingResetSeverityResponseFindingSeverity = "Low"
)

func (r CasbPostureFindingResetSeverityResponseFindingSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingResetSeverityResponseFindingSeverityCritical, CasbPostureFindingResetSeverityResponseFindingSeverityHigh, CasbPostureFindingResetSeverityResponseFindingSeverityMedium, CasbPostureFindingResetSeverityResponseFindingSeverityLow:
		return true
	}
	return false
}

// Remediation guide information for a finding.
type CasbPostureFindingResetSeverityResponseFindingRemediation struct {
	// Remediation Id.
	ID string `json:"id" api:"required" format:"uuid"`
	// Relevant Compliance Frameworks.
	Frameworks []string `json:"frameworks" api:"required"`
	// Remediation guide text.
	Guide string `json:"guide" api:"required"`
	// Description of the potential impact.
	Impact string `json:"impact" api:"required"`
	// I18N Locale.
	Locale string `json:"locale" api:"required"`
	// Description of the threat.
	Threat string                                                        `json:"threat" api:"required"`
	JSON   casbPostureFindingResetSeverityResponseFindingRemediationJSON `json:"-"`
}

// casbPostureFindingResetSeverityResponseFindingRemediationJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingResetSeverityResponseFindingRemediation]
type casbPostureFindingResetSeverityResponseFindingRemediationJSON struct {
	ID          apijson.Field
	Frameworks  apijson.Field
	Guide       apijson.Field
	Impact      apijson.Field
	Locale      apijson.Field
	Threat      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponseFindingRemediation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseFindingRemediationJSON) RawJSON() string {
	return r.raw
}

// Summary information about an integration.
type CasbPostureFindingResetSeverityResponseIntegration struct {
	// When entity was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// When were the integration credentials last updated.
	LastHydrated time.Time `json:"last_hydrated" api:"required" format:"date-time"`
	// Name of the integration.
	Name string `json:"name" api:"required"`
	// The vendor-specific permissions associated with the integration.
	Permissions []string `json:"permissions" api:"required"`
	// Policy configuration for an integration.
	Policy CasbPostureFindingResetSeverityResponseIntegrationPolicy `json:"policy" api:"required"`
	// Current status of the integration.
	Status string `json:"status" api:"required"`
	// Last entity was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Whether the integrations permissions can be updated.
	Upgradable bool `json:"upgradable" api:"required"`
	// Information about a vendor/service provider.
	Vendor CasbPostureFindingResetSeverityResponseIntegrationVendor `json:"vendor" api:"required"`
	// Zero Trust products associated with this integration.
	ZtEnrollments []CasbPostureFindingResetSeverityResponseIntegrationZtEnrollment `json:"zt_enrollments" api:"required"`
	// Integration ID.
	ID string `json:"id" format:"uuid"`
	// Health status of integration credentials.
	CredentialHealthStatus CasbPostureFindingResetSeverityResponseIntegrationCredentialHealthStatus `json:"credential_health_status"`
	// The date and time when the integration credentials will expire.
	CredentialsExpiry time.Time `json:"credentials_expiry" api:"nullable" format:"date-time"`
	// Whether the given integration is paused by the user.
	IsPaused bool `json:"is_paused"`
	// UI State as to whether a potential permissions upgrade has been dismissed.
	UpgradeDismissed bool                                                   `json:"upgrade_dismissed"`
	JSON             casbPostureFindingResetSeverityResponseIntegrationJSON `json:"-"`
}

// casbPostureFindingResetSeverityResponseIntegrationJSON contains the JSON
// metadata for the struct [CasbPostureFindingResetSeverityResponseIntegration]
type casbPostureFindingResetSeverityResponseIntegrationJSON struct {
	Created                apijson.Field
	LastHydrated           apijson.Field
	Name                   apijson.Field
	Permissions            apijson.Field
	Policy                 apijson.Field
	Status                 apijson.Field
	Updated                apijson.Field
	Upgradable             apijson.Field
	Vendor                 apijson.Field
	ZtEnrollments          apijson.Field
	ID                     apijson.Field
	CredentialHealthStatus apijson.Field
	CredentialsExpiry      apijson.Field
	IsPaused               apijson.Field
	UpgradeDismissed       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

// Policy configuration for an integration.
type CasbPostureFindingResetSeverityResponseIntegrationPolicy struct {
	// Policy identifier.
	ID string `json:"id" format:"uuid"`
	// OAuth client ID for the policy.
	ClientID string `json:"client_id" api:"nullable"`
	// Compliance level for the policy.
	ComplianceLevel string `json:"compliance_level"`
	// Whether DLP is enabled for this policy.
	DLPEnabled bool `json:"dlp_enabled"`
	// Link to policy documentation.
	Link string `json:"link" api:"nullable" format:"uri"`
	// Policy name.
	Name string `json:"name"`
	// List of permissions included in the policy.
	Permissions []string                                                     `json:"permissions"`
	JSON        casbPostureFindingResetSeverityResponseIntegrationPolicyJSON `json:"-"`
}

// casbPostureFindingResetSeverityResponseIntegrationPolicyJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingResetSeverityResponseIntegrationPolicy]
type casbPostureFindingResetSeverityResponseIntegrationPolicyJSON struct {
	ID              apijson.Field
	ClientID        apijson.Field
	ComplianceLevel apijson.Field
	DLPEnabled      apijson.Field
	Link            apijson.Field
	Name            apijson.Field
	Permissions     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponseIntegrationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseIntegrationPolicyJSON) RawJSON() string {
	return r.raw
}

// Information about a vendor/service provider.
type CasbPostureFindingResetSeverityResponseIntegrationVendor struct {
	// The id of the vendor.
	ID string `json:"id" api:"required"`
	// Detailed information about what kinds of issues are detected for this vendor.
	Description string `json:"description" api:"required,nullable"`
	// The display name of the vendor.
	DisplayName string `json:"display_name" api:"required"`
	// Logo URL for the vendor.
	Logo string `json:"logo" api:"required" format:"uri"`
	// The name of the vendor.
	Name string `json:"name" api:"required"`
	// Static logo URL for the vendor.
	StaticLogo string `json:"static_logo" api:"required" format:"uri"`
	// The vendor's compatible Zero Trust products.
	ZtEnrollments []string `json:"zt_enrollments" api:"required"`
	// The policies related to the vendor.
	Policies []map[string]interface{}                                     `json:"policies"`
	JSON     casbPostureFindingResetSeverityResponseIntegrationVendorJSON `json:"-"`
}

// casbPostureFindingResetSeverityResponseIntegrationVendorJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingResetSeverityResponseIntegrationVendor]
type casbPostureFindingResetSeverityResponseIntegrationVendorJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Logo          apijson.Field
	Name          apijson.Field
	StaticLogo    apijson.Field
	ZtEnrollments apijson.Field
	Policies      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponseIntegrationVendor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseIntegrationVendorJSON) RawJSON() string {
	return r.raw
}

// Information about a Zero Trust product integration.
type CasbPostureFindingResetSeverityResponseIntegrationZtEnrollment struct {
	// The internal identifier of the Zero Trust Product.
	ID string `json:"id"`
	// Brief description of the Zero Trust Product.
	Description string `json:"description"`
	// The verbose name of the Zero Trust Product.
	DisplayName string `json:"display_name"`
	// Flag to enable/disable access to the listed integration from the corresponding
	// Cloudflare product.
	Enabled bool                                                               `json:"enabled"`
	JSON    casbPostureFindingResetSeverityResponseIntegrationZtEnrollmentJSON `json:"-"`
}

// casbPostureFindingResetSeverityResponseIntegrationZtEnrollmentJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingResetSeverityResponseIntegrationZtEnrollment]
type casbPostureFindingResetSeverityResponseIntegrationZtEnrollmentJSON struct {
	ID          apijson.Field
	Description apijson.Field
	DisplayName apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponseIntegrationZtEnrollment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseIntegrationZtEnrollmentJSON) RawJSON() string {
	return r.raw
}

// Health status of integration credentials.
type CasbPostureFindingResetSeverityResponseIntegrationCredentialHealthStatus string

const (
	CasbPostureFindingResetSeverityResponseIntegrationCredentialHealthStatusInitializing CasbPostureFindingResetSeverityResponseIntegrationCredentialHealthStatus = "Initializing"
	CasbPostureFindingResetSeverityResponseIntegrationCredentialHealthStatusHealthy      CasbPostureFindingResetSeverityResponseIntegrationCredentialHealthStatus = "Healthy"
	CasbPostureFindingResetSeverityResponseIntegrationCredentialHealthStatusUnhealthy    CasbPostureFindingResetSeverityResponseIntegrationCredentialHealthStatus = "Unhealthy"
)

func (r CasbPostureFindingResetSeverityResponseIntegrationCredentialHealthStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingResetSeverityResponseIntegrationCredentialHealthStatusInitializing, CasbPostureFindingResetSeverityResponseIntegrationCredentialHealthStatusHealthy, CasbPostureFindingResetSeverityResponseIntegrationCredentialHealthStatusUnhealthy:
		return true
	}
	return false
}

// Override information for finding severity.
type CasbPostureFindingResetSeverityResponseSeverityOverride struct {
	// User ID who created the override.
	CreatedBy string `json:"created_by" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingResetSeverityResponseSeverityOverrideSeverity `json:"severity" api:"required"`
	JSON     casbPostureFindingResetSeverityResponseSeverityOverrideJSON     `json:"-"`
}

// casbPostureFindingResetSeverityResponseSeverityOverrideJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingResetSeverityResponseSeverityOverride]
type casbPostureFindingResetSeverityResponseSeverityOverrideJSON struct {
	CreatedBy   apijson.Field
	Severity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponseSeverityOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseSeverityOverrideJSON) RawJSON() string {
	return r.raw
}

// The severity level of a finding.
type CasbPostureFindingResetSeverityResponseSeverityOverrideSeverity string

const (
	CasbPostureFindingResetSeverityResponseSeverityOverrideSeverityCritical CasbPostureFindingResetSeverityResponseSeverityOverrideSeverity = "Critical"
	CasbPostureFindingResetSeverityResponseSeverityOverrideSeverityHigh     CasbPostureFindingResetSeverityResponseSeverityOverrideSeverity = "High"
	CasbPostureFindingResetSeverityResponseSeverityOverrideSeverityMedium   CasbPostureFindingResetSeverityResponseSeverityOverrideSeverity = "Medium"
	CasbPostureFindingResetSeverityResponseSeverityOverrideSeverityLow      CasbPostureFindingResetSeverityResponseSeverityOverrideSeverity = "Low"
)

func (r CasbPostureFindingResetSeverityResponseSeverityOverrideSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingResetSeverityResponseSeverityOverrideSeverityCritical, CasbPostureFindingResetSeverityResponseSeverityOverrideSeverityHigh, CasbPostureFindingResetSeverityResponseSeverityOverrideSeverityMedium, CasbPostureFindingResetSeverityResponseSeverityOverrideSeverityLow:
		return true
	}
	return false
}

// Aggregated finding information with counts and metadata. This is optimized for
// list API queries and represents a finding along with its instance statistics.
type CasbPostureFindingTuneSeverityResponse struct {
	// Base64 encoded identifier of the security finding.
	ID string `json:"id" api:"required" format:"byte"`
	// Number of active problematic instances identified in the security finding.
	ActiveCount int64 `json:"active_count" api:"required"`
	// Number of archived instances identified in the security finding.
	ArchivedCount int64 `json:"archived_count" api:"required"`
	// Basic finding type information.
	Finding CasbPostureFindingTuneSeverityResponseFinding `json:"finding" api:"required"`
	// Determines if finding is currently ignored.
	Ignored bool `json:"ignored" api:"required"`
	// Number of total (Active or archived) problematic instances identified in the
	// security finding.
	InstanceCount int64 `json:"instance_count" api:"required"`
	// Summary information about an integration.
	Integration CasbPostureFindingTuneSeverityResponseIntegration `json:"integration" api:"required"`
	// Timestamp of the latest affliction date of an active finding.
	LatestAfflictionDate time.Time `json:"latest_affliction_date" api:"required" format:"date-time"`
	// Override information for finding severity.
	SeverityOverride CasbPostureFindingTuneSeverityResponseSeverityOverride `json:"severity_override"`
	JSON             casbPostureFindingTuneSeverityResponseJSON             `json:"-"`
}

// casbPostureFindingTuneSeverityResponseJSON contains the JSON metadata for the
// struct [CasbPostureFindingTuneSeverityResponse]
type casbPostureFindingTuneSeverityResponseJSON struct {
	ID                   apijson.Field
	ActiveCount          apijson.Field
	ArchivedCount        apijson.Field
	Finding              apijson.Field
	Ignored              apijson.Field
	InstanceCount        apijson.Field
	Integration          apijson.Field
	LatestAfflictionDate apijson.Field
	SeverityOverride     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseJSON) RawJSON() string {
	return r.raw
}

// Basic finding type information.
type CasbPostureFindingTuneSeverityResponseFinding struct {
	// The unique identifier of the finding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Category information for a finding.
	Category CasbPostureFindingTuneSeverityResponseFindingCategory `json:"category" api:"required"`
	// The name of the finding.
	Name string `json:"name" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingTuneSeverityResponseFindingSeverity `json:"severity" api:"required"`
	// The SaaS/Cloud vendor of the platform with which the finding is associated.
	Vendor string `json:"vendor" api:"required"`
	// Detailed description of the finding.
	Description string `json:"description" api:"nullable"`
	// Remediation guide information for a finding.
	Remediation CasbPostureFindingTuneSeverityResponseFindingRemediation `json:"remediation" api:"nullable"`
	JSON        casbPostureFindingTuneSeverityResponseFindingJSON        `json:"-"`
}

// casbPostureFindingTuneSeverityResponseFindingJSON contains the JSON metadata for
// the struct [CasbPostureFindingTuneSeverityResponseFinding]
type casbPostureFindingTuneSeverityResponseFindingJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	Name        apijson.Field
	Severity    apijson.Field
	Vendor      apijson.Field
	Description apijson.Field
	Remediation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponseFinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseFindingJSON) RawJSON() string {
	return r.raw
}

// Category information for a finding.
type CasbPostureFindingTuneSeverityResponseFindingCategory struct {
	// The type of the observation.
	Observation CasbPostureFindingTuneSeverityResponseFindingCategoryObservation `json:"observation" api:"required"`
	// The product category.
	Product CasbPostureFindingTuneSeverityResponseFindingCategoryProduct `json:"product" api:"required"`
	// The type of the finding category.
	Type CasbPostureFindingTuneSeverityResponseFindingCategoryType `json:"type" api:"required"`
	JSON casbPostureFindingTuneSeverityResponseFindingCategoryJSON `json:"-"`
}

// casbPostureFindingTuneSeverityResponseFindingCategoryJSON contains the JSON
// metadata for the struct [CasbPostureFindingTuneSeverityResponseFindingCategory]
type casbPostureFindingTuneSeverityResponseFindingCategoryJSON struct {
	Observation apijson.Field
	Product     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponseFindingCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseFindingCategoryJSON) RawJSON() string {
	return r.raw
}

// The type of the observation.
type CasbPostureFindingTuneSeverityResponseFindingCategoryObservation string

const (
	CasbPostureFindingTuneSeverityResponseFindingCategoryObservationIssue    CasbPostureFindingTuneSeverityResponseFindingCategoryObservation = "Issue"
	CasbPostureFindingTuneSeverityResponseFindingCategoryObservationInsight  CasbPostureFindingTuneSeverityResponseFindingCategoryObservation = "Insight"
	CasbPostureFindingTuneSeverityResponseFindingCategoryObservationActivity CasbPostureFindingTuneSeverityResponseFindingCategoryObservation = "Activity"
)

func (r CasbPostureFindingTuneSeverityResponseFindingCategoryObservation) IsKnown() bool {
	switch r {
	case CasbPostureFindingTuneSeverityResponseFindingCategoryObservationIssue, CasbPostureFindingTuneSeverityResponseFindingCategoryObservationInsight, CasbPostureFindingTuneSeverityResponseFindingCategoryObservationActivity:
		return true
	}
	return false
}

// The product category.
type CasbPostureFindingTuneSeverityResponseFindingCategoryProduct string

const (
	CasbPostureFindingTuneSeverityResponseFindingCategoryProductSaaS  CasbPostureFindingTuneSeverityResponseFindingCategoryProduct = "SaaS"
	CasbPostureFindingTuneSeverityResponseFindingCategoryProductCloud CasbPostureFindingTuneSeverityResponseFindingCategoryProduct = "Cloud"
)

func (r CasbPostureFindingTuneSeverityResponseFindingCategoryProduct) IsKnown() bool {
	switch r {
	case CasbPostureFindingTuneSeverityResponseFindingCategoryProductSaaS, CasbPostureFindingTuneSeverityResponseFindingCategoryProductCloud:
		return true
	}
	return false
}

// The type of the finding category.
type CasbPostureFindingTuneSeverityResponseFindingCategoryType string

const (
	CasbPostureFindingTuneSeverityResponseFindingCategoryTypeContent CasbPostureFindingTuneSeverityResponseFindingCategoryType = "Content"
	CasbPostureFindingTuneSeverityResponseFindingCategoryTypePosture CasbPostureFindingTuneSeverityResponseFindingCategoryType = "Posture"
)

func (r CasbPostureFindingTuneSeverityResponseFindingCategoryType) IsKnown() bool {
	switch r {
	case CasbPostureFindingTuneSeverityResponseFindingCategoryTypeContent, CasbPostureFindingTuneSeverityResponseFindingCategoryTypePosture:
		return true
	}
	return false
}

// The severity level of a finding.
type CasbPostureFindingTuneSeverityResponseFindingSeverity string

const (
	CasbPostureFindingTuneSeverityResponseFindingSeverityCritical CasbPostureFindingTuneSeverityResponseFindingSeverity = "Critical"
	CasbPostureFindingTuneSeverityResponseFindingSeverityHigh     CasbPostureFindingTuneSeverityResponseFindingSeverity = "High"
	CasbPostureFindingTuneSeverityResponseFindingSeverityMedium   CasbPostureFindingTuneSeverityResponseFindingSeverity = "Medium"
	CasbPostureFindingTuneSeverityResponseFindingSeverityLow      CasbPostureFindingTuneSeverityResponseFindingSeverity = "Low"
)

func (r CasbPostureFindingTuneSeverityResponseFindingSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingTuneSeverityResponseFindingSeverityCritical, CasbPostureFindingTuneSeverityResponseFindingSeverityHigh, CasbPostureFindingTuneSeverityResponseFindingSeverityMedium, CasbPostureFindingTuneSeverityResponseFindingSeverityLow:
		return true
	}
	return false
}

// Remediation guide information for a finding.
type CasbPostureFindingTuneSeverityResponseFindingRemediation struct {
	// Remediation Id.
	ID string `json:"id" api:"required" format:"uuid"`
	// Relevant Compliance Frameworks.
	Frameworks []string `json:"frameworks" api:"required"`
	// Remediation guide text.
	Guide string `json:"guide" api:"required"`
	// Description of the potential impact.
	Impact string `json:"impact" api:"required"`
	// I18N Locale.
	Locale string `json:"locale" api:"required"`
	// Description of the threat.
	Threat string                                                       `json:"threat" api:"required"`
	JSON   casbPostureFindingTuneSeverityResponseFindingRemediationJSON `json:"-"`
}

// casbPostureFindingTuneSeverityResponseFindingRemediationJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingTuneSeverityResponseFindingRemediation]
type casbPostureFindingTuneSeverityResponseFindingRemediationJSON struct {
	ID          apijson.Field
	Frameworks  apijson.Field
	Guide       apijson.Field
	Impact      apijson.Field
	Locale      apijson.Field
	Threat      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponseFindingRemediation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseFindingRemediationJSON) RawJSON() string {
	return r.raw
}

// Summary information about an integration.
type CasbPostureFindingTuneSeverityResponseIntegration struct {
	// When entity was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// When were the integration credentials last updated.
	LastHydrated time.Time `json:"last_hydrated" api:"required" format:"date-time"`
	// Name of the integration.
	Name string `json:"name" api:"required"`
	// The vendor-specific permissions associated with the integration.
	Permissions []string `json:"permissions" api:"required"`
	// Policy configuration for an integration.
	Policy CasbPostureFindingTuneSeverityResponseIntegrationPolicy `json:"policy" api:"required"`
	// Current status of the integration.
	Status string `json:"status" api:"required"`
	// Last entity was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Whether the integrations permissions can be updated.
	Upgradable bool `json:"upgradable" api:"required"`
	// Information about a vendor/service provider.
	Vendor CasbPostureFindingTuneSeverityResponseIntegrationVendor `json:"vendor" api:"required"`
	// Zero Trust products associated with this integration.
	ZtEnrollments []CasbPostureFindingTuneSeverityResponseIntegrationZtEnrollment `json:"zt_enrollments" api:"required"`
	// Integration ID.
	ID string `json:"id" format:"uuid"`
	// Health status of integration credentials.
	CredentialHealthStatus CasbPostureFindingTuneSeverityResponseIntegrationCredentialHealthStatus `json:"credential_health_status"`
	// The date and time when the integration credentials will expire.
	CredentialsExpiry time.Time `json:"credentials_expiry" api:"nullable" format:"date-time"`
	// Whether the given integration is paused by the user.
	IsPaused bool `json:"is_paused"`
	// UI State as to whether a potential permissions upgrade has been dismissed.
	UpgradeDismissed bool                                                  `json:"upgrade_dismissed"`
	JSON             casbPostureFindingTuneSeverityResponseIntegrationJSON `json:"-"`
}

// casbPostureFindingTuneSeverityResponseIntegrationJSON contains the JSON metadata
// for the struct [CasbPostureFindingTuneSeverityResponseIntegration]
type casbPostureFindingTuneSeverityResponseIntegrationJSON struct {
	Created                apijson.Field
	LastHydrated           apijson.Field
	Name                   apijson.Field
	Permissions            apijson.Field
	Policy                 apijson.Field
	Status                 apijson.Field
	Updated                apijson.Field
	Upgradable             apijson.Field
	Vendor                 apijson.Field
	ZtEnrollments          apijson.Field
	ID                     apijson.Field
	CredentialHealthStatus apijson.Field
	CredentialsExpiry      apijson.Field
	IsPaused               apijson.Field
	UpgradeDismissed       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

// Policy configuration for an integration.
type CasbPostureFindingTuneSeverityResponseIntegrationPolicy struct {
	// Policy identifier.
	ID string `json:"id" format:"uuid"`
	// OAuth client ID for the policy.
	ClientID string `json:"client_id" api:"nullable"`
	// Compliance level for the policy.
	ComplianceLevel string `json:"compliance_level"`
	// Whether DLP is enabled for this policy.
	DLPEnabled bool `json:"dlp_enabled"`
	// Link to policy documentation.
	Link string `json:"link" api:"nullable" format:"uri"`
	// Policy name.
	Name string `json:"name"`
	// List of permissions included in the policy.
	Permissions []string                                                    `json:"permissions"`
	JSON        casbPostureFindingTuneSeverityResponseIntegrationPolicyJSON `json:"-"`
}

// casbPostureFindingTuneSeverityResponseIntegrationPolicyJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingTuneSeverityResponseIntegrationPolicy]
type casbPostureFindingTuneSeverityResponseIntegrationPolicyJSON struct {
	ID              apijson.Field
	ClientID        apijson.Field
	ComplianceLevel apijson.Field
	DLPEnabled      apijson.Field
	Link            apijson.Field
	Name            apijson.Field
	Permissions     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponseIntegrationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseIntegrationPolicyJSON) RawJSON() string {
	return r.raw
}

// Information about a vendor/service provider.
type CasbPostureFindingTuneSeverityResponseIntegrationVendor struct {
	// The id of the vendor.
	ID string `json:"id" api:"required"`
	// Detailed information about what kinds of issues are detected for this vendor.
	Description string `json:"description" api:"required,nullable"`
	// The display name of the vendor.
	DisplayName string `json:"display_name" api:"required"`
	// Logo URL for the vendor.
	Logo string `json:"logo" api:"required" format:"uri"`
	// The name of the vendor.
	Name string `json:"name" api:"required"`
	// Static logo URL for the vendor.
	StaticLogo string `json:"static_logo" api:"required" format:"uri"`
	// The vendor's compatible Zero Trust products.
	ZtEnrollments []string `json:"zt_enrollments" api:"required"`
	// The policies related to the vendor.
	Policies []map[string]interface{}                                    `json:"policies"`
	JSON     casbPostureFindingTuneSeverityResponseIntegrationVendorJSON `json:"-"`
}

// casbPostureFindingTuneSeverityResponseIntegrationVendorJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingTuneSeverityResponseIntegrationVendor]
type casbPostureFindingTuneSeverityResponseIntegrationVendorJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Logo          apijson.Field
	Name          apijson.Field
	StaticLogo    apijson.Field
	ZtEnrollments apijson.Field
	Policies      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponseIntegrationVendor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseIntegrationVendorJSON) RawJSON() string {
	return r.raw
}

// Information about a Zero Trust product integration.
type CasbPostureFindingTuneSeverityResponseIntegrationZtEnrollment struct {
	// The internal identifier of the Zero Trust Product.
	ID string `json:"id"`
	// Brief description of the Zero Trust Product.
	Description string `json:"description"`
	// The verbose name of the Zero Trust Product.
	DisplayName string `json:"display_name"`
	// Flag to enable/disable access to the listed integration from the corresponding
	// Cloudflare product.
	Enabled bool                                                              `json:"enabled"`
	JSON    casbPostureFindingTuneSeverityResponseIntegrationZtEnrollmentJSON `json:"-"`
}

// casbPostureFindingTuneSeverityResponseIntegrationZtEnrollmentJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingTuneSeverityResponseIntegrationZtEnrollment]
type casbPostureFindingTuneSeverityResponseIntegrationZtEnrollmentJSON struct {
	ID          apijson.Field
	Description apijson.Field
	DisplayName apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponseIntegrationZtEnrollment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseIntegrationZtEnrollmentJSON) RawJSON() string {
	return r.raw
}

// Health status of integration credentials.
type CasbPostureFindingTuneSeverityResponseIntegrationCredentialHealthStatus string

const (
	CasbPostureFindingTuneSeverityResponseIntegrationCredentialHealthStatusInitializing CasbPostureFindingTuneSeverityResponseIntegrationCredentialHealthStatus = "Initializing"
	CasbPostureFindingTuneSeverityResponseIntegrationCredentialHealthStatusHealthy      CasbPostureFindingTuneSeverityResponseIntegrationCredentialHealthStatus = "Healthy"
	CasbPostureFindingTuneSeverityResponseIntegrationCredentialHealthStatusUnhealthy    CasbPostureFindingTuneSeverityResponseIntegrationCredentialHealthStatus = "Unhealthy"
)

func (r CasbPostureFindingTuneSeverityResponseIntegrationCredentialHealthStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingTuneSeverityResponseIntegrationCredentialHealthStatusInitializing, CasbPostureFindingTuneSeverityResponseIntegrationCredentialHealthStatusHealthy, CasbPostureFindingTuneSeverityResponseIntegrationCredentialHealthStatusUnhealthy:
		return true
	}
	return false
}

// Override information for finding severity.
type CasbPostureFindingTuneSeverityResponseSeverityOverride struct {
	// User ID who created the override.
	CreatedBy string `json:"created_by" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverity `json:"severity" api:"required"`
	JSON     casbPostureFindingTuneSeverityResponseSeverityOverrideJSON     `json:"-"`
}

// casbPostureFindingTuneSeverityResponseSeverityOverrideJSON contains the JSON
// metadata for the struct [CasbPostureFindingTuneSeverityResponseSeverityOverride]
type casbPostureFindingTuneSeverityResponseSeverityOverrideJSON struct {
	CreatedBy   apijson.Field
	Severity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponseSeverityOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseSeverityOverrideJSON) RawJSON() string {
	return r.raw
}

// The severity level of a finding.
type CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverity string

const (
	CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverityCritical CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverity = "Critical"
	CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverityHigh     CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverity = "High"
	CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverityMedium   CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverity = "Medium"
	CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverityLow      CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverity = "Low"
)

func (r CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverityCritical, CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverityHigh, CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverityMedium, CasbPostureFindingTuneSeverityResponseSeverityOverrideSeverityLow:
		return true
	}
	return false
}

// Aggregated finding information with counts and metadata. This is optimized for
// list API queries and represents a finding along with its instance statistics.
type CasbPostureFindingUnignoreResponse struct {
	// Base64 encoded identifier of the security finding.
	ID string `json:"id" api:"required" format:"byte"`
	// Number of active problematic instances identified in the security finding.
	ActiveCount int64 `json:"active_count" api:"required"`
	// Number of archived instances identified in the security finding.
	ArchivedCount int64 `json:"archived_count" api:"required"`
	// Basic finding type information.
	Finding CasbPostureFindingUnignoreResponseFinding `json:"finding" api:"required"`
	// Determines if finding is currently ignored.
	Ignored bool `json:"ignored" api:"required"`
	// Number of total (Active or archived) problematic instances identified in the
	// security finding.
	InstanceCount int64 `json:"instance_count" api:"required"`
	// Summary information about an integration.
	Integration CasbPostureFindingUnignoreResponseIntegration `json:"integration" api:"required"`
	// Timestamp of the latest affliction date of an active finding.
	LatestAfflictionDate time.Time `json:"latest_affliction_date" api:"required" format:"date-time"`
	// Override information for finding severity.
	SeverityOverride CasbPostureFindingUnignoreResponseSeverityOverride `json:"severity_override"`
	JSON             casbPostureFindingUnignoreResponseJSON             `json:"-"`
}

// casbPostureFindingUnignoreResponseJSON contains the JSON metadata for the struct
// [CasbPostureFindingUnignoreResponse]
type casbPostureFindingUnignoreResponseJSON struct {
	ID                   apijson.Field
	ActiveCount          apijson.Field
	ArchivedCount        apijson.Field
	Finding              apijson.Field
	Ignored              apijson.Field
	InstanceCount        apijson.Field
	Integration          apijson.Field
	LatestAfflictionDate apijson.Field
	SeverityOverride     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseJSON) RawJSON() string {
	return r.raw
}

// Basic finding type information.
type CasbPostureFindingUnignoreResponseFinding struct {
	// The unique identifier of the finding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Category information for a finding.
	Category CasbPostureFindingUnignoreResponseFindingCategory `json:"category" api:"required"`
	// The name of the finding.
	Name string `json:"name" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingUnignoreResponseFindingSeverity `json:"severity" api:"required"`
	// The SaaS/Cloud vendor of the platform with which the finding is associated.
	Vendor string `json:"vendor" api:"required"`
	// Detailed description of the finding.
	Description string `json:"description" api:"nullable"`
	// Remediation guide information for a finding.
	Remediation CasbPostureFindingUnignoreResponseFindingRemediation `json:"remediation" api:"nullable"`
	JSON        casbPostureFindingUnignoreResponseFindingJSON        `json:"-"`
}

// casbPostureFindingUnignoreResponseFindingJSON contains the JSON metadata for the
// struct [CasbPostureFindingUnignoreResponseFinding]
type casbPostureFindingUnignoreResponseFindingJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	Name        apijson.Field
	Severity    apijson.Field
	Vendor      apijson.Field
	Description apijson.Field
	Remediation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponseFinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseFindingJSON) RawJSON() string {
	return r.raw
}

// Category information for a finding.
type CasbPostureFindingUnignoreResponseFindingCategory struct {
	// The type of the observation.
	Observation CasbPostureFindingUnignoreResponseFindingCategoryObservation `json:"observation" api:"required"`
	// The product category.
	Product CasbPostureFindingUnignoreResponseFindingCategoryProduct `json:"product" api:"required"`
	// The type of the finding category.
	Type CasbPostureFindingUnignoreResponseFindingCategoryType `json:"type" api:"required"`
	JSON casbPostureFindingUnignoreResponseFindingCategoryJSON `json:"-"`
}

// casbPostureFindingUnignoreResponseFindingCategoryJSON contains the JSON metadata
// for the struct [CasbPostureFindingUnignoreResponseFindingCategory]
type casbPostureFindingUnignoreResponseFindingCategoryJSON struct {
	Observation apijson.Field
	Product     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponseFindingCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseFindingCategoryJSON) RawJSON() string {
	return r.raw
}

// The type of the observation.
type CasbPostureFindingUnignoreResponseFindingCategoryObservation string

const (
	CasbPostureFindingUnignoreResponseFindingCategoryObservationIssue    CasbPostureFindingUnignoreResponseFindingCategoryObservation = "Issue"
	CasbPostureFindingUnignoreResponseFindingCategoryObservationInsight  CasbPostureFindingUnignoreResponseFindingCategoryObservation = "Insight"
	CasbPostureFindingUnignoreResponseFindingCategoryObservationActivity CasbPostureFindingUnignoreResponseFindingCategoryObservation = "Activity"
)

func (r CasbPostureFindingUnignoreResponseFindingCategoryObservation) IsKnown() bool {
	switch r {
	case CasbPostureFindingUnignoreResponseFindingCategoryObservationIssue, CasbPostureFindingUnignoreResponseFindingCategoryObservationInsight, CasbPostureFindingUnignoreResponseFindingCategoryObservationActivity:
		return true
	}
	return false
}

// The product category.
type CasbPostureFindingUnignoreResponseFindingCategoryProduct string

const (
	CasbPostureFindingUnignoreResponseFindingCategoryProductSaaS  CasbPostureFindingUnignoreResponseFindingCategoryProduct = "SaaS"
	CasbPostureFindingUnignoreResponseFindingCategoryProductCloud CasbPostureFindingUnignoreResponseFindingCategoryProduct = "Cloud"
)

func (r CasbPostureFindingUnignoreResponseFindingCategoryProduct) IsKnown() bool {
	switch r {
	case CasbPostureFindingUnignoreResponseFindingCategoryProductSaaS, CasbPostureFindingUnignoreResponseFindingCategoryProductCloud:
		return true
	}
	return false
}

// The type of the finding category.
type CasbPostureFindingUnignoreResponseFindingCategoryType string

const (
	CasbPostureFindingUnignoreResponseFindingCategoryTypeContent CasbPostureFindingUnignoreResponseFindingCategoryType = "Content"
	CasbPostureFindingUnignoreResponseFindingCategoryTypePosture CasbPostureFindingUnignoreResponseFindingCategoryType = "Posture"
)

func (r CasbPostureFindingUnignoreResponseFindingCategoryType) IsKnown() bool {
	switch r {
	case CasbPostureFindingUnignoreResponseFindingCategoryTypeContent, CasbPostureFindingUnignoreResponseFindingCategoryTypePosture:
		return true
	}
	return false
}

// The severity level of a finding.
type CasbPostureFindingUnignoreResponseFindingSeverity string

const (
	CasbPostureFindingUnignoreResponseFindingSeverityCritical CasbPostureFindingUnignoreResponseFindingSeverity = "Critical"
	CasbPostureFindingUnignoreResponseFindingSeverityHigh     CasbPostureFindingUnignoreResponseFindingSeverity = "High"
	CasbPostureFindingUnignoreResponseFindingSeverityMedium   CasbPostureFindingUnignoreResponseFindingSeverity = "Medium"
	CasbPostureFindingUnignoreResponseFindingSeverityLow      CasbPostureFindingUnignoreResponseFindingSeverity = "Low"
)

func (r CasbPostureFindingUnignoreResponseFindingSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingUnignoreResponseFindingSeverityCritical, CasbPostureFindingUnignoreResponseFindingSeverityHigh, CasbPostureFindingUnignoreResponseFindingSeverityMedium, CasbPostureFindingUnignoreResponseFindingSeverityLow:
		return true
	}
	return false
}

// Remediation guide information for a finding.
type CasbPostureFindingUnignoreResponseFindingRemediation struct {
	// Remediation Id.
	ID string `json:"id" api:"required" format:"uuid"`
	// Relevant Compliance Frameworks.
	Frameworks []string `json:"frameworks" api:"required"`
	// Remediation guide text.
	Guide string `json:"guide" api:"required"`
	// Description of the potential impact.
	Impact string `json:"impact" api:"required"`
	// I18N Locale.
	Locale string `json:"locale" api:"required"`
	// Description of the threat.
	Threat string                                                   `json:"threat" api:"required"`
	JSON   casbPostureFindingUnignoreResponseFindingRemediationJSON `json:"-"`
}

// casbPostureFindingUnignoreResponseFindingRemediationJSON contains the JSON
// metadata for the struct [CasbPostureFindingUnignoreResponseFindingRemediation]
type casbPostureFindingUnignoreResponseFindingRemediationJSON struct {
	ID          apijson.Field
	Frameworks  apijson.Field
	Guide       apijson.Field
	Impact      apijson.Field
	Locale      apijson.Field
	Threat      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponseFindingRemediation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseFindingRemediationJSON) RawJSON() string {
	return r.raw
}

// Summary information about an integration.
type CasbPostureFindingUnignoreResponseIntegration struct {
	// When entity was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// When were the integration credentials last updated.
	LastHydrated time.Time `json:"last_hydrated" api:"required" format:"date-time"`
	// Name of the integration.
	Name string `json:"name" api:"required"`
	// The vendor-specific permissions associated with the integration.
	Permissions []string `json:"permissions" api:"required"`
	// Policy configuration for an integration.
	Policy CasbPostureFindingUnignoreResponseIntegrationPolicy `json:"policy" api:"required"`
	// Current status of the integration.
	Status string `json:"status" api:"required"`
	// Last entity was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Whether the integrations permissions can be updated.
	Upgradable bool `json:"upgradable" api:"required"`
	// Information about a vendor/service provider.
	Vendor CasbPostureFindingUnignoreResponseIntegrationVendor `json:"vendor" api:"required"`
	// Zero Trust products associated with this integration.
	ZtEnrollments []CasbPostureFindingUnignoreResponseIntegrationZtEnrollment `json:"zt_enrollments" api:"required"`
	// Integration ID.
	ID string `json:"id" format:"uuid"`
	// Health status of integration credentials.
	CredentialHealthStatus CasbPostureFindingUnignoreResponseIntegrationCredentialHealthStatus `json:"credential_health_status"`
	// The date and time when the integration credentials will expire.
	CredentialsExpiry time.Time `json:"credentials_expiry" api:"nullable" format:"date-time"`
	// Whether the given integration is paused by the user.
	IsPaused bool `json:"is_paused"`
	// UI State as to whether a potential permissions upgrade has been dismissed.
	UpgradeDismissed bool                                              `json:"upgrade_dismissed"`
	JSON             casbPostureFindingUnignoreResponseIntegrationJSON `json:"-"`
}

// casbPostureFindingUnignoreResponseIntegrationJSON contains the JSON metadata for
// the struct [CasbPostureFindingUnignoreResponseIntegration]
type casbPostureFindingUnignoreResponseIntegrationJSON struct {
	Created                apijson.Field
	LastHydrated           apijson.Field
	Name                   apijson.Field
	Permissions            apijson.Field
	Policy                 apijson.Field
	Status                 apijson.Field
	Updated                apijson.Field
	Upgradable             apijson.Field
	Vendor                 apijson.Field
	ZtEnrollments          apijson.Field
	ID                     apijson.Field
	CredentialHealthStatus apijson.Field
	CredentialsExpiry      apijson.Field
	IsPaused               apijson.Field
	UpgradeDismissed       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

// Policy configuration for an integration.
type CasbPostureFindingUnignoreResponseIntegrationPolicy struct {
	// Policy identifier.
	ID string `json:"id" format:"uuid"`
	// OAuth client ID for the policy.
	ClientID string `json:"client_id" api:"nullable"`
	// Compliance level for the policy.
	ComplianceLevel string `json:"compliance_level"`
	// Whether DLP is enabled for this policy.
	DLPEnabled bool `json:"dlp_enabled"`
	// Link to policy documentation.
	Link string `json:"link" api:"nullable" format:"uri"`
	// Policy name.
	Name string `json:"name"`
	// List of permissions included in the policy.
	Permissions []string                                                `json:"permissions"`
	JSON        casbPostureFindingUnignoreResponseIntegrationPolicyJSON `json:"-"`
}

// casbPostureFindingUnignoreResponseIntegrationPolicyJSON contains the JSON
// metadata for the struct [CasbPostureFindingUnignoreResponseIntegrationPolicy]
type casbPostureFindingUnignoreResponseIntegrationPolicyJSON struct {
	ID              apijson.Field
	ClientID        apijson.Field
	ComplianceLevel apijson.Field
	DLPEnabled      apijson.Field
	Link            apijson.Field
	Name            apijson.Field
	Permissions     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponseIntegrationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseIntegrationPolicyJSON) RawJSON() string {
	return r.raw
}

// Information about a vendor/service provider.
type CasbPostureFindingUnignoreResponseIntegrationVendor struct {
	// The id of the vendor.
	ID string `json:"id" api:"required"`
	// Detailed information about what kinds of issues are detected for this vendor.
	Description string `json:"description" api:"required,nullable"`
	// The display name of the vendor.
	DisplayName string `json:"display_name" api:"required"`
	// Logo URL for the vendor.
	Logo string `json:"logo" api:"required" format:"uri"`
	// The name of the vendor.
	Name string `json:"name" api:"required"`
	// Static logo URL for the vendor.
	StaticLogo string `json:"static_logo" api:"required" format:"uri"`
	// The vendor's compatible Zero Trust products.
	ZtEnrollments []string `json:"zt_enrollments" api:"required"`
	// The policies related to the vendor.
	Policies []map[string]interface{}                                `json:"policies"`
	JSON     casbPostureFindingUnignoreResponseIntegrationVendorJSON `json:"-"`
}

// casbPostureFindingUnignoreResponseIntegrationVendorJSON contains the JSON
// metadata for the struct [CasbPostureFindingUnignoreResponseIntegrationVendor]
type casbPostureFindingUnignoreResponseIntegrationVendorJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Logo          apijson.Field
	Name          apijson.Field
	StaticLogo    apijson.Field
	ZtEnrollments apijson.Field
	Policies      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponseIntegrationVendor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseIntegrationVendorJSON) RawJSON() string {
	return r.raw
}

// Information about a Zero Trust product integration.
type CasbPostureFindingUnignoreResponseIntegrationZtEnrollment struct {
	// The internal identifier of the Zero Trust Product.
	ID string `json:"id"`
	// Brief description of the Zero Trust Product.
	Description string `json:"description"`
	// The verbose name of the Zero Trust Product.
	DisplayName string `json:"display_name"`
	// Flag to enable/disable access to the listed integration from the corresponding
	// Cloudflare product.
	Enabled bool                                                          `json:"enabled"`
	JSON    casbPostureFindingUnignoreResponseIntegrationZtEnrollmentJSON `json:"-"`
}

// casbPostureFindingUnignoreResponseIntegrationZtEnrollmentJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingUnignoreResponseIntegrationZtEnrollment]
type casbPostureFindingUnignoreResponseIntegrationZtEnrollmentJSON struct {
	ID          apijson.Field
	Description apijson.Field
	DisplayName apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponseIntegrationZtEnrollment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseIntegrationZtEnrollmentJSON) RawJSON() string {
	return r.raw
}

// Health status of integration credentials.
type CasbPostureFindingUnignoreResponseIntegrationCredentialHealthStatus string

const (
	CasbPostureFindingUnignoreResponseIntegrationCredentialHealthStatusInitializing CasbPostureFindingUnignoreResponseIntegrationCredentialHealthStatus = "Initializing"
	CasbPostureFindingUnignoreResponseIntegrationCredentialHealthStatusHealthy      CasbPostureFindingUnignoreResponseIntegrationCredentialHealthStatus = "Healthy"
	CasbPostureFindingUnignoreResponseIntegrationCredentialHealthStatusUnhealthy    CasbPostureFindingUnignoreResponseIntegrationCredentialHealthStatus = "Unhealthy"
)

func (r CasbPostureFindingUnignoreResponseIntegrationCredentialHealthStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingUnignoreResponseIntegrationCredentialHealthStatusInitializing, CasbPostureFindingUnignoreResponseIntegrationCredentialHealthStatusHealthy, CasbPostureFindingUnignoreResponseIntegrationCredentialHealthStatusUnhealthy:
		return true
	}
	return false
}

// Override information for finding severity.
type CasbPostureFindingUnignoreResponseSeverityOverride struct {
	// User ID who created the override.
	CreatedBy string `json:"created_by" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingUnignoreResponseSeverityOverrideSeverity `json:"severity" api:"required"`
	JSON     casbPostureFindingUnignoreResponseSeverityOverrideJSON     `json:"-"`
}

// casbPostureFindingUnignoreResponseSeverityOverrideJSON contains the JSON
// metadata for the struct [CasbPostureFindingUnignoreResponseSeverityOverride]
type casbPostureFindingUnignoreResponseSeverityOverrideJSON struct {
	CreatedBy   apijson.Field
	Severity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponseSeverityOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseSeverityOverrideJSON) RawJSON() string {
	return r.raw
}

// The severity level of a finding.
type CasbPostureFindingUnignoreResponseSeverityOverrideSeverity string

const (
	CasbPostureFindingUnignoreResponseSeverityOverrideSeverityCritical CasbPostureFindingUnignoreResponseSeverityOverrideSeverity = "Critical"
	CasbPostureFindingUnignoreResponseSeverityOverrideSeverityHigh     CasbPostureFindingUnignoreResponseSeverityOverrideSeverity = "High"
	CasbPostureFindingUnignoreResponseSeverityOverrideSeverityMedium   CasbPostureFindingUnignoreResponseSeverityOverrideSeverity = "Medium"
	CasbPostureFindingUnignoreResponseSeverityOverrideSeverityLow      CasbPostureFindingUnignoreResponseSeverityOverrideSeverity = "Low"
)

func (r CasbPostureFindingUnignoreResponseSeverityOverrideSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingUnignoreResponseSeverityOverrideSeverityCritical, CasbPostureFindingUnignoreResponseSeverityOverrideSeverityHigh, CasbPostureFindingUnignoreResponseSeverityOverrideSeverityMedium, CasbPostureFindingUnignoreResponseSeverityOverrideSeverityLow:
		return true
	}
	return false
}

type CasbPostureFindingListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A cursor for pagination. Obtained from the `result_info.cursor` field of a
	// previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Direction to order results.
	Direction param.Field[CasbPostureFindingListParamsDirection] `query:"direction"`
	// A comma separated list of UUIDs identifying the finding type(s).
	FindingTypeIDs param.Field[string] `query:"finding_type_ids" format:"uuid"`
	// Filter for only the ignored findings. Set to false to only see "active" items
	Ignored param.Field[bool] `query:"ignored"`
	// Filter by an integration ID
	IntegrationID param.Field[string] `query:"integration_id" format:"uuid"`
	// Filter to view findings that occurred on or before the affliction date. Can be a
	// date-time in ISO 8601 format or an epoch timestamp.
	MaxAfflictionDate param.Field[time.Time] `query:"max_affliction_date" format:"date-time"`
	// Filter to view findings that occurred on or after the affliction date. Can be a
	// date-time in ISO 8601 format or an epoch timestamp.
	MinAfflictionDate param.Field[time.Time] `query:"min_affliction_date" format:"date-time"`
	// Filter by observation type of the finding
	Observation param.Field[CasbPostureFindingListParamsObservation] `query:"observation"`
	// Which field to use when ordering the findings.
	Order param.Field[CasbPostureFindingListParamsOrder] `query:"order"`
	// A page number within the paginated result set.
	Page param.Field[int64] `query:"page"`
	// Number of results to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter by product category of the finding
	Product param.Field[CasbPostureFindingListParamsProduct] `query:"product"`
	// A search term.
	Search param.Field[string] `query:"search"`
	// Filter by severity
	Severity param.Field[CasbPostureFindingListParamsSeverity] `query:"severity"`
	// Filter by type of the finding
	Type param.Field[CasbPostureFindingListParamsType] `query:"type"`
	// Filter by vendor
	Vendor param.Field[CasbPostureFindingListParamsVendor] `query:"vendor"`
}

// URLQuery serializes [CasbPostureFindingListParams]'s query parameters as
// `url.Values`.
func (r CasbPostureFindingListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type CasbPostureFindingListParamsDirection string

const (
	CasbPostureFindingListParamsDirectionAsc  CasbPostureFindingListParamsDirection = "asc"
	CasbPostureFindingListParamsDirectionDesc CasbPostureFindingListParamsDirection = "desc"
)

func (r CasbPostureFindingListParamsDirection) IsKnown() bool {
	switch r {
	case CasbPostureFindingListParamsDirectionAsc, CasbPostureFindingListParamsDirectionDesc:
		return true
	}
	return false
}

// Filter by observation type of the finding
type CasbPostureFindingListParamsObservation string

const (
	CasbPostureFindingListParamsObservationActivity CasbPostureFindingListParamsObservation = "Activity"
	CasbPostureFindingListParamsObservationInsight  CasbPostureFindingListParamsObservation = "Insight"
	CasbPostureFindingListParamsObservationIssue    CasbPostureFindingListParamsObservation = "Issue"
)

func (r CasbPostureFindingListParamsObservation) IsKnown() bool {
	switch r {
	case CasbPostureFindingListParamsObservationActivity, CasbPostureFindingListParamsObservationInsight, CasbPostureFindingListParamsObservationIssue:
		return true
	}
	return false
}

// Which field to use when ordering the findings.
type CasbPostureFindingListParamsOrder string

const (
	CasbPostureFindingListParamsOrderFindingName          CasbPostureFindingListParamsOrder = "finding.name"
	CasbPostureFindingListParamsOrderInstanceCount        CasbPostureFindingListParamsOrder = "instance_count"
	CasbPostureFindingListParamsOrderIntegrationName      CasbPostureFindingListParamsOrder = "integration.name"
	CasbPostureFindingListParamsOrderLatestAfflictionDate CasbPostureFindingListParamsOrder = "latest_affliction_date"
	CasbPostureFindingListParamsOrderSeverity             CasbPostureFindingListParamsOrder = "severity"
)

func (r CasbPostureFindingListParamsOrder) IsKnown() bool {
	switch r {
	case CasbPostureFindingListParamsOrderFindingName, CasbPostureFindingListParamsOrderInstanceCount, CasbPostureFindingListParamsOrderIntegrationName, CasbPostureFindingListParamsOrderLatestAfflictionDate, CasbPostureFindingListParamsOrderSeverity:
		return true
	}
	return false
}

// Filter by product category of the finding
type CasbPostureFindingListParamsProduct string

const (
	CasbPostureFindingListParamsProductCloud CasbPostureFindingListParamsProduct = "Cloud"
	CasbPostureFindingListParamsProductSaaS  CasbPostureFindingListParamsProduct = "Saas"
)

func (r CasbPostureFindingListParamsProduct) IsKnown() bool {
	switch r {
	case CasbPostureFindingListParamsProductCloud, CasbPostureFindingListParamsProductSaaS:
		return true
	}
	return false
}

// Filter by severity
type CasbPostureFindingListParamsSeverity string

const (
	CasbPostureFindingListParamsSeverityCritical CasbPostureFindingListParamsSeverity = "Critical"
	CasbPostureFindingListParamsSeverityHigh     CasbPostureFindingListParamsSeverity = "High"
	CasbPostureFindingListParamsSeverityMedium   CasbPostureFindingListParamsSeverity = "Medium"
	CasbPostureFindingListParamsSeverityLow      CasbPostureFindingListParamsSeverity = "Low"
)

func (r CasbPostureFindingListParamsSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingListParamsSeverityCritical, CasbPostureFindingListParamsSeverityHigh, CasbPostureFindingListParamsSeverityMedium, CasbPostureFindingListParamsSeverityLow:
		return true
	}
	return false
}

// Filter by type of the finding
type CasbPostureFindingListParamsType string

const (
	CasbPostureFindingListParamsTypeContent CasbPostureFindingListParamsType = "Content"
	CasbPostureFindingListParamsTypePosture CasbPostureFindingListParamsType = "Posture"
)

func (r CasbPostureFindingListParamsType) IsKnown() bool {
	switch r {
	case CasbPostureFindingListParamsTypeContent, CasbPostureFindingListParamsTypePosture:
		return true
	}
	return false
}

// Filter by vendor
type CasbPostureFindingListParamsVendor string

const (
	CasbPostureFindingListParamsVendorAnthropic           CasbPostureFindingListParamsVendor = "ANTHROPIC"
	CasbPostureFindingListParamsVendorAws                 CasbPostureFindingListParamsVendor = "AWS"
	CasbPostureFindingListParamsVendorBitbucket           CasbPostureFindingListParamsVendor = "BITBUCKET"
	CasbPostureFindingListParamsVendorBox                 CasbPostureFindingListParamsVendor = "BOX"
	CasbPostureFindingListParamsVendorConfluence          CasbPostureFindingListParamsVendor = "CONFLUENCE"
	CasbPostureFindingListParamsVendorDropbox             CasbPostureFindingListParamsVendor = "DROPBOX"
	CasbPostureFindingListParamsVendorGitHub              CasbPostureFindingListParamsVendor = "GITHUB"
	CasbPostureFindingListParamsVendorGoogleCloudPlatform CasbPostureFindingListParamsVendor = "GOOGLE_CLOUD_PLATFORM"
	CasbPostureFindingListParamsVendorGoogleWorkspace     CasbPostureFindingListParamsVendor = "GOOGLE_WORKSPACE"
	CasbPostureFindingListParamsVendorJira                CasbPostureFindingListParamsVendor = "JIRA"
	CasbPostureFindingListParamsVendorMicrosoft           CasbPostureFindingListParamsVendor = "MICROSOFT"
	CasbPostureFindingListParamsVendorMicrosoftInternal   CasbPostureFindingListParamsVendor = "MICROSOFT_INTERNAL"
	CasbPostureFindingListParamsVendorOpenAI              CasbPostureFindingListParamsVendor = "OPENAI"
	CasbPostureFindingListParamsVendorSalesforce          CasbPostureFindingListParamsVendor = "SALESFORCE"
	CasbPostureFindingListParamsVendorServicenow          CasbPostureFindingListParamsVendor = "SERVICENOW"
	CasbPostureFindingListParamsVendorSlack               CasbPostureFindingListParamsVendor = "SLACK"
)

func (r CasbPostureFindingListParamsVendor) IsKnown() bool {
	switch r {
	case CasbPostureFindingListParamsVendorAnthropic, CasbPostureFindingListParamsVendorAws, CasbPostureFindingListParamsVendorBitbucket, CasbPostureFindingListParamsVendorBox, CasbPostureFindingListParamsVendorConfluence, CasbPostureFindingListParamsVendorDropbox, CasbPostureFindingListParamsVendorGitHub, CasbPostureFindingListParamsVendorGoogleCloudPlatform, CasbPostureFindingListParamsVendorGoogleWorkspace, CasbPostureFindingListParamsVendorJira, CasbPostureFindingListParamsVendorMicrosoft, CasbPostureFindingListParamsVendorMicrosoftInternal, CasbPostureFindingListParamsVendorOpenAI, CasbPostureFindingListParamsVendorSalesforce, CasbPostureFindingListParamsVendorServicenow, CasbPostureFindingListParamsVendorSlack:
		return true
	}
	return false
}

type CasbPostureFindingExportParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter for only the ignored findings. Set to false to only see active items.
	Ignored param.Field[bool] `json:"ignored"`
	// Filter by multiple integration IDs.
	IntegrationID param.Field[[]string] `json:"integration_id" format:"uuid"`
	// Filter to view findings that occurred on or before the affliction date. Can be a
	// date-time in ISO 8601 format or an epoch timestamp.
	MaxAfflictionDate param.Field[time.Time] `json:"max_affliction_date" format:"date-time"`
	// Filter to view findings that occurred on or after the affliction date. Can be a
	// date-time in ISO 8601 format or an epoch timestamp.
	MinAfflictionDate param.Field[time.Time] `json:"min_affliction_date" format:"date-time"`
	// Which fields to use when ordering the findings.
	Orders param.Field[[]CasbPostureFindingExportParamsOrder] `json:"orders"`
	// Filter by finding's category product.
	Product param.Field[CasbPostureFindingExportParamsProduct] `json:"product"`
	// A search term.
	Search param.Field[string] `json:"search"`
	// Filter by severity levels.
	Severities param.Field[[]CasbPostureFindingExportParamsSeverity] `json:"severities"`
	// Filter by vendor types.
	Vendors param.Field[[]CasbPostureFindingExportParamsVendor] `json:"vendors"`
}

func (r CasbPostureFindingExportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Order specification for finding exports.
type CasbPostureFindingExportParamsOrder struct {
	// Sort direction.
	Direction param.Field[CasbPostureFindingExportParamsOrdersDirection] `json:"direction" api:"required"`
	// Which field to use when ordering the findings.
	Name param.Field[CasbPostureFindingExportParamsOrdersName] `json:"name" api:"required"`
}

func (r CasbPostureFindingExportParamsOrder) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort direction.
type CasbPostureFindingExportParamsOrdersDirection string

const (
	CasbPostureFindingExportParamsOrdersDirectionAsc  CasbPostureFindingExportParamsOrdersDirection = "asc"
	CasbPostureFindingExportParamsOrdersDirectionDesc CasbPostureFindingExportParamsOrdersDirection = "desc"
)

func (r CasbPostureFindingExportParamsOrdersDirection) IsKnown() bool {
	switch r {
	case CasbPostureFindingExportParamsOrdersDirectionAsc, CasbPostureFindingExportParamsOrdersDirectionDesc:
		return true
	}
	return false
}

// Which field to use when ordering the findings.
type CasbPostureFindingExportParamsOrdersName string

const (
	CasbPostureFindingExportParamsOrdersNameInstanceCount        CasbPostureFindingExportParamsOrdersName = "instance_count"
	CasbPostureFindingExportParamsOrdersNameFindingName          CasbPostureFindingExportParamsOrdersName = "finding.name"
	CasbPostureFindingExportParamsOrdersNameIntegrationName      CasbPostureFindingExportParamsOrdersName = "integration.name"
	CasbPostureFindingExportParamsOrdersNameLatestAfflictionDate CasbPostureFindingExportParamsOrdersName = "latest_affliction_date"
	CasbPostureFindingExportParamsOrdersNameSeverity             CasbPostureFindingExportParamsOrdersName = "severity"
)

func (r CasbPostureFindingExportParamsOrdersName) IsKnown() bool {
	switch r {
	case CasbPostureFindingExportParamsOrdersNameInstanceCount, CasbPostureFindingExportParamsOrdersNameFindingName, CasbPostureFindingExportParamsOrdersNameIntegrationName, CasbPostureFindingExportParamsOrdersNameLatestAfflictionDate, CasbPostureFindingExportParamsOrdersNameSeverity:
		return true
	}
	return false
}

// Filter by finding's category product.
type CasbPostureFindingExportParamsProduct string

const (
	CasbPostureFindingExportParamsProductSaaS  CasbPostureFindingExportParamsProduct = "SaaS"
	CasbPostureFindingExportParamsProductCloud CasbPostureFindingExportParamsProduct = "Cloud"
)

func (r CasbPostureFindingExportParamsProduct) IsKnown() bool {
	switch r {
	case CasbPostureFindingExportParamsProductSaaS, CasbPostureFindingExportParamsProductCloud:
		return true
	}
	return false
}

// The severity level for export filters.
type CasbPostureFindingExportParamsSeverity string

const (
	CasbPostureFindingExportParamsSeverityCritical CasbPostureFindingExportParamsSeverity = "CRITICAL"
	CasbPostureFindingExportParamsSeverityHigh     CasbPostureFindingExportParamsSeverity = "HIGH"
	CasbPostureFindingExportParamsSeverityMedium   CasbPostureFindingExportParamsSeverity = "MEDIUM"
	CasbPostureFindingExportParamsSeverityLow      CasbPostureFindingExportParamsSeverity = "LOW"
)

func (r CasbPostureFindingExportParamsSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingExportParamsSeverityCritical, CasbPostureFindingExportParamsSeverityHigh, CasbPostureFindingExportParamsSeverityMedium, CasbPostureFindingExportParamsSeverityLow:
		return true
	}
	return false
}

// Supported vendor types for integrations.
type CasbPostureFindingExportParamsVendor string

const (
	CasbPostureFindingExportParamsVendorAnthropic           CasbPostureFindingExportParamsVendor = "ANTHROPIC"
	CasbPostureFindingExportParamsVendorAws                 CasbPostureFindingExportParamsVendor = "AWS"
	CasbPostureFindingExportParamsVendorBitbucket           CasbPostureFindingExportParamsVendor = "BITBUCKET"
	CasbPostureFindingExportParamsVendorBox                 CasbPostureFindingExportParamsVendor = "BOX"
	CasbPostureFindingExportParamsVendorConfluence          CasbPostureFindingExportParamsVendor = "CONFLUENCE"
	CasbPostureFindingExportParamsVendorDropbox             CasbPostureFindingExportParamsVendor = "DROPBOX"
	CasbPostureFindingExportParamsVendorGitHub              CasbPostureFindingExportParamsVendor = "GITHUB"
	CasbPostureFindingExportParamsVendorGoogleCloudPlatform CasbPostureFindingExportParamsVendor = "GOOGLE_CLOUD_PLATFORM"
	CasbPostureFindingExportParamsVendorGoogleWorkspace     CasbPostureFindingExportParamsVendor = "GOOGLE_WORKSPACE"
	CasbPostureFindingExportParamsVendorJira                CasbPostureFindingExportParamsVendor = "JIRA"
	CasbPostureFindingExportParamsVendorMicrosoft           CasbPostureFindingExportParamsVendor = "MICROSOFT"
	CasbPostureFindingExportParamsVendorMicrosoftInternal   CasbPostureFindingExportParamsVendor = "MICROSOFT_INTERNAL"
	CasbPostureFindingExportParamsVendorOpenAI              CasbPostureFindingExportParamsVendor = "OPENAI"
	CasbPostureFindingExportParamsVendorSalesforce          CasbPostureFindingExportParamsVendor = "SALESFORCE"
	CasbPostureFindingExportParamsVendorServicenow          CasbPostureFindingExportParamsVendor = "SERVICENOW"
	CasbPostureFindingExportParamsVendorSlack               CasbPostureFindingExportParamsVendor = "SLACK"
)

func (r CasbPostureFindingExportParamsVendor) IsKnown() bool {
	switch r {
	case CasbPostureFindingExportParamsVendorAnthropic, CasbPostureFindingExportParamsVendorAws, CasbPostureFindingExportParamsVendorBitbucket, CasbPostureFindingExportParamsVendorBox, CasbPostureFindingExportParamsVendorConfluence, CasbPostureFindingExportParamsVendorDropbox, CasbPostureFindingExportParamsVendorGitHub, CasbPostureFindingExportParamsVendorGoogleCloudPlatform, CasbPostureFindingExportParamsVendorGoogleWorkspace, CasbPostureFindingExportParamsVendorJira, CasbPostureFindingExportParamsVendorMicrosoft, CasbPostureFindingExportParamsVendorMicrosoftInternal, CasbPostureFindingExportParamsVendorOpenAI, CasbPostureFindingExportParamsVendorSalesforce, CasbPostureFindingExportParamsVendorServicenow, CasbPostureFindingExportParamsVendorSlack:
		return true
	}
	return false
}

// Common response structure for all API endpoints.
type CasbPostureFindingExportResponseEnvelope struct {
	Errors   []CasbPostureFindingExportResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingExportResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Information about an export job.
	Result CasbPostureFindingExportResponse             `json:"result"`
	JSON   casbPostureFindingExportResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingExportResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPostureFindingExportResponseEnvelope]
type casbPostureFindingExportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingExportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingExportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingExportResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                               `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingExportResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingExportResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingExportResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CasbPostureFindingExportResponseEnvelopeErrors]
type casbPostureFindingExportResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingExportResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingExportResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingExportResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                   `json:"pointer"`
	JSON    casbPostureFindingExportResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingExportResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CasbPostureFindingExportResponseEnvelopeErrorsSource]
type casbPostureFindingExportResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingExportResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingExportResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingExportResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                 `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingExportResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingExportResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingExportResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CasbPostureFindingExportResponseEnvelopeMessages]
type casbPostureFindingExportResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingExportResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingExportResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingExportResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                     `json:"pointer"`
	JSON    casbPostureFindingExportResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingExportResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CasbPostureFindingExportResponseEnvelopeMessagesSource]
type casbPostureFindingExportResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingExportResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingExportResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

// Common response structure for all API endpoints.
type CasbPostureFindingGetResponseEnvelope struct {
	Errors   []CasbPostureFindingGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Aggregated finding information with counts and metadata. This is optimized for
	// list API queries and represents a finding along with its instance statistics.
	Result CasbPostureFindingGetResponse             `json:"result"`
	JSON   casbPostureFindingGetResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPostureFindingGetResponseEnvelope]
type casbPostureFindingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingGetResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                            `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CasbPostureFindingGetResponseEnvelopeErrors]
type casbPostureFindingGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingGetResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                `json:"pointer"`
	JSON    casbPostureFindingGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CasbPostureFindingGetResponseEnvelopeErrorsSource]
type casbPostureFindingGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingGetResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                              `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CasbPostureFindingGetResponseEnvelopeMessages]
type casbPostureFindingGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingGetResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                  `json:"pointer"`
	JSON    casbPostureFindingGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CasbPostureFindingGetResponseEnvelopeMessagesSource]
type casbPostureFindingGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingIgnoreParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A list of finding IDs to pass along.
	Checks param.Field[[]string] `json:"checks" api:"required"`
}

func (r CasbPostureFindingIgnoreParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Common response structure for all API endpoints.
type CasbPostureFindingIgnoreResponseEnvelope struct {
	Errors   []CasbPostureFindingIgnoreResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingIgnoreResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Aggregated finding information with counts and metadata. This is optimized for
	// list API queries and represents a finding along with its instance statistics.
	Result CasbPostureFindingIgnoreResponse             `json:"result"`
	JSON   casbPostureFindingIgnoreResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingIgnoreResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPostureFindingIgnoreResponseEnvelope]
type casbPostureFindingIgnoreResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingIgnoreResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                               `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingIgnoreResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingIgnoreResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingIgnoreResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CasbPostureFindingIgnoreResponseEnvelopeErrors]
type casbPostureFindingIgnoreResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingIgnoreResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                   `json:"pointer"`
	JSON    casbPostureFindingIgnoreResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingIgnoreResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CasbPostureFindingIgnoreResponseEnvelopeErrorsSource]
type casbPostureFindingIgnoreResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingIgnoreResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                 `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingIgnoreResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingIgnoreResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingIgnoreResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CasbPostureFindingIgnoreResponseEnvelopeMessages]
type casbPostureFindingIgnoreResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingIgnoreResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                     `json:"pointer"`
	JSON    casbPostureFindingIgnoreResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingIgnoreResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CasbPostureFindingIgnoreResponseEnvelopeMessagesSource]
type casbPostureFindingIgnoreResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingIgnoreResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingIgnoreResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingResetSeverityParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

// Common response structure for all API endpoints.
type CasbPostureFindingResetSeverityResponseEnvelope struct {
	Errors   []CasbPostureFindingResetSeverityResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingResetSeverityResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Aggregated finding information with counts and metadata. This is optimized for
	// list API queries and represents a finding along with its instance statistics.
	Result CasbPostureFindingResetSeverityResponse             `json:"result"`
	JSON   casbPostureFindingResetSeverityResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingResetSeverityResponseEnvelopeJSON contains the JSON metadata
// for the struct [CasbPostureFindingResetSeverityResponseEnvelope]
type casbPostureFindingResetSeverityResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingResetSeverityResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                      `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingResetSeverityResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingResetSeverityResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingResetSeverityResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CasbPostureFindingResetSeverityResponseEnvelopeErrors]
type casbPostureFindingResetSeverityResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingResetSeverityResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                          `json:"pointer"`
	JSON    casbPostureFindingResetSeverityResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingResetSeverityResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingResetSeverityResponseEnvelopeErrorsSource]
type casbPostureFindingResetSeverityResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingResetSeverityResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                        `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingResetSeverityResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingResetSeverityResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingResetSeverityResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingResetSeverityResponseEnvelopeMessages]
type casbPostureFindingResetSeverityResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingResetSeverityResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                            `json:"pointer"`
	JSON    casbPostureFindingResetSeverityResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingResetSeverityResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingResetSeverityResponseEnvelopeMessagesSource]
type casbPostureFindingResetSeverityResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingResetSeverityResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingResetSeverityResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTuneSeverityParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The numeric severity value to apply to the finding.
	NewSeverity param.Field[CasbPostureFindingTuneSeverityParamsNewSeverity] `json:"new_severity" api:"required"`
}

func (r CasbPostureFindingTuneSeverityParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The numeric severity value to apply to the finding.
type CasbPostureFindingTuneSeverityParamsNewSeverity int64

const (
	CasbPostureFindingTuneSeverityParamsNewSeverity1 CasbPostureFindingTuneSeverityParamsNewSeverity = 1
	CasbPostureFindingTuneSeverityParamsNewSeverity2 CasbPostureFindingTuneSeverityParamsNewSeverity = 2
	CasbPostureFindingTuneSeverityParamsNewSeverity3 CasbPostureFindingTuneSeverityParamsNewSeverity = 3
	CasbPostureFindingTuneSeverityParamsNewSeverity4 CasbPostureFindingTuneSeverityParamsNewSeverity = 4
)

func (r CasbPostureFindingTuneSeverityParamsNewSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingTuneSeverityParamsNewSeverity1, CasbPostureFindingTuneSeverityParamsNewSeverity2, CasbPostureFindingTuneSeverityParamsNewSeverity3, CasbPostureFindingTuneSeverityParamsNewSeverity4:
		return true
	}
	return false
}

// Common response structure for all API endpoints.
type CasbPostureFindingTuneSeverityResponseEnvelope struct {
	Errors   []CasbPostureFindingTuneSeverityResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingTuneSeverityResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Aggregated finding information with counts and metadata. This is optimized for
	// list API queries and represents a finding along with its instance statistics.
	Result CasbPostureFindingTuneSeverityResponse             `json:"result"`
	JSON   casbPostureFindingTuneSeverityResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingTuneSeverityResponseEnvelopeJSON contains the JSON metadata
// for the struct [CasbPostureFindingTuneSeverityResponseEnvelope]
type casbPostureFindingTuneSeverityResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTuneSeverityResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                     `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingTuneSeverityResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingTuneSeverityResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingTuneSeverityResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CasbPostureFindingTuneSeverityResponseEnvelopeErrors]
type casbPostureFindingTuneSeverityResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTuneSeverityResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                         `json:"pointer"`
	JSON    casbPostureFindingTuneSeverityResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingTuneSeverityResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingTuneSeverityResponseEnvelopeErrorsSource]
type casbPostureFindingTuneSeverityResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTuneSeverityResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                       `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingTuneSeverityResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingTuneSeverityResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingTuneSeverityResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CasbPostureFindingTuneSeverityResponseEnvelopeMessages]
type casbPostureFindingTuneSeverityResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTuneSeverityResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                           `json:"pointer"`
	JSON    casbPostureFindingTuneSeverityResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingTuneSeverityResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingTuneSeverityResponseEnvelopeMessagesSource]
type casbPostureFindingTuneSeverityResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTuneSeverityResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTuneSeverityResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingUnignoreParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A list of finding IDs to pass along.
	Checks param.Field[[]string] `json:"checks" api:"required"`
}

func (r CasbPostureFindingUnignoreParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Common response structure for all API endpoints.
type CasbPostureFindingUnignoreResponseEnvelope struct {
	Errors   []CasbPostureFindingUnignoreResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingUnignoreResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Aggregated finding information with counts and metadata. This is optimized for
	// list API queries and represents a finding along with its instance statistics.
	Result CasbPostureFindingUnignoreResponse             `json:"result"`
	JSON   casbPostureFindingUnignoreResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingUnignoreResponseEnvelopeJSON contains the JSON metadata for
// the struct [CasbPostureFindingUnignoreResponseEnvelope]
type casbPostureFindingUnignoreResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingUnignoreResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                 `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingUnignoreResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingUnignoreResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingUnignoreResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CasbPostureFindingUnignoreResponseEnvelopeErrors]
type casbPostureFindingUnignoreResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingUnignoreResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                     `json:"pointer"`
	JSON    casbPostureFindingUnignoreResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingUnignoreResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CasbPostureFindingUnignoreResponseEnvelopeErrorsSource]
type casbPostureFindingUnignoreResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingUnignoreResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                   `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingUnignoreResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingUnignoreResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingUnignoreResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CasbPostureFindingUnignoreResponseEnvelopeMessages]
type casbPostureFindingUnignoreResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingUnignoreResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                       `json:"pointer"`
	JSON    casbPostureFindingUnignoreResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingUnignoreResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingUnignoreResponseEnvelopeMessagesSource]
type casbPostureFindingUnignoreResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingUnignoreResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingUnignoreResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
