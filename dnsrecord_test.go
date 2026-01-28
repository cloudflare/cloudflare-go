// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

func TestDNSRecordListWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Records.List(context.TODO(), cloudflare.DNSRecordListParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Comment: cloudflare.F(cloudflare.DNSRecordListParamsComment{
			Absent:     cloudflare.F("absent"),
			Contains:   cloudflare.F("ello, worl"),
			Endswith:   cloudflare.F("o, world"),
			Exact:      cloudflare.F("Hello, world"),
			Present:    cloudflare.F("present"),
			Startswith: cloudflare.F("Hello, w"),
		}),
		Content: cloudflare.F(cloudflare.DNSRecordListParamsContent{
			Contains:   cloudflare.F("7.0.0."),
			Endswith:   cloudflare.F(".0.1"),
			Exact:      cloudflare.F("127.0.0.1"),
			Startswith: cloudflare.F("127.0."),
		}),
		Direction: cloudflare.F(cloudflare.DNSRecordListParamsDirectionAsc),
		Match:     cloudflare.F(cloudflare.DNSRecordListParamsMatchAny),
		Name: cloudflare.F(cloudflare.DNSRecordListParamsName{
			Contains:   cloudflare.F("w.example."),
			Endswith:   cloudflare.F(".example.com"),
			Exact:      cloudflare.F("www.example.com"),
			Startswith: cloudflare.F("www.example"),
		}),
		Order:   cloudflare.F(cloudflare.DNSRecordListParamsOrderType),
		Page:    cloudflare.F(1.000000),
		PerPage: cloudflare.F(5.000000),
		Proxied: cloudflare.F(true),
		Search:  cloudflare.F("www.cloudflare.com"),
		Tag: cloudflare.F(cloudflare.DNSRecordListParamsTag{
			Absent:     cloudflare.F("important"),
			Contains:   cloudflare.F("greeting:ello, worl"),
			Endswith:   cloudflare.F("greeting:o, world"),
			Exact:      cloudflare.F("greeting:Hello, world"),
			Present:    cloudflare.F("important"),
			Startswith: cloudflare.F("greeting:Hello, w"),
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
