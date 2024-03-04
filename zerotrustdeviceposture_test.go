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

func TestZeroTrustDevicePostureNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Postures.New(context.TODO(), cloudflare.ZeroTrustDevicePostureNewParams{
		AccountID:   cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
		Name:        cloudflare.F("Admin Serial Numbers"),
		Type:        cloudflare.F(cloudflare.ZeroTrustDevicePostureNewParamsTypeFile),
		Description: cloudflare.F("The rule for admin serial numbers"),
		Expiration:  cloudflare.F("1h"),
		Input: cloudflare.F[cloudflare.ZeroTrustDevicePostureNewParamsInput](cloudflare.ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequest(cloudflare.ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequest{
			Exists:          cloudflare.F(true),
			OperatingSystem: cloudflare.F(cloudflare.ZeroTrustDevicePostureNewParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux),
			Path:            cloudflare.F("/bin/cat"),
			Sha256:          cloudflare.F("https://api.us-2.crowdstrike.com"),
			Thumbprint:      cloudflare.F("0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"),
		})),
		Match: cloudflare.F([]cloudflare.ZeroTrustDevicePostureNewParamsMatch{{
			Platform: cloudflare.F(cloudflare.ZeroTrustDevicePostureNewParamsMatchPlatformWindows),
		}, {
			Platform: cloudflare.F(cloudflare.ZeroTrustDevicePostureNewParamsMatchPlatformWindows),
		}, {
			Platform: cloudflare.F(cloudflare.ZeroTrustDevicePostureNewParamsMatchPlatformWindows),
		}}),
		Schedule: cloudflare.F("1h"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZeroTrustDevicePostureUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Postures.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDevicePostureUpdateParams{
			AccountID:   cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
			Name:        cloudflare.F("Admin Serial Numbers"),
			Type:        cloudflare.F(cloudflare.ZeroTrustDevicePostureUpdateParamsTypeFile),
			Description: cloudflare.F("The rule for admin serial numbers"),
			Expiration:  cloudflare.F("1h"),
			Input: cloudflare.F[cloudflare.ZeroTrustDevicePostureUpdateParamsInput](cloudflare.ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequest(cloudflare.ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequest{
				Exists:          cloudflare.F(true),
				OperatingSystem: cloudflare.F(cloudflare.ZeroTrustDevicePostureUpdateParamsInputTeamsDevicesFileInputRequestOperatingSystemLinux),
				Path:            cloudflare.F("/bin/cat"),
				Sha256:          cloudflare.F("https://api.us-2.crowdstrike.com"),
				Thumbprint:      cloudflare.F("0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"),
			})),
			Match: cloudflare.F([]cloudflare.ZeroTrustDevicePostureUpdateParamsMatch{{
				Platform: cloudflare.F(cloudflare.ZeroTrustDevicePostureUpdateParamsMatchPlatformWindows),
			}, {
				Platform: cloudflare.F(cloudflare.ZeroTrustDevicePostureUpdateParamsMatchPlatformWindows),
			}, {
				Platform: cloudflare.F(cloudflare.ZeroTrustDevicePostureUpdateParamsMatchPlatformWindows),
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

func TestZeroTrustDevicePostureList(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Postures.List(context.TODO(), cloudflare.ZeroTrustDevicePostureListParams{
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

func TestZeroTrustDevicePostureDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Postures.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDevicePostureDeleteParams{
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

func TestZeroTrustDevicePostureGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Postures.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDevicePostureGetParams{
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
