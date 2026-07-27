// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_sandbox_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/registrar_sandbox"
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
	_, err := client.RegistrarSandbox.Registrations.New(context.TODO(), registrar_sandbox.RegistrationNewParams{
		AccountID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		DomainName: cloudflare.F("my-brand-example.io"),
		Acknowledgements: cloudflare.F(map[string]interface{}{
			"fees": "bar",
		}),
		AutoRenew: cloudflare.F(false),
		ContactExtensions: cloudflare.F(map[string]interface{}{
			"application_purpose": "bar",
			"nexus_category":      "bar",
		}),
		Contacts: cloudflare.F(registrar_sandbox.RegistrationNewParamsContacts{
			Administrator: cloudflare.F(registrar_sandbox.RegistrationNewParamsContactsAdministrator{
				Email: cloudflare.F("katherine@example.io"),
				Phone: cloudflare.F("+1.5555550102"),
				PostalInfo: cloudflare.F(registrar_sandbox.RegistrationNewParamsContactsAdministratorPostalInfo{
					Address: cloudflare.F(registrar_sandbox.RegistrationNewParamsContactsAdministratorPostalInfoAddress{
						City:        cloudflare.F("San Francisco"),
						CountryCode: cloudflare.F("US"),
						PostalCode:  cloudflare.F("94103"),
						State:       cloudflare.F("CA"),
						Street:      cloudflare.F("789 Mission St"),
					}),
					Name:         cloudflare.F("Katherine Johnson"),
					Organization: cloudflare.F("Example Admin Inc"),
				}),
				Fax: cloudflare.F("+1.5555555555"),
			}),
			Billing: cloudflare.F(registrar_sandbox.RegistrationNewParamsContactsBilling{
				Email: cloudflare.F("dorothy@example.io"),
				Phone: cloudflare.F("+1.5555550103"),
				PostalInfo: cloudflare.F(registrar_sandbox.RegistrationNewParamsContactsBillingPostalInfo{
					Address: cloudflare.F(registrar_sandbox.RegistrationNewParamsContactsBillingPostalInfoAddress{
						City:        cloudflare.F("San Francisco"),
						CountryCode: cloudflare.F("US"),
						PostalCode:  cloudflare.F("94105"),
						State:       cloudflare.F("CA"),
						Street:      cloudflare.F("101 Howard St"),
					}),
					Name:         cloudflare.F("Dorothy Vaughan"),
					Organization: cloudflare.F("Example Billing Inc"),
				}),
				Fax: cloudflare.F("+1.5555555555"),
			}),
			Registrant: cloudflare.F(registrar_sandbox.RegistrationNewParamsContactsRegistrant{
				Email: cloudflare.F("ada@example.io"),
				Phone: cloudflare.F("+1.5555555555"),
				PostalInfo: cloudflare.F(registrar_sandbox.RegistrationNewParamsContactsRegistrantPostalInfo{
					Address: cloudflare.F(registrar_sandbox.RegistrationNewParamsContactsRegistrantPostalInfoAddress{
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
			Technical: cloudflare.F(registrar_sandbox.RegistrationNewParamsContactsTechnical{
				Email: cloudflare.F("grace@example.io"),
				Phone: cloudflare.F("+1.5555550101"),
				PostalInfo: cloudflare.F(registrar_sandbox.RegistrationNewParamsContactsTechnicalPostalInfo{
					Address: cloudflare.F(registrar_sandbox.RegistrationNewParamsContactsTechnicalPostalInfoAddress{
						City:        cloudflare.F("San Francisco"),
						CountryCode: cloudflare.F("US"),
						PostalCode:  cloudflare.F("94105"),
						State:       cloudflare.F("CA"),
						Street:      cloudflare.F("456 Market St"),
					}),
					Name:         cloudflare.F("Grace Hopper"),
					Organization: cloudflare.F("Example Technical Inc"),
				}),
				Fax: cloudflare.F("+1.5555555555"),
			}),
		}),
		PrivacyMode: cloudflare.F(registrar_sandbox.RegistrationNewParamsPrivacyModeRedaction),
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
	_, err := client.RegistrarSandbox.Registrations.List(context.TODO(), registrar_sandbox.RegistrationListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Cursor:    cloudflare.F("cursor"),
		Direction: cloudflare.F(registrar_sandbox.RegistrationListParamsDirectionAsc),
		PerPage:   cloudflare.F(int64(1)),
		SortBy:    cloudflare.F(registrar_sandbox.RegistrationListParamsSortByRegistryCreatedAt),
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
	_, err := client.RegistrarSandbox.Registrations.Edit(
		context.TODO(),
		"example.com",
		registrar_sandbox.RegistrationEditParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AutoRenew: cloudflare.F(false),
			Prefer:    cloudflare.F(registrar_sandbox.RegistrationEditParamsPreferRespondAsync),
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
	_, err := client.RegistrarSandbox.Registrations.Get(
		context.TODO(),
		"example.com",
		registrar_sandbox.RegistrationGetParams{
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
