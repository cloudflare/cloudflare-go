// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/magic_transit"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestAppNew(t *testing.T) {
	t.Skip("prism errors - https://github.com/cloudflare/cloudflare-python/actions/runs/9360388260/job/25765690361?pr=482#step:5:7212")
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
	_, err := client.MagicTransit.Apps.New(context.TODO(), magic_transit.AppNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Name:      cloudflare.F("Cloudflare Dashboard"),
		Type:      cloudflare.F("Development"),
		Hostnames: cloudflare.F([]string{"auth.cloudflare.com"}),
		IPSubnets: cloudflare.F([]string{"1.1.1.1/32"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppUpdate(t *testing.T) {
	t.Skip("prism errors - https://github.com/cloudflare/cloudflare-python/actions/runs/9360388260/job/25765690361?pr=482#step:5:7212")
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
	_, err := client.MagicTransit.Apps.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.AppUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Hostnames: cloudflare.F([]string{"auth.cloudflare.com"}),
			IPSubnets: cloudflare.F([]string{"1.1.1.1/32"}),
			Name:      cloudflare.F("Cloudflare Dashboard"),
			Type:      cloudflare.F("Development"),
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

func TestAppList(t *testing.T) {
	t.Skip("prism errors - https://github.com/cloudflare/cloudflare-python/actions/runs/9360388260/job/25765690361?pr=482#step:5:7212")
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
	_, err := client.MagicTransit.Apps.List(context.TODO(), magic_transit.AppListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppDelete(t *testing.T) {
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
	_, err := client.MagicTransit.Apps.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.AppDeleteParams{
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
