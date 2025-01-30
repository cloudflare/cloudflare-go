// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_center_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/intel"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/security_center"
)

func TestInsightDismissWithOptionalParams(t *testing.T) {
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
	_, err := client.SecurityCenter.Insights.Dismiss(
		context.TODO(),
		"issue_id",
		security_center.InsightDismissParams{
			AccountID: cloudflare.F("account_id"),
			Dismiss:   cloudflare.F(true),
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

func TestInsightGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SecurityCenter.Insights.Get(context.TODO(), security_center.InsightGetParams{
		AccountID:     cloudflare.F("account_id"),
		Dismissed:     cloudflare.F(false),
		IssueClass:    cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueClassNeq: cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueType:     cloudflare.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
		IssueTypeNeq:  cloudflare.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
		Page:          cloudflare.F(int64(1)),
		PerPage:       cloudflare.F(int64(25)),
		Product:       cloudflare.F([]string{"access", "dns"}),
		ProductNeq:    cloudflare.F([]string{"access", "dns"}),
		Severity:      cloudflare.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
		SeverityNeq:   cloudflare.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
		Subject:       cloudflare.F([]string{"example.com"}),
		SubjectNeq:    cloudflare.F([]string{"example.com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
