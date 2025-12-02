// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_data_catalog

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// NamespaceService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceService] method instead.
type NamespaceService struct {
	Options []option.RequestOption
	Tables  *NamespaceTableService
}

// NewNamespaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceService(opts ...option.RequestOption) (r *NamespaceService) {
	r = &NamespaceService{}
	r.Options = opts
	r.Tables = NewNamespaceTableService(opts...)
	return
}

// Returns a list of namespaces in the specified R2 catalog. Supports hierarchical
// filtering and pagination for efficient traversal of large namespace hierarchies.
func (r *NamespaceService) List(ctx context.Context, bucketName string, params NamespaceListParams, opts ...option.RequestOption) (res *NamespaceListResponse, err error) {
	var env NamespaceListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2-catalog/%s/namespaces", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Contains the list of namespaces with optional pagination.
type NamespaceListResponse struct {
	// Lists namespaces in the catalog.
	Namespaces [][]string `json:"namespaces,required"`
	// Contains detailed metadata for each namespace when return_details is true. Each
	// object includes the namespace, UUID, and timestamps.
	Details []NamespaceListResponseDetail `json:"details,nullable"`
	// Contains UUIDs for each namespace when return_uuids is true. The order
	// corresponds to the namespaces array.
	NamespaceUUIDs []string `json:"namespace_uuids,nullable" format:"uuid"`
	// Use this opaque token to fetch the next page of results. A null or absent value
	// indicates the last page.
	NextPageToken string                    `json:"next_page_token,nullable"`
	JSON          namespaceListResponseJSON `json:"-"`
}

// namespaceListResponseJSON contains the JSON metadata for the struct
// [NamespaceListResponse]
type namespaceListResponseJSON struct {
	Namespaces     apijson.Field
	Details        apijson.Field
	NamespaceUUIDs apijson.Field
	NextPageToken  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponseJSON) RawJSON() string {
	return r.raw
}

// Contains namespace with metadata details.
type NamespaceListResponseDetail struct {
	// Specifies the hierarchical namespace parts as an array of strings. For example,
	// ["bronze", "analytics"] represents the namespace "bronze.analytics".
	Namespace []string `json:"namespace,required"`
	// Contains the UUID that persists across renames.
	NamespaceUUID string `json:"namespace_uuid,required" format:"uuid"`
	// Indicates the creation timestamp in ISO 8601 format.
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// Shows the last update timestamp in ISO 8601 format. Null if never updated.
	UpdatedAt time.Time                       `json:"updated_at,nullable" format:"date-time"`
	JSON      namespaceListResponseDetailJSON `json:"-"`
}

// namespaceListResponseDetailJSON contains the JSON metadata for the struct
// [NamespaceListResponseDetail]
type namespaceListResponseDetailJSON struct {
	Namespace     apijson.Field
	NamespaceUUID apijson.Field
	CreatedAt     apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *NamespaceListResponseDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponseDetailJSON) RawJSON() string {
	return r.raw
}

type NamespaceListParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Maximum number of namespaces to return per page. Defaults to 100, maximum 1000.
	PageSize param.Field[int64] `query:"page_size"`
	// Opaque pagination token from a previous response. Use this to fetch the next
	// page of results.
	PageToken param.Field[string] `query:"page_token"`
	// Parent namespace to filter by. Only returns direct children of this namespace.
	// For nested namespaces, use %1F as separator (e.g., "bronze%1Fanalytics"). Omit
	// this parameter to list top-level namespaces.
	Parent param.Field[string] `query:"parent"`
	// Whether to include additional metadata (timestamps). When true, response
	// includes created_at and updated_at arrays.
	ReturnDetails param.Field[bool] `query:"return_details"`
	// Whether to include namespace UUIDs in the response. Set to true to receive the
	// namespace_uuids array.
	ReturnUUIDs param.Field[bool] `query:"return_uuids"`
}

// URLQuery serializes [NamespaceListParams]'s query parameters as `url.Values`.
func (r NamespaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NamespaceListResponseEnvelope struct {
	// Contains errors if the API call was unsuccessful.
	Errors []NamespaceListResponseEnvelopeErrors `json:"errors,required"`
	// Contains informational messages.
	Messages []NamespaceListResponseEnvelopeMessages `json:"messages,required"`
	// Indicates whether the API call was successful.
	Success bool `json:"success,required"`
	// Contains the list of namespaces with optional pagination.
	Result NamespaceListResponse             `json:"result"`
	JSON   namespaceListResponseEnvelopeJSON `json:"-"`
}

// namespaceListResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceListResponseEnvelope]
type namespaceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceListResponseEnvelopeErrors struct {
	// Specifies the error code.
	Code int64 `json:"code,required"`
	// Describes the error.
	Message string                                  `json:"message,required"`
	JSON    namespaceListResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [NamespaceListResponseEnvelopeErrors]
type namespaceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceListResponseEnvelopeMessages struct {
	// Specifies the message code.
	Code int64 `json:"code,required"`
	// Contains the message text.
	Message string                                    `json:"message,required"`
	JSON    namespaceListResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [NamespaceListResponseEnvelopeMessages]
type namespaceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
