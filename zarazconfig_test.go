// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestZarazConfigGet(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Zaraz.Config.Get(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZarazConfigUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Zaraz.Config.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZarazConfigUpdateParams{
			DataLayer: cloudflare.F(true),
			DebugKey:  cloudflare.F("string"),
			Settings: cloudflare.F(cloudflare.ZarazConfigUpdateParamsSettings{
				AutoInjectScript: cloudflare.F(true),
				ContextEnricher: cloudflare.F(cloudflare.ZarazConfigUpdateParamsSettingsContextEnricher{
					EscapedWorkerName: cloudflare.F("string"),
					WorkerTag:         cloudflare.F("string"),
				}),
				CookieDomain:        cloudflare.F("string"),
				Ecommerce:           cloudflare.F(true),
				EventsAPIPath:       cloudflare.F("string"),
				HideExternalReferer: cloudflare.F(true),
				HideIPAddress:       cloudflare.F(true),
				HideQueryParams:     cloudflare.F(true),
				HideUserAgent:       cloudflare.F(true),
				InitPath:            cloudflare.F("string"),
				InjectIframes:       cloudflare.F(true),
				McRootPath:          cloudflare.F("string"),
				ScriptPath:          cloudflare.F("string"),
				TrackPath:           cloudflare.F("string"),
			}),
			Tools: cloudflare.F(map[string]cloudflare.ZarazConfigUpdateParamsTools{
				"foo": cloudflare.ZarazConfigUpdateParamsToolsZarazLegacyTool(cloudflare.ZarazConfigUpdateParamsToolsZarazLegacyTool{
					BlockingTriggers: cloudflare.F([]string{"string", "string", "string"}),
					DefaultFields: cloudflare.F(map[string]cloudflare.ZarazConfigUpdateParamsToolsZarazLegacyToolDefaultFields{
						"foo": shared.UnionString("string"),
					}),
					DefaultPurpose: cloudflare.F("string"),
					Enabled:        cloudflare.F(true),
					Name:           cloudflare.F("string"),
					Library:        cloudflare.F("string"),
					NeoEvents: cloudflare.F([]cloudflare.ZarazConfigUpdateParamsToolsZarazLegacyToolNeoEvent{{
						BlockingTriggers: cloudflare.F([]string{"string", "string", "string"}),
						Data:             cloudflare.F[any](map[string]interface{}{}),
						FiringTriggers:   cloudflare.F([]string{"string"}),
					}, {
						BlockingTriggers: cloudflare.F([]string{"string", "string", "string"}),
						Data:             cloudflare.F[any](map[string]interface{}{}),
						FiringTriggers:   cloudflare.F([]string{"string"}),
					}, {
						BlockingTriggers: cloudflare.F([]string{"string", "string", "string"}),
						Data:             cloudflare.F[any](map[string]interface{}{}),
						FiringTriggers:   cloudflare.F([]string{"string"}),
					}}),
					Type: cloudflare.F(cloudflare.ZarazConfigUpdateParamsToolsZarazLegacyToolTypeLibrary),
				}),
			}),
			Triggers: cloudflare.F(map[string]cloudflare.ZarazConfigUpdateParamsTriggers{
				"foo": {
					Description: cloudflare.F("string"),
					ExcludeRules: cloudflare.F([]cloudflare.ZarazConfigUpdateParamsTriggersExcludeRule{cloudflare.ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule(cloudflare.ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule{
						ID:    cloudflare.F("string"),
						Match: cloudflare.F("string"),
						Op:    cloudflare.F(cloudflare.ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpContains),
						Value: cloudflare.F("string"),
					}), cloudflare.ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule(cloudflare.ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule{
						ID:    cloudflare.F("string"),
						Match: cloudflare.F("string"),
						Op:    cloudflare.F(cloudflare.ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpContains),
						Value: cloudflare.F("string"),
					}), cloudflare.ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule(cloudflare.ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule{
						ID:    cloudflare.F("string"),
						Match: cloudflare.F("string"),
						Op:    cloudflare.F(cloudflare.ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpContains),
						Value: cloudflare.F("string"),
					})}),
					LoadRules: cloudflare.F([]cloudflare.ZarazConfigUpdateParamsTriggersLoadRule{cloudflare.ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule(cloudflare.ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule{
						ID:    cloudflare.F("string"),
						Match: cloudflare.F("string"),
						Op:    cloudflare.F(cloudflare.ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpContains),
						Value: cloudflare.F("string"),
					}), cloudflare.ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule(cloudflare.ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule{
						ID:    cloudflare.F("string"),
						Match: cloudflare.F("string"),
						Op:    cloudflare.F(cloudflare.ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpContains),
						Value: cloudflare.F("string"),
					}), cloudflare.ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule(cloudflare.ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule{
						ID:    cloudflare.F("string"),
						Match: cloudflare.F("string"),
						Op:    cloudflare.F(cloudflare.ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpContains),
						Value: cloudflare.F("string"),
					})}),
					Name:   cloudflare.F("string"),
					System: cloudflare.F(cloudflare.ZarazConfigUpdateParamsTriggersSystemPageload),
				},
			}),
			Variables: cloudflare.F(map[string]cloudflare.ZarazConfigUpdateParamsVariables{
				"foo": cloudflare.ZarazConfigUpdateParamsVariablesObject(cloudflare.ZarazConfigUpdateParamsVariablesObject{
					Name:  cloudflare.F("string"),
					Type:  cloudflare.F(cloudflare.ZarazConfigUpdateParamsVariablesObjectTypeString),
					Value: cloudflare.F("string"),
				}),
			}),
			ZarazVersion: cloudflare.F(int64(0)),
			Consent: cloudflare.F(cloudflare.ZarazConfigUpdateParamsConsent{
				ButtonTextTranslations: cloudflare.F(cloudflare.ZarazConfigUpdateParamsConsentButtonTextTranslations{
					AcceptAll: cloudflare.F(map[string]string{
						"foo": "string",
					}),
					ConfirmMyChoices: cloudflare.F(map[string]string{
						"foo": "string",
					}),
					RejectAll: cloudflare.F(map[string]string{
						"foo": "string",
					}),
				}),
				CompanyEmail:          cloudflare.F("string"),
				CompanyName:           cloudflare.F("string"),
				CompanyStreetAddress:  cloudflare.F("string"),
				ConsentModalIntroHTML: cloudflare.F("string"),
				ConsentModalIntroHTMLWithTranslations: cloudflare.F(map[string]string{
					"foo": "string",
				}),
				CookieName:                     cloudflare.F("string"),
				CustomCss:                      cloudflare.F("string"),
				CustomIntroDisclaimerDismissed: cloudflare.F(true),
				DefaultLanguage:                cloudflare.F("string"),
				Enabled:                        cloudflare.F(true),
				HideModal:                      cloudflare.F(true),
				Purposes: cloudflare.F(map[string]cloudflare.ZarazConfigUpdateParamsConsentPurposes{
					"foo": {
						Description: cloudflare.F("string"),
						Name:        cloudflare.F("string"),
					},
				}),
				PurposesWithTranslations: cloudflare.F(map[string]cloudflare.ZarazConfigUpdateParamsConsentPurposesWithTranslations{
					"foo": {
						Description: cloudflare.F(map[string]string{
							"foo": "string",
						}),
						Name: cloudflare.F(map[string]string{
							"foo": "string",
						}),
						Order: cloudflare.F(int64(0)),
					},
				}),
			}),
			HistoryChange: cloudflare.F(true),
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
