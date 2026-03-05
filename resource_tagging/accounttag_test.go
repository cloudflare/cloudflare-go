// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_tagging_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/resource_tagging"
)

func TestAccountTagUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ResourceTagging.AccountTags.Update(context.TODO(), resource_tagging.AccountTagUpdateParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body: resource_tagging.AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersion{
			ResourceID:   cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ResourceType: cloudflare.F(resource_tagging.AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeWorker),
			WorkerID:     cloudflare.F("3f72a691-44b3-4c11-8642-c18a88ddaa5e"),
			Tags: cloudflare.F(map[string]string{
				"environment": "production",
				"team":        "engineering",
			}),
		},
		IfMatch: cloudflare.F(`"v1:RBNvo1WzZ4oRRq0W9-hkng"`),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountTagDeleteWithOptionalParams(t *testing.T) {
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
	err := client.ResourceTagging.AccountTags.Delete(context.TODO(), resource_tagging.AccountTagDeleteParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		IfMatch:   cloudflare.F(`"v1:RBNvo1WzZ4oRRq0W9-hkng"`),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountTagGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ResourceTagging.AccountTags.Get(context.TODO(), resource_tagging.AccountTagGetParams{
		AccountID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ResourceID:   cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ResourceType: cloudflare.F(resource_tagging.AccountTagGetParamsResourceTypeWorker),
		WorkerID:     cloudflare.F("3f72a691-44b3-4c11-8642-c18a88ddaa5e"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
