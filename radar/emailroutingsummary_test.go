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
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]radar.EmailRoutingSummaryARCParamsDKIM{radar.EmailRoutingSummaryARCParamsDKIMPass, radar.EmailRoutingSummaryARCParamsDKIMNone, radar.EmailRoutingSummaryARCParamsDKIMFail}),
		DMARC:     cloudflare.F([]radar.EmailRoutingSummaryARCParamsDMARC{radar.EmailRoutingSummaryARCParamsDMARCPass, radar.EmailRoutingSummaryARCParamsDMARCNone, radar.EmailRoutingSummaryARCParamsDMARCFail}),
		Encrypted: cloudflare.F([]radar.EmailRoutingSummaryARCParamsEncrypted{radar.EmailRoutingSummaryARCParamsEncryptedEncrypted, radar.EmailRoutingSummaryARCParamsEncryptedNotEncrypted}),
		Format:    cloudflare.F(radar.EmailRoutingSummaryARCParamsFormatJson),
		IPVersion: cloudflare.F([]radar.EmailRoutingSummaryARCParamsIPVersion{radar.EmailRoutingSummaryARCParamsIPVersionIPv4, radar.EmailRoutingSummaryARCParamsIPVersionIPv6}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]radar.EmailRoutingSummaryARCParamsSPF{radar.EmailRoutingSummaryARCParamsSPFPass, radar.EmailRoutingSummaryARCParamsSPFNone, radar.EmailRoutingSummaryARCParamsSPFFail}),
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
		ARC:       cloudflare.F([]radar.EmailRoutingSummaryDKIMParamsARC{radar.EmailRoutingSummaryDKIMParamsARCPass, radar.EmailRoutingSummaryDKIMParamsARCNone, radar.EmailRoutingSummaryDKIMParamsARCFail}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DMARC:     cloudflare.F([]radar.EmailRoutingSummaryDKIMParamsDMARC{radar.EmailRoutingSummaryDKIMParamsDMARCPass, radar.EmailRoutingSummaryDKIMParamsDMARCNone, radar.EmailRoutingSummaryDKIMParamsDMARCFail}),
		Encrypted: cloudflare.F([]radar.EmailRoutingSummaryDKIMParamsEncrypted{radar.EmailRoutingSummaryDKIMParamsEncryptedEncrypted, radar.EmailRoutingSummaryDKIMParamsEncryptedNotEncrypted}),
		Format:    cloudflare.F(radar.EmailRoutingSummaryDKIMParamsFormatJson),
		IPVersion: cloudflare.F([]radar.EmailRoutingSummaryDKIMParamsIPVersion{radar.EmailRoutingSummaryDKIMParamsIPVersionIPv4, radar.EmailRoutingSummaryDKIMParamsIPVersionIPv6}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]radar.EmailRoutingSummaryDKIMParamsSPF{radar.EmailRoutingSummaryDKIMParamsSPFPass, radar.EmailRoutingSummaryDKIMParamsSPFNone, radar.EmailRoutingSummaryDKIMParamsSPFFail}),
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
		ARC:       cloudflare.F([]radar.EmailRoutingSummaryDMARCParamsARC{radar.EmailRoutingSummaryDMARCParamsARCPass, radar.EmailRoutingSummaryDMARCParamsARCNone, radar.EmailRoutingSummaryDMARCParamsARCFail}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]radar.EmailRoutingSummaryDMARCParamsDKIM{radar.EmailRoutingSummaryDMARCParamsDKIMPass, radar.EmailRoutingSummaryDMARCParamsDKIMNone, radar.EmailRoutingSummaryDMARCParamsDKIMFail}),
		Encrypted: cloudflare.F([]radar.EmailRoutingSummaryDMARCParamsEncrypted{radar.EmailRoutingSummaryDMARCParamsEncryptedEncrypted, radar.EmailRoutingSummaryDMARCParamsEncryptedNotEncrypted}),
		Format:    cloudflare.F(radar.EmailRoutingSummaryDMARCParamsFormatJson),
		IPVersion: cloudflare.F([]radar.EmailRoutingSummaryDMARCParamsIPVersion{radar.EmailRoutingSummaryDMARCParamsIPVersionIPv4, radar.EmailRoutingSummaryDMARCParamsIPVersionIPv6}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]radar.EmailRoutingSummaryDMARCParamsSPF{radar.EmailRoutingSummaryDMARCParamsSPFPass, radar.EmailRoutingSummaryDMARCParamsSPFNone, radar.EmailRoutingSummaryDMARCParamsSPFFail}),
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
		ARC:       cloudflare.F([]radar.EmailRoutingSummaryEncryptedParamsARC{radar.EmailRoutingSummaryEncryptedParamsARCPass, radar.EmailRoutingSummaryEncryptedParamsARCNone, radar.EmailRoutingSummaryEncryptedParamsARCFail}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]radar.EmailRoutingSummaryEncryptedParamsDKIM{radar.EmailRoutingSummaryEncryptedParamsDKIMPass, radar.EmailRoutingSummaryEncryptedParamsDKIMNone, radar.EmailRoutingSummaryEncryptedParamsDKIMFail}),
		DMARC:     cloudflare.F([]radar.EmailRoutingSummaryEncryptedParamsDMARC{radar.EmailRoutingSummaryEncryptedParamsDMARCPass, radar.EmailRoutingSummaryEncryptedParamsDMARCNone, radar.EmailRoutingSummaryEncryptedParamsDMARCFail}),
		Format:    cloudflare.F(radar.EmailRoutingSummaryEncryptedParamsFormatJson),
		IPVersion: cloudflare.F([]radar.EmailRoutingSummaryEncryptedParamsIPVersion{radar.EmailRoutingSummaryEncryptedParamsIPVersionIPv4, radar.EmailRoutingSummaryEncryptedParamsIPVersionIPv6}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]radar.EmailRoutingSummaryEncryptedParamsSPF{radar.EmailRoutingSummaryEncryptedParamsSPFPass, radar.EmailRoutingSummaryEncryptedParamsSPFNone, radar.EmailRoutingSummaryEncryptedParamsSPFFail}),
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
		ARC:       cloudflare.F([]radar.EmailRoutingSummaryIPVersionParamsARC{radar.EmailRoutingSummaryIPVersionParamsARCPass, radar.EmailRoutingSummaryIPVersionParamsARCNone, radar.EmailRoutingSummaryIPVersionParamsARCFail}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]radar.EmailRoutingSummaryIPVersionParamsDKIM{radar.EmailRoutingSummaryIPVersionParamsDKIMPass, radar.EmailRoutingSummaryIPVersionParamsDKIMNone, radar.EmailRoutingSummaryIPVersionParamsDKIMFail}),
		DMARC:     cloudflare.F([]radar.EmailRoutingSummaryIPVersionParamsDMARC{radar.EmailRoutingSummaryIPVersionParamsDMARCPass, radar.EmailRoutingSummaryIPVersionParamsDMARCNone, radar.EmailRoutingSummaryIPVersionParamsDMARCFail}),
		Encrypted: cloudflare.F([]radar.EmailRoutingSummaryIPVersionParamsEncrypted{radar.EmailRoutingSummaryIPVersionParamsEncryptedEncrypted, radar.EmailRoutingSummaryIPVersionParamsEncryptedNotEncrypted}),
		Format:    cloudflare.F(radar.EmailRoutingSummaryIPVersionParamsFormatJson),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]radar.EmailRoutingSummaryIPVersionParamsSPF{radar.EmailRoutingSummaryIPVersionParamsSPFPass, radar.EmailRoutingSummaryIPVersionParamsSPFNone, radar.EmailRoutingSummaryIPVersionParamsSPFFail}),
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
		ARC:       cloudflare.F([]radar.EmailRoutingSummarySPFParamsARC{radar.EmailRoutingSummarySPFParamsARCPass, radar.EmailRoutingSummarySPFParamsARCNone, radar.EmailRoutingSummarySPFParamsARCFail}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]radar.EmailRoutingSummarySPFParamsDKIM{radar.EmailRoutingSummarySPFParamsDKIMPass, radar.EmailRoutingSummarySPFParamsDKIMNone, radar.EmailRoutingSummarySPFParamsDKIMFail}),
		DMARC:     cloudflare.F([]radar.EmailRoutingSummarySPFParamsDMARC{radar.EmailRoutingSummarySPFParamsDMARCPass, radar.EmailRoutingSummarySPFParamsDMARCNone, radar.EmailRoutingSummarySPFParamsDMARCFail}),
		Encrypted: cloudflare.F([]radar.EmailRoutingSummarySPFParamsEncrypted{radar.EmailRoutingSummarySPFParamsEncryptedEncrypted, radar.EmailRoutingSummarySPFParamsEncryptedNotEncrypted}),
		Format:    cloudflare.F(radar.EmailRoutingSummarySPFParamsFormatJson),
		IPVersion: cloudflare.F([]radar.EmailRoutingSummarySPFParamsIPVersion{radar.EmailRoutingSummarySPFParamsIPVersionIPv4, radar.EmailRoutingSummarySPFParamsIPVersionIPv6}),
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
