// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// UserSchemaService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserSchemaService] method instead.
type UserSchemaService struct {
	Options    []option.RequestOption
	Operations *UserSchemaOperationService
}

// NewUserSchemaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserSchemaService(opts ...option.RequestOption) (r *UserSchemaService) {
	r = &UserSchemaService{}
	r.Options = opts
	r.Operations = NewUserSchemaOperationService(opts...)
	return
}

// Upload a schema to a zone
func (r *UserSchemaService) New(ctx context.Context, params UserSchemaNewParams, opts ...option.RequestOption) (res *SchemaUpload, err error) {
	opts = append(r.Options[:], opts...)
	var env UserSchemaNewResponseEnvelope
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve information about all schemas on a zone
func (r *UserSchemaService) List(ctx context.Context, params UserSchemaListParams, opts ...option.RequestOption) (res *pagination.SinglePage[PublicSchema], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve information about all schemas on a zone
func (r *UserSchemaService) ListAutoPaging(ctx context.Context, params UserSchemaListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PublicSchema] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Delete a schema
func (r *UserSchemaService) Delete(ctx context.Context, schemaID string, body UserSchemaDeleteParams, opts ...option.RequestOption) (res *UserSchemaDeleteResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env UserSchemaDeleteResponseEnvelope
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", body.ZoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable validation for a schema
func (r *UserSchemaService) Edit(ctx context.Context, schemaID string, params UserSchemaEditParams, opts ...option.RequestOption) (res *PublicSchema, err error) {
	opts = append(r.Options[:], opts...)
	var env UserSchemaEditResponseEnvelope
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", params.ZoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve information about a specific schema on a zone
func (r *UserSchemaService) Get(ctx context.Context, schemaID string, params UserSchemaGetParams, opts ...option.RequestOption) (res *PublicSchema, err error) {
	opts = append(r.Options[:], opts...)
	var env UserSchemaGetResponseEnvelope
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", params.ZoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Message []shared.ResponseInfo

type PublicSchema struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Kind of schema
	Kind PublicSchemaKind `json:"kind,required"`
	// Name of the schema
	Name string `json:"name,required"`
	// UUID identifier
	SchemaID string `json:"schema_id,required" format:"uuid"`
	// Source of the schema
	Source string `json:"source"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled bool             `json:"validation_enabled"`
	JSON              publicSchemaJSON `json:"-"`
}

// publicSchemaJSON contains the JSON metadata for the struct [PublicSchema]
type publicSchemaJSON struct {
	CreatedAt         apijson.Field
	Kind              apijson.Field
	Name              apijson.Field
	SchemaID          apijson.Field
	Source            apijson.Field
	ValidationEnabled apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PublicSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r publicSchemaJSON) RawJSON() string {
	return r.raw
}

// Kind of schema
type PublicSchemaKind string

const (
	PublicSchemaKindOpenAPIV3 PublicSchemaKind = "openapi_v3"
)

func (r PublicSchemaKind) IsKnown() bool {
	switch r {
	case PublicSchemaKindOpenAPIV3:
		return true
	}
	return false
}

type SchemaUpload struct {
	Schema        PublicSchema              `json:"schema,required"`
	UploadDetails SchemaUploadUploadDetails `json:"upload_details"`
	JSON          schemaUploadJSON          `json:"-"`
}

// schemaUploadJSON contains the JSON metadata for the struct [SchemaUpload]
type schemaUploadJSON struct {
	Schema        apijson.Field
	UploadDetails apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SchemaUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaUploadJSON) RawJSON() string {
	return r.raw
}

type SchemaUploadUploadDetails struct {
	// Diagnostic warning events that occurred during processing. These events are
	// non-critical errors found within the schema.
	Warnings []SchemaUploadUploadDetailsWarning `json:"warnings"`
	JSON     schemaUploadUploadDetailsJSON      `json:"-"`
}

// schemaUploadUploadDetailsJSON contains the JSON metadata for the struct
// [SchemaUploadUploadDetails]
type schemaUploadUploadDetailsJSON struct {
	Warnings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaUploadUploadDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaUploadUploadDetailsJSON) RawJSON() string {
	return r.raw
}

type SchemaUploadUploadDetailsWarning struct {
	// Code that identifies the event that occurred.
	Code int64 `json:"code,required"`
	// JSONPath location(s) in the schema where these events were encountered. See
	// [https://goessner.net/articles/JsonPath/](https://goessner.net/articles/JsonPath/)
	// for JSONPath specification.
	Locations []string `json:"locations"`
	// Diagnostic message that describes the event.
	Message string                               `json:"message"`
	JSON    schemaUploadUploadDetailsWarningJSON `json:"-"`
}

// schemaUploadUploadDetailsWarningJSON contains the JSON metadata for the struct
// [SchemaUploadUploadDetailsWarning]
type schemaUploadUploadDetailsWarningJSON struct {
	Code        apijson.Field
	Locations   apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaUploadUploadDetailsWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaUploadUploadDetailsWarningJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [api_gateway.UserSchemaDeleteResponseUnknown] or
// [shared.UnionString].
type UserSchemaDeleteResponseUnion interface {
	ImplementsAPIGatewayUserSchemaDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSchemaDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserSchemaNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Schema file bytes
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	// Kind of schema
	Kind param.Field[UserSchemaNewParamsKind] `json:"kind,required"`
	// Name of the schema
	Name param.Field[string] `json:"name"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[UserSchemaNewParamsValidationEnabled] `json:"validation_enabled"`
}

func (r UserSchemaNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type UserSchemaNewParamsKind string

const (
	UserSchemaNewParamsKindOpenAPIV3 UserSchemaNewParamsKind = "openapi_v3"
)

func (r UserSchemaNewParamsKind) IsKnown() bool {
	switch r {
	case UserSchemaNewParamsKindOpenAPIV3:
		return true
	}
	return false
}

// Flag whether schema is enabled for validation.
type UserSchemaNewParamsValidationEnabled string

const (
	UserSchemaNewParamsValidationEnabledTrue  UserSchemaNewParamsValidationEnabled = "true"
	UserSchemaNewParamsValidationEnabledFalse UserSchemaNewParamsValidationEnabled = "false"
)

func (r UserSchemaNewParamsValidationEnabled) IsKnown() bool {
	switch r {
	case UserSchemaNewParamsValidationEnabledTrue, UserSchemaNewParamsValidationEnabledFalse:
		return true
	}
	return false
}

type UserSchemaNewResponseEnvelope struct {
	Errors   Message      `json:"errors,required"`
	Messages Message      `json:"messages,required"`
	Result   SchemaUpload `json:"result,required"`
	// Whether the API call was successful
	Success bool                              `json:"success,required"`
	JSON    userSchemaNewResponseEnvelopeJSON `json:"-"`
}

// userSchemaNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserSchemaNewResponseEnvelope]
type userSchemaNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserSchemaListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Omit the source-files of schemas and only retrieve their meta-data.
	OmitSource param.Field[bool] `query:"omit_source"`
	// Page number of paginated results.
	Page param.Field[interface{}] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[interface{}] `query:"per_page"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[bool] `query:"validation_enabled"`
}

// URLQuery serializes [UserSchemaListParams]'s query parameters as `url.Values`.
func (r UserSchemaListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type UserSchemaDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type UserSchemaDeleteResponseEnvelope struct {
	Errors   Message                       `json:"errors,required"`
	Messages Message                       `json:"messages,required"`
	Result   UserSchemaDeleteResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success bool                                 `json:"success,required"`
	JSON    userSchemaDeleteResponseEnvelopeJSON `json:"-"`
}

// userSchemaDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserSchemaDeleteResponseEnvelope]
type userSchemaDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserSchemaEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[UserSchemaEditParamsValidationEnabled] `json:"validation_enabled"`
}

func (r UserSchemaEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Flag whether schema is enabled for validation.
type UserSchemaEditParamsValidationEnabled bool

const (
	UserSchemaEditParamsValidationEnabledTrue UserSchemaEditParamsValidationEnabled = true
)

func (r UserSchemaEditParamsValidationEnabled) IsKnown() bool {
	switch r {
	case UserSchemaEditParamsValidationEnabledTrue:
		return true
	}
	return false
}

type UserSchemaEditResponseEnvelope struct {
	Errors   Message      `json:"errors,required"`
	Messages Message      `json:"messages,required"`
	Result   PublicSchema `json:"result,required"`
	// Whether the API call was successful
	Success bool                               `json:"success,required"`
	JSON    userSchemaEditResponseEnvelopeJSON `json:"-"`
}

// userSchemaEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserSchemaEditResponseEnvelope]
type userSchemaEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserSchemaGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Omit the source-files of schemas and only retrieve their meta-data.
	OmitSource param.Field[bool] `query:"omit_source"`
}

// URLQuery serializes [UserSchemaGetParams]'s query parameters as `url.Values`.
func (r UserSchemaGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type UserSchemaGetResponseEnvelope struct {
	Errors   Message      `json:"errors,required"`
	Messages Message      `json:"messages,required"`
	Result   PublicSchema `json:"result,required"`
	// Whether the API call was successful
	Success bool                              `json:"success,required"`
	JSON    userSchemaGetResponseEnvelopeJSON `json:"-"`
}

// userSchemaGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserSchemaGetResponseEnvelope]
type userSchemaGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
