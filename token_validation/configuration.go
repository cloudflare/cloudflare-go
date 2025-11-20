// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/api_gateway"
	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/tidwall/gjson"
)

// ConfigurationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigurationService] method instead.
type ConfigurationService struct {
	Options     []option.RequestOption
	Credentials *ConfigurationCredentialService
}

// NewConfigurationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigurationService(opts ...option.RequestOption) (r *ConfigurationService) {
	r = &ConfigurationService{}
	r.Options = opts
	r.Credentials = NewConfigurationCredentialService(opts...)
	return
}

// Create a new Token Validation configuration
func (r *ConfigurationService) New(ctx context.Context, params ConfigurationNewParams, opts ...option.RequestOption) (res *TokenConfig, err error) {
	var env ConfigurationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/token_validation/config", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all token validation configurations for this zone
func (r *ConfigurationService) List(ctx context.Context, params ConfigurationListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[TokenConfig], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/token_validation/config", params.ZoneID)
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

// Lists all token validation configurations for this zone
func (r *ConfigurationService) ListAutoPaging(ctx context.Context, params ConfigurationListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[TokenConfig] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete Token Configuration
func (r *ConfigurationService) Delete(ctx context.Context, configID string, body ConfigurationDeleteParams, opts ...option.RequestOption) (res *ConfigurationDeleteResponse, err error) {
	var env ConfigurationDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/token_validation/config/%s", body.ZoneID, configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit fields of an existing Token Configuration
func (r *ConfigurationService) Edit(ctx context.Context, configID string, params ConfigurationEditParams, opts ...option.RequestOption) (res *ConfigurationEditResponse, err error) {
	var env ConfigurationEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/token_validation/config/%s", params.ZoneID, configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a single Token Configuration
func (r *ConfigurationService) Get(ctx context.Context, configID string, query ConfigurationGetParams, opts ...option.RequestOption) (res *TokenConfig, err error) {
	var env ConfigurationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/token_validation/config/%s", query.ZoneID, configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TokenConfig struct {
	// UUID.
	ID           string                 `json:"id,required"`
	CreatedAt    time.Time              `json:"created_at,required" format:"date-time"`
	Credentials  TokenConfigCredentials `json:"credentials,required"`
	Description  string                 `json:"description,required"`
	LastUpdated  time.Time              `json:"last_updated,required" format:"date-time"`
	Title        string                 `json:"title,required"`
	TokenSources []string               `json:"token_sources,required"`
	TokenType    TokenConfigTokenType   `json:"token_type,required"`
	JSON         tokenConfigJSON        `json:"-"`
}

// tokenConfigJSON contains the JSON metadata for the struct [TokenConfig]
type tokenConfigJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Credentials  apijson.Field
	Description  apijson.Field
	LastUpdated  apijson.Field
	Title        apijson.Field
	TokenSources apijson.Field
	TokenType    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TokenConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenConfigJSON) RawJSON() string {
	return r.raw
}

type TokenConfigCredentials struct {
	Keys []TokenConfigCredentialsKey `json:"keys,required"`
	JSON tokenConfigCredentialsJSON  `json:"-"`
}

// tokenConfigCredentialsJSON contains the JSON metadata for the struct
// [TokenConfigCredentials]
type tokenConfigCredentialsJSON struct {
	Keys        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenConfigCredentials) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenConfigCredentialsJSON) RawJSON() string {
	return r.raw
}

// JSON representation of a JWKS key.
type TokenConfigCredentialsKey struct {
	// Algorithm
	Alg TokenConfigCredentialsKeysAlg `json:"alg,required"`
	// Key ID
	Kid string `json:"kid,required"`
	// Key Type
	Kty TokenConfigCredentialsKeysKty `json:"kty,required"`
	// Curve
	Crv TokenConfigCredentialsKeysCrv `json:"crv"`
	// RSA exponent
	E string `json:"e"`
	// RSA modulus
	N string `json:"n"`
	// X EC coordinate
	X string `json:"x"`
	// Y EC coordinate
	Y     string                        `json:"y"`
	JSON  tokenConfigCredentialsKeyJSON `json:"-"`
	union TokenConfigCredentialsKeysUnion
}

// tokenConfigCredentialsKeyJSON contains the JSON metadata for the struct
// [TokenConfigCredentialsKey]
type tokenConfigCredentialsKeyJSON struct {
	Alg         apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	Crv         apijson.Field
	E           apijson.Field
	N           apijson.Field
	X           apijson.Field
	Y           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r tokenConfigCredentialsKeyJSON) RawJSON() string {
	return r.raw
}

func (r *TokenConfigCredentialsKey) UnmarshalJSON(data []byte) (err error) {
	*r = TokenConfigCredentialsKey{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TokenConfigCredentialsKeysUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSA],
// [TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256],
// [TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384].
func (r TokenConfigCredentialsKey) AsUnion() TokenConfigCredentialsKeysUnion {
	return r.union
}

// JSON representation of a JWKS key.
//
// Union satisfied by [TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSA],
// [TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256] or
// [TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384].
type TokenConfigCredentialsKeysUnion interface {
	implementsTokenConfigCredentialsKey()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TokenConfigCredentialsKeysUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSA{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384{}),
		},
	)
}

// JSON representation of an RSA key.
type TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSA struct {
	// Algorithm
	Alg TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg `json:"alg,required"`
	// RSA exponent
	E string `json:"e,required"`
	// Key ID
	Kid string `json:"kid,required"`
	// Key Type
	Kty TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAKty `json:"kty,required"`
	// RSA modulus
	N    string                                                      `json:"n,required"`
	JSON tokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAJSON `json:"-"`
}

// tokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAJSON contains the JSON
// metadata for the struct
// [TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSA]
type tokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAJSON struct {
	Alg         apijson.Field
	E           apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	N           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAJSON) RawJSON() string {
	return r.raw
}

func (r TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSA) implementsTokenConfigCredentialsKey() {
}

// Algorithm
type TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg string

const (
	TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgRs256 TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS256"
	TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgRs384 TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS384"
	TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgRs512 TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS512"
	TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgPs256 TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS256"
	TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgPs384 TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS384"
	TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgPs512 TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS512"
)

func (r TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg) IsKnown() bool {
	switch r {
	case TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgRs256, TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgRs384, TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgRs512, TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgPs256, TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgPs384, TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgPs512:
		return true
	}
	return false
}

// Key Type
type TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAKty string

const (
	TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAKtyRSA TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAKty = "RSA"
)

func (r TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAKty) IsKnown() bool {
	switch r {
	case TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyRSAKtyRSA:
		return true
	}
	return false
}

// JSON representation of an ES256 key
type TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256 struct {
	// Algorithm
	Alg TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Alg `json:"alg,required"`
	// Curve
	Crv TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Crv `json:"crv,required"`
	// Key ID
	Kid string `json:"kid,required"`
	// Key Type
	Kty TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Kty `json:"kty,required"`
	// X EC coordinate
	X string `json:"x,required"`
	// Y EC coordinate
	Y    string                                                          `json:"y,required"`
	JSON tokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256JSON `json:"-"`
}

// tokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256JSON contains the
// JSON metadata for the struct
// [TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256]
type tokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256JSON struct {
	Alg         apijson.Field
	Crv         apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	X           apijson.Field
	Y           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256JSON) RawJSON() string {
	return r.raw
}

func (r TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256) implementsTokenConfigCredentialsKey() {
}

// Algorithm
type TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Alg string

const (
	TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256AlgEs256 TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Alg = "ES256"
)

func (r TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Alg) IsKnown() bool {
	switch r {
	case TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256AlgEs256:
		return true
	}
	return false
}

// Curve
type TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Crv string

const (
	TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256CrvP256 TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Crv = "P-256"
)

func (r TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Crv) IsKnown() bool {
	switch r {
	case TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256CrvP256:
		return true
	}
	return false
}

// Key Type
type TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Kty string

const (
	TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256KtyEc TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Kty = "EC"
)

func (r TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Kty) IsKnown() bool {
	switch r {
	case TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256KtyEc:
		return true
	}
	return false
}

// JSON representation of an ES384 key
type TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384 struct {
	// Algorithm
	Alg TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Alg `json:"alg,required"`
	// Curve
	Crv TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Crv `json:"crv,required"`
	// Key ID
	Kid string `json:"kid,required"`
	// Key Type
	Kty TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Kty `json:"kty,required"`
	// X EC coordinate
	X string `json:"x,required"`
	// Y EC coordinate
	Y    string                                                          `json:"y,required"`
	JSON tokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384JSON `json:"-"`
}

// tokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384JSON contains the
// JSON metadata for the struct
// [TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384]
type tokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384JSON struct {
	Alg         apijson.Field
	Crv         apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	X           apijson.Field
	Y           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384JSON) RawJSON() string {
	return r.raw
}

func (r TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384) implementsTokenConfigCredentialsKey() {
}

// Algorithm
type TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Alg string

const (
	TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384AlgEs384 TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Alg = "ES384"
)

func (r TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Alg) IsKnown() bool {
	switch r {
	case TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384AlgEs384:
		return true
	}
	return false
}

// Curve
type TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Crv string

const (
	TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384CrvP384 TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Crv = "P-384"
)

func (r TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Crv) IsKnown() bool {
	switch r {
	case TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384CrvP384:
		return true
	}
	return false
}

// Key Type
type TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Kty string

const (
	TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384KtyEc TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Kty = "EC"
)

func (r TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Kty) IsKnown() bool {
	switch r {
	case TokenConfigCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384KtyEc:
		return true
	}
	return false
}

// Algorithm
type TokenConfigCredentialsKeysAlg string

const (
	TokenConfigCredentialsKeysAlgRs256 TokenConfigCredentialsKeysAlg = "RS256"
	TokenConfigCredentialsKeysAlgRs384 TokenConfigCredentialsKeysAlg = "RS384"
	TokenConfigCredentialsKeysAlgRs512 TokenConfigCredentialsKeysAlg = "RS512"
	TokenConfigCredentialsKeysAlgPs256 TokenConfigCredentialsKeysAlg = "PS256"
	TokenConfigCredentialsKeysAlgPs384 TokenConfigCredentialsKeysAlg = "PS384"
	TokenConfigCredentialsKeysAlgPs512 TokenConfigCredentialsKeysAlg = "PS512"
	TokenConfigCredentialsKeysAlgEs256 TokenConfigCredentialsKeysAlg = "ES256"
	TokenConfigCredentialsKeysAlgEs384 TokenConfigCredentialsKeysAlg = "ES384"
)

func (r TokenConfigCredentialsKeysAlg) IsKnown() bool {
	switch r {
	case TokenConfigCredentialsKeysAlgRs256, TokenConfigCredentialsKeysAlgRs384, TokenConfigCredentialsKeysAlgRs512, TokenConfigCredentialsKeysAlgPs256, TokenConfigCredentialsKeysAlgPs384, TokenConfigCredentialsKeysAlgPs512, TokenConfigCredentialsKeysAlgEs256, TokenConfigCredentialsKeysAlgEs384:
		return true
	}
	return false
}

// Key Type
type TokenConfigCredentialsKeysKty string

const (
	TokenConfigCredentialsKeysKtyRSA TokenConfigCredentialsKeysKty = "RSA"
	TokenConfigCredentialsKeysKtyEc  TokenConfigCredentialsKeysKty = "EC"
)

func (r TokenConfigCredentialsKeysKty) IsKnown() bool {
	switch r {
	case TokenConfigCredentialsKeysKtyRSA, TokenConfigCredentialsKeysKtyEc:
		return true
	}
	return false
}

// Curve
type TokenConfigCredentialsKeysCrv string

const (
	TokenConfigCredentialsKeysCrvP256 TokenConfigCredentialsKeysCrv = "P-256"
	TokenConfigCredentialsKeysCrvP384 TokenConfigCredentialsKeysCrv = "P-384"
)

func (r TokenConfigCredentialsKeysCrv) IsKnown() bool {
	switch r {
	case TokenConfigCredentialsKeysCrvP256, TokenConfigCredentialsKeysCrvP384:
		return true
	}
	return false
}

type TokenConfigTokenType string

const (
	TokenConfigTokenTypeJWT TokenConfigTokenType = "JWT"
)

func (r TokenConfigTokenType) IsKnown() bool {
	switch r {
	case TokenConfigTokenTypeJWT:
		return true
	}
	return false
}

type ConfigurationDeleteResponse struct {
	// UUID.
	ID   string                          `json:"id"`
	JSON configurationDeleteResponseJSON `json:"-"`
}

// configurationDeleteResponseJSON contains the JSON metadata for the struct
// [ConfigurationDeleteResponse]
type configurationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigurationEditResponse struct {
	Description  string                        `json:"description"`
	Title        string                        `json:"title"`
	TokenSources []string                      `json:"token_sources"`
	JSON         configurationEditResponseJSON `json:"-"`
}

// configurationEditResponseJSON contains the JSON metadata for the struct
// [ConfigurationEditResponse]
type configurationEditResponseJSON struct {
	Description  apijson.Field
	Title        apijson.Field
	TokenSources apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConfigurationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationEditResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigurationNewParams struct {
	// Identifier.
	ZoneID       param.Field[string]                            `path:"zone_id,required"`
	Credentials  param.Field[ConfigurationNewParamsCredentials] `json:"credentials,required"`
	Description  param.Field[string]                            `json:"description,required"`
	Title        param.Field[string]                            `json:"title,required"`
	TokenSources param.Field[[]string]                          `json:"token_sources,required"`
	TokenType    param.Field[ConfigurationNewParamsTokenType]   `json:"token_type,required"`
}

func (r ConfigurationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigurationNewParamsCredentials struct {
	Keys param.Field[[]ConfigurationNewParamsCredentialsKeyUnion] `json:"keys,required"`
}

func (r ConfigurationNewParamsCredentials) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// JSON representation of a JWKS key.
type ConfigurationNewParamsCredentialsKey struct {
	// Algorithm
	Alg param.Field[ConfigurationNewParamsCredentialsKeysAlg] `json:"alg,required"`
	// Key ID
	Kid param.Field[string] `json:"kid,required"`
	// Key Type
	Kty param.Field[ConfigurationNewParamsCredentialsKeysKty] `json:"kty,required"`
	// Curve
	Crv param.Field[ConfigurationNewParamsCredentialsKeysCrv] `json:"crv"`
	// RSA exponent
	E param.Field[string] `json:"e"`
	// RSA modulus
	N param.Field[string] `json:"n"`
	// X EC coordinate
	X param.Field[string] `json:"x"`
	// Y EC coordinate
	Y param.Field[string] `json:"y"`
}

func (r ConfigurationNewParamsCredentialsKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationNewParamsCredentialsKey) implementsConfigurationNewParamsCredentialsKeyUnion() {}

// JSON representation of a JWKS key.
//
// Satisfied by
// [token_validation.ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSA],
// [token_validation.ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256],
// [token_validation.ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384],
// [ConfigurationNewParamsCredentialsKey].
type ConfigurationNewParamsCredentialsKeyUnion interface {
	implementsConfigurationNewParamsCredentialsKeyUnion()
}

// JSON representation of an RSA key.
type ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSA struct {
	// Algorithm
	Alg param.Field[ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg] `json:"alg,required"`
	// RSA exponent
	E param.Field[string] `json:"e,required"`
	// Key ID
	Kid param.Field[string] `json:"kid,required"`
	// Key Type
	Kty param.Field[ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAKty] `json:"kty,required"`
	// RSA modulus
	N param.Field[string] `json:"n,required"`
}

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSA) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSA) implementsConfigurationNewParamsCredentialsKeyUnion() {
}

// Algorithm
type ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg string

const (
	ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgRs256 ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS256"
	ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgRs384 ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS384"
	ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgRs512 ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS512"
	ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgPs256 ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS256"
	ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgPs384 ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS384"
	ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgPs512 ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS512"
)

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlg) IsKnown() bool {
	switch r {
	case ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgRs256, ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgRs384, ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgRs512, ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgPs256, ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgPs384, ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAAlgPs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAKty string

const (
	ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAKtyRSA ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAKty = "RSA"
)

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAKty) IsKnown() bool {
	switch r {
	case ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyRSAKtyRSA:
		return true
	}
	return false
}

// JSON representation of an ES256 key
type ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256 struct {
	// Algorithm
	Alg param.Field[ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Alg] `json:"alg,required"`
	// Curve
	Crv param.Field[ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Crv] `json:"crv,required"`
	// Key ID
	Kid param.Field[string] `json:"kid,required"`
	// Key Type
	Kty param.Field[ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Kty] `json:"kty,required"`
	// X EC coordinate
	X param.Field[string] `json:"x,required"`
	// Y EC coordinate
	Y param.Field[string] `json:"y,required"`
}

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256) implementsConfigurationNewParamsCredentialsKeyUnion() {
}

// Algorithm
type ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Alg string

const (
	ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256AlgEs256 ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Alg = "ES256"
)

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Alg) IsKnown() bool {
	switch r {
	case ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256AlgEs256:
		return true
	}
	return false
}

