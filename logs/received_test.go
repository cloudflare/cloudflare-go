// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/logs"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

func TestReceivedGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Logs.Received.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		logs.ReceivedGetParams{
			End:        cloudflare.F[logs.ReceivedGetParamsEndUnion](shared.UnionString("2018-05-20T10:01:00Z")),
			Count:      cloudflare.F(int64(1)),
			Fields:     cloudflare.F("ClientIP,RayID,EdgeStartTimestamp"),
			Sample:     cloudflare.F(0.100000),
			Start:      cloudflare.F[logs.ReceivedGetParamsStartUnion](shared.UnionString("2018-05-20T10:00:00Z")),
			Timestamps: cloudflare.F(logs.ReceivedGetParamsTimestampsUnixnano),
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
