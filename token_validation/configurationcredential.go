// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/api_gateway"
	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
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

// Replaces the configuration's complete key set. Symmetric keys must include their
// key material.
func (r *ConfigurationCredentialService) Update(ctx context.Context, configID string, params ConfigurationCredentialUpdateParams, opts ...option.RequestOption) (res *ConfigurationCredentialUpdateResponse, err error) {
	var env ConfigurationCredentialUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/token_validation/config/%s/credentials", params.ZoneID, configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates the configuration's complete key set while allowing omitted fields on
// existing keys to retain stored values. Omitted key identities are removed.
func (r *ConfigurationCredentialService) Edit(ctx context.Context, configID string, params ConfigurationCredentialEditParams, opts ...option.RequestOption) (res *ConfigurationCredentialEditResponse, err error) {
	var env ConfigurationCredentialEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/token_validation/config/%s/credentials", params.ZoneID, configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ConfigurationCredentialUpdateResponse struct {
	Keys []ConfigurationCredentialUpdateResponseKey `json:"keys" api:"required"`
	JSON configurationCredentialUpdateResponseJSON  `json:"-"`
}

// configurationCredentialUpdateResponseJSON contains the JSON metadata for the
// struct [ConfigurationCredentialUpdateResponse]
type configurationCredentialUpdateResponseJSON struct {
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

// JSON representation of a JWKS key.
type ConfigurationCredentialUpdateResponseKey struct {
	// Algorithm
	Alg ConfigurationCredentialUpdateResponseKeysAlg `json:"alg" api:"required"`
	// Key ID
	Kid string `json:"kid" api:"required"`
	// Key Type
	Kty ConfigurationCredentialUpdateResponseKeysKty `json:"kty" api:"required"`
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
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384],
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponse].
func (r ConfigurationCredentialUpdateResponseKey) AsUnion() ConfigurationCredentialUpdateResponseKeysUnion {
	return r.union
}

// JSON representation of a JWKS key.
//
// Union satisfied by
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSA],
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256],
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384] or
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponse].
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
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponse{}),
		},
	)
}

