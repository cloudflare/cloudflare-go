// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestAccountAlertingV3PolicyGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Alerting.V3s.Policies.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAlertingV3PolicyUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Alerting.V3s.Policies.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccountAlertingV3PolicyUpdateParams{
			AlertType:   cloudflare.F("universal_ssl_event_type"),
			Description: cloudflare.F("Something describing the policy."),
			Enabled:     cloudflare.F(true),
			Filters: cloudflare.F[any](map[string]interface{}{
				"slo": map[string]interface{}{
					"0": "99.9",
				},
			}),
			Mechanisms: cloudflare.F[any](map[string]interface{}{
				"email": map[string]interface{}{
					"0": map[string]interface{}{
						"id": "test@example.com",
					},
				},
				"pagerduty": map[string]interface{}{
					"0": map[string]interface{}{
						"id": "e8133a15-00a4-4d69-aec1-32f70c51f6e5",
					},
				},
				"webhooks": map[string]interface{}{
					"0": map[string]interface{}{
						"id": "14cc1190-5d2b-4b98-a696-c424cb2ad05f",
					},
				},
			}),
			Name: cloudflare.F("SSL Notification Event Policy"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAlertingV3PolicyDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Alerting.V3s.Policies.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAlertingV3PolicyNotificationPoliciesNewANotificationPolicyWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Alerting.V3s.Policies.NotificationPoliciesNewANotificationPolicy(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountAlertingV3PolicyNotificationPoliciesNewANotificationPolicyParams{
			AlertType: cloudflare.F("universal_ssl_event_type"),
			Enabled:   cloudflare.F(true),
			Mechanisms: cloudflare.F[any](map[string]interface{}{
				"email": map[string]interface{}{
					"0": map[string]interface{}{
						"id": "test@example.com",
					},
				},
				"pagerduty": map[string]interface{}{
					"0": map[string]interface{}{
						"id": "e8133a15-00a4-4d69-aec1-32f70c51f6e5",
					},
				},
				"webhooks": map[string]interface{}{
					"0": map[string]interface{}{
						"id": "14cc1190-5d2b-4b98-a696-c424cb2ad05f",
					},
				},
			}),
			Name:        cloudflare.F("SSL Notification Event Policy"),
			Description: cloudflare.F("Something describing the policy."),
			Filters: cloudflare.F[any](map[string]interface{}{
				"slo": map[string]interface{}{
					"0": "99.9",
				},
			}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAlertingV3PolicyNotificationPoliciesListNotificationPolicies(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Alerting.V3s.Policies.NotificationPoliciesListNotificationPolicies(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
