// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestRadarBGPTopAseGetWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.BGP.Top.Ases.Get(context.TODO(), cloudflare.RadarBGPTopAseGetParams{
		ASN:        cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]cloudflare.RadarBGPTopAseGetParamsDateRange{cloudflare.RadarBGPTopAseGetParamsDateRange1d, cloudflare.RadarBGPTopAseGetParamsDateRange2d, cloudflare.RadarBGPTopAseGetParamsDateRange7d}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:     cloudflare.F(cloudflare.RadarBGPTopAseGetParamsFormatJson),
		Limit:      cloudflare.F(int64(5)),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		Prefix:     cloudflare.F([]string{"string", "string", "string"}),
		UpdateType: cloudflare.F([]cloudflare.RadarBGPTopAseGetParamsUpdateType{cloudflare.RadarBGPTopAseGetParamsUpdateTypeAnnouncement, cloudflare.RadarBGPTopAseGetParamsUpdateTypeWithdrawal}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarBGPTopAsePrefixesWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.BGP.Top.Ases.Prefixes(context.TODO(), cloudflare.RadarBGPTopAsePrefixesParams{
		Country: cloudflare.F("NZ"),
		Format:  cloudflare.F(cloudflare.RadarBGPTopAsePrefixesParamsFormatJson),
		Limit:   cloudflare.F(int64(10)),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
