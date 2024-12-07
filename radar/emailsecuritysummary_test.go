// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/radar"
)

func TestEmailSecuritySummaryARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.ARC(context.TODO(), radar.EmailSecuritySummaryARCParams{
		DateEnd:    cloudflare.F([]time.Time{time.Now()}),
		DateRange:  cloudflare.F([]string{"7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummaryARCParamsDKIM{radar.EmailSecuritySummaryARCParamsDKIMPass}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummaryARCParamsDMARC{radar.EmailSecuritySummaryARCParamsDMARCPass}),
		Format:     cloudflare.F(radar.EmailSecuritySummaryARCParamsFormatJson),
		Name:       cloudflare.F([]string{"string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummaryARCParamsSPF{radar.EmailSecuritySummaryARCParamsSPFPass}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummaryARCParamsTLSVersion{radar.EmailSecuritySummaryARCParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummaryDKIMWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.DKIM(context.TODO(), radar.EmailSecuritySummaryDKIMParams{
		ARC:        cloudflare.F([]radar.EmailSecuritySummaryDKIMParamsARC{radar.EmailSecuritySummaryDKIMParamsARCPass}),
		DateEnd:    cloudflare.F([]time.Time{time.Now()}),
		DateRange:  cloudflare.F([]string{"7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now()}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummaryDKIMParamsDMARC{radar.EmailSecuritySummaryDKIMParamsDMARCPass}),
		Format:     cloudflare.F(radar.EmailSecuritySummaryDKIMParamsFormatJson),
		Name:       cloudflare.F([]string{"string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummaryDKIMParamsSPF{radar.EmailSecuritySummaryDKIMParamsSPFPass}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummaryDKIMParamsTLSVersion{radar.EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummaryDMARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.DMARC(context.TODO(), radar.EmailSecuritySummaryDMARCParams{
		ARC:        cloudflare.F([]radar.EmailSecuritySummaryDMARCParamsARC{radar.EmailSecuritySummaryDMARCParamsARCPass}),
		DateEnd:    cloudflare.F([]time.Time{time.Now()}),
		DateRange:  cloudflare.F([]string{"7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummaryDMARCParamsDKIM{radar.EmailSecuritySummaryDMARCParamsDKIMPass}),
		Format:     cloudflare.F(radar.EmailSecuritySummaryDMARCParamsFormatJson),
		Name:       cloudflare.F([]string{"string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummaryDMARCParamsSPF{radar.EmailSecuritySummaryDMARCParamsSPFPass}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummaryDMARCParamsTLSVersion{radar.EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummaryMaliciousWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.Malicious(context.TODO(), radar.EmailSecuritySummaryMaliciousParams{
		ARC:        cloudflare.F([]radar.EmailSecuritySummaryMaliciousParamsARC{radar.EmailSecuritySummaryMaliciousParamsARCPass}),
		DateEnd:    cloudflare.F([]time.Time{time.Now()}),
		DateRange:  cloudflare.F([]string{"7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummaryMaliciousParamsDKIM{radar.EmailSecuritySummaryMaliciousParamsDKIMPass}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummaryMaliciousParamsDMARC{radar.EmailSecuritySummaryMaliciousParamsDMARCPass}),
		Format:     cloudflare.F(radar.EmailSecuritySummaryMaliciousParamsFormatJson),
		Name:       cloudflare.F([]string{"string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummaryMaliciousParamsSPF{radar.EmailSecuritySummaryMaliciousParamsSPFPass}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummaryMaliciousParamsTLSVersion{radar.EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummarySpamWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.Spam(context.TODO(), radar.EmailSecuritySummarySpamParams{
		ARC:        cloudflare.F([]radar.EmailSecuritySummarySpamParamsARC{radar.EmailSecuritySummarySpamParamsARCPass}),
		DateEnd:    cloudflare.F([]time.Time{time.Now()}),
		DateRange:  cloudflare.F([]string{"7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummarySpamParamsDKIM{radar.EmailSecuritySummarySpamParamsDKIMPass}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummarySpamParamsDMARC{radar.EmailSecuritySummarySpamParamsDMARCPass}),
		Format:     cloudflare.F(radar.EmailSecuritySummarySpamParamsFormatJson),
		Name:       cloudflare.F([]string{"string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummarySpamParamsSPF{radar.EmailSecuritySummarySpamParamsSPFPass}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummarySpamParamsTLSVersion{radar.EmailSecuritySummarySpamParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummarySPFWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.SPF(context.TODO(), radar.EmailSecuritySummarySPFParams{
		ARC:        cloudflare.F([]radar.EmailSecuritySummarySPFParamsARC{radar.EmailSecuritySummarySPFParamsARCPass}),
		DateEnd:    cloudflare.F([]time.Time{time.Now()}),
		DateRange:  cloudflare.F([]string{"7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummarySPFParamsDKIM{radar.EmailSecuritySummarySPFParamsDKIMPass}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummarySPFParamsDMARC{radar.EmailSecuritySummarySPFParamsDMARCPass}),
		Format:     cloudflare.F(radar.EmailSecuritySummarySPFParamsFormatJson),
		Name:       cloudflare.F([]string{"string"}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummarySPFParamsTLSVersion{radar.EmailSecuritySummarySPFParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummarySpoofWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.Spoof(context.TODO(), radar.EmailSecuritySummarySpoofParams{
		ARC:        cloudflare.F([]radar.EmailSecuritySummarySpoofParamsARC{radar.EmailSecuritySummarySpoofParamsARCPass}),
		DateEnd:    cloudflare.F([]time.Time{time.Now()}),
		DateRange:  cloudflare.F([]string{"7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummarySpoofParamsDKIM{radar.EmailSecuritySummarySpoofParamsDKIMPass}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummarySpoofParamsDMARC{radar.EmailSecuritySummarySpoofParamsDMARCPass}),
		Format:     cloudflare.F(radar.EmailSecuritySummarySpoofParamsFormatJson),
		Name:       cloudflare.F([]string{"string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummarySpoofParamsSPF{radar.EmailSecuritySummarySpoofParamsSPFPass}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummarySpoofParamsTLSVersion{radar.EmailSecuritySummarySpoofParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummaryThreatCategoryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.ThreatCategory(context.TODO(), radar.EmailSecuritySummaryThreatCategoryParams{
		ARC:        cloudflare.F([]radar.EmailSecuritySummaryThreatCategoryParamsARC{radar.EmailSecuritySummaryThreatCategoryParamsARCPass}),
		DateEnd:    cloudflare.F([]time.Time{time.Now()}),
		DateRange:  cloudflare.F([]string{"7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummaryThreatCategoryParamsDKIM{radar.EmailSecuritySummaryThreatCategoryParamsDKIMPass}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummaryThreatCategoryParamsDMARC{radar.EmailSecuritySummaryThreatCategoryParamsDMARCPass}),
		Format:     cloudflare.F(radar.EmailSecuritySummaryThreatCategoryParamsFormatJson),
		Name:       cloudflare.F([]string{"string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummaryThreatCategoryParamsSPF{radar.EmailSecuritySummaryThreatCategoryParamsSPFPass}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummaryThreatCategoryParamsTLSVersion{radar.EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSecuritySummaryTLSVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.TLSVersion(context.TODO(), radar.EmailSecuritySummaryTLSVersionParams{
		ARC:       cloudflare.F([]radar.EmailSecuritySummaryTLSVersionParamsARC{radar.EmailSecuritySummaryTLSVersionParamsARCPass}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		DKIM:      cloudflare.F([]radar.EmailSecuritySummaryTLSVersionParamsDKIM{radar.EmailSecuritySummaryTLSVersionParamsDKIMPass}),
		DMARC:     cloudflare.F([]radar.EmailSecuritySummaryTLSVersionParamsDMARC{radar.EmailSecuritySummaryTLSVersionParamsDMARCPass}),
		Format:    cloudflare.F(radar.EmailSecuritySummaryTLSVersionParamsFormatJson),
		Name:      cloudflare.F([]string{"string"}),
		SPF:       cloudflare.F([]radar.EmailSecuritySummaryTLSVersionParamsSPF{radar.EmailSecuritySummaryTLSVersionParamsSPFPass}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
