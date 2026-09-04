// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/browser_rendering"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestDevtoolBrowserLiveViewNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.BrowserRendering.Devtools.Browser.LiveView.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		browser_rendering.DevtoolBrowserLiveViewNewParams{
			AccountID:   cloudflare.F("account_id"),
			ExpiresInMs: cloudflare.F(60000.000000),
			Guardrails: cloudflare.F(browser_rendering.DevtoolBrowserLiveViewNewParamsGuardrails{
				Mode: cloudflare.F(browser_rendering.DevtoolBrowserLiveViewNewParamsGuardrailsModeReadonly),
			}),
			Mode:     cloudflare.F(browser_rendering.DevtoolBrowserLiveViewNewParamsModeDevtools),
			TargetID: cloudflare.F("targetId"),
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
