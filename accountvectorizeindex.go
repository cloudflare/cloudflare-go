// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountVectorizeIndexService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountVectorizeIndexService]
// method instead.
type AccountVectorizeIndexService struct {
	Options []option.RequestOption
}

// NewAccountVectorizeIndexService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountVectorizeIndexService(opts ...option.RequestOption) (r *AccountVectorizeIndexService) {
	r = &AccountVectorizeIndexService{}
	r.Options = opts
	return
}

// Creates and returns a new Vectorize Index.
func (r *AccountVectorizeIndexService) New(ctx context.Context, accountIdentifier string, body AccountVectorizeIndexNewParams, opts ...option.RequestOption) (res *AccountVectorizeIndexNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/vectorize/indexes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the specified Vectorize Index.
func (r *AccountVectorizeIndexService) Get(ctx context.Context, accountIdentifier string, indexName string, opts ...option.RequestOption) (res *AccountVectorizeIndexGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates and returns the specified Vectorize Index.
func (r *AccountVectorizeIndexService) Update(ctx context.Context, accountIdentifier string, indexName string, body AccountVectorizeIndexUpdateParams, opts ...option.RequestOption) (res *AccountVectorizeIndexUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a list of Vectorize Indexes
func (r *AccountVectorizeIndexService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountVectorizeIndexListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/vectorize/indexes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes the specified Vectorize Index.
func (r *AccountVectorizeIndexService) Delete(ctx context.Context, accountIdentifier string, indexName string, opts ...option.RequestOption) (res *AccountVectorizeIndexDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Inserts vectors into the specified index and returns the count of the vectors
// successfully inserted.
func (r *AccountVectorizeIndexService) Insert(ctx context.Context, accountIdentifier string, indexName string, opts ...option.RequestOption) (res *AccountVectorizeIndexInsertResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/insert", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Finds vectors closest to a given vector in an index.
func (r *AccountVectorizeIndexService) Query(ctx context.Context, accountIdentifier string, indexName string, body AccountVectorizeIndexQueryParams, opts ...option.RequestOption) (res *AccountVectorizeIndexQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/query", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Upserts vectors into the specified index, creating them if they do not exist and
// returns the count of values and ids successfully inserted.
func (r *AccountVectorizeIndexService) Upsert(ctx context.Context, accountIdentifier string, indexName string, opts ...option.RequestOption) (res *AccountVectorizeIndexUpsertResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/upsert", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccountVectorizeIndexNewResponse struct {
	Errors   []AccountVectorizeIndexNewResponseError   `json:"errors"`
	Messages []AccountVectorizeIndexNewResponseMessage `json:"messages"`
	Result   shared.VectorizeCreateIndexResponse       `json:"result"`
	// Whether the API call was successful
	Success AccountVectorizeIndexNewResponseSuccess `json:"success"`
	JSON    accountVectorizeIndexNewResponseJSON    `json:"-"`
}

// accountVectorizeIndexNewResponseJSON contains the JSON metadata for the struct
// [AccountVectorizeIndexNewResponse]
type accountVectorizeIndexNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexNewResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountVectorizeIndexNewResponseErrorJSON `json:"-"`
}

// accountVectorizeIndexNewResponseErrorJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexNewResponseError]
type accountVectorizeIndexNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexNewResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountVectorizeIndexNewResponseMessageJSON `json:"-"`
}

// accountVectorizeIndexNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexNewResponseMessage]
type accountVectorizeIndexNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountVectorizeIndexNewResponseSuccess bool

const (
	AccountVectorizeIndexNewResponseSuccessTrue AccountVectorizeIndexNewResponseSuccess = true
)

type AccountVectorizeIndexGetResponse struct {
	Errors   []AccountVectorizeIndexGetResponseError   `json:"errors"`
	Messages []AccountVectorizeIndexGetResponseMessage `json:"messages"`
	Result   shared.VectorizeCreateIndexResponse       `json:"result"`
	// Whether the API call was successful
	Success AccountVectorizeIndexGetResponseSuccess `json:"success"`
	JSON    accountVectorizeIndexGetResponseJSON    `json:"-"`
}

// accountVectorizeIndexGetResponseJSON contains the JSON metadata for the struct
// [AccountVectorizeIndexGetResponse]
type accountVectorizeIndexGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexGetResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountVectorizeIndexGetResponseErrorJSON `json:"-"`
}

// accountVectorizeIndexGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexGetResponseError]
type accountVectorizeIndexGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexGetResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountVectorizeIndexGetResponseMessageJSON `json:"-"`
}

// accountVectorizeIndexGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexGetResponseMessage]
type accountVectorizeIndexGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountVectorizeIndexGetResponseSuccess bool

const (
	AccountVectorizeIndexGetResponseSuccessTrue AccountVectorizeIndexGetResponseSuccess = true
)

type AccountVectorizeIndexUpdateResponse struct {
	Errors   []AccountVectorizeIndexUpdateResponseError   `json:"errors"`
	Messages []AccountVectorizeIndexUpdateResponseMessage `json:"messages"`
	Result   shared.VectorizeCreateIndexResponse          `json:"result"`
	// Whether the API call was successful
	Success AccountVectorizeIndexUpdateResponseSuccess `json:"success"`
	JSON    accountVectorizeIndexUpdateResponseJSON    `json:"-"`
}

// accountVectorizeIndexUpdateResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexUpdateResponse]
type accountVectorizeIndexUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexUpdateResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountVectorizeIndexUpdateResponseErrorJSON `json:"-"`
}

// accountVectorizeIndexUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexUpdateResponseError]
type accountVectorizeIndexUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexUpdateResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountVectorizeIndexUpdateResponseMessageJSON `json:"-"`
}

// accountVectorizeIndexUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountVectorizeIndexUpdateResponseMessage]
type accountVectorizeIndexUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountVectorizeIndexUpdateResponseSuccess bool

const (
	AccountVectorizeIndexUpdateResponseSuccessTrue AccountVectorizeIndexUpdateResponseSuccess = true
)

type AccountVectorizeIndexListResponse struct {
	Errors   []AccountVectorizeIndexListResponseError   `json:"errors"`
	Messages []AccountVectorizeIndexListResponseMessage `json:"messages"`
	Result   []shared.VectorizeCreateIndexResponse      `json:"result"`
	// Whether the API call was successful
	Success AccountVectorizeIndexListResponseSuccess `json:"success"`
	JSON    accountVectorizeIndexListResponseJSON    `json:"-"`
}

// accountVectorizeIndexListResponseJSON contains the JSON metadata for the struct
// [AccountVectorizeIndexListResponse]
type accountVectorizeIndexListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexListResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountVectorizeIndexListResponseErrorJSON `json:"-"`
}

// accountVectorizeIndexListResponseErrorJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexListResponseError]
type accountVectorizeIndexListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexListResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountVectorizeIndexListResponseMessageJSON `json:"-"`
}

// accountVectorizeIndexListResponseMessageJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexListResponseMessage]
type accountVectorizeIndexListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountVectorizeIndexListResponseSuccess bool

const (
	AccountVectorizeIndexListResponseSuccessTrue AccountVectorizeIndexListResponseSuccess = true
)

type AccountVectorizeIndexDeleteResponse struct {
	Errors   []AccountVectorizeIndexDeleteResponseError   `json:"errors"`
	Messages []AccountVectorizeIndexDeleteResponseMessage `json:"messages"`
	Result   interface{}                                  `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountVectorizeIndexDeleteResponseSuccess `json:"success"`
	JSON    accountVectorizeIndexDeleteResponseJSON    `json:"-"`
}

// accountVectorizeIndexDeleteResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexDeleteResponse]
type accountVectorizeIndexDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexDeleteResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountVectorizeIndexDeleteResponseErrorJSON `json:"-"`
}

// accountVectorizeIndexDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexDeleteResponseError]
type accountVectorizeIndexDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexDeleteResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountVectorizeIndexDeleteResponseMessageJSON `json:"-"`
}

// accountVectorizeIndexDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountVectorizeIndexDeleteResponseMessage]
type accountVectorizeIndexDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountVectorizeIndexDeleteResponseSuccess bool

const (
	AccountVectorizeIndexDeleteResponseSuccessTrue AccountVectorizeIndexDeleteResponseSuccess = true
)

type AccountVectorizeIndexInsertResponse struct {
	Errors   []AccountVectorizeIndexInsertResponseError   `json:"errors"`
	Messages []AccountVectorizeIndexInsertResponseMessage `json:"messages"`
	Result   AccountVectorizeIndexInsertResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountVectorizeIndexInsertResponseSuccess `json:"success"`
	JSON    accountVectorizeIndexInsertResponseJSON    `json:"-"`
}

// accountVectorizeIndexInsertResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexInsertResponse]
type accountVectorizeIndexInsertResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexInsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexInsertResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountVectorizeIndexInsertResponseErrorJSON `json:"-"`
}

// accountVectorizeIndexInsertResponseErrorJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexInsertResponseError]
type accountVectorizeIndexInsertResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexInsertResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexInsertResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountVectorizeIndexInsertResponseMessageJSON `json:"-"`
}

// accountVectorizeIndexInsertResponseMessageJSON contains the JSON metadata for
// the struct [AccountVectorizeIndexInsertResponseMessage]
type accountVectorizeIndexInsertResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexInsertResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexInsertResponseResult struct {
	// Specifies the count of the vectors successfully inserted.
	Count int64 `json:"count"`
	// Array of vector identifiers of the vectors successfully inserted.
	IDs  []string                                      `json:"ids"`
	JSON accountVectorizeIndexInsertResponseResultJSON `json:"-"`
}

// accountVectorizeIndexInsertResponseResultJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexInsertResponseResult]
type accountVectorizeIndexInsertResponseResultJSON struct {
	Count       apijson.Field
	IDs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexInsertResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountVectorizeIndexInsertResponseSuccess bool

const (
	AccountVectorizeIndexInsertResponseSuccessTrue AccountVectorizeIndexInsertResponseSuccess = true
)

type AccountVectorizeIndexQueryResponse struct {
	Errors   []AccountVectorizeIndexQueryResponseError   `json:"errors"`
	Messages []AccountVectorizeIndexQueryResponseMessage `json:"messages"`
	Result   AccountVectorizeIndexQueryResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountVectorizeIndexQueryResponseSuccess `json:"success"`
	JSON    accountVectorizeIndexQueryResponseJSON    `json:"-"`
}

// accountVectorizeIndexQueryResponseJSON contains the JSON metadata for the struct
// [AccountVectorizeIndexQueryResponse]
type accountVectorizeIndexQueryResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexQueryResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountVectorizeIndexQueryResponseErrorJSON `json:"-"`
}

// accountVectorizeIndexQueryResponseErrorJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexQueryResponseError]
type accountVectorizeIndexQueryResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexQueryResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexQueryResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountVectorizeIndexQueryResponseMessageJSON `json:"-"`
}

// accountVectorizeIndexQueryResponseMessageJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexQueryResponseMessage]
type accountVectorizeIndexQueryResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexQueryResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexQueryResponseResult struct {
	// Specifies the count of vectors returned by the search
	Count int64 `json:"count"`
	// Array of vectors matched by the search
	Matches []AccountVectorizeIndexQueryResponseResultMatch `json:"matches"`
	JSON    accountVectorizeIndexQueryResponseResultJSON    `json:"-"`
}

// accountVectorizeIndexQueryResponseResultJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexQueryResponseResult]
type accountVectorizeIndexQueryResponseResultJSON struct {
	Count       apijson.Field
	Matches     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexQueryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexQueryResponseResultMatch struct {
	// The score of the vector according to the index's distance metric
	Score float64 `json:"score"`
	// If the returnVectors option is set, the vector itself
	Vector AccountVectorizeIndexQueryResponseResultMatchesVector `json:"vector,nullable"`
	// Identifier
	VectorID string                                            `json:"vectorId"`
	JSON     accountVectorizeIndexQueryResponseResultMatchJSON `json:"-"`
}

// accountVectorizeIndexQueryResponseResultMatchJSON contains the JSON metadata for
// the struct [AccountVectorizeIndexQueryResponseResultMatch]
type accountVectorizeIndexQueryResponseResultMatchJSON struct {
	Score       apijson.Field
	Vector      apijson.Field
	VectorID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexQueryResponseResultMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If the returnVectors option is set, the vector itself
type AccountVectorizeIndexQueryResponseResultMatchesVector struct {
	// Identifier
	ID       string                                                    `json:"id"`
	Metadata interface{}                                               `json:"metadata"`
	Values   []float64                                                 `json:"values"`
	JSON     accountVectorizeIndexQueryResponseResultMatchesVectorJSON `json:"-"`
}

// accountVectorizeIndexQueryResponseResultMatchesVectorJSON contains the JSON
// metadata for the struct [AccountVectorizeIndexQueryResponseResultMatchesVector]
type accountVectorizeIndexQueryResponseResultMatchesVectorJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexQueryResponseResultMatchesVector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountVectorizeIndexQueryResponseSuccess bool

const (
	AccountVectorizeIndexQueryResponseSuccessTrue AccountVectorizeIndexQueryResponseSuccess = true
)

type AccountVectorizeIndexUpsertResponse struct {
	Errors   []AccountVectorizeIndexUpsertResponseError   `json:"errors"`
	Messages []AccountVectorizeIndexUpsertResponseMessage `json:"messages"`
	Result   AccountVectorizeIndexUpsertResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountVectorizeIndexUpsertResponseSuccess `json:"success"`
	JSON    accountVectorizeIndexUpsertResponseJSON    `json:"-"`
}

// accountVectorizeIndexUpsertResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexUpsertResponse]
type accountVectorizeIndexUpsertResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexUpsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexUpsertResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountVectorizeIndexUpsertResponseErrorJSON `json:"-"`
}

// accountVectorizeIndexUpsertResponseErrorJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexUpsertResponseError]
type accountVectorizeIndexUpsertResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexUpsertResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexUpsertResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountVectorizeIndexUpsertResponseMessageJSON `json:"-"`
}

// accountVectorizeIndexUpsertResponseMessageJSON contains the JSON metadata for
// the struct [AccountVectorizeIndexUpsertResponseMessage]
type accountVectorizeIndexUpsertResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexUpsertResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVectorizeIndexUpsertResponseResult struct {
	// Specifies the count of the vectors successfully inserted.
	Count int64 `json:"count"`
	// Array of vector identifiers of the vectors successfully inserted.
	IDs  []string                                      `json:"ids"`
	JSON accountVectorizeIndexUpsertResponseResultJSON `json:"-"`
}

// accountVectorizeIndexUpsertResponseResultJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexUpsertResponseResult]
type accountVectorizeIndexUpsertResponseResultJSON struct {
	Count       apijson.Field
	IDs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexUpsertResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountVectorizeIndexUpsertResponseSuccess bool

const (
	AccountVectorizeIndexUpsertResponseSuccessTrue AccountVectorizeIndexUpsertResponseSuccess = true
)

type AccountVectorizeIndexNewParams struct {
	Config param.Field[interface{}] `json:"config,required"`
	Name   param.Field[string]      `json:"name,required"`
	// Specifies the description of the index.
	Description param.Field[string] `json:"description"`
}

func (r AccountVectorizeIndexNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountVectorizeIndexUpdateParams struct {
	// Specifies the description of the index.
	Description param.Field[string] `json:"description,required"`
}

func (r AccountVectorizeIndexUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountVectorizeIndexQueryParams struct {
	// Whether to return the values of the closest vectors, or just their identifiers.
	ReturnVectors param.Field[bool] `json:"returnVectors"`
	// The number of results to return
	TopK param.Field[float64] `json:"topK"`
	// The vector to find neighbors of
	Vector param.Field[[]float64] `json:"vector"`
}

func (r AccountVectorizeIndexQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
