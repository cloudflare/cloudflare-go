// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/cloudflare/cloudflare-go/v4/workers"
)

func TestObservabilityTelemetryKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Observability.Telemetry.Keys.New(context.TODO(), workers.ObservabilityTelemetryKeyNewParams{
		AccountID: cloudflare.F("account_id"),
		Datasets:  cloudflare.F([]string{"string"}),
		Filters: cloudflare.F([]workers.ObservabilityTelemetryKeyNewParamsFilter{{
			Key:       cloudflare.F("key"),
			Operation: cloudflare.F(workers.ObservabilityTelemetryKeyNewParamsFiltersOperationIncludes),
			Type:      cloudflare.F(workers.ObservabilityTelemetryKeyNewParamsFiltersTypeString),
			Value:     cloudflare.F[workers.ObservabilityTelemetryKeyNewParamsFiltersValueUnion](shared.UnionString("string")),
		}}),
		KeyNeedle: cloudflare.F(workers.ObservabilityTelemetryKeyNewParamsKeyNeedle{
			Value:     cloudflare.F[workers.ObservabilityTelemetryKeyNewParamsKeyNeedleValueUnion](shared.UnionString("string")),
			IsRegex:   cloudflare.F(true),
			MatchCase: cloudflare.F(true),
		}),
		Limit: cloudflare.F(0.000000),
		Needle: cloudflare.F(workers.ObservabilityTelemetryKeyNewParamsNeedle{
			Value:     cloudflare.F[workers.ObservabilityTelemetryKeyNewParamsNeedleValueUnion](shared.UnionString("string")),
			IsRegex:   cloudflare.F(true),
			MatchCase: cloudflare.F(true),
		}),
		Timeframe: cloudflare.F(workers.ObservabilityTelemetryKeyNewParamsTimeframe{
			From: cloudflare.F(0.000000),
			To:   cloudflare.F(0.000000),
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
