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
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// RegistryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryService] method instead.
type RegistryService struct {
	Options     []option.RequestOption
	Credentials *RegistryCredentialService
}

// NewRegistryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegistryService(opts ...option.RequestOption) (r *RegistryService) {
	r = &RegistryService{}
	r.Options = opts
	r.Credentials = NewRegistryCredentialService(opts...)
	return
}

// Registers credentials for a supported private external image registry so
// Containers can pull images from it. This endpoint does not create a registry or
// upload an image. Public Docker Hub images and images in the Cloudflare managed
// registry do not require this configuration.
//
// Refer to
// [Image management](https://developers.cloudflare.com/containers/platform-details/image-management/)
// for supported registries and instructions for storing registry credentials.
func (r *RegistryService) New(ctx context.Context, params RegistryNewParams, opts ...option.RequestOption) (res *RegistryNewResponse, err error) {
	var env RegistryNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/registries", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get the list of configured registries in the account.
func (r *RegistryService) List(ctx context.Context, query RegistryListParams, opts ...option.RequestOption) (res *pagination.SinglePage[RegistryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/registries", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Get the list of configured registries in the account.
func (r *RegistryService) ListAutoPaging(ctx context.Context, query RegistryListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RegistryListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a registry from the account, this will prevent Containers from pulling
// images from the registry.
func (r *RegistryService) Delete(ctx context.Context, domain string, body RegistryDeleteParams, opts ...option.RequestOption) (res *RegistryDeleteResponse, err error) {
	var env RegistryDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domain == "" {
		err = errors.New("missing required domain parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/registries/%s", body.AccountID, domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// An image registry added in a customer account.
type RegistryNewResponse struct {
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// A string representation of a domain name. See RFC-1034
	// (https://www.ietf.org/rfc/rfc1034.txt). Consider that the limit of a domain name
	// is min 3 and max 253 ASCII characters.
	Domain string `json:"domain" api:"required"`
	// The type of registry that is being configured.
	Kind RegistryNewResponseKind `json:"kind"`
	// Public component of the registry credentials. For managed registries this is a
	// base64-encoded public key; for external registries the format depends on the
	// registry provider.
	PublicKey string                  `json:"public_key"`
	JSON      registryNewResponseJSON `json:"-"`
}

// registryNewResponseJSON contains the JSON metadata for the struct
// [RegistryNewResponse]
type registryNewResponseJSON struct {
	CreatedAt   apijson.Field
	Domain      apijson.Field
	Kind        apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryNewResponseJSON) RawJSON() string {
	return r.raw
}

// The type of registry that is being configured.
type RegistryNewResponseKind string

const (
	RegistryNewResponseKindEcr       RegistryNewResponseKind = "ECR"
	RegistryNewResponseKindDockerHub RegistryNewResponseKind = "DockerHub"
	RegistryNewResponseKindGar       RegistryNewResponseKind = "GAR"
	RegistryNewResponseKindDefault   RegistryNewResponseKind = "default"
)

func (r RegistryNewResponseKind) IsKnown() bool {
	switch r {
	case RegistryNewResponseKindEcr, RegistryNewResponseKindDockerHub, RegistryNewResponseKindGar, RegistryNewResponseKindDefault:
		return true
	}
	return false
}

// An image registry added in a customer account.
type RegistryListResponse struct {
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// A string representation of a domain name. See RFC-1034
	// (https://www.ietf.org/rfc/rfc1034.txt). Consider that the limit of a domain name
	// is min 3 and max 253 ASCII characters.
	Domain string `json:"domain" api:"required"`
	// The type of registry that is being configured.
	Kind RegistryListResponseKind `json:"kind"`
	// Public component of the registry credentials. For managed registries this is a
	// base64-encoded public key; for external registries the format depends on the
	// registry provider.
	PublicKey string                   `json:"public_key"`
	JSON      registryListResponseJSON `json:"-"`
}

// registryListResponseJSON contains the JSON metadata for the struct
// [RegistryListResponse]
type registryListResponseJSON struct {
	CreatedAt   apijson.Field
	Domain      apijson.Field
	Kind        apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryListResponseJSON) RawJSON() string {
	return r.raw
}

// The type of registry that is being configured.
type RegistryListResponseKind string

const (
	RegistryListResponseKindEcr       RegistryListResponseKind = "ECR"
	RegistryListResponseKindDockerHub RegistryListResponseKind = "DockerHub"
	RegistryListResponseKindGar       RegistryListResponseKind = "GAR"
	RegistryListResponseKindDefault   RegistryListResponseKind = "default"
)

func (r RegistryListResponseKind) IsKnown() bool {
	switch r {
	case RegistryListResponseKindEcr, RegistryListResponseKindDockerHub, RegistryListResponseKindGar, RegistryListResponseKindDefault:
		return true
	}
	return false
}

// Result of deleting an image registry from a Containers account.
type RegistryDeleteResponse struct {
	// A string representation of a domain name. See RFC-1034
	// (https://www.ietf.org/rfc/rfc1034.txt). Consider that the limit of a domain name
	// is min 3 and max 253 ASCII characters.
	Domain string                     `json:"domain" api:"required"`
	JSON   registryDeleteResponseJSON `json:"-"`
}

// registryDeleteResponseJSON contains the JSON metadata for the struct
// [RegistryDeleteResponse]
type registryDeleteResponseJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RegistryNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Credentials for authenticating to a private external image registry. Store the
	// private credential in
	// [Secrets Store](https://developers.cloudflare.com/secrets-store/) before calling
	// the API. Refer to
	// [Image management](https://developers.cloudflare.com/containers/platform-details/image-management/)
	// for the credential required by each supported registry provider.
	Auth param.Field[RegistryNewParamsAuth] `json:"auth" api:"required"`
	// Hostname of the private registry, without a scheme or image path. Supported
	// hostnames are `docker.io`, AWS ECR hostnames, and Google Artifact Registry
	// `*-docker.pkg.dev` hostnames.
	Domain param.Field[string] `json:"domain" api:"required"`
	// Registry provider. This must match `domain`: `DockerHub` for `docker.io`, `ECR`
	// for AWS ECR, or `GAR` for Google Artifact Registry.
	Kind param.Field[RegistryNewParamsKind] `json:"kind" api:"required"`
	// Omit this field or set it to `false`. Public Docker Hub images do not require
	// registry configuration and cannot be added with this endpoint.
	IsPublic param.Field[RegistryNewParamsIsPublic] `json:"is_public"`
}

func (r RegistryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Credentials for authenticating to a private external image registry. Store the
// private credential in
// [Secrets Store](https://developers.cloudflare.com/secrets-store/) before calling
// the API. Refer to
// [Image management](https://developers.cloudflare.com/containers/platform-details/image-management/)
// for the credential required by each supported registry provider.
type RegistryNewParamsAuth struct {
	// A reference to the private registry credential in Secrets Store. The referenced
	// secret must have the `containers` scope. Raw secret values are not accepted.
	PrivateCredential param.Field[RegistryNewParamsAuthPrivateCredential] `json:"private_credential" api:"required"`
	// The non-secret part of the registry credential: an AWS access key ID for ECR, a
	// username for Docker Hub, or a service account email for Google Artifact
	// Registry.
	PublicCredential param.Field[string] `json:"public_credential" api:"required"`
}

func (r RegistryNewParamsAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A reference to the private registry credential in Secrets Store. The referenced
// secret must have the `containers` scope. Raw secret values are not accepted.
type RegistryNewParamsAuthPrivateCredential struct {
	// Name of the secret within the store.
	SecretName param.Field[string] `json:"secret_name" api:"required"`
	// Identifier of the Secrets Store containing the secret.
	StoreID param.Field[string] `json:"store_id" api:"required"`
}

func (r RegistryNewParamsAuthPrivateCredential) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Registry provider. This must match `domain`: `DockerHub` for `docker.io`, `ECR`
// for AWS ECR, or `GAR` for Google Artifact Registry.
type RegistryNewParamsKind string

const (
	RegistryNewParamsKindEcr       RegistryNewParamsKind = "ECR"
	RegistryNewParamsKindDockerHub RegistryNewParamsKind = "DockerHub"
	RegistryNewParamsKindGar       RegistryNewParamsKind = "GAR"
)

func (r RegistryNewParamsKind) IsKnown() bool {
	switch r {
	case RegistryNewParamsKindEcr, RegistryNewParamsKindDockerHub, RegistryNewParamsKindGar:
		return true
	}
	return false
}

// Omit this field or set it to `false`. Public Docker Hub images do not require
// registry configuration and cannot be added with this endpoint.
type RegistryNewParamsIsPublic bool

const (
	RegistryNewParamsIsPublicFalse RegistryNewParamsIsPublic = false
)

func (r RegistryNewParamsIsPublic) IsKnown() bool {
	switch r {
	case RegistryNewParamsIsPublicFalse:
		return true
	}
	return false
}

type RegistryNewResponseEnvelope struct {
	Errors   []RegistryNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RegistryNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// An image registry added in a customer account.
	Result RegistryNewResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                            `json:"success" api:"required"`
	JSON    registryNewResponseEnvelopeJSON `json:"-"`
}

// registryNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RegistryNewResponseEnvelope]
type registryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegistryNewResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code" api:"required"`
	Message          string                                  `json:"message" api:"required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           RegistryNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             registryNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// registryNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RegistryNewResponseEnvelopeErrors]
type registryNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegistryNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RegistryNewResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    registryNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// registryNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RegistryNewResponseEnvelopeErrorsSource]
type registryNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegistryNewResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           RegistryNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             registryNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// registryNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RegistryNewResponseEnvelopeMessages]
type registryNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegistryNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RegistryNewResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    registryNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// registryNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RegistryNewResponseEnvelopeMessagesSource]
type registryNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type RegistryListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RegistryDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RegistryDeleteResponseEnvelope struct {
	Errors   []RegistryDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RegistryDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Result of deleting an image registry from a Containers account.
	Result RegistryDeleteResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                               `json:"success" api:"required"`
	JSON    registryDeleteResponseEnvelopeJSON `json:"-"`
}

// registryDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RegistryDeleteResponseEnvelope]
type registryDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegistryDeleteResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           RegistryDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             registryDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// registryDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RegistryDeleteResponseEnvelopeErrors]
type registryDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegistryDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RegistryDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    registryDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// registryDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RegistryDeleteResponseEnvelopeErrorsSource]
type registryDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegistryDeleteResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           RegistryDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             registryDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// registryDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RegistryDeleteResponseEnvelopeMessages]
type registryDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegistryDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RegistryDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    registryDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// registryDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RegistryDeleteResponseEnvelopeMessagesSource]
type registryDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
