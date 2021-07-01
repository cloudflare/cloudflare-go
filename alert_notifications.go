package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// MechanismData holds a single public facing mechanism data integation
type MechanismData struct {
	Name string `json:"id"`
	ID   string `json:"name"`
}

// MechanismIntegrations is a list of all the integrations of a certain mechanism type
// e.g. all email integrations
type MechanismIntegrations []MechanismData

// Policy represents the notification policy created along with the destinations
type Policy struct {
	ID          string                           `json:"id"`
	Name        string                           `json:"name"`
	Description string                           `json:"description"`
	Enabled     bool                             `json:"enabled"`
	AlertType   string                           `json:"alert_type"`
	Mechanisms  map[string]MechanismIntegrations `json:"mechanisms"`
	Created     time.Time                        `json:"created"`
	Modified    time.Time                        `json:"modified"`
	Conditions  map[string]interface{}           `json:"conditions"`
	Filters     map[string][]string              `json:"filters"`
}

// PoliciesResponse holds the response for listing all notification policies for an account
type PoliciesResponse struct {
	Response   Response
	ResultInfo ResultInfo
	Result     []Policy
}

// PolicyResponse holds the response type when a single policy is retrieved
type PolicyResponse struct {
	Response Response
	Result   Policy
}

// WebhookIntegration describes the webhook information along with its status
type WebhookIntegration struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	URL         string     `json:"url"`
	CreatedAt   time.Time  `json:"created_at"`
	LastSuccess *time.Time `json:"last_success"`
	LastFailure *time.Time `json:"last_failure"`
}

// WebhookResponse describes a single webhook retrieved
type WebhookResponse struct {
	Response   Response
	ResultInfo ResultInfo
	Result     WebhookIntegration
}

// WebhooksResponse describes a list of webhooks retrieved
type WebhooksResponse struct {
	Response   Response
	ResultInfo ResultInfo
	Result     []WebhookIntegration
}

// UpsertWebhooks describes a valid webhook request
type UpsertWebhooks struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Secret string `json:"secret"`
}

// PagerDutyResource describes a PagerDuty integration
type PagerDutyResource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// PagerDutyResponse describes the PagerDuty integration retrieved
type PagerDutyResponse struct {
	Response   Response
	ResultInfo ResultInfo
	Result     PagerDutyResource
}

// Resource describes the id of an inserted/updated/deleted resource
type Resource struct {
	ID string
}

// SaveResponse is returned when a resource is inserted/updated/deleted
type SaveResponse struct {
	Response Response
	Result   Resource
}

// MechanismMetaData represents the state of the delivery mechanism
type MechanismMetaData struct {
	Eligible bool   `json:"eligible"`
	Ready    bool   `json:"ready"`
	Type     string `json:"type"`
}

// Mechanisms are the different possible delivery mechanisms
type Mechanisms struct {
	Email     MechanismMetaData `json:"email"`
	PagerDuty MechanismMetaData `json:"pagerduty"`
	Webhooks  MechanismMetaData `json:"webhooks,omitempty"`
}

// EligibilityResponse describes the eligible mechanisms that can be configured for a notification
type EligibilityResponse struct {
	Response Response
	Result   Mechanisms
}

// NotificationsGroupedByProduct are grouped by products
type NotificationsGroupedByProduct map[string][]AlertWithDescription

// AlertWithDescription represents the alert/notification available
type AlertWithDescription struct {
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// AvailableAlertsResponse describes the available alerts/notifications grouped by products
type AvailableAlertsResponse struct {
	Response Response
	Result   NotificationsGroupedByProduct
}

// ListNotificationPolicies will return the notification policies
// created by a user for a specific account
//
// API Reference: https://api.cloudflare.com/#notification-policies-properties
func (api *API) ListNotificationPolicies(ctx context.Context, accountID string) (PoliciesResponse, error) {
	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/policies", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return PoliciesResponse{}, err
	}
	var r PoliciesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

// GetNotificationPolicy returns a specific created by a user, given the account id and the policy id
//
// API Reference: https://api.cloudflare.com/#notification-policies-properties
func (api *API) GetNotificationPolicy(ctx context.Context, accountID, policyID string) (PolicyResponse, error) {
	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/policies/%s", accountID, policyID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return PolicyResponse{}, err
	}
	var r PolicyResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

// CreateNotificationPolicy creates a notification policy for an account
//
// API Reference: https://api.cloudflare.com/#notification-policies-create-notification-policy
func (api *API) CreateNotificationPolicy(ctx context.Context, accountID string, policy Policy) (SaveResponse, error) {

	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/policies", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, baseURL, policy)
	if err != nil {
		return SaveResponse{}, err
	}
	return unmarshalSaveResponse(res)
}

// UpdateNotificationPolicy updates a notification policy, given the account id and the policy id and returns the policy id
//
// API Reference: https://api.cloudflare.com/#notification-policies-update-notification-policy
func (api *API) UpdateNotificationPolicy(ctx context.Context, accountID string, policy *Policy) (SaveResponse, error) {
	if policy == nil {
		return SaveResponse{}, fmt.Errorf("policy cannot be nil")
	}
	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/policies/%s", accountID, policy.ID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, baseURL, policy)
	if err != nil {
		return SaveResponse{}, err
	}
	return unmarshalSaveResponse(res)
}

// DeleteNotificationPolicy deletes a notification policy for an account
//
// API Reference: https://api.cloudflare.com/#notification-policies-delete-notification-policy
func (api *API) DeleteNotificationPolicy(ctx context.Context, accountID, policyID string) (SaveResponse, error) {
	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/policies/%s", accountID, policyID)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, baseURL, nil)
	if err != nil {
		return SaveResponse{}, err
	}
	return unmarshalSaveResponse(res)
}

// ListNotificationWebhooks will return the webhook destinations configured for an account
//
// API Reference: https://api.cloudflare.com/#notification-webhooks-list-webhooks
func (api *API) ListNotificationWebhooks(ctx context.Context, accountID string) (WebhooksResponse, error) {
	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/destinations/webhooks", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return WebhooksResponse{}, err
	}
	var r WebhooksResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}
	return r, nil

}

