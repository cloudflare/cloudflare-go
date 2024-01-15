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

func TestZoneLogReceivedReceivedGetLogsReceivedWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Logs.Receiveds.ReceivedGetLogsReceived(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneLogReceivedReceivedGetLogsReceivedParams{
			End:        cloudflare.F[cloudflare.ZoneLogReceivedReceivedGetLogsReceivedParamsEnd](shared.UnionString("2018-05-20T10:01:00Z")),
			Count:      cloudflare.F(int64(1)),
			Fields:     cloudflare.F("ClientIP,RayID,EdgeStartTimestamp"),
			Sample:     cloudflare.F(0.100000),
			Start:      cloudflare.F[cloudflare.ZoneLogReceivedReceivedGetLogsReceivedParamsStart](shared.UnionString("2018-05-20T10:00:00Z")),
			Timestamps: cloudflare.F(cloudflare.ZoneLogReceivedReceivedGetLogsReceivedParamsTimestampsUnixnano),
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
