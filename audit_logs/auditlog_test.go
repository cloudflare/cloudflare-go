// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audit_logs_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/audit_logs"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

func TestAuditLogListWithOptionalParams(t *testing.T) {
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
	_, err := client.AuditLogs.List(context.TODO(), audit_logs.AuditLogListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ID:        cloudflare.F("f174be97-19b1-40d6-954d-70cd5fbd52db"),
		Action: cloudflare.F(audit_logs.AuditLogListParamsAction{
			Type: cloudflare.F("add"),
		}),
		Actor: cloudflare.F(audit_logs.AuditLogListParamsActor{
			Email: cloudflare.F("alice@example.com"),
			IP:    cloudflare.F("17.168.228.63"),
		}),
		Before:       cloudflare.F[audit_logs.AuditLogListParamsBeforeUnion](shared.UnionTime(time.Now())),
		Direction:    cloudflare.F(audit_logs.AuditLogListParamsDirectionDesc),
		Export:       cloudflare.F(true),
		HideUserLogs: cloudflare.F(true),
		Page:         cloudflare.F(50.000000),
		PerPage:      cloudflare.F(25.000000),
		Since:        cloudflare.F[audit_logs.AuditLogListParamsSinceUnion](shared.UnionTime(time.Now())),
		Zone: cloudflare.F(audit_logs.AuditLogListParamsZone{
			Name: cloudflare.F("example.com"),
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