// Curve
type ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Crv string

const (
	ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256CrvP256 ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Crv = "P-256"
)

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Crv) IsKnown() bool {
	switch r {
	case ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256CrvP256:
		return true
	}
	return false
}

// Key Type
type ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Kty string

const (
	ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256KtyEc ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Kty = "EC"
)

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256Kty) IsKnown() bool {
	switch r {
	case ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs256KtyEc:
		return true
	}
	return false
}

// JSON representation of an ES384 key
type ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384 struct {
	// Algorithm
	Alg param.Field[ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Alg] `json:"alg,required"`
	// Curve
	Crv param.Field[ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Crv] `json:"crv,required"`
	// Key ID
	Kid param.Field[string] `json:"kid,required"`
	// Key Type
	Kty param.Field[ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Kty] `json:"kty,required"`
	// X EC coordinate
	X param.Field[string] `json:"x,required"`
	// Y EC coordinate
	Y param.Field[string] `json:"y,required"`
}

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384) implementsConfigurationNewParamsCredentialsKeyUnion() {
}

// Algorithm
type ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Alg string

const (
	ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384AlgEs384 ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Alg = "ES384"
)

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Alg) IsKnown() bool {
	switch r {
	case ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384AlgEs384:
		return true
	}
	return false
}

// Curve
type ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Crv string

