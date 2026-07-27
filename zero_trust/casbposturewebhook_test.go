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

func TestCasbPostureWebhookNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Webhooks.New(context.TODO(), zero_trust.CasbPostureWebhookNewParams{
		AccountID:          cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		AuthenticationType: cloudflare.F(zero_trust.CasbPostureWebhookNewParamsAuthenticationTypeBearerAuth),
		DestinationURL:     cloudflare.F("https://example.com/webhook"),
		Label:              cloudflare.F("Send to Slack"),
		Headers: cloudflare.F([]zero_trust.CasbPostureWebhookNewParamsHeader{{
			Key:   cloudflare.F("Authorization"),
			Value: cloudflare.F("Bearer token123"),
		}, {
			Key:   cloudflare.F("X-Custom-Header"),
			Value: cloudflare.F("value"),
		}}),
		SigningSecret: cloudflare.F("my-secret-key"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPostureWebhookUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Webhooks.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.CasbPostureWebhookUpdateParams{
			AccountID:          cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
			AuthenticationType: cloudflare.F(zero_trust.CasbPostureWebhookUpdateParamsAuthenticationTypeBearerAuth),
			DestinationURL:     cloudflare.F("https://example.com/webhook"),
			Label:              cloudflare.F("Send to Slack"),
			Status:             cloudflare.F(zero_trust.CasbPostureWebhookUpdateParamsStatusEnabled),
			Headers: cloudflare.F([]zero_trust.CasbPostureWebhookUpdateParamsHeader{{
				Key:   cloudflare.F("Authorization"),
				Value: cloudflare.F("Bearer token123"),
			}, {
				Key:   cloudflare.F("X-Custom-Header"),
				Value: cloudflare.F("value"),
			}}),
			SigningSecret: cloudflare.F("my-secret-key"),
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

func TestCasbPostureWebhookList(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Webhooks.List(context.TODO(), zero_trust.CasbPostureWebhookListParams{
		AccountID: cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPostureWebhookDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Webhooks.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.CasbPostureWebhookDeleteParams{
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

func TestCasbPostureWebhookEvaluateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Webhooks.Evaluate(context.TODO(), zero_trust.CasbPostureWebhookEvaluateParams{
		AccountID:          cloudflare.F("46148281d8a93d002ef242d8b0d5f9f6"),
		AuthenticationType: cloudflare.F(zero_trust.CasbPostureWebhookEvaluateParamsAuthenticationTypeBearerAuth),
		DestinationURL:     cloudflare.F("https://example.com/webhook"),
		Headers: cloudflare.F([]zero_trust.CasbPostureWebhookEvaluateParamsHeader{{
			Key:   cloudflare.F("Authorization"),
			Value: cloudflare.F("Bearer token123"),
		}, {
			Key:   cloudflare.F("X-Custom-Header"),
			Value: cloudflare.F("value"),
		}}),
		SigningSecret: cloudflare.F("my-secret-key"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCasbPostureWebhookEvaluateExisting(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Webhooks.EvaluateExisting(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.CasbPostureWebhookEvaluateExistingParams{
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

func TestCasbPostureWebhookGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Casb.Posture.Webhooks.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.CasbPostureWebhookGetParams{
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
