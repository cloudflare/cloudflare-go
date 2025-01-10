// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/ai"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestModelListWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Models.List(context.TODO(), ai.ModelListParams{
		AccountID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Author:           cloudflare.F("author"),
		HideExperimental: cloudflare.F(true),
		Page:             cloudflare.F(int64(0)),
		PerPage:          cloudflare.F(int64(0)),
		Search:           cloudflare.F("search"),
		Source:           cloudflare.F(0.000000),
		Task:             cloudflare.F("Text Generation"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
