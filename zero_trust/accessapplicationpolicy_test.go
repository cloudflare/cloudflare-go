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

func TestAccessApplicationPolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.Policies.New(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyNewParams{
			Decision: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsDecisionAllow),
			Include: cloudflare.F([]zero_trust.AccessApplicationPolicyNewParamsIncludeUnion{zero_trust.AccessApplicationPolicyNewParamsIncludeAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.AccessApplicationPolicyNewParamsIncludeAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.AccessApplicationPolicyNewParamsIncludeAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			Name:      cloudflare.F("Allow devs"),
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			ApprovalGroups: cloudflare.F([]zero_trust.AccessApplicationPolicyNewParamsApprovalGroup{{
				ApprovalsNeeded: cloudflare.F(1.000000),
				EmailAddresses:  cloudflare.F([]string{"test1@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUUID:   cloudflare.F("string"),
			}, {
				ApprovalsNeeded: cloudflare.F(3.000000),
				EmailAddresses:  cloudflare.F([]string{"test@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUUID:   cloudflare.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: cloudflare.F(true),
			Exclude: cloudflare.F([]zero_trust.AccessApplicationPolicyNewParamsExcludeUnion{zero_trust.AccessApplicationPolicyNewParamsExcludeAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.AccessApplicationPolicyNewParamsExcludeAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.AccessApplicationPolicyNewParamsExcludeAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			IsolationRequired:            cloudflare.F(false),
			Precedence:                   cloudflare.F(int64(0)),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			Require: cloudflare.F([]zero_trust.AccessApplicationPolicyNewParamsRequireUnion{zero_trust.AccessApplicationPolicyNewParamsRequireAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.AccessApplicationPolicyNewParamsRequireAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.AccessApplicationPolicyNewParamsRequireAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			SessionDuration: cloudflare.F("24h"),
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
	_, err := client.ZeroTrust.Access.Applications.Policies.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyUpdateParams{
			Decision: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsDecisionAllow),
			Include: cloudflare.F([]zero_trust.AccessApplicationPolicyUpdateParamsIncludeUnion{zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			Name:      cloudflare.F("Allow devs"),
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			ApprovalGroups: cloudflare.F([]zero_trust.AccessApplicationPolicyUpdateParamsApprovalGroup{{
				ApprovalsNeeded: cloudflare.F(1.000000),
				EmailAddresses:  cloudflare.F([]string{"test1@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUUID:   cloudflare.F("string"),
			}, {
				ApprovalsNeeded: cloudflare.F(3.000000),
				EmailAddresses:  cloudflare.F([]string{"test@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUUID:   cloudflare.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: cloudflare.F(true),
			Exclude: cloudflare.F([]zero_trust.AccessApplicationPolicyUpdateParamsExcludeUnion{zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			IsolationRequired:            cloudflare.F(false),
			Precedence:                   cloudflare.F(int64(0)),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			Require: cloudflare.F([]zero_trust.AccessApplicationPolicyUpdateParamsRequireUnion{zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			SessionDuration: cloudflare.F("24h"),
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
	_, err := client.ZeroTrust.Access.Applications.Policies.List(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyListParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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
	_, err := client.ZeroTrust.Access.Applications.Policies.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyDeleteParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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
	_, err := client.ZeroTrust.Access.Applications.Policies.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessApplicationPolicyGetParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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
