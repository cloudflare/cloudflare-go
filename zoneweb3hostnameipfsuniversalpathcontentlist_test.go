// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetails(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Web3s.Hostnames.IpfsUniversalPaths.ContentLists.Web3HostnameIpfsUniversalPathGatewayContentListDetails(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentList(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Web3s.Hostnames.IpfsUniversalPaths.ContentLists.Web3HostnameUpdateIpfsUniversalPathGatewayContentList(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParams{
			Action: cloudflare.F(cloudflare.ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsActionBlock),
			Entries: cloudflare.F([]cloudflare.ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntry{{
				Content:     cloudflare.F("QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB"),
				Description: cloudflare.F("this is my content list entry"),
				Type:        cloudflare.F(cloudflare.ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesTypeCid),
			}, {
				Content:     cloudflare.F("QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB"),
				Description: cloudflare.F("this is my content list entry"),
				Type:        cloudflare.F(cloudflare.ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesTypeCid),
			}, {
				Content:     cloudflare.F("QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB"),
				Description: cloudflare.F("this is my content list entry"),
				Type:        cloudflare.F(cloudflare.ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesTypeCid),
			}}),
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
