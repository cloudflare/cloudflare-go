// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/api_gateway"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestOperationNew(t *testing.T) {
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
	_, err := client.APIGateway.Operations.New(context.TODO(), api_gateway.OperationNewParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body: []api_gateway.OperationNewParamsBody{{
			Endpoint: cloudflare.F("/api/v1/users/{var1}"),
			Host:     cloudflare.F("www.example.com"),
			Method:   cloudflare.F(api_gateway.OperationNewParamsBodyMethodGet),
		}, {
			Endpoint: cloudflare.F("/api/v1/users/{var1}"),
			Host:     cloudflare.F("www.example.com"),
			Method:   cloudflare.F(api_gateway.OperationNewParamsBodyMethodGet),
		}, {
			Endpoint: cloudflare.F("/api/v1/users/{var1}"),
			Host:     cloudflare.F("www.example.com"),
			Method:   cloudflare.F(api_gateway.OperationNewParamsBodyMethodGet),
		}},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperationListWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.Operations.List(context.TODO(), api_gateway.OperationListParams{
		ZoneID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: cloudflare.F(api_gateway.OperationListParamsDirectionDesc),
		Endpoint:  cloudflare.F("/api/v1"),
		Feature:   cloudflare.F([]api_gateway.OperationListParamsFeature{api_gateway.OperationListParamsFeatureThresholds}),
		Host:      cloudflare.F([]string{"api.cloudflare.com"}),
		Method:    cloudflare.F([]string{"GET"}),
		Order:     cloudflare.F(api_gateway.OperationListParamsOrderMethod),
		Page:      cloudflare.F[any](map[string]interface{}{}),
		PerPage:   cloudflare.F(5.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperationDelete(t *testing.T) {
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
	_, err := client.APIGateway.Operations.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		api_gateway.OperationDeleteParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestOperationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.Operations.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		api_gateway.OperationGetParams{
			ZoneID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Feature: cloudflare.F([]api_gateway.OperationGetParamsFeature{api_gateway.OperationGetParamsFeatureThresholds}),
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
