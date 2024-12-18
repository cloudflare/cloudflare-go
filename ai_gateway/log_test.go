// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/ai_gateway"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

func TestLogListWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.Logs.List(
		context.TODO(),
		"my-gateway",
		ai_gateway.LogListParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
			Cached:    cloudflare.F(true),
			Direction: cloudflare.F(ai_gateway.LogListParamsDirectionAsc),
			EndDate:   cloudflare.F(time.Now()),
			Feedback:  cloudflare.F(ai_gateway.LogListParamsFeedback0),
			Filters: cloudflare.F([]ai_gateway.LogListParamsFilter{{
				Key:      cloudflare.F(ai_gateway.LogListParamsFiltersKeyID),
				Operator: cloudflare.F(ai_gateway.LogListParamsFiltersOperatorEq),
				Value:    cloudflare.F([]ai_gateway.LogListParamsFiltersValueUnion{shared.UnionString("string")}),
			}}),
			MaxCost:             cloudflare.F(0.000000),
			MaxDuration:         cloudflare.F(0.000000),
			MaxTokensIn:         cloudflare.F(0.000000),
			MaxTokensOut:        cloudflare.F(0.000000),
			MaxTotalTokens:      cloudflare.F(0.000000),
			MetaInfo:            cloudflare.F(true),
			MinCost:             cloudflare.F(0.000000),
			MinDuration:         cloudflare.F(0.000000),
			MinTokensIn:         cloudflare.F(0.000000),
			MinTokensOut:        cloudflare.F(0.000000),
			MinTotalTokens:      cloudflare.F(0.000000),
			Model:               cloudflare.F("model"),
			ModelType:           cloudflare.F("model_type"),
			OrderBy:             cloudflare.F(ai_gateway.LogListParamsOrderByCreatedAt),
			OrderByDirection:    cloudflare.F(ai_gateway.LogListParamsOrderByDirectionAsc),
			Page:                cloudflare.F(int64(1)),
			PerPage:             cloudflare.F(int64(1)),
			Provider:            cloudflare.F("provider"),
			RequestContentType:  cloudflare.F("request_content_type"),
			ResponseContentType: cloudflare.F("response_content_type"),
			Search:              cloudflare.F("search"),
			StartDate:           cloudflare.F(time.Now()),
			Success:             cloudflare.F(true),
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

func TestLogDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.Logs.Delete(
		context.TODO(),
		"my-gateway",
		ai_gateway.LogDeleteParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
			Filters: cloudflare.F([]ai_gateway.LogDeleteParamsFilter{{
				Key:      cloudflare.F(ai_gateway.LogDeleteParamsFiltersKeyID),
				Operator: cloudflare.F(ai_gateway.LogDeleteParamsFiltersOperatorEq),
				Value:    cloudflare.F([]ai_gateway.LogDeleteParamsFiltersValueUnion{shared.UnionString("string")}),
			}}),
			Limit:            cloudflare.F(int64(1)),
			OrderBy:          cloudflare.F(ai_gateway.LogDeleteParamsOrderByCreatedAt),
			OrderByDirection: cloudflare.F(ai_gateway.LogDeleteParamsOrderByDirectionAsc),
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

func TestLogEditWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.Logs.Edit(
		context.TODO(),
		"my-gateway",
		"id",
		ai_gateway.LogEditParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
			Feedback:  cloudflare.F(-1.000000),
			Metadata: cloudflare.F(map[string]ai_gateway.LogEditParamsMetadataUnion{
				"foo": shared.UnionString("string"),
			}),
			Score: cloudflare.F(0.000000),
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

func TestLogGet(t *testing.T) {
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
	_, err := client.AIGateway.Logs.Get(
		context.TODO(),
		"my-gateway",
		"id",
		ai_gateway.LogGetParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
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

func TestLogRequest(t *testing.T) {
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
	_, err := client.AIGateway.Logs.Request(
		context.TODO(),
		"my-gateway",
		"id",
		ai_gateway.LogRequestParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
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

func TestLogResponse(t *testing.T) {
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
	_, err := client.AIGateway.Logs.Response(
		context.TODO(),
		"my-gateway",
		"id",
		ai_gateway.LogResponseParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
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
