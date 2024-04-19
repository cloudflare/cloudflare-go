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

func TestAttackLayer3TopLocationOriginWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Attacks.Layer3.Top.Locations.Origin(context.TODO(), radar.AttackLayer3TopLocationOriginParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]radar.AttackLayer3TopLocationOriginParamsDateRange{radar.AttackLayer3TopLocationOriginParamsDateRange1d, radar.AttackLayer3TopLocationOriginParamsDateRange2d, radar.AttackLayer3TopLocationOriginParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(radar.AttackLayer3TopLocationOriginParamsFormatJson),
		IPVersion: cloudflare.F([]radar.AttackLayer3TopLocationOriginParamsIPVersion{radar.AttackLayer3TopLocationOriginParamsIPVersionIPv4, radar.AttackLayer3TopLocationOriginParamsIPVersionIPv6}),
		Limit:     cloudflare.F(int64(5)),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]radar.AttackLayer3TopLocationOriginParamsProtocol{radar.AttackLayer3TopLocationOriginParamsProtocolUdp, radar.AttackLayer3TopLocationOriginParamsProtocolTCP, radar.AttackLayer3TopLocationOriginParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TopLocationTargetWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Attacks.Layer3.Top.Locations.Target(context.TODO(), radar.AttackLayer3TopLocationTargetParams{
		Continent: cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]radar.AttackLayer3TopLocationTargetParamsDateRange{radar.AttackLayer3TopLocationTargetParamsDateRange1d, radar.AttackLayer3TopLocationTargetParamsDateRange2d, radar.AttackLayer3TopLocationTargetParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(radar.AttackLayer3TopLocationTargetParamsFormatJson),
		IPVersion: cloudflare.F([]radar.AttackLayer3TopLocationTargetParamsIPVersion{radar.AttackLayer3TopLocationTargetParamsIPVersionIPv4, radar.AttackLayer3TopLocationTargetParamsIPVersionIPv6}),
		Limit:     cloudflare.F(int64(5)),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]radar.AttackLayer3TopLocationTargetParamsProtocol{radar.AttackLayer3TopLocationTargetParamsProtocolUdp, radar.AttackLayer3TopLocationTargetParamsProtocolTCP, radar.AttackLayer3TopLocationTargetParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
