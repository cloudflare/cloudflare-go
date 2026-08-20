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

// AccessAIControlMcpServerService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessAIControlMcpServerService] method instead.
type AccessAIControlMcpServerService struct {
	Options []option.RequestOption
}

// NewAccessAIControlMcpServerService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessAIControlMcpServerService(opts ...option.RequestOption) (r *AccessAIControlMcpServerService) {
	r = &AccessAIControlMcpServerService{}
	r.Options = opts
	return
}

// Creates a new MCP server for connecting to an upstream MCP endpoint.
func (r *AccessAIControlMcpServerService) New(ctx context.Context, params AccessAIControlMcpServerNewParams, opts ...option.RequestOption) (res *AccessAIControlMcpServerNewResponse, err error) {
	var env AccessAIControlMcpServerNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/servers", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an MCP server's configuration and credentials.
func (r *AccessAIControlMcpServerService) Update(ctx context.Context, id string, params AccessAIControlMcpServerUpdateParams, opts ...option.RequestOption) (res *AccessAIControlMcpServerUpdateResponse, err error) {
	var env AccessAIControlMcpServerUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/servers/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all MCP servers configured for the account.
func (r *AccessAIControlMcpServerService) List(ctx context.Context, params AccessAIControlMcpServerListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessAIControlMcpServerListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/servers", params.AccountID)
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

// Lists all MCP servers configured for the account.
func (r *AccessAIControlMcpServerService) ListAutoPaging(ctx context.Context, params AccessAIControlMcpServerListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessAIControlMcpServerListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an MCP server from the account.
func (r *AccessAIControlMcpServerService) Delete(ctx context.Context, id string, body AccessAIControlMcpServerDeleteParams, opts ...option.RequestOption) (res *AccessAIControlMcpServerDeleteResponse, err error) {
	var env AccessAIControlMcpServerDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/servers/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves an MCP server's configuration and capability sync state.
func (r *AccessAIControlMcpServerService) Read(ctx context.Context, id string, query AccessAIControlMcpServerReadParams, opts ...option.RequestOption) (res *AccessAIControlMcpServerReadResponse, err error) {
	var env AccessAIControlMcpServerReadResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/servers/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Syncs an MCP server's capabilities and returns the updated server state,
// including any connection errors.
func (r *AccessAIControlMcpServerService) Sync(ctx context.Context, id string, body AccessAIControlMcpServerSyncParams, opts ...option.RequestOption) (res *AccessAIControlMcpServerSyncResponse, err error) {
	var env AccessAIControlMcpServerSyncResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/servers/%s/sync", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AccessAIControlMcpServerNewResponse struct {
	// Unique identifier for the MCP server.
	ID string `json:"id" api:"required"`
	// Authentication method used to connect to the upstream MCP server.
	AuthType AccessAIControlMcpServerNewResponseAuthType `json:"auth_type" api:"required"`
	// URL of the upstream MCP endpoint.
	Hostname string `json:"hostname" api:"required" format:"uri"`
	// Display name for the MCP server.
	Name    string                   `json:"name" api:"required"`
	Prompts []map[string]interface{} `json:"prompts" api:"required"`
	Tools   []map[string]interface{} `json:"tools" api:"required"`
	// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
	// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
	// endpoints + client_id for manual servers. Never includes the secret value.
	AuthConfigSummary AccessAIControlMcpServerNewResponseAuthConfigSummary `json:"auth_config_summary"`
	// Whether administrative authentication is required before capabilities can be
	// synced. Manual OAuth is user-managed and has no administrative authentication
	// flow.
	AuthenticationStatus AccessAIControlMcpServerNewResponseAuthenticationStatus `json:"authentication_status"`
	CreatedAt            time.Time                                               `json:"created_at" format:"date-time"`
	CreatedBy            string                                                  `json:"created_by"`
	// Optional description of the MCP server.
	Description  string                                          `json:"description" api:"nullable"`
	Error        string                                          `json:"error"`
	ErrorDetails AccessAIControlMcpServerNewResponseErrorDetails `json:"error_details"`
	// When true, the gateway worker uses the shared Cloudflare-owned OAuth callback
	// endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the
	// customer portal hostname. Defaults to false (off); opt in per server by setting
	// true.
	IsSharedOAuthCallbackEnabled bool      `json:"is_shared_oauth_callback_enabled"`
	LastSuccessfulSync           time.Time `json:"last_successful_sync" format:"date-time"`
	LastSynced                   time.Time `json:"last_synced" format:"date-time"`
	ModifiedAt                   time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy                   string    `json:"modified_by"`
	// Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway.
	SecureWebGateway bool `json:"secure_web_gateway"`
	// Current sync state of the server
	Status AccessAIControlMcpServerNewResponseStatus `json:"status"`
	// Server-wide prompt capability overrides.
	UpdatedPrompts []AccessAIControlMcpServerNewResponseUpdatedPrompt `json:"updated_prompts"`
	// Server-wide tool capability overrides.
	UpdatedTools []AccessAIControlMcpServerNewResponseUpdatedTool `json:"updated_tools"`
	JSON         accessAIControlMcpServerNewResponseJSON          `json:"-"`
}

// accessAIControlMcpServerNewResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpServerNewResponse]
type accessAIControlMcpServerNewResponseJSON struct {
	ID                           apijson.Field
	AuthType                     apijson.Field
	Hostname                     apijson.Field
	Name                         apijson.Field
	Prompts                      apijson.Field
	Tools                        apijson.Field
	AuthConfigSummary            apijson.Field
	AuthenticationStatus         apijson.Field
	CreatedAt                    apijson.Field
	CreatedBy                    apijson.Field
	Description                  apijson.Field
	Error                        apijson.Field
	ErrorDetails                 apijson.Field
	IsSharedOAuthCallbackEnabled apijson.Field
	LastSuccessfulSync           apijson.Field
	LastSynced                   apijson.Field
	ModifiedAt                   apijson.Field
	ModifiedBy                   apijson.Field
	SecureWebGateway             apijson.Field
	Status                       apijson.Field
	UpdatedPrompts               apijson.Field
	UpdatedTools                 apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccessAIControlMcpServerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerNewResponseJSON) RawJSON() string {
	return r.raw
}

// Authentication method used to connect to the upstream MCP server.
type AccessAIControlMcpServerNewResponseAuthType string

const (
	AccessAIControlMcpServerNewResponseAuthTypeOAuth           AccessAIControlMcpServerNewResponseAuthType = "oauth"
	AccessAIControlMcpServerNewResponseAuthTypeBearer          AccessAIControlMcpServerNewResponseAuthType = "bearer"
	AccessAIControlMcpServerNewResponseAuthTypeUnauthenticated AccessAIControlMcpServerNewResponseAuthType = "unauthenticated"
)

func (r AccessAIControlMcpServerNewResponseAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerNewResponseAuthTypeOAuth, AccessAIControlMcpServerNewResponseAuthTypeBearer, AccessAIControlMcpServerNewResponseAuthTypeUnauthenticated:
		return true
	}
	return false
}

// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
// endpoints + client_id for manual servers. Never includes the secret value.
type AccessAIControlMcpServerNewResponseAuthConfigSummary struct {
	AuthMode            AccessAIControlMcpServerNewResponseAuthConfigSummaryAuthMode         `json:"auth_mode"`
	ClientSecretVersion float64                                                              `json:"client_secret_version"`
	Config              AccessAIControlMcpServerNewResponseAuthConfigSummaryConfig           `json:"config"`
	HasClientSecret     bool                                                                 `json:"has_client_secret"`
	RegistrationInfo    AccessAIControlMcpServerNewResponseAuthConfigSummaryRegistrationInfo `json:"registration_info"`
	JSON                accessAIControlMcpServerNewResponseAuthConfigSummaryJSON             `json:"-"`
}

// accessAIControlMcpServerNewResponseAuthConfigSummaryJSON contains the JSON
// metadata for the struct [AccessAIControlMcpServerNewResponseAuthConfigSummary]
type accessAIControlMcpServerNewResponseAuthConfigSummaryJSON struct {
	AuthMode            apijson.Field
	ClientSecretVersion apijson.Field
	Config              apijson.Field
	HasClientSecret     apijson.Field
	RegistrationInfo    apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessAIControlMcpServerNewResponseAuthConfigSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerNewResponseAuthConfigSummaryJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerNewResponseAuthConfigSummaryAuthMode string

const (
	AccessAIControlMcpServerNewResponseAuthConfigSummaryAuthModeDcr    AccessAIControlMcpServerNewResponseAuthConfigSummaryAuthMode = "dcr"
	AccessAIControlMcpServerNewResponseAuthConfigSummaryAuthModeManual AccessAIControlMcpServerNewResponseAuthConfigSummaryAuthMode = "manual"
)

func (r AccessAIControlMcpServerNewResponseAuthConfigSummaryAuthMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerNewResponseAuthConfigSummaryAuthModeDcr, AccessAIControlMcpServerNewResponseAuthConfigSummaryAuthModeManual:
		return true
	}
	return false
}

type AccessAIControlMcpServerNewResponseAuthConfigSummaryConfig struct {
	AuthorizationEndpoint string                                                         `json:"authorization_endpoint"`
	Issuer                string                                                         `json:"issuer"`
	Resource              string                                                         `json:"resource"`
	RevocationEndpoint    string                                                         `json:"revocation_endpoint"`
	TokenEndpoint         string                                                         `json:"token_endpoint"`
	JSON                  accessAIControlMcpServerNewResponseAuthConfigSummaryConfigJSON `json:"-"`
}

// accessAIControlMcpServerNewResponseAuthConfigSummaryConfigJSON contains the JSON
// metadata for the struct
// [AccessAIControlMcpServerNewResponseAuthConfigSummaryConfig]
type accessAIControlMcpServerNewResponseAuthConfigSummaryConfigJSON struct {
	AuthorizationEndpoint apijson.Field
	Issuer                apijson.Field
	Resource              apijson.Field
	RevocationEndpoint    apijson.Field
	TokenEndpoint         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessAIControlMcpServerNewResponseAuthConfigSummaryConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerNewResponseAuthConfigSummaryConfigJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerNewResponseAuthConfigSummaryRegistrationInfo struct {
	ClientID                string                                                                   `json:"client_id"`
	RedirectURIs            []string                                                                 `json:"redirect_uris"`
	Scope                   string                                                                   `json:"scope"`
	TokenEndpointAuthMethod string                                                                   `json:"token_endpoint_auth_method"`
	JSON                    accessAIControlMcpServerNewResponseAuthConfigSummaryRegistrationInfoJSON `json:"-"`
}

// accessAIControlMcpServerNewResponseAuthConfigSummaryRegistrationInfoJSON
// contains the JSON metadata for the struct
// [AccessAIControlMcpServerNewResponseAuthConfigSummaryRegistrationInfo]
type accessAIControlMcpServerNewResponseAuthConfigSummaryRegistrationInfoJSON struct {
	ClientID                apijson.Field
	RedirectURIs            apijson.Field
	Scope                   apijson.Field
	TokenEndpointAuthMethod apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessAIControlMcpServerNewResponseAuthConfigSummaryRegistrationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerNewResponseAuthConfigSummaryRegistrationInfoJSON) RawJSON() string {
	return r.raw
}

// Whether administrative authentication is required before capabilities can be
// synced. Manual OAuth is user-managed and has no administrative authentication
// flow.
type AccessAIControlMcpServerNewResponseAuthenticationStatus string

const (
	AccessAIControlMcpServerNewResponseAuthenticationStatusNotRequired AccessAIControlMcpServerNewResponseAuthenticationStatus = "not_required"
	AccessAIControlMcpServerNewResponseAuthenticationStatusRequired    AccessAIControlMcpServerNewResponseAuthenticationStatus = "required"
	AccessAIControlMcpServerNewResponseAuthenticationStatusConnected   AccessAIControlMcpServerNewResponseAuthenticationStatus = "connected"
	AccessAIControlMcpServerNewResponseAuthenticationStatusStale       AccessAIControlMcpServerNewResponseAuthenticationStatus = "stale"
	AccessAIControlMcpServerNewResponseAuthenticationStatusManual      AccessAIControlMcpServerNewResponseAuthenticationStatus = "manual"
)

func (r AccessAIControlMcpServerNewResponseAuthenticationStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerNewResponseAuthenticationStatusNotRequired, AccessAIControlMcpServerNewResponseAuthenticationStatusRequired, AccessAIControlMcpServerNewResponseAuthenticationStatusConnected, AccessAIControlMcpServerNewResponseAuthenticationStatusStale, AccessAIControlMcpServerNewResponseAuthenticationStatusManual:
		return true
	}
	return false
}

type AccessAIControlMcpServerNewResponseErrorDetails struct {
	// Underlying error message
	Cause string `json:"cause"`
	// True = MCP server returned an error. False = couldn't reach the server
	IsUpstream bool `json:"is_upstream"`
	// MCP protocol error code
	McpCode float64 `json:"mcp_code"`
	// Whether the error is transient and worth retrying
	Retryable bool `json:"retryable"`
	// HTTP status code from the server
	StatusCode float64                                             `json:"status_code"`
	JSON       accessAIControlMcpServerNewResponseErrorDetailsJSON `json:"-"`
}

// accessAIControlMcpServerNewResponseErrorDetailsJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerNewResponseErrorDetails]
type accessAIControlMcpServerNewResponseErrorDetailsJSON struct {
	Cause       apijson.Field
	IsUpstream  apijson.Field
	McpCode     apijson.Field
	Retryable   apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerNewResponseErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerNewResponseErrorDetailsJSON) RawJSON() string {
	return r.raw
}

// Current sync state of the server
type AccessAIControlMcpServerNewResponseStatus string

const (
	AccessAIControlMcpServerNewResponseStatusWaiting AccessAIControlMcpServerNewResponseStatus = "waiting"
	AccessAIControlMcpServerNewResponseStatusReady   AccessAIControlMcpServerNewResponseStatus = "ready"
	AccessAIControlMcpServerNewResponseStatusStale   AccessAIControlMcpServerNewResponseStatus = "stale"
	AccessAIControlMcpServerNewResponseStatusError   AccessAIControlMcpServerNewResponseStatus = "error"
)

func (r AccessAIControlMcpServerNewResponseStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerNewResponseStatusWaiting, AccessAIControlMcpServerNewResponseStatusReady, AccessAIControlMcpServerNewResponseStatusStale, AccessAIControlMcpServerNewResponseStatusError:
		return true
	}
	return false
}

type AccessAIControlMcpServerNewResponseUpdatedPrompt struct {
	// Name of the tool or prompt capability to override.
	Name string `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias string `json:"alias"`
	// Custom description exposed for the capability.
	Description string `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled bool                                                 `json:"enabled"`
	JSON    accessAIControlMcpServerNewResponseUpdatedPromptJSON `json:"-"`
}

// accessAIControlMcpServerNewResponseUpdatedPromptJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerNewResponseUpdatedPrompt]
type accessAIControlMcpServerNewResponseUpdatedPromptJSON struct {
	Name        apijson.Field
	Alias       apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerNewResponseUpdatedPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerNewResponseUpdatedPromptJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerNewResponseUpdatedTool struct {
	// Name of the tool or prompt capability to override.
	Name string `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias string `json:"alias"`
	// Custom description exposed for the capability.
	Description string `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled bool                                               `json:"enabled"`
	JSON    accessAIControlMcpServerNewResponseUpdatedToolJSON `json:"-"`
}

// accessAIControlMcpServerNewResponseUpdatedToolJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerNewResponseUpdatedTool]
type accessAIControlMcpServerNewResponseUpdatedToolJSON struct {
	Name        apijson.Field
	Alias       apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerNewResponseUpdatedTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerNewResponseUpdatedToolJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerUpdateResponse struct {
	// Unique identifier for the MCP server.
	ID string `json:"id" api:"required"`
	// Authentication method used to connect to the upstream MCP server.
	AuthType AccessAIControlMcpServerUpdateResponseAuthType `json:"auth_type" api:"required"`
	// URL of the upstream MCP endpoint.
	Hostname string `json:"hostname" api:"required" format:"uri"`
	// Display name for the MCP server.
	Name    string                   `json:"name" api:"required"`
	Prompts []map[string]interface{} `json:"prompts" api:"required"`
	Tools   []map[string]interface{} `json:"tools" api:"required"`
	// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
	// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
	// endpoints + client_id for manual servers. Never includes the secret value.
	AuthConfigSummary AccessAIControlMcpServerUpdateResponseAuthConfigSummary `json:"auth_config_summary"`
	// Whether administrative authentication is required before capabilities can be
	// synced. Manual OAuth is user-managed and has no administrative authentication
	// flow.
	AuthenticationStatus AccessAIControlMcpServerUpdateResponseAuthenticationStatus `json:"authentication_status"`
	CreatedAt            time.Time                                                  `json:"created_at" format:"date-time"`
	CreatedBy            string                                                     `json:"created_by"`
	// Optional description of the MCP server.
	Description  string                                             `json:"description" api:"nullable"`
	Error        string                                             `json:"error"`
	ErrorDetails AccessAIControlMcpServerUpdateResponseErrorDetails `json:"error_details"`
	// When true, the gateway worker uses the shared Cloudflare-owned OAuth callback
	// endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the
	// customer portal hostname. Defaults to false (off); opt in per server by setting
	// true.
	IsSharedOAuthCallbackEnabled bool      `json:"is_shared_oauth_callback_enabled"`
	LastSuccessfulSync           time.Time `json:"last_successful_sync" format:"date-time"`
	LastSynced                   time.Time `json:"last_synced" format:"date-time"`
	ModifiedAt                   time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy                   string    `json:"modified_by"`
	// Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway.
	SecureWebGateway bool `json:"secure_web_gateway"`
	// Current sync state of the server
	Status AccessAIControlMcpServerUpdateResponseStatus `json:"status"`
	// Server-wide prompt capability overrides.
	UpdatedPrompts []AccessAIControlMcpServerUpdateResponseUpdatedPrompt `json:"updated_prompts"`
	// Server-wide tool capability overrides.
	UpdatedTools []AccessAIControlMcpServerUpdateResponseUpdatedTool `json:"updated_tools"`
	JSON         accessAIControlMcpServerUpdateResponseJSON          `json:"-"`
}

// accessAIControlMcpServerUpdateResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpServerUpdateResponse]
type accessAIControlMcpServerUpdateResponseJSON struct {
	ID                           apijson.Field
	AuthType                     apijson.Field
	Hostname                     apijson.Field
	Name                         apijson.Field
	Prompts                      apijson.Field
	Tools                        apijson.Field
	AuthConfigSummary            apijson.Field
	AuthenticationStatus         apijson.Field
	CreatedAt                    apijson.Field
	CreatedBy                    apijson.Field
	Description                  apijson.Field
	Error                        apijson.Field
	ErrorDetails                 apijson.Field
	IsSharedOAuthCallbackEnabled apijson.Field
	LastSuccessfulSync           apijson.Field
	LastSynced                   apijson.Field
	ModifiedAt                   apijson.Field
	ModifiedBy                   apijson.Field
	SecureWebGateway             apijson.Field
	Status                       apijson.Field
	UpdatedPrompts               apijson.Field
	UpdatedTools                 apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccessAIControlMcpServerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Authentication method used to connect to the upstream MCP server.
type AccessAIControlMcpServerUpdateResponseAuthType string

const (
	AccessAIControlMcpServerUpdateResponseAuthTypeOAuth           AccessAIControlMcpServerUpdateResponseAuthType = "oauth"
	AccessAIControlMcpServerUpdateResponseAuthTypeBearer          AccessAIControlMcpServerUpdateResponseAuthType = "bearer"
	AccessAIControlMcpServerUpdateResponseAuthTypeUnauthenticated AccessAIControlMcpServerUpdateResponseAuthType = "unauthenticated"
)

func (r AccessAIControlMcpServerUpdateResponseAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerUpdateResponseAuthTypeOAuth, AccessAIControlMcpServerUpdateResponseAuthTypeBearer, AccessAIControlMcpServerUpdateResponseAuthTypeUnauthenticated:
		return true
	}
	return false
}

// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
// endpoints + client_id for manual servers. Never includes the secret value.
type AccessAIControlMcpServerUpdateResponseAuthConfigSummary struct {
	AuthMode            AccessAIControlMcpServerUpdateResponseAuthConfigSummaryAuthMode         `json:"auth_mode"`
	ClientSecretVersion float64                                                                 `json:"client_secret_version"`
	Config              AccessAIControlMcpServerUpdateResponseAuthConfigSummaryConfig           `json:"config"`
	HasClientSecret     bool                                                                    `json:"has_client_secret"`
	RegistrationInfo    AccessAIControlMcpServerUpdateResponseAuthConfigSummaryRegistrationInfo `json:"registration_info"`
	JSON                accessAIControlMcpServerUpdateResponseAuthConfigSummaryJSON             `json:"-"`
}

// accessAIControlMcpServerUpdateResponseAuthConfigSummaryJSON contains the JSON
// metadata for the struct
// [AccessAIControlMcpServerUpdateResponseAuthConfigSummary]
type accessAIControlMcpServerUpdateResponseAuthConfigSummaryJSON struct {
	AuthMode            apijson.Field
	ClientSecretVersion apijson.Field
	Config              apijson.Field
	HasClientSecret     apijson.Field
	RegistrationInfo    apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessAIControlMcpServerUpdateResponseAuthConfigSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerUpdateResponseAuthConfigSummaryJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerUpdateResponseAuthConfigSummaryAuthMode string

const (
	AccessAIControlMcpServerUpdateResponseAuthConfigSummaryAuthModeDcr    AccessAIControlMcpServerUpdateResponseAuthConfigSummaryAuthMode = "dcr"
	AccessAIControlMcpServerUpdateResponseAuthConfigSummaryAuthModeManual AccessAIControlMcpServerUpdateResponseAuthConfigSummaryAuthMode = "manual"
)

func (r AccessAIControlMcpServerUpdateResponseAuthConfigSummaryAuthMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerUpdateResponseAuthConfigSummaryAuthModeDcr, AccessAIControlMcpServerUpdateResponseAuthConfigSummaryAuthModeManual:
		return true
	}
	return false
}

type AccessAIControlMcpServerUpdateResponseAuthConfigSummaryConfig struct {
	AuthorizationEndpoint string                                                            `json:"authorization_endpoint"`
	Issuer                string                                                            `json:"issuer"`
	Resource              string                                                            `json:"resource"`
	RevocationEndpoint    string                                                            `json:"revocation_endpoint"`
	TokenEndpoint         string                                                            `json:"token_endpoint"`
	JSON                  accessAIControlMcpServerUpdateResponseAuthConfigSummaryConfigJSON `json:"-"`
}

// accessAIControlMcpServerUpdateResponseAuthConfigSummaryConfigJSON contains the
// JSON metadata for the struct
// [AccessAIControlMcpServerUpdateResponseAuthConfigSummaryConfig]
type accessAIControlMcpServerUpdateResponseAuthConfigSummaryConfigJSON struct {
	AuthorizationEndpoint apijson.Field
	Issuer                apijson.Field
	Resource              apijson.Field
	RevocationEndpoint    apijson.Field
	TokenEndpoint         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessAIControlMcpServerUpdateResponseAuthConfigSummaryConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerUpdateResponseAuthConfigSummaryConfigJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerUpdateResponseAuthConfigSummaryRegistrationInfo struct {
	ClientID                string                                                                      `json:"client_id"`
	RedirectURIs            []string                                                                    `json:"redirect_uris"`
	Scope                   string                                                                      `json:"scope"`
	TokenEndpointAuthMethod string                                                                      `json:"token_endpoint_auth_method"`
	JSON                    accessAIControlMcpServerUpdateResponseAuthConfigSummaryRegistrationInfoJSON `json:"-"`
}

// accessAIControlMcpServerUpdateResponseAuthConfigSummaryRegistrationInfoJSON
// contains the JSON metadata for the struct
// [AccessAIControlMcpServerUpdateResponseAuthConfigSummaryRegistrationInfo]
type accessAIControlMcpServerUpdateResponseAuthConfigSummaryRegistrationInfoJSON struct {
	ClientID                apijson.Field
	RedirectURIs            apijson.Field
	Scope                   apijson.Field
	TokenEndpointAuthMethod apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessAIControlMcpServerUpdateResponseAuthConfigSummaryRegistrationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerUpdateResponseAuthConfigSummaryRegistrationInfoJSON) RawJSON() string {
	return r.raw
}

// Whether administrative authentication is required before capabilities can be
// synced. Manual OAuth is user-managed and has no administrative authentication
// flow.
type AccessAIControlMcpServerUpdateResponseAuthenticationStatus string

const (
	AccessAIControlMcpServerUpdateResponseAuthenticationStatusNotRequired AccessAIControlMcpServerUpdateResponseAuthenticationStatus = "not_required"
	AccessAIControlMcpServerUpdateResponseAuthenticationStatusRequired    AccessAIControlMcpServerUpdateResponseAuthenticationStatus = "required"
	AccessAIControlMcpServerUpdateResponseAuthenticationStatusConnected   AccessAIControlMcpServerUpdateResponseAuthenticationStatus = "connected"
	AccessAIControlMcpServerUpdateResponseAuthenticationStatusStale       AccessAIControlMcpServerUpdateResponseAuthenticationStatus = "stale"
	AccessAIControlMcpServerUpdateResponseAuthenticationStatusManual      AccessAIControlMcpServerUpdateResponseAuthenticationStatus = "manual"
)

func (r AccessAIControlMcpServerUpdateResponseAuthenticationStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerUpdateResponseAuthenticationStatusNotRequired, AccessAIControlMcpServerUpdateResponseAuthenticationStatusRequired, AccessAIControlMcpServerUpdateResponseAuthenticationStatusConnected, AccessAIControlMcpServerUpdateResponseAuthenticationStatusStale, AccessAIControlMcpServerUpdateResponseAuthenticationStatusManual:
		return true
	}
	return false
}

type AccessAIControlMcpServerUpdateResponseErrorDetails struct {
	// Underlying error message
	Cause string `json:"cause"`
	// True = MCP server returned an error. False = couldn't reach the server
	IsUpstream bool `json:"is_upstream"`
	// MCP protocol error code
	McpCode float64 `json:"mcp_code"`
	// Whether the error is transient and worth retrying
	Retryable bool `json:"retryable"`
	// HTTP status code from the server
	StatusCode float64                                                `json:"status_code"`
	JSON       accessAIControlMcpServerUpdateResponseErrorDetailsJSON `json:"-"`
}

// accessAIControlMcpServerUpdateResponseErrorDetailsJSON contains the JSON
// metadata for the struct [AccessAIControlMcpServerUpdateResponseErrorDetails]
type accessAIControlMcpServerUpdateResponseErrorDetailsJSON struct {
	Cause       apijson.Field
	IsUpstream  apijson.Field
	McpCode     apijson.Field
	Retryable   apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerUpdateResponseErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerUpdateResponseErrorDetailsJSON) RawJSON() string {
	return r.raw
}

// Current sync state of the server
type AccessAIControlMcpServerUpdateResponseStatus string

const (
	AccessAIControlMcpServerUpdateResponseStatusWaiting AccessAIControlMcpServerUpdateResponseStatus = "waiting"
	AccessAIControlMcpServerUpdateResponseStatusReady   AccessAIControlMcpServerUpdateResponseStatus = "ready"
	AccessAIControlMcpServerUpdateResponseStatusStale   AccessAIControlMcpServerUpdateResponseStatus = "stale"
	AccessAIControlMcpServerUpdateResponseStatusError   AccessAIControlMcpServerUpdateResponseStatus = "error"
)

func (r AccessAIControlMcpServerUpdateResponseStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerUpdateResponseStatusWaiting, AccessAIControlMcpServerUpdateResponseStatusReady, AccessAIControlMcpServerUpdateResponseStatusStale, AccessAIControlMcpServerUpdateResponseStatusError:
		return true
	}
	return false
}

type AccessAIControlMcpServerUpdateResponseUpdatedPrompt struct {
	// Name of the tool or prompt capability to override.
	Name string `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias string `json:"alias"`
	// Custom description exposed for the capability.
	Description string `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled bool                                                    `json:"enabled"`
	JSON    accessAIControlMcpServerUpdateResponseUpdatedPromptJSON `json:"-"`
}

// accessAIControlMcpServerUpdateResponseUpdatedPromptJSON contains the JSON
// metadata for the struct [AccessAIControlMcpServerUpdateResponseUpdatedPrompt]
type accessAIControlMcpServerUpdateResponseUpdatedPromptJSON struct {
	Name        apijson.Field
	Alias       apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerUpdateResponseUpdatedPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerUpdateResponseUpdatedPromptJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerUpdateResponseUpdatedTool struct {
	// Name of the tool or prompt capability to override.
	Name string `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias string `json:"alias"`
	// Custom description exposed for the capability.
	Description string `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled bool                                                  `json:"enabled"`
	JSON    accessAIControlMcpServerUpdateResponseUpdatedToolJSON `json:"-"`
}

// accessAIControlMcpServerUpdateResponseUpdatedToolJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerUpdateResponseUpdatedTool]
type accessAIControlMcpServerUpdateResponseUpdatedToolJSON struct {
	Name        apijson.Field
	Alias       apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerUpdateResponseUpdatedTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerUpdateResponseUpdatedToolJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerListResponse struct {
	// Unique identifier for the MCP server.
	ID string `json:"id" api:"required"`
	// Authentication method used to connect to the upstream MCP server.
	AuthType AccessAIControlMcpServerListResponseAuthType `json:"auth_type" api:"required"`
	// URL of the upstream MCP endpoint.
	Hostname string `json:"hostname" api:"required" format:"uri"`
	// Display name for the MCP server.
	Name    string                   `json:"name" api:"required"`
	Prompts []map[string]interface{} `json:"prompts" api:"required"`
	Tools   []map[string]interface{} `json:"tools" api:"required"`
	// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
	// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
	// endpoints + client_id for manual servers. Never includes the secret value.
	AuthConfigSummary AccessAIControlMcpServerListResponseAuthConfigSummary `json:"auth_config_summary"`
	// Whether administrative authentication is required before capabilities can be
	// synced. Manual OAuth is user-managed and has no administrative authentication
	// flow.
	AuthenticationStatus AccessAIControlMcpServerListResponseAuthenticationStatus `json:"authentication_status"`
	CreatedAt            time.Time                                                `json:"created_at" format:"date-time"`
	CreatedBy            string                                                   `json:"created_by"`
	// Optional description of the MCP server.
	Description  string                                           `json:"description" api:"nullable"`
	Error        string                                           `json:"error"`
	ErrorDetails AccessAIControlMcpServerListResponseErrorDetails `json:"error_details"`
	// When true, the gateway worker uses the shared Cloudflare-owned OAuth callback
	// endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the
	// customer portal hostname. Defaults to false (off); opt in per server by setting
	// true.
	IsSharedOAuthCallbackEnabled bool      `json:"is_shared_oauth_callback_enabled"`
	LastSuccessfulSync           time.Time `json:"last_successful_sync" format:"date-time"`
	LastSynced                   time.Time `json:"last_synced" format:"date-time"`
	ModifiedAt                   time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy                   string    `json:"modified_by"`
	// Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway.
	SecureWebGateway bool `json:"secure_web_gateway"`
	// Current sync state of the server
	Status AccessAIControlMcpServerListResponseStatus `json:"status"`
	// Server-wide prompt capability overrides.
	UpdatedPrompts []AccessAIControlMcpServerListResponseUpdatedPrompt `json:"updated_prompts"`
	// Server-wide tool capability overrides.
	UpdatedTools []AccessAIControlMcpServerListResponseUpdatedTool `json:"updated_tools"`
	JSON         accessAIControlMcpServerListResponseJSON          `json:"-"`
}

// accessAIControlMcpServerListResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpServerListResponse]
type accessAIControlMcpServerListResponseJSON struct {
	ID                           apijson.Field
	AuthType                     apijson.Field
	Hostname                     apijson.Field
	Name                         apijson.Field
	Prompts                      apijson.Field
	Tools                        apijson.Field
	AuthConfigSummary            apijson.Field
	AuthenticationStatus         apijson.Field
	CreatedAt                    apijson.Field
	CreatedBy                    apijson.Field
	Description                  apijson.Field
	Error                        apijson.Field
	ErrorDetails                 apijson.Field
	IsSharedOAuthCallbackEnabled apijson.Field
	LastSuccessfulSync           apijson.Field
	LastSynced                   apijson.Field
	ModifiedAt                   apijson.Field
	ModifiedBy                   apijson.Field
	SecureWebGateway             apijson.Field
	Status                       apijson.Field
	UpdatedPrompts               apijson.Field
	UpdatedTools                 apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccessAIControlMcpServerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerListResponseJSON) RawJSON() string {
	return r.raw
}

// Authentication method used to connect to the upstream MCP server.
type AccessAIControlMcpServerListResponseAuthType string

const (
	AccessAIControlMcpServerListResponseAuthTypeOAuth           AccessAIControlMcpServerListResponseAuthType = "oauth"
	AccessAIControlMcpServerListResponseAuthTypeBearer          AccessAIControlMcpServerListResponseAuthType = "bearer"
	AccessAIControlMcpServerListResponseAuthTypeUnauthenticated AccessAIControlMcpServerListResponseAuthType = "unauthenticated"
)

func (r AccessAIControlMcpServerListResponseAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerListResponseAuthTypeOAuth, AccessAIControlMcpServerListResponseAuthTypeBearer, AccessAIControlMcpServerListResponseAuthTypeUnauthenticated:
		return true
	}
	return false
}

// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
// endpoints + client_id for manual servers. Never includes the secret value.
type AccessAIControlMcpServerListResponseAuthConfigSummary struct {
	AuthMode            AccessAIControlMcpServerListResponseAuthConfigSummaryAuthMode         `json:"auth_mode"`
	ClientSecretVersion float64                                                               `json:"client_secret_version"`
	Config              AccessAIControlMcpServerListResponseAuthConfigSummaryConfig           `json:"config"`
	HasClientSecret     bool                                                                  `json:"has_client_secret"`
	RegistrationInfo    AccessAIControlMcpServerListResponseAuthConfigSummaryRegistrationInfo `json:"registration_info"`
	JSON                accessAIControlMcpServerListResponseAuthConfigSummaryJSON             `json:"-"`
}

// accessAIControlMcpServerListResponseAuthConfigSummaryJSON contains the JSON
// metadata for the struct [AccessAIControlMcpServerListResponseAuthConfigSummary]
type accessAIControlMcpServerListResponseAuthConfigSummaryJSON struct {
	AuthMode            apijson.Field
	ClientSecretVersion apijson.Field
	Config              apijson.Field
	HasClientSecret     apijson.Field
	RegistrationInfo    apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessAIControlMcpServerListResponseAuthConfigSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerListResponseAuthConfigSummaryJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerListResponseAuthConfigSummaryAuthMode string

const (
	AccessAIControlMcpServerListResponseAuthConfigSummaryAuthModeDcr    AccessAIControlMcpServerListResponseAuthConfigSummaryAuthMode = "dcr"
	AccessAIControlMcpServerListResponseAuthConfigSummaryAuthModeManual AccessAIControlMcpServerListResponseAuthConfigSummaryAuthMode = "manual"
)

func (r AccessAIControlMcpServerListResponseAuthConfigSummaryAuthMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerListResponseAuthConfigSummaryAuthModeDcr, AccessAIControlMcpServerListResponseAuthConfigSummaryAuthModeManual:
		return true
	}
	return false
}

type AccessAIControlMcpServerListResponseAuthConfigSummaryConfig struct {
	AuthorizationEndpoint string                                                          `json:"authorization_endpoint"`
	Issuer                string                                                          `json:"issuer"`
	Resource              string                                                          `json:"resource"`
	RevocationEndpoint    string                                                          `json:"revocation_endpoint"`
	TokenEndpoint         string                                                          `json:"token_endpoint"`
	JSON                  accessAIControlMcpServerListResponseAuthConfigSummaryConfigJSON `json:"-"`
}

// accessAIControlMcpServerListResponseAuthConfigSummaryConfigJSON contains the
// JSON metadata for the struct
// [AccessAIControlMcpServerListResponseAuthConfigSummaryConfig]
type accessAIControlMcpServerListResponseAuthConfigSummaryConfigJSON struct {
	AuthorizationEndpoint apijson.Field
	Issuer                apijson.Field
	Resource              apijson.Field
	RevocationEndpoint    apijson.Field
	TokenEndpoint         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessAIControlMcpServerListResponseAuthConfigSummaryConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerListResponseAuthConfigSummaryConfigJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerListResponseAuthConfigSummaryRegistrationInfo struct {
	ClientID                string                                                                    `json:"client_id"`
	RedirectURIs            []string                                                                  `json:"redirect_uris"`
	Scope                   string                                                                    `json:"scope"`
	TokenEndpointAuthMethod string                                                                    `json:"token_endpoint_auth_method"`
	JSON                    accessAIControlMcpServerListResponseAuthConfigSummaryRegistrationInfoJSON `json:"-"`
}

// accessAIControlMcpServerListResponseAuthConfigSummaryRegistrationInfoJSON
// contains the JSON metadata for the struct
// [AccessAIControlMcpServerListResponseAuthConfigSummaryRegistrationInfo]
type accessAIControlMcpServerListResponseAuthConfigSummaryRegistrationInfoJSON struct {
	ClientID                apijson.Field
	RedirectURIs            apijson.Field
	Scope                   apijson.Field
	TokenEndpointAuthMethod apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessAIControlMcpServerListResponseAuthConfigSummaryRegistrationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerListResponseAuthConfigSummaryRegistrationInfoJSON) RawJSON() string {
	return r.raw
}

// Whether administrative authentication is required before capabilities can be
// synced. Manual OAuth is user-managed and has no administrative authentication
// flow.
type AccessAIControlMcpServerListResponseAuthenticationStatus string

const (
	AccessAIControlMcpServerListResponseAuthenticationStatusNotRequired AccessAIControlMcpServerListResponseAuthenticationStatus = "not_required"
	AccessAIControlMcpServerListResponseAuthenticationStatusRequired    AccessAIControlMcpServerListResponseAuthenticationStatus = "required"
	AccessAIControlMcpServerListResponseAuthenticationStatusConnected   AccessAIControlMcpServerListResponseAuthenticationStatus = "connected"
	AccessAIControlMcpServerListResponseAuthenticationStatusStale       AccessAIControlMcpServerListResponseAuthenticationStatus = "stale"
	AccessAIControlMcpServerListResponseAuthenticationStatusManual      AccessAIControlMcpServerListResponseAuthenticationStatus = "manual"
)

func (r AccessAIControlMcpServerListResponseAuthenticationStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerListResponseAuthenticationStatusNotRequired, AccessAIControlMcpServerListResponseAuthenticationStatusRequired, AccessAIControlMcpServerListResponseAuthenticationStatusConnected, AccessAIControlMcpServerListResponseAuthenticationStatusStale, AccessAIControlMcpServerListResponseAuthenticationStatusManual:
		return true
	}
	return false
}

type AccessAIControlMcpServerListResponseErrorDetails struct {
	// Underlying error message
	Cause string `json:"cause"`
	// True = MCP server returned an error. False = couldn't reach the server
	IsUpstream bool `json:"is_upstream"`
	// MCP protocol error code
	McpCode float64 `json:"mcp_code"`
	// Whether the error is transient and worth retrying
	Retryable bool `json:"retryable"`
	// HTTP status code from the server
	StatusCode float64                                              `json:"status_code"`
	JSON       accessAIControlMcpServerListResponseErrorDetailsJSON `json:"-"`
}

// accessAIControlMcpServerListResponseErrorDetailsJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerListResponseErrorDetails]
type accessAIControlMcpServerListResponseErrorDetailsJSON struct {
	Cause       apijson.Field
	IsUpstream  apijson.Field
	McpCode     apijson.Field
	Retryable   apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerListResponseErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerListResponseErrorDetailsJSON) RawJSON() string {
	return r.raw
}

// Current sync state of the server
type AccessAIControlMcpServerListResponseStatus string

const (
	AccessAIControlMcpServerListResponseStatusWaiting AccessAIControlMcpServerListResponseStatus = "waiting"
	AccessAIControlMcpServerListResponseStatusReady   AccessAIControlMcpServerListResponseStatus = "ready"
	AccessAIControlMcpServerListResponseStatusStale   AccessAIControlMcpServerListResponseStatus = "stale"
	AccessAIControlMcpServerListResponseStatusError   AccessAIControlMcpServerListResponseStatus = "error"
)

func (r AccessAIControlMcpServerListResponseStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerListResponseStatusWaiting, AccessAIControlMcpServerListResponseStatusReady, AccessAIControlMcpServerListResponseStatusStale, AccessAIControlMcpServerListResponseStatusError:
		return true
	}
	return false
}

type AccessAIControlMcpServerListResponseUpdatedPrompt struct {
	// Name of the tool or prompt capability to override.
	Name string `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias string `json:"alias"`
	// Custom description exposed for the capability.
	Description string `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled bool                                                  `json:"enabled"`
	JSON    accessAIControlMcpServerListResponseUpdatedPromptJSON `json:"-"`
}

// accessAIControlMcpServerListResponseUpdatedPromptJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerListResponseUpdatedPrompt]
type accessAIControlMcpServerListResponseUpdatedPromptJSON struct {
	Name        apijson.Field
	Alias       apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerListResponseUpdatedPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerListResponseUpdatedPromptJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerListResponseUpdatedTool struct {
	// Name of the tool or prompt capability to override.
	Name string `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias string `json:"alias"`
	// Custom description exposed for the capability.
	Description string `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled bool                                                `json:"enabled"`
	JSON    accessAIControlMcpServerListResponseUpdatedToolJSON `json:"-"`
}

// accessAIControlMcpServerListResponseUpdatedToolJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerListResponseUpdatedTool]
type accessAIControlMcpServerListResponseUpdatedToolJSON struct {
	Name        apijson.Field
	Alias       apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerListResponseUpdatedTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerListResponseUpdatedToolJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerDeleteResponse struct {
	// Unique identifier for the MCP server.
	ID string `json:"id" api:"required"`
	// Authentication method used to connect to the upstream MCP server.
	AuthType AccessAIControlMcpServerDeleteResponseAuthType `json:"auth_type" api:"required"`
	// URL of the upstream MCP endpoint.
	Hostname string `json:"hostname" api:"required" format:"uri"`
	// Display name for the MCP server.
	Name    string                   `json:"name" api:"required"`
	Prompts []map[string]interface{} `json:"prompts" api:"required"`
	Tools   []map[string]interface{} `json:"tools" api:"required"`
	// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
	// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
	// endpoints + client_id for manual servers. Never includes the secret value.
	AuthConfigSummary AccessAIControlMcpServerDeleteResponseAuthConfigSummary `json:"auth_config_summary"`
	// Whether administrative authentication is required before capabilities can be
	// synced. Manual OAuth is user-managed and has no administrative authentication
	// flow.
	AuthenticationStatus AccessAIControlMcpServerDeleteResponseAuthenticationStatus `json:"authentication_status"`
	CreatedAt            time.Time                                                  `json:"created_at" format:"date-time"`
	CreatedBy            string                                                     `json:"created_by"`
	// Optional description of the MCP server.
	Description  string                                             `json:"description" api:"nullable"`
	Error        string                                             `json:"error"`
	ErrorDetails AccessAIControlMcpServerDeleteResponseErrorDetails `json:"error_details"`
	// When true, the gateway worker uses the shared Cloudflare-owned OAuth callback
	// endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the
	// customer portal hostname. Defaults to false (off); opt in per server by setting
	// true.
	IsSharedOAuthCallbackEnabled bool      `json:"is_shared_oauth_callback_enabled"`
	LastSuccessfulSync           time.Time `json:"last_successful_sync" format:"date-time"`
	LastSynced                   time.Time `json:"last_synced" format:"date-time"`
	ModifiedAt                   time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy                   string    `json:"modified_by"`
	// Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway.
	SecureWebGateway bool `json:"secure_web_gateway"`
	// Current sync state of the server
	Status AccessAIControlMcpServerDeleteResponseStatus `json:"status"`
	// Server-wide prompt capability overrides.
	UpdatedPrompts []AccessAIControlMcpServerDeleteResponseUpdatedPrompt `json:"updated_prompts"`
	// Server-wide tool capability overrides.
	UpdatedTools []AccessAIControlMcpServerDeleteResponseUpdatedTool `json:"updated_tools"`
	JSON         accessAIControlMcpServerDeleteResponseJSON          `json:"-"`
}

// accessAIControlMcpServerDeleteResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpServerDeleteResponse]
type accessAIControlMcpServerDeleteResponseJSON struct {
	ID                           apijson.Field
	AuthType                     apijson.Field
	Hostname                     apijson.Field
	Name                         apijson.Field
	Prompts                      apijson.Field
	Tools                        apijson.Field
	AuthConfigSummary            apijson.Field
	AuthenticationStatus         apijson.Field
	CreatedAt                    apijson.Field
	CreatedBy                    apijson.Field
	Description                  apijson.Field
	Error                        apijson.Field
	ErrorDetails                 apijson.Field
	IsSharedOAuthCallbackEnabled apijson.Field
	LastSuccessfulSync           apijson.Field
	LastSynced                   apijson.Field
	ModifiedAt                   apijson.Field
	ModifiedBy                   apijson.Field
	SecureWebGateway             apijson.Field
	Status                       apijson.Field
	UpdatedPrompts               apijson.Field
	UpdatedTools                 apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccessAIControlMcpServerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Authentication method used to connect to the upstream MCP server.
type AccessAIControlMcpServerDeleteResponseAuthType string

const (
	AccessAIControlMcpServerDeleteResponseAuthTypeOAuth           AccessAIControlMcpServerDeleteResponseAuthType = "oauth"
	AccessAIControlMcpServerDeleteResponseAuthTypeBearer          AccessAIControlMcpServerDeleteResponseAuthType = "bearer"
	AccessAIControlMcpServerDeleteResponseAuthTypeUnauthenticated AccessAIControlMcpServerDeleteResponseAuthType = "unauthenticated"
)

func (r AccessAIControlMcpServerDeleteResponseAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerDeleteResponseAuthTypeOAuth, AccessAIControlMcpServerDeleteResponseAuthTypeBearer, AccessAIControlMcpServerDeleteResponseAuthTypeUnauthenticated:
		return true
	}
	return false
}

// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
// endpoints + client_id for manual servers. Never includes the secret value.
type AccessAIControlMcpServerDeleteResponseAuthConfigSummary struct {
	AuthMode            AccessAIControlMcpServerDeleteResponseAuthConfigSummaryAuthMode         `json:"auth_mode"`
	ClientSecretVersion float64                                                                 `json:"client_secret_version"`
	Config              AccessAIControlMcpServerDeleteResponseAuthConfigSummaryConfig           `json:"config"`
	HasClientSecret     bool                                                                    `json:"has_client_secret"`
	RegistrationInfo    AccessAIControlMcpServerDeleteResponseAuthConfigSummaryRegistrationInfo `json:"registration_info"`
	JSON                accessAIControlMcpServerDeleteResponseAuthConfigSummaryJSON             `json:"-"`
}

// accessAIControlMcpServerDeleteResponseAuthConfigSummaryJSON contains the JSON
// metadata for the struct
// [AccessAIControlMcpServerDeleteResponseAuthConfigSummary]
type accessAIControlMcpServerDeleteResponseAuthConfigSummaryJSON struct {
	AuthMode            apijson.Field
	ClientSecretVersion apijson.Field
	Config              apijson.Field
	HasClientSecret     apijson.Field
	RegistrationInfo    apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessAIControlMcpServerDeleteResponseAuthConfigSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerDeleteResponseAuthConfigSummaryJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerDeleteResponseAuthConfigSummaryAuthMode string

const (
	AccessAIControlMcpServerDeleteResponseAuthConfigSummaryAuthModeDcr    AccessAIControlMcpServerDeleteResponseAuthConfigSummaryAuthMode = "dcr"
	AccessAIControlMcpServerDeleteResponseAuthConfigSummaryAuthModeManual AccessAIControlMcpServerDeleteResponseAuthConfigSummaryAuthMode = "manual"
)

func (r AccessAIControlMcpServerDeleteResponseAuthConfigSummaryAuthMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerDeleteResponseAuthConfigSummaryAuthModeDcr, AccessAIControlMcpServerDeleteResponseAuthConfigSummaryAuthModeManual:
		return true
	}
	return false
}

type AccessAIControlMcpServerDeleteResponseAuthConfigSummaryConfig struct {
	AuthorizationEndpoint string                                                            `json:"authorization_endpoint"`
	Issuer                string                                                            `json:"issuer"`
	Resource              string                                                            `json:"resource"`
	RevocationEndpoint    string                                                            `json:"revocation_endpoint"`
	TokenEndpoint         string                                                            `json:"token_endpoint"`
	JSON                  accessAIControlMcpServerDeleteResponseAuthConfigSummaryConfigJSON `json:"-"`
}

// accessAIControlMcpServerDeleteResponseAuthConfigSummaryConfigJSON contains the
// JSON metadata for the struct
// [AccessAIControlMcpServerDeleteResponseAuthConfigSummaryConfig]
type accessAIControlMcpServerDeleteResponseAuthConfigSummaryConfigJSON struct {
	AuthorizationEndpoint apijson.Field
	Issuer                apijson.Field
	Resource              apijson.Field
	RevocationEndpoint    apijson.Field
	TokenEndpoint         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessAIControlMcpServerDeleteResponseAuthConfigSummaryConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerDeleteResponseAuthConfigSummaryConfigJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerDeleteResponseAuthConfigSummaryRegistrationInfo struct {
	ClientID                string                                                                      `json:"client_id"`
	RedirectURIs            []string                                                                    `json:"redirect_uris"`
	Scope                   string                                                                      `json:"scope"`
	TokenEndpointAuthMethod string                                                                      `json:"token_endpoint_auth_method"`
	JSON                    accessAIControlMcpServerDeleteResponseAuthConfigSummaryRegistrationInfoJSON `json:"-"`
}

// accessAIControlMcpServerDeleteResponseAuthConfigSummaryRegistrationInfoJSON
// contains the JSON metadata for the struct
// [AccessAIControlMcpServerDeleteResponseAuthConfigSummaryRegistrationInfo]
type accessAIControlMcpServerDeleteResponseAuthConfigSummaryRegistrationInfoJSON struct {
	ClientID                apijson.Field
	RedirectURIs            apijson.Field
	Scope                   apijson.Field
	TokenEndpointAuthMethod apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessAIControlMcpServerDeleteResponseAuthConfigSummaryRegistrationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerDeleteResponseAuthConfigSummaryRegistrationInfoJSON) RawJSON() string {
	return r.raw
}

// Whether administrative authentication is required before capabilities can be
// synced. Manual OAuth is user-managed and has no administrative authentication
// flow.
type AccessAIControlMcpServerDeleteResponseAuthenticationStatus string

const (
	AccessAIControlMcpServerDeleteResponseAuthenticationStatusNotRequired AccessAIControlMcpServerDeleteResponseAuthenticationStatus = "not_required"
	AccessAIControlMcpServerDeleteResponseAuthenticationStatusRequired    AccessAIControlMcpServerDeleteResponseAuthenticationStatus = "required"
	AccessAIControlMcpServerDeleteResponseAuthenticationStatusConnected   AccessAIControlMcpServerDeleteResponseAuthenticationStatus = "connected"
	AccessAIControlMcpServerDeleteResponseAuthenticationStatusStale       AccessAIControlMcpServerDeleteResponseAuthenticationStatus = "stale"
	AccessAIControlMcpServerDeleteResponseAuthenticationStatusManual      AccessAIControlMcpServerDeleteResponseAuthenticationStatus = "manual"
)

func (r AccessAIControlMcpServerDeleteResponseAuthenticationStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerDeleteResponseAuthenticationStatusNotRequired, AccessAIControlMcpServerDeleteResponseAuthenticationStatusRequired, AccessAIControlMcpServerDeleteResponseAuthenticationStatusConnected, AccessAIControlMcpServerDeleteResponseAuthenticationStatusStale, AccessAIControlMcpServerDeleteResponseAuthenticationStatusManual:
		return true
	}
	return false
}

type AccessAIControlMcpServerDeleteResponseErrorDetails struct {
	// Underlying error message
	Cause string `json:"cause"`
	// True = MCP server returned an error. False = couldn't reach the server
	IsUpstream bool `json:"is_upstream"`
	// MCP protocol error code
	McpCode float64 `json:"mcp_code"`
	// Whether the error is transient and worth retrying
	Retryable bool `json:"retryable"`
	// HTTP status code from the server
	StatusCode float64                                                `json:"status_code"`
	JSON       accessAIControlMcpServerDeleteResponseErrorDetailsJSON `json:"-"`
}

// accessAIControlMcpServerDeleteResponseErrorDetailsJSON contains the JSON
// metadata for the struct [AccessAIControlMcpServerDeleteResponseErrorDetails]
type accessAIControlMcpServerDeleteResponseErrorDetailsJSON struct {
	Cause       apijson.Field
	IsUpstream  apijson.Field
	McpCode     apijson.Field
	Retryable   apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerDeleteResponseErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerDeleteResponseErrorDetailsJSON) RawJSON() string {
	return r.raw
}

// Current sync state of the server
type AccessAIControlMcpServerDeleteResponseStatus string

const (
	AccessAIControlMcpServerDeleteResponseStatusWaiting AccessAIControlMcpServerDeleteResponseStatus = "waiting"
	AccessAIControlMcpServerDeleteResponseStatusReady   AccessAIControlMcpServerDeleteResponseStatus = "ready"
	AccessAIControlMcpServerDeleteResponseStatusStale   AccessAIControlMcpServerDeleteResponseStatus = "stale"
	AccessAIControlMcpServerDeleteResponseStatusError   AccessAIControlMcpServerDeleteResponseStatus = "error"
)

func (r AccessAIControlMcpServerDeleteResponseStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerDeleteResponseStatusWaiting, AccessAIControlMcpServerDeleteResponseStatusReady, AccessAIControlMcpServerDeleteResponseStatusStale, AccessAIControlMcpServerDeleteResponseStatusError:
		return true
	}
	return false
}

type AccessAIControlMcpServerDeleteResponseUpdatedPrompt struct {
	// Name of the tool or prompt capability to override.
	Name string `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias string `json:"alias"`
	// Custom description exposed for the capability.
	Description string `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled bool                                                    `json:"enabled"`
	JSON    accessAIControlMcpServerDeleteResponseUpdatedPromptJSON `json:"-"`
}

// accessAIControlMcpServerDeleteResponseUpdatedPromptJSON contains the JSON
// metadata for the struct [AccessAIControlMcpServerDeleteResponseUpdatedPrompt]
type accessAIControlMcpServerDeleteResponseUpdatedPromptJSON struct {
	Name        apijson.Field
	Alias       apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerDeleteResponseUpdatedPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerDeleteResponseUpdatedPromptJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerDeleteResponseUpdatedTool struct {
	// Name of the tool or prompt capability to override.
	Name string `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias string `json:"alias"`
	// Custom description exposed for the capability.
	Description string `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled bool                                                  `json:"enabled"`
	JSON    accessAIControlMcpServerDeleteResponseUpdatedToolJSON `json:"-"`
}

// accessAIControlMcpServerDeleteResponseUpdatedToolJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerDeleteResponseUpdatedTool]
type accessAIControlMcpServerDeleteResponseUpdatedToolJSON struct {
	Name        apijson.Field
	Alias       apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerDeleteResponseUpdatedTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerDeleteResponseUpdatedToolJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerReadResponse struct {
	// Unique identifier for the MCP server.
	ID string `json:"id" api:"required"`
	// Authentication method used to connect to the upstream MCP server.
	AuthType AccessAIControlMcpServerReadResponseAuthType `json:"auth_type" api:"required"`
	// URL of the upstream MCP endpoint.
	Hostname string `json:"hostname" api:"required" format:"uri"`
	// Display name for the MCP server.
	Name    string                   `json:"name" api:"required"`
	Prompts []map[string]interface{} `json:"prompts" api:"required"`
	Tools   []map[string]interface{} `json:"tools" api:"required"`
	// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
	// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
	// endpoints + client_id for manual servers. Never includes the secret value.
	AuthConfigSummary AccessAIControlMcpServerReadResponseAuthConfigSummary `json:"auth_config_summary"`
	// Whether administrative authentication is required before capabilities can be
	// synced. Manual OAuth is user-managed and has no administrative authentication
	// flow.
	AuthenticationStatus AccessAIControlMcpServerReadResponseAuthenticationStatus `json:"authentication_status"`
	CreatedAt            time.Time                                                `json:"created_at" format:"date-time"`
	CreatedBy            string                                                   `json:"created_by"`
	// Optional description of the MCP server.
	Description  string                                           `json:"description" api:"nullable"`
	Error        string                                           `json:"error"`
	ErrorDetails AccessAIControlMcpServerReadResponseErrorDetails `json:"error_details"`
	// When true, the gateway worker uses the shared Cloudflare-owned OAuth callback
	// endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the
	// customer portal hostname. Defaults to false (off); opt in per server by setting
	// true.
	IsSharedOAuthCallbackEnabled bool      `json:"is_shared_oauth_callback_enabled"`
	LastSuccessfulSync           time.Time `json:"last_successful_sync" format:"date-time"`
	LastSynced                   time.Time `json:"last_synced" format:"date-time"`
	ModifiedAt                   time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy                   string    `json:"modified_by"`
	// Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway.
	SecureWebGateway bool `json:"secure_web_gateway"`
	// Current sync state of the server
	Status AccessAIControlMcpServerReadResponseStatus `json:"status"`
	// Server-wide prompt capability overrides.
	UpdatedPrompts []AccessAIControlMcpServerReadResponseUpdatedPrompt `json:"updated_prompts"`
	// Server-wide tool capability overrides.
	UpdatedTools []AccessAIControlMcpServerReadResponseUpdatedTool `json:"updated_tools"`
	JSON         accessAIControlMcpServerReadResponseJSON          `json:"-"`
}

// accessAIControlMcpServerReadResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpServerReadResponse]
type accessAIControlMcpServerReadResponseJSON struct {
	ID                           apijson.Field
	AuthType                     apijson.Field
	Hostname                     apijson.Field
	Name                         apijson.Field
	Prompts                      apijson.Field
	Tools                        apijson.Field
	AuthConfigSummary            apijson.Field
	AuthenticationStatus         apijson.Field
	CreatedAt                    apijson.Field
	CreatedBy                    apijson.Field
	Description                  apijson.Field
	Error                        apijson.Field
	ErrorDetails                 apijson.Field
	IsSharedOAuthCallbackEnabled apijson.Field
	LastSuccessfulSync           apijson.Field
	LastSynced                   apijson.Field
	ModifiedAt                   apijson.Field
	ModifiedBy                   apijson.Field
	SecureWebGateway             apijson.Field
	Status                       apijson.Field
	UpdatedPrompts               apijson.Field
	UpdatedTools                 apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccessAIControlMcpServerReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerReadResponseJSON) RawJSON() string {
	return r.raw
}

// Authentication method used to connect to the upstream MCP server.
type AccessAIControlMcpServerReadResponseAuthType string

const (
	AccessAIControlMcpServerReadResponseAuthTypeOAuth           AccessAIControlMcpServerReadResponseAuthType = "oauth"
	AccessAIControlMcpServerReadResponseAuthTypeBearer          AccessAIControlMcpServerReadResponseAuthType = "bearer"
	AccessAIControlMcpServerReadResponseAuthTypeUnauthenticated AccessAIControlMcpServerReadResponseAuthType = "unauthenticated"
)

func (r AccessAIControlMcpServerReadResponseAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerReadResponseAuthTypeOAuth, AccessAIControlMcpServerReadResponseAuthTypeBearer, AccessAIControlMcpServerReadResponseAuthTypeUnauthenticated:
		return true
	}
	return false
}

// Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode
// (dcr|manual), has_client_secret, client_secret_version, and the OAuth
// endpoints + client_id for manual servers. Never includes the secret value.
type AccessAIControlMcpServerReadResponseAuthConfigSummary struct {
	AuthMode            AccessAIControlMcpServerReadResponseAuthConfigSummaryAuthMode         `json:"auth_mode"`
	ClientSecretVersion float64                                                               `json:"client_secret_version"`
	Config              AccessAIControlMcpServerReadResponseAuthConfigSummaryConfig           `json:"config"`
	HasClientSecret     bool                                                                  `json:"has_client_secret"`
	RegistrationInfo    AccessAIControlMcpServerReadResponseAuthConfigSummaryRegistrationInfo `json:"registration_info"`
	JSON                accessAIControlMcpServerReadResponseAuthConfigSummaryJSON             `json:"-"`
}

// accessAIControlMcpServerReadResponseAuthConfigSummaryJSON contains the JSON
// metadata for the struct [AccessAIControlMcpServerReadResponseAuthConfigSummary]
type accessAIControlMcpServerReadResponseAuthConfigSummaryJSON struct {
	AuthMode            apijson.Field
	ClientSecretVersion apijson.Field
	Config              apijson.Field
	HasClientSecret     apijson.Field
	RegistrationInfo    apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessAIControlMcpServerReadResponseAuthConfigSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerReadResponseAuthConfigSummaryJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerReadResponseAuthConfigSummaryAuthMode string

const (
	AccessAIControlMcpServerReadResponseAuthConfigSummaryAuthModeDcr    AccessAIControlMcpServerReadResponseAuthConfigSummaryAuthMode = "dcr"
	AccessAIControlMcpServerReadResponseAuthConfigSummaryAuthModeManual AccessAIControlMcpServerReadResponseAuthConfigSummaryAuthMode = "manual"
)

func (r AccessAIControlMcpServerReadResponseAuthConfigSummaryAuthMode) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerReadResponseAuthConfigSummaryAuthModeDcr, AccessAIControlMcpServerReadResponseAuthConfigSummaryAuthModeManual:
		return true
	}
	return false
}

type AccessAIControlMcpServerReadResponseAuthConfigSummaryConfig struct {
	AuthorizationEndpoint string                                                          `json:"authorization_endpoint"`
	Issuer                string                                                          `json:"issuer"`
	Resource              string                                                          `json:"resource"`
	RevocationEndpoint    string                                                          `json:"revocation_endpoint"`
	TokenEndpoint         string                                                          `json:"token_endpoint"`
	JSON                  accessAIControlMcpServerReadResponseAuthConfigSummaryConfigJSON `json:"-"`
}

// accessAIControlMcpServerReadResponseAuthConfigSummaryConfigJSON contains the
// JSON metadata for the struct
// [AccessAIControlMcpServerReadResponseAuthConfigSummaryConfig]
type accessAIControlMcpServerReadResponseAuthConfigSummaryConfigJSON struct {
	AuthorizationEndpoint apijson.Field
	Issuer                apijson.Field
	Resource              apijson.Field
	RevocationEndpoint    apijson.Field
	TokenEndpoint         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessAIControlMcpServerReadResponseAuthConfigSummaryConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerReadResponseAuthConfigSummaryConfigJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerReadResponseAuthConfigSummaryRegistrationInfo struct {
	ClientID                string                                                                    `json:"client_id"`
	RedirectURIs            []string                                                                  `json:"redirect_uris"`
	Scope                   string                                                                    `json:"scope"`
	TokenEndpointAuthMethod string                                                                    `json:"token_endpoint_auth_method"`
	JSON                    accessAIControlMcpServerReadResponseAuthConfigSummaryRegistrationInfoJSON `json:"-"`
}

// accessAIControlMcpServerReadResponseAuthConfigSummaryRegistrationInfoJSON
// contains the JSON metadata for the struct
// [AccessAIControlMcpServerReadResponseAuthConfigSummaryRegistrationInfo]
type accessAIControlMcpServerReadResponseAuthConfigSummaryRegistrationInfoJSON struct {
	ClientID                apijson.Field
	RedirectURIs            apijson.Field
	Scope                   apijson.Field
	TokenEndpointAuthMethod apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessAIControlMcpServerReadResponseAuthConfigSummaryRegistrationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerReadResponseAuthConfigSummaryRegistrationInfoJSON) RawJSON() string {
	return r.raw
}

// Whether administrative authentication is required before capabilities can be
// synced. Manual OAuth is user-managed and has no administrative authentication
// flow.
type AccessAIControlMcpServerReadResponseAuthenticationStatus string

const (
	AccessAIControlMcpServerReadResponseAuthenticationStatusNotRequired AccessAIControlMcpServerReadResponseAuthenticationStatus = "not_required"
	AccessAIControlMcpServerReadResponseAuthenticationStatusRequired    AccessAIControlMcpServerReadResponseAuthenticationStatus = "required"
	AccessAIControlMcpServerReadResponseAuthenticationStatusConnected   AccessAIControlMcpServerReadResponseAuthenticationStatus = "connected"
	AccessAIControlMcpServerReadResponseAuthenticationStatusStale       AccessAIControlMcpServerReadResponseAuthenticationStatus = "stale"
	AccessAIControlMcpServerReadResponseAuthenticationStatusManual      AccessAIControlMcpServerReadResponseAuthenticationStatus = "manual"
)

func (r AccessAIControlMcpServerReadResponseAuthenticationStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerReadResponseAuthenticationStatusNotRequired, AccessAIControlMcpServerReadResponseAuthenticationStatusRequired, AccessAIControlMcpServerReadResponseAuthenticationStatusConnected, AccessAIControlMcpServerReadResponseAuthenticationStatusStale, AccessAIControlMcpServerReadResponseAuthenticationStatusManual:
		return true
	}
	return false
}

type AccessAIControlMcpServerReadResponseErrorDetails struct {
	// Underlying error message
	Cause string `json:"cause"`
	// True = MCP server returned an error. False = couldn't reach the server
	IsUpstream bool `json:"is_upstream"`
	// MCP protocol error code
	McpCode float64 `json:"mcp_code"`
	// Whether the error is transient and worth retrying
	Retryable bool `json:"retryable"`
	// HTTP status code from the server
	StatusCode float64                                              `json:"status_code"`
	JSON       accessAIControlMcpServerReadResponseErrorDetailsJSON `json:"-"`
}

// accessAIControlMcpServerReadResponseErrorDetailsJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerReadResponseErrorDetails]
type accessAIControlMcpServerReadResponseErrorDetailsJSON struct {
	Cause       apijson.Field
	IsUpstream  apijson.Field
	McpCode     apijson.Field
	Retryable   apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerReadResponseErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerReadResponseErrorDetailsJSON) RawJSON() string {
	return r.raw
}

// Current sync state of the server
type AccessAIControlMcpServerReadResponseStatus string

const (
	AccessAIControlMcpServerReadResponseStatusWaiting AccessAIControlMcpServerReadResponseStatus = "waiting"
	AccessAIControlMcpServerReadResponseStatusReady   AccessAIControlMcpServerReadResponseStatus = "ready"
	AccessAIControlMcpServerReadResponseStatusStale   AccessAIControlMcpServerReadResponseStatus = "stale"
	AccessAIControlMcpServerReadResponseStatusError   AccessAIControlMcpServerReadResponseStatus = "error"
)

func (r AccessAIControlMcpServerReadResponseStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerReadResponseStatusWaiting, AccessAIControlMcpServerReadResponseStatusReady, AccessAIControlMcpServerReadResponseStatusStale, AccessAIControlMcpServerReadResponseStatusError:
		return true
	}
	return false
}

type AccessAIControlMcpServerReadResponseUpdatedPrompt struct {
	// Name of the tool or prompt capability to override.
	Name string `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias string `json:"alias"`
	// Custom description exposed for the capability.
	Description string `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled bool                                                  `json:"enabled"`
	JSON    accessAIControlMcpServerReadResponseUpdatedPromptJSON `json:"-"`
}

// accessAIControlMcpServerReadResponseUpdatedPromptJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerReadResponseUpdatedPrompt]
type accessAIControlMcpServerReadResponseUpdatedPromptJSON struct {
	Name        apijson.Field
	Alias       apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerReadResponseUpdatedPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerReadResponseUpdatedPromptJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerReadResponseUpdatedTool struct {
	// Name of the tool or prompt capability to override.
	Name string `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias string `json:"alias"`
	// Custom description exposed for the capability.
	Description string `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled bool                                                `json:"enabled"`
	JSON    accessAIControlMcpServerReadResponseUpdatedToolJSON `json:"-"`
}

// accessAIControlMcpServerReadResponseUpdatedToolJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerReadResponseUpdatedTool]
type accessAIControlMcpServerReadResponseUpdatedToolJSON struct {
	Name        apijson.Field
	Alias       apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerReadResponseUpdatedTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerReadResponseUpdatedToolJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerSyncResponse struct {
	Error        string                                           `json:"error"`
	ErrorDetails AccessAIControlMcpServerSyncResponseErrorDetails `json:"error_details"`
	Status       AccessAIControlMcpServerSyncResponseStatus       `json:"status"`
	JSON         accessAIControlMcpServerSyncResponseJSON         `json:"-"`
}

// accessAIControlMcpServerSyncResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpServerSyncResponse]
type accessAIControlMcpServerSyncResponseJSON struct {
	Error        apijson.Field
	ErrorDetails apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAIControlMcpServerSyncResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerSyncResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerSyncResponseErrorDetails struct {
	// Underlying error message
	Cause string `json:"cause"`
	// True = MCP server returned an error. False = couldn't reach the server
	IsUpstream bool `json:"is_upstream"`
	// MCP protocol error code
	McpCode float64 `json:"mcp_code"`
	// Whether the error is transient and worth retrying
	Retryable bool `json:"retryable"`
	// HTTP status code from the server
	StatusCode float64                                              `json:"status_code"`
	JSON       accessAIControlMcpServerSyncResponseErrorDetailsJSON `json:"-"`
}

// accessAIControlMcpServerSyncResponseErrorDetailsJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerSyncResponseErrorDetails]
type accessAIControlMcpServerSyncResponseErrorDetailsJSON struct {
	Cause       apijson.Field
	IsUpstream  apijson.Field
	McpCode     apijson.Field
	Retryable   apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerSyncResponseErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerSyncResponseErrorDetailsJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerSyncResponseStatus string

const (
	AccessAIControlMcpServerSyncResponseStatusWaiting AccessAIControlMcpServerSyncResponseStatus = "waiting"
	AccessAIControlMcpServerSyncResponseStatusReady   AccessAIControlMcpServerSyncResponseStatus = "ready"
	AccessAIControlMcpServerSyncResponseStatusStale   AccessAIControlMcpServerSyncResponseStatus = "stale"
	AccessAIControlMcpServerSyncResponseStatusError   AccessAIControlMcpServerSyncResponseStatus = "error"
)

func (r AccessAIControlMcpServerSyncResponseStatus) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerSyncResponseStatusWaiting, AccessAIControlMcpServerSyncResponseStatusReady, AccessAIControlMcpServerSyncResponseStatusStale, AccessAIControlMcpServerSyncResponseStatusError:
		return true
	}
	return false
}

type AccessAIControlMcpServerNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Unique identifier for the MCP server.
	ID param.Field[string] `json:"id" api:"required"`
	// Authentication method used to connect to the upstream MCP server.
	AuthType param.Field[AccessAIControlMcpServerNewParamsAuthType] `json:"auth_type" api:"required"`
	// URL of the upstream MCP endpoint.
	Hostname param.Field[string] `json:"hostname" api:"required" format:"uri"`
	// Display name for the MCP server.
	Name param.Field[string] `json:"name" api:"required"`
	// Static credential for the upstream MCP server. For auth_type "bearer", either a
	// raw token string (e.g. "sk-abc123"), which is wrapped server-side as
	// `Authorization: Bearer <token>`, or a JSON-encoded object of the form
	// `{"headers":{"Header-Name":"value",...}}` for custom or multiple static headers
	// (e.g. Cloudflare Access service tokens:
	// `{"headers":{"cf-access-client-id":"...","cf-access-client-secret":"..."}}`).
	AuthCredentials param.Field[string] `json:"auth_credentials"`
	// Pre-registered OAuth client_secret. Write-only - accepted on create/update when
	// auth_credentials.auth_mode is 'manual'. Stored AES-GCM-encrypted in
	// server_oauth_secrets; never returned by read endpoints.
	ClientSecret param.Field[string] `json:"client_secret"`
	// Optional description of the MCP server.
	Description param.Field[string] `json:"description"`
	// When true, the gateway worker uses the shared Cloudflare-owned OAuth callback
	// endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the
	// customer portal hostname. Defaults to false (off); opt in per server by setting
	// true.
	IsSharedOAuthCallbackEnabled param.Field[bool] `json:"is_shared_oauth_callback_enabled"`
	// Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway.
	SecureWebGateway param.Field[bool] `json:"secure_web_gateway"`
	// Server-wide prompt capability overrides.
	UpdatedPrompts param.Field[[]AccessAIControlMcpServerNewParamsUpdatedPrompt] `json:"updated_prompts"`
	// Server-wide tool capability overrides.
	UpdatedTools param.Field[[]AccessAIControlMcpServerNewParamsUpdatedTool] `json:"updated_tools"`
}

func (r AccessAIControlMcpServerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Authentication method used to connect to the upstream MCP server.
type AccessAIControlMcpServerNewParamsAuthType string

const (
	AccessAIControlMcpServerNewParamsAuthTypeOAuth           AccessAIControlMcpServerNewParamsAuthType = "oauth"
	AccessAIControlMcpServerNewParamsAuthTypeBearer          AccessAIControlMcpServerNewParamsAuthType = "bearer"
	AccessAIControlMcpServerNewParamsAuthTypeUnauthenticated AccessAIControlMcpServerNewParamsAuthType = "unauthenticated"
)

func (r AccessAIControlMcpServerNewParamsAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerNewParamsAuthTypeOAuth, AccessAIControlMcpServerNewParamsAuthTypeBearer, AccessAIControlMcpServerNewParamsAuthTypeUnauthenticated:
		return true
	}
	return false
}

type AccessAIControlMcpServerNewParamsUpdatedPrompt struct {
	// Name of the tool or prompt capability to override.
	Name param.Field[string] `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias param.Field[string] `json:"alias"`
	// Custom description exposed for the capability.
	Description param.Field[string] `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccessAIControlMcpServerNewParamsUpdatedPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpServerNewParamsUpdatedTool struct {
	// Name of the tool or prompt capability to override.
	Name param.Field[string] `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias param.Field[string] `json:"alias"`
	// Custom description exposed for the capability.
	Description param.Field[string] `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccessAIControlMcpServerNewParamsUpdatedTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpServerNewResponseEnvelope struct {
	Result  AccessAIControlMcpServerNewResponse             `json:"result" api:"required"`
	Success bool                                            `json:"success" api:"required"`
	JSON    accessAIControlMcpServerNewResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpServerNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessAIControlMcpServerNewResponseEnvelope]
type accessAIControlMcpServerNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerUpdateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Static credential for the upstream MCP server. For auth_type "bearer", either a
	// raw token string (e.g. "sk-abc123"), which is wrapped server-side as
	// `Authorization: Bearer <token>`, or a JSON-encoded object of the form
	// `{"headers":{"Header-Name":"value",...}}` for custom or multiple static headers
	// (e.g. Cloudflare Access service tokens:
	// `{"headers":{"cf-access-client-id":"...","cf-access-client-secret":"..."}}`).
	AuthCredentials param.Field[string] `json:"auth_credentials"`
	// Pre-registered OAuth client_secret. Write-only - accepted on create/update when
	// auth_credentials.auth_mode is 'manual'. Stored AES-GCM-encrypted in
	// server_oauth_secrets; never returned by read endpoints.
	ClientSecret param.Field[string] `json:"client_secret"`
	// Optional description of the MCP server.
	Description param.Field[string] `json:"description"`
	// When true, the gateway worker uses the shared Cloudflare-owned OAuth callback
	// endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the
	// customer portal hostname. Defaults to false (off); opt in per server by setting
	// true.
	IsSharedOAuthCallbackEnabled param.Field[bool] `json:"is_shared_oauth_callback_enabled"`
	// Display name for the MCP server.
	Name param.Field[string] `json:"name"`
	// Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway.
	SecureWebGateway param.Field[bool] `json:"secure_web_gateway"`
	// Server-wide prompt capability overrides.
	UpdatedPrompts param.Field[[]AccessAIControlMcpServerUpdateParamsUpdatedPrompt] `json:"updated_prompts"`
	// Server-wide tool capability overrides.
	UpdatedTools param.Field[[]AccessAIControlMcpServerUpdateParamsUpdatedTool] `json:"updated_tools"`
}

func (r AccessAIControlMcpServerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpServerUpdateParamsUpdatedPrompt struct {
	// Name of the tool or prompt capability to override.
	Name param.Field[string] `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias param.Field[string] `json:"alias"`
	// Custom description exposed for the capability.
	Description param.Field[string] `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccessAIControlMcpServerUpdateParamsUpdatedPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpServerUpdateParamsUpdatedTool struct {
	// Name of the tool or prompt capability to override.
	Name param.Field[string] `json:"name" api:"required"`
	// Custom name exposed for the capability.
	Alias param.Field[string] `json:"alias"`
	// Custom description exposed for the capability.
	Description param.Field[string] `json:"description"`
	// Whether the capability is available through the MCP server.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccessAIControlMcpServerUpdateParamsUpdatedTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpServerUpdateResponseEnvelope struct {
	Result  AccessAIControlMcpServerUpdateResponse             `json:"result" api:"required"`
	Success bool                                               `json:"success" api:"required"`
	JSON    accessAIControlMcpServerUpdateResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpServerUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerUpdateResponseEnvelope]
type accessAIControlMcpServerUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
	// Search by id, name
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccessAIControlMcpServerListParams]'s query parameters as
// `url.Values`.
func (r AccessAIControlMcpServerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessAIControlMcpServerDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessAIControlMcpServerDeleteResponseEnvelope struct {
	Result  AccessAIControlMcpServerDeleteResponse             `json:"result" api:"required"`
	Success bool                                               `json:"success" api:"required"`
	JSON    accessAIControlMcpServerDeleteResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpServerDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerDeleteResponseEnvelope]
type accessAIControlMcpServerDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerReadParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessAIControlMcpServerReadResponseEnvelope struct {
	Result  AccessAIControlMcpServerReadResponse             `json:"result" api:"required"`
	Success bool                                             `json:"success" api:"required"`
	JSON    accessAIControlMcpServerReadResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpServerReadResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessAIControlMcpServerReadResponseEnvelope]
type accessAIControlMcpServerReadResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerReadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerReadResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerSyncParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessAIControlMcpServerSyncResponseEnvelope struct {
	Result  AccessAIControlMcpServerSyncResponse             `json:"result" api:"required"`
	Success bool                                             `json:"success" api:"required"`
	JSON    accessAIControlMcpServerSyncResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpServerSyncResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessAIControlMcpServerSyncResponseEnvelope]
type accessAIControlMcpServerSyncResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerSyncResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerSyncResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