// JSON representation of an RSA key.
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSA struct {
	// Algorithm
	Alg ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAAlg `json:"alg" api:"required"`
	// RSA exponent
	E string `json:"e" api:"required"`
	// Key ID
	Kid string `json:"kid" api:"required"`
	// Key Type
	Kty ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyRSAKty `json:"kty" api:"required"`
	// RSA modulus
	N    string                                                                     `json:"n" api:"required"`
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
	Alg ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Alg `json:"alg" api:"required"`
	// Curve
	Crv ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Crv `json:"crv" api:"required"`
	// Key ID
	Kid string `json:"kid" api:"required"`
	// Key Type
	Kty ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs256Kty `json:"kty" api:"required"`
	// X EC coordinate
	X string `json:"x" api:"required"`
	// Y EC coordinate
	Y    string                                                                         `json:"y" api:"required"`
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
	Alg ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Alg `json:"alg" api:"required"`
	// Curve
	Crv ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Crv `json:"crv" api:"required"`
	// Key ID
	Kid string `json:"kid" api:"required"`
	// Key Type
	Kty ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyEcEs384Kty `json:"kty" api:"required"`
	// X EC coordinate
	X string `json:"x" api:"required"`
	// Y EC coordinate
	Y    string                                                                         `json:"y" api:"required"`
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

// JSON representation of a symmetric verification key in API responses (secret
// material is redacted).
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponse struct {
	// Algorithm
	Alg ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlg `json:"alg" api:"required"`
	// Key ID
	Kid string `json:"kid" api:"required"`
	// Key Type
	Kty  ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseKty  `json:"kty" api:"required"`
	JSON configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseJSON `json:"-"`
}

// configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseJSON
// contains the JSON metadata for the struct
// [ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponse]
type configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseJSON struct {
	Alg         apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponse) implementsConfigurationCredentialUpdateResponseKey() {
}

// Algorithm
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlg string

const (
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlgHs256 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlg = "HS256"
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlgHs384 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlg = "HS384"
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlgHs512 ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlg = "HS512"
)

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlgHs256, ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlgHs384, ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlgHs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseKty string

const (
	ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseKtyOct ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseKty = "oct"
)

func (r ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAPIShieldCredentialsJWTKeyOctResponseKtyOct:
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
	ConfigurationCredentialUpdateResponseKeysAlgHs256 ConfigurationCredentialUpdateResponseKeysAlg = "HS256"
	ConfigurationCredentialUpdateResponseKeysAlgHs384 ConfigurationCredentialUpdateResponseKeysAlg = "HS384"
	ConfigurationCredentialUpdateResponseKeysAlgHs512 ConfigurationCredentialUpdateResponseKeysAlg = "HS512"
)

func (r ConfigurationCredentialUpdateResponseKeysAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysAlgRs256, ConfigurationCredentialUpdateResponseKeysAlgRs384, ConfigurationCredentialUpdateResponseKeysAlgRs512, ConfigurationCredentialUpdateResponseKeysAlgPs256, ConfigurationCredentialUpdateResponseKeysAlgPs384, ConfigurationCredentialUpdateResponseKeysAlgPs512, ConfigurationCredentialUpdateResponseKeysAlgEs256, ConfigurationCredentialUpdateResponseKeysAlgEs384, ConfigurationCredentialUpdateResponseKeysAlgHs256, ConfigurationCredentialUpdateResponseKeysAlgHs384, ConfigurationCredentialUpdateResponseKeysAlgHs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateResponseKeysKty string

const (
	ConfigurationCredentialUpdateResponseKeysKtyRSA ConfigurationCredentialUpdateResponseKeysKty = "RSA"
	ConfigurationCredentialUpdateResponseKeysKtyEc  ConfigurationCredentialUpdateResponseKeysKty = "EC"
	ConfigurationCredentialUpdateResponseKeysKtyOct ConfigurationCredentialUpdateResponseKeysKty = "oct"
)

func (r ConfigurationCredentialUpdateResponseKeysKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseKeysKtyRSA, ConfigurationCredentialUpdateResponseKeysKtyEc, ConfigurationCredentialUpdateResponseKeysKtyOct:
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

type ConfigurationCredentialEditResponse struct {
	Keys []ConfigurationCredentialEditResponseKey `json:"keys" api:"required"`
	JSON configurationCredentialEditResponseJSON  `json:"-"`
}

// configurationCredentialEditResponseJSON contains the JSON metadata for the
// struct [ConfigurationCredentialEditResponse]
type configurationCredentialEditResponseJSON struct {
	Keys        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialEditResponseJSON) RawJSON() string {
	return r.raw
}

// JSON representation of a JWKS key.
type ConfigurationCredentialEditResponseKey struct {
	// Algorithm
	Alg ConfigurationCredentialEditResponseKeysAlg `json:"alg" api:"required"`
	// Key ID
	Kid string `json:"kid" api:"required"`
	// Key Type
	Kty ConfigurationCredentialEditResponseKeysKty `json:"kty" api:"required"`
	// Curve
	Crv ConfigurationCredentialEditResponseKeysCrv `json:"crv"`
	// RSA exponent
	E string `json:"e"`
	// RSA modulus
	N string `json:"n"`
	// X EC coordinate
	X string `json:"x"`
	// Y EC coordinate
	Y     string                                     `json:"y"`
	JSON  configurationCredentialEditResponseKeyJSON `json:"-"`
	union ConfigurationCredentialEditResponseKeysUnion
}

// configurationCredentialEditResponseKeyJSON contains the JSON metadata for the
// struct [ConfigurationCredentialEditResponseKey]
type configurationCredentialEditResponseKeyJSON struct {
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

func (r configurationCredentialEditResponseKeyJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigurationCredentialEditResponseKey) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigurationCredentialEditResponseKey{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigurationCredentialEditResponseKeysUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSA],
// [ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256],
// [ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384],
// [ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponse].
func (r ConfigurationCredentialEditResponseKey) AsUnion() ConfigurationCredentialEditResponseKeysUnion {
	return r.union
}

// JSON representation of a JWKS key.
//
// Union satisfied by
// [ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSA],
// [ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256],
// [ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384] or
// [ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponse].
type ConfigurationCredentialEditResponseKeysUnion interface {
	implementsConfigurationCredentialEditResponseKey()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationCredentialEditResponseKeysUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSA{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponse{}),
		},
	)
}

// JSON representation of an RSA key.
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSA struct {
	// Algorithm
	Alg ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlg `json:"alg" api:"required"`
	// RSA exponent
	E string `json:"e" api:"required"`
	// Key ID
	Kid string `json:"kid" api:"required"`
	// Key Type
	Kty ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAKty `json:"kty" api:"required"`
	// RSA modulus
	N    string                                                                   `json:"n" api:"required"`
	JSON configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAJSON `json:"-"`
}

// configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAJSON
// contains the JSON metadata for the struct
// [ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSA]
type configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAJSON struct {
	Alg         apijson.Field
	E           apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	N           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSA) implementsConfigurationCredentialEditResponseKey() {
}

// Algorithm
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlg string

const (
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlgRs256 ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS256"
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlgRs384 ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS384"
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlgRs512 ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS512"
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlgPs256 ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS256"
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlgPs384 ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS384"
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlgPs512 ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS512"
)

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlgRs256, ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlgRs384, ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlgRs512, ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlgPs256, ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlgPs384, ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAAlgPs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAKty string

const (
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAKtyRSA ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAKty = "RSA"
)

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyRSAKtyRSA:
		return true
	}
	return false
}

