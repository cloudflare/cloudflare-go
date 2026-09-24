// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/registrar"
)

func TestTransferInNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Registrar.TransferIn.New(
		context.TODO(),
		"example.com",
		registrar.TransferInNewParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AuthCode:  cloudflare.F("bml4b3M+Pj5hcmNoLWxpbnV4"),
			AutoRenew: cloudflare.F(false),
			ContactExtensions: cloudflare.F(map[string]interface{}{
				"application_purpose": "bar",
				"nexus_category":      "bar",
			}),
			Contacts: cloudflare.F(registrar.TransferInNewParamsContacts{
				Administrator: cloudflare.F(registrar.TransferInNewParamsContactsAdministrator{
					Email: cloudflare.F("ada@example.com"),
					Phone: cloudflare.F("+1.5555555555"),
					PostalInfo: cloudflare.F(registrar.TransferInNewParamsContactsAdministratorPostalInfo{
						Address: cloudflare.F(registrar.TransferInNewParamsContactsAdministratorPostalInfoAddress{
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
				Billing: cloudflare.F(registrar.TransferInNewParamsContactsBilling{
					Email: cloudflare.F("ada@example.com"),
					Phone: cloudflare.F("+1.5555555555"),
					PostalInfo: cloudflare.F(registrar.TransferInNewParamsContactsBillingPostalInfo{
						Address: cloudflare.F(registrar.TransferInNewParamsContactsBillingPostalInfoAddress{
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
				Registrant: cloudflare.F(registrar.TransferInNewParamsContactsRegistrant{
					Email: cloudflare.F("ada@example.com"),
					Phone: cloudflare.F("+1.5555555555"),
					PostalInfo: cloudflare.F(registrar.TransferInNewParamsContactsRegistrantPostalInfo{
						Address: cloudflare.F(registrar.TransferInNewParamsContactsRegistrantPostalInfoAddress{
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
				Technical: cloudflare.F(registrar.TransferInNewParamsContactsTechnical{
					Email: cloudflare.F("ada@example.com"),
					Phone: cloudflare.F("+1.5555555555"),
					PostalInfo: cloudflare.F(registrar.TransferInNewParamsContactsTechnicalPostalInfo{
						Address: cloudflare.F(registrar.TransferInNewParamsContactsTechnicalPostalInfoAddress{
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
			PrivacyMode: cloudflare.F(registrar.TransferInNewParamsPrivacyModeRedaction),
			Prefer:      cloudflare.F("Prefer"),
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
