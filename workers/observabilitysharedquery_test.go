// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/cloudflare/cloudflare-go/v7/workers"
)

func TestObservabilitySharedQueryNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Observability.SharedQueries.New(context.TODO(), workers.ObservabilitySharedQueryNewParams{
		AccountID: cloudflare.F("account_id"),
		QueryID:   cloudflare.F("queryId"),
		Timeframe: cloudflare.F(workers.ObservabilitySharedQueryNewParamsTimeframe{
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
		Parameters: cloudflare.F(workers.ObservabilitySharedQueryNewParamsParameters{
			Calculations: cloudflare.F([]workers.ObservabilitySharedQueryNewParamsParametersCalculation{{
				Operator: cloudflare.F(workers.ObservabilitySharedQueryNewParamsParametersCalculationsOperatorUniq),
				Alias:    cloudflare.F("alias"),
				Key:      cloudflare.F("key"),
				KeyType:  cloudflare.F(workers.ObservabilitySharedQueryNewParamsParametersCalculationsKeyTypeString),
			}}),
			Datasets:          cloudflare.F([]string{"string"}),
			FilterCombination: cloudflare.F(workers.ObservabilitySharedQueryNewParamsParametersFilterCombinationAnd),
			Filters: cloudflare.F([]workers.ObservabilitySharedQueryNewParamsParametersFilterUnion{workers.ObservabilitySharedQueryNewParamsParametersFiltersObject{
				FilterCombination: cloudflare.F(workers.ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombinationAnd),
				Filters: cloudflare.F([]workers.ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterUnion{workers.ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObject{
					FilterCombination: cloudflare.F(workers.ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombinationAnd),
					Filters:           cloudflare.F([]interface{}{map[string]interface{}{}}),
					Kind:              cloudflare.F(workers.ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectKindGroup),
				}}),
				Kind: cloudflare.F(workers.ObservabilitySharedQueryNewParamsParametersFiltersObjectKindGroup),
			}}),
			GroupBys: cloudflare.F([]workers.ObservabilitySharedQueryNewParamsParametersGroupBy{{
				Type:  cloudflare.F(workers.ObservabilitySharedQueryNewParamsParametersGroupBysTypeString),
				Value: cloudflare.F("value"),
			}}),
			Havings: cloudflare.F([]workers.ObservabilitySharedQueryNewParamsParametersHaving{{
				Key:       cloudflare.F("key"),
				Operation: cloudflare.F(workers.ObservabilitySharedQueryNewParamsParametersHavingsOperationEq),
				Value:     cloudflare.F(0.000000),
			}}),
			Limit: cloudflare.F(int64(0)),
			Needle: cloudflare.F(workers.ObservabilitySharedQueryNewParamsParametersNeedle{
				Value:     cloudflare.F[workers.ObservabilitySharedQueryNewParamsParametersNeedleValueUnion](shared.UnionString("string")),
				IsRegex:   cloudflare.F(true),
				MatchCase: cloudflare.F(true),
			}),
			OrderBy: cloudflare.F(workers.ObservabilitySharedQueryNewParamsParametersOrderBy{
				Value: cloudflare.F("value"),
				Order: cloudflare.F(workers.ObservabilitySharedQueryNewParamsParametersOrderByOrderAsc),
			}),
		}),
		View: cloudflare.F(workers.ObservabilitySharedQueryNewParamsViewTraces),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObservabilitySharedQueryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Observability.SharedQueries.Get(
		context.TODO(),
		"id",
		workers.ObservabilitySharedQueryGetParams{
			AccountID: cloudflare.F("account_id"),
			View:      cloudflare.F(workers.ObservabilitySharedQueryGetParamsViewEvents),
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
