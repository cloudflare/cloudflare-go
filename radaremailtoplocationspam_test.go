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

func TestRadarEmailTopLocationSpamGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.Emails.Tops.Locations.Spams.Get(
		context.TODO(),
		cloudflare.RadarEmailTopLocationSpamGetParamsSpamSpam,
		cloudflare.RadarEmailTopLocationSpamGetParams{
			Arc:       cloudflare.F([]cloudflare.RadarEmailTopLocationSpamGetParamsArc{cloudflare.RadarEmailTopLocationSpamGetParamsArcPass, cloudflare.RadarEmailTopLocationSpamGetParamsArcNone, cloudflare.RadarEmailTopLocationSpamGetParamsArcFail}),
			ASN:       cloudflare.F([]string{"15169", "15169", "15169"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange: cloudflare.F([]cloudflare.RadarEmailTopLocationSpamGetParamsDateRange{cloudflare.RadarEmailTopLocationSpamGetParamsDateRange1d, cloudflare.RadarEmailTopLocationSpamGetParamsDateRange7d, cloudflare.RadarEmailTopLocationSpamGetParamsDateRange14d}),
			DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Dkim:      cloudflare.F([]cloudflare.RadarEmailTopLocationSpamGetParamsDkim{cloudflare.RadarEmailTopLocationSpamGetParamsDkimPass, cloudflare.RadarEmailTopLocationSpamGetParamsDkimNone, cloudflare.RadarEmailTopLocationSpamGetParamsDkimFail}),
			Dmarc:     cloudflare.F([]cloudflare.RadarEmailTopLocationSpamGetParamsDmarc{cloudflare.RadarEmailTopLocationSpamGetParamsDmarcPass, cloudflare.RadarEmailTopLocationSpamGetParamsDmarcNone, cloudflare.RadarEmailTopLocationSpamGetParamsDmarcFail}),
			Format:    cloudflare.F(cloudflare.RadarEmailTopLocationSpamGetParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:      cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			Spf:       cloudflare.F([]cloudflare.RadarEmailTopLocationSpamGetParamsSpf{cloudflare.RadarEmailTopLocationSpamGetParamsSpfPass, cloudflare.RadarEmailTopLocationSpamGetParamsSpfNone, cloudflare.RadarEmailTopLocationSpamGetParamsSpfFail}),
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
