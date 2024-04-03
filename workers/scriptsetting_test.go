// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/workers"
)

func TestScriptSettingEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Scripts.Settings.Edit(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptSettingEditParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Settings: cloudflare.F(workers.ScriptSettingEditParamsSettings{
				Bindings: cloudflare.F([]workers.ScriptSettingEditParamsSettingsBinding{workers.ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding(workers.ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding{
					Type: cloudflare.F(workers.ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingTypeKVNamespace),
				}), workers.ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding(workers.ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding{
					Type: cloudflare.F(workers.ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingTypeKVNamespace),
				}), workers.ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding(workers.ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding{
					Type: cloudflare.F(workers.ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingTypeKVNamespace),
				})}),
				CompatibilityDate:  cloudflare.F("2022-04-05"),
				CompatibilityFlags: cloudflare.F([]string{"formdata_parser_supports_files", "formdata_parser_supports_files", "formdata_parser_supports_files"}),
				Logpush:            cloudflare.F(false),
				Migrations: cloudflare.F[workers.ScriptSettingEditParamsSettingsMigrations](workers.ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrations(workers.ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrations{
					NewTag:         cloudflare.F("v2"),
					OldTag:         cloudflare.F("v1"),
					DeletedClasses: cloudflare.F([]string{"string", "string", "string"}),
					NewClasses:     cloudflare.F([]string{"string", "string", "string"}),
					RenamedClasses: cloudflare.F([]workers.ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsRenamedClass{{
						From: cloudflare.F("string"),
						To:   cloudflare.F("string"),
					}, {
						From: cloudflare.F("string"),
						To:   cloudflare.F("string"),
					}, {
						From: cloudflare.F("string"),
						To:   cloudflare.F("string"),
					}}),
					TransferredClasses: cloudflare.F([]workers.ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsTransferredClass{{
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
				Placement: cloudflare.F(workers.ScriptSettingEditParamsSettingsPlacement{
					Mode: cloudflare.F(workers.ScriptSettingEditParamsSettingsPlacementModeSmart),
				}),
				Tags: cloudflare.F([]string{"my-tag", "my-tag", "my-tag"}),
				TailConsumers: cloudflare.F([]workers.ScriptSettingEditParamsSettingsTailConsumer{{
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

func TestScriptSettingGet(t *testing.T) {
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
	_, err := client.Workers.Scripts.Settings.Get(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptSettingGetParams{
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
