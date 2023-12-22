// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountWorkerDurableObjectNamespaceObjectService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerDurableObjectNamespaceObjectService] method instead.
type AccountWorkerDurableObjectNamespaceObjectService struct {
	Options []option.RequestOption
}

// NewAccountWorkerDurableObjectNamespaceObjectService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerDurableObjectNamespaceObjectService(opts ...option.RequestOption) (r *AccountWorkerDurableObjectNamespaceObjectService) {
	r = &AccountWorkerDurableObjectNamespaceObjectService{}
	r.Options = opts
	return
}

// Returns the Durable Objects in a given namespace.
func (r *AccountWorkerDurableObjectNamespaceObjectService) List(ctx context.Context, accountIdentifier string, id string, query AccountWorkerDurableObjectNamespaceObjectListParams, opts ...option.RequestOption) (res *AccountWorkerDurableObjectNamespaceObjectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/durable_objects/namespaces/%s/objects", accountIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountWorkerDurableObjectNamespaceObjectListResponse struct {
	Errors     []AccountWorkerDurableObjectNamespaceObjectListResponseError    `json:"errors"`
	Messages   []AccountWorkerDurableObjectNamespaceObjectListResponseMessage  `json:"messages"`
	Result     []AccountWorkerDurableObjectNamespaceObjectListResponseResult   `json:"result"`
	ResultInfo AccountWorkerDurableObjectNamespaceObjectListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountWorkerDurableObjectNamespaceObjectListResponseSuccess `json:"success"`
	JSON    accountWorkerDurableObjectNamespaceObjectListResponseJSON    `json:"-"`
}

// accountWorkerDurableObjectNamespaceObjectListResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDurableObjectNamespaceObjectListResponse]
type accountWorkerDurableObjectNamespaceObjectListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceObjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDurableObjectNamespaceObjectListResponseError struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    accountWorkerDurableObjectNamespaceObjectListResponseErrorJSON `json:"-"`
}

// accountWorkerDurableObjectNamespaceObjectListResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountWorkerDurableObjectNamespaceObjectListResponseError]
type accountWorkerDurableObjectNamespaceObjectListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceObjectListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDurableObjectNamespaceObjectListResponseMessage struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    accountWorkerDurableObjectNamespaceObjectListResponseMessageJSON `json:"-"`
}

// accountWorkerDurableObjectNamespaceObjectListResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountWorkerDurableObjectNamespaceObjectListResponseMessage]
type accountWorkerDurableObjectNamespaceObjectListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceObjectListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDurableObjectNamespaceObjectListResponseResult struct {
	// ID of the Durable Object.
	ID string `json:"id"`
	// Whether the Durable Object has stored data.
	HasStoredData bool                                                            `json:"hasStoredData"`
	JSON          accountWorkerDurableObjectNamespaceObjectListResponseResultJSON `json:"-"`
}

// accountWorkerDurableObjectNamespaceObjectListResponseResultJSON contains the
// JSON metadata for the struct
// [AccountWorkerDurableObjectNamespaceObjectListResponseResult]
type accountWorkerDurableObjectNamespaceObjectListResponseResultJSON struct {
	ID            apijson.Field
	HasStoredData apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceObjectListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDurableObjectNamespaceObjectListResponseResultInfo struct {
	// Total results returned based on your list parameters.
	Count float64 `json:"count"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records. A valid value for the cursor can be obtained from the
	// cursors object in the result_info structure.
	Cursor string                                                              `json:"cursor"`
	JSON   accountWorkerDurableObjectNamespaceObjectListResponseResultInfoJSON `json:"-"`
}

// accountWorkerDurableObjectNamespaceObjectListResponseResultInfoJSON contains the
// JSON metadata for the struct
// [AccountWorkerDurableObjectNamespaceObjectListResponseResultInfo]
type accountWorkerDurableObjectNamespaceObjectListResponseResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceObjectListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerDurableObjectNamespaceObjectListResponseSuccess bool

const (
	AccountWorkerDurableObjectNamespaceObjectListResponseSuccessTrue AccountWorkerDurableObjectNamespaceObjectListResponseSuccess = true
)

type AccountWorkerDurableObjectNamespaceObjectListParams struct {
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records. A valid value for the cursor can be obtained from the
	// cursors object in the result_info structure.
	Cursor param.Field[string] `query:"cursor"`
	// The number of objects to return. The cursor attribute may be used to iterate
	// over the next batch of objects if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
}

// URLQuery serializes [AccountWorkerDurableObjectNamespaceObjectListParams]'s
// query parameters as `url.Values`.
func (r AccountWorkerDurableObjectNamespaceObjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
