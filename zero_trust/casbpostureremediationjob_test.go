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

func TestCasbPostureRemediationJobNew(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Remediations.Jobs.New(context.TODO(), zero_trust.CasbPostureRemediationJobNewParams{
		AccountID:          cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		FindingInstanceIDs: cloudflare.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		RemediationTypeID:  cloudflare.F("5a7d9e2f-1b3c-4d5e-8f6a-7b8c9d0e1f2a"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPostureRemediationJobListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Remediations.Jobs.List(context.TODO(), zero_trust.CasbPostureRemediationJobListParams{
		AccountID:        cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		Cursor:           cloudflare.F("cursor"),
		Direction:        cloudflare.F(zero_trust.CasbPostureRemediationJobListParamsDirectionAsc),
		IntegrationID:    cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		MaxUpdatedAt:     cloudflare.F(time.Now()),
		MinUpdatedAt:     cloudflare.F(time.Now()),
		Order:            cloudflare.F(zero_trust.CasbPostureRemediationJobListParamsOrderCreatedAt),
		Page:             cloudflare.F(int64(0)),
		PerPage:          cloudflare.F(int64(0)),
		Search:           cloudflare.F("search"),
		Status:           cloudflare.F(zero_trust.CasbPostureRemediationJobListParamsStatusPending),
		TriggeredByActor: cloudflare.F([]zero_trust.CasbPostureRemediationJobListParamsTriggeredByActor{zero_trust.CasbPostureRemediationJobListParamsTriggeredByActorUser}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPostureRemediationJobExportWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Remediations.Jobs.Export(context.TODO(), zero_trust.CasbPostureRemediationJobExportParams{
		AccountID:     cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		IntegrationID: cloudflare.F([]string{"55d7337e-1d0a-49fc-9826-925ba40df035"}),
		MaxUpdatedAt:  cloudflare.F(time.Now()),
		MinUpdatedAt:  cloudflare.F(time.Now()),
		Orders: cloudflare.F([]zero_trust.CasbPostureRemediationJobExportParamsOrder{{
			Direction: cloudflare.F(zero_trust.CasbPostureRemediationJobExportParamsOrdersDirectionAsc),
			Name:      cloudflare.F(zero_trust.CasbPostureRemediationJobExportParamsOrdersNameLastUpdatedAt),
		}}),
		Search: cloudflare.F("public access"),
		Status: cloudflare.F([]zero_trust.CasbPostureRemediationJobExportParamsStatus{zero_trust.CasbPostureRemediationJobExportParamsStatusPending, zero_trust.CasbPostureRemediationJobExportParamsStatusCompleted}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
