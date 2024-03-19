// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/intel"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestAttackSurfaceReportIssueListWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.List(context.TODO(), intel.AttackSurfaceReportIssueListParams{
		AccountID:     cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Dismissed:     cloudflare.F(false),
		IssueClass:    cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueClassNeq: cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueType:     cloudflare.F([]intel.AttackSurfaceReportIssueListParamsIssueType{intel.AttackSurfaceReportIssueListParamsIssueTypeComplianceViolation, intel.AttackSurfaceReportIssueListParamsIssueTypeEmailSecurity}),
		IssueTypeNeq:  cloudflare.F([]intel.AttackSurfaceReportIssueListParamsIssueTypeNeq{intel.AttackSurfaceReportIssueListParamsIssueTypeNeqComplianceViolation, intel.AttackSurfaceReportIssueListParamsIssueTypeNeqEmailSecurity}),
		Page:          cloudflare.F(int64(1)),
		PerPage:       cloudflare.F(int64(25)),
		Product:       cloudflare.F([]string{"access", "dns"}),
		ProductNeq:    cloudflare.F([]string{"access", "dns"}),
		Severity:      cloudflare.F([]intel.AttackSurfaceReportIssueListParamsSeverity{intel.AttackSurfaceReportIssueListParamsSeverityLow, intel.AttackSurfaceReportIssueListParamsSeverityModerate}),
		SeverityNeq:   cloudflare.F([]intel.AttackSurfaceReportIssueListParamsSeverityNeq{intel.AttackSurfaceReportIssueListParamsSeverityNeqLow, intel.AttackSurfaceReportIssueListParamsSeverityNeqModerate}),
		Subject:       cloudflare.F([]string{"example.com", "example.com", "example.com"}),
		SubjectNeq:    cloudflare.F([]string{"example.com", "example.com", "example.com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackSurfaceReportIssueClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.Class(context.TODO(), intel.AttackSurfaceReportIssueClassParams{
		AccountID:     cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Dismissed:     cloudflare.F(false),
		IssueClass:    cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueClassNeq: cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueType:     cloudflare.F([]intel.AttackSurfaceReportIssueClassParamsIssueType{intel.AttackSurfaceReportIssueClassParamsIssueTypeComplianceViolation, intel.AttackSurfaceReportIssueClassParamsIssueTypeEmailSecurity}),
		IssueTypeNeq:  cloudflare.F([]intel.AttackSurfaceReportIssueClassParamsIssueTypeNeq{intel.AttackSurfaceReportIssueClassParamsIssueTypeNeqComplianceViolation, intel.AttackSurfaceReportIssueClassParamsIssueTypeNeqEmailSecurity}),
		Product:       cloudflare.F([]string{"access", "dns"}),
		ProductNeq:    cloudflare.F([]string{"access", "dns"}),
		Severity:      cloudflare.F([]intel.AttackSurfaceReportIssueClassParamsSeverity{intel.AttackSurfaceReportIssueClassParamsSeverityLow, intel.AttackSurfaceReportIssueClassParamsSeverityModerate}),
		SeverityNeq:   cloudflare.F([]intel.AttackSurfaceReportIssueClassParamsSeverityNeq{intel.AttackSurfaceReportIssueClassParamsSeverityNeqLow, intel.AttackSurfaceReportIssueClassParamsSeverityNeqModerate}),
		Subject:       cloudflare.F([]string{"example.com", "example.com", "example.com"}),
		SubjectNeq:    cloudflare.F([]string{"example.com", "example.com", "example.com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackSurfaceReportIssueDismissWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.Dismiss(
		context.TODO(),
		"string",
		intel.AttackSurfaceReportIssueDismissParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestAttackSurfaceReportIssueSeverityWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.Severity(context.TODO(), intel.AttackSurfaceReportIssueSeverityParams{
		AccountID:     cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Dismissed:     cloudflare.F(false),
		IssueClass:    cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueClassNeq: cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueType:     cloudflare.F([]intel.AttackSurfaceReportIssueSeverityParamsIssueType{intel.AttackSurfaceReportIssueSeverityParamsIssueTypeComplianceViolation, intel.AttackSurfaceReportIssueSeverityParamsIssueTypeEmailSecurity}),
		IssueTypeNeq:  cloudflare.F([]intel.AttackSurfaceReportIssueSeverityParamsIssueTypeNeq{intel.AttackSurfaceReportIssueSeverityParamsIssueTypeNeqComplianceViolation, intel.AttackSurfaceReportIssueSeverityParamsIssueTypeNeqEmailSecurity}),
		Product:       cloudflare.F([]string{"access", "dns"}),
		ProductNeq:    cloudflare.F([]string{"access", "dns"}),
		Severity:      cloudflare.F([]intel.AttackSurfaceReportIssueSeverityParamsSeverity{intel.AttackSurfaceReportIssueSeverityParamsSeverityLow, intel.AttackSurfaceReportIssueSeverityParamsSeverityModerate}),
		SeverityNeq:   cloudflare.F([]intel.AttackSurfaceReportIssueSeverityParamsSeverityNeq{intel.AttackSurfaceReportIssueSeverityParamsSeverityNeqLow, intel.AttackSurfaceReportIssueSeverityParamsSeverityNeqModerate}),
		Subject:       cloudflare.F([]string{"example.com", "example.com", "example.com"}),
		SubjectNeq:    cloudflare.F([]string{"example.com", "example.com", "example.com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackSurfaceReportIssueTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.Type(context.TODO(), intel.AttackSurfaceReportIssueTypeParams{
		AccountID:     cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Dismissed:     cloudflare.F(false),
		IssueClass:    cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueClassNeq: cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
		IssueType:     cloudflare.F([]intel.AttackSurfaceReportIssueTypeParamsIssueType{intel.AttackSurfaceReportIssueTypeParamsIssueTypeComplianceViolation, intel.AttackSurfaceReportIssueTypeParamsIssueTypeEmailSecurity}),
		IssueTypeNeq:  cloudflare.F([]intel.AttackSurfaceReportIssueTypeParamsIssueTypeNeq{intel.AttackSurfaceReportIssueTypeParamsIssueTypeNeqComplianceViolation, intel.AttackSurfaceReportIssueTypeParamsIssueTypeNeqEmailSecurity}),
		Product:       cloudflare.F([]string{"access", "dns"}),
		ProductNeq:    cloudflare.F([]string{"access", "dns"}),
		Severity:      cloudflare.F([]intel.AttackSurfaceReportIssueTypeParamsSeverity{intel.AttackSurfaceReportIssueTypeParamsSeverityLow, intel.AttackSurfaceReportIssueTypeParamsSeverityModerate}),
		SeverityNeq:   cloudflare.F([]intel.AttackSurfaceReportIssueTypeParamsSeverityNeq{intel.AttackSurfaceReportIssueTypeParamsSeverityNeqLow, intel.AttackSurfaceReportIssueTypeParamsSeverityNeqModerate}),
		Subject:       cloudflare.F([]string{"example.com", "example.com", "example.com"}),
		SubjectNeq:    cloudflare.F([]string{"example.com", "example.com", "example.com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
