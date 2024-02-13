// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// CfdTunnelConnectionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCfdTunnelConnectionService]
// method instead.
type CfdTunnelConnectionService struct {
	Options []option.RequestOption
}

// NewCfdTunnelConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCfdTunnelConnectionService(opts ...option.RequestOption) (r *CfdTunnelConnectionService) {
	r = &CfdTunnelConnectionService{}
	r.Options = opts
	return
}

// Removes a connection (aka Cloudflare Tunnel Connector) from a Cloudflare Tunnel
// independently of its current state. If no connector id (client_id) is provided
// all connectors will be removed. We recommend running this command after rotating
// tokens.
func (r *CfdTunnelConnectionService) Delete(ctx context.Context, accountID string, tunnelID string, params CfdTunnelConnectionDeleteParams, opts ...option.RequestOption) (res *CfdTunnelConnectionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfdTunnelConnectionDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connections", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches connection details for a Cloudflare Tunnel.
func (r *CfdTunnelConnectionService) CloudflareTunnelListCloudflareTunnelConnections(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *[]CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connections", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [CfdTunnelConnectionDeleteResponseUnknown],
// [CfdTunnelConnectionDeleteResponseArray] or [shared.UnionString].
type CfdTunnelConnectionDeleteResponse interface {
	ImplementsCfdTunnelConnectionDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CfdTunnelConnectionDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CfdTunnelConnectionDeleteResponseArray []interface{}

func (r CfdTunnelConnectionDeleteResponseArray) ImplementsCfdTunnelConnectionDeleteResponse() {}

// A client (typically cloudflared) that maintains connections to a Cloudflare data
// center.
type CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponse struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// The cloudflared OS architecture used to establish this connection.
	Arch string `json:"arch"`
	// The version of the remote tunnel configuration. Used internally to sync
	// cloudflared with the Zero Trust dashboard.
	ConfigVersion int64 `json:"config_version"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Conns []CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseConn `json:"conns"`
	// Features enabled for the Cloudflare Tunnel.
	Features []string `json:"features"`
	// Timestamp of when the tunnel connection was started.
	RunAt time.Time `json:"run_at" format:"date-time"`
	// The cloudflared version used to establish this connection.
	Version string                                                                         `json:"version"`
	JSON    cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseJSON `json:"-"`
}

// cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseJSON
// contains the JSON metadata for the struct
// [CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponse]
type cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseJSON struct {
	ID            apijson.Field
	Arch          apijson.Field
	ConfigVersion apijson.Field
	Conns         apijson.Field
	Features      apijson.Field
	RunAt         apijson.Field
	Version       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseConn struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                                             `json:"uuid"`
	JSON cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseConnJSON `json:"-"`
}

// cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseConnJSON
// contains the JSON metadata for the struct
// [CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseConn]
type cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseConnJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseConn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelConnectionDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
	// UUID of the Cloudflare Tunnel Connector to disconnect.
	ClientID param.Field[string] `query:"client_id"`
}

func (r CfdTunnelConnectionDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [CfdTunnelConnectionDeleteParams]'s query parameters as
// `url.Values`.
func (r CfdTunnelConnectionDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CfdTunnelConnectionDeleteResponseEnvelope struct {
	Errors   []CfdTunnelConnectionDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CfdTunnelConnectionDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CfdTunnelConnectionDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CfdTunnelConnectionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfdTunnelConnectionDeleteResponseEnvelopeJSON    `json:"-"`
}

// cfdTunnelConnectionDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CfdTunnelConnectionDeleteResponseEnvelope]
type cfdTunnelConnectionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelConnectionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelConnectionDeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    cfdTunnelConnectionDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// cfdTunnelConnectionDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CfdTunnelConnectionDeleteResponseEnvelopeErrors]
type cfdTunnelConnectionDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelConnectionDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelConnectionDeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    cfdTunnelConnectionDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// cfdTunnelConnectionDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CfdTunnelConnectionDeleteResponseEnvelopeMessages]
type cfdTunnelConnectionDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelConnectionDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CfdTunnelConnectionDeleteResponseEnvelopeSuccess bool

const (
	CfdTunnelConnectionDeleteResponseEnvelopeSuccessTrue CfdTunnelConnectionDeleteResponseEnvelopeSuccess = true
)

type CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelope struct {
	Errors   []CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeMessages `json:"messages,required"`
	Result   []CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeJSON       `json:"-"`
}

// cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelope]
type cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeErrors struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeErrorsJSON `json:"-"`
}

// cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeErrors]
type cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeMessages struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeMessagesJSON `json:"-"`
}

// cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeMessages]
type cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeSuccess bool

const (
	CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeSuccessTrue CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeSuccess = true
)

type CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                          `json:"total_count"`
	JSON       cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeResultInfoJSON `json:"-"`
}

// cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeResultInfo]
type cfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelConnectionCloudflareTunnelListCloudflareTunnelConnectionsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
