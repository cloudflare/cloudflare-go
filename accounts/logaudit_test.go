// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/accounts"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestLogAuditListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Logs.Audit.List(context.TODO(), accounts.LogAuditListParams{
		AccountID:       cloudflare.F("a67e14daa5f8dceeb91fe5449ba496ef"),
		Before:          cloudflare.F(time.Now()),
		Since:           cloudflare.F(time.Now()),
		AccountName:     cloudflare.F("account_name"),
		ActionResult:    cloudflare.F(accounts.LogAuditListParamsActionResultSuccess),
		ActionType:      cloudflare.F(accounts.LogAuditListParamsActionTypeCreate),
		ActorContext:    cloudflare.F(accounts.LogAuditListParamsActorContextAPIKey),
		ActorEmail:      cloudflare.F("alice@example.com"),
		ActorID:         cloudflare.F("1d20c3afe174f18b642710cec6298a9d"),
		ActorIPAddress:  cloudflare.F("17.168.228.63"),
		ActorTokenID:    cloudflare.F("144cdb2e39c55e203cf225d8d8208647"),
		ActorTokenName:  cloudflare.F("Test Token"),
		ActorType:       cloudflare.F(accounts.LogAuditListParamsActorTypeCloudflareAdmin),
		AuditLogID:      cloudflare.F("f174be97-19b1-40d6-954d-70cd5fbd52db"),
		Cursor:          cloudflare.F("Q1buH-__DQqqig7SVYXT-SsMOTGY2Z3Y80W-fGgva7yaDdmPKveucH5ddOcHsJRhNb-xUK8agZQqkJSMAENGO8NU6g=="),
		Direction:       cloudflare.F(accounts.LogAuditListParamsDirectionDesc),
		Limit:           cloudflare.F(25.000000),
		RawCfRayID:      cloudflare.F("8e8dd2156ef28414"),
		RawMethod:       cloudflare.F("GET"),
		RawStatusCode:   cloudflare.F(int64(200)),
		RawURI:          cloudflare.F("raw_uri"),
		ResourceID:      cloudflare.F("resource_id"),
		ResourceProduct: cloudflare.F("Stream"),
		ResourceScope:   cloudflare.F(accounts.LogAuditListParamsResourceScopeAccounts),
		ResourceType:    cloudflare.F("Video"),
		ZoneID:          cloudflare.F("zone_id"),
		ZoneName:        cloudflare.F("example.com"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
