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

func TestObservabilityTelemetryValueNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Observability.Telemetry.Values.New(context.TODO(), workers.ObservabilityTelemetryValueNewParams{
		AccountID: cloudflare.F("account_id"),
		Datasets:  cloudflare.F([]string{"string"}),
		Key:       cloudflare.F("key"),
		Timeframe: cloudflare.F(workers.ObservabilityTelemetryValueNewParamsTimeframe{
			From: cloudflare.F(0.000000),
			To:   cloudflare.F(0.000000),
		}),
		Type: cloudflare.F(workers.ObservabilityTelemetryValueNewParamsTypeString),
		Filters: cloudflare.F([]workers.ObservabilityTelemetryValueNewParamsFilter{{
			Key:       cloudflare.F("key"),
			Operation: cloudflare.F(workers.ObservabilityTelemetryValueNewParamsFiltersOperationIncludes),
			Type:      cloudflare.F(workers.ObservabilityTelemetryValueNewParamsFiltersTypeString),
			Value:     cloudflare.F[workers.ObservabilityTelemetryValueNewParamsFiltersValueUnion](shared.UnionString("string")),
		}}),
		Limit: cloudflare.F(0.000000),
		Needle: cloudflare.F(workers.ObservabilityTelemetryValueNewParamsNeedle{
			Value:     cloudflare.F[workers.ObservabilityTelemetryValueNewParamsNeedleValueUnion](shared.UnionString("string")),
			IsRegex:   cloudflare.F(true),
			MatchCase: cloudflare.F(true),
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
