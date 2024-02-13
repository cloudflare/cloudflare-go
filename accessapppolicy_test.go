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

func TestAccessAppPolicyNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Access.Apps.Policies.New(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccessAppPolicyNewParams{
			Decision: cloudflare.F(cloudflare.AccessAppPolicyNewParamsDecisionAllow),
			Include: cloudflare.F([]cloudflare.AccessAppPolicyNewParamsInclude{cloudflare.AccessAppPolicyNewParamsIncludeAccessEmailRule(cloudflare.AccessAppPolicyNewParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyNewParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyNewParamsIncludeAccessEmailRule(cloudflare.AccessAppPolicyNewParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyNewParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyNewParamsIncludeAccessEmailRule(cloudflare.AccessAppPolicyNewParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyNewParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			ApprovalGroups: cloudflare.F([]cloudflare.AccessAppPolicyNewParamsApprovalGroup{{
				ApprovalsNeeded: cloudflare.F(1.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test1@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cloudflare.F("string"),
			}, {
				ApprovalsNeeded: cloudflare.F(3.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cloudflare.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: cloudflare.F(true),
			Exclude: cloudflare.F([]cloudflare.AccessAppPolicyNewParamsExclude{cloudflare.AccessAppPolicyNewParamsExcludeAccessEmailRule(cloudflare.AccessAppPolicyNewParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyNewParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyNewParamsExcludeAccessEmailRule(cloudflare.AccessAppPolicyNewParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyNewParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyNewParamsExcludeAccessEmailRule(cloudflare.AccessAppPolicyNewParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyNewParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			IsolationRequired:            cloudflare.F(false),
			Precedence:                   cloudflare.F(int64(0)),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			Require: cloudflare.F([]cloudflare.AccessAppPolicyNewParamsRequire{cloudflare.AccessAppPolicyNewParamsRequireAccessEmailRule(cloudflare.AccessAppPolicyNewParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyNewParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyNewParamsRequireAccessEmailRule(cloudflare.AccessAppPolicyNewParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyNewParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyNewParamsRequireAccessEmailRule(cloudflare.AccessAppPolicyNewParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyNewParamsRequireAccessEmailRuleEmail{
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

func TestAccessAppPolicyUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Access.Apps.Policies.Update(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccessAppPolicyUpdateParams{
			Decision: cloudflare.F(cloudflare.AccessAppPolicyUpdateParamsDecisionAllow),
			Include: cloudflare.F([]cloudflare.AccessAppPolicyUpdateParamsInclude{cloudflare.AccessAppPolicyUpdateParamsIncludeAccessEmailRule(cloudflare.AccessAppPolicyUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyUpdateParamsIncludeAccessEmailRule(cloudflare.AccessAppPolicyUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyUpdateParamsIncludeAccessEmailRule(cloudflare.AccessAppPolicyUpdateParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyUpdateParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			ApprovalGroups: cloudflare.F([]cloudflare.AccessAppPolicyUpdateParamsApprovalGroup{{
				ApprovalsNeeded: cloudflare.F(1.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test1@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cloudflare.F("string"),
			}, {
				ApprovalsNeeded: cloudflare.F(3.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cloudflare.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: cloudflare.F(true),
			Exclude: cloudflare.F([]cloudflare.AccessAppPolicyUpdateParamsExclude{cloudflare.AccessAppPolicyUpdateParamsExcludeAccessEmailRule(cloudflare.AccessAppPolicyUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyUpdateParamsExcludeAccessEmailRule(cloudflare.AccessAppPolicyUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyUpdateParamsExcludeAccessEmailRule(cloudflare.AccessAppPolicyUpdateParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyUpdateParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			IsolationRequired:            cloudflare.F(false),
			Precedence:                   cloudflare.F(int64(0)),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			Require: cloudflare.F([]cloudflare.AccessAppPolicyUpdateParamsRequire{cloudflare.AccessAppPolicyUpdateParamsRequireAccessEmailRule(cloudflare.AccessAppPolicyUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyUpdateParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyUpdateParamsRequireAccessEmailRule(cloudflare.AccessAppPolicyUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyUpdateParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyUpdateParamsRequireAccessEmailRule(cloudflare.AccessAppPolicyUpdateParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyUpdateParamsRequireAccessEmailRuleEmail{
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

func TestAccessAppPolicyDelete(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Access.Apps.Policies.Delete(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
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

func TestAccessAppPolicyAccessPoliciesNewAnAccessPolicyWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Access.Apps.Policies.AccessPoliciesNewAnAccessPolicy(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParams{
			Decision: cloudflare.F(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsDecisionAllow),
			Include: cloudflare.F([]cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude{cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRule(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRule(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRule(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			ApprovalGroups: cloudflare.F([]cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsApprovalGroup{{
				ApprovalsNeeded: cloudflare.F(1.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test1@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cloudflare.F("string"),
			}, {
				ApprovalsNeeded: cloudflare.F(3.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cloudflare.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: cloudflare.F(true),
			Exclude: cloudflare.F([]cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude{cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRule(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRule(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRule(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			IsolationRequired:            cloudflare.F(false),
			Precedence:                   cloudflare.F(int64(0)),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			Require: cloudflare.F([]cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire{cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRule(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRule(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRule(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRuleEmail{
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

func TestAccessAppPolicyAccessPoliciesListAccessPolicies(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Access.Apps.Policies.AccessPoliciesListAccessPolicies(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
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

func TestAccessAppPolicyGet(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Access.Apps.Policies.Get(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
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
