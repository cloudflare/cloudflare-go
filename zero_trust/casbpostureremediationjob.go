// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// CasbPostureRemediationJobService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbPostureRemediationJobService] method instead.
type CasbPostureRemediationJobService struct {
	Options []option.RequestOption
}

// NewCasbPostureRemediationJobService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCasbPostureRemediationJobService(opts ...option.RequestOption) (r *CasbPostureRemediationJobService) {
	r = &CasbPostureRemediationJobService{}
	r.Options = opts
	return
}

// Create one or more remediation jobs tied to a specific Cloudflare Account.
func (r *CasbPostureRemediationJobService) New(ctx context.Context, params CasbPostureRemediationJobNewParams, opts ...option.RequestOption) (res *CasbPostureRemediationJobNewResponse, err error) {
	var env CasbPostureRemediationJobNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/remediations/jobs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all remediation jobs tied to a specific Cloudflare Account. Note that
// `cursor` and `page` are mutually exclusive.
func (r *CasbPostureRemediationJobService) List(ctx context.Context, params CasbPostureRemediationJobListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[CasbPostureRemediationJobListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/remediations/jobs", params.AccountID)
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

// List all remediation jobs tied to a specific Cloudflare Account. Note that
// `cursor` and `page` are mutually exclusive.
func (r *CasbPostureRemediationJobService) ListAutoPaging(ctx context.Context, params CasbPostureRemediationJobListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[CasbPostureRemediationJobListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Creates a CSV export for remediation jobs and accepts optional filters in the
// payload.
func (r *CasbPostureRemediationJobService) Export(ctx context.Context, params CasbPostureRemediationJobExportParams, opts ...option.RequestOption) (res *CasbPostureRemediationJobExportResponse, err error) {
	var env CasbPostureRemediationJobExportResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/remediations/jobs/export", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type CasbPostureRemediationJobNewResponse struct {
	// Successfully created remediation jobs.
	Created []CasbPostureRemediationJobNewResponseCreated `json:"created" api:"required"`
	// Failed remediation job creation attempts.
	Failed []CasbPostureRemediationJobNewResponseFailed `json:"failed" api:"required"`
	JSON   casbPostureRemediationJobNewResponseJSON     `json:"-"`
}

// casbPostureRemediationJobNewResponseJSON contains the JSON metadata for the
// struct [CasbPostureRemediationJobNewResponse]
type casbPostureRemediationJobNewResponseJSON struct {
	Created     apijson.Field
	Failed      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureRemediationJobNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobNewResponseJSON) RawJSON() string {
	return r.raw
}

// Information about a remediation job.
type CasbPostureRemediationJobNewResponseCreated struct {
	// Unique identifier for the remediation job.
	ID string `json:"id" api:"required" format:"uuid"`
	// Asset information for a remediation job.
	Asset CasbPostureRemediationJobNewResponseCreatedAsset `json:"asset" api:"required"`
	// When the remediation job was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Encoded finding ID.
	FindingID string `json:"finding_id" api:"required"`
	// ID of the finding instance being remediated.
	FindingInstanceID string `json:"finding_instance_id" api:"required" format:"uuid"`
	// ID of the finding type.
	FindingTypeID string `json:"finding_type_id" api:"required" format:"uuid"`
	// Name of the finding type.
	FindingTypeName string `json:"finding_type_name" api:"required"`
	// Name of the integration.
	IntegrationName string `json:"integration_name" api:"required"`
	// When the remediation job was last updated.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Type of remediation being performed.
	RemediationType string `json:"remediation_type" api:"required"`
	// Status of a remediation job.
	Status CasbPostureRemediationJobNewResponseCreatedStatus `json:"status" api:"required"`
	// Email of the user who triggered the remediation. For account-token actors this
	// is the literal "Account API Token"; for policy actors this is empty.
	TriggeredByUser string `json:"triggered_by_user" api:"required"`
	// Type of actor that triggered the remediation job. Null on legacy rows created
	// before this column was populated.
	TriggeredByActor CasbPostureRemediationJobNewResponseCreatedTriggeredByActor `json:"triggered_by_actor" api:"nullable"`
	// ID of the actor that triggered the job. Meaning depends on triggered_by_actor.
	// Null on legacy rows.
	TriggeredByID string                                          `json:"triggered_by_id" api:"nullable"`
	JSON          casbPostureRemediationJobNewResponseCreatedJSON `json:"-"`
}

// casbPostureRemediationJobNewResponseCreatedJSON contains the JSON metadata for
// the struct [CasbPostureRemediationJobNewResponseCreated]
type casbPostureRemediationJobNewResponseCreatedJSON struct {
	ID                apijson.Field
	Asset             apijson.Field
	CreatedAt         apijson.Field
	FindingID         apijson.Field
	FindingInstanceID apijson.Field
	FindingTypeID     apijson.Field
	FindingTypeName   apijson.Field
	IntegrationName   apijson.Field
	LastUpdated       apijson.Field
	RemediationType   apijson.Field
	Status            apijson.Field
	TriggeredByUser   apijson.Field
	TriggeredByActor  apijson.Field
	TriggeredByID     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbPostureRemediationJobNewResponseCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobNewResponseCreatedJSON) RawJSON() string {
	return r.raw
}

// Asset information for a remediation job.
type CasbPostureRemediationJobNewResponseCreatedAsset struct {
	// Unique identifier for the asset.
	ID string `json:"id" api:"required" format:"uuid"`
	// Category information for a remediation job asset.
	Category CasbPostureRemediationJobNewResponseCreatedAssetCategory `json:"category" api:"required"`
	// External identifier from the source system.
	ExternalID string `json:"external_id" api:"required"`
	// Additional fields associated with the asset.
	Fields []CasbPostureRemediationJobNewResponseCreatedAssetField `json:"fields" api:"required"`
	// Human-readable name of the asset.
	Name string `json:"name" api:"required"`
	// Direct link to the asset.
	Link string                                               `json:"link" api:"nullable" format:"uri"`
	JSON casbPostureRemediationJobNewResponseCreatedAssetJSON `json:"-"`
}

// casbPostureRemediationJobNewResponseCreatedAssetJSON contains the JSON metadata
// for the struct [CasbPostureRemediationJobNewResponseCreatedAsset]
type casbPostureRemediationJobNewResponseCreatedAssetJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	ExternalID  apijson.Field
	Fields      apijson.Field
	Name        apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureRemediationJobNewResponseCreatedAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobNewResponseCreatedAssetJSON) RawJSON() string {
	return r.raw
}

// Category information for a remediation job asset.
type CasbPostureRemediationJobNewResponseCreatedAssetCategory struct {
	// Specific service within the vendor.
	Service string `json:"service" api:"required"`
	// Asset type.
	Type string `json:"type" api:"required"`
	// Display names for vendor types.
	Vendor CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor `json:"vendor" api:"required"`
	JSON   casbPostureRemediationJobNewResponseCreatedAssetCategoryJSON   `json:"-"`
}

// casbPostureRemediationJobNewResponseCreatedAssetCategoryJSON contains the JSON
// metadata for the struct
// [CasbPostureRemediationJobNewResponseCreatedAssetCategory]
type casbPostureRemediationJobNewResponseCreatedAssetCategoryJSON struct {
	Service     apijson.Field
	Type        apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureRemediationJobNewResponseCreatedAssetCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobNewResponseCreatedAssetCategoryJSON) RawJSON() string {
	return r.raw
}

// Display names for vendor types.
type CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor string

const (
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorAws                 CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "AWS"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorAnthropic           CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Anthropic"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorBitbucket           CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Bitbucket"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorBox                 CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Box"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorConfluence          CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Confluence"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorDropbox             CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Dropbox"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorGitHub              CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "GitHub"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorGoogleCloudPlatform CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Google Cloud Platform"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorGoogleWorkspace     CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Google Workspace"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorJira                CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Jira"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorMicrosoft           CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Microsoft"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorMicrosoftInternal   CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Microsoft Internal"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorOkta                CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Okta"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorOpenAI              CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "OpenAI"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorSlack               CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Slack"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorSalesforce          CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Salesforce"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorServiceNow          CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "ServiceNow"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorWorkday             CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Workday"
	CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorZoom                CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor = "Zoom"
)

func (r CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendor) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorAws, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorAnthropic, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorBitbucket, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorBox, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorConfluence, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorDropbox, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorGitHub, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorGoogleCloudPlatform, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorGoogleWorkspace, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorJira, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorMicrosoft, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorMicrosoftInternal, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorOkta, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorOpenAI, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorSlack, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorSalesforce, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorServiceNow, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorWorkday, CasbPostureRemediationJobNewResponseCreatedAssetCategoryVendorZoom:
		return true
	}
	return false
}

// Additional field information for a remediation job asset.
type CasbPostureRemediationJobNewResponseCreatedAssetField struct {
	// Field name.
	Name string `json:"name" api:"required"`
	// Field value (can be string, number, or boolean).
	Value CasbPostureRemediationJobNewResponseCreatedAssetFieldsValueUnion `json:"value" api:"required"`
	// Optional link associated with the field.
	Link string                                                    `json:"link" api:"nullable" format:"uri"`
	JSON casbPostureRemediationJobNewResponseCreatedAssetFieldJSON `json:"-"`
}

// casbPostureRemediationJobNewResponseCreatedAssetFieldJSON contains the JSON
// metadata for the struct [CasbPostureRemediationJobNewResponseCreatedAssetField]
type casbPostureRemediationJobNewResponseCreatedAssetFieldJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureRemediationJobNewResponseCreatedAssetField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobNewResponseCreatedAssetFieldJSON) RawJSON() string {
	return r.raw
}

// Field value (can be string, number, or boolean).
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type CasbPostureRemediationJobNewResponseCreatedAssetFieldsValueUnion interface {
	ImplementsCasbPostureRemediationJobNewResponseCreatedAssetFieldsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CasbPostureRemediationJobNewResponseCreatedAssetFieldsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// Status of a remediation job.
type CasbPostureRemediationJobNewResponseCreatedStatus string

const (
	CasbPostureRemediationJobNewResponseCreatedStatusPending    CasbPostureRemediationJobNewResponseCreatedStatus = "pending"
	CasbPostureRemediationJobNewResponseCreatedStatusProcessing CasbPostureRemediationJobNewResponseCreatedStatus = "processing"
	CasbPostureRemediationJobNewResponseCreatedStatusCompleted  CasbPostureRemediationJobNewResponseCreatedStatus = "completed"
	CasbPostureRemediationJobNewResponseCreatedStatusFailed     CasbPostureRemediationJobNewResponseCreatedStatus = "failed"
	CasbPostureRemediationJobNewResponseCreatedStatusValidating CasbPostureRemediationJobNewResponseCreatedStatus = "validating"
)

func (r CasbPostureRemediationJobNewResponseCreatedStatus) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobNewResponseCreatedStatusPending, CasbPostureRemediationJobNewResponseCreatedStatusProcessing, CasbPostureRemediationJobNewResponseCreatedStatusCompleted, CasbPostureRemediationJobNewResponseCreatedStatusFailed, CasbPostureRemediationJobNewResponseCreatedStatusValidating:
		return true
	}
	return false
}

