// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/load_balancers"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestSearchGetWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.LoadBalancers.Searches.Get(context.TODO(), load_balancers.SearchGetParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Page:      cloudflare.F(1.000000),
		PerPage:   cloudflare.F(1.000000),
		SearchParams: cloudflare.F(load_balancers.SearchGetParamsSearchParams{
			Query:      cloudflare.F("primary"),
			References: cloudflare.F(load_balancers.SearchGetParamsSearchParamsReferencesEmpty),
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