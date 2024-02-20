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

func TestRadarBGPTopAseListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.BGP.Tops.Ases.List(context.TODO(), cloudflare.RadarBGPTopAseListParams{
		Asn:        cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]cloudflare.RadarBGPTopAseListParamsDateRange{cloudflare.RadarBGPTopAseListParamsDateRange1d, cloudflare.RadarBGPTopAseListParamsDateRange2d, cloudflare.RadarBGPTopAseListParamsDateRange7d}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:     cloudflare.F(cloudflare.RadarBGPTopAseListParamsFormatJson),
		Limit:      cloudflare.F(int64(5)),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		Prefix:     cloudflare.F([]string{"string", "string", "string"}),
		UpdateType: cloudflare.F([]cloudflare.RadarBGPTopAseListParamsUpdateType{cloudflare.RadarBGPTopAseListParamsUpdateTypeAnnouncement, cloudflare.RadarBGPTopAseListParamsUpdateTypeWithdrawal}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