// Type of actor that triggered the remediation job. Null on legacy rows created
// before this column was populated.
type CasbPostureRemediationJobNewResponseCreatedTriggeredByActor string

const (
	CasbPostureRemediationJobNewResponseCreatedTriggeredByActorUser         CasbPostureRemediationJobNewResponseCreatedTriggeredByActor = "user"
	CasbPostureRemediationJobNewResponseCreatedTriggeredByActorAccountToken CasbPostureRemediationJobNewResponseCreatedTriggeredByActor = "account_token"
)

func (r CasbPostureRemediationJobNewResponseCreatedTriggeredByActor) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobNewResponseCreatedTriggeredByActorUser, CasbPostureRemediationJobNewResponseCreatedTriggeredByActorAccountToken:
		return true
	}
	return false
}

// Information about a failed remediation job creation.
type CasbPostureRemediationJobNewResponseFailed struct {
	// Error message describing the failure.
	Error string `json:"error" api:"required"`
	// ID of the finding instance that failed to create a remediation job.
	FindingInstanceID string                                         `json:"finding_instance_id" api:"required" format:"uuid"`
	JSON              casbPostureRemediationJobNewResponseFailedJSON `json:"-"`
}

// casbPostureRemediationJobNewResponseFailedJSON contains the JSON metadata for
// the struct [CasbPostureRemediationJobNewResponseFailed]
type casbPostureRemediationJobNewResponseFailedJSON struct {
	Error             apijson.Field
	FindingInstanceID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbPostureRemediationJobNewResponseFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobNewResponseFailedJSON) RawJSON() string {
	return r.raw
}

