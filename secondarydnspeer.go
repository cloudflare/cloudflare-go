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

// SecondaryDNSPeerService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSecondaryDNSPeerService] method
// instead.
type SecondaryDNSPeerService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSPeerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSecondaryDNSPeerService(opts ...option.RequestOption) (r *SecondaryDNSPeerService) {
	r = &SecondaryDNSPeerService{}
	r.Options = opts
	return
}

// Create Peer.
func (r *SecondaryDNSPeerService) New(ctx context.Context, params SecondaryDNSPeerNewParams, opts ...option.RequestOption) (res *SecondaryDNSPeer, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSPeerNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify Peer.
func (r *SecondaryDNSPeerService) Update(ctx context.Context, peerID interface{}, params SecondaryDNSPeerUpdateParams, opts ...option.RequestOption) (res *SecondaryDNSPeer, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSPeerUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers/%v", params.AccountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Peers.
func (r *SecondaryDNSPeerService) List(ctx context.Context, query SecondaryDNSPeerListParams, opts ...option.RequestOption) (res *[]SecondaryDNSPeer, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSPeerListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Peer.
func (r *SecondaryDNSPeerService) Delete(ctx context.Context, peerID interface{}, body SecondaryDNSPeerDeleteParams, opts ...option.RequestOption) (res *SecondaryDNSPeerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSPeerDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers/%v", body.AccountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Peer.
func (r *SecondaryDNSPeerService) Get(ctx context.Context, peerID interface{}, query SecondaryDNSPeerGetParams, opts ...option.RequestOption) (res *SecondaryDNSPeer, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSPeerGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers/%v", query.AccountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSPeer struct {
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
	TSIGID string               `json:"tsig_id"`
	JSON   secondaryDNSPeerJSON `json:"-"`
}

// secondaryDNSPeerJSON contains the JSON metadata for the struct
// [SecondaryDNSPeer]
type secondaryDNSPeerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TSIGID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerDeleteResponse struct {
	ID   interface{}                        `json:"id"`
	JSON secondaryDNSPeerDeleteResponseJSON `json:"-"`
}

// secondaryDNSPeerDeleteResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSPeerDeleteResponse]
type secondaryDNSPeerDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r SecondaryDNSPeerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SecondaryDNSPeerNewResponseEnvelope struct {
	Errors   []SecondaryDNSPeerNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSPeerNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSPeer                              `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSPeerNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSPeerNewResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSPeerNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSPeerNewResponseEnvelope]
type secondaryDNSPeerNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    secondaryDNSPeerNewResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSPeerNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SecondaryDNSPeerNewResponseEnvelopeErrors]
type secondaryDNSPeerNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    secondaryDNSPeerNewResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSPeerNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDNSPeerNewResponseEnvelopeMessages]
type secondaryDNSPeerNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSPeerNewResponseEnvelopeSuccess bool

const (
	SecondaryDNSPeerNewResponseEnvelopeSuccessTrue SecondaryDNSPeerNewResponseEnvelopeSuccess = true
)

type SecondaryDNSPeerUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
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
	TSIGID param.Field[string] `json:"tsig_id"`
}

func (r SecondaryDNSPeerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSPeerUpdateResponseEnvelope struct {
	Errors   []SecondaryDNSPeerUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSPeerUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSPeer                                 `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSPeerUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSPeerUpdateResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSPeerUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSPeerUpdateResponseEnvelope]
type secondaryDNSPeerUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    secondaryDNSPeerUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSPeerUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDNSPeerUpdateResponseEnvelopeErrors]
type secondaryDNSPeerUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    secondaryDNSPeerUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSPeerUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SecondaryDNSPeerUpdateResponseEnvelopeMessages]
type secondaryDNSPeerUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSPeerUpdateResponseEnvelopeSuccess bool

const (
	SecondaryDNSPeerUpdateResponseEnvelopeSuccessTrue SecondaryDNSPeerUpdateResponseEnvelopeSuccess = true
)

type SecondaryDNSPeerListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type SecondaryDNSPeerListResponseEnvelope struct {
	Errors   []SecondaryDNSPeerListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSPeerListResponseEnvelopeMessages `json:"messages,required"`
	Result   []SecondaryDNSPeer                             `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SecondaryDNSPeerListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SecondaryDNSPeerListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       secondaryDNSPeerListResponseEnvelopeJSON       `json:"-"`
}

// secondaryDNSPeerListResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSPeerListResponseEnvelope]
type secondaryDNSPeerListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    secondaryDNSPeerListResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSPeerListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDNSPeerListResponseEnvelopeErrors]
type secondaryDNSPeerListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    secondaryDNSPeerListResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSPeerListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDNSPeerListResponseEnvelopeMessages]
type secondaryDNSPeerListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSPeerListResponseEnvelopeSuccess bool

const (
	SecondaryDNSPeerListResponseEnvelopeSuccessTrue SecondaryDNSPeerListResponseEnvelopeSuccess = true
)

type SecondaryDNSPeerListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       secondaryDNSPeerListResponseEnvelopeResultInfoJSON `json:"-"`
}

// secondaryDNSPeerListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [SecondaryDNSPeerListResponseEnvelopeResultInfo]
type secondaryDNSPeerListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type SecondaryDNSPeerDeleteResponseEnvelope struct {
	Errors   []SecondaryDNSPeerDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSPeerDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSPeerDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSPeerDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSPeerDeleteResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSPeerDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSPeerDeleteResponseEnvelope]
type secondaryDNSPeerDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    secondaryDNSPeerDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSPeerDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDNSPeerDeleteResponseEnvelopeErrors]
type secondaryDNSPeerDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    secondaryDNSPeerDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSPeerDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SecondaryDNSPeerDeleteResponseEnvelopeMessages]
type secondaryDNSPeerDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSPeerDeleteResponseEnvelopeSuccess bool

const (
	SecondaryDNSPeerDeleteResponseEnvelopeSuccessTrue SecondaryDNSPeerDeleteResponseEnvelopeSuccess = true
)

type SecondaryDNSPeerGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type SecondaryDNSPeerGetResponseEnvelope struct {
	Errors   []SecondaryDNSPeerGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSPeerGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSPeer                              `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSPeerGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSPeerGetResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSPeerGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSPeerGetResponseEnvelope]
type secondaryDNSPeerGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    secondaryDNSPeerGetResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSPeerGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SecondaryDNSPeerGetResponseEnvelopeErrors]
type secondaryDNSPeerGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    secondaryDNSPeerGetResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSPeerGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SecondaryDNSPeerGetResponseEnvelopeMessages]
type secondaryDNSPeerGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSPeerGetResponseEnvelopeSuccess bool

const (
	SecondaryDNSPeerGetResponseEnvelopeSuccessTrue SecondaryDNSPeerGetResponseEnvelopeSuccess = true
)
