// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestRadarEmailSecurityTimeseriesGroupARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.ARC(context.TODO(), cloudflare.RadarEmailSecurityTimeseriesGroupARCParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsAggInterval1h),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsDateRange{cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsDateRange1d, cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsDateRange2d, cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsDKIM{cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsDKIMPass, cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsDKIMNone, cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsDKIMFail}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsDMARC{cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsDMARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsDMARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsDMARCFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsSPF{cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsSPFPass, cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsSPFNone, cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsSPFFail}),
		TLSVersion:  cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsTLSVersion{cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupDKIMWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.DKIM(context.TODO(), cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsAggInterval1h),
		ARC:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsARC{cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange{cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange1d, cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange2d, cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsDMARC{cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsDMARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsDMARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsDMARCFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsSPF{cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsSPFPass, cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsSPFNone, cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsSPFFail}),
		TLSVersion:  cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersion{cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupDMARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.DMARC(context.TODO(), cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsAggInterval1h),
		ARC:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsARC{cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange{cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange1d, cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange2d, cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsDKIM{cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsDKIMPass, cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsDKIMNone, cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsDKIMFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsSPF{cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsSPFPass, cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsSPFNone, cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsSPFFail}),
		TLSVersion:  cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersion{cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupMaliciousWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.Malicious(context.TODO(), cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsAggInterval1h),
		ARC:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsARC{cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange{cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange1d, cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange2d, cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsDKIM{cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsDKIMPass, cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsDKIMNone, cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsDKIMFail}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsDMARC{cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsDMARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsDMARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsDMARCFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsSPF{cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsSPFPass, cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsSPFNone, cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsSPFFail}),
		TLSVersion:  cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersion{cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupSpamWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.Spam(context.TODO(), cloudflare.RadarEmailSecurityTimeseriesGroupSpamParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsAggInterval1h),
		ARC:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsARC{cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsDateRange{cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsDateRange1d, cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsDateRange2d, cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsDKIM{cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsDKIMPass, cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsDKIMNone, cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsDKIMFail}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsDMARC{cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsDMARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsDMARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsDMARCFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsSPF{cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsSPFPass, cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsSPFNone, cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsSPFFail}),
		TLSVersion:  cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersion{cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupSPFWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.SPF(context.TODO(), cloudflare.RadarEmailSecurityTimeseriesGroupSPFParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsAggInterval1h),
		ARC:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsARC{cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsDateRange{cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsDateRange1d, cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsDateRange2d, cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsDKIM{cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsDKIMPass, cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsDKIMNone, cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsDKIMFail}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsDMARC{cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsDMARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsDMARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsDMARCFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		TLSVersion:  cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersion{cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupSpoofWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.Spoof(context.TODO(), cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsAggInterval1h),
		ARC:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsARC{cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange{cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange1d, cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange2d, cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsDKIM{cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsDKIMPass, cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsDKIMNone, cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsDKIMFail}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsDMARC{cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsDMARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsDMARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsDMARCFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsSPF{cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsSPFPass, cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsSPFNone, cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsSPFFail}),
		TLSVersion:  cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersion{cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupThreatCategoryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.ThreatCategory(context.TODO(), cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1h),
		ARC:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsARC{cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange{cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange1d, cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange2d, cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDKIM{cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDKIMPass, cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDKIMNone, cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDKIMFail}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDMARC{cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDMARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDMARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsDMARCFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsSPF{cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsSPFPass, cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsSPFNone, cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsSPFFail}),
		TLSVersion:  cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion{cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupTLSVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.TLSVersion(context.TODO(), cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1h),
		ARC:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsARC{cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange{cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange1d, cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange2d, cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsDKIM{cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsDKIMPass, cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsDKIMNone, cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsDKIMFail}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsDMARC{cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsDMARCPass, cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsDMARCNone, cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsDMARCFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsSPF{cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsSPFPass, cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsSPFNone, cloudflare.RadarEmailSecurityTimeseriesGroupTLSVersionParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