// Information about a remediation job.
type CasbPostureRemediationJobListResponse struct {
	// Unique identifier for the remediation job.
	ID string `json:"id" api:"required" format:"uuid"`
	// Asset information for a remediation job.
	Asset CasbPostureRemediationJobListResponseAsset `json:"asset" api:"required"`
	// When the remediation job was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Encoded finding ID.
	FindingID string `json:"finding_id" api:"required"`
	// ID of the finding instance being remediated.
	FindingInstanceID string `json:"finding_instance_id" api:"required" format:"uuid"`
	// ID of the finding type.
	FindingTypeID string `json:"finding_type_id" api:"required" format:"uuid"`
	// Name of the finding type.
	FindingTypeName string `json:"finding_type_name" api:"required"`
	// Name of the integration.
	IntegrationName string `json:"integration_name" api:"required"`
	// When the remediation job was last updated.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Type of remediation being performed.
	RemediationType string `json:"remediation_type" api:"required"`
	// Status of a remediation job.
	Status CasbPostureRemediationJobListResponseStatus `json:"status" api:"required"`
	// Email of the user who triggered the remediation. For account-token actors this
	// is the literal "Account API Token"; for policy actors this is empty.
	TriggeredByUser string `json:"triggered_by_user" api:"required"`
	// Type of actor that triggered the remediation job. Null on legacy rows created
	// before this column was populated.
	TriggeredByActor CasbPostureRemediationJobListResponseTriggeredByActor `json:"triggered_by_actor" api:"nullable"`
	// ID of the actor that triggered the job. Meaning depends on triggered_by_actor.
	// Null on legacy rows.
	TriggeredByID string                                    `json:"triggered_by_id" api:"nullable"`
	JSON          casbPostureRemediationJobListResponseJSON `json:"-"`
}

// casbPostureRemediationJobListResponseJSON contains the JSON metadata for the
// struct [CasbPostureRemediationJobListResponse]
type casbPostureRemediationJobListResponseJSON struct {
	ID                apijson.Field
	Asset             apijson.Field
	CreatedAt         apijson.Field
	FindingID         apijson.Field
	FindingInstanceID apijson.Field
	FindingTypeID     apijson.Field
	FindingTypeName   apijson.Field
	IntegrationName   apijson.Field
	LastUpdated       apijson.Field
	RemediationType   apijson.Field
	Status            apijson.Field
	TriggeredByUser   apijson.Field
	TriggeredByActor  apijson.Field
	TriggeredByID     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbPostureRemediationJobListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobListResponseJSON) RawJSON() string {
	return r.raw
}