const (
	ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384CrvP384 ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Crv = "P-384"
)

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Crv) IsKnown() bool {
	switch r {
	case ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384CrvP384:
		return true
	}
	return false
}

// Key Type
type ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Kty string

const (
	ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384KtyEc ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Kty = "EC"
)

func (r ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384Kty) IsKnown() bool {
	switch r {
	case ConfigurationNewParamsCredentialsKeysAPIShieldCredentialsJWTKeyEcEs384KtyEc:
		return true
	}
	return false
}

// Algorithm
type ConfigurationNewParamsCredentialsKeysAlg string

const (
	ConfigurationNewParamsCredentialsKeysAlgRs256 ConfigurationNewParamsCredentialsKeysAlg = "RS256"
	ConfigurationNewParamsCredentialsKeysAlgRs384 ConfigurationNewParamsCredentialsKeysAlg = "RS384"
	ConfigurationNewParamsCredentialsKeysAlgRs512 ConfigurationNewParamsCredentialsKeysAlg = "RS512"
	ConfigurationNewParamsCredentialsKeysAlgPs256 ConfigurationNewParamsCredentialsKeysAlg = "PS256"
	ConfigurationNewParamsCredentialsKeysAlgPs384 ConfigurationNewParamsCredentialsKeysAlg = "PS384"
	ConfigurationNewParamsCredentialsKeysAlgPs512 ConfigurationNewParamsCredentialsKeysAlg = "PS512"
	ConfigurationNewParamsCredentialsKeysAlgEs256 ConfigurationNewParamsCredentialsKeysAlg = "ES256"
	ConfigurationNewParamsCredentialsKeysAlgEs384 ConfigurationNewParamsCredentialsKeysAlg = "ES384"
)