// JSON representation of an ES256 key
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256 struct {
	// Algorithm
	Alg ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256Alg `json:"alg" api:"required"`
	// Curve
	Crv ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256Crv `json:"crv" api:"required"`
	// Key ID
	Kid string `json:"kid" api:"required"`
	// Key Type
	Kty ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256Kty `json:"kty" api:"required"`
	// X EC coordinate
	X string `json:"x" api:"required"`
	// Y EC coordinate
	Y    string                                                                       `json:"y" api:"required"`
	JSON configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256JSON `json:"-"`
}

// configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256JSON
// contains the JSON metadata for the struct
// [ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256]
type configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256JSON struct {
	Alg         apijson.Field
	Crv         apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	X           apijson.Field
	Y           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256JSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256) implementsConfigurationCredentialEditResponseKey() {
}

// Algorithm
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256Alg string

const (
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256AlgEs256 ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256Alg = "ES256"
)

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256Alg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256AlgEs256:
		return true
	}
	return false
}

// Curve
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256Crv string

const (
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256CrvP256 ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256Crv = "P-256"
)

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256Crv) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256CrvP256:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256Kty string

const (
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256KtyEc ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256Kty = "EC"
)

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256Kty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs256KtyEc:
		return true
	}
	return false
}

// JSON representation of an ES384 key
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384 struct {
	// Algorithm
	Alg ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384Alg `json:"alg" api:"required"`
	// Curve
	Crv ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384Crv `json:"crv" api:"required"`
	// Key ID
	Kid string `json:"kid" api:"required"`
	// Key Type
	Kty ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384Kty `json:"kty" api:"required"`
	// X EC coordinate
	X string `json:"x" api:"required"`
	// Y EC coordinate
	Y    string                                                                       `json:"y" api:"required"`
	JSON configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384JSON `json:"-"`
}

// configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384JSON
// contains the JSON metadata for the struct
// [ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384]
type configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384JSON struct {
	Alg         apijson.Field
	Crv         apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	X           apijson.Field
	Y           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384JSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384) implementsConfigurationCredentialEditResponseKey() {
}

// Algorithm
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384Alg string

const (
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384AlgEs384 ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384Alg = "ES384"
)

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384Alg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384AlgEs384:
		return true
	}
	return false
}

// Curve
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384Crv string

const (
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384CrvP384 ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384Crv = "P-384"
)

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384Crv) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384CrvP384:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384Kty string

const (
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384KtyEc ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384Kty = "EC"
)

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384Kty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyEcEs384KtyEc:
		return true
	}
	return false
}

// JSON representation of a symmetric verification key in API responses (secret
// material is redacted).
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponse struct {
	// Algorithm
	Alg ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlg `json:"alg" api:"required"`
	// Key ID
	Kid string `json:"kid" api:"required"`
	// Key Type
	Kty  ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseKty  `json:"kty" api:"required"`
	JSON configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseJSON `json:"-"`
}

// configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseJSON
// contains the JSON metadata for the struct
// [ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponse]
type configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseJSON struct {
	Alg         apijson.Field
	Kid         apijson.Field
	Kty         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponse) implementsConfigurationCredentialEditResponseKey() {
}

// Algorithm
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlg string

const (
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlgHs256 ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlg = "HS256"
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlgHs384 ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlg = "HS384"
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlgHs512 ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlg = "HS512"
)

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlgHs256, ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlgHs384, ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseAlgHs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseKty string

const (
	ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseKtyOct ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseKty = "oct"
)

func (r ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseKeysAPIShieldCredentialsJWTKeyOctResponseKtyOct:
		return true
	}
	return false
}

// Algorithm
type ConfigurationCredentialEditResponseKeysAlg string

const (
	ConfigurationCredentialEditResponseKeysAlgRs256 ConfigurationCredentialEditResponseKeysAlg = "RS256"
	ConfigurationCredentialEditResponseKeysAlgRs384 ConfigurationCredentialEditResponseKeysAlg = "RS384"
	ConfigurationCredentialEditResponseKeysAlgRs512 ConfigurationCredentialEditResponseKeysAlg = "RS512"
	ConfigurationCredentialEditResponseKeysAlgPs256 ConfigurationCredentialEditResponseKeysAlg = "PS256"
	ConfigurationCredentialEditResponseKeysAlgPs384 ConfigurationCredentialEditResponseKeysAlg = "PS384"
	ConfigurationCredentialEditResponseKeysAlgPs512 ConfigurationCredentialEditResponseKeysAlg = "PS512"
	ConfigurationCredentialEditResponseKeysAlgEs256 ConfigurationCredentialEditResponseKeysAlg = "ES256"
	ConfigurationCredentialEditResponseKeysAlgEs384 ConfigurationCredentialEditResponseKeysAlg = "ES384"
	ConfigurationCredentialEditResponseKeysAlgHs256 ConfigurationCredentialEditResponseKeysAlg = "HS256"
	ConfigurationCredentialEditResponseKeysAlgHs384 ConfigurationCredentialEditResponseKeysAlg = "HS384"
	ConfigurationCredentialEditResponseKeysAlgHs512 ConfigurationCredentialEditResponseKeysAlg = "HS512"
)

func (r ConfigurationCredentialEditResponseKeysAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseKeysAlgRs256, ConfigurationCredentialEditResponseKeysAlgRs384, ConfigurationCredentialEditResponseKeysAlgRs512, ConfigurationCredentialEditResponseKeysAlgPs256, ConfigurationCredentialEditResponseKeysAlgPs384, ConfigurationCredentialEditResponseKeysAlgPs512, ConfigurationCredentialEditResponseKeysAlgEs256, ConfigurationCredentialEditResponseKeysAlgEs384, ConfigurationCredentialEditResponseKeysAlgHs256, ConfigurationCredentialEditResponseKeysAlgHs384, ConfigurationCredentialEditResponseKeysAlgHs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialEditResponseKeysKty string

const (
	ConfigurationCredentialEditResponseKeysKtyRSA ConfigurationCredentialEditResponseKeysKty = "RSA"
	ConfigurationCredentialEditResponseKeysKtyEc  ConfigurationCredentialEditResponseKeysKty = "EC"
	ConfigurationCredentialEditResponseKeysKtyOct ConfigurationCredentialEditResponseKeysKty = "oct"
)

func (r ConfigurationCredentialEditResponseKeysKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseKeysKtyRSA, ConfigurationCredentialEditResponseKeysKtyEc, ConfigurationCredentialEditResponseKeysKtyOct:
		return true
	}
	return false
}

// Curve
type ConfigurationCredentialEditResponseKeysCrv string

const (
	ConfigurationCredentialEditResponseKeysCrvP256 ConfigurationCredentialEditResponseKeysCrv = "P-256"
	ConfigurationCredentialEditResponseKeysCrvP384 ConfigurationCredentialEditResponseKeysCrv = "P-384"
)

func (r ConfigurationCredentialEditResponseKeysCrv) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseKeysCrvP256, ConfigurationCredentialEditResponseKeysCrvP384:
		return true
	}
	return false
}

type ConfigurationCredentialUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string]                                        `path:"zone_id" api:"required"`
	Keys   param.Field[[]ConfigurationCredentialUpdateParamsKeyUnion] `json:"keys" api:"required"`
}

func (r ConfigurationCredentialUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// JSON representation of a JWKS key for create and PUT requests.
type ConfigurationCredentialUpdateParamsKey struct {
	// Algorithm
	Alg param.Field[ConfigurationCredentialUpdateParamsKeysAlg] `json:"alg" api:"required"`
	// Key ID
	Kid param.Field[string] `json:"kid" api:"required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialUpdateParamsKeysKty] `json:"kty" api:"required"`
	// Curve
	Crv param.Field[ConfigurationCredentialUpdateParamsKeysCrv] `json:"crv"`
	// RSA exponent
	E param.Field[string] `json:"e"`
	// Symmetric key material. Required for create and PUT update requests.
	K param.Field[string] `json:"k"`
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

// JSON representation of a JWKS key for create and PUT requests.
//
// Satisfied by
// [token_validation.ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSA],
// [token_validation.ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256],
// [token_validation.ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384],
// [token_validation.ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequest],
// [ConfigurationCredentialUpdateParamsKey].
type ConfigurationCredentialUpdateParamsKeyUnion interface {
	implementsConfigurationCredentialUpdateParamsKeyUnion()
}

// JSON representation of an RSA key.
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSA struct {
	// Algorithm
	Alg param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAAlg] `json:"alg" api:"required"`
	// RSA exponent
	E param.Field[string] `json:"e" api:"required"`
	// Key ID
	Kid param.Field[string] `json:"kid" api:"required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyRSAKty] `json:"kty" api:"required"`
	// RSA modulus
	N param.Field[string] `json:"n" api:"required"`
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
	Alg param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Alg] `json:"alg" api:"required"`
	// Curve
	Crv param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Crv] `json:"crv" api:"required"`
	// Key ID
	Kid param.Field[string] `json:"kid" api:"required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256Kty] `json:"kty" api:"required"`
	// X EC coordinate
	X param.Field[string] `json:"x" api:"required"`
	// Y EC coordinate
	Y param.Field[string] `json:"y" api:"required"`
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
	Alg param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Alg] `json:"alg" api:"required"`
	// Curve
	Crv param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Crv] `json:"crv" api:"required"`
	// Key ID
	Kid param.Field[string] `json:"kid" api:"required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs384Kty] `json:"kty" api:"required"`
	// X EC coordinate
	X param.Field[string] `json:"x" api:"required"`
	// Y EC coordinate
	Y param.Field[string] `json:"y" api:"required"`
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