// Asset information for a remediation job.
type CasbPostureRemediationJobListResponseAsset struct {
	// Unique identifier for the asset.
	ID string `json:"id" api:"required" format:"uuid"`
	// Category information for a remediation job asset.
	Category CasbPostureRemediationJobListResponseAssetCategory `json:"category" api:"required"`
	// External identifier from the source system.
	ExternalID string `json:"external_id" api:"required"`
	// Additional fields associated with the asset.
	Fields []CasbPostureRemediationJobListResponseAssetField `json:"fields" api:"required"`
	// Human-readable name of the asset.
	Name string `json:"name" api:"required"`
	// Direct link to the asset.
	Link string                                         `json:"link" api:"nullable" format:"uri"`
	JSON casbPostureRemediationJobListResponseAssetJSON `json:"-"`
}

// casbPostureRemediationJobListResponseAssetJSON contains the JSON metadata for
// the struct [CasbPostureRemediationJobListResponseAsset]
type casbPostureRemediationJobListResponseAssetJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	ExternalID  apijson.Field
	Fields      apijson.Field
	Name        apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureRemediationJobListResponseAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobListResponseAssetJSON) RawJSON() string {
	return r.raw
}

// Category information for a remediation job asset.
type CasbPostureRemediationJobListResponseAssetCategory struct {
	// Specific service within the vendor.
	Service string `json:"service" api:"required"`
	// Asset type.
	Type string `json:"type" api:"required"`
	// Display names for vendor types.
	Vendor CasbPostureRemediationJobListResponseAssetCategoryVendor `json:"vendor" api:"required"`
	JSON   casbPostureRemediationJobListResponseAssetCategoryJSON   `json:"-"`
}

// casbPostureRemediationJobListResponseAssetCategoryJSON contains the JSON
// metadata for the struct [CasbPostureRemediationJobListResponseAssetCategory]
type casbPostureRemediationJobListResponseAssetCategoryJSON struct {
	Service     apijson.Field
	Type        apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureRemediationJobListResponseAssetCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobListResponseAssetCategoryJSON) RawJSON() string {
	return r.raw
}

// Display names for vendor types.
type CasbPostureRemediationJobListResponseAssetCategoryVendor string

const (
	CasbPostureRemediationJobListResponseAssetCategoryVendorAws                 CasbPostureRemediationJobListResponseAssetCategoryVendor = "AWS"
	CasbPostureRemediationJobListResponseAssetCategoryVendorAnthropic           CasbPostureRemediationJobListResponseAssetCategoryVendor = "Anthropic"
	CasbPostureRemediationJobListResponseAssetCategoryVendorBitbucket           CasbPostureRemediationJobListResponseAssetCategoryVendor = "Bitbucket"
	CasbPostureRemediationJobListResponseAssetCategoryVendorBox                 CasbPostureRemediationJobListResponseAssetCategoryVendor = "Box"
	CasbPostureRemediationJobListResponseAssetCategoryVendorConfluence          CasbPostureRemediationJobListResponseAssetCategoryVendor = "Confluence"
	CasbPostureRemediationJobListResponseAssetCategoryVendorDropbox             CasbPostureRemediationJobListResponseAssetCategoryVendor = "Dropbox"
	CasbPostureRemediationJobListResponseAssetCategoryVendorGitHub              CasbPostureRemediationJobListResponseAssetCategoryVendor = "GitHub"
	CasbPostureRemediationJobListResponseAssetCategoryVendorGoogleCloudPlatform CasbPostureRemediationJobListResponseAssetCategoryVendor = "Google Cloud Platform"
	CasbPostureRemediationJobListResponseAssetCategoryVendorGoogleWorkspace     CasbPostureRemediationJobListResponseAssetCategoryVendor = "Google Workspace"
	CasbPostureRemediationJobListResponseAssetCategoryVendorJira                CasbPostureRemediationJobListResponseAssetCategoryVendor = "Jira"
	CasbPostureRemediationJobListResponseAssetCategoryVendorMicrosoft           CasbPostureRemediationJobListResponseAssetCategoryVendor = "Microsoft"
	CasbPostureRemediationJobListResponseAssetCategoryVendorMicrosoftInternal   CasbPostureRemediationJobListResponseAssetCategoryVendor = "Microsoft Internal"
	CasbPostureRemediationJobListResponseAssetCategoryVendorOkta                CasbPostureRemediationJobListResponseAssetCategoryVendor = "Okta"
	CasbPostureRemediationJobListResponseAssetCategoryVendorOpenAI              CasbPostureRemediationJobListResponseAssetCategoryVendor = "OpenAI"
	CasbPostureRemediationJobListResponseAssetCategoryVendorSlack               CasbPostureRemediationJobListResponseAssetCategoryVendor = "Slack"
	CasbPostureRemediationJobListResponseAssetCategoryVendorSalesforce          CasbPostureRemediationJobListResponseAssetCategoryVendor = "Salesforce"
	CasbPostureRemediationJobListResponseAssetCategoryVendorServiceNow          CasbPostureRemediationJobListResponseAssetCategoryVendor = "ServiceNow"
	CasbPostureRemediationJobListResponseAssetCategoryVendorWorkday             CasbPostureRemediationJobListResponseAssetCategoryVendor = "Workday"
	CasbPostureRemediationJobListResponseAssetCategoryVendorZoom                CasbPostureRemediationJobListResponseAssetCategoryVendor = "Zoom"
)

