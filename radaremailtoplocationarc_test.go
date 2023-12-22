// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestRadarEmailTopLocationArcGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Radars.Emails.Tops.Locations.Arcs.Get(
		context.TODO(),
		cloudflare.RadarEmailTopLocationArcGetParamsArcPass,
		cloudflare.RadarEmailTopLocationArcGetParams{
			ASN:       cloudflare.F([]string{"15169", "15169", "15169"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange: cloudflare.F([]cloudflare.RadarEmailTopLocationArcGetParamsDateRange{cloudflare.RadarEmailTopLocationArcGetParamsDateRange1d, cloudflare.RadarEmailTopLocationArcGetParamsDateRange7d, cloudflare.RadarEmailTopLocationArcGetParamsDateRange14d}),
			DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Dkim:      cloudflare.F([]cloudflare.RadarEmailTopLocationArcGetParamsDkim{cloudflare.RadarEmailTopLocationArcGetParamsDkimPass, cloudflare.RadarEmailTopLocationArcGetParamsDkimNone, cloudflare.RadarEmailTopLocationArcGetParamsDkimFail}),
			Dmarc:     cloudflare.F([]cloudflare.RadarEmailTopLocationArcGetParamsDmarc{cloudflare.RadarEmailTopLocationArcGetParamsDmarcPass, cloudflare.RadarEmailTopLocationArcGetParamsDmarcNone, cloudflare.RadarEmailTopLocationArcGetParamsDmarcFail}),
			Format:    cloudflare.F(cloudflare.RadarEmailTopLocationArcGetParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:      cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			Spf:       cloudflare.F([]cloudflare.RadarEmailTopLocationArcGetParamsSpf{cloudflare.RadarEmailTopLocationArcGetParamsSpfPass, cloudflare.RadarEmailTopLocationArcGetParamsSpfNone, cloudflare.RadarEmailTopLocationArcGetParamsSpfFail}),
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
