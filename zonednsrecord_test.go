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

func TestZoneDNSRecordGet(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.Get(
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

func TestZoneDNSRecordUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneDNSRecordUpdateParamsMpBuiT95ARecord{
			Content: cloudflare.F("198.51.100.4"),
			Name:    cloudflare.F("example.com"),
			Type:    cloudflare.F(cloudflare.ZoneDNSRecordUpdateParamsMpBuiT95ARecordTypeA),
			Comment: cloudflare.F("Domain verification record"),
			Proxied: cloudflare.F(false),
			Tags:    cloudflare.F([]string{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
			Ttl:     cloudflare.F[cloudflare.ZoneDNSRecordUpdateParamsMpBuiT95ARecordTtl](shared.UnionFloat(3600.000000)),
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

func TestZoneDNSRecordDelete(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.Delete(
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

func TestZoneDNSRecordDNSRecordsForAZoneNewDNSRecordWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.DNSRecordsForAZoneNewDNSRecord(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecord{
			Content: cloudflare.F("198.51.100.4"),
			Name:    cloudflare.F("example.com"),
			Type:    cloudflare.F(cloudflare.ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecordTypeA),
			Comment: cloudflare.F("Domain verification record"),
			Proxied: cloudflare.F(false),
			Tags:    cloudflare.F([]string{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
			Ttl:     cloudflare.F[cloudflare.ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecordTtl](shared.UnionFloat(3600.000000)),
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

func TestZoneDNSRecordDNSRecordsForAZoneListDNSRecordsWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.DNSRecordsForAZoneListDNSRecords(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParams{
			Comment: cloudflare.F(cloudflare.ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsComment{
				Present:    cloudflare.F("string"),
				Absent:     cloudflare.F("string"),
				Exact:      cloudflare.F("Hello, world"),
				Contains:   cloudflare.F("ello, worl"),
				Startswith: cloudflare.F("Hello, w"),
				Endswith:   cloudflare.F("o, world"),
			}),
			Content:   cloudflare.F("127.0.0.1"),
			Direction: cloudflare.F(cloudflare.ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirectionAsc),
			Match:     cloudflare.F(cloudflare.ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatchAny),
			Name:      cloudflare.F("example.com"),
			Order:     cloudflare.F(cloudflare.ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrderType),
			Page:      cloudflare.F(1.000000),
			PerPage:   cloudflare.F(5.000000),
			Proxied:   cloudflare.F(false),
			Search:    cloudflare.F("www.cloudflare.com"),
			Tag: cloudflare.F(cloudflare.ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTag{
				Present:    cloudflare.F("important"),
				Absent:     cloudflare.F("important"),
				Exact:      cloudflare.F("greeting:Hello, world"),
				Contains:   cloudflare.F("greeting:ello, worl"),
				Startswith: cloudflare.F("greeting:Hello, w"),
				Endswith:   cloudflare.F("greeting:o, world"),
			}),
			TagMatch: cloudflare.F(cloudflare.ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatchAny),
			Type:     cloudflare.F(cloudflare.ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeA),
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

func TestZoneDNSRecordPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.Patch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneDNSRecordPatchParamsMpBuiT95ARecord{
			Content: cloudflare.F("198.51.100.4"),
			Name:    cloudflare.F("example.com"),
			Type:    cloudflare.F(cloudflare.ZoneDNSRecordPatchParamsMpBuiT95ARecordTypeA),
			Comment: cloudflare.F("Domain verification record"),
			Proxied: cloudflare.F(false),
			Tags:    cloudflare.F([]string{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
			Ttl:     cloudflare.F[cloudflare.ZoneDNSRecordPatchParamsMpBuiT95ARecordTtl](shared.UnionFloat(3600.000000)),
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
