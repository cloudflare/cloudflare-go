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

func TestUserLoadBalancerPoolGet(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Pools.Get(context.TODO(), "17b5962d775c646f3f9725cbc7a53df4")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerPoolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Pools.Update(
		context.TODO(),
		"17b5962d775c646f3f9725cbc7a53df4",
		cloudflare.UserLoadBalancerPoolUpdateParams{
			Name: cloudflare.F("primary-dc-1"),
			Origins: cloudflare.F([]cloudflare.UserLoadBalancerPoolUpdateParamsOrigin{{
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}}),
			CheckRegions: cloudflare.F([]cloudflare.UserLoadBalancerPoolUpdateParamsCheckRegion{cloudflare.UserLoadBalancerPoolUpdateParamsCheckRegionWeu, cloudflare.UserLoadBalancerPoolUpdateParamsCheckRegionEnam}),
			Description:  cloudflare.F("Primary data center - Provider XYZ"),
			Enabled:      cloudflare.F(false),
			Latitude:     cloudflare.F(0.000000),
			LoadShedding: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsLoadShedding{
				DefaultPercent: cloudflare.F(0.000000),
				DefaultPolicy:  cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicyRandom),
				SessionPercent: cloudflare.F(0.000000),
				SessionPolicy:  cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsLoadSheddingSessionPolicyHash),
			}),
			Longitude:         cloudflare.F(0.000000),
			MinimumOrigins:    cloudflare.F(int64(0)),
			Monitor:           cloudflare.F[any](map[string]interface{}{}),
			NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsNotificationFilter{
				Origin: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsNotificationFilterOrigin{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(true),
				}),
				Pool: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsNotificationFilterPool{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(false),
				}),
			}),
			OriginSteering: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsOriginSteering{
				Policy: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsOriginSteeringPolicyRandom),
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

func TestUserLoadBalancerPoolDelete(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Pools.Delete(context.TODO(), "17b5962d775c646f3f9725cbc7a53df4")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerPoolLoadBalancerPoolsNewPoolWithOptionalParams(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Pools.LoadBalancerPoolsNewPool(context.TODO(), cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParams{
		Name: cloudflare.F("primary-dc-1"),
		Origins: cloudflare.F([]cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOrigin{{
			Address: cloudflare.F("0.0.0.0"),
			Enabled: cloudflare.F(true),
			Header: cloudflare.F(cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginsHeader{
				Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			}),
			Name:             cloudflare.F("app-server-1"),
			VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           cloudflare.F(0.600000),
		}, {
			Address: cloudflare.F("0.0.0.0"),
			Enabled: cloudflare.F(true),
			Header: cloudflare.F(cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginsHeader{
				Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			}),
			Name:             cloudflare.F("app-server-1"),
			VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           cloudflare.F(0.600000),
		}, {
			Address: cloudflare.F("0.0.0.0"),
			Enabled: cloudflare.F(true),
			Header: cloudflare.F(cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginsHeader{
				Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			}),
			Name:             cloudflare.F("app-server-1"),
			VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           cloudflare.F(0.600000),
		}}),
		CheckRegions: cloudflare.F([]cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion{cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionWeu, cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionEnam}),
		Description:  cloudflare.F("Primary data center - Provider XYZ"),
		Enabled:      cloudflare.F(false),
		Latitude:     cloudflare.F(0.000000),
		LoadShedding: cloudflare.F(cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadShedding{
			DefaultPercent: cloudflare.F(0.000000),
			DefaultPolicy:  cloudflare.F(cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicyRandom),
			SessionPercent: cloudflare.F(0.000000),
			SessionPolicy:  cloudflare.F(cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicyHash),
		}),
		Longitude:         cloudflare.F(0.000000),
		MinimumOrigins:    cloudflare.F(int64(0)),
		Monitor:           cloudflare.F[any](map[string]interface{}{}),
		NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
		NotificationFilter: cloudflare.F(cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsNotificationFilter{
			Origin: cloudflare.F(cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsNotificationFilterOrigin{
				Disable: cloudflare.F(true),
				Healthy: cloudflare.F(true),
			}),
			Pool: cloudflare.F(cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsNotificationFilterPool{
				Disable: cloudflare.F(true),
				Healthy: cloudflare.F(false),
			}),
		}),
		OriginSteering: cloudflare.F(cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteering{
			Policy: cloudflare.F(cloudflare.UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyRandom),
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerPoolLoadBalancerPoolsListPoolsWithOptionalParams(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Pools.LoadBalancerPoolsListPools(context.TODO(), cloudflare.UserLoadBalancerPoolLoadBalancerPoolsListPoolsParams{
		Monitor: cloudflare.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerPoolLoadBalancerPoolsPatchPoolsWithOptionalParams(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Pools.LoadBalancerPoolsPatchPools(context.TODO(), cloudflare.UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsParams{
		NotificationEmail: cloudflare.F(cloudflare.UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsParamsNotificationEmailEmpty),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerPoolPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Pools.Patch(
		context.TODO(),
		"17b5962d775c646f3f9725cbc7a53df4",
		cloudflare.UserLoadBalancerPoolPatchParams{
			CheckRegions: cloudflare.F([]cloudflare.UserLoadBalancerPoolPatchParamsCheckRegion{cloudflare.UserLoadBalancerPoolPatchParamsCheckRegionWeu, cloudflare.UserLoadBalancerPoolPatchParamsCheckRegionEnam}),
			Description:  cloudflare.F("Primary data center - Provider XYZ"),
			Enabled:      cloudflare.F(false),
			Latitude:     cloudflare.F(0.000000),
			LoadShedding: cloudflare.F(cloudflare.UserLoadBalancerPoolPatchParamsLoadShedding{
				DefaultPercent: cloudflare.F(0.000000),
				DefaultPolicy:  cloudflare.F(cloudflare.UserLoadBalancerPoolPatchParamsLoadSheddingDefaultPolicyRandom),
				SessionPercent: cloudflare.F(0.000000),
				SessionPolicy:  cloudflare.F(cloudflare.UserLoadBalancerPoolPatchParamsLoadSheddingSessionPolicyHash),
			}),
			Longitude:         cloudflare.F(0.000000),
			MinimumOrigins:    cloudflare.F(int64(0)),
			Monitor:           cloudflare.F[any](map[string]interface{}{}),
			Name:              cloudflare.F("primary-dc-1"),
			NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cloudflare.F(cloudflare.UserLoadBalancerPoolPatchParamsNotificationFilter{
				Origin: cloudflare.F(cloudflare.UserLoadBalancerPoolPatchParamsNotificationFilterOrigin{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(true),
				}),
				Pool: cloudflare.F(cloudflare.UserLoadBalancerPoolPatchParamsNotificationFilterPool{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(false),
				}),
			}),
			OriginSteering: cloudflare.F(cloudflare.UserLoadBalancerPoolPatchParamsOriginSteering{
				Policy: cloudflare.F(cloudflare.UserLoadBalancerPoolPatchParamsOriginSteeringPolicyRandom),
			}),
			Origins: cloudflare.F([]cloudflare.UserLoadBalancerPoolPatchParamsOrigin{{
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.UserLoadBalancerPoolPatchParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.UserLoadBalancerPoolPatchParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.UserLoadBalancerPoolPatchParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
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
