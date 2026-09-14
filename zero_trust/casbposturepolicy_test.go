// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
)

func TestCasbPosturePolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Policies.New(context.TODO(), zero_trust.CasbPosturePolicyNewParams{
		AccountID: cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		Actions: cloudflare.F(zero_trust.CasbPosturePolicyNewParamsActions{
			RemediationTypes: cloudflare.F([]zero_trust.CasbPosturePolicyNewParamsActionsRemediationType{{
				RemediationTypeID: cloudflare.F("5a7d9e2f-1b3c-4d5e-8f6a-7b8c9d0e1f2a"),
			}}),
			WebhookConfigs: cloudflare.F([]zero_trust.CasbPosturePolicyNewParamsActionsWebhookConfig{{
				WebhookConfigID: cloudflare.F("3f7b8c9d-6e5a-4f3b-9c2d-1e0a8b7c6d5e"),
			}}),
		}),
		AppliesToAllIntegrations: cloudflare.F(false),
		DisplayName:              cloudflare.F("Auto-remediate public files"),
		Enabled:                  cloudflare.F(true),
		FindingTypeID:            cloudflare.F("5a7d9e2f-1b3c-4d5e-8f6a-7b8c9d0e1f2a"),
		Description:              cloudflare.F("Automatically remove public access from files when detected"),
		IntegrationIDs:           cloudflare.F([]string{"497f6eca-6276-4993-bfeb-53cbbbba6f08"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPosturePolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Policies.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.CasbPosturePolicyUpdateParams{
			AccountID: cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
			Actions: cloudflare.F(zero_trust.CasbPosturePolicyUpdateParamsActions{
				RemediationTypes: cloudflare.F([]zero_trust.CasbPosturePolicyUpdateParamsActionsRemediationType{{
					RemediationTypeID: cloudflare.F("5a7d9e2f-1b3c-4d5e-8f6a-7b8c9d0e1f2a"),
				}}),
				WebhookConfigs: cloudflare.F([]zero_trust.CasbPosturePolicyUpdateParamsActionsWebhookConfig{{
					WebhookConfigID: cloudflare.F("3f7b8c9d-6e5a-4f3b-9c2d-1e0a8b7c6d5e"),
				}}),
			}),
			AppliesToAllIntegrations: cloudflare.F(false),
			DisplayName:              cloudflare.F("Auto-remediate public files"),
			Enabled:                  cloudflare.F(true),
			Description:              cloudflare.F("Automatically remove public access from files when detected"),
			IntegrationIDs:           cloudflare.F([]string{"497f6eca-6276-4993-bfeb-53cbbbba6f08"}),
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

func TestCasbPosturePolicyListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Policies.List(context.TODO(), zero_trust.CasbPosturePolicyListParams{
		AccountID: cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		Cursor:    cloudflare.F("cursor"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPosturePolicyDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Policies.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.CasbPosturePolicyDeleteParams{
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

func TestCasbPosturePolicyGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Policies.Get(
		context.TODO(),
		"497f6eca-6276-4993-bfeb-53cbbbba6f08",
		zero_trust.CasbPosturePolicyGetParams{
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
