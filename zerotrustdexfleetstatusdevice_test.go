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

func TestZeroTrustDEXFleetStatusDeviceListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DEX.FleetStatus.Devices.List(context.TODO(), cloudflare.ZeroTrustDEXFleetStatusDeviceListParams{
		AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
		Page:      cloudflare.F(1.000000),
		PerPage:   cloudflare.F(10.000000),
		TimeEnd:   cloudflare.F("2023-10-11T00:00:00Z"),
		TimeStart: cloudflare.F("2023-10-11T00:00:00Z"),
		Colo:      cloudflare.F("SJC"),
		DeviceID:  cloudflare.F("cb49c27f-7f97-49c5-b6f3-f7c01ead0fd7"),
		Mode:      cloudflare.F("proxy"),
		Platform:  cloudflare.F("windows"),
		SortBy:    cloudflare.F(cloudflare.ZeroTrustDEXFleetStatusDeviceListParamsSortByColo),
		Status:    cloudflare.F("connected"),
		Version:   cloudflare.F("1.0.0"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
