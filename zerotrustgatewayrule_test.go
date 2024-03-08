// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestZeroTrustGatewayRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.New(context.TODO(), cloudflare.ZeroTrustGatewayRuleNewParams{
		AccountID:     cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
		Action:        cloudflare.F(cloudflare.ZeroTrustGatewayRuleNewParamsActionAllow),
		Name:          cloudflare.F("block bad websites"),
		Description:   cloudflare.F("Block bad websites based on their host name."),
		DevicePosture: cloudflare.F("any(device_posture.checks.passed[*] in {\"1308749e-fcfb-4ebc-b051-fe022b632644\"})"),
		Enabled:       cloudflare.F(true),
		Filters:       cloudflare.F([]cloudflare.ZeroTrustGatewayRuleNewParamsFilter{cloudflare.ZeroTrustGatewayRuleNewParamsFilterHTTP}),
		Identity:      cloudflare.F("any(identity.groups.name[*] in {\"finance\"})"),
		Precedence:    cloudflare.F(int64(0)),
		RuleSettings: cloudflare.F(cloudflare.ZeroTrustGatewayRuleNewParamsRuleSettings{
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
			AuditSSH: cloudflare.F(cloudflare.ZeroTrustGatewayRuleNewParamsRuleSettingsAuditSSH{
				CommandLogging: cloudflare.F(false),
			}),
			BisoAdminControls: cloudflare.F(cloudflare.ZeroTrustGatewayRuleNewParamsRuleSettingsBisoAdminControls{
				Dcp: cloudflare.F(false),
				Dd:  cloudflare.F(false),
				Dk:  cloudflare.F(false),
				Dp:  cloudflare.F(false),
				Du:  cloudflare.F(false),
			}),
			BlockPageEnabled: cloudflare.F(true),
			BlockReason:      cloudflare.F("This website is a security risk"),
			BypassParentRule: cloudflare.F(false),
			CheckSession: cloudflare.F(cloudflare.ZeroTrustGatewayRuleNewParamsRuleSettingsCheckSession{
				Duration: cloudflare.F("300s"),
				Enforce:  cloudflare.F(true),
			}),
			DNSResolvers: cloudflare.F(cloudflare.ZeroTrustGatewayRuleNewParamsRuleSettingsDNSResolvers{
				IPV4: cloudflare.F([]cloudflare.ZeroTrustGatewayRuleNewParamsRuleSettingsDNSResolversIPV4{{
					IP:                         cloudflare.F("2001:DB8::/64"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         cloudflare.F("2001:DB8::/64"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         cloudflare.F("2001:DB8::/64"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}}),
				IPV6: cloudflare.F([]cloudflare.ZeroTrustGatewayRuleNewParamsRuleSettingsDNSResolversIPV6{{
					IP:                         cloudflare.F("2001:DB8::/64"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         cloudflare.F("2001:DB8::/64"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         cloudflare.F("2001:DB8::/64"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}}),
			}),
			Egress: cloudflare.F(cloudflare.ZeroTrustGatewayRuleNewParamsRuleSettingsEgress{
				IPV4:         cloudflare.F("192.0.2.2"),
				IPV4Fallback: cloudflare.F("192.0.2.3"),
				IPV6:         cloudflare.F("2001:DB8::/64"),
			}),
			InsecureDisableDNSSECValidation: cloudflare.F(false),
			IPCategories:                    cloudflare.F(true),
			IPIndicatorFeeds:                cloudflare.F(true),
			L4override: cloudflare.F(cloudflare.ZeroTrustGatewayRuleNewParamsRuleSettingsL4override{
				IP:   cloudflare.F("1.1.1.1"),
				Port: cloudflare.F(int64(0)),
			}),
			NotificationSettings: cloudflare.F(cloudflare.ZeroTrustGatewayRuleNewParamsRuleSettingsNotificationSettings{
				Enabled:    cloudflare.F(true),
				Msg:        cloudflare.F("string"),
				SupportURL: cloudflare.F("string"),
			}),
			OverrideHost: cloudflare.F("example.com"),
			OverrideIPs:  cloudflare.F([]string{"1.1.1.1", "2.2.2.2"}),
			PayloadLog: cloudflare.F(cloudflare.ZeroTrustGatewayRuleNewParamsRuleSettingsPayloadLog{
				Enabled: cloudflare.F(true),
			}),
			ResolveDNSThroughCloudflare: cloudflare.F(true),
			UntrustedCert: cloudflare.F(cloudflare.ZeroTrustGatewayRuleNewParamsRuleSettingsUntrustedCert{
				Action: cloudflare.F(cloudflare.ZeroTrustGatewayRuleNewParamsRuleSettingsUntrustedCertActionError),
			}),
		}),
		Schedule: cloudflare.F(cloudflare.ZeroTrustGatewayRuleNewParamsSchedule{
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

func TestZeroTrustGatewayRuleUpdateWithOptionalParams(t *testing.T) {
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
		cloudflare.ZeroTrustGatewayRuleUpdateParams{
			AccountID:     cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
			Action:        cloudflare.F(cloudflare.ZeroTrustGatewayRuleUpdateParamsActionAllow),
			Name:          cloudflare.F("block bad websites"),
			Description:   cloudflare.F("Block bad websites based on their host name."),
			DevicePosture: cloudflare.F("any(device_posture.checks.passed[*] in {\"1308749e-fcfb-4ebc-b051-fe022b632644\"})"),
			Enabled:       cloudflare.F(true),
			Filters:       cloudflare.F([]cloudflare.ZeroTrustGatewayRuleUpdateParamsFilter{cloudflare.ZeroTrustGatewayRuleUpdateParamsFilterHTTP}),
			Identity:      cloudflare.F("any(identity.groups.name[*] in {\"finance\"})"),
			Precedence:    cloudflare.F(int64(0)),
			RuleSettings: cloudflare.F(cloudflare.ZeroTrustGatewayRuleUpdateParamsRuleSettings{
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
				AuditSSH: cloudflare.F(cloudflare.ZeroTrustGatewayRuleUpdateParamsRuleSettingsAuditSSH{
					CommandLogging: cloudflare.F(false),
				}),
				BisoAdminControls: cloudflare.F(cloudflare.ZeroTrustGatewayRuleUpdateParamsRuleSettingsBisoAdminControls{
					Dcp: cloudflare.F(false),
					Dd:  cloudflare.F(false),
					Dk:  cloudflare.F(false),
					Dp:  cloudflare.F(false),
					Du:  cloudflare.F(false),
				}),
				BlockPageEnabled: cloudflare.F(true),
				BlockReason:      cloudflare.F("This website is a security risk"),
				BypassParentRule: cloudflare.F(false),
				CheckSession: cloudflare.F(cloudflare.ZeroTrustGatewayRuleUpdateParamsRuleSettingsCheckSession{
					Duration: cloudflare.F("300s"),
					Enforce:  cloudflare.F(true),
				}),
				DNSResolvers: cloudflare.F(cloudflare.ZeroTrustGatewayRuleUpdateParamsRuleSettingsDNSResolvers{
					IPV4: cloudflare.F([]cloudflare.ZeroTrustGatewayRuleUpdateParamsRuleSettingsDNSResolversIPV4{{
						IP:                         cloudflare.F("2001:DB8::/64"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::/64"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::/64"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
					IPV6: cloudflare.F([]cloudflare.ZeroTrustGatewayRuleUpdateParamsRuleSettingsDNSResolversIPV6{{
						IP:                         cloudflare.F("2001:DB8::/64"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::/64"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::/64"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
				}),
				Egress: cloudflare.F(cloudflare.ZeroTrustGatewayRuleUpdateParamsRuleSettingsEgress{
					IPV4:         cloudflare.F("192.0.2.2"),
					IPV4Fallback: cloudflare.F("192.0.2.3"),
					IPV6:         cloudflare.F("2001:DB8::/64"),
				}),
				InsecureDisableDNSSECValidation: cloudflare.F(false),
				IPCategories:                    cloudflare.F(true),
				IPIndicatorFeeds:                cloudflare.F(true),
				L4override: cloudflare.F(cloudflare.ZeroTrustGatewayRuleUpdateParamsRuleSettingsL4override{
					IP:   cloudflare.F("1.1.1.1"),
					Port: cloudflare.F(int64(0)),
				}),
				NotificationSettings: cloudflare.F(cloudflare.ZeroTrustGatewayRuleUpdateParamsRuleSettingsNotificationSettings{
					Enabled:    cloudflare.F(true),
					Msg:        cloudflare.F("string"),
					SupportURL: cloudflare.F("string"),
				}),
				OverrideHost: cloudflare.F("example.com"),
				OverrideIPs:  cloudflare.F([]string{"1.1.1.1", "2.2.2.2"}),
				PayloadLog: cloudflare.F(cloudflare.ZeroTrustGatewayRuleUpdateParamsRuleSettingsPayloadLog{
					Enabled: cloudflare.F(true),
				}),
				ResolveDNSThroughCloudflare: cloudflare.F(true),
				UntrustedCert: cloudflare.F(cloudflare.ZeroTrustGatewayRuleUpdateParamsRuleSettingsUntrustedCert{
					Action: cloudflare.F(cloudflare.ZeroTrustGatewayRuleUpdateParamsRuleSettingsUntrustedCertActionError),
				}),
			}),
			Schedule: cloudflare.F(cloudflare.ZeroTrustGatewayRuleUpdateParamsSchedule{
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

func TestZeroTrustGatewayRuleList(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.List(context.TODO(), cloudflare.ZeroTrustGatewayRuleListParams{
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

func TestZeroTrustGatewayRuleDelete(t *testing.T) {
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
		cloudflare.ZeroTrustGatewayRuleDeleteParams{
			AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
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

func TestZeroTrustGatewayRuleGet(t *testing.T) {
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
		cloudflare.ZeroTrustGatewayRuleGetParams{
			AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
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
