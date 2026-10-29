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

// CasbPostureFindingTypeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbPostureFindingTypeService] method instead.
type CasbPostureFindingTypeService struct {
	Options          []option.RequestOption
	RemediationTypes *CasbPostureFindingTypeRemediationTypeService
}

// NewCasbPostureFindingTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCasbPostureFindingTypeService(opts ...option.RequestOption) (r *CasbPostureFindingTypeService) {
	r = &CasbPostureFindingTypeService{}
	r.Options = opts
	r.RemediationTypes = NewCasbPostureFindingTypeRemediationTypeService(opts...)
	return
}

// Create a new custom finding type with policy definition.
func (r *CasbPostureFindingTypeService) New(ctx context.Context, params CasbPostureFindingTypeNewParams, opts ...option.RequestOption) (res *CasbPostureFindingTypeNewResponse, err error) {
	var env CasbPostureFindingTypeNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/finding_types", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update an existing custom finding type.
//
// **Immutable fields**: asset_vendor, asset_class, expression, and category cannot
// be changed after creation.
//
// **Security**: Only custom finding types (not system-managed) can be updated.
func (r *CasbPostureFindingTypeService) Update(ctx context.Context, findingTypeID string, params CasbPostureFindingTypeUpdateParams, opts ...option.RequestOption) (res *CasbPostureFindingTypeUpdateResponse, err error) {
	var env CasbPostureFindingTypeUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if findingTypeID == "" {
		err = errors.New("missing required finding_type_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/finding_types/%s", params.AccountID, findingTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all available finding types with pagination support.
func (r *CasbPostureFindingTypeService) List(ctx context.Context, params CasbPostureFindingTypeListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[CasbPostureFindingTypeListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/finding_types", params.AccountID)
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

// List all available finding types with pagination support.
func (r *CasbPostureFindingTypeService) ListAutoPaging(ctx context.Context, params CasbPostureFindingTypeListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[CasbPostureFindingTypeListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete an existing custom finding type.
//
// **Security**: Only custom finding types (not system-managed) can be deleted.
func (r *CasbPostureFindingTypeService) Delete(ctx context.Context, findingTypeID string, body CasbPostureFindingTypeDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	if findingTypeID == "" {
		err = errors.New("missing required finding_type_id parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/finding_types/%s", body.AccountID, findingTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve a specific finding type by its unique identifier.
func (r *CasbPostureFindingTypeService) Get(ctx context.Context, findingTypeID string, query CasbPostureFindingTypeGetParams, opts ...option.RequestOption) (res *CasbPostureFindingTypeGetResponse, err error) {
	var env CasbPostureFindingTypeGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if findingTypeID == "" {
		err = errors.New("missing required finding_type_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/finding_types/%s", query.AccountID, findingTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Check a candidate custom finding type without creating it, so problems can be
// corrected before publishing. The request body is the same shape as
// `CreateFindingType`, and the candidate is checked against the same rules, so a
// valid result means the definition itself is sound. It is not a reservation: a
// create can still be rejected for a reason that does not belong to the
// definition, such as the account already holding the maximum number of custom
// finding types.
//
// A candidate that fails validation is still a **200** with `result.valid: false`
// and one entry in `result.errors` per problem — that is the answer this endpoint
// exists to give. `4xx` is reserved for a request that could not be processed at
// all, such as a body that will not parse.
func (r *CasbPostureFindingTypeService) Validate(ctx context.Context, params CasbPostureFindingTypeValidateParams, opts ...option.RequestOption) (res *CasbPostureFindingTypeValidateResponse, err error) {
	var env CasbPostureFindingTypeValidateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/finding_types/validate", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Basic finding type information.
type CasbPostureFindingTypeNewResponse struct {
	// The unique identifier of the finding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether this finding type applies to every current and future integration of the
	// vendor.
	AppliesToAllIntegrations bool `json:"applies_to_all_integrations" api:"required"`
	// Asset class the finding type's policy evaluates.
	AssetClass string `json:"asset_class" api:"required"`
	// Canonical vendor enum value.
	AssetVendor CasbPostureFindingTypeNewResponseAssetVendor `json:"asset_vendor" api:"required"`
	// Category information for a finding.
	Category CasbPostureFindingTypeNewResponseCategory `json:"category" api:"required"`
	// Timestamp when the finding type was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Detailed description of the finding.
	Description string `json:"description" api:"required,nullable"`
	// Rego expression backing this finding type. For custom finding types, generated
	// package and import declarations are omitted. System expressions are returned
	// unchanged.
	Expression string `json:"expression" api:"required"`
	// The name of the finding.
	Name string `json:"name" api:"required"`
	// Whether the finding type is system-managed or customer-created.
	Origin CasbPostureFindingTypeNewResponseOrigin `json:"origin" api:"required"`
	// Default severity used when no integration-specific severity override exists.
	Severity CasbPostureFindingTypeNewResponseSeverity `json:"severity" api:"required"`
	// Timestamp when the finding type was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The SaaS/Cloud vendor of the platform with which the finding is associated.
	Vendor string `json:"vendor" api:"required"`
	// Current version of the finding type. Incremented on each update.
	Version int64 `json:"version" api:"required"`
	// Integrations this finding type applies to when applies_to_all_integrations is
	// false.
	IntegrationIDs []string `json:"integration_ids" format:"uuid"`
	// Number of currently active finding instances across the account's integrations.
	// Zero when none have been detected.
	MatchCount int64 `json:"match_count"`
	// Whether evaluation of this finding type is suspended by an operator. Present
	// only when true.
	Quarantined bool                                  `json:"quarantined"`
	JSON        casbPostureFindingTypeNewResponseJSON `json:"-"`
}

// casbPostureFindingTypeNewResponseJSON contains the JSON metadata for the struct
// [CasbPostureFindingTypeNewResponse]
type casbPostureFindingTypeNewResponseJSON struct {
	ID                       apijson.Field
	AppliesToAllIntegrations apijson.Field
	AssetClass               apijson.Field
	AssetVendor              apijson.Field
	Category                 apijson.Field
	CreatedAt                apijson.Field
	Description              apijson.Field
	Expression               apijson.Field
	Name                     apijson.Field
	Origin                   apijson.Field
	Severity                 apijson.Field
	UpdatedAt                apijson.Field
	Vendor                   apijson.Field
	Version                  apijson.Field
	IntegrationIDs           apijson.Field
	MatchCount               apijson.Field
	Quarantined              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CasbPostureFindingTypeNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeNewResponseJSON) RawJSON() string {
	return r.raw
}

// Canonical vendor enum value.
type CasbPostureFindingTypeNewResponseAssetVendor string

const (
	CasbPostureFindingTypeNewResponseAssetVendorAnthropic           CasbPostureFindingTypeNewResponseAssetVendor = "ANTHROPIC"
	CasbPostureFindingTypeNewResponseAssetVendorAws                 CasbPostureFindingTypeNewResponseAssetVendor = "AWS"
	CasbPostureFindingTypeNewResponseAssetVendorBitbucket           CasbPostureFindingTypeNewResponseAssetVendor = "BITBUCKET"
	CasbPostureFindingTypeNewResponseAssetVendorBox                 CasbPostureFindingTypeNewResponseAssetVendor = "BOX"
	CasbPostureFindingTypeNewResponseAssetVendorConfluence          CasbPostureFindingTypeNewResponseAssetVendor = "CONFLUENCE"
	CasbPostureFindingTypeNewResponseAssetVendorDropbox             CasbPostureFindingTypeNewResponseAssetVendor = "DROPBOX"
	CasbPostureFindingTypeNewResponseAssetVendorGitHub              CasbPostureFindingTypeNewResponseAssetVendor = "GITHUB"
	CasbPostureFindingTypeNewResponseAssetVendorGoogleCloudPlatform CasbPostureFindingTypeNewResponseAssetVendor = "GOOGLE_CLOUD_PLATFORM"
	CasbPostureFindingTypeNewResponseAssetVendorGoogleWorkspace     CasbPostureFindingTypeNewResponseAssetVendor = "GOOGLE_WORKSPACE"
	CasbPostureFindingTypeNewResponseAssetVendorJira                CasbPostureFindingTypeNewResponseAssetVendor = "JIRA"
	CasbPostureFindingTypeNewResponseAssetVendorMicrosoft           CasbPostureFindingTypeNewResponseAssetVendor = "MICROSOFT"
	CasbPostureFindingTypeNewResponseAssetVendorOpenAI              CasbPostureFindingTypeNewResponseAssetVendor = "OPENAI"
	CasbPostureFindingTypeNewResponseAssetVendorSalesforce          CasbPostureFindingTypeNewResponseAssetVendor = "SALESFORCE"
	CasbPostureFindingTypeNewResponseAssetVendorServicenow          CasbPostureFindingTypeNewResponseAssetVendor = "SERVICENOW"
	CasbPostureFindingTypeNewResponseAssetVendorSlack               CasbPostureFindingTypeNewResponseAssetVendor = "SLACK"
	CasbPostureFindingTypeNewResponseAssetVendorZoom                CasbPostureFindingTypeNewResponseAssetVendor = "ZOOM"
)

func (r CasbPostureFindingTypeNewResponseAssetVendor) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeNewResponseAssetVendorAnthropic, CasbPostureFindingTypeNewResponseAssetVendorAws, CasbPostureFindingTypeNewResponseAssetVendorBitbucket, CasbPostureFindingTypeNewResponseAssetVendorBox, CasbPostureFindingTypeNewResponseAssetVendorConfluence, CasbPostureFindingTypeNewResponseAssetVendorDropbox, CasbPostureFindingTypeNewResponseAssetVendorGitHub, CasbPostureFindingTypeNewResponseAssetVendorGoogleCloudPlatform, CasbPostureFindingTypeNewResponseAssetVendorGoogleWorkspace, CasbPostureFindingTypeNewResponseAssetVendorJira, CasbPostureFindingTypeNewResponseAssetVendorMicrosoft, CasbPostureFindingTypeNewResponseAssetVendorOpenAI, CasbPostureFindingTypeNewResponseAssetVendorSalesforce, CasbPostureFindingTypeNewResponseAssetVendorServicenow, CasbPostureFindingTypeNewResponseAssetVendorSlack, CasbPostureFindingTypeNewResponseAssetVendorZoom:
		return true
	}
	return false
}

// Category information for a finding.
type CasbPostureFindingTypeNewResponseCategory struct {
	// The type of the observation.
	Observation CasbPostureFindingTypeNewResponseCategoryObservation `json:"observation" api:"required"`
	// The product category.
	Product CasbPostureFindingTypeNewResponseCategoryProduct `json:"product" api:"required"`
	// The type of the finding category.
	Type CasbPostureFindingTypeNewResponseCategoryType `json:"type" api:"required"`
	JSON casbPostureFindingTypeNewResponseCategoryJSON `json:"-"`
}

// casbPostureFindingTypeNewResponseCategoryJSON contains the JSON metadata for the
// struct [CasbPostureFindingTypeNewResponseCategory]
type casbPostureFindingTypeNewResponseCategoryJSON struct {
	Observation apijson.Field
	Product     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeNewResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeNewResponseCategoryJSON) RawJSON() string {
	return r.raw
}

// The type of the observation.
type CasbPostureFindingTypeNewResponseCategoryObservation string

const (
	CasbPostureFindingTypeNewResponseCategoryObservationIssue    CasbPostureFindingTypeNewResponseCategoryObservation = "Issue"
	CasbPostureFindingTypeNewResponseCategoryObservationInsight  CasbPostureFindingTypeNewResponseCategoryObservation = "Insight"
	CasbPostureFindingTypeNewResponseCategoryObservationActivity CasbPostureFindingTypeNewResponseCategoryObservation = "Activity"
)

func (r CasbPostureFindingTypeNewResponseCategoryObservation) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeNewResponseCategoryObservationIssue, CasbPostureFindingTypeNewResponseCategoryObservationInsight, CasbPostureFindingTypeNewResponseCategoryObservationActivity:
		return true
	}
	return false
}

// The product category.
type CasbPostureFindingTypeNewResponseCategoryProduct string

const (
	CasbPostureFindingTypeNewResponseCategoryProductSaaS  CasbPostureFindingTypeNewResponseCategoryProduct = "SaaS"
	CasbPostureFindingTypeNewResponseCategoryProductCloud CasbPostureFindingTypeNewResponseCategoryProduct = "Cloud"
)

func (r CasbPostureFindingTypeNewResponseCategoryProduct) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeNewResponseCategoryProductSaaS, CasbPostureFindingTypeNewResponseCategoryProductCloud:
		return true
	}
	return false
}

// The type of the finding category.
type CasbPostureFindingTypeNewResponseCategoryType string

const (
	CasbPostureFindingTypeNewResponseCategoryTypeContent CasbPostureFindingTypeNewResponseCategoryType = "Content"
	CasbPostureFindingTypeNewResponseCategoryTypePosture CasbPostureFindingTypeNewResponseCategoryType = "Posture"
)

func (r CasbPostureFindingTypeNewResponseCategoryType) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeNewResponseCategoryTypeContent, CasbPostureFindingTypeNewResponseCategoryTypePosture:
		return true
	}
	return false
}

// Whether the finding type is system-managed or customer-created.
type CasbPostureFindingTypeNewResponseOrigin string

const (
	CasbPostureFindingTypeNewResponseOriginSystem CasbPostureFindingTypeNewResponseOrigin = "system"
	CasbPostureFindingTypeNewResponseOriginCustom CasbPostureFindingTypeNewResponseOrigin = "custom"
)

func (r CasbPostureFindingTypeNewResponseOrigin) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeNewResponseOriginSystem, CasbPostureFindingTypeNewResponseOriginCustom:
		return true
	}
	return false
}

// Default severity used when no integration-specific severity override exists.
type CasbPostureFindingTypeNewResponseSeverity string

const (
	CasbPostureFindingTypeNewResponseSeverityCritical CasbPostureFindingTypeNewResponseSeverity = "Critical"
	CasbPostureFindingTypeNewResponseSeverityHigh     CasbPostureFindingTypeNewResponseSeverity = "High"
	CasbPostureFindingTypeNewResponseSeverityMedium   CasbPostureFindingTypeNewResponseSeverity = "Medium"
	CasbPostureFindingTypeNewResponseSeverityLow      CasbPostureFindingTypeNewResponseSeverity = "Low"
)

func (r CasbPostureFindingTypeNewResponseSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeNewResponseSeverityCritical, CasbPostureFindingTypeNewResponseSeverityHigh, CasbPostureFindingTypeNewResponseSeverityMedium, CasbPostureFindingTypeNewResponseSeverityLow:
		return true
	}
	return false
}

// Basic finding type information.
type CasbPostureFindingTypeUpdateResponse struct {
	// The unique identifier of the finding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether this finding type applies to every current and future integration of the
	// vendor.
	AppliesToAllIntegrations bool `json:"applies_to_all_integrations" api:"required"`
	// Asset class the finding type's policy evaluates.
	AssetClass string `json:"asset_class" api:"required"`
	// Canonical vendor enum value.
	AssetVendor CasbPostureFindingTypeUpdateResponseAssetVendor `json:"asset_vendor" api:"required"`
	// Category information for a finding.
	Category CasbPostureFindingTypeUpdateResponseCategory `json:"category" api:"required"`
	// Timestamp when the finding type was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Detailed description of the finding.
	Description string `json:"description" api:"required,nullable"`
	// Rego expression backing this finding type. For custom finding types, generated
	// package and import declarations are omitted. System expressions are returned
	// unchanged.
	Expression string `json:"expression" api:"required"`
	// The name of the finding.
	Name string `json:"name" api:"required"`
	// Whether the finding type is system-managed or customer-created.
	Origin CasbPostureFindingTypeUpdateResponseOrigin `json:"origin" api:"required"`
	// Default severity used when no integration-specific severity override exists.
	Severity CasbPostureFindingTypeUpdateResponseSeverity `json:"severity" api:"required"`
	// Timestamp when the finding type was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The SaaS/Cloud vendor of the platform with which the finding is associated.
	Vendor string `json:"vendor" api:"required"`
	// Current version of the finding type. Incremented on each update.
	Version int64 `json:"version" api:"required"`
	// Integrations this finding type applies to when applies_to_all_integrations is
	// false.
	IntegrationIDs []string `json:"integration_ids" format:"uuid"`
	// Number of currently active finding instances across the account's integrations.
	// Zero when none have been detected.
	MatchCount int64 `json:"match_count"`
	// Whether evaluation of this finding type is suspended by an operator. Present
	// only when true.
	Quarantined bool                                     `json:"quarantined"`
	JSON        casbPostureFindingTypeUpdateResponseJSON `json:"-"`
}

// casbPostureFindingTypeUpdateResponseJSON contains the JSON metadata for the
// struct [CasbPostureFindingTypeUpdateResponse]
type casbPostureFindingTypeUpdateResponseJSON struct {
	ID                       apijson.Field
	AppliesToAllIntegrations apijson.Field
	AssetClass               apijson.Field
	AssetVendor              apijson.Field
	Category                 apijson.Field
	CreatedAt                apijson.Field
	Description              apijson.Field
	Expression               apijson.Field
	Name                     apijson.Field
	Origin                   apijson.Field
	Severity                 apijson.Field
	UpdatedAt                apijson.Field
	Vendor                   apijson.Field
	Version                  apijson.Field
	IntegrationIDs           apijson.Field
	MatchCount               apijson.Field
	Quarantined              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CasbPostureFindingTypeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Canonical vendor enum value.
type CasbPostureFindingTypeUpdateResponseAssetVendor string

const (
	CasbPostureFindingTypeUpdateResponseAssetVendorAnthropic           CasbPostureFindingTypeUpdateResponseAssetVendor = "ANTHROPIC"
	CasbPostureFindingTypeUpdateResponseAssetVendorAws                 CasbPostureFindingTypeUpdateResponseAssetVendor = "AWS"
	CasbPostureFindingTypeUpdateResponseAssetVendorBitbucket           CasbPostureFindingTypeUpdateResponseAssetVendor = "BITBUCKET"
	CasbPostureFindingTypeUpdateResponseAssetVendorBox                 CasbPostureFindingTypeUpdateResponseAssetVendor = "BOX"
	CasbPostureFindingTypeUpdateResponseAssetVendorConfluence          CasbPostureFindingTypeUpdateResponseAssetVendor = "CONFLUENCE"
	CasbPostureFindingTypeUpdateResponseAssetVendorDropbox             CasbPostureFindingTypeUpdateResponseAssetVendor = "DROPBOX"
	CasbPostureFindingTypeUpdateResponseAssetVendorGitHub              CasbPostureFindingTypeUpdateResponseAssetVendor = "GITHUB"
	CasbPostureFindingTypeUpdateResponseAssetVendorGoogleCloudPlatform CasbPostureFindingTypeUpdateResponseAssetVendor = "GOOGLE_CLOUD_PLATFORM"
	CasbPostureFindingTypeUpdateResponseAssetVendorGoogleWorkspace     CasbPostureFindingTypeUpdateResponseAssetVendor = "GOOGLE_WORKSPACE"
	CasbPostureFindingTypeUpdateResponseAssetVendorJira                CasbPostureFindingTypeUpdateResponseAssetVendor = "JIRA"
	CasbPostureFindingTypeUpdateResponseAssetVendorMicrosoft           CasbPostureFindingTypeUpdateResponseAssetVendor = "MICROSOFT"
	CasbPostureFindingTypeUpdateResponseAssetVendorOpenAI              CasbPostureFindingTypeUpdateResponseAssetVendor = "OPENAI"
	CasbPostureFindingTypeUpdateResponseAssetVendorSalesforce          CasbPostureFindingTypeUpdateResponseAssetVendor = "SALESFORCE"
	CasbPostureFindingTypeUpdateResponseAssetVendorServicenow          CasbPostureFindingTypeUpdateResponseAssetVendor = "SERVICENOW"
	CasbPostureFindingTypeUpdateResponseAssetVendorSlack               CasbPostureFindingTypeUpdateResponseAssetVendor = "SLACK"
	CasbPostureFindingTypeUpdateResponseAssetVendorZoom                CasbPostureFindingTypeUpdateResponseAssetVendor = "ZOOM"
)

func (r CasbPostureFindingTypeUpdateResponseAssetVendor) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeUpdateResponseAssetVendorAnthropic, CasbPostureFindingTypeUpdateResponseAssetVendorAws, CasbPostureFindingTypeUpdateResponseAssetVendorBitbucket, CasbPostureFindingTypeUpdateResponseAssetVendorBox, CasbPostureFindingTypeUpdateResponseAssetVendorConfluence, CasbPostureFindingTypeUpdateResponseAssetVendorDropbox, CasbPostureFindingTypeUpdateResponseAssetVendorGitHub, CasbPostureFindingTypeUpdateResponseAssetVendorGoogleCloudPlatform, CasbPostureFindingTypeUpdateResponseAssetVendorGoogleWorkspace, CasbPostureFindingTypeUpdateResponseAssetVendorJira, CasbPostureFindingTypeUpdateResponseAssetVendorMicrosoft, CasbPostureFindingTypeUpdateResponseAssetVendorOpenAI, CasbPostureFindingTypeUpdateResponseAssetVendorSalesforce, CasbPostureFindingTypeUpdateResponseAssetVendorServicenow, CasbPostureFindingTypeUpdateResponseAssetVendorSlack, CasbPostureFindingTypeUpdateResponseAssetVendorZoom:
		return true
	}
	return false
}

// Category information for a finding.
type CasbPostureFindingTypeUpdateResponseCategory struct {
	// The type of the observation.
	Observation CasbPostureFindingTypeUpdateResponseCategoryObservation `json:"observation" api:"required"`
	// The product category.
	Product CasbPostureFindingTypeUpdateResponseCategoryProduct `json:"product" api:"required"`
	// The type of the finding category.
	Type CasbPostureFindingTypeUpdateResponseCategoryType `json:"type" api:"required"`
	JSON casbPostureFindingTypeUpdateResponseCategoryJSON `json:"-"`
}

// casbPostureFindingTypeUpdateResponseCategoryJSON contains the JSON metadata for
// the struct [CasbPostureFindingTypeUpdateResponseCategory]
type casbPostureFindingTypeUpdateResponseCategoryJSON struct {
	Observation apijson.Field
	Product     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeUpdateResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeUpdateResponseCategoryJSON) RawJSON() string {
	return r.raw
}

// The type of the observation.
type CasbPostureFindingTypeUpdateResponseCategoryObservation string

const (
	CasbPostureFindingTypeUpdateResponseCategoryObservationIssue    CasbPostureFindingTypeUpdateResponseCategoryObservation = "Issue"
	CasbPostureFindingTypeUpdateResponseCategoryObservationInsight  CasbPostureFindingTypeUpdateResponseCategoryObservation = "Insight"
	CasbPostureFindingTypeUpdateResponseCategoryObservationActivity CasbPostureFindingTypeUpdateResponseCategoryObservation = "Activity"
)

func (r CasbPostureFindingTypeUpdateResponseCategoryObservation) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeUpdateResponseCategoryObservationIssue, CasbPostureFindingTypeUpdateResponseCategoryObservationInsight, CasbPostureFindingTypeUpdateResponseCategoryObservationActivity:
		return true
	}
	return false
}

// The product category.
type CasbPostureFindingTypeUpdateResponseCategoryProduct string

const (
	CasbPostureFindingTypeUpdateResponseCategoryProductSaaS  CasbPostureFindingTypeUpdateResponseCategoryProduct = "SaaS"
	CasbPostureFindingTypeUpdateResponseCategoryProductCloud CasbPostureFindingTypeUpdateResponseCategoryProduct = "Cloud"
)

func (r CasbPostureFindingTypeUpdateResponseCategoryProduct) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeUpdateResponseCategoryProductSaaS, CasbPostureFindingTypeUpdateResponseCategoryProductCloud:
		return true
	}
	return false
}

