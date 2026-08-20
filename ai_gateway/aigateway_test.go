// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/ai_gateway"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestAIGatewayNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.New(context.TODO(), ai_gateway.AIGatewayNewParams{
		AccountID:               cloudflare.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
		ID:                      cloudflare.F("my-gateway"),
		CacheInvalidateOnUpdate: cloudflare.F(true),
		CacheTTL:                cloudflare.F(int64(0)),
		CollectLogs:             cloudflare.F(true),
		RateLimitingInterval:    cloudflare.F(int64(0)),
		RateLimitingLimit:       cloudflare.F(int64(0)),
		Authentication:          cloudflare.F(true),
		LogManagement:           cloudflare.F(int64(10000)),
		LogManagementStrategy:   cloudflare.F(ai_gateway.AIGatewayNewParamsLogManagementStrategyStopInserting),
		Logpush:                 cloudflare.F(true),
		LogpushPublicKey:        cloudflare.F("xxxxxxxxxxxxxxxx"),
		RateLimitingTechnique:   cloudflare.F(ai_gateway.AIGatewayNewParamsRateLimitingTechniqueFixed),
		RetryBackoff:            cloudflare.F(ai_gateway.AIGatewayNewParamsRetryBackoffConstant),
		RetryDelay:              cloudflare.F(int64(0)),
		RetryMaxAttempts:        cloudflare.F(int64(1)),
		WorkersAIBillingMode:    cloudflare.F(ai_gateway.AIGatewayNewParamsWorkersAIBillingModePostpaid),
		Zdr:                     cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIGatewayUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.Update(
		context.TODO(),
		"my-gateway",
		ai_gateway.AIGatewayUpdateParams{
			AccountID:               cloudflare.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
			CacheInvalidateOnUpdate: cloudflare.F(true),
			CacheTTL:                cloudflare.F(int64(0)),
			CollectLogs:             cloudflare.F(true),
			RateLimitingInterval:    cloudflare.F(int64(0)),
			RateLimitingLimit:       cloudflare.F(int64(0)),
			Authentication:          cloudflare.F(true),
			DLP: cloudflare.F[ai_gateway.AIGatewayUpdateParamsDLPUnion](ai_gateway.AIGatewayUpdateParamsDLPObject{
				Action:   cloudflare.F(ai_gateway.AIGatewayUpdateParamsDLPObjectActionBlock),
				Enabled:  cloudflare.F(true),
				Profiles: cloudflare.F([]string{"string"}),
			}),
			Guardrails: cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrails{
				Prompt: cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPrompt{
					P1:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptP1Flag),
					S1:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptS1Flag),
					S10: cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptS10Flag),
					S11: cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptS11Flag),
					S12: cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptS12Flag),
					S13: cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptS13Flag),
					S2:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptS2Flag),
					S3:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptS3Flag),
					S4:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptS4Flag),
					S5:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptS5Flag),
					S6:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptS6Flag),
					S7:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptS7Flag),
					S8:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptS8Flag),
					S9:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsPromptS9Flag),
				}),
				Response: cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponse{
					P1:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseP1Flag),
					S1:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseS1Flag),
					S10: cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseS10Flag),
					S11: cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseS11Flag),
					S12: cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseS12Flag),
					S13: cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseS13Flag),
					S2:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseS2Flag),
					S3:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseS3Flag),
					S4:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseS4Flag),
					S5:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseS5Flag),
					S6:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseS6Flag),
					S7:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseS7Flag),
					S8:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseS8Flag),
					S9:  cloudflare.F(ai_gateway.AIGatewayUpdateParamsGuardrailsResponseS9Flag),
				}),
			}),
			LogManagement:         cloudflare.F(int64(10000)),
			LogManagementStrategy: cloudflare.F(ai_gateway.AIGatewayUpdateParamsLogManagementStrategyStopInserting),
			Logpush:               cloudflare.F(true),
			LogpushPublicKey:      cloudflare.F("xxxxxxxxxxxxxxxx"),
			Otel: cloudflare.F([]ai_gateway.AIGatewayUpdateParamsOtel{{
				Headers: cloudflare.F(map[string]string{
					"foo": "string",
				}),
				URL:           cloudflare.F("https://example.com"),
				Authorization: cloudflare.F("authorization"),
				ContentType:   cloudflare.F(ai_gateway.AIGatewayUpdateParamsOtelContentTypeJson),
			}}),
			RateLimitingTechnique: cloudflare.F(ai_gateway.AIGatewayUpdateParamsRateLimitingTechniqueFixed),
			RetryBackoff:          cloudflare.F(ai_gateway.AIGatewayUpdateParamsRetryBackoffConstant),
			RetryDelay:            cloudflare.F(int64(0)),
			RetryMaxAttempts:      cloudflare.F(int64(1)),
			SpendLimits: cloudflare.F(ai_gateway.AIGatewayUpdateParamsSpendLimits{
				Enabled: cloudflare.F(true),
				Rules: cloudflare.F([]ai_gateway.AIGatewayUpdateParamsSpendLimitsRule{{
					Limit:     cloudflare.F(1.000000),
					LimitType: cloudflare.F(ai_gateway.AIGatewayUpdateParamsSpendLimitsRulesLimitTypeCost),
					Window:    cloudflare.F(int64(1)),
					ID:        cloudflare.F("x"),
					Enabled:   cloudflare.F(true),
					Metadata: cloudflare.F(map[string]ai_gateway.AIGatewayUpdateParamsSpendLimitsRulesMetadataUnion{
						"foo": ai_gateway.AIGatewayUpdateParamsSpendLimitsRulesMetadataMode{
							Mode: cloudflare.F(ai_gateway.AIGatewayUpdateParamsSpendLimitsRulesMetadataModeModePartition),
						},
					}),
					Model: cloudflare.F(ai_gateway.AIGatewayUpdateParamsSpendLimitsRulesModel{
						Mode:   cloudflare.F(ai_gateway.AIGatewayUpdateParamsSpendLimitsRulesModelModeFilter),
						Values: cloudflare.F([]string{"string"}),
					}),
					Provider: cloudflare.F(ai_gateway.AIGatewayUpdateParamsSpendLimitsRulesProvider{
						Mode:   cloudflare.F(ai_gateway.AIGatewayUpdateParamsSpendLimitsRulesProviderModeFilter),
						Values: cloudflare.F([]string{"string"}),
					}),
					Technique: cloudflare.F(ai_gateway.AIGatewayUpdateParamsSpendLimitsRulesTechniqueFixed),
				}}),
			}),
			StoreID: cloudflare.F("store_id"),
			Stripe: cloudflare.F(ai_gateway.AIGatewayUpdateParamsStripe{
				Authorization: cloudflare.F("authorization"),
				UsageEvents: cloudflare.F([]ai_gateway.AIGatewayUpdateParamsStripeUsageEvent{{
					Payload: cloudflare.F("payload"),
				}}),
			}),
			WorkersAIBillingMode: cloudflare.F(ai_gateway.AIGatewayUpdateParamsWorkersAIBillingModePostpaid),
			Zdr:                  cloudflare.F(true),
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

func TestAIGatewayListWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.List(context.TODO(), ai_gateway.AIGatewayListParams{
		AccountID: cloudflare.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(1)),
		Search:    cloudflare.F("search"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIGatewayDelete(t *testing.T) {
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
	_, err := client.AIGateway.Delete(
		context.TODO(),
		"my-gateway",
		ai_gateway.AIGatewayDeleteParams{
			AccountID: cloudflare.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
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

func TestAIGatewayGet(t *testing.T) {
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
	_, err := client.AIGateway.Get(
		context.TODO(),
		"my-gateway",
		ai_gateway.AIGatewayGetParams{
			AccountID: cloudflare.F("3ebbcb006d4d46d7bb6a8c7f14676cb0"),
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
