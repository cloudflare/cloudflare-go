// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/registrar"
)

func TestRegistrationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Registrar.Registrations.New(context.TODO(), registrar.RegistrationNewParams{
		AccountID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		DomainName: cloudflare.F("my-new-startup.com"),
		AutoRenew:  cloudflare.F(false),
		Contacts: cloudflare.F(registrar.RegistrationNewParamsContacts{
			Registrant: cloudflare.F(registrar.RegistrationNewParamsContactsRegistrant{
				Email: cloudflare.F("ada@example.com"),
				Phone: cloudflare.F("+1.5555555555"),
				PostalInfo: cloudflare.F(registrar.RegistrationNewParamsContactsRegistrantPostalInfo{
					Address: cloudflare.F(registrar.RegistrationNewParamsContactsRegistrantPostalInfoAddress{
						City:        cloudflare.F("Austin"),
						CountryCode: cloudflare.F("US"),
						PostalCode:  cloudflare.F("78701"),
						State:       cloudflare.F("TX"),
						Street:      cloudflare.F("123 Main St"),
					}),
					Name:         cloudflare.F("Ada Lovelace"),
					Organization: cloudflare.F("Example Inc"),
				}),
				Fax: cloudflare.F("+1.5555555555"),
			}),
		}),
		PrivacyMode: cloudflare.F(registrar.RegistrationNewParamsPrivacyModeRedaction),
		Years:       cloudflare.F(int64(1)),
		Prefer:      cloudflare.F("Prefer"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRegistrationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Registrar.Registrations.List(context.TODO(), registrar.RegistrationListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Cursor:    cloudflare.F("cursor"),
		Direction: cloudflare.F(registrar.RegistrationListParamsDirectionAsc),
		PerPage:   cloudflare.F(int64(1)),
		SortBy:    cloudflare.F(registrar.RegistrationListParamsSortByRegistryCreatedAt),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRegistrationEditWithOptionalParams(t *testing.T) {
	t.Skip("422: Prism mock rejects test fixture domain with 'Domain not found'")
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
	_, err := client.Registrar.Registrations.Edit(
		context.TODO(),
		"example.com",
		registrar.RegistrationEditParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AutoRenew: cloudflare.F(false),
			Prefer:    cloudflare.F(registrar.RegistrationEditParamsPreferRespondAsync),
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

func TestRegistrationGet(t *testing.T) {
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
	_, err := client.Registrar.Registrations.Get(
		context.TODO(),
		"example.com",
		registrar.RegistrationGetParams{
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
