// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestSettingEditWithOptionalParams(t *testing.T) {
	t.Skip("HTTP 422 from prism")
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
	_, err := client.DNS.Settings.Edit(context.TODO(), dns.SettingEditParams{
		AccountID: cloudflare.F("account_id"),
		ZoneDefaults: cloudflare.F(dns.SettingEditParamsZoneDefaults{
			FlattenAllCNAMEs: cloudflare.F(false),
			FoundationDNS:    cloudflare.F(false),
			MultiProvider:    cloudflare.F(false),
			Nameservers: cloudflare.F(dns.SettingEditParamsZoneDefaultsNameservers{
				Type: cloudflare.F(dns.SettingEditParamsZoneDefaultsNameserversTypeCloudflareStandard),
			}),
			NSTTL:              cloudflare.F(86400.000000),
			SecondaryOverrides: cloudflare.F(false),
			SOA: cloudflare.F(dns.SettingEditParamsZoneDefaultsSOA{
				Expire:  cloudflare.F(604800.000000),
				MinTTL:  cloudflare.F(1800.000000),
				MNAME:   cloudflare.F("kristina.ns.cloudflare.com"),
				Refresh: cloudflare.F(10000.000000),
				Retry:   cloudflare.F(2400.000000),
				RNAME:   cloudflare.F("admin.example.com"),
				TTL:     cloudflare.F(3600.000000),
			}),
			ZoneMode: cloudflare.F(dns.SettingEditParamsZoneDefaultsZoneModeStandard),
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

func TestSettingGetWithOptionalParams(t *testing.T) {
	t.Skip("HTTP 422 from prism")
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
	_, err := client.DNS.Settings.Get(context.TODO(), dns.SettingGetParams{
		AccountID: cloudflare.F("account_id"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
