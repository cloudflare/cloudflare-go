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
	Errors   api_gateway.Message                        `json:"errors,required"`
	Keys     []ConfigurationCredentialUpdateResponseKey `json:"keys,required"`
	Messages api_gateway.Message                        `json:"messages,required"`
	// Whether the API call was successful.
	Success ConfigurationCredentialUpdateResponseSuccess `json:"success,required"`
	JSON    configurationCredentialUpdateResponseJSON    `json:"-"`
}

// configurationCredentialUpdateResponseJSON contains the JSON metadata for the
// struct [ConfigurationCredentialUpdateResponse]
type configurationCredentialUpdateResponseJSON struct {
	Errors      apijson.Field
	Keys        apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// JSON representation of a JWKS key.
type ConfigurationCredentialUpdateResponseKey struct {
	// Algorithm
	Alg ConfigurationCredentialUpdateResponseKeysAlg `json:"alg,required"`
	// Key ID
	Kid string `json:"kid,required"`
	// Key Type
	Kty ConfigurationCredentialUpdateResponseKeysKty `json:"kty,required"`
	// Curve
	Crv ConfigurationCredentialUpdateResponseKeysCrv `json:"crv"`
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
	Crv         apijson.Field
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
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSA],
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256],
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384].
func (r ConfigurationCredentialUpdateResponseKey) AsUnion() ConfigurationCredentialUpdateResponseKeysUnion {
	return r.union
}

// JSON representation of a JWKS key.
//
// Union satisfied by
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSA],
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256] or
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384].
type ConfigurationCredentialUpdateResponseKeysUnion interface {
	implementsConfigurationCredentialUpdateResponseKey()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationCredentialUpdateResponseKeysUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSA{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384{}),
		},
	)
}

// JSON representation of an RSA key.
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

// JSON representation of an ES256 key
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256 struct {
	// Algorithm
	Alg ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Alg `json:"alg,required"`
	// Curve
	Crv ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Crv `json:"crv,required"`
	// Key ID
	Kid string `json:"kid,required"`
	// Key Type
	Kty ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Kty `json:"kty,required"`
	// X EC coordinate
	X string `json:"x,required"`
	// Y EC coordinate
	Y    string                                                                         `json:"y,required"`
	JSON configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256JSON `json:"-"`
}

// configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256JSON
// contains the JSON metadata for the struct
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256]
type configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256JSON struct {
	Alg         apijson.Field
	Crv         apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	X           apijson.Field
	Y           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256JSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256) implementsConfigurationCredentialUpdateResponseKey() {
}

// Algorithm
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Alg string

const (
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256AlgEs256 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Alg = "ES256"
)

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Alg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256AlgEs256:
		return true
	}
	return false
}

// Curve
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Crv string

const (
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256CrvP256 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Crv = "P-256"
)

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Crv) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256CrvP256:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Kty string

const (
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256KtyEc ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Kty = "EC"
)

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Kty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256KtyEc:
		return true
	}
	return false
}

// JSON representation of an ES384 key
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384 struct {
	// Algorithm
	Alg ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Alg `json:"alg,required"`
	// Curve
	Crv ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Crv `json:"crv,required"`
	// Key ID
	Kid string `json:"kid,required"`
	// Key Type
	Kty ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Kty `json:"kty,required"`
	// X EC coordinate
	X string `json:"x,required"`
	// Y EC coordinate
	Y    string                                                                         `json:"y,required"`
	JSON configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384JSON `json:"-"`
}

// configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384JSON
// contains the JSON metadata for the struct
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384]
type configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384JSON struct {
	Alg         apijson.Field
	Crv         apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	X           apijson.Field
	Y           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384JSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384) implementsConfigurationCredentialUpdateResponseKey() {
}

// Algorithm
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Alg string

const (
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384AlgEs384 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Alg = "ES384"
)

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Alg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384AlgEs384:
		return true
	}
	return false
}

// Curve
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Crv string

const (
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384CrvP384 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Crv = "P-384"
)

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Crv) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384CrvP384:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Kty string

const (
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384KtyEc ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Kty = "EC"
)

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Kty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384KtyEc:
		return true
	}
	return false
}

