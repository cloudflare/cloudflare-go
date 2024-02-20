// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestAlertingV3PolicyNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Alerting.V3.Policies.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AlertingV3PolicyNewParams{
			AlertType: cloudflare.F(cloudflare.AlertingV3PolicyNewParamsAlertTypeUniversalSSLEventType),
			Enabled:   cloudflare.F(true),
			Mechanisms: cloudflare.F(map[string][]cloudflare.AlertingV3PolicyNewParamsMechanisms{
				"email": {{
					ID: cloudflare.F[cloudflare.AlertingV3PolicyNewParamsMechanismsID](shared.UnionString("test@example.com")),
				}},
				"pagerduty": {{
					ID: cloudflare.F[cloudflare.AlertingV3PolicyNewParamsMechanismsID](shared.UnionString("e8133a15-00a4-4d69-aec1-32f70c51f6e5")),
				}},
				"webhooks": {{
					ID: cloudflare.F[cloudflare.AlertingV3PolicyNewParamsMechanismsID](shared.UnionString("14cc1190-5d2b-4b98-a696-c424cb2ad05f")),
				}},
			}),
			Name:        cloudflare.F("SSL Notification Event Policy"),
			Description: cloudflare.F("Something describing the policy."),
			Filters: cloudflare.F(cloudflare.AlertingV3PolicyNewParamsFilters{
				Actions:                      cloudflare.F([]string{"string", "string", "string"}),
				AffectedAsns:                 cloudflare.F([]string{"string", "string", "string"}),
				AffectedComponents:           cloudflare.F([]string{"string", "string", "string"}),
				AffectedLocations:            cloudflare.F([]string{"string", "string", "string"}),
				AirportCode:                  cloudflare.F([]string{"string", "string", "string"}),
				AlertTriggerPreferences:      cloudflare.F([]string{"string", "string", "string"}),
				AlertTriggerPreferencesValue: cloudflare.F([]cloudflare.AlertingV3PolicyNewParamsFiltersAlertTriggerPreferencesValue{cloudflare.AlertingV3PolicyNewParamsFiltersAlertTriggerPreferencesValue99_0, cloudflare.AlertingV3PolicyNewParamsFiltersAlertTriggerPreferencesValue98_0, cloudflare.AlertingV3PolicyNewParamsFiltersAlertTriggerPreferencesValue97_0}),
				Enabled:                      cloudflare.F([]string{"string", "string", "string"}),
				Environment:                  cloudflare.F([]string{"string", "string", "string"}),
				Event:                        cloudflare.F([]string{"string", "string", "string"}),
				EventSource:                  cloudflare.F([]string{"string", "string", "string"}),
				EventType:                    cloudflare.F([]string{"string", "string", "string"}),
				GroupBy:                      cloudflare.F([]string{"string", "string", "string"}),
				HealthCheckID:                cloudflare.F([]string{"string", "string", "string"}),
				IncidentImpact:               cloudflare.F([]cloudflare.AlertingV3PolicyNewParamsFiltersIncidentImpact{cloudflare.AlertingV3PolicyNewParamsFiltersIncidentImpactIncidentImpactNone, cloudflare.AlertingV3PolicyNewParamsFiltersIncidentImpactIncidentImpactMinor, cloudflare.AlertingV3PolicyNewParamsFiltersIncidentImpactIncidentImpactMajor}),
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
				TrafficExclusions:            cloudflare.F([]cloudflare.AlertingV3PolicyNewParamsFiltersTrafficExclusion{cloudflare.AlertingV3PolicyNewParamsFiltersTrafficExclusionSecurityEvents}),
				TunnelID:                     cloudflare.F([]string{"string", "string", "string"}),
				TunnelName:                   cloudflare.F([]string{"string", "string", "string"}),
				Where:                        cloudflare.F([]string{"string", "string", "string"}),
				Zones:                        cloudflare.F([]string{"string", "string", "string"}),
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

func TestAlertingV3PolicyUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Alerting.V3.Policies.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0da2b59e-f118-439d-8097-bdfb215203c9",
		cloudflare.AlertingV3PolicyUpdateParams{
			AlertType:   cloudflare.F(cloudflare.AlertingV3PolicyUpdateParamsAlertTypeUniversalSSLEventType),
			Description: cloudflare.F("Something describing the policy."),
			Enabled:     cloudflare.F(true),
			Filters: cloudflare.F(cloudflare.AlertingV3PolicyUpdateParamsFilters{
				Actions:                      cloudflare.F([]string{"string", "string", "string"}),
				AffectedAsns:                 cloudflare.F([]string{"string", "string", "string"}),
				AffectedComponents:           cloudflare.F([]string{"string", "string", "string"}),
				AffectedLocations:            cloudflare.F([]string{"string", "string", "string"}),
				AirportCode:                  cloudflare.F([]string{"string", "string", "string"}),
				AlertTriggerPreferences:      cloudflare.F([]string{"string", "string", "string"}),
				AlertTriggerPreferencesValue: cloudflare.F([]cloudflare.AlertingV3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue{cloudflare.AlertingV3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue99_0, cloudflare.AlertingV3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue98_0, cloudflare.AlertingV3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue97_0}),
				Enabled:                      cloudflare.F([]string{"string", "string", "string"}),
				Environment:                  cloudflare.F([]string{"string", "string", "string"}),
				Event:                        cloudflare.F([]string{"string", "string", "string"}),
				EventSource:                  cloudflare.F([]string{"string", "string", "string"}),
				EventType:                    cloudflare.F([]string{"string", "string", "string"}),
				GroupBy:                      cloudflare.F([]string{"string", "string", "string"}),
				HealthCheckID:                cloudflare.F([]string{"string", "string", "string"}),
				IncidentImpact:               cloudflare.F([]cloudflare.AlertingV3PolicyUpdateParamsFiltersIncidentImpact{cloudflare.AlertingV3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactNone, cloudflare.AlertingV3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactMinor, cloudflare.AlertingV3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactMajor}),
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
				TrafficExclusions:            cloudflare.F([]cloudflare.AlertingV3PolicyUpdateParamsFiltersTrafficExclusion{cloudflare.AlertingV3PolicyUpdateParamsFiltersTrafficExclusionSecurityEvents}),
				TunnelID:                     cloudflare.F([]string{"string", "string", "string"}),
				TunnelName:                   cloudflare.F([]string{"string", "string", "string"}),
				Where:                        cloudflare.F([]string{"string", "string", "string"}),
				Zones:                        cloudflare.F([]string{"string", "string", "string"}),
			}),
			Mechanisms: cloudflare.F(map[string][]cloudflare.AlertingV3PolicyUpdateParamsMechanisms{
				"email": {{
					ID: cloudflare.F[cloudflare.AlertingV3PolicyUpdateParamsMechanismsID](shared.UnionString("test@example.com")),
				}},
				"pagerduty": {{
					ID: cloudflare.F[cloudflare.AlertingV3PolicyUpdateParamsMechanismsID](shared.UnionString("e8133a15-00a4-4d69-aec1-32f70c51f6e5")),
				}},
				"webhooks": {{
					ID: cloudflare.F[cloudflare.AlertingV3PolicyUpdateParamsMechanismsID](shared.UnionString("14cc1190-5d2b-4b98-a696-c424cb2ad05f")),
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

func TestAlertingV3PolicyList(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Alerting.V3.Policies.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAlertingV3PolicyDelete(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Alerting.V3.Policies.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0da2b59e-f118-439d-8097-bdfb215203c9",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAlertingV3PolicyGet(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Alerting.V3.Policies.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0da2b59e-f118-439d-8097-bdfb215203c9",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
