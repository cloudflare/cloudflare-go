// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/cloudforce_one"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

func TestThreatEventNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism")
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
	_, err := client.CloudforceOne.ThreatEvents.New(context.TODO(), cloudforce_one.ThreatEventNewParams{
		PathAccountID: cloudflare.F("account_id"),
		Category:      cloudflare.F("Domain Resolution"),
		Date:          cloudflare.F(time.Now()),
		Event:         cloudflare.F("An attacker registered the domain domain.com"),
		Raw: cloudflare.F(cloudforce_one.ThreatEventNewParamsRaw{
			Data: cloudflare.F(map[string]interface{}{
				"foo": "bar",
			}),
			Source: cloudflare.F("example.com"),
			TLP:    cloudflare.F("amber"),
		}),
		TLP:             cloudflare.F("amber"),
		BodyAccountID:   cloudflare.F(123456.000000),
		Attacker:        cloudflare.F("Flying Yeti"),
		AttackerCountry: cloudflare.F("CN"),
		DatasetID:       cloudflare.F("durableObjectName"),
		Indicator:       cloudflare.F("domain.com"),
		Indicators: cloudflare.F([]cloudforce_one.ThreatEventNewParamsIndicator{{
			IndicatorType: cloudflare.F("domain"),
			Value:         cloudflare.F("malicious.com"),
		}}),
		IndicatorType:  cloudflare.F("domain"),
		Insight:        cloudflare.F("This domain was likely registered for phishing purposes"),
		Tags:           cloudflare.F([]string{"malware"}),
		TargetCountry:  cloudflare.F("US"),
		TargetIndustry: cloudflare.F("Agriculture"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestThreatEventListWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism")
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
	_, err := client.CloudforceOne.ThreatEvents.List(context.TODO(), cloudforce_one.ThreatEventListParams{
		AccountID:    cloudflare.F("account_id"),
		Cursor:       cloudflare.F("eyJ2ZXJzaW9uIjoxLCJwb3NpdGlvbiI6eyJkYXRlIjoiMjAyNC0wMS0xMlQxMDowMDowMFoiLCJ1dWlkIjoiYWJjMTIzIn19"),
		DatasetID:    cloudflare.F([]string{"string"}),
		ForceRefresh: cloudflare.F(true),
		Format:       cloudflare.F(cloudforce_one.ThreatEventListParamsFormatJson),
		Order:        cloudflare.F(cloudforce_one.ThreatEventListParamsOrderAsc),
		OrderBy:      cloudflare.F("orderBy"),
		Page:         cloudflare.F(0.000000),
		PageSize:     cloudflare.F(0.000000),
		Search: cloudflare.F([]cloudforce_one.ThreatEventListParamsSearch{{
			Field: cloudflare.F("attackerCountry"),
			Op:    cloudflare.F(cloudforce_one.ThreatEventListParamsSearchOpEquals),
			Value: cloudflare.F[cloudforce_one.ThreatEventListParamsSearchValueUnion](shared.UnionString("usa")),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestThreatEventDelete(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism")
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
	_, err := client.CloudforceOne.ThreatEvents.Delete(
		context.TODO(),
		"event_id",
		cloudforce_one.ThreatEventDeleteParams{
			AccountID: cloudflare.F("account_id"),
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

func TestThreatEventBulkNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism")
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
	_, err := client.CloudforceOne.ThreatEvents.BulkNew(context.TODO(), cloudforce_one.ThreatEventBulkNewParams{
		AccountID: cloudflare.F("account_id"),
		Data: cloudflare.F([]cloudforce_one.ThreatEventBulkNewParamsData{{
			Category: cloudflare.F("Domain Resolution"),
			Date:     cloudflare.F(time.Now()),
			Event:    cloudflare.F("An attacker registered the domain domain.com"),
			Raw: cloudflare.F(cloudforce_one.ThreatEventBulkNewParamsDataRaw{
				Data: cloudflare.F(map[string]interface{}{
					"foo": "bar",
				}),
				Source: cloudflare.F("example.com"),
				TLP:    cloudflare.F("amber"),
			}),
			TLP:             cloudflare.F("amber"),
			AccountID:       cloudflare.F(123456.000000),
			Attacker:        cloudflare.F("Flying Yeti"),
			AttackerCountry: cloudflare.F("CN"),
			DatasetID:       cloudflare.F("durableObjectName"),
			Indicator:       cloudflare.F("domain.com"),
			Indicators: cloudflare.F([]cloudforce_one.ThreatEventBulkNewParamsDataIndicator{{
				IndicatorType: cloudflare.F("domain"),
				Value:         cloudflare.F("malicious.com"),
			}}),
			IndicatorType:  cloudflare.F("domain"),
			Insight:        cloudflare.F("This domain was likely registered for phishing purposes"),
			Tags:           cloudflare.F([]string{"malware"}),
			TargetCountry:  cloudflare.F("US"),
			TargetIndustry: cloudflare.F("Agriculture"),
		}}),
		DatasetID:            cloudflare.F("durableObjectName"),
		IncludeCreatedEvents: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestThreatEventEditWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism")
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
	_, err := client.CloudforceOne.ThreatEvents.Edit(
		context.TODO(),
		"event_id",
		cloudforce_one.ThreatEventEditParams{
			AccountID:       cloudflare.F("account_id"),
			Attacker:        cloudflare.F("Flying Yeti"),
			AttackerCountry: cloudflare.F("CN"),
			Category:        cloudflare.F("Domain Resolution"),
			CreatedAt:       cloudflare.F(time.Now()),
			DatasetID:       cloudflare.F("9b769969-a211-466c-8ac3-cb91266a066a"),
			Date:            cloudflare.F(time.Now()),
			Event:           cloudflare.F("An attacker registered the domain domain.com"),
			Indicator:       cloudflare.F("domain2.com"),
			IndicatorType:   cloudflare.F("domain"),
			Insight:         cloudflare.F("new insight"),
			Raw: cloudflare.F(cloudforce_one.ThreatEventEditParamsRaw{
				Data: cloudflare.F(map[string]interface{}{
					"foo": "bar",
				}),
				Source: cloudflare.F("example.com"),
				TLP:    cloudflare.F("amber"),
			}),
			TargetCountry:  cloudflare.F("US"),
			TargetIndustry: cloudflare.F("Insurance"),
			TLP:            cloudflare.F("amber"),
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

func TestThreatEventGet(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism")
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
	_, err := client.CloudforceOne.ThreatEvents.Get(
		context.TODO(),
		"event_id",
		cloudforce_one.ThreatEventGetParams{
			AccountID: cloudflare.F("account_id"),
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