func (r ConfigurationNewParamsCredentialsKeysAlg) IsKnown() bool {
	switch r {
	case ConfigurationNewParamsCredentialsKeysAlgRs256, ConfigurationNewParamsCredentialsKeysAlgRs384, ConfigurationNewParamsCredentialsKeysAlgRs512, ConfigurationNewParamsCredentialsKeysAlgPs256, ConfigurationNewParamsCredentialsKeysAlgPs384, ConfigurationNewParamsCredentialsKeysAlgPs512, ConfigurationNewParamsCredentialsKeysAlgEs256, ConfigurationNewParamsCredentialsKeysAlgEs384:
		return true
	}
	return false
}

// Key Type
type ConfigurationNewParamsCredentialsKeysKty string

const (
	ConfigurationNewParamsCredentialsKeysKtyRSA ConfigurationNewParamsCredentialsKeysKty = "RSA"
	ConfigurationNewParamsCredentialsKeysKtyEc  ConfigurationNewParamsCredentialsKeysKty = "EC"
)

func (r ConfigurationNewParamsCredentialsKeysKty) IsKnown() bool {
	switch r {
	case ConfigurationNewParamsCredentialsKeysKtyRSA, ConfigurationNewParamsCredentialsKeysKtyEc:
		return true
	}
	return false
}

