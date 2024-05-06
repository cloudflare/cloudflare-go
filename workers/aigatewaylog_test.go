// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/workers"
)

func TestAIGatewayLogGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.AI.Gateways.Logs.Get(
		context.TODO(),
		"0d37909e38d3e99c29fa2cd343ac421a",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		workers.AIGatewayLogGetParams{
			Cached:    cloudflare.F(true),
			Direction: cloudflare.F(workers.AIGatewayLogGetParamsDirectionAsc),
			EndDate:   cloudflare.F(time.Now()),
			OrderBy:   cloudflare.F(workers.AIGatewayLogGetParamsOrderByCreatedAt),
			Page:      cloudflare.F(int64(1)),
			PerPage:   cloudflare.F(int64(5)),
			Search:    cloudflare.F("string"),
			StartDate: cloudflare.F(time.Now()),
			Success:   cloudflare.F(true),
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
