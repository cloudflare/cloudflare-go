// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
)

func TestGatewayLocationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Locations.New(context.TODO(), zero_trust.GatewayLocationNewParams{
		AccountID:           cloudflare.F("699d98642c564d2e855e9661899b7252"),
		Name:                cloudflare.F("Austin Office Location"),
		ClientDefault:       cloudflare.F(false),
		DNSDestinationIPsID: cloudflare.F("0e4a32c6-6fb8-4858-9296-98f51631e8e6"),
		ECSSupport:          cloudflare.F(false),
		Endpoints: cloudflare.F(zero_trust.EndpointParam{
			DOH: cloudflare.F(zero_trust.DOHEndpointParam{
				Enabled: cloudflare.F(true),
				Networks: cloudflare.F([]zero_trust.IPNetworkParam{{
					Network: cloudflare.F("2001:85a3::/64"),
				}}),
				RequireToken: cloudflare.F(true),
			}),
			DOT: cloudflare.F(zero_trust.DOTEndpointParam{
				Enabled: cloudflare.F(true),
				Networks: cloudflare.F([]zero_trust.IPNetworkParam{{
					Network: cloudflare.F("2001:85a3::/64"),
				}}),
			}),
			IPV4: cloudflare.F(zero_trust.IPV4EndpointParam{
				Enabled: cloudflare.F(true),
			}),
			IPV6: cloudflare.F(zero_trust.IPV6EndpointParam{
				Enabled: cloudflare.F(true),
				Networks: cloudflare.F([]zero_trust.IPV6NetworkParam{{
					Network: cloudflare.F("2001:85a3::/64"),
				}}),
			}),
		}),
		Networks: cloudflare.F([]zero_trust.GatewayLocationNewParamsNetwork{{
			Network: cloudflare.F("192.0.2.1/32"),
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

func TestGatewayLocationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Locations.Update(
		context.TODO(),
		"ed35569b41ce4d1facfe683550f54086",
		zero_trust.GatewayLocationUpdateParams{
			AccountID:           cloudflare.F("699d98642c564d2e855e9661899b7252"),
			Name:                cloudflare.F("Austin Office Location"),
			ClientDefault:       cloudflare.F(false),
			DNSDestinationIPsID: cloudflare.F("0e4a32c6-6fb8-4858-9296-98f51631e8e6"),
			ECSSupport:          cloudflare.F(false),
			Endpoints: cloudflare.F(zero_trust.EndpointParam{
				DOH: cloudflare.F(zero_trust.DOHEndpointParam{
					Enabled: cloudflare.F(true),
					Networks: cloudflare.F([]zero_trust.IPNetworkParam{{
						Network: cloudflare.F("2001:85a3::/64"),
					}}),
					RequireToken: cloudflare.F(true),
				}),
				DOT: cloudflare.F(zero_trust.DOTEndpointParam{
					Enabled: cloudflare.F(true),
					Networks: cloudflare.F([]zero_trust.IPNetworkParam{{
						Network: cloudflare.F("2001:85a3::/64"),
					}}),
				}),
				IPV4: cloudflare.F(zero_trust.IPV4EndpointParam{
					Enabled: cloudflare.F(true),
				}),
				IPV6: cloudflare.F(zero_trust.IPV6EndpointParam{
					Enabled: cloudflare.F(true),
					Networks: cloudflare.F([]zero_trust.IPV6NetworkParam{{
						Network: cloudflare.F("2001:85a3::/64"),
					}}),
				}),
			}),
			Networks: cloudflare.F([]zero_trust.GatewayLocationUpdateParamsNetwork{{
				Network: cloudflare.F("192.0.2.1/32"),
			}}),
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

func TestGatewayLocationList(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Locations.List(context.TODO(), zero_trust.GatewayLocationListParams{
		AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayLocationDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Locations.Delete(
		context.TODO(),
		"ed35569b41ce4d1facfe683550f54086",
		zero_trust.GatewayLocationDeleteParams{
			AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
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

func TestGatewayLocationGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Locations.Get(
		context.TODO(),
		"ed35569b41ce4d1facfe683550f54086",
		zero_trust.GatewayLocationGetParams{
			AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
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
