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

func TestRadarEmailSecurityTopLocationBySpamListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Locations.BySpam.List(
		context.TODO(),
		cloudflare.RadarEmailSecurityTopLocationBySpamListParamsSpamSpam,
		cloudflare.RadarEmailSecurityTopLocationBySpamListParams{
			Arc:       cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationBySpamListParamsArc{cloudflare.RadarEmailSecurityTopLocationBySpamListParamsArcPass, cloudflare.RadarEmailSecurityTopLocationBySpamListParamsArcNone, cloudflare.RadarEmailSecurityTopLocationBySpamListParamsArcFail}),
			ASN:       cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange: cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationBySpamListParamsDateRange{cloudflare.RadarEmailSecurityTopLocationBySpamListParamsDateRange1d, cloudflare.RadarEmailSecurityTopLocationBySpamListParamsDateRange2d, cloudflare.RadarEmailSecurityTopLocationBySpamListParamsDateRange7d}),
			DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Dkim:      cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationBySpamListParamsDkim{cloudflare.RadarEmailSecurityTopLocationBySpamListParamsDkimPass, cloudflare.RadarEmailSecurityTopLocationBySpamListParamsDkimNone, cloudflare.RadarEmailSecurityTopLocationBySpamListParamsDkimFail}),
			Dmarc:     cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationBySpamListParamsDmarc{cloudflare.RadarEmailSecurityTopLocationBySpamListParamsDmarcPass, cloudflare.RadarEmailSecurityTopLocationBySpamListParamsDmarcNone, cloudflare.RadarEmailSecurityTopLocationBySpamListParamsDmarcFail}),
			Format:    cloudflare.F(cloudflare.RadarEmailSecurityTopLocationBySpamListParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"string", "string", "string"}),
			Name:      cloudflare.F([]string{"string", "string", "string"}),
			Spf:       cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationBySpamListParamsSpf{cloudflare.RadarEmailSecurityTopLocationBySpamListParamsSpfPass, cloudflare.RadarEmailSecurityTopLocationBySpamListParamsSpfNone, cloudflare.RadarEmailSecurityTopLocationBySpamListParamsSpfFail}),
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