// The type of the finding category.
type CasbPostureFindingTypeUpdateResponseCategoryType string

const (
	CasbPostureFindingTypeUpdateResponseCategoryTypeContent CasbPostureFindingTypeUpdateResponseCategoryType = "Content"
	CasbPostureFindingTypeUpdateResponseCategoryTypePosture CasbPostureFindingTypeUpdateResponseCategoryType = "Posture"
)

func (r CasbPostureFindingTypeUpdateResponseCategoryType) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeUpdateResponseCategoryTypeContent, CasbPostureFindingTypeUpdateResponseCategoryTypePosture:
		return true
	}
	return false
}

// Whether the finding type is system-managed or customer-created.
type CasbPostureFindingTypeUpdateResponseOrigin string

const (
	CasbPostureFindingTypeUpdateResponseOriginSystem CasbPostureFindingTypeUpdateResponseOrigin = "system"
	CasbPostureFindingTypeUpdateResponseOriginCustom CasbPostureFindingTypeUpdateResponseOrigin = "custom"
)

func (r CasbPostureFindingTypeUpdateResponseOrigin) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeUpdateResponseOriginSystem, CasbPostureFindingTypeUpdateResponseOriginCustom:
		return true
	}
	return false
}

// Default severity used when no integration-specific severity override exists.
type CasbPostureFindingTypeUpdateResponseSeverity string

