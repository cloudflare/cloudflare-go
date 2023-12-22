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

func TestRadarEmailTopAseSpfGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.Emails.Tops.Ases.Spfs.Get(
		context.TODO(),
		cloudflare.RadarEmailTopAseSpfGetParamsSpfPass,
		cloudflare.RadarEmailTopAseSpfGetParams{
			Arc:       cloudflare.F([]cloudflare.RadarEmailTopAseSpfGetParamsArc{cloudflare.RadarEmailTopAseSpfGetParamsArcPass, cloudflare.RadarEmailTopAseSpfGetParamsArcNone, cloudflare.RadarEmailTopAseSpfGetParamsArcFail}),
			ASN:       cloudflare.F([]string{"15169", "15169", "15169"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange: cloudflare.F([]cloudflare.RadarEmailTopAseSpfGetParamsDateRange{cloudflare.RadarEmailTopAseSpfGetParamsDateRange1d, cloudflare.RadarEmailTopAseSpfGetParamsDateRange7d, cloudflare.RadarEmailTopAseSpfGetParamsDateRange14d}),
			DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Dkim:      cloudflare.F([]cloudflare.RadarEmailTopAseSpfGetParamsDkim{cloudflare.RadarEmailTopAseSpfGetParamsDkimPass, cloudflare.RadarEmailTopAseSpfGetParamsDkimNone, cloudflare.RadarEmailTopAseSpfGetParamsDkimFail}),
			Dmarc:     cloudflare.F([]cloudflare.RadarEmailTopAseSpfGetParamsDmarc{cloudflare.RadarEmailTopAseSpfGetParamsDmarcPass, cloudflare.RadarEmailTopAseSpfGetParamsDmarcNone, cloudflare.RadarEmailTopAseSpfGetParamsDmarcFail}),
			Format:    cloudflare.F(cloudflare.RadarEmailTopAseSpfGetParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
			Name:      cloudflare.F([]string{"main_series", "main_series", "main_series"}),
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
