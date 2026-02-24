// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_test

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
	"github.com/cloudflare/cloudflare-go/v6/workers_for_platforms"
)

func TestDispatchNamespaceScriptUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Update(
		context.TODO(),
		"my-dispatch-namespace",
		"this-is_my_script-01",
		workers_for_platforms.DispatchNamespaceScriptUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Metadata: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadata{
				Assets: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataAssets{
					Config: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataAssetsConfig{
						Headers:          cloudflare.F("/dashboard/*\nX-Frame-Options: DENY\n\n/static/*\nAccess-Control-Allow-Origin: *"),
						Redirects:        cloudflare.F("/foo /bar 301\n/news/* /blog/:splat"),
						HTMLHandling:     cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash),
						NotFoundHandling: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling404Page),
						RunWorkerFirst:   cloudflare.F[workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion](workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstArray([]string{"string"})),
						ServeDirectly:    cloudflare.F(true),
					}),
					JWT: cloudflare.F("jwt"),
				}),
				Bindings: cloudflare.F([]workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingUnion{workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText{
					Name: cloudflare.F("MY_ENV_VAR"),
					Text: cloudflare.F("my_data"),
					Type: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText),
				}}),
				BodyPart:           cloudflare.F("worker.js"),
				CompatibilityDate:  cloudflare.F("2021-01-01"),
				CompatibilityFlags: cloudflare.F([]string{"nodejs_compat"}),
				KeepAssets:         cloudflare.F(false),
				KeepBindings:       cloudflare.F([]string{"string"}),
				Limits: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataLimits{
					CPUMs: cloudflare.F(int64(50)),
				}),
				Logpush:    cloudflare.F(false),
				MainModule: cloudflare.F("worker.js"),
				Migrations: cloudflare.F[workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion](workers.SingleStepMigrationParam{
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
				Observability: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObservability{
					Enabled:          cloudflare.F(true),
					HeadSamplingRate: cloudflare.F(0.100000),
					Logs: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObservabilityLogs{
						Enabled:          cloudflare.F(true),
						InvocationLogs:   cloudflare.F(true),
						Destinations:     cloudflare.F([]string{"cloudflare"}),
						HeadSamplingRate: cloudflare.F(0.100000),
						Persist:          cloudflare.F(true),
					}),
				}),
				Placement: cloudflare.F[workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataPlacementUnion](workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataPlacementObject{
					Mode: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectModeSmart),
				}),
				Tags: cloudflare.F([]string{"string"}),
				TailConsumers: cloudflare.F([]workers.ConsumerScriptParam{{
					Service:     cloudflare.F("my-log-consumer"),
					Environment: cloudflare.F("production"),
					Namespace:   cloudflare.F("my-namespace"),
				}}),
				UsageModel: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataUsageModelStandard),
			}),
			BindingsInherit: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsBindingsInheritStrict),
			Files:           cloudflare.F([]io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents")))}),
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

func TestDispatchNamespaceScriptDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Delete(
		context.TODO(),
		"my-dispatch-namespace",
		"this-is_my_script-01",
		workers_for_platforms.DispatchNamespaceScriptDeleteParams{
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

func TestDispatchNamespaceScriptGet(t *testing.T) {
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
	_, err := client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Get(
		context.TODO(),
		"my-dispatch-namespace",
		"this-is_my_script-01",
		workers_for_platforms.DispatchNamespaceScriptGetParams{
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
