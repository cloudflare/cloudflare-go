package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	createdAndModifiedTimestamp, _ = time.Parse(time.RFC3339, "2018-08-28T17:26:26Z")
	expiresAtTimestamp, _          = time.Parse(time.RFC3339, "2019-08-28T23:59:59Z")
	expectedRegistrarTransferIn    = RegistrarTransferIn{
		UnlockDomain:      "ok",
		DisablePrivacy:    "ok",
		EnterAuthCode:     "needed",
		ApproveTransfer:   "unknown",
		AcceptFoa:         "needed",
		CanCancelTransfer: true,
	}
	expectedRegistrarContact = RegistrantContact{
		ID:           "ea95132c15732412d22c1476fa83f27a",
		FirstName:    "John",
		LastName:     "Appleseed",
		Organization: "Cloudflare, Inc.",
		Address:      "123 Sesame St.",
		Address2:     "Suite 430",
		City:         "Austin",
		State:        "TX",
		Zip:          "12345",
		Country:      "US",
		Phone:        "+1 123-123-1234",
		Email:        "user@example.com",
		Fax:          "123-867-5309",
	}
	expectedRegistrarDomain = RegistrarDomain{
		RegistrarDomainConfiguration: RegistrarDomainConfiguration{
			Locked:      false,
			NameServers: []string{"ns1.cloudflare.com", "ns2.cloudflare.com"},
			Privacy:     true,
			AutoRenew:   true,
		},
		ID:                "ea95132c15732412d22c1476fa83f27a",
		Name:              "cloudflare.com",
		Available:         false,
		SupportedTLD:      true,
		CanRegister:       false,
		TransferIn:        expectedRegistrarTransferIn,
		CurrentRegistrar:  "Cloudflare",
		ExpiresAt:         expiresAtTimestamp,
		RegistryStatuses:  "ok,serverTransferProhibited",
		CreatedAt:         createdAndModifiedTimestamp,
		UpdatedAt:         createdAndModifiedTimestamp,
		RegistrantContact: expectedRegistrarContact,
		Permissions:       []string{"show_private_data"},
	}
)

