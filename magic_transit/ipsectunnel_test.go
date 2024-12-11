// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/magic_transit"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestIPSECTunnelNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.IPSECTunnels.New(context.TODO(), magic_transit.IPSECTunnelNewParams{
		AccountID:          cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		CloudflareEndpoint: cloudflare.F("203.0.113.1"),
		InterfaceAddress:   cloudflare.F("192.0.2.0/31"),
		Name:               cloudflare.F("IPsec_1"),
		CustomerEndpoint:   cloudflare.F("203.0.113.1"),
		Description:        cloudflare.F("Tunnel for ISP X"),
		HealthCheck: cloudflare.F(magic_transit.IPSECTunnelNewParamsHealthCheck{
			Direction: cloudflare.F(magic_transit.IPSECTunnelNewParamsHealthCheckDirectionUnidirectional),
			Enabled:   cloudflare.F(true),
			Rate:      cloudflare.F(magic_transit.HealthCheckRateLow),
			Target: cloudflare.F[magic_transit.IPSECTunnelNewParamsHealthCheckTargetUnion](magic_transit.IPSECTunnelNewParamsHealthCheckTargetMagicHealthCheckTarget{
				Saved: cloudflare.F("203.0.113.1"),
			}),
			Type: cloudflare.F(magic_transit.HealthCheckTypeReply),
		}),
		PSK:               cloudflare.F("O3bwKSjnaoCxDoUxjcq4Rk8ZKkezQUiy"),
		ReplayProtection:  cloudflare.F(false),
		XMagicNewHcTarget: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIPSECTunnelUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.IPSECTunnels.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.IPSECTunnelUpdateParams{
			AccountID:          cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			CloudflareEndpoint: cloudflare.F("203.0.113.1"),
			InterfaceAddress:   cloudflare.F("192.0.2.0/31"),
			Name:               cloudflare.F("IPsec_1"),
			CustomerEndpoint:   cloudflare.F("203.0.113.1"),
			Description:        cloudflare.F("Tunnel for ISP X"),
			HealthCheck: cloudflare.F(magic_transit.IPSECTunnelUpdateParamsHealthCheck{
				Direction: cloudflare.F(magic_transit.IPSECTunnelUpdateParamsHealthCheckDirectionUnidirectional),
				Enabled:   cloudflare.F(true),
				Rate:      cloudflare.F(magic_transit.HealthCheckRateLow),
				Target: cloudflare.F[magic_transit.IPSECTunnelUpdateParamsHealthCheckTargetUnion](magic_transit.IPSECTunnelUpdateParamsHealthCheckTargetMagicHealthCheckTarget{
					Saved: cloudflare.F("203.0.113.1"),
				}),
				Type: cloudflare.F(magic_transit.HealthCheckTypeReply),
			}),
			PSK:               cloudflare.F("O3bwKSjnaoCxDoUxjcq4Rk8ZKkezQUiy"),
			ReplayProtection:  cloudflare.F(false),
			XMagicNewHcTarget: cloudflare.F(true),
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

func TestIPSECTunnelListWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.IPSECTunnels.List(context.TODO(), magic_transit.IPSECTunnelListParams{
		AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		XMagicNewHcTarget: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIPSECTunnelDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.IPSECTunnels.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.IPSECTunnelDeleteParams{
			AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			XMagicNewHcTarget: cloudflare.F(true),
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

func TestIPSECTunnelBulkUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.IPSECTunnels.BulkUpdate(context.TODO(), magic_transit.IPSECTunnelBulkUpdateParams{
		AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body:              map[string]interface{}{},
		XMagicNewHcTarget: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIPSECTunnelGetWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.IPSECTunnels.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.IPSECTunnelGetParams{
			AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			XMagicNewHcTarget: cloudflare.F(true),
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

func TestIPSECTunnelPSKGenerate(t *testing.T) {
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
	_, err := client.MagicTransit.IPSECTunnels.PSKGenerate(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.IPSECTunnelPSKGenerateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body:      map[string]interface{}{},
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
