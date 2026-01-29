// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package realtime_kit_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/realtime_kit"
)

func TestLivestreamNewIndependentLivestreamWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Livestreams.NewIndependentLivestream(
		context.TODO(),
		"app_id",
		realtime_kit.LivestreamNewIndependentLivestreamParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:      cloudflare.F("prdmmp-xhycsl"),
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

func TestLivestreamGetActiveLivestreamsForLivestreamID(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Livestreams.GetActiveLivestreamsForLivestreamID(
		context.TODO(),
		"app_id",
		"livestream_id",
		realtime_kit.LivestreamGetActiveLivestreamsForLivestreamIDParams{
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

func TestLivestreamGetAllLivestreamsWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Livestreams.GetAllLivestreams(
		context.TODO(),
		"app_id",
		realtime_kit.LivestreamGetAllLivestreamsParams{
			AccountID:       cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			EndTime:         cloudflare.F(time.Now()),
			ExcludeMeetings: cloudflare.F(true),
			PageNo:          cloudflare.F(int64(0)),
			PerPage:         cloudflare.F(int64(0)),
			SortOrder:       cloudflare.F(realtime_kit.LivestreamGetAllLivestreamsParamsSortOrderAsc),
			StartTime:       cloudflare.F(time.Now()),
			Status:          cloudflare.F(realtime_kit.LivestreamGetAllLivestreamsParamsStatusLive),
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

func TestLivestreamGetLivestreamAnalyticsCompleteWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Livestreams.GetLivestreamAnalyticsComplete(
		context.TODO(),
		"app_id",
		realtime_kit.LivestreamGetLivestreamAnalyticsCompleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			EndTime:   cloudflare.F(time.Now()),
			StartTime: cloudflare.F(time.Now()),
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

func TestLivestreamGetLivestreamSessionDetailsForSessionID(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Livestreams.GetLivestreamSessionDetailsForSessionID(
		context.TODO(),
		"app_id",
		"livestream-session-id",
		realtime_kit.LivestreamGetLivestreamSessionDetailsForSessionIDParams{
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

func TestLivestreamGetLivestreamSessionForLivestreamIDWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Livestreams.GetLivestreamSessionForLivestreamID(
		context.TODO(),
		"app_id",
		"livestream_id",
		realtime_kit.LivestreamGetLivestreamSessionForLivestreamIDParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			PageNo:    cloudflare.F(int64(0)),
			PerPage:   cloudflare.F(int64(0)),
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

func TestLivestreamGetMeetingActiveLivestreams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Livestreams.GetMeetingActiveLivestreams(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.LivestreamGetMeetingActiveLivestreamsParams{
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

func TestLivestreamGetOrgAnalyticsWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Livestreams.GetOrgAnalytics(
		context.TODO(),
		"app_id",
		realtime_kit.LivestreamGetOrgAnalyticsParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			EndDate:   cloudflare.F("2022-09-22"),
			StartDate: cloudflare.F("2022-09-01"),
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

func TestLivestreamStartLivestreamingAMeetingWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Livestreams.StartLivestreamingAMeeting(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.LivestreamStartLivestreamingAMeetingParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:      cloudflare.F("prdmmp-xhycsl"),
			VideoConfig: cloudflare.F(realtime_kit.LivestreamStartLivestreamingAMeetingParamsVideoConfig{
				Height: cloudflare.F(int64(0)),
				Width:  cloudflare.F(int64(0)),
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

func TestLivestreamStopLivestreamingAMeeting(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Livestreams.StopLivestreamingAMeeting(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.LivestreamStopLivestreamingAMeetingParams{
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