// Algorithm
type ConfigurationCredentialUpdateResponseKeysAlg string

const (
	ConfigurationCredentialUpdateResponseKeysAlgRs256 ConfigurationCredentialUpdateResponseKeysAlg = "RS256"
	ConfigurationCredentialUpdateResponseKeysAlgRs384 ConfigurationCredentialUpdateResponseKeysAlg = "RS384"
	ConfigurationCredentialUpdateResponseKeysAlgRs512 ConfigurationCredentialUpdateResponseKeysAlg = "RS512"
	ConfigurationCredentialUpdateResponseKeysAlgPs256 ConfigurationCredentialUpdateResponseKeysAlg = "PS256"
	ConfigurationCredentialUpdateResponseKeysAlgPs384 ConfigurationCredentialUpdateResponseKeysAlg = "PS384"
	ConfigurationCredentialUpdateResponseKeysAlgPs512 ConfigurationCredentialUpdateResponseKeysAlg = "PS512"
	ConfigurationCredentialUpdateResponseKeysAlgEs256 ConfigurationCredentialUpdateResponseKeysAlg = "ES256"
	ConfigurationCredentialUpdateResponseKeysAlgEs384 ConfigurationCredentialUpdateResponseKeysAlg = "ES384"
)

func (r ConfigurationCredentialUpdateResponseKeysAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAlgRs256, ConfigurationCredentialUpdateResponseKeysAlgRs384, ConfigurationCredentialUpdateResponseKeysAlgRs512, ConfigurationCredentialUpdateResponseKeysAlgPs256, ConfigurationCredentialUpdateResponseKeysAlgPs384, ConfigurationCredentialUpdateResponseKeysAlgPs512, ConfigurationCredentialUpdateResponseKeysAlgEs256, ConfigurationCredentialUpdateResponseKeysAlgEs384:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateResponseKeysKty string

const (
	ConfigurationCredentialUpdateResponseKeysKtyRSA ConfigurationCredentialUpdateResponseKeysKty = "RSA"
	ConfigurationCredentialUpdateResponseKeysKtyEc  ConfigurationCredentialUpdateResponseKeysKty = "EC"
)

func (r ConfigurationCredentialUpdateResponseKeysKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysKtyRSA, ConfigurationCredentialUpdateResponseKeysKtyEc:
		return true
	}
	return false
}

// Curve
type ConfigurationCredentialUpdateResponseKeysCrv string

const (
	ConfigurationCredentialUpdateResponseKeysCrvP256 ConfigurationCredentialUpdateResponseKeysCrv = "P-256"
	ConfigurationCredentialUpdateResponseKeysCrvP384 ConfigurationCredentialUpdateResponseKeysCrv = "P-384"
)

func (r ConfigurationCredentialUpdateResponseKeysCrv) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysCrvP256, ConfigurationCredentialUpdateResponseKeysCrvP384:
		return true
	}
	return false
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

type ConfigurationCredentialUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string]                                        `path:"zone_id,required"`
	Keys   param.Field[[]ConfigurationCredentialUpdateParamsKeyUnion] `json:"keys,required"`
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
	// Curve
	Crv param.Field[ConfigurationCredentialUpdateParamsKeysCrv] `json:"crv"`
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
// [token_validation.ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSA],
// [token_validation.ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256],
// [token_validation.ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384],
// [ConfigurationCredentialUpdateParamsKey].
type ConfigurationCredentialUpdateParamsKeyUnion interface {
	implementsConfigurationCredentialUpdateParamsKeyUnion()
}

// JSON representation of an RSA key.
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

