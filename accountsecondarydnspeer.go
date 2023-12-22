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
func (r *AccountSecondaryDNSPeerService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *SchemasSingleResponseNZNnIuCk, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify Peer.
func (r *AccountSecondaryDNSPeerService) Update(ctx context.Context, accountIdentifier interface{}, identifier interface{}, body AccountSecondaryDNSPeerUpdateParams, opts ...option.RequestOption) (res *SchemasSingleResponseNZNnIuCk, err error) {
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
func (r *AccountSecondaryDNSPeerService) SecondaryDNSPeerNewPeer(ctx context.Context, accountIdentifier interface{}, body AccountSecondaryDNSPeerSecondaryDNSPeerNewPeerParams, opts ...option.RequestOption) (res *SchemasSingleResponseNZNnIuCk, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Peers.
func (r *AccountSecondaryDNSPeerService) SecondaryDNSPeerListPeers(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *SchemasResponseCollectionSCt9wUKt, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SchemasResponseCollectionSCt9wUKt struct {
	Errors     []SchemasResponseCollectionSCt9wUKtError    `json:"errors"`
	Messages   []SchemasResponseCollectionSCt9wUKtMessage  `json:"messages"`
	Result     []SchemasResponseCollectionSCt9wUKtResult   `json:"result"`
	ResultInfo SchemasResponseCollectionSCt9wUKtResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SchemasResponseCollectionSCt9wUKtSuccess `json:"success"`
	JSON    schemasResponseCollectionSCt9wUKtJSON    `json:"-"`
}

// schemasResponseCollectionSCt9wUKtJSON contains the JSON metadata for the struct
// [SchemasResponseCollectionSCt9wUKt]
type schemasResponseCollectionSCt9wUKtJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionSCt9wUKt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionSCt9wUKtError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    schemasResponseCollectionSCt9wUKtErrorJSON `json:"-"`
}

// schemasResponseCollectionSCt9wUKtErrorJSON contains the JSON metadata for the
// struct [SchemasResponseCollectionSCt9wUKtError]
type schemasResponseCollectionSCt9wUKtErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionSCt9wUKtError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionSCt9wUKtMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    schemasResponseCollectionSCt9wUKtMessageJSON `json:"-"`
}

// schemasResponseCollectionSCt9wUKtMessageJSON contains the JSON metadata for the
// struct [SchemasResponseCollectionSCt9wUKtMessage]
type schemasResponseCollectionSCt9wUKtMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionSCt9wUKtMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionSCt9wUKtResult struct {
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
	TsigID string                                      `json:"tsig_id"`
	JSON   schemasResponseCollectionSCt9wUKtResultJSON `json:"-"`
}

// schemasResponseCollectionSCt9wUKtResultJSON contains the JSON metadata for the
// struct [SchemasResponseCollectionSCt9wUKtResult]
type schemasResponseCollectionSCt9wUKtResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TsigID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionSCt9wUKtResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionSCt9wUKtResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       schemasResponseCollectionSCt9wUKtResultInfoJSON `json:"-"`
}

// schemasResponseCollectionSCt9wUKtResultInfoJSON contains the JSON metadata for
// the struct [SchemasResponseCollectionSCt9wUKtResultInfo]
type schemasResponseCollectionSCt9wUKtResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionSCt9wUKtResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasResponseCollectionSCt9wUKtSuccess bool

const (
	SchemasResponseCollectionSCt9wUKtSuccessTrue SchemasResponseCollectionSCt9wUKtSuccess = true
)

type SchemasSingleResponseNZNnIuCk struct {
	Errors   []SchemasSingleResponseNZNnIuCkError   `json:"errors"`
	Messages []SchemasSingleResponseNZNnIuCkMessage `json:"messages"`
	Result   SchemasSingleResponseNZNnIuCkResult    `json:"result"`
	// Whether the API call was successful
	Success SchemasSingleResponseNZNnIuCkSuccess `json:"success"`
	JSON    schemasSingleResponseNzNnIuCkJSON    `json:"-"`
}

// schemasSingleResponseNzNnIuCkJSON contains the JSON metadata for the struct
// [SchemasSingleResponseNZNnIuCk]
type schemasSingleResponseNzNnIuCkJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseNZNnIuCk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseNZNnIuCkError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    schemasSingleResponseNzNnIuCkErrorJSON `json:"-"`
}

// schemasSingleResponseNzNnIuCkErrorJSON contains the JSON metadata for the struct
// [SchemasSingleResponseNZNnIuCkError]
type schemasSingleResponseNzNnIuCkErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseNZNnIuCkError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseNZNnIuCkMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    schemasSingleResponseNzNnIuCkMessageJSON `json:"-"`
}

// schemasSingleResponseNzNnIuCkMessageJSON contains the JSON metadata for the
// struct [SchemasSingleResponseNZNnIuCkMessage]
type schemasSingleResponseNzNnIuCkMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseNZNnIuCkMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseNZNnIuCkResult struct {
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
	TsigID string                                  `json:"tsig_id"`
	JSON   schemasSingleResponseNzNnIuCkResultJSON `json:"-"`
}

// schemasSingleResponseNzNnIuCkResultJSON contains the JSON metadata for the
// struct [SchemasSingleResponseNZNnIuCkResult]
type schemasSingleResponseNzNnIuCkResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TsigID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseNZNnIuCkResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasSingleResponseNZNnIuCkSuccess bool

const (
	SchemasSingleResponseNZNnIuCkSuccessTrue SchemasSingleResponseNZNnIuCkSuccess = true
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
