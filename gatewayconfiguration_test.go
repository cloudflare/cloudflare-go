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
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Gateways.Configurations.Update(context.TODO(), cloudflare.GatewayConfigurationUpdateParams{
		AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Gateways.Configurations.Edit(context.TODO(), cloudflare.GatewayConfigurationEditParams{
		AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
		Settings: cloudflare.F(cloudflare.GatewayConfigurationEditParamsSettings{
			ActivityLog: cloudflare.F(cloudflare.GatewayConfigurationEditParamsSettingsActivityLog{
				Enabled: cloudflare.F(true),
			}),
			Antivirus: cloudflare.F(cloudflare.GatewayConfigurationEditParamsSettingsAntivirus{
				EnabledDownloadPhase: cloudflare.F(false),
				EnabledUploadPhase:   cloudflare.F(false),
				FailClosed:           cloudflare.F(false),
				NotificationSettings: cloudflare.F(cloudflare.GatewayConfigurationEditParamsSettingsAntivirusNotificationSettings{
					Enabled:    cloudflare.F(true),
					Msg:        cloudflare.F("string"),
					SupportURL: cloudflare.F("string"),
				}),
			}),
			BlockPage: cloudflare.F(cloudflare.GatewayConfigurationEditParamsSettingsBlockPage{
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
			BodyScanning: cloudflare.F(cloudflare.GatewayConfigurationEditParamsSettingsBodyScanning{
				InspectionMode: cloudflare.F("deep"),
			}),
			BrowserIsolation: cloudflare.F(cloudflare.GatewayConfigurationEditParamsSettingsBrowserIsolation{
				NonIdentityEnabled:         cloudflare.F(true),
				URLBrowserIsolationEnabled: cloudflare.F(true),
			}),
			CustomCertificate: cloudflare.F(cloudflare.GatewayConfigurationEditParamsSettingsCustomCertificate{
				Enabled: cloudflare.F(true),
				ID:      cloudflare.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
			}),
			ExtendedEmailMatching: cloudflare.F(cloudflare.GatewayConfigurationEditParamsSettingsExtendedEmailMatching{
				Enabled: cloudflare.F(true),
			}),
			Fips: cloudflare.F(cloudflare.GatewayConfigurationEditParamsSettingsFips{
				TLS: cloudflare.F(true),
			}),
			ProtocolDetection: cloudflare.F(cloudflare.GatewayConfigurationEditParamsSettingsProtocolDetection{
				Enabled: cloudflare.F(true),
			}),
			TLSDecrypt: cloudflare.F(cloudflare.GatewayConfigurationEditParamsSettingsTLSDecrypt{
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Gateways.Configurations.Get(context.TODO(), cloudflare.GatewayConfigurationGetParams{
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
