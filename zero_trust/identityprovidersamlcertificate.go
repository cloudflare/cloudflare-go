// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// IdentityProviderSAMLCertificateService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityProviderSAMLCertificateService] method instead.
type IdentityProviderSAMLCertificateService struct {
	Options []option.RequestOption
}

// NewIdentityProviderSAMLCertificateService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewIdentityProviderSAMLCertificateService(opts ...option.RequestOption) (r *IdentityProviderSAMLCertificateService) {
	r = &IdentityProviderSAMLCertificateService{}
	r.Options = opts
	return
}

// Creates a new SAML encryption certificate set and assigns it to the specified
// SAML Identity Provider. This endpoint is idempotent - if the IdP already has a
// certificate set assigned, the existing certificate set is returned with a 200
// status.
//
// **Workflow for enabling SAML encryption:**
//
// 1. Call this endpoint to create and assign a certificate set to the IdP
// 2. Update the IdP configuration (PUT `/identity_providers/{id}`) with:
//   - `config.enable_encryption: true`
//   - `saml_certificate_set_id: <uid from step 1>`
//     3. Configure the certificate's public key in your external SAML Identity
//     Provider
func (r *IdentityProviderSAMLCertificateService) New(ctx context.Context, identityProviderID string, body IdentityProviderSAMLCertificateNewParams, opts ...option.RequestOption) (res *IdentityProviderSAMLCertificateNewResponse, err error) {
	var env IdentityProviderSAMLCertificateNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s/saml_certificate", body.AccountID, identityProviderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// IdentityProviderSAMLCertificateNewResponse is a SAML encryption certificate set containing current and optionally previous
// certificates for encryption key rotation.
type IdentityProviderSAMLCertificateNewResponse struct {
	// Timestamp when the certificate set was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unique identifier for the certificate set
	UID string `json:"uid" api:"required" format:"uuid"`
	// Timestamp when the certificate set was last updated (e.g., during rotation)
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The currently active certificate used for encrypting SAML assertions
	CurrentCertificate IdentityProviderSAMLCertificateNewResponseCurrentCertificate `json:"current_certificate"`
	// The previous certificate, maintained during rotation to ensure continuity. Null
	// if no rotation has occurred. Mirrors the structure of `saml_certificate`.
	PreviousCertificate interface{}                                    `json:"previous_certificate" api:"nullable"`
	JSON                identityProviderSAMLCertificateNewResponseJSON `json:"-"`
}

// identityProviderSAMLCertificateNewResponseJSON contains the JSON metadata for
// the struct [IdentityProviderSAMLCertificateNewResponse]
type identityProviderSAMLCertificateNewResponseJSON struct {
	CreatedAt           apijson.Field
	UID                 apijson.Field
	UpdatedAt           apijson.Field
	CurrentCertificate  apijson.Field
	PreviousCertificate apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *IdentityProviderSAMLCertificateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSAMLCertificateNewResponseJSON) RawJSON() string {
	return r.raw
}

// IdentityProviderSAMLCertificateNewResponseCurrentCertificate is the currently active certificate used for encrypting SAML assertions
type IdentityProviderSAMLCertificateNewResponseCurrentCertificate struct {
	// Indicates whether this is the currently active certificate
	IsCurrent bool `json:"is_current" api:"required"`
	// Certificate expiration date. Certificates are automatically rotated 30 days
	// before expiration.
	NotAfter time.Time `json:"not_after" api:"required" format:"date-time"`
	// PEM-encoded X.509 certificate containing the public key. Configure this
	// certificate in your external SAML Identity Provider to enable encryption.
	PublicCertificate string `json:"public_certificate" api:"required"`
	// Unique identifier for the certificate
	UID  string                                                           `json:"uid" api:"required" format:"uuid"`
	JSON identityProviderSAMLCertificateNewResponseCurrentCertificateJSON `json:"-"`
}

// identityProviderSAMLCertificateNewResponseCurrentCertificateJSON contains the
// JSON metadata for the struct
// [IdentityProviderSAMLCertificateNewResponseCurrentCertificate]
type identityProviderSAMLCertificateNewResponseCurrentCertificateJSON struct {
	IsCurrent         apijson.Field
	NotAfter          apijson.Field
	PublicCertificate apijson.Field
	UID               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *IdentityProviderSAMLCertificateNewResponseCurrentCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSAMLCertificateNewResponseCurrentCertificateJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSAMLCertificateNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type IdentityProviderSAMLCertificateNewResponseEnvelope struct {
	Errors   []IdentityProviderSAMLCertificateNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []IdentityProviderSAMLCertificateNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success IdentityProviderSAMLCertificateNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// A SAML encryption certificate set containing current and optionally previous
	// certificates for encryption key rotation.
	Result IdentityProviderSAMLCertificateNewResponse             `json:"result"`
	JSON   identityProviderSAMLCertificateNewResponseEnvelopeJSON `json:"-"`
}

// identityProviderSAMLCertificateNewResponseEnvelopeJSON contains the JSON
// metadata for the struct [IdentityProviderSAMLCertificateNewResponseEnvelope]
type identityProviderSAMLCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderSAMLCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSAMLCertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSAMLCertificateNewResponseEnvelopeErrors struct {
	Code             int64                                                          `json:"code" api:"required"`
	Message          string                                                         `json:"message" api:"required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           IdentityProviderSAMLCertificateNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             identityProviderSAMLCertificateNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// identityProviderSAMLCertificateNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [IdentityProviderSAMLCertificateNewResponseEnvelopeErrors]
type identityProviderSAMLCertificateNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IdentityProviderSAMLCertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSAMLCertificateNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSAMLCertificateNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    identityProviderSAMLCertificateNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// identityProviderSAMLCertificateNewResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [IdentityProviderSAMLCertificateNewResponseEnvelopeErrorsSource]
type identityProviderSAMLCertificateNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderSAMLCertificateNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSAMLCertificateNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSAMLCertificateNewResponseEnvelopeMessages struct {
	Code             int64                                                            `json:"code" api:"required"`
	Message          string                                                           `json:"message" api:"required"`
	DocumentationURL string                                                           `json:"documentation_url"`
	Source           IdentityProviderSAMLCertificateNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             identityProviderSAMLCertificateNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// identityProviderSAMLCertificateNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [IdentityProviderSAMLCertificateNewResponseEnvelopeMessages]
type identityProviderSAMLCertificateNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IdentityProviderSAMLCertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSAMLCertificateNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSAMLCertificateNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                               `json:"pointer"`
	JSON    identityProviderSAMLCertificateNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// identityProviderSAMLCertificateNewResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [IdentityProviderSAMLCertificateNewResponseEnvelopeMessagesSource]
type identityProviderSAMLCertificateNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderSAMLCertificateNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSAMLCertificateNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// IdentityProviderSAMLCertificateNewResponseEnvelopeSuccess indicates whether the API call was successful.
type IdentityProviderSAMLCertificateNewResponseEnvelopeSuccess bool

const (
	IdentityProviderSAMLCertificateNewResponseEnvelopeSuccessTrue IdentityProviderSAMLCertificateNewResponseEnvelopeSuccess = true
)

func (r IdentityProviderSAMLCertificateNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IdentityProviderSAMLCertificateNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