const (
	CasbPostureFindingTypeUpdateResponseSeverityCritical CasbPostureFindingTypeUpdateResponseSeverity = "Critical"
	CasbPostureFindingTypeUpdateResponseSeverityHigh     CasbPostureFindingTypeUpdateResponseSeverity = "High"
	CasbPostureFindingTypeUpdateResponseSeverityMedium   CasbPostureFindingTypeUpdateResponseSeverity = "Medium"
	CasbPostureFindingTypeUpdateResponseSeverityLow      CasbPostureFindingTypeUpdateResponseSeverity = "Low"
)

func (r CasbPostureFindingTypeUpdateResponseSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeUpdateResponseSeverityCritical, CasbPostureFindingTypeUpdateResponseSeverityHigh, CasbPostureFindingTypeUpdateResponseSeverityMedium, CasbPostureFindingTypeUpdateResponseSeverityLow:
		return true
	}
	return false
}

// Basic finding type information.
type CasbPostureFindingTypeListResponse struct {
	// The unique identifier of the finding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Category information for a finding.
	Category CasbPostureFindingTypeListResponseCategory `json:"category" api:"required"`
	// The name of the finding.
	Name string `json:"name" api:"required"`
	// Default severity used when no integration-specific severity override exists.
	Severity CasbPostureFindingTypeListResponseSeverity `json:"severity" api:"required"`
	// The SaaS/Cloud vendor of the platform with which the finding is associated.
	Vendor string `json:"vendor" api:"required"`
	// Whether this finding type applies to every current and future integration of the
	// vendor.
	AppliesToAllIntegrations bool `json:"applies_to_all_integrations"`
	// Asset class the finding type's policy evaluates.
	AssetClass string `json:"asset_class"`
	// Canonical vendor enum value.
	AssetVendor CasbPostureFindingTypeListResponseAssetVendor `json:"asset_vendor"`
	// Timestamp when the finding type was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Detailed description of the finding.
	Description string `json:"description" api:"nullable"`
	// Rego expression backing this finding type. For custom finding types, generated
	// package and import declarations are omitted. System expressions are returned
	// unchanged.
	Expression string `json:"expression"`
	// Integrations this finding type applies to when applies_to_all_integrations is
	// false.
	IntegrationIDs []string `json:"integration_ids" format:"uuid"`
	// Number of currently active finding instances across the account's integrations.
	// Zero when none have been detected.
	MatchCount int64 `json:"match_count"`
	// Whether the finding type is system-managed or customer-created.
	Origin CasbPostureFindingTypeListResponseOrigin `json:"origin"`
	// Whether evaluation of this finding type is suspended by an operator. Present
	// only when true.
	Quarantined bool `json:"quarantined"`
	// Timestamp when the finding type was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Current version of the finding type. Incremented on each update.
	Version int64                                  `json:"version"`
	JSON    casbPostureFindingTypeListResponseJSON `json:"-"`
}

// casbPostureFindingTypeListResponseJSON contains the JSON metadata for the struct
// [CasbPostureFindingTypeListResponse]
type casbPostureFindingTypeListResponseJSON struct {
	ID                       apijson.Field
	Category                 apijson.Field
	Name                     apijson.Field
	Severity                 apijson.Field
	Vendor                   apijson.Field
	AppliesToAllIntegrations apijson.Field
	AssetClass               apijson.Field
	AssetVendor              apijson.Field
	CreatedAt                apijson.Field
	Description              apijson.Field
	Expression               apijson.Field
	IntegrationIDs           apijson.Field
	MatchCount               apijson.Field
	Origin                   apijson.Field
	Quarantined              apijson.Field
	UpdatedAt                apijson.Field
	Version                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CasbPostureFindingTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeListResponseJSON) RawJSON() string {
	return r.raw
}

// Category information for a finding.
type CasbPostureFindingTypeListResponseCategory struct {
	// The type of the observation.
	Observation CasbPostureFindingTypeListResponseCategoryObservation `json:"observation" api:"required"`
	// The product category.
	Product CasbPostureFindingTypeListResponseCategoryProduct `json:"product" api:"required"`
	// The type of the finding category.
	Type CasbPostureFindingTypeListResponseCategoryType `json:"type" api:"required"`
	JSON casbPostureFindingTypeListResponseCategoryJSON `json:"-"`
}

// casbPostureFindingTypeListResponseCategoryJSON contains the JSON metadata for
// the struct [CasbPostureFindingTypeListResponseCategory]
type casbPostureFindingTypeListResponseCategoryJSON struct {
	Observation apijson.Field
	Product     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeListResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeListResponseCategoryJSON) RawJSON() string {
	return r.raw
}

// The type of the observation.
type CasbPostureFindingTypeListResponseCategoryObservation string

