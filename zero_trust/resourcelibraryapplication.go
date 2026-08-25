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

// Create a custom application for an account.
func (r *ResourceLibraryApplicationService) New(ctx context.Context, params ResourceLibraryApplicationNewParams, opts ...option.RequestOption) (res *ResourceLibraryApplicationNewResponse, err error) {
	var env ResourceLibraryApplicationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/resource-library/applications", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replace the network matchers for a custom application and create a new version.
func (r *ResourceLibraryApplicationService) Update(ctx context.Context, id int64, params ResourceLibraryApplicationUpdateParams, opts ...option.RequestOption) (res *ResourceLibraryApplicationUpdateResponse, err error) {
	var env ResourceLibraryApplicationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/resource-library/applications/%v", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List the applications available to an account, both the applications Cloudflare
// curates and the custom applications the account has defined.
//
// Results are paginated. Use `filter` and `search` to narrow the list, `order_by`
// to sort it, and `fields` to reduce each result to only the properties you need.
//
// The authenticated principal must have access to the account identified by
// `account_id`.
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

// List the applications available to an account, both the applications Cloudflare
// curates and the custom applications the account has defined.
//
// Results are paginated. Use `filter` and `search` to narrow the list, `order_by`
// to sort it, and `fields` to reduce each result to only the properties you need.
//
// The authenticated principal must have access to the account identified by
// `account_id`.
func (r *ResourceLibraryApplicationService) ListAutoPaging(ctx context.Context, params ResourceLibraryApplicationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ResourceLibraryApplicationListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Delete a custom application and all of its versions. Deletion is rejected when
// other resources reference the application.
func (r *ResourceLibraryApplicationService) Delete(ctx context.Context, id int64, body ResourceLibraryApplicationDeleteParams, opts ...option.RequestOption) (res *ResourceLibraryApplicationDeleteResponse, err error) {
	var env ResourceLibraryApplicationDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/resource-library/applications/%v", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get application by ID.
func (r *ResourceLibraryApplicationService) Get(ctx context.Context, id int64, query ResourceLibraryApplicationGetParams, opts ...option.RequestOption) (res *ResourceLibraryApplicationGetResponse, err error) {
	var env ResourceLibraryApplicationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/resource-library/applications/%v", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ResourceLibraryApplicationNewResponse struct {
	// Returns the application ID.
	ID int64 `json:"id" api:"required"`
	// Confidence score for the application. Returns -1 when no score is available.
	ApplicationConfidenceScore float64 `json:"application_confidence_score" api:"required"`
	// Returns the application source.
	ApplicationSource string `json:"application_source" api:"required"`
	// Returns the application type.
	ApplicationType string `json:"application_type" api:"required"`
	// Returns the application type description.
	ApplicationTypeDescription string `json:"application_type_description" api:"required"`
	// Returns the category ID.
	CategoryID int64 `json:"category_id" api:"required"`
	// Returns the application creation time.
	CreatedAt string `json:"created_at" api:"required"`
	// GenAI score for the application. Returns -1 when no score is available.
	GenAIScore float64 `json:"gen_ai_score" api:"required"`
	// Hostnames matched by the application.
	Hostnames []string `json:"hostnames" api:"required"`
	// Returns the human readable ID.
	HumanID string `json:"human_id" api:"required"`
	// IP subnets matched by the application.
	IPSubnets []string `json:"ip_subnets" api:"required"`
	// Returns the application name.
	Name string `json:"name" api:"required"`
	// Port and protocol pairs matched by the application.
	PortProtocols []string `json:"port_protocols" api:"required"`
	// Support domains matched by the application.
	SupportDomains []string `json:"support_domains" api:"required"`
	// Cloudflare products that support this application.
	Supported []ResourceLibraryApplicationNewResponseSupported `json:"supported" api:"required"`
	// Returns the application update time.
	UpdatedAt string `json:"updated_at" api:"required"`
	// Returns the application version.
	Version string `json:"version" api:"required"`
	// Returns the score composition breakdown for the application.
	ApplicationScoreComposition interface{}                               `json:"application_score_composition" api:"nullable"`
	JSON                        resourceLibraryApplicationNewResponseJSON `json:"-"`
}

// resourceLibraryApplicationNewResponseJSON contains the JSON metadata for the
// struct [ResourceLibraryApplicationNewResponse]
type resourceLibraryApplicationNewResponseJSON struct {
	ID                          apijson.Field
	ApplicationConfidenceScore  apijson.Field
	ApplicationSource           apijson.Field
	ApplicationType             apijson.Field
	ApplicationTypeDescription  apijson.Field
	CategoryID                  apijson.Field
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
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ResourceLibraryApplicationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationNewResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationNewResponseSupported string

const (
	ResourceLibraryApplicationNewResponseSupportedGateway ResourceLibraryApplicationNewResponseSupported = "GATEWAY"
	ResourceLibraryApplicationNewResponseSupportedAccess  ResourceLibraryApplicationNewResponseSupported = "ACCESS"
	ResourceLibraryApplicationNewResponseSupportedCasb    ResourceLibraryApplicationNewResponseSupported = "CASB"
)

func (r ResourceLibraryApplicationNewResponseSupported) IsKnown() bool {
	switch r {
	case ResourceLibraryApplicationNewResponseSupportedGateway, ResourceLibraryApplicationNewResponseSupportedAccess, ResourceLibraryApplicationNewResponseSupportedCasb:
		return true
	}
	return false
}

type ResourceLibraryApplicationUpdateResponse struct {
	// Returns the application ID.
	ID int64 `json:"id" api:"required"`
	// Confidence score for the application. Returns -1 when no score is available.
	ApplicationConfidenceScore float64 `json:"application_confidence_score" api:"required"`
	// Returns the application source.
	ApplicationSource string `json:"application_source" api:"required"`
	// Returns the application type.
	ApplicationType string `json:"application_type" api:"required"`
	// Returns the application type description.
	ApplicationTypeDescription string `json:"application_type_description" api:"required"`
	// Returns the category ID.
	CategoryID int64 `json:"category_id" api:"required"`
	// Returns the application creation time.
	CreatedAt string `json:"created_at" api:"required"`
	// GenAI score for the application. Returns -1 when no score is available.
	GenAIScore float64 `json:"gen_ai_score" api:"required"`
	// Hostnames matched by the application.
	Hostnames []string `json:"hostnames" api:"required"`
	// Returns the human readable ID.
	HumanID string `json:"human_id" api:"required"`
	// IP subnets matched by the application.
	IPSubnets []string `json:"ip_subnets" api:"required"`
	// Returns the application name.
	Name string `json:"name" api:"required"`
	// Port and protocol pairs matched by the application.
	PortProtocols []string `json:"port_protocols" api:"required"`
	// Support domains matched by the application.
	SupportDomains []string `json:"support_domains" api:"required"`
	// Cloudflare products that support this application.
	Supported []ResourceLibraryApplicationUpdateResponseSupported `json:"supported" api:"required"`
	// Returns the application update time.
	UpdatedAt string `json:"updated_at" api:"required"`
	// Returns the application version.
	Version string `json:"version" api:"required"`
	// Returns the score composition breakdown for the application.
	ApplicationScoreComposition interface{}                                  `json:"application_score_composition" api:"nullable"`
	JSON                        resourceLibraryApplicationUpdateResponseJSON `json:"-"`
}

// resourceLibraryApplicationUpdateResponseJSON contains the JSON metadata for the
// struct [ResourceLibraryApplicationUpdateResponse]
type resourceLibraryApplicationUpdateResponseJSON struct {
	ID                          apijson.Field
	ApplicationConfidenceScore  apijson.Field
	ApplicationSource           apijson.Field
	ApplicationType             apijson.Field
	ApplicationTypeDescription  apijson.Field
	CategoryID                  apijson.Field
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
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ResourceLibraryApplicationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationUpdateResponseSupported string

const (
	ResourceLibraryApplicationUpdateResponseSupportedGateway ResourceLibraryApplicationUpdateResponseSupported = "GATEWAY"
	ResourceLibraryApplicationUpdateResponseSupportedAccess  ResourceLibraryApplicationUpdateResponseSupported = "ACCESS"
	ResourceLibraryApplicationUpdateResponseSupportedCasb    ResourceLibraryApplicationUpdateResponseSupported = "CASB"
)

func (r ResourceLibraryApplicationUpdateResponseSupported) IsKnown() bool {
	switch r {
	case ResourceLibraryApplicationUpdateResponseSupportedGateway, ResourceLibraryApplicationUpdateResponseSupportedAccess, ResourceLibraryApplicationUpdateResponseSupportedCasb:
		return true
	}
	return false
}

// Describes one application in a list response. This endpoint returns every
// property below unless the `fields` query parameter narrows the response, so
// treat all of them except `id` as optional.
type ResourceLibraryApplicationListResponse struct {
	// Returns the application ID.
	ID int64 `json:"id" api:"required"`
	// Confidence score for the application. Returns -1 when no score is available.
	ApplicationConfidenceScore float64 `json:"application_confidence_score"`
	// Returns the score composition breakdown for the application.
	ApplicationScoreComposition interface{} `json:"application_score_composition" api:"nullable"`
	// Returns the application source.
	ApplicationSource string `json:"application_source"`
	// Returns the application type.
	ApplicationType string `json:"application_type"`
	// Returns the application type description.
	ApplicationTypeDescription string `json:"application_type_description"`
	// Returns the category ID.
	CategoryID int64 `json:"category_id"`
	// Returns the application creation time.
	CreatedAt string `json:"created_at"`
	// GenAI score for the application. Returns -1 when no score is available.
	GenAIScore float64 `json:"gen_ai_score"`
	// Hostnames matched by the application.
	Hostnames []string `json:"hostnames"`
	// Returns the human readable ID.
	HumanID string `json:"human_id"`
	// IP subnets matched by the application.
	IPSubnets []string `json:"ip_subnets"`
	// Returns the application name.
	Name string `json:"name"`
	// Port and protocol pairs matched by the application.
	PortProtocols []string `json:"port_protocols"`
	// The account-specific Gateway review status. Applications with no assigned review
	// status are returned as `unreviewed`.
	ReviewStatus ResourceLibraryApplicationListResponseReviewStatus `json:"review_status"`
	// Support domains matched by the application.
	SupportDomains []string `json:"support_domains"`
	// Cloudflare products that support this application.
	Supported []ResourceLibraryApplicationListResponseSupported `json:"supported"`
	// Returns the application update time.
	UpdatedAt string `json:"updated_at"`
	// Returns the application version.
	Version string                                     `json:"version"`
	JSON    resourceLibraryApplicationListResponseJSON `json:"-"`
}

// resourceLibraryApplicationListResponseJSON contains the JSON metadata for the
// struct [ResourceLibraryApplicationListResponse]
type resourceLibraryApplicationListResponseJSON struct {
	ID                          apijson.Field
	ApplicationConfidenceScore  apijson.Field
	ApplicationScoreComposition apijson.Field
	ApplicationSource           apijson.Field
	ApplicationType             apijson.Field
	ApplicationTypeDescription  apijson.Field
	CategoryID                  apijson.Field
	CreatedAt                   apijson.Field
	GenAIScore                  apijson.Field
	Hostnames                   apijson.Field
	HumanID                     apijson.Field
	IPSubnets                   apijson.Field
	Name                        apijson.Field
	PortProtocols               apijson.Field
	ReviewStatus                apijson.Field
	SupportDomains              apijson.Field
	Supported                   apijson.Field
	UpdatedAt                   apijson.Field
	Version                     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ResourceLibraryApplicationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationListResponseJSON) RawJSON() string {
	return r.raw
}

// The account-specific Gateway review status. Applications with no assigned review
// status are returned as `unreviewed`.
type ResourceLibraryApplicationListResponseReviewStatus string

const (
	ResourceLibraryApplicationListResponseReviewStatusApproved   ResourceLibraryApplicationListResponseReviewStatus = "approved"
	ResourceLibraryApplicationListResponseReviewStatusUnapproved ResourceLibraryApplicationListResponseReviewStatus = "unapproved"
	ResourceLibraryApplicationListResponseReviewStatusInReview   ResourceLibraryApplicationListResponseReviewStatus = "in_review"
	ResourceLibraryApplicationListResponseReviewStatusUnreviewed ResourceLibraryApplicationListResponseReviewStatus = "unreviewed"
)

func (r ResourceLibraryApplicationListResponseReviewStatus) IsKnown() bool {
	switch r {
	case ResourceLibraryApplicationListResponseReviewStatusApproved, ResourceLibraryApplicationListResponseReviewStatusUnapproved, ResourceLibraryApplicationListResponseReviewStatusInReview, ResourceLibraryApplicationListResponseReviewStatusUnreviewed:
		return true
	}
	return false
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

type ResourceLibraryApplicationDeleteResponse = interface{}

type ResourceLibraryApplicationGetResponse struct {
	// Returns the application ID.
	ID int64 `json:"id" api:"required"`
	// Confidence score for the application. Returns -1 when no score is available.
	ApplicationConfidenceScore float64 `json:"application_confidence_score" api:"required"`
	// Returns the application source.
	ApplicationSource string `json:"application_source" api:"required"`
	// Returns the application type.
	ApplicationType string `json:"application_type" api:"required"`
	// Returns the application type description.
	ApplicationTypeDescription string `json:"application_type_description" api:"required"`
	// Returns the category ID.
	CategoryID int64 `json:"category_id" api:"required"`
	// Returns the application creation time.
	CreatedAt string `json:"created_at" api:"required"`
	// GenAI score for the application. Returns -1 when no score is available.
	GenAIScore float64 `json:"gen_ai_score" api:"required"`
	// Hostnames matched by the application.
	Hostnames []string `json:"hostnames" api:"required"`
	// Returns the human readable ID.
	HumanID string `json:"human_id" api:"required"`
	// IP subnets matched by the application.
	IPSubnets []string `json:"ip_subnets" api:"required"`
	// Returns the application name.
	Name string `json:"name" api:"required"`
	// Port and protocol pairs matched by the application.
	PortProtocols []string `json:"port_protocols" api:"required"`
	// Support domains matched by the application.
	SupportDomains []string `json:"support_domains" api:"required"`
	// Cloudflare products that support this application.
	Supported []ResourceLibraryApplicationGetResponseSupported `json:"supported" api:"required"`
	// Returns the application update time.
	UpdatedAt string `json:"updated_at" api:"required"`
	// Returns the application version.
	Version string `json:"version" api:"required"`
	// Returns the score composition breakdown for the application.
	ApplicationScoreComposition interface{}                               `json:"application_score_composition" api:"nullable"`
	JSON                        resourceLibraryApplicationGetResponseJSON `json:"-"`
}

// resourceLibraryApplicationGetResponseJSON contains the JSON metadata for the
// struct [ResourceLibraryApplicationGetResponse]
type resourceLibraryApplicationGetResponseJSON struct {
	ID                          apijson.Field
	ApplicationConfidenceScore  apijson.Field
	ApplicationSource           apijson.Field
	ApplicationType             apijson.Field
	ApplicationTypeDescription  apijson.Field
	CategoryID                  apijson.Field
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

type ResourceLibraryApplicationNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Returns the category ID.
	CategoryID param.Field[int64] `json:"category_id" api:"required"`
	// Returns the human readable ID.
	HumanID param.Field[string] `json:"human_id" api:"required"`
	// Returns the application name.
	Name param.Field[string] `json:"name" api:"required"`
	// Hostnames matched by the application.
	Hostnames param.Field[[]string] `json:"hostnames"`
	// IP subnets matched by the application.
	IPSubnets param.Field[[]string] `json:"ip_subnets"`
	// Port and protocol pairs matched by the application.
	PortProtocols param.Field[[]string] `json:"port_protocols"`
	// Support domains matched by the application.
	SupportDomains param.Field[[]string] `json:"support_domains"`
}

func (r ResourceLibraryApplicationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceLibraryApplicationNewResponseEnvelope struct {
	Errors   []ResourceLibraryApplicationNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ResourceLibraryApplicationNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Indicates whether the API call was successful.
	Success ResourceLibraryApplicationNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  ResourceLibraryApplicationNewResponse                `json:"result"`
	JSON    resourceLibraryApplicationNewResponseEnvelopeJSON    `json:"-"`
}

// resourceLibraryApplicationNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ResourceLibraryApplicationNewResponseEnvelope]
type resourceLibraryApplicationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryApplicationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationNewResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           ResourceLibraryApplicationNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             resourceLibraryApplicationNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// resourceLibraryApplicationNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ResourceLibraryApplicationNewResponseEnvelopeErrors]
type resourceLibraryApplicationNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceLibraryApplicationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    resourceLibraryApplicationNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// resourceLibraryApplicationNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [ResourceLibraryApplicationNewResponseEnvelopeErrorsSource]
type resourceLibraryApplicationNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryApplicationNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationNewResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           ResourceLibraryApplicationNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             resourceLibraryApplicationNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// resourceLibraryApplicationNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ResourceLibraryApplicationNewResponseEnvelopeMessages]
type resourceLibraryApplicationNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceLibraryApplicationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    resourceLibraryApplicationNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// resourceLibraryApplicationNewResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [ResourceLibraryApplicationNewResponseEnvelopeMessagesSource]
type resourceLibraryApplicationNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryApplicationNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type ResourceLibraryApplicationNewResponseEnvelopeSuccess bool

const (
	ResourceLibraryApplicationNewResponseEnvelopeSuccessTrue ResourceLibraryApplicationNewResponseEnvelopeSuccess = true
)

func (r ResourceLibraryApplicationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ResourceLibraryApplicationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ResourceLibraryApplicationUpdateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Hostnames matched by the application.
	Hostnames param.Field[[]string] `json:"hostnames"`
	// IP subnets matched by the application.
	IPSubnets param.Field[[]string] `json:"ip_subnets"`
	// Port and protocol pairs matched by the application.
	PortProtocols param.Field[[]string] `json:"port_protocols"`
	// Support domains matched by the application.
	SupportDomains param.Field[[]string] `json:"support_domains"`
}

func (r ResourceLibraryApplicationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceLibraryApplicationUpdateResponseEnvelope struct {
	Errors   []ResourceLibraryApplicationUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ResourceLibraryApplicationUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Indicates whether the API call was successful.
	Success ResourceLibraryApplicationUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  ResourceLibraryApplicationUpdateResponse                `json:"result"`
	JSON    resourceLibraryApplicationUpdateResponseEnvelopeJSON    `json:"-"`
}

// resourceLibraryApplicationUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ResourceLibraryApplicationUpdateResponseEnvelope]
type resourceLibraryApplicationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryApplicationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationUpdateResponseEnvelopeErrors struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           ResourceLibraryApplicationUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             resourceLibraryApplicationUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// resourceLibraryApplicationUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ResourceLibraryApplicationUpdateResponseEnvelopeErrors]
type resourceLibraryApplicationUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceLibraryApplicationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    resourceLibraryApplicationUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// resourceLibraryApplicationUpdateResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [ResourceLibraryApplicationUpdateResponseEnvelopeErrorsSource]
type resourceLibraryApplicationUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryApplicationUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationUpdateResponseEnvelopeMessages struct {
	Code             int64                                                          `json:"code" api:"required"`
	Message          string                                                         `json:"message" api:"required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           ResourceLibraryApplicationUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             resourceLibraryApplicationUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// resourceLibraryApplicationUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ResourceLibraryApplicationUpdateResponseEnvelopeMessages]
type resourceLibraryApplicationUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceLibraryApplicationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    resourceLibraryApplicationUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// resourceLibraryApplicationUpdateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [ResourceLibraryApplicationUpdateResponseEnvelopeMessagesSource]
type resourceLibraryApplicationUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryApplicationUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type ResourceLibraryApplicationUpdateResponseEnvelopeSuccess bool

const (
	ResourceLibraryApplicationUpdateResponseEnvelopeSuccessTrue ResourceLibraryApplicationUpdateResponseEnvelopeSuccess = true
)

func (r ResourceLibraryApplicationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ResourceLibraryApplicationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ResourceLibraryApplicationListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Return only the listed properties on each application, as a comma-separated
	// list. Use this to keep responses small when you only need part of each
	// application — for example populating a picker with `fields=id,name` instead of
	// downloading every hostname and IP subnet.
	//
	// Omit this parameter to receive the full application object.
	//
	// `id` is always returned.
	//
	// Selectable properties: `id`, `name`, `human_id`, `version`, `hostnames`,
	// `support_domains`, `ip_subnets`, `port_protocols`, `supported`, `gen_ai_score`,
	// `application_confidence_score`, `created_at`, `updated_at`, `review_status`.
	//
	// Unknown or empty property names return `400`.
	Fields param.Field[string] `query:"fields"`
	// Filter applications using key:value format. Supported filter keys:
	//
	//   - name: Filter by application name (e.g., name:HR)
	//   - id: Filter by application ID (e.g., id:498)
	//   - human_id: Filter by human-readable ID (e.g., human_id:HR)
	//   - hostname: Filter by hostname or support domain (e.g.,
	//     hostname:portal.example.com)
	//   - source: Filter by application source name (e.g., source:cloudflare)
	//   - ip_subnet: Filter by IP subnet using CIDR containment — returns applications
	//     where any stored subnet contains the search value (e.g., ip_subnet:10.0.1.5/32
	//     matches apps with 10.0.0.0/16)
	//   - category_id: Filter by category ID (e.g., category_id:12).
	//   - category_name: Filter by category name (e.g., category_name:HR).
	//   - supported: Filter by supported Cloudflare product (e.g., supported:ACCESS).
	//     Values: GATEWAY, ACCESS, CASB.
	//   - review_status: Filter by the account's Gateway review status. Values:
	//     approved, unapproved, in_review, unreviewed. .
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

type ResourceLibraryApplicationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ResourceLibraryApplicationDeleteResponseEnvelope struct {
	Errors   []ResourceLibraryApplicationDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ResourceLibraryApplicationDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ResourceLibraryApplicationDeleteResponse                   `json:"result" api:"required,nullable"`
	// Indicates whether the API call was successful.
	Success ResourceLibraryApplicationDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    resourceLibraryApplicationDeleteResponseEnvelopeJSON    `json:"-"`
}

// resourceLibraryApplicationDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [ResourceLibraryApplicationDeleteResponseEnvelope]
type resourceLibraryApplicationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryApplicationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationDeleteResponseEnvelopeErrors struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           ResourceLibraryApplicationDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             resourceLibraryApplicationDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// resourceLibraryApplicationDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ResourceLibraryApplicationDeleteResponseEnvelopeErrors]
type resourceLibraryApplicationDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceLibraryApplicationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    resourceLibraryApplicationDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// resourceLibraryApplicationDeleteResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [ResourceLibraryApplicationDeleteResponseEnvelopeErrorsSource]
type resourceLibraryApplicationDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryApplicationDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationDeleteResponseEnvelopeMessages struct {
	Code             int64                                                          `json:"code" api:"required"`
	Message          string                                                         `json:"message" api:"required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           ResourceLibraryApplicationDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             resourceLibraryApplicationDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// resourceLibraryApplicationDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ResourceLibraryApplicationDeleteResponseEnvelopeMessages]
type resourceLibraryApplicationDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceLibraryApplicationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryApplicationDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    resourceLibraryApplicationDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// resourceLibraryApplicationDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [ResourceLibraryApplicationDeleteResponseEnvelopeMessagesSource]
type resourceLibraryApplicationDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryApplicationDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryApplicationDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type ResourceLibraryApplicationDeleteResponseEnvelopeSuccess bool

const (
	ResourceLibraryApplicationDeleteResponseEnvelopeSuccessTrue ResourceLibraryApplicationDeleteResponseEnvelopeSuccess = true
)

func (r ResourceLibraryApplicationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ResourceLibraryApplicationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
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

// Indicates whether the API call was successful.
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
