// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/queues"
)

func TestConsumerNewWithOptionalParams(t *testing.T) {
	t.Skip("422 status codes in prism tests")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Queues.Consumers.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		queues.ConsumerNewParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: queues.ConsumerNewParamsBodyMqNotificationConsumerRequest{
				Settings: cloudflare.F[queues.ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsUnion](queues.ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObject{
					Email: cloudflare.F([]queues.ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObjectEmail{{
						ID: cloudflare.F("user@example.com"),
					}}),
					Pagerduty: cloudflare.F([]queues.ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObjectPagerduty{{
						ID: cloudflare.F("fedcba9876543210fedcba9876543210"),
					}}),
					Webhooks: cloudflare.F([]queues.ConsumerNewParamsBodyMqNotificationConsumerRequestSettingsObjectWebhook{{
						ID: cloudflare.F("0123456789abcdef0123456789abcdef"),
					}}),
				}),
				Type:            cloudflare.F(queues.ConsumerNewParamsBodyMqNotificationConsumerRequestTypeNotification),
				DeadLetterQueue: cloudflare.F("example-queue"),
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
	t.Skip("422 status codes in prism tests")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Queues.Consumers.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		queues.ConsumerUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: queues.ConsumerUpdateParamsBodyMqNotificationConsumerRequest{
				Settings: cloudflare.F[queues.ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsUnion](queues.ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObject{
					Email: cloudflare.F([]queues.ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObjectEmail{{
						ID: cloudflare.F("user@example.com"),
					}}),
					Pagerduty: cloudflare.F([]queues.ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObjectPagerduty{{
						ID: cloudflare.F("fedcba9876543210fedcba9876543210"),
					}}),
					Webhooks: cloudflare.F([]queues.ConsumerUpdateParamsBodyMqNotificationConsumerRequestSettingsObjectWebhook{{
						ID: cloudflare.F("0123456789abcdef0123456789abcdef"),
					}}),
				}),
				Type:            cloudflare.F(queues.ConsumerUpdateParamsBodyMqNotificationConsumerRequestTypeNotification),
				DeadLetterQueue: cloudflare.F("example-queue"),
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

func TestConsumerList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Queues.Consumers.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		queues.ConsumerListParams{
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Queues.Consumers.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
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
