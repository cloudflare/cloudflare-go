// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
)

func TestDLPProfilePredefinedUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Predefined.Update(
		context.TODO(),
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		zero_trust.DLPProfilePredefinedUpdateParams{
			AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AllowedMatchCount: cloudflare.F(5.000000),
			ContextAwareness: cloudflare.F(zero_trust.DLPProfilePredefinedUpdateParamsContextAwareness{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(zero_trust.DLPProfilePredefinedUpdateParamsContextAwarenessSkip{
					Files: cloudflare.F(true),
				}),
			}),
			Entries: cloudflare.F([]zero_trust.DLPProfilePredefinedUpdateParamsEntry{{
				Enabled: cloudflare.F(true),
			}, {
				Enabled: cloudflare.F(true),
			}, {
				Enabled: cloudflare.F(true),
			}}),
			OcrEnabled: cloudflare.F(true),
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

func TestDLPProfilePredefinedGet(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Predefined.Get(
		context.TODO(),
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		zero_trust.DLPProfilePredefinedGetParams{
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
