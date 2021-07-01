package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testWebhookID = "fe49ee055d23404e9d58f9110b210c8d"
	testPolicyID  = "6ec8a5145d0d2263a36fad55c03cb43d"
)

var (
	notificationTimestamp = time.Date(2021, 05, 01, 10, 47, 01, 01, time.UTC)
)

func TestAPI_GetEligibleNotificationDestinations(t *testing.T) {
	setup()
	defer teardown()

	expected := Mechanisms{
		Email:     MechanismMetaData{true, true, "email"},
		PagerDuty: MechanismMetaData{true, true, "pagerduty"},
		Webhooks:  MechanismMetaData{true, true, "webhooks"},
	}
	b, err := json.Marshal(expected)
	require.NoError(t, err)
	require.NotEmpty(t, b)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result":%s
}`, fmt.Sprintf(string(b)))
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/destinations/eligible", handler)

	actual, err := client.GetEligibleNotificationDestinations(context.Background(), testAccountID)
	require.Nil(t, err)
	require.NotNil(t, actual)
	assert.Equal(t, expected, actual.Result)
}
func TestAPI_GetAvailableNotificationTypes(t *testing.T) {
	setup()
	defer teardown()

	expected := make(NotificationsGroupedByProduct, 1)
	alert1 := AlertWithDescription{Type: "secondary_dns_zone_successfully_updated", DisplayName: "Secondary DNS Successfully Updated", Description: "Secondary zone transfers are succeeding, the zone has been updated."}
	alert2 := AlertWithDescription{Type: "secondary_dns_zone_validation_warning", DisplayName: "Secondary DNSSEC Validation Warning", Description: "The transferred DNSSEC zone is incorrectly configured."}
	expected["DNS"] = []AlertWithDescription{alert1, alert2}

	b, err := json.Marshal(expected)
	require.NoError(t, err)
	require.NotEmpty(t, b)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  									"success": true,
  									"errors": [],
  									"messages": [],
  									"result":%s
								}`,
			fmt.Sprintf(string(b)))
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/available_alerts", handler)

	actual, err := client.GetAvailableNotificationTypes(context.Background(), testAccountID)
	require.Nil(t, err)
	require.NotNil(t, actual)
	assert.Equal(t, expected, actual.Result)
}
func TestAPI_ListPagerDutyDestinations(t *testing.T) {
	setup()
	defer teardown()

	expected := PagerDutyResource{ID: "valid-uuid", Name: "my pagerduty connection"}
	b, err := json.Marshal(expected)
	require.NoError(t, err)
	require.NotEmpty(t, b)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		_, err = fmt.Fprintf(w, `{
  									"success": true,
  									"errors": [],
  									"messages": [],
  									"result":%s
								}`,
			fmt.Sprintf(string(b)))
		require.NoError(t, err)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/destinations/pagerduty", handler)

	actual, err := client.ListPagerDutyDestinations(context.Background(), testAccountID)
	require.Nil(t, err)
	require.NotNil(t, actual)
	assert.Equal(t, expected, actual.Result)
}
func TestAPI_DeletePagerDutyDestinations(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		_, err := fmt.Fprintf(w, `{
  									"success": true,
  									"errors": [],
  									"messages": []
								}`,
		)
		require.NoError(t, err)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/destinations/pagerduty", handler)

	actual, err := client.DeletePagerDutyDestinations(context.Background(), testAccountID)
	require.Nil(t, err)
	require.NotNil(t, actual)
	assert.True(t, actual.Success)
}
func TestAPI_CreateNotificationPolicy(t *testing.T) {
	setup()
	defer teardown()

	mechanisms := make(map[string]MechanismIntegrations)
	mechanisms["email"] = []MechanismData{{Name: "email to send notification", ID: "test@gmail.com"}}
	policy := Policy{
		Description: "Notifies when my zones are under attack",
		Name:        "CF DOS attack alert - L4",
		Enabled:     true,
		AlertType:   "dos_attack_l4",
		Mechanisms:  mechanisms,
		Conditions:  nil,
		Filters:     nil,
	}
	b, err := json.Marshal(policy)
	require.NoError(t, err)
	require.NotEmpty(t, b)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		_, err = fmt.Fprintf(w, `{
  									"success": true,
  									"errors": [],
  									"messages": [],
  									"result":%s
								}`,
			fmt.Sprintf(string(b)))
		require.NoError(t, err)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/policies", handler)
	res, err := client.CreateNotificationPolicy(context.Background(), testAccountID, policy)
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestAPI_GetNotificationPolicy(t *testing.T) {
	setup()
	defer teardown()

	mechanisms := make(map[string]MechanismIntegrations)
	mechanisms["email"] = []MechanismData{{Name: "email to send notification", ID: "test@gmail.com"}}
	policy := Policy{
		ID:          testPolicyID,
		Description: "Notifies when my zones are under attack",
		Name:        "CF DOS attack alert - L4",
		Enabled:     true,
		AlertType:   "dos_attack_l4",
		Mechanisms:  mechanisms,
		Conditions:  nil,
		Filters:     nil,
		Created:     notificationTimestamp,
		Modified:    notificationTimestamp,
	}
	b, err := json.Marshal(policy)
	require.NoError(t, err)
	require.NotEmpty(t, b)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		_, err = fmt.Fprintf(w, `{
  									"success": true,
  									"errors": [],
  									"messages": [],
  									"result":%s
								}`,
			fmt.Sprintf(string(b)))
		require.NoError(t, err)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/policies/"+testPolicyID, handler)

	res, err := client.GetNotificationPolicy(context.Background(), testAccountID, testPolicyID)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, policy, res.Result)

}

