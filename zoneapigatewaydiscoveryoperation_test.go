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

func TestZoneAPIGatewayDiscoveryOperationUpdate(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Discovery.Operations.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneAPIGatewayDiscoveryOperationUpdateParams{
			Body: cloudflare.F[any](map[string]interface{}{
				"3818d821-5901-4147-a474-f5f5aec1d54e": map[string]interface{}{
					"state": "ignored",
				},
				"b17c8043-99a0-4202-b7d9-8f7cdbee02cd": map[string]interface{}{
					"state": "review",
				},
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

func TestZoneAPIGatewayDiscoveryOperationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Discovery.Operations.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneAPIGatewayDiscoveryOperationListParams{
			Diff:      cloudflare.F(true),
			Direction: cloudflare.F(cloudflare.ZoneAPIGatewayDiscoveryOperationListParamsDirectionDesc),
			Endpoint:  cloudflare.F("/api/v1"),
			Host:      cloudflare.F([]string{"api.cloudflare.com"}),
			Method:    cloudflare.F([]string{"GET"}),
			Order:     cloudflare.F(cloudflare.ZoneAPIGatewayDiscoveryOperationListParamsOrderMethod),
			Origin:    cloudflare.F(cloudflare.ZoneAPIGatewayDiscoveryOperationListParamsOriginMl),
			Page:      cloudflare.F[any](map[string]interface{}{}),
			PerPage:   cloudflare.F[any](map[string]interface{}{}),
			State:     cloudflare.F(cloudflare.ZoneAPIGatewayDiscoveryOperationListParamsStateReview),
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

func TestZoneAPIGatewayDiscoveryOperationUpdateStateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Discovery.Operations.UpdateState(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0d9bf70c-92e1-4bb3-9411-34a3bcc59003",
		cloudflare.ZoneAPIGatewayDiscoveryOperationUpdateStateParams{
			State: cloudflare.F(cloudflare.ZoneAPIGatewayDiscoveryOperationUpdateStateParamsStateReview),
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
