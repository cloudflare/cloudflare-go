// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/api_gateway"
	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/tidwall/gjson"
)

// ConfigurationCredentialService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigurationCredentialService] method instead.
type ConfigurationCredentialService struct {
	Options []option.RequestOption
}

// NewConfigurationCredentialService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConfigurationCredentialService(opts ...option.RequestOption) (r *ConfigurationCredentialService) {
	r = &ConfigurationCredentialService{}
	r.Options = opts
	return
}

// Update Token Configuration credentials
func (r *ConfigurationCredentialService) Update(ctx context.Context, configID string, params ConfigurationCredentialUpdateParams, opts ...option.RequestOption) (res *ConfigurationCredentialUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/token_validation/config/%s/credentials", params.ZoneID, configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type ConfigurationCredentialUpdateResponse struct {
	Errors   api_gateway.Message `json:"errors,required"`
	Messages api_gateway.Message `json:"messages,required"`
	// Whether the API call was successful.
	Success ConfigurationCredentialUpdateResponseSuccess `json:"success,required"`
	Keys    []ConfigurationCredentialUpdateResponseKey   `json:"keys"`
	JSON    configurationCredentialUpdateResponseJSON    `json:"-"`
}

// configurationCredentialUpdateResponseJSON contains the JSON metadata for the
// struct [ConfigurationCredentialUpdateResponse]
type configurationCredentialUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Keys        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConfigurationCredentialUpdateResponseSuccess bool

const (
	ConfigurationCredentialUpdateResponseSuccessTrue ConfigurationCredentialUpdateResponseSuccess = true
)

func (r ConfigurationCredentialUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseSuccessTrue:
		return true
	}
	return false
}

// JSON representation of a JWKS key.
type ConfigurationCredentialUpdateResponseKey struct {
	// Algorithm
	Alg ConfigurationCredentialUpdateResponseKeysAlg `json:"alg,required"`
	// Key ID
	Kid string `json:"kid,required"`
	// Key Type
	Kty ConfigurationCredentialUpdateResponseKeysKty `json:"kty,required"`
	// RSA exponent
	E string `json:"e"`
	// RSA modulus
	N string `json:"n"`
	// X EC coordinate
	X string `json:"x"`
	// Y EC coordinate
	Y     string                                       `json:"y"`
	JSON  configurationCredentialUpdateResponseKeyJSON `json:"-"`
	union ConfigurationCredentialUpdateResponseKeysUnion
}

// configurationCredentialUpdateResponseKeyJSON contains the JSON metadata for the
// struct [ConfigurationCredentialUpdateResponseKey]
type configurationCredentialUpdateResponseKeyJSON struct {
	Alg         apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	E           apijson.Field
	N           apijson.Field
	X           apijson.Field
	Y           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configurationCredentialUpdateResponseKeyJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigurationCredentialUpdateResponseKey) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigurationCredentialUpdateResponseKey{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigurationCredentialUpdateResponseKeysUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEc],
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSA].
func (r ConfigurationCredentialUpdateResponseKey) AsUnion() ConfigurationCredentialUpdateResponseKeysUnion {
	return r.union
}

// JSON representation of a JWKS key.
//
// Union satisfied by
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEc] or
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSA].
type ConfigurationCredentialUpdateResponseKeysUnion interface {
	implementsConfigurationCredentialUpdateResponseKey()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationCredentialUpdateResponseKeysUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEc{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSA{}),
		},
	)
}

// Common properties of a JWT key.
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEc struct {
	// Algorithm
	Alg ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcAlg `json:"alg,required"`
	// Key ID
	Kid string `json:"kid,required"`
	// Key Type
	Kty ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcKty `json:"kty,required"`
	// X EC coordinate
	X string `json:"x,required"`
	// Y EC coordinate
	Y    string                                                                    `json:"y,required"`
	JSON configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcJSON `json:"-"`
}

// configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcJSON
// contains the JSON metadata for the struct
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEc]
type configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcJSON struct {
	Alg         apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	X           apijson.Field
	Y           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEc) implementsConfigurationCredentialUpdateResponseKey() {
}

// Algorithm
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcAlg string

const (
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcAlgEs256 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcAlg = "ES256"
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcAlgEs384 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcAlg = "ES384"
)

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcAlgEs256, ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcAlgEs384:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcKty string

const (
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcKtyEc ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcKty = "EC"
)

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcKtyEc:
		return true
	}
	return false
}