// Curve
type ConfigurationNewParamsCredentialsKeysCrv string

const (
	ConfigurationNewParamsCredentialsKeysCrvP256 ConfigurationNewParamsCredentialsKeysCrv = "P-256"
	ConfigurationNewParamsCredentialsKeysCrvP384 ConfigurationNewParamsCredentialsKeysCrv = "P-384"
)

func (r ConfigurationNewParamsCredentialsKeysCrv) IsKnown() bool {
	switch r {
	case ConfigurationNewParamsCredentialsKeysCrvP256, ConfigurationNewParamsCredentialsKeysCrvP384:
		return true
	}
	return false
}

type ConfigurationNewParamsTokenType string

const (
	ConfigurationNewParamsTokenTypeJWT ConfigurationNewParamsTokenType = "JWT"
)

func (r ConfigurationNewParamsTokenType) IsKnown() bool {
	switch r {
	case ConfigurationNewParamsTokenTypeJWT:
		return true
	}
	return false
}

type ConfigurationNewResponseEnvelope struct {
	Errors   api_gateway.Message `json:"errors,required"`
	Messages api_gateway.Message `json:"messages,required"`
	Result   TokenConfig         `json:"result,required"`
	// Whether the API call was successful.
	Success ConfigurationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    configurationNewResponseEnvelopeJSON    `json:"-"`
}

// configurationNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigurationNewResponseEnvelope]
type configurationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConfigurationNewResponseEnvelopeSuccess bool

const (
	ConfigurationNewResponseEnvelopeSuccessTrue ConfigurationNewResponseEnvelopeSuccess = true
)

func (r ConfigurationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigurationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigurationListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ConfigurationListParams]'s query parameters as
// `url.Values`.
func (r ConfigurationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ConfigurationDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ConfigurationDeleteResponseEnvelope struct {
	Errors   api_gateway.Message         `json:"errors,required"`
	Messages api_gateway.Message         `json:"messages,required"`
	Result   ConfigurationDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success ConfigurationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    configurationDeleteResponseEnvelopeJSON    `json:"-"`
}

// configurationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ConfigurationDeleteResponseEnvelope]
type configurationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConfigurationDeleteResponseEnvelopeSuccess bool

const (
	ConfigurationDeleteResponseEnvelopeSuccessTrue ConfigurationDeleteResponseEnvelopeSuccess = true
)

func (r ConfigurationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigurationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigurationEditParams struct {
	// Identifier.
	ZoneID       param.Field[string]   `path:"zone_id,required"`
	Description  param.Field[string]   `json:"description"`
	Title        param.Field[string]   `json:"title"`
	TokenSources param.Field[[]string] `json:"token_sources"`
}

func (r ConfigurationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigurationEditResponseEnvelope struct {
	Errors   api_gateway.Message       `json:"errors,required"`
	Messages api_gateway.Message       `json:"messages,required"`
	Result   ConfigurationEditResponse `json:"result,required"`
	// Whether the API call was successful.
	Success ConfigurationEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    configurationEditResponseEnvelopeJSON    `json:"-"`
}

// configurationEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigurationEditResponseEnvelope]
type configurationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConfigurationEditResponseEnvelopeSuccess bool

const (
	ConfigurationEditResponseEnvelopeSuccessTrue ConfigurationEditResponseEnvelopeSuccess = true
)

func (r ConfigurationEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigurationEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigurationGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ConfigurationGetResponseEnvelope struct {
	Errors   api_gateway.Message `json:"errors,required"`
	Messages api_gateway.Message `json:"messages,required"`
	Result   TokenConfig         `json:"result,required"`
	// Whether the API call was successful.
	Success ConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    configurationGetResponseEnvelopeJSON    `json:"-"`
}

// configurationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigurationGetResponseEnvelope]
type configurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConfigurationGetResponseEnvelopeSuccess bool

const (
	ConfigurationGetResponseEnvelopeSuccessTrue ConfigurationGetResponseEnvelopeSuccess = true
)

func (r ConfigurationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigurationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
