// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountSecondaryDNSPeerService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountSecondaryDNSPeerService] method instead.
type AccountSecondaryDNSPeerService struct {
	Options []option.RequestOption
}

// NewAccountSecondaryDNSPeerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountSecondaryDNSPeerService(opts ...option.RequestOption) (r *AccountSecondaryDNSPeerService) {
	r = &AccountSecondaryDNSPeerService{}
	r.Options = opts
	return
}

// Get Peer.
func (r *AccountSecondaryDNSPeerService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *AccountSecondaryDNSPeerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify Peer.
func (r *AccountSecondaryDNSPeerService) Update(ctx context.Context, accountIdentifier interface{}, identifier interface{}, body AccountSecondaryDNSPeerUpdateParams, opts ...option.RequestOption) (res *AccountSecondaryDNSPeerUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete Peer.
func (r *AccountSecondaryDNSPeerService) Delete(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *AccountSecondaryDNSPeerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Peer.
func (r *AccountSecondaryDNSPeerService) SecondaryDNSPeerNewPeer(ctx context.Context, accountIdentifier interface{}, body AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerParams, opts ...option.RequestOption) (res *AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Peers.
func (r *AccountSecondaryDNSPeerService) SecondaryDNSPeerListPeers(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountSecondaryDNSPeerGetResponse struct {
	Errors   []AccountSecondaryDNSPeerGetResponseError   `json:"errors"`
	Messages []AccountSecondaryDNSPeerGetResponseMessage `json:"messages"`
	Result   AccountSecondaryDNSPeerGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSecondaryDNSPeerGetResponseSuccess `json:"success"`
	JSON    accountSecondaryDNSPeerGetResponseJSON    `json:"-"`
}

// accountSecondaryDNSPeerGetResponseJSON contains the JSON metadata for the struct
// [AccountSecondaryDNSPeerGetResponse]
type accountSecondaryDNSPeerGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerGetResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountSecondaryDNSPeerGetResponseErrorJSON `json:"-"`
}

// accountSecondaryDNSPeerGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountSecondaryDNSPeerGetResponseError]
type accountSecondaryDNSPeerGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerGetResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountSecondaryDNSPeerGetResponseMessageJSON `json:"-"`
}

// accountSecondaryDNSPeerGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountSecondaryDNSPeerGetResponseMessage]
type accountSecondaryDNSPeerGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerGetResponseResult struct {
	ID interface{} `json:"id,required"`
	// The name of the peer.
	Name string `json:"name,required"`
	// IPv4/IPv6 address of primary or secondary nameserver, depending on what zone
	// this peer is linked to. For primary zones this IP defines the IP of the
	// secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary
	// zones this IP defines the IP of the primary nameserver Cloudflare will send
	// AXFR/IXFR requests to.
	IP string `json:"ip"`
	// Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary
	// zones.
	IxfrEnable bool `json:"ixfr_enable"`
	// DNS port of primary or secondary nameserver, depending on what zone this peer is
	// linked to.
	Port float64 `json:"port"`
	// TSIG authentication will be used for zone transfer if configured.
	TsigID string                                       `json:"tsig_id"`
	JSON   accountSecondaryDNSPeerGetResponseResultJSON `json:"-"`
}

// accountSecondaryDNSPeerGetResponseResultJSON contains the JSON metadata for the
// struct [AccountSecondaryDNSPeerGetResponseResult]
type accountSecondaryDNSPeerGetResponseResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TsigID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDNSPeerGetResponseSuccess bool

const (
	AccountSecondaryDNSPeerGetResponseSuccessTrue AccountSecondaryDNSPeerGetResponseSuccess = true
)

type AccountSecondaryDNSPeerUpdateResponse struct {
	Errors   []AccountSecondaryDNSPeerUpdateResponseError   `json:"errors"`
	Messages []AccountSecondaryDNSPeerUpdateResponseMessage `json:"messages"`
	Result   AccountSecondaryDNSPeerUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSecondaryDNSPeerUpdateResponseSuccess `json:"success"`
	JSON    accountSecondaryDNSPeerUpdateResponseJSON    `json:"-"`
}

// accountSecondaryDNSPeerUpdateResponseJSON contains the JSON metadata for the
// struct [AccountSecondaryDNSPeerUpdateResponse]
type accountSecondaryDNSPeerUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerUpdateResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountSecondaryDNSPeerUpdateResponseErrorJSON `json:"-"`
}

// accountSecondaryDNSPeerUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountSecondaryDNSPeerUpdateResponseError]
type accountSecondaryDNSPeerUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerUpdateResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountSecondaryDNSPeerUpdateResponseMessageJSON `json:"-"`
}

// accountSecondaryDNSPeerUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountSecondaryDNSPeerUpdateResponseMessage]
type accountSecondaryDNSPeerUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerUpdateResponseResult struct {
	ID interface{} `json:"id,required"`
	// The name of the peer.
	Name string `json:"name,required"`
	// IPv4/IPv6 address of primary or secondary nameserver, depending on what zone
	// this peer is linked to. For primary zones this IP defines the IP of the
	// secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary
	// zones this IP defines the IP of the primary nameserver Cloudflare will send
	// AXFR/IXFR requests to.
	IP string `json:"ip"`
	// Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary
	// zones.
	IxfrEnable bool `json:"ixfr_enable"`
	// DNS port of primary or secondary nameserver, depending on what zone this peer is
	// linked to.
	Port float64 `json:"port"`
	// TSIG authentication will be used for zone transfer if configured.
	TsigID string                                          `json:"tsig_id"`
	JSON   accountSecondaryDNSPeerUpdateResponseResultJSON `json:"-"`
}

// accountSecondaryDNSPeerUpdateResponseResultJSON contains the JSON metadata for
// the struct [AccountSecondaryDNSPeerUpdateResponseResult]
type accountSecondaryDNSPeerUpdateResponseResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TsigID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDNSPeerUpdateResponseSuccess bool

const (
	AccountSecondaryDNSPeerUpdateResponseSuccessTrue AccountSecondaryDNSPeerUpdateResponseSuccess = true
)

type AccountSecondaryDNSPeerDeleteResponse struct {
	Errors   []AccountSecondaryDNSPeerDeleteResponseError   `json:"errors"`
	Messages []AccountSecondaryDNSPeerDeleteResponseMessage `json:"messages"`
	Result   AccountSecondaryDNSPeerDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSecondaryDNSPeerDeleteResponseSuccess `json:"success"`
	JSON    accountSecondaryDNSPeerDeleteResponseJSON    `json:"-"`
}

// accountSecondaryDNSPeerDeleteResponseJSON contains the JSON metadata for the
// struct [AccountSecondaryDNSPeerDeleteResponse]
type accountSecondaryDNSPeerDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerDeleteResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountSecondaryDNSPeerDeleteResponseErrorJSON `json:"-"`
}

// accountSecondaryDNSPeerDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountSecondaryDNSPeerDeleteResponseError]
type accountSecondaryDNSPeerDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerDeleteResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountSecondaryDNSPeerDeleteResponseMessageJSON `json:"-"`
}

// accountSecondaryDNSPeerDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountSecondaryDNSPeerDeleteResponseMessage]
type accountSecondaryDNSPeerDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerDeleteResponseResult struct {
	ID   interface{}                                     `json:"id"`
	JSON accountSecondaryDNSPeerDeleteResponseResultJSON `json:"-"`
}

// accountSecondaryDNSPeerDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountSecondaryDNSPeerDeleteResponseResult]
type accountSecondaryDNSPeerDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDNSPeerDeleteResponseSuccess bool

const (
	AccountSecondaryDNSPeerDeleteResponseSuccessTrue AccountSecondaryDNSPeerDeleteResponseSuccess = true
)

type AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponse struct {
	Errors   []AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseError   `json:"errors"`
	Messages []AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseMessage `json:"messages"`
	Result   AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseSuccess `json:"success"`
	JSON    accountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseJSON    `json:"-"`
}

// accountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseJSON contains the JSON
// metadata for the struct [AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponse]
type accountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseError struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseErrorJSON `json:"-"`
}

// accountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseError]
type accountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseMessage struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseMessageJSON `json:"-"`
}

// accountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseMessage]
type accountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseResult struct {
	ID interface{} `json:"id,required"`
	// The name of the peer.
	Name string `json:"name,required"`
	// IPv4/IPv6 address of primary or secondary nameserver, depending on what zone
	// this peer is linked to. For primary zones this IP defines the IP of the
	// secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary
	// zones this IP defines the IP of the primary nameserver Cloudflare will send
	// AXFR/IXFR requests to.
	IP string `json:"ip"`
	// Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary
	// zones.
	IxfrEnable bool `json:"ixfr_enable"`
	// DNS port of primary or secondary nameserver, depending on what zone this peer is
	// linked to.
	Port float64 `json:"port"`
	// TSIG authentication will be used for zone transfer if configured.
	TsigID string                                                           `json:"tsig_id"`
	JSON   accountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseResultJSON `json:"-"`
}

// accountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseResultJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseResult]
type accountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TsigID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseSuccess bool

const (
	AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseSuccessTrue AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerResponseSuccess = true
)

type AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponse struct {
	Errors     []AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseError    `json:"errors"`
	Messages   []AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseMessage  `json:"messages"`
	Result     []AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResult   `json:"result"`
	ResultInfo AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseSuccess `json:"success"`
	JSON    accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseJSON    `json:"-"`
}

// accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseJSON contains the JSON
// metadata for the struct
// [AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponse]
type accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseErrorJSON `json:"-"`
}

// accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseError]
type accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseMessageJSON `json:"-"`
}

// accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseMessage]
type accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResult struct {
	ID interface{} `json:"id,required"`
	// The name of the peer.
	Name string `json:"name,required"`
	// IPv4/IPv6 address of primary or secondary nameserver, depending on what zone
	// this peer is linked to. For primary zones this IP defines the IP of the
	// secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary
	// zones this IP defines the IP of the primary nameserver Cloudflare will send
	// AXFR/IXFR requests to.
	IP string `json:"ip"`
	// Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary
	// zones.
	IxfrEnable bool `json:"ixfr_enable"`
	// DNS port of primary or secondary nameserver, depending on what zone this peer is
	// linked to.
	Port float64 `json:"port"`
	// TSIG authentication will be used for zone transfer if configured.
	TsigID string                                                             `json:"tsig_id"`
	JSON   accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResultJSON `json:"-"`
}

// accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResultJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResult]
type accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TsigID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                `json:"total_count"`
	JSON       accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResultInfoJSON `json:"-"`
}

// accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResultInfoJSON contains
// the JSON metadata for the struct
// [AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResultInfo]
type accountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseSuccess bool

const (
	AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseSuccessTrue AccountSecondaryDNSPeerSecondaryDNSPeerListPeersResponseSuccess = true
)

type AccountSecondaryDNSPeerUpdateParams struct {
	// The name of the peer.
	Name param.Field[string] `json:"name,required"`
	// IPv4/IPv6 address of primary or secondary nameserver, depending on what zone
	// this peer is linked to. For primary zones this IP defines the IP of the
	// secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary
	// zones this IP defines the IP of the primary nameserver Cloudflare will send
	// AXFR/IXFR requests to.
	IP param.Field[string] `json:"ip"`
	// Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary
	// zones.
	IxfrEnable param.Field[bool] `json:"ixfr_enable"`
	// DNS port of primary or secondary nameserver, depending on what zone this peer is
	// linked to.
	Port param.Field[float64] `json:"port"`
	// TSIG authentication will be used for zone transfer if configured.
	TsigID param.Field[string] `json:"tsig_id"`
}

func (r AccountSecondaryDNSPeerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
