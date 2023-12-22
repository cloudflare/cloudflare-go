// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountTunnelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountTunnelService] method
// instead.
type AccountTunnelService struct {
	Options     []option.RequestOption
	Connections *AccountTunnelConnectionService
}

// NewAccountTunnelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountTunnelService(opts ...option.RequestOption) (r *AccountTunnelService) {
	r = &AccountTunnelService{}
	r.Options = opts
	r.Connections = NewAccountTunnelConnectionService(opts...)
	return
}

// Fetches a single Argo Tunnel.
func (r *AccountTunnelService) Get(ctx context.Context, accountIdentifier string, tunnelID string, opts ...option.RequestOption) (res *TunnelResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/tunnels/%s", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an Argo Tunnel from an account.
func (r *AccountTunnelService) Delete(ctx context.Context, accountIdentifier string, tunnelID string, body AccountTunnelDeleteParams, opts ...option.RequestOption) (res *TunnelResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/tunnels/%s", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Creates a new Argo Tunnel in an account.
func (r *AccountTunnelService) ArgoTunnelNewAnArgoTunnel(ctx context.Context, accountIdentifier string, body AccountTunnelArgoTunnelNewAnArgoTunnelParams, opts ...option.RequestOption) (res *TunnelResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists and filters Argo Tunnels in an account.
func (r *AccountTunnelService) ArgoTunnelListArgoTunnels(ctx context.Context, accountIdentifier string, query AccountTunnelArgoTunnelListArgoTunnelsParams, opts ...option.RequestOption) (res *TunnelResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/tunnels", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TunnelResponseCollection struct {
	Errors     []TunnelResponseCollectionError    `json:"errors"`
	Messages   []TunnelResponseCollectionMessage  `json:"messages"`
	Result     []TunnelResponseCollectionResult   `json:"result"`
	ResultInfo TunnelResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success TunnelResponseCollectionSuccess `json:"success"`
	JSON    tunnelResponseCollectionJSON    `json:"-"`
}

// tunnelResponseCollectionJSON contains the JSON metadata for the struct
// [TunnelResponseCollection]
type tunnelResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelResponseCollectionError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    tunnelResponseCollectionErrorJSON `json:"-"`
}

// tunnelResponseCollectionErrorJSON contains the JSON metadata for the struct
// [TunnelResponseCollectionError]
type tunnelResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelResponseCollectionMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    tunnelResponseCollectionMessageJSON `json:"-"`
}

// tunnelResponseCollectionMessageJSON contains the JSON metadata for the struct
// [TunnelResponseCollectionMessage]
type tunnelResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelResponseCollectionResult struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelResponseCollectionResultConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time                          `json:"deleted_at,nullable" format:"date-time"`
	JSON      tunnelResponseCollectionResultJSON `json:"-"`
}

// tunnelResponseCollectionResultJSON contains the JSON metadata for the struct
// [TunnelResponseCollectionResult]
type tunnelResponseCollectionResultJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelResponseCollectionResultConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                       `json:"uuid"`
	JSON tunnelResponseCollectionResultConnectionJSON `json:"-"`
}

// tunnelResponseCollectionResultConnectionJSON contains the JSON metadata for the
// struct [TunnelResponseCollectionResultConnection]
type tunnelResponseCollectionResultConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelResponseCollectionResultConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       tunnelResponseCollectionResultInfoJSON `json:"-"`
}

// tunnelResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [TunnelResponseCollectionResultInfo]
type tunnelResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelResponseCollectionSuccess bool

const (
	TunnelResponseCollectionSuccessTrue TunnelResponseCollectionSuccess = true
)

type TunnelResponseSingle struct {
	Errors   []TunnelResponseSingleError   `json:"errors"`
	Messages []TunnelResponseSingleMessage `json:"messages"`
	Result   interface{}                   `json:"result"`
	// Whether the API call was successful
	Success TunnelResponseSingleSuccess `json:"success"`
	JSON    tunnelResponseSingleJSON    `json:"-"`
}

// tunnelResponseSingleJSON contains the JSON metadata for the struct
// [TunnelResponseSingle]
type tunnelResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelResponseSingleError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    tunnelResponseSingleErrorJSON `json:"-"`
}

// tunnelResponseSingleErrorJSON contains the JSON metadata for the struct
// [TunnelResponseSingleError]
type tunnelResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelResponseSingleMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    tunnelResponseSingleMessageJSON `json:"-"`
}

// tunnelResponseSingleMessageJSON contains the JSON metadata for the struct
// [TunnelResponseSingleMessage]
type tunnelResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelResponseSingleSuccess bool

const (
	TunnelResponseSingleSuccessTrue TunnelResponseSingleSuccess = true
)

type AccountTunnelDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountTunnelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountTunnelArgoTunnelNewAnArgoTunnelParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name,required"`
	// Sets the password required to run the tunnel. Must be at least 32 bytes and
	// encoded as a base64 string.
	TunnelSecret param.Field[interface{}] `json:"tunnel_secret,required"`
}

func (r AccountTunnelArgoTunnelNewAnArgoTunnelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountTunnelArgoTunnelListArgoTunnelsParams struct {
	// If provided, include only tunnels that were created (and not deleted) before
	// this time.
	ExistedAt param.Field[time.Time] `query:"existed_at" format:"date-time"`
	// If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If
	// empty, all tunnels will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// A user-friendly name for the tunnel.
	TunnelName param.Field[string] `query:"tunnel_name"`
}

// URLQuery serializes [AccountTunnelArgoTunnelListArgoTunnelsParams]'s query
// parameters as `url.Values`.
func (r AccountTunnelArgoTunnelListArgoTunnelsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
