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

func TestRadarEmailSecurityTopLocationDMARCGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Email.Security.Top.Locations.DMARC.Get(
		context.TODO(),
		cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsDMARCPass,
		cloudflare.RadarEmailSecurityTopLocationDMARCGetParams{
			ARC:       cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsARC{cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsARCPass, cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsARCNone, cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsARCFail}),
			ASN:       cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange: cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsDateRange{cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsDateRange1d, cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsDateRange2d, cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsDateRange7d}),
			DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DKIM:      cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsDKIM{cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsDKIMPass, cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsDKIMNone, cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsDKIMFail}),
			Format:    cloudflare.F(cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"string", "string", "string"}),
			Name:      cloudflare.F([]string{"string", "string", "string"}),
			SPF:       cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsSPF{cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsSPFPass, cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsSPFNone, cloudflare.RadarEmailSecurityTopLocationDMARCGetParamsSPFFail}),
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
