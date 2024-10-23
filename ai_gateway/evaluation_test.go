// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/ai_gateway"
	"github.com/cloudflare/cloudflare-go/v3/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v3/option"
)

func TestEvaluationNew(t *testing.T) {
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
	_, err := client.AIGateway.Evaluations.New(
		context.TODO(),
		"gateway_id",
		ai_gateway.EvaluationNewParams{
			AccountID:         cloudflare.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
			DatasetIDs:        cloudflare.F([]string{"string", "string", "string"}),
			EvaluationTypeIDs: cloudflare.F([]string{"string", "string", "string"}),
			Name:              cloudflare.F("name"),
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

func TestEvaluationListWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.Evaluations.List(
		context.TODO(),
		"gateway_id",
		ai_gateway.EvaluationListParams{
			AccountID:        cloudflare.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
			ID:               cloudflare.F("id"),
			Name:             cloudflare.F("name"),
			OrderBy:          cloudflare.F("order_by"),
			OrderByDirection: cloudflare.F(ai_gateway.EvaluationListParamsOrderByDirectionAsc),
			Page:             cloudflare.F(int64(1)),
			PerPage:          cloudflare.F(int64(5)),
			Processed:        cloudflare.F(true),
			Search:           cloudflare.F("search"),
			TotalLogs:        cloudflare.F(0.000000),
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

func TestEvaluationDelete(t *testing.T) {
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
	_, err := client.AIGateway.Evaluations.Delete(
		context.TODO(),
		"gateway_id",
		"id",
		ai_gateway.EvaluationDeleteParams{
			AccountID: cloudflare.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
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

func TestEvaluationGet(t *testing.T) {
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
	_, err := client.AIGateway.Evaluations.Get(
		context.TODO(),
		"gateway_id",
		"id",
		ai_gateway.EvaluationGetParams{
			AccountID: cloudflare.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
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
