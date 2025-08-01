// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/dns"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

func TestSettingZoneEditWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Settings.Zone.Edit(context.TODO(), dns.SettingZoneEditParams{
		ZoneID:           cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		FlattenAllCNAMEs: cloudflare.F(false),
		FoundationDNS:    cloudflare.F(false),
		InternalDNS: cloudflare.F(dns.SettingZoneEditParamsInternalDNS{
			ReferenceZoneID: cloudflare.F("reference_zone_id"),
		}),
		MultiProvider: cloudflare.F(false),
		Nameservers: cloudflare.F(dns.SettingZoneEditParamsNameservers{
			Type:  cloudflare.F(dns.SettingZoneEditParamsNameserversTypeCloudflareStandard),
			NSSet: cloudflare.F(int64(1)),
		}),
		NSTTL:              cloudflare.F(86400.000000),
		SecondaryOverrides: cloudflare.F(false),
		SOA: cloudflare.F(dns.SettingZoneEditParamsSOA{
			Expire:  cloudflare.F(604800.000000),
			MinTTL:  cloudflare.F(1800.000000),
			MNAME:   cloudflare.F("kristina.ns.cloudflare.com"),
			Refresh: cloudflare.F(10000.000000),
			Retry:   cloudflare.F(2400.000000),
			RNAME:   cloudflare.F("admin.example.com"),
			TTL:     cloudflare.F(3600.000000),
		}),
		ZoneMode: cloudflare.F(dns.SettingZoneEditParamsZoneModeDNSOnly),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingZoneGet(t *testing.T) {
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
	_, err := client.DNS.Settings.Zone.Get(context.TODO(), dns.SettingZoneGetParams{
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
