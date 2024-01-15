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

func TestZoneHealthcheckPreviewGet(t *testing.T) {
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
	_, err := client.Zones.Healthchecks.Previews.Get(
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

func TestZoneHealthcheckPreviewDelete(t *testing.T) {
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
	_, err := client.Zones.Healthchecks.Previews.Delete(
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

func TestZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Healthchecks.Previews.HealthChecksNewPreviewHealthCheck(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParams{
			Address:              cloudflare.F("www.example.com"),
			Name:                 cloudflare.F("server-1"),
			CheckRegions:         cloudflare.F([]cloudflare.ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion{cloudflare.ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionWeu, cloudflare.ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionEnam}),
			ConsecutiveFails:     cloudflare.F(int64(0)),
			ConsecutiveSuccesses: cloudflare.F(int64(0)),
			Description:          cloudflare.F("Health check for www.example.com"),
			HTTPConfig: cloudflare.F(cloudflare.ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsHTTPConfig{
				AllowInsecure:   cloudflare.F(true),
				ExpectedBody:    cloudflare.F("success"),
				ExpectedCodes:   cloudflare.F([]string{"2xx", "302"}),
				FollowRedirects: cloudflare.F(true),
				Header: cloudflare.F[any](map[string]interface{}{
					"Host": map[string]interface{}{
						"0": "example.com",
					},
					"X-App-ID": map[string]interface{}{
						"0": "abc123",
					},
				}),
				Method: cloudflare.F(cloudflare.ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsHTTPConfigMethodGet),
				Path:   cloudflare.F("/health"),
				Port:   cloudflare.F(int64(0)),
			}),
			Interval:  cloudflare.F(int64(0)),
			Retries:   cloudflare.F(int64(0)),
			Suspended: cloudflare.F(true),
			TcpConfig: cloudflare.F(cloudflare.ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsTcpConfig{
				Method: cloudflare.F(cloudflare.ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsTcpConfigMethodConnectionEstablished),
				Port:   cloudflare.F(int64(0)),
			}),
			Timeout: cloudflare.F(int64(0)),
			Type:    cloudflare.F("HTTPS"),
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
