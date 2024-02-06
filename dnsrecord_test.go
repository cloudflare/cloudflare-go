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

func TestDNSRecordGet(t *testing.T) {
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
	_, err := client.DNSRecords.Get(
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

func TestDNSRecordUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.DNSRecords.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.DNSRecordUpdateParamsDNSRecordsARecord{
			Content: cloudflare.F("198.51.100.4"),
			Name:    cloudflare.F("example.com"),
			Type:    cloudflare.F(cloudflare.DNSRecordUpdateParamsDNSRecordsARecordTypeA),
			Comment: cloudflare.F("Domain verification record"),
			Proxied: cloudflare.F(false),
			Tags:    cloudflare.F([]string{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
			TTL:     cloudflare.F[cloudflare.DNSRecordUpdateParamsDNSRecordsARecordTTL](shared.UnionFloat(3600.000000)),
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

func TestDNSRecordDelete(t *testing.T) {
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
	_, err := client.DNSRecords.Delete(
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

func TestDNSRecordDNSRecordsForAZoneNewDNSRecordWithOptionalParams(t *testing.T) {
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
	_, err := client.DNSRecords.DNSRecordsForAZoneNewDNSRecord(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecord{
			Content: cloudflare.F("198.51.100.4"),
			Name:    cloudflare.F("example.com"),
			Type:    cloudflare.F(cloudflare.DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecordTypeA),
			Comment: cloudflare.F("Domain verification record"),
			Proxied: cloudflare.F(false),
			Tags:    cloudflare.F([]string{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
			TTL:     cloudflare.F[cloudflare.DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecordTTL](shared.UnionFloat(3600.000000)),
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

func TestDNSRecordDNSRecordsForAZoneListDNSRecordsWithOptionalParams(t *testing.T) {
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
	_, err := client.DNSRecords.DNSRecordsForAZoneListDNSRecords(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.DNSRecordDNSRecordsForAZoneListDNSRecordsParams{
			Comment: cloudflare.F(cloudflare.DNSRecordDNSRecordsForAZoneListDNSRecordsParamsComment{
				Present:    cloudflare.F("string"),
				Absent:     cloudflare.F("string"),
				Exact:      cloudflare.F("Hello, world"),
				Contains:   cloudflare.F("ello, worl"),
				Startswith: cloudflare.F("Hello, w"),
				Endswith:   cloudflare.F("o, world"),
			}),
			Content:   cloudflare.F("127.0.0.1"),
			Direction: cloudflare.F(cloudflare.DNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirectionAsc),
			Match:     cloudflare.F(cloudflare.DNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatchAny),
			Name:      cloudflare.F("example.com"),
			Order:     cloudflare.F(cloudflare.DNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrderType),
			Page:      cloudflare.F(1.000000),
			PerPage:   cloudflare.F(5.000000),
			Proxied:   cloudflare.F(false),
			Search:    cloudflare.F("www.cloudflare.com"),
			Tag: cloudflare.F(cloudflare.DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTag{
				Present:    cloudflare.F("important"),
				Absent:     cloudflare.F("important"),
				Exact:      cloudflare.F("greeting:Hello, world"),
				Contains:   cloudflare.F("greeting:ello, worl"),
				Startswith: cloudflare.F("greeting:Hello, w"),
				Endswith:   cloudflare.F("greeting:o, world"),
			}),
			TagMatch: cloudflare.F(cloudflare.DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatchAny),
			Type:     cloudflare.F(cloudflare.DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeA),
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