const (
	CasbPostureFindingTypeListResponseCategoryObservationIssue    CasbPostureFindingTypeListResponseCategoryObservation = "Issue"
	CasbPostureFindingTypeListResponseCategoryObservationInsight  CasbPostureFindingTypeListResponseCategoryObservation = "Insight"
	CasbPostureFindingTypeListResponseCategoryObservationActivity CasbPostureFindingTypeListResponseCategoryObservation = "Activity"
)

func (r CasbPostureFindingTypeListResponseCategoryObservation) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeListResponseCategoryObservationIssue, CasbPostureFindingTypeListResponseCategoryObservationInsight, CasbPostureFindingTypeListResponseCategoryObservationActivity:
		return true
	}
	return false
}

// The product category.
type CasbPostureFindingTypeListResponseCategoryProduct string

const (
	CasbPostureFindingTypeListResponseCategoryProductSaaS  CasbPostureFindingTypeListResponseCategoryProduct = "SaaS"
	CasbPostureFindingTypeListResponseCategoryProductCloud CasbPostureFindingTypeListResponseCategoryProduct = "Cloud"
)

func (r CasbPostureFindingTypeListResponseCategoryProduct) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeListResponseCategoryProductSaaS, CasbPostureFindingTypeListResponseCategoryProductCloud:
		return true
	}
	return false
}

// The type of the finding category.
type CasbPostureFindingTypeListResponseCategoryType string

const (
	CasbPostureFindingTypeListResponseCategoryTypeContent CasbPostureFindingTypeListResponseCategoryType = "Content"
	CasbPostureFindingTypeListResponseCategoryTypePosture CasbPostureFindingTypeListResponseCategoryType = "Posture"
)

func (r CasbPostureFindingTypeListResponseCategoryType) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeListResponseCategoryTypeContent, CasbPostureFindingTypeListResponseCategoryTypePosture:
		return true
	}
	return false
}

// Default severity used when no integration-specific severity override exists.
type CasbPostureFindingTypeListResponseSeverity string

const (
	CasbPostureFindingTypeListResponseSeverityCritical CasbPostureFindingTypeListResponseSeverity = "Critical"
	CasbPostureFindingTypeListResponseSeverityHigh     CasbPostureFindingTypeListResponseSeverity = "High"
	CasbPostureFindingTypeListResponseSeverityMedium   CasbPostureFindingTypeListResponseSeverity = "Medium"
	CasbPostureFindingTypeListResponseSeverityLow      CasbPostureFindingTypeListResponseSeverity = "Low"
)

func (r CasbPostureFindingTypeListResponseSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeListResponseSeverityCritical, CasbPostureFindingTypeListResponseSeverityHigh, CasbPostureFindingTypeListResponseSeverityMedium, CasbPostureFindingTypeListResponseSeverityLow:
		return true
	}
	return false
}

// Canonical vendor enum value.
type CasbPostureFindingTypeListResponseAssetVendor string

const (
	CasbPostureFindingTypeListResponseAssetVendorAnthropic           CasbPostureFindingTypeListResponseAssetVendor = "ANTHROPIC"
	CasbPostureFindingTypeListResponseAssetVendorAws                 CasbPostureFindingTypeListResponseAssetVendor = "AWS"
	CasbPostureFindingTypeListResponseAssetVendorBitbucket           CasbPostureFindingTypeListResponseAssetVendor = "BITBUCKET"
	CasbPostureFindingTypeListResponseAssetVendorBox                 CasbPostureFindingTypeListResponseAssetVendor = "BOX"
	CasbPostureFindingTypeListResponseAssetVendorConfluence          CasbPostureFindingTypeListResponseAssetVendor = "CONFLUENCE"
	CasbPostureFindingTypeListResponseAssetVendorDropbox             CasbPostureFindingTypeListResponseAssetVendor = "DROPBOX"
	CasbPostureFindingTypeListResponseAssetVendorGitHub              CasbPostureFindingTypeListResponseAssetVendor = "GITHUB"
	CasbPostureFindingTypeListResponseAssetVendorGoogleCloudPlatform CasbPostureFindingTypeListResponseAssetVendor = "GOOGLE_CLOUD_PLATFORM"
	CasbPostureFindingTypeListResponseAssetVendorGoogleWorkspace     CasbPostureFindingTypeListResponseAssetVendor = "GOOGLE_WORKSPACE"
	CasbPostureFindingTypeListResponseAssetVendorJira                CasbPostureFindingTypeListResponseAssetVendor = "JIRA"
	CasbPostureFindingTypeListResponseAssetVendorMicrosoft           CasbPostureFindingTypeListResponseAssetVendor = "MICROSOFT"
	CasbPostureFindingTypeListResponseAssetVendorOpenAI              CasbPostureFindingTypeListResponseAssetVendor = "OPENAI"
	CasbPostureFindingTypeListResponseAssetVendorSalesforce          CasbPostureFindingTypeListResponseAssetVendor = "SALESFORCE"
	CasbPostureFindingTypeListResponseAssetVendorServicenow          CasbPostureFindingTypeListResponseAssetVendor = "SERVICENOW"
	CasbPostureFindingTypeListResponseAssetVendorSlack               CasbPostureFindingTypeListResponseAssetVendor = "SLACK"
	CasbPostureFindingTypeListResponseAssetVendorZoom                CasbPostureFindingTypeListResponseAssetVendor = "ZOOM"
)

func (r CasbPostureFindingTypeListResponseAssetVendor) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeListResponseAssetVendorAnthropic, CasbPostureFindingTypeListResponseAssetVendorAws, CasbPostureFindingTypeListResponseAssetVendorBitbucket, CasbPostureFindingTypeListResponseAssetVendorBox, CasbPostureFindingTypeListResponseAssetVendorConfluence, CasbPostureFindingTypeListResponseAssetVendorDropbox, CasbPostureFindingTypeListResponseAssetVendorGitHub, CasbPostureFindingTypeListResponseAssetVendorGoogleCloudPlatform, CasbPostureFindingTypeListResponseAssetVendorGoogleWorkspace, CasbPostureFindingTypeListResponseAssetVendorJira, CasbPostureFindingTypeListResponseAssetVendorMicrosoft, CasbPostureFindingTypeListResponseAssetVendorOpenAI, CasbPostureFindingTypeListResponseAssetVendorSalesforce, CasbPostureFindingTypeListResponseAssetVendorServicenow, CasbPostureFindingTypeListResponseAssetVendorSlack, CasbPostureFindingTypeListResponseAssetVendorZoom:
		return true
	}
	return false
}

// Whether the finding type is system-managed or customer-created.
type CasbPostureFindingTypeListResponseOrigin string

const (
	CasbPostureFindingTypeListResponseOriginSystem CasbPostureFindingTypeListResponseOrigin = "system"
	CasbPostureFindingTypeListResponseOriginCustom CasbPostureFindingTypeListResponseOrigin = "custom"
)

func (r CasbPostureFindingTypeListResponseOrigin) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeListResponseOriginSystem, CasbPostureFindingTypeListResponseOriginCustom:
		return true
	}
	return false
}

// Basic finding type information.
type CasbPostureFindingTypeGetResponse struct {
	// The unique identifier of the finding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Category information for a finding.
	Category CasbPostureFindingTypeGetResponseCategory `json:"category" api:"required"`
	// The name of the finding.
	Name string `json:"name" api:"required"`
	// Default severity used when no integration-specific severity override exists.
	Severity CasbPostureFindingTypeGetResponseSeverity `json:"severity" api:"required"`
	// The SaaS/Cloud vendor of the platform with which the finding is associated.
	Vendor string `json:"vendor" api:"required"`
	// Whether this finding type applies to every current and future integration of the
	// vendor.
	AppliesToAllIntegrations bool `json:"applies_to_all_integrations"`
	// Asset class the finding type's policy evaluates.
	AssetClass string `json:"asset_class"`
	// Canonical vendor enum value.
	AssetVendor CasbPostureFindingTypeGetResponseAssetVendor `json:"asset_vendor"`
	// Timestamp when the finding type was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Detailed description of the finding.
	Description string `json:"description" api:"nullable"`
	// Rego expression backing this finding type. For custom finding types, generated
	// package and import declarations are omitted. System expressions are returned
	// unchanged.
	Expression string `json:"expression"`
	// Integrations this finding type applies to when applies_to_all_integrations is
	// false.
	IntegrationIDs []string `json:"integration_ids" format:"uuid"`
	// Number of currently active finding instances across the account's integrations.
	// Zero when none have been detected.
	MatchCount int64 `json:"match_count"`
	// Whether the finding type is system-managed or customer-created.
	Origin CasbPostureFindingTypeGetResponseOrigin `json:"origin"`
	// Whether evaluation of this finding type is suspended by an operator. Present
	// only when true.
	Quarantined bool `json:"quarantined"`
	// Timestamp when the finding type was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Current version of the finding type. Incremented on each update.
	Version int64                                 `json:"version"`
	JSON    casbPostureFindingTypeGetResponseJSON `json:"-"`
}

// casbPostureFindingTypeGetResponseJSON contains the JSON metadata for the struct
// [CasbPostureFindingTypeGetResponse]
type casbPostureFindingTypeGetResponseJSON struct {
	ID                       apijson.Field
	Category                 apijson.Field
	Name                     apijson.Field
	Severity                 apijson.Field
	Vendor                   apijson.Field
	AppliesToAllIntegrations apijson.Field
	AssetClass               apijson.Field
	AssetVendor              apijson.Field
	CreatedAt                apijson.Field
	Description              apijson.Field
	Expression               apijson.Field
	IntegrationIDs           apijson.Field
	MatchCount               apijson.Field
	Origin                   apijson.Field
	Quarantined              apijson.Field
	UpdatedAt                apijson.Field
	Version                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CasbPostureFindingTypeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeGetResponseJSON) RawJSON() string {
	return r.raw
}

