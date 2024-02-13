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

func TestAPIGatewayOperationUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.APIGateways.Operations.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0d9bf70c-92e1-4bb3-9411-34a3bcc59003",
		cloudflare.APIGatewayOperationUpdateParams{
			State: cloudflare.F(cloudflare.APIGatewayOperationUpdateParamsStateReview),
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

func TestAPIGatewayOperationListWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.APIGateways.Operations.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.APIGatewayOperationListParams{
			Diff:      cloudflare.F(true),
			Direction: cloudflare.F(cloudflare.APIGatewayOperationListParamsDirectionDesc),
			Endpoint:  cloudflare.F("/api/v1"),
			Host:      cloudflare.F([]string{"api.cloudflare.com"}),
			Method:    cloudflare.F([]string{"GET"}),
			Order:     cloudflare.F(cloudflare.APIGatewayOperationListParamsOrderMethod),
			Origin:    cloudflare.F(cloudflare.APIGatewayOperationListParamsOriginMl),
			Page:      cloudflare.F[any](map[string]interface{}{}),
			PerPage:   cloudflare.F[any](map[string]interface{}{}),
			State:     cloudflare.F(cloudflare.APIGatewayOperationListParamsStateReview),
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

func TestAPIGatewayOperationDelete(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.APIGateways.Operations.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZone(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.APIGateways.Operations.APIShieldEndpointManagementAddOperationsToAZone(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParams{
			Body: cloudflare.F([]cloudflare.APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBody{{
				Endpoint: cloudflare.F("/api/v1/users/{var1}"),
				Host:     cloudflare.F("www.example.com"),
				Method:   cloudflare.F(cloudflare.APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethodGet),
			}, {
				Endpoint: cloudflare.F("/api/v1/users/{var1}"),
				Host:     cloudflare.F("www.example.com"),
				Method:   cloudflare.F(cloudflare.APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethodGet),
			}, {
				Endpoint: cloudflare.F("/api/v1/users/{var1}"),
				Host:     cloudflare.F("www.example.com"),
				Method:   cloudflare.F(cloudflare.APIGatewayOperationAPIShieldEndpointManagementAddOperationsToAZoneParamsBodyMethodGet),
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

func TestAPIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.APIGateways.Operations.APIShieldEndpointManagementGetInformationAboutAllOperationsOnAZone(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParams{
			Direction: cloudflare.F(cloudflare.APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsDirectionDesc),
			Endpoint:  cloudflare.F("/api/v1"),
			Feature:   cloudflare.F([]cloudflare.APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsFeature{cloudflare.APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsFeatureThresholds}),
			Host:      cloudflare.F([]string{"api.cloudflare.com"}),
			Method:    cloudflare.F([]string{"GET"}),
			Order:     cloudflare.F(cloudflare.APIGatewayOperationAPIShieldEndpointManagementGetInformationAboutAllOperationsOnAZoneParamsOrderMethod),
			Page:      cloudflare.F[any](map[string]interface{}{}),
			PerPage:   cloudflare.F(5.000000),
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

func TestAPIGatewayOperationGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.APIGateways.Operations.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.APIGatewayOperationGetParams{
			Feature: cloudflare.F([]cloudflare.APIGatewayOperationGetParamsFeature{cloudflare.APIGatewayOperationGetParamsFeatureThresholds}),
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
