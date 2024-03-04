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

func TestLogReceivedGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Logs.Received.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.LogReceivedGetParams{
			End:        cloudflare.F[cloudflare.LogReceivedGetParamsEnd](shared.UnionString("2018-05-20T10:01:00Z")),
			Count:      cloudflare.F(int64(1)),
			Fields:     cloudflare.F("ClientIP,RayID,EdgeStartTimestamp"),
			Sample:     cloudflare.F(0.100000),
			Start:      cloudflare.F[cloudflare.LogReceivedGetParamsStart](shared.UnionString("2018-05-20T10:00:00Z")),
			Timestamps: cloudflare.F(cloudflare.LogReceivedGetParamsTimestampsUnixnano),
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
