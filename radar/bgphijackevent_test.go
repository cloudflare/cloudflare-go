// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/radar"
)

func TestBGPHijackEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.BGP.Hijacks.Events.List(context.TODO(), radar.BGPHijackEventListParams{
		DateEnd:         cloudflare.F(time.Now()),
		DateRange:       cloudflare.F("7d"),
		DateStart:       cloudflare.F(time.Now()),
		EventID:         cloudflare.F(int64(0)),
		Format:          cloudflare.F(radar.BGPHijackEventListParamsFormatJson),
		HijackerASN:     cloudflare.F(int64(0)),
		InvolvedASN:     cloudflare.F(int64(0)),
		InvolvedCountry: cloudflare.F("involvedCountry"),
		MaxConfidence:   cloudflare.F(int64(0)),
		MinConfidence:   cloudflare.F(int64(0)),
		Page:            cloudflare.F(int64(0)),
		PerPage:         cloudflare.F(int64(0)),
		Prefix:          cloudflare.F("1.1.1.0/24"),
		SortBy:          cloudflare.F(radar.BGPHijackEventListParamsSortByID),
		SortOrder:       cloudflare.F(radar.BGPHijackEventListParamsSortOrderAsc),
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
