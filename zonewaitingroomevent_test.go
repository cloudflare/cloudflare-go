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

func TestZoneWaitingRoomEventGet(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Events.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneWaitingRoomEventUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Events.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		cloudflare.ZoneWaitingRoomEventUpdateParams{
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

func TestZoneWaitingRoomEventDelete(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Events.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneWaitingRoomEventPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Events.Patch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		cloudflare.ZoneWaitingRoomEventPatchParams{
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

func TestZoneWaitingRoomEventWaitingRoomNewEventWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Events.WaitingRoomNewEvent(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		cloudflare.ZoneWaitingRoomEventWaitingRoomNewEventParams{
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

func TestZoneWaitingRoomEventWaitingRoomListEvents(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Events.WaitingRoomListEvents(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
