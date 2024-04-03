// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/load_balancers"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestPoolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Pools.New(context.TODO(), load_balancers.PoolNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Name:      cloudflare.F("primary-dc-1"),
		Origins: cloudflare.F([]load_balancers.PoolNewParamsOrigin{{
			Address: cloudflare.F("0.0.0.0"),
			Enabled: cloudflare.F(true),
			Header: cloudflare.F(load_balancers.PoolNewParamsOriginsHeader{
				Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			}),
			Name:             cloudflare.F("app-server-1"),
			VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           cloudflare.F(0.600000),
		}, {
			Address: cloudflare.F("0.0.0.0"),
			Enabled: cloudflare.F(true),
			Header: cloudflare.F(load_balancers.PoolNewParamsOriginsHeader{
				Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			}),
			Name:             cloudflare.F("app-server-1"),
			VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           cloudflare.F(0.600000),
		}, {
			Address: cloudflare.F("0.0.0.0"),
			Enabled: cloudflare.F(true),
			Header: cloudflare.F(load_balancers.PoolNewParamsOriginsHeader{
				Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			}),
			Name:             cloudflare.F("app-server-1"),
			VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           cloudflare.F(0.600000),
		}}),
		Description: cloudflare.F("Primary data center - Provider XYZ"),
		Enabled:     cloudflare.F(false),
		Latitude:    cloudflare.F(0.000000),
		LoadShedding: cloudflare.F(load_balancers.PoolNewParamsLoadShedding{
			DefaultPercent: cloudflare.F(0.000000),
			DefaultPolicy:  cloudflare.F(load_balancers.PoolNewParamsLoadSheddingDefaultPolicyRandom),
			SessionPercent: cloudflare.F(0.000000),
			SessionPolicy:  cloudflare.F(load_balancers.PoolNewParamsLoadSheddingSessionPolicyHash),
		}),
		Longitude:         cloudflare.F(0.000000),
		MinimumOrigins:    cloudflare.F(int64(0)),
		Monitor:           cloudflare.F[any](map[string]interface{}{}),
		NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
		NotificationFilter: cloudflare.F(load_balancers.PoolNewParamsNotificationFilter{
			Origin: cloudflare.F(load_balancers.PoolNewParamsNotificationFilterOrigin{
				Disable: cloudflare.F(true),
				Healthy: cloudflare.F(true),
			}),
			Pool: cloudflare.F(load_balancers.PoolNewParamsNotificationFilterPool{
				Disable: cloudflare.F(true),
				Healthy: cloudflare.F(false),
			}),
		}),
		OriginSteering: cloudflare.F(load_balancers.PoolNewParamsOriginSteering{
			Policy: cloudflare.F(load_balancers.PoolNewParamsOriginSteeringPolicyRandom),
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

func TestPoolUpdateWithOptionalParams(t *testing.T) {
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
		load_balancers.PoolUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:      cloudflare.F("primary-dc-1"),
			Origins: cloudflare.F([]load_balancers.PoolUpdateParamsOrigin{{
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(load_balancers.PoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(load_balancers.PoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(load_balancers.PoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}}),
			CheckRegions: cloudflare.F([]load_balancers.PoolUpdateParamsCheckRegion{load_balancers.PoolUpdateParamsCheckRegionWeu, load_balancers.PoolUpdateParamsCheckRegionEnam}),
			Description:  cloudflare.F("Primary data center - Provider XYZ"),
			Enabled:      cloudflare.F(false),
			Latitude:     cloudflare.F(0.000000),
			LoadShedding: cloudflare.F(load_balancers.PoolUpdateParamsLoadShedding{
				DefaultPercent: cloudflare.F(0.000000),
				DefaultPolicy:  cloudflare.F(load_balancers.PoolUpdateParamsLoadSheddingDefaultPolicyRandom),
				SessionPercent: cloudflare.F(0.000000),
				SessionPolicy:  cloudflare.F(load_balancers.PoolUpdateParamsLoadSheddingSessionPolicyHash),
			}),
			Longitude:         cloudflare.F(0.000000),
			MinimumOrigins:    cloudflare.F(int64(0)),
			Monitor:           cloudflare.F[any](map[string]interface{}{}),
			NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cloudflare.F(load_balancers.PoolUpdateParamsNotificationFilter{
				Origin: cloudflare.F(load_balancers.PoolUpdateParamsNotificationFilterOrigin{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(true),
				}),
				Pool: cloudflare.F(load_balancers.PoolUpdateParamsNotificationFilterPool{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(false),
				}),
			}),
			OriginSteering: cloudflare.F(load_balancers.PoolUpdateParamsOriginSteering{
				Policy: cloudflare.F(load_balancers.PoolUpdateParamsOriginSteeringPolicyRandom),
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

func TestPoolListWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Pools.List(context.TODO(), load_balancers.PoolListParams{
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

func TestPoolDelete(t *testing.T) {
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
		load_balancers.PoolDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body:      cloudflare.F[any](map[string]interface{}{}),
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

func TestPoolEditWithOptionalParams(t *testing.T) {
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
		load_balancers.PoolEditParams{
			AccountID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			CheckRegions: cloudflare.F([]load_balancers.PoolEditParamsCheckRegion{load_balancers.PoolEditParamsCheckRegionWeu, load_balancers.PoolEditParamsCheckRegionEnam}),
			Description:  cloudflare.F("Primary data center - Provider XYZ"),
			Enabled:      cloudflare.F(false),
			Latitude:     cloudflare.F(0.000000),
			LoadShedding: cloudflare.F(load_balancers.PoolEditParamsLoadShedding{
				DefaultPercent: cloudflare.F(0.000000),
				DefaultPolicy:  cloudflare.F(load_balancers.PoolEditParamsLoadSheddingDefaultPolicyRandom),
				SessionPercent: cloudflare.F(0.000000),
				SessionPolicy:  cloudflare.F(load_balancers.PoolEditParamsLoadSheddingSessionPolicyHash),
			}),
			Longitude:         cloudflare.F(0.000000),
			MinimumOrigins:    cloudflare.F(int64(0)),
			Monitor:           cloudflare.F[any](map[string]interface{}{}),
			Name:              cloudflare.F("primary-dc-1"),
			NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cloudflare.F(load_balancers.PoolEditParamsNotificationFilter{
				Origin: cloudflare.F(load_balancers.PoolEditParamsNotificationFilterOrigin{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(true),
				}),
				Pool: cloudflare.F(load_balancers.PoolEditParamsNotificationFilterPool{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(false),
				}),
			}),
			OriginSteering: cloudflare.F(load_balancers.PoolEditParamsOriginSteering{
				Policy: cloudflare.F(load_balancers.PoolEditParamsOriginSteeringPolicyRandom),
			}),
			Origins: cloudflare.F([]load_balancers.PoolEditParamsOrigin{{
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(load_balancers.PoolEditParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(load_balancers.PoolEditParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(load_balancers.PoolEditParamsOriginsHeader{
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

func TestPoolGet(t *testing.T) {
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
		load_balancers.PoolGetParams{
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