// JSON representation of a symmetric key for create/PUT requests.
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequest struct {
	// Algorithm
	Alg param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestAlg] `json:"alg" api:"required"`
	// Symmetric key material. Required for create and PUT update requests.
	K param.Field[string] `json:"k" api:"required"`
	// Key ID
	Kid param.Field[string] `json:"kid" api:"required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestKty] `json:"kty" api:"required"`
}

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequest) implementsConfigurationCredentialUpdateParamsKeyUnion() {
}

// Algorithm
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestAlg string

const (
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestAlgHs256 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestAlg = "HS256"
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestAlgHs384 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestAlg = "HS384"
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestAlgHs512 ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestAlg = "HS512"
)

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestAlgHs256, ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestAlgHs384, ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestAlgHs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestKty string

const (
	ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestKtyOct ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestKty = "oct"
)

func (r ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyOctRequestKtyOct:
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
	ConfigurationCredentialUpdateParamsKeysAlgHs256 ConfigurationCredentialUpdateParamsKeysAlg = "HS256"
	ConfigurationCredentialUpdateParamsKeysAlgHs384 ConfigurationCredentialUpdateParamsKeysAlg = "HS384"
	ConfigurationCredentialUpdateParamsKeysAlgHs512 ConfigurationCredentialUpdateParamsKeysAlg = "HS512"
)

func (r ConfigurationCredentialUpdateParamsKeysAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysAlgRs256, ConfigurationCredentialUpdateParamsKeysAlgRs384, ConfigurationCredentialUpdateParamsKeysAlgRs512, ConfigurationCredentialUpdateParamsKeysAlgPs256, ConfigurationCredentialUpdateParamsKeysAlgPs384, ConfigurationCredentialUpdateParamsKeysAlgPs512, ConfigurationCredentialUpdateParamsKeysAlgEs256, ConfigurationCredentialUpdateParamsKeysAlgEs384, ConfigurationCredentialUpdateParamsKeysAlgHs256, ConfigurationCredentialUpdateParamsKeysAlgHs384, ConfigurationCredentialUpdateParamsKeysAlgHs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialUpdateParamsKeysKty string

const (
	ConfigurationCredentialUpdateParamsKeysKtyRSA ConfigurationCredentialUpdateParamsKeysKty = "RSA"
	ConfigurationCredentialUpdateParamsKeysKtyEc  ConfigurationCredentialUpdateParamsKeysKty = "EC"
	ConfigurationCredentialUpdateParamsKeysKtyOct ConfigurationCredentialUpdateParamsKeysKty = "oct"
)

func (r ConfigurationCredentialUpdateParamsKeysKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateParamsKeysKtyRSA, ConfigurationCredentialUpdateParamsKeysKtyEc, ConfigurationCredentialUpdateParamsKeysKtyOct:
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

type ConfigurationCredentialUpdateResponseEnvelope struct {
	Errors   api_gateway.Message                   `json:"errors" api:"required"`
	Messages api_gateway.Message                   `json:"messages" api:"required"`
	Result   ConfigurationCredentialUpdateResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ConfigurationCredentialUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    configurationCredentialUpdateResponseEnvelopeJSON    `json:"-"`
}

// configurationCredentialUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ConfigurationCredentialUpdateResponseEnvelope]
type configurationCredentialUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConfigurationCredentialUpdateResponseEnvelopeSuccess bool

const (
	ConfigurationCredentialUpdateResponseEnvelopeSuccessTrue ConfigurationCredentialUpdateResponseEnvelopeSuccess = true
)

func (r ConfigurationCredentialUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigurationCredentialUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigurationCredentialEditParams struct {
	// Identifier.
	ZoneID param.Field[string]                                      `path:"zone_id" api:"required"`
	Keys   param.Field[[]ConfigurationCredentialEditParamsKeyUnion] `json:"keys" api:"required"`
}

func (r ConfigurationCredentialEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// JSON representation of a JWKS key for PATCH credentials requests.
type ConfigurationCredentialEditParamsKey struct {
	// Algorithm
	Alg param.Field[ConfigurationCredentialEditParamsKeysAlg] `json:"alg" api:"required"`
	// Key ID
	Kid param.Field[string] `json:"kid" api:"required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialEditParamsKeysKty] `json:"kty" api:"required"`
	// Curve
	Crv param.Field[ConfigurationCredentialEditParamsKeysCrv] `json:"crv"`
	// RSA exponent
	E param.Field[string] `json:"e"`
	// Symmetric key material. Optional for PATCH: omit to preserve existing secret for
	// matching `{alg,kid}`; send a string to rotate. `k: null` is invalid.
	K param.Field[string] `json:"k"`
	// RSA modulus
	N param.Field[string] `json:"n"`
	// X EC coordinate
	X param.Field[string] `json:"x"`
	// Y EC coordinate
	Y param.Field[string] `json:"y"`
}

func (r ConfigurationCredentialEditParamsKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationCredentialEditParamsKey) implementsConfigurationCredentialEditParamsKeyUnion() {}

// JSON representation of a JWKS key for PATCH credentials requests.
//
// Satisfied by
// [token_validation.ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSA],
// [token_validation.ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256],
// [token_validation.ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384],
// [token_validation.ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequest],
// [ConfigurationCredentialEditParamsKey].
type ConfigurationCredentialEditParamsKeyUnion interface {
	implementsConfigurationCredentialEditParamsKeyUnion()
}

