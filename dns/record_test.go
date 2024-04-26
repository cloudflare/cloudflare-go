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
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

func TestRecordNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.DNS.Records.New(context.TODO(), dns.RecordNewParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Record: dns.ARecordParam{
			Content: cloudflare.F("198.51.100.4"),
			Name:    cloudflare.F("example.com"),
			Proxied: cloudflare.F(false),
			Type:    cloudflare.F(dns.ARecordTypeA),
			Comment: cloudflare.F("Domain verification record"),
			Tags:    cloudflare.F([]dns.RecordTagsParam{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
			TTL:     cloudflare.F[dns.TTLUnionParam](shared.UnionFloat(3600.000000)),
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

func TestRecordUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.DNS.Records.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.RecordUpdateParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Record: dns.ARecordParam{
				Content: cloudflare.F("198.51.100.4"),
				Name:    cloudflare.F("example.com"),
				Proxied: cloudflare.F(false),
				Type:    cloudflare.F(dns.ARecordTypeA),
				Comment: cloudflare.F("Domain verification record"),
				Tags:    cloudflare.F([]dns.RecordTagsParam{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
				TTL:     cloudflare.F[dns.TTLUnionParam](shared.UnionFloat(3600.000000)),
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

func TestRecordListWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.DNS.Records.List(context.TODO(), dns.RecordListParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Comment: cloudflare.F(dns.RecordListParamsComment{
			Present:    cloudflare.F("string"),
			Absent:     cloudflare.F("string"),
			Exact:      cloudflare.F("Hello, world"),
			Contains:   cloudflare.F("ello, worl"),
			Startswith: cloudflare.F("Hello, w"),
			Endswith:   cloudflare.F("o, world"),
		}),
		Content:   cloudflare.F("127.0.0.1"),
		Direction: cloudflare.F(dns.RecordListParamsDirectionAsc),
		Match:     cloudflare.F(dns.RecordListParamsMatchAny),
		Name:      cloudflare.F("example.com"),
		Order:     cloudflare.F(dns.RecordListParamsOrderType),
		Page:      cloudflare.F(1.000000),
		PerPage:   cloudflare.F(5.000000),
		Proxied:   cloudflare.F(false),
		Search:    cloudflare.F("www.cloudflare.com"),
		Tag: cloudflare.F(dns.RecordListParamsTag{
			Present:    cloudflare.F("important"),
			Absent:     cloudflare.F("important"),
			Exact:      cloudflare.F("greeting:Hello, world"),
			Contains:   cloudflare.F("greeting:ello, worl"),
			Startswith: cloudflare.F("greeting:Hello, w"),
			Endswith:   cloudflare.F("greeting:o, world"),
		}),
		TagMatch: cloudflare.F(dns.RecordListParamsTagMatchAny),
		Type:     cloudflare.F(dns.RecordListParamsTypeA),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecordDelete(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.DNS.Records.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.RecordDeleteParams{
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

func TestRecordEditWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.DNS.Records.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.RecordEditParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Record: dns.ARecordParam{
				Content: cloudflare.F("198.51.100.4"),
				Name:    cloudflare.F("example.com"),
				Proxied: cloudflare.F(false),
				Type:    cloudflare.F(dns.ARecordTypeA),
				Comment: cloudflare.F("Domain verification record"),
				Tags:    cloudflare.F([]dns.RecordTagsParam{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
				TTL:     cloudflare.F[dns.TTLUnionParam](shared.UnionFloat(3600.000000)),
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

func TestRecordExport(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.DNS.Records.Export(context.TODO(), dns.RecordExportParams{
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

func TestRecordGet(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.DNS.Records.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.RecordGetParams{
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

func TestRecordImportWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.DNS.Records.Import(context.TODO(), dns.RecordImportParams{
		ZoneID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		File:    cloudflare.F("www.example.com. 300 IN  A 127.0.0.1"),
		Proxied: cloudflare.F("true"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecordScan(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.DNS.Records.Scan(context.TODO(), dns.RecordScanParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body:   map[string]interface{}{},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
