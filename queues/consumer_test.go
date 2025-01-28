// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/queues"
)

func TestConsumerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Queues.Consumers.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		queues.ConsumerNewParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: queues.ConsumerNewParamsBodyMqWorkerConsumer{
				DeadLetterQueue: cloudflare.F("example-queue"),
				ScriptName:      cloudflare.F("my-consumer-worker"),
				Settings: cloudflare.F(queues.ConsumerNewParamsBodyMqWorkerConsumerSettings{
					BatchSize:      cloudflare.F(50.000000),
					MaxConcurrency: cloudflare.F(10.000000),
					MaxRetries:     cloudflare.F(3.000000),
					MaxWaitTimeMs:  cloudflare.F(5000.000000),
					RetryDelay:     cloudflare.F(10.000000),
				}),
				Type: cloudflare.F(queues.ConsumerNewParamsBodyMqWorkerConsumerTypeWorker),
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

func TestConsumerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Queues.Consumers.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		queues.ConsumerUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: queues.ConsumerUpdateParamsBodyMqWorkerConsumer{
				DeadLetterQueue: cloudflare.F("example-queue"),
				ScriptName:      cloudflare.F("my-consumer-worker"),
				Settings: cloudflare.F(queues.ConsumerUpdateParamsBodyMqWorkerConsumerSettings{
					BatchSize:      cloudflare.F(50.000000),
					MaxConcurrency: cloudflare.F(10.000000),
					MaxRetries:     cloudflare.F(3.000000),
					MaxWaitTimeMs:  cloudflare.F(5000.000000),
					RetryDelay:     cloudflare.F(10.000000),
				}),
				Type: cloudflare.F(queues.ConsumerUpdateParamsBodyMqWorkerConsumerTypeWorker),
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

func TestConsumerDelete(t *testing.T) {
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
	_, err := client.Queues.Consumers.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		queues.ConsumerDeleteParams{
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

func TestConsumerGet(t *testing.T) {
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
	_, err := client.Queues.Consumers.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		queues.ConsumerGetParams{
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
