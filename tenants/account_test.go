// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tenants_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/tenants"
)

func TestAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.Tenants.Accounts.List(
		context.TODO(),
		"tenant_id",
		tenants.AccountListParams{
			AccountPubname: cloudflare.F(tenants.AccountListParamsAccountPubname{
				Contains:   cloudflare.F("contains"),
				EndsWith:   cloudflare.F("endsWith"),
				StartsWith: cloudflare.F("startsWith"),
			}),
			Direction:    cloudflare.F(tenants.AccountListParamsDirectionAsc),
			IncludeTotal: cloudflare.F(true),
			Name: cloudflare.F(tenants.AccountListParamsName{
				Contains:   cloudflare.F("contains"),
				EndsWith:   cloudflare.F("endsWith"),
				StartsWith: cloudflare.F("startsWith"),
			}),
			OrderBy:   cloudflare.F(tenants.AccountListParamsOrderByAccountName),
			PageSize:  cloudflare.F(int64(0)),
			PageToken: cloudflare.F("page_token"),
			Status:    cloudflare.F(tenants.AccountListParamsStatusActive),
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
