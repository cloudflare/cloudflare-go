// File generated from our OpenAPI spec by Stainless.

package dns_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/dns"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
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
		ZoneID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Name:    cloudflare.F("example.com"),
		Type:    cloudflare.F(dns.RecordNewParamsTypeURI),
		Comment: cloudflare.F("Domain verification record"),
		Data: cloudflare.F(dns.RecordNewParamsData{
			Flags:         cloudflare.F("string"),
			Tag:           cloudflare.F("issue"),
			Value:         cloudflare.F("alpn=\"h3,h2\" ipv4hint=\"127.0.0.1\" ipv6hint=\"::1\""),
			Algorithm:     cloudflare.F(2.000000),
			Certificate:   cloudflare.F("string"),
			KeyTag:        cloudflare.F(1.000000),
			Type:          cloudflare.F(1.000000),
			Protocol:      cloudflare.F(3.000000),
			PublicKey:     cloudflare.F("string"),
			Digest:        cloudflare.F("string"),
			DigestType:    cloudflare.F(1.000000),
			Priority:      cloudflare.F(1.000000),
			Target:        cloudflare.F("."),
			Altitude:      cloudflare.F(0.000000),
			LatDegrees:    cloudflare.F(37.000000),
			LatDirection:  cloudflare.F(dns.RecordNewParamsDataLatDirectionN),
			LatMinutes:    cloudflare.F(46.000000),
			LatSeconds:    cloudflare.F(46.000000),
			LongDegrees:   cloudflare.F(122.000000),
			LongDirection: cloudflare.F(dns.RecordNewParamsDataLongDirectionW),
			LongMinutes:   cloudflare.F(23.000000),
			LongSeconds:   cloudflare.F(35.000000),
			PrecisionHorz: cloudflare.F(0.000000),
			PrecisionVert: cloudflare.F(0.000000),
			Size:          cloudflare.F(100.000000),
			Order:         cloudflare.F(100.000000),
			Preference:    cloudflare.F(10.000000),
			Regex:         cloudflare.F("string"),
			Replacement:   cloudflare.F("string"),
			Service:       cloudflare.F("_sip"),
			MatchingType:  cloudflare.F(1.000000),
			Selector:      cloudflare.F(0.000000),
			Usage:         cloudflare.F(0.000000),
			Name:          cloudflare.F("example.com"),
			Port:          cloudflare.F(8806.000000),
			Proto:         cloudflare.F("_tcp"),
			Weight:        cloudflare.F(20.000000),
			Fingerprint:   cloudflare.F("string"),
			Content:       cloudflare.F("http://example.com/example.html"),
		}),
		Meta: cloudflare.F(dns.RecordNewParamsMeta{
			AutoAdded: cloudflare.F(true),
			Source:    cloudflare.F("primary"),
		}),
		Priority: cloudflare.F(10.000000),
		Proxied:  cloudflare.F(false),
		Tags:     cloudflare.F([]string{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
		TTL:      cloudflare.F[dns.RecordNewParamsTTL](shared.UnionFloat(3600.000000)),
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
			ZoneID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:    cloudflare.F("example.com"),
			Type:    cloudflare.F(dns.RecordUpdateParamsTypeURI),
			Comment: cloudflare.F("Domain verification record"),
			Data: cloudflare.F(dns.RecordUpdateParamsData{
				Flags:         cloudflare.F("string"),
				Tag:           cloudflare.F("issue"),
				Value:         cloudflare.F("alpn=\"h3,h2\" ipv4hint=\"127.0.0.1\" ipv6hint=\"::1\""),
				Algorithm:     cloudflare.F(2.000000),
				Certificate:   cloudflare.F("string"),
				KeyTag:        cloudflare.F(1.000000),
				Type:          cloudflare.F(1.000000),
				Protocol:      cloudflare.F(3.000000),
				PublicKey:     cloudflare.F("string"),
				Digest:        cloudflare.F("string"),
				DigestType:    cloudflare.F(1.000000),
				Priority:      cloudflare.F(1.000000),
				Target:        cloudflare.F("."),
				Altitude:      cloudflare.F(0.000000),
				LatDegrees:    cloudflare.F(37.000000),
				LatDirection:  cloudflare.F(dns.RecordUpdateParamsDataLatDirectionN),
				LatMinutes:    cloudflare.F(46.000000),
				LatSeconds:    cloudflare.F(46.000000),
				LongDegrees:   cloudflare.F(122.000000),
				LongDirection: cloudflare.F(dns.RecordUpdateParamsDataLongDirectionW),
				LongMinutes:   cloudflare.F(23.000000),
				LongSeconds:   cloudflare.F(35.000000),
				PrecisionHorz: cloudflare.F(0.000000),
				PrecisionVert: cloudflare.F(0.000000),
				Size:          cloudflare.F(100.000000),
				Order:         cloudflare.F(100.000000),
				Preference:    cloudflare.F(10.000000),
				Regex:         cloudflare.F("string"),
				Replacement:   cloudflare.F("string"),
				Service:       cloudflare.F("_sip"),
				MatchingType:  cloudflare.F(1.000000),
				Selector:      cloudflare.F(0.000000),
				Usage:         cloudflare.F(0.000000),
				Name:          cloudflare.F("example.com"),
				Port:          cloudflare.F(8806.000000),
				Proto:         cloudflare.F("_tcp"),
				Weight:        cloudflare.F(20.000000),
				Fingerprint:   cloudflare.F("string"),
				Content:       cloudflare.F("http://example.com/example.html"),
			}),
			Meta: cloudflare.F(dns.RecordUpdateParamsMeta{
				AutoAdded: cloudflare.F(true),
				Source:    cloudflare.F("primary"),
			}),
			Priority: cloudflare.F(10.000000),
			Proxied:  cloudflare.F(false),
			Tags:     cloudflare.F([]string{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
			TTL:      cloudflare.F[dns.RecordUpdateParamsTTL](shared.UnionFloat(3600.000000)),
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
			ZoneID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:    cloudflare.F("example.com"),
			Type:    cloudflare.F(dns.RecordEditParamsTypeURI),
			Comment: cloudflare.F("Domain verification record"),
			Data: cloudflare.F(dns.RecordEditParamsData{
				Flags:         cloudflare.F("string"),
				Tag:           cloudflare.F("issue"),
				Value:         cloudflare.F("alpn=\"h3,h2\" ipv4hint=\"127.0.0.1\" ipv6hint=\"::1\""),
				Algorithm:     cloudflare.F(2.000000),
				Certificate:   cloudflare.F("string"),
				KeyTag:        cloudflare.F(1.000000),
				Type:          cloudflare.F(1.000000),
				Protocol:      cloudflare.F(3.000000),
				PublicKey:     cloudflare.F("string"),
				Digest:        cloudflare.F("string"),
				DigestType:    cloudflare.F(1.000000),
				Priority:      cloudflare.F(1.000000),
				Target:        cloudflare.F("."),
				Altitude:      cloudflare.F(0.000000),
				LatDegrees:    cloudflare.F(37.000000),
				LatDirection:  cloudflare.F(dns.RecordEditParamsDataLatDirectionN),
				LatMinutes:    cloudflare.F(46.000000),
				LatSeconds:    cloudflare.F(46.000000),
				LongDegrees:   cloudflare.F(122.000000),
				LongDirection: cloudflare.F(dns.RecordEditParamsDataLongDirectionW),
				LongMinutes:   cloudflare.F(23.000000),
				LongSeconds:   cloudflare.F(35.000000),
				PrecisionHorz: cloudflare.F(0.000000),
				PrecisionVert: cloudflare.F(0.000000),
				Size:          cloudflare.F(100.000000),
				Order:         cloudflare.F(100.000000),
				Preference:    cloudflare.F(10.000000),
				Regex:         cloudflare.F("string"),
				Replacement:   cloudflare.F("string"),
				Service:       cloudflare.F("_sip"),
				MatchingType:  cloudflare.F(1.000000),
				Selector:      cloudflare.F(0.000000),
				Usage:         cloudflare.F(0.000000),
				Name:          cloudflare.F("example.com"),
				Port:          cloudflare.F(8806.000000),
				Proto:         cloudflare.F("_tcp"),
				Weight:        cloudflare.F(20.000000),
				Fingerprint:   cloudflare.F("string"),
				Content:       cloudflare.F("http://example.com/example.html"),
			}),
			Meta: cloudflare.F(dns.RecordEditParamsMeta{
				AutoAdded: cloudflare.F(true),
				Source:    cloudflare.F("primary"),
			}),
			Priority: cloudflare.F(10.000000),
			Proxied:  cloudflare.F(false),
			Tags:     cloudflare.F([]string{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
			TTL:      cloudflare.F[dns.RecordEditParamsTTL](shared.UnionFloat(3600.000000)),
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
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
