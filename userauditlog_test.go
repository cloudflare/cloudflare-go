// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestUserAuditLogListWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.User.AuditLogs.List(context.TODO(), cloudflare.UserAuditLogListParams{
		ID: cloudflare.F("f174be97-19b1-40d6-954d-70cd5fbd52db"),
		Action: cloudflare.F(cloudflare.UserAuditLogListParamsAction{
			Type: cloudflare.F("add"),
		}),
		Actor: cloudflare.F(cloudflare.UserAuditLogListParamsActor{
			IP:    cloudflare.F("17.168.228.63"),
			Email: cloudflare.F("alice@example.com"),
		}),
		Before:       cloudflare.F(time.Now()),
		Direction:    cloudflare.F(cloudflare.UserAuditLogListParamsDirectionDesc),
		Export:       cloudflare.F(true),
		HideUserLogs: cloudflare.F(true),
		Page:         cloudflare.F(50.000000),
		PerPage:      cloudflare.F(25.000000),
		Since:        cloudflare.F(time.Now()),
		Zone: cloudflare.F(cloudflare.UserAuditLogListParamsZone{
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
