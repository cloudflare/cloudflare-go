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

// ResourceLibraryApplicationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceLibraryApplicationService] method instead.
type ResourceLibraryApplicationService struct {
	Options []option.RequestOption
}

// NewResourceLibraryApplicationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewResourceLibraryApplicationService(opts ...option.RequestOption) (r *ResourceLibraryApplicationService) {
	r = &ResourceLibraryApplicationService{}
	r.Options = opts
	return
}

// List applications with different filters.
func (r *ResourceLibraryApplicationService) List(ctx context.Context, params ResourceLibraryApplicationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ResourceLibraryApplicationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/resource-library/applications", params.AccountID)
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

// List applications with different filters.
func (r *ResourceLibraryApplicationService) ListAutoPaging(ctx context.Context, params ResourceLibraryApplicationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ResourceLibraryApplicationListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Get application by ID.
func (r *ResourceLibraryApplicationService) Get(ctx context.Context, id string, query ResourceLibraryApplicationGetParams, opts ...option.RequestOption) (res *ResourceLibraryApplicationGetResponse, err error) {
	var env ResourceLibraryApplicationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/resource-library/applications/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ResourceLibraryApplicationListResponse struct {
	// Returns the application ID.
	ID string `json:"id" api:"required"`
	// Confidence score for the application. Returns -1 when no score is available.
	ApplicationConfidenceScore float64 `json:"application_confidence_score" api:"required"`
	// Returns the application source.
	ApplicationSource string `json:"application_source" api:"required"`
	// Returns the application type.
	ApplicationType string `json:"application_type" api:"required"`
	// Returns the application type description.
	ApplicationTypeDescription string `json:"application_type_description" api:"required"`
	// Returns the application creation time.
	CreatedAt string `json:"created_at" api:"required"`
	// GenAI score for the application. Returns -1 when no score is available.
	GenAIScore float64 `json:"gen_ai_score" api:"required"`
	// Returns the list of hostnames for the application.
	Hostnames []string `json:"hostnames" api:"required"`
	// Returns the human readable ID.
	HumanID string `json:"human_id" api:"required"`
	// Returns the list of IP subnets for the application.
	IPSubnets []string `json:"ip_subnets" api:"required"`
	// Returns the application name.
	Name string `json:"name" api:"required"`
	// Returns the list of port protocols for the application.
	PortProtocols []string `json:"port_protocols" api:"required"`
	// Returns the list of support domains for the application.
	SupportDomains []string `json:"support_domains" api:"required"`
	// Cloudflare products that support this application.
	Supported []ResourceLibraryApplicationListResponseSupported `json:"supported" api:"required"`
	// Returns the application update time.
	UpdatedAt string `json:"updated_at" api:"required"`
	// Returns the application version.
	Version string `json:"version" api:"required"`
	// Returns the score composition breakdown for the application.
	ApplicationScoreComposition interface{} `json:"application_score_composition" api:"nullable"`
	// Returns the Intel API ID for the application.
	IntelID int64                                      `json:"intel_id" api:"nullable"`
	JSON    resourceLibraryApplicationListResponseJSON `json:"-"`
}

// resourceLibraryApplicationListResponseJSON contains the JSON metadata for the
// struct [ResourceLibraryApplicationListResponse]
type resourceLibraryApplicationListResponseJSON struct {
	ID                          apijson.Field
	ApplicationConfidenceScore  apijson.Field
	ApplicationSource           apijson.Field
	ApplicationType             apijson.Field
	ApplicationTypeDescription  apijson.Field
	CreatedAt                   apijson.Field
	GenAIScore                  apijson.Field
	Hostnames                   apijson.Field
	HumanID                     apijson.Field
	IPSubnets                   apijson.Field
	Name                        apijson.Field
	PortProtocols               apijson.Field
	SupportDomains              apijson.Field
	Supported                   apijson.Field
	UpdatedAt                   apijson.Field
	Version                     apijson.Field
	ApplicationScoreComposition apijson.Field
	IntelID                     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ResourceLibraryApplicationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationListResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationListResponseSupported string

const (
	ResourceLibraryApplicationListResponseSupportedGateway ResourceLibraryApplicationListResponseSupported = "GATEWAY"
	ResourceLibraryApplicationListResponseSupportedAccess  ResourceLibraryApplicationListResponseSupported = "ACCESS"
	ResourceLibraryApplicationListResponseSupportedCasb    ResourceLibraryApplicationListResponseSupported = "CASB"
)

func (r ResourceLibraryApplicationListResponseSupported) IsKnown() bool {
	switch r {
	case ResourceLibraryApplicationListResponseSupportedGateway, ResourceLibraryApplicationListResponseSupportedAccess, ResourceLibraryApplicationListResponseSupportedCasb:
		return true
	}
	return false
}

type ResourceLibraryApplicationGetResponse struct {
	// Returns the application ID.
	ID string `json:"id" api:"required"`
	// Confidence score for the application. Returns -1 when no score is available.
	ApplicationConfidenceScore float64 `json:"application_confidence_score" api:"required"`
	// Returns the application source.
	ApplicationSource string `json:"application_source" api:"required"`
	// Returns the application type.
	ApplicationType string `json:"application_type" api:"required"`
	// Returns the application type description.
	ApplicationTypeDescription string `json:"application_type_description" api:"required"`
	// Returns the application creation time.
	CreatedAt string `json:"created_at" api:"required"`
	// GenAI score for the application. Returns -1 when no score is available.
	GenAIScore float64 `json:"gen_ai_score" api:"required"`
	// Returns the list of hostnames for the application.
	Hostnames []string `json:"hostnames" api:"required"`
	// Returns the human readable ID.
	HumanID string `json:"human_id" api:"required"`
	// Returns the list of IP subnets for the application.
	IPSubnets []string `json:"ip_subnets" api:"required"`
	// Returns the application name.
	Name string `json:"name" api:"required"`
	// Returns the list of port protocols for the application.
	PortProtocols []string `json:"port_protocols" api:"required"`
	// Returns the list of support domains for the application.
	SupportDomains []string `json:"support_domains" api:"required"`
	// Cloudflare products that support this application.
	Supported []ResourceLibraryApplicationGetResponseSupported `json:"supported" api:"required"`
	// Returns the application update time.
	UpdatedAt string `json:"updated_at" api:"required"`
	// Returns the application version.
	Version string `json:"version" api:"required"`
	// Returns the score composition breakdown for the application.
	ApplicationScoreComposition interface{} `json:"application_score_composition" api:"nullable"`
	// Returns the Intel API ID for the application.
	IntelID int64                                     `json:"intel_id" api:"nullable"`
	JSON    resourceLibraryApplicationGetResponseJSON `json:"-"`
}

// resourceLibraryApplicationGetResponseJSON contains the JSON metadata for the
// struct [ResourceLibraryApplicationGetResponse]
type resourceLibraryApplicationGetResponseJSON struct {
	ID                          apijson.Field
	ApplicationConfidenceScore  apijson.Field
	ApplicationSource           apijson.Field
	ApplicationType             apijson.Field
	ApplicationTypeDescription  apijson.Field
	CreatedAt                   apijson.Field
	GenAIScore                  apijson.Field
	Hostnames                   apijson.Field
	HumanID                     apijson.Field
	IPSubnets                   apijson.Field
	Name                        apijson.Field
	PortProtocols               apijson.Field
	SupportDomains              apijson.Field
	Supported                   apijson.Field
	UpdatedAt                   apijson.Field
	Version                     apijson.Field
	ApplicationScoreComposition apijson.Field
	IntelID                     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ResourceLibraryApplicationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationGetResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationGetResponseSupported string

const (
	ResourceLibraryApplicationGetResponseSupportedGateway ResourceLibraryApplicationGetResponseSupported = "GATEWAY"
	ResourceLibraryApplicationGetResponseSupportedAccess  ResourceLibraryApplicationGetResponseSupported = "ACCESS"
	ResourceLibraryApplicationGetResponseSupportedCasb    ResourceLibraryApplicationGetResponseSupported = "CASB"
)

func (r ResourceLibraryApplicationGetResponseSupported) IsKnown() bool {
	switch r {
	case ResourceLibraryApplicationGetResponseSupportedGateway, ResourceLibraryApplicationGetResponseSupportedAccess, ResourceLibraryApplicationGetResponseSupportedCasb:
		return true
	}
	return false
}

type ResourceLibraryApplicationListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter applications using key:value format. Supported filter keys:
	//
	//   - name: Filter by application name (e.g., name:HR)
	//   - id: Filter by application ID (e.g., id:0b63249c-95bf-4cc0-a7cc-d7faaaf1dac0)
	//   - human_id: Filter by human-readable ID (e.g., human_id:HR)
	//   - hostname: Filter by hostname or support domain (e.g.,
	//     hostname:portal.example.com)
	//   - source: Filter by application source name (e.g., source:cloudflare)
	//   - ip_subnet: Filter by IP subnet using CIDR containment — returns applications
	//     where any stored subnet contains the search value (e.g., ip_subnet:10.0.1.5/32
	//     matches apps with 10.0.0.0/16)
	//   - intel_id: Filter by Intel API ID (e.g., intel_id:498). also supports multiple
	//     values (e.g., intel_id:498,1001)
	//   - category_id: Filter by category ID (e.g.,
	//     category_id:37f8ec03-8766-49d4-9a15-369b044c842c).
	//   - category_name: Filter by category name (e.g., category_name:HR).
	//   - supported: Filter by supported Cloudflare product (e.g., supported:ACCESS).
	//     Values: GATEWAY, ACCESS, CASB. .
	Filter param.Field[string] `query:"filter"`
	// Limit of number of results to return (max 250).
	Limit param.Field[int64] `query:"limit"`
	// Offset of results to return.
	Offset param.Field[int64] `query:"offset"`
	// Order results by field name and direction (e.g., name:asc). Ignored when search
	// is provided; results are ranked by relevance instead.
	OrderBy param.Field[string] `query:"order_by"`
	// Fuzzy search across application name and hostnames. Results are ranked by
	// relevance. Must be between 2 and 200 characters. Can be combined with filter
	// parameters.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [ResourceLibraryApplicationListParams]'s query parameters as
// `url.Values`.
func (r ResourceLibraryApplicationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ResourceLibraryApplicationGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ResourceLibraryApplicationGetResponseEnvelope struct {
	Errors   []ResourceLibraryApplicationGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ResourceLibraryApplicationGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Indicates whether the API call was successful.
	Success ResourceLibraryApplicationGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  ResourceLibraryApplicationGetResponse                `json:"result"`
	JSON    resourceLibraryApplicationGetResponseEnvelopeJSON    `json:"-"`
}

// resourceLibraryApplicationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ResourceLibraryApplicationGetResponseEnvelope]
type resourceLibraryApplicationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryApplicationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationGetResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           ResourceLibraryApplicationGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             resourceLibraryApplicationGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// resourceLibraryApplicationGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ResourceLibraryApplicationGetResponseEnvelopeErrors]
type resourceLibraryApplicationGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceLibraryApplicationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    resourceLibraryApplicationGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// resourceLibraryApplicationGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [ResourceLibraryApplicationGetResponseEnvelopeErrorsSource]
type resourceLibraryApplicationGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryApplicationGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationGetResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           ResourceLibraryApplicationGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             resourceLibraryApplicationGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// resourceLibraryApplicationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ResourceLibraryApplicationGetResponseEnvelopeMessages]
type resourceLibraryApplicationGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceLibraryApplicationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    resourceLibraryApplicationGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// resourceLibraryApplicationGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [ResourceLibraryApplicationGetResponseEnvelopeMessagesSource]
type resourceLibraryApplicationGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryApplicationGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// ResourceLibraryApplicationGetResponseEnvelopeSuccess indicates whether the API call was successful.
type ResourceLibraryApplicationGetResponseEnvelopeSuccess bool

const (
	ResourceLibraryApplicationGetResponseEnvelopeSuccessTrue ResourceLibraryApplicationGetResponseEnvelopeSuccess = true
)

func (r ResourceLibraryApplicationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ResourceLibraryApplicationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
