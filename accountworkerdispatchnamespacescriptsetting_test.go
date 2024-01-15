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

func TestAccountWorkerDispatchNamespaceScriptSettingGet(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Workers.Dispatch.Namespaces.Scripts.Settings.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"my-dispatch-namespace",
		"this-is_my_script-01",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountWorkerDispatchNamespaceScriptSettingUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Workers.Dispatch.Namespaces.Scripts.Settings.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"my-dispatch-namespace",
		"this-is_my_script-01",
		cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParams{
			Bindings: cloudflare.F([]cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBinding{cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding(cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding{
				Type: cloudflare.F(cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace),
			}), cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding(cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding{
				Type: cloudflare.F(cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace),
			}), cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding(cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding{
				Type: cloudflare.F(cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace),
			})}),
			CompatibilityDate:  cloudflare.F("2022-04-05"),
			CompatibilityFlags: cloudflare.F([]string{"formdata_parser_supports_files", "formdata_parser_supports_files", "formdata_parser_supports_files"}),
			Logpush:            cloudflare.F(false),
			Migrations: cloudflare.F[cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrations](cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrations(cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrations{
				NewTag:         cloudflare.F("v2"),
				OldTag:         cloudflare.F("v1"),
				DeletedClasses: cloudflare.F([]string{"string", "string", "string"}),
				NewClasses:     cloudflare.F([]string{"string", "string", "string"}),
				RenamedClasses: cloudflare.F([]cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsRenamedClass{{
					From: cloudflare.F("string"),
					To:   cloudflare.F("string"),
				}, {
					From: cloudflare.F("string"),
					To:   cloudflare.F("string"),
				}, {
					From: cloudflare.F("string"),
					To:   cloudflare.F("string"),
				}}),
				TransferredClasses: cloudflare.F([]cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsTransferredClass{{
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
			Placement: cloudflare.F(cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsPlacement{
				Mode: cloudflare.F(cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsPlacementModeSmart),
			}),
			Tags: cloudflare.F([]string{"my-tag", "my-tag", "my-tag"}),
			TailConsumers: cloudflare.F([]cloudflare.AccountWorkerDispatchNamespaceScriptSettingUpdateParamsTailConsumer{{
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
