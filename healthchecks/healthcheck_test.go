// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthchecks_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/healthchecks"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestHealthcheckNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Healthchecks.New(context.TODO(), healthchecks.HealthcheckNewParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		QueryHealthcheck: healthchecks.QueryHealthcheckParam{
			Address:              cloudflare.F("www.example.com"),
			Name:                 cloudflare.F("server-1"),
			CheckRegions:         cloudflare.F([]healthchecks.CheckRegion{healthchecks.CheckRegionWnam, healthchecks.CheckRegionEnam}),
			ConsecutiveFails:     cloudflare.F(int64(0)),
			ConsecutiveSuccesses: cloudflare.F(int64(0)),
			Description:          cloudflare.F("Health check for www.example.com"),
			HTTPConfig: cloudflare.F(healthchecks.HTTPConfigurationParam{
				AllowInsecure:   cloudflare.F(true),
				ExpectedBody:    cloudflare.F("success"),
				ExpectedCodes:   cloudflare.F([]string{"2xx", "302"}),
				FollowRedirects: cloudflare.F(true),
				Header: cloudflare.F(map[string][]string{
					"Host":     {"example.com"},
					"X-App-ID": {"abc123"},
				}),
				Method: cloudflare.F(healthchecks.HTTPConfigurationMethodGet),
				Path:   cloudflare.F("/health"),
				Port:   cloudflare.F(int64(0)),
			}),
			Interval:  cloudflare.F(int64(0)),
			Retries:   cloudflare.F(int64(0)),
			Suspended: cloudflare.F(true),
			TCPConfig: cloudflare.F(healthchecks.TCPConfigurationParam{
				Method: cloudflare.F(healthchecks.TCPConfigurationMethodConnectionEstablished),
				Port:   cloudflare.F(int64(0)),
			}),
			Timeout: cloudflare.F(int64(0)),
			Type:    cloudflare.F("HTTPS"),
		},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHealthcheckUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Healthchecks.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		healthchecks.HealthcheckUpdateParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			QueryHealthcheck: healthchecks.QueryHealthcheckParam{
				Address:              cloudflare.F("www.example.com"),
				Name:                 cloudflare.F("server-1"),
				CheckRegions:         cloudflare.F([]healthchecks.CheckRegion{healthchecks.CheckRegionWnam, healthchecks.CheckRegionEnam}),
				ConsecutiveFails:     cloudflare.F(int64(0)),
				ConsecutiveSuccesses: cloudflare.F(int64(0)),
				Description:          cloudflare.F("Health check for www.example.com"),
				HTTPConfig: cloudflare.F(healthchecks.HTTPConfigurationParam{
					AllowInsecure:   cloudflare.F(true),
					ExpectedBody:    cloudflare.F("success"),
					ExpectedCodes:   cloudflare.F([]string{"2xx", "302"}),
					FollowRedirects: cloudflare.F(true),
					Header: cloudflare.F(map[string][]string{
						"Host":     {"example.com"},
						"X-App-ID": {"abc123"},
					}),
					Method: cloudflare.F(healthchecks.HTTPConfigurationMethodGet),
					Path:   cloudflare.F("/health"),
					Port:   cloudflare.F(int64(0)),
				}),
				Interval:  cloudflare.F(int64(0)),
				Retries:   cloudflare.F(int64(0)),
				Suspended: cloudflare.F(true),
				TCPConfig: cloudflare.F(healthchecks.TCPConfigurationParam{
					Method: cloudflare.F(healthchecks.TCPConfigurationMethodConnectionEstablished),
					Port:   cloudflare.F(int64(0)),
				}),
				Timeout: cloudflare.F(int64(0)),
				Type:    cloudflare.F("HTTPS"),
			},
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

func TestHealthcheckListWithOptionalParams(t *testing.T) {
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
	_, err := client.Healthchecks.List(context.TODO(), healthchecks.HealthcheckListParams{
		ZoneID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Page:    cloudflare.F(1.000000),
		PerPage: cloudflare.F(5.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHealthcheckDelete(t *testing.T) {
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
	_, err := client.Healthchecks.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		healthchecks.HealthcheckDeleteParams{
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

func TestHealthcheckEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Healthchecks.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		healthchecks.HealthcheckEditParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			QueryHealthcheck: healthchecks.QueryHealthcheckParam{
				Address:              cloudflare.F("www.example.com"),
				Name:                 cloudflare.F("server-1"),
				CheckRegions:         cloudflare.F([]healthchecks.CheckRegion{healthchecks.CheckRegionWnam, healthchecks.CheckRegionEnam}),
				ConsecutiveFails:     cloudflare.F(int64(0)),
				ConsecutiveSuccesses: cloudflare.F(int64(0)),
				Description:          cloudflare.F("Health check for www.example.com"),
				HTTPConfig: cloudflare.F(healthchecks.HTTPConfigurationParam{
					AllowInsecure:   cloudflare.F(true),
					ExpectedBody:    cloudflare.F("success"),
					ExpectedCodes:   cloudflare.F([]string{"2xx", "302"}),
					FollowRedirects: cloudflare.F(true),
					Header: cloudflare.F(map[string][]string{
						"Host":     {"example.com"},
						"X-App-ID": {"abc123"},
					}),
					Method: cloudflare.F(healthchecks.HTTPConfigurationMethodGet),
					Path:   cloudflare.F("/health"),
					Port:   cloudflare.F(int64(0)),
				}),
				Interval:  cloudflare.F(int64(0)),
				Retries:   cloudflare.F(int64(0)),
				Suspended: cloudflare.F(true),
				TCPConfig: cloudflare.F(healthchecks.TCPConfigurationParam{
					Method: cloudflare.F(healthchecks.TCPConfigurationMethodConnectionEstablished),
					Port:   cloudflare.F(int64(0)),
				}),
				Timeout: cloudflare.F(int64(0)),
				Type:    cloudflare.F("HTTPS"),
			},
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

func TestHealthcheckGet(t *testing.T) {
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
	_, err := client.Healthchecks.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		healthchecks.HealthcheckGetParams{
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
