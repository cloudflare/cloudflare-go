// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestZeroTrustDEXTracerouteTestGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DEX.TracerouteTests.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDEXTracerouteTestGetParams{
			AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
			Interval:  cloudflare.F(cloudflare.ZeroTrustDEXTracerouteTestGetParamsIntervalMinute),
			TimeEnd:   cloudflare.F("string"),
			TimeStart: cloudflare.F("string"),
			Colo:      cloudflare.F("string"),
			DeviceID:  cloudflare.F([]string{"string", "string", "string"}),
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

func TestZeroTrustDEXTracerouteTestNetworkPath(t *testing.T) {
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
	_, err := client.ZeroTrust.DEX.TracerouteTests.NetworkPath(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDEXTracerouteTestNetworkPathParams{
			AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
			DeviceID:  cloudflare.F("string"),
			Interval:  cloudflare.F(cloudflare.ZeroTrustDEXTracerouteTestNetworkPathParamsIntervalMinute),
			TimeEnd:   cloudflare.F("string"),
			TimeStart: cloudflare.F("string"),
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

func TestZeroTrustDEXTracerouteTestPercentilesWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DEX.TracerouteTests.Percentiles(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDEXTracerouteTestPercentilesParams{
			AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
			TimeEnd:   cloudflare.F("2023-09-20T17:00:00Z"),
			TimeStart: cloudflare.F("2023-09-20T17:00:00Z"),
			Colo:      cloudflare.F("string"),
			DeviceID:  cloudflare.F([]string{"string", "string", "string"}),
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
