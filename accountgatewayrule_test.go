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

func TestAccountGatewayRuleGet(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Gateways.Rules.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGatewayRuleUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Gateways.Rules.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccountGatewayRuleUpdateParams{
			Action:        cloudflare.F(cloudflare.AccountGatewayRuleUpdateParamsActionAllow),
			Name:          cloudflare.F("block bad websites"),
			Description:   cloudflare.F("Block the bad websites based on host name"),
			DevicePosture: cloudflare.F("any(device_posture.checks.passed[*] in {\"1308749e-fcfb-4ebc-b051-fe022b632644\"})"),
			Enabled:       cloudflare.F(true),
			Filters:       cloudflare.F([]cloudflare.AccountGatewayRuleUpdateParamsFilter{cloudflare.AccountGatewayRuleUpdateParamsFilterHTTP}),
			Identity:      cloudflare.F("any(identity.groups.name[*] in {\"finance\"})"),
			Precedence:    cloudflare.F(int64(0)),
			RuleSettings: cloudflare.F(cloudflare.AccountGatewayRuleUpdateParamsRuleSettings{
				AddHeaders: cloudflare.F[any](map[string]interface{}{
					"My-Next-Header": map[string]interface{}{
						"0": "foo",
						"1": "bar",
					},
					"X-Custom-Header-Name": map[string]interface{}{
						"0": "somecustomvalue",
					},
				}),
				AllowChildBypass: cloudflare.F(false),
				AuditSSH: cloudflare.F(cloudflare.AccountGatewayRuleUpdateParamsRuleSettingsAuditSSH{
					CommandLogging: cloudflare.F(false),
				}),
				BisoAdminControls: cloudflare.F(cloudflare.AccountGatewayRuleUpdateParamsRuleSettingsBisoAdminControls{
					Dcp: cloudflare.F(false),
					Dd:  cloudflare.F(false),
					Dk:  cloudflare.F(false),
					Dp:  cloudflare.F(false),
					Du:  cloudflare.F(false),
				}),
				BlockPageEnabled: cloudflare.F(true),
				BlockReason:      cloudflare.F("This website is a security risk"),
				BypassParentRule: cloudflare.F(false),
				CheckSession: cloudflare.F(cloudflare.AccountGatewayRuleUpdateParamsRuleSettingsCheckSession{
					Duration: cloudflare.F("300s"),
					Enforce:  cloudflare.F(true),
				}),
				Egress: cloudflare.F(cloudflare.AccountGatewayRuleUpdateParamsRuleSettingsEgress{
					Ipv4:         cloudflare.F("192.0.2.2"),
					Ipv4Fallback: cloudflare.F("192.0.2.3"),
					Ipv6:         cloudflare.F("2001:DB8::/64"),
				}),
				InsecureDisableDnssecValidation: cloudflare.F(false),
				IPCategories:                    cloudflare.F(true),
				L4override: cloudflare.F(cloudflare.AccountGatewayRuleUpdateParamsRuleSettingsL4override{
					IP:   cloudflare.F("1.1.1.1"),
					Port: cloudflare.F(int64(0)),
				}),
				OverrideHost: cloudflare.F("example.com"),
				OverrideIPs:  cloudflare.F([]string{"1.1.1.1", "2.2.2.2"}),
				PayloadLog: cloudflare.F(cloudflare.AccountGatewayRuleUpdateParamsRuleSettingsPayloadLog{
					Enabled: cloudflare.F(true),
				}),
				UntrustedCert: cloudflare.F(cloudflare.AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCert{
					Action: cloudflare.F(cloudflare.AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertActionError),
				}),
			}),
			Schedule: cloudflare.F(cloudflare.AccountGatewayRuleUpdateParamsSchedule{
				Fri:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Mon:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Sat:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Sun:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Thu:      cloudflare.F("08:00-12:30,13:30-17:00"),
				TimeZone: cloudflare.F("America/New York"),
				Tue:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Wed:      cloudflare.F("08:00-12:30,13:30-17:00"),
			}),
			Traffic: cloudflare.F("http.request.uri matches \".*a/partial/uri.*\" and http.request.host in $01302951-49f9-47c9-a400-0297e60b6a10"),
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

func TestAccountGatewayRuleDelete(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Gateways.Rules.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Gateways.Rules.ZeroTrustGatewayRulesNewZeroTrustGatewayRule(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParams{
			Action:        cloudflare.F(cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionAllow),
			Name:          cloudflare.F("block bad websites"),
			Description:   cloudflare.F("Block the bad websites based on host name"),
			DevicePosture: cloudflare.F("any(device_posture.checks.passed[*] in {\"1308749e-fcfb-4ebc-b051-fe022b632644\"})"),
			Enabled:       cloudflare.F(true),
			Filters:       cloudflare.F([]cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter{cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilterHTTP}),
			Identity:      cloudflare.F("any(identity.groups.name[*] in {\"finance\"})"),
			Precedence:    cloudflare.F(int64(0)),
			RuleSettings: cloudflare.F(cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettings{
				AddHeaders: cloudflare.F[any](map[string]interface{}{
					"My-Next-Header": map[string]interface{}{
						"0": "foo",
						"1": "bar",
					},
					"X-Custom-Header-Name": map[string]interface{}{
						"0": "somecustomvalue",
					},
				}),
				AllowChildBypass: cloudflare.F(false),
				AuditSSH: cloudflare.F(cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsAuditSSH{
					CommandLogging: cloudflare.F(false),
				}),
				BisoAdminControls: cloudflare.F(cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsBisoAdminControls{
					Dcp: cloudflare.F(false),
					Dd:  cloudflare.F(false),
					Dk:  cloudflare.F(false),
					Dp:  cloudflare.F(false),
					Du:  cloudflare.F(false),
				}),
				BlockPageEnabled: cloudflare.F(true),
				BlockReason:      cloudflare.F("This website is a security risk"),
				BypassParentRule: cloudflare.F(false),
				CheckSession: cloudflare.F(cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsCheckSession{
					Duration: cloudflare.F("300s"),
					Enforce:  cloudflare.F(true),
				}),
				Egress: cloudflare.F(cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsEgress{
					Ipv4:         cloudflare.F("192.0.2.2"),
					Ipv4Fallback: cloudflare.F("192.0.2.3"),
					Ipv6:         cloudflare.F("2001:DB8::/64"),
				}),
				InsecureDisableDnssecValidation: cloudflare.F(false),
				IPCategories:                    cloudflare.F(true),
				L4override: cloudflare.F(cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsL4override{
					IP:   cloudflare.F("1.1.1.1"),
					Port: cloudflare.F(int64(0)),
				}),
				OverrideHost: cloudflare.F("example.com"),
				OverrideIPs:  cloudflare.F([]string{"1.1.1.1", "2.2.2.2"}),
				PayloadLog: cloudflare.F(cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsPayloadLog{
					Enabled: cloudflare.F(true),
				}),
				UntrustedCert: cloudflare.F(cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCert{
					Action: cloudflare.F(cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertActionError),
				}),
			}),
			Schedule: cloudflare.F(cloudflare.AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsSchedule{
				Fri:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Mon:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Sat:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Sun:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Thu:      cloudflare.F("08:00-12:30,13:30-17:00"),
				TimeZone: cloudflare.F("America/New York"),
				Tue:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Wed:      cloudflare.F("08:00-12:30,13:30-17:00"),
			}),
			Traffic: cloudflare.F("http.request.uri matches \".*a/partial/uri.*\" and http.request.host in $01302951-49f9-47c9-a400-0297e60b6a10"),
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

func TestAccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRules(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Gateways.Rules.ZeroTrustGatewayRulesListZeroTrustGatewayRules(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
