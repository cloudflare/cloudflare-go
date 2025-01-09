// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/addressing"
	"github.com/cloudflare/cloudflare-go/v3/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v3/option"
)

func TestAddressMapZoneUpdate(t *testing.T) {
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
	_, err := client.Addressing.AddressMaps.Zones.Update(
		context.TODO(),
		"055817b111884e0227e1be16a0be6ee0",
		addressing.AddressMapZoneUpdateParams{
			ZoneID:    cloudflare.F("8ac8489932db6327334c9b6d58544cfe"),
			AccountID: cloudflare.F("258def64c72dae45f3e4c8516e2111f2"),
			Body:      map[string]interface{}{},
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

func TestAddressMapZoneDelete(t *testing.T) {
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
	_, err := client.Addressing.AddressMaps.Zones.Delete(
		context.TODO(),
		"055817b111884e0227e1be16a0be6ee0",
		addressing.AddressMapZoneDeleteParams{
			ZoneID:    cloudflare.F("8ac8489932db6327334c9b6d58544cfe"),
			AccountID: cloudflare.F("258def64c72dae45f3e4c8516e2111f2"),
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