func TestAPI_ListNotificationPolicies(t *testing.T) {
	setup()
	defer teardown()

	mechanisms := make(map[string]MechanismIntegrations)
	mechanisms["email"] = []MechanismData{{Name: "email to send notification", ID: "test@gmail.com"}}
	policy := Policy{
		ID:          testPolicyID,
		Description: "Notifies when my zones are under attack",
		Name:        "CF DOS attack alert - L4",
		Enabled:     true,
		AlertType:   "dos_attack_l4",
		Mechanisms:  mechanisms,
		Conditions:  nil,
		Filters:     nil,
		Created:     time.Date(2021, 05, 01, 10, 47, 01, 01, time.UTC),
		Modified:    time.Date(2021, 05, 01, 10, 47, 01, 01, time.UTC),
	}
	policies := []Policy{
		policy,
	}
	b, err := json.Marshal(policies)
	require.NoError(t, err)
	require.NotEmpty(t, b)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		_, err = fmt.Fprintf(w, `{
  									"success": true,
  									"errors": [],
  									"messages": [],
  									"result":%s
								}`,
			fmt.Sprintf(string(b)))
		require.NoError(t, err)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/policies", handler)

	res, err := client.ListNotificationPolicies(context.Background(), testAccountID)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, policies, res.Result)
}

