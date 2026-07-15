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
)

// CasbApplicationSetupFlowService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbApplicationSetupFlowService] method instead.
type CasbApplicationSetupFlowService struct {
	Options []option.RequestOption
}

// NewCasbApplicationSetupFlowService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCasbApplicationSetupFlowService(opts ...option.RequestOption) (r *CasbApplicationSetupFlowService) {
	r = &CasbApplicationSetupFlowService{}
	r.Options = opts
	return
}

// Returns all available setup flows for the application, one per auth method.
func (r *CasbApplicationSetupFlowService) List(ctx context.Context, slug string, params CasbApplicationSetupFlowListParams, opts ...option.RequestOption) (res *[]CasbApplicationSetupFlowListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/one/applications/%s/setup-flows", params.AccountID, slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Setup flow for an application auth method.
type CasbApplicationSetupFlowListResponse struct {
	// Setup flow identifier.
	ID string `json:"id" api:"required"`
	// Whether this is the default auth method.
	Default bool `json:"default" api:"required"`
	// Flow description.
	Description string `json:"description" api:"required"`
	// Human-readable flow name.
	Name string `json:"name" api:"required"`
	// Ordered list of setup steps.
	Steps []CasbApplicationSetupFlowListResponseStep `json:"steps" api:"required"`
	// Environments this auth method supports (standard, fedramp).
	SupportedEnvironments []string `json:"supported_environments" api:"required"`
	// OAuth configuration (present for OAuth-based flows).
	AuthConfig CasbApplicationSetupFlowListResponseAuthConfig `json:"auth_config"`
	JSON       casbApplicationSetupFlowListResponseJSON       `json:"-"`
}

// casbApplicationSetupFlowListResponseJSON contains the JSON metadata for the
// struct [CasbApplicationSetupFlowListResponse]
type casbApplicationSetupFlowListResponseJSON struct {
	ID                    apijson.Field
	Default               apijson.Field
	Description           apijson.Field
	Name                  apijson.Field
	Steps                 apijson.Field
	SupportedEnvironments apijson.Field
	AuthConfig            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CasbApplicationSetupFlowListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationSetupFlowListResponseJSON) RawJSON() string {
	return r.raw
}

// A single step in the setup flow. Polymorphic based on type.
type CasbApplicationSetupFlowListResponseStep struct {
	// Step type.
	//
	// - `component` - component
	// - `instruction` - instruction
	// - `form_input` - form_input
	// - `oauth_redirect` - oauth_redirect
	Type CasbApplicationSetupFlowListResponseStepsType `json:"type" api:"required"`
	// Component identifier (for component type).
	ComponentID string `json:"component_id" api:"nullable"`
	// Step description with markdown support.
	Description string `json:"description" api:"nullable"`
	// Dynamic content blocks (for instruction/form_input).
	DynamicContent []CasbApplicationSetupFlowListResponseStepsDynamicContent `json:"dynamic_content" api:"nullable"`
	// Form fields (for form_input).
	FormFields []CasbApplicationSetupFlowListResponseStepsFormField `json:"form_fields"`
	// Whether step is required (for form_input).
	IsRequired bool `json:"is_required"`
	// Component parameters (for component type).
	Parameters map[string]string `json:"parameters" api:"nullable"`
	// Step title (for instruction/form_input/oauth_redirect).
	Title string                                       `json:"title"`
	JSON  casbApplicationSetupFlowListResponseStepJSON `json:"-"`
}

// casbApplicationSetupFlowListResponseStepJSON contains the JSON metadata for the
// struct [CasbApplicationSetupFlowListResponseStep]
type casbApplicationSetupFlowListResponseStepJSON struct {
	Type           apijson.Field
	ComponentID    apijson.Field
	Description    apijson.Field
	DynamicContent apijson.Field
	FormFields     apijson.Field
	IsRequired     apijson.Field
	Parameters     apijson.Field
	Title          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CasbApplicationSetupFlowListResponseStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationSetupFlowListResponseStepJSON) RawJSON() string {
	return r.raw
}

// Step type.
//
// - `component` - component
// - `instruction` - instruction
// - `form_input` - form_input
// - `oauth_redirect` - oauth_redirect
type CasbApplicationSetupFlowListResponseStepsType string

const (
	CasbApplicationSetupFlowListResponseStepsTypeComponent     CasbApplicationSetupFlowListResponseStepsType = "component"
	CasbApplicationSetupFlowListResponseStepsTypeInstruction   CasbApplicationSetupFlowListResponseStepsType = "instruction"
	CasbApplicationSetupFlowListResponseStepsTypeFormInput     CasbApplicationSetupFlowListResponseStepsType = "form_input"
	CasbApplicationSetupFlowListResponseStepsTypeOAuthRedirect CasbApplicationSetupFlowListResponseStepsType = "oauth_redirect"
)

func (r CasbApplicationSetupFlowListResponseStepsType) IsKnown() bool {
	switch r {
	case CasbApplicationSetupFlowListResponseStepsTypeComponent, CasbApplicationSetupFlowListResponseStepsTypeInstruction, CasbApplicationSetupFlowListResponseStepsTypeFormInput, CasbApplicationSetupFlowListResponseStepsTypeOAuthRedirect:
		return true
	}
	return false
}

// Dynamic content for instruction/form_input steps.
type CasbApplicationSetupFlowListResponseStepsDynamicContent struct {
	// Display label.
	Label string `json:"label" api:"required"`
	// Content type.
	//
	// - `copy_block` - copy_block
	// - `external_link` - external_link
	Type CasbApplicationSetupFlowListResponseStepsDynamicContentType `json:"type" api:"required"`
	// URL template with {{ variable }} interpolation (for external_link).
	URLTemplate string `json:"url_template" api:"nullable"`
	// Field path to get value from (for copy_block).
	ValueFrom string                                                      `json:"value_from" api:"nullable"`
	JSON      casbApplicationSetupFlowListResponseStepsDynamicContentJSON `json:"-"`
}

// casbApplicationSetupFlowListResponseStepsDynamicContentJSON contains the JSON
// metadata for the struct
// [CasbApplicationSetupFlowListResponseStepsDynamicContent]
type casbApplicationSetupFlowListResponseStepsDynamicContentJSON struct {
	Label       apijson.Field
	Type        apijson.Field
	URLTemplate apijson.Field
	ValueFrom   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbApplicationSetupFlowListResponseStepsDynamicContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationSetupFlowListResponseStepsDynamicContentJSON) RawJSON() string {
	return r.raw
}

// Content type.
//
// - `copy_block` - copy_block
// - `external_link` - external_link
type CasbApplicationSetupFlowListResponseStepsDynamicContentType string

const (
	CasbApplicationSetupFlowListResponseStepsDynamicContentTypeCopyBlock    CasbApplicationSetupFlowListResponseStepsDynamicContentType = "copy_block"
	CasbApplicationSetupFlowListResponseStepsDynamicContentTypeExternalLink CasbApplicationSetupFlowListResponseStepsDynamicContentType = "external_link"
)

func (r CasbApplicationSetupFlowListResponseStepsDynamicContentType) IsKnown() bool {
	switch r {
	case CasbApplicationSetupFlowListResponseStepsDynamicContentTypeCopyBlock, CasbApplicationSetupFlowListResponseStepsDynamicContentTypeExternalLink:
		return true
	}
	return false
}

// A form field within a form_input step.
type CasbApplicationSetupFlowListResponseStepsFormField struct {
	// Human-readable field label.
	Label string `json:"label" api:"required"`
	// Field identifier (maps to credentials key).
	Name string `json:"name" api:"required"`
	// Placeholder text.
	Placeholder string `json:"placeholder" api:"required,nullable"`
	// Whether field is required.
	Required bool `json:"required" api:"required"`
	// Allowed file extensions for file_upload type.
	SupportedFileTypes []string `json:"supported_file_types" api:"required,nullable"`
	// Field input type.
	//
	// - `text` - text
	// - `password` - password
	// - `email` - email
	// - `file_upload` - file_upload
	Type CasbApplicationSetupFlowListResponseStepsFormFieldsType `json:"type" api:"required"`
	JSON casbApplicationSetupFlowListResponseStepsFormFieldJSON  `json:"-"`
}

// casbApplicationSetupFlowListResponseStepsFormFieldJSON contains the JSON
// metadata for the struct [CasbApplicationSetupFlowListResponseStepsFormField]
type casbApplicationSetupFlowListResponseStepsFormFieldJSON struct {
	Label              apijson.Field
	Name               apijson.Field
	Placeholder        apijson.Field
	Required           apijson.Field
	SupportedFileTypes apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CasbApplicationSetupFlowListResponseStepsFormField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationSetupFlowListResponseStepsFormFieldJSON) RawJSON() string {
	return r.raw
}

// Field input type.
//
// - `text` - text
// - `password` - password
// - `email` - email
// - `file_upload` - file_upload
type CasbApplicationSetupFlowListResponseStepsFormFieldsType string

const (
	CasbApplicationSetupFlowListResponseStepsFormFieldsTypeText       CasbApplicationSetupFlowListResponseStepsFormFieldsType = "text"
	CasbApplicationSetupFlowListResponseStepsFormFieldsTypePassword   CasbApplicationSetupFlowListResponseStepsFormFieldsType = "password"
	CasbApplicationSetupFlowListResponseStepsFormFieldsTypeEmail      CasbApplicationSetupFlowListResponseStepsFormFieldsType = "email"
	CasbApplicationSetupFlowListResponseStepsFormFieldsTypeFileUpload CasbApplicationSetupFlowListResponseStepsFormFieldsType = "file_upload"
)

func (r CasbApplicationSetupFlowListResponseStepsFormFieldsType) IsKnown() bool {
	switch r {
	case CasbApplicationSetupFlowListResponseStepsFormFieldsTypeText, CasbApplicationSetupFlowListResponseStepsFormFieldsTypePassword, CasbApplicationSetupFlowListResponseStepsFormFieldsTypeEmail, CasbApplicationSetupFlowListResponseStepsFormFieldsTypeFileUpload:
		return true
	}
	return false
}

// OAuth configuration (present for OAuth-based flows).
type CasbApplicationSetupFlowListResponseAuthConfig struct {
	// Authorization URL for the requested environment.
	AuthorizationURL string `json:"authorization_url" api:"required,nullable"`
	// OAuth client ID.
	ClientID string `json:"client_id" api:"required,nullable"`
	// Whether PKCE is required.
	RequiresPKCE bool `json:"requires_pkce" api:"required"`
	// OAuth scopes to request.
	Scopes []string `json:"scopes" api:"required"`
	// Placeholders in authorization URL that frontend must fill.
	URLPlaceholders []string                                           `json:"url_placeholders" api:"required"`
	JSON            casbApplicationSetupFlowListResponseAuthConfigJSON `json:"-"`
}

// casbApplicationSetupFlowListResponseAuthConfigJSON contains the JSON metadata
// for the struct [CasbApplicationSetupFlowListResponseAuthConfig]
type casbApplicationSetupFlowListResponseAuthConfigJSON struct {
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	RequiresPKCE     apijson.Field
	Scopes           apijson.Field
	URLPlaceholders  apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbApplicationSetupFlowListResponseAuthConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationSetupFlowListResponseAuthConfigJSON) RawJSON() string {
	return r.raw
}

type CasbApplicationSetupFlowListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter by auth method slug. Get available slugs from GET /v2/applications.
	AuthMethod param.Field[string] `query:"auth_method"`
	// Filter by environment.
	Environment param.Field[CasbApplicationSetupFlowListParamsEnvironment] `query:"environment"`
}

// URLQuery serializes [CasbApplicationSetupFlowListParams]'s query parameters as
// `url.Values`.
func (r CasbApplicationSetupFlowListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by environment.
type CasbApplicationSetupFlowListParamsEnvironment string

const (
	CasbApplicationSetupFlowListParamsEnvironmentFedramp  CasbApplicationSetupFlowListParamsEnvironment = "fedramp"
	CasbApplicationSetupFlowListParamsEnvironmentStandard CasbApplicationSetupFlowListParamsEnvironment = "standard"
)

func (r CasbApplicationSetupFlowListParamsEnvironment) IsKnown() bool {
	switch r {
	case CasbApplicationSetupFlowListParamsEnvironmentFedramp, CasbApplicationSetupFlowListParamsEnvironmentStandard:
		return true
	}
	return false
}
