// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_cloud_networking_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/magic_cloud_networking"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

func TestCatalogSyncPrebuiltPolicyListWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicCloudNetworking.CatalogSyncs.PrebuiltPolicies.List(context.TODO(), magic_cloud_networking.CatalogSyncPrebuiltPolicyListParams{
		AccountID:       cloudflare.F("account_id"),
		DestinationType: cloudflare.F(magic_cloud_networking.CatalogSyncPrebuiltPolicyListParamsDestinationTypeNone),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
