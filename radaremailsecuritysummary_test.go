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

func TestRadarEmailSecuritySummaryARCWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Email.Security.Summary.ARC(context.TODO(), cloudflare.RadarEmailSecuritySummaryARCParams{
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]cloudflare.RadarEmailSecuritySummaryARCParamsDateRange{cloudflare.RadarEmailSecuritySummaryARCParamsDateRange1d, cloudflare.RadarEmailSecuritySummaryARCParamsDateRange2d, cloudflare.RadarEmailSecuritySummaryARCParamsDateRange7d}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryARCParamsDKIM{cloudflare.RadarEmailSecuritySummaryARCParamsDKIMPass, cloudflare.RadarEmailSecuritySummaryARCParamsDKIMNone, cloudflare.RadarEmailSecuritySummaryARCParamsDKIMFail}),
		DMARC:      cloudflare.F([]cloudflare.RadarEmailSecuritySummaryARCParamsDMARC{cloudflare.RadarEmailSecuritySummaryARCParamsDMARCPass, cloudflare.RadarEmailSecuritySummaryARCParamsDMARCNone, cloudflare.RadarEmailSecuritySummaryARCParamsDMARCFail}),
		Format:     cloudflare.F(cloudflare.RadarEmailSecuritySummaryARCParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]cloudflare.RadarEmailSecuritySummaryARCParamsSPF{cloudflare.RadarEmailSecuritySummaryARCParamsSPFPass, cloudflare.RadarEmailSecuritySummaryARCParamsSPFNone, cloudflare.RadarEmailSecuritySummaryARCParamsSPFFail}),
		TLSVersion: cloudflare.F([]cloudflare.RadarEmailSecuritySummaryARCParamsTLSVersion{cloudflare.RadarEmailSecuritySummaryARCParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecuritySummaryARCParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecuritySummaryARCParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryDKIMWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Email.Security.Summary.DKIM(context.TODO(), cloudflare.RadarEmailSecuritySummaryDKIMParams{
		ARC:        cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDKIMParamsARC{cloudflare.RadarEmailSecuritySummaryDKIMParamsARCPass, cloudflare.RadarEmailSecuritySummaryDKIMParamsARCNone, cloudflare.RadarEmailSecuritySummaryDKIMParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDKIMParamsDateRange{cloudflare.RadarEmailSecuritySummaryDKIMParamsDateRange1d, cloudflare.RadarEmailSecuritySummaryDKIMParamsDateRange2d, cloudflare.RadarEmailSecuritySummaryDKIMParamsDateRange7d}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DMARC:      cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDKIMParamsDMARC{cloudflare.RadarEmailSecuritySummaryDKIMParamsDMARCPass, cloudflare.RadarEmailSecuritySummaryDKIMParamsDMARCNone, cloudflare.RadarEmailSecuritySummaryDKIMParamsDMARCFail}),
		Format:     cloudflare.F(cloudflare.RadarEmailSecuritySummaryDKIMParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDKIMParamsSPF{cloudflare.RadarEmailSecuritySummaryDKIMParamsSPFPass, cloudflare.RadarEmailSecuritySummaryDKIMParamsSPFNone, cloudflare.RadarEmailSecuritySummaryDKIMParamsSPFFail}),
		TLSVersion: cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDKIMParamsTLSVersion{cloudflare.RadarEmailSecuritySummaryDKIMParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecuritySummaryDKIMParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecuritySummaryDKIMParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryDMARCWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Email.Security.Summary.DMARC(context.TODO(), cloudflare.RadarEmailSecuritySummaryDMARCParams{
		ARC:        cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDMARCParamsARC{cloudflare.RadarEmailSecuritySummaryDMARCParamsARCPass, cloudflare.RadarEmailSecuritySummaryDMARCParamsARCNone, cloudflare.RadarEmailSecuritySummaryDMARCParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDMARCParamsDateRange{cloudflare.RadarEmailSecuritySummaryDMARCParamsDateRange1d, cloudflare.RadarEmailSecuritySummaryDMARCParamsDateRange2d, cloudflare.RadarEmailSecuritySummaryDMARCParamsDateRange7d}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDMARCParamsDKIM{cloudflare.RadarEmailSecuritySummaryDMARCParamsDKIMPass, cloudflare.RadarEmailSecuritySummaryDMARCParamsDKIMNone, cloudflare.RadarEmailSecuritySummaryDMARCParamsDKIMFail}),
		Format:     cloudflare.F(cloudflare.RadarEmailSecuritySummaryDMARCParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDMARCParamsSPF{cloudflare.RadarEmailSecuritySummaryDMARCParamsSPFPass, cloudflare.RadarEmailSecuritySummaryDMARCParamsSPFNone, cloudflare.RadarEmailSecuritySummaryDMARCParamsSPFFail}),
		TLSVersion: cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDMARCParamsTLSVersion{cloudflare.RadarEmailSecuritySummaryDMARCParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecuritySummaryDMARCParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecuritySummaryDMARCParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryMaliciousWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Email.Security.Summary.Malicious(context.TODO(), cloudflare.RadarEmailSecuritySummaryMaliciousParams{
		ARC:        cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousParamsARC{cloudflare.RadarEmailSecuritySummaryMaliciousParamsARCPass, cloudflare.RadarEmailSecuritySummaryMaliciousParamsARCNone, cloudflare.RadarEmailSecuritySummaryMaliciousParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousParamsDateRange{cloudflare.RadarEmailSecuritySummaryMaliciousParamsDateRange1d, cloudflare.RadarEmailSecuritySummaryMaliciousParamsDateRange2d, cloudflare.RadarEmailSecuritySummaryMaliciousParamsDateRange7d}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousParamsDKIM{cloudflare.RadarEmailSecuritySummaryMaliciousParamsDKIMPass, cloudflare.RadarEmailSecuritySummaryMaliciousParamsDKIMNone, cloudflare.RadarEmailSecuritySummaryMaliciousParamsDKIMFail}),
		DMARC:      cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousParamsDMARC{cloudflare.RadarEmailSecuritySummaryMaliciousParamsDMARCPass, cloudflare.RadarEmailSecuritySummaryMaliciousParamsDMARCNone, cloudflare.RadarEmailSecuritySummaryMaliciousParamsDMARCFail}),
		Format:     cloudflare.F(cloudflare.RadarEmailSecuritySummaryMaliciousParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousParamsSPF{cloudflare.RadarEmailSecuritySummaryMaliciousParamsSPFPass, cloudflare.RadarEmailSecuritySummaryMaliciousParamsSPFNone, cloudflare.RadarEmailSecuritySummaryMaliciousParamsSPFFail}),
		TLSVersion: cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousParamsTLSVersion{cloudflare.RadarEmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummarySpamWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Email.Security.Summary.Spam(context.TODO(), cloudflare.RadarEmailSecuritySummarySpamParams{
		ARC:        cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamParamsARC{cloudflare.RadarEmailSecuritySummarySpamParamsARCPass, cloudflare.RadarEmailSecuritySummarySpamParamsARCNone, cloudflare.RadarEmailSecuritySummarySpamParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamParamsDateRange{cloudflare.RadarEmailSecuritySummarySpamParamsDateRange1d, cloudflare.RadarEmailSecuritySummarySpamParamsDateRange2d, cloudflare.RadarEmailSecuritySummarySpamParamsDateRange7d}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamParamsDKIM{cloudflare.RadarEmailSecuritySummarySpamParamsDKIMPass, cloudflare.RadarEmailSecuritySummarySpamParamsDKIMNone, cloudflare.RadarEmailSecuritySummarySpamParamsDKIMFail}),
		DMARC:      cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamParamsDMARC{cloudflare.RadarEmailSecuritySummarySpamParamsDMARCPass, cloudflare.RadarEmailSecuritySummarySpamParamsDMARCNone, cloudflare.RadarEmailSecuritySummarySpamParamsDMARCFail}),
		Format:     cloudflare.F(cloudflare.RadarEmailSecuritySummarySpamParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamParamsSPF{cloudflare.RadarEmailSecuritySummarySpamParamsSPFPass, cloudflare.RadarEmailSecuritySummarySpamParamsSPFNone, cloudflare.RadarEmailSecuritySummarySpamParamsSPFFail}),
		TLSVersion: cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamParamsTLSVersion{cloudflare.RadarEmailSecuritySummarySpamParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecuritySummarySpamParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecuritySummarySpamParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummarySPFWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Email.Security.Summary.SPF(context.TODO(), cloudflare.RadarEmailSecuritySummarySPFParams{
		ARC:        cloudflare.F([]cloudflare.RadarEmailSecuritySummarySPFParamsARC{cloudflare.RadarEmailSecuritySummarySPFParamsARCPass, cloudflare.RadarEmailSecuritySummarySPFParamsARCNone, cloudflare.RadarEmailSecuritySummarySPFParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]cloudflare.RadarEmailSecuritySummarySPFParamsDateRange{cloudflare.RadarEmailSecuritySummarySPFParamsDateRange1d, cloudflare.RadarEmailSecuritySummarySPFParamsDateRange2d, cloudflare.RadarEmailSecuritySummarySPFParamsDateRange7d}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]cloudflare.RadarEmailSecuritySummarySPFParamsDKIM{cloudflare.RadarEmailSecuritySummarySPFParamsDKIMPass, cloudflare.RadarEmailSecuritySummarySPFParamsDKIMNone, cloudflare.RadarEmailSecuritySummarySPFParamsDKIMFail}),
		DMARC:      cloudflare.F([]cloudflare.RadarEmailSecuritySummarySPFParamsDMARC{cloudflare.RadarEmailSecuritySummarySPFParamsDMARCPass, cloudflare.RadarEmailSecuritySummarySPFParamsDMARCNone, cloudflare.RadarEmailSecuritySummarySPFParamsDMARCFail}),
		Format:     cloudflare.F(cloudflare.RadarEmailSecuritySummarySPFParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		TLSVersion: cloudflare.F([]cloudflare.RadarEmailSecuritySummarySPFParamsTLSVersion{cloudflare.RadarEmailSecuritySummarySPFParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecuritySummarySPFParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecuritySummarySPFParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummarySpoofWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Email.Security.Summary.Spoof(context.TODO(), cloudflare.RadarEmailSecuritySummarySpoofParams{
		ARC:        cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpoofParamsARC{cloudflare.RadarEmailSecuritySummarySpoofParamsARCPass, cloudflare.RadarEmailSecuritySummarySpoofParamsARCNone, cloudflare.RadarEmailSecuritySummarySpoofParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpoofParamsDateRange{cloudflare.RadarEmailSecuritySummarySpoofParamsDateRange1d, cloudflare.RadarEmailSecuritySummarySpoofParamsDateRange2d, cloudflare.RadarEmailSecuritySummarySpoofParamsDateRange7d}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpoofParamsDKIM{cloudflare.RadarEmailSecuritySummarySpoofParamsDKIMPass, cloudflare.RadarEmailSecuritySummarySpoofParamsDKIMNone, cloudflare.RadarEmailSecuritySummarySpoofParamsDKIMFail}),
		DMARC:      cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpoofParamsDMARC{cloudflare.RadarEmailSecuritySummarySpoofParamsDMARCPass, cloudflare.RadarEmailSecuritySummarySpoofParamsDMARCNone, cloudflare.RadarEmailSecuritySummarySpoofParamsDMARCFail}),
		Format:     cloudflare.F(cloudflare.RadarEmailSecuritySummarySpoofParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpoofParamsSPF{cloudflare.RadarEmailSecuritySummarySpoofParamsSPFPass, cloudflare.RadarEmailSecuritySummarySpoofParamsSPFNone, cloudflare.RadarEmailSecuritySummarySpoofParamsSPFFail}),
		TLSVersion: cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpoofParamsTLSVersion{cloudflare.RadarEmailSecuritySummarySpoofParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecuritySummarySpoofParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecuritySummarySpoofParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryThreatCategoryWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Email.Security.Summary.ThreatCategory(context.TODO(), cloudflare.RadarEmailSecuritySummaryThreatCategoryParams{
		ARC:        cloudflare.F([]cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsARC{cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsARCPass, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsARCNone, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDateRange{cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDateRange1d, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDateRange2d, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDateRange7d}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDKIM{cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDKIMPass, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDKIMNone, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDKIMFail}),
		DMARC:      cloudflare.F([]cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDMARC{cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDMARCPass, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDMARCNone, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDMARCFail}),
		Format:     cloudflare.F(cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsSPF{cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsSPFPass, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsSPFNone, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsSPFFail}),
		TLSVersion: cloudflare.F([]cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsTLSVersion{cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryTLSVersionWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Radar.Email.Security.Summary.TLSVersion(context.TODO(), cloudflare.RadarEmailSecuritySummaryTLSVersionParams{
		ARC:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryTLSVersionParamsARC{cloudflare.RadarEmailSecuritySummaryTLSVersionParamsARCPass, cloudflare.RadarEmailSecuritySummaryTLSVersionParamsARCNone, cloudflare.RadarEmailSecuritySummaryTLSVersionParamsARCFail}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailSecuritySummaryTLSVersionParamsDateRange{cloudflare.RadarEmailSecuritySummaryTLSVersionParamsDateRange1d, cloudflare.RadarEmailSecuritySummaryTLSVersionParamsDateRange2d, cloudflare.RadarEmailSecuritySummaryTLSVersionParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailSecuritySummaryTLSVersionParamsDKIM{cloudflare.RadarEmailSecuritySummaryTLSVersionParamsDKIMPass, cloudflare.RadarEmailSecuritySummaryTLSVersionParamsDKIMNone, cloudflare.RadarEmailSecuritySummaryTLSVersionParamsDKIMFail}),
		DMARC:     cloudflare.F([]cloudflare.RadarEmailSecuritySummaryTLSVersionParamsDMARC{cloudflare.RadarEmailSecuritySummaryTLSVersionParamsDMARCPass, cloudflare.RadarEmailSecuritySummaryTLSVersionParamsDMARCNone, cloudflare.RadarEmailSecuritySummaryTLSVersionParamsDMARCFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailSecuritySummaryTLSVersionParamsFormatJson),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryTLSVersionParamsSPF{cloudflare.RadarEmailSecuritySummaryTLSVersionParamsSPFPass, cloudflare.RadarEmailSecuritySummaryTLSVersionParamsSPFNone, cloudflare.RadarEmailSecuritySummaryTLSVersionParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