// JSON representation of an RSA key.
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSA struct {
	// Algorithm
	Alg param.Field[ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlg] `json:"alg" api:"required"`
	// RSA exponent
	E param.Field[string] `json:"e" api:"required"`
	// Key ID
	Kid param.Field[string] `json:"kid" api:"required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAKty] `json:"kty" api:"required"`
	// RSA modulus
	N param.Field[string] `json:"n" api:"required"`
}

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSA) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSA) implementsConfigurationCredentialEditParamsKeyUnion() {
}

// Algorithm
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlg string

const (
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlgRs256 ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS256"
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlgRs384 ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS384"
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlgRs512 ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlg = "RS512"
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlgPs256 ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS256"
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlgPs384 ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS384"
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlgPs512 ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlg = "PS512"
)

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlgRs256, ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlgRs384, ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlgRs512, ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlgPs256, ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlgPs384, ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAAlgPs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAKty string

const (
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAKtyRSA ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAKty = "RSA"
)

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyRSAKtyRSA:
		return true
	}
	return false
}

// JSON representation of an ES256 key
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256 struct {
	// Algorithm
	Alg param.Field[ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256Alg] `json:"alg" api:"required"`
	// Curve
	Crv param.Field[ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256Crv] `json:"crv" api:"required"`
	// Key ID
	Kid param.Field[string] `json:"kid" api:"required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256Kty] `json:"kty" api:"required"`
	// X EC coordinate
	X param.Field[string] `json:"x" api:"required"`
	// Y EC coordinate
	Y param.Field[string] `json:"y" api:"required"`
}

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256) implementsConfigurationCredentialEditParamsKeyUnion() {
}

// Algorithm
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256Alg string

const (
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256AlgEs256 ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256Alg = "ES256"
)

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256Alg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256AlgEs256:
		return true
	}
	return false
}

// Curve
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256Crv string

const (
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256CrvP256 ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256Crv = "P-256"
)

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256Crv) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256CrvP256:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256Kty string

const (
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256KtyEc ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256Kty = "EC"
)

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256Kty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs256KtyEc:
		return true
	}
	return false
}

