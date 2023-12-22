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

func TestAccountLoadBalancerPoolGet(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.LoadBalancers.Pools.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"17b5962d775c646f3f9725cbc7a53df4",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountLoadBalancerPoolUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.LoadBalancers.Pools.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"17b5962d775c646f3f9725cbc7a53df4",
		cloudflare.AccountLoadBalancerPoolUpdateParams{
			Name: cloudflare.F("primary-dc-1"),
			Origins: cloudflare.F([]cloudflare.AccountLoadBalancerPoolUpdateParamsOrigin{{
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.AccountLoadBalancerPoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.AccountLoadBalancerPoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.AccountLoadBalancerPoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}}),
			CheckRegions: cloudflare.F([]cloudflare.AccountLoadBalancerPoolUpdateParamsCheckRegion{cloudflare.AccountLoadBalancerPoolUpdateParamsCheckRegionWeu, cloudflare.AccountLoadBalancerPoolUpdateParamsCheckRegionEnam}),
			Description:  cloudflare.F("Primary data center - Provider XYZ"),
			Enabled:      cloudflare.F(false),
			Latitude:     cloudflare.F(0.000000),
			LoadShedding: cloudflare.F(cloudflare.AccountLoadBalancerPoolUpdateParamsLoadShedding{
				DefaultPercent: cloudflare.F(0.000000),
				DefaultPolicy:  cloudflare.F(cloudflare.AccountLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicyRandom),
				SessionPercent: cloudflare.F(0.000000),
				SessionPolicy:  cloudflare.F(cloudflare.AccountLoadBalancerPoolUpdateParamsLoadSheddingSessionPolicyHash),
			}),
			Longitude:         cloudflare.F(0.000000),
			MinimumOrigins:    cloudflare.F(int64(0)),
			Monitor:           cloudflare.F[any](map[string]interface{}{}),
			NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cloudflare.F(cloudflare.AccountLoadBalancerPoolUpdateParamsNotificationFilter{
				Origin: cloudflare.F(cloudflare.AccountLoadBalancerPoolUpdateParamsNotificationFilterOrigin{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(true),
				}),
				Pool: cloudflare.F(cloudflare.AccountLoadBalancerPoolUpdateParamsNotificationFilterPool{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(false),
				}),
			}),
			OriginSteering: cloudflare.F(cloudflare.AccountLoadBalancerPoolUpdateParamsOriginSteering{
				Policy: cloudflare.F(cloudflare.AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicyRandom),
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

func TestAccountLoadBalancerPoolDelete(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.LoadBalancers.Pools.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"17b5962d775c646f3f9725cbc7a53df4",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.LoadBalancers.Pools.AccountLoadBalancerPoolsNewPool(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParams{
			Name: cloudflare.F("primary-dc-1"),
			Origins: cloudflare.F([]cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOrigin{{
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}}),
			Description: cloudflare.F("Primary data center - Provider XYZ"),
			Enabled:     cloudflare.F(false),
			Latitude:    cloudflare.F(0.000000),
			LoadShedding: cloudflare.F(cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadShedding{
				DefaultPercent: cloudflare.F(0.000000),
				DefaultPolicy:  cloudflare.F(cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicyRandom),
				SessionPercent: cloudflare.F(0.000000),
				SessionPolicy:  cloudflare.F(cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicyHash),
			}),
			Longitude:         cloudflare.F(0.000000),
			MinimumOrigins:    cloudflare.F(int64(0)),
			Monitor:           cloudflare.F[any](map[string]interface{}{}),
			NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cloudflare.F(cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilter{
				Origin: cloudflare.F(cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterOrigin{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(true),
				}),
				Pool: cloudflare.F(cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterPool{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(false),
				}),
			}),
			OriginSteering: cloudflare.F(cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteering{
				Policy: cloudflare.F(cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyRandom),
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

func TestAccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.LoadBalancers.Pools.AccountLoadBalancerPoolsListPools(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsParams{
			Monitor: cloudflare.F[any](map[string]interface{}{}),
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

func TestAccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.LoadBalancers.Pools.AccountLoadBalancerPoolsPatchPools(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParams{
			NotificationEmail: cloudflare.F(cloudflare.AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParamsNotificationEmailEmpty),
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
