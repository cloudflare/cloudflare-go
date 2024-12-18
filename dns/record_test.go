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
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

func TestRecordNewWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Records.New(context.TODO(), dns.RecordNewParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Record: dns.ARecordParam{
			Comment: cloudflare.F("Domain verification record"),
			Content: cloudflare.F("198.51.100.4"),
			Name:    cloudflare.F("example.com"),
			Proxied: cloudflare.F(true),
			Settings: cloudflare.F(dns.ARecordSettingsParam{
				IPV4Only: cloudflare.F(true),
				IPV6Only: cloudflare.F(true),
			}),
			Tags: cloudflare.F([]dns.RecordTagsParam{"owner:dns-team"}),
			TTL:  cloudflare.F(dns.TTL1),
			Type: cloudflare.F(dns.ARecordTypeA),
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
	_, err := client.DNS.Records.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.RecordUpdateParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Record: dns.ARecordParam{
				Comment: cloudflare.F("Domain verification record"),
				Content: cloudflare.F("198.51.100.4"),
				Name:    cloudflare.F("example.com"),
				Proxied: cloudflare.F(true),
				Settings: cloudflare.F(dns.ARecordSettingsParam{
					IPV4Only: cloudflare.F(true),
					IPV6Only: cloudflare.F(true),
				}),
				Tags: cloudflare.F([]dns.RecordTagsParam{"owner:dns-team"}),
				TTL:  cloudflare.F(dns.TTL1),
				Type: cloudflare.F(dns.ARecordTypeA),
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
			Absent:     cloudflare.F("absent"),
			Contains:   cloudflare.F("ello, worl"),
			Endswith:   cloudflare.F("o, world"),
			Exact:      cloudflare.F("Hello, world"),
			Present:    cloudflare.F("present"),
			Startswith: cloudflare.F("Hello, w"),
		}),
		Content: cloudflare.F(dns.RecordListParamsContent{
			Contains:   cloudflare.F("7.0.0."),
			Endswith:   cloudflare.F(".0.1"),
			Exact:      cloudflare.F("127.0.0.1"),
			Startswith: cloudflare.F("127.0."),
		}),
		Direction: cloudflare.F(shared.SortDirectionAsc),
		Match:     cloudflare.F(dns.RecordListParamsMatchAny),
		Name: cloudflare.F(dns.RecordListParamsName{
			Contains:   cloudflare.F("w.example."),
			Endswith:   cloudflare.F(".example.com"),
			Exact:      cloudflare.F("www.example.com"),
			Startswith: cloudflare.F("www.example"),
		}),
		Order:   cloudflare.F(dns.RecordListParamsOrderType),
		Page:    cloudflare.F(1.000000),
		PerPage: cloudflare.F(5.000000),
		Proxied: cloudflare.F(true),
		Search:  cloudflare.F("www.cloudflare.com"),
		Tag: cloudflare.F(dns.RecordListParamsTag{
			Absent:     cloudflare.F("important"),
			Contains:   cloudflare.F("greeting:ello, worl"),
			Endswith:   cloudflare.F("greeting:o, world"),
			Exact:      cloudflare.F("greeting:Hello, world"),
			Present:    cloudflare.F("important"),
			Startswith: cloudflare.F("greeting:Hello, w"),
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

func TestRecordBatchWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate auth errors on test suite")
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
	_, err := client.DNS.Records.Batch(context.TODO(), dns.RecordBatchParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Deletes: cloudflare.F([]dns.RecordBatchParamsDelete{{
			ID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		}}),
		Patches: cloudflare.F([]dns.RecordUnionParam{dns.ARecordParam{
			Comment: cloudflare.F("Domain verification record"),
			Content: cloudflare.F("198.51.100.4"),
			Name:    cloudflare.F("example.com"),
			Proxied: cloudflare.F(true),
			Settings: cloudflare.F(dns.ARecordSettingsParam{
				IPV4Only: cloudflare.F(true),
				IPV6Only: cloudflare.F(true),
			}),
			Tags: cloudflare.F([]dns.RecordTagsParam{"owner:dns-team"}),
			TTL:  cloudflare.F(dns.TTL1),
			Type: cloudflare.F(dns.ARecordTypeA),
		}}),
		Posts: cloudflare.F([]dns.RecordUnionParam{dns.ARecordParam{
			Comment: cloudflare.F("Domain verification record"),
			Content: cloudflare.F("198.51.100.4"),
			Name:    cloudflare.F("example.com"),
			Proxied: cloudflare.F(true),
			Settings: cloudflare.F(dns.ARecordSettingsParam{
				IPV4Only: cloudflare.F(true),
				IPV6Only: cloudflare.F(true),
			}),
			Tags: cloudflare.F([]dns.RecordTagsParam{"owner:dns-team"}),
			TTL:  cloudflare.F(dns.TTL1),
			Type: cloudflare.F(dns.ARecordTypeA),
		}}),
		Puts: cloudflare.F([]dns.RecordUnionParam{dns.ARecordParam{
			Comment: cloudflare.F("Domain verification record"),
			Content: cloudflare.F("198.51.100.4"),
			Name:    cloudflare.F("example.com"),
			Proxied: cloudflare.F(true),
			Settings: cloudflare.F(dns.ARecordSettingsParam{
				IPV4Only: cloudflare.F(true),
				IPV6Only: cloudflare.F(true),
			}),
			Tags: cloudflare.F([]dns.RecordTagsParam{"owner:dns-team"}),
			TTL:  cloudflare.F(dns.TTL1),
			Type: cloudflare.F(dns.ARecordTypeA),
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

func TestRecordEditWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Records.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.RecordEditParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Record: dns.ARecordParam{
				Comment: cloudflare.F("Domain verification record"),
				Content: cloudflare.F("198.51.100.4"),
				Name:    cloudflare.F("example.com"),
				Proxied: cloudflare.F(true),
				Settings: cloudflare.F(dns.ARecordSettingsParam{
					IPV4Only: cloudflare.F(true),
					IPV6Only: cloudflare.F(true),
				}),
				Tags: cloudflare.F([]dns.RecordTagsParam{"owner:dns-team"}),
				TTL:  cloudflare.F(dns.TTL1),
				Type: cloudflare.F(dns.ARecordTypeA),
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
