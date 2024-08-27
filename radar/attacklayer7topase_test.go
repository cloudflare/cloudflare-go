// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/radar"
)

func TestAttackLayer7TopAseOriginWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Attacks.Layer7.Top.Ases.Origin(context.TODO(), radar.AttackLayer7TopAseOriginParams{
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TopAseOriginParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TopAseOriginParamsHTTPMethod{radar.AttackLayer7TopAseOriginParamsHTTPMethodGet, radar.AttackLayer7TopAseOriginParamsHTTPMethodPost, radar.AttackLayer7TopAseOriginParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TopAseOriginParamsHTTPVersion{radar.AttackLayer7TopAseOriginParamsHTTPVersionHttPv1, radar.AttackLayer7TopAseOriginParamsHTTPVersionHttPv2, radar.AttackLayer7TopAseOriginParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TopAseOriginParamsIPVersion{radar.AttackLayer7TopAseOriginParamsIPVersionIPv4, radar.AttackLayer7TopAseOriginParamsIPVersionIPv6}),
		Limit:             cloudflare.F(int64(5)),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TopAseOriginParamsMitigationProduct{radar.AttackLayer7TopAseOriginParamsMitigationProductDDoS, radar.AttackLayer7TopAseOriginParamsMitigationProductWAF, radar.AttackLayer7TopAseOriginParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
