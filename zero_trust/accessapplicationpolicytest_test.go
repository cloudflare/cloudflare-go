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

func TestAccessApplicationPolicyTestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.PolicyTests.New(context.TODO(), zero_trust.AccessApplicationPolicyTestNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Policies: cloudflare.F([]zero_trust.AccessApplicationPolicyTestNewParamsPolicyUnion{zero_trust.AccessApplicationPolicyTestNewParamsPoliciesObject{
			Decision: cloudflare.F(zero_trust.DecisionAllow),
			Include: cloudflare.F([]zero_trust.AccessRuleUnionParam{zero_trust.GroupRuleParam{
				Group: cloudflare.F(zero_trust.GroupRuleGroupParam{
					ID: cloudflare.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
			}}),
			Name: cloudflare.F("Allow devs"),
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
			Exclude: cloudflare.F([]zero_trust.AccessRuleUnionParam{zero_trust.GroupRuleParam{
				Group: cloudflare.F(zero_trust.GroupRuleGroupParam{
					ID: cloudflare.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
			}}),
			IsolationRequired: cloudflare.F(false),
			MfaConfig: cloudflare.F(zero_trust.AccessApplicationPolicyTestNewParamsPoliciesObjectMfaConfig{
				AllowedAuthenticators: cloudflare.F([]zero_trust.AccessApplicationPolicyTestNewParamsPoliciesObjectMfaConfigAllowedAuthenticator{zero_trust.AccessApplicationPolicyTestNewParamsPoliciesObjectMfaConfigAllowedAuthenticatorTotp, zero_trust.AccessApplicationPolicyTestNewParamsPoliciesObjectMfaConfigAllowedAuthenticatorBiometrics, zero_trust.AccessApplicationPolicyTestNewParamsPoliciesObjectMfaConfigAllowedAuthenticatorSecurityKey}),
				MfaBypass:             cloudflare.F(false),
				SessionDuration:       cloudflare.F("24h"),
			}),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			Require: cloudflare.F([]zero_trust.AccessRuleUnionParam{zero_trust.GroupRuleParam{
				Group: cloudflare.F(zero_trust.GroupRuleGroupParam{
					ID: cloudflare.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
			}}),
			SessionDuration: cloudflare.F("24h"),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessApplicationPolicyTestGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.PolicyTests.Get(
		context.TODO(),
		"f1a8b3c9d4e5f6789a0b1c2d3e4f5678a9b0c1d2e3f4a5b67890c1d2e3f4b5a6",
		zero_trust.AccessApplicationPolicyTestGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
