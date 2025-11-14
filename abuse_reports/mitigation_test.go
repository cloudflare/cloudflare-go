// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package abuse_reports_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/abuse_reports"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

func TestMitigationListWithOptionalParams(t *testing.T) {
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
	_, err := client.AbuseReports.Mitigations.List(
		context.TODO(),
		"report_id",
		abuse_reports.MitigationListParams{
			AccountID:       cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			EffectiveAfter:  cloudflare.F("2009-11-10T23:00:00Z"),
			EffectiveBefore: cloudflare.F("2009-11-10T23:00:00Z"),
			EntityType:      cloudflare.F(abuse_reports.MitigationListParamsEntityTypeURLPattern),
			Page:            cloudflare.F(int64(0)),
			PerPage:         cloudflare.F(int64(0)),
			Sort:            cloudflare.F(abuse_reports.MitigationListParamsSortTypeAsc),
			Status:          cloudflare.F(abuse_reports.MitigationListParamsStatusPending),
			Type:            cloudflare.F(abuse_reports.MitigationListParamsTypeLegalBlock),
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

func TestMitigationReview(t *testing.T) {
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
	_, err := client.AbuseReports.Mitigations.Review(
		context.TODO(),
		"report_id",
		abuse_reports.MitigationReviewParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Appeals: cloudflare.F([]abuse_reports.MitigationReviewParamsAppeal{{
				ID:     cloudflare.F("id"),
				Reason: cloudflare.F(abuse_reports.MitigationReviewParamsAppealsReasonMisclassified),
			}}),
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
