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

func TestObservabilityTelemetryQueryNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Observability.Telemetry.Query.New(context.TODO(), workers.ObservabilityTelemetryQueryNewParams{
		AccountID: cloudflare.F("account_id"),
		QueryID:   cloudflare.F("queryId"),
		Timeframe: cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsTimeframe{
			From: cloudflare.F(0.000000),
			To:   cloudflare.F(0.000000),
		}),
		Chart:           cloudflare.F(true),
		Compare:         cloudflare.F(true),
		Dry:             cloudflare.F(true),
		Granularity:     cloudflare.F(0.000000),
		IgnoreSeries:    cloudflare.F(true),
		Limit:           cloudflare.F(100.000000),
		Offset:          cloudflare.F("offset"),
		OffsetBy:        cloudflare.F(0.000000),
		OffsetDirection: cloudflare.F("offsetDirection"),
		Parameters: cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsParameters{
			Calculations: cloudflare.F([]workers.ObservabilityTelemetryQueryNewParamsParametersCalculation{{
				Operator: cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorUniq),
				Alias:    cloudflare.F("alias"),
				Key:      cloudflare.F("key"),
				KeyType:  cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsParametersCalculationsKeyTypeString),
			}}),
			Datasets:          cloudflare.F([]string{"string"}),
			FilterCombination: cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsParametersFilterCombinationAnd),
			Filters: cloudflare.F([]workers.ObservabilityTelemetryQueryNewParamsParametersFilter{{
				Key:       cloudflare.F("key"),
				Operation: cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsParametersFiltersOperationIncludes),
				Type:      cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsParametersFiltersTypeString),
				Value:     cloudflare.F[workers.ObservabilityTelemetryQueryNewParamsParametersFiltersValueUnion](shared.UnionString("string")),
			}}),
			GroupBys: cloudflare.F([]workers.ObservabilityTelemetryQueryNewParamsParametersGroupBy{{
				Type:  cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsParametersGroupBysTypeString),
				Value: cloudflare.F("value"),
			}}),
			Havings: cloudflare.F([]workers.ObservabilityTelemetryQueryNewParamsParametersHaving{{
				Key:       cloudflare.F("key"),
				Operation: cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsParametersHavingsOperationEq),
				Value:     cloudflare.F(0.000000),
			}}),
			Limit: cloudflare.F(int64(0)),
			Needle: cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsParametersNeedle{
				Value:     cloudflare.F[workers.ObservabilityTelemetryQueryNewParamsParametersNeedleValueUnion](shared.UnionString("string")),
				IsRegex:   cloudflare.F(true),
				MatchCase: cloudflare.F(true),
			}),
			OrderBy: cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsParametersOrderBy{
				Value: cloudflare.F("value"),
				Order: cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsParametersOrderByOrderAsc),
			}),
		}),
		PatternType: cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsPatternTypeMessage),
		View:        cloudflare.F(workers.ObservabilityTelemetryQueryNewParamsViewTraces),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
