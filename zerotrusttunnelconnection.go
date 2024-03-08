// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// ZeroTrustTunnelConnectionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustTunnelConnectionService] method instead.
type ZeroTrustTunnelConnectionService struct {
	Options []option.RequestOption
}

// NewZeroTrustTunnelConnectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustTunnelConnectionService(opts ...option.RequestOption) (r *ZeroTrustTunnelConnectionService) {
	r = &ZeroTrustTunnelConnectionService{}
	r.Options = opts
	return
}

// Removes connections that are in a disconnected or pending reconnect state. We
// recommend running this command after shutting down a tunnel.
func (r *ZeroTrustTunnelConnectionService) Delete(ctx context.Context, tunnelID string, params ZeroTrustTunnelConnectionDeleteParams, opts ...option.RequestOption) (res *ZeroTrustTunnelConnectionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustTunnelConnectionDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels/%s/connections", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches connection details for a Cloudflare Tunnel.
func (r *ZeroTrustTunnelConnectionService) Get(ctx context.Context, tunnelID string, query ZeroTrustTunnelConnectionGetParams, opts ...option.RequestOption) (res *[]ZeroTrustTunnelConnectionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustTunnelConnectionGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connections", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ZeroTrustTunnelConnectionDeleteResponseUnknown],
// [ZeroTrustTunnelConnectionDeleteResponseArray] or [shared.UnionString].
type ZeroTrustTunnelConnectionDeleteResponse interface {
	ImplementsZeroTrustTunnelConnectionDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustTunnelConnectionDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustTunnelConnectionDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustTunnelConnectionDeleteResponseArray []interface{}

func (r ZeroTrustTunnelConnectionDeleteResponseArray) ImplementsZeroTrustTunnelConnectionDeleteResponse() {
}

// A client (typically cloudflared) that maintains connections to a Cloudflare data
// center.
type ZeroTrustTunnelConnectionGetResponse struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// The cloudflared OS architecture used to establish this connection.
	Arch string `json:"arch"`
	// The version of the remote tunnel configuration. Used internally to sync
	// cloudflared with the Zero Trust dashboard.
	ConfigVersion int64 `json:"config_version"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Conns []ZeroTrustTunnelConnectionGetResponseConn `json:"conns"`
	// Features enabled for the Cloudflare Tunnel.
	Features []string `json:"features"`
	// Timestamp of when the tunnel connection was started.
	RunAt time.Time `json:"run_at" format:"date-time"`
	// The cloudflared version used to establish this connection.
	Version string                                   `json:"version"`
	JSON    zeroTrustTunnelConnectionGetResponseJSON `json:"-"`
}

// zeroTrustTunnelConnectionGetResponseJSON contains the JSON metadata for the
// struct [ZeroTrustTunnelConnectionGetResponse]
type zeroTrustTunnelConnectionGetResponseJSON struct {
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

func (r *ZeroTrustTunnelConnectionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelConnectionGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelConnectionGetResponseConn struct {
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
	UUID string                                       `json:"uuid"`
	JSON zeroTrustTunnelConnectionGetResponseConnJSON `json:"-"`
}

// zeroTrustTunnelConnectionGetResponseConnJSON contains the JSON metadata for the
// struct [ZeroTrustTunnelConnectionGetResponseConn]
type zeroTrustTunnelConnectionGetResponseConnJSON struct {
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

func (r *ZeroTrustTunnelConnectionGetResponseConn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelConnectionGetResponseConnJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelConnectionDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ZeroTrustTunnelConnectionDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZeroTrustTunnelConnectionDeleteResponseEnvelope struct {
	Errors   []ZeroTrustTunnelConnectionDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustTunnelConnectionDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustTunnelConnectionDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustTunnelConnectionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustTunnelConnectionDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustTunnelConnectionDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustTunnelConnectionDeleteResponseEnvelope]
type zeroTrustTunnelConnectionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConnectionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelConnectionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelConnectionDeleteResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustTunnelConnectionDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustTunnelConnectionDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustTunnelConnectionDeleteResponseEnvelopeErrors]
type zeroTrustTunnelConnectionDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConnectionDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelConnectionDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelConnectionDeleteResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustTunnelConnectionDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustTunnelConnectionDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustTunnelConnectionDeleteResponseEnvelopeMessages]
type zeroTrustTunnelConnectionDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConnectionDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelConnectionDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustTunnelConnectionDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustTunnelConnectionDeleteResponseEnvelopeSuccessTrue ZeroTrustTunnelConnectionDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustTunnelConnectionGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustTunnelConnectionGetResponseEnvelope struct {
	Errors   []ZeroTrustTunnelConnectionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustTunnelConnectionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustTunnelConnectionGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustTunnelConnectionGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustTunnelConnectionGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustTunnelConnectionGetResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustTunnelConnectionGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustTunnelConnectionGetResponseEnvelope]
type zeroTrustTunnelConnectionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConnectionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelConnectionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelConnectionGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustTunnelConnectionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustTunnelConnectionGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustTunnelConnectionGetResponseEnvelopeErrors]
type zeroTrustTunnelConnectionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConnectionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelConnectionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelConnectionGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustTunnelConnectionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustTunnelConnectionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustTunnelConnectionGetResponseEnvelopeMessages]
type zeroTrustTunnelConnectionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConnectionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelConnectionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustTunnelConnectionGetResponseEnvelopeSuccess bool

const (
	ZeroTrustTunnelConnectionGetResponseEnvelopeSuccessTrue ZeroTrustTunnelConnectionGetResponseEnvelopeSuccess = true
)

type ZeroTrustTunnelConnectionGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       zeroTrustTunnelConnectionGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustTunnelConnectionGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustTunnelConnectionGetResponseEnvelopeResultInfo]
type zeroTrustTunnelConnectionGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelConnectionGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelConnectionGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
