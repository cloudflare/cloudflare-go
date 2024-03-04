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

func TestZeroTrustDeviceDEXTestNewWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.ZeroTrust.Devices.DEXTests.New(context.TODO(), cloudflare.ZeroTrustDeviceDEXTestNewParams{
		AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
		Data: cloudflare.F(cloudflare.ZeroTrustDeviceDEXTestNewParamsData{
			Host:   cloudflare.F("https://dash.cloudflare.com"),
			Kind:   cloudflare.F("http"),
			Method: cloudflare.F("GET"),
		}),
		Enabled:     cloudflare.F(true),
		Interval:    cloudflare.F("30m"),
		Name:        cloudflare.F("HTTP dash health check"),
		Description: cloudflare.F("Checks the dash endpoint every 30 minutes"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZeroTrustDeviceDEXTestUpdateWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.ZeroTrust.Devices.DEXTests.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDeviceDEXTestUpdateParams{
			AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
			Data: cloudflare.F(cloudflare.ZeroTrustDeviceDEXTestUpdateParamsData{
				Host:   cloudflare.F("https://dash.cloudflare.com"),
				Kind:   cloudflare.F("http"),
				Method: cloudflare.F("GET"),
			}),
			Enabled:     cloudflare.F(true),
			Interval:    cloudflare.F("30m"),
			Name:        cloudflare.F("HTTP dash health check"),
			Description: cloudflare.F("Checks the dash endpoint every 30 minutes"),
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

func TestZeroTrustDeviceDEXTestList(t *testing.T) {
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
	)
	_, err := client.ZeroTrust.Devices.DEXTests.List(context.TODO(), cloudflare.ZeroTrustDeviceDEXTestListParams{
		AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZeroTrustDeviceDEXTestDelete(t *testing.T) {
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
	)
	_, err := client.ZeroTrust.Devices.DEXTests.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDeviceDEXTestDeleteParams{
			AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
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

func TestZeroTrustDeviceDEXTestGet(t *testing.T) {
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
	)
	_, err := client.ZeroTrust.Devices.DEXTests.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDeviceDEXTestGetParams{
			AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
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