// Category information for a finding.
type CasbPostureFindingTypeGetResponseCategory struct {
	// The type of the observation.
	Observation CasbPostureFindingTypeGetResponseCategoryObservation `json:"observation" api:"required"`
	// The product category.
	Product CasbPostureFindingTypeGetResponseCategoryProduct `json:"product" api:"required"`
	// The type of the finding category.
	Type CasbPostureFindingTypeGetResponseCategoryType `json:"type" api:"required"`
	JSON casbPostureFindingTypeGetResponseCategoryJSON `json:"-"`
}

// casbPostureFindingTypeGetResponseCategoryJSON contains the JSON metadata for the
// struct [CasbPostureFindingTypeGetResponseCategory]
type casbPostureFindingTypeGetResponseCategoryJSON struct {
	Observation apijson.Field
	Product     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeGetResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeGetResponseCategoryJSON) RawJSON() string {
	return r.raw
}

// The type of the observation.
type CasbPostureFindingTypeGetResponseCategoryObservation string

const (
	CasbPostureFindingTypeGetResponseCategoryObservationIssue    CasbPostureFindingTypeGetResponseCategoryObservation = "Issue"
	CasbPostureFindingTypeGetResponseCategoryObservationInsight  CasbPostureFindingTypeGetResponseCategoryObservation = "Insight"
	CasbPostureFindingTypeGetResponseCategoryObservationActivity CasbPostureFindingTypeGetResponseCategoryObservation = "Activity"
)

func (r CasbPostureFindingTypeGetResponseCategoryObservation) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeGetResponseCategoryObservationIssue, CasbPostureFindingTypeGetResponseCategoryObservationInsight, CasbPostureFindingTypeGetResponseCategoryObservationActivity:
		return true
	}
	return false
}

// The product category.
type CasbPostureFindingTypeGetResponseCategoryProduct string

const (
	CasbPostureFindingTypeGetResponseCategoryProductSaaS  CasbPostureFindingTypeGetResponseCategoryProduct = "SaaS"
	CasbPostureFindingTypeGetResponseCategoryProductCloud CasbPostureFindingTypeGetResponseCategoryProduct = "Cloud"
)

func (r CasbPostureFindingTypeGetResponseCategoryProduct) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeGetResponseCategoryProductSaaS, CasbPostureFindingTypeGetResponseCategoryProductCloud:
		return true
	}
	return false
}

// The type of the finding category.
type CasbPostureFindingTypeGetResponseCategoryType string

const (
	CasbPostureFindingTypeGetResponseCategoryTypeContent CasbPostureFindingTypeGetResponseCategoryType = "Content"
	CasbPostureFindingTypeGetResponseCategoryTypePosture CasbPostureFindingTypeGetResponseCategoryType = "Posture"
)

func (r CasbPostureFindingTypeGetResponseCategoryType) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeGetResponseCategoryTypeContent, CasbPostureFindingTypeGetResponseCategoryTypePosture:
		return true
	}
	return false
}

// Default severity used when no integration-specific severity override exists.
type CasbPostureFindingTypeGetResponseSeverity string

const (
	CasbPostureFindingTypeGetResponseSeverityCritical CasbPostureFindingTypeGetResponseSeverity = "Critical"
	CasbPostureFindingTypeGetResponseSeverityHigh     CasbPostureFindingTypeGetResponseSeverity = "High"
	CasbPostureFindingTypeGetResponseSeverityMedium   CasbPostureFindingTypeGetResponseSeverity = "Medium"
	CasbPostureFindingTypeGetResponseSeverityLow      CasbPostureFindingTypeGetResponseSeverity = "Low"
)

func (r CasbPostureFindingTypeGetResponseSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeGetResponseSeverityCritical, CasbPostureFindingTypeGetResponseSeverityHigh, CasbPostureFindingTypeGetResponseSeverityMedium, CasbPostureFindingTypeGetResponseSeverityLow:
		return true
	}
	return false
}

// Canonical vendor enum value.
type CasbPostureFindingTypeGetResponseAssetVendor string

const (
	CasbPostureFindingTypeGetResponseAssetVendorAnthropic           CasbPostureFindingTypeGetResponseAssetVendor = "ANTHROPIC"
	CasbPostureFindingTypeGetResponseAssetVendorAws                 CasbPostureFindingTypeGetResponseAssetVendor = "AWS"
	CasbPostureFindingTypeGetResponseAssetVendorBitbucket           CasbPostureFindingTypeGetResponseAssetVendor = "BITBUCKET"
	CasbPostureFindingTypeGetResponseAssetVendorBox                 CasbPostureFindingTypeGetResponseAssetVendor = "BOX"
	CasbPostureFindingTypeGetResponseAssetVendorConfluence          CasbPostureFindingTypeGetResponseAssetVendor = "CONFLUENCE"
	CasbPostureFindingTypeGetResponseAssetVendorDropbox             CasbPostureFindingTypeGetResponseAssetVendor = "DROPBOX"
	CasbPostureFindingTypeGetResponseAssetVendorGitHub              CasbPostureFindingTypeGetResponseAssetVendor = "GITHUB"
	CasbPostureFindingTypeGetResponseAssetVendorGoogleCloudPlatform CasbPostureFindingTypeGetResponseAssetVendor = "GOOGLE_CLOUD_PLATFORM"
	CasbPostureFindingTypeGetResponseAssetVendorGoogleWorkspace     CasbPostureFindingTypeGetResponseAssetVendor = "GOOGLE_WORKSPACE"
	CasbPostureFindingTypeGetResponseAssetVendorJira                CasbPostureFindingTypeGetResponseAssetVendor = "JIRA"
	CasbPostureFindingTypeGetResponseAssetVendorMicrosoft           CasbPostureFindingTypeGetResponseAssetVendor = "MICROSOFT"
	CasbPostureFindingTypeGetResponseAssetVendorOpenAI              CasbPostureFindingTypeGetResponseAssetVendor = "OPENAI"
	CasbPostureFindingTypeGetResponseAssetVendorSalesforce          CasbPostureFindingTypeGetResponseAssetVendor = "SALESFORCE"
	CasbPostureFindingTypeGetResponseAssetVendorServicenow          CasbPostureFindingTypeGetResponseAssetVendor = "SERVICENOW"
	CasbPostureFindingTypeGetResponseAssetVendorSlack               CasbPostureFindingTypeGetResponseAssetVendor = "SLACK"
	CasbPostureFindingTypeGetResponseAssetVendorZoom                CasbPostureFindingTypeGetResponseAssetVendor = "ZOOM"
)

func (r CasbPostureFindingTypeGetResponseAssetVendor) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeGetResponseAssetVendorAnthropic, CasbPostureFindingTypeGetResponseAssetVendorAws, CasbPostureFindingTypeGetResponseAssetVendorBitbucket, CasbPostureFindingTypeGetResponseAssetVendorBox, CasbPostureFindingTypeGetResponseAssetVendorConfluence, CasbPostureFindingTypeGetResponseAssetVendorDropbox, CasbPostureFindingTypeGetResponseAssetVendorGitHub, CasbPostureFindingTypeGetResponseAssetVendorGoogleCloudPlatform, CasbPostureFindingTypeGetResponseAssetVendorGoogleWorkspace, CasbPostureFindingTypeGetResponseAssetVendorJira, CasbPostureFindingTypeGetResponseAssetVendorMicrosoft, CasbPostureFindingTypeGetResponseAssetVendorOpenAI, CasbPostureFindingTypeGetResponseAssetVendorSalesforce, CasbPostureFindingTypeGetResponseAssetVendorServicenow, CasbPostureFindingTypeGetResponseAssetVendorSlack, CasbPostureFindingTypeGetResponseAssetVendorZoom:
		return true
	}
	return false
}

// Whether the finding type is system-managed or customer-created.
type CasbPostureFindingTypeGetResponseOrigin string

const (
	CasbPostureFindingTypeGetResponseOriginSystem CasbPostureFindingTypeGetResponseOrigin = "system"
	CasbPostureFindingTypeGetResponseOriginCustom CasbPostureFindingTypeGetResponseOrigin = "custom"
)

func (r CasbPostureFindingTypeGetResponseOrigin) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeGetResponseOriginSystem, CasbPostureFindingTypeGetResponseOriginCustom:
		return true
	}
	return false
}

type CasbPostureFindingTypeValidateResponse struct {
	// Every problem found with the candidate, so all of them can be surfaced at once.
	// Empty when `valid` is true. Problems without a location come first, followed by
	// located problems ordered by `line`, then `column`.
	Errors []CasbPostureFindingTypeValidateResponseError `json:"errors" api:"required"`
	// Whether the candidate finding type would be accepted as-is.
	Valid bool                                       `json:"valid" api:"required"`
	JSON  casbPostureFindingTypeValidateResponseJSON `json:"-"`
}

