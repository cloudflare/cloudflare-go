// File generated from our OpenAPI spec by Stainless.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Gateway.Configurations.Update(context.TODO(), zero_trust.GatewayConfigurationUpdateParams{
		AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
		Settings: cloudflare.F(zero_trust.GatewayConfigurationUpdateParamsSettings{
			ActivityLog: cloudflare.F(zero_trust.GatewayConfigurationUpdateParamsSettingsActivityLog{
				Enabled: cloudflare.F(true),
			}),
			Antivirus: cloudflare.F(zero_trust.GatewayConfigurationUpdateParamsSettingsAntivirus{
				EnabledDownloadPhase: cloudflare.F(false),
				EnabledUploadPhase:   cloudflare.F(false),
				FailClosed:           cloudflare.F(false),
				NotificationSettings: cloudflare.F(zero_trust.GatewayConfigurationUpdateParamsSettingsAntivirusNotificationSettings{
					Enabled:    cloudflare.F(true),
					Msg:        cloudflare.F("string"),
					SupportURL: cloudflare.F("string"),
				}),
			}),
			BlockPage: cloudflare.F(zero_trust.GatewayConfigurationUpdateParamsSettingsBlockPage{
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
			BodyScanning: cloudflare.F(zero_trust.GatewayConfigurationUpdateParamsSettingsBodyScanning{
				InspectionMode: cloudflare.F("deep"),
			}),
			BrowserIsolation: cloudflare.F(zero_trust.GatewayConfigurationUpdateParamsSettingsBrowserIsolation{
				NonIdentityEnabled:         cloudflare.F(true),
				URLBrowserIsolationEnabled: cloudflare.F(true),
			}),
			CustomCertificate: cloudflare.F(zero_trust.GatewayConfigurationUpdateParamsSettingsCustomCertificate{
				Enabled: cloudflare.F(true),
				ID:      cloudflare.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
			}),
			ExtendedEmailMatching: cloudflare.F(zero_trust.GatewayConfigurationUpdateParamsSettingsExtendedEmailMatching{
				Enabled: cloudflare.F(true),
			}),
			Fips: cloudflare.F(zero_trust.GatewayConfigurationUpdateParamsSettingsFips{
				TLS: cloudflare.F(true),
			}),
			ProtocolDetection: cloudflare.F(zero_trust.GatewayConfigurationUpdateParamsSettingsProtocolDetection{
				Enabled: cloudflare.F(true),
			}),
			TLSDecrypt: cloudflare.F(zero_trust.GatewayConfigurationUpdateParamsSettingsTLSDecrypt{
				Enabled: cloudflare.F(true),
			}),
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayConfigurationEditWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Gateway.Configurations.Edit(context.TODO(), zero_trust.GatewayConfigurationEditParams{
		AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
		Settings: cloudflare.F(zero_trust.GatewayConfigurationEditParamsSettings{
			ActivityLog: cloudflare.F(zero_trust.GatewayConfigurationEditParamsSettingsActivityLog{
				Enabled: cloudflare.F(true),
			}),
			Antivirus: cloudflare.F(zero_trust.GatewayConfigurationEditParamsSettingsAntivirus{
				EnabledDownloadPhase: cloudflare.F(false),
				EnabledUploadPhase:   cloudflare.F(false),
				FailClosed:           cloudflare.F(false),
				NotificationSettings: cloudflare.F(zero_trust.GatewayConfigurationEditParamsSettingsAntivirusNotificationSettings{
					Enabled:    cloudflare.F(true),
					Msg:        cloudflare.F("string"),
					SupportURL: cloudflare.F("string"),
				}),
			}),
			BlockPage: cloudflare.F(zero_trust.GatewayConfigurationEditParamsSettingsBlockPage{
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
			BodyScanning: cloudflare.F(zero_trust.GatewayConfigurationEditParamsSettingsBodyScanning{
				InspectionMode: cloudflare.F("deep"),
			}),
			BrowserIsolation: cloudflare.F(zero_trust.GatewayConfigurationEditParamsSettingsBrowserIsolation{
				NonIdentityEnabled:         cloudflare.F(true),
				URLBrowserIsolationEnabled: cloudflare.F(true),
			}),
			CustomCertificate: cloudflare.F(zero_trust.GatewayConfigurationEditParamsSettingsCustomCertificate{
				Enabled: cloudflare.F(true),
				ID:      cloudflare.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
			}),
			ExtendedEmailMatching: cloudflare.F(zero_trust.GatewayConfigurationEditParamsSettingsExtendedEmailMatching{
				Enabled: cloudflare.F(true),
			}),
			Fips: cloudflare.F(zero_trust.GatewayConfigurationEditParamsSettingsFips{
				TLS: cloudflare.F(true),
			}),
			ProtocolDetection: cloudflare.F(zero_trust.GatewayConfigurationEditParamsSettingsProtocolDetection{
				Enabled: cloudflare.F(true),
			}),
			TLSDecrypt: cloudflare.F(zero_trust.GatewayConfigurationEditParamsSettingsTLSDecrypt{
				Enabled: cloudflare.F(true),
			}),
		}),
	})
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Gateway.Configurations.Get(context.TODO(), zero_trust.GatewayConfigurationGetParams{
		AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
