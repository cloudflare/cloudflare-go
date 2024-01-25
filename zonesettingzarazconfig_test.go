// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestZoneSettingZarazConfigGet(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Settings.Zaraz.Config.Get(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneSettingZarazConfigUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Settings.Zaraz.Config.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneSettingZarazConfigUpdateParams{
			Consent: cloudflare.F(cloudflare.ZoneSettingZarazConfigUpdateParamsConsent{
				ButtonTextTranslations: cloudflare.F(cloudflare.ZoneSettingZarazConfigUpdateParamsConsentButtonTextTranslations{
					AcceptAll:        cloudflare.F[any](map[string]interface{}{}),
					ConfirmMyChoices: cloudflare.F[any](map[string]interface{}{}),
					RejectAll:        cloudflare.F[any](map[string]interface{}{}),
				}),
				CompanyEmail:                          cloudflare.F("string"),
				CompanyName:                           cloudflare.F("string"),
				CompanyStreetAddress:                  cloudflare.F("string"),
				ConsentModalIntroHTML:                 cloudflare.F("string"),
				ConsentModalIntroHTMLWithTranslations: cloudflare.F[any](map[string]interface{}{}),
				CookieName:                            cloudflare.F("string"),
				CustomCss:                             cloudflare.F("string"),
				CustomIntroDisclaimerDismissed:        cloudflare.F(true),
				DefaultLanguage:                       cloudflare.F("string"),
				Enabled:                               cloudflare.F(true),
				HideModal:                             cloudflare.F(true),
				Purposes:                              cloudflare.F[any](map[string]interface{}{}),
				PurposesWithTranslations:              cloudflare.F[any](map[string]interface{}{}),
			}),
			DataLayer:     cloudflare.F(true),
			DebugKey:      cloudflare.F("string"),
			HistoryChange: cloudflare.F(true),
			Settings: cloudflare.F(cloudflare.ZoneSettingZarazConfigUpdateParamsSettings{
				AutoInjectScript: cloudflare.F(true),
				ContextEnricher: cloudflare.F(cloudflare.ZoneSettingZarazConfigUpdateParamsSettingsContextEnricher{
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
			Tools:        cloudflare.F[any](map[string]interface{}{}),
			Triggers:     cloudflare.F[any](map[string]interface{}{}),
			Variables:    cloudflare.F[any](map[string]interface{}{}),
			ZarazVersion: cloudflare.F(int64(0)),
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
