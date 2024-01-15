// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apiform"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZoneAPIGatewayUserSchemaService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneAPIGatewayUserSchemaService] method instead.
type ZoneAPIGatewayUserSchemaService struct {
	Options    []option.RequestOption
	Operations *ZoneAPIGatewayUserSchemaOperationService
}

// NewZoneAPIGatewayUserSchemaService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAPIGatewayUserSchemaService(opts ...option.RequestOption) (r *ZoneAPIGatewayUserSchemaService) {
	r = &ZoneAPIGatewayUserSchemaService{}
	r.Options = opts
	r.Operations = NewZoneAPIGatewayUserSchemaOperationService(opts...)
	return
}

// Upload a schema to a zone
func (r *ZoneAPIGatewayUserSchemaService) New(ctx context.Context, zoneID string, body ZoneAPIGatewayUserSchemaNewParams, opts ...option.RequestOption) (res *ZoneAPIGatewayUserSchemaNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve information about a specific schema on a zone
func (r *ZoneAPIGatewayUserSchemaService) Get(ctx context.Context, zoneID string, schemaID string, query ZoneAPIGatewayUserSchemaGetParams, opts ...option.RequestOption) (res *ZoneAPIGatewayUserSchemaGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", zoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Enable validation for a schema
func (r *ZoneAPIGatewayUserSchemaService) Update(ctx context.Context, zoneID string, schemaID string, body ZoneAPIGatewayUserSchemaUpdateParams, opts ...option.RequestOption) (res *ZoneAPIGatewayUserSchemaUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", zoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve information about all schemas on a zone
func (r *ZoneAPIGatewayUserSchemaService) List(ctx context.Context, zoneID string, query ZoneAPIGatewayUserSchemaListParams, opts ...option.RequestOption) (res *ZoneAPIGatewayUserSchemaListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a schema
func (r *ZoneAPIGatewayUserSchemaService) Delete(ctx context.Context, zoneID string, schemaID string, opts ...option.RequestOption) (res *ZoneAPIGatewayUserSchemaDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", zoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ZoneAPIGatewayUserSchemaNewResponse struct {
	Errors   []ZoneAPIGatewayUserSchemaNewResponseError   `json:"errors"`
	Messages []ZoneAPIGatewayUserSchemaNewResponseMessage `json:"messages"`
	Result   ZoneAPIGatewayUserSchemaNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                                    `json:"success"`
	JSON    zoneAPIGatewayUserSchemaNewResponseJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaNewResponseJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaNewResponse]
type zoneAPIGatewayUserSchemaNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaNewResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneAPIGatewayUserSchemaNewResponseErrorJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaNewResponseErrorJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaNewResponseError]
type zoneAPIGatewayUserSchemaNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaNewResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneAPIGatewayUserSchemaNewResponseMessageJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaNewResponseMessageJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayUserSchemaNewResponseMessage]
type zoneAPIGatewayUserSchemaNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaNewResponseResult struct {
	Schema        ZoneAPIGatewayUserSchemaNewResponseResultSchema        `json:"schema,required"`
	UploadDetails ZoneAPIGatewayUserSchemaNewResponseResultUploadDetails `json:"upload_details"`
	JSON          zoneAPIGatewayUserSchemaNewResponseResultJSON          `json:"-"`
}

// zoneAPIGatewayUserSchemaNewResponseResultJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaNewResponseResult]
type zoneAPIGatewayUserSchemaNewResponseResultJSON struct {
	Schema        apijson.Field
	UploadDetails apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaNewResponseResultSchema struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Kind of schema
	Kind ZoneAPIGatewayUserSchemaNewResponseResultSchemaKind `json:"kind,required"`
	// Name of the schema
	Name string `json:"name,required"`
	// UUID identifier
	SchemaID string `json:"schema_id,required" format:"uuid"`
	// Source of the schema
	Source string `json:"source"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled bool                                                `json:"validation_enabled"`
	JSON              zoneAPIGatewayUserSchemaNewResponseResultSchemaJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaNewResponseResultSchemaJSON contains the JSON metadata
// for the struct [ZoneAPIGatewayUserSchemaNewResponseResultSchema]
type zoneAPIGatewayUserSchemaNewResponseResultSchemaJSON struct {
	CreatedAt         apijson.Field
	Kind              apijson.Field
	Name              apijson.Field
	SchemaID          apijson.Field
	Source            apijson.Field
	ValidationEnabled apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaNewResponseResultSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Kind of schema
type ZoneAPIGatewayUserSchemaNewResponseResultSchemaKind string

const (
	ZoneAPIGatewayUserSchemaNewResponseResultSchemaKindOpenapiV3 ZoneAPIGatewayUserSchemaNewResponseResultSchemaKind = "openapi_v3"
)

type ZoneAPIGatewayUserSchemaNewResponseResultUploadDetails struct {
	// Diagnostic warning events that occurred during processing. These events are
	// non-critical errors found within the schema.
	Warnings []ZoneAPIGatewayUserSchemaNewResponseResultUploadDetailsWarning `json:"warnings"`
	JSON     zoneAPIGatewayUserSchemaNewResponseResultUploadDetailsJSON      `json:"-"`
}

// zoneAPIGatewayUserSchemaNewResponseResultUploadDetailsJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayUserSchemaNewResponseResultUploadDetails]
type zoneAPIGatewayUserSchemaNewResponseResultUploadDetailsJSON struct {
	Warnings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaNewResponseResultUploadDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaNewResponseResultUploadDetailsWarning struct {
	// Code that identifies the event that occurred.
	Code int64 `json:"code,required"`
	// JSONPath location(s) in the schema where these events were encountered. See
	// [https://goessner.net/articles/JsonPath/](https://goessner.net/articles/JsonPath/)
	// for JSONPath specification.
	Locations []string `json:"locations"`
	// Diagnostic message that describes the event.
	Message string                                                            `json:"message"`
	JSON    zoneAPIGatewayUserSchemaNewResponseResultUploadDetailsWarningJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaNewResponseResultUploadDetailsWarningJSON contains the
// JSON metadata for the struct
// [ZoneAPIGatewayUserSchemaNewResponseResultUploadDetailsWarning]
type zoneAPIGatewayUserSchemaNewResponseResultUploadDetailsWarningJSON struct {
	Code        apijson.Field
	Locations   apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaNewResponseResultUploadDetailsWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaGetResponse struct {
	Errors   []ZoneAPIGatewayUserSchemaGetResponseError   `json:"errors"`
	Messages []ZoneAPIGatewayUserSchemaGetResponseMessage `json:"messages"`
	Result   ZoneAPIGatewayUserSchemaGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                                    `json:"success"`
	JSON    zoneAPIGatewayUserSchemaGetResponseJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaGetResponseJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaGetResponse]
type zoneAPIGatewayUserSchemaGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaGetResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneAPIGatewayUserSchemaGetResponseErrorJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaGetResponseErrorJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaGetResponseError]
type zoneAPIGatewayUserSchemaGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaGetResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneAPIGatewayUserSchemaGetResponseMessageJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaGetResponseMessageJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayUserSchemaGetResponseMessage]
type zoneAPIGatewayUserSchemaGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaGetResponseResult struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Kind of schema
	Kind ZoneAPIGatewayUserSchemaGetResponseResultKind `json:"kind,required"`
	// Name of the schema
	Name string `json:"name,required"`
	// UUID identifier
	SchemaID string `json:"schema_id,required" format:"uuid"`
	// Source of the schema
	Source string `json:"source"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled bool                                          `json:"validation_enabled"`
	JSON              zoneAPIGatewayUserSchemaGetResponseResultJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaGetResponseResult]
type zoneAPIGatewayUserSchemaGetResponseResultJSON struct {
	CreatedAt         apijson.Field
	Kind              apijson.Field
	Name              apijson.Field
	SchemaID          apijson.Field
	Source            apijson.Field
	ValidationEnabled apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Kind of schema
type ZoneAPIGatewayUserSchemaGetResponseResultKind string

const (
	ZoneAPIGatewayUserSchemaGetResponseResultKindOpenapiV3 ZoneAPIGatewayUserSchemaGetResponseResultKind = "openapi_v3"
)

type ZoneAPIGatewayUserSchemaUpdateResponse struct {
	Errors   []ZoneAPIGatewayUserSchemaUpdateResponseError   `json:"errors"`
	Messages []ZoneAPIGatewayUserSchemaUpdateResponseMessage `json:"messages"`
	Result   ZoneAPIGatewayUserSchemaUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                                       `json:"success"`
	JSON    zoneAPIGatewayUserSchemaUpdateResponseJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaUpdateResponse]
type zoneAPIGatewayUserSchemaUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaUpdateResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneAPIGatewayUserSchemaUpdateResponseErrorJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayUserSchemaUpdateResponseError]
type zoneAPIGatewayUserSchemaUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaUpdateResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneAPIGatewayUserSchemaUpdateResponseMessageJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayUserSchemaUpdateResponseMessage]
type zoneAPIGatewayUserSchemaUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaUpdateResponseResult struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Kind of schema
	Kind ZoneAPIGatewayUserSchemaUpdateResponseResultKind `json:"kind,required"`
	// Name of the schema
	Name string `json:"name,required"`
	// UUID identifier
	SchemaID string `json:"schema_id,required" format:"uuid"`
	// Source of the schema
	Source string `json:"source"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled bool                                             `json:"validation_enabled"`
	JSON              zoneAPIGatewayUserSchemaUpdateResponseResultJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayUserSchemaUpdateResponseResult]
type zoneAPIGatewayUserSchemaUpdateResponseResultJSON struct {
	CreatedAt         apijson.Field
	Kind              apijson.Field
	Name              apijson.Field
	SchemaID          apijson.Field
	Source            apijson.Field
	ValidationEnabled apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Kind of schema
type ZoneAPIGatewayUserSchemaUpdateResponseResultKind string

const (
	ZoneAPIGatewayUserSchemaUpdateResponseResultKindOpenapiV3 ZoneAPIGatewayUserSchemaUpdateResponseResultKind = "openapi_v3"
)

type ZoneAPIGatewayUserSchemaListResponse struct {
	Errors     []ZoneAPIGatewayUserSchemaListResponseError    `json:"errors"`
	Messages   []ZoneAPIGatewayUserSchemaListResponseMessage  `json:"messages"`
	Result     []ZoneAPIGatewayUserSchemaListResponseResult   `json:"result"`
	ResultInfo ZoneAPIGatewayUserSchemaListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success bool                                     `json:"success"`
	JSON    zoneAPIGatewayUserSchemaListResponseJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaListResponseJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaListResponse]
type zoneAPIGatewayUserSchemaListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaListResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneAPIGatewayUserSchemaListResponseErrorJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaListResponseError]
type zoneAPIGatewayUserSchemaListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaListResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneAPIGatewayUserSchemaListResponseMessageJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayUserSchemaListResponseMessage]
type zoneAPIGatewayUserSchemaListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaListResponseResult struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Kind of schema
	Kind ZoneAPIGatewayUserSchemaListResponseResultKind `json:"kind,required"`
	// Name of the schema
	Name string `json:"name,required"`
	// UUID identifier
	SchemaID string `json:"schema_id,required" format:"uuid"`
	// Source of the schema
	Source string `json:"source"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled bool                                           `json:"validation_enabled"`
	JSON              zoneAPIGatewayUserSchemaListResponseResultJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaListResponseResultJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayUserSchemaListResponseResult]
