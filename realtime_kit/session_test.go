// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package realtime_kit_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/realtime_kit"
)

func TestSessionGenerateSummaryOfTranscripts(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Sessions.GenerateSummaryOfTranscripts(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.SessionGenerateSummaryOfTranscriptsParams{
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

func TestSessionGetParticipantDataFromPeerIDWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Sessions.GetParticipantDataFromPeerID(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.SessionGetParticipantDataFromPeerIDParams{
			AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Filters:           cloudflare.F(realtime_kit.SessionGetParticipantDataFromPeerIDParamsFiltersDeviceInfo),
			IncludePeerEvents: cloudflare.F(true),
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

func TestSessionGetSessionChat(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Sessions.GetSessionChat(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.SessionGetSessionChatParams{
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

func TestSessionGetSessionDetailsWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Sessions.GetSessionDetails(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.SessionGetSessionDetailsParams{
			AccountID:            cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			IncludeBreakoutRooms: cloudflare.F(true),
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

func TestSessionGetSessionParticipantDetailsWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Sessions.GetSessionParticipantDetails(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.SessionGetSessionParticipantDetailsParams{
			AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Filters:           cloudflare.F(realtime_kit.SessionGetSessionParticipantDetailsParamsFiltersDeviceInfo),
			IncludePeerEvents: cloudflare.F(true),
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

func TestSessionGetSessionParticipantsWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Sessions.GetSessionParticipants(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.SessionGetSessionParticipantsParams{
			AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			IncludePeerEvents: cloudflare.F(true),
			PageNo:            cloudflare.F(0.000000),
			PerPage:           cloudflare.F(0.000000),
			Search:            cloudflare.F("search"),
			SortBy:            cloudflare.F(realtime_kit.SessionGetSessionParticipantsParamsSortByJoinedAt),
			SortOrder:         cloudflare.F(realtime_kit.SessionGetSessionParticipantsParamsSortOrderAsc),
			View:              cloudflare.F(realtime_kit.SessionGetSessionParticipantsParamsViewRaw),
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

func TestSessionGetSessionSummary(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Sessions.GetSessionSummary(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.SessionGetSessionSummaryParams{
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

func TestSessionGetSessionTranscriptsWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Sessions.GetSessionTranscripts(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.SessionGetSessionTranscriptsParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Format:    cloudflare.F(realtime_kit.SessionGetSessionTranscriptsParamsFormatSrt),
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

func TestSessionGetSessionsWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Sessions.GetSessions(
		context.TODO(),
		"app_id",
		realtime_kit.SessionGetSessionsParams{
			AccountID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AssociatedID: cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			EndTime:      cloudflare.F(time.Now()),
			PageNo:       cloudflare.F(0.000000),
			Participants: cloudflare.F("1:10"),
			PerPage:      cloudflare.F(0.000000),
			Search:       cloudflare.F("search"),
			SortBy:       cloudflare.F(realtime_kit.SessionGetSessionsParamsSortByMinutesConsumed),
			SortOrder:    cloudflare.F(realtime_kit.SessionGetSessionsParamsSortOrderAsc),
			StartTime:    cloudflare.F(time.Now()),
			Status:       cloudflare.F(realtime_kit.SessionGetSessionsParamsStatusLive),
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
