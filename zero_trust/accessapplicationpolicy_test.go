// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
)

func TestAccessApplicationPolicyNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.Policies.New(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyNewParams{
			Body: zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequest{
				ApprovalGroups: cloudflare.F([]zero_trust.ApprovalGroupParam{{
					ApprovalsNeeded: cloudflare.F(1.000000),
					EmailAddresses:  cloudflare.F([]string{"test1@cloudflare.com", "test2@cloudflare.com"}),
					EmailListUUID:   cloudflare.F("email_list_uuid"),
				}, {
					ApprovalsNeeded: cloudflare.F(3.000000),
					EmailAddresses:  cloudflare.F([]string{"test@cloudflare.com", "test2@cloudflare.com"}),
					EmailListUUID:   cloudflare.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
				}}),
				ApprovalRequired: cloudflare.F(true),
				ConnectionRules: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRules{
					RDP: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDP{
						AllowedClipboardLocalToRemoteFormats: cloudflare.F([]zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormat{zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormatText, zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormatFile}),
						AllowedClipboardRemoteToLocalFormats: cloudflare.F([]zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormat{zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormatText, zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormatFile}),
					}),
				}),
				IsolationRequired: cloudflare.F(false),
				MfaConfig: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfig{
					AllowedAuthenticators: cloudflare.F([]zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticator{zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorTotp, zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorBiometrics, zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorSecurityKey}),
					MfaDisabled:           cloudflare.F(false),
					SessionDuration:       cloudflare.F("24h"),
				}),
				Precedence:                   cloudflare.F(int64(0)),
				PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
				PurposeJustificationRequired: cloudflare.F(true),
				SessionDuration:              cloudflare.F("24h"),
			},
			AccountID: cloudflare.F("account_id"),
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

func TestAccessApplicationPolicyUpdateWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.Policies.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyUpdateParams{
			Body: zero_trust.AccessApplicationPolicyUpdateParamsBodyObject{
				Decision: cloudflare.F(zero_trust.DecisionAllow),
				Include: cloudflare.F([]zero_trust.AccessRuleUnionParam{zero_trust.EveryoneRuleParam{
					Everyone: cloudflare.F(zero_trust.EveryoneRuleEveryoneParam{}),
				}}),
				Name: cloudflare.F("Allow SSH users with a FIDO2 key"),
				ConnectionRules: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsBodyObjectConnectionRules{
					SSH: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsBodyObjectConnectionRulesSSH{
						Usernames:       cloudflare.F([]string{"root", "ubuntu"}),
						AllowEmailAlias: cloudflare.F(true),
					}),
				}),
				Exclude: cloudflare.F([]zero_trust.AccessRuleUnionParam{zero_trust.CertificateRuleParam{
					Certificate: cloudflare.F(zero_trust.CertificateRuleCertificateParam{}),
				}}),
				MfaConfig: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsBodyObjectMfaConfig{
					AllowedAuthenticators: cloudflare.F([]zero_trust.AccessApplicationPolicyUpdateParamsBodyObjectMfaConfigAllowedAuthenticator{zero_trust.AccessApplicationPolicyUpdateParamsBodyObjectMfaConfigAllowedAuthenticatorSSHFido2Key}),
					MfaDisabled:           cloudflare.F(false),
					SessionDuration:       cloudflare.F("24h"),
				}),
				Require: cloudflare.F([]zero_trust.AccessRuleUnionParam{zero_trust.CertificateRuleParam{
					Certificate: cloudflare.F(zero_trust.CertificateRuleCertificateParam{}),
				}}),
			},
			AccountID: cloudflare.F("account_id"),
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

func TestAccessApplicationPolicyListWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.Policies.List(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyListParams{
			AccountID: cloudflare.F("account_id"),
			Page:      cloudflare.F(int64(0)),
			PerPage:   cloudflare.F(int64(1000)),
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

func TestAccessApplicationPolicyDeleteWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.Policies.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyDeleteParams{
			AccountID: cloudflare.F("account_id"),
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

func TestAccessApplicationPolicyGetWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.ZeroTrust.Access.Applications.Policies.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyGetParams{
			AccountID: cloudflare.F("account_id"),
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