// Common properties of a JWT key.
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSA struct {
	// Algorithm
	Alg ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlg `json:"alg,required"`
	// RSA exponent
	E string `json:"e,required"`
	// Key ID
	Kid string `json:"kid,required"`
	// Key Type
	Kty ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAKty `json:"kty,required"`
	// RSA modulus
	N    string                                                                     `json:"n,required"`
	JSON configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAJSON `json:"-"`
}

// configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAJSON
// contains the JSON metadata for the struct
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSA]
type configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAJSON struct {
	Alg         apijson.Field
	E           apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	N           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSA) implementsConfigurationCredentialUpdateResponseKey() {
}

// Algorithm
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlg string

const (
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlgRs256 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS256"
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlgRs384 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS384"
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlgRs512 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS512"
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlgPs256 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS256"
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlgPs384 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS384"
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlgPs512 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS512"
)

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlgRs256, ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlgRs384, ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlgRs512, ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlgPs256, ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlgPs384, ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlgPs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAKty string

const (
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAKtyRSA ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAKty = "RSA"
)

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAKtyRSA:
		return true
	}
	return false
}

// Algorithm
type ConfigurationCredentialUpdateResponseKeysAlg string

const (
	ConfigurationCredentialUpdateResponseKeysAlgEs256 ConfigurationCredentialUpdateResponseKeysAlg = "ES256"
	ConfigurationCredentialUpdateResponseKeysAlgEs384 ConfigurationCredentialUpdateResponseKeysAlg = "ES384"
	ConfigurationCredentialUpdateResponseKeysAlgRs256 ConfigurationCredentialUpdateResponseKeysAlg = "RS256"
	ConfigurationCredentialUpdateResponseKeysAlgRs384 ConfigurationCredentialUpdateResponseKeysAlg = "RS384"
	ConfigurationCredentialUpdateResponseKeysAlgRs512 ConfigurationCredentialUpdateResponseKeysAlg = "RS512"
	ConfigurationCredentialUpdateResponseKeysAlgPs256 ConfigurationCredentialUpdateResponseKeysAlg = "PS256"
	ConfigurationCredentialUpdateResponseKeysAlgPs384 ConfigurationCredentialUpdateResponseKeysAlg = "PS384"
	ConfigurationCredentialUpdateResponseKeysAlgPs512 ConfigurationCredentialUpdateResponseKeysAlg = "PS512"
)

func (r ConfigurationCredentialUpdateResponseKeysAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAlgEs256, ConfigurationCredentialUpdateResponseKeysAlgEs384, ConfigurationCredentialUpdateResponseKeysAlgRs256, ConfigurationCredentialUpdateResponseKeysAlgRs384, ConfigurationCredentialUpdateResponseKeysAlgRs512, ConfigurationCredentialUpdateResponseKeysAlgPs256, ConfigurationCredentialUpdateResponseKeysAlgPs384, ConfigurationCredentialUpdateResponseKeysAlgPs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateResponseKeysKty string

const (
	ConfigurationCredentialUpdateResponseKeysKtyEc  ConfigurationCredentialUpdateResponseKeysKty = "EC"
	ConfigurationCredentialUpdateResponseKeysKtyRSA ConfigurationCredentialUpdateResponseKeysKty = "RSA"
)

func (r ConfigurationCredentialUpdateResponseKeysKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysKtyEc, ConfigurationCredentialUpdateResponseKeysKtyRSA:
		return true
	}
	return false
}

type ConfigurationCredentialUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string]                                        `path:"zone_id,required"`
	Keys   param.Field[[]ConfigurationCredentialUpdateParamsKeyUnion] `json:"keys"`
}

