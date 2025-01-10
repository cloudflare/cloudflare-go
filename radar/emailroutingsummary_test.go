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

func TestEmailRoutingSummaryARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.ARC(context.TODO(), radar.EmailRoutingSummaryARCParams{
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		DKIM:      cloudflare.F([]radar.EmailRoutingSummaryARCParamsDKIM{radar.EmailRoutingSummaryARCParamsDKIMPass}),
		DMARC:     cloudflare.F([]radar.EmailRoutingSummaryARCParamsDMARC{radar.EmailRoutingSummaryARCParamsDMARCPass}),
		Encrypted: cloudflare.F([]radar.EmailRoutingSummaryARCParamsEncrypted{radar.EmailRoutingSummaryARCParamsEncryptedEncrypted}),
		Format:    cloudflare.F(radar.EmailRoutingSummaryARCParamsFormatJson),
		IPVersion: cloudflare.F([]radar.EmailRoutingSummaryARCParamsIPVersion{radar.EmailRoutingSummaryARCParamsIPVersionIPv4}),
		Name:      cloudflare.F([]string{"string"}),
		SPF:       cloudflare.F([]radar.EmailRoutingSummaryARCParamsSPF{radar.EmailRoutingSummaryARCParamsSPFPass}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingSummaryDKIMWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.DKIM(context.TODO(), radar.EmailRoutingSummaryDKIMParams{
		ARC:       cloudflare.F([]radar.EmailRoutingSummaryDKIMParamsARC{radar.EmailRoutingSummaryDKIMParamsARCPass}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		DMARC:     cloudflare.F([]radar.EmailRoutingSummaryDKIMParamsDMARC{radar.EmailRoutingSummaryDKIMParamsDMARCPass}),
		Encrypted: cloudflare.F([]radar.EmailRoutingSummaryDKIMParamsEncrypted{radar.EmailRoutingSummaryDKIMParamsEncryptedEncrypted}),
		Format:    cloudflare.F(radar.EmailRoutingSummaryDKIMParamsFormatJson),
		IPVersion: cloudflare.F([]radar.EmailRoutingSummaryDKIMParamsIPVersion{radar.EmailRoutingSummaryDKIMParamsIPVersionIPv4}),
		Name:      cloudflare.F([]string{"string"}),
		SPF:       cloudflare.F([]radar.EmailRoutingSummaryDKIMParamsSPF{radar.EmailRoutingSummaryDKIMParamsSPFPass}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingSummaryDMARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.DMARC(context.TODO(), radar.EmailRoutingSummaryDMARCParams{
		ARC:       cloudflare.F([]radar.EmailRoutingSummaryDMARCParamsARC{radar.EmailRoutingSummaryDMARCParamsARCPass}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		DKIM:      cloudflare.F([]radar.EmailRoutingSummaryDMARCParamsDKIM{radar.EmailRoutingSummaryDMARCParamsDKIMPass}),
		Encrypted: cloudflare.F([]radar.EmailRoutingSummaryDMARCParamsEncrypted{radar.EmailRoutingSummaryDMARCParamsEncryptedEncrypted}),
		Format:    cloudflare.F(radar.EmailRoutingSummaryDMARCParamsFormatJson),
		IPVersion: cloudflare.F([]radar.EmailRoutingSummaryDMARCParamsIPVersion{radar.EmailRoutingSummaryDMARCParamsIPVersionIPv4}),
		Name:      cloudflare.F([]string{"string"}),
		SPF:       cloudflare.F([]radar.EmailRoutingSummaryDMARCParamsSPF{radar.EmailRoutingSummaryDMARCParamsSPFPass}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingSummaryEncryptedWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.Encrypted(context.TODO(), radar.EmailRoutingSummaryEncryptedParams{
		ARC:       cloudflare.F([]radar.EmailRoutingSummaryEncryptedParamsARC{radar.EmailRoutingSummaryEncryptedParamsARCPass}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		DKIM:      cloudflare.F([]radar.EmailRoutingSummaryEncryptedParamsDKIM{radar.EmailRoutingSummaryEncryptedParamsDKIMPass}),
		DMARC:     cloudflare.F([]radar.EmailRoutingSummaryEncryptedParamsDMARC{radar.EmailRoutingSummaryEncryptedParamsDMARCPass}),
		Format:    cloudflare.F(radar.EmailRoutingSummaryEncryptedParamsFormatJson),
		IPVersion: cloudflare.F([]radar.EmailRoutingSummaryEncryptedParamsIPVersion{radar.EmailRoutingSummaryEncryptedParamsIPVersionIPv4}),
		Name:      cloudflare.F([]string{"string"}),
		SPF:       cloudflare.F([]radar.EmailRoutingSummaryEncryptedParamsSPF{radar.EmailRoutingSummaryEncryptedParamsSPFPass}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingSummaryIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.IPVersion(context.TODO(), radar.EmailRoutingSummaryIPVersionParams{
		ARC:       cloudflare.F([]radar.EmailRoutingSummaryIPVersionParamsARC{radar.EmailRoutingSummaryIPVersionParamsARCPass}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		DKIM:      cloudflare.F([]radar.EmailRoutingSummaryIPVersionParamsDKIM{radar.EmailRoutingSummaryIPVersionParamsDKIMPass}),
		DMARC:     cloudflare.F([]radar.EmailRoutingSummaryIPVersionParamsDMARC{radar.EmailRoutingSummaryIPVersionParamsDMARCPass}),
		Encrypted: cloudflare.F([]radar.EmailRoutingSummaryIPVersionParamsEncrypted{radar.EmailRoutingSummaryIPVersionParamsEncryptedEncrypted}),
		Format:    cloudflare.F(radar.EmailRoutingSummaryIPVersionParamsFormatJson),
		Name:      cloudflare.F([]string{"string"}),
		SPF:       cloudflare.F([]radar.EmailRoutingSummaryIPVersionParamsSPF{radar.EmailRoutingSummaryIPVersionParamsSPFPass}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingSummarySPFWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.SPF(context.TODO(), radar.EmailRoutingSummarySPFParams{
		ARC:       cloudflare.F([]radar.EmailRoutingSummarySPFParamsARC{radar.EmailRoutingSummarySPFParamsARCPass}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		DKIM:      cloudflare.F([]radar.EmailRoutingSummarySPFParamsDKIM{radar.EmailRoutingSummarySPFParamsDKIMPass}),
		DMARC:     cloudflare.F([]radar.EmailRoutingSummarySPFParamsDMARC{radar.EmailRoutingSummarySPFParamsDMARCPass}),
		Encrypted: cloudflare.F([]radar.EmailRoutingSummarySPFParamsEncrypted{radar.EmailRoutingSummarySPFParamsEncryptedEncrypted}),
		Format:    cloudflare.F(radar.EmailRoutingSummarySPFParamsFormatJson),
		IPVersion: cloudflare.F([]radar.EmailRoutingSummarySPFParamsIPVersion{radar.EmailRoutingSummarySPFParamsIPVersionIPv4}),
		Name:      cloudflare.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
