// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestWorkersForPlatformDispatchNamespaceScriptContentSettingEditWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Content.Settings.Edit(
		context.TODO(),
		"my-dispatch-namespace",
		"this-is_my_script-01",
		cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Errors: cloudflare.F([]cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsError{{
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}, {
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}, {
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}}),
			Messages: cloudflare.F([]cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsMessage{{
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}, {
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}, {
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}}),
			Result: cloudflare.F(cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResult{
				Bindings: cloudflare.F([]cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBinding{cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding(cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding{
					Type: cloudflare.F(cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace),
				}), cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding(cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding{
					Type: cloudflare.F(cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace),
				}), cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding(cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding{
					Type: cloudflare.F(cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace),
				})}),
				CompatibilityDate:  cloudflare.F("2022-04-05"),
				CompatibilityFlags: cloudflare.F([]string{"formdata_parser_supports_files", "formdata_parser_supports_files", "formdata_parser_supports_files"}),
				Logpush:            cloudflare.F(false),
				Migrations: cloudflare.F[cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrations](cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrations(cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrations{
					NewTag:         cloudflare.F("v2"),
					OldTag:         cloudflare.F("v1"),
					DeletedClasses: cloudflare.F([]string{"string", "string", "string"}),
					NewClasses:     cloudflare.F([]string{"string", "string", "string"}),
					RenamedClasses: cloudflare.F([]cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass{{
						From: cloudflare.F("string"),
						To:   cloudflare.F("string"),
					}, {
						From: cloudflare.F("string"),
						To:   cloudflare.F("string"),
					}, {
						From: cloudflare.F("string"),
						To:   cloudflare.F("string"),
					}}),
					TransferredClasses: cloudflare.F([]cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass{{
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
				Placement: cloudflare.F(cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultPlacement{
					Mode: cloudflare.F(cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultPlacementModeSmart),
				}),
				Tags: cloudflare.F([]string{"my-tag", "my-tag", "my-tag"}),
				TailConsumers: cloudflare.F([]cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultTailConsumer{{
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
			Success: cloudflare.F(cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsSuccessTrue),
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

func TestWorkersForPlatformDispatchNamespaceScriptContentSettingGet(t *testing.T) {
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
	)
	_, err := client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Content.Settings.Get(
		context.TODO(),
		"my-dispatch-namespace",
		"this-is_my_script-01",
		cloudflare.WorkersForPlatformDispatchNamespaceScriptContentSettingGetParams{
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