func (r CasbPostureRemediationJobListResponseAssetCategoryVendor) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobListResponseAssetCategoryVendorAws, CasbPostureRemediationJobListResponseAssetCategoryVendorAnthropic, CasbPostureRemediationJobListResponseAssetCategoryVendorBitbucket, CasbPostureRemediationJobListResponseAssetCategoryVendorBox, CasbPostureRemediationJobListResponseAssetCategoryVendorConfluence, CasbPostureRemediationJobListResponseAssetCategoryVendorDropbox, CasbPostureRemediationJobListResponseAssetCategoryVendorGitHub, CasbPostureRemediationJobListResponseAssetCategoryVendorGoogleCloudPlatform, CasbPostureRemediationJobListResponseAssetCategoryVendorGoogleWorkspace, CasbPostureRemediationJobListResponseAssetCategoryVendorJira, CasbPostureRemediationJobListResponseAssetCategoryVendorMicrosoft, CasbPostureRemediationJobListResponseAssetCategoryVendorMicrosoftInternal, CasbPostureRemediationJobListResponseAssetCategoryVendorOkta, CasbPostureRemediationJobListResponseAssetCategoryVendorOpenAI, CasbPostureRemediationJobListResponseAssetCategoryVendorSlack, CasbPostureRemediationJobListResponseAssetCategoryVendorSalesforce, CasbPostureRemediationJobListResponseAssetCategoryVendorServiceNow, CasbPostureRemediationJobListResponseAssetCategoryVendorWorkday, CasbPostureRemediationJobListResponseAssetCategoryVendorZoom:
		return true
	}
	return false
}

// Additional field information for a remediation job asset.
type CasbPostureRemediationJobListResponseAssetField struct {
	// Field name.
	Name string `json:"name" api:"required"`
	// Field value (can be string, number, or boolean).
	Value CasbPostureRemediationJobListResponseAssetFieldsValueUnion `json:"value" api:"required"`
	// Optional link associated with the field.
	Link string                                              `json:"link" api:"nullable" format:"uri"`
	JSON casbPostureRemediationJobListResponseAssetFieldJSON `json:"-"`
}

// casbPostureRemediationJobListResponseAssetFieldJSON contains the JSON metadata
// for the struct [CasbPostureRemediationJobListResponseAssetField]
type casbPostureRemediationJobListResponseAssetFieldJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureRemediationJobListResponseAssetField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobListResponseAssetFieldJSON) RawJSON() string {
	return r.raw
}

// Field value (can be string, number, or boolean).
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type CasbPostureRemediationJobListResponseAssetFieldsValueUnion interface {
	ImplementsCasbPostureRemediationJobListResponseAssetFieldsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CasbPostureRemediationJobListResponseAssetFieldsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// Status of a remediation job.
type CasbPostureRemediationJobListResponseStatus string

const (
	CasbPostureRemediationJobListResponseStatusPending    CasbPostureRemediationJobListResponseStatus = "pending"
	CasbPostureRemediationJobListResponseStatusProcessing CasbPostureRemediationJobListResponseStatus = "processing"
	CasbPostureRemediationJobListResponseStatusCompleted  CasbPostureRemediationJobListResponseStatus = "completed"
	CasbPostureRemediationJobListResponseStatusFailed     CasbPostureRemediationJobListResponseStatus = "failed"
	CasbPostureRemediationJobListResponseStatusValidating CasbPostureRemediationJobListResponseStatus = "validating"
)

func (r CasbPostureRemediationJobListResponseStatus) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobListResponseStatusPending, CasbPostureRemediationJobListResponseStatusProcessing, CasbPostureRemediationJobListResponseStatusCompleted, CasbPostureRemediationJobListResponseStatusFailed, CasbPostureRemediationJobListResponseStatusValidating:
		return true
	}
	return false
}

// Type of actor that triggered the remediation job. Null on legacy rows created
// before this column was populated.
type CasbPostureRemediationJobListResponseTriggeredByActor string

const (
	CasbPostureRemediationJobListResponseTriggeredByActorUser         CasbPostureRemediationJobListResponseTriggeredByActor = "user"
	CasbPostureRemediationJobListResponseTriggeredByActorAccountToken CasbPostureRemediationJobListResponseTriggeredByActor = "account_token"
)

func (r CasbPostureRemediationJobListResponseTriggeredByActor) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobListResponseTriggeredByActorUser, CasbPostureRemediationJobListResponseTriggeredByActorAccountToken:
		return true
	}
	return false
}

// Information about an export job.
type CasbPostureRemediationJobExportResponse struct {
	// Unique identifier for the export job.
	ID string `json:"id" api:"required" format:"uuid"`
	// Status of an export job.
	Status CasbPostureRemediationJobExportResponseStatus `json:"status" api:"required"`
	// Type of export job.
	Type CasbPostureRemediationJobExportResponseType `json:"type" api:"required"`
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
	FilePath string                                      `json:"file_path" api:"nullable"`
	JSON     casbPostureRemediationJobExportResponseJSON `json:"-"`
}

