// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/organizations"
)

func TestOrganizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.New(context.TODO(), organizations.OrganizationNewParams{
		Organization: organizations.OrganizationParam{
			Name: cloudflare.F("name"),
			Parent: cloudflare.F(organizations.OrganizationParentParam{
				ID: cloudflare.F("a7b9c3d2e8f4g1h5i6j0k9l2m3n7o4p8"),
			}),
			Profile: cloudflare.F(accounts.AccountProfileParam{
				BusinessAddress:  cloudflare.F("business_address"),
				BusinessEmail:    cloudflare.F("business_email"),
				BusinessName:     cloudflare.F("business_name"),
				BusinessPhone:    cloudflare.F("business_phone"),
				ExternalMetadata: cloudflare.F("external_metadata"),
			}),
		},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.Update(
		context.TODO(),
		"a7b9c3d2e8f4g1h5i6j0k9l2m3n7o4p8",
		organizations.OrganizationUpdateParams{
			Organization: organizations.OrganizationParam{
				Name: cloudflare.F("name"),
				Parent: cloudflare.F(organizations.OrganizationParentParam{
					ID: cloudflare.F("a7b9c3d2e8f4g1h5i6j0k9l2m3n7o4p8"),
				}),
				Profile: cloudflare.F(accounts.AccountProfileParam{
					BusinessAddress:  cloudflare.F("business_address"),
					BusinessEmail:    cloudflare.F("business_email"),
					BusinessName:     cloudflare.F("business_name"),
					BusinessPhone:    cloudflare.F("business_phone"),
					ExternalMetadata: cloudflare.F("external_metadata"),
				}),
			},
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

func TestOrganizationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.List(context.TODO(), organizations.OrganizationListParams{
		ID: cloudflare.F([]string{"a7b9c3d2e8f4g1h5i6j0k9l2m3n7o4p8"}),
		Containing: cloudflare.F(organizations.OrganizationListParamsContaining{
			Account:      cloudflare.F("account"),
			Organization: cloudflare.F("organization"),
			User:         cloudflare.F("user"),
		}),
		Name: cloudflare.F(organizations.OrganizationListParamsName{
			Contains:   cloudflare.F("contains"),
			EndsWith:   cloudflare.F("endsWith"),
			StartsWith: cloudflare.F("startsWith"),
		}),
		PageSize:  cloudflare.F(int64(0)),
		PageToken: cloudflare.F("page_token"),
		Parent: cloudflare.F(organizations.OrganizationListParamsParent{
			ID: cloudflare.F(organizations.OrganizationListParamsParentIDNull),
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationDelete(t *testing.T) {
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
	err := client.Organizations.Delete(context.TODO(), "a7b9c3d2e8f4g1h5i6j0k9l2m3n7o4p8")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationGet(t *testing.T) {
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
	_, err := client.Organizations.Get(context.TODO(), "a7b9c3d2e8f4g1h5i6j0k9l2m3n7o4p8")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
