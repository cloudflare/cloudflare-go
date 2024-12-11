// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/api_gateway"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestUserSchemaOperationListWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.UserSchemas.Operations.List(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		api_gateway.UserSchemaOperationListParams{
			ZoneID:          cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Endpoint:        cloudflare.F("/api/v1"),
			Feature:         cloudflare.F([]api_gateway.UserSchemaOperationListParamsFeature{api_gateway.UserSchemaOperationListParamsFeatureThresholds}),
			Host:            cloudflare.F([]string{"api.cloudflare.com"}),
			Method:          cloudflare.F([]string{"GET"}),
			OperationStatus: cloudflare.F(api_gateway.UserSchemaOperationListParamsOperationStatusNew),
			Page:            cloudflare.F(int64(1)),
			PerPage:         cloudflare.F(int64(5)),
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
