// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/organizations"
)

func TestLogAuditListWithOptionalParams(t *testing.T) {
	t.Skip("TODO: required params 'since' and 'before' not populated by Prism mock")
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
	_, err := client.Organizations.Logs.Audit.List(
		context.TODO(),
		"a67e14daa5f8dceeb91fe5449ba496ef",
		organizations.LogAuditListParams{
			Before: cloudflare.F(time.Now()),
			Since:  cloudflare.F(time.Now()),
			ID: cloudflare.F(organizations.LogAuditListParamsID{
				Not: cloudflare.F([]string{"f174be97-19b1-40d6-954d-70cd5fbd52db"}),
			}),
			ActionResult: cloudflare.F(organizations.LogAuditListParamsActionResult{
				Not: cloudflare.F([]organizations.LogAuditListParamsActionResultNot{organizations.LogAuditListParamsActionResultNotSuccess}),
			}),
			ActionType: cloudflare.F(organizations.LogAuditListParamsActionType{
				Not: cloudflare.F([]organizations.LogAuditListParamsActionTypeNot{organizations.LogAuditListParamsActionTypeNotCreate}),
			}),
			ActorContext: cloudflare.F(organizations.LogAuditListParamsActorContext{
				Not: cloudflare.F([]organizations.LogAuditListParamsActorContextNot{organizations.LogAuditListParamsActorContextNotAPIKey}),
			}),
			ActorEmail: cloudflare.F(organizations.LogAuditListParamsActorEmail{
				Not: cloudflare.F([]string{"alice@example.com"}),
			}),
			ActorID: cloudflare.F(organizations.LogAuditListParamsActorID{
				Not: cloudflare.F([]string{"1d20c3afe174f18b642710cec6298a9d"}),
			}),
			ActorIPAddress: cloudflare.F(organizations.LogAuditListParamsActorIPAddress{
				Not: cloudflare.F([]string{"17.168.228.63"}),
			}),
			ActorTokenID: cloudflare.F(organizations.LogAuditListParamsActorTokenID{
				Not: cloudflare.F([]string{"144cdb2e39c55e203cf225d8d8208647"}),
			}),
			ActorTokenName: cloudflare.F(organizations.LogAuditListParamsActorTokenName{
				Not: cloudflare.F([]string{"Test Token"}),
			}),
			ActorType: cloudflare.F(organizations.LogAuditListParamsActorType{
				Not: cloudflare.F([]organizations.LogAuditListParamsActorTypeNot{organizations.LogAuditListParamsActorTypeNotUser}),
			}),
			Cursor:    cloudflare.F("Q1buH-__DQqqig7SVYXT-SsMOTGY2Z3Y80W-fGgva7yaDdmPKveucH5ddOcHsJRhNb-xUK8agZQqkJSMAENGO8NU6g=="),
			Direction: cloudflare.F(organizations.LogAuditListParamsDirectionDesc),
			Limit:     cloudflare.F(25.000000),
			RawCfRayID: cloudflare.F(organizations.LogAuditListParamsRawCfRayID{
				Not: cloudflare.F([]string{"8e8dd2156ef28414"}),
			}),
			RawMethod: cloudflare.F(organizations.LogAuditListParamsRawMethod{
				Not: cloudflare.F([]string{"GET"}),
			}),
			RawStatusCode: cloudflare.F(organizations.LogAuditListParamsRawStatusCode{
				Not: cloudflare.F([]int64{int64(200)}),
			}),
			RawURI: cloudflare.F(organizations.LogAuditListParamsRawURI{
				Not: cloudflare.F([]string{"string"}),
			}),
			ResourceID: cloudflare.F(organizations.LogAuditListParamsResourceID{
				Not: cloudflare.F([]string{"string"}),
			}),
			ResourceProduct: cloudflare.F(organizations.LogAuditListParamsResourceProduct{
				Not: cloudflare.F([]string{"Stream"}),
			}),
			ResourceScope: cloudflare.F(organizations.LogAuditListParamsResourceScope{
				Not: cloudflare.F([]organizations.LogAuditListParamsResourceScopeNot{organizations.LogAuditListParamsResourceScopeNotOrganizations}),
			}),
			ResourceType: cloudflare.F(organizations.LogAuditListParamsResourceType{
				Not: cloudflare.F([]string{"Video"}),
			}),
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
