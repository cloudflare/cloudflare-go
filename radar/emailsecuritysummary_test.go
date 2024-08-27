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
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummaryARCParamsDKIM{radar.EmailSecuritySummaryARCParamsDKIMPass, radar.EmailSecuritySummaryARCParamsDKIMNone, radar.EmailSecuritySummaryARCParamsDKIMFail}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummaryARCParamsDMARC{radar.EmailSecuritySummaryARCParamsDMARCPass, radar.EmailSecuritySummaryARCParamsDMARCNone, radar.EmailSecuritySummaryARCParamsDMARCFail}),
		Format:     cloudflare.F(radar.EmailSecuritySummaryARCParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummaryARCParamsSPF{radar.EmailSecuritySummaryARCParamsSPFPass, radar.EmailSecuritySummaryARCParamsSPFNone, radar.EmailSecuritySummaryARCParamsSPFFail}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummaryARCParamsTLSVersion{radar.EmailSecuritySummaryARCParamsTLSVersionTlSv1_0, radar.EmailSecuritySummaryARCParamsTLSVersionTlSv1_1, radar.EmailSecuritySummaryARCParamsTLSVersionTlSv1_2}),
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
		ARC:        cloudflare.F([]radar.EmailSecuritySummaryDKIMParamsARC{radar.EmailSecuritySummaryDKIMParamsARCPass, radar.EmailSecuritySummaryDKIMParamsARCNone, radar.EmailSecuritySummaryDKIMParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummaryDKIMParamsDMARC{radar.EmailSecuritySummaryDKIMParamsDMARCPass, radar.EmailSecuritySummaryDKIMParamsDMARCNone, radar.EmailSecuritySummaryDKIMParamsDMARCFail}),
		Format:     cloudflare.F(radar.EmailSecuritySummaryDKIMParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummaryDKIMParamsSPF{radar.EmailSecuritySummaryDKIMParamsSPFPass, radar.EmailSecuritySummaryDKIMParamsSPFNone, radar.EmailSecuritySummaryDKIMParamsSPFFail}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummaryDKIMParamsTLSVersion{radar.EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_0, radar.EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_1, radar.EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_2}),
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
		ARC:        cloudflare.F([]radar.EmailSecuritySummaryDMARCParamsARC{radar.EmailSecuritySummaryDMARCParamsARCPass, radar.EmailSecuritySummaryDMARCParamsARCNone, radar.EmailSecuritySummaryDMARCParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummaryDMARCParamsDKIM{radar.EmailSecuritySummaryDMARCParamsDKIMPass, radar.EmailSecuritySummaryDMARCParamsDKIMNone, radar.EmailSecuritySummaryDMARCParamsDKIMFail}),
		Format:     cloudflare.F(radar.EmailSecuritySummaryDMARCParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummaryDMARCParamsSPF{radar.EmailSecuritySummaryDMARCParamsSPFPass, radar.EmailSecuritySummaryDMARCParamsSPFNone, radar.EmailSecuritySummaryDMARCParamsSPFFail}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummaryDMARCParamsTLSVersion{radar.EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_0, radar.EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_1, radar.EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_2}),
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
		ARC:        cloudflare.F([]radar.EmailSecuritySummaryMaliciousParamsARC{radar.EmailSecuritySummaryMaliciousParamsARCPass, radar.EmailSecuritySummaryMaliciousParamsARCNone, radar.EmailSecuritySummaryMaliciousParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummaryMaliciousParamsDKIM{radar.EmailSecuritySummaryMaliciousParamsDKIMPass, radar.EmailSecuritySummaryMaliciousParamsDKIMNone, radar.EmailSecuritySummaryMaliciousParamsDKIMFail}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummaryMaliciousParamsDMARC{radar.EmailSecuritySummaryMaliciousParamsDMARCPass, radar.EmailSecuritySummaryMaliciousParamsDMARCNone, radar.EmailSecuritySummaryMaliciousParamsDMARCFail}),
		Format:     cloudflare.F(radar.EmailSecuritySummaryMaliciousParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummaryMaliciousParamsSPF{radar.EmailSecuritySummaryMaliciousParamsSPFPass, radar.EmailSecuritySummaryMaliciousParamsSPFNone, radar.EmailSecuritySummaryMaliciousParamsSPFFail}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummaryMaliciousParamsTLSVersion{radar.EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_0, radar.EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_1, radar.EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_2}),
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
		ARC:        cloudflare.F([]radar.EmailSecuritySummarySpamParamsARC{radar.EmailSecuritySummarySpamParamsARCPass, radar.EmailSecuritySummarySpamParamsARCNone, radar.EmailSecuritySummarySpamParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummarySpamParamsDKIM{radar.EmailSecuritySummarySpamParamsDKIMPass, radar.EmailSecuritySummarySpamParamsDKIMNone, radar.EmailSecuritySummarySpamParamsDKIMFail}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummarySpamParamsDMARC{radar.EmailSecuritySummarySpamParamsDMARCPass, radar.EmailSecuritySummarySpamParamsDMARCNone, radar.EmailSecuritySummarySpamParamsDMARCFail}),
		Format:     cloudflare.F(radar.EmailSecuritySummarySpamParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummarySpamParamsSPF{radar.EmailSecuritySummarySpamParamsSPFPass, radar.EmailSecuritySummarySpamParamsSPFNone, radar.EmailSecuritySummarySpamParamsSPFFail}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummarySpamParamsTLSVersion{radar.EmailSecuritySummarySpamParamsTLSVersionTlSv1_0, radar.EmailSecuritySummarySpamParamsTLSVersionTlSv1_1, radar.EmailSecuritySummarySpamParamsTLSVersionTlSv1_2}),
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
		ARC:        cloudflare.F([]radar.EmailSecuritySummarySPFParamsARC{radar.EmailSecuritySummarySPFParamsARCPass, radar.EmailSecuritySummarySPFParamsARCNone, radar.EmailSecuritySummarySPFParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummarySPFParamsDKIM{radar.EmailSecuritySummarySPFParamsDKIMPass, radar.EmailSecuritySummarySPFParamsDKIMNone, radar.EmailSecuritySummarySPFParamsDKIMFail}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummarySPFParamsDMARC{radar.EmailSecuritySummarySPFParamsDMARCPass, radar.EmailSecuritySummarySPFParamsDMARCNone, radar.EmailSecuritySummarySPFParamsDMARCFail}),
		Format:     cloudflare.F(radar.EmailSecuritySummarySPFParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummarySPFParamsTLSVersion{radar.EmailSecuritySummarySPFParamsTLSVersionTlSv1_0, radar.EmailSecuritySummarySPFParamsTLSVersionTlSv1_1, radar.EmailSecuritySummarySPFParamsTLSVersionTlSv1_2}),
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
		ARC:        cloudflare.F([]radar.EmailSecuritySummarySpoofParamsARC{radar.EmailSecuritySummarySpoofParamsARCPass, radar.EmailSecuritySummarySpoofParamsARCNone, radar.EmailSecuritySummarySpoofParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummarySpoofParamsDKIM{radar.EmailSecuritySummarySpoofParamsDKIMPass, radar.EmailSecuritySummarySpoofParamsDKIMNone, radar.EmailSecuritySummarySpoofParamsDKIMFail}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummarySpoofParamsDMARC{radar.EmailSecuritySummarySpoofParamsDMARCPass, radar.EmailSecuritySummarySpoofParamsDMARCNone, radar.EmailSecuritySummarySpoofParamsDMARCFail}),
		Format:     cloudflare.F(radar.EmailSecuritySummarySpoofParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummarySpoofParamsSPF{radar.EmailSecuritySummarySpoofParamsSPFPass, radar.EmailSecuritySummarySpoofParamsSPFNone, radar.EmailSecuritySummarySpoofParamsSPFFail}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummarySpoofParamsTLSVersion{radar.EmailSecuritySummarySpoofParamsTLSVersionTlSv1_0, radar.EmailSecuritySummarySpoofParamsTLSVersionTlSv1_1, radar.EmailSecuritySummarySpoofParamsTLSVersionTlSv1_2}),
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
		ARC:        cloudflare.F([]radar.EmailSecuritySummaryThreatCategoryParamsARC{radar.EmailSecuritySummaryThreatCategoryParamsARCPass, radar.EmailSecuritySummaryThreatCategoryParamsARCNone, radar.EmailSecuritySummaryThreatCategoryParamsARCFail}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:       cloudflare.F([]radar.EmailSecuritySummaryThreatCategoryParamsDKIM{radar.EmailSecuritySummaryThreatCategoryParamsDKIMPass, radar.EmailSecuritySummaryThreatCategoryParamsDKIMNone, radar.EmailSecuritySummaryThreatCategoryParamsDKIMFail}),
		DMARC:      cloudflare.F([]radar.EmailSecuritySummaryThreatCategoryParamsDMARC{radar.EmailSecuritySummaryThreatCategoryParamsDMARCPass, radar.EmailSecuritySummaryThreatCategoryParamsDMARCNone, radar.EmailSecuritySummaryThreatCategoryParamsDMARCFail}),
		Format:     cloudflare.F(radar.EmailSecuritySummaryThreatCategoryParamsFormatJson),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		SPF:        cloudflare.F([]radar.EmailSecuritySummaryThreatCategoryParamsSPF{radar.EmailSecuritySummaryThreatCategoryParamsSPFPass, radar.EmailSecuritySummaryThreatCategoryParamsSPFNone, radar.EmailSecuritySummaryThreatCategoryParamsSPFFail}),
		TLSVersion: cloudflare.F([]radar.EmailSecuritySummaryThreatCategoryParamsTLSVersion{radar.EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_0, radar.EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_1, radar.EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_2}),
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
		ARC:       cloudflare.F([]radar.EmailSecuritySummaryTLSVersionParamsARC{radar.EmailSecuritySummaryTLSVersionParamsARCPass, radar.EmailSecuritySummaryTLSVersionParamsARCNone, radar.EmailSecuritySummaryTLSVersionParamsARCFail}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]radar.EmailSecuritySummaryTLSVersionParamsDKIM{radar.EmailSecuritySummaryTLSVersionParamsDKIMPass, radar.EmailSecuritySummaryTLSVersionParamsDKIMNone, radar.EmailSecuritySummaryTLSVersionParamsDKIMFail}),
		DMARC:     cloudflare.F([]radar.EmailSecuritySummaryTLSVersionParamsDMARC{radar.EmailSecuritySummaryTLSVersionParamsDMARCPass, radar.EmailSecuritySummaryTLSVersionParamsDMARCNone, radar.EmailSecuritySummaryTLSVersionParamsDMARCFail}),
		Format:    cloudflare.F(radar.EmailSecuritySummaryTLSVersionParamsFormatJson),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]radar.EmailSecuritySummaryTLSVersionParamsSPF{radar.EmailSecuritySummaryTLSVersionParamsSPFPass, radar.EmailSecuritySummaryTLSVersionParamsSPFNone, radar.EmailSecuritySummaryTLSVersionParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
