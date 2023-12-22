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

func TestRadarEmailTopAseDkimGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.Emails.Tops.Ases.Dkims.Get(
		context.TODO(),
		cloudflare.RadarEmailTopAseDkimGetParamsDkimPass,
		cloudflare.RadarEmailTopAseDkimGetParams{
			Arc:       cloudflare.F([]cloudflare.RadarEmailTopAseDkimGetParamsArc{cloudflare.RadarEmailTopAseDkimGetParamsArcPass, cloudflare.RadarEmailTopAseDkimGetParamsArcNone, cloudflare.RadarEmailTopAseDkimGetParamsArcFail}),
			ASN:       cloudflare.F([]string{"15169", "15169", "15169"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange: cloudflare.F([]cloudflare.RadarEmailTopAseDkimGetParamsDateRange{cloudflare.RadarEmailTopAseDkimGetParamsDateRange1d, cloudflare.RadarEmailTopAseDkimGetParamsDateRange7d, cloudflare.RadarEmailTopAseDkimGetParamsDateRange14d}),
			DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Dmarc:     cloudflare.F([]cloudflare.RadarEmailTopAseDkimGetParamsDmarc{cloudflare.RadarEmailTopAseDkimGetParamsDmarcPass, cloudflare.RadarEmailTopAseDkimGetParamsDmarcNone, cloudflare.RadarEmailTopAseDkimGetParamsDmarcFail}),
			Format:    cloudflare.F(cloudflare.RadarEmailTopAseDkimGetParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:      cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			Spf:       cloudflare.F([]cloudflare.RadarEmailTopAseDkimGetParamsSpf{cloudflare.RadarEmailTopAseDkimGetParamsSpfPass, cloudflare.RadarEmailTopAseDkimGetParamsSpfNone, cloudflare.RadarEmailTopAseDkimGetParamsSpfFail}),
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
