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
	t.Skip("TODO:investigate broken test")
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
		AccountID: cloudflare.F("a67e14daa5f8dceeb91fe5449ba496ef"),
		Before:    cloudflare.F(time.Now()),
		Since:     cloudflare.F(time.Now()),
		AccountName: cloudflare.F(accounts.LogAuditListParamsAccountName{
			Not: cloudflare.F([]string{"string"}),
		}),
		ActionResult: cloudflare.F(accounts.LogAuditListParamsActionResult{
			Not: cloudflare.F([]accounts.LogAuditListParamsActionResultNot{accounts.LogAuditListParamsActionResultNotSuccess}),
		}),
		ActionType: cloudflare.F(accounts.LogAuditListParamsActionType{
			Not: cloudflare.F([]accounts.LogAuditListParamsActionTypeNot{accounts.LogAuditListParamsActionTypeNotCreate}),
		}),
		ActorContext: cloudflare.F(accounts.LogAuditListParamsActorContext{
			Not: cloudflare.F([]accounts.LogAuditListParamsActorContextNot{accounts.LogAuditListParamsActorContextNotAPIKey}),
		}),
		ActorEmail: cloudflare.F(accounts.LogAuditListParamsActorEmail{
			Not: cloudflare.F([]string{"alice@example.com"}),
		}),
		ActorID: cloudflare.F(accounts.LogAuditListParamsActorID{
			Not: cloudflare.F([]string{"1d20c3afe174f18b642710cec6298a9d"}),
		}),
		ActorIPAddress: cloudflare.F(accounts.LogAuditListParamsActorIPAddress{
			Not: cloudflare.F([]string{"17.168.228.63"}),
		}),
		ActorTokenID: cloudflare.F(accounts.LogAuditListParamsActorTokenID{
			Not: cloudflare.F([]string{"144cdb2e39c55e203cf225d8d8208647"}),
		}),
		ActorTokenName: cloudflare.F(accounts.LogAuditListParamsActorTokenName{
			Not: cloudflare.F([]string{"Test Token"}),
		}),
		ActorType: cloudflare.F(accounts.LogAuditListParamsActorType{
			Not: cloudflare.F([]accounts.LogAuditListParamsActorTypeNot{accounts.LogAuditListParamsActorTypeNotAccount}),
		}),
		AuditLogID: cloudflare.F(accounts.LogAuditListParamsAuditLogID{
			Not: cloudflare.F([]string{"f174be97-19b1-40d6-954d-70cd5fbd52db"}),
		}),
		Cursor:    cloudflare.F("Q1buH-__DQqqig7SVYXT-SsMOTGY2Z3Y80W-fGgva7yaDdmPKveucH5ddOcHsJRhNb-xUK8agZQqkJSMAENGO8NU6g=="),
		Direction: cloudflare.F(accounts.LogAuditListParamsDirectionDesc),
		Limit:     cloudflare.F(25.000000),
		RawCfRayID: cloudflare.F(accounts.LogAuditListParamsRawCfRayID{
			Not: cloudflare.F([]string{"8e8dd2156ef28414"}),
		}),
		RawMethod: cloudflare.F(accounts.LogAuditListParamsRawMethod{
			Not: cloudflare.F([]string{"GET"}),
		}),
		RawStatusCode: cloudflare.F(accounts.LogAuditListParamsRawStatusCode{
			Not: cloudflare.F([]int64{int64(200)}),
		}),
		RawURI: cloudflare.F(accounts.LogAuditListParamsRawURI{
			Not: cloudflare.F([]string{"string"}),
		}),
		ResourceID: cloudflare.F(accounts.LogAuditListParamsResourceID{
			Not: cloudflare.F([]string{"string"}),
		}),
		ResourceProduct: cloudflare.F(accounts.LogAuditListParamsResourceProduct{
			Not: cloudflare.F([]string{"Stream"}),
		}),
		ResourceScope: cloudflare.F(accounts.LogAuditListParamsResourceScope{
			Not: cloudflare.F([]accounts.LogAuditListParamsResourceScopeNot{accounts.LogAuditListParamsResourceScopeNotAccounts}),
		}),
		ResourceType: cloudflare.F(accounts.LogAuditListParamsResourceType{
			Not: cloudflare.F([]string{"Video"}),
		}),
		ZoneID: cloudflare.F(accounts.LogAuditListParamsZoneID{
			Not: cloudflare.F([]string{"string"}),
		}),
		ZoneName: cloudflare.F(accounts.LogAuditListParamsZoneName{
			Not: cloudflare.F([]string{"example.com"}),
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
