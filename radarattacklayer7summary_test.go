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

func TestRadarAttackLayer7SummaryBitrateWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.Bitrate(context.TODO(), cloudflare.RadarAttackLayer7SummaryBitrateParams{
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryBitrateParamsDateRange{cloudflare.RadarAttackLayer7SummaryBitrateParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryBitrateParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryBitrateParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(cloudflare.RadarAttackLayer7SummaryBitrateParamsDirectionOrigin),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer7SummaryBitrateParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryBitrateParamsIPVersion{cloudflare.RadarAttackLayer7SummaryBitrateParamsIPVersionIPv4, cloudflare.RadarAttackLayer7SummaryBitrateParamsIPVersionIPv6}),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]cloudflare.RadarAttackLayer7SummaryBitrateParamsProtocol{cloudflare.RadarAttackLayer7SummaryBitrateParamsProtocolUdp, cloudflare.RadarAttackLayer7SummaryBitrateParamsProtocolTcp, cloudflare.RadarAttackLayer7SummaryBitrateParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryDurationWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.Duration(context.TODO(), cloudflare.RadarAttackLayer7SummaryDurationParams{
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryDurationParamsDateRange{cloudflare.RadarAttackLayer7SummaryDurationParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryDurationParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryDurationParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(cloudflare.RadarAttackLayer7SummaryDurationParamsDirectionOrigin),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer7SummaryDurationParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryDurationParamsIPVersion{cloudflare.RadarAttackLayer7SummaryDurationParamsIPVersionIPv4, cloudflare.RadarAttackLayer7SummaryDurationParamsIPVersionIPv6}),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]cloudflare.RadarAttackLayer7SummaryDurationParamsProtocol{cloudflare.RadarAttackLayer7SummaryDurationParamsProtocolUdp, cloudflare.RadarAttackLayer7SummaryDurationParamsProtocolTcp, cloudflare.RadarAttackLayer7SummaryDurationParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.Get(context.TODO(), cloudflare.RadarAttackLayer7SummaryGetParams{
		Asn:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryGetParamsDateRange{cloudflare.RadarAttackLayer7SummaryGetParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryGetParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryGetParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer7SummaryGetParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.IPVersion(context.TODO(), cloudflare.RadarAttackLayer7SummaryIPVersionParams{
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryIPVersionParamsDateRange{cloudflare.RadarAttackLayer7SummaryIPVersionParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryIPVersionParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryIPVersionParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(cloudflare.RadarAttackLayer7SummaryIPVersionParamsDirectionOrigin),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer7SummaryIPVersionParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]cloudflare.RadarAttackLayer7SummaryIPVersionParamsProtocol{cloudflare.RadarAttackLayer7SummaryIPVersionParamsProtocolUdp, cloudflare.RadarAttackLayer7SummaryIPVersionParamsProtocolTcp, cloudflare.RadarAttackLayer7SummaryIPVersionParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.Protocol(context.TODO(), cloudflare.RadarAttackLayer7SummaryProtocolParams{
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryProtocolParamsDateRange{cloudflare.RadarAttackLayer7SummaryProtocolParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryProtocolParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryProtocolParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(cloudflare.RadarAttackLayer7SummaryProtocolParamsDirectionOrigin),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer7SummaryProtocolParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryProtocolParamsIPVersion{cloudflare.RadarAttackLayer7SummaryProtocolParamsIPVersionIPv4, cloudflare.RadarAttackLayer7SummaryProtocolParamsIPVersionIPv6}),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryVectorWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.Vector(context.TODO(), cloudflare.RadarAttackLayer7SummaryVectorParams{
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryVectorParamsDateRange{cloudflare.RadarAttackLayer7SummaryVectorParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryVectorParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryVectorParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: cloudflare.F(cloudflare.RadarAttackLayer7SummaryVectorParamsDirectionOrigin),
		Format:    cloudflare.F(cloudflare.RadarAttackLayer7SummaryVectorParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryVectorParamsIPVersion{cloudflare.RadarAttackLayer7SummaryVectorParamsIPVersionIPv4, cloudflare.RadarAttackLayer7SummaryVectorParamsIPVersionIPv6}),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Protocol:  cloudflare.F([]cloudflare.RadarAttackLayer7SummaryVectorParamsProtocol{cloudflare.RadarAttackLayer7SummaryVectorParamsProtocolUdp, cloudflare.RadarAttackLayer7SummaryVectorParamsProtocolTcp, cloudflare.RadarAttackLayer7SummaryVectorParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
