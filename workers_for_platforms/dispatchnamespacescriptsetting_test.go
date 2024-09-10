// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/workers"
	"github.com/cloudflare/cloudflare-go/workers_for_platforms"
)

func TestDispatchNamespaceScriptSettingEditWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Settings.Edit(
		context.TODO(),
		"my-dispatch-namespace",
		"this-is_my_script-01",
		workers_for_platforms.DispatchNamespaceScriptSettingEditParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Settings: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettings{
				Bindings: cloudflare.F([]workers.BindingUnionParam{workers.KVNamespaceBindingParam{
					Type: cloudflare.F(workers.KVNamespaceBindingTypeKVNamespace),
				}, workers.KVNamespaceBindingParam{
					Type: cloudflare.F(workers.KVNamespaceBindingTypeKVNamespace),
				}, workers.KVNamespaceBindingParam{
					Type: cloudflare.F(workers.KVNamespaceBindingTypeKVNamespace),
				}}),
				CompatibilityDate:  cloudflare.F("2022-04-05"),
				CompatibilityFlags: cloudflare.F([]string{"formdata_parser_supports_files", "formdata_parser_supports_files", "formdata_parser_supports_files"}),
				Limits: cloudflare.F(workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsLimits{
					CPUMs: cloudflare.F(int64(50)),
				}),
				Logpush: cloudflare.F(false),
				Migrations: cloudflare.F[workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion](workers.SingleStepMigrationParam{
					DeletedClasses:   cloudflare.F([]string{"string", "string", "string"}),
					NewClasses:       cloudflare.F([]string{"string", "string", "string"}),
					NewSqliteClasses: cloudflare.F([]string{"string", "string", "string"}),
					NewTag:           cloudflare.F("v2"),
					OldTag:           cloudflare.F("v1"),
					RenamedClasses: cloudflare.F([]workers.SingleStepMigrationRenamedClassParam{{
						From: cloudflare.F("from"),
						To:   cloudflare.F("to"),
					}, {
						From: cloudflare.F("from"),
						To:   cloudflare.F("to"),
					}, {
						From: cloudflare.F("from"),
						To:   cloudflare.F("to"),
					}}),
					TransferredClasses: cloudflare.F([]workers.SingleStepMigrationTransferredClassParam{{
						From:       cloudflare.F("from"),
						FromScript: cloudflare.F("from_script"),
						To:         cloudflare.F("to"),
					}, {
						From:       cloudflare.F("from"),
						FromScript: cloudflare.F("from_script"),
						To:         cloudflare.F("to"),
					}, {
						From:       cloudflare.F("from"),
						FromScript: cloudflare.F("from_script"),
						To:         cloudflare.F("to"),
					}}),
				}),
				Placement: cloudflare.F(workers.PlacementConfigurationParam{
					Mode: cloudflare.F(workers.PlacementConfigurationModeSmart),
				}),
				Tags: cloudflare.F([]string{"my-tag", "my-tag", "my-tag"}),
				TailConsumers: cloudflare.F([]workers.ConsumerScriptParam{{
					Service:     cloudflare.F("my-log-consumer"),
					Environment: cloudflare.F("production"),
					Namespace:   cloudflare.F("my-namespace"),
				}, {
					Service:     cloudflare.F("my-log-consumer"),
					Environment: cloudflare.F("production"),
					Namespace:   cloudflare.F("my-namespace"),
				}, {
					Service:     cloudflare.F("my-log-consumer"),
					Environment: cloudflare.F("production"),
					Namespace:   cloudflare.F("my-namespace"),
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

func TestDispatchNamespaceScriptSettingGet(t *testing.T) {
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
	_, err := client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Settings.Get(
		context.TODO(),
		"my-dispatch-namespace",
		"this-is_my_script-01",
		workers_for_platforms.DispatchNamespaceScriptSettingGetParams{
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
