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

func TestZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Zones.APIGateways.Configurations.APIShieldSettingsGetInformationAboutSpecificConfigurationProperties(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParams{
			Properties: cloudflare.F([]cloudflare.ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParamsProperty{cloudflare.ZoneAPIGatewayConfigurationAPIShieldSettingsGetInformationAboutSpecificConfigurationPropertiesParamsPropertyAuthIDCharacteristics}),
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

func TestZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Zones.APIGateways.Configurations.APIShieldSettingsSetConfigurationProperties(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParams{
			AuthIDCharacteristics: cloudflare.F([]cloudflare.ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristic{{
				Name: cloudflare.F("authorization"),
				Type: cloudflare.F(cloudflare.ZoneAPIGatewayConfigurationAPIShieldSettingsSetConfigurationPropertiesParamsAuthIDCharacteristicsTypeHeader),
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