// JSON representation of an ES256 key
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256 struct {
	// Algorithm
	Alg param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Alg] `json:"alg,required"`
	// Curve
	Crv param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Crv] `json:"crv,required"`
	// Key ID
	Kid param.Field[string] `json:"kid,required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Kty] `json:"kty,required"`
	// X EC coordinate
	X param.Field[string] `json:"x,required"`
	// Y EC coordinate
	Y param.Field[string] `json:"y,required"`
}

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256) implementsConfigurationCredentialUpdateParamsKeyUnion() {
}

// Algorithm
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Alg string

const (
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256AlgEs256 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Alg = "ES256"
)

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Alg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256AlgEs256:
		return true
	}
	return false
}

// Curve
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Crv string

const (
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256CrvP256 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Crv = "P-256"
)

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Crv) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256CrvP256:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Kty string

const (
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256KtyEc ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Kty = "EC"
)

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Kty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256KtyEc:
		return true
	}
	return false
}

// JSON representation of an ES384 key
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384 struct {
	// Algorithm
	Alg param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Alg] `json:"alg,required"`
	// Curve
	Crv param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Crv] `json:"crv,required"`
	// Key ID
	Kid param.Field[string] `json:"kid,required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Kty] `json:"kty,required"`
	// X EC coordinate
	X param.Field[string] `json:"x,required"`
	// Y EC coordinate
	Y param.Field[string] `json:"y,required"`
}

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384) implementsConfigurationCredentialUpdateParamsKeyUnion() {
}

// Algorithm
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Alg string

const (
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384AlgEs384 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Alg = "ES384"
)

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Alg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384AlgEs384:
		return true
	}
	return false
}

// Curve
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Crv string

const (
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384CrvP384 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Crv = "P-384"
)

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Crv) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384CrvP384:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Kty string

const (
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384KtyEc ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Kty = "EC"
)

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Kty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384KtyEc:
		return true
	}
	return false
}

// Algorithm
type ConfigurationCredentialUpdateParamsKeysAlg string

const (
	ConfigurationCredentialUpdateParamsKeysAlgRs256 ConfigurationCredentialUpdateParamsKeysAlg = "RS256"
	ConfigurationCredentialUpdateParamsKeysAlgRs384 ConfigurationCredentialUpdateParamsKeysAlg = "RS384"
	ConfigurationCredentialUpdateParamsKeysAlgRs512 ConfigurationCredentialUpdateParamsKeysAlg = "RS512"
	ConfigurationCredentialUpdateParamsKeysAlgPs256 ConfigurationCredentialUpdateParamsKeysAlg = "PS256"
	ConfigurationCredentialUpdateParamsKeysAlgPs384 ConfigurationCredentialUpdateParamsKeysAlg = "PS384"
	ConfigurationCredentialUpdateParamsKeysAlgPs512 ConfigurationCredentialUpdateParamsKeysAlg = "PS512"
	ConfigurationCredentialUpdateParamsKeysAlgEs256 ConfigurationCredentialUpdateParamsKeysAlg = "ES256"
	ConfigurationCredentialUpdateParamsKeysAlgEs384 ConfigurationCredentialUpdateParamsKeysAlg = "ES384"
)

func (r ConfigurationCredentialUpdateParamsKeysAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAlgRs256, ConfigurationCredentialUpdateParamsKeysAlgRs384, ConfigurationCredentialUpdateParamsKeysAlgRs512, ConfigurationCredentialUpdateParamsKeysAlgPs256, ConfigurationCredentialUpdateParamsKeysAlgPs384, ConfigurationCredentialUpdateParamsKeysAlgPs512, ConfigurationCredentialUpdateParamsKeysAlgEs256, ConfigurationCredentialUpdateParamsKeysAlgEs384:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateParamsKeysKty string

const (
	ConfigurationCredentialUpdateParamsKeysKtyRSA ConfigurationCredentialUpdateParamsKeysKty = "RSA"
	ConfigurationCredentialUpdateParamsKeysKtyEc  ConfigurationCredentialUpdateParamsKeysKty = "EC"
)

func (r ConfigurationCredentialUpdateParamsKeysKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysKtyRSA, ConfigurationCredentialUpdateParamsKeysKtyEc:
		return true
	}
	return false
}

// Curve
type ConfigurationCredentialUpdateParamsKeysCrv string

const (
	ConfigurationCredentialUpdateParamsKeysCrvP256 ConfigurationCredentialUpdateParamsKeysCrv = "P-256"
	ConfigurationCredentialUpdateParamsKeysCrvP384 ConfigurationCredentialUpdateParamsKeysCrv = "P-384"
)

func (r ConfigurationCredentialUpdateParamsKeysCrv) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysCrvP256, ConfigurationCredentialUpdateParamsKeysCrvP384:
		return true
	}
	return false
}
