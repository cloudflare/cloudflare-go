// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/workers"
)

func TestScriptVersionNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Workers.Scripts.Versions.New(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptVersionNewParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Metadata: cloudflare.F(workers.ScriptVersionNewParamsMetadata{
				MainModule: cloudflare.F("worker.js"),
				Annotations: cloudflare.F(workers.ScriptVersionNewParamsMetadataAnnotations{
					WorkersMessage: cloudflare.F("Fixed worker code."),
					WorkersTag:     cloudflare.F("workers/tag"),
				}),
				Bindings: cloudflare.F([]workers.ScriptVersionNewParamsMetadataBindingUnion{workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAny{
					Name: cloudflare.F("MY_ENV_VAR"),
					Type: cloudflare.F("plain_text"),
				}}),
				CompatibilityDate:  cloudflare.F("2021-01-01"),
				CompatibilityFlags: cloudflare.F([]string{"nodejs_compat"}),
				KeepBindings:       cloudflare.F([]string{"string"}),
				UsageModel:         cloudflare.F(workers.ScriptVersionNewParamsMetadataUsageModelStandard),
			}),
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

func TestScriptVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Scripts.Versions.List(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptVersionListParams{
			AccountID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Deployable: cloudflare.F(true),
			Page:       cloudflare.F(int64(0)),
			PerPage:    cloudflare.F(int64(0)),
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

func TestScriptVersionGet(t *testing.T) {
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
	_, err := client.Workers.Scripts.Versions.Get(
		context.TODO(),
		"this-is_my_script-01",
		"bcf48806-b317-4351-9ee7-36e7d557d4de",
		workers.ScriptVersionGetParams{
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
