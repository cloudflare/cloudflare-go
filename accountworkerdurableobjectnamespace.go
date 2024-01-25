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

// AccountWorkerDurableObjectNamespaceService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountWorkerDurableObjectNamespaceService] method instead.
type AccountWorkerDurableObjectNamespaceService struct {
	Options []option.RequestOption
	Objects *AccountWorkerDurableObjectNamespaceObjectService
}

// NewAccountWorkerDurableObjectNamespaceService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerDurableObjectNamespaceService(opts ...option.RequestOption) (r *AccountWorkerDurableObjectNamespaceService) {
	r = &AccountWorkerDurableObjectNamespaceService{}
	r.Options = opts
	r.Objects = NewAccountWorkerDurableObjectNamespaceObjectService(opts...)
	return
}

// Returns the Durable Object namespaces owned by an account.
func (r *AccountWorkerDurableObjectNamespaceService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountWorkerDurableObjectNamespaceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/durable_objects/namespaces", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountWorkerDurableObjectNamespaceListResponse struct {
	Errors     []AccountWorkerDurableObjectNamespaceListResponseError    `json:"errors"`
	Messages   []AccountWorkerDurableObjectNamespaceListResponseMessage  `json:"messages"`
	Result     []AccountWorkerDurableObjectNamespaceListResponseResult   `json:"result"`
	ResultInfo AccountWorkerDurableObjectNamespaceListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountWorkerDurableObjectNamespaceListResponseSuccess `json:"success"`
	JSON    accountWorkerDurableObjectNamespaceListResponseJSON    `json:"-"`
}

// accountWorkerDurableObjectNamespaceListResponseJSON contains the JSON metadata
// for the struct [AccountWorkerDurableObjectNamespaceListResponse]
type accountWorkerDurableObjectNamespaceListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDurableObjectNamespaceListResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountWorkerDurableObjectNamespaceListResponseErrorJSON `json:"-"`
}

// accountWorkerDurableObjectNamespaceListResponseErrorJSON contains the JSON
// metadata for the struct [AccountWorkerDurableObjectNamespaceListResponseError]
type accountWorkerDurableObjectNamespaceListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDurableObjectNamespaceListResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountWorkerDurableObjectNamespaceListResponseMessageJSON `json:"-"`
}

// accountWorkerDurableObjectNamespaceListResponseMessageJSON contains the JSON
// metadata for the struct [AccountWorkerDurableObjectNamespaceListResponseMessage]
type accountWorkerDurableObjectNamespaceListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDurableObjectNamespaceListResponseResult struct {
	ID     interface{}                                               `json:"id"`
	Class  interface{}                                               `json:"class"`
	Name   interface{}                                               `json:"name"`
	Script interface{}                                               `json:"script"`
	JSON   accountWorkerDurableObjectNamespaceListResponseResultJSON `json:"-"`
}

// accountWorkerDurableObjectNamespaceListResponseResultJSON contains the JSON
// metadata for the struct [AccountWorkerDurableObjectNamespaceListResponseResult]
type accountWorkerDurableObjectNamespaceListResponseResultJSON struct {
	ID          apijson.Field
	Class       apijson.Field
	Name        apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDurableObjectNamespaceListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       accountWorkerDurableObjectNamespaceListResponseResultInfoJSON `json:"-"`
}

// accountWorkerDurableObjectNamespaceListResponseResultInfoJSON contains the JSON
// metadata for the struct
// [AccountWorkerDurableObjectNamespaceListResponseResultInfo]
type accountWorkerDurableObjectNamespaceListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerDurableObjectNamespaceListResponseSuccess bool

const (
	AccountWorkerDurableObjectNamespaceListResponseSuccessTrue AccountWorkerDurableObjectNamespaceListResponseSuccess = true
)
