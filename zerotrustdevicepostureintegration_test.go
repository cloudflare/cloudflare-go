// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestZeroTrustDevicePostureIntegrationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Postures.Integrations.New(context.TODO(), cloudflare.ZeroTrustDevicePostureIntegrationNewParams{
		AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
		Config: cloudflare.F[cloudflare.ZeroTrustDevicePostureIntegrationNewParamsConfig](cloudflare.ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest(cloudflare.ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest{
			APIURL:       cloudflare.F("https://as123.awmdm.com/API"),
			AuthURL:      cloudflare.F("https://na.uemauth.vmwservices.com/connect/token"),
			ClientID:     cloudflare.F("example client id"),
			ClientSecret: cloudflare.F("example client secret"),
		})),
		Interval: cloudflare.F("10m"),
		Name:     cloudflare.F("My Workspace One Integration"),
		Type:     cloudflare.F(cloudflare.ZeroTrustDevicePostureIntegrationNewParamsTypeWorkspaceOne),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZeroTrustDevicePostureIntegrationList(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Postures.Integrations.List(context.TODO(), cloudflare.ZeroTrustDevicePostureIntegrationListParams{
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

func TestZeroTrustDevicePostureIntegrationDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Postures.Integrations.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDevicePostureIntegrationDeleteParams{
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

func TestZeroTrustDevicePostureIntegrationEditWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Postures.Integrations.Edit(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDevicePostureIntegrationEditParams{
			AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
			Config: cloudflare.F[cloudflare.ZeroTrustDevicePostureIntegrationEditParamsConfig](cloudflare.ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest(cloudflare.ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest{
				APIURL:       cloudflare.F("https://as123.awmdm.com/API"),
				AuthURL:      cloudflare.F("https://na.uemauth.vmwservices.com/connect/token"),
				ClientID:     cloudflare.F("example client id"),
				ClientSecret: cloudflare.F("example client secret"),
			})),
			Interval: cloudflare.F("10m"),
			Name:     cloudflare.F("My Workspace One Integration"),
			Type:     cloudflare.F(cloudflare.ZeroTrustDevicePostureIntegrationEditParamsTypeWorkspaceOne),
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

func TestZeroTrustDevicePostureIntegrationGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Postures.Integrations.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDevicePostureIntegrationGetParams{
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
