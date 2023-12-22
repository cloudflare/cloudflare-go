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

func TestRadarEmailTopLocationDmarcGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.Emails.Tops.Locations.Dmarcs.Get(
		context.TODO(),
		cloudflare.RadarEmailTopLocationDmarcGetParamsDmarcPass,
		cloudflare.RadarEmailTopLocationDmarcGetParams{
			Arc:       cloudflare.F([]cloudflare.RadarEmailTopLocationDmarcGetParamsArc{cloudflare.RadarEmailTopLocationDmarcGetParamsArcPass, cloudflare.RadarEmailTopLocationDmarcGetParamsArcNone, cloudflare.RadarEmailTopLocationDmarcGetParamsArcFail}),
			ASN:       cloudflare.F([]string{"15169", "15169", "15169"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange: cloudflare.F([]cloudflare.RadarEmailTopLocationDmarcGetParamsDateRange{cloudflare.RadarEmailTopLocationDmarcGetParamsDateRange1d, cloudflare.RadarEmailTopLocationDmarcGetParamsDateRange7d, cloudflare.RadarEmailTopLocationDmarcGetParamsDateRange14d}),
			DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Dkim:      cloudflare.F([]cloudflare.RadarEmailTopLocationDmarcGetParamsDkim{cloudflare.RadarEmailTopLocationDmarcGetParamsDkimPass, cloudflare.RadarEmailTopLocationDmarcGetParamsDkimNone, cloudflare.RadarEmailTopLocationDmarcGetParamsDkimFail}),
			Format:    cloudflare.F(cloudflare.RadarEmailTopLocationDmarcGetParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:      cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			Spf:       cloudflare.F([]cloudflare.RadarEmailTopLocationDmarcGetParamsSpf{cloudflare.RadarEmailTopLocationDmarcGetParamsSpfPass, cloudflare.RadarEmailTopLocationDmarcGetParamsSpfNone, cloudflare.RadarEmailTopLocationDmarcGetParamsSpfFail}),
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
