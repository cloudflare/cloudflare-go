// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/email_security"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

func TestSettingImpersonationRegistryNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.EmailSecurity.Settings.ImpersonationRegistry.New(context.TODO(), email_security.SettingImpersonationRegistryNewParams{
		AccountID:               cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Email:                   cloudflare.F("john.doe@example.com"),
		IsEmailRegex:            cloudflare.F(false),
		Name:                    cloudflare.F("John Doe"),
		Comments:                cloudflare.F("comments"),
		DirectoryID:             cloudflare.F(int64(0)),
		DirectoryNodeID:         cloudflare.F(int64(0)),
		ExternalDirectoryNodeID: cloudflare.F("external_directory_node_id"),
		Provenance:              cloudflare.F(email_security.SettingImpersonationRegistryNewParamsProvenanceA1SInternal),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingImpersonationRegistryListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.EmailSecurity.Settings.ImpersonationRegistry.List(context.TODO(), email_security.SettingImpersonationRegistryListParams{
		AccountID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction:  cloudflare.F(email_security.SettingImpersonationRegistryListParamsDirectionAsc),
		Order:      cloudflare.F(email_security.SettingImpersonationRegistryListParamsOrderName),
		Page:       cloudflare.F(int64(1)),
		PerPage:    cloudflare.F(int64(20)),
		Provenance: cloudflare.F(email_security.SettingImpersonationRegistryListParamsProvenanceA1SInternal),
		Search:     cloudflare.F("search"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
