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

// AccessUserPasskeyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessUserPasskeyService] method instead.
type AccessUserPasskeyService struct {
	Options []option.RequestOption
}

// NewAccessUserPasskeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessUserPasskeyService(opts ...option.RequestOption) (r *AccessUserPasskeyService) {
	r = &AccessUserPasskeyService{}
	r.Options = opts
	return
}

// Alias of the authenticator deletion endpoint for passkey management. The
// identifier is the passkey's authenticator ID from the last-seen identity
// response. Successful deletion revokes the enrollment and returns a null result.
func (r *AccessUserPasskeyService) Delete(ctx context.Context, userID string, authenticatorID string, body AccessUserPasskeyDeleteParams, opts ...option.RequestOption) (res *AccessUserPasskeyDeleteResponse, err error) {
	var env AccessUserPasskeyDeleteResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/access/users/%s/passkeys/%s", body.AccountID, userID, authenticatorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AccessUserPasskeyDeleteResponse = interface{}

type AccessUserPasskeyDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessUserPasskeyDeleteResponseEnvelope struct {
	Errors   []AccessUserPasskeyDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessUserPasskeyDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessUserPasskeyDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessUserPasskeyDeleteResponse                `json:"result" api:"nullable"`
	JSON    accessUserPasskeyDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessUserPasskeyDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessUserPasskeyDeleteResponseEnvelope]
type accessUserPasskeyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPasskeyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPasskeyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserPasskeyDeleteResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           AccessUserPasskeyDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserPasskeyDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserPasskeyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessUserPasskeyDeleteResponseEnvelopeErrors]
type accessUserPasskeyDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPasskeyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPasskeyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserPasskeyDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    accessUserPasskeyDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserPasskeyDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessUserPasskeyDeleteResponseEnvelopeErrorsSource]
type accessUserPasskeyDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPasskeyDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPasskeyDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserPasskeyDeleteResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           AccessUserPasskeyDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserPasskeyDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserPasskeyDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessUserPasskeyDeleteResponseEnvelopeMessages]
type accessUserPasskeyDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPasskeyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPasskeyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserPasskeyDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    accessUserPasskeyDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserPasskeyDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessUserPasskeyDeleteResponseEnvelopeMessagesSource]
type accessUserPasskeyDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPasskeyDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPasskeyDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserPasskeyDeleteResponseEnvelopeSuccess bool

const (
	AccessUserPasskeyDeleteResponseEnvelopeSuccessTrue AccessUserPasskeyDeleteResponseEnvelopeSuccess = true
)

func (r AccessUserPasskeyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserPasskeyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
