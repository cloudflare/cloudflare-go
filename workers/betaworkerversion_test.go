// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/workers"
)

func TestBetaWorkerVersionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Beta.Workers.Versions.New(
		context.TODO(),
		"worker_id",
		workers.BetaWorkerVersionNewParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Version: workers.VersionParam{
				Annotations: cloudflare.F(workers.VersionAnnotationsParam{
					WorkersMessage: cloudflare.F("Fixed bug."),
					WorkersTag:     cloudflare.F("v1.0.1"),
				}),
				Assets: cloudflare.F(workers.VersionAssetsParam{
					Config: cloudflare.F(workers.VersionAssetsConfigParam{
						HTMLHandling:     cloudflare.F(workers.VersionAssetsConfigHTMLHandlingAutoTrailingSlash),
						NotFoundHandling: cloudflare.F(workers.VersionAssetsConfigNotFoundHandling404Page),
						RunWorkerFirst:   cloudflare.F[workers.VersionAssetsConfigRunWorkerFirstUnionParam](workers.VersionAssetsConfigRunWorkerFirstArrayParam([]string{"string"})),
					}),
					JWT: cloudflare.F("jwt"),
				}),
				Bindings: cloudflare.F([]workers.VersionBindingsUnionParam{workers.VersionBindingsWorkersBindingKindPlainTextParam{
					Name: cloudflare.F("MY_ENV_VAR"),
					Text: cloudflare.F("my_data"),
					Type: cloudflare.F(workers.VersionBindingsWorkersBindingKindPlainTextTypePlainText),
				}}),
				CompatibilityDate:  cloudflare.F("2021-01-01"),
				CompatibilityFlags: cloudflare.F([]string{"nodejs_compat"}),
				Limits: cloudflare.F(workers.VersionLimitsParam{
					CPUMs: cloudflare.F(int64(50)),
				}),
				MainModule: cloudflare.F("index.js"),
				Migrations: cloudflare.F[workers.VersionMigrationsUnionParam](workers.SingleStepMigrationParam{
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
				Modules: cloudflare.F([]workers.VersionModuleParam{{
					ContentBase64: cloudflare.F("ZXhwb3J0IGRlZmF1bHQgewogIGFzeW5jIGZldGNoKHJlcXVlc3QsIGVudiwgY3R4KSB7CiAgICByZXR1cm4gbmV3IFJlc3BvbnNlKCdIZWxsbyBXb3JsZCEnKQogIH0KfQ=="),
					ContentType:   cloudflare.F("application/javascript+module"),
					Name:          cloudflare.F("index.js"),
				}}),
				Placement: cloudflare.F[workers.VersionPlacementUnionParam](workers.VersionPlacementModeParam{
					Mode: cloudflare.F(workers.VersionPlacementModeModeSmart),
				}),
				UsageModel: cloudflare.F(workers.VersionUsageModelStandard),
			},
			Deploy: cloudflare.F(true),
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

func TestBetaWorkerVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Beta.Workers.Versions.List(
		context.TODO(),
		"worker_id",
		workers.BetaWorkerVersionListParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Page:      cloudflare.F(int64(1)),
			PerPage:   cloudflare.F(int64(1)),
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

func TestBetaWorkerVersionDelete(t *testing.T) {
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
	_, err := client.Workers.Beta.Workers.Versions.Delete(
		context.TODO(),
		"worker_id",
		"version_id",
		workers.BetaWorkerVersionDeleteParams{
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

func TestBetaWorkerVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Beta.Workers.Versions.Get(
		context.TODO(),
		"worker_id",
		"version_id",
		workers.BetaWorkerVersionGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Include:   cloudflare.F(workers.BetaWorkerVersionGetParamsIncludeModules),
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