// casbPostureRemediationJobExportResponseJSON contains the JSON metadata for the
// struct [CasbPostureRemediationJobExportResponse]
type casbPostureRemediationJobExportResponseJSON struct {
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

func (r *CasbPostureRemediationJobExportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobExportResponseJSON) RawJSON() string {
	return r.raw
}

// Status of an export job.
type CasbPostureRemediationJobExportResponseStatus string

const (
	CasbPostureRemediationJobExportResponseStatusPending     CasbPostureRemediationJobExportResponseStatus = "Pending"
	CasbPostureRemediationJobExportResponseStatusSuccess     CasbPostureRemediationJobExportResponseStatus = "Success"
	CasbPostureRemediationJobExportResponseStatusFailure     CasbPostureRemediationJobExportResponseStatus = "Failure"
	CasbPostureRemediationJobExportResponseStatusRescheduled CasbPostureRemediationJobExportResponseStatus = "Rescheduled"
	CasbPostureRemediationJobExportResponseStatusInProgress  CasbPostureRemediationJobExportResponseStatus = "In-Progress"
)

func (r CasbPostureRemediationJobExportResponseStatus) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobExportResponseStatusPending, CasbPostureRemediationJobExportResponseStatusSuccess, CasbPostureRemediationJobExportResponseStatusFailure, CasbPostureRemediationJobExportResponseStatusRescheduled, CasbPostureRemediationJobExportResponseStatusInProgress:
		return true
	}
	return false
}

// Type of export job.
type CasbPostureRemediationJobExportResponseType string

const (
	CasbPostureRemediationJobExportResponseTypeFinding         CasbPostureRemediationJobExportResponseType = "finding"
	CasbPostureRemediationJobExportResponseTypeFindingInstance CasbPostureRemediationJobExportResponseType = "findingInstance"
	CasbPostureRemediationJobExportResponseTypeContent         CasbPostureRemediationJobExportResponseType = "content"
	CasbPostureRemediationJobExportResponseTypeRemediationJob  CasbPostureRemediationJobExportResponseType = "remediationJob"
)

func (r CasbPostureRemediationJobExportResponseType) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobExportResponseTypeFinding, CasbPostureRemediationJobExportResponseTypeFindingInstance, CasbPostureRemediationJobExportResponseTypeContent, CasbPostureRemediationJobExportResponseTypeRemediationJob:
		return true
	}
	return false
}

type CasbPostureRemediationJobNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// UUIDs identifying Finding Instances.
	FindingInstanceIDs param.Field[[]string] `json:"finding_instance_ids" api:"required" format:"uuid"`
	// A UUID identifying this Remediation Type.
	RemediationTypeID param.Field[string] `json:"remediation_type_id" api:"required" format:"uuid"`
}

func (r CasbPostureRemediationJobNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response for remediation job creation requests.
type CasbPostureRemediationJobNewResponseEnvelope struct {
	// Array of error messages.
	Errors []interface{} `json:"errors" api:"required"`
	// Array of informational messages.
	Messages []interface{}                        `json:"messages" api:"required"`
	Result   CasbPostureRemediationJobNewResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                             `json:"success" api:"required"`
	JSON    casbPostureRemediationJobNewResponseEnvelopeJSON `json:"-"`
}

// casbPostureRemediationJobNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [CasbPostureRemediationJobNewResponseEnvelope]
type casbPostureRemediationJobNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureRemediationJobNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureRemediationJobListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A cursor for pagination.
	Cursor param.Field[string] `query:"cursor"`
	// Direction to order results.
	Direction param.Field[CasbPostureRemediationJobListParamsDirection] `query:"direction"`
	// Filter by an integration ID
	IntegrationID param.Field[string] `query:"integration_id" format:"uuid"`
	// Filter to view remediations updated on or before the max updated datetime. Can
	// be a date-time in ISO 8601 format or an epoch timestamp.
	MaxUpdatedAt param.Field[time.Time] `query:"max_updated_at" format:"date-time"`
	// Filter to view remediations updated on or after the min updated datetime. Can be
	// a date-time in ISO 8601 format or an epoch timestamp.
	MinUpdatedAt param.Field[time.Time] `query:"min_updated_at" format:"date-time"`
	// An optional param to sort the results by the given field.
	Order param.Field[CasbPostureRemediationJobListParamsOrder] `query:"order"`
	// A page number within the paginated result set.
	Page param.Field[int64] `query:"page"`
	// Number of results to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// A search term.
	Search param.Field[string] `query:"search"`
	// Filter to view remediations with the given status.
	Status param.Field[CasbPostureRemediationJobListParamsStatus] `query:"status"`
	// Filter remediations by what kind of actor triggered them. Supports multiple
	// comma-separated values.
	TriggeredByActor param.Field[[]CasbPostureRemediationJobListParamsTriggeredByActor] `query:"triggered_by_actor"`
}

// URLQuery serializes [CasbPostureRemediationJobListParams]'s query parameters as
// `url.Values`.
func (r CasbPostureRemediationJobListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type CasbPostureRemediationJobListParamsDirection string

const (
	CasbPostureRemediationJobListParamsDirectionAsc  CasbPostureRemediationJobListParamsDirection = "asc"
	CasbPostureRemediationJobListParamsDirectionDesc CasbPostureRemediationJobListParamsDirection = "desc"
)

func (r CasbPostureRemediationJobListParamsDirection) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobListParamsDirectionAsc, CasbPostureRemediationJobListParamsDirectionDesc:
		return true
	}
	return false
}

