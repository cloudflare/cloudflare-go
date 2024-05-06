// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/ai_gateway"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestAIGatewayNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.New(context.TODO(), ai_gateway.AIGatewayNewParams{
		AccountID:               cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
		CacheInvalidateOnUpdate: cloudflare.F(true),
		CacheTTL:                cloudflare.F(int64(0)),
		CollectLogs:             cloudflare.F(true),
		Name:                    cloudflare.F("string"),
		Slug:                    cloudflare.F("string"),
		RateLimitingInterval:    cloudflare.F(int64(0)),
		RateLimitingLimit:       cloudflare.F(int64(0)),
		RateLimitingTechnique:   cloudflare.F("string"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIGatewayUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		ai_gateway.AIGatewayUpdateParams{
			AccountID:               cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
			CacheInvalidateOnUpdate: cloudflare.F(true),
			CacheTTL:                cloudflare.F(int64(0)),
			CollectLogs:             cloudflare.F(true),
			Name:                    cloudflare.F("string"),
			Slug:                    cloudflare.F("string"),
			RateLimitingInterval:    cloudflare.F(int64(0)),
			RateLimitingLimit:       cloudflare.F(int64(0)),
			RateLimitingTechnique:   cloudflare.F("string"),
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

func TestAIGatewayListWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.List(context.TODO(), ai_gateway.AIGatewayListParams{
		AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
		ID:        cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		OrderBy:   cloudflare.F("string"),
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(5)),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIGatewayDelete(t *testing.T) {
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
	_, err := client.AIGateway.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		ai_gateway.AIGatewayDeleteParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
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

func TestAIGatewayGet(t *testing.T) {
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
	_, err := client.AIGateway.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		ai_gateway.AIGatewayGetParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
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
