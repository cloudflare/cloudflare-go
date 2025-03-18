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

func TestOnRampNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicCloudNetworking.OnRamps.New(context.TODO(), magic_cloud_networking.OnRampNewParams{
		AccountID:                 cloudflare.F("account_id"),
		CloudType:                 cloudflare.F(magic_cloud_networking.OnRampNewParamsCloudTypeAws),
		InstallRoutesInCloud:      cloudflare.F(true),
		InstallRoutesInMagicWAN:   cloudflare.F(true),
		Name:                      cloudflare.F("name"),
		Type:                      cloudflare.F(magic_cloud_networking.OnRampNewParamsTypeOnrampTypeSingle),
		AdoptedHubID:              cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		AttachedHubs:              cloudflare.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		AttachedVPCs:              cloudflare.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		Description:               cloudflare.F("description"),
		HubProviderID:             cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ManageHubToHubAttachments: cloudflare.F(true),
		ManageVPCToHubAttachments: cloudflare.F(true),
		Region:                    cloudflare.F("region"),
		VPC:                       cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Forwarded:                 cloudflare.F("forwarded"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnRampUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicCloudNetworking.OnRamps.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		magic_cloud_networking.OnRampUpdateParams{
			AccountID:                 cloudflare.F("account_id"),
			AttachedHubs:              cloudflare.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			AttachedVPCs:              cloudflare.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			Description:               cloudflare.F("description"),
			InstallRoutesInCloud:      cloudflare.F(true),
			InstallRoutesInMagicWAN:   cloudflare.F(true),
			ManageHubToHubAttachments: cloudflare.F(true),
			ManageVPCToHubAttachments: cloudflare.F(true),
			Name:                      cloudflare.F("name"),
			VPC:                       cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestOnRampListWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicCloudNetworking.OnRamps.List(context.TODO(), magic_cloud_networking.OnRampListParams{
		AccountID: cloudflare.F("account_id"),
		Desc:      cloudflare.F(true),
		OrderBy:   cloudflare.F("order_by"),
		Status:    cloudflare.F(true),
		VPCs:      cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOnRampDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicCloudNetworking.OnRamps.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		magic_cloud_networking.OnRampDeleteParams{
			AccountID: cloudflare.F("account_id"),
			Destroy:   cloudflare.F(true),
			Force:     cloudflare.F(true),
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

func TestOnRampApply(t *testing.T) {
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
	_, err := client.MagicCloudNetworking.OnRamps.Apply(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		magic_cloud_networking.OnRampApplyParams{
			AccountID: cloudflare.F("account_id"),
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

func TestOnRampEditWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicCloudNetworking.OnRamps.Edit(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		magic_cloud_networking.OnRampEditParams{
			AccountID:                 cloudflare.F("account_id"),
			AttachedHubs:              cloudflare.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			AttachedVPCs:              cloudflare.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			Description:               cloudflare.F("description"),
			InstallRoutesInCloud:      cloudflare.F(true),
			InstallRoutesInMagicWAN:   cloudflare.F(true),
			ManageHubToHubAttachments: cloudflare.F(true),
			ManageVPCToHubAttachments: cloudflare.F(true),
			Name:                      cloudflare.F("name"),
			VPC:                       cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestOnRampExport(t *testing.T) {
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
	resp, err := client.MagicCloudNetworking.OnRamps.Export(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		magic_cloud_networking.OnRampExportParams{
			AccountID: cloudflare.F("account_id"),
		},
	)
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

func TestOnRampGetWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicCloudNetworking.OnRamps.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		magic_cloud_networking.OnRampGetParams{
			AccountID:          cloudflare.F("account_id"),
			PlannedResources:   cloudflare.F(true),
			PostApplyResources: cloudflare.F(true),
			Status:             cloudflare.F(true),
			VPCs:               cloudflare.F(true),
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

func TestOnRampPlan(t *testing.T) {
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
	_, err := client.MagicCloudNetworking.OnRamps.Plan(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		magic_cloud_networking.OnRampPlanParams{
			AccountID: cloudflare.F("account_id"),
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
