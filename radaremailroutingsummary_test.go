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

func TestRadarEmailRoutingSummaryARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.ARC(context.TODO(), cloudflare.RadarEmailRoutingSummaryARCParams{
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailRoutingSummaryARCParamsDateRange{cloudflare.RadarEmailRoutingSummaryARCParamsDateRange1d, cloudflare.RadarEmailRoutingSummaryARCParamsDateRange2d, cloudflare.RadarEmailRoutingSummaryARCParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailRoutingSummaryARCParamsDKIM{cloudflare.RadarEmailRoutingSummaryARCParamsDKIMPass, cloudflare.RadarEmailRoutingSummaryARCParamsDKIMNone, cloudflare.RadarEmailRoutingSummaryARCParamsDKIMFail}),
		DMARC:     cloudflare.F([]cloudflare.RadarEmailRoutingSummaryARCParamsDMARC{cloudflare.RadarEmailRoutingSummaryARCParamsDMARCPass, cloudflare.RadarEmailRoutingSummaryARCParamsDMARCNone, cloudflare.RadarEmailRoutingSummaryARCParamsDMARCFail}),
		Encrypted: cloudflare.F([]cloudflare.RadarEmailRoutingSummaryARCParamsEncrypted{cloudflare.RadarEmailRoutingSummaryARCParamsEncryptedEncrypted, cloudflare.RadarEmailRoutingSummaryARCParamsEncryptedNotEncrypted}),
		Format:    cloudflare.F(cloudflare.RadarEmailRoutingSummaryARCParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarEmailRoutingSummaryARCParamsIPVersion{cloudflare.RadarEmailRoutingSummaryARCParamsIPVersionIPv4, cloudflare.RadarEmailRoutingSummaryARCParamsIPVersionIPv6}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailRoutingSummaryARCParamsSPF{cloudflare.RadarEmailRoutingSummaryARCParamsSPFPass, cloudflare.RadarEmailRoutingSummaryARCParamsSPFNone, cloudflare.RadarEmailRoutingSummaryARCParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingSummaryDKIMWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.DKIM(context.TODO(), cloudflare.RadarEmailRoutingSummaryDKIMParams{
		ARC:       cloudflare.F([]cloudflare.RadarEmailRoutingSummaryDKIMParamsARC{cloudflare.RadarEmailRoutingSummaryDKIMParamsARCPass, cloudflare.RadarEmailRoutingSummaryDKIMParamsARCNone, cloudflare.RadarEmailRoutingSummaryDKIMParamsARCFail}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailRoutingSummaryDKIMParamsDateRange{cloudflare.RadarEmailRoutingSummaryDKIMParamsDateRange1d, cloudflare.RadarEmailRoutingSummaryDKIMParamsDateRange2d, cloudflare.RadarEmailRoutingSummaryDKIMParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DMARC:     cloudflare.F([]cloudflare.RadarEmailRoutingSummaryDKIMParamsDMARC{cloudflare.RadarEmailRoutingSummaryDKIMParamsDMARCPass, cloudflare.RadarEmailRoutingSummaryDKIMParamsDMARCNone, cloudflare.RadarEmailRoutingSummaryDKIMParamsDMARCFail}),
		Encrypted: cloudflare.F([]cloudflare.RadarEmailRoutingSummaryDKIMParamsEncrypted{cloudflare.RadarEmailRoutingSummaryDKIMParamsEncryptedEncrypted, cloudflare.RadarEmailRoutingSummaryDKIMParamsEncryptedNotEncrypted}),
		Format:    cloudflare.F(cloudflare.RadarEmailRoutingSummaryDKIMParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarEmailRoutingSummaryDKIMParamsIPVersion{cloudflare.RadarEmailRoutingSummaryDKIMParamsIPVersionIPv4, cloudflare.RadarEmailRoutingSummaryDKIMParamsIPVersionIPv6}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailRoutingSummaryDKIMParamsSPF{cloudflare.RadarEmailRoutingSummaryDKIMParamsSPFPass, cloudflare.RadarEmailRoutingSummaryDKIMParamsSPFNone, cloudflare.RadarEmailRoutingSummaryDKIMParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingSummaryDMARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.DMARC(context.TODO(), cloudflare.RadarEmailRoutingSummaryDMARCParams{
		ARC:       cloudflare.F([]cloudflare.RadarEmailRoutingSummaryDMARCParamsARC{cloudflare.RadarEmailRoutingSummaryDMARCParamsARCPass, cloudflare.RadarEmailRoutingSummaryDMARCParamsARCNone, cloudflare.RadarEmailRoutingSummaryDMARCParamsARCFail}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailRoutingSummaryDMARCParamsDateRange{cloudflare.RadarEmailRoutingSummaryDMARCParamsDateRange1d, cloudflare.RadarEmailRoutingSummaryDMARCParamsDateRange2d, cloudflare.RadarEmailRoutingSummaryDMARCParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailRoutingSummaryDMARCParamsDKIM{cloudflare.RadarEmailRoutingSummaryDMARCParamsDKIMPass, cloudflare.RadarEmailRoutingSummaryDMARCParamsDKIMNone, cloudflare.RadarEmailRoutingSummaryDMARCParamsDKIMFail}),
		Encrypted: cloudflare.F([]cloudflare.RadarEmailRoutingSummaryDMARCParamsEncrypted{cloudflare.RadarEmailRoutingSummaryDMARCParamsEncryptedEncrypted, cloudflare.RadarEmailRoutingSummaryDMARCParamsEncryptedNotEncrypted}),
		Format:    cloudflare.F(cloudflare.RadarEmailRoutingSummaryDMARCParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarEmailRoutingSummaryDMARCParamsIPVersion{cloudflare.RadarEmailRoutingSummaryDMARCParamsIPVersionIPv4, cloudflare.RadarEmailRoutingSummaryDMARCParamsIPVersionIPv6}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailRoutingSummaryDMARCParamsSPF{cloudflare.RadarEmailRoutingSummaryDMARCParamsSPFPass, cloudflare.RadarEmailRoutingSummaryDMARCParamsSPFNone, cloudflare.RadarEmailRoutingSummaryDMARCParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingSummaryEncryptedWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.Encrypted(context.TODO(), cloudflare.RadarEmailRoutingSummaryEncryptedParams{
		ARC:       cloudflare.F([]cloudflare.RadarEmailRoutingSummaryEncryptedParamsARC{cloudflare.RadarEmailRoutingSummaryEncryptedParamsARCPass, cloudflare.RadarEmailRoutingSummaryEncryptedParamsARCNone, cloudflare.RadarEmailRoutingSummaryEncryptedParamsARCFail}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailRoutingSummaryEncryptedParamsDateRange{cloudflare.RadarEmailRoutingSummaryEncryptedParamsDateRange1d, cloudflare.RadarEmailRoutingSummaryEncryptedParamsDateRange2d, cloudflare.RadarEmailRoutingSummaryEncryptedParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailRoutingSummaryEncryptedParamsDKIM{cloudflare.RadarEmailRoutingSummaryEncryptedParamsDKIMPass, cloudflare.RadarEmailRoutingSummaryEncryptedParamsDKIMNone, cloudflare.RadarEmailRoutingSummaryEncryptedParamsDKIMFail}),
		DMARC:     cloudflare.F([]cloudflare.RadarEmailRoutingSummaryEncryptedParamsDMARC{cloudflare.RadarEmailRoutingSummaryEncryptedParamsDMARCPass, cloudflare.RadarEmailRoutingSummaryEncryptedParamsDMARCNone, cloudflare.RadarEmailRoutingSummaryEncryptedParamsDMARCFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailRoutingSummaryEncryptedParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarEmailRoutingSummaryEncryptedParamsIPVersion{cloudflare.RadarEmailRoutingSummaryEncryptedParamsIPVersionIPv4, cloudflare.RadarEmailRoutingSummaryEncryptedParamsIPVersionIPv6}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailRoutingSummaryEncryptedParamsSPF{cloudflare.RadarEmailRoutingSummaryEncryptedParamsSPFPass, cloudflare.RadarEmailRoutingSummaryEncryptedParamsSPFNone, cloudflare.RadarEmailRoutingSummaryEncryptedParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingSummaryIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.IPVersion(context.TODO(), cloudflare.RadarEmailRoutingSummaryIPVersionParams{
		ARC:       cloudflare.F([]cloudflare.RadarEmailRoutingSummaryIPVersionParamsARC{cloudflare.RadarEmailRoutingSummaryIPVersionParamsARCPass, cloudflare.RadarEmailRoutingSummaryIPVersionParamsARCNone, cloudflare.RadarEmailRoutingSummaryIPVersionParamsARCFail}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailRoutingSummaryIPVersionParamsDateRange{cloudflare.RadarEmailRoutingSummaryIPVersionParamsDateRange1d, cloudflare.RadarEmailRoutingSummaryIPVersionParamsDateRange2d, cloudflare.RadarEmailRoutingSummaryIPVersionParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailRoutingSummaryIPVersionParamsDKIM{cloudflare.RadarEmailRoutingSummaryIPVersionParamsDKIMPass, cloudflare.RadarEmailRoutingSummaryIPVersionParamsDKIMNone, cloudflare.RadarEmailRoutingSummaryIPVersionParamsDKIMFail}),
		DMARC:     cloudflare.F([]cloudflare.RadarEmailRoutingSummaryIPVersionParamsDMARC{cloudflare.RadarEmailRoutingSummaryIPVersionParamsDMARCPass, cloudflare.RadarEmailRoutingSummaryIPVersionParamsDMARCNone, cloudflare.RadarEmailRoutingSummaryIPVersionParamsDMARCFail}),
		Encrypted: cloudflare.F([]cloudflare.RadarEmailRoutingSummaryIPVersionParamsEncrypted{cloudflare.RadarEmailRoutingSummaryIPVersionParamsEncryptedEncrypted, cloudflare.RadarEmailRoutingSummaryIPVersionParamsEncryptedNotEncrypted}),
		Format:    cloudflare.F(cloudflare.RadarEmailRoutingSummaryIPVersionParamsFormatJson),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailRoutingSummaryIPVersionParamsSPF{cloudflare.RadarEmailRoutingSummaryIPVersionParamsSPFPass, cloudflare.RadarEmailRoutingSummaryIPVersionParamsSPFNone, cloudflare.RadarEmailRoutingSummaryIPVersionParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingSummarySPFWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.SPF(context.TODO(), cloudflare.RadarEmailRoutingSummarySPFParams{
		ARC:       cloudflare.F([]cloudflare.RadarEmailRoutingSummarySPFParamsARC{cloudflare.RadarEmailRoutingSummarySPFParamsARCPass, cloudflare.RadarEmailRoutingSummarySPFParamsARCNone, cloudflare.RadarEmailRoutingSummarySPFParamsARCFail}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailRoutingSummarySPFParamsDateRange{cloudflare.RadarEmailRoutingSummarySPFParamsDateRange1d, cloudflare.RadarEmailRoutingSummarySPFParamsDateRange2d, cloudflare.RadarEmailRoutingSummarySPFParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailRoutingSummarySPFParamsDKIM{cloudflare.RadarEmailRoutingSummarySPFParamsDKIMPass, cloudflare.RadarEmailRoutingSummarySPFParamsDKIMNone, cloudflare.RadarEmailRoutingSummarySPFParamsDKIMFail}),
		DMARC:     cloudflare.F([]cloudflare.RadarEmailRoutingSummarySPFParamsDMARC{cloudflare.RadarEmailRoutingSummarySPFParamsDMARCPass, cloudflare.RadarEmailRoutingSummarySPFParamsDMARCNone, cloudflare.RadarEmailRoutingSummarySPFParamsDMARCFail}),
		Encrypted: cloudflare.F([]cloudflare.RadarEmailRoutingSummarySPFParamsEncrypted{cloudflare.RadarEmailRoutingSummarySPFParamsEncryptedEncrypted, cloudflare.RadarEmailRoutingSummarySPFParamsEncryptedNotEncrypted}),
		Format:    cloudflare.F(cloudflare.RadarEmailRoutingSummarySPFParamsFormatJson),
		IPVersion: cloudflare.F([]cloudflare.RadarEmailRoutingSummarySPFParamsIPVersion{cloudflare.RadarEmailRoutingSummarySPFParamsIPVersionIPv4, cloudflare.RadarEmailRoutingSummarySPFParamsIPVersionIPv6}),
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
