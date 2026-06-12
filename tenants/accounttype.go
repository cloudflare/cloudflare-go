// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tenants

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// AccountTypeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountTypeService] method instead.
type AccountTypeService struct {
	Options []option.RequestOption
}

// NewAccountTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountTypeService(opts ...option.RequestOption) (r *AccountTypeService) {
	r = &AccountTypeService{}
	r.Options = opts
	return
}

// List of account types available for the Tenant to provision accounts.
func (r *AccountTypeService) List(ctx context.Context, tenantID string, opts ...option.RequestOption) (res *pagination.SinglePage[string], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if tenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tenants/%s/account_types", tenantID)
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

// List of account types available for the Tenant to provision accounts.
func (r *AccountTypeService) ListAutoPaging(ctx context.Context, tenantID string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[string] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, tenantID, opts...))
}
