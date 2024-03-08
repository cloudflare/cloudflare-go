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

func TestWorkerServiceEnvironmentSettingEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Services.Environments.Settings.Edit(
		context.TODO(),
		"my-worker",
		"production",
		cloudflare.WorkerServiceEnvironmentSettingEditParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Errors: cloudflare.F([]cloudflare.WorkerServiceEnvironmentSettingEditParamsError{{
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}, {
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}, {
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}}),
			Messages: cloudflare.F([]cloudflare.WorkerServiceEnvironmentSettingEditParamsMessage{{
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}, {
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}, {
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}}),
			Result: cloudflare.F(cloudflare.WorkerServiceEnvironmentSettingEditParamsResult{
				Bindings: cloudflare.F([]cloudflare.WorkerServiceEnvironmentSettingEditParamsResultBinding{cloudflare.WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding(cloudflare.WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding{
					Type: cloudflare.F(cloudflare.WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace),
				}), cloudflare.WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding(cloudflare.WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding{
					Type: cloudflare.F(cloudflare.WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace),
				}), cloudflare.WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding(cloudflare.WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding{
					Type: cloudflare.F(cloudflare.WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace),
				})}),
				CompatibilityDate:  cloudflare.F("2022-04-05"),
				CompatibilityFlags: cloudflare.F([]string{"formdata_parser_supports_files", "formdata_parser_supports_files", "formdata_parser_supports_files"}),
				Logpush:            cloudflare.F(false),
				Migrations: cloudflare.F[cloudflare.WorkerServiceEnvironmentSettingEditParamsResultMigrations](cloudflare.WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrations(cloudflare.WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrations{
					NewTag:         cloudflare.F("v2"),
					OldTag:         cloudflare.F("v1"),
					DeletedClasses: cloudflare.F([]string{"string", "string", "string"}),
					NewClasses:     cloudflare.F([]string{"string", "string", "string"}),
					RenamedClasses: cloudflare.F([]cloudflare.WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass{{
						From: cloudflare.F("string"),
						To:   cloudflare.F("string"),
					}, {
						From: cloudflare.F("string"),
						To:   cloudflare.F("string"),
					}, {
						From: cloudflare.F("string"),
						To:   cloudflare.F("string"),
					}}),
					TransferredClasses: cloudflare.F([]cloudflare.WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass{{
						From:       cloudflare.F("string"),
						FromScript: cloudflare.F("string"),
						To:         cloudflare.F("string"),
					}, {
						From:       cloudflare.F("string"),
						FromScript: cloudflare.F("string"),
						To:         cloudflare.F("string"),
					}, {
						From:       cloudflare.F("string"),
						FromScript: cloudflare.F("string"),
						To:         cloudflare.F("string"),
					}}),
				})),
				Placement: cloudflare.F(cloudflare.WorkerServiceEnvironmentSettingEditParamsResultPlacement{
					Mode: cloudflare.F(cloudflare.WorkerServiceEnvironmentSettingEditParamsResultPlacementModeSmart),
				}),
				Tags: cloudflare.F([]string{"my-tag", "my-tag", "my-tag"}),
				TailConsumers: cloudflare.F([]cloudflare.WorkerServiceEnvironmentSettingEditParamsResultTailConsumer{{
					Environment: cloudflare.F("production"),
					Namespace:   cloudflare.F("my-namespace"),
					Service:     cloudflare.F("my-log-consumer"),
				}, {
					Environment: cloudflare.F("production"),
					Namespace:   cloudflare.F("my-namespace"),
					Service:     cloudflare.F("my-log-consumer"),
				}, {
					Environment: cloudflare.F("production"),
					Namespace:   cloudflare.F("my-namespace"),
					Service:     cloudflare.F("my-log-consumer"),
				}}),
				UsageModel: cloudflare.F("unbound"),
			}),
			Success: cloudflare.F(cloudflare.WorkerServiceEnvironmentSettingEditParamsSuccessTrue),
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

func TestWorkerServiceEnvironmentSettingGet(t *testing.T) {
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
	_, err := client.Workers.Services.Environments.Settings.Get(
		context.TODO(),
		"my-worker",
		"production",
		cloudflare.WorkerServiceEnvironmentSettingGetParams{
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
