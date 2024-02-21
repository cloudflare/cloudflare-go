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

func TestRadarEmailSecuritySummaryArcWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Email.Security.Summary.Arc(context.TODO(), cloudflare.RadarEmailSecuritySummaryArcParams{
		Asn:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailSecuritySummaryArcParamsDateRange{cloudflare.RadarEmailSecuritySummaryArcParamsDateRange1d, cloudflare.RadarEmailSecuritySummaryArcParamsDateRange2d, cloudflare.RadarEmailSecuritySummaryArcParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailSecuritySummaryArcParamsDKIM{cloudflare.RadarEmailSecuritySummaryArcParamsDKIMPass, cloudflare.RadarEmailSecuritySummaryArcParamsDKIMNone, cloudflare.RadarEmailSecuritySummaryArcParamsDKIMFail}),
		Dmarc:     cloudflare.F([]cloudflare.RadarEmailSecuritySummaryArcParamsDmarc{cloudflare.RadarEmailSecuritySummaryArcParamsDmarcPass, cloudflare.RadarEmailSecuritySummaryArcParamsDmarcNone, cloudflare.RadarEmailSecuritySummaryArcParamsDmarcFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailSecuritySummaryArcParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryArcParamsSPF{cloudflare.RadarEmailSecuritySummaryArcParamsSPFPass, cloudflare.RadarEmailSecuritySummaryArcParamsSPFNone, cloudflare.RadarEmailSecuritySummaryArcParamsSPFFail}),
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Email.Security.Summary.DKIM(context.TODO(), cloudflare.RadarEmailSecuritySummaryDKIMParams{
		Arc:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDKIMParamsArc{cloudflare.RadarEmailSecuritySummaryDKIMParamsArcPass, cloudflare.RadarEmailSecuritySummaryDKIMParamsArcNone, cloudflare.RadarEmailSecuritySummaryDKIMParamsArcFail}),
		Asn:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDKIMParamsDateRange{cloudflare.RadarEmailSecuritySummaryDKIMParamsDateRange1d, cloudflare.RadarEmailSecuritySummaryDKIMParamsDateRange2d, cloudflare.RadarEmailSecuritySummaryDKIMParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Dmarc:     cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDKIMParamsDmarc{cloudflare.RadarEmailSecuritySummaryDKIMParamsDmarcPass, cloudflare.RadarEmailSecuritySummaryDKIMParamsDmarcNone, cloudflare.RadarEmailSecuritySummaryDKIMParamsDmarcFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailSecuritySummaryDKIMParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDKIMParamsSPF{cloudflare.RadarEmailSecuritySummaryDKIMParamsSPFPass, cloudflare.RadarEmailSecuritySummaryDKIMParamsSPFNone, cloudflare.RadarEmailSecuritySummaryDKIMParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryDmarcWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Email.Security.Summary.Dmarc(context.TODO(), cloudflare.RadarEmailSecuritySummaryDmarcParams{
		Arc:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDmarcParamsArc{cloudflare.RadarEmailSecuritySummaryDmarcParamsArcPass, cloudflare.RadarEmailSecuritySummaryDmarcParamsArcNone, cloudflare.RadarEmailSecuritySummaryDmarcParamsArcFail}),
		Asn:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDmarcParamsDateRange{cloudflare.RadarEmailSecuritySummaryDmarcParamsDateRange1d, cloudflare.RadarEmailSecuritySummaryDmarcParamsDateRange2d, cloudflare.RadarEmailSecuritySummaryDmarcParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDmarcParamsDKIM{cloudflare.RadarEmailSecuritySummaryDmarcParamsDKIMPass, cloudflare.RadarEmailSecuritySummaryDmarcParamsDKIMNone, cloudflare.RadarEmailSecuritySummaryDmarcParamsDKIMFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailSecuritySummaryDmarcParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryDmarcParamsSPF{cloudflare.RadarEmailSecuritySummaryDmarcParamsSPFPass, cloudflare.RadarEmailSecuritySummaryDmarcParamsSPFNone, cloudflare.RadarEmailSecuritySummaryDmarcParamsSPFFail}),
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Email.Security.Summary.Malicious(context.TODO(), cloudflare.RadarEmailSecuritySummaryMaliciousParams{
		Arc:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousParamsArc{cloudflare.RadarEmailSecuritySummaryMaliciousParamsArcPass, cloudflare.RadarEmailSecuritySummaryMaliciousParamsArcNone, cloudflare.RadarEmailSecuritySummaryMaliciousParamsArcFail}),
		Asn:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousParamsDateRange{cloudflare.RadarEmailSecuritySummaryMaliciousParamsDateRange1d, cloudflare.RadarEmailSecuritySummaryMaliciousParamsDateRange2d, cloudflare.RadarEmailSecuritySummaryMaliciousParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousParamsDKIM{cloudflare.RadarEmailSecuritySummaryMaliciousParamsDKIMPass, cloudflare.RadarEmailSecuritySummaryMaliciousParamsDKIMNone, cloudflare.RadarEmailSecuritySummaryMaliciousParamsDKIMFail}),
		Dmarc:     cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousParamsDmarc{cloudflare.RadarEmailSecuritySummaryMaliciousParamsDmarcPass, cloudflare.RadarEmailSecuritySummaryMaliciousParamsDmarcNone, cloudflare.RadarEmailSecuritySummaryMaliciousParamsDmarcFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailSecuritySummaryMaliciousParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousParamsSPF{cloudflare.RadarEmailSecuritySummaryMaliciousParamsSPFPass, cloudflare.RadarEmailSecuritySummaryMaliciousParamsSPFNone, cloudflare.RadarEmailSecuritySummaryMaliciousParamsSPFFail}),
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Email.Security.Summary.Spam(context.TODO(), cloudflare.RadarEmailSecuritySummarySpamParams{
		Arc:       cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamParamsArc{cloudflare.RadarEmailSecuritySummarySpamParamsArcPass, cloudflare.RadarEmailSecuritySummarySpamParamsArcNone, cloudflare.RadarEmailSecuritySummarySpamParamsArcFail}),
		Asn:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamParamsDateRange{cloudflare.RadarEmailSecuritySummarySpamParamsDateRange1d, cloudflare.RadarEmailSecuritySummarySpamParamsDateRange2d, cloudflare.RadarEmailSecuritySummarySpamParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamParamsDKIM{cloudflare.RadarEmailSecuritySummarySpamParamsDKIMPass, cloudflare.RadarEmailSecuritySummarySpamParamsDKIMNone, cloudflare.RadarEmailSecuritySummarySpamParamsDKIMFail}),
		Dmarc:     cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamParamsDmarc{cloudflare.RadarEmailSecuritySummarySpamParamsDmarcPass, cloudflare.RadarEmailSecuritySummarySpamParamsDmarcNone, cloudflare.RadarEmailSecuritySummarySpamParamsDmarcFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailSecuritySummarySpamParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamParamsSPF{cloudflare.RadarEmailSecuritySummarySpamParamsSPFPass, cloudflare.RadarEmailSecuritySummarySpamParamsSPFNone, cloudflare.RadarEmailSecuritySummarySpamParamsSPFFail}),
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Email.Security.Summary.SPF(context.TODO(), cloudflare.RadarEmailSecuritySummarySPFParams{
		Arc:       cloudflare.F([]cloudflare.RadarEmailSecuritySummarySPFParamsArc{cloudflare.RadarEmailSecuritySummarySPFParamsArcPass, cloudflare.RadarEmailSecuritySummarySPFParamsArcNone, cloudflare.RadarEmailSecuritySummarySPFParamsArcFail}),
		Asn:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailSecuritySummarySPFParamsDateRange{cloudflare.RadarEmailSecuritySummarySPFParamsDateRange1d, cloudflare.RadarEmailSecuritySummarySPFParamsDateRange2d, cloudflare.RadarEmailSecuritySummarySPFParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailSecuritySummarySPFParamsDKIM{cloudflare.RadarEmailSecuritySummarySPFParamsDKIMPass, cloudflare.RadarEmailSecuritySummarySPFParamsDKIMNone, cloudflare.RadarEmailSecuritySummarySPFParamsDKIMFail}),
		Dmarc:     cloudflare.F([]cloudflare.RadarEmailSecuritySummarySPFParamsDmarc{cloudflare.RadarEmailSecuritySummarySPFParamsDmarcPass, cloudflare.RadarEmailSecuritySummarySPFParamsDmarcNone, cloudflare.RadarEmailSecuritySummarySPFParamsDmarcFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailSecuritySummarySPFParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Email.Security.Summary.ThreatCategory(context.TODO(), cloudflare.RadarEmailSecuritySummaryThreatCategoryParams{
		Arc:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsArc{cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsArcPass, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsArcNone, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsArcFail}),
		Asn:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDateRange{cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDateRange1d, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDateRange2d, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDKIM{cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDKIMPass, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDKIMNone, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDKIMFail}),
		Dmarc:     cloudflare.F([]cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDmarc{cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDmarcPass, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDmarcNone, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsDmarcFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsSPF{cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsSPFPass, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsSPFNone, cloudflare.RadarEmailSecuritySummaryThreatCategoryParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