type zoneAPIGatewayUserSchemaListResponseResultJSON struct {
	CreatedAt         apijson.Field
	Kind              apijson.Field
	Name              apijson.Field
	SchemaID          apijson.Field
	Source            apijson.Field
	ValidationEnabled apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Kind of schema
type ZoneAPIGatewayUserSchemaListResponseResultKind string

const (
	ZoneAPIGatewayUserSchemaListResponseResultKindOpenapiV3 ZoneAPIGatewayUserSchemaListResponseResultKind = "openapi_v3"
)

type ZoneAPIGatewayUserSchemaListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       zoneAPIGatewayUserSchemaListResponseResultInfoJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaListResponseResultInfoJSON contains the JSON metadata
// for the struct [ZoneAPIGatewayUserSchemaListResponseResultInfo]
type zoneAPIGatewayUserSchemaListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaDeleteResponse struct {
	Errors   []ZoneAPIGatewayUserSchemaDeleteResponseError   `json:"errors"`
	Messages []ZoneAPIGatewayUserSchemaDeleteResponseMessage `json:"messages"`
	Result   ZoneAPIGatewayUserSchemaDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                                       `json:"success"`
	JSON    zoneAPIGatewayUserSchemaDeleteResponseJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaDeleteResponse]
type zoneAPIGatewayUserSchemaDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaDeleteResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneAPIGatewayUserSchemaDeleteResponseErrorJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaDeleteResponseErrorJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayUserSchemaDeleteResponseError]
type zoneAPIGatewayUserSchemaDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayUserSchemaDeleteResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneAPIGatewayUserSchemaDeleteResponseMessageJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaDeleteResponseMessageJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayUserSchemaDeleteResponseMessage]
type zoneAPIGatewayUserSchemaDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZoneAPIGatewayUserSchemaDeleteResponseResultUnknown] or
// [shared.UnionString].
type ZoneAPIGatewayUserSchemaDeleteResponseResult interface {
	ImplementsZoneAPIGatewayUserSchemaDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAPIGatewayUserSchemaDeleteResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneAPIGatewayUserSchemaNewParams struct {
	// Schema file bytes
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	// Kind of schema
	Kind param.Field[ZoneAPIGatewayUserSchemaNewParamsKind] `json:"kind,required"`
	// Name of the schema
	Name param.Field[string] `json:"name"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[ZoneAPIGatewayUserSchemaNewParamsValidationEnabled] `json:"validation_enabled"`
}

func (r ZoneAPIGatewayUserSchemaNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Kind of schema
type ZoneAPIGatewayUserSchemaNewParamsKind string

const (
	ZoneAPIGatewayUserSchemaNewParamsKindOpenapiV3 ZoneAPIGatewayUserSchemaNewParamsKind = "openapi_v3"
)

// Flag whether schema is enabled for validation.
type ZoneAPIGatewayUserSchemaNewParamsValidationEnabled string

const (
	ZoneAPIGatewayUserSchemaNewParamsValidationEnabledTrue  ZoneAPIGatewayUserSchemaNewParamsValidationEnabled = "true"
	ZoneAPIGatewayUserSchemaNewParamsValidationEnabledFalse ZoneAPIGatewayUserSchemaNewParamsValidationEnabled = "false"
)

type ZoneAPIGatewayUserSchemaGetParams struct {
	// Omit the source-files of schemas and only retrieve their meta-data.
	OmitSource param.Field[bool] `query:"omit_source"`
}

// URLQuery serializes [ZoneAPIGatewayUserSchemaGetParams]'s query parameters as
// `url.Values`.
func (r ZoneAPIGatewayUserSchemaGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneAPIGatewayUserSchemaUpdateParams struct {
	ValidationEnabled param.Field[interface{}] `json:"validation_enabled"`
}

func (r ZoneAPIGatewayUserSchemaUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAPIGatewayUserSchemaListParams struct {
	// Omit the source-files of schemas and only retrieve their meta-data.
	OmitSource param.Field[bool] `query:"omit_source"`
	// Page number of paginated results.
	Page param.Field[interface{}] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[interface{}] `query:"per_page"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[bool] `query:"validation_enabled"`
}

// URLQuery serializes [ZoneAPIGatewayUserSchemaListParams]'s query parameters as
// `url.Values`.
func (r ZoneAPIGatewayUserSchemaListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
