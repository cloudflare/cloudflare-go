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

// NamespaceTableService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceTableService] method instead.
type NamespaceTableService struct {
	Options            []option.RequestOption
	MaintenanceConfigs *NamespaceTableMaintenanceConfigService
}

// NewNamespaceTableService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceTableService(opts ...option.RequestOption) (r *NamespaceTableService) {
	r = &NamespaceTableService{}
	r.Options = opts
	r.MaintenanceConfigs = NewNamespaceTableMaintenanceConfigService(opts...)
	return
}

// Returns a list of tables in the specified namespace within an R2 catalog.
// Supports pagination for efficient traversal of large table collections.
func (r *NamespaceTableService) List(ctx context.Context, bucketName string, namespace string, params NamespaceTableListParams, opts ...option.RequestOption) (res *NamespaceTableListResponse, err error) {
	var env NamespaceTableListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	if namespace == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2-catalog/%s/namespaces/%s/tables", params.AccountID, bucketName, namespace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Contains the list of tables with optional pagination.
type NamespaceTableListResponse struct {
	// Lists tables in the namespace.
	Identifiers []NamespaceTableListResponseIdentifier `json:"identifiers,required"`
	// Contains detailed metadata for each table when return_details is true. Each
	// object includes identifier, UUID, timestamps, and locations.
	Details []NamespaceTableListResponseDetail `json:"details,nullable"`
	// Use this opaque token to fetch the next page of results. A null or absent value
	// indicates the last page.
	NextPageToken string `json:"next_page_token,nullable"`
	// Contains UUIDs for each table when return_uuids is true. The order corresponds
	// to the identifiers array.
	TableUUIDs []string                       `json:"table_uuids,nullable" format:"uuid"`
	JSON       namespaceTableListResponseJSON `json:"-"`
}

// namespaceTableListResponseJSON contains the JSON metadata for the struct
// [NamespaceTableListResponse]
type namespaceTableListResponseJSON struct {
	Identifiers   apijson.Field
	Details       apijson.Field
	NextPageToken apijson.Field
	TableUUIDs    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *NamespaceTableListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableListResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies a unique table identifier within a catalog.
type NamespaceTableListResponseIdentifier struct {
	// Specifies the table name.
	Name string `json:"name,required"`
	// Specifies the hierarchical namespace parts as an array of strings. For example,
	// ["bronze", "analytics"] represents the namespace "bronze.analytics".
	Namespace []string                                 `json:"namespace,required"`
	JSON      namespaceTableListResponseIdentifierJSON `json:"-"`
}

// namespaceTableListResponseIdentifierJSON contains the JSON metadata for the
// struct [NamespaceTableListResponseIdentifier]
type namespaceTableListResponseIdentifierJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceTableListResponseIdentifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableListResponseIdentifierJSON) RawJSON() string {
	return r.raw
}

// Contains table with metadata.
type NamespaceTableListResponseDetail struct {
	// Specifies a unique table identifier within a catalog.
	Identifier NamespaceTableListResponseDetailsIdentifier `json:"identifier,required"`
	// Contains the UUID that persists across renames.
	TableUUID string `json:"table_uuid,required" format:"uuid"`
	// Indicates the creation timestamp in ISO 8601 format.
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// Specifies the base S3 URI for table storage location.
	Location string `json:"location,nullable"`
	// Contains the S3 URI to table metadata file. Null for staged tables.
	MetadataLocation string `json:"metadata_location,nullable"`
	// Shows the last update timestamp in ISO 8601 format. Null if never updated.
	UpdatedAt time.Time                            `json:"updated_at,nullable" format:"date-time"`
	JSON      namespaceTableListResponseDetailJSON `json:"-"`
}

// namespaceTableListResponseDetailJSON contains the JSON metadata for the struct
// [NamespaceTableListResponseDetail]
type namespaceTableListResponseDetailJSON struct {
	Identifier       apijson.Field
	TableUUID        apijson.Field
	CreatedAt        apijson.Field
	Location         apijson.Field
	MetadataLocation apijson.Field
	UpdatedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NamespaceTableListResponseDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableListResponseDetailJSON) RawJSON() string {
	return r.raw
}

// Specifies a unique table identifier within a catalog.
type NamespaceTableListResponseDetailsIdentifier struct {
	// Specifies the table name.
	Name string `json:"name,required"`
	// Specifies the hierarchical namespace parts as an array of strings. For example,
	// ["bronze", "analytics"] represents the namespace "bronze.analytics".
	Namespace []string                                        `json:"namespace,required"`
	JSON      namespaceTableListResponseDetailsIdentifierJSON `json:"-"`
}

// namespaceTableListResponseDetailsIdentifierJSON contains the JSON metadata for
// the struct [NamespaceTableListResponseDetailsIdentifier]
type namespaceTableListResponseDetailsIdentifierJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceTableListResponseDetailsIdentifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableListResponseDetailsIdentifierJSON) RawJSON() string {
	return r.raw
}

type NamespaceTableListParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Maximum number of tables to return per page. Defaults to 100, maximum 1000.
	PageSize param.Field[int64] `query:"page_size"`
	// Opaque pagination token from a previous response. Use this to fetch the next
	// page of results.
	PageToken param.Field[string] `query:"page_token"`
	// Whether to include additional metadata (timestamps, locations). When true,
	// response includes created_at, updated_at, metadata_locations, and locations
	// arrays.
	ReturnDetails param.Field[bool] `query:"return_details"`
	// Whether to include table UUIDs in the response. Set to true to receive the
	// table_uuids array.
	ReturnUUIDs param.Field[bool] `query:"return_uuids"`
}

// URLQuery serializes [NamespaceTableListParams]'s query parameters as
// `url.Values`.
func (r NamespaceTableListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NamespaceTableListResponseEnvelope struct {
	// Contains errors if the API call was unsuccessful.
	Errors []NamespaceTableListResponseEnvelopeErrors `json:"errors,required"`
	// Contains informational messages.
	Messages []NamespaceTableListResponseEnvelopeMessages `json:"messages,required"`
	// Indicates whether the API call was successful.
	Success bool `json:"success,required"`
	// Contains the list of tables with optional pagination.
	Result NamespaceTableListResponse             `json:"result"`
	JSON   namespaceTableListResponseEnvelopeJSON `json:"-"`
}

// namespaceTableListResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceTableListResponseEnvelope]
type namespaceTableListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceTableListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceTableListResponseEnvelopeErrors struct {
	// Specifies the error code.
	Code int64 `json:"code,required"`
	// Describes the error.
	Message string                                       `json:"message,required"`
	JSON    namespaceTableListResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceTableListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [NamespaceTableListResponseEnvelopeErrors]
type namespaceTableListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceTableListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceTableListResponseEnvelopeMessages struct {
	// Specifies the message code.
	Code int64 `json:"code,required"`
	// Contains the message text.
	Message string                                         `json:"message,required"`
	JSON    namespaceTableListResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceTableListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [NamespaceTableListResponseEnvelopeMessages]
type namespaceTableListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceTableListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
