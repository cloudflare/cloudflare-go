// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/ai_gateway"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

func TestDatasetNew(t *testing.T) {
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
	_, err := client.AIGateway.Datasets.New(
		context.TODO(),
		"my-gateway",
		ai_gateway.DatasetNewParams{
			AccountID: cloudflare.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
			Enable:    cloudflare.F(true),
			Filters: cloudflare.F([]ai_gateway.DatasetNewParamsFilter{{
				Key:      cloudflare.F(ai_gateway.DatasetNewParamsFiltersKeyCreatedAt),
				Operator: cloudflare.F(ai_gateway.DatasetNewParamsFiltersOperatorEq),
				Value:    cloudflare.F([]ai_gateway.DatasetNewParamsFiltersValueUnion{shared.UnionString("string")}),
			}}),
			Name: cloudflare.F("name"),
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

func TestDatasetUpdate(t *testing.T) {
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
	_, err := client.AIGateway.Datasets.Update(
		context.TODO(),
		"my-gateway",
		"id",
		ai_gateway.DatasetUpdateParams{
			AccountID: cloudflare.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
			Enable:    cloudflare.F(true),
			Filters: cloudflare.F([]ai_gateway.DatasetUpdateParamsFilter{{
				Key:      cloudflare.F(ai_gateway.DatasetUpdateParamsFiltersKeyCreatedAt),
				Operator: cloudflare.F(ai_gateway.DatasetUpdateParamsFiltersOperatorEq),
				Value:    cloudflare.F([]ai_gateway.DatasetUpdateParamsFiltersValueUnion{shared.UnionString("string")}),
			}}),
			Name: cloudflare.F("name"),
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

func TestDatasetListWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.Datasets.List(
		context.TODO(),
		"my-gateway",
		ai_gateway.DatasetListParams{
			AccountID:        cloudflare.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
			Enable:           cloudflare.F(true),
			Name:             cloudflare.F("name"),
			OrderBy:          cloudflare.F("order_by"),
			OrderByDirection: cloudflare.F(ai_gateway.DatasetListParamsOrderByDirectionAsc),
			Page:             cloudflare.F(int64(1)),
			PerPage:          cloudflare.F(int64(1)),
			Search:           cloudflare.F("search"),
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

func TestDatasetDelete(t *testing.T) {
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
	_, err := client.AIGateway.Datasets.Delete(
		context.TODO(),
		"my-gateway",
		"id",
		ai_gateway.DatasetDeleteParams{
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

func TestDatasetGet(t *testing.T) {
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
	_, err := client.AIGateway.Datasets.Get(
		context.TODO(),
		"my-gateway",
		"id",
		ai_gateway.DatasetGetParams{
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