// CreateNotificationWebhooks will help connect a webhooks destination.
// A test message will be sent to the webhooks endpoint during creation.
// If added successfully, the webhooks can be setup as a destination mechanism while creating policies.
// Notifications will be posted to this URL.
//
// API Reference: https://api.cloudflare.com/#notification-webhooks-create-webhook
func (api *API) CreateNotificationWebhooks(ctx context.Context, accountID string, webhooks *UpsertWebhooks) (SaveResponse, error) {
	if webhooks == nil {
		return SaveResponse{}, fmt.Errorf("webhooks cannot be nil")
	}
	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/destinations/webhooks", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, baseURL, webhooks)
	if err != nil {
		return SaveResponse{}, err
	}

	return unmarshalSaveResponse(res)
}

// GetNotificationWebhooks will return a specific webhook destination, given the account and webhooks ids
//
// API Reference: https://api.cloudflare.com/#notification-webhooks-get-webhook
func (api *API) GetNotificationWebhooks(ctx context.Context, accountID, webhookID string) (WebhookResponse, error) {
	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/destinations/webhooks/%s", accountID, webhookID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return WebhookResponse{}, err
	}
	var r WebhookResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

// UpdateNotificationWebhooks will update a particular webhook's name, given the account and webhooks ids.
// The webhook url and secret cannot be updated
//
// API Reference: https://api.cloudflare.com/#notification-webhooks-update-webhook
func (api *API) UpdateNotificationWebhooks(ctx context.Context, accountID, webhookID string, webhooks *UpsertWebhooks) (SaveResponse, error) {
	if webhooks == nil {
		return SaveResponse{}, fmt.Errorf("webhooks cannot be nil")
	}
	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/destinations/webhooks/%s", accountID, webhookID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, baseURL, webhooks)
	if err != nil {
		return SaveResponse{}, err
	}

	return unmarshalSaveResponse(res)
}

// DeleteNotificationWebhooks will delete a webhook, given the account and webhooks ids.
// Deleting the webhooks will remove it from any connected notification policies.
//
// API Reference: https://api.cloudflare.com/#notification-webhooks-delete-webhook
func (api *API) DeleteNotificationWebhooks(ctx context.Context, accountID, webhookID string) (SaveResponse, error) {
	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/destinations/webhooks/%s", accountID, webhookID)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, baseURL, nil)
	if err != nil {
		return SaveResponse{}, err
	}

	return unmarshalSaveResponse(res)
}

// ListPagerDutyDestinations will return the pagerduty destinations configured for an account
//
// API Reference: https://api.cloudflare.com/#notification-destinations-with-pagerduty-list-pagerduty-services
func (api *API) ListPagerDutyDestinations(ctx context.Context, accountID string) (PagerDutyResponse, error) {
	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/destinations/pagerduty", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return PagerDutyResponse{}, err
	}
	var r PagerDutyResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

// DeletePagerDutyDestinations will delete the pagerduty destination connected for an account
//
// API Reference: https://api.cloudflare.com/#notification-destinations-with-pagerduty-delete-pagerduty-destinations
func (api *API) DeletePagerDutyDestinations(ctx context.Context, accountID string) (Response, error) {
	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/destinations/pagerduty", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, baseURL, nil)
	if err != nil {
		return Response{}, err
	}
	var r Response
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

// GetEligibleNotificationDestinations will return the types of destinations an account is eligible to configure
//
// API Reference: https://api.cloudflare.com/#notification-mechanism-eligibility-properties
func (api *API) GetEligibleNotificationDestinations(ctx context.Context, accountID string) (EligibilityResponse, error) {
	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/destinations/eligible", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return EligibilityResponse{}, err
	}
	var r EligibilityResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

// GetAvailableNotificationTypes will return the alert types available for a given account
//
// API Reference: https://api.cloudflare.com/#notification-mechanism-eligibility-properties
func (api *API) GetAvailableNotificationTypes(ctx context.Context, accountID string) (AvailableAlertsResponse, error) {
	baseURL := fmt.Sprintf("/accounts/%s/alerting/v3/available_alerts", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return AvailableAlertsResponse{}, err
	}
	var r AvailableAlertsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

// unmarshal will unmarshal bytes and return a SaveResponse
func unmarshalSaveResponse(res []byte) (SaveResponse, error) {
	var r SaveResponse
	err := json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}
