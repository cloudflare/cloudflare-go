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

func TestRadarEmailTopAseMaliciousGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.Emails.Tops.Ases.Malicious.Get(
		context.TODO(),
		cloudflare.RadarEmailTopAseMaliciousGetParamsMaliciousMalicious,
		cloudflare.RadarEmailTopAseMaliciousGetParams{
			Arc:       cloudflare.F([]cloudflare.RadarEmailTopAseMaliciousGetParamsArc{cloudflare.RadarEmailTopAseMaliciousGetParamsArcPass, cloudflare.RadarEmailTopAseMaliciousGetParamsArcNone, cloudflare.RadarEmailTopAseMaliciousGetParamsArcFail}),
			ASN:       cloudflare.F([]string{"15169", "15169", "15169"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange: cloudflare.F([]cloudflare.RadarEmailTopAseMaliciousGetParamsDateRange{cloudflare.RadarEmailTopAseMaliciousGetParamsDateRange1d, cloudflare.RadarEmailTopAseMaliciousGetParamsDateRange7d, cloudflare.RadarEmailTopAseMaliciousGetParamsDateRange14d}),
			DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Dkim:      cloudflare.F([]cloudflare.RadarEmailTopAseMaliciousGetParamsDkim{cloudflare.RadarEmailTopAseMaliciousGetParamsDkimPass, cloudflare.RadarEmailTopAseMaliciousGetParamsDkimNone, cloudflare.RadarEmailTopAseMaliciousGetParamsDkimFail}),
			Dmarc:     cloudflare.F([]cloudflare.RadarEmailTopAseMaliciousGetParamsDmarc{cloudflare.RadarEmailTopAseMaliciousGetParamsDmarcPass, cloudflare.RadarEmailTopAseMaliciousGetParamsDmarcNone, cloudflare.RadarEmailTopAseMaliciousGetParamsDmarcFail}),
			Format:    cloudflare.F(cloudflare.RadarEmailTopAseMaliciousGetParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:      cloudflare.F([]string{"main_series", "main_series", "main_series"}),
			Spf:       cloudflare.F([]cloudflare.RadarEmailTopAseMaliciousGetParamsSpf{cloudflare.RadarEmailTopAseMaliciousGetParamsSpfPass, cloudflare.RadarEmailTopAseMaliciousGetParamsSpfNone, cloudflare.RadarEmailTopAseMaliciousGetParamsSpfFail}),
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
