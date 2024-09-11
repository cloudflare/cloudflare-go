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

func TestAttackLayer7TopLocationOriginWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.Locations.Origin(context.TODO(), radar.AttackLayer7TopLocationOriginParams{
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TopLocationOriginParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TopLocationOriginParamsHTTPMethod{radar.AttackLayer7TopLocationOriginParamsHTTPMethodGet, radar.AttackLayer7TopLocationOriginParamsHTTPMethodPost, radar.AttackLayer7TopLocationOriginParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TopLocationOriginParamsHTTPVersion{radar.AttackLayer7TopLocationOriginParamsHTTPVersionHttPv1, radar.AttackLayer7TopLocationOriginParamsHTTPVersionHttPv2, radar.AttackLayer7TopLocationOriginParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TopLocationOriginParamsIPVersion{radar.AttackLayer7TopLocationOriginParamsIPVersionIPv4, radar.AttackLayer7TopLocationOriginParamsIPVersionIPv6}),
		Limit:             cloudflare.F(int64(5)),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TopLocationOriginParamsMitigationProduct{radar.AttackLayer7TopLocationOriginParamsMitigationProductDDoS, radar.AttackLayer7TopLocationOriginParamsMitigationProductWAF, radar.AttackLayer7TopLocationOriginParamsMitigationProductBotManagement}),
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

func TestAttackLayer7TopLocationTargetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.Locations.Target(context.TODO(), radar.AttackLayer7TopLocationTargetParams{
		Continent:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TopLocationTargetParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TopLocationTargetParamsHTTPMethod{radar.AttackLayer7TopLocationTargetParamsHTTPMethodGet, radar.AttackLayer7TopLocationTargetParamsHTTPMethodPost, radar.AttackLayer7TopLocationTargetParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TopLocationTargetParamsHTTPVersion{radar.AttackLayer7TopLocationTargetParamsHTTPVersionHttPv1, radar.AttackLayer7TopLocationTargetParamsHTTPVersionHttPv2, radar.AttackLayer7TopLocationTargetParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TopLocationTargetParamsIPVersion{radar.AttackLayer7TopLocationTargetParamsIPVersionIPv4, radar.AttackLayer7TopLocationTargetParamsIPVersionIPv6}),
		Limit:             cloudflare.F(int64(5)),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TopLocationTargetParamsMitigationProduct{radar.AttackLayer7TopLocationTargetParamsMitigationProductDDoS, radar.AttackLayer7TopLocationTargetParamsMitigationProductWAF, radar.AttackLayer7TopLocationTargetParamsMitigationProductBotManagement}),
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
