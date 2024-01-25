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

func TestAccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfiguration(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Configurations.ZeroTrustAccountsGetZeroTrustAccountConfiguration(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Configurations.ZeroTrustAccountsPatchZeroTrustAccountConfiguration(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParams{
			Settings: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettings{
				ActivityLog: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsActivityLog{
					Enabled: cloudflare.F(true),
				}),
				Antivirus: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirus{
					EnabledDownloadPhase: cloudflare.F(false),
					EnabledUploadPhase:   cloudflare.F(false),
					FailClosed:           cloudflare.F(false),
				}),
				BlockPage: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBlockPage{
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
				BodyScanning: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBodyScanning{
					InspectionMode: cloudflare.F("deep"),
				}),
				BrowserIsolation: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBrowserIsolation{
					NonIdentityEnabled:         cloudflare.F(true),
					URLBrowserIsolationEnabled: cloudflare.F(true),
				}),
				CustomCertificate: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsCustomCertificate{
					Enabled: cloudflare.F(true),
					ID:      cloudflare.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
				}),
				Fips: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsFips{
					Tls: cloudflare.F(true),
				}),
				ProtocolDetection: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsProtocolDetection{
					Enabled: cloudflare.F(true),
				}),
				TlsDecrypt: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsTlsDecrypt{
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

func TestAccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Configurations.ZeroTrustAccountsUpdateZeroTrustAccountConfiguration(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParams{
			Settings: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettings{
				ActivityLog: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsActivityLog{
					Enabled: cloudflare.F(true),
				}),
				Antivirus: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirus{
					EnabledDownloadPhase: cloudflare.F(false),
					EnabledUploadPhase:   cloudflare.F(false),
					FailClosed:           cloudflare.F(false),
				}),
				BlockPage: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBlockPage{
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
				BodyScanning: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBodyScanning{
					InspectionMode: cloudflare.F("deep"),
				}),
				BrowserIsolation: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBrowserIsolation{
					NonIdentityEnabled:         cloudflare.F(true),
					URLBrowserIsolationEnabled: cloudflare.F(true),
				}),
				CustomCertificate: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsCustomCertificate{
					Enabled: cloudflare.F(true),
					ID:      cloudflare.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
				}),
				Fips: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsFips{
					Tls: cloudflare.F(true),
				}),
				ProtocolDetection: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsProtocolDetection{
					Enabled: cloudflare.F(true),
				}),
				TlsDecrypt: cloudflare.F(cloudflare.AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsTlsDecrypt{
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
