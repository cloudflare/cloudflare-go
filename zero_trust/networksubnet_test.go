// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
)

func TestNetworkSubnetListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.Subnets.List(context.TODO(), zero_trust.NetworkSubnetListParams{
		AccountID:        cloudflare.F("699d98642c564d2e855e9661899b7252"),
		AddressFamily:    cloudflare.F(zero_trust.NetworkSubnetListParamsAddressFamilyV4),
		Comment:          cloudflare.F("example%20comment"),
		ExistedAt:        cloudflare.F("2019-10-12T07%3A20%3A50.52Z"),
		IsDefaultNetwork: cloudflare.F(true),
		IsDeleted:        cloudflare.F(true),
		Name:             cloudflare.F("IPv4%20Cloudflare%20Source%20IPs"),
		Network:          cloudflare.F("172.16.0.0%2F16"),
		Page:             cloudflare.F(1.000000),
		PerPage:          cloudflare.F(1.000000),
		SortOrder:        cloudflare.F(zero_trust.NetworkSubnetListParamsSortOrderAsc),
		SubnetTypes:      cloudflare.F(zero_trust.NetworkSubnetListParamsSubnetTypesCloudflareSource),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
