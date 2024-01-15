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

func TestAccountWorkerServiceEnvironmentSettingGet(t *testing.T) {
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
	_, err := client.Accounts.Workers.Services.Environments.Settings.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"my-worker",
		"production",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountWorkerServiceEnvironmentSettingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Workers.Services.Environments.Settings.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"my-worker",
		"production",
		cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParams{
			Bindings: cloudflare.F([]cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsBinding{cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding(cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding{
				Type: cloudflare.F(cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace),
			}), cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding(cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding{
				Type: cloudflare.F(cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace),
			}), cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding(cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding{
				Type: cloudflare.F(cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace),
			})}),
			CompatibilityDate:  cloudflare.F("2022-04-05"),
			CompatibilityFlags: cloudflare.F([]string{"formdata_parser_supports_files", "formdata_parser_supports_files", "formdata_parser_supports_files"}),
			Logpush:            cloudflare.F(false),
			Migrations: cloudflare.F[cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsMigrations](cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrations(cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrations{
				NewTag:         cloudflare.F("v2"),
				OldTag:         cloudflare.F("v1"),
				DeletedClasses: cloudflare.F([]string{"string", "string", "string"}),
				NewClasses:     cloudflare.F([]string{"string", "string", "string"}),
				RenamedClasses: cloudflare.F([]cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsRenamedClass{{
					From: cloudflare.F("string"),
					To:   cloudflare.F("string"),
				}, {
					From: cloudflare.F("string"),
					To:   cloudflare.F("string"),
				}, {
					From: cloudflare.F("string"),
					To:   cloudflare.F("string"),
				}}),
				TransferredClasses: cloudflare.F([]cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsTransferredClass{{
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
			Placement: cloudflare.F(cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsPlacement{
				Mode: cloudflare.F(cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsPlacementModeSmart),
			}),
			Tags: cloudflare.F([]string{"my-tag", "my-tag", "my-tag"}),
			TailConsumers: cloudflare.F([]cloudflare.AccountWorkerServiceEnvironmentSettingUpdateParamsTailConsumer{{
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