func TestRegistrarDomain(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "ea95132c15732412d22c1476fa83f27a",
				"administrator_contact_id": 12345,
				"auto_renew": true,
				"available": false,
				"supported_tld": true,
				"billing_contact_id": 12345,
				"can_register": false,
				"transfer_in": {
					"unlock_domain": "ok",
					"disable_privacy": "ok",
					"enter_auth_code": "needed",
					"approve_transfer": "unknown",
					"accept_foa": "needed",
					"can_cancel_transfer": true
				},
				"cloudflare_dns": false,
				"cloudflare_registration": true,
				"contacts": {
					"administrator_id": 12345,
					"billing_id": 12345,
					"registrant_id": 12345,
					"technical_id": 12345
				},
				"contacts_updated_at": "2018-08-28T17:26:26Z",
				"created_registrar": "Cloudflare",
				"current_registrar": "Cloudflare",
				"expires_at": "2019-08-28T23:59:59Z",
				"registry_statuses": "ok,serverTransferProhibited",
				"locked": false,
				"created_at": "2018-08-28T17:26:26Z",
				"updated_at": "2018-08-28T17:26:26Z",
				"registrant_contact": {
					"id": "ea95132c15732412d22c1476fa83f27a",
					"first_name": "John",
					"last_name": "Appleseed",
					"organization": "Cloudflare, Inc.",
					"address": "123 Sesame St.",
					"address2": "Suite 430",
					"city": "Austin",
					"state": "TX",
					"zip": "12345",
					"country": "US",
					"phone": "+1 123-123-1234",
					"email": "user@example.com",
					"fax": "123-867-5309"
				},
				"dns": [],
				"ds_records": [],
				"email_verified": true,
				"fees": {
					"icann_fee": 0,
					"redemption_fee": 0,
					"registration_fee": 0,
					"renewal_fee": 0,
					"transfer_fee": 0
				},
				"last_known_status": "registrationActive",
				"last_transferred_at": "2020-07-26T04:18:28Z",
				"name": "cloudflare.com",
				"name_servers": [
					"ns1.cloudflare.com",
					"ns2.cloudflare.com"
				],
				"payment_expires_at": "2022-09-10T05:16:19Z",
				"pending_transfer": false,
				"permissions": [
					"show_private_data"
				],
				"previous_registrar": "",
				"privacy": true,
				"registered_at": null,
				"registrant_contact_id": 12345,
				"registry_object_id": "SOME_ID_FROM_REGISTRY",
				"technical_contact_id": 12345,
				"transfer_conditions": {
					"exists": true,
					"not_premium": true,
					"not_secure": true,
					"not_started": true,
					"not_waiting": true,
					"supported_tld": true
				},
				"updated_registrar": "Cloudflare",
				"using_created_registrar_nameservers": false,
				"using_current_registrar_nameservers": false,
				"using_previous_registrar_nameservers": false,
				"using_updated_registrar_nameservers": false,
				"whois": {
					"administrator": {
						"address": "DATA REDACTED",
						"city": "DATA REDACTED",
						"country": "US",
						"email": "https://domaincontact.cloudflareregistrar.com/cloudflare.com",
						"fax": "DATA REDACTED",
						"first_name": "DATA REDACTED",
						"last_name": "DATA REDACTED",
						"organization": "DATA REDACTED",
						"phone": "DATA REDACTED",
						"state": "TX",
						"zip": "DATA REDACTED"
					},
					"billing": {
						"address": "DATA REDACTED",
						"city": "DATA REDACTED",
						"country": "US",
						"email": "https://domaincontact.cloudflareregistrar.com/cloudflare.com",
						"fax": "DATA REDACTED",
						"first_name": "DATA REDACTED",
						"last_name": "DATA REDACTED",
						"organization": "DATA REDACTED",
						"phone": "DATA REDACTED",
						"state": "TX",
						"zip": "DATA REDACTED"
					},
					"privacy": true,
					"raw": "A really long string with all the WHOIS data in a pretty format",
					"registrant": {
						"address": "DATA REDACTED",
						"city": "DATA REDACTED",
						"country": "US",
						"email": "https://domaincontact.cloudflareregistrar.com/cloudflare.com",
						"fax": "DATA REDACTED",
						"first_name": "DATA REDACTED",
						"last_name": "DATA REDACTED",
						"organization": "DATA REDACTED",
						"phone": "DATA REDACTED",
						"state": "TX",
						"zip": "DATA REDACTED"
					},
					"technical": {
						"address": "DATA REDACTED",
						"city": "DATA REDACTED",
						"country": "US",
						"email": "https://domaincontact.cloudflareregistrar.com/cloudflare.com",
						"fax": "DATA REDACTED",
						"first_name": "DATA REDACTED",
						"last_name": "DATA REDACTED",
						"organization": "DATA REDACTED",
						"phone": "DATA REDACTED",
						"state": "TX",
						"zip": "DATA REDACTED"
					}
				}
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/registrar/domains/cloudflare.com", handler)

	actual, err := client.RegistrarDomain(context.Background(), "01a7362d577a6c3019a474fd6f485823", "cloudflare.com")

	if assert.NoError(t, err) {
		assert.Equal(t, expectedRegistrarDomain, actual)
	}
}

