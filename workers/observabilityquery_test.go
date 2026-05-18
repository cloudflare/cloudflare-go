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

func TestObservabilityQueryNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Observability.Queries.New(context.TODO(), workers.ObservabilityQueryNewParams{
		AccountID:   cloudflare.F("account_id"),
		Description: cloudflare.F("Query description"),
		Name:        cloudflare.F("x"),
		Parameters: cloudflare.F(workers.ObservabilityQueryNewParamsParameters{
			Calculations: cloudflare.F([]workers.ObservabilityQueryNewParamsParametersCalculation{{
				Operator: cloudflare.F(workers.ObservabilityQueryNewParamsParametersCalculationsOperatorUniq),
				Alias:    cloudflare.F("alias"),
				Key:      cloudflare.F("key"),
				KeyType:  cloudflare.F(workers.ObservabilityQueryNewParamsParametersCalculationsKeyTypeString),
			}}),
			Datasets:          cloudflare.F([]string{"string"}),
			FilterCombination: cloudflare.F(workers.ObservabilityQueryNewParamsParametersFilterCombinationAnd),
			Filters: cloudflare.F([]workers.ObservabilityQueryNewParamsParametersFilterUnion{workers.ObservabilityQueryNewParamsParametersFiltersObject{
				FilterCombination: cloudflare.F(workers.ObservabilityQueryNewParamsParametersFiltersObjectFilterCombinationAnd),
				Filters:           cloudflare.F([]interface{}{map[string]interface{}{}}),
				Kind:              cloudflare.F(workers.ObservabilityQueryNewParamsParametersFiltersObjectKindGroup),
			}}),
			GroupBys: cloudflare.F([]workers.ObservabilityQueryNewParamsParametersGroupBy{{
				Type:  cloudflare.F(workers.ObservabilityQueryNewParamsParametersGroupBysTypeString),
				Value: cloudflare.F("value"),
			}}),
			Havings: cloudflare.F([]workers.ObservabilityQueryNewParamsParametersHaving{{
				Key:       cloudflare.F("key"),
				Operation: cloudflare.F(workers.ObservabilityQueryNewParamsParametersHavingsOperationEq),
				Value:     cloudflare.F(0.000000),
			}}),
			Limit: cloudflare.F(int64(0)),
			Needle: cloudflare.F(workers.ObservabilityQueryNewParamsParametersNeedle{
				Value:     cloudflare.F[workers.ObservabilityQueryNewParamsParametersNeedleValueUnion](shared.UnionString("string")),
				IsRegex:   cloudflare.F(true),
				MatchCase: cloudflare.F(true),
			}),
			OrderBy: cloudflare.F(workers.ObservabilityQueryNewParamsParametersOrderBy{
				Value: cloudflare.F("value"),
				Order: cloudflare.F(workers.ObservabilityQueryNewParamsParametersOrderByOrderAsc),
			}),
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

func TestObservabilityQueryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Observability.Queries.List(context.TODO(), workers.ObservabilityQueryListParams{
		AccountID: cloudflare.F("account_id"),
		Order:     cloudflare.F(workers.ObservabilityQueryListParamsOrderAsc),
		OrderBy:   cloudflare.F(workers.ObservabilityQueryListParamsOrderByCreated),
		Page:      cloudflare.F(1.000000),
		PerPage:   cloudflare.F(5.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
