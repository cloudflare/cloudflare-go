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

// CasbPostureContentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbPostureContentService] method instead.
type CasbPostureContentService struct {
	Options []option.RequestOption
}

// NewCasbPostureContentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCasbPostureContentService(opts ...option.RequestOption) (r *CasbPostureContentService) {
	r = &CasbPostureContentService{}
	r.Options = opts
	return
}

// List DLP content findings
func (r *CasbPostureContentService) List(ctx context.Context, params CasbPostureContentListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[CasbPostureContentListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/content", params.AccountID)
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

// List DLP content findings
func (r *CasbPostureContentService) ListAutoPaging(ctx context.Context, params CasbPostureContentListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[CasbPostureContentListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Creates a CSV export for content and accepts optional filters in the payload.
func (r *CasbPostureContentService) Export(ctx context.Context, params CasbPostureContentExportParams, opts ...option.RequestOption) (res *CasbPostureContentExportResponse, err error) {
	var env CasbPostureContentExportResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/content/export", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Content asset with DLP information.
type CasbPostureContentListResponse struct {
	// Unique identifier for the asset.
	AssetID string `json:"asset_id" api:"required" format:"uuid"`
	// Name of the asset.
	AssetName string `json:"asset_name" api:"required"`
	// DLP context information for this asset.
	DLPContexts []CasbPostureContentListResponseDLPContext `json:"dlp_contexts" api:"required"`
	// Number of DLP profiles that flagged this asset.
	DLPProfileCount int64 `json:"dlp_profile_count" api:"required"`
	// IDs of DLP profiles that flagged this asset.
	DLPProfileIDs []string `json:"dlp_profile_ids" api:"required" format:"uuid"`
	// Summary information about an integration.
	Integration CasbPostureContentListResponseIntegration `json:"integration" api:"required"`
	// Most recent date this asset was flagged.
	LatestAfflictionDate time.Time                          `json:"latest_affliction_date" api:"required" format:"date-time"`
	JSON                 casbPostureContentListResponseJSON `json:"-"`
}

// casbPostureContentListResponseJSON contains the JSON metadata for the struct
// [CasbPostureContentListResponse]
type casbPostureContentListResponseJSON struct {
	AssetID              apijson.Field
	AssetName            apijson.Field
	DLPContexts          apijson.Field
	DLPProfileCount      apijson.Field
	DLPProfileIDs        apijson.Field
	Integration          apijson.Field
	LatestAfflictionDate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CasbPostureContentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureContentListResponseJSON) RawJSON() string {
	return r.raw
}

// DLP context information for a finding.
type CasbPostureContentListResponseDLPContext struct {
	// When the DLP context was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// DLP Entry IDs.
	EntryIDs []string `json:"entry_ids" api:"required" format:"uuid"`
	// DLP Profile ID.
	ProfileID string `json:"profile_id" api:"required" format:"uuid"`
	// When the DLP context was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Unique identifier for the DLP context.
	ID string `json:"id" format:"uuid"`
	// When the DLP context was deleted.
	Deleted time.Time `json:"deleted" api:"nullable" format:"date-time"`
	// DLP Right Boundary of match context.
	MatchContextMaxExtent int64 `json:"match_context_max_extent" api:"nullable"`
	// DLP Left Boundary of match context.
	MatchContextMinExtent int64 `json:"match_context_min_extent" api:"nullable"`
	// DLP Match context payload that matched the profile in question.
	MatchContextPayload map[string]interface{}                       `json:"match_context_payload" api:"nullable"`
	JSON                casbPostureContentListResponseDLPContextJSON `json:"-"`
}

// casbPostureContentListResponseDLPContextJSON contains the JSON metadata for the
// struct [CasbPostureContentListResponseDLPContext]
type casbPostureContentListResponseDLPContextJSON struct {
	Created               apijson.Field
	EntryIDs              apijson.Field
	ProfileID             apijson.Field
	Updated               apijson.Field
	ID                    apijson.Field
	Deleted               apijson.Field
	MatchContextMaxExtent apijson.Field
	MatchContextMinExtent apijson.Field
	MatchContextPayload   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CasbPostureContentListResponseDLPContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureContentListResponseDLPContextJSON) RawJSON() string {
	return r.raw
}

// Summary information about an integration.
type CasbPostureContentListResponseIntegration struct {
	// When entity was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// When were the integration credentials last updated.
	LastHydrated time.Time `json:"last_hydrated" api:"required" format:"date-time"`
	// Name of the integration.
	Name string `json:"name" api:"required"`
	// The vendor-specific permissions associated with the integration.
	Permissions []string `json:"permissions" api:"required"`
	// Policy configuration for an integration.
	Policy CasbPostureContentListResponseIntegrationPolicy `json:"policy" api:"required"`
	// Current status of the integration.
	Status string `json:"status" api:"required"`
	// Last entity was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Whether the integrations permissions can be updated.
	Upgradable bool `json:"upgradable" api:"required"`
	// Information about a vendor/service provider.
	Vendor CasbPostureContentListResponseIntegrationVendor `json:"vendor" api:"required"`
	// Zero Trust products associated with this integration.
	ZtEnrollments []CasbPostureContentListResponseIntegrationZtEnrollment `json:"zt_enrollments" api:"required"`
	// Integration ID.
	ID string `json:"id" format:"uuid"`
	// Health status of integration credentials.
	CredentialHealthStatus CasbPostureContentListResponseIntegrationCredentialHealthStatus `json:"credential_health_status"`
	// The date and time when the integration credentials will expire.
	CredentialsExpiry time.Time `json:"credentials_expiry" api:"nullable" format:"date-time"`
	// Whether the given integration is paused by the user.
	IsPaused bool `json:"is_paused"`
	// UI State as to whether a potential permissions upgrade has been dismissed.
	UpgradeDismissed bool                                          `json:"upgrade_dismissed"`
	JSON             casbPostureContentListResponseIntegrationJSON `json:"-"`
}

// casbPostureContentListResponseIntegrationJSON contains the JSON metadata for the
// struct [CasbPostureContentListResponseIntegration]
type casbPostureContentListResponseIntegrationJSON struct {
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

func (r *CasbPostureContentListResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureContentListResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

// Policy configuration for an integration.
type CasbPostureContentListResponseIntegrationPolicy struct {
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
	JSON        casbPostureContentListResponseIntegrationPolicyJSON `json:"-"`
}

// casbPostureContentListResponseIntegrationPolicyJSON contains the JSON metadata
// for the struct [CasbPostureContentListResponseIntegrationPolicy]
type casbPostureContentListResponseIntegrationPolicyJSON struct {
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

func (r *CasbPostureContentListResponseIntegrationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureContentListResponseIntegrationPolicyJSON) RawJSON() string {
	return r.raw
}

// Information about a vendor/service provider.
type CasbPostureContentListResponseIntegrationVendor struct {
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
	JSON     casbPostureContentListResponseIntegrationVendorJSON `json:"-"`
}

// casbPostureContentListResponseIntegrationVendorJSON contains the JSON metadata
// for the struct [CasbPostureContentListResponseIntegrationVendor]
type casbPostureContentListResponseIntegrationVendorJSON struct {
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

func (r *CasbPostureContentListResponseIntegrationVendor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureContentListResponseIntegrationVendorJSON) RawJSON() string {
	return r.raw
}

// Information about a Zero Trust product integration.
type CasbPostureContentListResponseIntegrationZtEnrollment struct {
	// The internal identifier of the Zero Trust Product.
	ID string `json:"id"`
	// Brief description of the Zero Trust Product.
	Description string `json:"description"`
	// The verbose name of the Zero Trust Product.
	DisplayName string `json:"display_name"`
	// Flag to enable/disable access to the listed integration from the corresponding
	// Cloudflare product.
	Enabled bool                                                      `json:"enabled"`
	JSON    casbPostureContentListResponseIntegrationZtEnrollmentJSON `json:"-"`
}

// casbPostureContentListResponseIntegrationZtEnrollmentJSON contains the JSON
// metadata for the struct [CasbPostureContentListResponseIntegrationZtEnrollment]
type casbPostureContentListResponseIntegrationZtEnrollmentJSON struct {
	ID          apijson.Field
	Description apijson.Field
	DisplayName apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureContentListResponseIntegrationZtEnrollment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureContentListResponseIntegrationZtEnrollmentJSON) RawJSON() string {
	return r.raw
}

// Health status of integration credentials.
type CasbPostureContentListResponseIntegrationCredentialHealthStatus string

const (
	CasbPostureContentListResponseIntegrationCredentialHealthStatusInitializing CasbPostureContentListResponseIntegrationCredentialHealthStatus = "Initializing"
	CasbPostureContentListResponseIntegrationCredentialHealthStatusHealthy      CasbPostureContentListResponseIntegrationCredentialHealthStatus = "Healthy"
	CasbPostureContentListResponseIntegrationCredentialHealthStatusUnhealthy    CasbPostureContentListResponseIntegrationCredentialHealthStatus = "Unhealthy"
)

func (r CasbPostureContentListResponseIntegrationCredentialHealthStatus) IsKnown() bool {
	switch r {
	case CasbPostureContentListResponseIntegrationCredentialHealthStatusInitializing, CasbPostureContentListResponseIntegrationCredentialHealthStatusHealthy, CasbPostureContentListResponseIntegrationCredentialHealthStatusUnhealthy:
		return true
	}
	return false
}

// Information about an export job.
type CasbPostureContentExportResponse struct {
	// Unique identifier for the export job.
	ID string `json:"id" api:"required" format:"uuid"`
	// Status of an export job.
	Status CasbPostureContentExportResponseStatus `json:"status" api:"required"`
	// Type of export job.
	Type CasbPostureContentExportResponseType `json:"type" api:"required"`
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
	JSON     casbPostureContentExportResponseJSON `json:"-"`
}

// casbPostureContentExportResponseJSON contains the JSON metadata for the struct
// [CasbPostureContentExportResponse]
type casbPostureContentExportResponseJSON struct {
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

func (r *CasbPostureContentExportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureContentExportResponseJSON) RawJSON() string {
	return r.raw
}

// Status of an export job.
type CasbPostureContentExportResponseStatus string

const (
	CasbPostureContentExportResponseStatusPending     CasbPostureContentExportResponseStatus = "Pending"
	CasbPostureContentExportResponseStatusSuccess     CasbPostureContentExportResponseStatus = "Success"
	CasbPostureContentExportResponseStatusFailure     CasbPostureContentExportResponseStatus = "Failure"
	CasbPostureContentExportResponseStatusRescheduled CasbPostureContentExportResponseStatus = "Rescheduled"
	CasbPostureContentExportResponseStatusInProgress  CasbPostureContentExportResponseStatus = "In-Progress"
)

func (r CasbPostureContentExportResponseStatus) IsKnown() bool {
	switch r {
	case CasbPostureContentExportResponseStatusPending, CasbPostureContentExportResponseStatusSuccess, CasbPostureContentExportResponseStatusFailure, CasbPostureContentExportResponseStatusRescheduled, CasbPostureContentExportResponseStatusInProgress:
		return true
	}
	return false
}

// Type of export job.
type CasbPostureContentExportResponseType string

const (
	CasbPostureContentExportResponseTypeFinding         CasbPostureContentExportResponseType = "finding"
	CasbPostureContentExportResponseTypeFindingInstance CasbPostureContentExportResponseType = "findingInstance"
	CasbPostureContentExportResponseTypeContent         CasbPostureContentExportResponseType = "content"
	CasbPostureContentExportResponseTypeRemediationJob  CasbPostureContentExportResponseType = "remediationJob"
)

func (r CasbPostureContentExportResponseType) IsKnown() bool {
	switch r {
	case CasbPostureContentExportResponseTypeFinding, CasbPostureContentExportResponseTypeFindingInstance, CasbPostureContentExportResponseTypeContent, CasbPostureContentExportResponseTypeRemediationJob:
		return true
	}
	return false
}

type CasbPostureContentListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Direction to order results.
	Direction param.Field[CasbPostureContentListParamsDirection] `query:"direction"`
	// Filter by an DLP profile ID
	DLPProfileID param.Field[string] `query:"dlp_profile_id" format:"uuid"`
	// Filter by an integration ID
	IntegrationID param.Field[string] `query:"integration_id" format:"uuid"`
	// Filter to view findings that occurred on or before the affliction date. Can be a
	// date-time in ISO 8601 format or an epoch timestamp.
	MaxAfflictionDate param.Field[time.Time] `query:"max_affliction_date" format:"date-time"`
	// Filter to view findings that occurred on or after the affliction date. Can be a
	// date-time in ISO 8601 format or an epoch timestamp.
	MinAfflictionDate param.Field[time.Time] `query:"min_affliction_date" format:"date-time"`
	// Which field to use when ordering content assets.
	Order param.Field[CasbPostureContentListParamsOrder] `query:"order"`
	// A page number within the paginated result set.
	Page param.Field[int64] `query:"page"`
	// Number of results to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// A search term.
	Search param.Field[string] `query:"search"`
	// Filter by vendor
	Vendor param.Field[CasbPostureContentListParamsVendor] `query:"vendor"`
}

// URLQuery serializes [CasbPostureContentListParams]'s query parameters as
// `url.Values`.
func (r CasbPostureContentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type CasbPostureContentListParamsDirection string

const (
	CasbPostureContentListParamsDirectionAsc  CasbPostureContentListParamsDirection = "asc"
	CasbPostureContentListParamsDirectionDesc CasbPostureContentListParamsDirection = "desc"
)

func (r CasbPostureContentListParamsDirection) IsKnown() bool {
	switch r {
	case CasbPostureContentListParamsDirectionAsc, CasbPostureContentListParamsDirectionDesc:
		return true
	}
	return false
}

// Which field to use when ordering content assets.
type CasbPostureContentListParamsOrder string

const (
	CasbPostureContentListParamsOrderAssetName            CasbPostureContentListParamsOrder = "asset_name"
	CasbPostureContentListParamsOrderDLPProfileCount      CasbPostureContentListParamsOrder = "dlp_profile_count"
	CasbPostureContentListParamsOrderIntegrationName      CasbPostureContentListParamsOrder = "integration_name"
	CasbPostureContentListParamsOrderLatestAfflictionDate CasbPostureContentListParamsOrder = "latest_affliction_date"
)

func (r CasbPostureContentListParamsOrder) IsKnown() bool {
	switch r {
	case CasbPostureContentListParamsOrderAssetName, CasbPostureContentListParamsOrderDLPProfileCount, CasbPostureContentListParamsOrderIntegrationName, CasbPostureContentListParamsOrderLatestAfflictionDate:
		return true
	}
	return false
}

// Filter by vendor
type CasbPostureContentListParamsVendor string

const (
	CasbPostureContentListParamsVendorAnthropic           CasbPostureContentListParamsVendor = "ANTHROPIC"
	CasbPostureContentListParamsVendorAws                 CasbPostureContentListParamsVendor = "AWS"
	CasbPostureContentListParamsVendorBitbucket           CasbPostureContentListParamsVendor = "BITBUCKET"
	CasbPostureContentListParamsVendorBox                 CasbPostureContentListParamsVendor = "BOX"
	CasbPostureContentListParamsVendorConfluence          CasbPostureContentListParamsVendor = "CONFLUENCE"
	CasbPostureContentListParamsVendorDropbox             CasbPostureContentListParamsVendor = "DROPBOX"
	CasbPostureContentListParamsVendorGitHub              CasbPostureContentListParamsVendor = "GITHUB"
	CasbPostureContentListParamsVendorGoogleCloudPlatform CasbPostureContentListParamsVendor = "GOOGLE_CLOUD_PLATFORM"
	CasbPostureContentListParamsVendorGoogleWorkspace     CasbPostureContentListParamsVendor = "GOOGLE_WORKSPACE"
	CasbPostureContentListParamsVendorJira                CasbPostureContentListParamsVendor = "JIRA"
	CasbPostureContentListParamsVendorMicrosoft           CasbPostureContentListParamsVendor = "MICROSOFT"
	CasbPostureContentListParamsVendorMicrosoftInternal   CasbPostureContentListParamsVendor = "MICROSOFT_INTERNAL"
	CasbPostureContentListParamsVendorOpenAI              CasbPostureContentListParamsVendor = "OPENAI"
	CasbPostureContentListParamsVendorSalesforce          CasbPostureContentListParamsVendor = "SALESFORCE"
	CasbPostureContentListParamsVendorServicenow          CasbPostureContentListParamsVendor = "SERVICENOW"
	CasbPostureContentListParamsVendorSlack               CasbPostureContentListParamsVendor = "SLACK"
)

func (r CasbPostureContentListParamsVendor) IsKnown() bool {
	switch r {
	case CasbPostureContentListParamsVendorAnthropic, CasbPostureContentListParamsVendorAws, CasbPostureContentListParamsVendorBitbucket, CasbPostureContentListParamsVendorBox, CasbPostureContentListParamsVendorConfluence, CasbPostureContentListParamsVendorDropbox, CasbPostureContentListParamsVendorGitHub, CasbPostureContentListParamsVendorGoogleCloudPlatform, CasbPostureContentListParamsVendorGoogleWorkspace, CasbPostureContentListParamsVendorJira, CasbPostureContentListParamsVendorMicrosoft, CasbPostureContentListParamsVendorMicrosoftInternal, CasbPostureContentListParamsVendorOpenAI, CasbPostureContentListParamsVendorSalesforce, CasbPostureContentListParamsVendorServicenow, CasbPostureContentListParamsVendorSlack:
		return true
	}
	return false
}

type CasbPostureContentExportParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// DLP profile metadata for the export.
	DLPProfileInformation param.Field[[]CasbPostureContentExportParamsDLPProfileInformation] `json:"dlp_profile_information" api:"required"`
	// Filter by DLP profile IDs.
	DLPProfileID param.Field[[]string] `json:"dlp_profile_id" format:"uuid"`
	// Filter by integration IDs.
	IntegrationID param.Field[[]string] `json:"integration_id" format:"uuid"`
	// Filter to view content flagged on or before this date.
	MaxAfflictionDate param.Field[time.Time] `json:"max_affliction_date" format:"date-time"`
	// Filter to view content flagged on or after this date.
	MinAfflictionDate param.Field[time.Time] `json:"min_affliction_date" format:"date-time"`
	// Ordering specifications for the export.
	Orders param.Field[[]CasbPostureContentExportParamsOrder] `json:"orders"`
	// Search term to filter content.
	Search param.Field[string] `json:"search"`
	// Filter by vendor types.
	Vendors param.Field[[]CasbPostureContentExportParamsVendor] `json:"vendors"`
}

func (r CasbPostureContentExportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP profile configuration.
type CasbPostureContentExportParamsDLPProfileInformation struct {
	// Unique identifier for the DLP profile.
	ID param.Field[string] `json:"id" api:"required" format:"uuid"`
	// Entries contained within this DLP profile.
	Entries param.Field[[]CasbPostureContentExportParamsDLPProfileInformationEntry] `json:"entries" api:"required"`
	// Name of the DLP profile.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r CasbPostureContentExportParamsDLPProfileInformation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Entry within a DLP profile.
type CasbPostureContentExportParamsDLPProfileInformationEntry struct {
	// Unique identifier for the DLP profile entry.
	ID param.Field[string] `json:"id" api:"required" format:"uuid"`
	// Name of the DLP profile entry.
	Name param.Field[string] `json:"name" api:"required"`
	// ID of the parent DLP profile.
	ProfileID param.Field[string] `json:"profile_id" api:"required" format:"uuid"`
}

func (r CasbPostureContentExportParamsDLPProfileInformationEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Generic ordering specification.
type CasbPostureContentExportParamsOrder struct {
	// Sort direction.
	Direction param.Field[CasbPostureContentExportParamsOrdersDirection] `json:"direction" api:"required"`
	// Content-specific field names for ordering.
	Name param.Field[CasbPostureContentExportParamsOrdersName] `json:"name" api:"required"`
}

func (r CasbPostureContentExportParamsOrder) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort direction.
type CasbPostureContentExportParamsOrdersDirection string

const (
	CasbPostureContentExportParamsOrdersDirectionAsc  CasbPostureContentExportParamsOrdersDirection = "asc"
	CasbPostureContentExportParamsOrdersDirectionDesc CasbPostureContentExportParamsOrdersDirection = "desc"
)

func (r CasbPostureContentExportParamsOrdersDirection) IsKnown() bool {
	switch r {
	case CasbPostureContentExportParamsOrdersDirectionAsc, CasbPostureContentExportParamsOrdersDirectionDesc:
		return true
	}
	return false
}

// Content-specific field names for ordering.
type CasbPostureContentExportParamsOrdersName string

const (
	CasbPostureContentExportParamsOrdersNameAssetName            CasbPostureContentExportParamsOrdersName = "asset_name"
	CasbPostureContentExportParamsOrdersNameDLPProfileCount      CasbPostureContentExportParamsOrdersName = "dlp_profile_count"
	CasbPostureContentExportParamsOrdersNameIntegrationName      CasbPostureContentExportParamsOrdersName = "integration_name"
	CasbPostureContentExportParamsOrdersNameLatestAfflictionDate CasbPostureContentExportParamsOrdersName = "latest_affliction_date"
)

func (r CasbPostureContentExportParamsOrdersName) IsKnown() bool {
	switch r {
	case CasbPostureContentExportParamsOrdersNameAssetName, CasbPostureContentExportParamsOrdersNameDLPProfileCount, CasbPostureContentExportParamsOrdersNameIntegrationName, CasbPostureContentExportParamsOrdersNameLatestAfflictionDate:
		return true
	}
	return false
}

// Supported vendor types for integrations.
type CasbPostureContentExportParamsVendor string

const (
	CasbPostureContentExportParamsVendorAnthropic           CasbPostureContentExportParamsVendor = "ANTHROPIC"
	CasbPostureContentExportParamsVendorAws                 CasbPostureContentExportParamsVendor = "AWS"
	CasbPostureContentExportParamsVendorBitbucket           CasbPostureContentExportParamsVendor = "BITBUCKET"
	CasbPostureContentExportParamsVendorBox                 CasbPostureContentExportParamsVendor = "BOX"
	CasbPostureContentExportParamsVendorConfluence          CasbPostureContentExportParamsVendor = "CONFLUENCE"
	CasbPostureContentExportParamsVendorDropbox             CasbPostureContentExportParamsVendor = "DROPBOX"
	CasbPostureContentExportParamsVendorGitHub              CasbPostureContentExportParamsVendor = "GITHUB"
	CasbPostureContentExportParamsVendorGoogleCloudPlatform CasbPostureContentExportParamsVendor = "GOOGLE_CLOUD_PLATFORM"
	CasbPostureContentExportParamsVendorGoogleWorkspace     CasbPostureContentExportParamsVendor = "GOOGLE_WORKSPACE"
	CasbPostureContentExportParamsVendorJira                CasbPostureContentExportParamsVendor = "JIRA"
	CasbPostureContentExportParamsVendorMicrosoft           CasbPostureContentExportParamsVendor = "MICROSOFT"
	CasbPostureContentExportParamsVendorMicrosoftInternal   CasbPostureContentExportParamsVendor = "MICROSOFT_INTERNAL"
	CasbPostureContentExportParamsVendorOpenAI              CasbPostureContentExportParamsVendor = "OPENAI"
	CasbPostureContentExportParamsVendorSalesforce          CasbPostureContentExportParamsVendor = "SALESFORCE"
	CasbPostureContentExportParamsVendorServicenow          CasbPostureContentExportParamsVendor = "SERVICENOW"
	CasbPostureContentExportParamsVendorSlack               CasbPostureContentExportParamsVendor = "SLACK"
)

func (r CasbPostureContentExportParamsVendor) IsKnown() bool {
	switch r {
	case CasbPostureContentExportParamsVendorAnthropic, CasbPostureContentExportParamsVendorAws, CasbPostureContentExportParamsVendorBitbucket, CasbPostureContentExportParamsVendorBox, CasbPostureContentExportParamsVendorConfluence, CasbPostureContentExportParamsVendorDropbox, CasbPostureContentExportParamsVendorGitHub, CasbPostureContentExportParamsVendorGoogleCloudPlatform, CasbPostureContentExportParamsVendorGoogleWorkspace, CasbPostureContentExportParamsVendorJira, CasbPostureContentExportParamsVendorMicrosoft, CasbPostureContentExportParamsVendorMicrosoftInternal, CasbPostureContentExportParamsVendorOpenAI, CasbPostureContentExportParamsVendorSalesforce, CasbPostureContentExportParamsVendorServicenow, CasbPostureContentExportParamsVendorSlack:
		return true
	}
	return false
}

// Common response structure for all API endpoints.
type CasbPostureContentExportResponseEnvelope struct {
	Errors   []CasbPostureContentExportResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureContentExportResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Information about an export job.
	Result CasbPostureContentExportResponse             `json:"result"`
	JSON   casbPostureContentExportResponseEnvelopeJSON `json:"-"`
}

// casbPostureContentExportResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPostureContentExportResponseEnvelope]
type casbPostureContentExportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureContentExportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureContentExportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureContentExportResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                               `json:"documentation_url" format:"uri"`
	Source           CasbPostureContentExportResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureContentExportResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureContentExportResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CasbPostureContentExportResponseEnvelopeErrors]
type casbPostureContentExportResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureContentExportResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureContentExportResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureContentExportResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                   `json:"pointer"`
	JSON    casbPostureContentExportResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureContentExportResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CasbPostureContentExportResponseEnvelopeErrorsSource]
type casbPostureContentExportResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureContentExportResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureContentExportResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureContentExportResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                 `json:"documentation_url" format:"uri"`
	Source           CasbPostureContentExportResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureContentExportResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureContentExportResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CasbPostureContentExportResponseEnvelopeMessages]
type casbPostureContentExportResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureContentExportResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureContentExportResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureContentExportResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                     `json:"pointer"`
	JSON    casbPostureContentExportResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureContentExportResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CasbPostureContentExportResponseEnvelopeMessagesSource]
type casbPostureContentExportResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureContentExportResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureContentExportResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
