// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *CasbApplicationAuthMethodService) List(ctx context.Context, applicationID CasbApplicationAuthMethodListParamsApplicationID, query CasbApplicationAuthMethodListParams, opts ...option.RequestOption) (res *[]CasbApplicationAuthMethodListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/one/applications/%v/auth-methods", query.AccountID, applicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
}

type CasbApplicationAuthMethodListParamsApplicationID string

const (
	CasbApplicationAuthMethodListParamsApplicationIDBitbucket         CasbApplicationAuthMethodListParamsApplicationID = "BITBUCKET"
	CasbApplicationAuthMethodListParamsApplicationIDBox               CasbApplicationAuthMethodListParamsApplicationID = "BOX"
	CasbApplicationAuthMethodListParamsApplicationIDConfluence        CasbApplicationAuthMethodListParamsApplicationID = "CONFLUENCE"
	CasbApplicationAuthMethodListParamsApplicationIDDropbox           CasbApplicationAuthMethodListParamsApplicationID = "DROPBOX"
	CasbApplicationAuthMethodListParamsApplicationIDGitHub            CasbApplicationAuthMethodListParamsApplicationID = "GITHUB"
	CasbApplicationAuthMethodListParamsApplicationIDGoogleWorkspace   CasbApplicationAuthMethodListParamsApplicationID = "GOOGLE_WORKSPACE"
	CasbApplicationAuthMethodListParamsApplicationIDJira              CasbApplicationAuthMethodListParamsApplicationID = "JIRA"
	CasbApplicationAuthMethodListParamsApplicationIDMicrosoftInternal CasbApplicationAuthMethodListParamsApplicationID = "MICROSOFT_INTERNAL"
	CasbApplicationAuthMethodListParamsApplicationIDSalesforce        CasbApplicationAuthMethodListParamsApplicationID = "SALESFORCE"
	CasbApplicationAuthMethodListParamsApplicationIDSlack             CasbApplicationAuthMethodListParamsApplicationID = "SLACK"
)

func (r CasbApplicationAuthMethodListParamsApplicationID) IsKnown() bool {
	switch r {
	case CasbApplicationAuthMethodListParamsApplicationIDBitbucket, CasbApplicationAuthMethodListParamsApplicationIDBox, CasbApplicationAuthMethodListParamsApplicationIDConfluence, CasbApplicationAuthMethodListParamsApplicationIDDropbox, CasbApplicationAuthMethodListParamsApplicationIDGitHub, CasbApplicationAuthMethodListParamsApplicationIDGoogleWorkspace, CasbApplicationAuthMethodListParamsApplicationIDJira, CasbApplicationAuthMethodListParamsApplicationIDMicrosoftInternal, CasbApplicationAuthMethodListParamsApplicationIDSalesforce, CasbApplicationAuthMethodListParamsApplicationIDSlack:
		return true
	}
	return false
}
