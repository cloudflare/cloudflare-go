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

func TestRadarAttackLayer7TimeseriesGroupBitrateWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Bitrate(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsAggInterval1h),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsDirectionOrigin),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsFormatJson),
		IPVersion:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsIPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsIPVersionIPv4, cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsIPVersionIPv6}),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsNormalizationPercentage),
		Protocol:      cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsProtocol{cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsProtocolUdp, cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsProtocolTcp, cloudflare.RadarAttackLayer7TimeseriesGroupBitrateParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupDurationWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Duration(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupDurationParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsAggInterval1h),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsDirectionOrigin),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsFormatJson),
		IPVersion:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsIPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsIPVersionIPv4, cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsIPVersionIPv6}),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsNormalizationPercentage),
		Protocol:      cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsProtocol{cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsProtocolUdp, cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsProtocolTcp, cloudflare.RadarAttackLayer7TimeseriesGroupDurationParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Get(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupGetParams{
		AggInterval: cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupGetParamsAggInterval1h),
		Asn:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupGetParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupGetParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupGetParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupGetParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupGetParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Industry(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsAggInterval1h),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsDirectionOrigin),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsFormatJson),
		IPVersion:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv4, cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv6}),
		LimitPerGroup: cloudflare.F(int64(4)),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIndustryParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.IPVersion(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsAggInterval1h),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsDirectionOrigin),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsFormatJson),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsNormalizationPercentage),
		Protocol:      cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocol{cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocolUdp, cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocolTcp, cloudflare.RadarAttackLayer7TimeseriesGroupIPVersionParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Protocol(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupProtocolParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupProtocolParamsAggInterval1h),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupProtocolParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupProtocolParamsDirectionOrigin),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupProtocolParamsFormatJson),
		IPVersion:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupProtocolParamsIPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupProtocolParamsIPVersionIPv4, cloudflare.RadarAttackLayer7TimeseriesGroupProtocolParamsIPVersionIPv6}),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupProtocolParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupVectorWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Vector(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupVectorParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsAggInterval1h),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsDirectionOrigin),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsFormatJson),
		IPVersion:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsIPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsIPVersionIPv4, cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsIPVersionIPv6}),
		LimitPerGroup: cloudflare.F(int64(4)),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsNormalizationPercentage),
		Protocol:      cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsProtocol{cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsProtocolUdp, cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsProtocolTcp, cloudflare.RadarAttackLayer7TimeseriesGroupVectorParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupVerticalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.Vertical(context.TODO(), cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsAggInterval1h),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange{cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction:     cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsDirectionOrigin),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsFormatJson),
		IPVersion:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersion{cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv4, cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv6}),
		LimitPerGroup: cloudflare.F(int64(4)),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer7TimeseriesGroupVerticalParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
