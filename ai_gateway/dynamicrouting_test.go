// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/ai_gateway"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

func TestDynamicRoutingNew(t *testing.T) {
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
	_, err := client.AIGateway.DynamicRouting.New(
		context.TODO(),
		"54442216",
		ai_gateway.DynamicRoutingNewParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
			Elements: cloudflare.F([]ai_gateway.DynamicRoutingNewParamsElementUnion{ai_gateway.DynamicRoutingNewParamsElementsObject{
				ID: cloudflare.F("id"),
				Outputs: cloudflare.F(ai_gateway.DynamicRoutingNewParamsElementsObjectOutputs{
					Next: cloudflare.F(ai_gateway.DynamicRoutingNewParamsElementsObjectOutputsNext{
						ElementID: cloudflare.F("elementId"),
					}),
				}),
				Type: cloudflare.F(ai_gateway.DynamicRoutingNewParamsElementsObjectTypeStart),
			}}),
			Name: cloudflare.F("name"),
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

func TestDynamicRoutingUpdate(t *testing.T) {
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
	_, err := client.AIGateway.DynamicRouting.Update(
		context.TODO(),
		"54442216",
		"54442216",
		ai_gateway.DynamicRoutingUpdateParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
			Name:      cloudflare.F("Route Name"),
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

func TestDynamicRoutingList(t *testing.T) {
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
	_, err := client.AIGateway.DynamicRouting.List(
		context.TODO(),
		"54442216",
		ai_gateway.DynamicRoutingListParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
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

func TestDynamicRoutingDelete(t *testing.T) {
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
	_, err := client.AIGateway.DynamicRouting.Delete(
		context.TODO(),
		"54442216",
		"54442216",
		ai_gateway.DynamicRoutingDeleteParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
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

func TestDynamicRoutingNewDeployment(t *testing.T) {
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
	_, err := client.AIGateway.DynamicRouting.NewDeployment(
		context.TODO(),
		"54442216",
		"54442216",
		ai_gateway.DynamicRoutingNewDeploymentParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
			Comment:   cloudflare.F("Route Deployment Comment"),
			VersionID: cloudflare.F("54442216"),
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

func TestDynamicRoutingNewVersion(t *testing.T) {
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
	_, err := client.AIGateway.DynamicRouting.NewVersion(
		context.TODO(),
		"54442216",
		"54442216",
		ai_gateway.DynamicRoutingNewVersionParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
			Comment:   cloudflare.F("Route Version Comment"),
			Elements: cloudflare.F([]ai_gateway.DynamicRoutingNewVersionParamsElementUnion{ai_gateway.DynamicRoutingNewVersionParamsElementsObject{
				ID: cloudflare.F("id"),
				Outputs: cloudflare.F(ai_gateway.DynamicRoutingNewVersionParamsElementsObjectOutputs{
					Next: cloudflare.F(ai_gateway.DynamicRoutingNewVersionParamsElementsObjectOutputsNext{
						ElementID: cloudflare.F("elementId"),
					}),
				}),
				Type: cloudflare.F(ai_gateway.DynamicRoutingNewVersionParamsElementsObjectTypeStart),
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

func TestDynamicRoutingGet(t *testing.T) {
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
	_, err := client.AIGateway.DynamicRouting.Get(
		context.TODO(),
		"54442216",
		"54442216",
		ai_gateway.DynamicRoutingGetParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
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

func TestDynamicRoutingGetVersion(t *testing.T) {
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
	_, err := client.AIGateway.DynamicRouting.GetVersion(
		context.TODO(),
		"54442216",
		"54442216",
		"54442216",
		ai_gateway.DynamicRoutingGetVersionParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
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

func TestDynamicRoutingListDeployments(t *testing.T) {
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
	_, err := client.AIGateway.DynamicRouting.ListDeployments(
		context.TODO(),
		"54442216",
		"54442216",
		ai_gateway.DynamicRoutingListDeploymentsParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
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

func TestDynamicRoutingListVersions(t *testing.T) {
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
	_, err := client.AIGateway.DynamicRouting.ListVersions(
		context.TODO(),
		"54442216",
		"54442216",
		ai_gateway.DynamicRoutingListVersionsParams{
			AccountID: cloudflare.F("0d37909e38d3e99c29fa2cd343ac421a"),
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
