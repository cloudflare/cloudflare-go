// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.Policies.New(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyNewParams{
			AccountID: cloudflare.F("account_id"),
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
			ConnectionRules: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsConnectionRules{
				Rdp: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsConnectionRulesRdp{
					AllowedClipboardLocalToRemoteFormats: cloudflare.F([]zero_trust.AccessApplicationPolicyNewParamsConnectionRulesRdpAllowedClipboardLocalToRemoteFormat{zero_trust.AccessApplicationPolicyNewParamsConnectionRulesRdpAllowedClipboardLocalToRemoteFormatText}),
					AllowedClipboardRemoteToLocalFormats: cloudflare.F([]zero_trust.AccessApplicationPolicyNewParamsConnectionRulesRdpAllowedClipboardRemoteToLocalFormat{zero_trust.AccessApplicationPolicyNewParamsConnectionRulesRdpAllowedClipboardRemoteToLocalFormatText}),
				}),
			}),
			IsolationRequired: cloudflare.F(false),
			MfaConfig: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsMfaConfig{
				AllowedAuthenticators: cloudflare.F([]zero_trust.AccessApplicationPolicyNewParamsMfaConfigAllowedAuthenticator{zero_trust.AccessApplicationPolicyNewParamsMfaConfigAllowedAuthenticatorTotp, zero_trust.AccessApplicationPolicyNewParamsMfaConfigAllowedAuthenticatorBiometrics, zero_trust.AccessApplicationPolicyNewParamsMfaConfigAllowedAuthenticatorSecurityKey}),
				MfaBypass:             cloudflare.F(false),
				SessionDuration:       cloudflare.F("24h"),
			}),
			Precedence:                   cloudflare.F(int64(0)),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			SessionDuration:              cloudflare.F("24h"),
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.Policies.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyUpdateParams{
			AccountID: cloudflare.F("account_id"),
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
			ConnectionRules: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsConnectionRules{
				Rdp: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsConnectionRulesRdp{
					AllowedClipboardLocalToRemoteFormats: cloudflare.F([]zero_trust.AccessApplicationPolicyUpdateParamsConnectionRulesRdpAllowedClipboardLocalToRemoteFormat{zero_trust.AccessApplicationPolicyUpdateParamsConnectionRulesRdpAllowedClipboardLocalToRemoteFormatText}),
					AllowedClipboardRemoteToLocalFormats: cloudflare.F([]zero_trust.AccessApplicationPolicyUpdateParamsConnectionRulesRdpAllowedClipboardRemoteToLocalFormat{zero_trust.AccessApplicationPolicyUpdateParamsConnectionRulesRdpAllowedClipboardRemoteToLocalFormatText}),
				}),
			}),
			IsolationRequired: cloudflare.F(false),
			MfaConfig: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsMfaConfig{
				AllowedAuthenticators: cloudflare.F([]zero_trust.AccessApplicationPolicyUpdateParamsMfaConfigAllowedAuthenticator{zero_trust.AccessApplicationPolicyUpdateParamsMfaConfigAllowedAuthenticatorTotp, zero_trust.AccessApplicationPolicyUpdateParamsMfaConfigAllowedAuthenticatorBiometrics, zero_trust.AccessApplicationPolicyUpdateParamsMfaConfigAllowedAuthenticatorSecurityKey}),
				MfaBypass:             cloudflare.F(false),
				SessionDuration:       cloudflare.F("24h"),
			}),
			Precedence:                   cloudflare.F(int64(0)),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			SessionDuration:              cloudflare.F("24h"),
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.Policies.List(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyListParams{
			AccountID: cloudflare.F("account_id"),
			Page:      cloudflare.F(int64(0)),
			PerPage:   cloudflare.F(int64(0)),
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
