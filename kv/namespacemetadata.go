// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kv

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// NamespaceMetadataService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewNamespaceMetadataService] method
// instead.
type NamespaceMetadataService struct {
	Options []option.RequestOption
}

// NewNamespaceMetadataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNamespaceMetadataService(opts ...option.RequestOption) (r *NamespaceMetadataService) {
	r = &NamespaceMetadataService{}
	r.Options = opts
	return
}

// Returns the metadata associated with the given key in the given namespace. Use
// URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key
// name.
func (r *NamespaceMetadataService) Get(ctx context.Context, namespaceID string, keyName string, query NamespaceMetadataGetParams, opts ...option.RequestOption) (res *NamespaceMetadataGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NamespaceMetadataGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/metadata/%s", query.AccountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NamespaceMetadataGetResponse = interface{}

type NamespaceMetadataGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type NamespaceMetadataGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Arbitrary JSON that is associated with a key.
	Result NamespaceMetadataGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success NamespaceMetadataGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    namespaceMetadataGetResponseEnvelopeJSON    `json:"-"`
}

// namespaceMetadataGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceMetadataGetResponseEnvelope]
type namespaceMetadataGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceMetadataGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceMetadataGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceMetadataGetResponseEnvelopeSuccess bool

const (
	NamespaceMetadataGetResponseEnvelopeSuccessTrue NamespaceMetadataGetResponseEnvelopeSuccess = true
)

func (r NamespaceMetadataGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceMetadataGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
