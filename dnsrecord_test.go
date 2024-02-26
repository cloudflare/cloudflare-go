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

func TestDNSRecordNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.DNS.Records.New(context.TODO(), cloudflare.DNSRecordNewParams{
		ZoneID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Name:    cloudflare.F("example.com"),
		Type:    cloudflare.F(cloudflare.DNSRecordNewParamsTypeUri),
		Comment: cloudflare.F("Domain verification record"),
		Data: cloudflare.F(cloudflare.DNSRecordNewParamsData{
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
			LatDirection:  cloudflare.F(cloudflare.DNSRecordNewParamsDataLatDirectionN),
			LatMinutes:    cloudflare.F(46.000000),
			LatSeconds:    cloudflare.F(46.000000),
			LongDegrees:   cloudflare.F(122.000000),
			LongDirection: cloudflare.F(cloudflare.DNSRecordNewParamsDataLongDirectionW),
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
		Meta: cloudflare.F(cloudflare.DNSRecordNewParamsMeta{
			AutoAdded: cloudflare.F(true),
			Source:    cloudflare.F("primary"),
		}),
		Priority: cloudflare.F(10.000000),
		Proxied:  cloudflare.F(false),
		Tags:     cloudflare.F([]string{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
		TTL:      cloudflare.F[cloudflare.DNSRecordNewParamsTTL](shared.UnionFloat(3600.000000)),
	})
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
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.DNS.Records.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.DNSRecordUpdateParams{
			ZoneID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:    cloudflare.F("example.com"),
			Type:    cloudflare.F(cloudflare.DNSRecordUpdateParamsTypeUri),
			Comment: cloudflare.F("Domain verification record"),
			Data: cloudflare.F(cloudflare.DNSRecordUpdateParamsData{
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
				LatDirection:  cloudflare.F(cloudflare.DNSRecordUpdateParamsDataLatDirectionN),
				LatMinutes:    cloudflare.F(46.000000),
				LatSeconds:    cloudflare.F(46.000000),
				LongDegrees:   cloudflare.F(122.000000),
				LongDirection: cloudflare.F(cloudflare.DNSRecordUpdateParamsDataLongDirectionW),
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
			Meta: cloudflare.F(cloudflare.DNSRecordUpdateParamsMeta{
				AutoAdded: cloudflare.F(true),
				Source:    cloudflare.F("primary"),
			}),
			Priority: cloudflare.F(10.000000),
			Proxied:  cloudflare.F(false),
			Tags:     cloudflare.F([]string{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
			TTL:      cloudflare.F[cloudflare.DNSRecordUpdateParamsTTL](shared.UnionFloat(3600.000000)),
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

func TestDNSRecordListWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.DNS.Records.List(context.TODO(), cloudflare.DNSRecordListParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Comment: cloudflare.F(cloudflare.DNSRecordListParamsComment{
			Present:    cloudflare.F("string"),
			Absent:     cloudflare.F("string"),
			Exact:      cloudflare.F("Hello, world"),
			Contains:   cloudflare.F("ello, worl"),
			Startswith: cloudflare.F("Hello, w"),
			Endswith:   cloudflare.F("o, world"),
		}),
		Content:   cloudflare.F("127.0.0.1"),
		Direction: cloudflare.F(cloudflare.DNSRecordListParamsDirectionAsc),
		Match:     cloudflare.F(cloudflare.DNSRecordListParamsMatchAny),
		Name:      cloudflare.F("example.com"),
		Order:     cloudflare.F(cloudflare.DNSRecordListParamsOrderType),
		Page:      cloudflare.F(1.000000),
		PerPage:   cloudflare.F(5.000000),
		Proxied:   cloudflare.F(false),
		Search:    cloudflare.F("www.cloudflare.com"),
		Tag: cloudflare.F(cloudflare.DNSRecordListParamsTag{
			Present:    cloudflare.F("important"),
			Absent:     cloudflare.F("important"),
			Exact:      cloudflare.F("greeting:Hello, world"),
			Contains:   cloudflare.F("greeting:ello, worl"),
			Startswith: cloudflare.F("greeting:Hello, w"),
			Endswith:   cloudflare.F("greeting:o, world"),
		}),
		TagMatch: cloudflare.F(cloudflare.DNSRecordListParamsTagMatchAny),
		Type:     cloudflare.F(cloudflare.DNSRecordListParamsTypeA),
	})
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
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.DNS.Records.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.DNSRecordDeleteParams{
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

func TestDNSRecordEditWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.DNS.Records.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.DNSRecordEditParams{
			ZoneID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:    cloudflare.F("example.com"),
			Type:    cloudflare.F(cloudflare.DNSRecordEditParamsTypeUri),
			Comment: cloudflare.F("Domain verification record"),
			Data: cloudflare.F(cloudflare.DNSRecordEditParamsData{
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
				LatDirection:  cloudflare.F(cloudflare.DNSRecordEditParamsDataLatDirectionN),
				LatMinutes:    cloudflare.F(46.000000),
				LatSeconds:    cloudflare.F(46.000000),
				LongDegrees:   cloudflare.F(122.000000),
				LongDirection: cloudflare.F(cloudflare.DNSRecordEditParamsDataLongDirectionW),
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
			Meta: cloudflare.F(cloudflare.DNSRecordEditParamsMeta{
				AutoAdded: cloudflare.F(true),
				Source:    cloudflare.F("primary"),
			}),
			Priority: cloudflare.F(10.000000),
			Proxied:  cloudflare.F(false),
			Tags:     cloudflare.F([]string{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
			TTL:      cloudflare.F[cloudflare.DNSRecordEditParamsTTL](shared.UnionFloat(3600.000000)),
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

func TestDNSRecordExport(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.DNS.Records.Export(context.TODO(), cloudflare.DNSRecordExportParams{
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
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.DNS.Records.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.DNSRecordGetParams{
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

func TestDNSRecordImportWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.DNS.Records.Import(context.TODO(), cloudflare.DNSRecordImportParams{
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

func TestDNSRecordScan(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.DNS.Records.Scan(context.TODO(), cloudflare.DNSRecordScanParams{
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
