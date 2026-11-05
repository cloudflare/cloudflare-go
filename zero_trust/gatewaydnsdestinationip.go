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

// GatewayDNSDestinationIPService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayDNSDestinationIPService] method instead.
type GatewayDNSDestinationIPService struct {
	Options []option.RequestOption
}

// NewGatewayDNSDestinationIPService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGatewayDNSDestinationIPService(opts ...option.RequestOption) (r *GatewayDNSDestinationIPService) {
	r = &GatewayDNSDestinationIPService{}
	r.Options = opts
	return
}

// List Zero Trust Gateway IPv4 address pairs for an account.
func (r *GatewayDNSDestinationIPService) List(ctx context.Context, query GatewayDNSDestinationIPListParams, opts ...option.RequestOption) (res *pagination.SinglePage[GatewayDNSDestinationIPListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/gateway/dns_destination_ips", query.AccountID)
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

// List Zero Trust Gateway IPv4 address pairs for an account.
func (r *GatewayDNSDestinationIPService) ListAutoPaging(ctx context.Context, query GatewayDNSDestinationIPListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[GatewayDNSDestinationIPListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type DNSDestinationIPsID = string

type GatewayDNSDestinationIPListResponse struct {
	ID       string `json:"id" api:"required"`
	BackupIP string `json:"backup_ip" api:"required"`
	// Specify whether the pair shared across multiple accounts (shared) or available
	// exclusively to this account. Non-shared pairs can contain Cloudflare-owned IPs
	// (dedicated) or customer-provided IPs (byoip).
	PairType  GatewayDNSDestinationIPListResponsePairType `json:"pair_type" api:"required"`
	PrimaryIP string                                      `json:"primary_ip" api:"required"`
	JSON      gatewayDNSDestinationIPListResponseJSON     `json:"-"`
}

// gatewayDNSDestinationIPListResponseJSON contains the JSON metadata for the
// struct [GatewayDNSDestinationIPListResponse]
type gatewayDNSDestinationIPListResponseJSON struct {
	ID          apijson.Field
	BackupIP    apijson.Field
	PairType    apijson.Field
	PrimaryIP   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayDNSDestinationIPListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayDNSDestinationIPListResponseJSON) RawJSON() string {
	return r.raw
}

// Specify whether the pair shared across multiple accounts (shared) or available
// exclusively to this account. Non-shared pairs can contain Cloudflare-owned IPs
// (dedicated) or customer-provided IPs (byoip).
type GatewayDNSDestinationIPListResponsePairType string

const (
	GatewayDNSDestinationIPListResponsePairTypeShared    GatewayDNSDestinationIPListResponsePairType = "shared"
	GatewayDNSDestinationIPListResponsePairTypeDedicated GatewayDNSDestinationIPListResponsePairType = "dedicated"
	GatewayDNSDestinationIPListResponsePairTypeByoip     GatewayDNSDestinationIPListResponsePairType = "byoip"
)

func (r GatewayDNSDestinationIPListResponsePairType) IsKnown() bool {
	switch r {
	case GatewayDNSDestinationIPListResponsePairTypeShared, GatewayDNSDestinationIPListResponsePairTypeDedicated, GatewayDNSDestinationIPListResponsePairTypeByoip:
		return true
	}
	return false
}

type GatewayDNSDestinationIPListParams struct {
	// Specify the Cloudflare account identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
