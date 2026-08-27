// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

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

// Basic finding type information.
type CasbPostureFindingTypeListResponse struct {
	// The unique identifier of the finding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Category information for a finding.
	Category CasbPostureFindingTypeListResponseCategory `json:"category" api:"required"`
	// The name of the finding.
	Name string `json:"name" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingTypeListResponseSeverity `json:"severity" api:"required"`
	// The SaaS/Cloud vendor of the platform with which the finding is associated.
	Vendor string                                 `json:"vendor" api:"required"`
	JSON   casbPostureFindingTypeListResponseJSON `json:"-"`
}

// casbPostureFindingTypeListResponseJSON contains the JSON metadata for the struct
// [CasbPostureFindingTypeListResponse]
type casbPostureFindingTypeListResponseJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	Name        apijson.Field
	Severity    apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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

// The severity level of a finding.
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

// Basic finding type information.
type CasbPostureFindingTypeGetResponse struct {
	// The unique identifier of the finding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Category information for a finding.
	Category CasbPostureFindingTypeGetResponseCategory `json:"category" api:"required"`
	// The name of the finding.
	Name string `json:"name" api:"required"`
	// The severity level of a finding.
	Severity CasbPostureFindingTypeGetResponseSeverity `json:"severity" api:"required"`
	// The SaaS/Cloud vendor of the platform with which the finding is associated.
	Vendor string                                `json:"vendor" api:"required"`
	JSON   casbPostureFindingTypeGetResponseJSON `json:"-"`
}

// casbPostureFindingTypeGetResponseJSON contains the JSON metadata for the struct
// [CasbPostureFindingTypeGetResponse]
type casbPostureFindingTypeGetResponseJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	Name        apijson.Field
	Severity    apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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

// The severity level of a finding.
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

type CasbPostureFindingTypeListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A page number within the paginated result set.
	Page param.Field[int64] `query:"page"`
	// Number of results to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter finding types by name or ID (case-insensitive substring match).
	Search param.Field[string] `query:"search"`
	// Filter finding types by vendor. Supports multiple comma-separated values.
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

// Supported vendor types for integrations.
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
	CasbPostureFindingTypeListParamsVendorMicrosoftInternal   CasbPostureFindingTypeListParamsVendor = "MICROSOFT_INTERNAL"
	CasbPostureFindingTypeListParamsVendorOpenAI              CasbPostureFindingTypeListParamsVendor = "OPENAI"
	CasbPostureFindingTypeListParamsVendorSalesforce          CasbPostureFindingTypeListParamsVendor = "SALESFORCE"
	CasbPostureFindingTypeListParamsVendorServicenow          CasbPostureFindingTypeListParamsVendor = "SERVICENOW"
	CasbPostureFindingTypeListParamsVendorSlack               CasbPostureFindingTypeListParamsVendor = "SLACK"
)

func (r CasbPostureFindingTypeListParamsVendor) IsKnown() bool {
	switch r {
	case CasbPostureFindingTypeListParamsVendorAnthropic, CasbPostureFindingTypeListParamsVendorAws, CasbPostureFindingTypeListParamsVendorBitbucket, CasbPostureFindingTypeListParamsVendorBox, CasbPostureFindingTypeListParamsVendorConfluence, CasbPostureFindingTypeListParamsVendorDropbox, CasbPostureFindingTypeListParamsVendorGitHub, CasbPostureFindingTypeListParamsVendorGoogleCloudPlatform, CasbPostureFindingTypeListParamsVendorGoogleWorkspace, CasbPostureFindingTypeListParamsVendorJira, CasbPostureFindingTypeListParamsVendorMicrosoft, CasbPostureFindingTypeListParamsVendorMicrosoftInternal, CasbPostureFindingTypeListParamsVendorOpenAI, CasbPostureFindingTypeListParamsVendorSalesforce, CasbPostureFindingTypeListParamsVendorServicenow, CasbPostureFindingTypeListParamsVendorSlack:
		return true
	}
	return false
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
