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

func TestBetaWorkerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Beta.Workers.New(context.TODO(), workers.BetaWorkerNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Worker: workers.WorkerParam{
			Name:    cloudflare.F("my-worker"),
			Logpush: cloudflare.F(true),
			Observability: cloudflare.F(workers.WorkerObservabilityParam{
				Enabled:          cloudflare.F(true),
				HeadSamplingRate: cloudflare.F(0.000000),
				Logs: cloudflare.F(workers.WorkerObservabilityLogsParam{
					Enabled:          cloudflare.F(true),
					HeadSamplingRate: cloudflare.F(0.000000),
					InvocationLogs:   cloudflare.F(true),
				}),
			}),
			Subdomain: cloudflare.F(workers.WorkerSubdomainParam{
				Enabled:         cloudflare.F(true),
				PreviewsEnabled: cloudflare.F(true),
			}),
			Tags: cloudflare.F([]string{"my-team", "my-public-api"}),
			TailConsumers: cloudflare.F([]workers.WorkerTailConsumerParam{{
				Name: cloudflare.F("my-tail-consumer"),
			}}),
		},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaWorkerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Beta.Workers.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		workers.BetaWorkerUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Worker: workers.WorkerParam{
				Name:    cloudflare.F("my-worker"),
				Logpush: cloudflare.F(true),
				Observability: cloudflare.F(workers.WorkerObservabilityParam{
					Enabled:          cloudflare.F(true),
					HeadSamplingRate: cloudflare.F(0.000000),
					Logs: cloudflare.F(workers.WorkerObservabilityLogsParam{
						Enabled:          cloudflare.F(true),
						HeadSamplingRate: cloudflare.F(0.000000),
						InvocationLogs:   cloudflare.F(true),
					}),
				}),
				Subdomain: cloudflare.F(workers.WorkerSubdomainParam{
					Enabled:         cloudflare.F(true),
					PreviewsEnabled: cloudflare.F(true),
				}),
				Tags: cloudflare.F([]string{"my-team", "my-public-api"}),
				TailConsumers: cloudflare.F([]workers.WorkerTailConsumerParam{{
					Name: cloudflare.F("my-tail-consumer"),
				}}),
			},
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

func TestBetaWorkerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Beta.Workers.List(context.TODO(), workers.BetaWorkerListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(1)),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaWorkerDelete(t *testing.T) {
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
	_, err := client.Workers.Beta.Workers.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		workers.BetaWorkerDeleteParams{
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

func TestBetaWorkerGet(t *testing.T) {
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
	_, err := client.Workers.Beta.Workers.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		workers.BetaWorkerGetParams{
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
