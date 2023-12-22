// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountImageV1KeyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountImageV1KeyService] method
// instead.
type AccountImageV1KeyService struct {
	Options []option.RequestOption
}

// NewAccountImageV1KeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountImageV1KeyService(opts ...option.RequestOption) (r *AccountImageV1KeyService) {
	r = &AccountImageV1KeyService{}
	r.Options = opts
	return
}

// Lists your signing keys. These can be found on your Cloudflare Images dashboard.
func (r *AccountImageV1KeyService) CloudflareImagesKeysListSigningKeys(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *KeyResponseCollectionNUct6mbe, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/keys", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type KeyResponseCollectionNUct6mbe struct {
	Errors     []KeyResponseCollectionNUct6mbeError    `json:"errors"`
	Messages   []KeyResponseCollectionNUct6mbeMessage  `json:"messages"`
	Result     KeyResponseCollectionNUct6mbeResult     `json:"result"`
	ResultInfo KeyResponseCollectionNUct6mbeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success KeyResponseCollectionNUct6mbeSuccess `json:"success"`
	JSON    keyResponseCollectionNUct6mbeJSON    `json:"-"`
}

// keyResponseCollectionNUct6mbeJSON contains the JSON metadata for the struct
// [KeyResponseCollectionNUct6mbe]
type keyResponseCollectionNUct6mbeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyResponseCollectionNUct6mbe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeyResponseCollectionNUct6mbeError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    keyResponseCollectionNUct6mbeErrorJSON `json:"-"`
}

// keyResponseCollectionNUct6mbeErrorJSON contains the JSON metadata for the struct
// [KeyResponseCollectionNUct6mbeError]
type keyResponseCollectionNUct6mbeErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyResponseCollectionNUct6mbeError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeyResponseCollectionNUct6mbeMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    keyResponseCollectionNUct6mbeMessageJSON `json:"-"`
}

// keyResponseCollectionNUct6mbeMessageJSON contains the JSON metadata for the
// struct [KeyResponseCollectionNUct6mbeMessage]
type keyResponseCollectionNUct6mbeMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyResponseCollectionNUct6mbeMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeyResponseCollectionNUct6mbeResult struct {
	Keys []KeyResponseCollectionNUct6mbeResultKey `json:"keys"`
	JSON keyResponseCollectionNUct6mbeResultJSON  `json:"-"`
}

// keyResponseCollectionNUct6mbeResultJSON contains the JSON metadata for the
// struct [KeyResponseCollectionNUct6mbeResult]
type keyResponseCollectionNUct6mbeResultJSON struct {
	Keys        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyResponseCollectionNUct6mbeResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeyResponseCollectionNUct6mbeResultKey struct {
	// Key name.
	Name string `json:"name"`
	// Key value.
	Value string                                     `json:"value"`
	JSON  keyResponseCollectionNUct6mbeResultKeyJSON `json:"-"`
}

// keyResponseCollectionNUct6mbeResultKeyJSON contains the JSON metadata for the
// struct [KeyResponseCollectionNUct6mbeResultKey]
type keyResponseCollectionNUct6mbeResultKeyJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyResponseCollectionNUct6mbeResultKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeyResponseCollectionNUct6mbeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       keyResponseCollectionNUct6mbeResultInfoJSON `json:"-"`
}

// keyResponseCollectionNUct6mbeResultInfoJSON contains the JSON metadata for the
// struct [KeyResponseCollectionNUct6mbeResultInfo]
type keyResponseCollectionNUct6mbeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyResponseCollectionNUct6mbeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type KeyResponseCollectionNUct6mbeSuccess bool

const (
	KeyResponseCollectionNUct6mbeSuccessTrue KeyResponseCollectionNUct6mbeSuccess = true
)
