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

// Modify Peer.
func (r *SecondaryDNSPeerService) Update(ctx context.Context, accountIdentifier interface{}, identifier interface{}, body SecondaryDNSPeerUpdateParams, opts ...option.RequestOption) (res *SecondaryDNSPeerUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSPeerUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Peer.
func (r *SecondaryDNSPeerService) Delete(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *SecondaryDNSPeerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSPeerDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Peer.
func (r *SecondaryDNSPeerService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *SecondaryDNSPeerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSPeerGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create Peer.
func (r *SecondaryDNSPeerService) SecondaryDNSPeerNewPeer(ctx context.Context, accountIdentifier interface{}, body SecondaryDNSPeerSecondaryDNSPeerNewPeerParams, opts ...option.RequestOption) (res *SecondaryDNSPeerSecondaryDNSPeerNewPeerResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Peers.
func (r *SecondaryDNSPeerService) SecondaryDNSPeerListPeers(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *[]SecondaryDNSPeerSecondaryDNSPeerListPeersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelope
	path := fmt.Sprintf("accounts/%v/secondary_dns/peers", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSPeerUpdateResponse struct {
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
	TsigID string                             `json:"tsig_id"`
	JSON   secondaryDNSPeerUpdateResponseJSON `json:"-"`
}

// secondaryDNSPeerUpdateResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSPeerUpdateResponse]
type secondaryDNSPeerUpdateResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TsigID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
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

type SecondaryDNSPeerGetResponse struct {
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
	TsigID string                          `json:"tsig_id"`
	JSON   secondaryDNSPeerGetResponseJSON `json:"-"`
}

// secondaryDNSPeerGetResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSPeerGetResponse]
type secondaryDNSPeerGetResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TsigID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerSecondaryDNSPeerNewPeerResponse struct {
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
	TsigID string                                              `json:"tsig_id"`
	JSON   secondaryDNSPeerSecondaryDNSPeerNewPeerResponseJSON `json:"-"`
}

// secondaryDNSPeerSecondaryDNSPeerNewPeerResponseJSON contains the JSON metadata
// for the struct [SecondaryDNSPeerSecondaryDNSPeerNewPeerResponse]
type secondaryDNSPeerSecondaryDNSPeerNewPeerResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TsigID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerSecondaryDNSPeerNewPeerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerSecondaryDNSPeerListPeersResponse struct {
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
	TsigID string                                                `json:"tsig_id"`
	JSON   secondaryDNSPeerSecondaryDNSPeerListPeersResponseJSON `json:"-"`
}

// secondaryDNSPeerSecondaryDNSPeerListPeersResponseJSON contains the JSON metadata
// for the struct [SecondaryDNSPeerSecondaryDNSPeerListPeersResponse]
type secondaryDNSPeerSecondaryDNSPeerListPeersResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TsigID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerSecondaryDNSPeerListPeersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerUpdateParams struct {
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

func (r SecondaryDNSPeerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSPeerUpdateResponseEnvelope struct {
	Errors   []SecondaryDNSPeerUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSPeerUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSPeerUpdateResponse                   `json:"result,required"`
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

type SecondaryDNSPeerGetResponseEnvelope struct {
	Errors   []SecondaryDNSPeerGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSPeerGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSPeerGetResponse                   `json:"result,required"`
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

type SecondaryDNSPeerSecondaryDNSPeerNewPeerParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r SecondaryDNSPeerSecondaryDNSPeerNewPeerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelope struct {
	Errors   []SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSPeerSecondaryDNSPeerNewPeerResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelope]
type secondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    secondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeErrors]
type secondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    secondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeMessages]
type secondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeSuccess bool

const (
	SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeSuccessTrue SecondaryDNSPeerSecondaryDNSPeerNewPeerResponseEnvelopeSuccess = true
)

type SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelope struct {
	Errors   []SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeMessages `json:"messages,required"`
	Result   []SecondaryDNSPeerSecondaryDNSPeerListPeersResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeResultInfo `json:"result_info"`
	JSON       secondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeJSON       `json:"-"`
}

// secondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelope]
type secondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeErrors struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    secondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeErrors]
type secondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeMessages struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    secondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeMessages]
type secondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeSuccess bool

const (
	SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeSuccessTrue SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeSuccess = true
)

type SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                 `json:"total_count"`
	JSON       secondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeResultInfoJSON `json:"-"`
}

// secondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeResultInfo]
type secondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSPeerSecondaryDNSPeerListPeersResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
