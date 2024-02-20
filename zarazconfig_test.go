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

func TestZarazConfigReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Zaraz.Config.Replace(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZarazConfigReplaceParams{
			DataLayer: cloudflare.F(true),
			DebugKey:  cloudflare.F("my-debug-key"),
			Settings: cloudflare.F(cloudflare.ZarazConfigReplaceParamsSettings{
				AutoInjectScript: cloudflare.F(true),
				ContextEnricher: cloudflare.F(cloudflare.ZarazConfigReplaceParamsSettingsContextEnricher{
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
				InitPath:            cloudflare.F("/i"),
				InjectIframes:       cloudflare.F(true),
				McRootPath:          cloudflare.F("string"),
				ScriptPath:          cloudflare.F("string"),
				TrackPath:           cloudflare.F("string"),
			}),
			Tools: cloudflare.F(map[string]cloudflare.ZarazConfigReplaceParamsTools{
				"aJvt": cloudflare.ZarazConfigReplaceParamsToolsZarazLegacyTool(cloudflare.ZarazConfigReplaceParamsToolsZarazLegacyTool{
					BlockingTriggers: cloudflare.F([]string{"string", "string", "string"}),
					DefaultFields: cloudflare.F(map[string]cloudflare.ZarazConfigReplaceParamsToolsZarazLegacyToolDefaultFields{
						"testKey": shared.UnionString("TEST123456"),
					}),
					DefaultPurpose: cloudflare.F("string"),
					Enabled:        cloudflare.F(true),
					Name:           cloudflare.F("Facebook Pixel"),
					Library:        cloudflare.F("string"),
					NeoEvents: cloudflare.F([]cloudflare.ZarazConfigReplaceParamsToolsZarazLegacyToolNeoEvent{{
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
					Type: cloudflare.F(cloudflare.ZarazConfigReplaceParamsToolsZarazLegacyToolTypeLibrary),
				}),
			}),
			Triggers: cloudflare.F(map[string]cloudflare.ZarazConfigReplaceParamsTriggers{
				"ktBn": {
					Description: cloudflare.F("string"),
					ExcludeRules: cloudflare.F([]cloudflare.ZarazConfigReplaceParamsTriggersExcludeRule{cloudflare.ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRule(cloudflare.ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRule{
						ID:    cloudflare.F("string"),
						Match: cloudflare.F("string"),
						Op:    cloudflare.F(cloudflare.ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOpContains),
						Value: cloudflare.F("string"),
					}), cloudflare.ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRule(cloudflare.ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRule{
						ID:    cloudflare.F("string"),
						Match: cloudflare.F("string"),
						Op:    cloudflare.F(cloudflare.ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOpContains),
						Value: cloudflare.F("string"),
					}), cloudflare.ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRule(cloudflare.ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRule{
						ID:    cloudflare.F("string"),
						Match: cloudflare.F("string"),
						Op:    cloudflare.F(cloudflare.ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOpContains),
						Value: cloudflare.F("string"),
					})}),
					LoadRules: cloudflare.F([]cloudflare.ZarazConfigReplaceParamsTriggersLoadRule{cloudflare.ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRule(cloudflare.ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRule{
						ID:    cloudflare.F("string"),
						Match: cloudflare.F("string"),
						Op:    cloudflare.F(cloudflare.ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOpContains),
						Value: cloudflare.F("string"),
					}), cloudflare.ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRule(cloudflare.ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRule{
						ID:    cloudflare.F("string"),
						Match: cloudflare.F("string"),
						Op:    cloudflare.F(cloudflare.ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOpContains),
						Value: cloudflare.F("string"),
					}), cloudflare.ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRule(cloudflare.ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRule{
						ID:    cloudflare.F("string"),
						Match: cloudflare.F("string"),
						Op:    cloudflare.F(cloudflare.ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOpContains),
						Value: cloudflare.F("string"),
					})}),
					Name:   cloudflare.F("string"),
					System: cloudflare.F(cloudflare.ZarazConfigReplaceParamsTriggersSystemPageload),
				},
			}),
			Variables: cloudflare.F(map[string]cloudflare.ZarazConfigReplaceParamsVariables{
				"Autd": cloudflare.ZarazConfigReplaceParamsVariablesObject(cloudflare.ZarazConfigReplaceParamsVariablesObject{
					Name:  cloudflare.F("ip"),
					Type:  cloudflare.F(cloudflare.ZarazConfigReplaceParamsVariablesObjectTypeString),
					Value: cloudflare.F("{{ system.device.ip }}"),
				}),
			}),
			ZarazVersion: cloudflare.F(int64(43)),
			Consent: cloudflare.F(cloudflare.ZarazConfigReplaceParamsConsent{
				ButtonTextTranslations: cloudflare.F(cloudflare.ZarazConfigReplaceParamsConsentButtonTextTranslations{
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
				CookieName:                     cloudflare.F("zaraz-consent"),
				CustomCss:                      cloudflare.F("string"),
				CustomIntroDisclaimerDismissed: cloudflare.F(true),
				DefaultLanguage:                cloudflare.F("string"),
				Enabled:                        cloudflare.F(false),
				HideModal:                      cloudflare.F(true),
				Purposes: cloudflare.F(map[string]cloudflare.ZarazConfigReplaceParamsConsentPurposes{
					"foo": {
						Description: cloudflare.F("string"),
						Name:        cloudflare.F("string"),
					},
				}),
				PurposesWithTranslations: cloudflare.F(map[string]cloudflare.ZarazConfigReplaceParamsConsentPurposesWithTranslations{
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