// An optional param to sort the results by the given field.
type CasbPostureRemediationJobListParamsOrder string

const (
	CasbPostureRemediationJobListParamsOrderCreatedAt       CasbPostureRemediationJobListParamsOrder = "created_at"
	CasbPostureRemediationJobListParamsOrderAfflictionDate  CasbPostureRemediationJobListParamsOrder = "affliction_date"
	CasbPostureRemediationJobListParamsOrderIntegrationName CasbPostureRemediationJobListParamsOrder = "integration_name"
	CasbPostureRemediationJobListParamsOrderStatus          CasbPostureRemediationJobListParamsOrder = "status"
	CasbPostureRemediationJobListParamsOrderLastUpdatedAt   CasbPostureRemediationJobListParamsOrder = "last_updated_at"
	CasbPostureRemediationJobListParamsOrderAssetName       CasbPostureRemediationJobListParamsOrder = "asset_name"
	CasbPostureRemediationJobListParamsOrderFindingTypeName CasbPostureRemediationJobListParamsOrder = "finding_type_name"
)

func (r CasbPostureRemediationJobListParamsOrder) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobListParamsOrderCreatedAt, CasbPostureRemediationJobListParamsOrderAfflictionDate, CasbPostureRemediationJobListParamsOrderIntegrationName, CasbPostureRemediationJobListParamsOrderStatus, CasbPostureRemediationJobListParamsOrderLastUpdatedAt, CasbPostureRemediationJobListParamsOrderAssetName, CasbPostureRemediationJobListParamsOrderFindingTypeName:
		return true
	}
	return false
}

// Filter to view remediations with the given status.
type CasbPostureRemediationJobListParamsStatus string

const (
	CasbPostureRemediationJobListParamsStatusPending    CasbPostureRemediationJobListParamsStatus = "pending"
	CasbPostureRemediationJobListParamsStatusProcessing CasbPostureRemediationJobListParamsStatus = "processing"
	CasbPostureRemediationJobListParamsStatusCompleted  CasbPostureRemediationJobListParamsStatus = "completed"
	CasbPostureRemediationJobListParamsStatusFailed     CasbPostureRemediationJobListParamsStatus = "failed"
	CasbPostureRemediationJobListParamsStatusValidating CasbPostureRemediationJobListParamsStatus = "validating"
)

func (r CasbPostureRemediationJobListParamsStatus) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobListParamsStatusPending, CasbPostureRemediationJobListParamsStatusProcessing, CasbPostureRemediationJobListParamsStatusCompleted, CasbPostureRemediationJobListParamsStatusFailed, CasbPostureRemediationJobListParamsStatusValidating:
		return true
	}
	return false
}

// Type of actor that triggered the remediation job.
type CasbPostureRemediationJobListParamsTriggeredByActor string

const (
	CasbPostureRemediationJobListParamsTriggeredByActorUser         CasbPostureRemediationJobListParamsTriggeredByActor = "user"
	CasbPostureRemediationJobListParamsTriggeredByActorAccountToken CasbPostureRemediationJobListParamsTriggeredByActor = "account_token"
)

func (r CasbPostureRemediationJobListParamsTriggeredByActor) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobListParamsTriggeredByActorUser, CasbPostureRemediationJobListParamsTriggeredByActorAccountToken:
		return true
	}
	return false
}

type CasbPostureRemediationJobExportParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter by multiple integration IDs.
	IntegrationID param.Field[[]string] `json:"integration_id" format:"uuid"`
	// Filter to view remediation jobs updated on or before this datetime. Can be a
	// date-time in ISO 8601 format or an epoch timestamp.
	MaxUpdatedAt param.Field[time.Time] `json:"max_updated_at" format:"date-time"`
	// Filter to view remediation jobs updated on or after this datetime. Can be a
	// date-time in ISO 8601 format or an epoch timestamp.
	MinUpdatedAt param.Field[time.Time] `json:"min_updated_at" format:"date-time"`
	// Ordering specifications for the export.
	Orders param.Field[[]CasbPostureRemediationJobExportParamsOrder] `json:"orders"`
	// A search term.
	Search param.Field[string] `json:"search"`
	// Filter by remediation job status.
	Status param.Field[[]CasbPostureRemediationJobExportParamsStatus] `json:"status"`
}

func (r CasbPostureRemediationJobExportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Order specification for remediation jobs exports.
type CasbPostureRemediationJobExportParamsOrder struct {
	// Sort direction.
	Direction param.Field[CasbPostureRemediationJobExportParamsOrdersDirection] `json:"direction" api:"required"`
	// Which field to use when ordering the remediation jobs.
	Name param.Field[CasbPostureRemediationJobExportParamsOrdersName] `json:"name" api:"required"`
}

func (r CasbPostureRemediationJobExportParamsOrder) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort direction.
type CasbPostureRemediationJobExportParamsOrdersDirection string

const (
	CasbPostureRemediationJobExportParamsOrdersDirectionAsc  CasbPostureRemediationJobExportParamsOrdersDirection = "asc"
	CasbPostureRemediationJobExportParamsOrdersDirectionDesc CasbPostureRemediationJobExportParamsOrdersDirection = "desc"
)

func (r CasbPostureRemediationJobExportParamsOrdersDirection) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobExportParamsOrdersDirectionAsc, CasbPostureRemediationJobExportParamsOrdersDirectionDesc:
		return true
	}
	return false
}

