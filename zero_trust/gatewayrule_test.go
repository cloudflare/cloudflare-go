// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

func TestGatewayRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.New(context.TODO(), zero_trust.GatewayRuleNewParams{
		AccountID:     cloudflare.F("699d98642c564d2e855e9661899b7252"),
		Action:        cloudflare.F(zero_trust.GatewayRuleNewParamsActionAllow),
		Name:          cloudflare.F("block bad websites"),
		Description:   cloudflare.F("Block bad websites based on their host name."),
		DevicePosture: cloudflare.F("any(device_posture.checks.passed[*] in {\"1308749e-fcfb-4ebc-b051-fe022b632644\"})"),
		Enabled:       cloudflare.F(true),
		Filters:       cloudflare.F([]zero_trust.GatewayRuleNewParamsFilter{zero_trust.GatewayRuleNewParamsFilterHTTP}),
		Identity:      cloudflare.F("any(identity.groups.name[*] in {\"finance\"})"),
		Precedence:    cloudflare.F(int64(0)),
		RuleSettings: cloudflare.F(zero_trust.GatewayRuleNewParamsRuleSettings{
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
			AuditSSH: cloudflare.F(zero_trust.GatewayRuleNewParamsRuleSettingsAuditSSH{
				CommandLogging: cloudflare.F(false),
			}),
			BisoAdminControls: cloudflare.F(zero_trust.GatewayRuleNewParamsRuleSettingsBisoAdminControls{
				Dcp: cloudflare.F(false),
				Dd:  cloudflare.F(false),
				Dk:  cloudflare.F(false),
				Dp:  cloudflare.F(false),
				Du:  cloudflare.F(false),
			}),
			BlockPageEnabled: cloudflare.F(true),
			BlockReason:      cloudflare.F("This website is a security risk"),
			BypassParentRule: cloudflare.F(false),
			CheckSession: cloudflare.F(zero_trust.GatewayRuleNewParamsRuleSettingsCheckSession{
				Duration: cloudflare.F("300s"),
				Enforce:  cloudflare.F(true),
			}),
			DNSResolvers: cloudflare.F(zero_trust.GatewayRuleNewParamsRuleSettingsDNSResolvers{
				IPV4: cloudflare.F([]zero_trust.GatewayRuleNewParamsRuleSettingsDNSResolversIPV4{{
					IP:                         cloudflare.F("2.2.2.2"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         cloudflare.F("2.2.2.2"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         cloudflare.F("2.2.2.2"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}}),
				IPV6: cloudflare.F([]zero_trust.GatewayRuleNewParamsRuleSettingsDNSResolversIPV6{{
					IP:                         cloudflare.F("2001:DB8::"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         cloudflare.F("2001:DB8::"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         cloudflare.F("2001:DB8::"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}}),
			}),
			Egress: cloudflare.F(zero_trust.GatewayRuleNewParamsRuleSettingsEgress{
				IPV4:         cloudflare.F("192.0.2.2"),
				IPV4Fallback: cloudflare.F("192.0.2.3"),
				IPV6:         cloudflare.F("2001:DB8::/64"),
			}),
			InsecureDisableDNSSECValidation: cloudflare.F(false),
			IPCategories:                    cloudflare.F(true),
			IPIndicatorFeeds:                cloudflare.F(true),
			L4override: cloudflare.F(zero_trust.GatewayRuleNewParamsRuleSettingsL4override{
				IP:   cloudflare.F("1.1.1.1"),
				Port: cloudflare.F(int64(0)),
			}),
			NotificationSettings: cloudflare.F(zero_trust.GatewayRuleNewParamsRuleSettingsNotificationSettings{
				Enabled:    cloudflare.F(true),
				Msg:        cloudflare.F("string"),
				SupportURL: cloudflare.F("string"),
			}),
			OverrideHost: cloudflare.F("example.com"),
			OverrideIPs:  cloudflare.F([]string{"1.1.1.1", "2.2.2.2"}),
			PayloadLog: cloudflare.F(zero_trust.GatewayRuleNewParamsRuleSettingsPayloadLog{
				Enabled: cloudflare.F(true),
			}),
			ResolveDNSThroughCloudflare: cloudflare.F(true),
			UntrustedCERT: cloudflare.F(zero_trust.GatewayRuleNewParamsRuleSettingsUntrustedCERT{
				Action: cloudflare.F(zero_trust.GatewayRuleNewParamsRuleSettingsUntrustedCERTActionError),
			}),
		}),
		Schedule: cloudflare.F(zero_trust.GatewayRuleNewParamsSchedule{
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
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.GatewayRuleUpdateParams{
			AccountID:     cloudflare.F("699d98642c564d2e855e9661899b7252"),
			Action:        cloudflare.F(zero_trust.GatewayRuleUpdateParamsActionAllow),
			Name:          cloudflare.F("block bad websites"),
			Description:   cloudflare.F("Block bad websites based on their host name."),
			DevicePosture: cloudflare.F("any(device_posture.checks.passed[*] in {\"1308749e-fcfb-4ebc-b051-fe022b632644\"})"),
			Enabled:       cloudflare.F(true),
			Filters:       cloudflare.F([]zero_trust.GatewayRuleUpdateParamsFilter{zero_trust.GatewayRuleUpdateParamsFilterHTTP}),
			Identity:      cloudflare.F("any(identity.groups.name[*] in {\"finance\"})"),
			Precedence:    cloudflare.F(int64(0)),
			RuleSettings: cloudflare.F(zero_trust.GatewayRuleUpdateParamsRuleSettings{
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
				AuditSSH: cloudflare.F(zero_trust.GatewayRuleUpdateParamsRuleSettingsAuditSSH{
					CommandLogging: cloudflare.F(false),
				}),
				BisoAdminControls: cloudflare.F(zero_trust.GatewayRuleUpdateParamsRuleSettingsBisoAdminControls{
					Dcp: cloudflare.F(false),
					Dd:  cloudflare.F(false),
					Dk:  cloudflare.F(false),
					Dp:  cloudflare.F(false),
					Du:  cloudflare.F(false),
				}),
				BlockPageEnabled: cloudflare.F(true),
				BlockReason:      cloudflare.F("This website is a security risk"),
				BypassParentRule: cloudflare.F(false),
				CheckSession: cloudflare.F(zero_trust.GatewayRuleUpdateParamsRuleSettingsCheckSession{
					Duration: cloudflare.F("300s"),
					Enforce:  cloudflare.F(true),
				}),
				DNSResolvers: cloudflare.F(zero_trust.GatewayRuleUpdateParamsRuleSettingsDNSResolvers{
					IPV4: cloudflare.F([]zero_trust.GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV4{{
						IP:                         cloudflare.F("2.2.2.2"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2.2.2.2"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2.2.2.2"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
					IPV6: cloudflare.F([]zero_trust.GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV6{{
						IP:                         cloudflare.F("2001:DB8::"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
				}),
				Egress: cloudflare.F(zero_trust.GatewayRuleUpdateParamsRuleSettingsEgress{
					IPV4:         cloudflare.F("192.0.2.2"),
					IPV4Fallback: cloudflare.F("192.0.2.3"),
					IPV6:         cloudflare.F("2001:DB8::/64"),
				}),
				InsecureDisableDNSSECValidation: cloudflare.F(false),
				IPCategories:                    cloudflare.F(true),
				IPIndicatorFeeds:                cloudflare.F(true),
				L4override: cloudflare.F(zero_trust.GatewayRuleUpdateParamsRuleSettingsL4override{
					IP:   cloudflare.F("1.1.1.1"),
					Port: cloudflare.F(int64(0)),
				}),
				NotificationSettings: cloudflare.F(zero_trust.GatewayRuleUpdateParamsRuleSettingsNotificationSettings{
					Enabled:    cloudflare.F(true),
					Msg:        cloudflare.F("string"),
					SupportURL: cloudflare.F("string"),
				}),
				OverrideHost: cloudflare.F("example.com"),
				OverrideIPs:  cloudflare.F([]string{"1.1.1.1", "2.2.2.2"}),
				PayloadLog: cloudflare.F(zero_trust.GatewayRuleUpdateParamsRuleSettingsPayloadLog{
					Enabled: cloudflare.F(true),
				}),
				ResolveDNSThroughCloudflare: cloudflare.F(true),
				UntrustedCERT: cloudflare.F(zero_trust.GatewayRuleUpdateParamsRuleSettingsUntrustedCERT{
					Action: cloudflare.F(zero_trust.GatewayRuleUpdateParamsRuleSettingsUntrustedCERTActionError),
				}),
			}),
			Schedule: cloudflare.F(zero_trust.GatewayRuleUpdateParamsSchedule{
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

func TestGatewayRuleList(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.List(context.TODO(), zero_trust.GatewayRuleListParams{
		AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayRuleDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.GatewayRuleDeleteParams{
			AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
			Body:      cloudflare.F[any](map[string]interface{}{}),
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

func TestGatewayRuleGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.GatewayRuleGetParams{
			AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
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
