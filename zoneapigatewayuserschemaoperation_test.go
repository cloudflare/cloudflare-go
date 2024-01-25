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

func TestZoneAPIGatewayUserSchemaOperationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.UserSchemas.Operations.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZoneAPIGatewayUserSchemaOperationListParams{
			Endpoint:        cloudflare.F("/api/v1"),
			Feature:         cloudflare.F([]cloudflare.ZoneAPIGatewayUserSchemaOperationListParamsFeature{cloudflare.ZoneAPIGatewayUserSchemaOperationListParamsFeatureThresholds}),
			Host:            cloudflare.F([]string{"api.cloudflare.com"}),
			Method:          cloudflare.F([]string{"GET"}),
			OperationStatus: cloudflare.F(cloudflare.ZoneAPIGatewayUserSchemaOperationListParamsOperationStatusNew),
			Page:            cloudflare.F[any](map[string]interface{}{}),
			PerPage:         cloudflare.F[any](map[string]interface{}{}),
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
