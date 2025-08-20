// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/workers"
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
			Metadata: cloudflare.F[workers.ScriptUpdateParamsMetadataUnion](workers.ScriptUpdateParamsMetadataObject{
				MainModule: cloudflare.F("worker.js"),
				Assets: cloudflare.F(workers.ScriptUpdateParamsMetadataObjectAssets{
					Config: cloudflare.F(workers.ScriptUpdateParamsMetadataObjectAssetsConfig{
						Headers:          cloudflare.F("/dashboard/*\nX-Frame-Options: DENY\n\n/static/*\nAccess-Control-Allow-Origin: *"),
						Redirects:        cloudflare.F("/foo /bar 301\n/news/* /blog/:splat"),
						HTMLHandling:     cloudflare.F(workers.ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingAutoTrailingSlash),
						NotFoundHandling: cloudflare.F(workers.ScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling404Page),
						RunWorkerFirst:   cloudflare.F[workers.ScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstUnion](workers.ScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstArray([]string{"string"})),
						ServeDirectly:    cloudflare.F(true),
					}),
					JWT: cloudflare.F("jwt"),
				}),
				Bindings: cloudflare.F([]workers.ScriptUpdateParamsMetadataObjectBindingUnion{workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainText{
					Name: cloudflare.F("MY_ENV_VAR"),
					Text: cloudflare.F("my_data"),
					Type: cloudflare.F(workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainTextTypePlainText),
				}}),
				CompatibilityDate:  cloudflare.F("2021-01-01"),
				CompatibilityFlags: cloudflare.F([]string{"nodejs_compat"}),
				KeepAssets:         cloudflare.F(false),
				KeepBindings:       cloudflare.F([]string{"string"}),
				Logpush:            cloudflare.F(false),
				Migrations: cloudflare.F[workers.ScriptUpdateParamsMetadataObjectMigrationsUnion](workers.SingleStepMigrationParam{
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
				Observability: cloudflare.F(workers.ScriptUpdateParamsMetadataObjectObservability{
					Enabled:          cloudflare.F(true),
					HeadSamplingRate: cloudflare.F(0.100000),
					Logs: cloudflare.F(workers.ScriptUpdateParamsMetadataObjectObservabilityLogs{
						Enabled:          cloudflare.F(true),
						InvocationLogs:   cloudflare.F(true),
						HeadSamplingRate: cloudflare.F(0.100000),
					}),
				}),
				Placement: cloudflare.F(workers.ScriptUpdateParamsMetadataObjectPlacement{
					Mode: cloudflare.F(workers.ScriptUpdateParamsMetadataObjectPlacementModeSmart),
				}),
				Tags: cloudflare.F([]string{"string"}),
				TailConsumers: cloudflare.F([]workers.ConsumerScriptParam{{
					Service:     cloudflare.F("my-log-consumer"),
					Environment: cloudflare.F("production"),
					Namespace:   cloudflare.F("my-namespace"),
				}}),
				UsageModel: cloudflare.F(workers.ScriptUpdateParamsMetadataObjectUsageModelStandard),
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
