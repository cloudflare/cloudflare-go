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

// CasbApplicationAuthMethodService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbApplicationAuthMethodService] method instead.
type CasbApplicationAuthMethodService struct {
	Options []option.RequestOption
}

// NewCasbApplicationAuthMethodService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCasbApplicationAuthMethodService(opts ...option.RequestOption) (r *CasbApplicationAuthMethodService) {
	r = &CasbApplicationAuthMethodService{}
	r.Options = opts
	return
}

// Returns available auth methods for the specified vendor, including credential
// schema, instructions, and example payloads. Use this to understand what
// credentials are required before calling POST /v2/integrations.
func (r *CasbApplicationAuthMethodService) List(ctx context.Context, applicationID CasbApplicationAuthMethodListParamsApplicationID, params CasbApplicationAuthMethodListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CasbApplicationAuthMethodListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/one/applications/%v/auth-methods", params.AccountID, applicationID)
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

// Returns available auth methods for the specified vendor, including credential
// schema, instructions, and example payloads. Use this to understand what
// credentials are required before calling POST /v2/integrations.
func (r *CasbApplicationAuthMethodService) ListAutoPaging(ctx context.Context, applicationID CasbApplicationAuthMethodListParamsApplicationID, params CasbApplicationAuthMethodListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CasbApplicationAuthMethodListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, applicationID, params, opts...))
}

// Detailed auth method info including credentials schema and instructions.
type CasbApplicationAuthMethodListResponse struct {
	// Auth method identifier.
	ID string `json:"id" api:"required"`
	// Human-readable auth method name.
	DisplayName string `json:"display_name" api:"required"`
	// Whether setup requires human interaction or integration can be created purely
	// using API (e.g., For OAuth can not be created without user interaction).
	HumanInteractionRequired bool `json:"human_interaction_required" api:"required"`
	// Step-by-step instructions for obtaining credentials.
	Instructions CasbApplicationAuthMethodListResponseInstructions `json:"instructions" api:"required"`
	// Example credentials payload with placeholder values.
	PayloadExample map[string]interface{} `json:"payload_example" api:"required,nullable"`
	// JSON Schema for the credentials object in POST /v2/integrations request.
	PayloadSchema map[string]interface{} `json:"payload_schema" api:"required,nullable"`
	// OAuth redirect URL for vendors requiring human interaction.
	RedirectURL string                                    `json:"redirect_url" api:"required,nullable"`
	JSON        casbApplicationAuthMethodListResponseJSON `json:"-"`
}

// casbApplicationAuthMethodListResponseJSON contains the JSON metadata for the
// struct [CasbApplicationAuthMethodListResponse]
type casbApplicationAuthMethodListResponseJSON struct {
	ID                       apijson.Field
	DisplayName              apijson.Field
	HumanInteractionRequired apijson.Field
	Instructions             apijson.Field
	PayloadExample           apijson.Field
	PayloadSchema            apijson.Field
	RedirectURL              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CasbApplicationAuthMethodListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationAuthMethodListResponseJSON) RawJSON() string {
	return r.raw
}

// Step-by-step instructions for obtaining credentials.
type CasbApplicationAuthMethodListResponseInstructions struct {
	// Detailed instructions in markdown format.
	Markdown string                                                `json:"markdown" api:"required"`
	JSON     casbApplicationAuthMethodListResponseInstructionsJSON `json:"-"`
}

// casbApplicationAuthMethodListResponseInstructionsJSON contains the JSON metadata
// for the struct [CasbApplicationAuthMethodListResponseInstructions]
type casbApplicationAuthMethodListResponseInstructionsJSON struct {
	Markdown    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbApplicationAuthMethodListResponseInstructions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationAuthMethodListResponseInstructionsJSON) RawJSON() string {
	return r.raw
}

type CasbApplicationAuthMethodListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A page number within the paginated result set.
	Page param.Field[int64] `query:"page"`
	// Number of results to return per page.
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [CasbApplicationAuthMethodListParams]'s query parameters as
// `url.Values`.
func (r CasbApplicationAuthMethodListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CasbApplicationAuthMethodListParamsApplicationID string

const (
	CasbApplicationAuthMethodListParamsApplicationIDAnthropic           CasbApplicationAuthMethodListParamsApplicationID = "ANTHROPIC"
	CasbApplicationAuthMethodListParamsApplicationIDAws                 CasbApplicationAuthMethodListParamsApplicationID = "AWS"
	CasbApplicationAuthMethodListParamsApplicationIDBitbucket           CasbApplicationAuthMethodListParamsApplicationID = "BITBUCKET"
	CasbApplicationAuthMethodListParamsApplicationIDBox                 CasbApplicationAuthMethodListParamsApplicationID = "BOX"
	CasbApplicationAuthMethodListParamsApplicationIDConfluence          CasbApplicationAuthMethodListParamsApplicationID = "CONFLUENCE"
	CasbApplicationAuthMethodListParamsApplicationIDDropbox             CasbApplicationAuthMethodListParamsApplicationID = "DROPBOX"
	CasbApplicationAuthMethodListParamsApplicationIDGitHub              CasbApplicationAuthMethodListParamsApplicationID = "GITHUB"
	CasbApplicationAuthMethodListParamsApplicationIDGoogleCloudPlatform CasbApplicationAuthMethodListParamsApplicationID = "GOOGLE_CLOUD_PLATFORM"
	CasbApplicationAuthMethodListParamsApplicationIDGoogleWorkspace     CasbApplicationAuthMethodListParamsApplicationID = "GOOGLE_WORKSPACE"
	CasbApplicationAuthMethodListParamsApplicationIDJira                CasbApplicationAuthMethodListParamsApplicationID = "JIRA"
	CasbApplicationAuthMethodListParamsApplicationIDMicrosoftInternal   CasbApplicationAuthMethodListParamsApplicationID = "MICROSOFT_INTERNAL"
	CasbApplicationAuthMethodListParamsApplicationIDOpenAI              CasbApplicationAuthMethodListParamsApplicationID = "OPENAI"
	CasbApplicationAuthMethodListParamsApplicationIDSalesforce          CasbApplicationAuthMethodListParamsApplicationID = "SALESFORCE"
	CasbApplicationAuthMethodListParamsApplicationIDServicenow          CasbApplicationAuthMethodListParamsApplicationID = "SERVICENOW"
	CasbApplicationAuthMethodListParamsApplicationIDSlack               CasbApplicationAuthMethodListParamsApplicationID = "SLACK"
)

func (r CasbApplicationAuthMethodListParamsApplicationID) IsKnown() bool {
	switch r {
	case CasbApplicationAuthMethodListParamsApplicationIDAnthropic, CasbApplicationAuthMethodListParamsApplicationIDAws, CasbApplicationAuthMethodListParamsApplicationIDBitbucket, CasbApplicationAuthMethodListParamsApplicationIDBox, CasbApplicationAuthMethodListParamsApplicationIDConfluence, CasbApplicationAuthMethodListParamsApplicationIDDropbox, CasbApplicationAuthMethodListParamsApplicationIDGitHub, CasbApplicationAuthMethodListParamsApplicationIDGoogleCloudPlatform, CasbApplicationAuthMethodListParamsApplicationIDGoogleWorkspace, CasbApplicationAuthMethodListParamsApplicationIDJira, CasbApplicationAuthMethodListParamsApplicationIDMicrosoftInternal, CasbApplicationAuthMethodListParamsApplicationIDOpenAI, CasbApplicationAuthMethodListParamsApplicationIDSalesforce, CasbApplicationAuthMethodListParamsApplicationIDServicenow, CasbApplicationAuthMethodListParamsApplicationIDSlack:
		return true
	}
	return false
}
