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

func TestRadarAttackLayer7TopAttackListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Tops.Attacks.List(context.TODO(), cloudflare.RadarAttackLayer7TopAttackListParams{
		ASN:              cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:          cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:        cloudflare.F([]cloudflare.RadarAttackLayer7TopAttackListParamsDateRange{cloudflare.RadarAttackLayer7TopAttackListParamsDateRange1d, cloudflare.RadarAttackLayer7TopAttackListParamsDateRange2d, cloudflare.RadarAttackLayer7TopAttackListParamsDateRange7d}),
		DateStart:        cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:           cloudflare.F(cloudflare.RadarAttackLayer7TopAttackListParamsFormatJson),
		Limit:            cloudflare.F(int64(5)),
		LimitDirection:   cloudflare.F(cloudflare.RadarAttackLayer7TopAttackListParamsLimitDirectionOrigin),
		LimitPerLocation: cloudflare.F(int64(10)),
		Location:         cloudflare.F([]string{"string", "string", "string"}),
		Magnitude:        cloudflare.F(cloudflare.RadarAttackLayer7TopAttackListParamsMagnitudeMitigatedRequests),
		Name:             cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
