// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/workers"
	"github.com/cloudflare/cloudflare-go/v4/workers_for_platforms"
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
			Body: workers_for_platforms.DispatchNamespaceScriptUpdateParamsBodyMetadata{
				Metadata: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsBodyMetadataMetadata{
					Assets: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssets{
						Config: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfig{
							HTMLHandling:     cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingAutoTrailingSlash),
							NotFoundHandling: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandlingNone),
							ServeDirectly:    cloudflare.F(true),
						}),
						JWT: cloudflare.F("jwt"),
					}),
					Bindings: cloudflare.F([]workers_for_platforms.DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataBinding{{
						Name: cloudflare.F("MY_ENV_VAR"),
						Type: cloudflare.F("plain_text"),
					}}),
					BodyPart:           cloudflare.F("worker.js"),
					CompatibilityDate:  cloudflare.F("2021-01-01"),
					CompatibilityFlags: cloudflare.F([]string{"nodejs_compat"}),
					KeepAssets:         cloudflare.F(false),
					KeepBindings:       cloudflare.F([]string{"string"}),
					Logpush:            cloudflare.F(false),
					MainModule:         cloudflare.F("worker.js"),
					Migrations: cloudflare.F[workers_for_platforms.DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrationsUnion](workers.SingleStepMigrationParam{
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
					Observability: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataObservability{
						Enabled:          cloudflare.F(true),
						HeadSamplingRate: cloudflare.F(0.100000),
					}),
					Placement: cloudflare.F(workers.PlacementConfigurationParam{
						Mode: cloudflare.F(workers.PlacementConfigurationModeSmart),
					}),
					Tags: cloudflare.F([]string{"string"}),
					TailConsumers: cloudflare.F([]workers.ConsumerScriptParam{{
						Service:     cloudflare.F("my-log-consumer"),
						Environment: cloudflare.F("production"),
						Namespace:   cloudflare.F("my-namespace"),
					}}),
					UsageModel: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataUsageModelBundled),
					VersionTags: cloudflare.F(map[string]string{
						"foo": "string",
					}),
				}),
			},
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
	err := client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Delete(
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
