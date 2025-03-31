// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zaraz_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/cloudflare/cloudflare-go/v4/zaraz"
)

func TestConfigUpdateWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Zaraz.Config.Update(context.TODO(), zaraz.ConfigUpdateParams{
		ZoneID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		DataLayer: cloudflare.F(true),
		DebugKey:  cloudflare.F("debugKey"),
		Settings: cloudflare.F(zaraz.ConfigUpdateParamsSettings{
			AutoInjectScript: cloudflare.F(true),
			ContextEnricher: cloudflare.F(zaraz.ConfigUpdateParamsSettingsContextEnricher{
				EscapedWorkerName: cloudflare.F("escapedWorkerName"),
				WorkerTag:         cloudflare.F("workerTag"),
			}),
			CookieDomain:        cloudflare.F("cookieDomain"),
			Ecommerce:           cloudflare.F(true),
			EventsAPIPath:       cloudflare.F("eventsApiPath"),
			HideExternalReferer: cloudflare.F(true),
			HideIPAddress:       cloudflare.F(true),
			HideQueryParams:     cloudflare.F(true),
			HideUserAgent:       cloudflare.F(true),
			InitPath:            cloudflare.F("initPath"),
			InjectIframes:       cloudflare.F(true),
			McRootPath:          cloudflare.F("mcRootPath"),
			ScriptPath:          cloudflare.F("scriptPath"),
			TrackPath:           cloudflare.F("trackPath"),
		}),
		Tools: cloudflare.F(map[string]zaraz.ConfigUpdateParamsToolsUnion{
			"foo": zaraz.ConfigUpdateParamsToolsZarazManagedComponent{
				BlockingTriggers: cloudflare.F([]string{"string"}),
				Component:        cloudflare.F("component"),
				DefaultFields: cloudflare.F(map[string]zaraz.ConfigUpdateParamsToolsZarazManagedComponentDefaultFieldsUnion{
					"foo": shared.UnionString("string"),
				}),
				Enabled:     cloudflare.F(true),
				Name:        cloudflare.F("name"),
				Permissions: cloudflare.F([]string{"string"}),
				Settings: cloudflare.F(map[string]zaraz.ConfigUpdateParamsToolsZarazManagedComponentSettingsUnion{
					"foo": shared.UnionString("string"),
				}),
				Type:            cloudflare.F(zaraz.ConfigUpdateParamsToolsZarazManagedComponentTypeComponent),
				Actions:         cloudflare.F[any](map[string]interface{}{}),
				DefaultPurpose:  cloudflare.F("defaultPurpose"),
				NeoEvents:       cloudflare.F([]interface{}{map[string]interface{}{}}),
				VendorName:      cloudflare.F("vendorName"),
				VendorPolicyURL: cloudflare.F("vendorPolicyUrl"),
			},
		}),
		Triggers: cloudflare.F(map[string]zaraz.ConfigUpdateParamsTriggers{
			"foo": {
				ExcludeRules: cloudflare.F([]zaraz.ConfigUpdateParamsTriggersExcludeRuleUnion{zaraz.ConfigUpdateParamsTriggersExcludeRulesZarazLoadRule{
					ID:    cloudflare.F("id"),
					Match: cloudflare.F("match"),
					Op:    cloudflare.F(zaraz.ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpContains),
					Value: cloudflare.F("value"),
				}}),
				LoadRules: cloudflare.F([]zaraz.ConfigUpdateParamsTriggersLoadRuleUnion{zaraz.ConfigUpdateParamsTriggersLoadRulesZarazLoadRule{
					ID:    cloudflare.F("id"),
					Match: cloudflare.F("match"),
					Op:    cloudflare.F(zaraz.ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpContains),
					Value: cloudflare.F("value"),
				}}),
				Name:        cloudflare.F("name"),
				Description: cloudflare.F("description"),
				System:      cloudflare.F(zaraz.ConfigUpdateParamsTriggersSystemPageload),
			},
		}),
		Variables: cloudflare.F(map[string]zaraz.ConfigUpdateParamsVariablesUnion{
			"foo": zaraz.ConfigUpdateParamsVariablesZarazStringVariable{
				Name:  cloudflare.F("name"),
				Type:  cloudflare.F(zaraz.ConfigUpdateParamsVariablesZarazStringVariableTypeString),
				Value: cloudflare.F("value"),
			},
		}),
		ZarazVersion: cloudflare.F(int64(0)),
		Analytics: cloudflare.F(zaraz.ConfigUpdateParamsAnalytics{
			DefaultPurpose: cloudflare.F("defaultPurpose"),
			Enabled:        cloudflare.F(true),
			SessionExpTime: cloudflare.F(int64(60)),
		}),
		Consent: cloudflare.F(zaraz.ConfigUpdateParamsConsent{
			Enabled:                cloudflare.F(true),
			ButtonTextTranslations: cloudflare.F[any](map[string]interface{}{}),
			CompanyEmail:           cloudflare.F("companyEmail"),
			CompanyName:            cloudflare.F("companyName"),
			CompanyStreetAddress:   cloudflare.F("companyStreetAddress"),
			ConsentModalIntroHTML:  cloudflare.F("consentModalIntroHTML"),
			ConsentModalIntroHTMLWithTranslations: cloudflare.F(map[string]string{
				"foo": "string",
			}),
			CookieName:                     cloudflare.F("cookieName"),
			CustomCSS:                      cloudflare.F("customCSS"),
			CustomIntroDisclaimerDismissed: cloudflare.F(true),
			DefaultLanguage:                cloudflare.F("defaultLanguage"),
			HideModal:                      cloudflare.F(true),
			Purposes: cloudflare.F(map[string]zaraz.ConfigUpdateParamsConsentPurposes{
				"foo": {
					Description: cloudflare.F("description"),
					Name:        cloudflare.F("name"),
				},
			}),
			PurposesWithTranslations: cloudflare.F(map[string]zaraz.ConfigUpdateParamsConsentPurposesWithTranslations{
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
			TcfCompliant: cloudflare.F(true),
		}),
		HistoryChange: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigGet(t *testing.T) {
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
	)
	_, err := client.Zaraz.Config.Get(context.TODO(), zaraz.ConfigGetParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
