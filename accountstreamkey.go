// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountStreamKeyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamKeyService] method
// instead.
type AccountStreamKeyService struct {
	Options []option.RequestOption
}

// NewAccountStreamKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamKeyService(opts ...option.RequestOption) (r *AccountStreamKeyService) {
	r = &AccountStreamKeyService{}
	r.Options = opts
	return
}

// Deletes signing keys and revokes all signed URLs generated with the key.
func (r *AccountStreamKeyService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *DeletedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/keys/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates an RSA private key in PEM and JWK formats. Key files are only displayed
// once after creation. Keys are created, used, and deleted independently of
// videos, and every key can sign any video.
func (r *AccountStreamKeyService) StreamSigningKeysNewSigningKeys(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *KeyGenerationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/keys", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Lists the video ID and creation date and time when a signing key was created.
func (r *AccountStreamKeyService) StreamSigningKeysListSigningKeys(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *KeyResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/keys", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type KeyGenerationResponse struct {
	Errors     []KeyGenerationResponseError    `json:"errors"`
	Messages   []KeyGenerationResponseMessage  `json:"messages"`
	Result     KeyGenerationResponseResult     `json:"result"`
	ResultInfo KeyGenerationResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success KeyGenerationResponseSuccess `json:"success"`
	JSON    keyGenerationResponseJSON    `json:"-"`
}

// keyGenerationResponseJSON contains the JSON metadata for the struct
// [KeyGenerationResponse]
type keyGenerationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyGenerationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeyGenerationResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    keyGenerationResponseErrorJSON `json:"-"`
}

// keyGenerationResponseErrorJSON contains the JSON metadata for the struct
// [KeyGenerationResponseError]
type keyGenerationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyGenerationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeyGenerationResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    keyGenerationResponseMessageJSON `json:"-"`
}

// keyGenerationResponseMessageJSON contains the JSON metadata for the struct
// [KeyGenerationResponseMessage]
type keyGenerationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyGenerationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeyGenerationResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time `json:"created" format:"date-time"`
	// The signing key in JWK format.
	Jwk string `json:"jwk"`
	// The signing key in PEM format.
	Pem  string                          `json:"pem"`
	JSON keyGenerationResponseResultJSON `json:"-"`
}

// keyGenerationResponseResultJSON contains the JSON metadata for the struct
// [KeyGenerationResponseResult]
type keyGenerationResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Jwk         apijson.Field
	Pem         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyGenerationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeyGenerationResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                             `json:"total_count"`
	JSON       keyGenerationResponseResultInfoJSON `json:"-"`
}

// keyGenerationResponseResultInfoJSON contains the JSON metadata for the struct
// [KeyGenerationResponseResultInfo]
type keyGenerationResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyGenerationResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type KeyGenerationResponseSuccess bool

const (
	KeyGenerationResponseSuccessTrue KeyGenerationResponseSuccess = true
)

type KeyResponseCollection struct {
	Errors     []KeyResponseCollectionError    `json:"errors"`
	Messages   []KeyResponseCollectionMessage  `json:"messages"`
	Result     []KeyResponseCollectionResult   `json:"result"`
	ResultInfo KeyResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success KeyResponseCollectionSuccess `json:"success"`
	JSON    keyResponseCollectionJSON    `json:"-"`
}

// keyResponseCollectionJSON contains the JSON metadata for the struct
// [KeyResponseCollection]
type keyResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeyResponseCollectionError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    keyResponseCollectionErrorJSON `json:"-"`
}

// keyResponseCollectionErrorJSON contains the JSON metadata for the struct
// [KeyResponseCollectionError]
type keyResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeyResponseCollectionMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    keyResponseCollectionMessageJSON `json:"-"`
}

// keyResponseCollectionMessageJSON contains the JSON metadata for the struct
// [KeyResponseCollectionMessage]
type keyResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeyResponseCollectionResult struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time                       `json:"created" format:"date-time"`
	JSON    keyResponseCollectionResultJSON `json:"-"`
}

// keyResponseCollectionResultJSON contains the JSON metadata for the struct
// [KeyResponseCollectionResult]
type keyResponseCollectionResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeyResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                             `json:"total_count"`
	JSON       keyResponseCollectionResultInfoJSON `json:"-"`
}

// keyResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [KeyResponseCollectionResultInfo]
type keyResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type KeyResponseCollectionSuccess bool

const (
	KeyResponseCollectionSuccessTrue KeyResponseCollectionSuccess = true
)

type DeletedResponse struct {
	Errors   []DeletedResponseError   `json:"errors"`
	Messages []DeletedResponseMessage `json:"messages"`
	Result   string                   `json:"result"`
	// Whether the API call was successful
	Success DeletedResponseSuccess `json:"success"`
	JSON    deletedResponseJSON    `json:"-"`
}

// deletedResponseJSON contains the JSON metadata for the struct [DeletedResponse]
type deletedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeletedResponseError struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    deletedResponseErrorJSON `json:"-"`
}

// deletedResponseErrorJSON contains the JSON metadata for the struct
// [DeletedResponseError]
type deletedResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeletedResponseMessage struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    deletedResponseMessageJSON `json:"-"`
}

// deletedResponseMessageJSON contains the JSON metadata for the struct
// [DeletedResponseMessage]
type deletedResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DeletedResponseSuccess bool

const (
	DeletedResponseSuccessTrue DeletedResponseSuccess = true
)
