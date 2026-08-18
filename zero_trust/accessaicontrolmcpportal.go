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

// AccessAIControlMcpPortalService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessAIControlMcpPortalService] method instead.
type AccessAIControlMcpPortalService struct {
	Options []option.RequestOption
}

// NewAccessAIControlMcpPortalService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessAIControlMcpPortalService(opts ...option.RequestOption) (r *AccessAIControlMcpPortalService) {
	r = &AccessAIControlMcpPortalService{}
	r.Options = opts
	return
}

// Creates a new MCP portal for managing AI tool access through Cloudflare Access.
func (r *AccessAIControlMcpPortalService) New(ctx context.Context, params AccessAIControlMcpPortalNewParams, opts ...option.RequestOption) (res *AccessAIControlMcpPortalNewResponse, err error) {
	var env AccessAIControlMcpPortalNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/portals", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an MCP portal configuration.
func (r *AccessAIControlMcpPortalService) Update(ctx context.Context, id string, params AccessAIControlMcpPortalUpdateParams, opts ...option.RequestOption) (res *AccessAIControlMcpPortalUpdateResponse, err error) {
	var env AccessAIControlMcpPortalUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/portals/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all MCP portals configured for the account.
func (r *AccessAIControlMcpPortalService) List(ctx context.Context, params AccessAIControlMcpPortalListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessAIControlMcpPortalListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/portals", params.AccountID)
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

// Lists all MCP portals configured for the account.
func (r *AccessAIControlMcpPortalService) ListAutoPaging(ctx context.Context, params AccessAIControlMcpPortalListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessAIControlMcpPortalListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an MCP portal from the account.
func (r *AccessAIControlMcpPortalService) Delete(ctx context.Context, id string, body AccessAIControlMcpPortalDeleteParams, opts ...option.RequestOption) (res *AccessAIControlMcpPortalDeleteResponse, err error) {
	var env AccessAIControlMcpPortalDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/portals/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Read the details of a single MCP Portal, including its configured servers.
func (r *AccessAIControlMcpPortalService) Read(ctx context.Context, id string, query AccessAIControlMcpPortalReadParams, opts ...option.RequestOption) (res *AccessAIControlMcpPortalReadResponse, err error) {
	var env AccessAIControlMcpPortalReadResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/portals/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AccessAIControlMcpPortalNewResponse struct {
	// Unique identifier for the MCP portal.
	ID string `json:"id" api:"required"`
	// Hostname where the MCP portal is available.
	Hostname string `json:"hostname" api:"required"`
	// Display name for the MCP portal.
	Name    string                                      `json:"name" api:"required"`
	Servers []AccessAIControlMcpPortalNewResponseServer `json:"servers" api:"required"`
	// Deprecated: use `code_mode` for new integrations. `true` maps to any non-off
	// Code Mode policy; `false` maps to `code_mode: off`. If both fields are sent,
	// they must be consistent or the request returns a 400.
	//
	// Deprecated: deprecated
	AllowCodeMode bool `json:"allow_code_mode"`
	// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
	// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
	// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
	// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
	// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
	// create. If both `code_mode` and `allow_code_mode` are sent, they must be
	// consistent or the request returns a 400.
	CodeMode  AccessAIControlMcpPortalNewResponseCodeMode `json:"code_mode"`
	CreatedAt time.Time                                   `json:"created_at" format:"date-time"`
	CreatedBy string                                      `json:"created_by"`
	// Optional description of the MCP portal.
	Description string    `json:"description"`
	ModifiedAt  time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy  string    `json:"modified_by"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway.
	SecureWebGateway bool                                    `json:"secure_web_gateway"`
	JSON             accessAIControlMcpPortalNewResponseJSON `json:"-"`
}

// accessAIControlMcpPortalNewResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpPortalNewResponse]
type accessAIControlMcpPortalNewResponseJSON struct {
	ID               apijson.Field
	Hostname         apijson.Field
	Name             apijson.Field
	Servers          apijson.Field
	AllowCodeMode    apijson.Field
	CodeMode         apijson.Field
	CreatedAt        apijson.Field
	CreatedBy        apijson.Field
	Description      apijson.Field
	ModifiedAt       apijson.Field
	ModifiedBy       apijson.Field
	SecureWebGateway apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalNewResponseServer struct {
	// Unique identifier for the MCP server.
	ID string `json:"id" api:"required"`
	// Authentication method used to connect to the upstream MCP server.
	AuthType AccessAIControlMcpPortalNewResponseServersAuthType `json:"auth_type" api:"required"`
	// URL of the upstream MCP endpoint.
	Hostname string `json:"hostname" api:"required" format:"uri"`
	// Display name for the MCP server.
	Name    string                   `json:"name" api:"required"`
	Prompts []map[string]interface{} `json:"prompts" api:"required"`
	// Unique identifier for the MCP server.
	ServerID string                   `json:"server_id" api:"required"`
	Tools    []map[string]interface{} `json:"tools" api:"required"`
	// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
	// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
	// endpoints + client_id for manual servers. Never includes the secret value.
	AuthConfigSummary AccessAIControlMcpPortalNewResponseServersAuthConfigSummary `json:"auth_config_summary"`
	// Whether administrative authentication is required before capabilities can be
	// synced. Manual OAuth is user-managed and has no administrative authentication
	// flow.
	AuthenticationStatus AccessAIControlMcpPortalNewResponseServersAuthenticationStatus `json:"authentication_status"`
	CreatedAt            time.Time                                                      `json:"created_at" format:"date-time"`
	CreatedBy            string                                                         `json:"created_by"`
	DefaultDisabled      bool                                                           `json:"default_disabled"`
	// Optional description of the MCP server.
	Description  string                                                 `json:"description" api:"nullable"`
	Error        string                                                 `json:"error"`
	ErrorDetails AccessAIControlMcpPortalNewResponseServersErrorDetails `json:"error_details"`
	// When true, the gateway worker uses the shared Cloudflare-owned OAuth callback
	// endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the
	// customer portal hostname. Defaults to false (off); opt in per server by setting
	// true.
	IsSharedOAuthCallbackEnabled bool      `json:"is_shared_oauth_callback_enabled"`
	LastSuccessfulSync           time.Time `json:"last_successful_sync" format:"date-time"`
	LastSynced                   time.Time `json:"last_synced" format:"date-time"`
	ModifiedAt                   time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy                   string    `json:"modified_by"`
	OnBehalf                     bool      `json:"on_behalf"`
	// Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway.
	SecureWebGateway bool `json:"secure_web_gateway"`
	// Current sync state of the server
	Status         AccessAIControlMcpPortalNewResponseServersStatus          `json:"status"`
	UpdatedPrompts []AccessAIControlMcpPortalNewResponseServersUpdatedPrompt `json:"updated_prompts"`
	UpdatedTools   []AccessAIControlMcpPortalNewResponseServersUpdatedTool   `json:"updated_tools"`
	JSON           accessAIControlMcpPortalNewResponseServerJSON             `json:"-"`
}

// accessAIControlMcpPortalNewResponseServerJSON contains the JSON metadata for the
// struct [AccessAIControlMcpPortalNewResponseServer]
type accessAIControlMcpPortalNewResponseServerJSON struct {
	ID                           apijson.Field
	AuthType                     apijson.Field
	Hostname                     apijson.Field
	Name                         apijson.Field
	Prompts                      apijson.Field
	ServerID                     apijson.Field
	Tools                        apijson.Field
	AuthConfigSummary            apijson.Field
	AuthenticationStatus         apijson.Field
	CreatedAt                    apijson.Field
	CreatedBy                    apijson.Field
	DefaultDisabled              apijson.Field
	Description                  apijson.Field
	Error                        apijson.Field
	ErrorDetails                 apijson.Field
	IsSharedOAuthCallbackEnabled apijson.Field
	LastSuccessfulSync           apijson.Field
	LastSynced                   apijson.Field
	ModifiedAt                   apijson.Field
	ModifiedBy                   apijson.Field
	OnBehalf                     apijson.Field
	SecureWebGateway             apijson.Field
	Status                       apijson.Field
	UpdatedPrompts               apijson.Field
	UpdatedTools                 apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalNewResponseServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalNewResponseServerJSON) RawJSON() string {
	return r.raw
}

// Authentication method used to connect to the upstream MCP server.
type AccessAIControlMcpPortalNewResponseServersAuthType string

const (
	AccessAIControlMcpPortalNewResponseServersAuthTypeOAuth           AccessAIControlMcpPortalNewResponseServersAuthType = "oauth"
	AccessAIControlMcpPortalNewResponseServersAuthTypeBearer          AccessAIControlMcpPortalNewResponseServersAuthType = "bearer"
	AccessAIControlMcpPortalNewResponseServersAuthTypeUnauthenticated AccessAIControlMcpPortalNewResponseServersAuthType = "unauthenticated"
)

func (r AccessAIControlMcpPortalNewResponseServersAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalNewResponseServersAuthTypeOAuth, AccessAIControlMcpPortalNewResponseServersAuthTypeBearer, AccessAIControlMcpPortalNewResponseServersAuthTypeUnauthenticated:
		return true
	}
	return false
}

// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
// endpoints + client_id for manual servers. Never includes the secret value.
type AccessAIControlMcpPortalNewResponseServersAuthConfigSummary struct {
	AuthMode            AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryAuthMode         `json:"auth_mode"`
	ClientSecretVersion float64                                                                     `json:"client_secret_version"`
	Config              AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryConfig           `json:"config"`
	HasClientSecret     bool                                                                        `json:"has_client_secret"`
	RegistrationInfo    AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryRegistrationInfo `json:"registration_info"`
	JSON                accessAIControlMcpPortalNewResponseServersAuthConfigSummaryJSON             `json:"-"`
}

// accessAIControlMcpPortalNewResponseServersAuthConfigSummaryJSON contains the
// JSON metadata for the struct
// [AccessAIControlMcpPortalNewResponseServersAuthConfigSummary]
type accessAIControlMcpPortalNewResponseServersAuthConfigSummaryJSON struct {
	AuthMode            apijson.Field
	ClientSecretVersion apijson.Field
	Config              apijson.Field
	HasClientSecret     apijson.Field
	RegistrationInfo    apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalNewResponseServersAuthConfigSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalNewResponseServersAuthConfigSummaryJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryAuthMode string

const (
	AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryAuthModeDcr    AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryAuthMode = "dcr"
	AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryAuthModeManual AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryAuthMode = "manual"
)

func (r AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryAuthMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryAuthModeDcr, AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryAuthModeManual:
		return true
	}
	return false
}

type AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryConfig struct {
	AuthorizationEndpoint string                                                                `json:"authorization_endpoint"`
	Issuer                string                                                                `json:"issuer"`
	Resource              string                                                                `json:"resource"`
	RevocationEndpoint    string                                                                `json:"revocation_endpoint"`
	TokenEndpoint         string                                                                `json:"token_endpoint"`
	JSON                  accessAIControlMcpPortalNewResponseServersAuthConfigSummaryConfigJSON `json:"-"`
}

// accessAIControlMcpPortalNewResponseServersAuthConfigSummaryConfigJSON contains
// the JSON metadata for the struct
// [AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryConfig]
type accessAIControlMcpPortalNewResponseServersAuthConfigSummaryConfigJSON struct {
	AuthorizationEndpoint apijson.Field
	Issuer                apijson.Field
	Resource              apijson.Field
	RevocationEndpoint    apijson.Field
	TokenEndpoint         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalNewResponseServersAuthConfigSummaryConfigJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryRegistrationInfo struct {
	ClientID                string                                                                          `json:"client_id"`
	RedirectURIs            []string                                                                        `json:"redirect_uris"`
	Scope                   string                                                                          `json:"scope"`
	TokenEndpointAuthMethod string                                                                          `json:"token_endpoint_auth_method"`
	JSON                    accessAIControlMcpPortalNewResponseServersAuthConfigSummaryRegistrationInfoJSON `json:"-"`
}

// accessAIControlMcpPortalNewResponseServersAuthConfigSummaryRegistrationInfoJSON
// contains the JSON metadata for the struct
// [AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryRegistrationInfo]
type accessAIControlMcpPortalNewResponseServersAuthConfigSummaryRegistrationInfoJSON struct {
	ClientID                apijson.Field
	RedirectURIs            apijson.Field
	Scope                   apijson.Field
	TokenEndpointAuthMethod apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalNewResponseServersAuthConfigSummaryRegistrationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalNewResponseServersAuthConfigSummaryRegistrationInfoJSON) RawJSON() string {
	return r.raw
}

// Whether administrative authentication is required before capabilities can be
// synced. Manual OAuth is user-managed and has no administrative authentication
// flow.
type AccessAIControlMcpPortalNewResponseServersAuthenticationStatus string

const (
	AccessAIControlMcpPortalNewResponseServersAuthenticationStatusNotRequired AccessAIControlMcpPortalNewResponseServersAuthenticationStatus = "not_required"
	AccessAIControlMcpPortalNewResponseServersAuthenticationStatusRequired    AccessAIControlMcpPortalNewResponseServersAuthenticationStatus = "required"
	AccessAIControlMcpPortalNewResponseServersAuthenticationStatusConnected   AccessAIControlMcpPortalNewResponseServersAuthenticationStatus = "connected"
	AccessAIControlMcpPortalNewResponseServersAuthenticationStatusStale       AccessAIControlMcpPortalNewResponseServersAuthenticationStatus = "stale"
	AccessAIControlMcpPortalNewResponseServersAuthenticationStatusManual      AccessAIControlMcpPortalNewResponseServersAuthenticationStatus = "manual"
)

func (r AccessAIControlMcpPortalNewResponseServersAuthenticationStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalNewResponseServersAuthenticationStatusNotRequired, AccessAIControlMcpPortalNewResponseServersAuthenticationStatusRequired, AccessAIControlMcpPortalNewResponseServersAuthenticationStatusConnected, AccessAIControlMcpPortalNewResponseServersAuthenticationStatusStale, AccessAIControlMcpPortalNewResponseServersAuthenticationStatusManual:
		return true
	}
	return false
}

type AccessAIControlMcpPortalNewResponseServersErrorDetails struct {
	// Underlying error message
	Cause string `json:"cause"`
	// True = MCP server returned an error. False = couldn't reach the server
	IsUpstream bool `json:"is_upstream"`
	// MCP protocol error code
	McpCode float64 `json:"mcp_code"`
	// Whether the error is transient and worth retrying
	Retryable bool `json:"retryable"`
	// HTTP status code from the server
	StatusCode float64                                                    `json:"status_code"`
	JSON       accessAIControlMcpPortalNewResponseServersErrorDetailsJSON `json:"-"`
}

// accessAIControlMcpPortalNewResponseServersErrorDetailsJSON contains the JSON
// metadata for the struct [AccessAIControlMcpPortalNewResponseServersErrorDetails]
type accessAIControlMcpPortalNewResponseServersErrorDetailsJSON struct {
	Cause       apijson.Field
	IsUpstream  apijson.Field
	McpCode     apijson.Field
	Retryable   apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalNewResponseServersErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalNewResponseServersErrorDetailsJSON) RawJSON() string {
	return r.raw
}

// Current sync state of the server
type AccessAIControlMcpPortalNewResponseServersStatus string

const (
	AccessAIControlMcpPortalNewResponseServersStatusWaiting AccessAIControlMcpPortalNewResponseServersStatus = "waiting"
	AccessAIControlMcpPortalNewResponseServersStatusReady   AccessAIControlMcpPortalNewResponseServersStatus = "ready"
	AccessAIControlMcpPortalNewResponseServersStatusStale   AccessAIControlMcpPortalNewResponseServersStatus = "stale"
	AccessAIControlMcpPortalNewResponseServersStatusError   AccessAIControlMcpPortalNewResponseServersStatus = "error"
)

func (r AccessAIControlMcpPortalNewResponseServersStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalNewResponseServersStatusWaiting, AccessAIControlMcpPortalNewResponseServersStatusReady, AccessAIControlMcpPortalNewResponseServersStatusStale, AccessAIControlMcpPortalNewResponseServersStatusError:
		return true
	}
	return false
}

type AccessAIControlMcpPortalNewResponseServersUpdatedPrompt struct {
	Name              string                                                      `json:"name" api:"required"`
	Enabled           bool                                                        `json:"enabled"`
	PortalAlias       string                                                      `json:"portal_alias"`
	PortalDescription string                                                      `json:"portal_description"`
	ServerAlias       string                                                      `json:"server_alias"`
	ServerDescription string                                                      `json:"server_description"`
	JSON              accessAIControlMcpPortalNewResponseServersUpdatedPromptJSON `json:"-"`
}

// accessAIControlMcpPortalNewResponseServersUpdatedPromptJSON contains the JSON
// metadata for the struct
// [AccessAIControlMcpPortalNewResponseServersUpdatedPrompt]
type accessAIControlMcpPortalNewResponseServersUpdatedPromptJSON struct {
	Name              apijson.Field
	Enabled           apijson.Field
	PortalAlias       apijson.Field
	PortalDescription apijson.Field
	ServerAlias       apijson.Field
	ServerDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalNewResponseServersUpdatedPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalNewResponseServersUpdatedPromptJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalNewResponseServersUpdatedTool struct {
	Name              string                                                    `json:"name" api:"required"`
	Enabled           bool                                                      `json:"enabled"`
	PortalAlias       string                                                    `json:"portal_alias"`
	PortalDescription string                                                    `json:"portal_description"`
	ServerAlias       string                                                    `json:"server_alias"`
	ServerDescription string                                                    `json:"server_description"`
	JSON              accessAIControlMcpPortalNewResponseServersUpdatedToolJSON `json:"-"`
}

// accessAIControlMcpPortalNewResponseServersUpdatedToolJSON contains the JSON
// metadata for the struct [AccessAIControlMcpPortalNewResponseServersUpdatedTool]
type accessAIControlMcpPortalNewResponseServersUpdatedToolJSON struct {
	Name              apijson.Field
	Enabled           apijson.Field
	PortalAlias       apijson.Field
	PortalDescription apijson.Field
	ServerAlias       apijson.Field
	ServerDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalNewResponseServersUpdatedTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalNewResponseServersUpdatedToolJSON) RawJSON() string {
	return r.raw
}

// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
// create. If both `code_mode` and `allow_code_mode` are sent, they must be
// consistent or the request returns a 400.
type AccessAIControlMcpPortalNewResponseCodeMode string

const (
	AccessAIControlMcpPortalNewResponseCodeModeOff       AccessAIControlMcpPortalNewResponseCodeMode = "off"
	AccessAIControlMcpPortalNewResponseCodeModeOptIn     AccessAIControlMcpPortalNewResponseCodeMode = "opt_in"
	AccessAIControlMcpPortalNewResponseCodeModeDefaultOn AccessAIControlMcpPortalNewResponseCodeMode = "default_on"
	AccessAIControlMcpPortalNewResponseCodeModeEnforced  AccessAIControlMcpPortalNewResponseCodeMode = "enforced"
)

func (r AccessAIControlMcpPortalNewResponseCodeMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalNewResponseCodeModeOff, AccessAIControlMcpPortalNewResponseCodeModeOptIn, AccessAIControlMcpPortalNewResponseCodeModeDefaultOn, AccessAIControlMcpPortalNewResponseCodeModeEnforced:
		return true
	}
	return false
}

type AccessAIControlMcpPortalUpdateResponse struct {
	// Unique identifier for the MCP portal.
	ID string `json:"id" api:"required"`
	// Hostname where the MCP portal is available.
	Hostname string `json:"hostname" api:"required"`
	// Display name for the MCP portal.
	Name    string                                         `json:"name" api:"required"`
	Servers []AccessAIControlMcpPortalUpdateResponseServer `json:"servers" api:"required"`
	// Deprecated: use `code_mode` for new integrations. `true` maps to any non-off
	// Code Mode policy; `false` maps to `code_mode: off`. If both fields are sent,
	// they must be consistent or the request returns a 400.
	//
	// Deprecated: deprecated
	AllowCodeMode bool `json:"allow_code_mode"`
	// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
	// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
	// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
	// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
	// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
	// create. If both `code_mode` and `allow_code_mode` are sent, they must be
	// consistent or the request returns a 400.
	CodeMode  AccessAIControlMcpPortalUpdateResponseCodeMode `json:"code_mode"`
	CreatedAt time.Time                                      `json:"created_at" format:"date-time"`
	CreatedBy string                                         `json:"created_by"`
	// Optional description of the MCP portal.
	Description string    `json:"description"`
	ModifiedAt  time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy  string    `json:"modified_by"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway.
	SecureWebGateway bool                                       `json:"secure_web_gateway"`
	JSON             accessAIControlMcpPortalUpdateResponseJSON `json:"-"`
}

// accessAIControlMcpPortalUpdateResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpPortalUpdateResponse]
type accessAIControlMcpPortalUpdateResponseJSON struct {
	ID               apijson.Field
	Hostname         apijson.Field
	Name             apijson.Field
	Servers          apijson.Field
	AllowCodeMode    apijson.Field
	CodeMode         apijson.Field
	CreatedAt        apijson.Field
	CreatedBy        apijson.Field
	Description      apijson.Field
	ModifiedAt       apijson.Field
	ModifiedBy       apijson.Field
	SecureWebGateway apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalUpdateResponseServer struct {
	// Unique identifier for the MCP server.
	ID string `json:"id" api:"required"`
	// Authentication method used to connect to the upstream MCP server.
	AuthType AccessAIControlMcpPortalUpdateResponseServersAuthType `json:"auth_type" api:"required"`
	// URL of the upstream MCP endpoint.
	Hostname string `json:"hostname" api:"required" format:"uri"`
	// Display name for the MCP server.
	Name    string                   `json:"name" api:"required"`
	Prompts []map[string]interface{} `json:"prompts" api:"required"`
	// Unique identifier for the MCP server.
	ServerID string                   `json:"server_id" api:"required"`
	Tools    []map[string]interface{} `json:"tools" api:"required"`
	// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
	// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
	// endpoints + client_id for manual servers. Never includes the secret value.
	AuthConfigSummary AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummary `json:"auth_config_summary"`
	// Whether administrative authentication is required before capabilities can be
	// synced. Manual OAuth is user-managed and has no administrative authentication
	// flow.
	AuthenticationStatus AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatus `json:"authentication_status"`
	CreatedAt            time.Time                                                         `json:"created_at" format:"date-time"`
	CreatedBy            string                                                            `json:"created_by"`
	DefaultDisabled      bool                                                              `json:"default_disabled"`
	// Optional description of the MCP server.
	Description  string                                                    `json:"description" api:"nullable"`
	Error        string                                                    `json:"error"`
	ErrorDetails AccessAIControlMcpPortalUpdateResponseServersErrorDetails `json:"error_details"`
	// When true, the gateway worker uses the shared Cloudflare-owned OAuth callback
	// endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the
	// customer portal hostname. Defaults to false (off); opt in per server by setting
	// true.
	IsSharedOAuthCallbackEnabled bool      `json:"is_shared_oauth_callback_enabled"`
	LastSuccessfulSync           time.Time `json:"last_successful_sync" format:"date-time"`
	LastSynced                   time.Time `json:"last_synced" format:"date-time"`
	ModifiedAt                   time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy                   string    `json:"modified_by"`
	OnBehalf                     bool      `json:"on_behalf"`
	// Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway.
	SecureWebGateway bool `json:"secure_web_gateway"`
	// Current sync state of the server
	Status         AccessAIControlMcpPortalUpdateResponseServersStatus          `json:"status"`
	UpdatedPrompts []AccessAIControlMcpPortalUpdateResponseServersUpdatedPrompt `json:"updated_prompts"`
	UpdatedTools   []AccessAIControlMcpPortalUpdateResponseServersUpdatedTool   `json:"updated_tools"`
	JSON           accessAIControlMcpPortalUpdateResponseServerJSON             `json:"-"`
}

// accessAIControlMcpPortalUpdateResponseServerJSON contains the JSON metadata for
// the struct [AccessAIControlMcpPortalUpdateResponseServer]
type accessAIControlMcpPortalUpdateResponseServerJSON struct {
	ID                           apijson.Field
	AuthType                     apijson.Field
	Hostname                     apijson.Field
	Name                         apijson.Field
	Prompts                      apijson.Field
	ServerID                     apijson.Field
	Tools                        apijson.Field
	AuthConfigSummary            apijson.Field
	AuthenticationStatus         apijson.Field
	CreatedAt                    apijson.Field
	CreatedBy                    apijson.Field
	DefaultDisabled              apijson.Field
	Description                  apijson.Field
	Error                        apijson.Field
	ErrorDetails                 apijson.Field
	IsSharedOAuthCallbackEnabled apijson.Field
	LastSuccessfulSync           apijson.Field
	LastSynced                   apijson.Field
	ModifiedAt                   apijson.Field
	ModifiedBy                   apijson.Field
	OnBehalf                     apijson.Field
	SecureWebGateway             apijson.Field
	Status                       apijson.Field
	UpdatedPrompts               apijson.Field
	UpdatedTools                 apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalUpdateResponseServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalUpdateResponseServerJSON) RawJSON() string {
	return r.raw
}

// Authentication method used to connect to the upstream MCP server.
type AccessAIControlMcpPortalUpdateResponseServersAuthType string

const (
	AccessAIControlMcpPortalUpdateResponseServersAuthTypeOAuth           AccessAIControlMcpPortalUpdateResponseServersAuthType = "oauth"
	AccessAIControlMcpPortalUpdateResponseServersAuthTypeBearer          AccessAIControlMcpPortalUpdateResponseServersAuthType = "bearer"
	AccessAIControlMcpPortalUpdateResponseServersAuthTypeUnauthenticated AccessAIControlMcpPortalUpdateResponseServersAuthType = "unauthenticated"
)

func (r AccessAIControlMcpPortalUpdateResponseServersAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalUpdateResponseServersAuthTypeOAuth, AccessAIControlMcpPortalUpdateResponseServersAuthTypeBearer, AccessAIControlMcpPortalUpdateResponseServersAuthTypeUnauthenticated:
		return true
	}
	return false
}

// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
// endpoints + client_id for manual servers. Never includes the secret value.
type AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummary struct {
	AuthMode            AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryAuthMode         `json:"auth_mode"`
	ClientSecretVersion float64                                                                        `json:"client_secret_version"`
	Config              AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryConfig           `json:"config"`
	HasClientSecret     bool                                                                           `json:"has_client_secret"`
	RegistrationInfo    AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryRegistrationInfo `json:"registration_info"`
	JSON                accessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryJSON             `json:"-"`
}

// accessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryJSON contains the
// JSON metadata for the struct
// [AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummary]
type accessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryJSON struct {
	AuthMode            apijson.Field
	ClientSecretVersion apijson.Field
	Config              apijson.Field
	HasClientSecret     apijson.Field
	RegistrationInfo    apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryAuthMode string

const (
	AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryAuthModeDcr    AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryAuthMode = "dcr"
	AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryAuthModeManual AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryAuthMode = "manual"
)

func (r AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryAuthMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryAuthModeDcr, AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryAuthModeManual:
		return true
	}
	return false
}

type AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryConfig struct {
	AuthorizationEndpoint string                                                                   `json:"authorization_endpoint"`
	Issuer                string                                                                   `json:"issuer"`
	Resource              string                                                                   `json:"resource"`
	RevocationEndpoint    string                                                                   `json:"revocation_endpoint"`
	TokenEndpoint         string                                                                   `json:"token_endpoint"`
	JSON                  accessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryConfigJSON `json:"-"`
}

// accessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryConfigJSON
// contains the JSON metadata for the struct
// [AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryConfig]
type accessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryConfigJSON struct {
	AuthorizationEndpoint apijson.Field
	Issuer                apijson.Field
	Resource              apijson.Field
	RevocationEndpoint    apijson.Field
	TokenEndpoint         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryConfigJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryRegistrationInfo struct {
	ClientID                string                                                                             `json:"client_id"`
	RedirectURIs            []string                                                                           `json:"redirect_uris"`
	Scope                   string                                                                             `json:"scope"`
	TokenEndpointAuthMethod string                                                                             `json:"token_endpoint_auth_method"`
	JSON                    accessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryRegistrationInfoJSON `json:"-"`
}

// accessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryRegistrationInfoJSON
// contains the JSON metadata for the struct
// [AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryRegistrationInfo]
type accessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryRegistrationInfoJSON struct {
	ClientID                apijson.Field
	RedirectURIs            apijson.Field
	Scope                   apijson.Field
	TokenEndpointAuthMethod apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryRegistrationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalUpdateResponseServersAuthConfigSummaryRegistrationInfoJSON) RawJSON() string {
	return r.raw
}

// Whether administrative authentication is required before capabilities can be
// synced. Manual OAuth is user-managed and has no administrative authentication
// flow.
type AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatus string

const (
	AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatusNotRequired AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatus = "not_required"
	AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatusRequired    AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatus = "required"
	AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatusConnected   AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatus = "connected"
	AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatusStale       AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatus = "stale"
	AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatusManual      AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatus = "manual"
)

func (r AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatusNotRequired, AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatusRequired, AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatusConnected, AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatusStale, AccessAIControlMcpPortalUpdateResponseServersAuthenticationStatusManual:
		return true
	}
	return false
}

type AccessAIControlMcpPortalUpdateResponseServersErrorDetails struct {
	// Underlying error message
	Cause string `json:"cause"`
	// True = MCP server returned an error. False = couldn't reach the server
	IsUpstream bool `json:"is_upstream"`
	// MCP protocol error code
	McpCode float64 `json:"mcp_code"`
	// Whether the error is transient and worth retrying
	Retryable bool `json:"retryable"`
	// HTTP status code from the server
	StatusCode float64                                                       `json:"status_code"`
	JSON       accessAIControlMcpPortalUpdateResponseServersErrorDetailsJSON `json:"-"`
}

// accessAIControlMcpPortalUpdateResponseServersErrorDetailsJSON contains the JSON
// metadata for the struct
// [AccessAIControlMcpPortalUpdateResponseServersErrorDetails]
type accessAIControlMcpPortalUpdateResponseServersErrorDetailsJSON struct {
	Cause       apijson.Field
	IsUpstream  apijson.Field
	McpCode     apijson.Field
	Retryable   apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalUpdateResponseServersErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalUpdateResponseServersErrorDetailsJSON) RawJSON() string {
	return r.raw
}

// Current sync state of the server
type AccessAIControlMcpPortalUpdateResponseServersStatus string

const (
	AccessAIControlMcpPortalUpdateResponseServersStatusWaiting AccessAIControlMcpPortalUpdateResponseServersStatus = "waiting"
	AccessAIControlMcpPortalUpdateResponseServersStatusReady   AccessAIControlMcpPortalUpdateResponseServersStatus = "ready"
	AccessAIControlMcpPortalUpdateResponseServersStatusStale   AccessAIControlMcpPortalUpdateResponseServersStatus = "stale"
	AccessAIControlMcpPortalUpdateResponseServersStatusError   AccessAIControlMcpPortalUpdateResponseServersStatus = "error"
)

func (r AccessAIControlMcpPortalUpdateResponseServersStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalUpdateResponseServersStatusWaiting, AccessAIControlMcpPortalUpdateResponseServersStatusReady, AccessAIControlMcpPortalUpdateResponseServersStatusStale, AccessAIControlMcpPortalUpdateResponseServersStatusError:
		return true
	}
	return false
}

type AccessAIControlMcpPortalUpdateResponseServersUpdatedPrompt struct {
	Name              string                                                         `json:"name" api:"required"`
	Enabled           bool                                                           `json:"enabled"`
	PortalAlias       string                                                         `json:"portal_alias"`
	PortalDescription string                                                         `json:"portal_description"`
	ServerAlias       string                                                         `json:"server_alias"`
	ServerDescription string                                                         `json:"server_description"`
	JSON              accessAIControlMcpPortalUpdateResponseServersUpdatedPromptJSON `json:"-"`
}

// accessAIControlMcpPortalUpdateResponseServersUpdatedPromptJSON contains the JSON
// metadata for the struct
// [AccessAIControlMcpPortalUpdateResponseServersUpdatedPrompt]
type accessAIControlMcpPortalUpdateResponseServersUpdatedPromptJSON struct {
	Name              apijson.Field
	Enabled           apijson.Field
	PortalAlias       apijson.Field
	PortalDescription apijson.Field
	ServerAlias       apijson.Field
	ServerDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalUpdateResponseServersUpdatedPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalUpdateResponseServersUpdatedPromptJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalUpdateResponseServersUpdatedTool struct {
	Name              string                                                       `json:"name" api:"required"`
	Enabled           bool                                                         `json:"enabled"`
	PortalAlias       string                                                       `json:"portal_alias"`
	PortalDescription string                                                       `json:"portal_description"`
	ServerAlias       string                                                       `json:"server_alias"`
	ServerDescription string                                                       `json:"server_description"`
	JSON              accessAIControlMcpPortalUpdateResponseServersUpdatedToolJSON `json:"-"`
}

// accessAIControlMcpPortalUpdateResponseServersUpdatedToolJSON contains the JSON
// metadata for the struct
// [AccessAIControlMcpPortalUpdateResponseServersUpdatedTool]
type accessAIControlMcpPortalUpdateResponseServersUpdatedToolJSON struct {
	Name              apijson.Field
	Enabled           apijson.Field
	PortalAlias       apijson.Field
	PortalDescription apijson.Field
	ServerAlias       apijson.Field
	ServerDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalUpdateResponseServersUpdatedTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalUpdateResponseServersUpdatedToolJSON) RawJSON() string {
	return r.raw
}

// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
// create. If both `code_mode` and `allow_code_mode` are sent, they must be
// consistent or the request returns a 400.
type AccessAIControlMcpPortalUpdateResponseCodeMode string

const (
	AccessAIControlMcpPortalUpdateResponseCodeModeOff       AccessAIControlMcpPortalUpdateResponseCodeMode = "off"
	AccessAIControlMcpPortalUpdateResponseCodeModeOptIn     AccessAIControlMcpPortalUpdateResponseCodeMode = "opt_in"
	AccessAIControlMcpPortalUpdateResponseCodeModeDefaultOn AccessAIControlMcpPortalUpdateResponseCodeMode = "default_on"
	AccessAIControlMcpPortalUpdateResponseCodeModeEnforced  AccessAIControlMcpPortalUpdateResponseCodeMode = "enforced"
)

func (r AccessAIControlMcpPortalUpdateResponseCodeMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalUpdateResponseCodeModeOff, AccessAIControlMcpPortalUpdateResponseCodeModeOptIn, AccessAIControlMcpPortalUpdateResponseCodeModeDefaultOn, AccessAIControlMcpPortalUpdateResponseCodeModeEnforced:
		return true
	}
	return false
}

type AccessAIControlMcpPortalListResponse struct {
	// Unique identifier for the MCP portal.
	ID string `json:"id" api:"required"`
	// Hostname where the MCP portal is available.
	Hostname string `json:"hostname" api:"required"`
	// Display name for the MCP portal.
	Name    string                                       `json:"name" api:"required"`
	Servers []AccessAIControlMcpPortalListResponseServer `json:"servers" api:"required"`
	// Deprecated: use `code_mode` for new integrations. `true` maps to any non-off
	// Code Mode policy; `false` maps to `code_mode: off`. If both fields are sent,
	// they must be consistent or the request returns a 400.
	//
	// Deprecated: deprecated
	AllowCodeMode bool `json:"allow_code_mode"`
	// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
	// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
	// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
	// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
	// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
	// create. If both `code_mode` and `allow_code_mode` are sent, they must be
	// consistent or the request returns a 400.
	CodeMode  AccessAIControlMcpPortalListResponseCodeMode `json:"code_mode"`
	CreatedAt time.Time                                    `json:"created_at" format:"date-time"`
	CreatedBy string                                       `json:"created_by"`
	// Optional description of the MCP portal.
	Description string    `json:"description"`
	ModifiedAt  time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy  string    `json:"modified_by"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway.
	SecureWebGateway bool                                     `json:"secure_web_gateway"`
	JSON             accessAIControlMcpPortalListResponseJSON `json:"-"`
}

// accessAIControlMcpPortalListResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpPortalListResponse]
type accessAIControlMcpPortalListResponseJSON struct {
	ID               apijson.Field
	Hostname         apijson.Field
	Name             apijson.Field
	Servers          apijson.Field
	AllowCodeMode    apijson.Field
	CodeMode         apijson.Field
	CreatedAt        apijson.Field
	CreatedBy        apijson.Field
	Description      apijson.Field
	ModifiedAt       apijson.Field
	ModifiedBy       apijson.Field
	SecureWebGateway apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalListResponseServer struct {
	// Unique identifier for the MCP server.
	ID string `json:"id" api:"required"`
	// Authentication method used to connect to the upstream MCP server.
	AuthType AccessAIControlMcpPortalListResponseServersAuthType `json:"auth_type" api:"required"`
	// URL of the upstream MCP endpoint.
	Hostname string `json:"hostname" api:"required" format:"uri"`
	// Display name for the MCP server.
	Name    string                   `json:"name" api:"required"`
	Prompts []map[string]interface{} `json:"prompts" api:"required"`
	// Unique identifier for the MCP server.
	ServerID string                   `json:"server_id" api:"required"`
	Tools    []map[string]interface{} `json:"tools" api:"required"`
	// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
	// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
	// endpoints + client_id for manual servers. Never includes the secret value.
	AuthConfigSummary AccessAIControlMcpPortalListResponseServersAuthConfigSummary `json:"auth_config_summary"`
	// Whether administrative authentication is required before capabilities can be
	// synced. Manual OAuth is user-managed and has no administrative authentication
	// flow.
	AuthenticationStatus AccessAIControlMcpPortalListResponseServersAuthenticationStatus `json:"authentication_status"`
	CreatedAt            time.Time                                                       `json:"created_at" format:"date-time"`
	CreatedBy            string                                                          `json:"created_by"`
	DefaultDisabled      bool                                                            `json:"default_disabled"`
	// Optional description of the MCP server.
	Description  string                                                  `json:"description" api:"nullable"`
	Error        string                                                  `json:"error"`
	ErrorDetails AccessAIControlMcpPortalListResponseServersErrorDetails `json:"error_details"`
	// When true, the gateway worker uses the shared Cloudflare-owned OAuth callback
	// endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the
	// customer portal hostname. Defaults to false (off); opt in per server by setting
	// true.
	IsSharedOAuthCallbackEnabled bool      `json:"is_shared_oauth_callback_enabled"`
	LastSuccessfulSync           time.Time `json:"last_successful_sync" format:"date-time"`
	LastSynced                   time.Time `json:"last_synced" format:"date-time"`
	ModifiedAt                   time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy                   string    `json:"modified_by"`
	OnBehalf                     bool      `json:"on_behalf"`
	// Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway.
	SecureWebGateway bool `json:"secure_web_gateway"`
	// Current sync state of the server
	Status         AccessAIControlMcpPortalListResponseServersStatus          `json:"status"`
	UpdatedPrompts []AccessAIControlMcpPortalListResponseServersUpdatedPrompt `json:"updated_prompts"`
	UpdatedTools   []AccessAIControlMcpPortalListResponseServersUpdatedTool   `json:"updated_tools"`
	JSON           accessAIControlMcpPortalListResponseServerJSON             `json:"-"`
}

// accessAIControlMcpPortalListResponseServerJSON contains the JSON metadata for
// the struct [AccessAIControlMcpPortalListResponseServer]
type accessAIControlMcpPortalListResponseServerJSON struct {
	ID                           apijson.Field
	AuthType                     apijson.Field
	Hostname                     apijson.Field
	Name                         apijson.Field
	Prompts                      apijson.Field
	ServerID                     apijson.Field
	Tools                        apijson.Field
	AuthConfigSummary            apijson.Field
	AuthenticationStatus         apijson.Field
	CreatedAt                    apijson.Field
	CreatedBy                    apijson.Field
	DefaultDisabled              apijson.Field
	Description                  apijson.Field
	Error                        apijson.Field
	ErrorDetails                 apijson.Field
	IsSharedOAuthCallbackEnabled apijson.Field
	LastSuccessfulSync           apijson.Field
	LastSynced                   apijson.Field
	ModifiedAt                   apijson.Field
	ModifiedBy                   apijson.Field
	OnBehalf                     apijson.Field
	SecureWebGateway             apijson.Field
	Status                       apijson.Field
	UpdatedPrompts               apijson.Field
	UpdatedTools                 apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalListResponseServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalListResponseServerJSON) RawJSON() string {
	return r.raw
}

// Authentication method used to connect to the upstream MCP server.
type AccessAIControlMcpPortalListResponseServersAuthType string

const (
	AccessAIControlMcpPortalListResponseServersAuthTypeOAuth           AccessAIControlMcpPortalListResponseServersAuthType = "oauth"
	AccessAIControlMcpPortalListResponseServersAuthTypeBearer          AccessAIControlMcpPortalListResponseServersAuthType = "bearer"
	AccessAIControlMcpPortalListResponseServersAuthTypeUnauthenticated AccessAIControlMcpPortalListResponseServersAuthType = "unauthenticated"
)

func (r AccessAIControlMcpPortalListResponseServersAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalListResponseServersAuthTypeOAuth, AccessAIControlMcpPortalListResponseServersAuthTypeBearer, AccessAIControlMcpPortalListResponseServersAuthTypeUnauthenticated:
		return true
	}
	return false
}

// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
// endpoints + client_id for manual servers. Never includes the secret value.
type AccessAIControlMcpPortalListResponseServersAuthConfigSummary struct {
	AuthMode            AccessAIControlMcpPortalListResponseServersAuthConfigSummaryAuthMode         `json:"auth_mode"`
	ClientSecretVersion float64                                                                      `json:"client_secret_version"`
	Config              AccessAIControlMcpPortalListResponseServersAuthConfigSummaryConfig           `json:"config"`
	HasClientSecret     bool                                                                         `json:"has_client_secret"`
	RegistrationInfo    AccessAIControlMcpPortalListResponseServersAuthConfigSummaryRegistrationInfo `json:"registration_info"`
	JSON                accessAIControlMcpPortalListResponseServersAuthConfigSummaryJSON             `json:"-"`
}

// accessAIControlMcpPortalListResponseServersAuthConfigSummaryJSON contains the
// JSON metadata for the struct
// [AccessAIControlMcpPortalListResponseServersAuthConfigSummary]
type accessAIControlMcpPortalListResponseServersAuthConfigSummaryJSON struct {
	AuthMode            apijson.Field
	ClientSecretVersion apijson.Field
	Config              apijson.Field
	HasClientSecret     apijson.Field
	RegistrationInfo    apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalListResponseServersAuthConfigSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalListResponseServersAuthConfigSummaryJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalListResponseServersAuthConfigSummaryAuthMode string

const (
	AccessAIControlMcpPortalListResponseServersAuthConfigSummaryAuthModeDcr    AccessAIControlMcpPortalListResponseServersAuthConfigSummaryAuthMode = "dcr"
	AccessAIControlMcpPortalListResponseServersAuthConfigSummaryAuthModeManual AccessAIControlMcpPortalListResponseServersAuthConfigSummaryAuthMode = "manual"
)

func (r AccessAIControlMcpPortalListResponseServersAuthConfigSummaryAuthMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalListResponseServersAuthConfigSummaryAuthModeDcr, AccessAIControlMcpPortalListResponseServersAuthConfigSummaryAuthModeManual:
		return true
	}
	return false
}

type AccessAIControlMcpPortalListResponseServersAuthConfigSummaryConfig struct {
	AuthorizationEndpoint string                                                                 `json:"authorization_endpoint"`
	Issuer                string                                                                 `json:"issuer"`
	Resource              string                                                                 `json:"resource"`
	RevocationEndpoint    string                                                                 `json:"revocation_endpoint"`
	TokenEndpoint         string                                                                 `json:"token_endpoint"`
	JSON                  accessAIControlMcpPortalListResponseServersAuthConfigSummaryConfigJSON `json:"-"`
}

// accessAIControlMcpPortalListResponseServersAuthConfigSummaryConfigJSON contains
// the JSON metadata for the struct
// [AccessAIControlMcpPortalListResponseServersAuthConfigSummaryConfig]
type accessAIControlMcpPortalListResponseServersAuthConfigSummaryConfigJSON struct {
	AuthorizationEndpoint apijson.Field
	Issuer                apijson.Field
	Resource              apijson.Field
	RevocationEndpoint    apijson.Field
	TokenEndpoint         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalListResponseServersAuthConfigSummaryConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalListResponseServersAuthConfigSummaryConfigJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalListResponseServersAuthConfigSummaryRegistrationInfo struct {
	ClientID                string                                                                           `json:"client_id"`
	RedirectURIs            []string                                                                         `json:"redirect_uris"`
	Scope                   string                                                                           `json:"scope"`
	TokenEndpointAuthMethod string                                                                           `json:"token_endpoint_auth_method"`
	JSON                    accessAIControlMcpPortalListResponseServersAuthConfigSummaryRegistrationInfoJSON `json:"-"`
}

// accessAIControlMcpPortalListResponseServersAuthConfigSummaryRegistrationInfoJSON
// contains the JSON metadata for the struct
// [AccessAIControlMcpPortalListResponseServersAuthConfigSummaryRegistrationInfo]
type accessAIControlMcpPortalListResponseServersAuthConfigSummaryRegistrationInfoJSON struct {
	ClientID                apijson.Field
	RedirectURIs            apijson.Field
	Scope                   apijson.Field
	TokenEndpointAuthMethod apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalListResponseServersAuthConfigSummaryRegistrationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalListResponseServersAuthConfigSummaryRegistrationInfoJSON) RawJSON() string {
	return r.raw
}

// Whether administrative authentication is required before capabilities can be
// synced. Manual OAuth is user-managed and has no administrative authentication
// flow.
type AccessAIControlMcpPortalListResponseServersAuthenticationStatus string

const (
	AccessAIControlMcpPortalListResponseServersAuthenticationStatusNotRequired AccessAIControlMcpPortalListResponseServersAuthenticationStatus = "not_required"
	AccessAIControlMcpPortalListResponseServersAuthenticationStatusRequired    AccessAIControlMcpPortalListResponseServersAuthenticationStatus = "required"
	AccessAIControlMcpPortalListResponseServersAuthenticationStatusConnected   AccessAIControlMcpPortalListResponseServersAuthenticationStatus = "connected"
	AccessAIControlMcpPortalListResponseServersAuthenticationStatusStale       AccessAIControlMcpPortalListResponseServersAuthenticationStatus = "stale"
	AccessAIControlMcpPortalListResponseServersAuthenticationStatusManual      AccessAIControlMcpPortalListResponseServersAuthenticationStatus = "manual"
)

func (r AccessAIControlMcpPortalListResponseServersAuthenticationStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalListResponseServersAuthenticationStatusNotRequired, AccessAIControlMcpPortalListResponseServersAuthenticationStatusRequired, AccessAIControlMcpPortalListResponseServersAuthenticationStatusConnected, AccessAIControlMcpPortalListResponseServersAuthenticationStatusStale, AccessAIControlMcpPortalListResponseServersAuthenticationStatusManual:
		return true
	}
	return false
}

type AccessAIControlMcpPortalListResponseServersErrorDetails struct {
	// Underlying error message
	Cause string `json:"cause"`
	// True = MCP server returned an error. False = couldn't reach the server
	IsUpstream bool `json:"is_upstream"`
	// MCP protocol error code
	McpCode float64 `json:"mcp_code"`
	// Whether the error is transient and worth retrying
	Retryable bool `json:"retryable"`
	// HTTP status code from the server
	StatusCode float64                                                     `json:"status_code"`
	JSON       accessAIControlMcpPortalListResponseServersErrorDetailsJSON `json:"-"`
}

// accessAIControlMcpPortalListResponseServersErrorDetailsJSON contains the JSON
// metadata for the struct
// [AccessAIControlMcpPortalListResponseServersErrorDetails]
type accessAIControlMcpPortalListResponseServersErrorDetailsJSON struct {
	Cause       apijson.Field
	IsUpstream  apijson.Field
	McpCode     apijson.Field
	Retryable   apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalListResponseServersErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalListResponseServersErrorDetailsJSON) RawJSON() string {
	return r.raw
}

// Current sync state of the server
type AccessAIControlMcpPortalListResponseServersStatus string

const (
	AccessAIControlMcpPortalListResponseServersStatusWaiting AccessAIControlMcpPortalListResponseServersStatus = "waiting"
	AccessAIControlMcpPortalListResponseServersStatusReady   AccessAIControlMcpPortalListResponseServersStatus = "ready"
	AccessAIControlMcpPortalListResponseServersStatusStale   AccessAIControlMcpPortalListResponseServersStatus = "stale"
	AccessAIControlMcpPortalListResponseServersStatusError   AccessAIControlMcpPortalListResponseServersStatus = "error"
)

func (r AccessAIControlMcpPortalListResponseServersStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalListResponseServersStatusWaiting, AccessAIControlMcpPortalListResponseServersStatusReady, AccessAIControlMcpPortalListResponseServersStatusStale, AccessAIControlMcpPortalListResponseServersStatusError:
		return true
	}
	return false
}

type AccessAIControlMcpPortalListResponseServersUpdatedPrompt struct {
	Name              string                                                       `json:"name" api:"required"`
	Enabled           bool                                                         `json:"enabled"`
	PortalAlias       string                                                       `json:"portal_alias"`
	PortalDescription string                                                       `json:"portal_description"`
	ServerAlias       string                                                       `json:"server_alias"`
	ServerDescription string                                                       `json:"server_description"`
	JSON              accessAIControlMcpPortalListResponseServersUpdatedPromptJSON `json:"-"`
}

// accessAIControlMcpPortalListResponseServersUpdatedPromptJSON contains the JSON
// metadata for the struct
// [AccessAIControlMcpPortalListResponseServersUpdatedPrompt]
type accessAIControlMcpPortalListResponseServersUpdatedPromptJSON struct {
	Name              apijson.Field
	Enabled           apijson.Field
	PortalAlias       apijson.Field
	PortalDescription apijson.Field
	ServerAlias       apijson.Field
	ServerDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalListResponseServersUpdatedPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalListResponseServersUpdatedPromptJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalListResponseServersUpdatedTool struct {
	Name              string                                                     `json:"name" api:"required"`
	Enabled           bool                                                       `json:"enabled"`
	PortalAlias       string                                                     `json:"portal_alias"`
	PortalDescription string                                                     `json:"portal_description"`
	ServerAlias       string                                                     `json:"server_alias"`
	ServerDescription string                                                     `json:"server_description"`
	JSON              accessAIControlMcpPortalListResponseServersUpdatedToolJSON `json:"-"`
}

// accessAIControlMcpPortalListResponseServersUpdatedToolJSON contains the JSON
// metadata for the struct [AccessAIControlMcpPortalListResponseServersUpdatedTool]
type accessAIControlMcpPortalListResponseServersUpdatedToolJSON struct {
	Name              apijson.Field
	Enabled           apijson.Field
	PortalAlias       apijson.Field
	PortalDescription apijson.Field
	ServerAlias       apijson.Field
	ServerDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalListResponseServersUpdatedTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalListResponseServersUpdatedToolJSON) RawJSON() string {
	return r.raw
}

// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
// create. If both `code_mode` and `allow_code_mode` are sent, they must be
// consistent or the request returns a 400.
type AccessAIControlMcpPortalListResponseCodeMode string

const (
	AccessAIControlMcpPortalListResponseCodeModeOff       AccessAIControlMcpPortalListResponseCodeMode = "off"
	AccessAIControlMcpPortalListResponseCodeModeOptIn     AccessAIControlMcpPortalListResponseCodeMode = "opt_in"
	AccessAIControlMcpPortalListResponseCodeModeDefaultOn AccessAIControlMcpPortalListResponseCodeMode = "default_on"
	AccessAIControlMcpPortalListResponseCodeModeEnforced  AccessAIControlMcpPortalListResponseCodeMode = "enforced"
)

func (r AccessAIControlMcpPortalListResponseCodeMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalListResponseCodeModeOff, AccessAIControlMcpPortalListResponseCodeModeOptIn, AccessAIControlMcpPortalListResponseCodeModeDefaultOn, AccessAIControlMcpPortalListResponseCodeModeEnforced:
		return true
	}
	return false
}

type AccessAIControlMcpPortalDeleteResponse struct {
	// Unique identifier for the MCP portal.
	ID string `json:"id" api:"required"`
	// Hostname where the MCP portal is available.
	Hostname string `json:"hostname" api:"required"`
	// Display name for the MCP portal.
	Name string `json:"name" api:"required"`
	// Deprecated: use `code_mode` for new integrations. `true` maps to any non-off
	// Code Mode policy; `false` maps to `code_mode: off`. If both fields are sent,
	// they must be consistent or the request returns a 400.
	//
	// Deprecated: deprecated
	AllowCodeMode bool `json:"allow_code_mode"`
	// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
	// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
	// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
	// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
	// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
	// create. If both `code_mode` and `allow_code_mode` are sent, they must be
	// consistent or the request returns a 400.
	CodeMode  AccessAIControlMcpPortalDeleteResponseCodeMode `json:"code_mode"`
	CreatedAt time.Time                                      `json:"created_at" format:"date-time"`
	CreatedBy string                                         `json:"created_by"`
	// Optional description of the MCP portal.
	Description string    `json:"description"`
	ModifiedAt  time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy  string    `json:"modified_by"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway.
	SecureWebGateway bool                                       `json:"secure_web_gateway"`
	JSON             accessAIControlMcpPortalDeleteResponseJSON `json:"-"`
}

// accessAIControlMcpPortalDeleteResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpPortalDeleteResponse]
type accessAIControlMcpPortalDeleteResponseJSON struct {
	ID               apijson.Field
	Hostname         apijson.Field
	Name             apijson.Field
	AllowCodeMode    apijson.Field
	CodeMode         apijson.Field
	CreatedAt        apijson.Field
	CreatedBy        apijson.Field
	Description      apijson.Field
	ModifiedAt       apijson.Field
	ModifiedBy       apijson.Field
	SecureWebGateway apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
// create. If both `code_mode` and `allow_code_mode` are sent, they must be
// consistent or the request returns a 400.
type AccessAIControlMcpPortalDeleteResponseCodeMode string

const (
	AccessAIControlMcpPortalDeleteResponseCodeModeOff       AccessAIControlMcpPortalDeleteResponseCodeMode = "off"
	AccessAIControlMcpPortalDeleteResponseCodeModeOptIn     AccessAIControlMcpPortalDeleteResponseCodeMode = "opt_in"
	AccessAIControlMcpPortalDeleteResponseCodeModeDefaultOn AccessAIControlMcpPortalDeleteResponseCodeMode = "default_on"
	AccessAIControlMcpPortalDeleteResponseCodeModeEnforced  AccessAIControlMcpPortalDeleteResponseCodeMode = "enforced"
)

func (r AccessAIControlMcpPortalDeleteResponseCodeMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalDeleteResponseCodeModeOff, AccessAIControlMcpPortalDeleteResponseCodeModeOptIn, AccessAIControlMcpPortalDeleteResponseCodeModeDefaultOn, AccessAIControlMcpPortalDeleteResponseCodeModeEnforced:
		return true
	}
	return false
}

type AccessAIControlMcpPortalReadResponse struct {
	// Unique identifier for the MCP portal.
	ID string `json:"id" api:"required"`
	// Hostname where the MCP portal is available.
	Hostname string `json:"hostname" api:"required"`
	// Display name for the MCP portal.
	Name    string                                       `json:"name" api:"required"`
	Servers []AccessAIControlMcpPortalReadResponseServer `json:"servers" api:"required"`
	// Deprecated: use `code_mode` for new integrations. `true` maps to any non-off
	// Code Mode policy; `false` maps to `code_mode: off`. If both fields are sent,
	// they must be consistent or the request returns a 400.
	//
	// Deprecated: deprecated
	AllowCodeMode bool `json:"allow_code_mode"`
	// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
	// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
	// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
	// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
	// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
	// create. If both `code_mode` and `allow_code_mode` are sent, they must be
	// consistent or the request returns a 400.
	CodeMode  AccessAIControlMcpPortalReadResponseCodeMode `json:"code_mode"`
	CreatedAt time.Time                                    `json:"created_at" format:"date-time"`
	CreatedBy string                                       `json:"created_by"`
	// Optional description of the MCP portal.
	Description string    `json:"description"`
	ModifiedAt  time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy  string    `json:"modified_by"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway.
	SecureWebGateway bool                                     `json:"secure_web_gateway"`
	JSON             accessAIControlMcpPortalReadResponseJSON `json:"-"`
}

// accessAIControlMcpPortalReadResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpPortalReadResponse]
type accessAIControlMcpPortalReadResponseJSON struct {
	ID               apijson.Field
	Hostname         apijson.Field
	Name             apijson.Field
	Servers          apijson.Field
	AllowCodeMode    apijson.Field
	CodeMode         apijson.Field
	CreatedAt        apijson.Field
	CreatedBy        apijson.Field
	Description      apijson.Field
	ModifiedAt       apijson.Field
	ModifiedBy       apijson.Field
	SecureWebGateway apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalReadResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalReadResponseServer struct {
	// Unique identifier for the MCP server.
	ID string `json:"id" api:"required"`
	// Authentication method used to connect to the upstream MCP server.
	AuthType AccessAIControlMcpPortalReadResponseServersAuthType `json:"auth_type" api:"required"`
	// URL of the upstream MCP endpoint.
	Hostname string `json:"hostname" api:"required" format:"uri"`
	// Display name for the MCP server.
	Name    string                   `json:"name" api:"required"`
	Prompts []map[string]interface{} `json:"prompts" api:"required"`
	// Unique identifier for the MCP server.
	ServerID string                   `json:"server_id" api:"required"`
	Tools    []map[string]interface{} `json:"tools" api:"required"`
	// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
	// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
	// endpoints + client_id for manual servers. Never includes the secret value.
	AuthConfigSummary AccessAIControlMcpPortalReadResponseServersAuthConfigSummary `json:"auth_config_summary"`
	// Whether administrative authentication is required before capabilities can be
	// synced. Manual OAuth is user-managed and has no administrative authentication
	// flow.
	AuthenticationStatus AccessAIControlMcpPortalReadResponseServersAuthenticationStatus `json:"authentication_status"`
	CreatedAt            time.Time                                                       `json:"created_at" format:"date-time"`
	CreatedBy            string                                                          `json:"created_by"`
	DefaultDisabled      bool                                                            `json:"default_disabled"`
	// Optional description of the MCP server.
	Description  string                                                  `json:"description" api:"nullable"`
	Error        string                                                  `json:"error"`
	ErrorDetails AccessAIControlMcpPortalReadResponseServersErrorDetails `json:"error_details"`
	// When true, the gateway worker uses the shared Cloudflare-owned OAuth callback
	// endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the
	// customer portal hostname. Defaults to false (off); opt in per server by setting
	// true.
	IsSharedOAuthCallbackEnabled bool      `json:"is_shared_oauth_callback_enabled"`
	LastSuccessfulSync           time.Time `json:"last_successful_sync" format:"date-time"`
	LastSynced                   time.Time `json:"last_synced" format:"date-time"`
	ModifiedAt                   time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy                   string    `json:"modified_by"`
	OnBehalf                     bool      `json:"on_behalf"`
	// Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway.
	SecureWebGateway bool `json:"secure_web_gateway"`
	// Current sync state of the server
	Status         AccessAIControlMcpPortalReadResponseServersStatus          `json:"status"`
	UpdatedPrompts []AccessAIControlMcpPortalReadResponseServersUpdatedPrompt `json:"updated_prompts"`
	UpdatedTools   []AccessAIControlMcpPortalReadResponseServersUpdatedTool   `json:"updated_tools"`
	JSON           accessAIControlMcpPortalReadResponseServerJSON             `json:"-"`
}

// accessAIControlMcpPortalReadResponseServerJSON contains the JSON metadata for
// the struct [AccessAIControlMcpPortalReadResponseServer]
type accessAIControlMcpPortalReadResponseServerJSON struct {
	ID                           apijson.Field
	AuthType                     apijson.Field
	Hostname                     apijson.Field
	Name                         apijson.Field
	Prompts                      apijson.Field
	ServerID                     apijson.Field
	Tools                        apijson.Field
	AuthConfigSummary            apijson.Field
	AuthenticationStatus         apijson.Field
	CreatedAt                    apijson.Field
	CreatedBy                    apijson.Field
	DefaultDisabled              apijson.Field
	Description                  apijson.Field
	Error                        apijson.Field
	ErrorDetails                 apijson.Field
	IsSharedOAuthCallbackEnabled apijson.Field
	LastSuccessfulSync           apijson.Field
	LastSynced                   apijson.Field
	ModifiedAt                   apijson.Field
	ModifiedBy                   apijson.Field
	OnBehalf                     apijson.Field
	SecureWebGateway             apijson.Field
	Status                       apijson.Field
	UpdatedPrompts               apijson.Field
	UpdatedTools                 apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalReadResponseServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalReadResponseServerJSON) RawJSON() string {
	return r.raw
}

// Authentication method used to connect to the upstream MCP server.
type AccessAIControlMcpPortalReadResponseServersAuthType string

const (
	AccessAIControlMcpPortalReadResponseServersAuthTypeOAuth           AccessAIControlMcpPortalReadResponseServersAuthType = "oauth"
	AccessAIControlMcpPortalReadResponseServersAuthTypeBearer          AccessAIControlMcpPortalReadResponseServersAuthType = "bearer"
	AccessAIControlMcpPortalReadResponseServersAuthTypeUnauthenticated AccessAIControlMcpPortalReadResponseServersAuthType = "unauthenticated"
)

func (r AccessAIControlMcpPortalReadResponseServersAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalReadResponseServersAuthTypeOAuth, AccessAIControlMcpPortalReadResponseServersAuthTypeBearer, AccessAIControlMcpPortalReadResponseServersAuthTypeUnauthenticated:
		return true
	}
	return false
}

// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
// endpoints + client_id for manual servers. Never includes the secret value.
type AccessAIControlMcpPortalReadResponseServersAuthConfigSummary struct {
	AuthMode            AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryAuthMode         `json:"auth_mode"`
	ClientSecretVersion float64                                                                      `json:"client_secret_version"`
	Config              AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryConfig           `json:"config"`
	HasClientSecret     bool                                                                         `json:"has_client_secret"`
	RegistrationInfo    AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryRegistrationInfo `json:"registration_info"`
	JSON                accessAIControlMcpPortalReadResponseServersAuthConfigSummaryJSON             `json:"-"`
}

// accessAIControlMcpPortalReadResponseServersAuthConfigSummaryJSON contains the
// JSON metadata for the struct
// [AccessAIControlMcpPortalReadResponseServersAuthConfigSummary]
type accessAIControlMcpPortalReadResponseServersAuthConfigSummaryJSON struct {
	AuthMode            apijson.Field
	ClientSecretVersion apijson.Field
	Config              apijson.Field
	HasClientSecret     apijson.Field
	RegistrationInfo    apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalReadResponseServersAuthConfigSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalReadResponseServersAuthConfigSummaryJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryAuthMode string

const (
	AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryAuthModeDcr    AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryAuthMode = "dcr"
	AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryAuthModeManual AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryAuthMode = "manual"
)

func (r AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryAuthMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryAuthModeDcr, AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryAuthModeManual:
		return true
	}
	return false
}

type AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryConfig struct {
	AuthorizationEndpoint string                                                                 `json:"authorization_endpoint"`
	Issuer                string                                                                 `json:"issuer"`
	Resource              string                                                                 `json:"resource"`
	RevocationEndpoint    string                                                                 `json:"revocation_endpoint"`
	TokenEndpoint         string                                                                 `json:"token_endpoint"`
	JSON                  accessAIControlMcpPortalReadResponseServersAuthConfigSummaryConfigJSON `json:"-"`
}

// accessAIControlMcpPortalReadResponseServersAuthConfigSummaryConfigJSON contains
// the JSON metadata for the struct
// [AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryConfig]
type accessAIControlMcpPortalReadResponseServersAuthConfigSummaryConfigJSON struct {
	AuthorizationEndpoint apijson.Field
	Issuer                apijson.Field
	Resource              apijson.Field
	RevocationEndpoint    apijson.Field
	TokenEndpoint         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalReadResponseServersAuthConfigSummaryConfigJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryRegistrationInfo struct {
	ClientID                string                                                                           `json:"client_id"`
	RedirectURIs            []string                                                                         `json:"redirect_uris"`
	Scope                   string                                                                           `json:"scope"`
	TokenEndpointAuthMethod string                                                                           `json:"token_endpoint_auth_method"`
	JSON                    accessAIControlMcpPortalReadResponseServersAuthConfigSummaryRegistrationInfoJSON `json:"-"`
}

// accessAIControlMcpPortalReadResponseServersAuthConfigSummaryRegistrationInfoJSON
// contains the JSON metadata for the struct
// [AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryRegistrationInfo]
type accessAIControlMcpPortalReadResponseServersAuthConfigSummaryRegistrationInfoJSON struct {
	ClientID                apijson.Field
	RedirectURIs            apijson.Field
	Scope                   apijson.Field
	TokenEndpointAuthMethod apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalReadResponseServersAuthConfigSummaryRegistrationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalReadResponseServersAuthConfigSummaryRegistrationInfoJSON) RawJSON() string {
	return r.raw
}

// Whether administrative authentication is required before capabilities can be
// synced. Manual OAuth is user-managed and has no administrative authentication
// flow.
type AccessAIControlMcpPortalReadResponseServersAuthenticationStatus string

const (
	AccessAIControlMcpPortalReadResponseServersAuthenticationStatusNotRequired AccessAIControlMcpPortalReadResponseServersAuthenticationStatus = "not_required"
	AccessAIControlMcpPortalReadResponseServersAuthenticationStatusRequired    AccessAIControlMcpPortalReadResponseServersAuthenticationStatus = "required"
	AccessAIControlMcpPortalReadResponseServersAuthenticationStatusConnected   AccessAIControlMcpPortalReadResponseServersAuthenticationStatus = "connected"
	AccessAIControlMcpPortalReadResponseServersAuthenticationStatusStale       AccessAIControlMcpPortalReadResponseServersAuthenticationStatus = "stale"
	AccessAIControlMcpPortalReadResponseServersAuthenticationStatusManual      AccessAIControlMcpPortalReadResponseServersAuthenticationStatus = "manual"
)

func (r AccessAIControlMcpPortalReadResponseServersAuthenticationStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalReadResponseServersAuthenticationStatusNotRequired, AccessAIControlMcpPortalReadResponseServersAuthenticationStatusRequired, AccessAIControlMcpPortalReadResponseServersAuthenticationStatusConnected, AccessAIControlMcpPortalReadResponseServersAuthenticationStatusStale, AccessAIControlMcpPortalReadResponseServersAuthenticationStatusManual:
		return true
	}
	return false
}

type AccessAIControlMcpPortalReadResponseServersErrorDetails struct {
	// Underlying error message
	Cause string `json:"cause"`
	// True = MCP server returned an error. False = couldn't reach the server
	IsUpstream bool `json:"is_upstream"`
	// MCP protocol error code
	McpCode float64 `json:"mcp_code"`
	// Whether the error is transient and worth retrying
	Retryable bool `json:"retryable"`
	// HTTP status code from the server
	StatusCode float64                                                     `json:"status_code"`
	JSON       accessAIControlMcpPortalReadResponseServersErrorDetailsJSON `json:"-"`
}

// accessAIControlMcpPortalReadResponseServersErrorDetailsJSON contains the JSON
// metadata for the struct
// [AccessAIControlMcpPortalReadResponseServersErrorDetails]
type accessAIControlMcpPortalReadResponseServersErrorDetailsJSON struct {
	Cause       apijson.Field
	IsUpstream  apijson.Field
	McpCode     apijson.Field
	Retryable   apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalReadResponseServersErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalReadResponseServersErrorDetailsJSON) RawJSON() string {
	return r.raw
}

// Current sync state of the server
type AccessAIControlMcpPortalReadResponseServersStatus string

const (
	AccessAIControlMcpPortalReadResponseServersStatusWaiting AccessAIControlMcpPortalReadResponseServersStatus = "waiting"
	AccessAIControlMcpPortalReadResponseServersStatusReady   AccessAIControlMcpPortalReadResponseServersStatus = "ready"
	AccessAIControlMcpPortalReadResponseServersStatusStale   AccessAIControlMcpPortalReadResponseServersStatus = "stale"
	AccessAIControlMcpPortalReadResponseServersStatusError   AccessAIControlMcpPortalReadResponseServersStatus = "error"
)

func (r AccessAIControlMcpPortalReadResponseServersStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalReadResponseServersStatusWaiting, AccessAIControlMcpPortalReadResponseServersStatusReady, AccessAIControlMcpPortalReadResponseServersStatusStale, AccessAIControlMcpPortalReadResponseServersStatusError:
		return true
	}
	return false
}

type AccessAIControlMcpPortalReadResponseServersUpdatedPrompt struct {
	Name              string                                                       `json:"name" api:"required"`
	Enabled           bool                                                         `json:"enabled"`
	PortalAlias       string                                                       `json:"portal_alias"`
	PortalDescription string                                                       `json:"portal_description"`
	ServerAlias       string                                                       `json:"server_alias"`
	ServerDescription string                                                       `json:"server_description"`
	JSON              accessAIControlMcpPortalReadResponseServersUpdatedPromptJSON `json:"-"`
}

// accessAIControlMcpPortalReadResponseServersUpdatedPromptJSON contains the JSON
// metadata for the struct
// [AccessAIControlMcpPortalReadResponseServersUpdatedPrompt]
type accessAIControlMcpPortalReadResponseServersUpdatedPromptJSON struct {
	Name              apijson.Field
	Enabled           apijson.Field
	PortalAlias       apijson.Field
	PortalDescription apijson.Field
	ServerAlias       apijson.Field
	ServerDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalReadResponseServersUpdatedPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalReadResponseServersUpdatedPromptJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalReadResponseServersUpdatedTool struct {
	Name              string                                                     `json:"name" api:"required"`
	Enabled           bool                                                       `json:"enabled"`
	PortalAlias       string                                                     `json:"portal_alias"`
	PortalDescription string                                                     `json:"portal_description"`
	ServerAlias       string                                                     `json:"server_alias"`
	ServerDescription string                                                     `json:"server_description"`
	JSON              accessAIControlMcpPortalReadResponseServersUpdatedToolJSON `json:"-"`
}

// accessAIControlMcpPortalReadResponseServersUpdatedToolJSON contains the JSON
// metadata for the struct [AccessAIControlMcpPortalReadResponseServersUpdatedTool]
type accessAIControlMcpPortalReadResponseServersUpdatedToolJSON struct {
	Name              apijson.Field
	Enabled           apijson.Field
	PortalAlias       apijson.Field
	PortalDescription apijson.Field
	ServerAlias       apijson.Field
	ServerDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalReadResponseServersUpdatedTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalReadResponseServersUpdatedToolJSON) RawJSON() string {
	return r.raw
}

// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
// create. If both `code_mode` and `allow_code_mode` are sent, they must be
// consistent or the request returns a 400.
type AccessAIControlMcpPortalReadResponseCodeMode string

const (
	AccessAIControlMcpPortalReadResponseCodeModeOff       AccessAIControlMcpPortalReadResponseCodeMode = "off"
	AccessAIControlMcpPortalReadResponseCodeModeOptIn     AccessAIControlMcpPortalReadResponseCodeMode = "opt_in"
	AccessAIControlMcpPortalReadResponseCodeModeDefaultOn AccessAIControlMcpPortalReadResponseCodeMode = "default_on"
	AccessAIControlMcpPortalReadResponseCodeModeEnforced  AccessAIControlMcpPortalReadResponseCodeMode = "enforced"
)

func (r AccessAIControlMcpPortalReadResponseCodeMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalReadResponseCodeModeOff, AccessAIControlMcpPortalReadResponseCodeModeOptIn, AccessAIControlMcpPortalReadResponseCodeModeDefaultOn, AccessAIControlMcpPortalReadResponseCodeModeEnforced:
		return true
	}
	return false
}

type AccessAIControlMcpPortalNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Unique identifier for the MCP portal.
	ID param.Field[string] `json:"id" api:"required"`
	// Hostname where the MCP portal is available.
	Hostname param.Field[string] `json:"hostname" api:"required"`
	// Display name for the MCP portal.
	Name param.Field[string] `json:"name" api:"required"`
	// Deprecated: use `code_mode` for new integrations. `true` maps to any non-off
	// Code Mode policy; `false` maps to `code_mode: off`. If both fields are sent,
	// they must be consistent or the request returns a 400.
	AllowCodeMode param.Field[bool] `json:"allow_code_mode"`
	// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
	// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
	// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
	// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
	// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
	// create. If both `code_mode` and `allow_code_mode` are sent, they must be
	// consistent or the request returns a 400.
	CodeMode param.Field[AccessAIControlMcpPortalNewParamsCodeMode] `json:"code_mode"`
	// Optional description of the MCP portal.
	Description param.Field[string] `json:"description"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway.
	SecureWebGateway param.Field[bool] `json:"secure_web_gateway"`
	// MCP servers attached to the portal and their portal-specific settings.
	Servers param.Field[[]AccessAIControlMcpPortalNewParamsServer] `json:"servers"`
}

func (r AccessAIControlMcpPortalNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
// create. If both `code_mode` and `allow_code_mode` are sent, they must be
// consistent or the request returns a 400.
type AccessAIControlMcpPortalNewParamsCodeMode string

const (
	AccessAIControlMcpPortalNewParamsCodeModeOff       AccessAIControlMcpPortalNewParamsCodeMode = "off"
	AccessAIControlMcpPortalNewParamsCodeModeOptIn     AccessAIControlMcpPortalNewParamsCodeMode = "opt_in"
	AccessAIControlMcpPortalNewParamsCodeModeDefaultOn AccessAIControlMcpPortalNewParamsCodeMode = "default_on"
	AccessAIControlMcpPortalNewParamsCodeModeEnforced  AccessAIControlMcpPortalNewParamsCodeMode = "enforced"
)

func (r AccessAIControlMcpPortalNewParamsCodeMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalNewParamsCodeModeOff, AccessAIControlMcpPortalNewParamsCodeModeOptIn, AccessAIControlMcpPortalNewParamsCodeModeDefaultOn, AccessAIControlMcpPortalNewParamsCodeModeEnforced:
		return true
	}
	return false
}

type AccessAIControlMcpPortalNewParamsServer struct {
	// Unique identifier for the MCP server.
	ServerID param.Field[string] `json:"server_id" api:"required"`
	// Disable this server by default for clients connecting through the portal.
	DefaultDisabled param.Field[bool] `json:"default_disabled"`
	// Use end-user OAuth credentials when connecting this server to the portal.
	OnBehalf param.Field[bool] `json:"on_behalf"`
	// Portal-specific prompt overrides.
	UpdatedPrompts param.Field[[]AccessAIControlMcpPortalNewParamsServersUpdatedPrompt] `json:"updated_prompts"`
	// Portal-specific tool overrides.
	UpdatedTools param.Field[[]AccessAIControlMcpPortalNewParamsServersUpdatedTool] `json:"updated_tools"`
}

func (r AccessAIControlMcpPortalNewParamsServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalNewParamsServersUpdatedPrompt struct {
	// Name of the tool or prompt capability to override.
	Name param.Field[string] `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias param.Field[string] `json:"alias"`
	// Custom description exposed for the capability.
	Description param.Field[string] `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccessAIControlMcpPortalNewParamsServersUpdatedPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalNewParamsServersUpdatedTool struct {
	// Name of the tool or prompt capability to override.
	Name param.Field[string] `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias param.Field[string] `json:"alias"`
	// Custom description exposed for the capability.
	Description param.Field[string] `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccessAIControlMcpPortalNewParamsServersUpdatedTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalNewResponseEnvelope struct {
	Result  AccessAIControlMcpPortalNewResponse             `json:"result" api:"required"`
	Success bool                                            `json:"success" api:"required"`
	JSON    accessAIControlMcpPortalNewResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpPortalNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessAIControlMcpPortalNewResponseEnvelope]
type accessAIControlMcpPortalNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalUpdateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Deprecated: use `code_mode` for new integrations. `true` maps to any non-off
	// Code Mode policy; `false` maps to `code_mode: off`. If both fields are sent,
	// they must be consistent or the request returns a 400.
	AllowCodeMode param.Field[bool] `json:"allow_code_mode"`
	// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
	// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
	// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
	// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
	// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
	// create. If both `code_mode` and `allow_code_mode` are sent, they must be
	// consistent or the request returns a 400.
	CodeMode param.Field[AccessAIControlMcpPortalUpdateParamsCodeMode] `json:"code_mode"`
	// Optional description of the MCP portal.
	Description param.Field[string] `json:"description"`
	// Hostname where the MCP portal is available.
	Hostname param.Field[string] `json:"hostname"`
	// Display name for the MCP portal.
	Name param.Field[string] `json:"name"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway.
	SecureWebGateway param.Field[bool] `json:"secure_web_gateway"`
	// MCP servers attached to the portal and their portal-specific settings.
	Servers param.Field[[]AccessAIControlMcpPortalUpdateParamsServer] `json:"servers"`
}

func (r AccessAIControlMcpPortalUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Code Mode policy for this portal. `off`: Code Mode is unavailable; query
// parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it
// on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by
// default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is
// always on; query parameters are ignored. Defaults to `opt_in` when omitted on
// create. If both `code_mode` and `allow_code_mode` are sent, they must be
// consistent or the request returns a 400.
type AccessAIControlMcpPortalUpdateParamsCodeMode string

const (
	AccessAIControlMcpPortalUpdateParamsCodeModeOff       AccessAIControlMcpPortalUpdateParamsCodeMode = "off"
	AccessAIControlMcpPortalUpdateParamsCodeModeOptIn     AccessAIControlMcpPortalUpdateParamsCodeMode = "opt_in"
	AccessAIControlMcpPortalUpdateParamsCodeModeDefaultOn AccessAIControlMcpPortalUpdateParamsCodeMode = "default_on"
	AccessAIControlMcpPortalUpdateParamsCodeModeEnforced  AccessAIControlMcpPortalUpdateParamsCodeMode = "enforced"
)

func (r AccessAIControlMcpPortalUpdateParamsCodeMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalUpdateParamsCodeModeOff, AccessAIControlMcpPortalUpdateParamsCodeModeOptIn, AccessAIControlMcpPortalUpdateParamsCodeModeDefaultOn, AccessAIControlMcpPortalUpdateParamsCodeModeEnforced:
		return true
	}
	return false
}

type AccessAIControlMcpPortalUpdateParamsServer struct {
	// Unique identifier for the MCP server.
	ServerID param.Field[string] `json:"server_id" api:"required"`
	// Disable this server by default for clients connecting through the portal.
	DefaultDisabled param.Field[bool] `json:"default_disabled"`
	// Use end-user OAuth credentials when connecting this server to the portal.
	OnBehalf param.Field[bool] `json:"on_behalf"`
	// Portal-specific prompt overrides.
	UpdatedPrompts param.Field[[]AccessAIControlMcpPortalUpdateParamsServersUpdatedPrompt] `json:"updated_prompts"`
	// Portal-specific tool overrides.
	UpdatedTools param.Field[[]AccessAIControlMcpPortalUpdateParamsServersUpdatedTool] `json:"updated_tools"`
}

func (r AccessAIControlMcpPortalUpdateParamsServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalUpdateParamsServersUpdatedPrompt struct {
	// Name of the tool or prompt capability to override.
	Name param.Field[string] `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias param.Field[string] `json:"alias"`
	// Custom description exposed for the capability.
	Description param.Field[string] `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccessAIControlMcpPortalUpdateParamsServersUpdatedPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalUpdateParamsServersUpdatedTool struct {
	// Name of the tool or prompt capability to override.
	Name param.Field[string] `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias param.Field[string] `json:"alias"`
	// Custom description exposed for the capability.
	Description param.Field[string] `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccessAIControlMcpPortalUpdateParamsServersUpdatedTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalUpdateResponseEnvelope struct {
	Result  AccessAIControlMcpPortalUpdateResponse             `json:"result" api:"required"`
	Success bool                                               `json:"success" api:"required"`
	JSON    accessAIControlMcpPortalUpdateResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpPortalUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessAIControlMcpPortalUpdateResponseEnvelope]
type accessAIControlMcpPortalUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
	// Search by id, name, hostname
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccessAIControlMcpPortalListParams]'s query parameters as
// `url.Values`.
func (r AccessAIControlMcpPortalListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessAIControlMcpPortalDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessAIControlMcpPortalDeleteResponseEnvelope struct {
	Result  AccessAIControlMcpPortalDeleteResponse             `json:"result" api:"required"`
	Success bool                                               `json:"success" api:"required"`
	JSON    accessAIControlMcpPortalDeleteResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpPortalDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessAIControlMcpPortalDeleteResponseEnvelope]
type accessAIControlMcpPortalDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalReadParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessAIControlMcpPortalReadResponseEnvelope struct {
	Result  AccessAIControlMcpPortalReadResponse             `json:"result" api:"required"`
	Success bool                                             `json:"success" api:"required"`
	JSON    accessAIControlMcpPortalReadResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpPortalReadResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessAIControlMcpPortalReadResponseEnvelope]
type accessAIControlMcpPortalReadResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalReadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalReadResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
