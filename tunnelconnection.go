// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// TunnelConnectionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTunnelConnectionService] method
// instead.
type TunnelConnectionService struct {
	Options []option.RequestOption
}

// NewTunnelConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTunnelConnectionService(opts ...option.RequestOption) (r *TunnelConnectionService) {
	r = &TunnelConnectionService{}
	r.Options = opts
	return
}

// Fetches connection details for a Cloudflare Tunnel.
func (r *TunnelConnectionService) List(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *[]TunnelConnectionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelConnectionListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connections", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Removes connections that are in a disconnected or pending reconnect state. We
// recommend running this command after shutting down a tunnel.
func (r *TunnelConnectionService) Delete(ctx context.Context, accountID string, tunnelID string, body TunnelConnectionDeleteParams, opts ...option.RequestOption) (res *TunnelConnectionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelConnectionDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels/%s/connections", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A client (typically cloudflared) that maintains connections to a Cloudflare data
// center.
type TunnelConnectionListResponse struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// The cloudflared OS architecture used to establish this connection.
	Arch string `json:"arch"`
	// The version of the remote tunnel configuration. Used internally to sync
	// cloudflared with the Zero Trust dashboard.
	ConfigVersion int64 `json:"config_version"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Conns []TunnelConnectionListResponseConn `json:"conns"`
	// Features enabled for the Cloudflare Tunnel.
	Features []string `json:"features"`
	// Timestamp of when the tunnel connection was started.
	RunAt time.Time `json:"run_at" format:"date-time"`
	// The cloudflared version used to establish this connection.
	Version string                           `json:"version"`
	JSON    tunnelConnectionListResponseJSON `json:"-"`
}

// tunnelConnectionListResponseJSON contains the JSON metadata for the struct
// [TunnelConnectionListResponse]
type tunnelConnectionListResponseJSON struct {
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

func (r *TunnelConnectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionListResponseConn struct {
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
	UUID string                               `json:"uuid"`
	JSON tunnelConnectionListResponseConnJSON `json:"-"`
}

// tunnelConnectionListResponseConnJSON contains the JSON metadata for the struct
// [TunnelConnectionListResponseConn]
type tunnelConnectionListResponseConnJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelConnectionListResponseConn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [TunnelConnectionDeleteResponseUnknown],
// [TunnelConnectionDeleteResponseArray] or [shared.UnionString].
type TunnelConnectionDeleteResponse interface {
	ImplementsTunnelConnectionDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelConnectionDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TunnelConnectionDeleteResponseArray []interface{}

func (r TunnelConnectionDeleteResponseArray) ImplementsTunnelConnectionDeleteResponse() {}

type TunnelConnectionListResponseEnvelope struct {
	Errors   []TunnelConnectionListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelConnectionListResponseEnvelopeMessages `json:"messages,required"`
	Result   []TunnelConnectionListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    TunnelConnectionListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo TunnelConnectionListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       tunnelConnectionListResponseEnvelopeJSON       `json:"-"`
}

// tunnelConnectionListResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelConnectionListResponseEnvelope]
type tunnelConnectionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    tunnelConnectionListResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelConnectionListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [TunnelConnectionListResponseEnvelopeErrors]
type tunnelConnectionListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    tunnelConnectionListResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelConnectionListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [TunnelConnectionListResponseEnvelopeMessages]
type tunnelConnectionListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelConnectionListResponseEnvelopeSuccess bool

const (
	TunnelConnectionListResponseEnvelopeSuccessTrue TunnelConnectionListResponseEnvelopeSuccess = true
)

type TunnelConnectionListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       tunnelConnectionListResponseEnvelopeResultInfoJSON `json:"-"`
}

// tunnelConnectionListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [TunnelConnectionListResponseEnvelopeResultInfo]
type tunnelConnectionListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r TunnelConnectionDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type TunnelConnectionDeleteResponseEnvelope struct {
	Errors   []TunnelConnectionDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelConnectionDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelConnectionDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TunnelConnectionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelConnectionDeleteResponseEnvelopeJSON    `json:"-"`
}

// tunnelConnectionDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelConnectionDeleteResponseEnvelope]
type tunnelConnectionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    tunnelConnectionDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelConnectionDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [TunnelConnectionDeleteResponseEnvelopeErrors]
type tunnelConnectionDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    tunnelConnectionDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelConnectionDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [TunnelConnectionDeleteResponseEnvelopeMessages]
type tunnelConnectionDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelConnectionDeleteResponseEnvelopeSuccess bool

const (
	TunnelConnectionDeleteResponseEnvelopeSuccessTrue TunnelConnectionDeleteResponseEnvelopeSuccess = true
)
