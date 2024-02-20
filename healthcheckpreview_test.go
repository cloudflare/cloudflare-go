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

func TestHealthcheckPreviewNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Healthchecks.Previews.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.HealthcheckPreviewNewParams{
			Address:              cloudflare.F("www.example.com"),
			Name:                 cloudflare.F("server-1"),
			CheckRegions:         cloudflare.F([]cloudflare.HealthcheckPreviewNewParamsCheckRegion{cloudflare.HealthcheckPreviewNewParamsCheckRegionWeu, cloudflare.HealthcheckPreviewNewParamsCheckRegionEnam}),
			ConsecutiveFails:     cloudflare.F(int64(0)),
			ConsecutiveSuccesses: cloudflare.F(int64(0)),
			Description:          cloudflare.F("Health check for www.example.com"),
			HTTPConfig: cloudflare.F(cloudflare.HealthcheckPreviewNewParamsHTTPConfig{
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
				Method: cloudflare.F(cloudflare.HealthcheckPreviewNewParamsHTTPConfigMethodGet),
				Path:   cloudflare.F("/health"),
				Port:   cloudflare.F(int64(0)),
			}),
			Interval:  cloudflare.F(int64(0)),
			Retries:   cloudflare.F(int64(0)),
			Suspended: cloudflare.F(true),
			TcpConfig: cloudflare.F(cloudflare.HealthcheckPreviewNewParamsTcpConfig{
				Method: cloudflare.F(cloudflare.HealthcheckPreviewNewParamsTcpConfigMethodConnectionEstablished),
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

func TestHealthcheckPreviewDelete(t *testing.T) {
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
	_, err := client.Healthchecks.Previews.Delete(
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

func TestHealthcheckPreviewGet(t *testing.T) {
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
	_, err := client.Healthchecks.Previews.Get(
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
