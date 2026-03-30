// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_sending

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/email_routing"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// SubdomainDNSService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubdomainDNSService] method instead.
type SubdomainDNSService struct {
	Options []option.RequestOption
}

// NewSubdomainDNSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubdomainDNSService(opts ...option.RequestOption) (r *SubdomainDNSService) {
	r = &SubdomainDNSService{}
	r.Options = opts
	return
}

// Returns the expected DNS records for a sending subdomain.
func (r *SubdomainDNSService) Get(ctx context.Context, subdomainID string, query SubdomainDNSGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[email_routing.DNSRecord], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if subdomainID == "" {
		err = errors.New("missing required subdomain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/email/sending/subdomains/%s/dns", query.ZoneID, subdomainID)
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

// Returns the expected DNS records for a sending subdomain.
func (r *SubdomainDNSService) GetAutoPaging(ctx context.Context, subdomainID string, query SubdomainDNSGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[email_routing.DNSRecord] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, subdomainID, query, opts...))
}

type SubdomainDNSGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}
