// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/managed_transforms"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestManagedTransformList(t *testing.T) {
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
	_, err := client.ManagedTransforms.List(context.TODO(), managed_transforms.ManagedTransformListParams{
		ZoneID: cloudflare.F("9f1839b6152d298aca64c4e906b6d074"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestManagedTransformDelete(t *testing.T) {
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
	err := client.ManagedTransforms.Delete(context.TODO(), managed_transforms.ManagedTransformDeleteParams{
		ZoneID: cloudflare.F("9f1839b6152d298aca64c4e906b6d074"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestManagedTransformEdit(t *testing.T) {
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
	_, err := client.ManagedTransforms.Edit(context.TODO(), managed_transforms.ManagedTransformEditParams{
		ZoneID: cloudflare.F("9f1839b6152d298aca64c4e906b6d074"),
		ManagedRequestHeaders: cloudflare.F([]managed_transforms.ManagedTransformEditParamsManagedRequestHeader{{
			ID:      cloudflare.F("add_bot_protection_headers"),
			Enabled: cloudflare.F(true),
		}}),
		ManagedResponseHeaders: cloudflare.F([]managed_transforms.ManagedTransformEditParamsManagedResponseHeader{{
			ID:      cloudflare.F("add_security_headers"),
			Enabled: cloudflare.F(true),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
