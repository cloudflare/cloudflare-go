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

func TestAccountDevicePostureGet(t *testing.T) {
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
	_, err := client.Accounts.Devices.Postures.Get(
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

func TestAccountDevicePostureUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Devices.Postures.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccountDevicePostureUpdateParams{
			Name:        cloudflare.F("Admin Serial Numbers"),
			Type:        cloudflare.F(cloudflare.AccountDevicePostureUpdateParamsTypeFile),
			Description: cloudflare.F("The rule for admin serial numbers"),
			Expiration:  cloudflare.F("1h"),
			Input: cloudflare.F[cloudflare.AccountDevicePostureUpdateParamsInput](cloudflare.AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequest(cloudflare.AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequest{
				Exists:          cloudflare.F(true),
				OperatingSystem: cloudflare.F(cloudflare.AccountDevicePostureUpdateParamsInputZR7Sv6YhFileInputRequestOperatingSystemMac),
				Path:            cloudflare.F("/bin/cat"),
				Sha256:          cloudflare.F("https://api.us-2.crowdstrike.com"),
				Thumbprint:      cloudflare.F("0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"),
			})),
			Match: cloudflare.F([]cloudflare.AccountDevicePostureUpdateParamsMatch{{
				Platform: cloudflare.F(cloudflare.AccountDevicePostureUpdateParamsMatchPlatformWindows),
			}, {
				Platform: cloudflare.F(cloudflare.AccountDevicePostureUpdateParamsMatchPlatformWindows),
			}, {
				Platform: cloudflare.F(cloudflare.AccountDevicePostureUpdateParamsMatchPlatformWindows),
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

func TestAccountDevicePostureDelete(t *testing.T) {
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
	_, err := client.Accounts.Devices.Postures.Delete(
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

func TestAccountDevicePostureDevicePostureRulesNewDevicePostureRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Devices.Postures.DevicePostureRulesNewDevicePostureRule(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParams{
			Name:        cloudflare.F("Admin Serial Numbers"),
			Type:        cloudflare.F(cloudflare.AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeFile),
			Description: cloudflare.F("The rule for admin serial numbers"),
			Expiration:  cloudflare.F("1h"),
			Input: cloudflare.F[cloudflare.AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput](cloudflare.AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequest(cloudflare.AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequest{
				Exists:          cloudflare.F(true),
				OperatingSystem: cloudflare.F(cloudflare.AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInputZR7Sv6YhFileInputRequestOperatingSystemMac),
				Path:            cloudflare.F("/bin/cat"),
				Sha256:          cloudflare.F("https://api.us-2.crowdstrike.com"),
				Thumbprint:      cloudflare.F("0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"),
			})),
			Match: cloudflare.F([]cloudflare.AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatch{{
				Platform: cloudflare.F(cloudflare.AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformWindows),
			}, {
				Platform: cloudflare.F(cloudflare.AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformWindows),
			}, {
				Platform: cloudflare.F(cloudflare.AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformWindows),
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

func TestAccountDevicePostureDevicePostureRulesListDevicePostureRules(t *testing.T) {
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
	_, err := client.Accounts.Devices.Postures.DevicePostureRulesListDevicePostureRules(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
