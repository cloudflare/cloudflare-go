// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/ai_gateway"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
			AccountID:           cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
			Cached:              cloudflare.F(true),
			Direction:           cloudflare.F(ai_gateway.LogListParamsDirectionAsc),
			EndDate:             cloudflare.F(time.Now()),
			Feedback:            cloudflare.F(ai_gateway.LogListParamsFeedback0),
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
			PerPage:             cloudflare.F(int64(5)),
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
