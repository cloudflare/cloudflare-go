// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
)

func TestDeviceDEXTestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.DEXTests.New(context.TODO(), zero_trust.DeviceDEXTestNewParams{
		AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
		Data: cloudflare.F(zero_trust.SchemaDataParam{
			Host:   cloudflare.F("https://dash.cloudflare.com"),
			Kind:   cloudflare.F("http"),
			Method: cloudflare.F("GET"),
		}),
		Enabled:     cloudflare.F(true),
		Interval:    cloudflare.F("30m"),
		Name:        cloudflare.F("HTTP dash health check"),
		Description: cloudflare.F("Checks the dash endpoint every 30 minutes"),
		TargetPolicies: cloudflare.F([]zero_trust.DeviceDEXTestNewParamsTargetPolicy{{
			Default: cloudflare.F(true),
			ID:      cloudflare.F("string"),
			Name:    cloudflare.F("string"),
		}, {
			Default: cloudflare.F(true),
			ID:      cloudflare.F("string"),
			Name:    cloudflare.F("string"),
		}, {
			Default: cloudflare.F(true),
			ID:      cloudflare.F("string"),
			Name:    cloudflare.F("string"),
		}}),
		Targeted: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceDEXTestUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.DEXTests.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DeviceDEXTestUpdateParams{
			AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
			Data: cloudflare.F(zero_trust.SchemaDataParam{
				Host:   cloudflare.F("https://dash.cloudflare.com"),
				Kind:   cloudflare.F("http"),
				Method: cloudflare.F("GET"),
			}),
			Enabled:     cloudflare.F(true),
			Interval:    cloudflare.F("30m"),
			Name:        cloudflare.F("HTTP dash health check"),
			Description: cloudflare.F("Checks the dash endpoint every 30 minutes"),
			TargetPolicies: cloudflare.F([]zero_trust.DeviceDEXTestUpdateParamsTargetPolicy{{
				Default: cloudflare.F(true),
				ID:      cloudflare.F("string"),
				Name:    cloudflare.F("string"),
			}, {
				Default: cloudflare.F(true),
				ID:      cloudflare.F("string"),
				Name:    cloudflare.F("string"),
			}, {
				Default: cloudflare.F(true),
				ID:      cloudflare.F("string"),
				Name:    cloudflare.F("string"),
			}}),
			Targeted: cloudflare.F(true),
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

func TestDeviceDEXTestList(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.DEXTests.List(context.TODO(), zero_trust.DeviceDEXTestListParams{
		AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceDEXTestDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.DEXTests.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DeviceDEXTestDeleteParams{
			AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
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

func TestDeviceDEXTestGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.DEXTests.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DeviceDEXTestGetParams{
			AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
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
