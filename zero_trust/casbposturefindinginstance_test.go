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

func TestCasbPostureFindingInstanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Findings.Instances.List(
		context.TODO(),
		"U3RhaW5sZXNzIHJvY2tz",
		zero_trust.CasbPostureFindingInstanceListParams{
			AccountID:           cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
			Archived:            cloudflare.F(true),
			AssetIDs:            cloudflare.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			Cursor:              cloudflare.F("cursor"),
			Direction:           cloudflare.F(zero_trust.CasbPostureFindingInstanceListParamsDirectionAsc),
			FindingInstanceIDs:  cloudflare.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			MaxAfflictionDate:   cloudflare.F(time.Now()),
			MinAfflictionDate:   cloudflare.F(time.Now()),
			Order:               cloudflare.F(zero_trust.CasbPostureFindingInstanceListParamsOrderAfflictionDate),
			Page:                cloudflare.F(int64(0)),
			PerPage:             cloudflare.F(int64(0)),
			RemediationStatuses: cloudflare.F([]zero_trust.CasbPostureFindingInstanceListParamsRemediationStatus{zero_trust.CasbPostureFindingInstanceListParamsRemediationStatusNone}),
			Search:              cloudflare.F("search"),
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

func TestCasbPostureFindingInstanceArchive(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Findings.Instances.Archive(
		context.TODO(),
		"U3RhaW5sZXNzIHJvY2tz",
		zero_trust.CasbPostureFindingInstanceArchiveParams{
			AccountID:      cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
			CheckInstances: cloudflare.F([]string{"497f6eca-6276-4993-bfeb-53cbbbba6f08"}),
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

func TestCasbPostureFindingInstanceExportWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Findings.Instances.Export(
		context.TODO(),
		"00000000-0000-0000-0000-000000000001-00000000-0000-0000-0000-000000000002",
		zero_trust.CasbPostureFindingInstanceExportParams{
			AccountID:         cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
			Archived:          cloudflare.F(false),
			MaxAfflictionDate: cloudflare.F(time.Now()),
			MinAfflictionDate: cloudflare.F(time.Now()),
			Orders: cloudflare.F([]zero_trust.CasbPostureFindingInstanceExportParamsOrder{{
				Direction: cloudflare.F(zero_trust.CasbPostureFindingInstanceExportParamsOrdersDirectionAsc),
				Name:      cloudflare.F(zero_trust.CasbPostureFindingInstanceExportParamsOrdersNameAssetName),
			}}),
			Search: cloudflare.F("sensitive data"),
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

func TestCasbPostureFindingInstanceGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Findings.Instances.Get(
		context.TODO(),
		"U3RhaW5sZXNzIHJvY2tz",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.CasbPostureFindingInstanceGetParams{
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

func TestCasbPostureFindingInstanceUnarchive(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Findings.Instances.Unarchive(
		context.TODO(),
		"U3RhaW5sZXNzIHJvY2tz",
		zero_trust.CasbPostureFindingInstanceUnarchiveParams{
			AccountID:      cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
			CheckInstances: cloudflare.F([]string{"497f6eca-6276-4993-bfeb-53cbbbba6f08"}),
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
