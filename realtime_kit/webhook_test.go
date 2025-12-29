// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package realtime_kit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/realtime_kit"
)

func TestWebhookNewWebhookWithOptionalParams(t *testing.T) {
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
	_, err := client.RealtimeKit.Webhooks.NewWebhook(
		context.TODO(),
		"app_id",
		realtime_kit.WebhookNewWebhookParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Events:    cloudflare.F([]realtime_kit.WebhookNewWebhookParamsEvent{realtime_kit.WebhookNewWebhookParamsEventMeetingStarted, realtime_kit.WebhookNewWebhookParamsEventMeetingEnded, realtime_kit.WebhookNewWebhookParamsEventMeetingParticipantJoined, realtime_kit.WebhookNewWebhookParamsEventMeetingParticipantLeft, realtime_kit.WebhookNewWebhookParamsEventMeetingChatSynced, realtime_kit.WebhookNewWebhookParamsEventRecordingStatusUpdate, realtime_kit.WebhookNewWebhookParamsEventLivestreamingStatusUpdate, realtime_kit.WebhookNewWebhookParamsEventMeetingTranscript, realtime_kit.WebhookNewWebhookParamsEventMeetingSummary}),
			Name:      cloudflare.F("All events webhook"),
			URL:       cloudflare.F("https://webhook.site/b23a5bbd-c7b0-4ced-a9e2-78ae7889897e"),
			Enabled:   cloudflare.F(true),
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

func TestWebhookDeleteWebhook(t *testing.T) {
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
	_, err := client.RealtimeKit.Webhooks.DeleteWebhook(
		context.TODO(),
		"app_id",
		"webhook_id",
		realtime_kit.WebhookDeleteWebhookParams{
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

func TestWebhookEditWebhookWithOptionalParams(t *testing.T) {
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
	_, err := client.RealtimeKit.Webhooks.EditWebhook(
		context.TODO(),
		"app_id",
		"webhook_id",
		realtime_kit.WebhookEditWebhookParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Enabled:   cloudflare.F(true),
			Events:    cloudflare.F([]realtime_kit.WebhookEditWebhookParamsEvent{realtime_kit.WebhookEditWebhookParamsEventMeetingStarted}),
			Name:      cloudflare.F("name"),
			URL:       cloudflare.F("https://webhook.site/b23a5bbd-c7b0-4ced-a9e2-78ae7889897e"),
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

func TestWebhookGetWebhookByID(t *testing.T) {
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
	_, err := client.RealtimeKit.Webhooks.GetWebhookByID(
		context.TODO(),
		"app_id",
		"webhook_id",
		realtime_kit.WebhookGetWebhookByIDParams{
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

func TestWebhookGetWebhooks(t *testing.T) {
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
	_, err := client.RealtimeKit.Webhooks.GetWebhooks(
		context.TODO(),
		"app_id",
		realtime_kit.WebhookGetWebhooksParams{
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

func TestWebhookReplaceWebhookWithOptionalParams(t *testing.T) {
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
	_, err := client.RealtimeKit.Webhooks.ReplaceWebhook(
		context.TODO(),
		"app_id",
		"webhook_id",
		realtime_kit.WebhookReplaceWebhookParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Events:    cloudflare.F([]realtime_kit.WebhookReplaceWebhookParamsEvent{realtime_kit.WebhookReplaceWebhookParamsEventMeetingStarted, realtime_kit.WebhookReplaceWebhookParamsEventMeetingEnded, realtime_kit.WebhookReplaceWebhookParamsEventMeetingParticipantJoined, realtime_kit.WebhookReplaceWebhookParamsEventMeetingParticipantLeft, realtime_kit.WebhookReplaceWebhookParamsEventMeetingChatSynced, realtime_kit.WebhookReplaceWebhookParamsEventRecordingStatusUpdate, realtime_kit.WebhookReplaceWebhookParamsEventLivestreamingStatusUpdate, realtime_kit.WebhookReplaceWebhookParamsEventMeetingTranscript, realtime_kit.WebhookReplaceWebhookParamsEventMeetingSummary}),
			Name:      cloudflare.F("All events webhook"),
			URL:       cloudflare.F("https://webhook.site/b23a5bbd-c7b0-4ced-a9e2-78ae7889897e"),
			Enabled:   cloudflare.F(true),
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
