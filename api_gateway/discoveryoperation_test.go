// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/api_gateway"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

func TestDiscoveryOperationListWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.Discovery.Operations.List(context.TODO(), api_gateway.DiscoveryOperationListParams{
		ZoneID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Diff:      cloudflare.F(true),
		Direction: cloudflare.F(api_gateway.DiscoveryOperationListParamsDirectionDesc),
		Endpoint:  cloudflare.F("/api/v1"),
		Host:      cloudflare.F([]string{"api.cloudflare.com"}),
		Method:    cloudflare.F([]string{"GET"}),
		Order:     cloudflare.F(api_gateway.DiscoveryOperationListParamsOrderMethod),
		Origin:    cloudflare.F(api_gateway.DiscoveryOperationListParamsOriginMl),
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(5)),
		State:     cloudflare.F(api_gateway.DiscoveryOperationListParamsStateReview),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDiscoveryOperationBulkEdit(t *testing.T) {
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
	_, err := client.APIGateway.Discovery.Operations.BulkEdit(context.TODO(), api_gateway.DiscoveryOperationBulkEditParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body: map[string]api_gateway.DiscoveryOperationBulkEditParamsBody{
			"3818d821-5901-4147-a474-f5f5aec1d54e": {
				State: cloudflare.F(api_gateway.DiscoveryOperationBulkEditParamsBodyStateIgnored),
			},
			"b17c8043-99a0-4202-b7d9-8f7cdbee02cd": {
				State: cloudflare.F(api_gateway.DiscoveryOperationBulkEditParamsBodyStateReview),
			},
		},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDiscoveryOperationEditWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.Discovery.Operations.Edit(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		api_gateway.DiscoveryOperationEditParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			State:  cloudflare.F(api_gateway.DiscoveryOperationEditParamsStateReview),
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
