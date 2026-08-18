// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// DNSSECZskService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSSECZskService] method instead.
type DNSSECZskService struct {
	Options []option.RequestOption
}

// NewDNSSECZskService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSSECZskService(opts ...option.RequestOption) (r *DNSSECZskService) {
	r = &DNSSECZskService{}
	r.Options = opts
	return
}

// List the Zone Signing Keys (ZSKs) that DNSSEC uses for the zone.
func (r *DNSSECZskService) List(ctx context.Context, query DNSSECZskListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DNSSECZskListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/dnssec/zsk", query.ZoneID)
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

// List the Zone Signing Keys (ZSKs) that DNSSEC uses for the zone.
func (r *DNSSECZskService) ListAutoPaging(ctx context.Context, query DNSSECZskListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DNSSECZskListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type DNSSECZskListResponse struct {
	DNSKEY DNSSECZskListResponseDNSKEY `json:"DNSKEY"`
	// Storage backend where the DNSSEC key material is stored.
	Location DNSSECZskListResponseLocation `json:"Location"`
	// Internal key name for the ZSK.
	Name       string                          `json:"Name"`
	SigningKey DNSSECZskListResponseSigningKey `json:"SigningKey"`
	// Lifecycle state tag attached to the DNSSEC key.
	Tag  DNSSECZskListResponseTag  `json:"Tag"`
	JSON dnssecZskListResponseJSON `json:"-"`
}

// dnssecZskListResponseJSON contains the JSON metadata for the struct
// [DNSSECZskListResponse]
type dnssecZskListResponseJSON struct {
	DNSKEY      apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	SigningKey  apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSECZskListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnssecZskListResponseJSON) RawJSON() string {
	return r.raw
}

type DNSSECZskListResponseDNSKEY struct {
	Algorithm int64                           `json:"Algorithm" api:"nullable"`
	Flags     int64                           `json:"Flags" api:"nullable"`
	Hdr       DNSSECZskListResponseDNSKEYHdr  `json:"Hdr"`
	Protocol  int64                           `json:"Protocol" api:"nullable"`
	PublicKey string                          `json:"PublicKey" api:"nullable"`
	JSON      dnssecZskListResponseDNSKEYJSON `json:"-"`
}

// dnssecZskListResponseDNSKEYJSON contains the JSON metadata for the struct
// [DNSSECZskListResponseDNSKEY]
type dnssecZskListResponseDNSKEYJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Hdr         apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSECZskListResponseDNSKEY) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnssecZskListResponseDNSKEYJSON) RawJSON() string {
	return r.raw
}

type DNSSECZskListResponseDNSKEYHdr struct {
	Class    int64                              `json:"Class" api:"nullable"`
	Name     string                             `json:"Name" api:"nullable"`
	Rdlength int64                              `json:"Rdlength" api:"nullable"`
	Rrtype   int64                              `json:"Rrtype" api:"nullable"`
	TTL      int64                              `json:"Ttl" api:"nullable"`
	JSON     dnssecZskListResponseDNSKEYHdrJSON `json:"-"`
}

// dnssecZskListResponseDNSKEYHdrJSON contains the JSON metadata for the struct
// [DNSSECZskListResponseDNSKEYHdr]
type dnssecZskListResponseDNSKEYHdrJSON struct {
	Class       apijson.Field
	Name        apijson.Field
	Rdlength    apijson.Field
	Rrtype      apijson.Field
	TTL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSECZskListResponseDNSKEYHdr) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnssecZskListResponseDNSKEYHdrJSON) RawJSON() string {
	return r.raw
}

// Storage backend where the DNSSEC key material is stored.
type DNSSECZskListResponseLocation string

const (
	DNSSECZskListResponseLocationDatabase DNSSECZskListResponseLocation = "database"
	DNSSECZskListResponseLocationVault    DNSSECZskListResponseLocation = "vault"
)

func (r DNSSECZskListResponseLocation) IsKnown() bool {
	switch r {
	case DNSSECZskListResponseLocationDatabase, DNSSECZskListResponseLocationVault:
		return true
	}
	return false
}

type DNSSECZskListResponseSigningKey struct {
	// Key encryption key name used to encrypt the private key.
	Kek string `json:"kek" api:"nullable"`
	// Encrypted private key material for the signing key.
	Privkey string `json:"privkey" api:"nullable" format:"byte"`
	// Public key content associated with the signing key.
	Pubkey string                              `json:"pubkey" api:"nullable"`
	JSON   dnssecZskListResponseSigningKeyJSON `json:"-"`
}

// dnssecZskListResponseSigningKeyJSON contains the JSON metadata for the struct
// [DNSSECZskListResponseSigningKey]
type dnssecZskListResponseSigningKeyJSON struct {
	Kek         apijson.Field
	Privkey     apijson.Field
	Pubkey      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSECZskListResponseSigningKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnssecZskListResponseSigningKeyJSON) RawJSON() string {
	return r.raw
}

// Lifecycle state tag attached to the DNSSEC key.
type DNSSECZskListResponseTag string

const (
	DNSSECZskListResponseTagActive   DNSSECZskListResponseTag = "active"
	DNSSECZskListResponseTagPublish  DNSSECZskListResponseTag = "publish"
	DNSSECZskListResponseTagExternal DNSSECZskListResponseTag = "external"
	DNSSECZskListResponseTagRetired  DNSSECZskListResponseTag = "retired"
	DNSSECZskListResponseTagRevoked  DNSSECZskListResponseTag = "revoked"
	DNSSECZskListResponseTagRemoved  DNSSECZskListResponseTag = "removed"
)

func (r DNSSECZskListResponseTag) IsKnown() bool {
	switch r {
	case DNSSECZskListResponseTagActive, DNSSECZskListResponseTagPublish, DNSSECZskListResponseTagExternal, DNSSECZskListResponseTagRetired, DNSSECZskListResponseTagRevoked, DNSSECZskListResponseTagRemoved:
		return true
	}
	return false
}

type DNSSECZskListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}
