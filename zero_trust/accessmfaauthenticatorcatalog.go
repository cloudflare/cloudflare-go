// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// AccessMfaAuthenticatorCatalogService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessMfaAuthenticatorCatalogService] method instead.
type AccessMfaAuthenticatorCatalogService struct {
	Options []option.RequestOption
}

// NewAccessMfaAuthenticatorCatalogService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessMfaAuthenticatorCatalogService(opts ...option.RequestOption) (r *AccessMfaAuthenticatorCatalogService) {
	r = &AccessMfaAuthenticatorCatalogService{}
	r.Options = opts
	return
}

// Returns a list of Authenticator Device AAGUIDs for MFA configuration.
func (r *AccessMfaAuthenticatorCatalogService) List(ctx context.Context, query AccessMfaAuthenticatorCatalogListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AccessMfaAuthenticatorCatalogListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/authenticator_device_aaguids", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a list of Authenticator Device AAGUIDs for MFA configuration.
func (r *AccessMfaAuthenticatorCatalogService) ListAutoPaging(ctx context.Context, query AccessMfaAuthenticatorCatalogListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AccessMfaAuthenticatorCatalogListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// A FIDO2 authenticator device AAGUID entry
type AccessMfaAuthenticatorCatalogListResponse struct {
	// The Authenticator Attestation GUID (AAGUID) uniquely identifying a FIDO2
	// authenticator model
	Aaguid string `json:"aaguid" api:"required" format:"uuid"`
	// The human-readable name of the FIDO2 authenticator
	Name string                                        `json:"name" api:"required"`
	JSON accessMfaAuthenticatorCatalogListResponseJSON `json:"-"`
}

// accessMfaAuthenticatorCatalogListResponseJSON contains the JSON metadata for the
// struct [AccessMfaAuthenticatorCatalogListResponse]
type accessMfaAuthenticatorCatalogListResponseJSON struct {
	Aaguid      apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessMfaAuthenticatorCatalogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessMfaAuthenticatorCatalogListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessMfaAuthenticatorCatalogListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