// Which field to use when ordering the remediation jobs.
type CasbPostureRemediationJobExportParamsOrdersName string

const (
	CasbPostureRemediationJobExportParamsOrdersNameAssetName       CasbPostureRemediationJobExportParamsOrdersName = "asset_name"
	CasbPostureRemediationJobExportParamsOrdersNameFindingTypeName CasbPostureRemediationJobExportParamsOrdersName = "finding_type_name"
	CasbPostureRemediationJobExportParamsOrdersNameIntegrationName CasbPostureRemediationJobExportParamsOrdersName = "integration_name"
	CasbPostureRemediationJobExportParamsOrdersNameStatus          CasbPostureRemediationJobExportParamsOrdersName = "status"
	CasbPostureRemediationJobExportParamsOrdersNameLastUpdatedAt   CasbPostureRemediationJobExportParamsOrdersName = "last_updated_at"
	CasbPostureRemediationJobExportParamsOrdersNameAfflictionDate  CasbPostureRemediationJobExportParamsOrdersName = "affliction_date"
)

func (r CasbPostureRemediationJobExportParamsOrdersName) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobExportParamsOrdersNameAssetName, CasbPostureRemediationJobExportParamsOrdersNameFindingTypeName, CasbPostureRemediationJobExportParamsOrdersNameIntegrationName, CasbPostureRemediationJobExportParamsOrdersNameStatus, CasbPostureRemediationJobExportParamsOrdersNameLastUpdatedAt, CasbPostureRemediationJobExportParamsOrdersNameAfflictionDate:
		return true
	}
	return false
}

// Status of a remediation job.
type CasbPostureRemediationJobExportParamsStatus string

const (
	CasbPostureRemediationJobExportParamsStatusPending    CasbPostureRemediationJobExportParamsStatus = "pending"
	CasbPostureRemediationJobExportParamsStatusProcessing CasbPostureRemediationJobExportParamsStatus = "processing"
	CasbPostureRemediationJobExportParamsStatusCompleted  CasbPostureRemediationJobExportParamsStatus = "completed"
	CasbPostureRemediationJobExportParamsStatusFailed     CasbPostureRemediationJobExportParamsStatus = "failed"
	CasbPostureRemediationJobExportParamsStatusValidating CasbPostureRemediationJobExportParamsStatus = "validating"
)

func (r CasbPostureRemediationJobExportParamsStatus) IsKnown() bool {
	switch r {
	case CasbPostureRemediationJobExportParamsStatusPending, CasbPostureRemediationJobExportParamsStatusProcessing, CasbPostureRemediationJobExportParamsStatusCompleted, CasbPostureRemediationJobExportParamsStatusFailed, CasbPostureRemediationJobExportParamsStatusValidating:
		return true
	}
	return false
}

// Common response structure for all API endpoints.
type CasbPostureRemediationJobExportResponseEnvelope struct {
	Errors   []CasbPostureRemediationJobExportResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureRemediationJobExportResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Information about an export job.
	Result CasbPostureRemediationJobExportResponse             `json:"result"`
	JSON   casbPostureRemediationJobExportResponseEnvelopeJSON `json:"-"`
}

// casbPostureRemediationJobExportResponseEnvelopeJSON contains the JSON metadata
// for the struct [CasbPostureRemediationJobExportResponseEnvelope]
type casbPostureRemediationJobExportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureRemediationJobExportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobExportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureRemediationJobExportResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                      `json:"documentation_url" format:"uri"`
	Source           CasbPostureRemediationJobExportResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureRemediationJobExportResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureRemediationJobExportResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CasbPostureRemediationJobExportResponseEnvelopeErrors]
type casbPostureRemediationJobExportResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureRemediationJobExportResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobExportResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureRemediationJobExportResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                          `json:"pointer"`
	JSON    casbPostureRemediationJobExportResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureRemediationJobExportResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [CasbPostureRemediationJobExportResponseEnvelopeErrorsSource]
type casbPostureRemediationJobExportResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureRemediationJobExportResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobExportResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureRemediationJobExportResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                        `json:"documentation_url" format:"uri"`
	Source           CasbPostureRemediationJobExportResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureRemediationJobExportResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureRemediationJobExportResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CasbPostureRemediationJobExportResponseEnvelopeMessages]
type casbPostureRemediationJobExportResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureRemediationJobExportResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobExportResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureRemediationJobExportResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                            `json:"pointer"`
	JSON    casbPostureRemediationJobExportResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureRemediationJobExportResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [CasbPostureRemediationJobExportResponseEnvelopeMessagesSource]
type casbPostureRemediationJobExportResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureRemediationJobExportResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureRemediationJobExportResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
