// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// CWSRegoMetadataService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCWSRegoMetadataService] method instead.
type CWSRegoMetadataService struct {
	Options []option.RequestOption
}

// NewCWSRegoMetadataService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCWSRegoMetadataService(opts ...option.RequestOption) (r *CWSRegoMetadataService) {
	r = &CWSRegoMetadataService{}
	r.Options = opts
	return
}

// Creates a JSON document that Rego can access through its import path.
func (r *CWSRegoMetadataService) New(ctx context.Context, params CWSRegoMetadataNewParams, opts ...option.RequestOption) (res *RegoMetadata, err error) {
	var env CWSRegoMetadataNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/rego-metadata", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces the metadata document's name, import path, and JSON data.
func (r *CWSRegoMetadataService) Update(ctx context.Context, regoMetadataID string, params CWSRegoMetadataUpdateParams, opts ...option.RequestOption) (res *RegoMetadata, err error) {
	var env CWSRegoMetadataUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if regoMetadataID == "" {
		err = errors.New("missing required rego_metadata_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/rego-metadata/%s", params.AccountID, regoMetadataID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of Rego metadata documents in the account.
func (r *CWSRegoMetadataService) List(ctx context.Context, params CWSRegoMetadataListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[RegoMetadata], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/rego-metadata", params.AccountID)
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

// Returns a paginated list of Rego metadata documents in the account.
func (r *CWSRegoMetadataService) ListAutoPaging(ctx context.Context, params CWSRegoMetadataListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[RegoMetadata] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes the specified Rego metadata document.
func (r *CWSRegoMetadataService) Delete(ctx context.Context, regoMetadataID string, body CWSRegoMetadataDeleteParams, opts ...option.RequestOption) (res *CWSRegoMetadataDeleteResponse, err error) {
	var env CWSRegoMetadataDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if regoMetadataID == "" {
		err = errors.New("missing required rego_metadata_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/rego-metadata/%s", body.AccountID, regoMetadataID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the Rego metadata document with the specified ID.
func (r *CWSRegoMetadataService) Get(ctx context.Context, regoMetadataID string, query CWSRegoMetadataGetParams, opts ...option.RequestOption) (res *RegoMetadata, err error) {
	var env CWSRegoMetadataGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if regoMetadataID == "" {
		err = errors.New("missing required rego_metadata_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/rego-metadata/%s", query.AccountID, regoMetadataID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type RegoMetadata struct {
	// Server-assigned UUID
	ID        string      `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time   `json:"created_at" api:"required" format:"date-time"`
	Data      interface{} `json:"data" api:"required"`
	// Unique Rego import path in the form metadata.<segment>, with one alphanumeric or
	// underscore segment. Reference it as data.<import_path>.
	ImportPath string `json:"import_path" api:"required"`
	// Human-readable name; unique per account
	Name      string           `json:"name" api:"required"`
	UpdatedAt time.Time        `json:"updated_at" api:"required" format:"date-time"`
	JSON      regoMetadataJSON `json:"-"`
}

// regoMetadataJSON contains the JSON metadata for the struct [RegoMetadata]
type regoMetadataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Data        apijson.Field
	ImportPath  apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegoMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regoMetadataJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataDeleteResponse struct {
	ID   string                            `json:"id" api:"required" format:"uuid"`
	JSON cwsRegoMetadataDeleteResponseJSON `json:"-"`
}

// cwsRegoMetadataDeleteResponseJSON contains the JSON metadata for the struct
// [CWSRegoMetadataDeleteResponse]
type cwsRegoMetadataDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataNewParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string]      `path:"account_id" api:"required"`
	Data      param.Field[interface{}] `json:"data" api:"required"`
	// Unique Rego import path in the form metadata.<segment>, with one alphanumeric or
	// underscore segment. Reference it as data.<import_path>.
	ImportPath param.Field[string] `json:"import_path" api:"required"`
	// Name of the rego metadata document (unique per account)
	Name param.Field[string] `json:"name" api:"required"`
}

func (r CWSRegoMetadataNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSRegoMetadataNewResponseEnvelope struct {
	Errors   []CWSRegoMetadataNewResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSRegoMetadataNewResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   RegoMetadata                                 `json:"result" api:"required"`
	Success  bool                                         `json:"success" api:"required"`
	JSON     cwsRegoMetadataNewResponseEnvelopeJSON       `json:"-"`
}

// cwsRegoMetadataNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSRegoMetadataNewResponseEnvelope]
type cwsRegoMetadataNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataNewResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	Source  CWSRegoMetadataNewResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsRegoMetadataNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsRegoMetadataNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSRegoMetadataNewResponseEnvelopeErrors]
type cwsRegoMetadataNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataNewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer" api:"required"`
	JSON    cwsRegoMetadataNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsRegoMetadataNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSRegoMetadataNewResponseEnvelopeErrorsSource]
type cwsRegoMetadataNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataNewResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code" api:"required"`
	Message string                                           `json:"message" api:"required"`
	Source  CWSRegoMetadataNewResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsRegoMetadataNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsRegoMetadataNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSRegoMetadataNewResponseEnvelopeMessages]
type cwsRegoMetadataNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataNewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer" api:"required"`
	JSON    cwsRegoMetadataNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsRegoMetadataNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSRegoMetadataNewResponseEnvelopeMessagesSource]
type cwsRegoMetadataNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataUpdateParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string]      `path:"account_id" api:"required"`
	Data      param.Field[interface{}] `json:"data" api:"required"`
	// Unique Rego import path in the form metadata.<segment>, with one alphanumeric or
	// underscore segment. Reference it as data.<import_path>.
	ImportPath param.Field[string] `json:"import_path" api:"required"`
	// Name of the rego metadata document (unique per account)
	Name param.Field[string] `json:"name" api:"required"`
}

func (r CWSRegoMetadataUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSRegoMetadataUpdateResponseEnvelope struct {
	Errors   []CWSRegoMetadataUpdateResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSRegoMetadataUpdateResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   RegoMetadata                                    `json:"result" api:"required"`
	Success  bool                                            `json:"success" api:"required"`
	JSON     cwsRegoMetadataUpdateResponseEnvelopeJSON       `json:"-"`
}

// cwsRegoMetadataUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSRegoMetadataUpdateResponseEnvelope]
type cwsRegoMetadataUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataUpdateResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	Source  CWSRegoMetadataUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsRegoMetadataUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsRegoMetadataUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CWSRegoMetadataUpdateResponseEnvelopeErrors]
type cwsRegoMetadataUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer" api:"required"`
	JSON    cwsRegoMetadataUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsRegoMetadataUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSRegoMetadataUpdateResponseEnvelopeErrorsSource]
type cwsRegoMetadataUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataUpdateResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code" api:"required"`
	Message string                                              `json:"message" api:"required"`
	Source  CWSRegoMetadataUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsRegoMetadataUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsRegoMetadataUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSRegoMetadataUpdateResponseEnvelopeMessages]
type cwsRegoMetadataUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer" api:"required"`
	JSON    cwsRegoMetadataUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsRegoMetadataUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CWSRegoMetadataUpdateResponseEnvelopeMessagesSource]
type cwsRegoMetadataUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataListParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Sort direction (default: asc)
	Direction param.Field[CWSRegoMetadataListParamsDirection] `query:"direction"`
	// Filter by import path (case-sensitive literal substring match)
	ImportPath param.Field[string] `query:"import_path"`
	// Filter by rego metadata name (case-insensitive substring match)
	Name param.Field[string] `query:"name"`
	// Field to order results by (default: created_at)
	Order param.Field[CWSRegoMetadataListParamsOrder] `query:"order"`
	// Number of items per page (1-100, default 20)
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [CWSRegoMetadataListParams]'s query parameters as
// `url.Values`.
func (r CWSRegoMetadataListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction (default: asc)
type CWSRegoMetadataListParamsDirection string

const (
	CWSRegoMetadataListParamsDirectionAsc  CWSRegoMetadataListParamsDirection = "asc"
	CWSRegoMetadataListParamsDirectionDesc CWSRegoMetadataListParamsDirection = "desc"
)

func (r CWSRegoMetadataListParamsDirection) IsKnown() bool {
	switch r {
	case CWSRegoMetadataListParamsDirectionAsc, CWSRegoMetadataListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by (default: created_at)
type CWSRegoMetadataListParamsOrder string

const (
	CWSRegoMetadataListParamsOrderName      CWSRegoMetadataListParamsOrder = "name"
	CWSRegoMetadataListParamsOrderCreatedAt CWSRegoMetadataListParamsOrder = "created_at"
	CWSRegoMetadataListParamsOrderUpdatedAt CWSRegoMetadataListParamsOrder = "updated_at"
)

func (r CWSRegoMetadataListParamsOrder) IsKnown() bool {
	switch r {
	case CWSRegoMetadataListParamsOrderName, CWSRegoMetadataListParamsOrderCreatedAt, CWSRegoMetadataListParamsOrderUpdatedAt:
		return true
	}
	return false
}

type CWSRegoMetadataDeleteParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSRegoMetadataDeleteResponseEnvelope struct {
	Errors   []CWSRegoMetadataDeleteResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSRegoMetadataDeleteResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   CWSRegoMetadataDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                            `json:"success" api:"required"`
	JSON     cwsRegoMetadataDeleteResponseEnvelopeJSON       `json:"-"`
}

// cwsRegoMetadataDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSRegoMetadataDeleteResponseEnvelope]
type cwsRegoMetadataDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataDeleteResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	Source  CWSRegoMetadataDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsRegoMetadataDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsRegoMetadataDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CWSRegoMetadataDeleteResponseEnvelopeErrors]
type cwsRegoMetadataDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer" api:"required"`
	JSON    cwsRegoMetadataDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsRegoMetadataDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSRegoMetadataDeleteResponseEnvelopeErrorsSource]
type cwsRegoMetadataDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataDeleteResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code" api:"required"`
	Message string                                              `json:"message" api:"required"`
	Source  CWSRegoMetadataDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsRegoMetadataDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsRegoMetadataDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSRegoMetadataDeleteResponseEnvelopeMessages]
type cwsRegoMetadataDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer" api:"required"`
	JSON    cwsRegoMetadataDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsRegoMetadataDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CWSRegoMetadataDeleteResponseEnvelopeMessagesSource]
type cwsRegoMetadataDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataGetParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSRegoMetadataGetResponseEnvelope struct {
	Errors   []CWSRegoMetadataGetResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSRegoMetadataGetResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   RegoMetadata                                 `json:"result" api:"required"`
	Success  bool                                         `json:"success" api:"required"`
	JSON     cwsRegoMetadataGetResponseEnvelopeJSON       `json:"-"`
}

// cwsRegoMetadataGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSRegoMetadataGetResponseEnvelope]
type cwsRegoMetadataGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	Source  CWSRegoMetadataGetResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsRegoMetadataGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsRegoMetadataGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSRegoMetadataGetResponseEnvelopeErrors]
type cwsRegoMetadataGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer" api:"required"`
	JSON    cwsRegoMetadataGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsRegoMetadataGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSRegoMetadataGetResponseEnvelopeErrorsSource]
type cwsRegoMetadataGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code" api:"required"`
	Message string                                           `json:"message" api:"required"`
	Source  CWSRegoMetadataGetResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsRegoMetadataGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsRegoMetadataGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSRegoMetadataGetResponseEnvelopeMessages]
type cwsRegoMetadataGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSRegoMetadataGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer" api:"required"`
	JSON    cwsRegoMetadataGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsRegoMetadataGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSRegoMetadataGetResponseEnvelopeMessagesSource]
type cwsRegoMetadataGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoMetadataGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoMetadataGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
