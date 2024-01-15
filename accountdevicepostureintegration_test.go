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

func TestAccountDevicePostureIntegrationGet(t *testing.T) {
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
	_, err := client.Accounts.Devices.Postures.Integrations.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDevicePostureIntegrationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Devices.Postures.Integrations.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccountDevicePostureIntegrationUpdateParams{
			Config: cloudflare.F[cloudflare.AccountDevicePostureIntegrationUpdateParamsConfig](cloudflare.AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhWorkspaceOneConfigRequest(cloudflare.AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhWorkspaceOneConfigRequest{
				APIURL:       cloudflare.F("https://as123.awmdm.com/API"),
				AuthURL:      cloudflare.F("https://na.uemauth.vmwservices.com/connect/token"),
				ClientID:     cloudflare.F("example client id"),
				ClientSecret: cloudflare.F("example client secret"),
			})),
			Interval: cloudflare.F("10m"),
			Name:     cloudflare.F("My Workspace One Integration"),
			Type:     cloudflare.F(cloudflare.AccountDevicePostureIntegrationUpdateParamsTypeWorkspaceOne),
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

func TestAccountDevicePostureIntegrationDelete(t *testing.T) {
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
	_, err := client.Accounts.Devices.Postures.Integrations.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Devices.Postures.Integrations.DevicePostureIntegrationsNewDevicePostureIntegration(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParams{
			Config: cloudflare.F[cloudflare.AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig](cloudflare.AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhWorkspaceOneConfigRequest(cloudflare.AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhWorkspaceOneConfigRequest{
				APIURL:       cloudflare.F("https://as123.awmdm.com/API"),
				AuthURL:      cloudflare.F("https://na.uemauth.vmwservices.com/connect/token"),
				ClientID:     cloudflare.F("example client id"),
				ClientSecret: cloudflare.F("example client secret"),
			})),
			Interval: cloudflare.F("10m"),
			Name:     cloudflare.F("My Workspace One Integration"),
			Type:     cloudflare.F(cloudflare.AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeWorkspaceOne),
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

func TestAccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrations(t *testing.T) {
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
	_, err := client.Accounts.Devices.Postures.Integrations.DevicePostureIntegrationsListDevicePostureIntegrations(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
