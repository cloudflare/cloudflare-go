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

func TestZeroTrustNetworkVirtualNetworkNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.VirtualNetworks.New(context.TODO(), cloudflare.ZeroTrustNetworkVirtualNetworkNewParams{
		AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
		Name:      cloudflare.F("us-east-1-vpc"),
		Comment:   cloudflare.F("Staging VPC for data science"),
		IsDefault: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZeroTrustNetworkVirtualNetworkListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.VirtualNetworks.List(context.TODO(), cloudflare.ZeroTrustNetworkVirtualNetworkListParams{
		AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
		IsDefault: cloudflare.F[any](map[string]interface{}{}),
		IsDeleted: cloudflare.F[any](map[string]interface{}{}),
		Name:      cloudflare.F("us-east-1-vpc"),
		VnetName:  cloudflare.F("us-east-1-vpc"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZeroTrustNetworkVirtualNetworkDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.VirtualNetworks.Delete(
		context.TODO(),
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustNetworkVirtualNetworkDeleteParams{
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

func TestZeroTrustNetworkVirtualNetworkEditWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.VirtualNetworks.Edit(
		context.TODO(),
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustNetworkVirtualNetworkEditParams{
			AccountID:        cloudflare.F("699d98642c564d2e855e9661899b7252"),
			Comment:          cloudflare.F("Staging VPC for data science"),
			IsDefaultNetwork: cloudflare.F(true),
			Name:             cloudflare.F("us-east-1-vpc"),
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
