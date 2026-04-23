// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// TunnelWARPConnectorConnectionService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelWARPConnectorConnectionService] method instead.
type TunnelWARPConnectorConnectionService struct {
	Options []option.RequestOption
}

// NewTunnelWARPConnectorConnectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTunnelWARPConnectorConnectionService(opts ...option.RequestOption) (r *TunnelWARPConnectorConnectionService) {
	r = &TunnelWARPConnectorConnectionService{}
	r.Options = opts
	return
}

// Fetches connection details for a WARP Connector Tunnel.
func (r *TunnelWARPConnectorConnectionService) Get(ctx context.Context, tunnelID string, query TunnelWARPConnectorConnectionGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[TunnelWARPConnectorConnectionGetResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/warp_connector/%s/connections", query.AccountID, tunnelID)
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

// Fetches connection details for a WARP Connector Tunnel.
func (r *TunnelWARPConnectorConnectionService) GetAutoPaging(ctx context.Context, tunnelID string, query TunnelWARPConnectorConnectionGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TunnelWARPConnectorConnectionGetResponse] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, tunnelID, query, opts...))
}

// A WARP Connector client that maintains a connection to a Cloudflare data center.
type TunnelWARPConnectorConnectionGetResponse struct {
	// UUID of the Cloudflare Tunnel connector.
	ID string `json:"id" format:"uuid"`
	// The cloudflared OS architecture used to establish this connection.
	Arch string `json:"arch"`
	// The WARP Connector Tunnel connections between your origin and Cloudflare's edge.
	Conns []TunnelWARPConnectorConnectionGetResponseConn `json:"conns"`
	// Features enabled for the Cloudflare Tunnel.
	Features []string `json:"features"`
	// The HA status of a WARP Connector client.
	HaStatus TunnelWARPConnectorConnectionGetResponseHaStatus `json:"ha_status"`
	// Timestamp of when the tunnel connection was started.
	RunAt time.Time `json:"run_at" format:"date-time"`
	// The cloudflared version used to establish this connection.
	Version string                                       `json:"version"`
	JSON    tunnelWARPConnectorConnectionGetResponseJSON `json:"-"`
}

// tunnelWARPConnectorConnectionGetResponseJSON contains the JSON metadata for the
// struct [TunnelWARPConnectorConnectionGetResponse]
type tunnelWARPConnectorConnectionGetResponseJSON struct {
	ID          apijson.Field
	Arch        apijson.Field
	Conns       apijson.Field
	Features    apijson.Field
	HaStatus    apijson.Field
	RunAt       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConnectionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConnectionGetResponseJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConnectionGetResponseConn struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID string `json:"client_id" format:"uuid"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running WARP Connector.
	OriginIP string                                           `json:"origin_ip"`
	JSON     tunnelWARPConnectorConnectionGetResponseConnJSON `json:"-"`
}

// tunnelWARPConnectorConnectionGetResponseConnJSON contains the JSON metadata for
// the struct [TunnelWARPConnectorConnectionGetResponseConn]
type tunnelWARPConnectorConnectionGetResponseConnJSON struct {
	ID            apijson.Field
	ClientID      apijson.Field
	ClientVersion apijson.Field
	ColoName      apijson.Field
	OpenedAt      apijson.Field
	OriginIP      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelWARPConnectorConnectionGetResponseConn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConnectionGetResponseConnJSON) RawJSON() string {
	return r.raw
}

// The HA status of a WARP Connector client.
type TunnelWARPConnectorConnectionGetResponseHaStatus string

const (
	TunnelWARPConnectorConnectionGetResponseHaStatusOffline TunnelWARPConnectorConnectionGetResponseHaStatus = "offline"
	TunnelWARPConnectorConnectionGetResponseHaStatusPassive TunnelWARPConnectorConnectionGetResponseHaStatus = "passive"
	TunnelWARPConnectorConnectionGetResponseHaStatusActive  TunnelWARPConnectorConnectionGetResponseHaStatus = "active"
)

func (r TunnelWARPConnectorConnectionGetResponseHaStatus) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorConnectionGetResponseHaStatusOffline, TunnelWARPConnectorConnectionGetResponseHaStatusPassive, TunnelWARPConnectorConnectionGetResponseHaStatusActive:
		return true
	}
	return false
}

type TunnelWARPConnectorConnectionGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
