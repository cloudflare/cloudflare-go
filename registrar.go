package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type RegistrarDomainContacts struct {
	Administrator RegistrantContact `json:"administrator"`
	Billing       RegistrantContact `json:"billing"`
	Registrant    RegistrantContact `json:"registrant"`
	Technical     RegistrantContact `json:"technical"`
}

type RegistrarFees struct {
	Icann        float64 `json:"icann_fee"`
	Redemption   float64 `json:"redemption_fee"`
	Registration float64 `json:"registration_fee"`
	Renewal      float64 `json:"renewal_fee"`
	Transfer     float64 `json:"transfer_fee"`
}

// RegistrarDomain is the structure of the API response for a new
// Cloudflare Registrar domain.
type RegistrarDomain struct {
	RegistrarDomainConfiguration
	ID                string              `json:"id"`
	Available         bool                `json:"available"`
	SupportedTLD      bool                `json:"supported_tld"`
	CanRegister       bool                `json:"can_register"`
	TransferIn        RegistrarTransferIn `json:"transfer_in"`
	CurrentRegistrar  string              `json:"current_registrar"`
	ExpiresAt         time.Time           `json:"expires_at"`
	RegistryStatuses  string              `json:"registry_statuses"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
	RegistrantContact RegistrantContact   `json:"registrant_contact"`
	Fees              RegistrarFees       `json:"fees"`
	Name              string              `json:"name"`
	PendingTransfer   bool                `json:"pending_transfer"`
	Permissions       []string            `json:"permissions"`
}

// RegistrarTransferIn contains the structure for a domain transfer in
// request.
type RegistrarTransferIn struct {
	UnlockDomain      string `json:"unlock_domain"`
	DisablePrivacy    string `json:"disable_privacy"`
	EnterAuthCode     string `json:"enter_auth_code"`
	ApproveTransfer   string `json:"approve_transfer"`
	AcceptFoa         string `json:"accept_foa"`
	CanCancelTransfer bool   `json:"can_cancel_transfer"`
}

// RegistrantContact is the contact details for the domain registration.
type RegistrantContact struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Organization string `json:"organization"`
	Address      string `json:"address"`
	Address2     string `json:"address2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
	Country      string `json:"country"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Fax          string `json:"fax"`
	Verified     bool   `json:"email_verified"`
	Label        string `json:"label"`
}

// RegistrarDomainConfiguration is the structure for making updates to
// and existing domain.
type RegistrarDomainConfiguration struct {
	NameServers []string `json:"name_servers"`
	Privacy     bool     `json:"privacy"`
	Locked      bool     `json:"locked"`
	AutoRenew   bool     `json:"auto_renew"`
}

// RegistrarDomainDetailResponse is the structure of the detailed
// response from the API for a single domain.
type RegistrarDomainDetailResponse struct {
	Response
	Result RegistrarDomain `json:"result"`
}

// RegistrarDomainsDetailResponse is the structure of the detailed
// response from the API.
type RegistrarDomainsDetailResponse struct {
	Response
	Result []RegistrarDomain `json:"result"`
}

type RegistrarContactsDetailResponse struct {
	Response
	Result []RegistrantContact `json:"result"`
}

type RegistrarContactDetailResponse struct {
	Response
	Result RegistrantContact `json:"result"`
}

// RegistrarDomain returns a single domain based on the account ID and
// domain name.
//
// API reference: https://api.cloudflare.com/#registrar-domains-get-domain
func (api *API) RegistrarDomain(ctx context.Context, accountID, domainName string) (RegistrarDomain, error) {
	uri := fmt.Sprintf("/accounts/%s/registrar/domains/%s", accountID, domainName)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return RegistrarDomain{}, err
	}

	var r RegistrarDomainDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return RegistrarDomain{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// RegistrarDomains returns all registrar domains based on the account
// ID.
//
// API reference: https://api.cloudflare.com/#registrar-domains-list-domains
func (api *API) RegistrarDomains(ctx context.Context, accountID string) ([]RegistrarDomain, error) {
	uri := fmt.Sprintf("/accounts/%s/registrar/domains", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return []RegistrarDomain{}, err
	}

	var r RegistrarDomainsDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []RegistrarDomain{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// TransferRegistrarDomain initiates the transfer from another registrar
// to Cloudflare Registrar.
//
// API reference: https://api.cloudflare.com/#registrar-domains-transfer-domain
func (api *API) TransferRegistrarDomain(ctx context.Context, accountID, domainName string) ([]RegistrarDomain, error) {
	uri := fmt.Sprintf("/accounts/%s/registrar/domains/%s/transfer", accountID, domainName)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return []RegistrarDomain{}, err
	}

	var r RegistrarDomainsDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []RegistrarDomain{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// CancelRegistrarDomainTransfer cancels a pending domain transfer.
//
// API reference: https://api.cloudflare.com/#registrar-domains-cancel-transfer
func (api *API) CancelRegistrarDomainTransfer(ctx context.Context, accountID, domainName string) ([]RegistrarDomain, error) {
	uri := fmt.Sprintf("/accounts/%s/registrar/domains/%s/cancel_transfer", accountID, domainName)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return []RegistrarDomain{}, err
	}

	var r RegistrarDomainsDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []RegistrarDomain{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// UpdateRegistrarDomain updates an existing Registrar Domain configuration.
//
// API reference: https://api.cloudflare.com/#registrar-domains-update-domain
func (api *API) UpdateRegistrarDomain(ctx context.Context, accountID, domainName string, domainConfiguration RegistrarDomainConfiguration) (RegistrarDomain, error) {
	uri := fmt.Sprintf("/accounts/%s/registrar/domains/%s", accountID, domainName)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, domainConfiguration)
	if err != nil {
		return RegistrarDomain{}, err
	}

	var r RegistrarDomainDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return RegistrarDomain{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

func (api *API) getContacts(ctx context.Context, uri string) ([]RegistrantContact, error) {
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []RegistrantContact{}, err
	}

	var r RegistrarContactsDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []RegistrantContact{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}

// RegistrarContacts gets a list of all Contacts based on the account ID
//
// API reference: undocumented
func (api *API) RegistrarContacts(ctx context.Context, accountID string) ([]RegistrantContact, error) {
	uri := fmt.Sprintf("/accounts/%s/registrar/contacts", accountID)
	return api.getContacts(ctx, uri)
}

// RegistrarContact gets a RegistrantContact based on the account ID and contact ID
//
// API reference: undocumented
func (api *API) RegistrarContact(ctx context.Context, accountID, contactID string) (RegistrantContact, error) {
	uri := fmt.Sprintf("/accounts/%s/registrar/contacts/%s", accountID, contactID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return RegistrantContact{}, err
	}

	var r RegistrarContactDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return RegistrantContact{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}

// RegistrarDomainContacts gets a list of all Contacts based on the account and domain ID
//
// API reference: undocumented
func (api *API) RegistrarDomainContacts(ctx context.Context, accountID, domainName string) ([]RegistrantContact, error) {
	uri := fmt.Sprintf("/accounts/%s/registrar/domains/%s/contacts", accountID, domainName)
	return api.getContacts(ctx, uri)
}
