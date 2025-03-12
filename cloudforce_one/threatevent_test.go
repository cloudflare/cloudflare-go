// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/cloudforce_one"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

func TestThreatEventNewWithOptionalParams(t *testing.T) {
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
		AccountID: cloudflare.F(0.000000),
		DatasetID: cloudflare.F([]string{"7632a037-fdef-4899-9b12-148470aae772"}),
		Order:     cloudflare.F(cloudforce_one.ThreatEventNewParamsOrderAsc),
		OrderBy:   cloudflare.F("created"),
		Page:      cloudflare.F(1.000000),
		PageSize:  cloudflare.F(100.000000),
		Search: cloudflare.F([]cloudforce_one.ThreatEventNewParamsSearch{{
			Field: cloudflare.F("attackerCountry"),
			Op:    cloudflare.F("equals"),
			Value: cloudflare.F[cloudforce_one.ThreatEventNewParamsSearchValueUnion](shared.UnionString("usa")),
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
			AccountID: cloudflare.F(0.000000),
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

func TestThreatEventBulkNew(t *testing.T) {
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
		AccountID: cloudflare.F(0.000000),
		Data: cloudflare.F([]cloudforce_one.ThreatEventBulkNewParamsData{{
			Attacker:        cloudflare.F("Flying Yeti"),
			AttackerCountry: cloudflare.F("CN"),
			Category:        cloudflare.F("Domain Resolution"),
			Date:            cloudflare.F(time.Now()),
			Event:           cloudflare.F("An attacker registered the domain domain.com"),
			IndicatorType:   cloudflare.F("domain"),
			Raw: cloudflare.F(cloudforce_one.ThreatEventBulkNewParamsDataRaw{
				Data:   cloudflare.F[any](map[string]interface{}{}),
				Source: cloudflare.F("example.com"),
				TLP:    cloudflare.F("amber"),
			}),
			TLP:            cloudflare.F("amber"),
			AccountID:      cloudflare.F(123456.000000),
			DatasetID:      cloudflare.F("durableObjectName"),
			Indicator:      cloudflare.F("domain.com"),
			Tags:           cloudflare.F([]string{"malware"}),
			TargetCountry:  cloudflare.F("US"),
			TargetIndustry: cloudflare.F("Agriculture"),
		}}),
		DatasetID: cloudflare.F("durableObjectName"),
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
			AccountID:       cloudflare.F(0.000000),
			Attacker:        cloudflare.F("Flying Yeti"),
			AttackerCountry: cloudflare.F("CN"),
			Category:        cloudflare.F("Domain Resolution"),
			Date:            cloudflare.F(time.Now()),
			Event:           cloudflare.F("An attacker registered the domain domain.com"),
			Indicator:       cloudflare.F("domain2.com"),
			IndicatorType:   cloudflare.F("sha256"),
			TargetCountry:   cloudflare.F("US"),
			TargetIndustry:  cloudflare.F("Insurance"),
			TLP:             cloudflare.F("amber"),
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
			AccountID: cloudflare.F(0.000000),
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
