// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/radar"
)

func TestBGPHijackEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.BGP.Hijacks.Events(context.TODO(), radar.BGPHijackEventsParams{
		DateEnd:         cloudflare.F(time.Now()),
		DateRange:       cloudflare.F(radar.BGPHijackEventsParamsDateRange7d),
		DateStart:       cloudflare.F(time.Now()),
		EventID:         cloudflare.F(int64(0)),
		Format:          cloudflare.F(radar.BGPHijackEventsParamsFormatJson),
		HijackerASN:     cloudflare.F(int64(0)),
		InvolvedASN:     cloudflare.F(int64(0)),
		InvolvedCountry: cloudflare.F("string"),
		MaxConfidence:   cloudflare.F(int64(0)),
		MinConfidence:   cloudflare.F(int64(0)),
		Page:            cloudflare.F(int64(0)),
		PerPage:         cloudflare.F(int64(0)),
		Prefix:          cloudflare.F("string"),
		SortBy:          cloudflare.F(radar.BGPHijackEventsParamsSortByTime),
		SortOrder:       cloudflare.F(radar.BGPHijackEventsParamsSortOrderDesc),
		VictimASN:       cloudflare.F(int64(0)),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