func TestRegistrarDomains(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "ea95132c15732412d22c1476fa83f27a",
					"administrator_contact_id": 12345,
					"auto_renew": true,
					"available": false,
					"supported_tld": true,
					"billing_contact_id": 12345,
					"can_register": false,
					"transfer_in": {
						"unlock_domain": "ok",
						"disable_privacy": "ok",
						"enter_auth_code": "needed",
						"approve_transfer": "unknown",
						"accept_foa": "needed",
						"can_cancel_transfer": true
					},
					"cloudflare_dns": false,
					"cloudflare_registration": true,
					"contacts": {
						"administrator_id": 12345,
						"billing_id": 12345,
						"registrant_id": 12345,
						"technical_id": 12345
					},
					"contacts_updated_at": "2018-08-28T17:26:26Z",
					"created_registrar": "Cloudflare",
					"current_registrar": "Cloudflare",
					"expires_at": "2019-08-28T23:59:59Z",
					"registry_statuses": "ok,serverTransferProhibited",
					"locked": false,
					"created_at": "2018-08-28T17:26:26Z",
					"updated_at": "2018-08-28T17:26:26Z",
					"registrant_contact": {
						"id": "ea95132c15732412d22c1476fa83f27a",
						"first_name": "John",
						"last_name": "Appleseed",
						"organization": "Cloudflare, Inc.",
						"address": "123 Sesame St.",
						"address2": "Suite 430",
						"city": "Austin",
						"state": "TX",
						"zip": "12345",
						"country": "US",
						"phone": "+1 123-123-1234",
						"email": "user@example.com",
						"fax": "123-867-5309"
					},
					"dns": [],
					"ds_records": [],
					"email_verified": true,
					"fees": {
						"icann_fee": 0,
						"redemption_fee": 0,
						"registration_fee": 0,
						"renewal_fee": 0,
						"transfer_fee": 0
					},
					"last_known_status": "registrationActive",
					"last_transferred_at": "2020-07-26T04:18:28Z",
					"name": "cloudflare.com",
					"name_servers": [
						"ns1.cloudflare.com",
						"ns2.cloudflare.com"
					],
					"payment_expires_at": "2022-09-10T05:16:19Z",
					"pending_transfer": false,
					"permissions": [
						"show_private_data"
					],
					"previous_registrar": "",
					"privacy": true,
					"registered_at": null,
					"registrant_contact_id": 12345,
					"registry_object_id": "SOME_ID_FROM_REGISTRY",
					"technical_contact_id": 12345,
					"transfer_conditions": {
						"exists": true,
						"not_premium": true,
						"not_secure": true,
						"not_started": true,
						"not_waiting": true,
						"supported_tld": true
					},
					"updated_registrar": "Cloudflare",
					"using_created_registrar_nameservers": false,
					"using_current_registrar_nameservers": false,
					"using_previous_registrar_nameservers": false,
					"using_updated_registrar_nameservers": false,
					"whois": {
						"administrator": {
							"address": "DATA REDACTED",
							"city": "DATA REDACTED",
							"country": "US",
							"email": "https://domaincontact.cloudflareregistrar.com/cloudflare.com",
							"fax": "DATA REDACTED",
							"first_name": "DATA REDACTED",
							"last_name": "DATA REDACTED",
							"organization": "DATA REDACTED",
							"phone": "DATA REDACTED",
							"state": "TX",
							"zip": "DATA REDACTED"
						},
						"billing": {
							"address": "DATA REDACTED",
							"city": "DATA REDACTED",
							"country": "US",
							"email": "https://domaincontact.cloudflareregistrar.com/cloudflare.com",
							"fax": "DATA REDACTED",
							"first_name": "DATA REDACTED",
							"last_name": "DATA REDACTED",
							"organization": "DATA REDACTED",
							"phone": "DATA REDACTED",
							"state": "TX",
							"zip": "DATA REDACTED"
						},
						"privacy": true,
						"raw": "A really long string with all the WHOIS data in a pretty format",
						"registrant": {
							"address": "DATA REDACTED",
							"city": "DATA REDACTED",
							"country": "US",
							"email": "https://domaincontact.cloudflareregistrar.com/cloudflare.com",
							"fax": "DATA REDACTED",
							"first_name": "DATA REDACTED",
							"last_name": "DATA REDACTED",
							"organization": "DATA REDACTED",
							"phone": "DATA REDACTED",
							"state": "TX",
							"zip": "DATA REDACTED"
						},
						"technical": {
							"address": "DATA REDACTED",
							"city": "DATA REDACTED",
							"country": "US",
							"email": "https://domaincontact.cloudflareregistrar.com/cloudflare.com",
							"fax": "DATA REDACTED",
							"first_name": "DATA REDACTED",
							"last_name": "DATA REDACTED",
							"organization": "DATA REDACTED",
							"phone": "DATA REDACTED",
							"state": "TX",
							"zip": "DATA REDACTED"
						}
					}
				}
			],
			"result_info": {
				"page": 1,
				"per_page": 20,
				"count": 1,
				"total_count": 2000
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/registrar/domains", handler)

	actual, err := client.RegistrarDomains(context.Background(), "01a7362d577a6c3019a474fd6f485823")

	if assert.NoError(t, err) {
		assert.Equal(t, []RegistrarDomain{expectedRegistrarDomain}, actual)
	}
}

