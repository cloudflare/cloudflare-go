// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alerting_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/alerting"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestPolicyNewWithOptionalParams(t *testing.T) {
	t.Skip("prism errors - https://github.com/cloudflare/cloudflare-python/actions/runs/9327225061/job/25676826349?pr=482#step:5:4274")
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
	_, err := client.Alerting.Policies.New(context.TODO(), alerting.PolicyNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		AlertType: cloudflare.F(alerting.PolicyNewParamsAlertTypeAccessCustomCertificateExpirationType),
		Enabled:   cloudflare.F(true),
		Mechanisms: cloudflare.F(alerting.MechanismParam{
			"email": []alerting.MechanismItemParam{{
				ID: cloudflare.F("test@example.com"),
			}},
			"pagerduty": []alerting.MechanismItemParam{{
				ID: cloudflare.F("e8133a15-00a4-4d69-aec1-32f70c51f6e5"),
			}},
			"webhooks": []alerting.MechanismItemParam{{
				ID: cloudflare.F("14cc1190-5d2b-4b98-a696-c424cb2ad05f"),
			}},
		}),
		Name:          cloudflare.F("SSL Notification Event Policy"),
		AlertInterval: cloudflare.F("30m"),
		Description:   cloudflare.F("Something describing the policy."),
		Filters: cloudflare.F(alerting.PolicyFilterParam{
			Actions:                      cloudflare.F([]string{"string"}),
			AffectedASNs:                 cloudflare.F([]string{"string"}),
			AffectedComponents:           cloudflare.F([]string{"string"}),
			AffectedLocations:            cloudflare.F([]string{"string"}),
			AirportCode:                  cloudflare.F([]string{"string"}),
			AlertTriggerPreferences:      cloudflare.F([]string{"string"}),
			AlertTriggerPreferencesValue: cloudflare.F([]string{"string"}),
			Enabled:                      cloudflare.F([]string{"string"}),
			Environment:                  cloudflare.F([]string{"string"}),
			Event:                        cloudflare.F([]string{"string"}),
			EventSource:                  cloudflare.F([]string{"string"}),
			EventType:                    cloudflare.F([]string{"string"}),
			GroupBy:                      cloudflare.F([]string{"string"}),
			HealthCheckID:                cloudflare.F([]string{"string"}),
			IncidentImpact:               cloudflare.F([]alerting.PolicyFilterIncidentImpact{alerting.PolicyFilterIncidentImpactIncidentImpactNone}),
			InputID:                      cloudflare.F([]string{"string"}),
			Limit:                        cloudflare.F([]string{"string"}),
			LogoTag:                      cloudflare.F([]string{"string"}),
			MegabitsPerSecond:            cloudflare.F([]string{"string"}),
			NewHealth:                    cloudflare.F([]string{"string"}),
			NewStatus:                    cloudflare.F([]string{"string"}),
			PacketsPerSecond:             cloudflare.F([]string{"string"}),
			PoolID:                       cloudflare.F([]string{"string"}),
			POPName:                      cloudflare.F([]string{"string"}),
			Product:                      cloudflare.F([]string{"string"}),
			ProjectID:                    cloudflare.F([]string{"string"}),
			Protocol:                     cloudflare.F([]string{"string"}),
			QueryTag:                     cloudflare.F([]string{"string"}),
			RequestsPerSecond:            cloudflare.F([]string{"string"}),
			Selectors:                    cloudflare.F([]string{"string"}),
			Services:                     cloudflare.F([]string{"string"}),
			Slo:                          cloudflare.F([]string{"99.9"}),
			Status:                       cloudflare.F([]string{"string"}),
			TargetHostname:               cloudflare.F([]string{"string"}),
			TargetIP:                     cloudflare.F([]string{"string"}),
			TargetZoneName:               cloudflare.F([]string{"string"}),
			TrafficExclusions:            cloudflare.F([]alerting.PolicyFilterTrafficExclusion{alerting.PolicyFilterTrafficExclusionSecurityEvents}),
			TunnelID:                     cloudflare.F([]string{"string"}),
			TunnelName:                   cloudflare.F([]string{"string"}),
			Where:                        cloudflare.F([]string{"string"}),
			Zones:                        cloudflare.F([]string{"string"}),
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

func TestPolicyUpdateWithOptionalParams(t *testing.T) {
	t.Skip("prism errors - https://github.com/cloudflare/cloudflare-python/actions/runs/9327225061/job/25676826349?pr=482#step:5:4274")
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
	_, err := client.Alerting.Policies.Update(
		context.TODO(),
		"0da2b59e-f118-439d-8097-bdfb215203c9",
		alerting.PolicyUpdateParams{
			AccountID:     cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AlertInterval: cloudflare.F("30m"),
			AlertType:     cloudflare.F(alerting.PolicyUpdateParamsAlertTypeAccessCustomCertificateExpirationType),
			Description:   cloudflare.F("Something describing the policy."),
			Enabled:       cloudflare.F(true),
			Filters: cloudflare.F(alerting.PolicyFilterParam{
				Actions:                      cloudflare.F([]string{"string"}),
				AffectedASNs:                 cloudflare.F([]string{"string"}),
				AffectedComponents:           cloudflare.F([]string{"string"}),
				AffectedLocations:            cloudflare.F([]string{"string"}),
				AirportCode:                  cloudflare.F([]string{"string"}),
				AlertTriggerPreferences:      cloudflare.F([]string{"string"}),
				AlertTriggerPreferencesValue: cloudflare.F([]string{"string"}),
				Enabled:                      cloudflare.F([]string{"string"}),
				Environment:                  cloudflare.F([]string{"string"}),
				Event:                        cloudflare.F([]string{"string"}),
				EventSource:                  cloudflare.F([]string{"string"}),
				EventType:                    cloudflare.F([]string{"string"}),
				GroupBy:                      cloudflare.F([]string{"string"}),
				HealthCheckID:                cloudflare.F([]string{"string"}),
				IncidentImpact:               cloudflare.F([]alerting.PolicyFilterIncidentImpact{alerting.PolicyFilterIncidentImpactIncidentImpactNone}),
				InputID:                      cloudflare.F([]string{"string"}),
				Limit:                        cloudflare.F([]string{"string"}),
				LogoTag:                      cloudflare.F([]string{"string"}),
				MegabitsPerSecond:            cloudflare.F([]string{"string"}),
				NewHealth:                    cloudflare.F([]string{"string"}),
				NewStatus:                    cloudflare.F([]string{"string"}),
				PacketsPerSecond:             cloudflare.F([]string{"string"}),
				PoolID:                       cloudflare.F([]string{"string"}),
				POPName:                      cloudflare.F([]string{"string"}),
				Product:                      cloudflare.F([]string{"string"}),
				ProjectID:                    cloudflare.F([]string{"string"}),
				Protocol:                     cloudflare.F([]string{"string"}),
				QueryTag:                     cloudflare.F([]string{"string"}),
				RequestsPerSecond:            cloudflare.F([]string{"string"}),
				Selectors:                    cloudflare.F([]string{"string"}),
				Services:                     cloudflare.F([]string{"string"}),
				Slo:                          cloudflare.F([]string{"99.9"}),
				Status:                       cloudflare.F([]string{"string"}),
				TargetHostname:               cloudflare.F([]string{"string"}),
				TargetIP:                     cloudflare.F([]string{"string"}),
				TargetZoneName:               cloudflare.F([]string{"string"}),
				TrafficExclusions:            cloudflare.F([]alerting.PolicyFilterTrafficExclusion{alerting.PolicyFilterTrafficExclusionSecurityEvents}),
				TunnelID:                     cloudflare.F([]string{"string"}),
				TunnelName:                   cloudflare.F([]string{"string"}),
				Where:                        cloudflare.F([]string{"string"}),
				Zones:                        cloudflare.F([]string{"string"}),
			}),
			Mechanisms: cloudflare.F(alerting.MechanismParam{
				"email": []alerting.MechanismItemParam{{
					ID: cloudflare.F("test@example.com"),
				}},
				"pagerduty": []alerting.MechanismItemParam{{
					ID: cloudflare.F("e8133a15-00a4-4d69-aec1-32f70c51f6e5"),
				}},
				"webhooks": []alerting.MechanismItemParam{{
					ID: cloudflare.F("14cc1190-5d2b-4b98-a696-c424cb2ad05f"),
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

func TestPolicyList(t *testing.T) {
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
	_, err := client.Alerting.Policies.List(context.TODO(), alerting.PolicyListParams{
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

func TestPolicyDelete(t *testing.T) {
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
	_, err := client.Alerting.Policies.Delete(
		context.TODO(),
		"0da2b59e-f118-439d-8097-bdfb215203c9",
		alerting.PolicyDeleteParams{
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

func TestPolicyGet(t *testing.T) {
	t.Skip("prism errors - https://github.com/cloudflare/cloudflare-python/actions/runs/9327225061/job/25676826349?pr=482#step:5:4274")
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
	_, err := client.Alerting.Policies.Get(
		context.TODO(),
		"0da2b59e-f118-439d-8097-bdfb215203c9",
		alerting.PolicyGetParams{
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
