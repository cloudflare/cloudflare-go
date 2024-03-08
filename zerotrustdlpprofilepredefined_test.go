// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestZeroTrustDLPProfilePredefinedUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.ZeroTrust.DLP.Profiles.Predefineds.Update(
		context.TODO(),
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		cloudflare.ZeroTrustDLPProfilePredefinedUpdateParams{
			AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AllowedMatchCount: cloudflare.F(5.000000),
			ContextAwareness: cloudflare.F(cloudflare.ZeroTrustDLPProfilePredefinedUpdateParamsContextAwareness{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(cloudflare.ZeroTrustDLPProfilePredefinedUpdateParamsContextAwarenessSkip{
					Files: cloudflare.F(true),
				}),
			}),
			Entries: cloudflare.F([]cloudflare.ZeroTrustDLPProfilePredefinedUpdateParamsEntry{{
				Enabled: cloudflare.F(true),
			}, {
				Enabled: cloudflare.F(true),
			}, {
				Enabled: cloudflare.F(true),
			}}),
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

func TestZeroTrustDLPProfilePredefinedGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.ZeroTrust.DLP.Profiles.Predefineds.Get(
		context.TODO(),
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		cloudflare.ZeroTrustDLPProfilePredefinedGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
