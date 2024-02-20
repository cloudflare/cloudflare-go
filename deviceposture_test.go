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

func TestDevicePostureNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Devices.Postures.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.DevicePostureNewParams{
			Name:        cloudflare.F("Admin Serial Numbers"),
			Type:        cloudflare.F(cloudflare.DevicePostureNewParamsTypeFile),
			Description: cloudflare.F("The rule for admin serial numbers"),
			Expiration:  cloudflare.F("1h"),
			Input: cloudflare.F[cloudflare.DevicePostureNewParamsInput](cloudflare.DevicePostureNewParamsInputTeamsDevicesFileInputRequest(cloudflare.DevicePostureNewParamsInputTeamsDevicesFileInputRequest{
				Exists:          cloudflare.F(true),
				OperatingSystem: cloudflare.F(cloudflare.DevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux),
				Path:            cloudflare.F("/bin/cat"),
				Sha256:          cloudflare.F("https://api.us-2.crowdstrike.com"),
				Thumbprint:      cloudflare.F("0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"),
			})),
			Match: cloudflare.F([]cloudflare.DevicePostureNewParamsMatch{{
				Platform: cloudflare.F(cloudflare.DevicePostureNewParamsMatchPlatformWindows),
			}, {
				Platform: cloudflare.F(cloudflare.DevicePostureNewParamsMatchPlatformWindows),
			}, {
				Platform: cloudflare.F(cloudflare.DevicePostureNewParamsMatchPlatformWindows),
			}}),
			Schedule: cloudflare.F("1h"),
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

func TestDevicePostureUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Devices.Postures.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.DevicePostureUpdateParams{
			Name:        cloudflare.F("Admin Serial Numbers"),
			Type:        cloudflare.F(cloudflare.DevicePostureUpdateParamsTypeFile),
			Description: cloudflare.F("The rule for admin serial numbers"),
			Expiration:  cloudflare.F("1h"),
			Input: cloudflare.F[cloudflare.DevicePostureUpdateParamsInput](cloudflare.DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest(cloudflare.DevicePostureUpdateParamsInputTeamsDevicesFileInputRequest{
				Exists:          cloudflare.F(true),
				OperatingSystem: cloudflare.F(cloudflare.DevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux),
				Path:            cloudflare.F("/bin/cat"),
				Sha256:          cloudflare.F("https://api.us-2.crowdstrike.com"),
				Thumbprint:      cloudflare.F("0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"),
			})),
			Match: cloudflare.F([]cloudflare.DevicePostureUpdateParamsMatch{{
				Platform: cloudflare.F(cloudflare.DevicePostureUpdateParamsMatchPlatformWindows),
			}, {
				Platform: cloudflare.F(cloudflare.DevicePostureUpdateParamsMatchPlatformWindows),
			}, {
				Platform: cloudflare.F(cloudflare.DevicePostureUpdateParamsMatchPlatformWindows),
			}}),
			Schedule: cloudflare.F("1h"),
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

func TestDevicePostureList(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Devices.Postures.List(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDevicePostureDelete(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Devices.Postures.Delete(
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

func TestDevicePostureGet(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Devices.Postures.Get(
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
