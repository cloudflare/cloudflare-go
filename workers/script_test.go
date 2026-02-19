// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/workers"
)

func TestScriptUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Scripts.Update(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Metadata: cloudflare.F(workers.ScriptUpdateParamsMetadata{
				Annotations: cloudflare.F(workers.ScriptUpdateParamsMetadataAnnotations{
					WorkersMessage: cloudflare.F("Fixed bug."),
					WorkersTag:     cloudflare.F("v1.0.1"),
				}),
				Assets: cloudflare.F(workers.ScriptUpdateParamsMetadataAssets{
					Config: cloudflare.F(workers.ScriptUpdateParamsMetadataAssetsConfig{
						Headers:          cloudflare.F("/dashboard/*\nX-Frame-Options: DENY\n\n/static/*\nAccess-Control-Allow-Origin: *"),
						Redirects:        cloudflare.F("/foo /bar 301\n/news/* /blog/:splat"),
						HTMLHandling:     cloudflare.F(workers.ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash),
						NotFoundHandling: cloudflare.F(workers.ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling404Page),
						RunWorkerFirst:   cloudflare.F[workers.ScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion](workers.ScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstArray([]string{"string"})),
						ServeDirectly:    cloudflare.F(true),
					}),
					JWT: cloudflare.F("jwt"),
				}),
				Bindings: cloudflare.F([]workers.ScriptUpdateParamsMetadataBindingUnion{workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText{
					Name: cloudflare.F("MY_ENV_VAR"),
					Text: cloudflare.F("my_data"),
					Type: cloudflare.F(workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText),
				}}),
				BodyPart:           cloudflare.F("worker.js"),
				CompatibilityDate:  cloudflare.F("2021-01-01"),
				CompatibilityFlags: cloudflare.F([]string{"nodejs_compat"}),
				KeepAssets:         cloudflare.F(false),
				KeepBindings:       cloudflare.F([]string{"string"}),
				Limits: cloudflare.F(workers.ScriptUpdateParamsMetadataLimits{
					CPUMs: cloudflare.F(int64(50)),
				}),
				Logpush:    cloudflare.F(false),
				MainModule: cloudflare.F("worker.js"),
				Migrations: cloudflare.F[workers.ScriptUpdateParamsMetadataMigrationsUnion](workers.SingleStepMigrationParam{
					DeletedClasses:   cloudflare.F([]string{"string"}),
					NewClasses:       cloudflare.F([]string{"string"}),
					NewSqliteClasses: cloudflare.F([]string{"string"}),
					NewTag:           cloudflare.F("v2"),
					OldTag:           cloudflare.F("v1"),
					RenamedClasses: cloudflare.F([]workers.SingleStepMigrationRenamedClassParam{{
						From: cloudflare.F("from"),
						To:   cloudflare.F("to"),
					}}),
					TransferredClasses: cloudflare.F([]workers.SingleStepMigrationTransferredClassParam{{
						From:       cloudflare.F("from"),
						FromScript: cloudflare.F("from_script"),
						To:         cloudflare.F("to"),
					}}),
				}),
				Observability: cloudflare.F(workers.ScriptUpdateParamsMetadataObservability{
					Enabled:          cloudflare.F(true),
					HeadSamplingRate: cloudflare.F(0.100000),
					Logs: cloudflare.F(workers.ScriptUpdateParamsMetadataObservabilityLogs{
						Enabled:          cloudflare.F(true),
						InvocationLogs:   cloudflare.F(true),
						Destinations:     cloudflare.F([]string{"cloudflare"}),
						HeadSamplingRate: cloudflare.F(0.100000),
						Persist:          cloudflare.F(true),
					}),
				}),
				Placement: cloudflare.F[workers.ScriptUpdateParamsMetadataPlacementUnion](workers.ScriptUpdateParamsMetadataPlacementObject{
					Mode: cloudflare.F(workers.ScriptUpdateParamsMetadataPlacementObjectModeSmart),
				}),
				Tags: cloudflare.F([]string{"string"}),
				TailConsumers: cloudflare.F([]workers.ConsumerScriptParam{{
					Service:     cloudflare.F("my-log-consumer"),
					Environment: cloudflare.F("production"),
					Namespace:   cloudflare.F("my-namespace"),
				}}),
				UsageModel: cloudflare.F(workers.ScriptUpdateParamsMetadataUsageModelStandard),
			}),
			Files: cloudflare.F([]io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents")))}),
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

func TestScriptListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Scripts.List(context.TODO(), workers.ScriptListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Tags:      cloudflare.F("production:yes,staging:no"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScriptDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Scripts.Delete(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Force:     cloudflare.F(true),
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

func TestScriptGet(t *testing.T) {
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
	_, err := client.Workers.Scripts.Get(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptGetParams{
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

func TestScriptSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Scripts.Search(context.TODO(), workers.ScriptSearchParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ID:        cloudflare.F("bdf3567828824b74aadd550004cf4913"),
		Name:      cloudflare.F("my-worker"),
		OrderBy:   cloudflare.F(workers.ScriptSearchParamsOrderByCreatedOn),
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(1)),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
