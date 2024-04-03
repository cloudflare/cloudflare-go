// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// PeerService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPeerService] method instead.
type PeerService struct {
	Options []option.RequestOption
}

// NewPeerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPeerService(opts ...option.RequestOption) (r *PeerService) {
	r = &PeerService{}
	r.Options = opts
	return
}

// Create Peer.
func (r *PeerService) New(ctx context.Context, params PeerNewParams, opts ...option.RequestOption) (res *SecondaryDNSPeer, err error) {
	opts = append(r.Options[:], opts...)
	var env PeerNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify Peer.
func (r *PeerService) Update(ctx context.Context, peerID string, params PeerUpdateParams, opts ...option.RequestOption) (res *SecondaryDNSPeer, err error) {
	opts = append(r.Options[:], opts...)
	var env PeerUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers/%s", params.AccountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Peers.
func (r *PeerService) List(ctx context.Context, query PeerListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SecondaryDNSPeer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// List Peers.
func (r *PeerService) ListAutoPaging(ctx context.Context, query PeerListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SecondaryDNSPeer] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete Peer.
func (r *PeerService) Delete(ctx context.Context, peerID string, params PeerDeleteParams, opts ...option.RequestOption) (res *PeerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PeerDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers/%s", params.AccountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Peer.
func (r *PeerService) Get(ctx context.Context, peerID string, query PeerGetParams, opts ...option.RequestOption) (res *SecondaryDNSPeer, err error) {
	opts = append(r.Options[:], opts...)
	var env PeerGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers/%s", query.AccountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSPeer struct {
	ID string `json:"id,required"`
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

func (r secondaryDNSPeerJSON) RawJSON() string {
	return r.raw
}

type PeerDeleteResponse struct {
	ID   string                 `json:"id"`
	JSON peerDeleteResponseJSON `json:"-"`
}

// peerDeleteResponseJSON contains the JSON metadata for the struct
// [PeerDeleteResponse]
type peerDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PeerNewParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r PeerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type PeerNewResponseEnvelope struct {
	Errors   []PeerNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PeerNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSPeer                  `json:"result,required"`
	// Whether the API call was successful
	Success PeerNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    peerNewResponseEnvelopeJSON    `json:"-"`
}

// peerNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PeerNewResponseEnvelope]
type peerNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PeerNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    peerNewResponseEnvelopeErrorsJSON `json:"-"`
}

// peerNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PeerNewResponseEnvelopeErrors]
type peerNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PeerNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    peerNewResponseEnvelopeMessagesJSON `json:"-"`
}

// peerNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PeerNewResponseEnvelopeMessages]
type peerNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PeerNewResponseEnvelopeSuccess bool

const (
	PeerNewResponseEnvelopeSuccessTrue PeerNewResponseEnvelopeSuccess = true
)

func (r PeerNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PeerNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PeerUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
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

func (r PeerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PeerUpdateResponseEnvelope struct {
	Errors   []PeerUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PeerUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSPeer                     `json:"result,required"`
	// Whether the API call was successful
	Success PeerUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    peerUpdateResponseEnvelopeJSON    `json:"-"`
}

// peerUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PeerUpdateResponseEnvelope]
type peerUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PeerUpdateResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    peerUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// peerUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PeerUpdateResponseEnvelopeErrors]
type peerUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PeerUpdateResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    peerUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// peerUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PeerUpdateResponseEnvelopeMessages]
type peerUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PeerUpdateResponseEnvelopeSuccess bool

const (
	PeerUpdateResponseEnvelopeSuccessTrue PeerUpdateResponseEnvelopeSuccess = true
)

func (r PeerUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PeerUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PeerListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type PeerDeleteParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r PeerDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type PeerDeleteResponseEnvelope struct {
	Errors   []PeerDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PeerDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   PeerDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PeerDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    peerDeleteResponseEnvelopeJSON    `json:"-"`
}

// peerDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [PeerDeleteResponseEnvelope]
type peerDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PeerDeleteResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    peerDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// peerDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PeerDeleteResponseEnvelopeErrors]
type peerDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PeerDeleteResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    peerDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// peerDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PeerDeleteResponseEnvelopeMessages]
type peerDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PeerDeleteResponseEnvelopeSuccess bool

const (
	PeerDeleteResponseEnvelopeSuccessTrue PeerDeleteResponseEnvelopeSuccess = true
)

func (r PeerDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PeerDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PeerGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type PeerGetResponseEnvelope struct {
	Errors   []PeerGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PeerGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSPeer                  `json:"result,required"`
	// Whether the API call was successful
	Success PeerGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    peerGetResponseEnvelopeJSON    `json:"-"`
}

// peerGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PeerGetResponseEnvelope]
type peerGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PeerGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    peerGetResponseEnvelopeErrorsJSON `json:"-"`
}

// peerGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PeerGetResponseEnvelopeErrors]
type peerGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PeerGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    peerGetResponseEnvelopeMessagesJSON `json:"-"`
}

// peerGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PeerGetResponseEnvelopeMessages]
type peerGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PeerGetResponseEnvelopeSuccess bool

const (
	PeerGetResponseEnvelopeSuccessTrue PeerGetResponseEnvelopeSuccess = true
)

func (r PeerGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PeerGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
