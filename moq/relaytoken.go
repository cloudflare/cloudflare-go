// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moq

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// RelayTokenService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRelayTokenService] method instead.
type RelayTokenService struct {
	Options []option.RequestOption
}

// NewRelayTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRelayTokenService(opts ...option.RequestOption) (r *RelayTokenService) {
	r = &RelayTokenService{}
	r.Options = opts
	return
}

// Mints a new relay-scoped token and adds it to the relay's accepted-auth
// registry. The token value (secret) is shown once in the response. A relay may
// hold up to 10 tokens; creating an 11th is rejected.
func (r *RelayTokenService) New(ctx context.Context, relayID string, params RelayTokenNewParams, opts ...option.RequestOption) (res *RelayTokenNewResponse, err error) {
	var env RelayTokenNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if relayID == "" {
		err = errors.New("missing required relay_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/moq/relays/%s/tokens", params.AccountID, relayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns metadata for every token in the relay's registry. Secrets are never
// returned. The dashboard derives an `expired` flag by comparing each token's
// `expires` to the current time.
func (r *RelayTokenService) List(ctx context.Context, relayID string, query RelayTokenListParams, opts ...option.RequestOption) (res *RelayTokenListResponse, err error) {
	var env RelayTokenListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if relayID == "" {
		err = errors.New("missing required relay_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/moq/relays/%s/tokens", query.AccountID, relayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Revokes a token by removing it from the relay's registry. crique rejects the
// token within the cache TTL. Idempotent — revoking an unknown token succeeds.
func (r *RelayTokenService) Delete(ctx context.Context, relayID string, jti string, body RelayTokenDeleteParams, opts ...option.RequestOption) (res *RelayTokenDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if relayID == "" {
		err = errors.New("missing required relay_id parameter")
		return nil, err
	}
	if jti == "" {
		err = errors.New("missing required jti parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/moq/relays/%s/tokens/%s", body.AccountID, relayID, jti)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A relay's token collection, keyed on issuer `type` (a discriminated union). V1
// ships exactly one arm (`cloudflare_jwt`). Clients iterate `issuers`, switch on
// `type`, and ignore unknown types — that contract is what makes adding or
// removing an arm non-breaking.
type RelayTokenNewResponse struct {
	Issuers []RelayTokenNewResponseIssuer `json:"issuers" api:"required"`
	JSON    relayTokenNewResponseJSON     `json:"-"`
}

// relayTokenNewResponseJSON contains the JSON metadata for the struct
// [RelayTokenNewResponse]
type relayTokenNewResponseJSON struct {
	Issuers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenNewResponseJSON) RawJSON() string {
	return r.raw
}

// One arm of the discriminated-union token collection.
type RelayTokenNewResponseIssuer struct {
	// Always present ([] when empty).
	CloudflareTokens []RelayTokenNewResponseIssuersCloudflareToken `json:"cloudflare_tokens" api:"required"`
	Issuer           RelayTokenNewResponseIssuersIssuer            `json:"issuer" api:"required"`
	Type             RelayTokenNewResponseIssuersType              `json:"type" api:"required"`
	JSON             relayTokenNewResponseIssuerJSON               `json:"-"`
}

// relayTokenNewResponseIssuerJSON contains the JSON metadata for the struct
// [RelayTokenNewResponseIssuer]
type relayTokenNewResponseIssuerJSON struct {
	CloudflareTokens apijson.Field
	Issuer           apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RelayTokenNewResponseIssuer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenNewResponseIssuerJSON) RawJSON() string {
	return r.raw
}

type RelayTokenNewResponseIssuersCloudflareToken struct {
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Mandatory; no more than 1 year after `created`.
	Expires time.Time `json:"expires" api:"required" format:"date-time"`
	// Token identity and registry key (32 hex chars).
	Jti string `json:"jti" api:"required"`
	// Signed allowlist of what the token may do. V1 coarse roles; the array form
	// extends to fine-grained MoQT message names later without a breaking change.
	Operations []RelayTokenNewResponseIssuersCloudflareTokensOperation `json:"operations" api:"required"`
	// Optional, customer-set.
	Label string `json:"label"`
	// The signed JWT. Present ONLY in create / auto-create responses (shown once);
	// never returned by list, never stored.
	Secret string                                          `json:"secret"`
	JSON   relayTokenNewResponseIssuersCloudflareTokenJSON `json:"-"`
}

// relayTokenNewResponseIssuersCloudflareTokenJSON contains the JSON metadata for
// the struct [RelayTokenNewResponseIssuersCloudflareToken]
type relayTokenNewResponseIssuersCloudflareTokenJSON struct {
	Created     apijson.Field
	Expires     apijson.Field
	Jti         apijson.Field
	Operations  apijson.Field
	Label       apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenNewResponseIssuersCloudflareToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenNewResponseIssuersCloudflareTokenJSON) RawJSON() string {
	return r.raw
}

type RelayTokenNewResponseIssuersCloudflareTokensOperation string

const (
	RelayTokenNewResponseIssuersCloudflareTokensOperationPublish   RelayTokenNewResponseIssuersCloudflareTokensOperation = "publish"
	RelayTokenNewResponseIssuersCloudflareTokensOperationSubscribe RelayTokenNewResponseIssuersCloudflareTokensOperation = "subscribe"
)

func (r RelayTokenNewResponseIssuersCloudflareTokensOperation) IsKnown() bool {
	switch r {
	case RelayTokenNewResponseIssuersCloudflareTokensOperationPublish, RelayTokenNewResponseIssuersCloudflareTokensOperationSubscribe:
		return true
	}
	return false
}

type RelayTokenNewResponseIssuersIssuer string

const (
	RelayTokenNewResponseIssuersIssuerCloudflare RelayTokenNewResponseIssuersIssuer = "cloudflare"
)

func (r RelayTokenNewResponseIssuersIssuer) IsKnown() bool {
	switch r {
	case RelayTokenNewResponseIssuersIssuerCloudflare:
		return true
	}
	return false
}

type RelayTokenNewResponseIssuersType string

const (
	RelayTokenNewResponseIssuersTypeCloudflareJWT RelayTokenNewResponseIssuersType = "cloudflare_jwt"
)

func (r RelayTokenNewResponseIssuersType) IsKnown() bool {
	switch r {
	case RelayTokenNewResponseIssuersTypeCloudflareJWT:
		return true
	}
	return false
}

// A relay's token collection, keyed on issuer `type` (a discriminated union). V1
// ships exactly one arm (`cloudflare_jwt`). Clients iterate `issuers`, switch on
// `type`, and ignore unknown types — that contract is what makes adding or
// removing an arm non-breaking.
type RelayTokenListResponse struct {
	Issuers []RelayTokenListResponseIssuer `json:"issuers" api:"required"`
	JSON    relayTokenListResponseJSON     `json:"-"`
}

// relayTokenListResponseJSON contains the JSON metadata for the struct
// [RelayTokenListResponse]
type relayTokenListResponseJSON struct {
	Issuers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenListResponseJSON) RawJSON() string {
	return r.raw
}

// One arm of the discriminated-union token collection.
type RelayTokenListResponseIssuer struct {
	// Always present ([] when empty).
	CloudflareTokens []RelayTokenListResponseIssuersCloudflareToken `json:"cloudflare_tokens" api:"required"`
	Issuer           RelayTokenListResponseIssuersIssuer            `json:"issuer" api:"required"`
	Type             RelayTokenListResponseIssuersType              `json:"type" api:"required"`
	JSON             relayTokenListResponseIssuerJSON               `json:"-"`
}

// relayTokenListResponseIssuerJSON contains the JSON metadata for the struct
// [RelayTokenListResponseIssuer]
type relayTokenListResponseIssuerJSON struct {
	CloudflareTokens apijson.Field
	Issuer           apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RelayTokenListResponseIssuer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenListResponseIssuerJSON) RawJSON() string {
	return r.raw
}

type RelayTokenListResponseIssuersCloudflareToken struct {
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Mandatory; no more than 1 year after `created`.
	Expires time.Time `json:"expires" api:"required" format:"date-time"`
	// Token identity and registry key (32 hex chars).
	Jti string `json:"jti" api:"required"`
	// Signed allowlist of what the token may do. V1 coarse roles; the array form
	// extends to fine-grained MoQT message names later without a breaking change.
	Operations []RelayTokenListResponseIssuersCloudflareTokensOperation `json:"operations" api:"required"`
	// Optional, customer-set.
	Label string `json:"label"`
	// The signed JWT. Present ONLY in create / auto-create responses (shown once);
	// never returned by list, never stored.
	Secret string                                           `json:"secret"`
	JSON   relayTokenListResponseIssuersCloudflareTokenJSON `json:"-"`
}

// relayTokenListResponseIssuersCloudflareTokenJSON contains the JSON metadata for
// the struct [RelayTokenListResponseIssuersCloudflareToken]
type relayTokenListResponseIssuersCloudflareTokenJSON struct {
	Created     apijson.Field
	Expires     apijson.Field
	Jti         apijson.Field
	Operations  apijson.Field
	Label       apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenListResponseIssuersCloudflareToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenListResponseIssuersCloudflareTokenJSON) RawJSON() string {
	return r.raw
}

type RelayTokenListResponseIssuersCloudflareTokensOperation string

const (
	RelayTokenListResponseIssuersCloudflareTokensOperationPublish   RelayTokenListResponseIssuersCloudflareTokensOperation = "publish"
	RelayTokenListResponseIssuersCloudflareTokensOperationSubscribe RelayTokenListResponseIssuersCloudflareTokensOperation = "subscribe"
)

func (r RelayTokenListResponseIssuersCloudflareTokensOperation) IsKnown() bool {
	switch r {
	case RelayTokenListResponseIssuersCloudflareTokensOperationPublish, RelayTokenListResponseIssuersCloudflareTokensOperationSubscribe:
		return true
	}
	return false
}

type RelayTokenListResponseIssuersIssuer string

const (
	RelayTokenListResponseIssuersIssuerCloudflare RelayTokenListResponseIssuersIssuer = "cloudflare"
)

func (r RelayTokenListResponseIssuersIssuer) IsKnown() bool {
	switch r {
	case RelayTokenListResponseIssuersIssuerCloudflare:
		return true
	}
	return false
}

type RelayTokenListResponseIssuersType string

const (
	RelayTokenListResponseIssuersTypeCloudflareJWT RelayTokenListResponseIssuersType = "cloudflare_jwt"
)

func (r RelayTokenListResponseIssuersType) IsKnown() bool {
	switch r {
	case RelayTokenListResponseIssuersTypeCloudflareJWT:
		return true
	}
	return false
}

type RelayTokenDeleteResponse struct {
	Errors   []RelayTokenDeleteResponseError   `json:"errors" api:"required"`
	Messages []RelayTokenDeleteResponseMessage `json:"messages" api:"required"`
	Success  bool                              `json:"success" api:"required"`
	JSON     relayTokenDeleteResponseJSON      `json:"-"`
}

// relayTokenDeleteResponseJSON contains the JSON metadata for the struct
// [RelayTokenDeleteResponse]
type relayTokenDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RelayTokenDeleteResponseError struct {
	Code    int64                             `json:"code"`
	Message string                            `json:"message"`
	JSON    relayTokenDeleteResponseErrorJSON `json:"-"`
}

// relayTokenDeleteResponseErrorJSON contains the JSON metadata for the struct
// [RelayTokenDeleteResponseError]
type relayTokenDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type RelayTokenDeleteResponseMessage struct {
	Code    int64                               `json:"code"`
	Message string                              `json:"message"`
	JSON    relayTokenDeleteResponseMessageJSON `json:"-"`
}

// relayTokenDeleteResponseMessageJSON contains the JSON metadata for the struct
// [RelayTokenDeleteResponseMessage]
type relayTokenDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type RelayTokenNewParams struct {
	// Cloudflare account identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Non-empty subset of the V1 roles the token is allowed to perform. Signed into
	// the token.
	Operations param.Field[[]RelayTokenNewParamsOperation] `json:"operations" api:"required"`
	// Optional expiry (RFC 3339). Defaults to 1 year from creation; rejected if more
	// than 1 year in the future.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// Optional, customer-set label.
	Label param.Field[string] `json:"label"`
}

func (r RelayTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RelayTokenNewParamsOperation string

const (
	RelayTokenNewParamsOperationPublish   RelayTokenNewParamsOperation = "publish"
	RelayTokenNewParamsOperationSubscribe RelayTokenNewParamsOperation = "subscribe"
)

func (r RelayTokenNewParamsOperation) IsKnown() bool {
	switch r {
	case RelayTokenNewParamsOperationPublish, RelayTokenNewParamsOperationSubscribe:
		return true
	}
	return false
}

type RelayTokenNewResponseEnvelope struct {
	Errors   []RelayTokenNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RelayTokenNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Success  bool                                    `json:"success" api:"required"`
	// A relay's token collection, keyed on issuer `type` (a discriminated union). V1
	// ships exactly one arm (`cloudflare_jwt`). Clients iterate `issuers`, switch on
	// `type`, and ignore unknown types — that contract is what makes adding or
	// removing an arm non-breaking.
	Result RelayTokenNewResponse             `json:"result"`
	JSON   relayTokenNewResponseEnvelopeJSON `json:"-"`
}

// relayTokenNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RelayTokenNewResponseEnvelope]
type relayTokenNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RelayTokenNewResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code"`
	Message string                                  `json:"message"`
	JSON    relayTokenNewResponseEnvelopeErrorsJSON `json:"-"`
}

// relayTokenNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RelayTokenNewResponseEnvelopeErrors]
type relayTokenNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RelayTokenNewResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code"`
	Message string                                    `json:"message"`
	JSON    relayTokenNewResponseEnvelopeMessagesJSON `json:"-"`
}

// relayTokenNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RelayTokenNewResponseEnvelopeMessages]
type relayTokenNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RelayTokenListParams struct {
	// Cloudflare account identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RelayTokenListResponseEnvelope struct {
	Errors   []RelayTokenListResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RelayTokenListResponseEnvelopeMessages `json:"messages" api:"required"`
	Success  bool                                     `json:"success" api:"required"`
	// A relay's token collection, keyed on issuer `type` (a discriminated union). V1
	// ships exactly one arm (`cloudflare_jwt`). Clients iterate `issuers`, switch on
	// `type`, and ignore unknown types — that contract is what makes adding or
	// removing an arm non-breaking.
	Result RelayTokenListResponse             `json:"result"`
	JSON   relayTokenListResponseEnvelopeJSON `json:"-"`
}

// relayTokenListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RelayTokenListResponseEnvelope]
type relayTokenListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RelayTokenListResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code"`
	Message string                                   `json:"message"`
	JSON    relayTokenListResponseEnvelopeErrorsJSON `json:"-"`
}

// relayTokenListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RelayTokenListResponseEnvelopeErrors]
type relayTokenListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RelayTokenListResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code"`
	Message string                                     `json:"message"`
	JSON    relayTokenListResponseEnvelopeMessagesJSON `json:"-"`
}

// relayTokenListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RelayTokenListResponseEnvelopeMessages]
type relayTokenListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RelayTokenDeleteParams struct {
	// Cloudflare account identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
