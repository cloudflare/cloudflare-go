// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/radar"
)

func TestBGPRPKIASPAChangesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.BGP.RPKI.ASPA.Changes(context.TODO(), radar.BGPRPKIASPAChangesParams{
		ASN:            cloudflare.F(int64(13335)),
		DateEnd:        cloudflare.F(time.Now()),
		DateStart:      cloudflare.F(time.Now()),
		Format:         cloudflare.F(radar.BgprpkiaspaChangesParamsFormatJson),
		IncludeASNInfo: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBGPRPKIASPASnapshotWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.BGP.RPKI.ASPA.Snapshot(context.TODO(), radar.BGPRPKIASPASnapshotParams{
		CustomerASN:    cloudflare.F(int64(13335)),
		Date:           cloudflare.F(time.Now()),
		Format:         cloudflare.F(radar.BgprpkiaspaSnapshotParamsFormatJson),
		IncludeASNInfo: cloudflare.F(true),
		ProviderASN:    cloudflare.F(int64(174)),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBGPRPKIASPATimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.BGP.RPKI.ASPA.Timeseries(context.TODO(), radar.BGPRPKIASPATimeseriesParams{
		DateEnd:   cloudflare.F(time.Now()),
		DateStart: cloudflare.F(time.Now()),
		Format:    cloudflare.F(radar.BgprpkiaspaTimeseriesParamsFormatJson),
		Location:  cloudflare.F([]string{"string"}),
		Name:      cloudflare.F([]string{"main_series"}),
		Rir:       cloudflare.F([]radar.BgprpkiaspaTimeseriesParamsRir{radar.BgprpkiaspaTimeseriesParamsRirRipeNcc}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
