// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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
)

// AccessUserMfaAuthenticatorService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessUserMfaAuthenticatorService] method instead.
type AccessUserMfaAuthenticatorService struct {
	Options []option.RequestOption
}

// NewAccessUserMfaAuthenticatorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessUserMfaAuthenticatorService(opts ...option.RequestOption) (r *AccessUserMfaAuthenticatorService) {
	r = &AccessUserMfaAuthenticatorService{}
	r.Options = opts
	return
}

// Deletes a specific MFA device or passkey, including a PIV key or FIDO2 key
// enrollment, for a user. Admin deletion remains available when MFA is disabled.
// Successful deletion revokes the enrollment and returns a null result.
func (r *AccessUserMfaAuthenticatorService) Delete(ctx context.Context, userID string, authenticatorID string, body AccessUserMfaAuthenticatorDeleteParams, opts ...option.RequestOption) (res *AccessUserMfaAuthenticatorDeleteResponse, err error) {
	var env AccessUserMfaAuthenticatorDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	if authenticatorID == "" {
		err = errors.New("missing required authenticator_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/users/%s/mfa_authenticators/%s", body.AccountID, userID, authenticatorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AccessUserMfaAuthenticatorDeleteResponse = interface{}

type AccessUserMfaAuthenticatorDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessUserMfaAuthenticatorDeleteResponseEnvelope struct {
	Errors   []AccessUserMfaAuthenticatorDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessUserMfaAuthenticatorDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessUserMfaAuthenticatorDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessUserMfaAuthenticatorDeleteResponse                `json:"result" api:"nullable"`
	JSON    accessUserMfaAuthenticatorDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessUserMfaAuthenticatorDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessUserMfaAuthenticatorDeleteResponseEnvelope]
type accessUserMfaAuthenticatorDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserMfaAuthenticatorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserMfaAuthenticatorDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserMfaAuthenticatorDeleteResponseEnvelopeErrors struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           AccessUserMfaAuthenticatorDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserMfaAuthenticatorDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserMfaAuthenticatorDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessUserMfaAuthenticatorDeleteResponseEnvelopeErrors]
type accessUserMfaAuthenticatorDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserMfaAuthenticatorDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserMfaAuthenticatorDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserMfaAuthenticatorDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    accessUserMfaAuthenticatorDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserMfaAuthenticatorDeleteResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [AccessUserMfaAuthenticatorDeleteResponseEnvelopeErrorsSource]
type accessUserMfaAuthenticatorDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserMfaAuthenticatorDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserMfaAuthenticatorDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserMfaAuthenticatorDeleteResponseEnvelopeMessages struct {
	Code             int64                                                          `json:"code" api:"required"`
	Message          string                                                         `json:"message" api:"required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           AccessUserMfaAuthenticatorDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserMfaAuthenticatorDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserMfaAuthenticatorDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AccessUserMfaAuthenticatorDeleteResponseEnvelopeMessages]
type accessUserMfaAuthenticatorDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserMfaAuthenticatorDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserMfaAuthenticatorDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserMfaAuthenticatorDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    accessUserMfaAuthenticatorDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserMfaAuthenticatorDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccessUserMfaAuthenticatorDeleteResponseEnvelopeMessagesSource]
type accessUserMfaAuthenticatorDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserMfaAuthenticatorDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserMfaAuthenticatorDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserMfaAuthenticatorDeleteResponseEnvelopeSuccess bool

const (
	AccessUserMfaAuthenticatorDeleteResponseEnvelopeSuccessTrue AccessUserMfaAuthenticatorDeleteResponseEnvelopeSuccess = true
)

func (r AccessUserMfaAuthenticatorDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserMfaAuthenticatorDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
