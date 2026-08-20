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

// CasbIntegrationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbIntegrationService] method instead.
type CasbIntegrationService struct {
	Options []option.RequestOption
}

// NewCasbIntegrationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCasbIntegrationService(opts ...option.RequestOption) (r *CasbIntegrationService) {
	r = &CasbIntegrationService{}
	r.Options = opts
	return
}

// Creates a new integration for the specified application. Integration creation
// with OAuth is not supported by API at the moment. For other auth methods, use
// `GET /v2/applications/{application_id}/credential-guide` to see the required
// credential structure and example payloads for each vendor.
func (r *CasbIntegrationService) New(ctx context.Context, params CasbIntegrationNewParams, opts ...option.RequestOption) (res *CasbIntegrationNewResponse, err error) {
	var env CasbIntegrationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/one/integrations", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an integration's name, permissions, DLP profiles, use cases, or
// credentials.
func (r *CasbIntegrationService) Update(ctx context.Context, id string, params CasbIntegrationUpdateParams, opts ...option.RequestOption) (res *CasbIntegrationUpdateResponse, err error) {
	var env CasbIntegrationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/one/integrations/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of integrations for the account.
func (r *CasbIntegrationService) List(ctx context.Context, params CasbIntegrationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CasbIntegrationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/one/integrations", params.AccountID)
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

// Returns a paginated list of integrations for the account.
func (r *CasbIntegrationService) ListAutoPaging(ctx context.Context, params CasbIntegrationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CasbIntegrationListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Delete an integration by soft-deleting it.
func (r *CasbIntegrationService) Delete(ctx context.Context, id string, body CasbIntegrationDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/one/integrations/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns full integration details including use cases and permissions.
func (r *CasbIntegrationService) Get(ctx context.Context, id string, query CasbIntegrationGetParams, opts ...option.RequestOption) (res *CasbIntegrationGetResponse, err error) {
	var env CasbIntegrationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/one/integrations/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Pauses an integration, stopping all crawlers.
func (r *CasbIntegrationService) Pause(ctx context.Context, id string, body CasbIntegrationPauseParams, opts ...option.RequestOption) (res *CasbIntegrationPauseResponse, err error) {
	var env CasbIntegrationPauseResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/one/integrations/%s/pause", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Resumes a paused integration, restarting crawlers.
func (r *CasbIntegrationService) Resume(ctx context.Context, id string, body CasbIntegrationResumeParams, opts ...option.RequestOption) (res *CasbIntegrationResumeResponse, err error) {
	var env CasbIntegrationResumeResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/one/integrations/%s/resume", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// The requested item.
type CasbIntegrationNewResponse struct {
	// Integration ID.
	ID          string            `json:"id" api:"required" format:"uuid"`
	Application map[string]string `json:"application" api:"required"`
	// The integration's authentication method.
	AuthMethod map[string]string `json:"auth_method" api:"required,nullable"`
	// Authorization link for the integration.
	AuthorizationLink CasbIntegrationNewResponseAuthorizationLink `json:"authorization_link" api:"required,nullable"`
	// When the integration was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Credentials expiry time.
	CredentialsExpiry time.Time `json:"credentials_expiry" api:"required" format:"date-time"`
	// DLP Profiles enabled for the integration.
	DLPProfiles []string `json:"dlp_profiles" api:"required" format:"uuid"`
	// Health details with remediation hints.
	HealthDetails []map[string]interface{} `json:"health_details" api:"required"`
	// Whether the user paused the integration.
	IsPaused bool `json:"is_paused" api:"required"`
	// Last time the integration was hydrated.
	LastHydrated time.Time `json:"last_hydrated" api:"required" format:"date-time"`
	// Name of the integration.
	Name string `json:"name" api:"required"`
	// Integration status.
	Status string `json:"status" api:"required"`
	// When the integration was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Use cases enabled for the integration.
	UseCases []map[string]interface{}       `json:"use_cases" api:"required"`
	JSON     casbIntegrationNewResponseJSON `json:"-"`
}

// casbIntegrationNewResponseJSON contains the JSON metadata for the struct
// [CasbIntegrationNewResponse]
type casbIntegrationNewResponseJSON struct {
	ID                apijson.Field
	Application       apijson.Field
	AuthMethod        apijson.Field
	AuthorizationLink apijson.Field
	Created           apijson.Field
	CredentialsExpiry apijson.Field
	DLPProfiles       apijson.Field
	HealthDetails     apijson.Field
	IsPaused          apijson.Field
	LastHydrated      apijson.Field
	Name              apijson.Field
	Status            apijson.Field
	Updated           apijson.Field
	UseCases          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbIntegrationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationNewResponseJSON) RawJSON() string {
	return r.raw
}

// Authorization link for the integration.
type CasbIntegrationNewResponseAuthorizationLink struct {
	Components map[string]interface{}                          `json:"components" api:"required,nullable"`
	Link       string                                          `json:"link" api:"required,nullable"`
	JSON       casbIntegrationNewResponseAuthorizationLinkJSON `json:"-"`
}

// casbIntegrationNewResponseAuthorizationLinkJSON contains the JSON metadata for
// the struct [CasbIntegrationNewResponseAuthorizationLink]
type casbIntegrationNewResponseAuthorizationLinkJSON struct {
	Components  apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbIntegrationNewResponseAuthorizationLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationNewResponseAuthorizationLinkJSON) RawJSON() string {
	return r.raw
}

// The requested item.
type CasbIntegrationUpdateResponse struct {
	// Integration ID.
	ID          string            `json:"id" api:"required" format:"uuid"`
	Application map[string]string `json:"application" api:"required"`
	// The integration's authentication method.
	AuthMethod map[string]string `json:"auth_method" api:"required,nullable"`
	// Authorization link for the integration.
	AuthorizationLink CasbIntegrationUpdateResponseAuthorizationLink `json:"authorization_link" api:"required,nullable"`
	// When the integration was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Credentials expiry time.
	CredentialsExpiry time.Time `json:"credentials_expiry" api:"required" format:"date-time"`
	// DLP Profiles enabled for the integration.
	DLPProfiles []string `json:"dlp_profiles" api:"required" format:"uuid"`
	// Health details with remediation hints.
	HealthDetails []map[string]interface{} `json:"health_details" api:"required"`
	// Whether the user paused the integration.
	IsPaused bool `json:"is_paused" api:"required"`
	// Last time the integration was hydrated.
	LastHydrated time.Time `json:"last_hydrated" api:"required" format:"date-time"`
	// Name of the integration.
	Name string `json:"name" api:"required"`
	// Integration status.
	Status string `json:"status" api:"required"`
	// When the integration was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Use cases enabled for the integration.
	UseCases []map[string]interface{}          `json:"use_cases" api:"required"`
	JSON     casbIntegrationUpdateResponseJSON `json:"-"`
}

// casbIntegrationUpdateResponseJSON contains the JSON metadata for the struct
// [CasbIntegrationUpdateResponse]
type casbIntegrationUpdateResponseJSON struct {
	ID                apijson.Field
	Application       apijson.Field
	AuthMethod        apijson.Field
	AuthorizationLink apijson.Field
	Created           apijson.Field
	CredentialsExpiry apijson.Field
	DLPProfiles       apijson.Field
	HealthDetails     apijson.Field
	IsPaused          apijson.Field
	LastHydrated      apijson.Field
	Name              apijson.Field
	Status            apijson.Field
	Updated           apijson.Field
	UseCases          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbIntegrationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Authorization link for the integration.
type CasbIntegrationUpdateResponseAuthorizationLink struct {
	Components map[string]interface{}                             `json:"components" api:"required,nullable"`
	Link       string                                             `json:"link" api:"required,nullable"`
	JSON       casbIntegrationUpdateResponseAuthorizationLinkJSON `json:"-"`
}

// casbIntegrationUpdateResponseAuthorizationLinkJSON contains the JSON metadata
// for the struct [CasbIntegrationUpdateResponseAuthorizationLink]
type casbIntegrationUpdateResponseAuthorizationLinkJSON struct {
	Components  apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbIntegrationUpdateResponseAuthorizationLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationUpdateResponseAuthorizationLinkJSON) RawJSON() string {
	return r.raw
}

// Serializer for v2 integration list responses.
type CasbIntegrationListResponse struct {
	// Integration ID.
	ID          string            `json:"id" api:"required" format:"uuid"`
	Application map[string]string `json:"application" api:"required"`
	// When the integration was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Whether the user paused the integration.
	IsPaused bool `json:"is_paused" api:"required"`
	// Name of the integration.
	Name string `json:"name" api:"required"`
	// Integration status.
	Status string `json:"status" api:"required"`
	// When the integration was last updated.
	Updated time.Time                       `json:"updated" api:"required" format:"date-time"`
	JSON    casbIntegrationListResponseJSON `json:"-"`
}

// casbIntegrationListResponseJSON contains the JSON metadata for the struct
// [CasbIntegrationListResponse]
type casbIntegrationListResponseJSON struct {
	ID          apijson.Field
	Application apijson.Field
	Created     apijson.Field
	IsPaused    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbIntegrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationListResponseJSON) RawJSON() string {
	return r.raw
}

// The requested item.
type CasbIntegrationGetResponse struct {
	// Integration ID.
	ID          string            `json:"id" api:"required" format:"uuid"`
	Application map[string]string `json:"application" api:"required"`
	// The integration's authentication method.
	AuthMethod map[string]string `json:"auth_method" api:"required,nullable"`
	// Authorization link for the integration.
	AuthorizationLink CasbIntegrationGetResponseAuthorizationLink `json:"authorization_link" api:"required,nullable"`
	// When the integration was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Credentials expiry time.
	CredentialsExpiry time.Time `json:"credentials_expiry" api:"required" format:"date-time"`
	// DLP Profiles enabled for the integration.
	DLPProfiles []string `json:"dlp_profiles" api:"required" format:"uuid"`
	// Health details with remediation hints.
	HealthDetails []map[string]interface{} `json:"health_details" api:"required"`
	// Whether the user paused the integration.
	IsPaused bool `json:"is_paused" api:"required"`
	// Last time the integration was hydrated.
	LastHydrated time.Time `json:"last_hydrated" api:"required" format:"date-time"`
	// Name of the integration.
	Name string `json:"name" api:"required"`
	// Integration status.
	Status string `json:"status" api:"required"`
	// When the integration was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Use cases enabled for the integration.
	UseCases []map[string]interface{}       `json:"use_cases" api:"required"`
	JSON     casbIntegrationGetResponseJSON `json:"-"`
}

// casbIntegrationGetResponseJSON contains the JSON metadata for the struct
// [CasbIntegrationGetResponse]
type casbIntegrationGetResponseJSON struct {
	ID                apijson.Field
	Application       apijson.Field
	AuthMethod        apijson.Field
	AuthorizationLink apijson.Field
	Created           apijson.Field
	CredentialsExpiry apijson.Field
	DLPProfiles       apijson.Field
	HealthDetails     apijson.Field
	IsPaused          apijson.Field
	LastHydrated      apijson.Field
	Name              apijson.Field
	Status            apijson.Field
	Updated           apijson.Field
	UseCases          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbIntegrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Authorization link for the integration.
type CasbIntegrationGetResponseAuthorizationLink struct {
	Components map[string]interface{}                          `json:"components" api:"required,nullable"`
	Link       string                                          `json:"link" api:"required,nullable"`
	JSON       casbIntegrationGetResponseAuthorizationLinkJSON `json:"-"`
}

// casbIntegrationGetResponseAuthorizationLinkJSON contains the JSON metadata for
// the struct [CasbIntegrationGetResponseAuthorizationLink]
type casbIntegrationGetResponseAuthorizationLinkJSON struct {
	Components  apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbIntegrationGetResponseAuthorizationLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationGetResponseAuthorizationLinkJSON) RawJSON() string {
	return r.raw
}

// The requested item.
type CasbIntegrationPauseResponse struct {
	// Integration ID.
	ID          string            `json:"id" api:"required" format:"uuid"`
	Application map[string]string `json:"application" api:"required"`
	// The integration's authentication method.
	AuthMethod map[string]string `json:"auth_method" api:"required,nullable"`
	// Authorization link for the integration.
	AuthorizationLink CasbIntegrationPauseResponseAuthorizationLink `json:"authorization_link" api:"required,nullable"`
	// When the integration was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Credentials expiry time.
	CredentialsExpiry time.Time `json:"credentials_expiry" api:"required" format:"date-time"`
	// DLP Profiles enabled for the integration.
	DLPProfiles []string `json:"dlp_profiles" api:"required" format:"uuid"`
	// Health details with remediation hints.
	HealthDetails []map[string]interface{} `json:"health_details" api:"required"`
	// Whether the user paused the integration.
	IsPaused bool `json:"is_paused" api:"required"`
	// Last time the integration was hydrated.
	LastHydrated time.Time `json:"last_hydrated" api:"required" format:"date-time"`
	// Name of the integration.
	Name string `json:"name" api:"required"`
	// Integration status.
	Status string `json:"status" api:"required"`
	// When the integration was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Use cases enabled for the integration.
	UseCases []map[string]interface{}         `json:"use_cases" api:"required"`
	JSON     casbIntegrationPauseResponseJSON `json:"-"`
}

// casbIntegrationPauseResponseJSON contains the JSON metadata for the struct
// [CasbIntegrationPauseResponse]
type casbIntegrationPauseResponseJSON struct {
	ID                apijson.Field
	Application       apijson.Field
	AuthMethod        apijson.Field
	AuthorizationLink apijson.Field
	Created           apijson.Field
	CredentialsExpiry apijson.Field
	DLPProfiles       apijson.Field
	HealthDetails     apijson.Field
	IsPaused          apijson.Field
	LastHydrated      apijson.Field
	Name              apijson.Field
	Status            apijson.Field
	Updated           apijson.Field
	UseCases          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbIntegrationPauseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationPauseResponseJSON) RawJSON() string {
	return r.raw
}

// Authorization link for the integration.
type CasbIntegrationPauseResponseAuthorizationLink struct {
	Components map[string]interface{}                            `json:"components" api:"required,nullable"`
	Link       string                                            `json:"link" api:"required,nullable"`
	JSON       casbIntegrationPauseResponseAuthorizationLinkJSON `json:"-"`
}

// casbIntegrationPauseResponseAuthorizationLinkJSON contains the JSON metadata for
// the struct [CasbIntegrationPauseResponseAuthorizationLink]
type casbIntegrationPauseResponseAuthorizationLinkJSON struct {
	Components  apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbIntegrationPauseResponseAuthorizationLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationPauseResponseAuthorizationLinkJSON) RawJSON() string {
	return r.raw
}

// The requested item.
type CasbIntegrationResumeResponse struct {
	// Integration ID.
	ID          string            `json:"id" api:"required" format:"uuid"`
	Application map[string]string `json:"application" api:"required"`
	// The integration's authentication method.
	AuthMethod map[string]string `json:"auth_method" api:"required,nullable"`
	// Authorization link for the integration.
	AuthorizationLink CasbIntegrationResumeResponseAuthorizationLink `json:"authorization_link" api:"required,nullable"`
	// When the integration was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Credentials expiry time.
	CredentialsExpiry time.Time `json:"credentials_expiry" api:"required" format:"date-time"`
	// DLP Profiles enabled for the integration.
	DLPProfiles []string `json:"dlp_profiles" api:"required" format:"uuid"`
	// Health details with remediation hints.
	HealthDetails []map[string]interface{} `json:"health_details" api:"required"`
	// Whether the user paused the integration.
	IsPaused bool `json:"is_paused" api:"required"`
	// Last time the integration was hydrated.
	LastHydrated time.Time `json:"last_hydrated" api:"required" format:"date-time"`
	// Name of the integration.
	Name string `json:"name" api:"required"`
	// Integration status.
	Status string `json:"status" api:"required"`
	// When the integration was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Use cases enabled for the integration.
	UseCases []map[string]interface{}          `json:"use_cases" api:"required"`
	JSON     casbIntegrationResumeResponseJSON `json:"-"`
}

// casbIntegrationResumeResponseJSON contains the JSON metadata for the struct
// [CasbIntegrationResumeResponse]
type casbIntegrationResumeResponseJSON struct {
	ID                apijson.Field
	Application       apijson.Field
	AuthMethod        apijson.Field
	AuthorizationLink apijson.Field
	Created           apijson.Field
	CredentialsExpiry apijson.Field
	DLPProfiles       apijson.Field
	HealthDetails     apijson.Field
	IsPaused          apijson.Field
	LastHydrated      apijson.Field
	Name              apijson.Field
	Status            apijson.Field
	Updated           apijson.Field
	UseCases          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbIntegrationResumeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationResumeResponseJSON) RawJSON() string {
	return r.raw
}

// Authorization link for the integration.
type CasbIntegrationResumeResponseAuthorizationLink struct {
	Components map[string]interface{}                             `json:"components" api:"required,nullable"`
	Link       string                                             `json:"link" api:"required,nullable"`
	JSON       casbIntegrationResumeResponseAuthorizationLinkJSON `json:"-"`
}

// casbIntegrationResumeResponseAuthorizationLinkJSON contains the JSON metadata
// for the struct [CasbIntegrationResumeResponseAuthorizationLink]
type casbIntegrationResumeResponseAuthorizationLinkJSON struct {
	Components  apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbIntegrationResumeResponseAuthorizationLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationResumeResponseAuthorizationLinkJSON) RawJSON() string {
	return r.raw
}

type CasbIntegrationNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Vendor/application slug (e.g., GOOGLE_WORKSPACE).
	//
	// - `ANTHROPIC` - ANTHROPIC
	// - `AWS` - AWS
	// - `BITBUCKET` - BITBUCKET
	// - `BOX` - BOX
	// - `CONFLUENCE` - CONFLUENCE
	// - `DROPBOX` - DROPBOX
	// - `GITHUB` - GITHUB
	// - `GOOGLE_CLOUD_PLATFORM` - GOOGLE_CLOUD_PLATFORM
	// - `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE
	// - `JIRA` - JIRA
	// - `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL
	// - `OPENAI` - OPENAI
	// - `SALESFORCE` - SALESFORCE
	// - `SERVICENOW` - SERVICENOW
	// - `SLACK` - SLACK
	Application param.Field[CasbIntegrationNewParamsApplication] `json:"application" api:"required"`
	// Credentials for the integration.
	Credentials param.Field[map[string]interface{}] `json:"credentials" api:"required"`
	// Name of the integration.
	Name param.Field[string] `json:"name" api:"required"`
	// Authentication method slug (uses default if omitted).
	AuthMethod param.Field[string] `json:"auth_method"`
	// List of DLP profile IDs to associate.
	DLPProfiles param.Field[[]string] `json:"dlp_profiles" format:"uuid"`
	// List of permission scopes (uses policy defaults if empty).
	Permissions param.Field[[]string] `json:"permissions"`
	// List of use case or feature slugs to enroll (e.g., ['casb', 'ces',
	// 'auto_remediation']).
	UseCases param.Field[[]CasbIntegrationNewParamsUseCase] `json:"use_cases"`
}

func (r CasbIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vendor/application slug (e.g., GOOGLE_WORKSPACE).
//
// - `ANTHROPIC` - ANTHROPIC
// - `AWS` - AWS
// - `BITBUCKET` - BITBUCKET
// - `BOX` - BOX
// - `CONFLUENCE` - CONFLUENCE
// - `DROPBOX` - DROPBOX
// - `GITHUB` - GITHUB
// - `GOOGLE_CLOUD_PLATFORM` - GOOGLE_CLOUD_PLATFORM
// - `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE
// - `JIRA` - JIRA
// - `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL
// - `OPENAI` - OPENAI
// - `SALESFORCE` - SALESFORCE
// - `SERVICENOW` - SERVICENOW
// - `SLACK` - SLACK
type CasbIntegrationNewParamsApplication string

const (
	CasbIntegrationNewParamsApplicationAnthropic           CasbIntegrationNewParamsApplication = "ANTHROPIC"
	CasbIntegrationNewParamsApplicationAws                 CasbIntegrationNewParamsApplication = "AWS"
	CasbIntegrationNewParamsApplicationBitbucket           CasbIntegrationNewParamsApplication = "BITBUCKET"
	CasbIntegrationNewParamsApplicationBox                 CasbIntegrationNewParamsApplication = "BOX"
	CasbIntegrationNewParamsApplicationConfluence          CasbIntegrationNewParamsApplication = "CONFLUENCE"
	CasbIntegrationNewParamsApplicationDropbox             CasbIntegrationNewParamsApplication = "DROPBOX"
	CasbIntegrationNewParamsApplicationGitHub              CasbIntegrationNewParamsApplication = "GITHUB"
	CasbIntegrationNewParamsApplicationGoogleCloudPlatform CasbIntegrationNewParamsApplication = "GOOGLE_CLOUD_PLATFORM"
	CasbIntegrationNewParamsApplicationGoogleWorkspace     CasbIntegrationNewParamsApplication = "GOOGLE_WORKSPACE"
	CasbIntegrationNewParamsApplicationJira                CasbIntegrationNewParamsApplication = "JIRA"
	CasbIntegrationNewParamsApplicationMicrosoftInternal   CasbIntegrationNewParamsApplication = "MICROSOFT_INTERNAL"
	CasbIntegrationNewParamsApplicationOpenAI              CasbIntegrationNewParamsApplication = "OPENAI"
	CasbIntegrationNewParamsApplicationSalesforce          CasbIntegrationNewParamsApplication = "SALESFORCE"
	CasbIntegrationNewParamsApplicationServicenow          CasbIntegrationNewParamsApplication = "SERVICENOW"
	CasbIntegrationNewParamsApplicationSlack               CasbIntegrationNewParamsApplication = "SLACK"
)

func (r CasbIntegrationNewParamsApplication) IsKnown() bool {
	switch r {
	case CasbIntegrationNewParamsApplicationAnthropic, CasbIntegrationNewParamsApplicationAws, CasbIntegrationNewParamsApplicationBitbucket, CasbIntegrationNewParamsApplicationBox, CasbIntegrationNewParamsApplicationConfluence, CasbIntegrationNewParamsApplicationDropbox, CasbIntegrationNewParamsApplicationGitHub, CasbIntegrationNewParamsApplicationGoogleCloudPlatform, CasbIntegrationNewParamsApplicationGoogleWorkspace, CasbIntegrationNewParamsApplicationJira, CasbIntegrationNewParamsApplicationMicrosoftInternal, CasbIntegrationNewParamsApplicationOpenAI, CasbIntegrationNewParamsApplicationSalesforce, CasbIntegrationNewParamsApplicationServicenow, CasbIntegrationNewParamsApplicationSlack:
		return true
	}
	return false
}

// - `casb` - casb
// - `ces` - ces
// - `auto_remediation` - auto_remediation
type CasbIntegrationNewParamsUseCase string

const (
	CasbIntegrationNewParamsUseCaseCasb            CasbIntegrationNewParamsUseCase = "casb"
	CasbIntegrationNewParamsUseCaseCes             CasbIntegrationNewParamsUseCase = "ces"
	CasbIntegrationNewParamsUseCaseAutoRemediation CasbIntegrationNewParamsUseCase = "auto_remediation"
)

func (r CasbIntegrationNewParamsUseCase) IsKnown() bool {
	switch r {
	case CasbIntegrationNewParamsUseCaseCasb, CasbIntegrationNewParamsUseCaseCes, CasbIntegrationNewParamsUseCaseAutoRemediation:
		return true
	}
	return false
}

type CasbIntegrationNewResponseEnvelope struct {
	// The requested item.
	Result CasbIntegrationNewResponse `json:"result" api:"required"`
	// Whether the request succeeded.
	Success bool `json:"success" api:"required"`
	// List of errors.
	Errors []map[string]interface{} `json:"errors"`
	// List of messages.
	Messages []string                               `json:"messages"`
	JSON     casbIntegrationNewResponseEnvelopeJSON `json:"-"`
}

// casbIntegrationNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CasbIntegrationNewResponseEnvelope]
type casbIntegrationNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbIntegrationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbIntegrationUpdateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Partial credential fields to merge with existing.
	Credentials param.Field[map[string]interface{}] `json:"credentials"`
	// List of DLP profile IDs to associate with the integration.
	DLPProfiles param.Field[[]string] `json:"dlp_profiles" format:"uuid"`
	// Name of the integration.
	Name param.Field[string] `json:"name"`
	// List of permission scopes granted to the integration.
	Permissions param.Field[[]string] `json:"permissions"`
	// List of use case or feature slugs to enroll (e.g., ['casb', 'ces',
	// 'auto_remediation']).
	UseCases param.Field[[]CasbIntegrationUpdateParamsUseCase] `json:"use_cases"`
}

func (r CasbIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// - `casb` - casb
// - `ces` - ces
// - `auto_remediation` - auto_remediation
type CasbIntegrationUpdateParamsUseCase string

const (
	CasbIntegrationUpdateParamsUseCaseCasb            CasbIntegrationUpdateParamsUseCase = "casb"
	CasbIntegrationUpdateParamsUseCaseCes             CasbIntegrationUpdateParamsUseCase = "ces"
	CasbIntegrationUpdateParamsUseCaseAutoRemediation CasbIntegrationUpdateParamsUseCase = "auto_remediation"
)

func (r CasbIntegrationUpdateParamsUseCase) IsKnown() bool {
	switch r {
	case CasbIntegrationUpdateParamsUseCaseCasb, CasbIntegrationUpdateParamsUseCaseCes, CasbIntegrationUpdateParamsUseCaseAutoRemediation:
		return true
	}
	return false
}

type CasbIntegrationUpdateResponseEnvelope struct {
	// The requested item.
	Result CasbIntegrationUpdateResponse `json:"result" api:"required"`
	// Whether the request succeeded.
	Success bool `json:"success" api:"required"`
	// List of errors.
	Errors []map[string]interface{} `json:"errors"`
	// List of messages.
	Messages []string                                  `json:"messages"`
	JSON     casbIntegrationUpdateResponseEnvelopeJSON `json:"-"`
}

// casbIntegrationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbIntegrationUpdateResponseEnvelope]
type casbIntegrationUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbIntegrationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbIntegrationListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter by application/vendor (e.g., GOOGLE_WORKSPACE, MICROSOFT_INTERNAL).
	Application param.Field[string] `query:"application"`
	// Direction to order results.
	Direction param.Field[CasbIntegrationListParamsDirection] `query:"direction"`
	// Filter by DLP enabled status (true/false).
	DLPEnabled param.Field[bool] `query:"dlp_enabled"`
	// Field to order results by.
	Order param.Field[CasbIntegrationListParamsOrder] `query:"order"`
	// Page number within the paginated result set.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PageSize param.Field[int64] `query:"page_size"`
	// Search integrations by name or application.
	Search param.Field[string] `query:"search"`
	// Filter by integration status.
	Status param.Field[CasbIntegrationListParamsStatus] `query:"status"`
	// Filter by enabled use cases (e.g., casb, ces). Matches integrations enrolled in
	// any of the specified values. Can be specified multiple times.
	UseCases param.Field[string] `query:"use_cases"`
}

// URLQuery serializes [CasbIntegrationListParams]'s query parameters as
// `url.Values`.
func (r CasbIntegrationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type CasbIntegrationListParamsDirection string

const (
	CasbIntegrationListParamsDirectionAsc  CasbIntegrationListParamsDirection = "asc"
	CasbIntegrationListParamsDirectionDesc CasbIntegrationListParamsDirection = "desc"
)

func (r CasbIntegrationListParamsDirection) IsKnown() bool {
	switch r {
	case CasbIntegrationListParamsDirectionAsc, CasbIntegrationListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by.
type CasbIntegrationListParamsOrder string

const (
	CasbIntegrationListParamsOrderApplication CasbIntegrationListParamsOrder = "application"
	CasbIntegrationListParamsOrderCreated     CasbIntegrationListParamsOrder = "created"
	CasbIntegrationListParamsOrderName        CasbIntegrationListParamsOrder = "name"
	CasbIntegrationListParamsOrderStatus      CasbIntegrationListParamsOrder = "status"
)

func (r CasbIntegrationListParamsOrder) IsKnown() bool {
	switch r {
	case CasbIntegrationListParamsOrderApplication, CasbIntegrationListParamsOrderCreated, CasbIntegrationListParamsOrderName, CasbIntegrationListParamsOrderStatus:
		return true
	}
	return false
}

// Filter by integration status.
type CasbIntegrationListParamsStatus string

const (
	CasbIntegrationListParamsStatusHealthy      CasbIntegrationListParamsStatus = "Healthy"
	CasbIntegrationListParamsStatusInitializing CasbIntegrationListParamsStatus = "Initializing"
	CasbIntegrationListParamsStatusOffline      CasbIntegrationListParamsStatus = "Offline"
	CasbIntegrationListParamsStatusUnhealthy    CasbIntegrationListParamsStatus = "Unhealthy"
)

func (r CasbIntegrationListParamsStatus) IsKnown() bool {
	switch r {
	case CasbIntegrationListParamsStatusHealthy, CasbIntegrationListParamsStatusInitializing, CasbIntegrationListParamsStatusOffline, CasbIntegrationListParamsStatusUnhealthy:
		return true
	}
	return false
}

type CasbIntegrationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CasbIntegrationGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CasbIntegrationGetResponseEnvelope struct {
	// The requested item.
	Result CasbIntegrationGetResponse `json:"result" api:"required"`
	// Whether the request succeeded.
	Success bool `json:"success" api:"required"`
	// List of errors.
	Errors []map[string]interface{} `json:"errors"`
	// List of messages.
	Messages []string                               `json:"messages"`
	JSON     casbIntegrationGetResponseEnvelopeJSON `json:"-"`
}

// casbIntegrationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CasbIntegrationGetResponseEnvelope]
type casbIntegrationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbIntegrationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbIntegrationPauseParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CasbIntegrationPauseResponseEnvelope struct {
	// The requested item.
	Result CasbIntegrationPauseResponse `json:"result" api:"required"`
	// Whether the request succeeded.
	Success bool `json:"success" api:"required"`
	// List of errors.
	Errors []map[string]interface{} `json:"errors"`
	// List of messages.
	Messages []string                                 `json:"messages"`
	JSON     casbIntegrationPauseResponseEnvelopeJSON `json:"-"`
}

// casbIntegrationPauseResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbIntegrationPauseResponseEnvelope]
type casbIntegrationPauseResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbIntegrationPauseResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationPauseResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbIntegrationResumeParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CasbIntegrationResumeResponseEnvelope struct {
	// The requested item.
	Result CasbIntegrationResumeResponse `json:"result" api:"required"`
	// Whether the request succeeded.
	Success bool `json:"success" api:"required"`
	// List of errors.
	Errors []map[string]interface{} `json:"errors"`
	// List of messages.
	Messages []string                                  `json:"messages"`
	JSON     casbIntegrationResumeResponseEnvelopeJSON `json:"-"`
}

// casbIntegrationResumeResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbIntegrationResumeResponseEnvelope]
type casbIntegrationResumeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbIntegrationResumeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbIntegrationResumeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
