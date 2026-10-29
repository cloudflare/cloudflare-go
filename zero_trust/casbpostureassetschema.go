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

// CasbPostureAssetSchemaService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbPostureAssetSchemaService] method instead.
type CasbPostureAssetSchemaService struct {
	Options []option.RequestOption
}

// NewCasbPostureAssetSchemaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCasbPostureAssetSchemaService(opts ...option.RequestOption) (r *CasbPostureAssetSchemaService) {
	r = &CasbPostureAssetSchemaService{}
	r.Options = opts
	return
}

// Returns the catalog of asset classes available when authoring custom findings,
// grouped by vendor. Optionally filter to a single vendor with the `vendor` query
// parameter.
func (r *CasbPostureAssetSchemaService) List(ctx context.Context, params CasbPostureAssetSchemaListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CasbPostureAssetSchemaListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/asset_schemas", params.AccountID)
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

// Returns the catalog of asset classes available when authoring custom findings,
// grouped by vendor. Optionally filter to a single vendor with the `vendor` query
// parameter.
func (r *CasbPostureAssetSchemaService) ListAutoPaging(ctx context.Context, params CasbPostureAssetSchemaListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CasbPostureAssetSchemaListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Returns the fields and related attributes a custom finding type may reference
// for the given asset class (JSON Schema). Used by the custom finding types
// builder.
func (r *CasbPostureAssetSchemaService) Get(ctx context.Context, assetClass string, query CasbPostureAssetSchemaGetParams, opts ...option.RequestOption) (res *CasbPostureAssetSchemaGetResponse, err error) {
	var env CasbPostureAssetSchemaGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if assetClass == "" {
		err = errors.New("missing required asset_class parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/asset_schemas/%s", query.AccountID, assetClass)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type CasbPostureAssetSchemaListResponse struct {
	// Asset classes belonging to this vendor, sorted by asset_class.
	AssetClasses []CasbPostureAssetSchemaListResponseAssetClass `json:"asset_classes" api:"required"`
	// Canonical vendor enum value, e.g. GOOGLE_WORKSPACE.
	Vendor string                                 `json:"vendor" api:"required"`
	JSON   casbPostureAssetSchemaListResponseJSON `json:"-"`
}

// casbPostureAssetSchemaListResponseJSON contains the JSON metadata for the struct
// [CasbPostureAssetSchemaListResponse]
type casbPostureAssetSchemaListResponseJSON struct {
	AssetClasses apijson.Field
	Vendor       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CasbPostureAssetSchemaListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureAssetSchemaListResponseJSON) RawJSON() string {
	return r.raw
}

type CasbPostureAssetSchemaListResponseAssetClass struct {
	// Asset class identifier (PascalCase), e.g. AnthropicUser.
	AssetClass string                                           `json:"asset_class" api:"required"`
	JSON       casbPostureAssetSchemaListResponseAssetClassJSON `json:"-"`
}

// casbPostureAssetSchemaListResponseAssetClassJSON contains the JSON metadata for
// the struct [CasbPostureAssetSchemaListResponseAssetClass]
type casbPostureAssetSchemaListResponseAssetClassJSON struct {
	AssetClass  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureAssetSchemaListResponseAssetClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureAssetSchemaListResponseAssetClassJSON) RawJSON() string {
	return r.raw
}

type CasbPostureAssetSchemaGetResponse struct {
	// Asset class identifier (PascalCase), e.g. AnthropicApiKey.
	AssetClass string `json:"asset_class" api:"required"`
	// Fields and related attributes a custom finding type may reference for this asset
	// class, as a JSON Schema. Top-level properties are `asset` and, when related
	// resources exist, `associated`.
	Schema map[string]interface{}                `json:"schema" api:"required"`
	JSON   casbPostureAssetSchemaGetResponseJSON `json:"-"`
}

// casbPostureAssetSchemaGetResponseJSON contains the JSON metadata for the struct
// [CasbPostureAssetSchemaGetResponse]
type casbPostureAssetSchemaGetResponseJSON struct {
	AssetClass  apijson.Field
	Schema      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureAssetSchemaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureAssetSchemaGetResponseJSON) RawJSON() string {
	return r.raw
}

type CasbPostureAssetSchemaListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Restrict the result to a single vendor. Unknown values return 400.
	Vendor param.Field[CasbPostureAssetSchemaListParamsVendor] `query:"vendor"`
}

// URLQuery serializes [CasbPostureAssetSchemaListParams]'s query parameters as
// `url.Values`.
func (r CasbPostureAssetSchemaListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Restrict the result to a single vendor. Unknown values return 400.
type CasbPostureAssetSchemaListParamsVendor string

const (
	CasbPostureAssetSchemaListParamsVendorAnthropic           CasbPostureAssetSchemaListParamsVendor = "ANTHROPIC"
	CasbPostureAssetSchemaListParamsVendorAws                 CasbPostureAssetSchemaListParamsVendor = "AWS"
	CasbPostureAssetSchemaListParamsVendorBitbucket           CasbPostureAssetSchemaListParamsVendor = "BITBUCKET"
	CasbPostureAssetSchemaListParamsVendorBox                 CasbPostureAssetSchemaListParamsVendor = "BOX"
	CasbPostureAssetSchemaListParamsVendorConfluence          CasbPostureAssetSchemaListParamsVendor = "CONFLUENCE"
	CasbPostureAssetSchemaListParamsVendorDropbox             CasbPostureAssetSchemaListParamsVendor = "DROPBOX"
	CasbPostureAssetSchemaListParamsVendorGitHub              CasbPostureAssetSchemaListParamsVendor = "GITHUB"
	CasbPostureAssetSchemaListParamsVendorGoogleCloudPlatform CasbPostureAssetSchemaListParamsVendor = "GOOGLE_CLOUD_PLATFORM"
	CasbPostureAssetSchemaListParamsVendorGoogleWorkspace     CasbPostureAssetSchemaListParamsVendor = "GOOGLE_WORKSPACE"
	CasbPostureAssetSchemaListParamsVendorJira                CasbPostureAssetSchemaListParamsVendor = "JIRA"
	CasbPostureAssetSchemaListParamsVendorMicrosoft           CasbPostureAssetSchemaListParamsVendor = "MICROSOFT"
	CasbPostureAssetSchemaListParamsVendorOpenAI              CasbPostureAssetSchemaListParamsVendor = "OPENAI"
	CasbPostureAssetSchemaListParamsVendorSalesforce          CasbPostureAssetSchemaListParamsVendor = "SALESFORCE"
	CasbPostureAssetSchemaListParamsVendorServicenow          CasbPostureAssetSchemaListParamsVendor = "SERVICENOW"
	CasbPostureAssetSchemaListParamsVendorSlack               CasbPostureAssetSchemaListParamsVendor = "SLACK"
	CasbPostureAssetSchemaListParamsVendorWorkday             CasbPostureAssetSchemaListParamsVendor = "WORKDAY"
	CasbPostureAssetSchemaListParamsVendorZoom                CasbPostureAssetSchemaListParamsVendor = "ZOOM"
)

func (r CasbPostureAssetSchemaListParamsVendor) IsKnown() bool {
	switch r {
	case CasbPostureAssetSchemaListParamsVendorAnthropic, CasbPostureAssetSchemaListParamsVendorAws, CasbPostureAssetSchemaListParamsVendorBitbucket, CasbPostureAssetSchemaListParamsVendorBox, CasbPostureAssetSchemaListParamsVendorConfluence, CasbPostureAssetSchemaListParamsVendorDropbox, CasbPostureAssetSchemaListParamsVendorGitHub, CasbPostureAssetSchemaListParamsVendorGoogleCloudPlatform, CasbPostureAssetSchemaListParamsVendorGoogleWorkspace, CasbPostureAssetSchemaListParamsVendorJira, CasbPostureAssetSchemaListParamsVendorMicrosoft, CasbPostureAssetSchemaListParamsVendorOpenAI, CasbPostureAssetSchemaListParamsVendorSalesforce, CasbPostureAssetSchemaListParamsVendorServicenow, CasbPostureAssetSchemaListParamsVendorSlack, CasbPostureAssetSchemaListParamsVendorWorkday, CasbPostureAssetSchemaListParamsVendorZoom:
		return true
	}
	return false
}

type CasbPostureAssetSchemaGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

// Common response structure for all API endpoints.
type CasbPostureAssetSchemaGetResponseEnvelope struct {
	Errors   []CasbPostureAssetSchemaGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureAssetSchemaGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool                                          `json:"success" api:"required"`
	Result  CasbPostureAssetSchemaGetResponse             `json:"result"`
	JSON    casbPostureAssetSchemaGetResponseEnvelopeJSON `json:"-"`
}

// casbPostureAssetSchemaGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPostureAssetSchemaGetResponseEnvelope]
type casbPostureAssetSchemaGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureAssetSchemaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureAssetSchemaGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureAssetSchemaGetResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                `json:"documentation_url" format:"uri"`
	Source           CasbPostureAssetSchemaGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureAssetSchemaGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureAssetSchemaGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CasbPostureAssetSchemaGetResponseEnvelopeErrors]
type casbPostureAssetSchemaGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureAssetSchemaGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureAssetSchemaGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureAssetSchemaGetResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                    `json:"pointer"`
	JSON    casbPostureAssetSchemaGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureAssetSchemaGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CasbPostureAssetSchemaGetResponseEnvelopeErrorsSource]
type casbPostureAssetSchemaGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureAssetSchemaGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureAssetSchemaGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureAssetSchemaGetResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                  `json:"documentation_url" format:"uri"`
	Source           CasbPostureAssetSchemaGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureAssetSchemaGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureAssetSchemaGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CasbPostureAssetSchemaGetResponseEnvelopeMessages]
type casbPostureAssetSchemaGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureAssetSchemaGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureAssetSchemaGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureAssetSchemaGetResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                      `json:"pointer"`
	JSON    casbPostureAssetSchemaGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureAssetSchemaGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CasbPostureAssetSchemaGetResponseEnvelopeMessagesSource]
type casbPostureAssetSchemaGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureAssetSchemaGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureAssetSchemaGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
