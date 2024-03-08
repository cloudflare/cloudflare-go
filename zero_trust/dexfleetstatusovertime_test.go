// File generated from our OpenAPI spec by Stainless.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
)

func TestDEXFleetStatusOverTimeListWithOptionalParams(t *testing.T) {
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
	err := client.ZeroTrust.DEX.FleetStatus.OverTime.List(context.TODO(), zero_trust.DEXFleetStatusOverTimeListParams{
		AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
		TimeEnd:   cloudflare.F("2023-10-11T00:00:00Z"),
		TimeStart: cloudflare.F("2023-10-11T00:00:00Z"),
		Colo:      cloudflare.F("SJC"),
		DeviceID:  cloudflare.F("cb49c27f-7f97-49c5-b6f3-f7c01ead0fd7"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
