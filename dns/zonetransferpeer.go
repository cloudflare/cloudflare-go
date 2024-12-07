// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ZoneTransferPeerService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneTransferPeerService] method instead.
type ZoneTransferPeerService struct {
	Options []option.RequestOption
}

// NewZoneTransferPeerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneTransferPeerService(opts ...option.RequestOption) (r *ZoneTransferPeerService) {
	r = &ZoneTransferPeerService{}
	r.Options = opts
	return
}

// Create Peer.
func (r *ZoneTransferPeerService) New(ctx context.Context, params ZoneTransferPeerNewParams, opts ...option.RequestOption) (res *Peer, err error) {
	var env ZoneTransferPeerNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify Peer.
func (r *ZoneTransferPeerService) Update(ctx context.Context, peerID string, params ZoneTransferPeerUpdateParams, opts ...option.RequestOption) (res *Peer, err error) {
	var env ZoneTransferPeerUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if peerID == "" {
		err = errors.New("missing required peer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers/%s", params.AccountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Peers.
func (r *ZoneTransferPeerService) List(ctx context.Context, query ZoneTransferPeerListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Peer], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *ZoneTransferPeerService) ListAutoPaging(ctx context.Context, query ZoneTransferPeerListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Peer] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete Peer.
func (r *ZoneTransferPeerService) Delete(ctx context.Context, peerID string, body ZoneTransferPeerDeleteParams, opts ...option.RequestOption) (res *ZoneTransferPeerDeleteResponse, err error) {
	var env ZoneTransferPeerDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if peerID == "" {
		err = errors.New("missing required peer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers/%s", body.AccountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Peer.
func (r *ZoneTransferPeerService) Get(ctx context.Context, peerID string, query ZoneTransferPeerGetParams, opts ...option.RequestOption) (res *Peer, err error) {
	var env ZoneTransferPeerGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if peerID == "" {
		err = errors.New("missing required peer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers/%s", query.AccountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Peer struct {
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
	TSIGID string   `json:"tsig_id"`
	JSON   peerJSON `json:"-"`
}

// peerJSON contains the JSON metadata for the struct [Peer]
type peerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TSIGID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Peer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerJSON) RawJSON() string {
	return r.raw
}

type PeerParam struct {
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

func (r PeerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneTransferPeerDeleteResponse struct {
	ID   string                             `json:"id"`
	JSON zoneTransferPeerDeleteResponseJSON `json:"-"`
}

// zoneTransferPeerDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneTransferPeerDeleteResponse]
type zoneTransferPeerDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferPeerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferPeerDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferPeerNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the peer.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneTransferPeerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneTransferPeerNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferPeerNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Peer                                       `json:"result"`
	JSON    zoneTransferPeerNewResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferPeerNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferPeerNewResponseEnvelope]
type zoneTransferPeerNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferPeerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferPeerNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferPeerNewResponseEnvelopeSuccess bool

const (
	ZoneTransferPeerNewResponseEnvelopeSuccessTrue ZoneTransferPeerNewResponseEnvelopeSuccess = true
)

func (r ZoneTransferPeerNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferPeerNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferPeerUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Peer      PeerParam           `json:"peer,required"`
}

func (r ZoneTransferPeerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Peer)
}

type ZoneTransferPeerUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferPeerUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Peer                                          `json:"result"`
	JSON    zoneTransferPeerUpdateResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferPeerUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferPeerUpdateResponseEnvelope]
type zoneTransferPeerUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferPeerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferPeerUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferPeerUpdateResponseEnvelopeSuccess bool

const (
	ZoneTransferPeerUpdateResponseEnvelopeSuccessTrue ZoneTransferPeerUpdateResponseEnvelopeSuccess = true
)

func (r ZoneTransferPeerUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferPeerUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferPeerListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZoneTransferPeerDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZoneTransferPeerDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferPeerDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferPeerDeleteResponse                `json:"result"`
	JSON    zoneTransferPeerDeleteResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferPeerDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferPeerDeleteResponseEnvelope]
type zoneTransferPeerDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferPeerDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferPeerDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferPeerDeleteResponseEnvelopeSuccess bool

const (
	ZoneTransferPeerDeleteResponseEnvelopeSuccessTrue ZoneTransferPeerDeleteResponseEnvelopeSuccess = true
)

func (r ZoneTransferPeerDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferPeerDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferPeerGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZoneTransferPeerGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferPeerGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Peer                                       `json:"result"`
	JSON    zoneTransferPeerGetResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferPeerGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferPeerGetResponseEnvelope]
type zoneTransferPeerGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferPeerGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferPeerGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferPeerGetResponseEnvelopeSuccess bool

const (
	ZoneTransferPeerGetResponseEnvelopeSuccessTrue ZoneTransferPeerGetResponseEnvelopeSuccess = true
)

func (r ZoneTransferPeerGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferPeerGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
