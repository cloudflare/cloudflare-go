// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AccountStorageKvNamespaceValueService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountStorageKvNamespaceValueService] method instead.
type AccountStorageKvNamespaceValueService struct {
	Options []option.RequestOption
}

// NewAccountStorageKvNamespaceValueService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStorageKvNamespaceValueService(opts ...option.RequestOption) (r *AccountStorageKvNamespaceValueService) {
	r = &AccountStorageKvNamespaceValueService{}
	r.Options = opts
	return
}

// Returns the value associated with the given key in the given namespace. Use
// URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key
// name. If the KV-pair is set to expire at some point, the expiration time as
// measured in seconds since the UNIX epoch will be returned in the `expiration`
// response header.
func (r *AccountStorageKvNamespaceValueService) Get(ctx context.Context, accountIdentifier string, namespaceIdentifier string, keyName string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountIdentifier, namespaceIdentifier, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Write a value identified by a key. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name. Body should be the value to be
// stored along with JSON metadata to be associated with the key/value pair.
// Existing values, expirations, and metadata will be overwritten. If neither
// `expiration` nor `expiration_ttl` is specified, the key-value pair will never
// expire. If both are set, `expiration_ttl` is used and `expiration` is ignored.
func (r *AccountStorageKvNamespaceValueService) Update(ctx context.Context, accountIdentifier string, namespaceIdentifier string, keyName string, body AccountStorageKvNamespaceValueUpdateParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceValueUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountIdentifier, namespaceIdentifier, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Remove a KV pair from the namespace. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name.
func (r *AccountStorageKvNamespaceValueService) Delete(ctx context.Context, accountIdentifier string, namespaceIdentifier string, keyName string, opts ...option.RequestOption) (res *AccountStorageKvNamespaceValueDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountIdentifier, namespaceIdentifier, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountStorageKvNamespaceValueUpdateResponse struct {
	Errors   []AccountStorageKvNamespaceValueUpdateResponseError   `json:"errors"`
	Messages []AccountStorageKvNamespaceValueUpdateResponseMessage `json:"messages"`
	Result   AccountStorageKvNamespaceValueUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStorageKvNamespaceValueUpdateResponseSuccess `json:"success"`
	JSON    accountStorageKvNamespaceValueUpdateResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceValueUpdateResponseJSON contains the JSON metadata for
// the struct [AccountStorageKvNamespaceValueUpdateResponse]
type accountStorageKvNamespaceValueUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceValueUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceValueUpdateResponseError struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountStorageKvNamespaceValueUpdateResponseErrorJSON `json:"-"`
}

// accountStorageKvNamespaceValueUpdateResponseErrorJSON contains the JSON metadata
// for the struct [AccountStorageKvNamespaceValueUpdateResponseError]
type accountStorageKvNamespaceValueUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceValueUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceValueUpdateResponseMessage struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accountStorageKvNamespaceValueUpdateResponseMessageJSON `json:"-"`
}

// accountStorageKvNamespaceValueUpdateResponseMessageJSON contains the JSON
// metadata for the struct [AccountStorageKvNamespaceValueUpdateResponseMessage]
type accountStorageKvNamespaceValueUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceValueUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccountStorageKvNamespaceValueUpdateResponseResultUnknown]
// or [shared.UnionString].
type AccountStorageKvNamespaceValueUpdateResponseResult interface {
	ImplementsAccountStorageKvNamespaceValueUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountStorageKvNamespaceValueUpdateResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type AccountStorageKvNamespaceValueUpdateResponseSuccess bool

const (
	AccountStorageKvNamespaceValueUpdateResponseSuccessTrue AccountStorageKvNamespaceValueUpdateResponseSuccess = true
)

type AccountStorageKvNamespaceValueDeleteResponse struct {
	Errors   []AccountStorageKvNamespaceValueDeleteResponseError   `json:"errors"`
	Messages []AccountStorageKvNamespaceValueDeleteResponseMessage `json:"messages"`
	Result   AccountStorageKvNamespaceValueDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStorageKvNamespaceValueDeleteResponseSuccess `json:"success"`
	JSON    accountStorageKvNamespaceValueDeleteResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceValueDeleteResponseJSON contains the JSON metadata for
// the struct [AccountStorageKvNamespaceValueDeleteResponse]
type accountStorageKvNamespaceValueDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceValueDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceValueDeleteResponseError struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountStorageKvNamespaceValueDeleteResponseErrorJSON `json:"-"`
}

// accountStorageKvNamespaceValueDeleteResponseErrorJSON contains the JSON metadata
// for the struct [AccountStorageKvNamespaceValueDeleteResponseError]
type accountStorageKvNamespaceValueDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceValueDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageKvNamespaceValueDeleteResponseMessage struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accountStorageKvNamespaceValueDeleteResponseMessageJSON `json:"-"`
}

// accountStorageKvNamespaceValueDeleteResponseMessageJSON contains the JSON
// metadata for the struct [AccountStorageKvNamespaceValueDeleteResponseMessage]
type accountStorageKvNamespaceValueDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceValueDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccountStorageKvNamespaceValueDeleteResponseResultUnknown]
// or [shared.UnionString].
type AccountStorageKvNamespaceValueDeleteResponseResult interface {
	ImplementsAccountStorageKvNamespaceValueDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountStorageKvNamespaceValueDeleteResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type AccountStorageKvNamespaceValueDeleteResponseSuccess bool

const (
	AccountStorageKvNamespaceValueDeleteResponseSuccessTrue AccountStorageKvNamespaceValueDeleteResponseSuccess = true
)

type AccountStorageKvNamespaceValueUpdateParams struct {
	// Arbitrary JSON to be associated with a key/value pair.
	Metadata param.Field[string] `json:"metadata,required"`
	// A byte sequence to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value,required"`
}

func (r AccountStorageKvNamespaceValueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
