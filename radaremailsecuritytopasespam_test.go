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

func TestRadarEmailSecurityTopAseSpamGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Ases.Spam.Get(
		context.TODO(),
		cloudflare.RadarEmailSecurityTopAseSpamGetParamsSpamSpam,
		cloudflare.RadarEmailSecurityTopAseSpamGetParams{
			Arc:       cloudflare.F([]cloudflare.RadarEmailSecurityTopAseSpamGetParamsArc{cloudflare.RadarEmailSecurityTopAseSpamGetParamsArcPass, cloudflare.RadarEmailSecurityTopAseSpamGetParamsArcNone, cloudflare.RadarEmailSecurityTopAseSpamGetParamsArcFail}),
			ASN:       cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange: cloudflare.F([]cloudflare.RadarEmailSecurityTopAseSpamGetParamsDateRange{cloudflare.RadarEmailSecurityTopAseSpamGetParamsDateRange1d, cloudflare.RadarEmailSecurityTopAseSpamGetParamsDateRange2d, cloudflare.RadarEmailSecurityTopAseSpamGetParamsDateRange7d}),
			DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Dkim:      cloudflare.F([]cloudflare.RadarEmailSecurityTopAseSpamGetParamsDkim{cloudflare.RadarEmailSecurityTopAseSpamGetParamsDkimPass, cloudflare.RadarEmailSecurityTopAseSpamGetParamsDkimNone, cloudflare.RadarEmailSecurityTopAseSpamGetParamsDkimFail}),
			Dmarc:     cloudflare.F([]cloudflare.RadarEmailSecurityTopAseSpamGetParamsDmarc{cloudflare.RadarEmailSecurityTopAseSpamGetParamsDmarcPass, cloudflare.RadarEmailSecurityTopAseSpamGetParamsDmarcNone, cloudflare.RadarEmailSecurityTopAseSpamGetParamsDmarcFail}),
			Format:    cloudflare.F(cloudflare.RadarEmailSecurityTopAseSpamGetParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"string", "string", "string"}),
			Name:      cloudflare.F([]string{"string", "string", "string"}),
			Spf:       cloudflare.F([]cloudflare.RadarEmailSecurityTopAseSpamGetParamsSpf{cloudflare.RadarEmailSecurityTopAseSpamGetParamsSpfPass, cloudflare.RadarEmailSecurityTopAseSpamGetParamsSpfNone, cloudflare.RadarEmailSecurityTopAseSpamGetParamsSpfFail}),
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