// casbPostureFindingTypeValidateResponseJSON contains the JSON metadata for the
// struct [CasbPostureFindingTypeValidateResponse]
type casbPostureFindingTypeValidateResponseJSON struct {
	Errors      apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeValidateResponseJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeValidateResponseError struct {
	// Description of the problem.
	Detail string `json:"detail" api:"required"`
	// Deprecated: use `detail`.
	//
	// Deprecated: deprecated
	Message string `json:"message" api:"required"`
	// Machine-readable identifier for the class of problem, when one is available.
	// Omitted otherwise.
	Code string `json:"code"`
	// 1-based column on `line`. Omitted when the problem could not be located in the
	// submitted source.
	Column int64 `json:"column"`
	// Request field the problem belongs to, matching the property names of
	// `ValidateFindingTypeRequest`. Array properties are reported per entry with a
	// zero-based index, for example `integration_ids[0]`. Empty when the problem
	// cannot be attributed to a single submitted field.
	Field string `json:"field"`
	// 1-based line in `expression` where the problem was found. Omitted when the
	// problem could not be located in the submitted source.
	Line int64                                           `json:"line"`
	JSON casbPostureFindingTypeValidateResponseErrorJSON `json:"-"`
}

// casbPostureFindingTypeValidateResponseErrorJSON contains the JSON metadata for
// the struct [CasbPostureFindingTypeValidateResponseError]
type casbPostureFindingTypeValidateResponseErrorJSON struct {
	Detail      apijson.Field
	Message     apijson.Field
	Code        apijson.Field
	Column      apijson.Field
	Field       apijson.Field
	Line        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeValidateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeValidateResponseErrorJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// When true, this finding type applies to every current and future integration of
	// the vendor.
	AppliesToAllIntegrations param.Field[bool] `json:"applies_to_all_integrations" api:"required"`
	// Asset class this finding type applies to (PascalCase, e.g. GoogleWorkspaceFile)
	AssetClass param.Field[string] `json:"asset_class" api:"required"`
	// Canonical vendor enum for the platform whose assets this finding type evaluates.
	AssetVendor param.Field[CasbPostureFindingTypeNewParamsAssetVendor] `json:"asset_vendor" api:"required"`
	// Human-readable description of the finding type. Cannot be empty. Maximum 4096
	// characters.
	Description param.Field[string] `json:"description" api:"required"`
	// Rego expression body for the policy. `package cloudflare.casb.custom` and
	// `import rego.v1` are added automatically. Do not include package or import
	// declarations. This field cannot be changed after creation. Maximum 4096
	// characters.
	Expression param.Field[string] `json:"expression" api:"required"`
	// Human-readable name for the finding type. Maximum 256 characters.
	Name param.Field[string] `json:"name" api:"required"`
	// Default severity used when no integration-specific severity override exists.
	Severity param.Field[CasbPostureFindingTypeNewParamsSeverity] `json:"severity" api:"required"`
	// Required and non-empty when applies_to_all_integrations is false; omit when it
	// is true.
	IntegrationIDs param.Field[[]string] `json:"integration_ids" format:"uuid"`
}

func (r CasbPostureFindingTypeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Canonical vendor enum for the platform whose assets this finding type evaluates.
type CasbPostureFindingTypeNewParamsAssetVendor string

const (
	CasbPostureFindingTypeNewParamsAssetVendorAnthropic           CasbPostureFindingTypeNewParamsAssetVendor = "ANTHROPIC"
	CasbPostureFindingTypeNewParamsAssetVendorAws                 CasbPostureFindingTypeNewParamsAssetVendor = "AWS"
	CasbPostureFindingTypeNewParamsAssetVendorBitbucket           CasbPostureFindingTypeNewParamsAssetVendor = "BITBUCKET"
	CasbPostureFindingTypeNewParamsAssetVendorBox                 CasbPostureFindingTypeNewParamsAssetVendor = "BOX"
	CasbPostureFindingTypeNewParamsAssetVendorConfluence          CasbPostureFindingTypeNewParamsAssetVendor = "CONFLUENCE"
	CasbPostureFindingTypeNewParamsAssetVendorDropbox             CasbPostureFindingTypeNewParamsAssetVendor = "DROPBOX"
	CasbPostureFindingTypeNewParamsAssetVendorGitHub              CasbPostureFindingTypeNewParamsAssetVendor = "GITHUB"
	CasbPostureFindingTypeNewParamsAssetVendorGoogleCloudPlatform CasbPostureFindingTypeNewParamsAssetVendor = "GOOGLE_CLOUD_PLATFORM"
	CasbPostureFindingTypeNewParamsAssetVendorGoogleWorkspace     CasbPostureFindingTypeNewParamsAssetVendor = "GOOGLE_WORKSPACE"
	CasbPostureFindingTypeNewParamsAssetVendorJira                CasbPostureFindingTypeNewParamsAssetVendor = "JIRA"
	CasbPostureFindingTypeNewParamsAssetVendorMicrosoft           CasbPostureFindingTypeNewParamsAssetVendor = "MICROSOFT"
	CasbPostureFindingTypeNewParamsAssetVendorOpenAI              CasbPostureFindingTypeNewParamsAssetVendor = "OPENAI"
	CasbPostureFindingTypeNewParamsAssetVendorSalesforce          CasbPostureFindingTypeNewParamsAssetVendor = "SALESFORCE"
	CasbPostureFindingTypeNewParamsAssetVendorServicenow          CasbPostureFindingTypeNewParamsAssetVendor = "SERVICENOW"
	CasbPostureFindingTypeNewParamsAssetVendorSlack               CasbPostureFindingTypeNewParamsAssetVendor = "SLACK"
	CasbPostureFindingTypeNewParamsAssetVendorZoom                CasbPostureFindingTypeNewParamsAssetVendor = "ZOOM"
)

func (r CasbPostureFindingTypeNewParamsAssetVendor) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeNewParamsAssetVendorAnthropic, CasbPostureFindingTypeNewParamsAssetVendorAws, CasbPostureFindingTypeNewParamsAssetVendorBitbucket, CasbPostureFindingTypeNewParamsAssetVendorBox, CasbPostureFindingTypeNewParamsAssetVendorConfluence, CasbPostureFindingTypeNewParamsAssetVendorDropbox, CasbPostureFindingTypeNewParamsAssetVendorGitHub, CasbPostureFindingTypeNewParamsAssetVendorGoogleCloudPlatform, CasbPostureFindingTypeNewParamsAssetVendorGoogleWorkspace, CasbPostureFindingTypeNewParamsAssetVendorJira, CasbPostureFindingTypeNewParamsAssetVendorMicrosoft, CasbPostureFindingTypeNewParamsAssetVendorOpenAI, CasbPostureFindingTypeNewParamsAssetVendorSalesforce, CasbPostureFindingTypeNewParamsAssetVendorServicenow, CasbPostureFindingTypeNewParamsAssetVendorSlack, CasbPostureFindingTypeNewParamsAssetVendorZoom:
		return true
	}
	return false
}

// Default severity used when no integration-specific severity override exists.
type CasbPostureFindingTypeNewParamsSeverity string

const (
	CasbPostureFindingTypeNewParamsSeverityCritical CasbPostureFindingTypeNewParamsSeverity = "Critical"
	CasbPostureFindingTypeNewParamsSeverityHigh     CasbPostureFindingTypeNewParamsSeverity = "High"
	CasbPostureFindingTypeNewParamsSeverityMedium   CasbPostureFindingTypeNewParamsSeverity = "Medium"
	CasbPostureFindingTypeNewParamsSeverityLow      CasbPostureFindingTypeNewParamsSeverity = "Low"
)

func (r CasbPostureFindingTypeNewParamsSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeNewParamsSeverityCritical, CasbPostureFindingTypeNewParamsSeverityHigh, CasbPostureFindingTypeNewParamsSeverityMedium, CasbPostureFindingTypeNewParamsSeverityLow:
		return true
	}
	return false
}

// Common response structure for all API endpoints.
type CasbPostureFindingTypeNewResponseEnvelope struct {
	Errors   []CasbPostureFindingTypeNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingTypeNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Basic finding type information.
	Result CasbPostureFindingTypeNewResponse             `json:"result"`
	JSON   casbPostureFindingTypeNewResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingTypeNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPostureFindingTypeNewResponseEnvelope]
type casbPostureFindingTypeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeNewResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingTypeNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingTypeNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingTypeNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CasbPostureFindingTypeNewResponseEnvelopeErrors]
type casbPostureFindingTypeNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingTypeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeNewResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                    `json:"pointer"`
	JSON    casbPostureFindingTypeNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingTypeNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CasbPostureFindingTypeNewResponseEnvelopeErrorsSource]
type casbPostureFindingTypeNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeNewResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                  `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingTypeNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingTypeNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingTypeNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CasbPostureFindingTypeNewResponseEnvelopeMessages]
type casbPostureFindingTypeNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingTypeNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeNewResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                      `json:"pointer"`
	JSON    casbPostureFindingTypeNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingTypeNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingTypeNewResponseEnvelopeMessagesSource]
type casbPostureFindingTypeNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeUpdateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// When true, this finding type applies to every current and future integration of
	// the vendor.
	AppliesToAllIntegrations param.Field[bool] `json:"applies_to_all_integrations" api:"required"`
	// Optional description of the finding type. Cannot be empty if provided. Maximum
	// 4096 characters.
	Description param.Field[string] `json:"description"`
	// Required and non-empty when applies_to_all_integrations is false; omit when it
	// is true.
	IntegrationIDs param.Field[[]string] `json:"integration_ids" format:"uuid"`
	// Human-readable name for the finding type. Cannot be empty if provided. Maximum
	// 256 characters.
	Name param.Field[string] `json:"name"`
	// Default severity used when no integration-specific severity override exists.
	// Changing it reclassifies existing and future findings that do not have an
	// override.
	Severity param.Field[CasbPostureFindingTypeUpdateParamsSeverity] `json:"severity"`
}

func (r CasbPostureFindingTypeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Default severity used when no integration-specific severity override exists.
// Changing it reclassifies existing and future findings that do not have an
// override.
type CasbPostureFindingTypeUpdateParamsSeverity string

const (
	CasbPostureFindingTypeUpdateParamsSeverityCritical CasbPostureFindingTypeUpdateParamsSeverity = "Critical"
	CasbPostureFindingTypeUpdateParamsSeverityHigh     CasbPostureFindingTypeUpdateParamsSeverity = "High"
	CasbPostureFindingTypeUpdateParamsSeverityMedium   CasbPostureFindingTypeUpdateParamsSeverity = "Medium"
	CasbPostureFindingTypeUpdateParamsSeverityLow      CasbPostureFindingTypeUpdateParamsSeverity = "Low"
)

func (r CasbPostureFindingTypeUpdateParamsSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeUpdateParamsSeverityCritical, CasbPostureFindingTypeUpdateParamsSeverityHigh, CasbPostureFindingTypeUpdateParamsSeverityMedium, CasbPostureFindingTypeUpdateParamsSeverityLow:
		return true
	}
	return false
}

// Common response structure for all API endpoints.
type CasbPostureFindingTypeUpdateResponseEnvelope struct {
	Errors   []CasbPostureFindingTypeUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingTypeUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Basic finding type information.
	Result CasbPostureFindingTypeUpdateResponse             `json:"result"`
	JSON   casbPostureFindingTypeUpdateResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingTypeUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [CasbPostureFindingTypeUpdateResponseEnvelope]
type casbPostureFindingTypeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeUpdateResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                   `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingTypeUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingTypeUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingTypeUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CasbPostureFindingTypeUpdateResponseEnvelopeErrors]
type casbPostureFindingTypeUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingTypeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeUpdateResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                       `json:"pointer"`
	JSON    casbPostureFindingTypeUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingTypeUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingTypeUpdateResponseEnvelopeErrorsSource]
type casbPostureFindingTypeUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeUpdateResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                     `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingTypeUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingTypeUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingTypeUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CasbPostureFindingTypeUpdateResponseEnvelopeMessages]
type casbPostureFindingTypeUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingTypeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeUpdateResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                         `json:"pointer"`
	JSON    casbPostureFindingTypeUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingTypeUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingTypeUpdateResponseEnvelopeMessagesSource]
type casbPostureFindingTypeUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Sort direction for the primary field. Defaults to 'desc'. Ties are resolved by
	// case-insensitive name, then ID, both ascending regardless of direction.
	Direction param.Field[CasbPostureFindingTypeListParamsDirection] `query:"direction"`
	// Return only finding types updated at or before this RFC 3339 timestamp
	// (inclusive).
	MaxUpdatedAt param.Field[time.Time] `query:"max_updated_at" format:"date-time"`
	// Return only finding types updated at or after this RFC 3339 timestamp
	// (inclusive).
	MinUpdatedAt param.Field[time.Time] `query:"min_updated_at" format:"date-time"`
	// Field to sort by (case-insensitive). Defaults to 'created'. Name and vendor
	// sorting is case-insensitive; vendor uses friendly provider names. Ascending
	// order: type Content then Posture; origin custom then system; severity Low,
	// Medium, High, Critical.
	Order param.Field[CasbPostureFindingTypeListParamsOrder] `query:"order"`
	// Filter finding types by origin. Supports multiple comma-separated,
	// case-insensitive values.
	Origins param.Field[[]CasbPostureFindingTypeListParamsOrigin] `query:"origins"`
	// A page number within the paginated result set.
	Page param.Field[int64] `query:"page"`
	// Number of results to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// When true, return only quarantined finding types; when false, return only
	// non-quarantined finding types. Omit to return both.
	Quarantined param.Field[bool] `query:"quarantined"`
	// Filter finding types by name or ID (case-insensitive substring match).
	Search param.Field[string] `query:"search"`
	// Filter finding types by severity. Supports multiple comma-separated,
	// case-insensitive values.
	Severities param.Field[[]CasbPostureFindingTypeListParamsSeverity] `query:"severities"`
	// Filter finding types by category type. Supports multiple comma-separated,
	// case-insensitive values.
	Types param.Field[[]CasbPostureFindingTypeListParamsType] `query:"types"`
	// Filter finding types by vendor. Supports multiple comma-separated values. Each
	// value may be either the vendor enum (GOOGLE_WORKSPACE) or the friendly name
	// returned by the read routes (Google Workspace).
	Vendors param.Field[[]CasbPostureFindingTypeListParamsVendor] `query:"vendors"`
}

// URLQuery serializes [CasbPostureFindingTypeListParams]'s query parameters as
// `url.Values`.
func (r CasbPostureFindingTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction for the primary field. Defaults to 'desc'. Ties are resolved by
// case-insensitive name, then ID, both ascending regardless of direction.
type CasbPostureFindingTypeListParamsDirection string

const (
	CasbPostureFindingTypeListParamsDirectionAsc  CasbPostureFindingTypeListParamsDirection = "asc"
	CasbPostureFindingTypeListParamsDirectionDesc CasbPostureFindingTypeListParamsDirection = "desc"
)

func (r CasbPostureFindingTypeListParamsDirection) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeListParamsDirectionAsc, CasbPostureFindingTypeListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to sort by (case-insensitive). Defaults to 'created'. Name and vendor
// sorting is case-insensitive; vendor uses friendly provider names. Ascending
// order: type Content then Posture; origin custom then system; severity Low,
// Medium, High, Critical.
type CasbPostureFindingTypeListParamsOrder string

const (
	CasbPostureFindingTypeListParamsOrderCreated    CasbPostureFindingTypeListParamsOrder = "created"
	CasbPostureFindingTypeListParamsOrderUpdated    CasbPostureFindingTypeListParamsOrder = "updated"
	CasbPostureFindingTypeListParamsOrderMatchCount CasbPostureFindingTypeListParamsOrder = "match_count"
	CasbPostureFindingTypeListParamsOrderName       CasbPostureFindingTypeListParamsOrder = "name"
	CasbPostureFindingTypeListParamsOrderVendor     CasbPostureFindingTypeListParamsOrder = "vendor"
	CasbPostureFindingTypeListParamsOrderType       CasbPostureFindingTypeListParamsOrder = "type"
	CasbPostureFindingTypeListParamsOrderOrigin     CasbPostureFindingTypeListParamsOrder = "origin"
	CasbPostureFindingTypeListParamsOrderSeverity   CasbPostureFindingTypeListParamsOrder = "severity"
)

func (r CasbPostureFindingTypeListParamsOrder) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeListParamsOrderCreated, CasbPostureFindingTypeListParamsOrderUpdated, CasbPostureFindingTypeListParamsOrderMatchCount, CasbPostureFindingTypeListParamsOrderName, CasbPostureFindingTypeListParamsOrderVendor, CasbPostureFindingTypeListParamsOrderType, CasbPostureFindingTypeListParamsOrderOrigin, CasbPostureFindingTypeListParamsOrderSeverity:
		return true
	}
	return false
}

// Finding type origins for list filters.
type CasbPostureFindingTypeListParamsOrigin string

const (
	CasbPostureFindingTypeListParamsOriginCustom CasbPostureFindingTypeListParamsOrigin = "custom"
	CasbPostureFindingTypeListParamsOriginSystem CasbPostureFindingTypeListParamsOrigin = "system"
)

func (r CasbPostureFindingTypeListParamsOrigin) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeListParamsOriginCustom, CasbPostureFindingTypeListParamsOriginSystem:
		return true
	}
	return false
}

// Finding type severities for list filters.
type CasbPostureFindingTypeListParamsSeverity string

const (
	CasbPostureFindingTypeListParamsSeverityCritical CasbPostureFindingTypeListParamsSeverity = "critical"
	CasbPostureFindingTypeListParamsSeverityHigh     CasbPostureFindingTypeListParamsSeverity = "high"
	CasbPostureFindingTypeListParamsSeverityMedium   CasbPostureFindingTypeListParamsSeverity = "medium"
	CasbPostureFindingTypeListParamsSeverityLow      CasbPostureFindingTypeListParamsSeverity = "low"
)

func (r CasbPostureFindingTypeListParamsSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeListParamsSeverityCritical, CasbPostureFindingTypeListParamsSeverityHigh, CasbPostureFindingTypeListParamsSeverityMedium, CasbPostureFindingTypeListParamsSeverityLow:
		return true
	}
	return false
}

// The type of the finding category.
type CasbPostureFindingTypeListParamsType string

const (
	CasbPostureFindingTypeListParamsTypeContent CasbPostureFindingTypeListParamsType = "Content"
	CasbPostureFindingTypeListParamsTypePosture CasbPostureFindingTypeListParamsType = "Posture"
)

func (r CasbPostureFindingTypeListParamsType) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeListParamsTypeContent, CasbPostureFindingTypeListParamsTypePosture:
		return true
	}
	return false
}

// Vendors whose assets a finding type can evaluate.
type CasbPostureFindingTypeListParamsVendor string

const (
	CasbPostureFindingTypeListParamsVendorAnthropic           CasbPostureFindingTypeListParamsVendor = "ANTHROPIC"
	CasbPostureFindingTypeListParamsVendorAws                 CasbPostureFindingTypeListParamsVendor = "AWS"
	CasbPostureFindingTypeListParamsVendorBitbucket           CasbPostureFindingTypeListParamsVendor = "BITBUCKET"
	CasbPostureFindingTypeListParamsVendorBox                 CasbPostureFindingTypeListParamsVendor = "BOX"
	CasbPostureFindingTypeListParamsVendorConfluence          CasbPostureFindingTypeListParamsVendor = "CONFLUENCE"
	CasbPostureFindingTypeListParamsVendorDropbox             CasbPostureFindingTypeListParamsVendor = "DROPBOX"
	CasbPostureFindingTypeListParamsVendorGitHub              CasbPostureFindingTypeListParamsVendor = "GITHUB"
	CasbPostureFindingTypeListParamsVendorGoogleCloudPlatform CasbPostureFindingTypeListParamsVendor = "GOOGLE_CLOUD_PLATFORM"
	CasbPostureFindingTypeListParamsVendorGoogleWorkspace     CasbPostureFindingTypeListParamsVendor = "GOOGLE_WORKSPACE"
	CasbPostureFindingTypeListParamsVendorJira                CasbPostureFindingTypeListParamsVendor = "JIRA"
	CasbPostureFindingTypeListParamsVendorMicrosoft           CasbPostureFindingTypeListParamsVendor = "MICROSOFT"
	CasbPostureFindingTypeListParamsVendorOpenAI              CasbPostureFindingTypeListParamsVendor = "OPENAI"
	CasbPostureFindingTypeListParamsVendorSalesforce          CasbPostureFindingTypeListParamsVendor = "SALESFORCE"
	CasbPostureFindingTypeListParamsVendorServicenow          CasbPostureFindingTypeListParamsVendor = "SERVICENOW"
	CasbPostureFindingTypeListParamsVendorSlack               CasbPostureFindingTypeListParamsVendor = "SLACK"
	CasbPostureFindingTypeListParamsVendorZoom                CasbPostureFindingTypeListParamsVendor = "ZOOM"
)

func (r CasbPostureFindingTypeListParamsVendor) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeListParamsVendorAnthropic, CasbPostureFindingTypeListParamsVendorAws, CasbPostureFindingTypeListParamsVendorBitbucket, CasbPostureFindingTypeListParamsVendorBox, CasbPostureFindingTypeListParamsVendorConfluence, CasbPostureFindingTypeListParamsVendorDropbox, CasbPostureFindingTypeListParamsVendorGitHub, CasbPostureFindingTypeListParamsVendorGoogleCloudPlatform, CasbPostureFindingTypeListParamsVendorGoogleWorkspace, CasbPostureFindingTypeListParamsVendorJira, CasbPostureFindingTypeListParamsVendorMicrosoft, CasbPostureFindingTypeListParamsVendorOpenAI, CasbPostureFindingTypeListParamsVendorSalesforce, CasbPostureFindingTypeListParamsVendorServicenow, CasbPostureFindingTypeListParamsVendorSlack, CasbPostureFindingTypeListParamsVendorZoom:
		return true
	}
	return false
}

type CasbPostureFindingTypeDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CasbPostureFindingTypeGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

// Common response structure for all API endpoints.
type CasbPostureFindingTypeGetResponseEnvelope struct {
	Errors   []CasbPostureFindingTypeGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingTypeGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Basic finding type information.
	Result CasbPostureFindingTypeGetResponse             `json:"result"`
	JSON   casbPostureFindingTypeGetResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingTypeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPostureFindingTypeGetResponseEnvelope]
type casbPostureFindingTypeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeGetResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingTypeGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingTypeGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingTypeGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CasbPostureFindingTypeGetResponseEnvelopeErrors]
type casbPostureFindingTypeGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingTypeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeGetResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                    `json:"pointer"`
	JSON    casbPostureFindingTypeGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingTypeGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CasbPostureFindingTypeGetResponseEnvelopeErrorsSource]
type casbPostureFindingTypeGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeGetResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                  `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingTypeGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingTypeGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingTypeGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CasbPostureFindingTypeGetResponseEnvelopeMessages]
type casbPostureFindingTypeGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingTypeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeGetResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                      `json:"pointer"`
	JSON    casbPostureFindingTypeGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingTypeGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingTypeGetResponseEnvelopeMessagesSource]
type casbPostureFindingTypeGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeValidateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// When true, this finding type applies to every current and future integration of
	// the vendor.
	AppliesToAllIntegrations param.Field[bool] `json:"applies_to_all_integrations" api:"required"`
	// Asset class this finding type applies to (PascalCase, e.g. GoogleWorkspaceFile)
	AssetClass param.Field[string] `json:"asset_class" api:"required"`
	// Canonical vendor enum for the platform whose assets this finding type evaluates.
	AssetVendor param.Field[CasbPostureFindingTypeValidateParamsAssetVendor] `json:"asset_vendor" api:"required"`
	// Human-readable description of the finding type. Cannot be empty. Maximum 4096
	// characters.
	Description param.Field[string] `json:"description" api:"required"`
	// Rego expression body for the policy. `package cloudflare.casb.custom` and
	// `import rego.v1` are added automatically. Do not include package or import
	// declarations. This field cannot be changed after creation. Maximum 4096
	// characters.
	Expression param.Field[string] `json:"expression" api:"required"`
	// Human-readable name for the finding type. Maximum 256 characters.
	Name param.Field[string] `json:"name" api:"required"`
	// Default severity used when no integration-specific severity override exists.
	Severity param.Field[CasbPostureFindingTypeValidateParamsSeverity] `json:"severity" api:"required"`
	// Required and non-empty when applies_to_all_integrations is false; omit when it
	// is true.
	IntegrationIDs param.Field[[]string] `json:"integration_ids" format:"uuid"`
}

func (r CasbPostureFindingTypeValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Canonical vendor enum for the platform whose assets this finding type evaluates.
type CasbPostureFindingTypeValidateParamsAssetVendor string

const (
	CasbPostureFindingTypeValidateParamsAssetVendorAnthropic           CasbPostureFindingTypeValidateParamsAssetVendor = "ANTHROPIC"
	CasbPostureFindingTypeValidateParamsAssetVendorAws                 CasbPostureFindingTypeValidateParamsAssetVendor = "AWS"
	CasbPostureFindingTypeValidateParamsAssetVendorBitbucket           CasbPostureFindingTypeValidateParamsAssetVendor = "BITBUCKET"
	CasbPostureFindingTypeValidateParamsAssetVendorBox                 CasbPostureFindingTypeValidateParamsAssetVendor = "BOX"
	CasbPostureFindingTypeValidateParamsAssetVendorConfluence          CasbPostureFindingTypeValidateParamsAssetVendor = "CONFLUENCE"
	CasbPostureFindingTypeValidateParamsAssetVendorDropbox             CasbPostureFindingTypeValidateParamsAssetVendor = "DROPBOX"
	CasbPostureFindingTypeValidateParamsAssetVendorGitHub              CasbPostureFindingTypeValidateParamsAssetVendor = "GITHUB"
	CasbPostureFindingTypeValidateParamsAssetVendorGoogleCloudPlatform CasbPostureFindingTypeValidateParamsAssetVendor = "GOOGLE_CLOUD_PLATFORM"
	CasbPostureFindingTypeValidateParamsAssetVendorGoogleWorkspace     CasbPostureFindingTypeValidateParamsAssetVendor = "GOOGLE_WORKSPACE"
	CasbPostureFindingTypeValidateParamsAssetVendorJira                CasbPostureFindingTypeValidateParamsAssetVendor = "JIRA"
	CasbPostureFindingTypeValidateParamsAssetVendorMicrosoft           CasbPostureFindingTypeValidateParamsAssetVendor = "MICROSOFT"
	CasbPostureFindingTypeValidateParamsAssetVendorOpenAI              CasbPostureFindingTypeValidateParamsAssetVendor = "OPENAI"
	CasbPostureFindingTypeValidateParamsAssetVendorSalesforce          CasbPostureFindingTypeValidateParamsAssetVendor = "SALESFORCE"
	CasbPostureFindingTypeValidateParamsAssetVendorServicenow          CasbPostureFindingTypeValidateParamsAssetVendor = "SERVICENOW"
	CasbPostureFindingTypeValidateParamsAssetVendorSlack               CasbPostureFindingTypeValidateParamsAssetVendor = "SLACK"
	CasbPostureFindingTypeValidateParamsAssetVendorZoom                CasbPostureFindingTypeValidateParamsAssetVendor = "ZOOM"
)

func (r CasbPostureFindingTypeValidateParamsAssetVendor) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeValidateParamsAssetVendorAnthropic, CasbPostureFindingTypeValidateParamsAssetVendorAws, CasbPostureFindingTypeValidateParamsAssetVendorBitbucket, CasbPostureFindingTypeValidateParamsAssetVendorBox, CasbPostureFindingTypeValidateParamsAssetVendorConfluence, CasbPostureFindingTypeValidateParamsAssetVendorDropbox, CasbPostureFindingTypeValidateParamsAssetVendorGitHub, CasbPostureFindingTypeValidateParamsAssetVendorGoogleCloudPlatform, CasbPostureFindingTypeValidateParamsAssetVendorGoogleWorkspace, CasbPostureFindingTypeValidateParamsAssetVendorJira, CasbPostureFindingTypeValidateParamsAssetVendorMicrosoft, CasbPostureFindingTypeValidateParamsAssetVendorOpenAI, CasbPostureFindingTypeValidateParamsAssetVendorSalesforce, CasbPostureFindingTypeValidateParamsAssetVendorServicenow, CasbPostureFindingTypeValidateParamsAssetVendorSlack, CasbPostureFindingTypeValidateParamsAssetVendorZoom:
		return true
	}
	return false
}

// Default severity used when no integration-specific severity override exists.
type CasbPostureFindingTypeValidateParamsSeverity string

const (
	CasbPostureFindingTypeValidateParamsSeverityCritical CasbPostureFindingTypeValidateParamsSeverity = "Critical"
	CasbPostureFindingTypeValidateParamsSeverityHigh     CasbPostureFindingTypeValidateParamsSeverity = "High"
	CasbPostureFindingTypeValidateParamsSeverityMedium   CasbPostureFindingTypeValidateParamsSeverity = "Medium"
	CasbPostureFindingTypeValidateParamsSeverityLow      CasbPostureFindingTypeValidateParamsSeverity = "Low"
)

func (r CasbPostureFindingTypeValidateParamsSeverity) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeValidateParamsSeverityCritical, CasbPostureFindingTypeValidateParamsSeverityHigh, CasbPostureFindingTypeValidateParamsSeverityMedium, CasbPostureFindingTypeValidateParamsSeverityLow:
		return true
	}
	return false
}

// Common response structure for all API endpoints.
type CasbPostureFindingTypeValidateResponseEnvelope struct {
	Errors   []CasbPostureFindingTypeValidateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingTypeValidateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool                                               `json:"success" api:"required"`
	Result  CasbPostureFindingTypeValidateResponse             `json:"result"`
	JSON    casbPostureFindingTypeValidateResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingTypeValidateResponseEnvelopeJSON contains the JSON metadata
// for the struct [CasbPostureFindingTypeValidateResponseEnvelope]
type casbPostureFindingTypeValidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeValidateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeValidateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeValidateResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                     `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingTypeValidateResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingTypeValidateResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingTypeValidateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CasbPostureFindingTypeValidateResponseEnvelopeErrors]
type casbPostureFindingTypeValidateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingTypeValidateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeValidateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeValidateResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                         `json:"pointer"`
	JSON    casbPostureFindingTypeValidateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingTypeValidateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingTypeValidateResponseEnvelopeErrorsSource]
type casbPostureFindingTypeValidateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeValidateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeValidateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeValidateResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                       `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingTypeValidateResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingTypeValidateResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingTypeValidateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CasbPostureFindingTypeValidateResponseEnvelopeMessages]
type casbPostureFindingTypeValidateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingTypeValidateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeValidateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeValidateResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                           `json:"pointer"`
	JSON    casbPostureFindingTypeValidateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingTypeValidateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingTypeValidateResponseEnvelopeMessagesSource]
type casbPostureFindingTypeValidateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingTypeValidateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeValidateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
