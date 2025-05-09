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

func TestScriptScriptAndVersionSettingEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Scripts.ScriptAndVersionSettings.Edit(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptScriptAndVersionSettingEditParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Settings: cloudflare.F(workers.ScriptScriptAndVersionSettingEditParamsSettings{
				Bindings: cloudflare.F([]workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingUnion{workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainText{
					Name: cloudflare.F("MY_ENV_VAR"),
					Text: cloudflare.F("my_data"),
					Type: cloudflare.F(workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypePlainText),
				}}),
				CompatibilityDate:  cloudflare.F("2021-01-01"),
				CompatibilityFlags: cloudflare.F([]string{"nodejs_compat"}),
				Limits: cloudflare.F(workers.ScriptScriptAndVersionSettingEditParamsSettingsLimits{
					CPUMs: cloudflare.F(int64(50)),
				}),
				Logpush: cloudflare.F(false),
				Migrations: cloudflare.F[workers.ScriptScriptAndVersionSettingEditParamsSettingsMigrationsUnion](workers.SingleStepMigrationParam{
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
				Observability: cloudflare.F(workers.ScriptScriptAndVersionSettingEditParamsSettingsObservability{
					Enabled:          cloudflare.F(true),
					HeadSamplingRate: cloudflare.F(0.100000),
				}),
				Placement: cloudflare.F(workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacement{
					Mode: cloudflare.F(workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementModeSmart),
				}),
				Tags: cloudflare.F([]string{"my-tag"}),
				TailConsumers: cloudflare.F([]workers.ConsumerScriptParam{{
					Service:     cloudflare.F("my-log-consumer"),
					Environment: cloudflare.F("production"),
					Namespace:   cloudflare.F("my-namespace"),
				}}),
				UsageModel: cloudflare.F(workers.ScriptScriptAndVersionSettingEditParamsSettingsUsageModelStandard),
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

func TestScriptScriptAndVersionSettingGet(t *testing.T) {
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
	_, err := client.Workers.Scripts.ScriptAndVersionSettings.Get(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptScriptAndVersionSettingGetParams{
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
