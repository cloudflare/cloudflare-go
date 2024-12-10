// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limits_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/rate_limits"
)

func TestRateLimitNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.RateLimits.New(context.TODO(), rate_limits.RateLimitNewParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Action: cloudflare.F(rate_limits.RateLimitNewParamsAction{
			Mode: cloudflare.F(rate_limits.RateLimitNewParamsActionModeSimulate),
			Response: cloudflare.F(rate_limits.RateLimitNewParamsActionResponse{
				Body:        cloudflare.F("<error>This request has been rate-limited.</error>"),
				ContentType: cloudflare.F("text/xml"),
			}),
			Timeout: cloudflare.F(86400.000000),
		}),
		Match: cloudflare.F(rate_limits.RateLimitNewParamsMatch{
			Headers: cloudflare.F([]rate_limits.RateLimitNewParamsMatchHeader{{
				Name:  cloudflare.F("Cf-Cache-Status"),
				Op:    cloudflare.F(rate_limits.RateLimitNewParamsMatchHeadersOpEq),
				Value: cloudflare.F("HIT"),
			}}),
			Request: cloudflare.F(rate_limits.RateLimitNewParamsMatchRequest{
				Methods: cloudflare.F([]rate_limits.Methods{rate_limits.MethodsGet, rate_limits.MethodsPost}),
				Schemes: cloudflare.F([]string{"HTTP", "HTTPS"}),
				URL:     cloudflare.F("*.example.org/path*"),
			}),
			Response: cloudflare.F(rate_limits.RateLimitNewParamsMatchResponse{
				OriginTraffic: cloudflare.F(true),
			}),
		}),
		Period:    cloudflare.F(900.000000),
		Threshold: cloudflare.F(60.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRateLimitListWithOptionalParams(t *testing.T) {
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
	_, err := client.RateLimits.List(context.TODO(), rate_limits.RateLimitListParams{
		ZoneID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Page:    cloudflare.F(1.000000),
		PerPage: cloudflare.F(1.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRateLimitDelete(t *testing.T) {
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
	_, err := client.RateLimits.Delete(
		context.TODO(),
		"372e67954025e0ba6aaa6d586b9e0b59",
		rate_limits.RateLimitDeleteParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestRateLimitEditWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.RateLimits.Edit(
		context.TODO(),
		"372e67954025e0ba6aaa6d586b9e0b59",
		rate_limits.RateLimitEditParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Action: cloudflare.F(rate_limits.RateLimitEditParamsAction{
				Mode: cloudflare.F(rate_limits.RateLimitEditParamsActionModeSimulate),
				Response: cloudflare.F(rate_limits.RateLimitEditParamsActionResponse{
					Body:        cloudflare.F("<error>This request has been rate-limited.</error>"),
					ContentType: cloudflare.F("text/xml"),
				}),
				Timeout: cloudflare.F(86400.000000),
			}),
			Match: cloudflare.F(rate_limits.RateLimitEditParamsMatch{
				Headers: cloudflare.F([]rate_limits.RateLimitEditParamsMatchHeader{{
					Name:  cloudflare.F("Cf-Cache-Status"),
					Op:    cloudflare.F(rate_limits.RateLimitEditParamsMatchHeadersOpEq),
					Value: cloudflare.F("HIT"),
				}}),
				Request: cloudflare.F(rate_limits.RateLimitEditParamsMatchRequest{
					Methods: cloudflare.F([]rate_limits.Methods{rate_limits.MethodsGet, rate_limits.MethodsPost}),
					Schemes: cloudflare.F([]string{"HTTP", "HTTPS"}),
					URL:     cloudflare.F("*.example.org/path*"),
				}),
				Response: cloudflare.F(rate_limits.RateLimitEditParamsMatchResponse{
					OriginTraffic: cloudflare.F(true),
				}),
			}),
			Period:    cloudflare.F(900.000000),
			Threshold: cloudflare.F(60.000000),
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

func TestRateLimitGet(t *testing.T) {
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
	_, err := client.RateLimits.Get(
		context.TODO(),
		"372e67954025e0ba6aaa6d586b9e0b59",
		rate_limits.RateLimitGetParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