// JSON representation of an ES384 key
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384 struct {
	// Algorithm
	Alg param.Field[ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384Alg] `json:"alg" api:"required"`
	// Curve
	Crv param.Field[ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384Crv] `json:"crv" api:"required"`
	// Key ID
	Kid param.Field[string] `json:"kid" api:"required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384Kty] `json:"kty" api:"required"`
	// X EC coordinate
	X param.Field[string] `json:"x" api:"required"`
	// Y EC coordinate
	Y param.Field[string] `json:"y" api:"required"`
}

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384) implementsConfigurationCredentialEditParamsKeyUnion() {
}

// Algorithm
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384Alg string

const (
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384AlgEs384 ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384Alg = "ES384"
)

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384Alg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384AlgEs384:
		return true
	}
	return false
}

// Curve
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384Crv string

const (
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384CrvP384 ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384Crv = "P-384"
)

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384Crv) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384CrvP384:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384Kty string

const (
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384KtyEc ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384Kty = "EC"
)

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384Kty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyEcEs384KtyEc:
		return true
	}
	return false
}

// JSON representation of a symmetric key for PATCH requests.
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequest struct {
	// Algorithm
	Alg param.Field[ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestAlg] `json:"alg" api:"required"`
	// Key ID
	Kid param.Field[string] `json:"kid" api:"required"`
	// Key Type
	Kty param.Field[ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestKty] `json:"kty" api:"required"`
	// Symmetric key material. Optional for PATCH: omit to preserve existing secret for
	// matching `{alg,kid}`; send a string to rotate. `k: null` is invalid.
	K param.Field[string] `json:"k"`
}

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequest) implementsConfigurationCredentialEditParamsKeyUnion() {
}

// Algorithm
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestAlg string

const (
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestAlgHs256 ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestAlg = "HS256"
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestAlgHs384 ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestAlg = "HS384"
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestAlgHs512 ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestAlg = "HS512"
)

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestAlgHs256, ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestAlgHs384, ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestAlgHs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestKty string

const (
	ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestKtyOct ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestKty = "oct"
)

func (r ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditParamsKeysAPIShieldCredentialsJWTKeyOctPatchRequestKtyOct:
		return true
	}
	return false
}

// Algorithm
type ConfigurationCredentialEditParamsKeysAlg string

const (
	ConfigurationCredentialEditParamsKeysAlgRs256 ConfigurationCredentialEditParamsKeysAlg = "RS256"
	ConfigurationCredentialEditParamsKeysAlgRs384 ConfigurationCredentialEditParamsKeysAlg = "RS384"
	ConfigurationCredentialEditParamsKeysAlgRs512 ConfigurationCredentialEditParamsKeysAlg = "RS512"
	ConfigurationCredentialEditParamsKeysAlgPs256 ConfigurationCredentialEditParamsKeysAlg = "PS256"
	ConfigurationCredentialEditParamsKeysAlgPs384 ConfigurationCredentialEditParamsKeysAlg = "PS384"
	ConfigurationCredentialEditParamsKeysAlgPs512 ConfigurationCredentialEditParamsKeysAlg = "PS512"
	ConfigurationCredentialEditParamsKeysAlgEs256 ConfigurationCredentialEditParamsKeysAlg = "ES256"
	ConfigurationCredentialEditParamsKeysAlgEs384 ConfigurationCredentialEditParamsKeysAlg = "ES384"
	ConfigurationCredentialEditParamsKeysAlgHs256 ConfigurationCredentialEditParamsKeysAlg = "HS256"
	ConfigurationCredentialEditParamsKeysAlgHs384 ConfigurationCredentialEditParamsKeysAlg = "HS384"
	ConfigurationCredentialEditParamsKeysAlgHs512 ConfigurationCredentialEditParamsKeysAlg = "HS512"
)

func (r ConfigurationCredentialEditParamsKeysAlg) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditParamsKeysAlgRs256, ConfigurationCredentialEditParamsKeysAlgRs384, ConfigurationCredentialEditParamsKeysAlgRs512, ConfigurationCredentialEditParamsKeysAlgPs256, ConfigurationCredentialEditParamsKeysAlgPs384, ConfigurationCredentialEditParamsKeysAlgPs512, ConfigurationCredentialEditParamsKeysAlgEs256, ConfigurationCredentialEditParamsKeysAlgEs384, ConfigurationCredentialEditParamsKeysAlgHs256, ConfigurationCredentialEditParamsKeysAlgHs384, ConfigurationCredentialEditParamsKeysAlgHs512:
		return true
	}
	return false
}

// Key Type
type ConfigurationCredentialEditParamsKeysKty string

const (
	ConfigurationCredentialEditParamsKeysKtyRSA ConfigurationCredentialEditParamsKeysKty = "RSA"
	ConfigurationCredentialEditParamsKeysKtyEc  ConfigurationCredentialEditParamsKeysKty = "EC"
	ConfigurationCredentialEditParamsKeysKtyOct ConfigurationCredentialEditParamsKeysKty = "oct"
)

func (r ConfigurationCredentialEditParamsKeysKty) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditParamsKeysKtyRSA, ConfigurationCredentialEditParamsKeysKtyEc, ConfigurationCredentialEditParamsKeysKtyOct:
		return true
	}
	return false
}

// Curve
type ConfigurationCredentialEditParamsKeysCrv string

const (
	ConfigurationCredentialEditParamsKeysCrvP256 ConfigurationCredentialEditParamsKeysCrv = "P-256"
	ConfigurationCredentialEditParamsKeysCrvP384 ConfigurationCredentialEditParamsKeysCrv = "P-384"
)

func (r ConfigurationCredentialEditParamsKeysCrv) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditParamsKeysCrvP256, ConfigurationCredentialEditParamsKeysCrvP384:
		return true
	}
	return false
}

type ConfigurationCredentialEditResponseEnvelope struct {
	Errors   api_gateway.Message                 `json:"errors" api:"required"`
	Messages api_gateway.Message                 `json:"messages" api:"required"`
	Result   ConfigurationCredentialEditResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ConfigurationCredentialEditResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    configurationCredentialEditResponseEnvelopeJSON    `json:"-"`
}

// configurationCredentialEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ConfigurationCredentialEditResponseEnvelope]
type configurationCredentialEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationCredentialEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationCredentialEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConfigurationCredentialEditResponseEnvelopeSuccess bool

const (
	ConfigurationCredentialEditResponseEnvelopeSuccessTrue ConfigurationCredentialEditResponseEnvelopeSuccess = true
)

func (r ConfigurationCredentialEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigurationCredentialEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
