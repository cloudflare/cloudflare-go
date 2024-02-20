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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Access.Applications.Policies.New(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccessApplicationPolicyNewParams{
			Decision: cloudflare.F(cloudflare.AccessApplicationPolicyNewParamsDecisionAllow),
			Include: cloudflare.F([]cloudflare.AccessApplicationPolicyNewParamsInclude{cloudflare.AccessApplicationPolicyNewParamsIncludeAccessEmailRule(cloudflare.AccessApplicationPolicyNewParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessApplicationPolicyNewParamsIncludeAccessEmailRule(cloudflare.AccessApplicationPolicyNewParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessApplicationPolicyNewParamsIncludeAccessEmailRule(cloudflare.AccessApplicationPolicyNewParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			ApprovalGroups: cloudflare.F([]cloudflare.AccessApplicationPolicyNewParamsApprovalGroup{{
				ApprovalsNeeded: cloudflare.F(1.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test1@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cloudflare.F("string"),
			}, {
				ApprovalsNeeded: cloudflare.F(3.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cloudflare.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: cloudflare.F(true),
			Exclude: cloudflare.F([]cloudflare.AccessApplicationPolicyNewParamsExclude{cloudflare.AccessApplicationPolicyNewParamsExcludeAccessEmailRule(cloudflare.AccessApplicationPolicyNewParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessApplicationPolicyNewParamsExcludeAccessEmailRule(cloudflare.AccessApplicationPolicyNewParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessApplicationPolicyNewParamsExcludeAccessEmailRule(cloudflare.AccessApplicationPolicyNewParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			IsolationRequired:            cloudflare.F(false),
			Precedence:                   cloudflare.F(int64(0)),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			Require: cloudflare.F([]cloudflare.AccessApplicationPolicyNewParamsRequire{cloudflare.AccessApplicationPolicyNewParamsRequireAccessEmailRule(cloudflare.AccessApplicationPolicyNewParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessApplicationPolicyNewParamsRequireAccessEmailRule(cloudflare.AccessApplicationPolicyNewParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessApplicationPolicyNewParamsRequireAccessEmailRule(cloudflare.AccessApplicationPolicyNewParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail{
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

func TestAccessApplicationPolicyList(t *testing.T) {
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
	_, err := client.Access.Applications.Policies.List(
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

func TestAccessApplicationPolicyDelete(t *testing.T) {
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
	_, err := client.Access.Applications.Policies.Delete(
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

func TestAccessApplicationPolicyGet(t *testing.T) {
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
	_, err := client.Access.Applications.Policies.Get(
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

func TestAccessApplicationPolicyReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Access.Applications.Policies.Replace(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccessApplicationPolicyReplaceParams{
			Decision: cloudflare.F(cloudflare.AccessApplicationPolicyReplaceParamsDecisionAllow),
			Include: cloudflare.F([]cloudflare.AccessApplicationPolicyReplaceParamsInclude{cloudflare.AccessApplicationPolicyReplaceParamsIncludeAccessEmailRule(cloudflare.AccessApplicationPolicyReplaceParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyReplaceParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessApplicationPolicyReplaceParamsIncludeAccessEmailRule(cloudflare.AccessApplicationPolicyReplaceParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyReplaceParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessApplicationPolicyReplaceParamsIncludeAccessEmailRule(cloudflare.AccessApplicationPolicyReplaceParamsIncludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyReplaceParamsIncludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			Name: cloudflare.F("Allow devs"),
			ApprovalGroups: cloudflare.F([]cloudflare.AccessApplicationPolicyReplaceParamsApprovalGroup{{
				ApprovalsNeeded: cloudflare.F(1.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test1@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cloudflare.F("string"),
			}, {
				ApprovalsNeeded: cloudflare.F(3.000000),
				EmailAddresses:  cloudflare.F([]interface{}{"test@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cloudflare.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: cloudflare.F(true),
			Exclude: cloudflare.F([]cloudflare.AccessApplicationPolicyReplaceParamsExclude{cloudflare.AccessApplicationPolicyReplaceParamsExcludeAccessEmailRule(cloudflare.AccessApplicationPolicyReplaceParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyReplaceParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessApplicationPolicyReplaceParamsExcludeAccessEmailRule(cloudflare.AccessApplicationPolicyReplaceParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyReplaceParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessApplicationPolicyReplaceParamsExcludeAccessEmailRule(cloudflare.AccessApplicationPolicyReplaceParamsExcludeAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyReplaceParamsExcludeAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			})}),
			IsolationRequired:            cloudflare.F(false),
			Precedence:                   cloudflare.F(int64(0)),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			Require: cloudflare.F([]cloudflare.AccessApplicationPolicyReplaceParamsRequire{cloudflare.AccessApplicationPolicyReplaceParamsRequireAccessEmailRule(cloudflare.AccessApplicationPolicyReplaceParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyReplaceParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessApplicationPolicyReplaceParamsRequireAccessEmailRule(cloudflare.AccessApplicationPolicyReplaceParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyReplaceParamsRequireAccessEmailRuleEmail{
					Email: cloudflare.F("test@example.com"),
				}),
			}), cloudflare.AccessApplicationPolicyReplaceParamsRequireAccessEmailRule(cloudflare.AccessApplicationPolicyReplaceParamsRequireAccessEmailRule{
				Email: cloudflare.F(cloudflare.AccessApplicationPolicyReplaceParamsRequireAccessEmailRuleEmail{
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
