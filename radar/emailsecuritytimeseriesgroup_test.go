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

func TestEmailSecurityTimeseriesGroupARCWithOptionalParams(t *testing.T) {
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
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupARCParamsAggInterval15m),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupARCParamsDKIM{radar.EmailSecurityTimeseriesGroupARCParamsDKIMPass}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupARCParamsDMARC{radar.EmailSecurityTimeseriesGroupARCParamsDMARCPass}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupARCParamsFormatJson),
		Name:        cloudflare.F([]string{"string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupARCParamsSPF{radar.EmailSecurityTimeseriesGroupARCParamsSPFPass}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupARCParamsTLSVersion{radar.EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_0}),
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
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupDKIMParamsAggInterval15m),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupDKIMParamsARC{radar.EmailSecurityTimeseriesGroupDKIMParamsARCPass}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupDKIMParamsDMARC{radar.EmailSecurityTimeseriesGroupDKIMParamsDMARCPass}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupDKIMParamsFormatJson),
		Name:        cloudflare.F([]string{"string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupDKIMParamsSPF{radar.EmailSecurityTimeseriesGroupDKIMParamsSPFPass}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupDKIMParamsTLSVersion{radar.EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_0}),
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
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupDMARCParamsAggInterval15m),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupDMARCParamsARC{radar.EmailSecurityTimeseriesGroupDMARCParamsARCPass}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupDMARCParamsDKIM{radar.EmailSecurityTimeseriesGroupDMARCParamsDKIMPass}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupDMARCParamsFormatJson),
		Name:        cloudflare.F([]string{"string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupDMARCParamsSPF{radar.EmailSecurityTimeseriesGroupDMARCParamsSPFPass}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupDMARCParamsTLSVersion{radar.EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_0}),
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
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupMaliciousParamsAggInterval15m),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsARC{radar.EmailSecurityTimeseriesGroupMaliciousParamsARCPass}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsDKIM{radar.EmailSecurityTimeseriesGroupMaliciousParamsDKIMPass}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsDMARC{radar.EmailSecurityTimeseriesGroupMaliciousParamsDMARCPass}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupMaliciousParamsFormatJson),
		Name:        cloudflare.F([]string{"string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsSPF{radar.EmailSecurityTimeseriesGroupMaliciousParamsSPFPass}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion{radar.EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_0}),
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
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupSpamParamsAggInterval15m),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpamParamsARC{radar.EmailSecurityTimeseriesGroupSpamParamsARCPass}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpamParamsDKIM{radar.EmailSecurityTimeseriesGroupSpamParamsDKIMPass}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpamParamsDMARC{radar.EmailSecurityTimeseriesGroupSpamParamsDMARCPass}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupSpamParamsFormatJson),
		Name:        cloudflare.F([]string{"string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpamParamsSPF{radar.EmailSecurityTimeseriesGroupSpamParamsSPFPass}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpamParamsTLSVersion{radar.EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_0}),
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
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupSPFParamsAggInterval15m),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupSPFParamsARC{radar.EmailSecurityTimeseriesGroupSPFParamsARCPass}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupSPFParamsDKIM{radar.EmailSecurityTimeseriesGroupSPFParamsDKIMPass}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupSPFParamsDMARC{radar.EmailSecurityTimeseriesGroupSPFParamsDMARCPass}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupSPFParamsFormatJson),
		Name:        cloudflare.F([]string{"string"}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupSPFParamsTLSVersion{radar.EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_0}),
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
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupSpoofParamsAggInterval15m),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsARC{radar.EmailSecurityTimeseriesGroupSpoofParamsARCPass}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsDKIM{radar.EmailSecurityTimeseriesGroupSpoofParamsDKIMPass}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsDMARC{radar.EmailSecurityTimeseriesGroupSpoofParamsDMARCPass}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupSpoofParamsFormatJson),
		Name:        cloudflare.F([]string{"string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsSPF{radar.EmailSecurityTimeseriesGroupSpoofParamsSPFPass}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupSpoofParamsTLSVersion{radar.EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_0}),
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
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval15m),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsARC{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsARCPass}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMPass}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCPass}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupThreatCategoryParamsFormatJson),
		Name:        cloudflare.F([]string{"string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsSPF{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsSPFPass}),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion{radar.EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_0}),
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
		AggInterval: cloudflare.F(radar.EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval15m),
		ARC:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupTLSVersionParamsARC{radar.EmailSecurityTimeseriesGroupTLSVersionParamsARCPass}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTimeseriesGroupTLSVersionParamsDKIM{radar.EmailSecurityTimeseriesGroupTLSVersionParamsDKIMPass}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTimeseriesGroupTLSVersionParamsDMARC{radar.EmailSecurityTimeseriesGroupTLSVersionParamsDMARCPass}),
		Format:      cloudflare.F(radar.EmailSecurityTimeseriesGroupTLSVersionParamsFormatJson),
		Name:        cloudflare.F([]string{"string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTimeseriesGroupTLSVersionParamsSPF{radar.EmailSecurityTimeseriesGroupTLSVersionParamsSPFPass}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
