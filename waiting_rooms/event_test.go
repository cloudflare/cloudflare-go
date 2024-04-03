// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/waiting_rooms"
)

func TestEventNewWithOptionalParams(t *testing.T) {
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
	_, err := client.WaitingRooms.Events.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.EventNewParams{
			ZoneID:                cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			EventEndTime:          cloudflare.F("2021-09-28T17:00:00.000Z"),
			EventStartTime:        cloudflare.F("2021-09-28T15:30:00.000Z"),
			Name:                  cloudflare.F("production_webinar_event"),
			CustomPageHTML:        cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Event is prequeueing / Queue all enabled {{/waitTimeKnown}}"),
			Description:           cloudflare.F("Production event - DO NOT MODIFY"),
			DisableSessionRenewal: cloudflare.F(true),
			NewUsersPerMinute:     cloudflare.F(int64(200)),
			PrequeueStartTime:     cloudflare.F("2021-09-28T15:00:00.000Z"),
			QueueingMethod:        cloudflare.F("random"),
			SessionDuration:       cloudflare.F(int64(1)),
			ShuffleAtEventStart:   cloudflare.F(true),
			Suspended:             cloudflare.F(true),
			TotalActiveUsers:      cloudflare.F(int64(200)),
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

func TestEventUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.WaitingRooms.Events.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		waiting_rooms.EventUpdateParams{
			ZoneID:                cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			EventEndTime:          cloudflare.F("2021-09-28T17:00:00.000Z"),
			EventStartTime:        cloudflare.F("2021-09-28T15:30:00.000Z"),
			Name:                  cloudflare.F("production_webinar_event"),
			CustomPageHTML:        cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Event is prequeueing / Queue all enabled {{/waitTimeKnown}}"),
			Description:           cloudflare.F("Production event - DO NOT MODIFY"),
			DisableSessionRenewal: cloudflare.F(true),
			NewUsersPerMinute:     cloudflare.F(int64(200)),
			PrequeueStartTime:     cloudflare.F("2021-09-28T15:00:00.000Z"),
			QueueingMethod:        cloudflare.F("random"),
			SessionDuration:       cloudflare.F(int64(1)),
			ShuffleAtEventStart:   cloudflare.F(true),
			Suspended:             cloudflare.F(true),
			TotalActiveUsers:      cloudflare.F(int64(200)),
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

func TestEventList(t *testing.T) {
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
	_, err := client.WaitingRooms.Events.List(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.EventListParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestEventDelete(t *testing.T) {
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
	_, err := client.WaitingRooms.Events.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		waiting_rooms.EventDeleteParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body:   cloudflare.F[any](map[string]interface{}{}),
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

func TestEventEditWithOptionalParams(t *testing.T) {
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
	_, err := client.WaitingRooms.Events.Edit(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		waiting_rooms.EventEditParams{
			ZoneID:                cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			EventEndTime:          cloudflare.F("2021-09-28T17:00:00.000Z"),
			EventStartTime:        cloudflare.F("2021-09-28T15:30:00.000Z"),
			Name:                  cloudflare.F("production_webinar_event"),
			CustomPageHTML:        cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Event is prequeueing / Queue all enabled {{/waitTimeKnown}}"),
			Description:           cloudflare.F("Production event - DO NOT MODIFY"),
			DisableSessionRenewal: cloudflare.F(true),
			NewUsersPerMinute:     cloudflare.F(int64(200)),
			PrequeueStartTime:     cloudflare.F("2021-09-28T15:00:00.000Z"),
			QueueingMethod:        cloudflare.F("random"),
			SessionDuration:       cloudflare.F(int64(1)),
			ShuffleAtEventStart:   cloudflare.F(true),
			Suspended:             cloudflare.F(true),
			TotalActiveUsers:      cloudflare.F(int64(200)),
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

func TestEventGet(t *testing.T) {
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
	_, err := client.WaitingRooms.Events.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		waiting_rooms.EventGetParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
