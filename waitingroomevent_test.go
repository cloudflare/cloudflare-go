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

func TestWaitingRoomEventNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.WaitingRooms.Events.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		cloudflare.WaitingRoomEventNewParams{
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

func TestWaitingRoomEventUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.WaitingRooms.Events.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		cloudflare.WaitingRoomEventUpdateParams{
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

func TestWaitingRoomEventList(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.WaitingRooms.Events.List(
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

func TestWaitingRoomEventDelete(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.WaitingRooms.Events.Delete(
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

func TestWaitingRoomEventGet(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.WaitingRooms.Events.Get(
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
