// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_test

import (
  "context"
  "errors"
  "os"
  "testing"

  "github.com/cloudflare/cloudflare-go/v2"
  "github.com/cloudflare/cloudflare-go/v2/dns"
  "github.com/cloudflare/cloudflare-go/v2/internal/testutil"
  "github.com/cloudflare/cloudflare-go/v2/option"
)

func TestFirewallNewWithOptionalParams(t *testing.T) {
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
  _, err := client.DNS.Firewall.New(context.TODO(), dns.FirewallNewParams{
    AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
    Name: cloudflare.F("My Awesome DNS Firewall cluster"),
    UpstreamIPs: cloudflare.F([]dns.UpstreamIPsParam{"192.0.2.1", "198.51.100.1", "string"}),
    AttackMitigation: cloudflare.F(dns.AttackMitigationParam{
      Enabled: cloudflare.F(true),
      OnlyWhenUpstreamUnhealthy: cloudflare.F(false),
    }),
    DeprecateAnyRequests: cloudflare.F(true),
    ECSFallback: cloudflare.F(false),
    MaximumCacheTTL: cloudflare.F(900.000000),
    MinimumCacheTTL: cloudflare.F(60.000000),
    NegativeCacheTTL: cloudflare.F(900.000000),
    Ratelimit: cloudflare.F(600.000000),
    Retries: cloudflare.F(2.000000),
  })
  if err != nil {
    var apierr *cloudflare.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestFirewallListWithOptionalParams(t *testing.T) {
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
  _, err := client.DNS.Firewall.List(context.TODO(), dns.FirewallListParams{
    AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
    Page: cloudflare.F(1.000000),
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

func TestFirewallDelete(t *testing.T) {
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
  _, err := client.DNS.Firewall.Delete(
    context.TODO(),
    "023e105f4ecef8ad9ca31a8372d0c353",
    dns.FirewallDeleteParams{
      AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestFirewallEditWithOptionalParams(t *testing.T) {
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
  _, err := client.DNS.Firewall.Edit(
    context.TODO(),
    "023e105f4ecef8ad9ca31a8372d0c353",
    dns.FirewallEditParams{
      AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
      AttackMitigation: cloudflare.F(dns.AttackMitigationParam{
        Enabled: cloudflare.F(true),
        OnlyWhenUpstreamUnhealthy: cloudflare.F(false),
      }),
      DeprecateAnyRequests: cloudflare.F(true),
      ECSFallback: cloudflare.F(false),
      MaximumCacheTTL: cloudflare.F(900.000000),
      MinimumCacheTTL: cloudflare.F(60.000000),
      Name: cloudflare.F("My Awesome DNS Firewall cluster"),
      NegativeCacheTTL: cloudflare.F(900.000000),
      Ratelimit: cloudflare.F(600.000000),
      Retries: cloudflare.F(2.000000),
      UpstreamIPs: cloudflare.F([]dns.UpstreamIPsParam{"192.0.2.1", "198.51.100.1", "string"}),
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

func TestFirewallGet(t *testing.T) {
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
  _, err := client.DNS.Firewall.Get(
    context.TODO(),
    "023e105f4ecef8ad9ca31a8372d0c353",
    dns.FirewallGetParams{
      AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
