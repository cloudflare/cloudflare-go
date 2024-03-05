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

func TestZeroTrustAccessApplicationPolicyNewWithOptionalParams(t *testing.T) {
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
		cloudflare.ZeroTrustAccessApplicationPolicyNewParams{
			Decision: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsDecisionAllow),
			Include: cloudflare.F([]cloudflare.ZeroTrustAccessApplicationPolicyNewParamsInclude{cloudflare.ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name:      cloudflare.F("Allow devs"),
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			ApprovalGroups: cloudflare.F([]cloudflare.ZeroTrustAccessApplicationPolicyNewParamsApprovalGroup{{
				ApprovalsNeeded: cloudflare.F(1.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test1@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUUID:   cloudflare.F("string"),
			}, {
				ApprovalsNeeded: cloudflare.F(3.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUUID:   cloudflare.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: cloudflare.F(true),
			Exclude: cloudflare.F([]cloudflare.ZeroTrustAccessApplicationPolicyNewParamsExclude{cloudflare.ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			IsolationRequired:            cloudflare.F(false),
			Precedence:                   cloudflare.F(int64(0)),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			Require: cloudflare.F([]cloudflare.ZeroTrustAccessApplicationPolicyNewParamsRequire{cloudflare.ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
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

func TestZeroTrustAccessApplicationPolicyUpdateWithOptionalParams(t *testing.T) {
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
		cloudflare.ZeroTrustAccessApplicationPolicyUpdateParams{
			Decision: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsDecisionAllow),
			Include: cloudflare.F([]cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsInclude{cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name:      cloudflare.F("Allow devs"),
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			ApprovalGroups: cloudflare.F([]cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsApprovalGroup{{
				ApprovalsNeeded: cloudflare.F(1.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test1@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUUID:   cloudflare.F("string"),
			}, {
				ApprovalsNeeded: cloudflare.F(3.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUUID:   cloudflare.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: cloudflare.F(true),
			Exclude: cloudflare.F([]cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsExclude{cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			IsolationRequired:            cloudflare.F(false),
			Precedence:                   cloudflare.F(int64(0)),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			Require: cloudflare.F([]cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsRequire{cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRule(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
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

func TestZeroTrustAccessApplicationPolicyListWithOptionalParams(t *testing.T) {
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
		cloudflare.ZeroTrustAccessApplicationPolicyListParams{
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

func TestZeroTrustAccessApplicationPolicyDeleteWithOptionalParams(t *testing.T) {
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
		cloudflare.ZeroTrustAccessApplicationPolicyDeleteParams{
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

func TestZeroTrustAccessApplicationPolicyGetWithOptionalParams(t *testing.T) {
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
		cloudflare.ZeroTrustAccessApplicationPolicyGetParams{
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
