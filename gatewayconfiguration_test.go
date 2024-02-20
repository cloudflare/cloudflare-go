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

func TestGatewayConfigurationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Gateways.Configurations.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.GatewayConfigurationUpdateParams{
			Settings: cloudflare.F(cloudflare.GatewayConfigurationUpdateParamsSettings{
				ActivityLog: cloudflare.F(cloudflare.GatewayConfigurationUpdateParamsSettingsActivityLog{
					Enabled: cloudflare.F(true),
				}),
				Antivirus: cloudflare.F(cloudflare.GatewayConfigurationUpdateParamsSettingsAntivirus{
					EnabledDownloadPhase: cloudflare.F(false),
					EnabledUploadPhase:   cloudflare.F(false),
					FailClosed:           cloudflare.F(false),
					NotificationSettings: cloudflare.F(cloudflare.GatewayConfigurationUpdateParamsSettingsAntivirusNotificationSettings{
						Enabled:    cloudflare.F(true),
						Msg:        cloudflare.F("string"),
						SupportURL: cloudflare.F("string"),
					}),
				}),
				BlockPage: cloudflare.F(cloudflare.GatewayConfigurationUpdateParamsSettingsBlockPage{
					BackgroundColor: cloudflare.F("string"),
					Enabled:         cloudflare.F(true),
					FooterText:      cloudflare.F("--footer--"),
					HeaderText:      cloudflare.F("--header--"),
					LogoPath:        cloudflare.F("https://logos.com/a.png"),
					MailtoAddress:   cloudflare.F("admin@example.com"),
					MailtoSubject:   cloudflare.F("Blocked User Inquiry"),
					Name:            cloudflare.F("Cloudflare"),
					SuppressFooter:  cloudflare.F(false),
				}),
				BodyScanning: cloudflare.F(cloudflare.GatewayConfigurationUpdateParamsSettingsBodyScanning{
					InspectionMode: cloudflare.F("deep"),
				}),
				BrowserIsolation: cloudflare.F(cloudflare.GatewayConfigurationUpdateParamsSettingsBrowserIsolation{
					NonIdentityEnabled:         cloudflare.F(true),
					URLBrowserIsolationEnabled: cloudflare.F(true),
				}),
				CustomCertificate: cloudflare.F(cloudflare.GatewayConfigurationUpdateParamsSettingsCustomCertificate{
					Enabled: cloudflare.F(true),
					ID:      cloudflare.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
				}),
				ExtendedEmailMatching: cloudflare.F(cloudflare.GatewayConfigurationUpdateParamsSettingsExtendedEmailMatching{
					Enabled: cloudflare.F(true),
				}),
				Fips: cloudflare.F(cloudflare.GatewayConfigurationUpdateParamsSettingsFips{
					TLS: cloudflare.F(true),
				}),
				ProtocolDetection: cloudflare.F(cloudflare.GatewayConfigurationUpdateParamsSettingsProtocolDetection{
					Enabled: cloudflare.F(true),
				}),
				TLSDecrypt: cloudflare.F(cloudflare.GatewayConfigurationUpdateParamsSettingsTLSDecrypt{
					Enabled: cloudflare.F(true),
				}),
			}),
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

func TestGatewayConfigurationGet(t *testing.T) {
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
	_, err := client.Gateways.Configurations.Get(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayConfigurationReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Gateways.Configurations.Replace(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.GatewayConfigurationReplaceParams{
			Settings: cloudflare.F(cloudflare.GatewayConfigurationReplaceParamsSettings{
				ActivityLog: cloudflare.F(cloudflare.GatewayConfigurationReplaceParamsSettingsActivityLog{
					Enabled: cloudflare.F(true),
				}),
				Antivirus: cloudflare.F(cloudflare.GatewayConfigurationReplaceParamsSettingsAntivirus{
					EnabledDownloadPhase: cloudflare.F(false),
					EnabledUploadPhase:   cloudflare.F(false),
					FailClosed:           cloudflare.F(false),
					NotificationSettings: cloudflare.F(cloudflare.GatewayConfigurationReplaceParamsSettingsAntivirusNotificationSettings{
						Enabled:    cloudflare.F(true),
						Msg:        cloudflare.F("string"),
						SupportURL: cloudflare.F("string"),
					}),
				}),
				BlockPage: cloudflare.F(cloudflare.GatewayConfigurationReplaceParamsSettingsBlockPage{
					BackgroundColor: cloudflare.F("string"),
					Enabled:         cloudflare.F(true),
					FooterText:      cloudflare.F("--footer--"),
					HeaderText:      cloudflare.F("--header--"),
					LogoPath:        cloudflare.F("https://logos.com/a.png"),
					MailtoAddress:   cloudflare.F("admin@example.com"),
					MailtoSubject:   cloudflare.F("Blocked User Inquiry"),
					Name:            cloudflare.F("Cloudflare"),
					SuppressFooter:  cloudflare.F(false),
				}),
				BodyScanning: cloudflare.F(cloudflare.GatewayConfigurationReplaceParamsSettingsBodyScanning{
					InspectionMode: cloudflare.F("deep"),
				}),
				BrowserIsolation: cloudflare.F(cloudflare.GatewayConfigurationReplaceParamsSettingsBrowserIsolation{
					NonIdentityEnabled:         cloudflare.F(true),
					URLBrowserIsolationEnabled: cloudflare.F(true),
				}),
				CustomCertificate: cloudflare.F(cloudflare.GatewayConfigurationReplaceParamsSettingsCustomCertificate{
					Enabled: cloudflare.F(true),
					ID:      cloudflare.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
				}),
				ExtendedEmailMatching: cloudflare.F(cloudflare.GatewayConfigurationReplaceParamsSettingsExtendedEmailMatching{
					Enabled: cloudflare.F(true),
				}),
				Fips: cloudflare.F(cloudflare.GatewayConfigurationReplaceParamsSettingsFips{
					TLS: cloudflare.F(true),
				}),
				ProtocolDetection: cloudflare.F(cloudflare.GatewayConfigurationReplaceParamsSettingsProtocolDetection{
					Enabled: cloudflare.F(true),
				}),
				TLSDecrypt: cloudflare.F(cloudflare.GatewayConfigurationReplaceParamsSettingsTLSDecrypt{
					Enabled: cloudflare.F(true),
				}),
			}),
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
