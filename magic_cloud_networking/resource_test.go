// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_cloud_networking_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/magic_cloud_networking"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestResourceListWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicCloudNetworking.Resources.List(context.TODO(), magic_cloud_networking.ResourceListParams{
		AccountID:     cloudflare.F("account_id"),
		Cloudflare:    cloudflare.F(true),
		Desc:          cloudflare.F(true),
		Managed:       cloudflare.F(true),
		OrderBy:       cloudflare.F("order_by"),
		Page:          cloudflare.F(int64(1)),
		PerPage:       cloudflare.F(int64(1)),
		ProviderID:    cloudflare.F("provider_id"),
		Region:        cloudflare.F("region"),
		ResourceGroup: cloudflare.F("resource_group"),
		ResourceID:    cloudflare.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		ResourceType:  cloudflare.F([]magic_cloud_networking.ResourceListParamsResourceType{magic_cloud_networking.ResourceListParamsResourceTypeAwsCustomerGateway}),
		Search:        cloudflare.F([]string{"string"}),
		V2:            cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResourceExportWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	resp, err := client.MagicCloudNetworking.Resources.Export(context.TODO(), magic_cloud_networking.ResourceExportParams{
		AccountID:     cloudflare.F("account_id"),
		Desc:          cloudflare.F(true),
		OrderBy:       cloudflare.F("order_by"),
		ProviderID:    cloudflare.F("provider_id"),
		Region:        cloudflare.F("region"),
		ResourceGroup: cloudflare.F("resource_group"),
		ResourceID:    cloudflare.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		ResourceType:  cloudflare.F([]magic_cloud_networking.ResourceExportParamsResourceType{magic_cloud_networking.ResourceExportParamsResourceTypeAwsCustomerGateway}),
		Search:        cloudflare.F([]string{"string"}),
		V2:            cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestResourceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicCloudNetworking.Resources.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		magic_cloud_networking.ResourceGetParams{
			AccountID: cloudflare.F("account_id"),
			V2:        cloudflare.F(true),
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

func TestResourcePolicyPreview(t *testing.T) {
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
	_, err := client.MagicCloudNetworking.Resources.PolicyPreview(context.TODO(), magic_cloud_networking.ResourcePolicyPreviewParams{
		AccountID: cloudflare.F("account_id"),
		Policy:    cloudflare.F("policy"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
