// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/radar"
)

func TestAS112TopDNSSECWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AS112.Top.DNSSEC(
		context.TODO(),
		radar.AS112TopDNSSECParamsDNSSECSupported,
		radar.AS112TopDNSSECParams{
			ASN:       cloudflare.F([]string{"string"}),
			Continent: cloudflare.F([]string{"string"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now()}),
			DateRange: cloudflare.F([]string{"7d"}),
			DateStart: cloudflare.F([]time.Time{time.Now()}),
			Format:    cloudflare.F(radar.AS112TopDNSSECParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"string"}),
			Name:      cloudflare.F([]string{"string"}),
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

func TestAS112TopEdnsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AS112.Top.Edns(
		context.TODO(),
		radar.AS112TopEdnsParamsEdnsSupported,
		radar.AS112TopEdnsParams{
			ASN:       cloudflare.F([]string{"string"}),
			Continent: cloudflare.F([]string{"string"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now()}),
			DateRange: cloudflare.F([]string{"7d"}),
			DateStart: cloudflare.F([]time.Time{time.Now()}),
			Format:    cloudflare.F(radar.AS112TopEdnsParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"string"}),
			Name:      cloudflare.F([]string{"string"}),
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

func TestAS112TopIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AS112.Top.IPVersion(
		context.TODO(),
		radar.AS112TopIPVersionParamsIPVersionIPv4,
		radar.AS112TopIPVersionParams{
			ASN:       cloudflare.F([]string{"string"}),
			Continent: cloudflare.F([]string{"string"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now()}),
			DateRange: cloudflare.F([]string{"7d"}),
			DateStart: cloudflare.F([]time.Time{time.Now()}),
			Format:    cloudflare.F(radar.AS112TopIPVersionParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"string"}),
			Name:      cloudflare.F([]string{"string"}),
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

func TestAS112TopLocationsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AS112.Top.Locations(context.TODO(), radar.AS112TopLocationsParams{
		ASN:       cloudflare.F([]string{"string"}),
		Continent: cloudflare.F([]string{"string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		Format:    cloudflare.F(radar.AS112TopLocationsParamsFormatJson),
		Limit:     cloudflare.F(int64(5)),
		Location:  cloudflare.F([]string{"string"}),
		Name:      cloudflare.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
