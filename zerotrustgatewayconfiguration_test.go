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

func TestZeroTrustGatewayConfigurationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Configurations.Update(context.TODO(), cloudflare.ZeroTrustGatewayConfigurationUpdateParams{
		AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
		Settings: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationUpdateParamsSettings{
			ActivityLog: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationUpdateParamsSettingsActivityLog{
				Enabled: cloudflare.F(true),
			}),
			Antivirus: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationUpdateParamsSettingsAntivirus{
				EnabledDownloadPhase: cloudflare.F(false),
				EnabledUploadPhase:   cloudflare.F(false),
				FailClosed:           cloudflare.F(false),
				NotificationSettings: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationUpdateParamsSettingsAntivirusNotificationSettings{
					Enabled:    cloudflare.F(true),
					Msg:        cloudflare.F("string"),
					SupportURL: cloudflare.F("string"),
				}),
			}),
			BlockPage: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationUpdateParamsSettingsBlockPage{
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
			BodyScanning: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationUpdateParamsSettingsBodyScanning{
				InspectionMode: cloudflare.F("deep"),
			}),
			BrowserIsolation: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationUpdateParamsSettingsBrowserIsolation{
				NonIdentityEnabled:         cloudflare.F(true),
				URLBrowserIsolationEnabled: cloudflare.F(true),
			}),
			CustomCertificate: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationUpdateParamsSettingsCustomCertificate{
				Enabled: cloudflare.F(true),
				ID:      cloudflare.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
			}),
			ExtendedEmailMatching: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationUpdateParamsSettingsExtendedEmailMatching{
				Enabled: cloudflare.F(true),
			}),
			Fips: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationUpdateParamsSettingsFips{
				TLS: cloudflare.F(true),
			}),
			ProtocolDetection: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationUpdateParamsSettingsProtocolDetection{
				Enabled: cloudflare.F(true),
			}),
			TLSDecrypt: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationUpdateParamsSettingsTLSDecrypt{
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

func TestZeroTrustGatewayConfigurationEditWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Configurations.Edit(context.TODO(), cloudflare.ZeroTrustGatewayConfigurationEditParams{
		AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
		Settings: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationEditParamsSettings{
			ActivityLog: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationEditParamsSettingsActivityLog{
				Enabled: cloudflare.F(true),
			}),
			Antivirus: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationEditParamsSettingsAntivirus{
				EnabledDownloadPhase: cloudflare.F(false),
				EnabledUploadPhase:   cloudflare.F(false),
				FailClosed:           cloudflare.F(false),
				NotificationSettings: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationEditParamsSettingsAntivirusNotificationSettings{
					Enabled:    cloudflare.F(true),
					Msg:        cloudflare.F("string"),
					SupportURL: cloudflare.F("string"),
				}),
			}),
			BlockPage: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationEditParamsSettingsBlockPage{
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
			BodyScanning: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationEditParamsSettingsBodyScanning{
				InspectionMode: cloudflare.F("deep"),
			}),
			BrowserIsolation: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationEditParamsSettingsBrowserIsolation{
				NonIdentityEnabled:         cloudflare.F(true),
				URLBrowserIsolationEnabled: cloudflare.F(true),
			}),
			CustomCertificate: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationEditParamsSettingsCustomCertificate{
				Enabled: cloudflare.F(true),
				ID:      cloudflare.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
			}),
			ExtendedEmailMatching: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationEditParamsSettingsExtendedEmailMatching{
				Enabled: cloudflare.F(true),
			}),
			Fips: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationEditParamsSettingsFips{
				TLS: cloudflare.F(true),
			}),
			ProtocolDetection: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationEditParamsSettingsProtocolDetection{
				Enabled: cloudflare.F(true),
			}),
			TLSDecrypt: cloudflare.F(cloudflare.ZeroTrustGatewayConfigurationEditParamsSettingsTLSDecrypt{
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

func TestZeroTrustGatewayConfigurationGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Configurations.Get(context.TODO(), cloudflare.ZeroTrustGatewayConfigurationGetParams{
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
