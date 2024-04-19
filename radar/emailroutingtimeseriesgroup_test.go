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
		AggInterval: cloudflare.F(radar.EmailRoutingTimeseriesGroupARCParamsAggInterval1h),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupARCParamsDateRange{radar.EmailRoutingTimeseriesGroupARCParamsDateRange1d, radar.EmailRoutingTimeseriesGroupARCParamsDateRange2d, radar.EmailRoutingTimeseriesGroupARCParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailRoutingTimeseriesGroupARCParamsDKIM{radar.EmailRoutingTimeseriesGroupARCParamsDKIMPass, radar.EmailRoutingTimeseriesGroupARCParamsDKIMNone, radar.EmailRoutingTimeseriesGroupARCParamsDKIMFail}),
		DMARC:       cloudflare.F([]radar.EmailRoutingTimeseriesGroupARCParamsDMARC{radar.EmailRoutingTimeseriesGroupARCParamsDMARCPass, radar.EmailRoutingTimeseriesGroupARCParamsDMARCNone, radar.EmailRoutingTimeseriesGroupARCParamsDMARCFail}),
		Encrypted:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupARCParamsEncrypted{radar.EmailRoutingTimeseriesGroupARCParamsEncryptedEncrypted, radar.EmailRoutingTimeseriesGroupARCParamsEncryptedNotEncrypted}),
		Format:      cloudflare.F(radar.EmailRoutingTimeseriesGroupARCParamsFormatJson),
		IPVersion:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupARCParamsIPVersion{radar.EmailRoutingTimeseriesGroupARCParamsIPVersionIPv4, radar.EmailRoutingTimeseriesGroupARCParamsIPVersionIPv6}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupARCParamsSPF{radar.EmailRoutingTimeseriesGroupARCParamsSPFPass, radar.EmailRoutingTimeseriesGroupARCParamsSPFNone, radar.EmailRoutingTimeseriesGroupARCParamsSPFFail}),
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
		AggInterval: cloudflare.F(radar.EmailRoutingTimeseriesGroupDKIMParamsAggInterval1h),
		ARC:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsARC{radar.EmailRoutingTimeseriesGroupDKIMParamsARCPass, radar.EmailRoutingTimeseriesGroupDKIMParamsARCNone, radar.EmailRoutingTimeseriesGroupDKIMParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsDateRange{radar.EmailRoutingTimeseriesGroupDKIMParamsDateRange1d, radar.EmailRoutingTimeseriesGroupDKIMParamsDateRange2d, radar.EmailRoutingTimeseriesGroupDKIMParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DMARC:       cloudflare.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsDMARC{radar.EmailRoutingTimeseriesGroupDKIMParamsDMARCPass, radar.EmailRoutingTimeseriesGroupDKIMParamsDMARCNone, radar.EmailRoutingTimeseriesGroupDKIMParamsDMARCFail}),
		Encrypted:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsEncrypted{radar.EmailRoutingTimeseriesGroupDKIMParamsEncryptedEncrypted, radar.EmailRoutingTimeseriesGroupDKIMParamsEncryptedNotEncrypted}),
		Format:      cloudflare.F(radar.EmailRoutingTimeseriesGroupDKIMParamsFormatJson),
		IPVersion:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsIPVersion{radar.EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv4, radar.EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv6}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupDKIMParamsSPF{radar.EmailRoutingTimeseriesGroupDKIMParamsSPFPass, radar.EmailRoutingTimeseriesGroupDKIMParamsSPFNone, radar.EmailRoutingTimeseriesGroupDKIMParamsSPFFail}),
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
		AggInterval: cloudflare.F(radar.EmailRoutingTimeseriesGroupDMARCParamsAggInterval1h),
		ARC:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsARC{radar.EmailRoutingTimeseriesGroupDMARCParamsARCPass, radar.EmailRoutingTimeseriesGroupDMARCParamsARCNone, radar.EmailRoutingTimeseriesGroupDMARCParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsDateRange{radar.EmailRoutingTimeseriesGroupDMARCParamsDateRange1d, radar.EmailRoutingTimeseriesGroupDMARCParamsDateRange2d, radar.EmailRoutingTimeseriesGroupDMARCParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsDKIM{radar.EmailRoutingTimeseriesGroupDMARCParamsDKIMPass, radar.EmailRoutingTimeseriesGroupDMARCParamsDKIMNone, radar.EmailRoutingTimeseriesGroupDMARCParamsDKIMFail}),
		Encrypted:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsEncrypted{radar.EmailRoutingTimeseriesGroupDMARCParamsEncryptedEncrypted, radar.EmailRoutingTimeseriesGroupDMARCParamsEncryptedNotEncrypted}),
		Format:      cloudflare.F(radar.EmailRoutingTimeseriesGroupDMARCParamsFormatJson),
		IPVersion:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsIPVersion{radar.EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv4, radar.EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv6}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupDMARCParamsSPF{radar.EmailRoutingTimeseriesGroupDMARCParamsSPFPass, radar.EmailRoutingTimeseriesGroupDMARCParamsSPFNone, radar.EmailRoutingTimeseriesGroupDMARCParamsSPFFail}),
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
		AggInterval: cloudflare.F(radar.EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1h),
		ARC:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsARC{radar.EmailRoutingTimeseriesGroupEncryptedParamsARCPass, radar.EmailRoutingTimeseriesGroupEncryptedParamsARCNone, radar.EmailRoutingTimeseriesGroupEncryptedParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsDateRange{radar.EmailRoutingTimeseriesGroupEncryptedParamsDateRange1d, radar.EmailRoutingTimeseriesGroupEncryptedParamsDateRange2d, radar.EmailRoutingTimeseriesGroupEncryptedParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsDKIM{radar.EmailRoutingTimeseriesGroupEncryptedParamsDKIMPass, radar.EmailRoutingTimeseriesGroupEncryptedParamsDKIMNone, radar.EmailRoutingTimeseriesGroupEncryptedParamsDKIMFail}),
		DMARC:       cloudflare.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsDMARC{radar.EmailRoutingTimeseriesGroupEncryptedParamsDMARCPass, radar.EmailRoutingTimeseriesGroupEncryptedParamsDMARCNone, radar.EmailRoutingTimeseriesGroupEncryptedParamsDMARCFail}),
		Format:      cloudflare.F(radar.EmailRoutingTimeseriesGroupEncryptedParamsFormatJson),
		IPVersion:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsIPVersion{radar.EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv4, radar.EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv6}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupEncryptedParamsSPF{radar.EmailRoutingTimeseriesGroupEncryptedParamsSPFPass, radar.EmailRoutingTimeseriesGroupEncryptedParamsSPFNone, radar.EmailRoutingTimeseriesGroupEncryptedParamsSPFFail}),
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
		AggInterval: cloudflare.F(radar.EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1h),
		ARC:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsARC{radar.EmailRoutingTimeseriesGroupIPVersionParamsARCPass, radar.EmailRoutingTimeseriesGroupIPVersionParamsARCNone, radar.EmailRoutingTimeseriesGroupIPVersionParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsDateRange{radar.EmailRoutingTimeseriesGroupIPVersionParamsDateRange1d, radar.EmailRoutingTimeseriesGroupIPVersionParamsDateRange2d, radar.EmailRoutingTimeseriesGroupIPVersionParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsDKIM{radar.EmailRoutingTimeseriesGroupIPVersionParamsDKIMPass, radar.EmailRoutingTimeseriesGroupIPVersionParamsDKIMNone, radar.EmailRoutingTimeseriesGroupIPVersionParamsDKIMFail}),
		DMARC:       cloudflare.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsDMARC{radar.EmailRoutingTimeseriesGroupIPVersionParamsDMARCPass, radar.EmailRoutingTimeseriesGroupIPVersionParamsDMARCNone, radar.EmailRoutingTimeseriesGroupIPVersionParamsDMARCFail}),
		Encrypted:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsEncrypted{radar.EmailRoutingTimeseriesGroupIPVersionParamsEncryptedEncrypted, radar.EmailRoutingTimeseriesGroupIPVersionParamsEncryptedNotEncrypted}),
		Format:      cloudflare.F(radar.EmailRoutingTimeseriesGroupIPVersionParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupIPVersionParamsSPF{radar.EmailRoutingTimeseriesGroupIPVersionParamsSPFPass, radar.EmailRoutingTimeseriesGroupIPVersionParamsSPFNone, radar.EmailRoutingTimeseriesGroupIPVersionParamsSPFFail}),
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
		AggInterval: cloudflare.F(radar.EmailRoutingTimeseriesGroupSPFParamsAggInterval1h),
		ARC:         cloudflare.F([]radar.EmailRoutingTimeseriesGroupSPFParamsARC{radar.EmailRoutingTimeseriesGroupSPFParamsARCPass, radar.EmailRoutingTimeseriesGroupSPFParamsARCNone, radar.EmailRoutingTimeseriesGroupSPFParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupSPFParamsDateRange{radar.EmailRoutingTimeseriesGroupSPFParamsDateRange1d, radar.EmailRoutingTimeseriesGroupSPFParamsDateRange2d, radar.EmailRoutingTimeseriesGroupSPFParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailRoutingTimeseriesGroupSPFParamsDKIM{radar.EmailRoutingTimeseriesGroupSPFParamsDKIMPass, radar.EmailRoutingTimeseriesGroupSPFParamsDKIMNone, radar.EmailRoutingTimeseriesGroupSPFParamsDKIMFail}),
		DMARC:       cloudflare.F([]radar.EmailRoutingTimeseriesGroupSPFParamsDMARC{radar.EmailRoutingTimeseriesGroupSPFParamsDMARCPass, radar.EmailRoutingTimeseriesGroupSPFParamsDMARCNone, radar.EmailRoutingTimeseriesGroupSPFParamsDMARCFail}),
		Encrypted:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupSPFParamsEncrypted{radar.EmailRoutingTimeseriesGroupSPFParamsEncryptedEncrypted, radar.EmailRoutingTimeseriesGroupSPFParamsEncryptedNotEncrypted}),
		Format:      cloudflare.F(radar.EmailRoutingTimeseriesGroupSPFParamsFormatJson),
		IPVersion:   cloudflare.F([]radar.EmailRoutingTimeseriesGroupSPFParamsIPVersion{radar.EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv4, radar.EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv6}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