func TestTransferRegistrarDomain(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "ea95132c15732412d22c1476fa83f27a",
					"available": false,
					"supported_tld": true,
					"can_register": false,
					"transfer_in": {
						"unlock_domain": "ok",
						"disable_privacy": "ok",
						"enter_auth_code": "needed",
						"approve_transfer": "unknown",
						"accept_foa": "needed",
						"can_cancel_transfer": true
					},
					"current_registrar": "Cloudflare",
					"expires_at": "2019-08-28T23:59:59Z",
					"registry_statuses": "ok,serverTransferProhibited",
					"locked": false,
					"created_at": "2018-08-28T17:26:26Z",
					"updated_at": "2018-08-28T17:26:26Z",
					"registrant_contact": {
						"id": "ea95132c15732412d22c1476fa83f27a",
						"first_name": "John",
						"last_name": "Appleseed",
						"organization": "Cloudflare, Inc.",
						"address": "123 Sesame St.",
						"address2": "Suite 430",
						"city": "Austin",
						"state": "TX",
						"zip": "12345",
						"country": "US",
						"phone": "+1 123-123-1234",
						"email": "user@example.com",
						"fax": "123-867-5309"
					},
					"name": "cloudflare.com",
					"name_servers": ["ns1.cloudflare.com", "ns2.cloudflare.com"],
					"privacy": true,
					"auto_renew": true,
					"permissions": [
						"show_private_data"
					]
				}
			],
			"result_info": {
				"page": 1,
				"per_page": 20,
				"count": 1,
				"total_count": 2000
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/registrar/domains/cloudflare.com/transfer", handler)

	actual, err := client.TransferRegistrarDomain(context.Background(), "01a7362d577a6c3019a474fd6f485823", "cloudflare.com")

	if assert.NoError(t, err) {
		assert.Equal(t, []RegistrarDomain{expectedRegistrarDomain}, actual)
	}
}

func TestCancelRegistrarDomainTransfer(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "ea95132c15732412d22c1476fa83f27a",
					"available": false,
					"supported_tld": true,
					"can_register": false,
					"transfer_in": {
						"unlock_domain": "ok",
						"disable_privacy": "ok",
						"enter_auth_code": "needed",
						"approve_transfer": "unknown",
						"accept_foa": "needed",
						"can_cancel_transfer": true
					},
					"current_registrar": "Cloudflare",
					"expires_at": "2019-08-28T23:59:59Z",
					"registry_statuses": "ok,serverTransferProhibited",
					"locked": false,
					"created_at": "2018-08-28T17:26:26Z",
					"updated_at": "2018-08-28T17:26:26Z",
					"registrant_contact": {
						"id": "ea95132c15732412d22c1476fa83f27a",
						"first_name": "John",
						"last_name": "Appleseed",
						"organization": "Cloudflare, Inc.",
						"address": "123 Sesame St.",
						"address2": "Suite 430",
						"city": "Austin",
						"state": "TX",
						"zip": "12345",
						"country": "US",
						"phone": "+1 123-123-1234",
						"email": "user@example.com",
						"fax": "123-867-5309"
					},
					"name": "cloudflare.com",
					"name_servers": ["ns1.cloudflare.com", "ns2.cloudflare.com"],
					"privacy": true,
					"auto_renew": true,
					"permissions": [
						"show_private_data"
					]
				}
			],
			"result_info": {
				"page": 1,
				"per_page": 20,
				"count": 1,
				"total_count": 2000
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/registrar/domains/cloudflare.com/cancel_transfer", handler)

	actual, err := client.CancelRegistrarDomainTransfer(context.Background(), "01a7362d577a6c3019a474fd6f485823", "cloudflare.com")

	if assert.NoError(t, err) {
		assert.Equal(t, []RegistrarDomain{expectedRegistrarDomain}, actual)
	}
}

