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

func TestZonePageruleGet(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Pagerules.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZonePageruleUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Pagerules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZonePageruleUpdateParams{
			Actions: cloudflare.F([]cloudflare.ZonePageruleUpdateParamsAction{{
				Name: cloudflare.F(cloudflare.ZonePageruleUpdateParamsActionsNameForwardURL),
				Value: cloudflare.F(cloudflare.ZonePageruleUpdateParamsActionsValue{
					Type: cloudflare.F(cloudflare.ZonePageruleUpdateParamsActionsValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}, {
				Name: cloudflare.F(cloudflare.ZonePageruleUpdateParamsActionsNameForwardURL),
				Value: cloudflare.F(cloudflare.ZonePageruleUpdateParamsActionsValue{
					Type: cloudflare.F(cloudflare.ZonePageruleUpdateParamsActionsValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}, {
				Name: cloudflare.F(cloudflare.ZonePageruleUpdateParamsActionsNameForwardURL),
				Value: cloudflare.F(cloudflare.ZonePageruleUpdateParamsActionsValue{
					Type: cloudflare.F(cloudflare.ZonePageruleUpdateParamsActionsValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}}),
			Targets: cloudflare.F([]cloudflare.ZonePageruleUpdateParamsTarget{{
				Constraint: cloudflare.F(cloudflare.ZonePageruleUpdateParamsTargetsConstraint{
					Operator: cloudflare.F(cloudflare.ZonePageruleUpdateParamsTargetsConstraintOperatorMatches),
					Value:    cloudflare.F("*example.com/images/*"),
				}),
				Target: cloudflare.F(cloudflare.ZonePageruleUpdateParamsTargetsTargetURL),
			}}),
			Priority: cloudflare.F(int64(0)),
			Status:   cloudflare.F(cloudflare.ZonePageruleUpdateParamsStatusActive),
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

func TestZonePageruleDelete(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Pagerules.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZonePagerulePageRulesNewAPageRuleWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Pagerules.PageRulesNewAPageRule(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZonePagerulePageRulesNewAPageRuleParams{
			Actions: cloudflare.F([]cloudflare.ZonePagerulePageRulesNewAPageRuleParamsAction{{
				Name: cloudflare.F(cloudflare.ZonePagerulePageRulesNewAPageRuleParamsActionsNameForwardURL),
				Value: cloudflare.F(cloudflare.ZonePagerulePageRulesNewAPageRuleParamsActionsValue{
					Type: cloudflare.F(cloudflare.ZonePagerulePageRulesNewAPageRuleParamsActionsValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}, {
				Name: cloudflare.F(cloudflare.ZonePagerulePageRulesNewAPageRuleParamsActionsNameForwardURL),
				Value: cloudflare.F(cloudflare.ZonePagerulePageRulesNewAPageRuleParamsActionsValue{
					Type: cloudflare.F(cloudflare.ZonePagerulePageRulesNewAPageRuleParamsActionsValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}, {
				Name: cloudflare.F(cloudflare.ZonePagerulePageRulesNewAPageRuleParamsActionsNameForwardURL),
				Value: cloudflare.F(cloudflare.ZonePagerulePageRulesNewAPageRuleParamsActionsValue{
					Type: cloudflare.F(cloudflare.ZonePagerulePageRulesNewAPageRuleParamsActionsValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}}),
			Targets: cloudflare.F([]cloudflare.ZonePagerulePageRulesNewAPageRuleParamsTarget{{
				Constraint: cloudflare.F(cloudflare.ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraint{
					Operator: cloudflare.F(cloudflare.ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraintOperatorMatches),
					Value:    cloudflare.F("*example.com/images/*"),
				}),
				Target: cloudflare.F(cloudflare.ZonePagerulePageRulesNewAPageRuleParamsTargetsTargetURL),
			}}),
			Priority: cloudflare.F(int64(0)),
			Status:   cloudflare.F(cloudflare.ZonePagerulePageRulesNewAPageRuleParamsStatusActive),
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

func TestZonePagerulePageRulesListPageRulesWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Pagerules.PageRulesListPageRules(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZonePagerulePageRulesListPageRulesParams{
			Direction: cloudflare.F(cloudflare.ZonePagerulePageRulesListPageRulesParamsDirectionDesc),
			Match:     cloudflare.F(cloudflare.ZonePagerulePageRulesListPageRulesParamsMatchAny),
			Order:     cloudflare.F(cloudflare.ZonePagerulePageRulesListPageRulesParamsOrderStatus),
			Status:    cloudflare.F(cloudflare.ZonePagerulePageRulesListPageRulesParamsStatusActive),
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
