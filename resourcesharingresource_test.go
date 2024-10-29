// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v3/option"
)

func TestResourceSharingResourceNew(t *testing.T) {
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
	_, err := client.ResourceSharing.Resources.New(
		context.TODO(),
		"3fd85f74b32742f1bff64a85009dda07",
		cloudflare.ResourceSharingResourceNewParams{
			AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Meta:              cloudflare.F[any](map[string]interface{}{}),
			ResourceAccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ResourceID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ResourceType:      cloudflare.F(cloudflare.ResourceSharingResourceNewParamsResourceTypeCustomRuleset),
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

func TestResourceSharingResourceUpdate(t *testing.T) {
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
	_, err := client.ResourceSharing.Resources.Update(
		context.TODO(),
		"3fd85f74b32742f1bff64a85009dda07",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ResourceSharingResourceUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Meta:      cloudflare.F[any](map[string]interface{}{}),
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

func TestResourceSharingResourceListWithOptionalParams(t *testing.T) {
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
	_, err := client.ResourceSharing.Resources.List(
		context.TODO(),
		"3fd85f74b32742f1bff64a85009dda07",
		cloudflare.ResourceSharingResourceListParams{
			AccountID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Page:         cloudflare.F(int64(2)),
			PerPage:      cloudflare.F(int64(20)),
			ResourceType: cloudflare.F(cloudflare.ResourceSharingResourceListParamsResourceTypeCustomRuleset),
			Status:       cloudflare.F(cloudflare.ResourceSharingResourceListParamsStatusActive),
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

func TestResourceSharingResourceDelete(t *testing.T) {
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
	_, err := client.ResourceSharing.Resources.Delete(
		context.TODO(),
		"3fd85f74b32742f1bff64a85009dda07",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ResourceSharingResourceDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestResourceSharingResourceGet(t *testing.T) {
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
	_, err := client.ResourceSharing.Resources.Get(
		context.TODO(),
		"3fd85f74b32742f1bff64a85009dda07",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ResourceSharingResourceGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
