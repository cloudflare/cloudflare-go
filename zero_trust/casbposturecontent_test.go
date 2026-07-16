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

func TestCasbPostureContentListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Content.List(context.TODO(), zero_trust.CasbPostureContentListParams{
		AccountID:         cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		Direction:         cloudflare.F(zero_trust.CasbPostureContentListParamsDirectionAsc),
		DLPProfileID:      cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IntegrationID:     cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		MaxAfflictionDate: cloudflare.F(time.Now()),
		MinAfflictionDate: cloudflare.F(time.Now()),
		Order:             cloudflare.F(zero_trust.CasbPostureContentListParamsOrderAssetName),
		Page:              cloudflare.F(int64(0)),
		PerPage:           cloudflare.F(int64(0)),
		Search:            cloudflare.F("search"),
		Vendor:            cloudflare.F(zero_trust.CasbPostureContentListParamsVendorGoogleWorkspace),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPostureContentExportWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Content.Export(context.TODO(), zero_trust.CasbPostureContentExportParams{
		AccountID: cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		DLPProfileInformation: cloudflare.F([]zero_trust.CasbPostureContentExportParamsDLPProfileInformation{{
			ID: cloudflare.F("e91a2360-da51-4fdf-9711-bcdecd462614"),
			Entries: cloudflare.F([]zero_trust.CasbPostureContentExportParamsDLPProfileInformationEntry{{
				ID:        cloudflare.F("55ba2c6c-8ef4-4b2e-9148-e75e8b6ccac1"),
				Name:      cloudflare.F("Credit Card Numbers"),
				ProfileID: cloudflare.F("e91a2360-da51-4fdf-9711-bcdecd462614"),
			}}),
			Name: cloudflare.F("Financial Information"),
		}}),
		DLPProfileID:      cloudflare.F([]string{"e91a2360-da51-4fdf-9711-bcdecd462614"}),
		IntegrationID:     cloudflare.F([]string{"c416bc38-75dc-425f-ae25-c37b5df5c37f"}),
		MaxAfflictionDate: cloudflare.F(time.Now()),
		MinAfflictionDate: cloudflare.F(time.Now()),
		Orders: cloudflare.F([]zero_trust.CasbPostureContentExportParamsOrder{{
			Direction: cloudflare.F(zero_trust.CasbPostureContentExportParamsOrdersDirectionAsc),
			Name:      cloudflare.F(zero_trust.CasbPostureContentExportParamsOrdersNameAssetName),
		}}),
		Search:  cloudflare.F("sensitive"),
		Vendors: cloudflare.F([]zero_trust.CasbPostureContentExportParamsVendor{zero_trust.CasbPostureContentExportParamsVendorGoogleWorkspace}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
