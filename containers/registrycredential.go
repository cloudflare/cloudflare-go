// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package containers

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

// RegistryCredentialService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryCredentialService] method instead.
type RegistryCredentialService struct {
	Options []option.RequestOption
}

// NewRegistryCredentialService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegistryCredentialService(opts ...option.RequestOption) (r *RegistryCredentialService) {
	r = &RegistryCredentialService{}
	r.Options = opts
	return
}

// Generates credentials for accessing a configured container image registry.
func (r *RegistryCredentialService) Generate(ctx context.Context, domain string, params RegistryCredentialGenerateParams, opts ...option.RequestOption) (res *RegistryCredentialGenerateResponse, err error) {
	var env RegistryCredentialGenerateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domain == "" {
		err = errors.New("missing required domain parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/registries/%s/credentials", params.AccountID, domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Credentials returned for an authenticated registry configured on a Containers
// account.
type RegistryCredentialGenerateResponse struct {
	// A unique identifier for the user's account.
	AccountID string `json:"account_id" api:"required"`
	// The password to use when authenticating to the image registry.
	Password string `json:"password" api:"required"`
	// The domain of the image registry these credentials target.
	RegistryHost string `json:"registry_host" api:"required"`
	// The username to use when authenticating to the image registry.
	Username string                                 `json:"username" api:"required"`
	JSON     registryCredentialGenerateResponseJSON `json:"-"`
}

// registryCredentialGenerateResponseJSON contains the JSON metadata for the struct
// [RegistryCredentialGenerateResponse]
type registryCredentialGenerateResponseJSON struct {
	AccountID    apijson.Field
	Password     apijson.Field
	RegistryHost apijson.Field
	Username     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RegistryCredentialGenerateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryCredentialGenerateResponseJSON) RawJSON() string {
	return r.raw
}

type RegistryCredentialGenerateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The number of minutes Cloudflare managed registry credentials stay valid.
	// Required for managed registries and must remain positive. Cloudflare ignores
	// this value for external registries.
	ExpirationMinutes param.Field[int64] `json:"expiration_minutes"`
	// The permissions for Cloudflare managed registry credentials. Required for
	// managed registries. Cloudflare ignores this value for external registries.
	Permissions param.Field[[]RegistryCredentialGenerateParamsPermission] `json:"permissions"`
}

func (r RegistryCredentialGenerateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies what permissions the credentials carry.
type RegistryCredentialGenerateParamsPermission string

const (
	RegistryCredentialGenerateParamsPermissionPull RegistryCredentialGenerateParamsPermission = "pull"
	RegistryCredentialGenerateParamsPermissionPush RegistryCredentialGenerateParamsPermission = "push"
	RegistryCredentialGenerateParamsPermissionList RegistryCredentialGenerateParamsPermission = "list"
)

func (r RegistryCredentialGenerateParamsPermission) IsKnown() bool {
	switch r {
	case RegistryCredentialGenerateParamsPermissionPull, RegistryCredentialGenerateParamsPermissionPush, RegistryCredentialGenerateParamsPermissionList:
		return true
	}
	return false
}

type RegistryCredentialGenerateResponseEnvelope struct {
	Errors   []RegistryCredentialGenerateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RegistryCredentialGenerateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Credentials returned for an authenticated registry configured on a Containers
	// account.
	Result RegistryCredentialGenerateResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                           `json:"success" api:"required"`
	JSON    registryCredentialGenerateResponseEnvelopeJSON `json:"-"`
}

// registryCredentialGenerateResponseEnvelopeJSON contains the JSON metadata for
// the struct [RegistryCredentialGenerateResponseEnvelope]
type registryCredentialGenerateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryCredentialGenerateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryCredentialGenerateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegistryCredentialGenerateResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           RegistryCredentialGenerateResponseEnvelopeErrorsSource `json:"source"`
	JSON             registryCredentialGenerateResponseEnvelopeErrorsJSON   `json:"-"`
}

// registryCredentialGenerateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [RegistryCredentialGenerateResponseEnvelopeErrors]
type registryCredentialGenerateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegistryCredentialGenerateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryCredentialGenerateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RegistryCredentialGenerateResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    registryCredentialGenerateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// registryCredentialGenerateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [RegistryCredentialGenerateResponseEnvelopeErrorsSource]
type registryCredentialGenerateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryCredentialGenerateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryCredentialGenerateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegistryCredentialGenerateResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           RegistryCredentialGenerateResponseEnvelopeMessagesSource `json:"source"`
	JSON             registryCredentialGenerateResponseEnvelopeMessagesJSON   `json:"-"`
}

// registryCredentialGenerateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [RegistryCredentialGenerateResponseEnvelopeMessages]
type registryCredentialGenerateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegistryCredentialGenerateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryCredentialGenerateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RegistryCredentialGenerateResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    registryCredentialGenerateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// registryCredentialGenerateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [RegistryCredentialGenerateResponseEnvelopeMessagesSource]
type registryCredentialGenerateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryCredentialGenerateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryCredentialGenerateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
