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

func TestResourceSharingNew(t *testing.T) {
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
	_, err := client.ResourceSharing.New(context.TODO(), cloudflare.ResourceSharingNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Name:      cloudflare.F("My Shared WAF Managed Rule"),
		Recipients: cloudflare.F([]cloudflare.ResourceSharingNewParamsRecipient{{
			AccountID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			OrganizationID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		}, {
			AccountID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			OrganizationID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		}, {
			AccountID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			OrganizationID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		}}),
		Resources: cloudflare.F([]cloudflare.ResourceSharingNewParamsResource{{
			Meta:              cloudflare.F[any](map[string]interface{}{}),
			ResourceAccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ResourceID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ResourceType:      cloudflare.F(cloudflare.ResourceSharingNewParamsResourcesResourceTypeCustomRuleset),
		}, {
			Meta:              cloudflare.F[any](map[string]interface{}{}),
			ResourceAccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ResourceID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ResourceType:      cloudflare.F(cloudflare.ResourceSharingNewParamsResourcesResourceTypeCustomRuleset),
		}, {
			Meta:              cloudflare.F[any](map[string]interface{}{}),
			ResourceAccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ResourceID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ResourceType:      cloudflare.F(cloudflare.ResourceSharingNewParamsResourcesResourceTypeCustomRuleset),
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

func TestResourceSharingUpdate(t *testing.T) {
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
	_, err := client.ResourceSharing.Update(
		context.TODO(),
		"3fd85f74b32742f1bff64a85009dda07",
		cloudflare.ResourceSharingUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:      cloudflare.F("My Shared WAF Managed Rule"),
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

func TestResourceSharingListWithOptionalParams(t *testing.T) {
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
	_, err := client.ResourceSharing.List(context.TODO(), cloudflare.ResourceSharingListParams{
		AccountID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction:  cloudflare.F(cloudflare.ResourceSharingListParamsDirectionAsc),
		Kind:       cloudflare.F(cloudflare.ResourceSharingListParamsKindSent),
		Order:      cloudflare.F(cloudflare.ResourceSharingListParamsOrderName),
		Page:       cloudflare.F(int64(2)),
		PerPage:    cloudflare.F(int64(20)),
		Status:     cloudflare.F(cloudflare.ResourceSharingListParamsStatusActive),
		TargetType: cloudflare.F(cloudflare.ResourceSharingListParamsTargetTypeAccount),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResourceSharingDelete(t *testing.T) {
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
	_, err := client.ResourceSharing.Delete(
		context.TODO(),
		"3fd85f74b32742f1bff64a85009dda07",
		cloudflare.ResourceSharingDeleteParams{
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

func TestResourceSharingGet(t *testing.T) {
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
	_, err := client.ResourceSharing.Get(
		context.TODO(),
		"3fd85f74b32742f1bff64a85009dda07",
		cloudflare.ResourceSharingGetParams{
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