func TestAPI_UpdateNotificationPolicy(t *testing.T) {
	setup()
	defer teardown()

	mechanisms := make(map[string]MechanismIntegrations)
	mechanisms["email"] = []MechanismData{{Name: "email to send notification", ID: "test@gmail.com"}}
	policy := Policy{
		ID:          testPolicyID,
		Description: "Notifies when my zones are under attack",
		Name:        "CF DOS attack alert - L4",
		Enabled:     true,
		AlertType:   "dos_attack_l4",
		Mechanisms:  mechanisms,
		Conditions:  nil,
		Filters:     nil,
		Created:     time.Date(2021, 05, 01, 10, 47, 01, 01, time.UTC),
		Modified:    time.Date(2021, 05, 01, 10, 47, 01, 01, time.UTC),
	}
	b, err := json.Marshal(policy)
	require.NoError(t, err)
	require.NotEmpty(t, b)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		_, err = fmt.Fprintf(w, `{
  									"success": true,
  									"errors": [],
  									"messages": [],
  									"result":%s
								}`,
			fmt.Sprintf(string(b)))
		require.NoError(t, err)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/policies/"+testPolicyID, handler)

	res, err := client.UpdateNotificationPolicy(context.Background(), testAccountID, &policy)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, testPolicyID, res.Result.ID)

}
func TestAPI_DeleteNotificationPolicy(t *testing.T) {
	setup()
	defer teardown()

	result := Resource{ID: testPolicyID}
	b, err := json.Marshal(result)
	require.NoError(t, err)
	require.NotNil(t, b)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		_, err = fmt.Fprintf(w, `{
  									"success": true,
  									"errors": [],
  									"messages": [],
  									"result":%s
								}`,
			fmt.Sprintf(string(b)))
		require.NoError(t, err)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/policies/"+testPolicyID, handler)

	res, err := client.DeleteNotificationPolicy(context.Background(), testAccountID, testPolicyID)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, testPolicyID, res.Result.ID)

}
func TestAPI_CreateNotificationWebhooks(t *testing.T) {
	setup()
	defer teardown()

	webhook := UpsertWebhooks{
		Name:   "my test webhook",
		URL:    "https://example.com",
		Secret: "mischief-managed", // optional
	}

	result := Resource{ID: testWebhookID}

	b, err := json.Marshal(result)
	require.NoError(t, err)
	require.NotEmpty(t, b)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		_, err = fmt.Fprintf(w, `{
  									"success": true,
  									"errors": [],
  									"messages": [],
  									"result":%s
								}`,
			fmt.Sprintf(string(b)))
		require.NoError(t, err)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/destinations/webhooks", handler)

	res, err := client.CreateNotificationWebhooks(context.Background(), testAccountID, &webhook)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, testWebhookID, res.Result.ID)

}
func TestAPI_ListNotificationWebhooks(t *testing.T) {
	setup()
	defer teardown()

	webhook := WebhookIntegration{
		ID:          testWebhookID,
		Name:        "my test webhook",
		URL:         "https://example.com",
		Type:        "generic",
		CreatedAt:   notificationTimestamp,
		LastSuccess: &notificationTimestamp,
		LastFailure: &notificationTimestamp,
	}
	webhooks := []WebhookIntegration{webhook}
	b, err := json.Marshal(webhooks)
	require.NoError(t, err)
	require.NotEmpty(t, b)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		_, err = fmt.Fprintf(w, `{
  									"success": true,
  									"errors": [],
  									"messages": [],
  									"result":%s
								}`,
			fmt.Sprintf(string(b)))
		require.NoError(t, err)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/destinations/webhooks", handler)

	res, err := client.ListNotificationWebhooks(context.Background(), testAccountID)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, webhooks, res.Result)
}

func TestAPI_GetNotificationWebhooks(t *testing.T) {
	setup()
	defer teardown()

	webhook := WebhookIntegration{
		ID:          testWebhookID,
		Name:        "my test webhook",
		URL:         "https://example.com",
		Type:        "generic",
		CreatedAt:   notificationTimestamp,
		LastSuccess: &notificationTimestamp,
		LastFailure: &notificationTimestamp,
	}
	b, err := json.Marshal(webhook)
	require.NoError(t, err)
	require.NotEmpty(t, b)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		_, err = fmt.Fprintf(w, `{
  									"success": true,
  									"errors": [],
  									"messages": [],
  									"result":%s
								}`,
			fmt.Sprintf(string(b)))
		require.NoError(t, err)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/destinations/webhooks/"+testWebhookID, handler)

	res, err := client.GetNotificationWebhooks(context.Background(), testAccountID, testWebhookID)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, webhook, res.Result)
}

func TestAPI_UpdateNotificationWebhooks(t *testing.T) {
	setup()
	defer teardown()

	result := Resource{ID: testWebhookID}
	b, err := json.Marshal(result)
	require.NoError(t, err)
	require.NotEmpty(t, b)

	webhook := UpsertWebhooks{
		Name:   "my test webhook with a new name",
		URL:    "https://example.com",
		Secret: "mischief-managed",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		_, err = fmt.Fprintf(w, `{
  									"success": true,
  									"errors": [],
  									"messages": [],
  									"result":%s
								}`,
			fmt.Sprintf(string(b)))
		require.NoError(t, err)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/destinations/webhooks/"+testWebhookID, handler)

	res, err := client.UpdateNotificationWebhooks(context.Background(), testAccountID, testWebhookID, &webhook)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, testWebhookID, res.Result.ID)

}

func TestAPI_DeleteNotificationWebhooks(t *testing.T) {
	setup()
	defer teardown()

	result := Resource{ID: testWebhookID}
	b, err := json.Marshal(result)
	require.NoError(t, err)
	require.NotEmpty(t, b)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		_, err = fmt.Fprintf(w, `{
  									"success": true,
  									"errors": [],
  									"messages": [],
  									"result":%s
								}`,
			fmt.Sprintf(string(b)))
		require.NoError(t, err)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/alerting/v3/destinations/webhooks/"+testWebhookID, handler)

	res, err := client.DeleteNotificationWebhooks(context.Background(), testAccountID, testWebhookID)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, testWebhookID, res.Result.ID)
}
