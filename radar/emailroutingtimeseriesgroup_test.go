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

func TestEmailRoutingTimeseriesGroupARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.ARC(context.TODO(), radar.EmailRoutingTimeseriesGroupARCParams{
		AggInterval: cloudflare.F(radar.EmailRoutingTimeseriesGroupARCParamsAggInterval15m),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailRoutingTimeseriesGroupARCParamsDKIM{radar.EmailRoutingTimeseriesGroupARCParamsDKIMPass}),
		DMARC:       cloudflare.F([]radar.EmailRoutingTimeseriesGroupARCParamsDMARC{radar.EmailRoutingTimeseriesGroupARCParamsDMARCPass}),
		Encrypted:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupARCParamsEncrypted{radar.EmailRoutingTimeseriesGroupARCParamsEncryptedEncrypted}),
		Format:      cloudflare.F(radar.EmailRoutingTimeseriesGroupARCParamsFormatJson),
		IPVersion:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupARCParamsIPVersion{radar.EmailRoutingTimeseriesGroupARCParamsIPVersionIPv4}),
		Name:        cloudflare.F([]string{"string"}),
		SPF:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupARCParamsSPF{radar.EmailRoutingTimeseriesGroupARCParamsSPFPass}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingTimeseriesGroupDKIMWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.DKIM(context.TODO(), radar.EmailRoutingTimeseriesGroupDKIMParams{
		AggInterval: cloudflare.F(radar.EmailRoutingTimeseriesGroupDKIMParamsAggInterval15m),
		ARC:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsARC{radar.EmailRoutingTimeseriesGroupDKIMParamsARCPass}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DMARC:       cloudflare.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsDMARC{radar.EmailRoutingTimeseriesGroupDKIMParamsDMARCPass}),
		Encrypted:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsEncrypted{radar.EmailRoutingTimeseriesGroupDKIMParamsEncryptedEncrypted}),
		Format:      cloudflare.F(radar.EmailRoutingTimeseriesGroupDKIMParamsFormatJson),
		IPVersion:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsIPVersion{radar.EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv4}),
		Name:        cloudflare.F([]string{"string"}),
		SPF:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsSPF{radar.EmailRoutingTimeseriesGroupDKIMParamsSPFPass}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingTimeseriesGroupDMARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.DMARC(context.TODO(), radar.EmailRoutingTimeseriesGroupDMARCParams{
		AggInterval: cloudflare.F(radar.EmailRoutingTimeseriesGroupDMARCParamsAggInterval15m),
		ARC:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsARC{radar.EmailRoutingTimeseriesGroupDMARCParamsARCPass}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsDKIM{radar.EmailRoutingTimeseriesGroupDMARCParamsDKIMPass}),
		Encrypted:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsEncrypted{radar.EmailRoutingTimeseriesGroupDMARCParamsEncryptedEncrypted}),
		Format:      cloudflare.F(radar.EmailRoutingTimeseriesGroupDMARCParamsFormatJson),
		IPVersion:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsIPVersion{radar.EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv4}),
		Name:        cloudflare.F([]string{"string"}),
		SPF:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsSPF{radar.EmailRoutingTimeseriesGroupDMARCParamsSPFPass}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingTimeseriesGroupEncryptedWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.Encrypted(context.TODO(), radar.EmailRoutingTimeseriesGroupEncryptedParams{
		AggInterval: cloudflare.F(radar.EmailRoutingTimeseriesGroupEncryptedParamsAggInterval15m),
		ARC:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsARC{radar.EmailRoutingTimeseriesGroupEncryptedParamsARCPass}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsDKIM{radar.EmailRoutingTimeseriesGroupEncryptedParamsDKIMPass}),
		DMARC:       cloudflare.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsDMARC{radar.EmailRoutingTimeseriesGroupEncryptedParamsDMARCPass}),
		Format:      cloudflare.F(radar.EmailRoutingTimeseriesGroupEncryptedParamsFormatJson),
		IPVersion:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsIPVersion{radar.EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv4}),
		Name:        cloudflare.F([]string{"string"}),
		SPF:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsSPF{radar.EmailRoutingTimeseriesGroupEncryptedParamsSPFPass}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingTimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.IPVersion(context.TODO(), radar.EmailRoutingTimeseriesGroupIPVersionParams{
		AggInterval: cloudflare.F(radar.EmailRoutingTimeseriesGroupIPVersionParamsAggInterval15m),
		ARC:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsARC{radar.EmailRoutingTimeseriesGroupIPVersionParamsARCPass}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsDKIM{radar.EmailRoutingTimeseriesGroupIPVersionParamsDKIMPass}),
		DMARC:       cloudflare.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsDMARC{radar.EmailRoutingTimeseriesGroupIPVersionParamsDMARCPass}),
		Encrypted:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsEncrypted{radar.EmailRoutingTimeseriesGroupIPVersionParamsEncryptedEncrypted}),
		Format:      cloudflare.F(radar.EmailRoutingTimeseriesGroupIPVersionParamsFormatJson),
		Name:        cloudflare.F([]string{"string"}),
		SPF:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsSPF{radar.EmailRoutingTimeseriesGroupIPVersionParamsSPFPass}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingTimeseriesGroupSPFWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.SPF(context.TODO(), radar.EmailRoutingTimeseriesGroupSPFParams{
		AggInterval: cloudflare.F(radar.EmailRoutingTimeseriesGroupSPFParamsAggInterval15m),
		ARC:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupSPFParamsARC{radar.EmailRoutingTimeseriesGroupSPFParamsARCPass}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailRoutingTimeseriesGroupSPFParamsDKIM{radar.EmailRoutingTimeseriesGroupSPFParamsDKIMPass}),
		DMARC:       cloudflare.F([]radar.EmailRoutingTimeseriesGroupSPFParamsDMARC{radar.EmailRoutingTimeseriesGroupSPFParamsDMARCPass}),
		Encrypted:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupSPFParamsEncrypted{radar.EmailRoutingTimeseriesGroupSPFParamsEncryptedEncrypted}),
		Format:      cloudflare.F(radar.EmailRoutingTimeseriesGroupSPFParamsFormatJson),
		IPVersion:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupSPFParamsIPVersion{radar.EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv4}),
		Name:        cloudflare.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
