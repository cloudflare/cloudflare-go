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

// APIGatewayUserSchemaService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAPIGatewayUserSchemaService]
// method instead.
type APIGatewayUserSchemaService struct {
	Options    []option.RequestOption
	Operations *APIGatewayUserSchemaOperationService
}

// NewAPIGatewayUserSchemaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIGatewayUserSchemaService(opts ...option.RequestOption) (r *APIGatewayUserSchemaService) {
	r = &APIGatewayUserSchemaService{}
	r.Options = opts
	r.Operations = NewAPIGatewayUserSchemaOperationService(opts...)
	return
}

// Upload a schema to a zone
func (r *APIGatewayUserSchemaService) New(ctx context.Context, zoneID string, body APIGatewayUserSchemaNewParams, opts ...option.RequestOption) (res *APIShieldSchemaUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewayUserSchemaNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable validation for a schema
func (r *APIGatewayUserSchemaService) Update(ctx context.Context, zoneID string, schemaID string, body APIGatewayUserSchemaUpdateParams, opts ...option.RequestOption) (res *APIShieldPublicSchema, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewayUserSchemaUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", zoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve information about all schemas on a zone
func (r *APIGatewayUserSchemaService) List(ctx context.Context, zoneID string, query APIGatewayUserSchemaListParams, opts ...option.RequestOption) (res *[]APIShieldPublicSchema, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewayUserSchemaListResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a schema
func (r *APIGatewayUserSchemaService) Delete(ctx context.Context, zoneID string, schemaID string, opts ...option.RequestOption) (res *APIGatewayUserSchemaDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIShieldAPIResponseSingle
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", zoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve information about a specific schema on a zone
func (r *APIGatewayUserSchemaService) Get(ctx context.Context, zoneID string, schemaID string, query APIGatewayUserSchemaGetParams, opts ...option.RequestOption) (res *APIShieldPublicSchema, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewayUserSchemaGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", zoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type APIShieldAPIResponseSingle struct {
	Errors   APIShieldMessages                `json:"errors,required"`
	Messages APIShieldMessages                `json:"messages,required"`
	Result   APIShieldAPIResponseSingleResult `json:"result,required"`
	// Whether the API call was successful
	Success bool                           `json:"success,required"`
	JSON    apiShieldAPIResponseSingleJSON `json:"-"`
}

// apiShieldAPIResponseSingleJSON contains the JSON metadata for the struct
// [APIShieldAPIResponseSingle]
type apiShieldAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [APIShieldAPIResponseSingleResultUnknown] or
// [shared.UnionString].
type APIShieldAPIResponseSingleResult interface {
	ImplementsAPIShieldAPIResponseSingleResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIShieldAPIResponseSingleResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIShieldMessages []APIShieldMessage

type APIShieldPublicSchema struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Kind of schema
	Kind APIShieldPublicSchemaKind `json:"kind,required"`
	// Name of the schema
	Name string `json:"name,required"`
	// UUID identifier
	SchemaID string `json:"schema_id,required" format:"uuid"`
	// Source of the schema
	Source string `json:"source"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled bool                      `json:"validation_enabled"`
	JSON              apiShieldPublicSchemaJSON `json:"-"`
}

// apiShieldPublicSchemaJSON contains the JSON metadata for the struct
// [APIShieldPublicSchema]
type apiShieldPublicSchemaJSON struct {
	CreatedAt         apijson.Field
	Kind              apijson.Field
	Name              apijson.Field
	SchemaID          apijson.Field
	Source            apijson.Field
	ValidationEnabled apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *APIShieldPublicSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Kind of schema
type APIShieldPublicSchemaKind string

const (
	APIShieldPublicSchemaKindOpenapiV3 APIShieldPublicSchemaKind = "openapi_v3"
)

type APIShieldSchemaUploadResponse struct {
	Schema        APIShieldPublicSchema                      `json:"schema,required"`
	UploadDetails APIShieldSchemaUploadResponseUploadDetails `json:"upload_details"`
	JSON          apiShieldSchemaUploadResponseJSON          `json:"-"`
}

// apiShieldSchemaUploadResponseJSON contains the JSON metadata for the struct
// [APIShieldSchemaUploadResponse]
type apiShieldSchemaUploadResponseJSON struct {
	Schema        apijson.Field
	UploadDetails apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIShieldSchemaUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIShieldSchemaUploadResponseUploadDetails struct {
	// Diagnostic warning events that occurred during processing. These events are
	// non-critical errors found within the schema.
	Warnings []APIShieldSchemaUploadResponseUploadDetailsWarning `json:"warnings"`
	JSON     apiShieldSchemaUploadResponseUploadDetailsJSON      `json:"-"`
}

// apiShieldSchemaUploadResponseUploadDetailsJSON contains the JSON metadata for
// the struct [APIShieldSchemaUploadResponseUploadDetails]
type apiShieldSchemaUploadResponseUploadDetailsJSON struct {
	Warnings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldSchemaUploadResponseUploadDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIShieldSchemaUploadResponseUploadDetailsWarning struct {
	// Code that identifies the event that occurred.
	Code int64 `json:"code,required"`
	// JSONPath location(s) in the schema where these events were encountered. See
	// [https://goessner.net/articles/JsonPath/](https://goessner.net/articles/JsonPath/)
	// for JSONPath specification.
	Locations []string `json:"locations"`
	// Diagnostic message that describes the event.
	Message string                                                `json:"message"`
	JSON    apiShieldSchemaUploadResponseUploadDetailsWarningJSON `json:"-"`
}

// apiShieldSchemaUploadResponseUploadDetailsWarningJSON contains the JSON metadata
// for the struct [APIShieldSchemaUploadResponseUploadDetailsWarning]
type apiShieldSchemaUploadResponseUploadDetailsWarningJSON struct {
	Code        apijson.Field
	Locations   apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIShieldSchemaUploadResponseUploadDetailsWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [APIGatewayUserSchemaDeleteResponseUnknown] or
// [shared.UnionString].
type APIGatewayUserSchemaDeleteResponse interface {
	ImplementsAPIGatewayUserSchemaDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIGatewayUserSchemaDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIGatewayUserSchemaNewParams struct {
	// Schema file bytes
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	// Kind of schema
	Kind param.Field[APIGatewayUserSchemaNewParamsKind] `json:"kind,required"`
	// Name of the schema
	Name param.Field[string] `json:"name"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[APIGatewayUserSchemaNewParamsValidationEnabled] `json:"validation_enabled"`
}

func (r APIGatewayUserSchemaNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type APIGatewayUserSchemaNewParamsKind string

const (
	APIGatewayUserSchemaNewParamsKindOpenapiV3 APIGatewayUserSchemaNewParamsKind = "openapi_v3"
)

// Flag whether schema is enabled for validation.
type APIGatewayUserSchemaNewParamsValidationEnabled string

const (
	APIGatewayUserSchemaNewParamsValidationEnabledTrue  APIGatewayUserSchemaNewParamsValidationEnabled = "true"
	APIGatewayUserSchemaNewParamsValidationEnabledFalse APIGatewayUserSchemaNewParamsValidationEnabled = "false"
)

type APIGatewayUserSchemaNewResponseEnvelope struct {
	Errors   APIShieldMessages             `json:"errors,required"`
	Messages APIShieldMessages             `json:"messages,required"`
	Result   APIShieldSchemaUploadResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                        `json:"success,required"`
	JSON    apiGatewayUserSchemaNewResponseEnvelopeJSON `json:"-"`
}

// apiGatewayUserSchemaNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [APIGatewayUserSchemaNewResponseEnvelope]
type apiGatewayUserSchemaNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayUserSchemaNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayUserSchemaUpdateParams struct {
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[APIGatewayUserSchemaUpdateParamsValidationEnabled] `json:"validation_enabled"`
}

func (r APIGatewayUserSchemaUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Flag whether schema is enabled for validation.
type APIGatewayUserSchemaUpdateParamsValidationEnabled bool

const (
	APIGatewayUserSchemaUpdateParamsValidationEnabledTrue APIGatewayUserSchemaUpdateParamsValidationEnabled = true
)

type APIGatewayUserSchemaUpdateResponseEnvelope struct {
	Errors   APIShieldMessages     `json:"errors,required"`
	Messages APIShieldMessages     `json:"messages,required"`
	Result   APIShieldPublicSchema `json:"result,required"`
	// Whether the API call was successful
	Success bool                                           `json:"success,required"`
	JSON    apiGatewayUserSchemaUpdateResponseEnvelopeJSON `json:"-"`
}

// apiGatewayUserSchemaUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [APIGatewayUserSchemaUpdateResponseEnvelope]
type apiGatewayUserSchemaUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayUserSchemaUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayUserSchemaListParams struct {
	// Omit the source-files of schemas and only retrieve their meta-data.
	OmitSource param.Field[bool] `query:"omit_source"`
	// Page number of paginated results.
	Page param.Field[interface{}] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[interface{}] `query:"per_page"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[bool] `query:"validation_enabled"`
}

// URLQuery serializes [APIGatewayUserSchemaListParams]'s query parameters as
// `url.Values`.
func (r APIGatewayUserSchemaListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIGatewayUserSchemaListResponseEnvelope struct {
	Errors   APIShieldMessages       `json:"errors,required"`
	Messages APIShieldMessages       `json:"messages,required"`
	Result   []APIShieldPublicSchema `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    bool                                               `json:"success,required"`
	ResultInfo APIGatewayUserSchemaListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       apiGatewayUserSchemaListResponseEnvelopeJSON       `json:"-"`
}

// apiGatewayUserSchemaListResponseEnvelopeJSON contains the JSON metadata for the
// struct [APIGatewayUserSchemaListResponseEnvelope]
type apiGatewayUserSchemaListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayUserSchemaListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayUserSchemaListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       apiGatewayUserSchemaListResponseEnvelopeResultInfoJSON `json:"-"`
}

// apiGatewayUserSchemaListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [APIGatewayUserSchemaListResponseEnvelopeResultInfo]
type apiGatewayUserSchemaListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayUserSchemaListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayUserSchemaGetParams struct {
	// Omit the source-files of schemas and only retrieve their meta-data.
	OmitSource param.Field[bool] `query:"omit_source"`
}

// URLQuery serializes [APIGatewayUserSchemaGetParams]'s query parameters as
// `url.Values`.
func (r APIGatewayUserSchemaGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIGatewayUserSchemaGetResponseEnvelope struct {
	Errors   APIShieldMessages     `json:"errors,required"`
	Messages APIShieldMessages     `json:"messages,required"`
	Result   APIShieldPublicSchema `json:"result,required"`
	// Whether the API call was successful
	Success bool                                        `json:"success,required"`
	JSON    apiGatewayUserSchemaGetResponseEnvelopeJSON `json:"-"`
}

// apiGatewayUserSchemaGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [APIGatewayUserSchemaGetResponseEnvelope]
type apiGatewayUserSchemaGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayUserSchemaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
