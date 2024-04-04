// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/workers"
)

func TestServiceEnvironmentSettingEditWithOptionalParams(t *testing.T) {
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
		workers.ServiceEnvironmentSettingEditParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Errors: cloudflare.F([]shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72Param{{
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}, {
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}, {
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}}),
			Messages: cloudflare.F([]shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72Param{{
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}, {
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}, {
				Code:    cloudflare.F(int64(1000)),
				Message: cloudflare.F("string"),
			}}),
			Result: cloudflare.F(workers.ServiceEnvironmentSettingEditParamsResult{
				Logpush: cloudflare.F(false),
				TailConsumers: cloudflare.F([]workers.ServiceEnvironmentSettingEditParamsResultTailConsumer{{
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
			}),
			Success: cloudflare.F(workers.ServiceEnvironmentSettingEditParamsSuccessTrue),
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

func TestServiceEnvironmentSettingGet(t *testing.T) {
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
		workers.ServiceEnvironmentSettingGetParams{
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
