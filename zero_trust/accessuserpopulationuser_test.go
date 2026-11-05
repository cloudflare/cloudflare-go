// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
)

func TestAccessUserPopulationUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.UserPopulations.Users.New(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessUserPopulationUserNewParams{
			AccountID:   cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Active:      cloudflare.F(true),
			DisplayName: cloudflare.F("displayName"),
			Emails: cloudflare.F([]zero_trust.AccessUserPopulationUserNewParamsEmail{{
				Primary: cloudflare.F(true),
				Type:    cloudflare.F("type"),
				Value:   cloudflare.F("value"),
			}}),
			ExternalID: cloudflare.F("externalId"),
			Meta: cloudflare.F(zero_trust.AccessUserPopulationUserNewParamsMeta{
				Created:      cloudflare.F(time.Now()),
				LastModified: cloudflare.F(time.Now()),
				Location:     cloudflare.F("https://example.com"),
				ResourceType: cloudflare.F("resourceType"),
				Version:      cloudflare.F("version"),
			}),
			Name: cloudflare.F(zero_trust.AccessUserPopulationUserNewParamsName{
				FamilyName:      cloudflare.F("familyName"),
				Formatted:       cloudflare.F("formatted"),
				GivenName:       cloudflare.F("givenName"),
				HonorificPrefix: cloudflare.F("honorificPrefix"),
				HonorificSuffix: cloudflare.F("honorificSuffix"),
				MiddleName:      cloudflare.F("middleName"),
			}),
			Schemas:  cloudflare.F([]string{"string"}),
			UserName: cloudflare.F("userName"),
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

func TestAccessUserPopulationUserUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.UserPopulations.Users.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"user_id",
		zero_trust.AccessUserPopulationUserUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Operations: cloudflare.F([]zero_trust.AccessUserPopulationUserUpdateParamsOperation{{
				Op:    cloudflare.F(zero_trust.AccessUserPopulationUserUpdateParamsOperationsOpAdd),
				Path:  cloudflare.F("path"),
				Value: cloudflare.F[zero_trust.AccessUserPopulationUserUpdateParamsOperationsValueUnion](shared.UnionString("string")),
			}}),
			Schemas: cloudflare.F([]string{"string"}),
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

func TestAccessUserPopulationUserListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.UserPopulations.Users.List(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessUserPopulationUserListParams{
			AccountID:          cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Attributes:         cloudflare.F("attributes"),
			Count:              cloudflare.F(int64(0)),
			ExcludedAttributes: cloudflare.F("excludedAttributes"),
			Filter:             cloudflare.F("filter"),
			StartIndex:         cloudflare.F(int64(1)),
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

func TestAccessUserPopulationUserDelete(t *testing.T) {
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
	err := client.ZeroTrust.Access.UserPopulations.Users.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"user_id",
		zero_trust.AccessUserPopulationUserDeleteParams{
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

func TestAccessUserPopulationUserGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.UserPopulations.Users.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"user_id",
		zero_trust.AccessUserPopulationUserGetParams{
			AccountID:          cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Attributes:         cloudflare.F("attributes"),
			ExcludedAttributes: cloudflare.F("excludedAttributes"),
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

func TestAccessUserPopulationUserReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.UserPopulations.Users.Replace(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"user_id",
		zero_trust.AccessUserPopulationUserReplaceParams{
			AccountID:   cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Active:      cloudflare.F(true),
			DisplayName: cloudflare.F("displayName"),
			Emails: cloudflare.F([]zero_trust.AccessUserPopulationUserReplaceParamsEmail{{
				Primary: cloudflare.F(true),
				Type:    cloudflare.F("type"),
				Value:   cloudflare.F("value"),
			}}),
			ExternalID: cloudflare.F("externalId"),
			Meta: cloudflare.F(zero_trust.AccessUserPopulationUserReplaceParamsMeta{
				Created:      cloudflare.F(time.Now()),
				LastModified: cloudflare.F(time.Now()),
				Location:     cloudflare.F("https://example.com"),
				ResourceType: cloudflare.F("resourceType"),
				Version:      cloudflare.F("version"),
			}),
			Name: cloudflare.F(zero_trust.AccessUserPopulationUserReplaceParamsName{
				FamilyName:      cloudflare.F("familyName"),
				Formatted:       cloudflare.F("formatted"),
				GivenName:       cloudflare.F("givenName"),
				HonorificPrefix: cloudflare.F("honorificPrefix"),
				HonorificSuffix: cloudflare.F("honorificSuffix"),
				MiddleName:      cloudflare.F("middleName"),
			}),
			Schemas:  cloudflare.F([]string{"string"}),
			UserName: cloudflare.F("userName"),
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
