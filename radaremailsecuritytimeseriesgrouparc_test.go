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

func TestRadarEmailSecurityTimeseriesGroupArcListWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Radar.Email.Security.TimeseriesGroups.Arcs.List(context.TODO(), cloudflare.RadarEmailSecurityTimeseriesGroupArcListParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsAggInterval1h),
		Asn:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsDateRange{cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsDateRange1d, cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsDateRange2d, cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsDKIM{cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsDKIMPass, cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsDKIMNone, cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsDKIMFail}),
		Dmarc:       cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsDmarc{cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsDmarcPass, cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsDmarcNone, cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsDmarcFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsSPF{cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsSPFPass, cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsSPFNone, cloudflare.RadarEmailSecurityTimeseriesGroupArcListParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
