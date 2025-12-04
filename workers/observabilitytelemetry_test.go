// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
	"github.com/cloudflare/cloudflare-go/v6/workers"
)

func TestObservabilityTelemetryKeysWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Observability.Telemetry.Keys(context.TODO(), workers.ObservabilityTelemetryKeysParams{
		AccountID: cloudflare.F("account_id"),
		Datasets:  cloudflare.F([]string{"string"}),
		Filters: cloudflare.F([]workers.ObservabilityTelemetryKeysParamsFilter{{
			Key:       cloudflare.F("key"),
			Operation: cloudflare.F(workers.ObservabilityTelemetryKeysParamsFiltersOperationIncludes),
			Type:      cloudflare.F(workers.ObservabilityTelemetryKeysParamsFiltersTypeString),
			Value:     cloudflare.F[workers.ObservabilityTelemetryKeysParamsFiltersValueUnion](shared.UnionString("string")),
		}}),
		From: cloudflare.F(0.000000),
		KeyNeedle: cloudflare.F(workers.ObservabilityTelemetryKeysParamsKeyNeedle{
			Value:     cloudflare.F[workers.ObservabilityTelemetryKeysParamsKeyNeedleValueUnion](shared.UnionString("string")),
			IsRegex:   cloudflare.F(true),
			MatchCase: cloudflare.F(true),
		}),
		Limit: cloudflare.F(0.000000),
		Needle: cloudflare.F(workers.ObservabilityTelemetryKeysParamsNeedle{
			Value:     cloudflare.F[workers.ObservabilityTelemetryKeysParamsNeedleValueUnion](shared.UnionString("string")),
			IsRegex:   cloudflare.F(true),
			MatchCase: cloudflare.F(true),
		}),
		To: cloudflare.F(0.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservabilityTelemetryQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Observability.Telemetry.Query(context.TODO(), workers.ObservabilityTelemetryQueryParams{
		AccountID: cloudflare.F("account_id"),
		QueryID:   cloudflare.F("queryId"),
		Timeframe: cloudflare.F(workers.ObservabilityTelemetryQueryParamsTimeframe{
			From: cloudflare.F(0.000000),
			To:   cloudflare.F(0.000000),
		}),
		Chart:           cloudflare.F(true),
		Compare:         cloudflare.F(true),
		Dry:             cloudflare.F(true),
		Granularity:     cloudflare.F(0.000000),
		IgnoreSeries:    cloudflare.F(true),
		Limit:           cloudflare.F(2000.000000),
		Offset:          cloudflare.F("offset"),
		OffsetBy:        cloudflare.F(0.000000),
		OffsetDirection: cloudflare.F("offsetDirection"),
		Parameters: cloudflare.F(workers.ObservabilityTelemetryQueryParamsParameters{
			Calculations: cloudflare.F([]workers.ObservabilityTelemetryQueryParamsParametersCalculation{{
				Operator: cloudflare.F(workers.ObservabilityTelemetryQueryParamsParametersCalculationsOperatorUniq),
				Alias:    cloudflare.F("alias"),
				Key:      cloudflare.F("key"),
				KeyType:  cloudflare.F(workers.ObservabilityTelemetryQueryParamsParametersCalculationsKeyTypeString),
			}}),
			Datasets:          cloudflare.F([]string{"string"}),
			FilterCombination: cloudflare.F(workers.ObservabilityTelemetryQueryParamsParametersFilterCombinationAnd),
			Filters: cloudflare.F([]workers.ObservabilityTelemetryQueryParamsParametersFilter{{
				Key:       cloudflare.F("key"),
				Operation: cloudflare.F(workers.ObservabilityTelemetryQueryParamsParametersFiltersOperationIncludes),
				Type:      cloudflare.F(workers.ObservabilityTelemetryQueryParamsParametersFiltersTypeString),
				Value:     cloudflare.F[workers.ObservabilityTelemetryQueryParamsParametersFiltersValueUnion](shared.UnionString("string")),
			}}),
			GroupBys: cloudflare.F([]workers.ObservabilityTelemetryQueryParamsParametersGroupBy{{
				Type:  cloudflare.F(workers.ObservabilityTelemetryQueryParamsParametersGroupBysTypeString),
				Value: cloudflare.F("value"),
			}}),
			Havings: cloudflare.F([]workers.ObservabilityTelemetryQueryParamsParametersHaving{{
				Key:       cloudflare.F("key"),
				Operation: cloudflare.F(workers.ObservabilityTelemetryQueryParamsParametersHavingsOperationEq),
				Value:     cloudflare.F(0.000000),
			}}),
			Limit: cloudflare.F(int64(0)),
			Needle: cloudflare.F(workers.ObservabilityTelemetryQueryParamsParametersNeedle{
				Value:     cloudflare.F[workers.ObservabilityTelemetryQueryParamsParametersNeedleValueUnion](shared.UnionString("string")),
				IsRegex:   cloudflare.F(true),
				MatchCase: cloudflare.F(true),
			}),
			OrderBy: cloudflare.F(workers.ObservabilityTelemetryQueryParamsParametersOrderBy{
				Value: cloudflare.F("value"),
				Order: cloudflare.F(workers.ObservabilityTelemetryQueryParamsParametersOrderByOrderAsc),
			}),
		}),
		PatternType: cloudflare.F(workers.ObservabilityTelemetryQueryParamsPatternTypeMessage),
		View:        cloudflare.F(workers.ObservabilityTelemetryQueryParamsViewTraces),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservabilityTelemetryValuesWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Observability.Telemetry.Values(context.TODO(), workers.ObservabilityTelemetryValuesParams{
		AccountID: cloudflare.F("account_id"),
		Datasets:  cloudflare.F([]string{"string"}),
		Key:       cloudflare.F("key"),
		Timeframe: cloudflare.F(workers.ObservabilityTelemetryValuesParamsTimeframe{
			From: cloudflare.F(0.000000),
			To:   cloudflare.F(0.000000),
		}),
		Type: cloudflare.F(workers.ObservabilityTelemetryValuesParamsTypeString),
		Filters: cloudflare.F([]workers.ObservabilityTelemetryValuesParamsFilter{{
			Key:       cloudflare.F("key"),
			Operation: cloudflare.F(workers.ObservabilityTelemetryValuesParamsFiltersOperationIncludes),
			Type:      cloudflare.F(workers.ObservabilityTelemetryValuesParamsFiltersTypeString),
			Value:     cloudflare.F[workers.ObservabilityTelemetryValuesParamsFiltersValueUnion](shared.UnionString("string")),
		}}),
		Limit: cloudflare.F(0.000000),
		Needle: cloudflare.F(workers.ObservabilityTelemetryValuesParamsNeedle{
			Value:     cloudflare.F[workers.ObservabilityTelemetryValuesParamsNeedleValueUnion](shared.UnionString("string")),
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
