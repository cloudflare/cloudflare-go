// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/load_balancers"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestPoolNewWithOptionalParams(t *testing.T) {
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
		Origins: cloudflare.F([]load_balancers.OriginParam{{
			Address: cloudflare.F("0.0.0.0"),
			Enabled: cloudflare.F(true),
			Header: cloudflare.F(load_balancers.HeaderParam{
				Host: cloudflare.F([]load_balancers.HostParam{"example.com"}),
			}),
			Name:             cloudflare.F("app-server-1"),
			VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           cloudflare.F(0.600000),
		}}),
		Description: cloudflare.F("Primary data center - Provider XYZ"),
		Enabled:     cloudflare.F(false),
		Latitude:    cloudflare.F(0.000000),
		LoadShedding: cloudflare.F(load_balancers.LoadSheddingParam{
			DefaultPercent: cloudflare.F(0.000000),
			DefaultPolicy:  cloudflare.F(load_balancers.LoadSheddingDefaultPolicyRandom),
			SessionPercent: cloudflare.F(0.000000),
			SessionPolicy:  cloudflare.F(load_balancers.LoadSheddingSessionPolicyHash),
		}),
		Longitude:         cloudflare.F(0.000000),
		MinimumOrigins:    cloudflare.F(int64(0)),
		Monitor:           cloudflare.F("monitor"),
		NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
		NotificationFilter: cloudflare.F(load_balancers.NotificationFilterParam{
			Origin: cloudflare.F(load_balancers.FilterOptionsParam{
				Disable: cloudflare.F(true),
				Healthy: cloudflare.F(true),
			}),
			Pool: cloudflare.F(load_balancers.FilterOptionsParam{
				Disable: cloudflare.F(true),
				Healthy: cloudflare.F(false),
			}),
		}),
		OriginSteering: cloudflare.F(load_balancers.OriginSteeringParam{
			Policy: cloudflare.F(load_balancers.OriginSteeringPolicyRandom),
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
			Origins: cloudflare.F([]load_balancers.OriginParam{{
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(load_balancers.HeaderParam{
					Host: cloudflare.F([]load_balancers.HostParam{"example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}}),
			CheckRegions: cloudflare.F([]load_balancers.CheckRegion{load_balancers.CheckRegionWnam, load_balancers.CheckRegionEnam}),
			Description:  cloudflare.F("Primary data center - Provider XYZ"),
			Enabled:      cloudflare.F(false),
			Latitude:     cloudflare.F(0.000000),
			LoadShedding: cloudflare.F(load_balancers.LoadSheddingParam{
				DefaultPercent: cloudflare.F(0.000000),
				DefaultPolicy:  cloudflare.F(load_balancers.LoadSheddingDefaultPolicyRandom),
				SessionPercent: cloudflare.F(0.000000),
				SessionPolicy:  cloudflare.F(load_balancers.LoadSheddingSessionPolicyHash),
			}),
			Longitude:         cloudflare.F(0.000000),
			MinimumOrigins:    cloudflare.F(int64(0)),
			Monitor:           cloudflare.F("monitor"),
			NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cloudflare.F(load_balancers.NotificationFilterParam{
				Origin: cloudflare.F(load_balancers.FilterOptionsParam{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(true),
				}),
				Pool: cloudflare.F(load_balancers.FilterOptionsParam{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(false),
				}),
			}),
			OriginSteering: cloudflare.F(load_balancers.OriginSteeringParam{
				Policy: cloudflare.F(load_balancers.OriginSteeringPolicyRandom),
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
		Monitor:   cloudflare.F("monitor"),
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

func TestPoolBulkEditWithOptionalParams(t *testing.T) {
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
	_, err := client.LoadBalancers.Pools.BulkEdit(context.TODO(), load_balancers.PoolBulkEditParams{
		AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		NotificationEmail: cloudflare.F(load_balancers.PoolBulkEditParamsNotificationEmailEmpty),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPoolEditWithOptionalParams(t *testing.T) {
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
			CheckRegions: cloudflare.F([]load_balancers.CheckRegion{load_balancers.CheckRegionWnam, load_balancers.CheckRegionEnam}),
			Description:  cloudflare.F("Primary data center - Provider XYZ"),
			Enabled:      cloudflare.F(false),
			Latitude:     cloudflare.F(0.000000),
			LoadShedding: cloudflare.F(load_balancers.LoadSheddingParam{
				DefaultPercent: cloudflare.F(0.000000),
				DefaultPolicy:  cloudflare.F(load_balancers.LoadSheddingDefaultPolicyRandom),
				SessionPercent: cloudflare.F(0.000000),
				SessionPolicy:  cloudflare.F(load_balancers.LoadSheddingSessionPolicyHash),
			}),
			Longitude:         cloudflare.F(0.000000),
			MinimumOrigins:    cloudflare.F(int64(0)),
			Monitor:           cloudflare.F("monitor"),
			Name:              cloudflare.F("primary-dc-1"),
			NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cloudflare.F(load_balancers.NotificationFilterParam{
				Origin: cloudflare.F(load_balancers.FilterOptionsParam{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(true),
				}),
				Pool: cloudflare.F(load_balancers.FilterOptionsParam{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(false),
				}),
			}),
			OriginSteering: cloudflare.F(load_balancers.OriginSteeringParam{
				Policy: cloudflare.F(load_balancers.OriginSteeringPolicyRandom),
			}),
			Origins: cloudflare.F([]load_balancers.OriginParam{{
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(load_balancers.HeaderParam{
					Host: cloudflare.F([]load_balancers.HostParam{"example.com"}),
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
