// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moq

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// RelayService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRelayService] method instead.
type RelayService struct {
	Options []option.RequestOption
	Tokens  *RelayTokenService
}

// NewRelayService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRelayService(opts ...option.RequestOption) (r *RelayService) {
	r = &RelayService{}
	r.Options = opts
	r.Tokens = NewRelayTokenService(opts...)
	return
}

// Provisions a new MoQ relay instance. Auto-creates a publish+subscribe token and
// a subscribe-only token. Token values are included in the response (shown once).
// Config is always set to defaults (upstreams off) and cannot be supplied here —
// sending a non-empty `config` is rejected (21014); `null` or `{}` is accepted as
// absent. Use PUT to configure the relay after it exists.
func (r *RelayService) New(ctx context.Context, params RelayNewParams, opts ...option.RequestOption) (res *RelayNewResponse, err error) {
	var env RelayNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/moq/relays", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates a relay's name and/or configuration. The relay ID goes in the URL path —
// `PUT /accounts/{account_id}/moq/relays/{relay_id}` — not the request body; there
// is no collection-level update endpoint. This is also the only way to set a
// relay's config (config cannot be set at create time). Partial updates: omitted
// fields are preserved; config sub-objects replace as whole objects when present.
func (r *RelayService) Update(ctx context.Context, relayID string, params RelayUpdateParams, opts ...option.RequestOption) (res *RelayUpdateResponse, err error) {
	var env RelayUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if relayID == "" {
		err = errors.New("missing required relay_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/moq/relays/%s", params.AccountID, relayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all MoQ relays for the account. Returns only metadata. Config, status, and
// tokens are omitted.
//
// Results are cursor-paginated (keyset on the `created` timestamp). Use
// `created_before` / `created_after` with the `created` value of the first/last
// item in a page to fetch the adjacent page. `result_info` reports the page
// `count` and the `total` matching the cursor filters.
func (r *RelayService) List(ctx context.Context, params RelayListParams, opts ...option.RequestOption) (res *pagination.SinglePage[RelayListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/moq/relays", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Lists all MoQ relays for the account. Returns only metadata. Config, status, and
// tokens are omitted.
//
// Results are cursor-paginated (keyset on the `created` timestamp). Use
// `created_before` / `created_after` with the `created` value of the first/last
// item in a page to fetch the adjacent page. `result_info` reports the page
// `count` and the `total` matching the cursor filters.
func (r *RelayService) ListAutoPaging(ctx context.Context, params RelayListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RelayListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Soft-deletes a MoQ relay. The relay ID goes in the URL path —
// `DELETE /accounts/{account_id}/moq/relays/{relay_id}` — not the request body;
// there is no collection-level delete endpoint.
func (r *RelayService) Delete(ctx context.Context, relayID string, body RelayDeleteParams, opts ...option.RequestOption) (res *RelayDeleteResponse, err error) {
	var env RelayDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if relayID == "" {
		err = errors.New("missing required relay_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/moq/relays/%s", body.AccountID, relayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves a single MoQ relay including config and status. Tokens are NOT
// included.
func (r *RelayService) Get(ctx context.Context, relayID string, query RelayGetParams, opts ...option.RequestOption) (res *RelayGetResponse, err error) {
	var env RelayGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if relayID == "" {
		err = errors.New("missing required relay_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/moq/relays/%s", query.AccountID, relayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Relay with its auto-created default token pair (one full-access [publish,
// subscribe] and one [subscribe]-only), each with its one-time secret, wrapped in
// the issuers envelope.
type RelayNewResponse struct {
	Config  RelayNewResponseConfig `json:"config" api:"required"`
	Created time.Time              `json:"created" api:"required" format:"date-time"`
	// Token collection (discriminated union on `type`). On create this holds the
	// auto-created default pair, each including its one-time secret.
	Issuers  []RelayNewResponseIssuer `json:"issuers" api:"required"`
	Modified time.Time                `json:"modified" api:"required" format:"date-time"`
	Name     string                   `json:"name" api:"required"`
	// Server-generated unique identifier (32 hex chars).
	UID  string               `json:"uid" api:"required"`
	JSON relayNewResponseJSON `json:"-"`
}

// relayNewResponseJSON contains the JSON metadata for the struct
// [RelayNewResponse]
type relayNewResponseJSON struct {
	Config      apijson.Field
	Created     apijson.Field
	Issuers     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayNewResponseJSON) RawJSON() string {
	return r.raw
}

type RelayNewResponseConfig struct {
	// Upstreams are external MOQT server publishers that a relay falls back to when it
	// has no local publisher for a requested namespace/track.
	Upstreams RelayNewResponseConfigUpstreams `json:"upstreams"`
	JSON      relayNewResponseConfigJSON      `json:"-"`
}

// relayNewResponseConfigJSON contains the JSON metadata for the struct
// [RelayNewResponseConfig]
type relayNewResponseConfigJSON struct {
	Upstreams   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayNewResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayNewResponseConfigJSON) RawJSON() string {
	return r.raw
}

// Upstreams are external MOQT server publishers that a relay falls back to when it
// has no local publisher for a requested namespace/track.
type RelayNewResponseConfigUpstreams struct {
	Enabled bool `json:"enabled"`
	// Ordered list of upstream MOQT server publishers. Each entry is an object (not a
	// bare string) so per-upstream configuration can be added in the future without
	// another breaking change.
	Upstreams []RelayNewResponseConfigUpstreamsUpstream `json:"upstreams"`
	JSON      relayNewResponseConfigUpstreamsJSON       `json:"-"`
}

// relayNewResponseConfigUpstreamsJSON contains the JSON metadata for the struct
// [RelayNewResponseConfigUpstreams]
type relayNewResponseConfigUpstreamsJSON struct {
	Enabled     apijson.Field
	Upstreams   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayNewResponseConfigUpstreams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayNewResponseConfigUpstreamsJSON) RawJSON() string {
	return r.raw
}

// A single upstream MOQT server publisher.
type RelayNewResponseConfigUpstreamsUpstream struct {
	// Upstream MOQT server publisher URL. Must be an absolute URL with a host and a
	// scheme the relay can dial: moqt:// (raw QUIC) or https:// (WebTransport).
	// Validated on update (PUT); rejected with 21013.
	URL  string                                      `json:"url" api:"required" format:"uri"`
	JSON relayNewResponseConfigUpstreamsUpstreamJSON `json:"-"`
}

// relayNewResponseConfigUpstreamsUpstreamJSON contains the JSON metadata for the
// struct [RelayNewResponseConfigUpstreamsUpstream]
type relayNewResponseConfigUpstreamsUpstreamJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayNewResponseConfigUpstreamsUpstream) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayNewResponseConfigUpstreamsUpstreamJSON) RawJSON() string {
	return r.raw
}

// One arm of the discriminated-union token collection.
type RelayNewResponseIssuer struct {
	// Always present ([] when empty).
	CloudflareTokens []RelayNewResponseIssuersCloudflareToken `json:"cloudflare_tokens" api:"required"`
	Issuer           RelayNewResponseIssuersIssuer            `json:"issuer" api:"required"`
	Type             RelayNewResponseIssuersType              `json:"type" api:"required"`
	JSON             relayNewResponseIssuerJSON               `json:"-"`
}

// relayNewResponseIssuerJSON contains the JSON metadata for the struct
// [RelayNewResponseIssuer]
type relayNewResponseIssuerJSON struct {
	CloudflareTokens apijson.Field
	Issuer           apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RelayNewResponseIssuer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayNewResponseIssuerJSON) RawJSON() string {
	return r.raw
}

type RelayNewResponseIssuersCloudflareToken struct {
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Mandatory; no more than 1 year after `created`.
	Expires time.Time `json:"expires" api:"required" format:"date-time"`
	// Token identity and registry key (32 hex chars).
	Jti string `json:"jti" api:"required"`
	// Signed allowlist of what the token may do. V1 coarse roles; the array form
	// extends to fine-grained MoQT message names later without a breaking change.
	Operations []RelayNewResponseIssuersCloudflareTokensOperation `json:"operations" api:"required"`
	// Optional, customer-set.
	Label string `json:"label"`
	// The signed JWT. Present ONLY in create / auto-create responses (shown once);
	// never returned by list, never stored.
	Secret string                                     `json:"secret"`
	JSON   relayNewResponseIssuersCloudflareTokenJSON `json:"-"`
}

// relayNewResponseIssuersCloudflareTokenJSON contains the JSON metadata for the
// struct [RelayNewResponseIssuersCloudflareToken]
type relayNewResponseIssuersCloudflareTokenJSON struct {
	Created     apijson.Field
	Expires     apijson.Field
	Jti         apijson.Field
	Operations  apijson.Field
	Label       apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayNewResponseIssuersCloudflareToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayNewResponseIssuersCloudflareTokenJSON) RawJSON() string {
	return r.raw
}

type RelayNewResponseIssuersCloudflareTokensOperation string

const (
	RelayNewResponseIssuersCloudflareTokensOperationPublish   RelayNewResponseIssuersCloudflareTokensOperation = "publish"
	RelayNewResponseIssuersCloudflareTokensOperationSubscribe RelayNewResponseIssuersCloudflareTokensOperation = "subscribe"
)

func (r RelayNewResponseIssuersCloudflareTokensOperation) IsKnown() bool {
	switch r {
	case RelayNewResponseIssuersCloudflareTokensOperationPublish, RelayNewResponseIssuersCloudflareTokensOperationSubscribe:
		return true
	}
	return false
}

type RelayNewResponseIssuersIssuer string

const (
	RelayNewResponseIssuersIssuerCloudflare RelayNewResponseIssuersIssuer = "cloudflare"
)

func (r RelayNewResponseIssuersIssuer) IsKnown() bool {
	switch r {
	case RelayNewResponseIssuersIssuerCloudflare:
		return true
	}
	return false
}

type RelayNewResponseIssuersType string

const (
	RelayNewResponseIssuersTypeCloudflareJWT RelayNewResponseIssuersType = "cloudflare_jwt"
)

func (r RelayNewResponseIssuersType) IsKnown() bool {
	switch r {
	case RelayNewResponseIssuersTypeCloudflareJWT:
		return true
	}
	return false
}

// Full relay details (no tokens).
type RelayUpdateResponse struct {
	Config   RelayUpdateResponseConfig `json:"config" api:"required"`
	Created  time.Time                 `json:"created" api:"required" format:"date-time"`
	Modified time.Time                 `json:"modified" api:"required" format:"date-time"`
	Name     string                    `json:"name" api:"required"`
	UID      string                    `json:"uid" api:"required"`
	// "connected" when active, omitted otherwise.
	Status RelayUpdateResponseStatus `json:"status"`
	JSON   relayUpdateResponseJSON   `json:"-"`
}

// relayUpdateResponseJSON contains the JSON metadata for the struct
// [RelayUpdateResponse]
type relayUpdateResponseJSON struct {
	Config      apijson.Field
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RelayUpdateResponseConfig struct {
	// Upstreams are external MOQT server publishers that a relay falls back to when it
	// has no local publisher for a requested namespace/track.
	Upstreams RelayUpdateResponseConfigUpstreams `json:"upstreams"`
	JSON      relayUpdateResponseConfigJSON      `json:"-"`
}

// relayUpdateResponseConfigJSON contains the JSON metadata for the struct
// [RelayUpdateResponseConfig]
type relayUpdateResponseConfigJSON struct {
	Upstreams   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayUpdateResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayUpdateResponseConfigJSON) RawJSON() string {
	return r.raw
}

// Upstreams are external MOQT server publishers that a relay falls back to when it
// has no local publisher for a requested namespace/track.
type RelayUpdateResponseConfigUpstreams struct {
	Enabled bool `json:"enabled"`
	// Ordered list of upstream MOQT server publishers. Each entry is an object (not a
	// bare string) so per-upstream configuration can be added in the future without
	// another breaking change.
	Upstreams []RelayUpdateResponseConfigUpstreamsUpstream `json:"upstreams"`
	JSON      relayUpdateResponseConfigUpstreamsJSON       `json:"-"`
}

// relayUpdateResponseConfigUpstreamsJSON contains the JSON metadata for the struct
// [RelayUpdateResponseConfigUpstreams]
type relayUpdateResponseConfigUpstreamsJSON struct {
	Enabled     apijson.Field
	Upstreams   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayUpdateResponseConfigUpstreams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayUpdateResponseConfigUpstreamsJSON) RawJSON() string {
	return r.raw
}

// A single upstream MOQT server publisher.
type RelayUpdateResponseConfigUpstreamsUpstream struct {
	// Upstream MOQT server publisher URL. Must be an absolute URL with a host and a
	// scheme the relay can dial: moqt:// (raw QUIC) or https:// (WebTransport).
	// Validated on update (PUT); rejected with 21013.
	URL  string                                         `json:"url" api:"required" format:"uri"`
	JSON relayUpdateResponseConfigUpstreamsUpstreamJSON `json:"-"`
}

// relayUpdateResponseConfigUpstreamsUpstreamJSON contains the JSON metadata for
// the struct [RelayUpdateResponseConfigUpstreamsUpstream]
type relayUpdateResponseConfigUpstreamsUpstreamJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayUpdateResponseConfigUpstreamsUpstream) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayUpdateResponseConfigUpstreamsUpstreamJSON) RawJSON() string {
	return r.raw
}

// "connected" when active, omitted otherwise.
type RelayUpdateResponseStatus string

const (
	RelayUpdateResponseStatusConnected RelayUpdateResponseStatus = "connected"
)

func (r RelayUpdateResponseStatus) IsKnown() bool {
	switch r {
	case RelayUpdateResponseStatusConnected:
		return true
	}
	return false
}

// Abbreviated relay for list responses.
type RelayListResponse struct {
	Created  time.Time             `json:"created" api:"required" format:"date-time"`
	Modified time.Time             `json:"modified" api:"required" format:"date-time"`
	Name     string                `json:"name" api:"required"`
	UID      string                `json:"uid" api:"required"`
	JSON     relayListResponseJSON `json:"-"`
}

// relayListResponseJSON contains the JSON metadata for the struct
// [RelayListResponse]
type relayListResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayListResponseJSON) RawJSON() string {
	return r.raw
}

type RelayDeleteResponse = interface{}

// Full relay details (no tokens).
type RelayGetResponse struct {
	Config   RelayGetResponseConfig `json:"config" api:"required"`
	Created  time.Time              `json:"created" api:"required" format:"date-time"`
	Modified time.Time              `json:"modified" api:"required" format:"date-time"`
	Name     string                 `json:"name" api:"required"`
	UID      string                 `json:"uid" api:"required"`
	// "connected" when active, omitted otherwise.
	Status RelayGetResponseStatus `json:"status"`
	JSON   relayGetResponseJSON   `json:"-"`
}

// relayGetResponseJSON contains the JSON metadata for the struct
// [RelayGetResponse]
type relayGetResponseJSON struct {
	Config      apijson.Field
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayGetResponseJSON) RawJSON() string {
	return r.raw
}

type RelayGetResponseConfig struct {
	// Upstreams are external MOQT server publishers that a relay falls back to when it
	// has no local publisher for a requested namespace/track.
	Upstreams RelayGetResponseConfigUpstreams `json:"upstreams"`
	JSON      relayGetResponseConfigJSON      `json:"-"`
}

// relayGetResponseConfigJSON contains the JSON metadata for the struct
// [RelayGetResponseConfig]
type relayGetResponseConfigJSON struct {
	Upstreams   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayGetResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayGetResponseConfigJSON) RawJSON() string {
	return r.raw
}

// Upstreams are external MOQT server publishers that a relay falls back to when it
// has no local publisher for a requested namespace/track.
type RelayGetResponseConfigUpstreams struct {
	Enabled bool `json:"enabled"`
	// Ordered list of upstream MOQT server publishers. Each entry is an object (not a
	// bare string) so per-upstream configuration can be added in the future without
	// another breaking change.
	Upstreams []RelayGetResponseConfigUpstreamsUpstream `json:"upstreams"`
	JSON      relayGetResponseConfigUpstreamsJSON       `json:"-"`
}

// relayGetResponseConfigUpstreamsJSON contains the JSON metadata for the struct
// [RelayGetResponseConfigUpstreams]
type relayGetResponseConfigUpstreamsJSON struct {
	Enabled     apijson.Field
	Upstreams   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayGetResponseConfigUpstreams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayGetResponseConfigUpstreamsJSON) RawJSON() string {
	return r.raw
}

// A single upstream MOQT server publisher.
type RelayGetResponseConfigUpstreamsUpstream struct {
	// Upstream MOQT server publisher URL. Must be an absolute URL with a host and a
	// scheme the relay can dial: moqt:// (raw QUIC) or https:// (WebTransport).
	// Validated on update (PUT); rejected with 21013.
	URL  string                                      `json:"url" api:"required" format:"uri"`
	JSON relayGetResponseConfigUpstreamsUpstreamJSON `json:"-"`
}

// relayGetResponseConfigUpstreamsUpstreamJSON contains the JSON metadata for the
// struct [RelayGetResponseConfigUpstreamsUpstream]
type relayGetResponseConfigUpstreamsUpstreamJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayGetResponseConfigUpstreamsUpstream) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayGetResponseConfigUpstreamsUpstreamJSON) RawJSON() string {
	return r.raw
}

// "connected" when active, omitted otherwise.
type RelayGetResponseStatus string

const (
	RelayGetResponseStatusConnected RelayGetResponseStatus = "connected"
)

func (r RelayGetResponseStatus) IsKnown() bool {
	switch r {
	case RelayGetResponseStatusConnected:
		return true
	}
	return false
}

type RelayNewParams struct {
	// Cloudflare account identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Human-readable name for the relay.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r RelayNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RelayNewResponseEnvelope struct {
	Errors   []RelayNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RelayNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Success  bool                               `json:"success" api:"required"`
	// Relay with its auto-created default token pair (one full-access [publish,
	// subscribe] and one [subscribe]-only), each with its one-time secret, wrapped in
	// the issuers envelope.
	Result RelayNewResponse             `json:"result"`
	JSON   relayNewResponseEnvelopeJSON `json:"-"`
}

// relayNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RelayNewResponseEnvelope]
type relayNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RelayNewResponseEnvelopeErrors struct {
	Code    int64                              `json:"code"`
	Message string                             `json:"message"`
	JSON    relayNewResponseEnvelopeErrorsJSON `json:"-"`
}

// relayNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RelayNewResponseEnvelopeErrors]
type relayNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RelayNewResponseEnvelopeMessages struct {
	Code    int64                                `json:"code"`
	Message string                               `json:"message"`
	JSON    relayNewResponseEnvelopeMessagesJSON `json:"-"`
}

// relayNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RelayNewResponseEnvelopeMessages]
type relayNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RelayUpdateParams struct {
	// Cloudflare account identifier.
	AccountID param.Field[string]                  `path:"account_id" api:"required"`
	Config    param.Field[RelayUpdateParamsConfig] `json:"config"`
	Name      param.Field[string]                  `json:"name"`
}

func (r RelayUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RelayUpdateParamsConfig struct {
	// Upstreams are external MOQT server publishers that a relay falls back to when it
	// has no local publisher for a requested namespace/track.
	Upstreams param.Field[RelayUpdateParamsConfigUpstreams] `json:"upstreams"`
}

func (r RelayUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Upstreams are external MOQT server publishers that a relay falls back to when it
// has no local publisher for a requested namespace/track.
type RelayUpdateParamsConfigUpstreams struct {
	Enabled param.Field[bool] `json:"enabled"`
	// Ordered list of upstream MOQT server publishers. Each entry is an object (not a
	// bare string) so per-upstream configuration can be added in the future without
	// another breaking change.
	Upstreams param.Field[[]RelayUpdateParamsConfigUpstreamsUpstream] `json:"upstreams"`
}

func (r RelayUpdateParamsConfigUpstreams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A single upstream MOQT server publisher.
type RelayUpdateParamsConfigUpstreamsUpstream struct {
	// Upstream MOQT server publisher URL. Must be an absolute URL with a host and a
	// scheme the relay can dial: moqt:// (raw QUIC) or https:// (WebTransport).
	// Validated on update (PUT); rejected with 21013.
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
}

func (r RelayUpdateParamsConfigUpstreamsUpstream) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RelayUpdateResponseEnvelope struct {
	Errors   []RelayUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RelayUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Success  bool                                  `json:"success" api:"required"`
	// Full relay details (no tokens).
	Result RelayUpdateResponse             `json:"result"`
	JSON   relayUpdateResponseEnvelopeJSON `json:"-"`
}

// relayUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RelayUpdateResponseEnvelope]
type relayUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RelayUpdateResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code"`
	Message string                                `json:"message"`
	JSON    relayUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// relayUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RelayUpdateResponseEnvelopeErrors]
type relayUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RelayUpdateResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code"`
	Message string                                  `json:"message"`
	JSON    relayUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// relayUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RelayUpdateResponseEnvelopeMessages]
type relayUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RelayListParams struct {
	// Cloudflare account identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Sort order by `created`. When true, results are returned oldest-first
	// (ascending); otherwise newest-first (descending, the default).
	Asc param.Field[bool] `query:"asc"`
	// Cursor for pagination. Returns relays created strictly after this RFC 3339
	// timestamp (typically the `created` value of the last item on the current page,
	// to fetch the next page).
	CreatedAfter param.Field[time.Time] `query:"created_after" format:"date-time"`
	// Cursor for pagination. Returns relays created strictly before this RFC 3339
	// timestamp (typically the `created` value of the first item on the current page,
	// to fetch the previous page).
	CreatedBefore param.Field[time.Time] `query:"created_before" format:"date-time"`
	// Maximum number of relays to return per page. Values above the maximum are
	// clamped to it rather than rejected.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [RelayListParams]'s query parameters as `url.Values`.
func (r RelayListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RelayDeleteParams struct {
	// Cloudflare account identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RelayDeleteResponseEnvelope struct {
	Errors   []RelayDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RelayDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Success  bool                                  `json:"success" api:"required"`
	Result   RelayDeleteResponse                   `json:"result" api:"nullable"`
	JSON     relayDeleteResponseEnvelopeJSON       `json:"-"`
}

// relayDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RelayDeleteResponseEnvelope]
type relayDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RelayDeleteResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code"`
	Message string                                `json:"message"`
	JSON    relayDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// relayDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RelayDeleteResponseEnvelopeErrors]
type relayDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RelayDeleteResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code"`
	Message string                                  `json:"message"`
	JSON    relayDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// relayDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RelayDeleteResponseEnvelopeMessages]
type relayDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RelayGetParams struct {
	// Cloudflare account identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RelayGetResponseEnvelope struct {
	Errors   []RelayGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RelayGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Success  bool                               `json:"success" api:"required"`
	// Full relay details (no tokens).
	Result RelayGetResponse             `json:"result"`
	JSON   relayGetResponseEnvelopeJSON `json:"-"`
}

// relayGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RelayGetResponseEnvelope]
type relayGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RelayGetResponseEnvelopeErrors struct {
	Code    int64                              `json:"code"`
	Message string                             `json:"message"`
	JSON    relayGetResponseEnvelopeErrorsJSON `json:"-"`
}

// relayGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RelayGetResponseEnvelopeErrors]
type relayGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RelayGetResponseEnvelopeMessages struct {
	Code    int64                                `json:"code"`
	Message string                               `json:"message"`
	JSON    relayGetResponseEnvelopeMessagesJSON `json:"-"`
}

// relayGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RelayGetResponseEnvelopeMessages]
type relayGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
