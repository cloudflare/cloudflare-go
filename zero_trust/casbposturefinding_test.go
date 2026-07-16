// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
)

func TestCasbPostureFindingListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Casb.Posture.Findings.List(context.TODO(), zero_trust.CasbPostureFindingListParams{
		AccountID:         cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		Cursor:            cloudflare.F("cursor"),
		Direction:         cloudflare.F(zero_trust.CasbPostureFindingListParamsDirectionAsc),
		FindingTypeIDs:    cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Ignored:           cloudflare.F(true),
		IntegrationID:     cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		MaxAfflictionDate: cloudflare.F(time.Now()),
		MinAfflictionDate: cloudflare.F(time.Now()),
		Observation:       cloudflare.F(zero_trust.CasbPostureFindingListParamsObservationActivity),
		Order:             cloudflare.F(zero_trust.CasbPostureFindingListParamsOrderFindingName),
		Page:              cloudflare.F(int64(0)),
		PerPage:           cloudflare.F(int64(0)),
		Product:           cloudflare.F(zero_trust.CasbPostureFindingListParamsProductCloud),
		Search:            cloudflare.F("search"),
		Severity:          cloudflare.F(zero_trust.CasbPostureFindingListParamsSeverityCritical),
		Type:              cloudflare.F(zero_trust.CasbPostureFindingListParamsTypeContent),
		Vendor:            cloudflare.F(zero_trust.CasbPostureFindingListParamsVendorGoogleWorkspace),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPostureFindingExportWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Casb.Posture.Findings.Export(context.TODO(), zero_trust.CasbPostureFindingExportParams{
		AccountID:         cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		Ignored:           cloudflare.F(true),
		IntegrationID:     cloudflare.F([]string{"55d7337e-1d0a-49fc-9826-925ba40df035"}),
		MaxAfflictionDate: cloudflare.F(time.Now()),
		MinAfflictionDate: cloudflare.F(time.Now()),
		Orders: cloudflare.F([]zero_trust.CasbPostureFindingExportParamsOrder{{
			Direction: cloudflare.F(zero_trust.CasbPostureFindingExportParamsOrdersDirectionAsc),
			Name:      cloudflare.F(zero_trust.CasbPostureFindingExportParamsOrdersNameInstanceCount),
		}}),
		Product:    cloudflare.F(zero_trust.CasbPostureFindingExportParamsProductSaaS),
		Search:     cloudflare.F("public access"),
		Severities: cloudflare.F([]zero_trust.CasbPostureFindingExportParamsSeverity{zero_trust.CasbPostureFindingExportParamsSeverityCritical}),
		Vendors:    cloudflare.F([]zero_trust.CasbPostureFindingExportParamsVendor{zero_trust.CasbPostureFindingExportParamsVendorAws}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPostureFindingGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Casb.Posture.Findings.Get(
		context.TODO(),
		"U3RhaW5sZXNzIHJvY2tz",
		zero_trust.CasbPostureFindingGetParams{
			AccountID: cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPostureFindingIgnore(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Casb.Posture.Findings.Ignore(context.TODO(), zero_trust.CasbPostureFindingIgnoreParams{
		AccountID: cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		Checks:    cloudflare.F([]string{"MDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAxOjAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMgo="}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPostureFindingResetSeverity(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Casb.Posture.Findings.ResetSeverity(
		context.TODO(),
		"U3RhaW5sZXNzIHJvY2tz",
		zero_trust.CasbPostureFindingResetSeverityParams{
			AccountID: cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPostureFindingTuneSeverity(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Casb.Posture.Findings.TuneSeverity(
		context.TODO(),
		"U3RhaW5sZXNzIHJvY2tz",
		zero_trust.CasbPostureFindingTuneSeverityParams{
			AccountID:   cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
			NewSeverity: cloudflare.F(zero_trust.CasbPostureFindingTuneSeverityParamsNewSeverity1),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPostureFindingUnignore(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Casb.Posture.Findings.Unignore(context.TODO(), zero_trust.CasbPostureFindingUnignoreParams{
		AccountID: cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		Checks:    cloudflare.F([]string{"MDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAxOjAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMgo="}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
