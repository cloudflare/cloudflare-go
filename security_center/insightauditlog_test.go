// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_center_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/security_center"
)

func TestInsightAuditLogListWithOptionalParams(t *testing.T) {
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
	_, err := client.SecurityCenter.Insights.AuditLogs.List(context.TODO(), security_center.InsightAuditLogListParams{
		AccountID:    cloudflare.F("account_id"),
		Before:       cloudflare.F(time.Now()),
		ChangedBy:    cloudflare.F("changed_by"),
		Cursor:       cloudflare.F("cursor"),
		FieldChanged: cloudflare.F(security_center.InsightAuditLogListParamsFieldChangedStatus),
		Order:        cloudflare.F(security_center.InsightAuditLogListParamsOrderAsc),
		PerPage:      cloudflare.F(int64(1)),
		Since:        cloudflare.F(time.Now()),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInsightAuditLogListByInsightWithOptionalParams(t *testing.T) {
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
	_, err := client.SecurityCenter.Insights.AuditLogs.ListByInsight(
		context.TODO(),
		"issue_id",
		security_center.InsightAuditLogListByInsightParams{
			AccountID:    cloudflare.F("account_id"),
			Before:       cloudflare.F(time.Now()),
			ChangedBy:    cloudflare.F("changed_by"),
			Cursor:       cloudflare.F("cursor"),
			FieldChanged: cloudflare.F(security_center.InsightAuditLogListByInsightParamsFieldChangedStatus),
			Order:        cloudflare.F(security_center.InsightAuditLogListByInsightParamsOrderAsc),
			PerPage:      cloudflare.F(int64(1)),
			Since:        cloudflare.F(time.Now()),
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
