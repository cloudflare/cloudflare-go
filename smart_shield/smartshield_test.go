// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package smart_shield_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
	"github.com/cloudflare/cloudflare-go/v6/smart_shield"
)

func TestSmartShieldUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.SmartShield.Update(context.TODO(), smart_shield.SmartShieldUpdateParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		CacheReserve: cloudflare.F(smart_shield.SmartShieldUpdateParamsCacheReserve{
			Value: cloudflare.F(smart_shield.SmartShieldUpdateParamsCacheReserveValueOn),
		}),
		RegionalTieredCache: cloudflare.F(smart_shield.SmartShieldUpdateParamsRegionalTieredCache{
			Value: cloudflare.F(smart_shield.SmartShieldUpdateParamsRegionalTieredCacheValueOn),
		}),
		SmartRouting: cloudflare.F(smart_shield.SmartShieldUpdateParamsSmartRouting{
			Value: cloudflare.F(smart_shield.SmartShieldUpdateParamsSmartRoutingValueOn),
		}),
		SmartTieredCache: cloudflare.F(smart_shield.SmartShieldUpdateParamsSmartTieredCache{
			Value: cloudflare.F(smart_shield.SmartShieldUpdateParamsSmartTieredCacheValueOn),
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

func TestSmartShieldNewHealthcheckWithOptionalParams(t *testing.T) {
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
	_, err := client.SmartShield.NewHealthcheck(context.TODO(), smart_shield.SmartShieldNewHealthcheckParams{
		ZoneID:               cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Address:              cloudflare.F("www.example.com"),
		Name:                 cloudflare.F("server-1"),
		CheckRegions:         cloudflare.F([]smart_shield.SmartShieldNewHealthcheckParamsCheckRegion{smart_shield.SmartShieldNewHealthcheckParamsCheckRegionWeu, smart_shield.SmartShieldNewHealthcheckParamsCheckRegionEnam}),
		ConsecutiveFails:     cloudflare.F(int64(0)),
		ConsecutiveSuccesses: cloudflare.F(int64(0)),
		Description:          cloudflare.F("Health check for www.example.com"),
		HTTPConfig: cloudflare.F(smart_shield.SmartShieldNewHealthcheckParamsHTTPConfig{
			AllowInsecure:   cloudflare.F(true),
			ExpectedBody:    cloudflare.F("success"),
			ExpectedCodes:   cloudflare.F([]string{"2xx", "302"}),
			FollowRedirects: cloudflare.F(true),
			Header: cloudflare.F(map[string][]string{
				"Host":     {"example.com"},
				"X-App-ID": {"abc123"},
			}),
			Method: cloudflare.F(smart_shield.SmartShieldNewHealthcheckParamsHTTPConfigMethodGet),
			Path:   cloudflare.F("/health"),
			Port:   cloudflare.F(int64(0)),
		}),
		Interval:  cloudflare.F(int64(0)),
		Retries:   cloudflare.F(int64(0)),
		Suspended: cloudflare.F(true),
		TCPConfig: cloudflare.F(smart_shield.SmartShieldNewHealthcheckParamsTCPConfig{
			Method: cloudflare.F(smart_shield.SmartShieldNewHealthcheckParamsTCPConfigMethodConnectionEstablished),
			Port:   cloudflare.F(int64(0)),
		}),
		Timeout: cloudflare.F(int64(0)),
		Type:    cloudflare.F("HTTPS"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSmartShieldDeleteHealthcheck(t *testing.T) {
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
	_, err := client.SmartShield.DeleteHealthcheck(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		smart_shield.SmartShieldDeleteHealthcheckParams{
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

func TestSmartShieldEditHealthcheckWithOptionalParams(t *testing.T) {
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
	_, err := client.SmartShield.EditHealthcheck(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		smart_shield.SmartShieldEditHealthcheckParams{
			ZoneID:               cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Address:              cloudflare.F("www.example.com"),
			Name:                 cloudflare.F("server-1"),
			CheckRegions:         cloudflare.F([]smart_shield.SmartShieldEditHealthcheckParamsCheckRegion{smart_shield.SmartShieldEditHealthcheckParamsCheckRegionWeu, smart_shield.SmartShieldEditHealthcheckParamsCheckRegionEnam}),
			ConsecutiveFails:     cloudflare.F(int64(0)),
			ConsecutiveSuccesses: cloudflare.F(int64(0)),
			Description:          cloudflare.F("Health check for www.example.com"),
			HTTPConfig: cloudflare.F(smart_shield.SmartShieldEditHealthcheckParamsHTTPConfig{
				AllowInsecure:   cloudflare.F(true),
				ExpectedBody:    cloudflare.F("success"),
				ExpectedCodes:   cloudflare.F([]string{"2xx", "302"}),
				FollowRedirects: cloudflare.F(true),
				Header: cloudflare.F(map[string][]string{
					"Host":     {"example.com"},
					"X-App-ID": {"abc123"},
				}),
				Method: cloudflare.F(smart_shield.SmartShieldEditHealthcheckParamsHTTPConfigMethodGet),
				Path:   cloudflare.F("/health"),
				Port:   cloudflare.F(int64(0)),
			}),
			Interval:  cloudflare.F(int64(0)),
			Retries:   cloudflare.F(int64(0)),
			Suspended: cloudflare.F(true),
			TCPConfig: cloudflare.F(smart_shield.SmartShieldEditHealthcheckParamsTCPConfig{
				Method: cloudflare.F(smart_shield.SmartShieldEditHealthcheckParamsTCPConfigMethodConnectionEstablished),
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

func TestSmartShieldGet(t *testing.T) {
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
	_, err := client.SmartShield.Get(context.TODO(), smart_shield.SmartShieldGetParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSmartShieldGetHealthcheck(t *testing.T) {
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
	_, err := client.SmartShield.GetHealthcheck(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		smart_shield.SmartShieldGetHealthcheckParams{
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

func TestSmartShieldListHealthchecksWithOptionalParams(t *testing.T) {
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
	_, err := client.SmartShield.ListHealthchecks(context.TODO(), smart_shield.SmartShieldListHealthchecksParams{
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

func TestSmartShieldUpdateHealthcheckWithOptionalParams(t *testing.T) {
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
	_, err := client.SmartShield.UpdateHealthcheck(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		smart_shield.SmartShieldUpdateHealthcheckParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Errors: cloudflare.F([]shared.ResponseInfoParam{{
				Code:             cloudflare.F(int64(1000)),
				Message:          cloudflare.F("message"),
				DocumentationURL: cloudflare.F("documentation_url"),
				Source: cloudflare.F(shared.ResponseInfoSourceParam{
					Pointer: cloudflare.F("pointer"),
				}),
			}}),
			Messages: cloudflare.F([]shared.ResponseInfoParam{{
				Code:             cloudflare.F(int64(1000)),
				Message:          cloudflare.F("message"),
				DocumentationURL: cloudflare.F("documentation_url"),
				Source: cloudflare.F(shared.ResponseInfoSourceParam{
					Pointer: cloudflare.F("pointer"),
				}),
			}}),
			Result: cloudflare.F(smart_shield.SmartShieldUpdateHealthcheckParamsResult{
				Address:              cloudflare.F("www.example.com"),
				CheckRegions:         cloudflare.F([]smart_shield.SmartShieldUpdateHealthcheckParamsResultCheckRegion{smart_shield.SmartShieldUpdateHealthcheckParamsResultCheckRegionWeu, smart_shield.SmartShieldUpdateHealthcheckParamsResultCheckRegionEnam}),
				ConsecutiveFails:     cloudflare.F(int64(0)),
				ConsecutiveSuccesses: cloudflare.F(int64(0)),
				Description:          cloudflare.F("Health check for www.example.com"),
				HTTPConfig: cloudflare.F(smart_shield.SmartShieldUpdateHealthcheckParamsResultHTTPConfig{
					AllowInsecure:   cloudflare.F(true),
					ExpectedBody:    cloudflare.F("success"),
					ExpectedCodes:   cloudflare.F([]string{"2xx", "302"}),
					FollowRedirects: cloudflare.F(true),
					Header: cloudflare.F(map[string][]string{
						"Host":     {"example.com"},
						"X-App-ID": {"abc123"},
					}),
					Method: cloudflare.F(smart_shield.SmartShieldUpdateHealthcheckParamsResultHTTPConfigMethodGet),
					Path:   cloudflare.F("/health"),
					Port:   cloudflare.F(int64(0)),
				}),
				Interval:  cloudflare.F(int64(0)),
				Name:      cloudflare.F("server-1"),
				Retries:   cloudflare.F(int64(0)),
				Suspended: cloudflare.F(true),
				TCPConfig: cloudflare.F(smart_shield.SmartShieldUpdateHealthcheckParamsResultTCPConfig{
					Method: cloudflare.F(smart_shield.SmartShieldUpdateHealthcheckParamsResultTCPConfigMethodConnectionEstablished),
					Port:   cloudflare.F(int64(0)),
				}),
				Timeout: cloudflare.F(int64(0)),
				Type:    cloudflare.F("HTTPS"),
			}),
			Success: cloudflare.F(smart_shield.SmartShieldUpdateHealthcheckParamsSuccessTrue),
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
