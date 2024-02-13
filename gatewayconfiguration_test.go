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

func TestGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfiguration(t *testing.T) {
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
	_, err := client.Gateways.Configurations.ZeroTrustAccountsGetZeroTrustAccountConfiguration(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationWithOptionalParams(t *testing.T) {
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
	_, err := client.Gateways.Configurations.ZeroTrustAccountsPatchZeroTrustAccountConfiguration(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParams{
			Settings: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettings{
				ActivityLog: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsActivityLog{
					Enabled: cloudflare.F(true),
				}),
				Antivirus: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirus{
					EnabledDownloadPhase: cloudflare.F(false),
					EnabledUploadPhase:   cloudflare.F(false),
					FailClosed:           cloudflare.F(false),
					NotificationSettings: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirusNotificationSettings{
						Enabled:    cloudflare.F(true),
						Msg:        cloudflare.F("string"),
						SupportURL: cloudflare.F("string"),
					}),
				}),
				BlockPage: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBlockPage{
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
				BodyScanning: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBodyScanning{
					InspectionMode: cloudflare.F("deep"),
				}),
				BrowserIsolation: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBrowserIsolation{
					NonIdentityEnabled:         cloudflare.F(true),
					URLBrowserIsolationEnabled: cloudflare.F(true),
				}),
				CustomCertificate: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsCustomCertificate{
					Enabled: cloudflare.F(true),
					ID:      cloudflare.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
				}),
				ExtendedEmailMatching: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsExtendedEmailMatching{
					Enabled: cloudflare.F(true),
				}),
				Fips: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsFips{
					TLS: cloudflare.F(true),
				}),
				ProtocolDetection: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsProtocolDetection{
					Enabled: cloudflare.F(true),
				}),
				TLSDecrypt: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsTLSDecrypt{
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

func TestGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationWithOptionalParams(t *testing.T) {
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
	_, err := client.Gateways.Configurations.ZeroTrustAccountsUpdateZeroTrustAccountConfiguration(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParams{
			Settings: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettings{
				ActivityLog: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsActivityLog{
					Enabled: cloudflare.F(true),
				}),
				Antivirus: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirus{
					EnabledDownloadPhase: cloudflare.F(false),
					EnabledUploadPhase:   cloudflare.F(false),
					FailClosed:           cloudflare.F(false),
					NotificationSettings: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirusNotificationSettings{
						Enabled:    cloudflare.F(true),
						Msg:        cloudflare.F("string"),
						SupportURL: cloudflare.F("string"),
					}),
				}),
				BlockPage: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBlockPage{
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
				BodyScanning: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBodyScanning{
					InspectionMode: cloudflare.F("deep"),
				}),
				BrowserIsolation: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBrowserIsolation{
					NonIdentityEnabled:         cloudflare.F(true),
					URLBrowserIsolationEnabled: cloudflare.F(true),
				}),
				CustomCertificate: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsCustomCertificate{
					Enabled: cloudflare.F(true),
					ID:      cloudflare.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
				}),
				ExtendedEmailMatching: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsExtendedEmailMatching{
					Enabled: cloudflare.F(true),
				}),
				Fips: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsFips{
					TLS: cloudflare.F(true),
				}),
				ProtocolDetection: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsProtocolDetection{
					Enabled: cloudflare.F(true),
				}),
				TLSDecrypt: cloudflare.F(cloudflare.GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsTLSDecrypt{
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
