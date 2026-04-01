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
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// TunnelWARPConnectorConnectorService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelWARPConnectorConnectorService] method instead.
type TunnelWARPConnectorConnectorService struct {
	Options []option.RequestOption
}

// NewTunnelWARPConnectorConnectorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTunnelWARPConnectorConnectorService(opts ...option.RequestOption) (r *TunnelWARPConnectorConnectorService) {
	r = &TunnelWARPConnectorConnectorService{}
	r.Options = opts
	return
}

// Fetches connector and connection details for a WARP Connector Tunnel.
func (r *TunnelWARPConnectorConnectorService) Get(ctx context.Context, tunnelID string, connectorID string, query TunnelWARPConnectorConnectorGetParams, opts ...option.RequestOption) (res *TunnelWARPConnectorConnectorGetResponse, err error) {
	var env TunnelWARPConnectorConnectorGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return nil, err
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/warp_connector/%s/connectors/%s", query.AccountID, tunnelID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A WARP Connector client that maintains a connection to a Cloudflare data center.
type TunnelWARPConnectorConnectorGetResponse struct {
	// UUID of the Cloudflare Tunnel connector.
	ID string `json:"id" format:"uuid"`
	// The cloudflared OS architecture used to establish this connection.
	Arch string `json:"arch"`
	// The WARP Connector Tunnel connections between your origin and Cloudflare's edge.
	Conns []TunnelWARPConnectorConnectorGetResponseConn `json:"conns"`
	// Features enabled for the Cloudflare Tunnel.
	Features []string `json:"features"`
	// The HA status of a WARP Connector client.
	HaStatus TunnelWARPConnectorConnectorGetResponseHaStatus `json:"ha_status"`
	// Timestamp of when the tunnel connection was started.
	RunAt time.Time `json:"run_at" format:"date-time"`
	// The cloudflared version used to establish this connection.
	Version string                                      `json:"version"`
	JSON    tunnelWARPConnectorConnectorGetResponseJSON `json:"-"`
}

// tunnelWARPConnectorConnectorGetResponseJSON contains the JSON metadata for the
// struct [TunnelWARPConnectorConnectorGetResponse]
type tunnelWARPConnectorConnectorGetResponseJSON struct {
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

func (r *TunnelWARPConnectorConnectorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConnectorGetResponseJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConnectorGetResponseConn struct {
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
	OriginIP string                                          `json:"origin_ip"`
	JSON     tunnelWARPConnectorConnectorGetResponseConnJSON `json:"-"`
}

// tunnelWARPConnectorConnectorGetResponseConnJSON contains the JSON metadata for
// the struct [TunnelWARPConnectorConnectorGetResponseConn]
type tunnelWARPConnectorConnectorGetResponseConnJSON struct {
	ID            apijson.Field
	ClientID      apijson.Field
	ClientVersion apijson.Field
	ColoName      apijson.Field
	OpenedAt      apijson.Field
	OriginIP      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelWARPConnectorConnectorGetResponseConn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConnectorGetResponseConnJSON) RawJSON() string {
	return r.raw
}

// The HA status of a WARP Connector client.
type TunnelWARPConnectorConnectorGetResponseHaStatus string

const (
	TunnelWARPConnectorConnectorGetResponseHaStatusOffline TunnelWARPConnectorConnectorGetResponseHaStatus = "offline"
	TunnelWARPConnectorConnectorGetResponseHaStatusPassive TunnelWARPConnectorConnectorGetResponseHaStatus = "passive"
	TunnelWARPConnectorConnectorGetResponseHaStatusActive  TunnelWARPConnectorConnectorGetResponseHaStatus = "active"
)

func (r TunnelWARPConnectorConnectorGetResponseHaStatus) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorConnectorGetResponseHaStatusOffline, TunnelWARPConnectorConnectorGetResponseHaStatusPassive, TunnelWARPConnectorConnectorGetResponseHaStatusActive:
		return true
	}
	return false
}

type TunnelWARPConnectorConnectorGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type TunnelWARPConnectorConnectorGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// A WARP Connector client that maintains a connection to a Cloudflare data center.
	Result TunnelWARPConnectorConnectorGetResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success TunnelWARPConnectorConnectorGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    tunnelWARPConnectorConnectorGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelWARPConnectorConnectorGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [TunnelWARPConnectorConnectorGetResponseEnvelope]
type tunnelWARPConnectorConnectorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConnectorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConnectorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelWARPConnectorConnectorGetResponseEnvelopeSuccess bool

const (
	TunnelWARPConnectorConnectorGetResponseEnvelopeSuccessTrue TunnelWARPConnectorConnectorGetResponseEnvelopeSuccess = true
)

func (r TunnelWARPConnectorConnectorGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorConnectorGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
