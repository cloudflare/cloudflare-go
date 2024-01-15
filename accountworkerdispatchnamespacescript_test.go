// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestAccountWorkerDispatchNamespaceScriptNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Workers.Dispatch.Namespaces.Scripts.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"my-dispatch-namespace",
		"this-is_my_script-01",
		cloudflare.AccountWorkerDispatchNamespaceScriptNewParamsVariant0{
			AnyPartName: cloudflare.F([]io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents"))), io.Reader(bytes.NewBuffer([]byte("some file contents"))), io.Reader(bytes.NewBuffer([]byte("some file contents")))}),
			Metadata: cloudflare.F(cloudflare.AccountWorkerDispatchNamespaceScriptNewParamsVariant0Metadata{
				Bindings:           cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				BodyPart:           cloudflare.F("worker.js"),
				CompatibilityDate:  cloudflare.F("2023-07-25"),
				CompatibilityFlags: cloudflare.F([]string{"string", "string", "string"}),
				KeepBindings:       cloudflare.F([]string{"string", "string", "string"}),
				Logpush:            cloudflare.F(false),
				MainModule:         cloudflare.F("worker.js"),
				Migrations: cloudflare.F[cloudflare.AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrations](cloudflare.AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrations(cloudflare.AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrations{
					NewTag:         cloudflare.F("v2"),
					OldTag:         cloudflare.F("v1"),
					DeletedClasses: cloudflare.F([]string{"string", "string", "string"}),
					NewClasses:     cloudflare.F([]string{"string", "string", "string"}),
					RenamedClasses: cloudflare.F([]cloudflare.AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsRenamedClass{{
						From: cloudflare.F("string"),
						To:   cloudflare.F("string"),
					}, {
						From: cloudflare.F("string"),
						To:   cloudflare.F("string"),
					}, {
						From: cloudflare.F("string"),
						To:   cloudflare.F("string"),
					}}),
					TransferredClasses: cloudflare.F([]cloudflare.AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsTransferredClass{{
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
				Placement: cloudflare.F(cloudflare.AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataPlacement{
					Mode: cloudflare.F(cloudflare.AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataPlacementModeSmart),
				}),
				Tags: cloudflare.F([]string{"string", "string", "string"}),
				TailConsumers: cloudflare.F([]cloudflare.AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataTailConsumer{{
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
				UsageModel:  cloudflare.F(cloudflare.AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataUsageModelBundled),
				VersionTags: cloudflare.F[any](map[string]interface{}{}),
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

func TestAccountWorkerDispatchNamespaceScriptGet(t *testing.T) {
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
	_, err := client.Accounts.Workers.Dispatch.Namespaces.Scripts.Get(
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

func TestAccountWorkerDispatchNamespaceScriptDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Accounts.Workers.Dispatch.Namespaces.Scripts.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"my-dispatch-namespace",
		"this-is_my_script-01",
		cloudflare.AccountWorkerDispatchNamespaceScriptDeleteParams{
			Force: cloudflare.F(true),
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
