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

func TestLoadBalancerPoolNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Pools.New(context.TODO(), cloudflare.LoadBalancerPoolNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Name:      cloudflare.F("primary-dc-1"),
		Origins: cloudflare.F([]cloudflare.LoadBalancerPoolNewParamsOrigin{{
			Address: cloudflare.F("0.0.0.0"),
			Enabled: cloudflare.F(true),
			Header: cloudflare.F(cloudflare.LoadBalancerPoolNewParamsOriginsHeader{
				Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			}),
			Name:             cloudflare.F("app-server-1"),
			VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           cloudflare.F(0.600000),
		}, {
			Address: cloudflare.F("0.0.0.0"),
			Enabled: cloudflare.F(true),
			Header: cloudflare.F(cloudflare.LoadBalancerPoolNewParamsOriginsHeader{
				Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			}),
			Name:             cloudflare.F("app-server-1"),
			VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           cloudflare.F(0.600000),
		}, {
			Address: cloudflare.F("0.0.0.0"),
			Enabled: cloudflare.F(true),
			Header: cloudflare.F(cloudflare.LoadBalancerPoolNewParamsOriginsHeader{
				Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			}),
			Name:             cloudflare.F("app-server-1"),
			VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           cloudflare.F(0.600000),
		}}),
		Description: cloudflare.F("Primary data center - Provider XYZ"),
		Enabled:     cloudflare.F(false),
		Latitude:    cloudflare.F(0.000000),
		LoadShedding: cloudflare.F(cloudflare.LoadBalancerPoolNewParamsLoadShedding{
			DefaultPercent: cloudflare.F(0.000000),
			DefaultPolicy:  cloudflare.F(cloudflare.LoadBalancerPoolNewParamsLoadSheddingDefaultPolicyRandom),
			SessionPercent: cloudflare.F(0.000000),
			SessionPolicy:  cloudflare.F(cloudflare.LoadBalancerPoolNewParamsLoadSheddingSessionPolicyHash),
		}),
		Longitude:         cloudflare.F(0.000000),
		MinimumOrigins:    cloudflare.F(int64(0)),
		Monitor:           cloudflare.F[any](map[string]interface{}{}),
		NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
		NotificationFilter: cloudflare.F(cloudflare.LoadBalancerPoolNewParamsNotificationFilter{
			Origin: cloudflare.F(cloudflare.LoadBalancerPoolNewParamsNotificationFilterOrigin{
				Disable: cloudflare.F(true),
				Healthy: cloudflare.F(true),
			}),
			Pool: cloudflare.F(cloudflare.LoadBalancerPoolNewParamsNotificationFilterPool{
				Disable: cloudflare.F(true),
				Healthy: cloudflare.F(false),
			}),
		}),
		OriginSteering: cloudflare.F(cloudflare.LoadBalancerPoolNewParamsOriginSteering{
			Policy: cloudflare.F(cloudflare.LoadBalancerPoolNewParamsOriginSteeringPolicyRandom),
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

func TestLoadBalancerPoolUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Pools.Update(
		context.TODO(),
		"17b5962d775c646f3f9725cbc7a53df4",
		cloudflare.LoadBalancerPoolUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:      cloudflare.F("primary-dc-1"),
			Origins: cloudflare.F([]cloudflare.LoadBalancerPoolUpdateParamsOrigin{{
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.LoadBalancerPoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.LoadBalancerPoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.LoadBalancerPoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}}),
			CheckRegions: cloudflare.F([]cloudflare.LoadBalancerPoolUpdateParamsCheckRegion{cloudflare.LoadBalancerPoolUpdateParamsCheckRegionWeu, cloudflare.LoadBalancerPoolUpdateParamsCheckRegionEnam}),
			Description:  cloudflare.F("Primary data center - Provider XYZ"),
			Enabled:      cloudflare.F(false),
			Latitude:     cloudflare.F(0.000000),
			LoadShedding: cloudflare.F(cloudflare.LoadBalancerPoolUpdateParamsLoadShedding{
				DefaultPercent: cloudflare.F(0.000000),
				DefaultPolicy:  cloudflare.F(cloudflare.LoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicyRandom),
				SessionPercent: cloudflare.F(0.000000),
				SessionPolicy:  cloudflare.F(cloudflare.LoadBalancerPoolUpdateParamsLoadSheddingSessionPolicyHash),
			}),
			Longitude:         cloudflare.F(0.000000),
			MinimumOrigins:    cloudflare.F(int64(0)),
			Monitor:           cloudflare.F[any](map[string]interface{}{}),
			NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cloudflare.F(cloudflare.LoadBalancerPoolUpdateParamsNotificationFilter{
				Origin: cloudflare.F(cloudflare.LoadBalancerPoolUpdateParamsNotificationFilterOrigin{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(true),
				}),
				Pool: cloudflare.F(cloudflare.LoadBalancerPoolUpdateParamsNotificationFilterPool{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(false),
				}),
			}),
			OriginSteering: cloudflare.F(cloudflare.LoadBalancerPoolUpdateParamsOriginSteering{
				Policy: cloudflare.F(cloudflare.LoadBalancerPoolUpdateParamsOriginSteeringPolicyRandom),
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

func TestLoadBalancerPoolListWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Pools.List(context.TODO(), cloudflare.LoadBalancerPoolListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Monitor:   cloudflare.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerPoolDelete(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Pools.Delete(
		context.TODO(),
		"17b5962d775c646f3f9725cbc7a53df4",
		cloudflare.LoadBalancerPoolDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestLoadBalancerPoolEditWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Pools.Edit(
		context.TODO(),
		"17b5962d775c646f3f9725cbc7a53df4",
		cloudflare.LoadBalancerPoolEditParams{
			AccountID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			CheckRegions: cloudflare.F([]cloudflare.LoadBalancerPoolEditParamsCheckRegion{cloudflare.LoadBalancerPoolEditParamsCheckRegionWeu, cloudflare.LoadBalancerPoolEditParamsCheckRegionEnam}),
			Description:  cloudflare.F("Primary data center - Provider XYZ"),
			Enabled:      cloudflare.F(false),
			Latitude:     cloudflare.F(0.000000),
			LoadShedding: cloudflare.F(cloudflare.LoadBalancerPoolEditParamsLoadShedding{
				DefaultPercent: cloudflare.F(0.000000),
				DefaultPolicy:  cloudflare.F(cloudflare.LoadBalancerPoolEditParamsLoadSheddingDefaultPolicyRandom),
				SessionPercent: cloudflare.F(0.000000),
				SessionPolicy:  cloudflare.F(cloudflare.LoadBalancerPoolEditParamsLoadSheddingSessionPolicyHash),
			}),
			Longitude:         cloudflare.F(0.000000),
			MinimumOrigins:    cloudflare.F(int64(0)),
			Monitor:           cloudflare.F[any](map[string]interface{}{}),
			Name:              cloudflare.F("primary-dc-1"),
			NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cloudflare.F(cloudflare.LoadBalancerPoolEditParamsNotificationFilter{
				Origin: cloudflare.F(cloudflare.LoadBalancerPoolEditParamsNotificationFilterOrigin{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(true),
				}),
				Pool: cloudflare.F(cloudflare.LoadBalancerPoolEditParamsNotificationFilterPool{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(false),
				}),
			}),
			OriginSteering: cloudflare.F(cloudflare.LoadBalancerPoolEditParamsOriginSteering{
				Policy: cloudflare.F(cloudflare.LoadBalancerPoolEditParamsOriginSteeringPolicyRandom),
			}),
			Origins: cloudflare.F([]cloudflare.LoadBalancerPoolEditParamsOrigin{{
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.LoadBalancerPoolEditParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.LoadBalancerPoolEditParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.LoadBalancerPoolEditParamsOriginsHeader{
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

func TestLoadBalancerPoolGet(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Pools.Get(
		context.TODO(),
		"17b5962d775c646f3f9725cbc7a53df4",
		cloudflare.LoadBalancerPoolGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
