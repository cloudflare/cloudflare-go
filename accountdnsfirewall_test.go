// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestAccountDNSFirewallGet(t *testing.T) {
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
	_, err := client.Accounts.DNSFirewalls.Get(
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

func TestAccountDNSFirewallUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.DNSFirewalls.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountDNSFirewallUpdateParams{
			DeprecateAnyRequests: cloudflare.F(true),
			DNSFirewallIPs:       cloudflare.F([]cloudflare.AccountDNSFirewallUpdateParamsDNSFirewallIP{shared.UnionString("203.0.113.1"), shared.UnionString("203.0.113.254"), shared.UnionString("2001:DB8:AB::CF"), shared.UnionString("2001:DB8:CD::CF")}),
			EcsFallback:          cloudflare.F(false),
			MaximumCacheTtl:      cloudflare.F(900.000000),
			MinimumCacheTtl:      cloudflare.F(60.000000),
			Name:                 cloudflare.F("My Awesome DNS Firewall cluster"),
			UpstreamIPs:          cloudflare.F([]cloudflare.AccountDNSFirewallUpdateParamsUpstreamIP{shared.UnionString("192.0.2.1"), shared.UnionString("198.51.100.1"), shared.UnionString("2001:DB8:100::CF")}),
			AttackMitigation: cloudflare.F(cloudflare.AccountDNSFirewallUpdateParamsAttackMitigation{
				Enabled:                   cloudflare.F(true),
				OnlyWhenOriginUnhealthy:   cloudflare.F[any](map[string]interface{}{}),
				OnlyWhenUpstreamUnhealthy: cloudflare.F(false),
			}),
			NegativeCacheTtl: cloudflare.F(900.000000),
			OriginIPs:        cloudflare.F[any](map[string]interface{}{}),
			Ratelimit:        cloudflare.F(600.000000),
			Retries:          cloudflare.F(2.000000),
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

func TestAccountDNSFirewallDelete(t *testing.T) {
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
	_, err := client.Accounts.DNSFirewalls.Delete(
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

func TestAccountDNSFirewallDNSFirewallNewDNSFirewallClusterWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.DNSFirewalls.DNSFirewallNewDNSFirewallCluster(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParams{
			Name:        cloudflare.F("My Awesome DNS Firewall cluster"),
			UpstreamIPs: cloudflare.F([]cloudflare.AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsUpstreamIP{shared.UnionString("192.0.2.1"), shared.UnionString("198.51.100.1"), shared.UnionString("2001:DB8:100::CF")}),
			AttackMitigation: cloudflare.F(cloudflare.AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsAttackMitigation{
				Enabled:                   cloudflare.F(true),
				OnlyWhenOriginUnhealthy:   cloudflare.F[any](map[string]interface{}{}),
				OnlyWhenUpstreamUnhealthy: cloudflare.F(false),
			}),
			DeprecateAnyRequests: cloudflare.F(true),
			EcsFallback:          cloudflare.F(false),
			MaximumCacheTtl:      cloudflare.F(900.000000),
			MinimumCacheTtl:      cloudflare.F(60.000000),
			NegativeCacheTtl:     cloudflare.F(900.000000),
			OriginIPs:            cloudflare.F[any](map[string]interface{}{}),
			Ratelimit:            cloudflare.F(600.000000),
			Retries:              cloudflare.F(2.000000),
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

func TestAccountDNSFirewallDNSFirewallListDNSFirewallClustersWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.DNSFirewalls.DNSFirewallListDNSFirewallClusters(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountDNSFirewallDNSFirewallListDNSFirewallClustersParams{
			Page:    cloudflare.F(1.000000),
			PerPage: cloudflare.F(1.000000),
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
