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

func TestUserLoadBalancerMonitorGet(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.Get(context.TODO(), "f1aba936b94213e5b8dca0c0dbf1f9cc")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerMonitorUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.Update(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		cloudflare.UserLoadBalancerMonitorUpdateParams{
			AllowInsecure:   cloudflare.F(true),
			ConsecutiveDown: cloudflare.F(int64(0)),
			ConsecutiveUp:   cloudflare.F(int64(0)),
			Description:     cloudflare.F("Login page monitor"),
			ExpectedBody:    cloudflare.F("alive"),
			ExpectedCodes:   cloudflare.F("2xx"),
			FollowRedirects: cloudflare.F(true),
			Header: cloudflare.F[any](map[string]interface{}{
				"Host": map[string]interface{}{
					"0": "example.com",
				},
				"X-App-ID": map[string]interface{}{
					"0": "abc123",
				},
			}),
			Interval:  cloudflare.F(int64(0)),
			Method:    cloudflare.F("GET"),
			Path:      cloudflare.F("/health"),
			Port:      cloudflare.F(int64(0)),
			ProbeZone: cloudflare.F("example.com"),
			Retries:   cloudflare.F(int64(0)),
			Timeout:   cloudflare.F(int64(0)),
			Type:      cloudflare.F(cloudflare.UserLoadBalancerMonitorUpdateParamsTypeHTTPs),
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

func TestUserLoadBalancerMonitorDelete(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.Delete(context.TODO(), "f1aba936b94213e5b8dca0c0dbf1f9cc")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorWithOptionalParams(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.LoadBalancerMonitorsNewMonitor(context.TODO(), cloudflare.UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParams{
		AllowInsecure:   cloudflare.F(true),
		ConsecutiveDown: cloudflare.F(int64(0)),
		ConsecutiveUp:   cloudflare.F(int64(0)),
		Description:     cloudflare.F("Login page monitor"),
		ExpectedBody:    cloudflare.F("alive"),
		ExpectedCodes:   cloudflare.F("2xx"),
		FollowRedirects: cloudflare.F(true),
		Header: cloudflare.F[any](map[string]interface{}{
			"Host": map[string]interface{}{
				"0": "example.com",
			},
			"X-App-ID": map[string]interface{}{
				"0": "abc123",
			},
		}),
		Interval:  cloudflare.F(int64(0)),
		Method:    cloudflare.F("GET"),
		Path:      cloudflare.F("/health"),
		Port:      cloudflare.F(int64(0)),
		ProbeZone: cloudflare.F("example.com"),
		Retries:   cloudflare.F(int64(0)),
		Timeout:   cloudflare.F(int64(0)),
		Type:      cloudflare.F(cloudflare.UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsTypeHTTPs),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerMonitorLoadBalancerMonitorsListMonitors(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.LoadBalancerMonitorsListMonitors(context.TODO())
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerMonitorPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.Patch(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		cloudflare.UserLoadBalancerMonitorPatchParams{
			AllowInsecure:   cloudflare.F(true),
			ConsecutiveDown: cloudflare.F(int64(0)),
			ConsecutiveUp:   cloudflare.F(int64(0)),
			Description:     cloudflare.F("Login page monitor"),
			ExpectedBody:    cloudflare.F("alive"),
			ExpectedCodes:   cloudflare.F("2xx"),
			FollowRedirects: cloudflare.F(true),
			Header: cloudflare.F[any](map[string]interface{}{
				"Host": map[string]interface{}{
					"0": "example.com",
				},
				"X-App-ID": map[string]interface{}{
					"0": "abc123",
				},
			}),
			Interval:  cloudflare.F(int64(0)),
			Method:    cloudflare.F("GET"),
			Path:      cloudflare.F("/health"),
			Port:      cloudflare.F(int64(0)),
			ProbeZone: cloudflare.F("example.com"),
			Retries:   cloudflare.F(int64(0)),
			Timeout:   cloudflare.F(int64(0)),
			Type:      cloudflare.F(cloudflare.UserLoadBalancerMonitorPatchParamsTypeHTTPs),
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
