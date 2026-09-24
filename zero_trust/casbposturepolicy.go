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

// CasbPosturePolicyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbPosturePolicyService] method instead.
type CasbPosturePolicyService struct {
	Options []option.RequestOption
}

// NewCasbPosturePolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCasbPosturePolicyService(opts ...option.RequestOption) (r *CasbPosturePolicyService) {
	r = &CasbPosturePolicyService{}
	r.Options = opts
	return
}

// Creates a new policy configuration that defines automated actions to be executed
// when security findings are detected. A policy can include multiple remediation
// and/or webhook actions that will be triggered automatically.
func (r *CasbPosturePolicyService) New(ctx context.Context, params CasbPosturePolicyNewParams, opts ...option.RequestOption) (res *CasbPosturePolicyNewResponse, err error) {
	var env CasbPosturePolicyNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/policies", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an existing policy configuration and replaces its actions.
func (r *CasbPosturePolicyService) Update(ctx context.Context, policyID string, params CasbPosturePolicyUpdateParams, opts ...option.RequestOption) (res *CasbPosturePolicyUpdateResponse, err error) {
	var env CasbPosturePolicyUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/policies/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a list of integration-scoped policy configurations for the given
// account. This endpoint supports cursor based pagination. By default, results are
// returned in sorted order based on created_at.
func (r *CasbPosturePolicyService) List(ctx context.Context, params CasbPosturePolicyListParams, opts ...option.RequestOption) (res *pagination.CursorPaginationAfter[CasbPosturePolicyListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/policies", params.AccountID)
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

// Returns a list of integration-scoped policy configurations for the given
// account. This endpoint supports cursor based pagination. By default, results are
// returned in sorted order based on created_at.
func (r *CasbPosturePolicyService) ListAutoPaging(ctx context.Context, params CasbPosturePolicyListParams, opts ...option.RequestOption) *pagination.CursorPaginationAfterAutoPager[CasbPosturePolicyListResponse] {
	return pagination.NewCursorPaginationAfterAutoPager(r.List(ctx, params, opts...))
}

// Deletes a policy configuration.
func (r *CasbPosturePolicyService) Delete(ctx context.Context, policyID string, body CasbPosturePolicyDeleteParams, opts ...option.RequestOption) (res *CasbPosturePolicyDeleteResponse, err error) {
	var env CasbPosturePolicyDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/policies/%s", body.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves the details of a specific policy configuration, including its
// associated remediation and webhook actions.
func (r *CasbPosturePolicyService) Get(ctx context.Context, policyID string, query CasbPosturePolicyGetParams, opts ...option.RequestOption) (res *CasbPosturePolicyGetResponse, err error) {
	var env CasbPosturePolicyGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/policies/%s", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Response body for a policy configuration.
type CasbPosturePolicyNewResponse struct {
	// Unique identifier for the policy configuration.
	ID string `json:"id" api:"required" format:"uuid"`
	// The actions configured for this policy.
	Actions CasbPosturePolicyNewResponseActions `json:"actions" api:"required"`
	// When true, the policy applies to all integrations for the account. When false,
	// it applies only to the specified integration_ids.
	AppliesToAllIntegrations bool `json:"applies_to_all_integrations" api:"required"`
	// Timestamp when the policy was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// User-set description of what this policy does. Limited to 1000 characters.
	Description string `json:"description" api:"required"`
	// Display name for the policy configuration. Limited to 255 characters.
	DisplayName string `json:"display_name" api:"required"`
	// Whether the policy is enabled. Derived from disabled_at (enabled when
	// disabled_at is unset).
	Enabled bool `json:"enabled" api:"required"`
	// The finding type this policy is associated with. Immutable after creation;
	// changing it replaces the policy.
	FindingTypeID string `json:"finding_type_id" api:"required" format:"uuid"`
	// The integrations this policy applies to.
	IntegrationIDs []string `json:"integration_ids" api:"required" format:"uuid"`
	// Timestamp when the policy was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Timestamp when the policy was disabled. Omitted from the response when the
	// policy is enabled.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Timestamp of the most recent successful policy invocation. Omitted from the
	// response when the policy has never been successfully triggered. Only populated
	// on GET responses; absent on responses from create/update endpoints.
	LastTriggeredAt time.Time                        `json:"last_triggered_at" format:"date-time"`
	JSON            casbPosturePolicyNewResponseJSON `json:"-"`
}

// casbPosturePolicyNewResponseJSON contains the JSON metadata for the struct
// [CasbPosturePolicyNewResponse]
type casbPosturePolicyNewResponseJSON struct {
	ID                       apijson.Field
	Actions                  apijson.Field
	AppliesToAllIntegrations apijson.Field
	CreatedAt                apijson.Field
	Description              apijson.Field
	DisplayName              apijson.Field
	Enabled                  apijson.Field
	FindingTypeID            apijson.Field
	IntegrationIDs           apijson.Field
	UpdatedAt                apijson.Field
	DisabledAt               apijson.Field
	LastTriggeredAt          apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CasbPosturePolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

// The actions configured for this policy.
type CasbPosturePolicyNewResponseActions struct {
	// List of remediation types that will be executed.
	RemediationTypes []CasbPosturePolicyNewResponseActionsRemediationType `json:"remediation_types" api:"required"`
	// List of webhook configurations that will be triggered.
	WebhookConfigs []CasbPosturePolicyNewResponseActionsWebhookConfig `json:"webhook_configs" api:"required"`
	JSON           casbPosturePolicyNewResponseActionsJSON            `json:"-"`
}

// casbPosturePolicyNewResponseActionsJSON contains the JSON metadata for the
// struct [CasbPosturePolicyNewResponseActions]
type casbPosturePolicyNewResponseActionsJSON struct {
	RemediationTypes apijson.Field
	WebhookConfigs   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPosturePolicyNewResponseActions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyNewResponseActionsJSON) RawJSON() string {
	return r.raw
}

// A remediation type configured for the policy.
type CasbPosturePolicyNewResponseActionsRemediationType struct {
	// Display name/label of the remediation type.
	DisplayName string `json:"display_name" api:"required"`
	// The system name of the remediation type.
	RemediationType string `json:"remediation_type" api:"required"`
	// Unique identifier for the remediation type.
	RemediationTypeID string                                                 `json:"remediation_type_id" api:"required" format:"uuid"`
	JSON              casbPosturePolicyNewResponseActionsRemediationTypeJSON `json:"-"`
}

// casbPosturePolicyNewResponseActionsRemediationTypeJSON contains the JSON
// metadata for the struct [CasbPosturePolicyNewResponseActionsRemediationType]
type casbPosturePolicyNewResponseActionsRemediationTypeJSON struct {
	DisplayName       apijson.Field
	RemediationType   apijson.Field
	RemediationTypeID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbPosturePolicyNewResponseActionsRemediationType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyNewResponseActionsRemediationTypeJSON) RawJSON() string {
	return r.raw
}

// A webhook configuration associated with the policy.
type CasbPosturePolicyNewResponseActionsWebhookConfig struct {
	// Display name/label of the webhook configuration.
	DisplayName string `json:"display_name" api:"required"`
	// Unique identifier for the webhook configuration.
	WebhookConfigID string                                               `json:"webhook_config_id" api:"required" format:"uuid"`
	JSON            casbPosturePolicyNewResponseActionsWebhookConfigJSON `json:"-"`
}

// casbPosturePolicyNewResponseActionsWebhookConfigJSON contains the JSON metadata
// for the struct [CasbPosturePolicyNewResponseActionsWebhookConfig]
type casbPosturePolicyNewResponseActionsWebhookConfigJSON struct {
	DisplayName     apijson.Field
	WebhookConfigID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CasbPosturePolicyNewResponseActionsWebhookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyNewResponseActionsWebhookConfigJSON) RawJSON() string {
	return r.raw
}

// Response body for a policy configuration.
type CasbPosturePolicyUpdateResponse struct {
	// Unique identifier for the policy configuration.
	ID string `json:"id" api:"required" format:"uuid"`
	// The actions configured for this policy.
	Actions CasbPosturePolicyUpdateResponseActions `json:"actions" api:"required"`
	// When true, the policy applies to all integrations for the account. When false,
	// it applies only to the specified integration_ids.
	AppliesToAllIntegrations bool `json:"applies_to_all_integrations" api:"required"`
	// Timestamp when the policy was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// User-set description of what this policy does. Limited to 1000 characters.
	Description string `json:"description" api:"required"`
	// Display name for the policy configuration. Limited to 255 characters.
	DisplayName string `json:"display_name" api:"required"`
	// Whether the policy is enabled. Derived from disabled_at (enabled when
	// disabled_at is unset).
	Enabled bool `json:"enabled" api:"required"`
	// The finding type this policy is associated with. Immutable after creation;
	// changing it replaces the policy.
	FindingTypeID string `json:"finding_type_id" api:"required" format:"uuid"`
	// The integrations this policy applies to.
	IntegrationIDs []string `json:"integration_ids" api:"required" format:"uuid"`
	// Timestamp when the policy was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Timestamp when the policy was disabled. Omitted from the response when the
	// policy is enabled.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Timestamp of the most recent successful policy invocation. Omitted from the
	// response when the policy has never been successfully triggered. Only populated
	// on GET responses; absent on responses from create/update endpoints.
	LastTriggeredAt time.Time                           `json:"last_triggered_at" format:"date-time"`
	JSON            casbPosturePolicyUpdateResponseJSON `json:"-"`
}

// casbPosturePolicyUpdateResponseJSON contains the JSON metadata for the struct
// [CasbPosturePolicyUpdateResponse]
type casbPosturePolicyUpdateResponseJSON struct {
	ID                       apijson.Field
	Actions                  apijson.Field
	AppliesToAllIntegrations apijson.Field
	CreatedAt                apijson.Field
	Description              apijson.Field
	DisplayName              apijson.Field
	Enabled                  apijson.Field
	FindingTypeID            apijson.Field
	IntegrationIDs           apijson.Field
	UpdatedAt                apijson.Field
	DisabledAt               apijson.Field
	LastTriggeredAt          apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CasbPosturePolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The actions configured for this policy.
type CasbPosturePolicyUpdateResponseActions struct {
	// List of remediation types that will be executed.
	RemediationTypes []CasbPosturePolicyUpdateResponseActionsRemediationType `json:"remediation_types" api:"required"`
	// List of webhook configurations that will be triggered.
	WebhookConfigs []CasbPosturePolicyUpdateResponseActionsWebhookConfig `json:"webhook_configs" api:"required"`
	JSON           casbPosturePolicyUpdateResponseActionsJSON            `json:"-"`
}

// casbPosturePolicyUpdateResponseActionsJSON contains the JSON metadata for the
// struct [CasbPosturePolicyUpdateResponseActions]
type casbPosturePolicyUpdateResponseActionsJSON struct {
	RemediationTypes apijson.Field
	WebhookConfigs   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPosturePolicyUpdateResponseActions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyUpdateResponseActionsJSON) RawJSON() string {
	return r.raw
}

// A remediation type configured for the policy.
type CasbPosturePolicyUpdateResponseActionsRemediationType struct {
	// Display name/label of the remediation type.
	DisplayName string `json:"display_name" api:"required"`
	// The system name of the remediation type.
	RemediationType string `json:"remediation_type" api:"required"`
	// Unique identifier for the remediation type.
	RemediationTypeID string                                                    `json:"remediation_type_id" api:"required" format:"uuid"`
	JSON              casbPosturePolicyUpdateResponseActionsRemediationTypeJSON `json:"-"`
}

// casbPosturePolicyUpdateResponseActionsRemediationTypeJSON contains the JSON
// metadata for the struct [CasbPosturePolicyUpdateResponseActionsRemediationType]
type casbPosturePolicyUpdateResponseActionsRemediationTypeJSON struct {
	DisplayName       apijson.Field
	RemediationType   apijson.Field
	RemediationTypeID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbPosturePolicyUpdateResponseActionsRemediationType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyUpdateResponseActionsRemediationTypeJSON) RawJSON() string {
	return r.raw
}

// A webhook configuration associated with the policy.
type CasbPosturePolicyUpdateResponseActionsWebhookConfig struct {
	// Display name/label of the webhook configuration.
	DisplayName string `json:"display_name" api:"required"`
	// Unique identifier for the webhook configuration.
	WebhookConfigID string                                                  `json:"webhook_config_id" api:"required" format:"uuid"`
	JSON            casbPosturePolicyUpdateResponseActionsWebhookConfigJSON `json:"-"`
}

// casbPosturePolicyUpdateResponseActionsWebhookConfigJSON contains the JSON
// metadata for the struct [CasbPosturePolicyUpdateResponseActionsWebhookConfig]
type casbPosturePolicyUpdateResponseActionsWebhookConfigJSON struct {
	DisplayName     apijson.Field
	WebhookConfigID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CasbPosturePolicyUpdateResponseActionsWebhookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyUpdateResponseActionsWebhookConfigJSON) RawJSON() string {
	return r.raw
}

// Response body for a policy configuration.
type CasbPosturePolicyListResponse struct {
	// Unique identifier for the policy configuration.
	ID string `json:"id" api:"required" format:"uuid"`
	// The actions configured for this policy.
	Actions CasbPosturePolicyListResponseActions `json:"actions" api:"required"`
	// When true, the policy applies to all integrations for the account. When false,
	// it applies only to the specified integration_ids.
	AppliesToAllIntegrations bool `json:"applies_to_all_integrations" api:"required"`
	// Timestamp when the policy was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// User-set description of what this policy does. Limited to 1000 characters.
	Description string `json:"description" api:"required"`
	// Display name for the policy configuration. Limited to 255 characters.
	DisplayName string `json:"display_name" api:"required"`
	// Whether the policy is enabled. Derived from disabled_at (enabled when
	// disabled_at is unset).
	Enabled bool `json:"enabled" api:"required"`
	// The finding type this policy is associated with. Immutable after creation;
	// changing it replaces the policy.
	FindingTypeID string `json:"finding_type_id" api:"required" format:"uuid"`
	// The integrations this policy applies to.
	IntegrationIDs []string `json:"integration_ids" api:"required" format:"uuid"`
	// Timestamp when the policy was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Timestamp when the policy was disabled. Omitted from the response when the
	// policy is enabled.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Timestamp of the most recent successful policy invocation. Omitted from the
	// response when the policy has never been successfully triggered. Only populated
	// on GET responses; absent on responses from create/update endpoints.
	LastTriggeredAt time.Time                         `json:"last_triggered_at" format:"date-time"`
	JSON            casbPosturePolicyListResponseJSON `json:"-"`
}

// casbPosturePolicyListResponseJSON contains the JSON metadata for the struct
// [CasbPosturePolicyListResponse]
type casbPosturePolicyListResponseJSON struct {
	ID                       apijson.Field
	Actions                  apijson.Field
	AppliesToAllIntegrations apijson.Field
	CreatedAt                apijson.Field
	Description              apijson.Field
	DisplayName              apijson.Field
	Enabled                  apijson.Field
	FindingTypeID            apijson.Field
	IntegrationIDs           apijson.Field
	UpdatedAt                apijson.Field
	DisabledAt               apijson.Field
	LastTriggeredAt          apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CasbPosturePolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyListResponseJSON) RawJSON() string {
	return r.raw
}

// The actions configured for this policy.
type CasbPosturePolicyListResponseActions struct {
	// List of remediation types that will be executed.
	RemediationTypes []CasbPosturePolicyListResponseActionsRemediationType `json:"remediation_types" api:"required"`
	// List of webhook configurations that will be triggered.
	WebhookConfigs []CasbPosturePolicyListResponseActionsWebhookConfig `json:"webhook_configs" api:"required"`
	JSON           casbPosturePolicyListResponseActionsJSON            `json:"-"`
}

// casbPosturePolicyListResponseActionsJSON contains the JSON metadata for the
// struct [CasbPosturePolicyListResponseActions]
type casbPosturePolicyListResponseActionsJSON struct {
	RemediationTypes apijson.Field
	WebhookConfigs   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPosturePolicyListResponseActions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyListResponseActionsJSON) RawJSON() string {
	return r.raw
}

// A remediation type configured for the policy.
type CasbPosturePolicyListResponseActionsRemediationType struct {
	// Display name/label of the remediation type.
	DisplayName string `json:"display_name" api:"required"`
	// The system name of the remediation type.
	RemediationType string `json:"remediation_type" api:"required"`
	// Unique identifier for the remediation type.
	RemediationTypeID string                                                  `json:"remediation_type_id" api:"required" format:"uuid"`
	JSON              casbPosturePolicyListResponseActionsRemediationTypeJSON `json:"-"`
}

// casbPosturePolicyListResponseActionsRemediationTypeJSON contains the JSON
// metadata for the struct [CasbPosturePolicyListResponseActionsRemediationType]
type casbPosturePolicyListResponseActionsRemediationTypeJSON struct {
	DisplayName       apijson.Field
	RemediationType   apijson.Field
	RemediationTypeID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbPosturePolicyListResponseActionsRemediationType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyListResponseActionsRemediationTypeJSON) RawJSON() string {
	return r.raw
}

// A webhook configuration associated with the policy.
type CasbPosturePolicyListResponseActionsWebhookConfig struct {
	// Display name/label of the webhook configuration.
	DisplayName string `json:"display_name" api:"required"`
	// Unique identifier for the webhook configuration.
	WebhookConfigID string                                                `json:"webhook_config_id" api:"required" format:"uuid"`
	JSON            casbPosturePolicyListResponseActionsWebhookConfigJSON `json:"-"`
}

// casbPosturePolicyListResponseActionsWebhookConfigJSON contains the JSON metadata
// for the struct [CasbPosturePolicyListResponseActionsWebhookConfig]
type casbPosturePolicyListResponseActionsWebhookConfigJSON struct {
	DisplayName     apijson.Field
	WebhookConfigID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CasbPosturePolicyListResponseActionsWebhookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyListResponseActionsWebhookConfigJSON) RawJSON() string {
	return r.raw
}

// Response from DeletePolicy operation.
type CasbPosturePolicyDeleteResponse struct {
	// ID of the policy deleted.
	ID   string                              `json:"id" api:"required" format:"uuid"`
	JSON casbPosturePolicyDeleteResponseJSON `json:"-"`
}

// casbPosturePolicyDeleteResponseJSON contains the JSON metadata for the struct
// [CasbPosturePolicyDeleteResponse]
type casbPosturePolicyDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPosturePolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Response body for a policy configuration.
type CasbPosturePolicyGetResponse struct {
	// Unique identifier for the policy configuration.
	ID string `json:"id" api:"required" format:"uuid"`
	// The actions configured for this policy.
	Actions CasbPosturePolicyGetResponseActions `json:"actions" api:"required"`
	// When true, the policy applies to all integrations for the account. When false,
	// it applies only to the specified integration_ids.
	AppliesToAllIntegrations bool `json:"applies_to_all_integrations" api:"required"`
	// Timestamp when the policy was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// User-set description of what this policy does. Limited to 1000 characters.
	Description string `json:"description" api:"required"`
	// Display name for the policy configuration. Limited to 255 characters.
	DisplayName string `json:"display_name" api:"required"`
	// Whether the policy is enabled. Derived from disabled_at (enabled when
	// disabled_at is unset).
	Enabled bool `json:"enabled" api:"required"`
	// The finding type this policy is associated with. Immutable after creation;
	// changing it replaces the policy.
	FindingTypeID string `json:"finding_type_id" api:"required" format:"uuid"`
	// The integrations this policy applies to.
	IntegrationIDs []string `json:"integration_ids" api:"required" format:"uuid"`
	// Timestamp when the policy was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Timestamp when the policy was disabled. Omitted from the response when the
	// policy is enabled.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Timestamp of the most recent successful policy invocation. Omitted from the
	// response when the policy has never been successfully triggered. Only populated
	// on GET responses; absent on responses from create/update endpoints.
	LastTriggeredAt time.Time                        `json:"last_triggered_at" format:"date-time"`
	JSON            casbPosturePolicyGetResponseJSON `json:"-"`
}

// casbPosturePolicyGetResponseJSON contains the JSON metadata for the struct
// [CasbPosturePolicyGetResponse]
type casbPosturePolicyGetResponseJSON struct {
	ID                       apijson.Field
	Actions                  apijson.Field
	AppliesToAllIntegrations apijson.Field
	CreatedAt                apijson.Field
	Description              apijson.Field
	DisplayName              apijson.Field
	Enabled                  apijson.Field
	FindingTypeID            apijson.Field
	IntegrationIDs           apijson.Field
	UpdatedAt                apijson.Field
	DisabledAt               apijson.Field
	LastTriggeredAt          apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CasbPosturePolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

// The actions configured for this policy.
type CasbPosturePolicyGetResponseActions struct {
	// List of remediation types that will be executed.
	RemediationTypes []CasbPosturePolicyGetResponseActionsRemediationType `json:"remediation_types" api:"required"`
	// List of webhook configurations that will be triggered.
	WebhookConfigs []CasbPosturePolicyGetResponseActionsWebhookConfig `json:"webhook_configs" api:"required"`
	JSON           casbPosturePolicyGetResponseActionsJSON            `json:"-"`
}

// casbPosturePolicyGetResponseActionsJSON contains the JSON metadata for the
// struct [CasbPosturePolicyGetResponseActions]
type casbPosturePolicyGetResponseActionsJSON struct {
	RemediationTypes apijson.Field
	WebhookConfigs   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPosturePolicyGetResponseActions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyGetResponseActionsJSON) RawJSON() string {
	return r.raw
}

// A remediation type configured for the policy.
type CasbPosturePolicyGetResponseActionsRemediationType struct {
	// Display name/label of the remediation type.
	DisplayName string `json:"display_name" api:"required"`
	// The system name of the remediation type.
	RemediationType string `json:"remediation_type" api:"required"`
	// Unique identifier for the remediation type.
	RemediationTypeID string                                                 `json:"remediation_type_id" api:"required" format:"uuid"`
	JSON              casbPosturePolicyGetResponseActionsRemediationTypeJSON `json:"-"`
}

// casbPosturePolicyGetResponseActionsRemediationTypeJSON contains the JSON
// metadata for the struct [CasbPosturePolicyGetResponseActionsRemediationType]
type casbPosturePolicyGetResponseActionsRemediationTypeJSON struct {
	DisplayName       apijson.Field
	RemediationType   apijson.Field
	RemediationTypeID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbPosturePolicyGetResponseActionsRemediationType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyGetResponseActionsRemediationTypeJSON) RawJSON() string {
	return r.raw
}

// A webhook configuration associated with the policy.
type CasbPosturePolicyGetResponseActionsWebhookConfig struct {
	// Display name/label of the webhook configuration.
	DisplayName string `json:"display_name" api:"required"`
	// Unique identifier for the webhook configuration.
	WebhookConfigID string                                               `json:"webhook_config_id" api:"required" format:"uuid"`
	JSON            casbPosturePolicyGetResponseActionsWebhookConfigJSON `json:"-"`
}

// casbPosturePolicyGetResponseActionsWebhookConfigJSON contains the JSON metadata
// for the struct [CasbPosturePolicyGetResponseActionsWebhookConfig]
type casbPosturePolicyGetResponseActionsWebhookConfigJSON struct {
	DisplayName     apijson.Field
	WebhookConfigID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CasbPosturePolicyGetResponseActionsWebhookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyGetResponseActionsWebhookConfigJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Actions to execute when this policy is triggered, grouped by action type. A
	// policy must contain at least one action across all groups and may include at
	// most one remediation.
	Actions param.Field[CasbPosturePolicyNewParamsActions] `json:"actions" api:"required"`
	// When true, the policy applies to all integrations for the account. When false,
	// integration_ids must be provided.
	AppliesToAllIntegrations param.Field[bool] `json:"applies_to_all_integrations" api:"required"`
	// Display name for the policy configuration.
	DisplayName param.Field[string] `json:"display_name" api:"required"`
	// Boolean specifying if the policy is enabled or disabled.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// The finding type this policy is associated with. All remediation actions must
	// match this finding type.
	FindingTypeID param.Field[string] `json:"finding_type_id" api:"required" format:"uuid"`
	// Optional description of what this policy does.
	Description param.Field[string] `json:"description"`
	// The integrations this policy applies to. Required when
	// applies_to_all_integrations is false.
	IntegrationIDs param.Field[[]string] `json:"integration_ids" format:"uuid"`
}

func (r CasbPosturePolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions to execute when this policy is triggered, grouped by action type. A
// policy must contain at least one action across all groups and may include at
// most one remediation.
type CasbPosturePolicyNewParamsActions struct {
	// Remediation actions to execute (at most one).
	RemediationTypes param.Field[[]CasbPosturePolicyNewParamsActionsRemediationType] `json:"remediation_types"`
	// Webhook actions to execute.
	WebhookConfigs param.Field[[]CasbPosturePolicyNewParamsActionsWebhookConfig] `json:"webhook_configs"`
}

func (r CasbPosturePolicyNewParamsActions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A remediation action to be executed.
type CasbPosturePolicyNewParamsActionsRemediationType struct {
	// The ID of the remediation type to execute.
	RemediationTypeID param.Field[string] `json:"remediation_type_id" api:"required" format:"uuid"`
}

func (r CasbPosturePolicyNewParamsActionsRemediationType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A webhook action to be executed.
type CasbPosturePolicyNewParamsActionsWebhookConfig struct {
	// The ID of the webhook configuration to use.
	WebhookConfigID param.Field[string] `json:"webhook_config_id" api:"required" format:"uuid"`
}

func (r CasbPosturePolicyNewParamsActionsWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Common response structure for all API endpoints.
type CasbPosturePolicyNewResponseEnvelope struct {
	Errors   []CasbPosturePolicyNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPosturePolicyNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Response body for a policy configuration.
	Result CasbPosturePolicyNewResponse             `json:"result"`
	JSON   casbPosturePolicyNewResponseEnvelopeJSON `json:"-"`
}

// casbPosturePolicyNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPosturePolicyNewResponseEnvelope]
type casbPosturePolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPosturePolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyNewResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                           `json:"documentation_url" format:"uri"`
	Source           CasbPosturePolicyNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPosturePolicyNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPosturePolicyNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CasbPosturePolicyNewResponseEnvelopeErrors]
type casbPosturePolicyNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPosturePolicyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyNewResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                               `json:"pointer"`
	JSON    casbPosturePolicyNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPosturePolicyNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CasbPosturePolicyNewResponseEnvelopeErrorsSource]
type casbPosturePolicyNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPosturePolicyNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyNewResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                             `json:"documentation_url" format:"uri"`
	Source           CasbPosturePolicyNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPosturePolicyNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPosturePolicyNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CasbPosturePolicyNewResponseEnvelopeMessages]
type casbPosturePolicyNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPosturePolicyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyNewResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                 `json:"pointer"`
	JSON    casbPosturePolicyNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPosturePolicyNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CasbPosturePolicyNewResponseEnvelopeMessagesSource]
type casbPosturePolicyNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPosturePolicyNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyUpdateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Actions to execute when this policy is triggered, grouped by action type. A
	// policy must contain at least one action across all groups and may include at
	// most one remediation.
	Actions param.Field[CasbPosturePolicyUpdateParamsActions] `json:"actions" api:"required"`
	// When true, the policy applies to all integrations for the account. When false,
	// integration_ids must be provided.
	AppliesToAllIntegrations param.Field[bool] `json:"applies_to_all_integrations" api:"required"`
	// Display name for the policy configuration.
	DisplayName param.Field[string] `json:"display_name" api:"required"`
	// Boolean specifying if the policy is enabled or disabled.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Optional description of what this policy does.
	Description param.Field[string] `json:"description"`
	// The integrations this policy applies to. Required when
	// applies_to_all_integrations is false.
	IntegrationIDs param.Field[[]string] `json:"integration_ids" format:"uuid"`
}

func (r CasbPosturePolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions to execute when this policy is triggered, grouped by action type. A
// policy must contain at least one action across all groups and may include at
// most one remediation.
type CasbPosturePolicyUpdateParamsActions struct {
	// Remediation actions to execute (at most one).
	RemediationTypes param.Field[[]CasbPosturePolicyUpdateParamsActionsRemediationType] `json:"remediation_types"`
	// Webhook actions to execute.
	WebhookConfigs param.Field[[]CasbPosturePolicyUpdateParamsActionsWebhookConfig] `json:"webhook_configs"`
}

func (r CasbPosturePolicyUpdateParamsActions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A remediation action to be executed.
type CasbPosturePolicyUpdateParamsActionsRemediationType struct {
	// The ID of the remediation type to execute.
	RemediationTypeID param.Field[string] `json:"remediation_type_id" api:"required" format:"uuid"`
}

func (r CasbPosturePolicyUpdateParamsActionsRemediationType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A webhook action to be executed.
type CasbPosturePolicyUpdateParamsActionsWebhookConfig struct {
	// The ID of the webhook configuration to use.
	WebhookConfigID param.Field[string] `json:"webhook_config_id" api:"required" format:"uuid"`
}

func (r CasbPosturePolicyUpdateParamsActionsWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Common response structure for all API endpoints.
type CasbPosturePolicyUpdateResponseEnvelope struct {
	Errors   []CasbPosturePolicyUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPosturePolicyUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Response body for a policy configuration.
	Result CasbPosturePolicyUpdateResponse             `json:"result"`
	JSON   casbPosturePolicyUpdateResponseEnvelopeJSON `json:"-"`
}

// casbPosturePolicyUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPosturePolicyUpdateResponseEnvelope]
type casbPosturePolicyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPosturePolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyUpdateResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                              `json:"documentation_url" format:"uri"`
	Source           CasbPosturePolicyUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPosturePolicyUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPosturePolicyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CasbPosturePolicyUpdateResponseEnvelopeErrors]
type casbPosturePolicyUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPosturePolicyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyUpdateResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                  `json:"pointer"`
	JSON    casbPosturePolicyUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPosturePolicyUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CasbPosturePolicyUpdateResponseEnvelopeErrorsSource]
type casbPosturePolicyUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPosturePolicyUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyUpdateResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                `json:"documentation_url" format:"uri"`
	Source           CasbPosturePolicyUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPosturePolicyUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPosturePolicyUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CasbPosturePolicyUpdateResponseEnvelopeMessages]
type casbPosturePolicyUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPosturePolicyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyUpdateResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                    `json:"pointer"`
	JSON    casbPosturePolicyUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPosturePolicyUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CasbPosturePolicyUpdateResponseEnvelopeMessagesSource]
type casbPosturePolicyUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPosturePolicyUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Cursor for pagination. Obtained from the `result_info.cursor` field of a
	// previous response.
	Cursor param.Field[string] `query:"cursor"`
}

// URLQuery serializes [CasbPosturePolicyListParams]'s query parameters as
// `url.Values`.
func (r CasbPosturePolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CasbPosturePolicyDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

// Common response structure for all API endpoints.
type CasbPosturePolicyDeleteResponseEnvelope struct {
	Errors   []CasbPosturePolicyDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPosturePolicyDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Response from DeletePolicy operation.
	Result CasbPosturePolicyDeleteResponse             `json:"result"`
	JSON   casbPosturePolicyDeleteResponseEnvelopeJSON `json:"-"`
}

// casbPosturePolicyDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPosturePolicyDeleteResponseEnvelope]
type casbPosturePolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPosturePolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyDeleteResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                              `json:"documentation_url" format:"uri"`
	Source           CasbPosturePolicyDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPosturePolicyDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPosturePolicyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CasbPosturePolicyDeleteResponseEnvelopeErrors]
type casbPosturePolicyDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPosturePolicyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyDeleteResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                  `json:"pointer"`
	JSON    casbPosturePolicyDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPosturePolicyDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CasbPosturePolicyDeleteResponseEnvelopeErrorsSource]
type casbPosturePolicyDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPosturePolicyDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyDeleteResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                `json:"documentation_url" format:"uri"`
	Source           CasbPosturePolicyDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPosturePolicyDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPosturePolicyDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CasbPosturePolicyDeleteResponseEnvelopeMessages]
type casbPosturePolicyDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPosturePolicyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyDeleteResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                    `json:"pointer"`
	JSON    casbPosturePolicyDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPosturePolicyDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CasbPosturePolicyDeleteResponseEnvelopeMessagesSource]
type casbPosturePolicyDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPosturePolicyDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

// Common response structure for all API endpoints.
type CasbPosturePolicyGetResponseEnvelope struct {
	Errors   []CasbPosturePolicyGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPosturePolicyGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Response body for a policy configuration.
	Result CasbPosturePolicyGetResponse             `json:"result"`
	JSON   casbPosturePolicyGetResponseEnvelopeJSON `json:"-"`
}

// casbPosturePolicyGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPosturePolicyGetResponseEnvelope]
type casbPosturePolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPosturePolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyGetResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                           `json:"documentation_url" format:"uri"`
	Source           CasbPosturePolicyGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPosturePolicyGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPosturePolicyGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CasbPosturePolicyGetResponseEnvelopeErrors]
type casbPosturePolicyGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPosturePolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyGetResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                               `json:"pointer"`
	JSON    casbPosturePolicyGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPosturePolicyGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CasbPosturePolicyGetResponseEnvelopeErrorsSource]
type casbPosturePolicyGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPosturePolicyGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyGetResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                             `json:"documentation_url" format:"uri"`
	Source           CasbPosturePolicyGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPosturePolicyGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPosturePolicyGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CasbPosturePolicyGetResponseEnvelopeMessages]
type casbPosturePolicyGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPosturePolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPosturePolicyGetResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                 `json:"pointer"`
	JSON    casbPosturePolicyGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPosturePolicyGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CasbPosturePolicyGetResponseEnvelopeMessagesSource]
type casbPosturePolicyGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPosturePolicyGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPosturePolicyGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
