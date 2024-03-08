// File generated from our OpenAPI spec by Stainless.

package alerting_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/alerting"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestV3PolicyNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Alerting.V3.Policies.New(context.TODO(), alerting.V3PolicyNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		AlertType: cloudflare.F(alerting.V3PolicyNewParamsAlertTypeUniversalSSLEventType),
		Enabled:   cloudflare.F(true),
		Mechanisms: cloudflare.F(map[string][]alerting.V3PolicyNewParamsMechanisms{
			"email": {{
				ID: cloudflare.F[alerting.V3PolicyNewParamsMechanismsID](shared.UnionString("test@example.com")),
			}},
			"pagerduty": {{
				ID: cloudflare.F[alerting.V3PolicyNewParamsMechanismsID](shared.UnionString("e8133a15-00a4-4d69-aec1-32f70c51f6e5")),
			}},
			"webhooks": {{
				ID: cloudflare.F[alerting.V3PolicyNewParamsMechanismsID](shared.UnionString("14cc1190-5d2b-4b98-a696-c424cb2ad05f")),
			}},
		}),
		Name:        cloudflare.F("SSL Notification Event Policy"),
		Description: cloudflare.F("Something describing the policy."),
		Filters: cloudflare.F(alerting.V3PolicyNewParamsFilters{
			Actions:                      cloudflare.F([]string{"string", "string", "string"}),
			AffectedASNs:                 cloudflare.F([]string{"string", "string", "string"}),
			AffectedComponents:           cloudflare.F([]string{"string", "string", "string"}),
			AffectedLocations:            cloudflare.F([]string{"string", "string", "string"}),
			AirportCode:                  cloudflare.F([]string{"string", "string", "string"}),
			AlertTriggerPreferences:      cloudflare.F([]string{"string", "string", "string"}),
			AlertTriggerPreferencesValue: cloudflare.F([]alerting.V3PolicyNewParamsFiltersAlertTriggerPreferencesValue{alerting.V3PolicyNewParamsFiltersAlertTriggerPreferencesValue99_0, alerting.V3PolicyNewParamsFiltersAlertTriggerPreferencesValue98_0, alerting.V3PolicyNewParamsFiltersAlertTriggerPreferencesValue97_0}),
			Enabled:                      cloudflare.F([]string{"string", "string", "string"}),
			Environment:                  cloudflare.F([]string{"string", "string", "string"}),
			Event:                        cloudflare.F([]string{"string", "string", "string"}),
			EventSource:                  cloudflare.F([]string{"string", "string", "string"}),
			EventType:                    cloudflare.F([]string{"string", "string", "string"}),
			GroupBy:                      cloudflare.F([]string{"string", "string", "string"}),
			HealthCheckID:                cloudflare.F([]string{"string", "string", "string"}),
			IncidentImpact:               cloudflare.F([]alerting.V3PolicyNewParamsFiltersIncidentImpact{alerting.V3PolicyNewParamsFiltersIncidentImpactIncidentImpactNone, alerting.V3PolicyNewParamsFiltersIncidentImpactIncidentImpactMinor, alerting.V3PolicyNewParamsFiltersIncidentImpactIncidentImpactMajor}),
			InputID:                      cloudflare.F([]string{"string", "string", "string"}),
			Limit:                        cloudflare.F([]string{"string", "string", "string"}),
			LogoTag:                      cloudflare.F([]string{"string", "string", "string"}),
			MegabitsPerSecond:            cloudflare.F([]string{"string", "string", "string"}),
			NewHealth:                    cloudflare.F([]string{"string", "string", "string"}),
			NewStatus:                    cloudflare.F([]string{"string", "string", "string"}),
			PacketsPerSecond:             cloudflare.F([]string{"string", "string", "string"}),
			PoolID:                       cloudflare.F([]string{"string", "string", "string"}),
			Product:                      cloudflare.F([]string{"string", "string", "string"}),
			ProjectID:                    cloudflare.F([]string{"string", "string", "string"}),
			Protocol:                     cloudflare.F([]string{"string", "string", "string"}),
			QueryTag:                     cloudflare.F([]string{"string", "string", "string"}),
			RequestsPerSecond:            cloudflare.F([]string{"string", "string", "string"}),
			Selectors:                    cloudflare.F([]string{"string", "string", "string"}),
			Services:                     cloudflare.F([]string{"string", "string", "string"}),
			Slo:                          cloudflare.F([]string{"99.9"}),
			Status:                       cloudflare.F([]string{"string", "string", "string"}),
			TargetHostname:               cloudflare.F([]string{"string", "string", "string"}),
			TargetIP:                     cloudflare.F([]string{"string", "string", "string"}),
			TargetZoneName:               cloudflare.F([]string{"string", "string", "string"}),
			TrafficExclusions:            cloudflare.F([]alerting.V3PolicyNewParamsFiltersTrafficExclusion{alerting.V3PolicyNewParamsFiltersTrafficExclusionSecurityEvents}),
			TunnelID:                     cloudflare.F([]string{"string", "string", "string"}),
			TunnelName:                   cloudflare.F([]string{"string", "string", "string"}),
			Where:                        cloudflare.F([]string{"string", "string", "string"}),
			Zones:                        cloudflare.F([]string{"string", "string", "string"}),
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV3PolicyUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Alerting.V3.Policies.Update(
		context.TODO(),
		"0da2b59e-f118-439d-8097-bdfb215203c9",
		alerting.V3PolicyUpdateParams{
			AccountID:   cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AlertType:   cloudflare.F(alerting.V3PolicyUpdateParamsAlertTypeUniversalSSLEventType),
			Description: cloudflare.F("Something describing the policy."),
			Enabled:     cloudflare.F(true),
			Filters: cloudflare.F(alerting.V3PolicyUpdateParamsFilters{
				Actions:                      cloudflare.F([]string{"string", "string", "string"}),
				AffectedASNs:                 cloudflare.F([]string{"string", "string", "string"}),
				AffectedComponents:           cloudflare.F([]string{"string", "string", "string"}),
				AffectedLocations:            cloudflare.F([]string{"string", "string", "string"}),
				AirportCode:                  cloudflare.F([]string{"string", "string", "string"}),
				AlertTriggerPreferences:      cloudflare.F([]string{"string", "string", "string"}),
				AlertTriggerPreferencesValue: cloudflare.F([]alerting.V3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue{alerting.V3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue99_0, alerting.V3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue98_0, alerting.V3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue97_0}),
				Enabled:                      cloudflare.F([]string{"string", "string", "string"}),
				Environment:                  cloudflare.F([]string{"string", "string", "string"}),
				Event:                        cloudflare.F([]string{"string", "string", "string"}),
				EventSource:                  cloudflare.F([]string{"string", "string", "string"}),
				EventType:                    cloudflare.F([]string{"string", "string", "string"}),
				GroupBy:                      cloudflare.F([]string{"string", "string", "string"}),
				HealthCheckID:                cloudflare.F([]string{"string", "string", "string"}),
				IncidentImpact:               cloudflare.F([]alerting.V3PolicyUpdateParamsFiltersIncidentImpact{alerting.V3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactNone, alerting.V3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactMinor, alerting.V3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactMajor}),
				InputID:                      cloudflare.F([]string{"string", "string", "string"}),
				Limit:                        cloudflare.F([]string{"string", "string", "string"}),
				LogoTag:                      cloudflare.F([]string{"string", "string", "string"}),
				MegabitsPerSecond:            cloudflare.F([]string{"string", "string", "string"}),
				NewHealth:                    cloudflare.F([]string{"string", "string", "string"}),
				NewStatus:                    cloudflare.F([]string{"string", "string", "string"}),
				PacketsPerSecond:             cloudflare.F([]string{"string", "string", "string"}),
				PoolID:                       cloudflare.F([]string{"string", "string", "string"}),
				Product:                      cloudflare.F([]string{"string", "string", "string"}),
				ProjectID:                    cloudflare.F([]string{"string", "string", "string"}),
				Protocol:                     cloudflare.F([]string{"string", "string", "string"}),
				QueryTag:                     cloudflare.F([]string{"string", "string", "string"}),
				RequestsPerSecond:            cloudflare.F([]string{"string", "string", "string"}),
				Selectors:                    cloudflare.F([]string{"string", "string", "string"}),
				Services:                     cloudflare.F([]string{"string", "string", "string"}),
				Slo:                          cloudflare.F([]string{"99.9"}),
				Status:                       cloudflare.F([]string{"string", "string", "string"}),
				TargetHostname:               cloudflare.F([]string{"string", "string", "string"}),
				TargetIP:                     cloudflare.F([]string{"string", "string", "string"}),
				TargetZoneName:               cloudflare.F([]string{"string", "string", "string"}),
				TrafficExclusions:            cloudflare.F([]alerting.V3PolicyUpdateParamsFiltersTrafficExclusion{alerting.V3PolicyUpdateParamsFiltersTrafficExclusionSecurityEvents}),
				TunnelID:                     cloudflare.F([]string{"string", "string", "string"}),
				TunnelName:                   cloudflare.F([]string{"string", "string", "string"}),
				Where:                        cloudflare.F([]string{"string", "string", "string"}),
				Zones:                        cloudflare.F([]string{"string", "string", "string"}),
			}),
			Mechanisms: cloudflare.F(map[string][]alerting.V3PolicyUpdateParamsMechanisms{
				"email": {{
					ID: cloudflare.F[alerting.V3PolicyUpdateParamsMechanismsID](shared.UnionString("test@example.com")),
				}},
				"pagerduty": {{
					ID: cloudflare.F[alerting.V3PolicyUpdateParamsMechanismsID](shared.UnionString("e8133a15-00a4-4d69-aec1-32f70c51f6e5")),
				}},
				"webhooks": {{
					ID: cloudflare.F[alerting.V3PolicyUpdateParamsMechanismsID](shared.UnionString("14cc1190-5d2b-4b98-a696-c424cb2ad05f")),
				}},
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

func TestV3PolicyList(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Alerting.V3.Policies.List(context.TODO(), alerting.V3PolicyListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV3PolicyDelete(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Alerting.V3.Policies.Delete(
		context.TODO(),
		"0da2b59e-f118-439d-8097-bdfb215203c9",
		alerting.V3PolicyDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestV3PolicyGet(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Alerting.V3.Policies.Get(
		context.TODO(),
		"0da2b59e-f118-439d-8097-bdfb215203c9",
		alerting.V3PolicyGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
