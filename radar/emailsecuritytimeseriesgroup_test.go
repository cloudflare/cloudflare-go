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

func TestEmailSecurityTimeseriesGroupARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.ARC(context.TODO(), radar.EmailSecurityTimeseriesGroupARCParams{
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupARCParamsAggInterval1h),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailSecurityTimeseriesGroupARCParamsDateRange{radar.EmailSecurityTimeseriesGroupARCParamsDateRange1d, radar.EmailSecurityTimeseriesGroupARCParamsDateRange2d, radar.EmailSecurityTimeseriesGroupARCParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupARCParamsDKIM{radar.EmailSecurityTimeseriesGroupARCParamsDKIMPass, radar.EmailSecurityTimeseriesGroupARCParamsDKIMNone, radar.EmailSecurityTimeseriesGroupARCParamsDKIMFail}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupARCParamsDMARC{radar.EmailSecurityTimeseriesGroupARCParamsDMARCPass, radar.EmailSecurityTimeseriesGroupARCParamsDMARCNone, radar.EmailSecurityTimeseriesGroupARCParamsDMARCFail}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupARCParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupARCParamsSPF{radar.EmailSecurityTimeseriesGroupARCParamsSPFPass, radar.EmailSecurityTimeseriesGroupARCParamsSPFNone, radar.EmailSecurityTimeseriesGroupARCParamsSPFFail}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupARCParamsTLSVersion{radar.EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupDKIMWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.DKIM(context.TODO(), radar.EmailSecurityTimeseriesGroupDKIMParams{
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupDKIMParamsAggInterval1h),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupDKIMParamsARC{radar.EmailSecurityTimeseriesGroupDKIMParamsARCPass, radar.EmailSecurityTimeseriesGroupDKIMParamsARCNone, radar.EmailSecurityTimeseriesGroupDKIMParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailSecurityTimeseriesGroupDKIMParamsDateRange{radar.EmailSecurityTimeseriesGroupDKIMParamsDateRange1d, radar.EmailSecurityTimeseriesGroupDKIMParamsDateRange2d, radar.EmailSecurityTimeseriesGroupDKIMParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupDKIMParamsDMARC{radar.EmailSecurityTimeseriesGroupDKIMParamsDMARCPass, radar.EmailSecurityTimeseriesGroupDKIMParamsDMARCNone, radar.EmailSecurityTimeseriesGroupDKIMParamsDMARCFail}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupDKIMParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupDKIMParamsSPF{radar.EmailSecurityTimeseriesGroupDKIMParamsSPFPass, radar.EmailSecurityTimeseriesGroupDKIMParamsSPFNone, radar.EmailSecurityTimeseriesGroupDKIMParamsSPFFail}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupDKIMParamsTLSVersion{radar.EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupDMARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.DMARC(context.TODO(), radar.EmailSecurityTimeseriesGroupDMARCParams{
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupDMARCParamsAggInterval1h),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupDMARCParamsARC{radar.EmailSecurityTimeseriesGroupDMARCParamsARCPass, radar.EmailSecurityTimeseriesGroupDMARCParamsARCNone, radar.EmailSecurityTimeseriesGroupDMARCParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailSecurityTimeseriesGroupDMARCParamsDateRange{radar.EmailSecurityTimeseriesGroupDMARCParamsDateRange1d, radar.EmailSecurityTimeseriesGroupDMARCParamsDateRange2d, radar.EmailSecurityTimeseriesGroupDMARCParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupDMARCParamsDKIM{radar.EmailSecurityTimeseriesGroupDMARCParamsDKIMPass, radar.EmailSecurityTimeseriesGroupDMARCParamsDKIMNone, radar.EmailSecurityTimeseriesGroupDMARCParamsDKIMFail}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupDMARCParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupDMARCParamsSPF{radar.EmailSecurityTimeseriesGroupDMARCParamsSPFPass, radar.EmailSecurityTimeseriesGroupDMARCParamsSPFNone, radar.EmailSecurityTimeseriesGroupDMARCParamsSPFFail}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupDMARCParamsTLSVersion{radar.EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupMaliciousWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.Malicious(context.TODO(), radar.EmailSecurityTimeseriesGroupMaliciousParams{
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupMaliciousParamsAggInterval1h),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsARC{radar.EmailSecurityTimeseriesGroupMaliciousParamsARCPass, radar.EmailSecurityTimeseriesGroupMaliciousParamsARCNone, radar.EmailSecurityTimeseriesGroupMaliciousParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsDateRange{radar.EmailSecurityTimeseriesGroupMaliciousParamsDateRange1d, radar.EmailSecurityTimeseriesGroupMaliciousParamsDateRange2d, radar.EmailSecurityTimeseriesGroupMaliciousParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsDKIM{radar.EmailSecurityTimeseriesGroupMaliciousParamsDKIMPass, radar.EmailSecurityTimeseriesGroupMaliciousParamsDKIMNone, radar.EmailSecurityTimeseriesGroupMaliciousParamsDKIMFail}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsDMARC{radar.EmailSecurityTimeseriesGroupMaliciousParamsDMARCPass, radar.EmailSecurityTimeseriesGroupMaliciousParamsDMARCNone, radar.EmailSecurityTimeseriesGroupMaliciousParamsDMARCFail}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupMaliciousParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsSPF{radar.EmailSecurityTimeseriesGroupMaliciousParamsSPFPass, radar.EmailSecurityTimeseriesGroupMaliciousParamsSPFNone, radar.EmailSecurityTimeseriesGroupMaliciousParamsSPFFail}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion{radar.EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupSpamWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.Spam(context.TODO(), radar.EmailSecurityTimeseriesGroupSpamParams{
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupSpamParamsAggInterval1h),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpamParamsARC{radar.EmailSecurityTimeseriesGroupSpamParamsARCPass, radar.EmailSecurityTimeseriesGroupSpamParamsARCNone, radar.EmailSecurityTimeseriesGroupSpamParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpamParamsDateRange{radar.EmailSecurityTimeseriesGroupSpamParamsDateRange1d, radar.EmailSecurityTimeseriesGroupSpamParamsDateRange2d, radar.EmailSecurityTimeseriesGroupSpamParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpamParamsDKIM{radar.EmailSecurityTimeseriesGroupSpamParamsDKIMPass, radar.EmailSecurityTimeseriesGroupSpamParamsDKIMNone, radar.EmailSecurityTimeseriesGroupSpamParamsDKIMFail}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpamParamsDMARC{radar.EmailSecurityTimeseriesGroupSpamParamsDMARCPass, radar.EmailSecurityTimeseriesGroupSpamParamsDMARCNone, radar.EmailSecurityTimeseriesGroupSpamParamsDMARCFail}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupSpamParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpamParamsSPF{radar.EmailSecurityTimeseriesGroupSpamParamsSPFPass, radar.EmailSecurityTimeseriesGroupSpamParamsSPFNone, radar.EmailSecurityTimeseriesGroupSpamParamsSPFFail}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpamParamsTLSVersion{radar.EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupSPFWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.SPF(context.TODO(), radar.EmailSecurityTimeseriesGroupSPFParams{
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupSPFParamsAggInterval1h),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupSPFParamsARC{radar.EmailSecurityTimeseriesGroupSPFParamsARCPass, radar.EmailSecurityTimeseriesGroupSPFParamsARCNone, radar.EmailSecurityTimeseriesGroupSPFParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailSecurityTimeseriesGroupSPFParamsDateRange{radar.EmailSecurityTimeseriesGroupSPFParamsDateRange1d, radar.EmailSecurityTimeseriesGroupSPFParamsDateRange2d, radar.EmailSecurityTimeseriesGroupSPFParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupSPFParamsDKIM{radar.EmailSecurityTimeseriesGroupSPFParamsDKIMPass, radar.EmailSecurityTimeseriesGroupSPFParamsDKIMNone, radar.EmailSecurityTimeseriesGroupSPFParamsDKIMFail}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupSPFParamsDMARC{radar.EmailSecurityTimeseriesGroupSPFParamsDMARCPass, radar.EmailSecurityTimeseriesGroupSPFParamsDMARCNone, radar.EmailSecurityTimeseriesGroupSPFParamsDMARCFail}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupSPFParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupSPFParamsTLSVersion{radar.EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupSpoofWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.Spoof(context.TODO(), radar.EmailSecurityTimeseriesGroupSpoofParams{
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupSpoofParamsAggInterval1h),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsARC{radar.EmailSecurityTimeseriesGroupSpoofParamsARCPass, radar.EmailSecurityTimeseriesGroupSpoofParamsARCNone, radar.EmailSecurityTimeseriesGroupSpoofParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsDateRange{radar.EmailSecurityTimeseriesGroupSpoofParamsDateRange1d, radar.EmailSecurityTimeseriesGroupSpoofParamsDateRange2d, radar.EmailSecurityTimeseriesGroupSpoofParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsDKIM{radar.EmailSecurityTimeseriesGroupSpoofParamsDKIMPass, radar.EmailSecurityTimeseriesGroupSpoofParamsDKIMNone, radar.EmailSecurityTimeseriesGroupSpoofParamsDKIMFail}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsDMARC{radar.EmailSecurityTimeseriesGroupSpoofParamsDMARCPass, radar.EmailSecurityTimeseriesGroupSpoofParamsDMARCNone, radar.EmailSecurityTimeseriesGroupSpoofParamsDMARCFail}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupSpoofParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsSPF{radar.EmailSecurityTimeseriesGroupSpoofParamsSPFPass, radar.EmailSecurityTimeseriesGroupSpoofParamsSPFNone, radar.EmailSecurityTimeseriesGroupSpoofParamsSPFFail}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsTLSVersion{radar.EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupThreatCategoryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.ThreatCategory(context.TODO(), radar.EmailSecurityTimeseriesGroupThreatCategoryParams{
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1h),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsARC{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsARCPass, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsARCNone, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange1d, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange2d, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMPass, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMNone, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMFail}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCPass, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCNone, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCFail}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupThreatCategoryParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsSPF{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsSPFPass, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsSPFNone, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsSPFFail}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_0, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_1, radar.EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecurityTimeseriesGroupTLSVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.TLSVersion(context.TODO(), radar.EmailSecurityTimeseriesGroupTLSVersionParams{
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1h),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupTLSVersionParamsARC{radar.EmailSecurityTimeseriesGroupTLSVersionParamsARCPass, radar.EmailSecurityTimeseriesGroupTLSVersionParamsARCNone, radar.EmailSecurityTimeseriesGroupTLSVersionParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailSecurityTimeseriesGroupTLSVersionParamsDateRange{radar.EmailSecurityTimeseriesGroupTLSVersionParamsDateRange1d, radar.EmailSecurityTimeseriesGroupTLSVersionParamsDateRange2d, radar.EmailSecurityTimeseriesGroupTLSVersionParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupTLSVersionParamsDKIM{radar.EmailSecurityTimeseriesGroupTLSVersionParamsDKIMPass, radar.EmailSecurityTimeseriesGroupTLSVersionParamsDKIMNone, radar.EmailSecurityTimeseriesGroupTLSVersionParamsDKIMFail}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupTLSVersionParamsDMARC{radar.EmailSecurityTimeseriesGroupTLSVersionParamsDMARCPass, radar.EmailSecurityTimeseriesGroupTLSVersionParamsDMARCNone, radar.EmailSecurityTimeseriesGroupTLSVersionParamsDMARCFail}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupTLSVersionParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupTLSVersionParamsSPF{radar.EmailSecurityTimeseriesGroupTLSVersionParamsSPFPass, radar.EmailSecurityTimeseriesGroupTLSVersionParamsSPFNone, radar.EmailSecurityTimeseriesGroupTLSVersionParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
