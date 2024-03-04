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

func TestRadarEmailRoutingTimeseriesGroupARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.ARC(context.TODO(), cloudflare.RadarEmailRoutingTimeseriesGroupARCParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsAggInterval1h),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsDateRange{cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsDateRange1d, cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsDateRange2d, cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsDKIM{cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsDKIMPass, cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsDKIMNone, cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsDKIMFail}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsDMARC{cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsDMARCPass, cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsDMARCNone, cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsDMARCFail}),
		Encrypted:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsEncrypted{cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsEncryptedEncrypted, cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsEncryptedNotEncrypted}),
		Format:      cloudflare.F(cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsFormatJson),
		IPVersion:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsIPVersion{cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsIPVersionIPv4, cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsIPVersionIPv6}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsSPF{cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsSPFPass, cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsSPFNone, cloudflare.RadarEmailRoutingTimeseriesGroupARCParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingTimeseriesGroupDKIMWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.DKIM(context.TODO(), cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsAggInterval1h),
		ARC:         cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsARC{cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsARCPass, cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsARCNone, cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange{cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange1d, cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange2d, cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsDMARC{cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsDMARCPass, cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsDMARCNone, cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsDMARCFail}),
		Encrypted:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsEncrypted{cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsEncryptedEncrypted, cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsEncryptedNotEncrypted}),
		Format:      cloudflare.F(cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsFormatJson),
		IPVersion:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsIPVersion{cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv4, cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv6}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsSPF{cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsSPFPass, cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsSPFNone, cloudflare.RadarEmailRoutingTimeseriesGroupDKIMParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingTimeseriesGroupDMARCWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.DMARC(context.TODO(), cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsAggInterval1h),
		ARC:         cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsARC{cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsARCPass, cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsARCNone, cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange{cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange1d, cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange2d, cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsDKIM{cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsDKIMPass, cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsDKIMNone, cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsDKIMFail}),
		Encrypted:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsEncrypted{cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsEncryptedEncrypted, cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsEncryptedNotEncrypted}),
		Format:      cloudflare.F(cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsFormatJson),
		IPVersion:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsIPVersion{cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv4, cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv6}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsSPF{cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsSPFPass, cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsSPFNone, cloudflare.RadarEmailRoutingTimeseriesGroupDMARCParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingTimeseriesGroupEncryptedWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.Encrypted(context.TODO(), cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsAggInterval1h),
		ARC:         cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsARC{cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsARCPass, cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsARCNone, cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange{cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange1d, cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange2d, cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsDKIM{cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsDKIMPass, cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsDKIMNone, cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsDKIMFail}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsDMARC{cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsDMARCPass, cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsDMARCNone, cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsDMARCFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsFormatJson),
		IPVersion:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsIPVersion{cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv4, cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv6}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsSPF{cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsSPFPass, cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsSPFNone, cloudflare.RadarEmailRoutingTimeseriesGroupEncryptedParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingTimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.IPVersion(context.TODO(), cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsAggInterval1h),
		ARC:         cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsARC{cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsARCPass, cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsARCNone, cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange{cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange1d, cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange2d, cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsDKIM{cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsDKIMPass, cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsDKIMNone, cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsDKIMFail}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsDMARC{cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsDMARCPass, cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsDMARCNone, cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsDMARCFail}),
		Encrypted:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsEncrypted{cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsEncryptedEncrypted, cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsEncryptedNotEncrypted}),
		Format:      cloudflare.F(cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsSPF{cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsSPFPass, cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsSPFNone, cloudflare.RadarEmailRoutingTimeseriesGroupIPVersionParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingTimeseriesGroupSPFWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.SPF(context.TODO(), cloudflare.RadarEmailRoutingTimeseriesGroupSPFParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsAggInterval1h),
		ARC:         cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsARC{cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsARCPass, cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsARCNone, cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsDateRange{cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsDateRange1d, cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsDateRange2d, cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsDKIM{cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsDKIMPass, cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsDKIMNone, cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsDKIMFail}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsDMARC{cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsDMARCPass, cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsDMARCNone, cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsDMARCFail}),
		Encrypted:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsEncrypted{cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsEncryptedEncrypted, cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsEncryptedNotEncrypted}),
		Format:      cloudflare.F(cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsFormatJson),
		IPVersion:   cloudflare.F([]cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsIPVersion{cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsIPVersionIPv4, cloudflare.RadarEmailRoutingTimeseriesGroupSPFParamsIPVersionIPv6}),
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