func TestUpdateRegistrarDomain(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "ea95132c15732412d22c1476fa83f27a",
				"available": false,
				"supported_tld": true,
				"can_register": false,
				"transfer_in": {
					"unlock_domain": "ok",
					"disable_privacy": "ok",
					"enter_auth_code": "needed",
					"approve_transfer": "unknown",
					"accept_foa": "needed",
					"can_cancel_transfer": true
				},
				"current_registrar": "Cloudflare",
				"expires_at": "2019-08-28T23:59:59Z",
				"registry_statuses": "ok,serverTransferProhibited",
				"locked": false,
				"created_at": "2018-08-28T17:26:26Z",
				"updated_at": "2018-08-28T17:26:26Z",
				"registrant_contact": {
					"id": "ea95132c15732412d22c1476fa83f27a",
					"first_name": "John",
					"last_name": "Appleseed",
					"organization": "Cloudflare, Inc.",
					"address": "123 Sesame St.",
					"address2": "Suite 430",
					"city": "Austin",
					"state": "TX",
					"zip": "12345",
					"country": "US",
					"phone": "+1 123-123-1234",
					"email": "user@example.com",
					"fax": "123-867-5309"
				},
				"name": "cloudflare.com",
				"name_servers": ["ns1.cloudflare.com", "ns2.cloudflare.com"],
				"privacy": true,
				"auto_renew": true,
				"permissions": [
					"show_private_data"
				]
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/registrar/domains/cloudflare.com", handler)

	actual, err := client.UpdateRegistrarDomain(context.Background(), "01a7362d577a6c3019a474fd6f485823", "cloudflare.com", RegistrarDomainConfiguration{
		NameServers: []string{"ns1.cloudflare.com", "ns2.cloudflare.com"},
		Locked:      true,
	})

	if assert.NoError(t, err) {
		assert.Equal(t, expectedRegistrarDomain, actual)
	}
}

func TestGetRegistrarContact(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "ea95132c15732412d22c1476fa83f27a",
				"first_name": "John",
				"last_name": "Appleseed",
				"organization": "Cloudflare, Inc.",
				"address": "123 Sesame St.",
				"address2": "Suite 430",
				"city": "Austin",
				"state": "TX",
				"zip": "12345",
				"country": "US",
				"phone": "+1 123-123-1234",
				"email": "user@example.com",
				"fax": "123-867-5309"
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/registrar/contacts/01a7362d577a6c3019a474fd6f485823", handler)

	actual, err := client.RegistrarContact(context.Background(), "01a7362d577a6c3019a474fd6f485823", "01a7362d577a6c3019a474fd6f485823")

	if assert.NoError(t, err) {
		assert.Equal(t, expectedRegistrarContact, actual)
	}
}

func TestGetRegistrarContacts(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [{
				"id": "ea95132c15732412d22c1476fa83f27a",
				"first_name": "John",
				"last_name": "Appleseed",
				"organization": "Cloudflare, Inc.",
				"address": "123 Sesame St.",
				"address2": "Suite 430",
				"city": "Austin",
				"state": "TX",
				"zip": "12345",
				"country": "US",
				"phone": "+1 123-123-1234",
				"email": "user@example.com",
				"fax": "123-867-5309"
			}]
		}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/registrar/contacts", handler)

	actual, err := client.RegistrarContacts(context.Background(), "01a7362d577a6c3019a474fd6f485823")

	if assert.NoError(t, err) {
		assert.Equal(t, []RegistrantContact{expectedRegistrarContact}, actual)
	}
}
