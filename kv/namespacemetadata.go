// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kv

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// NamespaceMetadataService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceMetadataService] method instead.
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
	var env NamespaceMetadataGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/metadata/%s", query.AccountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NamespaceMetadataGetResponse map[string]interface{}

type NamespaceMetadataGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type NamespaceMetadataGetResponseEnvelope struct {
	Errors   []NamespaceMetadataGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NamespaceMetadataGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success NamespaceMetadataGetResponseEnvelopeSuccess `json:"success,required"`
	// Arbitrary JSON that is associated with a key.
	Result NamespaceMetadataGetResponse             `json:"result"`
	JSON   namespaceMetadataGetResponseEnvelopeJSON `json:"-"`
}

// namespaceMetadataGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceMetadataGetResponseEnvelope]
type namespaceMetadataGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceMetadataGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceMetadataGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceMetadataGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    namespaceMetadataGetResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceMetadataGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [NamespaceMetadataGetResponseEnvelopeErrors]
type namespaceMetadataGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceMetadataGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceMetadataGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceMetadataGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    namespaceMetadataGetResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceMetadataGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [NamespaceMetadataGetResponseEnvelopeMessages]
type namespaceMetadataGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceMetadataGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceMetadataGetResponseEnvelopeMessagesJSON) RawJSON() string {
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