func (r ConfigurationCredentialUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// JSON representation of a JWKS key.
type ConfigurationCredentialUpdateParamsKey struct {
	// Algorithm
	Alg param.Field[ConfigurationCredentialUpdateParamsKeysAlg] `json:"alg,required"`
	// Key ID
	Kid param.Field[string] `json:"kid,required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialUpdateParamsKeysKty] `json:"kty,required"`
	// RSA exponent
	E param.Field[string] `json:"e"`
	// RSA modulus
	N param.Field[string] `json:"n"`
	// X EC coordinate
	X param.Field[string] `json:"x"`
	// Y EC coordinate
	Y param.Field[string] `json:"y"`
}

func (r ConfigurationCredentialUpdateParamsKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationCredentialUpdateParamsKey) implementsConfigurationCredentialUpdateParamsKeyUnion() {
}

// JSON representation of a JWKS key.
//
// Satisfied by
// [token_validation.ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEc],
// [token_validation.ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSA],
// [ConfigurationCredentialUpdateParamsKey].
type ConfigurationCredentialUpdateParamsKeyUnion interface {
	implementsConfigurationCredentialUpdateParamsKeyUnion()
}

// Common properties of a JWT key.
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEc struct {
	// Algorithm
	Alg param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcAlg] `json:"alg,required"`
	// Key ID
	Kid param.Field[string] `json:"kid,required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcKty] `json:"kty,required"`
	// X EC coordinate
	X param.Field[string] `json:"x,required"`
	// Y EC coordinate
	Y param.Field[string] `json:"y,required"`
}

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEc) implementsConfigurationCredentialUpdateParamsKeyUnion() {
}

// Algorithm
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcAlg string

const (
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcAlgEs256 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcAlg = "ES256"
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcAlgEs384 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcAlg = "ES384"
)

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcAlgEs256, ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcAlgEs384:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcKty string

const (
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcKtyEc ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcKty = "EC"
)

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcKtyEc:
		return true
	}
	return false
}

// Common properties of a JWT key.
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSA struct {
	// Algorithm
	Alg param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlg] `json:"alg,required"`
	// RSA exponent
	E param.Field[string] `json:"e,required"`
	// Key ID
	Kid param.Field[string] `json:"kid,required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAKty] `json:"kty,required"`
	// RSA modulus
	N param.Field[string] `json:"n,required"`
}

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSA) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSA) implementsConfigurationCredentialUpdateParamsKeyUnion() {
}

// Algorithm
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlg string

const (
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlgRs256 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS256"
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlgRs384 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS384"
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlgRs512 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS512"
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlgPs256 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS256"
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlgPs384 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS384"
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlgPs512 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS512"
)

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlgRs256, ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlgRs384, ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlgRs512, ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlgPs256, ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlgPs384, ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlgPs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAKty string

const (
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAKtyRSA ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAKty = "RSA"
)

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAKtyRSA:
		return true
	}
	return false
}

// Algorithm
type ConfigurationCredentialUpdateParamsKeysAlg string

const (
	ConfigurationCredentialUpdateParamsKeysAlgEs256 ConfigurationCredentialUpdateParamsKeysAlg = "ES256"
	ConfigurationCredentialUpdateParamsKeysAlgEs384 ConfigurationCredentialUpdateParamsKeysAlg = "ES384"
	ConfigurationCredentialUpdateParamsKeysAlgRs256 ConfigurationCredentialUpdateParamsKeysAlg = "RS256"
	ConfigurationCredentialUpdateParamsKeysAlgRs384 ConfigurationCredentialUpdateParamsKeysAlg = "RS384"
	ConfigurationCredentialUpdateParamsKeysAlgRs512 ConfigurationCredentialUpdateParamsKeysAlg = "RS512"
	ConfigurationCredentialUpdateParamsKeysAlgPs256 ConfigurationCredentialUpdateParamsKeysAlg = "PS256"
	ConfigurationCredentialUpdateParamsKeysAlgPs384 ConfigurationCredentialUpdateParamsKeysAlg = "PS384"
	ConfigurationCredentialUpdateParamsKeysAlgPs512 ConfigurationCredentialUpdateParamsKeysAlg = "PS512"
)

func (r ConfigurationCredentialUpdateParamsKeysAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAlgEs256, ConfigurationCredentialUpdateParamsKeysAlgEs384, ConfigurationCredentialUpdateParamsKeysAlgRs256, ConfigurationCredentialUpdateParamsKeysAlgRs384, ConfigurationCredentialUpdateParamsKeysAlgRs512, ConfigurationCredentialUpdateParamsKeysAlgPs256, ConfigurationCredentialUpdateParamsKeysAlgPs384, ConfigurationCredentialUpdateParamsKeysAlgPs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateParamsKeysKty string

const (
	ConfigurationCredentialUpdateParamsKeysKtyEc  ConfigurationCredentialUpdateParamsKeysKty = "EC"
	ConfigurationCredentialUpdateParamsKeysKtyRSA ConfigurationCredentialUpdateParamsKeysKty = "RSA"
)

func (r ConfigurationCredentialUpdateParamsKeysKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysKtyEc, ConfigurationCredentialUpdateParamsKeysKtyRSA:
		return true
	}
	return false
}
