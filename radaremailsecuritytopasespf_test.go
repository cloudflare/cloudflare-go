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

func TestRadarEmailSecurityTopAseSpfGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Radar.Email.Security.Top.Ases.Spf.Get(
		context.TODO(),
		cloudflare.RadarEmailSecurityTopAseSpfGetParamsSpfPass,
		cloudflare.RadarEmailSecurityTopAseSpfGetParams{
			Arc:       cloudflare.F([]cloudflare.RadarEmailSecurityTopAseSpfGetParamsArc{cloudflare.RadarEmailSecurityTopAseSpfGetParamsArcPass, cloudflare.RadarEmailSecurityTopAseSpfGetParamsArcNone, cloudflare.RadarEmailSecurityTopAseSpfGetParamsArcFail}),
			ASN:       cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange: cloudflare.F([]cloudflare.RadarEmailSecurityTopAseSpfGetParamsDateRange{cloudflare.RadarEmailSecurityTopAseSpfGetParamsDateRange1d, cloudflare.RadarEmailSecurityTopAseSpfGetParamsDateRange2d, cloudflare.RadarEmailSecurityTopAseSpfGetParamsDateRange7d}),
			DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Dkim:      cloudflare.F([]cloudflare.RadarEmailSecurityTopAseSpfGetParamsDkim{cloudflare.RadarEmailSecurityTopAseSpfGetParamsDkimPass, cloudflare.RadarEmailSecurityTopAseSpfGetParamsDkimNone, cloudflare.RadarEmailSecurityTopAseSpfGetParamsDkimFail}),
			Dmarc:     cloudflare.F([]cloudflare.RadarEmailSecurityTopAseSpfGetParamsDmarc{cloudflare.RadarEmailSecurityTopAseSpfGetParamsDmarcPass, cloudflare.RadarEmailSecurityTopAseSpfGetParamsDmarcNone, cloudflare.RadarEmailSecurityTopAseSpfGetParamsDmarcFail}),
			Format:    cloudflare.F(cloudflare.RadarEmailSecurityTopAseSpfGetParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"string", "string", "string"}),
			Name:      cloudflare.F([]string{"string", "string", "string"}),
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
