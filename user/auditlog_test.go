// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/cloudflare/cloudflare-go/v4/user"
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
	_, err := client.User.AuditLogs.List(context.TODO(), user.AuditLogListParams{
		ID: cloudflare.F("f174be97-19b1-40d6-954d-70cd5fbd52db"),
		Action: cloudflare.F(user.AuditLogListParamsAction{
			Type: cloudflare.F("add"),
		}),
		Actor: cloudflare.F(user.AuditLogListParamsActor{
			Email: cloudflare.F("alice@example.com"),
			IP:    cloudflare.F("17.168.228.63"),
		}),
		Before:       cloudflare.F[user.AuditLogListParamsBeforeUnion](shared.UnionTime(time.Now())),
		Direction:    cloudflare.F(user.AuditLogListParamsDirectionDesc),
		Export:       cloudflare.F(true),
		HideUserLogs: cloudflare.F(true),
		Page:         cloudflare.F(50.000000),
		PerPage:      cloudflare.F(25.000000),
		Since:        cloudflare.F[user.AuditLogListParamsSinceUnion](shared.UnionTime(time.Now())),
		Zone: cloudflare.F(user.AuditLogListParamsZone{
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
